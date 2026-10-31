package isureal

import (
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func putProfileImageHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		if c.Request().Header.Get("Content-Type") != "image/jpeg" {
			return NewAPIError(http.StatusUnsupportedMediaType, "unsupported_media_type")
		}
		data, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}
		width, height, ok := checkJPEG(data)
		if !ok {
			c.Set("abuse", "invalid_image")
			return NewAPIError(http.StatusBadRequest, "invalid_profile_image")
		}
		if width > 512 || height > 512 {
			return NewAPIError(http.StatusBadRequest, "invalid_profile_image")
		}

		ctx := c.Request().Context()
		userID := GetSession(c).UserID
		s := h.Store
		unlock := lockUsers(userID)
		defer unlock()
		s.mu.RLock()
		u := s.users[userID]
		var oldID *string
		var oldCreatedAt any
		if u != nil {
			oldID = u.ProfileImageID
			if u.ProfileImageCreatedAt != nil {
				oldCreatedAt = *u.ProfileImageCreatedAt
			}
		}
		s.mu.RUnlock()
		if u == nil {
			return NewAPIError(http.StatusNotFound, "account_not_found")
		}
		imageID := uuid.NewString()
		createdAt := now()
		tx, err := h.DB.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
		defer tx.Rollback()
		if _, err := tx.ExecContext(ctx, `INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,'image/jpeg',$4)`, imageID, userID, data, createdAt); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `UPDATE users SET profile_image_id=$2,updated_at=$3 WHERE id=$1`, userID, imageID, createdAt); err != nil {
			return err
		}
		if oldID != nil {
			if oldCreatedAt != nil {
				if _, err := tx.ExecContext(ctx, `DELETE FROM image_objects WHERE id=$1 AND created_at=$2`, *oldID, oldCreatedAt); err != nil {
					return err
				}
			} else if _, err := tx.ExecContext(ctx, `DELETE FROM image_objects WHERE id=$1`, *oldID); err != nil {
				return err
			}
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		s.mu.Lock()
		if oldID != nil {
			delete(s.profileByImage, *oldID)
		}
		u.ProfileImageID = &imageID
		u.ProfileImageCreatedAt = &createdAt
		s.profileByImage[imageID] = u
		s.mu.Unlock()
		if oldID != nil {
			imagesInvalidated(*oldID)
		}
		imageCache.put(imageID, data, "image/jpeg")
		return c.NoContent(http.StatusNoContent)
	}
}

func deleteProfileImageHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := c.Request().Context()
		userID := GetSession(c).UserID
		s := h.Store
		unlock := lockUsers(userID)
		defer unlock()
		s.mu.RLock()
		u := s.users[userID]
		var imageID *string
		var createdAt any
		if u != nil {
			imageID = u.ProfileImageID
			if u.ProfileImageCreatedAt != nil {
				createdAt = *u.ProfileImageCreatedAt
			}
		}
		s.mu.RUnlock()
		if u == nil || imageID == nil {
			return NewAPIError(http.StatusNotFound, "account_not_found")
		}
		tx, err := h.DB.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
		defer tx.Rollback()
		if _, err := tx.ExecContext(ctx, `UPDATE users SET profile_image_id=NULL,updated_at=$2 WHERE id=$1`, userID, now()); err != nil {
			return err
		}
		if createdAt != nil {
			if _, err := tx.ExecContext(ctx, `DELETE FROM image_objects WHERE id=$1 AND created_at=$2`, *imageID, createdAt); err != nil {
				return err
			}
		} else if _, err := tx.ExecContext(ctx, `DELETE FROM image_objects WHERE id=$1`, *imageID); err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		s.mu.Lock()
		delete(s.profileByImage, *imageID)
		u.ProfileImageID = nil
		u.ProfileImageCreatedAt = nil
		s.mu.Unlock()
		imagesInvalidated(*imageID)
		return c.NoContent(http.StatusNoContent)
	}
}
