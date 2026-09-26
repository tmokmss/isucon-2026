package main

import (
	"database/sql"
	"net/http"
	"slices"
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

// matchRides は空き椅子と待っているライドの全組み合わせを、乗車地点までの時間（迎えの距離 ÷ 速度）の短い順に割り当てる。
// 椅子が足りずライドが余っている間は、迎えの移動（得点は距離の 0.1 倍）を減らすほど椅子が多くのライドを回せる。
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

	type pair struct {
		ride, chair int
		cost        int // 乗車地点までの tick 数 × 速度の最小公倍数に合わせた整数（比較にだけ使う）
	}
	pairs := make([]pair, 0, len(store.unmatched)*len(free))
	for ri, ride := range store.unmatched {
		for ci, c := range free {
			d := calculateDistance(locs[ci].Latitude, locs[ci].Longitude, ride.PickupLatitude, ride.PickupLongitude)
			// 速度は 2〜7。420（2〜7 の最小公倍数）を掛けて整数のまま割る
			pairs = append(pairs, pair{ri, ci, d * 420 / chairSpeed(c.Model)})
		}
	}
	slices.SortFunc(pairs, func(a, b pair) int {
		if a.cost != b.cost {
			return a.cost - b.cost
		}
		return a.ride - b.ride // 同じなら待たせている方を先に
	})

	rideDone := make([]bool, len(store.unmatched))
	chairDone := make([]bool, len(free))
	left := min(len(store.unmatched), len(free))
	for _, p := range pairs {
		if left == 0 {
			break
		}
		if rideDone[p.ride] || chairDone[p.chair] {
			continue
		}
		rideDone[p.ride], chairDone[p.chair] = true, true
		left--
		ride, matched := store.unmatched[p.ride], free[p.chair]
		ride.ChairID = sql.NullString{String: matched.ID, Valid: true}
		ride.UpdatedAt = now()
		store.ridesByChair[matched.ID] = append(store.ridesByChair[matched.ID], ride)
		store.wakeRide(ride) // 椅子に MATCHING を送る
		dirty.Store(true)
	}

	rest := store.unmatched[:0:0]
	for i, ride := range store.unmatched {
		if !rideDone[i] {
			rest = append(rest, ride)
		}
	}
	store.unmatched = rest
}
