# run 20261031-111050-93a8224

## Summary

| | |
|---|---|
| score | 24028 |
| status | pass |
| ref | offload-cpu |
| sha | 93a8224f0122 |
| 所要時間 | 588s |
| 前回比 | +8360 (53.4%) vs 20261031-110450-8f720c2 |
| ベスト比 | -773 (-3.1%) vs 20261031-103811-cef5fe3 = 24801 |
| note | offload-jpeg-sign-to-eu-us |

## Bench

```
{"id":"3443","status":"done","portalStatus":"finished","passed":true,"score":24028,"raw":24028,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:19:27.167Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:19:28.340Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:19:31.349Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:19:35.878Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T02:19:37.332Z level=INFO msg="europe 2 日目: 9 人"
time=2026-10-31T02:19:40.331Z level=INFO msg="america 2 日目: 12 人"
time=2026-10-31T02:19:44.108Z level=INFO msg="asia 3 日目: 21 人"
time=2026-10-31T02:19:46.166Z level=INFO msg="user moved" user=ivymu.5183 from=america to=asia
time=2026-10-31T02:19:46.329Z level=INFO msg="europe 3 日目: 13 人"
time=2026-10-31T02:19:46.641Z level=INFO msg="user moved" user=ivymu.5183 from=asia to=europe
time=2026-10-31T02:19:49.360Z level=INFO msg="america 3 日目: 26 人"
time=2026-10-31T02:19:53.014Z level=INFO msg="asia 4 日目: 39 人"
time=2026-10-31T02:19:55.352Z level=INFO msg="europe 4 日目: 28 人"
time=2026-10-31T02:19:58.353Z level=INFO msg="america 4 日目: 45 人"
time=2026-10-31T02:20:03.894Z level=INFO msg="asia 5 日目: 60 人"
time=2026-10-31T02:20:03.940Z level=INFO msg="user moved" user=ivymu.5183 from=europe to=asia
time=2026-10-31T02:20:04.672Z level=INFO msg="europe 5 日目: 27 人"
time=2026-10-31T02:20:07.489Z level=INFO msg="america 5 日目: 47 人"
time=2026-10-31T02:20:16.930Z level=INFO msg="america 最終日（グランプリ）: 47 人"
time=2026-10-31T02:20:19.027Z level=INFO msg="asia 最終日（グランプリ）: 60 人"
time=2026-10-31T02:20:19.621Z level=INFO msg="europe 最終日（グランプリ）: 26 人"
time=2026-10-31T02:20:25.312Z level=INFO msg="グランプリ: アジア 6 人, ヨーロッパ 0 人, アメリカ 34 人"
{
  "pass": true,
  "score": 24028,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
52.9%	178	0	0	GET	/api/notifications	1761.043	9.894	15.934	16.223	175
12.9%	206	7	0	POST	/api/posts	430.519	2.090	6.897	6.987	608
10.3%	187	160	0	PUT	/api/me/profile-image	343.080	1.835	7.301	7.712	24
7.5%	1707	695	0	POST	/api/auth/login	250.633	0.147	1.280	3.999	158
5.4%	3380	240	0	GET	^/api/images/[^/]+	180.580	0.053	0.712	3.570	74175
3.1%	806	0	0	GET	/api/memories	104.492	0.130	0.739	2.208	13695
1.6%	74	0	0	GET	/api/me/export	51.806	0.700	3.757	3.757	1841675
0.9%	72	0	0	POST	^/api/posts/[^/]+/realmojis$	30.882	0.429	2.640	2.640	465
0.9%	1496	0	0	GET	^/api/feeds/[^/]+	29.339	0.020	0.270	1.414	2757
0.7%	259	0	0	GET	/api/posting-window/current	22.585	0.087	1.041	1.330	237
0.5%	76	1	0	POST	/api/auth/register	16.440	0.216	1.832	1.832	231
0.4%	18	0	0	POST	/api/admin/posting-window	14.867	0.826	3.307	3.307	176
0.4%	330	17	0	POST	/api/friend-requests	13.263	0.040	0.768	1.822	111
0.4%	205	0	0	POST	/api/auth/logout	13.063	0.064	1.161	2.354	0
0.4%	1	0	0	POST	/api/initialize	12.206	12.206	12.206	12.206	90
0.3%	299	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	11.373	0.038	0.494	0.541	116
0.3%	189	0	0	GET	^/api/memories/[^/]+	9.503	0.050	0.392	0.396	586
0.2%	153	1	0	GET	^/api/posts/me/current$	5.974	0.039	0.572	1.619	622
0.1%	220	0	0	GET	^/api/relationships/[^/]+	3.792	0.017	0.119	0.337	2195
0.1%	315	3	0	GET	^/api/posts/[^/]+$	3.324	0.011	0.103	0.324	597
0.1%	144	0	0	GET	/api/users	2.191	0.015	0.161	0.237	251
0.1%	155	0	0	GET	^/api/posts/[^/]+/realmojis$	2.048	0.013	0.206	0.215	217
0.1%	68	0	0	POST	^/api/posts/[^/]+/comments$	1.914	0.028	0.271	0.271	349
0.1%	158	0	0	GET	^/api/users/[^/]+$	1.867	0.012	0.219	0.255	184
0.1%	46	0	0	PUT	^/api/follows/[^/]+$	1.687	0.037	0.436	0.436	0
0.1%	123	0	0	GET	^/api/posts/[^/]+/comments$	1.666	0.014	0.139	0.145	615
0.0%	22	0	0	PATCH	^/api/posts/[^/]+$	1.500	0.068	0.357	0.357	622
0.0%	19	0	0	PATCH	/api/me	1.294	0.068	0.394	0.394	290
0.0%	1	0	0	POST	/api/admin/auth/logout	1.079	1.079	1.079	1.079	0
0.0%	34	3	0	GET	/api/me	0.968	0.028	0.460	0.460	271
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
48.7%	92	0	0	GET	/api/notifications	867.940	9.434	14.933	14.933	177
14.9%	103	2	0	POST	/api/posts	265.900	2.582	6.165	6.651	636
7.2%	703	6	0	GET	^/api/images/[^/]+	128.195	0.182	0.950	1.800	157743
5.1%	546	0	0	GET	^/api/feeds/[^/]+	91.486	0.168	0.430	0.643	2961
4.2%	183	0	0	POST	/api/auth/login	75.580	0.413	4.042	4.276	252
3.5%	211	0	0	GET	/api/memories	62.856	0.298	1.121	3.633	12686
2.5%	126	0	0	GET	/api/posting-window/current	44.410	0.352	2.142	2.451	239
1.9%	143	7	0	POST	/api/friend-requests	33.695	0.236	2.301	2.314	111
1.4%	40	0	0	POST	^/api/posts/[^/]+/realmojis$	24.624	0.616	2.167	2.167	442
1.3%	38	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	23.849	0.628	2.459	2.459	116
0.9%	90	2	0	GET	^/api/posts/[^/]+$	16.520	0.184	0.759	0.759	601
0.9%	76	0	0	GET	^/api/relationships/[^/]+	15.496	0.204	1.209	1.209	922
0.9%	10	0	0	GET	/api/me/export	15.432	1.543	3.219	3.219	2341474
0.8%	78	2	0	GET	^/api/posts/me/current$	13.942	0.179	0.551	0.551	628
0.8%	27	0	0	POST	^/api/posts/[^/]+/comments$	13.740	0.509	2.175	2.175	366
0.7%	34	0	0	POST	/api/auth/logout	12.631	0.371	4.096	4.096	0
0.6%	66	0	0	GET	^/api/users/[^/]+$	10.655	0.161	0.416	0.416	181
0.6%	24	0	0	PUT	^/api/follows/[^/]+$	10.567	0.440	2.859	2.859	0
0.5%	35	0	0	GET	^/api/memories/[^/]+	8.492	0.243	0.649	0.649	940
0.5%	25	0	0	POST	/api/auth/register	8.317	0.333	0.980	0.980	228
0.5%	56	0	0	GET	/api/users	8.047	0.144	0.429	0.429	272
0.3%	14	0	0	PUT	/api/me/profile-image	6.090	0.435	0.700	0.700	0
0.3%	32	0	0	GET	^/api/posts/[^/]+/realmojis$	5.265	0.165	0.365	0.365	414
0.2%	19	0	0	GET	^/api/posts/[^/]+/comments$	3.881	0.204	0.404	0.404	522
0.2%	13	0	0	PATCH	/api/me	3.656	0.281	0.738	0.738	302
0.2%	8	0	0	PATCH	^/api/posts/[^/]+$	2.941	0.368	1.051	1.051	657
0.2%	19	2	0	GET	/api/me	2.829	0.149	0.243	0.243	259
0.1%	8	0	0	DELETE	^/api/comments/[^/]+$	1.732	0.216	0.325	0.325	0
0.1%	4	0	0	DELETE	^/api/realmojis/[^/]+$	1.297	0.324	0.796	0.796	0
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	0.907	0.302	0.682	0.682	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.0%	142	0	0	GET	/api/notifications	1093.587	7.701	9.497	9.577	179
15.4%	174	0	0	POST	/api/posts	382.998	2.201	7.833	8.706	638
7.8%	1004	5	0	GET	^/api/images/[^/]+	193.476	0.193	1.261	1.426	170448
5.3%	579	0	0	GET	^/api/feeds/[^/]+	132.918	0.230	1.350	2.172	2787
4.4%	251	0	0	POST	/api/auth/login	109.675	0.437	4.247	4.660	252
4.0%	347	0	0	GET	/api/memories	98.651	0.284	1.324	1.921	11793
3.4%	292	11	0	POST	/api/friend-requests	83.905	0.287	2.278	3.446	112
1.8%	15	0	0	GET	/api/me/export	45.405	3.027	7.675	7.675	3161135
1.8%	161	0	0	GET	^/api/posts/me/current$	45.363	0.282	1.391	2.451	635
1.7%	222	0	0	GET	/api/posting-window/current	41.668	0.188	0.641	0.705	240
1.3%	83	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	32.554	0.392	2.251	2.251	116
1.2%	28	0	0	POST	^/api/posts/[^/]+/realmojis$	29.626	1.058	4.189	4.189	441
1.1%	134	0	0	GET	^/api/relationships/[^/]+	27.309	0.204	1.149	1.721	1133
0.9%	47	0	0	POST	/api/auth/register	23.184	0.493	2.158	2.158	229
0.9%	129	0	0	GET	/api/users	22.069	0.171	0.539	0.945	260
0.7%	42	1	0	POST	^/api/posts/[^/]+/comments$	16.206	0.386	2.981	2.981	275
0.6%	77	9	0	GET	^/api/posts/[^/]+$	16.018	0.208	2.353	2.353	528
0.6%	59	0	0	GET	^/api/users/[^/]+$	14.916	0.253	2.103	2.103	172
0.6%	35	0	0	PUT	^/api/follows/[^/]+$	14.057	0.402	1.859	1.859	0
0.5%	24	0	0	PUT	/api/me/profile-image	12.869	0.536	2.246	2.246	0
0.4%	44	0	0	GET	^/api/memories/[^/]+	9.651	0.219	0.551	0.551	1403
0.4%	36	0	0	POST	/api/auth/logout	9.123	0.253	0.878	0.878	0
0.2%	19	0	0	PATCH	/api/me	5.651	0.297	0.817	0.817	289
0.2%	5	0	0	DELETE	^/api/posts/[^/]+$	5.228	1.046	3.957	3.957	0
0.2%	7	0	0	PATCH	^/api/posts/[^/]+$	5.201	0.743	3.543	3.543	606
0.2%	27	0	0	GET	/api/me	4.423	0.164	0.377	0.377	295
0.2%	26	0	0	GET	^/api/posts/[^/]+/comments$	3.941	0.152	0.403	0.403	884
0.1%	18	0	0	GET	^/api/posts/[^/]+/realmojis$	2.887	0.160	0.288	0.288	459
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	1.380	0.460	0.714	0.714	0
0.1%	8	0	0	DELETE	^/api/comments/[^/]+$	1.321	0.165	0.312	0.312	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 17447 クエリ / 83 種類 / 合計 70.025s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	43.6%	99	30.557	308.66	1301.82	6742.8	133.1	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	15.7%	1458	10.997	7.54	326.84	114.8	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
3	10.7%	680	7.464	10.98	140.65	660.1	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
4	9.9%	18	6.927	384.84	1830.58	32650.7	1651.4	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
5	6.8%	2639	4.790	1.81	7.89	11.6	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
6	3.3%	1	2.344	2344.39	2344.39	29325.0	1199397.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	2.2%	1	1.530	1530.14	1530.14	19623.0	625294.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	1.6%	1	1.086	1085.57	1085.57	11528.0	529643.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	1.2%	475	0.805	1.70	11.56	32.1	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.8%	2	0.532	265.99	371.99	212153.5	2495.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.6%	140	0.403	2.88	13.42	38.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
12	0.5%	2	0.384	191.88	374.99	20236.0	2467.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
13	0.3%	1	0.229	229.41	229.41	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
14	0.3%	1	0.215	215.45	215.45	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
15	0.3%	1	0.211	211.40	211.40	69652.0	29540.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
-- rank 3
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: a8abf32e07c2e96bc84ab1f248ab39e0169daf37
Type: cpu
Time: 2026-10-31 11:19:07 JST
Duration: 60.05s, Total samples = 64.32s (107.11%)
Showing nodes accounting for 53.95s, 83.88% of 64.32s total
Dropped 908 nodes (cum <= 0.32s)
Showing top 40 nodes out of 216
      flat  flat%   sum%        cum   cum%
    12.93s 20.10% 20.10%     12.93s 20.10%  runtime.spanClass.sizeclass (inline)
     7.82s 12.16% 32.26%      7.82s 12.16%  crypto/internal/fips140/sha256.blockAVX2
     4.73s  7.35% 39.61%     18.65s 29.00%  runtime.tryDeferToSpanScan
     3.56s  5.53% 45.15%      3.56s  5.53%  runtime.memclrNoHeapPointers
     2.42s  3.76% 48.91%      3.70s  5.75%  image/jpeg.(*decoder).reconstructBlock
     2.21s  3.44% 52.35%      2.21s  3.44%  runtime.memmove
     2.15s  3.34% 55.69%      8.80s 13.68%  runtime.scanObjectsSmall
     2.08s  3.23% 58.92%     15.14s 23.54%  runtime.scanObject
     2.03s  3.16% 62.08%      2.03s  3.16%  internal/runtime/syscall/linux.Syscall6
     1.10s  1.71% 63.79%      5.35s  8.32%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.02s  1.59% 65.38%      1.02s  1.59%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.99s  1.54% 66.92%      8.21s 12.76%  image/jpeg.(*decoder).processSOS
     0.92s  1.43% 68.35%      1.55s  2.41%  image/jpeg.(*decoder).decodeHuffman
     0.87s  1.35% 69.70%      0.87s  1.35%  aeshashbody
     0.84s  1.31% 71.00%      0.84s  1.31%  memeqbody
     0.82s  1.27% 72.28%     10.38s 16.14%  runtime.scanSpan
     0.61s  0.95% 73.23%      0.62s  0.96%  image/jpeg.(*decoder).receiveExtend
     0.58s   0.9% 74.13%      0.63s  0.98%  runtime.(*spanInlineMarkBits).tryAcquire
     0.52s  0.81% 74.94%      0.68s  1.06%  image/jpeg.idctCols
     0.51s  0.79% 75.73%      0.51s  0.79%  runtime.extractHeapBitsSmall
     0.47s  0.73% 76.46%      0.64s     1%  image/jpeg.(*decoder).ensureNBits
     0.43s  0.67% 77.13%      8.55s 13.29%  crypto/internal/fips140/sha256.(*Digest).Write
     0.43s  0.67% 77.80%      8.75s 13.60%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.41s  0.64% 78.44%      0.57s  0.89%  example.com/isureal-company/isureal.postLess
     0.39s  0.61% 79.04%      0.56s  0.87%  image/jpeg.idctRows
     0.35s  0.54% 79.59%     10.02s 15.58%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.35s  0.54% 80.13%      2.18s  3.39%  runtime.mapaccess1_faststr
     0.34s  0.53% 80.66%      9.27s 14.41%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.33s  0.51% 81.17%      0.33s  0.51%  image/jpeg.dctBox (inline)
     0.33s  0.51% 81.69%      0.33s  0.51%  runtime.nextFreeFast (inline)
     0.20s  0.31% 82.00%      0.65s  1.01%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.20s  0.31% 82.31%      0.59s  0.92%  runtime.scanObjectSmall
     0.17s  0.26% 82.57%      1.08s  1.68%  runtime.mallocgcSmallScanNoHeader
     0.15s  0.23% 82.80%      0.56s  0.87%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.15s  0.23% 83.04%     10.82s 16.82%  example.com/isureal-company/isureal.computeImageSignature
     0.13s   0.2% 83.24%     14.60s 22.70%  runtime.mallocgc
     0.11s  0.17% 83.41%      0.55s  0.86%  database/sql.convertAssignRows
     0.10s  0.16% 83.57%      7.92s 12.31%  crypto/internal/fips140/sha256.block
     0.10s  0.16% 83.72%      1.22s  1.90%  runtime.mapassign_faststr
     0.10s  0.16% 83.88%      0.37s  0.58%  runtime.wbBufFlush1
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 64.9% / 100% | 49% | 0.4% | 7 | 156 MB | CPU が飽和 |
| isu2 (2 cpu) | 7.9% / 59% | 0% | 0.4% | 3 | 206 MB |  |
| isu3 (2 cpu) | 7.9% / 65% | 0% | 0.6% | 3 | 185 MB |  |
| isu4 (2 cpu) | 32.9% / 84% | 0% | 20.2% | 3 | 105 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 111% / nginx 12% |
| isu2 | isureal 14% |
| isu3 | isureal 13% / nginx 1% |
| isu4 | postgres 22% |

## Diff from previous run

スコア: 15668 → 24028 (+8360)

`8f720c2..93a8224`

```
 conf/isu1/home/isucon/env.sh                    |   1 +
 conf/isu4/etc/systemd/system/isureal-go.service |  17 ++
 webapp/go/cmd/image-export/main.go              |  11 +-
 webapp/go/cmd/isureal/main.go                   |   6 +
 webapp/go/export_api.go                         |   9 +
 webapp/go/image_local.go                        |  37 +++-
 webapp/go/image_signature.go                    |  33 +++-
 webapp/go/memory_api.go                         |  18 ++
 webapp/go/offload.go                            | 250 ++++++++++++++++++++++++
 webapp/go/post_api.go                           |   8 +-
 webapp/go/profile_image_api.go                  |  11 +-
 webapp/go/realmoji_api.go                       |   8 +-
 12 files changed, 363 insertions(+), 46 deletions(-)
```
