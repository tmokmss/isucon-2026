# run 20261031-151339-5a230a4

## Summary

| | |
|---|---|
| score | 191816 |
| status | pass |
| ref | 5a230a4 |
| sha | 5a230a478ba5 |
| 所要時間 | 140s |
| 前回比 | +54361 (39.5%) vs 20261031-145836-9c97c71 |
| ベスト比 | +17782 (10.2%) vs 20261031-144734-fd0f88e = 174034 **ベスト更新** |
| note | main: +#85 画像をリージョン内で完結（確認） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	146	422	1141	1144
europe	7	20	62	181	472	471
america	6	19	59	173	305	304
```

```
{"id":"21481","status":"done","portalStatus":"finished","passed":true,"score":191816,"raw":191816,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:14:05.851Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:14:08.846Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:14:11.846Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:14:14.873Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T06:14:17.360Z level=INFO msg="user moved" user=toshi.5295 from=asia to=america
time=2026-10-31T06:14:17.846Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T06:14:20.847Z level=INFO msg="america 2 日目: 19 人"
time=2026-10-31T06:14:21.617Z level=INFO msg="user moved" user=maxtor5407 from=europe to=asia
time=2026-10-31T06:14:23.854Z level=INFO msg="asia 3 日目: 146 人"
time=2026-10-31T06:14:25.242Z level=INFO msg="user moved" user=nishihe-5633 from=asia to=europe
time=2026-10-31T06:14:25.320Z level=INFO msg="user moved" user=ulchior5802 from=asia to=america
time=2026-10-31T06:14:26.846Z level=INFO msg="europe 3 日目: 62 人"
time=2026-10-31T06:14:27.315Z level=INFO msg="user moved" user=martoha-6182 from=asia to=america
time=2026-10-31T06:14:27.589Z level=INFO msg="user moved" user=shievaka-6268 from=asia to=america
time=2026-10-31T06:14:28.147Z level=INFO msg="user moved" user=ulchior5802 from=america to=europe
time=2026-10-31T06:14:28.226Z level=INFO msg="user moved" user=shievaka-6268 from=america to=asia
time=2026-10-31T06:14:29.846Z level=INFO msg="america 3 日目: 59 人"
time=2026-10-31T06:14:30.730Z level=INFO msg="user moved" user=solfumax6253 from=europe to=america
time=2026-10-31T06:14:32.856Z level=INFO msg="asia 4 日目: 422 人"
time=2026-10-31T06:14:33.355Z level=INFO msg="user moved" user=danor_6372 from=europe to=asia
time=2026-10-31T06:14:35.012Z level=INFO msg="user moved" user=shievaka-6268 from=asia to=america
time=2026-10-31T06:14:36.235Z level=INFO msg="europe 4 日目: 181 人"
time=2026-10-31T06:14:36.299Z level=INFO msg="user moved" user=tsusolchi7513 from=asia to=europe
time=2026-10-31T06:14:37.163Z level=INFO msg="user moved" user=shimi7258 from=asia to=america
time=2026-10-31T06:14:37.342Z level=INFO msg="user moved" user=toshi.5295 from=america to=asia
time=2026-10-31T06:14:37.936Z level=INFO msg="user moved" user=andan7893 from=asia to=america
time=2026-10-31T06:14:38.145Z level=INFO msg="user moved" user=hame7932 from=asia to=europe
time=2026-10-31T06:14:38.648Z level=INFO msg="user moved" user=rinoa-7024 from=asia to=america
time=2026-10-31T06:14:38.942Z level=INFO msg="america 4 日目: 173 人"
time=2026-10-31T06:14:39.396Z level=INFO msg="user moved" user=yakiya-8712 from=asia to=america
time=2026-10-31T06:14:39.822Z level=INFO msg="user moved" user=toshi.5295 from=asia to=america
time=2026-10-31T06:14:40.500Z level=INFO msg="user moved" user=miadanra.7341 from=asia to=america
time=2026-10-31T06:14:41.122Z level=INFO msg="user moved" user=danor_6372 from=asia to=america
time=2026-10-31T06:14:41.871Z level=INFO msg="asia 5 日目: 1141 人"
time=2026-10-31T06:14:43.518Z level=INFO msg="user moved" user=wamarhe.9627 from=america to=asia
time=2026-10-31T06:14:45.627Z level=INFO msg="user moved" user=noamia.8969 from=europe to=asia
time=2026-10-31T06:14:46.250Z level=INFO msg="europe 5 日目: 472 人"
time=2026-10-31T06:14:47.112Z level=INFO msg="user moved" user=ulchior5802 from=europe to=asia
time=2026-10-31T06:14:48.737Z level=INFO msg="america 5 日目: 305 人"
time=2026-10-31T06:14:57.506Z level=INFO msg="europe 最終日（グランプリ）: 471 人"
time=2026-10-31T06:14:57.615Z level=INFO msg="asia 最終日（グランプリ）: 1144 人"
time=2026-10-31T06:14:58.101Z level=INFO msg="america 最終日（グランプリ）: 304 人"
time=2026-10-31T06:15:05.969Z level=INFO msg="グランプリ: アジア 486 人, ヨーロッパ 200 人, アメリカ 82 人"
{
  "pass": true,
  "score": 191816,
  "messages": [
    "muremax17607: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rudan8595: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "musoku.12046: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "vinor_14935: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kefumia-8673: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nori-7923: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "munu-11540: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mekomax.5473: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "vinsura_17158: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rayuzoe17911: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "suri-17734: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rumu-16581: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "zoetevin_16167: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sakako18097: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ivyka-10408: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yoriya17423: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "bentavin.9574: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "benelre_11288: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maxno.18306: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "neteis_16973: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
78.9%	2929	0	0	GET	/api/notifications	82882.421	28.297	60.002	60.391	107
4.4%	2902	176	0	POST	/api/posts	4631.535	1.596	6.643	64.761	615
3.9%	6474	51	0	POST	/api/friend-requests	4145.055	0.640	2.226	3.476	114
1.9%	7146	581	0	POST	/api/auth/login	2013.830	0.282	1.858	3.444	226
1.8%	20666	348	0	GET	^/api/images/[^/]+	1881.109	0.091	1.060	3.808	123859
1.4%	3008	15	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1521.073	0.506	2.196	3.977	116
1.2%	4869	0	0	GET	/api/memories	1236.392	0.254	2.649	7.852	12483
0.9%	15760	0	0	GET	^/api/feeds/[^/]+	964.732	0.061	0.558	1.991	3178
0.7%	1168	1	0	POST	/api/auth/register	730.257	0.625	1.661	1.816	227
0.7%	777	129	0	PUT	/api/me/profile-image	711.936	0.916	9.401	13.413	3
0.6%	718	7	0	POST	^/api/posts/[^/]+/realmojis$	680.440	0.948	4.348	5.765	474
0.5%	849	3	0	POST	^/api/posts/[^/]+/comments$	553.130	0.652	3.532	4.885	354
0.4%	796	0	0	PUT	^/api/follows/[^/]+$	407.310	0.512	1.771	2.441	0
0.3%	1107	0	0	POST	/api/auth/logout	315.992	0.285	1.931	3.903	0
0.3%	4142	0	0	GET	/api/posting-window/current	307.871	0.074	0.503	1.299	237
0.3%	425	0	0	GET	/api/me/export	304.762	0.717	6.388	9.141	1938065
0.3%	3133	0	0	GET	/api/users	278.723	0.089	0.519	1.967	242
0.3%	558	0	0	PATCH	/api/me	278.632	0.499	1.943	3.095	276
0.2%	2170	19	0	GET	^/api/posts/me/current$	174.436	0.080	0.964	1.427	648
0.2%	235	5	0	PATCH	^/api/posts/[^/]+$	167.260	0.712	2.656	3.481	634
0.1%	1932	0	0	GET	^/api/relationships/[^/]+	125.050	0.065	0.657	1.913	1009
0.1%	2414	40	0	GET	^/api/posts/[^/]+$	123.703	0.051	0.462	1.426	618
0.1%	1044	0	0	GET	^/api/memories/[^/]+	114.345	0.110	1.378	2.685	660
0.1%	110	0	0	DELETE	^/api/posts/[^/]+$	113.798	1.035	3.148	3.436	0
0.1%	661	2	0	GET	/api/me	88.646	0.134	0.488	0.751	302
0.1%	154	0	0	DELETE	^/api/comments/[^/]+$	74.707	0.485	2.783	2.895	0
0.1%	123	0	0	DELETE	^/api/realmojis/[^/]+$	70.297	0.572	2.228	2.492	0
0.0%	1145	0	0	GET	^/api/users/[^/]+$	49.179	0.043	0.366	0.980	203
0.0%	1015	0	0	GET	^/api/posts/[^/]+/realmojis$	43.856	0.043	0.449	1.353	1403
0.0%	906	0	0	GET	^/api/posts/[^/]+/comments$	39.427	0.044	0.376	1.245	2160
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
75.4%	969	1	0	GET	/api/notifications	23124.761	23.865	60.498	60.603	100
5.1%	2295	24	0	POST	/api/friend-requests	1566.751	0.683	2.407	3.507	114
3.7%	990	62	0	POST	/api/posts	1124.491	1.136	6.483	64.263	604
2.4%	1338	0	0	POST	/api/auth/login	742.708	0.555	2.043	4.289	245
1.9%	4235	0	0	GET	^/api/feeds/[^/]+	592.682	0.140	1.553	2.599	2999
1.9%	4082	26	0	GET	^/api/images/[^/]+	578.922	0.142	1.477	2.898	175081
1.2%	470	0	0	POST	/api/auth/register	374.732	0.797	1.913	2.012	229
1.0%	526	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	314.442	0.598	2.240	2.472	116
0.9%	319	0	0	PUT	^/api/follows/[^/]+$	287.396	0.901	2.777	3.820	0
0.9%	249	0	0	PUT	/api/me/profile-image	275.662	1.107	3.575	5.042	0
0.7%	1034	0	0	GET	/api/memories	225.209	0.218	2.717	3.777	12504
0.7%	231	0	0	PATCH	/api/me	214.787	0.930	2.786	2.980	281
0.6%	1498	0	0	GET	/api/posting-window/current	169.252	0.113	0.945	2.331	237
0.5%	154	5	0	POST	^/api/posts/[^/]+/realmojis$	153.911	0.999	3.569	4.015	456
0.5%	713	35	0	GET	^/api/posts/me/current$	148.327	0.208	2.269	2.412	613
0.5%	1131	0	0	GET	/api/users	145.260	0.128	1.585	2.532	241
0.4%	215	3	0	POST	^/api/posts/[^/]+/comments$	128.294	0.597	2.560	2.725	345
0.3%	175	0	0	POST	/api/auth/logout	90.070	0.515	2.580	2.585	0
0.2%	522	0	0	GET	^/api/relationships/[^/]+	73.042	0.140	1.408	2.546	683
0.2%	514	16	0	GET	^/api/posts/[^/]+$	63.944	0.124	1.177	2.260	608
0.2%	74	9	0	PATCH	^/api/posts/[^/]+$	52.672	0.712	3.039	3.039	567
0.1%	33	0	0	DELETE	^/api/posts/[^/]+$	38.540	1.168	3.452	3.452	0
0.1%	297	0	0	GET	^/api/users/[^/]+$	32.325	0.109	0.794	1.151	198
0.1%	257	1	0	GET	/api/me	27.343	0.106	0.531	0.952	302
0.1%	38	0	0	DELETE	^/api/comments/[^/]+$	27.113	0.713	4.127	4.127	0
0.1%	187	0	0	GET	^/api/memories/[^/]+	26.058	0.139	1.900	1.916	932
0.1%	202	0	0	GET	^/api/posts/[^/]+/realmojis$	24.304	0.120	0.823	1.727	1353
0.1%	209	0	0	GET	^/api/posts/[^/]+/comments$	20.977	0.100	0.794	0.825	2347
0.1%	59	0	0	GET	/api/me/export	19.702	0.334	5.596	5.596	2543776
0.0%	15	0	0	DELETE	^/api/realmojis/[^/]+$	7.571	0.505	1.575	1.575	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
72.3%	676	2	0	GET	/api/notifications	13649.282	20.191	60.103	60.590	114
4.8%	702	41	0	POST	/api/posts	915.964	1.305	6.361	63.861	624
4.8%	1468	14	0	POST	/api/friend-requests	901.627	0.614	2.379	3.467	114
3.2%	1107	0	0	POST	/api/auth/login	597.737	0.540	2.267	3.450	249
2.6%	3679	16	0	GET	^/api/images/[^/]+	495.066	0.135	1.194	3.598	142789
2.1%	2715	0	0	GET	^/api/feeds/[^/]+	403.466	0.149	1.755	2.056	3347
1.3%	416	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	251.206	0.604	2.188	2.780	116
1.2%	295	0	0	POST	/api/auth/register	218.019	0.739	1.944	2.010	230
1.0%	181	0	0	PUT	/api/me/profile-image	181.423	1.002	3.550	4.075	0
0.8%	875	0	0	GET	/api/memories	150.914	0.172	2.533	3.398	17580
0.8%	1038	0	0	GET	/api/posting-window/current	143.757	0.138	0.853	1.929	240
0.6%	110	2	0	POST	^/api/posts/[^/]+/realmojis$	119.135	1.083	3.707	3.741	456
0.6%	177	0	0	PUT	^/api/follows/[^/]+$	109.648	0.619	2.047	2.049	0
0.6%	131	0	0	PATCH	/api/me	108.767	0.830	4.242	4.374	283
0.5%	152	0	0	POST	^/api/posts/[^/]+/comments$	97.111	0.639	3.388	4.233	347
0.5%	750	0	0	GET	/api/users	95.460	0.127	1.134	2.361	245
0.5%	164	0	0	POST	/api/auth/logout	88.512	0.540	2.599	3.048	0
0.4%	480	15	0	GET	^/api/posts/me/current$	84.303	0.176	2.092	2.371	643
0.3%	362	0	0	GET	^/api/relationships/[^/]+	51.393	0.142	1.800	1.968	830
0.2%	388	26	0	GET	^/api/posts/[^/]+$	38.688	0.100	0.401	1.844	592
0.2%	48	1	0	PATCH	^/api/posts/[^/]+$	36.024	0.751	1.802	1.802	636
0.1%	201	0	0	GET	^/api/users/[^/]+$	22.057	0.110	0.725	1.207	201
0.1%	187	1	0	GET	/api/me	20.728	0.111	1.449	1.449	303
0.1%	18	0	0	DELETE	^/api/posts/[^/]+$	20.240	1.124	2.498	2.498	0
0.1%	149	0	0	GET	^/api/memories/[^/]+	19.560	0.131	1.591	1.950	1100
0.1%	123	0	0	GET	^/api/posts/[^/]+/comments$	13.528	0.110	1.228	1.232	2142
0.1%	57	0	0	GET	/api/me/export	10.925	0.192	1.736	1.736	3054364
0.1%	26	0	0	DELETE	^/api/comments/[^/]+$	10.568	0.406	1.374	1.374	0
0.1%	105	0	0	GET	^/api/posts/[^/]+/realmojis$	10.372	0.099	0.338	1.142	1563
0.0%	13	0	0	DELETE	^/api/realmojis/[^/]+$	9.142	0.703	1.598	1.598	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 17073 クエリ / 48 種類 / 合計 157.300s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	56.5%	429	88.913	207.26	1257.66	1233.9	126.3	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	13.5%	3	21.268	7089.36	9548.22	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.6%	14642	18.290	1.25	9.87	9.3	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.3%	3	16.278	5425.86	7261.87	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.0%	3	7.911	2637.05	3562.21	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	1.6%	4	2.483	620.75	1972.55	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.2%	4	0.350	87.56	157.38	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.2%	4	0.346	86.41	156.83	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.253	253.31	253.31	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	0.2%	1	0.244	244.41	244.41	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.1%	1	0.229	229.08	229.08	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.152	152.20	152.20	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.115	114.71	114.71	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.082	82.13	82.13	2412.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.073	73.01	73.01	1503.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
Build ID: 5a17d239a328636b742bcc054004140a756dcf2d
Type: cpu
Time: 2026-10-31 15:13:46 JST
Duration: 60s, Total samples = 47.09s (78.48%)
Showing nodes accounting for 33.97s, 72.14% of 47.09s total
Dropped 1108 nodes (cum <= 0.24s)
Showing top 40 nodes out of 302
      flat  flat%   sum%        cum   cum%
     5.70s 12.10% 12.10%      5.70s 12.10%  crypto/internal/fips140/sha256.blockAVX2
     5.47s 11.62% 23.72%      5.47s 11.62%  internal/runtime/syscall/linux.Syscall6
     2.42s  5.14% 28.86%      2.42s  5.14%  runtime.memmove
     2.10s  4.46% 33.32%      2.66s  5.65%  runtime.findObject
     1.87s  3.97% 37.29%      7.61s 16.16%  runtime.scanObject
     1.77s  3.76% 41.05%      3.02s  6.41%  runtime.tryDeferToSpanScan
     1.69s  3.59% 44.64%      5.73s 12.17%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.67s  3.55% 48.18%      1.67s  3.55%  runtime.memclrNoHeapPointers
     1.46s  3.10% 51.28%      1.46s  3.10%  aeshashbody
     1.33s  2.82% 54.11%      1.33s  2.82%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.98s  2.08% 56.19%      0.98s  2.08%  runtime.spanClass.sizeclass (inline)
     0.61s  1.30% 57.49%      2.32s  4.93%  runtime.mapaccess1_faststr
     0.47s     1% 58.48%      0.47s     1%  memeqbody
     0.46s  0.98% 59.46%      0.46s  0.98%  runtime.nextFreeFast (inline)
     0.36s  0.76% 60.23%      6.40s 13.59%  crypto/internal/fips140/sha256.(*Digest).Write
     0.34s  0.72% 60.95%      6.51s 13.82%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.33s   0.7% 61.65%      0.51s  1.08%  example.com/isureal-company/isureal.postLess
     0.30s  0.64% 62.28%      0.64s  1.36%  database/sql.convertAssignRows
     0.30s  0.64% 62.92%      0.66s  1.40%  runtime.greyobject
     0.30s  0.64% 63.56%      0.31s  0.66%  runtime.spanOf (inline)
     0.29s  0.62% 64.17%      6.91s 14.67%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.29s  0.62% 64.79%      4.56s  9.68%  runtime.mallocgc
     0.29s  0.62% 65.41%      1.22s  2.59%  runtime.mallocgcSmallScanNoHeader
     0.28s  0.59% 66.00%      0.28s  0.59%  runtime.futex
     0.28s  0.59% 66.60%      0.28s  0.59%  sync/atomic.(*Int32).Add
     0.26s  0.55% 67.15%      0.34s  0.72%  runtime.(*mspan).markBitsForIndex (inline)
     0.24s  0.51% 67.66%      0.25s  0.53%  runtime.(*mspan).writeHeapBitsSmall
     0.22s  0.47% 68.12%      0.41s  0.87%  internal/runtime/maps.(*Iter).Next
     0.21s  0.45% 68.57%      0.48s  1.02%  compress/flate.(*compressor).deflate
     0.21s  0.45% 69.02%      1.88s  3.99%  runtime.mapassign_faststr
     0.21s  0.45% 69.46%      0.61s  1.30%  runtime.scanObjectsSmall
     0.21s  0.45% 69.91%      0.24s  0.51%  runtime.typePointers.next
     0.19s   0.4% 70.31%      0.26s  0.55%  encoding/json.appendString[go.shape.string]
     0.15s  0.32% 70.63%      7.40s 15.71%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.14s   0.3% 70.93%      0.81s  1.72%  encoding/json.structEncoder.encode
     0.13s  0.28% 71.20%      0.75s  1.59%  runtime.wbBufFlush1
     0.12s  0.25% 71.46%      3.13s  6.65%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s  0.23% 71.69%      7.90s 16.78%  example.com/isureal-company/isureal.computeImageSignature
     0.11s  0.23% 71.93%      0.38s  0.81%  runtime.mapaccess2_faststr
     0.10s  0.21% 72.14%      5.80s 12.32%  crypto/internal/fips140/sha256.block
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 42.4% / 100% | 27% | 0.6% | 8 | 132 MB | run queue が積もっている |
| isu2 (2 cpu) | 27.0% / 100% | 5% | 1.2% | 4 | 463 MB |  |
| isu3 (2 cpu) | 24.5% / 100% | 6% | 0.7% | 5 | 144 MB | run queue が積もっている |
| isu4 (2 cpu) | 36.4% / 100% | 16% | 13.3% | 18 | 103 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 69% / nginx 12% |
| isu2 | isureal 49% / nginx 2% |
| isu3 | isureal 46% / nginx 2% |
| isu4 | isureal 30% / postgres 14% |

## Diff from previous run

スコア: 137455 → 191816 (+54361)

`9c97c71..5a230a4`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf |   2 +
 conf/isu1/home/isucon/env.sh                   |   2 +-
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |   4 +
 conf/isu2/home/isucon/env.sh                   |   1 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |   4 +
 conf/isu3/home/isucon/env.sh                   |   1 +
 webapp/go/export_api.go                        |   4 +-
 webapp/go/image_file.go                        |  11 +-
 webapp/go/image_held.go                        | 172 +++++++++++++++++++++++++
 webapp/go/image_local.go                       |  20 +--
 webapp/go/offload.go                           |   2 +-
 webapp/go/origin_store.go                      |  59 +++++----
 webapp/go/post_api.go                          |  20 ++-
 webapp/go/post_forward.go                      |  47 ++-----
 webapp/go/repl.go                              |  17 ++-
 webapp/go/replica_server.go                    |  36 ++++--
 webapp/go/server.go                            |   1 +
 webapp/go/store.go                             |   1 +
 18 files changed, 309 insertions(+), 95 deletions(-)
```
