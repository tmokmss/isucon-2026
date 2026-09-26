package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
)

const (
	initialFare     = 500
	farePerDistance = 100
)

type ownerPostOwnersRequest struct {
	Name string `json:"name"`
}

type ownerPostOwnersResponse struct {
	ID                 string `json:"id"`
	ChairRegisterToken string `json:"chair_register_token"`
}

func ownerPostOwners(w http.ResponseWriter, r *http.Request) {
	req := &ownerPostOwnersRequest{}
	if err := bindJSON(r, req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusBadRequest, errors.New("some of required fields(name) are empty"))
		return
	}

	ownerID := ulid.Make().String()
	accessToken := secureRandomStr(32)
	chairRegisterToken := secureRandomStr(32)

	unlock := lockTables([]table{tOwners}, nil)
	if _, dup := store.ownerNames[req.Name]; dup {
		unlock()
		writeError(w, http.StatusInternalServerError, errors.New("Error 1062 (23000): Duplicate entry for key 'owners.name'"))
		return
	}
	t := now()
	store.addOwner(&Owner{ID: ownerID, Name: req.Name, AccessToken: accessToken, ChairRegisterToken: chairRegisterToken, CreatedAt: t, UpdatedAt: t})
	dirty.Store(true)
	unlock()

	http.SetCookie(w, &http.Cookie{
		Path:  "/",
		Name:  "owner_session",
		Value: accessToken,
	})

	writeJSON(w, http.StatusCreated, &ownerPostOwnersResponse{
		ID:                 ownerID,
		ChairRegisterToken: chairRegisterToken,
	})
}

type chairSales struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Sales int    `json:"sales"`
}

type modelSales struct {
	Model string `json:"model"`
	Sales int    `json:"sales"`
}

type ownerGetSalesResponse struct {
	TotalSales int          `json:"total_sales"`
	Chairs     []chairSales `json:"chairs"`
	Models     []modelSales `json:"models"`
}

func ownerGetSales(w http.ResponseWriter, r *http.Request) {
	since := time.Unix(0, 0)
	until := time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
	if r.URL.Query().Get("since") != "" {
		parsed, err := strconv.ParseInt(r.URL.Query().Get("since"), 10, 64)
		if err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}
		since = time.UnixMilli(parsed)
	}
	if r.URL.Query().Get("until") != "" {
		parsed, err := strconv.ParseInt(r.URL.Query().Get("until"), 10, 64)
		if err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}
		until = time.UnixMilli(parsed)
	}

	owner := r.Context().Value("owner").(*Owner)

	untilInclusive := until.Add(999 * time.Microsecond)

	res := ownerGetSalesResponse{
		TotalSales: 0,
	}

	modelSalesByModel := map[string]int{}
	unlock := lockTables(nil, []table{tChairs, tRides, tRideStatuses})
	for _, chair := range store.chairsByOwner[owner.ID] {
		sales := 0
		for _, ride := range store.ridesByChair[chair.ID] {
			if ride.UpdatedAt.Before(since) || ride.UpdatedAt.After(untilInclusive) {
				continue
			}
			if store.rideCompleted(ride) {
				sales += calculateSale(*ride)
			}
		}
		res.TotalSales += sales

		res.Chairs = append(res.Chairs, chairSales{
			ID:    chair.ID,
			Name:  chair.Name,
			Sales: sales,
		})

		modelSalesByModel[chair.Model] += sales
	}
	unlock()

	models := []modelSales{}
	for model, sales := range modelSalesByModel {
		models = append(models, modelSales{
			Model: model,
			Sales: sales,
		})
	}
	res.Models = models

	writeJSON(w, http.StatusOK, res)
}

func sumSales(rides []Ride) int {
	sale := 0
	for _, ride := range rides {
		sale += calculateSale(ride)
	}
	return sale
}

func calculateSale(ride Ride) int {
	return calculateFare(ride.PickupLatitude, ride.PickupLongitude, ride.DestinationLatitude, ride.DestinationLongitude)
}

type chairWithDetail struct {
	ID                     string       `db:"id"`
	OwnerID                string       `db:"owner_id"`
	Name                   string       `db:"name"`
	AccessToken            string       `db:"access_token"`
	Model                  string       `db:"model"`
	IsActive               bool         `db:"is_active"`
	CreatedAt              time.Time    `db:"created_at"`
	UpdatedAt              time.Time    `db:"updated_at"`
	TotalDistance          int          `db:"total_distance"`
	TotalDistanceUpdatedAt sql.NullTime `db:"total_distance_updated_at"`
}

type ownerGetChairResponse struct {
	Chairs []ownerGetChairResponseChair `json:"chairs"`
}

type ownerGetChairResponseChair struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	Model                  string `json:"model"`
	Active                 bool   `json:"active"`
	RegisteredAt           int64  `json:"registered_at"`
	TotalDistance          int    `json:"total_distance"`
	TotalDistanceUpdatedAt *int64 `json:"total_distance_updated_at,omitempty"`
}

func ownerGetChairs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	owner := ctx.Value("owner").(*Owner)

	unlock := lockTables(nil, []table{tChairs, tChairLocs})
	chairs := make([]chairWithDetail, 0, len(store.chairsByOwner[owner.ID]))
	for _, c := range store.chairsByOwner[owner.ID] {
		d := chairWithDetail{
			ID: c.ID, OwnerID: c.OwnerID, Name: c.Name, AccessToken: c.AccessToken, Model: c.Model,
			IsActive: c.IsActive, CreatedAt: c.CreatedAt, UpdatedAt: c.UpdatedAt,
		}
		if agg := store.chairLoc[c.ID]; agg != nil {
			if l, ok := agg.get(); ok {
				d.TotalDistance = l.TotalDistance
				d.TotalDistanceUpdatedAt = sql.NullTime{Time: l.UpdatedAt, Valid: true}
			}
		}
		chairs = append(chairs, d)
	}
	unlock()

	res := ownerGetChairResponse{}
	for _, chair := range chairs {
		c := ownerGetChairResponseChair{
			ID:            chair.ID,
			Name:          chair.Name,
			Model:         chair.Model,
			Active:        chair.IsActive,
			RegisteredAt:  chair.CreatedAt.UnixMilli(),
			TotalDistance: chair.TotalDistance,
		}
		if chair.TotalDistanceUpdatedAt.Valid {
			t := chair.TotalDistanceUpdatedAt.Time.UnixMilli()
			c.TotalDistanceUpdatedAt = &t
		}
		res.Chairs = append(res.Chairs, c)
	}
	writeJSON(w, http.StatusOK, res)
}
