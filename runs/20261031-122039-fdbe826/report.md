# run 20261031-122039-fdbe826

## Summary

| | |
|---|---|
| score | 65904 |
| status | pass |
| ref | origin/main |
| sha | fdbe82670a24 |
| 所要時間 | 113s |
| 前回比 | +36365 (123.1%) vs 20261031-121631-55a9d63 |
| ベスト比 | +5703 (9.5%) vs 20261031-115040-1c73ee4 = 60201 **ベスト更新** |
| note | 低い版の試走の直後の main |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	151	431	429	429
europe	7	15	41	91	93	94
america	8	14	34	57	55	54
```

```
{"id":"7686","status":"done","portalStatus":"finished","passed":true,"score":65904,"raw":65904,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:21:01.268Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:21:02.519Z level=INFO msg="user moved" user=kekiben5044 from=asia to=america
time=2026-10-31T03:21:03.908Z level=INFO msg="user moved" user=saeva_5023 from=asia to=america
time=2026-10-31T03:21:04.238Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:21:07.234Z level=INFO msg="america 1 日目: 8 人"
time=2026-10-31T03:21:10.257Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T03:21:13.244Z level=INFO msg="europe 2 日目: 15 人"
time=2026-10-31T03:21:16.235Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T03:21:19.256Z level=INFO msg="asia 3 日目: 151 人"
time=2026-10-31T03:21:20.006Z level=INFO msg="user moved" user=mewa_5506 from=asia to=europe
time=2026-10-31T03:21:22.263Z level=INFO msg="europe 3 日目: 41 人"
time=2026-10-31T03:21:22.326Z level=INFO msg="user moved" user=mewa_5506 from=europe to=asia
time=2026-10-31T03:21:24.064Z level=INFO msg="user moved" user=benmiwa-6046 from=asia to=america
time=2026-10-31T03:21:24.197Z level=INFO msg="user moved" user=kekiben5044 from=america to=asia
time=2026-10-31T03:21:25.147Z level=INFO msg="user moved" user=kekiben5044 from=asia to=america
time=2026-10-31T03:21:25.257Z level=INFO msg="america 3 日目: 34 人"
time=2026-10-31T03:21:28.300Z level=INFO msg="asia 4 日目: 431 人"
time=2026-10-31T03:21:29.977Z level=INFO msg="user moved" user=chievane_6374 from=europe to=america
time=2026-10-31T03:21:31.018Z level=INFO msg="user moved" user=mewa_5506 from=asia to=europe
time=2026-10-31T03:21:31.243Z level=INFO msg="europe 4 日目: 91 人"
time=2026-10-31T03:21:34.288Z level=INFO msg="america 4 日目: 57 人"
time=2026-10-31T03:21:36.015Z level=INFO msg="user moved" user=hiha-6587 from=asia to=europe
time=2026-10-31T03:21:37.637Z level=INFO msg="asia 5 日目: 429 人"
time=2026-10-31T03:21:37.734Z level=INFO msg="user moved" user=chievane_6374 from=america to=europe
time=2026-10-31T03:21:41.247Z level=INFO msg="europe 5 日目: 93 人"
time=2026-10-31T03:21:42.999Z level=INFO msg="user moved" user=mukomia.7459 from=asia to=europe
time=2026-10-31T03:21:43.475Z level=INFO msg="user moved" user=shinesol_6703 from=america to=asia
time=2026-10-31T03:21:43.554Z level=INFO msg="america 5 日目: 55 人"
time=2026-10-31T03:21:52.496Z level=INFO msg="europe 最終日（グランプリ）: 94 人"
time=2026-10-31T03:21:52.527Z level=INFO msg="america 最終日（グランプリ）: 54 人"
time=2026-10-31T03:21:52.593Z level=INFO msg="asia 最終日（グランプリ）: 429 人"
time=2026-10-31T03:22:01.241Z level=INFO msg="グランプリ: アジア 15 人, ヨーロッパ 1 人, アメリカ 0 人"
{
  "pass": true,
  "score": 65904,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.8%	981	0	0	GET	/api/notifications	9121.127	9.298	14.950	15.287	175
21.7%	1314	41	0	POST	/api/posts	5692.500	4.332	14.366	15.116	633
11.6%	9647	1016	0	GET	^/api/images/[^/]+	3030.448	0.314	6.151	15.091	126838
8.0%	2368	59	0	POST	/api/friend-requests	2088.410	0.882	7.741	12.898	113
6.9%	3849	1216	0	POST	/api/auth/login	1821.687	0.473	7.451	14.659	177
3.2%	1096	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	834.363	0.761	8.347	14.994	116
2.2%	1785	0	0	GET	/api/posting-window/current	568.569	0.319	6.861	6.925	236
1.7%	521	13	0	POST	^/api/posts/[^/]+/realmojis$	443.471	0.851	7.639	9.524	483
1.4%	383	131	0	PUT	/api/me/profile-image	364.059	0.951	7.727	10.403	8
1.3%	450	1	0	POST	/api/auth/register	347.629	0.773	3.496	12.022	227
1.0%	440	4	0	POST	^/api/posts/[^/]+/comments$	272.076	0.618	5.604	9.826	343
1.0%	336	0	0	PUT	^/api/follows/[^/]+$	260.913	0.777	8.222	11.525	0
0.9%	2317	0	0	GET	/api/memories	248.173	0.107	0.659	3.749	11158
0.8%	6963	0	0	GET	^/api/feeds/[^/]+	214.788	0.031	0.166	8.249	2569
0.6%	223	0	0	PATCH	/api/me	156.907	0.704	6.672	8.645	275
0.5%	161	0	0	GET	/api/me/export	138.998	0.863	12.507	13.966	1917687
0.5%	461	0	0	POST	/api/auth/logout	126.061	0.273	5.190	13.624	0
0.3%	106	5	0	PATCH	^/api/posts/[^/]+$	77.949	0.735	4.454	5.211	616
0.2%	66	0	0	DELETE	^/api/comments/[^/]+$	57.668	0.874	9.868	9.868	0
0.2%	88	0	0	DELETE	^/api/realmojis/[^/]+$	53.764	0.611	5.566	5.566	0
0.2%	47	0	0	DELETE	^/api/posts/[^/]+$	44.112	0.939	3.655	3.655	0
0.2%	953	49	0	GET	^/api/posts/me/current$	43.913	0.046	0.498	4.930	619
0.1%	1137	0	0	GET	/api/users	38.625	0.034	0.187	0.489	244
0.1%	1133	15	0	GET	^/api/posts/[^/]+$	30.681	0.027	0.151	0.284	621
0.1%	429	0	0	GET	^/api/memories/[^/]+	30.559	0.071	0.358	3.752	676
0.1%	980	0	0	GET	^/api/relationships/[^/]+	23.601	0.024	0.152	0.836	1060
0.1%	520	0	0	GET	^/api/users/[^/]+$	13.133	0.025	0.147	0.181	202
0.0%	14	0	0	PUT	^/api/admin/users/[^/]+/certification$	12.861	0.919	9.900	9.900	0
0.0%	441	1	0	GET	^/api/posts/[^/]+/comments$	11.645	0.026	0.168	0.459	1366
0.0%	1	0	0	POST	/api/initialize	11.529	11.529	11.529	11.529	90
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.7%	214	0	0	GET	/api/notifications	1748.412	8.170	11.346	11.819	175
22.6%	270	5	0	POST	/api/posts	1206.128	4.467	14.742	14.845	628
9.5%	516	6	0	POST	/api/friend-requests	507.823	0.984	8.800	12.585	114
7.7%	348	0	0	POST	/api/auth/login	409.532	1.177	12.571	13.247	249
3.8%	1244	6	0	GET	^/api/images/[^/]+	204.015	0.164	1.052	1.309	181569
3.5%	1179	0	0	GET	^/api/feeds/[^/]+	187.601	0.159	0.496	2.582	2626
3.5%	374	0	0	GET	/api/posting-window/current	184.840	0.494	7.049	7.118	238
3.4%	166	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	180.730	1.089	12.812	14.803	115
1.5%	367	0	0	GET	/api/memories	80.297	0.219	1.181	4.046	12229
1.2%	56	6	0	POST	^/api/posts/[^/]+/realmojis$	66.273	1.183	7.104	7.104	410
1.2%	44	0	0	PUT	/api/me/profile-image	65.138	1.480	7.613	7.613	0
1.2%	89	0	0	POST	/api/auth/register	64.818	0.728	8.120	8.120	228
1.1%	64	0	0	PUT	^/api/follows/[^/]+$	60.695	0.948	7.106	7.106	0
0.9%	23	0	0	GET	/api/me/export	48.471	2.107	13.046	13.046	2448920
0.8%	259	0	0	GET	/api/users	44.450	0.172	0.409	5.192	243
0.8%	50	0	0	POST	/api/auth/logout	44.207	0.884	12.019	12.019	0
0.8%	46	0	0	PATCH	/api/me	42.159	0.917	6.294	6.294	260
0.6%	190	5	0	GET	^/api/posts/me/current$	29.573	0.156	0.320	0.673	622
0.5%	189	9	0	GET	^/api/posts/[^/]+$	28.484	0.151	0.387	0.667	583
0.5%	192	0	0	GET	^/api/relationships/[^/]+	26.665	0.139	0.407	0.475	943
0.4%	38	0	0	POST	^/api/posts/[^/]+/comments$	21.660	0.570	5.076	5.076	372
0.3%	101	0	0	GET	^/api/users/[^/]+$	16.750	0.166	0.467	1.130	198
0.3%	96	1	0	GET	^/api/posts/[^/]+/realmojis$	14.086	0.147	0.337	0.337	722
0.3%	21	0	0	PATCH	^/api/posts/[^/]+$	14.033	0.668	4.994	4.994	645
0.2%	72	0	0	GET	^/api/posts/[^/]+/comments$	11.546	0.160	0.910	0.910	1030
0.1%	49	1	0	GET	/api/me	7.457	0.152	0.274	0.274	292
0.1%	41	0	0	GET	^/api/memories/[^/]+	7.319	0.179	0.354	0.354	782
0.1%	6	0	0	DELETE	^/api/posts/[^/]+$	6.627	1.105	3.361	3.361	0
0.1%	10	0	0	DELETE	^/api/realmojis/[^/]+$	5.846	0.585	2.238	2.238	0
0.0%	7	0	0	DELETE	^/api/comments/[^/]+$	1.916	0.274	0.620	0.620	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.7%	156	1	0	GET	/api/notifications	1165.652	7.472	9.260	9.260	175
21.7%	179	3	0	POST	/api/posts	729.938	4.078	13.649	14.179	617
8.9%	345	9	0	POST	/api/friend-requests	298.529	0.865	7.079	14.558	113
8.3%	227	0	0	POST	/api/auth/login	279.755	1.232	12.550	14.460	249
4.5%	897	6	0	GET	^/api/images/[^/]+	152.690	0.170	1.060	1.368	165940
3.8%	120	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	126.297	1.052	13.274	14.690	115
3.0%	578	0	0	GET	^/api/feeds/[^/]+	100.727	0.174	0.491	11.699	2551
2.5%	241	0	0	GET	/api/posting-window/current	82.768	0.343	4.407	7.058	238
1.7%	274	0	0	GET	/api/memories	58.307	0.213	1.116	1.613	14176
1.6%	52	0	0	POST	/api/auth/register	54.382	1.046	6.776	6.776	229
1.1%	26	0	0	PUT	/api/me/profile-image	36.443	1.402	5.581	5.581	0
1.0%	42	0	0	POST	^/api/posts/[^/]+/realmojis$	33.714	0.803	2.211	2.211	465
0.8%	12	0	0	GET	/api/me/export	27.184	2.265	7.482	7.482	2784863
0.8%	173	0	0	GET	/api/users	26.821	0.155	0.436	0.667	238
0.7%	44	0	0	POST	^/api/posts/[^/]+/comments$	24.765	0.563	4.552	4.552	305
0.7%	134	5	0	GET	^/api/posts/me/current$	23.495	0.175	0.908	1.099	601
0.7%	29	0	0	PUT	^/api/follows/[^/]+$	23.103	0.797	5.368	5.368	0
0.6%	20	0	0	PATCH	/api/me	19.079	0.954	7.856	7.856	279
0.5%	118	0	0	GET	^/api/relationships/[^/]+	17.455	0.148	0.557	0.756	924
0.4%	29	0	0	POST	/api/auth/logout	14.537	0.501	5.260	5.260	0
0.4%	17	2	0	PATCH	^/api/posts/[^/]+$	13.547	0.797	3.957	3.957	575
0.4%	88	2	0	GET	^/api/posts/[^/]+$	12.538	0.142	0.372	0.372	609
0.2%	7	0	0	DELETE	^/api/realmojis/[^/]+$	7.339	1.048	2.964	2.964	0
0.2%	6	0	0	DELETE	^/api/posts/[^/]+$	5.774	0.962	1.874	1.874	0
0.2%	31	0	0	GET	^/api/memories/[^/]+	5.413	0.175	0.470	0.470	946
0.1%	30	1	0	GET	/api/me	4.998	0.167	0.333	0.333	286
0.1%	30	0	0	GET	^/api/users/[^/]+$	4.832	0.161	0.430	0.430	206
0.1%	29	0	0	GET	^/api/posts/[^/]+/realmojis$	4.701	0.162	0.346	0.346	2127
0.1%	20	0	0	GET	^/api/posts/[^/]+/comments$	2.659	0.133	0.251	0.251	1334
0.0%	9	0	0	DELETE	^/api/comments/[^/]+$	1.400	0.156	0.363	0.363	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 45010 クエリ / 84 種類 / 合計 265.674s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	59.7%	2640	158.591	60.07	432.51	607.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	23.4%	196	62.196	317.33	2080.90	6696.7	130.4	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.9%	1827	15.789	8.64	451.49	118.1	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	5.0%	6191	13.180	2.13	56.48	10.6	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.5%	1718	4.071	2.37	111.03	26.1	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.9%	1	2.295	2294.82	2294.82	29829.0	1199538.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.8%	600	2.018	3.36	13.41	28.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
8	0.4%	1	1.158	1157.75	1157.75	32748.0	1.0	select pg_size_pretty(avg(length(data))) from (select data from image_objects order by created_at desc limit $1) t
9	0.4%	518	0.943	1.82	22.92	20.7	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
10	0.3%	2	0.756	378.01	748.11	16796.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.3%	1	0.731	730.71	730.71	11539.0	529848.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
12	0.3%	3139	0.683	0.22	18.98	20.3	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
13	0.2%	2	0.427	213.33	277.81	216191.0	2543.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.2%	3208	0.399	0.12	13.47	11.3	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
15	0.1%	1340	0.344	0.26	7.64	21.7	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 93cebbb64943b8a20fb2be2ee828625a8aafe067
Type: cpu
Time: 2026-10-31 12:20:44 JST
Duration: 60.05s, Total samples = 52.54s (87.49%)
Showing nodes accounting for 40.21s, 76.53% of 52.54s total
Dropped 1022 nodes (cum <= 0.26s)
Showing top 40 nodes out of 275
      flat  flat%   sum%        cum   cum%
     3.36s  6.40%  6.40%      3.36s  6.40%  internal/runtime/syscall/linux.Syscall6
     3.36s  6.40% 12.79%      3.36s  6.40%  runtime.memclrNoHeapPointers
     3.08s  5.86% 18.65%     10.71s 20.38%  runtime.scanObject
     2.82s  5.37% 24.02%      9.12s 17.36%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     2.59s  4.93% 28.95%      3.39s  6.45%  runtime.findObject
     2.37s  4.51% 33.46%      2.47s  4.70%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     2.36s  4.49% 37.95%      2.36s  4.49%  runtime.memmove
     2.30s  4.38% 42.33%      3.82s  7.27%  runtime.tryDeferToSpanScan
     2.28s  4.34% 46.67%      2.28s  4.34%  crypto/internal/fips140/sha256.blockAVX2
     2.27s  4.32% 50.99%      3.82s  7.27%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.54s  2.93% 53.92%      5.68s 10.81%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.48s  2.82% 56.74%      1.48s  2.82%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.23s  2.34% 59.08%      1.23s  2.34%  aeshashbody
     1.13s  2.15% 61.23%      1.62s  3.08%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     1.09s  2.07% 63.30%      1.09s  2.07%  runtime.spanClass.sizeclass (inline)
     0.90s  1.71% 65.02%      0.90s  1.71%  example.com/isureal-company/isureal.(*UserRow).active (partial-inline)
     0.59s  1.12% 66.14%      0.59s  1.12%  memeqbody
     0.58s  1.10% 67.24%      2.14s  4.07%  runtime.mapaccess1_faststr
     0.47s  0.89% 68.14%      0.49s  0.93%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.39s  0.74% 68.88%      0.60s  1.14%  example.com/isureal-company/isureal.postLess
     0.38s  0.72% 69.60%      0.38s  0.72%  runtime.madvise
     0.35s  0.67% 70.27%      0.41s  0.78%  runtime.spanOf (inline)
     0.30s  0.57% 70.84%      0.81s  1.54%  runtime.scanObjectsSmall
     0.29s  0.55% 71.39%      0.33s  0.63%  runtime.typePointers.next
     0.29s  0.55% 71.95%      0.29s  0.55%  runtime.typePointers.nextFast (inline)
     0.27s  0.51% 72.46%      0.27s  0.51%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.23s  0.44% 72.90%      0.63s  1.20%  runtime.greyobject
     0.21s   0.4% 73.30%      0.32s  0.61%  runtime.heapArenaOf (inline)
     0.20s  0.38% 73.68%      1.90s  3.62%  runtime.mapassign_faststr
     0.17s  0.32% 74.00%      2.54s  4.83%  crypto/internal/fips140/sha256.(*Digest).Write
     0.17s  0.32% 74.32%      0.46s  0.88%  database/sql.convertAssignRows
     0.17s  0.32% 74.65%      1.02s  1.94%  runtime.mallocgcSmallScanNoHeader
     0.15s  0.29% 74.93%      2.61s  4.97%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.15s  0.29% 75.22%      0.30s  0.57%  runtime.scanObjectSmall
     0.13s  0.25% 75.47%      0.31s  0.59%  internal/runtime/maps.(*Iter).Next
     0.13s  0.25% 75.71%      1.28s  2.44%  runtime.scanSpan
     0.12s  0.23% 75.94%      0.52s  0.99%  encoding/json.structEncoder.encode
     0.11s  0.21% 76.15%      1.51s  2.87%  example.com/isureal-company/isureal.NewServer.getFeedHandler.func50
     0.10s  0.19% 76.34%      1.63s  3.10%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.10s  0.19% 76.53%      7.44s 14.16%  runtime.mallocgc
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 52.2% / 100% | 38% | 0.1% | 8 | 176 MB | CPU が飽和 |
| isu2 (2 cpu) | 10.4% / 30% | 0% | 0.5% | 3 | 2762 MB |  |
| isu3 (2 cpu) | 7.3% / 35% | 0% | 0.5% | 3 | 2834 MB |  |
| isu4 (2 cpu) | 44.8% / 100% | 18% | 13.5% | 26 | 108 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 85% / nginx 13% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 19% |
| isu3 | isureal 13% / nginx 1% |
| isu4 | postgres 49% / isureal 9% |

## Diff from previous run

スコア: 29539 → 65904 (+36365)

`55a9d63..fdbe826`

```
 webapp/go/account_deletion_api.go |   2 +-
 webapp/go/admin_user_api.go       |   2 +-
 webapp/go/comment_api.go          |   2 +-
 webapp/go/realmoji_api.go         |   2 +-
 webapp/go/relationship.go         |  10 +++-
 webapp/go/server.go               |   1 -
 webapp/go/slow.go                 |  21 --------
 webapp/go/store.go                | 100 ++++++++++++++++++++++++++++++++------
 8 files changed, 96 insertions(+), 44 deletions(-)
```
