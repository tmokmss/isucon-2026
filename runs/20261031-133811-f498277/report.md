# run 20261031-133811-f498277

## Summary

| | |
|---|---|
| score | 75212 |
| status | pass |
| ref | image-files |
| sha | f498277673d9 |
| 所要時間 | 485s |
| 前回比 | +10512 (16.2%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +10512 (16.2%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | image bodies to files (fix passive SQL) + main 7e687d9 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	46	131	361	467	467
europe	7	17	47	129	129	130
america	6	14	45	117	114	113
```

```
{"id":"14000","status":"done","portalStatus":"finished","passed":true,"score":75212,"raw":75212,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:44:45.472Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:44:48.442Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:44:51.443Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:44:54.468Z level=INFO msg="asia 2 日目: 46 人"
time=2026-10-31T04:44:57.454Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T04:45:00.486Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T04:45:03.462Z level=INFO msg="asia 3 日目: 131 人"
time=2026-10-31T04:45:06.473Z level=INFO msg="europe 3 日目: 47 人"
time=2026-10-31T04:45:06.995Z level=INFO msg="user moved" user=muyu5962 from=asia to=america
time=2026-10-31T04:45:07.037Z level=INFO msg="user moved" user=resa.5845 from=asia to=america
time=2026-10-31T04:45:07.043Z level=INFO msg="user moved" user=kemuko-5922 from=asia to=america
time=2026-10-31T04:45:07.190Z level=INFO msg="user moved" user=walialeo5859 from=asia to=europe
time=2026-10-31T04:45:08.340Z level=INFO msg="user moved" user=muyu5962 from=america to=europe
time=2026-10-31T04:45:09.564Z level=INFO msg="america 3 日目: 45 人"
time=2026-10-31T04:45:12.500Z level=INFO msg="asia 4 日目: 361 人"
time=2026-10-31T04:45:12.862Z level=INFO msg="user moved" user=miaulshi6163 from=europe to=asia
time=2026-10-31T04:45:16.209Z level=INFO msg="user moved" user=fuyoleo_6620 from=asia to=america
time=2026-10-31T04:45:17.715Z level=INFO msg="user moved" user=fuyoleo_6620 from=america to=europe
time=2026-10-31T04:45:18.048Z level=INFO msg="europe 4 日目: 129 人"
time=2026-10-31T04:45:19.047Z level=INFO msg="america 4 日目: 117 人"
time=2026-10-31T04:45:22.807Z level=INFO msg="user moved" user=muyu5962 from=europe to=america
time=2026-10-31T04:45:24.963Z level=INFO msg="user moved" user=walialeo5859 from=europe to=america
time=2026-10-31T04:45:25.434Z level=INFO msg="user moved" user=andanan8311 from=america to=asia
time=2026-10-31T04:45:26.369Z level=INFO msg="user moved" user=rokoso_8681 from=europe to=asia
time=2026-10-31T04:45:26.654Z level=INFO msg="user moved" user=kemuko-5922 from=america to=europe
time=2026-10-31T04:45:26.889Z level=INFO msg="user moved" user=resa.5845 from=america to=asia
time=2026-10-31T04:45:27.368Z level=INFO msg="user moved" user=ivyri-6404 from=america to=asia
time=2026-10-31T04:45:28.402Z level=INFO msg="user moved" user=mewamu-6366 from=america to=europe
time=2026-10-31T04:45:31.183Z level=INFO msg="user moved" user=kiyo-8670 from=america to=asia
time=2026-10-31T04:45:31.512Z level=INFO msg="user moved" user=ivykaki.7250 from=asia to=america
time=2026-10-31T04:45:32.048Z level=INFO msg="user moved" user=miaulshi6163 from=asia to=europe
time=2026-10-31T04:45:32.287Z level=INFO msg="europe 5 日目: 129 人"
time=2026-10-31T04:45:32.304Z level=INFO msg="user moved" user=isleo_6915 from=asia to=europe
time=2026-10-31T04:45:32.625Z level=INFO msg="america 5 日目: 114 人"
time=2026-10-31T04:45:32.807Z level=INFO msg="asia 5 日目: 467 人"
time=2026-10-31T04:45:38.317Z level=INFO msg="america 最終日（グランプリ）: 113 人"
time=2026-10-31T04:45:38.317Z level=INFO msg="europe 最終日（グランプリ）: 130 人"
time=2026-10-31T04:45:38.317Z level=INFO msg="asia 最終日（グランプリ）: 467 人"
time=2026-10-31T04:45:45.467Z level=INFO msg="グランプリ: アジア 74 人, ヨーロッパ 125 人, アメリカ 110 人"
{
  "pass": true,
  "score": 75212,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
52.3%	887	6	0	GET	/api/notifications	8224.710	9.273	19.745	19.748	173
9.0%	6307	931	0	GET	^/api/images/[^/]+	1419.478	0.225	3.104	7.653	118743
8.3%	843	3	0	POST	/api/posts	1313.203	1.558	7.103	7.103	636
4.3%	4088	1	0	GET	^/api/feeds/[^/]+	674.555	0.165	1.600	6.865	2564
3.6%	2022	29	0	POST	/api/friend-requests	559.909	0.277	5.332	6.865	114
3.2%	1281	1	0	GET	/api/posting-window/current	510.068	0.398	2.625	6.865	236
3.2%	105	0	0	GET	/api/me/export	508.838	4.846	20.402	23.356	2040794
2.8%	2791	982	0	POST	/api/auth/login	443.174	0.159	2.625	6.865	166
2.6%	326	120	0	PUT	/api/me/profile-image	414.540	1.272	13.362	15.201	6
2.3%	1569	0	0	GET	/api/memories	355.152	0.226	2.197	7.217	10386
1.8%	944	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	278.195	0.295	4.228	6.865	116
1.6%	662	59	0	GET	^/api/posts/me/current$	259.150	0.391	6.865	6.865	587
0.8%	969	0	0	GET	/api/users	133.169	0.137	1.572	6.865	240
0.5%	172	1	0	POST	^/api/posts/[^/]+/realmojis$	76.196	0.443	1.975	2.195	460
0.4%	485	18	0	POST	/api/auth/register	67.229	0.139	2.197	4.554	219
0.4%	691	0	0	GET	^/api/relationships/[^/]+	64.643	0.094	0.840	6.865	1360
0.4%	546	13	0	GET	^/api/posts/[^/]+$	64.196	0.118	0.886	6.865	586
0.4%	248	0	0	PUT	^/api/follows/[^/]+$	64.071	0.258	3.924	6.865	0
0.3%	283	0	0	POST	/api/auth/logout	49.522	0.175	3.344	5.332	0
0.3%	264	0	0	GET	^/api/memories/[^/]+	45.590	0.173	2.138	6.865	692
0.3%	182	0	0	PATCH	/api/me	41.098	0.226	4.147	4.147	267
0.2%	217	0	0	POST	^/api/posts/[^/]+/comments$	36.433	0.168	1.388	1.983	311
0.2%	214	2	0	GET	/api/me	27.361	0.128	0.950	1.875	298
0.1%	333	0	0	GET	^/api/users/[^/]+$	22.176	0.067	0.517	0.728	181
0.1%	1	0	0	POST	/api/initialize	21.376	21.376	21.376	21.376	90
0.1%	76	7	0	PATCH	^/api/posts/[^/]+$	18.306	0.241	5.587	5.587	577
0.0%	233	0	0	GET	^/api/posts/[^/]+/realmojis$	7.447	0.032	0.423	0.446	757
0.0%	31	0	0	DELETE	^/api/posts/[^/]+$	7.392	0.238	1.208	1.208	0
0.0%	41	0	0	DELETE	^/api/realmojis/[^/]+$	7.194	0.175	0.733	0.733	0
0.0%	188	0	0	GET	^/api/posts/[^/]+/comments$	7.185	0.038	0.352	0.423	1137
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.6%	356	0	0	GET	/api/notifications	2923.193	8.211	14.340	14.340	178
13.6%	2520	18	0	GET	^/api/images/[^/]+	869.260	0.345	2.559	4.082	179234
13.2%	454	0	0	POST	/api/posts	848.275	1.868	5.339	5.889	650
6.0%	2504	0	0	GET	^/api/feeds/[^/]+	384.751	0.154	0.798	1.900	2823
3.1%	812	0	0	GET	/api/memories	197.366	0.243	1.239	1.840	9259
3.0%	58	0	0	GET	/api/me/export	194.602	3.355	17.324	17.324	1793922
2.8%	730	0	0	POST	/api/auth/login	180.519	0.247	3.410	4.405	250
2.1%	605	0	0	GET	/api/posting-window/current	132.047	0.218	1.532	2.335	239
1.8%	748	15	0	POST	/api/friend-requests	114.520	0.153	0.689	4.649	113
1.6%	445	10	0	GET	^/api/posts/me/current$	101.234	0.227	1.171	1.854	637
0.8%	378	0	0	GET	/api/users	53.465	0.141	0.727	0.830	240
0.8%	406	4	0	GET	^/api/posts/[^/]+$	49.973	0.123	0.495	0.623	602
0.8%	318	0	0	GET	^/api/relationships/[^/]+	48.142	0.151	0.815	0.839	676
0.6%	125	0	0	POST	/api/auth/register	38.210	0.306	1.299	1.302	228
0.6%	252	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	37.956	0.151	0.638	0.847	116
0.6%	67	0	0	PUT	/api/me/profile-image	37.589	0.561	1.912	1.912	0
0.5%	141	0	0	POST	/api/auth/logout	34.726	0.246	2.935	4.248	0
0.4%	49	1	0	POST	^/api/posts/[^/]+/realmojis$	25.588	0.522	1.599	1.599	475
0.4%	190	0	0	GET	^/api/users/[^/]+$	24.557	0.129	0.737	0.841	194
0.3%	24	0	0	DELETE	^/api/posts/[^/]+$	20.487	0.854	3.248	3.248	0
0.3%	108	0	0	GET	^/api/memories/[^/]+	19.731	0.183	0.816	1.540	548
0.2%	94	0	0	PUT	^/api/follows/[^/]+$	14.341	0.153	0.974	0.974	0
0.2%	72	0	0	PATCH	/api/me	14.119	0.196	1.052	1.052	272
0.2%	42	8	0	PATCH	^/api/posts/[^/]+$	12.327	0.293	1.346	1.346	524
0.2%	71	1	0	GET	/api/me	9.678	0.136	0.374	0.374	298
0.1%	67	0	0	GET	^/api/posts/[^/]+/comments$	8.147	0.122	0.511	0.511	961
0.1%	68	0	0	GET	^/api/posts/[^/]+/realmojis$	7.955	0.117	0.507	0.507	496
0.1%	38	0	0	POST	^/api/posts/[^/]+/comments$	6.094	0.160	1.501	1.501	325
0.0%	6	0	0	DELETE	^/api/realmojis/[^/]+$	1.450	0.242	0.301	0.301	0
0.0%	9	0	0	DELETE	^/api/comments/[^/]+$	1.089	0.121	0.174	0.174	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.0%	315	1	0	GET	/api/notifications	2429.178	7.712	13.950	13.950	175
13.1%	397	0	0	POST	/api/posts	724.344	1.825	5.304	5.674	662
13.0%	2012	10	0	GET	^/api/images/[^/]+	718.425	0.357	2.602	4.362	216920
6.9%	2597	0	0	GET	^/api/feeds/[^/]+	379.616	0.146	0.688	2.930	3167
4.8%	34	0	0	GET	/api/me/export	262.793	7.729	20.955	20.955	2932301
2.8%	633	0	0	GET	/api/memories	155.080	0.245	1.349	2.920	10596
2.2%	497	0	0	POST	/api/auth/login	122.583	0.247	3.488	4.532	252
1.9%	701	10	0	POST	/api/friend-requests	102.460	0.146	0.641	4.200	114
1.8%	529	0	0	GET	/api/posting-window/current	97.878	0.185	0.904	1.673	240
1.6%	392	8	0	GET	^/api/posts/me/current$	89.525	0.228	1.168	1.587	650
1.0%	450	7	0	GET	^/api/posts/[^/]+$	55.049	0.122	0.388	1.136	599
0.9%	308	0	0	GET	^/api/relationships/[^/]+	46.931	0.152	0.717	0.821	693
0.8%	322	0	0	GET	/api/users	43.987	0.137	0.644	0.717	244
0.7%	74	0	0	PUT	/api/me/profile-image	40.852	0.552	1.355	1.355	0
0.7%	229	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	37.871	0.165	0.741	3.596	116
0.6%	53	1	0	POST	^/api/posts/[^/]+/realmojis$	33.073	0.624	1.357	1.357	492
0.6%	114	0	0	POST	/api/auth/register	33.069	0.290	1.567	1.665	229
0.5%	208	0	0	GET	^/api/users/[^/]+$	25.927	0.125	0.411	0.640	191
0.4%	79	0	0	POST	/api/auth/logout	21.812	0.276	4.360	4.360	0
0.3%	73	0	0	GET	^/api/memories/[^/]+	15.336	0.210	0.638	0.638	1094
0.2%	15	0	0	DELETE	^/api/posts/[^/]+$	12.211	0.814	4.677	4.677	0
0.2%	86	0	0	PUT	^/api/follows/[^/]+$	11.815	0.137	0.938	0.938	0
0.2%	36	5	0	PATCH	^/api/posts/[^/]+$	10.852	0.301	2.389	2.389	585
0.2%	79	0	0	GET	^/api/posts/[^/]+/realmojis$	9.803	0.124	0.384	0.384	369
0.2%	78	1	0	GET	/api/me	9.590	0.123	0.316	0.316	300
0.2%	68	0	0	PATCH	/api/me	9.118	0.134	0.281	0.281	282
0.2%	40	0	0	POST	^/api/posts/[^/]+/comments$	8.881	0.222	2.899	2.899	304
0.1%	58	0	0	GET	^/api/posts/[^/]+/comments$	7.256	0.125	0.412	0.412	827
0.1%	12	0	0	DELETE	^/api/realmojis/[^/]+$	2.991	0.249	0.364	0.364	0
0.0%	7	0	0	PUT	^/api/blocks/[^/]+$	0.808	0.115	0.141	0.141	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 38437 クエリ / 79 種類 / 合計 25.123s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	24.9%	3525	6.264	1.78	12.47	97.4	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $2) AS is_profile, EXISTS (SELECT $3 FROM posts p WHERE p.image_id…
2	22.2%	2	5.569	2784.67	5421.03	240457.5	2832.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
3	10.4%	1	2.618	2617.50	2617.50	30823.0	1199695.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
4	8.8%	1691	2.211	1.31	7.21	25.9	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
5	7.6%	1	1.906	1906.19	1906.19	23889.0	628094.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
6	5.2%	1	1.313	1313.31	1313.31	12529.0	530149.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
7	3.6%	1	0.898	897.78	897.78	29856.0	2983.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	2.6%	271	0.650	2.40	7.06	32.9	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	1.7%	3399	0.436	0.13	10.11	20.2	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	1.5%	295	0.385	1.31	7.42	23.7	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	1.1%	2	0.285	142.43	277.15	16782.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	1.1%	3042	0.281	0.09	5.87	11.7	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
13	1.0%	1	0.248	247.58	247.58	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	1.0%	1	0.242	242.05	242.05	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
15	0.9%	2309	0.217	0.09	5.68	8.9	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$4::bytea,$5,$3)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $2) AS is_profile, EXISTS (SELECT $3 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $4 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
-- rank 2
SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id;
-- rank 3
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 7cc013328fd4ea09b4c29a087fc389f4a422d022
Type: cpu
Time: 2026-10-31 13:44:19 JST
Duration: 60.39s, Total samples = 30.67s (50.79%)
Showing nodes accounting for 22.80s, 74.34% of 30.67s total
Dropped 858 nodes (cum <= 0.15s)
Showing top 40 nodes out of 281
      flat  flat%   sum%        cum   cum%
     4.06s 13.24% 13.24%      4.06s 13.24%  internal/runtime/syscall/linux.Syscall6
     1.91s  6.23% 19.47%      1.91s  6.23%  crypto/internal/fips140/sha256.blockAVX2
     1.60s  5.22% 24.68%      6.37s 20.77%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.59s  5.18% 29.87%      1.59s  5.18%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.56s  5.09% 34.95%      1.56s  5.09%  runtime.memclrNoHeapPointers
     1.12s  3.65% 38.60%      1.12s  3.65%  aeshashbody
     1.03s  3.36% 41.96%      1.37s  4.47%  runtime.findObject
     0.99s  3.23% 45.19%      4.11s 13.40%  runtime.scanObject
     0.97s  3.16% 48.35%      0.97s  3.16%  runtime.memmove
     0.83s  2.71% 51.06%      1.40s  4.56%  runtime.tryDeferToSpanScan
     0.59s  1.92% 52.98%      0.59s  1.92%  example.com/isureal-company/isureal.(*UserRow).active
     0.57s  1.86% 54.84%      0.57s  1.86%  memeqbody
     0.49s  1.60% 56.44%      2.17s  7.08%  runtime.mapaccess1_faststr
     0.49s  1.60% 58.04%      0.49s  1.60%  runtime.spanClass.sizeclass (inline)
     0.36s  1.17% 59.21%      0.36s  1.17%  runtime.nextFreeFast (inline)
     0.31s  1.01% 60.22%      0.61s  1.99%  example.com/isureal-company/isureal.postLess
     0.31s  1.01% 61.23%      1.30s  4.24%  runtime.mallocgcSmallScanNoHeader
     0.30s  0.98% 62.21%      0.38s  1.24%  database/sql.convertAssignRows
     0.30s  0.98% 63.19%      0.30s  0.98%  indexbody
     0.23s  0.75% 63.94%      2.78s  9.06%  runtime.mallocgc
     0.22s  0.72% 64.66%      0.25s  0.82%  runtime.(*mspan).writeHeapBitsSmall
     0.22s  0.72% 65.37%      0.22s  0.72%  runtime.futex
     0.22s  0.72% 66.09%      0.26s  0.85%  runtime.typePointers.next
     0.22s  0.72% 66.81%      0.22s  0.72%  sync/atomic.(*Int32).Add (inline)
     0.21s  0.68% 67.49%      0.26s  0.85%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.20s  0.65% 68.14%      1.71s  5.58%  runtime.mapassign_faststr
     0.19s  0.62% 68.76%      0.19s  0.62%  cmpbody
     0.18s  0.59% 69.35%      0.19s  0.62%  runtime.spanOf (inline)
     0.17s  0.55% 69.91%      1.36s  4.43%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.15s  0.49% 70.39%      0.17s  0.55%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
     0.15s  0.49% 70.88%      0.30s  0.98%  runtime.scanObjectsSmall
     0.14s  0.46% 71.34%      0.39s  1.27%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.14s  0.46% 71.80%      2.54s  8.28%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.14s  0.46% 72.25%      0.31s  1.01%  runtime.greyobject
     0.13s  0.42% 72.68%      0.28s  0.91%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.13s  0.42% 73.10%      0.23s  0.75%  internal/runtime/maps.(*Iter).Next
     0.12s  0.39% 73.49%      0.49s  1.60%  runtime.mapaccess2_faststr
     0.09s  0.29% 73.79%      4.64s 15.13%  example.com/isureal-company/isureal.loadBig
     0.09s  0.29% 74.08%      0.47s  1.53%  runtime.mallocgcSmallNoscan
     0.08s  0.26% 74.34%      0.46s  1.50%  database/sql.(*Rows).scanLocked
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 84.3% / 100% | 61% | 37.8% | 12 | 101 MB | CPU が飽和 |
| isu2 (2 cpu) | 13.9% / 98% | 1% | 0.5% | 3 | 1917 MB |  |
| isu3 (2 cpu) | 11.1% / 99% | 1% | 0.7% | 3 | 1660 MB |  |
| isu4 (2 cpu) | 12.5% / 63% | 0% | 5.8% | 17 | 324 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 50% / find 14% / nginx 6% / kswapd0 4% / journalctl 2% |
| isu2 | isureal 25% / nginx 1% |
| isu3 | isureal 19% / nginx 2% |
| isu4 | postgres 10% |

## Diff from previous run

スコア: 64700 → 75212 (+10512)

`e0e4cb3..f498277`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf     |   3 +
 conf/isu1/home/isucon/env.sh                       |   1 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf     |  14 +
 conf/isu2/home/isucon/env.sh                       |   2 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf     |  14 +
 conf/isu3/home/isucon/env.sh                       |   2 +
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
 webapp/go/image_local.go                           |  43 ++-
 webapp/go/initialize.go                            |   1 +
 webapp/go/offload.go                               |  27 +-
 webapp/go/post_api.go                              |  69 ++++-
 webapp/go/post_forward.go                          | 113 +++++++
 webapp/go/post_forward_test.go                     |  94 ++++++
 webapp/go/profile_image_api.go                     |   5 +-
 webapp/go/realmoji_api.go                          |   7 +-
 webapp/go/relationship.go                          |  10 +-
 webapp/go/server.go                                |   1 +
 webapp/go/store.go                                 | 113 ++++++-
 27 files changed, 1047 insertions(+), 152 deletions(-)
```
