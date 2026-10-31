package isureal

import (
	"net/http"
	"time"
)

type FriendRequest struct {
	ID          string
	RequesterID string
	RecipientID string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// lockRelationship は 2 人の間の関係の書き込みを直列にし、2 人とも有効かを確かめる（元の beginRelationshipTx に相当）。
// 2 人の状態（退会・認定）は共有で取るので、同じ利用者が関わる別の組の書き込みは並列に進む
func (h *HandlerContext) lockRelationship(actor, target string) (func(), error) {
	if actor == target || target == "" {
		return nil, NewAPIError(http.StatusBadRequest, "invalid_relationship")
	}
	unlockStatus := lockStripes(&statusLocks, nil, []string{actor, target})
	unlockPair := lockPair(actor, target)
	unlock := func() {
		unlockPair()
		unlockStatus()
	}
	s := h.Store
	s.mu.RLock()
	ok := s.activeUser(actor) != nil && s.activeUser(target) != nil
	s.mu.RUnlock()
	if !ok {
		unlock()
		return nil, NewAPIError(http.StatusNotFound, "relationship_not_found")
	}
	return unlock, nil
}

func (s *Store) addFriendshipLocked(f *FriendshipRow) {
	s.friendships = append(s.friendships, f)
	s.users[f.LowID].friends[f.HighID] = f
	s.users[f.HighID].friends[f.LowID] = f
}

func (s *Store) removeFriendshipLocked(a, b string) {
	lo, hi := pairKey(a, b)
	f := s.users[lo].friends[hi]
	if f == nil {
		return
	}
	delete(s.users[lo].friends, hi)
	delete(s.users[hi].friends, lo)
	for i, x := range s.friendships {
		if x == f {
			s.friendships[i] = s.friendships[len(s.friendships)-1]
			s.friendships = s.friendships[:len(s.friendships)-1]
			break
		}
	}
}

func (s *Store) addFollowLocked(f *FollowRow) {
	s.follows = append(s.follows, f)
	s.users[f.FollowerID].following[f.FolloweeID] = f
	s.users[f.FolloweeID].followers[f.FollowerID] = f
}

func (s *Store) removeFollowLocked(follower, followee string) {
	f := s.users[follower].following[followee]
	if f == nil {
		return
	}
	delete(s.users[follower].following, followee)
	delete(s.users[followee].followers, follower)
	for i, x := range s.follows {
		if x == f {
			s.follows[i] = s.follows[len(s.follows)-1]
			s.follows = s.follows[:len(s.follows)-1]
			break
		}
	}
}

func (s *Store) setRequestStatusLocked(r *FriendRequest, status string, at time.Time) {
	if r.Status == "pending" && status != "pending" {
		delete(s.users[r.RequesterID].pendingOut, r.RecipientID)
		delete(s.users[r.RecipientID].pendingIn, r.RequesterID)
	}
	r.Status = status
	r.UpdatedAt = at
}

// pendingBetweenLocked は 2 人の間の保留中の友達申請を返す
func (s *Store) pendingBetweenLocked(a, b string) *FriendRequest {
	u := s.users[a]
	if r := u.pendingOut[b]; r != nil {
		return r
	}
	return u.pendingIn[b]
}
