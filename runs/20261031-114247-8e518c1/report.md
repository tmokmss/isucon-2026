# run 20261031-114247-8e518c1

## Summary

| | |
|---|---|
| score | 26395 |
| status | pass |
| ref | limiter-test-reg |
| sha | 8e518c1e77cf |
| 所要時間 | 433s |
| 前回比 | +9323 (54.6%) vs 20261031-113441-6a7a33a |
| ベスト比 | +1594 (6.4%) vs 20261031-103811-cef5fe3 = 24801 **ベスト更新** |
| note | limiter post4 mem2 heavy1 + register fix 1376dd7 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	134	134	134	133
europe	7	17	25	25	25	24
america	6	15	15	15	15	15
```

```
{"id":"5454","status":"done","portalStatus":"finished","passed":true,"score":26395,"raw":26395,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:48:47.536Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:48:50.471Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:48:53.486Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:48:56.500Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T02:48:59.571Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T02:49:01.591Z level=INFO msg="user moved" user=suseru.5364 from=asia to=europe
time=2026-10-31T02:49:02.524Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T02:49:05.491Z level=INFO msg="asia 3 日目: 134 人"
time=2026-10-31T02:49:08.660Z level=INFO msg="europe 3 日目: 25 人"
time=2026-10-31T02:49:11.593Z level=INFO msg="america 3 日目: 15 人"
time=2026-10-31T02:49:14.751Z level=INFO msg="asia 4 日目: 134 人"
time=2026-10-31T02:49:17.597Z level=INFO msg="europe 4 日目: 25 人"
time=2026-10-31T02:49:20.527Z level=INFO msg="america 4 日目: 15 人"
time=2026-10-31T02:49:23.547Z level=INFO msg="asia 5 日目: 134 人"
time=2026-10-31T02:49:26.565Z level=INFO msg="europe 5 日目: 25 人"
time=2026-10-31T02:49:29.504Z level=INFO msg="america 5 日目: 15 人"
time=2026-10-31T02:49:38.471Z level=INFO msg="america 最終日（グランプリ）: 15 人"
time=2026-10-31T02:49:38.521Z level=INFO msg="asia 最終日（グランプリ）: 133 人"
time=2026-10-31T02:49:38.563Z level=INFO msg="europe 最終日（グランプリ）: 24 人"
time=2026-10-31T02:49:47.465Z level=INFO msg="グランプリ: アジア 84 人, ヨーロッパ 17 人, アメリカ 8 人"
{
  "pass": true,
  "score": 26395,
  "messages": [
    "nukino5371: プロフィール画像を変えられません: PUT /api/me/profile-image: ステータスコード 503 が返りました（期待値 204、busy）（5 回試行）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.7%	380	0	0	GET	/api/notifications	3538.896	9.313	14.962	14.962	175
28.4%	375	80	238	PUT	/api/me/profile-image	2736.947	7.299	9.189	9.900	18
21.0%	431	1	29	POST	/api/posts	2020.857	4.689	8.674	8.802	591
7.6%	108	0	65	GET	/api/me/export	734.008	6.796	9.399	9.581	729604
0.9%	2832	210	0	GET	^/api/images/[^/]+	90.990	0.032	0.152	0.389	176963
0.8%	1206	478	0	POST	/api/auth/login	77.903	0.065	0.350	0.530	159
0.8%	3260	0	0	GET	^/api/feeds/[^/]+	77.700	0.024	0.100	0.193	2720
0.7%	534	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	69.380	0.130	0.511	0.677	116
0.5%	722	22	0	POST	/api/friend-requests	47.848	0.066	0.577	0.860	112
0.5%	874	0	0	GET	/api/memories	45.560	0.052	0.372	0.528	8249
0.3%	172	0	0	POST	^/api/posts/[^/]+/realmojis$	32.606	0.190	0.827	0.835	499
0.2%	225	1	0	POST	^/api/posts/[^/]+/comments$	16.777	0.075	0.494	0.518	317
0.2%	647	5	0	GET	^/api/posts/[^/]+$	15.635	0.024	0.114	0.146	613
0.1%	564	0	0	GET	/api/posting-window/current	14.393	0.026	0.110	0.121	237
0.1%	490	0	0	GET	^/api/relationships/[^/]+	13.086	0.027	0.092	0.128	1442
0.1%	1	0	0	POST	/api/initialize	12.550	12.550	12.550	12.550	90
0.1%	151	1	0	POST	/api/auth/register	10.206	0.068	0.366	0.393	228
0.1%	379	6	0	GET	^/api/posts/me/current$	9.527	0.025	0.097	0.111	623
0.1%	332	0	0	GET	^/api/users/[^/]+$	7.816	0.024	0.097	0.125	182
0.1%	304	0	0	GET	/api/users	7.332	0.024	0.111	0.169	207
0.1%	286	0	0	GET	^/api/posts/[^/]+/realmojis$	6.539	0.023	0.100	0.125	494
0.1%	35	0	0	DELETE	^/api/realmojis/[^/]+$	6.351	0.181	0.327	0.327	0
0.1%	134	0	0	POST	/api/auth/logout	6.236	0.047	0.172	0.193	0
0.1%	268	0	0	GET	^/api/posts/[^/]+/comments$	6.144	0.023	0.094	0.126	851
0.1%	77	0	0	PUT	^/api/follows/[^/]+$	5.115	0.066	0.693	0.693	0
0.0%	57	0	0	DELETE	^/api/comments/[^/]+$	3.480	0.061	0.157	0.157	0
0.0%	123	0	0	GET	^/api/memories/[^/]+	3.444	0.028	0.249	0.407	658
0.0%	16	0	0	DELETE	^/api/posts/[^/]+$	3.081	0.193	0.282	0.282	0
0.0%	52	0	0	PATCH	/api/me	2.607	0.050	0.156	0.156	264
0.0%	45	2	0	PATCH	^/api/posts/[^/]+$	2.483	0.055	0.187	0.187	611
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.9%	100	0	0	GET	/api/notifications	873.179	8.732	11.930	11.931	179
25.2%	92	1	3	POST	/api/posts	479.434	5.211	8.743	8.743	588
6.4%	16	0	9	GET	/api/me/export	122.229	7.639	9.632	9.632	1129497
4.7%	12	0	6	PUT	/api/me/profile-image	89.568	7.464	8.209	8.209	9
4.5%	618	2	0	GET	^/api/images/[^/]+	85.831	0.139	0.928	0.950	138696
2.4%	370	0	0	GET	^/api/feeds/[^/]+	45.713	0.124	0.205	0.240	1684
2.1%	176	0	0	POST	/api/auth/login	39.695	0.226	0.537	0.557	250
1.9%	215	0	0	GET	/api/memories	36.378	0.169	0.796	1.215	12249
1.2%	146	5	0	POST	/api/friend-requests	23.402	0.160	0.567	0.729	112
0.9%	127	0	0	GET	/api/posting-window/current	16.960	0.134	0.251	0.309	239
0.6%	57	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	12.228	0.215	0.803	0.803	116
0.5%	85	4	0	GET	^/api/posts/me/current$	10.085	0.119	0.214	0.214	589
0.5%	37	0	0	POST	^/api/posts/[^/]+/realmojis$	9.569	0.259	0.560	0.560	443
0.4%	69	0	0	GET	^/api/relationships/[^/]+	8.512	0.123	0.187	0.187	1314
0.4%	66	3	0	GET	^/api/posts/[^/]+$	7.971	0.121	0.218	0.218	566
0.4%	60	0	0	GET	/api/users	7.499	0.125	0.279	0.279	215
0.3%	37	0	0	POST	/api/auth/logout	6.225	0.168	0.378	0.378	0
0.3%	40	0	0	GET	^/api/posts/[^/]+/realmojis$	5.032	0.126	0.193	0.193	466
0.3%	36	0	0	GET	^/api/memories/[^/]+	4.941	0.137	0.251	0.251	623
0.2%	22	0	0	POST	/api/auth/register	4.727	0.215	0.384	0.384	228
0.2%	16	0	0	PUT	^/api/follows/[^/]+$	3.182	0.199	0.689	0.689	0
0.1%	22	0	0	GET	^/api/users/[^/]+$	2.794	0.127	0.161	0.161	168
0.1%	11	2	0	GET	/api/me	1.356	0.123	0.202	0.202	223
0.1%	10	0	0	GET	^/api/posts/[^/]+/comments$	1.346	0.135	0.231	0.231	745
0.1%	9	2	0	PATCH	^/api/posts/[^/]+$	1.292	0.144	0.280	0.280	479
0.1%	4	0	0	DELETE	^/api/realmojis/[^/]+$	1.144	0.286	0.396	0.396	0
0.1%	6	0	0	PATCH	/api/me	1.095	0.182	0.275	0.275	237
0.0%	3	0	0	DELETE	^/api/posts/[^/]+$	0.819	0.273	0.318	0.318	0
0.0%	1	0	0	DELETE	^/api/follows/[^/]+$	0.380	0.380	0.380	0.380	0
0.0%	1	0	0	DELETE	^/api/friendships/[^/]+$	0.322	0.322	0.322	0.322	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.3%	69	0	0	GET	/api/notifications	563.625	8.168	9.076	9.076	179
27.1%	73	0	1	POST	/api/posts	344.771	4.723	8.968	8.968	608
7.1%	13	0	6	GET	/api/me/export	90.120	6.932	9.991	9.991	1803462
3.9%	294	3	0	GET	^/api/images/[^/]+	49.108	0.167	0.951	0.983	265234
3.8%	6	0	2	PUT	/api/me/profile-image	48.674	8.112	8.247	8.247	6
3.2%	316	0	0	GET	^/api/feeds/[^/]+	40.886	0.129	0.229	0.277	2161
1.6%	113	7	0	POST	/api/friend-requests	20.318	0.180	0.777	0.834	110
1.5%	119	0	0	GET	/api/memories	19.055	0.160	0.570	0.828	8010
1.2%	55	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	15.579	0.283	0.548	0.548	116
1.1%	61	0	0	POST	/api/auth/login	13.476	0.221	0.423	0.423	251
0.9%	90	0	0	GET	/api/posting-window/current	11.982	0.133	0.248	0.248	240
0.7%	68	7	0	GET	^/api/posts/me/current$	8.345	0.123	0.239	0.239	553
0.6%	61	7	0	GET	^/api/posts/[^/]+$	7.645	0.125	0.177	0.177	546
0.6%	36	0	0	POST	^/api/posts/[^/]+/comments$	7.114	0.198	0.504	0.504	277
0.6%	55	0	0	GET	^/api/relationships/[^/]+	7.107	0.129	0.214	0.214	1628
0.5%	46	0	0	GET	/api/users	5.907	0.128	0.213	0.213	212
0.2%	14	0	0	POST	/api/auth/register	3.111	0.222	0.331	0.331	229
0.2%	24	0	0	GET	^/api/users/[^/]+$	3.050	0.127	0.156	0.156	172
0.2%	23	0	0	GET	^/api/posts/[^/]+/comments$	2.869	0.125	0.153	0.153	738
0.1%	12	0	0	GET	^/api/posts/[^/]+/realmojis$	1.525	0.127	0.168	0.168	459
0.1%	9	0	0	POST	/api/auth/logout	1.503	0.167	0.302	0.302	0
0.1%	9	0	0	DELETE	^/api/comments/[^/]+$	1.417	0.157	0.233	0.233	0
0.1%	4	0	0	DELETE	^/api/posts/[^/]+$	1.175	0.294	0.326	0.326	0
0.1%	6	0	0	PUT	^/api/follows/[^/]+$	1.104	0.184	0.356	0.356	0
0.1%	6	0	0	GET	^/api/memories/[^/]+	0.969	0.162	0.287	0.287	786
0.1%	8	1	0	GET	/api/me	0.920	0.115	0.134	0.134	231
0.1%	6	2	0	PATCH	^/api/posts/[^/]+$	0.915	0.152	0.239	0.239	421
0.1%	4	0	0	PATCH	/api/me	0.665	0.166	0.202	0.202	282
0.0%	1	0	0	POST	/api/me/account-deletion	0.208	0.208	0.208	0.208	0
0.0%	1	0	0	DELETE	^/api/friendships/[^/]+$	0.202	0.202	0.202	0.202	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 15687 クエリ / 82 種類 / 合計 39.414s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	42.2%	57	16.643	291.98	801.12	7185.8	134.7	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	13.8%	837	5.433	6.49	22.72	625.8	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	12.9%	772	5.102	6.61	248.48	123.8	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	7.1%	1572	2.811	1.79	5.29	13.1	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	6.3%	1	2.482	2481.69	2481.69	29503.0	1199786.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	4.1%	1	1.607	1607.02	1607.02	20164.0	626110.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	2.8%	1	1.117	1116.62	1116.62	11528.0	529833.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	1.9%	561	0.758	1.35	3.77	28.4	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
9	1.2%	209	0.485	2.32	5.37	31.9	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.9%	2	0.335	167.73	179.90	217208.5	2555.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.7%	260	0.292	1.12	4.06	24.6	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.7%	2	0.263	131.32	254.17	16757.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
13	0.6%	1	0.241	240.76	240.76	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
14	0.5%	1	0.195	195.41	195.41	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
15	0.4%	1	0.176	176.31	176.31	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: 345855c6969b3e2bf523a5ed2de993a7bc39daef
Type: cpu
Time: 2026-10-31 11:48:28 JST
Duration: 60.04s, Total samples = 75.14s (125.16%)
Showing nodes accounting for 66.03s, 87.88% of 75.14s total
Dropped 832 nodes (cum <= 0.38s)
Showing top 40 nodes out of 170
      flat  flat%   sum%        cum   cum%
    11.74s 15.62% 15.62%     11.74s 15.62%  crypto/internal/fips140/sha256.blockAVX2
    10.51s 13.99% 29.61%     10.51s 13.99%  runtime.spanClass.sizeclass (inline)
     7.22s  9.61% 39.22%     11.65s 15.50%  image/jpeg.(*decoder).reconstructBlock
     3.62s  4.82% 44.04%     14.95s 19.90%  runtime.tryDeferToSpanScan
     3.58s  4.76% 48.80%     24.82s 33.03%  image/jpeg.(*decoder).processSOS
     2.76s  3.67% 52.48%      4.55s  6.06%  image/jpeg.(*decoder).decodeHuffman
     2.28s  3.03% 55.51%      2.35s  3.13%  image/jpeg.(*decoder).receiveExtend
     2.11s  2.81% 58.32%      2.11s  2.81%  runtime.memmove
     2.02s  2.69% 61.01%      2.02s  2.69%  runtime.memclrNoHeapPointers
     1.88s  2.50% 63.51%      2.37s  3.15%  image/jpeg.idctCols
     1.61s  2.14% 65.65%      6.98s  9.29%  runtime.scanObjectsSmall
     1.41s  1.88% 67.53%         2s  2.66%  image/jpeg.idctRows
     1.38s  1.84% 69.36%      1.82s  2.42%  image/jpeg.(*decoder).ensureNBits
     1.33s  1.77% 71.13%      5.48s  7.29%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.30s  1.73% 72.86%      1.30s  1.73%  internal/runtime/syscall/linux.Syscall6
     1.24s  1.65% 74.51%     11.69s 15.56%  runtime.scanObject
     1.22s  1.62% 76.14%      1.22s  1.62%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.08s  1.44% 77.58%      1.08s  1.44%  image/jpeg.dctBox (inline)
     0.87s  1.16% 78.73%      0.87s  1.16%  memeqbody
     0.84s  1.12% 79.85%      0.84s  1.12%  aeshashbody
     0.56s  0.75% 80.60%     13.04s 17.35%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.55s  0.73% 81.33%      8.24s 10.97%  runtime.scanSpan
     0.51s  0.68% 82.01%     12.86s 17.11%  crypto/internal/fips140/sha256.(*Digest).Write
     0.44s  0.59% 82.59%     13.66s 18.18%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.44s  0.59% 83.18%      0.44s  0.59%  runtime.extractHeapBitsSmall
     0.39s  0.52% 83.70%      0.44s  0.59%  image/jpeg.(*decoder).readByteStuffedByte
     0.37s  0.49% 84.19%     14.62s 19.46%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.36s  0.48% 84.67%      0.60s   0.8%  example.com/isureal-company/isureal.postLess
     0.35s  0.47% 85.13%      0.47s  0.63%  runtime.(*spanInlineMarkBits).tryAcquire
     0.30s   0.4% 85.53%      0.38s  0.51%  runtime.findObject
     0.28s  0.37% 85.91%      2.23s  2.97%  runtime.mapaccess1_faststr
     0.20s  0.27% 86.17%      0.79s  1.05%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.20s  0.27% 86.44%     11.94s 15.89%  crypto/internal/fips140/sha256.block
     0.20s  0.27% 86.70%      0.96s  1.28%  runtime.mallocgcSmallScanNoHeader
     0.19s  0.25% 86.96%     10.13s 13.48%  runtime.mallocgc
     0.18s  0.24% 87.20%     15.77s 20.99%  example.com/isureal-company/isureal.imageSignature
     0.17s  0.23% 87.42%      0.63s  0.84%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.14s  0.19% 87.61%      0.55s  0.73%  runtime.scanObjectSmall
     0.10s  0.13% 87.74%      0.62s  0.83%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.10s  0.13% 87.88%      2.63s  3.50%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 75.3% / 100% | 62% | 1.5% | 8 | 175 MB | CPU が飽和 |
| isu2 (2 cpu) | 1.2% / 5% | 0% | 0.3% | 2 | 416 MB |  |
| isu3 (2 cpu) | 1.3% / 4% | 0% | 0.2% | 3 | 220 MB |  |
| isu4 (2 cpu) | 23.4% / 69% | 0% | 13.4% | 3 | 107 MB |  |

**isu1（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 127% / nginx 14% |
| isu2 | 目立つプロセス無し |
| isu3 | isureal 1% |
| isu4 | postgres 18% |

## Diff from previous run

スコア: 17072 → 26395 (+9323)

`6a7a33a..8e518c1`

```
 webapp/go/limiter.go  | 37 +++++++++++++++++++++++++++++++++++++
 webapp/go/server.go   | 10 +++++-----
 webapp/go/user_api.go | 17 ++++++++++++++---
 3 files changed, 56 insertions(+), 8 deletions(-)
```
