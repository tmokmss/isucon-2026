# run 20261031-162121-94b75aa

## Summary

| | |
|---|---|
| score | 197693 |
| status | pass |
| ref | presign-next-minute |
| sha | 94b75aa284de |
| 所要時間 | 206s |
| 前回比 | +197693 (—) vs 20261031-161627-94b75aa |
| ベスト比 | -3804 (-1.9%) vs 20261031-160409-40e65ef = 201497 |
| note | rerun: sig cache by exp generation only |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	142	383	736	735
europe	7	21	63	189	538	537
america	6	18	53	152	420	419
```

```
{"id":"27877","status":"done","portalStatus":"finished","passed":true,"score":197693,"raw":197693,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:23:17.925Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:23:20.913Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:23:23.910Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:23:26.924Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T07:23:29.923Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T07:23:32.911Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T07:23:35.928Z level=INFO msg="asia 3 日目: 142 人"
time=2026-10-31T07:23:38.911Z level=INFO msg="europe 3 日目: 63 人"
time=2026-10-31T07:23:39.433Z level=INFO msg="user moved" user=yuse.5545 from=america to=asia
time=2026-10-31T07:23:39.661Z level=INFO msg="user moved" user=reketor.6073 from=asia to=europe
time=2026-10-31T07:23:41.911Z level=INFO msg="america 3 日目: 53 人"
time=2026-10-31T07:23:44.927Z level=INFO msg="asia 4 日目: 383 人"
time=2026-10-31T07:23:47.912Z level=INFO msg="europe 4 日目: 189 人"
time=2026-10-31T07:23:49.307Z level=INFO msg="user moved" user=ivysolsu-7159 from=asia to=america
time=2026-10-31T07:23:49.947Z level=INFO msg="user moved" user=ivysolsu-7159 from=america to=asia
time=2026-10-31T07:23:50.453Z level=INFO msg="user moved" user=reketor.6073 from=europe to=asia
time=2026-10-31T07:23:50.954Z level=INFO msg="america 4 日目: 152 人"
time=2026-10-31T07:23:54.240Z level=INFO msg="asia 5 日目: 736 人"
time=2026-10-31T07:23:56.961Z level=INFO msg="user moved" user=noaeva.9544 from=america to=europe
time=2026-10-31T07:23:57.038Z level=INFO msg="europe 5 日目: 538 人"
time=2026-10-31T07:24:00.139Z level=INFO msg="america 5 日目: 420 人"
time=2026-10-31T07:24:08.933Z level=INFO msg="europe 最終日（グランプリ）: 537 人"
time=2026-10-31T07:24:09.091Z level=INFO msg="asia 最終日（グランプリ）: 735 人"
time=2026-10-31T07:24:09.091Z level=INFO msg="america 最終日（グランプリ）: 419 人"
time=2026-10-31T07:24:18.011Z level=INFO msg="グランプリ: アジア 380 人, ヨーロッパ 232 人, アメリカ 248 人"
{
  "pass": true,
  "score": 197693,
  "messages": [
    "miben-5087: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "keleotor-13268: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "homu15157: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ischi_11533: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "solto-13521: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.5%	1375	1	0	GET	/api/notifications	13657.346	9.933	14.881	14.891	171
10.8%	2056	85	0	POST	/api/posts	3227.403	1.570	7.228	7.949	625
9.0%	10409	0	0	GET	^/api/feeds/[^/]+	2692.252	0.259	1.510	2.559	4297
7.6%	2882	0	0	GET	/api/memories	2292.890	0.796	5.357	14.214	11054
5.5%	2865	0	0	GET	/api/posting-window/current	1639.597	0.572	2.187	2.739	237
3.1%	2027	0	0	GET	/api/users	939.659	0.464	1.802	2.554	240
3.1%	1646	2	0	GET	^/api/posts/me/current$	928.176	0.564	1.589	1.934	651
2.2%	218	0	0	GET	/api/me/export	672.644	3.086	13.750	15.650	2064123
1.9%	11839	284	0	GET	^/api/images/[^/]+	571.827	0.048	0.741	2.358	165121
1.4%	537	127	0	PUT	/api/me/profile-image	422.302	0.786	11.165	17.048	4
1.3%	4243	52	0	POST	/api/friend-requests	389.244	0.092	0.804	1.338	114
1.1%	1246	0	0	GET	^/api/relationships/[^/]+	336.030	0.270	1.402	2.988	1479
1.1%	675	9	0	POST	^/api/posts/[^/]+/realmojis$	331.753	0.491	5.615	6.101	458
1.0%	3414	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	296.707	0.087	0.791	1.391	116
1.0%	422	3	0	GET	/api/me	293.127	0.695	2.277	2.574	300
0.8%	790	2	0	GET	^/api/posts/[^/]+/realmojis$	231.790	0.293	1.386	1.852	1722
0.8%	1772	71	0	GET	^/api/posts/[^/]+$	228.337	0.129	1.019	1.457	611
0.6%	553	0	0	GET	^/api/memories/[^/]+	190.027	0.344	1.975	2.580	668
0.5%	4005	471	0	POST	/api/auth/login	145.348	0.036	0.738	1.479	218
0.4%	913	0	0	GET	^/api/users/[^/]+$	120.127	0.132	0.985	1.411	208
0.4%	688	0	0	GET	^/api/posts/[^/]+/comments$	109.241	0.159	1.084	1.420	2264
0.3%	649	0	0	POST	^/api/posts/[^/]+/comments$	75.091	0.116	0.836	1.482	337
0.1%	80	0	0	DELETE	^/api/posts/[^/]+$	39.853	0.498	1.593	1.593	0
0.1%	588	0	0	POST	/api/auth/logout	31.921	0.054	0.791	0.974	0
0.1%	179	2	0	PATCH	^/api/posts/[^/]+$	25.355	0.142	0.802	0.826	642
0.1%	99	0	0	DELETE	^/api/realmojis/[^/]+$	23.160	0.234	0.728	0.728	0
0.1%	1	0	0	POST	/api/initialize	22.756	22.756	22.756	22.756	90
0.1%	520	0	0	PUT	^/api/follows/[^/]+$	17.922	0.034	0.681	0.949	0
0.0%	351	0	0	PATCH	/api/me	11.424	0.033	0.268	0.854	274
0.0%	122	0	0	DELETE	^/api/comments/[^/]+$	8.768	0.072	0.850	1.212	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.1%	841	1	0	GET	/api/notifications	5969.743	7.098	11.846	12.019	174
9.4%	1355	50	0	POST	/api/posts	1242.344	0.917	7.568	8.060	629
7.6%	3013	30	0	POST	/api/friend-requests	1009.185	0.335	1.126	1.924	114
7.2%	6374	0	0	GET	^/api/feeds/[^/]+	946.463	0.148	1.952	2.708	3292
5.3%	6084	51	0	GET	^/api/images/[^/]+	701.076	0.115	1.053	3.305	174396
3.8%	299	0	0	PUT	/api/me/profile-image	508.430	1.700	6.156	6.295	0
3.2%	1880	0	0	POST	/api/auth/login	424.170	0.226	0.970	1.976	247
2.3%	1498	0	0	GET	/api/memories	309.671	0.207	2.199	3.227	13610
1.8%	1936	0	0	GET	/api/posting-window/current	232.613	0.120	0.785	2.677	238
1.6%	986	9	0	GET	^/api/posts/me/current$	215.143	0.218	1.869	2.357	646
1.4%	334	11	0	POST	^/api/posts/[^/]+/realmojis$	189.931	0.569	5.796	6.002	452
1.4%	572	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	187.862	0.328	1.698	2.045	116
1.4%	535	0	0	POST	/api/auth/register	187.320	0.350	0.840	0.869	229
1.3%	1447	0	0	GET	/api/users	175.273	0.121	0.767	2.683	243
1.1%	355	0	0	PUT	^/api/follows/[^/]+$	140.342	0.395	1.326	1.477	0
0.9%	840	0	0	GET	^/api/relationships/[^/]+	116.893	0.139	1.908	2.526	524
0.7%	1041	25	0	GET	^/api/posts/[^/]+$	96.932	0.093	0.255	1.841	624
0.7%	389	2	0	POST	^/api/posts/[^/]+/comments$	91.317	0.235	0.985	1.442	332
0.7%	263	0	0	PATCH	/api/me	87.376	0.332	0.932	0.957	276
0.5%	289	0	0	POST	/api/auth/logout	60.426	0.209	0.949	1.605	0
0.4%	307	2	0	GET	/api/me	47.420	0.154	0.833	0.844	303
0.3%	446	0	0	GET	^/api/users/[^/]+$	41.905	0.094	0.272	0.366	208
0.3%	419	1	0	GET	^/api/posts/[^/]+/realmojis$	38.977	0.093	0.303	1.241	1935
0.3%	110	5	0	PATCH	^/api/posts/[^/]+$	36.007	0.327	1.821	1.868	624
0.3%	372	0	0	GET	^/api/posts/[^/]+/comments$	34.473	0.093	0.259	0.293	3008
0.2%	55	0	0	DELETE	^/api/posts/[^/]+$	33.058	0.601	2.408	2.408	0
0.2%	114	0	0	GET	/api/me/export	28.727	0.252	5.192	5.528	2416848
0.2%	277	0	0	GET	^/api/memories/[^/]+	27.161	0.098	0.857	1.142	805
0.1%	84	0	0	DELETE	^/api/comments/[^/]+$	17.522	0.209	0.987	0.987	0
0.1%	43	0	0	DELETE	^/api/realmojis/[^/]+$	15.880	0.369	1.288	1.288	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.2%	652	1	0	GET	/api/notifications	3810.432	5.844	9.177	9.199	173
9.7%	1019	20	0	POST	/api/posts	879.515	0.863	6.940	7.703	634
7.3%	3818	0	0	GET	^/api/feeds/[^/]+	663.705	0.174	1.732	2.362	3157
7.0%	5045	31	0	GET	^/api/images/[^/]+	630.937	0.125	1.062	3.243	142949
6.7%	2352	21	0	POST	/api/friend-requests	600.741	0.255	1.083	1.879	114
5.2%	1571	0	0	POST	/api/auth/login	465.158	0.296	1.263	1.872	249
2.7%	1364	0	0	GET	/api/memories	244.084	0.179	2.006	3.562	16111
2.6%	421	0	0	POST	/api/auth/register	235.878	0.560	1.246	1.406	229
2.4%	1475	0	0	GET	/api/posting-window/current	218.434	0.148	1.108	1.937	240
2.2%	213	0	0	PUT	/api/me/profile-image	200.898	0.943	6.261	6.298	0
1.8%	1192	0	0	GET	/api/users	166.690	0.140	0.912	2.043	239
1.6%	767	26	0	GET	^/api/posts/me/current$	144.108	0.188	1.739	2.194	622
1.0%	352	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	89.641	0.255	1.011	1.230	116
0.9%	275	0	0	PUT	^/api/follows/[^/]+$	82.858	0.301	0.975	1.282	0
0.9%	499	0	0	GET	^/api/relationships/[^/]+	77.894	0.156	1.836	2.313	582
0.8%	108	10	0	POST	^/api/posts/[^/]+/realmojis$	71.864	0.665	5.851	5.878	402
0.7%	248	0	0	POST	/api/auth/logout	60.191	0.243	1.070	1.608	0
0.7%	180	0	0	PATCH	/api/me	60.067	0.334	1.083	1.101	274
0.6%	519	18	0	GET	^/api/posts/[^/]+$	56.571	0.109	0.373	2.282	610
0.6%	226	3	0	POST	^/api/posts/[^/]+/comments$	53.045	0.235	0.931	1.072	318
0.4%	80	9	0	PATCH	^/api/posts/[^/]+$	34.485	0.431	1.620	1.620	585
0.3%	246	0	0	GET	^/api/users/[^/]+$	30.610	0.124	0.835	0.915	202
0.3%	225	1	0	GET	^/api/posts/[^/]+/realmojis$	26.076	0.116	0.560	1.007	1522
0.3%	233	0	0	GET	^/api/posts/[^/]+/comments$	24.389	0.105	0.396	0.735	1962
0.3%	218	1	0	GET	/api/me	24.197	0.111	0.498	0.800	303
0.3%	218	0	0	GET	^/api/memories/[^/]+	23.703	0.109	0.472	1.995	921
0.2%	34	0	0	DELETE	^/api/posts/[^/]+$	21.562	0.634	2.263	2.263	0
0.2%	74	0	0	GET	/api/me/export	17.225	0.233	4.218	4.218	2811002
0.1%	33	0	0	DELETE	^/api/comments/[^/]+$	6.288	0.191	0.436	0.436	0
0.1%	11	0	0	DELETE	^/api/realmojis/[^/]+$	5.192	0.472	1.004	1.004	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 10842 クエリ / 48 種類 / 合計 140.916s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	44.8%	224	63.197	282.13	2286.01	1353.6	138.0	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	18.3%	4	25.738	6434.47	9548.51	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	13.7%	4	19.278	4819.61	7264.86	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.4%	7387	11.790	1.60	10.12	10.2	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.3%	4	8.851	2212.73	3440.71	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	4.9%	7	6.943	991.87	4008.51	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	1.0%	1	1.390	1390.24	1390.24	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.8%	7	1.072	153.20	615.58	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.4%	7	0.620	88.53	163.16	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.3%	1	0.427	427.32	427.32	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.343	342.58	342.58	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.2%	1	0.271	270.73	270.73	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.188	187.64	187.64	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.117	116.81	116.81	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	7	0.112	16.05	32.53	166.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: 86ee2acaab68161eeb395570172a7829f04d19ba
Type: cpu
Time: 2026-10-31 16:22:49 JST
Duration: 60s, Total samples = 23.76s (39.60%)
Showing nodes accounting for 18.15s, 76.39% of 23.76s total
Dropped 764 nodes (cum <= 0.12s)
Showing top 40 nodes out of 243
      flat  flat%   sum%        cum   cum%
     3.94s 16.58% 16.58%      3.94s 16.58%  crypto/internal/fips140/sha256.blockAVX2
     2.25s  9.47% 26.05%      2.25s  9.47%  internal/runtime/syscall/linux.Syscall6
     1.57s  6.61% 32.66%      5.80s 24.41%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.39s  5.85% 38.51%      1.39s  5.85%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.21s  5.09% 43.60%      1.21s  5.09%  aeshashbody
     1.10s  4.63% 48.23%      1.10s  4.63%  runtime.memmove
     0.72s  3.03% 51.26%      0.72s  3.03%  runtime.memclrNoHeapPointers
     0.58s  2.44% 53.70%      2.20s  9.26%  runtime.scanObject
     0.56s  2.36% 56.06%      0.73s  3.07%  runtime.tryDeferToSpanScan
     0.49s  2.06% 58.12%      1.87s  7.87%  runtime.mapaccess1_faststr
     0.46s  1.94% 60.06%      0.60s  2.53%  runtime.findObject
     0.30s  1.26% 61.32%      0.49s  2.06%  example.com/isureal-company/isureal.postLess
     0.29s  1.22% 62.54%      0.43s  1.81%  database/sql.convertAssignRows
     0.25s  1.05% 63.59%      0.25s  1.05%  memeqbody
     0.23s  0.97% 64.56%      4.39s 18.48%  crypto/internal/fips140/sha256.(*Digest).Write
     0.23s  0.97% 65.53%      4.49s 18.90%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.23s  0.97% 66.50%      0.23s  0.97%  sync/atomic.(*Int32).Add
     0.18s  0.76% 67.26%      0.18s  0.76%  runtime.futex
     0.17s  0.72% 67.97%      4.73s 19.91%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.17s  0.72% 68.69%      1.40s  5.89%  runtime.mallocgc
     0.16s  0.67% 69.36%      0.20s  0.84%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.15s  0.63% 69.99%      1.86s  7.83%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.15s  0.63% 70.62%      0.17s  0.72%  runtime.(*mspan).writeHeapBitsSmall
     0.14s  0.59% 71.21%      1.65s  6.94%  runtime.mapassign_faststr
     0.13s  0.55% 71.76%      0.61s  2.57%  runtime.mallocgcSmallScanNoHeader
     0.12s  0.51% 72.26%      0.12s  0.51%  runtime.(*spanInlineMarkBits).init
     0.11s  0.46% 72.73%      0.18s  0.76%  internal/runtime/maps.(*Iter).Next
     0.11s  0.46% 73.19%      0.12s  0.51%  runtime.(*mspan).markBitsForIndex (inline)
     0.10s  0.42% 73.61%      0.12s  0.51%  runtime.typePointers.next
     0.08s  0.34% 73.95%      0.84s  3.54%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.08s  0.34% 74.28%      0.27s  1.14%  runtime.greyobject
     0.08s  0.34% 74.62%      0.29s  1.22%  runtime.mapaccess2_faststr
     0.06s  0.25% 74.87%      4.92s 20.71%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.06s  0.25% 75.13%      0.16s  0.67%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.05s  0.21% 75.34%      0.14s  0.59%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.05s  0.21% 75.55%      3.99s 16.79%  crypto/internal/fips140/sha256.block
     0.05s  0.21% 75.76%      0.48s  2.02%  database/sql.(*Rows).scanLocked
     0.05s  0.21% 75.97%      0.38s  1.60%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func3
     0.05s  0.21% 76.18%      0.47s  1.98%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
     0.05s  0.21% 76.39%      0.12s  0.51%  runtime.newobject
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 41.2% / 100% | 21% | 2.9% | 6 | 109 MB | run queue が積もっている |
| isu2 (2 cpu) | 27.8% / 100% | 9% | 0.7% | 4 | 137 MB |  |
| isu3 (2 cpu) | 27.1% / 100% | 8% | 0.9% | 4 | 189 MB |  |
| isu4 (2 cpu) | 67.3% / 100% | 59% | 3.3% | 4 | 115 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 60% / nginx 10% / find 2% |
| isu2 | isureal 51% / nginx 3% |
| isu3 | isureal 50% / nginx 2% |
| isu4 | isureal 109% / postgres 17% |

## Diff from previous run

スコア: 0 → 197693 (+197693)

コードの差分は無い（同じ sha 94b75aa284de）。
スコアの違いはブレか、サーバ側の設定変更による。
