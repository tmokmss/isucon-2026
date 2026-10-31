# run 20261031-160409-40e65ef

## Summary

| | |
|---|---|
| score | 201497 |
| status | pass |
| ref | asia2-replica |
| sha | 40e65ef68604 |
| 所要時間 | 543s |
| 前回比 | +18166 (9.9%) vs 20261031-152137-8c13252 |
| ベスト比 | +9681 (5.0%) vs 20261031-151339-5a230a4 = 191816 **ベスト更新** |
| note | asia2 as asia read replica, asia1 nginx sends GET to asia2 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	149	401	679	679
europe	7	18	49	144	405	404
america	6	18	52	151	403	402
```

```
{"id":"26713","status":"done","portalStatus":"finished","passed":true,"score":201497,"raw":201497,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:11:32.712Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:11:35.704Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:11:38.705Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:11:40.161Z level=INFO msg="user moved" user=ishi-5127 from=america to=europe
time=2026-10-31T07:11:41.740Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T07:11:44.719Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T07:11:45.878Z level=INFO msg="user moved" user=ishi-5127 from=europe to=america
time=2026-10-31T07:11:47.705Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T07:11:50.716Z level=INFO msg="asia 3 日目: 149 人"
time=2026-10-31T07:11:52.940Z level=INFO msg="user moved" user=suni.5533 from=asia to=america
time=2026-10-31T07:11:53.704Z level=INFO msg="europe 3 日目: 49 人"
time=2026-10-31T07:11:53.727Z level=INFO msg="user moved" user=suni.5533 from=america to=europe
time=2026-10-31T07:11:54.447Z level=INFO msg="user moved" user=ivymanoa_5495 from=america to=asia
time=2026-10-31T07:11:56.704Z level=INFO msg="america 3 日目: 52 人"
time=2026-10-31T07:11:57.401Z level=INFO msg="user moved" user=niel_5858 from=asia to=europe
time=2026-10-31T07:11:57.616Z level=INFO msg="user moved" user=ivymanoa_5495 from=asia to=europe
time=2026-10-31T07:11:58.312Z level=INFO msg="user moved" user=chita.6255 from=europe to=america
time=2026-10-31T07:11:59.718Z level=INFO msg="asia 4 日目: 401 人"
time=2026-10-31T07:11:59.947Z level=INFO msg="user moved" user=tsuyuho6463 from=europe to=america
time=2026-10-31T07:12:02.706Z level=INFO msg="europe 4 日目: 144 人"
time=2026-10-31T07:12:04.257Z level=INFO msg="user moved" user=suni.5533 from=europe to=asia
time=2026-10-31T07:12:06.116Z level=INFO msg="america 4 日目: 151 人"
time=2026-10-31T07:12:07.823Z level=INFO msg="user moved" user=rivinso_8725 from=europe to=america
time=2026-10-31T07:12:07.979Z level=INFO msg="user moved" user=liasaeva_8418 from=asia to=america
time=2026-10-31T07:12:08.085Z level=INFO msg="user moved" user=solfu_6708 from=asia to=europe
time=2026-10-31T07:12:08.665Z level=INFO msg="user moved" user=chifumia9113 from=europe to=america
time=2026-10-31T07:12:08.721Z level=INFO msg="asia 5 日目: 679 人"
time=2026-10-31T07:12:09.227Z level=INFO msg="user moved" user=niel_5858 from=europe to=america
time=2026-10-31T07:12:11.349Z level=INFO msg="user moved" user=chita.6255 from=america to=europe
time=2026-10-31T07:12:12.421Z level=INFO msg="europe 5 日目: 405 人"
time=2026-10-31T07:12:13.061Z level=INFO msg="user moved" user=tsuyuho6463 from=america to=asia
time=2026-10-31T07:12:14.901Z level=INFO msg="america 5 日目: 403 人"
time=2026-10-31T07:12:23.960Z level=INFO msg="asia 最終日（グランプリ）: 679 人"
time=2026-10-31T07:12:24.124Z level=INFO msg="america 最終日（グランプリ）: 402 人"
time=2026-10-31T07:12:24.152Z level=INFO msg="europe 最終日（グランプリ）: 404 人"
time=2026-10-31T07:12:32.787Z level=INFO msg="グランプリ: アジア 373 人, ヨーロッパ 393 人, アメリカ 267 人"
{
  "pass": true,
  "score": 201497,
  "messages": [
    "kakaeva-5108: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "isheki-5039: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "danchi.14262: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "fusozoe_14401: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marmina_14023: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tormema-15135: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "miwaki_13779: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "liaho_8793: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "shirata_12780: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sardines_en_cocotte_de_terre_2092: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "roro.5605: ログインし直せません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "mekazoe_11470: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.36.139:54358-\u003e52.69.202.109:8080: write tcp 10.27.36.139:54358-\u003e52.69.202.109:8080: write: broken pipe",
    "chiel.13168: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "tsuyuho6463: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "tenu6924: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.6%	1344	1	0	GET	/api/notifications	13467.242	10.020	15.268	15.287	172
15.8%	1986	73	0	POST	/api/posts	4879.851	2.457	61.395	63.977	629
8.6%	10345	0	0	GET	^/api/feeds/[^/]+	2650.016	0.256	1.434	2.729	5261
7.3%	2992	0	0	GET	/api/memories	2238.422	0.748	6.202	11.164	11014
4.7%	2701	0	0	GET	/api/posting-window/current	1443.742	0.535	1.673	2.728	236
2.7%	1908	0	0	GET	/api/users	847.218	0.444	1.933	3.024	241
2.6%	1633	1	0	GET	^/api/posts/me/current$	806.407	0.494	1.558	2.796	652
2.1%	11961	276	0	GET	^/api/images/[^/]+	647.612	0.054	0.673	1.877	159447
2.0%	218	0	0	GET	/api/me/export	628.038	2.881	16.838	19.336	2010169
1.8%	511	123	0	PUT	/api/me/profile-image	542.809	1.062	12.171	15.390	4
1.3%	629	11	0	POST	^/api/posts/[^/]+/realmojis$	386.883	0.615	6.379	6.665	468
1.0%	1294	0	0	GET	^/api/relationships/[^/]+	313.329	0.242	1.437	2.063	1567
1.0%	4041	45	0	POST	/api/friend-requests	298.597	0.074	0.555	1.224	114
0.9%	1824	77	0	GET	^/api/posts/[^/]+$	275.518	0.151	0.928	1.516	598
0.7%	400	3	0	GET	/api/me	230.231	0.576	1.697	2.253	299
0.7%	3349	8	0	POST	^/api/friend-requests/[^/]+/[^/]+$	222.772	0.067	0.525	0.841	116
0.7%	863	5	0	GET	^/api/posts/[^/]+/realmojis$	207.448	0.240	1.647	2.346	1598
0.5%	495	0	0	GET	^/api/memories/[^/]+	167.653	0.339	2.017	3.198	608
0.4%	907	0	0	GET	^/api/users/[^/]+$	125.135	0.138	0.917	1.105	200
0.4%	4033	593	0	POST	/api/auth/login	121.025	0.030	0.487	0.737	211
0.3%	756	6	0	GET	^/api/posts/[^/]+/comments$	104.093	0.138	0.864	1.028	2587
0.2%	731	1	0	POST	^/api/posts/[^/]+/comments$	70.036	0.096	0.492	0.810	348
0.1%	76	0	0	DELETE	^/api/posts/[^/]+$	34.067	0.448	1.164	1.164	0
0.1%	698	1	0	POST	/api/auth/register	24.885	0.036	0.584	0.699	227
0.1%	538	0	0	POST	/api/auth/logout	23.859	0.044	0.474	0.633	0
0.1%	1	0	0	POST	/api/initialize	22.975	22.975	22.975	22.975	90
0.1%	94	0	0	DELETE	^/api/realmojis/[^/]+$	20.040	0.213	0.831	0.831	0
0.1%	165	0	0	PATCH	^/api/posts/[^/]+$	16.967	0.103	0.588	0.625	660
0.0%	464	0	0	PUT	^/api/follows/[^/]+$	14.932	0.032	0.490	0.603	0
0.0%	340	0	0	PATCH	/api/me	12.245	0.036	0.558	0.633	276
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.6%	651	1	0	GET	/api/notifications	4727.637	7.262	11.753	11.754	166
11.9%	1023	14	0	POST	/api/posts	1265.646	1.237	60.326	63.463	644
6.4%	6273	0	0	GET	^/api/feeds/[^/]+	681.921	0.109	1.259	1.946	3987
5.8%	6423	24	0	GET	^/api/images/[^/]+	617.386	0.096	0.852	3.500	171973
5.5%	228	0	0	PUT	/api/me/profile-image	580.421	2.546	6.208	6.380	0
5.4%	2459	23	0	POST	/api/friend-requests	569.275	0.232	0.742	1.573	114
3.4%	1855	0	0	POST	/api/auth/login	358.119	0.193	0.738	1.442	249
2.8%	375	9	0	POST	^/api/posts/[^/]+/realmojis$	295.550	0.788	6.257	6.718	459
2.2%	1664	0	0	GET	/api/memories	236.341	0.142	1.616	2.532	13668
1.2%	998	2	0	GET	^/api/posts/me/current$	130.853	0.131	1.522	1.928	652
1.2%	521	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	122.650	0.235	0.860	1.576	116
1.2%	1464	0	0	GET	/api/posting-window/current	122.389	0.084	0.339	1.736	237
1.0%	1177	0	0	GET	/api/users	101.239	0.086	0.264	1.209	241
0.9%	404	0	0	POST	/api/auth/register	93.600	0.232	0.683	0.754	228
0.9%	394	4	0	POST	^/api/posts/[^/]+/comments$	90.913	0.231	0.804	1.323	351
0.8%	981	22	0	GET	^/api/posts/[^/]+$	84.973	0.087	0.280	1.577	610
0.7%	792	0	0	GET	^/api/relationships/[^/]+	75.380	0.095	0.611	1.762	494
0.6%	323	0	0	POST	/api/auth/logout	60.338	0.187	0.714	1.066	0
0.5%	240	0	0	PUT	^/api/follows/[^/]+$	56.273	0.234	0.660	0.779	0
0.5%	220	0	0	PATCH	/api/me	49.371	0.224	0.653	0.810	280
0.4%	515	0	0	GET	^/api/users/[^/]+$	42.448	0.082	0.272	0.951	205
0.4%	447	0	0	GET	^/api/posts/[^/]+/realmojis$	38.225	0.086	0.565	1.321	2032
0.3%	387	0	0	GET	^/api/posts/[^/]+/comments$	32.877	0.085	0.278	1.181	3020
0.3%	107	1	0	PATCH	^/api/posts/[^/]+$	30.110	0.281	1.185	1.335	646
0.3%	40	0	0	DELETE	^/api/posts/[^/]+$	27.018	0.675	1.746	1.746	0
0.2%	274	0	0	GET	^/api/memories/[^/]+	24.923	0.091	0.339	1.859	998
0.2%	61	0	0	DELETE	^/api/realmojis/[^/]+$	21.425	0.351	0.889	0.889	0
0.2%	91	0	0	DELETE	^/api/comments/[^/]+$	21.177	0.233	0.709	0.709	0
0.2%	237	2	0	GET	/api/me	19.724	0.083	0.204	0.235	300
0.1%	106	0	0	GET	/api/me/export	11.663	0.110	0.824	0.958	2413117
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.1%	642	0	0	GET	/api/notifications	4056.666	6.319	9.454	9.454	169
8.3%	1005	19	0	POST	/api/posts	779.607	0.776	2.954	7.671	643
8.0%	4797	0	0	GET	^/api/feeds/[^/]+	753.081	0.157	1.939	2.623	3403
7.6%	5543	31	0	GET	^/api/images/[^/]+	717.130	0.129	0.981	3.861	152148
5.5%	2274	23	0	POST	/api/friend-requests	516.933	0.227	0.793	2.119	114
3.9%	229	0	0	PUT	/api/me/profile-image	370.269	1.617	6.284	6.435	0
3.9%	1619	0	0	POST	/api/auth/login	368.076	0.227	0.829	2.217	250
3.0%	1388	0	0	GET	/api/memories	282.086	0.203	2.309	3.496	15586
2.6%	325	29	0	POST	^/api/posts/[^/]+/realmojis$	249.134	0.767	5.665	7.242	437
1.9%	1441	0	0	GET	/api/posting-window/current	177.157	0.123	0.636	2.538	240
1.6%	793	9	0	GET	^/api/posts/me/current$	153.755	0.194	1.946	2.390	648
1.6%	1097	0	0	GET	/api/users	149.575	0.136	1.566	2.511	239
1.4%	399	0	0	POST	/api/auth/register	132.123	0.331	0.976	1.055	230
0.9%	408	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	88.211	0.216	0.846	1.712	116
0.9%	576	0	0	GET	^/api/relationships/[^/]+	84.785	0.147	2.107	2.569	611
0.8%	703	35	0	GET	^/api/posts/[^/]+$	76.089	0.108	0.502	2.190	592
0.7%	285	1	0	POST	^/api/posts/[^/]+/comments$	65.741	0.231	0.646	1.206	344
0.6%	267	0	0	POST	/api/auth/logout	59.440	0.223	1.555	1.981	0
0.5%	234	0	0	PUT	^/api/follows/[^/]+$	47.720	0.204	0.617	0.793	0
0.5%	198	0	0	PATCH	/api/me	43.465	0.220	0.853	0.853	279
0.4%	318	0	0	GET	^/api/users/[^/]+$	34.348	0.108	0.299	1.959	195
0.3%	112	0	0	GET	/api/me/export	32.066	0.286	3.707	4.632	2858538
0.3%	255	0	0	GET	^/api/posts/[^/]+/realmojis$	26.719	0.105	0.309	0.329	2590
0.3%	247	0	0	GET	^/api/posts/[^/]+/comments$	26.360	0.107	0.590	1.310	3423
0.2%	225	0	0	GET	^/api/memories/[^/]+	22.868	0.102	0.357	2.274	1064
0.2%	50	0	0	DELETE	^/api/realmojis/[^/]+$	22.464	0.449	2.263	2.263	0
0.2%	234	1	0	GET	/api/me	22.156	0.095	0.259	0.297	304
0.2%	75	1	0	PATCH	^/api/posts/[^/]+$	20.072	0.268	1.486	1.486	635
0.2%	34	0	0	DELETE	^/api/posts/[^/]+$	19.084	0.561	2.291	2.291	0
0.1%	46	0	0	DELETE	^/api/comments/[^/]+$	10.113	0.220	0.664	0.664	0
…（上位 30 / 全 36 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 11199 クエリ / 48 種類 / 合計 137.801s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	45.3%	225	62.460	277.60	1409.24	1300.4	132.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	18.8%	4	25.846	6461.52	9540.75	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	13.9%	4	19.091	4772.82	7217.61	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.6%	7458	11.903	1.60	9.45	10.3	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.6%	4	9.145	2286.29	3443.22	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	4.1%	6	5.583	930.54	3636.37	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	1	0.562	561.92	561.92	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.3%	6	0.454	75.62	162.45	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.3%	6	0.448	74.62	157.77	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
10	0.3%	1	0.409	408.52	408.52	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.3%	1	0.374	373.61	373.61	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.3%	1	0.362	362.47	362.47	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.2%	1	0.333	333.00	333.00	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.1%	1	0.137	136.59	136.59	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
15	0.1%	1	0.117	117.12	117.12	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: 3bbcaac7e85d7e3cf4d9adc5d7e6e2cea5907c6d
Type: cpu
Time: 2026-10-31 16:11:05 JST
Duration: 60.10s, Total samples = 27.53s (45.81%)
Showing nodes accounting for 21.09s, 76.61% of 27.53s total
Dropped 853 nodes (cum <= 0.14s)
Showing top 40 nodes out of 245
      flat  flat%   sum%        cum   cum%
     4.12s 14.97% 14.97%      4.12s 14.97%  crypto/internal/fips140/sha256.blockAVX2
     2.51s  9.12% 24.08%      2.51s  9.12%  internal/runtime/syscall/linux.Syscall6
     1.59s  5.78% 29.86%      5.75s 20.89%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.51s  5.48% 35.34%      1.51s  5.48%  runtime.memmove
     1.34s  4.87% 40.21%      1.34s  4.87%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.07s  3.89% 44.10%      1.07s  3.89%  aeshashbody
     0.81s  2.94% 47.04%         1s  3.63%  runtime.findObject
     0.71s  2.58% 49.62%      0.71s  2.58%  runtime.memclrNoHeapPointers
     0.61s  2.22% 51.83%      2.62s  9.52%  runtime.scanObject
     0.61s  2.22% 54.05%      1.01s  3.67%  runtime.tryDeferToSpanScan
     0.49s  1.78% 55.83%      0.70s  2.54%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.43s  1.56% 57.39%      1.93s  7.01%  runtime.mapaccess1_faststr
     0.41s  1.49% 58.88%      1.39s  5.05%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.37s  1.34% 60.23%      0.37s  1.34%  runtime.futex
     0.36s  1.31% 61.53%      0.58s  2.11%  example.com/isureal-company/isureal.postLess
     0.32s  1.16% 62.70%      0.32s  1.16%  memeqbody
     0.31s  1.13% 63.82%      0.43s  1.56%  database/sql.convertAssignRows
     0.27s  0.98% 64.80%      0.27s  0.98%  runtime.spanClass.sizeclass (inline)
     0.26s  0.94% 65.75%      0.28s  1.02%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.22s   0.8% 66.55%      1.40s  5.09%  runtime.mapassign_faststr
     0.22s   0.8% 67.34%      0.22s   0.8%  runtime.nextFreeFast (inline)
     0.22s   0.8% 68.14%      0.22s   0.8%  sync/atomic.(*Int32).Add (inline)
     0.21s  0.76% 68.91%      1.77s  6.43%  runtime.mallocgc
     0.20s  0.73% 69.63%      4.58s 16.64%  crypto/internal/fips140/sha256.(*Digest).Write
     0.18s  0.65% 70.29%      0.19s  0.69%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.18s  0.65% 70.94%      0.22s   0.8%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.17s  0.62% 71.56%      4.57s 16.60%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.16s  0.58% 72.14%      4.81s 17.47%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.16s  0.58% 72.72%      0.82s  2.98%  runtime.mallocgcSmallScanNoHeader
     0.14s  0.51% 73.23%      0.17s  0.62%  runtime.(*mspan).writeHeapBitsSmall
     0.13s  0.47% 73.70%      5.13s 18.63%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.12s  0.44% 74.14%      1.90s  6.90%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s   0.4% 74.54%      1.12s  4.07%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.10s  0.36% 74.90%      3.99s 14.49%  example.com/isureal-company/isureal.loadBig
     0.10s  0.36% 75.26%      0.43s  1.56%  runtime.mapaccess2_faststr
     0.09s  0.33% 75.59%      0.22s   0.8%  internal/runtime/maps.(*Iter).Next
     0.09s  0.33% 75.92%      0.45s  1.63%  runtime.mallocgcSmallNoscan
     0.07s  0.25% 76.17%      0.16s  0.58%  runtime.greyobject
     0.06s  0.22% 76.39%      0.49s  1.78%  database/sql.(*Rows).scanLocked
     0.06s  0.22% 76.61%      0.16s  0.58%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 40.3% / 100% | 18% | 0.5% | 7 | 113 MB | run queue が積もっている |
| isu2 (2 cpu) | 22.6% / 100% | 3% | 0.6% | 5 | 202 MB | run queue が積もっている |
| isu3 (2 cpu) | 25.6% / 100% | 7% | 1.0% | 5 | 187 MB | run queue が積もっている |
| isu4 (2 cpu) | 63.6% / 100% | 54% | 2.5% | 11 | 105 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 61% / nginx 9% |
| isu2 | isureal 41% / nginx 2% |
| isu3 | isureal 47% / nginx 2% |
| isu4 | isureal 101% / postgres 15% |

## Diff from previous run

スコア: 183331 → 201497 (+18166)

`8c13252..40e65ef`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf  | 25 ++++++++++++-
 conf/isu1/home/isucon/env.sh                    |  2 +-
 conf/isu4/etc/systemd/system/isureal-go.service | 11 +++---
 conf/isu4/home/isucon/env.sh                    | 12 +++++++
 webapp/go/admin_api.go                          | 27 +++++++++++---
 webapp/go/export_api.go                         | 12 +++++--
 webapp/go/image_api.go                          | 23 ++++++------
 webapp/go/image_file.go                         | 33 +++++++++--------
 webapp/go/image_held.go                         | 10 ++----
 webapp/go/image_local.go                        | 47 ++++++++++---------------
 webapp/go/ip_ban.go                             | 11 ++++++
 webapp/go/persist.go                            |  2 +-
 webapp/go/realmoji_api.go                       | 41 +++++++++------------
 webapp/go/repl.go                               |  1 +
 webapp/go/replica_server.go                     | 13 +++----
 webapp/go/store.go                              |  3 +-
 16 files changed, 163 insertions(+), 110 deletions(-)
```
