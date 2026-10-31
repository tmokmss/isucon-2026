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
		u.Status = "deleted"
		s.refreshUserSearchLocked(u)
		for hash, sess := range u.sessions {
			sess.Active = false
			delete(s.sessions, hash)
		}
		u.sessions = map[string]*LoginSessionRow{}
		s.mu.Unlock()
		userImagesInvalidated(userID)
		return c.NoContent(http.StatusNoContent)
	}
}
