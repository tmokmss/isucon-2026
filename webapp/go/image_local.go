package isureal

// europe1・america1（ISUREAL_PASSIVE=1）で GET /api/images/:imageId だけを返す。
// - 画像の本体: images/init/（初期データ。setup/image-export.sh で書き出した分と、asia1 から取った初期データの画像）、
//   images/new/（asia1 から取った新しい画像。initialize で消す）の順に探し、無ければ asia1 の /_internal/image から取って置く
// - 取得してよいかの判定: PostgreSQL に 1 回問い合わせ、結果を依存キー（img:<画像 ID>、user:<利用者 ID>）付きで手元に持つ
// - 無効化: asia1 が画像を取れなくしたときに、img:/user: のキーを各台の :8081 へ POST し、各台が判定を捨てた受領を待ってから応答する。
//   受領が imageInvalidateTimeout 以内に来ない台は、その台へのストリームを切る（その台は切断かハートビートの途絶で判定を全部捨てる）。
//   同じキーはストリームにも流す
// - ストリームのハートビートが imageEventLease を超えて途絶えたら、手元の判定は使わず毎回 PostgreSQL で判定する。
//   切れたときは判定を全部捨てる（届かなかった無効化があっても、手元の判定が残らない）
// - initialize（asia1）は "*" を送り、受けた台は判定を全部捨てて images/new/ を消す

import (
	"bufio"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

const (
	imageEventAddr         = ":8081"
	imageEventHeartbeat    = 300 * time.Millisecond
	imageEventLease        = 2 * time.Second
	imageInvalidateTimeout = time.Second
	imageRootDir           = "/home/isucon/images"
)

var localImageIDPattern = regexp.MustCompile(`^[0-9A-Za-z_-]{1,64}$`)

// ---- asia1: 無効化を流す側 ----

type imageEventHub struct {
	mu   sync.Mutex
	subs map[chan string]string // 購読者の IP
}

var imageEvents = &imageEventHub{subs: map[chan string]string{}}

// dropPeer はその IP の台へのストリームを切る
func (hub *imageEventHub) dropPeer(ip string) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	for ch, subIP := range hub.subs {
		if subIP == ip {
			close(ch)
			delete(hub.subs, ch)
		}
	}
}

var imageInvalidateClient = &http.Client{Timeout: imageInvalidateTimeout}

// invalidateOnPeers は各台へキーを送り、受領を待つ。受領が来なかった台へのストリームは切る
func invalidateOnPeers(peers []string, lines []string) {
	imageEvents.publish(lines...)
	body := strings.Join(lines, "\n")
	var wg sync.WaitGroup
	for _, peer := range peers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := imageInvalidateClient.Post("http://"+peer+"/_internal/invalidate", "text/plain", strings.NewReader(body))
			if err == nil {
				resp.Body.Close()
				if resp.StatusCode != http.StatusNoContent {
					err = fmt.Errorf("status %d", resp.StatusCode)
				}
			}
			if err != nil {
				log.Printf("image invalidation to %s failed: %v", peer, err)
				host, _, _ := net.SplitHostPort(peer)
				imageEvents.dropPeer(host)
			}
		}()
	}
	wg.Wait()
}

func privateRemote(r *http.Request) (string, bool) {
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	ip := net.ParseIP(host)
	return host, ip != nil && ip.IsPrivate()
}

func (hub *imageEventHub) publish(lines ...string) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	for ch := range hub.subs {
		if !sendAll(ch, lines) {
			// 詰まった購読者は切る。受けた側は切断を検知して判定を全部捨てる
			close(ch)
			delete(hub.subs, ch)
		}
	}
}

func sendAll(ch chan string, lines []string) bool {
	for _, line := range lines {
		select {
		case ch <- line:
		default:
			return false
		}
	}
	return true
}

