# run 20261031-102655-8cd1279

## Summary

| | |
|---|---|
| score | 6389 |
| status | pass |
| ref | HEAD |
| sha | 8cd127924913 |
| 所要時間 | 300s |
| 前回比 | +3715 (138.9%) vs 20261031-100914-e240203 |
| ベスト比 | +3715 (138.9%) vs 20261031-100914-e240203 = 2674 **ベスト更新** |
| note | pprof を入れた（初期状態のプロファイル） |

## Bench

```
{"id":"844","status":"done","portalStatus":"finished","passed":true,"score":6389,"raw":6389,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:30:06.255Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:30:07.985Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:30:10.843Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:30:16.405Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T01:30:16.411Z level=INFO msg="user moved" user=shito.5039 from=asia to=america
time=2026-10-31T01:30:18.832Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T01:30:21.170Z level=INFO msg="america 2 日目: 7 人"
time=2026-10-31T01:30:26.341Z level=INFO msg="asia 3 日目: 17 人"
time=2026-10-31T01:30:26.393Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T01:30:30.489Z level=INFO msg="america 3 日目: 7 人"
time=2026-10-31T01:30:32.915Z level=INFO msg="asia 4 日目: 17 人"
time=2026-10-31T01:30:35.107Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T01:30:41.620Z level=INFO msg="america 4 日目: 7 人"
time=2026-10-31T01:30:43.807Z level=INFO msg="asia 5 日目: 18 人"
time=2026-10-31T01:30:45.101Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T01:30:47.006Z level=INFO msg="america 5 日目: 7 人"
time=2026-10-31T01:30:57.958Z level=INFO msg="america 最終日（グランプリ）: 7 人"
time=2026-10-31T01:30:57.987Z level=INFO msg="europe 最終日（グランプリ）: 6 人"
time=2026-10-31T01:30:59.279Z level=INFO msg="asia 最終日（グランプリ）: 17 人"
time=2026-10-31T01:31:04.706Z level=INFO msg="グランプリ: アジア 17 人, ヨーロッパ 2 人, アメリカ 3 人"
{
  "pass": true,
  "score": 6389,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
55.3%	105	0	0	GET	/api/notifications	1034.460	9.852	15.623	17.096	176
11.8%	1496	198	0	GET	^/api/images/[^/]+	221.323	0.148	0.398	0.556	102881
5.6%	819	0	0	GET	^/api/feeds/[^/]+	104.257	0.127	0.419	0.961	2051
4.6%	90	0	0	POST	/api/posts	85.542	0.950	2.228	2.228	588
4.3%	351	0	0	GET	/api/memories	79.945	0.228	1.076	1.289	12587
3.0%	961	558	0	POST	/api/auth/login	55.559	0.058	0.241	0.373	122
2.8%	102	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	52.556	0.515	4.725	5.480	116
2.2%	34	0	0	GET	/api/me/export	40.445	1.190	3.071	3.071	1978360
2.0%	112	5	0	POST	/api/friend-requests	37.753	0.337	2.990	3.090	111
1.8%	18	0	0	POST	/api/admin/posting-window	32.765	1.820	3.912	3.912	177
1.0%	35	1	0	POST	/api/auth/register	18.219	0.521	3.381	3.381	237
0.9%	173	1	0	GET	^/api/posts/[^/]+$	17.180	0.099	0.353	0.412	588
0.8%	102	99	2	PUT	/api/me/profile-image	14.798	0.145	0.477	0.630	26
0.6%	93	0	0	GET	/api/posting-window/current	11.870	0.128	0.485	0.485	237
0.4%	72	0	0	GET	^/api/posts/[^/]+/realmojis$	8.277	0.115	0.377	0.377	131
0.4%	47	0	0	GET	/api/users	7.439	0.158	0.291	0.291	235
0.4%	91	1	0	GET	^/api/posts/me/current$	7.239	0.080	0.313	0.313	582
0.3%	49	0	0	GET	^/api/posts/[^/]+/comments$	6.022	0.123	0.339	0.339	51
0.3%	88	0	0	GET	^/api/users/[^/]+$	5.455	0.062	0.264	0.264	162
0.3%	114	0	0	GET	^/api/relationships/[^/]+	5.136	0.045	0.159	0.164	2218
0.2%	18	0	0	POST	^/api/posts/[^/]+/realmojis$	3.561	0.198	0.493	0.493	410
0.2%	70	0	0	GET	^/api/memories/[^/]+	3.437	0.049	0.244	0.244	366
0.2%	71	0	0	POST	/api/auth/logout	3.436	0.048	0.193	0.193	0
0.2%	9	0	0	POST	^/api/posts/[^/]+/comments$	2.925	0.325	1.085	1.085	275
0.2%	1	0	0	POST	/api/initialize	2.855	2.855	2.855	2.855	69
0.1%	3	0	0	DELETE	^/api/friendships/[^/]+$	2.041	0.680	1.928	1.928	0
0.1%	13	0	0	PUT	^/api/follows/[^/]+$	1.900	0.146	1.155	1.155	0
0.1%	4	0	0	PATCH	/api/me	1.678	0.419	1.157	1.157	211
0.1%	9	0	0	PATCH	^/api/posts/[^/]+$	1.478	0.164	0.395	0.395	591
0.0%	4	0	0	DELETE	^/api/posts/[^/]+$	0.351	0.088	0.112	0.112	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.6%	42	0	0	GET	/api/notifications	404.505	9.631	13.316	13.316	177
11.5%	31	0	0	POST	/api/posts	106.244	3.427	6.972	6.972	615
8.6%	98	3	0	GET	^/api/images/[^/]+	80.188	0.818	1.462	1.462	192620
7.7%	42	1	0	POST	/api/friend-requests	71.206	1.695	3.983	3.983	113
5.8%	54	0	0	GET	/api/memories	54.165	1.003	2.422	2.422	13073
5.2%	37	0	0	GET	^/api/feeds/[^/]+	48.062	1.299	2.306	2.306	719
3.9%	26	0	0	POST	/api/auth/login	36.329	1.397	2.002	2.002	257
2.3%	5	0	0	POST	/api/auth/register	21.062	4.212	6.059	6.059	226
2.2%	22	0	0	GET	/api/users	20.404	0.927	1.953	1.953	257
2.1%	5	0	0	PUT	^/api/follows/[^/]+$	19.360	3.872	5.546	5.546	0
1.9%	34	0	0	GET	/api/posting-window/current	17.961	0.528	1.474	1.474	239
1.6%	27	10	0	GET	^/api/posts/me/current$	14.842	0.550	1.027	1.027	412
1.0%	3	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	9.709	3.236	4.451	4.451	116
0.6%	13	0	0	GET	^/api/relationships/[^/]+	5.481	0.422	0.608	0.608	1842
0.5%	3	0	0	GET	/api/me/export	4.258	1.419	1.955	1.955	1369568
0.4%	3	0	0	DELETE	^/api/posts/[^/]+$	3.490	1.163	1.337	1.337	0
0.3%	4	0	0	GET	^/api/memories/[^/]+	2.436	0.609	0.751	0.751	189
0.2%	5	2	0	GET	/api/me	2.231	0.446	0.509	0.509	125
0.2%	1	0	0	PUT	^/api/blocks/[^/]+$	1.718	1.718	1.718	1.718	0
0.1%	1	0	0	GET	^/api/posts/[^/]+$	1.125	1.125	1.125	1.125	576
0.1%	1	0	0	POST	/api/me/account-deletion	1.022	1.022	1.022	1.022	0
0.1%	1	1	0	PATCH	^/api/posts/[^/]+$	0.810	0.810	0.810	0.810	30
0.1%	1	0	0	POST	/api/auth/logout	0.727	0.727	0.727	0.727	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.5%	42	0	0	GET	/api/notifications	400.397	9.533	12.471	12.471	178
11.5%	30	0	0	POST	/api/posts	108.719	3.624	7.074	7.074	595
9.5%	110	1	0	GET	^/api/images/[^/]+	89.933	0.818	1.483	1.580	187905
6.2%	33	1	0	POST	/api/friend-requests	58.076	1.760	4.530	4.530	112
6.1%	59	0	0	GET	/api/memories	57.068	0.967	1.803	1.803	10045
5.6%	38	0	0	GET	^/api/feeds/[^/]+	52.538	1.383	2.314	2.314	865
3.5%	24	0	0	POST	/api/auth/login	32.639	1.360	1.938	1.938	252
2.8%	11	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	26.034	2.367	5.265	5.265	116
2.3%	5	0	0	POST	/api/auth/register	21.790	4.358	7.485	7.485	228
2.3%	35	0	0	GET	/api/posting-window/current	21.551	0.616	1.299	1.299	240
1.9%	4	0	0	PUT	^/api/follows/[^/]+$	18.353	4.588	6.158	6.158	0
1.8%	28	10	0	GET	^/api/posts/me/current$	16.914	0.604	1.269	1.269	396
1.4%	13	0	0	GET	/api/users	12.778	0.983	1.971	1.971	218
1.0%	22	0	0	GET	^/api/relationships/[^/]+	9.118	0.414	0.619	0.619	1178
0.5%	5	0	0	POST	/api/auth/logout	4.716	0.943	1.035	1.035	0
0.3%	4	0	0	GET	^/api/memories/[^/]+	2.946	0.736	0.968	0.968	937
0.2%	2	2	0	PATCH	^/api/posts/[^/]+$	1.832	0.916	0.918	0.918	30
0.2%	1	0	0	PUT	^/api/blocks/[^/]+$	1.618	1.618	1.618	1.618	0
0.1%	1	0	0	DELETE	^/api/posts/[^/]+$	1.338	1.338	1.338	1.338	0
0.1%	4	1	0	GET	/api/me	1.116	0.279	0.406	0.406	150
0.1%	1	0	0	GET	/api/me/export	1.103	1.103	1.103	1.103	588558
0.1%	1	1	0	GET	^/api/posts/[^/]+$	1.019	1.019	1.019	1.019	27
0.1%	1	0	0	POST	/api/me/account-deletion	0.941	0.941	0.941	0.941	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 41371 クエリ / 203 種類 / 合計 432.664s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	27.5%	1161	118.888	102.40	5443.49	4.3	1.0	SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE
2	25.5%	154	110.414	716.98	4195.97	3.1	1.0	SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE
3	11.3%	63	48.888	775.99	5840.19	4.8	0.0	LOCK TABLE users IN SHARE ROW EXCLUSIVE MODE
4	8.9%	1510	38.595	25.56	223.31	357.5	1.0	SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXI…
5	5.2%	38	22.283	586.39	1636.43	12274.2	122.6	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
6	2.9%	18	12.451	691.71	3580.60	32830.4	1639.2	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT gen_random_uuid()::text, id, $2, $3 FROM users WHERE region = $1 AND status = $4
7	2.8%	82	11.961	145.87	396.67	225.9	4917.1	SELECT id, username, secret_key_hash, display_name, bio, profile_image_id, region, account_type, status FROM users WHERE status=$4 AND id<>$1 AND ($2=$5 OR (use…
8	2.7%	18	11.892	660.64	3849.81	13.7	0.8	UPDATE posting_windows SET ended_at = $2 WHERE region = $1 AND ended_at IS NULL
9	2.0%	1691	8.611	5.09	47.29	87.6	1.0	SELECT u.secret_key_hash FROM image_objects i JOIN users u ON u.id = i.owner_user_id WHERE i.id = $1 AND u.status = $2
10	1.8%	151	7.973	52.80	267.85	957.6	1.0	INSERT INTO image_objects (id, owner_user_id, data, media_type, created_at) VALUES ($1, $2, $3, $5, $4)
11	1.7%	894	7.263	8.12	88.66	127.9	14.2	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
12	1.3%	270	5.685	21.05	1008.90	4.3	1.0	SELECT id FROM users WHERE id = $1 FOR NO KEY UPDATE
13	0.7%	542	3.128	5.77	52.26	109.1	5.8	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
14	0.4%	37	1.739	47.00	137.26	610.3	1924.0	SELECT user_low_id, user_high_id FROM friendships WHERE user_low_id IN ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $2…
15	0.3%	30	1.453	48.43	132.94	611.8	1957.6	SELECT user_low_id, user_high_id FROM friendships WHERE user_low_id IN ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $2…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE;
-- rank 2
SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE;
-- rank 3
LOCK TABLE users IN SHARE ROW EXCLUSIVE MODE;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 608916f0d472811c6b9cc33615e090cf8ac99f63
Type: cpu
Time: 2026-10-31 10:29:35 JST
Duration: 60s, Total samples = 46.56s (77.60%)
Showing nodes accounting for 38.73s, 83.18% of 46.56s total
Dropped 890 nodes (cum <= 0.23s)
Showing top 40 nodes out of 226
      flat  flat%   sum%        cum   cum%
    18.22s 39.13% 39.13%     18.25s 39.20%  crypto/internal/fips140/sha256.blockAVX2
     2.55s  5.48% 44.61%     12.58s 27.02%  crypto/internal/fips140/hmac.New[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
     1.99s  4.27% 48.88%      1.99s  4.27%  runtime.memclrNoHeapPointers
     1.55s  3.33% 52.21%      1.55s  3.33%  runtime.memmove
     1.30s  2.79% 55.00%      2.15s  4.62%  image/jpeg.(*decoder).reconstructBlock
     1.25s  2.68% 57.69%      1.25s  2.68%  internal/runtime/syscall/linux.Syscall6
     0.71s  1.52% 59.21%     19.75s 42.42%  crypto/internal/fips140/sha256.(*Digest).Write
     0.68s  1.46% 60.67%      3.13s  6.72%  runtime.mallocgcSmallNoscan
     0.65s  1.40% 62.07%      0.65s  1.40%  runtime.nextFreeFast (inline)
     0.63s  1.35% 63.42%     12.71s 27.30%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.60s  1.29% 64.71%      6.25s 13.42%  runtime.mallocgc
     0.59s  1.27% 65.98%      4.05s  8.70%  image/jpeg.(*decoder).processSOS
     0.57s  1.22% 67.20%      0.57s  1.22%  runtime.madvise
     0.46s  0.99% 68.19%      0.49s  1.05%  image/jpeg.(*decoder).receiveExtend
     0.45s  0.97% 69.16%      0.63s  1.35%  runtime.tryDeferToSpanScan
     0.44s  0.95% 70.10%      0.44s  0.95%  runtime.procyieldAsm
     0.39s  0.84% 70.94%      0.74s  1.59%  image/jpeg.(*decoder).decodeHuffman
     0.39s  0.84% 71.78%      1.04s  2.23%  runtime.scanObject
     0.36s  0.77% 72.55%      5.74s 12.33%  crypto/internal/fips140/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
     0.35s  0.75% 73.30%        14s 30.07%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.34s  0.73% 74.03%      0.49s  1.05%  image/jpeg.idctRows
     0.34s  0.73% 74.76%      0.69s  1.48%  runtime.scanObjectsSmall
     0.30s  0.64% 75.41%      0.30s  0.64%  crypto/internal/fips140/sha256.(*Digest).Reset
     0.30s  0.64% 76.05%     18.55s 39.84%  crypto/internal/fips140/sha256.block
     0.30s  0.64% 76.70%      0.37s  0.79%  image/jpeg.(*decoder).ensureNBits
     0.28s   0.6% 77.30%     18.09s 38.85%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.28s   0.6% 77.90%      0.28s   0.6%  runtime.(*spanInlineMarkBits).init
     0.27s  0.58% 78.48%      0.27s  0.58%  runtime.(*mspan).writeHeapBitsSmall
     0.25s  0.54% 79.02%      0.26s  0.56%  image/jpeg.dctBox (inline)
     0.24s  0.52% 79.53%      0.36s  0.77%  image/jpeg.idctCols
     0.24s  0.52% 80.05%      0.29s  0.62%  runtime.typePointers.next
     0.23s  0.49% 80.54%      2.88s  6.19%  crypto/hmac.New.UnwrapNew[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }].func1
     0.22s  0.47% 81.01%      0.37s  0.79%  runtime.gcmarknewobject
     0.18s  0.39% 81.40%      0.98s  2.10%  runtime.growslice
     0.18s  0.39% 81.79%      1.15s  2.47%  runtime.mallocgcSmallScanNoHeader
     0.17s  0.37% 82.15%      3.71s  7.97%  runtime.newobject
     0.15s  0.32% 82.47%      0.28s   0.6%  runtime.scanblock
     0.12s  0.26% 82.73%      0.32s  0.69%  runtime.(*unwinder).resolveInternal
     0.12s  0.26% 82.99%      0.35s  0.75%  runtime.pcvalue
     0.09s  0.19% 83.18%      0.29s  0.62%  crypto/internal/fips140/hmac.New[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }].func1
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 36.0% / 99% | 6% | 0.0% | 5 | 2011 MB | run queue が積もっている |
| isu2 (2 cpu) | 2.4% / 16% | 0% | 0.0% | 3 | 3205 MB |  |
| isu3 (2 cpu) | 2.8% / 18% | 0% | 0.0% | 3 | 3220 MB |  |
| isu4 (2 cpu) | 42.2% / 99% | 18% | 4.8% | 24 | 90 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 65% / nginx 5% |
| isu2 | isureal 4% |
| isu3 | isureal 4% |
| isu4 | postgres 15% |

## Diff from previous run

スコア: 2674 → 6389 (+3715)

`e240203..8cd1279`

```
 webapp/go/cmd/isureal/pprof.go | 11 +++++++++++
 1 file changed, 11 insertions(+)
```
