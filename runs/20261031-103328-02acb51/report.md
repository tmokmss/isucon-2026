# run 20261031-103328-02acb51

## Summary

| | |
|---|---|
| score | 7675 |
| status | pass |
| ref | HEAD |
| sha | 02acb51f3856 |
| 所要時間 | 202s |
| 前回比 | +1286 (20.1%) vs 20261031-102655-8cd1279 |
| ベスト比 | +1286 (20.1%) vs 20261031-102655-8cd1279 = 6389 **ベスト更新** |
| note | 画像の署名: hmac の使い回し・キャッシュ・exp を 1 分単位に |

## Bench

```
{"id":"1019","status":"done","portalStatus":"finished","passed":true,"score":7675,"raw":7675,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:34:55.316Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:34:57.396Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:35:00.201Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:35:04.538Z level=INFO msg="asia 2 日目: 19 人"
time=2026-10-31T01:35:10.513Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T01:35:10.559Z level=INFO msg="america 2 日目: 6 人"
time=2026-10-31T01:35:13.707Z level=INFO msg="asia 3 日目: 21 人"
time=2026-10-31T01:35:16.328Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T01:35:20.554Z level=INFO msg="america 3 日目: 6 人"
time=2026-10-31T01:35:23.630Z level=INFO msg="asia 4 日目: 23 人"
time=2026-10-31T01:35:25.232Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T01:35:27.224Z level=INFO msg="america 4 日目: 6 人"
time=2026-10-31T01:35:32.734Z level=INFO msg="asia 5 日目: 24 人"
time=2026-10-31T01:35:34.858Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T01:35:36.505Z level=INFO msg="america 5 日目: 6 人"
time=2026-10-31T01:35:45.080Z level=INFO msg="europe 最終日（グランプリ）: 6 人"
time=2026-10-31T01:35:45.832Z level=INFO msg="asia 最終日（グランプリ）: 23 人"
time=2026-10-31T01:35:45.869Z level=INFO msg="america 最終日（グランプリ）: 6 人"
time=2026-10-31T01:35:54.029Z level=INFO msg="グランプリ: アジア 23 人, ヨーロッパ 3 人, アメリカ 3 人"
{
  "pass": true,
  "score": 7675,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
74.2%	146	7	0	GET	/api/notifications	2548.903	17.458	60.001	60.001	148
7.9%	1910	192	0	GET	^/api/images/[^/]+	270.403	0.142	0.406	0.576	87792
3.9%	115	0	0	POST	/api/posts	134.490	1.169	2.790	2.949	600
2.9%	907	0	0	GET	^/api/feeds/[^/]+	101.175	0.112	0.319	0.576	1999
1.6%	144	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	55.165	0.383	3.026	4.075	116
1.5%	453	0	0	GET	/api/memories	53.170	0.117	0.419	0.651	14707
1.3%	150	14	0	POST	/api/friend-requests	44.817	0.299	2.916	3.674	107
1.3%	1036	514	0	POST	/api/auth/login	44.685	0.043	0.213	0.349	140
1.0%	45	0	0	GET	/api/me/export	33.258	0.739	2.806	2.806	2050620
0.7%	18	0	0	POST	/api/admin/posting-window	25.512	1.417	4.481	4.481	177
0.5%	40	1	0	POST	/api/auth/register	15.628	0.391	3.989	3.989	235
0.4%	211	2	0	GET	^/api/posts/[^/]+$	12.925	0.061	0.255	0.322	584
0.3%	125	0	0	GET	/api/posting-window/current	11.997	0.096	0.349	0.350	237
0.3%	107	99	5	PUT	/api/me/profile-image	11.594	0.108	0.703	0.941	29
0.3%	29	0	0	POST	^/api/posts/[^/]+/realmojis$	10.258	0.354	3.008	3.008	409
0.3%	58	0	0	GET	/api/users	9.663	0.167	0.385	0.385	219
0.2%	116	1	0	GET	^/api/posts/me/current$	7.785	0.067	0.199	0.249	595
0.2%	97	0	0	GET	^/api/posts/[^/]+/realmojis$	6.615	0.068	0.212	0.212	159
0.2%	170	0	0	GET	^/api/relationships/[^/]+	6.046	0.036	0.185	0.191	1733
0.1%	108	0	0	POST	/api/auth/logout	4.450	0.041	0.240	0.271	0
0.1%	1	0	0	POST	/api/initialize	4.133	4.133	4.133	4.133	69
0.1%	53	0	0	GET	^/api/posts/[^/]+/comments$	4.004	0.076	0.302	0.302	29
0.1%	102	0	0	GET	^/api/users/[^/]+$	3.865	0.038	0.150	0.159	164
0.1%	96	0	0	GET	^/api/memories/[^/]+	3.582	0.037	0.245	0.245	793
0.1%	19	0	0	PUT	^/api/follows/[^/]+$	2.721	0.143	1.039	1.039	0
0.1%	3	0	0	DELETE	^/api/follows/[^/]+$	2.655	0.885	2.611	2.611	0
0.1%	24	0	0	PATCH	^/api/posts/[^/]+$	2.522	0.105	1.051	1.051	597
0.0%	3	0	0	DELETE	^/api/posts/[^/]+$	0.498	0.166	0.419	0.419	0
0.0%	7	0	0	DELETE	^/api/realmojis/[^/]+$	0.249	0.036	0.052	0.052	0
0.0%	4	0	0	PATCH	/api/me	0.182	0.045	0.132	0.132	215
…（上位 30 / 全 37 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
56.4%	48	6	0	GET	/api/notifications	752.048	15.668	60.002	60.002	155
10.1%	31	0	0	POST	/api/posts	134.393	4.335	10.199	10.199	603
7.2%	116	1	0	GET	^/api/images/[^/]+	95.991	0.828	1.540	1.573	210052
5.5%	40	6	0	POST	/api/friend-requests	72.842	1.821	3.543	3.543	103
4.5%	42	0	0	GET	^/api/feeds/[^/]+	60.391	1.438	2.432	2.432	634
4.1%	58	0	0	GET	/api/memories	54.595	0.941	1.785	1.785	7010
2.5%	25	0	0	POST	/api/auth/login	33.370	1.335	1.963	1.963	245
1.6%	33	0	0	GET	/api/posting-window/current	21.072	0.639	1.576	1.576	239
1.4%	22	0	0	GET	/api/users	18.793	0.854	1.255	1.255	239
1.4%	5	0	0	POST	/api/auth/register	18.008	3.602	7.436	7.436	227
1.3%	4	0	0	PUT	^/api/follows/[^/]+$	17.095	4.274	5.718	5.718	0
1.2%	29	8	0	GET	^/api/posts/me/current$	16.660	0.574	1.027	1.027	444
0.7%	4	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	9.526	2.381	4.950	4.950	116
0.7%	20	0	0	GET	^/api/relationships/[^/]+	8.904	0.445	0.618	0.618	1189
0.4%	3	0	0	GET	/api/me/export	5.294	1.765	2.658	2.658	3014552
0.3%	5	0	0	POST	/api/auth/logout	4.356	0.871	1.034	1.034	0
0.1%	5	2	0	GET	/api/me	1.929	0.386	0.511	0.511	125
0.1%	2	0	0	GET	^/api/memories/[^/]+	1.690	0.845	1.058	1.058	1080
0.1%	1	0	0	PUT	^/api/blocks/[^/]+$	1.618	1.618	1.618	1.618	0
0.1%	1	0	0	DELETE	^/api/posts/[^/]+$	1.535	1.535	1.535	1.535	0
0.1%	1	1	0	PATCH	^/api/posts/[^/]+$	1.124	1.124	1.124	1.124	30
0.1%	1	0	0	POST	/api/me/account-deletion	1.119	1.119	1.119	1.119	0
0.1%	1	1	0	GET	^/api/posts/[^/]+$	0.935	0.935	0.935	0.935	27
0.0%	1	0	0	GET	^/api/users/[^/]+$	0.614	0.614	0.614	0.614	167

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
58.4%	41	5	0	GET	/api/notifications	642.981	15.682	60.002	60.002	155
7.9%	101	1	0	GET	^/api/images/[^/]+	87.368	0.865	1.544	1.786	162984
6.5%	24	0	0	POST	/api/posts	71.834	2.993	5.311	5.311	594
5.1%	32	2	0	POST	/api/friend-requests	56.422	1.763	3.616	3.616	110
4.6%	52	0	0	GET	/api/memories	50.139	0.964	1.917	1.917	7098
4.2%	32	0	0	GET	^/api/feeds/[^/]+	46.136	1.442	2.582	2.582	944
2.9%	24	0	0	POST	/api/auth/login	32.113	1.338	2.028	2.028	245
2.2%	5	0	0	POST	/api/auth/register	24.396	4.879	6.724	6.724	227
1.8%	4	0	0	PUT	^/api/follows/[^/]+$	19.764	4.941	6.842	6.842	0
1.4%	19	0	0	GET	/api/users	15.435	0.812	1.939	1.939	243
1.4%	26	0	0	GET	/api/posting-window/current	14.959	0.575	1.337	1.337	240
1.3%	22	6	0	GET	^/api/posts/me/current$	13.864	0.630	1.701	1.701	443
0.6%	3	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	6.169	2.056	2.732	2.732	116
0.5%	12	0	0	GET	^/api/relationships/[^/]+	5.787	0.482	0.609	0.609	1726
0.3%	4	0	0	POST	/api/auth/logout	3.335	0.834	0.848	0.848	0
0.2%	4	0	0	GET	^/api/memories/[^/]+	2.656	0.664	1.070	1.070	966
0.2%	1	0	0	PUT	^/api/blocks/[^/]+$	1.821	1.821	1.821	1.821	0
0.1%	1	0	0	POST	/api/me/account-deletion	1.378	1.378	1.378	1.378	0
0.1%	1	0	0	DELETE	^/api/posts/[^/]+$	1.235	1.235	1.235	1.235	0
0.1%	4	1	0	GET	/api/me	1.216	0.304	0.405	0.405	150
0.1%	1	0	0	GET	/api/me/export	1.052	1.052	1.052	1.052	588591
0.1%	1	1	0	PATCH	^/api/posts/[^/]+$	0.738	0.738	0.738	0.738	30
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 48136 クエリ / 203 種類 / 合計 544.322s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	28.0%	177	152.636	862.35	6171.55	3.1	1.0	SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE
2	23.7%	1452	129.058	88.88	4493.15	4.3	1.0	SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE
3	11.8%	1938	64.248	33.15	271.30	355.0	1.0	SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXI…
4	7.7%	68	42.098	619.09	5773.03	6.8	0.0	LOCK TABLE users IN SHARE ROW EXCLUSIVE MODE
5	5.9%	981	31.919	32.54	97.98	250.5	17.6	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
6	5.4%	49	29.408	600.17	2396.53	12830.9	133.6	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
7	2.8%	170	15.447	90.86	311.44	978.0	1.0	INSERT INTO image_objects (id, owner_user_id, data, media_type, created_at) VALUES ($1, $2, $3, $5, $4)
8	2.5%	99	13.481	136.17	320.55	4598.9	4920.3	SELECT id, username, secret_key_hash, display_name, bio, profile_image_id, region, account_type, status FROM users WHERE status=$4 AND id<>$1 AND ($2=$5 OR (use…
9	2.0%	18	10.672	592.88	4318.09	13.1	0.8	UPDATE posting_windows SET ended_at = $2 WHERE region = $1 AND ended_at IS NULL
10	2.0%	18	10.627	590.36	2668.82	32734.2	1639.9	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT gen_random_uuid()::text, id, $2, $3 FROM users WHERE region = $1 AND status = $4
11	1.8%	2115	9.761	4.61	35.49	87.3	1.0	SELECT u.secret_key_hash FROM image_objects i JOIN users u ON u.id = i.owner_user_id WHERE i.id = $1 AND u.status = $2
12	1.2%	325	6.674	20.53	1084.17	4.3	1.0	SELECT id FROM users WHERE id = $1 FOR NO KEY UPDATE
13	0.7%	665	3.721	5.60	45.72	116.7	6.5	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
14	0.3%	52	1.470	28.28	64.22	984.0	1924.0	SELECT user_low_id, user_high_id FROM friendships WHERE user_low_id IN ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $2…
15	0.3%	3893	1.468	0.38	25.75	12.1	1.0	SELECT s.user_id, s.token_hash, s.created_at, u.region FROM login_sessions s JOIN users u ON u.id = s.user_id WHERE s.token_hash = $1 AND s.active = $2 AND u.st…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE;
-- rank 2
SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE;
-- rank 3
SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXISTS ( SELECT $3 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL ) OR EXISTS ( SELECT $4 FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.autho…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 2ac609560257914e5194e35b427eb626df732178
Type: cpu
Time: 2026-10-31 10:34:24 JST
Duration: 60.05s, Total samples = 24.65s (41.05%)
Showing nodes accounting for 20.19s, 81.91% of 24.65s total
Dropped 845 nodes (cum <= 0.12s)
Showing top 40 nodes out of 259
      flat  flat%   sum%        cum   cum%
     9.16s 37.16% 37.16%      9.18s 37.24%  crypto/internal/fips140/sha256.blockAVX2
     1.34s  5.44% 42.60%      2.04s  8.28%  image/jpeg.(*decoder).reconstructBlock
     1.25s  5.07% 47.67%      1.25s  5.07%  runtime.memclrNoHeapPointers
     1.22s  4.95% 52.62%      1.22s  4.95%  internal/runtime/syscall/linux.Syscall6
     1.08s  4.38% 57.00%      1.08s  4.38%  runtime.memmove
     0.52s  2.11% 59.11%     10.30s 41.78%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.49s  1.99% 61.10%     10.08s 40.89%  crypto/internal/fips140/sha256.(*Digest).Write
     0.42s  1.70% 62.80%      3.54s 14.36%  image/jpeg.(*decoder).processSOS
     0.39s  1.58% 64.38%      0.41s  1.66%  image/jpeg.(*decoder).receiveExtend
     0.36s  1.46% 65.84%      0.65s  2.64%  image/jpeg.(*decoder).decodeHuffman
     0.33s  1.34% 67.18%      5.83s 23.65%  crypto/internal/fips140/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
     0.31s  1.26% 68.44%     10.77s 43.69%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.28s  1.14% 69.57%      0.36s  1.46%  image/jpeg.idctCols
     0.26s  1.05% 70.63%      0.34s  1.38%  image/jpeg.idctRows
     0.24s  0.97% 71.60%     11.46s 46.49%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.24s  0.97% 72.58%      0.31s  1.26%  image/jpeg.(*decoder).ensureNBits
     0.22s  0.89% 73.47%      0.22s  0.89%  runtime.procyieldAsm
     0.21s  0.85% 74.32%      0.50s  2.03%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.19s  0.77% 75.09%      0.37s  1.50%  runtime.scanObject
     0.16s  0.65% 75.74%      0.16s  0.65%  image/jpeg.dctBox (inline)
     0.16s  0.65% 76.39%      0.16s  0.65%  runtime.futex
     0.16s  0.65% 77.04%      0.16s  0.65%  runtime.madvise
     0.15s  0.61% 77.65%      9.33s 37.85%  crypto/internal/fips140/sha256.block
     0.14s  0.57% 78.22%      0.17s  0.69%  runtime.tryDeferToSpanScan
     0.13s  0.53% 78.74%      0.36s  1.46%  runtime.pcvalue
     0.13s  0.53% 79.27%      0.18s  0.73%  runtime.step
     0.09s  0.37% 79.63%      0.33s  1.34%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.08s  0.32% 79.96%      0.33s  1.34%  runtime.mallocgcSmallScanNoHeader
     0.07s  0.28% 80.24%      0.15s  0.61%  runtime.scanObjectsSmall
     0.06s  0.24% 80.49%      1.10s  4.46%  runtime.mallocgc
     0.05s   0.2% 80.69%      0.19s  0.77%  crypto/internal/fips140/hmac.(*HMAC).Write (partial-inline)
     0.05s   0.2% 80.89%      0.31s  1.26%  runtime.newobject
     0.04s  0.16% 81.05%      6.55s 26.57%  example.com/isureal-company/isureal.imageSignature
     0.04s  0.16% 81.22%      0.24s  0.97%  runtime.(*unwinder).resolveInternal
     0.03s  0.12% 81.34%      0.64s  2.60%  crypto/tls.(*Conn).readRecordOrCCS
     0.03s  0.12% 81.46%      0.14s  0.57%  database/sql.convertAssignRows
     0.03s  0.12% 81.58%      0.18s  0.73%  golang.org/x/text/cases.(*caseFolder).Transform
     0.03s  0.12% 81.70%      0.40s  1.62%  runtime.findRunnable
     0.03s  0.12% 81.83%      0.22s  0.89%  runtime.scanSpan
     0.02s 0.081% 81.91%      0.17s  0.69%  crypto/tls.(*halfConn).decrypt
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 20.7% / 96% | 2% | 0.0% | 8 | 2019 MB | run queue が積もっている |
| isu2 (2 cpu) | 1.7% / 10% | 0% | 0.0% | 3 | 3139 MB |  |
| isu3 (2 cpu) | 2.0% / 8% | 0% | 0.0% | 2 | 3224 MB |  |
| isu4 (2 cpu) | 45.3% / 100% | 35% | 3.6% | 28 | 109 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 34% / nginx 4% |
| isu2 | isureal 3% |
| isu3 | isureal 3% |
| isu4 | postgres 17% |

## Diff from previous run

スコア: 6389 → 7675 (+1286)

`8cd1279..02acb51`

```
 webapp/go/image_signature.go | 34 +++++++++++++++++++++++++++++++---
 1 file changed, 31 insertions(+), 3 deletions(-)
```
