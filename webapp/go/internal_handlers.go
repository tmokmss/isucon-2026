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

// matchRides は待っているライドを、待たせている順に、配車位置に最も近い空き椅子へ割り当てる。
// 1 回の呼び出しで割り当てられるだけ割り当てる
func matchRides() {
	unlock := lockTables([]table{tRides}, []table{tChairs, tChairLocs, tRideStatuses})
	defer unlock()

	if len(store.unmatched) == 0 {
		return
	}

	// 空いている椅子（位置が分かるもの）と、その位置
	free := make([]*Chair, 0, len(store.chairList))
	locs := make([]chairLocSnap, 0, len(store.chairList))
	for _, c := range store.chairList {
		if !c.IsActive || !store.chairIsFree(c.ID) {
			continue
		}
		agg := store.chairLoc[c.ID]
		if agg == nil {
			continue
		}
		if l, ok := agg.get(); ok {
			free = append(free, c)
			locs = append(locs, l)
		}
	}
	if len(free) == 0 {
		return
	}

	rest := store.unmatched[:0:0]
	for _, ride := range store.unmatched {
		best := -1
		bestDist := 0
		for i, c := range free {
			if c == nil {
				continue
			}
			loc := locs[i]
			d := calculateDistance(loc.Latitude, loc.Longitude, ride.PickupLatitude, ride.PickupLongitude)
			if best < 0 || d < bestDist {
				best, bestDist = i, d
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
		store.wakeRide(ride) // 椅子に MATCHING を送る
		dirty.Store(true)
	}
	store.unmatched = rest
}
