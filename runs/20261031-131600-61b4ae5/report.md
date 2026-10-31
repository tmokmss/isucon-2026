# run 20261031-131600-61b4ae5

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | asia1-cpu-small |
| sha | 61b4ae59ef7f |
| 所要時間 | 843s |
| 前回比 | -64700 (-100.0%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | -64700 (-100.0%) vs 20261031-121235-e0e4cb3 = 64700 |
| note | groupA-asia1-cpu-small |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/agent-a6a6b494091f487c9/runs/20261031-131600-61b4ae5/bench.log
```

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	149	416	968	969
europe	7	14	34	74	76	76
america	6	17	41	75	76	76
```

```
{"id":"12749","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"post-validation evidence"}

--- stdout ---
time=2026-10-31T04:28:55.304Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:28:58.271Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:29:01.267Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:29:04.293Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T04:29:07.270Z level=INFO msg="europe 2 日目: 14 人"
time=2026-10-31T04:29:07.457Z level=INFO msg="user moved" user=yuhi5338 from=asia to=america
time=2026-10-31T04:29:09.019Z level=INFO msg="user moved" user=yuhi5338 from=america to=asia
time=2026-10-31T04:29:10.278Z level=INFO msg="america 2 日目: 17 人"
time=2026-10-31T04:29:11.770Z level=INFO msg="user moved" user=solse_5403 from=europe to=asia
time=2026-10-31T04:29:13.300Z level=INFO msg="asia 3 日目: 149 人"
time=2026-10-31T04:29:15.732Z level=INFO msg="user moved" user=solansol_5839 from=asia to=europe
time=2026-10-31T04:29:16.091Z level=INFO msg="user moved" user=solse_5403 from=asia to=europe
time=2026-10-31T04:29:16.156Z level=INFO msg="user moved" user=numia.6006 from=asia to=europe
time=2026-10-31T04:29:16.268Z level=INFO msg="europe 3 日目: 34 人"
time=2026-10-31T04:29:19.267Z level=INFO msg="america 3 日目: 41 人"
time=2026-10-31T04:29:20.161Z level=INFO msg="user moved" user=hakatsu-6221 from=europe to=asia
time=2026-10-31T04:29:22.328Z level=INFO msg="asia 4 日目: 416 人"
time=2026-10-31T04:29:25.296Z level=INFO msg="europe 4 日目: 74 人"
time=2026-10-31T04:29:25.429Z level=INFO msg="user moved" user=hakatsu-6221 from=asia to=america
time=2026-10-31T04:29:25.522Z level=INFO msg="user moved" user=yuhi5338 from=asia to=america
time=2026-10-31T04:29:26.731Z level=INFO msg="user moved" user=nakochi-7722 from=asia to=europe
time=2026-10-31T04:29:27.031Z level=INFO msg="user moved" user=sanoako.6746 from=asia to=europe
time=2026-10-31T04:29:27.339Z level=INFO msg="user moved" user=noasu6666 from=asia to=america
time=2026-10-31T04:29:28.330Z level=INFO msg="america 4 日目: 75 人"
time=2026-10-31T04:29:29.332Z level=INFO msg="user moved" user=orsamo_7726 from=asia to=america
time=2026-10-31T04:29:30.488Z level=INFO msg="user moved" user=numia.6006 from=europe to=asia
time=2026-10-31T04:29:30.630Z level=INFO msg="user moved" user=hoso.8349 from=asia to=europe
time=2026-10-31T04:29:31.556Z level=INFO msg="asia 5 日目: 968 人"
time=2026-10-31T04:29:31.845Z level=INFO msg="user moved" user=solse_5403 from=europe to=asia
time=2026-10-31T04:29:34.366Z level=INFO msg="europe 5 日目: 76 人"
time=2026-10-31T04:29:37.357Z level=INFO msg="america 5 日目: 76 人"
time=2026-10-31T04:29:46.325Z level=INFO msg="europe 最終日（グランプリ）: 76 人"
time=2026-10-31T04:29:46.371Z level=INFO msg="america 最終日（グランプリ）: 76 人"
time=2026-10-31T04:29:46.634Z level=INFO msg="asia 最終日（グランプリ）: 969 人"
time=2026-10-31T04:29:55.303Z level=INFO msg="グランプリ: アジア 481 人, ヨーロッパ 21 人, アメリカ 17 人"
{
  "pass": false,
  "score": 0,
  "messages": [
    "post-validation evidence: GET /api/images/474382c5-accb-4d0c-95f6-8d2f5e538cf3?exp=1793507400\u0026sig=yWdo5nxeKf4a0Ono7svCxZf3n33gXi_qXmQzvT8f02M: ステータスコード 403 が返りました（期待値 200、image_forbidden）",
    "dantari.5068: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.8%	1589	1	0	GET	/api/notifications	13102.115	8.246	15.074	15.098	166
18.5%	2394	51	0	POST	/api/posts	6976.136	2.914	9.103	12.349	638
12.5%	5464	70	0	POST	/api/friend-requests	4721.874	0.864	6.001	11.236	114
8.6%	13957	733	0	GET	^/api/images/[^/]+	3231.724	0.232	3.673	10.073	145611
5.1%	6238	2059	0	POST	/api/auth/login	1927.589	0.309	4.083	9.238	173
4.4%	1941	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1667.086	0.859	6.255	9.488	116
3.6%	693	122	0	PUT	/api/me/profile-image	1358.160	1.960	7.875	8.352	4
2.0%	637	27	0	POST	^/api/posts/[^/]+/realmojis$	766.972	1.204	6.846	7.935	465
1.4%	655	15	0	POST	^/api/posts/[^/]+/comments$	536.400	0.819	5.718	9.376	327
1.3%	991	1	0	POST	/api/auth/register	489.668	0.494	4.582	6.885	227
1.3%	647	0	0	PUT	^/api/follows/[^/]+$	479.010	0.740	5.584	7.398	0
1.0%	3540	0	0	GET	/api/memories	389.100	0.110	0.652	1.481	10897
0.9%	458	0	0	PATCH	/api/me	352.222	0.769	4.244	8.330	277
0.7%	256	0	0	GET	/api/me/export	270.959	1.058	6.773	8.649	2089789
0.7%	664	0	0	POST	/api/auth/logout	267.195	0.402	5.556	10.669	0
0.7%	230	9	0	PATCH	^/api/posts/[^/]+$	262.495	1.141	6.739	8.167	626
0.6%	13117	0	0	GET	^/api/feeds/[^/]+	228.219	0.017	0.169	9.142	2519
0.4%	101	0	0	DELETE	^/api/posts/[^/]+$	154.400	1.529	5.998	7.301	0
0.2%	123	0	0	DELETE	^/api/comments/[^/]+$	72.869	0.592	5.015	6.788	0
0.2%	3447	0	0	GET	/api/posting-window/current	71.112	0.021	0.122	0.712	237
0.2%	97	0	0	DELETE	^/api/realmojis/[^/]+$	66.225	0.683	7.998	7.998	0
0.1%	1950	76	0	GET	^/api/posts/me/current$	38.743	0.020	0.188	0.713	626
0.1%	2623	0	0	GET	/api/users	32.216	0.012	0.111	0.419	240
0.1%	584	0	0	GET	^/api/memories/[^/]+	31.092	0.053	0.428	0.607	631
0.1%	45	0	0	PUT	^/api/blocks/[^/]+$	30.227	0.672	4.576	4.576	0
0.1%	1758	0	0	GET	^/api/relationships/[^/]+	27.406	0.016	0.179	0.527	845
0.1%	1774	16	0	GET	^/api/posts/[^/]+$	21.733	0.012	0.136	0.306	633
0.0%	1	0	0	POST	/api/initialize	12.290	12.290	12.290	12.290	90
0.0%	942	0	0	GET	^/api/users/[^/]+$	10.884	0.012	0.131	0.237	207
0.0%	826	0	0	GET	^/api/posts/[^/]+/realmojis$	10.406	0.013	0.128	2.301	1127
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.9%	216	0	0	GET	/api/notifications	1811.173	8.385	11.803	12.161	177
19.0%	245	6	0	POST	/api/posts	841.553	3.435	9.396	10.252	633
7.9%	1215	5	0	GET	^/api/images/[^/]+	348.932	0.287	1.995	2.881	169895
5.5%	435	8	0	POST	/api/friend-requests	242.753	0.558	6.272	6.933	113
4.8%	345	0	0	POST	/api/auth/login	214.820	0.623	4.587	6.650	249
4.7%	211	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	209.433	0.993	5.723	8.262	116
3.8%	983	0	0	GET	^/api/feeds/[^/]+	167.147	0.170	0.607	0.903	3532
2.4%	331	0	0	GET	/api/memories	104.473	0.316	1.255	1.605	15321
1.8%	334	0	0	GET	/api/posting-window/current	81.521	0.244	0.665	1.090	237
1.3%	62	5	0	POST	^/api/posts/[^/]+/realmojis$	55.535	0.896	4.578	4.578	382
1.0%	197	3	0	GET	^/api/posts/me/current$	44.092	0.224	0.922	0.946	639
0.8%	208	0	0	GET	/api/users	35.564	0.171	0.545	0.901	235
0.7%	20	0	0	GET	/api/me/export	32.066	1.603	3.978	3.978	2137774
0.7%	128	0	0	GET	^/api/relationships/[^/]+	28.955	0.226	0.599	0.684	1293
0.6%	46	0	0	POST	/api/auth/logout	25.162	0.547	3.429	3.429	0
0.6%	72	0	0	POST	/api/auth/register	24.726	0.343	1.378	1.378	228
0.5%	42	0	0	PUT	/api/me/profile-image	23.942	0.570	1.901	1.901	0
0.4%	131	8	0	GET	^/api/posts/[^/]+$	18.100	0.138	0.452	0.469	596
0.3%	54	0	0	GET	^/api/memories/[^/]+	13.479	0.250	1.213	1.213	1010
0.3%	55	0	0	PUT	^/api/follows/[^/]+$	12.766	0.232	1.080	1.080	0
0.3%	21	0	0	POST	^/api/posts/[^/]+/comments$	12.431	0.592	1.916	1.916	349
0.3%	17	0	0	DELETE	^/api/realmojis/[^/]+$	12.299	0.723	4.067	4.067	0
0.3%	23	1	0	PATCH	^/api/posts/[^/]+$	11.959	0.520	2.171	2.171	625
0.2%	69	0	0	GET	^/api/posts/[^/]+/realmojis$	10.349	0.150	0.508	0.508	999
0.2%	62	0	0	GET	^/api/users/[^/]+$	9.419	0.152	0.469	0.469	193
0.2%	28	0	0	PATCH	/api/me	7.991	0.285	0.870	0.870	296
0.1%	49	1	0	GET	^/api/posts/[^/]+/comments$	6.620	0.135	0.352	0.352	1334
0.1%	44	1	0	GET	/api/me	6.350	0.144	0.464	0.464	296
0.1%	5	0	0	DELETE	^/api/posts/[^/]+$	6.347	1.269	3.677	3.677	0
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	3.665	1.222	2.218	2.218	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.0%	209	0	0	GET	/api/notifications	1607.011	7.689	9.175	9.177	179
19.2%	236	6	0	POST	/api/posts	810.197	3.433	10.449	12.415	634
8.4%	1040	6	0	GET	^/api/images/[^/]+	353.558	0.340	3.068	4.419	173174
6.7%	396	4	0	POST	/api/friend-requests	282.079	0.712	5.210	11.826	114
4.7%	206	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	197.018	0.956	8.414	8.619	115
3.9%	300	0	0	POST	/api/auth/login	163.846	0.546	6.030	7.231	250
3.8%	817	0	0	GET	^/api/feeds/[^/]+	160.983	0.197	0.647	1.713	3729
2.4%	292	0	0	GET	/api/memories	99.970	0.342	1.303	1.448	14192
1.9%	24	0	0	GET	/api/me/export	80.843	3.368	9.154	9.154	3278297
1.9%	325	0	0	GET	/api/posting-window/current	80.197	0.247	0.684	1.122	240
0.9%	180	8	0	GET	^/api/posts/me/current$	39.908	0.222	0.723	1.501	621
0.9%	39	0	0	PUT	/api/me/profile-image	39.153	1.004	4.396	4.396	0
0.9%	47	0	0	POST	/api/auth/logout	38.658	0.823	8.328	8.328	0
0.8%	187	0	0	GET	/api/users	35.077	0.188	0.633	0.635	233
0.7%	42	3	0	POST	^/api/posts/[^/]+/comments$	29.691	0.707	5.356	5.356	357
0.6%	147	5	0	GET	^/api/posts/[^/]+$	24.637	0.168	0.475	0.547	610
0.6%	100	0	0	GET	^/api/relationships/[^/]+	23.922	0.239	0.559	0.559	1245
0.5%	71	0	0	POST	/api/auth/register	23.047	0.325	3.234	3.234	229
0.5%	24	1	0	POST	^/api/posts/[^/]+/realmojis$	21.670	0.903	3.846	3.846	425
0.5%	21	2	0	PATCH	^/api/posts/[^/]+$	19.764	0.941	5.782	5.782	607
0.5%	44	0	0	PUT	^/api/follows/[^/]+$	19.357	0.440	4.540	4.540	0
0.4%	34	0	0	PATCH	/api/me	17.048	0.501	3.668	3.668	267
0.3%	64	0	0	GET	^/api/users/[^/]+$	11.424	0.178	0.476	0.476	198
0.3%	8	0	0	DELETE	^/api/posts/[^/]+$	11.098	1.387	3.463	3.463	0
0.2%	43	0	0	GET	^/api/memories/[^/]+	9.921	0.231	0.852	0.852	1209
0.2%	40	0	0	GET	/api/me	7.982	0.200	0.599	0.599	303
0.2%	51	0	0	GET	^/api/posts/[^/]+/realmojis$	7.652	0.150	0.413	0.413	862
0.2%	43	0	0	GET	^/api/posts/[^/]+/comments$	7.081	0.165	0.419	0.419	1229
0.1%	4	0	0	DELETE	^/api/comments/[^/]+$	3.099	0.775	1.304	1.304	0
0.0%	1	0	0	DELETE	/api/me/profile-image	1.358	1.358	1.358	1.358	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 68877 クエリ / 81 種類 / 合計 665.502s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	68.5%	4155	455.773	109.69	593.01	682.0	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	19.8%	300	131.592	438.64	2045.98	7951.5	140.3	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	4.2%	1951	27.837	14.27	520.29	118.7	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	3.7%	8880	24.697	2.78	45.58	10.7	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.4%	2813	9.460	3.36	78.01	24.9	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.4%	690	2.964	4.30	29.41	28.1	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
7	0.4%	1	2.382	2381.71	2381.71	30603.0	1199910.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.2%	700	1.486	2.12	21.09	20.5	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
9	0.2%	6200	1.353	0.22	80.55	19.6	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	0.2%	2	1.327	663.56	1319.25	16784.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.2%	2	1.078	539.01	934.55	234371.0	2760.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
12	0.2%	1	1.026	1025.85	1025.85	12273.0	530240.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
13	0.1%	2288	0.839	0.37	52.11	20.7	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
14	0.1%	4824	0.629	0.13	90.14	11.2	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
15	0.0%	4825	0.279	0.06	100.47	3.6	0.2	UPDATE login_sessions SET active = $3, invalidated_at = $2, reason = $4 WHERE user_id = $1 AND active = $5

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 5c4fa41bf03490a011f2d82e4f5ba854d3ae7857
Type: cpu
Time: 2026-10-31 13:28:37 JST
Duration: 60.10s, Total samples = 40.41s (67.24%)
Showing nodes accounting for 28.20s, 69.78% of 40.41s total
Dropped 1052 nodes (cum <= 0.20s)
Showing top 40 nodes out of 327
      flat  flat%   sum%        cum   cum%
     4.29s 10.62% 10.62%      4.29s 10.62%  internal/runtime/syscall/linux.Syscall6
     3.01s  7.45% 18.06%      3.01s  7.45%  crypto/internal/fips140/sha256.blockAVX2
     2.23s  5.52% 23.58%      7.59s 18.78%  runtime.scanObject
     1.99s  4.92% 28.51%      2.51s  6.21%  runtime.findObject
     1.82s  4.50% 33.01%      5.96s 14.75%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.77s  4.38% 37.39%      1.77s  4.38%  runtime.memclrNoHeapPointers
     1.75s  4.33% 41.72%      1.75s  4.33%  runtime.memmove
     1.43s  3.54% 45.26%      2.66s  6.58%  runtime.tryDeferToSpanScan
     1.31s  3.24% 48.50%      1.31s  3.24%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.22s  3.02% 51.52%      1.22s  3.02%  aeshashbody
     0.91s  2.25% 53.77%      0.91s  2.25%  runtime.spanClass.sizeclass (inline)
     0.49s  1.21% 54.99%      2.12s  5.25%  runtime.mapaccess1_faststr
     0.38s  0.94% 55.93%      0.38s  0.94%  memeqbody
     0.38s  0.94% 56.87%      0.38s  0.94%  runtime.futex
     0.34s  0.84% 57.71%      0.34s  0.84%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.31s  0.77% 58.48%      0.53s  1.31%  example.com/isureal-company/isureal.postLess
     0.28s  0.69% 59.17%      0.28s  0.69%  crypto/internal/fips140/aes/gcm.gcmAesEnc
     0.28s  0.69% 59.86%      0.28s  0.69%  indexbody
     0.28s  0.69% 60.55%      0.28s  0.69%  runtime.nextFreeFast (inline)
     0.26s  0.64% 61.20%      0.33s  0.82%  encoding/json.appendString[go.shape.string]
     0.26s  0.64% 61.84%      1.71s  4.23%  runtime.mapassign_faststr
     0.25s  0.62% 62.46%      0.56s  1.39%  runtime.greyobject
     0.25s  0.62% 63.08%      0.28s  0.69%  runtime.spanOf (inline)
     0.25s  0.62% 63.70%      0.28s  0.69%  runtime.typePointers.next
     0.23s  0.57% 64.27%      0.23s  0.57%  runtime.typePointers.nextFast (inline)
     0.22s  0.54% 64.81%      3.38s  8.36%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.22s  0.54% 65.36%      0.28s  0.69%  runtime.(*mspan).markBitsForIndex (inline)
     0.21s  0.52% 65.87%      0.44s  1.09%  compress/flate.(*compressor).deflate
     0.16s   0.4% 66.27%      0.21s  0.52%  runtime.(*mspan).writeHeapBitsSmall
     0.16s   0.4% 66.67%      0.25s  0.62%  runtime.heapArenaOf (inline)
     0.16s   0.4% 67.06%      5.63s 13.93%  runtime.mallocgc
     0.15s  0.37% 67.43%      0.42s  1.04%  database/sql.convertAssignRows
     0.14s  0.35% 67.78%      0.34s  0.84%  internal/runtime/maps.(*Iter).Next
     0.13s  0.32% 68.10%      3.57s  8.83%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.12s   0.3% 68.40%      3.23s  7.99%  crypto/internal/fips140/sha256.(*Digest).Write
     0.12s   0.3% 68.70%      0.85s  2.10%  runtime.newobject
     0.12s   0.3% 68.99%      0.25s  0.62%  runtime.scanObjectSmall
     0.11s  0.27% 69.27%      0.48s  1.19%  runtime.mapaccess2_faststr
     0.11s  0.27% 69.54%      0.54s  1.34%  runtime.scanObjectsSmall
     0.10s  0.25% 69.78%      0.49s  1.21%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 64.6% / 100% | 38% | 1.9% | 11 | 101 MB | CPU が飽和 |
| isu2 (2 cpu) | 26.1% / 60% | 0% | 0.5% | 5 | 2043 MB | run queue が積もっている |
| isu3 (2 cpu) | 23.4% / 59% | 0% | 0.4% | 4 | 2467 MB |  |
| isu4 (2 cpu) | 68.9% / 100% | 42% | 19.9% | 36 | 107 MB | CPU が飽和 |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 78% / nginx 22% / kswapd0 3% |
| isu2 | isureal 50% / nginx 1% |
| isu3 | isureal 44% / nginx 1% |
| isu4 | postgres 86% / isureal 5% / kcompactd0 1% / kswapd0 1% |

## Diff from previous run

スコア: 64700 → — (fail)

`e0e4cb3..61b4ae5`

```
 webapp/go/account_deletion_api.go |   3 +-
 webapp/go/action_log.go           |  13 +-
 webapp/go/action_log_test.go      |  67 ++++++++
 webapp/go/admin_user_api.go       |   2 +-
 webapp/go/comment_api.go          |   2 +-
 webapp/go/feed_api.go             | 113 +++++++++----
 webapp/go/feed_test.go            | 331 ++++++++++++++++++++++++++++++++++++++
 webapp/go/offload.go              |  46 +++++-
 webapp/go/post_api.go             |  62 +++++--
 webapp/go/profile_image_api.go    |   7 +-
 webapp/go/realmoji_api.go         |   2 +-
 webapp/go/relationship.go         |  10 +-
 webapp/go/store.go                | 166 ++++++++++++++++---
 webapp/go/user_api.go             |   1 +
 webapp/go/user_read_api.go        |  58 ++++---
 webapp/go/user_search_test.go     |  94 +++++++++++
 16 files changed, 876 insertions(+), 101 deletions(-)
```
