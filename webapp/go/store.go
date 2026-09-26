package main

import (
	"bufio"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// データはすべてメモリに持つ。MySQL は使わない。
//
// 永続化は「変更があれば 1 秒ごとにスナップショットを fsync まで書く」+「停止シグナルで同期的に書く」。
// 書き込みのたびにはディスクへ触らない。kill -9・OOM・電源断では最後のスナップショット以降（最大 1 秒）が失われる。
// 起動時はスナップショットがあれば読み、無ければ初期データ（../sql/3-initial-data.sql.gz）を読む。
// /api/initialize には依存しない。

// chairLocAgg は chair_locations の代わり。全件は持たず、API が必要とする集約だけを持つ。
// 椅子の登録時に作り（valid=false）、フィールドは行ロック mu で守る。
// 座標の送信は chairLocs の読みロック + この行ロックで書くので、別の椅子の座標の送信とは並列に走る。
type chairLocAgg struct {
	mu            sync.RWMutex
	valid         bool // 位置を 1 回でも受け取ったか
	latitude      int
	longitude     int
	totalDistance int
	updatedAt     time.Time // 最後の位置の created_at（= total_distance_updated_at）
}

// chairLocSnap はスナップショット用（ロックを含まない）。フィールド名は旧 chairLocAgg と同じなので、古いスナップショットも読める
type chairLocSnap struct {
	Latitude      int
	Longitude     int
	TotalDistance int
	UpdatedAt     time.Time
}

// get は位置の集約を読む。ok=false なら位置がまだ無い
func (a *chairLocAgg) get() (snap chairLocSnap, ok bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return chairLocSnap{Latitude: a.latitude, Longitude: a.longitude, TotalDistance: a.totalDistance, UpdatedAt: a.updatedAt}, a.valid
}

// add は位置を 1 件足す（累計距離は直前の位置からのマンハッタン距離）
func (a *chairLocAgg) add(lat, lon int, at time.Time) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.valid {
		a.totalDistance += abs(lat-a.latitude) + abs(lon-a.longitude)
	}
	a.valid = true
	a.latitude, a.longitude, a.updatedAt = lat, lon, at
}

type chairStats struct {
	Count         int
	EvaluationSum int
}

// Store は全テーブルを持つ。ロックはテーブルごと（locks.go の lockTables）。フィールドの横にどのテーブルに属するかを書く
type Store struct {
	paymentGatewayURL string // tSettings

	users             map[string]*User         // tUsers
	usersByToken      map[string]*User         // tUsers
	usersByInvitation map[string]*User         // tUsers
	usernames         map[string]struct{}      // tUsers
	owners            map[string]*Owner        // tOwners
	ownersByToken     map[string]*Owner        // tOwners
	ownersByRegister  map[string]*Owner        // tOwners
	ownerNames        map[string]struct{}      // tOwners
	chairs            map[string]*Chair        // tChairs（Chair.IsActive / UpdatedAt も）
	chairList         []*Chair                 // tChairs。id 順（= SELECT * FROM chairs の順）
	chairsByToken     map[string]*Chair        // tChairs
	chairsByOwner     map[string][]*Chair      // tChairs
	chairLoc          map[string]*chairLocAgg  // tChairLocs（map）。中身は行ロック
	chairStat         map[string]*chairStats   // tChairStats
	paymentTokens     map[string]*PaymentToken // tPaymentTokens
	rides             map[string]*Ride         // tRides（Ride.ChairID / UpdatedAt / Evaluation も）
	ridesByUser       map[string][]*Ride       // tRides。created_at 順
	ridesByChair      map[string][]*Ride       // tRides。割り当て順
	unmatched         []*Ride                  // tRides。chair_id IS NULL、created_at 順
	statuses          map[string][]*RideStatus // tRideStatuses。ride_id → created_at 順
	statusRow         map[string]*sync.Mutex   // tRideStatuses（map）。ride_id → 行ロック。送信済みの印（AppSentAt / ChairSentAt）を守る
	coupons           map[string][]*Coupon     // tCoupons。user_id → created_at 順（Coupon.UsedBy も）
	couponKeys        map[string]struct{}      // tCoupons。user_id + "\x00" + code
	couponCodeCount   map[string]int           // tCoupons。code → 件数（招待コードの上限判定）
	couponByRide      map[string]*Coupon       // tCoupons。used_by → coupon
}

