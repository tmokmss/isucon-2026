# run 20261031-152604-1884ea4

## Summary

| | |
|---|---|
| score | 185078 |
| status | pass |
| ref | final-no-metrics |
| sha | 1884ea4d7c35 |
| 所要時間 | 123s |
| 前回比 | +1747 (1.0%) vs 20261031-152137-8c13252 |
| ベスト比 | -6738 (-3.5%) vs 20261031-151339-5a230a4 = 191816 |
| note | final: access log and pprof off (on main 710ad2d) |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	143	423	1145	1146
europe	7	19	58	167	447	447
america	6	18	54	153	292	292
```

```
{"id":"22542","status":"done","portalStatus":"finished","passed":true,"score":185078,"raw":185078,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:26:31.661Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:26:34.660Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:26:37.657Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:26:40.668Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T06:26:43.663Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T06:26:46.657Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T06:26:49.668Z level=INFO msg="asia 3 日目: 143 人"
time=2026-10-31T06:26:50.015Z level=INFO msg="user moved" user=chiko-5580 from=asia to=europe
time=2026-10-31T06:26:52.660Z level=INFO msg="europe 3 日目: 58 人"
time=2026-10-31T06:26:53.402Z level=INFO msg="user moved" user=ismeho.5508 from=america to=asia
time=2026-10-31T06:26:53.985Z level=INFO msg="user moved" user=ismeho.5508 from=asia to=america
time=2026-10-31T06:26:55.657Z level=INFO msg="america 3 日目: 54 人"
time=2026-10-31T06:26:56.848Z level=INFO msg="user moved" user=rufuko6322 from=europe to=america
time=2026-10-31T06:26:58.667Z level=INFO msg="asia 4 日目: 423 人"
time=2026-10-31T06:27:01.735Z level=INFO msg="europe 4 日目: 167 人"
time=2026-10-31T06:27:02.916Z level=INFO msg="user moved" user=kiso6992 from=asia to=america
time=2026-10-31T06:27:04.860Z level=INFO msg="america 4 日目: 153 人"
time=2026-10-31T06:27:05.444Z level=INFO msg="user moved" user=miamiasu.7825 from=asia to=america
time=2026-10-31T06:27:05.790Z level=INFO msg="user moved" user=tatezoe_6963 from=asia to=europe
time=2026-10-31T06:27:07.180Z level=INFO msg="user moved" user=chiko-5580 from=europe to=asia
time=2026-10-31T06:27:07.745Z level=INFO msg="asia 5 日目: 1145 人"
time=2026-10-31T06:27:09.498Z level=INFO msg="user moved" user=kebenmia_9543 from=america to=asia
time=2026-10-31T06:27:11.327Z level=INFO msg="europe 5 日目: 447 人"
time=2026-10-31T06:27:11.564Z level=INFO msg="user moved" user=rufuko6322 from=america to=asia
time=2026-10-31T06:27:14.774Z level=INFO msg="america 5 日目: 292 人"
time=2026-10-31T06:27:24.071Z level=INFO msg="asia 最終日（グランプリ）: 1146 人"
time=2026-10-31T06:27:24.623Z level=INFO msg="america 最終日（グランプリ）: 292 人"
time=2026-10-31T06:27:24.730Z level=INFO msg="europe 最終日（グランプリ）: 447 人"
time=2026-10-31T06:27:31.772Z level=INFO msg="グランプリ: アジア 534 人, ヨーロッパ 112 人, アメリカ 66 人"
{
  "pass": true,
  "score": 185078,
  "messages": [
    "nera.5025: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "mohore_17405: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ulnoro.16747: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "teno.11802: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "isra.16815: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "benhemu9257: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "muzoe15909: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kuwawa-17669: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "evamezoe_16177: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "neor-17311: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "vinvinivy7455: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nobenso16846: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "motsu_15974: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "danrichi_16859: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "koniro-17098: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "katois_17301: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chimoho17422: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "miadan-18013: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "anre15568: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nehisu-17592: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
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
全 17631 クエリ / 48 種類 / 合計 153.917s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	55.9%	416	86.029	206.80	1359.56	1196.3	122.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	13.8%	3	21.289	7096.40	9523.54	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.2%	13888	17.202	1.24	8.80	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.4%	3	15.960	5319.86	7218.38	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.1%	3	7.922	2640.67	3683.97	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.2%	4	3.329	832.37	2842.14	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.2%	4	0.350	87.52	156.19	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.2%	4	0.348	87.04	157.28	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.259	259.25	259.25	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.249	249.34	249.34	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.1%	1	0.216	215.94	215.94	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.166	165.58	165.58	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.104	103.83	103.83	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.084	83.57	83.57	2486.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.067	67.04	67.04	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
未取得: metrics.pprofEnabled = false（アプリに pprof を仕込んでいない）
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 44.4% / 100% | 28% | 0.5% | 4 | 137 MB |  |
| isu2 (2 cpu) | 29.7% / 100% | 5% | 0.6% | 5 | 182 MB | run queue が積もっている |
| isu3 (2 cpu) | 25.9% / 100% | 6% | 0.7% | 5 | 390 MB | run queue が積もっている |
| isu4 (2 cpu) | 41.1% / 100% | 15% | 16.3% | 7 | 107 MB | run queue が積もっている |

**isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 73% / nginx 12% |
| isu2 | isureal 56% / nginx 2% |
| isu3 | isureal 50% / nginx 1% |
| isu4 | isureal 33% / postgres 15% |

## Diff from previous run

スコア: 183331 → 185078 (+1747)

`8c13252..1884ea4`

```
 conf/isu1/etc/nginx/conf.d/isuh-log.conf |  1 +
 conf/isu2/etc/nginx/conf.d/isuh-log.conf |  1 +
 conf/isu3/etc/nginx/conf.d/isuh-log.conf |  1 +
 webapp/go/cmd/isureal/pprof.go           | 11 ---------
 webapp/go/image_held.go                  |  5 +---
 webapp/go/image_local.go                 |  1 -
 webapp/go/persist.go                     |  2 +-
 webapp/go/realmoji_api.go                | 41 +++++++++++++-------------------
 webapp/go/replica_server.go              |  8 +------
 webapp/go/store.go                       |  3 +--
 10 files changed, 23 insertions(+), 51 deletions(-)
```
