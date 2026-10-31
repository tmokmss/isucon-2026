# run 20261031-171010-53bad13

## Summary

| | |
|---|---|
| score | 219729 |
| status | pass |
| ref | asia2-gogc |
| sha | 53bad13c4a4b |
| 所要時間 | 152s |
| 前回比 | -3746 (-1.7%) vs 20261031-170001-a1cea2f |
| ベスト比 | -9862 (-4.3%) vs 20261031-164120-892a769 = 229591 |
| note | asia2 GOGC=400 GOMEMLIMIT=1400MiB MemoryMax=1800M |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	137	365	678	678
europe	7	21	63	191	545	544
america	6	18	58	169	466	465
```

```
{"id":"32598","status":"done","portalStatus":"finished","passed":true,"score":219729,"raw":219729,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:10:41.413Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:10:44.401Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:10:47.400Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:10:50.414Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T08:10:53.400Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T08:10:56.400Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T08:10:59.415Z level=INFO msg="asia 3 日目: 137 人"
time=2026-10-31T08:11:01.897Z level=INFO msg="user moved" user=ullia5571 from=asia to=america
time=2026-10-31T08:11:02.401Z level=INFO msg="europe 3 日目: 63 人"
time=2026-10-31T08:11:03.168Z level=INFO msg="user moved" user=ullia5571 from=america to=europe
time=2026-10-31T08:11:03.502Z level=INFO msg="user moved" user=seyuno-5840 from=asia to=america
time=2026-10-31T08:11:04.318Z level=INFO msg="user moved" user=seyuno-5840 from=america to=europe
time=2026-10-31T08:11:04.584Z level=INFO msg="user moved" user=kazoeor-5635 from=asia to=europe
time=2026-10-31T08:11:05.401Z level=INFO msg="america 3 日目: 58 人"
time=2026-10-31T08:11:08.412Z level=INFO msg="asia 4 日目: 365 人"
time=2026-10-31T08:11:09.843Z level=INFO msg="user moved" user=maxsa-6691 from=america to=europe
time=2026-10-31T08:11:11.401Z level=INFO msg="europe 4 日目: 191 人"
time=2026-10-31T08:11:12.380Z level=INFO msg="user moved" user=maxsa-6691 from=europe to=america
time=2026-10-31T08:11:12.834Z level=INFO msg="user moved" user=ullia5571 from=europe to=asia
time=2026-10-31T08:11:14.413Z level=INFO msg="america 4 日目: 169 人"
time=2026-10-31T08:11:14.442Z level=INFO msg="user moved" user=seyuno-5840 from=europe to=asia
time=2026-10-31T08:11:16.407Z level=INFO msg="user moved" user=kazoeor-5635 from=europe to=asia
time=2026-10-31T08:11:17.191Z level=INFO msg="user moved" user=moulshi_9410 from=america to=asia
time=2026-10-31T08:11:17.418Z level=INFO msg="asia 5 日目: 678 人"
time=2026-10-31T08:11:19.305Z level=INFO msg="user moved" user=sukute_9863 from=america to=asia
time=2026-10-31T08:11:20.433Z level=INFO msg="europe 5 日目: 545 人"
time=2026-10-31T08:11:23.552Z level=INFO msg="america 5 日目: 466 人"
time=2026-10-31T08:11:32.477Z level=INFO msg="europe 最終日（グランプリ）: 544 人"
time=2026-10-31T08:11:32.493Z level=INFO msg="america 最終日（グランプリ）: 465 人"
time=2026-10-31T08:11:32.704Z level=INFO msg="asia 最終日（グランプリ）: 678 人"
time=2026-10-31T08:11:41.509Z level=INFO msg="グランプリ: アジア 407 人, ヨーロッパ 297 人, アメリカ 268 人"
{
  "pass": true,
  "score": 219729,
  "messages": [
    "leokimo.13620: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "fuha.8177: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nonu-15015: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tadanlia.7449: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rosu_14981: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nitsutsu7780: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rohi_9614: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "dalbam_mechurial_991: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "nunoa-5118: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.42.172:56616-\u003e52.69.202.109:8080: write tcp 10.27.42.172:56616-\u003e52.69.202.109:8080: use of closed network connection",
    "isisfu15321: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.42.172:55672-\u003e52.69.202.109:8080: write tcp 10.27.42.172:55672-\u003e52.69.202.109:8080: use of closed network connection",
    "hobakbit_sujokgwan_3266: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "reya12923: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.42.172:50358-\u003e52.69.202.109:8080: write tcp 10.27.42.172:50358-\u003e52.69.202.109:8080: write: broken pipe",
    "kianko12849: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
80.0%	1947	0	0	GET	/api/notifications	52971.982	27.207	60.001	60.002	114
4.4%	9977	0	0	GET	^/api/feeds/[^/]+	2884.553	0.289	1.521	2.483	5524
3.1%	3236	0	0	GET	/api/memories	2029.948	0.627	4.553	8.375	11281
2.0%	1906	51	0	POST	/api/posts	1316.140	0.691	6.811	7.262	636
1.8%	2645	0	0	GET	/api/posting-window/current	1219.718	0.461	1.681	1.844	236
1.4%	1953	0	0	GET	/api/users	931.032	0.477	1.659	2.054	239
1.1%	1533	1	0	GET	^/api/posts/me/current$	727.780	0.475	1.632	1.982	653
0.7%	243	0	0	GET	/api/me/export	479.853	1.975	9.025	9.471	1679326
0.7%	517	127	0	PUT	/api/me/profile-image	479.839	0.928	11.443	15.863	4
0.6%	12737	247	0	GET	^/api/images/[^/]+	428.245	0.034	0.584	1.939	140257
0.5%	1310	0	0	GET	^/api/relationships/[^/]+	358.261	0.273	1.411	2.782	1654
0.5%	4019	55	0	POST	/api/friend-requests	319.076	0.079	1.047	1.964	114
0.4%	698	9	0	GET	^/api/posts/[^/]+/realmojis$	290.999	0.417	2.712	3.641	2026
0.4%	3979	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	278.242	0.070	1.007	2.349	116
0.4%	1594	70	0	GET	^/api/posts/[^/]+$	270.296	0.170	1.085	1.357	597
0.4%	398	3	0	GET	/api/me	249.391	0.627	1.727	1.983	301
0.3%	679	0	0	GET	^/api/memories/[^/]+	205.328	0.302	1.588	1.957	658
0.2%	557	18	0	POST	^/api/posts/[^/]+/realmojis$	164.247	0.295	1.865	2.351	467
0.2%	4337	460	0	POST	/api/auth/login	137.151	0.032	0.959	2.498	221
0.2%	677	7	0	GET	^/api/posts/[^/]+/comments$	131.671	0.194	1.005	1.526	3034
0.2%	787	0	0	GET	^/api/users/[^/]+$	119.370	0.152	0.963	1.552	199
0.1%	771	1	0	POST	^/api/posts/[^/]+/comments$	83.386	0.108	1.445	2.674	359
0.1%	68	0	0	DELETE	^/api/posts/[^/]+$	34.231	0.503	1.536	1.536	0
0.0%	697	0	0	POST	/api/auth/logout	31.440	0.045	1.011	1.442	0
0.0%	159	1	0	PATCH	^/api/posts/[^/]+$	29.015	0.182	1.264	1.834	650
0.0%	79	0	0	DELETE	^/api/realmojis/[^/]+$	24.175	0.306	2.000	2.000	0
0.0%	1	0	0	POST	/api/initialize	21.865	21.865	21.865	21.865	90
0.0%	136	0	0	DELETE	^/api/comments/[^/]+$	7.234	0.053	0.951	1.003	0
0.0%	425	0	0	PUT	^/api/follows/[^/]+$	6.083	0.014	0.377	0.998	0
0.0%	321	0	0	PATCH	/api/me	4.951	0.015	0.696	1.176	277
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
85.3%	1383	1	0	GET	/api/notifications	37627.189	27.207	60.101	60.377	105
2.3%	8465	0	0	GET	^/api/feeds/[^/]+	1015.794	0.120	1.216	2.287	4402
1.9%	3120	37	0	POST	/api/friend-requests	855.197	0.274	1.289	1.936	114
1.9%	1371	30	0	POST	/api/posts	852.170	0.622	6.972	7.712	638
1.7%	7015	36	0	GET	^/api/images/[^/]+	755.127	0.108	1.017	3.057	187278
1.0%	2013	0	0	POST	/api/auth/login	437.007	0.217	1.253	1.856	248
0.6%	1637	0	0	GET	/api/memories	283.602	0.173	1.940	3.076	13322
0.5%	504	15	0	POST	^/api/posts/[^/]+/realmojis$	209.904	0.416	2.398	2.806	463
0.4%	1997	0	0	GET	/api/posting-window/current	197.930	0.099	0.417	1.339	239
0.4%	1071	0	0	GET	^/api/posts/me/current$	192.881	0.180	1.383	2.077	652
0.4%	543	0	0	POST	/api/auth/register	186.723	0.344	1.432	1.501	229
0.4%	1430	0	0	GET	/api/users	161.799	0.113	0.468	1.958	243
0.4%	308	0	0	PUT	/api/me/profile-image	160.339	0.521	1.774	2.422	0
0.3%	642	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	145.170	0.226	1.220	1.912	116
0.3%	1381	39	0	GET	^/api/posts/[^/]+$	125.944	0.091	0.448	1.940	606
0.3%	386	0	0	PUT	^/api/follows/[^/]+$	124.443	0.322	1.318	1.734	0
0.3%	535	7	0	POST	^/api/posts/[^/]+/comments$	120.232	0.225	1.338	2.457	349
0.3%	981	0	0	GET	^/api/relationships/[^/]+	118.272	0.121	1.420	1.960	547
0.2%	245	0	0	PATCH	/api/me	92.192	0.376	1.395	2.002	276
0.2%	718	0	0	GET	^/api/users/[^/]+$	66.744	0.093	0.447	1.591	202
0.1%	356	0	0	POST	/api/auth/logout	63.575	0.179	1.229	1.911	0
0.1%	619	1	0	GET	^/api/posts/[^/]+/realmojis$	55.742	0.090	0.459	0.724	3174
0.1%	558	0	0	GET	^/api/posts/[^/]+/comments$	50.431	0.090	0.395	1.640	3450
0.1%	109	1	0	PATCH	^/api/posts/[^/]+$	40.826	0.375	1.458	1.630	639
0.1%	318	2	0	GET	/api/me	35.748	0.112	0.463	0.602	302
0.1%	79	0	0	DELETE	^/api/realmojis/[^/]+$	30.241	0.383	1.347	1.347	0
0.1%	277	0	0	GET	^/api/memories/[^/]+	28.933	0.104	0.785	2.118	898
0.1%	41	0	0	DELETE	^/api/posts/[^/]+$	24.271	0.592	2.302	2.302	0
0.0%	114	0	0	DELETE	^/api/comments/[^/]+$	17.938	0.157	1.063	1.068	0
0.0%	110	0	0	GET	/api/me/export	11.767	0.107	0.597	0.643	2260830
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
85.9%	1191	0	0	GET	/api/notifications	32183.040	27.022	60.004	60.143	107
2.4%	1181	23	0	POST	/api/posts	890.441	0.754	7.119	7.738	649
2.1%	5181	0	0	GET	^/api/feeds/[^/]+	796.070	0.154	1.474	2.290	3430
2.0%	6039	37	0	GET	^/api/images/[^/]+	760.915	0.126	1.464	2.744	152844
1.3%	2581	18	0	POST	/api/friend-requests	469.934	0.182	1.176	1.392	114
1.1%	1812	0	0	POST	/api/auth/login	428.386	0.236	1.284	1.848	250
0.7%	467	0	0	POST	/api/auth/register	263.038	0.563	1.426	1.481	230
0.6%	1476	0	0	GET	/api/memories	238.281	0.161	1.675	2.866	17005
0.6%	1694	0	0	GET	/api/posting-window/current	225.965	0.133	1.349	1.589	240
0.4%	1292	0	0	GET	/api/users	156.606	0.121	1.255	1.669	242
0.4%	917	2	0	GET	^/api/posts/me/current$	150.860	0.165	1.408	1.834	661
0.3%	268	19	0	POST	^/api/posts/[^/]+/realmojis$	97.150	0.363	1.330	2.140	448
0.2%	585	0	0	GET	^/api/relationships/[^/]+	93.544	0.160	1.502	1.954	565
0.2%	425	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	84.448	0.199	1.165	1.487	116
0.2%	277	0	0	PUT	/api/me/profile-image	80.671	0.291	0.757	0.770	0
0.2%	676	24	0	GET	^/api/posts/[^/]+$	74.493	0.110	1.133	1.401	601
0.2%	351	18	0	POST	^/api/posts/[^/]+/comments$	72.017	0.205	1.177	1.250	343
0.2%	305	0	0	POST	/api/auth/logout	59.575	0.195	1.178	1.322	0
0.1%	295	0	0	PUT	^/api/follows/[^/]+$	43.732	0.148	0.628	0.635	0
0.1%	312	0	0	GET	^/api/users/[^/]+$	39.135	0.125	1.360	1.390	197
0.1%	279	0	0	GET	^/api/posts/[^/]+/comments$	35.007	0.125	1.213	1.397	3204
0.1%	225	0	0	PATCH	/api/me	32.342	0.144	0.391	0.576	281
0.1%	259	4	0	GET	^/api/posts/[^/]+/realmojis$	32.279	0.125	1.319	1.948	3163
0.1%	262	0	0	GET	^/api/memories/[^/]+	28.083	0.107	0.696	2.412	985
0.1%	83	1	0	PATCH	^/api/posts/[^/]+$	26.183	0.315	1.642	1.642	647
0.1%	45	0	0	DELETE	^/api/posts/[^/]+$	24.801	0.551	2.146	2.146	0
0.1%	109	0	0	GET	/api/me/export	24.470	0.224	2.376	5.832	2952880
0.1%	43	0	0	DELETE	^/api/realmojis/[^/]+$	21.637	0.503	1.510	1.510	0
0.1%	284	1	0	GET	/api/me	21.492	0.076	0.328	0.341	304
0.0%	47	0	0	DELETE	^/api/comments/[^/]+$	10.957	0.233	1.234	1.234	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 12913 クエリ / 48 種類 / 合計 136.679s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	42.0%	248	57.443	231.62	840.15	1070.5	109.0	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	18.7%	4	25.572	6392.91	9532.81	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.3%	4	19.527	4881.72	7189.11	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	10.8%	8500	14.739	1.73	16.18	10.0	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.7%	4	9.152	2288.08	3426.09	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	4.2%	7	5.673	810.42	3083.33	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.6%	1	0.810	810.14	810.14	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.5%	7	0.627	89.54	159.05	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.4%	7	0.598	85.39	161.68	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.4%	1	0.539	538.55	538.55	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.3%	1	0.370	370.41	370.41	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.3%	1	0.346	346.45	346.45	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
13	0.2%	1	0.327	327.09	327.09	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.173	173.37	173.37	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
15	0.1%	7	0.135	19.31	28.91	167.4	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM posts;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: e5079c07c958e9d908206ed53f118a8472dc7377
Type: cpu
Time: 2026-10-31 17:10:15 JST
Duration: 60s, Total samples = 25.87s (43.11%)
Showing nodes accounting for 20.14s, 77.85% of 25.87s total
Dropped 856 nodes (cum <= 0.13s)
Showing top 40 nodes out of 236
      flat  flat%   sum%        cum   cum%
     4.47s 17.28% 17.28%      4.47s 17.28%  crypto/internal/fips140/sha256.blockAVX2
     2.79s 10.78% 28.06%      2.79s 10.78%  internal/runtime/syscall/linux.Syscall6
     1.70s  6.57% 34.63%      5.68s 21.96%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.37s  5.30% 39.93%      1.37s  5.30%  runtime.memmove
     1.35s  5.22% 45.15%      1.35s  5.22%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.93s  3.59% 48.74%      0.93s  3.59%  aeshashbody
     0.82s  3.17% 51.91%      0.82s  3.17%  runtime.memclrNoHeapPointers
     0.57s  2.20% 54.12%      2.16s  8.35%  runtime.scanObject
     0.53s  2.05% 56.17%      0.80s  3.09%  runtime.tryDeferToSpanScan
     0.47s  1.82% 57.98%      0.64s  2.47%  runtime.findObject
     0.47s  1.82% 59.80%      1.99s  7.69%  runtime.mapaccess1_faststr
     0.40s  1.55% 61.35%      0.53s  2.05%  example.com/isureal-company/isureal.postLess
     0.39s  1.51% 62.85%      0.39s  1.51%  memeqbody
     0.32s  1.24% 64.09%      0.32s  1.24%  runtime.futex
     0.28s  1.08% 65.17%      0.28s  1.08%  runtime.nextFreeFast (inline)
     0.27s  1.04% 66.22%      0.46s  1.78%  database/sql.convertAssignRows
     0.26s  1.01% 67.22%      5.06s 19.56%  crypto/internal/fips140/sha256.(*Digest).Write
     0.21s  0.81% 68.03%      5.35s 20.68%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.20s  0.77% 68.81%      0.20s  0.77%  runtime.spanClass.sizeclass (inline)
     0.20s  0.77% 69.58%      0.20s  0.77%  sync/atomic.(*Int32).Add (inline)
     0.18s   0.7% 70.27%      5.11s 19.75%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.17s  0.66% 70.93%      0.19s  0.73%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.17s  0.66% 71.59%      0.87s  3.36%  runtime.mallocgcSmallScanNoHeader
     0.15s  0.58% 72.17%      1.80s  6.96%  runtime.mallocgc
     0.13s   0.5% 72.67%      0.15s  0.58%  runtime.(*mspan).writeHeapBitsSmall
     0.13s   0.5% 73.17%      1.33s  5.14%  runtime.mapassign_faststr
     0.12s  0.46% 73.64%      5.66s 21.88%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.12s  0.46% 74.10%      0.28s  1.08%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.11s  0.43% 74.53%      3.74s 14.46%  example.com/isureal-company/isureal.loadBig
     0.11s  0.43% 74.95%      0.43s  1.66%  runtime.mallocgcSmallNoscan
     0.10s  0.39% 75.34%      0.16s  0.62%  runtime.greyobject
     0.09s  0.35% 75.69%      6.06s 23.42%  example.com/isureal-company/isureal.computeImageSignature
     0.09s  0.35% 76.03%      0.63s  2.44%  runtime.slicebytetostring
     0.08s  0.31% 76.34%      4.55s 17.59%  crypto/internal/fips140/sha256.block
     0.08s  0.31% 76.65%      1.84s  7.11%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.08s  0.31% 76.96%      0.20s  0.77%  internal/runtime/maps.(*Iter).Next
     0.07s  0.27% 77.23%      0.19s  0.73%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.07s  0.27% 77.50%      0.34s  1.31%  runtime.mapaccess2_faststr
     0.05s  0.19% 77.70%      0.30s  1.16%  sort.insertionSort_func
     0.04s  0.15% 77.85%      0.50s  1.93%  database/sql.(*Rows).scanLocked
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 25.7% / 100% | 4% | 0.7% | 5 | 150 MB | run queue が積もっている |
| isu2 (2 cpu) | 21.3% / 100% | 5% | 0.7% | 6 | 154 MB | run queue が積もっている |
| isu3 (2 cpu) | 21.2% / 100% | 7% | 0.7% | 5 | 147 MB | run queue が積もっている |
| isu4 (2 cpu) | 54.7% / 100% | 46% | 2.6% | 9 | 102 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 40% / nginx 8% |
| isu2 | isureal 38% / nginx 3% |
| isu3 | isureal 39% / nginx 2% |
| isu4 | isureal 84% / postgres 13% |

## Diff from previous run

スコア: 223475 → 219729 (-3746)

`a1cea2f..53bad13`

```
 conf/isu4/etc/systemd/system/isureal-go.service |   5 +-
 conf/isu4/home/isucon/env.sh                    |   1 -
 webapp/go/cmd/isureal/main.go                   |   3 -
 webapp/go/image_signature.go                    |   1 -
 webapp/go/offload.go                            |   4 -
 webapp/go/sign_share.go                         | 105 ------------------------
 6 files changed, 3 insertions(+), 116 deletions(-)
```
