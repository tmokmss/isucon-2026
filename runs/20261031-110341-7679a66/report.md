# run 20261031-110341-7679a66

## Summary

| | |
|---|---|
| score | 16152 |
| status | pass |
| ref | HEAD |
| sha | 7679a66b952d |
| 所要時間 | 109s |
| 前回比 | +621 (4.0%) vs 20261031-105807-4478bd9 |
| ベスト比 | -8649 (-34.9%) vs 20261031-103811-cef5fe3 = 24801 |
| note | 画像: PostgreSQL から取った初期データの画像を init/ に貯める |

## Bench

```
{"id":"2526","status":"done","portalStatus":"finished","passed":true,"score":16152,"raw":16152,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:04:18.358Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:04:21.099Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:04:23.684Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:04:27.705Z level=INFO msg="asia 2 日目: 40 人"
time=2026-10-31T02:04:29.878Z level=INFO msg="europe 2 日目: 12 人"
time=2026-10-31T02:04:33.371Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T02:04:38.036Z level=INFO msg="asia 3 日目: 96 人"
time=2026-10-31T02:04:39.351Z level=INFO msg="europe 3 日目: 20 人"
time=2026-10-31T02:04:43.114Z level=INFO msg="america 3 日目: 23 人"
time=2026-10-31T02:04:46.619Z level=INFO msg="asia 4 日目: 96 人"
time=2026-10-31T02:04:48.823Z level=INFO msg="europe 4 日目: 20 人"
time=2026-10-31T02:04:52.262Z level=INFO msg="america 4 日目: 23 人"
time=2026-10-31T02:04:57.519Z level=INFO msg="asia 5 日目: 96 人"
time=2026-10-31T02:05:00.976Z level=INFO msg="europe 5 日目: 20 人"
time=2026-10-31T02:05:05.212Z level=INFO msg="america 5 日目: 23 人"
time=2026-10-31T02:05:12.878Z level=INFO msg="america 最終日（グランプリ）: 23 人"
time=2026-10-31T02:05:16.126Z level=INFO msg="asia 最終日（グランプリ）: 96 人"
time=2026-10-31T02:05:17.671Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 1 人"
{
  "pass": true,
  "score": 16152,
  "messages": [
    "ulzoe_5684: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.9%	326	0	0	GET	/api/notifications	3745.811	11.490	21.602	22.533	177
11.3%	289	2	0	POST	/api/posts	984.058	3.405	12.647	13.934	633
8.9%	2914	614	0	GET	^/api/images/[^/]+	773.964	0.266	2.307	6.253	99411
7.3%	1680	965	0	POST	/api/auth/login	639.285	0.381	3.103	10.649	121
6.1%	576	23	0	POST	/api/friend-requests	536.403	0.931	8.613	16.354	112
5.6%	257	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	492.501	1.916	8.578	12.541	116
4.1%	195	149	0	PUT	/api/me/profile-image	355.803	1.825	8.965	9.106	20
3.2%	1614	0	0	GET	^/api/feeds/[^/]+	275.602	0.171	1.662	1.981	1100
1.6%	752	0	0	GET	/api/memories	138.779	0.185	1.198	3.889	10455
1.5%	112	1	0	POST	/api/auth/register	132.210	1.180	5.315	5.375	229
1.1%	402	0	0	GET	/api/posting-window/current	93.271	0.232	1.396	2.418	237
1.0%	57	0	0	GET	/api/me/export	89.400	1.568	8.541	8.541	2055347
0.7%	269	15	0	GET	^/api/posts/me/current$	59.879	0.223	1.747	1.955	602
0.6%	69	0	0	PUT	^/api/follows/[^/]+$	48.077	0.697	5.696	5.696	0
0.5%	18	0	0	POST	/api/admin/posting-window	47.351	2.631	9.824	9.824	177
0.5%	308	0	0	GET	^/api/relationships/[^/]+	42.886	0.139	0.812	1.477	1582
0.4%	262	0	0	GET	/api/users	35.219	0.134	1.236	1.769	229
0.4%	39	4	0	PATCH	^/api/posts/[^/]+$	33.733	0.865	3.810	3.810	561
0.4%	123	0	0	POST	/api/auth/logout	32.685	0.266	2.246	5.174	0
0.4%	27	1	0	POST	^/api/posts/[^/]+/comments$	31.409	1.163	6.999	6.999	286
0.3%	197	13	0	GET	^/api/posts/[^/]+$	28.487	0.145	1.015	1.737	590
0.3%	41	0	0	POST	^/api/posts/[^/]+/realmojis$	26.268	0.641	5.775	5.775	484
0.2%	11	0	0	DELETE	^/api/posts/[^/]+$	20.272	1.843	5.504	5.504	0
0.2%	41	0	0	PATCH	/api/me	15.215	0.371	2.678	2.678	252
0.1%	92	0	0	GET	^/api/users/[^/]+$	10.811	0.118	1.825	1.825	210
0.1%	1	0	0	POST	/api/initialize	9.873	9.873	9.873	9.873	90
0.1%	4	0	0	DELETE	^/api/follows/[^/]+$	7.030	1.757	3.112	3.112	0
0.1%	8	0	0	PUT	^/api/admin/users/[^/]+/certification$	6.923	0.865	6.282	6.282	0
0.1%	112	0	0	GET	^/api/memories/[^/]+	6.275	0.056	0.372	1.104	615
0.0%	53	3	0	GET	/api/me	4.244	0.080	0.375	0.375	282
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.7%	77	0	0	GET	/api/notifications	838.993	10.896	20.104	20.104	181
13.7%	66	0	0	POST	/api/posts	251.866	3.816	10.545	10.545	649
5.9%	96	3	0	POST	/api/friend-requests	107.786	1.123	8.055	8.055	112
5.5%	53	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	100.387	1.894	9.214	9.214	116
5.2%	367	0	0	GET	^/api/feeds/[^/]+	95.847	0.261	1.067	1.706	1187
5.1%	104	0	0	POST	/api/auth/login	93.021	0.894	5.195	5.950	250
3.4%	412	3	0	GET	^/api/images/[^/]+	63.174	0.153	0.939	0.944	174213
2.9%	147	0	0	GET	/api/memories	53.333	0.363	1.919	3.157	8765
1.8%	81	0	0	GET	/api/posting-window/current	33.239	0.410	2.485	2.485	239
1.6%	18	0	0	POST	/api/auth/register	30.072	1.671	5.512	5.512	228
1.3%	64	4	0	GET	^/api/posts/me/current$	23.819	0.372	1.506	1.506	605
0.9%	64	0	0	GET	^/api/relationships/[^/]+	17.290	0.270	2.099	2.099	1432
0.8%	20	0	0	POST	/api/auth/logout	15.444	0.772	2.817	2.817	0
0.7%	19	0	0	PUT	^/api/follows/[^/]+$	13.680	0.720	3.383	3.383	0
0.7%	8	0	0	GET	/api/me/export	12.018	1.502	3.539	3.539	1478409
0.6%	43	0	0	GET	/api/users	11.675	0.272	1.552	1.552	243
0.5%	8	1	0	PATCH	^/api/posts/[^/]+$	9.510	1.189	2.762	2.762	566
0.5%	8	0	0	POST	^/api/posts/[^/]+/realmojis$	9.322	1.165	4.053	4.053	551
0.5%	10	0	0	PUT	/api/me/profile-image	8.963	0.896	1.969	1.969	0
0.5%	43	4	0	GET	^/api/posts/[^/]+$	8.524	0.198	0.896	0.896	552
0.4%	5	0	0	DELETE	^/api/posts/[^/]+$	7.874	1.575	4.780	4.780	0
0.3%	11	0	0	PATCH	/api/me	5.976	0.543	1.137	1.137	299
0.3%	31	0	0	GET	^/api/users/[^/]+$	5.383	0.174	0.473	0.473	180
0.3%	9	0	0	POST	^/api/posts/[^/]+/comments$	5.040	0.560	3.483	3.483	324
0.2%	15	0	0	GET	^/api/memories/[^/]+	3.798	0.253	0.476	0.476	563
0.2%	14	1	0	GET	/api/me	2.871	0.205	0.413	0.413	261
0.1%	16	0	0	GET	^/api/posts/[^/]+/comments$	2.495	0.156	0.468	0.468	223
0.1%	12	0	0	GET	^/api/posts/[^/]+/realmojis$	1.901	0.158	0.265	0.265	331
0.0%	2	0	0	DELETE	^/api/realmojis/[^/]+$	0.492	0.246	0.311	0.311	0
0.0%	1	0	0	DELETE	^/api/comments/[^/]+$	0.290	0.290	0.290	0.290	0
…（上位 30 / 全 31 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.1%	81	0	0	GET	/api/notifications	712.362	8.795	12.834	12.834	178
17.4%	74	1	0	POST	/api/posts	309.522	4.183	13.517	13.517	611
8.0%	119	5	0	POST	/api/friend-requests	141.712	1.191	10.005	14.045	112
5.4%	42	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	95.968	2.285	10.685	10.685	116
4.7%	97	0	0	POST	/api/auth/login	82.773	0.853	4.854	4.854	250
3.4%	367	2	0	GET	^/api/images/[^/]+	60.959	0.166	0.947	1.243	153778
3.3%	192	0	0	GET	^/api/feeds/[^/]+	58.555	0.305	1.639	1.705	885
2.9%	144	0	0	GET	/api/memories	51.938	0.361	1.644	2.035	12958
2.5%	8	0	0	GET	/api/me/export	43.543	5.443	16.140	16.140	2927914
2.4%	22	0	0	POST	/api/auth/register	43.407	1.973	5.112	5.112	227
2.0%	95	0	0	GET	/api/posting-window/current	35.655	0.375	1.528	1.528	240
1.5%	14	0	0	PUT	^/api/follows/[^/]+$	26.158	1.868	3.790	3.790	0
1.4%	59	16	0	GET	^/api/posts/me/current$	24.308	0.412	1.537	1.537	456
1.0%	56	0	0	GET	/api/users	17.587	0.314	1.535	1.535	258
0.9%	57	0	0	GET	^/api/relationships/[^/]+	15.676	0.275	2.008	2.008	1268
0.7%	9	0	0	PUT	/api/me/profile-image	11.815	1.313	2.381	2.381	0
0.4%	12	0	0	POST	^/api/posts/[^/]+/realmojis$	6.793	0.566	2.202	2.202	404
0.4%	5	1	0	PATCH	^/api/posts/[^/]+$	6.463	1.293	5.019	5.019	487
0.3%	12	0	0	POST	/api/auth/logout	5.156	0.430	1.444	1.444	0
0.2%	6	0	0	PATCH	/api/me	4.334	0.722	1.685	1.685	234
0.2%	21	1	0	GET	^/api/posts/[^/]+$	4.167	0.198	0.356	0.356	576
0.2%	8	0	0	POST	^/api/posts/[^/]+/comments$	3.518	0.440	2.078	2.078	270
0.2%	3	0	0	DELETE	^/api/posts/[^/]+$	3.269	1.090	2.092	2.092	0
0.1%	12	0	0	GET	/api/me	2.411	0.201	0.417	0.417	277
0.1%	11	0	0	GET	^/api/memories/[^/]+	2.264	0.206	0.372	0.372	787
0.1%	9	0	0	GET	^/api/posts/[^/]+/comments$	1.731	0.192	0.324	0.324	329
0.1%	4	0	0	GET	^/api/users/[^/]+$	1.510	0.378	0.583	0.583	216
0.1%	4	0	0	GET	^/api/posts/[^/]+/realmojis$	1.292	0.323	0.919	0.919	470
0.1%	2	0	0	DELETE	^/api/comments/[^/]+$	1.269	0.635	1.165	1.165	0
0.0%	1	0	0	DELETE	^/api/realmojis/[^/]+$	0.144	0.144	0.144	0.144	0
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 12564 クエリ / 77 種類 / 合計 61.661s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	56.5%	73	34.825	477.06	3603.56	6264.6	133.4	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	11.9%	18	7.362	409.00	2199.43	32745.7	1658.5	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
3	8.9%	667	5.498	8.24	227.05	120.2	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	7.7%	552	4.763	8.63	39.01	721.7	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
5	5.0%	1686	3.105	1.84	7.66	12.5	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
6	3.7%	1	2.269	2268.94	2268.94	29314.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	1.1%	1	0.698	697.94	697.94	11528.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	1.1%	426	0.676	1.59	4.27	33.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
9	0.5%	2	0.308	154.19	160.64	210567.5	2491.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
10	0.5%	1	0.288	287.78	287.78	40112.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.4%	1	0.230	229.63	229.63	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
12	0.3%	61	0.180	2.95	5.12	50.7	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
13	0.3%	1	0.180	179.68	179.68	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.3%	746	0.156	0.21	2.29	23.0	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
15	0.2%	1	0.149	148.79	148.79	29315.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4;
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: f78333ec298668bba832efe02dfede12b95a01fa
Type: cpu
Time: 2026-10-31 11:04:02 JST
Duration: 60.79s, Total samples = 87.24s (143.51%)
Showing nodes accounting for 77.85s, 89.24% of 87.24s total
Dropped 818 nodes (cum <= 0.44s)
Showing top 40 nodes out of 175
      flat  flat%   sum%        cum   cum%
    15.73s 18.03% 18.03%     15.74s 18.04%  crypto/internal/fips140/sha256.blockAVX2
    13.89s 15.92% 33.95%     13.89s 15.92%  runtime.spanClass.sizeclass (inline)
     7.48s  8.57% 42.53%     12.09s 13.86%  image/jpeg.(*decoder).reconstructBlock
     4.70s  5.39% 47.91%     19.41s 22.25%  runtime.tryDeferToSpanScan
     3.46s  3.97% 51.88%     26.05s 29.86%  image/jpeg.(*decoder).processSOS
     2.89s  3.31% 55.19%      2.89s  3.31%  runtime.memclrNoHeapPointers
     2.82s  3.23% 58.43%      4.65s  5.33%  image/jpeg.(*decoder).decodeHuffman
     2.39s  2.74% 61.16%      2.39s  2.74%  runtime.memmove
     2.35s  2.69% 63.86%      2.48s  2.84%  image/jpeg.(*decoder).receiveExtend
     2.30s  2.64% 66.49%      9.35s 10.72%  runtime.scanObjectsSmall
     2.03s  2.33% 68.82%      2.51s  2.88%  image/jpeg.idctCols
     1.80s  2.06% 70.88%     15.07s 17.27%  runtime.scanObject
     1.50s  1.72% 72.60%      2.02s  2.32%  image/jpeg.idctRows
     1.34s  1.54% 74.14%      1.34s  1.54%  internal/runtime/syscall/linux.Syscall6
     1.26s  1.44% 75.58%      1.95s  2.24%  image/jpeg.(*decoder).ensureNBits
     1.17s  1.34% 76.93%      5.43s  6.22%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.08s  1.24% 78.16%      1.08s  1.24%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.98s  1.12% 79.29%      0.98s  1.12%  image/jpeg.dctBox (inline)
     0.92s  1.05% 80.34%     17.70s 20.29%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.78s  0.89% 81.24%     17.18s 19.69%  crypto/internal/fips140/sha256.(*Digest).Write
     0.78s  0.89% 82.13%     11.13s 12.76%  runtime.scanSpan
     0.75s  0.86% 82.99%      0.75s  0.86%  aeshashbody
     0.65s  0.75% 83.73%      0.65s  0.75%  memeqbody
     0.64s  0.73% 84.47%      0.69s  0.79%  image/jpeg.(*decoder).readByteStuffedByte
     0.63s  0.72% 85.19%      0.63s  0.72%  runtime.extractHeapBitsSmall
     0.51s  0.58% 85.77%     19.65s 22.52%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.47s  0.54% 86.31%     18.47s 21.17%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.38s  0.44% 86.75%      2.13s  2.44%  runtime.mapaccess1_faststr
     0.34s  0.39% 87.14%      0.57s  0.65%  example.com/isureal-company/isureal.postLess
     0.26s   0.3% 87.44%     20.91s 23.97%  example.com/isureal-company/isureal.imageSignature
     0.26s   0.3% 87.73%      0.93s  1.07%  runtime.scanObjectSmall
     0.22s  0.25% 87.99%      0.63s  0.72%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.21s  0.24% 88.23%      0.77s  0.88%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.20s  0.23% 88.46%     15.94s 18.27%  crypto/internal/fips140/sha256.block
     0.20s  0.23% 88.69%      1.08s  1.24%  runtime.mapassign_faststr
     0.12s  0.14% 88.82%     16.94s 19.42%  runtime.mallocgc
     0.11s  0.13% 88.95%      0.79s  0.91%  runtime.mallocgcSmallScanNoHeader
     0.09s   0.1% 89.05%      0.53s  0.61%  runtime.sweepone
     0.08s 0.092% 89.14%      0.45s  0.52%  database/sql.convertAssignRows
     0.08s 0.092% 89.24%      1.01s  1.16%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 78.2% / 100% | 67% | 0.0% | 4 | 122 MB | CPU が飽和 |
| isu2 (2 cpu) | 0.9% / 3% | 0% | 0.1% | 1 | 289 MB |  |
| isu3 (2 cpu) | 1.2% / 4% | 0% | 0.1% | 3 | 221 MB |  |
| isu4 (2 cpu) | 27.8% / 78% | 0% | 17.4% | 5 | 92 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 140% / nginx 12% |
| isu2 | 目立つプロセス無し |
| isu3 | isureal 1% |
| isu4 | postgres 17% |

## Diff from previous run

スコア: 15531 → 16152 (+621)

`4478bd9..7679a66`

```
 webapp/go/image_local.go | 37 +++++++++++++++++++++++++++++--------
 1 file changed, 29 insertions(+), 8 deletions(-)
```
