# run 20261031-163105-6f2cd6f

## Summary

| | |
|---|---|
| score | 223742 |
| status | pass |
| ref | jpegcheck-fast-ac |
| sha | 6f2cd6fea333 |
| 所要時間 | 268s |
| 前回比 | +22245 (11.0%) vs 20261031-160409-40e65ef |
| ベスト比 | +22245 (11.0%) vs 20261031-160409-40e65ef = 201497 **ベスト更新** |
| note | jpegcheck の AC を 10 ビット表で読み飛ばす（main 73b6aff の上） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	149	418	658	659
europe	7	23	70	202	564	563
america	6	18	56	168	435	434
```

```
{"id":"28881","status":"done","portalStatus":"finished","passed":true,"score":223742,"raw":223742,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:33:30.359Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:33:33.350Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:33:36.345Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:33:39.389Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T07:33:40.786Z level=INFO msg="user moved" user=muyo-5183 from=asia to=europe
time=2026-10-31T07:33:41.542Z level=INFO msg="user moved" user=muyo-5183 from=europe to=asia
time=2026-10-31T07:33:42.346Z level=INFO msg="europe 2 日目: 23 人"
time=2026-10-31T07:33:45.346Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T07:33:48.370Z level=INFO msg="asia 3 日目: 149 人"
time=2026-10-31T07:33:50.131Z level=INFO msg="user moved" user=muyo-5183 from=asia to=europe
time=2026-10-31T07:33:51.345Z level=INFO msg="europe 3 日目: 70 人"
time=2026-10-31T07:33:52.908Z level=INFO msg="user moved" user=sokime_5895 from=asia to=europe
time=2026-10-31T07:33:53.323Z level=INFO msg="user moved" user=runuke.5627 from=asia to=america
time=2026-10-31T07:33:53.863Z level=INFO msg="user moved" user=iszoeku5547 from=america to=asia
time=2026-10-31T07:33:54.347Z level=INFO msg="america 3 日目: 56 人"
time=2026-10-31T07:33:55.094Z level=INFO msg="user moved" user=sukueva_6303 from=europe to=america
time=2026-10-31T07:33:55.762Z level=INFO msg="user moved" user=fudan6391 from=europe to=america
time=2026-10-31T07:33:57.388Z level=INFO msg="asia 4 日目: 418 人"
time=2026-10-31T07:34:00.386Z level=INFO msg="europe 4 日目: 202 人"
time=2026-10-31T07:34:03.342Z level=INFO msg="user moved" user=toor8719 from=europe to=america
time=2026-10-31T07:34:03.392Z level=INFO msg="america 4 日目: 168 人"
time=2026-10-31T07:34:04.038Z level=INFO msg="user moved" user=muyo-5183 from=europe to=asia
time=2026-10-31T07:34:05.365Z level=INFO msg="user moved" user=heta-7793 from=asia to=america
time=2026-10-31T07:34:05.592Z level=INFO msg="user moved" user=ornofu_8477 from=asia to=america
time=2026-10-31T07:34:06.480Z level=INFO msg="asia 5 日目: 658 人"
time=2026-10-31T07:34:06.834Z level=INFO msg="user moved" user=sokime_5895 from=europe to=asia
time=2026-10-31T07:34:09.414Z level=INFO msg="europe 5 日目: 564 人"
time=2026-10-31T07:34:10.163Z level=INFO msg="user moved" user=sukueva_6303 from=america to=asia
time=2026-10-31T07:34:12.416Z level=INFO msg="america 5 日目: 435 人"
time=2026-10-31T07:34:21.367Z level=INFO msg="america 最終日（グランプリ）: 434 人"
time=2026-10-31T07:34:21.386Z level=INFO msg="europe 最終日（グランプリ）: 563 人"
time=2026-10-31T07:34:21.424Z level=INFO msg="asia 最終日（グランプリ）: 659 人"
time=2026-10-31T07:34:30.397Z level=INFO msg="グランプリ: アジア 442 人, ヨーロッパ 249 人, アメリカ 413 人"
{
  "pass": true,
  "score": 223742,
  "messages": [
    "mihaki-5102: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "mashiri_12082: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sare_7004: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "useon_jinjeoeil_4956: ログアウトできません: POST /api/auth/logout: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/logout\": EOF",
    "miaro_10722: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sesuma_6494: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
79.3%	1977	0	0	GET	/api/notifications	51168.952	25.882	60.003	60.003	116
4.1%	10498	0	0	GET	^/api/feeds/[^/]+	2668.405	0.254	1.611	2.113	5405
3.4%	3002	0	0	GET	/api/memories	2173.172	0.724	6.047	11.506	11076
2.2%	1960	48	0	POST	/api/posts	1415.055	0.722	6.517	7.360	633
2.2%	2684	0	0	GET	/api/posting-window/current	1414.304	0.527	1.749	2.112	237
1.2%	1964	0	0	GET	/api/users	801.553	0.408	1.753	1.900	239
1.0%	1654	12	0	GET	^/api/posts/me/current$	657.205	0.397	1.444	1.734	644
1.0%	262	0	0	GET	/api/me/export	647.630	2.472	11.263	12.693	1720973
0.8%	487	128	0	PUT	/api/me/profile-image	487.043	1.000	11.609	14.656	4
0.7%	11902	242	0	GET	^/api/images/[^/]+	435.075	0.037	0.608	1.771	153107
0.6%	4097	69	0	POST	/api/friend-requests	374.524	0.091	0.993	1.696	114
0.4%	1329	0	0	GET	^/api/relationships/[^/]+	284.982	0.214	1.169	1.707	1664
0.4%	3576	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	277.981	0.078	0.991	1.907	116
0.4%	373	3	0	GET	/api/me	254.295	0.682	1.863	2.046	298
0.4%	836	7	0	GET	^/api/posts/[^/]+/realmojis$	235.839	0.282	1.543	2.159	1928
0.3%	1908	63	0	GET	^/api/posts/[^/]+$	224.790	0.118	0.896	1.281	617
0.3%	615	5	0	POST	^/api/posts/[^/]+/realmojis$	220.769	0.359	2.236	2.975	468
0.2%	535	0	0	GET	^/api/memories/[^/]+	159.347	0.298	1.837	2.521	703
0.2%	4054	583	0	POST	/api/auth/login	129.504	0.032	0.948	1.503	213
0.2%	935	0	0	GET	^/api/users/[^/]+$	113.023	0.121	0.980	1.071	211
0.2%	788	6	0	GET	^/api/posts/[^/]+/comments$	96.862	0.123	0.949	1.279	2601
0.1%	620	1	0	POST	^/api/posts/[^/]+/comments$	92.778	0.150	1.345	1.698	355
0.1%	78	0	0	DELETE	^/api/posts/[^/]+$	42.821	0.549	1.246	1.246	0
0.1%	627	0	0	POST	/api/auth/logout	33.433	0.053	0.950	1.108	0
0.0%	115	0	0	DELETE	^/api/realmojis/[^/]+$	32.137	0.279	1.642	1.823	0
0.0%	177	0	0	PATCH	^/api/posts/[^/]+$	27.225	0.154	1.021	1.538	645
0.0%	1	0	0	POST	/api/initialize	21.752	21.752	21.752	21.752	90
0.0%	414	0	0	PUT	^/api/follows/[^/]+$	15.300	0.037	0.840	1.008	0
0.0%	304	0	0	PATCH	/api/me	11.397	0.037	0.800	1.059	274
0.0%	138	0	0	DELETE	^/api/comments/[^/]+$	7.415	0.054	1.034	1.434	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
84.4%	1440	1	0	GET	/api/notifications	39014.636	27.093	60.116	60.443	105
2.5%	7946	0	0	GET	^/api/feeds/[^/]+	1138.252	0.143	1.620	2.503	3593
2.3%	1418	44	0	POST	/api/posts	1045.042	0.737	3.078	7.652	638
2.0%	3196	38	0	POST	/api/friend-requests	903.116	0.283	1.275	1.787	114
1.7%	6755	30	0	GET	^/api/images/[^/]+	794.735	0.118	1.096	3.569	186414
1.1%	1909	0	0	POST	/api/auth/login	501.082	0.262	1.391	2.150	247
0.7%	1610	0	0	GET	/api/memories	323.160	0.201	1.994	4.333	13134
0.5%	2026	0	0	GET	/api/posting-window/current	240.129	0.119	0.644	1.587	239
0.5%	340	0	0	PUT	/api/me/profile-image	235.088	0.691	2.721	2.886	0
0.5%	565	0	0	POST	/api/auth/register	228.125	0.404	1.488	1.550	229
0.5%	1038	0	0	GET	^/api/posts/me/current$	216.258	0.208	1.762	2.310	658
0.4%	1521	0	0	GET	/api/users	205.584	0.135	0.807	2.274	240
0.4%	706	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	179.250	0.254	1.311	1.553	116
0.4%	453	14	0	POST	^/api/posts/[^/]+/realmojis$	165.841	0.366	2.183	2.976	475
0.3%	357	0	0	PUT	^/api/follows/[^/]+$	122.296	0.343	1.298	1.805	0
0.3%	980	0	0	GET	^/api/relationships/[^/]+	120.997	0.123	1.050	2.275	596
0.3%	1145	34	0	GET	^/api/posts/[^/]+$	119.435	0.104	0.432	2.069	611
0.2%	283	0	0	PATCH	/api/me	114.103	0.403	1.667	1.774	283
0.2%	494	4	0	POST	^/api/posts/[^/]+/comments$	106.043	0.215	1.124	2.023	340
0.1%	324	0	0	POST	/api/auth/logout	66.498	0.205	1.289	1.731	0
0.1%	613	0	0	GET	^/api/users/[^/]+$	62.335	0.102	0.463	1.432	205
0.1%	558	2	0	GET	^/api/posts/[^/]+/realmojis$	58.768	0.105	0.448	1.704	3304
0.1%	116	2	0	PATCH	^/api/posts/[^/]+$	46.844	0.404	1.419	1.639	639
0.1%	479	0	0	GET	^/api/posts/[^/]+/comments$	44.087	0.092	0.351	0.436	4027
0.1%	351	2	0	GET	/api/me	40.412	0.115	0.669	0.811	303
0.1%	250	0	0	GET	^/api/memories/[^/]+	31.176	0.125	1.737	2.272	943
0.1%	73	0	0	DELETE	^/api/realmojis/[^/]+$	23.498	0.322	0.875	0.875	0
0.0%	113	0	0	GET	/api/me/export	20.879	0.185	3.158	3.212	2236989
0.0%	34	0	0	DELETE	^/api/posts/[^/]+$	17.742	0.522	1.996	1.996	0
0.0%	79	0	0	DELETE	^/api/comments/[^/]+$	12.794	0.162	0.541	0.541	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
85.5%	1129	1	0	GET	/api/notifications	29492.715	26.123	60.001	60.193	107
2.3%	5383	0	0	GET	^/api/feeds/[^/]+	802.231	0.149	1.296	1.906	3266
2.1%	1112	11	0	POST	/api/posts	728.886	0.655	2.447	7.735	653
2.1%	6130	32	0	GET	^/api/images/[^/]+	711.537	0.116	1.144	3.186	172251
1.4%	2494	20	0	POST	/api/friend-requests	477.910	0.192	0.875	1.549	114
1.2%	1691	0	0	POST	/api/auth/login	420.178	0.248	1.183	1.725	250
0.7%	1591	0	0	GET	/api/memories	235.741	0.148	1.551	2.420	14572
0.6%	430	0	0	POST	/api/auth/register	219.779	0.511	1.233	1.291	230
0.5%	1586	0	0	GET	/api/posting-window/current	176.254	0.111	0.867	1.548	240
0.5%	1079	0	0	GET	^/api/posts/me/current$	158.558	0.147	1.404	1.876	659
0.4%	1248	0	0	GET	/api/users	155.179	0.124	0.889	1.495	241
0.4%	294	16	0	POST	^/api/posts/[^/]+/realmojis$	122.441	0.416	2.113	2.640	458
0.3%	499	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	102.695	0.206	0.965	1.290	116
0.3%	636	0	0	GET	^/api/relationships/[^/]+	92.306	0.145	1.405	1.591	617
0.2%	733	29	0	GET	^/api/posts/[^/]+$	85.649	0.117	0.792	1.099	604
0.2%	251	0	0	PUT	/api/me/profile-image	75.035	0.299	1.185	1.263	0
0.2%	350	3	0	POST	^/api/posts/[^/]+/comments$	65.328	0.187	0.835	1.280	347
0.1%	318	0	0	PUT	^/api/follows/[^/]+$	50.423	0.159	0.567	0.731	0
0.1%	280	0	0	POST	/api/auth/logout	49.355	0.176	0.921	1.197	0
0.1%	210	0	0	PATCH	/api/me	41.176	0.196	0.874	1.230	282
0.1%	302	6	0	GET	^/api/posts/[^/]+/comments$	36.797	0.122	0.762	1.447	4041
0.1%	298	9	0	GET	^/api/posts/[^/]+/realmojis$	33.318	0.112	0.892	1.121	3104
0.1%	331	0	0	GET	^/api/users/[^/]+$	33.162	0.100	0.391	0.468	198
0.1%	214	0	0	GET	^/api/memories/[^/]+	27.818	0.130	1.043	2.356	911
0.1%	109	0	0	GET	/api/me/export	22.748	0.209	1.819	3.164	2972294
0.1%	43	0	0	DELETE	^/api/posts/[^/]+$	22.270	0.518	1.858	1.858	0
0.1%	84	0	0	PATCH	^/api/posts/[^/]+$	21.063	0.251	1.624	1.624	656
0.1%	53	0	0	DELETE	^/api/realmojis/[^/]+$	19.062	0.360	1.019	1.019	0
0.1%	260	1	0	GET	/api/me	18.878	0.073	0.153	0.345	303
0.0%	34	0	0	DELETE	^/api/comments/[^/]+$	5.350	0.157	0.342	0.342	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 11900 クエリ / 48 種類 / 合計 134.017s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	44.1%	270	59.159	219.11	828.22	1092.4	111.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	19.3%	4	25.851	6462.71	9529.69	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.2%	4	19.003	4750.85	7186.91	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.7%	7427	11.647	1.57	8.06	10.3	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.8%	4	9.063	2265.80	3449.74	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	4.1%	7	5.519	788.50	3145.14	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.6%	1	0.828	828.46	828.46	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
8	0.5%	7	0.617	88.10	161.76	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.4%	7	0.573	81.91	163.26	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.3%	1	0.403	403.47	403.47	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.2%	1	0.227	227.29	227.29	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.172	171.96	171.96	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.141	140.70	140.70	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.122	121.61	121.61	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
15	0.1%	7	0.113	16.17	26.88	166.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: 4671dd9b2d7b30d72cf0614948ec68d5f167a54b
Type: cpu
Time: 2026-10-31 16:33:04 JST
Duration: 60s, Total samples = 27.91s (46.52%)
Showing nodes accounting for 21.76s, 77.96% of 27.91s total
Dropped 877 nodes (cum <= 0.14s)
Showing top 40 nodes out of 236
      flat  flat%   sum%        cum   cum%
     4.65s 16.66% 16.66%      4.65s 16.66%  crypto/internal/fips140/sha256.blockAVX2
     2.90s 10.39% 27.05%      2.90s 10.39%  internal/runtime/syscall/linux.Syscall6
     1.75s  6.27% 33.32%      6.07s 21.75%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.53s  5.48% 38.80%      1.53s  5.48%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.50s  5.37% 44.18%      1.50s  5.37%  runtime.memmove
     1.12s  4.01% 48.19%      1.12s  4.01%  aeshashbody
     0.73s  2.62% 50.81%      0.73s  2.62%  runtime.memclrNoHeapPointers
     0.66s  2.36% 53.17%      0.90s  3.22%  example.com/isureal-company/isureal/jpegcheck.(*decoder).skipAC
     0.49s  1.76% 54.93%      1.92s  6.88%  runtime.mapaccess1_faststr
     0.46s  1.65% 56.57%      0.62s  2.22%  runtime.findObject
     0.46s  1.65% 58.22%      1.83s  6.56%  runtime.scanObject
     0.46s  1.65% 59.87%      0.58s  2.08%  runtime.tryDeferToSpanScan
     0.41s  1.47% 61.34%      0.62s  2.22%  example.com/isureal-company/isureal.postLess
     0.31s  1.11% 62.45%      5.30s 18.99%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.30s  1.07% 63.53%      0.30s  1.07%  sync/atomic.(*Int32).Add
     0.29s  1.04% 64.56%      0.29s  1.04%  runtime.futex
     0.26s  0.93% 65.50%      0.26s  0.93%  memeqbody
     0.26s  0.93% 66.43%      1.79s  6.41%  runtime.mapassign_faststr
     0.25s   0.9% 67.32%      5.13s 18.38%  crypto/internal/fips140/sha256.(*Digest).Write
     0.24s  0.86% 68.18%      0.24s  0.86%  example.com/isureal-company/isureal/jpegcheck.(*decoder).prefetch
     0.24s  0.86% 69.04%      1.26s  4.51%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.22s  0.79% 69.83%      0.36s  1.29%  database/sql.convertAssignRows
     0.20s  0.72% 70.55%      0.20s  0.72%  runtime.nextFreeFast (inline)
     0.18s  0.64% 71.19%      0.20s  0.72%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.18s  0.64% 71.84%      1.52s  5.45%  runtime.mallocgc
     0.16s  0.57% 72.41%      0.18s  0.64%  example.com/isureal-company/isureal.rmLess
     0.15s  0.54% 72.95%      0.15s  0.54%  cmpbody
     0.15s  0.54% 73.49%      0.35s  1.25%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.14s   0.5% 73.99%      5.53s 19.81%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.14s   0.5% 74.49%      0.61s  2.19%  runtime.mallocgcSmallScanNoHeader
     0.13s  0.47% 74.96%      0.15s  0.54%  runtime.(*mspan).writeHeapBitsSmall
     0.12s  0.43% 75.39%      0.43s  1.54%  runtime.mallocgcSmallNoscan
     0.11s  0.39% 75.78%      5.89s 21.10%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.10s  0.36% 76.14%      0.26s  0.93%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.10s  0.36% 76.50%      0.19s  0.68%  internal/runtime/maps.(*Iter).Next
     0.09s  0.32% 76.82%      0.27s  0.97%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.08s  0.29% 77.10%      4.73s 16.95%  crypto/internal/fips140/sha256.block
     0.08s  0.29% 77.39%      0.28s  1.00%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func12
     0.08s  0.29% 77.68%      0.14s   0.5%  runtime.greyobject
     0.08s  0.29% 77.96%      0.31s  1.11%  runtime.mapaccess2_faststr
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 28.8% / 100% | 7% | 0.9% | 6 | 157 MB | run queue が積もっている |
| isu2 (2 cpu) | 21.6% / 100% | 6% | 0.9% | 5 | 109 MB | run queue が積もっている |
| isu3 (2 cpu) | 21.1% / 100% | 5% | 1.0% | 4 | 167 MB |  |
| isu4 (2 cpu) | 53.2% / 100% | 45% | 2.0% | 8 | 102 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 43% / nginx 8% |
| isu2 | isureal 38% / nginx 3% |
| isu3 | isureal 38% / nginx 2% |
| isu4 | isureal 87% / postgres 13% |

## Diff from previous run

スコア: 201497 → 223742 (+22245)

`40e65ef..6f2cd6f`

```
 webapp/go/jpegcheck/check_test.go |  16 +++++
 webapp/go/jpegcheck/huffman.go    |  56 +++++++++++++++++-
 webapp/go/jpegcheck/scan.go       | 121 ++++++++++++++++++++++++++++++--------
 3 files changed, 165 insertions(+), 28 deletions(-)
```
