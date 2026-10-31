# run 20261031-155036-b19d232

## Summary

| | |
|---|---|
| score | 170564 |
| status | pass |
| ref | presign-next-minute |
| sha | b19d23279da5 |
| 所要時間 | 204s |
| 前回比 | -8052 (-4.5%) vs 20261031-154736-f1153ff |
| ベスト比 | -21252 (-11.1%) vs 20261031-151339-5a230a4 = 191816 |
| note | on b3440cf: sig cache + presign + JPEG prefers asia2 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	52	155	450	1013	1013
europe	7	19	55	154	432	432
america	6	16	52	147	418	418
```

```
{"id":"24955","status":"done","portalStatus":"finished","passed":true,"score":170564,"raw":170564,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:52:18.230Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:52:21.235Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:52:24.220Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:52:27.234Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T06:52:30.227Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T06:52:33.222Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T06:52:34.469Z level=INFO msg="user moved" user=here-5386 from=europe to=asia
time=2026-10-31T06:52:36.234Z level=INFO msg="asia 3 日目: 155 人"
time=2026-10-31T06:52:37.658Z level=INFO msg="user moved" user=here-5386 from=asia to=america
time=2026-10-31T06:52:38.381Z level=INFO msg="user moved" user=fumarya.5723 from=asia to=america
time=2026-10-31T06:52:38.940Z level=INFO msg="user moved" user=hosemi.6026 from=asia to=europe
time=2026-10-31T06:52:39.245Z level=INFO msg="europe 3 日目: 55 人"
time=2026-10-31T06:52:39.972Z level=INFO msg="user moved" user=manudan6095 from=asia to=america
time=2026-10-31T06:52:41.940Z level=INFO msg="user moved" user=fumarya.5723 from=america to=asia
time=2026-10-31T06:52:42.222Z level=INFO msg="america 3 日目: 52 人"
time=2026-10-31T06:52:43.583Z level=INFO msg="user moved" user=neshi_6351 from=europe to=america
time=2026-10-31T06:52:45.253Z level=INFO msg="asia 4 日目: 450 人"
time=2026-10-31T06:52:48.358Z level=INFO msg="europe 4 日目: 154 人"
time=2026-10-31T06:52:49.225Z level=INFO msg="user moved" user=yare_7712 from=asia to=america
time=2026-10-31T06:52:50.423Z level=INFO msg="user moved" user=yare_7712 from=america to=asia
time=2026-10-31T06:52:50.964Z level=INFO msg="user moved" user=hosemi.6026 from=europe to=asia
time=2026-10-31T06:52:51.795Z level=INFO msg="america 4 日目: 147 人"
time=2026-10-31T06:52:53.676Z level=INFO msg="user moved" user=kose.9557 from=europe to=asia
time=2026-10-31T06:52:53.733Z level=INFO msg="user moved" user=seri_9553 from=america to=asia
time=2026-10-31T06:52:54.416Z level=INFO msg="user moved" user=haleo.9013 from=europe to=america
time=2026-10-31T06:52:54.812Z level=INFO msg="user moved" user=maevaha-9488 from=europe to=asia
time=2026-10-31T06:52:56.251Z level=INFO msg="user moved" user=neshi_6351 from=america to=europe
time=2026-10-31T06:52:56.806Z level=INFO msg="user moved" user=here-5386 from=america to=asia
time=2026-10-31T06:52:57.257Z level=INFO msg="user moved" user=helia9710 from=america to=asia
time=2026-10-31T06:52:58.027Z level=INFO msg="user moved" user=manudan6095 from=america to=europe
time=2026-10-31T06:52:58.851Z level=INFO msg="asia 5 日目: 1013 人"
time=2026-10-31T06:52:58.881Z level=INFO msg="europe 5 日目: 432 人"
time=2026-10-31T06:53:01.092Z level=INFO msg="america 5 日目: 418 人"
time=2026-10-31T06:53:10.078Z level=INFO msg="america 最終日（グランプリ）: 418 人"
time=2026-10-31T06:53:10.104Z level=INFO msg="asia 最終日（グランプリ）: 1013 人"
time=2026-10-31T06:53:10.164Z level=INFO msg="europe 最終日（グランプリ）: 432 人"
time=2026-10-31T06:53:18.344Z level=INFO msg="グランプリ: アジア 473 人, ヨーロッパ 202 人, アメリカ 230 人"
{
  "pass": true,
  "score": 170564,
  "messages": [
    "tordan_17266: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yator9774: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hetaho11361: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yusolro14536: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "torleoke.13945: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yutamu-13158: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "zoehi-14175: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tormeru-13489: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mana-13591: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "shiroivy-11977: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rimaxnoa.17103: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tekuwa15237: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "vinsayo_7270: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kachi.17436: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nowake-17177: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "resama_10889: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "roan-15527: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kanu12820: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "heisvin-16056: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "moho.12779: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
47.7%	1626	1	0	GET	/api/notifications	14567.701	8.959	13.704	13.711	152
11.1%	2240	141	0	POST	/api/posts	3376.695	1.507	7.016	65.732	603
6.9%	5060	53	0	POST	/api/friend-requests	2113.384	0.418	4.164	4.960	114
6.3%	15914	383	0	GET	^/api/images/[^/]+	1919.887	0.121	1.603	5.332	100152
3.8%	695	139	0	PUT	/api/me/profile-image	1168.797	1.682	16.939	66.636	3
3.4%	3943	0	0	GET	/api/memories	1047.605	0.266	1.858	8.028	12387
3.3%	6034	674	0	POST	/api/auth/login	1022.027	0.169	1.785	4.402	219
3.1%	8087	0	0	GET	^/api/feeds/[^/]+	933.829	0.115	1.165	3.111	2638
2.6%	2515	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	796.598	0.317	3.410	4.864	116
1.3%	1025	1	0	POST	/api/auth/register	402.556	0.393	1.264	4.375	227
1.1%	2488	0	0	GET	/api/users	332.628	0.134	1.165	1.630	238
1.0%	3319	0	0	GET	/api/posting-window/current	316.653	0.095	0.644	1.806	236
1.0%	450	1	0	POST	^/api/posts/[^/]+/realmojis$	303.149	0.674	4.533	5.183	472
1.0%	639	0	0	PUT	^/api/follows/[^/]+$	293.514	0.459	3.410	4.857	0
1.0%	335	0	0	GET	/api/me/export	293.073	0.875	8.889	12.296	1773454
0.9%	447	0	0	PATCH	/api/me	288.462	0.645	4.382	4.763	278
0.6%	568	3	0	GET	/api/me	181.266	0.319	0.965	0.984	300
0.6%	1743	110	0	GET	^/api/posts/me/current$	178.886	0.103	0.595	1.806	610
0.4%	906	0	0	POST	/api/auth/logout	133.957	0.148	1.382	4.697	0
0.4%	296	0	0	POST	^/api/posts/[^/]+/comments$	132.836	0.449	4.381	4.474	334
0.4%	1063	49	0	GET	^/api/posts/[^/]+$	129.857	0.122	1.165	1.772	593
0.4%	1138	0	0	GET	^/api/relationships/[^/]+	118.185	0.104	0.981	1.603	1457
0.3%	862	0	0	GET	^/api/memories/[^/]+	95.717	0.111	1.210	1.713	677
0.3%	185	10	0	PATCH	^/api/posts/[^/]+$	94.630	0.512	4.330	4.400	628
0.3%	83	0	0	DELETE	^/api/posts/[^/]+$	93.517	1.127	4.329	4.329	0
0.2%	472	0	0	GET	^/api/users/[^/]+$	60.004	0.127	1.308	1.903	195
0.1%	402	0	0	GET	^/api/posts/[^/]+/realmojis$	40.545	0.101	0.986	1.921	1458
0.1%	94	0	0	DELETE	^/api/realmojis/[^/]+$	38.669	0.411	3.159	3.159	0
0.1%	302	0	0	GET	^/api/posts/[^/]+/comments$	29.379	0.097	0.779	1.630	1407
0.1%	1	0	0	POST	/api/initialize	15.882	15.882	15.882	15.882	90
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.5%	679	0	0	GET	/api/notifications	4318.081	6.359	11.304	11.553	168
11.3%	2150	25	0	POST	/api/friend-requests	1153.265	0.536	3.079	3.817	114
9.5%	912	50	0	POST	/api/posts	967.682	1.061	7.145	65.690	615
5.3%	5526	23	0	GET	^/api/images/[^/]+	543.591	0.098	1.028	2.226	125906
5.0%	1765	0	0	POST	/api/auth/login	503.869	0.285	1.962	3.501	248
4.6%	3845	0	0	GET	^/api/feeds/[^/]+	464.417	0.121	0.926	1.706	3046
2.7%	433	0	0	POST	/api/auth/register	270.763	0.625	1.732	1.775	229
2.2%	248	0	0	PUT	/api/me/profile-image	219.201	0.884	2.307	2.710	0
2.0%	1424	0	0	GET	/api/posting-window/current	205.722	0.144	1.097	1.672	239
1.8%	267	0	0	PUT	^/api/follows/[^/]+$	187.765	0.703	3.082	3.786	0
1.8%	1297	0	0	GET	/api/memories	182.703	0.141	1.744	4.548	15866
1.6%	1039	0	0	GET	/api/users	166.716	0.160	1.170	1.583	242
1.5%	210	0	0	PATCH	/api/me	156.428	0.745	3.115	3.183	277
1.4%	444	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	144.852	0.326	2.434	2.773	116
1.1%	703	25	0	GET	^/api/posts/me/current$	106.763	0.152	1.311	1.671	627
0.8%	536	0	0	GET	^/api/relationships/[^/]+	84.419	0.157	1.543	1.659	656
0.7%	181	0	0	POST	^/api/posts/[^/]+/comments$	76.062	0.420	3.423	3.901	361
0.7%	157	0	0	POST	^/api/posts/[^/]+/realmojis$	67.422	0.429	3.248	3.599	485
0.5%	297	0	0	POST	/api/auth/logout	53.126	0.179	1.336	2.109	0
0.5%	520	13	0	GET	^/api/posts/[^/]+$	47.644	0.092	0.417	1.078	618
0.4%	254	1	0	GET	/api/me	40.081	0.158	1.162	1.300	304
0.4%	76	5	0	PATCH	^/api/posts/[^/]+$	36.390	0.479	1.850	1.850	611
0.3%	28	0	0	DELETE	^/api/posts/[^/]+$	28.455	1.016	2.925	2.925	0
0.3%	260	0	0	GET	^/api/memories/[^/]+	27.261	0.105	1.005	1.916	946
0.3%	290	0	0	GET	^/api/users/[^/]+$	25.840	0.089	0.501	1.098	199
0.2%	242	0	0	GET	^/api/posts/[^/]+/realmojis$	21.277	0.088	0.688	0.839	2022
0.2%	114	0	0	GET	/api/me/export	19.957	0.175	2.132	2.700	2171893
0.1%	195	0	0	GET	^/api/posts/[^/]+/comments$	14.777	0.076	0.557	0.889	2233
0.1%	28	0	0	DELETE	^/api/realmojis/[^/]+$	10.513	0.375	2.695	2.695	0
0.1%	36	0	0	DELETE	^/api/comments/[^/]+$	5.942	0.165	0.685	0.685	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.1%	625	0	0	GET	/api/notifications	3658.217	5.853	9.905	9.907	154
8.8%	2064	11	0	POST	/api/friend-requests	763.049	0.370	2.859	3.542	114
8.3%	797	40	0	POST	/api/posts	721.848	0.906	6.822	65.428	617
6.8%	1510	0	0	POST	/api/auth/login	587.513	0.389	2.310	3.809	249
5.8%	4551	21	0	GET	^/api/images/[^/]+	500.672	0.110	0.999	2.774	133759
4.4%	3150	0	0	GET	^/api/feeds/[^/]+	378.870	0.120	1.195	1.738	2939
3.4%	418	0	0	POST	/api/auth/register	292.424	0.700	1.545	1.593	230
2.7%	1303	0	0	GET	/api/posting-window/current	233.982	0.180	1.088	1.522	240
2.6%	214	0	0	PUT	/api/me/profile-image	227.156	1.061	3.465	3.587	0
2.1%	194	0	0	PATCH	/api/me	180.921	0.933	3.132	4.159	275
1.8%	1119	0	0	GET	/api/memories	158.363	0.142	1.498	6.008	16621
1.8%	236	0	0	PUT	^/api/follows/[^/]+$	155.233	0.658	3.188	3.597	0
1.6%	1026	0	0	GET	/api/users	143.304	0.140	1.241	1.478	232
1.0%	334	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	90.401	0.271	2.824	2.919	116
1.0%	637	23	0	GET	^/api/posts/me/current$	89.312	0.140	1.282	1.400	625
0.8%	173	3	0	POST	^/api/posts/[^/]+/realmojis$	72.801	0.421	3.789	3.807	474
0.8%	65	2	0	PATCH	^/api/posts/[^/]+$	72.439	1.114	3.045	3.045	632
0.8%	242	0	0	POST	/api/auth/logout	69.308	0.286	2.806	2.972	0
0.7%	420	0	0	GET	^/api/relationships/[^/]+	58.363	0.139	1.329	1.620	560
0.4%	219	0	0	GET	/api/me	37.831	0.173	0.798	0.885	305
0.4%	438	12	0	GET	^/api/posts/[^/]+$	35.684	0.081	0.575	0.975	608
0.4%	137	0	0	POST	^/api/posts/[^/]+/comments$	32.439	0.237	2.566	2.863	357
0.4%	29	0	0	DELETE	^/api/posts/[^/]+$	31.396	1.083	3.378	3.378	0
0.3%	85	0	0	GET	/api/me/export	21.959	0.258	2.231	2.231	3033169
0.2%	216	0	0	GET	^/api/users/[^/]+$	20.880	0.097	0.812	1.191	192
0.2%	216	0	0	GET	^/api/memories/[^/]+	19.342	0.090	0.593	1.094	1080
0.1%	145	0	0	GET	^/api/posts/[^/]+/realmojis$	11.604	0.080	0.700	0.921	2329
0.1%	30	0	0	DELETE	^/api/realmojis/[^/]+$	10.071	0.336	1.184	1.184	0
0.1%	112	0	0	GET	^/api/posts/[^/]+/comments$	9.872	0.088	0.751	0.867	2786
0.1%	36	0	0	DELETE	^/api/comments/[^/]+$	4.813	0.134	0.246	0.246	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 14493 クエリ / 48 種類 / 合計 134.300s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	50.7%	346	68.070	196.73	1335.69	1125.2	114.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	16.5%	3	22.222	7407.26	10025.59	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	12.1%	3	16.251	5417.02	7241.78	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	10.9%	11737	14.690	1.25	6.57	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	5.9%	3	7.899	2633.02	3559.24	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.2%	4	2.948	736.91	2458.91	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.3%	4	0.349	87.29	157.51	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.3%	4	0.344	85.93	156.92	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.273	272.93	272.93	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.253	252.95	252.95	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.2%	1	0.236	235.93	235.93	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.166	166.45	166.45	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.107	106.72	106.72	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.082	81.95	81.95	2486.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.1%	1	0.077	76.92	76.92	1549.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
Build ID: 31a31944abd3be23b8889ea358b16e290eefa632
Type: cpu
Time: 2026-10-31 15:51:57 JST
Duration: 60s, Total samples = 36.99s (61.65%)
Showing nodes accounting for 26.07s, 70.48% of 36.99s total
Dropped 1076 nodes (cum <= 0.18s)
Showing top 40 nodes out of 286
      flat  flat%   sum%        cum   cum%
     5.15s 13.92% 13.92%      5.15s 13.92%  internal/runtime/syscall/linux.Syscall6
        2s  5.41% 19.33%         2s  5.41%  runtime.memmove
     1.93s  5.22% 24.55%      1.93s  5.22%  crypto/internal/fips140/sha256.blockAVX2
     1.81s  4.89% 29.44%      1.81s  4.89%  runtime.memclrNoHeapPointers
     1.57s  4.24% 33.68%      5.91s 15.98%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.54s  4.16% 37.85%      1.54s  4.16%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.37s  3.70% 41.55%      1.73s  4.68%  runtime.findObject
     1.36s  3.68% 45.23%      1.36s  3.68%  aeshashbody
     1.34s  3.62% 48.85%      5.42s 14.65%  runtime.scanObject
     1.29s  3.49% 52.34%      2.08s  5.62%  runtime.tryDeferToSpanScan
     0.67s  1.81% 54.15%      2.37s  6.41%  runtime.mapaccess1_faststr
     0.58s  1.57% 55.72%      0.58s  1.57%  runtime.spanClass.sizeclass (inline)
     0.44s  1.19% 56.91%      0.44s  1.19%  memeqbody
     0.39s  1.05% 57.96%      0.60s  1.62%  example.com/isureal-company/isureal.postLess
     0.35s  0.95% 58.91%      0.35s  0.95%  runtime.nextFreeFast (inline)
     0.33s  0.89% 59.80%      0.83s  2.24%  database/sql.convertAssignRows
     0.30s  0.81% 60.61%      0.30s  0.81%  runtime.futex
     0.27s  0.73% 61.34%      1.09s  2.95%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.65% 61.99%      0.24s  0.65%  sync/atomic.(*Int32).Add (inline)
     0.23s  0.62% 62.61%      0.40s  1.08%  compress/flate.(*compressor).deflate
     0.22s  0.59% 63.21%      0.46s  1.24%  runtime.greyobject
     0.20s  0.54% 63.75%      0.20s  0.54%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.19s  0.51% 64.26%      3.76s 10.16%  runtime.mallocgc
     0.19s  0.51% 64.77%      0.19s  0.51%  runtime.typePointers.nextFast (inline)
     0.18s  0.49% 65.26%      1.79s  4.84%  runtime.mapassign_faststr
     0.17s  0.46% 65.72%      0.56s  1.51%  runtime.mapaccess2_faststr
     0.15s  0.41% 66.13%      0.20s  0.54%  runtime.spanOf (inline)
     0.14s  0.38% 66.50%      2.43s  6.57%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.14s  0.38% 66.88%      2.24s  6.06%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.14s  0.38% 67.26%      3.06s  8.27%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.14s  0.38% 67.64%      1.29s  3.49%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.14s  0.38% 68.02%      0.20s  0.54%  runtime.(*mspan).markBitsForIndex (inline)
     0.14s  0.38% 68.40%      0.38s  1.03%  runtime.scanObjectsSmall
     0.13s  0.35% 68.75%      0.24s  0.65%  internal/runtime/maps.(*Iter).Next
     0.12s  0.32% 69.07%      0.80s  2.16%  runtime.mallocgcSmallNoscan
     0.11s   0.3% 69.37%      2.57s  6.95%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.11s   0.3% 69.67%      0.94s  2.54%  database/sql.(*Rows).scanLocked
     0.11s   0.3% 69.96%      4.19s 11.33%  example.com/isureal-company/isureal.loadBig
     0.10s  0.27% 70.24%      0.62s  1.68%  runtime.newobject
     0.09s  0.24% 70.48%      1.11s  3.00%  encoding/gob.(*Encoder).encodeStruct
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 75.1% / 100% | 36% | 31.4% | 8 | 109 MB | CPU が飽和 |
| isu2 (2 cpu) | 30.5% / 100% | 3% | 0.7% | 4 | 170 MB |  |
| isu3 (2 cpu) | 23.7% / 100% | 7% | 0.5% | 4 | 194 MB |  |
| isu4 (2 cpu) | 36.3% / 100% | 17% | 11.3% | 7 | 102 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 59% / nginx 11% / find 10% / kswapd0 1% |
| isu2 | isureal 57% / nginx 2% |
| isu3 | isureal 45% / nginx 2% |
| isu4 | isureal 35% / postgres 13% |

## Diff from previous run

スコア: 178616 → 170564 (-8052)

`f1153ff..b19d232`

```
 webapp/go/offload.go | 43 ++++++++++++++++++++++++++++++++++++-------
 1 file changed, 36 insertions(+), 7 deletions(-)
```
