package isureal

import (
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
		width, height, err := checkJPEG(data)
		if err != nil {
			return NewAPIError(http.StatusBadRequest, "invalid_realmoji")
		}
		if width > 512 || height > 512 {
			return NewAPIError(http.StatusBadRequest, "invalid_realmoji")
		}

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
		unlock := lockInteraction(actor, p.AuthorID)
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
		if err := storeNewImage(r.ImageID, data); err != nil {
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
		for _, old := range replaced {
			imagesInvalidated(old.ImageID)
		}
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
		s.mu.Lock()
		r.DeletedAt = &deletedAt
		s.mu.Unlock()
		imagesInvalidated(r.ImageID)
		return c.NoContent(http.StatusNoContent)
	}
}
