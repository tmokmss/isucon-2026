package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var erroredUpstream = errors.New("errored upstream")

type paymentGatewayPostPaymentRequest struct {
	Amount int `json:"amount"`
}

const (
	paymentMaxAttempts = 20
	paymentRetryWait   = 50 * time.Millisecond
)

// requestPaymentGatewayPostPayment は決済を依頼する。
// Idempotency-Key（ライドの ID）を付けるので、失敗したら支払履歴を確かめずに同じキーで再送してよい
// （同じキーのリクエストは決済サービスが冪等に処理する。ISURIDE.md「Idempotency-Key ヘッダを利用したリクエストの重複防止」）
func requestPaymentGatewayPostPayment(ctx context.Context, paymentGatewayURL string, token string, idempotencyKey string, param *paymentGatewayPostPaymentRequest) error {
	b, err := json.Marshal(param)
	if err != nil {
		return err
	}

	var lastErr error
	for attempt := 0; attempt < paymentMaxAttempts; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(paymentRetryWait):
			}
		}
		lastErr = func() error {
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, paymentGatewayURL+"/payments", bytes.NewBuffer(b))
			if err != nil {
				return err
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Idempotency-Key", idempotencyKey)

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer res.Body.Close()
			if res.StatusCode != http.StatusNoContent {
				return fmt.Errorf("[POST /payments] unexpected status code (%d): %w", res.StatusCode, erroredUpstream)
			}
			return nil
		}()
		if lastErr == nil {
			return nil
		}
	}
	return lastErr
}
