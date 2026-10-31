package isureal

import (
	"bytes"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type replaceRealMojiRequest struct {
	PostID       string                `param:"postId" json:"-"`
	Image        *multipart.FileHeader `form:"image"`
	ReactionType string                `form:"reactionType"`
}

type replaceRealMojiResponse struct {
	ID           string                   `json:"id"`
	Reactor      relationshipUserResponse `json:"reactor"`
	ReactionType string                   `json:"reactionType"`
	ImageURL     string                   `json:"imageUrl"`
	CreatedAt    time.Time                `json:"createdAt"`
}

func newRealMojiResponse(reaction RealMoji, user User) replaceRealMojiResponse {
	return replaceRealMojiResponse{ID: reaction.ID, Reactor: newRelationshipUserResponse(user), ReactionType: reaction.ReactionType, ImageURL: signedImageURL(user, reaction.ImageID), CreatedAt: reaction.CreatedAt}
}

func replaceRealMojiHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input replaceRealMojiRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		if input.Image == nil {
			return NewAPIError(http.StatusBadRequest, "invalid_realmoji")
		}
		switch input.ReactionType {
		case "like", "smile", "surprise", "heart", "laugh", "instant":
		default:
			return NewAPIError(http.StatusBadRequest, "invalid_realmoji")
		}
		file, err := input.Image.Open()
		if err != nil {
			return err
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}
		img, err := jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			return NewAPIError(http.StatusBadRequest, "invalid_realmoji")
		}
		if img.Bounds().Dx() > 512 || img.Bounds().Dy() > 512 {
			return NewAPIError(http.StatusBadRequest, "invalid_realmoji")
		}

		ctx := c.Request().Context()
		actor := GetSession(c).UserID
		s := h.Store
		s.mu.RLock()
		p := s.livePostLocked(input.PostID)
		s.mu.RUnlock()
		if p == nil {
			return NewAPIError(http.StatusForbidden, "interaction_forbidden")
		}
		if actor == p.AuthorID {
			return NewAPIError(http.StatusForbidden, "interaction_forbidden")
		}
		unlock := lockUsers(actor, p.AuthorID)
		defer unlock()
		s.mu.RLock()
		actorRow, authorRow := s.activeUser(actor), s.activeUser(p.AuthorID)
		if actorRow == nil || authorRow == nil {
			s.mu.RUnlock()
			return NewAPIError(http.StatusNotFound, "relationship_not_found")
		}
		allowed := s.canViewPostLocked(actor, p.ID)
		user := actorRow.User
		var replaced []*RealMojiRow
		for _, r := range p.realmojis {
			if r.ReactorID == actor && r.current() {
				replaced = append(replaced, r)
			}
		}
		s.mu.RUnlock()
		if !allowed {
			return NewAPIError(http.StatusForbidden, "interaction_forbidden")
		}
		createdAt := now()
		r := &RealMojiRow{RealMoji: RealMoji{ID: uuid.NewString(), PostID: p.ID, ReactorID: actor, ImageID: uuid.NewString(), ReactionType: input.ReactionType, CreatedAt: createdAt}}
		tx, err := h.DB.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
		defer tx.Rollback()
		if len(replaced) > 0 {
			if _, err := tx.ExecContext(ctx, `UPDATE realmojis SET replaced_at=$3 WHERE post_id=$1 AND reactor_id=$2 AND replaced_at IS NULL AND deleted_at IS NULL`, p.ID, actor, createdAt); err != nil {
				return err
			}
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES($1,$2,$3,'image/jpeg',$4)`, r.ImageID, actor, data, createdAt); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)`, r.ID, p.ID, actor, r.ImageID, r.ReactionType, createdAt); err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		imageCache.put(r.ImageID, data, "image/jpeg")
		s.mu.Lock()
		for _, old := range replaced {
			t := createdAt
			old.ReplacedAt = &t
		}
		s.addRealMojiLocked(r)
		s.mu.Unlock()
		return c.JSON(http.StatusCreated, newRealMojiResponse(r.RealMoji, user))
	}
}

type getRealMojisRequest struct {
	PostID          string     `param:"postId"`
	BeforeCreatedAt *time.Time `query:"beforeCreatedAt"`
	BeforeID        string     `query:"beforeId"`
}

type getRealMojisResponse struct {
	Items   []replaceRealMojiResponse `json:"items"`
	HasMore bool                      `json:"hasMore"`
}

func getRealMojisHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input getRealMojisRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		if (input.BeforeCreatedAt == nil) != (input.BeforeID == "") {
			return NewAPIError(http.StatusBadRequest, "invalid_pagination")
		}
		s := h.Store
		s.mu.RLock()
		if !s.canViewPostLocked(GetSession(c).UserID, input.PostID) {
			s.mu.RUnlock()
			return NewAPIError(http.StatusForbidden, "interaction_forbidden")
		}
		reactions := s.currentRealMojisLocked(s.posts[input.PostID], input.BeforeCreatedAt, input.BeforeID, 51)
		s.mu.RUnlock()
		items := make([]replaceRealMojiResponse, 0, len(reactions))
		for _, reaction := range reactions {
			items = append(items, newRealMojiResponse(reaction.reaction, reaction.reactor))
		}
		hasMore := len(items) > 50
		if hasMore {
			items = items[:50]
		}
		return c.JSON(http.StatusOK, getRealMojisResponse{Items: items, HasMore: hasMore})
	}
}

type deleteRealMojiRequest struct {
	RealMojiID string `param:"realMojiId" json:"-"`
}

func deleteRealMojiHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input deleteRealMojiRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		ctx := c.Request().Context()
		actor := GetSession(c).UserID
		s := h.Store
		unlock := lockUsers(actor)
		defer unlock()
		s.mu.RLock()
		r := s.realmojis[input.RealMojiID]
		var reactorID, authorID string
		if r != nil && r.DeletedAt == nil {
			reactorID, authorID = r.ReactorID, r.post.AuthorID
		}
		s.mu.RUnlock()
		if reactorID == "" {
			return NewAPIError(http.StatusNotFound, "interaction_not_found")
		}
		if actor != reactorID && actor != authorID {
			return NewAPIError(http.StatusForbidden, "interaction_forbidden")
		}
		deletedAt := now()
		if _, err := h.DB.ExecContext(ctx, `UPDATE realmojis SET deleted_at=$2 WHERE id=$1`, r.ID, deletedAt); err != nil {
			return err
		}
		s.mu.Lock()
		r.DeletedAt = &deletedAt
		s.mu.Unlock()
		return c.NoContent(http.StatusNoContent)
	}
}
