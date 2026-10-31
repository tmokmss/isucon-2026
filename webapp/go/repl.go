package isureal

// asia1 のストアへの書き込みを europe1・america1 の複製へ流す。
// - asia1 は s.mu の書きロックの中で書き込みを 1 件の変更（replOp）として当て、同じロックの中で通し番号と asia1 の時刻 T を振って各台へ送る。
//   書き込みの応答は複製の返事を待たない。代わりに replHeartbeat ごとに時刻だけの印を流す
// - 複製側は、読み取りを受けた時刻（+ replClockMargin）以上の T の印まで当て終えてから手元で返す（replica.waitCaughtUp）。
//   書き込みの応答は、その変更を流した後に asia1 を出る。読み取りはその後に複製の台へ着くので、着いた時刻より後の T の印より前に、その変更は流れている。
//   台の時計は Amazon Time Sync に合っている（chronyc tracking で 1ms 未満）。replReadWait 以内に追いつかなければ asia1 へ転送する
// - 複製側は購読を先に始めて変更を貯め、PostgreSQL から全部を読んだ後に当て直す（変更は ID で上書きするので、読み込みに入っていても同じ結果になる）。
//   貯めた分を当て終えたら ready を送り、asia1 が流れの中に置く live の印を受けた後から手元で返す
// - 複製側は replLease を超えて何も受けなければ手元で返さない（asia1 は replHeartbeat ごとに送る）。切れたら全部読み直す
// - initialize は asia1 が自分の読み直しの後に reset を流し、各台が小さい表を読み直した返事を待つ

import (
	"bufio"
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	replHeartbeat    = 5 * time.Millisecond
	replClockMargin  = 2 * time.Millisecond
	replReadWait     = 300 * time.Millisecond
	replLease        = time.Second
	replLeaseMargin  = 500 * time.Millisecond
	replAckTimeout   = 3 * time.Second // リージョン間の帯域が画像の転送で埋まると返事が 1 秒を超える
	replResetTimeout = 40 * time.Second
	replMaxPending   = 256 << 20
)

type replOp interface{ apply(s *Store) }

type replMsg struct {
	Seq uint64
	T   int64 // asia1 で流した時刻（UnixNano）。これより前に asia1 を出た書き込みの応答の変更は、すべてこれより前に流れている
	Op  replOp
}

// ---- 変更の種類。apply は s.mu の書きロックを持って呼ぶ ----

type opUserAdd struct{ U *UserRow }

func (o *opUserAdd) apply(s *Store) {
	if s.users[o.U.ID] != nil {
		return
	}
	o.U.setNormalized()
	s.addUserLocked(o.U)
}

type opLogin struct {
	UserID string
	Row    *LoginSessionRow
}

func (o *opLogin) apply(s *Store) {
	u := s.users[o.UserID]
	if u == nil {
		return
	}
	for hash, old := range u.sessions {
		old.Active = false
		delete(s.sessions, hash)
	}
	u.sessions = map[string]*LoginSessionRow{o.Row.TokenHash: o.Row}
	s.sessions[o.Row.TokenHash] = o.Row
}

type opLogout struct{ UserID, TokenHash string }

func (o *opLogout) apply(s *Store) {
	if row := s.sessions[o.TokenHash]; row != nil && row.UserID == o.UserID {
		row.Active = false
		delete(s.sessions, o.TokenHash)
		if u := s.users[o.UserID]; u != nil {
			delete(u.sessions, o.TokenHash)
		}
	}
}

type opUserDelete struct{ UserID string }

func (o *opUserDelete) apply(s *Store) {
	u := s.users[o.UserID]
	if u == nil {
		return
	}
	u.Status = "deleted"
	for hash, sess := range u.sessions {
		sess.Active = false
		delete(s.sessions, hash)
	}
	u.sessions = map[string]*LoginSessionRow{}
}

type opCertify struct{ UserID string }

