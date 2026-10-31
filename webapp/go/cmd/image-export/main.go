// 初期データの画像を、手元の PostgreSQL（AMI から来た止まったデータを一時的に起動したもの）から
// /home/isucon/images/init/<id の先頭 2 文字>/<id> に書き出す。既にあるファイルは書き直さないので、途中で止めても再実行できる。
// 他のベンチとディスクを取り合うので、ベンチのロック（isu1 の /tmp/isucon-bench.lock）を取ってから流す。手順は setup/image-export.sh。
package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
)

func main() {
	const dsn = "postgres://isureal:isureal@127.0.0.1:5433/isureal"
	const dir = "/home/isucon/images/init"

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := conn.Query(ctx, `SELECT id, data FROM image_objects WHERE created_at <= (SELECT created_at FROM initial_data_boundary)`)
	if err != nil {
		log.Fatal(err)
	}
	seen, written := 0, 0
	for rows.Next() {
		var id string
		var data []byte
		if err := rows.Scan(&id, &data); err != nil {
			log.Fatal(err)
		}
		seen++
		if seen%100000 == 0 {
			log.Printf("%d seen, %d written", seen, written)
		}
		sub := filepath.Join(dir, id[:2])
		path := filepath.Join(sub, id)
		if _, err := os.Stat(path); err == nil {
			continue
		}
		if err := os.MkdirAll(sub, 0o755); err != nil {
			log.Fatal(err)
		}
		tmp := filepath.Join(sub, "."+id+".tmp")
		if err := os.WriteFile(tmp, data, 0o644); err != nil {
			log.Fatal(err)
		}
		if err := os.Rename(tmp, path); err != nil {
			log.Fatal(err)
		}
		written++
	}
	if rows.Err() != nil {
		log.Fatal(rows.Err())
	}
	log.Printf("done: %d seen, %d written", seen, written)
}
