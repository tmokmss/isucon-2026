# run 20261031-110450-8f720c2

## Summary

| | |
|---|---|
| score | 15668 |
| status | pass |
| ref | origin/main |
| sha | 8f720c2f3486 |
| 所要時間 | 144s |
| 前回比 | +15668 (—) vs 20261031-110242-f264c69 |
| ベスト比 | -9133 (-36.8%) vs 20261031-103811-cef5fe3 = 24801 |
| note | main: インメモリ化 + 画像を手元から（#62）+ 署名 1 分単位 |

## Bench

```
{"id":"2618","status":"done","portalStatus":"finished","passed":true,"score":15668,"raw":15668,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:06:02.936Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:06:05.285Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:06:08.037Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:06:12.202Z level=INFO msg="asia 2 日目: 29 人"
time=2026-10-31T02:06:14.006Z level=INFO msg="europe 2 日目: 13 人"
time=2026-10-31T02:06:17.162Z level=INFO msg="america 2 日目: 12 人"
time=2026-10-31T02:06:20.836Z level=INFO msg="asia 3 日目: 62 人"
time=2026-10-31T02:06:23.998Z level=INFO msg="europe 3 日目: 30 人"
time=2026-10-31T02:06:24.600Z level=INFO msg="user moved" user=leoivymo5304 from=america to=europe
time=2026-10-31T02:06:27.818Z level=INFO msg="america 3 日目: 24 人"
time=2026-10-31T02:06:29.728Z level=INFO msg="user moved" user=solya_5510 from=asia to=america
time=2026-10-31T02:06:30.940Z level=INFO msg="asia 4 日目: 72 人"
time=2026-10-31T02:06:31.303Z level=INFO msg="user moved" user=leoivymo5304 from=europe to=america
time=2026-10-31T02:06:36.455Z level=INFO msg="europe 4 日目: 30 人"
time=2026-10-31T02:06:39.098Z level=INFO msg="america 4 日目: 26 人"
time=2026-10-31T02:06:41.281Z level=INFO msg="asia 5 日目: 72 人"
time=2026-10-31T02:06:42.896Z level=INFO msg="europe 5 日目: 30 人"
time=2026-10-31T02:06:50.390Z level=INFO msg="america 5 日目: 26 人"
time=2026-10-31T02:06:58.123Z level=INFO msg="europe 最終日（グランプリ）: 30 人"
time=2026-10-31T02:07:00.142Z level=INFO msg="america 最終日（グランプリ）: 26 人"
time=2026-10-31T02:07:01.961Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 15668,
  "messages": [
    "mekedan.5723: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "anmaxra5300: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "nenero_5076: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "asia の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.7%	259	1	0	GET	/api/notifications	3057.610	11.805	22.622	22.624	172
12.3%	2156	1430	0	POST	/api/auth/login	923.000	0.428	2.317	7.923	102
11.9%	230	2	0	POST	/api/posts	892.057	3.879	13.125	14.908	636
8.4%	2856	631	0	GET	^/api/images/[^/]+	632.167	0.221	1.707	7.838	82970
5.5%	258	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	415.716	1.611	7.735	10.365	116
5.4%	190	154	0	PUT	/api/me/profile-image	406.580	2.140	8.942	9.399	19
4.1%	458	18	0	POST	/api/friend-requests	309.056	0.675	6.037	13.045	112
2.7%	1302	0	0	GET	^/api/feeds/[^/]+	204.447	0.157	1.315	1.743	1195
1.3%	704	0	0	GET	/api/memories	100.828	0.143	1.273	2.262	10302
1.3%	89	1	0	POST	/api/auth/register	100.813	1.133	6.622	6.622	230
0.9%	59	0	0	GET	/api/me/export	66.660	1.130	14.838	14.838	1723599
0.8%	300	0	0	GET	/api/posting-window/current	62.359	0.208	0.824	1.491	237
0.7%	18	0	0	POST	/api/admin/posting-window	52.240	2.902	10.997	10.997	177
0.6%	227	3	0	GET	^/api/posts/me/current$	45.389	0.200	1.109	1.336	633
0.5%	138	0	0	POST	/api/auth/logout	35.793	0.259	3.046	4.914	0
0.4%	248	0	0	GET	^/api/relationships/[^/]+	33.394	0.135	1.392	1.510	1916
0.4%	28	2	0	PATCH	^/api/posts/[^/]+$	29.574	1.056	5.862	5.862	608
0.4%	54	0	0	PUT	^/api/follows/[^/]+$	27.351	0.506	5.970	5.970	0
0.4%	222	0	0	GET	/api/users	26.297	0.118	0.852	1.097	232
0.3%	189	10	0	GET	^/api/posts/[^/]+$	20.430	0.108	1.023	1.461	563
0.2%	107	0	0	GET	^/api/users/[^/]+$	13.140	0.123	1.512	1.516	179
0.2%	1	0	0	POST	/api/initialize	12.121	12.121	12.121	12.121	90
0.1%	9	0	0	DELETE	^/api/posts/[^/]+$	9.816	1.091	3.467	3.467	0
0.1%	30	0	0	PATCH	/api/me	8.311	0.277	1.622	1.622	299
0.1%	113	0	0	GET	^/api/memories/[^/]+	6.154	0.054	0.919	1.115	789
0.1%	42	0	0	POST	^/api/posts/[^/]+/comments$	5.436	0.129	2.664	2.664	321
0.0%	1	0	0	POST	^/api/posts/[^/]+/realmojis$	2.575	2.575	2.575	2.575	404
0.0%	42	2	0	GET	/api/me	2.503	0.060	0.414	0.414	279
0.0%	4	0	0	PUT	^/api/blocks/[^/]+$	2.397	0.599	2.377	2.377	0
0.0%	53	0	0	GET	^/api/posts/[^/]+/comments$	1.671	0.032	0.275	0.275	623
…（上位 30 / 全 37 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.2%	109	0	0	GET	/api/notifications	1059.225	9.718	15.150	15.150	179
19.3%	108	2	0	POST	/api/posts	463.649	4.293	12.946	13.732	632
7.7%	171	5	0	POST	/api/friend-requests	184.230	1.077	10.635	12.145	113
4.2%	128	0	0	POST	/api/auth/login	101.199	0.791	4.613	8.364	248
3.8%	356	0	0	GET	^/api/feeds/[^/]+	90.981	0.256	1.295	1.848	1261
3.1%	189	0	0	GET	/api/memories	75.429	0.399	1.895	9.964	9789
3.1%	467	2	0	GET	^/api/images/[^/]+	74.279	0.159	0.938	1.234	185757
2.8%	35	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	67.750	1.936	13.429	13.429	116
2.1%	139	0	0	GET	/api/posting-window/current	51.272	0.369	1.395	1.540	239
1.3%	84	23	0	GET	^/api/posts/me/current$	32.198	0.383	1.520	1.520	472
1.3%	8	0	0	GET	/api/me/export	31.398	3.925	12.471	12.471	1765654
1.0%	28	0	0	POST	/api/auth/register	23.623	0.844	2.635	2.635	228
0.8%	78	0	0	GET	/api/users	20.089	0.258	1.072	1.072	230
0.8%	15	0	0	PUT	/api/me/profile-image	18.563	1.238	4.993	4.993	0
0.7%	77	0	0	GET	^/api/relationships/[^/]+	16.454	0.214	1.321	1.321	844
0.7%	27	0	0	POST	/api/auth/logout	16.335	0.605	3.928	3.928	0
0.5%	30	0	0	POST	^/api/posts/[^/]+/realmojis$	12.012	0.400	3.262	3.262	407
0.4%	19	0	0	PUT	^/api/follows/[^/]+$	10.248	0.539	3.276	3.276	0
0.4%	44	3	0	GET	^/api/posts/[^/]+$	10.215	0.232	0.884	0.884	557
0.4%	9	5	0	PATCH	^/api/posts/[^/]+$	9.479	1.053	4.686	4.686	298
0.3%	25	0	0	GET	^/api/users/[^/]+$	6.479	0.259	1.164	1.164	167
0.2%	19	1	0	GET	/api/me	4.639	0.244	0.836	0.836	274
0.2%	23	0	0	GET	^/api/memories/[^/]+	4.293	0.187	0.486	0.486	554
0.1%	10	0	0	PATCH	/api/me	3.293	0.329	0.721	0.721	295
0.1%	21	0	0	GET	^/api/posts/[^/]+/realmojis$	2.957	0.141	0.319	0.319	435
0.1%	3	0	0	DELETE	^/api/realmojis/[^/]+$	2.687	0.896	1.653	1.653	0
0.1%	3	0	0	DELETE	^/api/posts/[^/]+$	2.612	0.871	1.923	1.923	0
0.1%	12	0	0	GET	^/api/posts/[^/]+/comments$	1.543	0.129	0.199	0.199	160
0.0%	2	0	0	PUT	^/api/blocks/[^/]+$	0.858	0.429	0.754	0.754	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.2%	92	0	0	GET	/api/notifications	899.742	9.780	12.014	12.014	180
17.8%	83	4	0	POST	/api/posts	361.214	4.352	15.428	15.428	599
8.5%	133	3	0	POST	/api/friend-requests	173.188	1.302	7.132	9.145	113
4.5%	94	0	0	POST	/api/auth/login	92.114	0.980	8.431	8.431	252
3.5%	272	0	0	GET	^/api/feeds/[^/]+	71.075	0.261	1.129	1.482	1456
3.3%	362	2	0	GET	^/api/images/[^/]+	67.686	0.187	0.954	0.957	187630
2.9%	145	0	0	GET	/api/memories	59.312	0.409	3.110	4.382	11634
2.5%	27	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	50.652	1.876	7.938	7.938	116
1.9%	111	0	0	GET	/api/posting-window/current	39.604	0.357	1.386	1.443	240
1.6%	12	0	0	GET	/api/me/export	31.581	2.632	7.475	7.475	3259078
1.1%	67	0	0	GET	/api/users	22.339	0.333	1.486	1.486	235
1.0%	66	14	0	GET	^/api/posts/me/current$	20.812	0.315	0.836	0.836	496
0.9%	64	0	0	GET	^/api/relationships/[^/]+	17.904	0.280	1.473	1.473	1022
0.8%	28	0	0	POST	^/api/posts/[^/]+/comments$	15.506	0.554	5.563	5.563	281
0.7%	24	0	0	POST	/api/auth/register	13.892	0.579	1.501	1.501	229
0.7%	9	3	0	PATCH	^/api/posts/[^/]+$	13.331	1.481	5.784	5.784	416
0.6%	13	0	0	POST	/api/auth/logout	12.994	1.000	4.987	4.987	0
0.6%	17	0	0	PUT	^/api/follows/[^/]+$	11.596	0.682	3.408	3.408	0
0.5%	41	3	0	GET	^/api/posts/[^/]+$	9.631	0.235	0.715	0.715	554
0.4%	11	0	0	PUT	/api/me/profile-image	7.367	0.670	3.843	3.843	0
0.4%	13	0	0	PATCH	/api/me	7.245	0.557	1.737	1.737	284
0.3%	38	0	0	GET	^/api/users/[^/]+$	6.637	0.175	0.519	0.519	170
0.3%	3	0	0	POST	^/api/posts/[^/]+/realmojis$	5.149	1.716	3.471	3.471	541
0.2%	3	0	0	PUT	^/api/blocks/[^/]+$	4.420	1.473	2.903	2.903	0
0.2%	13	0	0	DELETE	^/api/comments/[^/]+$	4.145	0.319	1.303	1.303	0
0.2%	10	0	0	GET	^/api/memories/[^/]+	3.406	0.341	1.036	1.036	1272
0.2%	2	0	0	DELETE	^/api/posts/[^/]+$	3.368	1.684	2.236	2.236	0
0.1%	1	0	0	DELETE	^/api/friendships/[^/]+$	2.463	2.463	2.463	2.463	0
0.1%	14	0	0	GET	/api/me	2.446	0.175	0.451	0.451	291
0.1%	17	0	0	GET	^/api/posts/[^/]+/comments$	2.133	0.125	0.199	0.199	769
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 12573 クエリ / 77 種類 / 合計 61.819s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	53.4%	79	33.035	418.16	3769.00	6104.4	124.7	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	10.7%	18	6.621	367.84	1843.28	32927.6	1655.9	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
3	8.8%	706	5.411	7.66	209.31	121.6	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	8.0%	512	4.947	9.66	51.84	763.4	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
5	5.2%	1729	3.214	1.86	7.75	12.2	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
6	3.8%	1	2.373	2372.89	2372.89	29320.0	1199409.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	2.5%	1	1.546	1546.37	1546.37	19577.0	625537.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	1.8%	1	1.093	1093.26	1093.26	11528.0	529666.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	1.1%	416	0.677	1.63	5.59	33.4	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.7%	2	0.409	204.44	251.78	214278.5	2520.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.5%	2	0.313	156.32	304.48	20239.0	2518.5	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.4%	1	0.230	229.91	229.91	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
13	0.4%	1	0.218	217.90	217.90	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.3%	1	0.212	212.18	212.18	69965.0	29853.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.2%	1	0.146	146.25	146.25	29382.0	61.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4;
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 440768d8ff35e48d773ba1e33cdc2b370b65294e
Type: cpu
Time: 2026-10-31 11:05:44 JST
Duration: 60.77s, Total samples = 83.22s (136.95%)
Showing nodes accounting for 73.72s, 88.58% of 83.22s total
Dropped 841 nodes (cum <= 0.42s)
Showing top 40 nodes out of 182
      flat  flat%   sum%        cum   cum%
    14.76s 17.74% 17.74%     14.76s 17.74%  crypto/internal/fips140/sha256.blockAVX2
    12.91s 15.51% 33.25%     12.91s 15.51%  runtime.spanClass.sizeclass (inline)
     6.80s  8.17% 41.42%     10.71s 12.87%  image/jpeg.(*decoder).reconstructBlock
     4.81s  5.78% 47.20%     18.52s 22.25%  runtime.tryDeferToSpanScan
     3.02s  3.63% 50.83%     23.74s 28.53%  image/jpeg.(*decoder).processSOS
     2.83s  3.40% 54.23%      4.56s  5.48%  image/jpeg.(*decoder).decodeHuffman
     2.57s  3.09% 57.32%      2.57s  3.09%  runtime.memclrNoHeapPointers
     2.32s  2.79% 60.11%      2.32s  2.79%  runtime.memmove
     2.22s  2.67% 62.77%      2.31s  2.78%  image/jpeg.(*decoder).receiveExtend
     2.06s  2.48% 65.25%     14.96s 17.98%  runtime.scanObject
     1.86s  2.24% 67.48%      8.23s  9.89%  runtime.scanObjectsSmall
     1.55s  1.86% 69.35%      2.02s  2.43%  image/jpeg.idctCols
     1.48s  1.78% 71.12%      1.48s  1.78%  internal/runtime/syscall/linux.Syscall6
     1.32s  1.59% 72.71%      1.87s  2.25%  image/jpeg.idctRows
     1.28s  1.54% 74.25%      1.81s  2.17%  image/jpeg.(*decoder).ensureNBits
     1.18s  1.42% 75.67%      5.39s  6.48%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.02s  1.23% 76.89%      1.02s  1.23%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.01s  1.21% 78.11%      1.02s  1.23%  image/jpeg.dctBox (inline)
     0.95s  1.14% 79.25%      0.95s  1.14%  aeshashbody
     0.91s  1.09% 80.34%     16.61s 19.96%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.73s  0.88% 81.22%     16.14s 19.39%  crypto/internal/fips140/sha256.(*Digest).Write
     0.62s  0.75% 81.96%      0.62s  0.75%  memeqbody
     0.61s  0.73% 82.70%     17.45s 20.97%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.56s  0.67% 83.37%      9.47s 11.38%  runtime.scanSpan
     0.52s  0.62% 83.99%      0.52s  0.62%  runtime.extractHeapBitsSmall
     0.49s  0.59% 84.58%      0.52s  0.62%  image/jpeg.(*decoder).readByteStuffedByte
     0.43s  0.52% 85.10%      0.50s   0.6%  runtime.(*spanInlineMarkBits).tryAcquire
     0.40s  0.48% 85.58%      2.15s  2.58%  runtime.mapaccess1_faststr
     0.39s  0.47% 86.05%     18.50s 22.23%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.30s  0.36% 86.41%      0.86s  1.03%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.30s  0.36% 86.77%      0.52s  0.62%  example.com/isureal-company/isureal.postLess
     0.29s  0.35% 87.12%      0.80s  0.96%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.21s  0.25% 87.37%     19.88s 23.89%  example.com/isureal-company/isureal.imageSignature
     0.20s  0.24% 87.61%     14.96s 17.98%  crypto/internal/fips140/sha256.block
     0.18s  0.22% 87.83%      0.56s  0.67%  database/sql.convertAssignRows
     0.16s  0.19% 88.02%     16.27s 19.55%  runtime.mallocgc
     0.13s  0.16% 88.18%      0.89s  1.07%  runtime.mallocgcSmallScanNoHeader
     0.12s  0.14% 88.32%      0.58s   0.7%  runtime.scanObjectSmall
     0.11s  0.13% 88.45%      1.13s  1.36%  runtime.mapassign_faststr
     0.11s  0.13% 88.58%      0.47s  0.56%  runtime.sweepone
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 74.9% / 100% | 62% | 0.1% | 6 | 107 MB | CPU が飽和 |
| isu2 (2 cpu) | 2.6% / 51% | 0% | 0.2% | 5 | 264 MB | run queue が積もっている |
| isu3 (2 cpu) | 2.6% / 51% | 0% | 0.1% | 2 | 200 MB |  |
| isu4 (2 cpu) | 29.1% / 88% | 0% | 17.8% | 5 | 109 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に find）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 134% / nginx 11% |
| isu2 | find 2% |
| isu3 | find 2% / isureal 1% |
| isu4 | postgres 18% |

## Diff from previous run

スコア: 0 → 15668 (+15668)

`f264c69..8f720c2`

```
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |  13 +
 conf/isu2/home/isucon/env.sh                   |   2 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |  13 +
 conf/isu3/home/isucon/env.sh                   |   2 +
 webapp/go/cmd/image-export/main.go             |  69 ++++
 webapp/go/cmd/isureal/main.go                  |  16 +-
 webapp/go/image_local.go                       | 511 +++++++++++++++++++++++++
 webapp/go/image_signature.go                   |   4 +-
 webapp/go/initialize.go                        |   1 +
 9 files changed, 620 insertions(+), 11 deletions(-)
```