func (o *opCertify) apply(s *Store) {
	if u := s.users[o.UserID]; u != nil {
		u.AccountType = "certified_public"
	}
}

// gob は空文字列を指すポインタを nil に戻すので、省略できる文字列は optStr で送る
type optStr struct {
	Set bool
	V   string
}

func opt(p *string) optStr {
	if p == nil {
		return optStr{}
	}
	return optStr{Set: true, V: *p}
}

type opUpdateMe struct {
	UserID                   string
	DisplayName, Bio, Region optStr
}

func (o *opUpdateMe) apply(s *Store) {
	u := s.users[o.UserID]
	if u == nil {
		return
	}
	if o.DisplayName.Set {
		u.DisplayName = o.DisplayName.V
		u.normDisplayName = normalizeSearchText(u.DisplayName)
	}
	if o.Bio.Set {
		u.Bio = o.Bio.V
	}
	if o.Region.Set {
		u.Region = o.Region.V
	}
}

// opProfileImage は ImageID が nil なら削除
type opProfileImage struct {
	UserID    string
	ImageID   *string
	CreatedAt *time.Time
}

func (o *opProfileImage) apply(s *Store) {
	u := s.users[o.UserID]
	if u == nil {
		return
	}
	if u.ProfileImageID != nil {
		delete(s.profileByImage, *u.ProfileImageID)
	}
	u.ProfileImageID = o.ImageID
	u.ProfileImageCreatedAt = o.CreatedAt
	if o.ImageID != nil {
		s.profileByImage[*o.ImageID] = u
	}
}

type opWindowOpen struct {
	W        *WindowRow
	Notified map[string]notificationRow
}

// 読み込んだ後に当て直すときは、読み込みに入っていた新しい投稿期間を終わらせない
func (o *opWindowOpen) apply(s *Store) {
	region := o.W.Region
	w := s.windows[o.W.ID]
	if w == nil {
		w = o.W
		if w.notified == nil {
			w.notified = o.Notified
			if w.notified == nil {
				w.notified = map[string]notificationRow{}
			}
		}
		s.windows[w.ID] = w
	}
	if w.EndedAt != nil {
		return
	}
	if old := s.activeWindows[region]; old != nil && old != w {
		if old.StartsAt.After(w.StartsAt) {
			return
		}
		if old.EndedAt == nil {
			t := w.StartsAt
			old.EndedAt = &t
			// 終わった投稿期間の通知はもう読まれない
			old.notified = map[string]notificationRow{}
		}
	}
	s.activeWindows[region] = w
	if d, ok := s.maxLocalDate[region]; !ok || w.LocalDate.After(d) {
		s.maxLocalDate[region] = w.LocalDate
	}
	s.broadcastLocked()
}

// gob は 0 や空文字列を指すポインタを nil に戻すので、位置のポインタが nil でないことを別に送る
type opPostAdd struct {
	P                     *PostRow
	HasLat, HasLng, HasLo bool
}

func newOpPostAdd(p *PostRow) *opPostAdd {
	return &opPostAdd{P: p, HasLat: p.Location.Latitude != nil, HasLng: p.Location.Longitude != nil, HasLo: p.Location.Locality != nil}
}

func (o *opPostAdd) apply(s *Store) {
	if s.posts[o.P.ID] != nil {
		return
	}
	l := &o.P.Location
	if o.HasLat && l.Latitude == nil {
		l.Latitude = new(float64)
	}
	if o.HasLng && l.Longitude == nil {
		l.Longitude = new(float64)
	}
	if o.HasLo && l.Locality == nil {
		l.Locality = new(string)
	}
	s.addPostLocked(o.P)
}

type opPostUpdate struct {
	ID                  string
	Caption, Visibility optStr
}

func (o *opPostUpdate) apply(s *Store) {
	p := s.posts[o.ID]
	if p == nil {
		return
	}
	if o.Caption.Set {
		p.Caption = o.Caption.V
	}
	if o.Visibility.Set {
		p.Visibility = o.Visibility.V
	}
}

