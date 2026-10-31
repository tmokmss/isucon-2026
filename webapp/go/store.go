package isureal

// 状態はこのプロセスのメモリに持ち、書き込みは先に PostgreSQL へ同期で書いてからメモリに反映する。
// 起動時と /api/initialize の後に PostgreSQL から読み直すので、再起動しても状態は失われない。
// 状態を持つのは 1 台（asia1）だけで、他の台の nginx は /api/ をこの台へ転送する。

import (
	"context"
	"hash/fnv"
	"log"
	"sort"
	"sync"
	"time"
	"unsafe"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"
)

type UserRow struct {
	User
	CreatedAt             time.Time
	ProfileImageCreatedAt *time.Time
	normUsername          string
	normDisplayName       string
	posts                 []*PostRow
	friends               map[string]*FriendshipRow
	following             map[string]*FollowRow
	followers             map[string]*FollowRow
	blockOut              map[string]*BlockRow
	blockIn               map[string]*BlockRow
	pendingOut            map[string]*FriendRequest
	pendingIn             map[string]*FriendRequest
	sessions              map[string]*LoginSessionRow
}

func (u *UserRow) active() bool { return u.Status == "active" }

func (u *UserRow) setNormalized() {
	u.normUsername = normalizeSearchText(u.Username)
	u.normDisplayName = normalizeSearchText(u.DisplayName)
}

type WindowRow struct {
	PostingWindow
	CreatedAt time.Time
	posts     []*PostRow
	notified  map[string]notificationRow
}

type notificationRow struct {
	ID        string
	CreatedAt time.Time
}

type PostRow struct {
	ID         string
	AuthorID   string
	WindowID   string
	ImageID    string
	Caption    string
	Visibility string
	Location   PostLocation
	PostedAt   time.Time
	CreatedAt  time.Time
	DeletedAt  *time.Time
	window     *WindowRow
	realmojis  []*RealMojiRow
	comments   []*CommentRow
}

func (p *PostRow) post() Post {
	return Post{
		ID: p.ID, AuthorID: p.AuthorID, PostingWindowID: p.WindowID, ImageID: p.ImageID,
		Caption: p.Caption, Visibility: p.Visibility, PostLocation: p.Location, PostedAt: p.PostedAt,
		OnTimeDeadline: p.window.OnTimeDeadline, WindowEndedAt: p.window.EndedAt,
	}
}

type RealMojiRow struct {
	RealMoji
	post *PostRow
}

func (r *RealMojiRow) current() bool { return r.ReplacedAt == nil && r.DeletedAt == nil }

type CommentRow struct {
	Comment
	post *PostRow
}

