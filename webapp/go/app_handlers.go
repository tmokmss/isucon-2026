package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
)

type appPostUsersRequest struct {
	Username       string  `json:"username"`
	FirstName      string  `json:"firstname"`
	LastName       string  `json:"lastname"`
	DateOfBirth    string  `json:"date_of_birth"`
	InvitationCode *string `json:"invitation_code"`
}

type appPostUsersResponse struct {
	ID             string `json:"id"`
	InvitationCode string `json:"invitation_code"`
}

type appPostPaymentMethodsRequest struct {
	Token string `json:"token"`
}

type getAppRidesResponse struct {
	Rides []getAppRidesResponseItem `json:"rides"`
}

type getAppRidesResponseItem struct {
	ID                    string                       `json:"id"`
	PickupCoordinate      Coordinate                   `json:"pickup_coordinate"`
	DestinationCoordinate Coordinate                   `json:"destination_coordinate"`
	Chair                 getAppRidesResponseItemChair `json:"chair"`
	Fare                  int                          `json:"fare"`
	Evaluation            int                          `json:"evaluation"`
	RequestedAt           int64                        `json:"requested_at"`
	CompletedAt           int64                        `json:"completed_at"`
}

type getAppRidesResponseItemChair struct {
	ID    string `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
	Model string `json:"model"`
}

type appPostRidesRequest struct {
	PickupCoordinate      *Coordinate `json:"pickup_coordinate"`
	DestinationCoordinate *Coordinate `json:"destination_coordinate"`
}

type appPostRidesResponse struct {
	RideID string `json:"ride_id"`
	Fare   int    `json:"fare"`
}

type appPostRidesEstimatedFareRequest struct {
	PickupCoordinate      *Coordinate `json:"pickup_coordinate"`
	DestinationCoordinate *Coordinate `json:"destination_coordinate"`
}

type appPostRidesEstimatedFareResponse struct {
	Fare     int `json:"fare"`
	Discount int `json:"discount"`
}

type appPostRideEvaluationRequest struct {
	Evaluation int `json:"evaluation"`
}

type appPostRideEvaluationResponse struct {
	CompletedAt int64 `json:"completed_at"`
}

type appGetNotificationResponse struct {
	Data         *appGetNotificationResponseData `json:"data"`
	RetryAfterMs int                             `json:"retry_after_ms"`
}

type appGetNotificationResponseData struct {
	RideID                string                           `json:"ride_id"`
	PickupCoordinate      Coordinate                       `json:"pickup_coordinate"`
	DestinationCoordinate Coordinate                       `json:"destination_coordinate"`
	Fare                  int                              `json:"fare"`
	Status                string                           `json:"status"`
	Chair                 *appGetNotificationResponseChair `json:"chair,omitempty"`
	CreatedAt             int64                            `json:"created_at"`
	UpdateAt              int64                            `json:"updated_at"`
}

type appGetNotificationResponseChair struct {
	ID    string                               `json:"id"`
	Name  string                               `json:"name"`
	Model string                               `json:"model"`
	Stats appGetNotificationResponseChairStats `json:"stats"`
}

type appGetNotificationResponseChairStats struct {
	TotalRidesCount    int     `json:"total_rides_count"`
	TotalEvaluationAvg float64 `json:"total_evaluation_avg"`
}

type appGetNearbyChairsResponse struct {
	Chairs      []appGetNearbyChairsResponseChair `json:"chairs"`
	RetrievedAt int64                             `json:"retrieved_at"`
}

type appGetNearbyChairsResponseChair struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	Model             string     `json:"model"`
	CurrentCoordinate Coordinate `json:"current_coordinate"`
}

func appPostUsers(w http.ResponseWriter, r *http.Request) {
	req := &appPostUsersRequest{}
	if err := bindJSON(r, req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if req.Username == "" || req.FirstName == "" || req.LastName == "" || req.DateOfBirth == "" {
		writeError(w, http.StatusBadRequest, errors.New("required fields(username, firstname, lastname, date_of_birth) are empty"))
		return
	}

	userID := ulid.Make().String()
	accessToken := secureRandomStr(32)
	invitationCode := secureRandomStr(15)

	unlock := lockTables([]table{tUsers, tCoupons}, nil)
	defer unlock()

	if _, dup := store.usernames[req.Username]; dup {
		writeError(w, http.StatusInternalServerError, errors.New("Error 1062 (23000): Duplicate entry for key 'users.username'"))
		return
	}

	// 招待コードを使った登録（検証を先に済ませ、通ったら書く）
	var inviter *User
	if req.InvitationCode != nil && *req.InvitationCode != "" {
		// 招待する側の招待数をチェック
		if store.couponCodeCount["INV_"+*req.InvitationCode] >= 3 {
			writeError(w, http.StatusBadRequest, errors.New("この招待コードは使用できません。"))
			return
		}
		// ユーザーチェック
		inviter = store.usersByInvitation[*req.InvitationCode]
		if inviter == nil {
			writeError(w, http.StatusBadRequest, errors.New("この招待コードは使用できません。"))
			return
		}
	}

	var rewardCode string
	if inviter != nil {
		rewardCode = fmt.Sprintf("RWD_%s_%d", *req.InvitationCode, time.Now().UnixMilli())
		if _, dup := store.couponKeys[couponKey(inviter.ID, rewardCode)]; dup {
			writeError(w, http.StatusInternalServerError, errors.New("Error 1062 (23000): Duplicate entry for key 'coupons.PRIMARY'"))
			return
		}
	}

	t := now()
	store.addUser(&User{
		ID: userID, Username: req.Username, Firstname: req.FirstName, Lastname: req.LastName,
		DateOfBirth: req.DateOfBirth, AccessToken: accessToken, InvitationCode: invitationCode,
		CreatedAt: t, UpdatedAt: t,
	})
	// 初回登録キャンペーンのクーポンを付与
	store.addCoupon(&Coupon{UserID: userID, Code: "CP_NEW2024", Discount: 3000, CreatedAt: t})
	if inviter != nil {
		// 招待クーポン付与
		store.addCoupon(&Coupon{UserID: userID, Code: "INV_" + *req.InvitationCode, Discount: 1500, CreatedAt: now()})
		// 招待した人にもRewardを付与
		store.addCoupon(&Coupon{UserID: inviter.ID, Code: rewardCode, Discount: 1000, CreatedAt: now()})
	}
	dirty.Store(true)

	http.SetCookie(w, &http.Cookie{
		Path:  "/",
		Name:  "app_session",
		Value: accessToken,
	})

	writeJSON(w, http.StatusCreated, &appPostUsersResponse{
		ID:             userID,
		InvitationCode: invitationCode,
	})
}

func appPostPaymentMethods(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &appPostPaymentMethodsRequest{}
	if err := bindJSON(r, req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if req.Token == "" {
		writeError(w, http.StatusBadRequest, errors.New("token is required but was empty"))
		return
	}

	user := ctx.Value("user").(*User)

	unlock := lockTables([]table{tPaymentTokens}, nil)
	if _, dup := store.paymentTokens[user.ID]; dup {
		unlock()
		writeError(w, http.StatusInternalServerError, errors.New("Error 1062 (23000): Duplicate entry for key 'payment_tokens.PRIMARY'"))
		return
	}
	store.paymentTokens[user.ID] = &PaymentToken{UserID: user.ID, Token: req.Token, CreatedAt: now()}
	dirty.Store(true)
	unlock()

	w.WriteHeader(http.StatusNoContent)
}

func appGetRides(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(*User)

	items := []getAppRidesResponseItem{}

	unlock := lockTables(nil, []table{tOwners, tChairs, tRides, tRideStatuses, tCoupons})
	rides := store.ridesByUser[user.ID]
	// ORDER BY created_at DESC
	for i := len(rides) - 1; i >= 0; i-- {
		ride := rides[i]
		if !store.rideCompleted(ride) {
			continue
		}

		fare := store.discountedFare(user.ID, ride, ride.PickupLatitude, ride.PickupLongitude, ride.DestinationLatitude, ride.DestinationLongitude)

		item := getAppRidesResponseItem{
			ID:                    ride.ID,
			PickupCoordinate:      Coordinate{Latitude: ride.PickupLatitude, Longitude: ride.PickupLongitude},
			DestinationCoordinate: Coordinate{Latitude: ride.DestinationLatitude, Longitude: ride.DestinationLongitude},
			Fare:                  fare,
			Evaluation:            *ride.Evaluation,
			RequestedAt:           ride.CreatedAt.UnixMilli(),
			CompletedAt:           ride.UpdatedAt.UnixMilli(),
		}

		chair := store.chairs[ride.ChairID.String]
		if chair == nil {
			unlock()
			writeError(w, http.StatusInternalServerError, errors.New("sql: no rows in result set"))
			return
		}
		item.Chair = getAppRidesResponseItemChair{ID: chair.ID, Name: chair.Name, Model: chair.Model}
		owner := store.owners[chair.OwnerID]
		if owner == nil {
			unlock()
			writeError(w, http.StatusInternalServerError, errors.New("sql: no rows in result set"))
			return
		}
		item.Chair.Owner = owner.Name

		items = append(items, item)
	}
	unlock()

	writeJSON(w, http.StatusOK, &getAppRidesResponse{
		Rides: items,
	})
}

func appPostRides(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &appPostRidesRequest{}
	if err := bindJSON(r, req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if req.PickupCoordinate == nil || req.DestinationCoordinate == nil {
		writeError(w, http.StatusBadRequest, errors.New("required fields(pickup_coordinate, destination_coordinate) are empty"))
		return
	}

	user := ctx.Value("user").(*User)
	rideID := ulid.Make().String()

	unlock := lockTables([]table{tRides, tRideStatuses, tCoupons}, nil)
	for _, ride := range store.ridesByUser[user.ID] {
		if !store.rideCompleted(ride) {
			unlock()
			writeError(w, http.StatusConflict, errors.New("ride already exists"))
			return
		}
	}

	t := now()
	ride := &Ride{
		ID: rideID, UserID: user.ID,
		PickupLatitude: req.PickupCoordinate.Latitude, PickupLongitude: req.PickupCoordinate.Longitude,
		DestinationLatitude: req.DestinationCoordinate.Latitude, DestinationLongitude: req.DestinationCoordinate.Longitude,
		CreatedAt: t, UpdatedAt: t,
	}
	store.addRide(ride)
	store.addStatus(&RideStatus{ID: ulid.Make().String(), RideID: rideID, Status: "MATCHING", CreatedAt: now()})

	var coupon *Coupon
	if len(store.ridesByUser[user.ID]) == 1 {
		// 初回利用で、初回利用クーポンがあれば必ず使う
		coupon = store.unusedCoupon(user.ID, "CP_NEW2024")
	}
	if coupon == nil {
		// 他のクーポンを付与された順番に使う
		coupon = store.unusedCoupon(user.ID, "")
	}
	if coupon != nil {
		id := rideID
		coupon.UsedBy = &id
		store.couponByRide[rideID] = coupon
	}

	fare := store.discountedFare(user.ID, ride, req.PickupCoordinate.Latitude, req.PickupCoordinate.Longitude, req.DestinationCoordinate.Latitude, req.DestinationCoordinate.Longitude)
	dirty.Store(true)
	unlock()

	writeJSON(w, http.StatusAccepted, &appPostRidesResponse{
		RideID: rideID,
		Fare:   fare,
	})
}

func appPostRidesEstimatedFare(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &appPostRidesEstimatedFareRequest{}
	if err := bindJSON(r, req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if req.PickupCoordinate == nil || req.DestinationCoordinate == nil {
		writeError(w, http.StatusBadRequest, errors.New("required fields(pickup_coordinate, destination_coordinate) are empty"))
		return
	}

	user := ctx.Value("user").(*User)

	unlock := lockTables(nil, []table{tCoupons})
	discounted := store.discountedFare(user.ID, nil, req.PickupCoordinate.Latitude, req.PickupCoordinate.Longitude, req.DestinationCoordinate.Latitude, req.DestinationCoordinate.Longitude)
	unlock()

	writeJSON(w, http.StatusOK, &appPostRidesEstimatedFareResponse{
		Fare:     discounted,
		Discount: calculateFare(req.PickupCoordinate.Latitude, req.PickupCoordinate.Longitude, req.DestinationCoordinate.Latitude, req.DestinationCoordinate.Longitude) - discounted,
	})
}

// マンハッタン距離を求める
func calculateDistance(aLatitude, aLongitude, bLatitude, bLongitude int) int {
	return abs(aLatitude-bLatitude) + abs(aLongitude-bLongitude)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 評価のレスポンスを返してから COMPLETED を足すまでの待ち
const completeDelay = 50 * time.Millisecond

func appPostRideEvaluatation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rideID := r.PathValue("ride_id")

	req := &appPostRideEvaluationRequest{}
	if err := bindJSON(r, req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if req.Evaluation < 1 || req.Evaluation > 5 {
		writeError(w, http.StatusBadRequest, errors.New("evaluation must be between 1 and 5"))
		return
	}

	// 検証と決済額の計算はロック下で、決済ゲートウェイの呼び出しはロックの外で行う
	unlock := lockTables(nil, []table{tPaymentTokens, tRides, tRideStatuses, tCoupons, tSettings})
	ride := store.rides[rideID]
	if ride == nil {
		unlock()
		writeError(w, http.StatusNotFound, errors.New("ride not found"))
		return
	}
	if status, _ := store.latestStatus(ride.ID); status != "ARRIVED" {
		unlock()
		writeError(w, http.StatusBadRequest, errors.New("not arrived yet"))
		return
	}
	paymentToken := store.paymentTokens[ride.UserID]
	if paymentToken == nil {
		unlock()
		writeError(w, http.StatusBadRequest, errors.New("payment token not registered"))
		return
	}
	fare := store.discountedFare(ride.UserID, ride, ride.PickupLatitude, ride.PickupLongitude, ride.DestinationLatitude, ride.DestinationLongitude)
	paymentGatewayURL := store.paymentGatewayURL
	token := paymentToken.Token
	unlock()

	paymentGatewayRequest := &paymentGatewayPostPaymentRequest{
		Amount: fare,
	}
	if err := requestPaymentGatewayPostPayment(ctx, paymentGatewayURL, token, rideID, paymentGatewayRequest); err != nil {
		if errors.Is(err, erroredUpstream) {
			writeError(w, http.StatusBadGateway, err)
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	// 評価を記録してレスポンスを返しきってから、COMPLETED を足す。
	// 椅子が空き（付近の椅子に出る・次のライドが割り当たる）になるのは COMPLETED から。
	// ベンチは評価のレスポンスを受け取るまでその椅子をライド中とみなすので、先に COMPLETED を足すと、
	// レスポンスが届くまでの間に付近の椅子を頼んだユーザーが CODE 30 を出す
	unlock = lockTables([]table{tRides}, []table{tRideStatuses})
	if status, _ := store.latestStatus(ride.ID); status != "ARRIVED" || ride.Evaluation != nil {
		unlock()
		writeError(w, http.StatusBadRequest, errors.New("not arrived yet"))
		return
	}
	evaluation := req.Evaluation
	ride.Evaluation = &evaluation
	ride.UpdatedAt = now()
	completedAt := ride.UpdatedAt
	dirty.Store(true)
	unlock()

	buf, _ := json.Marshal(&appPostRideEvaluationResponse{CompletedAt: completedAt.UnixMilli()})
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	// ベンチは、評価のレスポンスを受け取ってから椅子を空きとみなす。負荷が高いとそれが遅れるので、少し待ってから足す
	time.AfterFunc(completeDelay, func() {
		unlock := lockTables([]table{tRideStatuses, tChairStats}, []table{tRides})
		defer unlock()
		// /api/initialize でストアが置き換わっていたら何もしない
		if store.rides[rideID] != ride {
			return
		}
		if st, _ := store.latestStatus(rideID); st != "ARRIVED" {
			return
		}
		store.addStatus(&RideStatus{ID: ulid.Make().String(), RideID: rideID, Status: "COMPLETED", CreatedAt: completedAt})
		if ride.ChairID.Valid {
			st := store.chairStat[ride.ChairID.String]
			if st == nil {
				st = &chairStats{}
				store.chairStat[ride.ChairID.String] = st
			}
			st.Count++
			st.EvaluationSum += evaluation
		}
		dirty.Store(true)
	})
}

func (s *Store) chairStatsOf(chairID string) appGetNotificationResponseChairStats {
	stats := appGetNotificationResponseChairStats{}
	if st := s.chairStat[chairID]; st != nil && st.Count > 0 {
		stats.TotalRidesCount = st.Count
		stats.TotalEvaluationAvg = float64(st.EvaluationSum) / float64(st.Count)
	}
	return stats
}

func appGetNearbyChairs(w http.ResponseWriter, r *http.Request) {
	latStr := r.URL.Query().Get("latitude")
	lonStr := r.URL.Query().Get("longitude")
	distanceStr := r.URL.Query().Get("distance")
	if latStr == "" || lonStr == "" {
		writeError(w, http.StatusBadRequest, errors.New("latitude or longitude is empty"))
		return
	}

	lat, err := strconv.Atoi(latStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("latitude is invalid"))
		return
	}

	lon, err := strconv.Atoi(lonStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("longitude is invalid"))
		return
	}

	distance := 50
	if distanceStr != "" {
		distance, err = strconv.Atoi(distanceStr)
		if err != nil {
			writeError(w, http.StatusBadRequest, errors.New("distance is invalid"))
			return
		}
	}

	coordinate := Coordinate{Latitude: lat, Longitude: lon}

	nearbyChairs := []appGetNearbyChairsResponseChair{}
	unlock := lockTables(nil, []table{tChairs, tChairLocs, tRides, tRideStatuses})
	for _, chair := range store.chairList {
		if !chair.IsActive {
			continue
		}

		// 過去にライドが存在し、かつ、それが完了していない場合はスキップ
		// （前のライドが完了してから次が割り当てられるので、最新のライドだけを見れば足りる）
		if ride := store.latestRideOfChair(chair.ID); ride != nil {
			if status, _ := store.latestStatus(ride.ID); status != "COMPLETED" {
				continue
			}
		}

		// 最新の位置情報（行ロックで読む）
		agg := store.chairLoc[chair.ID]
		if agg == nil {
			continue
		}
		loc, ok := agg.get()
		if !ok {
			continue
		}

		if calculateDistance(coordinate.Latitude, coordinate.Longitude, loc.Latitude, loc.Longitude) <= distance {
			nearbyChairs = append(nearbyChairs, appGetNearbyChairsResponseChair{
				ID:    chair.ID,
				Name:  chair.Name,
				Model: chair.Model,
				CurrentCoordinate: Coordinate{
					Latitude:  loc.Latitude,
					Longitude: loc.Longitude,
				},
			})
		}
	}
	unlock()

	writeJSON(w, http.StatusOK, &appGetNearbyChairsResponse{
		Chairs:      nearbyChairs,
		RetrievedAt: now().UnixMilli(),
	})
}

func calculateFare(pickupLatitude, pickupLongitude, destLatitude, destLongitude int) int {
	meteredFare := farePerDistance * calculateDistance(pickupLatitude, pickupLongitude, destLatitude, destLongitude)
	return initialFare + meteredFare
}

// unusedCoupon はユーザーの未使用クーポンを付与順に探す。code が空なら種類を問わない。
func (s *Store) unusedCoupon(userID, code string) *Coupon {
	for _, c := range s.coupons[userID] {
		if c.UsedBy == nil && (code == "" || c.Code == code) {
			return c
		}
	}
	return nil
}

// discountedFare は calculateDiscountedFare と同じ。呼び出し側が mu を握る。
func (s *Store) discountedFare(userID string, ride *Ride, pickupLatitude, pickupLongitude, destLatitude, destLongitude int) int {
	discount := 0
	if ride != nil {
		destLatitude = ride.DestinationLatitude
		destLongitude = ride.DestinationLongitude
		pickupLatitude = ride.PickupLatitude
		pickupLongitude = ride.PickupLongitude

		// すでにクーポンが紐づいているならそれの割引額を参照
		if c := s.couponByRide[ride.ID]; c != nil {
			discount = c.Discount
		}
	} else {
		// 初回利用クーポンを最優先で使う。無いなら他のクーポンを付与された順番に使う
		if c := s.unusedCoupon(userID, "CP_NEW2024"); c != nil {
			discount = c.Discount
		} else if c := s.unusedCoupon(userID, ""); c != nil {
			discount = c.Discount
		}
	}

	meteredFare := farePerDistance * calculateDistance(pickupLatitude, pickupLongitude, destLatitude, destLongitude)
	discountedMeteredFare := max(meteredFare-discount, 0)

	return initialFare + discountedMeteredFare
}
