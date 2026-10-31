package main

import (
	"net/http"
	_ "net/http/pprof"
)

// 競技の最後に外す（isuh finalize が確認する）
func init() {
	go func() { _ = http.ListenAndServe("localhost:6060", nil) }()
}
