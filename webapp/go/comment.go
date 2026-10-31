package isureal

import (
	"time"
)

type Comment struct {
	ID        string
	PostID    string
	AuthorID  string
	Body      string
	CreatedAt time.Time
	DeletedAt *time.Time
}

func (s *Store) canCommentLocked(viewerID, postID string) bool {
	p, _ := s.livePostWithAuthorLocked(postID)
	if p == nil {
		return false
	}
	viewer, ok := s.loadViewerContextLocked(viewerID)
	if !ok {
		return false
	}
	post := p.post()
	if !viewer.canView(post) || post.WindowEndedAt != nil {
		return false
	}
	if viewer.ID == post.AuthorID {
		return viewer.HasPostedToday
	}
	if post.Visibility == "public" {
		return true
	}
	return viewer.Friends[post.AuthorID]
}
