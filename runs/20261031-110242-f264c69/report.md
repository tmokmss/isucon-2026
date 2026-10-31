# run 20261031-110242-f264c69

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | HEAD |
| sha | f264c69ca06c |
| 所要時間 | 76s |
| 前回比 | -10532 (-100.0%) vs 20261031-105911-5aad88a |
| ベスト比 | -24801 (-100.0%) vs 20261031-103811-cef5fe3 = 24801 |
| note | 署名: exp を 10 分単位に（main + 1 コミット） |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/contest-day-hosts/runs/20261031-110242-f264c69/bench.log
```

## Bench

```
{"id":"2454","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"初期データの検証に失敗しました"}

--- stdout ---
{
  "pass": false,
  "score": 0,
  "messages": [
    "初期データの検証に失敗しました: 初期データのユーザー panya_shugyo_386: Memories（2025-11）: 画像 URL \"/api/images/64237636-7c5c-4cbf-a9f0-03429cdd82b6?exp=1793499000\u0026sig=SuI7dH7ybPnjzrEmmtOPEpVhLiUlCEEPO-s1kg5_BXs\" の有効期限 2026-11-01T02:10:00Z が範囲外です",
    "初期データのユーザー fan_yang_1180: Memories（2026-05）: 画像 URL \"/api/images/ff080c95-7009-4712-a867-0abc5ec0a9f0?exp=1793499000\u0026sig=dxIf88useKlxdsdqT84fNYlAdAM5jt3OQLNwtZzZvEY\" の有効期限 2026-11-01T02:10:00Z が範囲外です",
    "初期データのユーザー charcoal_mini_turkey_2371: Memories（2026-04）: 画像 URL \"/api/images/213d28c7-b216-44b6-882c-a5bec6a69201?exp=1793499000\u0026sig=PUPto52V2f32hvtuFNfRtKiNB5fZS6k8PX5JcEu4Zrg\" の有効期限 2026-11-01T02:10:00Z が範囲外です",
    "初期データのユーザー sardinhas_na_panela_de_barro_4757: Memories（2026-08）: 画像 URL \"/api/images/3ad6df78-04e5-4d25-bb3a-2a5fc3ca36da?exp=1793499000\u0026sig=6TmXc6yYe4ymI4Sx4AEfJhKa6G_vfYtldULE-FK9YqI\" の有効期限 2026-11-01T02:10:00Z が範囲外です",
    "初期データのユーザー jihun_park_2: Memories（2025-11）: 画像 URL \"/api/images/777d5e88-9eb9-4084-ac7d-b4c3743ff421?exp=1793499000\u0026sig=gLtrElf772e77jwzbh3xXBBH0fxc_iWHepyEKfgXgoc\" の有効期限 2026-11-01T02:10:00Z が範囲外です"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
87.1%	1	0	0	POST	/api/initialize	12.168	12.168	12.168	12.168	90
5.8%	24	1	0	POST	/api/auth/login	0.805	0.034	0.085	0.085	221
4.7%	12	1	0	POST	/api/auth/register	0.651	0.054	0.092	0.092	210
1.5%	4	0	0	GET	/api/memories	0.209	0.052	0.174	0.174	47530
0.3%	8	0	0	PUT	^/api/follows/[^/]+$	0.041	0.005	0.008	0.008	0
0.3%	7	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	0.040	0.006	0.007	0.007	116
0.2%	7	0	0	POST	/api/friend-requests	0.021	0.003	0.004	0.004	115
0.1%	2	0	0	PUT	^/api/blocks/[^/]+$	0.011	0.005	0.006	0.006	0
0.1%	4	0	0	PUT	^/api/admin/users/[^/]+/certification$	0.009	0.002	0.003	0.003	0
0.1%	3	2	0	PUT	/api/me/profile-image	0.009	0.003	0.007	0.007	23
0.0%	1	0	0	POST	/api/admin/auth/login	0.004	0.004	0.004	0.004	106
0.0%	1	0	0	POST	/api/auth/logout	0.003	0.003	0.003	0.003	0
0.0%	4	2	0	GET	/api/me	0.001	0.000	0.001	0.001	137
0.0%	1	1	0	DELETE	^/api/admin/users/[^/]+/certification$	0.001	0.001	0.001	0.001	39
0.0%	3	2	0	GET	^/api/images/[^/]+	0.001	0.000	0.001	0.001	4400
0.0%	1	1	0	GET	^/api/posts/me/current$	0.000	0.000	0.000	0.000	27
0.0%	1	0	0	GET	^/api/feeds/[^/]+	0.000	0.000	0.000	0.000	29

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.4%	8	0	0	POST	/api/auth/login	1.320	0.165	0.207	0.207	240
26.2%	5	0	0	POST	/api/auth/register	1.004	0.201	0.217	0.217	231
11.0%	4	0	0	PUT	^/api/follows/[^/]+$	0.420	0.105	0.107	0.107	0
8.2%	3	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	0.315	0.105	0.106	0.106	116
8.1%	3	0	0	POST	/api/friend-requests	0.311	0.104	0.104	0.104	115
5.3%	1	1	0	GET	/api/me	0.202	0.202	0.202	0.202	28
4.0%	1	0	0	GET	/api/memories	0.155	0.155	0.155	0.155	29350
2.7%	1	0	0	PUT	^/api/blocks/[^/]+$	0.105	0.105	0.105	0.105	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
31.5%	7	0	0	POST	/api/auth/login	1.134	0.162	0.207	0.207	232
29.2%	5	0	0	POST	/api/auth/register	1.050	0.210	0.218	0.218	228
14.5%	4	0	0	PUT	^/api/follows/[^/]+$	0.522	0.131	0.204	0.204	0
8.7%	3	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	0.313	0.104	0.105	0.105	116
8.6%	3	0	0	POST	/api/friend-requests	0.311	0.104	0.104	0.104	115
4.5%	1	0	0	GET	/api/memories	0.162	0.162	0.162	0.162	39043
2.9%	1	0	0	PUT	^/api/blocks/[^/]+$	0.106	0.106	0.106	0.106	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 344 クエリ / 57 種類 / 合計 7.287s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	33.6%	1	2.445	2445.42	2445.42	29314.0	1199408.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
2	21.0%	1	1.532	1531.83	1531.83	19538.0	625426.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
3	14.7%	1	1.071	1070.84	1070.84	11528.0	529643.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
4	6.7%	2	0.489	244.48	337.88	213896.0	2516.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
5	4.4%	2	0.322	161.23	314.50	20233.5	2506.5	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
6	3.2%	1	0.231	231.12	231.12	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
7	2.9%	1	0.213	212.94	212.94	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
8	2.9%	1	0.210	209.79	209.79	69905.0	29793.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
9	2.1%	1	0.151	151.33	151.33	29375.0	60.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	1.4%	1	0.103	103.49	103.49	2433.0	0.0	DELETE FROM image_objects WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	1.3%	1	0.094	93.58	93.58	20169.0	315.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	1.0%	1	0.071	71.32	71.32	2594.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
13	0.9%	1	0.065	65.18	65.18	11542.0	13.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.9%	1	0.063	62.59	62.59	1531.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations
15	0.9%	1	0.062	62.08	62.08	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 2
SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM posts;
-- rank 3
SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 8236bb6b322e24f11b151e37777952b81bf4206b
Type: cpu
Time: 2026-10-31 11:02:56 JST
Duration: 60.02s, Total samples = 10.83s (18.05%)
Showing nodes accounting for 9s, 83.10% of 10.83s total
Dropped 191 nodes (cum <= 0.05s)
Showing top 40 nodes out of 164
      flat  flat%   sum%        cum   cum%
     1.24s 11.45% 11.45%      5.46s 50.42%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     0.84s  7.76% 19.21%      0.84s  7.76%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.80s  7.39% 26.59%      0.80s  7.39%  memeqbody
     0.64s  5.91% 32.50%      0.64s  5.91%  aeshashbody
     0.43s  3.97% 36.47%      0.43s  3.97%  runtime.spanClass.sizeclass (inline)
     0.40s  3.69% 40.17%      2.21s 20.41%  runtime.mapaccess1_faststr
     0.37s  3.42% 43.58%      0.37s  3.42%  runtime.memmove
     0.33s  3.05% 46.63%      0.58s  5.36%  example.com/isureal-company/isureal.postLess
     0.30s  2.77% 49.40%      0.79s  7.29%  runtime.tryDeferToSpanScan
     0.25s  2.31% 51.71%      0.25s  2.31%  sync/atomic.(*Int32).Add (inline)
     0.23s  2.12% 53.83%      0.23s  2.12%  crypto/internal/fips140/sha256.blockAVX2
     0.22s  2.03% 55.86%      0.22s  2.03%  runtime.nextFreeFast (inline)
     0.19s  1.75% 57.62%      0.19s  1.75%  internal/runtime/syscall/linux.Syscall6
     0.18s  1.66% 59.28%      0.18s  1.66%  cmpbody
     0.18s  1.66% 60.94%      0.44s  4.06%  runtime.scanObjectsSmall
     0.16s  1.48% 62.42%      0.20s  1.85%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.16s  1.48% 63.90%      0.77s  7.11%  runtime.mapassign_faststr
     0.16s  1.48% 65.37%      0.16s  1.48%  runtime.memclrNoHeapPointers
     0.15s  1.39% 66.76%      0.15s  1.39%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.14s  1.29% 68.05%      0.14s  1.29%  runtime.(*spanInlineMarkBits).init
     0.13s  1.20% 69.25%      0.22s  2.03%  database/sql.convertAssignRows
     0.13s  1.20% 70.45%      1.26s 11.63%  runtime.mallocgc
     0.12s  1.11% 71.56%      0.13s  1.20%  example.com/isureal-company/isureal.rmLess
     0.12s  1.11% 72.67%      0.19s  1.75%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
     0.12s  1.11% 73.78%      0.12s  1.11%  internal/runtime/maps.(*ctrlGroup).setEmpty (inline)
     0.11s  1.02% 74.79%      2.28s 21.05%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.92% 75.72%      0.12s  1.11%  runtime.(*mspan).writeHeapBitsSmall
     0.10s  0.92% 76.64%      0.74s  6.83%  runtime.mallocgcSmallScanNoHeader
     0.09s  0.83% 77.47%      1.12s 10.34%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.08s  0.74% 78.21%      0.56s  5.17%  runtime.scanObject
     0.07s  0.65% 78.86%      0.10s  0.92%  internal/runtime/maps.(*Iter).Next
     0.06s  0.55% 79.41%      0.07s  0.65%  example.com/isureal-company/isureal.commentLess
     0.06s  0.55% 79.96%      0.06s  0.55%  internal/sync.(*Mutex).Unlock (inline)
     0.06s  0.55% 80.52%      0.07s  0.65%  time.Time.Before
     0.05s  0.46% 80.98%      2.42s 22.35%  database/sql.(*Rows).nextLocked
     0.05s  0.46% 81.44%      0.44s  4.06%  runtime.convTstring
     0.05s  0.46% 81.90%      0.15s  1.39%  runtime.mapIterNext
     0.05s  0.46% 82.36%      0.18s  1.66%  runtime.wbBufFlush1
     0.04s  0.37% 82.73%      0.11s  1.02%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func5
     0.04s  0.37% 83.10%      3.39s 31.30%  example.com/isureal-company/isureal.loadBig
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 42.0% / 77% | 0% | 0.6% | 2 | 1798 MB |  |
| isu2 (2 cpu) | 0.3% / 1% | 0% | 0.0% | 0 | 329 MB |  |
| isu3 (2 cpu) | 0.2% / 1% | 0% | 0.0% | 1 | 282 MB |  |
| isu4 (2 cpu) | 25.5% / 64% | 0% | 10.6% | 1 | 132 MB |  |

どのホストも飽和していない。**ボトルネックはアプリのロジックか DB のクエリ側**にある可能性が高い。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 82% / kcompactd0 1% |
| isu2 | 目立つプロセス無し |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 24% |

## Diff from previous run

スコア: 10532 → — (fail)

`5aad88a..f264c69`

```
 conf/isu1/etc/systemd/system/isureal-go.service |  2 +-
 webapp/go/config.go                             |  2 --
 webapp/go/export_api.go                         |  4 +---
 webapp/go/image_signature.go                    |  4 ++--
 webapp/go/store.go                              | 11 -----------
 5 files changed, 4 insertions(+), 19 deletions(-)
```
