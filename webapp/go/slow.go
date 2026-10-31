package isureal

import (
	"strings"
	"time"

	"github.com/labstack/echo/v5"
)

// /api/ のリクエストに一様な待ちを入れる。直近 10 分以内に始まった投稿期間が開いている間だけ待つ
// （初期データの投稿期間はすべて終わっているので、initialize の直後は待たない）。
// /api/initialize・/api/admin/・/api/images/ には入れない（制限時間・タイムアウトで FAIL になる。画像は europe1・america1 の別の入口が返す）。
const slowDelay = 3000 * time.Millisecond

func slowActive(s *Store) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, w := range s.activeWindows {
		if time.Since(w.StartsAt) < 10*time.Minute {
			return true
		}
	}
	return false
}

func SlowMiddleware(h *HandlerContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			p := c.Request().URL.Path
			if p != "/api/initialize" && !strings.HasPrefix(p, "/api/admin/") && !strings.HasPrefix(p, "/api/images/") && slowActive(h.Store) {
				time.Sleep(slowDelay)
			}
			return next(c)
		}
	}
}
