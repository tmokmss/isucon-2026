# run 20261031-160039-1f0e740

## Summary

| | |
|---|---|
| score | 181170 |
| status | pass |
| ref | presign-next-minute |
| sha | 1f0e74082a6f |
| 所要時間 | 359s |
| 前回比 | -3660 (-2.0%) vs 20261031-155429-b3440cf |
| ベスト比 | -10646 (-5.6%) vs 20261031-151339-5a230a4 = 191816 |
| note | on b3440cf: sig cache by exp generation + presign (JPEG unchanged), heap capture |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	135	390	1071	1072
europe	7	21	69	197	542	541
america	6	16	47	129	273	272
```

```
{"id":"26116","status":"done","portalStatus":"finished","passed":true,"score":181170,"raw":181170,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:05:01.147Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:05:04.149Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:05:07.142Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:05:10.157Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T07:05:13.154Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T07:05:16.142Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T07:05:18.347Z level=INFO msg="user moved" user=nivin5391 from=europe to=asia
time=2026-10-31T07:05:19.152Z level=INFO msg="asia 3 日目: 135 人"
time=2026-10-31T07:05:20.004Z level=INFO msg="user moved" user=wamewa.5656 from=asia to=europe
time=2026-10-31T07:05:20.254Z level=INFO msg="user moved" user=benkasu5603 from=asia to=europe
time=2026-10-31T07:05:20.702Z level=INFO msg="user moved" user=wamewa.5656 from=europe to=asia
time=2026-10-31T07:05:20.819Z level=INFO msg="user moved" user=nivin5391 from=asia to=america
time=2026-10-31T07:05:20.884Z level=INFO msg="user moved" user=hois-5635 from=asia to=europe
time=2026-10-31T07:05:22.144Z level=INFO msg="europe 3 日目: 69 人"
time=2026-10-31T07:05:22.151Z level=INFO msg="user moved" user=benkasu5603 from=europe to=asia
time=2026-10-31T07:05:25.142Z level=INFO msg="america 3 日目: 47 人"
time=2026-10-31T07:05:28.188Z level=INFO msg="asia 4 日目: 390 人"
time=2026-10-31T07:05:28.813Z level=INFO msg="user moved" user=noulben.6327 from=europe to=america
time=2026-10-31T07:05:31.199Z level=INFO msg="user moved" user=benkasu5603 from=asia to=europe
time=2026-10-31T07:05:31.220Z level=INFO msg="europe 4 日目: 197 人"
time=2026-10-31T07:05:31.266Z level=INFO msg="user moved" user=wamewa.5656 from=asia to=europe
time=2026-10-31T07:05:31.321Z level=INFO msg="user moved" user=hianko_6451 from=america to=europe
time=2026-10-31T07:05:34.312Z level=INFO msg="america 4 日目: 129 人"
time=2026-10-31T07:05:34.721Z level=INFO msg="user moved" user=rona7501 from=asia to=america
time=2026-10-31T07:05:34.876Z level=INFO msg="user moved" user=heku_7125 from=asia to=america
time=2026-10-31T07:05:35.181Z level=INFO msg="user moved" user=yare-7945 from=asia to=america
time=2026-10-31T07:05:37.083Z level=INFO msg="user moved" user=orki9430 from=europe to=asia
time=2026-10-31T07:05:37.167Z level=INFO msg="asia 5 日目: 1071 人"
time=2026-10-31T07:05:40.817Z level=INFO msg="europe 5 日目: 542 人"
time=2026-10-31T07:05:41.997Z level=INFO msg="user moved" user=nivin5391 from=america to=asia
time=2026-10-31T07:05:44.755Z level=INFO msg="america 5 日目: 273 人"
time=2026-10-31T07:05:51.848Z level=INFO msg="user moved" user=nusa-10418 from=america to=asia
time=2026-10-31T07:05:52.733Z level=INFO msg="america 最終日（グランプリ）: 272 人"
time=2026-10-31T07:05:52.741Z level=INFO msg="europe 最終日（グランプリ）: 541 人"
time=2026-10-31T07:05:52.862Z level=INFO msg="asia 最終日（グランプリ）: 1072 人"
time=2026-10-31T07:06:01.228Z level=INFO msg="グランプリ: アジア 373 人, ヨーロッパ 227 人, アメリカ 130 人"
{
  "pass": true,
  "score": 181170,
  "messages": [
    "bentoso_5049: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "rishi17963: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nuivy17584: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nuulko_17825: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "elri17889: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marnior.17635: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marhi_18058: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hewa-17707: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "semareva_15870: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tsuchi-10943: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "zoechi_16629: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "neor-17077: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yaku18067: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "benkite17870: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nimu.18095: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mewa-5709: 友達申請を承認できません: POST /api/friend-requests/fcb85c25-aa10-44bc-acd2-19fa15f29fc6/accept: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests/fcb85c25-aa10-44bc-acd2-19fa15f29fc6/accept\": EOF",
    "rureke-7373: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "gansik_yesan_2619: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "yui_kato_3967: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "doyun_cho_4427: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.3%	1707	1	0	GET	/api/notifications	15664.184	9.176	16.090	16.096	171
11.5%	2693	196	0	POST	/api/posts	4088.304	1.518	6.594	63.854	606
8.1%	17907	289	0	GET	^/api/images/[^/]+	2861.855	0.160	1.680	3.352	116832
7.7%	5948	93	0	POST	/api/friend-requests	2724.584	0.458	2.564	3.600	114
3.7%	6561	649	0	POST	/api/auth/login	1324.223	0.202	1.741	3.452	222
3.5%	11678	0	0	GET	^/api/feeds/[^/]+	1243.190	0.106	0.742	1.601	3118
3.1%	2827	6	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1087.140	0.385	2.589	3.505	116
3.0%	4194	0	0	GET	/api/memories	1046.190	0.249	2.398	4.135	13046
2.0%	736	127	0	PUT	/api/me/profile-image	693.962	0.943	10.130	15.015	3
1.6%	713	9	0	POST	^/api/posts/[^/]+/realmojis$	558.265	0.783	3.988	4.683	467
1.5%	1092	1	0	POST	/api/auth/register	519.951	0.476	1.767	2.013	227
1.3%	3851	0	0	GET	/api/posting-window/current	444.481	0.115	0.832	1.437	237
1.0%	2891	0	0	GET	/api/users	362.227	0.125	0.673	1.202	240
0.9%	702	0	0	PUT	^/api/follows/[^/]+$	332.753	0.474	1.923	3.065	0
0.8%	498	0	0	PATCH	/api/me	294.475	0.591	2.097	2.525	279
0.8%	387	0	0	GET	/api/me/export	285.277	0.737	6.231	8.935	1932226
0.7%	499	1	0	POST	^/api/posts/[^/]+/comments$	263.259	0.528	3.031	4.006	341
0.7%	1996	8	0	GET	^/api/posts/me/current$	231.854	0.116	0.742	1.178	651
0.5%	993	0	0	POST	/api/auth/logout	193.826	0.195	1.852	3.586	0
0.5%	620	3	0	GET	/api/me	171.005	0.276	1.041	1.107	301
0.4%	1713	50	0	GET	^/api/posts/[^/]+$	159.270	0.093	0.640	0.952	612
0.4%	1548	0	0	GET	^/api/relationships/[^/]+	156.791	0.101	0.669	1.029	1100
0.4%	229	2	0	PATCH	^/api/posts/[^/]+$	145.667	0.636	2.708	2.752	650
0.3%	102	0	0	DELETE	^/api/posts/[^/]+$	121.525	1.191	3.155	3.785	0
0.3%	987	0	0	GET	^/api/memories/[^/]+	112.209	0.114	0.860	2.210	634
0.2%	844	0	0	GET	^/api/users/[^/]+$	73.695	0.087	0.625	0.742	199
0.2%	140	0	0	DELETE	^/api/realmojis/[^/]+$	62.978	0.450	1.599	2.907	0
0.2%	821	0	0	GET	^/api/posts/[^/]+/realmojis$	58.983	0.072	0.544	0.882	1676
0.1%	609	0	0	GET	^/api/posts/[^/]+/comments$	49.104	0.081	0.535	0.863	1959
0.1%	90	0	0	DELETE	^/api/comments/[^/]+$	26.168	0.291	2.576	2.576	0
…（上位 30 / 全 43 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.1%	821	0	0	GET	/api/notifications	5605.753	6.828	11.905	11.907	164
11.4%	2744	26	0	POST	/api/friend-requests	1591.220	0.580	2.757	3.871	114
10.0%	1203	100	0	POST	/api/posts	1393.234	1.158	7.086	7.472	594
5.8%	4604	31	0	GET	^/api/images/[^/]+	804.915	0.175	1.750	3.623	182063
5.2%	4309	0	0	GET	^/api/feeds/[^/]+	721.146	0.167	1.434	2.513	2926
3.9%	1470	0	0	POST	/api/auth/login	546.420	0.372	2.275	3.029	244
3.4%	539	0	0	POST	/api/auth/register	477.173	0.885	2.322	2.349	229
2.0%	293	0	0	PUT	/api/me/profile-image	277.299	0.946	2.655	2.730	0
2.0%	1825	0	0	GET	/api/posting-window/current	275.434	0.151	0.928	1.967	239
1.8%	1355	0	0	GET	/api/users	257.624	0.190	1.487	2.333	240
1.7%	478	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	231.220	0.484	2.539	3.551	115
1.7%	353	0	0	PUT	^/api/follows/[^/]+$	230.462	0.653	2.743	3.175	0
1.6%	268	0	0	PATCH	/api/me	226.775	0.846	2.895	3.102	281
1.6%	1191	0	0	GET	/api/memories	220.877	0.185	2.101	2.876	12195
1.2%	882	49	0	GET	^/api/posts/me/current$	170.854	0.194	2.016	2.360	612
1.1%	281	1	0	POST	^/api/posts/[^/]+/comments$	155.237	0.552	3.094	3.617	329
0.9%	144	2	0	POST	^/api/posts/[^/]+/realmojis$	126.986	0.882	3.864	5.119	465
0.8%	619	0	0	GET	^/api/relationships/[^/]+	112.157	0.181	1.870	2.190	550
0.7%	663	28	0	GET	^/api/posts/[^/]+$	100.111	0.151	1.000	2.123	603
0.5%	93	8	0	PATCH	^/api/posts/[^/]+$	67.659	0.728	2.995	2.995	582
0.5%	300	2	0	GET	/api/me	65.820	0.219	1.584	1.585	302
0.4%	198	0	0	POST	/api/auth/logout	61.904	0.313	2.182	2.472	0
0.3%	312	0	0	GET	^/api/users/[^/]+$	43.509	0.139	0.817	2.003	199
0.3%	37	0	0	DELETE	^/api/posts/[^/]+$	42.452	1.147	3.353	3.353	0
0.2%	203	0	0	GET	^/api/posts/[^/]+/realmojis$	33.103	0.163	1.012	1.487	1766
0.2%	234	0	0	GET	^/api/posts/[^/]+/comments$	33.076	0.141	1.086	1.092	1941
0.2%	60	0	0	DELETE	^/api/comments/[^/]+$	21.992	0.367	2.073	2.073	0
0.2%	26	0	0	DELETE	^/api/realmojis/[^/]+$	21.777	0.838	2.985	2.985	0
0.1%	73	0	0	GET	/api/me/export	19.146	0.262	4.485	4.485	2210525
0.1%	179	0	0	GET	^/api/memories/[^/]+	19.010	0.106	0.710	1.723	824
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.2%	415	0	0	GET	/api/notifications	3138.858	7.564	10.449	10.457	145
9.5%	1441	15	0	POST	/api/friend-requests	673.571	0.467	2.699	2.888	114
8.8%	577	45	0	POST	/api/posts	623.147	1.080	6.915	7.761	603
6.4%	1023	0	0	POST	/api/auth/login	455.071	0.445	2.860	3.098	250
6.1%	3301	20	0	GET	^/api/images/[^/]+	431.370	0.131	1.077	3.390	134289
4.4%	2246	0	0	GET	^/api/feeds/[^/]+	309.190	0.138	1.092	1.302	3164
3.2%	269	0	0	POST	/api/auth/register	228.520	0.850	2.883	2.911	230
2.0%	150	0	0	PUT	/api/me/profile-image	139.787	0.932	3.785	3.849	0
1.9%	304	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	132.906	0.437	2.653	2.815	116
1.6%	806	0	0	GET	/api/memories	114.719	0.142	1.525	2.096	17288
1.4%	709	0	0	GET	/api/users	95.916	0.135	1.074	1.233	241
1.3%	132	3	0	POST	^/api/posts/[^/]+/realmojis$	94.396	0.715	2.585	2.650	476
1.2%	857	0	0	GET	/api/posting-window/current	87.581	0.102	0.762	1.279	240
1.0%	121	0	0	PATCH	/api/me	73.259	0.605	2.744	2.744	279
0.9%	441	18	0	GET	^/api/posts/me/current$	66.249	0.150	1.205	1.393	628
0.9%	149	0	0	POST	/api/auth/logout	64.918	0.436	2.530	2.838	0
0.8%	138	0	0	PUT	^/api/follows/[^/]+$	57.573	0.417	2.315	2.881	0
0.7%	102	3	0	POST	^/api/posts/[^/]+/comments$	50.090	0.491	2.606	2.690	319
0.7%	308	0	0	GET	^/api/relationships/[^/]+	46.814	0.152	1.063	1.125	753
0.6%	324	14	0	GET	^/api/posts/[^/]+$	40.060	0.124	0.957	1.125	606
0.5%	25	0	0	DELETE	^/api/posts/[^/]+$	34.161	1.366	3.159	3.159	0
0.4%	41	1	0	PATCH	^/api/posts/[^/]+$	30.466	0.743	2.891	2.891	625
0.3%	153	1	0	GET	/api/me	22.612	0.148	1.235	1.243	303
0.3%	29	0	0	DELETE	^/api/realmojis/[^/]+$	20.319	0.701	2.362	2.362	0
0.3%	137	0	0	GET	^/api/memories/[^/]+	17.910	0.131	1.244	1.323	1026
0.2%	52	0	0	GET	/api/me/export	12.850	0.247	6.156	6.156	2766651
0.2%	125	0	0	GET	^/api/users/[^/]+$	12.333	0.099	0.357	0.679	199
0.2%	119	0	0	GET	^/api/posts/[^/]+/realmojis$	11.515	0.097	0.254	1.062	2001
0.1%	77	0	0	GET	^/api/posts/[^/]+/comments$	8.857	0.115	1.000	1.000	1727
0.0%	14	0	0	DELETE	^/api/comments/[^/]+$	3.370	0.241	0.662	0.662	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 15493 クエリ / 48 種類 / 合計 146.655s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	53.6%	395	78.650	199.11	908.21	1213.9	124.2	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	15.2%	3	22.336	7445.43	10063.55	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.1%	3	16.287	5429.10	7309.80	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	10.8%	12877	15.910	1.24	9.74	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	5.3%	3	7.757	2585.76	3480.97	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.3%	4	3.444	861.05	2884.90	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.3%	4	0.390	97.58	157.56	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
8	0.3%	4	0.371	92.82	157.45	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.2%	1	0.269	269.22	269.22	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.250	250.03	250.03	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.2%	1	0.229	229.30	229.30	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.159	159.07	159.07	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.104	103.55	103.55	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.084	83.99	83.99	2486.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.1%	1	0.078	78.13	78.13	1549.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
Build ID: f484fa262972f991e07aed1a54eb28ef3ea2570e
Type: cpu
Time: 2026-10-31 16:04:41 JST
Duration: 60.24s, Total samples = 44.98s (74.67%)
Showing nodes accounting for 32.81s, 72.94% of 44.98s total
Dropped 1137 nodes (cum <= 0.22s)
Showing top 40 nodes out of 273
      flat  flat%   sum%        cum   cum%
     6.79s 15.10% 15.10%      6.79s 15.10%  crypto/internal/fips140/sha256.blockAVX2
     5.04s 11.20% 26.30%      5.04s 11.20%  internal/runtime/syscall/linux.Syscall6
     2.03s  4.51% 30.81%      2.03s  4.51%  runtime.memmove
     1.81s  4.02% 34.84%      6.52s 14.50%  runtime.scanObject
     1.77s  3.94% 38.77%      2.36s  5.25%  runtime.findObject
     1.63s  3.62% 42.40%      1.63s  3.62%  runtime.memclrNoHeapPointers
     1.57s  3.49% 45.89%      5.69s 12.65%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.49s  3.31% 49.20%      1.49s  3.31%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.41s  3.13% 52.33%      2.47s  5.49%  runtime.tryDeferToSpanScan
     1.22s  2.71% 55.05%      1.22s  2.71%  aeshashbody
     0.75s  1.67% 56.71%      0.75s  1.67%  runtime.spanClass.sizeclass (inline)
     0.61s  1.36% 58.07%      2.12s  4.71%  runtime.mapaccess1_faststr
     0.47s  1.04% 59.12%      0.65s  1.45%  example.com/isureal-company/isureal.postLess
     0.43s  0.96% 60.07%      7.69s 17.10%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.35s  0.78% 60.85%      0.35s  0.78%  memeqbody
     0.34s  0.76% 61.61%      7.43s 16.52%  crypto/internal/fips140/sha256.(*Digest).Write
     0.34s  0.76% 62.36%      0.34s  0.76%  runtime.nextFreeFast (inline)
     0.32s  0.71% 63.07%      0.62s  1.38%  database/sql.convertAssignRows
     0.28s  0.62% 63.69%      0.29s  0.64%  runtime.(*mspan).writeHeapBitsSmall
     0.28s  0.62% 64.32%      0.28s  0.62%  sync/atomic.(*Int32).Add (inline)
     0.27s   0.6% 64.92%      8.67s 19.28%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.26s  0.58% 65.50%      8.07s 17.94%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.26s  0.58% 66.07%      0.57s  1.27%  runtime.greyobject
     0.24s  0.53% 66.61%      0.26s  0.58%  encoding/json.appendString[go.shape.string]
     0.24s  0.53% 67.14%      1.24s  2.76%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.53% 67.67%      0.27s   0.6%  runtime.spanOf (inline)
     0.23s  0.51% 68.19%      0.23s  0.51%  runtime.futex
     0.22s  0.49% 68.67%      4.45s  9.89%  runtime.mallocgc
     0.21s  0.47% 69.14%      0.49s  1.09%  compress/flate.(*compressor).deflate
     0.20s  0.44% 69.59%         2s  4.45%  runtime.mapassign_faststr
     0.18s   0.4% 69.99%      0.25s  0.56%  runtime.(*mspan).markBitsForIndex (inline)
     0.18s   0.4% 70.39%      0.54s  1.20%  runtime.scanObjectsSmall
     0.17s  0.38% 70.76%      0.43s  0.96%  internal/stringslite.Index
     0.16s  0.36% 71.12%      9.12s 20.28%  example.com/isureal-company/isureal.computeImageSignature
     0.15s  0.33% 71.45%      0.75s  1.67%  encoding/json.structEncoder.encode
     0.14s  0.31% 71.77%      0.28s  0.62%  internal/runtime/maps.(*Iter).Next
     0.14s  0.31% 72.08%      0.24s  0.53%  runtime.heapArenaOf (inline)
     0.13s  0.29% 72.37%      0.75s  1.67%  database/sql.(*Rows).scanLocked
     0.13s  0.29% 72.65%      3.77s  8.38%  example.com/isureal-company/isureal.loadBig
     0.13s  0.29% 72.94%      1.18s  2.62%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 46.6% / 100% | 30% | 1.4% | 7 | 115 MB | run queue が積もっている |
| isu2 (2 cpu) | 35.6% / 100% | 10% | 0.7% | 4 | 199 MB |  |
| isu3 (2 cpu) | 24.9% / 99% | 2% | 0.6% | 4 | 720 MB |  |
| isu4 (2 cpu) | 41.6% / 100% | 11% | 16.2% | 9 | 103 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 71% / nginx 12% |
| isu2 | isureal 67% / nginx 2% |
| isu3 | isureal 47% / nginx 1% |
| isu4 | isureal 30% / postgres 15% |

## Diff from previous run

スコア: 184830 → 181170 (-3660)

`b3440cf..1f0e740`

```
 webapp/go/image_signature.go | 67 ++++++++++++++++++++++++++++++++++----------
 webapp/go/offload.go         | 63 +++++++++++++++++++++++++++++++----------
 webapp/go/server.go          |  1 +
 3 files changed, 101 insertions(+), 30 deletions(-)
```
