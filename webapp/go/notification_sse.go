package main

import (
	"fmt"
	json "github.com/goccy/go-json"
	"net/http"
	"sync"
	"time"
)

// 通知は SSE（text/event-stream）で送る（ISURIDE.md「通知エンドポイント」で認められている）。
// 接続したら即座に最新の状態を送り、以降は状態が変わったときだけ送る。
// 未送信の状態のうち最も古いものから 1 件ずつ送って送信済みにするので、
// すべての状態遷移を発生した順に少なくとも 1 回送る（ポーリング時代と同じ判定）。

// 取りこぼしに備えた見直しの間隔（仕様は状態が変わってから 3 秒以内）
const notificationRecheck = 500 * time.Millisecond

// 状態が変わったことを、待っている SSE の接続に知らせる。ストアの中身とは独立に持つ（/initialize でも消えない）
var (
	wakeMu    sync.Mutex
	userWake  = map[string]chan struct{}{}
	chairWake = map[string]chan struct{}{}
)

func wakeChan(m map[string]chan struct{}, id string) chan struct{} {
	wakeMu.Lock()
	defer wakeMu.Unlock()
	ch := m[id]
	if ch == nil {
		ch = make(chan struct{}, 1)
		m[id] = ch
	}
	return ch
}

func wake(m map[string]chan struct{}, id string) {
	select {
	case wakeChan(m, id) <- struct{}{}:
	default:
	}
}

// wakeRide はライドの状態や割り当てが変わったときに、そのユーザーと椅子の接続を起こす。tRides の読みロック以上で呼ぶ
func (s *Store) wakeRide(ride *Ride) {
	wake(userWake, ride.UserID)
	if ride.ChairID.Valid {
		wake(chairWake, ride.ChairID.String)
	}
}

// appNextNotification は次に送るユーザー向けの通知を作る。未送信の状態があれば送信済みにして sent=true を返す。
// ライドが無ければ nil
func appNextNotification(user *User) (data *appGetNotificationResponseData, sent bool) {
	// rideStatuses は読みロック。送信済みの印はライドの行ロックで書く（別のライドの通知と並列に走る）
	unlock := lockTables(nil, []table{tChairs, tRides, tRideStatuses, tCoupons, tChairStats})
	defer unlock()

	rides := store.ridesByUser[user.ID]
	if len(rides) == 0 {
		return nil, false
	}
	ride := rides[len(rides)-1]

	row := store.statusRow[ride.ID]
	if row == nil {
		return nil, false
	}
	row.Lock()
	defer row.Unlock()
	var yetSent *RideStatus
	for _, st := range store.statuses[ride.ID] {
		if st.AppSentAt == nil {
			yetSent = st
			break
		}
	}
	status := ""
	if yetSent != nil {
		status = yetSent.Status
	} else {
		status, _ = store.latestStatus(ride.ID)
	}

	fare := store.discountedFare(user.ID, ride, ride.PickupLatitude, ride.PickupLongitude, ride.DestinationLatitude, ride.DestinationLongitude)
	data = &appGetNotificationResponseData{
		RideID:                ride.ID,
		PickupCoordinate:      Coordinate{Latitude: ride.PickupLatitude, Longitude: ride.PickupLongitude},
		DestinationCoordinate: Coordinate{Latitude: ride.DestinationLatitude, Longitude: ride.DestinationLongitude},
		Fare:                  fare,
		Status:                status,
		CreatedAt:             ride.CreatedAt.UnixMilli(),
		UpdateAt:              ride.UpdatedAt.UnixMilli(),
	}
	if ride.ChairID.Valid {
		if chair := store.chairs[ride.ChairID.String]; chair != nil {
			data.Chair = &appGetNotificationResponseChair{
				ID:    chair.ID,
				Name:  chair.Name,
				Model: chair.Model,
				Stats: store.chairStatsOf(chair.ID),
			}
		}
	}

	if yetSent != nil {
		t := now()
		yetSent.AppSentAt = &t
		dirty.Store(true)
		return data, true
	}
	return data, false
}

// chairNextNotification は次に送る椅子向けの通知を作る。未送信の状態があれば送信済みにして sent=true を返す。
// ライドが無ければ nil
func chairNextNotification(chair *Chair) (data *chairGetNotificationResponseData, sent bool) {
	unlock := lockTables(nil, []table{tUsers, tRides, tRideStatuses})
	defer unlock()

	ride := store.latestRideOfChair(chair.ID)
	if ride == nil {
		return nil, false
	}
	row := store.statusRow[ride.ID]
	if row == nil {
		return nil, false
	}
	row.Lock()
	defer row.Unlock()
	var yetSent *RideStatus
	for _, st := range store.statuses[ride.ID] {
		if st.ChairSentAt == nil {
			yetSent = st
			break
		}
	}
	status := ""
	if yetSent != nil {
		status = yetSent.Status
	} else {
		status, _ = store.latestStatus(ride.ID)
	}
	user := store.users[ride.UserID]
	if user == nil {
		return nil, false
	}

	data = &chairGetNotificationResponseData{
		RideID:                ride.ID,
		User:                  simpleUser{ID: user.ID, Name: fmt.Sprintf("%s %s", user.Firstname, user.Lastname)},
		PickupCoordinate:      Coordinate{Latitude: ride.PickupLatitude, Longitude: ride.PickupLongitude},
		DestinationCoordinate: Coordinate{Latitude: ride.DestinationLatitude, Longitude: ride.DestinationLongitude},
		Status:                status,
	}
	if yetSent != nil {
		t := now()
		yetSent.ChairSentAt = &t
		dirty.Store(true)
		return data, true
	}
	return data, false
}

// serveSSE は next を呼んで得た通知を SSE で送り続ける。
// 接続直後は最新の状態を必ず送り、以降は未送信の状態があるときだけ送る
func serveSSE[T any](w http.ResponseWriter, r *http.Request, wakeCh chan struct{}, next func() (*T, bool)) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("streaming unsupported"))
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("X-Accel-Buffering", "no") // nginx にバッファさせない
	w.WriteHeader(http.StatusOK)
	flusher.Flush()

	send := func(data *T) bool {
		b, err := json.Marshal(data)
		if err != nil {
			return false
		}
		if _, err := fmt.Fprintf(w, "data: %s\n\n", b); err != nil {
			return false
		}
		flusher.Flush()
		return true
	}

	first := true
	ticker := time.NewTicker(notificationRecheck)
	defer ticker.Stop()
	for {
		// 未送信が無くなるまで続けて送る
		for {
			data, sent := next()
			if data == nil {
				break
			}
			if sent || first {
				if !send(data) {
					return
				}
				first = false
			}
			if !sent {
				break
			}
		}
		select {
		case <-r.Context().Done():
			return
		case <-wakeCh:
		case <-ticker.C:
		}
	}
}

func appGetNotification(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*User)
	serveSSE(w, r, wakeChan(userWake, user.ID), func() (*appGetNotificationResponseData, bool) {
		return appNextNotification(user)
	})
}

func chairGetNotification(w http.ResponseWriter, r *http.Request) {
	chair := r.Context().Value("chair").(*Chair)
	serveSSE(w, r, wakeChan(chairWake, chair.ID), func() (*chairGetNotificationResponseData, bool) {
		return chairNextNotification(chair)
	})
}
