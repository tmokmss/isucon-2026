package isureal

// asia2 が出した画像 URL の取得は、画像を持つ asia1 に来て、asia1 が同じ署名を検証のために計算し直す。
// そこで、この台で計算した署名を ISUREAL_SIGN_SHARE_PEER（asia1 の :8090）へ 10ms ごと・最大 1000 件ずつ送り、asia1 の署名のキャッシュに入れる。
// 1 件 = キャッシュのキーの長さ(2) キー 署名(32)。送れなかった分は捨てる（asia1 は今まで通り手元で計算する）

import (
	"encoding/binary"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	signShareBatch    = 1000
	signShareInterval = 10 * time.Millisecond
)

var signShareCh = func() chan signShareEntry {
	if os.Getenv("ISUREAL_SIGN_SHARE_PEER") == "" {
		return nil
	}
	ch := make(chan signShareEntry, 100_000)
	go signShareForever(ch, "http://"+os.Getenv("ISUREAL_SIGN_SHARE_PEER")+"/signcache")
	return ch
}()

type signShareEntry struct {
	key string
	sig []byte
}

func shareSignature(cacheKey string, sig []byte) {
	if signShareCh == nil {
		return
	}
	select {
	case signShareCh <- signShareEntry{cacheKey, sig}:
	default:
	}
}

func signShareForever(ch chan signShareEntry, url string) {
	var body []byte
	n := 0
	tick := time.NewTicker(signShareInterval)
	for {
		select {
		case e := <-ch:
			body = binary.BigEndian.AppendUint16(body, uint16(len(e.key)))
			body = append(body, e.key...)
			body = append(body, e.sig...)
			n++
			if n < signShareBatch {
				continue
			}
		case <-tick.C:
			if n == 0 {
				continue
			}
		}
		offloadPost(url, body, time.Second)
		body = body[:0]
		n = 0
	}
}

func parseSignShare(b []byte) (keys []string, sigs [][]byte, err error) {
	for len(b) > 0 {
		if len(b) < 2 {
			return nil, nil, errors.New("short request")
		}
		m := int(binary.BigEndian.Uint16(b))
		if len(b) < 2+m+32 {
			return nil, nil, errors.New("short request")
		}
		keys = append(keys, string(b[2:2+m]))
		sigs = append(sigs, append([]byte(nil), b[2+m:2+m+32]...))
		b = b[2+m+32:]
	}
	return keys, sigs, nil
}

func signShareHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	keys, sigs, err := parseSignShare(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	imageSignatureCacheMu.Lock()
	if len(imageSignatureCache)+len(keys) > imageSignatureCacheMax {
		imageSignatureCache = make(map[string][]byte)
	}
	for i, k := range keys {
		imageSignatureCache[k] = sigs[i]
	}
	imageSignatureCacheMu.Unlock()
}
