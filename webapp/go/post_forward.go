package isureal

// europe1・america1 が受けた投稿は、その台で画像を jpegcheck.Check で検証してから asia1 へ送る。
// 検証に通ったときだけ、幅・高さと画像の本体に付けた HMAC（鍵は ISUREAL_JPEG_VERIFY_KEY、台の間で共有）を X-Isureal-Jpeg に入れる。
// asia1 は HMAC が合えば検証を省き、合わない・無いときは今までどおり検証する。各台の nginx はクライアントの X-Isureal-Jpeg を消す。

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v5"

	"example.com/isureal-company/isureal/jpegcheck"
)

const jpegVerifiedHeader = "X-Isureal-Jpeg"

var jpegVerifyKey = []byte(os.Getenv("ISUREAL_JPEG_VERIFY_KEY"))

func jpegVerifiedMAC(width, height int, image []byte) []byte {
	m := hmac.New(sha256.New, jpegVerifyKey)
	var b [8]byte
	binary.BigEndian.PutUint32(b[0:4], uint32(width))
	binary.BigEndian.PutUint32(b[4:8], uint32(height))
	m.Write(b[:])
	m.Write(image)
	return m.Sum(nil)
}

// 投稿の画像を jpegcheck.Check と同じ判定で検証する。europe1・america1 の検証済みの印が合えば、その結果を使う
func checkPostJPEG(req *http.Request, image []byte) (width, height int, err error) {
	if v := req.Header.Get(jpegVerifiedHeader); v != "" && len(jpegVerifyKey) > 0 {
		if parts := strings.Split(v, "."); len(parts) == 3 {
			w, errW := strconv.Atoi(parts[0])
			h, errH := strconv.Atoi(parts[1])
			mac, errM := hex.DecodeString(parts[2])
			if errW == nil && errH == nil && errM == nil && w >= 0 && h >= 0 && hmac.Equal(mac, jpegVerifiedMAC(w, h, image)) {
				return w, h, nil
			}
		}
	}
	return checkJPEG(image)
}

// readPostImage と同じパートを画像として取り出す
func postImageFromBody(header http.Header, body []byte) ([]byte, bool) {
	if !strings.HasPrefix(header.Get(echo.HeaderContentType), echo.MIMEMultipartForm) {
		return nil, false
	}
	mr, err := (&http.Request{Header: header, Body: io.NopCloser(bytes.NewReader(body))}).MultipartReader()
	if err != nil {
		return nil, false
	}
	for {
		p, err := mr.NextPart()
		if err != nil {
			return nil, false
		}
		if p.FormName() == "image" && p.FileName() != "" {
			image, err := io.ReadAll(p)
			return image, err == nil
		}
	}
}

// /api/posts を asia1（ISUREAL_POST_UPSTREAM）へ送る。POST の画像はこの台で検証して印を付ける
func postForwardHandler() echo.HandlerFunc {
	target := &url.URL{Scheme: "http", Host: os.Getenv("ISUREAL_POST_UPSTREAM")}
	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(target)
			pr.Out.Host = pr.In.Host
		},
		Transport: &http.Transport{
			DialContext:         (&net.Dialer{Timeout: 3 * time.Second}).DialContext,
			MaxIdleConns:        256,
			MaxIdleConnsPerHost: 256,
			IdleConnTimeout:     90 * time.Second,
		},
	}
	return func(c *echo.Context) error {
		req := c.Request()
		req.Header.Del(jpegVerifiedHeader)
		var body bytes.Buffer
		body.Grow(max(int(req.ContentLength), 0) + bytes.MinRead)
		if _, err := body.ReadFrom(req.Body); err != nil {
			return err
		}
		if req.Method == http.MethodPost {
			if image, ok := postImageFromBody(req.Header, body.Bytes()); ok {
				if w, h, err := jpegcheck.Check(bytes.NewReader(image)); err == nil {
					req.Header.Set(jpegVerifiedHeader, strconv.Itoa(w)+"."+strconv.Itoa(h)+"."+hex.EncodeToString(jpegVerifiedMAC(w, h, image)))
				}
			}
		}
		req.Body = io.NopCloser(bytes.NewReader(body.Bytes()))
		req.ContentLength = int64(body.Len())
		proxy.ServeHTTP(c.Response(), req)
		return nil
	}
}
