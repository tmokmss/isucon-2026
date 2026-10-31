package isureal

// 画像の本体はファイルに置き、PostgreSQL の image_objects.data には新しい画像では空のバイト列を入れる（行は判定と初期化のために残す）。
// - asia1: 書き込みは new/ に一時ファイル → rename をしてから DB を commit する（ディスクへの同期は裏で 1 秒ごと）。
//   読み出しは init/ → new/ の順にファイルを探し、無ければ PostgreSQL の data（初期データの画像だけが中身を持つ）を読んで init/ に置く
// - europe1・america1: 手元に無い画像は asia1 の :8081 /_internal/image から取る（PostgreSQL の data は読まない）
// - initialize で asia1 の new/ も消す

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
)

var originImageDB atomic.Pointer[sqlx.DB]

var newImagesDirty atomic.Bool

// 書き込みの経路では fsync しない。fsync を投稿ごとに待つと、最終日の一斉の投稿で asia1 のディスクが詰まり、
// nginx のログの書き込みまで止まる。代わりに裏で 1 秒ごとに sync する（PostgreSQL の synchronous_commit=off と同じ程度の幅）
func init() {
	go func() {
		for range time.Tick(time.Second) {
			if newImagesDirty.Swap(false) {
				syscall.Sync()
			}
		}
	}()
}

// storeNewImage は new/ へ一時ファイルを書いてから rename する。DB の commit より前に呼ぶ
func storeNewImage(id string, data []byte) error {
	sub := filepath.Join(imageRootDir, "new", id[:2])
	if err := os.MkdirAll(sub, 0o755); err != nil {
		return err
	}
	f, err := os.CreateTemp(sub, "."+id+".tmp*")
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	if err == nil {
		err = os.Rename(f.Name(), filepath.Join(sub, id))
	}
	if err != nil {
		_ = os.Remove(f.Name())
		return err
	}
	newImagesDirty.Store(true)
	return nil
}

// loadImage は asia1 で画像の本体を返す。initial は初期データの画像か（init/ か PostgreSQL から読んだか）
func loadImage(ctx context.Context, db *sqlx.DB, id string, createdAt *time.Time) (data []byte, initial bool, err error) {
	if !localImageIDPattern.MatchString(id) || len(id) < 2 {
		return nil, false, sql.ErrNoRows
	}
	if data, err := os.ReadFile(filepath.Join(imageRootDir, "init", id[:2], id)); err == nil {
		return data, true, nil
	}
	if data, err := os.ReadFile(filepath.Join(imageRootDir, "new", id[:2], id)); err == nil {
		return data, false, nil
	}
	err = sql.ErrNoRows
	if createdAt != nil {
		err = db.GetContext(ctx, &data, `SELECT data FROM image_objects WHERE id = $1 AND created_at = $2`, id, *createdAt)
	}
	if errors.Is(err, sql.ErrNoRows) {
		err = db.GetContext(ctx, &data, `SELECT data FROM image_objects WHERE id = $1`, id)
	}
	if err != nil {
		return nil, false, err
	}
	if len(data) == 0 {
		return nil, false, fmt.Errorf("image %s: no file and no data in PostgreSQL", id)
	}
	writeLocalImage(id, data, true)
	return data, true, nil
}

// serveOriginImage は europe1・america1 からの本体の取得に答える（asia1 の :8081）。取得してよいかの判定は呼び出し側が済ませている
func serveOriginImage(w http.ResponseWriter, r *http.Request) {
	if _, ok := privateRemote(r); !ok {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	db := originImageDB.Load()
	if db == nil {
		http.Error(w, "not ready", http.StatusServiceUnavailable)
		return
	}
	data, initial, err := loadImage(r.Context(), db, r.URL.Query().Get("id"), nil)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if initial {
		w.Header().Set("X-Image-Initial", "1")
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	_, _ = w.Write(data)
}

var originImageClient = &http.Client{Timeout: 10 * time.Second}

// fetchOriginImage は europe1・america1 で、手元に無い画像の本体を asia1 から取る。asia1 に無ければ sql.ErrNoRows
func fetchOriginImage(ctx context.Context, origin, id string) (data []byte, initial bool, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://"+origin+"/_internal/image?id="+id, nil)
	if err != nil {
		return nil, false, err
	}
	resp, err := originImageClient.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, false, sql.ErrNoRows
	}
	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("origin image %s: status %d", id, resp.StatusCode)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}
	return data, resp.Header.Get("X-Image-Initial") == "1", nil
}
