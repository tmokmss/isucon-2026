package isureal

import (
	"archive/zip"
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

type exportPostResponse struct {
	createPostResponse
	ImagePath string `json:"imagePath"`
}

func exportPostsHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := c.Request().Context()
		userID := GetSession(c).UserID
		s := h.Store
		s.mu.RLock()
		u := s.activeUser(userID)
		if u == nil {
			s.mu.RUnlock()
			return NewAPIError(http.StatusInternalServerError, "user_not_found")
		}
		author := u.User
		var posts []Post
		for _, p := range u.posts {
			if p.DeletedAt == nil {
				posts = append(posts, p.post())
			}
		}
		s.mu.RUnlock()

		images := make(map[string][]byte, len(posts))
		for _, p := range posts {
			if data, _, ok := imageCache.get(p.ImageID); ok {
				images[p.ImageID] = data
				continue
			}
			postedAt := p.PostedAt
			data, _, err := loadImage(ctx, h.DB, p.ImageID, &postedAt)
			if errors.Is(err, sql.ErrNoRows) {
				continue
			}
			if err != nil {
				return err
			}
			images[p.ImageID] = data
		}

		targets := make([]signTarget, 0, len(posts)+1)
		for _, post := range posts {
			targets = append(targets, signTarget{author.ImageKey, post.ImageID})
		}
		if author.ProfileImageID != nil {
			targets = append(targets, signTarget{author.ImageKey, *author.ProfileImageID})
		}
		prefetchImageSignatures(targets)

		var body bytes.Buffer
		writer := zip.NewWriter(&body)
		defer writer.Close()
		items := make([]exportPostResponse, 0, len(posts))
		for _, post := range posts {
			data, ok := images[post.ImageID]
			if !ok {
				continue
			}
			imagePath := fmt.Sprintf("images/%06d.jpg", len(items)+1)
			image, err := writer.CreateHeader(&zip.FileHeader{Name: imagePath, Method: zip.Store})
			if err != nil {
				return err
			}
			if _, err := image.Write(data); err != nil {
				return err
			}
			items = append(items, exportPostResponse{createPostResponse: newPostResponse(post, author), ImagePath: imagePath})
		}
		manifest, err := writer.Create("posts.json")
		if err != nil {
			return err
		}
		if err := json.NewEncoder(manifest).Encode(items); err != nil {
			return err
		}
		if err := writer.Close(); err != nil {
			return err
		}
		c.Response().Header().Set(echo.HeaderContentDisposition, `attachment; filename="isureal-posts.zip"`)
		c.Response().Header().Set(echo.HeaderCacheControl, "no-store")
		return c.Blob(http.StatusOK, "application/zip", body.Bytes())
	}
}
