# run 20261031-113320-54afe51

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | limiter-priority |
| sha | 54afe515e703 |
| 所要時間 | 111s |
| 前回比 | -18045 (-100.0%) vs 20261031-112932-d2d4b7d |
| ベスト比 | -24801 (-100.0%) vs 20261031-103811-cef5fe3 = 24801 |
| note | limiter post4-8s heavy1 on main-4764fb5 |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/agent-a4cbd2c10b54f68d3/runs/20261031-113320-54afe51/bench.log
```

## Bench

```
{"id":"4443","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"初期化に失敗しました"}

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
100.0%	1	1	0	POST	/api/initialize	60.002	60.002	60.002	60.002	0

== isu2 ==
（リクエストなし）

== isu3 ==
（リクエストなし）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 626 クエリ / 16 種類 / 合計 84.704s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	90.4%	1	76.556	76556.03	76556.03	705183.0	150273.0	COPY public.image_objects_202511 (id, owner_user_id, data, media_type, created_at) TO stdout
2	4.7%	1	3.987	3986.75	3986.75	29389.0	1199365.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	3.2%	1	2.699	2698.71	2698.71	19882.0	625231.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	1.3%	1	1.126	1125.85	1125.85	11528.0	529637.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
5	0.3%	1	0.267	266.88	266.88	214151.0	2519.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
6	0.0%	1	0.022	21.75	21.75	617.0	25057.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
7	0.0%	611	0.016	0.03	0.14	2.0	25.0	SELECT ip, banned_until FROM ip_bans WHERE banned_until > clock_timestamp()
8	0.0%	1	0.015	15.14	15.14	613.0	29242.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.0%	1	0.008	7.76	7.76	194.0	10056.0	SELECT id, follower_id, followee_id, created_at, updated_at FROM follows
10	0.0%	1	0.005	4.77	4.77	225.0	5081.0	SELECT id, username, secret_key_hash, display_name, bio, profile_image_id, region, account_type, status, created_at FROM users
11	0.0%	1	0.002	2.48	2.48	26.0	1108.0	SELECT id, region, local_date, starts_at, on_time_deadline, ended_at, created_at FROM posting_windows
12	0.0%	1	0.001	1.03	1.03	25.0	303.0	SELECT id, user_id, token_hash, created_at FROM login_sessions WHERE active = $1
13	0.0%	1	0.001	0.60	0.60	1.0	1.0	SELECT id, administrator_id, token_hash, created_at FROM administrator_sessions WHERE active = $1
14	0.0%	1	0.000	0.13	0.13	5.0	255.0	SELECT id, blocker_id, blocked_id, created_at FROM block_relations
15	0.0%	1	0.000	0.01	0.01	1.0	1.0	SELECT id, username, password_hash FROM administrators

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
COPY public.image_objects_202511 (id, owner_user_id, data, media_type, created_at) TO stdout;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM posts;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: efa3dc4f0319658b11bdd13a69ddb475fe8694f5
Type: cpu
Time: 2026-10-31 11:34:09 JST
Duration: 60s, Total samples = 6.89s (11.48%)
Showing nodes accounting for 5.71s, 82.87% of 6.89s total
Dropped 101 nodes (cum <= 0.03s)
Showing top 40 nodes out of 115
      flat  flat%   sum%        cum   cum%
     0.61s  8.85%  8.85%      0.61s  8.85%  runtime.spanClass.sizeclass (inline)
     0.43s  6.24% 15.09%      0.43s  6.24%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.33s  4.79% 19.88%      1.02s 14.80%  runtime.tryDeferToSpanScan
     0.26s  3.77% 23.66%      0.27s  3.92%  sync/atomic.(*Int32).Add (inline)
     0.25s  3.63% 27.29%      0.25s  3.63%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.25s  3.63% 30.91%      0.44s  6.39%  database/sql.convertAssignRows
     0.25s  3.63% 34.54%      0.25s  3.63%  runtime.memmove
     0.22s  3.19% 37.74%      0.22s  3.19%  internal/runtime/syscall/linux.Syscall6
     0.21s  3.05% 40.78%      3.47s 50.36%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.21s  3.05% 43.83%         1s 14.51%  runtime.mallocgcSmallScanNoHeader
     0.20s  2.90% 46.73%      0.20s  2.90%  runtime.(*spanInlineMarkBits).init
     0.19s  2.76% 49.49%      0.19s  2.76%  runtime.memclrNoHeapPointers
     0.18s  2.61% 52.10%      0.21s  3.05%  runtime.(*mspan).writeHeapBitsSmall
     0.18s  2.61% 54.72%      1.77s 25.69%  runtime.mallocgc
     0.15s  2.18% 56.89%      0.44s  6.39%  runtime.scanObjectsSmall
     0.13s  1.89% 58.78%      0.66s  9.58%  runtime.mapassign_faststr
     0.12s  1.74% 60.52%      0.12s  1.74%  runtime.nextFreeFast (inline)
     0.11s  1.60% 62.12%      0.19s  2.76%  github.com/jackc/pgx/v5/pgtype.(*scanPlanBinaryTimestamptzToTimestamptzScanner).Scan
     0.10s  1.45% 63.57%      5.42s 78.66%  example.com/isureal-company/isureal.loadBig
     0.10s  1.45% 65.02%      0.10s  1.45%  runtime.acquirem (inline)
     0.10s  1.45% 66.47%      0.68s  9.87%  runtime.scanObject
     0.09s  1.31% 67.78%      0.33s  4.79%  gcWriteBarrier
     0.09s  1.31% 69.09%      0.15s  2.18%  github.com/jackc/pgx/v5/pgproto3.(*DataRow).Decode
     0.09s  1.31% 70.39%      0.49s  7.11%  runtime.mallocgcSmallNoscan
     0.08s  1.16% 71.55%      1.88s 27.29%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.08s  1.16% 72.71%      0.84s 12.19%  runtime.convTstring
     0.08s  1.16% 73.88%      0.52s  7.55%  runtime.scanSpan
     0.06s  0.87% 74.75%      0.06s  0.87%  internal/runtime/maps.(*ctrlGroup).setEmpty (inline)
     0.06s  0.87% 75.62%      0.07s  1.02%  runtime.mallocgcTiny
     0.06s  0.87% 76.49%      0.91s 13.21%  runtime.slicebytetostring
     0.05s  0.73% 77.21%      0.46s  6.68%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func12
     0.05s  0.73% 77.94%      0.06s  0.87%  internal/sync.(*Mutex).Lock (inline)
     0.05s  0.73% 78.66%      0.10s  1.45%  runtime.gcmarknewobject
     0.05s  0.73% 79.39%      0.05s  0.73%  runtime.getMCache (inline)
     0.04s  0.58% 79.97%      0.48s  6.97%  database/sql.(*Rows).scanLocked
     0.04s  0.58% 80.55%      0.80s 11.61%  github.com/jackc/pgx/v5/pgconn.(*PgConn).receiveMessage
     0.04s  0.58% 81.13%      0.56s  8.13%  github.com/jackc/pgx/v5/pgproto3.(*chunkReader).Next
     0.04s  0.58% 81.71%      0.04s  0.58%  internal/runtime/maps.(*ctrlGroup).get (inline)
     0.04s  0.58% 82.29%      0.04s  0.58%  internal/sync.(*Mutex).Unlock (inline)
     0.04s  0.58% 82.87%      0.04s  0.58%  runtime.bool2int (inline)
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 61.5% / 100% | 12% | 0.1% | 4 | 2122 MB |  |
| isu2 (2 cpu) | 0.5% / 15% | 0% | 0.3% | 2 | 1284 MB |  |
| isu3 (2 cpu) | 0.6% / 1% | 0% | 0.0% | 2 | 316 MB |  |
| isu4 (2 cpu) | 11.6% / 53% | 0% | 3.8% | 1 | 114 MB |  |

どのホストも飽和していない。**ボトルネックはアプリのロジックか DB のクエリ側**にある可能性が高い。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | pg_dump 99% / isureal 11% / sshd 9% |
| isu2 | 目立つプロセス無し |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 14% |

## Diff from previous run

スコア: 18045 → — (fail)

`d2d4b7d..54afe51`

```
 webapp/go/limiter.go | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)
```
