package isureal

import (
	"time"

	"github.com/labstack/echo/v5"
)

// すべての /api/ のリクエストに一定の待ちを入れる。/api/initialize は 60 秒の制限があるので除く。
const slowDelay = 300 * time.Millisecond

func SlowMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			if c.Request().URL.Path != "/api/initialize" {
				time.Sleep(slowDelay)
			}
			return next(c)
		}
	}
}
