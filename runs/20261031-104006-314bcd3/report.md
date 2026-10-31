# run 20261031-104006-314bcd3

## Summary

| | |
|---|---|
| score | 7705 |
| status | pass |
| ref | db-drop-table-locks |
| sha | 314bcd364985 |
| 所要時間 | 379s |
| 前回比 | +30 (0.4%) vs 20261031-103328-02acb51 |
| ベスト比 | +30 (0.4%) vs 20261031-103328-02acb51 = 7675 **ベスト更新** |
| note | 投稿の投稿期間ロックをFOR SHAREに（+投稿期間の開始のadvisory lock） |

## Bench

```
{"id":"1448","status":"done","portalStatus":"finished","passed":true,"score":7705,"raw":7705,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:44:29.464Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:44:29.855Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:44:32.831Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:44:37.838Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T01:44:38.989Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T01:44:41.835Z level=INFO msg="america 2 日目: 6 人"
time=2026-10-31T01:44:46.760Z level=INFO msg="asia 3 日目: 18 人"
time=2026-10-31T01:44:48.655Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T01:44:50.875Z level=INFO msg="america 3 日目: 6 人"
time=2026-10-31T01:44:56.312Z level=INFO msg="asia 4 日目: 18 人"
time=2026-10-31T01:44:56.727Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T01:44:59.842Z level=INFO msg="america 4 日目: 6 人"
time=2026-10-31T01:45:04.402Z level=INFO msg="asia 5 日目: 18 人"
time=2026-10-31T01:45:06.017Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T01:45:08.809Z level=INFO msg="america 5 日目: 6 人"
time=2026-10-31T01:45:17.797Z level=INFO msg="america 最終日（グランプリ）: 6 人"
time=2026-10-31T01:45:17.846Z level=INFO msg="europe 最終日（グランプリ）: 6 人"
time=2026-10-31T01:45:20.204Z level=INFO msg="asia 最終日（グランプリ）: 17 人"
time=2026-10-31T01:45:26.678Z level=INFO msg="グランプリ: アジア 16 人, ヨーロッパ 6 人, アメリカ 5 人"
{
  "pass": true,
  "score": 7705,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
73.5%	125	10	0	GET	/api/notifications	2070.029	16.560	60.001	60.002	153
10.4%	1874	228	0	GET	^/api/images/[^/]+	293.712	0.157	0.436	0.560	77913
3.9%	845	0	0	GET	^/api/feeds/[^/]+	109.796	0.130	0.384	0.494	2285
1.8%	107	0	0	POST	/api/posts	51.458	0.481	1.400	1.852	594
1.8%	444	0	0	GET	/api/memories	49.676	0.112	0.367	0.654	12923
1.6%	1097	580	0	POST	/api/auth/login	46.421	0.042	0.221	0.306	134
1.2%	51	0	0	GET	/api/me/export	34.320	0.673	2.355	2.355	1763222
0.9%	116	8	0	POST	/api/friend-requests	26.714	0.230	1.836	2.146	109
0.9%	110	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	25.979	0.236	1.316	1.995	116
0.6%	18	0	0	POST	/api/admin/posting-window	16.801	0.933	2.786	2.786	177
0.5%	154	2	0	GET	^/api/posts/[^/]+$	12.730	0.083	0.222	0.246	585
0.4%	52	0	0	POST	^/api/posts/[^/]+/comments$	11.074	0.213	1.497	1.497	274
0.3%	108	1	0	GET	^/api/posts/me/current$	8.904	0.082	0.257	0.271	588
0.3%	103	91	11	PUT	/api/me/profile-image	8.348	0.081	0.261	0.357	38
0.3%	39	0	0	GET	/api/users	7.917	0.203	0.430	0.430	281
0.3%	110	0	0	GET	/api/posting-window/current	7.703	0.070	0.261	0.282	237
0.2%	52	0	0	GET	^/api/posts/[^/]+/realmojis$	5.719	0.110	0.327	0.327	29
0.2%	146	0	0	GET	^/api/relationships/[^/]+	5.197	0.036	0.210	0.239	1877
0.1%	53	0	0	GET	^/api/posts/[^/]+/comments$	4.057	0.077	0.270	0.270	209
0.1%	1	0	0	POST	/api/initialize	3.992	3.992	3.992	3.992	69
0.1%	97	0	0	POST	/api/auth/logout	3.733	0.038	0.151	0.151	0
0.1%	85	0	0	GET	^/api/memories/[^/]+	3.667	0.043	0.273	0.273	710
0.1%	85	0	0	GET	^/api/users/[^/]+$	3.300	0.039	0.152	0.152	163
0.1%	34	1	0	POST	/api/auth/register	2.940	0.086	0.277	0.277	237
0.0%	13	0	0	PATCH	^/api/posts/[^/]+$	1.255	0.097	0.226	0.226	604
0.0%	18	0	0	DELETE	^/api/comments/[^/]+$	0.828	0.046	0.243	0.243	0
0.0%	13	0	0	PUT	^/api/follows/[^/]+$	0.791	0.061	0.092	0.092	0
0.0%	5	0	0	DELETE	^/api/posts/[^/]+$	0.344	0.069	0.110	0.110	0
0.0%	4	0	0	PUT	^/api/blocks/[^/]+$	0.073	0.018	0.031	0.031	0
0.0%	1	0	0	POST	/api/admin/auth/logout	0.049	0.049	0.049	0.049	0
…（上位 30 / 全 37 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
57.1%	48	6	0	GET	/api/notifications	754.772	15.724	60.002	60.002	155
7.5%	122	0	0	GET	^/api/images/[^/]+	98.722	0.809	1.522	1.618	220728
6.7%	67	0	0	GET	^/api/feeds/[^/]+	87.837	1.311	2.000	2.000	1101
6.5%	36	0	0	POST	/api/posts	86.473	2.402	3.415	3.415	610
4.7%	67	0	0	GET	/api/memories	61.866	0.923	1.895	1.895	6315
4.2%	31	1	0	POST	/api/friend-requests	55.569	1.793	3.900	3.900	112
2.7%	26	0	0	POST	/api/auth/login	35.150	1.352	2.220	2.220	248
1.8%	39	0	0	GET	/api/posting-window/current	23.390	0.600	1.347	1.347	239
1.6%	35	4	0	GET	^/api/posts/me/current$	21.562	0.616	0.941	0.941	546
1.5%	5	0	0	POST	/api/auth/register	19.320	3.864	6.704	6.704	230
1.3%	8	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	17.205	2.151	4.448	4.448	116
1.2%	4	0	0	PUT	^/api/follows/[^/]+$	16.184	4.046	5.439	5.439	0
1.0%	17	0	0	GET	/api/users	13.215	0.777	1.051	1.051	220
0.6%	18	0	0	GET	^/api/relationships/[^/]+	7.773	0.432	0.608	0.608	1466
0.4%	4	2	0	GET	^/api/posts/[^/]+$	5.001	1.250	1.444	1.444	297
0.3%	3	2	0	PATCH	^/api/posts/[^/]+$	3.757	1.252	1.346	1.346	245
0.2%	5	0	0	GET	^/api/memories/[^/]+	3.082	0.616	0.849	0.849	963
0.2%	3	0	0	POST	/api/auth/logout	2.263	0.754	0.829	0.829	0
0.2%	5	2	0	GET	/api/me	2.026	0.405	0.507	0.507	127
0.1%	1	0	0	PUT	^/api/blocks/[^/]+$	1.718	1.718	1.718	1.718	0
0.1%	2	0	0	GET	^/api/users/[^/]+$	1.645	0.823	1.027	1.027	160
0.1%	1	0	0	GET	/api/me/export	1.142	1.142	1.142	1.142	2753246
0.1%	1	0	0	POST	/api/me/account-deletion	1.028	1.028	1.028	1.028	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
55.4%	42	5	0	GET	/api/notifications	643.340	15.318	60.001	60.001	156
8.8%	118	1	0	GET	^/api/images/[^/]+	101.811	0.863	1.578	2.454	220346
6.7%	30	0	0	POST	/api/posts	77.652	2.588	4.091	4.091	596
5.1%	36	1	0	POST	/api/friend-requests	59.415	1.650	2.713	2.713	113
4.8%	61	0	0	GET	/api/memories	56.187	0.921	1.708	1.708	6998
4.7%	38	0	0	GET	^/api/feeds/[^/]+	54.900	1.445	2.414	2.414	1106
2.7%	24	0	0	POST	/api/auth/login	30.998	1.292	1.938	1.938	252
2.1%	12	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	24.438	2.037	5.265	5.265	116
2.0%	5	0	0	POST	/api/auth/register	22.873	4.575	7.414	7.414	227
1.9%	34	0	0	GET	/api/posting-window/current	21.956	0.646	1.485	1.485	240
1.5%	4	0	0	PUT	^/api/follows/[^/]+$	17.927	4.482	6.051	6.051	0
1.3%	30	12	0	GET	^/api/posts/me/current$	15.365	0.512	0.928	0.928	372
1.1%	16	0	0	GET	/api/users	13.208	0.826	1.268	1.268	230
0.8%	19	0	0	GET	^/api/relationships/[^/]+	9.359	0.493	1.262	1.262	1435
0.2%	2	0	0	GET	/api/me/export	2.739	1.369	1.789	1.789	1809193
0.1%	2	0	0	GET	^/api/memories/[^/]+	1.673	0.837	0.953	0.953	1225
0.1%	1	0	0	PUT	^/api/blocks/[^/]+$	1.620	1.620	1.620	1.620	0
0.1%	4	1	0	GET	/api/me	1.525	0.381	0.508	0.508	150
0.1%	1	1	0	PATCH	^/api/posts/[^/]+$	1.264	1.264	1.264	1.264	30
0.1%	1	0	0	DELETE	^/api/posts/[^/]+$	1.174	1.174	1.174	1.174	0
0.1%	1	0	0	POST	/api/me/account-deletion	0.944	0.944	0.944	0.944	0
0.1%	1	0	0	POST	/api/auth/logout	0.830	0.830	0.830	0.830	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 44870 クエリ / 216 種類 / 合計 386.404s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	18.6%	1252	71.905	57.43	3746.52	4.2	1.0	SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE
2	18.6%	1889	71.786	38.00	264.34	354.8	1.0	SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXI…
3	10.9%	1	42.054	42053.98	42053.98	1845286.0	1.0	SELECT count(*), md5(string_agg(id || $1 || owner_user_id || $2 || media_type || $3 || created_at::text, $4 ORDER BY id)) FROM image_objects WHERE created_at <=…
4	9.3%	950	35.951	37.84	122.34	259.7	17.9	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
5	7.1%	54	27.467	508.65	1866.33	11120.8	115.3	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
6	7.1%	44	27.365	621.93	5756.62	8.3	0.0	LOCK TABLE users IN SHARE ROW EXCLUSIVE MODE
7	5.7%	173	22.187	128.25	459.58	999.2	1.0	INSERT INTO image_objects (id, owner_user_id, data, media_type, created_at) VALUES ($1, $2, $3, $5, $4)
8	3.6%	18	14.039	779.96	2755.21	32133.5	1639.1	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT gen_random_uuid()::text, id, $2, $3 FROM users WHERE region = $1 AND status = $4
9	3.0%	72	11.546	160.36	315.15	4597.1	4915.4	SELECT id, username, secret_key_hash, display_name, bio, profile_image_id, region, account_type, status FROM users WHERE status=$4 AND id<>$1 AND ($2=$5 OR (use…
10	3.0%	419	11.503	27.45	1444.38	4.1	1.0	SELECT id FROM users WHERE id = $1 FOR NO KEY UPDATE
11	2.7%	2102	10.623	5.05	38.88	87.3	1.0	SELECT u.secret_key_hash FROM image_objects i JOIN users u ON u.id = i.owner_user_id WHERE i.id = $1 AND u.status = $2
12	1.3%	664	4.836	7.28	77.92	100.3	6.0	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
13	0.6%	67	2.340	34.92	112.62	1027.8	2015.3	SELECT user_low_id, user_high_id FROM friendships WHERE user_low_id IN ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $2…
14	0.5%	3621	1.872	0.52	29.49	12.9	1.0	SELECT s.user_id, s.token_hash, s.created_at, u.region FROM login_sessions s JOIN users u ON u.id = s.user_id WHERE s.token_hash = $1 AND s.active = $2 AND u.st…
15	0.4%	44	1.519	34.51	88.04	1006.1	1953.6	SELECT user_low_id, user_high_id FROM friendships WHERE user_low_id IN ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $2…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE;
-- rank 2
SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXISTS ( SELECT $3 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL ) OR EXISTS ( SELECT $4 FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.autho…;
-- rank 3
SELECT count(*), md5(string_agg(id || $1 || owner_user_id || $2 || media_type || $3 || created_at::text, $4 ORDER BY id)) FROM image_objects WHERE created_at <= (SELECT created_at FROM initial_data_boundary);
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 6316c911685873d5f177951a3169a2aeda130131
Type: cpu
Time: 2026-10-31 10:43:56 JST
Duration: 60s, Total samples = 22.98s (38.30%)
Showing nodes accounting for 18.66s, 81.20% of 22.98s total
Dropped 831 nodes (cum <= 0.11s)
Showing top 40 nodes out of 268
      flat  flat%   sum%        cum   cum%
     7.34s 31.94% 31.94%      7.35s 31.98%  crypto/internal/fips140/sha256.blockAVX2
     1.29s  5.61% 37.55%      1.29s  5.61%  internal/runtime/syscall/linux.Syscall6
     1.28s  5.57% 43.12%      1.91s  8.31%  image/jpeg.(*decoder).reconstructBlock
     1.28s  5.57% 48.69%      1.28s  5.57%  runtime.memclrNoHeapPointers
     1.19s  5.18% 53.87%      1.19s  5.18%  runtime.memmove
     0.61s  2.65% 56.53%      3.63s 15.80%  image/jpeg.(*decoder).processSOS
     0.60s  2.61% 59.14%      8.41s 36.60%  crypto/internal/fips140/sha256.(*Digest).Write
     0.46s  2.00% 61.14%      0.72s  3.13%  image/jpeg.(*decoder).decodeHuffman
     0.41s  1.78% 62.92%      8.41s 36.60%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.33s  1.44% 64.36%      8.94s 38.90%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.30s  1.31% 65.67%      0.31s  1.35%  image/jpeg.(*decoder).receiveExtend
     0.28s  1.22% 66.88%      5.23s 22.76%  crypto/internal/fips140/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
     0.26s  1.13% 68.02%      0.30s  1.31%  image/jpeg.idctCols
     0.26s  1.13% 69.15%      0.26s  1.13%  runtime.madvise
     0.23s  1.00% 70.15%      9.77s 42.52%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.23s  1.00% 71.15%      0.31s  1.35%  image/jpeg.idctRows
     0.21s  0.91% 72.06%      0.21s  0.91%  runtime.procyieldAsm
     0.20s  0.87% 72.93%      0.27s  1.17%  image/jpeg.(*decoder).ensureNBits
     0.20s  0.87% 73.80%      0.33s  1.44%  runtime.tryDeferToSpanScan
     0.19s  0.83% 74.63%      0.19s  0.83%  internal/byteorder.BEUint32 (inline)
     0.17s  0.74% 75.37%      0.56s  2.44%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.15s  0.65% 76.02%      0.66s  2.87%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.15s  0.65% 76.68%      0.15s  0.65%  runtime.futex
     0.15s  0.65% 77.33%      0.32s  1.39%  runtime.scanObject
     0.13s  0.57% 77.89%      0.13s  0.57%  runtime.nextFreeFast (inline)
     0.11s  0.48% 78.37%      0.34s  1.48%  runtime.mallocgcSmallScanNoHeader
     0.09s  0.39% 78.76%      7.44s 32.38%  crypto/internal/fips140/sha256.block
     0.07s   0.3% 79.07%      0.26s  1.13%  crypto/internal/fips140deps/byteorder.BEUint32 (inline)
     0.07s   0.3% 79.37%      1.11s  4.83%  runtime.mallocgc
     0.06s  0.26% 79.63%      0.15s  0.65%  runtime.(*unwinder).resolveInternal
     0.06s  0.26% 79.90%      0.32s  1.39%  runtime.newobject
     0.05s  0.22% 80.11%      0.43s  1.87%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.05s  0.22% 80.33%      0.20s  0.87%  runtime.scanObjectsSmall
     0.03s  0.13% 80.46%      0.16s   0.7%  crypto/internal/fips140/mlkem.kemKeyGen
     0.03s  0.13% 80.59%      0.29s  1.26%  crypto/internal/fips140/sha256.consumeUint32 (inline)
     0.03s  0.13% 80.72%      0.36s  1.57%  crypto/tls.(*atLeastReader).Read
     0.03s  0.13% 80.85%      5.55s 24.15%  example.com/isureal-company/isureal.imageSignature
     0.03s  0.13% 80.98%      0.80s  3.48%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.03s  0.13% 81.11%      0.23s  1.00%  github.com/jmoiron/sqlx.fieldsByTraversal
     0.02s 0.087% 81.20%      5.41s 23.54%  github.com/jackc/pgx/v5/pgconn.(*PgConn).scramAuth
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 19.5% / 84% | 0% | 0.1% | 7 | 2862 MB | run queue が積もっている |
| isu2 (2 cpu) | 6.5% / 94% | 1% | 1.9% | 4 | 322 MB |  |
| isu3 (2 cpu) | 2.1% / 11% | 0% | 0.0% | 3 | 3212 MB |  |
| isu4 (2 cpu) | 46.1% / 100% | 39% | 3.3% | 36 | 106 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 31% / nginx 4% |
| isu2 | postgres 5% / isureal 3% |
| isu3 | isureal 3% |
| isu4 | postgres 22% |

## Diff from previous run

スコア: 7675 → 7705 (+30)

`02acb51..314bcd3`

```
 conf/isu4/etc/postgresql/16/main/conf.d/zz-tuning.conf |  7 +++++++
 webapp/go/admin_posting_window_api.go                  |  3 ++-
 webapp/go/post_api.go                                  |  5 +++--
 webapp/sql/002-indexes.sql                             | 10 ++++++++++
 webapp/sql/reset.sql                                   |  2 ++
 5 files changed, 24 insertions(+), 3 deletions(-)
```
