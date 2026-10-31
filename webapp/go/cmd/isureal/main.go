package main

import (
	"os"
	"time"

	"example.com/isureal-company/isureal"
)

func main() {
	// 状態を持たない台（nginx が /api/ を asia1 へ転送する台）ではアプリを動かさない
	if os.Getenv("ISUREAL_PASSIVE") == "1" {
		for {
			time.Sleep(time.Hour)
		}
	}
	config := isureal.ConfigFromEnv()
	server := isureal.NewServer(config)
	if err := server.Start(":8080"); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
