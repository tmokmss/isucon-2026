# ISUCON2026 当日の記録

振り返りは tmokmss/isucon-harness の Issue #107。ここにはその根拠になった生の記録を置く。

- `clarifications.md`: 公開質問と回答（競技後にポータルから取得）
- `leaderboard/`: 上位 30 チームとうちのスコアの履歴（ポータルの /api/audience/dashboard/teams/{id}）。`flyhigh-jumps.tsv` は各チームの「直前の 0 点でないランからの上昇幅」の最大
- `pprof-asia2/`: asia2 の CPU プロファイル（17:28 は initialize の読み直しを含む 30 秒、17:29 は負荷の最中の 20 秒）。`go tool pprof -top -cum isu4b.pprof` で読む
- `sigbench/main.go`: 画像 URL の署名（HMAC-SHA256 を 1,000 回）の速さを台ごとに比べる。`GOOS=linux GOARCH=amd64 go build` して各台で動かす。c5 で 0.461ms、c7i（SHA-NI あり）で 0.143ms だった
- `isureal-guard/`: 運営配布の isureal-guard のソース（各台の /opt/isureal-guard/src）
- `pg_schema.sql`・`pg_tables.txt`: asia2 の PostgreSQL のスキーマと、競技終了時の表ごとの行数
