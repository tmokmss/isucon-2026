package isureal

import (
	"archive/zip"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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
				// asia1 の初期データの画像は readHeldLocal（new/ だけ）では見つからない
				initials = append(initials, selfInternal == "" && !p.CreatedAt.After(s.boundary))
			}
		}
		s.mu.RUnlock()

		images := make(map[string][]byte, len(posts))
		var missingIDs []string
		var missingAt []time.Time
		var fileIDs []string
		var fileAt []time.Time
		for i, p := range posts {
			if data, _, ok := imageCache.get(p.ImageID); ok {
				images[p.ImageID] = data
			} else if initials[i] {
				// 初期データの画像は半分を asia1 の init/ から、残りを PostgreSQL から並行して読む。asia2 のディスクの読み出しが頭打ちになるため
				if (len(fileIDs)+len(missingIDs))%2 == 0 {
					fileIDs = append(fileIDs, p.ImageID)
					fileAt = append(fileAt, p.PostedAt)
				} else {
					missingIDs = append(missingIDs, p.ImageID)
					missingAt = append(missingAt, p.PostedAt)
				}
			} else if data, ok := readHeldLocal(p.ImageID); ok {
				images[p.ImageID] = data
			} else if data, ok, _ := fetchHeldImage(ctx, p.ImageID, holders[i]); ok {
				images[p.ImageID] = data
			} else {
				missingIDs = append(missingIDs, p.ImageID)
				missingAt = append(missingAt, p.PostedAt)
			}
		}
		fromFiles := readInitImages(fileIDs)
		dbImages, err := h.loadImagesFromDB(ctx, missingIDs, missingAt)
		files := <-fromFiles
		if err != nil {
			return err
		}
		for id, data := range dbImages {
			images[id] = data
		}
		var fileMissIDs []string
		var fileMissAt []time.Time
		for i, id := range fileIDs {
			if data, ok := files[id]; ok {
				images[id] = data
			} else {
				fileMissIDs = append(fileMissIDs, id)
				fileMissAt = append(fileMissAt, fileAt[i])
			}
		}
		if len(fileMissIDs) > 0 {
			dbImages, err := h.loadImagesFromDB(ctx, fileMissIDs, fileMissAt)
			if err != nil {
				return err
			}
			for id, data := range dbImages {
				images[id] = data
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

// readInitImages は init/ の画像を 4 並列で読み、読めた分を返す
func readInitImages(ids []string) <-chan map[string][]byte {
	out := make(chan map[string][]byte, 1)
	if len(ids) == 0 {
		out <- nil
		return out
	}
	go func() {
		datas := make([][]byte, len(ids))
		var wg sync.WaitGroup
		next := make(chan int)
		for range 4 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := range next {
					if data, err := os.ReadFile(filepath.Join(imageRootDir, "init", ids[i][:2], ids[i])); err == nil {
						datas[i] = data
					}
				}
			}()
		}
		for i, id := range ids {
			if localImageIDPattern.MatchString(id) && len(id) >= 2 {
				next <- i
			}
		}
		close(next)
		wg.Wait()
		m := make(map[string][]byte, len(ids))
		for i, d := range datas {
			if d != nil {
				m[ids[i]] = d
			}
		}
		out <- m
	}()
	return out
}

// loadImagesFromDB は PostgreSQL の image_objects.data を (id, created_at) で引き、無かった分を id だけで引き直す
func (h *HandlerContext) loadImagesFromDB(ctx context.Context, ids []string, at []time.Time) (map[string][]byte, error) {
	images := make(map[string][]byte, len(ids))
	if len(ids) == 0 {
		return images, nil
	}
	scan := func(rows *sql.Rows, err error) error {
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var id string
			var data []byte
			if err := rows.Scan(&id, &data); err != nil {
				return err
			}
			images[id] = data
		}
		return rows.Err()
	}
	if err := scan(h.DB.QueryContext(ctx, `
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at)
CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i`, ids, at)); err != nil {
		return nil, err
	}
	var stillMissing []string
	for _, id := range ids {
		if _, ok := images[id]; !ok {
			stillMissing = append(stillMissing, id)
		}
	}
	if len(stillMissing) > 0 {
		if err := scan(h.DB.QueryContext(ctx, `SELECT id, data FROM image_objects WHERE id = ANY($1)`, stillMissing)); err != nil {
			return nil, err
		}
	}
	return images, nil
}
