# run 20261031-165402-07a4cc3

## Summary

| | |
|---|---|
| score | 1561 |
| status | pass |
| ref | slow-before-1700-v4 |
| sha | 07a4cc334e28 |
| 所要時間 | 300s |
| 前回比 | -230549 (-99.3%) vs 20261031-164055-bdf05cd |
| ベスト比 | -230549 (-99.3%) vs 20261031-164055-bdf05cd = 232110 |
| note | low version before 17:00 (3000ms wait during fresh posting windows) |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	18	18	18	18	18
europe	7	7	7	7	7	7
america	6	6	6	6	6	6
```

```
{"id":"30967","status":"done","portalStatus":"finished","passed":true,"score":1561,"raw":1561,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:54:36.297Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:54:39.287Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:54:42.287Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:54:45.299Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T07:54:48.286Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T07:54:51.287Z level=INFO msg="america 2 日目: 6 人"
time=2026-10-31T07:54:54.297Z level=INFO msg="asia 3 日目: 18 人"
time=2026-10-31T07:54:57.287Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T07:55:00.285Z level=INFO msg="america 3 日目: 6 人"
time=2026-10-31T07:55:03.377Z level=INFO msg="asia 4 日目: 18 人"
time=2026-10-31T07:55:06.299Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T07:55:09.287Z level=INFO msg="america 4 日目: 6 人"
time=2026-10-31T07:55:12.297Z level=INFO msg="asia 5 日目: 18 人"
time=2026-10-31T07:55:15.287Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T07:55:18.286Z level=INFO msg="america 5 日目: 6 人"
time=2026-10-31T07:55:27.287Z level=INFO msg="europe 最終日（グランプリ）: 7 人"
time=2026-10-31T07:55:27.287Z level=INFO msg="america 最終日（グランプリ）: 6 人"
time=2026-10-31T07:55:27.298Z level=INFO msg="asia 最終日（グランプリ）: 18 人"
time=2026-10-31T07:55:36.289Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 2 人, アメリカ 1 人"
{
  "pass": true,
  "score": 1561,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
89.3%	12693	12617	0	POST	/api/auth/login	32613.998	2.569	3.002	3.028	15
5.6%	126	0	0	GET	/api/notifications	2055.805	16.316	63.001	63.001	153
1.1%	172	171	0	PUT	/api/me/profile-image	418.312	2.432	4.417	4.453	16
0.8%	101	0	0	GET	/api/memories	286.051	2.832	3.043	3.048	5405
0.5%	64	0	0	GET	/api/posting-window/current	192.387	3.006	3.009	3.009	237
0.4%	61	3	0	POST	/api/friend-requests	162.032	2.656	3.002	3.002	111
0.4%	54	0	0	GET	^/api/feeds/[^/]+	159.307	2.950	3.009	3.009	114
0.4%	51	1	0	POST	/api/posts	157.453	3.087	3.622	3.622	595
0.4%	47	44	0	GET	^/api/posts/me/current$	138.297	2.942	3.009	3.009	63
0.3%	41	0	0	GET	^/api/relationships/[^/]+	102.525	2.501	3.028	3.028	4203
0.2%	22	0	0	GET	/api/users	66.131	3.006	3.011	3.011	250
0.2%	34	1	0	POST	/api/auth/register	66.038	1.942	3.002	3.002	237
0.1%	17	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	30.011	1.765	3.001	3.001	116
0.1%	1	0	0	POST	/api/initialize	21.069	21.069	21.069	21.069	90
0.0%	7	0	0	GET	/api/me/export	17.441	2.492	3.821	3.821	2120208
0.0%	6	0	0	POST	/api/auth/logout	15.004	2.501	3.001	3.001	0
0.0%	4	3	0	PATCH	^/api/posts/[^/]+$	12.003	3.001	3.001	3.001	173
0.0%	7	2	0	GET	/api/me	9.032	1.290	3.007	3.007	161
0.0%	3	0	0	GET	^/api/memories/[^/]+	9.021	3.007	3.011	3.011	879
0.0%	1	0	0	DELETE	^/api/posts/[^/]+$	3.102	3.102	3.102	3.102	0
0.0%	348	179	0	GET	^/api/images/[^/]+	0.630	0.002	0.024	0.034	79225
0.0%	18	0	0	POST	/api/admin/posting-window	0.088	0.005	0.014	0.014	176
0.0%	2	0	0	PUT	^/api/blocks/[^/]+$	0.001	0.001	0.001	0.001	0
0.0%	4	0	0	PUT	^/api/admin/users/[^/]+/certification$	0.000	0.000	0.000	0.000	0
0.0%	8	0	0	PUT	^/api/follows/[^/]+$	0.000	0.000	0.000	0.000	0
0.0%	1	0	0	POST	/api/admin/auth/login	0.000	0.000	0.000	0.000	106
0.0%	1	1	0	DELETE	^/api/admin/users/[^/]+/certification$	0.000	0.000	0.000	0.000	39
0.0%	1	0	0	POST	/api/admin/auth/logout	0.000	0.000	0.000	0.000	0

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
56.8%	49	0	0	GET	/api/notifications	800.430	16.335	63.103	63.103	155
9.0%	45	0	0	GET	/api/memories	127.019	2.823	3.106	3.106	15512
5.0%	23	0	0	GET	/api/posting-window/current	70.572	3.068	3.102	3.102	239
4.6%	21	0	0	GET	^/api/feeds/[^/]+	64.232	3.059	3.102	3.102	83
4.2%	19	0	0	POST	/api/posts	59.841	3.150	3.214	3.214	610
4.2%	22	1	0	POST	/api/friend-requests	59.248	2.693	3.107	3.107	111
3.9%	18	12	0	GET	^/api/posts/me/current$	55.020	3.057	3.058	3.058	214
3.7%	19	0	0	GET	^/api/relationships/[^/]+	52.177	2.746	3.059	3.059	844
2.8%	13	0	0	GET	/api/users	39.739	3.057	3.059	3.059	208
2.1%	18	0	0	POST	/api/auth/login	29.286	1.627	3.103	3.103	251
1.1%	6	0	0	GET	/api/me/export	14.961	2.494	3.095	3.095	2667455
0.7%	4	1	0	GET	/api/me	9.376	2.344	3.058	3.058	150
0.5%	5	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	6.629	1.326	3.102	3.102	116
0.5%	68	1	0	GET	^/api/images/[^/]+	6.554	0.096	0.206	0.206	177277
0.5%	2	0	0	DELETE	^/api/posts/[^/]+$	6.409	3.204	3.206	3.206	0
0.4%	2	0	0	POST	/api/auth/logout	6.204	3.102	3.102	3.102	0
0.1%	5	0	0	POST	/api/auth/register	0.978	0.196	0.217	0.217	228
0.0%	4	0	0	PUT	^/api/follows/[^/]+$	0.405	0.101	0.102	0.102	0
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.102	0.102	0.102	0.102	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
56.9%	42	0	0	GET	/api/notifications	685.853	16.330	63.103	63.103	156
9.2%	38	0	0	GET	/api/memories	110.471	2.907	3.144	3.144	5851
5.7%	25	0	0	POST	/api/friend-requests	68.548	2.742	3.103	3.103	115
4.6%	18	0	0	GET	/api/posting-window/current	55.024	3.057	3.058	3.058	240
3.9%	15	0	0	POST	/api/posts	47.312	3.154	3.219	3.219	581
3.8%	15	0	0	GET	^/api/feeds/[^/]+	45.842	3.056	3.059	3.059	29
3.6%	15	0	0	GET	^/api/relationships/[^/]+	42.912	2.861	3.073	3.073	621
3.5%	14	13	0	GET	^/api/posts/me/current$	42.797	3.057	3.065	3.065	69
3.0%	12	0	0	GET	/api/users	36.672	3.056	3.058	3.058	189
2.6%	17	0	0	POST	/api/auth/login	31.930	1.878	3.102	3.102	249
0.9%	4	0	0	GET	/api/me/export	10.928	2.732	3.146	3.146	4162862
0.8%	3	0	0	GET	/api/me	9.170	3.057	3.058	3.058	194
0.4%	57	1	0	GET	^/api/images/[^/]+	4.322	0.076	0.206	0.206	210563
0.3%	1	0	0	DELETE	^/api/posts/[^/]+$	3.202	3.202	3.202	3.202	0
0.3%	1	1	0	PATCH	^/api/posts/[^/]+$	3.102	3.102	3.102	3.102	30
0.3%	1	0	0	POST	/api/auth/logout	3.102	3.102	3.102	3.102	0
0.3%	1	0	0	GET	^/api/memories/[^/]+	3.053	3.053	3.053	3.053	974
0.1%	5	0	0	POST	/api/auth/register	1.013	0.203	0.203	0.203	231
0.0%	4	0	0	PUT	^/api/follows/[^/]+$	0.404	0.101	0.101	0.101	0
0.0%	3	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	0.305	0.102	0.102	0.102	116
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.101	0.101	0.101	0.101	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 9035 クエリ / 48 種類 / 合計 66.771s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	39.2%	4	26.150	6537.44	10080.33	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
2	28.3%	4	18.884	4720.98	7185.66	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
3	13.7%	4	9.179	2294.78	3437.49	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
4	8.2%	7	5.484	783.43	2986.26	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
5	5.1%	10	3.432	343.19	717.24	1880.3	184.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
6	1.0%	7	0.638	91.20	159.48	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
7	1.0%	7	0.636	90.79	156.09	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
8	0.7%	1	0.449	449.26	449.26	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
9	0.5%	1	0.350	350.32	350.32	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.4%	1	0.248	247.94	247.94	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
11	0.3%	1	0.226	226.08	226.08	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.2%	1	0.165	164.91	164.91	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.2%	8668	0.156	0.02	0.05	0.9	17.3	SELECT ip, banned_until FROM ip_bans WHERE banned_until > clock_timestamp()
14	0.2%	115	0.134	1.17	3.14	14.2	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
15	0.2%	1	0.108	107.72	107.72	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: fca71ccc778fb2a5e1b02960c444ea7340bc30c3
Type: cpu
Time: 2026-10-31 16:54:08 JST
Duration: 60s, Total samples = 15.74s (26.23%)
Showing nodes accounting for 12.47s, 79.22% of 15.74s total
Dropped 493 nodes (cum <= 0.08s)
Showing top 40 nodes out of 219
      flat  flat%   sum%        cum   cum%
     1.63s 10.36% 10.36%      5.60s 35.58%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.37s  8.70% 19.06%      1.37s  8.70%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.37s  8.70% 27.76%      1.37s  8.70%  internal/runtime/syscall/linux.Syscall6
     0.98s  6.23% 33.99%      0.98s  6.23%  aeshashbody
     0.80s  5.08% 39.07%      0.80s  5.08%  runtime.memmove
     0.58s  3.68% 42.76%      0.58s  3.68%  runtime.memclrNoHeapPointers
     0.54s  3.43% 46.19%      1.88s 11.94%  runtime.scanObject
     0.45s  2.86% 49.05%      0.64s  4.07%  runtime.tryDeferToSpanScan
     0.40s  2.54% 51.59%      0.48s  3.05%  runtime.findObject
     0.36s  2.29% 53.88%      1.61s 10.23%  runtime.mapaccess1_faststr
     0.32s  2.03% 55.91%      0.53s  3.37%  example.com/isureal-company/isureal.postLess
     0.30s  1.91% 57.81%      0.30s  1.91%  runtime.nextFreeFast (inline)
     0.28s  1.78% 59.59%      0.43s  2.73%  database/sql.convertAssignRows
     0.27s  1.72% 61.31%      0.27s  1.72%  runtime.futex
     0.20s  1.27% 62.58%      0.20s  1.27%  memeqbody
     0.18s  1.14% 63.72%      0.20s  1.27%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.17s  1.08% 64.80%      1.42s  9.02%  runtime.mallocgc
     0.17s  1.08% 65.88%      0.75s  4.76%  runtime.mallocgcSmallScanNoHeader
     0.16s  1.02% 66.90%      1.49s  9.47%  runtime.mapassign_faststr
     0.15s  0.95% 67.85%      0.15s  0.95%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
     0.12s  0.76% 68.61%      0.21s  1.33%  runtime.greyobject
     0.12s  0.76% 69.38%      0.12s  0.76%  runtime.spanClass.sizeclass (inline)
     0.11s   0.7% 70.08%      0.11s   0.7%  crypto/internal/fips140/sha256.blockAVX2
     0.11s   0.7% 70.78%      2.02s 12.83%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s   0.7% 71.47%      0.32s  2.03%  runtime.mapaccess2_faststr
     0.11s   0.7% 72.17%      0.11s   0.7%  sync/atomic.(*Int32).Add (inline)
     0.10s  0.64% 72.81%      0.20s  1.27%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.10s  0.64% 73.44%      0.10s  0.64%  example.com/isureal-company/isureal.rmLess
     0.09s  0.57% 74.02%      0.09s  0.57%  cmpbody
     0.09s  0.57% 74.59%      0.52s  3.30%  database/sql.(*Rows).scanLocked
     0.09s  0.57% 75.16%      0.13s  0.83%  internal/runtime/maps.(*Iter).Next
     0.09s  0.57% 75.73%      0.09s  0.57%  runtime.acquirem (inline)
     0.08s  0.51% 76.24%      3.85s 24.46%  example.com/isureal-company/isureal.loadBig
     0.08s  0.51% 76.75%      1.19s  7.56%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.08s  0.51% 77.26%      0.08s  0.51%  runtime.(*spanInlineMarkBits).init
     0.08s  0.51% 77.76%      0.08s  0.51%  time.Time.Before
     0.06s  0.38% 78.14%      0.55s  3.49%  runtime.convTstring
     0.06s  0.38% 78.53%      0.73s  4.64%  runtime.growslice
     0.06s  0.38% 78.91%      0.38s  2.41%  runtime.mallocgcSmallNoscan
     0.05s  0.32% 79.22%      0.08s  0.51%  runtime.(*mspan).markBitsForIndex (inline)
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 8.2% / 63% | 0% | 1.0% | 4 | 191 MB |  |
| isu2 (2 cpu) | 3.4% / 52% | 0% | 0.0% | 3 | 1732 MB |  |
| isu3 (2 cpu) | 3.3% / 51% | 0% | 0.0% | 3 | 1718 MB |  |
| isu4 (2 cpu) | 7.4% / 100% | 3% | 0.9% | 5 | 99 MB | run queue が積もっている |

**isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 10% / nginx 3% |
| isu2 | isureal 6% |
| isu3 | isureal 7% |
| isu4 | isureal 8% / postgres 4% |

## Diff from previous run

スコア: 232110 → 1561 (-230549)

`bdf05cd..07a4cc3`

```
 webapp/go/replica_server.go |  1 +
 webapp/go/server.go         |  1 +
 webapp/go/slow.go           | 36 ++++++++++++++++++++++++++++++++++++
 3 files changed, 38 insertions(+)
```
