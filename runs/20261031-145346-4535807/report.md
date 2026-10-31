# run 20261031-145346-4535807

## Summary

| | |
|---|---|
| score | 133663 |
| status | pass |
| ref | 4535807 |
| sha | 45358075366f |
| 所要時間 | 276s |
| 前回比 | -2381 (-1.8%) vs 20261031-144830-3c8fc33 |
| ベスト比 | -7672 (-5.4%) vs 20261031-141027-bd0fe45 = 141335 |
| note | JPEG の検証を asia2 にも逃がす（main 6ad2e36 の上） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	151	441	1169	1168
europe	7	20	65	185	349	348
america	6	13	35	95	127	126
```

```
{"id":"19942","status":"done","portalStatus":"finished","passed":true,"score":133663,"raw":133663,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:56:30.679Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:56:33.670Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:56:36.668Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:56:39.702Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T05:56:42.669Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T05:56:45.669Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T05:56:48.682Z level=INFO msg="asia 3 日目: 151 人"
time=2026-10-31T05:56:50.237Z level=INFO msg="user moved" user=elorma5695 from=asia to=europe
time=2026-10-31T05:56:50.484Z level=INFO msg="user moved" user=koke-5813 from=asia to=europe
time=2026-10-31T05:56:51.670Z level=INFO msg="europe 3 日目: 65 人"
time=2026-10-31T05:56:54.668Z level=INFO msg="america 3 日目: 35 人"
time=2026-10-31T05:56:57.725Z level=INFO msg="asia 4 日目: 441 人"
time=2026-10-31T05:57:00.794Z level=INFO msg="europe 4 日目: 185 人"
time=2026-10-31T05:57:00.848Z level=INFO msg="user moved" user=yomar-6420 from=europe to=asia
time=2026-10-31T05:57:03.805Z level=INFO msg="america 4 日目: 95 人"
time=2026-10-31T05:57:04.302Z level=INFO msg="user moved" user=koeva7485 from=asia to=america
time=2026-10-31T05:57:04.487Z level=INFO msg="user moved" user=kovin-7136 from=asia to=europe
time=2026-10-31T05:57:04.629Z level=INFO msg="user moved" user=maxni6846 from=asia to=europe
time=2026-10-31T05:57:05.585Z level=INFO msg="user moved" user=anremar-8826 from=asia to=europe
time=2026-10-31T05:57:05.817Z level=INFO msg="user moved" user=isko8039 from=asia to=europe
time=2026-10-31T05:57:06.273Z level=INFO msg="user moved" user=noeva-8630 from=asia to=europe
time=2026-10-31T05:57:06.360Z level=INFO msg="user moved" user=yomar-6420 from=asia to=america
time=2026-10-31T05:57:06.912Z level=INFO msg="asia 5 日目: 1169 人"
time=2026-10-31T05:57:09.826Z level=INFO msg="europe 5 日目: 349 人"
time=2026-10-31T05:57:14.456Z level=INFO msg="america 5 日目: 127 人"
time=2026-10-31T05:57:25.561Z level=INFO msg="asia 最終日（グランプリ）: 1168 人"
time=2026-10-31T05:57:25.807Z level=INFO msg="america 最終日（グランプリ）: 126 人"
time=2026-10-31T05:57:25.808Z level=INFO msg="europe 最終日（グランプリ）: 348 人"
time=2026-10-31T05:57:30.816Z level=INFO msg="グランプリ: アジア 18 人, ヨーロッパ 151 人, アメリカ 121 人"
{
  "pass": true,
  "score": 133663,
  "messages": [
    "maleo5061: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "maevawa-5100: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "yasu.5125: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.7%	1832	65	0	GET	/api/notifications	17987.660	9.819	19.158	19.594	151
9.0%	15590	394	0	GET	^/api/images/[^/]+	3902.298	0.250	3.066	6.505	109293
8.6%	5331	64	0	POST	/api/friend-requests	3703.472	0.695	3.698	4.393	114
5.9%	1796	34	0	POST	/api/posts	2545.327	1.417	8.636	9.778	630
4.7%	9884	0	0	GET	^/api/feeds/[^/]+	2048.782	0.207	2.930	4.016	3199
4.3%	6066	763	0	POST	/api/auth/login	1866.099	0.308	2.521	4.016	215
4.1%	738	120	0	PUT	/api/me/profile-image	1773.418	2.403	9.670	11.799	3
3.6%	3316	2	0	GET	/api/posting-window/current	1542.336	0.465	3.437	4.016	236
3.0%	3639	0	0	GET	/api/memories	1301.236	0.358	3.220	8.890	13159
2.3%	1894	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1010.073	0.533	3.505	5.365	116
2.3%	2624	0	0	GET	/api/users	980.054	0.373	3.505	4.016	236
1.4%	1555	90	0	GET	^/api/posts/me/current$	588.026	0.378	4.016	4.016	606
1.2%	716	0	0	PUT	^/api/follows/[^/]+$	512.327	0.716	3.505	4.016	0
1.1%	1193	1	0	POST	/api/auth/register	471.802	0.395	1.334	3.013	227
0.9%	427	8	0	POST	^/api/posts/[^/]+/realmojis$	373.327	0.874	6.588	8.976	471
0.8%	312	0	0	GET	/api/me/export	344.465	1.104	8.336	11.436	2016605
0.8%	537	0	0	PATCH	/api/me	333.952	0.622	3.024	4.016	277
0.6%	1602	40	0	GET	^/api/posts/[^/]+$	252.534	0.158	1.943	4.016	609
0.6%	627	3	0	GET	/api/me	249.938	0.399	2.119	3.505	301
0.6%	885	0	0	POST	/api/auth/logout	238.713	0.270	2.922	4.393	0
0.5%	511	0	0	POST	^/api/posts/[^/]+/comments$	214.080	0.419	3.024	4.393	334
0.5%	1332	0	0	GET	^/api/relationships/[^/]+	202.090	0.152	2.417	4.016	1012
0.3%	790	0	0	GET	^/api/memories/[^/]+	144.769	0.183	2.123	3.437	622
0.3%	166	20	0	PATCH	^/api/posts/[^/]+$	137.821	0.830	4.016	4.016	560
0.2%	737	0	0	GET	^/api/users/[^/]+$	101.701	0.138	1.943	4.016	192
0.2%	71	0	0	DELETE	^/api/posts/[^/]+$	82.689	1.165	5.125	5.125	0
0.2%	621	0	0	GET	^/api/posts/[^/]+/comments$	76.588	0.123	1.943	2.222	1704
0.1%	673	0	0	GET	^/api/posts/[^/]+/realmojis$	62.440	0.093	1.120	3.505	1122
0.1%	68	0	0	DELETE	^/api/realmojis/[^/]+$	55.726	0.820	5.285	5.285	0
0.1%	94	0	0	DELETE	^/api/comments/[^/]+$	33.517	0.357	3.505	3.505	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.1%	634	1	0	GET	/api/notifications	5719.111	9.021	16.278	16.652	174
17.9%	3432	22	0	GET	^/api/images/[^/]+	2688.010	0.783	4.953	7.237	186162
11.8%	916	9	0	POST	/api/posts	1775.022	1.938	6.248	6.943	645
6.8%	1795	14	0	POST	/api/friend-requests	1018.992	0.568	3.037	3.557	114
4.1%	2930	0	0	GET	^/api/feeds/[^/]+	622.059	0.212	1.331	1.872	2440
3.8%	1019	0	0	POST	/api/auth/login	564.139	0.554	2.277	2.970	245
2.0%	341	0	0	POST	/api/auth/register	303.313	0.889	2.463	2.605	228
2.0%	188	0	0	PUT	/api/me/profile-image	299.353	1.592	8.755	8.762	0
1.9%	1288	0	0	GET	/api/posting-window/current	279.525	0.217	1.367	1.584	239
1.8%	489	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	263.343	0.539	2.841	3.137	116
1.6%	902	0	0	GET	/api/users	239.236	0.265	1.516	2.132	233
1.5%	959	0	0	GET	/api/memories	230.862	0.241	1.936	3.371	12412
1.1%	738	40	0	GET	^/api/posts/me/current$	159.727	0.216	1.343	1.705	615
0.7%	217	0	0	PUT	^/api/follows/[^/]+$	106.666	0.492	3.062	3.137	0
0.7%	150	0	0	PATCH	/api/me	103.015	0.687	2.409	2.435	277
0.6%	394	0	0	GET	^/api/relationships/[^/]+	92.296	0.234	1.551	2.107	913
0.4%	96	3	0	POST	^/api/posts/[^/]+/realmojis$	65.040	0.677	5.517	5.517	452
0.4%	194	2	0	GET	/api/me	62.177	0.320	2.049	2.095	301
0.4%	126	0	0	POST	/api/auth/logout	62.105	0.493	2.270	2.330	0
0.4%	69	9	0	PATCH	^/api/posts/[^/]+$	59.919	0.868	2.470	2.470	571
0.4%	348	15	0	GET	^/api/posts/[^/]+$	54.204	0.156	0.974	1.155	603
0.3%	143	0	0	POST	^/api/posts/[^/]+/comments$	46.056	0.322	1.442	2.051	337
0.3%	29	0	0	DELETE	^/api/posts/[^/]+$	37.816	1.304	3.149	3.149	0
0.2%	58	0	0	GET	/api/me/export	33.035	0.570	10.801	10.801	2508938
0.2%	195	0	0	GET	^/api/users/[^/]+$	32.185	0.165	1.406	1.590	192
0.2%	119	0	0	GET	^/api/memories/[^/]+	25.360	0.213	1.981	2.125	893
0.1%	98	2	0	GET	^/api/posts/[^/]+/comments$	14.360	0.147	1.218	1.218	1385
0.1%	13	0	0	DELETE	^/api/realmojis/[^/]+$	13.930	1.072	3.112	3.112	0
0.1%	92	0	0	GET	^/api/posts/[^/]+/realmojis$	11.283	0.123	0.980	0.980	992
0.0%	24	0	0	DELETE	^/api/comments/[^/]+$	6.030	0.251	0.714	0.714	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.6%	266	1	0	GET	/api/notifications	2244.164	8.437	12.722	12.754	178
18.1%	2329	9	0	GET	^/api/images/[^/]+	977.341	0.420	2.274	3.031	156530
7.1%	822	13	0	POST	/api/friend-requests	384.811	0.468	3.028	3.599	114
7.1%	359	0	0	POST	/api/posts	383.290	1.068	4.466	4.530	647
4.7%	639	0	0	POST	/api/auth/login	255.495	0.400	3.061	3.118	252
4.6%	1873	0	0	GET	^/api/feeds/[^/]+	250.511	0.134	0.830	1.431	3311
2.3%	309	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	125.195	0.405	2.746	3.257	116
1.9%	618	0	0	GET	/api/memories	102.916	0.167	2.078	3.336	15574
1.4%	493	0	0	GET	/api/posting-window/current	76.515	0.155	0.853	1.127	239
1.2%	406	0	0	GET	/api/users	66.787	0.164	0.975	1.058	234
1.2%	62	5	0	POST	^/api/posts/[^/]+/realmojis$	64.846	1.046	8.314	8.314	401
1.1%	124	0	0	POST	/api/auth/register	59.706	0.481	3.013	3.124	230
1.1%	357	5	0	GET	^/api/posts/me/current$	57.291	0.160	1.581	1.749	638
0.8%	268	0	0	GET	^/api/relationships/[^/]+	45.580	0.170	0.922	1.240	869
0.8%	111	0	0	POST	/api/auth/logout	44.773	0.403	3.021	3.454	0
0.7%	65	0	0	PUT	/api/me/profile-image	39.083	0.601	3.650	3.650	0
0.7%	74	0	0	PUT	^/api/follows/[^/]+$	37.140	0.502	2.547	2.547	0
0.6%	269	7	0	GET	^/api/posts/[^/]+$	30.671	0.114	0.793	1.238	610
0.4%	49	0	0	PATCH	/api/me	24.087	0.492	2.750	2.750	284
0.3%	38	0	0	POST	^/api/posts/[^/]+/comments$	17.166	0.452	2.525	2.525	348
0.3%	142	0	0	GET	^/api/users/[^/]+$	15.988	0.113	0.743	0.835	198
0.3%	96	0	0	GET	^/api/memories/[^/]+	15.088	0.157	1.546	1.546	961
0.3%	37	0	0	PATCH	^/api/posts/[^/]+$	13.893	0.375	1.337	1.337	634
0.2%	48	0	0	GET	/api/me/export	11.765	0.245	2.107	2.107	3481525
0.2%	70	1	0	GET	/api/me	10.978	0.157	1.137	1.137	296
0.2%	13	0	0	DELETE	^/api/realmojis/[^/]+$	9.668	0.744	2.076	2.076	0
0.2%	12	0	0	DELETE	^/api/posts/[^/]+$	8.611	0.718	1.234	1.234	0
0.1%	59	0	0	GET	^/api/posts/[^/]+/realmojis$	6.145	0.104	0.487	0.487	1161
0.1%	38	0	0	GET	^/api/posts/[^/]+/comments$	4.087	0.108	0.721	0.721	1642
0.1%	10	0	0	DELETE	^/api/comments/[^/]+$	3.252	0.325	0.920	0.920	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 13469 クエリ / 48 種類 / 合計 133.117s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	50.1%	320	66.697	208.43	1013.11	1288.0	131.6	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	16.0%	3	21.247	7082.40	9504.03	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	12.0%	3	15.915	5304.91	7199.72	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	10.3%	11231	13.727	1.22	6.61	9.5	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	5.9%	3	7.851	2616.85	3555.08	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.9%	6	3.850	641.60	2976.78	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.9%	1	1.233	1233.21	1233.21	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.5%	6	0.675	112.58	157.50	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.5%	6	0.664	110.73	157.32	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.2%	1	0.253	253.46	253.46	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.217	217.27	217.27	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.157	157.19	157.19	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.107	107.47	107.47	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.081	81.18	81.18	2470.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.1%	1	0.078	78.30	78.30	1525.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
File: isureal
Build ID: fb7838af5f90f3801d0265d34f1d74629618a7d1
Type: cpu
Time: 2026-10-31 14:56:09 JST
Duration: 60.15s, Total samples = 44.33s (73.70%)
Showing nodes accounting for 31.81s, 71.76% of 44.33s total
Dropped 1175 nodes (cum <= 0.22s)
Showing top 40 nodes out of 281
      flat  flat%   sum%        cum   cum%
     6.39s 14.41% 14.41%      6.39s 14.41%  crypto/internal/fips140/sha256.blockAVX2
     5.36s 12.09% 26.51%      5.36s 12.09%  internal/runtime/syscall/linux.Syscall6
     1.87s  4.22% 30.72%      1.87s  4.22%  runtime.memmove
     1.72s  3.88% 34.60%      1.72s  3.88%  runtime.memclrNoHeapPointers
     1.71s  3.86% 38.46%      6.44s 14.53%  runtime.scanObject
     1.69s  3.81% 42.27%      5.73s 12.93%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.63s  3.68% 45.95%      2.11s  4.76%  runtime.findObject
     1.47s  3.32% 49.27%      1.47s  3.32%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.34s  3.02% 52.29%      1.34s  3.02%  aeshashbody
     1.15s  2.59% 54.88%      2.27s  5.12%  runtime.tryDeferToSpanScan
     0.63s  1.42% 56.30%      0.63s  1.42%  runtime.spanClass.sizeclass (inline)
     0.50s  1.13% 57.43%      2.14s  4.83%  runtime.mapaccess1_faststr
     0.43s  0.97% 58.40%      0.43s  0.97%  memeqbody
     0.35s  0.79% 59.19%      0.43s  0.97%  runtime.heapArenaOf (inline)
     0.34s  0.77% 59.96%      6.26s 14.12%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.34s  0.77% 60.73%      0.74s  1.67%  database/sql.convertAssignRows
     0.34s  0.77% 61.49%      0.34s  0.77%  runtime.nextFreeFast (inline)
     0.32s  0.72% 62.22%      0.52s  1.17%  example.com/isureal-company/isureal.postLess
     0.31s   0.7% 62.91%      6.95s 15.68%  crypto/internal/fips140/sha256.(*Digest).Write
     0.29s  0.65% 63.57%      1.99s  4.49%  runtime.mapassign_faststr
     0.27s  0.61% 64.18%      0.29s  0.65%  runtime.(*mspan).writeHeapBitsSmall
     0.26s  0.59% 64.76%      0.26s  0.59%  sync/atomic.(*Int32).Add
     0.23s  0.52% 65.28%      0.23s  0.52%  runtime.futex
     0.22s   0.5% 65.78%      0.49s  1.11%  runtime.greyobject
     0.21s  0.47% 66.25%      0.45s  1.02%  compress/flate.(*compressor).deflate
     0.20s  0.45% 66.70%      0.23s  0.52%  encoding/json.appendString[go.shape.string]
     0.20s  0.45% 67.16%      3.06s  6.90%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.20s  0.45% 67.61%      4.24s  9.56%  runtime.mallocgc
     0.20s  0.45% 68.06%      0.48s  1.08%  runtime.scanObjectsSmall
     0.20s  0.45% 68.51%      0.27s  0.61%  runtime.typePointers.next
     0.19s  0.43% 68.94%      6.58s 14.84%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.19s  0.43% 69.37%      1.11s  2.50%  runtime.mallocgcSmallScanNoHeader
     0.17s  0.38% 69.75%         7s 15.79%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.17s  0.38% 70.13%      4.79s 10.81%  example.com/isureal-company/isureal.listMemories
     0.16s  0.36% 70.49%      0.34s  0.77%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.13s  0.29% 70.79%      1.13s  2.55%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.27% 71.06%      0.46s  1.04%  runtime.mapaccess2_faststr
     0.11s  0.25% 71.31%      0.25s  0.56%  internal/runtime/maps.(*Iter).Next
     0.10s  0.23% 71.53%      0.31s   0.7%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
     0.10s  0.23% 71.76%      0.86s  1.94%  runtime.mallocgcSmallNoscan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 41.8% / 100% | 28% | 5.4% | 6 | 157 MB | run queue が積もっている |
| isu2 (2 cpu) | 23.2% / 100% | 4% | 0.6% | 6 | 244 MB | run queue が積もっている |
| isu3 (2 cpu) | 19.1% / 90% | 1% | 0.5% | 3 | 662 MB |  |
| isu4 (2 cpu) | 30.4% / 99% | 7% | 12.8% | 5 | 106 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 58% / nginx 9% |
| isu2 | isureal 43% / nginx 1% |
| isu3 | isureal 36% / nginx 1% |
| isu4 | isureal 22% / postgres 12% |

## Diff from previous run

スコア: 136044 → 133663 (-2381)

`3c8fc33..4535807`

```
 conf/isu1/home/isucon/env.sh | 2 +-
 webapp/go/offload.go         | 2 +-
 2 files changed, 2 insertions(+), 2 deletions(-)
```
