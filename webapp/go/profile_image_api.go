package isureal

import (
	"bytes"
	"image/jpeg"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func putProfileImageHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		if c.Request().Header.Get("Content-Type") != "image/jpeg" {
			return NewAPIError(http.StatusUnsupportedMediaType, "unsupported_media_type")
		}
		var body bytes.Buffer
		body.Grow(max(int(c.Request().ContentLength), 0) + bytes.MinRead)
		if _, err := body.ReadFrom(c.Request().Body); err != nil {
			return err
		}
		data := body.Bytes()
		cfg, err := jpeg.DecodeConfig(bytes.NewReader(data))
		if err == nil {
			_, _, err = checkJPEG(data)
		}
		if err != nil {
			c.Set("abuse", "invalid_image")
			return NewAPIError(http.StatusBadRequest, "invalid_profile_image")
		}
		if cfg.Width > 512 || cfg.Height > 512 {
			return NewAPIError(http.StatusBadRequest, "invalid_profile_image")
		}

		userID := GetSession(c).UserID
		s := h.Store
		unlock := lockUsers(userID)
		defer unlock()
		s.mu.RLock()
		u := s.users[userID]
		var oldID *string
		if u != nil {
			oldID = u.ProfileImageID
		}
		s.mu.RUnlock()
		if u == nil {
			return NewAPIError(http.StatusNotFound, "account_not_found")
		}
		imageID := uuid.NewString()
		createdAt := now()
		if err := storeNewImage(imageID, data); err != nil {
			return err
		}
		s.mu.Lock()
		rw := s.commitLocked(&opProfileImage{UserID: u.ID, ImageID: &imageID, CreatedAt: &createdAt})
		s.mu.Unlock()
		rw.wait()
		if oldID != nil {
			imagesInvalidated(*oldID)
		}
		imageCache.put(imageID, data, "image/jpeg")
		return c.NoContent(http.StatusNoContent)
	}
}

func deleteProfileImageHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		userID := GetSession(c).UserID
		s := h.Store
		unlock := lockUsers(userID)
		defer unlock()
		s.mu.RLock()
		u := s.users[userID]
		var imageID *string
		if u != nil {
			imageID = u.ProfileImageID
		}
		s.mu.RUnlock()
		if u == nil || imageID == nil {
			return NewAPIError(http.StatusNotFound, "account_not_found")
		}
		s.mu.Lock()
		rw := s.commitLocked(&opProfileImage{UserID: u.ID})
		s.mu.Unlock()
		rw.wait()
		imagesInvalidated(*imageID)
		return c.NoContent(http.StatusNoContent)
	}
}
