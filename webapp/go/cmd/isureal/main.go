package main

import (
	"net/http"
	"os"

	"example.com/isureal-company/isureal"
)

func main() {
	// 逃がし先の台（asia2）は、署名の計算だけを受け付ける。DB にはつながず、データも読み込まない
	if os.Getenv("ISUREAL_OFFLOAD_WORKER") == "1" {
		panic(isureal.ServeOffload())
	}
	config := isureal.ConfigFromEnv()
	// 複製を持つ台（europe1・america1）は、asia1 から変更を受けて読み取りの一部を手元で返す
	if os.Getenv("ISUREAL_PASSIVE") == "1" && os.Getenv("ISUREAL_REPLICA_MODE") != "" {
		go func() { panic(isureal.ServeOffload()) }()
		panic(http.ListenAndServe(":8080", isureal.NewReplicaServer(config)))
	}
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
	// echo は SIGINT・SIGTERM で受け付けを止めて戻る。スナップショットを書き終えるまで終わらない
	isureal.WaitStopSnapshot()
}
