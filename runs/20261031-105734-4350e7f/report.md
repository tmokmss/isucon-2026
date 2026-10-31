# run 20261031-105734-4350e7f

## Summary

| | |
|---|---|
| score | 14209 |
| status | pass |
| ref | origin/main |
| sha | 4350e7f0be16 |
| 所要時間 | 94s |
| 前回比 | -3401 (-19.3%) vs 20261031-104710-4350e7f |
| ベスト比 | -10592 (-42.7%) vs 20261031-103811-cef5fe3 = 24801 |
| note | main 基準（インメモリ化、書き出しはロック中停止） |

## Bench

```
{"id":"2148","status":"done","portalStatus":"finished","passed":true,"score":14209,"raw":14209,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:57:52.999Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:57:55.388Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:57:58.108Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:58:01.791Z level=INFO msg="asia 2 日目: 37 人"
time=2026-10-31T01:58:04.678Z level=INFO msg="europe 2 日目: 13 人"
time=2026-10-31T01:58:07.099Z level=INFO msg="america 2 日目: 12 人"
time=2026-10-31T01:58:11.353Z level=INFO msg="asia 3 日目: 82 人"
time=2026-10-31T01:58:14.008Z level=INFO msg="europe 3 日目: 23 人"
time=2026-10-31T01:58:16.434Z level=INFO msg="america 3 日目: 17 人"
time=2026-10-31T01:58:16.687Z level=INFO msg="user moved" user=liavinso_5211 from=asia to=america
time=2026-10-31T01:58:21.599Z level=INFO msg="asia 4 日目: 90 人"
time=2026-10-31T01:58:26.759Z level=INFO msg="europe 4 日目: 23 人"
time=2026-10-31T01:58:30.594Z level=INFO msg="america 4 日目: 18 人"
time=2026-10-31T01:58:33.012Z level=INFO msg="asia 5 日目: 90 人"
time=2026-10-31T01:58:35.993Z level=INFO msg="europe 5 日目: 23 人"
time=2026-10-31T01:58:40.431Z level=INFO msg="america 5 日目: 18 人"
time=2026-10-31T01:58:48.857Z level=INFO msg="asia 最終日（グランプリ）: 90 人"
time=2026-10-31T01:58:52.083Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 14209,
  "messages": [
    "niso-5733: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "niso-5733: 投稿できるはずの投稿期間で投稿が拒否されました: POST /api/posts: ステータスコード 409 が返りました（期待値 201、post_slot_unavailable）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.6%	315	0	0	GET	/api/notifications	3344.011	10.616	17.096	18.407	176
16.3%	314	7	0	POST	/api/posts	1489.455	4.743	14.085	14.844	623
13.1%	2410	1659	0	POST	/api/auth/login	1192.938	0.495	3.440	10.439	97
10.7%	3128	803	0	GET	^/api/images/[^/]+	972.536	0.311	2.891	8.716	84590
4.7%	184	141	0	PUT	/api/me/profile-image	425.879	2.315	10.405	14.421	18
4.5%	285	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	414.123	1.453	10.985	13.285	116
3.9%	532	14	0	POST	/api/friend-requests	353.150	0.664	6.435	14.429	113
1.9%	1448	0	0	GET	^/api/feeds/[^/]+	173.420	0.120	1.261	2.327	1491
1.5%	107	1	0	POST	/api/auth/register	138.614	1.295	8.236	8.875	230
1.5%	771	0	0	GET	/api/memories	132.925	0.172	1.815	3.169	10520
1.2%	411	0	0	GET	/api/posting-window/current	113.534	0.276	1.857	2.329	237
0.7%	18	0	0	POST	/api/admin/posting-window	61.875	3.438	11.997	11.997	177
0.5%	265	18	0	GET	^/api/posts/me/current$	46.192	0.174	1.641	2.720	595
0.4%	45	0	0	GET	/api/me/export	38.674	0.859	4.383	4.383	1941091
0.4%	31	4	0	PATCH	^/api/posts/[^/]+$	34.851	1.124	10.381	10.381	549
0.3%	134	0	0	POST	/api/auth/logout	29.694	0.222	2.956	4.867	0
0.3%	244	0	0	GET	/api/users	27.387	0.112	0.646	1.325	246
0.3%	285	0	0	GET	^/api/relationships/[^/]+	25.285	0.089	0.563	1.916	1670
0.2%	226	10	0	GET	^/api/posts/[^/]+$	18.240	0.081	0.565	1.553	575
0.1%	27	0	0	PATCH	/api/me	13.364	0.495	5.162	5.162	277
0.1%	43	0	0	POST	^/api/posts/[^/]+/comments$	11.926	0.277	1.891	1.891	325
0.1%	34	0	0	POST	^/api/posts/[^/]+/realmojis$	11.689	0.344	1.296	1.296	472
0.1%	1	0	0	POST	/api/initialize	9.219	9.219	9.219	9.219	90
0.1%	31	0	0	PUT	^/api/follows/[^/]+$	8.310	0.268	2.068	2.068	0
0.1%	122	0	0	GET	^/api/memories/[^/]+	7.505	0.062	0.869	1.728	541
0.1%	14	0	0	DELETE	^/api/posts/[^/]+$	6.765	0.483	1.356	1.356	0
0.1%	91	0	0	GET	^/api/users/[^/]+$	6.469	0.071	0.293	0.293	170
0.0%	50	2	0	GET	/api/me	3.811	0.076	0.310	0.310	282
0.0%	5	0	0	PUT	^/api/blocks/[^/]+$	3.504	0.701	1.626	1.626	0
0.0%	1	0	0	POST	/api/me/account-deletion	3.486	3.486	3.486	3.486	0
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
49.0%	92	0	0	GET	/api/notifications	1018.978	11.076	18.894	18.894	179
15.7%	73	1	0	POST	/api/posts	325.238	4.455	15.708	15.708	633
7.6%	346	4	0	GET	^/api/images/[^/]+	158.598	0.458	3.128	5.285	208329
4.4%	43	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	92.300	2.147	8.550	8.550	116
4.4%	120	6	0	POST	/api/friend-requests	91.213	0.760	4.013	7.975	111
3.6%	92	0	0	POST	/api/auth/login	74.422	0.809	8.303	8.303	250
3.3%	309	0	0	GET	^/api/feeds/[^/]+	69.255	0.224	0.872	1.676	1095
2.7%	146	0	0	GET	/api/memories	55.173	0.378	3.416	6.295	9676
1.8%	97	0	0	GET	/api/posting-window/current	37.030	0.382	2.044	2.044	239
1.0%	21	0	0	POST	/api/auth/register	20.452	0.974	2.439	2.439	228
0.9%	69	4	0	GET	^/api/posts/me/current$	19.006	0.275	1.488	1.488	607
0.9%	9	0	0	GET	/api/me/export	18.667	2.074	7.448	7.448	2405710
0.7%	67	0	0	GET	^/api/relationships/[^/]+	15.191	0.227	1.866	1.866	1048
0.7%	59	0	0	GET	/api/users	13.935	0.236	1.497	1.497	252
0.6%	5	0	0	DELETE	^/api/posts/[^/]+$	12.537	2.507	8.491	8.491	0
0.4%	36	2	0	GET	^/api/posts/[^/]+$	8.027	0.223	1.466	1.466	566
0.4%	18	0	0	POST	/api/auth/logout	7.635	0.424	2.669	2.669	0
0.3%	16	0	0	PUT	^/api/follows/[^/]+$	5.787	0.362	1.534	1.534	0
0.3%	9	0	0	PUT	/api/me/profile-image	5.325	0.592	1.084	1.084	0
0.2%	14	0	0	POST	^/api/posts/[^/]+/comments$	5.162	0.369	2.039	2.039	318
0.2%	27	0	0	GET	^/api/users/[^/]+$	5.023	0.186	0.768	0.768	165
0.2%	12	0	0	GET	^/api/memories/[^/]+	4.607	0.384	1.193	1.193	921
0.2%	7	0	0	PATCH	/api/me	4.373	0.625	2.672	2.672	287
0.2%	1	0	0	POST	/api/me/account-deletion	3.648	3.648	3.648	3.648	0
0.1%	14	0	0	GET	^/api/posts/[^/]+/comments$	2.295	0.164	0.478	0.478	644
0.1%	13	1	0	GET	/api/me	2.276	0.175	0.319	0.319	261
0.0%	4	0	0	PATCH	^/api/posts/[^/]+$	0.577	0.144	0.199	0.199	618
0.0%	4	0	0	GET	^/api/posts/[^/]+/realmojis$	0.515	0.129	0.158	0.158	159
0.0%	2	0	0	DELETE	^/api/comments/[^/]+$	0.421	0.210	0.286	0.286	0
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.105	0.105	0.105	0.105	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
48.2%	73	0	0	GET	/api/notifications	804.944	11.027	14.119	14.119	182
13.8%	52	0	0	POST	/api/posts	230.575	4.434	14.404	14.404	630
9.7%	272	2	0	GET	^/api/images/[^/]+	162.644	0.598	5.992	10.425	167165
5.5%	88	2	0	POST	/api/friend-requests	92.245	1.048	9.886	9.886	113
4.4%	27	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	73.026	2.705	11.505	11.505	116
3.8%	80	0	0	POST	/api/auth/login	63.628	0.795	8.028	8.028	252
3.2%	163	0	0	GET	^/api/feeds/[^/]+	54.192	0.332	1.469	2.628	767
2.5%	114	0	0	GET	/api/memories	41.510	0.364	2.434	2.439	9017
1.6%	66	1	0	GET	/api/posting-window/current	26.500	0.402	2.046	2.046	236
1.5%	9	0	0	GET	/api/me/export	24.359	2.707	8.345	8.345	2971007
1.1%	44	7	0	GET	^/api/posts/me/current$	18.820	0.428	2.046	2.046	529
0.9%	47	0	0	GET	/api/users	15.313	0.326	1.682	1.682	258
0.8%	4	0	0	DELETE	^/api/posts/[^/]+$	12.975	3.244	9.810	9.810	0
0.5%	37	0	0	GET	^/api/relationships/[^/]+	7.832	0.212	1.030	1.030	1431
0.4%	16	0	0	POST	/api/auth/register	7.388	0.462	1.384	1.384	229
0.4%	11	0	0	POST	/api/auth/logout	6.835	0.621	3.077	3.077	0
0.2%	3	0	0	PATCH	^/api/posts/[^/]+$	3.865	1.288	3.114	3.114	622
0.2%	6	0	0	PUT	/api/me/profile-image	3.505	0.584	1.142	1.142	0
0.2%	9	0	0	POST	^/api/posts/[^/]+/comments$	2.970	0.330	0.572	0.572	269
0.2%	5	0	0	PATCH	/api/me	2.834	0.567	2.099	2.099	311
0.2%	12	0	0	GET	^/api/memories/[^/]+	2.658	0.222	0.671	0.671	708
0.2%	15	3	0	GET	^/api/posts/[^/]+$	2.556	0.170	0.313	0.313	471
0.1%	10	0	0	GET	^/api/posts/[^/]+/comments$	1.682	0.168	0.262	0.262	488
0.1%	9	0	0	GET	/api/me	1.375	0.153	0.262	0.262	268
0.1%	8	0	0	PUT	^/api/follows/[^/]+$	1.350	0.169	0.257	0.257	0
0.1%	7	0	0	GET	^/api/users/[^/]+$	1.311	0.187	0.259	0.259	156
0.1%	4	0	0	DELETE	^/api/comments/[^/]+$	1.063	0.266	0.337	0.337	0
0.0%	4	0	0	GET	^/api/posts/[^/]+/realmojis$	0.719	0.180	0.281	0.281	159
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.106	0.106	0.106	0.106	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 11504 クエリ / 77 種類 / 合計 40.982s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	51.0%	63	20.892	331.61	1259.89	6464.0	136.7	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	12.8%	523	5.253	10.04	54.26	767.9	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	11.7%	18	4.803	266.85	896.51	32694.2	1656.8	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
4	8.8%	2144	3.626	1.69	6.88	12.9	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	5.8%	1	2.393	2393.23	2393.23	29299.0	1199418.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	1.7%	431	0.713	1.65	4.98	34.2	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
7	1.7%	1	0.696	695.51	695.51	11528.0	529651.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	0.8%	2	0.322	160.93	163.85	213981.0	2517.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
9	0.7%	2	0.282	140.99	274.23	20229.5	2507.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
10	0.5%	1	0.222	221.95	221.95	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.5%	1	0.211	210.67	210.67	69809.0	29697.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.5%	1	0.187	187.05	187.05	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
13	0.4%	704	0.149	0.21	3.66	23.8	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
14	0.4%	1	0.145	145.48	145.48	29370.0	70.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.3%	34	0.125	3.67	8.67	59.1	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)

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
Build ID: 225e8b7fa1bf1e5efb352e5bd07bb9f4804c0bd2
Type: cpu
Time: 2026-10-31 10:57:38 JST
Duration: 60.86s, Total samples = 88.73s (145.78%)
Showing nodes accounting for 79.71s, 89.83% of 88.73s total
Dropped 826 nodes (cum <= 0.44s)
Showing top 40 nodes out of 167
      flat  flat%   sum%        cum   cum%
    15.30s 17.24% 17.24%     15.30s 17.24%  crypto/internal/fips140/sha256.blockAVX2
    14.54s 16.39% 33.63%     14.54s 16.39%  runtime.spanClass.sizeclass (inline)
        8s  9.02% 42.65%     12.76s 14.38%  image/jpeg.(*decoder).reconstructBlock
     4.67s  5.26% 47.91%     19.97s 22.51%  runtime.tryDeferToSpanScan
     3.57s  4.02% 51.93%     27.26s 30.72%  image/jpeg.(*decoder).processSOS
     2.96s  3.34% 55.27%      4.97s  5.60%  image/jpeg.(*decoder).decodeHuffman
     2.71s  3.05% 58.32%      2.71s  3.05%  runtime.memclrNoHeapPointers
     2.52s  2.84% 61.16%      2.52s  2.84%  runtime.memmove
     2.43s  2.74% 63.90%      2.54s  2.86%  image/jpeg.(*decoder).receiveExtend
     2.12s  2.39% 66.29%      9.12s 10.28%  runtime.scanObjectsSmall
     1.92s  2.16% 68.45%      2.44s  2.75%  image/jpeg.idctCols
     1.90s  2.14% 70.60%     15.84s 17.85%  runtime.scanObject
     1.66s  1.87% 72.47%      2.23s  2.51%  image/jpeg.idctRows
     1.56s  1.76% 74.23%      2.10s  2.37%  image/jpeg.(*decoder).ensureNBits
     1.41s  1.59% 75.81%      1.41s  1.59%  internal/runtime/syscall/linux.Syscall6
     1.30s  1.47% 77.28%      5.61s  6.32%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.08s  1.22% 78.50%      1.09s  1.23%  image/jpeg.dctBox (inline)
     0.96s  1.08% 79.58%      0.96s  1.08%  memeqbody
     0.93s  1.05% 80.63%      0.93s  1.05%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.90s  1.01% 81.64%      0.90s  1.01%  aeshashbody
     0.87s  0.98% 82.62%     16.96s 19.11%  crypto/internal/fips140/sha256.(*Digest).Write
     0.80s   0.9% 83.52%     17.25s 19.44%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.78s  0.88% 84.40%     10.95s 12.34%  runtime.scanSpan
     0.61s  0.69% 85.09%      0.61s  0.69%  runtime.extractHeapBitsSmall
     0.56s  0.63% 85.72%     18.21s 20.52%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.48s  0.54% 86.26%      0.54s  0.61%  image/jpeg.(*decoder).readByteStuffedByte
     0.46s  0.52% 86.78%      0.52s  0.59%  runtime.(*spanInlineMarkBits).tryAcquire
     0.42s  0.47% 87.25%     19.39s 21.85%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.36s  0.41% 87.66%      0.55s  0.62%  example.com/isureal-company/isureal.postLess
     0.29s  0.33% 87.99%      0.92s  1.04%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.29s  0.33% 88.31%      2.23s  2.51%  runtime.mapaccess1_faststr
     0.25s  0.28% 88.59%      0.94s  1.06%  runtime.scanObjectSmall
     0.22s  0.25% 88.84%      0.76s  0.86%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.21s  0.24% 89.08%     20.93s 23.59%  example.com/isureal-company/isureal.imageSignature
     0.17s  0.19% 89.27%     15.47s 17.43%  crypto/internal/fips140/sha256.block
     0.13s  0.15% 89.42%      1.06s  1.19%  runtime.mapassign_faststr
     0.11s  0.12% 89.54%      0.73s  0.82%  runtime.mallocgcSmallScanNoHeader
     0.10s  0.11% 89.65%      0.57s  0.64%  runtime.sweepone
     0.08s  0.09% 89.74%      4.75s  5.35%  image/jpeg.idct
     0.08s  0.09% 89.83%     13.62s 15.35%  runtime.gcDrainN
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 76.1% / 100% | 66% | 0.5% | 7 | 111 MB | CPU が飽和 |
| isu2 (2 cpu) | 2.4% / 30% | 0% | 0.4% | 1 | 194 MB |  |
| isu3 (2 cpu) | 8.2% / 56% | 0% | 3.2% | 3 | 175 MB |  |
| isu4 (2 cpu) | 22.9% / 79% | 0% | 14.0% | 8 | 102 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 130% / nginx 10% |
| isu2 | 目立つプロセス無し |
| isu3 | postgres 9% |
| isu4 | postgres 15% |

## Diff from previous run

スコア: 17610 → 14209 (-3401)

コードの差分は無い（同じ sha 4350e7f0be16）。
スコアの違いはブレか、サーバ側の設定変更による。
