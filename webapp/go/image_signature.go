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

// 同じ (鍵, 画像 ID, exp) の署名を使い回す。exp を 1 分単位に切り上げるので、同じ 1 分の間は同じ URL になる
const imageSignatureCacheMax = 1_000_000

var (
	imageSignatureCacheMu sync.RWMutex
	imageSignatureCache   = make(map[string][]byte)
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
	cached, ok := imageSignatureCache[cacheKey]
	imageSignatureCacheMu.RUnlock()
	if ok {
		return cached
	}

	value := computeImageSignature(key, imageID, exp)

	imageSignatureCacheMu.Lock()
	if len(imageSignatureCache) >= imageSignatureCacheMax {
		imageSignatureCache = make(map[string][]byte)
	}
	imageSignatureCache[cacheKey] = value
	imageSignatureCacheMu.Unlock()
	shareSignature(cacheKey, value)
	return value
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
