package isureal

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

type currentPostingWindowResponse struct {
	ID             string    `json:"id"`
	Region         string    `json:"region"`
	LocalDate      string    `json:"localDate"`
	StartsAt       time.Time `json:"startsAt"`
	OnTimeDeadline time.Time `json:"onTimeDeadline"`
}

type getCurrentPostingWindowResponse struct {
	Window        *currentPostingWindowResponse `json:"window"`
	HasPosted     bool                          `json:"hasPosted"`
	CanPost       bool                          `json:"canPost"`
	DeletionCount int                           `json:"deletionCount"`
}

func getCurrentPostingWindowHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		s := h.Store
		s.mu.RLock()
		u := s.users[GetSession(c).UserID]
		var w *WindowRow
		if u != nil {
			w = s.activeWindows[u.Region]
		}
		if w == nil {
			s.mu.RUnlock()
			return c.JSON(http.StatusOK, getCurrentPostingWindowResponse{})
		}
		window := w.PostingWindow
		active, deleted := postCountsLocked(u, w.LocalDate)
		s.mu.RUnlock()
		return c.JSON(http.StatusOK, getCurrentPostingWindowResponse{
			Window: &currentPostingWindowResponse{
				ID: window.ID, Region: window.Region, LocalDate: window.LocalDate.Format(time.DateOnly),
				StartsAt: window.StartsAt, OnTimeDeadline: window.OnTimeDeadline,
			},
			HasPosted:     active > 0,
			CanPost:       active == 0 && deleted < 2 && !time.Now().Before(window.StartsAt),
			DeletionCount: deleted,
		})
	}
}
