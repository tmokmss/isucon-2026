package main

import (
	"database/sql"
	"net/http"
	"time"
)

// マッチングの間隔。isuride-matcher（0.5 秒ごと）を待たずにアプリの中で回す
const matchingInterval = 100 * time.Millisecond

// startMatcher はアプリの中でマッチングを定期的に回す。isuride-matcher からの呼び出しもそのまま受け付ける
func startMatcher() {
	go func() {
		t := time.NewTicker(matchingInterval)
		for range t.C {
			matchRides()
		}
	}()
}

// このAPIをインスタンス内から一定間隔で叩かせることで、椅子とライドをマッチングさせる
func internalGetMatching(w http.ResponseWriter, r *http.Request) {
	matchRides()
	w.WriteHeader(http.StatusNoContent)
}

// matchRides は待っているライドを、待たせている順に、完了までの時間（迎えの距離 + 乗車距離）÷ 速度 が最小の空き椅子へ割り当てる。
// 1 回の呼び出しで割り当てられるだけ割り当てる
func matchRides() {
	store.mu.Lock()
	defer store.mu.Unlock()

	if len(store.unmatched) == 0 {
		return
	}

	// 空いている椅子（位置が分かるもの）
	free := make([]*Chair, 0, len(store.chairList))
	for _, c := range store.chairList {
		if c.IsActive && store.chairLoc[c.ID] != nil && store.chairIsFree(c.ID) {
			free = append(free, c)
		}
	}
	if len(free) == 0 {
		return
	}

	rest := store.unmatched[:0:0]
	for _, ride := range store.unmatched {
		rideDist := calculateDistance(ride.PickupLatitude, ride.PickupLongitude, ride.DestinationLatitude, ride.DestinationLongitude)
		best := -1
		bestCost := 0.0
		for i, c := range free {
			if c == nil {
				continue
			}
			loc := store.chairLoc[c.ID]
			pickupDist := calculateDistance(loc.Latitude, loc.Longitude, ride.PickupLatitude, ride.PickupLongitude)
			// 乗車地点までの時間 + 目的地までの時間（速い椅子は遠くても早く着く）
			cost := float64(pickupDist+rideDist) / float64(chairSpeed(c.Model))
			if best < 0 || cost < bestCost {
				best, bestCost = i, cost
			}
		}
		if best < 0 {
			rest = append(rest, ride)
			continue
		}
		matched := free[best]
		free[best] = nil
		ride.ChairID = sql.NullString{String: matched.ID, Valid: true}
		ride.UpdatedAt = now()
		store.ridesByChair[matched.ID] = append(store.ridesByChair[matched.ID], ride)
		dirty.Store(true)
	}
	store.unmatched = rest
}
