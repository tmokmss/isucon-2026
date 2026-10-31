package isureal

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

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
		var holders []string
		var initials []bool
		for _, p := range u.posts {
			if p.DeletedAt == nil {
				posts = append(posts, p.post())
				holders = append(holders, p.Holder)
				// asia1 の手元のファイルには初期データの画像が無い
				initials = append(initials, selfInternal == "" && !p.CreatedAt.After(s.boundary))
			}
		}
		s.mu.RUnlock()

		images := make(map[string][]byte, len(posts))
		var missingIDs []string
		var missingAt []time.Time
		for i, p := range posts {
			if data, _, ok := imageCache.get(p.ImageID); ok {
				images[p.ImageID] = data
			} else if initials[i] {
				missingIDs = append(missingIDs, p.ImageID)
				missingAt = append(missingAt, p.PostedAt)
			} else if data, ok := readHeldLocal(p.ImageID); ok {
				images[p.ImageID] = data
			} else if data, ok, _ := fetchHeldImage(ctx, p.ImageID, holders[i]); ok {
				images[p.ImageID] = data
			} else {
				missingIDs = append(missingIDs, p.ImageID)
				missingAt = append(missingAt, p.PostedAt)
			}
		}
		if len(missingIDs) > 0 {
			// 初期データの画像は asia2 のディスクから 1 行ずつ読まれ、1 本の問い合わせでは読み出しが直列になる（キューの深さが 1）。分けて並列に投げる
			const chunk = 16
			var wg sync.WaitGroup
			var mu sync.Mutex
			var firstErr error
			for lo := 0; lo < len(missingIDs); lo += chunk {
				hi := min(lo+chunk, len(missingIDs))
				wg.Add(1)
				go func(ids []string, at []time.Time) {
					defer wg.Done()
					got, err := queryInitialImages(ctx, h, ids, at)
					mu.Lock()
					defer mu.Unlock()
					if err != nil {
						if firstErr == nil {
							firstErr = err
						}
						return
					}
					for id, data := range got {
						images[id] = data
					}
				}(missingIDs[lo:hi], missingAt[lo:hi])
			}
			wg.Wait()
			if firstErr != nil {
				return firstErr
			}
			var stillMissing []string
			for _, id := range missingIDs {
				if _, ok := images[id]; !ok {
					stillMissing = append(stillMissing, id)
				}
			}
			if len(stillMissing) > 0 {
				rows, err := h.DB.QueryContext(ctx, `SELECT id, data FROM image_objects WHERE id = ANY($1)`, stillMissing)
				if err != nil {
					return err
				}
				for rows.Next() {
					var id string
					var data []byte
					if err := rows.Scan(&id, &data); err != nil {
						rows.Close()
						return err
					}
					images[id] = data
				}
				rows.Close()
				if err := rows.Err(); err != nil {
					return err
				}
			}
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

func queryInitialImages(ctx context.Context, h *HandlerContext, ids []string, at []time.Time) (map[string][]byte, error) {
	rows, err := h.DB.QueryContext(ctx, `
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at)
CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i`, ids, at)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	got := make(map[string][]byte, len(ids))
	for rows.Next() {
		var id string
		var data []byte
		if err := rows.Scan(&id, &data); err != nil {
			return nil, err
		}
		got[id] = data
	}
	return got, rows.Err()
}
