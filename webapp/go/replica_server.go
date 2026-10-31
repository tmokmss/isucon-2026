package isureal

// europe1・america1 で ISUREAL_REPLICA_MODE を設定したときの入口。nginx は /api/（SSE を除く）をこの台へ送る。
// - /api/images/ は今まで通り image_local.go が返す
// - 複製が追いついている間（replica.fresh）は、replicaLocalRoute の GET を手元のストアから返す
//   - shadow: asia1 へ転送した応答を返し、手元で前後 2 回作った応答のどちらとも違えば repl-diff をログに出す
//   - local: 手元の応答を返す
// - それ以外は asia1 へ転送する
// ISUREAL_REPLICA_MODE を空にすると、image_local.go だけの台に戻る（nginx は /api/ を asia1 へ直接送る設定に戻す）

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func replicaLocalRoute(path string) bool {
	switch path {
	case "/api/posting-window/current", "/api/posts/me/current", "/api/users", "/api/me", "/api/me/export", "/api/memories":
		return true
	}
	parts := strings.Split(strings.TrimPrefix(path, "/api/"), "/")
	switch {
	case len(parts) == 2 && (parts[0] == "feeds" || parts[0] == "users" || parts[0] == "relationships" || parts[0] == "memories") && parts[1] != "":
		return true
	case len(parts) == 2 && parts[0] == "posts" && parts[1] != "":
		return true
	case len(parts) == 3 && parts[0] == "posts" && (parts[2] == "comments" || parts[2] == "realmojis"):
		return true
	}
	return false
}

// exportLocalImage は複製の台で、エクスポートの画像を手元のファイル（image_local.go が置いたもの）から読む
var exportLocalImage = func(id string) ([]byte, bool) { return nil, false }

type replicaFront struct {
	mode   string
	rep    *replica
	images http.Handler
	local  http.Handler
	proxy  *httputil.ReverseProxy

	shadowTotal atomic.Int64
	shadowDiff  atomic.Int64
	shadowLate  atomic.Int64
	routeMu     sync.Mutex
	routeCounts map[string]*[3]int64 // 比べた数・食い違い・届く順の差
	localTotal  atomic.Int64
	// 追いつくのを待った時間の合計と、待ちきれずに転送した数
	waitedTotal   atomic.Int64
	fallbackTotal atomic.Int64
}

