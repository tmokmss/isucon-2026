# run 20261031-163606-15a1e6c

## Summary

| | |
|---|---|
| score | 220540 |
| status | pass |
| ref | jpegcheck-fast-ac |
| sha | 15a1e6cee1bd |
| 所要時間 | 140s |
| 前回比 | +13761 (6.7%) vs 20261031-163017-8a13cd2 |
| ベスト比 | -3202 (-1.4%) vs 20261031-163105-6f2cd6f = 223742 |
| note | jpegcheck の AC を 10 ビット表で読み飛ばす（main c001e38 を取り込み） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	151	395	840	841
europe	7	19	54	158	439	438
america	6	21	65	194	533	532
```

```
{"id":"29192","status":"done","portalStatus":"finished","passed":true,"score":220540,"raw":220540,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:36:40.222Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:36:43.210Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:36:46.205Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:36:49.219Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T07:36:51.826Z level=INFO msg="user moved" user=roroto_5185 from=asia to=america
time=2026-10-31T07:36:52.205Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T07:36:55.205Z level=INFO msg="america 2 日目: 21 人"
time=2026-10-31T07:36:55.790Z level=INFO msg="user moved" user=hasani_5405 from=europe to=asia
time=2026-10-31T07:36:58.216Z level=INFO msg="asia 3 日目: 151 人"
time=2026-10-31T07:37:01.207Z level=INFO msg="europe 3 日目: 54 人"
time=2026-10-31T07:37:02.721Z level=INFO msg="user moved" user=yunadan-5943 from=asia to=america
time=2026-10-31T07:37:03.293Z level=INFO msg="user moved" user=yunadan-5943 from=america to=europe
time=2026-10-31T07:37:03.807Z level=INFO msg="user moved" user=merusa-6337 from=europe to=america
time=2026-10-31T07:37:04.205Z level=INFO msg="america 3 日目: 65 人"
time=2026-10-31T07:37:05.108Z level=INFO msg="user moved" user=hasani_5405 from=asia to=america
time=2026-10-31T07:37:05.468Z level=INFO msg="user moved" user=elel5932 from=asia to=america
time=2026-10-31T07:37:06.061Z level=INFO msg="user moved" user=tsudansa5991 from=asia to=america
time=2026-10-31T07:37:07.219Z level=INFO msg="asia 4 日目: 395 人"
time=2026-10-31T07:37:08.250Z level=INFO msg="user moved" user=neyone.6319 from=europe to=asia
time=2026-10-31T07:37:08.883Z level=INFO msg="user moved" user=sene6704 from=america to=asia
time=2026-10-31T07:37:09.806Z level=INFO msg="user moved" user=roroto_5185 from=america to=europe
time=2026-10-31T07:37:10.206Z level=INFO msg="europe 4 日目: 158 人"
time=2026-10-31T07:37:11.822Z level=INFO msg="user moved" user=noroben7626 from=asia to=america
time=2026-10-31T07:37:12.637Z level=INFO msg="user moved" user=sene6704 from=asia to=europe
time=2026-10-31T07:37:12.639Z level=INFO msg="user moved" user=kean.8465 from=asia to=america
time=2026-10-31T07:37:12.751Z level=INFO msg="user moved" user=noroben7626 from=america to=europe
time=2026-10-31T07:37:13.218Z level=INFO msg="america 4 日目: 194 人"
time=2026-10-31T07:37:13.369Z level=INFO msg="user moved" user=zoemu7907 from=asia to=europe
time=2026-10-31T07:37:14.484Z level=INFO msg="user moved" user=yunadan-5943 from=europe to=america
time=2026-10-31T07:37:15.049Z level=INFO msg="user moved" user=martasu.8999 from=europe to=asia
time=2026-10-31T07:37:15.330Z level=INFO msg="user moved" user=shiheis_7190 from=asia to=europe
time=2026-10-31T07:37:15.520Z level=INFO msg="user moved" user=seyo_8089 from=asia to=europe
time=2026-10-31T07:37:15.771Z level=INFO msg="user moved" user=kaelivy_9077 from=europe to=asia
time=2026-10-31T07:37:15.965Z level=INFO msg="user moved" user=wator-7024 from=asia to=america
time=2026-10-31T07:37:16.222Z level=INFO msg="asia 5 日目: 840 人"
time=2026-10-31T07:37:16.548Z level=INFO msg="user moved" user=niha9059 from=europe to=america
time=2026-10-31T07:37:19.216Z level=INFO msg="europe 5 日目: 439 人"
time=2026-10-31T07:37:20.417Z level=INFO msg="user moved" user=tsudansa5991 from=america to=asia
time=2026-10-31T07:37:20.512Z level=INFO msg="user moved" user=elel5932 from=america to=asia
time=2026-10-31T07:37:22.488Z level=INFO msg="america 5 日目: 533 人"
time=2026-10-31T07:37:31.252Z level=INFO msg="europe 最終日（グランプリ）: 438 人"
time=2026-10-31T07:37:31.285Z level=INFO msg="america 最終日（グランプリ）: 532 人"
time=2026-10-31T07:37:31.462Z level=INFO msg="asia 最終日（グランプリ）: 841 人"
time=2026-10-31T07:37:40.345Z level=INFO msg="グランプリ: アジア 454 人, ヨーロッパ 356 人, アメリカ 256 人"
{
  "pass": true,
  "score": 220540,
  "messages": [
    "isyois_5774: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maxnina6381: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yuxin_wu_160: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "hobakbit_sujokgwan_6: 履歴を読めません: GET /api/memories?month=2026-02: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
47.0%	1499	1	0	GET	/api/notifications	15098.283	10.072	15.310	15.316	172
9.4%	13101	0	0	GET	^/api/feeds/[^/]+	3020.290	0.231	1.478	3.070	5140
7.8%	2970	0	0	GET	/api/memories	2511.992	0.846	7.376	14.854	10965
7.7%	2307	86	0	POST	/api/posts	2461.716	1.067	7.366	7.929	628
5.1%	3224	0	0	GET	/api/posting-window/current	1625.741	0.504	1.667	2.153	237
3.0%	2344	0	0	GET	/api/users	955.636	0.408	1.693	2.026	241
2.7%	220	0	0	GET	/api/me/export	863.222	3.924	20.065	24.705	1882373
2.5%	1855	1	0	GET	^/api/posts/me/current$	809.661	0.436	1.649	1.968	652
2.0%	595	130	0	PUT	/api/me/profile-image	638.533	1.073	15.800	18.037	3
2.0%	12191	250	0	GET	^/api/images/[^/]+	635.002	0.052	0.879	2.181	182259
1.9%	4917	60	0	POST	/api/friend-requests	620.295	0.126	1.207	2.111	114
1.3%	1002	8	0	GET	^/api/posts/[^/]+/realmojis$	404.016	0.403	3.271	3.826	2011
1.2%	3716	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	379.314	0.102	1.171	1.838	116
1.1%	482	3	0	GET	/api/me	343.964	0.714	1.667	1.722	300
0.9%	717	3	0	POST	^/api/posts/[^/]+/realmojis$	283.177	0.395	3.215	3.343	478
0.8%	1540	0	0	GET	^/api/relationships/[^/]+	269.916	0.175	1.228	4.793	1354
0.8%	2316	92	0	GET	^/api/posts/[^/]+$	249.148	0.108	0.653	1.242	596
0.6%	566	0	0	GET	^/api/memories/[^/]+	207.912	0.367	1.892	2.462	660
0.4%	4053	488	0	POST	/api/auth/login	143.617	0.035	0.962	1.639	216
0.4%	1129	0	0	GET	^/api/users/[^/]+$	122.030	0.108	0.786	1.233	193
0.4%	841	0	0	POST	^/api/posts/[^/]+/comments$	121.222	0.144	1.319	1.962	346
0.4%	934	2	0	GET	^/api/posts/[^/]+/comments$	120.067	0.129	0.980	1.228	3009
0.2%	96	0	0	DELETE	^/api/posts/[^/]+$	56.472	0.588	2.122	2.122	0
0.1%	539	0	0	POST	/api/auth/logout	40.000	0.074	1.196	1.251	0
0.1%	213	4	0	PATCH	^/api/posts/[^/]+$	38.359	0.180	1.205	1.263	644
0.1%	116	0	0	DELETE	^/api/realmojis/[^/]+$	36.713	0.316	1.550	1.785	0
0.1%	1	0	0	POST	/api/initialize	22.248	22.248	22.248	22.248	90
0.1%	593	0	0	PUT	^/api/follows/[^/]+$	21.687	0.037	0.863	1.570	0
0.0%	176	0	0	DELETE	^/api/comments/[^/]+$	12.353	0.070	1.160	1.168	0
0.0%	863	1	0	POST	/api/auth/register	12.351	0.014	0.234	0.278	227
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
49.9%	709	0	0	GET	/api/notifications	5237.502	7.387	11.876	11.877	170
7.6%	2478	16	0	POST	/api/friend-requests	801.406	0.323	1.299	1.912	114
6.9%	6512	0	0	GET	^/api/feeds/[^/]+	725.912	0.111	1.091	1.542	3640
6.2%	6171	40	0	GET	^/api/images/[^/]+	653.328	0.106	1.097	3.179	174815
5.7%	1112	23	0	POST	/api/posts	603.295	0.543	7.426	7.919	636
3.8%	1800	0	0	POST	/api/auth/login	401.685	0.223	1.373	2.356	249
2.3%	1558	0	0	GET	/api/memories	241.096	0.155	1.218	2.320	13144
1.9%	427	13	0	POST	^/api/posts/[^/]+/realmojis$	196.763	0.461	3.424	3.571	464
1.9%	242	0	0	PUT	/api/me/profile-image	196.240	0.811	3.466	3.591	0
1.4%	1592	0	0	GET	/api/posting-window/current	151.286	0.095	0.896	1.238	239
1.4%	963	0	0	GET	^/api/posts/me/current$	142.295	0.148	1.164	1.286	648
1.2%	567	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	127.487	0.225	1.298	1.378	116
1.2%	1195	0	0	GET	/api/users	125.083	0.105	0.688	1.431	240
1.1%	437	0	0	POST	/api/auth/register	117.975	0.270	1.238	1.526	229
1.0%	307	0	0	PUT	^/api/follows/[^/]+$	103.207	0.336	1.275	1.373	0
0.9%	396	4	0	POST	^/api/posts/[^/]+/comments$	93.386	0.236	1.443	2.277	335
0.7%	972	20	0	GET	^/api/posts/[^/]+$	78.548	0.081	0.277	1.235	607
0.7%	770	0	0	GET	^/api/relationships/[^/]+	76.182	0.099	1.032	1.542	592
0.7%	202	0	0	PATCH	/api/me	75.590	0.374	1.514	2.151	281
0.5%	295	0	0	POST	/api/auth/logout	55.369	0.188	1.049	1.217	0
0.4%	104	0	0	PATCH	^/api/posts/[^/]+$	40.304	0.388	1.349	1.448	647
0.3%	446	0	0	GET	^/api/users/[^/]+$	36.023	0.081	0.283	1.259	197
0.3%	418	1	0	GET	^/api/posts/[^/]+/realmojis$	32.403	0.078	0.299	1.024	2970
0.3%	395	0	0	GET	^/api/posts/[^/]+/comments$	31.201	0.079	0.270	1.102	3807
0.3%	43	0	0	DELETE	^/api/posts/[^/]+$	30.385	0.707	2.261	2.261	0
0.3%	277	0	0	GET	^/api/memories/[^/]+	27.080	0.098	1.085	1.951	975
0.3%	249	2	0	GET	/api/me	26.302	0.106	0.288	0.996	302
0.2%	59	0	0	DELETE	^/api/realmojis/[^/]+$	23.104	0.392	0.877	0.877	0
0.2%	118	0	0	GET	/api/me/export	20.328	0.172	1.995	5.520	2280977
0.1%	76	0	0	DELETE	^/api/comments/[^/]+$	13.495	0.178	1.256	1.256	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.3%	835	0	0	GET	/api/notifications	4866.295	5.828	9.482	9.484	175
8.9%	1286	53	0	POST	/api/posts	1004.565	0.781	7.373	7.856	624
7.5%	5005	0	0	GET	^/api/feeds/[^/]+	847.724	0.169	1.458	1.918	3167
6.9%	5914	36	0	GET	^/api/images/[^/]+	779.415	0.132	1.253	3.014	160574
6.2%	2878	20	0	POST	/api/friend-requests	691.290	0.240	0.774	1.596	114
4.8%	1818	0	0	POST	/api/auth/login	543.681	0.299	1.481	2.215	248
3.1%	527	0	0	POST	/api/auth/register	350.330	0.665	1.544	1.580	229
2.8%	1882	0	0	GET	/api/posting-window/current	315.968	0.168	1.012	1.968	239
2.5%	1483	0	0	GET	/api/memories	281.463	0.190	1.787	2.445	16066
1.9%	1431	0	0	GET	/api/users	213.976	0.150	0.814	1.539	246
1.8%	970	60	0	GET	^/api/posts/me/current$	205.575	0.212	1.449	1.875	609
1.4%	632	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	157.626	0.249	1.392	2.371	116
1.3%	271	0	0	PUT	/api/me/profile-image	140.433	0.518	3.519	3.520	0
0.9%	225	15	0	POST	^/api/posts/[^/]+/realmojis$	98.847	0.439	1.653	1.963	430
0.9%	639	0	0	GET	^/api/relationships/[^/]+	96.822	0.152	1.387	1.717	619
0.8%	347	0	0	PUT	^/api/follows/[^/]+$	91.736	0.264	1.252	1.337	0
0.7%	682	26	0	GET	^/api/posts/[^/]+$	82.948	0.122	0.659	1.429	595
0.7%	243	0	0	PATCH	/api/me	73.018	0.300	1.326	1.346	275
0.6%	287	0	0	POST	/api/auth/logout	62.245	0.217	1.383	1.653	0
0.5%	252	5	0	POST	^/api/posts/[^/]+/comments$	58.916	0.234	0.679	0.715	338
0.4%	341	0	0	GET	^/api/users/[^/]+$	44.570	0.131	0.645	0.835	187
0.3%	98	20	0	PATCH	^/api/posts/[^/]+$	34.073	0.348	1.287	1.287	522
0.3%	259	4	0	GET	^/api/posts/[^/]+/realmojis$	33.707	0.130	0.691	0.748	2121
0.3%	245	1	0	GET	^/api/posts/[^/]+/comments$	30.400	0.124	0.751	0.845	3443
0.3%	42	0	0	DELETE	^/api/posts/[^/]+$	28.780	0.685	2.257	2.257	0
0.2%	236	0	0	GET	^/api/memories/[^/]+	24.704	0.105	0.600	0.903	1052
0.2%	278	1	0	GET	/api/me	21.969	0.079	0.351	0.401	304
0.2%	95	0	0	GET	/api/me/export	20.340	0.214	2.746	2.746	2861775
0.1%	35	0	0	DELETE	^/api/realmojis/[^/]+$	13.710	0.392	0.904	0.904	0
0.1%	29	0	0	DELETE	^/api/comments/[^/]+$	6.578	0.227	0.593	0.593	0
…（上位 30 / 全 36 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 11116 クエリ / 48 種類 / 合計 130.301s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	41.9%	223	54.555	244.64	802.89	1211.1	123.3	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	19.9%	4	25.975	6493.78	9549.38	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.4%	4	18.822	4705.40	7203.69	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.6%	7116	11.152	1.57	13.20	10.3	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.9%	4	8.939	2234.81	3415.42	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	5.6%	7	7.253	1036.20	3890.81	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.5%	7	0.618	88.28	157.52	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
8	0.5%	7	0.613	87.63	159.87	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.4%	1	0.553	552.86	552.86	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.4%	1	0.503	502.61	502.61	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.2%	1	0.235	234.90	234.90	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.164	163.95	163.95	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.139	139.08	139.08	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
14	0.1%	7	0.138	19.70	38.49	167.4	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	0.1%	1	0.118	118.23	118.23	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: cca27f47043982ab7863aaa499bbe78ebbaf3155
Type: cpu
Time: 2026-10-31 16:36:11 JST
Duration: 60.02s, Total samples = 24.47s (40.77%)
Showing nodes accounting for 18.69s, 76.38% of 24.47s total
Dropped 768 nodes (cum <= 0.12s)
Showing top 40 nodes out of 237
      flat  flat%   sum%        cum   cum%
     3.90s 15.94% 15.94%      3.90s 15.94%  crypto/internal/fips140/sha256.blockAVX2
     2.13s  8.70% 24.64%      2.13s  8.70%  internal/runtime/syscall/linux.Syscall6
     1.73s  7.07% 31.71%      5.96s 24.36%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.39s  5.68% 37.39%      1.39s  5.68%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.18s  4.82% 42.21%      1.18s  4.82%  runtime.memmove
     1.12s  4.58% 46.79%      1.12s  4.58%  aeshashbody
     0.67s  2.74% 49.53%      0.67s  2.74%  runtime.memclrNoHeapPointers
     0.50s  2.04% 51.57%      2.10s  8.58%  runtime.scanObject
     0.45s  1.84% 53.41%      0.75s  3.06%  runtime.findObject
     0.40s  1.63% 55.05%      1.92s  7.85%  runtime.mapaccess1_faststr
     0.38s  1.55% 56.60%      0.38s  1.55%  memeqbody
     0.36s  1.47% 58.07%      0.59s  2.41%  runtime.tryDeferToSpanScan
     0.35s  1.43% 59.50%      0.54s  2.21%  example.com/isureal-company/isureal.postLess
     0.29s  1.19% 60.69%      0.39s  1.59%  database/sql.convertAssignRows
     0.26s  1.06% 61.75%      0.26s  1.06%  runtime.futex
     0.25s  1.02% 62.77%      4.50s 18.39%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.23s  0.94% 63.71%      4.31s 17.61%  crypto/internal/fips140/sha256.(*Digest).Write
     0.22s   0.9% 64.61%      1.53s  6.25%  runtime.mapassign_faststr
     0.21s  0.86% 65.47%      0.21s  0.86%  runtime.nextFreeFast (inline)
     0.20s  0.82% 66.29%      0.27s  1.10%  example.com/isureal-company/isureal/jpegcheck.(*decoder).skipAC
     0.19s  0.78% 67.06%      0.21s  0.86%  runtime.(*mspan).writeHeapBitsSmall
     0.19s  0.78% 67.84%      0.85s  3.47%  runtime.mallocgcSmallScanNoHeader
     0.17s  0.69% 68.53%      0.17s  0.69%  sync/atomic.(*Int32).Add (inline)
     0.16s  0.65% 69.19%      0.16s  0.65%  example.com/isureal-company/isureal.rmLess
     0.16s  0.65% 69.84%      1.30s  5.31%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.15s  0.61% 70.45%      4.99s 20.39%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.15s  0.61% 71.07%      0.19s  0.78%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.13s  0.53% 71.60%      0.13s  0.53%  runtime.(*spanInlineMarkBits).init
     0.13s  0.53% 72.13%      1.70s  6.95%  runtime.mallocgc
     0.13s  0.53% 72.66%      0.13s  0.53%  runtime.typePointers.nextFast (inline)
     0.11s  0.45% 73.11%      2.04s  8.34%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s  0.45% 73.56%      0.18s  0.74%  internal/runtime/maps.(*Iter).Next
     0.11s  0.45% 74.01%      0.17s  0.69%  runtime.spanOf (inline)
     0.10s  0.41% 74.42%      4.71s 19.25%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.10s  0.41% 74.83%      3.83s 15.65%  example.com/isureal-company/isureal.loadBig
     0.09s  0.37% 75.19%      0.46s  1.88%  runtime.convTstring
     0.08s  0.33% 75.52%      0.39s  1.59%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.07s  0.29% 75.81%      0.20s  0.82%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.07s  0.29% 76.09%      3.97s 16.22%  crypto/internal/fips140/sha256.block
     0.07s  0.29% 76.38%      0.21s  0.86%  runtime.greyobject
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 32.1% / 100% | 11% | 0.7% | 6 | 179 MB | run queue が積もっている |
| isu2 (2 cpu) | 22.0% / 100% | 3% | 1.4% | 5 | 156 MB | run queue が積もっている |
| isu3 (2 cpu) | 23.4% / 100% | 8% | 1.0% | 5 | 121 MB | run queue が積もっている |
| isu4 (2 cpu) | 58.5% / 100% | 50% | 2.6% | 7 | 104 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 50% / nginx 10% |
| isu2 | isureal 38% / nginx 2% |
| isu3 | isureal 43% / nginx 2% |
| isu4 | isureal 96% / postgres 14% |

## Diff from previous run

スコア: 206779 → 220540 (+13761)

`8a13cd2..15a1e6c`

```
 webapp/go/jpegcheck/check_test.go |  16 +++++
 webapp/go/jpegcheck/huffman.go    |  56 +++++++++++++++++-
 webapp/go/jpegcheck/scan.go       | 121 ++++++++++++++++++++++++++++++--------
 3 files changed, 165 insertions(+), 28 deletions(-)
```
