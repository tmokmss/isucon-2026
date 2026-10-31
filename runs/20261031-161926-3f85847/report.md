# run 20261031-161926-3f85847

## Summary

| | |
|---|---|
| score | 202575 |
| status | pass |
| ref | asia2-offload-to-asia1 |
| sha | 3f85847efd18 |
| 所要時間 | 199s |
| 前回比 | +8395 (4.3%) vs 20261031-161619-3f85847 |
| ベスト比 | +1078 (0.5%) vs 20261031-160409-40e65ef = 201497 **ベスト更新** |
| note | asia2 signing offload to europe1/america1 (2nd) |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	150	439	1028	1029
europe	7	21	64	195	540	540
america	6	21	67	191	429	427
```

```
{"id":"27631","status":"done","portalStatus":"finished","passed":true,"score":202575,"raw":202575,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:20:53.433Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:20:56.421Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:20:59.421Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:21:02.435Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T07:21:05.421Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T07:21:05.552Z level=INFO msg="user moved" user=homia5310 from=asia to=america
time=2026-10-31T07:21:08.422Z level=INFO msg="america 2 日目: 21 人"
time=2026-10-31T07:21:08.437Z level=INFO msg="user moved" user=homia5310 from=america to=europe
time=2026-10-31T07:21:11.433Z level=INFO msg="asia 3 日目: 150 人"
time=2026-10-31T07:21:14.318Z level=INFO msg="user moved" user=maelben.5917 from=asia to=america
time=2026-10-31T07:21:14.423Z level=INFO msg="europe 3 日目: 64 人"
time=2026-10-31T07:21:14.762Z level=INFO msg="user moved" user=marma-6187 from=asia to=america
time=2026-10-31T07:21:15.402Z level=INFO msg="user moved" user=marma-6187 from=america to=europe
time=2026-10-31T07:21:15.862Z level=INFO msg="user moved" user=keho-6288 from=asia to=america
time=2026-10-31T07:21:17.310Z level=INFO msg="user moved" user=maelben.5917 from=america to=europe
time=2026-10-31T07:21:17.422Z level=INFO msg="america 3 日目: 67 人"
time=2026-10-31T07:21:19.604Z level=INFO msg="user moved" user=homia5310 from=europe to=asia
time=2026-10-31T07:21:20.432Z level=INFO msg="asia 4 日目: 439 人"
time=2026-10-31T07:21:21.663Z level=INFO msg="user moved" user=teno-6960 from=asia to=europe
time=2026-10-31T07:21:22.285Z level=INFO msg="user moved" user=ronoais-7046 from=asia to=america
time=2026-10-31T07:21:22.532Z level=INFO msg="user moved" user=teno-6960 from=europe to=america
time=2026-10-31T07:21:23.118Z level=INFO msg="user moved" user=ronoais-7046 from=america to=europe
time=2026-10-31T07:21:23.421Z level=INFO msg="europe 4 日目: 195 人"
time=2026-10-31T07:21:24.372Z level=INFO msg="user moved" user=tsuhaul.6776 from=america to=asia
time=2026-10-31T07:21:25.622Z level=INFO msg="user moved" user=himo-6684 from=america to=europe
time=2026-10-31T07:21:26.812Z level=INFO msg="america 4 日目: 191 人"
time=2026-10-31T07:21:26.827Z level=INFO msg="user moved" user=marma-6187 from=europe to=america
time=2026-10-31T07:21:26.843Z level=INFO msg="user moved" user=orvinho.8054 from=asia to=europe
time=2026-10-31T07:21:27.571Z level=INFO msg="user moved" user=niho.8172 from=asia to=america
time=2026-10-31T07:21:28.559Z level=INFO msg="user moved" user=homia5310 from=asia to=america
time=2026-10-31T07:21:29.129Z level=INFO msg="user moved" user=fureka.9527 from=europe to=america
time=2026-10-31T07:21:29.437Z level=INFO msg="asia 5 日目: 1028 人"
time=2026-10-31T07:21:31.085Z level=INFO msg="user moved" user=maelben.5917 from=europe to=asia
time=2026-10-31T07:21:32.474Z level=INFO msg="europe 5 日目: 540 人"
time=2026-10-31T07:21:36.979Z level=INFO msg="america 5 日目: 429 人"
time=2026-10-31T07:21:37.969Z level=INFO msg="user moved" user=naivy-10282 from=america to=asia
time=2026-10-31T07:21:45.074Z level=INFO msg="asia 最終日（グランプリ）: 1029 人"
time=2026-10-31T07:21:45.551Z level=INFO msg="america 最終日（グランプリ）: 427 人"
time=2026-10-31T07:21:45.852Z level=INFO msg="europe 最終日（グランプリ）: 540 人"
time=2026-10-31T07:21:53.545Z level=INFO msg="グランプリ: アジア 458 人, ヨーロッパ 152 人, アメリカ 154 人"
{
  "pass": true,
  "score": 202575,
  "messages": [
    "mochinoa-5045: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "susete-13863: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kite.14573: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "seleo_17338: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sabenivy.18277: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "elni.18460: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hezoe18111: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "benne.8152: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chimi-18525: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tohake_16324: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ratonu16447: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "wamaya17773: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kikimax-18564: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ansu-10387: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "roto-6244: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hiyuso_17917: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rivin.11035: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kichi13946: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noane.7377: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kose12252: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
51.2%	1789	1	0	GET	/api/notifications	20493.311	11.455	60.001	60.006	165
9.8%	2640	185	0	POST	/api/posts	3932.441	1.490	7.773	8.675	610
6.0%	13031	0	0	GET	^/api/feeds/[^/]+	2404.654	0.185	1.524	2.356	4664
5.0%	4217	0	0	GET	/api/memories	2004.626	0.475	3.045	6.602	12790
4.6%	5912	91	0	POST	/api/friend-requests	1853.158	0.313	1.705	2.910	114
2.9%	17794	263	0	GET	^/api/images/[^/]+	1175.040	0.066	0.741	1.882	128787
2.3%	3801	0	0	GET	/api/posting-window/current	925.711	0.244	1.794	2.445	236
2.3%	2838	0	0	GET	/api/users	904.628	0.319	1.876	2.297	239
2.2%	730	126	0	PUT	/api/me/profile-image	864.448	1.184	12.573	15.843	3
2.0%	3588	9	0	POST	^/api/friend-requests/[^/]+/[^/]+$	787.129	0.219	1.513	3.351	116
1.8%	632	57	0	POST	^/api/posts/[^/]+/realmojis$	704.177	1.114	8.153	8.892	424
1.7%	6139	492	0	POST	/api/auth/login	695.921	0.113	1.452	2.393	226
1.5%	2005	17	0	GET	^/api/posts/me/current$	589.784	0.294	1.546	2.308	650
1.0%	358	0	0	GET	/api/me/export	388.741	1.086	5.364	7.314	1898631
0.9%	618	3	0	GET	/api/me	348.151	0.563	1.907	2.276	301
0.7%	665	5	0	POST	^/api/posts/[^/]+/comments$	282.733	0.425	1.966	3.006	339
0.6%	1564	0	0	GET	^/api/relationships/[^/]+	251.896	0.161	1.401	2.037	1354
0.6%	2105	61	0	GET	^/api/posts/[^/]+$	232.521	0.110	1.008	1.407	616
0.5%	914	0	0	GET	^/api/memories/[^/]+	192.836	0.211	1.717	2.335	631
0.5%	952	0	0	POST	/api/auth/logout	183.105	0.192	1.629	2.052	0
0.3%	245	4	0	PATCH	^/api/posts/[^/]+$	126.367	0.516	2.164	2.630	640
0.3%	891	1	0	GET	^/api/posts/[^/]+/realmojis$	124.987	0.140	1.173	1.483	1420
0.3%	639	0	0	PUT	^/api/follows/[^/]+$	108.844	0.170	1.586	2.307	0
0.3%	987	0	0	GET	^/api/users/[^/]+$	105.573	0.107	0.848	1.423	204
0.2%	718	0	0	GET	^/api/posts/[^/]+/comments$	90.933	0.127	1.302	1.828	2571
0.2%	515	0	0	PATCH	/api/me	85.529	0.166	1.620	2.194	275
0.2%	86	0	0	DELETE	^/api/posts/[^/]+$	74.564	0.867	2.171	2.171	0
0.1%	102	0	0	DELETE	^/api/realmojis/[^/]+$	33.012	0.324	1.307	1.707	0
0.1%	111	0	0	DELETE	^/api/comments/[^/]+$	27.478	0.248	1.442	1.496	0
0.1%	1	0	0	POST	/api/initialize	20.851	20.851	20.851	20.851	90
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.1%	809	1	0	GET	/api/notifications	6124.211	7.570	13.186	13.195	163
9.5%	2763	26	0	POST	/api/friend-requests	1374.236	0.497	1.886	3.041	114
9.2%	1209	98	0	POST	/api/posts	1343.383	1.111	6.820	7.250	601
5.5%	5575	0	0	GET	^/api/feeds/[^/]+	795.879	0.143	1.808	2.626	3607
5.0%	5410	29	0	GET	^/api/images/[^/]+	720.973	0.133	1.349	3.257	179937
4.9%	324	0	0	PUT	/api/me/profile-image	711.123	2.195	7.929	8.681	0
4.2%	1628	0	0	POST	/api/auth/login	607.212	0.373	1.912	2.523	245
2.7%	536	0	0	POST	/api/auth/register	393.167	0.734	1.923	2.002	229
2.1%	1805	0	0	GET	/api/posting-window/current	305.082	0.169	1.709	2.527	238
2.1%	292	26	0	POST	^/api/posts/[^/]+/realmojis$	303.747	1.040	7.888	7.946	447
1.6%	1226	0	0	GET	/api/memories	225.334	0.184	1.800	2.761	14510
1.5%	1374	0	0	GET	/api/users	219.948	0.160	1.485	2.381	243
1.4%	560	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	207.055	0.370	1.748	2.412	116
1.2%	369	0	0	PUT	^/api/follows/[^/]+$	180.189	0.488	1.876	2.085	0
1.1%	327	1	0	POST	^/api/posts/[^/]+/comments$	159.231	0.487	2.270	2.667	333
1.0%	857	10	0	GET	^/api/posts/me/current$	147.962	0.173	1.938	2.502	644
0.8%	263	0	0	PATCH	/api/me	120.719	0.459	1.912	2.204	289
0.6%	862	42	0	GET	^/api/posts/[^/]+$	92.482	0.107	0.839	1.800	599
0.6%	684	0	0	GET	^/api/relationships/[^/]+	90.845	0.133	1.498	2.219	625
0.4%	251	0	0	POST	/api/auth/logout	65.142	0.260	1.659	1.781	0
0.4%	88	0	0	PATCH	^/api/posts/[^/]+$	61.431	0.698	2.331	2.331	658
0.3%	413	0	0	GET	^/api/users/[^/]+$	45.166	0.109	0.877	1.047	202
0.3%	34	0	0	DELETE	^/api/posts/[^/]+$	44.629	1.313	2.738	2.738	0
0.3%	330	2	0	GET	/api/me	43.284	0.131	0.938	2.239	303
0.2%	299	0	0	GET	^/api/posts/[^/]+/realmojis$	31.399	0.105	0.669	1.083	2160
0.2%	290	0	0	GET	^/api/posts/[^/]+/comments$	31.207	0.108	0.848	0.878	3140
0.2%	209	0	0	GET	^/api/memories/[^/]+	23.753	0.114	0.858	1.081	881
0.2%	43	0	0	DELETE	^/api/realmojis/[^/]+$	22.431	0.522	2.425	2.425	0
0.1%	60	0	0	DELETE	^/api/comments/[^/]+$	21.014	0.350	1.398	1.398	0
0.1%	89	0	0	GET	/api/me/export	16.368	0.184	3.334	3.334	2188432
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.9%	667	1	0	GET	/api/notifications	4671.237	7.003	10.177	10.257	144
9.6%	2082	20	0	POST	/api/friend-requests	1021.209	0.490	1.738	2.692	114
8.0%	903	56	0	POST	/api/posts	846.594	0.938	6.990	7.243	612
6.3%	4362	21	0	GET	^/api/images/[^/]+	671.826	0.154	1.567	2.917	158917
5.0%	1377	0	0	POST	/api/auth/login	532.466	0.387	1.770	3.243	249
5.0%	3623	0	0	GET	^/api/feeds/[^/]+	532.259	0.147	1.568	2.169	3230
3.0%	239	0	0	PUT	/api/me/profile-image	322.931	1.351	7.432	7.934	0
1.9%	458	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	206.370	0.451	1.713	1.911	116
1.9%	1076	0	0	GET	/api/memories	205.414	0.191	1.967	2.816	15979
1.9%	424	0	0	POST	/api/auth/register	205.188	0.484	1.409	1.530	230
1.8%	1355	0	0	GET	/api/posting-window/current	189.853	0.140	1.294	2.129	240
1.5%	138	20	0	POST	^/api/posts/[^/]+/realmojis$	156.958	1.137	4.586	4.656	383
1.4%	656	9	0	GET	^/api/posts/me/current$	146.181	0.223	1.822	2.791	640
1.4%	959	0	0	GET	/api/users	145.541	0.152	1.464	2.087	241
1.2%	274	0	0	PUT	^/api/follows/[^/]+$	131.374	0.479	1.897	1.969	0
1.1%	217	0	0	PATCH	/api/me	122.170	0.563	1.795	2.184	285
0.8%	134	2	0	POST	^/api/posts/[^/]+/comments$	82.451	0.615	2.138	2.280	323
0.7%	215	0	0	POST	/api/auth/logout	74.726	0.348	1.878	2.403	0
0.7%	452	0	0	GET	^/api/relationships/[^/]+	70.614	0.156	1.827	2.007	823
0.6%	564	21	0	GET	^/api/posts/[^/]+$	59.006	0.105	0.569	1.916	608
0.5%	77	5	0	PATCH	^/api/posts/[^/]+$	49.967	0.649	2.261	2.261	598
0.3%	243	1	0	GET	/api/me	34.884	0.144	0.960	1.827	305
0.3%	259	0	0	GET	^/api/users/[^/]+$	26.611	0.103	0.576	0.666	210
0.2%	191	0	0	GET	^/api/memories/[^/]+	24.789	0.130	1.234	1.578	1033
0.2%	26	0	0	DELETE	^/api/posts/[^/]+$	23.590	0.907	2.021	2.021	0
0.2%	151	0	0	GET	^/api/posts/[^/]+/comments$	17.247	0.114	0.592	0.632	2541
0.1%	81	0	0	GET	/api/me/export	15.147	0.187	1.175	1.175	3044471
0.1%	25	0	0	DELETE	^/api/comments/[^/]+$	14.911	0.596	2.147	2.147	0
0.1%	149	0	0	GET	^/api/posts/[^/]+/realmojis$	14.554	0.098	0.589	0.613	1428
0.1%	22	0	0	DELETE	^/api/realmojis/[^/]+$	14.310	0.650	1.396	1.396	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16693 クエリ / 48 種類 / 合計 177.092s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	53.1%	363	94.017	259.00	2802.12	1220.2	124.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.6%	4	25.859	6464.72	9552.51	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.3%	12412	19.974	1.61	11.49	9.7	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.8%	4	19.101	4775.13	7162.55	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.1%	4	8.979	2244.68	3439.29	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	3.2%	7	5.637	805.29	3344.10	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	1	0.621	621.16	621.16	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
8	0.3%	7	0.618	88.30	156.95	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.3%	7	0.580	82.84	155.95	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.2%	1	0.409	408.71	408.71	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.1%	1	0.222	222.45	222.45	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.222	222.15	222.15	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
13	0.1%	1	0.169	168.59	168.59	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.106	105.94	105.94	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	3672	0.090	0.02	1.78	0.9	29.1	SELECT ip, banned_until FROM ip_bans WHERE banned_until > clock_timestamp()

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
Build ID: 6e933b8e5ca1b29ae10efd66673e1aa465da5495
Type: cpu
Time: 2026-10-31 16:20:28 JST
Duration: 60s, Total samples = 30.94s (51.56%)
Showing nodes accounting for 23.90s, 77.25% of 30.94s total
Dropped 909 nodes (cum <= 0.15s)
Showing top 40 nodes out of 241
      flat  flat%   sum%        cum   cum%
     6.30s 20.36% 20.36%      6.30s 20.36%  crypto/internal/fips140/sha256.blockAVX2
     3.23s 10.44% 30.80%      3.23s 10.44%  internal/runtime/syscall/linux.Syscall6
     1.60s  5.17% 35.97%      5.70s 18.42%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.48s  4.78% 40.76%      1.48s  4.78%  runtime.memmove
     1.29s  4.17% 44.93%      1.29s  4.17%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.12s  3.62% 48.55%      1.12s  3.62%  aeshashbody
     0.95s  3.07% 51.62%      0.95s  3.07%  runtime.memclrNoHeapPointers
     0.77s  2.49% 54.10%      3.01s  9.73%  runtime.scanObject
     0.70s  2.26% 56.37%      0.99s  3.20%  runtime.tryDeferToSpanScan
     0.68s  2.20% 58.56%      0.92s  2.97%  runtime.findObject
     0.49s  1.58% 60.15%      1.86s  6.01%  runtime.mapaccess1_faststr
     0.38s  1.23% 61.38%      7.16s 23.14%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.38s  1.23% 62.61%      0.54s  1.75%  database/sql.convertAssignRows
     0.34s  1.10% 63.70%      0.54s  1.75%  example.com/isureal-company/isureal.postLess
     0.32s  1.03% 64.74%      6.94s 22.43%  crypto/internal/fips140/sha256.(*Digest).Write
     0.32s  1.03% 65.77%      0.32s  1.03%  runtime.nextFreeFast (inline)
     0.28s   0.9% 66.68%      0.28s   0.9%  runtime.futex
     0.26s  0.84% 67.52%      0.81s  2.62%  runtime.mallocgcSmallScanNoHeader
     0.25s  0.81% 68.33%      7.57s 24.47%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.24s  0.78% 69.10%      0.24s  0.78%  runtime.spanClass.sizeclass (inline)
     0.23s  0.74% 69.84%      0.23s  0.74%  memeqbody
     0.23s  0.74% 70.59%      1.55s  5.01%  runtime.mapassign_faststr
     0.21s  0.68% 71.27%      8.08s 26.12%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.20s  0.65% 71.91%      2.37s  7.66%  runtime.mallocgc
     0.20s  0.65% 72.56%      0.20s  0.65%  sync/atomic.(*Int32).Add
     0.16s  0.52% 73.08%      0.20s  0.65%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.14s  0.45% 73.53%      0.39s  1.26%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.14s  0.45% 73.98%      0.17s  0.55%  runtime.typePointers.next
     0.12s  0.39% 74.37%      2.03s  6.56%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s  0.36% 74.73%      0.25s  0.81%  runtime.greyobject
     0.10s  0.32% 75.05%      0.23s  0.74%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.09s  0.29% 75.34%      6.39s 20.65%  crypto/internal/fips140/sha256.block
     0.09s  0.29% 75.63%      1.17s  3.78%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.09s  0.29% 75.92%      0.49s  1.58%  runtime.mallocgcSmallNoscan
     0.09s  0.29% 76.21%      0.27s  0.87%  runtime.mapaccess2_faststr
     0.07s  0.23% 76.44%      0.27s  0.87%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.07s  0.23% 76.66%      0.92s  2.97%  encoding/gob.(*Encoder).encodeStruct
     0.07s  0.23% 76.89%      0.79s  2.55%  runtime.slicebytetostring
     0.06s  0.19% 77.08%      8.44s 27.28%  example.com/isureal-company/isureal.computeImageSignature
     0.05s  0.16% 77.25%      0.49s  1.58%  runtime.netpoll
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 36.5% / 100% | 15% | 1.0% | 5 | 118 MB | run queue が積もっている |
| isu2 (2 cpu) | 28.0% / 100% | 7% | 0.7% | 5 | 174 MB | run queue が積もっている |
| isu3 (2 cpu) | 24.9% / 100% | 8% | 1.4% | 6 | 294 MB | run queue が積もっている |
| isu4 (2 cpu) | 47.9% / 100% | 31% | 9.0% | 22 | 113 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 55% / nginx 10% |
| isu2 | isureal 52% / nginx 2% |
| isu3 | isureal 45% / nginx 2% |
| isu4 | isureal 58% / postgres 17% |

## Diff from previous run

スコア: 194180 → 202575 (+8395)

コードの差分は無い（同じ sha 3f85847efd18）。
スコアの違いはブレか、サーバ側の設定変更による。
