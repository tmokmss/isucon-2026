package isureal

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

// 受付中のすべての投稿をたどって並べ替える実装。索引を使う実装と結果を比べる
func (s *Store) feedCandidatesScanLocked(viewer viewerContext, following map[string]bool, feed string, ranked bool, beforePostedAt *time.Time, beforeID string) []feedCandidate {
	var rows []*PostRow
	for _, w := range s.activeWindows {
		for _, p := range w.posts {
			if p.DeletedAt != nil || p.AuthorID == viewer.ID {
				continue
			}
			if beforePostedAt != nil && !(p.PostedAt.Before(*beforePostedAt) || (p.PostedAt.Equal(*beforePostedAt) && p.ID < beforeID)) {
				continue
			}
			rows = append(rows, p)
		}
	}
	sort.Slice(rows, func(i, j int) bool { return postLess(rows[j], rows[i]) })
	candidates := make([]feedCandidate, 0)
	for _, p := range rows {
		author := s.activeUser(p.AuthorID)
		if author == nil {
			continue
		}
		post := p.post()
		if viewer.canView(post) && matchesFeed(viewer, following, post, feed) {
			candidates = append(candidates, feedCandidate{post: post, author: author.User, row: p})
			if !ranked && len(candidates) == 51 {
				break
			}
		}
	}
	return candidates
}

var feedTestRegions = []string{"asia", "europe", "america"}

type feedFixture struct {
	s      *Store
	rng    *rand.Rand
	base   time.Time
	nextID int
	ids    []string
}

func (f *feedFixture) id(prefix string) string {
	f.nextID++
	return fmt.Sprintf("%s-%06d", prefix, f.nextID)
}

// 同じ時刻の投稿が並ぶように、時刻は粗く選ぶ
func (f *feedFixture) postedAt() time.Time {
	return f.base.Add(time.Duration(f.rng.Intn(200)) * time.Second)
}

func (f *feedFixture) visibility() string {
	return []string{"public", "friends", "friends_of_friends"}[f.rng.Intn(3)]
}

func (f *feedFixture) user() string { return f.ids[f.rng.Intn(len(f.ids))] }

func newFeedFixture(seed int64, nUsers int) *feedFixture {
	f := &feedFixture{rng: rand.New(rand.NewSource(seed)), base: time.Date(2026, 10, 31, 0, 0, 0, 0, time.UTC)}
	s := NewStore()
	f.s = s
	s.users = map[string]*UserRow{}
	s.windows = map[string]*WindowRow{}
	s.posts = map[string]*PostRow{}
	s.realmojis = map[string]*RealMojiRow{}
	s.comments = map[string]*CommentRow{}
	s.friendRequests = map[string]*FriendRequest{}
	s.sessions = map[string]*LoginSessionRow{}
	for _, region := range feedTestRegions {
		for day := 0; day < 2; day++ {
			w := &WindowRow{PostingWindow: PostingWindow{ID: f.id("w"), Region: region, LocalDate: f.base.AddDate(0, 0, day-1)}, notified: map[string]notificationRow{}}
			if day == 0 {
				t := f.base
				w.EndedAt = &t
			}
			s.windows[w.ID] = w
		}
	}
	for i := 0; i < nUsers; i++ {
		u := &UserRow{User: User{ID: f.id("u"), Username: fmt.Sprintf("user%d", i), Region: feedTestRegions[f.rng.Intn(3)], Status: "active"}}
		if f.rng.Intn(20) == 0 {
			u.Status = "deleted"
		}
		s.users[u.ID] = u
		f.ids = append(f.ids, u.ID)
	}
	for i := 0; i < nUsers*4; i++ {
		a, b := f.user(), f.user()
		if a == b {
			continue
		}
		lo, hi := pairKey(a, b)
		s.friendships = append(s.friendships, &FriendshipRow{ID: f.id("f"), LowID: lo, HighID: hi})
	}
	for i := 0; i < nUsers*2; i++ {
		a, b := f.user(), f.user()
		if a != b {
			s.follows = append(s.follows, &FollowRow{ID: f.id("fo"), FollowerID: a, FolloweeID: b})
		}
	}
	for i := 0; i < nUsers/5; i++ {
		a, b := f.user(), f.user()
		if a != b {
			s.blocks = append(s.blocks, &BlockRow{ID: f.id("b"), BlockerID: a, BlockedID: b})
		}
	}
	for _, w := range s.windows {
		for _, uid := range f.ids {
			if s.users[uid].Region != w.Region || f.rng.Intn(4) == 0 {
				continue
			}
			for k := 0; k < 1+f.rng.Intn(2); k++ {
				p := &PostRow{ID: f.id("p"), AuthorID: uid, WindowID: w.ID, ImageID: f.id("img"), Visibility: f.visibility(), Location: PostLocation{Kind: "none"}, PostedAt: f.postedAt()}
				if f.rng.Intn(6) == 0 {
					t := p.PostedAt
					p.DeletedAt = &t
				}
				s.posts[p.ID] = p
			}
		}
	}
	s.rebuildIndexes()
	return f
}

