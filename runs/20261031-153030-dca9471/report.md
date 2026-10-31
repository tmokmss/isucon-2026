# run 20261031-153030-dca9471

## Summary

| | |
|---|---|
| score | 186885 |
| status | pass |
| ref | presign-next-minute |
| sha | dca9471d652e |
| 所要時間 | 134s |
| 前回比 | +3554 (1.9%) vs 20261031-152137-8c13252 |
| ベスト比 | -4931 (-2.6%) vs 20261031-151339-5a230a4 = 191816 |
| note | presign next-minute image signatures on peers |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	141	413	1123	1124
europe	7	20	59	173	456	456
america	7	18	58	166	307	305
```

```
{"id":"22914","status":"done","portalStatus":"finished","passed":true,"score":186885,"raw":186885,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:30:56.164Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:30:59.161Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:31:01.218Z level=INFO msg="user moved" user=noaliazoe5088 from=europe to=america
time=2026-10-31T06:31:02.160Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T06:31:05.217Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T06:31:08.159Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T06:31:11.160Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T06:31:13.516Z level=INFO msg="user moved" user=noaliazoe5088 from=america to=europe
time=2026-10-31T06:31:14.173Z level=INFO msg="asia 3 日目: 141 人"
time=2026-10-31T06:31:14.183Z level=INFO msg="user moved" user=nohata.5432 from=europe to=america
time=2026-10-31T06:31:14.353Z level=INFO msg="user moved" user=nifuhi5407 from=europe to=asia
time=2026-10-31T06:31:15.995Z level=INFO msg="user moved" user=nifuhi5407 from=asia to=america
time=2026-10-31T06:31:17.161Z level=INFO msg="europe 3 日目: 59 人"
time=2026-10-31T06:31:17.595Z level=INFO msg="user moved" user=muni-6175 from=asia to=america
time=2026-10-31T06:31:18.181Z level=INFO msg="user moved" user=muni-6175 from=america to=asia
time=2026-10-31T06:31:19.635Z level=INFO msg="user moved" user=noaliazoe5088 from=europe to=america
time=2026-10-31T06:31:20.160Z level=INFO msg="america 3 日目: 58 人"
time=2026-10-31T06:31:23.172Z level=INFO msg="asia 4 日目: 413 人"
time=2026-10-31T06:31:25.225Z level=INFO msg="user moved" user=ivyzoeleo.6530 from=america to=europe
time=2026-10-31T06:31:25.468Z level=INFO msg="user moved" user=nohata.5432 from=america to=asia
time=2026-10-31T06:31:26.303Z level=INFO msg="user moved" user=yohael_6467 from=europe to=asia
time=2026-10-31T06:31:26.386Z level=INFO msg="europe 4 日目: 173 人"
time=2026-10-31T06:31:27.974Z level=INFO msg="user moved" user=nohata.5432 from=asia to=america
time=2026-10-31T06:31:27.992Z level=INFO msg="user moved" user=kuta.7431 from=asia to=america
time=2026-10-31T06:31:28.240Z level=INFO msg="user moved" user=fuvinka_7116 from=asia to=america
time=2026-10-31T06:31:29.111Z level=INFO msg="user moved" user=sovindan_8297 from=asia to=europe
time=2026-10-31T06:31:29.147Z level=INFO msg="user moved" user=shizoere.8182 from=asia to=america
time=2026-10-31T06:31:29.273Z level=INFO msg="america 4 日目: 166 人"
time=2026-10-31T06:31:29.332Z level=INFO msg="user moved" user=nois8137 from=asia to=europe
time=2026-10-31T06:31:29.667Z level=INFO msg="user moved" user=yohael_6467 from=asia to=america
time=2026-10-31T06:31:29.867Z level=INFO msg="user moved" user=solel_7419 from=asia to=america
time=2026-10-31T06:31:30.627Z level=INFO msg="user moved" user=ivyzoeleo.6530 from=europe to=america
time=2026-10-31T06:31:31.417Z level=INFO msg="user moved" user=moshi8694 from=asia to=america
time=2026-10-31T06:31:32.196Z level=INFO msg="asia 5 日目: 1123 人"
time=2026-10-31T06:31:36.251Z level=INFO msg="europe 5 日目: 456 人"
time=2026-10-31T06:31:37.340Z level=INFO msg="user moved" user=shihaha-8744 from=europe to=asia
time=2026-10-31T06:31:38.352Z level=INFO msg="america 5 日目: 307 人"
time=2026-10-31T06:31:38.848Z level=INFO msg="user moved" user=noaliazoe5088 from=america to=europe
time=2026-10-31T06:31:47.863Z level=INFO msg="europe 最終日（グランプリ）: 456 人"
time=2026-10-31T06:31:48.695Z level=INFO msg="asia 最終日（グランプリ）: 1124 人"
time=2026-10-31T06:31:49.193Z level=INFO msg="america 最終日（グランプリ）: 305 人"
time=2026-10-31T06:31:56.272Z level=INFO msg="グランプリ: アジア 381 人, ヨーロッパ 232 人, アメリカ 95 人"
{
  "pass": true,
  "score": 186885,
  "messages": [
    "karemo-5130: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "roto.17185: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "anna_becker_2684: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "liasuke.15410: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nasudan18122: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "keka18113: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "orshian11228: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ulto-16174: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "bennufu16771: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hoyamia_17577: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marmia.17196: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "istona_17930: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "muanul10267: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "bensumu.17539: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hame.16120: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nusemia-16271: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "moru.15653: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yana-17790: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "issama.17478: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "wayo-9222: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
（リクエストなし）

== isu2 ==
（リクエストなし）

== isu3 ==
（リクエストなし）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 18341 クエリ / 48 種類 / 合計 153.745s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	55.4%	430	85.122	197.96	743.09	1219.5	124.9	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	13.9%	3	21.360	7120.06	9514.62	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.7%	14273	17.982	1.26	15.62	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.4%	3	15.927	5309.16	7160.43	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.0%	3	7.717	2572.18	3453.68	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.2%	4	3.351	837.87	2863.74	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.2%	4	0.360	89.89	156.64	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
8	0.2%	4	0.358	89.49	161.71	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.2%	1	0.282	281.95	281.95	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.254	254.02	254.02	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.2%	1	0.252	251.94	251.94	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.167	166.57	166.57	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.106	105.73	105.73	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.081	80.59	80.59	2463.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.072	71.72	71.72	1519.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT data FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 7c6eba147a18bfe6b4f33aeb72878d8433f28cc2
Type: cpu
Time: 2026-10-31 15:30:36 JST
Duration: 60.07s, Total samples = 47.04s (78.31%)
Showing nodes accounting for 33.22s, 70.62% of 47.04s total
Dropped 1116 nodes (cum <= 0.24s)
Showing top 40 nodes out of 289
      flat  flat%   sum%        cum   cum%
     6.28s 13.35% 13.35%      6.28s 13.35%  crypto/internal/fips140/sha256.blockAVX2
     5.16s 10.97% 24.32%      5.16s 10.97%  internal/runtime/syscall/linux.Syscall6
     2.45s  5.21% 29.53%      2.45s  5.21%  runtime.memmove
     1.84s  3.91% 33.44%      7.29s 15.50%  runtime.scanObject
     1.83s  3.89% 37.33%      2.40s  5.10%  runtime.findObject
     1.78s  3.78% 41.11%      2.96s  6.29%  runtime.tryDeferToSpanScan
     1.77s  3.76% 44.88%      5.85s 12.44%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.49s  3.17% 48.04%      1.49s  3.17%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.47s  3.12% 51.17%      1.47s  3.12%  runtime.memclrNoHeapPointers
     1.40s  2.98% 54.15%      1.40s  2.98%  aeshashbody
     0.90s  1.91% 56.06%      0.90s  1.91%  runtime.spanClass.sizeclass (inline)
     0.56s  1.19% 57.25%      0.56s  1.19%  memeqbody
     0.50s  1.06% 58.31%      2.27s  4.83%  runtime.mapaccess1_faststr
     0.43s  0.91% 59.23%      0.57s  1.21%  example.com/isureal-company/isureal.postLess
     0.42s  0.89% 60.12%      0.42s  0.89%  runtime.nextFreeFast (inline)
     0.34s  0.72% 60.84%         7s 14.88%  crypto/internal/fips140/sha256.(*Digest).Write
     0.31s  0.66% 61.50%      7.14s 15.18%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.30s  0.64% 62.14%      0.75s  1.59%  database/sql.convertAssignRows
     0.27s  0.57% 62.71%      0.27s  0.57%  indexbody
     0.27s  0.57% 63.29%      0.27s  0.57%  runtime.futex
     0.25s  0.53% 63.82%      0.62s  1.32%  runtime.greyobject
     0.23s  0.49% 64.31%      4.55s  9.67%  runtime.mallocgc
     0.22s  0.47% 64.77%      0.44s  0.94%  compress/flate.(*compressor).deflate
     0.22s  0.47% 65.24%      1.16s  2.47%  runtime.mallocgcSmallScanNoHeader
     0.22s  0.47% 65.71%      0.26s  0.55%  runtime.spanOf (inline)
     0.21s  0.45% 66.16%      7.51s 15.97%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.21s  0.45% 66.60%      0.26s  0.55%  encoding/json.appendString[go.shape.string]
     0.20s  0.43% 67.03%      0.89s  1.89%  runtime.wbBufFlush1
     0.19s   0.4% 67.43%      0.59s  1.25%  runtime.mapaccess2_faststr
     0.19s   0.4% 67.84%      0.56s  1.19%  runtime.scanObjectsSmall
     0.18s  0.38% 68.22%      0.27s  0.57%  runtime.(*mspan).markBitsForIndex (inline)
     0.17s  0.36% 68.58%      2.10s  4.46%  runtime.mapassign_faststr
     0.16s  0.34% 68.92%      0.39s  0.83%  internal/runtime/maps.(*Iter).Next
     0.14s   0.3% 69.22%      0.84s  1.79%  encoding/json.structEncoder.encode
     0.12s  0.26% 69.47%      7.79s 16.56%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.12s  0.26% 69.73%      1.11s  2.36%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.11s  0.23% 69.96%      4.94s 10.50%  example.com/isureal-company/isureal.listMemories
     0.11s  0.23% 70.20%      2.74s  5.82%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.21% 70.41%      1.11s  2.36%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
     0.10s  0.21% 70.62%      0.88s  1.87%  runtime.scanSpan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 42.8% / 100% | 28% | 0.7% | 6 | 103 MB | run queue が積もっている |
| isu2 (2 cpu) | 35.2% / 100% | 8% | 0.7% | 6 | 178 MB | run queue が積もっている |
| isu3 (2 cpu) | 25.0% / 100% | 6% | 0.5% | 5 | 212 MB | run queue が積もっている |
| isu4 (2 cpu) | 38.8% / 100% | 14% | 14.9% | 21 | 105 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 69% / nginx 11% |
| isu2 | isureal 66% / nginx 2% |
| isu3 | isureal 47% / nginx 1% |
| isu4 | isureal 31% / postgres 15% |

## Diff from previous run

スコア: 183331 → 186885 (+3554)

`8c13252..dca9471`

```
 webapp/go/image_held.go     |  5 +---
 webapp/go/image_local.go    |  1 -
 webapp/go/offload.go        | 59 +++++++++++++++++++++++++++++++++++++++++++++
 webapp/go/persist.go        |  2 +-
 webapp/go/realmoji_api.go   | 41 ++++++++++++-------------------
 webapp/go/replica_server.go |  8 +-----
 webapp/go/server.go         |  1 +
 webapp/go/store.go          |  3 +--
 8 files changed, 80 insertions(+), 40 deletions(-)
```
