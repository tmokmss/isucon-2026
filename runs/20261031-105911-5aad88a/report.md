# run 20261031-105911-5aad88a

## Summary

| | |
|---|---|
| score | 10532 |
| status | pass |
| ref | origin/inmemory-2 |
| sha | 5aad88a3aaab |
| 所要時間 | 208s |
| 前回比 | -3677 (-25.9%) vs 20261031-105734-4350e7f |
| ベスト比 | -14269 (-57.5%) vs 20261031-103811-cef5fe3 = 24801 |
| note | PR #63: TLS 無効化・エクスポートを主キーで・メモリ削減 |

## Bench

```
{"id":"2354","status":"done","portalStatus":"finished","passed":true,"score":10532,"raw":10532,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:01:26.429Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:01:28.812Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:01:31.787Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:01:35.575Z level=INFO msg="asia 2 日目: 38 人"
time=2026-10-31T02:01:38.057Z level=INFO msg="europe 2 日目: 14 人"
time=2026-10-31T02:01:40.845Z level=INFO msg="america 2 日目: 10 人"
time=2026-10-31T02:01:45.267Z level=INFO msg="asia 3 日目: 79 人"
time=2026-10-31T02:01:47.416Z level=INFO msg="europe 3 日目: 22 人"
time=2026-10-31T02:01:51.823Z level=INFO msg="america 3 日目: 16 人"
time=2026-10-31T02:01:55.684Z level=INFO msg="asia 4 日目: 79 人"
time=2026-10-31T02:02:00.919Z level=INFO msg="europe 4 日目: 22 人"
time=2026-10-31T02:02:03.799Z level=INFO msg="america 4 日目: 16 人"
time=2026-10-31T02:02:08.727Z level=INFO msg="asia 5 日目: 79 人"
time=2026-10-31T02:02:10.927Z level=INFO msg="europe 5 日目: 22 人"
time=2026-10-31T02:02:25.278Z level=INFO msg="asia 最終日（グランプリ）: 79 人"
time=2026-10-31T02:02:25.683Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 10532,
  "messages": [
    "america の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました",
    "europe の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.3%	295	0	0	GET	/api/notifications	3330.608	11.290	16.073	20.778	178
12.3%	1936	1340	0	POST	/api/auth/login	1015.353	0.524	3.699	5.966	95
11.7%	216	0	0	POST	/api/posts	966.030	4.472	13.776	14.038	646
10.2%	2604	728	0	GET	^/api/images/[^/]+	841.342	0.323	2.687	6.430	74728
5.6%	183	144	0	PUT	/api/me/profile-image	460.732	2.518	10.412	10.496	23
5.1%	411	18	0	POST	/api/friend-requests	422.451	1.028	9.807	11.721	111
4.6%	163	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	382.480	2.347	10.734	14.219	116
1.6%	903	0	0	GET	^/api/feeds/[^/]+	133.678	0.148	1.594	2.233	919
1.6%	624	0	0	GET	/api/memories	132.949	0.213	2.681	3.288	10379
1.1%	299	0	0	GET	/api/posting-window/current	89.553	0.300	1.937	2.209	236
0.9%	18	0	0	POST	/api/admin/posting-window	72.689	4.038	10.746	10.746	176
0.7%	95	1	0	POST	/api/auth/register	60.571	0.638	6.728	6.728	230
0.7%	212	36	0	GET	^/api/posts/me/current$	54.876	0.259	1.644	1.646	540
0.6%	46	0	0	GET	/api/me/export	48.222	1.048	7.676	7.676	2026348
0.5%	221	0	0	GET	^/api/relationships/[^/]+	41.072	0.186	1.926	2.353	1675
0.4%	193	0	0	GET	/api/users	32.536	0.169	1.917	2.725	227
0.4%	22	8	0	PATCH	^/api/posts/[^/]+$	32.043	1.456	14.187	14.187	428
0.3%	48	0	0	PUT	^/api/follows/[^/]+$	27.499	0.573	4.001	4.001	0
0.3%	42	0	0	POST	^/api/posts/[^/]+/realmojis$	24.277	0.578	4.065	4.065	423
0.3%	90	0	0	POST	/api/auth/logout	21.796	0.242	5.717	5.717	0
0.2%	1	0	0	POST	/api/initialize	12.614	12.614	12.614	12.614	90
0.1%	122	7	0	GET	^/api/posts/[^/]+$	11.276	0.092	0.413	0.734	565
0.1%	11	0	0	DELETE	^/api/posts/[^/]+$	8.643	0.786	3.920	3.920	0
0.1%	26	0	0	PATCH	/api/me	7.897	0.304	2.264	2.264	284
0.1%	89	0	0	GET	^/api/memories/[^/]+	5.736	0.064	1.086	1.086	760
0.1%	53	0	0	GET	^/api/users/[^/]+$	4.669	0.088	0.304	0.304	173
0.0%	2	0	0	DELETE	^/api/follows/[^/]+$	3.338	1.669	3.143	3.143	0
0.0%	7	0	0	POST	^/api/posts/[^/]+/comments$	2.915	0.416	2.244	2.244	389
0.0%	45	2	0	GET	/api/me	2.826	0.063	0.332	0.332	280
0.0%	9	0	0	DELETE	^/api/realmojis/[^/]+$	2.774	0.308	0.798	0.798	0
…（上位 30 / 全 38 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
46.8%	85	0	0	GET	/api/notifications	926.465	10.900	16.306	16.306	181
14.1%	65	0	0	POST	/api/posts	278.767	4.289	11.278	11.278	607
6.7%	121	7	0	POST	/api/friend-requests	132.077	1.092	9.791	10.296	110
6.6%	49	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	130.488	2.663	12.345	12.345	116
6.2%	271	2	0	GET	^/api/images/[^/]+	122.262	0.451	2.372	2.713	169347
4.2%	92	0	0	POST	/api/auth/login	82.155	0.893	8.799	8.799	247
2.5%	134	0	0	GET	/api/memories	49.815	0.372	2.509	2.768	5691
2.2%	8	0	0	GET	/api/me/export	44.372	5.547	15.105	15.105	1796335
2.2%	145	0	0	GET	^/api/feeds/[^/]+	43.308	0.299	1.683	2.000	797
2.2%	89	0	0	GET	/api/posting-window/current	43.091	0.484	2.979	2.979	239
1.2%	60	14	0	GET	^/api/posts/me/current$	22.834	0.381	1.896	1.896	468
0.8%	58	0	0	GET	/api/users	15.528	0.268	2.529	2.529	240
0.7%	20	0	0	POST	/api/auth/register	14.001	0.700	2.282	2.282	228
0.7%	55	0	0	GET	^/api/relationships/[^/]+	13.925	0.253	1.405	1.405	1338
0.7%	12	0	0	PUT	^/api/follows/[^/]+$	13.812	1.151	5.062	5.062	0
0.5%	14	0	0	POST	/api/auth/logout	9.992	0.714	4.409	4.409	0
0.4%	6	0	0	PUT	/api/me/profile-image	8.611	1.435	2.164	2.164	0
0.2%	10	0	0	POST	^/api/posts/[^/]+/realmojis$	4.678	0.468	1.284	1.284	398
0.2%	3	0	0	DELETE	^/api/posts/[^/]+$	4.598	1.533	2.372	2.372	0
0.2%	13	0	0	GET	^/api/posts/[^/]+$	3.818	0.294	1.611	1.611	601
0.1%	6	2	0	PATCH	^/api/posts/[^/]+$	2.911	0.485	0.902	0.902	406
0.1%	10	1	0	GET	/api/me	2.622	0.262	0.497	0.497	243
0.1%	14	0	0	GET	^/api/memories/[^/]+	2.217	0.158	0.382	0.382	780
0.1%	4	0	0	PATCH	/api/me	1.982	0.495	0.855	0.855	246
0.1%	4	0	0	POST	^/api/posts/[^/]+/comments$	1.267	0.317	0.639	0.639	278
0.0%	2	0	0	DELETE	^/api/comments/[^/]+$	0.834	0.417	0.659	0.659	0
0.0%	5	0	0	GET	^/api/users/[^/]+$	0.640	0.128	0.182	0.182	162
0.0%	2	0	0	DELETE	^/api/realmojis/[^/]+$	0.528	0.264	0.305	0.305	0
0.0%	3	0	0	GET	^/api/posts/[^/]+/comments$	0.425	0.142	0.157	0.157	211
0.0%	3	0	0	GET	^/api/posts/[^/]+/realmojis$	0.367	0.122	0.144	0.144	425
…（上位 30 / 全 31 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
49.0%	59	0	0	GET	/api/notifications	612.564	10.382	19.422	19.422	176
11.9%	34	0	0	POST	/api/posts	148.525	4.368	11.309	11.309	623
9.3%	217	3	0	GET	^/api/images/[^/]+	116.752	0.538	3.509	4.095	133887
4.7%	69	0	0	POST	/api/auth/login	58.284	0.845	5.213	5.213	248
4.0%	50	1	0	POST	/api/friend-requests	50.506	1.010	9.741	9.741	113
3.5%	97	0	0	GET	/api/memories	43.164	0.445	9.879	9.879	9867
3.3%	15	0	0	PUT	^/api/follows/[^/]+$	40.787	2.719	10.044	10.044	0
2.2%	30	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	27.155	0.905	5.861	5.861	116
1.8%	79	0	0	GET	^/api/feeds/[^/]+	22.946	0.290	1.412	1.412	678
1.7%	53	0	0	GET	/api/posting-window/current	21.516	0.406	1.335	1.335	240
1.4%	31	2	0	GET	^/api/posts/me/current$	17.662	0.570	2.179	2.179	576
1.2%	5	0	0	DELETE	^/api/posts/[^/]+$	14.986	2.997	5.294	5.294	0
1.1%	5	0	0	PUT	/api/me/profile-image	13.875	2.775	7.021	7.021	0
1.0%	6	0	0	GET	/api/me/export	12.014	2.002	5.561	5.561	2591626
0.7%	15	0	0	POST	/api/auth/register	8.809	0.587	1.413	1.413	229
0.6%	13	0	0	POST	/api/auth/logout	7.829	0.602	4.065	4.065	0
0.6%	7	0	0	PATCH	/api/me	7.459	1.066	3.395	3.395	285
0.5%	36	0	0	GET	^/api/relationships/[^/]+	6.852	0.190	0.613	0.613	1125
0.4%	23	0	0	GET	/api/users	4.684	0.204	0.586	0.586	255
0.3%	8	0	0	POST	^/api/posts/[^/]+/realmojis$	3.794	0.474	1.208	1.208	404
0.2%	8	0	0	GET	/api/me	3.034	0.379	1.218	1.218	263
0.1%	2	0	0	POST	^/api/posts/[^/]+/comments$	1.598	0.799	0.907	0.907	408
0.1%	5	1	0	GET	^/api/posts/[^/]+$	1.516	0.303	0.845	0.845	469
0.1%	8	0	0	GET	^/api/memories/[^/]+	1.054	0.132	0.194	0.194	517
0.1%	3	0	0	PATCH	^/api/posts/[^/]+$	0.745	0.248	0.356	0.356	581
0.0%	2	0	0	GET	^/api/posts/[^/]+/realmojis$	0.324	0.162	0.185	0.185	230
0.0%	2	0	0	GET	^/api/posts/[^/]+/comments$	0.251	0.126	0.136	0.136	29
0.0%	1	0	0	DELETE	^/api/friendships/[^/]+$	0.238	0.238	0.238	0.238	0
0.0%	1	0	0	GET	^/api/users/[^/]+$	0.213	0.213	0.213	0.213	179
0.0%	1	0	0	DELETE	^/api/realmojis/[^/]+$	0.163	0.163	0.163	0.163	0
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 9240 クエリ / 75 種類 / 合計 47.771s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	56.8%	60	27.154	452.57	4295.92	1331.8	130.4	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	12.3%	18	5.859	325.51	1432.28	32327.4	1655.2	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
3	7.2%	425	3.431	8.07	39.11	697.9	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
4	6.5%	1781	3.096	1.74	12.16	12.5	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	4.9%	1	2.330	2330.38	2330.38	29308.0	1199411.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	3.3%	1	1.561	1561.32	1561.32	19504.0	625523.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	2.2%	1	1.058	1058.46	1058.46	11528.0	529676.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	1.1%	315	0.531	1.69	4.69	38.7	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
9	0.8%	2	0.364	181.83	213.45	215128.5	2530.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
10	0.6%	2	0.302	151.17	294.61	20231.0	2526.5	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.5%	1	0.235	235.29	235.29	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
12	0.4%	1	0.214	213.96	213.96	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
13	0.4%	1	0.207	206.76	206.76	70008.0	29896.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.4%	60	0.180	3.01	5.20	61.0	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
15	0.3%	1	0.143	143.38	143.38	29372.0	63.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: 86a462d96bb24764fb76e2fc6b8a530d0a31b701
Type: cpu
Time: 2026-10-31 11:01:09 JST
Duration: 60.39s, Total samples = 86.39s (143.06%)
Showing nodes accounting for 76.55s, 88.61% of 86.39s total
Dropped 794 nodes (cum <= 0.43s)
Showing top 40 nodes out of 174
      flat  flat%   sum%        cum   cum%
    19.15s 22.17% 22.17%     19.15s 22.17%  runtime.spanClass.sizeclass (inline)
    12.31s 14.25% 36.42%     12.32s 14.26%  crypto/internal/fips140/sha256.blockAVX2
     5.58s  6.46% 42.88%      8.79s 10.17%  image/jpeg.(*decoder).reconstructBlock
     5.29s  6.12% 49.00%     25.37s 29.37%  runtime.tryDeferToSpanScan
     3.28s  3.80% 52.80%      3.28s  3.80%  runtime.memclrNoHeapPointers
     2.59s  3.00% 55.79%     11.32s 13.10%  runtime.scanObjectsSmall
     2.47s  2.86% 58.65%     19.18s 22.20%  image/jpeg.(*decoder).processSOS
     2.36s  2.73% 61.38%      2.36s  2.73%  runtime.memmove
     2.21s  2.56% 63.94%     19.99s 23.14%  runtime.scanObject
     2.17s  2.51% 66.45%      3.65s  4.23%  image/jpeg.(*decoder).decodeHuffman
     1.77s  2.05% 68.50%      6.19s  7.17%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.77s  2.05% 70.55%      1.83s  2.12%  image/jpeg.(*decoder).receiveExtend
     1.35s  1.56% 72.11%      1.35s  1.56%  internal/runtime/syscall/linux.Syscall6
     1.23s  1.42% 73.54%      1.57s  1.82%  image/jpeg.idctCols
     1.17s  1.35% 74.89%      1.60s  1.85%  image/jpeg.idctRows
     1.07s  1.24% 76.13%      1.53s  1.77%  image/jpeg.(*decoder).ensureNBits
     1.06s  1.23% 77.36%     13.49s 15.62%  runtime.scanSpan
     1.05s  1.22% 78.57%      1.05s  1.22%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.99s  1.15% 79.72%      0.99s  1.15%  aeshashbody
     0.76s  0.88% 80.60%      0.76s  0.88%  image/jpeg.dctBox (inline)
     0.71s  0.82% 81.42%      0.71s  0.82%  memeqbody
     0.62s  0.72% 82.14%     13.74s 15.90%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.61s  0.71% 82.85%      0.61s  0.71%  runtime.extractHeapBitsSmall
     0.58s  0.67% 83.52%     13.46s 15.58%  crypto/internal/fips140/sha256.(*Digest).Write
     0.47s  0.54% 84.06%      0.48s  0.56%  runtime.(*spanInlineMarkBits).tryAcquire
     0.46s  0.53% 84.59%     14.46s 16.74%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.42s  0.49% 85.08%      0.47s  0.54%  runtime.findObject
     0.34s  0.39% 85.47%     15.36s 17.78%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.33s  0.38% 85.85%      0.45s  0.52%  image/jpeg.(*decoder).readByteStuffedByte
     0.32s  0.37% 86.23%      0.62s  0.72%  example.com/isureal-company/isureal.postLess
     0.31s  0.36% 86.58%      2.04s  2.36%  runtime.mapaccess1_faststr
     0.28s  0.32% 86.91%      0.86s     1%  database/sql.convertAssignRows
     0.27s  0.31% 87.22%      0.70s  0.81%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.25s  0.29% 87.51%      0.95s  1.10%  runtime.scanObjectSmall
     0.19s  0.22% 87.73%     12.51s 14.48%  crypto/internal/fips140/sha256.block
     0.18s  0.21% 87.94%      0.78s   0.9%  runtime.sweepone
     0.16s  0.19% 88.12%     16.48s 19.08%  example.com/isureal-company/isureal.imageSignature
     0.14s  0.16% 88.29%      0.52s   0.6%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.14s  0.16% 88.45%      1.15s  1.33%  runtime.mallocgcSmallScanNoHeader
     0.14s  0.16% 88.61%      1.57s  1.82%  runtime.mapassign_faststr
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 77.5% / 100% | 66% | 0.2% | 7 | 113 MB | CPU が飽和 |
| isu2 (2 cpu) | 4.8% / 100% | 1% | 0.2% | 3 | 193 MB |  |
| isu3 (2 cpu) | 2.0% / 42% | 0% | 0.6% | 1 | 186 MB |  |
| isu4 (2 cpu) | 22.2% / 91% | 1% | 14.5% | 2 | 107 MB |  |

**isu1（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 135% / nginx 11% |
| isu2 | kworker/u8:6+events_unbound 3% / kworker/1:0+inode_switch_wbs 2% |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 13% |

## Diff from previous run

スコア: 14209 → 10532 (-3677)

`4350e7f..5aad88a`

```
 conf/isu1/etc/systemd/system/isureal-go.service |  2 +-
 webapp/go/config.go                             |  2 ++
 webapp/go/export_api.go                         |  4 +++-
 webapp/go/store.go                              | 11 +++++++++++
 4 files changed, 17 insertions(+), 2 deletions(-)
```