// mutate はハンドラがメモリに反映するのと同じ書き換えを 1 つ行う。書きロックを持って呼ぶ
func (f *feedFixture) mutate() {
	s := f.s
	switch f.rng.Intn(9) {
	case 0, 1: // 投稿
		u := s.users[f.user()]
		w := s.activeWindows[u.Region]
		s.addPostLocked(&PostRow{ID: f.id("p"), AuthorID: u.ID, WindowID: w.ID, ImageID: f.id("img"), Visibility: f.visibility(), Location: PostLocation{Kind: "none"}, PostedAt: f.postedAt()})
	case 2: // 削除
		for _, p := range s.posts {
			if p.DeletedAt == nil {
				t := p.PostedAt
				p.DeletedAt = &t
				break
			}
		}
	case 3: // 公開範囲の変更
		for _, p := range s.posts {
			p.Visibility = f.visibility()
			break
		}
	case 4: // 友達の追加
		a, b := f.user(), f.user()
		if a != b && s.users[a].friends[b] == nil {
			lo, hi := pairKey(a, b)
			s.addFriendshipLocked(&FriendshipRow{ID: f.id("f"), LowID: lo, HighID: hi})
		}
	case 5: // 友達の解除
		a := s.users[f.user()]
		for b := range a.friends {
			s.removeFriendshipLocked(a.ID, b)
			break
		}
	case 6: // ブロック
		a, b := f.user(), f.user()
		if a != b && s.users[a].blockOut[b] == nil {
			bl := &BlockRow{ID: f.id("b"), BlockerID: a, BlockedID: b}
			s.blocks = append(s.blocks, bl)
			s.users[a].blockOut[b] = bl
			s.users[b].blockIn[a] = bl
			s.removeFriendshipLocked(a, b)
			s.removeFollowLocked(a, b)
			s.removeFollowLocked(b, a)
		}
	case 7: // 退会
		s.users[f.user()].Status = "deleted"
	case 8: // 投稿期間の切り替え
		region := feedTestRegions[f.rng.Intn(3)]
		old := s.activeWindows[region]
		t := f.base
		old.EndedAt = &t
		w := &WindowRow{PostingWindow: PostingWindow{ID: f.id("w"), Region: region, LocalDate: old.LocalDate.AddDate(0, 0, 1)}, notified: map[string]notificationRow{}}
		s.windows[w.ID] = w
		s.activeWindows[region] = w
	}
}

func candidateIDs(cands []feedCandidate) string {
	out := ""
	for _, c := range cands {
		out += fmt.Sprintf("%s/%s/%s/%s,", c.post.ID, c.author.ID, c.post.Visibility, c.post.PostedAt.Format(time.RFC3339Nano))
	}
	return out
}

