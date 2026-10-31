# run 20261031-124102-04641dc

## Summary

| | |
|---|---|
| score | 64161 |
| status | pass |
| ref | origin/main |
| sha | 04641dc401ec |
| 所要時間 | 230s |
| 前回比 | +58443 (1022.1%) vs 20261031-123602-3d5209e |
| ベスト比 | -2556 (-3.8%) vs 20261031-122401-fdbe826 = 66717 |
| note | 低い版の試走の直後の main |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	52	154	428	634	634
europe	7	17	45	99	99	99
america	6	14	34	61	61	61
```

```
{"id":"9332","status":"done","portalStatus":"finished","passed":true,"score":64161,"raw":64161,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:43:23.823Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:43:26.785Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:43:29.782Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:43:32.807Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T03:43:35.783Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T03:43:36.203Z level=INFO msg="user moved" user=mumaku5345 from=asia to=europe
time=2026-10-31T03:43:38.786Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T03:43:41.818Z level=INFO msg="asia 3 日目: 154 人"
time=2026-10-31T03:43:44.834Z level=INFO msg="europe 3 日目: 45 人"
time=2026-10-31T03:43:47.789Z level=INFO msg="america 3 日目: 34 人"
time=2026-10-31T03:43:50.819Z level=INFO msg="asia 4 日目: 428 人"
time=2026-10-31T03:43:53.903Z level=INFO msg="europe 4 日目: 99 人"
time=2026-10-31T03:43:56.884Z level=INFO msg="america 4 日目: 61 人"
time=2026-10-31T03:44:00.041Z level=INFO msg="asia 5 日目: 634 人"
time=2026-10-31T03:44:03.698Z level=INFO msg="europe 5 日目: 99 人"
time=2026-10-31T03:44:05.974Z level=INFO msg="america 5 日目: 61 人"
time=2026-10-31T03:44:14.972Z level=INFO msg="asia 最終日（グランプリ）: 634 人"
time=2026-10-31T03:44:15.045Z level=INFO msg="europe 最終日（グランプリ）: 99 人"
time=2026-10-31T03:44:15.081Z level=INFO msg="america 最終日（グランプリ）: 61 人"
time=2026-10-31T03:44:23.805Z level=INFO msg="グランプリ: アジア 15 人, ヨーロッパ 1 人, アメリカ 0 人"
{
  "pass": true,
  "score": 64161,
  "messages": [
    "taku_5101: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "danrute_7530: ログアウトできません: POST /api/auth/logout: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
29.6%	1147	0	0	GET	/api/notifications	9903.500	8.634	14.942	15.212	173
19.0%	1488	76	0	POST	/api/posts	6369.875	4.281	13.731	15.093	619
12.8%	3173	52	0	POST	/api/friend-requests	4288.711	1.352	13.711	17.220	114
10.4%	8886	958	0	GET	^/api/images/[^/]+	3468.294	0.390	10.011	15.306	126494
7.8%	3761	1065	0	POST	/api/auth/login	2593.865	0.690	10.151	14.487	182
4.8%	1057	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1617.684	1.530	14.171	17.044	116
3.7%	651	1	0	POST	/api/auth/register	1225.825	1.883	9.439	13.414	227
1.7%	483	122	0	PUT	/api/me/profile-image	577.747	1.196	10.176	15.420	6
1.4%	422	0	0	PUT	^/api/follows/[^/]+$	473.308	1.122	14.146	16.187	0
1.3%	312	17	0	POST	^/api/posts/[^/]+/realmojis$	431.419	1.383	10.136	16.417	454
1.3%	2068	0	0	GET	/api/posting-window/current	418.322	0.202	2.717	12.081	237
1.0%	290	4	0	POST	^/api/posts/[^/]+/comments$	322.543	1.112	12.544	16.633	340
0.8%	2281	0	0	GET	/api/memories	282.756	0.124	0.960	6.259	10466
0.8%	317	0	0	PATCH	/api/me	275.944	0.870	4.691	13.850	281
0.7%	416	0	0	POST	/api/auth/logout	232.940	0.560	12.070	19.770	0
0.6%	6340	0	0	GET	^/api/feeds/[^/]+	187.787	0.030	0.156	3.438	1814
0.5%	158	0	0	GET	/api/me/export	151.944	0.962	8.047	11.728	2079835
0.4%	127	9	0	PATCH	^/api/posts/[^/]+$	121.998	0.961	5.114	6.706	620
0.3%	54	0	0	DELETE	^/api/posts/[^/]+$	99.274	1.838	12.865	12.865	0
0.2%	46	0	0	DELETE	^/api/comments/[^/]+$	66.806	1.452	11.485	11.485	0
0.2%	1486	0	0	GET	/api/users	63.344	0.043	0.406	3.576	244
0.2%	1084	82	0	GET	^/api/posts/me/current$	59.003	0.054	0.802	3.438	604
0.1%	38	0	0	DELETE	^/api/realmojis/[^/]+$	46.807	1.232	11.070	11.070	0
0.1%	428	0	0	GET	^/api/memories/[^/]+	26.429	0.062	0.548	1.518	607
0.1%	969	0	0	GET	^/api/relationships/[^/]+	25.333	0.026	0.158	3.939	1046
0.1%	11	0	0	PUT	^/api/admin/users/[^/]+/certification$	21.450	1.950	13.302	13.302	0
0.1%	1	0	0	POST	/api/initialize	18.428	18.428	18.428	18.428	90
0.1%	889	11	0	GET	^/api/posts/[^/]+$	18.081	0.020	0.110	0.188	615
0.0%	488	0	0	GET	^/api/users/[^/]+$	10.792	0.022	0.122	0.172	193
0.0%	365	2	0	GET	/api/me	10.342	0.028	0.137	1.070	301
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.3%	232	1	0	GET	/api/notifications	1886.102	8.130	11.497	11.499	177
19.6%	269	9	0	POST	/api/posts	1106.334	4.113	14.920	17.864	628
13.6%	516	9	0	POST	/api/friend-requests	771.931	1.496	12.421	14.202	113
7.5%	369	0	0	POST	/api/auth/login	426.957	1.157	12.730	15.292	247
3.6%	1273	8	0	GET	^/api/images/[^/]+	202.281	0.159	0.964	1.434	177368
3.4%	156	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	190.022	1.218	14.113	15.694	115
3.1%	1079	0	0	GET	^/api/feeds/[^/]+	174.673	0.162	0.544	1.295	2088
2.1%	374	0	0	GET	/api/posting-window/current	121.091	0.324	2.793	3.524	239
1.5%	79	1	0	POST	^/api/posts/[^/]+/comments$	86.588	1.096	12.833	12.833	346
1.4%	358	0	0	GET	/api/memories	81.318	0.227	1.509	2.946	12046
1.3%	69	0	0	PUT	^/api/follows/[^/]+$	72.439	1.050	9.478	9.478	0
1.0%	62	0	0	POST	/api/auth/logout	58.552	0.944	13.726	13.726	0
0.9%	96	0	0	POST	/api/auth/register	53.020	0.552	3.115	3.115	228
0.9%	54	0	0	PUT	/api/me/profile-image	51.607	0.956	7.749	7.749	0
0.9%	47	1	0	POST	^/api/posts/[^/]+/realmojis$	51.545	1.097	6.960	6.960	452
0.9%	48	0	0	PATCH	/api/me	50.294	1.048	7.935	7.935	271
0.7%	237	0	0	GET	/api/users	39.113	0.165	1.039	1.539	250
0.6%	203	17	0	GET	^/api/posts/me/current$	31.120	0.153	0.484	0.683	598
0.5%	17	1	0	PATCH	^/api/posts/[^/]+$	27.247	1.603	14.846	14.846	602
0.4%	14	0	0	GET	/api/me/export	24.365	1.740	3.878	3.878	2578302
0.4%	9	0	0	DELETE	^/api/posts/[^/]+$	23.410	2.601	11.021	11.021	0
0.4%	156	0	0	GET	^/api/relationships/[^/]+	22.996	0.147	0.511	0.842	1027
0.4%	140	5	0	GET	^/api/posts/[^/]+$	21.255	0.152	0.474	0.492	594
0.3%	12	0	0	DELETE	^/api/realmojis/[^/]+$	19.211	1.601	13.878	13.878	0
0.2%	58	0	0	GET	^/api/memories/[^/]+	12.113	0.209	1.147	1.147	812
0.2%	59	1	0	GET	^/api/posts/[^/]+/comments$	10.307	0.175	2.158	2.158	837
0.2%	63	0	0	GET	^/api/users/[^/]+$	9.707	0.154	0.405	0.405	190
0.2%	17	0	0	DELETE	^/api/comments/[^/]+$	9.268	0.545	5.980	5.980	0
0.2%	59	0	0	GET	^/api/posts/[^/]+/realmojis$	8.934	0.151	0.295	0.295	614
0.2%	59	2	0	GET	/api/me	8.899	0.151	0.408	0.408	291
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.4%	148	0	0	GET	/api/notifications	1145.683	7.741	9.130	9.348	178
20.5%	164	5	0	POST	/api/posts	726.824	4.432	13.693	14.368	621
11.2%	285	3	0	POST	/api/friend-requests	394.845	1.385	12.785	14.697	114
8.4%	226	0	0	POST	/api/auth/login	295.638	1.308	12.316	14.638	249
5.3%	132	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	186.751	1.415	12.178	12.225	116
3.9%	793	9	0	GET	^/api/images/[^/]+	136.442	0.172	0.983	1.196	154779
2.2%	512	0	0	GET	^/api/feeds/[^/]+	77.952	0.152	0.474	1.126	2158
1.7%	228	0	0	GET	/api/posting-window/current	60.953	0.267	2.574	3.095	239
1.7%	51	0	0	PUT	^/api/follows/[^/]+$	59.858	1.174	5.318	5.318	0
1.6%	241	0	0	GET	/api/memories	58.081	0.241	0.910	2.946	13965
1.6%	60	0	0	POST	/api/auth/register	58.035	0.967	7.944	7.944	229
1.3%	30	0	0	PUT	/api/me/profile-image	46.987	1.566	6.240	6.240	0
1.1%	9	0	0	DELETE	^/api/posts/[^/]+$	38.105	4.234	16.177	16.177	0
0.9%	37	0	0	POST	/api/auth/logout	32.068	0.867	10.963	10.963	0
0.9%	27	0	0	PATCH	/api/me	30.616	1.134	11.259	11.259	276
0.6%	139	0	0	GET	/api/users	22.727	0.164	0.562	1.749	248
0.6%	14	2	0	POST	^/api/posts/[^/]+/realmojis$	22.288	1.592	4.838	4.838	356
0.6%	123	16	0	GET	^/api/posts/me/current$	20.463	0.166	0.836	1.219	557
0.5%	5	0	0	PUT	^/api/blocks/[^/]+$	18.866	3.773	8.607	8.607	0
0.5%	8	0	0	GET	/api/me/export	17.357	2.170	4.859	4.859	3244369
0.4%	94	0	0	GET	^/api/relationships/[^/]+	15.466	0.165	1.642	1.642	1419
0.4%	99	9	0	GET	^/api/posts/[^/]+$	13.758	0.139	0.381	0.381	564
0.4%	8	2	0	POST	^/api/posts/[^/]+/comments$	12.461	1.558	2.920	2.920	242
0.3%	11	3	0	PATCH	^/api/posts/[^/]+$	12.083	1.098	6.924	6.924	458
0.3%	2	0	0	DELETE	^/api/comments/[^/]+$	9.238	4.619	8.811	8.811	0
0.2%	54	0	0	GET	^/api/users/[^/]+$	8.327	0.154	0.727	0.727	191
0.2%	32	0	0	GET	^/api/memories/[^/]+	5.579	0.174	0.329	0.329	792
0.1%	33	0	0	GET	/api/me	5.229	0.158	0.449	0.449	297
0.1%	29	0	0	GET	^/api/posts/[^/]+/realmojis$	3.540	0.122	0.207	0.207	733
0.1%	22	0	0	GET	^/api/posts/[^/]+/comments$	2.946	0.134	0.239	0.239	1176
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 45404 クエリ / 81 種類 / 合計 275.197s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	58.6%	2632	161.265	61.27	600.55	638.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	23.8%	180	65.611	364.50	2230.17	7118.9	139.3	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.1%	1776	13.985	7.87	310.78	116.5	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.4%	5531	12.208	2.21	59.80	10.6	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	2.0%	2	5.631	2815.57	5368.49	236089.5	2780.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
6	1.6%	1834	4.487	2.45	17.82	27.0	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
7	0.8%	1	2.166	2166.11	2166.11	30049.0	1199935.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.7%	1	1.828	1828.42	1828.42	21660.0	627702.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.5%	353	1.294	3.67	14.94	30.6	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.3%	3889	0.858	0.22	29.70	20.4	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.3%	2	0.812	406.18	804.26	16794.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.3%	1	0.762	761.89	761.89	11744.0	530030.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
13	0.3%	370	0.730	1.97	10.48	22.8	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
14	0.2%	1	0.519	518.63	518.63	26843.0	2591.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.2%	3291	0.425	0.13	10.54	11.4	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 47a7f57247a2e8d633423637a42a8f5663afabb9
Type: cpu
Time: 2026-10-31 12:43:00 JST
Duration: 60.08s, Total samples = 43.66s (72.67%)
Showing nodes accounting for 33.26s, 76.18% of 43.66s total
Dropped 945 nodes (cum <= 0.22s)
Showing top 40 nodes out of 280
      flat  flat%   sum%        cum   cum%
     3.02s  6.92%  6.92%      3.02s  6.92%  internal/runtime/syscall/linux.Syscall6
     2.66s  6.09% 13.01%      2.66s  6.09%  crypto/internal/fips140/sha256.blockAVX2
     2.50s  5.73% 18.74%      7.78s 17.82%  runtime.scanObject
     2.40s  5.50% 24.23%      2.40s  5.50%  runtime.memclrNoHeapPointers
     2.07s  4.74% 28.97%      2.07s  4.74%  runtime.memmove
     1.84s  4.21% 33.19%      6.21s 14.22%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.84s  4.21% 37.40%      2.36s  5.41%  runtime.findObject
     1.78s  4.08% 41.48%      2.93s  6.71%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.74s  3.99% 45.46%      5.58s 12.78%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.66s  3.80% 49.27%      2.62s  6.00%  runtime.tryDeferToSpanScan
     1.36s  3.11% 52.38%      1.43s  3.28%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     1.21s  2.77% 55.15%      1.21s  2.77%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.05s  2.40% 57.56%      1.05s  2.40%  aeshashbody
     0.91s  2.08% 59.64%      1.20s  2.75%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.64s  1.47% 61.11%      0.64s  1.47%  example.com/isureal-company/isureal.(*UserRow).active
     0.64s  1.47% 62.57%      0.64s  1.47%  runtime.spanClass.sizeclass (inline)
     0.52s  1.19% 63.77%      0.52s  1.19%  memeqbody
     0.49s  1.12% 64.89%      1.69s  3.87%  runtime.mapaccess1_faststr
     0.37s  0.85% 65.74%      0.37s  0.85%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.37s  0.85% 66.58%      0.55s  1.26%  example.com/isureal-company/isureal.postLess
     0.36s  0.82% 67.41%      0.36s  0.82%  indexbody
     0.31s  0.71% 68.12%      0.31s  0.71%  runtime.nextFreeFast (inline)
     0.28s  0.64% 68.76%      0.64s  1.47%  runtime.greyobject
     0.28s  0.64% 69.40%      1.03s  2.36%  runtime.mallocgcSmallScanNoHeader
     0.27s  0.62% 70.02%      0.64s  1.47%  database/sql.convertAssignRows
     0.27s  0.62% 70.64%      0.28s  0.64%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.25s  0.57% 71.21%      0.28s  0.64%  runtime.spanOf (inline)
     0.22s   0.5% 71.71%      0.22s   0.5%  runtime.madvise
     0.22s   0.5% 72.22%      0.22s   0.5%  runtime.typePointers.nextFast (inline)
     0.21s  0.48% 72.70%      0.28s  0.64%  runtime.(*mspan).markBitsForIndex (inline)
     0.21s  0.48% 73.18%      1.75s  4.01%  runtime.mapassign_faststr
     0.20s  0.46% 73.64%      0.70s  1.60%  runtime.scanObjectsSmall
     0.18s  0.41% 74.05%      3.03s  6.94%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.17s  0.39% 74.44%      0.47s  1.08%  runtime.mapaccess2_faststr
     0.15s  0.34% 74.78%      0.26s   0.6%  runtime.heapArenaOf (inline)
     0.13s   0.3% 75.08%      3.35s  7.67%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.13s   0.3% 75.38%      1.10s  2.52%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.27% 75.65%      2.91s  6.67%  crypto/internal/fips140/sha256.(*Digest).Write
     0.12s  0.27% 75.93%      3.13s  7.17%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s  0.25% 76.18%      0.24s  0.55%  internal/runtime/maps.(*Iter).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 50.4% / 100% | 38% | 0.1% | 7 | 172 MB | CPU が飽和 |
| isu2 (2 cpu) | 10.2% / 34% | 0% | 0.4% | 3 | 2614 MB |  |
| isu3 (2 cpu) | 7.1% / 34% | 0% | 0.4% | 3 | 2725 MB |  |
| isu4 (2 cpu) | 44.4% / 100% | 15% | 16.0% | 37 | 106 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 82% / nginx 12% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 18% |
| isu3 | isureal 13% |
| isu4 | postgres 45% / isureal 9% |

## Diff from previous run

スコア: 5718 → 64161 (+58443)

`3d5209e..04641dc`

```
 webapp/go/account_deletion_api.go |   2 +-
 webapp/go/admin_user_api.go       |   2 +-
 webapp/go/comment_api.go          |   2 +-
 webapp/go/feed_api.go             | 113 +++++++++----
 webapp/go/feed_test.go            | 331 ++++++++++++++++++++++++++++++++++++++
 webapp/go/realmoji_api.go         |   2 +-
 webapp/go/relationship.go         |  10 +-
 webapp/go/server.go               |   1 -
 webapp/go/slow.go                 |  21 ---
 webapp/go/store.go                | 113 +++++++++++--
 10 files changed, 526 insertions(+), 71 deletions(-)
```
