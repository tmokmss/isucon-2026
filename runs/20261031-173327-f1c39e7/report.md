# run 20261031-173327-f1c39e7

## Summary

| | |
|---|---|
| score | 205561 |
| status | pass |
| ref | origin/main |
| sha | f1c39e71a0b9 |
| 所要時間 | 137s |
| 前回比 | -79 (-0.0%) vs 20261031-172947-21d10c4 |
| ベスト比 | -26549 (-11.4%) vs 20261031-164055-bdf05cd = 232110 |
| note | final candidate: sse 20s + asia2 sign EU/US + sig cache + jpeg prefetch + metrics off |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	53	161	461	1039	1040
europe	7	23	72	211	595	594
america	6	16	46	129	313	313
```

```
{"id":"35053","status":"done","portalStatus":"finished","passed":true,"score":205561,"raw":205561,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:33:55.578Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:33:58.569Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:34:01.569Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:34:04.588Z level=INFO msg="asia 2 日目: 53 人"
time=2026-10-31T08:34:05.708Z level=INFO msg="user moved" user=tsufu.5119 from=america to=asia
time=2026-10-31T08:34:06.006Z level=INFO msg="user moved" user=tsufu.5119 from=asia to=america
time=2026-10-31T08:34:06.790Z level=INFO msg="user moved" user=ormaxro.5280 from=asia to=europe
time=2026-10-31T08:34:07.468Z level=INFO msg="user moved" user=ormaxro.5280 from=europe to=asia
time=2026-10-31T08:34:07.569Z level=INFO msg="europe 2 日目: 23 人"
time=2026-10-31T08:34:10.572Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T08:34:13.581Z level=INFO msg="asia 3 日目: 161 人"
time=2026-10-31T08:34:14.908Z level=INFO msg="user moved" user=seya-5711 from=asia to=america
time=2026-10-31T08:34:15.039Z level=INFO msg="user moved" user=romi-5796 from=asia to=europe
time=2026-10-31T08:34:15.550Z level=INFO msg="user moved" user=seya-5711 from=america to=asia
time=2026-10-31T08:34:15.767Z level=INFO msg="user moved" user=romi-5796 from=europe to=america
time=2026-10-31T08:34:16.379Z level=INFO msg="user moved" user=ormaxro.5280 from=asia to=europe
time=2026-10-31T08:34:16.570Z level=INFO msg="europe 3 日目: 72 人"
time=2026-10-31T08:34:17.941Z level=INFO msg="user moved" user=neivyka-6298 from=asia to=europe
time=2026-10-31T08:34:19.597Z level=INFO msg="america 3 日目: 46 人"
time=2026-10-31T08:34:22.582Z level=INFO msg="asia 4 日目: 461 人"
time=2026-10-31T08:34:23.076Z level=INFO msg="user moved" user=hihi-6739 from=america to=asia
time=2026-10-31T08:34:23.723Z level=INFO msg="user moved" user=seya-5711 from=asia to=america
time=2026-10-31T08:34:24.062Z level=INFO msg="user moved" user=tsufu.5119 from=america to=asia
time=2026-10-31T08:34:25.578Z level=INFO msg="europe 4 日目: 211 人"
time=2026-10-31T08:34:26.305Z level=INFO msg="user moved" user=hihi-6739 from=asia to=america
time=2026-10-31T08:34:26.537Z level=INFO msg="user moved" user=tsufu.5119 from=asia to=america
time=2026-10-31T08:34:28.207Z level=INFO msg="user moved" user=neivyka-6298 from=europe to=america
time=2026-10-31T08:34:28.708Z level=INFO msg="america 4 日目: 129 人"
time=2026-10-31T08:34:28.746Z level=INFO msg="user moved" user=sakael8976 from=asia to=america
time=2026-10-31T08:34:31.011Z level=INFO msg="user moved" user=riha-9686 from=europe to=america
time=2026-10-31T08:34:31.039Z level=INFO msg="user moved" user=romi-5796 from=america to=europe
time=2026-10-31T08:34:31.585Z level=INFO msg="asia 5 日目: 1039 人"
time=2026-10-31T08:34:33.290Z level=INFO msg="user moved" user=ormaxro.5280 from=europe to=asia
time=2026-10-31T08:34:33.750Z level=INFO msg="user moved" user=zoekanoa_9097 from=europe to=asia
time=2026-10-31T08:34:34.956Z level=INFO msg="europe 5 日目: 595 人"
time=2026-10-31T08:34:37.900Z level=INFO msg="america 5 日目: 313 人"
time=2026-10-31T08:34:46.871Z level=INFO msg="asia 最終日（グランプリ）: 1040 人"
time=2026-10-31T08:34:47.091Z level=INFO msg="europe 最終日（グランプリ）: 594 人"
time=2026-10-31T08:34:47.098Z level=INFO msg="america 最終日（グランプリ）: 313 人"
time=2026-10-31T08:34:55.754Z level=INFO msg="グランプリ: アジア 533 人, ヨーロッパ 227 人, アメリカ 209 人"
{
  "pass": true,
  "score": 205561,
  "messages": [
    "yuulro5084: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "fuzoe_5047: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "kasu.7845: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noato.18571: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tsumarte-9378: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "liahe.8363: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sasutor13640: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "isivyfu8619: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yuxin_wu_1185: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "danhire-10982: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "narishi.18394: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
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
全 16435 クエリ / 48 種類 / 合計 186.049s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	54.9%	378	102.065	270.01	2168.00	1205.2	123.1	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	13.7%	4	25.437	6359.31	9528.98	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.0%	12327	20.399	1.65	11.13	9.7	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.2%	4	19.005	4751.16	7209.72	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	4.8%	4	9.020	2254.93	3408.19	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	3.3%	7	6.055	865.04	3258.47	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	1	0.779	779.40	779.40	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.3%	7	0.631	90.18	157.26	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.3%	7	0.617	88.09	160.82	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.2%	1	0.407	407.46	407.46	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.337	336.81	336.81	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.2%	1	0.287	287.08	287.08	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.177	177.06	177.06	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.131	131.20	131.20	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	1	0.110	110.43	110.43	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows

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
| isu1 (2 cpu) | 34.9% / 100% | 12% | 0.7% | 7 | 119 MB | run queue が積もっている |
| isu2 (2 cpu) | 29.0% / 100% | 8% | 0.7% | 7 | 107 MB | run queue が積もっている |
| isu3 (2 cpu) | 19.6% / 100% | 2% | 0.7% | 5 | 349 MB | run queue が積もっている |
| isu4 (2 cpu) | 51.5% / 100% | 39% | 6.1% | 7 | 97 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 54% / nginx 10% |
| isu2 | isureal 54% / nginx 2% |
| isu3 | isureal 37% / nginx 2% |
| isu4 | isureal 70% / postgres 18% |

## Diff from previous run

スコア: 205640 → 205561 (-79)

`21d10c4..f1c39e7`

```
 conf/isu1/etc/nginx/conf.d/isuh-log.conf |  1 +
 conf/isu2/etc/nginx/conf.d/isuh-log.conf |  1 +
 conf/isu3/etc/nginx/conf.d/isuh-log.conf |  1 +
 webapp/go/cmd/isureal/pprof.go           | 11 --------
 webapp/go/image_signature.go             | 46 +++++++++++++++++++++++++-------
 webapp/go/jpegcheck/huffman.go           |  9 +++++++
 webapp/go/offload.go                     | 10 ++++---
 7 files changed, 55 insertions(+), 24 deletions(-)
```
