# run 20261031-122244-5443d83

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | slow-before-1700 |
| sha | 5443d831ac90 |
| 所要時間 | 76s |
| 前回比 | -65904 (-100.0%) vs 20261031-122039-fdbe826 |
| ベスト比 | -65904 (-100.0%) vs 20261031-122039-fdbe826 = 65904 |
| note | 17:00 直前用の低い点の版（試験 1000ms） |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/agent-af72e4d3a3f655e90/runs/20261031-122244-5443d83/bench.log
```

## Bench

```
{"id":"7848","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"初期化に失敗しました"}

--- stdout ---
{
  "pass": false,
  "score": 0,
  "messages": [
    "初期化に失敗しました: POST /api/initialize: ステータスコード 504 が返りました（期待値 200、unknown）（上限 1m0s）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
100.0%	1	1	0	POST	/api/initialize	60.001	60.001	60.001	60.001	0

== isu2 ==
（リクエストなし）

== isu3 ==
（リクエストなし）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 628 クエリ / 19 種類 / 合計 145.783s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	59.4%	1	86.604	86603.67	86603.67	8675981.0	1.0	select encode(data,'base64') from image_objects order by created_at desc limit 1
2	33.6%	1	49.028	49028.49	49028.49	237879.0	2799.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
3	2.6%	1	3.774	3773.97	3773.97	29864.0	1199948.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
4	2.3%	1	3.332	3331.85	3331.85	21119.0	626828.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	1.2%	1	1.733	1732.83	1732.83	11571.0	530068.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	0.8%	1	1.134	1133.81	1133.81	33393.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
7	0.0%	1	0.060	60.24	60.24	225.0	5590.0	SELECT id, username, secret_key_hash, display_name, bio, profile_image_id, region, account_type, status, created_at FROM users
8	0.0%	1	0.034	34.18	34.18	661.0	32109.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.0%	1	0.032	32.46	32.46	647.0	26318.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
10	0.0%	610	0.014	0.02	0.06	1.0	22.7	SELECT ip, banned_until FROM ip_bans WHERE banned_until > clock_timestamp()
11	0.0%	1	0.013	12.77	12.77	229.0	10446.0	SELECT id, follower_id, followee_id, created_at, updated_at FROM follows
12	0.0%	1	0.008	8.06	8.06	114.0	3.0	SELECT id, notification_ids, notified_user_ids FROM posting_windows WHERE ended_at IS NULL AND notification_ids IS NOT NULL
13	0.0%	1	0.006	5.92	5.92	82.0	1999.0	SELECT id, user_id, token_hash, created_at FROM login_sessions WHERE active = $1
14	0.0%	1	0.004	4.24	4.24	42.0	1113.0	SELECT id, region, local_date, starts_at, on_time_deadline, ended_at, created_at FROM posting_windows
15	0.0%	1	0.003	3.14	3.14	8.0	0.0	ALTER TABLE posting_windows ADD COLUMN IF NOT EXISTS notification_ids text[], ADD COLUMN IF NOT EXISTS notified_user_ids text[]

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
select encode(data,'base64') from image_objects order by created_at desc limit 1;
-- rank 2
SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id;
-- rank 3
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 466e94c217da39967ba3eaef15717ade3fb0fcf0
Type: cpu
Time: 2026-10-31 12:22:59 JST
Duration: 60.02s, Total samples = 6.59s (10.98%)
Showing nodes accounting for 5.57s, 84.52% of 6.59s total
Dropped 124 nodes (cum <= 0.03s)
Showing top 40 nodes out of 127
      flat  flat%   sum%        cum   cum%
     0.75s 11.38% 11.38%      0.75s 11.38%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.30s  4.55% 15.93%      0.42s  6.37%  database/sql.convertAssignRows
     0.29s  4.40% 20.33%      1.13s 17.15%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     0.27s  4.10% 24.43%      0.27s  4.10%  memeqbody
     0.27s  4.10% 28.53%      0.27s  4.10%  runtime.memmove
     0.25s  3.79% 32.32%      0.25s  3.79%  aeshashbody
     0.25s  3.79% 36.12%      0.78s 11.84%  runtime.scanObject
     0.22s  3.34% 39.45%      0.22s  3.34%  sync/atomic.(*Int32).Add (inline)
     0.20s  3.03% 42.49%      0.28s  4.25%  runtime.tryDeferToSpanScan
     0.18s  2.73% 45.22%      0.28s  4.25%  example.com/isureal-company/isureal.postLess
     0.17s  2.58% 47.80%      0.17s  2.58%  internal/runtime/syscall/linux.Syscall6
     0.16s  2.43% 50.23%      0.23s  3.49%  runtime.findObject
     0.16s  2.43% 52.66%      0.55s  8.35%  runtime.mapaccess1_faststr
     0.15s  2.28% 54.93%      0.18s  2.73%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
     0.15s  2.28% 57.21%      0.15s  2.28%  runtime.nextFreeFast (inline)
     0.14s  2.12% 59.33%      2.06s 31.26%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.13s  1.97% 61.31%      0.93s 14.11%  runtime.mallocgc
     0.12s  1.82% 63.13%      0.52s  7.89%  runtime.convTstring
     0.12s  1.82% 64.95%      0.45s  6.83%  runtime.mapaccess2_faststr
     0.10s  1.52% 66.46%      4.48s 67.98%  example.com/isureal-company/isureal.loadBig
     0.10s  1.52% 67.98%      0.48s  7.28%  runtime.mallocgcSmallScanNoHeader
     0.09s  1.37% 69.35%      0.12s  1.82%  github.com/jackc/pgx/v5/pgtype.(*scanPlanBinaryTimestamptzToTimestamptzScanner).Scan
     0.08s  1.21% 70.56%      0.08s  1.21%  cmpbody
     0.08s  1.21% 71.78%      0.08s  1.21%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.08s  1.21% 72.99%      1.08s 16.39%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.08s  1.21% 74.20%      0.08s  1.21%  internal/runtime/maps.(*ctrlGroup).setEmpty (inline)
     0.08s  1.21% 75.42%      0.08s  1.21%  runtime.memclrNoHeapPointers
     0.07s  1.06% 76.48%      0.48s  7.28%  github.com/jackc/pgx/v5/pgtype.scanPlanString.Scan
     0.06s  0.91% 77.39%      0.06s  0.91%  runtime.(*mspan).base (inline)
     0.06s  0.91% 78.30%      0.71s 10.77%  runtime.mapassign_faststr
     0.06s  0.91% 79.21%      0.06s  0.91%  runtime.spanClass.sizeclass (inline)
     0.06s  0.91% 80.12%      0.06s  0.91%  runtime.typePointers.nextFast (inline)
     0.05s  0.76% 80.88%      0.17s  2.58%  runtime.mallocgcSmallNoscan
     0.04s  0.61% 81.49%      0.46s  6.98%  database/sql.(*Rows).scanLocked
     0.04s  0.61% 82.09%      0.06s  0.91%  runtime.(*mspan).writeHeapBitsSmall
     0.04s  0.61% 82.70%      0.06s  0.91%  runtime.mallocgcTiny
     0.03s  0.46% 83.16%      2.36s 35.81%  database/sql.(*Rows).Next
     0.03s  0.46% 83.61%      0.17s  2.58%  example.com/isureal-company/isureal.(*interner).id (inline)
     0.03s  0.46% 84.07%      0.30s  4.55%  github.com/jackc/pgx/v5/pgproto3.(*chunkReader).Next
     0.03s  0.46% 84.52%      0.07s  1.06%  runtime.greyobject
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 5.8% / 60% | 0% | 0.1% | 2 | 2003 MB |  |
| isu2 (2 cpu) | 0.1% / 1% | 0% | 0.0% | 0 | 2792 MB |  |
| isu3 (2 cpu) | 0.1% / 1% | 0% | 0.0% | 2 | 2856 MB |  |
| isu4 (2 cpu) | 95.8% / 100% | 92% | 57.8% | 5 | 492 MB | CPU が飽和 |

**isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 13% |
| isu2 | 目立つプロセス無し |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 73% |

## Diff from previous run

スコア: 65904 → — (fail)

`fdbe826..5443d83`

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