func (hub *imageEventHub) serveHTTP(w http.ResponseWriter, r *http.Request) {
	host, ok := privateRemote(r)
	if !ok {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	flusher, ok := w.(http.Flusher)
	if !ok {
		return
	}
	ch := make(chan string, 4096)
	hub.mu.Lock()
	hub.subs[ch] = host
	hub.mu.Unlock()
	defer func() {
		hub.mu.Lock()
		if _, ok := hub.subs[ch]; ok {
			delete(hub.subs, ch)
			close(ch)
		}
		hub.mu.Unlock()
	}()
	ticker := time.NewTicker(imageEventHeartbeat)
	defer ticker.Stop()
	bw := bufio.NewWriter(w)
	for {
		select {
		case line, ok := <-ch:
			if !ok {
				return
			}
			bw.WriteString(line + "\n")
			for drained := false; !drained; {
				select {
				case more, ok := <-ch:
					if !ok {
						return
					}
					bw.WriteString(more + "\n")
				default:
					drained = true
				}
			}
		case <-ticker.C:
			bw.WriteString(".\n")
		case <-r.Context().Done():
			return
		}
		if bw.Flush() != nil {
			return
		}
		flusher.Flush()
	}
}

// StartImageEventPublisher は asia1 で呼ぶ。画像が取れなくなる箇所の呼び出し口を、各台へ知らせて受領を待つ処理につなぐ。
// 送り先は ISUREAL_IMAGE_PEERS（europe1・america1 の :8081）
func StartImageEventPublisher() {
	peers := envPeers("ISUREAL_IMAGE_PEERS")
	imagesInvalidated = func(imageIDs ...string) {
		lines := make([]string, len(imageIDs))
		for i, id := range imageIDs {
			lines[i] = "img:" + id
		}
		invalidateOnPeers(peers, lines)
	}
	userImagesInvalidated = func(userID string) { invalidateOnPeers(peers, []string{"user:" + userID}) }
	imagesReset = func() { invalidateOnPeers(peers, []string{"*"}) }
	mux := http.NewServeMux()
	mux.HandleFunc("/_internal/image-events", imageEvents.serveHTTP)
	mux.HandleFunc("/_internal/image", serveOriginImage)
	serveInternal(mux)
}

func serveInternal(mux *http.ServeMux) {
	go func() {
		for {
			err := http.ListenAndServe(imageEventAddr, mux)
			log.Printf("image internal server stopped: %v", err)
			time.Sleep(time.Second)
		}
	}()
}

func envPeers(name string) []string {
	var peers []string
	for _, p := range strings.Split(os.Getenv(name), ",") {
		if p = strings.TrimSpace(p); p != "" {
			peers = append(peers, p)
		}
	}
	return peers
}

// imagesReset は initialize の後に呼ぶ
var imagesReset = func() {}

// ---- europe1・america1: 受けて画像を返す側 ----

type localImageEntry struct {
	key       []byte
	mediaType string
	deps      []string
}

type localImageState struct {
	mu       sync.Mutex
	entries  map[string]*localImageEntry
	byDep    map[string][]string
	seq      uint64
	flushSeq uint64
	recent   []seqDep // 直近の無効化。問い合わせ中に来た無効化を、結果を入れるときに当てる
	lastSeen map[string]time.Time
}

type seqDep struct {
	seq uint64
	at  time.Time
	dep string
}

func newLocalImageState(peers []string) *localImageState {
	st := &localImageState{entries: map[string]*localImageEntry{}, byDep: map[string][]string{}, lastSeen: map[string]time.Time{}}
	for _, p := range peers {
		st.lastSeen[p] = time.Time{}
	}
	return st
}

func (st *localImageState) trustedLocked() bool {
	if len(st.lastSeen) == 0 {
		return false
	}
	now := time.Now()
	for _, t := range st.lastSeen {
		if now.Sub(t) > imageEventLease {
			return false
		}
	}
	return true
}

func (st *localImageState) flushLocked() {
	st.seq++
	st.flushSeq = st.seq
	st.entries = map[string]*localImageEntry{}
	st.byDep = map[string][]string{}
	st.recent = nil
}

func (st *localImageState) invalidateLocked(dep string) {
	st.seq++
	now := time.Now()
	st.recent = append(st.recent, seqDep{seq: st.seq, at: now, dep: dep})
	if len(st.recent) > 0 && now.Sub(st.recent[0].at) > time.Minute {
		i := 0
		for i < len(st.recent) && now.Sub(st.recent[i].at) > 10*time.Second {
			i++
		}
		st.recent = append([]seqDep(nil), st.recent[i:]...)
	}
	for _, id := range st.byDep[dep] {
		delete(st.entries, id)
	}
	delete(st.byDep, dep)
}

// lookup は手元の判定を返す。trusted でなければ使わない
func (st *localImageState) lookup(id string) (entry *localImageEntry, seq uint64, trusted bool) {
	st.mu.Lock()
	defer st.mu.Unlock()
	trusted = st.trustedLocked()
	if trusted {
		entry = st.entries[id]
	}
	return entry, st.seq, trusted
}

func (st *localImageState) insert(id string, entry *localImageEntry, startSeq uint64) {
	st.mu.Lock()
	defer st.mu.Unlock()
	if !st.trustedLocked() || st.flushSeq > startSeq {
		return
	}
	for i := len(st.recent) - 1; i >= 0 && st.recent[i].seq > startSeq; i-- {
		for _, d := range entry.deps {
			if d == st.recent[i].dep {
				return
			}
		}
	}
	st.entries[id] = entry
	for _, d := range entry.deps {
		st.byDep[d] = append(st.byDep[d], id)
	}
}

func (st *localImageState) subscribe(peer string) {
	client := &http.Client{}
	for {
		err := st.readStream(client, peer)
		st.mu.Lock()
		st.lastSeen[peer] = time.Time{}
		st.flushLocked()
		st.mu.Unlock()
		log.Printf("image event stream from %s lost: %v", peer, err)
		time.Sleep(500 * time.Millisecond)
	}
}

func (st *localImageState) readStream(client *http.Client, peer string) error {
	resp, err := client.Get("http://" + peer + "/_internal/image-events")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status %d", resp.StatusCode)
	}
	// 届かなかった間の無効化は分からないので、つながり直したら全部捨ててから使い始める
	st.mu.Lock()
	st.flushLocked()
	st.mu.Unlock()
	sc := bufio.NewScanner(resp.Body)
	for sc.Scan() {
		line := sc.Text()
		st.mu.Lock()
		if line != "." {
			st.applyLocked(line)
		}
		st.lastSeen[peer] = time.Now()
		st.mu.Unlock()
	}
	if sc.Err() != nil {
		return sc.Err()
	}
	return errors.New("closed")
}

