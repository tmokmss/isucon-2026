package isureal

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func deleteAccountHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
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
		s.mu.Lock()
		rw := s.commitLocked(&opUserDelete{UserID: userID})
		s.mu.Unlock()
		rw.wait()
		userImagesInvalidated(userID)
		return c.NoContent(http.StatusNoContent)
	}
}
