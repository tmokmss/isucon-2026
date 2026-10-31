package isureal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

type getNotificationsResponse struct {
	ID              string    `json:"id"`
	PostingWindowID string    `json:"postingWindowId"`
	Region          string    `json:"region"`
	CreatedAt       time.Time `json:"createdAt"`
}

// pendingNotificationLocked は接続後に作られ、投稿期間がまだ終わっていない通知のうち最も古いものを返す
func (s *Store) pendingNotificationLocked(userID string, connectedAt time.Time) *getNotificationsResponse {
	var best *getNotificationsResponse
	for _, w := range s.activeWindows {
		n, ok := w.notified[userID]
		if !ok || !n.CreatedAt.After(connectedAt) {
			continue
		}
		if best == nil || n.CreatedAt.Before(best.CreatedAt) || (n.CreatedAt.Equal(best.CreatedAt) && n.ID < best.ID) {
			best = &getNotificationsResponse{ID: n.ID, PostingWindowID: w.ID, Region: w.Region, CreatedAt: n.CreatedAt}
		}
	}
	return best
}

func getNotificationsHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := c.Request().Context()
		session := GetSession(c)
		userID := session.UserID
		connectedAt := now()
		response := c.Response()
		response.Header().Set("Content-Type", "text/event-stream")
		response.WriteHeader(http.StatusOK)
		controller := http.NewResponseController(response)
		if err := controller.Flush(); err != nil {
			return nil
		}
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		s := h.Store
		for {
			s.mu.RLock()
			u := s.activeUser(userID)
			var region string
			if u != nil {
				region = u.Region
			}
			var notification *getNotificationsResponse
			if u != nil && region == session.Region {
				notification = s.pendingNotificationLocked(userID, connectedAt)
			}
			wait := s.notifyCh
			s.mu.RUnlock()
			if u == nil || region != session.Region {
				return nil
			}
			if notification != nil {
				data, err := json.Marshal(notification)
				if err != nil {
					c.Logger().Error("failed to encode notification", "error", err)
					return nil
				}
				if _, err := fmt.Fprintf(response, "data: %s\n\n", data); err != nil {
					return nil
				}
				_ = controller.Flush()
				return nil
			}
			select {
			case <-ctx.Done():
				return nil
			case <-wait:
			case <-ticker.C:
			}
		}
	}
}
