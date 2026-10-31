# run 20261031-104645-d630f63

## Summary

| | |
|---|---|
| score | 7550 |
| status | pass |
| ref | db-drop-table-locks |
| sha | d630f63a84c3 |
| 所要時間 | 248s |
| 前回比 | -125 (-1.6%) vs 20261031-103328-02acb51 |
| ベスト比 | -125 (-1.6%) vs 20261031-103328-02acb51 = 7675 |
| note | 登録のLOCK TABLEを一意インデックス+ON CONFLICTに |

## Bench

```
{"id":"1678","status":"done","portalStatus":"finished","passed":true,"score":7550,"raw":7550,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:49:03.051Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:49:04.039Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:49:07.035Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:49:12.547Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T01:49:12.970Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T01:49:16.012Z level=INFO msg="america 2 日目: 6 人"
time=2026-10-31T01:49:20.460Z level=INFO msg="asia 3 日目: 18 人"
time=2026-10-31T01:49:22.125Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T01:49:25.034Z level=INFO msg="america 3 日目: 6 人"
time=2026-10-31T01:49:29.659Z level=INFO msg="asia 4 日目: 19 人"
time=2026-10-31T01:49:31.890Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T01:49:34.738Z level=INFO msg="america 4 日目: 6 人"
time=2026-10-31T01:49:39.079Z level=INFO msg="asia 5 日目: 19 人"
time=2026-10-31T01:49:40.336Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T01:49:43.072Z level=INFO msg="america 5 日目: 6 人"
time=2026-10-31T01:49:52.302Z level=INFO msg="america 最終日（グランプリ）: 6 人"
time=2026-10-31T01:49:52.669Z level=INFO msg="europe 最終日（グランプリ）: 7 人"
time=2026-10-31T01:49:52.936Z level=INFO msg="asia 最終日（グランプリ）: 18 人"
time=2026-10-31T01:50:00.916Z level=INFO msg="グランプリ: アジア 16 人, ヨーロッパ 5 人, アメリカ 3 人"
{
  "pass": true,
  "score": 7550,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
64.2%	112	2	0	GET	/api/notifications	1171.526	10.460	60.000	60.002	174
15.3%	2202	249	0	GET	^/api/images/[^/]+	279.367	0.127	0.408	0.565	61637
4.0%	689	0	0	GET	^/api/feeds/[^/]+	72.552	0.105	0.343	0.452	2002
2.8%	513	0	0	GET	/api/memories	51.339	0.100	0.400	0.580	14476
2.2%	53	0	0	GET	/api/me/export	39.812	0.751	2.691	2.691	1968985
2.1%	1177	541	0	POST	/api/auth/login	37.428	0.032	0.201	0.670	149
2.0%	96	0	0	POST	/api/posts	36.418	0.379	1.228	1.228	614
1.2%	128	18	0	POST	/api/friend-requests	21.669	0.169	1.147	1.513	104
1.1%	102	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	20.258	0.199	1.881	2.474	116
0.8%	18	0	0	POST	/api/admin/posting-window	15.441	0.858	2.631	2.631	177
0.6%	54	0	0	POST	^/api/posts/[^/]+/realmojis$	10.615	0.197	1.300	1.300	459
0.5%	102	94	5	PUT	/api/me/profile-image	8.251	0.081	0.253	0.274	30
0.4%	44	0	0	GET	/api/users	7.553	0.172	0.335	0.335	265
0.3%	97	1	0	GET	^/api/posts/me/current$	6.160	0.064	0.332	0.332	608
0.3%	110	0	0	GET	^/api/posts/[^/]+$	6.016	0.055	0.239	0.426	605
0.3%	103	0	0	GET	^/api/posts/[^/]+/realmojis$	5.827	0.057	0.183	0.190	276
0.3%	103	0	0	GET	/api/posting-window/current	5.640	0.055	0.207	0.333	237
0.3%	1	0	0	POST	/api/initialize	4.613	4.613	4.613	4.613	69
0.2%	117	0	0	GET	^/api/relationships/[^/]+	4.041	0.035	0.205	0.234	2231
0.2%	123	0	0	POST	/api/auth/logout	3.431	0.028	0.186	0.193	0
0.2%	117	0	0	GET	^/api/memories/[^/]+	3.237	0.028	0.157	0.206	699
0.2%	75	0	0	GET	^/api/users/[^/]+$	2.936	0.039	0.191	0.191	177
0.1%	29	0	0	GET	^/api/posts/[^/]+/comments$	2.485	0.086	0.248	0.248	29
0.1%	4	0	0	DELETE	^/api/posts/[^/]+$	2.325	0.581	1.314	1.314	0
0.1%	35	1	0	POST	/api/auth/register	1.933	0.055	0.164	0.164	237
0.0%	15	0	0	PUT	^/api/follows/[^/]+$	0.810	0.054	0.097	0.097	0
0.0%	9	0	0	PATCH	^/api/posts/[^/]+$	0.794	0.088	0.298	0.298	634
0.0%	16	0	0	DELETE	^/api/realmojis/[^/]+$	0.594	0.037	0.121	0.121	0
0.0%	1	0	0	DELETE	^/api/friendships/[^/]+$	0.443	0.443	0.443	0.443	0
0.0%	4	0	0	PUT	^/api/blocks/[^/]+$	0.406	0.101	0.349	0.349	0
…（上位 30 / 全 39 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
52.4%	46	3	0	GET	/api/notifications	581.987	12.652	60.274	60.274	165
9.6%	120	3	0	GET	^/api/images/[^/]+	106.238	0.885	1.652	1.664	178441
7.6%	33	0	0	POST	/api/posts	84.218	2.552	4.036	4.036	602
5.7%	43	0	0	GET	^/api/feeds/[^/]+	63.644	1.480	2.568	2.568	835
5.5%	63	0	0	GET	/api/memories	61.451	0.975	2.608	2.608	8746
4.5%	31	2	0	POST	/api/friend-requests	49.611	1.600	2.679	2.679	110
3.4%	27	0	0	POST	/api/auth/login	37.694	1.396	2.025	2.025	253
2.0%	38	0	0	GET	/api/posting-window/current	22.524	0.593	1.524	1.524	239
1.9%	30	10	0	GET	^/api/posts/me/current$	20.715	0.690	1.733	1.733	418
1.8%	9	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	19.458	2.162	4.783	4.783	116
1.5%	4	0	0	PUT	^/api/follows/[^/]+$	16.874	4.219	5.667	5.667	0
1.0%	11	0	0	GET	/api/users	10.960	0.996	1.265	1.265	300
0.8%	19	0	0	GET	^/api/relationships/[^/]+	9.228	0.486	1.271	1.271	1158
0.6%	4	0	0	DELETE	^/api/posts/[^/]+$	6.596	1.649	2.282	2.282	0
0.5%	5	0	0	POST	/api/auth/register	6.051	1.210	1.340	1.340	227
0.5%	6	0	0	POST	/api/auth/logout	5.086	0.848	0.940	0.940	0
0.2%	5	2	0	GET	/api/me	2.346	0.469	0.509	0.509	125
0.1%	1	0	0	PUT	^/api/blocks/[^/]+$	1.617	1.617	1.617	1.617	0
0.1%	1	0	0	GET	/api/me/export	1.566	1.566	1.566	1.566	3097037
0.1%	1	0	0	POST	/api/me/account-deletion	1.288	1.288	1.288	1.288	0
0.1%	1	0	0	GET	^/api/users/[^/]+$	0.822	0.822	0.822	0.822	153
0.1%	1	0	0	GET	^/api/posts/[^/]+$	0.820	0.820	0.820	0.820	717
0.1%	1	0	0	GET	^/api/memories/[^/]+	0.641	0.641	0.641	0.641	29

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
56.7%	41	4	0	GET	/api/notifications	585.061	14.270	60.001	60.001	160
9.3%	122	2	0	GET	^/api/images/[^/]+	96.107	0.788	1.550	1.803	174189
5.9%	26	0	0	POST	/api/posts	60.536	2.328	3.239	3.239	588
5.8%	38	4	0	POST	/api/friend-requests	60.114	1.582	3.428	3.428	106
4.9%	56	0	0	GET	/api/memories	50.423	0.900	1.698	1.698	8125
4.2%	30	0	0	GET	^/api/feeds/[^/]+	43.855	1.462	2.369	2.369	969
3.2%	24	0	0	POST	/api/auth/login	32.624	1.359	1.947	1.947	247
2.0%	4	0	0	PUT	^/api/follows/[^/]+$	20.750	5.188	7.364	7.364	0
1.7%	29	0	0	GET	/api/posting-window/current	17.054	0.588	1.548	1.548	240
1.6%	21	0	0	GET	/api/users	16.073	0.765	1.065	1.065	288
1.4%	25	7	0	GET	^/api/posts/me/current$	14.832	0.593	0.956	0.956	434
0.7%	17	0	0	GET	^/api/relationships/[^/]+	7.089	0.417	0.610	0.610	1180
0.7%	4	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	7.008	1.752	1.924	1.924	116
0.6%	5	0	0	POST	/api/auth/register	6.700	1.340	1.344	1.344	230
0.3%	3	0	0	GET	^/api/memories/[^/]+	2.924	0.975	1.136	1.136	1809
0.3%	2	0	0	DELETE	^/api/posts/[^/]+$	2.911	1.456	1.573	1.573	0
0.2%	2	1	0	PATCH	^/api/posts/[^/]+$	2.469	1.235	1.332	1.332	321
0.2%	1	0	0	PUT	^/api/blocks/[^/]+$	1.824	1.824	1.824	1.824	0
0.1%	1	0	0	GET	/api/me/export	1.536	1.536	1.536	1.536	2913042
0.1%	1	0	0	POST	/api/me/account-deletion	1.037	1.037	1.037	1.037	0
0.1%	4	1	0	GET	/api/me	1.019	0.255	0.305	0.305	153
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 45066 クエリ / 200 種類 / 合計 266.457s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	24.6%	2199	65.564	29.82	271.46	352.2	1.0	SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXI…
2	21.7%	1487	57.789	38.86	5019.98	4.2	1.0	SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE
3	12.8%	55	34.177	621.41	2297.87	12460.5	129.6	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
4	9.0%	762	24.099	31.63	101.04	251.5	15.4	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
5	5.2%	155	13.838	89.28	331.99	954.7	1.0	INSERT INTO image_objects (id, owner_user_id, data, media_type, created_at) VALUES ($1, $2, $3, $5, $4)
6	4.9%	2432	13.070	5.37	51.80	87.4	1.0	SELECT u.secret_key_hash FROM image_objects i JOIN users u ON u.id = i.owner_user_id WHERE i.id = $1 AND u.status = $2
7	4.2%	18	11.245	624.74	2574.36	32205.0	1639.2	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT gen_random_uuid()::text, id, $2, $3 FROM users WHERE region = $1 AND status = $4
8	3.7%	76	9.981	131.33	322.63	4597.2	4915.2	SELECT id, username, secret_key_hash, display_name, bio, profile_image_id, region, account_type, status FROM users WHERE status=$4 AND id<>$1 AND ($2=$5 OR (use…
9	2.1%	321	5.524	17.21	1128.21	4.2	1.0	SELECT id FROM users WHERE id = $1 FOR NO KEY UPDATE
10	1.4%	753	3.616	4.80	38.31	117.0	6.4	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
11	1.0%	18	2.547	141.48	722.50	14.0	0.8	UPDATE posting_windows SET ended_at = $2 WHERE region = $1 AND ended_at IS NULL
12	0.6%	3455	1.674	0.48	34.30	14.4	1.0	SELECT s.user_id, s.token_hash, s.created_at, u.region FROM login_sessions s JOIN users u ON u.id = s.user_id WHERE s.token_hash = $1 AND s.active = $2 AND u.st…
13	0.5%	1	1.351	1350.67	1350.67	29346.0	68.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.5%	6571	1.234	0.19	24.92	17.1	0.0	SELECT n.id, n.posting_window_id, w.region, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE n.user_id = $1 AND n.cr…
15	0.3%	1	0.836	835.72	835.72	11654.0	125.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXISTS ( SELECT $3 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL ) OR EXISTS ( SELECT $4 FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.autho…;
-- rank 2
SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE;
-- rank 3
SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time_deadline, w.ended_at AS window_ended_at, i.data AS image_data FROM posts p JOIN posting_windows w ON w.id = p.posting_window_id JOIN image_objects i ON i.id = p.image_id WHERE p.author_id = $1 AND p.deleted_at IS NULL ORDER BY p.posted_at,…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 2e8c011cbb9a089ab9854d8341bba46f8a820968
Type: cpu
Time: 2026-10-31 10:48:36 JST
Duration: 60s, Total samples = 27.90s (46.50%)
Showing nodes accounting for 22.96s, 82.29% of 27.90s total
Dropped 897 nodes (cum <= 0.14s)
Showing top 40 nodes out of 262
      flat  flat%   sum%        cum   cum%
     9.84s 35.27% 35.27%      9.85s 35.30%  crypto/internal/fips140/sha256.blockAVX2
     1.54s  5.52% 40.79%      1.54s  5.52%  internal/runtime/syscall/linux.Syscall6
     1.51s  5.41% 46.20%      1.51s  5.41%  runtime.memclrNoHeapPointers
     1.37s  4.91% 51.11%      2.14s  7.67%  image/jpeg.(*decoder).reconstructBlock
     1.31s  4.70% 55.81%      1.31s  4.70%  runtime.memmove
     0.65s  2.33% 58.14%        11s 39.43%  crypto/internal/fips140/sha256.(*Digest).Write
     0.65s  2.33% 60.47%     11.25s 40.32%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.59s  2.11% 62.58%      4.03s 14.44%  image/jpeg.(*decoder).processSOS
     0.50s  1.79% 64.37%      0.84s  3.01%  image/jpeg.(*decoder).decodeHuffman
     0.43s  1.54% 65.91%     11.92s 42.72%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.38s  1.36% 67.28%      0.41s  1.47%  image/jpeg.(*decoder).receiveExtend
     0.37s  1.33% 68.60%     12.72s 45.59%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.32s  1.15% 69.75%      0.39s  1.40%  image/jpeg.idctCols
     0.32s  1.15% 70.90%      0.32s  1.15%  runtime.madvise
     0.28s  1.00% 71.90%      5.88s 21.08%  crypto/internal/fips140/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
     0.28s  1.00% 72.90%      0.36s  1.29%  image/jpeg.(*decoder).ensureNBits
     0.26s  0.93% 73.84%      0.38s  1.36%  image/jpeg.idctRows
     0.23s  0.82% 74.66%      0.23s  0.82%  runtime.futex
     0.21s  0.75% 75.41%      0.21s  0.75%  runtime.procyieldAsm
     0.20s  0.72% 76.13%      0.35s  1.25%  runtime.tryDeferToSpanScan
     0.19s  0.68% 76.81%      0.19s  0.68%  image/jpeg.dctBox (inline)
     0.17s  0.61% 77.42%      0.49s  1.76%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.17s  0.61% 78.03%      0.17s  0.61%  runtime.nextFreeFast (inline)
     0.13s  0.47% 78.49%      0.45s  1.61%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.13s  0.47% 78.96%      9.98s 35.77%  crypto/internal/fips140/sha256.block
     0.12s  0.43% 79.39%      0.52s  1.86%  runtime.mallocgcSmallScanNoHeader
     0.12s  0.43% 79.82%      0.35s  1.25%  runtime.scanObject
     0.07s  0.25% 80.07%      7.91s 28.35%  example.com/isureal-company/isureal.imageSignature
     0.07s  0.25% 80.32%      0.39s  1.40%  runtime.newobject
     0.07s  0.25% 80.57%      0.18s  0.65%  runtime.pcvalue
     0.06s  0.22% 80.79%      0.17s  0.61%  runtime.(*stkframe).getStackMap
     0.06s  0.22% 81.00%      1.35s  4.84%  runtime.mallocgc
     0.06s  0.22% 81.22%      0.18s  0.65%  runtime.scanObjectsSmall
     0.05s  0.18% 81.40%      0.32s  1.15%  crypto/internal/fips140/hmac.(*HMAC).Write (partial-inline)
     0.05s  0.18% 81.58%      0.82s  2.94%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.05s  0.18% 81.76%      0.14s   0.5%  runtime.(*unwinder).resolveInternal
     0.04s  0.14% 81.90%      0.14s   0.5%  crypto/internal/fips140deps/byteorder.BEUint32 (inline)
     0.04s  0.14% 82.04%      0.18s  0.65%  database/sql.convertAssignRows
     0.04s  0.14% 82.19%      0.28s  1.00%  runtime.mallocgcSmallNoscan
     0.03s  0.11% 82.29%      0.24s  0.86%  github.com/jackc/pgx/v5.(*baseRows).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 21.6% / 82% | 0% | 0.1% | 5 | 2981 MB | run queue が積もっている |
| isu2 (2 cpu) | 93.5% / 99% | 83% | 78.9% | 6 | 168 MB | CPU が飽和 |
| isu3 (2 cpu) | 2.2% / 8% | 0% | 0.0% | 3 | 3215 MB |  |
| isu4 (2 cpu) | 49.3% / 100% | 39% | 5.7% | 30 | 104 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（CPU が飽和、主に image-export）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 35% / nginx 5% |
| isu2 | image-export 14% / postgres 7% / isureal 3% / kworker/u8:1+flush-259:0 2% / kswapd0 2% |
| isu3 | isureal 3% |
| isu4 | postgres 18% |

## Diff from previous run

スコア: 7675 → 7550 (-125)

`02acb51..d630f63`

```
 conf/isu4/etc/postgresql/16/main/conf.d/zz-tuning.conf |  7 +++++++
 webapp/go/admin_posting_window_api.go                  |  3 ++-
 webapp/go/post_api.go                                  |  5 +++--
 webapp/go/user_api.go                                  | 18 +++++++-----------
 webapp/sql/002-indexes.sql                             | 10 ++++++++++
 webapp/sql/reset.sql                                   |  2 ++
 6 files changed, 31 insertions(+), 14 deletions(-)
```
