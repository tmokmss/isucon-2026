# run 20261031-140544-247167c

## Summary

| | |
|---|---|
| score | 123551 |
| status | pass |
| ref | image-storage-external |
| sha | 247167c3ec73 |
| 所要時間 | 418s |
| 前回比 | +9459 (8.3%) vs 20261031-135709-3bc4182 |
| ベスト比 | +4234 (3.5%) vs 20261031-135233-2d10aad = 119317 **ベスト更新** |
| note | 画像の列で圧縮を試さない（main 14640b7 の上） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	134	383	898	897
europe	7	18	52	147	220	219
america	6	15	45	103	132	132
```

```
{"id":"16108","status":"done","portalStatus":"finished","passed":true,"score":123551,"raw":123551,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:11:05.471Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:11:08.442Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:11:11.443Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:11:14.464Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T05:11:17.442Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T05:11:20.443Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T05:11:23.463Z level=INFO msg="asia 3 日目: 134 人"
time=2026-10-31T05:11:24.613Z level=INFO msg="user moved" user=kuniyo5636 from=asia to=europe
time=2026-10-31T05:11:25.974Z level=INFO msg="user moved" user=kuniyo5636 from=europe to=america
time=2026-10-31T05:11:26.444Z level=INFO msg="europe 3 日目: 52 人"
time=2026-10-31T05:11:27.295Z level=INFO msg="user moved" user=nireke_6135 from=asia to=america
time=2026-10-31T05:11:29.443Z level=INFO msg="america 3 日目: 45 人"
time=2026-10-31T05:11:32.478Z level=INFO msg="asia 4 日目: 383 人"
time=2026-10-31T05:11:35.563Z level=INFO msg="europe 4 日目: 147 人"
time=2026-10-31T05:11:38.125Z level=INFO msg="user moved" user=mafu_7277 from=asia to=europe
time=2026-10-31T05:11:38.505Z level=INFO msg="america 4 日目: 103 人"
time=2026-10-31T05:11:38.531Z level=INFO msg="user moved" user=zoechimax.6818 from=asia to=europe
time=2026-10-31T05:11:38.639Z level=INFO msg="user moved" user=niyu7693 from=asia to=europe
time=2026-10-31T05:11:39.843Z level=INFO msg="user moved" user=memutor_7698 from=asia to=europe
time=2026-10-31T05:11:40.253Z level=INFO msg="user moved" user=kuniyo5636 from=america to=asia
time=2026-10-31T05:11:41.537Z level=INFO msg="asia 5 日目: 898 人"
time=2026-10-31T05:11:44.554Z level=INFO msg="europe 5 日目: 220 人"
time=2026-10-31T05:11:47.536Z level=INFO msg="america 5 日目: 132 人"
time=2026-10-31T05:11:56.574Z level=INFO msg="asia 最終日（グランプリ）: 897 人"
time=2026-10-31T05:11:56.626Z level=INFO msg="europe 最終日（グランプリ）: 219 人"
time=2026-10-31T05:11:56.699Z level=INFO msg="america 最終日（グランプリ）: 132 人"
time=2026-10-31T05:12:05.484Z level=INFO msg="グランプリ: アジア 415 人, ヨーロッパ 95 人, アメリカ 70 人"
{
  "pass": true,
  "score": 123551,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.2%	1436	0	0	GET	/api/notifications	12088.216	8.418	15.053	15.062	162
21.3%	2154	59	0	POST	/api/posts	7528.810	3.495	11.114	14.518	634
9.9%	4870	77	0	POST	/api/friend-requests	3510.525	0.721	5.752	9.611	114
8.3%	14127	388	0	GET	^/api/images/[^/]+	2940.993	0.208	3.648	8.125	129331
5.3%	6270	1793	0	POST	/api/auth/login	1885.150	0.301	4.244	10.799	182
4.5%	651	130	0	PUT	/api/me/profile-image	1599.880	2.458	10.352	11.658	4
3.5%	1854	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1245.598	0.672	5.642	8.090	116
2.0%	479	34	0	POST	^/api/posts/[^/]+/realmojis$	724.203	1.512	11.003	11.994	457
1.7%	919	1	0	POST	/api/auth/register	586.608	0.638	5.866	8.017	227
1.4%	3537	0	0	GET	/api/memories	508.466	0.144	1.045	3.828	11496
1.1%	557	17	0	POST	^/api/posts/[^/]+/comments$	400.837	0.720	5.863	9.255	326
1.1%	456	0	0	PATCH	/api/me	398.368	0.874	5.262	8.552	277
1.1%	625	0	0	PUT	^/api/follows/[^/]+$	394.531	0.631	5.239	7.962	0
0.8%	11290	0	0	GET	^/api/feeds/[^/]+	282.165	0.025	0.222	1.069	2732
0.6%	278	0	0	GET	/api/me/export	221.767	0.798	5.100	5.309	1930066
0.6%	704	0	0	POST	/api/auth/logout	208.150	0.296	4.850	7.030	0
0.5%	193	3	0	PATCH	^/api/posts/[^/]+$	176.660	0.915	6.540	6.603	642
0.4%	89	0	0	DELETE	^/api/posts/[^/]+$	133.082	1.495	5.911	5.911	0
0.3%	3160	0	0	GET	/api/posting-window/current	89.258	0.028	0.191	0.358	237
0.2%	1684	41	0	GET	^/api/posts/me/current$	57.231	0.034	0.376	1.205	635
0.2%	2366	0	0	GET	/api/users	54.112	0.023	0.171	3.473	232
0.1%	97	0	0	DELETE	^/api/comments/[^/]+$	50.676	0.522	4.308	4.308	0
0.1%	1492	0	0	GET	^/api/relationships/[^/]+	42.633	0.029	0.348	1.159	932
0.1%	671	0	0	GET	^/api/memories/[^/]+	40.625	0.061	0.566	1.560	663
0.1%	74	0	0	DELETE	^/api/realmojis/[^/]+$	40.021	0.541	3.061	3.061	0
0.1%	1749	13	0	GET	^/api/posts/[^/]+$	39.890	0.023	0.231	0.580	611
0.1%	37	0	0	PUT	^/api/blocks/[^/]+$	25.604	0.692	4.114	4.114	0
0.1%	1	0	0	POST	/api/initialize	20.528	20.528	20.528	20.528	90
0.1%	816	0	0	GET	^/api/users/[^/]+$	18.472	0.023	0.224	1.460	197
0.0%	783	0	0	GET	^/api/posts/[^/]+/realmojis$	15.012	0.019	0.136	0.357	1074
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.6%	433	0	0	GET	/api/notifications	3504.557	8.094	12.171	12.178	166
13.6%	567	15	0	POST	/api/posts	1266.199	2.233	7.491	7.893	630
9.9%	1154	20	0	POST	/api/friend-requests	922.089	0.799	5.824	9.965	114
6.0%	672	0	0	POST	/api/auth/login	558.308	0.831	5.325	9.507	245
5.8%	2305	14	0	GET	^/api/images/[^/]+	541.052	0.235	1.404	2.627	198551
4.7%	2204	0	0	GET	^/api/feeds/[^/]+	443.174	0.201	0.648	1.212	2770
3.2%	303	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	300.841	0.993	4.571	6.966	116
2.8%	804	0	0	GET	/api/posting-window/current	262.214	0.326	0.717	1.223	239
1.9%	645	0	0	GET	/api/memories	181.757	0.282	1.478	1.872	12196
1.6%	214	0	0	POST	/api/auth/register	153.399	0.717	3.841	4.388	228
1.5%	111	0	0	PUT	/api/me/profile-image	142.160	1.281	6.710	8.158	0
1.3%	546	0	0	GET	/api/users	121.628	0.223	0.635	0.777	228
1.2%	105	9	0	POST	^/api/posts/[^/]+/realmojis$	115.793	1.103	7.264	8.542	442
1.1%	455	16	0	GET	^/api/posts/me/current$	105.456	0.232	0.811	1.761	626
0.9%	105	0	0	POST	/api/auth/logout	85.605	0.815	4.173	4.642	0
0.8%	132	0	0	PUT	^/api/follows/[^/]+$	76.856	0.582	3.593	4.188	0
0.7%	33	0	0	GET	/api/me/export	67.133	2.034	6.936	6.936	2343066
0.7%	105	0	0	PATCH	/api/me	65.359	0.622	3.224	3.232	272
0.7%	329	0	0	GET	^/api/relationships/[^/]+	65.022	0.198	0.636	1.146	667
0.7%	325	9	0	GET	^/api/posts/[^/]+$	61.479	0.189	0.671	0.719	597
0.7%	99	1	0	POST	^/api/posts/[^/]+/comments$	61.439	0.621	4.758	4.758	312
0.6%	57	6	0	PATCH	^/api/posts/[^/]+$	56.943	0.999	3.884	3.884	587
0.3%	19	0	0	DELETE	^/api/posts/[^/]+$	31.639	1.665	5.589	5.589	0
0.3%	156	0	0	GET	^/api/users/[^/]+$	28.638	0.184	0.662	0.687	189
0.2%	133	1	0	GET	^/api/posts/[^/]+/realmojis$	22.394	0.168	0.580	0.604	1104
0.2%	13	0	0	DELETE	^/api/realmojis/[^/]+$	20.219	1.555	3.996	3.996	0
0.2%	80	0	0	GET	^/api/memories/[^/]+	19.831	0.248	0.940	0.940	892
0.2%	119	2	0	GET	^/api/posts/[^/]+/comments$	19.250	0.162	0.547	0.629	1092
0.2%	115	2	0	GET	/api/me	18.704	0.163	0.491	0.581	298
0.1%	17	0	0	DELETE	^/api/comments/[^/]+$	8.436	0.496	2.926	2.926	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.0%	266	0	0	GET	/api/notifications	1979.053	7.440	9.598	9.599	175
12.5%	761	11	0	POST	/api/friend-requests	709.182	0.932	5.449	7.323	114
12.4%	352	6	0	POST	/api/posts	701.895	1.994	6.264	7.146	638
6.5%	1485	12	0	GET	^/api/images/[^/]+	366.765	0.247	1.476	2.028	185246
5.9%	426	0	0	POST	/api/auth/login	335.869	0.788	4.741	7.319	248
4.5%	1232	0	0	GET	^/api/feeds/[^/]+	252.704	0.205	0.684	0.820	3067
3.0%	214	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	169.922	0.794	4.986	5.482	116
2.5%	482	0	0	GET	/api/posting-window/current	142.873	0.296	0.757	1.372	240
2.2%	449	0	0	GET	/api/memories	124.226	0.277	1.378	3.371	12533
1.8%	49	6	0	POST	^/api/posts/[^/]+/realmojis$	101.398	2.069	9.656	9.656	383
1.5%	363	0	0	GET	/api/users	85.762	0.236	0.729	0.971	229
1.3%	68	0	0	PUT	/api/me/profile-image	72.540	1.067	5.821	5.821	0
1.3%	82	2	0	POST	^/api/posts/[^/]+/comments$	72.349	0.882	3.753	3.753	362
1.3%	130	0	0	POST	/api/auth/register	71.556	0.550	3.325	3.553	230
1.1%	296	20	0	GET	^/api/posts/me/current$	64.340	0.217	0.735	1.240	606
1.0%	67	0	0	PUT	^/api/follows/[^/]+$	56.737	0.847	6.663	6.663	0
1.0%	69	0	0	POST	/api/auth/logout	56.582	0.820	4.756	4.756	0
1.0%	26	0	0	GET	/api/me/export	53.864	2.072	6.841	6.841	2920120
0.7%	57	0	0	PATCH	/api/me	42.388	0.744	4.250	4.250	281
0.6%	182	13	0	GET	^/api/posts/[^/]+$	36.458	0.200	1.005	1.278	576
0.5%	173	0	0	GET	^/api/relationships/[^/]+	30.716	0.178	0.663	0.721	1041
0.5%	27	0	0	PATCH	^/api/posts/[^/]+$	27.033	1.001	3.778	3.778	643
0.3%	15	0	0	DELETE	^/api/posts/[^/]+$	18.521	1.235	4.580	4.580	0
0.3%	71	0	0	GET	^/api/users/[^/]+$	14.577	0.205	0.659	0.659	200
0.3%	72	0	0	GET	/api/me	14.362	0.199	0.911	0.911	304
0.2%	54	0	0	GET	^/api/memories/[^/]+	11.665	0.216	0.654	0.654	1064
0.2%	63	0	0	GET	^/api/posts/[^/]+/realmojis$	11.371	0.180	0.897	0.897	1022
0.2%	60	0	0	GET	^/api/posts/[^/]+/comments$	11.035	0.184	0.506	0.506	1638
0.2%	8	0	0	DELETE	^/api/realmojis/[^/]+$	8.829	1.104	2.123	2.123	0
0.1%	11	0	0	DELETE	^/api/comments/[^/]+$	6.083	0.553	2.068	2.068	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 75823 クエリ / 82 種類 / 合計 203.608s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	50.3%	337	102.471	304.07	1216.20	1284.0	131.1	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	15.0%	3171	30.470	9.61	725.83	119.9	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
3	10.6%	9879	21.496	2.18	45.05	10.6	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	10.4%	4280	21.210	4.96	145.34	540.5	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
5	3.5%	2	7.166	3583.12	7015.71	243608.0	2870.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
6	2.9%	2995	5.821	1.94	10.55	24.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
7	1.1%	1	2.334	2333.73	2333.73	31281.0	1199932.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.9%	1	1.876	1876.20	1876.20	24949.0	628276.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.9%	584	1.777	3.04	15.30	28.6	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.7%	6665	1.519	0.23	16.05	19.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.6%	718	1.129	1.57	13.78	20.2	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.4%	1	0.818	817.74	817.74	31280.0	3165.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.4%	1	0.776	776.06	776.06	12929.0	530177.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
14	0.4%	5575	0.758	0.14	9.22	11.4	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
15	0.3%	2288	0.703	0.31	6.91	20.7	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
-- rank 3
SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: a87ca7136f9b10f678cde8c3c83a130c4c15da03
Type: cpu
Time: 2026-10-31 14:10:39 JST
Duration: 60.02s, Total samples = 33.48s (55.78%)
Showing nodes accounting for 23.59s, 70.46% of 33.48s total
Dropped 932 nodes (cum <= 0.17s)
Showing top 40 nodes out of 311
      flat  flat%   sum%        cum   cum%
     3.46s 10.33% 10.33%      3.46s 10.33%  internal/runtime/syscall/linux.Syscall6
     2.45s  7.32% 17.65%      2.45s  7.32%  crypto/internal/fips140/sha256.blockAVX2
     1.86s  5.56% 23.21%      5.82s 17.38%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.60s  4.78% 27.99%      1.60s  4.78%  runtime.memmove
     1.42s  4.24% 32.23%      1.84s  5.50%  runtime.findObject
     1.35s  4.03% 36.26%      1.35s  4.03%  runtime.memclrNoHeapPointers
     1.33s  3.97% 40.23%      5.23s 15.62%  runtime.scanObject
     1.28s  3.82% 44.06%      1.28s  3.82%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.18s  3.52% 47.58%      1.18s  3.52%  aeshashbody
     1.14s  3.41% 50.99%      1.94s  5.79%  runtime.tryDeferToSpanScan
     0.59s  1.76% 52.75%      0.59s  1.76%  runtime.spanClass.sizeclass (inline)
     0.48s  1.43% 54.18%      1.81s  5.41%  runtime.mapaccess1_faststr
     0.41s  1.22% 55.41%      0.41s  1.22%  memeqbody
     0.40s  1.19% 56.60%      0.40s  1.19%  runtime.futex
     0.35s  1.05% 57.65%      0.35s  1.05%  runtime.nextFreeFast (inline)
     0.30s   0.9% 58.54%      0.50s  1.49%  example.com/isureal-company/isureal.postLess
     0.25s  0.75% 59.29%      3.81s 11.38%  runtime.mallocgc
     0.24s  0.72% 60.01%      1.01s  3.02%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.72% 60.72%      0.24s  0.72%  sync/atomic.(*Int32).Add
     0.22s  0.66% 61.38%      0.69s  2.06%  database/sql.convertAssignRows
     0.20s   0.6% 61.98%      0.20s   0.6%  runtime.(*mspan).divideByElemSize (inline)
     0.19s  0.57% 62.54%      0.40s  1.19%  compress/flate.(*compressor).deflate
     0.19s  0.57% 63.11%      0.32s  0.96%  internal/runtime/maps.(*Iter).Next
     0.19s  0.57% 63.68%      1.73s  5.17%  runtime.mapassign_faststr
     0.18s  0.54% 64.22%      0.21s  0.63%  runtime.(*mspan).markBitsForIndex (inline)
     0.18s  0.54% 64.76%      0.18s  0.54%  runtime.typePointers.nextFast (inline)
     0.17s  0.51% 65.26%      0.22s  0.66%  encoding/json.appendString[go.shape.string]
     0.17s  0.51% 65.77%      0.20s   0.6%  example.com/isureal-company/isureal.rmLess
     0.16s  0.48% 66.25%      3.14s  9.38%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.16s  0.48% 66.73%      0.45s  1.34%  runtime.greyobject
     0.15s  0.45% 67.17%      0.22s  0.66%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.14s  0.42% 67.59%      1.34s  4.00%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.13s  0.39% 67.98%      2.22s  6.63%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.12s  0.36% 68.34%      0.32s  0.96%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.12s  0.36% 68.70%      0.17s  0.51%  runtime.(*mspan).writeHeapBitsSmall
     0.12s  0.36% 69.06%      0.17s  0.51%  runtime.heapArenaOf (inline)
     0.12s  0.36% 69.41%      0.74s  2.21%  runtime.mallocgcSmallNoscan
     0.12s  0.36% 69.77%      0.59s  1.76%  runtime.mapaccess2_faststr
     0.12s  0.36% 70.13%      0.17s  0.51%  runtime.spanOf (inline)
     0.11s  0.33% 70.46%      1.03s  3.08%  github.com/jackc/pgx/v5/pgproto3.(*chunkReader).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 50.0% / 100% | 34% | 0.4% | 10 | 187 MB | CPU が飽和 |
| isu2 (2 cpu) | 18.1% / 100% | 2% | 0.7% | 4 | 1778 MB |  |
| isu3 (2 cpu) | 15.3% / 100% | 2% | 0.6% | 3 | 1678 MB |  |
| isu4 (2 cpu) | 47.1% / 100% | 15% | 22.6% | 35 | 116 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 76% / nginx 11% |
| isu2 | isureal 34% / nginx 1% |
| isu3 | isureal 28% / nginx 1% |
| isu4 | postgres 35% / isureal 10% |

## Diff from previous run

スコア: 114092 → 123551 (+9459)

`3bc4182..247167c`

```
 conf/isu1/home/isucon/env.sh   |  2 +-
 webapp/go/jpegcheck/huffman.go | 27 +++------------------------
 webapp/go/jpegcheck/reader.go  | 32 ++++++++++++--------------------
 webapp/go/jpegcheck/scan.go    |  7 +++----
 webapp/go/offload.go           |  2 +-
 webapp/sql/reset.sql           |  3 +++
 6 files changed, 23 insertions(+), 50 deletions(-)
```
