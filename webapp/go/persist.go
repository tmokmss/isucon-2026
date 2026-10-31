package isureal

// asia1 のストアは PostgreSQL へ書かず、ファイル（snapshotPath）に残す。
// - 初期データは PostgreSQL の表（init.sh で初期状態に戻す。アプリはもう表へ書かない）。起動時はそこから読み、スナップショットを重ねる
// - スナップショット: 小さい表（利用者・関係・セッション・投稿期間と通知）は全部、大きい表（投稿・RealMoji・コメント）は
//   初期データから変わった行だけ。変更があった秒だけ、読みロックの中で値を写し、ロックの外で gob に書く
//   （一時ファイル → fsync → rename → ディレクトリの fsync）。書く前に sync して、new/ の画像をスナップショットより先にディスクへ出す
// - SIGTERM・SIGINT・SIGHUP・SIGQUIT では書きロックを取って書き込みを止め、スナップショットを書いてから終わる
// - initialize は PostgreSQL と同じ初期化をメモリに行った後、その場でスナップショットを書き直す
// 失いうる幅: 正常な停止・再起動は 0。プロセスだけが落ちたときとハードリブートは最大 1 秒

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

const snapshotPath = "/home/isucon/isureal-snapshot.gob"

// gob は 0 や空文字列を指すポインタを nil に戻すので、位置の省略できる値は有無を別に持つ
type snapPost struct {
	Row                    PostRow
	HasLat, HasLng, HasLoc bool
}

type snapNotified struct {
	WindowID  string
	UserIDs   []string
	IDs       []string
	CreatedAt []time.Time
}

type snapshotData struct {
	Users          []UserRow
	Admins         []AdminRow
	AdminSessions  []AdminSessionRow
	Sessions       []LoginSessionRow
	Windows        []WindowRow
	Notified       []snapNotified
	FriendRequests []FriendRequest
	Friendships    []FriendshipRow
	Follows        []FollowRow
	Blocks         []BlockRow
	Posts          []snapPost
	RealMojis      []RealMojiRow
	Comments       []CommentRow
}

var snapshotFileMu sync.Mutex

func (s *Store) markPostLocked(p *PostRow) {
	s.changedPosts[p.ID] = p
	s.dirty.Store(true)
}

func (s *Store) markRealMojiLocked(r *RealMojiRow) {
	s.changedRealMojis[r.ID] = r
	s.dirty.Store(true)
}

func (s *Store) markCommentLocked(c *CommentRow) {
	s.changedComments[c.ID] = c
	s.dirty.Store(true)
}

