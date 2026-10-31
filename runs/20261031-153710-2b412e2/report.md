# run 20261031-153710-2b412e2

## Summary

| | |
|---|---|
| score | 196792 |
| status | pass |
| ref | 2b412e2 |
| sha | 2b412e2e146c |
| 所要時間 | 337s |
| 前回比 | +13461 (7.3%) vs 20261031-152137-8c13252 |
| ベスト比 | +4976 (2.6%) vs 20261031-151339-5a230a4 = 191816 **ベスト更新** |
| note | dedupe: image lookup once, shared ip_bans, admin session in oplog |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	46	131	374	1040	1042
europe	7	19	56	166	466	465
america	6	19	59	172	339	338
```

```
{"id":"23819","status":"done","portalStatus":"finished","passed":true,"score":196792,"raw":196792,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:40:51.280Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:40:54.270Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:40:57.270Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:41:00.283Z level=INFO msg="asia 2 日目: 46 人"
time=2026-10-31T06:41:00.813Z level=INFO msg="user moved" user=numarke.5105 from=europe to=america
time=2026-10-31T06:41:03.275Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T06:41:06.271Z level=INFO msg="america 2 日目: 19 人"
time=2026-10-31T06:41:09.282Z level=INFO msg="asia 3 日目: 131 人"
time=2026-10-31T06:41:10.802Z level=INFO msg="user moved" user=menomo5566 from=asia to=america
time=2026-10-31T06:41:12.221Z level=INFO msg="user moved" user=numarke.5105 from=america to=asia
time=2026-10-31T06:41:12.272Z level=INFO msg="europe 3 日目: 56 人"
time=2026-10-31T06:41:12.564Z level=INFO msg="user moved" user=numarke.5105 from=asia to=america
time=2026-10-31T06:41:14.022Z level=INFO msg="user moved" user=menomo5566 from=america to=europe
time=2026-10-31T06:41:15.271Z level=INFO msg="america 3 日目: 59 人"
time=2026-10-31T06:41:18.299Z level=INFO msg="asia 4 日目: 374 人"
time=2026-10-31T06:41:21.503Z level=INFO msg="europe 4 日目: 166 人"
time=2026-10-31T06:41:22.400Z level=INFO msg="user moved" user=kenoaso_7693 from=asia to=america
time=2026-10-31T06:41:23.484Z level=INFO msg="user moved" user=kenoaso_7693 from=america to=europe
time=2026-10-31T06:41:24.311Z level=INFO msg="america 4 日目: 172 人"
time=2026-10-31T06:41:27.344Z level=INFO msg="asia 5 日目: 1040 人"
time=2026-10-31T06:41:28.741Z level=INFO msg="user moved" user=antsute8597 from=europe to=asia
time=2026-10-31T06:41:30.723Z level=INFO msg="europe 5 日目: 466 人"
time=2026-10-31T06:41:32.194Z level=INFO msg="user moved" user=menomo5566 from=europe to=asia
time=2026-10-31T06:41:33.883Z level=INFO msg="user moved" user=tsuka.9496 from=america to=asia
time=2026-10-31T06:41:33.936Z level=INFO msg="america 5 日目: 339 人"
time=2026-10-31T06:41:42.937Z level=INFO msg="america 最終日（グランプリ）: 338 人"
time=2026-10-31T06:41:43.491Z level=INFO msg="europe 最終日（グランプリ）: 465 人"
time=2026-10-31T06:41:43.803Z level=INFO msg="asia 最終日（グランプリ）: 1042 人"
time=2026-10-31T06:41:51.342Z level=INFO msg="グランプリ: アジア 402 人, ヨーロッパ 201 人, アメリカ 189 人"
{
  "pass": true,
  "score": 196792,
  "messages": [
    "noamo15951: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rahaor6628: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "takomia_16537: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "solnoase-12220: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "toshitsu_17833: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "none.17209: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nuevator14539: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "solkatsu.16960: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hitsu_17060: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sume.17175: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kunoaben.13796: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "keevavin-17666: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "vinsoho-17079: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "issol.17758: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "isvin-7849: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "ajang_capybara_4864: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
81.2%	2680	0	0	GET	/api/notifications	76266.812	28.458	60.020	60.020	104
3.8%	5949	62	0	POST	/api/friend-requests	3543.413	0.596	2.103	2.980	114
3.6%	2616	165	0	POST	/api/posts	3357.649	1.284	6.173	6.582	612
1.9%	7264	537	0	POST	/api/auth/login	1739.695	0.239	1.820	4.041	228
1.4%	3114	6	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1350.208	0.434	1.956	2.976	116
1.4%	20767	304	0	GET	^/api/images/[^/]+	1337.966	0.064	0.802	2.983	117776
1.0%	4740	0	0	GET	/api/memories	901.232	0.190	1.796	6.768	13204
0.7%	714	122	0	PUT	/api/me/profile-image	673.044	0.943	9.148	10.967	3
0.7%	643	7	0	POST	^/api/posts/[^/]+/realmojis$	646.020	1.005	4.273	4.884	476
0.7%	16150	0	0	GET	^/api/feeds/[^/]+	616.797	0.038	0.391	0.907	3415
0.6%	750	6	0	POST	^/api/posts/[^/]+/comments$	566.028	0.755	3.411	4.952	344
0.4%	781	0	0	PUT	^/api/follows/[^/]+$	414.916	0.531	2.082	2.530	0
0.4%	1207	0	0	POST	/api/auth/logout	341.879	0.283	1.952	4.092	0
0.3%	1058	1	0	POST	/api/auth/register	313.332	0.296	1.187	1.243	227
0.3%	492	0	0	GET	/api/me/export	280.811	0.571	4.331	6.550	1927495
0.3%	516	0	0	PATCH	/api/me	252.255	0.489	1.994	2.654	277
0.2%	609	3	0	GET	/api/me	144.046	0.237	0.852	0.886	301
0.1%	2834	0	0	GET	/api/users	133.918	0.047	0.305	0.562	246
0.1%	3774	4	0	GET	/api/posting-window/current	127.896	0.034	0.257	0.912	237
0.1%	211	1	0	PATCH	^/api/posts/[^/]+$	127.597	0.605	1.899	2.139	647
0.1%	1997	1	0	GET	^/api/posts/me/current$	106.485	0.053	0.459	0.896	653
0.1%	93	0	0	DELETE	^/api/posts/[^/]+$	100.503	1.081	2.572	2.572	0
0.1%	131	0	0	DELETE	^/api/realmojis/[^/]+$	89.226	0.681	2.460	2.573	0
0.1%	159	0	0	DELETE	^/api/comments/[^/]+$	84.783	0.533	2.321	3.324	0
0.1%	2595	48	0	GET	^/api/posts/[^/]+$	83.493	0.032	0.325	0.676	610
0.1%	1123	0	0	GET	^/api/memories/[^/]+	82.577	0.074	0.664	1.240	672
0.1%	1978	0	0	GET	^/api/relationships/[^/]+	63.363	0.032	0.289	0.646	925
0.0%	1257	0	0	GET	^/api/users/[^/]+$	37.797	0.030	0.329	0.900	194
0.0%	995	0	0	GET	^/api/posts/[^/]+/realmojis$	30.646	0.031	0.326	0.608	1421
0.0%	48	0	0	PUT	^/api/blocks/[^/]+$	27.354	0.570	1.937	1.937	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
77.9%	1059	1	0	GET	/api/notifications	26502.432	25.026	60.101	60.342	102
4.8%	2483	21	0	POST	/api/friend-requests	1638.573	0.660	2.361	3.439	114
3.1%	1063	46	0	POST	/api/posts	1054.798	0.992	6.195	6.581	619
2.3%	1384	0	0	POST	/api/auth/login	767.328	0.554	2.150	2.693	245
1.8%	4575	35	0	GET	^/api/images/[^/]+	617.665	0.135	1.944	6.284	172697
1.7%	4613	0	0	GET	^/api/feeds/[^/]+	585.737	0.127	1.746	2.605	3329
1.3%	464	0	0	POST	/api/auth/register	427.299	0.921	2.047	2.073	228
0.9%	257	0	0	PUT	/api/me/profile-image	300.492	1.169	3.294	3.468	0
0.9%	472	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	298.032	0.631	2.508	2.766	116
0.7%	292	0	0	PUT	^/api/follows/[^/]+$	228.970	0.784	2.439	2.731	0
0.6%	230	0	0	PATCH	/api/me	210.536	0.915	2.622	3.687	277
0.5%	1080	0	0	GET	/api/memories	175.325	0.162	2.111	3.202	14799
0.5%	164	3	0	POST	^/api/posts/[^/]+/realmojis$	162.034	0.988	4.450	4.942	445
0.5%	212	2	0	POST	^/api/posts/[^/]+/comments$	160.185	0.756	3.725	4.516	355
0.5%	1561	6	0	GET	/api/posting-window/current	159.282	0.102	0.387	1.889	238
0.4%	1198	0	0	GET	/api/users	126.972	0.106	1.290	2.117	251
0.4%	798	45	0	GET	^/api/posts/me/current$	122.396	0.153	2.092	2.372	612
0.2%	570	0	0	GET	^/api/relationships/[^/]+	77.487	0.136	1.701	2.352	623
0.2%	191	0	0	POST	/api/auth/logout	76.404	0.400	1.934	2.510	0
0.2%	589	26	0	GET	^/api/posts/[^/]+$	63.022	0.107	1.054	2.109	591
0.2%	80	7	0	PATCH	^/api/posts/[^/]+$	53.709	0.671	2.136	2.136	592
0.1%	42	0	0	DELETE	^/api/posts/[^/]+$	45.232	1.077	2.782	2.782	0
0.1%	293	0	0	GET	^/api/users/[^/]+$	31.500	0.108	1.662	1.811	184
0.1%	31	0	0	DELETE	^/api/realmojis/[^/]+$	27.229	0.878	2.077	2.077	0
0.1%	264	2	0	GET	/api/me	22.755	0.086	0.247	0.304	302
0.1%	229	0	0	GET	^/api/posts/[^/]+/realmojis$	22.390	0.098	0.288	1.784	1849
0.1%	208	0	0	GET	^/api/posts/[^/]+/comments$	21.489	0.103	0.577	1.746	2097
0.1%	39	0	0	DELETE	^/api/comments/[^/]+$	20.381	0.523	3.059	3.059	0
0.0%	182	0	0	GET	^/api/memories/[^/]+	16.192	0.089	0.304	1.668	749
0.0%	78	0	0	GET	/api/me/export	11.345	0.145	1.147	1.147	2224831
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
76.2%	766	1	0	GET	/api/notifications	17210.919	22.469	60.102	60.407	100
4.6%	1686	21	0	POST	/api/friend-requests	1044.385	0.619	2.215	2.805	114
3.2%	757	29	0	POST	/api/posts	712.711	0.941	6.205	6.884	622
2.7%	1242	0	0	POST	/api/auth/login	619.878	0.499	2.179	3.385	249
2.5%	4060	31	0	GET	^/api/images/[^/]+	556.221	0.137	1.519	4.596	146107
1.8%	2880	0	0	GET	^/api/feeds/[^/]+	402.397	0.140	1.240	1.733	3534
1.2%	338	0	0	POST	/api/auth/register	262.812	0.778	2.143	2.190	229
1.0%	411	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	236.828	0.576	2.064	2.905	116
0.7%	230	0	0	PUT	^/api/follows/[^/]+$	160.886	0.700	2.261	2.568	0
0.7%	1015	0	0	GET	/api/memories	153.076	0.151	1.531	2.500	16761
0.7%	1127	2	0	GET	/api/posting-window/current	149.323	0.132	0.678	1.273	240
0.6%	178	0	0	PUT	/api/me/profile-image	134.713	0.757	1.754	1.775	0
0.6%	162	0	0	PATCH	/api/me	125.626	0.775	2.298	2.302	278
0.6%	118	4	0	POST	^/api/posts/[^/]+/realmojis$	124.391	1.054	4.467	5.350	456
0.5%	606	19	0	GET	^/api/posts/me/current$	107.331	0.177	1.367	1.838	625
0.5%	801	0	0	GET	/api/users	103.394	0.129	1.147	1.504	245
0.4%	117	5	0	POST	^/api/posts/[^/]+/comments$	85.118	0.728	2.788	3.231	322
0.3%	173	0	0	POST	/api/auth/logout	76.198	0.440	1.878	2.313	0
0.2%	388	0	0	GET	^/api/relationships/[^/]+	49.478	0.128	1.089	1.545	817
0.2%	366	14	0	GET	^/api/posts/[^/]+$	42.470	0.116	1.238	1.542	595
0.2%	64	3	0	PATCH	^/api/posts/[^/]+$	41.721	0.652	2.204	2.204	620
0.2%	34	0	0	DELETE	^/api/posts/[^/]+$	35.721	1.051	2.893	2.893	0
0.1%	230	0	0	GET	^/api/users/[^/]+$	25.110	0.109	1.383	1.534	187
0.1%	21	0	0	DELETE	^/api/realmojis/[^/]+$	21.598	1.028	2.489	2.489	0
0.1%	154	0	0	GET	^/api/posts/[^/]+/realmojis$	20.558	0.133	1.050	1.599	1270
0.1%	183	1	0	GET	/api/me	20.090	0.110	0.737	0.857	303
0.1%	177	0	0	GET	^/api/memories/[^/]+	18.881	0.107	0.823	2.045	1204
0.1%	71	0	0	GET	/api/me/export	18.856	0.266	3.902	3.902	2922714
0.1%	27	0	0	DELETE	^/api/comments/[^/]+$	15.531	0.575	1.321	1.321	0
0.0%	122	0	0	GET	^/api/posts/[^/]+/comments$	11.119	0.091	0.245	0.701	2069
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 18233 クエリ / 48 種類 / 合計 172.877s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	59.1%	503	102.203	203.19	846.07	1219.0	124.9	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	12.7%	3	22.027	7342.31	9904.73	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	10.7%	15060	18.551	1.23	8.01	9.3	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	9.3%	3	15.997	5332.25	7211.91	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.0%	3	8.708	2902.58	3943.49	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	1.5%	6	2.556	426.06	1714.50	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	6	0.685	114.19	162.21	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.4%	6	0.673	112.21	160.25	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.1%	1	0.255	255.07	255.07	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.1%	1	0.253	252.81	252.81	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.1%	1	0.192	191.71	191.71	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.158	157.94	157.94	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.105	105.42	105.42	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.0%	1	0.079	79.47	79.47	2486.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.079	79.30	79.30	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows

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
Build ID: 17ff8fa4fd6f531fa2ca71dd22e59668ed02ff65
Type: cpu
Time: 2026-10-31 15:40:32 JST
Duration: 60.05s, Total samples = 47.34s (78.83%)
Showing nodes accounting for 34.67s, 73.24% of 47.34s total
Dropped 1147 nodes (cum <= 0.24s)
Showing top 40 nodes out of 281
      flat  flat%   sum%        cum   cum%
     6.07s 12.82% 12.82%      6.08s 12.84%  crypto/internal/fips140/sha256.blockAVX2
     5.62s 11.87% 24.69%      5.62s 11.87%  internal/runtime/syscall/linux.Syscall6
     2.27s  4.80% 29.49%      2.27s  4.80%  runtime.memmove
     2.13s  4.50% 33.99%      7.53s 15.91%  runtime.scanObject
     1.96s  4.14% 38.13%      2.47s  5.22%  runtime.findObject
     1.86s  3.93% 42.06%      2.97s  6.27%  runtime.tryDeferToSpanScan
     1.71s  3.61% 45.67%      1.71s  3.61%  runtime.memclrNoHeapPointers
     1.66s  3.51% 49.18%      5.76s 12.17%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.54s  3.25% 52.43%      1.54s  3.25%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.32s  2.79% 55.22%      1.32s  2.79%  aeshashbody
     0.81s  1.71% 56.93%      0.81s  1.71%  runtime.spanClass.sizeclass (inline)
     0.56s  1.18% 58.11%      2.09s  4.41%  runtime.mapaccess1_faststr
     0.45s  0.95% 59.06%      0.45s  0.95%  memeqbody
     0.43s  0.91% 59.97%      0.62s  1.31%  example.com/isureal-company/isureal.postLess
     0.42s  0.89% 60.86%      6.94s 14.66%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.42s  0.89% 61.74%      0.42s  0.89%  runtime.futex
     0.37s  0.78% 62.53%      0.37s  0.78%  runtime.nextFreeFast (inline)
     0.33s   0.7% 63.22%      6.80s 14.36%  crypto/internal/fips140/sha256.(*Digest).Write
     0.30s  0.63% 63.86%      7.39s 15.61%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.30s  0.63% 64.49%      2.19s  4.63%  runtime.mapassign_faststr
     0.30s  0.63% 65.12%      0.34s  0.72%  runtime.typePointers.next
     0.29s  0.61% 65.74%      0.59s  1.25%  runtime.greyobject
     0.27s  0.57% 66.31%      0.32s  0.68%  runtime.spanOf (inline)
     0.25s  0.53% 66.84%      0.61s  1.29%  compress/flate.(*compressor).deflate
     0.25s  0.53% 67.36%      0.25s  0.53%  indexbody
     0.24s  0.51% 67.87%      0.24s  0.51%  sync/atomic.(*Int32).Add (inline)
     0.23s  0.49% 68.36%      0.30s  0.63%  runtime.(*mspan).markBitsForIndex (inline)
     0.22s  0.46% 68.82%      0.56s  1.18%  database/sql.convertAssignRows
     0.22s  0.46% 69.29%      0.24s  0.51%  runtime.(*mspan).writeHeapBitsSmall
     0.21s  0.44% 69.73%      0.24s  0.51%  encoding/json.appendString[go.shape.string]
     0.20s  0.42% 70.15%      1.25s  2.64%  runtime.mallocgcSmallScanNoHeader
     0.19s   0.4% 70.55%      0.35s  0.74%  internal/runtime/maps.(*Iter).Next
     0.19s   0.4% 70.95%      0.99s  2.09%  runtime.wbBufFlush1
     0.18s  0.38% 71.34%      4.27s  9.02%  runtime.mallocgc
     0.17s  0.36% 71.69%      2.94s  6.21%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.17s  0.36% 72.05%      0.26s  0.55%  runtime.gcmarknewobject
     0.16s  0.34% 72.39%      7.67s 16.20%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.15s  0.32% 72.71%      0.50s  1.06%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
     0.13s  0.27% 72.98%      0.52s  1.10%  runtime.mapaccess2_faststr
     0.12s  0.25% 73.24%      3.39s  7.16%  example.com/isureal-company/isureal.loadBig
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 40.6% / 100% | 26% | 0.6% | 5 | 145 MB | run queue が積もっている |
| isu2 (2 cpu) | 27.3% / 100% | 5% | 0.5% | 5 | 201 MB | run queue が積もっている |
| isu3 (2 cpu) | 23.7% / 100% | 6% | 0.5% | 5 | 211 MB | run queue が積もっている |
| isu4 (2 cpu) | 36.6% / 100% | 17% | 14.9% | 12 | 116 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 65% / nginx 12% |
| isu2 | isureal 51% / nginx 2% |
| isu3 | isureal 45% / nginx 2% |
| isu4 | isureal 26% / postgres 15% |

## Diff from previous run

スコア: 183331 → 196792 (+13461)

`8c13252..2b412e2`

```
 webapp/go/admin_api.go      | 27 +++++++++++++++++++++-----
 webapp/go/export_api.go     | 12 ++++++++++--
 webapp/go/image_api.go      | 23 ++++++++++++----------
 webapp/go/image_file.go     | 33 +++++++++++++++++--------------
 webapp/go/image_held.go     | 10 +++-------
 webapp/go/image_local.go    | 47 ++++++++++++++++++---------------------------
 webapp/go/ip_ban.go         | 11 +++++++++++
 webapp/go/persist.go        |  2 +-
 webapp/go/realmoji_api.go   | 41 +++++++++++++++------------------------
 webapp/go/repl.go           |  1 +
 webapp/go/replica_server.go | 13 +++++--------
 webapp/go/store.go          |  3 +--
 12 files changed, 121 insertions(+), 102 deletions(-)
```
