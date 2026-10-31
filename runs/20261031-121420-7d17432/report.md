# run 20261031-121420-7d17432

## Summary

| | |
|---|---|
| score | 67674 |
| status | pass |
| ref | relation-pair-locks |
| sha | 7d17432731d3 |
| 所要時間 | 221s |
| 前回比 | +67674 (—) vs 20261031-120912-249be25 |
| ベスト比 | +7473 (12.4%) vs 20261031-115040-1c73ee4 = 60201 **ベスト更新** |
| note | relation pair locks + max_connections 300 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	53	155	449	456	456
europe	7	20	49	102	102	102
america	6	14	38	49	49	49
```

```
{"id":"7366","status":"done","portalStatus":"finished","passed":true,"score":67674,"raw":67674,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:16:35.225Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:16:38.191Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:16:41.190Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:16:44.230Z level=INFO msg="asia 2 日目: 53 人"
time=2026-10-31T03:16:45.093Z level=INFO msg="user moved" user=kero.5164 from=asia to=europe
time=2026-10-31T03:16:47.191Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T03:16:47.728Z level=INFO msg="user moved" user=isre-5372 from=asia to=europe
time=2026-10-31T03:16:50.192Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T03:16:53.218Z level=INFO msg="asia 3 日目: 155 人"
time=2026-10-31T03:16:56.228Z level=INFO msg="europe 3 日目: 49 人"
time=2026-10-31T03:16:56.395Z level=INFO msg="user moved" user=heben5815 from=asia to=america
time=2026-10-31T03:16:58.800Z level=INFO msg="user moved" user=isre-5372 from=europe to=america
time=2026-10-31T03:16:59.270Z level=INFO msg="america 3 日目: 38 人"
time=2026-10-31T03:17:02.252Z level=INFO msg="asia 4 日目: 449 人"
time=2026-10-31T03:17:05.307Z level=INFO msg="europe 4 日目: 102 人"
time=2026-10-31T03:17:08.395Z level=INFO msg="america 4 日目: 49 人"
time=2026-10-31T03:17:11.405Z level=INFO msg="asia 5 日目: 456 人"
time=2026-10-31T03:17:14.867Z level=INFO msg="europe 5 日目: 102 人"
time=2026-10-31T03:17:17.244Z level=INFO msg="america 5 日目: 49 人"
time=2026-10-31T03:17:26.330Z level=INFO msg="asia 最終日（グランプリ）: 456 人"
time=2026-10-31T03:17:26.602Z level=INFO msg="america 最終日（グランプリ）: 49 人"
time=2026-10-31T03:17:26.767Z level=INFO msg="europe 最終日（グランプリ）: 102 人"
time=2026-10-31T03:17:35.200Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 67674,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.4%	1032	0	0	GET	/api/notifications	9542.844	9.247	15.065	15.326	175
22.4%	1383	45	0	POST	/api/posts	6206.881	4.488	14.272	15.261	629
10.1%	8739	659	0	GET	^/api/images/[^/]+	2802.053	0.321	5.456	15.258	142241
8.5%	2697	48	0	POST	/api/friend-requests	2365.879	0.877	6.993	13.993	113
7.7%	3933	1408	0	POST	/api/auth/login	2145.101	0.545	10.519	14.846	167
4.1%	1129	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1123.825	0.995	12.851	15.260	116
1.7%	1890	0	0	GET	/api/posting-window/current	475.883	0.252	3.645	12.437	237
1.4%	374	122	0	PUT	/api/me/profile-image	392.175	1.049	8.936	9.714	8
1.2%	475	1	0	POST	/api/auth/register	326.215	0.687	4.546	12.965	227
1.2%	381	12	0	POST	^/api/posts/[^/]+/realmojis$	324.126	0.851	5.411	6.243	462
1.1%	451	17	0	POST	^/api/posts/[^/]+/comments$	310.973	0.690	5.556	7.905	334
0.9%	7764	0	0	GET	^/api/feeds/[^/]+	250.400	0.032	0.203	3.608	1976
0.8%	294	0	0	PUT	^/api/follows/[^/]+$	227.301	0.773	6.396	7.205	0
0.8%	2170	0	0	GET	/api/memories	226.655	0.104	0.595	2.161	10019
0.6%	172	0	0	GET	/api/me/export	161.483	0.939	11.503	12.060	1950027
0.6%	459	0	0	POST	/api/auth/logout	159.260	0.347	5.479	11.975	0
0.5%	223	0	0	PATCH	/api/me	142.540	0.639	4.144	4.717	276
0.3%	102	8	0	PATCH	^/api/posts/[^/]+$	75.837	0.744	4.333	4.917	600
0.2%	69	0	0	DELETE	^/api/realmojis/[^/]+$	67.507	0.978	12.617	12.617	0
0.2%	52	0	0	DELETE	^/api/posts/[^/]+$	63.423	1.220	5.793	5.793	0
0.2%	83	0	0	DELETE	^/api/comments/[^/]+$	52.110	0.628	15.839	15.839	0
0.2%	1293	0	0	GET	/api/users	43.382	0.034	0.218	0.577	236
0.1%	1267	21	0	GET	^/api/posts/[^/]+$	36.859	0.029	0.180	4.286	618
0.1%	988	54	0	GET	^/api/posts/me/current$	30.759	0.031	0.234	0.392	615
0.1%	417	0	0	GET	^/api/memories/[^/]+	26.457	0.063	0.340	1.952	624
0.1%	1072	0	0	GET	^/api/relationships/[^/]+	24.870	0.023	0.185	2.219	987
0.1%	13	0	0	PUT	^/api/admin/users/[^/]+/certification$	19.888	1.530	13.513	13.513	0
0.1%	607	0	0	GET	^/api/users/[^/]+$	19.179	0.032	0.198	1.879	207
0.1%	1	0	0	POST	/api/initialize	13.898	13.898	13.898	13.898	90
0.0%	517	0	0	GET	^/api/posts/[^/]+/realmojis$	13.600	0.026	0.217	0.275	1460
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.4%	242	0	0	GET	/api/notifications	2013.145	8.319	11.977	11.978	179
20.9%	311	6	0	POST	/api/posts	1260.676	4.054	13.790	14.226	640
11.1%	639	19	0	POST	/api/friend-requests	669.333	1.047	12.127	14.499	112
6.7%	372	0	0	POST	/api/auth/login	400.503	1.077	13.563	14.672	247
4.1%	206	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	245.531	1.192	12.656	13.897	116
3.7%	1330	18	0	GET	^/api/images/[^/]+	220.914	0.166	1.006	1.385	185204
2.9%	1076	0	0	GET	^/api/feeds/[^/]+	172.112	0.160	0.386	1.288	2317
2.8%	409	0	0	GET	/api/posting-window/current	170.043	0.416	4.388	5.354	239
1.5%	421	0	0	GET	/api/memories	89.816	0.213	0.654	1.287	9244
1.4%	59	0	0	PUT	/api/me/profile-image	84.265	1.428	7.709	7.709	0
1.3%	99	6	0	POST	^/api/posts/[^/]+/comments$	77.334	0.781	5.838	5.838	317
1.3%	62	0	0	PUT	^/api/follows/[^/]+$	76.950	1.241	7.155	7.155	0
1.1%	47	0	0	PATCH	/api/me	65.091	1.385	10.461	10.461	289
1.0%	99	0	0	POST	/api/auth/register	60.475	0.611	3.223	3.223	228
1.0%	25	0	0	GET	/api/me/export	58.585	2.343	13.957	13.957	2788903
1.0%	57	0	0	POST	/api/auth/logout	57.470	1.008	12.495	12.495	0
0.8%	293	0	0	GET	/api/users	49.210	0.168	0.485	3.829	231
0.6%	230	3	0	GET	^/api/posts/me/current$	35.858	0.156	0.410	0.462	644
0.5%	29	2	0	POST	^/api/posts/[^/]+/realmojis$	27.427	0.946	4.440	4.440	464
0.5%	194	0	0	GET	^/api/relationships/[^/]+	27.427	0.141	0.486	0.612	1165
0.4%	161	19	0	GET	^/api/posts/[^/]+$	26.179	0.163	0.372	2.513	552
0.4%	4	0	0	DELETE	^/api/realmojis/[^/]+$	25.707	6.427	12.989	12.989	0
0.3%	24	1	0	PATCH	^/api/posts/[^/]+$	20.406	0.850	5.714	5.714	620
0.3%	100	0	0	GET	^/api/users/[^/]+$	15.272	0.153	0.370	0.510	201
0.2%	53	0	0	GET	^/api/memories/[^/]+	11.659	0.220	1.315	1.315	1262
0.2%	1	0	0	POST	/api/me/account-deletion	10.819	10.819	10.819	10.819	0
0.2%	64	1	0	GET	/api/me	10.377	0.162	0.389	0.389	293
0.2%	17	0	0	DELETE	^/api/posts/[^/]+$	10.104	0.594	1.738	1.738	0
0.2%	67	0	0	GET	^/api/posts/[^/]+/comments$	9.606	0.143	0.374	0.374	2451
0.1%	50	0	0	GET	^/api/posts/[^/]+/realmojis$	8.138	0.163	0.486	0.486	2462
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.9%	150	0	0	GET	/api/notifications	1199.754	7.998	9.371	9.371	181
21.6%	174	3	0	POST	/api/posts	701.334	4.031	14.295	14.828	627
6.5%	228	2	0	POST	/api/friend-requests	211.446	0.927	8.244	14.286	114
5.8%	208	0	0	POST	/api/auth/login	187.438	0.901	8.559	12.235	250
5.5%	134	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	179.258	1.338	14.735	14.952	115
4.3%	806	3	0	GET	^/api/images/[^/]+	138.591	0.172	1.006	1.312	179588
3.1%	222	0	0	GET	/api/posting-window/current	99.782	0.449	5.292	5.359	240
2.9%	596	0	0	GET	^/api/feeds/[^/]+	95.042	0.159	0.493	1.031	1918
1.7%	46	7	0	POST	^/api/posts/[^/]+/realmojis$	54.793	1.191	7.340	7.340	354
1.7%	251	0	0	GET	/api/memories	54.647	0.218	0.973	1.863	11365
1.4%	17	0	0	GET	/api/me/export	47.059	2.768	12.171	12.171	2856220
1.0%	27	0	0	POST	/api/auth/logout	33.260	1.232	12.788	12.788	0
0.9%	18	1	0	POST	^/api/posts/[^/]+/comments$	30.100	1.672	5.959	5.959	325
0.8%	22	0	0	PUT	/api/me/profile-image	26.368	1.199	6.196	6.196	0
0.8%	39	0	0	PUT	^/api/follows/[^/]+$	24.914	0.639	5.531	5.531	0
0.7%	46	0	0	POST	/api/auth/register	23.503	0.511	3.807	3.807	230
0.7%	134	10	0	GET	^/api/posts/me/current$	22.881	0.171	0.628	0.897	585
0.6%	16	0	0	PATCH	/api/me	20.266	1.267	7.174	7.174	299
0.5%	114	0	0	GET	/api/users	16.858	0.148	0.368	0.417	229
0.5%	106	0	0	GET	^/api/relationships/[^/]+	15.434	0.146	0.513	0.632	1578
0.5%	3	0	0	DELETE	^/api/comments/[^/]+$	15.267	5.089	14.811	14.811	0
0.4%	89	8	0	GET	^/api/posts/[^/]+$	13.778	0.155	0.719	0.719	565
0.4%	22	1	0	PATCH	^/api/posts/[^/]+$	12.066	0.548	2.621	2.621	583
0.2%	30	0	0	GET	^/api/memories/[^/]+	5.921	0.197	0.538	0.538	1350
0.2%	35	0	0	GET	^/api/posts/[^/]+/realmojis$	5.361	0.153	0.403	0.403	1115
0.1%	30	0	0	GET	^/api/users/[^/]+$	4.645	0.155	0.329	0.329	202
0.1%	30	0	0	GET	^/api/posts/[^/]+/comments$	4.344	0.145	0.403	0.403	2127
0.1%	3	0	0	DELETE	^/api/posts/[^/]+$	3.367	1.122	1.919	1.919	0
0.1%	25	0	0	GET	/api/me	3.219	0.129	0.240	0.240	293
0.1%	4	0	0	PUT	^/api/blocks/[^/]+$	2.480	0.620	1.873	1.873	0
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 44409 クエリ / 81 種類 / 合計 302.441s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	60.5%	2583	182.866	70.80	526.92	657.9	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	24.9%	214	75.313	351.93	5876.83	7122.8	135.4	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	4.8%	1858	14.628	7.87	240.05	118.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.2%	5592	12.632	2.26	37.95	11.1	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.4%	1815	4.158	2.29	17.98	25.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.7%	1	2.193	2192.82	2192.82	29787.0	1199656.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.5%	1	1.657	1656.67	1656.67	20905.0	626689.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.5%	435	1.589	3.65	16.97	30.5	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.5%	2	1.404	701.81	1395.59	16790.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
10	0.3%	544	0.941	1.73	11.26	21.2	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.2%	1	0.741	740.90	740.90	11528.0	529908.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
12	0.2%	3471	0.734	0.21	30.62	19.9	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
13	0.1%	2	0.452	226.13	302.03	225883.5	2659.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.1%	1	0.390	390.48	390.48	24062.0	1578.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	3105	0.388	0.12	24.73	11.2	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

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
Build ID: 5e6be6e3c20dd8902144f39800933cf4335257d4
Type: cpu
Time: 2026-10-31 12:16:15 JST
Duration: 60.02s, Total samples = 53.51s (89.16%)
Showing nodes accounting for 41.60s, 77.74% of 53.51s total
Dropped 1011 nodes (cum <= 0.27s)
Showing top 40 nodes out of 253
      flat  flat%   sum%        cum   cum%
     4.43s  8.28%  8.28%      4.43s  8.28%  crypto/internal/fips140/sha256.blockAVX2
     3.31s  6.19% 14.46%      3.31s  6.19%  internal/runtime/syscall/linux.Syscall6
     3.10s  5.79% 20.26%      9.81s 18.33%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     2.82s  5.27% 25.53%      9.58s 17.90%  runtime.scanObject
     2.67s  4.99% 30.52%      4.34s  8.11%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     2.62s  4.90% 35.41%      2.62s  4.90%  runtime.memclrNoHeapPointers
     2.24s  4.19% 39.60%      2.36s  4.41%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     2.20s  4.11% 43.71%      2.74s  5.12%  runtime.findObject
     2.19s  4.09% 47.80%      2.19s  4.09%  runtime.memmove
     1.94s  3.63% 51.43%      3.03s  5.66%  runtime.tryDeferToSpanScan
     1.79s  3.35% 54.77%      5.75s 10.75%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.43s  2.67% 57.45%      1.43s  2.67%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.25s  2.34% 59.78%      1.76s  3.29%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     1.14s  2.13% 61.91%      1.14s  2.13%  aeshashbody
     0.82s  1.53% 63.45%      0.82s  1.53%  example.com/isureal-company/isureal.(*UserRow).active
     0.72s  1.35% 64.79%      0.72s  1.35%  memeqbody
     0.72s  1.35% 66.14%      0.72s  1.35%  runtime.spanClass.sizeclass (inline)
     0.52s  0.97% 67.11%      0.86s  1.61%  runtime.greyobject
     0.48s   0.9% 68.01%      0.50s  0.93%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.44s  0.82% 68.83%      0.62s  1.16%  example.com/isureal-company/isureal.postLess
     0.42s  0.78% 69.61%      2.01s  3.76%  runtime.mapaccess1_faststr
     0.41s  0.77% 70.38%      0.41s  0.77%  runtime.typePointers.nextFast (inline)
     0.37s  0.69% 71.07%      0.37s  0.69%  runtime.nextFreeFast (inline)
     0.34s  0.64% 71.71%      0.34s  0.64%  indexbody
     0.30s  0.56% 72.27%      0.30s  0.56%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.29s  0.54% 72.81%      0.34s  0.64%  runtime.typePointers.next
     0.28s  0.52% 73.33%      0.28s  0.52%  runtime.spanOf (inline)
     0.26s  0.49% 73.82%      5.03s  9.40%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.25s  0.47% 74.29%      4.95s  9.25%  crypto/internal/fips140/sha256.(*Digest).Write
     0.24s  0.45% 74.73%      0.67s  1.25%  runtime.scanObjectsSmall
     0.23s  0.43% 75.16%      0.50s  0.93%  database/sql.convertAssignRows
     0.20s  0.37% 75.54%      1.15s  2.15%  runtime.scanSpan
     0.19s  0.36% 75.89%      1.06s  1.98%  runtime.mallocgcSmallScanNoHeader
     0.17s  0.32% 76.21%      1.72s  3.21%  runtime.mapassign_faststr
     0.15s  0.28% 76.49%      3.09s  5.77%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.14s  0.26% 76.75%      0.27s   0.5%  compress/flate.(*compressor).deflate
     0.14s  0.26% 77.01%      5.88s 10.99%  example.com/isureal-company/isureal.computeImageSignature
     0.13s  0.24% 77.26%      5.54s 10.35%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.13s  0.24% 77.50%      5.25s  9.81%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.13s  0.24% 77.74%      0.28s  0.52%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 56.1% / 100% | 42% | 0.3% | 6 | 110 MB | CPU が飽和 |
| isu2 (2 cpu) | 10.7% / 31% | 0% | 0.6% | 4 | 2768 MB |  |
| isu3 (2 cpu) | 7.4% / 29% | 0% | 0.5% | 3 | 2874 MB |  |
| isu4 (2 cpu) | 45.2% / 100% | 14% | 16.2% | 30 | 121 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 92% / nginx 13% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 19% / nginx 1% |
| isu3 | isureal 13% |
| isu4 | postgres 47% / isureal 7% |

## Diff from previous run

スコア: 0 → 67674 (+67674)

`249be25..7d17432`

```
 .../etc/postgresql/16/main/conf.d/zz-tuning.conf   |   2 +
 webapp/go/account_deletion_api.go                  |   2 +-
 webapp/go/admin_user_api.go                        |   2 +-
 webapp/go/comment_api.go                           |   2 +-
 webapp/go/realmoji_api.go                          |   2 +-
 webapp/go/relationship.go                          |  10 ++-
 webapp/go/store.go                                 | 100 +++++++++++++++++----
 7 files changed, 98 insertions(+), 22 deletions(-)
```
