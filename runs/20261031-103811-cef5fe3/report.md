# run 20261031-103811-cef5fe3

## Summary

| | |
|---|---|
| score | 24801 |
| status | pass |
| ref | HEAD |
| sha | cef5fe35a9fe |
| 所要時間 | 331s |
| 前回比 | +17126 (223.1%) vs 20261031-103328-02acb51 |
| ベスト比 | +17126 (223.1%) vs 20261031-103328-02acb51 = 7675 **ベスト更新** |
| note | インメモリ化（asia1に集約、他2台はnginxで転送） |

## Bench

```
{"id":"1368","status":"done","portalStatus":"finished","passed":true,"score":24801,"raw":24801,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:42:20.966Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:42:22.948Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:42:25.202Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:42:30.856Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T01:42:31.280Z level=INFO msg="europe 2 日目: 9 人"
time=2026-10-31T01:42:34.372Z level=INFO msg="america 2 日目: 10 人"
time=2026-10-31T01:42:39.893Z level=INFO msg="asia 3 日目: 18 人"
time=2026-10-31T01:42:40.334Z level=INFO msg="europe 3 日目: 16 人"
time=2026-10-31T01:42:43.180Z level=INFO msg="america 3 日目: 24 人"
time=2026-10-31T01:42:47.920Z level=INFO msg="asia 4 日目: 18 人"
time=2026-10-31T01:42:49.522Z level=INFO msg="europe 4 日目: 34 人"
time=2026-10-31T01:42:50.378Z level=INFO msg="user moved" user=miamia.5319 from=america to=asia
time=2026-10-31T01:42:51.048Z level=INFO msg="user moved" user=miamia.5319 from=asia to=america
time=2026-10-31T01:42:53.116Z level=INFO msg="america 4 日目: 50 人"
time=2026-10-31T01:42:56.508Z level=INFO msg="asia 5 日目: 20 人"
time=2026-10-31T01:42:58.860Z level=INFO msg="europe 5 日目: 55 人"
time=2026-10-31T01:43:02.155Z level=INFO msg="america 5 日目: 50 人"
time=2026-10-31T01:43:14.714Z level=INFO msg="asia 最終日（グランプリ）: 19 人"
time=2026-10-31T01:43:14.820Z level=INFO msg="europe 最終日（グランプリ）: 54 人"
time=2026-10-31T01:43:15.131Z level=INFO msg="america 最終日（グランプリ）: 49 人"
time=2026-10-31T01:43:19.149Z level=INFO msg="グランプリ: アジア 7 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 24801,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
56.7%	111	0	0	GET	/api/notifications	1022.532	9.212	18.196	18.196	176
9.3%	1598	673	0	POST	/api/auth/login	167.230	0.105	0.889	3.060	156
6.2%	326	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	112.386	0.345	3.040	4.185	115
6.1%	104	0	0	POST	/api/posts	109.589	1.054	5.342	5.787	612
6.0%	3033	261	0	GET	^/api/images/[^/]+	107.245	0.035	0.424	3.056	47129
3.6%	118	113	0	PUT	/api/me/profile-image	65.094	0.552	3.029	4.310	24
2.4%	69	0	0	GET	/api/me/export	43.435	0.629	3.760	3.760	2027646
2.1%	689	0	0	GET	/api/memories	38.214	0.055	0.423	1.117	15906
1.6%	18	0	0	POST	/api/admin/posting-window	29.046	1.614	4.978	4.978	177
1.5%	818	0	0	GET	^/api/feeds/[^/]+	26.645	0.033	0.425	1.821	3658
0.7%	115	11	0	POST	/api/friend-requests	11.717	0.102	0.715	2.075	107
0.6%	1	0	0	POST	/api/initialize	11.615	11.615	11.615	11.615	90
0.4%	36	1	0	POST	/api/auth/register	7.935	0.220	1.259	1.259	237
0.4%	116	0	0	GET	/api/posting-window/current	7.366	0.063	0.360	0.421	237
0.3%	95	1	0	GET	^/api/posts/me/current$	5.775	0.061	0.619	0.619	606
0.3%	129	0	0	GET	^/api/relationships/[^/]+	4.865	0.038	0.362	0.373	3607
0.2%	13	0	0	POST	^/api/posts/[^/]+/realmojis$	4.369	0.336	2.772	2.772	410
0.2%	185	0	0	POST	/api/auth/logout	3.912	0.021	0.233	0.768	0
0.2%	161	0	0	GET	^/api/posts/[^/]+$	3.786	0.024	0.284	0.288	598
0.2%	176	0	0	GET	^/api/memories/[^/]+	3.765	0.021	0.338	0.659	630
0.1%	1	0	0	POST	/api/admin/auth/logout	2.497	2.497	2.497	2.497	0
0.1%	74	0	0	GET	^/api/posts/[^/]+/comments$	2.291	0.031	0.511	0.511	130
0.1%	40	0	0	GET	/api/users	2.211	0.055	0.395	0.395	252
0.1%	73	0	0	GET	^/api/posts/[^/]+/realmojis$	2.084	0.029	0.563	0.563	318
0.1%	86	0	0	GET	^/api/users/[^/]+$	1.879	0.022	0.240	0.240	178
0.1%	6	0	0	DELETE	^/api/posts/[^/]+$	1.779	0.297	1.064	1.064	0
0.1%	14	0	0	PUT	^/api/admin/users/[^/]+/certification$	0.983	0.070	0.296	0.296	0
0.0%	4	0	0	PATCH	/api/me	0.839	0.210	0.533	0.533	258
0.0%	1	0	0	POST	/api/me/account-deletion	0.414	0.414	0.414	0.414	0
0.0%	16	0	0	PATCH	^/api/posts/[^/]+$	0.400	0.025	0.064	0.064	629
…（上位 30 / 全 37 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.8%	132	0	0	GET	/api/notifications	1227.454	9.299	15.967	15.967	180
15.8%	176	2	0	POST	/api/posts	465.141	2.643	6.463	6.932	649
9.8%	1069	8	0	GET	^/api/images/[^/]+	288.585	0.270	0.944	2.556	154920
5.9%	385	32	0	POST	/api/friend-requests	172.031	0.447	3.744	4.271	108
5.1%	913	0	0	GET	^/api/feeds/[^/]+	150.476	0.165	0.460	0.825	2940
4.6%	294	0	0	POST	/api/auth/login	134.331	0.457	2.503	3.233	246
2.3%	337	0	0	GET	/api/memories	68.637	0.204	0.842	0.959	10121
2.1%	76	1	0	POST	^/api/posts/[^/]+/realmojis$	62.616	0.824	4.690	4.690	501
1.7%	238	0	0	GET	/api/posting-window/current	50.950	0.214	0.685	0.919	239
1.6%	71	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	46.133	0.650	5.047	5.047	116
1.0%	172	0	0	GET	/api/users	29.658	0.172	0.574	0.820	252
0.8%	170	0	0	GET	^/api/relationships/[^/]+	24.568	0.145	0.367	0.381	914
0.8%	26	0	0	POST	^/api/posts/[^/]+/comments$	23.583	0.907	3.611	3.611	291
0.8%	53	0	0	POST	/api/auth/register	23.531	0.444	1.605	1.605	228
0.7%	18	0	0	GET	/api/me/export	21.443	1.191	3.987	3.987	1820728
0.7%	123	0	0	GET	^/api/posts/me/current$	20.669	0.168	0.429	0.439	657
0.7%	122	5	0	GET	^/api/posts/[^/]+$	20.608	0.169	0.422	0.587	595
0.6%	57	0	0	POST	/api/auth/logout	18.267	0.320	3.326	3.326	0
0.5%	27	0	0	PUT	^/api/follows/[^/]+$	15.783	0.585	4.387	4.387	0
0.5%	31	0	0	PUT	/api/me/profile-image	13.815	0.446	0.938	0.938	0
0.4%	32	0	0	PATCH	/api/me	10.578	0.331	1.354	1.354	286
0.4%	64	0	0	GET	^/api/users/[^/]+$	10.439	0.163	0.397	0.397	183
0.2%	40	0	0	GET	^/api/memories/[^/]+	7.034	0.176	0.470	0.470	1181
0.2%	43	0	0	GET	^/api/posts/[^/]+/realmojis$	6.598	0.153	0.352	0.352	597
0.2%	37	2	0	GET	/api/me	6.255	0.169	0.515	0.515	281
0.2%	16	0	0	DELETE	^/api/realmojis/[^/]+$	4.423	0.276	0.796	0.796	0
0.1%	13	1	0	PATCH	^/api/posts/[^/]+$	3.458	0.266	0.799	0.799	616
0.1%	18	0	0	GET	^/api/posts/[^/]+/comments$	2.932	0.163	0.284	0.284	303
0.1%	4	0	0	PUT	^/api/blocks/[^/]+$	2.772	0.693	1.624	1.624	0
0.1%	6	0	0	DELETE	^/api/comments/[^/]+$	2.233	0.372	0.859	0.859	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.6%	143	0	0	GET	/api/notifications	1320.108	9.232	13.764	13.976	180
16.9%	183	1	0	POST	/api/posts	510.897	2.792	6.464	6.682	647
10.3%	1131	11	0	GET	^/api/images/[^/]+	311.498	0.275	1.044	3.329	192971
5.9%	1061	0	0	GET	^/api/feeds/[^/]+	178.677	0.168	0.394	0.604	3217
4.0%	318	32	0	POST	/api/friend-requests	121.029	0.381	2.567	4.084	107
3.7%	270	0	0	POST	/api/auth/login	112.974	0.418	2.564	3.107	254
2.7%	337	0	0	GET	/api/memories	81.823	0.243	1.703	2.746	13288
1.7%	234	0	0	GET	/api/posting-window/current	51.800	0.221	0.634	1.671	240
1.1%	15	0	0	GET	/api/me/export	34.241	2.283	7.046	7.046	3389991
1.1%	203	10	0	GET	^/api/posts/[^/]+$	33.561	0.165	0.415	0.580	581
1.1%	62	1	0	POST	^/api/posts/[^/]+/realmojis$	33.305	0.537	2.131	2.131	434
0.9%	157	0	0	GET	/api/users	27.102	0.173	0.422	0.477	224
0.9%	134	3	0	GET	^/api/posts/me/current$	25.867	0.193	0.613	0.614	634
0.8%	159	0	0	GET	^/api/relationships/[^/]+	24.927	0.157	0.304	1.598	739
0.8%	54	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	23.572	0.437	1.895	1.895	116
0.7%	48	2	0	POST	^/api/posts/[^/]+/comments$	22.306	0.465	2.464	2.464	350
0.4%	84	0	0	GET	^/api/users/[^/]+$	13.459	0.160	0.341	0.341	190
0.4%	49	0	0	POST	/api/auth/register	12.054	0.246	0.507	0.507	229
0.4%	46	0	0	POST	/api/auth/logout	11.816	0.257	2.152	2.152	0
0.4%	31	0	0	PUT	/api/me/profile-image	11.575	0.373	0.881	0.881	0
0.3%	37	0	0	PUT	^/api/follows/[^/]+$	10.210	0.276	1.337	1.337	0
0.3%	57	0	0	GET	^/api/posts/[^/]+/realmojis$	8.481	0.149	0.400	0.400	604
0.2%	42	0	0	GET	^/api/memories/[^/]+	6.852	0.163	0.525	0.525	1120
0.2%	38	0	0	GET	^/api/posts/[^/]+/comments$	6.106	0.161	0.270	0.270	338
0.2%	23	0	0	PATCH	/api/me	6.046	0.263	0.584	0.584	303
0.2%	14	0	0	PATCH	^/api/posts/[^/]+$	5.749	0.411	0.922	0.922	628
0.2%	13	0	0	DELETE	^/api/comments/[^/]+$	5.461	0.420	1.678	1.678	0
0.2%	36	1	0	GET	/api/me	5.304	0.147	0.261	0.261	286
0.2%	11	0	0	DELETE	^/api/posts/[^/]+$	4.641	0.422	0.638	0.638	0
0.1%	12	0	0	DELETE	^/api/realmojis/[^/]+$	4.230	0.352	1.244	1.244	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16593 クエリ / 78 種類 / 合計 68.291s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	49.4%	102	33.722	330.61	2398.52	6775.7	137.5	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	16.7%	18	11.391	632.84	2708.20	32171.9	1650.0	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
3	12.1%	678	8.273	12.20	197.93	644.6	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
4	9.0%	3789	6.135	1.62	6.23	11.1	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	3.3%	1	2.255	2254.84	2254.84	29256.0	1199390.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	2.2%	1	1.493	1492.79	1492.79	19464.0	625302.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	1.4%	1	0.961	960.75	960.75	11528.0	529646.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	1.1%	461	0.763	1.65	22.66	32.4	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
9	0.7%	2	0.503	251.38	344.23	212153.5	2495.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
10	0.5%	149	0.343	2.30	4.46	38.1	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
11	0.5%	2	0.340	169.96	332.04	18870.5	2470.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.3%	1	0.232	231.64	231.64	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
13	0.3%	1	0.201	201.37	201.37	19847.0	191.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.3%	1	0.201	201.23	201.23	66940.0	29532.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.3%	1	0.197	197.40	197.40	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…

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
Build ID: dd2e7b75bd84e4475f3c3497c1412d23e56528a9
Type: cpu
Time: 2026-10-31 10:42:02 JST
Duration: 60.16s, Total samples = 83.38s (138.61%)
Showing nodes accounting for 73.44s, 88.08% of 83.38s total
Dropped 842 nodes (cum <= 0.42s)
Showing top 40 nodes out of 170
      flat  flat%   sum%        cum   cum%
    20.84s 24.99% 24.99%     20.84s 24.99%  crypto/internal/fips140/sha256.blockAVX2
     9.61s 11.53% 36.52%      9.61s 11.53%  runtime.spanClass.sizeclass (inline)
     6.15s  7.38% 43.90%      9.94s 11.92%  image/jpeg.(*decoder).reconstructBlock
     3.70s  4.44% 48.33%     13.81s 16.56%  runtime.tryDeferToSpanScan
     3.07s  3.68% 52.01%      3.07s  3.68%  runtime.memmove
     2.64s  3.17% 55.18%     20.73s 24.86%  image/jpeg.(*decoder).processSOS
     2.52s  3.02% 58.20%      4.17s  5.00%  image/jpeg.(*decoder).decodeHuffman
     2.23s  2.67% 60.88%      2.23s  2.67%  runtime.memclrNoHeapPointers
     1.71s  2.05% 62.93%      1.78s  2.13%  image/jpeg.(*decoder).receiveExtend
     1.63s  1.95% 64.88%      6.12s  7.34%  runtime.scanObjectsSmall
     1.62s  1.94% 66.83%     11.40s 13.67%  runtime.scanObject
     1.59s  1.91% 68.73%      2.02s  2.42%  image/jpeg.idctCols
     1.33s  1.60% 70.33%      5.63s  6.75%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.33s  1.60% 71.92%      1.33s  1.60%  internal/runtime/syscall/linux.Syscall6
     1.29s  1.55% 73.47%      1.74s  2.09%  image/jpeg.idctRows
     1.17s  1.40% 74.87%      1.71s  2.05%  image/jpeg.(*decoder).ensureNBits
     1.10s  1.32% 76.19%      1.10s  1.32%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.07s  1.28% 77.48%     23.48s 28.16%  crypto/internal/fips140/sha256.(*Digest).checkSum
     1.06s  1.27% 78.75%     22.91s 27.48%  crypto/internal/fips140/sha256.(*Digest).Write
     0.85s  1.02% 79.77%      0.85s  1.02%  image/jpeg.dctBox (inline)
     0.84s  1.01% 80.77%      0.84s  1.01%  aeshashbody
     0.82s  0.98% 81.76%     24.79s 29.73%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.77s  0.92% 82.68%      0.77s  0.92%  memeqbody
     0.61s  0.73% 83.41%     26.43s 31.70%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.55s  0.66% 84.07%      1.47s  1.76%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.50s   0.6% 84.67%      0.54s  0.65%  image/jpeg.(*decoder).readByteStuffedByte
     0.39s  0.47% 85.14%      0.55s  0.66%  example.com/isureal-company/isureal.postLess
     0.36s  0.43% 85.57%      6.98s  8.37%  runtime.scanSpan
     0.31s  0.37% 85.94%         2s  2.40%  runtime.mapaccess1_faststr
     0.29s  0.35% 86.29%     21.14s 25.35%  crypto/internal/fips140/sha256.block
     0.28s  0.34% 86.63%      1.06s  1.27%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.23s  0.28% 86.90%     28.52s 34.20%  example.com/isureal-company/isureal.imageSignature
     0.19s  0.23% 87.13%      9.28s 11.13%  runtime.mallocgc
     0.18s  0.22% 87.35%      1.25s  1.50%  runtime.mapassign_faststr
     0.14s  0.17% 87.51%      0.80s  0.96%  runtime.mallocgcSmallScanNoHeader
     0.12s  0.14% 87.66%      0.53s  0.64%  crypto/internal/fips140/hmac.(*HMAC).Write
     0.11s  0.13% 87.79%      0.43s  0.52%  runtime.scanObjectSmall
     0.09s  0.11% 87.90%      2.41s  2.89%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.08s 0.096% 87.99%      0.55s  0.66%  runtime.mallocgcSmallNoscan
     0.07s 0.084% 88.08%      1.04s  1.25%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 67.4% / 100% | 49% | 0.1% | 6 | 156 MB | CPU が飽和 |
| isu2 (2 cpu) | 1.1% / 16% | 0% | 0.0% | 0 | 3100 MB |  |
| isu3 (2 cpu) | 1.5% / 21% | 0% | 0.0% | 1 | 3212 MB |  |
| isu4 (2 cpu) | 29.9% / 98% | 3% | 19.6% | 6 | 107 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 121% / nginx 7% |
| isu2 | nginx 1% |
| isu3 | nginx 2% |
| isu4 | postgres 18% |

## Diff from previous run

スコア: 7675 → 24801 (+17126)

`02acb51..cef5fe3`

```
 conf/isu1/etc/nginx/nginx.conf                     |   3 +-
 conf/isu1/etc/nginx/sites-enabled/isureal.conf     |  20 +-
 conf/isu1/etc/systemd/system/isureal-go.service    |   2 +
 conf/isu1/home/isucon/env.sh                       |   1 +
 conf/isu2/etc/nginx/nginx.conf                     |   3 +-
 conf/isu2/etc/nginx/sites-enabled/isureal.conf     |  20 +-
 conf/isu2/home/isucon/env.sh                       |   1 +
 conf/isu3/etc/nginx/nginx.conf                     |   3 +-
 conf/isu3/etc/nginx/sites-enabled/isureal.conf     |  20 +-
 conf/isu3/home/isucon/env.sh                       |   1 +
 .../etc/postgresql/16/main/conf.d/zz-tuning.conf   |   7 +
 webapp/go/account_deletion_api.go                  |  32 +-
 webapp/go/admin_api.go                             |  48 +-
 webapp/go/admin_posting_window_api.go              |  86 ++-
 webapp/go/admin_session.go                         |  84 +--
 webapp/go/admin_user_api.go                        |  21 +-
 webapp/go/cmd/isureal/main.go                      |  13 +-
 webapp/go/comment.go                               |  75 +--
 webapp/go/comment_api.go                           | 145 ++---
 webapp/go/export_api.go                            |  92 ++-
 webapp/go/feed_api.go                              | 170 +++--
 webapp/go/friend_request_api.go                    | 140 +++--
 webapp/go/image_api.go                             | 105 ++--
 webapp/go/image_cache.go                           |  53 ++
 webapp/go/initialize.go                            |  25 +-
 webapp/go/memory_api.go                            | 124 ++--
 webapp/go/notification_api.go                      |  66 +-
 webapp/go/post.go                                  | 243 +++-----
 webapp/go/post_api.go                              | 255 +++-----
 webapp/go/posting_window.go                        |  21 +-
 webapp/go/posting_window_api.go                    |  41 +-
 webapp/go/profile_image_api.go                     |  94 ++-
 webapp/go/realmoji.go                              |  63 +-
 webapp/go/realmoji_api.go                          | 131 ++--
 webapp/go/region.go                                |  10 +-
 webapp/go/relationship.go                          | 118 ++--
 webapp/go/relationship_api.go                      | 154 +++--
 webapp/go/relationship_list_api.go                 |  88 ++-
 webapp/go/server.go                                |  74 ++-
 webapp/go/session.go                               | 123 +---
 webapp/go/store.go                                 | 686 +++++++++++++++++++++
 webapp/go/user.go                                  |  24 -
 webapp/go/user_api.go                              | 187 +++---
 webapp/go/user_read_api.go                         |  75 +--
 webapp/sql/002-indexes.sql                         |  10 +
 webapp/sql/reset.sql                               |   2 +
 46 files changed, 2323 insertions(+), 1436 deletions(-)
```