var store = newStore()

// dirty はスナップショット以降に変更があったか。書き込み側はテーブルのロックを握った状態で立てる。
var dirty atomic.Bool

func newStore() *Store {
	return &Store{
		users:             map[string]*User{},
		usersByToken:      map[string]*User{},
		usersByInvitation: map[string]*User{},
		usernames:         map[string]struct{}{},
		owners:            map[string]*Owner{},
		ownersByToken:     map[string]*Owner{},
		ownersByRegister:  map[string]*Owner{},
		ownerNames:        map[string]struct{}{},
		chairs:            map[string]*Chair{},
		chairsByToken:     map[string]*Chair{},
		chairsByOwner:     map[string][]*Chair{},
		chairLoc:          map[string]*chairLocAgg{},
		chairStat:         map[string]*chairStats{},
		paymentTokens:     map[string]*PaymentToken{},
		rides:             map[string]*Ride{},
		ridesByUser:       map[string][]*Ride{},
		ridesByChair:      map[string][]*Ride{},
		statuses:          map[string][]*RideStatus{},
		statusRow:         map[string]*sync.Mutex{},
		coupons:           map[string][]*Coupon{},
		couponKeys:        map[string]struct{}{},
		couponCodeCount:   map[string]int{},
		couponByRide:      map[string]*Coupon{},
	}
}

// now は MySQL の CURRENT_TIMESTAMP(6) に合わせてマイクロ秒に丸める。
func now() time.Time {
	return time.Now().UTC().Truncate(time.Microsecond)
}

// ---------------------------------------------------------------------------
// 追加（インデックスの更新をここに集める）。呼び出し側が該当テーブルの書きロックを握る
// ---------------------------------------------------------------------------

func (s *Store) addUser(u *User) {
	s.users[u.ID] = u
	s.usersByToken[u.AccessToken] = u
	s.usersByInvitation[u.InvitationCode] = u
	s.usernames[u.Username] = struct{}{}
}

func (s *Store) addOwner(o *Owner) {
	s.owners[o.ID] = o
	s.ownersByToken[o.AccessToken] = o
	s.ownersByRegister[o.ChairRegisterToken] = o
	s.ownerNames[o.Name] = struct{}{}
}

// addChair は tChairs と tChairLocs の書きロックで呼ぶ（位置の集約も作る）
func (s *Store) addChair(c *Chair) {
	if s.chairLoc[c.ID] == nil {
		s.chairLoc[c.ID] = &chairLocAgg{}
	}
	s.chairs[c.ID] = c
	s.chairList = append(s.chairList, c)
	s.chairsByToken[c.AccessToken] = c
	s.chairsByOwner[c.OwnerID] = append(s.chairsByOwner[c.OwnerID], c)
}

func (s *Store) addRide(r *Ride) {
	s.rides[r.ID] = r
	s.ridesByUser[r.UserID] = append(s.ridesByUser[r.UserID], r)
	if r.ChairID.Valid {
		s.ridesByChair[r.ChairID.String] = append(s.ridesByChair[r.ChairID.String], r)
	} else {
		s.unmatched = append(s.unmatched, r)
	}
}

// addStatus は tRideStatuses の書きロックと、tRides の読みロック以上で呼ぶ
func (s *Store) addStatus(st *RideStatus) {
	if s.statusRow[st.RideID] == nil {
		s.statusRow[st.RideID] = &sync.Mutex{}
	}
	s.statuses[st.RideID] = append(s.statuses[st.RideID], st)
	if ride := s.rides[st.RideID]; ride != nil {
		s.wakeRide(ride) // 通知の SSE を起こす
	}
}

