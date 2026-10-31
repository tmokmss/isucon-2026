package main

import (
	"os"

	"example.com/isureal-company/isureal"
)

func main() {
	config := isureal.ConfigFromEnv()
	// 状態を持たない台（nginx が /api/ を asia1 へ転送する台）は、/api/images/ だけを返す
	server := isureal.NewImageServer
	if os.Getenv("ISUREAL_PASSIVE") != "1" {
		server = isureal.NewServer
		isureal.StartImageEventPublisher()
	}
	if err := server(config).Start(":8080"); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
