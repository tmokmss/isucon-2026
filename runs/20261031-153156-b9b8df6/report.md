# run 20261031-153156-b9b8df6

## Summary

| | |
|---|---|
| score | 180192 |
| status | pass |
| ref | eu-us-wait |
| sha | b9b8df630aa8 |
| 所要時間 | 303s |
| 前回比 | -3139 (-1.7%) vs 20261031-152137-8c13252 |
| ベスト比 | -11624 (-6.1%) vs 20261031-151339-5a230a4 = 191816 |
| note | eu/us: 画像取得の同時要求を1本に・接続の使い回し・BBR |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	147	429	1169	1172
europe	7	19	54	158	424	423
america	6	19	52	149	287	287
```

```
{"id":"23300","status":"done","portalStatus":"finished","passed":true,"score":180192,"raw":180192,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:35:21.482Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:35:24.476Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:35:27.473Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:35:30.485Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T06:35:33.078Z level=INFO msg="user moved" user=wabente.5288 from=asia to=america
time=2026-10-31T06:35:33.474Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T06:35:36.473Z level=INFO msg="america 2 日目: 19 人"
time=2026-10-31T06:35:39.483Z level=INFO msg="asia 3 日目: 147 人"
time=2026-10-31T06:35:42.474Z level=INFO msg="europe 3 日目: 54 人"
time=2026-10-31T06:35:42.800Z level=INFO msg="user moved" user=romaryo-6143 from=asia to=europe
time=2026-10-31T06:35:45.477Z level=INFO msg="america 3 日目: 52 人"
time=2026-10-31T06:35:48.393Z level=INFO msg="user moved" user=mechisol_6126 from=europe to=america
time=2026-10-31T06:35:48.486Z level=INFO msg="asia 4 日目: 429 人"
time=2026-10-31T06:35:49.941Z level=INFO msg="user moved" user=wabente.5288 from=america to=asia
time=2026-10-31T06:35:51.753Z level=INFO msg="europe 4 日目: 158 人"
time=2026-10-31T06:35:53.336Z level=INFO msg="user moved" user=wabente.5288 from=asia to=america
time=2026-10-31T06:35:54.412Z level=INFO msg="user moved" user=rihe.6906 from=america to=asia
time=2026-10-31T06:35:54.610Z level=INFO msg="user moved" user=romaryo-6143 from=europe to=america
time=2026-10-31T06:35:54.712Z level=INFO msg="america 4 日目: 149 人"
time=2026-10-31T06:35:55.663Z level=INFO msg="user moved" user=keorivy7875 from=asia to=europe
time=2026-10-31T06:35:55.692Z level=INFO msg="user moved" user=vinbenno.7058 from=asia to=europe
time=2026-10-31T06:35:56.385Z level=INFO msg="user moved" user=rihe.6906 from=asia to=europe
time=2026-10-31T06:35:56.825Z level=INFO msg="user moved" user=nolia_9032 from=europe to=america
time=2026-10-31T06:35:57.523Z level=INFO msg="asia 5 日目: 1169 人"
time=2026-10-31T06:35:57.842Z level=INFO msg="user moved" user=anteor-8812 from=europe to=asia
time=2026-10-31T06:35:59.050Z level=INFO msg="user moved" user=maheto-9326 from=europe to=asia
time=2026-10-31T06:36:01.983Z level=INFO msg="europe 5 日目: 424 人"
time=2026-10-31T06:36:02.910Z level=INFO msg="user moved" user=tsuruan9251 from=europe to=asia
time=2026-10-31T06:36:05.349Z level=INFO msg="america 5 日目: 287 人"
time=2026-10-31T06:36:13.661Z level=INFO msg="asia 最終日（グランプリ）: 1172 人"
time=2026-10-31T06:36:14.597Z level=INFO msg="europe 最終日（グランプリ）: 423 人"
time=2026-10-31T06:36:15.189Z level=INFO msg="america 最終日（グランプリ）: 287 人"
time=2026-10-31T06:36:21.574Z level=INFO msg="グランプリ: アジア 510 人, ヨーロッパ 90 人, アメリカ 63 人"
{
  "pass": true,
  "score": 180192,
  "messages": [
    "kunofu15671: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ankusa-16232: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kuta_16422: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "honu.14432: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "leoketor7444: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hator6946: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hosu-15422: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tsudan16787: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sunoaivy_16128: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "leoshi-15711: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ulbenko15776: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nitanoa_16101: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "neme.11783: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tenohi15859: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kara-11992: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kimia_14335: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tormaben-5248: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "haeva12638: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "bensasu-11873: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marorna.16454: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
（リクエストなし）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
96.7%	122	0	0	GET	^/api/images/[^/]+	15.181	0.124	0.165	0.188	58110
1.5%	4	0	0	GET	/api/memories	0.233	0.058	0.061	0.061	14503
1.4%	4	0	0	GET	^/api/relationships/[^/]+	0.222	0.056	0.056	0.056	10572
0.4%	1	0	0	GET	/api/me	0.056	0.056	0.056	0.056	201

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
89.3%	93	0	0	GET	^/api/images/[^/]+	10.946	0.118	0.200	0.200	107340
5.6%	12	0	0	GET	/api/memories	0.681	0.057	0.058	0.058	4109
3.7%	8	0	0	GET	^/api/relationships/[^/]+	0.454	0.057	0.058	0.058	8473
1.4%	3	0	0	GET	/api/me	0.170	0.057	0.057	0.057	197
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 18187 クエリ / 48 種類 / 合計 147.131s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	51.0%	398	75.089	188.67	1074.03	1190.4	121.9	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	15.6%	3	22.991	7663.70	10343.47	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.9%	14329	17.500	1.22	9.54	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.9%	3	16.067	5355.58	7222.75	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.7%	3	8.432	2810.73	3870.16	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.8%	6	4.156	692.61	3190.28	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.5%	6	0.690	115.07	165.68	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.5%	6	0.689	114.88	162.58	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.251	251.19	251.19	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	0.2%	1	0.248	247.94	247.94	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.224	223.89	223.89	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.161	161.39	161.39	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.106	106.19	106.19	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.080	80.35	80.35	2486.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.071	70.91	70.91	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT data FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 66826fab865ca9633923e8703030c4c36feb158d
Type: cpu
Time: 2026-10-31 15:34:59 JST
Duration: 60.24s, Total samples = 43.34s (71.94%)
Showing nodes accounting for 30.12s, 69.50% of 43.34s total
Dropped 1165 nodes (cum <= 0.22s)
Showing top 40 nodes out of 293
      flat  flat%   sum%        cum   cum%
     4.95s 11.42% 11.42%      4.95s 11.42%  internal/runtime/syscall/linux.Syscall6
     4.26s  9.83% 21.25%      4.26s  9.83%  crypto/internal/fips140/sha256.blockAVX2
     2.02s  4.66% 25.91%      2.02s  4.66%  runtime.memmove
     1.89s  4.36% 30.27%      6.94s 16.01%  runtime.scanObject
     1.86s  4.29% 34.56%      5.65s 13.04%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.75s  4.04% 38.60%      2.26s  5.21%  runtime.findObject
     1.61s  3.71% 42.32%      1.61s  3.71%  aeshashbody
     1.47s  3.39% 45.71%      2.65s  6.11%  runtime.tryDeferToSpanScan
     1.40s  3.23% 48.94%      1.40s  3.23%  runtime.memclrNoHeapPointers
     1.37s  3.16% 52.10%      1.37s  3.16%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.83s  1.92% 54.01%      0.83s  1.92%  runtime.spanClass.sizeclass (inline)
     0.59s  1.36% 55.38%      1.98s  4.57%  runtime.mapaccess1_faststr
     0.45s  1.04% 56.41%      0.45s  1.04%  runtime.nextFreeFast (inline)
     0.40s  0.92% 57.34%      0.61s  1.41%  example.com/isureal-company/isureal.postLess
     0.37s  0.85% 58.19%      0.37s  0.85%  memeqbody
     0.34s  0.78% 58.98%      0.34s  0.78%  runtime.futex
     0.30s  0.69% 59.67%      0.77s  1.78%  database/sql.convertAssignRows
     0.29s  0.67% 60.34%      4.74s 10.94%  runtime.mallocgc
     0.28s  0.65% 60.98%      0.64s  1.48%  runtime.greyobject
     0.25s  0.58% 61.56%      5.44s 12.55%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.25s  0.58% 62.14%      4.85s 11.19%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.24s  0.55% 62.69%      0.26s   0.6%  runtime.(*mspan).markBitsForIndex (inline)
     0.23s  0.53% 63.22%      0.28s  0.65%  encoding/json.appendString[go.shape.string]
     0.23s  0.53% 63.75%      0.30s  0.69%  runtime.heapArenaOf (inline)
     0.22s  0.51% 64.26%      0.51s  1.18%  compress/flate.(*compressor).deflate
     0.20s  0.46% 64.72%      4.79s 11.05%  crypto/internal/fips140/sha256.(*Digest).Write
     0.20s  0.46% 65.18%      2.27s  5.24%  runtime.mapassign_faststr
     0.20s  0.46% 65.64%      0.22s  0.51%  runtime.spanOf (inline)
     0.19s  0.44% 66.08%      0.36s  0.83%  internal/runtime/maps.(*Iter).Next
     0.19s  0.44% 66.52%      1.25s  2.88%  runtime.mallocgcSmallScanNoHeader
     0.17s  0.39% 66.91%      2.98s  6.88%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.17s  0.39% 67.31%      0.65s  1.50%  runtime.mapaccess2_faststr
     0.14s  0.32% 67.63%      0.71s  1.64%  runtime.wbBufFlush1
     0.13s   0.3% 67.93%      0.24s  0.55%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.13s   0.3% 68.23%      0.96s  2.22%  runtime.mallocgcSmallNoscan
     0.11s  0.25% 68.48%      5.10s 11.77%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.11s  0.25% 68.74%      0.93s  2.15%  encoding/json.structEncoder.encode
     0.11s  0.25% 68.99%      3.93s  9.07%  example.com/isureal-company/isureal.loadBig
     0.11s  0.25% 69.24%      0.28s  0.65%  internal/stringslite.Index
     0.11s  0.25% 69.50%      0.45s  1.04%  runtime.scanObjectsSmall
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 42.5% / 100% | 26% | 0.5% | 5 | 115 MB | run queue が積もっている |
| isu2 (2 cpu) | 28.0% / 100% | 3% | 0.6% | 4 | 245 MB |  |
| isu3 (2 cpu) | 26.1% / 100% | 7% | 0.7% | 5 | 239 MB | run queue が積もっている |
| isu4 (2 cpu) | 40.7% / 100% | 13% | 15.6% | 4 | 106 MB |  |

**isu1（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 68% / nginx 12% |
| isu2 | isureal 52% / nginx 2% |
| isu3 | isureal 49% / nginx 1% |
| isu4 | isureal 33% / postgres 15% |

## Diff from previous run

スコア: 183331 → 180192 (-3139)

`8c13252..b9b8df6`

```
 conf/_all/etc/modules-load.d/bbr.conf |  1 +
 conf/_all/etc/sysctl.d/99-isucon.conf |  2 ++
 webapp/go/image_file.go               |  5 ++++-
 webapp/go/image_held.go               | 41 +++++++++++++++++++++++++++++++----
 webapp/go/image_local.go              |  1 -
 webapp/go/origin_store.go             |  2 +-
 webapp/go/persist.go                  |  2 +-
 webapp/go/realmoji_api.go             | 41 ++++++++++++++---------------------
 webapp/go/replica_server.go           |  8 +------
 webapp/go/store.go                    |  3 +--
 10 files changed, 64 insertions(+), 42 deletions(-)
```
