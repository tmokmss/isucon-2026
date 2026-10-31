package isureal

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// この台で署名した (鍵, 画像 ID) を覚えておき、毎分 35 秒を過ぎたら次の exp の署名を
// ISUREAL_PRESIGN_PEERS（SHA-NI のある europe1・america1）にまとめて計算させ、次の分のキャッシュに入れる。
// 計算が間に合わなかった分や失敗した分は、応答の経路が今まで通り手元で計算する
const (
	presignStartSecond = 35
	presignMax         = 300_000
	presignBatch       = 2000
	presignParallel    = 8
	presignTimeout     = 10 * time.Second
)

var (
	presignPeers = offloadPeersFromEnv("ISUREAL_PRESIGN_PEERS")

	presignMu      sync.Mutex
	presignTargets = map[string]signTarget{}

	presignHit      atomic.Int64
	presignCacheHit atomic.Int64
	presignLocal    atomic.Int64
)

func init() {
	if len(presignPeers) > 0 {
		go presignNextMinuteForever()
	}
}

func recordPresignTarget(key []byte, imageID string) {
	if len(presignPeers) == 0 {
		return
	}
	id := imageSignatureCacheKey(key, imageID)
	presignMu.Lock()
	if len(presignTargets) < presignMax {
		presignTargets[id] = signTarget{key, imageID}
	}
	presignMu.Unlock()
}

func presignNextMinuteForever() {
	var done int64
	for range time.Tick(time.Second) {
		if time.Now().Second() < presignStartSecond {
			continue
		}
		next := imageURLExp() + 60
		if done == next {
			continue
		}
		done = next

		presignMu.Lock()
		total := len(presignTargets)
		targets := make([]signTarget, 0, total)
		for _, t := range presignTargets {
			targets = append(targets, t)
		}
		presignMu.Unlock()

		start := time.Now()
		var ok atomic.Int64
		var wg sync.WaitGroup
		sem := make(chan struct{}, presignParallel)
		for len(targets) > 0 {
			n := min(presignBatch, len(targets))
			batch := targets[:n]
			targets = targets[n:]
			sem <- struct{}{}
			wg.Go(func() {
				defer func() { <-sem }()
				ok.Add(int64(presignImageSignatures(batch, next)))
			})
		}
		wg.Wait()
		log.Printf("presign exp=%d targets=%d signed=%d in %s; previous minute: presign_hit=%d cache_hit=%d local=%d",
			next, total, ok.Load(), time.Since(start).Round(time.Millisecond), presignHit.Swap(0), presignCacheHit.Swap(0), presignLocal.Swap(0))
	}
}

func presignImageSignatures(targets []signTarget, exp int64) int {
	var body []byte
	keys := make([]string, 0, len(targets))
	for _, t := range targets {
		keys = append(keys, imageSignatureCacheKey(t.key, t.imageID))
		body = appendSignRequest(body, t.key, t.imageID, exp)
	}
	res, err := offloadCall(presignPeers, "/sign", body, presignTimeout)
	if err != nil || len(res) != len(keys)*32 {
		return 0
	}
	storeImageSignatures(exp, keys, splitSignatures(res, len(keys)), true)
	return len(keys)
}
