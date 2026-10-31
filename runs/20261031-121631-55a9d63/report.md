# run 20261031-121631-55a9d63

## Summary

| | |
|---|---|
| score | 29539 |
| status | pass |
| ref | slow-before-1700 |
| sha | 55a9d638ae23 |
| 所要時間 | 241s |
| 前回比 | +29539 (—) vs 20261031-120912-249be25 |
| ベスト比 | -30662 (-50.9%) vs 20261031-115040-1c73ee4 = 60201 |
| note | 17:00 直前用の低い点の版（試験 300ms） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	28	52	101	189	188
europe	7	7	7	8	10	9
america	6	7	9	12	15	14
```

```
{"id":"7510","status":"done","portalStatus":"finished","passed":true,"score":29539,"raw":29539,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:18:44.131Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:18:47.081Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:18:50.078Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:18:53.109Z level=INFO msg="asia 2 日目: 28 人"
time=2026-10-31T03:18:56.079Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T03:18:59.078Z level=INFO msg="america 2 日目: 7 人"
time=2026-10-31T03:19:02.100Z level=INFO msg="asia 3 日目: 52 人"
time=2026-10-31T03:19:05.081Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T03:19:08.080Z level=INFO msg="america 3 日目: 9 人"
time=2026-10-31T03:19:11.100Z level=INFO msg="asia 4 日目: 101 人"
time=2026-10-31T03:19:14.077Z level=INFO msg="europe 4 日目: 8 人"
time=2026-10-31T03:19:17.085Z level=INFO msg="america 4 日目: 12 人"
time=2026-10-31T03:19:20.100Z level=INFO msg="asia 5 日目: 189 人"
time=2026-10-31T03:19:23.083Z level=INFO msg="europe 5 日目: 10 人"
time=2026-10-31T03:19:26.079Z level=INFO msg="america 5 日目: 15 人"
time=2026-10-31T03:19:35.078Z level=INFO msg="america 最終日（グランプリ）: 14 人"
time=2026-10-31T03:19:35.104Z level=INFO msg="asia 最終日（グランプリ）: 188 人"
time=2026-10-31T03:19:35.112Z level=INFO msg="europe 最終日（グランプリ）: 9 人"
time=2026-10-31T03:19:43.775Z level=INFO msg="グランプリ: アジア 187 人, ヨーロッパ 9 人, アメリカ 14 人"
{
  "pass": true,
  "score": 29539,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.9%	411	0	0	GET	/api/notifications	3827.486	9.313	15.006	15.006	175
12.6%	3764	416	0	GET	^/api/images/[^/]+	1236.625	0.329	1.045	3.390	171391
8.8%	2852	0	0	GET	^/api/feeds/[^/]+	867.057	0.304	0.382	0.642	3260
6.4%	571	0	0	POST	/api/posts	634.125	1.111	3.758	4.163	645
6.2%	1923	986	0	POST	/api/auth/login	613.022	0.319	0.674	3.323	132
4.3%	1221	0	0	GET	/api/memories	427.382	0.350	0.544	0.935	7921
3.5%	1144	31	0	POST	/api/friend-requests	349.466	0.305	0.345	0.703	113
2.6%	752	0	0	GET	/api/posting-window/current	255.794	0.340	0.770	3.347	237
2.1%	642	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	205.205	0.320	0.381	3.132	115
1.8%	571	2	0	GET	^/api/posts/me/current$	174.677	0.306	0.515	0.705	643
1.7%	545	0	0	GET	^/api/relationships/[^/]+	164.344	0.302	0.317	0.386	1597
1.6%	528	0	0	GET	/api/users	159.790	0.303	0.316	0.364	235
1.4%	467	30	0	GET	^/api/posts/[^/]+$	140.635	0.301	0.311	0.322	585
0.8%	198	97	0	PUT	/api/me/profile-image	82.178	0.415	1.181	1.759	12
0.7%	242	0	0	GET	^/api/users/[^/]+$	72.979	0.302	0.319	0.332	197
0.7%	239	0	0	POST	^/api/posts/[^/]+/comments$	72.841	0.305	0.329	0.386	360
0.7%	164	0	0	POST	^/api/posts/[^/]+/realmojis$	70.401	0.429	0.527	0.537	481
0.7%	205	1	0	POST	/api/auth/register	65.378	0.319	0.389	2.809	228
0.6%	196	0	0	GET	^/api/posts/[^/]+/comments$	59.024	0.301	0.312	0.320	1085
0.6%	188	0	0	GET	^/api/posts/[^/]+/realmojis$	56.685	0.302	0.331	0.333	585
0.6%	156	0	0	POST	/api/auth/logout	54.283	0.348	2.185	2.343	0
0.5%	139	0	0	GET	^/api/memories/[^/]+	47.417	0.341	0.454	0.638	622
0.4%	58	0	0	GET	/api/me/export	42.603	0.735	1.599	1.599	1989980
0.3%	109	3	0	GET	/api/me	32.904	0.302	0.324	0.325	292
0.3%	106	0	0	PUT	^/api/follows/[^/]+$	32.409	0.306	0.330	0.395	0
0.2%	75	0	0	PATCH	/api/me	22.828	0.304	0.336	0.336	267
0.2%	62	2	0	PATCH	^/api/posts/[^/]+$	19.137	0.309	0.620	0.620	620
0.1%	1	0	0	POST	/api/initialize	13.417	13.417	13.417	13.417	90
0.1%	30	0	0	DELETE	^/api/posts/[^/]+$	13.096	0.437	0.778	0.778	0
0.1%	32	0	0	DELETE	^/api/comments/[^/]+$	9.694	0.303	0.307	0.307	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.8%	46	0	0	GET	/api/notifications	395.673	8.602	12.024	12.024	177
10.6%	236	0	0	GET	^/api/feeds/[^/]+	95.656	0.405	0.444	0.453	2564
6.8%	360	2	0	GET	^/api/images/[^/]+	61.105	0.170	0.932	1.029	190428
6.2%	50	0	0	POST	/api/posts	56.054	1.121	3.925	3.925	616
5.6%	110	0	0	GET	/api/memories	50.888	0.463	0.650	0.729	9910
4.2%	79	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	37.943	0.480	2.571	2.571	115
3.8%	73	0	0	POST	/api/auth/login	34.721	0.476	2.548	2.548	255
2.7%	53	0	0	GET	/api/posting-window/current	23.942	0.452	0.949	0.949	239
2.5%	56	3	0	POST	/api/friend-requests	22.749	0.406	0.426	0.426	111
2.3%	50	0	0	GET	^/api/posts/me/current$	20.338	0.407	0.464	0.464	616
1.7%	37	0	0	GET	^/api/relationships/[^/]+	14.939	0.404	0.436	0.436	3021
1.2%	27	4	0	GET	^/api/posts/[^/]+$	10.921	0.404	0.451	0.451	533
1.1%	24	0	0	GET	/api/users	9.707	0.404	0.439	0.439	271
1.0%	22	0	0	GET	^/api/users/[^/]+$	8.877	0.403	0.423	0.423	191
0.8%	17	0	0	POST	/api/auth/logout	6.887	0.405	0.419	0.419	0
0.8%	4	0	0	GET	/api/me/export	6.811	1.703	2.473	2.473	3444018
0.7%	15	0	0	GET	^/api/memories/[^/]+	6.711	0.447	0.524	0.524	666
0.7%	16	0	0	GET	^/api/posts/[^/]+/realmojis$	6.462	0.404	0.419	0.419	563
0.6%	14	0	0	GET	^/api/posts/[^/]+/comments$	5.635	0.403	0.409	0.409	292
0.6%	10	0	0	POST	^/api/posts/[^/]+/realmojis$	5.225	0.522	0.608	0.608	465
0.4%	8	0	0	POST	/api/auth/register	3.665	0.458	0.508	0.508	227
0.4%	9	0	0	PUT	^/api/follows/[^/]+$	3.640	0.404	0.408	0.408	0
0.4%	8	0	0	POST	^/api/posts/[^/]+/comments$	3.289	0.411	0.436	0.436	270
0.3%	7	0	0	PATCH	^/api/posts/[^/]+$	2.934	0.419	0.465	0.465	612
0.3%	6	2	0	GET	/api/me	2.539	0.423	0.503	0.503	157
0.2%	3	0	0	DELETE	^/api/posts/[^/]+$	1.526	0.509	0.515	0.515	0
0.1%	2	0	0	DELETE	^/api/realmojis/[^/]+$	1.024	0.512	0.515	0.515	0
0.1%	2	0	0	PUT	^/api/blocks/[^/]+$	0.818	0.409	0.410	0.410	0
0.1%	1	0	0	POST	/api/me/account-deletion	0.515	0.515	0.515	0.515	0
0.1%	1	0	0	PUT	/api/me/profile-image	0.509	0.509	0.509	0.509	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.9%	55	0	0	GET	/api/notifications	438.419	7.971	9.001	9.001	179
7.5%	182	0	0	GET	^/api/feeds/[^/]+	74.473	0.409	0.819	0.819	2651
7.3%	399	1	0	GET	^/api/images/[^/]+	73.317	0.184	0.951	1.237	165396
7.1%	64	0	0	POST	/api/posts	71.225	1.113	3.672	3.672	627
6.3%	137	0	0	GET	/api/memories	63.020	0.460	0.684	0.753	13427
4.2%	94	0	0	POST	/api/auth/login	42.243	0.449	2.273	2.273	257
4.1%	102	4	0	POST	/api/friend-requests	41.434	0.406	0.462	0.505	112
3.2%	73	0	0	GET	/api/posting-window/current	31.933	0.437	0.716	0.716	240
2.6%	64	1	0	GET	^/api/posts/me/current$	25.917	0.405	0.474	0.474	617
2.1%	52	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	21.397	0.411	0.521	0.521	116
2.0%	48	0	0	GET	/api/users	19.543	0.407	0.503	0.503	247
1.7%	43	0	0	GET	^/api/relationships/[^/]+	17.359	0.404	0.464	0.464	1868
1.2%	8	0	0	GET	/api/me/export	12.339	1.542	2.662	2.662	3589812
0.9%	17	1	0	POST	^/api/posts/[^/]+/realmojis$	9.278	0.546	0.624	0.624	402
0.7%	17	1	0	GET	^/api/posts/[^/]+$	6.983	0.411	0.503	0.503	600
0.7%	15	0	0	GET	^/api/memories/[^/]+	6.610	0.441	0.510	0.510	520
0.6%	14	0	0	POST	/api/auth/register	6.485	0.463	0.509	0.509	230
0.5%	13	0	0	POST	/api/auth/logout	5.261	0.405	0.412	0.412	0
0.5%	12	0	0	PUT	^/api/follows/[^/]+$	4.865	0.405	0.416	0.416	0
0.4%	10	1	0	GET	/api/me	4.022	0.402	0.403	0.403	246
0.4%	10	0	0	GET	^/api/users/[^/]+$	4.019	0.402	0.403	0.403	160
0.4%	10	1	0	GET	^/api/posts/[^/]+/realmojis$	4.019	0.402	0.403	0.403	477
0.3%	6	0	0	PUT	/api/me/profile-image	3.093	0.515	0.552	0.552	0
0.3%	7	0	0	GET	^/api/posts/[^/]+/comments$	2.815	0.402	0.404	0.404	549
0.2%	5	0	0	PATCH	/api/me	2.063	0.413	0.440	0.440	288
0.2%	5	0	0	PATCH	^/api/posts/[^/]+$	2.039	0.408	0.422	0.422	604
0.2%	5	2	0	POST	^/api/posts/[^/]+/comments$	2.023	0.405	0.409	0.409	179
0.1%	2	0	0	DELETE	^/api/realmojis/[^/]+$	1.015	0.508	0.508	0.508	0
0.1%	1	0	0	DELETE	^/api/posts/[^/]+$	0.505	0.505	0.505	0.505	0
0.1%	1	0	0	POST	/api/me/account-deletion	0.505	0.505	0.505	0.505	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 18578 クエリ / 80 種類 / 合計 81.191s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	49.2%	983	39.927	40.62	440.51	667.7	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	27.3%	70	22.127	316.10	1057.42	7086.5	143.8	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.7%	659	4.649	7.05	212.28	123.6	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.6%	2016	3.703	1.84	23.89	11.2	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	2.7%	1	2.202	2201.84	2201.84	29821.0	1199783.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	2.1%	1	1.716	1716.29	1716.29	20990.0	626926.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	1.8%	685	1.428	2.08	22.24	30.5	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
8	1.7%	2	1.383	691.60	1234.43	225505.5	2654.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
9	0.9%	1	0.732	732.49	732.49	11529.0	530079.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
10	0.6%	190	0.470	2.47	6.44	30.6	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
11	0.4%	2	0.320	159.88	311.97	16794.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.4%	250	0.301	1.21	15.60	22.7	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
13	0.3%	1	0.230	230.30	230.30	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
14	0.2%	1	0.200	200.31	200.31	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
15	0.2%	1	0.185	184.73	184.73	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: c3a55a6f1ff851c82dba5d4440dcb93eb04299df
Type: cpu
Time: 2026-10-31 12:18:15 JST
Duration: 60s, Total samples = 18.30s (30.50%)
Showing nodes accounting for 14.55s, 79.51% of 18.30s total
Dropped 571 nodes (cum <= 0.09s)
Showing top 40 nodes out of 239
      flat  flat%   sum%        cum   cum%
     1.67s  9.13%  9.13%      5.74s 31.37%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.34s  7.32% 16.45%      1.34s  7.32%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.16s  6.34% 22.79%      1.16s  6.34%  runtime.memclrNoHeapPointers
     1.14s  6.23% 29.02%      1.14s  6.23%  internal/runtime/syscall/linux.Syscall6
        1s  5.46% 34.48%         1s  5.46%  aeshashbody
     0.94s  5.14% 39.62%      0.94s  5.14%  runtime.memmove
     0.93s  5.08% 44.70%      3.21s 17.54%  runtime.scanObject
     0.76s  4.15% 48.85%      1.06s  5.79%  runtime.findObject
     0.62s  3.39% 52.24%      0.94s  5.14%  runtime.tryDeferToSpanScan
     0.46s  2.51% 54.75%      1.75s  9.56%  runtime.mapaccess1_faststr
     0.33s  1.80% 56.56%      0.53s  2.90%  database/sql.convertAssignRows
     0.31s  1.69% 58.25%      0.31s  1.69%  memeqbody
     0.30s  1.64% 59.89%      0.53s  2.90%  example.com/isureal-company/isureal.postLess
     0.28s  1.53% 61.42%      0.28s  1.53%  crypto/internal/fips140/sha256.blockAVX2
     0.25s  1.37% 62.79%      0.25s  1.37%  runtime.spanClass.sizeclass (inline)
     0.24s  1.31% 64.10%      0.24s  1.31%  runtime.nextFreeFast (inline)
     0.23s  1.26% 65.36%      0.78s  4.26%  runtime.mallocgcSmallScanNoHeader
     0.18s  0.98% 66.34%      0.18s  0.98%  example.com/isureal-company/isureal.(*UserRow).active
     0.18s  0.98% 67.32%      1.45s  7.92%  runtime.mapassign_faststr
     0.18s  0.98% 68.31%      0.18s  0.98%  sync/atomic.(*Int32).Add
     0.16s  0.87% 69.18%      2.09s 11.42%  runtime.mallocgc
     0.14s  0.77% 69.95%      0.19s  1.04%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.13s  0.71% 70.66%      0.15s  0.82%  example.com/isureal-company/isureal.rmLess
     0.13s  0.71% 71.37%      0.13s  0.71%  runtime.typePointers.nextFast (inline)
     0.11s   0.6% 71.97%      0.11s   0.6%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.11s   0.6% 72.57%      0.11s   0.6%  runtime.(*mspan).divideByElemSize (inline)
     0.11s   0.6% 73.17%      0.16s  0.87%  runtime.spanOf (inline)
     0.10s  0.55% 73.72%      0.10s  0.55%  cmpbody
     0.10s  0.55% 74.26%      2.16s 11.80%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.55% 74.81%      0.17s  0.93%  internal/runtime/maps.(*Iter).Next
     0.10s  0.55% 75.36%      0.10s  0.55%  internal/runtime/maps.(*ctrlGroup).setEmpty (inline)
     0.10s  0.55% 75.90%      0.11s   0.6%  runtime.(*mspan).markBitsForIndex (inline)
     0.10s  0.55% 76.45%      0.12s  0.66%  runtime.(*mspan).writeHeapBitsSmall
     0.10s  0.55% 76.99%      0.27s  1.48%  runtime.greyobject
     0.10s  0.55% 77.54%      0.10s  0.55%  runtime.typePointers.next
     0.09s  0.49% 78.03%      0.10s  0.55%  time.Time.Before
     0.08s  0.44% 78.47%      1.24s  6.78%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.07s  0.38% 78.85%      0.52s  2.84%  runtime.slicebytetostring
     0.06s  0.33% 79.18%      0.59s  3.22%  database/sql.(*Rows).scanLocked
     0.06s  0.33% 79.51%      0.12s  0.66%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 18.0% / 100% | 1% | 0.2% | 4 | 514 MB |  |
| isu2 (2 cpu) | 5.5% / 29% | 0% | 0.1% | 4 | 2805 MB |  |
| isu3 (2 cpu) | 1.8% / 17% | 0% | 0.3% | 3 | 2911 MB |  |
| isu4 (2 cpu) | 17.6% / 97% | 1% | 9.7% | 27 | 114 MB | run queue が積もっている |

**isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 27% / nginx 6% |
| isu2 | isureal 11% |
| isu3 | isureal 3% |
| isu4 | postgres 14% |

## Diff from previous run

スコア: 0 → 29539 (+29539)

`249be25..55a9d63`

```
 .../etc/postgresql/16/main/conf.d/zz-tuning.conf    |  2 ++
 webapp/go/server.go                                 |  1 +
 webapp/go/slow.go                                   | 21 +++++++++++++++++++++
 3 files changed, 24 insertions(+)
```
