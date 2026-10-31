package isureal

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func deleteAccountHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := c.Request().Context()
		userID := GetSession(c).UserID
		s := h.Store
		unlock := lockStatus(userID)
		defer unlock()
		s.mu.RLock()
		u := s.activeUser(userID)
		s.mu.RUnlock()
		if u == nil {
			return NewAPIError(http.StatusNotFound, "account_not_found")
		}
		at := now()
		tx, err := h.DB.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
		defer tx.Rollback()
		if _, err := tx.ExecContext(ctx, `UPDATE users SET status='deleted', deleted_at=$2, updated_at=$2 WHERE id=$1`, userID, at); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `UPDATE login_sessions SET active=false, invalidated_at=$2, reason='account_deletion' WHERE user_id=$1 AND active=true`, userID, at); err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		s.mu.Lock()
		rw := s.commitLocked(&opUserDelete{UserID: userID})
		s.mu.Unlock()
		rw.wait()
		userImagesInvalidated(userID)
		return c.NoContent(http.StatusNoContent)
	}
}
