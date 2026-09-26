package main

import (
	"errors"
	"net/http"

	"github.com/oklog/ulid/v2"
)

type chairPostChairsRequest struct {
	Name               string `json:"name"`
	Model              string `json:"model"`
	ChairRegisterToken string `json:"chair_register_token"`
}

type chairPostChairsResponse struct {
	ID      string `json:"id"`
	OwnerID string `json:"owner_id"`
}

func chairPostChairs(w http.ResponseWriter, r *http.Request) {
	req := &chairPostChairsRequest{}
	if err := bindJSON(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.Model == "" || req.ChairRegisterToken == "" {
		writeError(w, http.StatusBadRequest, errors.New("some of required fields(name, model, chair_register_token) are empty"))
		return
	}

	chairID := ulid.Make().String()
	accessToken := secureRandomStr(32)

	unlock := lockTables([]table{tChairs, tChairLocs}, []table{tOwners})
	owner := store.ownersByRegister[req.ChairRegisterToken]
	if owner == nil {
		unlock()
		writeError(w, http.StatusUnauthorized, errors.New("invalid chair_register_token"))
		return
	}
	t := now()
	store.addChair(&Chair{
		ID: chairID, OwnerID: owner.ID, Name: req.Name, Model: req.Model,
		IsActive: false, AccessToken: accessToken, CreatedAt: t, UpdatedAt: t,
	})
	dirty.Store(true)
	unlock()

	http.SetCookie(w, &http.Cookie{
		Path:  "/",
		Name:  "chair_session",
		Value: accessToken,
	})

	writeJSON(w, http.StatusCreated, &chairPostChairsResponse{
		ID:      chairID,
		OwnerID: owner.ID,
	})
}

type postChairActivityRequest struct {
	IsActive bool `json:"is_active"`
}

func chairPostActivity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	chair := ctx.Value("chair").(*Chair)

	req := &postChairActivityRequest{}
	if err := bindJSON(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	unlock := lockTables([]table{tChairs}, nil)
	if chair.IsActive != req.IsActive {
		chair.IsActive = req.IsActive
		chair.UpdatedAt = now()
		dirty.Store(true)
	}
	unlock()

	w.WriteHeader(http.StatusNoContent)
}

type chairPostCoordinateResponse struct {
	RecordedAt int64 `json:"recorded_at"`
}

func chairPostCoordinate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &Coordinate{}
	if err := bindJSON(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chair := ctx.Value("chair").(*Chair)

	recordedAt := now()

	// 位置は chairLocs の読みロック + 行ロックで書く（別の椅子の座標の送信と並列に走る）。
	// ライドの状態は読むだけ。状態を足す必要があるときだけ、取り直して確かめ直す
	unlock := lockTables(nil, []table{tChairLocs, tRides, tRideStatuses})
	agg := store.chairLoc[chair.ID]
	if agg != nil {
		agg.add(req.Latitude, req.Longitude, recordedAt)
	}
	next := coordinateTransition(chair.ID, req)
	unlock()

	if agg == nil {
		// 位置の集約が無い（椅子の登録時に作るので通常は起きない）
		unlock := lockTables([]table{tChairLocs}, nil)
		if store.chairLoc[chair.ID] == nil {
			store.chairLoc[chair.ID] = &chairLocAgg{}
		}
		store.chairLoc[chair.ID].add(req.Latitude, req.Longitude, recordedAt)
		unlock()
	}
	if next != "" {
		unlock := lockTables([]table{tRideStatuses}, []table{tRides})
		// 取り直す間に状態が変わっているかもしれないので、確かめ直してから足す
		if coordinateTransition(chair.ID, req) == next {
			ride := store.latestRideOfChair(chair.ID)
			store.addStatus(&RideStatus{ID: ulid.Make().String(), RideID: ride.ID, Status: next, CreatedAt: now()})
		}
		unlock()
	}
	dirty.Store(true)

	writeJSON(w, http.StatusOK, &chairPostCoordinateResponse{
		RecordedAt: recordedAt.UnixMilli(),
	})
}

// coordinateTransition は、椅子がこの座標に着いたことで足すべき状態を返す（無ければ ""）。
// tRides と tRideStatuses の読みロック以上で呼ぶ
func coordinateTransition(chairID string, c *Coordinate) string {
	ride := store.latestRideOfChair(chairID)
	if ride == nil {
		return ""
	}
	status, _ := store.latestStatus(ride.ID)
	if c.Latitude == ride.PickupLatitude && c.Longitude == ride.PickupLongitude && status == "ENROUTE" {
		return "PICKUP"
	}
	if c.Latitude == ride.DestinationLatitude && c.Longitude == ride.DestinationLongitude && status == "CARRYING" {
		return "ARRIVED"
	}
	return ""
}

type simpleUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type chairGetNotificationResponse struct {
	Data         *chairGetNotificationResponseData `json:"data"`
	RetryAfterMs int                               `json:"retry_after_ms"`
}

type chairGetNotificationResponseData struct {
	RideID                string     `json:"ride_id"`
	User                  simpleUser `json:"user"`
	PickupCoordinate      Coordinate `json:"pickup_coordinate"`
	DestinationCoordinate Coordinate `json:"destination_coordinate"`
	Status                string     `json:"status"`
}

type postChairRidesRideIDStatusRequest struct {
	Status string `json:"status"`
}

func chairPostRideStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rideID := r.PathValue("ride_id")

	chair := ctx.Value("chair").(*Chair)

	req := &postChairRidesRideIDStatusRequest{}
	if err := bindJSON(r, req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	unlock := lockTables([]table{tRideStatuses}, []table{tRides})
	defer unlock()

	ride := store.rides[rideID]
	if ride == nil {
		writeError(w, http.StatusNotFound, errors.New("ride not found"))
		return
	}

	if ride.ChairID.String != chair.ID {
		writeError(w, http.StatusBadRequest, errors.New("not assigned to this ride"))
		return
	}

	switch req.Status {
	// Acknowledge the ride
	case "ENROUTE":
		store.addStatus(&RideStatus{ID: ulid.Make().String(), RideID: ride.ID, Status: "ENROUTE", CreatedAt: now()})
	// After Picking up user
	case "CARRYING":
		status, _ := store.latestStatus(ride.ID)
		if status != "PICKUP" {
			writeError(w, http.StatusBadRequest, errors.New("chair has not arrived yet"))
			return
		}
		store.addStatus(&RideStatus{ID: ulid.Make().String(), RideID: ride.ID, Status: "CARRYING", CreatedAt: now()})
	default:
		writeError(w, http.StatusBadRequest, errors.New("invalid status"))
		return
	}
	dirty.Store(true)

	w.WriteHeader(http.StatusNoContent)
}
