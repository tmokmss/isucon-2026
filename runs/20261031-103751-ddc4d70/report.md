# run 20261031-103751-ddc4d70

## Summary

| | |
|---|---|
| score | 8141 |
| status | pass |
| ref | HEAD |
| sha | ddc4d701940c |
| 所要時間 | 246s |
| 前回比 | +466 (6.1%) vs 20261031-103328-02acb51 |
| ベスト比 | +466 (6.1%) vs 20261031-103328-02acb51 = 7675 **ベスト更新** |
| note | main: 署名の修正 + インデックスと PostgreSQL の設定（#60） |

## Bench

```
{"id":"1270","status":"done","portalStatus":"finished","passed":true,"score":8141,"raw":8141,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:40:02.667Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:40:03.989Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:40:07.058Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:40:12.187Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T01:40:14.614Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T01:40:17.244Z level=INFO msg="america 2 日目: 6 人"
time=2026-10-31T01:40:20.774Z level=INFO msg="asia 3 日目: 21 人"
time=2026-10-31T01:40:22.888Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T01:40:26.139Z level=INFO msg="america 3 日目: 6 人"
time=2026-10-31T01:40:30.780Z level=INFO msg="asia 4 日目: 30 人"
time=2026-10-31T01:40:33.570Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T01:40:34.259Z level=INFO msg="america 4 日目: 6 人"
time=2026-10-31T01:40:41.873Z level=INFO msg="asia 5 日目: 31 人"
time=2026-10-31T01:40:44.894Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T01:40:45.059Z level=INFO msg="america 5 日目: 6 人"
time=2026-10-31T01:40:53.465Z level=INFO msg="america 最終日（グランプリ）: 6 人"
time=2026-10-31T01:40:55.453Z level=INFO msg="asia 最終日（グランプリ）: 30 人"
time=2026-10-31T01:40:58.092Z level=INFO msg="europe 最終日（グランプリ）: 7 人"
time=2026-10-31T01:41:00.879Z level=INFO msg="グランプリ: アジア 25 人, ヨーロッパ 0 人, アメリカ 1 人"
{
  "pass": true,
  "score": 8141,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
74.1%	168	11	0	GET	/api/notifications	3154.476	18.777	60.001	60.001	144
8.3%	1961	240	0	GET	^/api/images/[^/]+	354.523	0.181	0.481	0.604	100935
4.0%	1105	0	0	GET	^/api/feeds/[^/]+	168.735	0.153	0.408	0.540	1913
2.8%	130	0	0	POST	/api/posts	118.763	0.914	2.557	2.742	614
1.5%	466	0	0	GET	/api/memories	62.001	0.133	0.549	0.607	13464
1.4%	1196	672	0	POST	/api/auth/login	58.685	0.049	0.237	0.709	126
1.1%	47	1	0	POST	/api/auth/register	46.091	0.981	5.991	5.991	234
1.0%	18	0	0	POST	/api/admin/posting-window	41.162	2.287	6.213	6.213	177
0.8%	46	0	0	GET	/api/me/export	33.648	0.731	2.152	2.152	1919663
0.8%	131	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	33.172	0.253	2.911	5.434	115
0.7%	173	9	0	POST	/api/friend-requests	28.228	0.163	1.773	2.432	111
0.6%	266	2	0	GET	^/api/posts/[^/]+$	24.779	0.093	0.350	0.434	596
0.4%	24	0	0	PUT	^/api/follows/[^/]+$	15.583	0.649	4.982	4.982	0
0.3%	121	93	18	PUT	/api/me/profile-image	14.418	0.119	0.581	2.094	42
0.3%	71	0	0	GET	/api/users	13.847	0.195	0.455	0.455	238
0.3%	127	1	0	GET	^/api/posts/me/current$	10.749	0.085	0.301	0.360	610
0.2%	42	0	0	POST	^/api/posts/[^/]+/realmojis$	9.602	0.229	1.553	1.553	429
0.2%	154	0	0	GET	/api/posting-window/current	9.433	0.061	0.227	0.228	237
0.2%	174	0	0	GET	^/api/relationships/[^/]+	8.506	0.049	0.227	0.249	1766
0.2%	64	0	0	GET	^/api/posts/[^/]+/realmojis$	7.782	0.122	0.407	0.407	166
0.1%	21	0	0	POST	^/api/posts/[^/]+/comments$	6.228	0.297	2.302	2.302	359
0.1%	56	0	0	GET	^/api/posts/[^/]+/comments$	6.202	0.111	0.359	0.359	271
0.1%	102	0	0	GET	^/api/users/[^/]+$	5.407	0.053	0.241	0.242	181
0.1%	107	0	0	POST	/api/auth/logout	5.203	0.049	0.185	0.206	0
0.1%	90	0	0	GET	^/api/memories/[^/]+	4.383	0.049	0.267	0.267	642
0.1%	1	0	0	POST	/api/initialize	4.368	4.368	4.368	4.368	69
0.1%	7	0	0	PUT	^/api/admin/users/[^/]+/certification$	3.292	0.470	3.219	3.219	0
0.1%	15	0	0	PATCH	^/api/posts/[^/]+$	2.360	0.157	1.014	1.014	594
0.0%	1	0	0	POST	/api/me/account-deletion	0.798	0.798	0.798	0.798	0
0.0%	5	0	0	PATCH	/api/me	0.597	0.119	0.190	0.190	237
…（上位 30 / 全 39 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.1%	43	0	0	GET	/api/notifications	440.529	10.245	13.587	13.587	177
12.7%	32	0	0	POST	/api/posts	129.913	4.060	7.528	7.528	606
9.6%	111	0	0	GET	^/api/images/[^/]+	97.927	0.882	1.626	1.740	216352
7.8%	58	0	0	GET	^/api/feeds/[^/]+	80.085	1.381	2.432	2.432	309
5.5%	61	0	0	GET	/api/memories	56.538	0.927	1.932	1.932	11044
5.2%	32	0	0	POST	/api/friend-requests	52.794	1.650	2.883	2.883	115
3.3%	25	0	0	POST	/api/auth/login	33.574	1.343	1.953	1.953	258
2.0%	36	0	0	GET	/api/posting-window/current	20.239	0.562	1.562	1.562	239
1.8%	4	0	0	PUT	^/api/follows/[^/]+$	18.584	4.646	6.318	6.318	0
1.8%	30	12	0	GET	^/api/posts/me/current$	18.112	0.604	0.950	0.950	375
1.7%	5	0	0	POST	/api/auth/register	17.879	3.576	7.408	7.408	228
1.7%	10	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	17.605	1.760	2.433	2.433	116
1.1%	12	0	0	GET	/api/users	10.855	0.905	1.274	1.274	263
0.8%	18	0	0	GET	^/api/relationships/[^/]+	7.764	0.431	0.614	0.614	1342
0.6%	5	2	0	PATCH	^/api/posts/[^/]+$	6.117	1.223	1.451	1.451	394
0.5%	5	0	0	POST	/api/auth/logout	4.672	0.934	1.035	1.035	0
0.3%	2	0	0	GET	/api/me/export	2.676	1.338	1.595	1.595	2792505
0.2%	5	2	0	GET	/api/me	2.027	0.405	0.508	0.508	126
0.2%	1	0	0	PUT	^/api/blocks/[^/]+$	1.822	1.822	1.822	1.822	0
0.2%	3	0	0	GET	^/api/memories/[^/]+	1.788	0.596	0.740	0.740	924
0.1%	1	0	0	POST	/api/me/account-deletion	1.413	1.413	1.413	1.413	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
57.7%	42	6	0	GET	/api/notifications	711.058	16.930	60.002	60.002	156
8.7%	29	0	0	POST	/api/posts	106.752	3.681	5.775	5.775	587
7.3%	100	3	0	GET	^/api/images/[^/]+	90.391	0.904	1.705	1.738	198157
5.1%	35	0	0	POST	/api/friend-requests	62.401	1.783	3.608	3.608	115
4.0%	54	0	0	GET	/api/memories	49.721	0.921	1.810	1.810	5286
3.6%	30	0	0	GET	^/api/feeds/[^/]+	44.153	1.472	2.348	2.348	547
2.6%	25	0	0	POST	/api/auth/login	32.405	1.296	1.956	1.956	247
2.0%	5	0	0	POST	/api/auth/register	24.274	4.855	6.700	6.700	230
1.8%	9	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	21.587	2.399	4.655	4.655	116
1.6%	4	0	0	PUT	^/api/follows/[^/]+$	19.316	4.829	6.223	6.223	0
1.3%	31	0	0	GET	/api/posting-window/current	15.944	0.514	1.376	1.376	240
1.2%	25	9	0	GET	^/api/posts/me/current$	14.412	0.576	0.931	0.931	390
0.9%	15	0	0	GET	/api/users	11.103	0.740	1.048	1.048	223
0.5%	3	0	0	GET	/api/me/export	6.691	2.230	3.013	3.013	4580411
0.5%	16	0	0	GET	^/api/relationships/[^/]+	6.541	0.409	0.609	0.609	1570
0.3%	3	0	0	DELETE	^/api/posts/[^/]+$	3.950	1.317	1.479	1.479	0
0.3%	5	0	0	GET	^/api/memories/[^/]+	3.779	0.756	0.946	0.946	673
0.2%	2	1	0	PATCH	^/api/posts/[^/]+$	2.705	1.353	1.645	1.645	294
0.2%	3	0	0	POST	/api/auth/logout	2.597	0.866	1.021	1.021	0
0.1%	1	0	0	PUT	^/api/blocks/[^/]+$	1.721	1.721	1.721	1.721	0
0.1%	3	0	0	GET	/api/me	1.115	0.372	0.406	0.406	193
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 54621 クエリ / 212 種類 / 合計 617.382s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	24.9%	195	153.684	788.12	5361.42	3.1	1.0	SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE
2	14.2%	1936	87.876	45.39	326.93	357.5	1.0	SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXI…
3	14.1%	1516	86.953	57.36	5413.12	4.2	1.0	SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE
4	12.8%	75	78.828	1051.05	5858.33	6.2	0.0	LOCK TABLE users IN SHARE ROW EXCLUSIVE MODE
5	9.1%	1193	56.214	47.12	110.37	254.3	19.3	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
6	5.2%	51	32.250	632.36	2285.76	13002.1	135.5	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
7	3.4%	191	20.683	108.29	301.47	968.7	1.0	INSERT INTO image_objects (id, owner_user_id, data, media_type, created_at) VALUES ($1, $2, $3, $5, $4)
8	2.4%	18	15.119	839.92	4006.91	32123.8	1640.7	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT gen_random_uuid()::text, id, $2, $3 FROM users WHERE region = $1 AND status = $4
9	2.4%	18	14.874	826.32	2910.91	13.2	0.8	UPDATE posting_windows SET ended_at = $2 WHERE region = $1 AND ended_at IS NULL
10	2.2%	98	13.322	135.94	330.38	4600.6	4921.4	SELECT id, username, secret_key_hash, display_name, bio, profile_image_id, region, account_type, status FROM users WHERE status=$4 AND id<>$1 AND ($2=$5 OR (use…
11	2.1%	2159	12.898	5.97	46.31	87.4	1.0	SELECT u.secret_key_hash FROM image_objects i JOIN users u ON u.id = i.owner_user_id WHERE i.id = $1 AND u.status = $2
12	1.1%	389	6.955	17.88	2195.99	4.2	1.0	SELECT id FROM users WHERE id = $1 FOR NO KEY UPDATE
13	0.8%	679	5.136	7.56	75.27	100.6	6.1	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
14	0.4%	4308	2.467	0.57	40.34	17.1	1.0	SELECT s.user_id, s.token_hash, s.created_at, u.region FROM login_sessions s JOIN users u ON u.id = s.user_id WHERE s.token_hash = $1 AND s.active = $2 AND u.st…
15	0.3%	31	1.758	56.72	104.63	1023.0	2000.8	SELECT user_low_id, user_high_id FROM friendships WHERE user_low_id IN ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $2…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE;
-- rank 2
SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXISTS ( SELECT $3 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL ) OR EXISTS ( SELECT $4 FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.autho…;
-- rank 3
SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: a23ba3197135f2bdb5294626e505fd4c5761d0d3
Type: cpu
Time: 2026-10-31 10:39:29 JST
Duration: 60s, Total samples = 23.44s (39.07%)
Showing nodes accounting for 19.11s, 81.53% of 23.44s total
Dropped 814 nodes (cum <= 0.12s)
Showing top 40 nodes out of 274
      flat  flat%   sum%        cum   cum%
     8.49s 36.22% 36.22%      8.52s 36.35%  crypto/internal/fips140/sha256.blockAVX2
     1.31s  5.59% 41.81%      1.31s  5.59%  internal/runtime/syscall/linux.Syscall6
     1.22s  5.20% 47.01%      1.22s  5.20%  runtime.memclrNoHeapPointers
     1.18s  5.03% 52.05%      1.18s  5.03%  runtime.memmove
     0.99s  4.22% 56.27%      1.43s  6.10%  image/jpeg.(*decoder).reconstructBlock
     0.53s  2.26% 58.53%      9.65s 41.17%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.51s  2.18% 60.71%      2.98s 12.71%  image/jpeg.(*decoder).processSOS
     0.49s  2.09% 62.80%      9.44s 40.27%  crypto/internal/fips140/sha256.(*Digest).Write
     0.36s  1.54% 64.33%      0.65s  2.77%  image/jpeg.(*decoder).decodeHuffman
     0.32s  1.37% 65.70%     10.83s 46.20%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.30s  1.28% 66.98%      0.31s  1.32%  image/jpeg.(*decoder).receiveExtend
     0.27s  1.15% 68.13%     10.11s 43.13%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.27s  1.15% 69.28%      0.30s  1.28%  image/jpeg.(*decoder).ensureNBits
     0.25s  1.07% 70.35%      0.25s  1.07%  runtime.futex
     0.24s  1.02% 71.37%      4.93s 21.03%  crypto/internal/fips140/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
     0.22s  0.94% 72.31%      0.22s  0.94%  runtime.madvise
     0.21s   0.9% 73.21%      0.26s  1.11%  image/jpeg.idctCols
     0.18s  0.77% 73.98%      0.48s  2.05%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.18s  0.77% 74.74%      0.24s  1.02%  runtime.tryDeferToSpanScan
     0.17s  0.73% 75.47%      0.17s  0.73%  runtime.nextFreeFast (inline)
     0.13s  0.55% 76.02%      6.83s 29.14%  example.com/isureal-company/isureal.imageSignature
     0.13s  0.55% 76.58%      0.18s  0.77%  image/jpeg.idctRows
     0.13s  0.55% 77.13%      0.13s  0.55%  runtime.procyieldAsm
     0.13s  0.55% 77.69%      0.28s  1.19%  runtime.scanObject
     0.12s  0.51% 78.20%      0.39s  1.66%  runtime.mallocgcSmallScanNoHeader
     0.11s  0.47% 78.67%      0.22s  0.94%  runtime.scanObjectsSmall
     0.09s  0.38% 79.05%      1.11s  4.74%  runtime.mallocgc
     0.08s  0.34% 79.39%      0.13s  0.55%  crypto/internal/fips140.RecordApproved
     0.07s   0.3% 79.69%      0.32s  1.37%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.07s   0.3% 79.99%      8.59s 36.65%  crypto/internal/fips140/sha256.block
     0.07s   0.3% 80.29%      0.28s  1.19%  runtime.newobject
     0.07s   0.3% 80.59%      0.17s  0.73%  runtime.pcvalue
     0.04s  0.17% 80.76%      0.20s  0.85%  crypto/internal/fips140/hmac.(*HMAC).Write (partial-inline)
     0.04s  0.17% 80.93%      0.72s  3.07%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.03s  0.13% 81.06%      0.49s  2.09%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).receiveMessage
     0.03s  0.13% 81.19%      0.72s  3.07%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
     0.03s  0.13% 81.31%      0.14s   0.6%  golang.org/x/text/unicode/norm.(*formInfo).quickSpan
     0.02s 0.085% 81.40%      0.19s  0.81%  github.com/jackc/pgx/v5/pgtype.scanPlanString.Scan
     0.02s 0.085% 81.48%      1.17s  4.99%  github.com/jmoiron/sqlx.scanAll
     0.01s 0.043% 81.53%      0.12s  0.51%  crypto/ecdh.x25519ScalarMult
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 21.4% / 89% | 0% | 0.1% | 6 | 1923 MB | run queue が積もっている |
| isu2 (2 cpu) | 1.8% / 15% | 0% | 0.0% | 4 | 3140 MB |  |
| isu3 (2 cpu) | 2.2% / 21% | 0% | 0.0% | 3 | 3201 MB |  |
| isu4 (2 cpu) | 44.9% / 100% | 37% | 3.4% | 34 | 90 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 34% / nginx 5% |
| isu2 | isureal 3% |
| isu3 | isureal 3% |
| isu4 | postgres 20% |

## Diff from previous run

スコア: 7675 → 8141 (+466)

`02acb51..ddc4d70`

```
 conf/isu4/etc/postgresql/16/main/conf.d/zz-tuning.conf |  7 +++++++
 webapp/sql/002-indexes.sql                             | 10 ++++++++++
 webapp/sql/reset.sql                                   |  2 ++
 3 files changed, 19 insertions(+)
```
