package isureal

// asia1 の :8081 で、ストアの中身を europe1・america1 へ渡す（PostgreSQL の表はもう初期データのままなので）。
// - /_internal/snapshot: 複製が起動時に PostgreSQL の初期データへ重ねるスナップショット（persist.go と同じ形）
// - /_internal/image-check: 画像を取得してよいかの判定（image_api.go の imageOwnerLocked と同じ条件）。
//   asia1 が読み込み前・落ちているときは 503 やつながらないエラーになり、呼び出し側は許可しない

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var originStore atomic.Pointer[Store]

func serveOriginSnapshot(w http.ResponseWriter, r *http.Request) {
	if _, ok := privateRemote(r); !ok {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	s := originStore.Load()
	if s == nil {
		http.Error(w, "not ready", http.StatusServiceUnavailable)
		return
	}
	s.mu.RLock()
	d := s.snapshotLocked()
	s.mu.RUnlock()
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	_, _ = w.Write(buf.Bytes())
}

var originSnapshotClient = &http.Client{Timeout: 60 * time.Second}

// waitOriginSnapshot は asia1 が読み込みを終えて返せるようになるまで待つ（複製の台が先に起動したとき）
func waitOriginSnapshot(origin string) (*snapshotData, error) {
	for {
		d, err := fetchOriginSnapshot(origin)
		if err == nil {
			return d, nil
		}
		log.Printf("repl: waiting for the origin snapshot: %v", err)
		time.Sleep(time.Second)
	}
}

func fetchOriginSnapshot(origin string) (*snapshotData, error) {
	resp, err := originSnapshotClient.Get("http://" + origin + "/_internal/snapshot")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("origin snapshot: status %d", resp.StatusCode)
	}
	d := &snapshotData{}
	if err := gob.NewDecoder(resp.Body).Decode(d); err != nil {
		return nil, err
	}
	return d, nil
}

type imageCheckResponse struct {
	SecretKeyHash string `json:"s"`
	OwnerID       string `json:"o"`
	IsProfile     bool   `json:"p"`
	HasPost       bool   `json:"h"`
	RealMoji      string `json:"r"` // 投稿者 ID と投稿の画像 ID を空白でつないだもの
}

func serveOriginImageCheck(w http.ResponseWriter, r *http.Request) {
	if _, ok := privateRemote(r); !ok {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	s := originStore.Load()
	if s == nil {
		http.Error(w, "not ready", http.StatusServiceUnavailable)
		return
	}
	res, ok := s.imageCheck(r.URL.Query().Get("id"))
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&res)
}

// imageCheck は画像を取得してよいかの判定を返す。持ち主がいなければ ok=false
func (s *Store) imageCheck(id string) (res imageCheckResponse, ok bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	owner, available, _ := s.imageOwnerLocked(id)
	if owner == nil {
		return res, false
	}
	res.SecretKeyHash, res.OwnerID = owner.SecretKeyHash, owner.ID
	if available {
		switch {
		case s.profileByImage[id] != nil:
			res.IsProfile = true
		case s.postByImage[id] != nil:
			res.HasPost = true
		default:
			p := s.rmByImage[id].post
			res.RealMoji = p.AuthorID + " " + p.ImageID
		}
	}
	return res, true
}

func (res *imageCheckResponse) remote() *remoteImageCheck {
	return &remoteImageCheck{
		SecretKeyHash: res.SecretKeyHash, OwnerID: res.OwnerID, MediaType: "image/jpeg",
		IsProfile: res.IsProfile, HasPost: res.HasPost,
		RealMoji: sql.NullString{String: res.RealMoji, Valid: res.RealMoji != ""},
	}
}

var originImageCheckClient = &http.Client{Timeout: 5 * time.Second}

// fetchImageCheck は europe1・america1 で asia1 に判定を問い合わせる。持ち主がいなければ sql.ErrNoRows
func fetchImageCheck(r *http.Request, origin, id string) (*remoteImageCheck, error) {
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, "http://"+origin+"/_internal/image-check?id="+id, nil)
	if err != nil {
		return nil, err
	}
	resp, err := originImageCheckClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, sql.ErrNoRows
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("origin image check %s: status %d", id, resp.StatusCode)
	}
	var res imageCheckResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res.remote(), nil
}
