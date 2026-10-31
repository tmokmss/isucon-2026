package isureal

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

// requestRegion はリクエストを受けた台のリージョン。各台の nginx が X-Isureal-Region に自分のリージョンを入れて転送する
func requestRegion(c *echo.Context, h *HandlerContext) string {
	if region := c.Request().Header.Get("X-Isureal-Region"); validRegion(region) {
		return region
	}
	return h.Config.Region
}

func RegionMiddleware(h *HandlerContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			session := GetSession(c)
			if session == nil {
				return NewAPIError(http.StatusUnauthorized, "invalid_session")
			}
			if session.Region != requestRegion(c, h) {
				return NewAPIError(http.StatusForbidden, "region_mismatch")
			}
			return next(c)
		}
	}
}
