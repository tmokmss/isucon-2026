# run 20261031-161627-94b75aa

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | presign-next-minute |
| sha | 94b75aa284de |
| 所要時間 | 236s |
| 前回比 | -175643 (-100.0%) vs 20261031-160719-1f0e740 |
| ベスト比 | -201497 (-100.0%) vs 20261031-160409-40e65ef = 201497 |
| note | on 73b6aff: sig cache by exp generation, fixed-size keys only |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/agent-aff106d587106aae5/runs/20261031-161627-94b75aa/bench.log
```

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	149	399	699	700
europe	7	19	55	162	444	443
america	6	18	55	157	422	421
```

```
{"id":"27473","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"post-validation evidence"}

--- stdout ---
time=2026-10-31T07:19:16.150Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:19:19.137Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:19:22.138Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:19:25.183Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T07:19:27.048Z level=INFO msg="user moved" user=meyo_5223 from=asia to=america
time=2026-10-31T07:19:28.137Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T07:19:28.918Z level=INFO msg="user moved" user=meyo_5223 from=america to=europe
time=2026-10-31T07:19:31.138Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T07:19:32.074Z level=INFO msg="user moved" user=hohe.5405 from=europe to=america
time=2026-10-31T07:19:34.150Z level=INFO msg="asia 3 日目: 149 人"
time=2026-10-31T07:19:37.137Z level=INFO msg="europe 3 日目: 55 人"
time=2026-10-31T07:19:38.606Z level=INFO msg="user moved" user=chitsu-6010 from=asia to=america
time=2026-10-31T07:19:39.633Z level=INFO msg="user moved" user=chitsu-6010 from=america to=europe
time=2026-10-31T07:19:40.172Z level=INFO msg="america 3 日目: 55 人"
time=2026-10-31T07:19:40.347Z level=INFO msg="user moved" user=meyo_5223 from=europe to=asia
time=2026-10-31T07:19:40.768Z level=INFO msg="user moved" user=teteha_5629 from=asia to=europe
time=2026-10-31T07:19:43.178Z level=INFO msg="asia 4 日目: 399 人"
time=2026-10-31T07:19:45.734Z level=INFO msg="user moved" user=hohe.5405 from=america to=europe
time=2026-10-31T07:19:46.159Z level=INFO msg="europe 4 日目: 162 人"
time=2026-10-31T07:19:47.132Z level=INFO msg="user moved" user=hiwaha_6695 from=america to=asia
time=2026-10-31T07:19:48.809Z level=INFO msg="user moved" user=hiorul_8474 from=europe to=asia
time=2026-10-31T07:19:48.970Z level=INFO msg="user moved" user=nomarro_7563 from=asia to=america
time=2026-10-31T07:19:49.169Z level=INFO msg="america 4 日目: 157 人"
time=2026-10-31T07:19:50.049Z level=INFO msg="user moved" user=chitsu-6010 from=europe to=america
time=2026-10-31T07:19:51.870Z level=INFO msg="user moved" user=vinrimax-8639 from=europe to=asia
time=2026-10-31T07:19:52.154Z level=INFO msg="asia 5 日目: 699 人"
time=2026-10-31T07:19:52.722Z level=INFO msg="user moved" user=fuliaso-8849 from=europe to=asia
time=2026-10-31T07:19:54.243Z level=INFO msg="user moved" user=toranku_8447 from=europe to=asia
time=2026-10-31T07:19:55.178Z level=INFO msg="europe 5 日目: 444 人"
time=2026-10-31T07:19:58.872Z level=INFO msg="america 5 日目: 422 人"
time=2026-10-31T07:20:07.157Z level=INFO msg="america 最終日（グランプリ）: 421 人"
time=2026-10-31T07:20:07.178Z level=INFO msg="europe 最終日（グランプリ）: 443 人"
time=2026-10-31T07:20:07.343Z level=INFO msg="asia 最終日（グランプリ）: 700 人"
time=2026-10-31T07:20:16.226Z level=INFO msg="グランプリ: アジア 398 人, ヨーロッパ 342 人, アメリカ 286 人"
{
  "pass": false,
  "score": 0,
  "messages": [
    "post-validation evidence: GET /api/images/aee5f59d-333f-42aa-ba53-8a6c5c420c71?exp=1793517660\u0026sig=4x_32mGwLRxqXTUV4jcipThfa5WnI6X94J6W7_vCZss: ステータスコード 403 が返りました（期待値 200、image_forbidden）",
    "shihama.5127: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "vinma-15808: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chitsu-6010: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yuxin_wu_2929: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.7%	1354	0	0	GET	/api/notifications	13402.035	9.898	15.212	15.214	174
14.4%	2023	79	0	POST	/api/posts	4401.486	2.176	60.650	62.824	628
9.7%	9883	0	0	GET	^/api/feeds/[^/]+	2983.042	0.302	1.963	4.472	4705
7.3%	2870	0	0	GET	/api/memories	2246.534	0.783	5.268	9.448	10965
3.9%	2753	0	0	GET	/api/posting-window/current	1201.341	0.436	2.333	2.690	237
2.5%	1919	0	0	GET	/api/users	768.984	0.401	2.358	2.816	240
2.3%	1678	1	0	GET	^/api/posts/me/current$	699.587	0.417	1.580	2.421	653
2.2%	11599	251	0	GET	^/api/images/[^/]+	673.369	0.058	0.903	2.151	170863
2.2%	208	0	0	GET	/api/me/export	667.416	3.209	13.773	19.846	2035333
1.6%	4100	44	0	POST	/api/friend-requests	496.614	0.121	1.003	1.446	114
1.4%	536	130	0	PUT	/api/me/profile-image	420.683	0.785	11.131	13.049	4
1.2%	711	12	0	POST	^/api/posts/[^/]+/realmojis$	373.880	0.526	5.770	6.033	463
1.1%	692	12	0	GET	^/api/posts/[^/]+/realmojis$	347.079	0.502	3.754	4.288	2092
1.0%	413	3	0	GET	/api/me	308.575	0.747	1.925	2.731	301
1.0%	1268	0	0	GET	^/api/relationships/[^/]+	307.703	0.243	1.514	2.441	1498
0.9%	1725	78	0	GET	^/api/posts/[^/]+$	275.045	0.159	1.274	1.783	612
0.6%	3186	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	189.606	0.060	0.703	1.191	116
0.6%	520	0	0	GET	^/api/memories/[^/]+	186.515	0.359	2.460	2.890	721
0.5%	649	8	0	GET	^/api/posts/[^/]+/comments$	145.712	0.225	1.439	1.783	2757
0.4%	826	0	0	GET	^/api/users/[^/]+$	137.468	0.166	1.445	1.755	220
0.4%	3891	555	0	POST	/api/auth/login	121.450	0.031	0.678	1.267	213
0.2%	783	0	0	POST	^/api/posts/[^/]+/comments$	51.375	0.066	0.752	1.124	349
0.1%	71	0	0	DELETE	^/api/posts/[^/]+$	40.454	0.570	1.240	1.240	0
0.1%	452	0	0	PUT	^/api/follows/[^/]+$	39.595	0.088	1.022	1.209	0
0.1%	193	0	0	PATCH	^/api/posts/[^/]+$	37.722	0.195	1.065	1.137	649
0.1%	370	0	0	PATCH	/api/me	30.246	0.082	0.821	1.089	281
0.1%	544	0	0	POST	/api/auth/logout	28.992	0.053	0.906	1.208	0
0.1%	116	0	0	DELETE	^/api/realmojis/[^/]+$	25.726	0.222	0.595	0.614	0
0.1%	1	0	0	POST	/api/initialize	23.453	23.453	23.453	23.453	90
0.0%	131	0	0	DELETE	^/api/comments/[^/]+$	5.121	0.039	0.203	0.290	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.5%	713	1	0	GET	/api/notifications	5273.502	7.396	11.988	11.989	170
11.7%	1146	19	0	POST	/api/posts	1390.237	1.213	7.385	62.058	639
7.9%	7139	0	0	GET	^/api/feeds/[^/]+	932.497	0.131	1.581	2.283	4080
5.8%	6169	47	0	GET	^/api/images/[^/]+	690.952	0.112	1.061	2.821	185779
5.7%	2648	33	0	POST	/api/friend-requests	673.577	0.254	1.154	1.700	114
3.5%	1799	0	0	POST	/api/auth/login	415.581	0.231	1.102	1.851	247
2.4%	1562	0	0	GET	/api/memories	288.808	0.185	1.871	4.084	12380
2.0%	242	0	0	PUT	/api/me/profile-image	238.870	0.987	2.893	5.699	0
2.0%	1627	0	0	GET	/api/posting-window/current	237.150	0.146	1.219	2.084	239
1.8%	445	0	0	POST	/api/auth/register	211.748	0.476	1.277	1.324	228
1.7%	426	5	0	POST	^/api/posts/[^/]+/realmojis$	196.858	0.462	2.908	5.768	461
1.6%	1024	0	0	GET	^/api/posts/me/current$	184.757	0.180	1.721	2.218	649
1.4%	1314	0	0	GET	/api/users	162.692	0.124	0.534	1.291	246
1.1%	673	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	135.019	0.201	1.154	1.470	116
0.9%	1074	22	0	GET	^/api/posts/[^/]+$	102.761	0.096	0.558	1.634	629
0.8%	877	0	0	GET	^/api/relationships/[^/]+	97.889	0.112	0.855	2.017	536
0.8%	282	0	0	PUT	^/api/follows/[^/]+$	90.100	0.320	1.039	1.074	0
0.6%	387	2	0	POST	^/api/posts/[^/]+/comments$	72.121	0.186	0.978	1.219	329
0.6%	223	0	0	PATCH	/api/me	71.101	0.319	1.041	1.181	275
0.5%	302	0	0	POST	/api/auth/logout	56.056	0.186	1.132	1.646	0
0.4%	533	0	0	GET	^/api/users/[^/]+$	49.289	0.092	0.522	0.721	223
0.4%	521	0	0	GET	^/api/posts/[^/]+/realmojis$	47.188	0.091	0.560	0.887	3308
0.3%	57	0	0	DELETE	^/api/posts/[^/]+$	39.762	0.698	2.446	2.446	0
0.3%	427	0	0	GET	^/api/posts/[^/]+/comments$	38.110	0.089	0.426	0.707	3941
0.3%	93	0	0	PATCH	^/api/posts/[^/]+$	35.081	0.377	1.715	1.715	641
0.2%	70	0	0	DELETE	^/api/realmojis/[^/]+$	27.885	0.398	0.979	0.979	0
0.2%	246	2	0	GET	/api/me	27.861	0.113	0.390	0.535	303
0.2%	247	0	0	GET	^/api/memories/[^/]+	21.061	0.085	0.270	2.058	784
0.1%	106	0	0	GET	/api/me/export	15.079	0.142	1.003	2.077	2228752
0.1%	76	0	0	DELETE	^/api/comments/[^/]+$	12.300	0.162	0.479	0.479	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.4%	672	1	0	GET	/api/notifications	3907.263	5.814	9.716	9.720	175
12.1%	1031	33	0	POST	/api/posts	1141.220	1.107	6.756	62.835	633
8.0%	4711	32	0	GET	^/api/images/[^/]+	759.060	0.161	1.658	4.033	150310
7.6%	3458	0	0	GET	^/api/feeds/[^/]+	721.847	0.209	2.052	2.604	2865
5.9%	2295	17	0	POST	/api/friend-requests	553.001	0.241	0.902	2.211	114
3.8%	1500	0	0	POST	/api/auth/login	355.781	0.237	1.206	2.289	249
3.2%	1292	0	0	GET	/api/memories	297.695	0.230	2.442	4.543	14865
2.7%	1513	0	0	GET	/api/posting-window/current	253.301	0.167	0.951	2.404	238
2.2%	1069	0	0	GET	/api/users	208.981	0.195	0.743	2.530	246
2.0%	848	22	0	GET	^/api/posts/me/current$	192.094	0.227	1.889	2.613	637
1.6%	223	0	0	PUT	/api/me/profile-image	150.392	0.674	2.870	2.978	0
1.3%	420	0	0	POST	/api/auth/register	119.867	0.285	0.659	0.734	230
1.0%	365	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	92.148	0.252	1.590	1.872	116
0.9%	161	13	0	POST	^/api/posts/[^/]+/realmojis$	85.605	0.532	2.253	5.314	433
0.9%	488	0	0	GET	^/api/relationships/[^/]+	82.431	0.169	1.528	2.317	443
0.7%	284	0	0	PUT	^/api/follows/[^/]+$	70.356	0.248	0.575	1.305	0
0.7%	422	23	0	GET	^/api/posts/[^/]+$	64.333	0.152	0.763	1.727	605
0.6%	248	0	0	POST	/api/auth/logout	54.476	0.220	1.204	2.157	0
0.5%	189	4	0	POST	^/api/posts/[^/]+/comments$	47.139	0.249	1.098	1.211	320
0.5%	190	0	0	PATCH	/api/me	46.125	0.243	0.583	0.669	280
0.3%	226	1	0	GET	/api/me	32.622	0.144	0.521	0.699	305
0.3%	196	0	0	GET	^/api/memories/[^/]+	29.865	0.152	2.211	2.520	1051
0.3%	87	0	0	GET	/api/me/export	28.567	0.328	4.425	4.425	2779415
0.3%	195	0	0	GET	^/api/users/[^/]+$	28.370	0.145	0.727	0.886	216
0.3%	175	0	0	GET	^/api/posts/[^/]+/comments$	27.134	0.155	0.689	0.719	2557
0.3%	38	0	0	DELETE	^/api/posts/[^/]+$	25.684	0.676	2.390	2.390	0
0.2%	144	4	0	GET	^/api/posts/[^/]+/realmojis$	20.836	0.145	0.685	0.726	2311
0.2%	75	9	0	PATCH	^/api/posts/[^/]+$	19.958	0.266	1.127	1.127	581
0.1%	32	0	0	DELETE	^/api/comments/[^/]+$	9.868	0.308	1.118	1.118	0
0.1%	13	0	0	DELETE	^/api/realmojis/[^/]+$	6.141	0.472	1.126	1.126	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 9877 クエリ / 48 種類 / 合計 136.894s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	42.7%	214	58.504	273.38	1177.86	1301.6	132.7	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	19.5%	4	26.627	6656.80	10026.64	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.3%	4	19.635	4908.73	7302.29	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.3%	7070	11.349	1.61	9.04	10.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.5%	4	8.901	2225.34	3430.46	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	5.6%	6	7.665	1277.46	4195.62	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	1.0%	1	1.428	1428.45	1428.45	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.3%	6	0.441	73.53	162.78	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.3%	6	0.440	73.40	157.44	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
10	0.3%	1	0.403	403.25	403.25	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.325	325.35	325.35	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.2%	1	0.265	265.20	265.20	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.163	163.09	163.09	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.110	110.36	110.36	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	1	0.107	106.66	106.66	2486.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…

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
Build ID: 86ee2acaab68161eeb395570172a7829f04d19ba
Type: cpu
Time: 2026-10-31 16:18:47 JST
Duration: 60s, Total samples = 25.25s (42.08%)
Showing nodes accounting for 19.51s, 77.27% of 25.25s total
Dropped 722 nodes (cum <= 0.13s)
Showing top 40 nodes out of 245
      flat  flat%   sum%        cum   cum%
     3.65s 14.46% 14.46%      3.65s 14.46%  crypto/internal/fips140/sha256.blockAVX2
     2.50s  9.90% 24.36%      2.50s  9.90%  internal/runtime/syscall/linux.Syscall6
     1.68s  6.65% 31.01%      5.93s 23.49%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.39s  5.50% 36.51%      1.39s  5.50%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.33s  5.27% 41.78%      1.33s  5.27%  runtime.memmove
     1.02s  4.04% 45.82%      1.02s  4.04%  aeshashbody
     0.77s  3.05% 48.87%      2.76s 10.93%  runtime.scanObject
     0.67s  2.65% 51.52%      0.99s  3.92%  runtime.tryDeferToSpanScan
     0.60s  2.38% 53.90%      0.60s  2.38%  runtime.memclrNoHeapPointers
     0.58s  2.30% 56.20%      0.78s  3.09%  runtime.findObject
     0.44s  1.74% 57.94%      1.95s  7.72%  runtime.mapaccess1_faststr
     0.34s  1.35% 59.29%      0.34s  1.35%  runtime.nextFreeFast (inline)
     0.33s  1.31% 60.59%      0.33s  1.31%  memeqbody
     0.33s  1.31% 61.90%      0.33s  1.31%  runtime.futex
     0.32s  1.27% 63.17%      0.46s  1.82%  database/sql.convertAssignRows
     0.32s  1.27% 64.44%      0.49s  1.94%  example.com/isureal-company/isureal.postLess
     0.25s  0.99% 65.43%      0.78s  3.09%  runtime.mallocgcSmallScanNoHeader
     0.25s  0.99% 66.42%      0.25s  0.99%  sync/atomic.(*Int32).Add (inline)
     0.23s  0.91% 67.33%      4.12s 16.32%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.21s  0.83% 68.16%      0.21s  0.83%  runtime.spanClass.sizeclass (inline)
     0.20s  0.79% 68.95%      4.03s 15.96%  crypto/internal/fips140/sha256.(*Digest).Write
     0.18s  0.71% 69.66%      0.20s  0.79%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.15s  0.59% 70.26%      2.02s  8.00%  runtime.mallocgc
     0.15s  0.59% 70.85%      0.35s  1.39%  runtime.mapaccess2_faststr
     0.15s  0.59% 71.45%      1.41s  5.58%  runtime.mapassign_faststr
     0.14s  0.55% 72.00%      0.15s  0.59%  example.com/isureal-company/isureal.rmLess
     0.13s  0.51% 72.51%      2.10s  8.32%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.12s  0.48% 72.99%      4.53s 17.94%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.12s  0.48% 73.47%      4.26s 16.87%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.12s  0.48% 73.94%      0.16s  0.63%  runtime.(*mspan).markBitsForIndex (inline)
     0.11s  0.44% 74.38%      4.12s 16.32%  example.com/isureal-company/isureal.loadBig
     0.10s   0.4% 74.77%      0.14s  0.55%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.10s   0.4% 75.17%      1.10s  4.36%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.09s  0.36% 75.52%      1.58s  6.26%  encoding/gob.(*Encoder).encodeStruct
     0.09s  0.36% 75.88%      0.51s  2.02%  sort.insertionSort_func
     0.07s  0.28% 76.16%      0.21s  0.83%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.07s  0.28% 76.44%      4.83s 19.13%  example.com/isureal-company/isureal.computeImageSignature
     0.07s  0.28% 76.71%      0.27s  1.07%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.07s  0.28% 76.99%      0.40s  1.58%  github.com/jackc/pgx/v5/pgproto3.(*chunkReader).Next
     0.07s  0.28% 77.27%      0.47s  1.86%  runtime.convTstring
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 52.6% / 100% | 26% | 1.8% | 5 | 127 MB | run queue が積もっている |
| isu2 (2 cpu) | 32.6% / 100% | 11% | 0.7% | 4 | 138 MB |  |
| isu3 (2 cpu) | 33.6% / 100% | 14% | 1.4% | 4 | 133 MB |  |
| isu4 (2 cpu) | 82.3% / 100% | 72% | 4.3% | 4 | 105 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 74% / nginx 12% / md5sum 1% |
| isu2 | isureal 59% / nginx 3% |
| isu3 | isureal 61% / nginx 3% |
| isu4 | isureal 133% / postgres 19% |

## Diff from previous run

スコア: 175643 → — (fail)

`1f0e740..94b75aa`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf  | 25 ++++++++++++++-
 conf/isu1/home/isucon/env.sh                    |  2 +-
 conf/isu4/etc/systemd/system/isureal-go.service | 11 +++----
 conf/isu4/home/isucon/env.sh                    | 12 ++++++++
 webapp/go/image_signature.go                    |  2 +-
 webapp/go/offload.go                            | 41 ++-----------------------
 webapp/go/server.go                             |  1 -
 7 files changed, 45 insertions(+), 49 deletions(-)
```