func couponKey(userID, code string) string { return userID + "\x00" + code }

func (s *Store) addCoupon(c *Coupon) {
	s.coupons[c.UserID] = append(s.coupons[c.UserID], c)
	s.couponKeys[couponKey(c.UserID, c.Code)] = struct{}{}
	s.couponCodeCount[c.Code]++
	if c.UsedBy != nil {
		s.couponByRide[*c.UsedBy] = c
	}
}

// ---------------------------------------------------------------------------
// 参照
// ---------------------------------------------------------------------------

// latestStatus は SELECT status FROM ride_statuses WHERE ride_id = ? ORDER BY created_at DESC LIMIT 1。
func (s *Store) latestStatus(rideID string) (string, bool) {
	sts := s.statuses[rideID]
	if len(sts) == 0 {
		return "", false
	}
	return sts[len(sts)-1].Status, true
}

// latestRideOfChair は SELECT * FROM rides WHERE chair_id = ? ORDER BY updated_at DESC LIMIT 1。
// 椅子には前のライドが完了してから次が割り当てられるので、最後に割り当てたものが最新。
func (s *Store) latestRideOfChair(chairID string) *Ride {
	rs := s.ridesByChair[chairID]
	if len(rs) == 0 {
		return nil
	}
	return rs[len(rs)-1]
}

// chairIsFree は元の internalGetMatching の空き判定（ライドの 6 状態すべてを椅子に通知済み）。
// 前のライドが完了してから次を割り当てるので、最新のライドだけを見れば足りる。
func (s *Store) chairIsFree(chairID string) bool {
	r := s.latestRideOfChair(chairID)
	if r == nil {
		return true
	}
	row := s.statusRow[r.ID]
	if row == nil {
		return false
	}
	row.Lock()
	defer row.Unlock()
	sent := 0
	for _, st := range s.statuses[r.ID] {
		if st.ChairSentAt != nil {
			sent++
		}
	}
	return sent == 6
}

// ---------------------------------------------------------------------------
// スナップショット
// ---------------------------------------------------------------------------

type snapshot struct {
	PaymentGatewayURL string
	Users             []User
	Owners            []Owner
	Chairs            []Chair
	ChairLoc          map[string]chairLocSnap
	PaymentTokens     []PaymentToken
	Rides             []Ride
	Statuses          []RideStatus
	Coupons           []Coupon
}

// capture は全テーブルの読みロック以上を握った状態で呼ぶ。値をコピーするだけで、エンコードはロックの外でやる。
func (s *Store) capture() *snapshot {
	snap := &snapshot{
		PaymentGatewayURL: s.paymentGatewayURL,
		Users:             make([]User, 0, len(s.users)),
		Owners:            make([]Owner, 0, len(s.owners)),
		Chairs:            make([]Chair, 0, len(s.chairList)),
		ChairLoc:          make(map[string]chairLocSnap, len(s.chairLoc)),
		PaymentTokens:     make([]PaymentToken, 0, len(s.paymentTokens)),
		Rides:             make([]Ride, 0, len(s.rides)),
	}
	for _, u := range s.users {
		snap.Users = append(snap.Users, *u)
	}
	for _, o := range s.owners {
		snap.Owners = append(snap.Owners, *o)
	}
	for _, c := range s.chairList {
		snap.Chairs = append(snap.Chairs, *c)
	}
	for id, l := range s.chairLoc {
		if v, ok := l.get(); ok {
			snap.ChairLoc[id] = v
		}
	}
	for _, p := range s.paymentTokens {
		snap.PaymentTokens = append(snap.PaymentTokens, *p)
	}
	// ライドはユーザーごとに created_at 順で持っているので、その順で出せば復元時も順序が保たれる。
	// 椅子ごとの割り当て順は updated_at では復元できないので、Rides の順（= created_at）で近似する。
	for _, rs := range s.ridesByUser {
		for _, r := range rs {
			snap.Rides = append(snap.Rides, *r)
			if row := s.statusRow[r.ID]; row != nil {
				row.Lock()
				for _, st := range s.statuses[r.ID] {
					snap.Statuses = append(snap.Statuses, copyStatus(st))
				}
				row.Unlock()
			}
		}
	}
	for _, cs := range s.coupons {
		for _, c := range cs {
			cc := *c
			if c.UsedBy != nil {
				v := *c.UsedBy
				cc.UsedBy = &v
			}
			snap.Coupons = append(snap.Coupons, cc)
		}
	}
	return snap
}

