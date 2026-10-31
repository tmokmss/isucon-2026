# run 20261031-105006-a4d727c

## Summary

| | |
|---|---|
| score | 16943 |
| status | pass |
| ref | HEAD |
| sha | a4d727c84688 |
| 所要時間 | 277s |
| 前回比 | -3282 (-16.2%) vs 20261031-104641-d3eb91b |
| ベスト比 | -7858 (-31.7%) vs 20261031-103811-cef5fe3 = 24801 |
| note | TLS無効+export主キー+ID共有+GOMEMLIMIT2200 |

## Bench

```
{"id":"1898","status":"done","portalStatus":"finished","passed":true,"score":16943,"raw":16943,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:53:21.839Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:53:24.246Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:53:26.955Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:53:30.765Z level=INFO msg="asia 2 日目: 34 人"
time=2026-10-31T01:53:33.343Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T01:53:36.127Z level=INFO msg="america 2 日目: 7 人"
time=2026-10-31T01:53:39.573Z level=INFO msg="asia 3 日目: 70 人"
time=2026-10-31T01:53:42.336Z level=INFO msg="europe 3 日目: 11 人"
time=2026-10-31T01:53:45.942Z level=INFO msg="america 3 日目: 12 人"
time=2026-10-31T01:53:50.086Z level=INFO msg="asia 4 日目: 122 人"
time=2026-10-31T01:53:55.070Z level=INFO msg="europe 4 日目: 13 人"
time=2026-10-31T01:53:56.427Z level=INFO msg="america 4 日目: 12 人"
time=2026-10-31T01:53:59.082Z level=INFO msg="asia 5 日目: 122 人"
time=2026-10-31T01:54:04.566Z level=INFO msg="europe 5 日目: 13 人"
time=2026-10-31T01:54:06.714Z level=INFO msg="america 5 日目: 12 人"
time=2026-10-31T01:54:13.836Z level=INFO msg="asia 最終日（グランプリ）: 121 人"
time=2026-10-31T01:54:14.237Z level=INFO msg="europe 最終日（グランプリ）: 13 人"
time=2026-10-31T01:54:15.983Z level=INFO msg="america 最終日（グランプリ）: 12 人"
time=2026-10-31T01:54:20.935Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 16943,
  "messages": [
    "nuivymia.5947: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.8%	332	0	0	GET	/api/notifications	3177.007	9.569	16.010	16.889	175
17.7%	379	33	0	POST	/api/posts	1618.297	4.270	11.271	11.850	585
11.2%	2064	1146	0	POST	/api/auth/login	1025.048	0.497	3.483	6.873	124
7.1%	3459	442	0	GET	^/api/images/[^/]+	651.453	0.188	1.909	5.519	104144
5.4%	209	148	0	PUT	/api/me/profile-image	493.555	2.362	8.644	9.164	12
4.6%	138	1	0	POST	/api/auth/register	417.082	3.022	14.983	15.170	229
4.3%	663	21	0	POST	/api/friend-requests	392.524	0.592	5.577	6.619	112
3.4%	294	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	307.474	1.046	8.886	9.677	115
2.4%	2062	0	0	GET	^/api/feeds/[^/]+	218.888	0.106	0.588	1.550	2069
1.6%	918	0	0	GET	/api/memories	149.567	0.163	1.221	5.225	11476
1.3%	489	0	0	GET	/api/posting-window/current	117.861	0.241	1.145	1.484	237
1.1%	71	0	0	GET	/api/me/export	103.718	1.461	12.604	12.604	1865720
0.6%	35	4	0	POST	^/api/posts/[^/]+/realmojis$	58.839	1.681	6.599	6.599	442
0.6%	71	0	0	PUT	^/api/follows/[^/]+$	55.730	0.785	8.610	8.610	0
0.4%	275	8	0	GET	^/api/posts/me/current$	37.950	0.138	0.640	1.179	624
0.4%	357	0	0	GET	^/api/relationships/[^/]+	37.636	0.105	0.430	0.571	1538
0.4%	18	0	0	POST	/api/admin/posting-window	32.313	1.795	4.633	4.633	177
0.4%	304	0	0	GET	/api/users	31.966	0.105	0.588	0.670	243
0.3%	353	7	0	GET	^/api/posts/[^/]+$	31.139	0.088	0.309	0.454	601
0.3%	41	1	0	POST	^/api/posts/[^/]+/comments$	29.534	0.720	6.406	6.406	324
0.2%	45	0	0	PATCH	/api/me	22.810	0.507	3.374	3.374	274
0.2%	131	0	0	POST	/api/auth/logout	19.170	0.146	1.009	1.535	0
0.2%	182	0	0	GET	^/api/users/[^/]+$	15.189	0.083	0.298	0.339	191
0.1%	1	0	0	POST	/api/initialize	12.898	12.898	12.898	12.898	90
0.1%	29	2	0	PATCH	^/api/posts/[^/]+$	12.488	0.431	0.938	0.938	601
0.1%	112	0	0	GET	^/api/posts/[^/]+/realmojis$	8.377	0.075	0.343	0.369	173
0.1%	147	0	0	GET	^/api/memories/[^/]+	7.960	0.054	0.593	0.691	716
0.1%	14	0	0	DELETE	^/api/posts/[^/]+$	6.807	0.486	0.989	0.989	0
0.1%	7	0	0	PUT	^/api/blocks/[^/]+$	6.303	0.900	2.776	2.776	0
0.1%	69	3	0	GET	/api/me	6.196	0.090	0.631	0.631	284
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
57.4%	63	0	0	GET	/api/notifications	858.659	13.630	51.314	51.314	155
9.8%	46	1	0	POST	/api/posts	147.026	3.196	10.595	10.595	604
7.1%	247	1	0	GET	^/api/images/[^/]+	106.455	0.431	2.204	3.227	173147
5.1%	74	3	0	POST	/api/friend-requests	76.584	1.035	5.057	5.057	112
3.5%	40	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	52.649	1.316	6.703	6.703	116
3.0%	65	0	0	POST	/api/auth/login	44.279	0.681	6.137	6.137	253
2.9%	173	0	0	GET	^/api/feeds/[^/]+	44.086	0.255	0.824	0.840	1274
2.1%	11	0	0	POST	/api/auth/register	30.798	2.800	13.952	13.952	227
1.8%	98	0	0	GET	/api/memories	26.999	0.275	2.263	2.263	6842
1.5%	6	0	0	GET	/api/me/export	22.532	3.755	7.435	7.435	1737428
1.2%	51	0	0	GET	/api/posting-window/current	18.503	0.363	1.380	1.380	239
0.8%	37	0	0	GET	/api/users	11.730	0.317	1.407	1.407	238
0.7%	41	0	0	GET	^/api/relationships/[^/]+	9.897	0.241	0.740	0.740	1519
0.6%	38	0	0	GET	^/api/posts/me/current$	9.166	0.241	0.641	0.641	620
0.5%	1	0	0	POST	/api/me/account-deletion	7.844	7.844	7.844	7.844	0
0.5%	9	0	0	PUT	^/api/follows/[^/]+$	7.528	0.836	1.333	1.333	0
0.4%	28	4	0	GET	^/api/posts/[^/]+$	5.680	0.203	0.470	0.470	518
0.3%	8	0	0	POST	/api/auth/logout	3.913	0.489	1.376	1.376	0
0.2%	4	0	0	PATCH	^/api/posts/[^/]+$	3.169	0.792	2.236	2.236	637
0.1%	8	0	0	GET	^/api/users/[^/]+$	2.091	0.261	0.438	0.438	171
0.1%	9	0	0	GET	^/api/memories/[^/]+	1.576	0.175	0.293	0.293	780
0.1%	2	0	0	PUT	/api/me/profile-image	1.275	0.637	0.801	0.801	0
0.1%	6	1	0	GET	/api/me	0.910	0.152	0.231	0.231	202
0.0%	2	0	0	PATCH	/api/me	0.509	0.255	0.285	0.285	275
0.0%	3	0	0	GET	^/api/posts/[^/]+/comments$	0.360	0.120	0.155	0.155	29
0.0%	1	0	0	GET	^/api/posts/[^/]+/realmojis$	0.358	0.358	0.358	0.358	29
0.0%	1	0	0	DELETE	^/api/posts/[^/]+$	0.333	0.333	0.333	0.333	0
0.0%	2	0	0	PUT	^/api/blocks/[^/]+$	0.230	0.115	0.124	0.124	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.2%	54	0	0	GET	/api/notifications	476.906	8.832	10.597	10.597	179
13.1%	47	1	0	POST	/api/posts	141.185	3.004	8.249	8.249	605
8.2%	278	0	0	GET	^/api/images/[^/]+	88.781	0.319	1.613	2.659	186997
6.5%	47	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	70.662	1.503	9.654	9.654	116
6.4%	72	4	0	POST	/api/friend-requests	69.469	0.965	5.672	5.672	110
4.5%	68	0	0	POST	/api/auth/login	48.868	0.719	6.516	6.516	248
3.3%	135	0	0	GET	^/api/feeds/[^/]+	35.172	0.261	1.053	1.112	1479
2.6%	97	0	0	GET	/api/memories	27.735	0.286	1.790	1.790	9161
2.3%	59	0	0	GET	/api/posting-window/current	24.500	0.415	2.780	2.780	240
1.8%	6	0	0	GET	/api/me/export	19.303	3.217	5.784	5.784	3010437
1.1%	15	0	0	POST	^/api/posts/[^/]+/realmojis$	11.614	0.774	4.077	4.077	408
0.9%	42	3	0	GET	^/api/posts/me/current$	10.085	0.240	0.774	0.774	576
0.8%	42	0	0	GET	^/api/relationships/[^/]+	8.838	0.210	0.814	0.814	1774
0.8%	35	0	0	GET	/api/users	8.345	0.238	0.696	0.696	247
0.7%	10	0	0	PUT	^/api/follows/[^/]+$	7.048	0.705	3.499	3.499	0
0.5%	4	0	0	PATCH	^/api/posts/[^/]+$	5.171	1.293	3.384	3.384	645
0.4%	17	3	0	GET	^/api/posts/[^/]+$	4.686	0.276	0.600	0.600	480
0.4%	16	0	0	POST	/api/auth/logout	3.982	0.249	1.259	1.259	0
0.4%	11	0	0	POST	/api/auth/register	3.973	0.361	0.633	0.633	230
0.3%	13	0	0	GET	^/api/users/[^/]+$	3.567	0.274	1.027	1.027	169
0.2%	5	0	0	PUT	/api/me/profile-image	2.185	0.437	0.606	0.606	0
0.2%	10	0	0	GET	^/api/posts/[^/]+/realmojis$	1.706	0.171	0.335	0.335	396
0.1%	9	0	0	GET	^/api/memories/[^/]+	1.467	0.163	0.441	0.441	722
0.1%	3	0	0	PATCH	/api/me	1.391	0.464	0.793	0.793	296
0.1%	8	0	0	GET	/api/me	1.114	0.139	0.235	0.235	266
0.1%	1	0	0	DELETE	^/api/realmojis/[^/]+$	1.105	1.105	1.105	1.105	0
0.0%	1	0	0	GET	^/api/posts/[^/]+/comments$	0.151	0.151	0.151	0.151	29
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.106	0.106	0.106	0.106	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 12635 クエリ / 79 種類 / 合計 57.019s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	57.3%	83	32.673	393.65	2748.48	1265.2	124.6	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	10.7%	555	6.102	10.99	129.81	750.6	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	7.1%	2591	4.062	1.57	6.22	11.8	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	6.9%	15	3.936	262.41	887.28	32586.7	1649.4	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
5	4.1%	1	2.311	2311.36	2311.36	29288.0	1199439.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	2.8%	1	1.617	1616.96	1616.96	19464.0	625585.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	1.8%	1	1.021	1021.23	1021.23	11528.0	529653.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	1.4%	441	0.789	1.79	6.53	33.7	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
9	1.4%	2	0.770	385.02	414.35	214406.0	2522.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
10	1.3%	1	0.725	724.76	724.76	20413.0	474.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	1.1%	3	0.625	208.24	606.93	33441.7	1679.3	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
12	0.5%	2	0.281	140.45	273.26	20231.5	2517.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
13	0.4%	1	0.230	230.43	230.43	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
14	0.4%	1	0.221	221.29	221.29	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
15	0.4%	1	0.212	211.61	211.61	70449.0	29793.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 3
SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: a5fa48c1ada79d6f165e6873097af3aa1d5eef46
Type: cpu
Time: 2026-10-31 10:53:03 JST
Duration: 60.56s, Total samples = 86.28s (142.47%)
Showing nodes accounting for 76.41s, 88.56% of 86.28s total
Dropped 774 nodes (cum <= 0.43s)
Showing top 40 nodes out of 178
      flat  flat%   sum%        cum   cum%
    16.51s 19.14% 19.14%     16.51s 19.14%  crypto/internal/fips140/sha256.blockAVX2
    14.03s 16.26% 35.40%     14.03s 16.26%  runtime.spanClass.sizeclass (inline)
     6.41s  7.43% 42.83%     10.38s 12.03%  image/jpeg.(*decoder).reconstructBlock
     4.77s  5.53% 48.35%     19.63s 22.75%  runtime.tryDeferToSpanScan
     3.10s  3.59% 51.95%     23.69s 27.46%  image/jpeg.(*decoder).processSOS
     2.80s  3.25% 55.19%      2.80s  3.25%  runtime.memmove
     2.43s  2.82% 58.01%      4.19s  4.86%  image/jpeg.(*decoder).decodeHuffman
     2.27s  2.63% 60.64%      2.27s  2.63%  runtime.memclrNoHeapPointers
        2s  2.32% 62.96%      2.07s  2.40%  image/jpeg.(*decoder).receiveExtend
     1.96s  2.27% 65.23%      8.54s  9.90%  runtime.scanObjectsSmall
     1.77s  2.05% 67.28%     15.40s 17.85%  runtime.scanObject
     1.62s  1.88% 69.16%      1.62s  1.88%  internal/runtime/syscall/linux.Syscall6
     1.58s  1.83% 70.99%      1.98s  2.29%  image/jpeg.idctCols
     1.49s  1.73% 72.72%      5.71s  6.62%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.49s  1.73% 74.44%      1.88s  2.18%  image/jpeg.idctRows
     1.29s  1.50% 75.94%      1.82s  2.11%  image/jpeg.(*decoder).ensureNBits
     1.16s  1.34% 77.28%      1.16s  1.34%  aeshashbody
     0.98s  1.14% 78.42%      0.98s  1.14%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.91s  1.05% 79.47%     18.36s 21.28%  crypto/internal/fips140/sha256.(*Digest).Write
     0.84s  0.97% 80.45%     18.54s 21.49%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.79s  0.92% 81.36%      0.79s  0.92%  image/jpeg.dctBox (inline)
     0.70s  0.81% 82.17%     10.24s 11.87%  runtime.scanSpan
     0.53s  0.61% 82.79%      2.03s  2.35%  runtime.mapaccess1_faststr
     0.52s   0.6% 83.39%     19.42s 22.51%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.51s  0.59% 83.98%      0.53s  0.61%  image/jpeg.(*decoder).readByteStuffedByte
     0.45s  0.52% 84.50%     20.79s 24.10%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.45s  0.52% 85.03%      0.45s  0.52%  runtime.extractHeapBitsSmall
     0.42s  0.49% 85.51%      0.46s  0.53%  runtime.(*spanInlineMarkBits).tryAcquire
     0.35s  0.41% 85.92%      0.56s  0.65%  example.com/isureal-company/isureal.postLess
     0.30s  0.35% 86.27%      1.18s  1.37%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.29s  0.34% 86.60%     22.58s 26.17%  example.com/isureal-company/isureal.imageSignature
     0.27s  0.31% 86.91%      0.85s  0.99%  database/sql.convertAssignRows
     0.25s  0.29% 87.20%     16.76s 19.43%  crypto/internal/fips140/sha256.block
     0.25s  0.29% 87.49%      0.87s  1.01%  runtime.scanObjectSmall
     0.19s  0.22% 87.71%      0.84s  0.97%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.19s  0.22% 87.93%      1.09s  1.26%  runtime.mallocgcSmallScanNoHeader
     0.16s  0.19% 88.12%      0.61s  0.71%  runtime.sweepone
     0.15s  0.17% 88.29%     16.22s 18.80%  runtime.mallocgc
     0.13s  0.15% 88.44%      1.48s  1.72%  runtime.mapassign_faststr
     0.10s  0.12% 88.56%      3.96s  4.59%  image/jpeg.idct
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 71.1% / 100% | 62% | 0.1% | 6 | 198 MB | CPU が飽和 |
| isu2 (2 cpu) | 25.3% / 100% | 20% | 19.8% | 3 | 118 MB |  |
| isu3 (2 cpu) | 90.4% / 99% | 66% | 67.9% | 5 | 125 MB | CPU が飽和 |
| isu4 (2 cpu) | 23.6% / 96% | 1% | 15.2% | 9 | 104 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu3（CPU が飽和、主に image-export）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 129% / nginx 9% |
| isu2 | postgres 6% / image-export 2% / kswapd0 1% |
| isu3 | image-export 25% / postgres 11% / kswapd0 2% / kworker/u8:0+flush-259:0 2% / jbd2/nvme0n1p1-8 1% |
| isu4 | postgres 14% |

## Diff from previous run

スコア: 20225 → 16943 (-3282)

`d3eb91b..a4d727c`

```
 conf/isu1/etc/systemd/system/isureal-go.service |  2 +-
 webapp/go/admin_posting_window_api.go           | 17 ++++++-----------
 webapp/go/store.go                              | 11 +++++++++++
 3 files changed, 18 insertions(+), 12 deletions(-)
```
