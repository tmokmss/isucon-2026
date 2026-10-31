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

// 署名を exp（1 分単位に切り上げ）ごとの世代に分けて使い回す。直前の分と今の分だけを持ち、古い世代は丸ごと捨てる。
// キーと値はポインタを含まない固定長にし、GC が走査しないようにする。画像 ID が UUID の長さでなければ使い回さない
const imageIDLen = 36

type imageSigKey struct {
	key [sha256.Size]byte
	id  [imageIDLen]byte
}

var (
	imageSignatureCacheMu sync.RWMutex
	imageSignatureCache   = map[int64]map[imageSigKey][sha256.Size]byte{}
)

func makeImageSigKey(key []byte, imageID string) (imageSigKey, bool) {
	var k imageSigKey
	if len(key) != sha256.Size || len(imageID) != imageIDLen {
		return k, false
	}
	copy(k.key[:], key)
	copy(k.id[:], imageID)
	return k, true
}

// imageSignatureGenLocked は imageSignatureCacheMu の書きロックを持って呼ぶ。直前の分より古い世代はここで捨てる
func imageSignatureGenLocked(exp int64) map[imageSigKey][sha256.Size]byte {
	gen := imageSignatureCache[exp]
	if gen == nil {
		gen = make(map[imageSigKey][sha256.Size]byte)
		imageSignatureCache[exp] = gen
		oldest := imageURLExp() - 60
		for e := range imageSignatureCache {
			if e < oldest {
				delete(imageSignatureCache, e)
			}
		}
	}
	return gen
}

func putImageSignaturesLocked(exp int64, keys []imageSigKey, sigs []byte) {
	if exp < imageURLExp()-60 {
		return
	}
	gen := imageSignatureGenLocked(exp)
	for i, k := range keys {
		gen[k] = [sha256.Size]byte(sigs[i*sha256.Size : (i+1)*sha256.Size])
	}
}

func imageSigningKey(secretKeyHash string) ([]byte, error) {
	key, err := hex.DecodeString(secretKeyHash)
	if err != nil || len(key) != sha256.Size {
		return nil, fmt.Errorf("invalid image signing key")
	}
	return key, nil
}

func imageSignature(key []byte, imageID string, exp int64) []byte {
	k, cacheable := makeImageSigKey(key, imageID)
	if !cacheable {
		return computeImageSignature(key, imageID, exp)
	}
	imageSignatureCacheMu.RLock()
	cached, ok := imageSignatureCache[exp][k]
	imageSignatureCacheMu.RUnlock()
	if ok {
		return cached[:]
	}

	value := computeImageSignature(key, imageID, exp)

	imageSignatureCacheMu.Lock()
	putImageSignaturesLocked(exp, []imageSigKey{k}, value)
	imageSignatureCacheMu.Unlock()
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
