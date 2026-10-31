# run 20261031-171843-1662a5d

## Summary

| | |
|---|---|
| score | 219963 |
| status | pass |
| ref | sse-read-timeout |
| sha | 1662a5d5b0f1 |
| 所要時間 | 136s |
| 前回比 | +28955 (15.2%) vs 20261031-171626-1662a5d |
| ベスト比 | -12147 (-5.2%) vs 20261031-164055-bdf05cd = 232110 |
| note | SSE proxy_read_timeout 20s rerun |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	54	159	412	731	735
europe	7	22	63	182	523	523
america	6	19	58	172	472	471
```

```
{"id":"33494","status":"done","portalStatus":"finished","passed":true,"score":219963,"raw":219963,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:19:08.103Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:19:11.099Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:19:14.091Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:19:17.107Z level=INFO msg="asia 2 日目: 54 人"
time=2026-10-31T08:19:19.575Z level=INFO msg="user moved" user=wanu.5282 from=asia to=europe
time=2026-10-31T08:19:19.703Z level=INFO msg="user moved" user=vinis_5302 from=asia to=america
time=2026-10-31T08:19:20.091Z level=INFO msg="europe 2 日目: 22 人"
time=2026-10-31T08:19:20.171Z level=INFO msg="user moved" user=wanu.5282 from=europe to=asia
time=2026-10-31T08:19:20.334Z level=INFO msg="user moved" user=vinis_5302 from=america to=europe
time=2026-10-31T08:19:23.091Z level=INFO msg="america 2 日目: 19 人"
time=2026-10-31T08:19:26.104Z level=INFO msg="asia 3 日目: 159 人"
time=2026-10-31T08:19:28.329Z level=INFO msg="user moved" user=wanu.5282 from=asia to=europe
time=2026-10-31T08:19:29.094Z level=INFO msg="europe 3 日目: 63 人"
time=2026-10-31T08:19:31.245Z level=INFO msg="user moved" user=vinis_5302 from=europe to=america
time=2026-10-31T08:19:32.095Z level=INFO msg="america 3 日目: 58 人"
time=2026-10-31T08:19:32.851Z level=INFO msg="user moved" user=chini.6306 from=asia to=america
time=2026-10-31T08:19:33.179Z level=INFO msg="user moved" user=ruya_5913 from=asia to=america
time=2026-10-31T08:19:35.105Z level=INFO msg="asia 4 日目: 412 人"
time=2026-10-31T08:19:38.098Z level=INFO msg="europe 4 日目: 182 人"
time=2026-10-31T08:19:39.293Z level=INFO msg="user moved" user=wanu.5282 from=europe to=asia
time=2026-10-31T08:19:40.512Z level=INFO msg="user moved" user=hafushi7570 from=asia to=europe
time=2026-10-31T08:19:41.156Z level=INFO msg="america 4 日目: 172 人"
time=2026-10-31T08:19:42.222Z level=INFO msg="user moved" user=mamarben-8430 from=asia to=europe
time=2026-10-31T08:19:43.183Z level=INFO msg="user moved" user=kirumu.9565 from=america to=asia
time=2026-10-31T08:19:43.452Z level=INFO msg="user moved" user=ronaben.9588 from=europe to=america
time=2026-10-31T08:19:43.774Z level=INFO msg="user moved" user=hori.7910 from=asia to=america
time=2026-10-31T08:19:44.112Z level=INFO msg="asia 5 日目: 731 人"
time=2026-10-31T08:19:44.387Z level=INFO msg="user moved" user=todan9664 from=america to=asia
time=2026-10-31T08:19:44.481Z level=INFO msg="user moved" user=vinis_5302 from=america to=europe
time=2026-10-31T08:19:44.494Z level=INFO msg="user moved" user=koketor_9859 from=america to=asia
time=2026-10-31T08:19:47.138Z level=INFO msg="europe 5 日目: 523 人"
time=2026-10-31T08:19:48.547Z level=INFO msg="user moved" user=kuor.10215 from=america to=europe
time=2026-10-31T08:19:49.215Z level=INFO msg="user moved" user=maki.9770 from=america to=asia
time=2026-10-31T08:19:49.830Z level=INFO msg="user moved" user=chini.6306 from=america to=asia
time=2026-10-31T08:19:50.284Z level=INFO msg="user moved" user=orma_10123 from=america to=asia
time=2026-10-31T08:19:50.455Z level=INFO msg="america 5 日目: 472 人"
time=2026-10-31T08:19:59.970Z level=INFO msg="europe 最終日（グランプリ）: 523 人"
time=2026-10-31T08:20:01.656Z level=INFO msg="asia 最終日（グランプリ）: 735 人"
time=2026-10-31T08:20:01.850Z level=INFO msg="america 最終日（グランプリ）: 471 人"
time=2026-10-31T08:20:08.173Z level=INFO msg="グランプリ: アジア 239 人, ヨーロッパ 407 人, アメリカ 185 人"
{
  "pass": true,
  "score": 219963,
  "messages": [
    "seoyeon_lee_1482: ログアウトできません: POST /api/auth/logout: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/logout\": EOF",
    "fan_yang_1897: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
67.4%	2135	0	0	GET	/api/notifications	29667.032	13.896	20.549	20.549	115
6.6%	11254	0	0	GET	^/api/feeds/[^/]+	2922.391	0.260	1.732	2.840	5231
4.9%	2720	0	0	GET	/api/memories	2168.665	0.797	6.374	10.012	11393
3.7%	2087	98	0	POST	/api/posts	1647.169	0.789	5.124	5.596	623
2.6%	2873	0	0	GET	/api/posting-window/current	1160.860	0.404	1.788	2.421	237
1.9%	11601	277	0	GET	^/api/images/[^/]+	824.585	0.071	0.931	2.402	175860
1.8%	2013	0	0	GET	/api/users	791.842	0.393	1.630	2.737	234
1.7%	1504	1	0	GET	^/api/posts/me/current$	765.934	0.509	1.638	2.346	652
1.4%	193	0	0	GET	/api/me/export	635.493	3.293	28.065	30.290	1911353
1.2%	551	128	0	PUT	/api/me/profile-image	529.039	0.960	13.236	16.825	4
0.8%	4369	58	0	POST	/api/friend-requests	338.742	0.078	0.585	1.742	114
0.7%	972	16	0	GET	^/api/posts/[^/]+/realmojis$	311.579	0.321	2.252	3.880	2080
0.7%	2005	86	0	GET	^/api/posts/[^/]+$	303.616	0.151	1.176	2.737	605
0.7%	433	3	0	GET	/api/me	291.863	0.674	1.996	2.369	300
0.6%	1323	0	0	GET	^/api/relationships/[^/]+	281.567	0.213	1.287	3.152	1337
0.6%	3214	8	0	POST	^/api/friend-requests/[^/]+/[^/]+$	277.208	0.086	0.576	2.007	116
0.5%	764	5	0	POST	^/api/posts/[^/]+/realmojis$	230.168	0.301	1.242	3.340	475
0.4%	464	0	0	GET	^/api/memories/[^/]+	166.442	0.359	2.021	2.562	667
0.3%	3703	456	0	POST	/api/auth/login	146.422	0.040	0.616	2.007	216
0.3%	980	1	0	GET	^/api/users/[^/]+$	136.956	0.140	0.995	2.737	196
0.3%	799	9	0	GET	^/api/posts/[^/]+/comments$	116.340	0.146	0.905	2.737	2261
0.2%	736	1	0	POST	^/api/posts/[^/]+/comments$	88.592	0.120	0.856	1.742	343
0.1%	122	0	0	DELETE	^/api/realmojis/[^/]+$	44.152	0.362	2.737	2.737	0
0.1%	67	0	0	DELETE	^/api/posts/[^/]+$	30.823	0.460	1.264	1.264	0
0.1%	381	0	0	PATCH	/api/me	24.847	0.065	0.655	0.676	276
0.1%	528	0	0	PUT	^/api/follows/[^/]+$	23.235	0.044	0.589	0.738	0
0.0%	531	0	0	POST	/api/auth/logout	21.877	0.041	0.545	0.655	0
0.0%	157	2	0	PATCH	^/api/posts/[^/]+$	18.208	0.116	0.672	1.742	652
0.0%	1	0	0	POST	/api/initialize	16.742	16.742	16.742	16.742	90
0.0%	752	1	0	POST	/api/auth/register	14.507	0.019	0.340	0.995	227
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
71.8%	1334	1	0	GET	/api/notifications	16738.729	12.548	20.001	20.150	107
4.4%	8079	0	0	GET	^/api/feeds/[^/]+	1018.349	0.126	1.110	2.085	4637
3.8%	1322	33	0	POST	/api/posts	880.593	0.666	5.276	5.540	635
3.5%	2954	34	0	POST	/api/friend-requests	824.256	0.279	0.985	1.152	114
3.4%	6691	48	0	GET	^/api/images/[^/]+	792.225	0.118	1.343	4.823	208117
1.9%	1898	0	0	GET	/api/posting-window/current	454.518	0.239	1.677	1.955	239
1.8%	1966	0	0	POST	/api/auth/login	410.772	0.209	0.901	1.813	247
1.0%	1713	0	0	GET	/api/memories	238.969	0.140	1.533	2.031	13667
0.9%	549	7	0	POST	^/api/posts/[^/]+/realmojis$	218.223	0.397	1.423	2.839	462
0.8%	1391	0	0	GET	/api/users	185.706	0.134	0.976	1.181	240
0.7%	1174	0	0	GET	^/api/posts/me/current$	153.972	0.131	1.054	1.745	652
0.6%	673	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	149.777	0.223	0.966	1.098	116
0.6%	519	0	0	POST	/api/auth/register	147.922	0.285	0.779	0.826	229
0.6%	1382	41	0	GET	^/api/posts/[^/]+$	141.314	0.102	0.884	1.365	618
0.6%	955	0	0	GET	^/api/relationships/[^/]+	139.764	0.146	1.065	1.509	517
0.5%	292	0	0	PUT	/api/me/profile-image	115.265	0.395	1.053	1.083	0
0.4%	384	0	0	PUT	^/api/follows/[^/]+$	99.178	0.258	0.927	1.051	0
0.4%	378	1	0	POST	^/api/posts/[^/]+/comments$	84.736	0.224	1.037	1.095	347
0.3%	656	0	0	GET	^/api/posts/[^/]+/realmojis$	62.878	0.096	0.729	1.054	3201
0.3%	239	0	0	PATCH	/api/me	61.835	0.259	0.954	1.051	273
0.3%	651	0	0	GET	^/api/users/[^/]+$	61.465	0.094	0.865	1.088	205
0.2%	278	0	0	POST	/api/auth/logout	52.930	0.190	0.825	1.261	0
0.2%	141	0	0	PATCH	^/api/posts/[^/]+$	50.258	0.356	1.080	1.152	654
0.2%	478	2	0	GET	^/api/posts/[^/]+/comments$	45.284	0.095	0.662	1.088	3961
0.2%	62	0	0	DELETE	^/api/posts/[^/]+$	42.944	0.693	1.596	1.596	0
0.2%	299	2	0	GET	/api/me	40.325	0.135	0.379	0.505	302
0.2%	296	0	0	GET	^/api/memories/[^/]+	35.348	0.119	1.476	1.698	759
0.1%	86	0	0	DELETE	^/api/realmojis/[^/]+$	30.693	0.357	1.299	1.299	0
0.1%	132	0	0	GET	/api/me/export	26.743	0.203	3.316	4.053	2266203
0.1%	88	0	0	DELETE	^/api/comments/[^/]+$	14.800	0.168	0.940	0.940	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
71.1%	1215	0	0	GET	/api/notifications	14041.806	11.557	20.001	20.103	108
4.6%	1200	84	0	POST	/api/posts	906.533	0.755	6.152	6.219	613
4.2%	5516	0	0	GET	^/api/feeds/[^/]+	820.059	0.149	1.266	2.348	3421
3.9%	6287	34	0	GET	^/api/images/[^/]+	771.617	0.123	1.113	3.729	164342
3.5%	2608	30	0	POST	/api/friend-requests	697.417	0.267	0.850	1.625	114
2.1%	1801	0	0	POST	/api/auth/login	419.095	0.233	0.814	1.710	250
1.1%	1697	0	0	GET	/api/posting-window/current	215.591	0.127	0.934	1.002	239
1.1%	1407	0	0	GET	/api/memories	210.257	0.149	1.614	2.349	17166
0.9%	1303	0	0	GET	/api/users	184.850	0.142	0.960	1.813	241
0.8%	474	0	0	POST	/api/auth/register	160.212	0.338	0.729	0.820	230
0.7%	892	2	0	GET	^/api/posts/me/current$	146.333	0.164	1.742	2.024	658
0.7%	622	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	139.759	0.225	0.805	1.381	116
0.7%	268	0	0	PUT	/api/me/profile-image	133.324	0.497	1.436	1.961	0
0.6%	273	8	0	POST	^/api/posts/[^/]+/realmojis$	116.735	0.428	1.285	3.851	454
0.5%	335	0	0	PUT	^/api/follows/[^/]+$	107.180	0.320	0.791	0.811	0
0.5%	639	0	0	GET	^/api/relationships/[^/]+	96.424	0.151	1.635	1.903	739
0.5%	787	56	0	GET	^/api/posts/[^/]+$	93.092	0.118	0.679	1.347	592
0.4%	250	0	0	PATCH	/api/me	82.571	0.330	0.810	0.914	278
0.4%	242	1	0	POST	^/api/posts/[^/]+/comments$	69.560	0.287	1.094	1.560	351
0.3%	273	0	0	POST	/api/auth/logout	56.403	0.207	1.118	1.181	0
0.2%	382	0	0	GET	^/api/users/[^/]+$	43.841	0.115	0.680	0.808	204
0.2%	288	0	0	GET	^/api/posts/[^/]+/realmojis$	38.367	0.133	0.978	1.147	3088
0.2%	90	0	0	PATCH	^/api/posts/[^/]+$	34.697	0.386	1.721	1.721	661
0.2%	115	0	0	GET	/api/me/export	30.975	0.269	3.497	3.679	2928780
0.1%	251	0	0	GET	^/api/posts/[^/]+/comments$	28.656	0.114	0.683	1.070	3287
0.1%	254	0	0	GET	^/api/memories/[^/]+	28.546	0.112	0.997	1.253	1113
0.1%	275	1	0	GET	/api/me	27.543	0.100	0.388	0.468	304
0.1%	49	0	0	DELETE	^/api/realmojis/[^/]+$	16.607	0.339	0.963	0.963	0
0.1%	34	0	0	DELETE	^/api/posts/[^/]+$	16.349	0.481	2.165	2.165	0
0.1%	47	0	0	DELETE	^/api/comments/[^/]+$	10.328	0.220	0.879	0.879	0
…（上位 30 / 全 36 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 10818 クエリ / 48 種類 / 合計 129.197s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	43.2%	200	55.833	279.16	1606.53	1246.3	126.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	20.1%	4	25.968	6492.06	9547.34	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.6%	4	18.823	4705.68	7206.52	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.9%	6871	11.526	1.68	18.15	10.2	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	7.3%	4	9.449	2362.24	3590.44	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.5%	8	3.285	410.65	749.34	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.6%	8	0.817	102.15	171.35	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.6%	8	0.791	98.87	161.98	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.4%	1	0.488	487.93	487.93	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.3%	1	0.384	384.40	384.40	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.3%	1	0.362	361.56	361.56	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.3%	1	0.347	347.17	347.17	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.2%	1	0.262	261.93	261.93	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.129	128.58	128.58	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
15	0.1%	8	0.111	13.93	21.88	166.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: 0eb83cf1921478aee6a7b30e210bbe2042744bfc
Type: cpu
Time: 2026-10-31 17:18:46 JST
Duration: 60.05s, Total samples = 33.85s (56.37%)
Showing nodes accounting for 25.31s, 74.77% of 33.85s total
Dropped 974 nodes (cum <= 0.17s)
Showing top 40 nodes out of 249
      flat  flat%   sum%        cum   cum%
     4.66s 13.77% 13.77%      4.66s 13.77%  crypto/internal/fips140/sha256.blockAVX2
     3.67s 10.84% 24.61%      3.67s 10.84%  internal/runtime/syscall/linux.Syscall6
     1.73s  5.11% 29.72%      1.73s  5.11%  runtime.memmove
     1.72s  5.08% 34.80%      6.21s 18.35%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.45s  4.28% 39.08%      1.45s  4.28%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.30s  3.84% 42.92%      1.30s  3.84%  aeshashbody
        1s  2.95% 45.88%      1.22s  3.60%  runtime.findObject
     0.97s  2.87% 48.74%      0.97s  2.87%  runtime.memclrNoHeapPointers
     0.90s  2.66% 51.40%      3.82s 11.29%  runtime.scanObject
     0.82s  2.42% 53.83%      1.41s  4.17%  runtime.tryDeferToSpanScan
     0.57s  1.68% 55.51%      2.27s  6.71%  runtime.mapaccess1_faststr
     0.51s  1.51% 57.02%      0.68s  2.01%  example.com/isureal-company/isureal/jpegcheck.(*decoder).skipAC
     0.45s  1.33% 58.35%      0.67s  1.98%  example.com/isureal-company/isureal.postLess
     0.42s  1.24% 59.59%      0.42s  1.24%  memeqbody
     0.42s  1.24% 60.83%      0.42s  1.24%  runtime.spanClass.sizeclass (inline)
     0.37s  1.09% 61.92%      1.04s  3.07%  runtime.mallocgcSmallScanNoHeader
     0.34s  1.00% 62.92%      5.34s 15.78%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.32s  0.95% 63.87%      5.26s 15.54%  crypto/internal/fips140/sha256.(*Digest).Write
     0.30s  0.89% 64.76%      0.30s  0.89%  runtime.nextFreeFast (inline)
     0.28s  0.83% 65.58%      0.43s  1.27%  database/sql.convertAssignRows
     0.28s  0.83% 66.41%      0.28s  0.83%  runtime.futex
     0.27s   0.8% 67.21%      0.27s   0.8%  sync/atomic.(*Int32).Add (inline)
     0.22s  0.65% 67.86%      0.24s  0.71%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.21s  0.62% 68.48%      2.29s  6.77%  runtime.mallocgc
     0.20s  0.59% 69.07%      0.21s  0.62%  example.com/isureal-company/isureal.rmLess
     0.20s  0.59% 69.66%      0.39s  1.15%  runtime.greyobject
     0.20s  0.59% 70.25%      1.60s  4.73%  runtime.mapassign_faststr
     0.19s  0.56% 70.81%      0.98s  2.90%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.16s  0.47% 71.29%      2.62s  7.74%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.14s  0.41% 71.70%      6.01s 17.75%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.14s  0.41% 72.11%      5.59s 16.51%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.14s  0.41% 72.53%      0.37s  1.09%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.14s  0.41% 72.94%      0.46s  1.36%  runtime.mapaccess2_faststr
     0.11s  0.32% 73.26%      4.77s 14.09%  crypto/internal/fips140/sha256.block
     0.11s  0.32% 73.59%      0.63s  1.86%  runtime.mallocgcSmallNoscan
     0.10s   0.3% 73.88%      1.45s  4.28%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.09s  0.27% 74.15%      4.63s 13.68%  example.com/isureal-company/isureal.loadBig
     0.07s  0.21% 74.36%      0.50s  1.48%  database/sql.(*Rows).scanLocked
     0.07s  0.21% 74.56%      0.63s  1.86%  runtime.convTstring
     0.07s  0.21% 74.77%      0.19s  0.56%  runtime.scanObjectsSmall
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 72.1% / 100% | 34% | 30.6% | 7 | 100 MB | CPU が飽和 |
| isu2 (2 cpu) | 23.7% / 100% | 5% | 0.7% | 5 | 105 MB | run queue が積もっている |
| isu3 (2 cpu) | 23.5% / 100% | 7% | 1.0% | 4 | 102 MB |  |
| isu4 (2 cpu) | 56.7% / 100% | 49% | 1.7% | 7 | 103 MB | CPU が飽和 |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 50% / find 13% / nginx 10% / kswapd0 4% / kcompactd0 2% |
| isu2 | isureal 42% / nginx 3% |
| isu3 | isureal 43% / nginx 2% |
| isu4 | isureal 93% / postgres 14% |

## Diff from previous run

スコア: 191008 → 219963 (+28955)

コードの差分は無い（同じ sha 1662a5d5b0f1）。
スコアの違いはブレか、サーバ側の設定変更による。
