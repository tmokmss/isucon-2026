# run 20261031-143412-10c9074

## Summary

| | |
|---|---|
| score | 135425 |
| status | pass |
| ref | 10c9074 |
| sha | 10c90744e690 |
| 所要時間 | 122s |
| 前回比 | +11186 (9.0%) vs 20261031-143155-1ce8dee |
| ベスト比 | -5910 (-4.2%) vs 20261031-141027-bd0fe45 = 141335 |
| note | DB外し3: 変更の記録を追記 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	45	125	358	946	945
europe	7	19	53	154	312	311
america	6	18	53	149	219	219
```

```
{"id":"18097","status":"done","portalStatus":"finished","passed":true,"score":135425,"raw":135425,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:34:35.937Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:34:38.928Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:34:41.928Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:34:44.944Z level=INFO msg="asia 2 日目: 45 人"
time=2026-10-31T05:34:45.762Z level=INFO msg="user moved" user=keulne5164 from=asia to=america
time=2026-10-31T05:34:47.199Z level=INFO msg="user moved" user=keulne5164 from=america to=europe
time=2026-10-31T05:34:47.940Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T05:34:50.930Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T05:34:53.951Z level=INFO msg="asia 3 日目: 125 人"
time=2026-10-31T05:34:56.929Z level=INFO msg="europe 3 日目: 53 人"
time=2026-10-31T05:34:59.929Z level=INFO msg="america 3 日目: 53 人"
time=2026-10-31T05:35:00.399Z level=INFO msg="user moved" user=keulne5164 from=europe to=asia
time=2026-10-31T05:35:02.973Z level=INFO msg="asia 4 日目: 358 人"
time=2026-10-31T05:35:03.502Z level=INFO msg="user moved" user=toramia-6025 from=europe to=asia
time=2026-10-31T05:35:05.735Z level=INFO msg="user moved" user=keulne5164 from=asia to=america
time=2026-10-31T05:35:05.821Z level=INFO msg="user moved" user=tetoreva-6289 from=america to=europe
time=2026-10-31T05:35:06.028Z level=INFO msg="europe 4 日目: 154 人"
time=2026-10-31T05:35:06.754Z level=INFO msg="user moved" user=toramia-6025 from=asia to=europe
time=2026-10-31T05:35:06.754Z level=INFO msg="user moved" user=hiriis-7309 from=asia to=europe
time=2026-10-31T05:35:09.100Z level=INFO msg="user moved" user=rane-6965 from=asia to=europe
time=2026-10-31T05:35:09.236Z level=INFO msg="america 4 日目: 149 人"
time=2026-10-31T05:35:10.006Z level=INFO msg="user moved" user=mavinfu.7193 from=asia to=europe
time=2026-10-31T05:35:10.932Z level=INFO msg="user moved" user=somia.7454 from=asia to=europe
time=2026-10-31T05:35:11.952Z level=INFO msg="asia 5 日目: 946 人"
time=2026-10-31T05:35:14.496Z level=INFO msg="user moved" user=maxni.8559 from=europe to=america
time=2026-10-31T05:35:15.271Z level=INFO msg="europe 5 日目: 312 人"
time=2026-10-31T05:35:18.764Z level=INFO msg="america 5 日目: 219 人"
time=2026-10-31T05:35:33.462Z level=INFO msg="america 最終日（グランプリ）: 219 人"
time=2026-10-31T05:35:33.523Z level=INFO msg="asia 最終日（グランプリ）: 945 人"
time=2026-10-31T05:35:33.613Z level=INFO msg="europe 最終日（グランプリ）: 311 人"
time=2026-10-31T05:35:35.960Z level=INFO msg="グランプリ: アジア 61 人, ヨーロッパ 7 人, アメリカ 27 人"
{
  "pass": true,
  "score": 135425,
  "messages": [
    "miaan.5104: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
52.5%	1547	1	0	GET	/api/notifications	19506.196	12.609	21.560	21.568	172
9.6%	2161	367	0	POST	/api/posts	3546.836	1.641	4.218	5.328	542
6.8%	16836	347	0	GET	^/api/images/[^/]+	2542.009	0.151	1.301	4.364	115692
4.8%	5298	70	0	POST	/api/friend-requests	1785.260	0.337	1.609	2.095	114
4.7%	11986	0	0	GET	^/api/feeds/[^/]+	1727.507	0.144	1.231	1.598	3291
2.9%	6143	738	0	POST	/api/auth/login	1082.329	0.176	1.257	1.883	217
2.9%	3767	0	0	GET	/api/memories	1082.204	0.287	2.543	5.435	13976
2.7%	646	121	0	PUT	/api/me/profile-image	1006.533	1.558	5.476	5.620	4
2.2%	2527	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	806.009	0.319	1.529	2.095	116
1.8%	597	17	0	POST	^/api/posts/[^/]+/realmojis$	661.457	1.108	4.643	5.703	465
1.0%	2545	0	0	GET	/api/users	357.514	0.140	0.775	1.563	242
0.8%	3328	0	0	GET	/api/posting-window/current	302.969	0.091	0.541	1.233	237
0.8%	1983	61	0	GET	^/api/posts/[^/]+$	300.063	0.151	1.346	1.563	616
0.8%	967	1	0	POST	/api/auth/register	297.200	0.307	0.936	1.531	227
0.7%	339	0	0	GET	/api/me/export	259.104	0.764	4.993	9.652	1779751
0.7%	1535	8	0	GET	^/api/posts/me/current$	246.220	0.160	1.027	1.563	651
0.5%	593	0	0	PUT	^/api/follows/[^/]+$	199.167	0.336	1.127	1.715	0
0.5%	1512	0	0	GET	^/api/relationships/[^/]+	180.100	0.119	1.070	1.563	1025
0.5%	495	0	0	POST	^/api/posts/[^/]+/comments$	170.100	0.344	1.741	1.934	334
0.4%	454	0	0	PATCH	/api/me	153.175	0.337	1.121	1.532	272
0.4%	934	0	0	POST	/api/auth/logout	151.116	0.162	1.257	1.770	0
0.3%	867	0	0	GET	^/api/memories/[^/]+	125.778	0.145	1.226	1.960	651
0.3%	545	3	0	GET	/api/me	123.211	0.226	0.954	1.236	300
0.3%	952	0	0	GET	^/api/users/[^/]+$	117.195	0.123	1.109	1.563	206
0.2%	143	0	0	DELETE	^/api/realmojis/[^/]+$	72.898	0.510	1.948	2.158	0
0.2%	69	0	0	DELETE	^/api/posts/[^/]+$	72.301	1.048	3.362	3.362	0
0.2%	171	4	0	PATCH	^/api/posts/[^/]+$	71.587	0.419	1.775	1.909	633
0.2%	782	0	0	GET	^/api/posts/[^/]+/realmojis$	57.175	0.073	0.671	1.563	1388
0.1%	598	0	0	GET	^/api/posts/[^/]+/comments$	54.979	0.092	1.233	1.563	2016
0.1%	113	0	0	DELETE	^/api/comments/[^/]+$	34.002	0.301	1.600	1.610	0
…（上位 30 / 全 48 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.7%	568	3	0	GET	/api/notifications	5385.984	9.482	17.491	18.186	175
14.4%	2983	19	0	GET	^/api/images/[^/]+	1824.020	0.611	3.200	3.886	181681
7.2%	756	134	0	POST	/api/posts	914.657	1.210	6.542	7.180	537
6.0%	1637	16	0	POST	/api/friend-requests	755.678	0.462	1.800	2.391	114
5.4%	2421	0	0	GET	^/api/feeds/[^/]+	683.365	0.282	1.284	1.739	2864
3.6%	927	0	0	POST	/api/auth/login	460.114	0.496	1.877	1.916	246
2.9%	176	0	0	PUT	/api/me/profile-image	363.500	2.065	4.198	4.243	0
2.5%	1147	4	0	GET	/api/posting-window/current	313.724	0.274	1.169	1.400	238
2.4%	763	1	0	GET	/api/memories	305.038	0.400	2.030	4.535	11999
2.2%	818	0	0	GET	/api/users	283.207	0.346	1.291	1.695	243
1.7%	459	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	217.344	0.474	1.701	1.804	116
1.6%	306	0	0	POST	/api/auth/register	206.397	0.675	1.816	1.910	229
1.4%	544	17	0	GET	^/api/posts/me/current$	175.789	0.323	1.277	1.475	632
0.8%	203	0	0	PUT	^/api/follows/[^/]+$	107.075	0.527	1.794	1.843	0
0.7%	334	0	0	GET	^/api/relationships/[^/]+	94.227	0.282	1.153	1.467	955
0.7%	187	7	0	GET	/api/me	91.376	0.489	1.632	1.690	293
0.6%	144	0	0	PATCH	/api/me	73.707	0.512	1.872	1.889	279
0.5%	290	32	0	GET	^/api/posts/[^/]+$	62.111	0.214	1.072	1.183	560
0.4%	134	0	0	POST	/api/auth/logout	53.950	0.403	1.804	1.860	0
0.3%	116	0	0	GET	^/api/memories/[^/]+	43.568	0.376	1.949	2.859	957
0.3%	49	0	0	GET	/api/me/export	41.245	0.842	6.324	6.324	2313487
0.3%	166	0	0	GET	^/api/users/[^/]+$	35.546	0.214	1.310	1.400	202
0.3%	38	5	0	POST	^/api/posts/[^/]+/realmojis$	33.759	0.888	3.436	3.436	423
0.2%	29	0	0	DELETE	^/api/posts/[^/]+$	24.672	0.851	2.226	2.226	0
0.2%	48	2	0	PATCH	^/api/posts/[^/]+$	24.474	0.510	2.228	2.228	611
0.1%	62	0	0	POST	^/api/posts/[^/]+/comments$	16.041	0.259	1.467	1.467	332
0.1%	69	0	0	GET	^/api/posts/[^/]+/realmojis$	10.052	0.146	1.070	1.070	1246
0.1%	69	0	0	GET	^/api/posts/[^/]+/comments$	8.187	0.119	1.475	1.475	1607
0.1%	8	0	0	DELETE	^/api/realmojis/[^/]+$	6.739	0.842	1.805	1.805	0
0.0%	10	0	0	PUT	^/api/blocks/[^/]+$	3.141	0.314	1.440	1.440	0
…（上位 30 / 全 41 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.6%	424	1	0	GET	/api/notifications	3798.556	8.959	14.385	15.012	179
16.9%	2826	3	0	GET	^/api/images/[^/]+	1537.612	0.544	2.925	3.329	154271
7.4%	591	58	0	POST	/api/posts	677.014	1.146	3.969	4.768	590
5.8%	1184	9	0	POST	/api/friend-requests	525.159	0.444	1.942	2.367	114
5.4%	2000	0	0	GET	^/api/feeds/[^/]+	494.698	0.247	1.171	1.540	3564
3.8%	850	0	0	POST	/api/auth/login	343.487	0.404	2.321	2.412	250
2.8%	846	4	0	GET	/api/posting-window/current	254.278	0.301	1.165	1.390	238
2.7%	708	1	0	GET	/api/memories	247.379	0.349	2.588	5.662	16442
1.9%	547	0	0	GET	/api/users	173.553	0.317	1.163	1.641	247
1.7%	121	0	0	PUT	/api/me/profile-image	157.733	1.304	3.529	3.613	0
1.5%	354	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	140.538	0.397	1.579	2.296	116
1.3%	217	0	0	POST	/api/auth/register	118.415	0.546	2.243	2.408	229
1.0%	424	12	0	GET	^/api/posts/me/current$	93.549	0.221	1.156	1.390	634
0.8%	153	0	0	PUT	^/api/follows/[^/]+$	75.026	0.490	2.273	2.365	0
0.7%	252	0	0	GET	^/api/relationships/[^/]+	65.669	0.261	1.091	1.385	1062
0.6%	288	26	0	GET	^/api/posts/[^/]+$	57.705	0.200	1.369	1.388	568
0.5%	100	0	0	PATCH	/api/me	44.757	0.448	2.271	2.275	284
0.5%	112	0	0	POST	/api/auth/logout	42.765	0.382	2.271	2.273	0
0.5%	40	6	0	POST	^/api/posts/[^/]+/realmojis$	41.230	1.031	3.555	3.555	407
0.4%	127	2	0	GET	/api/me	36.663	0.289	1.325	1.426	299
0.4%	57	0	0	GET	/api/me/export	36.208	0.635	6.094	6.094	2888962
0.4%	122	0	0	GET	^/api/memories/[^/]+	34.603	0.284	1.723	2.009	775
0.3%	69	1	0	POST	^/api/posts/[^/]+/comments$	27.466	0.398	1.187	1.187	333
0.3%	132	0	0	GET	^/api/users/[^/]+$	23.133	0.175	0.910	1.174	207
0.2%	49	5	0	PATCH	^/api/posts/[^/]+$	22.493	0.459	1.599	1.599	576
0.1%	11	0	0	DELETE	^/api/posts/[^/]+$	12.156	1.105	1.541	1.541	0
0.1%	63	0	0	GET	^/api/posts/[^/]+/comments$	11.142	0.177	1.388	1.388	1647
0.1%	60	2	0	GET	^/api/posts/[^/]+/realmojis$	9.247	0.154	1.131	1.131	1114
0.1%	15	0	0	DELETE	^/api/comments/[^/]+$	7.498	0.500	1.468	1.468	0
0.1%	7	0	0	DELETE	^/api/realmojis/[^/]+$	5.255	0.751	1.824	1.824	0
…（上位 30 / 全 38 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 14192 クエリ / 48 種類 / 合計 143.154s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	53.6%	357	76.785	215.08	1752.92	1137.9	116.3	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.9%	3	21.304	7101.46	9564.31	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.6%	12161	16.535	1.36	9.01	9.5	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	11.2%	3	16.086	5361.99	7185.82	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.5%	3	7.871	2623.81	3546.39	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	1.4%	5	1.973	394.53	1249.18	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	5	0.527	105.43	161.71	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.4%	5	0.514	102.83	164.46	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.293	293.43	293.43	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.245	244.68	244.68	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.2%	1	0.235	234.69	234.69	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.158	157.82	157.82	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.125	125.31	125.31	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.083	83.22	83.22	2493.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.1%	1	0.073	72.66	72.66	1554.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
Build ID: a35fcfeea876c15cc6ffae5bdc1c212a0a191ef1
Type: cpu
Time: 2026-10-31 14:34:18 JST
Duration: 60.19s, Total samples = 48.26s (80.18%)
Showing nodes accounting for 34.65s, 71.80% of 48.26s total
Dropped 1193 nodes (cum <= 0.24s)
Showing top 40 nodes out of 277
      flat  flat%   sum%        cum   cum%
     7.95s 16.47% 16.47%      7.95s 16.47%  crypto/internal/fips140/sha256.blockAVX2
     5.55s 11.50% 27.97%      5.55s 11.50%  internal/runtime/syscall/linux.Syscall6
     1.91s  3.96% 31.93%      1.91s  3.96%  runtime.memmove
     1.67s  3.46% 35.39%      6.66s 13.80%  runtime.scanObject
     1.65s  3.42% 38.81%      1.65s  3.42%  runtime.memclrNoHeapPointers
     1.60s  3.32% 42.13%      2.26s  4.68%  runtime.findObject
     1.54s  3.19% 45.32%      1.54s  3.19%  aeshashbody
     1.52s  3.15% 48.47%      1.52s  3.15%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.49s  3.09% 51.55%      5.81s 12.04%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.33s  2.76% 54.31%      2.61s  5.41%  runtime.tryDeferToSpanScan
     0.95s  1.97% 56.28%      0.95s  1.97%  runtime.spanClass.sizeclass (inline)
     0.73s  1.51% 57.79%      2.48s  5.14%  runtime.mapaccess1_faststr
     0.48s  0.99% 58.79%      0.48s  0.99%  memeqbody
     0.45s  0.93% 59.72%      8.94s 18.52%  crypto/internal/fips140/sha256.(*Digest).Write
     0.40s  0.83% 60.55%      7.98s 16.54%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.40s  0.83% 61.38%      0.40s  0.83%  runtime.nextFreeFast (inline)
     0.33s  0.68% 62.06%      0.33s  0.68%  runtime.futex
     0.32s  0.66% 62.72%      0.68s  1.41%  database/sql.convertAssignRows
     0.31s  0.64% 63.37%      0.52s  1.08%  example.com/isureal-company/isureal.postLess
     0.30s  0.62% 63.99%      8.43s 17.47%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.29s   0.6% 64.59%      1.25s  2.59%  runtime.mallocgcSmallScanNoHeader
     0.27s  0.56% 65.15%      0.27s  0.56%  indexbody
     0.26s  0.54% 65.69%      2.19s  4.54%  runtime.mapassign_faststr
     0.26s  0.54% 66.22%      0.30s  0.62%  runtime.spanOf (inline)
     0.24s   0.5% 66.72%      0.47s  0.97%  compress/flate.(*compressor).deflate
     0.24s   0.5% 67.22%      0.26s  0.54%  encoding/json.appendString[go.shape.string]
     0.21s  0.44% 67.65%      8.16s 16.91%  crypto/internal/fips140/sha256.block
     0.20s  0.41% 68.07%      8.92s 18.48%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.20s  0.41% 68.48%      5.01s 10.38%  runtime.mallocgc
     0.19s  0.39% 68.88%      0.60s  1.24%  runtime.greyobject
     0.18s  0.37% 69.25%      0.28s  0.58%  runtime.(*mspan).markBitsForIndex (inline)
     0.18s  0.37% 69.62%      0.60s  1.24%  runtime.mapaccess2_faststr
     0.16s  0.33% 69.95%      0.47s  0.97%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.15s  0.31% 70.27%      1.26s  2.61%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.14s  0.29% 70.56%      0.42s  0.87%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.13s  0.27% 70.82%      0.63s  1.31%  runtime.wbBufFlush1
     0.12s  0.25% 71.07%      1.05s  2.18%  runtime.mallocgcSmallNoscan
     0.12s  0.25% 71.32%      0.43s  0.89%  runtime.scanObjectsSmall
     0.12s  0.25% 71.57%      0.74s  1.53%  runtime.scanSpan
     0.11s  0.23% 71.80%      0.78s  1.62%  encoding/json.structEncoder.encode
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 44.2% / 100% | 28% | 3.3% | 7 | 100 MB | run queue が積もっている |
| isu2 (2 cpu) | 25.7% / 99% | 2% | 0.6% | 4 | 435 MB |  |
| isu3 (2 cpu) | 27.6% / 99% | 2% | 1.3% | 5 | 610 MB | run queue が積もっている |
| isu4 (2 cpu) | 33.6% / 98% | 6% | 18.1% | 10 | 102 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 67% / nginx 11% |
| isu2 | isureal 49% / nginx 1% |
| isu3 | isureal 50% / nginx 2% |
| isu4 | isureal 15% / postgres 15% |

## Diff from previous run

スコア: 124239 → 135425 (+11186)

`1ce8dee..10c9074`

```
 webapp/go/admin_api.go  |   4 +-
 webapp/go/initialize.go |   4 +-
 webapp/go/persist.go    | 180 +++++++++++++++++++++++++++++++++++++++++-------
 webapp/go/repl.go       |   2 +-
 webapp/go/server.go     |   2 +-
 webapp/go/store.go      |   4 +-
 6 files changed, 163 insertions(+), 33 deletions(-)
```
