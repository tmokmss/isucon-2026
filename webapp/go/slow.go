package isureal

import (
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/labstack/echo/v5"
)

// /api/ のリクエストに一様な待ちを入れる。/api/initialize と /api/admin/ には入れない（制限時間・タイムアウトで FAIL になる）。
// 画像は europe1・america1 の台が initialize を受けずに返すので、待ちを入れない。
// initialize の後、最初に投稿期間が開かれるまでは待たない。投稿期間が開かれた後と、initialize を受けていないプロセスでは slowDelay 待つ。
const slowDelay = 3000 * time.Millisecond

// true の間は待たない
var slowPaused atomic.Bool

func SlowMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			p := c.Request().URL.Path
			switch {
			case p == "/api/initialize":
				err := next(c)
				if err == nil {
					slowPaused.Store(true)
				}
				return err
			case p == "/api/admin/posting-window" && c.Request().Method == http.MethodPost:
				slowPaused.Store(false)
				return next(c)
			case strings.HasPrefix(p, "/api/admin/") || strings.HasPrefix(p, "/api/images/"):
				return next(c)
			}
			if !slowPaused.Load() {
				time.Sleep(slowDelay)
			}
			return next(c)
		}
	}
}
