# run 20261031-103549-92ad1cc

## Summary

| | |
|---|---|
| score | 7316 |
| status | pass |
| ref | db-drop-table-locks |
| sha | 92ad1ccf072e |
| 所要時間 | 216s |
| 前回比 | +927 (14.5%) vs 20261031-102655-8cd1279 |
| ベスト比 | +927 (14.5%) vs 20261031-102655-8cd1279 = 6389 **ベスト更新** |
| note | 投稿期間の開始のLOCK TABLEをadvisory lockに（DBはインデックスと設定入り、署名修正入りのmainがbase） |

## Bench

```
{"id":"1139","status":"done","portalStatus":"finished","passed":true,"score":7316,"raw":7316,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:37:27.846Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:37:28.661Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:37:31.854Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:37:37.218Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T01:37:39.612Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T01:37:40.779Z level=INFO msg="america 2 日目: 6 人"
time=2026-10-31T01:37:44.808Z level=INFO msg="asia 3 日目: 18 人"
time=2026-10-31T01:37:47.897Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T01:37:51.625Z level=INFO msg="america 3 日目: 6 人"
time=2026-10-31T01:37:53.818Z level=INFO msg="asia 4 日目: 19 人"
time=2026-10-31T01:37:55.714Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T01:37:58.763Z level=INFO msg="america 4 日目: 6 人"
time=2026-10-31T01:38:03.218Z level=INFO msg="asia 5 日目: 23 人"
time=2026-10-31T01:38:05.963Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T01:38:07.754Z level=INFO msg="america 5 日目: 6 人"
time=2026-10-31T01:38:16.696Z level=INFO msg="europe 最終日（グランプリ）: 7 人"
time=2026-10-31T01:38:17.203Z level=INFO msg="america 最終日（グランプリ）: 6 人"
time=2026-10-31T01:38:17.777Z level=INFO msg="asia 最終日（グランプリ）: 22 人"
time=2026-10-31T01:38:25.629Z level=INFO msg="グランプリ: アジア 21 人, ヨーロッパ 4 人, アメリカ 2 人"
{
  "pass": true,
  "score": 7316,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
76.9%	136	6	0	GET	/api/notifications	2414.772	17.756	60.001	60.002	149
8.8%	2217	235	0	GET	^/api/images/[^/]+	277.554	0.125	0.374	0.510	68690
2.8%	819	0	0	GET	^/api/feeds/[^/]+	88.846	0.108	0.291	0.355	1053
2.6%	96	0	0	POST	/api/posts	81.076	0.845	2.457	2.457	611
1.7%	507	0	0	GET	/api/memories	52.271	0.103	0.378	0.463	15120
1.2%	1139	512	0	POST	/api/auth/login	38.631	0.034	0.166	0.287	150
0.9%	86	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	27.547	0.320	4.364	4.364	116
0.9%	42	0	0	GET	/api/me/export	27.524	0.655	2.469	2.469	1995492
0.9%	98	5	0	POST	/api/friend-requests	27.178	0.277	3.919	3.919	111
0.6%	18	0	0	POST	/api/admin/posting-window	17.852	0.992	2.588	2.588	177
0.4%	166	0	0	GET	^/api/posts/[^/]+$	11.918	0.072	0.253	0.285	593
0.3%	30	1	0	POST	^/api/posts/[^/]+/realmojis$	9.294	0.310	1.691	1.691	406
0.3%	103	91	7	PUT	/api/me/profile-image	7.889	0.077	0.314	0.325	31
0.2%	108	0	0	GET	/api/posting-window/current	6.651	0.062	0.200	0.272	237
0.2%	39	0	0	GET	/api/users	6.089	0.156	0.335	0.335	213
0.2%	74	0	0	GET	^/api/posts/[^/]+/realmojis$	5.518	0.075	0.240	0.240	147
0.2%	97	1	0	GET	^/api/posts/me/current$	5.212	0.054	0.231	0.231	605
0.2%	140	0	0	POST	/api/auth/logout	5.033	0.036	0.222	0.736	0
0.1%	1	0	0	POST	/api/initialize	4.613	4.613	4.613	4.613	69
0.1%	133	0	0	GET	^/api/relationships/[^/]+	4.132	0.031	0.159	0.194	1803
0.1%	55	0	0	GET	^/api/posts/[^/]+/comments$	3.703	0.067	0.250	0.250	64
0.1%	15	0	0	PUT	^/api/follows/[^/]+$	3.141	0.209	1.519	1.519	0
0.1%	113	0	0	GET	^/api/memories/[^/]+	2.783	0.025	0.174	0.193	460
0.1%	39	1	0	POST	/api/auth/register	2.591	0.066	0.207	0.207	236
0.1%	51	0	0	GET	^/api/users/[^/]+$	1.792	0.035	0.180	0.180	164
0.0%	6	0	0	DELETE	^/api/realmojis/[^/]+$	1.367	0.228	1.227	1.227	0
0.0%	1	0	0	POST	/api/me/account-deletion	1.163	1.163	1.163	1.163	0
0.0%	5	0	0	POST	^/api/posts/[^/]+/comments$	0.770	0.154	0.272	0.272	268
0.0%	11	0	0	PATCH	^/api/posts/[^/]+$	0.653	0.059	0.201	0.201	615
0.0%	4	0	0	DELETE	^/api/posts/[^/]+$	0.395	0.099	0.127	0.127	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
58.7%	49	5	0	GET	/api/notifications	757.887	15.467	60.002	60.002	155
8.3%	29	0	0	POST	/api/posts	106.894	3.686	7.592	7.592	611
6.5%	107	0	0	GET	^/api/images/[^/]+	83.860	0.784	1.743	2.054	203209
5.7%	46	1	0	POST	/api/friend-requests	73.545	1.599	3.293	3.293	113
4.2%	40	0	0	GET	^/api/feeds/[^/]+	54.742	1.369	1.719	1.719	668
4.2%	57	0	0	GET	/api/memories	53.905	0.946	1.828	1.828	5897
2.7%	27	0	0	POST	/api/auth/login	35.456	1.313	1.951	1.951	253
1.5%	4	0	0	PUT	^/api/follows/[^/]+$	19.444	4.861	6.733	6.733	0
1.3%	5	0	0	POST	/api/auth/register	17.200	3.440	6.733	6.733	227
1.3%	20	0	0	GET	/api/users	16.877	0.844	1.253	1.253	237
1.3%	29	0	0	GET	/api/posting-window/current	16.717	0.576	1.507	1.507	239
1.3%	27	7	0	GET	^/api/posts/me/current$	16.322	0.605	0.958	0.958	466
1.1%	8	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	14.412	1.801	3.681	3.681	116
0.7%	22	0	0	GET	^/api/relationships/[^/]+	8.402	0.382	0.623	0.623	1180
0.4%	4	0	0	GET	/api/me/export	4.526	1.131	1.694	1.694	1809969
0.2%	4	0	0	GET	^/api/memories/[^/]+	2.773	0.693	1.028	1.028	866
0.2%	3	0	0	POST	/api/auth/logout	2.477	0.826	0.830	0.830	0
0.1%	1	0	0	PUT	^/api/blocks/[^/]+$	1.922	1.922	1.922	1.922	0
0.1%	5	2	0	GET	/api/me	1.620	0.324	0.508	0.508	126
0.1%	1	0	0	POST	/api/me/account-deletion	1.035	1.035	1.035	1.035	0
0.1%	1	1	0	PATCH	^/api/posts/[^/]+$	0.737	0.737	0.737	0.737	30

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
58.0%	42	5	0	GET	/api/notifications	645.649	15.373	60.003	60.003	156
7.8%	111	1	0	GET	^/api/images/[^/]+	87.431	0.788	1.511	1.563	170735
6.7%	23	0	0	POST	/api/posts	75.032	3.262	6.752	6.752	582
6.4%	42	1	0	POST	/api/friend-requests	71.834	1.710	2.786	2.786	113
4.2%	49	0	0	GET	/api/memories	46.560	0.950	1.828	1.828	11371
3.3%	27	0	0	GET	^/api/feeds/[^/]+	36.880	1.366	1.606	1.606	619
3.0%	27	0	0	POST	/api/auth/login	33.873	1.255	1.947	1.947	260
2.2%	5	0	0	POST	/api/auth/register	25.055	5.011	7.443	7.443	231
1.7%	8	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	18.617	2.327	4.454	4.454	106
1.5%	4	0	0	PUT	^/api/follows/[^/]+$	16.220	4.055	5.422	5.422	0
1.4%	26	0	0	GET	/api/posting-window/current	15.143	0.582	1.475	1.475	240
1.2%	18	0	0	GET	/api/users	13.473	0.748	1.878	1.878	244
1.0%	20	9	0	GET	^/api/posts/me/current$	10.862	0.543	0.956	0.956	329
0.6%	15	0	0	GET	^/api/relationships/[^/]+	6.431	0.429	0.610	0.610	1678
0.3%	6	0	0	GET	^/api/memories/[^/]+	3.235	0.539	0.838	0.838	901
0.2%	1	0	0	PUT	^/api/blocks/[^/]+$	1.720	1.720	1.720	1.720	0
0.1%	2	0	0	POST	/api/auth/logout	1.556	0.778	0.834	0.834	0
0.1%	1	0	0	GET	/api/me/export	1.198	1.198	1.198	1.198	2807577
0.1%	1	0	0	DELETE	^/api/posts/[^/]+$	1.139	1.139	1.139	1.139	0
0.1%	4	1	0	GET	/api/me	1.118	0.279	0.406	0.406	152
0.1%	1	0	0	POST	/api/me/account-deletion	0.818	0.818	0.818	0.818	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 46571 クエリ / 199 種類 / 合計 411.892s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	23.9%	153	98.367	642.92	5114.76	3.1	1.0	SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE
2	21.3%	1390	87.634	63.05	4391.39	4.3	1.0	SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE
3	15.1%	2204	62.247	28.24	277.86	352.8	1.0	SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXI…
4	7.1%	886	29.165	32.92	104.61	250.2	14.8	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
5	6.5%	49	26.624	543.35	5493.28	6.4	0.0	LOCK TABLE users IN SHARE ROW EXCLUSIVE MODE
6	5.8%	47	23.723	504.74	2152.12	12309.7	128.3	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
7	3.0%	148	12.229	82.63	251.05	989.3	1.0	INSERT INTO image_objects (id, owner_user_id, data, media_type, created_at) VALUES ($1, $2, $3, $5, $4)
8	2.7%	77	11.036	143.32	318.75	4598.8	4919.0	SELECT id, username, secret_key_hash, display_name, bio, profile_image_id, region, account_type, status FROM users WHERE status=$4 AND id<>$1 AND ($2=$5 OR (use…
9	2.4%	2423	10.010	4.13	38.09	87.3	1.0	SELECT u.secret_key_hash FROM image_objects i JOIN users u ON u.id = i.owner_user_id WHERE i.id = $1 AND u.status = $2
10	2.4%	18	9.784	543.56	2489.06	32316.3	1639.4	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT gen_random_uuid()::text, id, $2, $3 FROM users WHERE region = $1 AND status = $4
11	1.6%	18	6.681	371.17	1897.24	13.3	0.8	UPDATE posting_windows SET ended_at = $2 WHERE region = $1 AND ended_at IS NULL
12	1.6%	327	6.580	20.12	1704.52	4.3	1.0	SELECT id FROM users WHERE id = $1 FOR NO KEY UPDATE
13	0.9%	736	3.718	5.05	70.81	116.2	6.7	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
14	0.4%	3553	1.761	0.50	38.07	15.5	1.0	SELECT s.user_id, s.token_hash, s.created_at, u.region FROM login_sessions s JOIN users u ON u.id = s.user_id WHERE s.token_hash = $1 AND s.active = $2 AND u.st…
15	0.4%	1	1.544	1543.84	1543.84	19805.0	170.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: 7cfec2d7f9b60c6be12b1c4e66e925b84c511dca
Type: cpu
Time: 2026-10-31 10:36:55 JST
Duration: 60.01s, Total samples = 23.05s (38.41%)
Showing nodes accounting for 19.06s, 82.69% of 23.05s total
Dropped 784 nodes (cum <= 0.12s)
Showing top 40 nodes out of 271
      flat  flat%   sum%        cum   cum%
     9.01s 39.09% 39.09%      9.03s 39.18%  crypto/internal/fips140/sha256.blockAVX2
     1.13s  4.90% 43.99%      1.13s  4.90%  internal/runtime/syscall/linux.Syscall6
     1.09s  4.73% 48.72%      1.09s  4.73%  runtime.memmove
     1.03s  4.47% 53.19%      1.03s  4.47%  runtime.memclrNoHeapPointers
     0.91s  3.95% 57.14%      1.48s  6.42%  image/jpeg.(*decoder).reconstructBlock
     0.53s  2.30% 59.44%      2.94s 12.75%  image/jpeg.(*decoder).processSOS
     0.52s  2.26% 61.69%        10s 43.38%  crypto/internal/fips140/sha256.(*Digest).Write
     0.45s  1.95% 63.64%     10.13s 43.95%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.42s  1.82% 65.47%     10.70s 46.42%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.35s  1.52% 66.98%      0.63s  2.73%  image/jpeg.(*decoder).decodeHuffman
     0.30s  1.30% 68.29%      5.23s 22.69%  crypto/internal/fips140/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
     0.30s  1.30% 69.59%      0.30s  1.30%  runtime.procyieldAsm
     0.26s  1.13% 70.72%     11.33s 49.15%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.25s  1.08% 71.80%      0.28s  1.21%  image/jpeg.(*decoder).receiveExtend
     0.23s     1% 72.80%      0.31s  1.34%  image/jpeg.(*decoder).ensureNBits
     0.22s  0.95% 73.75%      0.28s  1.21%  image/jpeg.idctCols
     0.22s  0.95% 74.71%      0.22s  0.95%  runtime.madvise
     0.21s  0.91% 75.62%      0.28s  1.21%  image/jpeg.idctRows
     0.17s  0.74% 76.36%      0.17s  0.74%  runtime.futex
     0.16s  0.69% 77.05%      9.19s 39.87%  crypto/internal/fips140/sha256.block
     0.15s  0.65% 77.70%      0.62s  2.69%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.13s  0.56% 78.26%      0.13s  0.56%  image/jpeg.dctBox (inline)
     0.12s  0.52% 78.79%      0.12s  0.52%  internal/byteorder.BEUint32 (inline)
     0.11s  0.48% 79.26%      0.20s  0.87%  runtime.tryDeferToSpanScan
     0.10s  0.43% 79.70%      0.12s  0.52%  runtime.step
     0.08s  0.35% 80.04%      7.18s 31.15%  example.com/isureal-company/isureal.imageSignature
     0.08s  0.35% 80.39%      0.21s  0.91%  runtime.scanObjectsSmall
     0.06s  0.26% 80.65%      0.41s  1.78%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.05s  0.22% 80.87%      0.81s  3.51%  runtime.mallocgc
     0.05s  0.22% 81.08%      0.22s  0.95%  runtime.mallocgcSmallScanNoHeader
     0.05s  0.22% 81.30%      0.20s  0.87%  runtime.pcvalue
     0.05s  0.22% 81.52%      0.21s  0.91%  runtime.scanObject
     0.04s  0.17% 81.69%      0.19s  0.82%  crypto/internal/fips140/sha256.consumeUint32 (inline)
     0.04s  0.17% 81.87%      0.15s  0.65%  database/sql.convertAssignRows
     0.04s  0.17% 82.04%      0.22s  0.95%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.03s  0.13% 82.17%      0.15s  0.65%  crypto/internal/fips140deps/byteorder.BEUint32 (inline)
     0.03s  0.13% 82.30%      0.42s  1.82%  github.com/jackc/pgx/v5/pgproto3.(*chunkReader).Next
     0.03s  0.13% 82.43%      0.13s  0.56%  runtime.makeslice
     0.03s  0.13% 82.56%      0.20s  0.87%  runtime.mallocgcSmallNoscan
     0.03s  0.13% 82.69%      0.29s  1.26%  runtime.newobject
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 19.4% / 78% | 0% | 0.0% | 8 | 2040 MB | run queue が積もっている |
| isu2 (2 cpu) | 1.6% / 10% | 0% | 0.0% | 3 | 3151 MB |  |
| isu3 (2 cpu) | 1.9% / 10% | 0% | 0.0% | 3 | 3225 MB |  |
| isu4 (2 cpu) | 43.8% / 100% | 34% | 3.7% | 30 | 104 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 32% / nginx 4% |
| isu2 | isureal 3% |
| isu3 | isureal 3% |
| isu4 | postgres 16% |

## Diff from previous run

スコア: 6389 → 7316 (+927)

`8cd1279..92ad1cc`

```
 webapp/go/admin_posting_window_api.go |  3 ++-
 webapp/go/image_signature.go          | 34 +++++++++++++++++++++++++++++++---
 2 files changed, 33 insertions(+), 4 deletions(-)
```
