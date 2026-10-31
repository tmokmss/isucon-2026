package isureal

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type adminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type adminLoginResponse struct {
	Token           string `json:"token"`
	AdministratorID string `json:"administratorId"`
}

func adminLoginHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input adminLoginRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		if !usernamePattern.MatchString(input.Username) {
			return NewAPIError(http.StatusUnauthorized, "invalid_credentials")
		}
		s := h.Store
		s.mu.RLock()
		admin := s.admins[input.Username]
		s.mu.RUnlock()
		if admin == nil {
			return NewAPIError(http.StatusUnauthorized, "invalid_credentials")
		}
		hash := sha256.Sum256([]byte(input.Password))
		if admin.PasswordHash != hex.EncodeToString(hash[:]) {
			return NewAPIError(http.StatusUnauthorized, "invalid_credentials")
		}
		token := uuid.NewString()
		row := &AdminSessionRow{ID: uuid.NewString(), AdministratorID: admin.ID, TokenHash: hashToken(token), Active: true, CreatedAt: now()}
		s.mu.Lock()
		s.adminSessions[row.TokenHash] = row
		s.needFull.Store(true)
		s.mu.Unlock()
		return c.JSON(http.StatusOK, adminLoginResponse{Token: token, AdministratorID: admin.ID})
	}
}

func adminLogoutHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		session := GetAdminSession(c)
		if session == nil {
			return NewAPIError(http.StatusUnauthorized, "invalid_administrator_session")
		}
		s := h.Store
		s.mu.Lock()
		if row := s.adminSessions[session.TokenHash]; row != nil && row.AdministratorID == session.AdministratorID {
			row.Active = false
			delete(s.adminSessions, session.TokenHash)
			s.needFull.Store(true)
		}
		s.mu.Unlock()
		return c.NoContent(http.StatusNoContent)
	}
}
