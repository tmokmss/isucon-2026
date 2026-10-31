# run 20261031-155347-d7fb2c5

## Summary

| | |
|---|---|
| score | 187910 |
| status | pass |
| ref | mem-headroom |
| sha | d7fb2c5a7b0e |
| 所要時間 | 265s |
| 前回比 | +2832 (1.5%) vs 20261031-152604-1884ea4 |
| ベスト比 | -3906 (-2.0%) vs 20261031-151339-5a230a4 = 191816 |
| note | mem headroom: GOMEMLIMIT 2200MiB, image cache 100MB (on main b3440cf) |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	137	399	1073	1073
europe	7	19	55	164	448	449
america	6	18	59	168	319	316
```

```
{"id":"25333","status":"done","portalStatus":"finished","passed":true,"score":187910,"raw":187910,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:56:33.181Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:56:36.180Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:56:39.176Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:56:42.185Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T06:56:45.175Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T06:56:48.176Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T06:56:51.190Z level=INFO msg="asia 3 日目: 137 人"
time=2026-10-31T06:56:52.827Z level=INFO msg="user moved" user=evahomia_5754 from=asia to=america
time=2026-10-31T06:56:54.177Z level=INFO msg="europe 3 日目: 55 人"
time=2026-10-31T06:56:54.766Z level=INFO msg="user moved" user=makini6119 from=asia to=america
time=2026-10-31T06:56:55.566Z level=INFO msg="user moved" user=evahomia_5754 from=america to=asia
time=2026-10-31T06:56:55.572Z level=INFO msg="user moved" user=ivymaxho.6191 from=asia to=europe
time=2026-10-31T06:56:57.177Z level=INFO msg="america 3 日目: 59 人"
time=2026-10-31T06:56:59.108Z level=INFO msg="user moved" user=hikunoa6249 from=europe to=asia
time=2026-10-31T06:57:00.287Z level=INFO msg="asia 4 日目: 399 人"
time=2026-10-31T06:57:02.790Z level=INFO msg="user moved" user=hikunoa6249 from=asia to=europe
time=2026-10-31T06:57:02.961Z level=INFO msg="user moved" user=ulnu-7330 from=asia to=europe
time=2026-10-31T06:57:03.264Z level=INFO msg="europe 4 日目: 164 人"
time=2026-10-31T06:57:04.491Z level=INFO msg="user moved" user=dansa-7049 from=asia to=america
time=2026-10-31T06:57:04.930Z level=INFO msg="user moved" user=evahomia_5754 from=asia to=america
time=2026-10-31T06:57:05.026Z level=INFO msg="user moved" user=chiya_7606 from=asia to=europe
time=2026-10-31T06:57:05.459Z level=INFO msg="user moved" user=natsu-6465 from=america to=asia
time=2026-10-31T06:57:06.280Z level=INFO msg="user moved" user=toorso-7055 from=asia to=europe
time=2026-10-31T06:57:06.286Z level=INFO msg="america 4 日目: 168 人"
time=2026-10-31T06:57:08.457Z level=INFO msg="user moved" user=ulhotsu-7477 from=asia to=europe
time=2026-10-31T06:57:08.540Z level=INFO msg="user moved" user=nuorku-8828 from=europe to=asia
time=2026-10-31T06:57:08.984Z level=INFO msg="user moved" user=natsu-6465 from=asia to=europe
time=2026-10-31T06:57:09.197Z level=INFO msg="asia 5 日目: 1073 人"
time=2026-10-31T06:57:12.801Z level=INFO msg="europe 5 日目: 448 人"
time=2026-10-31T06:57:13.408Z level=INFO msg="user moved" user=kushi.9214 from=america to=asia
time=2026-10-31T06:57:16.531Z level=INFO msg="america 5 日目: 319 人"
time=2026-10-31T06:57:16.676Z level=INFO msg="user moved" user=ruho_9556 from=america to=europe
time=2026-10-31T06:57:17.319Z level=INFO msg="user moved" user=setsuleo_9849 from=america to=europe
time=2026-10-31T06:57:24.893Z level=INFO msg="europe 最終日（グランプリ）: 449 人"
time=2026-10-31T06:57:25.414Z level=INFO msg="asia 最終日（グランプリ）: 1073 人"
time=2026-10-31T06:57:25.964Z level=INFO msg="america 最終日（グランプリ）: 316 人"
time=2026-10-31T06:57:33.267Z level=INFO msg="グランプリ: アジア 474 人, ヨーロッパ 227 人, アメリカ 99 人"
{
  "pass": true,
  "score": 187910,
  "messages": [
    "runi-15001: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "solhi.15440: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "liazoe.16823: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yane.15505: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "zoemaxor_15664: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noafuvin_16433: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sezoewa-15922: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tame_16365: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "homoan_12727: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "liais-15533: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chikori_16248: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "risa14945: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "solmofu.16706: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maxliaso-15631: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kuleonu.15784: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "matora.8109: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marseleo-16214: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tehi-13655: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kianul.5746: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mizoere_12043: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.5%	1724	1	0	GET	/api/notifications	15944.920	9.249	16.235	16.240	168
11.0%	2725	169	0	POST	/api/posts	3952.328	1.450	6.767	65.813	615
10.6%	6080	99	0	POST	/api/friend-requests	3809.886	0.627	2.423	3.467	114
4.9%	19989	293	0	GET	^/api/images/[^/]+	1768.972	0.088	0.859	3.465	125170
4.5%	6899	553	0	POST	/api/auth/login	1600.070	0.232	1.822	3.063	227
3.4%	2790	12	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1227.290	0.440	2.078	3.099	116
2.9%	4667	0	0	GET	/api/memories	1038.009	0.222	2.243	6.751	12705
2.6%	15745	0	0	GET	^/api/feeds/[^/]+	916.062	0.058	0.431	1.114	3038
2.0%	768	130	0	PUT	/api/me/profile-image	730.197	0.951	9.683	12.689	3
1.7%	767	3	0	POST	^/api/posts/[^/]+/comments$	591.203	0.771	3.803	4.631	335
1.6%	632	5	0	POST	^/api/posts/[^/]+/realmojis$	589.587	0.933	3.762	4.537	477
1.3%	795	0	0	PUT	^/api/follows/[^/]+$	454.261	0.571	2.061	2.811	0
1.2%	1096	1	0	POST	/api/auth/register	422.934	0.386	1.329	1.765	227
0.9%	1152	0	0	POST	/api/auth/logout	324.399	0.282	2.010	4.007	0
0.9%	510	0	0	PATCH	/api/me	320.160	0.628	2.680	3.478	275
0.8%	402	0	0	GET	/api/me/export	292.066	0.727	7.367	11.078	1837560
0.7%	3919	0	0	GET	/api/posting-window/current	266.887	0.068	0.374	0.837	237
0.7%	2931	0	0	GET	/api/users	247.794	0.085	0.425	1.055	241
0.5%	2132	6	0	GET	^/api/posts/me/current$	171.832	0.081	0.448	0.851	653
0.5%	246	1	0	PATCH	^/api/posts/[^/]+$	171.351	0.697	2.503	3.047	652
0.4%	656	3	0	GET	/api/me	157.480	0.240	0.806	0.990	301
0.4%	2489	35	0	GET	^/api/posts/[^/]+$	138.002	0.055	0.407	0.648	610
0.3%	1900	0	0	GET	^/api/relationships/[^/]+	112.646	0.059	0.500	1.048	941
0.3%	101	0	0	DELETE	^/api/posts/[^/]+$	103.319	1.023	2.846	3.072	0
0.2%	1051	0	0	GET	^/api/memories/[^/]+	88.464	0.084	0.765	1.426	643
0.2%	125	0	0	DELETE	^/api/realmojis/[^/]+$	67.594	0.541	2.110	2.238	0
0.2%	1256	0	0	GET	^/api/users/[^/]+$	63.546	0.051	0.369	0.716	197
0.2%	145	0	0	DELETE	^/api/comments/[^/]+$	62.207	0.429	2.232	2.849	0
0.2%	1068	0	0	GET	^/api/posts/[^/]+/realmojis$	54.969	0.051	0.380	0.609	1590
0.1%	834	0	0	GET	^/api/posts/[^/]+/comments$	44.087	0.053	0.353	0.820	2307
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.9%	644	0	0	GET	/api/notifications	4970.806	7.719	12.343	12.344	147
12.4%	2136	23	0	POST	/api/friend-requests	1508.120	0.706	2.449	3.814	114
9.0%	940	56	0	POST	/api/posts	1099.333	1.170	6.703	64.899	617
5.4%	1314	0	0	POST	/api/auth/login	661.276	0.503	2.163	3.552	246
4.7%	4056	0	0	GET	^/api/feeds/[^/]+	575.045	0.142	1.416	1.927	2382
4.4%	4229	19	0	GET	^/api/images/[^/]+	529.298	0.125	1.276	3.975	166908
3.7%	441	0	0	POST	/api/auth/register	446.660	1.013	2.195	2.284	229
2.5%	258	0	0	PUT	/api/me/profile-image	305.699	1.185	2.854	3.192	0
2.0%	300	0	0	PUT	^/api/follows/[^/]+$	242.483	0.808	2.506	2.829	0
1.9%	355	6	0	POST	^/api/friend-requests/[^/]+/[^/]+$	227.417	0.641	2.394	2.741	115
1.8%	1085	0	0	GET	/api/memories	213.045	0.196	1.840	2.831	13060
1.7%	226	0	0	PATCH	/api/me	202.183	0.895	2.800	2.913	283
1.3%	163	3	0	POST	^/api/posts/[^/]+/realmojis$	161.842	0.993	4.297	4.579	467
1.2%	1410	0	0	GET	/api/posting-window/current	142.515	0.101	0.426	1.558	239
1.0%	202	5	0	POST	^/api/posts/[^/]+/comments$	126.934	0.628	2.636	3.380	349
1.0%	1014	0	0	GET	/api/users	115.845	0.114	1.107	1.927	243
0.9%	729	41	0	GET	^/api/posts/me/current$	114.912	0.158	1.512	1.986	621
0.7%	195	0	0	POST	/api/auth/logout	83.022	0.426	2.227	2.362	0
0.7%	506	0	0	GET	^/api/relationships/[^/]+	79.195	0.157	1.396	1.710	508
0.5%	569	18	0	GET	^/api/posts/[^/]+$	66.819	0.117	0.848	1.205	599
0.5%	84	7	0	PATCH	^/api/posts/[^/]+$	57.846	0.689	2.114	2.114	598
0.3%	26	0	0	DELETE	^/api/posts/[^/]+$	34.358	1.321	2.957	2.957	0
0.3%	272	0	0	GET	^/api/users/[^/]+$	31.647	0.116	0.909	1.523	199
0.2%	41	0	0	DELETE	^/api/comments/[^/]+$	26.141	0.638	2.038	2.038	0
0.2%	219	0	0	GET	^/api/posts/[^/]+/comments$	23.422	0.107	0.907	1.836	2376
0.2%	265	2	0	GET	/api/me	22.559	0.085	0.207	0.230	302
0.2%	208	0	0	GET	^/api/posts/[^/]+/realmojis$	21.646	0.104	0.666	0.922	1722
0.2%	63	0	0	GET	/api/me/export	19.933	0.316	3.392	3.392	2148568
0.1%	169	0	0	GET	^/api/memories/[^/]+	15.537	0.092	0.723	1.108	980
0.1%	13	0	0	DELETE	^/api/realmojis/[^/]+$	13.892	1.069	2.193	2.193	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.7%	500	0	0	GET	/api/notifications	3721.120	7.442	10.318	11.054	155
10.6%	701	41	0	POST	/api/posts	967.147	1.380	6.421	63.784	613
10.3%	1608	11	0	POST	/api/friend-requests	937.992	0.583	2.131	3.031	114
6.1%	1136	0	0	POST	/api/auth/login	561.106	0.494	2.243	3.101	250
6.0%	3780	18	0	GET	^/api/images/[^/]+	551.543	0.146	2.071	3.125	142550
4.2%	2955	0	0	GET	^/api/feeds/[^/]+	379.445	0.128	1.452	2.331	3439
3.2%	514	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	291.087	0.566	2.253	2.993	116
2.7%	317	0	0	POST	/api/auth/register	242.399	0.765	2.176	2.273	230
2.1%	171	0	0	PUT	/api/me/profile-image	194.149	1.135	3.029	3.177	0
1.7%	863	0	0	GET	/api/memories	157.589	0.183	1.722	2.534	16566
1.4%	160	0	0	PATCH	/api/me	128.036	0.800	2.408	2.670	278
1.4%	1032	0	0	GET	/api/posting-window/current	124.057	0.120	0.915	1.711	240
1.3%	171	0	0	PUT	^/api/follows/[^/]+$	119.218	0.697	2.421	2.498	0
1.2%	497	18	0	GET	^/api/posts/me/current$	114.102	0.230	1.790	1.963	626
1.1%	107	5	0	POST	^/api/posts/[^/]+/realmojis$	104.504	0.977	3.422	3.777	439
1.0%	101	0	0	POST	^/api/posts/[^/]+/comments$	89.535	0.886	3.255	3.764	342
1.0%	781	0	0	GET	/api/users	87.767	0.112	1.065	1.854	239
0.8%	163	0	0	POST	/api/auth/logout	70.122	0.430	2.067	2.209	0
0.5%	379	0	0	GET	^/api/relationships/[^/]+	48.108	0.127	1.609	1.796	952
0.5%	51	4	0	PATCH	^/api/posts/[^/]+$	46.361	0.909	3.284	3.284	596
0.4%	429	23	0	GET	^/api/posts/[^/]+$	41.094	0.096	0.382	1.802	591
0.3%	240	0	0	GET	^/api/users/[^/]+$	24.639	0.103	0.449	2.019	196
0.3%	20	0	0	DELETE	^/api/posts/[^/]+$	23.785	1.189	2.792	2.792	0
0.2%	168	0	0	GET	^/api/memories/[^/]+	22.574	0.134	2.052	2.191	976
0.2%	60	0	0	GET	/api/me/export	17.026	0.284	3.007	3.007	3197278
0.2%	19	0	0	DELETE	^/api/realmojis/[^/]+$	16.807	0.885	2.727	2.727	0
0.2%	157	0	0	GET	^/api/posts/[^/]+/realmojis$	15.310	0.098	0.356	0.712	1328
0.2%	176	1	0	GET	/api/me	15.307	0.087	0.296	0.320	304
0.1%	112	0	0	GET	^/api/posts/[^/]+/comments$	11.357	0.101	0.384	1.142	2218
0.1%	21	0	0	DELETE	^/api/comments/[^/]+$	8.922	0.425	1.418	1.418	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16822 クエリ / 48 種類 / 合計 149.355s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	54.1%	410	80.841	197.17	1022.07	1175.7	120.3	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.2%	3	21.267	7088.97	9514.75	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.6%	14077	17.276	1.23	9.17	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.7%	3	15.944	5314.71	7213.38	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.4%	3	7.996	2665.37	3650.66	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.5%	4	3.725	931.31	3109.73	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.3%	4	0.389	97.32	155.99	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
8	0.2%	4	0.371	92.75	156.43	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.2%	1	0.282	282.02	282.02	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.259	259.04	259.04	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.2%	1	0.254	253.86	253.86	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.171	171.48	171.48	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.105	104.82	104.82	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.083	83.07	83.07	2412.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.069	68.62	68.62	1503.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
Build ID: e55f72fc0a48ad79e066719c8742d3d4ab0c3a99
Type: cpu
Time: 2026-10-31 15:56:11 JST
Duration: 60.17s, Total samples = 44.19s (73.45%)
Showing nodes accounting for 31.65s, 71.62% of 44.19s total
Dropped 1114 nodes (cum <= 0.22s)
Showing top 40 nodes out of 269
      flat  flat%   sum%        cum   cum%
     5.87s 13.28% 13.28%      5.87s 13.28%  crypto/internal/fips140/sha256.blockAVX2
     4.55s 10.30% 23.58%      4.55s 10.30%  internal/runtime/syscall/linux.Syscall6
     1.97s  4.46% 28.04%      7.98s 18.06%  runtime.scanObject
     1.96s  4.44% 32.47%      2.67s  6.04%  runtime.findObject
     1.91s  4.32% 36.80%      1.91s  4.32%  runtime.memmove
     1.76s  3.98% 40.78%      2.95s  6.68%  runtime.tryDeferToSpanScan
     1.69s  3.82% 44.60%      5.71s 12.92%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.52s  3.44% 48.04%      1.52s  3.44%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.27s  2.87% 50.92%      1.27s  2.87%  runtime.memclrNoHeapPointers
     1.15s  2.60% 53.52%      1.15s  2.60%  aeshashbody
     0.84s  1.90% 55.42%      0.84s  1.90%  runtime.spanClass.sizeclass (inline)
     0.59s  1.34% 56.75%      2.21s  5.00%  runtime.mapaccess1_faststr
     0.47s  1.06% 57.82%      0.63s  1.43%  example.com/isureal-company/isureal.postLess
     0.46s  1.04% 58.86%      0.46s  1.04%  runtime.nextFreeFast (inline)
     0.36s  0.81% 59.67%      0.36s  0.81%  memeqbody
     0.32s  0.72% 60.40%      6.68s 15.12%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.32s  0.72% 61.12%      0.32s  0.72%  runtime.typePointers.nextFast (inline)
     0.31s   0.7% 61.82%      0.31s   0.7%  runtime.(*mspan).divideByElemSize (inline)
     0.30s  0.68% 62.50%      0.77s  1.74%  database/sql.convertAssignRows
     0.28s  0.63% 63.14%      6.45s 14.60%  crypto/internal/fips140/sha256.(*Digest).Write
     0.28s  0.63% 63.77%      0.31s   0.7%  encoding/json.appendString[go.shape.string]
     0.28s  0.63% 64.40%      1.25s  2.83%  runtime.mallocgcSmallScanNoHeader
     0.27s  0.61% 65.01%      3.99s  9.03%  runtime.mallocgc
     0.25s  0.57% 65.58%      0.26s  0.59%  example.com/isureal-company/isureal.rmLess
     0.24s  0.54% 66.12%      0.53s  1.20%  runtime.greyobject
     0.23s  0.52% 66.64%      0.23s  0.52%  runtime.futex
     0.23s  0.52% 67.16%      0.28s  0.63%  runtime.spanOf (inline)
     0.22s   0.5% 67.66%      0.27s  0.61%  runtime.typePointers.next
     0.20s  0.45% 68.11%      1.82s  4.12%  runtime.mapassign_faststr
     0.17s  0.38% 68.50%      0.58s  1.31%  runtime.scanObjectsSmall
     0.16s  0.36% 68.86%      0.35s  0.79%  internal/runtime/maps.(*Iter).Next
     0.16s  0.36% 69.22%      0.25s  0.57%  runtime.heapArenaOf (inline)
     0.15s  0.34% 69.56%      0.35s  0.79%  compress/flate.(*compressor).deflate
     0.15s  0.34% 69.90%      6.93s 15.68%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.15s  0.34% 70.24%      0.24s  0.54%  runtime.(*mspan).markBitsForIndex (inline)
     0.13s  0.29% 70.54%      7.20s 16.29%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.13s  0.29% 70.83%      7.62s 17.24%  example.com/isureal-company/isureal.computeImageSignature
     0.13s  0.29% 71.12%      1.27s  2.87%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.27% 71.40%      2.94s  6.65%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.23% 71.62%      0.32s  0.72%  crypto/internal/fips140/hmac.(*HMAC).Reset
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 43.5% / 100% | 24% | 0.7% | 6 | 130 MB | run queue が積もっている |
| isu2 (2 cpu) | 29.0% / 100% | 4% | 0.6% | 4 | 411 MB |  |
| isu3 (2 cpu) | 25.7% / 100% | 8% | 0.7% | 4 | 441 MB |  |
| isu4 (2 cpu) | 40.4% / 100% | 15% | 16.2% | 6 | 104 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 67% / nginx 12% |
| isu2 | isureal 54% / nginx 2% |
| isu3 | isureal 48% / nginx 2% |
| isu4 | isureal 32% / postgres 15% |

## Diff from previous run

スコア: 185078 → 187910 (+2832)

`1884ea4..d7fb2c5`

```
 conf/isu1/etc/nginx/conf.d/isuh-log.conf        |  1 -
 conf/isu1/etc/systemd/system/isureal-go.service |  2 +-
 conf/isu2/etc/nginx/conf.d/isuh-log.conf        |  1 -
 conf/isu3/etc/nginx/conf.d/isuh-log.conf        |  1 -
 webapp/go/admin_api.go                          | 27 ++++++++++++---
 webapp/go/cmd/isureal/pprof.go                  | 11 ++++++
 webapp/go/export_api.go                         | 12 +++++--
 webapp/go/image_api.go                          | 23 +++++++------
 webapp/go/image_cache.go                        |  2 +-
 webapp/go/image_file.go                         | 33 ++++++++++--------
 webapp/go/image_held.go                         |  5 ++-
 webapp/go/image_local.go                        | 46 ++++++++++---------------
 webapp/go/ip_ban.go                             | 11 ++++++
 webapp/go/repl.go                               |  1 +
 webapp/go/replica_server.go                     |  5 ++-
 15 files changed, 114 insertions(+), 67 deletions(-)
```
