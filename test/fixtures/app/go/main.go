// 競技アプリの代役。deploy が「どの版のバイナリを配ったか」を外から確認できればよい。
// version は up.sh が version.go を、レスポンスの形は shape.go を書いて変える。
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func writeJSON(w http.ResponseWriter, value any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(value)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Println(version)
		return
	}

	addr := os.Getenv("ISUAPP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	http.HandleFunc("/version", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, version)
	})
	http.HandleFunc("/initialize", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{"language": "go"})
	})

	// check(#12) が形を見るエンドポイント
	http.HandleFunc("/api/items", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]any{
			"items": []Item{{ID: 1, Name: "isu", Tags: []string{"a", "b"}}},
			"total": 1,
		})
	})
	http.HandleFunc("/api/user", func(w http.ResponseWriter, _ *http.Request) {
		// email は null。null を取りうるフィールドで差分が出ないことの確認に使う。
		writeJSON(w, map[string]any{"id": 1, "name": "isucon", "email": nil})
	})

	fmt.Fprintf(os.Stderr, "isuapp %s listening on %s\n", version, addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
