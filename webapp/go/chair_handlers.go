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

	store.mu.Lock()
	owner := store.ownersByRegister[req.ChairRegisterToken]
	if owner == nil {
		store.mu.Unlock()
		writeError(w, http.StatusUnauthorized, errors.New("invalid chair_register_token"))
		return
	}
	t := now()
	store.addChair(&Chair{
		ID: chairID, OwnerID: owner.ID, Name: req.Name, Model: req.Model,
		IsActive: false, AccessToken: accessToken, CreatedAt: t, UpdatedAt: t,
	})
	dirty.Store(true)
	store.mu.Unlock()

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

	store.mu.Lock()
	if chair.IsActive != req.IsActive {
		chair.IsActive = req.IsActive
		chair.UpdatedAt = now()
		dirty.Store(true)
	}
	store.mu.Unlock()

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

	store.mu.Lock()
	recordedAt := now()
	store.addLocation(chair.ID, req.Latitude, req.Longitude, recordedAt)

	if ride := store.latestRideOfChair(chair.ID); ride != nil {
		status, _ := store.latestStatus(ride.ID)
		if status != "COMPLETED" && status != "CANCELED" {
			if req.Latitude == ride.PickupLatitude && req.Longitude == ride.PickupLongitude && status == "ENROUTE" {
				store.addStatus(&RideStatus{ID: ulid.Make().String(), RideID: ride.ID, Status: "PICKUP", CreatedAt: now()})
			}

			if req.Latitude == ride.DestinationLatitude && req.Longitude == ride.DestinationLongitude && status == "CARRYING" {
				store.addStatus(&RideStatus{ID: ulid.Make().String(), RideID: ride.ID, Status: "ARRIVED", CreatedAt: now()})
			}
		}
	}
	dirty.Store(true)
	store.mu.Unlock()

	writeJSON(w, http.StatusOK, &chairPostCoordinateResponse{
		RecordedAt: recordedAt.UnixMilli(),
	})
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

	store.mu.Lock()
	defer store.mu.Unlock()

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
