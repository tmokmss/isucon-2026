# run 20261031-172337-7534c3b

## Summary

| | |
|---|---|
| score | 220719 |
| status | pass |
| ref | HEAD |
| sha | 7534c3bfda5a |
| 所要時間 | 415s |
| 前回比 | +990 (0.5%) vs 20261031-171010-53bad13 |
| ベスト比 | -11391 (-4.9%) vs 20261031-164055-bdf05cd = 232110 |
| note | jpegcheck-prefetch2 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	148	395	639	638
europe	7	16	46	132	383	382
america	6	18	54	154	436	435
```

```
{"id":"34484","status":"done","portalStatus":"finished","passed":true,"score":220719,"raw":220719,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:28:39.595Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:28:42.584Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:28:45.585Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:28:47.918Z level=INFO msg="user moved" user=rebense5096 from=europe to=asia
time=2026-10-31T08:28:48.602Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T08:28:51.585Z level=INFO msg="europe 2 日目: 16 人"
time=2026-10-31T08:28:51.639Z level=INFO msg="user moved" user=rebense5096 from=asia to=europe
time=2026-10-31T08:28:54.587Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T08:28:57.598Z level=INFO msg="asia 3 日目: 148 人"
time=2026-10-31T08:29:00.267Z level=INFO msg="user moved" user=elre.5504 from=america to=europe
time=2026-10-31T08:29:00.584Z level=INFO msg="europe 3 日目: 46 人"
time=2026-10-31T08:29:03.359Z level=INFO msg="user moved" user=neevaan.6209 from=asia to=america
time=2026-10-31T08:29:03.618Z level=INFO msg="america 3 日目: 54 人"
time=2026-10-31T08:29:06.007Z level=INFO msg="user moved" user=elre.5504 from=europe to=asia
time=2026-10-31T08:29:06.602Z level=INFO msg="asia 4 日目: 395 人"
time=2026-10-31T08:29:06.985Z level=INFO msg="user moved" user=rebense5096 from=europe to=asia
time=2026-10-31T08:29:09.587Z level=INFO msg="europe 4 日目: 132 人"
time=2026-10-31T08:29:12.592Z level=INFO msg="america 4 日目: 154 人"
time=2026-10-31T08:29:13.418Z level=INFO msg="user moved" user=toor_7688 from=asia to=europe
time=2026-10-31T08:29:14.724Z level=INFO msg="user moved" user=mardan.8953 from=europe to=america
time=2026-10-31T08:29:15.603Z level=INFO msg="asia 5 日目: 639 人"
time=2026-10-31T08:29:16.522Z level=INFO msg="user moved" user=marsol.8993 from=america to=europe
time=2026-10-31T08:29:17.923Z level=INFO msg="user moved" user=honolia9973 from=america to=europe
time=2026-10-31T08:29:18.674Z level=INFO msg="europe 5 日目: 383 人"
time=2026-10-31T08:29:21.643Z level=INFO msg="america 5 日目: 436 人"
time=2026-10-31T08:29:30.601Z level=INFO msg="europe 最終日（グランプリ）: 382 人"
time=2026-10-31T08:29:30.633Z level=INFO msg="asia 最終日（グランプリ）: 638 人"
time=2026-10-31T08:29:30.689Z level=INFO msg="america 最終日（グランプリ）: 435 人"
time=2026-10-31T08:29:39.648Z level=INFO msg="グランプリ: アジア 460 人, ヨーロッパ 377 人, アメリカ 427 人"
{
  "pass": true,
  "score": 220719,
  "messages": [
    "noanoa.7674: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "iwashi_pilaf_stew_3670: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "rezoe7431: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "qiaomai_liuchang_rust_2363: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "tavin-6108: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "meyu8333: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
78.6%	1898	0	0	GET	/api/notifications	49345.258	25.999	60.005	60.005	118
4.8%	10966	0	0	GET	^/api/feeds/[^/]+	2997.314	0.273	1.435	3.308	5295
3.5%	2988	0	0	GET	/api/memories	2203.667	0.738	5.967	11.890	10747
2.3%	2592	0	0	GET	/api/posting-window/current	1453.060	0.561	1.460	3.156	237
2.3%	1890	27	0	POST	/api/posts	1427.572	0.755	3.675	8.015	640
1.4%	1692	1	0	GET	^/api/posts/me/current$	908.467	0.537	1.647	2.505	649
1.3%	1753	0	0	GET	/api/users	810.875	0.463	1.562	1.690	236
1.1%	205	0	0	GET	/api/me/export	679.896	3.317	18.507	21.250	2026913
0.7%	11902	243	0	GET	^/api/images/[^/]+	468.917	0.039	0.695	1.806	168673
0.7%	480	127	0	PUT	/api/me/profile-image	435.012	0.906	14.029	16.621	4
0.5%	1333	0	0	GET	^/api/relationships/[^/]+	336.626	0.253	1.328	2.119	1519
0.5%	938	4	0	GET	^/api/posts/[^/]+/realmojis$	324.598	0.346	1.415	2.357	2999
0.5%	1997	62	0	GET	^/api/posts/[^/]+$	288.619	0.145	0.948	1.426	608
0.3%	366	3	0	GET	/api/me	213.287	0.583	1.353	1.430	299
0.3%	797	10	0	POST	^/api/posts/[^/]+/realmojis$	166.759	0.209	2.364	3.343	468
0.3%	508	0	0	GET	^/api/memories/[^/]+	164.657	0.324	1.464	1.981	693
0.2%	1000	0	0	GET	^/api/users/[^/]+$	145.159	0.145	0.923	1.250	203
0.2%	766	5	0	GET	^/api/posts/[^/]+/comments$	122.999	0.161	0.946	1.464	3665
0.1%	3888	530	0	POST	/api/auth/login	65.355	0.017	0.494	1.107	214
0.1%	3069	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	60.168	0.020	0.419	0.965	116
0.1%	3763	62	0	POST	/api/friend-requests	58.562	0.016	0.215	0.767	114
0.0%	83	0	0	DELETE	^/api/posts/[^/]+$	29.357	0.354	1.420	1.420	0
0.0%	136	0	0	DELETE	^/api/realmojis/[^/]+$	25.669	0.189	0.540	0.624	0
0.0%	1	0	0	POST	/api/initialize	23.038	23.038	23.038	23.038	90
0.0%	171	0	0	PATCH	^/api/posts/[^/]+$	16.570	0.097	0.913	0.990	655
0.0%	676	0	0	POST	^/api/posts/[^/]+/comments$	13.109	0.019	0.265	0.691	339
0.0%	564	0	0	POST	/api/auth/logout	10.718	0.019	0.430	0.717	0
0.0%	454	0	0	PUT	^/api/follows/[^/]+$	7.523	0.017	0.144	0.543	0
0.0%	655	1	0	POST	/api/auth/register	7.260	0.011	0.170	0.250	227
0.0%	314	0	0	PATCH	/api/me	4.891	0.016	0.155	0.218	275
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
86.0%	969	1	0	GET	/api/notifications	26466.697	27.313	60.001	60.475	106
2.4%	7219	0	0	GET	^/api/feeds/[^/]+	747.853	0.104	0.859	1.327	4510
2.2%	6622	33	0	GET	^/api/images/[^/]+	663.101	0.100	0.960	2.216	179370
1.4%	2360	27	0	POST	/api/friend-requests	431.149	0.183	0.710	1.126	114
1.1%	970	3	0	POST	/api/posts	333.090	0.343	1.355	1.591	653
0.9%	1851	0	0	POST	/api/auth/login	279.869	0.151	0.578	0.955	249
0.8%	1627	0	0	GET	/api/memories	237.950	0.146	1.211	1.923	13887
0.7%	428	5	0	POST	^/api/posts/[^/]+/realmojis$	225.961	0.528	3.295	3.587	474
0.5%	957	0	0	GET	^/api/posts/me/current$	144.124	0.151	0.994	1.241	655
0.5%	1391	0	0	GET	/api/posting-window/current	143.671	0.103	0.565	1.195	239
0.4%	1055	0	0	GET	/api/users	124.231	0.118	0.582	1.166	244
0.4%	674	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	112.836	0.167	0.753	0.960	116
0.3%	1321	11	0	GET	^/api/posts/[^/]+$	103.192	0.078	0.371	0.978	622
0.3%	874	0	0	GET	^/api/relationships/[^/]+	83.813	0.096	0.686	1.092	650
0.3%	229	0	0	PUT	/api/me/profile-image	77.151	0.337	0.773	3.504	0
0.3%	380	0	0	POST	/api/auth/register	77.115	0.203	0.446	0.500	229
0.2%	388	2	0	POST	^/api/posts/[^/]+/comments$	72.897	0.188	0.854	0.884	345
0.2%	361	0	0	POST	/api/auth/logout	53.276	0.148	0.705	0.952	0
0.2%	278	0	0	PUT	^/api/follows/[^/]+$	50.356	0.181	0.638	0.759	0
0.2%	594	0	0	GET	^/api/users/[^/]+$	46.398	0.078	0.511	1.243	206
0.2%	594	0	0	GET	^/api/posts/[^/]+/realmojis$	46.233	0.078	0.378	1.188	3184
0.1%	120	0	0	PATCH	^/api/posts/[^/]+$	36.808	0.307	0.985	1.088	655
0.1%	470	0	0	GET	^/api/posts/[^/]+/comments$	35.263	0.075	0.362	0.565	4077
0.1%	52	0	0	DELETE	^/api/posts/[^/]+$	34.582	0.665	1.519	1.519	0
0.1%	199	0	0	PATCH	/api/me	34.580	0.174	0.653	0.844	278
0.1%	278	0	0	GET	^/api/memories/[^/]+	25.529	0.092	1.020	1.471	895
0.1%	74	0	0	DELETE	^/api/realmojis/[^/]+$	23.345	0.315	1.461	1.461	0
0.1%	132	0	0	GET	/api/me/export	23.139	0.175	2.260	3.543	2301228
0.1%	234	2	0	GET	/api/me	22.892	0.098	0.457	0.566	302
0.0%	66	0	0	DELETE	^/api/comments/[^/]+$	10.184	0.154	0.648	0.648	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
85.2%	1107	1	0	GET	/api/notifications	29507.139	26.655	60.001	60.001	108
2.5%	6502	0	0	GET	^/api/feeds/[^/]+	863.286	0.133	1.400	2.065	3118
2.3%	6817	34	0	GET	^/api/images/[^/]+	801.807	0.118	1.098	3.398	179789
2.0%	1110	0	0	POST	/api/posts	701.291	0.632	2.545	3.018	654
1.5%	2593	35	0	POST	/api/friend-requests	503.950	0.194	0.697	1.793	114
0.9%	1897	0	0	POST	/api/auth/login	312.230	0.165	0.729	1.799	250
0.8%	1741	0	0	GET	/api/memories	266.321	0.153	1.493	2.511	14654
0.5%	1585	0	0	GET	/api/posting-window/current	177.040	0.112	0.712	1.762	240
0.5%	1247	0	0	GET	/api/users	171.865	0.138	0.759	1.836	242
0.5%	1096	0	0	GET	^/api/posts/me/current$	170.961	0.156	1.408	1.789	654
0.5%	360	12	0	POST	^/api/posts/[^/]+/realmojis$	160.744	0.447	2.656	4.599	442
0.3%	869	0	0	GET	^/api/relationships/[^/]+	104.683	0.120	1.077	1.895	525
0.3%	553	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	101.912	0.184	0.790	1.693	116
0.3%	941	31	0	GET	^/api/posts/[^/]+$	99.910	0.106	0.679	1.980	603
0.3%	436	0	0	POST	/api/auth/register	95.259	0.218	0.685	0.733	230
0.2%	235	0	0	PUT	/api/me/profile-image	75.218	0.320	1.855	3.376	0
0.2%	358	3	0	POST	^/api/posts/[^/]+/comments$	63.665	0.178	0.686	1.062	344
0.2%	337	0	0	POST	/api/auth/logout	56.026	0.166	0.782	1.817	0
0.2%	312	0	0	PUT	^/api/follows/[^/]+$	54.547	0.175	0.672	0.729	0
0.2%	457	0	0	GET	^/api/users/[^/]+$	53.349	0.117	0.806	1.619	199
0.1%	134	0	0	PATCH	^/api/posts/[^/]+$	41.933	0.313	1.253	1.260	655
0.1%	210	0	0	PATCH	/api/me	37.325	0.178	0.723	0.808	280
0.1%	362	1	0	GET	^/api/posts/[^/]+/comments$	36.216	0.100	0.571	0.728	5251
0.1%	377	0	0	GET	^/api/posts/[^/]+/realmojis$	35.145	0.093	0.358	0.751	3976
0.1%	286	0	0	GET	^/api/memories/[^/]+	31.896	0.112	1.393	1.839	1058
0.1%	54	0	0	DELETE	^/api/posts/[^/]+$	31.629	0.586	2.129	2.129	0
0.1%	114	0	0	GET	/api/me/export	24.978	0.219	5.060	5.419	3156900
0.1%	237	1	0	GET	/api/me	23.223	0.098	0.594	0.682	304
0.0%	52	0	0	DELETE	^/api/realmojis/[^/]+$	14.826	0.285	0.616	0.616	0
0.0%	53	0	0	DELETE	^/api/comments/[^/]+$	8.500	0.160	0.657	0.657	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 11505 クエリ / 48 種類 / 合計 139.576s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	43.7%	214	61.002	285.06	2502.09	1298.8	132.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	19.1%	4	26.691	6672.81	10090.25	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	13.9%	4	19.338	4834.46	7225.84	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.8%	7292	12.266	1.68	13.68	10.2	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.5%	4	9.109	2277.21	3436.27	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	5.1%	7	7.180	1025.65	3925.45	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.5%	1	0.640	640.18	640.18	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
8	0.4%	7	0.613	87.62	161.69	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.4%	7	0.602	86.02	161.45	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.3%	1	0.448	447.76	447.76	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.3%	1	0.357	357.19	357.19	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.2%	1	0.247	247.31	247.31	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.207	207.36	207.36	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.147	146.52	146.52	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
15	0.1%	1	0.118	118.23	118.23	2412.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…

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
Build ID: 321b888ec927c8694b4575e8c54a6a9697b75911
Type: cpu
Time: 2026-10-31 17:28:11 JST
Duration: 60s, Total samples = 24.60s (41.00%)
Showing nodes accounting for 19.40s, 78.86% of 24.60s total
Dropped 723 nodes (cum <= 0.12s)
Showing top 40 nodes out of 231
      flat  flat%   sum%        cum   cum%
     4.39s 17.85% 17.85%      4.39s 17.85%  crypto/internal/fips140/sha256.blockAVX2
     2.23s  9.07% 26.91%      2.23s  9.07%  internal/runtime/syscall/linux.Syscall6
     1.48s  6.02% 32.93%      1.48s  6.02%  runtime.memmove
     1.32s  5.37% 38.29%      5.75s 23.37%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.31s  5.33% 43.62%      1.31s  5.33%  aeshashbody
     1.22s  4.96% 48.58%      1.22s  4.96%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.79s  3.21% 51.79%      0.79s  3.21%  runtime.memclrNoHeapPointers
     0.61s  2.48% 54.27%      0.79s  3.21%  runtime.findObject
     0.53s  2.15% 56.42%      2.26s  9.19%  runtime.scanObject
     0.48s  1.95% 58.37%      0.79s  3.21%  runtime.tryDeferToSpanScan
     0.38s  1.54% 59.92%      1.86s  7.56%  runtime.mapaccess1_faststr
     0.37s  1.50% 61.42%      0.57s  2.32%  example.com/isureal-company/isureal.postLess
     0.37s  1.50% 62.93%      0.37s  1.50%  memeqbody
     0.28s  1.14% 64.07%      0.28s  1.14%  sync/atomic.(*Int32).Add (inline)
     0.26s  1.06% 65.12%      0.34s  1.38%  database/sql.convertAssignRows
     0.23s  0.93% 66.06%      0.23s  0.93%  runtime.nextFreeFast (inline)
     0.20s  0.81% 66.87%      0.20s  0.81%  runtime.futex
     0.20s  0.81% 67.68%      0.87s  3.54%  runtime.mallocgcSmallScanNoHeader
     0.19s  0.77% 68.46%      4.87s 19.80%  crypto/internal/fips140/sha256.(*Digest).Write
     0.18s  0.73% 69.19%      4.87s 19.80%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.18s  0.73% 69.92%      0.21s  0.85%  runtime.(*mspan).writeHeapBitsSmall
     0.18s  0.73% 70.65%      0.18s  0.73%  runtime.spanClass.sizeclass (inline)
     0.17s  0.69% 71.34%      1.68s  6.83%  runtime.mallocgc
     0.16s  0.65% 71.99%      5.50s 22.36%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.15s  0.61% 72.60%      0.19s  0.77%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.15s  0.61% 73.21%      0.28s  1.14%  internal/runtime/maps.(*Iter).Next
     0.14s  0.57% 73.78%      5.18s 21.06%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.14s  0.57% 74.35%      1.20s  4.88%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.14s  0.57% 74.92%      0.41s  1.67%  runtime.mallocgcSmallNoscan
     0.13s  0.53% 75.45%      0.16s  0.65%  example.com/isureal-company/isureal.rmLess
     0.12s  0.49% 75.93%      1.53s  6.22%  runtime.mapassign_faststr
     0.11s  0.45% 76.38%      0.13s  0.53%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
     0.10s  0.41% 76.79%      2.23s  9.07%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.41% 77.20%      0.28s  1.14%  runtime.mapaccess2_faststr
     0.09s  0.37% 77.56%      0.22s  0.89%  runtime.greyobject
     0.08s  0.33% 77.89%      3.95s 16.06%  example.com/isureal-company/isureal.loadBig
     0.06s  0.24% 78.13%      0.16s  0.65%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.06s  0.24% 78.37%      5.83s 23.70%  example.com/isureal-company/isureal.computeImageSignature
     0.06s  0.24% 78.62%      0.61s  2.48%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
     0.06s  0.24% 78.86%      0.53s  2.15%  github.com/jackc/pgx/v5/pgtype.scanPlanString.Scan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 27.8% / 100% | 7% | 0.8% | 6 | 109 MB | run queue が積もっている |
| isu2 (2 cpu) | 19.1% / 100% | 1% | 0.7% | 5 | 186 MB | run queue が積もっている |
| isu3 (2 cpu) | 22.1% / 100% | 6% | 0.9% | 5 | 103 MB | run queue が積もっている |
| isu4 (2 cpu) | 56.6% / 100% | 50% | 2.8% | 15 | 106 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 42% / nginx 9% |
| isu2 | isureal 34% / nginx 2% |
| isu3 | isureal 40% / nginx 2% |
| isu4 | isureal 90% / postgres 14% |

## Diff from previous run

スコア: 219729 → 220719 (+990)

`53bad13..7534c3b`

```
 conf/isu1/home/isucon/env.sh                    | 2 +-
 conf/isu4/etc/systemd/system/isureal-go.service | 5 ++---
 webapp/go/jpegcheck/huffman.go                  | 9 +++++++++
 3 files changed, 12 insertions(+), 4 deletions(-)
```
