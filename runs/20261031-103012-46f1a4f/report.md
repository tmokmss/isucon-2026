# run 20261031-103012-46f1a4f

## Summary

| | |
|---|---|
| score | 6820 |
| status | pass |
| ref | db-indexes-pgconf |
| sha | 46f1a4f9b14c |
| 所要時間 | 248s |
| 前回比 | +1399 (25.8%) vs 20261031-102650-00ef856 |
| ベスト比 | +1399 (25.8%) vs 20261031-102650-00ef856 = 5421 **ベスト更新** |
| note | PostgreSQLの設定（shared_buffers 1GBほか） |

## Bench

```
{"id":"932","status":"done","portalStatus":"finished","passed":true,"score":6820,"raw":6820,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:32:31.120Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:32:32.948Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:32:35.907Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:32:41.170Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T01:32:44.946Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T01:32:45.955Z level=INFO msg="america 2 日目: 6 人"
time=2026-10-31T01:32:51.036Z level=INFO msg="asia 3 日目: 18 人"
time=2026-10-31T01:32:54.754Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T01:32:56.261Z level=INFO msg="america 3 日目: 6 人"
time=2026-10-31T01:32:58.048Z level=INFO msg="asia 4 日目: 18 人"
time=2026-10-31T01:33:01.211Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T01:33:04.625Z level=INFO msg="america 4 日目: 6 人"
time=2026-10-31T01:33:09.332Z level=INFO msg="asia 5 日目: 19 人"
time=2026-10-31T01:33:09.376Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T01:33:12.394Z level=INFO msg="america 5 日目: 6 人"
time=2026-10-31T01:33:20.660Z level=INFO msg="europe 最終日（グランプリ）: 7 人"
time=2026-10-31T01:33:21.363Z level=INFO msg="asia 最終日（グランプリ）: 18 人"
time=2026-10-31T01:33:21.451Z level=INFO msg="america 最終日（グランプリ）: 6 人"
time=2026-10-31T01:33:29.627Z level=INFO msg="グランプリ: アジア 18 人, ヨーロッパ 2 人, アメリカ 3 人"
{
  "pass": true,
  "score": 6820,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
54.4%	109	0	0	GET	/api/notifications	1032.858	9.476	13.029	13.581	177
12.1%	1509	235	0	GET	^/api/images/[^/]+	229.202	0.152	0.439	0.525	104164
6.4%	758	0	0	GET	^/api/feeds/[^/]+	120.930	0.160	0.445	0.600	2018
5.4%	106	0	0	POST	/api/posts	102.142	0.964	1.815	1.886	600
4.1%	348	0	0	GET	/api/memories	77.795	0.224	0.904	1.094	12389
3.2%	972	579	0	POST	/api/auth/login	60.238	0.062	0.312	0.459	119
1.9%	40	0	0	GET	/api/me/export	36.940	0.923	3.377	3.377	1658504
1.6%	18	0	0	POST	/api/admin/posting-window	31.224	1.735	4.126	4.126	177
1.4%	107	9	0	POST	/api/friend-requests	25.692	0.240	1.280	1.952	108
1.3%	92	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	25.263	0.275	2.497	2.497	116
1.2%	183	2	0	GET	^/api/posts/[^/]+$	22.534	0.123	0.406	0.435	588
1.1%	35	1	0	POST	/api/auth/register	20.626	0.589	3.902	3.902	237
0.8%	99	96	2	PUT	/api/me/profile-image	14.306	0.145	0.557	0.557	26
0.6%	110	0	0	GET	/api/posting-window/current	11.114	0.101	0.374	0.420	237
0.5%	43	0	0	POST	^/api/posts/[^/]+/realmojis$	10.420	0.242	0.969	0.969	415
0.5%	107	1	0	GET	^/api/posts/me/current$	9.139	0.085	0.278	0.299	595
0.4%	49	0	0	GET	/api/users	7.486	0.153	0.380	0.380	267
0.4%	59	0	0	GET	^/api/posts/[^/]+/realmojis$	7.349	0.125	0.361	0.361	253
0.4%	35	0	0	POST	^/api/posts/[^/]+/comments$	7.198	0.206	0.650	0.650	285
0.3%	108	0	0	GET	^/api/relationships/[^/]+	6.626	0.061	0.293	0.305	2306
0.3%	6	0	0	PATCH	/api/me	6.139	1.023	3.527	3.527	212
0.3%	80	0	0	GET	^/api/users/[^/]+$	6.008	0.075	0.282	0.282	172
0.3%	51	0	0	GET	^/api/posts/[^/]+/comments$	5.113	0.100	0.273	0.273	243
0.3%	73	0	0	POST	/api/auth/logout	4.806	0.066	0.305	0.305	0
0.2%	72	0	0	GET	^/api/memories/[^/]+	4.410	0.061	0.297	0.297	509
0.1%	1	0	0	POST	/api/initialize	2.479	2.479	2.479	2.479	69
0.1%	7	0	0	DELETE	^/api/comments/[^/]+$	2.002	0.286	1.300	1.300	0
0.1%	11	0	0	PATCH	^/api/posts/[^/]+$	1.899	0.173	1.144	1.144	583
0.1%	7	0	0	DELETE	^/api/realmojis/[^/]+$	1.346	0.192	0.856	0.856	0
0.1%	8	0	0	DELETE	^/api/posts/[^/]+$	1.154	0.144	0.346	0.346	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.5%	43	0	0	GET	/api/notifications	398.855	9.276	12.687	12.687	177
13.3%	36	0	0	POST	/api/posts	128.327	3.565	8.216	8.216	602
9.7%	104	2	0	GET	^/api/images/[^/]+	93.742	0.901	1.733	1.789	228384
6.4%	46	0	0	GET	^/api/feeds/[^/]+	61.672	1.341	2.575	2.575	515
5.7%	61	0	0	GET	/api/memories	54.688	0.897	1.826	1.826	8025
5.2%	31	0	0	POST	/api/friend-requests	50.443	1.627	2.647	2.647	115
3.2%	24	0	0	POST	/api/auth/login	30.566	1.274	1.939	1.939	247
2.3%	40	0	0	GET	/api/posting-window/current	22.437	0.561	1.468	1.468	239
2.3%	12	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	21.696	1.808	2.847	2.847	116
2.0%	5	0	0	POST	/api/auth/register	19.515	3.903	7.314	7.314	228
1.9%	4	0	0	PUT	^/api/follows/[^/]+$	18.612	4.653	6.347	6.347	0
1.8%	32	9	0	GET	^/api/posts/me/current$	17.377	0.543	0.813	0.813	447
1.2%	25	0	0	GET	^/api/relationships/[^/]+	11.371	0.455	1.119	1.119	1020
0.9%	11	0	0	GET	/api/users	8.977	0.816	1.364	1.364	198
0.6%	5	0	0	POST	/api/auth/logout	5.451	1.090	1.607	1.607	0
0.4%	4	3	0	PATCH	^/api/posts/[^/]+$	4.086	1.022	1.332	1.332	167
0.3%	4	0	0	GET	^/api/memories/[^/]+	3.126	0.781	0.962	0.962	2061
0.3%	2	0	0	GET	/api/me/export	2.801	1.401	1.645	1.645	1697984
0.3%	2	0	0	DELETE	^/api/posts/[^/]+$	2.574	1.287	1.339	1.339	0
0.2%	5	2	0	GET	/api/me	1.925	0.385	0.508	0.508	127
0.2%	1	0	0	PUT	^/api/blocks/[^/]+$	1.820	1.820	1.820	1.820	0
0.1%	1	0	0	GET	^/api/posts/[^/]+$	1.324	1.324	1.324	1.324	593
0.1%	1	0	0	POST	/api/me/account-deletion	0.812	0.812	0.812	0.812	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.8%	37	0	0	GET	/api/notifications	346.945	9.377	12.787	12.787	177
12.1%	121	1	0	GET	^/api/images/[^/]+	97.962	0.810	1.429	1.542	154031
10.3%	27	0	0	POST	/api/posts	83.897	3.107	5.646	5.646	593
6.6%	59	0	0	GET	/api/memories	53.427	0.906	1.980	1.980	12774
4.7%	28	0	0	GET	^/api/feeds/[^/]+	38.102	1.361	1.749	1.749	781
4.7%	25	3	0	POST	/api/friend-requests	37.839	1.514	2.232	2.232	105
4.0%	24	0	0	POST	/api/auth/login	32.814	1.367	1.955	1.955	255
3.5%	5	0	0	POST	/api/auth/register	28.747	5.749	8.022	8.022	231
2.3%	30	0	0	GET	/api/posting-window/current	18.926	0.631	1.471	1.471	240
2.0%	4	0	0	PUT	^/api/follows/[^/]+$	16.196	4.049	5.443	5.443	0
1.7%	25	10	0	GET	^/api/posts/me/current$	13.996	0.560	0.937	0.937	370
1.3%	12	0	0	GET	/api/users	10.268	0.856	1.254	1.254	256
1.2%	3	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	9.922	3.307	4.556	4.556	116
0.8%	16	0	0	GET	^/api/relationships/[^/]+	6.829	0.427	0.612	0.612	833
0.5%	2	0	0	GET	/api/me/export	3.702	1.851	1.852	1.852	5610978
0.5%	3	0	0	POST	/api/auth/logout	3.661	1.220	1.916	1.916	0
0.2%	1	0	0	POST	^/api/posts/[^/]+/comments$	1.965	1.965	1.965	1.965	277
0.2%	1	0	0	PUT	^/api/blocks/[^/]+$	1.721	1.721	1.721	1.721	0
0.2%	4	1	0	GET	/api/me	1.521	0.380	0.508	0.508	153
0.2%	1	0	0	DELETE	^/api/posts/[^/]+$	1.452	1.452	1.452	1.452	0
0.1%	1	0	0	POST	/api/me/account-deletion	1.024	1.024	1.024	1.024	0
0.1%	1	0	0	GET	^/api/memories/[^/]+	0.624	0.624	0.624	0.624	29
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 41830 クエリ / 204 種類 / 合計 387.072s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	30.2%	172	116.953	679.96	5899.09	3.1	1.0	SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE
2	15.3%	1132	59.246	52.34	4009.40	4.3	1.0	SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE
3	13.9%	63	53.696	852.31	6371.83	6.4	0.0	LOCK TABLE users IN SHARE ROW EXCLUSIVE MODE
4	11.1%	1505	42.898	28.50	309.87	358.1	1.0	SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXI…
5	5.7%	832	21.914	26.34	85.71	254.0	16.2	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
6	5.1%	44	19.557	444.47	1555.08	11377.0	118.1	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
7	3.3%	18	12.952	719.55	3564.43	32817.6	1639.2	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT gen_random_uuid()::text, id, $2, $3 FROM users WHERE region = $1 AND status = $4
8	3.0%	18	11.788	654.91	3116.37	13.4	0.8	UPDATE posting_windows SET ended_at = $2 WHERE region = $1 AND ended_at IS NULL
9	2.4%	169	9.323	55.17	304.75	949.2	1.0	INSERT INTO image_objects (id, owner_user_id, data, media_type, created_at) VALUES ($1, $2, $3, $5, $4)
10	2.0%	72	7.932	110.17	316.13	4595.6	4915.3	SELECT id, username, secret_key_hash, display_name, bio, profile_image_id, region, account_type, status FROM users WHERE status=$4 AND id<>$1 AND ($2=$5 OR (use…
11	1.8%	362	7.117	19.66	1229.47	4.3	1.0	SELECT id FROM users WHERE id = $1 FOR NO KEY UPDATE
12	1.5%	1718	5.726	3.33	27.84	87.4	1.0	SELECT u.secret_key_hash FROM image_objects i JOIN users u ON u.id = i.owner_user_id WHERE i.id = $1 AND u.status = $2
13	0.6%	545	2.343	4.30	31.45	109.4	6.0	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
14	0.2%	3326	0.959	0.29	17.27	9.5	1.0	SELECT s.user_id, s.token_hash, s.created_at, u.region FROM login_sessions s JOIN users u ON u.id = s.user_id WHERE s.token_hash = $1 AND s.active = $2 AND u.st…
15	0.2%	6483	0.855	0.13	16.12	15.4	0.0	SELECT n.id, n.posting_window_id, w.region, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE n.user_id = $1 AND n.cr…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE;
-- rank 2
SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE;
-- rank 3
LOCK TABLE users IN SHARE ROW EXCLUSIVE MODE;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: ec07723ca656ab0fccca1d0f01056f081e84b563
Type: cpu
Time: 2026-10-31 10:32:00 JST
Duration: 60.03s, Total samples = 43.79s (72.95%)
Showing nodes accounting for 36.13s, 82.51% of 43.79s total
Dropped 927 nodes (cum <= 0.22s)
Showing top 40 nodes out of 232
      flat  flat%   sum%        cum   cum%
    17.19s 39.26% 39.26%     17.22s 39.32%  crypto/internal/fips140/sha256.blockAVX2
     2.18s  4.98% 44.23%     10.82s 24.71%  crypto/internal/fips140/hmac.New[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
     1.85s  4.22% 48.46%      1.85s  4.22%  runtime.memclrNoHeapPointers
     1.63s  3.72% 52.18%      2.20s  5.02%  image/jpeg.(*decoder).reconstructBlock
     1.51s  3.45% 55.63%      1.51s  3.45%  runtime.memmove
     1.08s  2.47% 58.10%      1.08s  2.47%  internal/runtime/syscall/linux.Syscall6
     0.67s  1.53% 59.63%      4.35s  9.93%  image/jpeg.(*decoder).processSOS
     0.67s  1.53% 61.16%      0.67s  1.53%  runtime.nextFreeFast (inline)
     0.64s  1.46% 62.62%     18.55s 42.36%  crypto/internal/fips140/sha256.(*Digest).Write
     0.58s  1.32% 63.94%      5.84s 13.34%  runtime.mallocgc
     0.55s  1.26% 65.20%      2.85s  6.51%  runtime.mallocgcSmallNoscan
     0.54s  1.23% 66.43%      0.95s  2.17%  image/jpeg.(*decoder).decodeHuffman
     0.54s  1.23% 67.66%      0.70s  1.60%  runtime.tryDeferToSpanScan
     0.49s  1.12% 68.78%      1.09s  2.49%  runtime.scanObject
     0.45s  1.03% 69.81%      0.46s  1.05%  image/jpeg.(*decoder).receiveExtend
     0.44s  1.00% 70.82%     11.88s 27.13%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.42s  0.96% 71.77%      0.42s  0.96%  runtime.madvise
     0.38s  0.87% 72.64%     13.14s 30.01%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.30s  0.69% 73.33%      5.72s 13.06%  crypto/internal/fips140/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
     0.30s  0.69% 74.01%     17.52s 40.01%  crypto/internal/fips140/sha256.block
     0.29s  0.66% 74.67%      0.42s  0.96%  image/jpeg.(*decoder).ensureNBits
     0.29s  0.66% 75.34%      0.29s  0.66%  runtime.procyieldAsm
     0.27s  0.62% 75.95%     17.19s 39.26%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.26s  0.59% 76.55%      0.26s  0.59%  crypto/internal/fips140/sha256.(*Digest).Reset
     0.26s  0.59% 77.14%      0.39s  0.89%  runtime.gcmarknewobject
     0.26s  0.59% 77.73%      0.63s  1.44%  runtime.scanObjectsSmall
     0.22s   0.5% 78.24%      0.22s   0.5%  runtime.(*spanInlineMarkBits).init
     0.21s  0.48% 78.72%      0.24s  0.55%  image/jpeg.idctRows
     0.21s  0.48% 79.20%      0.24s  0.55%  runtime.typePointers.next
     0.20s  0.46% 79.65%      0.31s  0.71%  image/jpeg.idctCols
     0.20s  0.46% 80.11%      0.30s  0.69%  runtime.scanblock
     0.18s  0.41% 80.52%      2.55s  5.82%  crypto/hmac.New.UnwrapNew[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }].func1
     0.18s  0.41% 80.93%         1s  2.28%  runtime.mallocgcSmallScanNoHeader
     0.17s  0.39% 81.32%      0.42s  0.96%  runtime.pcvalue
     0.13s   0.3% 81.62%      0.36s  0.82%  runtime.(*unwinder).resolveInternal
     0.10s  0.23% 81.85%      3.23s  7.38%  runtime.newobject
     0.08s  0.18% 82.03%      0.27s  0.62%  crypto/internal/fips140/hmac.New[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }].func1
     0.07s  0.16% 82.19%      0.22s   0.5%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.07s  0.16% 82.35%      0.94s  2.15%  runtime.growslice
     0.07s  0.16% 82.51%      1.45s  3.31%  runtime.makeslice
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 35.3% / 99% | 11% | 0.1% | 7 | 2021 MB | run queue が積もっている |
| isu2 (2 cpu) | 2.3% / 21% | 0% | 0.1% | 3 | 3153 MB |  |
| isu3 (2 cpu) | 3.0% / 25% | 0% | 0.0% | 2 | 3213 MB |  |
| isu4 (2 cpu) | 41.5% / 99% | 17% | 3.7% | 31 | 166 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 63% / nginx 4% |
| isu2 | isureal 4% |
| isu3 | isureal 5% |
| isu4 | postgres 15% |

## Diff from previous run

スコア: 5421 → 6820 (+1399)

`00ef856..46f1a4f`

```
 conf/isu4/etc/postgresql/16/main/conf.d/zz-tuning.conf |  7 +++++++
 webapp/go/cmd/isureal/pprof.go                         | 11 +++++++++++
 2 files changed, 18 insertions(+)
```
