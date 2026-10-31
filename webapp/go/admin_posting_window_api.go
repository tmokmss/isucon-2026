package isureal

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

// 投稿期間の開始は、利用者のリクエストで共有のプールが埋まっていても待たされないよう、専用の接続で書く。
// 期間の終了・開始・通知の 3 文は 1 回の往復で送る（pgx の batch は 1 つの暗黙のトランザクションになる）
var postingWindowPool struct {
	once sync.Once
	pool *pgxpool.Pool
	err  error
}

func postingWindowDB(config *Config) (*pgxpool.Pool, error) {
	p := &postingWindowPool
	p.once.Do(func() {
		pc, err := pgxpool.ParseConfig(config.DatabaseDSN())
		if err != nil {
			p.err = err
			return
		}
		pc.MaxConns = 2
		p.pool, p.err = pgxpool.NewWithConfig(context.Background(), pc)
	})
	return p.pool, p.err
}

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
		ctx := c.Request().Context()
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
		notificationIDs := make([]string, len(userIDs))
		for i, id := range userIDs {
			notificationIDs[i] = uuid.NewString()
			w.notified[id] = notificationRow{ID: notificationIDs[i], CreatedAt: startsAt}
		}
		pool, err := postingWindowDB(h.Config)
		if err != nil {
			return err
		}
		batch := &pgx.Batch{}
		batch.Queue(`UPDATE posting_windows SET ended_at = $2 WHERE region = $1 AND ended_at IS NULL`, region, startsAt)
		batch.Queue(`
   INSERT INTO posting_windows (id, region, local_date, starts_at, on_time_deadline, created_at)
   VALUES ($1, $2, $3, $4, $5, $4)
  `, w.ID, region, localDate, startsAt, w.OnTimeDeadline)
		if len(userIDs) > 0 {
			batch.Queue(`
   INSERT INTO notifications (id, user_id, posting_window_id, created_at)
   SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
  `, notificationIDs, userIDs, w.ID, startsAt)
		}
		if err := pool.SendBatch(ctx, batch).Close(); err != nil {
			return err
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
