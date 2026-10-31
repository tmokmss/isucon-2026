# run 20261031-113441-6a7a33a

## Summary

| | |
|---|---|
| score | 17072 |
| status | pass |
| ref | origin/main |
| sha | 6a7a33a922e9 |
| 所要時間 | 161s |
| 前回比 | +17072 (—) vs 20261031-113154-1376dd7 |
| ベスト比 | -7729 (-31.2%) vs 20261031-103811-cef5fe3 = 24801 |
| note | main: 登録の並列化を戻した（#67 入り） |

## Bench

```
{"id":"4503","status":"done","portalStatus":"finished","passed":true,"score":17072,"raw":17072,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:36:06.689Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:36:09.706Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:36:12.738Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:36:15.660Z level=INFO msg="asia 2 日目: 46 人"
time=2026-10-31T02:36:18.861Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T02:36:21.418Z level=INFO msg="user moved" user=herite_5268 from=asia to=europe
time=2026-10-31T02:36:21.641Z level=INFO msg="america 2 日目: 10 人"
time=2026-10-31T02:36:21.651Z level=INFO msg="user moved" user=raben_5252 from=asia to=europe
time=2026-10-31T02:36:25.038Z level=INFO msg="asia 3 日目: 125 人"
time=2026-10-31T02:36:28.360Z level=INFO msg="europe 3 日目: 10 人"
time=2026-10-31T02:36:31.110Z level=INFO msg="america 3 日目: 14 人"
time=2026-10-31T02:36:33.984Z level=INFO msg="asia 4 日目: 125 人"
time=2026-10-31T02:36:37.462Z level=INFO msg="europe 4 日目: 10 人"
time=2026-10-31T02:36:40.066Z level=INFO msg="america 4 日目: 14 人"
time=2026-10-31T02:36:43.389Z level=INFO msg="asia 5 日目: 125 人"
time=2026-10-31T02:36:48.309Z level=INFO msg="europe 5 日目: 10 人"
time=2026-10-31T02:36:50.132Z level=INFO msg="america 5 日目: 14 人"
time=2026-10-31T02:36:58.001Z level=INFO msg="america 最終日（グランプリ）: 14 人"
time=2026-10-31T02:36:58.140Z level=INFO msg="asia 最終日（グランプリ）: 125 人"
time=2026-10-31T02:36:58.183Z level=INFO msg="europe 最終日（グランプリ）: 10 人"
time=2026-10-31T02:37:06.601Z level=INFO msg="グランプリ: アジア 10 人, ヨーロッパ 0 人, アメリカ 2 人"
{
  "pass": true,
  "score": 17072,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.3%	389	0	0	GET	/api/notifications	3582.873	9.210	14.730	15.131	173
16.9%	454	51	0	POST	/api/posts	1878.716	4.138	10.475	11.138	575
11.2%	2363	1620	0	POST	/api/auth/login	1240.523	0.525	3.485	7.667	97
9.9%	3207	785	0	GET	^/api/images/[^/]+	1099.928	0.343	2.783	6.229	110352
5.1%	222	150	0	PUT	/api/me/profile-image	567.519	2.556	10.111	10.433	10
4.8%	710	25	0	POST	/api/friend-requests	535.185	0.754	6.525	14.205	112
4.8%	143	1	0	POST	/api/auth/register	531.375	3.716	12.439	12.589	228
4.8%	375	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	531.205	1.417	9.047	10.941	116
1.6%	1679	0	0	GET	^/api/feeds/[^/]+	178.900	0.107	0.483	4.189	1496
1.4%	861	0	0	GET	/api/memories	155.871	0.181	1.379	6.316	9230
1.3%	593	0	0	GET	/api/posting-window/current	145.522	0.245	1.201	2.722	237
0.9%	45	0	0	GET	/api/me/export	104.957	2.332	12.317	12.317	2189596
0.8%	95	0	0	PUT	^/api/follows/[^/]+$	88.200	0.928	5.819	5.819	0
0.5%	350	16	0	GET	^/api/posts/me/current$	53.387	0.153	0.671	1.033	617
0.5%	134	0	0	POST	/api/auth/logout	51.173	0.382	3.536	4.943	0
0.4%	38	5	0	POST	^/api/posts/[^/]+/realmojis$	47.709	1.255	3.117	3.117	461
0.4%	1	0	0	POST	/api/initialize	44.755	44.755	44.755	44.755	90
0.4%	345	0	0	GET	/api/users	41.053	0.119	0.951	1.339	244
0.3%	348	0	0	GET	^/api/relationships/[^/]+	38.408	0.110	0.462	1.388	1673
0.3%	34	7	0	POST	^/api/posts/[^/]+/comments$	35.488	1.044	3.966	3.966	312
0.2%	252	5	0	GET	^/api/posts/[^/]+$	22.962	0.091	0.317	0.388	609
0.2%	37	3	0	PATCH	^/api/posts/[^/]+$	21.106	0.570	2.483	2.483	589
0.1%	55	0	0	PATCH	/api/me	14.509	0.264	1.408	1.408	278
0.1%	12	0	0	DELETE	^/api/posts/[^/]+$	14.161	1.180	2.692	2.692	0
0.1%	114	0	0	GET	^/api/memories/[^/]+	10.497	0.092	0.526	1.848	671
0.1%	18	0	0	POST	/api/admin/posting-window	10.453	0.581	2.677	2.677	176
0.1%	126	0	0	GET	^/api/posts/[^/]+/realmojis$	8.607	0.068	0.293	0.346	96
0.1%	100	0	0	GET	^/api/users/[^/]+$	7.084	0.071	0.241	0.271	196
0.1%	78	2	0	GET	/api/me	6.852	0.088	0.604	0.604	291
0.1%	83	0	0	GET	^/api/posts/[^/]+/comments$	5.632	0.068	0.429	0.429	166
…（上位 30 / 全 39 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.6%	51	0	0	GET	/api/notifications	441.490	8.657	10.872	10.872	181
14.4%	44	1	0	POST	/api/posts	146.260	3.324	10.589	10.589	597
10.0%	53	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	101.416	1.914	14.245	14.245	116
5.1%	271	0	0	GET	^/api/images/[^/]+	51.308	0.189	0.977	0.979	161893
4.5%	59	0	0	POST	/api/auth/login	45.403	0.770	4.662	4.662	251
3.9%	60	2	0	POST	/api/friend-requests	39.629	0.660	6.430	6.430	112
3.6%	162	0	0	GET	^/api/feeds/[^/]+	36.130	0.223	1.146	1.381	1321
2.9%	8	0	0	GET	/api/me/export	28.879	3.610	10.580	10.580	1761791
2.4%	97	0	0	GET	/api/memories	24.613	0.254	1.344	1.344	9623
1.6%	49	0	0	GET	/api/posting-window/current	15.715	0.321	1.072	1.072	239
1.4%	6	0	0	POST	/api/auth/register	14.366	2.394	12.189	12.189	231
1.3%	12	2	0	POST	^/api/posts/[^/]+/comments$	13.134	1.095	4.488	4.488	238
0.7%	9	0	0	POST	/api/auth/logout	7.143	0.794	1.688	1.688	0
0.7%	37	1	0	GET	^/api/posts/me/current$	6.844	0.185	0.524	0.524	592
0.7%	32	2	0	GET	^/api/posts/[^/]+$	6.663	0.208	0.472	0.472	560
0.6%	34	0	0	GET	^/api/relationships/[^/]+	5.951	0.175	0.421	0.421	2463
0.5%	26	0	0	GET	/api/users	4.661	0.179	0.563	0.563	279
0.3%	14	0	0	GET	^/api/users/[^/]+$	2.724	0.195	0.300	0.300	177
0.3%	12	0	0	GET	^/api/posts/[^/]+/realmojis$	2.679	0.223	0.797	0.797	118
0.2%	2	2	0	PATCH	^/api/posts/[^/]+$	2.432	1.216	2.222	2.222	30
0.2%	5	0	0	PUT	^/api/follows/[^/]+$	2.418	0.484	1.987	1.987	0
0.2%	1	0	0	POST	^/api/posts/[^/]+/realmojis$	2.393	2.393	2.393	2.393	548
0.2%	1	0	0	DELETE	^/api/realmojis/[^/]+$	2.274	2.274	2.274	2.274	0
0.2%	8	0	0	GET	^/api/posts/[^/]+/comments$	1.953	0.244	0.741	0.741	181
0.2%	7	0	0	GET	^/api/memories/[^/]+	1.937	0.277	0.543	0.543	989
0.1%	1	0	0	PUT	/api/me/profile-image	1.321	1.321	1.321	1.321	0
0.1%	5	1	0	GET	/api/me	1.018	0.204	0.512	0.512	184
0.1%	2	0	0	PUT	^/api/blocks/[^/]+$	0.647	0.324	0.537	0.537	0
0.1%	3	0	0	DELETE	^/api/comments/[^/]+$	0.614	0.205	0.312	0.312	0
0.0%	1	0	0	PATCH	/api/me	0.484	0.484	0.484	0.484	343

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.2%	59	0	0	GET	/api/notifications	464.047	7.865	10.234	10.234	182
20.7%	61	1	0	POST	/api/posts	244.620	4.010	11.196	11.196	603
4.9%	81	5	0	POST	/api/friend-requests	57.728	0.713	3.501	3.501	110
4.8%	73	0	0	POST	/api/auth/login	57.112	0.782	6.974	6.974	250
4.5%	287	5	0	GET	^/api/images/[^/]+	53.185	0.185	0.961	0.966	181126
3.6%	13	0	0	POST	/api/auth/register	42.546	3.273	10.703	10.703	229
3.5%	114	0	0	GET	/api/memories	41.517	0.364	3.563	4.466	12394
3.3%	150	0	0	GET	^/api/feeds/[^/]+	38.821	0.259	1.079	1.842	1753
3.0%	36	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	35.829	0.995	7.183	7.183	116
2.2%	76	0	0	GET	/api/posting-window/current	25.482	0.335	1.586	1.586	239
2.0%	7	0	0	GET	/api/me/export	23.660	3.380	9.615	9.615	2634805
1.7%	11	0	0	PUT	^/api/follows/[^/]+$	20.395	1.854	9.583	9.583	0
1.3%	50	4	0	GET	^/api/posts/me/current$	15.628	0.313	1.312	1.312	562
0.9%	46	0	0	GET	^/api/relationships/[^/]+	10.545	0.229	1.496	1.496	1476
0.8%	7	2	0	POST	^/api/posts/[^/]+/comments$	9.643	1.378	4.164	4.164	221
0.8%	5	0	0	DELETE	^/api/posts/[^/]+$	9.551	1.910	4.441	4.441	0
0.8%	43	0	0	GET	/api/users	9.222	0.214	0.615	0.615	235
0.3%	2	0	0	PATCH	^/api/posts/[^/]+$	3.780	1.890	2.925	2.925	590
0.3%	15	1	0	GET	^/api/posts/[^/]+$	3.666	0.244	0.465	0.465	547
0.3%	16	0	0	GET	^/api/users/[^/]+$	3.530	0.221	0.473	0.473	172
0.3%	10	0	0	POST	/api/auth/logout	3.505	0.351	0.667	0.667	0
0.2%	9	0	0	GET	^/api/posts/[^/]+/comments$	2.597	0.289	0.688	0.688	295
0.2%	4	0	0	PUT	/api/me/profile-image	1.871	0.468	0.758	0.758	0
0.2%	3	0	0	PATCH	/api/me	1.779	0.593	0.890	0.890	257
0.1%	7	0	0	GET	^/api/memories/[^/]+	1.741	0.249	0.691	0.691	1333
0.1%	8	0	0	GET	^/api/posts/[^/]+/realmojis$	1.253	0.157	0.270	0.270	29
0.1%	7	0	0	GET	/api/me	1.128	0.161	0.367	0.367	256
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.107	0.107	0.107	0.107	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 13495 クエリ / 169 種類 / 合計 186.884s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	48.5%	3	90.575	30191.76	90574.10	5.3	0.0	ALTER TABLE posting_windows ADD COLUMN IF NOT EXISTS notification_ids text[], ADD COLUMN IF NOT EXISTS notified_user_ids text[]
2	23.8%	60	44.485	741.41	9773.47	6357.4	139.3	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	6.1%	2	11.384	5691.96	7736.90	212961.0	2505.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
4	3.9%	1	7.348	7347.91	7347.91	29389.0	1199348.0	COPY public.realmojis (id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at) TO stdout
5	3.5%	623	6.557	10.53	90.95	766.0	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
6	2.1%	486	3.903	8.03	205.40	117.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
7	1.8%	1	3.363	3363.30	3363.30	29389.0	1199365.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	1.6%	1677	2.954	1.76	6.48	13.9	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
9	1.5%	1	2.874	2874.39	2874.39	19882.0	625111.0	COPY public.posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted…
10	1.3%	1	2.485	2484.52	2484.52	19882.0	625231.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
11	1.3%	1	2.465	2464.62	2464.62	11528.0	546011.0	COPY public.comments (id, post_id, author_id, body, created_at, deleted_at) TO stdout
12	1.2%	1	2.205	2204.74	2204.74	11528.0	529637.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
13	0.8%	1	1.518	1517.63	1517.63	29407.0	17.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.5%	507	0.859	1.69	8.58	32.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
15	0.4%	1	0.785	784.78	784.78	20123.0	120.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
ALTER TABLE posting_windows ADD COLUMN IF NOT EXISTS notification_ids text[], ADD COLUMN IF NOT EXISTS notified_user_ids text[];
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: c523838c70ed9068fc14a0bb4e188c38014daeaf
Type: cpu
Time: 2026-10-31 11:35:16 JST
Duration: 60.05s, Total samples = 27.35s (45.54%)
Showing nodes accounting for 22.19s, 81.13% of 27.35s total
Dropped 507 nodes (cum <= 0.14s)
Showing top 40 nodes out of 184
      flat  flat%   sum%        cum   cum%
     4.43s 16.20% 16.20%      4.43s 16.20%  crypto/internal/fips140/sha256.blockAVX2
     2.01s  7.35% 23.55%      2.01s  7.35%  runtime.spanClass.sizeclass (inline)
     1.42s  5.19% 28.74%      5.67s 20.73%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.16s  4.24% 32.98%      1.16s  4.24%  runtime.memmove
     1.08s  3.95% 36.93%      1.08s  3.95%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
        1s  3.66% 40.59%      1.62s  5.92%  image/jpeg.(*decoder).reconstructBlock
     0.98s  3.58% 44.17%      3.13s 11.44%  runtime.tryDeferToSpanScan
     0.93s  3.40% 47.57%      0.93s  3.40%  runtime.memclrNoHeapPointers
     0.80s  2.93% 50.49%      0.80s  2.93%  aeshashbody
     0.51s  1.86% 52.36%      0.80s  2.93%  image/jpeg.(*decoder).decodeHuffman
     0.51s  1.86% 54.22%      3.46s 12.65%  image/jpeg.(*decoder).processSOS
     0.50s  1.83% 56.05%      0.50s  1.83%  memeqbody
     0.48s  1.76% 57.81%      0.48s  1.76%  internal/runtime/syscall/linux.Syscall6
     0.44s  1.61% 59.41%      1.67s  6.11%  runtime.scanObjectsSmall
     0.39s  1.43% 60.84%      0.39s  1.43%  runtime.nextFreeFast (inline)
     0.36s  1.32% 62.16%      2.39s  8.74%  runtime.scanObject
     0.34s  1.24% 63.40%      0.55s  2.01%  example.com/isureal-company/isureal.postLess
     0.32s  1.17% 64.57%      0.35s  1.28%  image/jpeg.(*decoder).receiveExtend
     0.31s  1.13% 65.70%      1.88s  6.87%  runtime.mapaccess1_faststr
     0.28s  1.02% 66.73%      1.23s  4.50%  runtime.mallocgcSmallScanNoHeader
     0.27s  0.99% 67.71%      0.55s  2.01%  database/sql.convertAssignRows
     0.26s  0.95% 68.67%      0.32s  1.17%  image/jpeg.(*decoder).ensureNBits
     0.24s  0.88% 69.54%      5.04s 18.43%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.23s  0.84% 70.38%      0.23s  0.84%  runtime.extractHeapBitsSmall
     0.22s   0.8% 71.19%      0.22s   0.8%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.21s  0.77% 71.96%      4.90s 17.92%  crypto/internal/fips140/sha256.(*Digest).Write
     0.21s  0.77% 72.72%      0.31s  1.13%  image/jpeg.idctRows
     0.21s  0.77% 73.49%      1.17s  4.28%  runtime.mapassign_faststr
     0.21s  0.77% 74.26%      0.21s  0.77%  sync/atomic.(*Int32).Add
     0.20s  0.73% 74.99%      0.21s  0.77%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.20s  0.73% 75.72%      0.27s  0.99%  image/jpeg.idctCols
     0.20s  0.73% 76.45%      0.23s  0.84%  runtime.(*mspan).writeHeapBitsSmall
     0.18s  0.66% 77.11%      3.39s 12.39%  runtime.mallocgc
     0.17s  0.62% 77.73%      3.91s 14.30%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.16s  0.59% 78.32%      5.35s 19.56%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.16s  0.59% 78.90%      0.19s  0.69%  example.com/isureal-company/isureal.rmLess
     0.16s  0.59% 79.49%      0.73s  2.67%  runtime.mallocgcSmallNoscan
     0.15s  0.55% 80.04%      0.34s  1.24%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.15s  0.55% 80.59%      0.15s  0.55%  image/jpeg.dctBox (inline)
     0.15s  0.55% 81.13%      0.15s  0.55%  runtime.(*spanInlineMarkBits).init
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 74.5% / 100% | 58% | 0.2% | 6 | 108 MB | CPU が飽和 |
| isu2 (2 cpu) | 12.7% / 88% | 0% | 7.1% | 2 | 392 MB |  |
| isu3 (2 cpu) | 1.0% / 4% | 0% | 0.1% | 0 | 259 MB |  |
| isu4 (2 cpu) | 25.0% / 82% | 0% | 14.7% | 9 | 105 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 99% / pg_dump 33% / nginx 9% / sshd 3% |
| isu2 | find 8% / kswapd0 1% |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 18% |

## Diff from previous run

スコア: 0 → 17072 (+17072)

`1376dd7..6a7a33a`

```
 webapp/go/user_api.go | 17 +++--------------
 1 file changed, 3 insertions(+), 14 deletions(-)
```
