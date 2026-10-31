package isureal

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

// asia1（2 vCPU）で CPU を食う処理の同時実行数を絞り、投稿に CPU を回す。
// 待ちが maxWait を超えたら 503 を返す（ベンチは 503 を再試行する）。クライアントが切ったら待つのをやめる。
type limiter struct {
	slots   chan struct{}
	maxWait time.Duration
}

var (
	postLimiter  = &limiter{slots: make(chan struct{}, 4), maxWait: 8 * time.Second}
	heavyLimiter = &limiter{slots: make(chan struct{}, 1), maxWait: 4 * time.Second}
)

func (l *limiter) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		timer := time.NewTimer(l.maxWait)
		defer timer.Stop()
		select {
		case l.slots <- struct{}{}:
		case <-timer.C:
			return NewAPIError(http.StatusServiceUnavailable, "busy")
		case <-c.Request().Context().Done():
			return c.Request().Context().Err()
		}
		defer func() { <-l.slots }()
		return next(c)
	}
}
