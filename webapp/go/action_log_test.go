package isureal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v5"
)

// isureal-guard が BAN に数える失敗とエラーだけが書かれ、書かれる行の中身は変わらない
func TestActionLogWritesGuardFailuresAndErrors(t *testing.T) {
	cases := []struct {
		method, path, route string
		status              int
		abuse               string
		want                bool
	}{
		{http.MethodPost, "/api/auth/login", "/api/auth/login", http.StatusUnauthorized, "", true},
		{http.MethodPost, "/api/auth/login", "/api/auth/login", http.StatusOK, "", false},
		{http.MethodPost, "/api/auth/login", "/api/auth/login", http.StatusBadRequest, "", false},
		{http.MethodPut, "/api/me/profile-image", "/api/me/profile-image", http.StatusBadRequest, "invalid_image", true},
		{http.MethodPut, "/api/me/profile-image", "/api/me/profile-image", http.StatusBadRequest, "", false},
		{http.MethodPut, "/api/me/profile-image", "/api/me/profile-image", http.StatusOK, "", false},
		{http.MethodGet, "/api/images/x", "/api/images/:imageId", http.StatusForbidden, "invalid_signature", true},
		{http.MethodGet, "/api/images/x", "/api/images/:imageId", http.StatusForbidden, "", false},
		{http.MethodGet, "/api/images/x", "/api/images/:imageId", http.StatusOK, "", false},
		{http.MethodPost, "/api/posts", "/api/posts", http.StatusCreated, "", false},
		{http.MethodPost, "/api/posts", "/api/posts", http.StatusInternalServerError, "", true},
		{http.MethodGet, "/api/feeds/friends", "/api/feeds/:feed", http.StatusInternalServerError, "", true},
	}
	for _, tc := range cases {
		var out bytes.Buffer
		e := echo.New()
		e.Use(ActionLogMiddleware("info", &out))
		handler := func(c *echo.Context) error {
			if tc.abuse != "" {
				c.Set("abuse", tc.abuse)
			}
			return c.NoContent(tc.status)
		}
		e.Add(tc.method, tc.route, handler)
		req := httptest.NewRequest(tc.method, tc.path, nil)
		req.RemoteAddr = "192.0.2.1:1234"
		e.ServeHTTP(httptest.NewRecorder(), req)
		if got := out.Len() > 0; got != tc.want {
			t.Fatalf("%s %s %d %q: written=%v, want %v", tc.method, tc.path, tc.status, tc.abuse, got, tc.want)
		}
		if !tc.want {
			continue
		}
		var entry struct {
			Action *string `json:"action"`
			Status int     `json:"status"`
			IP     string  `json:"ip"`
			Abuse  string  `json:"abuse"`
		}
		if err := json.Unmarshal(out.Bytes(), &entry); err != nil {
			t.Fatal(err)
		}
		if entry.Action == nil || *entry.Action != actionNames[tc.method+" "+tc.route] || entry.Status != tc.status || entry.IP == "" || entry.Abuse != tc.abuse {
			t.Fatalf("%s %s: unexpected entry %s", tc.method, tc.path, out.String())
		}
	}
}
