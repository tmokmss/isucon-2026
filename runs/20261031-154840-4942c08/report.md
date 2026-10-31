# run 20261031-154840-4942c08

## Summary

| | |
|---|---|
| score | 183742 |
| status | pass |
| ref | eu-us-wait |
| sha | 4942c089b9e0 |
| 所要時間 | 191s |
| 前回比 | +183742 (—) vs 20261031-154510-4942c08 |
| ベスト比 | -8074 (-4.2%) vs 20261031-151339-5a230a4 = 191816 |
| note | eu/us: 画像取得の同時要求を1本に・接続の使い回し（main b3440cf 取り込み版、asia1 の OOM で fail した後の再計測） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	141	413	1119	1120
europe	7	21	62	186	512	510
america	6	18	55	162	289	289
```

```
{"id":"24764","status":"done","portalStatus":"finished","passed":true,"score":183742,"raw":183742,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:50:19.913Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:50:22.911Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:50:25.904Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:50:28.916Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T06:50:31.314Z level=INFO msg="user moved" user=hitoka_5286 from=asia to=america
time=2026-10-31T06:50:31.918Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T06:50:32.184Z level=INFO msg="user moved" user=hitoka_5286 from=america to=europe
time=2026-10-31T06:50:34.906Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T06:50:36.132Z level=INFO msg="user moved" user=raleo5452 from=europe to=america
time=2026-10-31T06:50:37.918Z level=INFO msg="asia 3 日目: 141 人"
time=2026-10-31T06:50:40.437Z level=INFO msg="user moved" user=anseku-5976 from=asia to=america
time=2026-10-31T06:50:40.904Z level=INFO msg="europe 3 日目: 62 人"
time=2026-10-31T06:50:41.042Z level=INFO msg="user moved" user=anseku-5976 from=america to=asia
time=2026-10-31T06:50:42.036Z level=INFO msg="user moved" user=hitoka_5286 from=europe to=asia
time=2026-10-31T06:50:43.905Z level=INFO msg="america 3 日目: 55 人"
time=2026-10-31T06:50:44.933Z level=INFO msg="user moved" user=raleo5452 from=america to=europe
time=2026-10-31T06:50:46.917Z level=INFO msg="asia 4 日目: 413 人"
time=2026-10-31T06:50:48.410Z level=INFO msg="user moved" user=hitoka_5286 from=asia to=america
time=2026-10-31T06:50:48.928Z level=INFO msg="user moved" user=anseku-5976 from=asia to=america
time=2026-10-31T06:50:49.009Z level=INFO msg="user moved" user=liafuma_6681 from=america to=europe
time=2026-10-31T06:50:50.038Z level=INFO msg="europe 4 日目: 186 人"
time=2026-10-31T06:50:51.135Z level=INFO msg="user moved" user=ivyeltsu.7182 from=asia to=america
time=2026-10-31T06:50:51.804Z level=INFO msg="user moved" user=tsusozoe-8274 from=asia to=europe
time=2026-10-31T06:50:52.405Z level=INFO msg="user moved" user=liafuma_6681 from=europe to=america
time=2026-10-31T06:50:52.607Z level=INFO msg="user moved" user=raleo5452 from=europe to=america
time=2026-10-31T06:50:52.838Z level=INFO msg="user moved" user=hianivy.7734 from=asia to=europe
time=2026-10-31T06:50:53.156Z level=INFO msg="america 4 日目: 162 人"
time=2026-10-31T06:50:53.166Z level=INFO msg="user moved" user=keelwa7308 from=asia to=europe
time=2026-10-31T06:50:53.242Z level=INFO msg="user moved" user=nuhe7647 from=asia to=america
time=2026-10-31T06:50:53.823Z level=INFO msg="user moved" user=torul-7893 from=asia to=america
time=2026-10-31T06:50:54.374Z level=INFO msg="user moved" user=mialiani.7922 from=asia to=america
time=2026-10-31T06:50:55.113Z level=INFO msg="user moved" user=ormaxsol.8549 from=europe to=asia
time=2026-10-31T06:50:55.511Z level=INFO msg="user moved" user=zoemia9025 from=europe to=america
time=2026-10-31T06:50:55.555Z level=INFO msg="user moved" user=mehi_7630 from=asia to=europe
time=2026-10-31T06:50:55.926Z level=INFO msg="asia 5 日目: 1119 人"
time=2026-10-31T06:50:58.587Z level=INFO msg="user moved" user=yasara.9478 from=europe to=asia
time=2026-10-31T06:51:00.201Z level=INFO msg="europe 5 日目: 512 人"
time=2026-10-31T06:51:00.808Z level=INFO msg="user moved" user=ruwa-9364 from=europe to=asia
time=2026-10-31T06:51:02.147Z level=INFO msg="america 5 日目: 289 人"
time=2026-10-31T06:51:11.683Z level=INFO msg="europe 最終日（グランプリ）: 510 人"
time=2026-10-31T06:51:12.502Z level=INFO msg="asia 最終日（グランプリ）: 1120 人"
time=2026-10-31T06:51:12.879Z level=INFO msg="america 最終日（グランプリ）: 289 人"
time=2026-10-31T06:51:20.028Z level=INFO msg="グランプリ: アジア 464 人, ヨーロッパ 241 人, アメリカ 102 人"
{
  "pass": true,
  "score": 183742,
  "messages": [
    "miasa9354: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nefu16588: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ruramu.17464: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ivyivyho12441: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "suwa-11219: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "keliaivy6634: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "benrome_15786: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "saramax_16609: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "himar.8013: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tesu.12634: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kibenvin_10804: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "koivyho.17707: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "zoera_18045: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "shinu6280: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yawa.16189: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ulsolmo.17720: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nahima16359: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "istome15505: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tewaor_16261: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "teka17499: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.1%	1789	0	0	GET	/api/notifications	16757.428	9.367	16.614	16.649	170
15.9%	2885	225	0	POST	/api/posts	6461.703	2.240	63.836	66.071	604
10.0%	6572	93	0	POST	/api/friend-requests	4053.673	0.617	2.639	4.348	114
4.5%	19539	343	0	GET	^/api/images/[^/]+	1852.126	0.095	0.796	2.173	124750
4.5%	6666	512	0	POST	/api/auth/login	1815.036	0.272	2.188	4.084	227
3.2%	4643	0	0	GET	/api/memories	1323.597	0.285	2.974	13.913	12945
3.0%	3043	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1212.643	0.399	2.549	3.728	116
2.9%	13991	0	0	GET	^/api/feeds/[^/]+	1198.405	0.086	0.785	4.349	3416
2.0%	807	134	0	PUT	/api/me/profile-image	828.283	1.026	9.327	69.580	3
1.7%	675	14	0	POST	^/api/posts/[^/]+/realmojis$	682.990	1.012	3.867	4.763	466
1.6%	1144	1	0	POST	/api/auth/register	647.209	0.566	1.396	1.450	227
1.3%	674	9	0	POST	^/api/posts/[^/]+/comments$	538.084	0.798	3.917	4.812	350
0.9%	3164	0	0	GET	/api/users	369.051	0.117	0.691	1.074	238
0.8%	678	0	0	PUT	^/api/follows/[^/]+$	341.865	0.504	2.384	3.289	0
0.8%	4063	0	0	GET	/api/posting-window/current	323.487	0.080	0.416	0.857	237
0.7%	1033	0	0	POST	/api/auth/logout	291.292	0.282	2.492	4.608	0
0.7%	369	0	0	GET	/api/me/export	282.509	0.766	5.161	14.051	1843576
0.7%	568	0	0	PATCH	/api/me	266.412	0.469	2.104	3.255	281
0.5%	685	3	0	GET	/api/me	212.826	0.311	0.842	1.060	301
0.5%	2192	23	0	GET	^/api/posts/me/current$	201.142	0.092	0.611	1.025	648
0.4%	239	0	0	PATCH	^/api/posts/[^/]+$	167.956	0.703	2.604	2.919	652
0.4%	2131	31	0	GET	^/api/posts/[^/]+$	156.076	0.073	0.596	0.945	614
0.4%	1766	0	0	GET	^/api/relationships/[^/]+	145.349	0.082	0.768	2.015	1024
0.3%	948	0	0	GET	^/api/memories/[^/]+	109.043	0.115	1.133	3.118	672
0.3%	94	0	0	DELETE	^/api/posts/[^/]+$	106.671	1.135	2.814	2.814	0
0.2%	156	0	0	DELETE	^/api/comments/[^/]+$	67.998	0.436	3.108	3.493	0
0.2%	1022	0	0	GET	^/api/users/[^/]+$	67.055	0.066	0.611	0.841	198
0.2%	116	0	0	DELETE	^/api/realmojis/[^/]+$	65.051	0.561	2.390	2.441	0
0.2%	767	0	0	GET	^/api/posts/[^/]+/realmojis$	63.833	0.083	0.680	1.811	1487
0.1%	652	0	0	GET	^/api/posts/[^/]+/comments$	51.292	0.079	0.606	1.576	2445
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.6%	693	1	0	GET	/api/notifications	5228.713	7.545	11.514	11.515	138
23.3%	1067	99	0	POST	/api/posts	3745.055	3.510	65.721	66.362	592
11.0%	2350	44	0	POST	/api/friend-requests	1771.476	0.754	3.128	4.274	113
4.9%	1383	0	0	POST	/api/auth/login	784.515	0.567	2.632	3.389	243
3.7%	3891	0	0	GET	^/api/feeds/[^/]+	598.812	0.154	1.604	2.718	3164
3.6%	4182	33	0	GET	^/api/images/[^/]+	575.108	0.138	1.309	2.789	172395
2.9%	510	0	0	POST	/api/auth/register	459.820	0.902	2.578	2.649	228
2.0%	367	0	0	PUT	^/api/follows/[^/]+$	323.274	0.881	3.021	3.659	0
2.0%	292	0	0	PUT	/api/me/profile-image	313.440	1.073	3.032	3.399	0
1.9%	546	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	298.800	0.547	2.866	3.894	115
1.6%	243	0	0	PATCH	/api/me	257.128	1.058	3.323	4.304	281
1.6%	1108	0	0	GET	/api/memories	256.842	0.232	2.115	3.170	13121
1.3%	193	7	0	POST	^/api/posts/[^/]+/realmojis$	210.737	1.092	4.088	5.609	464
1.3%	1599	0	0	GET	/api/posting-window/current	207.303	0.130	1.147	2.307	238
1.0%	774	33	0	GET	^/api/posts/me/current$	158.581	0.205	1.738	2.307	625
1.0%	1178	0	0	GET	/api/users	155.222	0.132	1.321	2.306	240
0.9%	173	2	0	POST	^/api/posts/[^/]+/comments$	137.249	0.793	4.433	4.727	328
0.6%	195	0	0	POST	/api/auth/logout	104.124	0.534	2.983	3.046	0
0.5%	535	0	0	GET	^/api/relationships/[^/]+	79.664	0.149	1.701	1.915	633
0.5%	83	1	0	PATCH	^/api/posts/[^/]+$	76.250	0.919	3.478	3.478	644
0.3%	490	14	0	GET	^/api/posts/[^/]+$	53.062	0.108	0.343	1.923	615
0.3%	39	0	0	DELETE	^/api/realmojis/[^/]+$	46.290	1.187	4.485	4.485	0
0.3%	38	0	0	DELETE	^/api/posts/[^/]+$	42.693	1.123	3.367	3.367	0
0.3%	297	2	0	GET	/api/me	40.226	0.135	0.343	1.020	303
0.2%	224	0	0	GET	^/api/users/[^/]+$	26.037	0.116	0.935	1.964	204
0.2%	29	0	0	DELETE	^/api/comments/[^/]+$	25.886	0.893	2.594	2.594	0
0.1%	205	0	0	GET	^/api/posts/[^/]+/realmojis$	23.353	0.114	0.807	1.406	1935
0.1%	168	0	0	GET	^/api/memories/[^/]+	22.208	0.132	1.856	2.502	1007
0.1%	153	0	0	GET	^/api/posts/[^/]+/comments$	16.355	0.107	0.793	0.803	2756
0.1%	62	0	0	GET	/api/me/export	15.495	0.250	4.897	4.897	1824219
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.3%	488	0	0	GET	/api/notifications	3641.080	7.461	10.762	11.077	156
31.3%	684	81	0	POST	/api/posts	3524.832	5.153	66.117	66.293	577
8.2%	1392	22	0	POST	/api/friend-requests	918.470	0.660	2.884	4.300	114
4.9%	1072	0	0	POST	/api/auth/login	552.286	0.515	2.640	4.646	249
3.8%	3378	12	0	GET	^/api/images/[^/]+	424.426	0.126	0.920	4.187	137212
2.7%	2234	0	0	GET	^/api/feeds/[^/]+	305.874	0.137	1.108	1.583	2737
2.0%	280	0	0	POST	/api/auth/register	230.745	0.824	2.515	2.561	230
1.9%	349	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	208.566	0.598	2.723	3.239	116
1.8%	875	0	0	GET	/api/memories	198.082	0.226	2.373	4.520	17561
1.4%	157	0	0	PUT	/api/me/profile-image	159.158	1.014	3.224	3.691	0
1.4%	148	0	0	PATCH	/api/me	152.340	1.029	3.489	3.542	282
1.2%	185	0	0	PUT	^/api/follows/[^/]+$	130.380	0.705	2.724	3.027	0
1.1%	980	0	0	GET	/api/posting-window/current	125.186	0.128	1.100	1.392	240
0.9%	174	0	0	POST	/api/auth/logout	98.740	0.567	2.725	3.110	0
0.9%	109	1	0	POST	^/api/posts/[^/]+/comments$	96.513	0.885	3.449	3.568	349
0.8%	469	27	0	GET	^/api/posts/me/current$	94.632	0.202	1.163	1.306	617
0.7%	685	0	0	GET	/api/users	83.389	0.122	0.966	1.108	239
0.7%	70	2	0	POST	^/api/posts/[^/]+/realmojis$	76.178	1.088	3.824	3.824	440
0.3%	38	3	0	PATCH	^/api/posts/[^/]+$	34.220	0.901	2.918	2.918	600
0.3%	275	0	0	GET	^/api/relationships/[^/]+	32.126	0.117	1.135	1.280	941
0.3%	276	18	0	GET	^/api/posts/[^/]+$	28.389	0.103	0.912	1.102	586
0.2%	15	0	0	DELETE	^/api/posts/[^/]+$	22.690	1.513	3.099	3.099	0
0.2%	161	0	0	GET	/api/me	20.253	0.126	1.008	1.012	304
0.2%	51	0	0	GET	/api/me/export	17.520	0.344	5.362	5.362	3093642
0.2%	137	0	0	GET	^/api/users/[^/]+$	17.165	0.125	0.814	1.291	199
0.1%	144	0	0	GET	^/api/memories/[^/]+	16.840	0.117	1.165	1.284	1081
0.1%	17	0	0	DELETE	^/api/realmojis/[^/]+$	11.393	0.670	1.423	1.423	0
0.1%	23	0	0	DELETE	^/api/comments/[^/]+$	10.226	0.445	1.992	1.992	0
0.1%	97	0	0	GET	^/api/posts/[^/]+/realmojis$	9.266	0.096	0.328	0.328	1673
0.1%	100	0	0	GET	^/api/posts/[^/]+/comments$	9.211	0.092	0.260	0.262	1654
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16261 クエリ / 48 種類 / 合計 142.201s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	52.0%	380	73.956	194.62	1017.95	1170.4	119.9	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	15.0%	3	21.389	7129.62	9576.76	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	12.0%	13737	17.078	1.24	6.85	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	11.2%	3	15.959	5319.56	7197.45	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.6%	3	7.931	2643.56	3616.09	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	1.9%	6	2.664	443.93	1802.61	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.7%	6	1.040	173.40	528.79	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.5%	6	0.673	112.24	158.18	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.251	251.05	251.05	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	0.2%	1	0.249	249.06	249.06	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.230	229.64	229.64	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.169	169.06	169.06	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.104	103.67	103.67	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.083	82.87	82.87	2412.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.070	69.95	69.95	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: 4652f906c9b030c81aabad690ac9e67388a28e0e
Type: cpu
Time: 2026-10-31 15:50:00 JST
Duration: 60.01s, Total samples = 43.69s (72.81%)
Showing nodes accounting for 31.32s, 71.69% of 43.69s total
Dropped 1075 nodes (cum <= 0.22s)
Showing top 40 nodes out of 281
      flat  flat%   sum%        cum   cum%
     6.17s 14.12% 14.12%      6.17s 14.12%  crypto/internal/fips140/sha256.blockAVX2
        5s 11.44% 25.57%         5s 11.44%  internal/runtime/syscall/linux.Syscall6
     1.90s  4.35% 29.92%      1.90s  4.35%  runtime.memmove
     1.63s  3.73% 33.65%      2.15s  4.92%  runtime.findObject
     1.59s  3.64% 37.29%      1.59s  3.64%  runtime.memclrNoHeapPointers
     1.57s  3.59% 40.88%      5.76s 13.18%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.56s  3.57% 44.45%      1.56s  3.57%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.52s  3.48% 47.93%      6.25s 14.31%  runtime.scanObject
     1.41s  3.23% 51.16%      1.41s  3.23%  aeshashbody
     1.41s  3.23% 54.38%      2.36s  5.40%  runtime.tryDeferToSpanScan
     0.74s  1.69% 56.08%      0.74s  1.69%  runtime.spanClass.sizeclass (inline)
     0.66s  1.51% 57.59%      2.34s  5.36%  runtime.mapaccess1_faststr
     0.53s  1.21% 58.80%      0.70s  1.60%  example.com/isureal-company/isureal.postLess
     0.43s  0.98% 59.78%      0.43s  0.98%  memeqbody
     0.34s  0.78% 60.56%      6.84s 15.66%  crypto/internal/fips140/sha256.(*Digest).Write
     0.33s  0.76% 61.32%      0.38s  0.87%  encoding/json.appendString[go.shape.string]
     0.32s  0.73% 62.05%      6.89s 15.77%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.30s  0.69% 62.74%      0.30s  0.69%  runtime.futex
     0.27s  0.62% 63.36%      0.27s  0.62%  runtime.nextFreeFast (inline)
     0.26s   0.6% 63.95%      0.26s   0.6%  sync/atomic.(*Int32).Add (inline)
     0.25s  0.57% 64.52%      4.21s  9.64%  runtime.mallocgc
     0.24s  0.55% 65.07%      0.68s  1.56%  database/sql.convertAssignRows
     0.24s  0.55% 65.62%      0.59s  1.35%  runtime.greyobject
     0.24s  0.55% 66.17%      0.29s  0.66%  runtime.typePointers.next
     0.23s  0.53% 66.70%      0.30s  0.69%  runtime.spanOf (inline)
     0.21s  0.48% 67.18%      0.48s  1.10%  compress/flate.(*compressor).deflate
     0.19s  0.43% 67.61%      7.74s 17.72%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.18s  0.41% 68.02%      0.23s  0.53%  runtime.(*mspan).markBitsForIndex (inline)
     0.18s  0.41% 68.44%      0.96s  2.20%  runtime.mallocgcSmallScanNoHeader
     0.18s  0.41% 68.85%      2.08s  4.76%  runtime.mapassign_faststr
     0.16s  0.37% 69.21%      0.27s  0.62%  internal/runtime/maps.(*Iter).Next
     0.15s  0.34% 69.56%      7.24s 16.57%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.15s  0.34% 69.90%      0.44s  1.01%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.15s  0.34% 70.24%      0.54s  1.24%  runtime.mapaccess2_faststr
     0.11s  0.25% 70.50%      8.35s 19.11%  example.com/isureal-company/isureal.computeImageSignature
     0.11s  0.25% 70.75%      4.55s 10.41%  example.com/isureal-company/isureal.listMemories
     0.11s  0.25% 71.00%      0.73s  1.67%  runtime.newobject
     0.10s  0.23% 71.23%      2.99s  6.84%  example.com/isureal-company/isureal.loadBig
     0.10s  0.23% 71.46%      2.45s  5.61%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.23% 71.69%      0.30s  0.69%  internal/stringslite.Index
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 53.0% / 100% | 28% | 7.0% | 6 | 263 MB | run queue が積もっている |
| isu2 (2 cpu) | 32.2% / 100% | 6% | 0.6% | 6 | 268 MB | run queue が積もっている |
| isu3 (2 cpu) | 27.6% / 100% | 4% | 0.6% | 5 | 320 MB | run queue が積もっている |
| isu4 (2 cpu) | 42.1% / 100% | 16% | 15.3% | 4 | 104 MB |  |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 72% / nginx 14% / find 1% |
| isu2 | isureal 61% / nginx 2% |
| isu3 | isureal 53% / nginx 2% |
| isu4 | isureal 36% / postgres 15% |

## Diff from previous run

スコア: 0 → 183742 (+183742)

コードの差分は無い（同じ sha 4942c089b9e0）。
スコアの違いはブレか、サーバ側の設定変更による。
