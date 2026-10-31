# run 20261031-142817-cb92b74

## Summary

| | |
|---|---|
| score | 111097 |
| status | pass |
| ref | cb92b74 |
| sha | cb92b742e108 |
| 所要時間 | 131s |
| 前回比 | -28933 (-20.7%) vs 20261031-141745-a28a701 |
| ベスト比 | -30238 (-21.4%) vs 20261031-141027-bd0fe45 = 141335 |
| note | DB外し: 書き込みをPGに送らずスナップショット |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	135	355	487	486
europe	7	17	52	133	133	132
america	7	18	44	87	88	87
```

```
{"id":"17606","status":"done","portalStatus":"finished","passed":true,"score":111097,"raw":111097,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:28:44.555Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:28:45.564Z level=INFO msg="user moved" user=keivychi_5045 from=asia to=america
time=2026-10-31T05:28:47.577Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:28:50.552Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T05:28:53.562Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T05:28:55.602Z level=INFO msg="user moved" user=haro-5245 from=asia to=america
time=2026-10-31T05:28:56.552Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T05:28:57.442Z level=INFO msg="user moved" user=haro-5245 from=america to=europe
time=2026-10-31T05:28:58.205Z level=INFO msg="user moved" user=orfu-5384 from=asia to=america
time=2026-10-31T05:28:59.560Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T05:29:02.582Z level=INFO msg="asia 3 日目: 135 人"
time=2026-10-31T05:29:04.299Z level=INFO msg="user moved" user=raya_5223 from=asia to=europe
time=2026-10-31T05:29:05.557Z level=INFO msg="europe 3 日目: 52 人"
time=2026-10-31T05:29:08.552Z level=INFO msg="america 3 日目: 44 人"
time=2026-10-31T05:29:08.733Z level=INFO msg="user moved" user=keivychi_5045 from=america to=asia
time=2026-10-31T05:29:09.077Z level=INFO msg="user moved" user=keivychi_5045 from=asia to=america
time=2026-10-31T05:29:09.300Z level=INFO msg="user moved" user=haro-5245 from=europe to=america
time=2026-10-31T05:29:11.602Z level=INFO msg="asia 4 日目: 355 人"
time=2026-10-31T05:29:14.613Z level=INFO msg="europe 4 日目: 133 人"
time=2026-10-31T05:29:17.646Z level=INFO msg="america 4 日目: 87 人"
time=2026-10-31T05:29:19.880Z level=INFO msg="user moved" user=vinsamo7251 from=asia to=america
time=2026-10-31T05:29:20.677Z level=INFO msg="asia 5 日目: 487 人"
time=2026-10-31T05:29:24.177Z level=INFO msg="europe 5 日目: 133 人"
time=2026-10-31T05:29:26.983Z level=INFO msg="america 5 日目: 88 人"
time=2026-10-31T05:29:35.673Z level=INFO msg="asia 最終日（グランプリ）: 486 人"
time=2026-10-31T05:29:35.674Z level=INFO msg="europe 最終日（グランプリ）: 132 人"
time=2026-10-31T05:29:35.682Z level=INFO msg="america 最終日（グランプリ）: 87 人"
time=2026-10-31T05:29:44.568Z level=INFO msg="グランプリ: アジア 333 人, ヨーロッパ 130 人, アメリカ 85 人"
{
  "pass": true,
  "score": 111097,
  "messages": [
    "solmiame.5030: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "benka.5125: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
47.3%	1060	1	0	GET	/api/notifications	10580.821	9.982	15.085	15.122	173
15.8%	1442	11	0	POST	/api/posts	3542.740	2.457	7.790	7.979	647
7.1%	13542	324	0	GET	^/api/images/[^/]+	1581.578	0.117	0.592	1.324	126914
4.6%	10843	0	0	GET	^/api/feeds/[^/]+	1037.195	0.096	0.492	0.998	5137
3.8%	535	43	0	POST	^/api/posts/[^/]+/realmojis$	844.988	1.579	7.599	7.756	441
3.2%	3385	0	0	GET	/api/memories	716.073	0.212	1.065	1.976	12070
3.1%	4900	875	0	POST	/api/auth/login	693.642	0.142	1.276	1.822	207
2.4%	3157	66	0	POST	/api/friend-requests	539.675	0.171	1.430	1.744	113
2.1%	2141	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	476.818	0.223	1.462	1.811	116
2.1%	411	122	0	PUT	/api/me/profile-image	472.812	1.150	5.917	7.679	7
1.0%	297	0	0	GET	/api/me/export	218.204	0.735	2.759	5.331	1984888
0.9%	2071	22	0	GET	^/api/posts/[^/]+$	192.829	0.093	0.404	1.040	613
0.8%	2043	0	0	GET	/api/posting-window/current	172.198	0.084	0.555	1.122	237
0.6%	546	0	0	POST	^/api/posts/[^/]+/comments$	143.883	0.264	1.299	1.826	357
0.5%	1279	5	0	GET	^/api/posts/me/current$	119.589	0.094	0.451	1.122	650
0.5%	800	0	0	POST	/api/auth/logout	117.131	0.146	1.103	1.625	0
0.5%	1482	0	0	GET	/api/users	116.594	0.079	0.421	1.070	244
0.5%	1386	0	0	GET	^/api/relationships/[^/]+	116.578	0.084	0.496	1.111	1266
0.4%	973	1	0	GET	^/api/users/[^/]+$	89.313	0.092	0.412	0.576	200
0.4%	904	0	0	GET	^/api/posts/[^/]+/realmojis$	83.583	0.092	0.459	0.613	807
0.4%	508	1	0	POST	/api/auth/register	82.606	0.163	1.343	1.457	227
0.4%	692	0	0	GET	^/api/memories/[^/]+	81.755	0.118	0.715	0.898	594
0.4%	842	0	0	GET	^/api/posts/[^/]+/comments$	80.652	0.096	0.500	0.577	1507
0.2%	348	0	0	PUT	^/api/follows/[^/]+$	49.999	0.144	1.029	1.451	0
0.2%	60	0	0	DELETE	^/api/posts/[^/]+$	43.772	0.730	1.986	1.986	0
0.2%	72	0	0	DELETE	^/api/realmojis/[^/]+$	37.693	0.524	1.439	1.439	0
0.1%	157	0	0	PATCH	^/api/posts/[^/]+$	31.738	0.202	1.525	1.535	651
0.1%	243	0	0	PATCH	/api/me	30.317	0.125	0.587	1.215	276
0.1%	114	0	0	DELETE	^/api/comments/[^/]+$	28.800	0.253	1.263	1.497	0
0.1%	294	3	0	GET	/api/me	27.220	0.093	0.476	0.498	299
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.6%	311	0	0	GET	/api/notifications	2596.243	8.348	11.500	11.510	177
16.8%	2183	10	0	GET	^/api/images/[^/]+	1158.785	0.531	2.283	3.435	190129
6.4%	426	0	0	POST	/api/posts	443.623	1.041	3.001	3.583	645
6.2%	1614	0	0	GET	^/api/feeds/[^/]+	429.961	0.266	0.913	1.499	4004
4.7%	791	15	0	POST	/api/friend-requests	325.063	0.411	1.571	1.834	113
4.3%	69	0	0	PUT	/api/me/profile-image	297.801	4.316	7.820	7.820	0
3.5%	596	0	0	POST	/api/auth/login	243.529	0.409	1.578	1.749	249
3.2%	679	0	0	GET	/api/memories	219.297	0.323	1.260	1.622	10342
3.1%	574	0	0	GET	/api/posting-window/current	215.209	0.375	1.011	1.166	238
1.9%	311	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	134.108	0.431	1.441	1.675	116
1.6%	395	0	0	GET	/api/users	112.220	0.284	0.749	1.497	244
1.6%	423	8	0	GET	^/api/posts/me/current$	111.273	0.263	0.888	1.031	633
1.2%	45	0	0	GET	/api/me/export	82.641	1.836	6.193	6.193	2458193
0.9%	56	2	0	POST	^/api/posts/[^/]+/realmojis$	62.119	1.109	5.956	5.956	450
0.8%	144	0	0	POST	^/api/posts/[^/]+/comments$	57.967	0.403	1.877	2.139	309
0.7%	245	0	0	GET	^/api/relationships/[^/]+	51.813	0.211	0.780	0.829	1167
0.7%	204	8	0	GET	^/api/posts/[^/]+$	49.418	0.242	0.813	0.864	595
0.6%	130	0	0	POST	/api/auth/register	39.593	0.305	0.661	0.704	228
0.6%	90	0	0	POST	/api/auth/logout	39.069	0.434	1.591	1.591	0
0.5%	99	0	0	PUT	^/api/follows/[^/]+$	36.113	0.365	1.415	1.415	0
0.5%	97	0	0	GET	^/api/memories/[^/]+	32.226	0.332	1.064	1.064	899
0.4%	128	0	0	GET	^/api/users/[^/]+$	31.032	0.242	0.968	0.970	192
0.4%	75	2	0	GET	/api/me	26.201	0.349	0.847	0.847	295
0.3%	63	0	0	PATCH	/api/me	23.006	0.365	1.350	1.350	280
0.3%	99	0	0	GET	^/api/posts/[^/]+/comments$	22.142	0.224	0.966	0.966	1551
0.3%	19	0	0	DELETE	^/api/posts/[^/]+$	20.460	1.077	1.921	1.921	0
0.3%	41	1	0	PATCH	^/api/posts/[^/]+$	20.084	0.490	1.344	1.344	647
0.2%	69	0	0	GET	^/api/posts/[^/]+/realmojis$	15.663	0.227	0.696	0.696	1303
0.1%	16	0	0	DELETE	^/api/comments/[^/]+$	5.480	0.343	1.062	1.062	0
0.1%	5	0	0	DELETE	^/api/realmojis/[^/]+$	3.513	0.703	0.885	0.885	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.8%	240	1	0	GET	/api/notifications	1758.814	7.328	9.479	9.545	175
17.7%	1685	11	0	GET	^/api/images/[^/]+	801.314	0.476	2.078	2.848	189299
6.5%	307	0	0	POST	/api/posts	292.854	0.954	2.752	3.582	650
5.4%	1344	0	0	GET	^/api/feeds/[^/]+	244.656	0.182	0.689	1.163	3874
5.2%	570	13	0	POST	/api/friend-requests	237.713	0.417	1.600	1.797	113
3.5%	412	0	0	POST	/api/auth/login	159.674	0.388	1.394	1.630	251
3.2%	412	0	0	GET	/api/posting-window/current	145.379	0.353	0.940	0.961	239
2.7%	512	0	0	GET	/api/memories	123.811	0.242	1.043	1.738	11595
2.7%	271	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	121.668	0.449	1.523	1.745	116
1.9%	47	0	0	PUT	/api/me/profile-image	87.793	1.868	7.682	7.682	0
1.8%	60	8	0	POST	^/api/posts/[^/]+/realmojis$	82.867	1.381	5.679	5.679	403
1.5%	277	0	0	GET	/api/users	69.313	0.250	0.621	0.787	242
1.5%	306	16	0	GET	^/api/posts/me/current$	68.311	0.223	0.810	1.310	617
1.2%	128	0	0	POST	^/api/posts/[^/]+/comments$	52.977	0.414	1.185	1.552	330
0.7%	84	0	0	POST	/api/auth/logout	32.923	0.392	1.261	1.261	0
0.7%	203	0	0	GET	^/api/relationships/[^/]+	31.569	0.156	1.465	1.519	1229
0.6%	22	0	0	GET	/api/me/export	27.781	1.263	3.184	3.184	2998442
0.5%	186	8	0	GET	^/api/posts/[^/]+$	24.846	0.134	0.531	0.648	598
0.5%	58	0	0	PUT	^/api/follows/[^/]+$	24.039	0.414	1.600	1.600	0
0.5%	83	0	0	POST	/api/auth/register	23.008	0.277	0.885	0.885	230
0.4%	59	0	0	GET	^/api/memories/[^/]+	17.291	0.293	1.344	1.344	1043
0.4%	53	1	0	GET	/api/me	16.196	0.306	0.612	0.612	292
0.3%	37	0	0	PATCH	/api/me	14.677	0.397	1.350	1.350	278
0.3%	99	0	0	GET	^/api/users/[^/]+$	13.708	0.138	0.554	0.554	204
0.3%	14	0	0	DELETE	^/api/posts/[^/]+$	12.702	0.907	2.134	2.134	0
0.3%	92	1	0	GET	^/api/posts/[^/]+/comments$	12.107	0.132	0.477	0.477	2428
0.2%	25	4	0	PATCH	^/api/posts/[^/]+$	10.692	0.428	1.471	1.471	537
0.2%	59	0	0	GET	^/api/posts/[^/]+/realmojis$	7.795	0.132	0.506	0.506	1531
0.2%	19	0	0	DELETE	^/api/comments/[^/]+$	7.496	0.395	0.939	0.939	0
0.1%	6	0	0	DELETE	^/api/realmojis/[^/]+$	4.608	0.768	1.199	1.199	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 11443 クエリ / 48 種類 / 合計 214.757s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	43.8%	348	94.048	270.25	3413.67	1302.4	133.0	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	22.9%	5	49.169	9833.73	15479.53	31603.0	1199816.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.7%	5	31.588	6317.66	7607.99	24949.0	627165.4	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.7%	5	18.790	3757.94	6239.75	13251.0	529745.6	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
5	6.6%	9443	14.187	1.50	8.53	9.7	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
6	1.5%	6	3.234	539.07	1619.65	215460.8	2624.7	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.5%	6	1.157	192.85	639.34	624.5	25422.2	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.4%	6	0.780	130.05	228.86	624.3	30284.2	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	6	0.334	55.62	283.71	5696.8	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
10	0.1%	1	0.273	273.29	273.29	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.1%	1	0.228	227.72	227.72	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.205	205.21	205.21	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.160	159.85	159.85	32384.0	780.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.110	109.53	109.53	31798.0	3424.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.0%	1	0.095	94.59	94.59	13932.0	680.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM posts;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 901a65ca747086356950e4a11a8e9c0b5b292f22
Type: cpu
Time: 2026-10-31 14:28:26 JST
Duration: 60.02s, Total samples = 52.41s (87.32%)
Showing nodes accounting for 37.96s, 72.43% of 52.41s total
Dropped 1117 nodes (cum <= 0.26s)
Showing top 40 nodes out of 273
      flat  flat%   sum%        cum   cum%
     5.26s 10.04% 10.04%      5.27s 10.06%  crypto/internal/fips140/sha256.blockAVX2
     4.54s  8.66% 18.70%      4.54s  8.66%  internal/runtime/syscall/linux.Syscall6
     3.42s  6.53% 25.22%      3.42s  6.53%  runtime.memmove
     2.37s  4.52% 29.75%      3.12s  5.95%  runtime.findObject
     2.26s  4.31% 34.06%      9.26s 17.67%  runtime.scanObject
     2.24s  4.27% 38.33%      3.57s  6.81%  runtime.tryDeferToSpanScan
     1.78s  3.40% 41.73%      1.78s  3.40%  runtime.memclrNoHeapPointers
     1.55s  2.96% 44.69%      5.66s 10.80%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.53s  2.92% 47.61%      1.53s  2.92%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.21s  2.31% 49.91%      3.55s  6.77%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.18s  2.25% 52.17%      1.18s  2.25%  aeshashbody
     1.14s  2.18% 54.34%      1.62s  3.09%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.92s  1.76% 56.10%      0.92s  1.76%  runtime.spanClass.sizeclass (inline)
     0.69s  1.32% 57.41%      0.72s  1.37%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.58s  1.11% 58.52%      0.58s  1.11%  memeqbody
     0.55s  1.05% 59.57%      2.14s  4.08%  runtime.mapaccess1_faststr
     0.51s  0.97% 60.54%      0.51s  0.97%  runtime.nextFreeFast (inline)
     0.43s  0.82% 61.36%      0.51s  0.97%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.41s  0.78% 62.14%      0.45s  0.86%  runtime.typePointers.next
     0.40s  0.76% 62.91%      0.76s  1.45%  runtime.greyobject
     0.38s  0.73% 63.63%      0.64s  1.22%  example.com/isureal-company/isureal.(*Store).snapshotLocked
     0.35s  0.67% 64.30%      0.85s  1.62%  database/sql.convertAssignRows
     0.35s  0.67% 64.97%      0.55s  1.05%  example.com/isureal-company/isureal.postLess
     0.31s  0.59% 65.56%      0.38s  0.73%  runtime.spanOf (inline)
     0.30s  0.57% 66.13%      5.81s 11.09%  crypto/internal/fips140/sha256.(*Digest).Write
     0.30s  0.57% 66.70%      5.81s 11.09%  runtime.mallocgc
     0.29s  0.55% 67.26%      5.25s 10.02%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.29s  0.55% 67.81%      1.44s  2.75%  runtime.mallocgcSmallScanNoHeader
     0.27s  0.52% 68.33%      0.36s  0.69%  runtime.(*mspan).markBitsForIndex (inline)
     0.27s  0.52% 68.84%      0.27s  0.52%  runtime.typePointers.nextFast (inline)
     0.25s  0.48% 69.32%      1.97s  3.76%  runtime.mapassign_faststr
     0.22s  0.42% 69.74%      0.41s  0.78%  internal/runtime/maps.(*Iter).Next
     0.21s   0.4% 70.14%      0.63s  1.20%  runtime.scanObjectsSmall
     0.20s  0.38% 70.52%      0.32s  0.61%  runtime.heapArenaOf (inline)
     0.19s  0.36% 70.88%      5.40s 10.30%  encoding/gob.(*Encoder).encodeStruct
     0.18s  0.34% 71.23%      0.39s  0.74%  compress/flate.(*compressor).deflate
     0.18s  0.34% 71.57%      1.15s  2.19%  runtime.scanSpan
     0.16s  0.31% 71.88%      0.31s  0.59%  runtime.scanObjectSmall
     0.15s  0.29% 72.16%      5.50s 10.49%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.14s  0.27% 72.43%      1.09s  2.08%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 46.1% / 100% | 29% | 1.3% | 5 | 102 MB | run queue が積もっている |
| isu2 (2 cpu) | 21.5% / 91% | 1% | 1.4% | 3 | 106 MB |  |
| isu3 (2 cpu) | 22.5% / 98% | 2% | 1.0% | 3 | 192 MB |  |
| isu4 (2 cpu) | 36.8% / 100% | 10% | 16.2% | 6 | 86 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 75% / nginx 10% |
| isu2 | isureal 38% |
| isu3 | isureal 41% / nginx 1% |
| isu4 | isureal 22% / postgres 17% |

## Diff from previous run

スコア: 140030 → 111097 (-28933)

`a28a701..cb92b74`

```
 webapp/go/account_deletion_api.go     |  16 --
 webapp/go/admin_api.go                |  15 +-
 webapp/go/admin_posting_window_api.go |  46 +----
 webapp/go/admin_user_api.go           |   3 -
 webapp/go/comment_api.go              |   8 -
 webapp/go/export_api.go               |   2 +
 webapp/go/friend_request_api.go       |  33 ----
 webapp/go/image_api.go                |  17 +-
 webapp/go/image_file.go               | 132 +++++++++++++
 webapp/go/image_local.go              |  66 +++----
 webapp/go/initialize.go               |   6 +
 webapp/go/origin_store.go             | 134 +++++++++++++
 webapp/go/persist.go                  | 344 ++++++++++++++++++++++++++++++++++
 webapp/go/post_api.go                 |  27 +--
 webapp/go/profile_image_api.go        |  50 +----
 webapp/go/realmoji_api.go             |  23 +--
 webapp/go/relationship_api.go         |  35 ----
 webapp/go/repl.go                     |  11 +-
 webapp/go/replica_server.go           |  13 +-
 webapp/go/server.go                   |   6 +-
 webapp/go/store.go                    |  20 +-
 webapp/go/user_api.go                 |  36 ----
 22 files changed, 697 insertions(+), 346 deletions(-)
```
