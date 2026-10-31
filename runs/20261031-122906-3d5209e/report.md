# run 20261031-122906-3d5209e

## Summary

| | |
|---|---|
| score | 5222 |
| status | pass |
| ref | slow-before-1700 |
| sha | 3d5209e81a33 |
| 所要時間 | 166s |
| 前回比 | -61495 (-92.2%) vs 20261031-122401-fdbe826 |
| ベスト比 | -61495 (-92.2%) vs 20261031-122401-fdbe826 = 66717 |
| note | 17:00 直前用の低い点の版（試験 600ms） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	18	19	19	19	18
europe	7	7	7	7	7	6
america	6	6	6	6	6	6
```

```
{"id":"8298","status":"done","portalStatus":"finished","passed":true,"score":5222,"raw":5222,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:29:53.071Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:29:56.043Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:29:59.038Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:30:02.067Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T03:30:05.036Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T03:30:08.039Z level=INFO msg="america 2 日目: 6 人"
time=2026-10-31T03:30:11.073Z level=INFO msg="asia 3 日目: 19 人"
time=2026-10-31T03:30:14.036Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T03:30:17.044Z level=INFO msg="america 3 日目: 6 人"
time=2026-10-31T03:30:20.057Z level=INFO msg="asia 4 日目: 19 人"
time=2026-10-31T03:30:23.036Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T03:30:26.037Z level=INFO msg="america 4 日目: 6 人"
time=2026-10-31T03:30:29.056Z level=INFO msg="asia 5 日目: 19 人"
time=2026-10-31T03:30:32.044Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T03:30:35.034Z level=INFO msg="america 5 日目: 6 人"
time=2026-10-31T03:30:44.057Z level=INFO msg="asia 最終日（グランプリ）: 18 人"
time=2026-10-31T03:30:44.058Z level=INFO msg="america 最終日（グランプリ）: 6 人"
time=2026-10-31T03:30:44.061Z level=INFO msg="europe 最終日（グランプリ）: 6 人"
time=2026-10-31T03:30:52.430Z level=INFO msg="グランプリ: アジア 18 人, ヨーロッパ 6 人, アメリカ 5 人"
{
  "pass": true,
  "score": 5222,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
49.2%	128	0	0	GET	/api/notifications	2048.497	16.004	60.603	60.603	152
16.6%	1152	1015	0	POST	/api/auth/login	690.109	0.599	0.609	0.695	51
13.7%	948	484	0	GET	^/api/images/[^/]+	571.145	0.602	0.617	0.643	86033
3.7%	259	0	0	GET	^/api/feeds/[^/]+	155.713	0.601	0.612	0.617	1352
2.8%	181	0	0	GET	/api/memories	116.860	0.646	0.776	0.791	8201
2.0%	93	0	0	POST	/api/posts	81.321	0.874	1.191	1.191	604
1.5%	115	112	0	PUT	/api/me/profile-image	63.575	0.553	0.865	0.954	22
1.4%	99	0	0	GET	/api/posting-window/current	59.499	0.601	0.609	0.609	237
1.4%	94	1	0	GET	^/api/posts/me/current$	56.478	0.601	0.603	0.603	598
1.3%	92	11	0	POST	/api/friend-requests	55.445	0.603	0.607	0.607	105
1.0%	68	0	0	GET	^/api/relationships/[^/]+	40.990	0.603	0.613	0.613	3064
0.7%	51	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	30.854	0.605	0.621	0.621	116
0.7%	45	0	0	GET	/api/users	27.108	0.602	0.613	0.613	235
0.5%	37	3	0	GET	^/api/posts/[^/]+$	22.237	0.601	0.611	0.611	545
0.5%	35	1	0	POST	/api/auth/register	21.808	0.623	0.694	0.694	237
0.4%	31	0	0	POST	/api/auth/logout	18.679	0.603	0.606	0.606	0
0.4%	1	0	0	POST	/api/initialize	17.264	17.264	17.264	17.264	90
0.3%	12	0	0	GET	/api/me/export	13.568	1.131	1.817	1.817	2152003
0.3%	18	0	0	POST	/api/admin/posting-window	11.073	0.615	0.640	0.640	176
0.2%	14	0	0	GET	^/api/memories/[^/]+	8.815	0.630	0.702	0.702	540
0.2%	12	0	0	POST	^/api/posts/[^/]+/comments$	7.278	0.607	0.613	0.613	282
0.2%	12	0	0	GET	^/api/users/[^/]+$	7.221	0.602	0.609	0.609	164
0.2%	10	0	0	POST	^/api/posts/[^/]+/realmojis$	7.217	0.722	0.812	0.812	410
0.1%	10	0	0	GET	^/api/posts/[^/]+/comments$	6.023	0.602	0.615	0.615	232
0.1%	10	3	0	GET	/api/me	6.014	0.601	0.602	0.602	175
0.1%	8	0	0	PUT	^/api/follows/[^/]+$	4.848	0.606	0.610	0.610	0
0.1%	7	0	0	GET	^/api/posts/[^/]+/realmojis$	4.210	0.601	0.603	0.603	87
0.1%	5	0	0	DELETE	^/api/posts/[^/]+$	3.519	0.704	0.705	0.705	0
0.1%	5	0	0	PUT	^/api/admin/users/[^/]+/certification$	3.019	0.604	0.605	0.605	0
0.1%	4	0	0	PATCH	^/api/posts/[^/]+$	2.428	0.607	0.620	0.620	617
…（上位 30 / 全 37 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
65.9%	49	1	0	GET	/api/notifications	726.386	14.824	60.704	60.704	152
5.6%	82	0	0	GET	/api/memories	61.296	0.748	0.937	0.937	10763
5.2%	82	0	0	GET	^/api/feeds/[^/]+	57.633	0.703	0.719	0.719	1061
3.7%	197	3	0	GET	^/api/images/[^/]+	40.613	0.206	0.946	1.237	152105
3.6%	37	0	0	POST	/api/posts	39.870	1.078	1.559	1.559	617
2.7%	42	0	0	POST	/api/auth/login	30.144	0.718	0.811	0.811	251
2.4%	37	0	0	GET	/api/posting-window/current	26.291	0.711	0.804	0.804	238
2.4%	37	0	0	GET	^/api/posts/me/current$	25.967	0.702	0.705	0.705	617
2.1%	33	7	0	POST	/api/friend-requests	23.213	0.703	0.711	0.711	98
1.7%	26	0	0	GET	^/api/relationships/[^/]+	18.268	0.703	0.713	0.713	809
1.0%	15	0	0	GET	/api/users	10.564	0.704	0.713	0.713	223
0.8%	13	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	9.189	0.707	0.724	0.724	116
0.4%	6	0	0	PATCH	^/api/posts/[^/]+$	4.224	0.704	0.705	0.705	633
0.4%	5	0	0	POST	/api/auth/register	3.943	0.789	0.808	0.808	230
0.3%	5	2	0	GET	/api/me	3.611	0.722	0.804	0.804	127
0.3%	5	0	0	PUT	^/api/follows/[^/]+$	3.526	0.705	0.707	0.707	0
0.3%	5	0	0	POST	/api/auth/logout	3.515	0.703	0.704	0.704	0
0.3%	4	0	0	DELETE	^/api/posts/[^/]+$	3.230	0.807	0.811	0.811	0
0.3%	4	0	0	GET	^/api/memories/[^/]+	3.111	0.778	0.803	0.803	1180
0.2%	3	0	0	GET	^/api/posts/[^/]+$	2.106	0.702	0.703	0.703	577
0.2%	3	0	0	GET	^/api/posts/[^/]+/realmojis$	2.106	0.702	0.702	0.702	29
0.2%	1	0	0	GET	/api/me/export	1.890	1.890	1.890	1.890	2861228
0.1%	1	0	0	PUT	^/api/blocks/[^/]+$	0.807	0.807	0.807	0.807	0
0.1%	1	0	0	POST	/api/me/account-deletion	0.806	0.806	0.806	0.806	0
0.1%	1	0	0	GET	^/api/users/[^/]+$	0.702	0.702	0.702	0.702	168

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
62.6%	42	1	0	GET	/api/notifications	619.027	14.739	60.704	60.704	157
5.7%	75	0	0	GET	/api/memories	56.296	0.751	0.939	0.939	11945
5.7%	80	0	0	GET	^/api/feeds/[^/]+	56.228	0.703	0.718	0.718	1366
4.1%	37	0	0	POST	/api/posts	40.288	1.089	1.587	1.587	594
3.9%	182	2	0	GET	^/api/images/[^/]+	38.687	0.213	0.951	1.140	208451
2.7%	38	0	0	POST	/api/auth/login	27.159	0.715	0.811	0.811	256
2.7%	37	0	0	GET	/api/posting-window/current	26.469	0.715	0.804	0.804	240
2.6%	37	0	0	GET	^/api/posts/me/current$	25.971	0.702	0.710	0.710	594
2.0%	28	3	0	POST	/api/friend-requests	19.708	0.704	0.708	0.708	106
1.6%	22	0	0	GET	^/api/relationships/[^/]+	15.462	0.703	0.718	0.718	1109
1.2%	17	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	11.986	0.705	0.707	0.707	116
1.1%	15	0	0	GET	/api/users	10.544	0.703	0.704	0.704	212
0.8%	11	0	0	POST	/api/auth/logout	7.746	0.704	0.707	0.707	0
0.5%	7	0	0	GET	^/api/memories/[^/]+	5.013	0.716	0.804	0.804	568
0.4%	2	0	0	GET	/api/me/export	4.168	2.084	2.268	2.268	4161135
0.4%	5	0	0	POST	/api/auth/register	4.042	0.808	0.809	0.809	225
0.4%	5	0	0	PUT	^/api/follows/[^/]+$	3.530	0.706	0.709	0.709	0
0.4%	5	0	0	PATCH	^/api/posts/[^/]+$	3.518	0.704	0.705	0.705	597
0.3%	4	0	0	DELETE	^/api/posts/[^/]+$	3.219	0.805	0.805	0.805	0
0.3%	4	1	0	GET	/api/me	2.807	0.702	0.702	0.702	147
0.2%	3	0	0	GET	^/api/posts/[^/]+/realmojis$	2.104	0.701	0.702	0.702	29
0.1%	2	1	0	GET	^/api/posts/[^/]+$	1.404	0.702	0.702	0.702	302
0.1%	1	0	0	PUT	^/api/blocks/[^/]+$	0.807	0.807	0.807	0.807	0
0.1%	1	0	0	POST	/api/me/account-deletion	0.806	0.806	0.806	0.806	0
0.1%	1	0	0	GET	^/api/posts/[^/]+/comments$	0.703	0.703	0.703	0.703	294
0.1%	1	0	0	GET	^/api/users/[^/]+$	0.703	0.703	0.703	0.703	164
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 5057 クエリ / 76 種類 / 合計 24.377s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	27.0%	15	6.576	438.39	1039.90	7262.8	157.3	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	15.6%	339	3.803	11.22	206.61	118.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
3	14.5%	2	3.532	1765.80	3045.22	230251.5	2711.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
4	9.2%	1	2.231	2230.61	2230.61	29933.0	1199780.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
5	7.6%	1	1.854	1853.76	1853.76	21292.0	627267.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
6	6.2%	180	1.502	8.35	18.82	872.8	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
7	4.7%	2	1.149	574.40	1140.51	16770.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
8	3.1%	1	0.752	752.10	752.10	11620.0	529962.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	2.4%	1	0.593	593.44	593.44	25605.0	2156.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	2.2%	311	0.534	1.72	3.86	9.6	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
11	1.0%	167	0.248	1.48	4.71	27.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
12	1.0%	1	0.233	233.10	233.10	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
13	0.9%	1	0.221	220.71	220.71	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.8%	1	0.186	186.18	186.18	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.6%	1	0.144	144.26	144.26	30366.0	432.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Time: 2026-10-31 12:29:12 JST
Duration: 60s, Total samples = 14.94s (24.90%)
Showing nodes accounting for 12.12s, 81.12% of 14.94s total
Dropped 426 nodes (cum <= 0.07s)
Showing top 40 nodes out of 199
      flat  flat%   sum%        cum   cum%
     1.76s 11.78% 11.78%      5.72s 38.29%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.25s  8.37% 20.15%      1.25s  8.37%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.90s  6.02% 26.17%      0.90s  6.02%  aeshashbody
     0.87s  5.82% 31.99%      0.87s  5.82%  runtime.memclrNoHeapPointers
     0.66s  4.42% 36.41%      0.66s  4.42%  runtime.memmove
     0.59s  3.95% 40.36%      2.35s 15.73%  runtime.scanObject
     0.57s  3.82% 44.18%      0.74s  4.95%  runtime.findObject
     0.50s  3.35% 47.52%      0.87s  5.82%  runtime.tryDeferToSpanScan
     0.49s  3.28% 50.80%      0.49s  3.28%  internal/runtime/syscall/linux.Syscall6
     0.46s  3.08% 53.88%      0.63s  4.22%  database/sql.convertAssignRows
     0.39s  2.61% 56.49%      0.55s  3.68%  example.com/isureal-company/isureal.postLess
     0.39s  2.61% 59.10%      1.57s 10.51%  runtime.mapaccess1_faststr
     0.28s  1.87% 60.98%      0.28s  1.87%  runtime.spanClass.sizeclass (inline)
     0.27s  1.81% 62.78%      0.27s  1.81%  memeqbody
     0.23s  1.54% 64.32%      0.23s  1.54%  sync/atomic.(*Int32).Add (inline)
     0.21s  1.41% 65.73%      1.61s 10.78%  runtime.mallocgc
     0.19s  1.27% 67.00%      0.19s  1.27%  crypto/internal/fips140/sha256.blockAVX2
     0.19s  1.27% 68.27%      0.20s  1.34%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.19s  1.27% 69.54%      0.60s  4.02%  runtime.mallocgcSmallScanNoHeader
     0.18s  1.20% 70.75%      1.20s  8.03%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.17s  1.14% 71.89%      1.46s  9.77%  runtime.mapassign_faststr
     0.12s   0.8% 72.69%      0.12s   0.8%  runtime.nextFreeFast (inline)
     0.11s  0.74% 73.43%      0.13s  0.87%  example.com/isureal-company/isureal.rmLess
     0.11s  0.74% 74.16%      0.12s   0.8%  runtime.(*mspan).markBitsForIndex (inline)
     0.09s   0.6% 74.77%      0.09s   0.6%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.09s   0.6% 75.37%      0.10s  0.67%  runtime.spanOf (inline)
     0.09s   0.6% 75.97%      0.09s   0.6%  runtime.typePointers.nextFast (inline)
     0.08s  0.54% 76.51%      3.99s 26.71%  example.com/isureal-company/isureal.loadBig
     0.08s  0.54% 77.04%      2.01s 13.45%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.08s  0.54% 77.58%      0.08s  0.54%  internal/runtime/maps.(*ctrlGroup).setEmpty (inline)
     0.08s  0.54% 78.11%      0.21s  1.41%  runtime.greyobject
     0.06s   0.4% 78.51%      0.31s  2.07%  runtime.mallocgcSmallNoscan
     0.06s   0.4% 78.92%      0.28s  1.87%  runtime.mapaccess2_faststr
     0.05s  0.33% 79.25%      0.55s  3.68%  github.com/jackc/pgx/v5/pgtype.scanPlanString.Scan
     0.05s  0.33% 79.59%      0.18s  1.20%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func12
     0.05s  0.33% 79.92%      0.15s  1.00%  internal/runtime/maps.(*Iter).Next
     0.05s  0.33% 80.25%      0.08s  0.54%  runtime.heapArenaOf (inline)
     0.05s  0.33% 80.59%      0.17s  1.14%  runtime.scanObjectsSmall
     0.04s  0.27% 80.86%      0.17s  1.14%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.04s  0.27% 81.12%      0.45s  3.01%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 9.2% / 83% | 0% | 0.4% | 6 | 513 MB | run queue が積もっている |
| isu2 (2 cpu) | 1.7% / 20% | 0% | 0.0% | 3 | 2761 MB |  |
| isu3 (2 cpu) | 0.5% / 2% | 0% | 0.1% | 2 | 2818 MB |  |
| isu4 (2 cpu) | 7.6% / 83% | 0% | 4.7% | 1 | 437 MB |  |

**isu1（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 13% / nginx 4% |
| isu2 | isureal 3% |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 5% |

## Diff from previous run

スコア: 66717 → 5222 (-61495)

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