type opPostDelete struct {
	ID string
	At time.Time
}

func (o *opPostDelete) apply(s *Store) {
	if p := s.posts[o.ID]; p != nil {
		t := o.At
		p.DeletedAt = &t
	}
}

type opCommentAdd struct{ C *CommentRow }

func (o *opCommentAdd) apply(s *Store) {
	if s.comments[o.C.ID] != nil || s.posts[o.C.PostID] == nil {
		return
	}
	s.addCommentLocked(o.C)
}

type opCommentDelete struct {
	ID string
	At time.Time
}

func (o *opCommentDelete) apply(s *Store) {
	if cm := s.comments[o.ID]; cm != nil {
		t := o.At
		cm.DeletedAt = &t
		delete(s.comments, o.ID)
	}
}

type opRealMojiAdd struct {
	R        *RealMojiRow
	Replaced []string
}

func (o *opRealMojiAdd) apply(s *Store) {
	for _, id := range o.Replaced {
		if old := s.realmojis[id]; old != nil {
			t := o.R.CreatedAt
			old.ReplacedAt = &t
		}
	}
	if s.realmojis[o.R.ID] != nil || s.posts[o.R.PostID] == nil {
		return
	}
	s.addRealMojiLocked(o.R)
}

type opRealMojiDelete struct {
	ID string
	At time.Time
}

func (o *opRealMojiDelete) apply(s *Store) {
	if r := s.realmojis[o.ID]; r != nil {
		t := o.At
		r.DeletedAt = &t
	}
}

type opFriendRequestAdd struct{ R *FriendRequest }

func (o *opFriendRequestAdd) apply(s *Store) {
	if s.friendRequests[o.R.ID] != nil {
		return
	}
	s.friendRequests[o.R.ID] = o.R
	s.users[o.R.RequesterID].pendingOut[o.R.RecipientID] = o.R
	s.users[o.R.RecipientID].pendingIn[o.R.RequesterID] = o.R
}

// opRequestStatus は申請の状態を変え、F があれば友達にする
type opRequestStatus struct {
	ID     string
	Status string
	At     time.Time
	F      *FriendshipRow
}

func (o *opRequestStatus) apply(s *Store) {
	if r := s.friendRequests[o.ID]; r != nil {
		s.setRequestStatusLocked(r, o.Status, o.At)
	}
	if o.F != nil {
		if cur := s.users[o.F.LowID].friends[o.F.HighID]; cur != nil && cur.ID == o.F.ID {
			return
		}
		s.removeFriendshipLocked(o.F.LowID, o.F.HighID)
		s.addFriendshipLocked(o.F)
	}
}

type opFriendshipRemove struct{ A, B string }

func (o *opFriendshipRemove) apply(s *Store) { s.removeFriendshipLocked(o.A, o.B) }

type opFollowAdd struct{ F *FollowRow }

func (o *opFollowAdd) apply(s *Store) {
	if cur := s.users[o.F.FollowerID].following[o.F.FolloweeID]; cur != nil && cur.ID == o.F.ID {
		return
	}
	s.removeFollowLocked(o.F.FollowerID, o.F.FolloweeID)
	s.addFollowLocked(o.F)
}

type opFollowRemove struct{ Follower, Followee string }

func (o *opFollowRemove) apply(s *Store) { s.removeFollowLocked(o.Follower, o.Followee) }

// Cancel は取り消す保留中の申請。当て直しのときに、読み込みに入っていた後の申請まで取り消さないよう asia1 で決める
type opBlockAdd struct {
	B      *BlockRow
	Cancel []string
}

func newOpBlockAddLocked(s *Store, b *BlockRow) *opBlockAdd {
	op := &opBlockAdd{B: b}
	u := s.users[b.BlockerID]
	for _, r := range []*FriendRequest{u.pendingOut[b.BlockedID], u.pendingIn[b.BlockedID]} {
		if r != nil {
			op.Cancel = append(op.Cancel, r.ID)
		}
	}
	return op
}

