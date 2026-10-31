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

// 署名は 1000 回の HMAC で重いので、同じ画像に同じ exp を使う間は使い回す。
// exp を 10 秒単位に切り下げて揃える（有効期限が最大 10 秒短くなるだけで、形式と検証は変わらない）
const imageExpGranularity = 10

type sigEntry struct {
	exp int64
	sig string
}

var signedCache sync.Map   // imageID -> sigEntry（発行したもの）
var verifiedCache sync.Map // imageID -> sigEntry（検証に通ったもの）

func imageSigningKey(secretKeyHash string) ([]byte, error) {
	key, err := hex.DecodeString(secretKeyHash)
	if err != nil || len(key) != sha256.Size {
		return nil, fmt.Errorf("invalid image signing key")
	}
	return key, nil
}

func imageSignature(key []byte, imageID string, exp int64) []byte {
	value := []byte(strconv.FormatInt(exp, 10) + ":" + imageID)
	mac := hmac.New(sha256.New, key)
	for range imageSignatureIterations {
		mac.Reset()
		mac.Write(value)
		value = mac.Sum(value[:0:0])
	}
	return value
}

func signedImageURL(user User, imageID string) string {
	exp := time.Now().Add(imageURLLifetime).Unix()
	exp -= exp % imageExpGranularity
	var sig string
	if v, ok := signedCache.Load(imageID); ok && v.(sigEntry).exp == exp {
		sig = v.(sigEntry).sig
	} else {
		sig = base64.RawURLEncoding.EncodeToString(imageSignature(user.ImageKey, imageID, exp))
		signedCache.Store(imageID, sigEntry{exp: exp, sig: sig})
	}
	return "/api/images/" + url.PathEscape(imageID) + "?exp=" + strconv.FormatInt(exp, 10) + "&sig=" + sig
}

// verifyImageSignature は sigText（base64url）と sig（デコード済み）が正しいかを返す
func verifyImageSignature(key []byte, imageID string, exp int64, sigText string, sig []byte) bool {
	if v, ok := signedCache.Load(imageID); ok && v.(sigEntry).exp == exp && v.(sigEntry).sig == sigText {
		return true
	}
	if v, ok := verifiedCache.Load(imageID); ok && v.(sigEntry).exp == exp && v.(sigEntry).sig == sigText {
		return true
	}
	if hmac.Equal(sig, imageSignature(key, imageID, exp)) {
		verifiedCache.Store(imageID, sigEntry{exp: exp, sig: sigText})
		return true
	}
	return false
}
