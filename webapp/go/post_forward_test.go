package isureal

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v5"

	"example.com/isureal-company/isureal/jpegcheck"
)

func postBody(t *testing.T, img []byte) (string, []byte) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	w.WriteField("caption", "hi")
	fw, _ := w.CreateFormFile("image", "a.jpg")
	fw.Write(img)
	w.WriteField("visibility", "public")
	w.Close()
	return w.FormDataContentType(), b.Bytes()
}

func TestPostForwardMarksVerifiedJPEG(t *testing.T) {
	jpegVerifyKey = []byte("test-key")
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, image.NewRGBA(image.Rect(0, 0, 33, 17)), nil); err != nil {
		t.Fatal(err)
	}
	valid := buf.Bytes()
	wantW, wantH, err := jpegcheck.Check(bytes.NewReader(valid))
	if err != nil {
		t.Fatal(err)
	}

	var got *http.Request
	var gotBody []byte
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r
		gotBody, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"ok":true}`))
	}))
	defer upstream.Close()
	u, _ := url.Parse(upstream.URL)
	t.Setenv("ISUREAL_POST_UPSTREAM", u.Host)
	e := echo.New()
	e.Any("/api/posts", postForwardHandler())

	send := func(img []byte, forged string) *httptest.ResponseRecorder {
		ct, body := postBody(t, img)
		req := httptest.NewRequest(http.MethodPost, "/api/posts", bytes.NewReader(body))
		req.Header.Set("Content-Type", ct)
		req.Header.Set("Cookie", "session=x")
		if forged != "" {
			req.Header.Set(jpegVerifiedHeader, forged)
		}
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		if !bytes.Equal(gotBody, body) || got.Header.Get("Cookie") != "session=x" {
			t.Fatalf("body or headers changed")
		}
		return rec
	}

	rec := send(valid, "")
	if rec.Code != http.StatusCreated || rec.Body.String() != `{"ok":true}` {
		t.Fatalf("response %d %s", rec.Code, rec.Body.String())
	}
	if got.Header.Get(jpegVerifiedHeader) == "" {
		t.Fatal("missing mark")
	}
	w, h, err := checkPostJPEG(got, valid)
	if err != nil || w != wantW || h != wantH {
		t.Fatalf("checkPostJPEG = %d %d %v", w, h, err)
	}
	// 印と違う画像は印を使わず検証する
	broken := append([]byte{}, valid[:len(valid)/2]...)
	if _, _, err := checkPostJPEG(got, broken); err == nil {
		t.Fatal("mark accepted for another image")
	}

	// 壊れた画像には印を付けず、クライアントが付けた印は消す
	send(broken, got.Header.Get(jpegVerifiedHeader))
	if got.Header.Get(jpegVerifiedHeader) != "" {
		t.Fatal("mark on invalid image")
	}
}
