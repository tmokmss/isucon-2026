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
		s.logLocalLocked(&opAdminLogin{Row: row})
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
			s.logLocalLocked(&opAdminLogout{TokenHash: session.TokenHash})
		}
		s.mu.Unlock()
		return c.NoContent(http.StatusNoContent)
	}
}

// 管理者のセッションは asia1 だけが使うので、複製へは流さず変更の記録にだけ残す
type opAdminLogin struct{ Row *AdminSessionRow }

func (o *opAdminLogin) apply(s *Store) { s.adminSessions[o.Row.TokenHash] = o.Row }

type opAdminLogout struct{ TokenHash string }

func (o *opAdminLogout) apply(s *Store) {
	if row := s.adminSessions[o.TokenHash]; row != nil {
		row.Active = false
		delete(s.adminSessions, o.TokenHash)
	}
}

// logLocalLocked は op をこのストアに当て、変更の記録に残す。s.mu の書きロックを持って呼ぶ
func (s *Store) logLocalLocked(op replOp) {
	op.apply(s)
	opLog.appendLocked(op)
}
