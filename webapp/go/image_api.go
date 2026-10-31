package isureal

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/labstack/echo/v5"
)

type getImageRequest struct {
	ImageID string `param:"imageId" json:"-"`
	Exp     string `query:"exp" json:"-"`
	Sig     string `query:"sig" json:"-"`
}

// imageOwnerLocked は画像の持ち主（有効な利用者）と、画像を今も取得できるか、画像の作成時刻を返す
func (s *Store) imageOwnerLocked(imageID string) (owner *UserRow, available bool, createdAt *time.Time) {
	if u := s.profileByImage[imageID]; u != nil {
		if !u.active() {
			return nil, false, nil
		}
		return u, true, u.ProfileImageCreatedAt
	}
	if p := s.postByImage[imageID]; p != nil {
		u := s.activeUser(p.AuthorID)
		if u == nil {
			return nil, false, nil
		}
		t := p.PostedAt
		return u, p.DeletedAt == nil, &t
	}
	if r := s.rmByImage[imageID]; r != nil {
		u := s.activeUser(r.ReactorID)
		if u == nil {
			return nil, false, nil
		}
		t := r.CreatedAt
		ok := r.current() && r.post.DeletedAt == nil && s.activeUser(r.post.AuthorID) != nil
		return u, ok, &t
	}
	return nil, false, nil
}

func getImageHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		c.Response().Header().Set("Cache-Control", "no-store")
		var input getImageRequest
		if err := c.Bind(&input); err != nil {
			return NewAPIError(http.StatusForbidden, "image_forbidden")
		}
		query, err := url.ParseQuery(c.Request().URL.RawQuery)
		if err != nil || len(query["exp"]) != 1 || len(query["sig"]) != 1 {
			return NewAPIError(http.StatusForbidden, "image_forbidden")
		}
		exp, err := strconv.ParseInt(input.Exp, 10, 64)
		if err != nil || strconv.FormatInt(exp, 10) != input.Exp || exp <= time.Now().Unix() {
			return NewAPIError(http.StatusForbidden, "image_forbidden")
		}
		sig, err := base64.RawURLEncoding.Strict().DecodeString(input.Sig)
		if err != nil || len(sig) != sha256.Size || base64.RawURLEncoding.EncodeToString(sig) != input.Sig {
			return NewAPIError(http.StatusForbidden, "image_forbidden")
		}

		s := h.Store
		s.mu.RLock()
		owner, available, createdAt := s.imageOwnerLocked(input.ImageID)
		var key []byte
		var holder string
		initial := false
		if owner != nil {
			key = owner.ImageKey
			if p := s.postByImage[input.ImageID]; p != nil {
				holder = p.Holder
				initial = !p.CreatedAt.After(s.boundary)
			} else {
				initial = createdAt != nil && !createdAt.After(s.boundary)
			}
		}
		s.mu.RUnlock()
		if owner == nil {
			return NewAPIError(http.StatusForbidden, "image_forbidden")
		}
		if !hmac.Equal(sig, imageSignature(key, input.ImageID, exp)) {
			c.Set("abuse", "invalid_signature")
			return NewAPIError(http.StatusForbidden, "image_forbidden")
		}
		if !available {
			return NewAPIError(http.StatusForbidden, "image_forbidden")
		}
		if imageNotModified(c, input.ImageID) {
			return c.NoContent(http.StatusNotModified)
		}
		data, mediaType, err := h.fetchImage(c, input.ImageID, createdAt, holder, initial)
		if errors.Is(err, sql.ErrNoRows) {
			return NewAPIError(http.StatusForbidden, "image_forbidden")
		}
		if err != nil {
			return err
		}
		if exp <= time.Now().Unix() {
			return NewAPIError(http.StatusForbidden, "image_forbidden")
		}
		return c.Blob(http.StatusOK, mediaType, data)
	}
}

// imageNotModified は Cache-Control と ETag を付け、条件付きの取得で手元の写しが使えるかを返す。
// 画像の ID ごとに中身は変わらない。no-cache なので取得のたびに署名・期限・削除の判定を通してから 304 を返す
func imageNotModified(c *echo.Context, imageID string) bool {
	etag := `"` + imageID + `"`
	h := c.Response().Header()
	h.Set("Cache-Control", "private, no-cache")
	h.Set("ETag", etag)
	return c.Request().Header.Get("If-None-Match") == etag
}

func (h *HandlerContext) fetchImage(c *echo.Context, imageID string, createdAt *time.Time, holder string, initial bool) ([]byte, string, error) {
	if data, mediaType, ok := imageCache.get(imageID); ok {
		return data, mediaType, nil
	}
	data, _, err := loadImage(c.Request().Context(), h.DB, imageID, createdAt, holder, initial)
	if err != nil {
		return nil, "", err
	}
	imageCache.put(imageID, data, "image/jpeg")
	return data, "image/jpeg", nil
}

// 画像が取得できなくなったときに呼ぶ。他の台へ無効化を知らせる処理はここに差し込む
var imagesInvalidated = func(imageIDs ...string) {}

// 退会で、その利用者のすべての画像が取得できなくなったときに呼ぶ
var userImagesInvalidated = func(userID string) {}
