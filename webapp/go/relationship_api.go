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
		lo, hi := pairKey(actor, input.UserID)
		if _, err := h.DB.ExecContext(c.Request().Context(), `DELETE FROM friendships WHERE user_low_id=$1 AND user_high_id=$2`, lo, hi); err != nil {
			return err
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
		if _, err := h.DB.ExecContext(c.Request().Context(), `INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$4)`, f.ID, actor, target.ID, at); err != nil {
			return err
		}
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
		if _, err := h.DB.ExecContext(c.Request().Context(), `DELETE FROM follows WHERE follower_id=$1 AND followee_id=$2`, actor, input.UserID); err != nil {
			return err
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
		ctx := c.Request().Context()
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
		lo, hi := pairKey(actor, target)
		tx, err := h.DB.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
		defer tx.Rollback()
		if _, err := tx.ExecContext(ctx, `INSERT INTO block_relations (id, blocker_id, blocked_id, created_at) VALUES ($1,$2,$3,$4)`, b.ID, actor, target, at); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `DELETE FROM friendships WHERE user_low_id=$1 AND user_high_id=$2`, lo, hi); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `DELETE FROM follows WHERE (follower_id=$1 AND followee_id=$2) OR (follower_id=$2 AND followee_id=$1)`, actor, target); err != nil {
			return err
		}
		if pending != nil {
			if _, err := tx.ExecContext(ctx, `UPDATE friend_requests SET status='cancelled', updated_at=$3 WHERE status='pending' AND ((requester_id=$1 AND recipient_id=$2) OR (requester_id=$2 AND recipient_id=$1))`, actor, target, at); err != nil {
				return err
			}
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		s.mu.Lock()
		rw := s.commitLocked(&opBlockAdd{B: b})
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
		if _, err := h.DB.ExecContext(c.Request().Context(), `DELETE FROM block_relations WHERE blocker_id=$1 AND blocked_id=$2`, actor, target); err != nil {
			return err
		}
		s.mu.Lock()
		rw := s.commitLocked(&opBlockRemove{Blocker: actor, Blocked: target})
		s.mu.Unlock()
		rw.wait()
		return c.NoContent(http.StatusNoContent)
	}
}