// compareAll は全員の 4 種類のフィードを、先頭と途中のページで比べる。読みロックを持って呼ぶ
func (f *feedFixture) compareAll(t *testing.T) int {
	s := f.s
	nonEmpty := 0
	for _, uid := range f.ids {
		viewer, ok := s.loadViewerContextLocked(uid)
		if !ok {
			continue
		}
		for _, feed := range []string{"friends", "friends-of-friends", "public", "following"} {
			following := map[string]bool{}
			if feed == "following" {
				for id := range s.users[uid].following {
					following[id] = true
				}
			}
			ranked := feed == "friends-of-friends"
			cursors := []*PostRow{nil}
			if !ranked {
				for _, p := range s.posts {
					cursors = append(cursors, p)
					if len(cursors) == 4 {
						break
					}
				}
			}
			for _, c := range cursors {
				var before *time.Time
				beforeID := ""
				if c != nil {
					tt := c.PostedAt
					before, beforeID = &tt, c.ID
				}
				got := candidateIDs(s.feedCandidatesLocked(viewer, following, feed, ranked, before, beforeID))
				want := candidateIDs(s.feedCandidatesScanLocked(viewer, following, feed, ranked, before, beforeID))
				if got != want {
					t.Fatalf("viewer=%s feed=%s cursor=%v\ngot  %s\nwant %s", uid, feed, c != nil, got, want)
				}
				if got != "" {
					nonEmpty++
				}
			}
		}
	}
	return nonEmpty
}

func TestFeedCandidatesMatchScan(t *testing.T) {
	for seed := int64(1); seed <= 5; seed++ {
		f := newFeedFixture(seed, 120)
		nonEmpty := f.compareAll(t)
		for i := 0; i < 300; i++ {
			f.mutate()
			if i%10 == 0 {
				nonEmpty += f.compareAll(t)
			}
		}
		// 再読み込み（起動時・initialize）と同じく索引を作り直しても一致する
		f.s.rebuildIndexes()
		nonEmpty += f.compareAll(t)
		if nonEmpty == 0 {
			t.Fatal("all feeds empty")
		}
	}
}

func TestFeedCandidatesConcurrent(t *testing.T) {
	f := newFeedFixture(42, 80)
	s := f.s
	var wg sync.WaitGroup
	stop := make(chan struct{})
	for g := 0; g < 4; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
				}
				s.mu.RLock()
				viewer, ok := s.loadViewerContextLocked(f.ids[g])
				if ok {
					for _, feed := range []string{"friends", "friends-of-friends", "public", "following"} {
						s.feedCandidatesLocked(viewer, map[string]bool{}, feed, feed == "friends-of-friends", nil, "")
					}
				}
				s.mu.RUnlock()
			}
		}(g)
	}
	for i := 0; i < 500; i++ {
		s.mu.Lock()
		f.mutate()
		s.mu.Unlock()
	}
	close(stop)
	wg.Wait()
}

func benchmarkFeed(b *testing.B, scan bool) {
	f := newFeedFixture(7, 1500)
	s := f.s
	viewers := make([]viewerContext, 0, len(f.ids))
	for _, uid := range f.ids {
		if v, ok := s.loadViewerContextLocked(uid); ok && v.HasPostedToday {
			viewers = append(viewers, v)
		}
	}
	feeds := []string{"friends", "friends-of-friends", "public", "following"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := viewers[i%len(viewers)]
		feed := feeds[i%len(feeds)]
		following := map[string]bool{}
		for id := range s.users[v.ID].following {
			following[id] = true
		}
		if scan {
			s.feedCandidatesScanLocked(v, following, feed, feed == "friends-of-friends", nil, "")
		} else {
			s.feedCandidatesLocked(v, following, feed, feed == "friends-of-friends", nil, "")
		}
	}
}

func BenchmarkFeedScan(b *testing.B)  { benchmarkFeed(b, true) }
func BenchmarkFeedIndex(b *testing.B) { benchmarkFeed(b, false) }
