# run 20261031-135709-3bc4182

## Summary

| | |
|---|---|
| score | 114092 |
| status | pass |
| ref | jpeg-peer-asia2 |
| sha | 3bc418272173 |
| 所要時間 | 386s |
| 前回比 | -5225 (-4.4%) vs 20261031-135233-2d10aad |
| ベスト比 | -5225 (-4.4%) vs 20261031-135233-2d10aad = 119317 |
| note | JPEG の検証を asia2 にも逃がす（jpegcheck-fast の上） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	145	404	851	851
europe	7	19	55	157	157	156
america	6	19	52	124	128	128
```

```
{"id":"15370","status":"done","portalStatus":"finished","passed":true,"score":114092,"raw":114092,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:02:00.046Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:02:03.027Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:02:06.017Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:02:09.043Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T05:02:11.691Z level=INFO msg="user moved" user=tsuso_5284 from=asia to=america
time=2026-10-31T05:02:12.016Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T05:02:15.019Z level=INFO msg="america 2 日目: 19 人"
time=2026-10-31T05:02:18.036Z level=INFO msg="asia 3 日目: 145 人"
time=2026-10-31T05:02:21.024Z level=INFO msg="europe 3 日目: 55 人"
time=2026-10-31T05:02:22.012Z level=INFO msg="user moved" user=neya.5982 from=asia to=europe
time=2026-10-31T05:02:22.047Z level=INFO msg="user moved" user=liakeha.6142 from=asia to=europe
time=2026-10-31T05:02:24.017Z level=INFO msg="america 3 日目: 52 人"
time=2026-10-31T05:02:27.056Z level=INFO msg="asia 4 日目: 404 人"
time=2026-10-31T05:02:30.151Z level=INFO msg="europe 4 日目: 157 人"
time=2026-10-31T05:02:31.016Z level=INFO msg="user moved" user=koanmax.7337 from=asia to=europe
time=2026-10-31T05:02:31.777Z level=INFO msg="user moved" user=tsuso_5284 from=america to=asia
time=2026-10-31T05:02:33.046Z level=INFO msg="america 4 日目: 124 人"
time=2026-10-31T05:02:33.491Z level=INFO msg="user moved" user=liakeha.6142 from=europe to=america
time=2026-10-31T05:02:34.198Z level=INFO msg="user moved" user=tsuso_5284 from=asia to=america
time=2026-10-31T05:02:34.663Z level=INFO msg="user moved" user=tsushian_8118 from=asia to=america
time=2026-10-31T05:02:34.748Z level=INFO msg="user moved" user=ivymu-7466 from=asia to=america
time=2026-10-31T05:02:36.284Z level=INFO msg="asia 5 日目: 851 人"
time=2026-10-31T05:02:39.081Z level=INFO msg="europe 5 日目: 157 人"
time=2026-10-31T05:02:42.291Z level=INFO msg="america 5 日目: 128 人"
time=2026-10-31T05:02:51.104Z level=INFO msg="europe 最終日（グランプリ）: 156 人"
time=2026-10-31T05:02:51.158Z level=INFO msg="america 最終日（グランプリ）: 128 人"
time=2026-10-31T05:02:51.278Z level=INFO msg="asia 最終日（グランプリ）: 851 人"
time=2026-10-31T05:03:00.052Z level=INFO msg="グランプリ: アジア 403 人, ヨーロッパ 85 人, アメリカ 64 人"
{
  "pass": true,
  "score": 114092,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.5%	1391	1	0	GET	/api/notifications	11709.651	8.418	15.030	15.046	169
16.4%	2026	52	0	POST	/api/posts	5740.695	2.834	9.755	12.154	635
13.8%	4588	57	0	POST	/api/friend-requests	4830.427	1.053	7.746	15.401	114
8.6%	11978	479	0	GET	^/api/images/[^/]+	3009.652	0.251	4.424	13.114	143251
5.9%	6326	2585	0	POST	/api/auth/login	2059.593	0.326	5.111	12.377	156
5.0%	1642	8	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1753.991	1.068	7.444	13.611	116
3.3%	603	114	0	PUT	/api/me/profile-image	1149.015	1.905	8.829	14.409	4
2.0%	399	34	0	POST	^/api/posts/[^/]+/realmojis$	685.620	1.718	10.762	13.023	432
1.9%	516	24	0	POST	^/api/posts/[^/]+/comments$	680.307	1.318	8.245	15.002	327
1.9%	873	1	0	POST	/api/auth/register	669.415	0.767	6.439	9.093	227
1.6%	599	0	0	PUT	^/api/follows/[^/]+$	560.201	0.935	6.195	8.746	0
1.0%	428	0	0	PATCH	/api/me	365.045	0.853	5.000	8.295	275
0.9%	3093	0	0	GET	/api/memories	312.264	0.101	0.604	2.462	10343
0.7%	267	0	0	GET	/api/me/export	236.121	0.884	4.652	8.749	1977217
0.6%	610	0	0	POST	/api/auth/logout	226.853	0.372	6.411	10.267	0
0.6%	186	9	0	PATCH	^/api/posts/[^/]+$	207.788	1.117	6.256	6.775	626
0.5%	10219	0	0	GET	^/api/feeds/[^/]+	171.570	0.017	0.122	0.947	1859
0.4%	77	0	0	DELETE	^/api/posts/[^/]+$	147.514	1.916	9.422	9.422	0
0.2%	2889	0	0	GET	/api/posting-window/current	57.301	0.020	0.140	2.142	237
0.2%	61	0	0	DELETE	^/api/realmojis/[^/]+$	54.867	0.899	4.608	4.608	0
0.1%	32	0	0	PUT	^/api/blocks/[^/]+$	43.313	1.354	6.764	6.764	0
0.1%	1687	77	0	GET	^/api/posts/me/current$	40.080	0.024	0.148	0.212	622
0.1%	74	0	0	DELETE	^/api/comments/[^/]+$	39.305	0.531	4.584	4.584	0
0.1%	2189	0	0	GET	/api/users	33.686	0.015	0.113	0.172	241
0.1%	577	0	0	GET	^/api/memories/[^/]+	32.256	0.056	0.479	1.699	646
0.1%	1456	0	0	GET	^/api/relationships/[^/]+	26.338	0.018	0.123	4.392	946
0.1%	1520	13	0	GET	^/api/posts/[^/]+$	19.789	0.013	0.104	0.137	620
0.1%	1	0	0	POST	/api/initialize	19.317	19.317	19.317	19.317	90
0.0%	692	0	0	GET	^/api/users/[^/]+$	10.081	0.015	0.121	0.197	203
0.0%	9	0	0	DELETE	^/api/follows/[^/]+$	9.341	1.038	2.834	2.834	0
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.5%	373	0	0	GET	/api/notifications	3078.410	8.253	12.045	12.239	176
15.0%	486	4	0	POST	/api/posts	1201.561	2.472	9.477	10.632	642
9.7%	935	11	0	POST	/api/friend-requests	776.603	0.831	5.927	7.028	114
6.2%	579	0	0	POST	/api/auth/login	494.385	0.854	5.876	10.611	248
5.0%	284	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	400.718	1.411	7.965	10.365	115
4.4%	2136	21	0	GET	^/api/images/[^/]+	352.910	0.165	1.036	1.605	188058
4.4%	1995	0	0	GET	^/api/feeds/[^/]+	349.826	0.175	0.654	0.932	2203
1.7%	608	0	0	GET	/api/memories	132.769	0.218	0.865	1.352	10848
1.6%	663	0	0	GET	/api/posting-window/current	124.534	0.188	0.522	0.896	238
1.3%	91	5	0	POST	^/api/posts/[^/]+/comments$	102.259	1.124	7.100	7.100	335
1.2%	56	4	0	POST	^/api/posts/[^/]+/realmojis$	96.594	1.725	5.581	5.581	425
1.2%	33	0	0	GET	/api/me/export	91.877	2.784	11.256	11.256	2531326
1.1%	425	21	0	GET	^/api/posts/me/current$	86.177	0.203	0.697	0.718	616
1.0%	104	0	0	PUT	^/api/follows/[^/]+$	77.802	0.748	6.635	6.782	0
1.0%	440	0	0	GET	/api/users	76.311	0.173	0.516	0.943	236
0.9%	153	0	0	POST	/api/auth/register	75.694	0.495	2.612	3.107	228
0.8%	44	9	0	PATCH	^/api/posts/[^/]+$	64.315	1.462	7.039	7.039	515
0.8%	85	0	0	PUT	/api/me/profile-image	62.431	0.734	4.643	4.643	0
0.7%	84	0	0	POST	/api/auth/logout	57.930	0.690	3.915	3.915	0
0.6%	72	0	0	PATCH	/api/me	45.317	0.629	2.650	2.650	280
0.6%	270	0	0	GET	^/api/relationships/[^/]+	45.039	0.167	0.612	0.977	1028
0.5%	274	8	0	GET	^/api/posts/[^/]+$	43.575	0.159	0.489	0.983	615
0.4%	22	0	0	DELETE	^/api/posts/[^/]+$	32.476	1.476	8.036	8.036	0
0.3%	141	0	0	GET	^/api/users/[^/]+$	24.135	0.171	0.518	0.835	199
0.2%	114	0	0	GET	^/api/posts/[^/]+/comments$	17.966	0.158	0.516	0.809	1451
0.2%	86	0	0	GET	^/api/memories/[^/]+	17.265	0.201	0.736	0.736	661
0.2%	109	0	0	GET	^/api/posts/[^/]+/realmojis$	16.765	0.154	0.474	0.809	367
0.2%	91	2	0	GET	/api/me	14.309	0.157	0.996	0.996	293
0.1%	8	0	0	DELETE	^/api/realmojis/[^/]+$	8.071	1.009	3.049	3.049	0
0.1%	7	0	0	PUT	^/api/blocks/[^/]+$	7.160	1.023	3.453	3.453	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.9%	285	0	0	GET	/api/notifications	2096.204	7.355	9.221	9.221	178
15.0%	376	5	0	POST	/api/posts	901.577	2.398	9.781	11.865	640
13.7%	701	12	0	POST	/api/friend-requests	822.986	1.174	8.413	11.646	113
6.1%	422	0	0	POST	/api/auth/login	369.115	0.875	5.998	9.529	248
4.5%	1533	15	0	GET	^/api/images/[^/]+	269.122	0.176	0.756	1.896	188265
4.4%	231	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	265.400	1.149	6.774	10.287	116
3.8%	1339	0	0	GET	^/api/feeds/[^/]+	230.158	0.172	0.611	1.232	1987
2.0%	51	12	0	POST	^/api/posts/[^/]+/realmojis$	117.727	2.308	7.652	7.652	364
1.7%	461	0	0	GET	/api/memories	100.374	0.218	0.744	1.258	10592
1.4%	508	0	0	GET	/api/posting-window/current	86.052	0.169	0.584	0.668	240
1.2%	63	3	0	POST	^/api/posts/[^/]+/comments$	71.795	1.140	5.941	5.941	314
1.2%	29	0	0	GET	/api/me/export	70.146	2.419	6.601	6.601	3473706
1.0%	341	0	0	GET	/api/users	60.085	0.176	0.482	0.630	235
1.0%	324	20	0	GET	^/api/posts/me/current$	59.882	0.185	0.615	0.636	612
1.0%	62	0	0	PUT	/api/me/profile-image	59.504	0.960	5.539	5.539	0
1.0%	123	0	0	POST	/api/auth/register	58.940	0.479	1.902	2.028	229
1.0%	56	0	0	POST	/api/auth/logout	58.772	1.050	6.130	6.130	0
0.9%	72	0	0	PUT	^/api/follows/[^/]+$	52.523	0.729	9.532	9.532	0
0.7%	52	0	0	PATCH	/api/me	41.117	0.791	3.920	3.920	272
0.7%	25	5	0	PATCH	^/api/posts/[^/]+$	39.654	1.586	6.073	6.073	535
0.6%	19	0	0	DELETE	^/api/posts/[^/]+$	33.778	1.778	4.724	4.724	0
0.5%	187	0	0	GET	^/api/relationships/[^/]+	27.319	0.146	0.435	0.480	1069
0.4%	159	9	0	GET	^/api/posts/[^/]+$	24.564	0.154	0.456	0.478	589
0.3%	106	0	0	GET	^/api/users/[^/]+$	17.214	0.162	0.393	0.439	199
0.2%	54	0	0	GET	^/api/memories/[^/]+	14.017	0.260	1.301	1.301	938
0.2%	66	1	0	GET	/api/me	13.581	0.206	0.458	0.458	299
0.2%	70	0	0	GET	^/api/posts/[^/]+/realmojis$	10.581	0.151	0.386	0.386	335
0.2%	6	0	0	DELETE	^/api/realmojis/[^/]+$	10.546	1.758	3.657	3.657	0
0.2%	58	0	0	GET	^/api/posts/[^/]+/comments$	9.487	0.164	0.412	0.412	1135
0.1%	3	0	0	DELETE	^/api/follows/[^/]+$	8.186	2.729	6.577	6.577	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 67046 クエリ / 81 種類 / 合計 618.640s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	63.4%	3919	392.036	100.03	807.46	683.8	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	22.8%	329	141.188	429.14	2057.76	1361.5	138.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	5.1%	3082	31.296	10.15	382.87	119.6	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	3.8%	7791	23.323	2.99	53.56	10.6	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.5%	2827	9.422	3.33	43.84	24.9	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.9%	2	5.373	2686.34	5215.26	233821.5	2753.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	1	2.438	2438.12	2438.12	31103.0	1200044.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.4%	456	2.213	4.85	35.93	30.5	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.3%	1	1.969	1968.82	1968.82	24709.0	627814.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
10	0.2%	6132	1.493	0.24	20.70	19.6	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.2%	638	1.306	2.05	13.99	21.7	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.2%	1	1.080	1080.01	1080.01	12760.0	530141.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
13	0.2%	1	0.986	985.59	985.59	30116.0	2703.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	2064	0.754	0.37	43.51	20.4	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	4742	0.704	0.15	25.42	11.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

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
Build ID: bba03a53a97b1c3907162179057cf86ee8696941
Type: cpu
Time: 2026-10-31 14:01:34 JST
Duration: 60s, Total samples = 36.18s (60.30%)
Showing nodes accounting for 26.01s, 71.89% of 36.18s total
Dropped 962 nodes (cum <= 0.18s)
Showing top 40 nodes out of 299
      flat  flat%   sum%        cum   cum%
     3.65s 10.09% 10.09%      3.65s 10.09%  internal/runtime/syscall/linux.Syscall6
     2.34s  6.47% 16.56%      2.34s  6.47%  crypto/internal/fips140/sha256.blockAVX2
     1.82s  5.03% 21.59%      5.99s 16.56%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.66s  4.59% 26.17%      1.66s  4.59%  runtime.memmove
     1.58s  4.37% 30.54%      1.58s  4.37%  runtime.memclrNoHeapPointers
     1.52s  4.20% 34.74%      1.98s  5.47%  runtime.findObject
     1.43s  3.95% 38.70%      1.43s  3.95%  aeshashbody
     1.40s  3.87% 42.56%      1.40s  3.87%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.39s  3.84% 46.41%      5.32s 14.70%  runtime.scanObject
     1.20s  3.32% 49.72%      1.98s  5.47%  runtime.tryDeferToSpanScan
     0.68s  1.88% 51.60%      1.06s  2.93%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.55s  1.52% 53.12%      2.09s  5.78%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.53s  1.46% 54.59%      0.53s  1.46%  runtime.spanClass.sizeclass (inline)
     0.48s  1.33% 55.91%      0.48s  1.33%  memeqbody
     0.45s  1.24% 57.16%      0.48s  1.33%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.43s  1.19% 58.35%      2.06s  5.69%  runtime.mapaccess1_faststr
     0.40s  1.11% 59.45%      0.60s  1.66%  example.com/isureal-company/isureal.postLess
     0.39s  1.08% 60.53%      0.40s  1.11%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.39s  1.08% 61.61%      0.39s  1.08%  runtime.nextFreeFast (inline)
     0.36s     1% 62.60%      0.76s  2.10%  database/sql.convertAssignRows
     0.30s  0.83% 63.43%      0.30s  0.83%  runtime.futex
     0.29s   0.8% 64.23%      1.18s  3.26%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.66% 64.90%      0.24s  0.66%  runtime.(*mspan).writeHeapBitsSmall
     0.22s  0.61% 65.51%      0.22s  0.61%  sync/atomic.(*Int32).Add
     0.21s  0.58% 66.09%      0.22s  0.61%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.20s  0.55% 66.64%      3.61s  9.98%  runtime.mallocgc
     0.20s  0.55% 67.19%      0.20s  0.55%  runtime.spanOf (inline)
     0.18s   0.5% 67.69%      1.88s  5.20%  runtime.mapassign_faststr
     0.17s  0.47% 68.16%      0.43s  1.19%  compress/flate.(*compressor).deflate
     0.17s  0.47% 68.63%      0.71s  1.96%  runtime.mapaccess2_faststr
     0.16s  0.44% 69.07%      2.97s  8.21%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.15s  0.41% 69.49%      1.47s  4.06%  example.com/isureal-company/isureal.listMemories
     0.14s  0.39% 69.87%      2.62s  7.24%  crypto/internal/fips140/sha256.(*Digest).Write
     0.12s  0.33% 70.20%      0.35s  0.97%  runtime.greyobject
     0.11s   0.3% 70.51%      2.18s  6.03%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.10s  0.28% 70.78%      0.26s  0.72%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.10s  0.28% 71.06%      1.22s  3.37%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.10s  0.28% 71.34%      0.25s  0.69%  internal/runtime/maps.(*Iter).Next
     0.10s  0.28% 71.61%      0.19s  0.53%  runtime.(*mspan).markBitsForIndex (inline)
     0.10s  0.28% 71.89%      0.85s  2.35%  runtime.mallocgcSmallNoscan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 47.1% / 100% | 28% | 0.4% | 10 | 208 MB | run queue が積もっている |
| isu2 (2 cpu) | 15.0% / 68% | 0% | 1.0% | 4 | 1910 MB |  |
| isu3 (2 cpu) | 11.1% / 75% | 0% | 1.0% | 3 | 1657 MB |  |
| isu4 (2 cpu) | 55.9% / 100% | 31% | 14.9% | 35 | 111 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 69% / nginx 10% |
| isu2 | isureal 27% / nginx 1% |
| isu3 | isureal 18% / nginx 1% |
| isu4 | postgres 59% / isureal 7% / kcompactd0 1% |

## Diff from previous run

スコア: 119317 → 114092 (-5225)

`2d10aad..3bc4182`

```
 conf/isu1/home/isucon/env.sh | 2 +-
 webapp/go/offload.go         | 2 +-
 2 files changed, 2 insertions(+), 2 deletions(-)
```
