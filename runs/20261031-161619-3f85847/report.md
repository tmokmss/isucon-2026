# run 20261031-161619-3f85847

## Summary

| | |
|---|---|
| score | 194180 |
| status | pass |
| ref | asia2-offload-to-asia1 |
| sha | 3f85847efd18 |
| 所要時間 | 141s |
| 前回比 | -7317 (-3.6%) vs 20261031-160409-40e65ef |
| ベスト比 | -7317 (-3.6%) vs 20261031-160409-40e65ef = 201497 |
| note | asia2 signing offload to europe1/america1, asia2 GOMEMLIMIT 1500MiB |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	139	411	1021	1022
europe	7	17	47	135	374	373
america	6	21	66	193	475	474
```

```
{"id":"27254","status":"done","portalStatus":"finished","passed":true,"score":194180,"raw":194180,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:16:51.427Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:16:54.421Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:16:57.420Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:17:00.430Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T07:17:03.419Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T07:17:04.033Z level=INFO msg="user moved" user=mute5288 from=asia to=america
time=2026-10-31T07:17:06.421Z level=INFO msg="america 2 日目: 21 人"
time=2026-10-31T07:17:09.207Z level=INFO msg="user moved" user=benniku-5410 from=europe to=america
time=2026-10-31T07:17:09.431Z level=INFO msg="asia 3 日目: 139 人"
time=2026-10-31T07:17:10.169Z level=INFO msg="user moved" user=shiyara_5499 from=america to=europe
time=2026-10-31T07:17:10.714Z level=INFO msg="user moved" user=tsufumu_5528 from=asia to=america
time=2026-10-31T07:17:11.873Z level=INFO msg="user moved" user=sunoso5883 from=asia to=america
time=2026-10-31T07:17:12.442Z level=INFO msg="europe 3 日目: 47 人"
time=2026-10-31T07:17:12.523Z level=INFO msg="user moved" user=sunoso5883 from=america to=asia
time=2026-10-31T07:17:12.810Z level=INFO msg="user moved" user=shiro-6105 from=asia to=america
time=2026-10-31T07:17:13.340Z level=INFO msg="user moved" user=shiyara_5499 from=europe to=asia
time=2026-10-31T07:17:13.482Z level=INFO msg="user moved" user=shiro-6105 from=america to=asia
time=2026-10-31T07:17:15.095Z level=INFO msg="user moved" user=tsufumu_5528 from=america to=europe
time=2026-10-31T07:17:15.420Z level=INFO msg="america 3 日目: 66 人"
time=2026-10-31T07:17:17.609Z level=INFO msg="user moved" user=danhemu.6317 from=europe to=asia
time=2026-10-31T07:17:18.376Z level=INFO msg="user moved" user=evator.6515 from=america to=asia
time=2026-10-31T07:17:18.376Z level=INFO msg="user moved" user=benniku-5410 from=america to=asia
time=2026-10-31T07:17:18.432Z level=INFO msg="asia 4 日目: 411 人"
time=2026-10-31T07:17:19.521Z level=INFO msg="user moved" user=shiyara_5499 from=asia to=europe
time=2026-10-31T07:17:20.662Z level=INFO msg="user moved" user=shiro-6105 from=asia to=america
time=2026-10-31T07:17:21.117Z level=INFO msg="user moved" user=evator.6515 from=asia to=europe
time=2026-10-31T07:17:21.480Z level=INFO msg="europe 4 日目: 135 人"
time=2026-10-31T07:17:21.688Z level=INFO msg="user moved" user=sunoso5883 from=asia to=america
time=2026-10-31T07:17:24.003Z level=INFO msg="user moved" user=neha_6941 from=asia to=europe
time=2026-10-31T07:17:24.367Z level=INFO msg="user moved" user=mute5288 from=america to=asia
time=2026-10-31T07:17:24.501Z level=INFO msg="america 4 日目: 193 人"
time=2026-10-31T07:17:24.571Z level=INFO msg="user moved" user=hiben.7178 from=asia to=america
time=2026-10-31T07:17:25.126Z level=INFO msg="user moved" user=rinoke.8269 from=asia to=europe
time=2026-10-31T07:17:25.211Z level=INFO msg="user moved" user=benniku-5410 from=asia to=america
time=2026-10-31T07:17:26.187Z level=INFO msg="user moved" user=danhemu.6317 from=asia to=america
time=2026-10-31T07:17:26.309Z level=INFO msg="user moved" user=elku-7596 from=asia to=america
time=2026-10-31T07:17:27.073Z level=INFO msg="user moved" user=himia_7419 from=asia to=america
time=2026-10-31T07:17:27.436Z level=INFO msg="asia 5 日目: 1021 人"
time=2026-10-31T07:17:31.143Z level=INFO msg="europe 5 日目: 374 人"
time=2026-10-31T07:17:32.993Z level=INFO msg="user moved" user=rosare9109 from=america to=asia
time=2026-10-31T07:17:33.460Z level=INFO msg="user moved" user=liatachi9709 from=america to=asia
time=2026-10-31T07:17:33.626Z level=INFO msg="america 5 日目: 475 人"
time=2026-10-31T07:17:46.473Z level=INFO msg="america 最終日（グランプリ）: 474 人"
time=2026-10-31T07:17:47.770Z level=INFO msg="asia 最終日（グランプリ）: 1022 人"
time=2026-10-31T07:17:47.792Z level=INFO msg="europe 最終日（グランプリ）: 373 人"
time=2026-10-31T07:17:51.542Z level=INFO msg="グランプリ: アジア 145 人, ヨーロッパ 46 人, アメリカ 194 人"
{
  "pass": true,
  "score": 194180,
  "messages": [
    "noaruel15444: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tsufumu_5528: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.33.159:52380-\u003e52.69.202.109:8080: write tcp 10.27.33.159:52380-\u003e52.69.202.109:8080: use of closed network connection"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
51.3%	1678	1	0	GET	/api/notifications	20779.674	12.384	20.292	20.314	172
13.1%	2477	449	0	POST	/api/posts	5314.987	2.146	13.129	13.244	539
5.3%	13574	0	0	GET	^/api/feeds/[^/]+	2135.656	0.157	1.013	1.993	4837
4.4%	17148	290	0	GET	^/api/images/[^/]+	1769.369	0.103	0.987	2.777	129942
4.0%	3741	0	0	GET	/api/memories	1615.285	0.432	3.640	7.059	13358
3.2%	5628	56	0	POST	/api/friend-requests	1309.220	0.233	1.355	2.495	114
2.3%	749	126	0	PUT	/api/me/profile-image	937.215	1.251	8.822	12.323	3
2.0%	6061	576	0	POST	/api/auth/login	818.099	0.135	1.250	1.792	223
1.9%	534	21	0	POST	^/api/posts/[^/]+/realmojis$	751.194	1.407	13.101	13.167	462
1.7%	3647	0	0	GET	/api/posting-window/current	691.238	0.190	1.483	1.932	237
1.7%	3279	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	680.163	0.207	1.353	2.419	116
1.4%	2666	0	0	GET	/api/users	584.895	0.219	1.183	1.990	255
1.1%	387	0	0	GET	/api/me/export	456.895	1.181	8.799	9.719	1972796
1.0%	1705	17	0	GET	^/api/posts/me/current$	397.085	0.233	1.074	1.897	652
0.7%	2313	54	0	GET	^/api/posts/[^/]+$	289.651	0.125	0.713	0.933	615
0.7%	632	3	0	GET	/api/me	272.788	0.432	1.857	2.012	302
0.5%	738	0	0	POST	^/api/posts/[^/]+/comments$	217.044	0.294	2.032	3.252	344
0.5%	1564	0	0	GET	^/api/relationships/[^/]+	204.027	0.130	0.789	1.564	1145
0.5%	1045	1	0	POST	/api/auth/register	185.741	0.178	1.019	1.584	227
0.4%	874	0	0	GET	^/api/memories/[^/]+	150.937	0.173	1.241	2.366	621
0.4%	1237	0	0	GET	^/api/users/[^/]+$	143.944	0.116	0.711	1.019	210
0.3%	925	0	0	POST	/api/auth/logout	140.845	0.152	1.339	2.284	0
0.3%	699	0	0	PUT	^/api/follows/[^/]+$	140.834	0.201	1.371	2.097	0
0.3%	494	0	0	PATCH	/api/me	122.909	0.249	1.591	1.772	282
0.3%	894	3	0	GET	^/api/posts/[^/]+/realmojis$	108.472	0.121	0.616	0.995	1182
0.2%	834	5	0	GET	^/api/posts/[^/]+/comments$	93.801	0.112	0.614	0.914	2554
0.1%	202	4	0	PATCH	^/api/posts/[^/]+$	59.997	0.297	1.703	1.711	643
0.1%	73	0	0	DELETE	^/api/posts/[^/]+$	47.171	0.646	1.811	1.811	0
0.1%	94	0	0	DELETE	^/api/realmojis/[^/]+$	36.140	0.384	1.979	1.979	0
0.1%	152	0	0	DELETE	^/api/comments/[^/]+$	33.087	0.218	1.592	1.643	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
52.0%	602	2	0	GET	/api/notifications	5556.212	9.230	16.023	16.024	168
8.4%	198	0	0	PUT	/api/me/profile-image	895.100	4.521	13.309	13.348	0
6.4%	836	169	0	POST	/api/posts	682.138	0.816	3.331	3.996	511
6.0%	1992	15	0	POST	/api/friend-requests	644.839	0.324	1.628	2.338	114
4.3%	1640	0	0	POST	/api/auth/login	463.441	0.283	1.240	2.691	249
4.2%	5256	0	0	GET	^/api/feeds/[^/]+	446.879	0.085	0.284	1.104	3780
4.1%	5356	25	0	GET	^/api/images/[^/]+	437.706	0.082	0.680	1.438	143874
2.1%	221	11	0	POST	^/api/posts/[^/]+/realmojis$	222.715	1.008	13.250	13.318	449
1.6%	369	0	0	POST	/api/auth/register	167.674	0.454	0.953	1.097	228
1.4%	495	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	152.050	0.307	1.601	2.114	115
1.1%	1150	0	0	GET	/api/memories	121.988	0.106	0.976	1.338	17058
1.1%	1284	0	0	GET	/api/posting-window/current	121.540	0.095	0.377	1.152	238
0.8%	239	0	0	POST	^/api/posts/[^/]+/comments$	89.938	0.376	2.094	2.645	341
0.8%	328	0	0	POST	/api/auth/logout	86.413	0.263	1.369	1.545	0
0.8%	925	0	0	GET	/api/users	82.985	0.090	0.261	0.421	255
0.6%	838	8	0	GET	^/api/posts/[^/]+$	68.336	0.082	0.242	0.628	616
0.6%	233	0	0	PUT	^/api/follows/[^/]+$	64.319	0.276	1.103	1.445	0
0.6%	190	0	0	PATCH	/api/me	60.779	0.320	1.093	1.676	268
0.5%	567	1	0	GET	^/api/posts/me/current$	54.805	0.097	0.297	0.370	640
0.5%	630	0	0	GET	^/api/relationships/[^/]+	51.088	0.081	0.227	0.297	522
0.4%	467	0	0	GET	^/api/users/[^/]+$	39.761	0.085	0.295	0.656	200
0.3%	64	0	0	PATCH	^/api/posts/[^/]+$	30.299	0.473	1.674	1.674	644
0.2%	51	0	0	DELETE	^/api/realmojis/[^/]+$	20.520	0.402	1.312	1.312	0
0.2%	28	0	0	DELETE	^/api/posts/[^/]+$	19.711	0.704	1.625	1.625	0
0.2%	219	0	0	GET	^/api/memories/[^/]+	19.090	0.087	0.290	0.716	858
0.2%	248	0	0	GET	^/api/posts/[^/]+/realmojis$	18.630	0.075	0.254	0.296	1363
0.2%	203	2	0	GET	/api/me	18.215	0.090	0.353	1.104	303
0.2%	231	1	0	GET	^/api/posts/[^/]+/comments$	18.002	0.078	0.297	0.298	2862
0.1%	91	0	0	GET	/api/me/export	13.427	0.148	2.589	2.589	2190192
0.1%	58	0	0	DELETE	^/api/comments/[^/]+$	12.287	0.212	1.083	1.083	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
46.1%	757	1	0	GET	/api/notifications	5893.028	7.785	12.969	13.353	172
7.7%	2443	18	0	POST	/api/friend-requests	984.425	0.403	1.691	2.240	114
6.6%	1126	142	0	POST	/api/posts	846.772	0.752	3.401	3.988	573
6.3%	279	0	0	PUT	/api/me/profile-image	811.724	2.909	13.351	13.359	0
6.3%	6239	0	0	GET	^/api/feeds/[^/]+	807.720	0.129	0.780	1.017	3733
5.7%	6079	33	0	GET	^/api/images/[^/]+	724.476	0.119	0.947	2.119	172122
4.2%	1703	0	0	POST	/api/auth/login	533.979	0.314	1.360	1.945	250
2.4%	1614	0	0	GET	/api/posting-window/current	303.061	0.188	0.844	0.929	240
1.8%	662	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	231.413	0.350	1.762	2.354	116
1.5%	1457	0	0	GET	/api/memories	198.178	0.136	0.858	1.298	16060
1.4%	471	0	0	POST	/api/auth/register	175.283	0.372	0.922	1.090	230
1.1%	1155	0	0	GET	/api/users	146.022	0.126	0.682	0.844	254
1.1%	126	13	0	POST	^/api/posts/[^/]+/realmojis$	136.594	1.084	7.545	7.741	426
1.0%	303	0	0	PUT	^/api/follows/[^/]+$	129.121	0.426	1.809	2.016	0
0.9%	885	4	0	GET	^/api/posts/me/current$	116.105	0.131	0.517	0.948	653
0.9%	961	15	0	GET	^/api/posts/[^/]+$	112.835	0.117	0.516	0.923	611
0.7%	193	3	0	POST	^/api/posts/[^/]+/comments$	93.511	0.485	2.580	2.688	333
0.7%	201	0	0	PATCH	/api/me	91.094	0.453	1.802	1.820	280
0.7%	735	0	0	GET	^/api/relationships/[^/]+	88.290	0.120	0.499	0.754	680
0.6%	269	0	0	POST	/api/auth/logout	82.369	0.306	1.290	2.239	0
0.4%	458	0	0	GET	^/api/users/[^/]+$	54.300	0.119	0.606	0.809	204
0.3%	97	2	0	PATCH	^/api/posts/[^/]+$	43.324	0.447	1.728	1.728	640
0.3%	283	1	0	GET	/api/me	32.183	0.114	0.897	0.910	305
0.2%	225	0	0	GET	^/api/memories/[^/]+	26.232	0.117	0.791	1.127	1060
0.2%	34	0	0	DELETE	^/api/posts/[^/]+$	25.101	0.738	2.051	2.051	0
0.2%	218	0	0	GET	^/api/posts/[^/]+/comments$	23.814	0.109	0.347	0.746	2364
0.2%	204	0	0	GET	^/api/posts/[^/]+/realmojis$	21.588	0.106	0.360	0.797	1010
0.1%	30	0	0	DELETE	^/api/realmojis/[^/]+$	16.937	0.565	1.999	1.999	0
0.1%	95	0	0	GET	/api/me/export	16.702	0.176	1.250	1.250	2964505
0.1%	43	0	0	DELETE	^/api/comments/[^/]+$	12.943	0.301	1.204	1.204	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16180 クエリ / 48 種類 / 合計 193.210s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	54.8%	391	105.846	270.71	876.09	1255.3	128.3	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	13.6%	4	26.357	6589.26	9530.79	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	10.5%	11937	20.358	1.71	9.63	9.7	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.2%	4	19.686	4921.48	7199.62	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	4.7%	4	9.048	2261.91	3440.46	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	4.0%	6	7.664	1277.36	3562.81	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.6%	1	1.200	1199.99	1199.99	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.3%	1	0.516	515.81	515.81	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
9	0.2%	6	0.465	77.49	156.66	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
10	0.2%	6	0.430	71.63	156.30	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
11	0.2%	1	0.393	392.66	392.66	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.260	260.20	260.20	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.166	166.42	166.42	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.158	158.48	158.48	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	1	0.113	113.26	113.26	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows

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
Time: 2026-10-31 16:16:25 JST
Duration: 60.01s, Total samples = 33.05s (55.08%)
Showing nodes accounting for 25.48s, 77.10% of 33.05s total
Dropped 893 nodes (cum <= 0.17s)
Showing top 40 nodes out of 235
      flat  flat%   sum%        cum   cum%
     5.92s 17.91% 17.91%      5.92s 17.91%  crypto/internal/fips140/sha256.blockAVX2
     3.07s  9.29% 27.20%      3.07s  9.29%  internal/runtime/syscall/linux.Syscall6
     1.76s  5.33% 32.53%      1.76s  5.33%  runtime.memmove
     1.72s  5.20% 37.73%      6.39s 19.33%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.49s  4.51% 42.24%      1.49s  4.51%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.03s  3.12% 45.36%      1.03s  3.12%  aeshashbody
     0.95s  2.87% 48.23%      3.50s 10.59%  runtime.scanObject
     0.91s  2.75% 50.98%      1.25s  3.78%  runtime.findObject
     0.89s  2.69% 53.68%      0.89s  2.69%  runtime.memclrNoHeapPointers
     0.79s  2.39% 56.07%      1.30s  3.93%  runtime.tryDeferToSpanScan
     0.51s  1.54% 57.61%      1.95s  5.90%  runtime.mapaccess1_faststr
     0.45s  1.36% 58.97%      0.69s  2.09%  example.com/isureal-company/isureal.postLess
     0.40s  1.21% 60.18%      0.65s  1.97%  database/sql.convertAssignRows
     0.36s  1.09% 61.27%      0.36s  1.09%  runtime.nextFreeFast (inline)
     0.36s  1.09% 62.36%      0.36s  1.09%  runtime.spanClass.sizeclass (inline)
     0.35s  1.06% 63.42%      6.68s 20.21%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.34s  1.03% 64.45%      0.34s  1.03%  memeqbody
     0.32s  0.97% 65.42%      0.32s  0.97%  sync/atomic.(*Int32).Add (inline)
     0.30s  0.91% 66.32%      6.50s 19.67%  crypto/internal/fips140/sha256.(*Digest).Write
     0.29s  0.88% 67.20%      0.29s  0.88%  runtime.futex
     0.24s  0.73% 67.93%      1.22s  3.69%  runtime.mallocgcSmallScanNoHeader
     0.23s   0.7% 68.62%      0.27s  0.82%  runtime.(*mspan).writeHeapBitsSmall
     0.22s  0.67% 69.29%      2.13s  6.44%  runtime.mallocgc
     0.21s  0.64% 69.92%      0.35s  1.06%  internal/runtime/maps.(*Iter).Next
     0.21s  0.64% 70.56%      0.21s  0.64%  runtime.(*mspan).divideByElemSize (inline)
     0.20s  0.61% 71.16%      7.39s 22.36%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.20s  0.61% 71.77%      6.97s 21.09%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.20s  0.61% 72.38%      0.39s  1.18%  runtime.greyobject
     0.19s  0.57% 72.95%      0.24s  0.73%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.15s  0.45% 73.40%      0.33s     1%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.15s  0.45% 73.86%      4.98s 15.07%  example.com/isureal-company/isureal.loadBig
     0.15s  0.45% 74.31%      2.47s  7.47%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.13s  0.39% 74.70%      7.94s 24.02%  example.com/isureal-company/isureal.computeImageSignature
     0.12s  0.36% 75.07%      1.43s  4.33%  encoding/gob.(*Encoder).encodeStruct
     0.12s  0.36% 75.43%      0.17s  0.51%  runtime.(*mspan).markBitsForIndex (inline)
     0.12s  0.36% 75.79%      0.48s  1.45%  runtime.mapaccess2_faststr
     0.12s  0.36% 76.16%      1.47s  4.45%  runtime.mapassign_faststr
     0.11s  0.33% 76.49%      0.76s  2.30%  database/sql.(*Rows).scanLocked
     0.11s  0.33% 76.82%      0.17s  0.51%  example.com/isureal-company/isureal.(*Store).snapshotLocked
     0.09s  0.27% 77.10%      0.25s  0.76%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 45.7% / 100% | 26% | 5.6% | 7 | 118 MB | run queue が積もっている |
| isu2 (2 cpu) | 23.3% / 100% | 2% | 0.4% | 5 | 630 MB | run queue が積もっている |
| isu3 (2 cpu) | 25.9% / 100% | 7% | 0.8% | 4 | 180 MB |  |
| isu4 (2 cpu) | 51.9% / 100% | 36% | 9.6% | 13 | 102 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 53% / nginx 11% / find 6% / kswapd0 3% |
| isu2 | isureal 44% / nginx 2% |
| isu3 | isureal 48% / nginx 2% |
| isu4 | isureal 61% / postgres 18% |

## Diff from previous run

スコア: 201497 → 194180 (-7317)

`40e65ef..3f85847`

```
 conf/isu1/home/isucon/env.sh                    | 2 +-
 conf/isu4/etc/systemd/system/isureal-go.service | 5 +++--
 conf/isu4/home/isucon/env.sh                    | 1 +
 3 files changed, 5 insertions(+), 3 deletions(-)
```
