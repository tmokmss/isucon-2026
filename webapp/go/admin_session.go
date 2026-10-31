package isureal

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
)

type AdminSession struct {
	Token           string
	AdministratorID string
	TokenHash       string
	CreatedAt       time.Time
}

func GetAdminSession(c *echo.Context) *AdminSession {
	session, _ := c.Get("admin_session").(*AdminSession)
	return session
}

func getAdminSessionFromHeader(c *echo.Context, s *Store) (*AdminSession, error) {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	if authHeader == "" {
		return nil, nil
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, NewAPIError(http.StatusUnauthorized, "invalid_administrator_session")
	}
	hash := hashToken(strings.TrimPrefix(authHeader, "Bearer "))
	s.mu.RLock()
	defer s.mu.RUnlock()
	row := s.adminSessions[hash]
	if row == nil || !row.Active {
		return nil, nil
	}
	return &AdminSession{AdministratorID: row.AdministratorID, TokenHash: row.TokenHash, CreatedAt: row.CreatedAt}, nil
}

func AdminSessionMiddleware(h *HandlerContext, authRequired bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			session, err := getAdminSessionFromHeader(c, h.Store)
			if err != nil {
				return err
			}
			if session == nil && authRequired {
				return NewAPIError(http.StatusUnauthorized, "invalid_administrator_session")
			}
			c.Set("admin_session", session)
			return next(c)
		}
	}
}
