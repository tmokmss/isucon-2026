package isureal

// europe1・america1 が受けた投稿の画像は、その台の new/ に置き、asia1 へは画像の ID と持ち主の台（:8081）だけを送る。
// - post_forward.go: 画像を検証して new/ に書いてから、画像のパートを空にした本文と X-Isureal-Image（ID・持ち主・幅・高さ・HMAC）を asia1 へ送る
// - asia1 の createPostHandler: 印が合えば、その ID と持ち主で投稿を作る（PostRow.Holder）。画像の本体は受けない
// - 本体を読むときは、手元に無ければ PostRow.Holder の台の /_internal/image から取り、手元の new/ に置く。Holder が空なら asia1 が持つ
// - 取得してよいかの判定はどの台もストア（asia1 は本体、europe1・america1 は複製）で行う。initialize では各台が new/ を消す

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

const heldImageHeader = "X-Isureal-Image"

// selfInternal はこの台の :8081（europe1・america1 の env.sh の ISUREAL_SELF_INTERNAL）。asia1 は空
var selfInternal = os.Getenv("ISUREAL_SELF_INTERNAL")

// imageHolderOf は画像の本体を持つ台を返す。ストアを持つ台が差し替える。空なら asia1
var imageHolderOf = func(id string) string { return "" }

// asia1 の :8081。europe1・america1 で、持ち主が asia1 の画像を取りに行く先
var asiaInternal = func() string {
	if peers := envPeers("ISUREAL_IMAGE_EVENT_PEERS"); len(peers) > 0 {
		return peers[0]
	}
	return ""
}()

func heldImageMAC(id, holder string, width, height int) []byte {
	m := hmac.New(sha256.New, jpegVerifyKey)
	m.Write([]byte("held\x00" + id + "\x00" + holder + "\x00" + strconv.Itoa(width) + "\x00" + strconv.Itoa(height)))
	return m.Sum(nil)
}

func heldImageHeaderValue(id, holder string, width, height int) string {
	return id + " " + holder + " " + strconv.Itoa(width) + " " + strconv.Itoa(height) + " " + hex.EncodeToString(heldImageMAC(id, holder, width, height))
}

// parseHeldImageHeader は asia1 で印を確かめる。合わなければ ok=false
func parseHeldImageHeader(req *http.Request) (id, holder string, width, height int, ok bool) {
	v := req.Header.Get(heldImageHeader)
	if v == "" || len(jpegVerifyKey) == 0 {
		return "", "", 0, 0, false
	}
	parts := strings.Split(v, " ")
	if len(parts) != 5 || !localImageIDPattern.MatchString(parts[0]) || len(parts[0]) < 2 || parts[1] == "" {
		return "", "", 0, 0, false
	}
	w, errW := strconv.Atoi(parts[2])
	h, errH := strconv.Atoi(parts[3])
	mac, errM := hex.DecodeString(parts[4])
	if errW != nil || errH != nil || errM != nil || !hmac.Equal(mac, heldImageMAC(parts[0], parts[1], w, h)) {
		return "", "", 0, 0, false
	}
	return parts[0], parts[1], w, h, true
}

// readHeldLocal はこの台の手元のファイルを読む。europe1・america1 は init/ と new/、asia1 は new/ だけ（初期データは PostgreSQL）
var readHeldLocal = readNewImage

// fetchHeldImage は、手元に無い画像の本体を持ち主の台（imageHolderOf の値）から取り、手元の new/ に置く。持ち主が無い・この台なら false
// 同じ画像を同時に取りに行く要求は、先に出た 1 本の結果を待つ（リージョン間の帯域を同じ本体で重ねて使わない）
type heldFetch struct {
	done chan struct{}
	data []byte
	ok   bool
	err  error
}

var (
	heldFetchMu sync.Mutex
	heldFetches = map[string]*heldFetch{}
)

func fetchHeldImage(ctx context.Context, id, holder string) ([]byte, bool, error) {
	heldFetchMu.Lock()
	if f := heldFetches[id]; f != nil {
		heldFetchMu.Unlock()
		select {
		case <-f.done:
			return f.data, f.ok, f.err
		case <-ctx.Done():
			return nil, false, ctx.Err()
		}
	}
	f := &heldFetch{done: make(chan struct{})}
	heldFetches[id] = f
	heldFetchMu.Unlock()
	f.data, f.ok, f.err = fetchHeldImageOnce(context.WithoutCancel(ctx), id, holder)
	heldFetchMu.Lock()
	delete(heldFetches, id)
	heldFetchMu.Unlock()
	close(f.done)
	return f.data, f.ok, f.err
}

func fetchHeldImageOnce(ctx context.Context, id, holder string) ([]byte, bool, error) {
	if holder == "" {
		holder = asiaInternal
	}
	if holder == "" || holder == selfInternal {
		return nil, false, nil
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	data, initial, err := fetchOriginImage(ctx, holder, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	writeLocalImage(id, data, initial)
	return data, true, nil
}

// serveHeldImage は europe1・america1 の :8081 で、手元のファイルにある本体だけを返す（ここから先は取りに行かない）
func serveHeldImage(w http.ResponseWriter, r *http.Request) {
	if _, ok := privateRemote(r); !ok {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	id := r.URL.Query().Get("id")
	if !localImageIDPattern.MatchString(id) || len(id) < 2 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	data, ok := readLocalImage(id)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	_, _ = w.Write(data)
}

// holdPostImage は europe1・america1 で、検証済みの投稿の画像を new/ に置き、画像のパートを空にした本文を作る。
// readPostImage と同じく最初の image のファイルのパートを画像とみなす。置けなければ ok=false（今まで通り本文ごと送る）
func holdPostImage(header http.Header, body []byte, width, height int) (held, contentType string, newBody []byte, ok bool) {
	if selfInternal == "" || len(jpegVerifyKey) == 0 {
		return "", "", nil, false
	}
	mr, err := (&http.Request{Header: header, Body: io.NopCloser(bytes.NewReader(body))}).MultipartReader()
	if err != nil {
		return "", "", nil, false
	}
	var out bytes.Buffer
	mw := multipart.NewWriter(&out)
	id := uuid.NewString()
	stored := false
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", "", nil, false
		}
		pw, err := mw.CreatePart(p.Header)
		if err != nil {
			return "", "", nil, false
		}
		if !stored && p.FormName() == "image" && p.FileName() != "" {
			image, err := io.ReadAll(p)
			if err != nil || storeNewImage(id, image) != nil {
				return "", "", nil, false
			}
			stored = true
			continue
		}
		if _, err := io.Copy(pw, p); err != nil {
			return "", "", nil, false
		}
	}
	if !stored || mw.Close() != nil {
		return "", "", nil, false
	}
	return heldImageHeaderValue(id, selfInternal, width, height), mw.FormDataContentType(), out.Bytes(), true
}

// storeImageHolderLookup は、ストアの投稿から画像の持ち主を引くようにする
func storeImageHolderLookup(s *Store) {
	imageHolderOf = func(id string) string {
		s.mu.RLock()
		defer s.mu.RUnlock()
		if p := s.postByImage[id]; p != nil {
			return p.Holder
		}
		return ""
	}
}