func (o *opBlockAdd) apply(s *Store) {
	actor, target, at := o.B.BlockerID, o.B.BlockedID, o.B.CreatedAt
	if cur := s.users[actor].blockOut[target]; cur == nil || cur.ID != o.B.ID {
		if cur != nil {
			s.removeBlockLocked(actor, target)
		}
		s.blocks = append(s.blocks, o.B)
		s.users[actor].blockOut[target] = o.B
		s.users[target].blockIn[actor] = o.B
	}
	s.removeFriendshipLocked(actor, target)
	s.removeFollowLocked(actor, target)
	s.removeFollowLocked(target, actor)
	for _, id := range o.Cancel {
		if r := s.friendRequests[id]; r != nil && r.Status == "pending" {
			s.setRequestStatusLocked(r, "cancelled", at)
		}
	}
}

type opBlockRemove struct{ Blocker, Blocked string }

func (o *opBlockRemove) apply(s *Store) { s.removeBlockLocked(o.Blocker, o.Blocked) }

func (s *Store) removeBlockLocked(blocker, blocked string) {
	b := s.users[blocker].blockOut[blocked]
	if b == nil {
		return
	}
	delete(s.users[blocker].blockOut, blocked)
	delete(s.users[blocked].blockIn, blocker)
	for i, x := range s.blocks {
		if x == b {
			s.blocks[i] = s.blocks[len(s.blocks)-1]
			s.blocks = s.blocks[:len(s.blocks)-1]
			break
		}
	}
}

// 以下は流れの制御に使い、ストアには当てない
type opReset struct{ X bool } // gob は公開フィールドの無い型を送れない

func (o *opReset) apply(s *Store) {}

type opHeartbeat struct{ X bool } // gob は公開フィールドの無い型を送れない

func (o *opHeartbeat) apply(s *Store) {}

type opLive struct{ X bool } // gob は公開フィールドの無い型を送れない

func (o *opLive) apply(s *Store) {}

func init() {
	for _, op := range []replOp{
		&opUserAdd{}, &opLogin{}, &opLogout{}, &opUserDelete{}, &opCertify{}, &opUpdateMe{}, &opProfileImage{},
		&opWindowOpen{}, &opPostAdd{}, &opPostUpdate{}, &opPostDelete{}, &opCommentAdd{}, &opCommentDelete{},
		&opRealMojiAdd{}, &opRealMojiDelete{}, &opFriendRequestAdd{}, &opRequestStatus{}, &opFriendshipRemove{},
		&opFollowAdd{}, &opFollowRemove{}, &opBlockAdd{}, &opBlockRemove{}, &opReset{}, &opHeartbeat{}, &opLive{},
	} {
		gob.Register(op)
	}
}

// ---- asia1: 変更を流す側 ----

type replSub struct {
	addr    string
	conn    net.Conn
	mu      sync.Mutex
	cond    *sync.Cond
	pending bytes.Buffer
	enc     *gob.Encoder
	live    bool
	closed  bool
	acked   uint64
	ackCh   chan struct{}
	dead    chan struct{}
}

type replHub struct {
	mu          sync.Mutex
	seq         uint64
	subs        map[*replSub]struct{}
	quietUntil  time.Time
	waitedTotal atomic.Int64
	waitedCount atomic.Int64
}

var repl = &replHub{subs: map[*replSub]struct{}{}}

type replWait struct {
	seq       uint64
	subs      []*replSub
	notBefore time.Time
	timeout   time.Duration
}

// commitLocked は op をこのストアに当て、各台へ流す。s.mu の書きロックを持って呼び、ロックを外した後に返り値の wait を呼ぶ
func (s *Store) commitLocked(op replOp) replWait {
	op.apply(s)
	return repl.publish(op, 0)
}

