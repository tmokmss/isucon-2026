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
// exp ごとに分けて持ち、今の exp より古いものは捨てる（今の分と、先回りで入れる次の分だけが残る）
type imageSignatureEntry struct {
	sig       []byte
	presigned bool
}

var (
	imageSignatureCacheMu sync.RWMutex
	imageSignatureCaches  = map[int64]map[string]imageSignatureEntry{}
)

func imageSigningKey(secretKeyHash string) ([]byte, error) {
	key, err := hex.DecodeString(secretKeyHash)
	if err != nil || len(key) != sha256.Size {
		return nil, fmt.Errorf("invalid image signing key")
	}
	return key, nil
}

func lookupImageSignature(key []byte, imageID string, exp int64) (imageSignatureEntry, bool) {
	imageSignatureCacheMu.RLock()
	e, ok := imageSignatureCaches[exp][imageSignatureCacheKey(key, imageID)]
	imageSignatureCacheMu.RUnlock()
	return e, ok
}

func storeImageSignatures(exp int64, keys []string, sigs [][]byte, presigned bool) {
	cur := imageURLExp()
	imageSignatureCacheMu.Lock()
	for e := range imageSignatureCaches {
		if e < cur {
			delete(imageSignatureCaches, e)
		}
	}
	if exp >= cur {
		m := imageSignatureCaches[exp]
		if m == nil {
			m = make(map[string]imageSignatureEntry, len(keys))
			imageSignatureCaches[exp] = m
		}
		for i, k := range keys {
			m[k] = imageSignatureEntry{sigs[i], presigned}
		}
	}
	imageSignatureCacheMu.Unlock()
}

func imageSignature(key []byte, imageID string, exp int64) []byte {
	if e, ok := lookupImageSignature(key, imageID, exp); ok {
		if e.presigned {
			presignHit.Add(1)
		} else {
			presignCacheHit.Add(1)
		}
		return e.sig
	}
	presignLocal.Add(1)
	recordPresignTarget(key, imageID)
	value := computeImageSignature(key, imageID, exp)
	storeImageSignatures(exp, []string{imageSignatureCacheKey(key, imageID)}, [][]byte{value}, false)
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

func imageSignatureCacheKey(key []byte, imageID string) string {
	return string(key) + imageID
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
