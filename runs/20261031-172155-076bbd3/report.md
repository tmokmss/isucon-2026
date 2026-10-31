# run 20261031-172155-076bbd3

## Summary

| | |
|---|---|
| score | 223494 |
| status | pass |
| ref | asia2-sign-euus |
| sha | 076bbd3dca1f |
| 所要時間 | 371s |
| 前回比 | +3765 (1.7%) vs 20261031-171010-53bad13 |
| ベスト比 | -8616 (-3.7%) vs 20261031-164055-bdf05cd = 232110 |
| note | asia2 sign offload to europe1/america1, 64 slots |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	53	158	466	1075	1074
europe	7	21	63	185	522	520
america	6	20	61	182	445	444
```

```
{"id":"34218","status":"done","portalStatus":"finished","passed":true,"score":223494,"raw":223494,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:26:05.392Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:26:08.383Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:26:11.382Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:26:14.395Z level=INFO msg="asia 2 日目: 53 人"
time=2026-10-31T08:26:15.941Z level=INFO msg="user moved" user=fueva.5229 from=asia to=america
time=2026-10-31T08:26:17.383Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T08:26:19.273Z level=INFO msg="user moved" user=fueva.5229 from=america to=europe
time=2026-10-31T08:26:20.383Z level=INFO msg="america 2 日目: 20 人"
time=2026-10-31T08:26:20.570Z level=INFO msg="user moved" user=senowa5426 from=europe to=america
time=2026-10-31T08:26:23.396Z level=INFO msg="asia 3 日目: 158 人"
time=2026-10-31T08:26:26.383Z level=INFO msg="europe 3 日目: 63 人"
time=2026-10-31T08:26:29.383Z level=INFO msg="america 3 日目: 61 人"
time=2026-10-31T08:26:32.155Z level=INFO msg="user moved" user=fueva.5229 from=europe to=asia
time=2026-10-31T08:26:32.398Z level=INFO msg="asia 4 日目: 466 人"
time=2026-10-31T08:26:33.723Z level=INFO msg="user moved" user=senowa5426 from=america to=asia
time=2026-10-31T08:26:35.390Z level=INFO msg="europe 4 日目: 185 人"
time=2026-10-31T08:26:35.776Z level=INFO msg="user moved" user=reheta7241 from=asia to=europe
time=2026-10-31T08:26:36.937Z level=INFO msg="user moved" user=tsumax-6822 from=america to=asia
time=2026-10-31T08:26:37.549Z level=INFO msg="user moved" user=miaku7928 from=asia to=europe
time=2026-10-31T08:26:37.912Z level=INFO msg="user moved" user=zoema-6956 from=asia to=america
time=2026-10-31T08:26:38.395Z level=INFO msg="america 4 日目: 182 人"
time=2026-10-31T08:26:39.005Z level=INFO msg="user moved" user=tosa_8764 from=asia to=america
time=2026-10-31T08:26:39.038Z level=INFO msg="user moved" user=orfu.9071 from=europe to=america
time=2026-10-31T08:26:40.574Z level=INFO msg="user moved" user=fueva.5229 from=asia to=america
time=2026-10-31T08:26:40.616Z level=INFO msg="user moved" user=muho_7638 from=asia to=america
time=2026-10-31T08:26:41.403Z level=INFO msg="asia 5 日目: 1075 人"
time=2026-10-31T08:26:44.399Z level=INFO msg="europe 5 日目: 522 人"
time=2026-10-31T08:26:45.079Z level=INFO msg="user moved" user=ishe.9311 from=europe to=america
time=2026-10-31T08:26:47.789Z level=INFO msg="america 5 日目: 445 人"
time=2026-10-31T08:26:56.459Z level=INFO msg="america 最終日（グランプリ）: 444 人"
time=2026-10-31T08:26:56.582Z level=INFO msg="asia 最終日（グランプリ）: 1074 人"
time=2026-10-31T08:26:56.716Z level=INFO msg="europe 最終日（グランプリ）: 520 人"
time=2026-10-31T08:27:05.525Z level=INFO msg="グランプリ: アジア 487 人, ヨーロッパ 228 人, アメリカ 209 人"
{
  "pass": true,
  "score": 223494,
  "messages": [
    "torka.5016: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "heivy-16931: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kevin_6430: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "koivyya.11285: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yuto_tanaka_2313: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "hoshi_12479: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.42.137:57222-\u003e52.69.202.109:8080: write tcp 10.27.42.137:57222-\u003e52.69.202.109:8080: write: broken pipe",
    "evachilia11196: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.42.137:53254-\u003e52.69.202.109:8080: write tcp 10.27.42.137:53254-\u003e52.69.202.109:8080: use of closed network connection",
    "elias_hoffmann_1411: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
80.4%	2866	1	0	GET	/api/notifications	81021.634	28.270	60.007	60.007	109
3.9%	12791	0	0	GET	^/api/feeds/[^/]+	3949.907	0.309	1.334	2.652	4250
2.8%	4548	0	0	GET	/api/memories	2848.494	0.626	2.480	3.873	12559
2.6%	2845	176	0	POST	/api/posts	2647.069	0.930	7.592	8.086	617
1.4%	4006	0	0	GET	/api/posting-window/current	1381.447	0.345	1.649	2.046	236
1.3%	3023	0	0	GET	/api/users	1277.774	0.423	1.622	2.110	255
1.1%	19163	308	0	GET	^/api/images/[^/]+	1126.888	0.059	0.694	2.846	130782
0.8%	790	131	0	PUT	/api/me/profile-image	806.832	1.021	13.921	18.224	2
0.8%	2106	1	0	GET	^/api/posts/me/current$	805.782	0.383	1.378	2.044	657
0.7%	6167	78	0	POST	/api/friend-requests	703.855	0.114	0.681	1.277	114
0.5%	390	0	0	GET	/api/me/export	536.266	1.375	5.588	10.515	1779669
0.5%	2000	106	0	GET	^/api/posts/[^/]+$	469.014	0.235	1.080	1.436	602
0.4%	1606	0	0	GET	^/api/relationships/[^/]+	436.499	0.272	1.154	1.703	1327
0.3%	671	3	0	GET	/api/me	339.616	0.506	1.871	2.167	302
0.3%	902	41	0	POST	^/api/posts/[^/]+/realmojis$	333.769	0.370	1.378	1.812	459
0.3%	910	0	0	GET	^/api/memories/[^/]+	311.408	0.342	1.816	2.617	643
0.3%	6367	446	0	POST	/api/auth/login	274.761	0.043	0.525	1.224	228
0.3%	868	9	0	GET	^/api/posts/[^/]+/realmojis$	270.555	0.312	1.629	1.983	2175
0.3%	3477	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	268.005	0.077	0.620	0.930	116
0.2%	1049	1	0	GET	^/api/users/[^/]+$	248.934	0.237	1.055	1.278	216
0.2%	742	7	0	GET	^/api/posts/[^/]+/comments$	190.554	0.257	0.960	1.361	2354
0.1%	851	1	0	POST	^/api/posts/[^/]+/comments$	139.787	0.164	0.789	1.587	357
0.1%	1003	0	0	POST	/api/auth/logout	60.235	0.060	0.631	0.876	0
0.1%	105	0	0	DELETE	^/api/posts/[^/]+$	55.315	0.527	1.273	1.388	0
0.0%	525	0	0	PATCH	/api/me	43.844	0.084	0.491	0.659	284
0.0%	127	0	0	DELETE	^/api/realmojis/[^/]+$	38.765	0.305	1.087	1.282	0
0.0%	732	0	0	PUT	^/api/follows/[^/]+$	38.261	0.052	0.470	0.627	0
0.0%	1095	1	0	POST	/api/auth/register	34.790	0.032	0.324	0.374	227
0.0%	227	1	0	PATCH	^/api/posts/[^/]+$	31.934	0.141	0.679	0.683	647
0.0%	1	0	0	POST	/api/initialize	22.310	22.310	22.310	22.310	90
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
83.3%	1327	1	0	GET	/api/notifications	36450.093	27.468	60.001	60.104	108
2.7%	1313	83	0	POST	/api/posts	1193.610	0.909	7.525	7.793	615
2.3%	5545	0	0	GET	^/api/feeds/[^/]+	1000.076	0.180	1.306	1.855	3360
2.1%	2897	37	0	POST	/api/friend-requests	933.562	0.322	1.156	1.403	114
2.0%	5340	26	0	GET	^/api/images/[^/]+	893.088	0.167	1.738	4.299	180506
1.1%	1619	0	0	POST	/api/auth/login	470.571	0.291	1.209	1.610	245
0.8%	1872	0	0	GET	/api/posting-window/current	330.409	0.177	1.115	1.515	239
0.8%	1316	0	0	GET	/api/memories	329.575	0.250	1.692	2.372	13012
0.6%	520	0	0	POST	/api/auth/register	265.560	0.511	1.467	1.524	229
0.6%	1335	0	0	GET	/api/users	260.361	0.195	1.125	1.379	256
0.5%	303	0	0	PUT	/api/me/profile-image	200.604	0.662	2.055	2.090	0
0.4%	970	12	0	GET	^/api/posts/me/current$	193.094	0.199	1.504	1.711	648
0.4%	632	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	170.035	0.269	1.041	1.308	116
0.4%	288	19	0	POST	^/api/posts/[^/]+/realmojis$	158.045	0.549	1.957	2.199	460
0.3%	345	0	0	PUT	^/api/follows/[^/]+$	116.172	0.337	1.291	1.372	0
0.2%	688	0	0	GET	^/api/relationships/[^/]+	100.714	0.146	1.081	1.612	671
0.2%	281	0	0	PATCH	/api/me	100.694	0.358	1.196	1.748	280
0.2%	670	46	0	GET	^/api/posts/[^/]+$	99.400	0.148	1.060	1.516	591
0.2%	270	0	0	POST	^/api/posts/[^/]+/comments$	81.143	0.301	1.227	1.934	336
0.2%	310	2	0	GET	/api/me	75.049	0.242	1.211	1.575	303
0.1%	236	0	0	POST	/api/auth/logout	51.983	0.220	0.852	1.007	0
0.1%	359	0	0	GET	^/api/users/[^/]+$	48.003	0.134	0.996	1.076	211
0.1%	112	3	0	PATCH	^/api/posts/[^/]+$	44.216	0.395	1.185	1.448	634
0.1%	283	1	0	GET	^/api/posts/[^/]+/realmojis$	41.951	0.148	1.250	1.574	2620
0.1%	227	0	0	GET	^/api/memories/[^/]+	31.462	0.139	1.306	1.828	998
0.1%	238	1	0	GET	^/api/posts/[^/]+/comments$	31.114	0.131	1.009	1.288	2859
0.1%	31	0	0	DELETE	^/api/posts/[^/]+$	24.294	0.784	1.796	1.796	0
0.0%	41	0	0	DELETE	^/api/realmojis/[^/]+$	19.265	0.470	0.886	0.886	0
0.0%	86	0	0	GET	/api/me/export	19.187	0.223	4.614	4.614	2325103
0.0%	56	0	0	DELETE	^/api/comments/[^/]+$	13.458	0.240	1.013	1.013	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
84.6%	1161	0	0	GET	/api/notifications	30032.052	25.867	60.306	60.314	109
2.2%	5039	44	0	GET	^/api/images/[^/]+	784.385	0.156	1.612	3.183	169747
2.2%	1112	45	0	POST	/api/posts	783.271	0.704	7.605	7.870	626
2.1%	2504	21	0	POST	/api/friend-requests	739.508	0.295	1.015	1.371	114
2.0%	4087	0	0	GET	^/api/feeds/[^/]+	715.418	0.175	1.346	1.746	3609
1.1%	1511	0	0	POST	/api/auth/login	376.556	0.249	0.912	1.170	249
0.7%	1245	0	0	GET	/api/memories	237.542	0.191	1.817	2.520	15475
0.6%	1626	0	0	GET	/api/posting-window/current	229.073	0.141	0.744	1.433	239
0.5%	1174	0	0	GET	/api/users	182.876	0.156	0.909	1.501	257
0.5%	439	0	0	POST	/api/auth/register	164.791	0.375	0.871	0.899	230
0.4%	546	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	156.494	0.287	0.971	1.315	116
0.4%	268	24	0	POST	^/api/posts/[^/]+/realmojis$	154.569	0.577	1.994	2.346	442
0.3%	245	0	0	PUT	/api/me/profile-image	124.130	0.507	1.264	2.267	0
0.3%	861	27	0	GET	^/api/posts/me/current$	122.004	0.142	1.079	1.733	633
0.3%	514	0	0	GET	^/api/relationships/[^/]+	97.695	0.190	1.388	1.751	773
0.2%	264	0	0	PUT	^/api/follows/[^/]+$	80.591	0.305	0.943	0.953	0
0.2%	555	25	0	GET	^/api/posts/[^/]+$	68.422	0.123	0.724	1.227	602
0.1%	214	0	0	POST	/api/auth/logout	50.984	0.238	0.988	1.409	0
0.1%	188	0	0	PATCH	/api/me	50.717	0.270	0.953	1.062	277
0.1%	250	1	0	GET	/api/me	47.411	0.190	0.588	0.764	304
0.1%	185	6	0	POST	^/api/posts/[^/]+/comments$	45.270	0.245	1.027	1.247	295
0.1%	51	0	0	DELETE	^/api/posts/[^/]+$	35.687	0.700	1.876	1.876	0
0.1%	253	1	0	GET	^/api/users/[^/]+$	32.566	0.129	0.859	1.611	209
0.1%	210	0	0	GET	^/api/memories/[^/]+	32.317	0.154	1.687	2.053	890
0.1%	92	3	0	PATCH	^/api/posts/[^/]+$	32.310	0.351	1.209	1.209	627
0.1%	232	0	0	GET	^/api/posts/[^/]+/realmojis$	30.393	0.131	0.635	1.611	2104
0.1%	91	0	0	GET	/api/me/export	29.377	0.323	4.217	4.217	2907256
0.1%	213	0	0	GET	^/api/posts/[^/]+/comments$	23.425	0.110	0.619	0.688	2159
0.1%	38	0	0	DELETE	^/api/realmojis/[^/]+$	21.782	0.573	1.083	1.083	0
0.0%	44	0	0	DELETE	^/api/comments/[^/]+$	9.763	0.222	1.039	1.039	0
…（上位 30 / 全 36 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 17761 クエリ / 48 種類 / 合計 181.768s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	52.1%	397	94.722	238.60	863.22	1135.3	116.0	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.5%	4	26.387	6596.65	10034.52	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	12.1%	13206	21.960	1.66	14.86	9.7	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	11.0%	4	19.980	4995.02	8485.06	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.0%	4	9.108	2276.91	3440.60	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	3.1%	7	5.654	807.75	3154.02	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	1	0.644	643.82	643.82	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
8	0.3%	7	0.614	87.75	160.64	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.3%	7	0.588	83.99	157.34	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.2%	1	0.381	381.44	381.44	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.2%	1	0.354	354.26	354.26	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.2%	1	0.327	327.34	327.34	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.233	233.49	233.49	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.127	126.90	126.90	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
15	0.1%	1	0.114	114.26	114.26	2463.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…

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
Build ID: f43d10437ba62e6fe2256ce011a10acf99c20477
Type: cpu
Time: 2026-10-31 17:25:38 JST
Duration: 60.04s, Total samples = 29.53s (49.19%)
Showing nodes accounting for 23.16s, 78.43% of 29.53s total
Dropped 862 nodes (cum <= 0.15s)
Showing top 40 nodes out of 244
      flat  flat%   sum%        cum   cum%
     6.19s 20.96% 20.96%      6.19s 20.96%  crypto/internal/fips140/sha256.blockAVX2
     3.17s 10.73% 31.70%      3.17s 10.73%  internal/runtime/syscall/linux.Syscall6
     1.71s  5.79% 37.49%      1.71s  5.79%  runtime.memmove
     1.64s  5.55% 43.04%      6.05s 20.49%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.59s  5.38% 48.43%      1.59s  5.38%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.02s  3.45% 51.88%      1.02s  3.45%  aeshashbody
     0.71s  2.40% 54.28%      0.71s  2.40%  runtime.memclrNoHeapPointers
     0.57s  1.93% 56.21%      2.24s  7.59%  runtime.scanObject
     0.56s  1.90% 58.11%      0.79s  2.68%  runtime.findObject
     0.47s  1.59% 59.70%      2.15s  7.28%  runtime.mapaccess1_faststr
     0.40s  1.35% 61.06%      0.67s  2.27%  runtime.tryDeferToSpanScan
     0.37s  1.25% 62.31%      7.08s 23.98%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.36s  1.22% 63.53%      0.36s  1.22%  memeqbody
     0.33s  1.12% 64.65%      6.94s 23.50%  crypto/internal/fips140/sha256.(*Digest).Write
     0.33s  1.12% 65.76%      0.50s  1.69%  example.com/isureal-company/isureal.postLess
     0.24s  0.81% 66.58%      7.47s 25.30%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.23s  0.78% 67.36%      8.01s 27.12%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.23s  0.78% 68.13%      0.26s  0.88%  runtime.(*mspan).writeHeapBitsSmall
     0.22s  0.75% 68.88%      1.05s  3.56%  runtime.mallocgcSmallScanNoHeader
     0.21s  0.71% 69.59%      0.37s  1.25%  database/sql.convertAssignRows
     0.21s  0.71% 70.30%      0.21s  0.71%  sync/atomic.(*Int32).Add (inline)
     0.20s  0.68% 70.98%      0.20s  0.68%  runtime.nextFreeFast (inline)
     0.20s  0.68% 71.66%      0.20s  0.68%  runtime.spanClass.sizeclass (inline)
     0.19s  0.64% 72.30%      1.56s  5.28%  runtime.mapassign_faststr
     0.18s  0.61% 72.91%      0.20s  0.68%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.18s  0.61% 73.52%      0.18s  0.61%  runtime.futex
     0.17s  0.58% 74.09%      1.93s  6.54%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.16s  0.54% 74.64%      0.17s  0.58%  example.com/isureal-company/isureal.rmLess
     0.15s  0.51% 75.14%      0.46s  1.56%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.11s  0.37% 75.52%      0.99s  3.35%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.11s  0.37% 75.89%      1.63s  5.52%  runtime.mallocgc
     0.10s  0.34% 76.23%      6.29s 21.30%  crypto/internal/fips140/sha256.block
     0.10s  0.34% 76.57%      0.27s  0.91%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.10s  0.34% 76.90%      8.50s 28.78%  example.com/isureal-company/isureal.computeImageSignature
     0.10s  0.34% 77.24%      3.74s 12.67%  example.com/isureal-company/isureal.loadBig
     0.08s  0.27% 77.51%      0.45s  1.52%  runtime.convTstring
     0.08s  0.27% 77.79%      0.23s  0.78%  runtime.greyobject
     0.07s  0.24% 78.02%      0.50s  1.69%  sort.partition_func
     0.06s   0.2% 78.23%      0.34s  1.15%  runtime.mallocgcSmallNoscan
     0.06s   0.2% 78.43%      0.35s  1.19%  runtime.netpoll
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 31.5% / 100% | 12% | 0.9% | 5 | 102 MB | run queue が積もっている |
| isu2 (2 cpu) | 28.1% / 100% | 8% | 0.7% | 5 | 124 MB | run queue が積もっている |
| isu3 (2 cpu) | 23.6% / 100% | 8% | 1.0% | 6 | 189 MB | run queue が積もっている |
| isu4 (2 cpu) | 47.6% / 100% | 34% | 6.7% | 8 | 102 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 47% / nginx 10% |
| isu2 | isureal 51% / nginx 2% |
| isu3 | isureal 43% / nginx 2% |
| isu4 | isureal 61% / postgres 17% |

## Diff from previous run

スコア: 219729 → 223494 (+3765)

`53bad13..076bbd3`

```
 conf/isu1/home/isucon/env.sh                    |  2 +-
 conf/isu4/etc/systemd/system/isureal-go.service |  5 ++---
 conf/isu4/home/isucon/env.sh                    |  2 ++
 webapp/go/offload.go                            | 10 +++++++++-
 4 files changed, 14 insertions(+), 5 deletions(-)
```
