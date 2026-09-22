// 競技アプリの代役。deploy が「どの版のバイナリを配ったか」を外から確認できればよい。
// version は up.sh が version.go を書き換えてコミットごとに変える。
package main

import (
	"fmt"
	"net/http"
	"os"
)

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
		fmt.Fprintln(w, `{"language":"go"}`)
	})

	fmt.Fprintf(os.Stderr, "isuapp %s listening on %s\n", version, addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
