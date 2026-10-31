# run 20261031-145836-9c97c71

## Summary

| | |
|---|---|
| score | 137455 |
| status | pass |
| ref | forward-fifo-2 |
| sha | 9c97c718b2f4 |
| 所要時間 | 253s |
| 前回比 | +4602 (3.5%) vs 20261031-141932-6717755 |
| ベスト比 | -3880 (-2.7%) vs 20261031-141027-bd0fe45 = 141335 |
| note | 投稿の本文を 16 件ずつ先着順に asia1 へ（main 6ad2e36 の上） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	135	385	986	986
europe	7	17	46	127	234	234
america	6	15	41	105	152	152
```

```
{"id":"20326","status":"done","portalStatus":"finished","passed":true,"score":137455,"raw":137455,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:01:08.217Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:01:11.208Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:01:14.206Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:01:17.218Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T06:01:17.777Z level=INFO msg="user moved" user=suwa_5175 from=asia to=america
time=2026-10-31T06:01:19.009Z level=INFO msg="user moved" user=suwa_5175 from=america to=asia
time=2026-10-31T06:01:20.206Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T06:01:23.208Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T06:01:26.217Z level=INFO msg="asia 3 日目: 135 人"
time=2026-10-31T06:01:28.986Z level=INFO msg="user moved" user=suwa_5175 from=asia to=america
time=2026-10-31T06:01:29.206Z level=INFO msg="europe 3 日目: 46 人"
time=2026-10-31T06:01:32.206Z level=INFO msg="america 3 日目: 41 人"
time=2026-10-31T06:01:35.218Z level=INFO msg="asia 4 日目: 385 人"
time=2026-10-31T06:01:36.508Z level=INFO msg="user moved" user=leoso6475 from=asia to=america
time=2026-10-31T06:01:38.100Z level=INFO msg="user moved" user=suma.6978 from=asia to=europe
time=2026-10-31T06:01:38.328Z level=INFO msg="europe 4 日目: 127 人"
time=2026-10-31T06:01:39.564Z level=INFO msg="user moved" user=kehi.6750 from=asia to=america
time=2026-10-31T06:01:41.308Z level=INFO msg="america 4 日目: 105 人"
time=2026-10-31T06:01:41.418Z level=INFO msg="user moved" user=leoso6475 from=america to=europe
time=2026-10-31T06:01:44.311Z level=INFO msg="asia 5 日目: 986 人"
time=2026-10-31T06:01:46.454Z level=INFO msg="user moved" user=zoeivyul.8594 from=europe to=asia
time=2026-10-31T06:01:47.394Z level=INFO msg="europe 5 日目: 234 人"
time=2026-10-31T06:01:51.654Z level=INFO msg="america 5 日目: 152 人"
time=2026-10-31T06:01:59.456Z level=INFO msg="europe 最終日（グランプリ）: 234 人"
time=2026-10-31T06:01:59.497Z level=INFO msg="america 最終日（グランプリ）: 152 人"
time=2026-10-31T06:01:59.724Z level=INFO msg="asia 最終日（グランプリ）: 986 人"
time=2026-10-31T06:02:08.257Z level=INFO msg="グランプリ: アジア 466 人, ヨーロッパ 55 人, アメリカ 48 人"
{
  "pass": true,
  "score": 137455,
  "messages": [
    "hetorivy-5021: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "karese8039: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "benshilia_6955: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "eltoya12060: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.35.24:33878-\u003e52.69.202.109:8080: write tcp 10.27.35.24:33878-\u003e52.69.202.109:8080: write: broken pipe",
    "nurike11986: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ivyketsu-9494: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kunimar-10591: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nemu_support_2279: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "yang_liu_3304: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "ren_watanabe_3399: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "xiangrikui_shousi_4628: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "hupo_shuizuguan_4951: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "yuzu_wakame_4204: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "tao_zhou_1952: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "raolu_huojian_4584: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "oyatsu_yosan_2671: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "yejun_jang_3468: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "yuxin_wu_1099: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "yotani-5076: 友達申請を承認できません: POST /api/friend-requests/2e67414f-ff1c-4a00-ab94-0d330f88baf2/accept: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests/2e67414f-ff1c-4a00-ab94-0d330f88baf2/accept\": EOF",
    "croquette_sodachi_326: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.7%	1614	1	0	GET	/api/notifications	14685.370	9.099	15.671	15.690	171
16.1%	2531	74	0	POST	/api/posts	5938.446	2.346	8.841	10.395	632
7.1%	5811	81	0	POST	/api/friend-requests	2637.176	0.454	1.848	2.621	114
6.6%	18041	316	0	GET	^/api/images/[^/]+	2436.099	0.135	1.254	2.674	122377
4.0%	4426	0	0	GET	/api/memories	1481.037	0.335	3.308	5.086	12606
4.0%	6447	642	0	POST	/api/auth/login	1474.966	0.229	1.640	2.970	223
3.5%	13361	0	0	GET	^/api/feeds/[^/]+	1300.039	0.097	0.601	2.046	2877
2.9%	2368	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1071.130	0.452	1.940	2.492	116
2.8%	673	126	0	PUT	/api/me/profile-image	1044.848	1.553	7.801	8.916	3
2.4%	659	16	0	POST	^/api/posts/[^/]+/realmojis$	870.437	1.321	8.795	10.423	454
1.5%	3632	1	0	GET	/api/posting-window/current	566.266	0.156	0.749	1.423	237
1.2%	1006	1	0	POST	/api/auth/register	426.985	0.424	1.342	1.804	227
1.1%	717	4	0	POST	^/api/posts/[^/]+/comments$	403.781	0.563	3.064	3.877	339
0.8%	627	0	0	PUT	^/api/follows/[^/]+$	288.302	0.460	1.809	2.880	0
0.8%	2842	0	0	GET	/api/users	278.253	0.098	0.591	1.422	239
0.7%	379	0	0	GET	/api/me/export	261.769	0.691	3.713	5.150	1756232
0.7%	491	0	0	PATCH	/api/me	252.175	0.514	1.903	2.347	272
0.7%	1004	0	0	POST	/api/auth/logout	247.983	0.247	1.605	3.097	0
0.6%	1964	10	0	GET	^/api/posts/me/current$	221.035	0.113	0.671	1.143	647
0.5%	2272	26	0	GET	^/api/posts/[^/]+$	185.175	0.082	0.538	1.177	627
0.5%	1709	0	0	GET	^/api/relationships/[^/]+	175.801	0.103	1.105	4.359	978
0.3%	957	0	0	GET	^/api/memories/[^/]+	114.522	0.120	1.018	1.774	686
0.3%	234	4	0	PATCH	^/api/posts/[^/]+$	110.037	0.470	1.944	2.645	631
0.2%	86	0	0	DELETE	^/api/posts/[^/]+$	85.139	0.990	2.652	2.652	0
0.2%	1031	0	0	GET	^/api/users/[^/]+$	74.466	0.072	0.543	1.011	206
0.2%	960	0	0	GET	^/api/posts/[^/]+/realmojis$	73.401	0.076	0.550	1.000	1427
0.2%	553	3	0	GET	/api/me	68.178	0.123	0.693	0.907	301
0.2%	834	0	0	GET	^/api/posts/[^/]+/comments$	61.620	0.074	0.442	0.740	2052
0.2%	92	0	0	DELETE	^/api/realmojis/[^/]+$	57.910	0.629	2.622	2.622	0
0.1%	135	0	0	DELETE	^/api/comments/[^/]+$	55.087	0.408	1.596	1.834	0
…（上位 30 / 全 43 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.0%	409	0	0	GET	/api/notifications	3265.796	7.985	12.310	12.314	158
15.7%	558	27	0	POST	/api/posts	1464.039	2.624	7.111	8.163	615
12.3%	2373	16	0	GET	^/api/images/[^/]+	1148.037	0.484	2.830	4.024	143911
8.2%	1097	10	0	POST	/api/friend-requests	764.631	0.697	2.440	2.692	114
4.7%	777	0	0	POST	/api/auth/login	441.216	0.568	2.391	2.914	246
3.9%	816	1	0	GET	/api/posting-window/current	364.030	0.446	1.309	1.660	239
3.7%	1586	0	0	GET	^/api/feeds/[^/]+	343.086	0.216	1.244	1.666	2261
2.3%	636	0	0	GET	/api/memories	218.572	0.344	2.429	5.323	13651
1.9%	231	0	0	POST	/api/auth/register	175.444	0.759	2.311	2.346	228
1.8%	555	0	0	GET	/api/users	168.128	0.303	1.406	2.138	243
1.8%	131	0	0	PUT	/api/me/profile-image	166.494	1.271	3.227	3.552	0
1.6%	250	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	151.660	0.607	2.415	2.455	116
1.1%	379	18	0	GET	^/api/posts/me/current$	99.207	0.262	1.279	2.013	612
0.9%	146	0	0	PUT	^/api/follows/[^/]+$	84.340	0.578	2.167	2.456	0
0.9%	103	0	0	PATCH	/api/me	82.308	0.799	2.808	2.878	286
0.6%	134	0	0	POST	/api/auth/logout	59.593	0.445	2.372	2.606	0
0.6%	230	3	0	GET	^/api/relationships/[^/]+	58.688	0.255	1.596	17.814	834
0.5%	136	2	0	GET	/api/me	46.405	0.341	1.510	1.711	299
0.4%	203	12	0	GET	^/api/posts/[^/]+$	35.942	0.177	1.243	1.484	582
0.3%	19	0	0	DELETE	^/api/posts/[^/]+$	29.937	1.576	3.952	3.952	0
0.3%	110	0	0	GET	^/api/memories/[^/]+	28.875	0.262	1.549	1.952	717
0.3%	42	0	0	GET	/api/me/export	27.010	0.643	4.668	4.668	2423015
0.2%	37	6	0	PATCH	^/api/posts/[^/]+$	22.865	0.618	1.752	1.752	544
0.2%	45	3	0	POST	^/api/posts/[^/]+/realmojis$	21.782	0.484	3.259	3.259	464
0.2%	95	1	0	GET	^/api/users/[^/]+$	14.102	0.148	1.328	1.328	188
0.1%	18	2	0	POST	^/api/posts/[^/]+/comments$	11.997	0.666	1.691	1.691	336
0.1%	85	1	0	GET	^/api/posts/[^/]+/realmojis$	11.328	0.133	1.385	1.385	1191
0.1%	69	0	0	GET	^/api/posts/[^/]+/comments$	7.851	0.114	0.937	0.937	1041
0.1%	10	0	0	DELETE	^/api/realmojis/[^/]+$	4.698	0.470	0.966	0.966	0
0.0%	3	0	0	DELETE	^/api/comments/[^/]+$	1.559	0.520	1.311	1.311	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.5%	286	1	0	GET	/api/notifications	2159.499	7.551	10.776	10.778	170
13.7%	1737	9	0	GET	^/api/images/[^/]+	834.863	0.481	2.328	4.269	124062
13.0%	360	20	0	POST	/api/posts	793.832	2.205	5.191	5.790	623
8.8%	787	7	0	POST	/api/friend-requests	534.298	0.679	2.320	2.761	114
4.9%	579	0	0	POST	/api/auth/login	300.874	0.520	2.318	2.464	250
3.9%	306	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	235.291	0.769	2.356	2.761	116
3.7%	518	0	0	GET	/api/posting-window/current	222.569	0.430	1.358	1.598	240
3.3%	879	0	0	GET	^/api/feeds/[^/]+	198.270	0.226	1.049	2.049	2994
2.8%	493	0	0	GET	/api/memories	171.485	0.348	2.360	3.130	16354
1.7%	389	0	0	GET	/api/users	101.800	0.262	1.261	1.891	239
1.2%	261	21	0	GET	^/api/posts/me/current$	71.613	0.274	1.342	1.375	607
1.1%	149	0	0	POST	/api/auth/register	69.763	0.468	1.698	2.152	230
1.1%	85	0	0	PUT	/api/me/profile-image	68.353	0.804	3.795	3.795	0
0.9%	85	0	0	PUT	^/api/follows/[^/]+$	54.007	0.635	2.761	2.761	0
0.7%	68	0	0	PATCH	/api/me	43.356	0.638	2.444	2.444	287
0.6%	79	0	0	POST	/api/auth/logout	39.306	0.498	2.445	2.445	0
0.4%	153	0	0	GET	^/api/relationships/[^/]+	27.318	0.179	1.244	1.360	1425
0.4%	88	0	0	GET	/api/me	26.847	0.305	1.640	1.640	303
0.4%	92	0	0	GET	^/api/memories/[^/]+	25.566	0.278	1.806	1.806	1153
0.3%	17	3	0	POST	^/api/posts/[^/]+/realmojis$	18.919	1.113	3.657	3.657	366
0.3%	29	7	0	PATCH	^/api/posts/[^/]+$	17.500	0.603	2.256	2.256	507
0.2%	97	6	0	GET	^/api/posts/[^/]+$	12.623	0.130	1.038	1.038	591
0.2%	9	0	0	DELETE	^/api/posts/[^/]+$	11.605	1.289	2.988	2.988	0
0.2%	30	0	0	POST	^/api/posts/[^/]+/comments$	10.913	0.364	1.522	1.522	349
0.2%	71	0	0	GET	^/api/users/[^/]+$	10.167	0.143	1.171	1.171	196
0.1%	31	0	0	GET	/api/me/export	7.909	0.255	2.323	2.323	3195044
0.1%	45	0	0	GET	^/api/posts/[^/]+/comments$	7.776	0.173	0.797	0.797	1401
0.1%	37	1	0	GET	^/api/posts/[^/]+/realmojis$	5.350	0.145	1.035	1.035	1322
0.0%	4	0	0	DELETE	^/api/comments/[^/]+$	2.704	0.676	1.469	1.469	0
0.0%	3	0	0	PUT	^/api/blocks/[^/]+$	0.325	0.108	0.115	0.115	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 15110 クエリ / 48 種類 / 合計 136.959s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	50.2%	391	68.710	175.73	1312.69	1113.9	114.1	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	15.6%	3	21.407	7135.79	9542.17	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.6%	3	15.950	5316.78	7190.78	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	11.2%	12942	15.280	1.18	6.25	9.5	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	5.7%	3	7.821	2606.95	3552.72	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.9%	6	3.959	659.89	3022.38	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.9%	1	1.234	1234.11	1234.11	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.5%	6	0.685	114.09	157.07	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.5%	6	0.663	110.53	157.07	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.2%	1	0.250	249.56	249.56	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.227	227.50	227.50	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.167	167.10	167.10	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.103	103.26	103.26	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.081	80.82	80.82	2495.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.1%	1	0.078	77.84	77.84	1554.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
Build ID: bc952a659b787e4a137204d7ae75deb6afa6eab1
Type: cpu
Time: 2026-10-31 15:00:46 JST
Duration: 60.44s, Total samples = 42.29s (69.97%)
Showing nodes accounting for 30.27s, 71.58% of 42.29s total
Dropped 1126 nodes (cum <= 0.21s)
Showing top 40 nodes out of 275
      flat  flat%   sum%        cum   cum%
     5.61s 13.27% 13.27%      5.61s 13.27%  crypto/internal/fips140/sha256.blockAVX2
     4.89s 11.56% 24.83%      4.89s 11.56%  internal/runtime/syscall/linux.Syscall6
     2.09s  4.94% 29.77%      2.09s  4.94%  runtime.memmove
     1.80s  4.26% 34.03%      2.34s  5.53%  runtime.findObject
     1.72s  4.07% 38.09%      6.02s 14.24%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.54s  3.64% 41.74%      1.54s  3.64%  runtime.memclrNoHeapPointers
     1.40s  3.31% 45.05%      1.40s  3.31%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.40s  3.31% 48.36%      6.29s 14.87%  runtime.scanObject
     1.31s  3.10% 51.45%      2.38s  5.63%  runtime.tryDeferToSpanScan
     1.24s  2.93% 54.39%      1.24s  2.93%  aeshashbody
     0.77s  1.82% 56.21%      0.77s  1.82%  runtime.spanClass.sizeclass (inline)
     0.67s  1.58% 57.79%      2.24s  5.30%  runtime.mapaccess1_faststr
     0.42s  0.99% 58.78%      0.42s  0.99%  runtime.nextFreeFast (inline)
     0.41s  0.97% 59.75%      0.41s  0.97%  memeqbody
     0.40s  0.95% 60.70%      0.56s  1.32%  example.com/isureal-company/isureal.postLess
     0.36s  0.85% 61.55%      6.18s 14.61%  crypto/internal/fips140/sha256.(*Digest).Write
     0.33s  0.78% 62.33%      0.33s  0.78%  runtime.futex
     0.27s  0.64% 62.97%      0.76s  1.80%  database/sql.convertAssignRows
     0.25s  0.59% 63.56%      5.33s 12.60%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.24s  0.57% 64.13%      0.42s  0.99%  compress/flate.(*compressor).deflate
     0.23s  0.54% 64.67%      1.93s  4.56%  runtime.mapassign_faststr
     0.22s  0.52% 65.19%      4.45s 10.52%  runtime.mallocgc
     0.21s   0.5% 65.69%      0.45s  1.06%  runtime.greyobject
     0.21s   0.5% 66.19%      0.25s  0.59%  runtime.heapArenaOf (inline)
     0.20s  0.47% 66.66%      1.18s  2.79%  runtime.mallocgcSmallScanNoHeader
     0.20s  0.47% 67.13%      0.23s  0.54%  runtime.spanOf (inline)
     0.19s  0.45% 67.58%      0.23s  0.54%  runtime.typePointers.next
     0.18s  0.43% 68.01%      6.06s 14.33%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.18s  0.43% 68.43%      0.22s  0.52%  encoding/json.appendString[go.shape.string]
     0.17s   0.4% 68.83%      0.36s  0.85%  internal/runtime/maps.(*Iter).Next
     0.17s   0.4% 69.24%      0.45s  1.06%  runtime.scanObjectsSmall
     0.15s  0.35% 69.59%      5.61s 13.27%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.13s  0.31% 69.90%      0.34s   0.8%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.12s  0.28% 70.18%      0.60s  1.42%  encoding/json.structEncoder.encode
     0.12s  0.28% 70.47%      0.26s  0.61%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.11s  0.26% 70.73%      0.52s  1.23%  runtime.mapaccess2_faststr
     0.10s  0.24% 70.96%      0.55s  1.30%  runtime.wbBufFlush1
     0.09s  0.21% 71.18%      1.14s  2.70%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.09s  0.21% 71.39%      0.22s  0.52%  runtime.scanObjectSmall
     0.08s  0.19% 71.58%      0.26s  0.61%  crypto/internal/fips140/hmac.(*HMAC).Reset
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 47.5% / 100% | 32% | 0.5% | 7 | 126 MB | CPU が飽和 |
| isu2 (2 cpu) | 23.4% / 100% | 2% | 0.4% | 5 | 664 MB | run queue が積もっている |
| isu3 (2 cpu) | 22.7% / 99% | 2% | 0.4% | 5 | 670 MB | run queue が積もっている |
| isu4 (2 cpu) | 31.3% / 100% | 5% | 16.6% | 5 | 99 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 73% / nginx 12% / landscape-sysin 1% |
| isu2 | isureal 45% / nginx 1% |
| isu3 | isureal 44% / nginx 1% |
| isu4 | isureal 14% / postgres 13% |

## Diff from previous run

スコア: 132853 → 137455 (+4602)

`6717755..9c97c71`

```
 conf/isu1/home/isucon/env.sh                   |   2 +-
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |   7 +-
 conf/isu2/home/isucon/env.sh                   |   3 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |   7 +-
 conf/isu3/home/isucon/env.sh                   |   3 +
 webapp/go/account_deletion_api.go              |  25 +-
 webapp/go/admin_api.go                         |  15 +-
 webapp/go/admin_posting_window_api.go          |  62 +-
 webapp/go/admin_user_api.go                    |   6 +-
 webapp/go/cmd/isureal/main.go                  |   8 +
 webapp/go/comment_api.go                       |  15 +-
 webapp/go/export_api.go                        |   4 +
 webapp/go/friend_request_api.go                |  48 +-
 webapp/go/image_api.go                         |  17 +-
 webapp/go/image_file.go                        | 132 ++++
 webapp/go/image_local.go                       |  67 +-
 webapp/go/initialize.go                        |  10 +-
 webapp/go/offload.go                           |   2 +-
 webapp/go/origin_store.go                      | 147 ++++
 webapp/go/persist.go                           | 479 +++++++++++++
 webapp/go/post_api.go                          |  41 +-
 webapp/go/post_forward.go                      |  33 +-
 webapp/go/profile_image_api.go                 |  63 +-
 webapp/go/realmoji_api.go                      |  37 +-
 webapp/go/relationship_api.go                  |  68 +-
 webapp/go/repl.go                              | 943 +++++++++++++++++++++++++
 webapp/go/repl_digest.go                       | 151 ++++
 webapp/go/repl_test.go                         |  62 ++
 webapp/go/replica_server.go                    | 266 +++++++
 webapp/go/server.go                            |   7 +-
 webapp/go/store.go                             | 126 ++--
 webapp/go/user_api.go                          |  71 +-
 32 files changed, 2426 insertions(+), 501 deletions(-)
```
