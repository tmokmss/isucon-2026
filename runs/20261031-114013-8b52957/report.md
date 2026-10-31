# run 20261031-114013-8b52957

## Summary

| | |
|---|---|
| score | 25636 |
| status | pass |
| ref | limiter-priority |
| sha | 8b529573d8ea |
| 所要時間 | 151s |
| 前回比 | +8564 (50.2%) vs 20261031-113441-6a7a33a |
| ベスト比 | +835 (3.4%) vs 20261031-103811-cef5fe3 = 24801 **ベスト更新** |
| note | limiter post4 mem2 heavy1 wait8s |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	124	124	124	123
europe	8	18	22	22	22	21
america	6	13	26	26	26	25
```

```
{"id":"4904","status":"done","portalStatus":"finished","passed":true,"score":25636,"raw":25636,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:41:30.179Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:41:31.885Z level=INFO msg="user moved" user=chishi_5054 from=asia to=europe
time=2026-10-31T02:41:33.140Z level=INFO msg="europe 1 日目: 8 人"
time=2026-10-31T02:41:36.128Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:41:39.189Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T02:41:42.166Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T02:41:45.145Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T02:41:48.298Z level=INFO msg="asia 3 日目: 124 人"
time=2026-10-31T02:41:51.177Z level=INFO msg="europe 3 日目: 22 人"
time=2026-10-31T02:41:54.301Z level=INFO msg="america 3 日目: 26 人"
time=2026-10-31T02:41:57.187Z level=INFO msg="asia 4 日目: 124 人"
time=2026-10-31T02:42:00.243Z level=INFO msg="europe 4 日目: 22 人"
time=2026-10-31T02:42:03.299Z level=INFO msg="america 4 日目: 26 人"
time=2026-10-31T02:42:06.274Z level=INFO msg="asia 5 日目: 124 人"
time=2026-10-31T02:42:09.293Z level=INFO msg="europe 5 日目: 22 人"
time=2026-10-31T02:42:12.239Z level=INFO msg="america 5 日目: 26 人"
time=2026-10-31T02:42:21.176Z level=INFO msg="america 最終日（グランプリ）: 25 人"
time=2026-10-31T02:42:21.439Z level=INFO msg="asia 最終日（グランプリ）: 123 人"
time=2026-10-31T02:42:21.563Z level=INFO msg="europe 最終日（グランプリ）: 21 人"
time=2026-10-31T02:42:30.130Z level=INFO msg="グランプリ: アジア 78 人, ヨーロッパ 11 人, アメリカ 14 人"
{
  "pass": true,
  "score": 25636,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.7%	372	0	0	GET	/api/notifications	3449.329	9.272	15.212	15.216	176
26.6%	337	75	218	PUT	/api/me/profile-image	2434.603	7.224	9.311	9.475	18
21.1%	408	2	24	POST	/api/posts	1929.950	4.730	8.398	8.501	591
7.2%	96	0	50	GET	/api/me/export	656.874	6.842	9.093	9.093	995763
1.2%	3027	243	0	GET	^/api/images/[^/]+	108.860	0.036	0.178	0.651	151898
1.1%	1461	660	0	POST	/api/auth/login	98.671	0.068	0.338	0.474	148
0.9%	522	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	81.915	0.157	0.670	0.991	116
0.8%	2794	0	0	GET	^/api/feeds/[^/]+	69.682	0.025	0.103	0.149	2557
0.6%	934	0	0	GET	/api/memories	53.350	0.057	0.281	0.375	9022
0.6%	237	1	0	POST	^/api/posts/[^/]+/realmojis$	53.331	0.225	1.007	1.283	486
0.5%	702	30	0	POST	/api/friend-requests	44.424	0.063	0.403	1.045	111
0.2%	141	1	0	POST	/api/auth/register	22.394	0.159	0.515	0.519	229
0.2%	468	0	0	GET	^/api/relationships/[^/]+	14.481	0.031	0.129	0.180	1513
0.2%	528	0	0	GET	/api/posting-window/current	14.459	0.027	0.108	0.127	237
0.1%	570	12	0	GET	^/api/posts/[^/]+$	13.598	0.024	0.097	0.123	601
0.1%	120	2	0	POST	^/api/posts/[^/]+/comments$	12.699	0.106	0.589	0.810	344
0.1%	1	0	0	POST	/api/initialize	12.195	12.195	12.195	12.195	90
0.1%	48	0	0	DELETE	^/api/realmojis/[^/]+$	10.543	0.220	0.590	0.590	0
0.1%	362	7	0	GET	^/api/posts/me/current$	10.084	0.028	0.098	0.230	618
0.1%	293	0	0	GET	/api/users	7.297	0.025	0.105	0.135	227
0.1%	292	0	0	GET	^/api/posts/[^/]+/realmojis$	7.136	0.024	0.093	0.143	587
0.1%	158	0	0	POST	/api/auth/logout	6.832	0.043	0.201	0.296	0
0.1%	273	0	0	GET	^/api/users/[^/]+$	6.808	0.025	0.090	0.130	192
0.1%	70	0	0	PUT	^/api/follows/[^/]+$	5.876	0.084	0.709	0.709	0
0.0%	156	0	0	GET	^/api/posts/[^/]+/comments$	3.701	0.024	0.088	0.090	654
0.0%	50	2	0	PATCH	^/api/posts/[^/]+$	2.978	0.060	0.249	0.249	594
0.0%	44	0	0	PATCH	/api/me	2.500	0.057	0.245	0.245	252
0.0%	9	0	0	DELETE	^/api/posts/[^/]+$	2.386	0.265	0.747	0.747	0
0.0%	125	0	0	GET	^/api/memories/[^/]+	2.339	0.019	0.120	0.138	593
0.0%	9	0	0	PUT	^/api/blocks/[^/]+$	2.150	0.239	0.843	0.843	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.8%	85	1	0	GET	/api/notifications	737.403	8.675	12.285	12.285	177
25.2%	81	0	0	POST	/api/posts	433.348	5.350	8.665	8.665	636
10.0%	22	0	11	PUT	/api/me/profile-image	171.825	7.810	8.251	8.251	9
7.2%	17	0	12	GET	/api/me/export	123.222	7.248	9.618	9.618	861202
3.5%	406	3	0	GET	^/api/images/[^/]+	60.943	0.150	0.945	1.171	176603
2.0%	279	0	0	GET	^/api/feeds/[^/]+	34.817	0.125	0.231	0.275	2181
1.6%	153	0	0	GET	/api/memories	26.738	0.175	0.606	0.961	12440
1.5%	164	3	0	POST	/api/friend-requests	25.456	0.155	0.391	0.477	113
1.0%	90	0	0	POST	/api/auth/login	17.515	0.195	0.468	0.468	251
0.9%	112	0	0	GET	/api/posting-window/current	15.065	0.135	0.227	0.266	237
0.7%	61	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	12.336	0.202	0.636	0.636	116
0.5%	71	0	0	GET	^/api/relationships/[^/]+	8.875	0.125	0.203	0.203	1561
0.5%	69	0	0	GET	/api/users	8.616	0.125	0.213	0.213	225
0.5%	72	4	0	GET	^/api/posts/me/current$	8.571	0.119	0.169	0.169	604
0.4%	50	0	0	POST	^/api/posts/[^/]+/comments$	7.479	0.150	0.329	0.329	311
0.4%	50	2	0	GET	^/api/posts/[^/]+$	6.089	0.122	0.179	0.179	578
0.2%	19	0	0	POST	/api/auth/register	4.016	0.211	0.385	0.385	227
0.2%	10	0	0	POST	^/api/posts/[^/]+/realmojis$	2.839	0.284	0.500	0.500	552
0.2%	23	0	0	GET	^/api/posts/[^/]+/comments$	2.783	0.121	0.178	0.178	945
0.1%	16	0	0	POST	/api/auth/logout	2.342	0.146	0.235	0.235	0
0.1%	18	0	0	GET	^/api/users/[^/]+$	2.267	0.126	0.174	0.174	174
0.1%	15	2	0	GET	/api/me	1.922	0.128	0.202	0.202	246
0.1%	11	0	0	PUT	^/api/follows/[^/]+$	1.610	0.146	0.269	0.269	0
0.1%	9	0	0	PATCH	/api/me	1.347	0.150	0.224	0.224	293
0.1%	11	0	0	GET	^/api/memories/[^/]+	1.335	0.121	0.179	0.179	1658
0.1%	7	0	0	DELETE	^/api/comments/[^/]+$	1.000	0.143	0.206	0.206	0
0.0%	5	0	0	PATCH	^/api/posts/[^/]+$	0.750	0.150	0.208	0.208	612
0.0%	3	0	0	DELETE	^/api/posts/[^/]+$	0.676	0.225	0.246	0.246	0
0.0%	5	0	0	GET	^/api/posts/[^/]+/realmojis$	0.606	0.121	0.175	0.175	470
0.0%	1	0	0	DELETE	^/api/realmojis/[^/]+$	0.386	0.386	0.386	0.386	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.1%	77	1	0	GET	/api/notifications	621.477	8.071	9.149	9.149	175
19.7%	43	0	39	PUT	/api/me/profile-image	349.715	8.133	8.326	8.326	15
19.6%	78	1	0	POST	/api/posts	347.081	4.450	8.565	8.565	589
6.6%	15	0	7	GET	/api/me/export	117.197	7.813	10.786	10.786	2394341
4.8%	605	0	0	GET	^/api/images/[^/]+	85.300	0.141	0.928	0.946	147750
2.8%	379	0	0	GET	^/api/feeds/[^/]+	49.283	0.130	0.217	0.223	2631
2.1%	208	0	0	GET	/api/memories	36.583	0.176	0.460	0.612	12173
1.9%	154	0	0	POST	/api/auth/login	34.137	0.222	0.592	0.606	252
1.1%	104	1	0	POST	/api/friend-requests	19.409	0.187	0.779	0.867	114
0.8%	50	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	14.371	0.287	0.835	0.835	116
0.8%	98	0	0	GET	/api/posting-window/current	13.566	0.138	0.339	0.339	240
0.5%	41	1	0	POST	^/api/posts/[^/]+/comments$	9.590	0.234	0.639	0.639	279
0.5%	71	0	0	GET	^/api/relationships/[^/]+	9.426	0.133	0.215	0.215	1210
0.5%	75	4	0	GET	^/api/posts/me/current$	9.423	0.126	0.180	0.180	566
0.5%	64	1	0	GET	^/api/posts/[^/]+$	8.805	0.138	0.210	0.210	601
0.4%	58	0	0	GET	/api/users	7.402	0.128	0.258	0.258	215
0.3%	25	0	0	POST	/api/auth/register	6.120	0.245	0.354	0.354	230
0.3%	46	0	0	GET	^/api/users/[^/]+$	5.884	0.128	0.182	0.182	175
0.3%	31	0	0	POST	/api/auth/logout	5.246	0.169	0.337	0.337	0
0.2%	29	0	0	GET	^/api/posts/[^/]+/comments$	4.071	0.140	0.239	0.239	744
0.2%	11	0	0	POST	^/api/posts/[^/]+/realmojis$	3.588	0.326	0.509	0.509	422
0.2%	22	0	0	GET	^/api/posts/[^/]+/realmojis$	3.008	0.137	0.243	0.243	375
0.2%	22	0	0	GET	^/api/memories/[^/]+	2.979	0.135	0.259	0.259	1082
0.1%	12	0	0	PUT	^/api/follows/[^/]+$	1.692	0.141	0.206	0.206	0
0.1%	7	0	0	PATCH	/api/me	1.229	0.176	0.262	0.262	214
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	1.056	0.352	0.480	0.480	0
0.0%	5	0	0	DELETE	^/api/comments/[^/]+$	0.850	0.170	0.221	0.221	0
0.0%	5	0	0	PATCH	^/api/posts/[^/]+$	0.798	0.160	0.241	0.241	580
0.0%	5	1	0	GET	/api/me	0.544	0.109	0.138	0.138	210
0.0%	1	0	0	DELETE	^/api/posts/[^/]+$	0.404	0.404	0.404	0.404	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16358 クエリ / 80 種類 / 合計 43.899s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	46.8%	59	20.541	348.15	810.52	7812.9	157.0	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	12.3%	856	5.408	6.32	35.76	581.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	12.3%	884	5.381	6.09	206.64	117.6	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	7.5%	1835	3.276	1.79	11.21	12.4	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	5.4%	1	2.392	2391.55	2391.55	29416.0	1199611.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	3.6%	1	1.567	1567.23	1567.23	19993.0	625750.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	2.5%	1	1.079	1079.24	1079.24	11528.0	529700.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	1.7%	540	0.751	1.39	5.44	29.5	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
9	1.4%	257	0.597	2.32	6.80	31.9	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.8%	2	0.330	164.93	180.43	215298.5	2532.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.6%	2	0.263	131.73	255.11	16766.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.6%	208	0.260	1.25	3.51	24.7	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
13	0.5%	1	0.236	236.23	236.23	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
14	0.5%	1	0.208	208.18	208.18	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
15	0.4%	1	0.180	180.27	180.27	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: e96aad80b391ccef1c9272d4caa2ab7d86cec3b2
Type: cpu
Time: 2026-10-31 11:41:11 JST
Duration: 60s, Total samples = 74.75s (124.58%)
Showing nodes accounting for 65.37s, 87.45% of 74.75s total
Dropped 837 nodes (cum <= 0.37s)
Showing top 40 nodes out of 168
      flat  flat%   sum%        cum   cum%
    13.70s 18.33% 18.33%     13.70s 18.33%  crypto/internal/fips140/sha256.blockAVX2
     9.44s 12.63% 30.96%      9.44s 12.63%  runtime.spanClass.sizeclass (inline)
     7.28s  9.74% 40.70%     11.45s 15.32%  image/jpeg.(*decoder).reconstructBlock
     3.40s  4.55% 45.24%     13.43s 17.97%  runtime.tryDeferToSpanScan
     3.15s  4.21% 49.46%     23.27s 31.13%  image/jpeg.(*decoder).processSOS
     2.85s  3.81% 53.27%      4.28s  5.73%  image/jpeg.(*decoder).decodeHuffman
     2.16s  2.89% 56.16%      2.26s  3.02%  image/jpeg.(*decoder).receiveExtend
     2.15s  2.88% 59.04%      2.15s  2.88%  runtime.memmove
     1.83s  2.45% 61.48%      2.36s  3.16%  image/jpeg.idctCols
     1.78s  2.38% 63.87%      1.78s  2.38%  runtime.memclrNoHeapPointers
     1.57s  2.10% 65.97%      6.33s  8.47%  runtime.scanObjectsSmall
     1.39s  1.86% 67.83%     10.42s 13.94%  runtime.scanObject
     1.33s  1.78% 69.61%      5.45s  7.29%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.32s  1.77% 71.37%      1.74s  2.33%  image/jpeg.idctRows
     1.32s  1.77% 73.14%      1.32s  1.77%  internal/runtime/syscall/linux.Syscall6
     1.19s  1.59% 74.73%      1.19s  1.59%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.09s  1.46% 76.19%      1.50s  2.01%  image/jpeg.(*decoder).ensureNBits
     0.91s  1.22% 77.40%      0.92s  1.23%  image/jpeg.dctBox (inline)
     0.88s  1.18% 78.58%     15.70s 21.00%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.73s  0.98% 79.56%      0.73s  0.98%  aeshashbody
     0.73s  0.98% 80.54%     15.16s 20.28%  crypto/internal/fips140/sha256.(*Digest).Write
     0.64s  0.86% 81.39%      0.64s  0.86%  memeqbody
     0.54s  0.72% 82.11%     16.47s 22.03%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.43s  0.58% 82.69%      7.26s  9.71%  runtime.scanSpan
     0.39s  0.52% 83.21%      0.52s   0.7%  example.com/isureal-company/isureal.postLess
     0.39s  0.52% 83.73%      0.41s  0.55%  image/jpeg.(*decoder).readByteStuffedByte
     0.34s  0.45% 84.19%      0.83s  1.11%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.32s  0.43% 84.62%      1.91s  2.56%  runtime.mapaccess1_faststr
     0.29s  0.39% 85.00%     17.35s 23.21%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.29s  0.39% 85.39%     18.83s 25.19%  example.com/isureal-company/isureal.imageSignature
     0.24s  0.32% 85.71%      0.73s  0.98%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.22s  0.29% 86.01%      0.88s  1.18%  runtime.mallocgcSmallScanNoHeader
     0.21s  0.28% 86.29%     13.91s 18.61%  crypto/internal/fips140/sha256.block
     0.19s  0.25% 86.54%      0.44s  0.59%  database/sql.convertAssignRows
     0.15s   0.2% 86.74%      8.28s 11.08%  runtime.mallocgc
     0.14s  0.19% 86.93%      1.24s  1.66%  runtime.mapassign_faststr
     0.13s  0.17% 87.10%      2.58s  3.45%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.09s  0.12% 87.22%      1.28s  1.71%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.09s  0.12% 87.34%      0.42s  0.56%  runtime.sweepone
     0.08s  0.11% 87.45%     11.73s 15.69%  runtime.gcDrain
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 72.6% / 100% | 57% | 0.8% | 6 | 147 MB | CPU が飽和 |
| isu2 (2 cpu) | 0.9% / 3% | 0% | 0.1% | 3 | 502 MB |  |
| isu3 (2 cpu) | 1.6% / 5% | 0% | 0.3% | 1 | 231 MB |  |
| isu4 (2 cpu) | 24.4% / 54% | 0% | 14.7% | 7 | 103 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 125% / nginx 13% |
| isu2 | 目立つプロセス無し |
| isu3 | isureal 1% |
| isu4 | postgres 17% |

## Diff from previous run

スコア: 17072 → 25636 (+8564)

`6a7a33a..8b52957`

```
 webapp/go/limiter.go | 37 +++++++++++++++++++++++++++++++++++++
 webapp/go/server.go  | 10 +++++-----
 2 files changed, 42 insertions(+), 5 deletions(-)
```
