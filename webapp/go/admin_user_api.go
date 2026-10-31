package isureal

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type certifyUserRequest struct {
	UserID string `param:"userId" json:"-"`
}

func certifyUserHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input certifyUserRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		s := h.Store
		unlock := lockStatus(input.UserID)
		defer unlock()
		s.mu.RLock()
		u := s.activeUser(input.UserID)
		ok := u != nil && u.AccountType == "regular"
		s.mu.RUnlock()
		if !ok {
			return NewAPIError(http.StatusConflict, "certification_conflict")
		}
		s.mu.Lock()
		rw := s.commitLocked(&opCertify{UserID: u.ID})
		s.mu.Unlock()
		rw.wait()
		return c.NoContent(http.StatusNoContent)
	}
}

// 2026-08-26: 認定されたユーザーを戻すことは禁止する
func deleteUserCertificationHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return NewAPIError(http.StatusConflict, "certification_irreversible")
	}
}
