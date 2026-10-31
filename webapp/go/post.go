package isureal

import (
	"time"
)

type PostLocation struct {
	Kind      string   `json:"kind" db:"location_kind"`
	Latitude  *float64 `json:"latitude,omitempty" db:"latitude"`
	Longitude *float64 `json:"longitude,omitempty" db:"longitude"`
	Locality  *string  `json:"locality,omitempty" db:"locality"`
}

func (location PostLocation) valid() bool {
	switch location.Kind {
	case "none":
		return location.Latitude == nil && location.Longitude == nil && location.Locality == nil
	case "exact":
		return location.Latitude != nil && location.Longitude != nil && location.Locality == nil &&
			*location.Latitude >= -90 && *location.Latitude <= 90 && *location.Longitude >= -180 && *location.Longitude <= 180
	case "approximate":
		return location.Latitude == nil && location.Longitude == nil && location.Locality != nil && validUserText(*location.Locality, 1, 100)
	default:
		return false
	}
}

type Post struct {
	ID              string
	AuthorID        string
	PostingWindowID string
	ImageID         string
	Caption         string
	Visibility      string
	PostLocation
	PostedAt       time.Time
	OnTimeDeadline time.Time
	WindowEndedAt  *time.Time
}

// viewerContext は閲覧者の友達・ブロックを手元の UserRow から直接引く。読みロックを持っている間だけ使う。
// 友達の友達と共通の友達の数は、使うときに一度だけ作る
type viewerContext struct {
	User
	HasPostedToday bool
	u              *UserRow
	s              *Store
	fof            *viewerFoF
}

type viewerFoF struct {
	all    map[string]bool
	mutual map[string]int
}

func (v viewerContext) isFriend(id string) bool {
	_, ok := v.u.friends[id]
	return ok
}

func (v viewerContext) isBlocked(id string) bool {
	if _, ok := v.u.blockOut[id]; ok {
		return true
	}
	_, ok := v.u.blockIn[id]
	return ok
}

func (v viewerContext) friendsOfFriends() *viewerFoF {
	if v.fof.all != nil {
		return v.fof
	}
	v.fof.all = map[string]bool{}
	v.fof.mutual = map[string]int{}
	for id := range v.u.friends {
		f := v.s.users[id]
		if f == nil {
			continue
		}
		for g := range f.friends {
			v.fof.all[g] = true
			if g != v.ID {
				v.fof.mutual[g]++
			}
		}
	}
	return v.fof
}

// livePostLocked は削除されていない投稿を返す
func (s *Store) livePostLocked(postID string) *PostRow {
	p := s.posts[postID]
	if p == nil || p.DeletedAt != nil {
		return nil
	}
	return p
}

// livePostWithAuthorLocked は削除されておらず、投稿者が有効な投稿を返す
func (s *Store) livePostWithAuthorLocked(postID string) (*PostRow, *UserRow) {
	p := s.livePostLocked(postID)
	if p == nil {
		return nil, nil
	}
	u := s.activeUser(p.AuthorID)
	if u == nil {
		return nil, nil
	}
	return p, u
}

// hasPostedTodayLocked は、利用者のリージョンの現在の投稿期間と同じ日付の投稿があるか
func (s *Store) hasPostedTodayLocked(u *UserRow) bool {
	return s.currentPostLocked(u) != nil
}

func (s *Store) currentPostLocked(u *UserRow) *PostRow {
	w := s.activeWindows[u.Region]
	if w == nil {
		return nil
	}
	for i := len(u.posts) - 1; i >= 0; i-- {
		p := u.posts[i]
		if p.DeletedAt == nil && p.window.LocalDate.Equal(w.LocalDate) {
			return p
		}
	}
	return nil
}

// postCountsLocked は利用者のその日付の投稿の数（有効・削除済み）
func postCountsLocked(u *UserRow, localDate time.Time) (active, deleted int) {
	for _, p := range u.posts {
		if p.window.LocalDate.Equal(localDate) {
			if p.DeletedAt == nil {
				active++
			} else {
				deleted++
			}
		}
	}
	return
}

func (s *Store) loadViewerContextLocked(viewerID string) (viewerContext, bool) {
	u := s.activeUser(viewerID)
	if u == nil {
		return viewerContext{}, false
	}
	return viewerContext{User: u.User, HasPostedToday: s.hasPostedTodayLocked(u), u: u, s: s, fof: &viewerFoF{}}, true
}

func (viewer viewerContext) canView(post Post) bool {
	if post.AuthorID == viewer.ID {
		return true
	}
	if post.WindowEndedAt != nil || !viewer.HasPostedToday || viewer.isBlocked(post.AuthorID) {
		return false
	}
	switch post.Visibility {
	case "public":
		return true
	case "friends":
		return viewer.isFriend(post.AuthorID)
	case "friends_of_friends":
		return viewer.isFriend(post.AuthorID) || viewer.friendsOfFriends().all[post.AuthorID]
	default:
		return false
	}
}

func (s *Store) canViewPostLocked(viewerID, postID string) bool {
	p, _ := s.livePostWithAuthorLocked(postID)
	if p == nil {
		return false
	}
	viewer, ok := s.loadViewerContextLocked(viewerID)
	if !ok {
		return false
	}
	return viewer.canView(p.post())
}