func NewReplicaServer(config *Config) http.Handler {
	upstream := os.Getenv("ISUREAL_REPLICA_UPSTREAM") // asia1 の :8080
	target, _ := url.Parse("http://" + upstream)
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = &http.Transport{
		DialContext:         (&net.Dialer{Timeout: 3 * time.Second, KeepAlive: 30 * time.Second}).DialContext,
		MaxIdleConns:        1024,
		MaxIdleConnsPerHost: 1024,
		IdleConnTimeout:     90 * time.Second,
	}

	h := &HandlerContext{Config: config, IPBans: &IPBanList{}, Store: NewStore()}
	rep := &replica{s: h.Store, upstream: os.Getenv("ISUREAL_REPLICA_STREAM")}
	advanced := make(chan struct{})
	rep.advanced.Store(&advanced)
	go func() {
		db := openDatabase(config)
		db.SetMaxOpenConns(20)
		db.SetMaxIdleConns(20)
		h.DB = db
		rep.db = db
		rep.run()
	}()

	e := echo.New()
	e.IPExtractor = extractClientIP
	e.Use(ActionLogMiddleware(config.LogLevel, os.Stdout))
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = APIErrorHandler
	e.Use(IPBanMiddleware(h))
	e.GET("/api/posts/:postId", getPostHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	e.GET("/api/feeds/:feed", getFeedHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	e.GET("/api/posts/:postId/comments", getCommentsHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	e.GET("/api/posts/:postId/realmojis", getRealMojisHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	e.GET("/api/posting-window/current", getCurrentPostingWindowHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	e.GET("/api/posts/me/current", getMyCurrentPostHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	e.GET("/api/users", searchUsersHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	e.GET("/api/users/:userId", getUserHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	e.GET("/api/me", getMeHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	e.GET("/api/me/export", exportPostsHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	e.GET("/api/memories", getMemoriesHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	e.GET("/api/memories/:date", getMemoriesByDateHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	e.GET("/api/relationships/:kind", getRelationshipsHandler(h), SessionMiddleware(h, true), RegionMiddleware(h))
	// 新しい画像の本体は PostgreSQL に無い（image_file.go）ので、手元に無ければ asia1 から取って置く
	exportLocalImage = func(id string) ([]byte, bool) {
		if !localImageIDPattern.MatchString(id) || len(id) < 2 {
			return nil, false
		}
		if data, ok := readLocalImage(id); ok {
			return data, true
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		data, initial, err := fetchOriginImage(ctx, rep.upstream, id)
		if err != nil {
			return nil, false
		}
		writeLocalImage(id, data, initial)
		return data, true
	}

	f := &replicaFront{mode: os.Getenv("ISUREAL_REPLICA_MODE"), rep: rep, images: NewImageServer(config), local: e, proxy: proxy}
	go f.statsLoop()
	return f
}

func (f *replicaFront) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 画像と投稿（画像を検証してから asia1 へ送る、post_forward.go）は image_local.go の入口が受ける
	if strings.HasPrefix(r.URL.Path, "/api/images/") || r.URL.Path == "/api/posts" {
		f.images.ServeHTTP(w, r)
		return
	}
	if r.Method == http.MethodGet && replicaLocalRoute(r.URL.Path) && f.rep.live.Load() {
		start := time.Now()
		if f.rep.waitCaughtUp(start.Add(replClockMargin).UnixNano()) {
			f.waitedTotal.Add(int64(time.Since(start)))
			switch f.mode {
			case "local":
				f.localTotal.Add(1)
				f.local.ServeHTTP(w, r)
				return
			case "shadow":
				f.shadow(w, r)
				return
			}
		} else {
			f.fallbackTotal.Add(1)
		}
	}
	f.proxy.ServeHTTP(w, r)
}

func (f *replicaFront) record(h http.Handler, r *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, r.Clone(context.Background()))
	return rec
}

func (f *replicaFront) shadow(w http.ResponseWriter, r *http.Request) {
	before := f.record(f.local, r)
	rec := httptest.NewRecorder()
	f.proxy.ServeHTTP(rec, r)
	fresh := f.rep.fresh()
	after := f.record(f.local, r)
	for k, v := range rec.Header() {
		w.Header()[k] = v
	}
	w.WriteHeader(rec.Code)
	w.Write(rec.Body.Bytes())

	f.shadowTotal.Add(1)
	route := routeKey(r.URL.Path)
	f.countRoute(route, 0)
	if !fresh || rec.Code == http.StatusBadGateway || sameResponse(rec, before) || sameResponse(rec, after) {
		return
	}
	// asia1 の応答と変更は同じ頃に届くので、少し待って作り直して一致すれば届く順の差として数える
	r2 := r.Clone(context.Background())
	go func() {
		for _, d := range []time.Duration{100 * time.Millisecond, 300 * time.Millisecond} {
			time.Sleep(d)
			if sameResponse(rec, f.record(f.local, r2)) {
				f.shadowLate.Add(1)
				f.countRoute(route, 2)
				return
			}
		}
		n := f.shadowDiff.Add(1)
		f.countRoute(route, 1)
		if n <= 1000 {
			log.Printf("repl-diff %s %s asia=%d local=%d/%d\n  asia:   %s\n  before: %s\n  after:  %s", r2.Method, r2.URL.RequestURI(), rec.Code, before.Code, after.Code,
				diffContext(rec.Body.Bytes(), before.Body.Bytes()), diffContext(before.Body.Bytes(), rec.Body.Bytes()), diffContext(after.Body.Bytes(), rec.Body.Bytes()))
		}
	}()
}

func sameResponse(a, b *httptest.ResponseRecorder) bool {
	return a.Code == b.Code && bytes.Equal(bytes.TrimSpace(a.Body.Bytes()), bytes.TrimSpace(b.Body.Bytes()))
}

// diffContext は a の、b と最初に食い違う位置の前後を返す
func diffContext(a, b []byte) string {
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	lo, hi := max(0, i-300), min(len(a), i+300)
	return string(a[lo:hi])
}

func routeKey(path string) string {
	parts := strings.Split(strings.TrimPrefix(path, "/api/"), "/")
	switch {
	case parts[0] == "feeds" || parts[0] == "relationships" || (parts[0] == "posts" && len(parts) == 2 && parts[1] == "me"):
		return strings.Join(parts, "/")
	case len(parts) >= 2 && parts[0] != "me" && parts[0] != "posting-window":
		parts[1] = ":id"
	}
	return strings.Join(parts, "/")
}

func (f *replicaFront) countRoute(route string, i int) {
	f.routeMu.Lock()
	defer f.routeMu.Unlock()
	if f.routeCounts == nil {
		f.routeCounts = map[string]*[3]int64{}
	}
	c := f.routeCounts[route]
	if c == nil {
		c = &[3]int64{}
		f.routeCounts[route] = c
	}
	c[i]++
}

func (f *replicaFront) statsLoop() {
	for range time.Tick(10 * time.Second) {
		n := f.localTotal.Load() + f.shadowTotal.Load()
		var avgWait time.Duration
		if n > 0 {
			avgWait = time.Duration(f.waitedTotal.Load() / n)
		}
		var avgLag time.Duration
		if m := f.rep.lagCount.Load(); m > 0 {
			avgLag = time.Duration(f.rep.lagTotal.Load() / m)
		}
		log.Printf("repl-stats fresh=%v shadow=%d diff=%d late=%d local=%d fallback=%d avg_wait=%s avg_lag=%s", f.rep.fresh(), f.shadowTotal.Load(), f.shadowDiff.Load(), f.shadowLate.Load(), f.localTotal.Load(), f.fallbackTotal.Load(), avgWait, avgLag)
		avg := func(total, count *atomic.Int64) time.Duration {
			n := count.Swap(0)
			s := total.Swap(0)
			if n == 0 {
				return 0
			}
			return time.Duration(s / n)
		}
		log.Printf("repl-lag recv=%s lock_wait=%s", avg(&f.rep.recvLagTotal, &f.rep.recvLagCount), avg(&f.rep.lockWaitTotal, &f.rep.lockWaitCount))
		f.routeMu.Lock()
		var b strings.Builder
		for k, c := range f.routeCounts {
			fmt.Fprintf(&b, " %s=%d/%d/%d", k, c[0], c[1], c[2])
		}
		f.routeMu.Unlock()
		if b.Len() > 0 {
			log.Printf("repl-routes (compared/diff/late)%s", b.String())
		}
	}
}
