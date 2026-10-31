package isureal

import (
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type createFriendRequestRequest struct {
	RecipientID string `json:"recipientId"`
}

type createFriendRequestResponse struct {
	RelationID string    `json:"relationId"`
	Status     string    `json:"status"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func createFriendRequestHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input createFriendRequestRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		ctx := c.Request().Context()
		actor := GetSession(c).UserID
		recipient := input.RecipientID
		unlock, err := h.lockRelationship(actor, recipient)
		if err != nil {
			return err
		}
		defer unlock()
		s := h.Store
		s.mu.RLock()
		blocked := s.isBlocked(actor, recipient)
		isFriend := s.users[actor].friends[recipient] != nil
		existing := s.pendingBetweenLocked(actor, recipient)
		var existingRequester string
		if existing != nil {
			existingRequester = existing.RequesterID
		}
		s.mu.RUnlock()
		if blocked || isFriend {
			return NewAPIError(http.StatusConflict, "relationship_conflict")
		}
		at := now()
		if existing != nil {
			if existingRequester == actor {
				return NewAPIError(http.StatusConflict, "relationship_conflict")
			}
			lo, hi := pairKey(existing.RequesterID, existing.RecipientID)
			srcID := existing.ID
			f := &FriendshipRow{ID: uuid.NewString(), LowID: lo, HighID: hi, SourceRequestID: &srcID, CreatedAt: at, UpdatedAt: at}
			tx, err := h.DB.BeginTxx(ctx, nil)
			if err != nil {
				return err
			}
			defer tx.Rollback()
			if _, err := tx.ExecContext(ctx, `UPDATE friend_requests SET status='accepted', updated_at=$2 WHERE id=$1`, existing.ID, at); err != nil {
				return err
			}
			if _, err := tx.ExecContext(ctx, `INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)`, f.ID, lo, hi, srcID, at); err != nil {
				return err
			}
			if err := tx.Commit(); err != nil {
				return err
			}
			s.mu.Lock()
			rw := s.commitLocked(&opRequestStatus{ID: existing.ID, Status: "accepted", At: at, F: f})
			s.mu.Unlock()
			rw.wait()
			return c.JSON(http.StatusCreated, createFriendRequestResponse{RelationID: existing.ID, Status: "accepted", UpdatedAt: at})
		}
		r := &FriendRequest{ID: uuid.NewString(), RequesterID: actor, RecipientID: recipient, Status: "pending", CreatedAt: at, UpdatedAt: at}
		if _, err := h.DB.ExecContext(ctx, `INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,'pending',$4,$4)`, r.ID, actor, recipient, at); err != nil {
			return err
		}
		s.mu.Lock()
		rw := s.commitLocked(&opFriendRequestAdd{R: r})
		s.mu.Unlock()
		rw.wait()
		return c.JSON(http.StatusCreated, createFriendRequestResponse{RelationID: r.ID, Status: r.Status, UpdatedAt: r.UpdatedAt})
	}
}

type transitionFriendRequestRequest struct {
	RequestID string `param:"requestId" json:"-"`
	Action    string `param:"action" json:"-"`
}

type transitionFriendRequestResponse createFriendRequestResponse

func transitionFriendRequestHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input transitionFriendRequestRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		var status string
		switch input.Action {
		case "accept":
			status = "accepted"
		case "reject":
			status = "rejected"
		case "cancel":
			status = "cancelled"
		default:
			return NewAPIError(http.StatusForbidden, "relationship_action_forbidden")
		}
		ctx := c.Request().Context()
		actor := GetSession(c).UserID
		s := h.Store
		s.mu.RLock()
		r := s.friendRequests[input.RequestID]
		var requesterID, recipientID string
		if r != nil {
			requesterID, recipientID = r.RequesterID, r.RecipientID
		}
		s.mu.RUnlock()
		if r == nil {
			return NewAPIError(http.StatusNotFound, "relationship_not_found")
		}
		if (status == "cancelled" && actor != requesterID) || (status != "cancelled" && actor != recipientID) {
			return NewAPIError(http.StatusForbidden, "relationship_action_forbidden")
		}
		unlock, err := h.lockRelationship(requesterID, recipientID)
		if err != nil {
			return err
		}
		defer unlock()
		s.mu.RLock()
		pending := r.Status == "pending"
		blocked := s.isBlocked(requesterID, recipientID)
		friendExists := s.users[requesterID].friends[recipientID] != nil
		s.mu.RUnlock()
		if !pending {
			return NewAPIError(http.StatusForbidden, "relationship_action_forbidden")
		}
		at := now()
		var f *FriendshipRow
		tx, err := h.DB.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
		defer tx.Rollback()
		if status == "accepted" {
			if blocked {
				return NewAPIError(http.StatusConflict, "relationship_conflict")
			}
			if friendExists {
				return errors.New("friendship already exists")
			}
			lo, hi := pairKey(requesterID, recipientID)
			srcID := r.ID
			f = &FriendshipRow{ID: uuid.NewString(), LowID: lo, HighID: hi, SourceRequestID: &srcID, CreatedAt: at, UpdatedAt: at}
			if _, err := tx.ExecContext(ctx, `INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)`, f.ID, lo, hi, srcID, at); err != nil {
				return err
			}
		}
		if _, err := tx.ExecContext(ctx, `UPDATE friend_requests SET status=$2, updated_at=$3 WHERE id=$1`, r.ID, status, at); err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		s.mu.Lock()
		rw := s.commitLocked(&opRequestStatus{ID: r.ID, Status: status, At: at, F: f})
		s.mu.Unlock()
		rw.wait()
		return c.JSON(http.StatusOK, transitionFriendRequestResponse{RelationID: r.ID, Status: status, UpdatedAt: at})
	}
}
