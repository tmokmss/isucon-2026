# run 20261031-104710-4350e7f

## Summary

| | |
|---|---|
| score | 17610 |
| status | pass |
| ref | HEAD |
| sha | 4350e7f0be16 |
| 所要時間 | 339s |
| 前回比 | -7191 (-29.0%) vs 20261031-103811-cef5fe3 |
| ベスト比 | -7191 (-29.0%) vs 20261031-103811-cef5fe3 = 24801 |
| note | main: インメモリ化（#61）を入れた |

## Bench

```
{"id":"1802","status":"done","portalStatus":"finished","passed":true,"score":17610,"raw":17610,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:51:17.713Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:51:18.611Z level=INFO msg="user moved" user=riku_5030 from=asia to=america
time=2026-10-31T01:51:19.968Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:51:22.553Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T01:51:26.615Z level=INFO msg="asia 2 日目: 29 人"
time=2026-10-31T01:51:28.960Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T01:51:29.192Z level=INFO msg="user moved" user=mamo.5142 from=america to=europe
time=2026-10-31T01:51:31.562Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T01:51:36.294Z level=INFO msg="asia 3 日目: 65 人"
time=2026-10-31T01:51:38.370Z level=INFO msg="europe 3 日目: 8 人"
time=2026-10-31T01:51:40.916Z level=INFO msg="america 3 日目: 25 人"
time=2026-10-31T01:51:44.845Z level=INFO msg="asia 4 日目: 103 人"
time=2026-10-31T01:51:50.189Z level=INFO msg="europe 4 日目: 8 人"
time=2026-10-31T01:51:50.737Z level=INFO msg="america 4 日目: 25 人"
time=2026-10-31T01:51:56.223Z level=INFO msg="asia 5 日目: 103 人"
time=2026-10-31T01:51:57.881Z level=INFO msg="europe 5 日目: 8 人"
time=2026-10-31T01:52:03.345Z level=INFO msg="america 5 日目: 25 人"
time=2026-10-31T01:52:10.583Z level=INFO msg="europe 最終日（グランプリ）: 8 人"
time=2026-10-31T01:52:12.198Z level=INFO msg="asia 最終日（グランプリ）: 102 人"
time=2026-10-31T01:52:14.074Z level=INFO msg="america 最終日（グランプリ）: 24 人"
time=2026-10-31T01:52:16.538Z level=INFO msg="グランプリ: アジア 1 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 17610,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.0%	305	0	0	GET	/api/notifications	3289.427	10.785	18.706	19.693	177
15.5%	354	5	0	POST	/api/posts	1345.607	3.801	10.947	12.833	626
9.0%	2197	1209	0	POST	/api/auth/login	782.569	0.356	2.838	6.358	127
8.9%	3843	649	0	GET	^/api/images/[^/]+	773.394	0.201	2.224	6.051	92151
5.2%	188	141	0	PUT	/api/me/profile-image	449.086	2.389	10.097	10.615	19
4.9%	355	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	420.727	1.185	8.202	11.191	116
4.0%	626	24	0	POST	/api/friend-requests	344.765	0.551	5.481	9.361	112
3.0%	1993	0	0	GET	^/api/feeds/[^/]+	263.652	0.132	1.187	2.343	1803
2.4%	120	1	0	POST	/api/auth/register	206.514	1.721	6.483	6.637	229
2.0%	967	0	0	GET	/api/memories	172.086	0.178	1.955	2.451	11136
1.3%	457	0	0	GET	/api/posting-window/current	116.268	0.254	1.933	2.173	237
0.7%	63	0	0	GET	/api/me/export	64.444	1.023	9.833	9.833	1906739
0.7%	69	0	0	POST	^/api/posts/[^/]+/realmojis$	56.971	0.826	4.458	4.458	467
0.7%	282	14	0	GET	^/api/posts/me/current$	56.475	0.200	1.953	2.373	602
0.6%	308	0	0	GET	^/api/relationships/[^/]+	47.735	0.155	0.784	1.427	1788
0.4%	18	0	0	POST	/api/admin/posting-window	36.945	2.053	6.504	6.504	176
0.4%	70	0	0	PUT	^/api/follows/[^/]+$	34.210	0.489	4.325	4.325	0
0.4%	308	0	0	GET	/api/users	33.341	0.108	0.743	2.075	229
0.3%	353	8	0	GET	^/api/posts/[^/]+$	29.848	0.085	0.482	1.010	590
0.3%	174	0	0	POST	/api/auth/logout	21.766	0.125	2.801	2.909	0
0.2%	170	0	0	GET	^/api/users/[^/]+$	17.905	0.105	1.586	1.681	181
0.1%	36	4	0	PATCH	^/api/posts/[^/]+$	11.554	0.321	2.222	2.222	568
0.1%	16	0	0	DELETE	^/api/posts/[^/]+$	10.576	0.661	4.131	4.131	0
0.1%	7	0	0	PUT	^/api/blocks/[^/]+$	10.301	1.472	5.100	5.100	0
0.1%	26	1	0	POST	^/api/posts/[^/]+/comments$	9.163	0.352	1.372	1.372	305
0.1%	35	0	0	PATCH	/api/me	8.449	0.241	0.696	0.696	273
0.1%	1	0	0	POST	/api/me/account-deletion	5.470	5.470	5.470	5.470	0
0.1%	55	3	0	GET	/api/me	5.090	0.093	0.551	0.551	277
0.1%	158	0	0	GET	^/api/memories/[^/]+	4.847	0.031	0.509	0.549	488
0.1%	1	0	0	POST	/api/initialize	4.619	4.619	4.619	4.619	90
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
49.2%	47	0	0	GET	/api/notifications	429.319	9.134	12.411	12.411	177
16.6%	39	1	0	POST	/api/posts	144.928	3.716	10.678	10.678	592
11.7%	134	2	0	GET	^/api/images/[^/]+	101.949	0.761	3.387	3.407	192109
4.1%	71	0	0	GET	/api/memories	35.550	0.501	2.733	2.733	11433
3.6%	59	0	0	GET	^/api/feeds/[^/]+	31.361	0.532	1.330	1.330	887
2.8%	40	0	0	GET	/api/posting-window/current	24.315	0.608	2.118	2.118	239
2.7%	29	0	0	POST	/api/auth/login	23.788	0.820	3.420	3.420	249
2.4%	29	0	0	POST	/api/friend-requests	20.719	0.714	2.608	2.608	115
1.7%	33	6	0	GET	^/api/posts/me/current$	14.865	0.450	2.035	2.035	506
1.4%	10	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	12.114	1.211	5.205	5.205	116
0.9%	17	0	0	GET	/api/users	7.773	0.457	1.131	1.131	216
0.7%	1	0	0	POST	/api/me/account-deletion	6.426	6.426	6.426	6.426	0
0.5%	20	0	0	GET	^/api/relationships/[^/]+	4.569	0.228	0.986	0.986	1323
0.3%	2	1	0	PATCH	^/api/posts/[^/]+$	2.870	1.435	2.569	2.569	358
0.3%	2	0	0	GET	/api/me/export	2.338	1.169	1.721	1.721	3178587
0.3%	3	0	0	POST	/api/auth/logout	2.223	0.741	1.165	1.165	0
0.2%	4	0	0	PUT	^/api/follows/[^/]+$	1.847	0.462	0.582	0.582	0
0.2%	5	0	0	GET	^/api/memories/[^/]+	1.811	0.362	1.259	1.259	945
0.2%	4	0	0	GET	^/api/posts/[^/]+$	1.312	0.328	0.515	0.515	588
0.1%	5	0	0	POST	/api/auth/register	1.092	0.218	0.224	0.224	231
0.1%	2	0	0	DELETE	^/api/posts/[^/]+$	0.686	0.343	0.547	0.547	0
0.1%	4	1	0	GET	/api/me	0.507	0.127	0.202	0.202	153
0.0%	1	0	0	GET	^/api/users/[^/]+$	0.162	0.162	0.162	0.162	158
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.104	0.104	0.104	0.104	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.4%	98	0	0	GET	/api/notifications	956.706	9.762	13.283	13.283	178
12.7%	87	0	0	POST	/api/posts	267.015	3.069	9.712	9.712	652
8.4%	492	4	0	GET	^/api/images/[^/]+	177.555	0.361	2.534	5.334	178134
6.0%	68	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	127.158	1.870	9.647	9.647	116
4.5%	147	1	0	POST	/api/friend-requests	95.134	0.647	5.764	9.451	114
4.5%	323	0	0	GET	^/api/feeds/[^/]+	94.726	0.293	1.190	6.143	1513
4.3%	127	0	0	POST	/api/auth/login	90.020	0.709	6.113	6.894	252
2.5%	177	0	0	GET	/api/memories	53.119	0.300	2.175	4.695	10246
1.9%	113	0	0	GET	/api/posting-window/current	40.153	0.355	1.774	1.915	238
1.4%	25	3	0	POST	^/api/posts/[^/]+/realmojis$	28.611	1.144	6.519	6.519	376
1.3%	10	0	0	GET	/api/me/export	26.422	2.642	11.497	11.497	2414566
0.8%	80	5	0	GET	^/api/posts/me/current$	17.769	0.222	1.699	1.699	609
0.8%	23	0	0	PUT	^/api/follows/[^/]+$	17.501	0.761	6.879	6.879	0
0.8%	71	0	0	GET	^/api/relationships/[^/]+	17.306	0.244	0.869	0.869	1559
0.7%	23	0	0	POST	/api/auth/logout	14.344	0.624	2.303	2.303	0
0.6%	67	0	0	GET	/api/users	13.391	0.200	0.795	0.795	227
0.5%	6	0	0	DELETE	^/api/posts/[^/]+$	10.482	1.747	7.347	7.347	0
0.5%	33	1	0	GET	^/api/posts/[^/]+$	9.806	0.297	1.040	1.040	582
0.5%	12	2	0	PATCH	^/api/posts/[^/]+$	9.594	0.800	5.336	5.336	546
0.4%	31	0	0	GET	^/api/users/[^/]+$	8.545	0.276	1.134	1.134	176
0.3%	16	0	0	PUT	/api/me/profile-image	5.587	0.349	0.838	0.838	0
0.3%	24	0	0	POST	/api/auth/register	5.470	0.228	0.412	0.412	230
0.2%	23	0	0	GET	^/api/memories/[^/]+	5.269	0.229	0.953	0.953	826
0.2%	17	0	0	GET	^/api/posts/[^/]+/realmojis$	4.969	0.292	0.992	0.992	418
0.2%	20	1	0	GET	/api/me	3.392	0.170	0.359	0.359	277
0.1%	6	0	0	PATCH	/api/me	1.867	0.311	0.553	0.553	295
0.1%	9	0	0	GET	^/api/posts/[^/]+/comments$	1.767	0.196	0.434	0.434	236
0.1%	5	0	0	POST	^/api/posts/[^/]+/comments$	1.437	0.287	0.607	0.607	269
0.1%	1	0	0	POST	/api/me/account-deletion	1.162	1.162	1.162	1.162	0
0.0%	5	0	0	DELETE	^/api/realmojis/[^/]+$	1.017	0.203	0.305	0.305	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 13944 クエリ / 76 種類 / 合計 51.679s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	57.0%	75	29.435	392.46	6518.83	6065.5	128.4	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	14.8%	18	7.670	426.11	2109.85	32285.7	1655.2	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
3	12.6%	628	6.528	10.40	105.83	716.3	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
4	9.6%	2848	4.956	1.74	9.01	12.0	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.4%	474	0.728	1.54	12.03	32.2	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.5%	91	0.253	2.78	8.60	48.1	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
7	0.5%	1	0.234	233.51	233.51	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
8	0.4%	1	0.217	217.19	217.19	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
9	0.4%	1	0.205	204.88	204.88	69075.0	29506.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	0.3%	1	0.158	157.82	157.82	211771.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.3%	1	0.152	151.54	151.54	29333.0	54.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.3%	760	0.144	0.19	2.28	22.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
13	0.2%	1144	0.126	0.11	1.90	11.7	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
14	0.2%	1	0.120	119.79	119.79	2433.0	0.0	DELETE FROM image_objects WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.2%	433	0.104	0.24	3.56	25.6	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4;
-- rank 3
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 225e8b7fa1bf1e5efb352e5bd07bb9f4804c0bd2
Type: cpu
Time: 2026-10-31 10:51:03 JST
Duration: 60.57s, Total samples = 85.51s (141.17%)
Showing nodes accounting for 77.66s, 90.82% of 85.51s total
Dropped 847 nodes (cum <= 0.43s)
Showing top 40 nodes out of 154
      flat  flat%   sum%        cum   cum%
    17.46s 20.42% 20.42%     17.46s 20.42%  crypto/internal/fips140/sha256.blockAVX2
    12.78s 14.95% 35.36%     12.78s 14.95%  runtime.spanClass.sizeclass (inline)
     8.26s  9.66% 45.02%     13.38s 15.65%  image/jpeg.(*decoder).reconstructBlock
     3.80s  4.44% 49.47%     28.77s 33.65%  image/jpeg.(*decoder).processSOS
     3.66s  4.28% 53.75%     16.94s 19.81%  runtime.tryDeferToSpanScan
     3.35s  3.92% 57.67%      5.60s  6.55%  image/jpeg.(*decoder).decodeHuffman
     2.57s  3.01% 60.67%      2.72s  3.18%  image/jpeg.(*decoder).receiveExtend
     2.33s  2.72% 63.40%      2.33s  2.72%  runtime.memmove
     2.11s  2.47% 65.86%      2.11s  2.47%  runtime.memclrNoHeapPointers
     2.10s  2.46% 68.32%      2.70s  3.16%  image/jpeg.idctCols
     1.73s  2.02% 70.34%     13.13s 15.35%  runtime.scanObject
     1.70s  1.99% 72.33%      2.39s  2.79%  image/jpeg.(*decoder).ensureNBits
     1.68s  1.96% 74.30%      2.32s  2.71%  image/jpeg.idctRows
     1.49s  1.74% 76.04%      7.92s  9.26%  runtime.scanObjectsSmall
     1.34s  1.57% 77.60%      1.34s  1.57%  internal/runtime/syscall/linux.Syscall6
     1.22s  1.43% 79.03%      1.22s  1.43%  image/jpeg.dctBox (inline)
     1.09s  1.27% 80.31%     19.82s 23.18%  crypto/internal/fips140/sha256.(*Digest).checkSum
     1.01s  1.18% 81.49%      3.86s  4.51%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     0.97s  1.13% 82.62%     19.33s 22.61%  crypto/internal/fips140/sha256.(*Digest).Write
     0.87s  1.02% 83.64%      9.66s 11.30%  runtime.scanSpan
     0.68s   0.8% 84.43%      0.68s   0.8%  runtime.extractHeapBitsSmall
     0.66s  0.77% 85.21%      0.69s  0.81%  image/jpeg.(*decoder).readByteStuffedByte
     0.64s  0.75% 85.95%      0.64s  0.75%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.63s  0.74% 86.69%     20.76s 24.28%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.62s  0.73% 87.42%      0.62s  0.73%  memeqbody
     0.45s  0.53% 87.94%     22.20s 25.96%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.43s   0.5% 88.45%      0.43s   0.5%  aeshashbody
     0.38s  0.44% 88.89%      1.23s  1.44%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.27s  0.32% 89.21%     17.73s 20.73%  crypto/internal/fips140/sha256.block
     0.25s  0.29% 89.50%     23.92s 27.97%  example.com/isureal-company/isureal.imageSignature
     0.22s  0.26% 89.76%      1.41s  1.65%  runtime.mapaccess1_faststr
     0.19s  0.22% 89.98%      0.85s  0.99%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.19s  0.22% 90.20%      0.73s  0.85%  runtime.scanObjectSmall
     0.12s  0.14% 90.34%      0.45s  0.53%  crypto/internal/fips140/hmac.(*HMAC).Write
     0.11s  0.13% 90.47%      0.46s  0.54%  runtime.sweepone
     0.08s 0.094% 90.56%      0.62s  0.73%  runtime.mapassign_faststr
     0.06s  0.07% 90.63%      5.08s  5.94%  image/jpeg.idct
     0.06s  0.07% 90.70%     10.51s 12.29%  runtime.gcDrainN
     0.06s  0.07% 90.77%     11.86s 13.87%  runtime.mallocgc
     0.04s 0.047% 90.82%     15.71s 18.37%  example.com/isureal-company/isureal.listMemories
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 65.0% / 100% | 55% | 0.1% | 5 | 112 MB | CPU が飽和 |
| isu2 (2 cpu) | 98.1% / 100% | 98% | 86.2% | 6 | 174 MB | CPU が飽和 |
| isu3 (2 cpu) | 4.9% / 100% | 1% | 1.2% | 3 | 2465 MB |  |
| isu4 (2 cpu) | 29.5% / 92% | 1% | 20.3% | 9 | 106 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu2（CPU が飽和、主に image-export）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 115% / nginx 9% |
| isu2 | image-export 13% / postgres 5% / kswapd0 1% |
| isu3 | postgres 5% |
| isu4 | postgres 16% |

## Diff from previous run

スコア: 24801 → 17610 (-7191)

`cef5fe3..4350e7f`

```
 webapp/go/account_deletion_api.go | 1 +
 webapp/go/image_api.go            | 6 ++++++
 webapp/go/post_api.go             | 1 +
 webapp/go/profile_image_api.go    | 4 ++++
 webapp/go/realmoji_api.go         | 4 ++++
 5 files changed, 16 insertions(+)
```
