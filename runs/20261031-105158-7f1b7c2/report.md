# run 20261031-105158-7f1b7c2

## Summary

| | |
|---|---|
| score | 14462 |
| status | pass |
| ref | jpeg-check-cache |
| sha | 7f1b7c2c9a2e |
| 所要時間 | 275s |
| 前回比 | -10339 (-41.7%) vs 20261031-103811-cef5fe3 |
| ベスト比 | -10339 (-41.7%) vs 20261031-103811-cef5fe3 = 24801 |
| note | JPEGの検証結果をsha256でキャッシュ |

## Bench

```
{"id":"2007","status":"done","portalStatus":"finished","passed":true,"score":14462,"raw":14462,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:55:21.289Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:55:23.091Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:55:25.703Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:55:29.434Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T01:55:31.597Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T01:55:34.579Z level=INFO msg="america 2 日目: 9 人"
time=2026-10-31T01:55:38.167Z level=INFO msg="asia 3 日目: 32 人"
time=2026-10-31T01:55:40.751Z level=INFO msg="europe 3 日目: 10 人"
time=2026-10-31T01:55:40.779Z level=INFO msg="user moved" user=rois.5207 from=asia to=europe
time=2026-10-31T01:55:43.626Z level=INFO msg="america 3 日目: 14 人"
time=2026-10-31T01:55:47.482Z level=INFO msg="asia 4 日目: 71 人"
time=2026-10-31T01:55:50.783Z level=INFO msg="europe 4 日目: 19 人"
time=2026-10-31T01:55:53.129Z level=INFO msg="america 4 日目: 19 人"
time=2026-10-31T01:55:59.660Z level=INFO msg="asia 5 日目: 78 人"
time=2026-10-31T01:56:03.573Z level=INFO msg="europe 5 日目: 19 人"
time=2026-10-31T01:56:07.770Z level=INFO msg="user moved" user=rois.5207 from=europe to=america
time=2026-10-31T01:56:08.326Z level=INFO msg="america 5 日目: 20 人"
time=2026-10-31T01:56:13.507Z level=INFO msg="america 最終日（グランプリ）: 20 人"
time=2026-10-31T01:56:15.825Z level=INFO msg="europe 最終日（グランプリ）: 18 人"
time=2026-10-31T01:56:19.344Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 1 人, アメリカ 6 人"
{
  "pass": true,
  "score": 14462,
  "messages": [
    "asia の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.5%	222	0	0	GET	/api/notifications	2701.692	12.170	23.709	23.851	174
9.9%	1994	1214	0	POST	/api/auth/login	626.789	0.314	2.077	5.443	115
9.2%	192	2	0	POST	/api/posts	586.255	3.053	10.880	11.084	639
8.4%	2905	471	0	GET	^/api/images/[^/]+	532.505	0.183	2.036	4.273	84768
8.2%	193	152	0	PUT	/api/me/profile-image	522.569	2.708	11.511	12.788	18
5.0%	448	22	0	POST	/api/friend-requests	314.921	0.703	4.567	8.009	111
4.0%	192	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	252.192	1.314	10.233	10.527	116
2.9%	1450	0	0	GET	^/api/feeds/[^/]+	182.713	0.126	0.663	1.327	1634
1.7%	95	1	0	POST	/api/auth/register	108.002	1.137	9.643	9.643	230
1.5%	719	0	0	GET	/api/memories	93.846	0.131	1.010	4.679	11535
0.7%	259	0	0	GET	/api/posting-window/current	45.712	0.176	0.980	1.092	237
0.7%	18	0	0	POST	/api/admin/posting-window	44.884	2.494	10.213	10.213	176
0.5%	138	0	0	POST	/api/auth/logout	33.845	0.245	3.815	10.386	0
0.5%	51	0	0	GET	/api/me/export	32.882	0.645	2.481	2.481	1889734
0.5%	190	7	0	GET	^/api/posts/me/current$	30.702	0.162	1.266	1.387	622
0.5%	204	0	0	GET	/api/users	30.645	0.150	0.899	1.143	242
0.5%	31	0	0	POST	^/api/posts/[^/]+/realmojis$	28.906	0.932	3.250	3.250	425
0.4%	233	0	0	GET	^/api/relationships/[^/]+	22.941	0.098	0.591	0.824	1722
0.3%	223	1	0	GET	^/api/posts/[^/]+$	22.104	0.099	0.432	0.789	598
0.3%	45	0	0	PUT	^/api/follows/[^/]+$	20.001	0.444	3.647	3.647	0
0.3%	27	2	0	PATCH	^/api/posts/[^/]+$	17.661	0.654	2.463	2.463	599
0.2%	1	0	0	POST	/api/initialize	14.297	14.297	14.297	14.297	90
0.2%	35	0	0	PATCH	/api/me	13.984	0.400	2.875	2.875	264
0.2%	20	0	0	POST	^/api/posts/[^/]+/comments$	12.428	0.621	3.531	3.531	381
0.2%	96	0	0	GET	^/api/users/[^/]+$	12.048	0.125	0.614	0.614	174
0.1%	13	0	0	DELETE	^/api/posts/[^/]+$	8.327	0.641	1.672	1.672	0
0.1%	141	0	0	GET	^/api/memories/[^/]+	6.827	0.048	0.394	0.583	515
0.1%	2	0	0	DELETE	^/api/follows/[^/]+$	6.411	3.206	6.408	6.408	0
0.1%	5	0	0	PUT	^/api/blocks/[^/]+$	4.951	0.990	4.164	4.164	0
0.1%	7	0	0	DELETE	^/api/comments/[^/]+$	4.515	0.645	1.224	1.224	0
…（上位 30 / 全 39 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.3%	61	0	0	GET	/api/notifications	602.670	9.880	16.302	16.302	181
13.1%	59	0	0	POST	/api/posts	186.555	3.162	10.317	10.317	623
9.8%	424	2	0	GET	^/api/images/[^/]+	139.828	0.330	1.716	3.365	185307
8.1%	116	6	0	POST	/api/friend-requests	115.457	0.995	7.462	10.427	111
5.0%	356	0	0	GET	^/api/feeds/[^/]+	71.872	0.202	0.953	1.214	1721
4.3%	101	0	0	POST	/api/auth/login	61.743	0.611	4.495	6.668	251
3.7%	37	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	53.046	1.434	6.593	6.593	116
2.2%	128	0	0	GET	/api/memories	31.645	0.247	1.076	1.841	11258
1.6%	73	0	0	GET	/api/posting-window/current	22.904	0.314	0.904	0.904	238
1.5%	8	0	0	GET	/api/me/export	21.341	2.668	7.865	7.865	2986467
0.9%	54	0	0	GET	/api/users	12.782	0.237	0.761	0.761	257
0.9%	45	2	0	GET	^/api/posts/me/current$	12.390	0.275	1.043	1.043	594
0.8%	58	0	0	GET	^/api/relationships/[^/]+	11.487	0.198	1.060	1.060	1169
0.7%	16	0	0	POST	/api/auth/register	9.484	0.593	1.781	1.781	228
0.7%	31	0	0	POST	^/api/posts/[^/]+/realmojis$	9.449	0.305	3.441	3.441	420
0.6%	63	2	0	GET	^/api/posts/[^/]+$	9.076	0.144	0.313	0.313	579
0.5%	12	0	0	PUT	^/api/follows/[^/]+$	7.421	0.618	3.023	3.023	0
0.5%	5	0	0	PATCH	/api/me	7.184	1.437	5.134	5.134	300
0.5%	26	0	0	GET	^/api/users/[^/]+$	7.156	0.275	1.145	1.145	172
0.4%	6	0	0	PUT	/api/me/profile-image	5.621	0.937	1.661	1.661	0
0.3%	5	0	0	POST	^/api/posts/[^/]+/comments$	4.620	0.924	2.483	2.483	275
0.3%	16	0	0	GET	^/api/memories/[^/]+	4.311	0.269	1.214	1.214	524
0.2%	11	0	0	POST	/api/auth/logout	3.458	0.314	1.331	1.331	0
0.2%	11	0	0	DELETE	^/api/realmojis/[^/]+$	3.262	0.297	0.816	0.816	0
0.2%	24	0	0	GET	^/api/posts/[^/]+/realmojis$	3.096	0.129	0.312	0.312	513
0.2%	10	1	0	GET	/api/me	2.244	0.224	0.443	0.443	246
0.1%	2	0	0	DELETE	^/api/posts/[^/]+$	1.548	0.774	1.445	1.445	0
0.1%	10	0	0	GET	^/api/posts/[^/]+/comments$	1.376	0.138	0.243	0.243	193
0.1%	5	0	0	PATCH	^/api/posts/[^/]+$	1.341	0.268	0.745	0.745	605
0.0%	2	0	0	PUT	^/api/blocks/[^/]+$	0.599	0.299	0.493	0.493	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.3%	66	0	0	GET	/api/notifications	587.176	8.897	15.328	15.328	179
15.5%	60	1	0	POST	/api/posts	210.421	3.507	11.289	11.289	620
12.4%	395	0	0	GET	^/api/images/[^/]+	167.885	0.425	3.609	4.949	120876
4.8%	85	3	0	POST	/api/friend-requests	65.363	0.769	9.694	9.694	112
4.3%	103	0	0	POST	/api/auth/login	58.382	0.567	2.562	5.040	249
3.9%	188	0	0	GET	^/api/feeds/[^/]+	52.262	0.278	1.402	1.404	927
3.3%	147	0	0	GET	/api/memories	45.156	0.307	1.971	2.497	16290
2.3%	26	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	31.737	1.221	6.737	6.737	116
1.7%	12	0	0	PUT	^/api/follows/[^/]+$	23.473	1.956	9.570	9.570	0
1.7%	80	0	0	GET	/api/posting-window/current	23.020	0.288	0.577	0.577	239
1.0%	48	9	0	GET	^/api/posts/me/current$	13.716	0.286	1.174	1.174	511
0.8%	54	0	0	GET	^/api/relationships/[^/]+	10.730	0.199	1.386	1.386	831
0.7%	19	0	0	POST	/api/auth/logout	9.956	0.524	1.968	1.968	0
0.7%	42	0	0	GET	/api/users	9.685	0.231	1.206	1.206	229
0.5%	20	2	0	GET	^/api/posts/[^/]+$	6.729	0.336	0.909	0.909	542
0.5%	18	0	0	POST	/api/auth/register	6.294	0.350	1.067	1.067	230
0.4%	2	0	0	PATCH	^/api/posts/[^/]+$	4.783	2.392	3.503	3.503	640
0.3%	6	0	0	PUT	/api/me/profile-image	4.430	0.738	2.227	2.227	0
0.3%	8	0	0	POST	^/api/posts/[^/]+/realmojis$	3.885	0.486	1.122	1.122	404
0.3%	3	0	0	GET	/api/me/export	3.804	1.268	1.653	1.653	2272285
0.3%	7	0	0	PATCH	/api/me	3.720	0.531	2.086	2.086	288
0.3%	17	0	0	GET	^/api/memories/[^/]+	3.594	0.211	0.460	0.460	1087
0.1%	2	0	0	POST	^/api/posts/[^/]+/comments$	2.011	1.006	1.124	1.124	415
0.1%	11	0	0	GET	^/api/users/[^/]+$	1.970	0.179	0.288	0.288	175
0.1%	1	0	0	DELETE	^/api/posts/[^/]+$	1.360	1.360	1.360	1.360	0
0.1%	9	0	0	GET	/api/me	1.299	0.144	0.314	0.314	273
0.1%	7	0	0	GET	^/api/posts/[^/]+/realmojis$	1.079	0.154	0.200	0.200	322
0.1%	3	0	0	DELETE	^/api/realmojis/[^/]+$	0.717	0.239	0.314	0.314	0
0.0%	1	0	0	DELETE	^/api/comments/[^/]+$	0.436	0.436	0.436	0.436	0
0.0%	1	0	0	GET	^/api/posts/[^/]+/comments$	0.192	0.192	0.192	0.192	29
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 11030 クエリ / 78 種類 / 合計 46.476s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	47.0%	62	21.848	352.39	2332.53	6360.2	131.2	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	13.5%	18	6.288	349.35	1875.91	32711.1	1649.8	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
3	8.7%	2452	4.065	1.66	6.68	11.9	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	6.9%	2	3.213	1606.42	3055.68	214533.5	2523.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
5	6.7%	432	3.118	7.22	27.75	668.6	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
6	5.2%	1	2.397	2397.34	2397.34	29290.0	1199394.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	3.3%	1	1.514	1514.00	1514.00	19464.0	625552.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	2.4%	1	1.110	1109.81	1109.81	11528.0	529660.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	1.1%	309	0.508	1.64	3.70	36.1	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.7%	2	0.311	155.60	303.29	20234.0	2522.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.5%	70	0.241	3.44	7.27	47.2	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
12	0.5%	1	0.230	230.09	230.09	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
13	0.5%	1	0.218	217.53	217.53	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.4%	1	0.207	206.65	206.65	69891.0	29779.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.3%	1	0.145	145.21	145.21	29337.0	46.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4;
-- rank 3
SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 0e225b64d79314fdf137a3646b056868c2f107c7
Type: cpu
Time: 2026-10-31 10:54:58 JST
Duration: 60.32s, Total samples = 77.43s (128.36%)
Showing nodes accounting for 68.70s, 88.73% of 77.43s total
Dropped 816 nodes (cum <= 0.39s)
Showing top 40 nodes out of 176
      flat  flat%   sum%        cum   cum%
    23.29s 30.08% 30.08%     23.29s 30.08%  crypto/internal/fips140/sha256.blockAVX2
    10.46s 13.51% 43.59%     10.46s 13.51%  runtime.spanClass.sizeclass (inline)
     4.15s  5.36% 48.95%      6.57s  8.49%  image/jpeg.(*decoder).reconstructBlock
     3.60s  4.65% 53.60%     14.77s 19.08%  runtime.tryDeferToSpanScan
     2.30s  2.97% 56.57%      2.30s  2.97%  runtime.memmove
     2.08s  2.69% 59.25%     14.17s 18.30%  image/jpeg.(*decoder).processSOS
     2.03s  2.62% 61.88%      2.03s  2.62%  runtime.memclrNoHeapPointers
     1.77s  2.29% 64.16%      2.75s  3.55%  image/jpeg.(*decoder).decodeHuffman
     1.59s  2.05% 66.21%      6.73s  8.69%  runtime.scanObjectsSmall
     1.33s  1.72% 67.93%      5.34s  6.90%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.23s  1.59% 69.52%      1.23s  1.59%  internal/runtime/syscall/linux.Syscall6
     1.23s  1.59% 71.11%     11.34s 14.65%  runtime.scanObject
     1.02s  1.32% 72.43%      1.02s  1.32%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
        1s  1.29% 73.72%         1s  1.29%  aeshashbody
        1s  1.29% 75.01%      1.09s  1.41%  image/jpeg.(*decoder).receiveExtend
     0.91s  1.18% 76.18%      1.22s  1.58%  image/jpeg.idctCols
     0.91s  1.18% 77.36%      1.16s  1.50%  image/jpeg.idctRows
     0.87s  1.12% 78.48%     24.89s 32.15%  crypto/internal/fips140/sha256.(*Digest).Write
     0.77s  0.99% 79.48%      1.06s  1.37%  image/jpeg.(*decoder).ensureNBits
     0.75s  0.97% 80.45%     17.02s 21.98%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.61s  0.79% 81.23%      0.61s  0.79%  memeqbody
     0.56s  0.72% 81.96%      0.56s  0.72%  image/jpeg.dctBox (inline)
     0.55s  0.71% 82.67%     17.78s 22.96%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.55s  0.71% 83.38%         8s 10.33%  runtime.scanSpan
     0.49s  0.63% 84.01%     18.93s 24.45%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.47s  0.61% 84.62%      0.47s  0.61%  runtime.extractHeapBitsSmall
     0.35s  0.45% 85.07%      0.51s  0.66%  example.com/isureal-company/isureal.postLess
     0.33s  0.43% 85.50%      0.41s  0.53%  runtime.(*spanInlineMarkBits).tryAcquire
     0.32s  0.41% 85.91%      0.92s  1.19%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.32s  0.41% 86.32%      2.09s  2.70%  runtime.mapaccess1_faststr
     0.25s  0.32% 86.65%      0.79s  1.02%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.23s   0.3% 86.94%     20.59s 26.59%  example.com/isureal-company/isureal.imageSignature
     0.22s  0.28% 87.23%      0.65s  0.84%  runtime.scanObjectSmall
     0.21s  0.27% 87.50%     23.51s 30.36%  crypto/internal/fips140/sha256.block
     0.20s  0.26% 87.76%      0.98s  1.27%  runtime.mallocgcSmallScanNoHeader
     0.20s  0.26% 88.01%      1.25s  1.61%  runtime.mapassign_faststr
     0.18s  0.23% 88.25%      0.45s  0.58%  database/sql.convertAssignRows
     0.14s  0.18% 88.43%      2.64s  3.41%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.12s  0.15% 88.58%      0.56s  0.72%  runtime.mallocgcSmallNoscan
     0.11s  0.14% 88.73%     11.30s 14.59%  runtime.mallocgc
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 71.2% / 100% | 60% | 0.3% | 4 | 157 MB | CPU が飽和 |
| isu2 (2 cpu) | 7.9% / 100% | 1% | 1.9% | 2 | 146 MB |  |
| isu3 (2 cpu) | 46.4% / 100% | 33% | 32.5% | 6 | 144 MB | CPU が飽和 |
| isu4 (2 cpu) | 24.8% / 97% | 1% | 16.7% | 2 | 98 MB |  |

**isu1（CPU が飽和、主に isureal）、isu3（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 126% / nginx 10% |
| isu2 | du 5% / postgres 2% |
| isu3 | postgres 9% / image-export 7% / du 5% / kswapd0 1% / kworker/u8:0+flush-259:0 1% |
| isu4 | postgres 14% |

## Diff from previous run

スコア: 24801 → 14462 (-10339)

`cef5fe3..7f1b7c2`

```
 webapp/go/account_deletion_api.go |  1 +
 webapp/go/image_api.go            |  6 ++++
 webapp/go/jpeg_check.go           | 62 +++++++++++++++++++++++++++++++++++++++
 webapp/go/post_api.go             |  5 ++--
 webapp/go/profile_image_api.go    | 11 ++++---
 webapp/go/realmoji_api.go         |  8 +++--
 6 files changed, 81 insertions(+), 12 deletions(-)
```
