package isureal

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type registerRequest struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Region      string `json:"region"`
}

type registerResponse struct {
	SecretKey string              `json:"secretKey"`
	User      sessionUserResponse `json:"user"`
}

type loginRequest struct {
	Username  string `json:"username"`
	SecretKey string `json:"secretKey"`
}

type loginResponse struct {
	Token string              `json:"token"`
	User  sessionUserResponse `json:"user"`
}

type sessionUserResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Bio         string `json:"bio"`
	Region      string `json:"region"`
	AccountType string `json:"accountType"`
	Status      string `json:"status"`
}

type getMeResponse struct {
	sessionUserResponse
	ProfileImageURL *string `json:"profileImageUrl"`
}

type updateMeRequest struct {
	DisplayName *string `json:"displayName"`
	Bio         *string `json:"bio"`
	Region      *string `json:"region"`
}

type updateMeResponse getMeResponse

func newSessionUserResponse(user User) sessionUserResponse {
	return sessionUserResponse{
		ID: user.ID, Username: user.Username, DisplayName: user.DisplayName,
		Bio: user.Bio, Region: user.Region, AccountType: user.AccountType, Status: user.Status,
	}
}

func profileImageURL(user User) *string {
	if user.ProfileImageID == nil {
		return nil
	}
	url := signedImageURL(user, *user.ProfileImageID)
	return &url
}

func newGetMeResponse(user User) getMeResponse {
	return getMeResponse{sessionUserResponse: newSessionUserResponse(user), ProfileImageURL: profileImageURL(user)}
}

// ユーザー名の一意性を守るため、登録は直列にする
var registerMu sync.Mutex

func registerHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input registerRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		c.Set("username", input.Username)
		if !usernamePattern.MatchString(input.Username) || !validUserText(input.DisplayName, 1, 100) || !validRegion(input.Region) {
			return NewAPIError(http.StatusBadRequest, "invalid_registration")
		}
		secretKey := uuid.NewString()
		hash := sha256.Sum256([]byte(secretKey))
		ctx := c.Request().Context()
		s := h.Store
		registerMu.Lock()
		defer registerMu.Unlock()
		s.mu.RLock()
		exists := s.usersByName[input.Username] != nil
		s.mu.RUnlock()
		if exists {
			return NewAPIError(http.StatusConflict, "username_taken")
		}
		userID := uuid.NewString()
		createdAt := now()
		u := &UserRow{
			User: User{
				ID: userID, Username: input.Username, SecretKeyHash: hex.EncodeToString(hash[:]), DisplayName: input.DisplayName,
				Bio: "", Region: input.Region, AccountType: "regular", Status: "active",
			},
			CreatedAt: createdAt,
		}
		if err := u.setImageKey(); err != nil {
			return err
		}
		u.setNormalized()
		if _, err := h.DB.ExecContext(ctx, `INSERT INTO users (id, username, secret_key_hash, display_name, region, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $6)`,
			userID, input.Username, u.SecretKeyHash, input.DisplayName, input.Region, createdAt); err != nil {
			return err
		}
		s.mu.Lock()
		s.addUserLocked(u)
		user := u.User
		s.mu.Unlock()
		return c.JSON(http.StatusCreated, registerResponse{SecretKey: secretKey, User: newSessionUserResponse(user)})
	}
}

func loginHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input loginRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		c.Set("username", input.Username)
		if !usernamePattern.MatchString(input.Username) {
			return NewAPIError(http.StatusUnauthorized, "invalid_credentials")
		}
		ctx := c.Request().Context()
		s := h.Store
		s.mu.RLock()
		u := s.usersByName[input.Username]
		s.mu.RUnlock()
		if u == nil {
			return NewAPIError(http.StatusUnauthorized, "invalid_credentials")
		}
		unlock := lockUsers(u.ID)
		defer unlock()
		s.mu.RLock()
		user := u.User
		s.mu.RUnlock()
		hash := sha256.Sum256([]byte(input.SecretKey))
		if user.Status != "active" || user.SecretKeyHash != hex.EncodeToString(hash[:]) {
			return NewAPIError(http.StatusUnauthorized, "invalid_credentials")
		}
		token := uuid.NewString()
		row := &LoginSessionRow{ID: uuid.NewString(), UserID: user.ID, TokenHash: hashToken(token), Active: true, CreatedAt: now()}
		tx, err := h.DB.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
		defer tx.Rollback()
		if _, err := tx.ExecContext(ctx, `
  UPDATE login_sessions
  SET active = false, invalidated_at = $2, reason = 'replaced'
  WHERE user_id = $1 AND active = true
 `, user.ID, row.CreatedAt); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)`, row.ID, row.UserID, row.TokenHash, row.CreatedAt); err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		s.mu.Lock()
		for hash, old := range u.sessions {
			old.Active = false
			delete(s.sessions, hash)
		}
		u.sessions = map[string]*LoginSessionRow{row.TokenHash: row}
		s.sessions[row.TokenHash] = row
		s.mu.Unlock()
		return c.JSON(http.StatusOK, loginResponse{Token: token, User: newSessionUserResponse(user)})
	}
}

func logoutHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		session := GetSession(c)
		if session == nil {
			return NewAPIError(http.StatusUnauthorized, "invalid_session")
		}
		s := h.Store
		unlock := lockUsers(session.UserID)
		defer unlock()
		if _, err := h.DB.ExecContext(c.Request().Context(), `
		UPDATE login_sessions
		SET active = false, invalidated_at = $3, reason = 'logout'
		WHERE user_id = $1 AND active = true AND token_hash = $2
	`, session.UserID, session.TokenHash, now()); err != nil {
			return err
		}
		s.mu.Lock()
		if row := s.sessions[session.TokenHash]; row != nil && row.UserID == session.UserID {
			row.Active = false
			delete(s.sessions, session.TokenHash)
			if u := s.users[session.UserID]; u != nil {
				delete(u.sessions, session.TokenHash)
			}
		}
		s.mu.Unlock()
		return c.NoContent(http.StatusNoContent)
	}
}

func getMeHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		session := GetSession(c)
		if session == nil {
			return NewAPIError(http.StatusUnauthorized, "invalid_session")
		}
		s := h.Store
		s.mu.RLock()
		u := s.activeUser(session.UserID)
		var user User
		if u != nil {
			user = u.User
		}
		s.mu.RUnlock()
		if u == nil {
			return NewAPIError(http.StatusNotFound, "account_not_found")
		}
		return c.JSON(http.StatusOK, newGetMeResponse(user))
	}
}

func updateMeHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		session := GetSession(c)
		if session == nil {
			return NewAPIError(http.StatusUnauthorized, "invalid_session")
		}
		var input updateMeRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		if (input.DisplayName == nil && input.Bio == nil && input.Region == nil) ||
			(input.DisplayName != nil && !validUserText(*input.DisplayName, 1, 100)) ||
			(input.Bio != nil && !validUserText(*input.Bio, 0, 500)) ||
			(input.Region != nil && !validRegion(*input.Region)) {
			return NewAPIError(http.StatusBadRequest, "invalid_profile")
		}
		s := h.Store
		unlock := lockUsers(session.UserID)
		defer unlock()
		if _, err := h.DB.ExecContext(c.Request().Context(), `
  UPDATE users SET display_name = COALESCE($2, display_name), bio = COALESCE($3, bio), region = COALESCE($4, region), updated_at = $5
  WHERE id = $1`, session.UserID, input.DisplayName, input.Bio, input.Region, now()); err != nil {
			return err
		}
		s.mu.Lock()
		u := s.users[session.UserID]
		if u != nil {
			if input.DisplayName != nil {
				u.DisplayName = *input.DisplayName
				u.normDisplayName = normalizeSearchText(u.DisplayName)
			}
			if input.Bio != nil {
				u.Bio = *input.Bio
			}
			if input.Region != nil {
				u.Region = *input.Region
			}
		}
		var user User
		ok := u != nil && u.active()
		if ok {
			user = u.User
		}
		s.mu.Unlock()
		if !ok {
			return NewAPIError(http.StatusNotFound, "account_not_found")
		}
		return c.JSON(http.StatusOK, updateMeResponse(newGetMeResponse(user)))
	}
}
