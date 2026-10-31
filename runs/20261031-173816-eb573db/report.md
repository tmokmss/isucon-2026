# run 20261031-173816-eb573db

## Summary

| | |
|---|---|
| score | 234286 |
| status | pass |
| ref | eb573db |
| sha | eb573dbfa002 |
| 所要時間 | 115s |
| 前回比 | +28725 (14.0%) vs 20261031-173327-f1c39e7 |
| ベスト比 | +2176 (0.9%) vs 20261031-164055-bdf05cd = 232110 **ベスト更新** |
| note | last run: code of 885b12c (233154) |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	52	151	413	636	637
europe	7	22	66	195	564	563
america	6	14	40	113	317	316
```

```
{"id":"35565","status":"done","portalStatus":"finished","passed":true,"score":234286,"raw":234286,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:38:45.291Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:38:48.286Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:38:51.280Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:38:54.294Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T08:38:57.069Z level=INFO msg="user moved" user=yarudan_5236 from=asia to=europe
time=2026-10-31T08:38:57.306Z level=INFO msg="europe 2 日目: 22 人"
time=2026-10-31T08:39:00.280Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T08:39:03.293Z level=INFO msg="asia 3 日目: 151 人"
time=2026-10-31T08:39:03.985Z level=INFO msg="user moved" user=koreta5480 from=america to=europe
time=2026-10-31T08:39:05.497Z level=INFO msg="user moved" user=miahe.5681 from=asia to=europe
time=2026-10-31T08:39:06.197Z level=INFO msg="user moved" user=miahe.5681 from=europe to=asia
time=2026-10-31T08:39:06.281Z level=INFO msg="europe 3 日目: 66 人"
time=2026-10-31T08:39:06.976Z level=INFO msg="user moved" user=benru-5926 from=asia to=europe
time=2026-10-31T08:39:08.670Z level=INFO msg="user moved" user=nokima.6116 from=asia to=america
time=2026-10-31T08:39:09.280Z level=INFO msg="america 3 日目: 40 人"
time=2026-10-31T08:39:10.162Z level=INFO msg="user moved" user=ismo-6374 from=europe to=asia
time=2026-10-31T08:39:10.471Z level=INFO msg="user moved" user=roorshi_6341 from=europe to=asia
time=2026-10-31T08:39:11.150Z level=INFO msg="user moved" user=yarudan_5236 from=europe to=asia
time=2026-10-31T08:39:12.174Z level=INFO msg="user moved" user=koreta5480 from=europe to=asia
time=2026-10-31T08:39:12.292Z level=INFO msg="asia 4 日目: 413 人"
time=2026-10-31T08:39:15.003Z level=INFO msg="user moved" user=miahe.5681 from=asia to=europe
time=2026-10-31T08:39:15.281Z level=INFO msg="europe 4 日目: 195 人"
time=2026-10-31T08:39:16.894Z level=INFO msg="user moved" user=rokani-7248 from=asia to=europe
time=2026-10-31T08:39:17.988Z level=INFO msg="user moved" user=ismo-6374 from=asia to=europe
time=2026-10-31T08:39:18.288Z level=INFO msg="america 4 日目: 113 人"
time=2026-10-31T08:39:19.149Z level=INFO msg="user moved" user=yoben-8105 from=asia to=america
time=2026-10-31T08:39:20.104Z level=INFO msg="user moved" user=benru-5926 from=europe to=america
time=2026-10-31T08:39:20.534Z level=INFO msg="user moved" user=temiaha.9057 from=europe to=america
time=2026-10-31T08:39:21.298Z level=INFO msg="asia 5 日目: 636 人"
time=2026-10-31T08:39:21.484Z level=INFO msg="user moved" user=tsudanvin-9014 from=europe to=asia
time=2026-10-31T08:39:24.307Z level=INFO msg="europe 5 日目: 564 人"
time=2026-10-31T08:39:25.979Z level=INFO msg="user moved" user=nokima.6116 from=america to=asia
time=2026-10-31T08:39:27.372Z level=INFO msg="america 5 日目: 317 人"
time=2026-10-31T08:39:36.297Z level=INFO msg="america 最終日（グランプリ）: 316 人"
time=2026-10-31T08:39:36.318Z level=INFO msg="europe 最終日（グランプリ）: 563 人"
time=2026-10-31T08:39:36.355Z level=INFO msg="asia 最終日（グランプリ）: 637 人"
time=2026-10-31T08:39:45.348Z level=INFO msg="グランプリ: アジア 525 人, ヨーロッパ 506 人, アメリカ 310 人"
{
  "pass": true,
  "score": 234286,
  "messages": [
    "mimiahi_5052: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "anretsu.6336: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "leokuyo-10214: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "niis8057: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maso_9970: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noul.11939: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sechi.5290: 友達申請を承認できません: POST /api/friend-requests/de6375f7-6ce2-4ec1-b80b-4442c1c259d6/accept: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests/de6375f7-6ce2-4ec1-b80b-4442c1c259d6/accept\": EOF",
    "fuki5487: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sardines_en_cocotte_de_terre_2092: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
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
全 11107 クエリ / 48 種類 / 合計 129.445s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	43.3%	203	56.017	275.95	1013.81	1143.7	116.2	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	18.9%	4	24.524	6130.95	9514.87	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	13.0%	3	16.774	5591.28	7227.35	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	11.7%	7857	15.114	1.92	9.16	10.1	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.7%	4	8.694	2173.52	3549.52	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	1.9%	8	2.491	311.40	400.91	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	1.3%	8	1.703	212.83	626.57	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	1.1%	1	1.392	1392.22	1392.22	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
9	0.6%	8	0.725	90.61	158.43	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.3%	1	0.420	419.71	419.71	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.3%	1	0.355	354.59	354.59	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.2%	1	0.213	213.30	213.30	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.146	145.52	145.52	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	8	0.137	17.15	26.15	166.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	0.1%	1	0.108	107.87	107.87	2433.0	0.0	DELETE FROM image_objects WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM posts;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
未取得: metrics.pprofEnabled = false（アプリに pprof を仕込んでいない）
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 34.0% / 100% | 7% | 1.1% | 5 | 158 MB | run queue が積もっている |
| isu2 (2 cpu) | 31.8% / 100% | 10% | 3.3% | 5 | 102 MB | run queue が積もっている |
| isu3 (2 cpu) | 28.8% / 100% | 7% | 4.6% | 8 | 120 MB | run queue が積もっている |
| isu4 (2 cpu) | 70.1% / 100% | 61% | 3.0% | 5 | 101 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 52% / nginx 10% |
| isu2 | isureal 53% / nginx 3% |
| isu3 | isureal 45% / nginx 2% |
| isu4 | isureal 116% / postgres 16% |

## Diff from previous run

スコア: 205561 → 234286 (+28725)

`f1c39e7..eb573db`

```
 conf/isu1/etc/nginx/conf.d/isuh-log.conf       |  1 -
 conf/isu1/etc/nginx/sites-enabled/isureal.conf |  1 -
 conf/isu2/etc/nginx/conf.d/isuh-log.conf       |  1 -
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |  1 -
 conf/isu3/etc/nginx/conf.d/isuh-log.conf       |  1 -
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |  1 -
 conf/isu4/home/isucon/env.sh                   |  2 --
 webapp/go/cmd/isureal/pprof.go                 | 11 ++++++
 webapp/go/image_signature.go                   | 46 +++++---------------------
 webapp/go/jpegcheck/huffman.go                 |  9 -----
 webapp/go/offload.go                           | 20 +++--------
 11 files changed, 25 insertions(+), 69 deletions(-)
```
