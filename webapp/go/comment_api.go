package isureal

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type createCommentRequest struct {
	PostID string `param:"postId" json:"-"`
	Body   string `json:"body"`
}

type createCommentResponse struct {
	ID        string                   `json:"id"`
	Author    relationshipUserResponse `json:"author"`
	Body      string                   `json:"body"`
	CreatedAt time.Time                `json:"createdAt"`
}

func newCommentResponse(comment Comment, user User) createCommentResponse {
	return createCommentResponse{ID: comment.ID, Author: newRelationshipUserResponse(user), Body: comment.Body, CreatedAt: comment.CreatedAt}
}

func createCommentHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input createCommentRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		if !validUserText(input.Body, 1, 1000) {
			return NewAPIError(http.StatusBadRequest, "invalid_interaction")
		}
		actor := GetSession(c).UserID
		s := h.Store
		s.mu.RLock()
		p := s.livePostLocked(input.PostID)
		s.mu.RUnlock()
		if p == nil {
			return NewAPIError(http.StatusForbidden, "interaction_forbidden")
		}
		unlock := lockInteraction(actor, p.AuthorID)
		defer unlock()
		s.mu.RLock()
		allowed := s.canCommentLocked(actor, p.ID)
		u := s.activeUser(actor)
		var user User
		if u != nil {
			user = u.User
		}
		s.mu.RUnlock()
		if !allowed {
			return NewAPIError(http.StatusForbidden, "interaction_forbidden")
		}
		if u == nil {
			return NewAPIError(http.StatusInternalServerError, "user_not_found")
		}
		cm := &CommentRow{Comment: Comment{ID: uuid.NewString(), PostID: p.ID, AuthorID: actor, Body: input.Body, CreatedAt: now()}}
		s.mu.Lock()
		rw := s.commitLocked(&opCommentAdd{C: cm})
		s.mu.Unlock()
		rw.wait()
		return c.JSON(http.StatusCreated, newCommentResponse(cm.Comment, user))
	}
}

type getCommentsRequest struct {
	PostID         string     `param:"postId"`
	AfterCreatedAt *time.Time `query:"afterCreatedAt"`
	AfterID        string     `query:"afterId"`
}

type getCommentsResponse struct {
	Items   []createCommentResponse `json:"items"`
	HasMore bool                    `json:"hasMore"`
}

type commentWithAuthor struct {
	comment Comment
	author  User
}

// commentsLocked は投稿の有効なコメントを (created_at, id) の昇順で返す
func (s *Store) commentsLocked(p *PostRow, afterAt *time.Time, afterID string, limit int) []commentWithAuthor {
	out := make([]commentWithAuthor, 0)
	for _, cm := range p.comments {
		if cm.DeletedAt != nil {
			continue
		}
		if afterAt != nil && !(cm.CreatedAt.After(*afterAt) || (cm.CreatedAt.Equal(*afterAt) && cm.ID > afterID)) {
			continue
		}
		a := s.activeUser(cm.AuthorID)
		if a == nil {
			continue
		}
		out = append(out, commentWithAuthor{comment: cm.Comment, author: a.User})
		if len(out) == limit {
			break
		}
	}
	return out
}

func getCommentsHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input getCommentsRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		if (input.AfterCreatedAt == nil) != (input.AfterID == "") {
			return NewAPIError(http.StatusBadRequest, "invalid_pagination")
		}
		s := h.Store
		s.mu.RLock()
		if !s.canCommentLocked(GetSession(c).UserID, input.PostID) {
			s.mu.RUnlock()
			return NewAPIError(http.StatusForbidden, "interaction_forbidden")
		}
		comments := s.commentsLocked(s.posts[input.PostID], input.AfterCreatedAt, input.AfterID, 51)
		s.mu.RUnlock()
		items := make([]createCommentResponse, 0, len(comments))
		for _, cm := range comments {
			items = append(items, newCommentResponse(cm.comment, cm.author))
		}
		hasMore := len(items) > 50
		if hasMore {
			items = items[:50]
		}
		return c.JSON(http.StatusOK, getCommentsResponse{Items: items, HasMore: hasMore})
	}
}

type deleteCommentRequest struct {
	CommentID string `param:"commentId" json:"-"`
}

func deleteCommentHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input deleteCommentRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		actor := GetSession(c).UserID
		s := h.Store
		unlock := lockUsers(actor)
		defer unlock()
		s.mu.RLock()
		cm := s.comments[input.CommentID]
		var authorID, postAuthorID string
		if cm != nil && cm.DeletedAt == nil {
			authorID, postAuthorID = cm.AuthorID, cm.post.AuthorID
		}
		s.mu.RUnlock()
		if cm == nil || authorID == "" {
			return NewAPIError(http.StatusNotFound, "interaction_not_found")
		}
		if actor != authorID && actor != postAuthorID {
			return NewAPIError(http.StatusForbidden, "interaction_forbidden")
		}
		deletedAt := now()
		s.mu.Lock()
		rw := s.commitLocked(&opCommentDelete{ID: cm.ID, At: deletedAt})
		s.mu.Unlock()
		rw.wait()
		return c.NoContent(http.StatusNoContent)
	}
}
