# run 20261031-170001-a1cea2f

## Summary

| | |
|---|---|
| score | 223475 |
| status | pass |
| ref | sign-cache-to-asia1 |
| sha | a1cea2f11404 |
| 所要時間 | 142s |
| 前回比 | +2935 (1.3%) vs 20261031-163606-15a1e6c |
| ベスト比 | -267 (-0.1%) vs 20261031-163105-6f2cd6f = 223742 |
| note | asia2 が計算した署名を asia1 のキャッシュへ送る（main bdf05cd の上） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	140	378	706	706
europe	7	21	60	173	484	483
america	6	18	56	167	462	461
```

```
{"id":"31574","status":"done","portalStatus":"finished","passed":true,"score":223475,"raw":223475,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:00:30.694Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:00:33.691Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:00:36.684Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:00:39.698Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T08:00:42.414Z level=INFO msg="user moved" user=yumihi.5267 from=asia to=europe
time=2026-10-31T08:00:42.687Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T08:00:45.684Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T08:00:48.708Z level=INFO msg="asia 3 日目: 140 人"
time=2026-10-31T08:00:51.686Z level=INFO msg="europe 3 日目: 60 人"
time=2026-10-31T08:00:53.867Z level=INFO msg="user moved" user=yumihi.5267 from=europe to=asia
time=2026-10-31T08:00:54.134Z level=INFO msg="user moved" user=yumian.6282 from=europe to=america
time=2026-10-31T08:00:54.708Z level=INFO msg="user moved" user=yumian.6282 from=america to=europe
time=2026-10-31T08:00:54.713Z level=INFO msg="america 3 日目: 56 人"
time=2026-10-31T08:00:57.727Z level=INFO msg="asia 4 日目: 378 人"
time=2026-10-31T08:01:00.684Z level=INFO msg="europe 4 日目: 173 人"
time=2026-10-31T08:01:00.954Z level=INFO msg="user moved" user=yumihi.5267 from=asia to=europe
time=2026-10-31T08:01:03.305Z level=INFO msg="user moved" user=yumian.6282 from=europe to=america
time=2026-10-31T08:01:03.654Z level=INFO msg="user moved" user=yalia6757 from=asia to=america
time=2026-10-31T08:01:03.707Z level=INFO msg="america 4 日目: 167 人"
time=2026-10-31T08:01:04.198Z level=INFO msg="user moved" user=nurase8526 from=asia to=europe
time=2026-10-31T08:01:04.529Z level=INFO msg="user moved" user=liamax7638 from=asia to=america
time=2026-10-31T08:01:04.656Z level=INFO msg="user moved" user=maxrushi-8278 from=asia to=europe
time=2026-10-31T08:01:06.694Z level=INFO msg="user moved" user=rochi.9000 from=europe to=asia
time=2026-10-31T08:01:06.713Z level=INFO msg="asia 5 日目: 706 人"
time=2026-10-31T08:01:09.535Z level=INFO msg="user moved" user=keshi8929 from=europe to=asia
time=2026-10-31T08:01:09.754Z level=INFO msg="europe 5 日目: 484 人"
time=2026-10-31T08:01:12.789Z level=INFO msg="america 5 日目: 462 人"
time=2026-10-31T08:01:21.711Z level=INFO msg="america 最終日（グランプリ）: 461 人"
time=2026-10-31T08:01:21.742Z level=INFO msg="asia 最終日（グランプリ）: 706 人"
time=2026-10-31T08:01:21.836Z level=INFO msg="europe 最終日（グランプリ）: 483 人"
time=2026-10-31T08:01:30.815Z level=INFO msg="グランプリ: アジア 453 人, ヨーロッパ 442 人, アメリカ 363 人"
{
  "pass": true,
  "score": 223475,
  "messages": [
    "leoha-5017: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "fura_8944: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "wachi-8774: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "natsu-5032: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "noaisul-11081: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.32.185:55322-\u003e52.69.202.109:8080: write tcp 10.27.32.185:55322-\u003e52.69.202.109:8080: use of closed network connection",
    "temaxshi_13775: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "miaozhen_gongfang_970: ログアウトできません: POST /api/auth/logout: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/logout\": EOF",
    "evafumi6564: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
79.6%	2016	1	0	GET	/api/notifications	54639.325	27.103	60.007	60.012	113
4.3%	11494	0	0	GET	^/api/feeds/[^/]+	2929.994	0.255	1.515	2.635	5510
3.3%	3013	0	0	GET	/api/memories	2282.318	0.757	5.758	9.050	10929
2.3%	2748	0	0	GET	/api/posting-window/current	1547.361	0.563	2.073	2.658	236
2.2%	1991	46	0	POST	/api/posts	1524.975	0.766	6.347	7.841	638
1.1%	2052	0	0	GET	/api/users	761.385	0.371	1.800	2.644	236
1.0%	221	0	0	GET	/api/me/export	697.464	3.156	17.343	20.443	1821274
1.0%	1633	1	0	GET	^/api/posts/me/current$	668.712	0.409	1.263	1.753	653
0.6%	12099	235	0	GET	^/api/images/[^/]+	419.428	0.035	0.633	1.960	174075
0.6%	544	128	0	PUT	/api/me/profile-image	402.689	0.740	11.638	15.675	4
0.5%	4402	60	0	POST	/api/friend-requests	318.356	0.072	0.745	1.388	114
0.5%	430	3	0	GET	/api/me	312.403	0.727	1.931	2.615	300
0.4%	1465	0	0	GET	^/api/relationships/[^/]+	298.476	0.204	1.265	2.144	1424
0.4%	844	14	0	GET	^/api/posts/[^/]+/realmojis$	297.988	0.353	2.330	2.884	2175
0.4%	720	2	0	POST	^/api/posts/[^/]+/realmojis$	267.387	0.371	2.592	2.809	482
0.4%	2001	76	0	GET	^/api/posts/[^/]+$	266.895	0.133	0.752	1.178	603
0.3%	3527	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	234.844	0.067	0.734	1.273	116
0.3%	524	0	0	GET	^/api/memories/[^/]+	179.440	0.342	1.849	2.872	558
0.2%	867	8	0	GET	^/api/posts/[^/]+/comments$	128.509	0.148	0.850	1.588	3078
0.2%	995	0	0	GET	^/api/users/[^/]+$	122.111	0.123	0.800	1.063	197
0.1%	784	0	0	POST	^/api/posts/[^/]+/comments$	98.978	0.126	1.083	1.488	361
0.1%	3973	515	0	POST	/api/auth/login	81.455	0.021	0.549	1.254	215
0.0%	78	0	0	DELETE	^/api/posts/[^/]+$	34.125	0.437	1.256	1.256	0
0.0%	100	0	0	DELETE	^/api/realmojis/[^/]+$	26.561	0.266	0.946	0.996	0
0.0%	174	0	0	PATCH	^/api/posts/[^/]+$	17.961	0.103	0.762	1.275	644
0.0%	617	0	0	POST	/api/auth/logout	17.840	0.029	0.616	0.771	0
0.0%	1	0	0	POST	/api/initialize	17.409	17.409	17.409	17.409	90
0.0%	433	0	0	PUT	^/api/follows/[^/]+$	9.701	0.022	0.511	0.718	0
0.0%	144	0	0	DELETE	^/api/comments/[^/]+$	8.463	0.059	0.797	1.130	0
0.0%	346	0	0	PATCH	/api/me	8.118	0.023	0.695	0.740	276
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
85.5%	1234	0	0	GET	/api/notifications	33507.398	27.153	60.002	60.101	108
2.1%	2826	29	0	POST	/api/friend-requests	803.955	0.284	1.016	1.533	114
2.0%	7364	0	0	GET	^/api/feeds/[^/]+	803.160	0.109	1.126	1.784	3897
1.8%	6851	42	0	GET	^/api/images/[^/]+	724.730	0.106	0.899	2.727	179772
1.7%	1222	14	0	POST	/api/posts	685.999	0.561	6.401	8.021	645
0.9%	1955	0	0	POST	/api/auth/login	366.656	0.188	0.824	1.379	248
0.7%	1780	0	0	GET	/api/memories	287.516	0.162	1.384	3.125	12794
0.6%	541	9	0	POST	^/api/posts/[^/]+/realmojis$	218.574	0.404	2.342	3.087	469
0.5%	274	0	0	PUT	/api/me/profile-image	210.667	0.769	2.313	2.561	0
0.5%	1092	0	0	GET	^/api/posts/me/current$	188.973	0.173	1.341	1.567	653
0.5%	1741	0	0	GET	/api/posting-window/current	177.038	0.102	0.470	1.719	239
0.4%	1356	0	0	GET	/api/users	139.958	0.103	0.474	1.445	237
0.3%	482	0	0	POST	/api/auth/register	136.437	0.283	0.663	0.903	229
0.3%	616	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	131.839	0.214	0.954	1.200	116
0.3%	906	0	0	GET	^/api/relationships/[^/]+	106.896	0.118	1.322	1.667	562
0.3%	1165	20	0	GET	^/api/posts/[^/]+$	100.638	0.086	0.421	1.258	623
0.2%	312	0	0	PUT	^/api/follows/[^/]+$	94.687	0.303	0.952	1.067	0
0.2%	385	3	0	POST	^/api/posts/[^/]+/comments$	76.984	0.200	1.319	1.780	351
0.2%	237	0	0	PATCH	/api/me	75.948	0.320	0.990	1.425	286
0.1%	342	0	0	POST	/api/auth/logout	57.686	0.169	0.930	1.072	0
0.1%	595	0	0	GET	^/api/users/[^/]+$	48.028	0.081	0.406	1.090	206
0.1%	571	0	0	GET	^/api/posts/[^/]+/realmojis$	45.517	0.080	0.361	0.499	3444
0.1%	119	1	0	PATCH	^/api/posts/[^/]+$	39.085	0.328	1.017	1.027	642
0.1%	283	2	0	GET	/api/me	32.396	0.114	0.420	0.436	302
0.1%	400	1	0	GET	^/api/posts/[^/]+/comments$	31.627	0.079	0.367	0.496	4985
0.1%	297	0	0	GET	^/api/memories/[^/]+	28.684	0.097	0.680	1.496	877
0.1%	80	0	0	DELETE	^/api/realmojis/[^/]+$	26.943	0.337	1.389	1.389	0
0.1%	46	0	0	DELETE	^/api/posts/[^/]+$	26.300	0.572	2.120	2.120	0
0.0%	114	0	0	GET	/api/me/export	15.757	0.138	0.827	1.853	2140055
0.0%	62	0	0	DELETE	^/api/comments/[^/]+$	8.690	0.140	0.786	0.786	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
85.7%	1170	0	0	GET	/api/notifications	31049.894	26.538	60.001	60.103	108
2.4%	1168	13	0	POST	/api/posts	851.755	0.729	2.355	8.003	647
2.1%	4972	0	0	GET	^/api/feeds/[^/]+	762.805	0.153	1.567	2.327	3319
2.0%	5868	32	0	GET	^/api/images/[^/]+	716.248	0.122	1.429	2.808	159149
1.4%	2601	34	0	POST	/api/friend-requests	512.711	0.197	1.484	1.592	114
1.2%	1745	0	0	POST	/api/auth/login	432.711	0.248	1.354	2.388	250
0.7%	1550	0	0	GET	/api/memories	244.572	0.158	1.711	2.540	14523
0.6%	1670	0	0	GET	/api/posting-window/current	212.506	0.127	1.512	1.819	240
0.6%	458	0	0	POST	/api/auth/register	204.712	0.447	1.055	1.149	230
0.5%	974	6	0	GET	^/api/posts/me/current$	164.608	0.169	1.485	1.771	650
0.4%	1216	0	0	GET	/api/users	149.211	0.123	1.407	1.636	238
0.3%	259	0	0	PUT	/api/me/profile-image	114.702	0.443	2.244	2.435	0
0.3%	253	11	0	POST	^/api/posts/[^/]+/realmojis$	105.784	0.418	1.986	3.323	454
0.3%	472	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	96.149	0.204	1.431	1.569	116
0.2%	631	0	0	GET	^/api/relationships/[^/]+	86.415	0.137	1.474	1.898	584
0.2%	677	36	0	GET	^/api/posts/[^/]+$	79.955	0.118	1.391	1.659	598
0.2%	215	0	0	PATCH	/api/me	55.568	0.258	1.084	1.126	281
0.1%	285	0	0	PUT	^/api/follows/[^/]+$	50.477	0.177	0.866	1.442	0
0.1%	271	0	0	POST	/api/auth/logout	47.054	0.174	0.958	1.465	0
0.1%	224	6	0	POST	^/api/posts/[^/]+/comments$	41.454	0.185	1.138	1.465	303
0.1%	108	0	0	GET	/api/me/export	33.784	0.313	5.261	5.400	3139508
0.1%	95	4	0	PATCH	^/api/posts/[^/]+$	32.534	0.342	1.679	1.679	617
0.1%	311	0	0	GET	^/api/users/[^/]+$	31.590	0.102	0.366	1.505	203
0.1%	268	1	0	GET	^/api/posts/[^/]+/realmojis$	31.410	0.117	1.486	1.663	3244
0.1%	235	2	0	GET	^/api/posts/[^/]+/comments$	28.050	0.119	1.343	1.521	4182
0.1%	234	0	0	GET	^/api/memories/[^/]+	24.786	0.106	1.148	2.137	1114
0.1%	41	0	0	DELETE	^/api/posts/[^/]+$	24.150	0.589	1.968	1.968	0
0.1%	267	1	0	GET	/api/me	20.372	0.076	0.283	0.340	303
0.0%	33	0	0	DELETE	^/api/realmojis/[^/]+$	12.531	0.380	1.790	1.790	0
0.0%	35	0	0	DELETE	^/api/comments/[^/]+$	4.933	0.141	0.353	0.353	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 11421 クエリ / 48 種類 / 合計 126.429s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	44.4%	229	56.086	244.92	1996.37	1192.1	121.2	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	20.1%	4	25.418	6354.46	10095.18	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.8%	4	18.689	4672.19	7201.36	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	9.3%	7327	11.754	1.60	9.90	10.2	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.8%	4	8.540	2134.92	3441.73	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	1.7%	7	2.121	302.98	455.28	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.6%	1	0.819	818.84	818.84	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
8	0.5%	7	0.618	88.27	161.86	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.5%	7	0.575	82.10	159.87	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.3%	1	0.384	384.12	384.12	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.2%	1	0.271	270.59	270.59	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.2%	1	0.210	209.94	209.94	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.178	177.64	177.64	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.162	162.13	162.13	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
15	0.1%	1	0.096	95.77	95.77	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: 5641a5b2d0ad69fa39783d8f97e6028e0bc2f5a6
Type: cpu
Time: 2026-10-31 17:00:07 JST
Duration: 60s, Total samples = 27.60s (46.00%)
Showing nodes accounting for 21.26s, 77.03% of 27.60s total
Dropped 844 nodes (cum <= 0.14s)
Showing top 40 nodes out of 252
      flat  flat%   sum%        cum   cum%
     3.69s 13.37% 13.37%      3.69s 13.37%  crypto/internal/fips140/sha256.blockAVX2
     2.96s 10.72% 24.09%      2.96s 10.72%  internal/runtime/syscall/linux.Syscall6
     1.64s  5.94% 30.04%      5.86s 21.23%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.57s  5.69% 35.72%      1.57s  5.69%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.30s  4.71% 40.43%      1.30s  4.71%  runtime.memmove
     1.12s  4.06% 44.49%      1.12s  4.06%  aeshashbody
     0.86s  3.12% 47.61%      3.19s 11.56%  runtime.scanObject
     0.81s  2.93% 50.54%      0.81s  2.93%  runtime.memclrNoHeapPointers
     0.74s  2.68% 53.22%      0.91s  3.30%  runtime.findObject
     0.65s  2.36% 55.58%      1.07s  3.88%  runtime.tryDeferToSpanScan
     0.58s  2.10% 57.68%      0.87s  3.15%  example.com/isureal-company/isureal/jpegcheck.(*decoder).skipAC
     0.46s  1.67% 59.35%      2.08s  7.54%  runtime.mapaccess1_faststr
     0.33s  1.20% 60.54%      0.47s  1.70%  example.com/isureal-company/isureal.postLess
     0.33s  1.20% 61.74%      0.33s  1.20%  memeqbody
     0.30s  1.09% 62.83%      0.30s  1.09%  example.com/isureal-company/isureal/jpegcheck.(*decoder).prefetch
     0.27s  0.98% 63.80%      4.28s 15.51%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.27s  0.98% 64.78%      1.29s  4.67%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.26s  0.94% 65.72%      0.26s  0.94%  runtime.spanClass.sizeclass (inline)
     0.25s  0.91% 66.63%      0.25s  0.91%  sync/atomic.(*Int32).Add
     0.24s  0.87% 67.50%      0.24s  0.87%  runtime.nextFreeFast (inline)
     0.20s  0.72% 68.22%      0.31s  1.12%  database/sql.convertAssignRows
     0.18s  0.65% 68.88%      0.20s  0.72%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.17s  0.62% 69.49%      0.73s  2.64%  runtime.mallocgcSmallScanNoHeader
     0.16s  0.58% 70.07%      0.17s  0.62%  runtime.(*mspan).writeHeapBitsSmall
     0.15s  0.54% 70.62%      4.49s 16.27%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.15s  0.54% 71.16%      4.09s 14.82%  crypto/internal/fips140/sha256.(*Digest).Write
     0.15s  0.54% 71.70%      0.96s  3.48%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.14s  0.51% 72.21%      0.14s  0.51%  runtime.futex
     0.14s  0.51% 72.72%      0.14s  0.51%  runtime.typePointers.nextFast (inline)
     0.12s  0.43% 73.15%      4.70s 17.03%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.12s  0.43% 73.59%      1.66s  6.01%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.12s  0.43% 74.02%      0.17s  0.62%  internal/runtime/maps.(*Iter).Next
     0.12s  0.43% 74.46%      1.49s  5.40%  runtime.mapassign_faststr
     0.11s   0.4% 74.86%      0.25s  0.91%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.11s   0.4% 75.25%      0.14s  0.51%  example.com/isureal-company/isureal.rmLess
     0.11s   0.4% 75.65%      0.30s  1.09%  runtime.greyobject
     0.11s   0.4% 76.05%      1.77s  6.41%  runtime.mallocgc
     0.10s  0.36% 76.41%      0.15s  0.54%  runtime.(*mspan).markBitsForIndex (inline)
     0.09s  0.33% 76.74%      0.37s  1.34%  runtime.netpoll
     0.08s  0.29% 77.03%      0.41s  1.49%  runtime.mallocgcSmallNoscan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 30.2% / 100% | 8% | 0.5% | 7 | 146 MB | run queue が積もっている |
| isu2 (2 cpu) | 21.9% / 100% | 5% | 0.7% | 5 | 104 MB | run queue が積もっている |
| isu3 (2 cpu) | 22.9% / 100% | 7% | 0.9% | 4 | 188 MB |  |
| isu4 (2 cpu) | 56.9% / 100% | 47% | 1.8% | 5 | 100 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 44% / nginx 9% |
| isu2 | isureal 39% / nginx 3% |
| isu3 | isureal 42% / nginx 2% |
| isu4 | isureal 95% / postgres 13% |

## Diff from previous run

スコア: 220540 → 223475 (+2935)

`15a1e6c..a1cea2f`

```
 conf/isu4/home/isucon/env.sh  |   1 +
 webapp/go/cmd/isureal/main.go |   3 ++
 webapp/go/image_signature.go  |   1 +
 webapp/go/offload.go          |   4 ++
 webapp/go/sign_share.go       | 105 ++++++++++++++++++++++++++++++++++++++++++
 5 files changed, 114 insertions(+)
```
