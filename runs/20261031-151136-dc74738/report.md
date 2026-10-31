# run 20261031-151136-dc74738

## Summary

| | |
|---|---|
| score | 170298 |
| status | pass |
| ref | lane2-tune |
| sha | dc74738f8010 |
| 所要時間 | 118s |
| 前回比 | +170298 (—) vs 20261031-150350-0343458 |
| ベスト比 | -3736 (-2.1%) vs 20261031-144734-fd0f88e = 174034 |
| note | ipban bg reload + image cache 100MB + GOMEMLIMIT 2200 on 5a230a4 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	141	402	1105	1106
europe	7	20	59	176	475	474
america	6	15	41	115	183	183
```

```
{"id":"21291","status":"done","portalStatus":"finished","passed":true,"score":170298,"raw":170298,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:11:58.544Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:12:01.537Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:12:04.537Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:12:07.546Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T06:12:08.548Z level=INFO msg="user moved" user=shito5208 from=asia to=europe
time=2026-10-31T06:12:10.162Z level=INFO msg="user moved" user=shito5208 from=europe to=america
time=2026-10-31T06:12:10.538Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T06:12:13.538Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T06:12:16.548Z level=INFO msg="asia 3 日目: 141 人"
time=2026-10-31T06:12:17.428Z level=INFO msg="user moved" user=meis_5561 from=asia to=europe
time=2026-10-31T06:12:18.382Z level=INFO msg="user moved" user=meis_5561 from=europe to=america
time=2026-10-31T06:12:19.538Z level=INFO msg="europe 3 日目: 59 人"
time=2026-10-31T06:12:20.221Z level=INFO msg="user moved" user=yotor.6152 from=asia to=america
time=2026-10-31T06:12:20.851Z level=INFO msg="user moved" user=yotor.6152 from=america to=europe
time=2026-10-31T06:12:22.542Z level=INFO msg="america 3 日目: 41 人"
time=2026-10-31T06:12:23.621Z level=INFO msg="user moved" user=shito5208 from=america to=asia
time=2026-10-31T06:12:25.550Z level=INFO msg="asia 4 日目: 402 人"
time=2026-10-31T06:12:25.689Z level=INFO msg="user moved" user=hitoya6426 from=america to=asia
time=2026-10-31T06:12:26.671Z level=INFO msg="user moved" user=shito5208 from=asia to=europe
time=2026-10-31T06:12:27.509Z level=INFO msg="user moved" user=maxsu-6656 from=asia to=america
time=2026-10-31T06:12:28.653Z level=INFO msg="europe 4 日目: 176 人"
time=2026-10-31T06:12:29.588Z level=INFO msg="user moved" user=fuis-7164 from=asia to=america
time=2026-10-31T06:12:29.905Z level=INFO msg="user moved" user=tamax_6840 from=asia to=europe
time=2026-10-31T06:12:30.011Z level=INFO msg="user moved" user=hitoya6426 from=asia to=europe
time=2026-10-31T06:12:30.202Z level=INFO msg="user moved" user=maxsu-6656 from=america to=asia
time=2026-10-31T06:12:31.392Z level=INFO msg="user moved" user=ankichi.8241 from=asia to=europe
time=2026-10-31T06:12:32.237Z level=INFO msg="america 4 日目: 115 人"
time=2026-10-31T06:12:32.552Z level=INFO msg="user moved" user=rivin.7077 from=asia to=america
time=2026-10-31T06:12:32.615Z level=INFO msg="user moved" user=hoel-7759 from=asia to=america
time=2026-10-31T06:12:32.682Z level=INFO msg="user moved" user=yotor.6152 from=europe to=america
time=2026-10-31T06:12:32.764Z level=INFO msg="user moved" user=wasasa8150 from=asia to=europe
time=2026-10-31T06:12:32.888Z level=INFO msg="user moved" user=elka.8566 from=europe to=asia
time=2026-10-31T06:12:33.943Z level=INFO msg="user moved" user=katesu-8462 from=asia to=europe
time=2026-10-31T06:12:34.136Z level=INFO msg="user moved" user=meis_5561 from=america to=europe
time=2026-10-31T06:12:34.630Z level=INFO msg="asia 5 日目: 1105 人"
time=2026-10-31T06:12:36.506Z level=INFO msg="user moved" user=chizoenoa8733 from=europe to=asia
time=2026-10-31T06:12:38.263Z level=INFO msg="europe 5 日目: 475 人"
time=2026-10-31T06:12:38.497Z level=INFO msg="user moved" user=yata_9391 from=america to=asia
time=2026-10-31T06:12:41.330Z level=INFO msg="america 5 日目: 183 人"
time=2026-10-31T06:12:50.482Z level=INFO msg="europe 最終日（グランプリ）: 474 人"
time=2026-10-31T06:12:51.126Z level=INFO msg="america 最終日（グランプリ）: 183 人"
time=2026-10-31T06:12:52.774Z level=INFO msg="asia 最終日（グランプリ）: 1106 人"
time=2026-10-31T06:12:58.624Z level=INFO msg="グランプリ: アジア 277 人, ヨーロッパ 295 人, アメリカ 90 人"
{
  "pass": true,
  "score": 170298,
  "messages": [
    "wari-5034: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "yahita.5103: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "releone.5125: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "nasolso.16925: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "vinte_17264: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tsuhi_17024: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chiheno_15678: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rureha-17180: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "doyun_cho_4194: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "wachi10813: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "marben_5056: 友達申請を承認できません: POST /api/friend-requests/c1222f1f-e007-4509-a6a2-793c316ba063/accept: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests/c1222f1f-e007-4509-a6a2-793c316ba063/accept\": EOF",
    "solhiri13518: 友達申請を送れません: POST /api/friend-requests: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.6%	1762	1	0	GET	/api/notifications	18212.321	10.336	18.227	18.282	171
10.7%	6263	84	0	POST	/api/friend-requests	4273.047	0.682	2.638	4.709	114
8.9%	2779	118	0	POST	/api/posts	3539.201	1.274	4.515	7.722	626
6.5%	18785	318	0	GET	^/api/images/[^/]+	2601.554	0.138	1.731	3.567	121903
4.8%	6528	566	0	POST	/api/auth/login	1920.401	0.294	2.122	5.026	224
3.3%	2684	8	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1312.818	0.489	2.447	3.577	116
3.2%	14703	0	0	GET	^/api/feeds/[^/]+	1270.292	0.086	0.537	1.364	3088
2.4%	4245	0	0	GET	/api/memories	944.014	0.222	1.801	6.053	12779
1.9%	748	127	0	PUT	/api/me/profile-image	746.224	0.998	7.189	9.120	3
1.6%	616	3	0	POST	^/api/posts/[^/]+/realmojis$	631.913	1.026	4.518	6.153	473
1.5%	1130	1	0	POST	/api/auth/register	596.941	0.528	1.825	1.864	227
1.1%	3996	0	0	GET	/api/posting-window/current	433.099	0.108	0.797	1.358	237
1.0%	538	3	0	POST	^/api/posts/[^/]+/comments$	411.965	0.766	3.611	4.692	357
1.0%	721	0	0	PUT	^/api/follows/[^/]+$	388.012	0.538	1.961	2.304	0
0.9%	3023	0	0	GET	/api/users	368.258	0.122	0.756	1.376	255
0.8%	566	0	0	PATCH	/api/me	335.578	0.593	2.104	2.382	277
0.7%	983	0	0	POST	/api/auth/logout	281.746	0.287	2.281	3.130	0
0.6%	370	0	0	GET	/api/me/export	251.346	0.679	3.321	6.825	1947766
0.5%	1991	14	0	GET	^/api/posts/me/current$	199.949	0.100	0.830	1.014	647
0.5%	211	2	0	PATCH	^/api/posts/[^/]+$	184.990	0.877	3.508	3.909	644
0.5%	2189	60	0	GET	^/api/posts/[^/]+$	182.780	0.083	0.590	0.965	625
0.3%	106	0	0	DELETE	^/api/posts/[^/]+$	131.846	1.244	3.115	3.144	0
0.3%	1812	0	0	GET	^/api/relationships/[^/]+	129.505	0.071	0.512	0.964	936
0.3%	635	3	0	GET	/api/me	111.502	0.176	0.802	0.975	301
0.3%	942	0	0	GET	^/api/memories/[^/]+	107.353	0.114	0.998	1.893	716
0.3%	120	0	0	DELETE	^/api/realmojis/[^/]+$	101.563	0.846	3.082	3.801	0
0.2%	1077	1	0	GET	^/api/users/[^/]+$	82.606	0.077	0.478	1.383	218
0.2%	121	0	0	DELETE	^/api/comments/[^/]+$	62.671	0.518	2.312	2.487	0
0.1%	830	0	0	GET	^/api/posts/[^/]+/realmojis$	54.233	0.065	0.487	1.015	2016
0.1%	676	0	0	GET	^/api/posts/[^/]+/comments$	37.151	0.055	0.340	0.842	2225
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.1%	628	2	0	GET	/api/notifications	4833.464	7.697	12.346	12.354	159
15.1%	2222	26	0	POST	/api/friend-requests	1973.033	0.888	2.809	4.257	114
7.8%	998	28	0	POST	/api/posts	1013.173	1.015	4.468	5.533	634
5.5%	1294	0	0	POST	/api/auth/login	714.723	0.552	2.377	3.281	244
5.1%	4088	22	0	GET	^/api/images/[^/]+	664.259	0.162	1.628	5.421	185967
4.6%	3942	0	0	GET	^/api/feeds/[^/]+	603.911	0.153	1.149	1.774	2581
2.9%	545	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	377.590	0.693	2.602	3.438	116
2.8%	468	0	0	POST	/api/auth/register	370.897	0.793	1.981	2.054	229
2.6%	311	0	0	PUT	^/api/follows/[^/]+$	334.489	1.076	3.019	3.600	0
2.1%	1502	0	0	GET	/api/posting-window/current	269.414	0.179	0.989	1.293	239
1.9%	267	0	0	PUT	/api/me/profile-image	252.741	0.947	2.978	3.571	0
1.8%	1169	0	0	GET	/api/memories	234.091	0.200	1.338	1.601	11856
1.8%	225	0	0	PATCH	/api/me	231.489	1.029	3.820	4.185	283
1.2%	1066	0	0	GET	/api/users	151.293	0.142	1.027	1.534	258
1.1%	146	1	0	POST	^/api/posts/[^/]+/comments$	143.476	0.983	3.891	4.775	341
1.1%	121	4	0	POST	^/api/posts/[^/]+/realmojis$	142.512	1.178	4.908	5.517	474
0.9%	837	43	0	GET	^/api/posts/me/current$	116.586	0.139	1.166	1.674	620
0.8%	564	0	0	GET	^/api/relationships/[^/]+	106.528	0.189	1.293	1.580	839
0.7%	184	0	0	POST	/api/auth/logout	90.570	0.492	3.105	4.560	0
0.5%	92	6	0	PATCH	^/api/posts/[^/]+$	69.649	0.757	2.285	2.285	613
0.5%	531	19	0	GET	^/api/posts/[^/]+$	66.399	0.125	1.012	1.281	618
0.4%	274	2	0	GET	/api/me	46.138	0.168	1.081	1.431	302
0.3%	34	0	0	DELETE	^/api/posts/[^/]+$	44.982	1.323	2.879	2.879	0
0.3%	245	0	0	GET	^/api/users/[^/]+$	36.959	0.151	1.086	1.636	211
0.2%	214	0	0	GET	^/api/posts/[^/]+/realmojis$	26.165	0.122	0.942	1.640	1796
0.2%	185	0	0	GET	^/api/posts/[^/]+/comments$	21.276	0.115	0.870	1.169	2028
0.2%	62	0	0	GET	/api/me/export	21.172	0.341	4.184	4.184	1754514
0.1%	179	0	0	GET	^/api/memories/[^/]+	18.727	0.105	0.571	0.825	905
0.1%	17	0	0	DELETE	^/api/realmojis/[^/]+$	17.003	1.000	2.987	2.987	0
0.1%	22	0	0	DELETE	^/api/comments/[^/]+$	16.324	0.742	1.697	1.697	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.8%	310	1	0	GET	/api/notifications	2379.234	7.675	10.357	10.478	159
11.9%	953	8	0	POST	/api/friend-requests	692.682	0.727	2.540	3.693	114
7.8%	738	0	0	POST	/api/auth/login	455.387	0.617	2.594	3.812	251
6.7%	425	11	0	POST	/api/posts	390.936	0.920	2.837	4.221	635
6.5%	2443	15	0	GET	^/api/images/[^/]+	377.192	0.154	1.402	4.834	133520
4.4%	360	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	257.651	0.716	2.452	3.253	115
3.7%	1575	0	0	GET	^/api/feeds/[^/]+	213.837	0.136	1.094	1.378	2164
2.3%	180	0	0	POST	/api/auth/register	133.463	0.741	2.510	2.544	230
2.0%	98	0	0	PUT	/api/me/profile-image	115.256	1.176	3.075	3.075	0
1.8%	618	0	0	GET	/api/memories	106.300	0.172	1.461	1.650	15946
1.5%	114	0	0	PUT	^/api/follows/[^/]+$	87.924	0.771	2.358	3.407	0
1.4%	630	0	0	GET	/api/posting-window/current	84.263	0.134	0.944	1.446	240
1.3%	121	0	0	POST	/api/auth/logout	77.034	0.637	2.290	2.441	0
1.2%	88	0	0	PATCH	/api/me	68.878	0.783	3.262	3.262	282
1.0%	347	12	0	GET	^/api/posts/me/current$	58.100	0.167	1.008	1.111	627
0.9%	455	0	0	GET	/api/users	54.617	0.120	0.785	1.064	254
0.8%	53	1	0	POST	^/api/posts/[^/]+/realmojis$	43.739	0.825	3.284	3.284	459
0.7%	34	2	0	POST	^/api/posts/[^/]+/comments$	40.522	1.192	4.233	4.233	340
0.6%	228	0	0	GET	^/api/relationships/[^/]+	33.396	0.146	1.148	1.324	1215
0.5%	17	0	0	DELETE	^/api/posts/[^/]+$	29.299	1.723	3.517	3.517	0
0.5%	32	1	0	PATCH	^/api/posts/[^/]+$	27.157	0.849	2.444	2.444	623
0.4%	43	0	0	GET	/api/me/export	22.581	0.525	9.486	9.486	3045394
0.3%	188	12	0	GET	^/api/posts/[^/]+$	19.731	0.105	0.815	1.079	585
0.2%	109	0	0	GET	^/api/users/[^/]+$	11.088	0.102	0.649	0.703	214
0.2%	102	0	0	GET	^/api/memories/[^/]+	10.893	0.107	0.522	0.716	994
0.2%	103	1	0	GET	/api/me	9.034	0.088	0.233	0.489	300
0.1%	69	1	0	GET	^/api/posts/[^/]+/realmojis$	7.297	0.106	0.748	0.748	1599
0.1%	8	0	0	DELETE	^/api/realmojis/[^/]+$	6.390	0.799	1.794	1.794	0
0.1%	52	0	0	GET	^/api/posts/[^/]+/comments$	4.732	0.091	0.296	0.296	2464
0.1%	4	0	0	PUT	^/api/blocks/[^/]+$	4.008	1.002	2.046	2.046	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16793 クエリ / 48 種類 / 合計 144.198s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	54.8%	385	79.005	205.21	753.01	1233.2	126.2	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	15.2%	3	21.866	7288.57	10239.30	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.4%	13265	16.484	1.24	9.80	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	11.0%	3	15.876	5292.02	7175.28	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.3%	3	7.706	2568.83	3518.86	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	0.6%	5	0.855	170.93	186.55	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	5	0.538	107.66	156.89	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.4%	5	0.523	104.60	156.30	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.247	247.24	247.24	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.1%	1	0.195	194.74	194.74	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.1%	1	0.193	192.83	192.83	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.156	155.64	155.64	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.105	105.24	105.24	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.0%	1	0.071	71.37	71.37	1519.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations
15	0.0%	1	0.070	69.90	69.90	2465.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…

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
Build ID: 3edc35153df61dc75d6e68c305326a7a03ff01ad
Type: cpu
Time: 2026-10-31 15:11:41 JST
Duration: 60.26s, Total samples = 49.14s (81.54%)
Showing nodes accounting for 34.94s, 71.10% of 49.14s total
Dropped 1194 nodes (cum <= 0.25s)
Showing top 40 nodes out of 288
      flat  flat%   sum%        cum   cum%
     5.99s 12.19% 12.19%         6s 12.21%  crypto/internal/fips140/sha256.blockAVX2
     5.80s 11.80% 23.99%      5.80s 11.80%  internal/runtime/syscall/linux.Syscall6
     2.20s  4.48% 28.47%      2.70s  5.49%  runtime.findObject
     2.08s  4.23% 32.70%      8.21s 16.71%  runtime.scanObject
     2.04s  4.15% 36.85%      2.04s  4.15%  runtime.memmove
     1.89s  3.85% 40.70%      1.89s  3.85%  runtime.memclrNoHeapPointers
     1.71s  3.48% 44.18%      1.71s  3.48%  aeshashbody
     1.63s  3.32% 47.50%      3.04s  6.19%  runtime.tryDeferToSpanScan
     1.57s  3.19% 50.69%      5.70s 11.60%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.38s  2.81% 53.50%      1.38s  2.81%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.02s  2.08% 55.58%      1.02s  2.08%  runtime.spanClass.sizeclass (inline)
     0.58s  1.18% 56.76%      2.36s  4.80%  runtime.mapaccess1_faststr
     0.44s   0.9% 57.65%      0.78s  1.59%  runtime.greyobject
     0.40s  0.81% 58.47%      6.69s 13.61%  crypto/internal/fips140/sha256.(*Digest).Write
     0.40s  0.81% 59.28%      0.55s  1.12%  example.com/isureal-company/isureal.postLess
     0.40s  0.81% 60.09%      0.40s  0.81%  memeqbody
     0.36s  0.73% 60.83%      6.75s 13.74%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.34s  0.69% 61.52%      1.40s  2.85%  runtime.mallocgcSmallScanNoHeader
     0.34s  0.69% 62.21%      0.34s  0.69%  runtime.nextFreeFast (inline)
     0.32s  0.65% 62.86%      0.32s  0.65%  indexbody
     0.29s  0.59% 63.45%      0.49s     1%  compress/flate.(*compressor).deflate
     0.26s  0.53% 63.98%      0.26s  0.53%  runtime.typePointers.nextFast (inline)
     0.25s  0.51% 64.49%      0.25s  0.51%  runtime.futex
     0.25s  0.51% 65.00%      2.29s  4.66%  runtime.mapassign_faststr
     0.23s  0.47% 65.47%      0.29s  0.59%  runtime.(*mspan).markBitsForIndex (inline)
     0.23s  0.47% 65.93%      0.25s  0.51%  runtime.spanOf (inline)
     0.22s  0.45% 66.38%      7.68s 15.63%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.21s  0.43% 66.81%      0.53s  1.08%  database/sql.convertAssignRows
     0.20s  0.41% 67.22%      0.80s  1.63%  runtime.wbBufFlush1
     0.19s  0.39% 67.60%      7.05s 14.35%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.19s  0.39% 67.99%      0.57s  1.16%  internal/stringslite.Index
     0.19s  0.39% 68.38%      0.25s  0.51%  runtime.heapArenaOf (inline)
     0.19s  0.39% 68.76%      0.67s  1.36%  runtime.mallocgcSmallNoscan
     0.18s  0.37% 69.13%      0.66s  1.34%  runtime.mapaccess2_faststr
     0.18s  0.37% 69.50%      0.55s  1.12%  runtime.scanObjectsSmall
     0.17s  0.35% 69.84%      0.34s  0.69%  internal/runtime/maps.(*Iter).Next
     0.17s  0.35% 70.19%      0.37s  0.75%  runtime.scanObjectSmall
     0.16s  0.33% 70.51%      5.66s 11.52%  runtime.mallocgc
     0.15s  0.31% 70.82%      0.74s  1.51%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
     0.14s  0.28% 71.10%      0.46s  0.94%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 45.5% / 100% | 29% | 0.6% | 7 | 153 MB | run queue が積もっている |
| isu2 (2 cpu) | 30.8% / 100% | 6% | 0.7% | 5 | 179 MB | run queue が積もっている |
| isu3 (2 cpu) | 24.3% / 90% | 1% | 0.6% | 6 | 439 MB | run queue が積もっている |
| isu4 (2 cpu) | 41.1% / 100% | 14% | 16.3% | 8 | 104 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 71% / nginx 13% |
| isu2 | isureal 57% / nginx 2% |
| isu3 | isureal 47% / nginx 1% |
| isu4 | isureal 33% / postgres 15% |

## Diff from previous run

スコア: 0 → 170298 (+170298)

`0343458..dc74738`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf |   2 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |   4 +
 conf/isu2/home/isucon/env.sh                   |   1 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |   4 +
 conf/isu3/home/isucon/env.sh                   |   1 +
 webapp/go/export_api.go                        |   4 +-
 webapp/go/image_file.go                        |  11 +-
 webapp/go/image_held.go                        | 172 +++++++++++++++++++++++++
 webapp/go/image_local.go                       |  20 +--
 webapp/go/image_signature.go                   |   6 +-
 webapp/go/origin_store.go                      |  59 +++++----
 webapp/go/post_api.go                          |  20 ++-
 webapp/go/post_forward.go                      |  14 +-
 webapp/go/repl.go                              |  17 ++-
 webapp/go/replica_server.go                    |  36 ++++--
 webapp/go/server.go                            |   1 +
 webapp/go/store.go                             |   1 +
 17 files changed, 308 insertions(+), 65 deletions(-)
```
