# run 20261031-150358-c416920

## Summary

| | |
|---|---|
| score | 1315 |
| status | pass |
| ref | origin/slow-before-1700-v2 |
| sha | c416920db0e2 |
| 所要時間 | 370s |
| 前回比 | -138715 (-99.1%) vs 20261031-141745-a28a701 |
| ベスト比 | -140020 (-99.1%) vs 20261031-141027-bd0fe45 = 141335 |
| note | slow-v2 3000ms after window opens |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	18	18	18	18	18
europe	7	7	7	7	7	7
america	6	6	6	6	6	6
```

```
{"id":"20728","status":"done","portalStatus":"finished","passed":true,"score":1315,"raw":1315,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:05:33.027Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:05:35.996Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:05:38.997Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:05:42.125Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T06:05:45.006Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T06:05:47.996Z level=INFO msg="america 2 日目: 6 人"
time=2026-10-31T06:05:51.023Z level=INFO msg="asia 3 日目: 18 人"
time=2026-10-31T06:05:53.997Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T06:05:57.008Z level=INFO msg="america 3 日目: 6 人"
time=2026-10-31T06:06:00.179Z level=INFO msg="asia 4 日目: 18 人"
time=2026-10-31T06:06:02.996Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T06:06:06.117Z level=INFO msg="america 4 日目: 6 人"
time=2026-10-31T06:06:09.020Z level=INFO msg="asia 5 日目: 18 人"
time=2026-10-31T06:06:12.018Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T06:06:14.995Z level=INFO msg="america 5 日目: 6 人"
time=2026-10-31T06:06:23.995Z level=INFO msg="america 最終日（グランプリ）: 6 人"
time=2026-10-31T06:06:24.019Z level=INFO msg="asia 最終日（グランプリ）: 18 人"
time=2026-10-31T06:06:24.021Z level=INFO msg="europe 最終日（グランプリ）: 7 人"
time=2026-10-31T06:06:32.996Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 2 人"
{
  "pass": true,
  "score": 1315,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
89.4%	12734	12657	0	POST	/api/auth/login	32574.281	2.558	3.003	3.020	15
5.6%	126	0	0	GET	/api/notifications	2055.892	16.317	63.001	63.001	153
1.1%	172	171	0	PUT	/api/me/profile-image	413.014	2.401	4.414	4.473	15
0.8%	102	0	0	GET	/api/memories	283.871	2.783	3.217	3.224	8761
0.5%	60	0	0	GET	/api/posting-window/current	180.045	3.001	3.002	3.002	236
0.4%	50	1	0	POST	/api/posts	159.854	3.197	3.671	3.671	588
0.4%	52	0	0	GET	^/api/feeds/[^/]+	153.040	2.943	3.002	3.002	51
0.4%	54	3	0	POST	/api/friend-requests	141.086	2.613	3.004	3.004	110
0.4%	46	42	0	GET	^/api/posts/me/current$	135.037	2.936	3.002	3.002	75
0.3%	40	0	0	GET	^/api/relationships/[^/]+	96.132	2.403	3.010	3.010	4402
0.2%	24	0	0	GET	/api/users	72.024	3.001	3.002	3.002	243
0.2%	34	1	0	POST	/api/auth/register	66.548	1.957	3.003	3.003	235
0.1%	19	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	36.070	1.898	3.005	3.005	116
0.0%	8	0	0	POST	/api/auth/logout	18.011	2.251	3.002	3.002	0
0.0%	1	0	0	POST	/api/initialize	15.884	15.884	15.884	15.884	90
0.0%	5	0	0	GET	^/api/memories/[^/]+	15.120	3.024	3.117	3.117	173
0.0%	6	0	0	GET	/api/me/export	11.241	1.873	3.481	3.481	2117499
0.0%	7	2	0	GET	/api/me	9.007	1.287	3.002	3.002	158
0.0%	2	0	0	DELETE	^/api/posts/[^/]+$	6.205	3.103	3.103	3.103	0
0.0%	2	2	0	PATCH	^/api/posts/[^/]+$	6.001	3.000	3.001	3.001	30
0.0%	357	172	0	GET	^/api/images/[^/]+	0.402	0.001	0.005	0.005	74000
0.0%	18	0	0	POST	/api/admin/posting-window	0.268	0.015	0.064	0.064	177
0.0%	8	0	0	PUT	^/api/follows/[^/]+$	0.016	0.002	0.003	0.003	0
0.0%	2	0	0	PUT	^/api/blocks/[^/]+$	0.010	0.005	0.005	0.005	0
0.0%	1	0	0	POST	/api/admin/auth/login	0.005	0.005	0.005	0.005	106
0.0%	4	0	0	PUT	^/api/admin/users/[^/]+/certification$	0.004	0.001	0.001	0.001	0
0.0%	1	0	0	POST	/api/admin/auth/logout	0.002	0.002	0.002	0.002	0
0.0%	1	1	0	DELETE	^/api/admin/users/[^/]+/certification$	0.000	0.000	0.000	0.000	39

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
55.4%	49	0	0	GET	/api/notifications	745.707	15.219	63.102	63.102	155
8.5%	38	0	0	GET	/api/memories	114.658	3.017	3.298	3.298	6369
5.2%	22	0	0	GET	/api/posting-window/current	69.451	3.157	3.362	3.362	239
4.9%	24	1	0	POST	/api/friend-requests	65.755	2.740	3.160	3.160	112
4.6%	20	0	0	GET	^/api/feeds/[^/]+	62.370	3.118	3.257	3.257	87
4.1%	16	0	0	POST	/api/posts	55.051	3.441	3.973	3.973	613
3.5%	16	0	0	GET	^/api/relationships/[^/]+	47.767	2.985	3.260	3.260	833
3.5%	15	12	0	GET	^/api/posts/me/current$	46.931	3.129	3.260	3.260	154
2.8%	12	0	0	GET	/api/users	37.764	3.147	3.259	3.259	196
2.2%	17	0	0	POST	/api/auth/login	29.413	1.730	3.169	3.169	242
1.2%	53	0	0	GET	^/api/images/[^/]+	16.692	0.315	0.984	0.984	209501
1.0%	5	2	0	GET	/api/me	13.099	2.620	3.259	3.259	127
0.7%	3	3	0	PATCH	^/api/posts/[^/]+$	9.392	3.131	3.189	3.189	30
0.6%	3	0	0	GET	/api/me/export	7.960	2.653	3.294	3.294	2149374
0.5%	5	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	6.641	1.328	3.104	3.104	116
0.5%	2	0	0	GET	^/api/memories/[^/]+	6.276	3.138	3.153	3.153	29
0.5%	2	0	0	POST	/api/auth/logout	6.204	3.102	3.102	3.102	0
0.2%	1	0	0	POST	/api/me/account-deletion	3.208	3.208	3.208	3.208	0
0.1%	5	0	0	POST	/api/auth/register	1.316	0.263	0.265	0.265	229
0.0%	4	0	0	PUT	^/api/follows/[^/]+$	0.573	0.143	0.165	0.165	0
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.115	0.115	0.115	0.115	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
58.6%	42	0	0	GET	/api/notifications	686.090	16.335	63.102	63.102	156
8.7%	34	0	0	GET	/api/memories	102.336	3.010	3.340	3.340	9232
4.6%	17	0	0	GET	/api/posting-window/current	53.355	3.139	3.361	3.361	240
4.0%	14	0	0	POST	/api/posts	47.072	3.362	3.863	3.863	585
3.7%	14	0	0	GET	^/api/feeds/[^/]+	43.649	3.118	3.259	3.259	29
3.5%	14	0	0	GET	^/api/relationships/[^/]+	41.455	2.961	3.259	3.259	674
3.5%	13	10	0	GET	^/api/posts/me/current$	40.948	3.150	3.260	3.260	159
3.5%	16	0	0	POST	/api/friend-requests	40.851	2.553	3.151	3.151	115
2.4%	9	0	0	GET	/api/users	27.945	3.105	3.259	3.259	230
2.2%	15	0	0	POST	/api/auth/login	26.047	1.736	3.125	3.125	248
1.4%	63	2	0	GET	^/api/images/[^/]+	16.184	0.257	0.826	0.826	160899
1.1%	4	0	0	POST	/api/auth/logout	12.494	3.123	3.153	3.153	0
0.8%	3	0	0	GET	/api/me	9.775	3.258	3.259	3.259	192
0.6%	2	0	0	DELETE	^/api/posts/[^/]+$	6.490	3.245	3.286	3.286	0
0.5%	2	0	0	GET	^/api/memories/[^/]+	6.313	3.157	3.258	3.258	898
0.3%	4	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	3.460	0.865	3.106	3.106	116
0.3%	1	1	0	PATCH	^/api/posts/[^/]+$	3.131	3.131	3.131	3.131	30
0.1%	1	0	0	GET	/api/me/export	1.755	1.755	1.755	1.755	5393917
0.1%	5	0	0	POST	/api/auth/register	1.461	0.292	0.295	0.295	228
0.0%	4	0	0	PUT	^/api/follows/[^/]+$	0.569	0.142	0.143	0.143	0
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.115	0.115	0.115	0.115	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 4875 クエリ / 71 種類 / 合計 97.358s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	41.8%	5	40.676	8135.16	9737.45	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
2	30.3%	5	29.508	5901.50	7175.37	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
3	14.9%	5	14.493	2898.62	3444.40	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
4	4.7%	6	4.542	756.92	3501.34	202894.0	2491.7	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
5	2.3%	9	2.275	252.76	623.55	1634.3	152.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
6	1.3%	6	1.276	212.74	824.21	11226.7	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
7	0.8%	1	0.795	795.45	795.45	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.7%	6	0.716	119.28	207.17	610.0	24993.7	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.7%	6	0.665	110.78	157.41	593.3	28988.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.4%	105	0.343	3.26	7.98	95.3	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $2) AS is_profile, EXISTS (SELECT $3 FROM posts p WHERE p.image_id…
11	0.3%	6	0.260	43.41	257.60	4.0	0.0	ALTER TABLE posting_windows ADD COLUMN IF NOT EXISTS notification_ids text[], ADD COLUMN IF NOT EXISTS notified_user_ids text[]
12	0.2%	1	0.243	242.91	242.91	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
13	0.2%	1	0.231	231.23	231.23	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.2%	79	0.197	2.49	7.28	34.0	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
15	0.2%	1	0.160	159.63	159.63	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 2
SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM posts;
-- rank 3
SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: c1262c71d5af71daeaef9a5e812bbc3a4e14d546
Type: cpu
Time: 2026-10-31 15:05:10 JST
Duration: 60s, Total samples = 14.29s (23.82%)
Showing nodes accounting for 11.35s, 79.43% of 14.29s total
Dropped 574 nodes (cum <= 0.07s)
Showing top 40 nodes out of 188
      flat  flat%   sum%        cum   cum%
     1.41s  9.87%  9.87%      5.49s 38.42%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.29s  9.03% 18.89%      1.29s  9.03%  internal/runtime/syscall/linux.Syscall6
     1.24s  8.68% 27.57%      1.24s  8.68%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.98s  6.86% 34.43%      0.98s  6.86%  aeshashbody
     0.63s  4.41% 38.84%      0.63s  4.41%  runtime.memclrNoHeapPointers
     0.47s  3.29% 42.13%      1.86s 13.02%  runtime.mapaccess1_faststr
     0.46s  3.22% 45.35%      0.60s  4.20%  example.com/isureal-company/isureal.postLess
     0.44s  3.08% 48.43%      0.44s  3.08%  runtime.memmove
     0.33s  2.31% 50.73%      0.43s  3.01%  runtime.tryDeferToSpanScan
     0.32s  2.24% 52.97%      1.23s  8.61%  runtime.scanObject
     0.26s  1.82% 54.79%      0.40s  2.80%  runtime.findObject
     0.24s  1.68% 56.47%      0.33s  2.31%  database/sql.convertAssignRows
     0.21s  1.47% 57.94%      0.22s  1.54%  example.com/isureal-company/isureal.rmLess
     0.21s  1.47% 59.41%      0.21s  1.47%  memeqbody
     0.21s  1.47% 60.88%      1.36s  9.52%  runtime.mallocgc
     0.21s  1.47% 62.35%      0.67s  4.69%  runtime.mallocgcSmallScanNoHeader
     0.20s  1.40% 63.75%      0.24s  1.68%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.18s  1.26% 65.01%      0.18s  1.26%  crypto/internal/fips140/sha256.blockAVX2
     0.18s  1.26% 66.27%      0.18s  1.26%  runtime.nextFreeFast (inline)
     0.17s  1.19% 67.46%      1.26s  8.82%  runtime.mapassign_faststr
     0.17s  1.19% 68.65%      0.17s  1.19%  sync/atomic.(*Int32).Add (inline)
     0.15s  1.05% 69.70%      0.15s  1.05%  runtime.futex
     0.13s  0.91% 70.61%      1.19s  8.33%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.84% 71.45%      0.17s  1.19%  internal/runtime/maps.(*Iter).Next
     0.11s  0.77% 72.22%      0.33s  2.31%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.11s  0.77% 72.99%      2.08s 14.56%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.09s  0.63% 73.62%      0.09s  0.63%  cmpbody
     0.09s  0.63% 74.25%      0.09s  0.63%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
     0.09s  0.63% 74.88%      0.09s  0.63%  internal/runtime/maps.(*ctrlGroup).setEmpty (inline)
     0.09s  0.63% 75.51%      0.34s  2.38%  runtime.mapaccess2_faststr
     0.08s  0.56% 76.07%      0.08s  0.56%  internal/runtime/maps.(*Map).directoryAt (inline)
     0.08s  0.56% 76.63%      0.11s  0.77%  runtime.(*mspan).writeHeapBitsSmall
     0.08s  0.56% 77.19%      0.08s  0.56%  runtime.spanClass.sizeclass (inline)
     0.06s  0.42% 77.61%      2.22s 15.54%  database/sql.(*Rows).nextLocked
     0.06s  0.42% 78.03%      0.28s  1.96%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func12
     0.05s  0.35% 78.38%      3.66s 25.61%  example.com/isureal-company/isureal.loadBig
     0.04s  0.28% 78.66%      0.37s  2.59%  database/sql.(*Rows).scanLocked
     0.04s  0.28% 78.94%      0.30s  2.10%  github.com/jackc/pgx/v5/pgproto3.(*chunkReader).Next
     0.04s  0.28% 79.22%      0.14s  0.98%  runtime.greyobject
     0.03s  0.21% 79.43%      0.22s  1.54%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func2
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 7.9% / 70% | 0% | 0.9% | 4 | 240 MB |  |
| isu2 (2 cpu) | 4.1% / 59% | 0% | 0.0% | 3 | 927 MB |  |
| isu3 (2 cpu) | 4.7% / 69% | 0% | 0.0% | 3 | 779 MB |  |
| isu4 (2 cpu) | 3.6% / 64% | 0% | 1.4% | 3 | 84 MB |  |

どのホストも飽和していない。**ボトルネックはアプリのロジックか DB のクエリ側**にある可能性が高い。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 9% / nginx 3% |
| isu2 | isureal 8% |
| isu3 | isureal 10% |
| isu4 | postgres 4% |

## Diff from previous run

スコア: 140030 → 1315 (-138715)

`a28a701..c416920`

```
 webapp/go/export_api.go        |   2 +
 webapp/go/image_api.go         |  17 +-----
 webapp/go/image_file.go        | 132 +++++++++++++++++++++++++++++++++++++++++
 webapp/go/image_local.go       |  40 ++++++-------
 webapp/go/initialize.go        |   1 +
 webapp/go/post_api.go          |   5 +-
 webapp/go/profile_image_api.go |   5 +-
 webapp/go/realmoji_api.go      |   5 +-
 webapp/go/replica_server.go    |  14 ++++-
 webapp/go/server.go            |   2 +
 webapp/go/slow.go              |  36 +++++++++++
 11 files changed, 218 insertions(+), 41 deletions(-)
```