func (hub *replHub) publish(op replOp, timeout time.Duration) replWait {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	hub.seq++
	w := replWait{seq: hub.seq, timeout: timeout}
	if timeout > 0 {
		w.notBefore = hub.quietUntil
	}
	m := replMsg{Seq: hub.seq, T: time.Now().UnixNano(), Op: op}
	for sub := range hub.subs {
		if sub.send(m) && sub.live && timeout > 0 {
			w.subs = append(w.subs, sub)
		}
	}
	return w
}

// send は hub.mu を持って呼ぶ。送れなければ（詰まった・切れた）false
func (sub *replSub) send(m replMsg) bool {
	sub.mu.Lock()
	defer sub.mu.Unlock()
	if sub.closed {
		return false
	}
	if sub.pending.Len() > replMaxPending {
		log.Printf("repl: %s is too far behind", sub.addr)
		go repl.drop(sub)
		return false
	}
	if err := sub.enc.Encode(&m); err != nil {
		log.Printf("repl: encode: %v", err)
		go repl.drop(sub)
		return false
	}
	sub.cond.Signal()
	return true
}

func (w replWait) wait() {
	if len(w.subs) > 0 {
		start := time.Now()
		deadline := start.Add(w.timeout)
		for _, sub := range w.subs {
			if !sub.waitAck(w.seq, deadline) {
				log.Printf("repl: no ack from %s for seq %d", sub.addr, w.seq)
				repl.drop(sub)
			}
		}
		repl.waitedTotal.Add(int64(time.Since(start)))
		repl.waitedCount.Add(1)
	}
	repl.mu.Lock()
	until := repl.quietUntil
	repl.mu.Unlock()
	if w.notBefore.After(until) {
		until = w.notBefore
	}
	if d := time.Until(until); d > 0 {
		time.Sleep(d)
	}
}

func (sub *replSub) waitAck(seq uint64, deadline time.Time) bool {
	for {
		sub.mu.Lock()
		if sub.acked >= seq {
			sub.mu.Unlock()
			return true
		}
		ch, dead := sub.ackCh, sub.dead
		sub.mu.Unlock()
		d := time.Until(deadline)
		if d <= 0 {
			return false
		}
		t := time.NewTimer(d)
		select {
		case <-ch:
			t.Stop()
		case <-dead:
			t.Stop()
			return false
		case <-t.C:
			return false
		}
	}
}

// drop はその台を切る。その台が手元の判定をやめるまで、以後の書き込みの応答を待たせる
func (hub *replHub) drop(sub *replSub) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	sub.mu.Lock()
	if sub.closed {
		sub.mu.Unlock()
		return
	}
	sub.closed = true
	wasLive := sub.live
	close(sub.dead)
	sub.cond.Broadcast()
	sub.mu.Unlock()
	sub.conn.Close()
	delete(hub.subs, sub)
	if wasLive {
		if t := time.Now().Add(replLease + replLeaseMargin); t.After(hub.quietUntil) {
			hub.quietUntil = t
		}
	}
	log.Printf("repl: dropped %s (live=%v)", sub.addr, wasLive)
}

func (hub *replHub) serveHTTP(w http.ResponseWriter, r *http.Request) {
	host, ok := privateRemote(r)
	if !ok {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	hj, ok := w.(http.Hijacker)
	if !ok {
		return
	}
	conn, brw, err := hj.Hijack()
	if err != nil {
		return
	}
	if _, err := conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\n\r\n")); err != nil {
		conn.Close()
		return
	}
	sub := &replSub{addr: host, conn: conn, ackCh: make(chan struct{}), dead: make(chan struct{})}
	sub.cond = sync.NewCond(&sub.mu)
	sub.enc = gob.NewEncoder(&sub.pending)
	hub.mu.Lock()
	hub.subs[sub] = struct{}{}
	hub.mu.Unlock()
	log.Printf("repl: %s joined", host)
	go sub.writeLoop()
	sub.readLoop(brw.Reader)
}

