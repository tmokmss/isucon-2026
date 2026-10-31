# run 20261031-155112-3420183

## Summary

| | |
|---|---|
| score | 184837 |
| status | pass |
| ref | export-parallel-initial-images |
| sha | 3420183f8d45 |
| 所要時間 | 293s |
| 前回比 | +1506 (0.8%) vs 20261031-152137-8c13252 |
| ベスト比 | -6979 (-3.6%) vs 20261031-151339-5a230a4 = 191816 |
| note | export-parallel-initial-images-16 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	151	439	1163	1162
europe	7	19	56	165	422	422
america	6	18	56	165	276	276
```

```
{"id":"25156","status":"done","portalStatus":"finished","passed":true,"score":184837,"raw":184837,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:54:29.719Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:54:32.708Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:54:35.707Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:54:38.722Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T06:54:41.707Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T06:54:44.708Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T06:54:47.717Z level=INFO msg="asia 3 日目: 151 人"
time=2026-10-31T06:54:49.542Z level=INFO msg="user moved" user=ulmu5788 from=asia to=america
time=2026-10-31T06:54:50.216Z level=INFO msg="user moved" user=ulmu5788 from=america to=europe
time=2026-10-31T06:54:50.708Z level=INFO msg="europe 3 日目: 56 人"
time=2026-10-31T06:54:53.708Z level=INFO msg="america 3 日目: 56 人"
time=2026-10-31T06:54:56.718Z level=INFO msg="asia 4 日目: 439 人"
time=2026-10-31T06:54:58.441Z level=INFO msg="user moved" user=anhaleo6683 from=america to=europe
time=2026-10-31T06:54:59.591Z level=INFO msg="user moved" user=miameku.7117 from=asia to=europe
time=2026-10-31T06:54:59.838Z level=INFO msg="europe 4 日目: 165 人"
time=2026-10-31T06:55:00.527Z level=INFO msg="user moved" user=lialia.6895 from=asia to=america
time=2026-10-31T06:55:01.887Z level=INFO msg="user moved" user=ulmu5788 from=europe to=asia
time=2026-10-31T06:55:02.059Z level=INFO msg="user moved" user=orki_5828 from=asia to=europe
time=2026-10-31T06:55:02.209Z level=INFO msg="user moved" user=tornolia-7648 from=asia to=america
time=2026-10-31T06:55:02.684Z level=INFO msg="user moved" user=lialia.6895 from=america to=europe
time=2026-10-31T06:55:02.852Z level=INFO msg="america 4 日目: 165 人"
time=2026-10-31T06:55:03.365Z level=INFO msg="user moved" user=tosano-7231 from=asia to=america
time=2026-10-31T06:55:04.366Z level=INFO msg="user moved" user=solyana7996 from=asia to=america
time=2026-10-31T06:55:04.646Z level=INFO msg="user moved" user=kean-8031 from=asia to=america
time=2026-10-31T06:55:05.820Z level=INFO msg="asia 5 日目: 1163 人"
time=2026-10-31T06:55:11.405Z level=INFO msg="europe 5 日目: 422 人"
time=2026-10-31T06:55:11.486Z level=INFO msg="user moved" user=kiya9728 from=america to=europe
time=2026-10-31T06:55:12.126Z level=INFO msg="america 5 日目: 276 人"
time=2026-10-31T06:55:22.103Z level=INFO msg="europe 最終日（グランプリ）: 422 人"
time=2026-10-31T06:55:22.131Z level=INFO msg="america 最終日（グランプリ）: 276 人"
time=2026-10-31T06:55:22.263Z level=INFO msg="asia 最終日（グランプリ）: 1162 人"
time=2026-10-31T06:55:29.794Z level=INFO msg="グランプリ: アジア 531 人, ヨーロッパ 186 人, アメリカ 130 人"
{
  "pass": true,
  "score": 184837,
  "messages": [
    "fusol-5085: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "rumono-13580: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maxmiake12788: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.9%	1864	1	0	GET	/api/notifications	16492.141	8.848	16.966	16.976	172
11.9%	2840	148	0	POST	/api/posts	4573.345	1.610	6.233	8.020	619
9.6%	6562	82	0	POST	/api/friend-requests	3692.958	0.563	2.726	3.665	114
7.1%	20089	342	0	GET	^/api/images/[^/]+	2731.371	0.136	1.933	3.569	115481
3.8%	7195	580	0	POST	/api/auth/login	1478.160	0.205	2.336	3.784	226
3.3%	13205	0	0	GET	^/api/feeds/[^/]+	1253.100	0.095	0.974	1.970	3043
3.2%	4873	0	0	GET	/api/memories	1245.942	0.256	2.586	4.029	13008
2.9%	3067	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1109.036	0.362	2.333	3.591	116
2.7%	1186	1	0	POST	/api/auth/register	1033.835	0.872	2.846	3.001	227
2.1%	801	135	0	PUT	/api/me/profile-image	817.463	1.021	8.352	15.663	3
1.3%	528	4	0	POST	^/api/posts/[^/]+/realmojis$	482.441	0.914	4.674	5.919	472
1.1%	4171	0	0	GET	/api/posting-window/current	405.891	0.097	1.027	1.946	237
1.0%	635	1	0	POST	^/api/posts/[^/]+/comments$	387.972	0.611	3.890	4.728	354
0.9%	3155	0	0	GET	/api/users	363.728	0.115	0.940	1.466	241
0.8%	749	0	0	PUT	^/api/follows/[^/]+$	313.846	0.419	2.731	3.455	0
0.7%	1150	0	0	POST	/api/auth/logout	251.856	0.219	2.091	3.221	0
0.6%	558	0	0	PATCH	/api/me	244.388	0.438	3.317	3.600	279
0.6%	438	0	0	GET	/api/me/export	242.450	0.554	4.217	4.623	1836772
0.6%	2227	32	0	GET	^/api/posts/me/current$	232.073	0.104	0.843	1.953	645
0.4%	1929	66	0	GET	^/api/posts/[^/]+$	155.587	0.081	0.754	1.934	615
0.4%	1617	0	0	GET	^/api/relationships/[^/]+	137.299	0.085	0.716	1.761	1144
0.3%	680	3	0	GET	/api/me	134.430	0.198	1.415	1.962	301
0.3%	127	0	0	DELETE	^/api/posts/[^/]+$	133.770	1.053	2.483	2.710	0
0.3%	251	6	0	PATCH	^/api/posts/[^/]+$	130.820	0.521	2.330	3.413	631
0.2%	1031	0	0	GET	^/api/memories/[^/]+	92.229	0.089	0.866	1.391	541
0.2%	923	0	0	GET	^/api/users/[^/]+$	70.892	0.077	0.680	1.237	218
0.1%	79	0	0	DELETE	^/api/realmojis/[^/]+$	56.019	0.709	3.455	3.455	0
0.1%	746	0	0	GET	^/api/posts/[^/]+/realmojis$	46.732	0.063	0.504	1.013	1288
0.1%	122	0	0	DELETE	^/api/comments/[^/]+$	40.014	0.328	2.086	2.255	0
0.1%	699	0	0	GET	^/api/posts/[^/]+/comments$	39.768	0.057	0.371	0.825	2086
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.9%	670	1	0	GET	/api/notifications	4538.203	6.773	11.297	11.587	173
10.9%	2117	17	0	POST	/api/friend-requests	1176.904	0.556	2.615	3.328	114
7.8%	983	36	0	POST	/api/posts	847.311	0.862	5.963	6.867	626
5.7%	4121	27	0	GET	^/api/images/[^/]+	612.685	0.149	1.399	3.278	173833
5.5%	1281	0	0	POST	/api/auth/login	599.497	0.468	3.169	3.917	246
5.3%	3603	0	0	GET	^/api/feeds/[^/]+	577.813	0.160	1.468	2.480	2943
2.8%	416	0	0	POST	/api/auth/register	298.664	0.718	2.408	2.451	229
2.3%	451	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	244.489	0.542	2.643	3.322	116
2.0%	1463	0	0	GET	/api/posting-window/current	214.536	0.147	1.058	2.033	238
1.8%	1089	0	0	GET	/api/memories	198.096	0.182	1.959	3.900	14191
1.7%	285	0	0	PUT	^/api/follows/[^/]+$	183.302	0.643	2.524	2.646	0
1.6%	225	0	0	PUT	/api/me/profile-image	178.314	0.793	2.689	3.042	0
1.4%	1065	0	0	GET	/api/users	149.432	0.140	1.452	1.922	244
1.3%	190	0	0	PATCH	/api/me	140.794	0.741	2.761	3.038	281
1.3%	132	3	0	POST	^/api/posts/[^/]+/realmojis$	140.066	1.061	3.552	3.965	474
1.2%	733	34	0	GET	^/api/posts/me/current$	127.813	0.174	1.478	1.958	621
1.1%	164	1	0	POST	^/api/posts/[^/]+/comments$	121.642	0.742	4.127	4.577	341
0.8%	505	0	0	GET	^/api/relationships/[^/]+	87.207	0.173	1.548	2.199	605
0.7%	166	0	0	POST	/api/auth/logout	80.518	0.485	3.329	3.330	0
0.7%	524	29	0	GET	^/api/posts/[^/]+$	76.861	0.147	0.953	1.727	599
0.4%	74	5	0	PATCH	^/api/posts/[^/]+$	45.829	0.619	2.654	2.654	602
0.3%	263	0	0	GET	^/api/users/[^/]+$	34.370	0.131	0.721	1.857	217
0.3%	33	0	0	DELETE	^/api/posts/[^/]+$	31.573	0.957	2.863	2.863	0
0.2%	232	2	0	GET	/api/me	25.160	0.108	0.310	0.313	301
0.2%	168	0	0	GET	^/api/posts/[^/]+/realmojis$	22.558	0.134	0.966	1.348	1576
0.2%	167	0	0	GET	^/api/posts/[^/]+/comments$	22.429	0.134	1.439	1.812	2728
0.2%	154	0	0	GET	^/api/memories/[^/]+	18.294	0.119	0.419	2.808	845
0.1%	27	0	0	DELETE	^/api/comments/[^/]+$	16.193	0.600	3.329	3.329	0
0.1%	64	0	0	GET	/api/me/export	13.196	0.206	3.227	3.227	2364048
0.1%	16	0	0	DELETE	^/api/realmojis/[^/]+$	8.163	0.510	1.082	1.082	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.0%	510	0	0	GET	/api/notifications	3503.739	6.870	10.467	10.485	169
8.9%	1506	20	0	POST	/api/friend-requests	727.112	0.483	3.192	3.356	114
8.3%	706	36	0	POST	/api/posts	679.132	0.962	5.951	6.899	622
6.7%	3743	12	0	GET	^/api/images/[^/]+	548.644	0.147	1.243	2.334	155475
6.2%	1092	0	0	POST	/api/auth/login	502.109	0.460	2.637	3.549	250
5.3%	2905	0	0	GET	^/api/feeds/[^/]+	434.519	0.150	1.448	2.274	3339
2.4%	462	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	194.274	0.421	2.666	3.334	116
2.2%	273	0	0	POST	/api/auth/register	176.937	0.648	2.395	2.456	230
2.1%	921	0	0	GET	/api/memories	168.609	0.183	1.758	3.442	17475
1.7%	1031	0	0	GET	/api/posting-window/current	142.344	0.138	0.740	1.478	239
1.4%	153	0	0	PUT	/api/me/profile-image	117.704	0.769	3.099	3.702	0
1.2%	129	0	0	POST	^/api/posts/[^/]+/comments$	100.343	0.778	3.405	3.612	367
1.2%	698	0	0	GET	/api/users	98.663	0.141	0.910	1.847	245
1.2%	163	0	0	PUT	^/api/follows/[^/]+$	95.991	0.589	2.638	2.639	0
1.2%	554	15	0	GET	^/api/posts/me/current$	94.545	0.171	1.353	1.890	637
1.1%	137	0	0	PATCH	/api/me	92.693	0.677	2.636	3.114	286
1.0%	81	1	0	POST	^/api/posts/[^/]+/realmojis$	82.368	1.017	3.982	3.982	440
0.8%	156	0	0	POST	/api/auth/logout	62.359	0.400	2.698	3.312	0
0.7%	356	0	0	GET	^/api/relationships/[^/]+	59.184	0.166	1.478	1.919	869
0.7%	458	20	0	GET	^/api/posts/[^/]+$	55.987	0.122	0.816	1.885	610
0.6%	52	5	0	PATCH	^/api/posts/[^/]+$	47.165	0.907	3.360	3.360	611
0.3%	206	0	0	GET	^/api/users/[^/]+$	24.020	0.117	0.623	1.891	211
0.3%	156	0	0	GET	/api/me	21.901	0.140	0.944	1.921	305
0.3%	17	0	0	DELETE	^/api/posts/[^/]+$	21.327	1.255	3.459	3.459	0
0.2%	65	0	0	GET	/api/me/export	20.311	0.312	2.797	2.797	3244700
0.2%	142	0	0	GET	^/api/memories/[^/]+	18.695	0.132	1.542	1.880	1117
0.2%	146	0	0	GET	^/api/posts/[^/]+/comments$	13.844	0.095	0.321	0.746	2196
0.2%	141	0	0	GET	^/api/posts/[^/]+/realmojis$	13.631	0.097	0.507	1.059	991
0.1%	21	0	0	DELETE	^/api/realmojis/[^/]+$	11.205	0.534	2.157	2.157	0
0.1%	29	0	0	DELETE	^/api/comments/[^/]+$	9.568	0.330	2.677	2.677	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 20437 クエリ / 48 種類 / 合計 177.061s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	60.2%	3537	106.599	30.14	78.84	147.7	15.0	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	12.6%	3	22.228	7409.28	10062.97	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	10.7%	14620	18.961	1.30	13.56	9.2	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	9.1%	3	16.102	5367.20	7194.42	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	4.6%	3	8.201	2733.59	3666.80	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	1.6%	4	2.805	701.27	2275.89	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.2%	4	0.355	88.74	157.12	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.2%	4	0.350	87.51	156.35	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.1%	1	0.249	248.61	248.61	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	0.1%	1	0.248	247.69	247.69	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.1%	1	0.219	218.67	218.67	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.158	157.79	157.79	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.123	123.07	123.07	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.0%	1	0.083	82.64	82.64	2463.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.068	68.40	68.40	1519.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT data FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 2a2b7ef17d7c224cb3cb013bdd08dd3a9f77d02d
Type: cpu
Time: 2026-10-31 15:54:06 JST
Duration: 60.06s, Total samples = 43.23s (71.98%)
Showing nodes accounting for 30.65s, 70.90% of 43.23s total
Dropped 1132 nodes (cum <= 0.22s)
Showing top 40 nodes out of 292
      flat  flat%   sum%        cum   cum%
     4.66s 10.78% 10.78%      4.66s 10.78%  crypto/internal/fips140/sha256.blockAVX2
     4.53s 10.48% 21.26%      4.53s 10.48%  internal/runtime/syscall/linux.Syscall6
     2.15s  4.97% 26.23%      2.15s  4.97%  runtime.memmove
     1.70s  3.93% 30.16%      1.70s  3.93%  runtime.memclrNoHeapPointers
     1.66s  3.84% 34.00%      6.70s 15.50%  runtime.scanObject
     1.65s  3.82% 37.82%      6.13s 14.18%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.65s  3.82% 41.64%      2.53s  5.85%  runtime.tryDeferToSpanScan
     1.61s  3.72% 45.36%      1.61s  3.72%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.58s  3.65% 49.02%      2.03s  4.70%  runtime.findObject
     1.46s  3.38% 52.39%      1.46s  3.38%  aeshashbody
     0.61s  1.41% 53.81%      2.35s  5.44%  runtime.mapaccess1_faststr
     0.57s  1.32% 55.12%      0.57s  1.32%  runtime.spanClass.sizeclass (inline)
     0.52s  1.20% 56.33%      0.52s  1.20%  memeqbody
     0.47s  1.09% 57.41%      0.79s  1.83%  example.com/isureal-company/isureal.postLess
     0.41s  0.95% 58.36%      0.41s  0.95%  runtime.nextFreeFast (inline)
     0.34s  0.79% 59.15%      5.21s 12.05%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.34s  0.79% 59.94%      0.82s  1.90%  database/sql.convertAssignRows
     0.32s  0.74% 60.68%      0.78s  1.80%  runtime.greyobject
     0.31s  0.72% 61.39%      0.36s  0.83%  runtime.(*mspan).markBitsForIndex (inline)
     0.29s  0.67% 62.06%      4.59s 10.62%  runtime.mallocgc
     0.28s  0.65% 62.71%      1.42s  3.28%  runtime.mallocgcSmallScanNoHeader
     0.26s   0.6% 63.31%      5.60s 12.95%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.26s   0.6% 63.91%      0.30s  0.69%  runtime.typePointers.next
     0.25s  0.58% 64.49%      5.13s 11.87%  crypto/internal/fips140/sha256.(*Digest).Write
     0.25s  0.58% 65.07%      0.25s  0.58%  runtime.futex
     0.23s  0.53% 65.60%      0.26s   0.6%  runtime.(*mspan).writeHeapBitsSmall
     0.22s  0.51% 66.11%      0.24s  0.56%  encoding/json.appendString[go.shape.string]
     0.22s  0.51% 66.62%      0.22s  0.51%  sync/atomic.(*Int32).Add
     0.21s  0.49% 67.11%      2.22s  5.14%  runtime.mapassign_faststr
     0.19s  0.44% 67.55%      0.54s  1.25%  compress/flate.(*compressor).deflate
     0.19s  0.44% 67.99%      0.24s  0.56%  runtime.heapArenaOf (inline)
     0.16s  0.37% 68.36%      0.80s  1.85%  encoding/json.structEncoder.encode
     0.16s  0.37% 68.73%      0.24s  0.56%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.16s  0.37% 69.10%      0.23s  0.53%  runtime.scanObjectSmall
     0.16s  0.37% 69.47%      0.51s  1.18%  runtime.scanObjectsSmall
     0.14s  0.32% 69.79%      0.64s  1.48%  runtime.mapaccess2_faststr
     0.13s   0.3% 70.09%      0.30s  0.69%  internal/runtime/maps.(*Iter).Next
     0.12s  0.28% 70.37%      3.06s  7.08%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.12s  0.28% 70.65%      0.67s  1.55%  runtime.wbBufFlush1
     0.11s  0.25% 70.90%      0.22s  0.51%  crypto/internal/fips140/hmac.(*HMAC).Reset
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 60.4% / 100% | 44% | 6.7% | 8 | 116 MB | CPU が飽和 |
| isu2 (2 cpu) | 31.1% / 100% | 7% | 0.8% | 6 | 168 MB | run queue が積もっている |
| isu3 (2 cpu) | 26.5% / 100% | 6% | 0.5% | 5 | 159 MB | run queue が積もっている |
| isu4 (2 cpu) | 32.8% / 100% | 7% | 5.8% | 18 | 106 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 76% / nginx 13% / find 7% / kswapd0 2% |
| isu2 | isureal 57% / nginx 2% |
| isu3 | isureal 50% / nginx 2% |
| isu4 | isureal 33% / postgres 18% |

## Diff from previous run

スコア: 183331 → 184837 (+1506)

`8c13252..3420183`

```
 webapp/go/admin_api.go      | 27 +++++++++++++---
 webapp/go/export_api.go     | 77 ++++++++++++++++++++++++++++++++++-----------
 webapp/go/image_api.go      | 23 ++++++++------
 webapp/go/image_file.go     | 33 ++++++++++---------
 webapp/go/image_held.go     | 10 ++----
 webapp/go/image_local.go    | 47 +++++++++++----------------
 webapp/go/ip_ban.go         | 11 +++++++
 webapp/go/persist.go        |  2 +-
 webapp/go/realmoji_api.go   | 41 ++++++++++--------------
 webapp/go/repl.go           |  1 +
 webapp/go/replica_server.go | 13 +++-----
 webapp/go/store.go          |  3 +-
 12 files changed, 169 insertions(+), 119 deletions(-)
```
