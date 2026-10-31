# run 20261031-153729-0ce5dec

## Summary

| | |
|---|---|
| score | 186373 |
| status | pass |
| ref | eu-us-wait |
| sha | 0ce5dec63a9b |
| 所要時間 | 448s |
| 前回比 | +6181 (3.4%) vs 20261031-153156-b9b8df6 |
| ベスト比 | -5443 (-2.8%) vs 20261031-151339-5a230a4 = 191816 |
| note | eu/us: 画像取得の同時要求を1本に・接続の使い回し（BBRなし、cubicに戻す） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	141	414	1130	1131
europe	7	20	62	173	444	443
america	6	20	62	183	332	331
```

```
{"id":"24055","status":"done","portalStatus":"finished","passed":true,"score":186373,"raw":186373,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:43:13.176Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:43:16.168Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:43:19.167Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:43:22.177Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T06:43:23.878Z level=INFO msg="user moved" user=kereme.5232 from=asia to=america
time=2026-10-31T06:43:25.003Z level=INFO msg="user moved" user=kereme.5232 from=america to=europe
time=2026-10-31T06:43:25.168Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T06:43:28.168Z level=INFO msg="america 2 日目: 20 人"
time=2026-10-31T06:43:31.180Z level=INFO msg="asia 3 日目: 141 人"
time=2026-10-31T06:43:31.784Z level=INFO msg="user moved" user=ivyzoe-5454 from=europe to=america
time=2026-10-31T06:43:32.239Z level=INFO msg="user moved" user=anchike5567 from=asia to=europe
time=2026-10-31T06:43:32.640Z level=INFO msg="user moved" user=nuheno.5686 from=asia to=europe
time=2026-10-31T06:43:32.937Z level=INFO msg="user moved" user=raor5786 from=asia to=europe
time=2026-10-31T06:43:33.481Z level=INFO msg="user moved" user=nuheno.5686 from=europe to=america
time=2026-10-31T06:43:34.169Z level=INFO msg="europe 3 日目: 62 人"
time=2026-10-31T06:43:34.750Z level=INFO msg="user moved" user=chimo-5999 from=asia to=europe
time=2026-10-31T06:43:36.579Z level=INFO msg="user moved" user=kereme.5232 from=europe to=asia
time=2026-10-31T06:43:37.168Z level=INFO msg="america 3 日目: 62 人"
time=2026-10-31T06:43:39.581Z level=INFO msg="user moved" user=rihami6488 from=europe to=asia
time=2026-10-31T06:43:40.181Z level=INFO msg="asia 4 日目: 414 人"
time=2026-10-31T06:43:42.189Z level=INFO msg="user moved" user=rihami6488 from=asia to=europe
time=2026-10-31T06:43:42.448Z level=INFO msg="user moved" user=kereme.5232 from=asia to=america
time=2026-10-31T06:43:43.622Z level=INFO msg="europe 4 日目: 173 人"
time=2026-10-31T06:43:44.118Z level=INFO msg="user moved" user=benya6781 from=asia to=europe
time=2026-10-31T06:43:44.242Z level=INFO msg="user moved" user=natorri7544 from=asia to=europe
time=2026-10-31T06:43:44.694Z level=INFO msg="user moved" user=waha.6859 from=asia to=america
time=2026-10-31T06:43:45.615Z level=INFO msg="user moved" user=vinliahi_7652 from=asia to=america
time=2026-10-31T06:43:45.921Z level=INFO msg="user moved" user=homi.7281 from=asia to=america
time=2026-10-31T06:43:46.044Z level=INFO msg="user moved" user=houldan.7877 from=asia to=america
time=2026-10-31T06:43:46.055Z level=INFO msg="user moved" user=ormar8458 from=asia to=europe
time=2026-10-31T06:43:46.322Z level=INFO msg="america 4 日目: 183 人"
time=2026-10-31T06:43:49.214Z level=INFO msg="asia 5 日目: 1130 人"
time=2026-10-31T06:43:49.437Z level=INFO msg="user moved" user=anchike5567 from=europe to=america
time=2026-10-31T06:43:53.544Z level=INFO msg="europe 5 日目: 444 人"
time=2026-10-31T06:43:54.020Z level=INFO msg="user moved" user=houldan.7877 from=america to=asia
time=2026-10-31T06:43:54.814Z level=INFO msg="user moved" user=raor5786 from=europe to=america
time=2026-10-31T06:43:55.120Z level=INFO msg="user moved" user=hakihi_9765 from=america to=asia
time=2026-10-31T06:43:55.927Z level=INFO msg="america 5 日目: 332 人"
time=2026-10-31T06:44:04.813Z level=INFO msg="asia 最終日（グランプリ）: 1131 人"
time=2026-10-31T06:44:05.287Z level=INFO msg="europe 最終日（グランプリ）: 443 人"
time=2026-10-31T06:44:05.751Z level=INFO msg="america 最終日（グランプリ）: 331 人"
time=2026-10-31T06:44:13.309Z level=INFO msg="グランプリ: アジア 579 人, ヨーロッパ 139 人, アメリカ 106 人"
{
  "pass": true,
  "score": 186373,
  "messages": [
    "hiyo_17187: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "isliavin_16101: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tormiasol.17587: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yosu.7252: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tsuliaivy.16821: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "surewa.6224: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rayu.13923: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "isul_15054: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "suho_16404: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rodanul16242: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "leoroul-6123: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tsusu_17248: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "momarmar-17309: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sasotsu15457: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tasaya_15574: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yomoivy-16162: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yamu17153: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "momax-17229: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "miko16443: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "isfu.18082: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.0%	1784	0	0	GET	/api/notifications	15820.157	8.868	15.712	15.720	169
13.0%	2849	170	0	POST	/api/posts	4909.417	1.723	7.374	9.477	616
10.9%	6392	76	0	POST	/api/friend-requests	4121.673	0.645	2.441	3.636	114
5.1%	7008	493	0	POST	/api/auth/login	1907.797	0.272	2.111	4.501	229
4.6%	20422	412	0	GET	^/api/images/[^/]+	1742.276	0.085	0.966	2.641	116692
3.5%	2939	14	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1315.810	0.448	2.427	4.087	116
3.2%	4903	0	0	GET	/api/memories	1221.530	0.249	2.712	4.631	12908
2.4%	808	136	0	PUT	/api/me/profile-image	887.287	1.098	8.415	11.286	3
2.2%	14386	0	0	GET	^/api/feeds/[^/]+	834.696	0.058	0.479	1.523	2714
2.2%	641	21	0	POST	^/api/posts/[^/]+/realmojis$	822.801	1.284	7.097	7.404	453
1.7%	827	29	0	POST	^/api/posts/[^/]+/comments$	635.068	0.768	4.058	5.467	343
1.3%	1158	1	0	POST	/api/auth/register	479.657	0.414	1.912	1.974	227
1.1%	712	0	0	PUT	^/api/follows/[^/]+$	420.215	0.590	2.372	3.177	0
0.8%	571	0	0	PATCH	/api/me	320.411	0.561	2.320	2.476	279
0.8%	1115	0	0	POST	/api/auth/logout	313.087	0.281	2.326	4.120	0
0.8%	430	0	0	GET	/api/me/export	309.608	0.720	5.502	6.493	1859907
0.6%	4081	0	0	GET	/api/posting-window/current	235.980	0.058	0.443	0.879	237
0.6%	3122	0	0	GET	/api/users	208.316	0.067	0.501	0.817	242
0.5%	2248	33	0	GET	^/api/posts/me/current$	172.478	0.077	0.485	0.809	645
0.4%	245	11	0	PATCH	^/api/posts/[^/]+$	157.854	0.644	3.184	4.048	623
0.3%	685	3	0	GET	/api/me	111.406	0.163	0.787	0.815	301
0.3%	105	0	0	DELETE	^/api/posts/[^/]+$	110.458	1.052	2.425	2.659	0
0.3%	1808	0	0	GET	^/api/relationships/[^/]+	107.948	0.060	0.494	1.563	995
0.3%	2075	35	0	GET	^/api/posts/[^/]+$	105.022	0.051	0.377	0.824	621
0.2%	1003	0	0	GET	^/api/memories/[^/]+	92.017	0.092	0.808	2.318	612
0.2%	100	0	0	DELETE	^/api/realmojis/[^/]+$	66.143	0.661	3.303	4.015	0
0.2%	135	0	0	DELETE	^/api/comments/[^/]+$	65.500	0.485	4.567	5.024	0
0.1%	1081	0	0	GET	^/api/users/[^/]+$	48.099	0.044	0.318	0.782	206
0.1%	913	1	0	GET	^/api/posts/[^/]+/realmojis$	39.072	0.043	0.346	0.880	1373
0.1%	830	0	0	GET	^/api/posts/[^/]+/comments$	38.465	0.046	0.355	0.803	1930
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.4%	649	0	0	GET	/api/notifications	4925.730	7.590	11.730	11.762	147
11.9%	2051	14	0	POST	/api/friend-requests	1457.748	0.711	2.880	4.269	114
9.4%	913	66	0	POST	/api/posts	1145.547	1.255	6.772	7.249	603
5.9%	1286	0	0	POST	/api/auth/login	722.740	0.562	2.608	4.375	245
5.4%	3537	0	0	GET	^/api/feeds/[^/]+	658.625	0.186	2.058	2.665	2645
4.3%	3827	24	0	GET	^/api/images/[^/]+	528.931	0.138	1.529	3.472	163847
3.3%	438	0	0	POST	/api/auth/register	401.006	0.916	2.265	2.342	229
2.3%	431	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	285.817	0.663	2.578	3.534	116
2.3%	241	0	0	PUT	/api/me/profile-image	278.361	1.155	3.863	4.463	0
2.2%	301	0	0	PUT	^/api/follows/[^/]+$	265.534	0.882	2.826	3.737	0
1.6%	976	0	0	GET	/api/memories	195.002	0.200	2.379	4.891	13331
1.5%	200	0	0	PATCH	/api/me	188.222	0.941	3.215	3.873	277
1.4%	1393	0	0	GET	/api/posting-window/current	169.688	0.122	1.520	2.200	239
1.3%	139	8	0	POST	^/api/posts/[^/]+/realmojis$	161.960	1.165	4.462	4.980	439
1.0%	1013	0	0	GET	/api/users	122.458	0.121	1.544	2.253	241
0.9%	616	38	0	GET	^/api/posts/me/current$	114.269	0.186	1.985	2.292	612
0.9%	161	4	0	POST	^/api/posts/[^/]+/comments$	109.461	0.680	4.478	5.400	337
0.6%	192	0	0	POST	/api/auth/logout	78.040	0.406	3.172	3.366	0
0.5%	444	0	0	GET	^/api/relationships/[^/]+	63.357	0.143	1.577	2.088	688
0.5%	460	18	0	GET	^/api/posts/[^/]+$	57.555	0.125	0.860	1.942	606
0.4%	68	5	0	PATCH	^/api/posts/[^/]+$	54.877	0.807	2.647	2.647	618
0.3%	29	0	0	DELETE	^/api/posts/[^/]+$	34.544	1.191	3.290	3.290	0
0.3%	278	0	0	GET	^/api/users/[^/]+$	31.804	0.114	0.874	1.288	209
0.3%	27	0	0	DELETE	^/api/realmojis/[^/]+$	31.416	1.164	2.994	2.994	0
0.2%	185	1	0	GET	^/api/posts/[^/]+/comments$	23.160	0.125	1.300	1.301	2263
0.2%	246	1	0	GET	/api/me	22.301	0.091	0.252	0.696	303
0.2%	185	0	0	GET	^/api/posts/[^/]+/realmojis$	20.622	0.111	0.894	1.285	1533
0.2%	174	0	0	GET	^/api/memories/[^/]+	19.292	0.111	1.605	2.656	817
0.1%	83	0	0	GET	/api/me/export	14.922	0.180	2.383	2.383	2000355
0.1%	25	0	0	DELETE	^/api/comments/[^/]+$	12.119	0.485	2.478	2.478	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.6%	537	2	0	GET	/api/notifications	3988.367	7.427	10.382	10.674	152
11.2%	1670	18	0	POST	/api/friend-requests	1070.679	0.641	2.595	3.747	114
9.1%	743	36	0	POST	/api/posts	871.289	1.173	6.055	6.945	622
6.7%	1148	0	0	POST	/api/auth/login	639.307	0.557	2.719	4.970	249
4.8%	3776	27	0	GET	^/api/images/[^/]+	456.860	0.121	1.156	4.368	147170
4.6%	2954	0	0	GET	^/api/feeds/[^/]+	440.399	0.149	1.575	2.362	3442
2.9%	442	6	0	POST	^/api/friend-requests/[^/]+/[^/]+$	273.578	0.619	2.845	3.546	115
2.4%	324	0	0	POST	/api/auth/register	233.364	0.720	1.919	1.950	230
2.2%	182	0	0	PUT	/api/me/profile-image	214.561	1.179	3.503	3.777	0
1.7%	928	0	0	GET	/api/memories	158.796	0.171	2.122	4.948	15783
1.5%	200	0	0	PUT	^/api/follows/[^/]+$	140.499	0.702	2.942	3.292	0
1.4%	144	0	0	PATCH	/api/me	137.387	0.954	3.476	3.507	280
1.4%	1108	0	0	GET	/api/posting-window/current	132.211	0.119	1.190	1.765	240
1.2%	89	4	0	POST	^/api/posts/[^/]+/realmojis$	117.728	1.323	7.861	7.861	441
1.0%	93	4	0	POST	^/api/posts/[^/]+/comments$	98.181	1.056	5.562	5.562	327
1.0%	532	20	0	GET	^/api/posts/me/current$	95.525	0.180	1.594	2.116	630
1.0%	789	0	0	GET	/api/users	91.105	0.115	0.869	1.646	244
0.9%	181	0	0	POST	/api/auth/logout	87.261	0.482	2.587	2.893	0
0.6%	409	0	0	GET	^/api/relationships/[^/]+	59.657	0.146	1.388	1.688	767
0.5%	53	4	0	PATCH	^/api/posts/[^/]+$	45.376	0.856	3.754	3.754	603
0.5%	408	29	0	GET	^/api/posts/[^/]+$	44.304	0.109	0.288	1.411	584
0.3%	29	0	0	DELETE	^/api/posts/[^/]+$	33.050	1.140	2.385	2.385	0
0.3%	260	0	0	GET	^/api/users/[^/]+$	30.974	0.119	0.783	0.831	206
0.2%	186	1	0	GET	/api/me	18.667	0.100	0.263	0.268	304
0.2%	155	0	0	GET	^/api/memories/[^/]+	18.503	0.119	1.859	1.885	1147
0.2%	152	0	0	GET	^/api/posts/[^/]+/comments$	17.684	0.116	0.830	0.855	2063
0.2%	168	0	0	GET	^/api/posts/[^/]+/realmojis$	16.867	0.100	0.857	0.858	1215
0.1%	14	0	0	PUT	^/api/blocks/[^/]+$	12.487	0.892	2.572	2.572	0
0.1%	10	0	0	DELETE	^/api/realmojis/[^/]+$	10.618	1.062	2.639	2.639	0
0.1%	62	0	0	GET	/api/me/export	10.065	0.162	0.880	0.880	2926279
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 18617 クエリ / 48 種類 / 合計 157.411s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	55.6%	440	87.460	198.77	1789.76	1176.9	120.7	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	13.6%	3	21.441	7146.83	9511.80	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.6%	14683	18.187	1.24	10.51	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.2%	3	16.082	5360.77	7148.09	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.3%	3	8.319	2772.96	4095.98	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.1%	5	3.344	668.74	2664.53	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.3%	5	0.521	104.24	158.28	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.3%	5	0.510	101.91	156.82	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.277	277.38	277.38	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.249	248.69	248.69	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.1%	1	0.224	224.44	224.44	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.153	153.21	153.21	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.120	119.85	119.85	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.081	81.45	81.45	2463.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.078	78.39	78.39	1519.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
Build ID: 7026b7f2d70dd44838b8403bad917c6e11d996ae
Type: cpu
Time: 2026-10-31 15:42:53 JST
Duration: 60.24s, Total samples = 48.16s (79.95%)
Showing nodes accounting for 34.39s, 71.41% of 48.16s total
Dropped 1131 nodes (cum <= 0.24s)
Showing top 40 nodes out of 289
      flat  flat%   sum%        cum   cum%
     6.47s 13.43% 13.43%      6.47s 13.43%  crypto/internal/fips140/sha256.blockAVX2
     5.46s 11.34% 24.77%      5.46s 11.34%  internal/runtime/syscall/linux.Syscall6
     2.06s  4.28% 29.05%      2.06s  4.28%  runtime.memmove
     1.94s  4.03% 33.08%      2.65s  5.50%  runtime.findObject
     1.87s  3.88% 36.96%      7.58s 15.74%  runtime.scanObject
     1.73s  3.59% 40.55%      3.12s  6.48%  runtime.tryDeferToSpanScan
     1.59s  3.30% 43.85%      5.61s 11.65%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.59s  3.30% 47.16%      1.59s  3.30%  runtime.memclrNoHeapPointers
     1.54s  3.20% 50.35%      1.54s  3.20%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.53s  3.18% 53.53%      1.53s  3.18%  aeshashbody
     0.85s  1.76% 55.29%      0.85s  1.76%  runtime.spanClass.sizeclass (inline)
     0.51s  1.06% 56.35%      2.28s  4.73%  runtime.mapaccess1_faststr
     0.42s  0.87% 57.23%      0.42s  0.87%  memeqbody
     0.41s  0.85% 58.08%      0.55s  1.14%  example.com/isureal-company/isureal.postLess
     0.41s  0.85% 58.93%      0.41s  0.85%  runtime.nextFreeFast (inline)
     0.37s  0.77% 59.70%      7.20s 14.95%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.35s  0.73% 60.42%      0.35s  0.73%  runtime.futex
     0.33s  0.69% 61.11%      0.33s  0.69%  sync/atomic.(*Int32).Add
     0.32s  0.66% 61.77%      0.37s  0.77%  runtime.typePointers.next
     0.30s  0.62% 62.40%      1.40s  2.91%  runtime.mallocgcSmallScanNoHeader
     0.29s   0.6% 63.00%      0.76s  1.58%  database/sql.convertAssignRows
     0.29s   0.6% 63.60%      0.36s  0.75%  runtime.spanOf (inline)
     0.27s  0.56% 64.16%      8.18s 16.99%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.27s  0.56% 64.72%      7.61s 15.80%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.25s  0.52% 65.24%      7.06s 14.66%  crypto/internal/fips140/sha256.(*Digest).Write
     0.25s  0.52% 65.76%      0.26s  0.54%  encoding/json.appendString[go.shape.string]
     0.25s  0.52% 66.28%      2.31s  4.80%  runtime.mapassign_faststr
     0.24s   0.5% 66.78%      0.26s  0.54%  runtime.(*mspan).writeHeapBitsSmall
     0.24s   0.5% 67.28%      1.11s  2.30%  runtime.wbBufFlush1
     0.23s  0.48% 67.75%      0.42s  0.87%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.23s  0.48% 68.23%      0.39s  0.81%  runtime.heapArenaOf (inline)
     0.21s  0.44% 68.67%      0.51s  1.06%  compress/flate.(*compressor).deflate
     0.19s  0.39% 69.06%      0.31s  0.64%  runtime.(*mspan).markBitsForIndex (inline)
     0.19s  0.39% 69.46%      0.57s  1.18%  runtime.greyobject
     0.18s  0.37% 69.83%      8.82s 18.31%  example.com/isureal-company/isureal.computeImageSignature
     0.16s  0.33% 70.16%      2.90s  6.02%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.15s  0.31% 70.47%      5.29s 10.98%  example.com/isureal-company/isureal.listMemories
     0.15s  0.31% 70.78%      4.23s  8.78%  runtime.mallocgc
     0.15s  0.31% 71.10%      0.28s  0.58%  runtime.scanObjectSmall
     0.15s  0.31% 71.41%      0.60s  1.25%  runtime.scanObjectsSmall
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 46.6% / 100% | 31% | 0.4% | 6 | 228 MB | CPU が飽和 |
| isu2 (2 cpu) | 27.8% / 100% | 5% | 0.5% | 5 | 179 MB | run queue が積もっている |
| isu3 (2 cpu) | 25.8% / 100% | 4% | 0.5% | 6 | 183 MB | run queue が積もっている |
| isu4 (2 cpu) | 39.0% / 100% | 16% | 14.8% | 5 | 104 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 77% / nginx 13% |
| isu2 | isureal 52% / nginx 2% |
| isu3 | isureal 49% / nginx 2% |
| isu4 | isureal 31% / postgres 15% |

## Diff from previous run

スコア: 180192 → 186373 (+6181)

`b9b8df6..0ce5dec`

```
 conf/_all/etc/modules-load.d/bbr.conf | 1 -
 conf/_all/etc/sysctl.d/99-isucon.conf | 4 ++--
 2 files changed, 2 insertions(+), 3 deletions(-)
```
