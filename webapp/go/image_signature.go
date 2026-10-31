package isureal

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"sync"
	"time"
)

const imageURLLifetime = 24 * time.Hour
const imageSignatureIterations = 1000

// 同じ (鍵, 画像 ID, exp) の署名を使い回す。exp を 1 分単位に切り上げるので、同じ 1 分の間は同じ URL になる。
// exp ごとに分けて、一番新しい exp とその前の exp の分だけを持つ。計算中の署名は、同じものを頼んだ他のリクエストに待たせる
type imageSignatureCall struct {
	done  chan struct{}
	value []byte
}

var (
	imageSignatureCacheMu sync.RWMutex
	imageSignatureCache   = map[int64]map[string][]byte{}
	imageSignatureCalls   = map[string]*imageSignatureCall{}
)

func imageSigningKey(secretKeyHash string) ([]byte, error) {
	key, err := hex.DecodeString(secretKeyHash)
	if err != nil || len(key) != sha256.Size {
		return nil, fmt.Errorf("invalid image signing key")
	}
	return key, nil
}

func imageSignature(key []byte, imageID string, exp int64) []byte {
	cacheKey := imageSignatureCacheKey(key, imageID, exp)

	imageSignatureCacheMu.RLock()
	cached, ok := imageSignatureCache[exp][cacheKey]
	imageSignatureCacheMu.RUnlock()
	if ok {
		return cached
	}

	imageSignatureCacheMu.Lock()
	if cached, ok := imageSignatureCache[exp][cacheKey]; ok {
		imageSignatureCacheMu.Unlock()
		return cached
	}
	if call, ok := imageSignatureCalls[cacheKey]; ok {
		imageSignatureCacheMu.Unlock()
		<-call.done
		return call.value
	}
	call := &imageSignatureCall{done: make(chan struct{})}
	imageSignatureCalls[cacheKey] = call
	imageSignatureCacheMu.Unlock()

	call.value = computeImageSignature(key, imageID, exp)

	imageSignatureCacheMu.Lock()
	m := imageSignatureCache[exp]
	if m == nil {
		m = make(map[string][]byte)
		imageSignatureCache[exp] = m
		for e := range imageSignatureCache {
			if e < exp-60 {
				delete(imageSignatureCache, e)
			}
		}
	}
	m[cacheKey] = call.value
	delete(imageSignatureCalls, cacheKey)
	imageSignatureCacheMu.Unlock()
	close(call.done)
	return call.value
}

func computeImageSignature(key []byte, imageID string, exp int64) []byte {
	value := []byte(strconv.FormatInt(exp, 10) + ":" + imageID)
	mac := hmac.New(sha256.New, key)
	sum := make([]byte, 0, sha256.Size)
	for range imageSignatureIterations {
		mac.Reset()
		mac.Write(value)
		sum = mac.Sum(sum[:0])
		value = append(value[:0], sum...)
	}
	return value
}

func imageSignatureCacheKey(key []byte, imageID string, exp int64) string {
	return string(key) + strconv.FormatInt(exp, 10) + ":" + imageID
}

func imageURLExp() int64 {
	return time.Now().Add(imageURLLifetime).Add(time.Minute - time.Nanosecond).Truncate(time.Minute).Unix()
}

func signedImageURL(user User, imageID string) string {
	exp := imageURLExp()
	query := url.Values{
		"exp": {strconv.FormatInt(exp, 10)},
		"sig": {base64.RawURLEncoding.EncodeToString(imageSignature(user.ImageKey, imageID, exp))},
	}
	return "/api/images/" + url.PathEscape(imageID) + "?" + query.Encode()
}
