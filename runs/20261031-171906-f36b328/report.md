# run 20261031-171906-f36b328

## Summary

| | |
|---|---|
| score | 198386 |
| status | pass |
| ref | f36b328 |
| sha | f36b328b59a9 |
| 所要時間 | 240s |
| 前回比 | -3199 (-1.6%) vs 20261031-171105-5c44717 |
| ベスト比 | -33724 (-14.5%) vs 20261031-164055-bdf05cd = 232110 |
| note | memories to EU US replicas, export on asia1 (main ac5b09b) |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	150	442	1184	1183
europe	7	23	73	210	475	473
america	6	21	60	177	351	352
```

```
{"id":"33734","status":"done","portalStatus":"finished","passed":true,"score":198386,"raw":198386,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:21:29.250Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:21:32.239Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:21:35.240Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:21:38.257Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T08:21:38.580Z level=INFO msg="user moved" user=ruke_5158 from=asia to=europe
time=2026-10-31T08:21:39.248Z level=INFO msg="user moved" user=ruke_5158 from=europe to=asia
time=2026-10-31T08:21:40.495Z level=INFO msg="user moved" user=fukoko_5179 from=asia to=america
time=2026-10-31T08:21:41.248Z level=INFO msg="europe 2 日目: 23 人"
time=2026-10-31T08:21:44.239Z level=INFO msg="america 2 日目: 21 人"
time=2026-10-31T08:21:47.252Z level=INFO msg="asia 3 日目: 150 人"
time=2026-10-31T08:21:48.848Z level=INFO msg="user moved" user=ruke_5158 from=asia to=europe
time=2026-10-31T08:21:49.754Z level=INFO msg="user moved" user=kanoame_6049 from=asia to=europe
time=2026-10-31T08:21:50.240Z level=INFO msg="europe 3 日目: 73 人"
time=2026-10-31T08:21:50.373Z level=INFO msg="user moved" user=kanoame_6049 from=europe to=asia
time=2026-10-31T08:21:50.802Z level=INFO msg="user moved" user=yono6224 from=asia to=europe
time=2026-10-31T08:21:53.240Z level=INFO msg="america 3 日目: 60 人"
time=2026-10-31T08:21:55.930Z level=INFO msg="user moved" user=muke-6385 from=europe to=asia
time=2026-10-31T08:21:56.272Z level=INFO msg="asia 4 日目: 442 人"
time=2026-10-31T08:21:58.068Z level=INFO msg="user moved" user=ivyramu_6798 from=america to=asia
time=2026-10-31T08:21:59.820Z level=INFO msg="user moved" user=kanoame_6049 from=asia to=europe
time=2026-10-31T08:21:59.870Z level=INFO msg="europe 4 日目: 210 人"
time=2026-10-31T08:21:59.871Z level=INFO msg="user moved" user=rosevin6521 from=europe to=asia
time=2026-10-31T08:21:59.995Z level=INFO msg="user moved" user=yaivyse-7010 from=asia to=america
time=2026-10-31T08:22:00.352Z level=INFO msg="user moved" user=wamemo7093 from=asia to=america
time=2026-10-31T08:22:00.924Z level=INFO msg="user moved" user=ivyramu_6798 from=asia to=europe
time=2026-10-31T08:22:01.488Z level=INFO msg="user moved" user=wamemo7093 from=america to=asia
time=2026-10-31T08:22:02.151Z level=INFO msg="user moved" user=yaivyse-7010 from=america to=asia
time=2026-10-31T08:22:02.249Z level=INFO msg="america 4 日目: 177 人"
time=2026-10-31T08:22:03.299Z level=INFO msg="user moved" user=ruke_5158 from=europe to=asia
time=2026-10-31T08:22:03.656Z level=INFO msg="user moved" user=mayoel.7581 from=asia to=europe
time=2026-10-31T08:22:04.512Z level=INFO msg="user moved" user=naka8508 from=asia to=america
time=2026-10-31T08:22:05.166Z level=INFO msg="user moved" user=orsu.9090 from=europe to=america
time=2026-10-31T08:22:05.272Z level=INFO msg="asia 5 日目: 1184 人"
time=2026-10-31T08:22:08.999Z level=INFO msg="europe 5 日目: 475 人"
time=2026-10-31T08:22:11.642Z level=INFO msg="america 5 日目: 351 人"
time=2026-10-31T08:22:17.116Z level=INFO msg="user moved" user=rekuke-10118 from=europe to=america
time=2026-10-31T08:22:21.163Z level=INFO msg="america 最終日（グランプリ）: 352 人"
time=2026-10-31T08:22:21.884Z level=INFO msg="asia 最終日（グランプリ）: 1183 人"
time=2026-10-31T08:22:22.217Z level=INFO msg="europe 最終日（グランプリ）: 473 人"
time=2026-10-31T08:22:29.376Z level=INFO msg="グランプリ: アジア 467 人, ヨーロッパ 185 人, アメリカ 174 人"
{
  "pass": true,
  "score": 198386,
  "messages": [
    "noasolmo-5025: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "kuho_16990: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "dannu_11532: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "anel_16636: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maraki_17450: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "niben_17841: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ivykalia.17974: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "danivy18376: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "elivy.18889: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sushi18160: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mosu18810: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "henu.17923: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "remata.18098: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yaliaki-18384: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yoru17533: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kuyanu-18334: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ridanki_18984: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ratorwa.18600: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "murumu_12359: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mozoe.18566: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
46.4%	1884	1	0	GET	/api/notifications	19119.842	10.149	16.639	16.671	172
9.2%	2924	136	0	POST	/api/posts	3808.459	1.302	6.868	8.215	623
8.4%	6788	73	0	POST	/api/friend-requests	3450.544	0.508	2.394	3.946	114
7.4%	17916	348	0	GET	^/api/images/[^/]+	3051.293	0.170	1.870	3.727	146798
3.6%	3724	12	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1481.822	0.398	2.217	3.735	116
3.6%	4157	0	0	GET	/api/memories	1480.037	0.356	2.618	6.634	12114
3.4%	6384	819	0	POST	/api/auth/login	1404.741	0.220	1.835	4.546	216
3.0%	16538	0	0	GET	^/api/feeds/[^/]+	1255.227	0.076	0.717	1.604	4047
2.3%	726	11	0	POST	^/api/posts/[^/]+/realmojis$	945.337	1.302	7.784	8.516	461
2.0%	833	130	0	PUT	/api/me/profile-image	814.624	0.978	8.360	72.712	3
1.4%	804	5	0	POST	^/api/posts/[^/]+/comments$	566.229	0.704	3.807	4.614	347
1.2%	4251	0	0	GET	/api/posting-window/current	503.256	0.118	0.805	1.205	237
1.1%	3200	0	0	GET	/api/users	474.133	0.148	0.943	1.254	242
0.9%	1203	1	0	POST	/api/auth/register	389.155	0.323	1.116	1.626	227
0.7%	397	0	0	GET	/api/me/export	277.873	0.700	3.801	7.777	1873336
0.7%	2205	20	0	GET	^/api/posts/me/current$	277.608	0.126	0.878	1.223	647
0.6%	918	0	0	POST	/api/auth/logout	262.442	0.286	1.962	6.133	0
0.6%	833	0	0	GET	^/api/memories/[^/]+	235.998	0.283	1.717	3.481	677
0.6%	795	0	0	PUT	^/api/follows/[^/]+$	234.733	0.295	2.169	2.987	0
0.4%	557	0	0	PATCH	/api/me	151.734	0.272	1.967	2.953	282
0.3%	2575	37	0	GET	^/api/posts/[^/]+$	140.989	0.055	0.555	1.165	626
0.3%	1967	0	0	GET	^/api/relationships/[^/]+	137.689	0.070	0.734	1.196	1104
0.3%	238	4	0	PATCH	^/api/posts/[^/]+$	128.469	0.540	2.425	2.467	639
0.2%	720	3	0	GET	/api/me	102.890	0.143	0.864	1.480	301
0.2%	147	0	0	DELETE	^/api/realmojis/[^/]+$	102.549	0.698	2.393	3.461	0
0.2%	102	0	0	DELETE	^/api/posts/[^/]+$	90.926	0.891	2.685	2.939	0
0.2%	1358	1	0	GET	^/api/users/[^/]+$	77.443	0.057	0.559	1.127	207
0.2%	181	0	0	DELETE	^/api/comments/[^/]+$	69.914	0.386	2.233	2.365	0
0.1%	1020	4	0	GET	^/api/posts/[^/]+/realmojis$	61.005	0.060	0.661	1.083	1443
0.1%	915	6	0	GET	^/api/posts/[^/]+/comments$	53.573	0.059	0.558	0.814	2495
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.3%	727	1	0	GET	/api/notifications	5596.472	7.698	12.818	14.010	165
13.0%	2411	24	0	POST	/api/friend-requests	1716.074	0.712	2.780	3.764	114
8.0%	1059	40	0	POST	/api/posts	1057.998	0.999	5.826	6.555	627
5.4%	1379	0	0	POST	/api/auth/login	713.555	0.517	2.668	3.593	244
4.7%	4427	35	0	GET	^/api/images/[^/]+	617.679	0.140	1.188	3.036	180020
4.6%	4475	0	0	GET	^/api/feeds/[^/]+	612.959	0.137	1.022	1.757	3311
2.7%	554	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	360.536	0.651	2.567	3.031	116
2.5%	260	0	0	PUT	/api/me/profile-image	336.748	1.295	6.438	6.980	0
2.2%	472	0	0	POST	/api/auth/register	291.387	0.617	2.570	2.582	228
1.9%	303	0	0	PUT	^/api/follows/[^/]+$	251.220	0.829	2.989	3.183	0
1.5%	229	0	0	PATCH	/api/me	197.405	0.862	3.117	3.768	277
1.4%	1154	0	0	GET	/api/memories	183.865	0.159	1.265	3.215	12226
1.4%	1586	0	0	GET	/api/posting-window/current	178.961	0.113	0.554	1.411	239
1.3%	1191	0	0	GET	/api/users	167.724	0.141	1.018	1.679	242
1.2%	192	5	0	POST	^/api/posts/[^/]+/comments$	156.289	0.814	3.908	4.316	341
1.0%	809	41	0	GET	^/api/posts/me/current$	138.767	0.172	1.145	1.693	620
0.9%	105	3	0	POST	^/api/posts/[^/]+/realmojis$	123.280	1.174	7.646	7.783	461
0.6%	187	0	0	POST	/api/auth/logout	74.568	0.399	2.242	2.747	0
0.5%	666	27	0	GET	^/api/posts/[^/]+$	68.753	0.103	0.527	1.156	611
0.5%	545	0	0	GET	^/api/relationships/[^/]+	67.249	0.123	0.978	1.157	774
0.5%	83	5	0	PATCH	^/api/posts/[^/]+$	63.180	0.761	2.374	2.374	614
0.3%	43	0	0	DELETE	^/api/posts/[^/]+$	43.136	1.003	3.025	3.025	0
0.3%	266	2	0	GET	/api/me	35.061	0.132	0.514	1.181	302
0.3%	343	0	0	GET	^/api/users/[^/]+$	34.378	0.100	0.749	0.861	204
0.2%	184	0	0	GET	^/api/memories/[^/]+	23.512	0.128	1.095	1.853	1008
0.2%	225	0	0	GET	^/api/posts/[^/]+/realmojis$	21.273	0.095	0.600	0.676	1499
0.1%	195	1	0	GET	^/api/posts/[^/]+/comments$	19.112	0.098	0.605	1.160	2936
0.1%	39	0	0	DELETE	^/api/comments/[^/]+$	17.539	0.450	1.750	1.750	0
0.1%	21	0	0	DELETE	^/api/realmojis/[^/]+$	17.432	0.830	2.074	2.074	0
0.1%	61	0	0	GET	/api/me/export	14.830	0.243	4.179	4.179	2167316
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.1%	507	0	0	GET	/api/notifications	3827.645	7.550	9.754	9.963	158
10.7%	1688	12	0	POST	/api/friend-requests	997.851	0.591	2.389	3.407	114
8.3%	742	25	0	POST	/api/posts	770.684	1.039	6.472	7.686	636
6.7%	1138	0	0	POST	/api/auth/login	626.178	0.550	2.546	3.326	248
5.2%	3791	19	0	GET	^/api/images/[^/]+	480.021	0.127	1.309	3.819	146745
4.0%	2885	0	0	GET	^/api/feeds/[^/]+	376.063	0.130	1.156	1.979	3311
3.5%	348	0	0	POST	/api/auth/register	326.549	0.938	2.401	2.544	230
3.3%	483	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	306.861	0.635	2.855	3.601	116
2.3%	208	0	0	PUT	/api/me/profile-image	218.443	1.050	6.296	8.116	0
1.7%	215	0	0	PUT	^/api/follows/[^/]+$	157.903	0.734	2.335	2.837	0
1.6%	166	0	0	PATCH	/api/me	150.798	0.908	3.651	4.494	286
1.5%	1107	0	0	GET	/api/posting-window/current	144.368	0.130	0.949	1.744	240
1.5%	142	5	0	POST	^/api/posts/[^/]+/realmojis$	135.811	0.956	3.763	4.459	439
1.4%	968	0	0	GET	/api/memories	132.595	0.137	1.264	2.848	15362
1.2%	820	0	0	GET	/api/users	114.680	0.140	1.132	1.673	242
1.1%	128	1	0	POST	^/api/posts/[^/]+/comments$	99.963	0.781	3.813	3.838	307
0.8%	591	12	0	GET	^/api/posts/me/current$	79.043	0.134	1.232	1.906	644
0.8%	152	0	0	POST	/api/auth/logout	78.225	0.515	3.421	4.350	0
0.6%	76	7	0	PATCH	^/api/posts/[^/]+$	51.542	0.678	2.966	2.966	599
0.5%	392	0	0	GET	^/api/relationships/[^/]+	46.656	0.119	1.077	1.404	909
0.4%	343	14	0	GET	^/api/posts/[^/]+$	40.632	0.118	1.097	1.582	605
0.2%	211	0	0	GET	/api/me	21.895	0.104	0.221	0.272	305
0.2%	64	0	0	GET	/api/me/export	20.705	0.324	4.237	4.237	2897981
0.2%	161	0	0	GET	^/api/users/[^/]+$	20.189	0.125	1.024	1.687	196
0.2%	22	0	0	DELETE	^/api/posts/[^/]+$	19.748	0.898	2.571	2.571	0
0.2%	140	0	0	GET	^/api/memories/[^/]+	16.314	0.117	1.166	1.192	1331
0.2%	144	0	0	GET	^/api/posts/[^/]+/realmojis$	15.014	0.104	0.536	0.939	1354
0.1%	103	0	0	GET	^/api/posts/[^/]+/comments$	12.856	0.125	1.058	1.129	2584
0.1%	21	0	0	DELETE	^/api/realmojis/[^/]+$	10.874	0.518	1.509	1.509	0
0.1%	19	0	0	DELETE	^/api/comments/[^/]+$	6.538	0.344	0.934	0.934	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 15558 クエリ / 48 種類 / 合計 184.910s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	55.6%	406	102.724	253.01	2027.06	1190.8	121.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.3%	4	26.393	6598.13	9994.43	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	10.5%	11674	19.417	1.66	20.63	9.5	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.1%	4	18.694	4673.50	7157.73	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	4.9%	4	9.011	2252.66	3468.02	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.8%	6	5.090	848.37	2441.36	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.3%	1	0.508	508.28	508.28	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.3%	6	0.465	77.58	158.41	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.2%	6	0.454	75.73	157.43	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.2%	1	0.422	422.13	422.13	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.395	395.09	395.09	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.2%	1	0.331	331.39	331.39	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
13	0.1%	1	0.171	170.69	170.69	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.165	164.83	164.83	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	1	0.150	150.01	150.01	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows

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
Build ID: 38b189f601d4939a44548c730b7993bb658a29c0
Type: cpu
Time: 2026-10-31 17:21:05 JST
Duration: 60.01s, Total samples = 38.57s (64.27%)
Showing nodes accounting for 29.57s, 76.67% of 38.57s total
Dropped 1047 nodes (cum <= 0.19s)
Showing top 40 nodes out of 242
      flat  flat%   sum%        cum   cum%
     6.69s 17.35% 17.35%      6.69s 17.35%  crypto/internal/fips140/sha256.blockAVX2
     3.80s  9.85% 27.20%      3.80s  9.85%  internal/runtime/syscall/linux.Syscall6
     2.09s  5.42% 32.62%      2.09s  5.42%  runtime.memmove
     1.79s  4.64% 37.26%      6.32s 16.39%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.48s  3.84% 41.09%      1.48s  3.84%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.37s  3.55% 44.65%      1.37s  3.55%  runtime.memclrNoHeapPointers
     1.34s  3.47% 48.12%      5.18s 13.43%  runtime.scanObject
     1.25s  3.24% 51.36%      1.55s  4.02%  runtime.findObject
     1.19s  3.09% 54.45%      1.19s  3.09%  aeshashbody
     1.09s  2.83% 57.27%      1.84s  4.77%  runtime.tryDeferToSpanScan
     0.60s  1.56% 58.83%      2.01s  5.21%  runtime.mapaccess1_faststr
     0.52s  1.35% 60.18%      0.52s  1.35%  runtime.spanClass.sizeclass (inline)
     0.46s  1.19% 61.37%      0.65s  1.69%  example.com/isureal-company/isureal.postLess
     0.46s  1.19% 62.56%      0.46s  1.19%  runtime.nextFreeFast (inline)
     0.45s  1.17% 63.73%      7.64s 19.81%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.40s  1.04% 64.77%      0.40s  1.04%  memeqbody
     0.36s  0.93% 65.70%      1.93s  5.00%  runtime.mapassign_faststr
     0.32s  0.83% 66.53%      7.36s 19.08%  crypto/internal/fips140/sha256.(*Digest).Write
     0.31s   0.8% 67.33%      0.73s  1.89%  database/sql.convertAssignRows
     0.30s  0.78% 68.11%      8.64s 22.40%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.28s  0.73% 68.84%      1.29s  3.34%  runtime.mallocgcSmallScanNoHeader
     0.27s   0.7% 69.54%      0.27s   0.7%  sync/atomic.(*Int32).Add
     0.23s   0.6% 70.13%      0.26s  0.67%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.22s  0.57% 70.70%      0.43s  1.11%  runtime.greyobject
     0.22s  0.57% 71.27%      0.27s   0.7%  runtime.typePointers.next
     0.21s  0.54% 71.82%      0.39s  1.01%  compress/flate.(*compressor).deflate
     0.21s  0.54% 72.36%      8.05s 20.87%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.21s  0.54% 72.91%      0.45s  1.17%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.21s  0.54% 73.45%      0.24s  0.62%  runtime.(*mspan).writeHeapBitsSmall
     0.20s  0.52% 73.97%      0.20s  0.52%  runtime.futex
     0.14s  0.36% 74.33%      0.37s  0.96%  runtime.scanObjectsSmall
     0.13s  0.34% 74.67%      2.88s  7.47%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.13s  0.34% 75.01%      0.48s  1.24%  runtime.mapaccess2_faststr
     0.10s  0.26% 75.27%      6.79s 17.60%  crypto/internal/fips140/sha256.block
     0.10s  0.26% 75.53%      0.71s  1.84%  runtime.mallocgcSmallNoscan
     0.10s  0.26% 75.78%      0.62s  1.61%  runtime.scanSpan
     0.09s  0.23% 76.02%      0.28s  0.73%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.09s  0.23% 76.25%      3.80s  9.85%  runtime.mallocgc
     0.08s  0.21% 76.46%      9.05s 23.46%  example.com/isureal-company/isureal.computeImageSignature
     0.08s  0.21% 76.67%      0.85s  2.20%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 50.5% / 100% | 33% | 4.7% | 7 | 109 MB | CPU が飽和 |
| isu2 (2 cpu) | 28.2% / 100% | 5% | 0.7% | 5 | 127 MB | run queue が積もっている |
| isu3 (2 cpu) | 25.2% / 100% | 3% | 0.9% | 5 | 640 MB | run queue が積もっている |
| isu4 (2 cpu) | 50.7% / 100% | 27% | 13.2% | 11 | 118 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 66% / nginx 13% / find 6% / kswapd0 2% |
| isu2 | isureal 53% / nginx 2% |
| isu3 | isureal 47% / nginx 2% |
| isu4 | isureal 53% / postgres 19% |

## Diff from previous run

スコア: 201585 → 198386 (-3199)

`5c44717..f36b328`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf | 11 +++++++++--
 1 file changed, 9 insertions(+), 2 deletions(-)
```
