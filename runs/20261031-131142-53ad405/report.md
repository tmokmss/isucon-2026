# run 20261031-131142-53ad405

## Summary

| | |
|---|---|
| score | 108397 |
| status | pass |
| ref | HEAD |
| sha | 53ad405390e9 |
| 所要時間 | 248s |
| 前回比 | -2057 (-1.9%) vs 20261031-130542-c75a0a9 |
| ベスト比 | -2057 (-1.9%) vs 20261031-130542-c75a0a9 = 110454 |
| note | tls-off-export-pk-sync-off-bodybuf |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	44	122	342	754	753
europe	7	15	39	98	109	108
america	7	17	42	82	83	82
```

```
{"id":"11638","status":"done","portalStatus":"finished","passed":true,"score":108397,"raw":108397,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:14:19.942Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:14:21.286Z level=INFO msg="user moved" user=oryodan5026 from=asia to=america
time=2026-10-31T04:14:22.908Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:14:25.909Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T04:14:28.940Z level=INFO msg="asia 2 日目: 44 人"
time=2026-10-31T04:14:31.912Z level=INFO msg="europe 2 日目: 15 人"
time=2026-10-31T04:14:31.974Z level=INFO msg="user moved" user=shiivyse5258 from=asia to=america
time=2026-10-31T04:14:34.911Z level=INFO msg="america 2 日目: 17 人"
time=2026-10-31T04:14:37.931Z level=INFO msg="asia 3 日目: 122 人"
time=2026-10-31T04:14:38.244Z level=INFO msg="user moved" user=oryodan5026 from=america to=europe
time=2026-10-31T04:14:40.372Z level=INFO msg="user moved" user=orleora_5696 from=asia to=america
time=2026-10-31T04:14:40.909Z level=INFO msg="europe 3 日目: 39 人"
time=2026-10-31T04:14:41.434Z level=INFO msg="user moved" user=isnoa5762 from=asia to=europe
time=2026-10-31T04:14:43.892Z level=INFO msg="user moved" user=orleora_5696 from=america to=asia
time=2026-10-31T04:14:43.957Z level=INFO msg="america 3 日目: 42 人"
time=2026-10-31T04:14:44.706Z level=INFO msg="user moved" user=oryodan5026 from=europe to=america
time=2026-10-31T04:14:46.963Z level=INFO msg="asia 4 日目: 342 人"
time=2026-10-31T04:14:48.660Z level=INFO msg="user moved" user=nanetsu-6453 from=asia to=europe
time=2026-10-31T04:14:49.577Z level=INFO msg="user moved" user=nurato-6531 from=asia to=europe
time=2026-10-31T04:14:49.999Z level=INFO msg="europe 4 日目: 98 人"
time=2026-10-31T04:14:50.192Z level=INFO msg="user moved" user=koto.6435 from=asia to=america
time=2026-10-31T04:14:50.212Z level=INFO msg="user moved" user=orleora_5696 from=asia to=america
time=2026-10-31T04:14:50.594Z level=INFO msg="user moved" user=shihi.6448 from=asia to=america
time=2026-10-31T04:14:50.974Z level=INFO msg="user moved" user=keulni6287 from=america to=europe
time=2026-10-31T04:14:53.045Z level=INFO msg="america 4 日目: 82 人"
time=2026-10-31T04:14:53.418Z level=INFO msg="user moved" user=niis-7470 from=asia to=europe
time=2026-10-31T04:14:54.400Z level=INFO msg="user moved" user=keulni6287 from=europe to=america
time=2026-10-31T04:14:56.084Z level=INFO msg="asia 5 日目: 754 人"
time=2026-10-31T04:14:59.146Z level=INFO msg="europe 5 日目: 109 人"
time=2026-10-31T04:15:01.925Z level=INFO msg="america 5 日目: 83 人"
time=2026-10-31T04:15:11.059Z level=INFO msg="europe 最終日（グランプリ）: 108 人"
time=2026-10-31T04:15:11.159Z level=INFO msg="america 最終日（グランプリ）: 82 人"
time=2026-10-31T04:15:11.295Z level=INFO msg="asia 最終日（グランプリ）: 753 人"
time=2026-10-31T04:15:19.939Z level=INFO msg="グランプリ: アジア 525 人, ヨーロッパ 39 人, アメリカ 24 人"
{
  "pass": true,
  "score": 108397,
  "messages": [
    "ivymose-5033: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.8%	1285	1	0	GET	/api/notifications	10810.964	8.413	15.333	15.340	171
14.9%	1936	26	0	POST	/api/posts	4393.656	2.269	7.739	9.316	647
11.7%	4470	79	0	POST	/api/friend-requests	3428.103	0.767	5.034	10.410	113
9.1%	13024	459	0	GET	^/api/images/[^/]+	2674.539	0.205	3.338	8.726	147379
5.5%	4682	864	0	POST	/api/auth/login	1620.354	0.346	4.183	10.041	204
4.7%	1652	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1390.099	0.841	5.001	10.018	116
2.5%	584	122	0	PUT	/api/me/profile-image	725.331	1.242	6.453	7.835	4
2.2%	654	17	0	POST	^/api/posts/[^/]+/comments$	639.281	0.977	6.224	7.283	335
1.9%	437	16	0	POST	^/api/posts/[^/]+/realmojis$	558.786	1.279	6.390	11.228	466
1.9%	779	1	0	POST	/api/auth/register	551.878	0.708	4.296	7.385	227
1.4%	3313	0	0	GET	/api/memories	413.404	0.125	0.666	1.135	10393
1.2%	506	0	0	PUT	^/api/follows/[^/]+$	366.045	0.723	4.258	7.323	0
1.1%	393	0	0	PATCH	/api/me	311.591	0.793	4.850	9.782	280
0.9%	12114	0	0	GET	^/api/feeds/[^/]+	271.737	0.022	0.172	2.830	2358
0.7%	192	3	0	PATCH	^/api/posts/[^/]+$	209.497	1.091	4.967	5.857	642
0.6%	632	0	0	POST	/api/auth/logout	188.683	0.299	3.363	7.691	0
0.6%	239	0	0	GET	/api/me/export	187.367	0.784	4.529	8.799	1964692
0.4%	69	0	0	DELETE	^/api/posts/[^/]+$	110.423	1.600	8.131	8.131	0
0.2%	2749	0	0	GET	/api/posting-window/current	70.565	0.026	0.170	0.857	237
0.2%	109	0	0	DELETE	^/api/comments/[^/]+$	63.107	0.579	4.360	5.201	0
0.2%	1739	40	0	GET	^/api/posts/me/current$	54.720	0.031	0.238	0.646	641
0.2%	2181	0	0	GET	/api/users	51.395	0.024	0.130	1.747	243
0.1%	61	0	0	DELETE	^/api/realmojis/[^/]+$	41.422	0.679	3.702	3.702	0
0.1%	34	0	0	PUT	^/api/blocks/[^/]+$	39.265	1.155	5.423	5.423	0
0.1%	1988	18	0	GET	^/api/posts/[^/]+$	39.117	0.020	0.116	2.088	615
0.1%	606	0	0	GET	^/api/memories/[^/]+	30.754	0.051	0.477	0.773	617
0.1%	1515	0	0	GET	^/api/relationships/[^/]+	30.473	0.020	0.213	0.397	885
0.1%	16	0	0	PUT	^/api/admin/users/[^/]+/certification$	22.542	1.409	10.560	10.560	0
0.1%	1	0	0	POST	/api/initialize	20.075	20.075	20.075	20.075	90
0.1%	932	0	0	GET	^/api/users/[^/]+$	16.815	0.018	0.128	0.384	196
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.1%	279	1	0	GET	/api/notifications	2315.483	8.299	11.646	11.680	175
17.7%	348	9	0	POST	/api/posts	1023.103	2.940	8.854	9.846	635
6.7%	1541	13	0	GET	^/api/images/[^/]+	386.260	0.251	1.252	1.985	182561
6.3%	634	13	0	POST	/api/friend-requests	360.747	0.569	4.133	5.307	113
5.1%	425	0	0	POST	/api/auth/login	292.540	0.688	4.756	8.108	248
4.2%	1361	0	0	GET	^/api/feeds/[^/]+	245.283	0.180	0.516	0.948	2576
2.7%	191	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	155.890	0.816	3.384	6.057	116
2.2%	436	0	0	GET	/api/memories	129.170	0.296	1.125	1.484	13494
2.1%	115	11	0	POST	^/api/posts/[^/]+/realmojis$	119.435	1.039	6.144	7.268	436
1.9%	472	0	0	GET	/api/posting-window/current	109.475	0.232	0.629	1.017	239
1.1%	78	0	0	POST	/api/auth/logout	62.310	0.799	5.096	5.096	0
1.1%	283	8	0	GET	^/api/posts/me/current$	60.661	0.214	0.598	0.609	634
1.0%	312	0	0	GET	/api/users	56.270	0.180	0.596	0.759	240
0.9%	58	1	0	POST	^/api/posts/[^/]+/comments$	51.714	0.892	4.525	4.525	351
0.9%	103	0	0	POST	/api/auth/register	51.390	0.499	2.780	2.838	228
0.9%	26	0	0	GET	/api/me/export	50.068	1.926	6.576	6.576	2394890
0.7%	236	10	0	GET	^/api/posts/[^/]+$	38.768	0.164	0.459	0.564	595
0.6%	58	0	0	PUT	/api/me/profile-image	36.012	0.621	3.954	3.954	0
0.6%	200	0	0	GET	^/api/relationships/[^/]+	33.226	0.166	0.508	1.085	885
0.5%	31	3	0	PATCH	^/api/posts/[^/]+$	29.242	0.943	4.790	4.790	596
0.4%	85	0	0	PUT	^/api/follows/[^/]+$	25.106	0.295	3.689	3.689	0
0.4%	46	0	0	PATCH	/api/me	24.487	0.532	4.036	4.036	274
0.3%	72	0	0	GET	^/api/memories/[^/]+	17.171	0.238	0.994	0.994	598
0.3%	13	0	0	DELETE	^/api/posts/[^/]+$	16.334	1.256	3.842	3.842	0
0.3%	106	0	0	GET	^/api/users/[^/]+$	15.955	0.151	0.431	0.466	193
0.3%	98	0	0	GET	^/api/posts/[^/]+/realmojis$	15.074	0.154	0.552	0.552	1027
0.2%	13	0	0	DELETE	^/api/comments/[^/]+$	12.034	0.926	3.312	3.312	0
0.2%	64	0	0	GET	^/api/posts/[^/]+/comments$	11.020	0.172	0.428	0.428	1402
0.2%	15	0	0	DELETE	^/api/realmojis/[^/]+$	10.379	0.692	2.040	2.040	0
0.2%	63	2	0	GET	/api/me	9.555	0.152	0.361	0.361	292
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.4%	231	1	0	GET	/api/notifications	1783.303	7.720	9.260	9.314	177
18.7%	280	7	0	POST	/api/posts	826.663	2.952	8.655	9.815	642
7.5%	1208	9	0	GET	^/api/images/[^/]+	332.040	0.275	1.690	2.408	181904
6.0%	437	7	0	POST	/api/friend-requests	265.006	0.606	4.624	5.049	114
4.7%	328	0	0	POST	/api/auth/login	208.969	0.637	5.167	8.146	252
4.1%	186	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	180.443	0.970	6.349	7.518	116
4.0%	865	0	0	GET	^/api/feeds/[^/]+	174.691	0.202	0.578	0.784	2780
2.4%	365	0	0	GET	/api/memories	105.422	0.289	1.144	1.687	12579
1.7%	374	0	0	GET	/api/posting-window/current	76.046	0.203	0.573	0.624	239
1.4%	24	0	0	GET	/api/me/export	59.993	2.500	7.380	7.380	3426805
1.1%	228	14	0	GET	^/api/posts/me/current$	50.236	0.220	0.598	0.899	618
0.9%	214	0	0	GET	/api/users	38.655	0.181	0.591	0.595	241
0.9%	41	4	0	POST	^/api/posts/[^/]+/realmojis$	37.924	0.925	5.857	5.857	427
0.8%	41	2	0	POST	^/api/posts/[^/]+/comments$	35.656	0.870	4.697	4.697	342
0.7%	52	0	0	POST	/api/auth/logout	30.436	0.585	3.610	3.610	0
0.7%	44	0	0	PUT	/api/me/profile-image	29.479	0.670	5.106	5.106	0
0.6%	147	0	0	GET	^/api/relationships/[^/]+	26.321	0.179	0.542	0.558	1170
0.4%	20	0	0	PATCH	^/api/posts/[^/]+$	19.287	0.964	4.714	4.714	653
0.4%	116	6	0	GET	^/api/posts/[^/]+$	18.931	0.163	0.422	0.493	590
0.4%	43	0	0	PUT	^/api/follows/[^/]+$	17.397	0.405	5.303	5.303	0
0.4%	77	0	0	POST	/api/auth/register	15.645	0.203	0.811	0.811	229
0.3%	35	0	0	PATCH	/api/me	13.563	0.388	2.136	2.136	277
0.3%	10	0	0	DELETE	^/api/posts/[^/]+$	13.431	1.343	8.126	8.126	0
0.3%	50	0	0	GET	^/api/memories/[^/]+	13.265	0.265	1.102	1.102	961
0.2%	8	0	0	DELETE	^/api/realmojis/[^/]+$	9.947	1.243	5.568	5.568	0
0.2%	48	0	0	GET	^/api/users/[^/]+$	8.318	0.173	0.376	0.376	189
0.2%	44	0	0	GET	^/api/posts/[^/]+/realmojis$	8.029	0.182	0.535	0.535	667
0.2%	48	1	0	GET	/api/me	7.269	0.151	0.502	0.502	297
0.2%	36	0	0	GET	^/api/posts/[^/]+/comments$	7.119	0.198	0.548	0.548	1540
0.0%	3	0	0	PUT	^/api/blocks/[^/]+$	1.962	0.654	1.726	1.726	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 63629 クエリ / 82 種類 / 合計 486.345s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	63.6%	3650	309.088	84.68	831.91	660.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	19.8%	289	96.240	333.01	1907.59	1336.8	135.9	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	6.1%	2422	29.701	12.26	768.47	124.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.0%	8514	19.646	2.31	70.14	10.7	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.7%	2522	8.419	3.34	119.59	25.3	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.6%	2	7.703	3851.63	7552.63	236718.5	2784.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.5%	1	2.424	2423.78	2423.78	23025.0	627646.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.4%	1	2.168	2168.04	2168.04	30564.0	1199802.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
9	0.4%	562	2.122	3.78	56.44	28.6	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.4%	5428	2.009	0.37	126.78	19.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.3%	733	1.283	1.75	15.12	20.3	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.2%	1969	0.969	0.49	117.18	20.3	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
13	0.2%	1	0.753	752.64	752.64	12235.0	530078.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
14	0.1%	2026	0.590	0.29	116.28	7.9	1.0	UPDATE friend_requests SET status=$2, updated_at=$3 WHERE id=$1
15	0.1%	4571	0.555	0.12	23.22	11.1	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 23a8803dc801ea10893e98ca3a04e9f2e98e2646
Type: cpu
Time: 2026-10-31 13:13:54 JST
Duration: 60s, Total samples = 35.85s (59.75%)
Showing nodes accounting for 26.08s, 72.75% of 35.85s total
Dropped 929 nodes (cum <= 0.18s)
Showing top 40 nodes out of 297
      flat  flat%   sum%        cum   cum%
     3.19s  8.90%  8.90%      3.19s  8.90%  internal/runtime/syscall/linux.Syscall6
     1.83s  5.10% 14.00%      6.18s 17.24%  runtime.scanObject
     1.77s  4.94% 18.94%      1.77s  4.94%  runtime.memmove
     1.75s  4.88% 23.82%      5.76s 16.07%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.75s  4.88% 28.70%      1.75s  4.88%  runtime.memclrNoHeapPointers
     1.71s  4.77% 33.47%      1.71s  4.77%  crypto/internal/fips140/sha256.blockAVX2
     1.70s  4.74% 38.21%      2.11s  5.89%  runtime.findObject
     1.47s  4.10% 42.32%      1.47s  4.10%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.28s  3.57% 45.89%      1.28s  3.57%  aeshashbody
     1.03s  2.87% 48.76%      1.93s  5.38%  runtime.tryDeferToSpanScan
     0.71s  1.98% 50.74%      0.71s  1.98%  runtime.spanClass.sizeclass (inline)
     0.70s  1.95% 52.69%      0.70s  1.95%  example.com/isureal-company/isureal.(*UserRow).active
     0.58s  1.62% 54.31%      0.58s  1.62%  memeqbody
     0.52s  1.45% 55.76%      2.15s  6.00%  runtime.mapaccess1_faststr
     0.43s  1.20% 56.96%      1.27s  3.54%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.41s  1.14% 58.10%      0.41s  1.14%  runtime.nextFreeFast (inline)
     0.39s  1.09% 59.19%      0.56s  1.56%  example.com/isureal-company/isureal.postLess
     0.35s  0.98% 60.17%      0.72s  2.01%  database/sql.convertAssignRows
     0.34s  0.95% 61.12%      0.34s  0.95%  runtime.futex
     0.30s  0.84% 61.95%      0.54s  1.51%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.29s  0.81% 62.76%      0.30s  0.84%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.26s  0.73% 63.49%      1.24s  3.46%  runtime.mallocgcSmallScanNoHeader
     0.25s   0.7% 64.18%      0.25s   0.7%  indexbody
     0.25s   0.7% 64.88%      0.57s  1.59%  runtime.greyobject
     0.25s   0.7% 65.58%      0.25s   0.7%  runtime.typePointers.nextFast (inline)
     0.21s  0.59% 66.16%      4.62s 12.89%  runtime.mallocgc
     0.19s  0.53% 66.69%      1.44s  4.02%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.19s  0.53% 67.22%      2.77s  7.73%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.19s  0.53% 67.75%      0.31s  0.86%  internal/runtime/maps.(*Iter).Next
     0.19s  0.53% 68.28%      0.24s  0.67%  runtime.(*mspan).markBitsForIndex (inline)
     0.19s  0.53% 68.81%      0.22s  0.61%  runtime.(*mspan).writeHeapBitsSmall
     0.19s  0.53% 69.34%      1.80s  5.02%  runtime.mapassign_faststr
     0.18s   0.5% 69.85%      0.18s   0.5%  runtime.(*mspan).divideByElemSize (inline)
     0.17s  0.47% 70.32%      0.25s   0.7%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.17s  0.47% 70.79%      0.86s  2.40%  runtime.scanSpan
     0.16s  0.45% 71.24%      0.45s  1.26%  runtime.mapaccess2_faststr
     0.15s  0.42% 71.66%      0.18s   0.5%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.15s  0.42% 72.08%      4.01s 11.19%  example.com/isureal-company/isureal.loadBig
     0.13s  0.36% 72.44%      0.18s   0.5%  runtime.heapArenaOf (inline)
     0.11s  0.31% 72.75%      1.22s  3.40%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 43.6% / 100% | 30% | 0.2% | 7 | 177 MB | run queue が積もっている |
| isu2 (2 cpu) | 16.7% / 64% | 0% | 0.6% | 5 | 1980 MB | run queue が積もっている |
| isu3 (2 cpu) | 14.8% / 64% | 0% | 0.4% | 5 | 2446 MB | run queue が積もっている |
| isu4 (2 cpu) | 50.6% / 100% | 22% | 18.5% | 30 | 108 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 67% / nginx 11% / journalctl 2% / systemd-journal 2% |
| isu2 | isureal 31% |
| isu3 | isureal 27% / nginx 1% |
| isu4 | postgres 54% / isureal 5% / kcompactd0 2% |

## Diff from previous run

スコア: 110454 → 108397 (-2057)

`c75a0a9..53ad405`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf         | 1 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf         | 1 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf         | 1 +
 conf/isu4/etc/postgresql/16/main/conf.d/zz-tuning.conf | 2 ++
 4 files changed, 5 insertions(+)
```
