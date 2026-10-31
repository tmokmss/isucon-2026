# run 20261031-123602-3d5209e

## Summary

| | |
|---|---|
| score | 5718 |
| status | pass |
| ref | slow-before-1700 |
| sha | 3d5209e81a33 |
| 所要時間 | 297s |
| 前回比 | -59319 (-91.2%) vs 20261031-123155-fdbe826 |
| ベスト比 | -60999 (-91.4%) vs 20261031-122401-fdbe826 = 66717 |
| note | 17:00 直前用の低い点の版（試験 600ms 2回目） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	18	18	18	18	17
europe	7	7	7	7	7	7
america	6	6	6	6	6	5
```

```
{"id":"8982","status":"done","portalStatus":"finished","passed":true,"score":5718,"raw":5718,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:38:57.347Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:39:00.314Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:39:03.313Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:39:06.343Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T03:39:09.313Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T03:39:12.325Z level=INFO msg="america 2 日目: 6 人"
time=2026-10-31T03:39:15.338Z level=INFO msg="asia 3 日目: 18 人"
time=2026-10-31T03:39:18.314Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T03:39:21.321Z level=INFO msg="america 3 日目: 6 人"
time=2026-10-31T03:39:24.333Z level=INFO msg="asia 4 日目: 18 人"
time=2026-10-31T03:39:27.316Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T03:39:30.314Z level=INFO msg="america 4 日目: 6 人"
time=2026-10-31T03:39:33.333Z level=INFO msg="asia 5 日目: 18 人"
time=2026-10-31T03:39:36.322Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T03:39:39.314Z level=INFO msg="america 5 日目: 6 人"
time=2026-10-31T03:39:48.314Z level=INFO msg="europe 最終日（グランプリ）: 7 人"
time=2026-10-31T03:39:48.337Z level=INFO msg="asia 最終日（グランプリ）: 17 人"
time=2026-10-31T03:39:48.339Z level=INFO msg="america 最終日（グランプリ）: 5 人"
time=2026-10-31T03:39:56.709Z level=INFO msg="グランプリ: アジア 17 人, ヨーロッパ 7 人, アメリカ 5 人"
{
  "pass": true,
  "score": 5718,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
46.8%	125	0	0	GET	/api/notifications	1967.578	15.741	60.603	60.603	153
17.9%	1250	1112	0	POST	/api/auth/login	751.656	0.601	0.625	0.696	50
13.7%	954	484	0	GET	^/api/images/[^/]+	574.327	0.602	0.613	0.618	98027
3.9%	272	0	0	GET	^/api/feeds/[^/]+	163.581	0.601	0.612	0.620	1820
2.9%	188	0	0	GET	/api/memories	120.804	0.643	0.736	0.805	9134
2.1%	102	0	0	POST	/api/posts	89.590	0.878	1.155	1.172	601
1.5%	106	11	0	POST	/api/friend-requests	63.938	0.603	0.610	0.630	107
1.5%	114	113	0	PUT	/api/me/profile-image	62.660	0.550	0.879	0.972	22
1.5%	103	1	0	GET	^/api/posts/me/current$	61.894	0.601	0.602	0.604	595
1.5%	102	0	0	GET	/api/posting-window/current	61.313	0.601	0.611	0.622	237
1.3%	88	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	53.187	0.604	0.617	0.617	116
1.1%	76	0	0	GET	^/api/relationships/[^/]+	45.800	0.603	0.612	0.612	3216
0.7%	48	0	0	GET	/api/users	28.888	0.602	0.604	0.604	255
0.5%	35	2	0	GET	^/api/posts/[^/]+$	21.058	0.602	0.613	0.613	568
0.5%	34	1	0	POST	/api/auth/register	21.006	0.618	0.684	0.684	237
0.4%	1	0	0	POST	/api/initialize	15.484	15.484	15.484	15.484	90
0.3%	21	0	0	POST	/api/auth/logout	12.660	0.603	0.604	0.604	0
0.3%	11	0	0	GET	/api/me/export	12.191	1.108	1.590	1.590	2330546
0.3%	19	0	0	GET	^/api/memories/[^/]+	11.724	0.617	0.707	0.707	295
0.3%	18	0	0	POST	/api/admin/posting-window	11.036	0.613	0.638	0.638	176
0.3%	18	0	0	GET	^/api/users/[^/]+$	10.817	0.601	0.602	0.602	164
0.2%	10	0	0	POST	^/api/posts/[^/]+/realmojis$	7.557	0.756	0.866	0.866	407
0.2%	11	0	0	PUT	^/api/follows/[^/]+$	6.656	0.605	0.609	0.609	0
0.1%	9	0	0	GET	^/api/posts/[^/]+/realmojis$	5.411	0.601	0.603	0.603	164
0.1%	8	3	0	GET	/api/me	4.811	0.601	0.602	0.602	143
0.1%	6	0	0	DELETE	^/api/posts/[^/]+$	4.239	0.706	0.720	0.720	0
0.1%	7	0	0	PATCH	^/api/posts/[^/]+$	4.220	0.603	0.604	0.604	604
0.1%	6	0	0	POST	^/api/posts/[^/]+/comments$	3.630	0.605	0.606	0.606	271
0.1%	5	0	0	PUT	^/api/admin/users/[^/]+/certification$	3.018	0.604	0.604	0.604	0
0.1%	5	0	0	GET	^/api/posts/[^/]+/comments$	3.004	0.601	0.601	0.601	137
…（上位 30 / 全 35 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
65.0%	49	0	0	GET	/api/notifications	788.806	16.098	60.703	60.703	155
6.0%	104	0	0	GET	^/api/feeds/[^/]+	73.012	0.702	0.708	0.714	1046
5.2%	85	0	0	GET	/api/memories	63.043	0.742	0.816	0.816	8359
3.4%	39	0	0	POST	/api/posts	41.833	1.073	1.456	1.456	601
3.2%	194	1	0	GET	^/api/images/[^/]+	39.010	0.201	0.941	0.942	162066
2.7%	45	0	0	POST	/api/auth/login	32.353	0.719	0.810	0.810	247
2.5%	43	0	0	GET	/api/posting-window/current	30.902	0.719	0.804	0.804	239
2.4%	41	6	0	POST	/api/friend-requests	28.917	0.705	0.739	0.739	103
2.3%	39	0	0	GET	^/api/posts/me/current$	27.383	0.702	0.712	0.712	601
1.7%	30	0	0	GET	^/api/relationships/[^/]+	21.057	0.702	0.705	0.705	1296
0.9%	16	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	11.293	0.706	0.717	0.717	116
0.9%	16	0	0	GET	/api/users	11.256	0.703	0.713	0.713	259
0.6%	9	0	0	POST	^/api/posts/[^/]+/realmojis$	7.554	0.839	0.913	0.913	409
0.5%	9	0	0	POST	/api/auth/logout	6.339	0.704	0.713	0.713	0
0.4%	3	0	0	GET	/api/me/export	5.282	1.761	2.263	2.263	2943740
0.3%	5	0	0	POST	/api/auth/register	3.938	0.788	0.806	0.806	229
0.3%	5	0	0	GET	^/api/memories/[^/]+	3.826	0.765	0.816	0.816	1029
0.3%	5	0	0	PATCH	^/api/posts/[^/]+$	3.539	0.708	0.725	0.725	592
0.2%	4	1	0	GET	/api/me	2.911	0.728	0.803	0.803	153
0.2%	4	0	0	PUT	^/api/follows/[^/]+$	2.826	0.706	0.708	0.708	0
0.2%	4	0	0	GET	^/api/posts/[^/]+$	2.806	0.702	0.702	0.702	627
0.1%	2	0	0	GET	^/api/posts/[^/]+/realmojis$	1.403	0.702	0.702	0.702	232
0.1%	1	0	0	DELETE	^/api/posts/[^/]+$	0.831	0.831	0.831	0.831	0
0.1%	1	0	0	DELETE	^/api/realmojis/[^/]+$	0.805	0.805	0.805	0.805	0
0.1%	1	0	0	PUT	^/api/blocks/[^/]+$	0.708	0.708	0.708	0.708	0
0.1%	1	0	0	GET	^/api/posts/[^/]+/comments$	0.702	0.702	0.702	0.702	29
0.1%	1	0	0	GET	^/api/users/[^/]+$	0.702	0.702	0.702	0.702	165

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
60.7%	41	0	0	GET	/api/notifications	605.428	14.767	60.703	60.703	146
5.5%	74	0	0	GET	/api/memories	55.160	0.745	0.940	0.940	11278
5.4%	77	0	0	GET	^/api/feeds/[^/]+	54.105	0.703	0.741	0.741	1365
4.4%	39	0	0	POST	/api/posts	43.485	1.115	1.589	1.589	598
3.9%	181	4	0	GET	^/api/images/[^/]+	38.505	0.213	0.951	0.951	186011
3.3%	46	0	0	POST	/api/auth/login	32.771	0.712	0.809	0.809	254
3.2%	44	0	0	GET	/api/posting-window/current	31.534	0.717	0.812	0.812	240
2.9%	41	5	0	POST	/api/friend-requests	28.877	0.704	0.727	0.727	105
2.7%	39	1	0	GET	^/api/posts/me/current$	27.379	0.702	0.707	0.707	584
1.3%	19	0	0	GET	^/api/relationships/[^/]+	13.339	0.702	0.706	0.706	1589
1.1%	15	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	10.578	0.705	0.710	0.710	116
1.0%	14	0	0	GET	/api/users	9.846	0.703	0.708	0.708	263
0.7%	4	0	0	GET	/api/me/export	7.258	1.815	2.236	2.236	3474121
0.6%	8	0	0	POST	/api/auth/logout	5.737	0.717	0.805	0.805	0
0.5%	7	0	0	GET	^/api/memories/[^/]+	5.352	0.765	0.824	0.824	1088
0.4%	5	0	0	POST	/api/auth/register	4.034	0.807	0.807	0.807	230
0.4%	5	1	0	GET	/api/me	3.511	0.702	0.703	0.703	183
0.3%	4	0	0	DELETE	^/api/posts/[^/]+$	3.249	0.812	0.821	0.821	0
0.3%	4	0	0	PUT	^/api/follows/[^/]+$	2.823	0.706	0.709	0.709	0
0.3%	4	0	0	GET	^/api/users/[^/]+$	2.806	0.702	0.702	0.702	186
0.3%	4	1	0	GET	^/api/posts/[^/]+$	2.805	0.701	0.702	0.702	454
0.2%	3	0	0	POST	^/api/posts/[^/]+/comments$	2.130	0.710	0.715	0.715	274
0.2%	3	0	0	GET	^/api/posts/[^/]+/comments$	2.105	0.702	0.702	0.702	120
0.1%	2	0	0	PATCH	^/api/posts/[^/]+$	1.407	0.704	0.704	0.704	643
0.1%	1	0	0	POST	/api/me/account-deletion	0.808	0.808	0.808	0.808	0
0.1%	1	0	0	PUT	/api/me/profile-image	0.806	0.806	0.806	0.806	0
0.1%	1	0	0	PUT	^/api/blocks/[^/]+$	0.706	0.706	0.706	0.706	0
0.1%	1	0	0	GET	^/api/posts/[^/]+/realmojis$	0.703	0.703	0.703	0.703	29
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 5371 クエリ / 75 種類 / 合計 23.087s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	31.2%	18	7.205	400.27	832.35	7798.9	169.9	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	14.6%	339	3.364	9.92	207.07	117.8	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
3	12.4%	2	2.874	1436.92	2720.51	232195.0	2732.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
4	9.8%	1	2.265	2265.27	2265.27	30014.0	1199679.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
5	7.5%	1	1.735	1734.90	1734.90	21549.0	626833.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
6	6.9%	201	1.595	7.94	19.89	825.4	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
7	3.3%	1	0.773	772.96	772.96	11708.0	529917.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	2.6%	2	0.595	297.35	586.58	16790.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
9	2.1%	289	0.492	1.70	4.39	9.8	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
10	1.4%	1	0.330	329.94	329.94	24994.0	1722.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	1.1%	1	0.257	257.29	257.29	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
12	1.0%	180	0.237	1.31	4.82	27.2	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
13	1.0%	1	0.224	224.46	224.46	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.8%	1	0.182	182.11	182.11	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.6%	1	0.142	142.11	142.11	30346.0	331.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
-- rank 3
SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: e0fbe50e925ae32fc3de3445ab2c4187bc62dcd5
Type: cpu
Time: 2026-10-31 12:38:18 JST
Duration: 60s, Total samples = 15.22s (25.37%)
Showing nodes accounting for 12.26s, 80.55% of 15.22s total
Dropped 432 nodes (cum <= 0.08s)
Showing top 40 nodes out of 200
      flat  flat%   sum%        cum   cum%
     1.72s 11.30% 11.30%      5.75s 37.78%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.22s  8.02% 19.32%      1.22s  8.02%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.99s  6.50% 25.82%      0.99s  6.50%  aeshashbody
     0.96s  6.31% 32.13%      0.96s  6.31%  runtime.memclrNoHeapPointers
     0.67s  4.40% 36.53%      0.67s  4.40%  runtime.memmove
     0.63s  4.14% 40.67%      2.37s 15.57%  runtime.scanObject
     0.62s  4.07% 44.74%      0.62s  4.07%  internal/runtime/syscall/linux.Syscall6
     0.60s  3.94% 48.69%      0.84s  5.52%  runtime.findObject
     0.44s  2.89% 51.58%      0.74s  4.86%  runtime.tryDeferToSpanScan
     0.40s  2.63% 54.20%      1.52s  9.99%  runtime.mapaccess1_faststr
     0.36s  2.37% 56.57%      0.52s  3.42%  example.com/isureal-company/isureal.postLess
     0.34s  2.23% 58.80%      0.45s  2.96%  database/sql.convertAssignRows
     0.24s  1.58% 60.38%      0.24s  1.58%  memeqbody
     0.23s  1.51% 61.89%      0.77s  5.06%  runtime.mallocgcSmallScanNoHeader
     0.22s  1.45% 63.34%      0.22s  1.45%  runtime.spanClass.sizeclass (inline)
     0.19s  1.25% 64.59%      0.19s  1.25%  example.com/isureal-company/isureal.rmLess
     0.18s  1.18% 65.77%      0.18s  1.18%  runtime.nextFreeFast (inline)
     0.16s  1.05% 66.82%      2.21s 14.52%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.15s  0.99% 67.81%      0.19s  1.25%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.13s  0.85% 68.66%      0.13s  0.85%  sync/atomic.(*Int32).Add (inline)
     0.12s  0.79% 69.45%      0.12s  0.79%  runtime.typePointers.nextFast (inline)
     0.11s  0.72% 70.17%      0.11s  0.72%  crypto/internal/fips140/sha256.blockAVX2
     0.11s  0.72% 70.89%      1.09s  7.16%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.11s  0.72% 71.62%      0.11s  0.72%  runtime.(*mspan).divideByElemSize (inline)
     0.10s  0.66% 72.27%      0.13s  0.85%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
     0.10s  0.66% 72.93%      0.10s  0.66%  runtime.(*spanInlineMarkBits).init
     0.10s  0.66% 73.59%      1.75s 11.50%  runtime.mallocgc
     0.10s  0.66% 74.24%      0.33s  2.17%  runtime.mapaccess2_faststr
     0.09s  0.59% 74.84%         4s 26.28%  example.com/isureal-company/isureal.loadBig
     0.09s  0.59% 75.43%      0.48s  3.15%  github.com/jackc/pgx/v5/pgtype.scanPlanString.Scan
     0.09s  0.59% 76.02%      0.11s  0.72%  runtime.spanOf (inline)
     0.08s  0.53% 76.54%      0.27s  1.77%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.08s  0.53% 77.07%      0.08s  0.53%  example.com/isureal-company/isureal.(*UserRow).active
     0.08s  0.53% 77.60%      0.16s  1.05%  internal/reflectlite.Swapper.func3
     0.08s  0.53% 78.12%      0.08s  0.53%  internal/runtime/maps.(*ctrlGroup).setEmpty (inline)
     0.08s  0.53% 78.65%      0.12s  0.79%  runtime.(*mspan).markBitsForIndex (inline)
     0.08s  0.53% 79.17%      0.12s  0.79%  runtime.(*mspan).writeHeapBitsSmall
     0.07s  0.46% 79.63%      0.52s  3.42%  database/sql.(*Rows).scanLocked
     0.07s  0.46% 80.09%      1.33s  8.74%  runtime.mapassign_faststr
     0.07s  0.46% 80.55%      0.65s  4.27%  sort.insertionSort_func
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 9.2% / 88% | 0% | 0.4% | 4 | 475 MB |  |
| isu2 (2 cpu) | 1.8% / 17% | 0% | 0.1% | 2 | 2705 MB |  |
| isu3 (2 cpu) | 0.5% / 3% | 0% | 0.1% | 2 | 2762 MB |  |
| isu4 (2 cpu) | 7.4% / 60% | 0% | 4.5% | 2 | 131 MB |  |

どのホストも飽和していない。**ボトルネックはアプリのロジックか DB のクエリ側**にある可能性が高い。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 13% / nginx 4% |
| isu2 | isureal 3% |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 5% |

## Diff from previous run

スコア: 65037 → 5718 (-59319)

`fdbe826..3d5209e`

```
 webapp/go/account_deletion_api.go |   2 +-
 webapp/go/admin_user_api.go       |   2 +-
 webapp/go/comment_api.go          |   2 +-
 webapp/go/realmoji_api.go         |   2 +-
 webapp/go/relationship.go         |  10 +---
 webapp/go/server.go               |   1 +
 webapp/go/slow.go                 |  21 ++++++++
 webapp/go/store.go                | 100 ++++++--------------------------------
 8 files changed, 44 insertions(+), 96 deletions(-)
```
