# run 20261031-160719-1f0e740

## Summary

| | |
|---|---|
| score | 175643 |
| status | pass |
| ref | presign-next-minute |
| sha | 1f0e74082a6f |
| 所要時間 | 470s |
| 前回比 | -5527 (-3.1%) vs 20261031-160039-1f0e740 |
| ベスト比 | -16173 (-8.4%) vs 20261031-151339-5a230a4 = 191816 |
| note | sig cache + presign, rerun with heap capture |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	150	441	1185	1184
europe	7	19	57	163	406	407
america	6	18	52	155	266	265
```

```
{"id":"26937","status":"done","portalStatus":"finished","passed":true,"score":175643,"raw":175643,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:13:39.220Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:13:42.218Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:13:45.224Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:13:48.223Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T07:13:50.243Z level=INFO msg="user moved" user=reke_5272 from=asia to=america
time=2026-10-31T07:13:51.215Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T07:13:51.484Z level=INFO msg="user moved" user=reke_5272 from=america to=europe
time=2026-10-31T07:13:54.216Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T07:13:57.224Z level=INFO msg="asia 3 日目: 150 人"
time=2026-10-31T07:13:59.889Z level=INFO msg="user moved" user=yoleo.6036 from=asia to=europe
time=2026-10-31T07:14:00.220Z level=INFO msg="europe 3 日目: 57 人"
time=2026-10-31T07:14:03.215Z level=INFO msg="america 3 日目: 52 人"
time=2026-10-31T07:14:03.636Z level=INFO msg="user moved" user=reke_5272 from=europe to=asia
time=2026-10-31T07:14:06.236Z level=INFO msg="asia 4 日目: 441 人"
time=2026-10-31T07:14:07.882Z level=INFO msg="user moved" user=rari.6460 from=europe to=america
time=2026-10-31T07:14:08.077Z level=INFO msg="user moved" user=mesu-7073 from=asia to=america
time=2026-10-31T07:14:09.734Z level=INFO msg="europe 4 日目: 163 人"
time=2026-10-31T07:14:10.723Z level=INFO msg="user moved" user=yoivy.6731 from=asia to=america
time=2026-10-31T07:14:10.763Z level=INFO msg="user moved" user=mesu-7073 from=america to=europe
time=2026-10-31T07:14:11.401Z level=INFO msg="user moved" user=maxra_8416 from=asia to=europe
time=2026-10-31T07:14:12.217Z level=INFO msg="user moved" user=yoleo.6036 from=europe to=america
time=2026-10-31T07:14:12.309Z level=INFO msg="america 4 日目: 155 人"
time=2026-10-31T07:14:12.504Z level=INFO msg="user moved" user=reke_5272 from=asia to=america
time=2026-10-31T07:14:12.842Z level=INFO msg="user moved" user=ruhemax_8048 from=asia to=america
time=2026-10-31T07:14:12.952Z level=INFO msg="user moved" user=meku-7777 from=asia to=europe
time=2026-10-31T07:14:13.087Z level=INFO msg="user moved" user=remimi_8282 from=asia to=america
time=2026-10-31T07:14:13.224Z level=INFO msg="user moved" user=rimelia8771 from=asia to=america
time=2026-10-31T07:14:14.394Z level=INFO msg="user moved" user=ulshitsu9399 from=america to=europe
time=2026-10-31T07:14:15.277Z level=INFO msg="asia 5 日目: 1185 人"
time=2026-10-31T07:14:16.024Z level=INFO msg="user moved" user=muvinke-8810 from=europe to=america
time=2026-10-31T07:14:19.162Z level=INFO msg="europe 5 日目: 406 人"
time=2026-10-31T07:14:22.155Z level=INFO msg="america 5 日目: 266 人"
time=2026-10-31T07:14:23.090Z level=INFO msg="user moved" user=maxme.9516 from=america to=europe
time=2026-10-31T07:14:30.873Z level=INFO msg="europe 最終日（グランプリ）: 407 人"
time=2026-10-31T07:14:30.933Z level=INFO msg="asia 最終日（グランプリ）: 1184 人"
time=2026-10-31T07:14:31.123Z level=INFO msg="america 最終日（グランプリ）: 265 人"
time=2026-10-31T07:14:39.310Z level=INFO msg="グランプリ: アジア 660 人, ヨーロッパ 155 人, アメリカ 95 人"
{
  "pass": true,
  "score": 175643,
  "messages": [
    "evasake_5066: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "nunoamar.5109: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "yaor.16591: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mamaxan_16077: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "miaki17475: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kiho-16295: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "meto16822: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maxse.17272: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hidanro16225: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "toben_17723: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rotorso.16701: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "torki_16804: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chimiaor_12895: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "semo-10587: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "harasu.17711: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maxeva15200: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "evaul.17891: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "taivy.17124: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yoro17261: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yorudan.6278: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.4%	1841	2	0	GET	/api/notifications	17119.749	9.299	15.839	15.845	163
14.1%	2772	104	0	POST	/api/posts	5813.354	2.097	12.262	12.889	624
9.0%	6402	79	0	POST	/api/friend-requests	3722.043	0.581	2.420	4.019	114
6.1%	19626	397	0	GET	^/api/images/[^/]+	2530.267	0.129	1.104	2.619	117762
4.4%	7084	719	0	POST	/api/auth/login	1829.153	0.258	1.968	4.485	222
3.6%	772	136	0	PUT	/api/me/profile-image	1476.175	1.912	12.250	12.317	3
3.4%	4799	0	0	GET	/api/memories	1386.821	0.289	3.397	7.995	12432
3.1%	2651	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1277.445	0.482	2.387	3.944	116
2.7%	13226	0	0	GET	^/api/feeds/[^/]+	1100.712	0.083	0.527	1.475	2686
2.0%	571	36	0	POST	^/api/posts/[^/]+/realmojis$	814.373	1.426	12.035	12.390	431
1.3%	1210	1	0	POST	/api/auth/register	555.363	0.459	2.006	2.390	227
1.1%	4093	0	0	GET	/api/posting-window/current	442.920	0.108	0.510	0.995	237
1.0%	538	12	0	POST	^/api/posts/[^/]+/comments$	405.336	0.753	4.048	4.417	345
0.8%	3160	0	0	GET	/api/users	338.125	0.107	0.507	0.987	239
0.8%	780	0	0	PUT	^/api/follows/[^/]+$	331.517	0.425	2.312	3.775	0
0.7%	405	0	0	GET	/api/me/export	303.231	0.749	4.642	9.345	1701058
0.6%	2291	97	0	GET	^/api/posts/me/current$	255.175	0.111	0.632	0.995	622
0.6%	1012	0	0	POST	/api/auth/logout	254.600	0.252	2.129	3.085	0
0.6%	565	0	0	PATCH	/api/me	251.835	0.446	2.059	3.082	275
0.4%	260	13	0	PATCH	^/api/posts/[^/]+$	172.054	0.662	3.048	4.622	607
0.3%	1931	33	0	GET	^/api/posts/[^/]+$	140.607	0.073	0.483	0.986	618
0.3%	647	3	0	GET	/api/me	136.161	0.210	0.940	1.184	301
0.3%	117	0	0	DELETE	^/api/posts/[^/]+$	131.382	1.123	2.220	3.376	0
0.3%	1632	0	0	GET	^/api/relationships/[^/]+	119.036	0.073	0.496	0.668	973
0.3%	987	0	0	GET	^/api/memories/[^/]+	111.313	0.113	0.808	2.070	661
0.2%	931	0	0	GET	^/api/users/[^/]+$	65.606	0.070	0.383	0.638	205
0.1%	851	3	0	GET	^/api/posts/[^/]+/realmojis$	56.695	0.067	0.517	0.983	1708
0.1%	696	1	0	GET	^/api/posts/[^/]+/comments$	44.597	0.064	0.420	0.579	2952
0.1%	100	0	0	DELETE	^/api/realmojis/[^/]+$	44.029	0.440	1.898	2.780	0
0.1%	77	0	0	DELETE	^/api/comments/[^/]+$	33.480	0.435	3.109	3.109	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.7%	619	1	0	GET	/api/notifications	4711.808	7.612	11.437	11.643	142
11.5%	1771	16	0	POST	/api/friend-requests	1331.749	0.752	2.737	3.884	114
8.5%	854	25	0	POST	/api/posts	989.537	1.159	3.384	7.066	625
5.7%	1250	0	0	POST	/api/auth/login	658.890	0.527	2.412	3.740	246
4.9%	210	0	0	PUT	/api/me/profile-image	569.378	2.711	11.925	11.940	0
4.4%	3469	0	0	GET	^/api/feeds/[^/]+	513.831	0.148	1.157	1.734	2270
4.0%	3741	29	0	GET	^/api/images/[^/]+	467.525	0.125	1.127	2.558	163294
2.5%	390	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	295.081	0.757	2.903	3.202	116
2.4%	402	0	0	POST	/api/auth/register	280.286	0.697	1.801	1.959	229
2.1%	289	0	0	PUT	^/api/follows/[^/]+$	239.760	0.830	2.415	3.084	0
1.9%	173	14	0	POST	^/api/posts/[^/]+/realmojis$	214.534	1.240	11.492	11.494	426
1.6%	1278	0	0	GET	/api/posting-window/current	190.374	0.149	1.147	1.730	239
1.4%	179	0	0	PATCH	/api/me	158.368	0.885	3.084	3.254	273
1.3%	1000	0	0	GET	/api/memories	146.933	0.147	1.601	3.684	13605
1.1%	883	0	0	GET	/api/users	125.356	0.142	1.145	1.615	241
1.0%	624	31	0	GET	^/api/posts/me/current$	114.490	0.183	1.332	1.708	611
0.9%	125	1	0	POST	^/api/posts/[^/]+/comments$	109.301	0.874	3.842	3.933	332
0.7%	203	0	0	POST	/api/auth/logout	83.610	0.412	2.015	2.294	0
0.6%	74	7	0	PATCH	^/api/posts/[^/]+$	66.156	0.894	3.127	3.127	580
0.5%	478	0	0	GET	^/api/relationships/[^/]+	61.520	0.129	1.078	1.538	583
0.4%	423	13	0	GET	^/api/posts/[^/]+$	48.853	0.115	0.870	1.118	607
0.4%	35	0	0	DELETE	^/api/posts/[^/]+$	46.345	1.324	2.869	2.869	0
0.2%	219	0	0	GET	^/api/users/[^/]+$	26.444	0.121	0.887	1.160	197
0.2%	181	0	0	GET	^/api/posts/[^/]+/comments$	22.804	0.126	0.952	1.157	2589
0.2%	217	2	0	GET	/api/me	22.030	0.102	0.339	0.632	301
0.2%	177	0	0	GET	^/api/posts/[^/]+/realmojis$	21.793	0.123	0.887	1.433	1848
0.2%	78	0	0	GET	/api/me/export	21.002	0.269	2.964	2.964	2415228
0.2%	160	0	0	GET	^/api/memories/[^/]+	19.454	0.122	1.464	2.194	899
0.1%	18	0	0	DELETE	^/api/realmojis/[^/]+$	12.221	0.679	2.409	2.409	0
0.1%	17	0	0	DELETE	^/api/comments/[^/]+$	8.293	0.488	1.490	1.490	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.1%	445	0	0	GET	/api/notifications	3370.777	7.575	9.945	10.090	151
10.4%	1358	15	0	POST	/api/friend-requests	851.426	0.627	2.369	3.256	114
8.9%	600	20	0	POST	/api/posts	733.800	1.223	5.925	6.762	636
7.0%	991	0	0	POST	/api/auth/login	573.360	0.579	2.533	4.282	250
5.6%	3254	11	0	GET	^/api/images/[^/]+	462.457	0.142	1.291	2.901	141269
4.7%	2509	0	0	GET	^/api/feeds/[^/]+	388.897	0.155	1.132	1.799	2884
3.4%	445	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	277.429	0.623	2.384	4.367	116
2.6%	258	0	0	POST	/api/auth/register	213.489	0.827	2.292	2.389	230
2.1%	81	9	0	POST	^/api/posts/[^/]+/realmojis$	168.641	2.082	12.244	12.244	420
1.7%	143	0	0	PUT	/api/me/profile-image	137.181	0.959	3.321	4.583	0
1.5%	871	0	0	GET	/api/posting-window/current	122.570	0.141	1.001	1.230	240
1.4%	173	0	0	PUT	^/api/follows/[^/]+$	117.212	0.678	2.182	3.154	0
1.4%	810	0	0	GET	/api/memories	114.742	0.142	1.870	3.825	18373
1.3%	651	0	0	GET	/api/users	105.059	0.161	1.137	1.375	238
1.2%	117	0	0	PATCH	/api/me	94.918	0.811	2.685	3.013	278
0.9%	157	0	0	POST	/api/auth/logout	71.216	0.454	2.848	3.345	0
0.9%	97	0	0	POST	^/api/posts/[^/]+/comments$	70.269	0.724	3.669	3.669	348
0.8%	439	11	0	GET	^/api/posts/me/current$	65.004	0.148	1.222	1.345	642
0.6%	49	3	0	PATCH	^/api/posts/[^/]+$	50.590	1.032	4.551	4.551	631
0.5%	316	0	0	GET	^/api/relationships/[^/]+	41.016	0.130	0.997	1.151	925
0.4%	329	16	0	GET	^/api/posts/[^/]+$	35.599	0.108	0.715	1.118	599
0.3%	55	0	0	GET	/api/me/export	21.380	0.389	4.733	4.733	3283809
0.2%	179	0	0	GET	^/api/users/[^/]+$	20.257	0.113	0.880	0.944	200
0.2%	148	0	0	GET	/api/me	20.080	0.136	1.372	1.372	303
0.2%	139	0	0	GET	^/api/memories/[^/]+	17.289	0.124	1.027	2.048	821
0.2%	144	0	0	GET	^/api/posts/[^/]+/realmojis$	14.353	0.100	0.314	1.254	1703
0.2%	122	0	0	GET	^/api/posts/[^/]+/comments$	13.821	0.113	0.711	1.219	1991
0.2%	12	0	0	DELETE	^/api/posts/[^/]+$	12.982	1.082	2.379	2.379	0
0.1%	8	0	0	DELETE	^/api/realmojis/[^/]+$	7.333	0.917	2.211	2.211	0
0.1%	17	0	0	DELETE	^/api/comments/[^/]+$	5.533	0.325	1.152	1.152	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16698 クエリ / 48 種類 / 合計 152.484s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	53.4%	417	81.474	195.38	833.51	1093.6	111.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.0%	3	21.305	7101.62	9516.47	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	12.5%	14112	19.063	1.35	11.19	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.5%	3	16.044	5348.13	7156.81	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.2%	3	7.988	2662.70	3634.61	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.6%	4	3.908	976.95	3283.02	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	1	0.558	557.68	557.68	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.3%	4	0.393	98.31	158.04	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.2%	4	0.379	94.78	157.25	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.2%	1	0.303	303.29	303.29	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.2%	1	0.269	269.23	269.23	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
12	0.1%	1	0.167	167.07	167.07	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.107	107.03	107.03	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.080	80.45	80.45	2463.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.073	73.21	73.21	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: f484fa262972f991e07aed1a54eb28ef3ea2570e
Type: cpu
Time: 2026-10-31 16:13:16 JST
Duration: 60.06s, Total samples = 42.58s (70.89%)
Showing nodes accounting for 30.25s, 71.04% of 42.58s total
Dropped 1091 nodes (cum <= 0.21s)
Showing top 40 nodes out of 287
      flat  flat%   sum%        cum   cum%
     5.95s 13.97% 13.97%      5.95s 13.97%  crypto/internal/fips140/sha256.blockAVX2
     4.35s 10.22% 24.19%      4.35s 10.22%  internal/runtime/syscall/linux.Syscall6
     1.95s  4.58% 28.77%      1.95s  4.58%  runtime.memmove
     1.77s  4.16% 32.93%      2.21s  5.19%  runtime.findObject
     1.71s  4.02% 36.94%      5.61s 13.18%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.62s  3.80% 40.75%      1.62s  3.80%  runtime.memclrNoHeapPointers
     1.55s  3.64% 44.39%      2.49s  5.85%  runtime.tryDeferToSpanScan
     1.44s  3.38% 47.77%      1.44s  3.38%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.43s  3.36% 51.13%      6.52s 15.31%  runtime.scanObject
     1.31s  3.08% 54.20%      1.31s  3.08%  aeshashbody
     0.64s  1.50% 55.71%      0.64s  1.50%  runtime.spanClass.sizeclass (inline)
     0.51s  1.20% 56.90%      2.18s  5.12%  runtime.mapaccess1_faststr
     0.37s  0.87% 57.77%      0.73s  1.71%  runtime.greyobject
     0.34s   0.8% 58.57%      0.54s  1.27%  example.com/isureal-company/isureal.postLess
     0.34s   0.8% 59.37%      0.34s   0.8%  memeqbody
     0.30s   0.7% 60.08%      6.58s 15.45%  crypto/internal/fips140/sha256.(*Digest).Write
     0.29s  0.68% 60.76%      6.66s 15.64%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.27s  0.63% 61.39%      0.27s  0.63%  runtime.futex
     0.27s  0.63% 62.02%      0.27s  0.63%  runtime.nextFreeFast (inline)
     0.26s  0.61% 62.64%      0.49s  1.15%  database/sql.convertAssignRows
     0.25s  0.59% 63.22%      0.25s  0.59%  sync/atomic.(*Int32).Add
     0.24s  0.56% 63.79%      0.27s  0.63%  runtime.(*mspan).writeHeapBitsSmall
     0.23s  0.54% 64.33%      0.23s  0.54%  indexbody
     0.23s  0.54% 64.87%      0.25s  0.59%  runtime.typePointers.next
     0.22s  0.52% 65.38%      0.27s  0.63%  encoding/json.appendString[go.shape.string]
     0.21s  0.49% 65.88%      0.49s  1.15%  compress/flate.(*compressor).deflate
     0.20s  0.47% 66.35%      0.25s  0.59%  runtime.(*mspan).markBitsForIndex (inline)
     0.20s  0.47% 66.82%      3.75s  8.81%  runtime.mallocgc
     0.19s  0.45% 67.26%      0.22s  0.52%  runtime.heapArenaOf (inline)
     0.19s  0.45% 67.71%      1.23s  2.89%  runtime.mallocgcSmallScanNoHeader
     0.17s   0.4% 68.11%      7.79s 18.29%  example.com/isureal-company/isureal.computeImageSignature
     0.17s   0.4% 68.51%      0.47s  1.10%  runtime.scanObjectsSmall
     0.16s  0.38% 68.88%      7.35s 17.26%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.16s  0.38% 69.26%      6.95s 16.32%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.16s  0.38% 69.63%      1.59s  3.73%  runtime.mapassign_faststr
     0.14s  0.33% 69.96%      0.37s  0.87%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.14s  0.33% 70.29%      0.80s  1.88%  runtime.mallocgcSmallNoscan
     0.13s  0.31% 70.60%      2.55s  5.99%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.23% 70.83%      0.74s  1.74%  runtime.wbBufFlush1
     0.09s  0.21% 71.04%      0.31s  0.73%  crypto/internal/fips140/hmac.(*HMAC).Reset
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 50.5% / 100% | 34% | 0.7% | 7 | 112 MB | CPU が飽和 |
| isu2 (2 cpu) | 34.4% / 100% | 5% | 0.5% | 4 | 561 MB |  |
| isu3 (2 cpu) | 27.2% / 100% | 4% | 1.3% | 5 | 717 MB | run queue が積もっている |
| isu4 (2 cpu) | 42.7% / 100% | 16% | 16.9% | 11 | 103 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 81% / nginx 13% |
| isu2 | isureal 66% / nginx 2% |
| isu3 | isureal 51% / nginx 1% |
| isu4 | isureal 33% / postgres 16% |

## Diff from previous run

スコア: 181170 → 175643 (-5527)

コードの差分は無い（同じ sha 1f0e74082a6f）。
スコアの違いはブレか、サーバ側の設定変更による。
