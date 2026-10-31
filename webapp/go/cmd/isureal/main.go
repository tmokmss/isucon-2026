package main

import (
	"os"

	"example.com/isureal-company/isureal"
)

func main() {
	// CPU の空いている台は、JPEG の検証と署名の計算だけを受け付ける
	if os.Getenv("ISUREAL_OFFLOAD_WORKER") == "1" {
		panic(isureal.ServeOffload())
	}
	config := isureal.ConfigFromEnv()
	// 状態を持たない台（nginx が /api/ を asia1 へ転送する台）は、/api/images/ だけを返す
	server := isureal.NewImageServer
	if os.Getenv("ISUREAL_PASSIVE") != "1" {
		server = isureal.NewServer
		isureal.StartImageEventPublisher()
	} else {
		go func() { panic(isureal.ServeOffload()) }()
	}
	if err := server(config).Start(":8080"); err != nil {
		panic("failed to start server: " + err.Error())
	}
}
