# run 20261031-162248-73b6aff

## Summary

| | |
|---|---|
| score | 205275 |
| status | pass |
| ref | origin/main |
| sha | 73b6aff45087 |
| 所要時間 | 244s |
| 前回比 | +2700 (1.3%) vs 20261031-161926-3f85847 |
| ベスト比 | +2700 (1.3%) vs 20261031-161926-3f85847 = 202575 **ベスト更新** |
| note | main 73b6aff baseline after #93 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	144	387	757	759
europe	7	19	55	160	452	451
america	6	20	61	180	494	493
```

```
{"id":"28079","status":"done","portalStatus":"finished","passed":true,"score":205275,"raw":205275,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:25:21.151Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:25:24.141Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:25:27.139Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:25:30.154Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T07:25:30.582Z level=INFO msg="user moved" user=leomi5162 from=asia to=america
time=2026-10-31T07:25:31.213Z level=INFO msg="user moved" user=leomi5162 from=america to=asia
time=2026-10-31T07:25:33.139Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T07:25:36.139Z level=INFO msg="america 2 日目: 20 人"
time=2026-10-31T07:25:39.157Z level=INFO msg="asia 3 日目: 144 人"
time=2026-10-31T07:25:41.869Z level=INFO msg="user moved" user=leomi5162 from=asia to=america
time=2026-10-31T07:25:42.141Z level=INFO msg="europe 3 日目: 55 人"
time=2026-10-31T07:25:45.140Z level=INFO msg="america 3 日目: 61 人"
time=2026-10-31T07:25:45.631Z level=INFO msg="user moved" user=suma.6332 from=europe to=america
time=2026-10-31T07:25:48.155Z level=INFO msg="asia 4 日目: 387 人"
time=2026-10-31T07:25:51.140Z level=INFO msg="europe 4 日目: 160 人"
time=2026-10-31T07:25:53.742Z level=INFO msg="user moved" user=yukose7241 from=asia to=europe
time=2026-10-31T07:25:54.154Z level=INFO msg="america 4 日目: 180 人"
time=2026-10-31T07:25:55.359Z level=INFO msg="user moved" user=muka-8432 from=asia to=europe
time=2026-10-31T07:25:55.667Z level=INFO msg="user moved" user=suma.6332 from=america to=asia
time=2026-10-31T07:25:55.958Z level=INFO msg="user moved" user=nime8014 from=asia to=europe
time=2026-10-31T07:25:57.044Z level=INFO msg="user moved" user=leomi5162 from=america to=asia
time=2026-10-31T07:25:57.161Z level=INFO msg="asia 5 日目: 757 人"
time=2026-10-31T07:25:59.012Z level=INFO msg="user moved" user=motorma9123 from=america to=europe
time=2026-10-31T07:26:00.174Z level=INFO msg="europe 5 日目: 452 人"
time=2026-10-31T07:26:00.657Z level=INFO msg="user moved" user=sashi.9336 from=america to=asia
time=2026-10-31T07:26:02.285Z level=INFO msg="user moved" user=muhiku9443 from=america to=asia
time=2026-10-31T07:26:02.803Z level=INFO msg="user moved" user=ulnoaleo-10059 from=america to=asia
time=2026-10-31T07:26:03.465Z level=INFO msg="america 5 日目: 494 人"
time=2026-10-31T07:26:12.174Z level=INFO msg="america 最終日（グランプリ）: 493 人"
time=2026-10-31T07:26:12.185Z level=INFO msg="europe 最終日（グランプリ）: 451 人"
time=2026-10-31T07:26:12.209Z level=INFO msg="asia 最終日（グランプリ）: 759 人"
time=2026-10-31T07:26:21.231Z level=INFO msg="グランプリ: アジア 451 人, ヨーロッパ 260 人, アメリカ 277 人"
{
  "pass": true,
  "score": 205275,
  "messages": [
    "hoorme_5051: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "yuivy.13220: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.42.38:33282-\u003e52.69.202.109:8080: write tcp 10.27.42.38:33282-\u003e52.69.202.109:8080: use of closed network connection"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
46.5%	1408	1	0	GET	/api/notifications	13790.262	9.794	15.063	15.082	170
9.5%	11245	0	0	GET	^/api/feeds/[^/]+	2805.773	0.250	1.651	2.950	4570
8.2%	2120	72	0	POST	/api/posts	2445.461	1.154	6.996	8.101	627
7.9%	3045	0	0	GET	/api/memories	2353.472	0.773	5.886	10.290	10850
5.6%	2957	0	0	GET	/api/posting-window/current	1655.557	0.560	1.983	2.466	237
3.6%	2100	0	0	GET	/api/users	1064.498	0.507	1.885	2.170	238
2.6%	1775	1	0	GET	^/api/posts/me/current$	770.766	0.434	1.536	2.217	650
2.5%	222	0	0	GET	/api/me/export	741.468	3.340	18.161	19.725	1849814
2.1%	12410	237	0	GET	^/api/images/[^/]+	619.691	0.050	0.911	2.400	168248
1.9%	551	132	0	PUT	/api/me/profile-image	554.000	1.005	13.334	16.178	4
1.2%	430	3	0	GET	/api/me	362.359	0.843	2.077	2.108	300
1.2%	1413	0	0	GET	^/api/relationships/[^/]+	350.445	0.248	1.670	2.936	1578
1.1%	862	21	0	GET	^/api/posts/[^/]+/realmojis$	339.302	0.394	3.131	4.035	1540
1.1%	4522	54	0	POST	/api/friend-requests	334.797	0.074	1.007	1.281	114
1.0%	655	6	0	POST	^/api/posts/[^/]+/realmojis$	305.824	0.467	4.066	4.118	469
0.8%	1984	66	0	GET	^/api/posts/[^/]+$	237.158	0.120	0.912	1.316	611
0.7%	584	0	0	GET	^/api/memories/[^/]+	207.688	0.356	1.977	3.407	612
0.7%	3852	9	0	POST	^/api/friend-requests/[^/]+/[^/]+$	198.828	0.052	1.003	1.256	116
0.4%	4239	624	0	POST	/api/auth/login	118.798	0.028	0.465	1.070	211
0.4%	765	16	0	GET	^/api/posts/[^/]+/comments$	106.149	0.139	0.999	1.476	2248
0.3%	958	0	0	GET	^/api/users/[^/]+$	98.909	0.103	0.722	1.330	200
0.1%	88	0	0	DELETE	^/api/posts/[^/]+$	44.325	0.504	1.360	1.360	0
0.1%	652	0	0	POST	^/api/posts/[^/]+/comments$	37.980	0.058	0.858	1.278	349
0.1%	198	0	0	PATCH	^/api/posts/[^/]+$	25.365	0.128	1.111	1.224	651
0.1%	601	0	0	POST	/api/auth/logout	22.858	0.038	0.561	1.077	0
0.1%	515	0	0	PUT	^/api/follows/[^/]+$	22.246	0.043	0.460	0.658	0
0.1%	1	0	0	POST	/api/initialize	21.375	21.375	21.375	21.375	90
0.1%	85	0	0	DELETE	^/api/realmojis/[^/]+$	17.188	0.202	0.501	0.501	0
0.0%	374	0	0	PATCH	/api/me	13.681	0.037	0.464	0.505	274
0.0%	775	1	0	POST	/api/auth/register	7.241	0.009	0.170	0.216	227
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
48.8%	714	1	0	GET	/api/notifications	5450.482	7.634	11.894	11.901	175
8.2%	1142	24	0	POST	/api/posts	919.847	0.805	7.346	8.120	642
7.7%	6580	0	0	GET	^/api/feeds/[^/]+	862.679	0.131	1.318	2.067	3844
6.5%	6230	35	0	GET	^/api/images/[^/]+	731.655	0.117	1.114	4.010	166897
5.1%	2483	41	0	POST	/api/friend-requests	565.346	0.228	1.309	1.530	114
3.4%	1808	0	0	POST	/api/auth/login	378.949	0.210	0.959	1.603	248
2.5%	1504	0	0	GET	/api/memories	282.833	0.188	1.737	2.610	13902
2.4%	270	0	0	PUT	/api/me/profile-image	267.449	0.991	3.703	3.705	0
1.9%	1619	0	0	GET	/api/posting-window/current	208.540	0.129	1.108	1.781	239
1.6%	867	0	0	GET	^/api/posts/me/current$	174.525	0.201	1.729	1.883	655
1.5%	374	10	0	POST	^/api/posts/[^/]+/realmojis$	170.652	0.456	3.575	4.979	460
1.5%	447	0	0	POST	/api/auth/register	168.051	0.376	0.999	1.339	228
1.3%	1213	0	0	GET	/api/users	149.022	0.123	0.845	1.879	237
0.9%	533	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	99.636	0.187	0.907	1.470	116
0.8%	1064	27	0	GET	^/api/posts/[^/]+$	94.451	0.089	0.367	1.108	615
0.8%	743	0	0	GET	^/api/relationships/[^/]+	88.010	0.118	1.144	1.893	531
0.8%	356	0	0	PUT	^/api/follows/[^/]+$	86.312	0.242	1.173	1.370	0
0.7%	412	0	0	POST	^/api/posts/[^/]+/comments$	73.144	0.178	1.185	1.342	344
0.6%	233	0	0	PATCH	/api/me	62.918	0.270	1.243	1.327	280
0.4%	281	2	0	GET	/api/me	47.107	0.168	1.322	1.323	301
0.4%	289	0	0	POST	/api/auth/logout	46.780	0.162	0.837	1.124	0
0.4%	466	0	0	GET	^/api/users/[^/]+$	45.173	0.097	0.748	1.141	205
0.3%	440	0	0	GET	^/api/posts/[^/]+/realmojis$	36.444	0.083	0.225	0.962	2583
0.3%	432	1	0	GET	^/api/posts/[^/]+/comments$	35.122	0.081	0.252	0.796	2858
0.3%	271	0	0	GET	^/api/memories/[^/]+	32.460	0.120	0.898	1.781	818
0.2%	93	0	0	PATCH	^/api/posts/[^/]+$	26.719	0.287	1.725	1.725	654
0.2%	38	0	0	DELETE	^/api/posts/[^/]+$	21.119	0.556	1.754	1.754	0
0.2%	45	0	0	DELETE	^/api/realmojis/[^/]+$	17.387	0.386	0.864	0.864	0
0.1%	100	0	0	GET	/api/me/export	13.633	0.136	1.091	1.271	2165986
0.1%	83	0	0	DELETE	^/api/comments/[^/]+$	11.601	0.140	0.312	0.312	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.9%	772	0	0	GET	/api/notifications	4405.797	5.707	9.312	9.313	172
8.9%	1163	32	0	POST	/api/posts	962.802	0.828	7.235	8.169	637
8.2%	5229	28	0	GET	^/api/images/[^/]+	885.326	0.169	1.579	2.748	154899
8.0%	3707	0	0	GET	^/api/feeds/[^/]+	862.802	0.233	1.747	2.785	2659
7.1%	2632	17	0	POST	/api/friend-requests	766.702	0.291	1.643	2.200	114
3.6%	1665	0	0	POST	/api/auth/login	390.223	0.234	1.430	2.021	249
3.5%	1384	0	0	GET	/api/memories	378.025	0.273	2.467	3.274	15866
2.8%	1730	0	0	GET	/api/posting-window/current	300.213	0.174	1.189	2.209	240
2.5%	1338	0	0	GET	/api/users	271.814	0.203	1.437	2.451	241
2.2%	874	20	0	GET	^/api/posts/me/current$	235.352	0.269	1.502	2.467	642
2.2%	497	0	0	POST	/api/auth/register	233.972	0.471	1.663	1.700	230
1.6%	280	0	0	PUT	/api/me/profile-image	173.562	0.620	3.459	3.490	0
1.3%	247	22	0	POST	^/api/posts/[^/]+/realmojis$	136.789	0.554	2.915	4.296	436
0.8%	408	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	89.015	0.218	0.737	1.477	116
0.8%	491	0	0	GET	^/api/relationships/[^/]+	85.708	0.175	1.479	2.527	544
0.7%	272	0	0	PUT	^/api/follows/[^/]+$	74.309	0.273	1.324	1.624	0
0.6%	471	28	0	GET	^/api/posts/[^/]+$	68.449	0.145	0.586	1.064	592
0.5%	258	0	0	PATCH	/api/me	58.928	0.228	1.324	1.624	279
0.5%	242	0	0	POST	/api/auth/logout	52.790	0.218	1.362	2.065	0
0.5%	207	2	0	POST	^/api/posts/[^/]+/comments$	49.166	0.238	0.747	1.660	355
0.4%	87	0	0	GET	/api/me/export	43.914	0.505	4.921	4.921	2637051
0.3%	284	1	0	GET	/api/me	36.394	0.128	0.599	0.719	305
0.3%	212	0	0	GET	^/api/users/[^/]+$	34.108	0.161	0.742	1.093	204
0.3%	212	0	0	GET	^/api/memories/[^/]+	32.414	0.153	1.529	3.001	1026
0.3%	92	4	0	PATCH	^/api/posts/[^/]+$	31.664	0.344	2.045	2.045	623
0.3%	188	2	0	GET	^/api/posts/[^/]+/realmojis$	28.118	0.150	0.636	0.657	2207
0.2%	184	3	0	GET	^/api/posts/[^/]+/comments$	26.180	0.142	0.619	0.686	2442
0.2%	32	0	0	DELETE	^/api/posts/[^/]+$	18.872	0.590	2.619	2.619	0
0.1%	37	0	0	DELETE	^/api/realmojis/[^/]+$	15.069	0.407	0.857	0.857	0
0.1%	36	0	0	DELETE	^/api/comments/[^/]+$	8.233	0.229	0.610	0.610	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 10973 クエリ / 50 種類 / 合計 134.108s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	44.1%	231	59.095	255.82	1577.91	1216.4	124.0	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	19.6%	4	26.336	6583.90	10070.92	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.1%	4	18.956	4739.00	7217.65	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.7%	7464	11.623	1.56	9.87	10.3	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.8%	4	9.100	2275.05	3450.44	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	4.1%	7	5.474	782.03	3230.02	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.5%	1	0.648	647.96	647.96	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
8	0.5%	7	0.616	88.00	163.08	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.5%	7	0.605	86.44	161.59	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
10	0.3%	1	0.420	420.34	420.34	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.233	232.76	232.76	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.171	171.09	171.09	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.128	128.37	128.37	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
14	0.1%	1	0.106	105.52	105.52	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	7	0.089	12.71	22.28	166.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: ea06a9096ca09da60bd90c8f50f07d5a5f26dca9
Type: cpu
Time: 2026-10-31 16:24:53 JST
Duration: 60.04s, Total samples = 25.57s (42.59%)
Showing nodes accounting for 19.91s, 77.86% of 25.57s total
Dropped 732 nodes (cum <= 0.13s)
Showing top 40 nodes out of 238
      flat  flat%   sum%        cum   cum%
     4.37s 17.09% 17.09%      4.37s 17.09%  crypto/internal/fips140/sha256.blockAVX2
     2.43s  9.50% 26.59%      2.43s  9.50%  internal/runtime/syscall/linux.Syscall6
     1.69s  6.61% 33.20%      5.75s 22.49%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.34s  5.24% 38.44%      1.34s  5.24%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.24s  4.85% 43.29%      1.24s  4.85%  runtime.memmove
     0.99s  3.87% 47.16%      0.99s  3.87%  aeshashbody
     0.70s  2.74% 49.90%      0.70s  2.74%  runtime.memclrNoHeapPointers
     0.52s  2.03% 51.94%      0.70s  2.74%  runtime.findObject
     0.48s  1.88% 53.81%      1.97s  7.70%  runtime.scanObject
     0.47s  1.84% 55.65%      0.68s  2.66%  runtime.tryDeferToSpanScan
     0.43s  1.68% 57.33%      1.76s  6.88%  runtime.mapaccess1_faststr
     0.37s  1.45% 58.78%      0.48s  1.88%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.37s  1.45% 60.23%      0.37s  1.45%  memeqbody
     0.36s  1.41% 61.63%      0.53s  2.07%  database/sql.convertAssignRows
     0.35s  1.37% 63.00%         1s  3.91%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.32s  1.25% 64.25%      0.55s  2.15%  example.com/isureal-company/isureal.postLess
     0.26s  1.02% 65.27%      0.26s  1.02%  runtime.nextFreeFast (inline)
     0.23s   0.9% 66.17%      0.72s  2.82%  runtime.mallocgcSmallScanNoHeader
     0.23s   0.9% 67.07%      0.23s   0.9%  sync/atomic.(*Int32).Add (inline)
     0.22s  0.86% 67.93%      4.77s 18.65%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.22s  0.86% 68.79%      0.22s  0.86%  example.com/isureal-company/isureal.rmLess
     0.20s  0.78% 69.57%      0.20s  0.78%  runtime.futex
     0.19s  0.74% 70.32%      0.19s  0.74%  runtime.spanClass.sizeclass (inline)
     0.18s   0.7% 71.02%      1.68s  6.57%  runtime.mallocgc
     0.17s  0.66% 71.69%      0.17s  0.66%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.15s  0.59% 72.27%      5.07s 19.83%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.15s  0.59% 72.86%      0.20s  0.78%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.13s  0.51% 73.37%      4.71s 18.42%  crypto/internal/fips140/sha256.(*Digest).Write
     0.13s  0.51% 73.88%      1.37s  5.36%  runtime.mapassign_faststr
     0.12s  0.47% 74.34%      1.25s  4.89%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.47% 74.81%      0.20s  0.78%  runtime.greyobject
     0.11s  0.43% 75.24%      0.23s   0.9%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.10s  0.39% 75.64%      4.11s 16.07%  example.com/isureal-company/isureal.loadBig
     0.10s  0.39% 76.03%      0.15s  0.59%  internal/runtime/maps.(*Iter).Next
     0.09s  0.35% 76.38%      5.32s 20.81%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.08s  0.31% 76.69%      0.23s   0.9%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.08s  0.31% 77.00%      0.50s  1.96%  runtime.mallocgcSmallNoscan
     0.08s  0.31% 77.32%      0.47s  1.84%  runtime.mapaccess2_faststr
     0.07s  0.27% 77.59%      0.29s  1.13%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.07s  0.27% 77.86%      0.50s  1.96%  runtime.convTstring
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 37.7% / 100% | 19% | 0.7% | 5 | 161 MB | run queue が積もっている |
| isu2 (2 cpu) | 26.7% / 100% | 9% | 0.8% | 5 | 114 MB | run queue が積もっている |
| isu3 (2 cpu) | 29.5% / 100% | 12% | 1.4% | 5 | 156 MB | run queue が積もっている |
| isu4 (2 cpu) | 67.1% / 100% | 58% | 2.7% | 6 | 106 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 61% / nginx 10% |
| isu2 | isureal 48% / nginx 3% |
| isu3 | isureal 54% / nginx 2% |
| isu4 | isureal 110% / postgres 16% |

## Diff from previous run

スコア: 202575 → 205275 (+2700)

`3f85847..73b6aff`

```
 conf/isu1/home/isucon/env.sh                    | 2 +-
 conf/isu4/etc/systemd/system/isureal-go.service | 5 ++---
 conf/isu4/home/isucon/env.sh                    | 1 -
 3 files changed, 3 insertions(+), 5 deletions(-)
```
