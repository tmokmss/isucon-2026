# run 20261031-132851-90a4ef9

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | image-files |
| sha | 90a4ef9e55ab |
| 所要時間 | 532s |
| 前回比 | -64700 (-100.0%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | -64700 (-100.0%) vs 20261031-121235-e0e4cb3 = 64700 |
| note | image bodies to files on asia1 (init/ new/), DB data empty |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/agent-ae5c3d5012bdce542/runs/20261031-132851-90a4ef9/bench.log
```

## Bench

```
{"id":"13372","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"初期データの検証に失敗しました"}

--- stdout ---
{
  "pass": false,
  "score": 0,
  "messages": [
    "初期データの検証に失敗しました: 初期データのユーザー charcoal_mini_turkey_2371: GET /api/images/1de3bc25-8f41-475b-b0a5-41eac65a9ce4?exp=1793507880\u0026sig=H53Vz55BeMnSCgNVmKRxQVVSlEaF_3GTjFd9b7GyUDk: ステータスコード 500 が返りました（期待値 200、ERROR: syntax error at or near \"FROM\" (SQLSTATE 42601)）",
    "初期データのユーザー davi_pereira_3221: GET /api/images/8353d8df-b8a5-481b-acc5-1e5fca566c66?exp=1793507880\u0026sig=2GUzxT5bhgSo-0ZP83fDdGeFNJyE1FWnx3zlyMfzoiA: ステータスコード 500 が返りました（期待値 200、ERROR: syntax error at or near \"FROM\" (SQLSTATE 42601)）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
87.1%	1	0	0	POST	/api/initialize	22.895	22.895	22.895	22.895	90
4.0%	6	0	0	GET	/api/memories	1.043	0.174	0.237	0.237	59095
3.3%	3	0	0	GET	/api/me/export	0.859	0.286	0.530	0.530	2237053
3.2%	24	1	0	POST	/api/auth/login	0.854	0.036	0.070	0.070	221
1.6%	12	1	0	POST	/api/auth/register	0.418	0.035	0.064	0.064	209
0.3%	8	0	0	GET	^/api/relationships/[^/]+	0.088	0.011	0.023	0.023	12444
0.1%	45	2	0	GET	^/api/images/[^/]+	0.038	0.001	0.003	0.003	6263
0.1%	7	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	0.022	0.003	0.004	0.004	116
0.1%	3	2	0	PUT	/api/me/profile-image	0.019	0.006	0.015	0.015	23
0.0%	7	0	0	POST	/api/friend-requests	0.013	0.002	0.003	0.003	115
0.0%	8	0	0	PUT	^/api/follows/[^/]+$	0.013	0.002	0.003	0.003	0
0.0%	2	0	0	PUT	^/api/blocks/[^/]+$	0.009	0.005	0.005	0.005	0
0.0%	4	0	0	PUT	^/api/admin/users/[^/]+/certification$	0.005	0.001	0.002	0.002	0
0.0%	1	0	0	POST	/api/admin/auth/login	0.004	0.004	0.004	0.004	106
0.0%	1	0	0	POST	/api/auth/logout	0.001	0.001	0.001	0.001	0
0.0%	2	2	0	OPTIONS	/	0.000	0.000	0.000	0.000	166
0.0%	3	0	0	GET	/	0.000	0.000	0.000	0.000	392
0.0%	1	1	0	DELETE	^/api/admin/users/[^/]+/certification$	0.000	0.000	0.000	0.000	39
0.0%	1	0	0	GET	/files/	0.000	0.000	0.000	0.000	392
0.0%	1	1	0	GET	^/api/posts/me/current$	0.000	0.000	0.000	0.000	27
0.0%	1	0	0	GET	^/api/feeds/[^/]+	0.000	0.000	0.000	0.000	29
0.0%	1	0	0	GET	/favicon.ico	0.000	0.000	0.000	0.000	392
0.0%	4	2	0	GET	/api/me	0.000	0.000	0.000	0.000	134
0.0%	1	1	0	OEPR	/	0.000	0.000	0.000	0.000	166
0.0%	1	0	0	GET	/pub/	0.000	0.000	0.000	0.000	392
0.0%	1	0	0	GET	/downloads/	0.000	0.000	0.000	0.000	392
0.0%	1	0	0	HEAD	/	0.000	0.000	0.000	0.000	0
0.0%	1	1	0	POST	/	0.000	0.000	0.000	0.000	166

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.8%	6	0	6	GET	^/api/images/[^/]+	2.650	0.442	0.511	0.511	69
13.4%	8	0	0	POST	/api/auth/login	1.047	0.131	0.207	0.207	237
13.1%	5	0	0	POST	/api/auth/register	1.023	0.205	0.206	0.206	226
12.3%	1	0	0	GET	/api/me/export	0.960	0.960	0.960	0.960	2900169
6.6%	2	0	0	GET	/api/memories	0.514	0.257	0.309	0.309	21716
6.5%	4	0	0	PUT	^/api/follows/[^/]+$	0.509	0.127	0.202	0.202	0
5.3%	3	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	0.419	0.140	0.206	0.206	116
3.9%	3	0	0	POST	/api/friend-requests	0.307	0.102	0.103	0.103	115
2.6%	1	1	0	GET	/api/me	0.202	0.202	0.202	0.202	28
1.3%	1	0	0	PUT	^/api/blocks/[^/]+$	0.104	0.104	0.104	0.104	0
1.3%	1	0	0	GET	^/api/relationships/[^/]+	0.101	0.101	0.101	0.101	2301

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.3%	6	0	6	GET	^/api/images/[^/]+	2.654	0.442	0.512	0.512	69
14.8%	5	0	0	POST	/api/auth/register	1.024	0.205	0.206	0.206	227
13.4%	7	0	0	POST	/api/auth/login	0.931	0.133	0.207	0.207	226
8.0%	1	0	0	GET	/api/me/export	0.553	0.553	0.553	0.553	579071
6.2%	2	0	0	GET	/api/memories	0.430	0.215	0.218	0.218	4886
6.0%	3	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	0.419	0.140	0.207	0.207	116
5.8%	4	0	0	PUT	^/api/follows/[^/]+$	0.405	0.101	0.102	0.102	0
4.4%	3	0	0	POST	/api/friend-requests	0.306	0.102	0.102	0.102	115
1.5%	1	0	0	PUT	^/api/blocks/[^/]+$	0.106	0.106	0.106	0.106	0
1.5%	1	0	0	GET	^/api/relationships/[^/]+	0.102	0.102	0.102	0.102	2223
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 927 クエリ / 59 種類 / 合計 16.932s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	41.3%	2	6.990	3495.07	6831.39	231740.0	2727.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
2	18.9%	1	3.194	3194.02	3194.02	30732.0	1199807.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.8%	1	2.001	2000.69	2000.69	23539.0	627389.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.0%	2	1.360	680.05	1351.46	16773.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
5	6.5%	1	1.093	1092.84	1092.84	12428.0	530010.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	4.7%	1	0.804	804.20	804.20	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
7	1.8%	1	0.309	309.12	309.12	28096.0	2278.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	1.4%	1	0.244	243.82	243.82	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
9	1.1%	1	0.191	191.23	191.23	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
10	0.9%	1	0.148	147.76	147.76	31192.0	459.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.6%	1	0.100	100.22	100.22	2646.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
12	0.5%	1	0.078	77.93	77.93	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
13	0.4%	1	0.074	74.43	74.43	1548.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations
14	0.4%	1	0.068	68.28	68.28	12902.0	473.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.3%	1	0.053	53.40	53.40	2433.0	0.0	DELETE FROM image_objects WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM posts;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: a05ce2b05f989f05151549365f35cd95f42128e9
Type: cpu
Time: 2026-10-31 13:36:42 JST
Duration: 60s, Total samples = 13.39s (22.32%)
Showing nodes accounting for 10.99s, 82.08% of 13.39s total
Dropped 278 nodes (cum <= 0.07s)
Showing top 40 nodes out of 151
      flat  flat%   sum%        cum   cum%
     1.91s 14.26% 14.26%      6.09s 45.48%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.22s  9.11% 23.38%      1.22s  9.11%  aeshashbody
     1.21s  9.04% 32.41%      1.21s  9.04%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.46s  3.44% 35.85%      1.32s  9.86%  runtime.scanObject
     0.45s  3.36% 39.21%      1.87s 13.97%  runtime.mapaccess1_faststr
     0.45s  3.36% 42.57%      0.45s  3.36%  runtime.memmove
     0.38s  2.84% 45.41%      0.61s  4.56%  example.com/isureal-company/isureal.postLess
     0.33s  2.46% 47.87%      0.55s  4.11%  database/sql.convertAssignRows
     0.33s  2.46% 50.34%      0.95s  7.09%  runtime.mallocgcSmallScanNoHeader
     0.33s  2.46% 52.80%      0.50s  3.73%  runtime.tryDeferToSpanScan
     0.30s  2.24% 55.04%      0.30s  2.24%  memeqbody
     0.26s  1.94% 56.98%      2.80s 20.91%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.25s  1.87% 58.85%      0.25s  1.87%  internal/runtime/syscall/linux.Syscall6
     0.25s  1.87% 60.72%      0.39s  2.91%  runtime.findObject
     0.20s  1.49% 62.21%      1.58s 11.80%  runtime.mallocgc
     0.19s  1.42% 63.63%      0.19s  1.42%  sync/atomic.(*Int32).Add (inline)
     0.18s  1.34% 64.97%      0.18s  1.34%  runtime.nextFreeFast (inline)
     0.16s  1.19% 66.17%      0.19s  1.42%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.15s  1.12% 67.29%      0.15s  1.12%  runtime.memclrNoHeapPointers
     0.14s  1.05% 68.33%      1.70s 12.70%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.14s  1.05% 69.38%      0.14s  1.05%  internal/runtime/maps.(*Map).directoryAt (inline)
     0.13s  0.97% 70.35%      0.15s  1.12%  runtime.(*mspan).writeHeapBitsSmall
     0.13s  0.97% 71.32%      0.82s  6.12%  runtime.convTstring
     0.12s   0.9% 72.22%      0.12s   0.9%  cmpbody
     0.12s   0.9% 73.11%      0.12s   0.9%  crypto/internal/fips140/sha256.blockAVX2
     0.12s   0.9% 74.01%      4.96s 37.04%  example.com/isureal-company/isureal.loadBig
     0.12s   0.9% 74.91%      0.15s  1.12%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
     0.11s  0.82% 75.73%      0.12s   0.9%  example.com/isureal-company/isureal.rmLess
     0.11s  0.82% 76.55%      1.41s 10.53%  runtime.mapassign_faststr
     0.09s  0.67% 77.22%      0.09s  0.67%  runtime.(*spanInlineMarkBits).init
     0.08s   0.6% 77.82%      0.49s  3.66%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
     0.08s   0.6% 78.42%      0.08s   0.6%  runtime.spanClass.sizeclass (inline)
     0.07s  0.52% 78.94%      0.67s  5.00%  runtime.slicebytetostring
     0.07s  0.52% 79.46%      0.07s  0.52%  runtime.typePointers.nextFast (inline)
     0.06s  0.45% 79.91%      0.10s  0.75%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func5
     0.06s  0.45% 80.36%      0.74s  5.53%  github.com/jackc/pgx/v5/pgtype.scanPlanString.Scan
     0.06s  0.45% 80.81%      0.27s  2.02%  runtime.mallocgcSmallNoscan
     0.06s  0.45% 81.25%      0.09s  0.67%  runtime.mallocgcTiny
     0.06s  0.45% 81.70%      0.28s  2.09%  runtime.mapaccess2_faststr
     0.05s  0.37% 82.08%      0.17s  1.27%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 57.6% / 99% | 19% | 12.1% | 5 | 294 MB | run queue が積もっている |
| isu2 (2 cpu) | 25.0% / 55% | 0% | 22.3% | 1 | 1992 MB | IO 待ちが重い |
| isu3 (2 cpu) | 0.3% / 1% | 0% | 0.0% | 0 | 2513 MB |  |
| isu4 (2 cpu) | 32.0% / 96% | 4% | 22.4% | 2 | 103 MB | IO 待ちが重い |

**isu1（run queue が積もっている、主に isureal）、isu2（IO 待ちが重い、主に find）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 50% / du 18% / kcompactd0 3% |
| isu2 | find 4% / isureal 1% |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 13% / kcompactd0 1% |

## Diff from previous run

スコア: 64700 → — (fail)

`e0e4cb3..90a4ef9`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf     |   1 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf     |   1 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf     |   1 +
 .../etc/postgresql/16/main/conf.d/zz-tuning.conf   |   2 +
 webapp/go/account_deletion_api.go                  |   2 +-
 webapp/go/admin_user_api.go                        |   2 +-
 webapp/go/comment_api.go                           |   2 +-
 webapp/go/config.go                                |   2 +
 webapp/go/export_api.go                            |  56 +---
 webapp/go/feed_api.go                              | 113 +++++--
 webapp/go/feed_test.go                             | 331 +++++++++++++++++++++
 webapp/go/image_api.go                             |  17 +-
 webapp/go/image_file.go                            | 153 ++++++++++
 webapp/go/image_local.go                           |  38 +--
 webapp/go/initialize.go                            |   1 +
 webapp/go/offload.go                               |  27 +-
 webapp/go/post_api.go                              |  67 ++++-
 webapp/go/profile_image_api.go                     |   5 +-
 webapp/go/realmoji_api.go                          |   7 +-
 webapp/go/relationship.go                          |  10 +-
 webapp/go/server.go                                |   1 +
 webapp/go/store.go                                 | 113 ++++++-
 22 files changed, 803 insertions(+), 149 deletions(-)
```
