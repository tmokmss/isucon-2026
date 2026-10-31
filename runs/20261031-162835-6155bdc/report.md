# run 20261031-162835-6155bdc

## Summary

| | |
|---|---|
| score | 198345 |
| status | pass |
| ref | jpeg-peers-asia2-last |
| sha | 6155bdc01658 |
| 所要時間 | 136s |
| 前回比 | -6930 (-3.4%) vs 20261031-162248-73b6aff |
| ベスト比 | -6930 (-3.4%) vs 20261031-162248-73b6aff = 205275 |
| note | asia1 JPEG offload peers: europe1/america1 only (drop asia2) |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	150	426	792	791
europe	7	19	56	161	428	428
america	6	18	61	178	423	422
```

```
{"id":"28453","status":"done","portalStatus":"finished","passed":true,"score":198345,"raw":198345,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:29:07.832Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:29:10.822Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:29:13.822Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:29:16.837Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T07:29:19.823Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T07:29:22.822Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T07:29:25.834Z level=INFO msg="asia 3 日目: 150 人"
time=2026-10-31T07:29:28.627Z level=INFO msg="user moved" user=kamuyo.5725 from=asia to=europe
time=2026-10-31T07:29:28.822Z level=INFO msg="europe 3 日目: 56 人"
time=2026-10-31T07:29:28.835Z level=INFO msg="user moved" user=shiisna-5909 from=asia to=america
time=2026-10-31T07:29:28.958Z level=INFO msg="user moved" user=fuhote.5730 from=asia to=america
time=2026-10-31T07:29:29.795Z level=INFO msg="user moved" user=shiisna-5909 from=america to=asia
time=2026-10-31T07:29:30.081Z level=INFO msg="user moved" user=seel_6283 from=asia to=america
time=2026-10-31T07:29:30.968Z level=INFO msg="user moved" user=seel_6283 from=america to=asia
time=2026-10-31T07:29:31.853Z level=INFO msg="america 3 日目: 61 人"
time=2026-10-31T07:29:33.399Z level=INFO msg="user moved" user=yaleodan_6456 from=europe to=america
time=2026-10-31T07:29:34.861Z level=INFO msg="asia 4 日目: 426 人"
time=2026-10-31T07:29:37.833Z level=INFO msg="europe 4 日目: 161 人"
time=2026-10-31T07:29:38.774Z level=INFO msg="user moved" user=seel_6283 from=asia to=america
time=2026-10-31T07:29:39.328Z level=INFO msg="user moved" user=tsuri6585 from=america to=asia
time=2026-10-31T07:29:39.554Z level=INFO msg="user moved" user=mechi_6822 from=asia to=america
time=2026-10-31T07:29:39.900Z level=INFO msg="user moved" user=evamoivy-7247 from=asia to=america
time=2026-10-31T07:29:40.844Z level=INFO msg="america 4 日目: 178 人"
time=2026-10-31T07:29:41.102Z level=INFO msg="user moved" user=shiisna-5909 from=asia to=america
time=2026-10-31T07:29:42.483Z level=INFO msg="user moved" user=suleo-8709 from=europe to=america
time=2026-10-31T07:29:43.614Z level=INFO msg="user moved" user=yaleodan_6456 from=america to=europe
time=2026-10-31T07:29:43.866Z level=INFO msg="asia 5 日目: 792 人"
time=2026-10-31T07:29:44.360Z level=INFO msg="user moved" user=bensonu.9412 from=europe to=america
time=2026-10-31T07:29:46.846Z level=INFO msg="europe 5 日目: 428 人"
time=2026-10-31T07:29:49.697Z level=INFO msg="user moved" user=noteel_9929 from=america to=europe
time=2026-10-31T07:29:50.018Z level=INFO msg="america 5 日目: 423 人"
time=2026-10-31T07:29:58.868Z level=INFO msg="europe 最終日（グランプリ）: 428 人"
time=2026-10-31T07:29:58.898Z level=INFO msg="asia 最終日（グランプリ）: 791 人"
time=2026-10-31T07:29:58.987Z level=INFO msg="america 最終日（グランプリ）: 422 人"
time=2026-10-31T07:30:07.923Z level=INFO msg="グランプリ: アジア 315 人, ヨーロッパ 387 人, アメリカ 186 人"
{
  "pass": true,
  "score": 198345,
  "messages": [
    "hesa-5030: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "yota.9433: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tohoke-5083: 友達申請を承認できません: POST /api/friend-requests/3da52371-3062-40fc-abed-7e9f71fc4e10/accept: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests/3da52371-3062-40fc-abed-7e9f71fc4e10/accept\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.9%	1471	1	0	GET	/api/notifications	14737.276	10.019	15.065	15.074	174
13.5%	2236	83	0	POST	/api/posts	4421.086	1.977	6.847	9.305	627
9.4%	9500	0	0	GET	^/api/feeds/[^/]+	3075.574	0.324	1.605	2.401	4298
7.1%	3205	0	0	GET	/api/memories	2325.190	0.725	3.698	7.725	11570
4.8%	3074	0	0	GET	/api/posting-window/current	1563.564	0.509	1.567	1.898	237
3.1%	2259	0	0	GET	/api/users	1018.733	0.451	1.563	1.671	235
2.8%	1725	17	0	GET	^/api/posts/me/current$	931.296	0.540	1.718	2.290	644
2.2%	232	0	0	GET	/api/me/export	723.646	3.119	14.648	17.224	2041002
2.0%	572	128	0	PUT	/api/me/profile-image	666.303	1.165	10.877	19.455	3
1.6%	574	10	0	POST	^/api/posts/[^/]+/realmojis$	537.018	0.936	6.608	6.796	463
1.5%	13162	282	0	GET	^/api/images/[^/]+	507.284	0.039	0.722	2.761	145470
1.0%	1621	78	0	GET	^/api/posts/[^/]+$	332.038	0.205	1.378	1.868	588
1.0%	1182	0	0	GET	^/api/relationships/[^/]+	327.989	0.277	1.463	2.498	1577
0.8%	456	3	0	GET	/api/me	272.563	0.598	1.657	1.818	300
0.7%	692	25	0	GET	^/api/posts/[^/]+/realmojis$	239.570	0.346	1.887	2.135	1543
0.7%	605	0	0	GET	^/api/memories/[^/]+	229.288	0.379	1.903	2.317	595
0.6%	4684	51	0	POST	/api/friend-requests	200.556	0.043	0.630	1.009	114
0.5%	782	0	0	GET	^/api/users/[^/]+$	158.515	0.203	1.472	2.281	189
0.4%	658	16	0	GET	^/api/posts/[^/]+/comments$	147.378	0.224	1.470	1.842	2507
0.4%	3129	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	128.275	0.041	0.624	0.779	116
0.3%	4496	551	0	POST	/api/auth/login	82.322	0.018	0.371	0.925	217
0.2%	700	0	0	POST	^/api/posts/[^/]+/comments$	50.459	0.072	0.728	1.110	335
0.1%	103	0	0	DELETE	^/api/posts/[^/]+$	48.941	0.475	1.032	1.037	0
0.1%	1	0	0	POST	/api/initialize	21.987	21.987	21.987	21.987	90
0.1%	664	0	0	POST	/api/auth/logout	17.507	0.026	0.586	0.775	0
0.0%	79	0	0	DELETE	^/api/realmojis/[^/]+$	15.854	0.201	0.625	0.625	0
0.0%	458	0	0	PUT	^/api/follows/[^/]+$	12.800	0.028	0.565	0.724	0
0.0%	106	0	0	DELETE	^/api/comments/[^/]+$	9.411	0.089	0.660	0.917	0
0.0%	813	1	0	POST	/api/auth/register	7.619	0.009	0.213	0.287	227
0.0%	383	0	0	PATCH	/api/me	7.008	0.018	0.211	0.676	274
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.9%	680	1	0	GET	/api/notifications	5045.879	7.420	11.756	11.757	171
8.4%	6496	0	0	GET	^/api/feeds/[^/]+	923.783	0.142	1.501	2.205	3838
6.4%	5509	25	0	GET	^/api/images/[^/]+	703.101	0.128	1.154	3.592	188100
6.4%	1072	12	0	POST	/api/posts	698.567	0.652	2.447	7.537	641
5.8%	237	0	0	PUT	/api/me/profile-image	632.847	2.670	6.723	6.800	0
5.5%	2458	25	0	POST	/api/friend-requests	601.351	0.245	0.897	1.595	114
3.2%	1626	0	0	POST	/api/auth/login	346.395	0.213	0.812	1.879	248
2.7%	1512	0	0	GET	/api/memories	294.783	0.195	2.204	3.064	12827
1.9%	296	14	0	POST	^/api/posts/[^/]+/realmojis$	206.894	0.699	7.370	7.777	452
1.7%	1538	0	0	GET	/api/posting-window/current	190.500	0.124	0.695	2.200	239
1.7%	969	0	0	GET	^/api/posts/me/current$	184.291	0.190	1.630	2.071	648
1.4%	1184	0	0	GET	/api/users	157.603	0.133	0.837	1.795	238
1.3%	427	0	0	POST	/api/auth/register	138.756	0.325	0.638	0.685	229
1.1%	521	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	125.030	0.240	0.894	1.591	116
0.9%	1037	28	0	GET	^/api/posts/[^/]+$	104.209	0.100	0.729	1.853	605
0.8%	731	0	0	GET	^/api/relationships/[^/]+	88.186	0.121	1.392	1.773	512
0.7%	335	1	0	POST	^/api/posts/[^/]+/comments$	75.222	0.225	1.236	1.541	346
0.6%	301	0	0	PUT	^/api/follows/[^/]+$	70.744	0.235	0.774	1.414	0
0.5%	287	0	0	POST	/api/auth/logout	51.696	0.180	0.746	0.920	0
0.5%	518	0	0	GET	^/api/users/[^/]+$	50.561	0.098	0.668	1.339	190
0.4%	194	0	0	PATCH	/api/me	48.330	0.249	0.831	0.834	278
0.4%	445	3	0	GET	^/api/posts/[^/]+/realmojis$	42.468	0.095	0.672	0.777	1904
0.3%	419	0	0	GET	^/api/posts/[^/]+/comments$	36.322	0.087	0.294	0.666	3098
0.3%	96	0	0	GET	/api/me/export	31.688	0.330	4.532	4.532	2144172
0.3%	246	2	0	GET	/api/me	30.595	0.124	0.692	0.755	301
0.2%	99	1	0	PATCH	^/api/posts/[^/]+$	25.752	0.260	1.223	1.223	642
0.2%	226	0	0	GET	^/api/memories/[^/]+	25.199	0.111	0.756	1.365	782
0.2%	35	0	0	DELETE	^/api/posts/[^/]+$	21.469	0.613	2.070	2.070	0
0.1%	35	0	0	DELETE	^/api/realmojis/[^/]+$	13.745	0.393	0.941	0.941	0
0.1%	55	0	0	DELETE	^/api/comments/[^/]+$	10.908	0.198	0.833	0.833	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.9%	711	0	0	GET	/api/notifications	4451.730	6.261	9.366	9.367	167
8.7%	4630	0	0	GET	^/api/feeds/[^/]+	886.055	0.191	2.160	3.122	3491
8.2%	1073	34	0	POST	/api/posts	834.440	0.778	2.889	7.599	633
8.2%	5221	36	0	GET	^/api/images/[^/]+	830.474	0.159	2.164	3.745	168360
5.2%	2279	16	0	POST	/api/friend-requests	523.956	0.230	0.867	1.978	114
3.4%	1542	0	0	POST	/api/auth/login	341.707	0.222	1.067	1.881	249
3.0%	228	0	0	PUT	/api/me/profile-image	300.495	1.318	6.748	6.966	0
2.9%	1285	0	0	GET	/api/memories	289.933	0.226	2.586	3.350	15981
2.2%	1554	0	0	GET	/api/posting-window/current	221.844	0.143	1.131	2.228	240
1.8%	240	26	0	POST	^/api/posts/[^/]+/realmojis$	182.041	0.759	5.474	6.789	416
1.7%	1084	0	0	GET	/api/users	176.044	0.162	0.954	2.507	237
1.5%	417	0	0	POST	/api/auth/register	151.625	0.364	1.236	1.398	230
1.4%	783	1	0	GET	^/api/posts/me/current$	145.785	0.186	1.940	2.586	653
1.2%	571	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	124.196	0.218	0.973	1.948	116
1.0%	624	0	0	GET	^/api/relationships/[^/]+	99.625	0.160	1.397	2.283	709
0.9%	626	32	0	GET	^/api/posts/[^/]+$	91.631	0.146	0.767	2.226	592
0.6%	280	0	0	PUT	^/api/follows/[^/]+$	61.150	0.218	0.876	1.118	0
0.6%	269	0	0	POST	/api/auth/logout	60.422	0.225	1.487	1.913	0
0.5%	232	1	0	POST	^/api/posts/[^/]+/comments$	50.780	0.219	0.596	2.136	349
0.5%	328	0	0	GET	^/api/users/[^/]+$	46.777	0.143	0.806	2.200	189
0.5%	102	0	0	GET	/api/me/export	45.981	0.451	4.583	5.105	2804047
0.4%	204	0	0	PATCH	/api/me	44.770	0.219	0.829	0.857	273
0.4%	239	1	0	GET	^/api/posts/[^/]+/realmojis$	37.749	0.158	1.039	1.653	2056
0.3%	199	0	0	GET	^/api/memories/[^/]+	26.540	0.133	1.309	1.381	1013
0.3%	211	0	0	GET	^/api/posts/[^/]+/comments$	26.381	0.125	0.700	0.954	3712
0.2%	39	0	0	DELETE	^/api/posts/[^/]+$	25.026	0.642	2.463	2.463	0
0.2%	233	1	0	GET	/api/me	24.272	0.104	0.434	0.434	305
0.2%	61	2	0	PATCH	^/api/posts/[^/]+$	15.831	0.260	1.916	1.916	623
0.1%	37	0	0	DELETE	^/api/realmojis/[^/]+$	15.039	0.406	0.851	0.851	0
0.1%	36	0	0	DELETE	^/api/comments/[^/]+$	7.873	0.219	0.571	0.571	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 12432 クエリ / 48 種類 / 合計 144.002s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	46.9%	251	67.563	269.17	859.15	1328.6	135.7	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	17.8%	4	25.694	6423.54	10097.06	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	13.0%	4	18.700	4674.97	7180.76	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	9.5%	8580	13.712	1.60	13.93	10.1	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.5%	4	9.366	2341.62	3416.15	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	3.7%	7	5.354	764.86	3468.22	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	7	0.642	91.72	163.31	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
8	0.4%	7	0.607	86.70	161.69	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.4%	1	0.554	553.72	553.72	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
10	0.3%	1	0.416	415.64	415.64	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.230	230.21	230.21	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.2%	1	0.221	220.52	220.52	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.174	173.55	173.55	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.156	156.34	156.34	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
15	0.1%	7	0.094	13.40	21.79	166.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: 55e616adfe49a98f8b775a891f73d18130d29ef4
Type: cpu
Time: 2026-10-31 16:28:41 JST
Duration: 60s, Total samples = 25.77s (42.95%)
Showing nodes accounting for 20.15s, 78.19% of 25.77s total
Dropped 842 nodes (cum <= 0.13s)
Showing top 40 nodes out of 229
      flat  flat%   sum%        cum   cum%
     4.58s 17.77% 17.77%      4.58s 17.77%  crypto/internal/fips140/sha256.blockAVX2
     2.67s 10.36% 28.13%      2.67s 10.36%  internal/runtime/syscall/linux.Syscall6
     1.83s  7.10% 35.23%      5.89s 22.86%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.29s  5.01% 40.24%      1.29s  5.01%  runtime.memmove
     1.28s  4.97% 45.21%      1.28s  4.97%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.12s  4.35% 49.55%      1.12s  4.35%  aeshashbody
     0.78s  3.03% 52.58%      0.78s  3.03%  runtime.memclrNoHeapPointers
     0.64s  2.48% 55.06%      0.79s  3.07%  runtime.findObject
     0.57s  2.21% 57.28%      2.26s  8.77%  runtime.scanObject
     0.55s  2.13% 59.41%      0.73s  2.83%  runtime.tryDeferToSpanScan
     0.41s  1.59% 61.00%      1.72s  6.67%  runtime.mapaccess1_faststr
     0.35s  1.36% 62.36%      0.57s  2.21%  example.com/isureal-company/isureal.postLess
     0.33s  1.28% 63.64%      0.45s  1.75%  database/sql.convertAssignRows
     0.27s  1.05% 64.69%      0.27s  1.05%  memeqbody
     0.27s  1.05% 65.74%      0.27s  1.05%  runtime.nextFreeFast (inline)
     0.24s  0.93% 66.67%      0.24s  0.93%  runtime.futex
     0.23s  0.89% 67.56%      0.83s  3.22%  runtime.mallocgcSmallScanNoHeader
     0.22s  0.85% 68.41%      5.01s 19.44%  crypto/internal/fips140/sha256.(*Digest).Write
     0.21s  0.81% 69.23%      5.39s 20.92%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.18s   0.7% 69.93%      1.47s  5.70%  runtime.mallocgc
     0.17s  0.66% 70.59%      5.06s 19.64%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.17s  0.66% 71.25%      1.63s  6.33%  runtime.mapassign_faststr
     0.17s  0.66% 71.91%      0.17s  0.66%  sync/atomic.(*Int32).Add (inline)
     0.15s  0.58% 72.49%      0.19s  0.74%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.15s  0.58% 73.07%      1.16s  4.50%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.14s  0.54% 73.61%      5.73s 22.24%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.14s  0.54% 74.16%      0.22s  0.85%  internal/runtime/maps.(*Iter).Next
     0.13s   0.5% 74.66%      1.89s  7.33%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.12s  0.47% 75.13%      0.13s   0.5%  example.com/isureal-company/isureal.rmLess
     0.10s  0.39% 75.51%      6.17s 23.94%  example.com/isureal-company/isureal.computeImageSignature
     0.10s  0.39% 75.90%      0.22s  0.85%  runtime.greyobject
     0.09s  0.35% 76.25%      0.26s  1.01%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.08s  0.31% 76.56%      4.66s 18.08%  crypto/internal/fips140/sha256.block
     0.07s  0.27% 76.83%      0.25s  0.97%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.07s  0.27% 77.11%      0.48s  1.86%  sort.insertionSort_func
     0.06s  0.23% 77.34%      0.19s  0.74%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.06s  0.23% 77.57%      3.80s 14.75%  example.com/isureal-company/isureal.loadBig
     0.06s  0.23% 77.80%      0.52s  2.02%  runtime.slicebytetostring
     0.05s  0.19% 78.00%      0.50s  1.94%  database/sql.(*Rows).scanLocked
     0.05s  0.19% 78.19%      0.31s  1.20%  runtime.mallocgcSmallNoscan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 35.9% / 100% | 16% | 0.6% | 5 | 102 MB | run queue が積もっている |
| isu2 (2 cpu) | 24.2% / 100% | 5% | 0.7% | 4 | 190 MB |  |
| isu3 (2 cpu) | 26.2% / 100% | 9% | 1.6% | 5 | 164 MB | run queue が積もっている |
| isu4 (2 cpu) | 60.4% / 100% | 54% | 2.5% | 11 | 99 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 58% / nginx 9% |
| isu2 | isureal 44% / nginx 2% |
| isu3 | isureal 48% / nginx 2% |
| isu4 | isureal 98% / postgres 16% |

## Diff from previous run

スコア: 205275 → 198345 (-6930)

`73b6aff..6155bdc`

```
 conf/isu1/home/isucon/env.sh | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)
```
