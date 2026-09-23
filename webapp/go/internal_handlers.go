package main

import (
	"database/sql"
	"math/rand/v2"
	"net/http"
)

// このAPIをインスタンス内から一定間隔で叩かせることで、椅子とライドをマッチングさせる
func internalGetMatching(w http.ResponseWriter, r *http.Request) {
	store.mu.Lock()
	defer store.mu.Unlock()

	// 最も待たせているリクエストに、空いている椅子をランダムに割り当てる
	if len(store.unmatched) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	ride := store.unmatched[0]

	var candidates []*Chair
	for _, c := range store.chairList {
		if c.IsActive && store.chairIsFree(c.ID) {
			candidates = append(candidates, c)
		}
	}
	if len(candidates) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	matched := candidates[rand.IntN(len(candidates))]

	store.unmatched = store.unmatched[1:]
	ride.ChairID = sql.NullString{String: matched.ID, Valid: true}
	ride.UpdatedAt = now()
	store.ridesByChair[matched.ID] = append(store.ridesByChair[matched.ID], ride)
	dirty.Store(true)

	w.WriteHeader(http.StatusNoContent)
}
