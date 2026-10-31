# run 20261031-172947-21d10c4

## Summary

| | |
|---|---|
| score | 205640 |
| status | pass |
| ref | combined |
| sha | 21d10c4d020a |
| 所要時間 | 185s |
| 前回比 | -17794 (-8.0%) vs 20261031-172124-4c078d8 |
| ベスト比 | -26470 (-11.4%) vs 20261031-164055-bdf05cd = 232110 |
| note | combined: sse read timeout 20s + asia2 sign offload to EU/US (64 slots) |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	146	421	1035	1038
europe	7	21	63	185	518	517
america	6	20	61	170	461	460
```

```
{"id":"34740","status":"done","portalStatus":"finished","passed":true,"score":205640,"raw":205640,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:31:04.027Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:31:07.018Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:31:10.017Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:31:13.031Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T08:31:14.061Z level=INFO msg="user moved" user=zoete5207 from=asia to=america
time=2026-10-31T08:31:14.804Z level=INFO msg="user moved" user=zoete5207 from=america to=asia
time=2026-10-31T08:31:16.018Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T08:31:19.017Z level=INFO msg="america 2 日目: 20 人"
time=2026-10-31T08:31:22.032Z level=INFO msg="asia 3 日目: 146 人"
time=2026-10-31T08:31:23.739Z level=INFO msg="user moved" user=zoete5207 from=asia to=america
time=2026-10-31T08:31:25.018Z level=INFO msg="europe 3 日目: 63 人"
time=2026-10-31T08:31:25.922Z level=INFO msg="user moved" user=danivysa_5817 from=asia to=europe
time=2026-10-31T08:31:28.017Z level=INFO msg="america 3 日目: 61 人"
time=2026-10-31T08:31:31.029Z level=INFO msg="asia 4 日目: 421 人"
time=2026-10-31T08:31:33.560Z level=INFO msg="user moved" user=kemuma-7031 from=america to=asia
time=2026-10-31T08:31:34.035Z level=INFO msg="europe 4 日目: 185 人"
time=2026-10-31T08:31:35.995Z level=INFO msg="user moved" user=kemuma-7031 from=asia to=europe
time=2026-10-31T08:31:36.412Z level=INFO msg="user moved" user=danku.8636 from=asia to=europe
time=2026-10-31T08:31:37.021Z level=INFO msg="america 4 日目: 170 人"
time=2026-10-31T08:31:37.224Z level=INFO msg="user moved" user=ramaxka.7110 from=asia to=europe
time=2026-10-31T08:31:39.899Z level=INFO msg="user moved" user=riyuan_8737 from=asia to=america
time=2026-10-31T08:31:39.975Z level=INFO msg="user moved" user=annoaru_7517 from=asia to=europe
time=2026-10-31T08:31:40.035Z level=INFO msg="asia 5 日目: 1035 人"
time=2026-10-31T08:31:40.437Z level=INFO msg="user moved" user=zoete5207 from=america to=asia
time=2026-10-31T08:31:40.467Z level=INFO msg="user moved" user=kota9011 from=europe to=asia
time=2026-10-31T08:31:40.875Z level=INFO msg="user moved" user=kakoan_9328 from=europe to=asia
time=2026-10-31T08:31:41.683Z level=INFO msg="user moved" user=danivysa_5817 from=europe to=asia
time=2026-10-31T08:31:43.051Z level=INFO msg="europe 5 日目: 518 人"
time=2026-10-31T08:31:46.993Z level=INFO msg="america 5 日目: 461 人"
time=2026-10-31T08:31:55.566Z level=INFO msg="asia 最終日（グランプリ）: 1038 人"
time=2026-10-31T08:31:55.896Z level=INFO msg="europe 最終日（グランプリ）: 517 人"
time=2026-10-31T08:31:56.015Z level=INFO msg="america 最終日（グランプリ）: 460 人"
time=2026-10-31T08:32:04.144Z level=INFO msg="グランプリ: アジア 464 人, ヨーロッパ 158 人, アメリカ 181 人"
{
  "pass": true,
  "score": 205640,
  "messages": [
    "leoleote5036: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "home-18404: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nabenhe-18636: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "miaor18798: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "niyo.17865: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "liasol-17160: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kihe-18341: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rohe.18169: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ishi-17092: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "orsol_18539: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "leotorchi19015: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "masozoe17298: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mizoe18107: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "manulia-10004: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sunito_18414: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mabenan-5058: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noan12680: ログインし直せません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "sasol.11778: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.40.144:60226-\u003e52.69.202.109:8080: write tcp 10.27.40.144:60226-\u003e52.69.202.109:8080: use of closed network connection",
    "swieoganeun_mogyoktang_4357: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "lei_zhao_1922: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
65.2%	2731	1	0	GET	/api/notifications	36905.380	13.514	20.002	20.004	107
5.3%	2652	128	0	POST	/api/posts	3002.056	1.132	7.429	7.792	622
4.3%	13030	0	0	GET	^/api/feeds/[^/]+	2444.753	0.188	1.504	2.258	5115
3.7%	5866	80	0	POST	/api/friend-requests	2087.266	0.356	2.061	2.883	114
3.6%	4249	0	0	GET	/api/memories	2043.373	0.481	2.217	3.723	12534
3.0%	17734	269	0	GET	^/api/images/[^/]+	1713.354	0.097	1.168	2.293	126898
2.2%	3824	0	0	GET	/api/posting-window/current	1260.311	0.330	1.916	2.613	237
1.6%	2869	0	0	GET	/api/users	882.968	0.308	1.868	2.275	237
1.3%	6321	733	0	POST	/api/auth/login	755.403	0.120	1.761	2.813	219
1.3%	714	122	0	PUT	/api/me/profile-image	729.031	1.021	13.907	16.705	2
1.0%	3432	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	589.727	0.172	1.382	2.120	116
1.0%	2004	12	0	GET	^/api/posts/me/current$	587.736	0.293	1.355	1.995	649
0.9%	669	33	0	POST	^/api/posts/[^/]+/realmojis$	486.889	0.728	2.976	3.411	450
0.7%	353	0	0	GET	/api/me/export	409.814	1.161	4.242	5.557	1964963
0.6%	602	3	0	GET	/api/me	362.758	0.603	2.262	2.394	301
0.5%	831	7	0	POST	^/api/posts/[^/]+/comments$	308.700	0.371	2.251	2.684	344
0.5%	2177	58	0	GET	^/api/posts/[^/]+$	271.692	0.125	1.154	1.643	614
0.4%	1542	0	0	GET	^/api/relationships/[^/]+	240.201	0.156	1.196	2.004	1310
0.4%	892	0	0	GET	^/api/memories/[^/]+	230.879	0.259	1.853	2.138	691
0.3%	674	0	0	PUT	^/api/follows/[^/]+$	167.330	0.248	1.982	2.526	0
0.3%	494	0	0	PATCH	/api/me	164.238	0.332	2.417	2.859	277
0.3%	929	0	0	POST	/api/auth/logout	153.532	0.165	1.765	2.169	0
0.2%	1096	0	0	GET	^/api/users/[^/]+$	139.695	0.127	1.172	1.753	209
0.2%	853	0	0	GET	^/api/posts/[^/]+/realmojis$	117.108	0.137	1.100	1.738	1593
0.2%	834	0	0	GET	^/api/posts/[^/]+/comments$	114.226	0.137	1.086	1.622	2356
0.2%	1057	1	0	POST	/api/auth/register	111.319	0.105	2.160	2.262	227
0.2%	202	1	0	PATCH	^/api/posts/[^/]+$	110.077	0.545	2.260	2.354	644
0.2%	91	0	0	DELETE	^/api/posts/[^/]+$	85.635	0.941	2.543	2.543	0
0.1%	121	0	0	DELETE	^/api/realmojis/[^/]+$	42.850	0.354	1.994	2.026	0
0.1%	170	0	0	DELETE	^/api/comments/[^/]+$	31.237	0.184	1.304	1.808	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
67.5%	1296	1	0	GET	/api/notifications	15730.374	12.138	20.101	20.104	106
5.6%	1223	56	0	POST	/api/posts	1307.972	1.069	7.118	7.460	620
5.4%	2745	31	0	POST	/api/friend-requests	1247.466	0.454	1.654	2.607	114
3.2%	5201	0	0	GET	^/api/feeds/[^/]+	734.936	0.141	1.317	2.242	3624
2.9%	1552	0	0	POST	/api/auth/login	681.702	0.439	2.148	2.912	245
2.8%	5211	25	0	GET	^/api/images/[^/]+	661.645	0.127	1.121	2.487	170434
2.2%	514	0	0	POST	/api/auth/register	512.140	0.996	2.322	2.390	229
1.2%	1800	0	0	GET	/api/posting-window/current	285.006	0.158	1.162	2.114	239
1.2%	283	0	0	PUT	/api/me/profile-image	269.542	0.952	2.566	2.772	0
0.9%	1233	0	0	GET	/api/memories	219.090	0.178	1.948	2.502	14489
0.9%	1317	0	0	GET	/api/users	215.183	0.163	1.037	2.033	237
0.8%	502	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	189.460	0.377	1.482	2.301	116
0.8%	240	12	0	POST	^/api/posts/[^/]+/realmojis$	180.040	0.750	2.836	3.274	455
0.7%	881	47	0	GET	^/api/posts/me/current$	153.264	0.174	1.682	2.080	613
0.6%	358	0	0	PUT	^/api/follows/[^/]+$	135.786	0.379	1.458	1.666	0
0.5%	241	0	0	PATCH	/api/me	117.545	0.488	1.671	1.688	275
0.5%	292	4	0	POST	^/api/posts/[^/]+/comments$	113.352	0.388	1.834	2.499	328
0.4%	602	0	0	GET	^/api/relationships/[^/]+	85.432	0.142	1.587	1.946	688
0.4%	754	26	0	GET	^/api/posts/[^/]+$	82.851	0.110	0.615	1.464	611
0.3%	205	0	0	POST	/api/auth/logout	58.771	0.287	1.501	2.139	0
0.2%	83	4	0	PATCH	^/api/posts/[^/]+$	54.207	0.653	1.693	1.693	608
0.2%	292	2	0	GET	/api/me	42.329	0.145	0.987	1.017	302
0.2%	381	0	0	GET	^/api/users/[^/]+$	39.806	0.104	0.569	1.092	213
0.2%	285	1	0	GET	^/api/posts/[^/]+/realmojis$	38.035	0.133	1.049	1.953	1857
0.1%	282	0	0	GET	^/api/posts/[^/]+/comments$	33.167	0.118	1.023	1.157	3249
0.1%	196	0	0	GET	^/api/memories/[^/]+	26.805	0.137	1.697	2.099	990
0.1%	61	0	0	DELETE	^/api/comments/[^/]+$	21.261	0.349	2.126	2.126	0
0.1%	28	0	0	DELETE	^/api/posts/[^/]+$	20.999	0.750	2.240	2.240	0
0.1%	41	0	0	DELETE	^/api/realmojis/[^/]+$	20.154	0.492	2.175	2.175	0
0.0%	72	0	0	GET	/api/me/export	10.109	0.140	1.602	1.602	2011227
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
69.3%	1061	2	0	GET	/api/notifications	12312.965	11.605	20.135	20.809	101
5.5%	2233	34	0	POST	/api/friend-requests	974.679	0.436	1.950	2.935	114
3.9%	975	43	0	POST	/api/posts	699.363	0.717	7.072	7.444	630
3.8%	4939	36	0	GET	^/api/images/[^/]+	673.704	0.136	1.464	2.208	141957
2.7%	3650	0	0	GET	^/api/feeds/[^/]+	488.455	0.134	1.165	1.659	3570
2.5%	1534	0	0	POST	/api/auth/login	450.417	0.294	1.649	2.187	248
1.2%	280	0	0	PUT	/api/me/profile-image	220.514	0.788	2.262	2.682	0
1.1%	460	0	0	POST	/api/auth/register	198.647	0.432	0.995	1.443	230
1.0%	484	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	185.316	0.383	1.882	2.160	116
1.0%	1178	0	0	GET	/api/memories	173.035	0.147	1.519	2.782	17478
1.0%	184	10	0	POST	^/api/posts/[^/]+/realmojis$	170.289	0.925	3.048	3.279	449
0.9%	1499	0	0	GET	/api/posting-window/current	160.834	0.107	0.924	1.340	240
0.9%	300	0	0	PUT	^/api/follows/[^/]+$	158.562	0.529	1.507	1.553	0
0.8%	1111	0	0	GET	/api/users	136.840	0.123	1.030	1.772	240
0.7%	221	0	0	PATCH	/api/me	120.146	0.544	1.837	2.456	283
0.6%	742	40	0	GET	^/api/posts/me/current$	110.521	0.149	1.198	1.303	622
0.4%	219	0	0	POST	/api/auth/logout	78.974	0.361	2.183	2.738	0
0.4%	157	3	0	POST	^/api/posts/[^/]+/comments$	76.905	0.490	2.145	2.595	337
0.4%	487	0	0	GET	^/api/relationships/[^/]+	65.695	0.135	1.014	1.330	831
0.4%	587	29	0	GET	^/api/posts/[^/]+$	62.507	0.106	0.894	1.110	601
0.3%	285	1	0	GET	/api/me	46.746	0.164	0.669	0.840	305
0.2%	77	4	0	PATCH	^/api/posts/[^/]+$	36.270	0.471	2.014	2.014	618
0.2%	39	0	0	DELETE	^/api/posts/[^/]+$	33.064	0.848	2.505	2.505	0
0.2%	249	0	0	GET	^/api/users/[^/]+$	27.042	0.109	0.614	1.064	212
0.1%	196	0	0	GET	^/api/posts/[^/]+/realmojis$	21.481	0.110	0.699	0.944	1714
0.1%	25	0	0	DELETE	^/api/realmojis/[^/]+$	19.472	0.779	2.495	2.495	0
0.1%	168	1	0	GET	^/api/posts/[^/]+/comments$	19.405	0.116	0.608	1.253	2734
0.1%	191	0	0	GET	^/api/memories/[^/]+	17.506	0.092	0.401	0.524	1039
0.1%	88	0	0	GET	/api/me/export	15.123	0.172	4.460	4.460	2558513
0.1%	27	0	0	DELETE	^/api/comments/[^/]+$	10.592	0.392	1.408	1.408	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16524 クエリ / 48 種類 / 合計 176.823s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	53.3%	357	94.171	263.78	1210.02	1245.4	127.1	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.6%	4	25.811	6452.86	10048.18	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.5%	12334	20.347	1.65	11.01	9.7	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.7%	4	18.849	4712.22	7177.51	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.1%	4	9.044	2261.06	3445.04	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.6%	7	4.569	652.74	3277.04	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	1	0.646	645.96	645.96	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.3%	7	0.585	83.53	161.24	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.3%	7	0.581	82.93	160.91	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.2%	1	0.439	439.03	439.03	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.406	406.39	406.39	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.2%	1	0.303	302.76	302.76	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.219	218.82	218.82	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.180	179.66	179.66	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	1	0.109	109.16	109.16	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows

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
Build ID: d0163c9fb6a23274a34cb1b844236f449fe443c6
Type: cpu
Time: 2026-10-31 17:30:37 JST
Duration: 60s, Total samples = 30.39s (50.65%)
Showing nodes accounting for 23.45s, 77.16% of 30.39s total
Dropped 862 nodes (cum <= 0.15s)
Showing top 40 nodes out of 245
      flat  flat%   sum%        cum   cum%
     5.86s 19.28% 19.28%      5.86s 19.28%  crypto/internal/fips140/sha256.blockAVX2
     3.28s 10.79% 30.08%      3.28s 10.79%  internal/runtime/syscall/linux.Syscall6
     1.64s  5.40% 35.47%      5.66s 18.62%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.30s  4.28% 39.75%      1.30s  4.28%  runtime.memmove
     1.26s  4.15% 43.90%      1.26s  4.15%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.01s  3.32% 47.22%      1.01s  3.32%  aeshashbody
     0.86s  2.83% 50.05%      3.34s 10.99%  runtime.scanObject
     0.81s  2.67% 52.71%      1.09s  3.59%  runtime.findObject
     0.79s  2.60% 55.31%      0.79s  2.60%  runtime.memclrNoHeapPointers
     0.66s  2.17% 57.49%      1.09s  3.59%  runtime.tryDeferToSpanScan
     0.65s  2.14% 59.62%      1.81s  5.96%  runtime.mapaccess1_faststr
     0.40s  1.32% 60.94%      6.76s 22.24%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.37s  1.22% 62.16%      6.58s 21.65%  crypto/internal/fips140/sha256.(*Digest).Write
     0.36s  1.18% 63.34%      0.59s  1.94%  example.com/isureal-company/isureal.postLess
     0.30s  0.99% 64.33%      0.30s  0.99%  runtime.nextFreeFast (inline)
     0.29s  0.95% 65.28%      0.29s  0.95%  runtime.spanClass.sizeclass (inline)
     0.27s  0.89% 66.17%      0.96s  3.16%  runtime.mallocgcSmallScanNoHeader
     0.26s  0.86% 67.03%      0.37s  1.22%  database/sql.convertAssignRows
     0.26s  0.86% 67.88%      0.26s  0.86%  runtime.futex
     0.25s  0.82% 68.71%      0.25s  0.82%  sync/atomic.(*Int32).Add (inline)
     0.22s  0.72% 69.43%      0.22s  0.72%  memeqbody
     0.19s  0.63% 70.06%      7.10s 23.36%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.18s  0.59% 70.65%      2.53s  8.33%  runtime.mallocgc
     0.18s  0.59% 71.24%      1.44s  4.74%  runtime.mapassign_faststr
     0.17s  0.56% 71.80%      7.53s 24.78%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.17s  0.56% 72.36%      0.20s  0.66%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.17s  0.56% 72.92%      0.18s  0.59%  runtime.(*mspan).writeHeapBitsSmall
     0.16s  0.53% 73.45%      0.34s  1.12%  runtime.greyobject
     0.14s  0.46% 73.91%         6s 19.74%  crypto/internal/fips140/sha256.block
     0.12s  0.39% 74.30%      1.84s  6.05%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s  0.36% 74.66%      1.07s  3.52%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.11s  0.36% 75.02%      0.22s  0.72%  internal/runtime/maps.(*Iter).Next
     0.10s  0.33% 75.35%      0.16s  0.53%  runtime.(*mspan).markBitsForIndex (inline)
     0.09s   0.3% 75.65%      0.28s  0.92%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.09s   0.3% 75.95%      3.70s 12.18%  example.com/isureal-company/isureal.loadBig
     0.09s   0.3% 76.24%      0.18s  0.59%  runtime.scanObjectsSmall
     0.08s  0.26% 76.51%      0.45s  1.48%  database/sql.(*Rows).scanLocked
     0.07s  0.23% 76.74%      0.23s  0.76%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.07s  0.23% 76.97%      0.37s  1.22%  runtime.mapaccess2_faststr
     0.06s   0.2% 77.16%      1.43s  4.71%  runtime.growslice
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 33.1% / 100% | 14% | 0.6% | 8 | 134 MB | run queue が積もっている |
| isu2 (2 cpu) | 28.9% / 100% | 6% | 0.6% | 4 | 164 MB |  |
| isu3 (2 cpu) | 21.4% / 100% | 4% | 0.8% | 5 | 152 MB | run queue が積もっている |
| isu4 (2 cpu) | 49.8% / 100% | 35% | 7.4% | 9 | 105 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 52% / nginx 10% |
| isu2 | isureal 54% / nginx 2% |
| isu3 | isureal 40% / nginx 2% |
| isu4 | isureal 64% / postgres 18% |

## Diff from previous run

スコア: 223434 → 205640 (-17794)

`4c078d8..21d10c4`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf |   1 +
 conf/isu1/home/isucon/env.sh                   |   1 -
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |   1 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |   1 +
 conf/isu4/home/isucon/env.sh                   |   3 +-
 webapp/go/image_signature.go                   |  61 ++++-----------
 webapp/go/offload.go                           |  32 +++++---
 webapp/go/presign.go                           | 103 -------------------------
 8 files changed, 42 insertions(+), 161 deletions(-)
```
