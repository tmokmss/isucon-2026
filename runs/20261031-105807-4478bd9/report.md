# run 20261031-105807-4478bd9

## Summary

| | |
|---|---|
| score | 15531 |
| status | pass |
| ref | HEAD |
| sha | 4478bd9d6667 |
| 所要時間 | 167s |
| 前回比 | -9270 (-37.4%) vs 20261031-103811-cef5fe3 |
| ベスト比 | -9270 (-37.4%) vs 20261031-103811-cef5fe3 = 24801 |
| note | 画像: europe1・america1 で手元から返す（書き出し途中、ベンチ中は書き出し停止） |

## Bench

```
{"id":"2245","status":"done","portalStatus":"finished","passed":true,"score":15531,"raw":15531,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:59:41.198Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:59:43.811Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:59:45.762Z level=INFO msg="user moved" user=nite.5022 from=asia to=america
time=2026-10-31T01:59:46.691Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T01:59:50.217Z level=INFO msg="asia 2 日目: 40 人"
time=2026-10-31T01:59:52.780Z level=INFO msg="europe 2 日目: 11 人"
time=2026-10-31T01:59:55.134Z level=INFO msg="user moved" user=soku_5290 from=asia to=europe
time=2026-10-31T01:59:55.584Z level=INFO msg="america 2 日目: 12 人"
time=2026-10-31T01:59:59.739Z level=INFO msg="asia 3 日目: 97 人"
time=2026-10-31T02:00:02.351Z level=INFO msg="europe 3 日目: 23 人"
time=2026-10-31T02:00:05.573Z level=INFO msg="america 3 日目: 22 人"
time=2026-10-31T02:00:12.747Z level=INFO msg="asia 4 日目: 109 人"
time=2026-10-31T02:00:16.001Z level=INFO msg="europe 4 日目: 23 人"
time=2026-10-31T02:00:24.614Z level=INFO msg="europe 5 日目: 23 人"
time=2026-10-31T02:00:26.929Z level=INFO msg="user moved" user=temero5778 from=asia to=america
time=2026-10-31T02:00:26.931Z level=INFO msg="asia 5 日目: 108 人"
time=2026-10-31T02:00:27.787Z level=INFO msg="america 5 日目: 23 人"
time=2026-10-31T02:00:38.971Z level=INFO msg="america 最終日（グランプリ）: 23 人"
time=2026-10-31T02:00:40.564Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 15531,
  "messages": [
    "america の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました",
    "danka5924: 開いていない投稿期間 83f9189d-3f6c-4546-894f-50d72ca3d2ae が通知されました",
    "nite.5022: 投稿が開いていない投稿期間 83f9189d-3f6c-4546-894f-50d72ca3d2ae に入りました",
    "mushi_5855: 開いていない投稿期間 83f9189d-3f6c-4546-894f-50d72ca3d2ae が通知されました",
    "kita.5891: 開いていない投稿期間 83f9189d-3f6c-4546-894f-50d72ca3d2ae が通知されました",
    "kechi5909: 開いていない投稿期間 83f9189d-3f6c-4546-894f-50d72ca3d2ae が通知されました",
    "martoel5871: 投稿が開いていない投稿期間 83f9189d-3f6c-4546-894f-50d72ca3d2ae に入りました",
    "leosota5910: 投稿が開いていない投稿期間 83f9189d-3f6c-4546-894f-50d72ca3d2ae に入りました",
    "hokamax-5879: 開いていない投稿期間 83f9189d-3f6c-4546-894f-50d72ca3d2ae が通知されました",
    "asia の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.7%	359	0	0	GET	/api/notifications	4542.146	12.652	18.472	19.247	175
14.6%	295	8	0	POST	/api/posts	1585.872	5.376	15.092	15.773	635
11.0%	2187	1351	0	POST	/api/auth/login	1193.588	0.546	4.288	10.142	112
10.2%	3683	930	0	GET	^/api/images/[^/]+	1107.318	0.301	2.722	9.635	87388
3.9%	199	132	0	PUT	/api/me/profile-image	421.900	2.120	13.519	14.288	15
3.5%	554	22	0	POST	/api/friend-requests	379.156	0.684	8.797	13.191	112
3.5%	275	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	378.154	1.375	10.814	13.210	115
2.1%	2070	0	0	GET	^/api/feeds/[^/]+	230.572	0.111	0.707	1.910	1353
1.8%	127	1	0	POST	/api/auth/register	198.277	1.561	9.240	9.495	229
1.3%	839	0	0	GET	/api/memories	143.154	0.171	1.514	2.966	10350
1.1%	93	0	0	PUT	^/api/follows/[^/]+$	118.183	1.271	10.975	10.975	0
0.9%	40	0	0	GET	/api/me/export	98.730	2.468	12.999	12.999	2304209
0.8%	396	0	0	GET	/api/posting-window/current	85.596	0.216	1.206	2.183	237
0.7%	18	0	0	POST	/api/admin/posting-window	73.000	4.056	12.916	12.916	176
0.5%	244	8	0	GET	^/api/posts/me/current$	53.528	0.219	1.052	1.489	627
0.3%	62	0	0	POST	^/api/posts/[^/]+/realmojis$	32.681	0.527	4.592	4.592	452
0.3%	307	0	0	GET	^/api/relationships/[^/]+	32.209	0.105	0.542	1.107	1520
0.3%	322	7	0	GET	^/api/posts/[^/]+$	27.796	0.086	0.408	0.985	606
0.2%	243	0	0	GET	/api/users	24.433	0.101	0.646	0.838	244
0.2%	26	1	0	POST	^/api/posts/[^/]+/comments$	21.271	0.818	6.866	6.866	375
0.2%	141	0	0	POST	/api/auth/logout	18.732	0.133	1.220	2.736	0
0.2%	28	0	0	PATCH	^/api/posts/[^/]+$	18.181	0.649	4.255	4.255	643
0.1%	52	0	0	PATCH	/api/me	15.958	0.307	2.529	2.529	283
0.1%	171	0	0	GET	^/api/users/[^/]+$	13.181	0.077	0.334	0.334	181
0.1%	1	0	0	POST	/api/initialize	12.538	12.538	12.538	12.538	90
0.1%	8	0	0	DELETE	^/api/posts/[^/]+$	12.127	1.516	3.662	3.662	0
0.1%	7	0	0	PUT	^/api/admin/users/[^/]+/certification$	7.260	1.037	7.240	7.240	0
0.1%	135	0	0	GET	^/api/memories/[^/]+	7.084	0.052	0.812	1.136	731
0.1%	74	2	0	GET	/api/me	6.855	0.093	0.588	0.588	290
0.0%	11	0	0	PUT	^/api/blocks/[^/]+$	5.062	0.460	2.095	2.095	0
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
50.0%	90	0	0	GET	/api/notifications	1028.095	11.423	19.995	19.995	179
16.4%	71	0	0	POST	/api/posts	338.171	4.763	16.200	16.200	614
5.3%	136	4	0	POST	/api/friend-requests	108.458	0.797	6.667	6.854	112
5.2%	282	0	0	GET	^/api/feeds/[^/]+	106.634	0.378	1.750	7.471	1322
4.8%	110	0	0	POST	/api/auth/login	98.731	0.898	6.835	7.844	250
3.1%	384	2	0	GET	^/api/images/[^/]+	63.247	0.165	0.943	0.951	171267
2.7%	153	0	0	GET	/api/memories	54.654	0.357	2.803	2.830	9777
1.9%	32	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	38.754	1.211	14.666	14.666	116
1.7%	93	0	0	GET	/api/posting-window/current	34.122	0.367	2.491	2.491	238
1.4%	20	0	0	POST	/api/auth/register	28.464	1.423	3.076	3.076	228
1.2%	65	19	0	GET	^/api/posts/me/current$	24.739	0.381	1.618	1.618	443
1.0%	12	0	0	PUT	^/api/follows/[^/]+$	21.475	1.790	11.286	11.286	0
1.0%	9	0	0	GET	/api/me/export	20.627	2.292	9.031	9.031	1984130
0.8%	64	0	0	GET	/api/users	16.745	0.262	1.714	1.714	230
0.7%	60	0	0	GET	^/api/relationships/[^/]+	13.918	0.232	0.863	0.863	1173
0.6%	18	0	0	POST	/api/auth/logout	12.183	0.677	5.187	5.187	0
0.6%	1	0	0	POST	/api/me/account-deletion	11.552	11.552	11.552	11.552	0
0.4%	26	0	0	GET	^/api/users/[^/]+$	7.401	0.285	1.336	1.336	186
0.3%	24	1	0	GET	^/api/posts/[^/]+$	6.389	0.266	0.725	0.725	581
0.2%	11	3	0	PATCH	^/api/posts/[^/]+$	3.989	0.363	1.058	1.058	443
0.1%	18	0	0	GET	^/api/memories/[^/]+	3.056	0.170	0.281	0.281	919
0.1%	7	0	0	PATCH	/api/me	2.981	0.426	0.684	0.684	236
0.1%	5	0	0	PUT	/api/me/profile-image	2.757	0.551	0.878	0.878	0
0.1%	1	0	0	POST	^/api/posts/[^/]+/realmojis$	2.682	2.682	2.682	2.682	409
0.1%	14	0	0	GET	^/api/posts/[^/]+/comments$	2.010	0.144	0.264	0.264	435
0.1%	9	1	0	GET	/api/me	1.559	0.173	0.321	0.321	236
0.1%	10	0	0	POST	^/api/posts/[^/]+/comments$	1.446	0.145	0.209	0.209	276
0.1%	2	0	0	DELETE	^/api/posts/[^/]+$	1.333	0.666	0.847	0.847	0
0.1%	7	0	0	GET	^/api/posts/[^/]+/realmojis$	1.148	0.164	0.308	0.308	484
0.0%	4	0	0	DELETE	^/api/comments/[^/]+$	0.680	0.170	0.312	0.312	0
…（上位 30 / 全 31 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
47.1%	84	0	0	GET	/api/notifications	802.514	9.554	16.261	16.261	182
17.5%	53	2	0	POST	/api/posts	298.014	5.623	14.705	14.705	594
7.3%	96	2	0	POST	/api/friend-requests	125.063	1.303	14.218	14.218	113
5.0%	97	0	0	POST	/api/auth/login	86.050	0.887	7.777	7.777	252
3.3%	35	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	56.995	1.628	6.658	6.658	116
3.2%	359	3	0	GET	^/api/images/[^/]+	54.961	0.153	0.956	0.974	132891
2.6%	192	0	0	GET	^/api/feeds/[^/]+	43.633	0.227	1.115	1.132	1252
2.1%	122	0	0	GET	/api/memories	35.038	0.287	1.702	2.462	12291
2.0%	17	0	0	PUT	^/api/follows/[^/]+$	34.133	2.008	14.243	14.243	0
1.9%	72	0	0	GET	/api/posting-window/current	32.254	0.448	1.747	1.747	238
1.8%	7	0	0	GET	/api/me/export	30.545	4.364	16.295	16.295	4598609
1.0%	22	0	0	POST	^/api/posts/[^/]+/comments$	17.334	0.788	9.289	9.289	282
0.7%	11	0	0	POST	/api/auth/logout	12.325	1.120	4.380	4.380	0
0.6%	54	0	0	GET	^/api/relationships/[^/]+	10.935	0.203	0.902	0.902	1316
0.6%	43	0	0	GET	/api/users	10.527	0.245	1.433	1.433	233
0.6%	45	4	0	GET	^/api/posts/[^/]+$	10.436	0.232	0.840	0.840	571
0.6%	20	0	0	POST	/api/auth/register	9.536	0.477	1.274	1.274	230
0.4%	31	1	0	GET	^/api/posts/me/current$	7.237	0.233	1.300	1.300	593
0.4%	9	0	0	PUT	/api/me/profile-image	6.988	0.776	1.495	1.495	0
0.3%	16	0	0	GET	^/api/memories/[^/]+	4.407	0.275	1.146	1.146	883
0.2%	12	0	0	GET	/api/me	2.680	0.223	0.564	0.564	279
0.1%	12	0	0	GET	^/api/users/[^/]+$	2.174	0.181	0.284	0.284	185
0.1%	5	0	0	PATCH	/api/me	1.904	0.381	0.735	0.735	313
0.1%	12	0	0	GET	^/api/posts/[^/]+/comments$	1.876	0.156	0.255	0.255	512
0.1%	3	0	0	DELETE	^/api/posts/[^/]+$	1.775	0.592	1.299	1.299	0
0.1%	3	0	0	PATCH	^/api/posts/[^/]+$	1.524	0.508	0.857	0.857	623
0.1%	2	0	0	PUT	^/api/blocks/[^/]+$	1.458	0.729	1.350	1.350	0
0.1%	3	0	0	DELETE	^/api/comments/[^/]+$	0.940	0.313	0.491	0.491	0
0.0%	3	0	0	GET	^/api/posts/[^/]+/realmojis$	0.716	0.239	0.275	0.275	29
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 13665 クエリ / 81 種類 / 合計 58.331s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	53.0%	56	30.900	551.78	3061.66	7471.7	161.0	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	10.6%	17	6.205	365.00	2500.71	34507.4	1756.0	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
3	7.5%	662	4.395	6.64	205.23	126.7	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	7.4%	556	4.289	7.71	39.05	686.5	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
5	6.4%	2053	3.755	1.83	6.01	12.5	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
6	4.1%	1	2.394	2394.44	2394.44	29299.0	1199382.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	2.9%	1	1.720	1719.53	1719.53	19464.0	625542.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	1.9%	1	1.105	1105.12	1105.12	11528.0	529686.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	1.2%	412	0.675	1.64	4.99	32.8	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.6%	2	0.349	174.70	199.53	214066.0	2518.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.5%	2	0.315	157.74	307.70	20239.0	2513.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.4%	1	0.245	245.20	245.20	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
13	0.4%	1	0.211	211.45	211.45	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.4%	1	0.206	205.67	205.67	69935.0	29823.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.3%	63	0.187	2.98	5.58	53.7	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)

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
Build ID: 10680fad2f1b13103cad26d3492e4a9c712de06d
Type: cpu
Time: 2026-10-31 10:59:22 JST
Duration: 60.36s, Total samples = 83.24s (137.90%)
Showing nodes accounting for 73.29s, 88.05% of 83.24s total
Dropped 786 nodes (cum <= 0.42s)
Showing top 40 nodes out of 180
      flat  flat%   sum%        cum   cum%
    15.55s 18.68% 18.68%     15.55s 18.68%  crypto/internal/fips140/sha256.blockAVX2
    12.87s 15.46% 34.14%     12.87s 15.46%  runtime.spanClass.sizeclass (inline)
     6.65s  7.99% 42.13%     10.39s 12.48%  image/jpeg.(*decoder).reconstructBlock
     4.32s  5.19% 47.32%     18.18s 21.84%  runtime.tryDeferToSpanScan
     2.93s  3.52% 50.84%     22.57s 27.11%  image/jpeg.(*decoder).processSOS
     2.50s  3.00% 53.84%      2.50s  3.00%  runtime.memclrNoHeapPointers
     2.49s  2.99% 56.84%      3.98s  4.78%  image/jpeg.(*decoder).decodeHuffman
     2.31s  2.78% 59.61%      2.31s  2.78%  runtime.memmove
     1.89s  2.27% 61.88%      8.35s 10.03%  runtime.scanObjectsSmall
     1.88s  2.26% 64.14%      1.97s  2.37%  image/jpeg.(*decoder).receiveExtend
     1.65s  1.98% 66.12%     14.52s 17.44%  runtime.scanObject
     1.58s  1.90% 68.02%      1.58s  1.90%  internal/runtime/syscall/linux.Syscall6
     1.39s  1.67% 69.69%      1.92s  2.31%  image/jpeg.idctCols
     1.37s  1.65% 71.34%      5.76s  6.92%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.27s  1.53% 72.86%      1.79s  2.15%  image/jpeg.idctRows
     1.15s  1.38% 74.24%      1.15s  1.38%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.11s  1.33% 75.58%      1.53s  1.84%  image/jpeg.(*decoder).ensureNBits
     1.04s  1.25% 76.83%      1.04s  1.25%  aeshashbody
     1.02s  1.23% 78.05%      1.02s  1.23%  image/jpeg.dctBox (inline)
     0.83s     1% 79.05%     16.99s 20.41%  crypto/internal/fips140/sha256.(*Digest).Write
     0.83s     1% 80.05%     17.27s 20.75%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.77s  0.93% 80.97%      0.77s  0.93%  memeqbody
     0.73s  0.88% 81.85%      9.82s 11.80%  runtime.scanSpan
     0.63s  0.76% 82.60%      0.63s  0.76%  runtime.extractHeapBitsSmall
     0.58s   0.7% 83.30%     18.14s 21.79%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.56s  0.67% 83.97%     19.55s 23.49%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.44s  0.53% 84.50%      0.56s  0.67%  runtime.(*spanInlineMarkBits).tryAcquire
     0.39s  0.47% 84.97%      0.42s   0.5%  image/jpeg.(*decoder).readByteStuffedByte
     0.38s  0.46% 85.43%      0.54s  0.65%  example.com/isureal-company/isureal.postLess
     0.31s  0.37% 85.80%      2.24s  2.69%  runtime.mapaccess1_faststr
     0.25s   0.3% 86.10%      0.92s  1.11%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.25s   0.3% 86.40%     20.97s 25.19%  example.com/isureal-company/isureal.imageSignature
     0.22s  0.26% 86.67%      0.69s  0.83%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.21s  0.25% 86.92%      0.57s  0.68%  database/sql.convertAssignRows
     0.20s  0.24% 87.16%     14.54s 17.47%  runtime.mallocgc
     0.17s   0.2% 87.36%      0.99s  1.19%  runtime.mallocgcSmallScanNoHeader
     0.16s  0.19% 87.55%      0.67s   0.8%  runtime.scanObjectSmall
     0.15s  0.18% 87.73%      2.55s  3.06%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.15s  0.18% 87.91%      0.46s  0.55%  runtime.sweepone
     0.11s  0.13% 88.05%      1.42s  1.71%  runtime.mapassign_faststr
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 78.2% / 100% | 66% | 0.3% | 7 | 108 MB | CPU が飽和 |
| isu2 (2 cpu) | 2.6% / 33% | 0% | 0.3% | 1 | 189 MB |  |
| isu3 (2 cpu) | 8.1% / 35% | 0% | 2.5% | 4 | 115 MB |  |
| isu4 (2 cpu) | 26.6% / 73% | 0% | 16.4% | 5 | 108 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 134% / nginx 11% |
| isu2 | 目立つプロセス無し |
| isu3 | postgres 9% |
| isu4 | postgres 18% |

## Diff from previous run

スコア: 24801 → 15531 (-9270)

`cef5fe3..4478bd9`

```
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |  13 +
 conf/isu2/home/isucon/env.sh                   |   2 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |  13 +
 conf/isu3/home/isucon/env.sh                   |   2 +
 webapp/go/account_deletion_api.go              |   1 +
 webapp/go/cmd/image-export/main.go             |  69 ++++
 webapp/go/cmd/isureal/main.go                  |  16 +-
 webapp/go/image_api.go                         |   6 +
 webapp/go/image_local.go                       | 511 +++++++++++++++++++++++++
 webapp/go/initialize.go                        |   1 +
 webapp/go/post_api.go                          |   1 +
 webapp/go/profile_image_api.go                 |   4 +
 webapp/go/realmoji_api.go                      |   4 +
 13 files changed, 634 insertions(+), 9 deletions(-)
```
