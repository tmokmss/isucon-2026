# run 20261031-122555-e42d0b5

## Summary

| | |
|---|---|
| score | 79369 |
| status | pass |
| ref | grandprix-post-fifo |
| sha | e42d0b56cef8 |
| 所要時間 | 119s |
| 前回比 | +14669 (22.7%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +14669 (22.7%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | jpeg fifo on 72 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	148	416	643	642
europe	7	20	45	100	101	101
america	6	13	33	62	63	63
```

```
{"id":"8085","status":"done","portalStatus":"finished","passed":true,"score":79369,"raw":79369,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:26:31.633Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:26:34.597Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:26:37.596Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:26:40.630Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T03:26:41.485Z level=INFO msg="user moved" user=evamero_5191 from=asia to=europe
time=2026-10-31T03:26:43.044Z level=INFO msg="user moved" user=yadanho5282 from=asia to=europe
time=2026-10-31T03:26:43.597Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T03:26:46.598Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T03:26:49.623Z level=INFO msg="asia 3 日目: 148 人"
time=2026-10-31T03:26:52.649Z level=INFO msg="europe 3 日目: 45 人"
time=2026-10-31T03:26:53.223Z level=INFO msg="user moved" user=andan-5995 from=asia to=america
time=2026-10-31T03:26:55.600Z level=INFO msg="america 3 日目: 33 人"
time=2026-10-31T03:26:58.750Z level=INFO msg="asia 4 日目: 416 人"
time=2026-10-31T03:26:59.749Z level=INFO msg="user moved" user=yadanho5282 from=europe to=america
time=2026-10-31T03:27:01.871Z level=INFO msg="europe 4 日目: 100 人"
time=2026-10-31T03:27:04.890Z level=INFO msg="america 4 日目: 62 人"
time=2026-10-31T03:27:05.826Z level=INFO msg="user moved" user=taku7178 from=asia to=europe
time=2026-10-31T03:27:07.840Z level=INFO msg="asia 5 日目: 643 人"
time=2026-10-31T03:27:08.419Z level=INFO msg="user moved" user=torkimax_7647 from=asia to=america
time=2026-10-31T03:27:10.691Z level=INFO msg="europe 5 日目: 101 人"
time=2026-10-31T03:27:13.778Z level=INFO msg="america 5 日目: 63 人"
time=2026-10-31T03:27:22.901Z level=INFO msg="europe 最終日（グランプリ）: 101 人"
time=2026-10-31T03:27:22.981Z level=INFO msg="america 最終日（グランプリ）: 63 人"
time=2026-10-31T03:27:23.157Z level=INFO msg="asia 最終日（グランプリ）: 642 人"
time=2026-10-31T03:27:31.620Z level=INFO msg="グランプリ: アジア 201 人, ヨーロッパ 18 人, アメリカ 16 人"
{
  "pass": true,
  "score": 79369,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.5%	1210	0	0	GET	/api/notifications	10711.150	8.852	15.353	15.570	173
17.9%	1715	45	0	POST	/api/posts	5401.912	3.150	10.146	12.315	634
10.7%	3450	48	0	POST	/api/friend-requests	3209.959	0.930	6.152	12.789	114
8.0%	9745	358	0	GET	^/api/images/[^/]+	2409.993	0.247	3.970	10.787	139229
6.0%	4414	1368	0	POST	/api/auth/login	1803.596	0.409	4.960	11.514	177
5.5%	1252	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1666.041	1.331	7.294	12.752	116
2.0%	663	1	0	POST	/api/auth/register	600.012	0.905	6.509	8.721	227
2.0%	396	35	0	POST	^/api/posts/[^/]+/realmojis$	597.412	1.509	9.923	10.944	453
1.7%	494	127	0	PUT	/api/me/profile-image	510.982	1.034	5.642	6.810	6
1.6%	344	25	0	POST	^/api/posts/[^/]+/comments$	473.019	1.375	10.683	12.369	318
1.3%	471	0	0	PUT	^/api/follows/[^/]+$	403.197	0.856	6.151	8.213	0
1.0%	8287	0	0	GET	^/api/feeds/[^/]+	309.843	0.037	0.226	1.855	1346
1.0%	318	0	0	PATCH	/api/me	308.085	0.969	5.213	7.862	279
0.9%	2621	0	0	GET	/api/memories	283.915	0.108	0.517	7.330	10500
0.8%	518	0	0	POST	/api/auth/logout	249.798	0.482	6.484	11.861	0
0.7%	172	7	0	PATCH	^/api/posts/[^/]+$	206.427	1.200	6.564	6.771	625
0.6%	2421	0	0	GET	/api/posting-window/current	193.726	0.080	0.650	2.424	237
0.6%	204	0	0	GET	/api/me/export	192.324	0.943	8.239	10.089	2046051
0.5%	78	0	0	DELETE	^/api/posts/[^/]+$	145.295	1.863	8.628	8.628	0
0.2%	1327	71	0	GET	^/api/posts/me/current$	66.575	0.050	0.494	0.995	615
0.2%	1664	0	0	GET	/api/users	61.807	0.037	0.247	0.997	250
0.2%	58	0	0	DELETE	^/api/realmojis/[^/]+$	54.809	0.945	5.192	5.192	0
0.2%	53	0	0	DELETE	^/api/comments/[^/]+$	54.192	1.022	13.944	13.944	0
0.1%	1193	0	0	GET	^/api/relationships/[^/]+	36.774	0.031	0.197	0.611	882
0.1%	25	0	0	PUT	^/api/blocks/[^/]+$	33.360	1.334	6.428	6.428	0
0.1%	1053	19	0	GET	^/api/posts/[^/]+$	31.125	0.030	0.174	0.442	614
0.1%	459	0	0	GET	^/api/memories/[^/]+	26.308	0.057	0.344	0.991	665
0.1%	626	0	0	GET	^/api/users/[^/]+$	20.316	0.032	0.203	1.003	205
0.1%	1	0	0	POST	/api/initialize	15.455	15.455	15.455	15.455	90
0.0%	456	0	0	GET	^/api/posts/[^/]+/realmojis$	12.509	0.027	0.151	0.233	1538
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.1%	245	0	0	GET	/api/notifications	2083.355	8.503	12.321	12.360	178
19.7%	298	5	0	POST	/api/posts	1105.773	3.711	11.390	12.173	639
11.2%	563	3	0	POST	/api/friend-requests	628.728	1.117	7.711	13.329	114
6.2%	397	0	0	POST	/api/auth/login	345.635	0.871	5.959	7.752	248
4.1%	1386	8	0	GET	^/api/images/[^/]+	228.879	0.165	1.063	1.441	161962
3.8%	209	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	213.613	1.022	5.955	8.520	116
3.3%	1247	0	0	GET	^/api/feeds/[^/]+	185.731	0.149	0.314	1.740	1777
1.6%	400	0	0	GET	/api/memories	89.053	0.223	0.604	0.896	12252
1.4%	407	0	0	GET	/api/posting-window/current	79.957	0.196	0.733	0.920	239
1.3%	60	4	0	POST	^/api/posts/[^/]+/realmojis$	70.267	1.171	10.074	10.074	448
1.2%	29	0	0	GET	/api/me/export	67.357	2.323	11.160	11.160	2642670
1.1%	97	0	0	POST	/api/auth/register	61.107	0.630	3.221	3.221	228
1.0%	35	4	0	POST	^/api/posts/[^/]+/comments$	55.427	1.584	7.331	7.331	263
0.9%	68	0	0	PUT	^/api/follows/[^/]+$	49.073	0.722	5.093	5.093	0
0.8%	297	0	0	GET	/api/users	45.166	0.152	0.360	0.429	255
0.8%	51	0	0	POST	/api/auth/logout	44.695	0.876	13.690	13.690	0
0.6%	51	0	0	PUT	/api/me/profile-image	36.463	0.715	3.424	3.424	0
0.6%	223	13	0	GET	^/api/posts/me/current$	35.309	0.158	0.465	0.596	610
0.5%	201	0	0	GET	^/api/relationships/[^/]+	29.794	0.148	0.383	0.540	1055
0.5%	26	2	0	PATCH	^/api/posts/[^/]+$	29.513	1.135	4.854	4.854	600
0.5%	44	0	0	PATCH	/api/me	27.080	0.615	2.443	2.443	286
0.4%	164	4	0	GET	^/api/posts/[^/]+$	23.911	0.146	0.362	0.409	599
0.2%	8	0	0	DELETE	^/api/posts/[^/]+$	13.111	1.639	5.485	5.485	0
0.2%	92	0	0	GET	^/api/posts/[^/]+/realmojis$	13.050	0.142	0.311	0.311	1117
0.2%	88	0	0	GET	^/api/users/[^/]+$	12.272	0.139	0.271	0.271	193
0.2%	50	0	0	GET	^/api/memories/[^/]+	9.256	0.185	0.534	0.534	788
0.2%	56	2	0	GET	/api/me	8.751	0.156	0.441	0.441	290
0.1%	49	0	0	GET	^/api/posts/[^/]+/comments$	7.312	0.149	0.328	0.328	1506
0.1%	5	0	0	DELETE	^/api/realmojis/[^/]+$	7.063	1.413	3.211	3.211	0
0.1%	1	0	0	POST	/api/me/account-deletion	4.701	4.701	4.701	4.701	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.1%	159	0	0	GET	/api/notifications	1238.254	7.788	9.299	9.299	177
19.5%	192	2	0	POST	/api/posts	669.115	3.485	10.298	10.792	638
9.6%	338	9	0	POST	/api/friend-requests	330.293	0.977	6.332	9.331	113
6.7%	225	0	0	POST	/api/auth/login	229.534	1.020	7.332	10.085	251
4.6%	882	4	0	GET	^/api/images/[^/]+	156.005	0.177	0.968	1.452	165905
3.8%	117	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	129.482	1.107	3.700	9.528	115
2.8%	597	0	0	GET	^/api/feeds/[^/]+	95.376	0.160	0.508	0.703	1677
2.0%	50	6	0	POST	^/api/posts/[^/]+/comments$	70.115	1.402	8.989	8.989	286
1.8%	272	0	0	GET	/api/memories	60.691	0.223	0.543	0.572	14001
1.5%	256	0	0	GET	/api/posting-window/current	51.376	0.201	0.817	1.012	240
1.5%	29	0	0	PUT	/api/me/profile-image	50.582	1.744	8.693	8.693	0
1.3%	21	6	0	POST	^/api/posts/[^/]+/realmojis$	45.145	2.150	8.693	8.693	302
1.3%	20	0	0	GET	/api/me/export	43.674	2.184	5.854	5.854	2859405
1.0%	59	0	0	POST	/api/auth/register	32.680	0.554	3.413	3.413	229
0.9%	36	0	0	POST	/api/auth/logout	30.426	0.845	3.873	3.873	0
0.8%	150	10	0	GET	^/api/posts/me/current$	28.751	0.192	0.811	0.811	602
0.8%	29	0	0	PATCH	/api/me	26.950	0.929	4.545	4.545	267
0.8%	175	0	0	GET	/api/users	26.854	0.153	0.362	0.621	249
0.7%	16	2	0	PATCH	^/api/posts/[^/]+$	22.827	1.427	5.161	5.161	554
0.5%	30	0	0	PUT	^/api/follows/[^/]+$	18.779	0.626	5.462	5.462	0
0.5%	111	0	0	GET	^/api/relationships/[^/]+	16.496	0.149	0.279	0.288	1104
0.3%	80	2	0	GET	^/api/posts/[^/]+$	11.623	0.145	0.251	0.251	605
0.3%	5	0	0	DELETE	^/api/posts/[^/]+$	9.499	1.900	3.185	3.185	0
0.2%	47	0	0	GET	^/api/users/[^/]+$	6.795	0.145	0.250	0.250	183
0.2%	39	0	0	GET	^/api/posts/[^/]+/comments$	5.642	0.145	0.316	0.316	1423
0.1%	32	0	0	GET	/api/me	5.003	0.156	0.289	0.289	296
0.1%	20	0	0	GET	^/api/memories/[^/]+	4.488	0.224	0.389	0.389	1319
0.1%	1	0	0	POST	/api/me/account-deletion	3.703	3.703	3.703	3.703	0
0.1%	20	0	0	GET	^/api/posts/[^/]+/realmojis$	2.644	0.132	0.237	0.237	1116
0.0%	5	0	0	DELETE	^/api/comments/[^/]+$	1.539	0.308	0.845	0.845	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 50942 クエリ / 80 種類 / 合計 278.459s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	50.0%	3035	139.333	45.91	473.10	611.5	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	32.5%	253	90.380	357.23	1294.71	7488.1	139.3	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.4%	1918	15.132	7.89	253.45	117.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	5.0%	6466	14.013	2.17	42.03	10.6	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.7%	2156	4.686	2.17	23.01	25.5	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.0%	2	2.789	1394.47	2632.78	227479.5	2678.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.8%	1	2.258	2257.78	2257.78	29904.0	1199881.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.6%	1	1.793	1792.74	1792.74	21199.0	626927.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.5%	432	1.438	3.33	18.81	29.4	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.3%	2	0.946	473.04	938.01	16764.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.3%	4278	0.859	0.20	13.80	19.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
12	0.3%	1	0.773	772.56	772.56	11600.0	529886.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
13	0.3%	394	0.699	1.77	9.65	21.9	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
14	0.2%	3668	0.464	0.13	17.14	11.4	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
15	0.1%	1512	0.355	0.23	10.32	20.7	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)

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
Build ID: 914bfaf5d2c97645b53090fbd8ce1b981377fb49
Type: cpu
Time: 2026-10-31 12:26:10 JST
Duration: 60.03s, Total samples = 47.88s (79.76%)
Showing nodes accounting for 36.04s, 75.27% of 47.88s total
Dropped 1025 nodes (cum <= 0.24s)
Showing top 40 nodes out of 292
      flat  flat%   sum%        cum   cum%
     3.38s  7.06%  7.06%      3.38s  7.06%  internal/runtime/syscall/linux.Syscall6
     2.89s  6.04% 13.10%      9.49s 19.82%  runtime.scanObject
     2.78s  5.81% 18.90%      2.78s  5.81%  runtime.memclrNoHeapPointers
     2.21s  4.62% 23.52%      2.95s  6.16%  runtime.findObject
     1.97s  4.11% 27.63%      1.97s  4.11%  runtime.memmove
     1.93s  4.03% 31.66%      3.04s  6.35%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.89s  3.95% 35.61%      6.36s 13.28%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.73s  3.61% 39.22%      1.73s  3.61%  crypto/internal/fips140/sha256.blockAVX2
     1.65s  3.45% 42.67%      5.72s 11.95%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.50s  3.13% 45.80%      2.90s  6.06%  runtime.tryDeferToSpanScan
     1.37s  2.86% 48.66%      1.42s  2.97%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     1.37s  2.86% 51.52%      1.37s  2.86%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.32s  2.76% 54.28%      1.32s  2.76%  aeshashbody
     1.04s  2.17% 56.45%      1.04s  2.17%  example.com/isureal-company/isureal.(*UserRow).active
        1s  2.09% 58.54%         1s  2.09%  runtime.spanClass.sizeclass (inline)
     0.84s  1.75% 60.30%      1.15s  2.40%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.63s  1.32% 61.61%      0.63s  1.32%  memeqbody
     0.57s  1.19% 62.80%      2.05s  4.28%  runtime.mapaccess1_faststr
     0.45s  0.94% 63.74%      0.45s  0.94%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.43s   0.9% 64.64%      0.68s  1.42%  example.com/isureal-company/isureal.postLess
     0.42s  0.88% 65.52%      0.81s  1.69%  runtime.greyobject
     0.39s  0.81% 66.33%      0.48s  1.00%  runtime.typePointers.next
     0.31s  0.65% 66.98%      0.31s  0.65%  runtime.nextFreeFast (inline)
     0.30s  0.63% 67.61%      0.30s  0.63%  runtime.(*mspan).divideByElemSize (inline)
     0.29s  0.61% 68.21%      0.29s  0.61%  indexbody
     0.28s  0.58% 68.80%      0.73s  1.52%  database/sql.convertAssignRows
     0.28s  0.58% 69.38%      0.30s  0.63%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.28s  0.58% 69.97%      0.28s  0.58%  runtime.typePointers.nextFast (inline)
     0.27s  0.56% 70.53%      0.27s  0.56%  runtime.futex
     0.27s  0.56% 71.09%      0.33s  0.69%  runtime.spanOf (inline)
     0.26s  0.54% 71.64%      0.26s  0.54%  encoding/json.appendString[go.shape.string]
     0.26s  0.54% 72.18%      0.26s  0.54%  runtime.madvise
     0.26s  0.54% 72.72%      0.79s  1.65%  runtime.scanObjectsSmall
     0.21s  0.44% 73.16%      1.98s  4.14%  runtime.mapassign_faststr
     0.19s   0.4% 73.56%      0.26s  0.54%  runtime.heapArenaOf (inline)
     0.18s  0.38% 73.93%      0.30s  0.63%  compress/flate.(*compressor).deflate
     0.17s  0.36% 74.29%      1.19s  2.49%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.17s  0.36% 74.64%      6.85s 14.31%  runtime.mallocgc
     0.17s  0.36% 75.00%      0.25s  0.52%  runtime.scanObjectSmall
     0.13s  0.27% 75.27%      1.08s  2.26%  runtime.mallocgcSmallScanNoHeader
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 53.3% / 100% | 39% | 0.1% | 6 | 182 MB | CPU が飽和 |
| isu2 (2 cpu) | 15.6% / 47% | 0% | 0.5% | 7 | 2706 MB | run queue が積もっている |
| isu3 (2 cpu) | 11.8% / 44% | 0% | 0.5% | 4 | 2777 MB |  |
| isu4 (2 cpu) | 51.1% / 100% | 17% | 20.7% | 28 | 111 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 85% / nginx 14% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 29% / nginx 1% |
| isu3 | isureal 22% / nginx 1% |
| isu4 | postgres 55% / isureal 2% |

## Diff from previous run

スコア: 64700 → 79369 (+14669)

`e0e4cb3..e42d0b5`

```
 webapp/go/account_deletion_api.go |   2 +-
 webapp/go/admin_user_api.go       |   2 +-
 webapp/go/comment_api.go          |   2 +-
 webapp/go/offload.go              |  27 ++++++++--
 webapp/go/realmoji_api.go         |   2 +-
 webapp/go/relationship.go         |  10 +++-
 webapp/go/store.go                | 100 ++++++++++++++++++++++++++++++++------
 7 files changed, 119 insertions(+), 26 deletions(-)
```