func copyStatus(st *RideStatus) RideStatus {
	c := *st
	if st.AppSentAt != nil {
		v := *st.AppSentAt
		c.AppSentAt = &v
	}
	if st.ChairSentAt != nil {
		v := *st.ChairSentAt
		c.ChairSentAt = &v
	}
	return c
}

// restore はスナップショットから Store を組み立てる（インデックスは作り直す）。
func restore(snap *snapshot) *Store {
	s := newStore()
	s.paymentGatewayURL = snap.PaymentGatewayURL
	for i := range snap.Users {
		s.addUser(&snap.Users[i])
	}
	for i := range snap.Owners {
		s.addOwner(&snap.Owners[i])
	}
	sortByID(snap.Chairs)
	for i := range snap.Chairs {
		s.addChair(&snap.Chairs[i])
	}
	for id, l := range snap.ChairLoc {
		agg := s.chairLoc[id]
		if agg == nil {
			agg = &chairLocAgg{}
			s.chairLoc[id] = agg
		}
		agg.valid = true
		agg.latitude, agg.longitude, agg.totalDistance, agg.updatedAt = l.Latitude, l.Longitude, l.TotalDistance, l.UpdatedAt
	}
	for i := range snap.PaymentTokens {
		s.paymentTokens[snap.PaymentTokens[i].UserID] = &snap.PaymentTokens[i]
	}
	sortRides(snap.Rides)
	for i := range snap.Rides {
		s.addRide(&snap.Rides[i])
	}
	sortStatuses(snap.Statuses)
	for i := range snap.Statuses {
		s.addStatus(&snap.Statuses[i])
	}
	sortCoupons(snap.Coupons)
	for i := range snap.Coupons {
		s.addCoupon(&snap.Coupons[i])
	}
	s.rebuildChairStats()
	return s
}

// rebuildChairStats は getChairStats の集計（ARRIVED・CARRYING・COMPLETED が揃ったライドの件数と評価の合計）を作り直す。
func (s *Store) rebuildChairStats() {
	s.chairStat = map[string]*chairStats{}
	for chairID, rs := range s.ridesByChair {
		st := &chairStats{}
		for _, r := range rs {
			var arrived, carrying, completed bool
			for _, x := range s.statuses[r.ID] {
				switch x.Status {
				case "ARRIVED":
					arrived = true
				case "CARRYING":
					carrying = true
				case "COMPLETED":
					completed = true
				}
			}
			if arrived && carrying && completed && r.Evaluation != nil {
				st.Count++
				st.EvaluationSum += *r.Evaluation
			}
		}
		s.chairStat[chairID] = st
	}
}

func snapshotPath() string {
	if p := os.Getenv("ISUCON_SNAPSHOT_PATH"); p != "" {
		return p
	}
	return "/home/isucon/isuride-snapshot.gob"
}

var saveMu sync.Mutex // スナップショットの書き出しを直列にする

func saveSnapshot() error {
	saveMu.Lock()
	defer saveMu.Unlock()

	if !dirty.Load() {
		return nil
	}
	unlock := lockTables(nil, allTables)
	snap := store.capture()
	dirty.Store(false)
	unlock()

	if err := writeSnapshot(snap); err != nil {
		dirty.Store(true) // 次の周期で書き直す
		return err
	}
	return nil
}