// snapshotLocked は読みロック以上を持って呼ぶ。値を写すだけで、エンコードはしない
func (s *Store) snapshotLocked() *snapshotData {
	d := &snapshotData{
		Users:          make([]UserRow, 0, len(s.users)),
		FriendRequests: make([]FriendRequest, 0, len(s.friendRequests)),
		Friendships:    make([]FriendshipRow, 0, len(s.friendships)),
		Follows:        make([]FollowRow, 0, len(s.follows)),
		Blocks:         make([]BlockRow, 0, len(s.blocks)),
		Posts:          make([]snapPost, 0, len(s.changedPosts)),
		RealMojis:      make([]RealMojiRow, 0, len(s.changedRealMojis)),
		Comments:       make([]CommentRow, 0, len(s.changedComments)),
	}
	for _, u := range s.users {
		d.Users = append(d.Users, UserRow{User: u.User, CreatedAt: u.CreatedAt, ProfileImageCreatedAt: u.ProfileImageCreatedAt})
	}
	for _, a := range s.admins {
		d.Admins = append(d.Admins, *a)
	}
	for _, a := range s.adminSessions {
		d.AdminSessions = append(d.AdminSessions, *a)
	}
	for _, x := range s.sessions {
		d.Sessions = append(d.Sessions, *x)
	}
	for _, w := range s.windows {
		d.Windows = append(d.Windows, WindowRow{PostingWindow: w.PostingWindow, CreatedAt: w.CreatedAt})
		if len(w.notified) == 0 {
			continue
		}
		n := snapNotified{WindowID: w.ID, UserIDs: make([]string, 0, len(w.notified)), IDs: make([]string, 0, len(w.notified)), CreatedAt: make([]time.Time, 0, len(w.notified))}
		for userID, row := range w.notified {
			n.UserIDs = append(n.UserIDs, userID)
			n.IDs = append(n.IDs, row.ID)
			n.CreatedAt = append(n.CreatedAt, row.CreatedAt)
		}
		d.Notified = append(d.Notified, n)
	}
	for _, r := range s.friendRequests {
		d.FriendRequests = append(d.FriendRequests, *r)
	}
	for _, f := range s.friendships {
		d.Friendships = append(d.Friendships, *f)
	}
	for _, f := range s.follows {
		d.Follows = append(d.Follows, *f)
	}
	for _, b := range s.blocks {
		d.Blocks = append(d.Blocks, *b)
	}
	for _, p := range s.changedPosts {
		row := *p
		row.window, row.realmojis, row.comments = nil, nil, nil
		d.Posts = append(d.Posts, snapPost{Row: row, HasLat: p.Location.Latitude != nil, HasLng: p.Location.Longitude != nil, HasLoc: p.Location.Locality != nil})
	}
	for _, r := range s.changedRealMojis {
		d.RealMojis = append(d.RealMojis, RealMojiRow{RealMoji: r.RealMoji})
	}
	for _, c := range s.changedComments {
		d.Comments = append(d.Comments, CommentRow{Comment: c.Comment})
	}
	return d
}

func writeSnapshot(d *snapshotData) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(d); err != nil {
		return err
	}
	snapshotFileMu.Lock()
	defer snapshotFileMu.Unlock()
	syscall.Sync()
	dir := filepath.Dir(snapshotPath)
	f, err := os.CreateTemp(dir, ".isureal-snapshot.tmp*")
	if err != nil {
		return err
	}
	_, err = f.Write(buf.Bytes())
	if err == nil {
		err = f.Sync()
	}
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	if err == nil {
		err = os.Rename(f.Name(), snapshotPath)
	}
	if err != nil {
		_ = os.Remove(f.Name())
		return err
	}
	df, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer df.Close()
	return df.Sync()
}

