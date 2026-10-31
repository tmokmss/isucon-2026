# run 20261031-104641-d3eb91b

## Summary

| | |
|---|---|
| score | 20225 |
| status | pass |
| ref | HEAD |
| sha | d3eb91bbfe86 |
| 所要時間 | 101s |
| 前回比 | -4576 (-18.5%) vs 20261031-103811-cef5fe3 |
| ベスト比 | -4576 (-18.5%) vs 20261031-103811-cef5fe3 = 24801 |
| note | インメモリ化+TLS無効+通知非同期+export主キー |

## Bench

```
{"id":"1586","status":"done","portalStatus":"finished","passed":true,"score":20225,"raw":20225,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:47:03.426Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:47:06.434Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:47:09.410Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:47:12.691Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T01:47:15.593Z level=INFO msg="europe 2 日目: 16 人"
time=2026-10-31T01:47:18.869Z level=INFO msg="america 2 日目: 12 人"
time=2026-10-31T01:47:19.640Z level=INFO msg="user moved" user=nesako_5243 from=asia to=america
time=2026-10-31T01:47:21.553Z level=INFO msg="asia 3 日目: 117 人"
time=2026-10-31T01:47:26.055Z level=INFO msg="europe 3 日目: 26 人"
time=2026-10-31T01:47:28.162Z level=INFO msg="america 3 日目: 14 人"
time=2026-10-31T01:47:31.037Z level=INFO msg="asia 4 日目: 117 人"
time=2026-10-31T01:47:36.595Z level=INFO msg="europe 4 日目: 26 人"
time=2026-10-31T01:47:37.988Z level=INFO msg="america 4 日目: 14 人"
time=2026-10-31T01:47:41.981Z level=INFO msg="asia 5 日目: 117 人"
time=2026-10-31T01:47:45.235Z level=INFO msg="europe 5 日目: 26 人"
time=2026-10-31T01:47:48.949Z level=INFO msg="america 5 日目: 14 人"
time=2026-10-31T01:47:54.879Z level=INFO msg="asia 最終日（グランプリ）: 117 人"
time=2026-10-31T01:47:55.931Z level=INFO msg="america 最終日（グランプリ）: 14 人"
time=2026-10-31T01:47:59.263Z level=INFO msg="europe 最終日（グランプリ）: 25 人"
time=2026-10-31T01:48:03.408Z level=INFO msg="グランプリ: アジア 3 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 20225,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.1%	389	0	0	GET	/api/notifications	3625.356	9.320	13.463	15.039	176
16.8%	442	33	0	POST	/api/posts	1791.252	4.053	10.466	11.111	591
9.9%	3755	856	0	GET	^/api/images/[^/]+	1058.294	0.282	2.755	6.268	105555
9.2%	1980	1118	0	POST	/api/auth/login	976.873	0.493	3.357	6.775	123
5.5%	374	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	580.269	1.552	8.130	10.504	116
5.2%	695	25	0	POST	/api/friend-requests	556.169	0.800	5.478	10.696	112
5.1%	206	153	0	PUT	/api/me/profile-image	540.547	2.624	10.166	10.231	12
4.0%	134	1	0	POST	/api/auth/register	420.335	3.137	12.036	13.237	229
1.9%	2137	0	0	GET	^/api/feeds/[^/]+	201.109	0.094	0.540	4.534	1448
1.3%	937	0	0	GET	/api/memories	139.410	0.149	1.067	2.399	9348
1.3%	559	0	0	GET	/api/posting-window/current	134.193	0.240	0.988	1.859	237
1.0%	63	0	0	GET	/api/me/export	110.807	1.759	9.035	9.035	1986410
0.6%	70	8	0	POST	^/api/posts/[^/]+/realmojis$	63.681	0.910	4.257	4.257	393
0.5%	110	8	0	POST	^/api/posts/[^/]+/comments$	50.379	0.458	2.648	4.079	287
0.5%	74	0	0	PUT	^/api/follows/[^/]+$	48.421	0.654	7.125	7.125	0
0.4%	343	14	0	GET	^/api/posts/me/current$	43.024	0.125	0.687	0.925	614
0.4%	145	0	0	POST	/api/auth/logout	42.146	0.291	3.848	4.593	0
0.3%	375	0	0	GET	^/api/relationships/[^/]+	35.532	0.095	0.359	0.659	1707
0.3%	331	0	0	GET	/api/users	34.257	0.103	0.683	0.776	233
0.3%	37	2	0	PATCH	^/api/posts/[^/]+$	29.991	0.811	5.534	5.534	610
0.2%	349	11	0	GET	^/api/posts/[^/]+$	25.164	0.072	0.319	0.414	608
0.2%	18	0	0	POST	/api/admin/posting-window	24.567	1.365	4.800	4.800	177
0.2%	50	0	0	PATCH	/api/me	19.845	0.397	3.143	3.143	260
0.1%	1	0	0	POST	/api/initialize	12.348	12.348	12.348	12.348	90
0.1%	156	0	0	GET	^/api/users/[^/]+$	11.693	0.075	0.293	0.320	194
0.1%	9	0	0	PUT	^/api/blocks/[^/]+$	11.256	1.251	3.157	3.157	0
0.1%	18	0	0	DELETE	^/api/posts/[^/]+$	10.386	0.577	2.010	2.010	0
0.1%	111	0	0	GET	^/api/posts/[^/]+/realmojis$	6.719	0.061	0.364	0.378	416
0.1%	129	0	0	GET	^/api/memories/[^/]+	6.638	0.051	0.430	0.464	782
0.1%	32	0	0	DELETE	^/api/comments/[^/]+$	6.055	0.189	0.818	0.818	0
…（上位 30 / 全 39 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.8%	88	0	0	GET	/api/notifications	849.833	9.657	14.341	14.341	179
11.6%	89	1	0	POST	/api/posts	253.596	2.849	10.062	10.062	639
9.8%	476	0	0	GET	^/api/images/[^/]+	214.142	0.450	3.119	7.551	180945
6.0%	154	3	0	POST	/api/friend-requests	132.310	0.859	7.371	7.568	113
6.0%	76	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	130.791	1.721	8.703	8.703	116
5.6%	24	0	0	POST	/api/auth/register	121.783	5.074	13.128	13.128	228
4.8%	119	0	0	POST	/api/auth/login	105.309	0.885	4.559	8.565	250
4.7%	355	0	0	GET	^/api/feeds/[^/]+	102.763	0.289	1.193	1.548	1134
2.1%	173	0	0	GET	/api/memories	45.345	0.262	0.871	1.584	10322
1.5%	109	0	0	GET	/api/posting-window/current	33.493	0.307	0.972	0.973	239
1.2%	11	0	0	GET	/api/me/export	27.317	2.483	8.186	8.186	2091376
1.0%	12	0	0	PUT	^/api/follows/[^/]+$	22.610	1.884	6.565	6.565	0
0.9%	80	0	0	GET	^/api/relationships/[^/]+	19.721	0.247	1.232	1.232	1418
0.9%	74	5	0	GET	^/api/posts/me/current$	19.101	0.258	0.742	0.742	601
0.8%	68	0	0	GET	/api/users	16.782	0.247	1.163	1.163	228
0.7%	13	0	0	PUT	/api/me/profile-image	14.377	1.106	2.696	2.696	0
0.6%	60	6	0	GET	^/api/posts/[^/]+$	14.139	0.236	0.876	0.876	544
0.6%	6	0	0	POST	^/api/posts/[^/]+/realmojis$	13.391	2.232	4.240	4.240	503
0.4%	9	0	0	PATCH	/api/me	9.715	1.079	3.268	3.268	274
0.4%	15	0	0	POST	/api/auth/logout	8.246	0.550	2.200	2.200	0
0.4%	13	2	0	POST	^/api/posts/[^/]+/comments$	8.075	0.621	3.024	3.024	273
0.2%	7	1	0	PATCH	^/api/posts/[^/]+$	4.264	0.609	1.117	1.117	536
0.2%	24	0	0	GET	^/api/users/[^/]+$	4.195	0.175	0.439	0.439	178
0.2%	18	2	0	GET	/api/me	3.941	0.219	0.521	0.521	256
0.2%	15	0	0	GET	^/api/memories/[^/]+	3.797	0.253	0.647	0.647	658
0.2%	17	0	0	GET	^/api/posts/[^/]+/realmojis$	3.351	0.197	0.821	0.821	156
0.1%	18	0	0	GET	^/api/posts/[^/]+/comments$	2.469	0.137	0.245	0.245	294
0.1%	2	0	0	DELETE	^/api/follows/[^/]+$	1.821	0.910	1.684	1.684	0
0.0%	1	0	0	POST	/api/me/account-deletion	0.902	0.902	0.902	0.902	0
0.0%	1	0	0	DELETE	^/api/comments/[^/]+$	0.148	0.148	0.148	0.148	0
…（上位 30 / 全 31 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.7%	65	0	0	GET	/api/notifications	556.110	8.556	10.891	10.891	182
16.2%	59	0	0	POST	/api/posts	206.322	3.497	9.650	9.650	616
11.2%	279	2	0	GET	^/api/images/[^/]+	142.397	0.510	3.992	5.627	167778
5.9%	34	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	74.724	2.198	10.501	10.501	116
4.4%	70	2	0	POST	/api/friend-requests	56.331	0.805	4.187	4.187	113
3.3%	165	0	0	GET	^/api/feeds/[^/]+	42.285	0.256	0.922	1.133	973
2.9%	65	0	0	POST	/api/auth/login	36.616	0.563	3.422	3.422	250
2.8%	108	0	0	GET	/api/memories	36.200	0.335	1.546	5.159	10229
1.6%	71	0	0	GET	/api/posting-window/current	20.469	0.288	1.676	1.676	240
1.1%	12	0	0	POST	/api/auth/register	14.568	1.214	11.280	11.280	228
1.1%	6	0	0	GET	/api/me/export	14.202	2.367	4.521	4.521	3405374
1.1%	49	4	0	GET	^/api/posts/me/current$	13.946	0.285	1.753	1.753	572
0.7%	43	0	0	GET	^/api/relationships/[^/]+	8.302	0.193	0.526	0.526	1399
0.6%	8	1	0	PATCH	^/api/posts/[^/]+$	8.217	1.027	2.773	2.773	565
0.6%	31	0	0	GET	/api/users	7.311	0.236	0.555	0.555	203
0.5%	12	0	0	POST	^/api/posts/[^/]+/comments$	6.013	0.501	1.644	1.644	300
0.4%	4	0	0	PUT	/api/me/profile-image	5.413	1.353	2.872	2.872	0
0.3%	20	3	0	GET	^/api/posts/[^/]+$	3.557	0.178	0.321	0.321	510
0.3%	13	0	0	PUT	^/api/follows/[^/]+$	3.385	0.260	1.293	1.293	0
0.3%	2	0	0	DELETE	^/api/posts/[^/]+$	3.287	1.643	3.088	3.088	0
0.2%	5	0	0	PATCH	/api/me	3.174	0.635	1.698	1.698	238
0.2%	17	0	0	GET	^/api/users/[^/]+$	2.966	0.174	0.284	0.284	168
0.2%	7	0	0	GET	/api/me	2.536	0.362	1.434	1.434	262
0.2%	8	0	0	POST	/api/auth/logout	2.271	0.284	1.253	1.253	0
0.1%	8	0	0	GET	^/api/memories/[^/]+	0.996	0.124	0.163	0.163	1152
0.1%	2	0	0	DELETE	^/api/comments/[^/]+$	0.781	0.390	0.420	0.420	0
0.1%	5	0	0	GET	^/api/posts/[^/]+/comments$	0.727	0.145	0.211	0.211	300
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.107	0.107	0.107	0.107	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 14010 クエリ / 81 種類 / 合計 71.505s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	46.4%	80	33.206	415.08	4938.48	1347.4	134.2	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.7%	16	10.541	658.79	2760.87	30381.3	1555.5	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
3	9.9%	699	7.107	10.17	47.12	753.4	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
4	6.3%	1	4.533	4533.34	4533.34	1845286.0	1.0	SELECT count(*), md5(string_agg(id || chr($1) || owner_user_id || chr($2) || media_type || chr($3) || created_at::text, chr($4) ORDER BY id)) FROM image_objects…
5	5.2%	2575	3.709	1.44	9.68	12.5	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
6	3.2%	1	2.270	2269.86	2269.86	29270.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	2.8%	1	2.005	2004.70	2004.70	19464.0	625284.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	2.2%	2	1.548	773.99	1532.75	48772.0	2503.5	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
9	1.3%	561	0.926	1.65	7.11	34.0	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	1.2%	1	0.879	878.82	878.82	34934.0	1.0	SELECT count(*), md5(string_agg(id || chr($1) || md5(data), chr($2) ORDER BY id)) FROM image_objects WHERE id LIKE chr($3)||chr($4)||chr($5)||chr($6) AND create…
11	1.2%	1	0.860	860.07	860.07	11528.0	529666.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
12	0.9%	2	0.639	319.34	345.90	211813.5	2491.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
13	0.8%	1	0.592	592.32	592.32	19811.0	173.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.6%	2	0.436	217.88	423.92	19413.5	2463.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	0.4%	1	0.305	304.91	304.91	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4;
-- rank 3
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: ef1183473a32c4c04bd7c96600bfcda3257cba68
Type: cpu
Time: 2026-10-31 10:46:46 JST
Duration: 60.31s, Total samples = 86.30s (143.10%)
Showing nodes accounting for 76.58s, 88.74% of 86.30s total
Dropped 820 nodes (cum <= 0.43s)
Showing top 40 nodes out of 164
      flat  flat%   sum%        cum   cum%
    15.29s 17.72% 17.72%     15.29s 17.72%  crypto/internal/fips140/sha256.blockAVX2
    11.27s 13.06% 30.78%     11.27s 13.06%  runtime.spanClass.sizeclass (inline)
     8.86s 10.27% 41.04%     13.91s 16.12%  image/jpeg.(*decoder).reconstructBlock
     3.96s  4.59% 45.63%     30.59s 35.45%  image/jpeg.(*decoder).processSOS
     3.95s  4.58% 50.21%     15.85s 18.37%  runtime.tryDeferToSpanScan
     3.82s  4.43% 54.63%      5.94s  6.88%  image/jpeg.(*decoder).decodeHuffman
     2.74s  3.17% 57.81%      2.88s  3.34%  image/jpeg.(*decoder).receiveExtend
     2.34s  2.71% 60.52%      2.34s  2.71%  runtime.memclrNoHeapPointers
     2.30s  2.67% 63.19%      2.30s  2.67%  runtime.memmove
     1.92s  2.22% 65.41%      2.57s  2.98%  image/jpeg.idctCols
     1.78s  2.06% 67.47%      2.37s  2.75%  image/jpeg.idctRows
     1.78s  2.06% 69.54%      7.18s  8.32%  runtime.scanObjectsSmall
     1.66s  1.92% 71.46%      2.24s  2.60%  image/jpeg.(*decoder).ensureNBits
     1.53s  1.77% 73.23%      1.53s  1.77%  internal/runtime/syscall/linux.Syscall6
     1.37s  1.59% 74.82%     12.74s 14.76%  runtime.scanObject
     1.23s  1.43% 76.25%      1.23s  1.43%  image/jpeg.dctBox (inline)
     1.19s  1.38% 77.62%      5.33s  6.18%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.03s  1.19% 78.82%      1.03s  1.19%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.92s  1.07% 79.88%     17.36s 20.12%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.82s  0.95% 80.83%      0.82s  0.95%  memeqbody
     0.77s  0.89% 81.73%     16.92s 19.61%  crypto/internal/fips140/sha256.(*Digest).Write
     0.69s   0.8% 82.53%     18.31s 21.22%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.57s  0.66% 83.19%      0.57s  0.66%  aeshashbody
     0.57s  0.66% 83.85%      8.37s  9.70%  runtime.scanSpan
     0.55s  0.64% 84.48%      0.57s  0.66%  image/jpeg.(*decoder).readByteStuffedByte
     0.45s  0.52% 85.01%      2.09s  2.42%  runtime.mapaccess1_faststr
     0.44s  0.51% 85.52%     19.52s 22.62%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.39s  0.45% 85.97%      0.59s  0.68%  example.com/isureal-company/isureal.postLess
     0.32s  0.37% 86.34%      0.87s  1.01%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.29s  0.34% 86.67%      0.93s  1.08%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.25s  0.29% 86.96%      1.11s  1.29%  runtime.mallocgcSmallScanNoHeader
     0.22s  0.25% 87.22%     14.48s 16.78%  runtime.mallocgc
     0.22s  0.25% 87.47%      0.54s  0.63%  runtime.scanObjectSmall
     0.21s  0.24% 87.72%      0.47s  0.54%  database/sql.convertAssignRows
     0.21s  0.24% 87.96%     20.99s 24.32%  example.com/isureal-company/isureal.imageSignature
     0.19s  0.22% 88.18%     15.49s 17.95%  crypto/internal/fips140/sha256.block
     0.17s   0.2% 88.38%      1.16s  1.34%  runtime.mapassign_faststr
     0.11s  0.13% 88.51%      2.50s  2.90%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s  0.13% 88.63%      0.49s  0.57%  runtime.(*sweepLocked).sweep
     0.09s   0.1% 88.74%      5.03s  5.83%  image/jpeg.idct
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 74.1% / 100% | 64% | 0.0% | 6 | 108 MB | CPU が飽和 |
| isu2 (2 cpu) | 86.3% / 100% | 82% | 71.3% | 7 | 141 MB | CPU が飽和 |
| isu3 (2 cpu) | 0.8% / 15% | 0% | 0.0% | 1 | 3216 MB |  |
| isu4 (2 cpu) | 32.3% / 91% | 1% | 19.6% | 4 | 114 MB |  |

**isu1（CPU が飽和、主に isureal）、isu2（CPU が飽和、主に image-export）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 132% / nginx 10% |
| isu2 | image-export 13% / postgres 12% / kworker/u8:1+flush-259:0 1% |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 23% |

## Diff from previous run

スコア: 24801 → 20225 (-4576)

`cef5fe3..d3eb91b`

```
 webapp/go/account_deletion_api.go     |  1 +
 webapp/go/admin_posting_window_api.go | 17 +++++++++++------
 webapp/go/config.go                   |  2 ++
 webapp/go/export_api.go               |  4 +++-
 webapp/go/image_api.go                |  6 ++++++
 webapp/go/post_api.go                 |  1 +
 webapp/go/profile_image_api.go        |  4 ++++
 webapp/go/realmoji_api.go             |  4 ++++
 8 files changed, 32 insertions(+), 7 deletions(-)
```
