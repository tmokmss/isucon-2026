package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	json "github.com/goccy/go-json"
	"io"
	"net/http"
	"time"
)

var erroredUpstream = errors.New("errored upstream")

type paymentGatewayPostPaymentRequest struct {
	Amount int `json:"amount"`
}

// 決済サービスは直近 3 秒の処理件数 / 100（最大 50%）の確率で失敗を返し、同じトークンで 4 回失敗すると次は必ず通す。
// 失敗率は試す回数に依らないので、待たずに試し直す（試すたびに 50ms 待っていたら、評価 API が平均 350ms 掛かっていた）
const (
	paymentMaxAttempts = 20
)

// 決済サービスへの接続を使い回す（http.DefaultClient は待機させる接続がホストあたり 2 本で、同時に数百件の決済のたびに張り直す）
var paymentClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        1024,
		MaxIdleConnsPerHost: 1024,
		IdleConnTimeout:     90 * time.Second,
	},
}

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
		if attempt > 0 && ctx.Err() != nil {
			return ctx.Err()
		}
		lastErr = func() error {
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, paymentGatewayURL+"/payments", bytes.NewBuffer(b))
			if err != nil {
				return err
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Idempotency-Key", idempotencyKey)

			res, err := paymentClient.Do(req)
			if err != nil {
				return err
			}
			defer func() {
				_, _ = io.Copy(io.Discard, res.Body) // 読み切ってから閉じると接続を使い回せる
				res.Body.Close()
			}()
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