func readSnapshot() (*snapshotData, error) {
	f, err := os.Open(snapshotPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()
	d := &snapshotData{}
	if err := gob.NewDecoder(f).Decode(d); err != nil {
		return nil, err
	}
	return d, nil
}

// applySnapshotLocked は PostgreSQL から読んだ初期データにスナップショットを重ねる。書きロックを持って呼び、後で rebuildIndexes を呼ぶ
func (s *Store) applySnapshotLocked(d *snapshotData) {
	s.users = make(map[string]*UserRow, len(d.Users))
	for i := range d.Users {
		u := &d.Users[i]
		u.setNormalized()
		s.users[u.ID] = u
	}
	s.admins = map[string]*AdminRow{}
	for i := range d.Admins {
		s.admins[d.Admins[i].Username] = &d.Admins[i]
	}
	s.adminSessions = map[string]*AdminSessionRow{}
	for i := range d.AdminSessions {
		s.adminSessions[d.AdminSessions[i].TokenHash] = &d.AdminSessions[i]
	}
	s.sessions = map[string]*LoginSessionRow{}
	for i := range d.Sessions {
		s.sessions[d.Sessions[i].TokenHash] = &d.Sessions[i]
	}
	s.windows = make(map[string]*WindowRow, len(d.Windows))
	for i := range d.Windows {
		w := &d.Windows[i]
		w.notified = map[string]notificationRow{}
		s.windows[w.ID] = w
	}
	for _, n := range d.Notified {
		if w := s.windows[n.WindowID]; w != nil {
			for i, userID := range n.UserIDs {
				w.notified[userID] = notificationRow{ID: n.IDs[i], CreatedAt: n.CreatedAt[i]}
			}
		}
	}
	s.friendRequests = make(map[string]*FriendRequest, len(d.FriendRequests))
	for i := range d.FriendRequests {
		s.friendRequests[d.FriendRequests[i].ID] = &d.FriendRequests[i]
	}
	s.friendships = make([]*FriendshipRow, len(d.Friendships))
	for i := range d.Friendships {
		s.friendships[i] = &d.Friendships[i]
	}
	s.follows = make([]*FollowRow, len(d.Follows))
	for i := range d.Follows {
		s.follows[i] = &d.Follows[i]
	}
	s.blocks = make([]*BlockRow, len(d.Blocks))
	for i := range d.Blocks {
		s.blocks[i] = &d.Blocks[i]
	}
	var zeroF float64
	var zeroS string
	for i := range d.Posts {
		sp := &d.Posts[i]
		loc := &sp.Row.Location
		if sp.HasLat && loc.Latitude == nil {
			v := zeroF
			loc.Latitude = &v
		}
		if sp.HasLng && loc.Longitude == nil {
			v := zeroF
			loc.Longitude = &v
		}
		if sp.HasLoc && loc.Locality == nil {
			v := zeroS
			loc.Locality = &v
		}
		p := s.posts[sp.Row.ID]
		if p == nil {
			p = &sp.Row
			s.posts[p.ID] = p
		} else {
			*p = sp.Row
		}
		s.changedPosts[p.ID] = p
	}
	for i := range d.RealMojis {
		r := s.realmojis[d.RealMojis[i].ID]
		if r == nil {
			r = &d.RealMojis[i]
			s.realmojis[r.ID] = r
		} else {
			*r = d.RealMojis[i]
		}
		s.changedRealMojis[r.ID] = r
	}
	for i := range d.Comments {
		c := &d.Comments[i]
		if c.DeletedAt != nil {
			delete(s.comments, c.ID)
		} else if old := s.comments[c.ID]; old != nil {
			*old = *c
			c = old
		} else {
			s.comments[c.ID] = c
		}
		s.changedComments[c.ID] = c
	}
}

// startPersist は読み込みが終わった後に呼ぶ
func (s *Store) startPersist() {
	go func() {
		for range time.Tick(time.Second) {
			if !s.dirty.Swap(false) {
				continue
			}
			s.mu.RLock()
			d := s.snapshotLocked()
			s.mu.RUnlock()
			if err := writeSnapshot(d); err != nil {
				s.dirty.Store(true)
				log.Printf("snapshot: %v", err)
			}
		}
	}()
}

// handleStopSignals は起動直後に呼ぶ。読み込みが終わる前のシグナルでは、手元のファイルを変えずに終わる
func (s *Store) handleStopSignals(loaded <-chan struct{}) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		sig := <-ch
		select {
		case <-loaded:
		default:
			log.Printf("got %s before the store was loaded; exiting without a snapshot", sig)
			os.Exit(0)
		}
		start := time.Now()
		s.mu.Lock()
		d := s.snapshotLocked()
		if err := writeSnapshot(d); err != nil {
			log.Printf("snapshot on %s: %v", sig, err)
			os.Exit(1)
		}
		log.Printf("snapshot on %s: users=%d posts=%d realmojis=%d comments=%d in %s", sig, len(d.Users), len(d.Posts), len(d.RealMojis), len(d.Comments), time.Since(start))
		os.Exit(0)
	}()
}

// resetChangesLocked は initialize で境界より後の行を捨てた後に呼ぶ。メモリに残した初期データの行の変更は持ち続ける
func (s *Store) resetChangesLocked() {
	for id, p := range s.changedPosts {
		if s.posts[id] != p {
			delete(s.changedPosts, id)
		}
	}
	for id, r := range s.changedRealMojis {
		if s.realmojis[id] != r {
			delete(s.changedRealMojis, id)
		}
	}
	for id, c := range s.changedComments {
		if c.CreatedAt.After(s.boundary) {
			delete(s.changedComments, id)
		}
	}
}
