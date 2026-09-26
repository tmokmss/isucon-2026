package main

import (
	"database/sql"
	"net/http"
	"slices"
	"sync/atomic"
	"time"
)

// マッチングの間隔。isuride-matcher（0.5 秒ごと）を待たずにアプリの中で回す
const matchingInterval = 100 * time.Millisecond

// これより長く待たせているライドは、乗車地点までの時間によらず先に割り当てる（ベンチは 30 秒で CODE 32）
const starvationLimit = 10 * time.Second

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

// freeChair は空いている椅子と、その位置（行ロックで最新を読む）
type freeChair struct {
	chair *Chair
	loc   *chairLocAgg
}

// freeChairs はマッチングのたびに作り直す、付近の椅子の API に出す椅子の一覧。
// 付近の椅子の API は、全椅子を見る代わりにこれを使う。
// ユーザーは付近に椅子が無いと配車をやめるので、出せる椅子はできるだけ出す（割り当て後に残った空き椅子だけにしたら、
// 依頼が減ってスコアが 61 万 → 25 万に落ちた）。ベンチが許すのは次の椅子:
//   - 空いている椅子（空いてから 3 秒以内なら、含まれていなくてもよい）
//   - ベンチがマッチを受理してから 3 秒以内の椅子。受理はこちらの割り当てより後なので、割り当てから 2 秒以内なら確実に収まる
var freeChairs atomic.Pointer[[]freeChair]

// 割り当ててからこの時間は、付近の椅子に出す（ベンチの許容は受理から 3 秒）
const nearbyMatchedGrace = 2 * time.Second

type matchedChair struct {
	freeChair
	at time.Time
}

// recentMatched は割り当ててから nearbyMatchedGrace 以内の椅子。matchRides だけが触る（tRides の書きロックの下）
var recentMatched []matchedChair

// matchRides は空き椅子と待っているライドの全組み合わせを、乗車地点までの時間（迎えの距離 ÷ 速度）の短い順に割り当てる。
// 椅子が足りずライドが余っている間は、迎えの移動（得点は距離の 0.1 倍）を減らすほど椅子が多くのライドを回せる。
// 1 回の呼び出しで割り当てられるだけ割り当てる
func matchRides() {
	unlock := lockTables([]table{tRides}, []table{tChairs, tChairLocs, tRideStatuses})
	defer unlock()

	// 空いている椅子（位置が分かるもの）と、その位置
	free := make([]*Chair, 0, len(store.chairList))
	aggs := make([]*chairLocAgg, 0, len(store.chairList))
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
			aggs = append(aggs, agg)
			locs = append(locs, l)
		}
	}
	chairDone := make([]bool, len(free))
	// 付近の椅子の API に渡す一覧: 割り当て前の空き椅子すべてと、割り当ててから nearbyMatchedGrace 以内の椅子
	defer func() {
		t := now()
		list := make([]freeChair, 0, len(free)+len(recentMatched))
		for ci, c := range free {
			list = append(list, freeChair{c, aggs[ci]})
			if chairDone[ci] {
				recentMatched = append(recentMatched, matchedChair{freeChair{c, aggs[ci]}, t})
			}
		}
		keep := recentMatched[:0]
		for _, m := range recentMatched {
			if t.Sub(m.at) < nearbyMatchedGrace {
				keep = append(keep, m)
			}
		}
		recentMatched = keep
		for _, m := range recentMatched {
			// 今回割り当てた椅子は free に入っている。/api/initialize の後は前のストアの椅子を出さない
			if m.at.Equal(t) || store.chairs[m.chair.ID] != m.chair || store.chairIsFree(m.chair.ID) {
				continue
			}
			list = append(list, m.freeChair)
		}
		freeChairs.Store(&list)
	}()
	if len(store.unmatched) == 0 || len(free) == 0 {
		return
	}

	rideDone := make([]bool, len(store.unmatched))
	assign := func(ri, ci int) {
		rideDone[ri], chairDone[ci] = true, true
		ride, matched := store.unmatched[ri], free[ci]
		ride.ChairID = sql.NullString{String: matched.ID, Valid: true}
		ride.UpdatedAt = now()
		store.ridesByChair[matched.ID] = append(store.ridesByChair[matched.ID], ride)
		store.wakeRide(ride) // 椅子に MATCHING を送る
		dirty.Store(true)
	}
	cost := func(ri, ci int) int {
		ride := store.unmatched[ri]
		d := calculateDistance(locs[ci].Latitude, locs[ci].Longitude, ride.PickupLatitude, ride.PickupLongitude)
		// 速度は 2〜7。420（2〜7 の最小公倍数）を掛けて整数のまま割る
		return d * 420 / chairSpeed(free[ci].Model)
	}

	// 待たせすぎているライドは、待たせている順に先に割り当てる。
	// ベンチは 30 秒マッチされないとクリティカルエラーにする（CODE 32）。貪欲法だけだと、椅子の少ない地域のライドが後回しになり続ける
	starveBefore := now().Add(-starvationLimit)
	for ri, ride := range store.unmatched {
		if !ride.CreatedAt.Before(starveBefore) {
			break // unmatched は待たせている順
		}
		best, bestCost := -1, 0
		for ci := range free {
			if chairDone[ci] {
				continue
			}
			if c := cost(ri, ci); best < 0 || c < bestCost {
				best, bestCost = ci, c
			}
		}
		if best < 0 {
			break
		}
		assign(ri, best)
	}

	type pair struct {
		ride, chair int
		cost        int // 乗車地点までの tick 数 × 速度の最小公倍数に合わせた整数（比較にだけ使う）
	}
	pairs := make([]pair, 0, len(store.unmatched)*len(free))
	for ri := range store.unmatched {
		if rideDone[ri] {
			continue
		}
		for ci := range free {
			if !chairDone[ci] {
				pairs = append(pairs, pair{ri, ci, cost(ri, ci)})
			}
		}
	}
	slices.SortFunc(pairs, func(a, b pair) int {
		if a.cost != b.cost {
			return a.cost - b.cost
		}
		return a.ride - b.ride // 同じなら待たせている方を先に
	})
	for _, p := range pairs {
		if rideDone[p.ride] || chairDone[p.chair] {
			continue
		}
		assign(p.ride, p.chair)
	}

	rest := store.unmatched[:0:0]
	for i, ride := range store.unmatched {
		if !rideDone[i] {
			rest = append(rest, ride)
		}
	}
	store.unmatched = rest
}
