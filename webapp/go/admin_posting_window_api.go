package isureal

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type openPostingWindowRequest struct {
	Region        string `json:"region"`
	LocalDate     string `json:"localDate"`
	OnTimeSeconds *int64 `json:"onTimeSeconds"`
}

type openPostingWindowResponse struct {
	ID             string    `json:"id"`
	Region         string    `json:"region"`
	LocalDate      string    `json:"localDate"`
	StartsAt       time.Time `json:"startsAt"`
	OnTimeDeadline time.Time `json:"onTimeDeadline"`
}

var postingWindowMu sync.Mutex

func openPostingWindowHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		var input openPostingWindowRequest
		if err := c.Bind(&input); err != nil {
			return err
		}
		region := input.Region
		localDate, err := time.Parse(time.DateOnly, input.LocalDate)
		if !validRegion(region) || err != nil {
			return NewAPIError(http.StatusBadRequest, "invalid_posting_window")
		}
		seconds := int64(24 * 60 * 60)
		if input.OnTimeSeconds != nil {
			seconds = *input.OnTimeSeconds
		}
		s := h.Store
		postingWindowMu.Lock()
		defer postingWindowMu.Unlock()
		s.mu.RLock()
		maxDate, hasAny := s.maxLocalDate[region]
		conflict := hasAny && !maxDate.Before(localDate)
		var userIDs []string
		if !conflict {
			for _, u := range s.users {
				if u.Region == region && u.active() {
					userIDs = append(userIDs, u.ID)
				}
			}
		}
		var oldWindow *WindowRow
		if w := s.activeWindows[region]; w != nil {
			oldWindow = w
		}
		s.mu.RUnlock()
		if conflict {
			return NewAPIError(http.StatusConflict, "posting_window_conflict")
		}
		startsAt := now()
		w := &WindowRow{
			PostingWindow: PostingWindow{
				ID: uuid.NewString(), Region: region, LocalDate: localDate, StartsAt: startsAt,
				OnTimeDeadline: startsAt.Add(time.Duration(seconds) * time.Second),
			},
			CreatedAt: startsAt,
			notified:  make(map[string]notificationRow, len(userIDs)),
		}
		for _, id := range userIDs {
			w.notified[id] = notificationRow{ID: uuid.NewString(), CreatedAt: startsAt}
		}
		s.mu.Lock()
		if oldWindow != nil && oldWindow.EndedAt == nil {
			t := startsAt
			oldWindow.EndedAt = &t
			// 終わった投稿期間の通知はもう読まれない
			oldWindow.notified = map[string]notificationRow{}
		}
		s.windows[w.ID] = w
		s.activeWindows[region] = w
		s.maxLocalDate[region] = localDate
		s.broadcastLocked()
		s.mu.Unlock()
		return c.JSON(http.StatusCreated, openPostingWindowResponse{
			ID: w.ID, Region: region, LocalDate: localDate.Format(time.DateOnly),
			StartsAt: startsAt, OnTimeDeadline: w.OnTimeDeadline,
		})
	}
}
