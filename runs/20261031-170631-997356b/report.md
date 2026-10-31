# run 20261031-170631-997356b

## Summary

| | |
|---|---|
| score | 233154 |
| status | pass |
| ref | sign-peers-euus |
| sha | 997356bce53e |
| 所要時間 | 158s |
| 前回比 | +231593 (14836.2%) vs 20261031-165402-07a4cc3 |
| ベスト比 | +1044 (0.4%) vs 20261031-164055-bdf05cd = 232110 **ベスト更新** |
| note | asia1 sign offload peers: europe1/america1 only (no asia2) |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	145	399	670	669
europe	7	21	66	191	530	532
america	6	18	56	165	445	442
```

```
{"id":"32275","status":"done","portalStatus":"finished","passed":true,"score":233154,"raw":233154,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:07:32.427Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:07:35.416Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:07:38.416Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:07:41.461Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T08:07:44.416Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T08:07:47.420Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T08:07:50.442Z level=INFO msg="asia 3 日目: 145 人"
time=2026-10-31T08:07:52.795Z level=INFO msg="user moved" user=rakoivy.5862 from=asia to=america
time=2026-10-31T08:07:52.837Z level=INFO msg="user moved" user=zoesu5700 from=asia to=europe
time=2026-10-31T08:07:53.419Z level=INFO msg="europe 3 日目: 66 人"
time=2026-10-31T08:07:53.471Z level=INFO msg="user moved" user=rakoivy.5862 from=america to=asia
time=2026-10-31T08:07:54.725Z level=INFO msg="user moved" user=oranivy.6101 from=asia to=europe
time=2026-10-31T08:07:56.449Z level=INFO msg="america 3 日目: 56 人"
time=2026-10-31T08:07:56.470Z level=INFO msg="user moved" user=sasashi-5977 from=asia to=america
time=2026-10-31T08:07:59.428Z level=INFO msg="asia 4 日目: 399 人"
time=2026-10-31T08:08:02.417Z level=INFO msg="europe 4 日目: 191 人"
time=2026-10-31T08:08:03.608Z level=INFO msg="user moved" user=keheme6904 from=asia to=europe
time=2026-10-31T08:08:04.701Z level=INFO msg="user moved" user=zoesu5700 from=europe to=america
time=2026-10-31T08:08:05.432Z level=INFO msg="america 4 日目: 165 人"
time=2026-10-31T08:08:05.534Z level=INFO msg="user moved" user=rakoivy.5862 from=asia to=america
time=2026-10-31T08:08:07.516Z level=INFO msg="user moved" user=keteke-8394 from=asia to=america
time=2026-10-31T08:08:08.070Z level=INFO msg="user moved" user=sona-8675 from=asia to=europe
time=2026-10-31T08:08:08.389Z level=INFO msg="user moved" user=miaevaso.9117 from=europe to=america
time=2026-10-31T08:08:08.539Z level=INFO msg="asia 5 日目: 670 人"
time=2026-10-31T08:08:11.425Z level=INFO msg="europe 5 日目: 530 人"
time=2026-10-31T08:08:12.583Z level=INFO msg="user moved" user=shiliavin-9815 from=america to=europe
time=2026-10-31T08:08:14.475Z level=INFO msg="america 5 日目: 445 人"
time=2026-10-31T08:08:14.715Z level=INFO msg="user moved" user=sasashi-5977 from=america to=europe
time=2026-10-31T08:08:14.719Z level=INFO msg="user moved" user=isyosa10435 from=america to=europe
time=2026-10-31T08:08:23.478Z level=INFO msg="america 最終日（グランプリ）: 442 人"
time=2026-10-31T08:08:23.531Z level=INFO msg="asia 最終日（グランプリ）: 669 人"
time=2026-10-31T08:08:23.637Z level=INFO msg="europe 最終日（グランプリ）: 532 人"
time=2026-10-31T08:08:32.532Z level=INFO msg="グランプリ: アジア 557 人, ヨーロッパ 471 人, アメリカ 399 人"
{
  "pass": true,
  "score": 233154,
  "messages": [
    "ivymi.5102: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "torse_12721: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "muni_15571: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "misa.16072: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "raisvin-8510: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.38.108:45238-\u003e52.69.202.109:8080: write tcp 10.27.38.108:45238-\u003e52.69.202.109:8080: use of closed network connection",
    "sakura_takahashi_269: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "chimu-8273: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.38.108:44584-\u003e52.69.202.109:8080: write tcp 10.27.38.108:44584-\u003e52.69.202.109:8080: use of closed network connection",
    "marhaul7217: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.38.108:45084-\u003e52.69.202.109:8080: write tcp 10.27.38.108:45084-\u003e52.69.202.109:8080: use of closed network connection",
    "kiyu-10468: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "aoi_yamamoto_3334: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "nomiko-10746: ログインし直せません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "toha.13163: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noasemo_8986: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rarima_15588: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rebenwa.11263: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "miasu.6963: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "miake.8666: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.38.108:45654-\u003e52.69.202.109:8080: write tcp 10.27.38.108:45654-\u003e52.69.202.109:8080: use of closed network connection"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
49.5%	1319	1	0	GET	/api/notifications	13231.053	10.031	15.021	15.022	171
11.3%	9890	0	0	GET	^/api/feeds/[^/]+	3020.323	0.305	1.432	2.645	4976
8.6%	3165	0	0	GET	/api/memories	2295.097	0.725	5.345	12.114	10823
5.3%	2691	0	0	GET	/api/posting-window/current	1421.213	0.528	1.480	2.016	237
3.7%	1950	23	0	POST	/api/posts	996.386	0.511	2.297	7.797	646
2.8%	226	0	0	GET	/api/me/export	744.140	3.293	21.067	21.502	1813294
2.7%	1929	0	0	GET	/api/users	732.327	0.380	1.500	1.964	240
2.6%	1788	2	0	GET	^/api/posts/me/current$	688.746	0.385	1.441	1.705	653
1.7%	509	129	0	PUT	/api/me/profile-image	467.359	0.918	10.893	13.062	4
1.6%	12020	214	0	GET	^/api/images/[^/]+	439.611	0.037	0.689	1.897	160329
1.5%	772	0	0	GET	^/api/posts/[^/]+/realmojis$	406.694	0.527	2.996	4.988	2676
1.3%	1349	0	0	GET	^/api/relationships/[^/]+	351.647	0.261	1.438	3.121	1601
1.1%	4044	47	0	POST	/api/friend-requests	294.460	0.073	0.699	1.141	114
1.1%	1620	53	0	GET	^/api/posts/[^/]+$	292.003	0.180	1.011	1.242	614
0.8%	394	3	0	GET	/api/me	227.092	0.576	1.546	1.799	299
0.8%	650	1	0	POST	^/api/posts/[^/]+/realmojis$	201.048	0.309	1.516	1.742	488
0.7%	3578	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	189.964	0.053	0.668	1.355	116
0.6%	515	0	0	GET	^/api/memories/[^/]+	163.523	0.318	1.604	2.298	654
0.5%	628	0	0	GET	^/api/posts/[^/]+/comments$	133.204	0.212	1.124	1.415	2987
0.5%	849	0	0	GET	^/api/users/[^/]+$	127.557	0.150	0.940	1.447	204
0.3%	4006	542	0	POST	/api/auth/login	74.888	0.019	0.435	0.955	214
0.2%	624	0	0	POST	^/api/posts/[^/]+/comments$	58.394	0.094	0.774	1.401	345
0.2%	86	0	0	DELETE	^/api/posts/[^/]+$	48.080	0.559	1.098	1.098	0
0.1%	191	1	0	PATCH	^/api/posts/[^/]+$	28.815	0.151	0.748	0.850	647
0.1%	1	0	0	POST	/api/initialize	22.838	22.838	22.838	22.838	90
0.1%	592	0	0	POST	/api/auth/logout	21.539	0.036	0.631	0.705	0
0.1%	87	0	0	DELETE	^/api/realmojis/[^/]+$	20.559	0.236	1.003	1.003	0
0.0%	340	0	0	PATCH	/api/me	8.111	0.024	0.523	0.637	276
0.0%	443	0	0	PUT	^/api/follows/[^/]+$	6.660	0.015	0.375	0.569	0
0.0%	30	0	0	PUT	^/api/blocks/[^/]+$	4.863	0.162	0.759	0.759	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
51.6%	845	1	0	GET	/api/notifications	6645.271	7.864	11.969	11.974	172
7.9%	8691	0	0	GET	^/api/feeds/[^/]+	1022.483	0.118	1.103	1.936	4018
6.2%	7078	42	0	GET	^/api/images/[^/]+	795.039	0.112	1.160	2.892	186030
6.1%	3083	28	0	POST	/api/friend-requests	784.488	0.254	0.988	1.475	114
5.8%	1343	20	0	POST	/api/posts	750.699	0.559	6.666	7.841	645
2.9%	1996	0	0	POST	/api/auth/login	374.831	0.188	0.944	1.543	247
2.1%	1838	0	0	GET	/api/memories	272.684	0.148	1.527	2.073	12453
1.9%	1916	0	0	GET	/api/posting-window/current	242.573	0.127	0.707	1.756	239
1.5%	1488	0	0	GET	/api/users	196.248	0.132	0.689	1.790	238
1.5%	1188	0	0	GET	^/api/posts/me/current$	191.125	0.161	1.411	1.643	654
1.4%	527	5	0	POST	^/api/posts/[^/]+/realmojis$	184.852	0.351	1.135	1.584	473
1.3%	526	0	0	POST	/api/auth/register	172.553	0.328	1.064	1.501	229
1.2%	307	0	0	PUT	/api/me/profile-image	159.569	0.520	1.258	1.647	0
1.2%	725	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	157.978	0.218	0.994	1.447	116
1.0%	1411	30	0	GET	^/api/posts/[^/]+$	133.769	0.095	0.367	1.548	625
1.0%	1057	0	0	GET	^/api/relationships/[^/]+	126.867	0.120	1.035	1.608	536
0.7%	348	0	0	PUT	^/api/follows/[^/]+$	93.861	0.270	0.975	1.048	0
0.6%	266	0	0	PATCH	/api/me	80.258	0.302	0.994	1.017	274
0.5%	717	0	0	GET	^/api/users/[^/]+$	64.569	0.090	0.442	1.235	213
0.5%	691	0	0	GET	^/api/posts/[^/]+/realmojis$	63.259	0.092	0.531	1.404	4300
0.5%	344	0	0	POST	^/api/posts/[^/]+/comments$	62.170	0.181	0.937	1.302	347
0.4%	298	0	0	POST	/api/auth/logout	50.401	0.169	0.952	1.291	0
0.4%	492	0	0	GET	^/api/posts/[^/]+/comments$	48.124	0.098	0.551	1.433	4424
0.3%	128	0	0	PATCH	^/api/posts/[^/]+$	42.316	0.331	1.580	1.638	651
0.3%	60	0	0	DELETE	^/api/posts/[^/]+$	41.538	0.692	2.100	2.100	0
0.3%	316	2	0	GET	/api/me	36.052	0.114	0.706	0.820	302
0.2%	279	0	0	GET	^/api/memories/[^/]+	28.338	0.102	1.128	1.666	898
0.2%	90	0	0	DELETE	^/api/realmojis/[^/]+$	27.775	0.309	0.868	0.868	0
0.1%	117	0	0	GET	/api/me/export	15.480	0.132	0.762	1.250	2114992
0.1%	56	0	0	DELETE	^/api/comments/[^/]+$	6.956	0.124	0.241	0.241	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
47.7%	711	0	0	GET	/api/notifications	4590.670	6.457	9.026	9.043	175
8.6%	5707	0	0	GET	^/api/feeds/[^/]+	828.000	0.145	1.359	2.262	3972
7.9%	1131	25	0	POST	/api/posts	762.868	0.675	7.201	7.899	642
7.6%	6062	29	0	GET	^/api/images/[^/]+	728.970	0.120	1.054	2.781	173252
4.9%	2582	26	0	POST	/api/friend-requests	475.619	0.184	0.977	1.488	114
3.6%	1716	0	0	POST	/api/auth/login	342.210	0.199	0.895	1.591	250
2.4%	1606	0	0	GET	/api/memories	230.983	0.144	1.227	2.482	14779
1.9%	1600	0	0	GET	/api/posting-window/current	182.592	0.114	0.974	1.725	240
1.9%	1252	0	0	GET	/api/users	178.162	0.142	1.120	1.686	239
1.8%	1057	0	0	GET	^/api/posts/me/current$	174.982	0.166	1.573	2.252	657
1.5%	440	0	0	POST	/api/auth/register	144.585	0.329	1.105	1.289	230
1.3%	348	20	0	POST	^/api/posts/[^/]+/realmojis$	129.293	0.372	1.187	1.369	453
1.0%	560	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	98.649	0.176	0.995	1.607	116
1.0%	698	0	0	GET	^/api/relationships/[^/]+	96.554	0.138	1.302	1.602	648
1.0%	805	46	0	GET	^/api/posts/[^/]+$	93.051	0.116	0.774	2.254	608
0.9%	246	0	0	PUT	/api/me/profile-image	81.996	0.333	1.249	1.742	0
0.7%	320	4	0	POST	^/api/posts/[^/]+/comments$	63.839	0.199	0.902	1.681	356
0.6%	450	0	0	GET	^/api/users/[^/]+$	53.130	0.118	0.812	1.606	215
0.5%	281	0	0	POST	/api/auth/logout	52.901	0.188	0.903	1.209	0
0.5%	269	0	0	PUT	^/api/follows/[^/]+$	47.989	0.178	0.880	1.531	0
0.4%	211	0	0	PATCH	/api/me	38.646	0.183	0.854	1.560	277
0.4%	337	0	0	GET	^/api/posts/[^/]+/realmojis$	37.756	0.112	0.702	1.098	5171
0.4%	299	2	0	GET	^/api/posts/[^/]+/comments$	34.364	0.115	1.003	1.325	4823
0.3%	253	1	0	GET	/api/me	29.109	0.115	0.345	0.620	304
0.3%	40	0	0	DELETE	^/api/posts/[^/]+$	28.648	0.716	2.360	2.360	0
0.3%	91	0	0	PATCH	^/api/posts/[^/]+$	26.579	0.292	1.561	1.561	646
0.3%	225	0	0	GET	^/api/memories/[^/]+	25.060	0.111	1.109	1.427	1052
0.2%	106	0	0	GET	/api/me/export	21.372	0.202	2.527	3.069	2865833
0.2%	48	0	0	DELETE	^/api/realmojis/[^/]+$	17.008	0.354	0.814	0.814	0
0.1%	47	0	0	DELETE	^/api/comments/[^/]+$	6.851	0.146	0.419	0.419	0
…（上位 30 / 全 36 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 11158 クエリ / 48 種類 / 合計 140.073s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	44.2%	233	61.875	265.56	861.21	1161.6	118.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	18.5%	4	25.896	6474.04	9541.24	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	13.5%	4	18.967	4741.78	7215.37	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	9.6%	7463	13.493	1.81	9.05	10.0	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.4%	4	8.931	2232.79	3423.12	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	4.3%	7	6.086	869.43	3384.29	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	1.1%	1	1.550	1549.66	1549.66	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.4%	7	0.614	87.77	161.50	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.4%	7	0.603	86.12	158.35	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
10	0.3%	1	0.375	375.21	375.21	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.347	347.42	347.42	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.2%	1	0.233	233.06	233.06	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.161	160.52	160.52	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.142	141.92	141.92	2433.0	0.0	DELETE FROM image_objects WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	1	0.129	129.15	129.15	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: c881fdd0fd0f4c6a603496621a4fb609aee3af46
Type: cpu
Time: 2026-10-31 17:07:05 JST
Duration: 60.01s, Total samples = 26.10s (43.49%)
Showing nodes accounting for 20.22s, 77.47% of 26.10s total
Dropped 826 nodes (cum <= 0.13s)
Showing top 40 nodes out of 224
      flat  flat%   sum%        cum   cum%
     4.66s 17.85% 17.85%      4.66s 17.85%  crypto/internal/fips140/sha256.blockAVX2
     2.40s  9.20% 27.05%      2.40s  9.20%  internal/runtime/syscall/linux.Syscall6
     1.41s  5.40% 32.45%      5.97s 22.87%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.35s  5.17% 37.62%      1.35s  5.17%  runtime.memmove
     1.32s  5.06% 42.68%      1.32s  5.06%  aeshashbody
     1.31s  5.02% 47.70%      1.31s  5.02%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.64s  2.45% 50.15%      0.64s  2.45%  runtime.memclrNoHeapPointers
     0.49s  1.88% 52.03%      0.69s  2.64%  example.com/isureal-company/isureal.postLess
     0.49s  1.88% 53.91%      0.66s  2.53%  runtime.findObject
     0.49s  1.88% 55.79%      0.73s  2.80%  runtime.tryDeferToSpanScan
     0.48s  1.84% 57.62%         2s  7.66%  runtime.scanObject
     0.44s  1.69% 59.31%      0.70s  2.68%  example.com/isureal-company/isureal/jpegcheck.(*decoder).skipAC
     0.44s  1.69% 61.00%      1.97s  7.55%  runtime.mapaccess1_faststr
     0.30s  1.15% 62.15%      5.27s 20.19%  crypto/internal/fips140/sha256.(*Digest).Write
     0.27s  1.03% 63.18%      0.27s  1.03%  memeqbody
     0.26s     1% 64.18%      0.26s     1%  sync/atomic.(*Int32).Add (inline)
     0.25s  0.96% 65.13%      5.36s 20.54%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.25s  0.96% 66.09%      0.25s  0.96%  example.com/isureal-company/isureal/jpegcheck.(*decoder).prefetch
     0.25s  0.96% 67.05%      0.25s  0.96%  runtime.nextFreeFast (inline)
     0.22s  0.84% 67.89%      0.22s  0.84%  runtime.futex
     0.21s   0.8% 68.70%      0.36s  1.38%  database/sql.convertAssignRows
     0.19s  0.73% 69.43%      0.98s  3.75%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.19s  0.73% 70.15%      0.74s  2.84%  runtime.mallocgcSmallScanNoHeader
     0.16s  0.61% 70.77%      0.19s  0.73%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.16s  0.61% 71.38%      0.17s  0.65%  example.com/isureal-company/isureal.rmLess
     0.14s  0.54% 71.92%      0.24s  0.92%  runtime.greyobject
     0.13s   0.5% 72.41%      5.80s 22.22%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.13s   0.5% 72.91%      0.17s  0.65%  runtime.(*mspan).writeHeapBitsSmall
     0.13s   0.5% 73.41%      0.40s  1.53%  runtime.mapaccess2_faststr
     0.12s  0.46% 73.87%      0.18s  0.69%  internal/runtime/maps.(*Iter).Next
     0.12s  0.46% 74.33%      1.50s  5.75%  runtime.mapassign_faststr
     0.11s  0.42% 74.75%      5.54s 21.23%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.11s  0.42% 75.17%      1.45s  5.56%  runtime.mallocgc
     0.10s  0.38% 75.56%      0.24s  0.92%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.10s  0.38% 75.94%      4.76s 18.24%  crypto/internal/fips140/sha256.block
     0.10s  0.38% 76.32%      1.04s  3.98%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.08s  0.31% 76.63%      6.22s 23.83%  example.com/isureal-company/isureal.computeImageSignature
     0.08s  0.31% 76.93%      1.84s  7.05%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.07s  0.27% 77.20%      0.43s  1.65%  database/sql.(*Rows).scanLocked
     0.07s  0.27% 77.47%      3.76s 14.41%  example.com/isureal-company/isureal.loadBig
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 31.6% / 100% | 9% | 0.7% | 6 | 178 MB | run queue が積もっている |
| isu2 (2 cpu) | 25.0% / 100% | 7% | 0.8% | 6 | 173 MB | run queue が積もっている |
| isu3 (2 cpu) | 24.5% / 100% | 8% | 0.8% | 5 | 145 MB | run queue が積もっている |
| isu4 (2 cpu) | 64.3% / 100% | 55% | 3.4% | 6 | 103 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 49% / nginx 10% |
| isu2 | isureal 45% / nginx 3% |
| isu3 | isureal 45% / nginx 3% |
| isu4 | isureal 103% / postgres 16% |

## Diff from previous run

スコア: 1561 → 233154 (+231593)

`07a4cc3..997356b`

```
 conf/isu1/home/isucon/env.sh |  2 +-
 webapp/go/replica_server.go  |  1 -
 webapp/go/server.go          |  1 -
 webapp/go/slow.go            | 36 ------------------------------------
 4 files changed, 1 insertion(+), 39 deletions(-)
```
