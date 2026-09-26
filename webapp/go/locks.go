package main

import "sync"

// ロックはテーブルごとに 1 つ（スキル isucon-improve の references/inmemory.md「ストアとロックの単位」）。
// 呼び出し側は Lock を直接呼ばず、必ず lockTables を通す。lockTables は順位順に取るので、待ちが循環しない。
// 守る決まり:
//   - 必要なロックは処理の最初にまとめて取る。途中で追加しない（要るなら全部放して取り直し、前提を確かめ直す）
//   - 読みから書きへ格上げしない
//   - 再入しない（ロックを持った状態で呼ぶ関数は、自分では取らない）
//   - 行のロック（chairLocAgg.mu）は、テーブルのロックの後に取る

type table int

const (
	tUsers table = iota // 順位 = 取る順
	tOwners
	tChairs
	tChairLocs
	tPaymentTokens
	tRides
	tRideStatuses
	tCoupons
	tChairStats
	tSettings
	nTables
)

var locks [nTables]sync.RWMutex

// lockTables は書くテーブルと読むテーブルのロックを順位順に取り、放す関数を返す
func lockTables(write, read []table) (unlock func()) {
	var mode [nTables]int // 0: 取らない, 1: 読み, 2: 書き
	for _, t := range read {
		mode[t] = max(mode[t], 1)
	}
	for _, t := range write {
		mode[t] = 2
	}
	for t := range nTables {
		switch mode[t] {
		case 1:
			locks[t].RLock()
		case 2:
			locks[t].Lock()
		}
	}
	return func() {
		for t := nTables - 1; t >= 0; t-- {
			switch mode[t] {
			case 1:
				locks[t].RUnlock()
			case 2:
				locks[t].Unlock()
			}
		}
	}
}

var allTables = func() []table {
	ts := make([]table, 0, nTables)
	for t := range nTables {
		ts = append(ts, t)
	}
	return ts
}()