func (st *localImageState) applyLocked(line string) {
	if line == "*" {
		st.flushLocked()
		go resetNewImages()
		return
	}
	st.invalidateLocked(line)
}

// serveInvalidate は asia1 からの無効化を当ててから 204 を返す（asia1 はこの受領を待って応答する）
func (st *localImageState) serveInvalidate(w http.ResponseWriter, r *http.Request) {
	if _, ok := privateRemote(r); !ok || r.Method != http.MethodPost {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	st.mu.Lock()
	for _, line := range strings.Split(string(body), "\n") {
		if line != "" {
			st.applyLocked(line)
		}
	}
	st.mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
}

func resetNewImages() {
	dir := filepath.Join(imageRootDir, "new")
	old := dir + ".old-" + strconv.FormatInt(time.Now().UnixNano(), 10)
	if err := os.Rename(dir, old); err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Printf("failed to reset new images: %v", err)
		return
	}
	_ = os.RemoveAll(old)
}

func readLocalImage(id string) ([]byte, bool) {
	for _, sub := range []string{"init", "new"} {
		if data, err := os.ReadFile(filepath.Join(imageRootDir, sub, id[:2], id)); err == nil {
			return data, true
		}
	}
	return nil, false
}

// 初期データの画像は initialize で消えない init/ に、新しい画像は new/ に置く。一時ファイルに書いてから rename する
func writeLocalImage(id string, data []byte, initial bool) {
	sub := filepath.Join(imageRootDir, "new", id[:2])
	if initial {
		sub = filepath.Join(imageRootDir, "init", id[:2])
	}
	if err := os.MkdirAll(sub, 0o755); err != nil {
		return
	}
	f, err := os.CreateTemp(sub, "."+id+".tmp*")
	if err != nil {
		return
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
	}
}

type remoteImageCheck struct {
	SecretKeyHash string         `db:"secret_key_hash"`
	OwnerID       string         `db:"owner_id"`
	MediaType     string         `db:"media_type"`
	IsProfile     bool           `db:"is_profile"`
	HasPost       bool           `db:"has_post"`
	RealMoji      sql.NullString `db:"realmoji"`
}

// 判定の SQL は image_api.go の元の判定と同じ条件。RealMoji のときは、投稿者と投稿の画像 ID も返す（その無効化で判定を捨てるため）
const remoteImageCheckQuery = `
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type,
  COALESCE(u.profile_image_id = i.id, false) AS is_profile,
  EXISTS (SELECT 1 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post,
  (SELECT p.author_id || ' ' || p.image_id FROM realmojis r
     JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL
     JOIN users author ON author.id = p.author_id AND author.status = 'active'
     WHERE r.image_id = i.id AND r.reactor_id = u.id AND r.replaced_at IS NULL AND r.deleted_at IS NULL LIMIT 1) AS realmoji,
FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = 'active'
WHERE i.id = $1`

type imageServer struct {
	db     atomic.Pointer[sqlx.DB]
	state  *localImageState
	origin string // asia1 の :8081。手元に無い画像の本体はここから取る
}

