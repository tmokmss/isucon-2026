# run 20261031-111738-5948f35

## Summary

| | |
|---|---|
| score | 19275 |
| status | pass |
| ref | origin/main |
| sha | 5948f352b680 |
| 所要時間 | 289s |
| 前回比 | +3607 (23.0%) vs 20261031-110450-8f720c2 |
| ベスト比 | -5526 (-22.3%) vs 20261031-103811-cef5fe3 = 24801 |
| note | main 測り直し（書き出し完了後） |

## Bench

```
{"id":"3535","status":"done","portalStatus":"finished","passed":true,"score":19275,"raw":19275,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:21:11.114Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:21:13.597Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:21:16.034Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:21:20.448Z level=INFO msg="asia 2 日目: 24 人"
time=2026-10-31T02:21:22.178Z level=INFO msg="europe 2 日目: 9 人"
time=2026-10-31T02:21:23.808Z level=INFO msg="user moved" user=hemuor.5192 from=asia to=america
time=2026-10-31T02:21:25.145Z level=INFO msg="america 2 日目: 9 人"
time=2026-10-31T02:21:28.837Z level=INFO msg="asia 3 日目: 42 人"
time=2026-10-31T02:21:31.409Z level=INFO msg="europe 3 日目: 17 人"
time=2026-10-31T02:21:34.170Z level=INFO msg="america 3 日目: 18 人"
time=2026-10-31T02:21:37.484Z level=INFO msg="user moved" user=hemuor.5192 from=america to=asia
time=2026-10-31T02:21:39.261Z level=INFO msg="asia 4 日目: 84 人"
time=2026-10-31T02:21:42.197Z level=INFO msg="europe 4 日目: 29 人"
time=2026-10-31T02:21:42.705Z level=INFO msg="user moved" user=hemax_5465 from=america to=europe
time=2026-10-31T02:21:43.371Z level=INFO msg="america 4 日目: 29 人"
time=2026-10-31T02:21:46.137Z level=INFO msg="user moved" user=hemax_5465 from=europe to=asia
time=2026-10-31T02:21:48.853Z level=INFO msg="asia 5 日目: 85 人"
time=2026-10-31T02:21:51.086Z level=INFO msg="europe 5 日目: 29 人"
time=2026-10-31T02:21:54.929Z level=INFO msg="america 5 日目: 29 人"
time=2026-10-31T02:22:03.273Z level=INFO msg="america 最終日（グランプリ）: 29 人"
time=2026-10-31T02:22:06.760Z level=INFO msg="asia 最終日（グランプリ）: 84 人"
time=2026-10-31T02:22:09.994Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 4 人"
{
  "pass": true,
  "score": 19275,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.6%	255	0	0	GET	/api/notifications	2764.300	10.840	18.836	20.054	177
13.7%	297	16	0	POST	/api/posts	885.535	2.982	8.118	8.392	599
9.1%	2214	1287	0	POST	/api/auth/login	587.376	0.265	2.054	5.394	121
6.3%	3350	380	0	GET	^/api/images/[^/]+	409.983	0.122	1.499	3.809	88342
5.8%	274	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	373.339	1.363	8.049	12.877	116
4.6%	480	19	0	POST	/api/friend-requests	300.483	0.626	5.339	6.308	112
4.1%	182	138	0	PUT	/api/me/profile-image	267.849	1.472	8.274	8.500	18
2.6%	1618	0	0	GET	^/api/feeds/[^/]+	170.733	0.106	0.786	1.498	1819
1.7%	867	0	0	GET	/api/memories	111.462	0.129	1.076	1.974	11567
1.6%	73	0	0	GET	/api/me/export	105.080	1.439	13.202	13.202	1987308
1.4%	372	0	0	GET	/api/posting-window/current	90.270	0.243	1.719	1.996	237
1.3%	100	1	0	POST	/api/auth/register	85.056	0.851	3.350	3.711	230
0.6%	294	0	0	GET	^/api/relationships/[^/]+	37.594	0.128	0.667	1.414	1661
0.6%	55	1	0	POST	^/api/posts/[^/]+/realmojis$	35.661	0.648	5.353	5.353	449
0.5%	18	0	0	POST	/api/admin/posting-window	34.413	1.912	8.719	8.719	177
0.4%	232	7	0	GET	^/api/posts/me/current$	26.008	0.112	0.464	0.476	612
0.4%	35	0	0	POST	^/api/posts/[^/]+/comments$	24.841	0.710	5.683	5.683	349
0.4%	263	5	0	GET	^/api/posts/[^/]+$	24.027	0.091	0.522	1.358	604
0.3%	226	0	0	GET	/api/users	20.748	0.092	0.434	0.511	250
0.3%	59	0	0	PUT	^/api/follows/[^/]+$	19.365	0.328	2.845	2.845	0
0.2%	163	0	0	POST	/api/auth/logout	14.733	0.090	1.023	1.325	0
0.2%	1	0	0	POST	/api/initialize	12.149	12.149	12.149	12.149	90
0.2%	114	0	0	GET	^/api/users/[^/]+$	11.821	0.104	0.987	1.340	191
0.1%	17	0	0	DELETE	^/api/posts/[^/]+$	8.855	0.521	2.378	2.378	0
0.1%	8	0	0	DELETE	^/api/comments/[^/]+$	7.774	0.972	2.932	2.932	0
0.1%	146	0	0	GET	^/api/memories/[^/]+	7.521	0.052	0.535	1.427	603
0.1%	18	1	0	PATCH	^/api/posts/[^/]+$	5.925	0.329	1.376	1.376	620
0.1%	82	0	0	GET	^/api/posts/[^/]+/realmojis$	5.577	0.068	0.540	0.540	237
0.1%	12	0	0	DELETE	^/api/realmojis/[^/]+$	5.547	0.462	1.781	1.781	0
0.1%	6	0	0	PUT	^/api/blocks/[^/]+$	5.077	0.846	4.547	4.547	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
47.4%	115	0	0	GET	/api/notifications	969.311	8.429	19.176	20.321	135
10.0%	74	0	0	POST	/api/posts	204.821	2.768	9.854	9.854	626
7.4%	174	9	0	POST	/api/friend-requests	152.015	0.874	5.944	6.681	111
5.9%	528	0	0	GET	^/api/feeds/[^/]+	121.209	0.230	1.282	1.541	2097
5.6%	157	0	0	POST	/api/auth/login	113.631	0.724	4.857	4.882	253
4.3%	556	1	0	GET	^/api/images/[^/]+	87.977	0.158	0.949	1.152	172316
4.2%	63	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	85.471	1.357	11.427	11.427	116
2.4%	194	0	0	GET	/api/memories	49.415	0.255	1.289	2.126	10933
1.4%	27	0	0	POST	/api/auth/register	28.366	1.051	3.618	3.618	228
1.4%	100	0	0	GET	/api/posting-window/current	28.337	0.283	1.334	1.541	239
1.3%	95	0	0	GET	^/api/relationships/[^/]+	26.521	0.279	1.270	1.270	1056
1.1%	12	0	0	GET	/api/me/export	23.486	1.957	6.942	6.942	2538611
1.1%	100	4	0	GET	^/api/posts/[^/]+$	21.729	0.217	1.148	1.390	578
0.8%	21	0	0	PUT	^/api/follows/[^/]+$	16.872	0.803	4.976	4.976	0
0.7%	31	0	0	POST	/api/auth/logout	15.246	0.492	3.310	3.310	0
0.7%	70	0	0	GET	^/api/posts/me/current$	15.126	0.216	0.834	0.834	624
0.7%	69	0	0	GET	/api/users	13.885	0.201	0.618	0.618	233
0.6%	49	0	0	GET	^/api/users/[^/]+$	11.534	0.235	1.544	1.544	183
0.5%	2	0	0	POST	^/api/posts/[^/]+/comments$	10.013	5.006	6.057	6.057	283
0.4%	11	0	0	PUT	/api/me/profile-image	9.025	0.820	1.765	1.765	0
0.4%	13	0	0	PATCH	/api/me	8.552	0.658	2.226	2.226	268
0.3%	26	0	0	GET	^/api/memories/[^/]+	6.609	0.254	1.579	1.579	705
0.2%	9	0	0	PATCH	^/api/posts/[^/]+$	5.011	0.557	1.547	1.547	605
0.2%	16	2	0	GET	/api/me	3.916	0.245	0.441	0.441	250
0.2%	24	0	0	GET	^/api/posts/[^/]+/comments$	3.527	0.147	0.324	0.324	121
0.2%	26	0	0	GET	^/api/posts/[^/]+/realmojis$	3.471	0.134	0.182	0.182	98
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	2.698	0.899	2.347	2.347	0
0.1%	1	0	0	POST	/api/me/account-deletion	2.631	2.631	2.631	2.631	0
0.1%	5	0	0	POST	^/api/posts/[^/]+/realmojis$	1.678	0.336	0.832	0.832	431
0.1%	1	0	0	DELETE	^/api/posts/[^/]+$	1.297	1.297	1.297	1.297	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.6%	82	0	0	GET	/api/notifications	679.962	8.292	11.428	11.428	171
19.1%	85	2	0	POST	/api/posts	345.310	4.062	9.657	9.657	624
10.2%	162	3	0	POST	/api/friend-requests	184.060	1.136	6.135	6.514	113
6.5%	152	0	0	POST	/api/auth/login	117.718	0.774	4.764	5.187	255
4.4%	532	2	0	GET	^/api/images/[^/]+	80.394	0.151	0.943	0.953	143458
3.4%	237	0	0	GET	^/api/feeds/[^/]+	60.966	0.257	1.147	1.850	1150
3.2%	183	0	0	GET	/api/memories	57.681	0.315	2.403	4.537	10996
2.3%	31	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	41.913	1.352	8.895	8.895	116
2.2%	29	0	0	POST	/api/auth/register	39.198	1.352	4.288	4.288	229
2.1%	114	0	0	GET	/api/posting-window/current	37.982	0.333	1.537	1.617	240
1.1%	81	0	0	GET	/api/users	20.340	0.251	1.257	1.257	244
0.9%	62	0	0	GET	^/api/relationships/[^/]+	16.648	0.269	1.596	1.596	1166
0.9%	66	11	0	GET	^/api/posts/me/current$	16.550	0.251	1.541	1.541	531
0.9%	11	0	0	GET	/api/me/export	16.266	1.479	3.108	3.108	2309014
0.8%	20	0	0	PUT	^/api/follows/[^/]+$	14.201	0.710	3.650	3.650	0
0.7%	15	0	0	PUT	/api/me/profile-image	13.526	0.902	2.880	2.880	0
0.6%	24	0	0	POST	/api/auth/logout	11.708	0.488	2.546	2.546	0
0.5%	12	0	0	PATCH	/api/me	9.348	0.779	3.897	3.897	307
0.5%	3	0	0	DELETE	^/api/posts/[^/]+$	8.373	2.791	3.961	3.961	0
0.4%	34	3	0	GET	^/api/posts/[^/]+$	7.886	0.232	0.618	0.618	549
0.3%	5	1	0	PATCH	^/api/posts/[^/]+$	6.082	1.216	4.165	4.165	488
0.3%	25	0	0	GET	^/api/memories/[^/]+	6.015	0.241	1.153	1.153	921
0.2%	18	0	0	GET	/api/me	3.974	0.221	0.450	0.450	288
0.2%	13	0	0	POST	^/api/posts/[^/]+/comments$	3.950	0.304	0.735	0.735	282
0.1%	13	0	0	GET	^/api/posts/[^/]+/realmojis$	2.377	0.183	0.350	0.350	174
0.1%	13	0	0	GET	^/api/users/[^/]+$	2.370	0.182	0.361	0.361	191
0.1%	13	0	0	GET	^/api/posts/[^/]+/comments$	2.192	0.169	0.513	0.513	266
0.0%	4	0	0	DELETE	^/api/comments/[^/]+$	0.678	0.169	0.210	0.210	0
0.0%	1	0	0	DELETE	^/api/follows/[^/]+$	0.616	0.616	0.616	0.616	0
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.108	0.108	0.108	0.108	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 15122 クエリ / 80 種類 / 合計 70.705s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	53.7%	96	37.994	395.77	3140.20	6533.8	133.7	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	10.7%	567	7.596	13.40	137.53	739.0	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	9.2%	18	6.538	363.21	1341.09	32499.3	1654.6	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
4	7.9%	926	5.560	6.00	305.62	114.1	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	6.0%	2368	4.234	1.79	7.31	11.3	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
6	3.4%	1	2.423	2422.61	2422.61	29343.0	1199488.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	2.2%	1	1.588	1588.09	1588.09	19659.0	625586.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	1.6%	1	1.097	1096.97	1096.97	11528.0	529738.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	1.0%	438	0.709	1.62	5.35	32.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.5%	2	0.370	184.88	209.04	214448.5	2522.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.4%	2	0.293	146.31	283.94	20237.0	2515.5	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.3%	1	0.227	227.45	227.45	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
13	0.3%	1	0.219	219.03	219.03	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.3%	1	0.207	206.62	206.62	69837.0	29725.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.3%	59	0.179	3.04	5.75	48.0	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 3
INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 0a386750252c141d57f089186c48674435db2347
Type: cpu
Time: 2026-10-31 11:20:53 JST
Duration: 60.29s, Total samples = 86.12s (142.84%)
Showing nodes accounting for 76.39s, 88.70% of 86.12s total
Dropped 823 nodes (cum <= 0.43s)
Showing top 40 nodes out of 169
      flat  flat%   sum%        cum   cum%
    19.59s 22.75% 22.75%     19.59s 22.75%  crypto/internal/fips140/sha256.blockAVX2
    10.86s 12.61% 35.36%     10.86s 12.61%  runtime.spanClass.sizeclass (inline)
     7.27s  8.44% 43.80%     11.56s 13.42%  image/jpeg.(*decoder).reconstructBlock
     3.64s  4.23% 48.03%     15.20s 17.65%  runtime.tryDeferToSpanScan
     3.21s  3.73% 51.75%     24.71s 28.69%  image/jpeg.(*decoder).processSOS
     2.96s  3.44% 55.19%      4.87s  5.65%  image/jpeg.(*decoder).decodeHuffman
     2.52s  2.93% 58.12%      2.52s  2.93%  runtime.memmove
     2.06s  2.39% 60.51%      2.06s  2.39%  runtime.memclrNoHeapPointers
     2.04s  2.37% 62.88%      2.18s  2.53%  image/jpeg.(*decoder).receiveExtend
     1.80s  2.09% 64.97%      7.27s  8.44%  runtime.scanObjectsSmall
     1.78s  2.07% 67.03%      2.34s  2.72%  image/jpeg.idctCols
     1.54s  1.79% 68.82%      2.05s  2.38%  image/jpeg.(*decoder).ensureNBits
     1.45s  1.68% 70.51%     11.91s 13.83%  runtime.scanObject
     1.41s  1.64% 72.14%      1.92s  2.23%  image/jpeg.idctRows
     1.36s  1.58% 73.72%      1.36s  1.58%  internal/runtime/syscall/linux.Syscall6
     1.22s  1.42% 75.14%      5.37s  6.24%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.07s  1.24% 76.38%      1.07s  1.24%  image/jpeg.dctBox (inline)
     1.02s  1.18% 77.57%     21.50s 24.97%  crypto/internal/fips140/sha256.(*Digest).Write
     0.98s  1.14% 78.70%      0.98s  1.14%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.95s  1.10% 79.81%     21.81s 25.33%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.80s  0.93% 80.74%      0.80s  0.93%  aeshashbody
     0.71s  0.82% 81.56%     22.94s 26.64%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.62s  0.72% 82.28%      0.62s  0.72%  memeqbody
     0.62s  0.72% 83.00%      8.79s 10.21%  runtime.scanSpan
     0.57s  0.66% 83.66%     24.46s 28.40%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.55s  0.64% 84.30%      0.55s  0.64%  runtime.extractHeapBitsSmall
     0.44s  0.51% 84.81%      0.50s  0.58%  image/jpeg.(*decoder).readByteStuffedByte
     0.44s  0.51% 85.32%      0.49s  0.57%  runtime.(*spanInlineMarkBits).tryAcquire
     0.38s  0.44% 85.76%     26.28s 30.52%  example.com/isureal-company/isureal.imageSignature
     0.37s  0.43% 86.19%      1.92s  2.23%  runtime.mapaccess1_faststr
     0.36s  0.42% 86.61%      1.08s  1.25%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.29s  0.34% 86.95%      0.51s  0.59%  example.com/isureal-company/isureal.postLess
     0.28s  0.33% 87.27%     19.87s 23.07%  crypto/internal/fips140/sha256.block
     0.26s   0.3% 87.58%      0.79s  0.92%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.24s  0.28% 87.85%      0.58s  0.67%  database/sql.convertAssignRows
     0.19s  0.22% 88.07%      0.78s  0.91%  runtime.scanObjectSmall
     0.17s   0.2% 88.27%      1.13s  1.31%  runtime.mapassign_faststr
     0.14s  0.16% 88.43%     12.10s 14.05%  runtime.mallocgc
     0.13s  0.15% 88.59%      0.85s  0.99%  runtime.mallocgcSmallScanNoHeader
     0.10s  0.12% 88.70%      1.21s  1.41%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 75.8% / 100% | 61% | 0.1% | 6 | 172 MB | CPU が飽和 |
| isu2 (2 cpu) | 1.1% / 4% | 0% | 0.2% | 2 | 271 MB |  |
| isu3 (2 cpu) | 1.6% / 6% | 0% | 0.4% | 0 | 281 MB |  |
| isu4 (2 cpu) | 29.7% / 89% | 0% | 18.1% | 8 | 100 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 137% / nginx 10% |
| isu2 | 目立つプロセス無し |
| isu3 | isureal 1% |
| isu4 | postgres 20% |

## Diff from previous run

スコア: 15668 → 19275 (+3607)

`8f720c2..5948f35`

```
 webapp/go/cmd/image-export/main.go | 11 +----------
 webapp/go/image_local.go           | 37 +++++++++++++++++++++++++++++--------
 2 files changed, 30 insertions(+), 18 deletions(-)
```
