# run 20261031-115409-773f053

## Summary

| | |
|---|---|
| score | 56553 |
| status | pass |
| ref | try-limiter-register |
| sha | 773f0539cdf5 |
| 所要時間 | 223s |
| 前回比 | -2326 (-4.0%) vs 20261031-115123-0904d44 |
| ベスト比 | -3648 (-6.1%) vs 20261031-115040-1c73ee4 = 60201 |
| note | main + #70 同時実行の制御 + 登録の並列化 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	52	155	232	268	268
europe	7	19	39	54	52	52
america	6	14	26	43	43	42
```

```
{"id":"5987","status":"done","portalStatus":"finished","passed":true,"score":56553,"raw":56553,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:56:36.022Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:56:38.980Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:56:42.009Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:56:44.996Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T02:56:47.988Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T02:56:50.992Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T02:56:51.742Z level=INFO msg="user moved" user=kuhe_5431 from=europe to=asia
time=2026-10-31T02:56:54.040Z level=INFO msg="asia 3 日目: 155 人"
time=2026-10-31T02:56:55.904Z level=INFO msg="user moved" user=kuhe_5431 from=asia to=europe
time=2026-10-31T02:56:57.015Z level=INFO msg="europe 3 日目: 39 人"
time=2026-10-31T02:56:59.983Z level=INFO msg="america 3 日目: 26 人"
time=2026-10-31T02:57:00.499Z level=INFO msg="user moved" user=yunayu5449 from=europe to=asia
time=2026-10-31T02:57:03.033Z level=INFO msg="asia 4 日目: 232 人"
time=2026-10-31T02:57:05.475Z level=INFO msg="user moved" user=yunayu5449 from=asia to=europe
time=2026-10-31T02:57:06.134Z level=INFO msg="europe 4 日目: 54 人"
time=2026-10-31T02:57:08.738Z level=INFO msg="user moved" user=ruelke6635 from=asia to=america
time=2026-10-31T02:57:08.984Z level=INFO msg="america 4 日目: 43 人"
time=2026-10-31T02:57:10.306Z level=INFO msg="user moved" user=kuhe_5431 from=europe to=asia
time=2026-10-31T02:57:12.013Z level=INFO msg="asia 5 日目: 268 人"
time=2026-10-31T02:57:12.180Z level=INFO msg="user moved" user=nishiya5458 from=europe to=asia
time=2026-10-31T02:57:15.102Z level=INFO msg="europe 5 日目: 52 人"
time=2026-10-31T02:57:18.103Z level=INFO msg="america 5 日目: 43 人"
time=2026-10-31T02:57:27.149Z level=INFO msg="europe 最終日（グランプリ）: 52 人"
time=2026-10-31T02:57:27.177Z level=INFO msg="america 最終日（グランプリ）: 42 人"
time=2026-10-31T02:57:27.379Z level=INFO msg="asia 最終日（グランプリ）: 268 人"
time=2026-10-31T02:57:35.984Z level=INFO msg="グランプリ: アジア 174 人, ヨーロッパ 36 人, アメリカ 28 人"
{
  "pass": true,
  "score": 56553,
  "messages": [
    "yomu.5527: プロフィール画像を変えられません: PUT /api/me/profile-image: ステータスコード 503 が返りました（期待値 204、busy）（5 回試行）",
    "ulto_5624: プロフィール画像を変えられません: PUT /api/me/profile-image: ステータスコード 503 が返りました（期待値 204、busy）（5 回試行）",
    "tefuvin5089: プロフィール画像を変えられません: PUT /api/me/profile-image: ステータスコード 503 が返りました（期待値 204、busy）（5 回試行）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.3%	558	0	0	GET	/api/notifications	5348.592	9.585	15.412	15.413	174
32.9%	593	86	400	PUT	/api/me/profile-image	4357.748	7.349	8.950	9.243	15
12.5%	718	3	0	POST	/api/posts	1653.038	2.302	6.234	6.340	626
9.7%	181	0	117	GET	/api/me/export	1284.877	7.099	9.350	9.475	698180
0.8%	7160	0	0	GET	^/api/feeds/[^/]+	99.507	0.014	0.088	0.349	4499
0.7%	5376	197	0	GET	^/api/images/[^/]+	87.641	0.016	0.140	0.462	190182
0.5%	1106	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	63.191	0.057	0.250	0.386	116
0.4%	1641	0	0	GET	/api/memories	59.658	0.036	0.293	0.429	9121
0.4%	1872	513	0	POST	/api/auth/login	47.144	0.025	0.175	0.844	185
0.3%	1473	44	0	POST	/api/friend-requests	44.084	0.030	0.183	0.242	112
0.3%	444	0	0	POST	^/api/posts/[^/]+/realmojis$	40.055	0.090	0.315	0.959	462
0.2%	1557	1	0	GET	^/api/posts/[^/]+$	21.551	0.014	0.094	0.149	613
0.1%	406	0	0	POST	^/api/posts/[^/]+/comments$	17.332	0.043	0.212	0.237	310
0.1%	95	0	0	DELETE	^/api/realmojis/[^/]+$	15.356	0.162	0.402	0.402	0
0.1%	1	0	0	POST	/api/initialize	13.023	13.023	13.023	13.023	90
0.1%	949	0	0	GET	/api/posting-window/current	12.254	0.013	0.099	0.132	237
0.1%	998	0	0	GET	^/api/relationships/[^/]+	11.951	0.012	0.078	0.205	1222
0.1%	691	1	0	GET	^/api/posts/me/current$	11.459	0.017	0.117	0.433	627
0.1%	720	0	0	GET	^/api/users/[^/]+$	10.086	0.014	0.091	0.166	190
0.1%	684	0	0	GET	/api/users	9.622	0.014	0.089	0.134	225
0.1%	734	0	0	GET	^/api/posts/[^/]+/realmojis$	9.426	0.013	0.078	0.134	782
0.1%	284	1	0	POST	/api/auth/register	8.021	0.028	0.147	0.336	227
0.1%	533	0	0	GET	^/api/posts/[^/]+/comments$	7.492	0.014	0.078	0.136	798
0.0%	185	0	0	GET	^/api/memories/[^/]+	5.241	0.028	0.326	0.349	631
0.0%	28	0	0	DELETE	^/api/posts/[^/]+$	4.791	0.171	0.282	0.282	0
0.0%	182	0	0	PUT	^/api/follows/[^/]+$	4.786	0.026	0.171	0.231	0
0.0%	235	0	0	POST	/api/auth/logout	4.356	0.019	0.087	0.099	0
0.0%	108	0	0	PATCH	/api/me	2.782	0.026	0.117	0.124	255
0.0%	89	0	0	PATCH	^/api/posts/[^/]+$	2.588	0.029	0.132	0.132	635
0.0%	96	0	0	DELETE	^/api/comments/[^/]+$	2.046	0.021	0.169	0.169	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.1%	136	0	0	GET	/api/notifications	1122.428	8.253	11.956	11.965	173
21.3%	82	0	58	PUT	/api/me/profile-image	646.306	7.882	8.287	8.287	12
13.1%	163	0	0	POST	/api/posts	397.322	2.438	6.367	7.178	628
8.2%	33	0	22	GET	/api/me/export	249.582	7.563	9.554	9.554	768001
5.0%	1399	0	0	GET	^/api/feeds/[^/]+	152.456	0.109	0.161	0.221	4506
4.1%	952	4	0	GET	^/api/images/[^/]+	124.564	0.131	0.937	1.250	246515
1.4%	305	0	0	GET	/api/memories	41.097	0.135	0.445	0.595	11574
1.2%	299	8	0	POST	/api/friend-requests	35.110	0.117	0.210	0.379	113
1.0%	219	0	0	POST	/api/auth/login	30.308	0.138	0.290	0.301	246
0.8%	215	0	0	GET	/api/posting-window/current	25.642	0.119	0.247	0.280	239
0.8%	219	2	0	GET	^/api/posts/[^/]+$	23.838	0.109	0.150	0.204	602
0.7%	206	0	0	GET	^/api/relationships/[^/]+	22.253	0.108	0.148	0.152	1019
0.7%	121	0	0	POST	^/api/posts/[^/]+/realmojis$	20.114	0.166	0.376	0.388	444
0.6%	149	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	18.344	0.123	0.240	0.276	116
0.6%	156	1	0	GET	^/api/posts/me/current$	16.915	0.108	0.152	0.210	624
0.5%	147	0	0	GET	/api/users	16.302	0.111	0.204	0.215	243
0.5%	132	0	0	GET	^/api/posts/[^/]+/realmojis$	14.326	0.109	0.148	0.167	726
0.4%	122	0	0	GET	^/api/users/[^/]+$	13.216	0.108	0.150	0.163	186
0.4%	101	0	0	GET	^/api/posts/[^/]+/comments$	10.946	0.108	0.136	0.147	752
0.3%	82	0	0	POST	^/api/posts/[^/]+/comments$	10.016	0.122	0.244	0.244	311
0.3%	52	0	0	POST	/api/auth/register	8.507	0.164	0.259	0.259	227
0.2%	23	0	0	DELETE	^/api/realmojis/[^/]+$	5.393	0.234	0.345	0.345	0
0.1%	33	0	0	GET	^/api/memories/[^/]+	4.497	0.136	0.451	0.451	961
0.1%	33	0	0	PUT	^/api/follows/[^/]+$	3.917	0.119	0.204	0.204	0
0.1%	34	0	0	POST	/api/auth/logout	3.865	0.114	0.151	0.151	0
0.1%	23	0	0	PATCH	^/api/posts/[^/]+$	2.723	0.118	0.173	0.173	616
0.1%	23	1	0	GET	/api/me	2.630	0.114	0.202	0.202	283
0.1%	18	0	0	PATCH	/api/me	2.198	0.122	0.229	0.229	264
0.1%	16	0	0	DELETE	^/api/comments/[^/]+$	1.813	0.113	0.151	0.151	0
0.0%	6	0	0	DELETE	^/api/posts/[^/]+$	1.458	0.243	0.319	0.319	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.0%	101	0	0	GET	/api/notifications	814.387	8.063	9.121	9.122	180
20.9%	57	0	45	PUT	/api/me/profile-image	460.441	8.078	8.254	8.254	13
12.3%	117	1	0	POST	/api/posts	269.701	2.305	6.224	6.384	610
8.2%	24	0	17	GET	/api/me/export	181.080	7.545	9.715	9.715	1187213
5.3%	827	5	0	GET	^/api/images/[^/]+	116.108	0.140	0.940	1.151	220087
5.0%	986	0	0	GET	^/api/feeds/[^/]+	110.041	0.112	0.162	0.211	4993
1.5%	263	0	0	GET	/api/memories	34.037	0.129	0.317	0.383	11069
1.3%	243	6	0	POST	/api/friend-requests	29.663	0.122	0.278	0.405	113
1.1%	225	7	0	GET	^/api/posts/[^/]+$	25.255	0.112	0.153	0.213	589
1.1%	188	0	0	POST	/api/auth/login	25.185	0.134	0.296	0.301	249
1.0%	162	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	22.892	0.141	0.322	0.377	116
0.8%	150	0	0	GET	/api/posting-window/current	16.935	0.113	0.203	0.203	240
0.7%	143	0	0	GET	^/api/relationships/[^/]+	15.585	0.109	0.158	0.230	1346
0.6%	113	0	0	GET	^/api/posts/me/current$	12.387	0.110	0.150	0.152	614
0.5%	100	0	0	GET	/api/users	10.934	0.109	0.149	0.161	234
0.5%	89	0	0	GET	^/api/posts/[^/]+/realmojis$	10.080	0.113	0.180	0.180	465
0.4%	78	0	0	GET	^/api/posts/[^/]+/comments$	8.793	0.113	0.203	0.203	661
0.3%	68	0	0	GET	^/api/users/[^/]+$	7.478	0.110	0.182	0.182	192
0.3%	41	0	0	POST	/api/auth/register	6.208	0.151	0.251	0.251	229
0.2%	35	0	0	POST	/api/auth/logout	4.361	0.125	0.205	0.205	0
0.1%	25	0	0	POST	^/api/posts/[^/]+/comments$	3.258	0.130	0.224	0.224	286
0.1%	17	0	0	POST	^/api/posts/[^/]+/realmojis$	2.979	0.175	0.404	0.404	429
0.1%	21	0	0	GET	^/api/memories/[^/]+	2.654	0.126	0.265	0.265	516
0.1%	16	0	0	PUT	^/api/follows/[^/]+$	1.832	0.115	0.206	0.206	0
0.1%	15	0	0	PATCH	^/api/posts/[^/]+$	1.828	0.122	0.183	0.183	644
0.1%	7	0	0	DELETE	^/api/posts/[^/]+$	1.669	0.238	0.305	0.305	0
0.1%	13	0	0	PATCH	/api/me	1.504	0.116	0.206	0.206	260
0.1%	12	1	0	GET	/api/me	1.275	0.106	0.125	0.125	257
0.0%	5	0	0	PUT	^/api/blocks/[^/]+$	0.774	0.155	0.228	0.228	0
0.0%	4	0	0	DELETE	^/api/comments/[^/]+$	0.470	0.117	0.129	0.129	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 29909 クエリ / 82 種類 / 合計 65.669s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	40.2%	82	26.423	322.23	872.48	7133.6	140.1	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	21.3%	1719	13.980	8.13	53.50	483.6	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	11.3%	1430	7.441	5.20	305.25	129.3	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	9.1%	3109	5.978	1.92	11.17	11.3	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	3.5%	1	2.320	2320.15	2320.15	29587.0	1199653.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	2.7%	1	1.760	1760.34	1760.34	20364.0	626266.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	2.2%	994	1.440	1.45	6.55	26.4	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
8	2.0%	582	1.339	2.30	13.84	28.5	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	1.7%	1	1.104	1104.10	1104.10	11528.0	529880.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
10	0.9%	513	0.569	1.11	7.54	19.8	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.5%	2	0.338	168.83	186.03	219904.0	2587.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
12	0.4%	1945	0.294	0.15	4.83	20.2	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
13	0.4%	1	0.271	271.17	271.17	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.4%	2	0.264	131.87	255.00	16770.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	0.4%	1	0.238	237.50	237.50	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 041040664263104d94ce7c7bf0c950ad4a7eeaa1
Type: cpu
Time: 2026-10-31 11:56:16 JST
Duration: 60.07s, Total samples = 45.87s (76.36%)
Showing nodes accounting for 36.90s, 80.44% of 45.87s total
Dropped 893 nodes (cum <= 0.23s)
Showing top 40 nodes out of 231
      flat  flat%   sum%        cum   cum%
     4.94s 10.77% 10.77%     15.63s 34.07%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     4.87s 10.62% 21.39%      7.21s 15.72%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     3.33s  7.26% 28.65%      3.46s  7.54%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     2.25s  4.91% 33.55%      2.25s  4.91%  internal/runtime/syscall/linux.Syscall6
     1.80s  3.92% 37.48%      1.80s  3.92%  runtime.memclrNoHeapPointers
     1.71s  3.73% 41.20%      2.44s  5.32%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     1.63s  3.55% 44.76%      6.49s 14.15%  runtime.scanObject
     1.59s  3.47% 48.22%      5.79s 12.62%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.58s  3.44% 51.67%      2.12s  4.62%  runtime.findObject
     1.53s  3.34% 55.00%      1.53s  3.34%  runtime.memmove
     1.39s  3.03% 58.03%      1.39s  3.03%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.31s  2.86% 60.89%      2.32s  5.06%  runtime.tryDeferToSpanScan
     1.21s  2.64% 63.53%      1.21s  2.64%  aeshashbody
     0.98s  2.14% 65.66%      0.98s  2.14%  crypto/internal/fips140/sha256.blockAVX2
     0.79s  1.72% 67.39%      0.79s  1.72%  runtime.spanClass.sizeclass (inline)
     0.66s  1.44% 68.82%      0.71s  1.55%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.61s  1.33% 70.15%      0.61s  1.33%  example.com/isureal-company/isureal.(*UserRow).active
     0.53s  1.16% 71.31%      2.04s  4.45%  runtime.mapaccess1_faststr
     0.40s  0.87% 72.18%      0.64s  1.40%  example.com/isureal-company/isureal.postLess
     0.38s  0.83% 73.01%      0.38s  0.83%  memeqbody
     0.29s  0.63% 73.64%      0.52s  1.13%  database/sql.convertAssignRows
     0.29s  0.63% 74.28%      0.29s  0.63%  indexbody
     0.29s  0.63% 74.91%      0.29s  0.63%  runtime.nextFreeFast (inline)
     0.27s  0.59% 75.50%      0.94s  2.05%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.52% 76.02%      1.80s  3.92%  runtime.mapassign_faststr
     0.22s  0.48% 76.50%      0.25s  0.55%  runtime.spanOf (inline)
     0.21s  0.46% 76.96%      0.52s  1.13%  runtime.greyobject
     0.20s  0.44% 77.39%      0.23s   0.5%  runtime.(*mspan).writeHeapBitsSmall
     0.20s  0.44% 77.83%      0.42s  0.92%  runtime.scanObjectsSmall
     0.17s  0.37% 78.20%      2.66s  5.80%  runtime.mallocgc
     0.14s  0.31% 78.50%      2.45s  5.34%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.14s  0.31% 78.81%      0.30s  0.65%  internal/runtime/maps.(*Iter).Next
     0.13s  0.28% 79.09%      1.26s  2.75%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.11s  0.24% 79.33%      4.11s  8.96%  example.com/isureal-company/isureal.loadBig
     0.11s  0.24% 79.57%      0.24s  0.52%  runtime.scanObjectSmall
     0.10s  0.22% 79.79%      1.10s  2.40%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.08s  0.17% 79.97%      0.37s  0.81%  runtime.mapaccess2_faststr
     0.08s  0.17% 80.14%      0.75s  1.64%  runtime.scanSpan
     0.07s  0.15% 80.29%      0.59s  1.29%  database/sql.(*Rows).scanLocked
     0.07s  0.15% 80.44%      1.07s  2.33%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 57.0% / 100% | 36% | 7.1% | 7 | 177 MB | CPU が飽和 |
| isu2 (2 cpu) | 1.8% / 6% | 0% | 0.3% | 3 | 390 MB |  |
| isu3 (2 cpu) | 2.2% / 6% | 0% | 0.4% | 2 | 236 MB |  |
| isu4 (2 cpu) | 49.8% / 92% | 2% | 13.8% | 6 | 108 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 77% / nginx 15% / systemd-journal 1% / journalctl 1% |
| isu2 | isureal 2% / nginx 1% |
| isu3 | isureal 2% / nginx 1% |
| isu4 | isureal 40% / postgres 28% |

## Diff from previous run

スコア: 58879 → 56553 (-2326)

`0904d44..773f053`

```
 webapp/go/limiter.go | 37 +++++++++++++++++++++++++++++++++++++
 webapp/go/server.go  | 10 +++++-----
 2 files changed, 42 insertions(+), 5 deletions(-)
```