func (sub *replSub) writeLoop() {
	for {
		sub.mu.Lock()
		for sub.pending.Len() == 0 && !sub.closed {
			sub.cond.Wait()
		}
		if sub.closed {
			sub.mu.Unlock()
			return
		}
		data := bytes.Clone(sub.pending.Bytes())
		sub.pending.Reset()
		sub.mu.Unlock()
		sub.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
		if _, err := sub.conn.Write(data); err != nil {
			log.Printf("repl: write to %s: %v", sub.addr, err)
			repl.drop(sub)
			return
		}
	}
}

// readLoop は複製側からの "a <seq>"（当て終えた番号）と "ready" を読む
func (sub *replSub) readLoop(br *bufio.Reader) {
	defer repl.drop(sub)
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			return
		}
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "a "):
			n, err := strconv.ParseUint(line[2:], 10, 64)
			if err != nil {
				return
			}
			sub.mu.Lock()
			if n > sub.acked {
				sub.acked = n
				close(sub.ackCh)
				sub.ackCh = make(chan struct{})
			}
			sub.mu.Unlock()
		case line == "ready":
			repl.mu.Lock()
			if sub.send(replMsg{Seq: repl.seq, T: time.Now().UnixNano(), Op: &opLive{}}) {
				sub.mu.Lock()
				sub.live = true
				sub.mu.Unlock()
				log.Printf("repl: %s is live at seq %d", sub.addr, repl.seq)
			}
			repl.mu.Unlock()
		}
	}
}

func (hub *replHub) heartbeatLoop() {
	t := time.NewTicker(replHeartbeat)
	last := time.Now()
	for range t.C {
		hub.mu.Lock()
		for sub := range hub.subs {
			sub.send(replMsg{Seq: hub.seq, T: time.Now().UnixNano(), Op: &opHeartbeat{}})
		}
		hub.mu.Unlock()
		if time.Since(last) > 10*time.Second {
			last = time.Now()
			if n := hub.waitedCount.Swap(0); n > 0 {
				log.Printf("repl: %d writes waited %s on average for acks in the last 10s", n, time.Duration(hub.waitedTotal.Swap(0)/n))
			}
		}
	}
}

// StartReplPublisher は asia1 で呼ぶ。mux は :8081 の内部向け
func StartReplPublisher(mux *http.ServeMux) {
	mux.HandleFunc("/_internal/repl", repl.serveHTTP)
	go repl.heartbeatLoop()
}

// resetReplicas は initialize の最後に呼ぶ（s.mu の書きロックを持たずに）
func resetReplicas() {
	repl.publish(&opReset{}, replResetTimeout).wait()
}

// ---- europe1・america1: 受けて当てる側 ----

type replica struct {
	s        *Store
	db       *sqlx.DB
	upstream string
	live     atomic.Bool
	lastRecv atomic.Int64
	appliedT atomic.Int64
	// asia1 で流してから当て終えるまでの時間の合計と件数（印ごと）
	lagTotal atomic.Int64
	lagCount atomic.Int64
	advanced atomic.Pointer[chan struct{}]
}

// waitCaughtUp は、時刻 t 以降に流れた印まで当て終えるのを待つ。間に合わなければ false
func (r *replica) waitCaughtUp(t int64) bool {
	var timer *time.Timer
	for {
		if !r.fresh() {
			return false
		}
		ch := r.advanced.Load()
		if r.appliedT.Load() >= t {
			return true
		}
		if timer == nil {
			timer = time.NewTimer(replReadWait)
			defer timer.Stop()
		}
		select {
		case <-*ch:
		case <-timer.C:
			return false
		}
	}
}

// fresh は手元の読み取りを返してよいとき true
func (r *replica) fresh() bool {
	return r.live.Load() && time.Since(time.Unix(0, r.lastRecv.Load())) < replLease
}

func (r *replica) run() {
	loaded := false
	for {
		err := r.safeSession(&loaded)
		r.live.Store(false)
		loaded = false
		log.Printf("repl: stream from %s lost: %v", r.upstream, err)
		time.Sleep(time.Second)
	}
}