type FriendshipRow struct {
	ID              string
	LowID           string
	HighID          string
	SourceRequestID *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type FollowRow struct {
	ID         string
	FollowerID string
	FolloweeID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type BlockRow struct {
	ID        string
	BlockerID string
	BlockedID string
	CreatedAt time.Time
}

type LoginSessionRow struct {
	ID        string
	UserID    string
	TokenHash string
	Active    bool
	CreatedAt time.Time
}

type AdminRow struct {
	ID           string
	Username     string
	PasswordHash string
}

type AdminSessionRow struct {
	ID              string
	AdministratorID string
	TokenHash       string
	Active          bool
	CreatedAt       time.Time
}

type Store struct {
	mu sync.RWMutex

	users          map[string]*UserRow
	usersByName    map[string]*UserRow
	usersSorted    []*UserRow
	profileByImage map[string]*UserRow
	sessions       map[string]*LoginSessionRow
	admins         map[string]*AdminRow
	adminSessions  map[string]*AdminSessionRow
	windows        map[string]*WindowRow
	activeWindows  map[string]*WindowRow
	maxLocalDate   map[string]time.Time
	posts          map[string]*PostRow
	postByImage    map[string]*PostRow
	realmojis      map[string]*RealMojiRow
	rmByImage      map[string]*RealMojiRow
	comments       map[string]*CommentRow
	friendRequests map[string]*FriendRequest
	friendships    []*FriendshipRow
	follows        []*FollowRow
	blocks         []*BlockRow
	boundary       time.Time

	notifyCh chan struct{}
}

func now() time.Time { return time.Now().UTC().Truncate(time.Microsecond) }

// 書き込みの流れを利用者ごとに直列にする（元の実装の users 行ロックに相当する）
var userLocks [4096]sync.Mutex

func lockUsers(ids ...string) func() {
	idx := make([]int, 0, len(ids))
	for _, id := range ids {
		h := fnv.New32a()
		h.Write([]byte(id))
		idx = append(idx, int(h.Sum32()%uint32(len(userLocks))))
	}
	sort.Ints(idx)
	uniq := idx[:0]
	for i, v := range idx {
		if i == 0 || v != idx[i-1] {
			uniq = append(uniq, v)
		}
	}
	for _, i := range uniq {
		userLocks[i].Lock()
	}
	return func() {
		for i := len(uniq) - 1; i >= 0; i-- {
			userLocks[uniq[i]].Unlock()
		}
	}
}

func (s *Store) broadcastLocked() {
	close(s.notifyCh)
	s.notifyCh = make(chan struct{})
}

func (s *Store) user(id string) *UserRow {
	return s.users[id]
}

func (s *Store) activeUser(id string) *UserRow {
	u := s.users[id]
	if u == nil || !u.active() {
		return nil
	}
	return u
}

func (s *Store) isBlocked(a, b string) bool {
	u := s.users[a]
	if u == nil {
		return false
	}
	return u.blockOut[b] != nil || u.blockIn[b] != nil
}

func pairKey(a, b string) (string, string) { return min(a, b), max(a, b) }

// 大きい表の行はまとめて確保し、繰り返し現れる文字列（利用者 ID・本文など）は 1 つを共有する。GC が辿るオブジェクトを減らすため
func nextSlot[T any](slab *[]T) *T {
	if len(*slab) == 0 {
		*slab = make([]T, 4096)
	}
	v := &(*slab)[0]
	*slab = (*slab)[1:]
	return v
}

type interner struct {
	strs map[string]string
	ptrs map[string]*string
	buf  []byte
}

// id は一意な文字列を大きな領域へ詰めて写す。書き込んだ後の領域は変更しない
func (in *interner) id(v string) string {
	if len(v) == 0 {
		return ""
	}
	if cap(in.buf)-len(in.buf) < len(v) {
		in.buf = make([]byte, 0, 1<<20)
	}
	start := len(in.buf)
	in.buf = append(in.buf, v...)
	return unsafe.String(&in.buf[start], len(v))
}

func (in *interner) str(v string) string {
	if u, ok := in.strs[v]; ok {
		return u
	}
	in.strs[v] = v
	return v
}

func (in *interner) ptr(v *string) *string {
	if v == nil {
		return nil
	}
	if u, ok := in.ptrs[*v]; ok {
		return u
	}
	in.ptrs[*v] = v
	return v
}

func loadBig(ctx context.Context, db *sqlx.DB, s *Store) error {
	start := time.Now()
	in := &interner{strs: map[string]string{}, ptrs: map[string]*string{}}
	s.posts = make(map[string]*PostRow, 700000)
	rows, err := db.QueryContext(ctx, `SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM posts`)
	if err != nil {
		return err
	}
	var postSlab []PostRow
	for rows.Next() {
		p := nextSlot(&postSlab)
		if err := rows.Scan(&p.ID, &p.AuthorID, &p.WindowID, &p.ImageID, &p.Caption, &p.Visibility, &p.Location.Kind, &p.Location.Latitude, &p.Location.Longitude, &p.Location.Locality, &p.PostedAt, &p.CreatedAt, &p.DeletedAt); err != nil {
			rows.Close()
			return err
		}
		p.ID = in.id(p.ID)
		p.ImageID = in.id(p.ImageID)
		p.AuthorID = in.str(p.AuthorID)
		p.WindowID = in.str(p.WindowID)
		p.Caption = in.str(p.Caption)
		p.Visibility = in.str(p.Visibility)
		p.Location.Kind = in.str(p.Location.Kind)
		p.Location.Locality = in.ptr(p.Location.Locality)
		p.PostedAt = p.PostedAt.UTC()
		p.CreatedAt = p.CreatedAt.UTC()
		s.posts[p.ID] = p
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}
	log.Printf("loaded posts=%d in %s", len(s.posts), time.Since(start))

	s.realmojis = make(map[string]*RealMojiRow, 1300000)
	rows, err = db.QueryContext(ctx, `SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis`)
	if err != nil {
		return err
	}
	var rmSlab []RealMojiRow
	for rows.Next() {
		r := nextSlot(&rmSlab)
		if err := rows.Scan(&r.ID, &r.PostID, &r.ReactorID, &r.ImageID, &r.ReactionType, &r.CreatedAt, &r.ReplacedAt, &r.DeletedAt); err != nil {
			rows.Close()
			return err
		}
		r.ID = in.id(r.ID)
		r.ImageID = in.id(r.ImageID)
		r.PostID = s.posts[r.PostID].ID
		r.ReactorID = in.str(r.ReactorID)
		r.ReactionType = in.str(r.ReactionType)
		r.CreatedAt = r.CreatedAt.UTC()
		s.realmojis[r.ID] = r
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}
	log.Printf("loaded realmojis=%d in %s", len(s.realmojis), time.Since(start))

	s.comments = make(map[string]*CommentRow, 600000)
	rows, err = db.QueryContext(ctx, `SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL`)
	if err != nil {
		return err
	}
	var commentSlab []CommentRow
	for rows.Next() {
		c := nextSlot(&commentSlab)
		if err := rows.Scan(&c.ID, &c.PostID, &c.AuthorID, &c.Body, &c.CreatedAt); err != nil {
			rows.Close()
			return err
		}
		c.ID = in.id(c.ID)
		c.PostID = s.posts[c.PostID].ID
		c.AuthorID = in.str(c.AuthorID)
		c.Body = in.str(c.Body)
		c.CreatedAt = c.CreatedAt.UTC()
		s.comments[c.ID] = c
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}
	log.Printf("loaded comments=%d in %s", len(s.comments), time.Since(start))
	return nil
}

func loadSmall(ctx context.Context, db *sqlx.DB, s *Store) error {
	if err := db.GetContext(ctx, &s.boundary, `SELECT created_at FROM initial_data_boundary`); err != nil {
		return err
	}
	s.boundary = s.boundary.UTC()

	s.users = map[string]*UserRow{}
	rows, err := db.QueryContext(ctx, `SELECT id, username, secret_key_hash, display_name, bio, profile_image_id, region, account_type, status, created_at FROM users`)
	if err != nil {
		return err
	}
	for rows.Next() {
		u := &UserRow{}
		if err := rows.Scan(&u.ID, &u.Username, &u.SecretKeyHash, &u.DisplayName, &u.Bio, &u.ProfileImageID, &u.Region, &u.AccountType, &u.Status, &u.CreatedAt); err != nil {
			rows.Close()
			return err
		}
		if err := u.setImageKey(); err != nil {
			rows.Close()
			return err
		}
		u.setNormalized()
		s.users[u.ID] = u
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}

	var profileImages []struct {
		ID        string    `db:"id"`
		CreatedAt time.Time `db:"created_at"`
	}
	if err := db.SelectContext(ctx, &profileImages, `SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id`); err != nil {
		return err
	}
	for _, pi := range profileImages {
		if u := s.users[pi.ID]; u != nil {
			t := pi.CreatedAt
			u.ProfileImageCreatedAt = &t
		}
	}

	s.friendships = nil
	if err := selectRows(ctx, db, `SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships`, func(sc scanner) error {
		f := &FriendshipRow{}
		if err := sc.Scan(&f.ID, &f.LowID, &f.HighID, &f.SourceRequestID, &f.CreatedAt, &f.UpdatedAt); err != nil {
			return err
		}
		s.friendships = append(s.friendships, f)
		return nil
	}); err != nil {
		return err
	}
	s.follows = nil
	if err := selectRows(ctx, db, `SELECT id, follower_id, followee_id, created_at, updated_at FROM follows`, func(sc scanner) error {
		f := &FollowRow{}
		if err := sc.Scan(&f.ID, &f.FollowerID, &f.FolloweeID, &f.CreatedAt, &f.UpdatedAt); err != nil {
			return err
		}
		s.follows = append(s.follows, f)
		return nil
	}); err != nil {
		return err
	}
	s.blocks = nil
	if err := selectRows(ctx, db, `SELECT id, blocker_id, blocked_id, created_at FROM block_relations`, func(sc scanner) error {
		b := &BlockRow{}
		if err := sc.Scan(&b.ID, &b.BlockerID, &b.BlockedID, &b.CreatedAt); err != nil {
			return err
		}
		s.blocks = append(s.blocks, b)
		return nil
	}); err != nil {
		return err
	}
	s.friendRequests = map[string]*FriendRequest{}
	if err := selectRows(ctx, db, `SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests`, func(sc scanner) error {
		r := &FriendRequest{}
		if err := sc.Scan(&r.ID, &r.RequesterID, &r.RecipientID, &r.Status, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return err
		}
		s.friendRequests[r.ID] = r
		return nil
	}); err != nil {
		return err
	}
	s.sessions = map[string]*LoginSessionRow{}
	if err := selectRows(ctx, db, `SELECT id, user_id, token_hash, created_at FROM login_sessions WHERE active = true`, func(sc scanner) error {
		r := &LoginSessionRow{Active: true}
		if err := sc.Scan(&r.ID, &r.UserID, &r.TokenHash, &r.CreatedAt); err != nil {
			return err
		}
		s.sessions[r.TokenHash] = r
		return nil
	}); err != nil {
		return err
	}
	s.admins = map[string]*AdminRow{}
	if err := selectRows(ctx, db, `SELECT id, username, password_hash FROM administrators`, func(sc scanner) error {
		r := &AdminRow{}
		if err := sc.Scan(&r.ID, &r.Username, &r.PasswordHash); err != nil {
			return err
		}
		s.admins[r.Username] = r
		return nil
	}); err != nil {
		return err
	}
	s.adminSessions = map[string]*AdminSessionRow{}
	if err := selectRows(ctx, db, `SELECT id, administrator_id, token_hash, created_at FROM administrator_sessions WHERE active = true`, func(sc scanner) error {
		r := &AdminSessionRow{Active: true}
		if err := sc.Scan(&r.ID, &r.AdministratorID, &r.TokenHash, &r.CreatedAt); err != nil {
			return err
		}
		s.adminSessions[r.TokenHash] = r
		return nil
	}); err != nil {
		return err
	}
	s.windows = map[string]*WindowRow{}
	if err := selectRows(ctx, db, `SELECT id, region, local_date, starts_at, on_time_deadline, ended_at, created_at FROM posting_windows`, func(sc scanner) error {
		w := &WindowRow{notified: map[string]notificationRow{}}
		if err := sc.Scan(&w.ID, &w.Region, &w.LocalDate, &w.StartsAt, &w.OnTimeDeadline, &w.EndedAt, &w.CreatedAt); err != nil {
			return err
		}
		s.windows[w.ID] = w
		return nil
	}); err != nil {
		return err
	}
	// 投稿期間を開いたときの通知は posting_windows の配列に持つ（初期データの通知は notifications にある）
	if _, err := db.ExecContext(ctx, `ALTER TABLE posting_windows ADD COLUMN IF NOT EXISTS notification_ids text[], ADD COLUMN IF NOT EXISTS notified_user_ids text[]`); err != nil {
		return err
	}
	typeMap := pgtype.NewMap()
	if err := selectRows(ctx, db, `SELECT id, notification_ids, notified_user_ids FROM posting_windows WHERE ended_at IS NULL AND notification_ids IS NOT NULL`, func(sc scanner) error {
		var windowID string
		var ids, userIDs []string
		if err := sc.Scan(&windowID, typeMap.SQLScanner(&ids), typeMap.SQLScanner(&userIDs)); err != nil {
			return err
		}
		if w := s.windows[windowID]; w != nil {
			for i, userID := range userIDs {
				w.notified[userID] = notificationRow{ID: ids[i], CreatedAt: w.StartsAt}
			}
		}
		return nil
	}); err != nil {
		return err
	}
	if err := selectRows(ctx, db, `
  SELECT n.id, n.user_id, n.posting_window_id, n.created_at
  FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id
  WHERE w.ended_at IS NULL`, func(sc scanner) error {
		var id, userID, windowID string
		var createdAt time.Time
		if err := sc.Scan(&id, &userID, &windowID, &createdAt); err != nil {
			return err
		}
		if w := s.windows[windowID]; w != nil {
			w.notified[userID] = notificationRow{ID: id, CreatedAt: createdAt}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

type scanner interface {
	Scan(dest ...any) error
}

func selectRows(ctx context.Context, db *sqlx.DB, query string, fn func(scanner) error) error {
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := fn(rows); err != nil {
			return err
		}
	}
	return rows.Err()
}

// rebuildIndexes は読み込んだ行から派生のインデックスを作り直す。書きロックを持って呼ぶ
func (s *Store) rebuildIndexes() {
	s.usersByName = make(map[string]*UserRow, len(s.users))
	s.usersSorted = make([]*UserRow, 0, len(s.users))
	s.profileByImage = make(map[string]*UserRow, len(s.users))
	for _, u := range s.users {
		u.posts = nil
		u.friends = map[string]*FriendshipRow{}
		u.following = map[string]*FollowRow{}
		u.followers = map[string]*FollowRow{}
		u.blockOut = map[string]*BlockRow{}
		u.blockIn = map[string]*BlockRow{}
		u.pendingOut = map[string]*FriendRequest{}
		u.pendingIn = map[string]*FriendRequest{}
		u.sessions = map[string]*LoginSessionRow{}
		s.usersByName[u.Username] = u
		s.usersSorted = append(s.usersSorted, u)
		if u.ProfileImageID != nil {
			s.profileByImage[*u.ProfileImageID] = u
		}
	}
	sort.Slice(s.usersSorted, func(i, j int) bool { return userLess(s.usersSorted[i], s.usersSorted[j]) })
	for _, f := range s.friendships {
		if lo, hi := s.users[f.LowID], s.users[f.HighID]; lo != nil && hi != nil {
			lo.friends[f.HighID] = f
			hi.friends[f.LowID] = f
		}
	}
	for _, f := range s.follows {
		if a, b := s.users[f.FollowerID], s.users[f.FolloweeID]; a != nil && b != nil {
			a.following[f.FolloweeID] = f
			b.followers[f.FollowerID] = f
		}
	}
	for _, bl := range s.blocks {
		if a, b := s.users[bl.BlockerID], s.users[bl.BlockedID]; a != nil && b != nil {
			a.blockOut[bl.BlockedID] = bl
			b.blockIn[bl.BlockerID] = bl
		}
	}
	for _, r := range s.friendRequests {
		if r.Status != "pending" {
			continue
		}
		if a, b := s.users[r.RequesterID], s.users[r.RecipientID]; a != nil && b != nil {
			a.pendingOut[r.RecipientID] = r
			b.pendingIn[r.RequesterID] = r
		}
	}
	for _, sess := range s.sessions {
		if u := s.users[sess.UserID]; u != nil {
			u.sessions[sess.TokenHash] = sess
		}
	}
	s.activeWindows = map[string]*WindowRow{}
	s.maxLocalDate = map[string]time.Time{}
	for _, w := range s.windows {
		w.posts = nil
		if w.EndedAt == nil {
			s.activeWindows[w.Region] = w
		}
		if d, ok := s.maxLocalDate[w.Region]; !ok || w.LocalDate.After(d) {
			s.maxLocalDate[w.Region] = w.LocalDate
		}
	}
	s.postByImage = make(map[string]*PostRow, len(s.posts))
	for _, p := range s.posts {
		p.realmojis = nil
		p.comments = nil
		p.window = s.windows[p.WindowID]
		if u := s.users[p.AuthorID]; u != nil {
			u.posts = append(u.posts, p)
		}
		p.window.posts = append(p.window.posts, p)
		s.postByImage[p.ImageID] = p
	}
	for _, u := range s.users {
		sort.Slice(u.posts, func(i, j int) bool { return postLess(u.posts[i], u.posts[j]) })
	}
	for _, w := range s.windows {
		sort.Slice(w.posts, func(i, j int) bool { return postLess(w.posts[i], w.posts[j]) })
	}
	s.rmByImage = make(map[string]*RealMojiRow, len(s.realmojis))
	for _, r := range s.realmojis {
		r.post = s.posts[r.PostID]
		r.post.realmojis = append(r.post.realmojis, r)
		s.rmByImage[r.ImageID] = r
	}
	for _, c := range s.comments {
		c.post = s.posts[c.PostID]
		c.post.comments = append(c.post.comments, c)
	}
	for _, p := range s.posts {
		if len(p.realmojis) > 1 {
			sort.Slice(p.realmojis, func(i, j int) bool { return rmLess(p.realmojis[i], p.realmojis[j]) })
		}
		if len(p.comments) > 1 {
			sort.Slice(p.comments, func(i, j int) bool { return commentLess(p.comments[i], p.comments[j]) })
		}
	}
}

func userLess(a, b *UserRow) bool {
	if a.Username != b.Username {
		return a.Username < b.Username
	}
	return a.ID < b.ID
}

func postLess(a, b *PostRow) bool {
	if !a.PostedAt.Equal(b.PostedAt) {
		return a.PostedAt.Before(b.PostedAt)
	}
	return a.ID < b.ID
}

func rmLess(a, b *RealMojiRow) bool {
	if !a.CreatedAt.Equal(b.CreatedAt) {
		return a.CreatedAt.Before(b.CreatedAt)
	}
	return a.ID < b.ID
}

func commentLess(a, b *CommentRow) bool {
	if !a.CreatedAt.Equal(b.CreatedAt) {
		return a.CreatedAt.Before(b.CreatedAt)
	}
	return a.ID < b.ID
}

// insertSorted は順序を保って末尾近くに挿入する（ほぼ時刻順に来るので末尾から探す）
func insertPostSorted(list []*PostRow, p *PostRow) []*PostRow {
	i := len(list)
	for i > 0 && postLess(p, list[i-1]) {
		i--
	}
	list = append(list, nil)
	copy(list[i+1:], list[i:])
	list[i] = p
	return list
}

func (s *Store) addPostLocked(p *PostRow) {
	p.window = s.windows[p.WindowID]
	s.posts[p.ID] = p
	s.postByImage[p.ImageID] = p
	if u := s.users[p.AuthorID]; u != nil {
		u.posts = insertPostSorted(u.posts, p)
	}
	p.window.posts = insertPostSorted(p.window.posts, p)
}

func (s *Store) addRealMojiLocked(r *RealMojiRow) {
	r.post = s.posts[r.PostID]
	s.realmojis[r.ID] = r
	s.rmByImage[r.ImageID] = r
	list := r.post.realmojis
	i := len(list)
	for i > 0 && rmLess(r, list[i-1]) {
		i--
	}
	list = append(list, nil)
	copy(list[i+1:], list[i:])
	list[i] = r
	r.post.realmojis = list
}

func (s *Store) addCommentLocked(c *CommentRow) {
	c.post = s.posts[c.PostID]
	s.comments[c.ID] = c
	list := c.post.comments
	i := len(list)
	for i > 0 && commentLess(c, list[i-1]) {
		i--
	}
	list = append(list, nil)
	copy(list[i+1:], list[i:])
	list[i] = c
	c.post.comments = list
}

func (s *Store) addUserLocked(u *UserRow) {
	u.friends = map[string]*FriendshipRow{}
	u.following = map[string]*FollowRow{}
	u.followers = map[string]*FollowRow{}
	u.blockOut = map[string]*BlockRow{}
	u.blockIn = map[string]*BlockRow{}
	u.pendingOut = map[string]*FriendRequest{}
	u.pendingIn = map[string]*FriendRequest{}
	u.sessions = map[string]*LoginSessionRow{}
	s.users[u.ID] = u
	s.usersByName[u.Username] = u
	i := sort.Search(len(s.usersSorted), func(i int) bool { return userLess(u, s.usersSorted[i]) })
	s.usersSorted = append(s.usersSorted, nil)
	copy(s.usersSorted[i+1:], s.usersSorted[i:])
	s.usersSorted[i] = u
}

// pruneAfterBoundary は reset.sql と同じく、初期データの境界より後に作られた投稿・RealMoji・コメントを捨てる
func (s *Store) pruneAfterBoundary() {
	for id, p := range s.posts {
		if p.CreatedAt.After(s.boundary) {
			delete(s.posts, id)
		}
	}
	for id, r := range s.realmojis {
		if r.CreatedAt.After(s.boundary) {
			delete(s.realmojis, id)
		}
	}
	for id, c := range s.comments {
		if c.CreatedAt.After(s.boundary) {
			delete(s.comments, id)
		}
	}
}

func NewStore() *Store {
	return &Store{notifyCh: make(chan struct{})}
}

func (s *Store) LoadAll(ctx context.Context, db *sqlx.DB) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := loadBig(ctx, db, s); err != nil {
		return err
	}
	if err := loadSmall(ctx, db, s); err != nil {
		return err
	}
	start := time.Now()
	s.rebuildIndexes()
	log.Printf("rebuilt indexes in %s", time.Since(start))
	return nil
}
