# run 20261031-122401-fdbe826

## Summary

| | |
|---|---|
| score | 66717 |
| status | pass |
| ref | origin/main |
| sha | fdbe82670a24 |
| 所要時間 | 106s |
| 前回比 | +66717 (—) vs 20261031-122244-5443d83 |
| ベスト比 | +813 (1.2%) vs 20261031-122039-fdbe826 = 65904 **ベスト更新** |
| note | 低い版の試走の直後の main |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	45	122	345	489	488
europe	7	20	48	115	116	115
america	6	16	39	73	72	72
```

```
{"id":"7922","status":"done","portalStatus":"finished","passed":true,"score":66717,"raw":66717,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:24:28.235Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:24:31.190Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:24:34.190Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:24:37.215Z level=INFO msg="asia 2 日目: 45 人"
time=2026-10-31T03:24:38.790Z level=INFO msg="user moved" user=wariru5214 from=asia to=america
time=2026-10-31T03:24:39.318Z level=INFO msg="user moved" user=benya5232 from=asia to=europe
time=2026-10-31T03:24:39.462Z level=INFO msg="user moved" user=yaya.5229 from=asia to=europe
time=2026-10-31T03:24:40.189Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T03:24:40.323Z level=INFO msg="user moved" user=wariru5214 from=america to=europe
time=2026-10-31T03:24:43.190Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T03:24:46.218Z level=INFO msg="asia 3 日目: 122 人"
time=2026-10-31T03:24:48.694Z level=INFO msg="user moved" user=dannana5544 from=asia to=europe
time=2026-10-31T03:24:49.200Z level=INFO msg="europe 3 日目: 48 人"
time=2026-10-31T03:24:49.350Z level=INFO msg="user moved" user=fusu_5591 from=asia to=europe
time=2026-10-31T03:24:50.113Z level=INFO msg="user moved" user=isanshi6058 from=asia to=europe
time=2026-10-31T03:24:52.014Z level=INFO msg="user moved" user=benya5232 from=europe to=asia
time=2026-10-31T03:24:52.192Z level=INFO msg="america 3 日目: 39 人"
time=2026-10-31T03:24:53.616Z level=INFO msg="user moved" user=wariru5214 from=europe to=america
time=2026-10-31T03:24:55.239Z level=INFO msg="asia 4 日目: 345 人"
time=2026-10-31T03:24:57.586Z level=INFO msg="user moved" user=benya5232 from=asia to=europe
time=2026-10-31T03:24:58.226Z level=INFO msg="europe 4 日目: 115 人"
time=2026-10-31T03:25:01.333Z level=INFO msg="america 4 日目: 73 人"
time=2026-10-31T03:25:04.501Z level=INFO msg="asia 5 日目: 489 人"
time=2026-10-31T03:25:07.399Z level=INFO msg="user moved" user=wariru5214 from=america to=europe
time=2026-10-31T03:25:08.359Z level=INFO msg="europe 5 日目: 116 人"
time=2026-10-31T03:25:10.311Z level=INFO msg="america 5 日目: 72 人"
time=2026-10-31T03:25:19.368Z level=INFO msg="europe 最終日（グランプリ）: 115 人"
time=2026-10-31T03:25:19.550Z level=INFO msg="asia 最終日（グランプリ）: 488 人"
time=2026-10-31T03:25:19.657Z level=INFO msg="america 最終日（グランプリ）: 72 人"
time=2026-10-31T03:25:28.202Z level=INFO msg="グランプリ: アジア 20 人, ヨーロッパ 0 人, アメリカ 3 人"
{
  "pass": true,
  "score": 66717,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
31.8%	973	2	0	GET	/api/notifications	8389.310	8.622	15.115	15.330	172
22.1%	1329	74	0	POST	/api/posts	5831.435	4.388	13.907	14.788	615
11.6%	8830	648	0	GET	^/api/images/[^/]+	3062.163	0.347	6.833	14.367	117133
8.2%	2648	44	0	POST	/api/friend-requests	2154.969	0.814	8.284	13.807	114
7.5%	3661	1030	0	POST	/api/auth/login	1974.722	0.539	7.368	15.685	183
3.8%	1046	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	990.153	0.947	11.796	15.164	116
2.8%	511	1	0	POST	/api/auth/register	741.710	1.451	7.509	11.364	227
1.8%	401	129	0	PUT	/api/me/profile-image	477.672	1.191	10.046	13.693	6
1.8%	448	19	0	POST	^/api/posts/[^/]+/realmojis$	469.790	1.049	5.570	8.519	456
1.5%	1848	0	0	GET	/api/posting-window/current	405.611	0.219	3.398	12.869	236
0.9%	2192	0	0	GET	/api/memories	245.893	0.112	0.750	5.760	11092
0.9%	233	10	0	POST	^/api/posts/[^/]+/comments$	240.620	1.033	5.763	10.339	340
0.7%	329	0	0	PUT	^/api/follows/[^/]+$	190.992	0.581	3.195	6.235	0
0.7%	460	0	0	POST	/api/auth/logout	183.844	0.400	10.170	12.609	0
0.7%	6201	0	0	GET	^/api/feeds/[^/]+	178.821	0.029	0.163	4.783	2044
0.6%	240	0	0	PATCH	/api/me	167.923	0.700	3.500	4.939	272
0.4%	124	5	0	PATCH	^/api/posts/[^/]+$	110.510	0.891	5.406	10.907	618
0.4%	153	0	0	GET	/api/me/export	108.942	0.712	2.366	13.052	1972259
0.2%	43	0	0	DELETE	^/api/posts/[^/]+$	56.024	1.303	11.183	11.183	0
0.2%	1281	0	0	GET	/api/users	49.137	0.038	0.288	3.396	247
0.2%	64	0	0	DELETE	^/api/realmojis/[^/]+$	47.124	0.736	6.127	6.127	0
0.2%	29	0	0	DELETE	^/api/comments/[^/]+$	43.556	1.502	13.114	13.114	0
0.2%	922	52	0	GET	^/api/posts/me/current$	41.548	0.045	0.386	3.345	614
0.1%	940	0	0	GET	^/api/relationships/[^/]+	28.953	0.031	0.208	3.434	1039
0.1%	15	0	0	PUT	^/api/admin/users/[^/]+/certification$	23.828	1.589	13.907	13.907	0
0.1%	417	0	0	GET	^/api/memories/[^/]+	20.975	0.050	0.292	0.612	638
0.1%	854	7	0	GET	^/api/posts/[^/]+$	20.844	0.024	0.144	0.645	635
0.1%	1	0	0	POST	/api/initialize	16.518	16.518	16.518	16.518	90
0.0%	467	0	0	GET	^/api/users/[^/]+$	11.388	0.024	0.148	0.397	217
0.0%	453	0	0	GET	^/api/posts/[^/]+/realmojis$	9.902	0.022	0.135	0.235	1572
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.1%	278	0	0	GET	/api/notifications	2266.381	8.152	11.091	12.016	176
24.1%	353	7	0	POST	/api/posts	1604.119	4.544	14.222	14.863	645
8.4%	566	8	0	POST	/api/friend-requests	559.327	0.988	9.693	13.180	114
6.7%	411	0	0	POST	/api/auth/login	448.253	1.091	11.668	14.435	247
4.2%	209	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	278.000	1.330	14.373	14.541	115
3.5%	1449	11	0	GET	^/api/images/[^/]+	229.572	0.158	0.987	1.462	185467
3.3%	1314	0	0	GET	^/api/feeds/[^/]+	220.025	0.167	0.714	1.112	2420
2.0%	484	0	0	GET	/api/posting-window/current	136.077	0.281	2.936	5.799	239
1.4%	440	0	0	GET	/api/memories	91.816	0.209	0.572	0.696	11541
1.3%	92	0	0	PUT	^/api/follows/[^/]+$	84.910	0.923	11.740	11.740	0
1.2%	71	0	0	POST	/api/auth/logout	81.460	1.147	13.588	13.588	0
1.2%	30	0	0	GET	/api/me/export	78.675	2.623	13.301	13.301	2347317
1.1%	71	0	0	PUT	/api/me/profile-image	71.283	1.004	5.433	5.433	0
1.0%	57	5	0	POST	^/api/posts/[^/]+/realmojis$	64.459	1.131	4.787	4.787	461
1.0%	108	0	0	POST	/api/auth/register	63.251	0.586	2.539	2.713	228
0.8%	48	0	0	PATCH	/api/me	52.669	1.097	5.755	5.755	297
0.7%	305	0	0	GET	/api/users	49.424	0.162	0.459	1.593	245
0.7%	54	3	0	POST	^/api/posts/[^/]+/comments$	45.139	0.836	9.237	9.237	320
0.6%	253	26	0	GET	^/api/posts/me/current$	37.504	0.148	0.497	0.548	592
0.5%	211	0	0	GET	^/api/relationships/[^/]+	32.215	0.153	0.555	0.683	937
0.4%	31	2	0	PATCH	^/api/posts/[^/]+$	29.065	0.938	5.418	5.418	595
0.4%	159	7	0	GET	^/api/posts/[^/]+$	25.933	0.163	0.709	0.734	615
0.3%	112	0	0	GET	^/api/users/[^/]+$	19.359	0.173	0.474	0.835	219
0.2%	11	0	0	DELETE	^/api/posts/[^/]+$	13.982	1.271	3.643	3.643	0
0.2%	77	2	0	GET	/api/me	13.518	0.176	0.868	0.868	292
0.2%	6	0	0	PUT	^/api/blocks/[^/]+$	11.243	1.874	7.912	7.912	0
0.2%	50	0	0	GET	^/api/memories/[^/]+	10.917	0.218	0.835	0.835	1167
0.1%	59	0	0	GET	^/api/posts/[^/]+/realmojis$	9.844	0.167	0.548	0.548	1780
0.1%	52	0	0	GET	^/api/posts/[^/]+/comments$	8.677	0.167	0.734	0.734	1261
0.1%	6	0	0	DELETE	^/api/comments/[^/]+$	8.118	1.353	2.726	2.726	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.5%	177	0	0	GET	/api/notifications	1331.941	7.525	9.350	9.350	174
20.7%	217	5	0	POST	/api/posts	847.742	3.907	13.609	13.907	645
11.2%	445	8	0	POST	/api/friend-requests	456.960	1.027	12.166	14.451	113
8.7%	276	0	0	POST	/api/auth/login	356.713	1.292	12.522	13.054	249
4.2%	990	11	0	GET	^/api/images/[^/]+	172.557	0.174	1.026	1.519	175447
3.4%	99	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	139.512	1.409	13.611	13.611	116
2.9%	749	0	0	GET	^/api/feeds/[^/]+	119.420	0.159	0.594	1.982	2400
1.6%	294	0	0	GET	/api/posting-window/current	66.628	0.227	2.925	3.077	240
1.6%	294	0	0	GET	/api/memories	65.961	0.224	1.108	1.640	14601
1.6%	42	0	0	PUT	/api/me/profile-image	64.508	1.536	6.967	6.967	0
1.5%	54	2	0	POST	^/api/posts/[^/]+/realmojis$	62.425	1.156	8.798	8.798	475
1.5%	71	0	0	POST	/api/auth/register	62.086	0.874	8.594	8.594	229
1.1%	210	0	0	GET	/api/users	43.214	0.206	1.539	5.473	258
0.8%	23	2	0	PATCH	^/api/posts/[^/]+$	32.798	1.426	5.716	5.716	583
0.8%	45	0	0	PUT	^/api/follows/[^/]+$	32.577	0.724	4.085	4.085	0
0.8%	41	0	0	POST	/api/auth/logout	31.337	0.764	11.781	11.781	0
0.8%	33	0	0	PATCH	/api/me	30.745	0.932	5.544	5.544	289
0.7%	16	0	0	GET	/api/me/export	30.406	1.900	7.274	7.274	2420405
0.7%	166	12	0	GET	^/api/posts/me/current$	29.605	0.178	0.687	0.720	612
0.6%	12	0	0	DELETE	^/api/posts/[^/]+$	26.379	2.198	12.146	12.146	0
0.6%	113	0	0	GET	^/api/relationships/[^/]+	23.730	0.210	3.007	3.070	930
0.3%	92	7	0	GET	^/api/posts/[^/]+$	13.769	0.150	0.458	0.458	614
0.3%	19	1	0	POST	^/api/posts/[^/]+/comments$	12.602	0.663	3.019	3.019	306
0.2%	60	0	0	GET	^/api/users/[^/]+$	9.470	0.158	0.458	0.458	227
0.2%	48	0	0	GET	^/api/posts/[^/]+/realmojis$	7.234	0.151	0.939	0.939	1695
0.2%	33	0	0	GET	^/api/memories/[^/]+	6.815	0.207	1.026	1.026	967
0.2%	46	1	0	GET	/api/me	6.391	0.139	0.515	0.515	296
0.2%	35	0	0	GET	^/api/posts/[^/]+/comments$	6.225	0.178	0.816	0.816	1277
0.1%	3	0	0	DELETE	^/api/comments/[^/]+$	2.381	0.794	2.155	2.155	0
0.0%	5	0	0	DELETE	^/api/realmojis/[^/]+$	1.597	0.319	0.635	0.635	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 45942 クエリ / 81 種類 / 合計 294.026s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	60.8%	2735	178.641	65.32	540.48	616.6	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	21.9%	199	64.304	323.14	1128.22	6973.5	131.7	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.6%	2059	16.347	7.94	306.23	118.5	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.3%	5988	12.615	2.11	111.31	10.5	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	2.1%	1816	6.282	3.46	87.70	26.1	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.8%	2	5.172	2585.92	5014.53	224825.0	2645.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.7%	1	2.168	2167.93	2167.93	29864.0	1199948.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.6%	1	1.637	1637.12	1637.12	21119.0	626828.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.5%	533	1.600	3.00	15.58	29.4	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.3%	1	0.750	750.24	750.24	11571.0	530068.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
11	0.2%	3583	0.670	0.19	24.79	19.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
12	0.2%	292	0.573	1.96	11.42	23.6	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
13	0.1%	3318	0.425	0.13	13.63	11.1	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
14	0.1%	1326	0.318	0.24	13.61	21.5	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	18	0.272	15.08	156.29	268.4	1.0	INSERT INTO posting_windows (id, region, local_date, starts_at, on_time_deadline, created_at, notification_ids, notified_user_ids) VALUES ($1, $2, $3, $4, $5, $…

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
Time: 2026-10-31 12:24:05 JST
Duration: 60.34s, Total samples = 42.69s (70.75%)
Showing nodes accounting for 31.94s, 74.82% of 42.69s total
Dropped 986 nodes (cum <= 0.21s)
Showing top 40 nodes out of 273
      flat  flat%   sum%        cum   cum%
     2.84s  6.65%  6.65%      2.84s  6.65%  internal/runtime/syscall/linux.Syscall6
     2.57s  6.02% 12.67%      2.57s  6.02%  runtime.memclrNoHeapPointers
     2.22s  5.20% 17.87%      7.78s 18.22%  runtime.scanObject
     2.09s  4.90% 22.77%      2.09s  4.90%  crypto/internal/fips140/sha256.blockAVX2
     2.01s  4.71% 27.48%      2.66s  6.23%  runtime.findObject
     1.90s  4.45% 31.93%      5.96s 13.96%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.84s  4.31% 36.24%      2.89s  6.77%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.82s  4.26% 40.50%      5.70s 13.35%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.69s  3.96% 44.46%      1.69s  3.96%  runtime.memmove
     1.57s  3.68% 48.14%      2.47s  5.79%  runtime.tryDeferToSpanScan
     1.28s  3.00% 51.14%      1.28s  3.00%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.20s  2.81% 53.95%      1.20s  2.81%  aeshashbody
     1.11s  2.60% 56.55%      1.17s  2.74%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.85s  1.99% 58.54%      0.85s  1.99%  example.com/isureal-company/isureal.(*UserRow).active
     0.78s  1.83% 60.37%      1.10s  2.58%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.74s  1.73% 62.10%      0.74s  1.73%  memeqbody
     0.66s  1.55% 63.64%      0.66s  1.55%  runtime.spanClass.sizeclass (inline)
     0.50s  1.17% 64.82%      1.95s  4.57%  runtime.mapaccess1_faststr
     0.35s  0.82% 65.64%      0.64s  1.50%  example.com/isureal-company/isureal.postLess
     0.32s  0.75% 66.39%      0.32s  0.75%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.29s  0.68% 67.06%      0.29s  0.68%  runtime.nextFreeFast (inline)
     0.26s  0.61% 67.67%      0.26s  0.61%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.26s  0.61% 68.28%      0.31s  0.73%  runtime.spanOf (inline)
     0.25s  0.59% 68.87%      0.25s  0.59%  indexbody
     0.25s  0.59% 69.45%      0.53s  1.24%  runtime.scanObjectsSmall
     0.23s  0.54% 69.99%      0.23s  0.54%  runtime.(*mspan).divideByElemSize (inline)
     0.20s  0.47% 70.46%      0.48s  1.12%  runtime.greyobject
     0.20s  0.47% 70.93%      0.24s  0.56%  runtime.typePointers.next
     0.18s  0.42% 71.35%      0.85s  1.99%  runtime.mallocgcSmallScanNoHeader
     0.17s   0.4% 71.75%      5.14s 12.04%  runtime.mallocgc
     0.17s   0.4% 72.15%      1.61s  3.77%  runtime.mapassign_faststr
     0.15s  0.35% 72.50%      2.35s  5.50%  crypto/internal/fips140/sha256.(*Digest).Write
     0.14s  0.33% 72.83%      0.28s  0.66%  compress/flate.(*compressor).deflate
     0.14s  0.33% 73.16%      2.43s  5.69%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.13s   0.3% 73.46%      0.36s  0.84%  database/sql.convertAssignRows
     0.13s   0.3% 73.76%      0.25s  0.59%  internal/runtime/maps.(*Iter).Next
     0.12s  0.28% 74.05%      1.63s  3.82%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.12s  0.28% 74.33%      2.54s  5.95%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s  0.26% 74.58%      0.24s  0.56%  runtime.scanObjectSmall
     0.10s  0.23% 74.82%      2.75s  6.44%  crypto/internal/fips140/hmac.(*HMAC).Sum
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 54.2% / 100% | 38% | 0.1% | 5 | 123 MB | CPU が飽和 |
| isu2 (2 cpu) | 11.1% / 32% | 0% | 0.5% | 4 | 2694 MB |  |
| isu3 (2 cpu) | 8.5% / 33% | 0% | 0.5% | 4 | 2780 MB |  |
| isu4 (2 cpu) | 47.0% / 100% | 19% | 16.5% | 35 | 103 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 89% / nginx 13% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 20% / nginx 1% |
| isu3 | isureal 15% / nginx 1% |
| isu4 | postgres 48% / isureal 10% |

## Diff from previous run

スコア: 0 → 66717 (+66717)

`5443d83..fdbe826`

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
