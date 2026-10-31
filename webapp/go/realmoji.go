package isureal

import (
	"time"
)

type RealMoji struct {
	ID           string
	PostID       string
	ReactorID    string
	ImageID      string
	ReactionType string
	CreatedAt    time.Time
	ReplacedAt   *time.Time
	DeletedAt    *time.Time
}

type realMojiWithReactor struct {
	reaction RealMoji
	reactor  User
}

// currentRealMojisLocked は投稿の現在の RealMoji を (created_at, id) の降順で返す
func (s *Store) currentRealMojisLocked(p *PostRow, beforeAt *time.Time, beforeID string, limit int) []realMojiWithReactor {
	out := make([]realMojiWithReactor, 0)
	for i := len(p.realmojis) - 1; i >= 0; i-- {
		r := p.realmojis[i]
		if !r.current() {
			continue
		}
		if beforeAt != nil && !(r.CreatedAt.Before(*beforeAt) || (r.CreatedAt.Equal(*beforeAt) && r.ID < beforeID)) {
			continue
		}
		u := s.activeUser(r.ReactorID)
		if u == nil {
			continue
		}
		out = append(out, realMojiWithReactor{reaction: r.RealMoji, reactor: u.User})
		if len(out) == limit {
			break
		}
	}
	return out
}
