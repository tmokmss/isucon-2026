# run 20261031-125104-d3f5685

## Summary

| | |
|---|---|
| score | 87494 |
| status | pass |
| ref | grandprix-tcp-idle |
| sha | d3f568540604 |
| 所要時間 | 239s |
| 前回比 | -13166 (-13.1%) vs 20261031-124715-46aeaca |
| ベスト比 | -13166 (-13.1%) vs 20261031-124715-46aeaca = 100660 |
| note | PR73 + tcp_slow_start_after_idle=0 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	41	111	308	774	774
europe	7	16	42	106	129	128
america	6	16	50	100	103	102
```

```
{"id":"10100","status":"done","portalStatus":"finished","passed":true,"score":87494,"raw":87494,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:53:36.722Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:53:39.687Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:53:42.682Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:53:44.252Z level=INFO msg="user moved" user=anka5100 from=europe to=america
time=2026-10-31T03:53:45.709Z level=INFO msg="asia 2 日目: 41 人"
time=2026-10-31T03:53:48.717Z level=INFO msg="europe 2 日目: 16 人"
time=2026-10-31T03:53:51.684Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T03:53:54.705Z level=INFO msg="asia 3 日目: 111 人"
time=2026-10-31T03:53:55.137Z level=INFO msg="user moved" user=noyoho_5450 from=asia to=america
time=2026-10-31T03:53:56.438Z level=INFO msg="user moved" user=noyoho_5450 from=america to=asia
time=2026-10-31T03:53:56.599Z level=INFO msg="user moved" user=chiroshi5638 from=asia to=america
time=2026-10-31T03:53:56.678Z level=INFO msg="user moved" user=anka5100 from=america to=europe
time=2026-10-31T03:53:57.485Z level=INFO msg="user moved" user=solevanu-5751 from=asia to=america
time=2026-10-31T03:53:57.687Z level=INFO msg="europe 3 日目: 42 人"
time=2026-10-31T03:53:58.199Z level=INFO msg="user moved" user=funezoe.5863 from=asia to=america
time=2026-10-31T03:53:58.317Z level=INFO msg="user moved" user=chiroshi5638 from=america to=asia
time=2026-10-31T03:54:00.683Z level=INFO msg="america 3 日目: 50 人"
time=2026-10-31T03:54:02.478Z level=INFO msg="user moved" user=anka5100 from=europe to=america
time=2026-10-31T03:54:03.727Z level=INFO msg="asia 4 日目: 308 人"
time=2026-10-31T03:54:05.395Z level=INFO msg="user moved" user=keyusa6346 from=asia to=europe
time=2026-10-31T03:54:06.481Z level=INFO msg="user moved" user=noyoho_5450 from=asia to=america
time=2026-10-31T03:54:06.728Z level=INFO msg="europe 4 日目: 106 人"
time=2026-10-31T03:54:06.817Z level=INFO msg="user moved" user=chiroshi5638 from=asia to=america
time=2026-10-31T03:54:07.909Z level=INFO msg="user moved" user=kune.6520 from=asia to=america
time=2026-10-31T03:54:08.767Z level=INFO msg="user moved" user=suorben6294 from=america to=asia
time=2026-10-31T03:54:08.885Z level=INFO msg="user moved" user=torhais6249 from=america to=asia
time=2026-10-31T03:54:09.048Z level=INFO msg="user moved" user=yatsu6898 from=asia to=america
time=2026-10-31T03:54:09.065Z level=INFO msg="user moved" user=uleva-7292 from=asia to=america
time=2026-10-31T03:54:09.700Z level=INFO msg="america 4 日目: 100 人"
time=2026-10-31T03:54:09.808Z level=INFO msg="user moved" user=ware_7155 from=asia to=europe
time=2026-10-31T03:54:09.816Z level=INFO msg="user moved" user=isku6791 from=asia to=america
time=2026-10-31T03:54:09.957Z level=INFO msg="user moved" user=suorben6294 from=asia to=america
time=2026-10-31T03:54:10.025Z level=INFO msg="user moved" user=vintsuku.7413 from=asia to=europe
time=2026-10-31T03:54:10.030Z level=INFO msg="user moved" user=torhais6249 from=asia to=europe
time=2026-10-31T03:54:10.075Z level=INFO msg="user moved" user=evayoni_7444 from=asia to=europe
time=2026-10-31T03:54:10.176Z level=INFO msg="user moved" user=maxul7489 from=asia to=america
time=2026-10-31T03:54:11.092Z level=INFO msg="user moved" user=maryu6613 from=asia to=europe
time=2026-10-31T03:54:11.766Z level=INFO msg="user moved" user=ulmu7717 from=asia to=europe
time=2026-10-31T03:54:12.787Z level=INFO msg="asia 5 日目: 774 人"
time=2026-10-31T03:54:16.826Z level=INFO msg="europe 5 日目: 129 人"
time=2026-10-31T03:54:18.703Z level=INFO msg="america 5 日目: 103 人"
time=2026-10-31T03:54:27.920Z level=INFO msg="america 最終日（グランプリ）: 102 人"
time=2026-10-31T03:54:28.043Z level=INFO msg="europe 最終日（グランプリ）: 128 人"
time=2026-10-31T03:54:28.192Z level=INFO msg="asia 最終日（グランプリ）: 774 人"
time=2026-10-31T03:54:36.711Z level=INFO msg="グランプリ: アジア 206 人, ヨーロッパ 17 人, アメリカ 10 人"
{
  "pass": true,
  "score": 87494,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
29.6%	1134	0	0	GET	/api/notifications	9187.003	8.101	15.508	15.524	163
17.6%	1627	50	0	POST	/api/posts	5463.048	3.358	10.998	12.129	633
13.0%	3840	64	0	POST	/api/friend-requests	4043.084	1.053	9.460	14.724	114
9.0%	11088	567	0	GET	^/api/images/[^/]+	2790.425	0.252	5.840	12.858	121986
8.5%	4906	1343	0	POST	/api/auth/login	2633.809	0.537	7.229	12.818	184
4.4%	805	1	0	POST	/api/auth/register	1362.221	1.692	8.645	11.441	227
4.3%	1357	6	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1337.093	0.985	9.619	12.304	116
2.8%	577	123	0	PUT	/api/me/profile-image	869.082	1.506	8.220	17.655	4
1.5%	504	0	0	PUT	^/api/follows/[^/]+$	471.995	0.936	6.263	16.230	0
1.4%	396	21	0	POST	^/api/posts/[^/]+/realmojis$	419.653	1.060	8.133	9.137	445
1.3%	372	0	0	PATCH	/api/me	402.969	1.083	7.729	9.293	279
1.1%	2758	0	0	GET	/api/memories	344.123	0.125	0.930	1.799	11697
1.0%	8524	0	0	GET	^/api/feeds/[^/]+	298.204	0.035	0.630	1.932	2107
0.8%	141	4	0	PATCH	^/api/posts/[^/]+$	243.961	1.730	9.484	9.634	630
0.7%	219	0	0	GET	/api/me/export	212.393	0.970	10.603	10.910	1812251
0.6%	344	10	0	POST	^/api/posts/[^/]+/comments$	197.763	0.575	4.926	8.973	319
0.6%	562	0	0	POST	/api/auth/logout	197.002	0.351	6.486	8.810	0
0.3%	47	0	0	DELETE	^/api/posts/[^/]+$	95.557	2.033	8.571	8.571	0
0.2%	2404	0	0	GET	/api/posting-window/current	74.461	0.031	0.410	1.808	237
0.2%	1209	64	0	GET	^/api/posts/me/current$	65.802	0.054	0.726	1.841	619
0.2%	1917	0	0	GET	/api/users	49.797	0.026	0.320	0.992	248
0.1%	1194	0	0	GET	^/api/relationships/[^/]+	39.312	0.033	0.652	1.089	920
0.1%	67	0	0	DELETE	^/api/realmojis/[^/]+$	36.273	0.541	3.098	3.098	0
0.1%	58	0	0	DELETE	^/api/comments/[^/]+$	31.247	0.539	12.476	12.476	0
0.1%	530	0	0	GET	^/api/memories/[^/]+	31.154	0.059	0.440	0.953	639
0.1%	1236	7	0	GET	^/api/posts/[^/]+$	27.575	0.022	0.295	1.836	649
0.0%	24	0	0	PUT	^/api/blocks/[^/]+$	14.076	0.587	3.684	3.684	0
0.0%	1	0	0	POST	/api/initialize	12.495	12.495	12.495	12.495	90
0.0%	1	0	0	POST	/api/me/account-deletion	12.253	12.253	12.253	12.253	0
0.0%	642	0	0	GET	^/api/users/[^/]+$	12.143	0.019	0.161	0.709	223
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.5%	301	1	0	GET	/api/notifications	2570.688	8.540	11.316	12.640	174
18.6%	367	12	0	POST	/api/posts	1241.944	3.384	10.823	12.166	633
7.8%	652	10	0	POST	/api/friend-requests	520.218	0.798	7.676	13.097	114
5.7%	459	0	0	POST	/api/auth/login	379.758	0.827	9.261	10.681	249
5.0%	1666	9	0	GET	^/api/images/[^/]+	332.800	0.200	1.090	1.920	190239
4.6%	1563	0	0	GET	^/api/feeds/[^/]+	306.621	0.196	0.641	1.084	2208
3.2%	159	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	214.180	1.347	8.539	9.224	116
1.8%	504	0	0	GET	/api/posting-window/current	118.873	0.236	0.733	1.538	239
1.7%	446	0	0	GET	/api/memories	114.300	0.256	1.349	1.718	12906
1.4%	76	7	0	POST	^/api/posts/[^/]+/realmojis$	91.782	1.208	10.944	10.944	461
1.3%	86	2	0	POST	^/api/posts/[^/]+/comments$	88.267	1.026	6.205	6.205	316
1.3%	64	0	0	POST	/api/auth/logout	87.576	1.368	9.814	9.814	0
1.1%	121	0	0	POST	/api/auth/register	71.040	0.587	4.626	5.664	228
0.9%	274	19	0	GET	^/api/posts/me/current$	59.188	0.216	0.993	1.568	612
0.9%	300	0	0	GET	/api/users	58.818	0.196	0.589	0.970	250
0.8%	32	3	0	PATCH	^/api/posts/[^/]+$	55.993	1.750	7.606	7.606	585
0.6%	22	0	0	GET	/api/me/export	43.385	1.972	5.490	5.490	2195650
0.6%	86	0	0	PUT	^/api/follows/[^/]+$	41.817	0.486	7.551	7.551	0
0.6%	74	0	0	PUT	/api/me/profile-image	41.413	0.560	4.195	4.195	0
0.5%	12	0	0	DELETE	^/api/posts/[^/]+$	36.001	3.000	9.071	9.071	0
0.5%	197	14	0	GET	^/api/posts/[^/]+$	34.491	0.175	0.558	0.569	591
0.5%	173	0	0	GET	^/api/relationships/[^/]+	33.635	0.194	0.906	1.003	773
0.4%	66	0	0	PATCH	/api/me	27.007	0.409	4.989	4.989	286
0.3%	114	0	0	GET	^/api/users/[^/]+$	21.230	0.186	0.571	0.606	206
0.3%	100	0	0	GET	^/api/posts/[^/]+/realmojis$	17.318	0.173	0.561	0.593	1119
0.2%	70	0	0	GET	^/api/memories/[^/]+	15.428	0.220	1.062	1.062	985
0.2%	81	2	0	GET	/api/me	13.439	0.166	0.526	0.526	293
0.2%	83	0	0	GET	^/api/posts/[^/]+/comments$	13.349	0.161	0.617	0.617	1277
0.1%	10	0	0	DELETE	^/api/realmojis/[^/]+$	9.731	0.973	5.547	5.547	0
0.1%	11	0	0	DELETE	^/api/comments/[^/]+$	5.104	0.464	1.841	1.841	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.7%	268	0	0	GET	/api/notifications	2032.733	7.585	9.282	9.283	176
19.8%	311	8	0	POST	/api/posts	1069.878	3.440	10.357	12.424	632
9.8%	518	4	0	POST	/api/friend-requests	528.473	1.020	8.137	11.495	114
5.5%	1275	15	0	GET	^/api/images/[^/]+	297.594	0.233	2.556	3.014	182270
5.1%	364	0	0	POST	/api/auth/login	277.073	0.761	8.100	10.116	250
3.7%	202	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	201.267	0.996	6.873	9.631	116
3.2%	921	0	0	GET	^/api/feeds/[^/]+	169.958	0.185	0.746	1.390	2241
1.9%	437	0	0	GET	/api/posting-window/current	101.747	0.233	0.730	1.266	240
1.8%	358	0	0	GET	/api/memories	95.716	0.267	1.179	2.259	14124
1.1%	56	0	0	POST	/api/auth/logout	61.237	1.094	6.896	6.896	0
1.0%	23	0	0	GET	/api/me/export	55.064	2.394	9.217	9.217	3186486
1.0%	233	15	0	GET	^/api/posts/me/current$	54.081	0.232	0.940	1.667	607
0.9%	250	0	0	GET	/api/users	48.901	0.196	0.715	0.781	247
0.8%	53	0	0	PUT	/api/me/profile-image	44.574	0.841	7.229	7.229	0
0.8%	46	3	0	POST	^/api/posts/[^/]+/realmojis$	41.327	0.898	4.049	4.049	452
0.8%	51	0	0	PUT	^/api/follows/[^/]+$	40.793	0.800	7.021	7.021	0
0.7%	16	0	0	DELETE	^/api/posts/[^/]+$	36.360	2.272	8.772	8.772	0
0.5%	23	2	0	PATCH	^/api/posts/[^/]+$	27.615	1.201	5.610	5.610	594
0.5%	154	0	0	GET	^/api/relationships/[^/]+	26.409	0.171	0.707	0.979	1021
0.5%	50	0	0	PATCH	/api/me	26.384	0.528	5.226	5.226	273
0.5%	93	0	0	POST	/api/auth/register	24.653	0.265	1.423	1.423	228
0.4%	10	0	0	DELETE	^/api/realmojis/[^/]+$	20.629	2.063	14.476	14.476	0
0.4%	130	6	0	GET	^/api/posts/[^/]+$	20.578	0.158	0.526	0.627	601
0.3%	9	1	0	POST	^/api/posts/[^/]+/comments$	18.515	2.057	6.519	6.519	342
0.3%	58	0	0	GET	^/api/memories/[^/]+	15.953	0.275	2.408	2.408	936
0.2%	75	3	0	GET	^/api/posts/[^/]+/realmojis$	11.179	0.149	0.534	0.534	1220
0.2%	2	0	0	DELETE	^/api/comments/[^/]+$	10.308	5.154	10.148	10.148	0
0.2%	57	1	0	GET	/api/me	9.835	0.173	0.880	0.880	299
0.2%	57	0	0	GET	^/api/users/[^/]+$	8.517	0.149	0.538	0.538	219
0.1%	1	0	0	DELETE	^/api/friendships/[^/]+$	5.101	5.101	5.101	5.101	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 58425 クエリ / 82 種類 / 合計 417.711s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	62.8%	3306	262.212	79.31	578.73	656.1	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	22.6%	264	94.232	356.94	1676.47	6863.8	125.2	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.5%	2541	23.003	9.05	562.31	123.4	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.4%	7604	18.421	2.42	44.63	10.3	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.7%	2237	7.289	3.26	67.80	25.3	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.5%	1	2.211	2210.58	2210.58	30229.0	1199800.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.4%	1	1.868	1867.92	1867.92	22094.0	626923.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.4%	487	1.862	3.82	48.00	29.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.2%	4925	0.950	0.19	21.16	19.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	0.2%	426	0.762	1.79	8.76	21.1	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.2%	1	0.748	747.75	747.75	11910.0	529988.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
12	0.1%	4386	0.531	0.12	14.24	11.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
13	0.1%	1654	0.501	0.30	27.03	20.6	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
14	0.1%	2	0.412	205.86	255.36	227117.0	2672.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
15	0.1%	2	0.262	131.04	254.03	16785.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: b4eac213b40b3648cf29a70b6a372487cec6463b
Type: cpu
Time: 2026-10-31 12:53:17 JST
Duration: 60.09s, Total samples = 45.83s (76.27%)
Showing nodes accounting for 33.92s, 74.01% of 45.83s total
Dropped 1081 nodes (cum <= 0.23s)
Showing top 40 nodes out of 288
      flat  flat%   sum%        cum   cum%
     4.19s  9.14%  9.14%      4.19s  9.14%  crypto/internal/fips140/sha256.blockAVX2
     3.42s  7.46% 16.60%      3.42s  7.46%  internal/runtime/syscall/linux.Syscall6
     2.21s  4.82% 21.43%      7.65s 16.69%  runtime.scanObject
     2.02s  4.41% 25.83%      2.02s  4.41%  runtime.memmove
     1.91s  4.17% 30.00%      2.50s  5.45%  runtime.findObject
     1.91s  4.17% 34.17%      1.91s  4.17%  runtime.memclrNoHeapPointers
     1.71s  3.73% 37.90%      5.70s 12.44%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.55s  3.38% 41.28%      1.55s  3.38%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.45s  3.16% 44.45%      2.53s  5.52%  runtime.tryDeferToSpanScan
     1.30s  2.84% 47.28%      2.12s  4.63%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.30s  2.84% 50.12%      4.40s  9.60%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.16s  2.53% 52.65%      1.16s  2.53%  aeshashbody
     0.93s  2.03% 54.68%      0.98s  2.14%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.85s  1.85% 56.54%      0.85s  1.85%  example.com/isureal-company/isureal.(*UserRow).active
     0.85s  1.85% 58.39%      0.85s  1.85%  runtime.spanClass.sizeclass (inline)
     0.68s  1.48% 59.87%      0.87s  1.90%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.56s  1.22% 61.10%      0.56s  1.22%  memeqbody
     0.51s  1.11% 62.21%      1.99s  4.34%  runtime.mapaccess1_faststr
     0.45s  0.98% 63.19%      0.45s  0.98%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.43s  0.94% 64.13%      0.43s  0.94%  indexbody
     0.39s  0.85% 64.98%      0.55s  1.20%  example.com/isureal-company/isureal.postLess
     0.37s  0.81% 65.79%      0.37s  0.81%  runtime.nextFreeFast (inline)
     0.35s  0.76% 66.55%      0.68s  1.48%  runtime.greyobject
     0.34s  0.74% 67.29%      0.34s  0.74%  runtime.futex
     0.32s   0.7% 67.99%      0.62s  1.35%  database/sql.convertAssignRows
     0.26s  0.57% 68.56%      0.26s  0.57%  runtime.typePointers.nextFast (inline)
     0.24s  0.52% 69.08%      0.24s  0.52%  runtime.(*mspan).divideByElemSize (inline)
     0.23s   0.5% 69.58%      4.61s 10.06%  crypto/internal/fips140/sha256.(*Digest).Write
     0.22s  0.48% 70.06%      1.87s  4.08%  runtime.mapassign_faststr
     0.21s  0.46% 70.52%      1.05s  2.29%  runtime.mallocgcSmallScanNoHeader
     0.20s  0.44% 70.96%      0.27s  0.59%  runtime.spanOf (inline)
     0.19s  0.41% 71.37%      4.89s 10.67%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.19s  0.41% 71.79%      0.23s   0.5%  runtime.(*mspan).markBitsForIndex (inline)
     0.17s  0.37% 72.16%      4.23s  9.23%  runtime.mallocgc
     0.16s  0.35% 72.51%      0.34s  0.74%  compress/flate.(*compressor).deflate
     0.16s  0.35% 72.86%      4.62s 10.08%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.16s  0.35% 73.21%      0.54s  1.18%  runtime.scanObjectsSmall
     0.13s  0.28% 73.49%      0.45s  0.98%  runtime.mapaccess2_faststr
     0.13s  0.28% 73.77%      0.82s  1.79%  runtime.slicebytetostring
     0.11s  0.24% 74.01%      5.13s 11.19%  crypto/internal/fips140/hmac.(*HMAC).Sum
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 48.2% / 100% | 32% | 0.1% | 8 | 181 MB | CPU が飽和 |
| isu2 (2 cpu) | 14.5% / 56% | 0% | 0.5% | 5 | 2527 MB | run queue が積もっている |
| isu3 (2 cpu) | 14.5% / 60% | 0% | 0.5% | 3 | 2558 MB |  |
| isu4 (2 cpu) | 49.0% / 100% | 24% | 16.1% | 22 | 109 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 76% / nginx 13% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 27% / nginx 1% |
| isu3 | isureal 26% / nginx 1% |
| isu4 | postgres 59% / isureal 3% |

## Diff from previous run

スコア: 100660 → 87494 (-13166)

`46aeaca..d3f5685`

```
 conf/_all/etc/sysctl.d/99-isucon.conf | 2 ++
 1 file changed, 2 insertions(+)
```