func writeSnapshot(snap *snapshot) error {
	path := snapshotPath()
	tmp := fmt.Sprintf("%s.tmp", path)
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	w := bufio.NewWriterSize(f, 1<<20)
	if err := gob.NewEncoder(w).Encode(snap); err != nil {
		f.Close()
		return err
	}
	if err := w.Flush(); err != nil {
		f.Close()
		return err
	}
	// innodb_flush_log_at_trx_commit=0 と同じ水準（1 秒ごとに fsync まで）にする。
	// カーネルの書き戻し（最大 30 秒ほど遅れる）に任せると、ハードリブートで失う幅が広がる
	if err := f.Sync(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		return err
	}
	// rename 自体をディスクに確定させる
	dir, err := os.Open(filepath.Dir(path))
	if err != nil {
		return err
	}
	defer dir.Close()
	return dir.Sync()
}

func loadSnapshot() (*snapshot, error) {
	f, err := os.Open(snapshotPath())
	if err != nil {
		return nil, err
	}
	defer f.Close()
	snap := &snapshot{}
	if err := gob.NewDecoder(bufio.NewReaderSize(f, 1<<20)).Decode(snap); err != nil {
		return nil, err
	}
	return snap, nil
}

// initStore は起動時に呼ぶ。スナップショットがあればそれを、無ければ初期データを読む。
func initStore() {
	if err := loadChairSpeeds(); err != nil {
		panic(err)
	}
	snap, err := loadSnapshot()
	if err == nil {
		store = restore(snap)
		slog.Info("snapshot loaded", "path", snapshotPath(), "rides", len(snap.Rides))
	} else {
		if !errors.Is(err, os.ErrNotExist) {
			slog.Error("failed to load snapshot; falling back to initial data", "err", err)
		}
		s, err := loadInitialStore()
		if err != nil {
			panic(err)
		}
		store = s
		dirty.Store(true)
	}

	go func() {
		t := time.NewTicker(time.Second)
		for range t.C {
			if err := saveSnapshot(); err != nil {
				slog.Error("failed to save snapshot", "err", err)
			}
		}
	}()

	// systemd の ExecStop は QUIT を送る。reboot では TERM。どちらでも書いてから落ちる。
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)
	go func() {
		s := <-sig
		// 以降の書き込みを止めてから書く（全テーブルの書きロックを取り、放さずに終了する）
		lockTables(allTables, nil)
		snap := store.capture()
		if err := writeSnapshot(snap); err != nil {
			slog.Error("failed to save snapshot on signal", "signal", s, "err", err)
			os.Exit(1)
		}
		slog.Info("snapshot saved on signal", "signal", s, "path", snapshotPath())
		os.Exit(0)
	}()
}

// replaceWith は /api/initialize 用。ポインタを差し替えると並行するハンドラと競合するので、ロック下で中身を入れ替える。
func (s *Store) replaceWith(n *Store) {
	unlock := lockTables(allTables, nil)
	defer unlock()
	s.paymentGatewayURL = n.paymentGatewayURL
	s.users, s.usersByToken, s.usersByInvitation, s.usernames = n.users, n.usersByToken, n.usersByInvitation, n.usernames
	s.owners, s.ownersByToken, s.ownersByRegister, s.ownerNames = n.owners, n.ownersByToken, n.ownersByRegister, n.ownerNames
	s.chairs, s.chairList, s.chairsByToken, s.chairsByOwner = n.chairs, n.chairList, n.chairsByToken, n.chairsByOwner
	s.chairLoc, s.chairStat, s.paymentTokens = n.chairLoc, n.chairStat, n.paymentTokens
	s.rides, s.ridesByUser, s.ridesByChair, s.unmatched, s.statuses, s.statusRow = n.rides, n.ridesByUser, n.ridesByChair, n.unmatched, n.statuses, n.statusRow
	s.coupons, s.couponKeys, s.couponCodeCount, s.couponByRide = n.coupons, n.couponKeys, n.couponCodeCount, n.couponByRide
	dirty.Store(true)
}

func gobEncode(w io.Writer, v any) error { return gob.NewEncoder(w).Encode(v) }
func gobDecode(r io.Reader, v any) error { return gob.NewDecoder(r).Decode(v) }
