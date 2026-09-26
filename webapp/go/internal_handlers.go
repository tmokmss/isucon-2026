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

// マッチングで椅子ごとに見るライドの候補数
const matchCandidates = 8

// startMatcher はアプリの中でマッチングを定期的に回す
func startMatcher() {
	go func() {
		t := time.NewTicker(matchingInterval)
		for range t.C {
			matchRides()
		}
	}()
}

// マッチングはアプリの中で回す（startMatcher）ので、isuride-matcher からの呼び出しでは何もしない。
// AMI の isuride-matcher が各台で 0.5 秒ごとに叩き、余分にマッチングを回していた（ISURIDE.md: 内部用なので仕様は自由）
func internalGetMatching(w http.ResponseWriter, r *http.Request) {
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
//   - ベンチが椅子にマッチを受理させてから 3 秒以内の椅子。受理は ENROUTE の送信なので、こちらが ENROUTE を受け取ってから
//     1.5 秒以内なら、通信の遅れを見込んでも収まる。受理の前は出さない（ベンチ側で前のライドの評価が済んでいるか分からない。
//     割り当てた時刻で 2 秒としたら、評価のレスポンスが遅れた椅子で CODE 30 が 1 秒に 35 件出た）
var freeChairs atomic.Pointer[[]freeChair]

// 完了（COMPLETED）してからこの時間は、空いていても付近の椅子に出さない。
// ベンチは評価のレスポンスを受け取るまで椅子をライド中とみなし、負荷が高いとそれが遅れる（出すと CODE 30）
const nearbyCompletedHide = 500 * time.Millisecond

// recentlyCompleted は椅子の直近のライドが nearbyCompletedHide 以内に完了したか。tRides と tRideStatuses の読みロック以上で呼ぶ
func (s *Store) recentlyCompleted(chairID string, t time.Time) bool {
	r := s.latestRideOfChair(chairID)
	if r == nil {
		return false
	}
	sts := s.statuses[r.ID]
	if len(sts) == 0 {
		return false
	}
	last := sts[len(sts)-1]
	return last.Status == "COMPLETED" && t.Sub(last.CreatedAt) < nearbyCompletedHide
}

// 椅子が受理（ENROUTE）してからこの時間は、付近の椅子に出す（ベンチの許容は受理から 3 秒）
const nearbyAcceptedGrace = 1500 * time.Millisecond

// 割り当てた椅子を、受理を待って付近の椅子に出す候補として持っておく時間
const recentMatchedKeep = 5 * time.Second

type matchedChair struct {
	freeChair
	ride *Ride
	at   time.Time
}

// recentMatched は割り当ててから recentMatchedKeep 以内の椅子。matchRides だけが触る（tRides の書きロックの下）
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
	matchedRide := make([]*Ride, len(free))
	// 完了直後かは割り当ての前に見る（割り当てると直近のライドが新しいライドに変わり、判定をすり抜けて CODE 30 を出していた）
	t0 := now()
	hide := make([]bool, len(free))
	for ci, c := range free {
		hide[ci] = store.recentlyCompleted(c.ID, t0)
	}
	// 付近の椅子の API に渡す一覧: 割り当て前の空き椅子すべてと、割り当ててから nearbyMatchedGrace 以内の椅子
	defer func() {
		t := now()
		list := make([]freeChair, 0, len(free)+len(recentMatched))
		for ci, c := range free {
			if !hide[ci] {
				list = append(list, freeChair{c, aggs[ci]})
			}
			if chairDone[ci] {
				recentMatched = append(recentMatched, matchedChair{freeChair{c, aggs[ci]}, matchedRide[ci], t})
			}
		}
		keep := recentMatched[:0]
		for _, m := range recentMatched {
			if t.Sub(m.at) < recentMatchedKeep {
				keep = append(keep, m)
			}
		}
		recentMatched = keep
		for _, m := range recentMatched {
			// 今回割り当てた椅子は free に入っている。/api/initialize の後は前のストアの椅子を出さない
			if m.at.Equal(t) || store.chairs[m.chair.ID] != m.chair || store.rides[m.ride.ID] != m.ride {
				continue
			}
			// 椅子が受理（ENROUTE）してから nearbyAcceptedGrace 以内で、受理の後に座標を送ってきた椅子だけ出す。
			// ベンチは受理のレスポンスを受け取ってから新しいライドを履歴に足し、それまでは前のライドで判定する
			// （前のライドの評価が済んでいないと CODE 30）。椅子は tick を順に処理するので、受理の後の座標は履歴を足した後に来る
			loc, ok := m.loc.get()
			if !ok {
				continue
			}
			for _, st := range store.statuses[m.ride.ID] {
				if st.Status == "ENROUTE" && t.Sub(st.CreatedAt) < nearbyAcceptedGrace && loc.UpdatedAt.After(st.CreatedAt) {
					list = append(list, m.freeChair)
					break
				}
			}
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
		matchedRide[ci] = ride
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
	// 椅子ごとに、乗車地点までの時間が短いライドを matchCandidates 件だけ候補にする（全組み合わせを並べ替えると
	// CPU プロファイルで 4.3%）。候補がすべて他の椅子に取られた椅子は、次の回に回す
	pairs := make([]pair, 0, len(free)*matchCandidates)
	var top [matchCandidates]pair
	for ci := range free {
		if chairDone[ci] {
			continue
		}
		n := 0
		for ri := range store.unmatched {
			if rideDone[ri] {
				continue
			}
			p := pair{ri, ci, cost(ri, ci)}
			if n == matchCandidates && p.cost >= top[n-1].cost {
				continue
			}
			// 挿入ソート（cost の昇順。同じなら先に見た = 待たせている方が前）
			j := min(n, matchCandidates-1)
			for j > 0 && top[j-1].cost > p.cost {
				top[j] = top[j-1]
				j--
			}
			top[j] = p
			if n < matchCandidates {
				n++
			}
		}
		pairs = append(pairs, top[:n]...)
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
