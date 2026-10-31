# run 20261031-154736-f1153ff

## Summary

| | |
|---|---|
| score | 178616 |
| status | pass |
| ref | presign-next-minute |
| sha | f1153ff4efe1 |
| 所要時間 | 139s |
| 前回比 | +178616 (—) vs 20261031-154148-dfccbc3 |
| ベスト比 | -13200 (-6.9%) vs 20261031-151339-5a230a4 = 191816 |
| note | on b3440cf: sig cache by exp generation + presign next minute |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	135	388	1065	1065
europe	7	21	65	187	514	513
america	6	18	57	164	335	334
```

```
{"id":"24550","status":"done","portalStatus":"finished","passed":true,"score":178616,"raw":178616,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:48:24.877Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:48:27.865Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:48:30.865Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:48:33.875Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T06:48:34.196Z level=INFO msg="user moved" user=sonoa5169 from=asia to=america
time=2026-10-31T06:48:35.219Z level=INFO msg="user moved" user=sonoa5169 from=america to=asia
time=2026-10-31T06:48:35.470Z level=INFO msg="user moved" user=soltavin_5219 from=asia to=europe
time=2026-10-31T06:48:36.585Z level=INFO msg="user moved" user=soltavin_5219 from=europe to=asia
time=2026-10-31T06:48:36.865Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T06:48:39.865Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T06:48:42.875Z level=INFO msg="asia 3 日目: 135 人"
time=2026-10-31T06:48:43.593Z level=INFO msg="user moved" user=sonoa5169 from=asia to=america
time=2026-10-31T06:48:44.167Z level=INFO msg="user moved" user=soltavin_5219 from=asia to=europe
time=2026-10-31T06:48:44.197Z level=INFO msg="user moved" user=naivy.5675 from=asia to=america
time=2026-10-31T06:48:44.607Z level=INFO msg="user moved" user=numu_5801 from=asia to=america
time=2026-10-31T06:48:44.643Z level=INFO msg="user moved" user=soorma.5815 from=asia to=europe
time=2026-10-31T06:48:45.110Z level=INFO msg="user moved" user=naivy.5675 from=america to=europe
time=2026-10-31T06:48:45.247Z level=INFO msg="user moved" user=numu_5801 from=america to=asia
time=2026-10-31T06:48:45.429Z level=INFO msg="user moved" user=soorma.5815 from=europe to=asia
time=2026-10-31T06:48:45.865Z level=INFO msg="europe 3 日目: 65 人"
time=2026-10-31T06:48:45.928Z level=INFO msg="user moved" user=maka_6015 from=asia to=europe
time=2026-10-31T06:48:48.865Z level=INFO msg="america 3 日目: 57 人"
time=2026-10-31T06:48:51.877Z level=INFO msg="asia 4 日目: 388 人"
time=2026-10-31T06:48:52.693Z level=INFO msg="user moved" user=soorma.5815 from=asia to=europe
time=2026-10-31T06:48:52.719Z level=INFO msg="user moved" user=numu_5801 from=asia to=america
time=2026-10-31T06:48:52.948Z level=INFO msg="user moved" user=temu-6415 from=europe to=america
time=2026-10-31T06:48:54.957Z level=INFO msg="europe 4 日目: 187 人"
time=2026-10-31T06:48:56.218Z level=INFO msg="user moved" user=ulel_7933 from=asia to=america
time=2026-10-31T06:48:56.653Z level=INFO msg="user moved" user=naivy.5675 from=europe to=asia
time=2026-10-31T06:48:57.117Z level=INFO msg="user moved" user=ulel_7933 from=america to=asia
time=2026-10-31T06:48:57.291Z level=INFO msg="user moved" user=soltavin_5219 from=europe to=asia
time=2026-10-31T06:48:58.190Z level=INFO msg="america 4 日目: 164 人"
time=2026-10-31T06:48:58.805Z level=INFO msg="user moved" user=daneva_7549 from=asia to=america
time=2026-10-31T06:49:00.451Z level=INFO msg="user moved" user=zoehidan.9072 from=europe to=america
time=2026-10-31T06:49:00.884Z level=INFO msg="asia 5 日目: 1065 人"
time=2026-10-31T06:49:01.160Z level=INFO msg="user moved" user=sonoa5169 from=america to=asia
time=2026-10-31T06:49:04.318Z level=INFO msg="europe 5 日目: 514 人"
time=2026-10-31T06:49:09.234Z level=INFO msg="america 5 日目: 335 人"
time=2026-10-31T06:49:18.232Z level=INFO msg="america 最終日（グランプリ）: 334 人"
time=2026-10-31T06:49:18.232Z level=INFO msg="europe 最終日（グランプリ）: 513 人"
time=2026-10-31T06:49:18.233Z level=INFO msg="asia 最終日（グランプリ）: 1065 人"
time=2026-10-31T06:49:24.996Z level=INFO msg="グランプリ: アジア 419 人, ヨーロッパ 141 人, アメリカ 108 人"
{
  "pass": true,
  "score": 178616,
  "messages": [
    "mudanho-14794: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "madanni_14003: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "semarto_11805: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "haka.13634: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "orwa-13708: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "raolu_huojian_28: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sasa-6136: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "holiatsu_8344: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "bennohe-6980: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "miatorro_12993: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "anivyri.10508: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sosumax12518: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "reheke_6827: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ronoaleo10072: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ivydan-8414: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "fuvinmu10630: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rashise_12291: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "xiangrikui_shousi_1186: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
48.5%	1681	0	0	GET	/api/notifications	17497.698	10.409	17.380	17.504	173
8.6%	5708	87	0	POST	/api/friend-requests	3091.662	0.542	2.344	3.295	114
8.3%	2552	125	0	POST	/api/posts	2983.522	1.169	5.532	64.812	620
6.5%	17518	374	0	GET	^/api/images/[^/]+	2329.562	0.133	1.278	3.664	121145
4.1%	11621	0	0	GET	^/api/feeds/[^/]+	1490.270	0.128	1.147	2.140	3294
3.5%	4128	0	0	GET	/api/memories	1254.866	0.304	3.733	8.229	12602
3.3%	6173	603	0	POST	/api/auth/login	1195.365	0.194	2.014	4.466	222
3.0%	2608	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1077.661	0.413	2.253	3.134	116
1.6%	708	124	0	PUT	/api/me/profile-image	561.259	0.793	9.934	14.412	3
1.3%	543	11	0	POST	^/api/posts/[^/]+/realmojis$	455.394	0.839	3.408	3.875	462
1.1%	645	8	0	POST	^/api/posts/[^/]+/comments$	408.050	0.633	2.954	3.618	336
1.1%	2726	0	0	GET	/api/users	392.201	0.144	0.747	2.140	231
1.1%	3731	0	0	GET	/api/posting-window/current	380.139	0.102	0.491	2.140	237
1.0%	1086	1	0	POST	/api/auth/register	377.052	0.347	1.097	2.006	227
0.9%	706	0	0	PUT	^/api/follows/[^/]+$	334.340	0.474	2.402	3.396	0
0.9%	350	0	0	GET	/api/me/export	324.540	0.927	9.129	9.585	1998834
0.7%	2058	30	0	GET	^/api/posts/me/current$	263.857	0.128	0.511	2.140	642
0.6%	523	0	0	PATCH	/api/me	224.085	0.428	2.461	2.939	275
0.6%	1840	46	0	GET	^/api/posts/[^/]+$	202.422	0.110	1.809	2.140	603
0.6%	925	0	0	POST	/api/auth/logout	199.308	0.215	2.135	2.445	0
0.4%	595	3	0	GET	/api/me	160.891	0.270	1.127	1.263	301
0.4%	1541	0	0	GET	^/api/relationships/[^/]+	159.436	0.103	0.813	2.140	1111
0.4%	217	3	0	PATCH	^/api/posts/[^/]+$	156.518	0.721	3.055	5.430	636
0.3%	846	0	0	GET	^/api/memories/[^/]+	99.241	0.117	0.997	1.884	679
0.3%	96	0	0	DELETE	^/api/posts/[^/]+$	97.494	1.016	3.058	3.058	0
0.2%	906	0	0	GET	^/api/users/[^/]+$	88.516	0.098	0.508	2.140	193
0.2%	696	0	0	GET	^/api/posts/[^/]+/comments$	68.484	0.098	0.589	0.995	2310
0.2%	756	0	0	GET	^/api/posts/[^/]+/realmojis$	66.700	0.088	0.647	0.995	1530
0.2%	102	0	0	DELETE	^/api/realmojis/[^/]+$	65.928	0.646	3.016	3.399	0
0.1%	119	0	0	DELETE	^/api/comments/[^/]+$	42.705	0.359	2.511	2.936	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.6%	728	1	0	GET	/api/notifications	5148.676	7.072	11.892	11.900	161
11.2%	1047	114	0	POST	/api/posts	1454.476	1.389	6.617	8.192	580
10.5%	2470	33	0	POST	/api/friend-requests	1366.952	0.553	2.119	3.171	114
5.9%	1589	0	0	POST	/api/auth/login	769.702	0.484	2.521	2.942	246
5.3%	4827	40	0	GET	^/api/images/[^/]+	694.951	0.144	1.118	6.030	159387
4.7%	4133	0	0	GET	^/api/feeds/[^/]+	604.762	0.146	1.448	2.657	3514
4.0%	512	0	0	POST	/api/auth/register	518.078	1.012	2.797	2.836	229
2.3%	299	0	0	PUT	/api/me/profile-image	304.203	1.017	3.287	4.427	0
2.1%	361	0	0	PUT	^/api/follows/[^/]+$	269.157	0.746	2.771	3.228	0
1.9%	246	0	0	PATCH	/api/me	246.865	1.004	2.480	2.688	278
1.7%	518	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	222.934	0.430	2.187	2.744	116
1.5%	1204	0	0	GET	/api/memories	188.546	0.157	1.704	2.984	14451
1.4%	1695	0	0	GET	/api/posting-window/current	178.279	0.105	0.595	1.688	238
1.3%	188	7	0	POST	^/api/posts/[^/]+/realmojis$	168.374	0.896	2.952	3.118	466
1.0%	1184	0	0	GET	/api/users	131.444	0.111	0.687	1.349	236
1.0%	763	82	0	GET	^/api/posts/me/current$	127.748	0.167	1.728	1.893	584
1.0%	223	7	0	POST	^/api/posts/[^/]+/comments$	123.476	0.554	2.870	4.594	337
0.6%	241	0	0	POST	/api/auth/logout	72.817	0.302	1.938	2.720	0
0.5%	524	0	0	GET	^/api/relationships/[^/]+	70.872	0.135	1.005	1.451	594
0.5%	543	30	0	GET	^/api/posts/[^/]+$	60.469	0.111	0.766	1.022	593
0.4%	44	0	0	DELETE	^/api/posts/[^/]+$	47.967	1.090	2.994	2.994	0
0.3%	71	12	0	PATCH	^/api/posts/[^/]+$	32.858	0.463	2.638	2.638	547
0.2%	305	2	0	GET	/api/me	32.466	0.106	0.427	1.046	302
0.2%	251	0	0	GET	^/api/users/[^/]+$	31.488	0.125	0.771	0.832	206
0.2%	223	0	0	GET	^/api/posts/[^/]+/realmojis$	25.856	0.116	0.687	0.832	1315
0.2%	224	0	0	GET	^/api/posts/[^/]+/comments$	25.273	0.113	0.723	1.221	2120
0.2%	207	0	0	GET	^/api/memories/[^/]+	21.691	0.105	0.850	1.554	773
0.1%	42	0	0	DELETE	^/api/comments/[^/]+$	19.429	0.463	3.520	3.520	0
0.1%	76	0	0	GET	/api/me/export	14.376	0.189	4.048	4.048	2358838
0.1%	25	0	0	DELETE	^/api/realmojis/[^/]+$	11.771	0.471	1.916	1.916	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.2%	483	0	0	GET	/api/notifications	3548.379	7.347	10.923	10.937	159
11.8%	683	89	0	POST	/api/posts	991.228	1.451	6.611	7.010	563
7.7%	1572	11	0	POST	/api/friend-requests	649.868	0.413	2.192	2.579	114
6.5%	1305	0	0	POST	/api/auth/login	548.253	0.420	2.074	2.966	250
6.3%	4279	17	0	GET	^/api/images/[^/]+	530.446	0.124	1.154	3.068	128783
4.4%	2859	0	0	GET	^/api/feeds/[^/]+	372.640	0.130	1.050	1.675	3537
2.9%	330	0	0	POST	/api/auth/register	244.835	0.742	1.824	1.871	230
2.2%	412	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	182.862	0.444	2.228	2.401	116
1.8%	1031	0	0	GET	/api/memories	151.715	0.147	1.630	5.215	18358
1.8%	172	0	0	PUT	/api/me/profile-image	150.257	0.874	2.478	2.573	0
1.4%	1082	0	0	GET	/api/posting-window/current	114.601	0.106	0.631	1.279	239
1.3%	251	0	0	PUT	^/api/follows/[^/]+$	108.342	0.432	2.126	2.807	0
1.2%	179	0	0	PATCH	/api/me	104.608	0.584	1.776	1.945	276
1.2%	107	0	0	POST	^/api/posts/[^/]+/realmojis$	101.377	0.947	3.947	4.014	458
1.0%	800	0	0	GET	/api/users	85.651	0.107	0.786	1.360	229
1.0%	502	13	0	GET	^/api/posts/me/current$	85.233	0.170	1.238	1.390	631
0.8%	199	0	0	POST	/api/auth/logout	67.018	0.337	2.109	2.886	0
0.8%	124	3	0	POST	^/api/posts/[^/]+/comments$	65.243	0.526	2.572	3.223	329
0.6%	443	26	0	GET	^/api/posts/[^/]+$	53.322	0.120	0.867	1.230	586
0.6%	361	0	0	GET	^/api/relationships/[^/]+	47.447	0.131	0.949	1.240	840
0.4%	51	5	0	PATCH	^/api/posts/[^/]+$	35.768	0.701	2.546	2.546	585
0.3%	204	0	0	GET	^/api/users/[^/]+$	25.083	0.123	0.735	1.389	197
0.3%	19	0	0	DELETE	^/api/posts/[^/]+$	22.713	1.195	2.373	2.373	0
0.3%	158	0	0	GET	^/api/posts/[^/]+/realmojis$	21.299	0.135	1.279	1.586	1365
0.2%	68	0	0	GET	/api/me/export	18.994	0.279	6.540	6.540	2617553
0.2%	178	1	0	GET	/api/me	18.480	0.104	0.242	1.489	302
0.2%	175	0	0	GET	^/api/memories/[^/]+	16.924	0.097	0.812	1.326	1012
0.2%	159	0	0	GET	^/api/posts/[^/]+/comments$	14.870	0.094	0.291	0.712	1864
0.2%	15	0	0	DELETE	^/api/realmojis/[^/]+$	14.221	0.948	3.278	3.278	0
0.1%	25	0	0	DELETE	^/api/comments/[^/]+$	9.048	0.362	2.121	2.121	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 14884 クエリ / 50 種類 / 合計 152.174s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	55.3%	355	84.080	236.85	4099.34	1276.7	130.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.0%	3	21.269	7089.79	9544.49	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	10.5%	3	15.945	5314.87	7191.40	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	9.7%	12314	14.781	1.20	10.49	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	5.2%	3	7.960	2653.47	3637.56	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.6%	6	3.897	649.51	2958.97	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.7%	6	1.050	174.93	525.70	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.6%	1	0.910	910.27	910.27	32628.0	1.0	select count(*), sum(length(data))/$1 from image_objects
9	0.4%	6	0.680	113.33	158.40	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.2%	1	0.309	309.14	309.14	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.253	253.43	253.43	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.2%	1	0.235	235.34	235.34	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
13	0.1%	1	0.172	171.97	171.97	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.108	108.06	108.06	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	1	0.082	82.20	82.20	2463.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…

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
Build ID: 6b74704f02e049ca1defe636aa09813615f43fc0
Type: cpu
Time: 2026-10-31 15:48:03 JST
Duration: 60.16s, Total samples = 40.70s (67.65%)
Showing nodes accounting for 29.29s, 71.97% of 40.70s total
Dropped 1110 nodes (cum <= 0.20s)
Showing top 40 nodes out of 278
      flat  flat%   sum%        cum   cum%
     5.18s 12.73% 12.73%      5.18s 12.73%  crypto/internal/fips140/sha256.blockAVX2
     4.81s 11.82% 24.55%      4.81s 11.82%  internal/runtime/syscall/linux.Syscall6
     1.68s  4.13% 28.67%      1.68s  4.13%  runtime.memmove
     1.65s  4.05% 32.73%      2.30s  5.65%  runtime.findObject
     1.63s  4.00% 36.73%      5.79s 14.23%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.57s  3.86% 40.59%      1.57s  3.86%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.48s  3.64% 44.23%      6.24s 15.33%  runtime.scanObject
     1.45s  3.56% 47.79%      1.45s  3.56%  runtime.memclrNoHeapPointers
     1.32s  3.24% 51.03%      2.31s  5.68%  runtime.tryDeferToSpanScan
     1.24s  3.05% 54.08%      1.24s  3.05%  aeshashbody
     0.76s  1.87% 55.95%      0.76s  1.87%  runtime.spanClass.sizeclass (inline)
     0.60s  1.47% 57.42%      2.12s  5.21%  runtime.mapaccess1_faststr
     0.45s  1.11% 58.53%      0.62s  1.52%  example.com/isureal-company/isureal.postLess
     0.43s  1.06% 59.58%      0.43s  1.06%  runtime.nextFreeFast (inline)
     0.40s  0.98% 60.57%      0.40s  0.98%  memeqbody
     0.31s  0.76% 61.33%      5.82s 14.30%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.30s  0.74% 62.06%      0.85s  2.09%  database/sql.convertAssignRows
     0.30s  0.74% 62.80%      0.30s  0.74%  runtime.futex
     0.30s  0.74% 63.54%      2.02s  4.96%  runtime.mapassign_faststr
     0.29s  0.71% 64.25%      0.35s  0.86%  runtime.spanOf (inline)
     0.27s  0.66% 64.91%      5.70s 14.00%  crypto/internal/fips140/sha256.(*Digest).Write
     0.24s  0.59% 65.50%      0.54s  1.33%  runtime.greyobject
     0.22s  0.54% 66.04%      1.15s  2.83%  runtime.mallocgcSmallScanNoHeader
     0.22s  0.54% 66.58%      0.22s  0.54%  sync/atomic.(*Int32).Add
     0.20s  0.49% 67.08%      0.41s  1.01%  compress/flate.(*compressor).deflate
     0.20s  0.49% 67.57%      0.22s  0.54%  runtime.(*mspan).writeHeapBitsSmall
     0.19s  0.47% 68.03%      6.06s 14.89%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.19s  0.47% 68.50%      0.28s  0.69%  internal/runtime/maps.(*Iter).Next
     0.19s  0.47% 68.97%      0.25s  0.61%  runtime.(*mspan).markBitsForIndex (inline)
     0.15s  0.37% 69.34%      6.41s 15.75%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.15s  0.37% 69.71%      4.53s 11.13%  runtime.mallocgc
     0.12s  0.29% 70.00%      1.03s  2.53%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.29% 70.29%      0.79s  1.94%  runtime.mallocgcSmallNoscan
     0.11s  0.27% 70.57%      6.83s 16.78%  example.com/isureal-company/isureal.computeImageSignature
     0.11s  0.27% 70.84%      2.88s  7.08%  example.com/isureal-company/isureal.listMemories
     0.10s  0.25% 71.08%      0.33s  0.81%  internal/stringslite.Index
     0.09s  0.22% 71.30%      0.30s  0.74%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.09s  0.22% 71.52%      2.75s  6.76%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.09s  0.22% 71.74%      0.31s  0.76%  runtime.bulkBarrierPreWrite
     0.09s  0.22% 71.97%      0.72s  1.77%  runtime.scanSpan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 43.1% / 100% | 26% | 2.2% | 5 | 249 MB | run queue が積もっている |
| isu2 (2 cpu) | 35.1% / 100% | 3% | 0.5% | 7 | 245 MB | run queue が積もっている |
| isu3 (2 cpu) | 27.2% / 100% | 5% | 0.7% | 5 | 261 MB | run queue が積もっている |
| isu4 (2 cpu) | 40.7% / 100% | 13% | 16.5% | 9 | 98 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 65% / nginx 13% |
| isu2 | isureal 66% / nginx 2% |
| isu3 | isureal 51% / nginx 2% |
| isu4 | isureal 30% / postgres 15% |

## Diff from previous run

スコア: 0 → 178616 (+178616)

`dfccbc3..f1153ff`

```
 webapp/go/admin_api.go      | 27 +++++++++++++++++++++-----
 webapp/go/export_api.go     | 12 ++++++++++--
 webapp/go/image_api.go      | 23 +++++++++++++----------
 webapp/go/image_file.go     | 33 ++++++++++++++++++--------------
 webapp/go/image_held.go     |  5 ++---
 webapp/go/image_local.go    | 46 +++++++++++++++++++--------------------------
 webapp/go/ip_ban.go         | 11 +++++++++++
 webapp/go/offload.go        | 43 +++++++-----------------------------------
 webapp/go/repl.go           |  1 +
 webapp/go/replica_server.go |  5 ++++-
 10 files changed, 108 insertions(+), 98 deletions(-)
```
