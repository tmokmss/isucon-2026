# run 20261031-171105-5c44717

## Summary

| | |
|---|---|
| score | 201585 |
| status | pass |
| ref | 5c44717 |
| sha | 5c4471724b9b |
| 所要時間 | 236s |
| 前回比 | -21890 (-9.8%) vs 20261031-170001-a1cea2f |
| ベスト比 | -30525 (-13.2%) vs 20261031-164055-bdf05cd = 232110 |
| note | asia1 GET memories export to EU US replicas on 885b12c |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	138	403	1137	1136
europe	7	19	57	170	465	465
america	6	18	55	159	347	346
```

```
{"id":"32861","status":"done","portalStatus":"finished","passed":true,"score":201585,"raw":201585,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:13:15.270Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:13:18.259Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:13:21.260Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:13:24.273Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T08:13:27.261Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T08:13:30.259Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T08:13:33.270Z level=INFO msg="asia 3 日目: 138 人"
time=2026-10-31T08:13:34.183Z level=INFO msg="user moved" user=wakere5623 from=asia to=europe
time=2026-10-31T08:13:34.820Z level=INFO msg="user moved" user=wakere5623 from=europe to=america
time=2026-10-31T08:13:36.260Z level=INFO msg="europe 3 日目: 57 人"
time=2026-10-31T08:13:36.578Z level=INFO msg="user moved" user=iskasol-6113 from=asia to=europe
time=2026-10-31T08:13:39.260Z level=INFO msg="america 3 日目: 55 人"
time=2026-10-31T08:13:42.271Z level=INFO msg="asia 4 日目: 403 人"
time=2026-10-31T08:13:44.516Z level=INFO msg="user moved" user=memu7282 from=asia to=america
time=2026-10-31T08:13:45.270Z level=INFO msg="europe 4 日目: 170 人"
time=2026-10-31T08:13:46.152Z level=INFO msg="user moved" user=wanayo_7994 from=asia to=america
time=2026-10-31T08:13:46.683Z level=INFO msg="user moved" user=leomar7810 from=asia to=europe
time=2026-10-31T08:13:46.855Z level=INFO msg="user moved" user=soko-8469 from=asia to=europe
time=2026-10-31T08:13:47.624Z level=INFO msg="user moved" user=iskasol-6113 from=europe to=asia
time=2026-10-31T08:13:47.630Z level=INFO msg="user moved" user=wanayo_7994 from=america to=europe
time=2026-10-31T08:13:47.833Z level=INFO msg="user moved" user=memu7282 from=america to=asia
time=2026-10-31T08:13:48.266Z level=INFO msg="america 4 日目: 159 人"
time=2026-10-31T08:13:50.904Z level=INFO msg="user moved" user=wakere5623 from=america to=asia
time=2026-10-31T08:13:51.278Z level=INFO msg="asia 5 日目: 1137 人"
time=2026-10-31T08:13:54.701Z level=INFO msg="europe 5 日目: 465 人"
time=2026-10-31T08:13:57.956Z level=INFO msg="america 5 日目: 347 人"
time=2026-10-31T08:14:07.012Z level=INFO msg="america 最終日（グランプリ）: 346 人"
time=2026-10-31T08:14:07.845Z level=INFO msg="asia 最終日（グランプリ）: 1136 人"
time=2026-10-31T08:14:08.181Z level=INFO msg="europe 最終日（グランプリ）: 465 人"
time=2026-10-31T08:14:15.385Z level=INFO msg="グランプリ: アジア 492 人, ヨーロッパ 139 人, アメリカ 201 人"
{
  "pass": true,
  "score": 201585,
  "messages": [
    "tsuivyhe.5053: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "evaivytor.5126: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "morumar18268: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "koyo.15308: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tsuanel_7930: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "shaguo_shadingyu_3710: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "sumibi_mini_turkey_3192: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "capybaras_en_promenade_4662: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
52.3%	1789	1	0	GET	/api/notifications	19430.148	10.861	16.613	16.636	172
7.9%	2794	236	0	POST	/api/posts	2927.405	1.048	6.702	7.314	600
6.2%	18286	382	0	GET	^/api/images/[^/]+	2296.678	0.126	1.840	3.635	137415
4.8%	15503	0	0	GET	^/api/feeds/[^/]+	1785.825	0.115	1.032	2.023	3620
4.7%	4274	0	0	GET	/api/memories	1759.042	0.412	2.312	3.046	12243
4.6%	6498	66	0	POST	/api/friend-requests	1692.961	0.261	1.569	2.631	114
2.3%	788	130	0	PUT	/api/me/profile-image	861.252	1.093	14.330	16.989	2
2.1%	3617	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	773.510	0.214	1.384	2.624	116
1.8%	6265	532	0	POST	/api/auth/login	679.097	0.108	1.067	2.259	225
1.8%	3082	0	0	GET	/api/users	660.124	0.214	1.099	1.566	245
1.6%	4059	0	0	GET	/api/posting-window/current	595.743	0.147	1.042	1.316	237
1.5%	867	7	0	POST	^/api/posts/[^/]+/realmojis$	550.550	0.635	2.959	3.392	470
1.4%	377	0	0	GET	/api/me/export	528.712	1.402	5.826	8.724	1875139
0.9%	900	0	0	GET	^/api/memories/[^/]+	326.779	0.363	1.737	2.044	599
0.9%	2075	2	0	GET	^/api/posts/me/current$	326.499	0.157	1.047	1.228	654
0.8%	759	0	0	POST	^/api/posts/[^/]+/comments$	301.571	0.397	2.413	3.064	347
0.5%	2346	53	0	GET	^/api/posts/[^/]+$	199.795	0.085	0.780	1.565	615
0.5%	1793	0	0	GET	^/api/relationships/[^/]+	195.437	0.109	1.059	1.336	1146
0.5%	672	3	0	GET	/api/me	176.960	0.263	1.171	1.784	301
0.4%	951	0	0	POST	/api/auth/logout	137.504	0.145	1.245	1.639	0
0.3%	1233	0	0	GET	^/api/users/[^/]+$	108.711	0.088	0.811	1.565	198
0.3%	1156	1	0	POST	/api/auth/register	105.190	0.091	0.493	0.749	227
0.3%	724	0	0	PUT	^/api/follows/[^/]+$	99.812	0.138	1.003	1.586	0
0.3%	163	0	0	DELETE	^/api/realmojis/[^/]+$	93.823	0.576	2.663	2.767	0
0.3%	581	0	0	PATCH	/api/me	93.034	0.160	1.207	2.016	278
0.2%	973	0	0	GET	^/api/posts/[^/]+/realmojis$	85.626	0.088	0.827	1.168	2199
0.2%	100	0	0	DELETE	^/api/posts/[^/]+$	80.610	0.806	2.013	3.230	0
0.2%	241	0	0	PATCH	^/api/posts/[^/]+$	79.967	0.332	1.989	2.101	651
0.2%	793	0	0	GET	^/api/posts/[^/]+/comments$	62.681	0.079	0.789	1.586	3154
0.1%	143	0	0	DELETE	^/api/comments/[^/]+$	47.985	0.336	1.638	2.114	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.7%	743	0	0	GET	/api/notifications	5607.221	7.547	13.088	13.092	173
9.0%	1131	72	0	POST	/api/posts	1149.077	1.016	6.330	6.745	615
8.6%	2494	25	0	POST	/api/friend-requests	1104.220	0.443	1.900	2.274	114
6.5%	4510	40	0	GET	^/api/images/[^/]+	839.299	0.186	1.708	2.615	185200
6.3%	4577	0	0	GET	^/api/feeds/[^/]+	802.163	0.175	1.329	1.805	2661
4.1%	1392	0	0	POST	/api/auth/login	522.556	0.375	1.769	2.442	245
2.4%	1649	0	0	GET	/api/posting-window/current	309.183	0.187	0.963	1.939	238
2.4%	460	0	0	POST	/api/auth/register	308.964	0.672	1.716	1.755	229
2.0%	269	0	0	PUT	/api/me/profile-image	254.803	0.947	2.400	2.548	0
2.0%	587	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	250.631	0.427	1.900	2.397	116
1.8%	1223	0	0	GET	/api/users	231.215	0.189	0.812	1.677	244
1.7%	1081	0	0	GET	/api/memories	213.468	0.197	1.952	3.693	13667
1.2%	286	0	0	PUT	^/api/follows/[^/]+$	155.759	0.545	1.991	2.295	0
1.2%	818	17	0	GET	^/api/posts/me/current$	149.763	0.183	1.456	1.790	643
1.1%	218	5	0	POST	^/api/posts/[^/]+/realmojis$	139.847	0.642	3.014	3.207	458
0.9%	230	0	0	PATCH	/api/me	121.310	0.527	2.167	2.568	282
0.7%	579	0	0	GET	^/api/relationships/[^/]+	94.857	0.164	1.425	1.757	669
0.7%	571	22	0	GET	^/api/posts/[^/]+$	91.027	0.159	1.013	1.755	605
0.7%	200	3	0	POST	^/api/posts/[^/]+/comments$	90.875	0.454	2.065	3.335	333
0.6%	223	0	0	POST	/api/auth/logout	72.254	0.324	1.846	2.044	0
0.5%	274	1	0	GET	/api/me	60.231	0.220	0.814	0.815	303
0.4%	79	1	0	PATCH	^/api/posts/[^/]+$	55.177	0.698	2.527	2.527	657
0.3%	42	0	0	DELETE	^/api/posts/[^/]+$	36.675	0.873	2.326	2.326	0
0.3%	270	0	0	GET	^/api/users/[^/]+$	35.681	0.132	0.704	0.937	198
0.2%	177	0	0	GET	^/api/memories/[^/]+	27.571	0.156	1.704	2.311	658
0.2%	177	0	0	GET	^/api/posts/[^/]+/comments$	25.520	0.144	0.815	1.040	2283
0.2%	182	0	0	GET	^/api/posts/[^/]+/realmojis$	25.013	0.137	0.822	0.855	1877
0.2%	31	0	0	DELETE	^/api/realmojis/[^/]+$	19.345	0.624	1.407	1.407	0
0.1%	40	0	0	DELETE	^/api/comments/[^/]+$	16.057	0.401	1.742	1.742	0
0.1%	76	0	0	GET	/api/me/export	16.015	0.211	2.762	2.762	1891111
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.1%	555	1	0	GET	/api/notifications	3692.686	6.653	9.951	9.951	171
8.6%	1847	13	0	POST	/api/friend-requests	735.283	0.398	1.763	2.037	114
8.1%	832	45	0	POST	/api/posts	692.382	0.832	6.536	6.885	622
7.8%	3959	14	0	GET	^/api/images/[^/]+	666.988	0.168	1.395	2.018	162022
6.0%	2779	0	0	GET	^/api/feeds/[^/]+	515.985	0.186	1.049	1.400	3209
4.8%	1199	0	0	POST	/api/auth/login	408.328	0.341	2.023	2.458	249
2.6%	346	0	0	POST	/api/auth/register	226.677	0.655	1.855	2.308	230
2.6%	1208	0	0	GET	/api/posting-window/current	218.907	0.181	0.942	1.106	239
2.1%	438	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	182.284	0.416	1.751	1.831	116
2.1%	912	0	0	GET	/api/users	177.399	0.195	0.898	1.373	247
2.0%	1015	0	0	GET	/api/memories	175.218	0.173	1.398	2.511	15832
1.2%	688	30	0	GET	^/api/posts/me/current$	102.105	0.148	0.911	1.398	629
1.1%	157	12	0	POST	^/api/posts/[^/]+/realmojis$	96.534	0.615	2.502	2.615	445
1.1%	196	0	0	PUT	/api/me/profile-image	94.130	0.480	2.178	2.262	0
0.9%	386	0	0	GET	^/api/relationships/[^/]+	76.051	0.197	1.170	1.345	820
0.7%	130	1	0	POST	^/api/posts/[^/]+/comments$	58.519	0.450	1.719	1.793	342
0.7%	224	0	0	PUT	^/api/follows/[^/]+$	57.105	0.255	1.517	1.798	0
0.7%	178	0	0	POST	/api/auth/logout	57.100	0.321	1.846	1.871	0
0.6%	149	0	0	PATCH	/api/me	53.214	0.357	1.733	2.055	283
0.6%	362	22	0	GET	^/api/posts/[^/]+$	47.976	0.133	0.526	0.659	596
0.6%	73	0	0	GET	/api/me/export	47.727	0.654	7.273	7.273	3120074
0.4%	201	1	0	GET	/api/me	32.333	0.161	0.695	0.879	303
0.3%	182	0	0	GET	^/api/memories/[^/]+	28.246	0.155	1.062	1.110	1010
0.3%	155	0	0	GET	^/api/users/[^/]+$	23.606	0.152	0.793	1.004	206
0.3%	48	11	0	PATCH	^/api/posts/[^/]+$	22.686	0.473	1.625	1.625	523
0.2%	21	0	0	DELETE	^/api/posts/[^/]+$	20.786	0.990	2.149	2.149	0
0.2%	22	0	0	DELETE	^/api/realmojis/[^/]+$	19.658	0.894	2.606	2.606	0
0.2%	123	0	0	GET	^/api/posts/[^/]+/realmojis$	18.969	0.154	0.588	0.831	1979
0.2%	98	0	0	GET	^/api/posts/[^/]+/comments$	13.509	0.138	0.895	0.895	2863
0.1%	19	0	0	DELETE	^/api/comments/[^/]+$	9.292	0.489	1.202	1.202	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16215 クエリ / 48 種類 / 合計 102.179s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	25.9%	4	26.474	6618.44	10040.83	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
2	19.6%	12323	20.073	1.63	9.66	9.6	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
3	18.7%	4	19.088	4772.03	7195.71	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	17.7%	31	18.103	583.98	1868.69	1561.9	155.7	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
5	8.8%	4	9.025	2256.20	3433.55	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	5.1%	7	5.218	745.41	3285.97	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	1.0%	7	1.070	152.90	608.30	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.7%	1	0.667	666.54	666.54	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
9	0.6%	7	0.604	86.23	161.17	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.4%	1	0.454	454.24	454.24	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.3%	1	0.352	352.18	352.18	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.2%	1	0.173	172.69	172.69	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.140	139.58	139.58	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.131	130.67	130.67	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
15	0.1%	1	0.092	92.14	92.14	2412.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 2
SELECT data FROM image_objects WHERE id = $1 AND created_at = $2;
-- rank 3
SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM posts;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: eb4bcb720093ba7b10c93ac3391a33aa4b45df5a
Type: cpu
Time: 2026-10-31 17:12:48 JST
Duration: 60s, Total samples = 26.27s (43.78%)
Showing nodes accounting for 20.37s, 77.54% of 26.27s total
Dropped 815 nodes (cum <= 0.13s)
Showing top 40 nodes out of 233
      flat  flat%   sum%        cum   cum%
     4.77s 18.16% 18.16%      4.77s 18.16%  crypto/internal/fips140/sha256.blockAVX2
     3.12s 11.88% 30.03%      3.12s 11.88%  internal/runtime/syscall/linux.Syscall6
     1.76s  6.70% 36.73%      5.92s 22.54%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.33s  5.06% 41.80%      1.33s  5.06%  runtime.memmove
     1.28s  4.87% 46.67%      1.28s  4.87%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.94s  3.58% 50.25%      0.94s  3.58%  aeshashbody
     0.81s  3.08% 53.33%      0.81s  3.08%  runtime.memclrNoHeapPointers
     0.59s  2.25% 55.58%      2.09s  7.96%  runtime.scanObject
     0.56s  2.13% 57.71%      0.74s  2.82%  runtime.findObject
     0.50s  1.90% 59.61%      1.87s  7.12%  runtime.mapaccess1_faststr
     0.41s  1.56% 61.17%      0.64s  2.44%  runtime.tryDeferToSpanScan
     0.35s  1.33% 62.50%      0.58s  2.21%  example.com/isureal-company/isureal.postLess
     0.27s  1.03% 63.53%      0.27s  1.03%  sync/atomic.(*Int32).Add (inline)
     0.24s  0.91% 64.45%      0.24s  0.91%  runtime.nextFreeFast (inline)
     0.23s  0.88% 65.32%      0.23s  0.88%  memeqbody
     0.22s  0.84% 66.16%      0.22s  0.84%  runtime.futex
     0.20s  0.76% 66.92%      5.17s 19.68%  crypto/internal/fips140/sha256.(*Digest).Write
     0.20s  0.76% 67.68%      0.32s  1.22%  database/sql.convertAssignRows
     0.18s  0.69% 68.37%      1.45s  5.52%  runtime.mapassign_faststr
     0.17s  0.65% 69.01%      5.47s 20.82%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.17s  0.65% 69.66%      0.79s  3.01%  runtime.mallocgcSmallScanNoHeader
     0.17s  0.65% 70.31%      0.17s  0.65%  runtime.spanClass.sizeclass (inline)
     0.16s  0.61% 70.92%      0.19s  0.72%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.16s  0.61% 71.53%      6.10s 23.22%  example.com/isureal-company/isureal.computeImageSignature
     0.15s  0.57% 72.10%      5.21s 19.83%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.15s  0.57% 72.67%      0.15s  0.57%  example.com/isureal-company/isureal.rmLess
     0.14s  0.53% 73.20%      5.76s 21.93%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.14s  0.53% 73.73%      0.14s  0.53%  internal/runtime/maps.(*Map).directoryAt (inline)
     0.14s  0.53% 74.27%      1.58s  6.01%  runtime.mallocgc
     0.13s  0.49% 74.76%      0.15s  0.57%  runtime.(*mspan).writeHeapBitsSmall
     0.10s  0.38% 75.14%      1.63s  6.20%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.09s  0.34% 75.49%      0.24s  0.91%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.09s  0.34% 75.83%      0.15s  0.57%  internal/runtime/maps.(*Iter).Next
     0.08s   0.3% 76.13%      0.98s  3.73%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.08s   0.3% 76.44%      0.18s  0.69%  runtime.greyobject
     0.07s  0.27% 76.70%      3.25s 12.37%  example.com/isureal-company/isureal.loadBig
     0.07s  0.27% 76.97%      0.21s   0.8%  runtime.mapaccess2_faststr
     0.05s  0.19% 77.16%      0.22s  0.84%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.05s  0.19% 77.35%      0.82s  3.12%  runtime.growslice
     0.05s  0.19% 77.54%      0.36s  1.37%  runtime.mallocgcSmallNoscan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 33.9% / 100% | 15% | 1.5% | 6 | 104 MB | run queue が積もっている |
| isu2 (2 cpu) | 33.0% / 100% | 13% | 7.4% | 9 | 104 MB | run queue が積もっている |
| isu3 (2 cpu) | 32.9% / 100% | 15% | 7.6% | 10 | 112 MB | run queue が積もっている |
| isu4 (2 cpu) | 33.5% / 100% | 14% | 4.6% | 5 | 99 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 50% / nginx 12% |
| isu2 | isureal 48% / nginx 2% |
| isu3 | isureal 49% / nginx 2% |
| isu4 | isureal 43% / postgres 12% |

## Diff from previous run

スコア: 223475 → 201585 (-21890)

`a1cea2f..5c44717`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf |  15 +++-
 conf/isu1/home/isucon/env.sh                   |   2 +-
 conf/isu4/home/isucon/env.sh                   |   1 -
 webapp/go/cmd/isureal/main.go                  |   3 -
 webapp/go/image_signature.go                   |   1 -
 webapp/go/offload.go                           |   4 -
 webapp/go/sign_share.go                        | 105 -------------------------
 7 files changed, 14 insertions(+), 117 deletions(-)
```
