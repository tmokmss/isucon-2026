# run 20261031-132303-5a138c8

## Summary

| | |
|---|---|
| score | 83744 |
| status | pass |
| ref | nginx-buffers-tcp |
| sha | 5a138c83d36b |
| 所要時間 | 804s |
| 前回比 | +19044 (29.4%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +19044 (29.4%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | nginx-no-tempfile + tcp_slow_start_after_idle=0 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	140	396	667	668
europe	8	18	48	122	120	120
america	6	16	43	82	84	84
```

```
{"id":"13209","status":"done","portalStatus":"finished","passed":true,"score":83744,"raw":83744,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:35:03.514Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:35:05.686Z level=INFO msg="user moved" user=solsa5029 from=asia to=europe
time=2026-10-31T04:35:06.491Z level=INFO msg="europe 1 日目: 8 人"
time=2026-10-31T04:35:09.484Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:35:12.511Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T04:35:15.486Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T04:35:18.484Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T04:35:21.521Z level=INFO msg="asia 3 日目: 140 人"
time=2026-10-31T04:35:23.818Z level=INFO msg="user moved" user=hima_5754 from=asia to=america
time=2026-10-31T04:35:24.566Z level=INFO msg="europe 3 日目: 48 人"
time=2026-10-31T04:35:25.615Z level=INFO msg="user moved" user=hima_5754 from=america to=europe
time=2026-10-31T04:35:25.726Z level=INFO msg="user moved" user=yoka.5983 from=asia to=america
time=2026-10-31T04:35:25.863Z level=INFO msg="user moved" user=yaormo.5935 from=asia to=europe
time=2026-10-31T04:35:27.500Z level=INFO msg="america 3 日目: 43 人"
time=2026-10-31T04:35:30.535Z level=INFO msg="asia 4 日目: 396 人"
time=2026-10-31T04:35:34.346Z level=INFO msg="user moved" user=suri.6415 from=america to=europe
time=2026-10-31T04:35:34.899Z level=INFO msg="europe 4 日目: 122 人"
time=2026-10-31T04:35:35.647Z level=INFO msg="user moved" user=solsa5029 from=europe to=america
time=2026-10-31T04:35:36.723Z level=INFO msg="america 4 日目: 82 人"
time=2026-10-31T04:35:38.206Z level=INFO msg="user moved" user=kimunoa_7107 from=asia to=america
time=2026-10-31T04:35:38.702Z level=INFO msg="user moved" user=muso7848 from=asia to=america
time=2026-10-31T04:35:39.955Z level=INFO msg="asia 5 日目: 667 人"
time=2026-10-31T04:35:40.262Z level=INFO msg="user moved" user=hima_5754 from=europe to=asia
time=2026-10-31T04:35:42.550Z level=INFO msg="europe 5 日目: 120 人"
time=2026-10-31T04:35:45.711Z level=INFO msg="america 5 日目: 84 人"
time=2026-10-31T04:35:54.602Z level=INFO msg="america 最終日（グランプリ）: 84 人"
time=2026-10-31T04:35:54.690Z level=INFO msg="asia 最終日（グランプリ）: 668 人"
time=2026-10-31T04:35:54.859Z level=INFO msg="europe 最終日（グランプリ）: 120 人"
time=2026-10-31T04:36:03.512Z level=INFO msg="グランプリ: アジア 191 人, ヨーロッパ 20 人, アメリカ 15 人"
{
  "pass": true,
  "score": 83744,
  "messages": [
    "konimu_5023: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
31.8%	1239	1	0	GET	/api/notifications	10205.062	8.237	14.713	15.118	169
20.2%	1711	48	0	POST	/api/posts	6463.341	3.778	13.447	14.757	632
10.9%	10585	646	0	GET	^/api/images/[^/]+	3489.655	0.330	8.837	14.471	125121
9.2%	3577	61	0	POST	/api/friend-requests	2953.300	0.826	10.966	15.069	114
6.6%	4448	1223	0	POST	/api/auth/login	2110.836	0.475	5.787	14.487	184
4.6%	1401	8	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1484.873	1.060	11.035	13.780	115
2.2%	508	130	0	PUT	/api/me/profile-image	712.357	1.402	7.926	9.048	5
2.2%	689	1	0	POST	/api/auth/register	696.094	1.010	7.024	9.156	227
1.7%	7543	0	0	GET	^/api/feeds/[^/]+	550.576	0.073	0.631	2.220	2263
1.3%	2662	0	0	GET	/api/memories	419.869	0.158	0.984	9.742	10532
1.2%	357	4	0	POST	^/api/posts/[^/]+/realmojis$	380.656	1.066	9.375	12.054	473
1.0%	397	5	0	POST	^/api/posts/[^/]+/comments$	332.028	0.836	8.854	12.627	340
1.0%	457	0	0	PUT	^/api/follows/[^/]+$	331.608	0.726	5.256	6.575	0
0.9%	154	10	0	PATCH	^/api/posts/[^/]+$	279.322	1.814	12.042	12.802	612
0.9%	331	0	0	PATCH	/api/me	277.995	0.840	5.269	6.550	274
0.8%	2452	0	0	GET	/api/posting-window/current	266.984	0.109	0.575	2.516	236
0.5%	191	0	0	GET	/api/me/export	172.996	0.906	10.435	14.833	1832694
0.4%	547	0	0	POST	/api/auth/logout	143.397	0.262	3.468	10.991	0
0.4%	1288	72	0	GET	^/api/posts/me/current$	139.842	0.109	0.707	1.920	613
0.4%	1679	0	0	GET	/api/users	129.228	0.077	0.433	1.351	236
0.4%	62	0	0	DELETE	^/api/posts/[^/]+$	118.244	1.907	12.577	12.577	0
0.2%	1064	0	0	GET	^/api/relationships/[^/]+	69.818	0.066	0.650	1.848	1136
0.2%	1190	33	0	GET	^/api/posts/[^/]+$	67.920	0.057	0.450	2.150	601
0.1%	514	0	0	GET	^/api/memories/[^/]+	44.226	0.086	0.582	2.945	637
0.1%	387	3	0	GET	/api/me	34.825	0.090	0.433	0.450	300
0.1%	58	0	0	DELETE	^/api/realmojis/[^/]+$	29.147	0.503	4.731	4.731	0
0.1%	535	0	0	GET	^/api/users/[^/]+$	27.296	0.051	0.433	0.625	193
0.1%	77	0	0	DELETE	^/api/comments/[^/]+$	27.275	0.354	2.629	2.629	0
0.1%	497	0	0	GET	^/api/posts/[^/]+/realmojis$	24.796	0.050	0.651	1.325	1042
0.1%	1	0	0	POST	/api/initialize	21.089	21.089	21.089	21.089	90
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.1%	289	0	0	GET	/api/notifications	2362.810	8.176	11.975	12.276	178
23.9%	357	5	0	POST	/api/posts	1611.627	4.514	14.753	14.872	636
8.0%	682	10	0	POST	/api/friend-requests	535.856	0.786	5.940	12.970	114
6.0%	480	0	0	POST	/api/auth/login	404.058	0.842	9.855	13.070	248
4.1%	1554	14	0	GET	^/api/images/[^/]+	274.136	0.176	0.924	2.609	168409
4.0%	223	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	272.080	1.220	11.564	13.585	116
3.6%	1280	0	0	GET	^/api/feeds/[^/]+	245.370	0.192	0.670	1.657	2149
1.9%	514	0	0	GET	/api/posting-window/current	124.795	0.243	0.702	2.212	239
1.8%	446	0	0	GET	/api/memories	122.906	0.276	2.119	5.400	11483
1.6%	88	4	0	POST	^/api/posts/[^/]+/realmojis$	106.069	1.205	10.865	10.865	461
1.2%	85	0	0	POST	/api/auth/logout	77.797	0.915	10.301	10.301	0
1.1%	32	0	0	GET	/api/me/export	76.118	2.379	11.999	11.999	2410530
1.0%	116	0	0	POST	/api/auth/register	70.478	0.608	2.378	3.347	228
0.9%	319	0	0	GET	/api/users	63.625	0.199	0.587	0.614	236
0.9%	266	4	0	GET	^/api/posts/me/current$	59.734	0.225	1.237	1.728	631
0.8%	68	0	0	PUT	^/api/follows/[^/]+$	57.235	0.842	10.477	10.477	0
0.5%	197	0	0	GET	^/api/relationships/[^/]+	36.554	0.186	0.845	0.866	1056
0.5%	61	0	0	PUT	/api/me/profile-image	34.060	0.558	1.498	1.498	0
0.4%	16	0	0	DELETE	^/api/posts/[^/]+$	28.804	1.800	8.657	8.657	0
0.4%	65	0	0	PATCH	/api/me	26.787	0.412	2.367	2.367	275
0.4%	138	7	0	GET	^/api/posts/[^/]+$	26.238	0.190	0.966	1.236	586
0.4%	46	1	0	POST	^/api/posts/[^/]+/comments$	25.715	0.559	6.698	6.698	284
0.2%	71	0	0	GET	^/api/memories/[^/]+	15.881	0.224	0.722	0.722	654
0.2%	85	0	0	GET	^/api/users/[^/]+$	13.734	0.162	0.593	0.593	180
0.2%	21	1	0	PATCH	^/api/posts/[^/]+$	12.023	0.573	2.085	2.085	618
0.2%	16	0	0	DELETE	^/api/realmojis/[^/]+$	11.961	0.748	3.350	3.350	0
0.2%	67	1	0	GET	/api/me	11.551	0.172	0.383	0.383	293
0.2%	59	1	0	GET	^/api/posts/[^/]+/realmojis$	11.276	0.191	0.979	0.979	1352
0.1%	57	0	0	GET	^/api/posts/[^/]+/comments$	8.774	0.154	0.634	0.634	1743
0.1%	12	0	0	DELETE	^/api/comments/[^/]+$	3.816	0.318	1.093	1.093	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.0%	218	0	0	GET	/api/notifications	1642.726	7.535	9.306	9.337	179
22.8%	270	7	0	POST	/api/posts	1068.470	3.957	14.120	14.984	638
9.6%	476	6	0	POST	/api/friend-requests	450.304	0.946	8.891	11.427	114
5.3%	318	0	0	POST	/api/auth/login	248.110	0.780	10.848	12.286	249
5.0%	1083	8	0	GET	^/api/images/[^/]+	236.332	0.218	1.204	3.179	180070
4.6%	186	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	213.637	1.149	11.532	12.224	116
3.3%	730	0	0	GET	^/api/feeds/[^/]+	153.128	0.210	0.728	7.556	2325
2.0%	342	0	0	GET	/api/memories	95.688	0.280	2.047	4.859	14042
1.7%	361	0	0	GET	/api/posting-window/current	77.367	0.214	0.752	1.682	240
1.6%	27	0	0	GET	/api/me/export	72.967	2.702	12.105	12.105	2970614
1.0%	236	0	0	GET	/api/users	47.780	0.202	0.697	1.404	241
1.0%	209	14	0	GET	^/api/posts/me/current$	45.920	0.220	0.819	1.400	614
0.9%	38	2	0	POST	^/api/posts/[^/]+/comments$	41.879	1.102	6.516	6.516	352
0.8%	24	1	0	PATCH	^/api/posts/[^/]+$	39.481	1.645	9.334	9.334	633
0.7%	48	0	0	PUT	/api/me/profile-image	32.631	0.680	4.585	4.585	0
0.7%	54	0	0	POST	/api/auth/logout	31.545	0.584	3.977	3.977	0
0.6%	80	0	0	POST	/api/auth/register	30.309	0.379	2.496	2.496	229
0.6%	25	3	0	POST	^/api/posts/[^/]+/realmojis$	28.646	1.146	3.479	3.479	465
0.6%	135	0	0	GET	^/api/relationships/[^/]+	26.136	0.194	0.666	0.830	1312
0.5%	52	0	0	PUT	^/api/follows/[^/]+$	25.766	0.495	4.967	4.967	0
0.4%	37	0	0	PATCH	/api/me	16.785	0.454	3.425	3.425	290
0.3%	72	6	0	GET	^/api/posts/[^/]+$	12.627	0.175	0.343	0.343	575
0.2%	45	0	0	GET	^/api/memories/[^/]+	10.818	0.240	1.850	1.850	995
0.2%	8	0	0	DELETE	^/api/posts/[^/]+$	8.891	1.111	2.107	2.107	0
0.2%	52	0	0	GET	/api/me	8.476	0.163	0.671	0.671	297
0.1%	36	0	0	GET	^/api/users/[^/]+$	6.779	0.188	0.720	0.720	195
0.1%	28	0	0	GET	^/api/posts/[^/]+/realmojis$	4.515	0.161	0.326	0.326	1099
0.1%	31	0	0	GET	^/api/posts/[^/]+/comments$	4.308	0.139	0.335	0.335	1368
0.1%	6	0	0	DELETE	^/api/comments/[^/]+$	2.792	0.465	1.718	1.718	0
0.0%	3	0	0	DELETE	^/api/realmojis/[^/]+$	2.071	0.690	1.104	1.104	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 55717 クエリ / 81 種類 / 合計 403.909s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	62.6%	3226	252.922	78.40	625.66	661.3	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	20.3%	250	82.039	328.16	1315.17	1269.5	129.4	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	6.2%	2232	25.001	11.20	1156.01	128.2	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.0%	7075	16.237	2.30	45.83	10.9	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.7%	2280	6.710	2.94	120.89	25.9	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.6%	2	6.344	3171.87	6184.90	236598.5	2786.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.6%	1	2.347	2346.59	2346.59	30703.0	1199899.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.4%	1	1.766	1766.05	1766.05	23420.0	627716.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.4%	459	1.554	3.39	22.47	29.9	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.3%	2	1.332	666.16	1324.47	16763.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.3%	1	1.039	1038.75	1038.75	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.2%	4632	0.942	0.20	38.68	19.9	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
13	0.2%	473	0.801	1.69	36.04	21.8	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
14	0.2%	1	0.778	778.11	778.11	12389.0	530029.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
15	0.2%	1	0.729	729.05	729.05	28631.0	2605.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 0ab73a4d871f9265109b043c9bf654ebba782ac5
Type: cpu
Time: 2026-10-31 13:34:36 JST
Duration: 60.04s, Total samples = 32.71s (54.48%)
Showing nodes accounting for 23.69s, 72.42% of 32.71s total
Dropped 975 nodes (cum <= 0.16s)
Showing top 40 nodes out of 296
      flat  flat%   sum%        cum   cum%
     2.96s  9.05%  9.05%      2.96s  9.05%  internal/runtime/syscall/linux.Syscall6
     1.76s  5.38% 14.43%      5.90s 18.04%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.70s  5.20% 19.63%      1.70s  5.20%  crypto/internal/fips140/sha256.blockAVX2
     1.64s  5.01% 24.64%      1.64s  5.01%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.63s  4.98% 29.62%      1.63s  4.98%  runtime.memclrNoHeapPointers
     1.48s  4.52% 34.15%      1.48s  4.52%  runtime.memmove
     1.44s  4.40% 38.55%      5.49s 16.78%  runtime.scanObject
     1.41s  4.31% 42.86%      1.77s  5.41%  runtime.findObject
     1.11s  3.39% 46.25%      1.98s  6.05%  runtime.tryDeferToSpanScan
     1.08s  3.30% 49.56%      1.08s  3.30%  aeshashbody
     0.66s  2.02% 51.57%      0.66s  2.02%  runtime.spanClass.sizeclass (inline)
     0.61s  1.86% 53.44%      0.61s  1.86%  example.com/isureal-company/isureal.(*UserRow).active
     0.54s  1.65% 55.09%      0.54s  1.65%  memeqbody
     0.53s  1.62% 56.71%      1.93s  5.90%  runtime.mapaccess1_faststr
     0.41s  1.25% 57.96%      0.60s  1.83%  example.com/isureal-company/isureal.postLess
     0.33s  1.01% 58.97%      0.33s  1.01%  runtime.nextFreeFast (inline)
     0.29s  0.89% 59.86%      0.29s  0.89%  indexbody
     0.27s  0.83% 60.68%      0.27s  0.83%  runtime.typePointers.nextFast (inline)
     0.26s  0.79% 61.48%      0.39s  1.19%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.25s  0.76% 62.24%      3.80s 11.62%  runtime.mallocgc
     0.22s  0.67% 62.92%      1.06s  3.24%  runtime.mallocgcSmallScanNoHeader
     0.20s  0.61% 63.53%      1.86s  5.69%  runtime.mapassign_faststr
     0.19s  0.58% 64.11%      0.19s  0.58%  runtime.futex
     0.19s  0.58% 64.69%      0.36s  1.10%  runtime.greyobject
     0.19s  0.58% 65.27%      0.44s  1.35%  runtime.scanObjectsSmall
     0.19s  0.58% 65.85%      0.19s  0.58%  sync/atomic.(*Int32).Add
     0.18s  0.55% 66.40%      0.47s  1.44%  database/sql.convertAssignRows
     0.18s  0.55% 66.95%      0.19s  0.58%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.18s  0.55% 67.50%      0.19s  0.58%  runtime.(*mspan).writeHeapBitsSmall
     0.18s  0.55% 68.05%      0.24s  0.73%  runtime.typePointers.next
     0.17s  0.52% 68.57%      0.19s  0.58%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.17s  0.52% 69.09%      0.17s  0.52%  runtime.(*mspan).divideByElemSize (inline)
     0.16s  0.49% 69.58%      0.74s  2.26%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.16s  0.49% 70.07%      1.28s  3.91%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.15s  0.46% 70.53%      0.17s  0.52%  runtime.spanOf (inline)
     0.14s  0.43% 70.96%      1.94s  5.93%  crypto/internal/fips140/sha256.(*Digest).Write
     0.14s  0.43% 71.38%      2.75s  8.41%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.13s   0.4% 71.78%      1.38s  4.22%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.11s  0.34% 72.12%      0.25s  0.76%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.10s  0.31% 72.42%      0.32s  0.98%  compress/flate.(*compressor).deflate
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 74.1% / 100% | 42% | 14.9% | 12 | 114 MB | CPU が飽和 |
| isu2 (2 cpu) | 14.6% / 64% | 0% | 0.4% | 4 | 1945 MB |  |
| isu3 (2 cpu) | 11.5% / 56% | 0% | 0.3% | 3 | 2447 MB |  |
| isu4 (2 cpu) | 49.9% / 100% | 17% | 20.9% | 37 | 108 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 75% / nginx 11% / du 8% / kswapd0 3% / kcompactd0 2% |
| isu2 | isureal 27% |
| isu3 | isureal 21% / nginx 1% |
| isu4 | postgres 50% / isureal 3% |

## Diff from previous run

スコア: 64700 → 83744 (+19044)

`e0e4cb3..5a138c8`

```
 conf/_all/etc/sysctl.d/99-isucon.conf              |   2 +
 conf/isu1/etc/nginx/sites-enabled/isureal.conf     |   6 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf     |   6 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf     |   6 +
 .../etc/postgresql/16/main/conf.d/zz-tuning.conf   |   2 +
 webapp/go/account_deletion_api.go                  |   2 +-
 webapp/go/admin_user_api.go                        |   2 +-
 webapp/go/comment_api.go                           |   2 +-
 webapp/go/config.go                                |   2 +
 webapp/go/export_api.go                            |   4 +-
 webapp/go/feed_api.go                              | 113 +++++--
 webapp/go/feed_test.go                             | 331 +++++++++++++++++++++
 webapp/go/offload.go                               |  27 +-
 webapp/go/post_api.go                              |  62 +++-
 webapp/go/realmoji_api.go                          |   2 +-
 webapp/go/relationship.go                          |  10 +-
 webapp/go/store.go                                 | 113 ++++++-
 17 files changed, 628 insertions(+), 64 deletions(-)
```
