# run 20261031-171626-1662a5d

## Summary

| | |
|---|---|
| score | 191008 |
| status | pass |
| ref | sse-read-timeout |
| sha | 1662a5d5b0f1 |
| 所要時間 | 118s |
| 前回比 | -28721 (-13.1%) vs 20261031-171010-53bad13 |
| ベスト比 | -41102 (-17.7%) vs 20261031-164055-bdf05cd = 232110 |
| note | SSE proxy_read_timeout 20s on main 885b12c |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	144	385	730	730
europe	7	21	66	197	568	568
america	6	14	38	114	318	316
```

```
{"id":"33249","status":"done","portalStatus":"finished","passed":true,"score":191008,"raw":191008,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:16:53.121Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:16:56.133Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:16:59.102Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:17:02.117Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T08:17:05.103Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T08:17:06.401Z level=INFO msg="user moved" user=leoshi_5239 from=asia to=europe
time=2026-10-31T08:17:08.104Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T08:17:11.118Z level=INFO msg="asia 3 日目: 144 人"
time=2026-10-31T08:17:12.758Z level=INFO msg="user moved" user=tachiku.5605 from=asia to=europe
time=2026-10-31T08:17:13.412Z level=INFO msg="user moved" user=tachiku.5605 from=europe to=asia
time=2026-10-31T08:17:14.105Z level=INFO msg="europe 3 日目: 66 人"
time=2026-10-31T08:17:17.103Z level=INFO msg="america 3 日目: 38 人"
time=2026-10-31T08:17:17.607Z level=INFO msg="user moved" user=ulleo.5832 from=asia to=europe
time=2026-10-31T08:17:20.115Z level=INFO msg="asia 4 日目: 385 人"
time=2026-10-31T08:17:21.993Z level=INFO msg="user moved" user=leoshi_5239 from=europe to=america
time=2026-10-31T08:17:23.103Z level=INFO msg="europe 4 日目: 197 人"
time=2026-10-31T08:17:24.398Z level=INFO msg="user moved" user=mitsu_6813 from=asia to=america
time=2026-10-31T08:17:25.335Z level=INFO msg="user moved" user=maxeva.7350 from=asia to=europe
time=2026-10-31T08:17:25.709Z level=INFO msg="user moved" user=tachiku.5605 from=asia to=europe
time=2026-10-31T08:17:26.118Z level=INFO msg="america 4 日目: 114 人"
time=2026-10-31T08:17:27.026Z level=INFO msg="user moved" user=tashino_7391 from=asia to=america
time=2026-10-31T08:17:29.180Z level=INFO msg="asia 5 日目: 730 人"
time=2026-10-31T08:17:29.298Z level=INFO msg="user moved" user=ulsa-9300 from=europe to=asia
time=2026-10-31T08:17:29.904Z level=INFO msg="user moved" user=ulleo.5832 from=europe to=america
time=2026-10-31T08:17:32.129Z level=INFO msg="europe 5 日目: 568 人"
time=2026-10-31T08:17:35.532Z level=INFO msg="america 5 日目: 318 人"
time=2026-10-31T08:17:35.984Z level=INFO msg="user moved" user=naelor-9608 from=america to=europe
time=2026-10-31T08:17:47.103Z level=INFO msg="asia 最終日（グランプリ）: 730 人"
time=2026-10-31T08:17:47.872Z level=INFO msg="america 最終日（グランプリ）: 316 人"
time=2026-10-31T08:17:50.853Z level=INFO msg="europe 最終日（グランプリ）: 568 人"
time=2026-10-31T08:17:53.187Z level=INFO msg="グランプリ: アジア 297 人, ヨーロッパ 0 人, アメリカ 243 人"
{
  "pass": true,
  "score": 191008,
  "messages": [
    "solmuhe5134: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
68.1%	2077	1	0	GET	/api/notifications	30022.098	14.455	21.121	21.121	114
6.2%	6351	0	0	GET	^/api/feeds/[^/]+	2738.778	0.431	1.691	3.345	3184
3.5%	2716	0	0	GET	/api/memories	1546.960	0.570	2.502	7.719	11099
3.4%	10058	347	0	GET	^/api/images/[^/]+	1482.661	0.147	1.467	6.807	136596
2.9%	1978	25	0	POST	/api/posts	1275.671	0.645	2.025	2.906	640
2.9%	2771	0	0	GET	/api/posting-window/current	1262.239	0.456	1.552	3.345	237
2.2%	1973	0	0	GET	/api/users	970.901	0.492	1.589	3.345	231
1.7%	1479	22	0	GET	^/api/posts/me/current$	748.918	0.506	1.467	1.922	637
1.4%	4003	45	0	POST	/api/friend-requests	623.583	0.156	1.328	3.345	114
1.2%	199	0	0	GET	/api/me/export	514.655	2.586	14.726	18.745	2055667
1.0%	1997	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	423.909	0.212	1.922	3.345	116
0.9%	502	120	1	PUT	/api/me/profile-image	408.665	0.814	10.849	14.266	5
0.7%	3897	637	0	POST	/api/auth/login	327.188	0.084	1.212	3.250	206
0.6%	878	0	0	GET	^/api/relationships/[^/]+	286.183	0.326	1.363	2.066	1390
0.6%	883	54	0	GET	^/api/posts/[^/]+$	264.032	0.299	1.922	3.250	577
0.5%	391	3	0	GET	/api/me	232.942	0.596	1.536	1.715	300
0.4%	495	0	0	GET	^/api/memories/[^/]+	186.627	0.377	1.922	3.250	667
0.3%	452	0	0	GET	^/api/users/[^/]+$	129.598	0.287	1.922	1.922	187
0.3%	340	5	0	POST	^/api/posts/[^/]+/realmojis$	118.573	0.349	1.922	3.250	453
0.3%	387	1	0	GET	^/api/posts/[^/]+/realmojis$	112.951	0.292	1.467	3.250	1263
0.2%	291	3	0	GET	^/api/posts/[^/]+/comments$	75.628	0.260	1.578	1.611	1687
0.1%	392	4	0	POST	^/api/posts/[^/]+/comments$	63.145	0.161	1.328	2.394	330
0.1%	544	0	0	POST	/api/auth/logout	58.328	0.107	1.328	2.394	0
0.1%	486	0	0	PUT	^/api/follows/[^/]+$	42.330	0.087	0.693	2.873	0
0.1%	752	1	0	POST	/api/auth/register	37.407	0.050	0.429	1.018	227
0.1%	335	0	0	PATCH	/api/me	33.960	0.101	0.702	1.212	273
0.1%	63	0	0	DELETE	^/api/posts/[^/]+$	31.375	0.498	1.400	1.400	0
0.1%	126	2	0	PATCH	^/api/posts/[^/]+$	22.574	0.179	0.990	1.018	630
0.0%	1	0	0	POST	/api/initialize	15.450	15.450	15.450	15.450	90
0.0%	80	0	0	DELETE	^/api/comments/[^/]+$	14.356	0.179	2.394	2.394	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
78.3%	1484	0	0	GET	/api/notifications	22244.907	14.990	21.047	21.151	106
3.5%	10033	0	0	GET	^/api/feeds/[^/]+	981.080	0.098	0.591	1.066	3342
2.6%	3356	48	0	POST	/api/friend-requests	729.826	0.217	0.786	1.169	114
2.5%	7720	46	0	GET	^/api/images/[^/]+	698.522	0.090	0.800	7.134	166054
2.0%	1996	0	0	GET	/api/posting-window/current	566.028	0.284	1.084	1.351	239
2.0%	1179	55	0	POST	/api/posts	564.191	0.479	1.286	1.551	621
1.7%	2358	0	0	POST	/api/auth/login	473.226	0.201	0.939	1.485	248
0.8%	565	0	0	POST	/api/auth/register	217.000	0.384	0.925	0.968	229
0.7%	1635	0	0	GET	/api/users	203.064	0.124	0.606	0.856	235
0.7%	1613	0	0	GET	/api/memories	202.965	0.126	0.981	1.729	16158
0.7%	328	0	0	PUT	/api/me/profile-image	191.437	0.584	1.525	1.618	0
0.5%	463	4	0	POST	^/api/posts/[^/]+/realmojis$	142.995	0.309	1.090	1.316	474
0.5%	1631	9	0	GET	^/api/posts/[^/]+$	138.469	0.085	0.543	0.690	620
0.4%	873	0	0	GET	^/api/posts/me/current$	122.208	0.140	0.576	0.812	651
0.4%	748	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	120.720	0.161	0.657	0.885	116
0.4%	1150	0	0	GET	^/api/relationships/[^/]+	104.265	0.091	0.564	0.853	460
0.3%	366	0	0	PUT	^/api/follows/[^/]+$	90.359	0.247	0.899	0.936	0
0.3%	303	0	0	PATCH	/api/me	87.232	0.288	0.791	0.935	278
0.3%	472	1	0	POST	^/api/posts/[^/]+/comments$	75.967	0.161	0.697	0.847	347
0.2%	811	0	0	GET	^/api/users/[^/]+$	69.244	0.085	0.550	0.629	201
0.2%	425	0	0	POST	/api/auth/logout	67.463	0.159	0.827	1.029	0
0.2%	671	0	0	GET	^/api/posts/[^/]+/realmojis$	57.922	0.086	0.481	0.565	2646
0.2%	628	0	0	GET	^/api/posts/[^/]+/comments$	54.512	0.087	0.479	0.666	3466
0.1%	335	2	0	GET	/api/me	37.279	0.111	0.341	0.475	304
0.1%	361	0	0	GET	^/api/memories/[^/]+	32.100	0.089	0.711	1.173	833
0.1%	123	0	0	PATCH	^/api/posts/[^/]+$	30.151	0.245	0.751	0.795	662
0.1%	91	0	0	DELETE	^/api/realmojis/[^/]+$	26.848	0.295	0.897	0.897	0
0.1%	141	0	0	GET	/api/me/export	22.481	0.159	2.155	3.158	2372791
0.1%	46	0	0	DELETE	^/api/posts/[^/]+$	20.075	0.436	1.118	1.118	0
0.0%	95	0	0	DELETE	^/api/comments/[^/]+$	13.115	0.138	0.608	0.608	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
74.2%	838	1	0	GET	/api/notifications	10388.687	12.397	20.001	20.002	108
4.3%	5994	21	0	GET	^/api/images/[^/]+	608.446	0.102	1.272	1.931	120986
4.3%	4963	0	0	GET	^/api/feeds/[^/]+	608.363	0.123	1.061	1.889	3011
2.6%	801	1	0	POST	/api/posts	361.672	0.452	1.826	2.172	656
2.3%	1833	15	0	POST	/api/friend-requests	325.551	0.178	0.714	1.311	114
2.1%	1840	0	0	POST	/api/auth/login	297.545	0.162	0.816	1.599	253
1.8%	1426	0	0	GET	/api/memories	258.546	0.181	1.220	2.225	17242
1.2%	675	0	0	GET	^/api/posts/me/current$	163.954	0.243	1.088	1.860	656
1.1%	1153	0	0	GET	/api/posting-window/current	147.926	0.128	1.339	1.889	240
0.7%	923	0	0	GET	/api/users	96.635	0.105	0.659	1.052	233
0.6%	264	4	0	POST	^/api/posts/[^/]+/realmojis$	91.014	0.345	1.311	1.704	468
0.5%	749	4	0	GET	^/api/posts/[^/]+$	65.316	0.087	0.697	0.907	620
0.5%	379	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	63.180	0.167	0.774	1.718	116
0.4%	583	0	0	GET	^/api/relationships/[^/]+	62.292	0.107	0.903	1.577	563
0.4%	313	0	0	POST	/api/auth/register	57.292	0.183	0.496	0.585	230
0.4%	337	0	0	POST	/api/auth/logout	55.313	0.164	1.231	1.705	0
0.4%	174	0	0	PUT	/api/me/profile-image	51.543	0.296	0.792	1.789	0
0.3%	275	0	0	POST	^/api/posts/[^/]+/comments$	45.628	0.166	0.773	0.868	328
0.2%	195	0	0	PUT	^/api/follows/[^/]+$	32.727	0.168	0.745	0.784	0
0.2%	328	0	0	GET	^/api/memories/[^/]+	30.590	0.093	0.736	2.105	1061
0.2%	353	0	0	GET	^/api/users/[^/]+$	28.353	0.080	0.323	0.748	202
0.2%	150	0	0	PATCH	/api/me	23.929	0.160	0.657	0.658	273
0.2%	287	0	0	GET	^/api/posts/[^/]+/realmojis$	22.871	0.080	0.305	0.343	3668
0.2%	122	0	0	GET	/api/me/export	22.619	0.185	1.375	2.070	2890702
0.2%	246	1	0	GET	^/api/posts/[^/]+/comments$	21.426	0.087	0.493	0.756	4201
0.1%	180	1	0	GET	/api/me	16.884	0.094	0.397	0.824	303
0.1%	46	0	0	DELETE	^/api/realmojis/[^/]+$	16.204	0.352	1.103	1.103	0
0.1%	55	0	0	PATCH	^/api/posts/[^/]+$	15.483	0.282	1.401	1.401	663
0.1%	25	0	0	DELETE	^/api/posts/[^/]+$	12.903	0.516	2.299	2.299	0
0.1%	50	0	0	DELETE	^/api/comments/[^/]+$	7.275	0.146	0.799	0.799	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 10014 クエリ / 48 種類 / 合計 127.393s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	43.7%	209	55.712	266.56	808.50	1314.2	133.6	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	20.5%	4	26.098	6524.51	10086.76	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.9%	4	18.925	4731.29	7211.03	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.5%	6643	10.868	1.64	10.38	10.2	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	7.2%	4	9.233	2308.27	3661.22	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.1%	7	2.624	374.82	535.74	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.9%	7	1.157	165.28	719.42	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.5%	7	0.599	85.51	162.78	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.4%	1	0.458	457.61	457.61	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.3%	1	0.353	352.64	352.64	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.3%	1	0.329	328.79	328.79	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.164	164.12	164.12	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	7	0.164	23.40	34.21	166.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
14	0.1%	1	0.110	110.42	110.42	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	1	0.103	103.39	103.39	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows

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
Build ID: 0eb83cf1921478aee6a7b30e210bbe2042744bfc
Type: cpu
Time: 2026-10-31 17:16:31 JST
Duration: 60.13s, Total samples = 33.52s (55.74%)
Showing nodes accounting for 25.22s, 75.24% of 33.52s total
Dropped 972 nodes (cum <= 0.17s)
Showing top 40 nodes out of 258
      flat  flat%   sum%        cum   cum%
     4.62s 13.78% 13.78%      4.62s 13.78%  crypto/internal/fips140/sha256.blockAVX2
     3.59s 10.71% 24.49%      3.59s 10.71%  internal/runtime/syscall/linux.Syscall6
     1.53s  4.56% 29.06%      6.21s 18.53%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.51s  4.50% 33.56%      1.51s  4.50%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.43s  4.27% 37.83%      1.43s  4.27%  runtime.memmove
     1.11s  3.31% 41.14%      1.11s  3.31%  aeshashbody
     1.02s  3.04% 44.18%      3.80s 11.34%  runtime.scanObject
     0.98s  2.92% 47.11%      1.24s  3.70%  runtime.findObject
     0.92s  2.74% 49.85%      0.92s  2.74%  runtime.memclrNoHeapPointers
     0.79s  2.36% 52.21%      1.46s  4.36%  runtime.tryDeferToSpanScan
     0.70s  2.09% 54.30%      2.29s  6.83%  runtime.mapaccess1_faststr
     0.52s  1.55% 55.85%      0.77s  2.30%  example.com/isureal-company/isureal/jpegcheck.(*decoder).skipAC
     0.47s  1.40% 57.25%      0.47s  1.40%  runtime.spanClass.sizeclass (inline)
     0.46s  1.37% 58.62%      0.65s  1.94%  example.com/isureal-company/isureal.postLess
     0.38s  1.13% 59.76%      0.60s  1.79%  database/sql.convertAssignRows
     0.36s  1.07% 60.83%      5.18s 15.45%  crypto/internal/fips140/sha256.(*Digest).Write
     0.32s  0.95% 61.78%      0.32s  0.95%  memeqbody
     0.31s  0.92% 62.71%      1.19s  3.55%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.31s  0.92% 63.63%      0.31s  0.92%  runtime.futex
     0.30s  0.89% 64.53%      5.25s 15.66%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.30s  0.89% 65.42%      1.07s  3.19%  runtime.mallocgcSmallScanNoHeader
     0.29s  0.87% 66.29%      0.29s  0.87%  sync/atomic.(*Int32).Add
     0.27s  0.81% 67.09%      1.49s  4.45%  runtime.mapassign_faststr
     0.25s  0.75% 67.84%      0.25s  0.75%  runtime.nextFreeFast (inline)
     0.24s  0.72% 68.56%      0.24s  0.72%  example.com/isureal-company/isureal/jpegcheck.(*decoder).prefetch
     0.23s  0.69% 69.24%      0.26s  0.78%  runtime.(*mspan).writeHeapBitsSmall
     0.21s  0.63% 69.87%      5.54s 16.53%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.21s  0.63% 70.50%      0.22s  0.66%  example.com/isureal-company/isureal.rmLess
     0.20s   0.6% 71.09%      2.31s  6.89%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.20s   0.6% 71.69%      2.12s  6.32%  runtime.mallocgc
     0.18s  0.54% 72.23%      0.24s  0.72%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.15s  0.45% 72.67%      0.47s  1.40%  runtime.mapaccess2_faststr
     0.14s  0.42% 73.09%      0.24s  0.72%  internal/runtime/maps.(*Iter).Next
     0.13s  0.39% 73.48%      0.17s  0.51%  runtime.typePointers.next
     0.12s  0.36% 73.84%      5.87s 17.51%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.11s  0.33% 74.16%      0.33s  0.98%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.11s  0.33% 74.49%      0.17s  0.51%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
     0.09s  0.27% 74.76%      0.23s  0.69%  runtime.greyobject
     0.08s  0.24% 75.00%      0.24s  0.72%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.08s  0.24% 75.24%      1.23s  3.67%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 74.2% / 100% | 36% | 30.1% | 6 | 105 MB | CPU が飽和 |
| isu2 (2 cpu) | 26.0% / 100% | 5% | 0.8% | 5 | 166 MB | run queue が積もっている |
| isu3 (2 cpu) | 25.0% / 100% | 6% | 2.2% | 4 | 177 MB |  |
| isu4 (2 cpu) | 64.5% / 100% | 55% | 3.0% | 5 | 103 MB | CPU が飽和 |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 52% / find 15% / nginx 10% / kswapd0 5% / kcompactd0 2% |
| isu2 | isureal 47% / nginx 3% |
| isu3 | isureal 44% / nginx 2% |
| isu4 | isureal 104% / postgres 16% |

## Diff from previous run

スコア: 219729 → 191008 (-28721)

`53bad13..1662a5d`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf  | 1 +
 conf/isu1/home/isucon/env.sh                    | 2 +-
 conf/isu2/etc/nginx/sites-enabled/isureal.conf  | 1 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf  | 1 +
 conf/isu4/etc/systemd/system/isureal-go.service | 5 ++---
 5 files changed, 6 insertions(+), 4 deletions(-)
```