type replIn struct {
	m   replMsg
	err error
}

func (r *replica) session(loaded *bool) error {
	conn, err := net.DialTimeout("tcp", r.upstream, 3*time.Second)
	if err != nil {
		return err
	}
	defer conn.Close()
	if _, err := conn.Write([]byte("GET /_internal/repl HTTP/1.1\r\nHost: repl\r\n\r\n")); err != nil {
		return err
	}
	br := bufio.NewReaderSize(conn, 1<<20)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	status, err := br.ReadString('\n')
	if err != nil {
		return err
	}
	if !strings.Contains(status, " 200 ") {
		return fmt.Errorf("status %q", strings.TrimSpace(status))
	}
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			return err
		}
		if line == "\r\n" {
			break
		}
	}
	in := make(chan replIn, 1<<20)
	go func() {
		dec := gob.NewDecoder(br)
		for {
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			var m replMsg
			if err := dec.Decode(&m); err != nil {
				in <- replIn{err: err}
				close(in)
				return
			}
			r.lastRecv.Store(time.Now().UnixNano())
			in <- replIn{m: m}
		}
	}()
	defer func() {
		conn.Close()
		for range in {
		}
	}()
	if !*loaded {
		start := time.Now()
		if err := r.s.LoadAll(context.Background(), r.db); err != nil {
			return err
		}
		*loaded = true
		log.Printf("repl: loaded store in %s (backlog %d)", time.Since(start), len(in))
	}
	readySent := false
	var lastSeq uint64
	for {
		x, ok := <-in
		if !ok {
			return errors.New("closed")
		}
		batch := []replIn{x}
		for len(batch) < 256 && len(in) > 0 {
			batch = append(batch, <-in)
		}
		if err := r.applyBatch(batch, &lastSeq); err != nil {
			return err
		}
		conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
		msg := "a " + strconv.FormatUint(lastSeq, 10) + "\n"
		// 貯めた分をほぼ当て終えたら ready を送る（手元で返し始めるのは、この後に届く live の印を当ててから）
		if !readySent && len(in) < 256 {
			msg += "ready\n"
			readySent = true
		}
		if _, err := conn.Write([]byte(msg)); err != nil {
			return err
		}
	}
}

// safeSession は当てる途中で落ちても台ごと止めず、読み直してつなぎ直す（転送中の書き込みを 502 にしないため）
func (r *replica) safeSession(loaded *bool) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("panic: %v", p)
		}
	}()
	return r.session(loaded)
}

func (r *replica) applyBatch(batch []replIn, lastSeq *uint64) error {
	r.s.mu.Lock()
	defer r.s.mu.Unlock()
	var lastT int64
	defer func() {
		if lastT > 0 {
			r.lagTotal.Add(time.Now().UnixNano() - lastT)
			r.lagCount.Add(1)
			r.appliedT.Store(lastT)
			next := make(chan struct{})
			if old := r.advanced.Swap(&next); old != nil {
				close(*old)
			}
		}
	}()
	for _, b := range batch {
		if b.err != nil {
			return b.err
		}
		switch op := b.m.Op.(type) {
		case *opReset:
			start := time.Now()
			if err := r.resetLocked(); err != nil {
				return err
			}
			log.Printf("repl: reset in %s", time.Since(start))
		case *opLive:
			r.live.Store(true)
			log.Printf("repl: live at seq %d", b.m.Seq)
		case *opHeartbeat:
		case *opDigest:
			r.checkDigestLocked(b.m.Seq, op.Sums)
		default:
			op.apply(r.s)
		}
		*lastSeq = b.m.Seq
		lastT = b.m.T
	}
	return nil
}

func (r *replica) resetLocked() error {
	if err := loadSmall(context.Background(), r.db, r.s); err != nil {
		return err
	}
	r.s.pruneAfterBoundary()
	r.s.rebuildIndexes()
	return nil
}
