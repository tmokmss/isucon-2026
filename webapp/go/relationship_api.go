package isureal

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type deleteFriendshipRequest struct {
	UserID string `param:"userId" json:"-"`
}

func deleteFriendshipHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input deleteFriendshipRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		actor := GetSession(c).UserID
		unlock, err := h.lockRelationship(actor, input.UserID)
		if err != nil {
			return err
		}
		defer unlock()
		s := h.Store
		s.mu.RLock()
		exists := s.users[actor].friends[input.UserID] != nil
		s.mu.RUnlock()
		if !exists {
			return NewAPIError(http.StatusConflict, "relationship_conflict")
		}
		s.mu.Lock()
		s.removeFriendshipLocked(actor, input.UserID)
		s.mu.Unlock()
		return c.NoContent(http.StatusNoContent)
	}
}

type followRequest struct {
	UserID string `param:"userId" json:"-"`
}

func followHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input followRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		actor := GetSession(c).UserID
		unlock, err := h.lockRelationship(actor, input.UserID)
		if err != nil {
			return err
		}
		defer unlock()
		s := h.Store
		s.mu.RLock()
		target := s.users[input.UserID]
		certified := target.AccountType == "certified_public"
		blocked := s.isBlocked(actor, target.ID)
		exists := s.users[actor].following[target.ID] != nil
		s.mu.RUnlock()
		if !certified || blocked || exists {
			return NewAPIError(http.StatusConflict, "relationship_conflict")
		}
		at := now()
		f := &FollowRow{ID: uuid.NewString(), FollowerID: actor, FolloweeID: target.ID, CreatedAt: at, UpdatedAt: at}
		s.mu.Lock()
		s.addFollowLocked(f)
		s.mu.Unlock()
		return c.NoContent(http.StatusNoContent)
	}
}

type unfollowRequest struct {
	UserID string `param:"userId" json:"-"`
}

func unfollowHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input unfollowRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		actor := GetSession(c).UserID
		unlock, err := h.lockRelationship(actor, input.UserID)
		if err != nil {
			return err
		}
		defer unlock()
		s := h.Store
		s.mu.RLock()
		exists := s.users[actor].following[input.UserID] != nil
		s.mu.RUnlock()
		if !exists {
			return NewAPIError(http.StatusConflict, "relationship_conflict")
		}
		s.mu.Lock()
		s.removeFollowLocked(actor, input.UserID)
		s.mu.Unlock()
		return c.NoContent(http.StatusNoContent)
	}
}

type blockRequest struct {
	UserID string `param:"userId" json:"-"`
}

func blockHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input blockRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		actor := GetSession(c).UserID
		target := input.UserID
		unlock, err := h.lockRelationship(actor, target)
		if err != nil {
			return err
		}
		defer unlock()
		s := h.Store
		s.mu.RLock()
		exists := s.users[actor].blockOut[target] != nil
		pending := s.pendingBetweenLocked(actor, target)
		s.mu.RUnlock()
		if exists {
			return NewAPIError(http.StatusConflict, "relationship_conflict")
		}
		at := now()
		b := &BlockRow{ID: uuid.NewString(), BlockerID: actor, BlockedID: target, CreatedAt: at}
		if pending != nil {
		}
		s.mu.Lock()
		s.blocks = append(s.blocks, b)
		s.users[actor].blockOut[target] = b
		s.users[target].blockIn[actor] = b
		s.removeFriendshipLocked(actor, target)
		s.removeFollowLocked(actor, target)
		s.removeFollowLocked(target, actor)
		for _, r := range []*FriendRequest{s.users[actor].pendingOut[target], s.users[actor].pendingIn[target]} {
			if r != nil {
				s.setRequestStatusLocked(r, "cancelled", at)
			}
		}
		s.mu.Unlock()
		return c.NoContent(http.StatusNoContent)
	}
}

type unblockRequest struct {
	UserID string `param:"userId" json:"-"`
}

func unblockHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input unblockRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		actor := GetSession(c).UserID
		target := input.UserID
		unlock, err := h.lockRelationship(actor, target)
		if err != nil {
			return err
		}
		defer unlock()
		s := h.Store
		s.mu.RLock()
		b := s.users[actor].blockOut[target]
		s.mu.RUnlock()
		if b == nil {
			return NewAPIError(http.StatusForbidden, "relationship_action_forbidden")
		}
		s.mu.Lock()
		delete(s.users[actor].blockOut, target)
		delete(s.users[target].blockIn, actor)
		for i, x := range s.blocks {
			if x == b {
				s.blocks[i] = s.blocks[len(s.blocks)-1]
				s.blocks = s.blocks[:len(s.blocks)-1]
				break
			}
		}
		s.mu.Unlock()
		return c.NoContent(http.StatusNoContent)
	}
}
