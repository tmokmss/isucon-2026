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
		rw := s.commitLocked(&opFriendshipRemove{A: actor, B: input.UserID})
		s.mu.Unlock()
		rw.wait()
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
		rw := s.commitLocked(&opFollowAdd{F: f})
		s.mu.Unlock()
		rw.wait()
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
		rw := s.commitLocked(&opFollowRemove{Follower: actor, Followee: input.UserID})
		s.mu.Unlock()
		rw.wait()
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
		rw := s.commitLocked(newOpBlockAddLocked(s, b))
		s.mu.Unlock()
		rw.wait()
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
		rw := s.commitLocked(&opBlockRemove{Blocker: actor, Blocked: target})
		s.mu.Unlock()
		rw.wait()
		return c.NoContent(http.StatusNoContent)
	}
}