func (s *imageServer) handler(c *echo.Context) error {
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
	id := input.ImageID
	ctx := c.Request().Context()
	db := s.db.Load()

	var data []byte
	haveData := false
	if localImageIDPattern.MatchString(id) && len(id) >= 2 {
		data, haveData = readLocalImage(id)
	} else {
		return NewAPIError(http.StatusForbidden, "image_forbidden")
	}

	entry, startSeq, trusted := s.state.lookup(id)
	valid := entry != nil
	if entry == nil {
		var row remoteImageCheck
		err := db.GetContext(ctx, &row, remoteImageCheckQuery, id)
		if errors.Is(err, sql.ErrNoRows) {
			return NewAPIError(http.StatusForbidden, "image_forbidden")
		}
		if err != nil {
			return err
		}
		key, err := imageSigningKey(row.SecretKeyHash)
		if err != nil {
			return err
		}
		entry = &localImageEntry{key: key, mediaType: row.MediaType}
		switch {
		case row.IsProfile, row.HasPost:
			valid = true
			entry.deps = []string{"img:" + id, "user:" + row.OwnerID}
		case row.RealMoji.Valid:
			valid = true
			author, postImage, _ := strings.Cut(row.RealMoji.String, " ")
			entry.deps = []string{"img:" + id, "user:" + row.OwnerID, "user:" + author, "img:" + postImage}
		}
		if valid && trusted {
			s.state.insert(id, entry, startSeq)
		}
	}
	if !hmac.Equal(sig, imageSignature(entry.key, id, exp)) {
		c.Set("abuse", "invalid_signature")
		return NewAPIError(http.StatusForbidden, "image_forbidden")
	}
	if !valid {
		return NewAPIError(http.StatusForbidden, "image_forbidden")
	}
	if !haveData {
		fetched, initial, err := fetchOriginImage(ctx, s.origin, id)
		if errors.Is(err, sql.ErrNoRows) {
			return NewAPIError(http.StatusForbidden, "image_forbidden")
		}
		if err != nil {
			return err
		}
		data = fetched
		writeLocalImage(id, data, initial)
	}
	if exp <= time.Now().Unix() {
		return NewAPIError(http.StatusForbidden, "image_forbidden")
	}
	return c.Blob(http.StatusOK, entry.mediaType, data)
}

// NewImageServer は ISUREAL_PASSIVE=1 の台で動かす。/api/images/:imageId を返し、/api/posts は画像を検証して asia1 へ送る（post_forward.go）。ほかの /api/ は nginx が asia1 へ送る
func NewImageServer(config *Config) *echo.Echo {
	e := echo.New()
	e.IPExtractor = extractClientIP
	e.Use(ActionLogMiddleware(config.LogLevel, os.Stdout))
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = APIErrorHandler

	peers := envPeers("ISUREAL_IMAGE_EVENT_PEERS")
	s := &imageServer{state: newLocalImageState(peers)}
	if len(peers) > 0 {
		s.origin = peers[0]
	}
	for _, p := range peers {
		go s.state.subscribe(p)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/_internal/invalidate", s.state.serveInvalidate)
	serveInternal(mux)
	bans := &IPBanList{}
	ready := make(chan struct{})
	go func() {
		db := openDatabase(config)
		db.SetMaxOpenConns(40)
		db.SetMaxIdleConns(40)
		s.db.Store(db)
		// ip_bans は裏で読み直し、リクエストでは DB を待たない
		for bans.snapshot.Load() == nil {
			if loaded, err := loadIPBans(context.Background(), db); err == nil {
				bans.snapshot.Store(loaded)
			} else {
				time.Sleep(time.Second)
			}
		}
		close(ready)
		for {
			if loaded, err := loadIPBans(context.Background(), db); err == nil {
				bans.snapshot.Store(loaded)
			}
			time.Sleep(ipBanReloadInterval / 2)
		}
	}()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			select {
			case <-ready:
			case <-c.Request().Context().Done():
				return c.Request().Context().Err()
			}
			return next(c)
		}
	})
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			if snapshot := bans.snapshot.Load(); snapshot != nil {
				if until, ok := snapshot.bans[c.RealIP()]; ok && time.Now().Before(until) {
					return NewAPIError(http.StatusForbidden, "ip_banned")
				}
			}
			return next(c)
		}
	})
	e.GET("/api/images/:imageId", s.handler)
	e.Any("/api/posts", postForwardHandler())
	return e
}
