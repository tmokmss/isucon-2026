# run 20261031-111006-d2acb70

## Summary

| | |
|---|---|
| score | 12786 |
| status | pass |
| ref | posting-window-fast |
| sha | d2acb70c3ccd |
| 所要時間 | 526s |
| 前回比 | -2882 (-18.4%) vs 20261031-110450-8f720c2 |
| ベスト比 | -12015 (-48.4%) vs 20261031-103811-cef5fe3 = 24801 |
| note | 投稿期間の開始を専用プール・1往復で |

## Bench

```
{"id":"3327","status":"done","portalStatus":"finished","passed":true,"score":12786,"raw":12786,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:17:46.346Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:17:46.418Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:17:49.271Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:18:05.701Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T02:18:12.048Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T02:18:12.147Z level=INFO msg="america 3 日目: 15 人"
time=2026-10-31T02:18:18.356Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T02:18:18.372Z level=INFO msg="america 4 日目: 15 人"
time=2026-10-31T02:18:18.547Z level=INFO msg="asia 4 日目: 18 人"
time=2026-10-31T02:18:22.081Z level=INFO msg="asia 5 日目: 18 人"
time=2026-10-31T02:18:22.246Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T02:18:25.316Z level=INFO msg="america 5 日目: 15 人"
time=2026-10-31T02:18:34.256Z level=INFO msg="europe 最終日（グランプリ）: 6 人"
time=2026-10-31T02:18:35.918Z level=INFO msg="asia 最終日（グランプリ）: 17 人"
time=2026-10-31T02:18:35.971Z level=INFO msg="america 最終日（グランプリ）: 14 人"
time=2026-10-31T02:18:43.231Z level=INFO msg="グランプリ: アジア 17 人, ヨーロッパ 6 人, アメリカ 14 人"
{
  "pass": true,
  "score": 12786,
  "messages": [
    "asia の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました",
    "europe の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました",
    "asia の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
68.8%	108	0	0	GET	/api/notifications	948.743	8.785	19.099	19.119	177
8.2%	91	0	0	GET	/api/me/export	112.806	1.240	9.645	9.645	2033729
5.8%	18	0	0	POST	/api/admin/posting-window	79.624	4.424	13.256	13.256	177
3.8%	3041	246	0	GET	^/api/images/[^/]+	53.072	0.017	0.219	0.447	36571
3.5%	71	0	0	POST	/api/posts	48.837	0.688	1.824	1.824	604
3.1%	1529	536	0	POST	/api/auth/login	43.028	0.028	0.406	0.999	172
2.0%	109	107	0	PUT	/api/me/profile-image	27.539	0.253	1.391	1.900	24
1.6%	702	0	0	GET	/api/memories	21.477	0.031	0.240	0.466	15355
0.8%	1	0	0	POST	/api/initialize	10.609	10.609	10.609	10.609	90
0.5%	103	8	0	POST	/api/friend-requests	6.748	0.066	0.635	0.707	109
0.5%	531	0	0	GET	^/api/feeds/[^/]+	6.346	0.012	0.155	0.267	2242
0.4%	99	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	5.315	0.054	0.693	0.693	116
0.3%	82	0	0	GET	/api/posting-window/current	4.346	0.053	0.269	0.269	237
0.2%	170	0	0	POST	/api/auth/logout	2.405	0.014	0.301	0.351	0
0.2%	72	2	0	GET	^/api/posts/me/current$	2.106	0.029	0.201	0.201	588
0.1%	34	1	0	POST	/api/auth/register	1.826	0.054	0.886	0.886	237
0.1%	92	0	0	GET	^/api/relationships/[^/]+	0.984	0.011	0.170	0.170	2836
0.1%	36	0	0	GET	/api/users	0.807	0.022	0.154	0.154	261
0.0%	197	0	0	GET	^/api/memories/[^/]+	0.582	0.003	0.058	0.096	589
0.0%	54	0	0	GET	^/api/users/[^/]+$	0.510	0.009	0.141	0.141	174
0.0%	68	0	0	GET	^/api/posts/[^/]+$	0.418	0.006	0.057	0.057	604
0.0%	3	0	0	PATCH	^/api/posts/[^/]+$	0.276	0.092	0.265	0.265	574
0.0%	7	0	0	PUT	^/api/admin/users/[^/]+/certification$	0.113	0.016	0.060	0.060	0
0.0%	4	0	0	DELETE	^/api/posts/[^/]+$	0.112	0.028	0.062	0.062	0
0.0%	1	0	0	PATCH	/api/me	0.099	0.099	0.099	0.099	236
0.0%	13	0	0	GET	^/api/posts/[^/]+/realmojis$	0.068	0.005	0.030	0.030	29
0.0%	13	0	0	PUT	^/api/follows/[^/]+$	0.061	0.005	0.008	0.008	0
0.0%	4	0	0	PUT	^/api/blocks/[^/]+$	0.053	0.013	0.021	0.021	0
0.0%	1	0	0	POST	/api/me/account-deletion	0.019	0.019	0.019	0.019	0
0.0%	1	0	0	POST	/api/admin/auth/logout	0.012	0.012	0.012	0.012	0
…（上位 30 / 全 34 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
64.6%	42	0	0	GET	/api/notifications	348.177	8.290	19.165	19.165	177
8.7%	311	0	0	GET	^/api/images/[^/]+	46.721	0.150	0.945	1.260	156440
5.6%	34	0	0	POST	/api/posts	30.274	0.890	1.802	1.802	618
5.0%	8	0	0	GET	/api/me/export	27.165	3.396	9.757	9.757	3311356
3.7%	158	0	0	GET	^/api/feeds/[^/]+	20.084	0.127	0.378	0.414	1145
2.8%	95	0	0	GET	/api/memories	15.295	0.161	1.820	1.820	16146
2.6%	79	0	0	POST	/api/auth/login	13.766	0.174	0.851	0.851	256
1.0%	37	0	0	GET	/api/posting-window/current	5.427	0.147	0.399	0.399	239
0.9%	34	0	0	GET	^/api/posts/me/current$	5.084	0.150	0.502	0.502	618
0.8%	36	3	0	POST	/api/friend-requests	4.543	0.126	0.282	0.282	108
0.6%	28	0	0	GET	^/api/relationships/[^/]+	3.382	0.121	0.347	0.347	1089
0.6%	23	1	0	GET	^/api/posts/[^/]+$	3.051	0.133	0.299	0.299	572
0.5%	20	0	0	GET	^/api/memories/[^/]+	2.594	0.130	0.365	0.365	1257
0.5%	22	0	0	GET	^/api/users/[^/]+$	2.488	0.113	0.182	0.182	168
0.4%	12	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	2.079	0.173	0.637	0.637	116
0.4%	19	0	0	GET	/api/users	2.040	0.107	0.136	0.136	251
0.3%	13	0	0	POST	/api/auth/logout	1.665	0.128	0.294	0.294	0
0.3%	5	0	0	POST	/api/auth/register	1.498	0.300	0.358	0.358	229
0.2%	5	0	0	POST	^/api/posts/[^/]+/realmojis$	0.972	0.194	0.310	0.310	409
0.1%	5	2	0	GET	/api/me	0.606	0.121	0.202	0.202	127
0.1%	4	0	0	PUT	^/api/follows/[^/]+$	0.604	0.151	0.160	0.160	0
0.1%	5	0	0	PATCH	^/api/posts/[^/]+$	0.604	0.121	0.149	0.149	602
0.1%	2	0	0	DELETE	^/api/realmojis/[^/]+$	0.372	0.186	0.263	0.263	0
0.0%	1	0	0	POST	/api/me/account-deletion	0.191	0.191	0.191	0.191	0
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.107	0.107	0.107	0.107	0
0.0%	1	0	0	GET	^/api/posts/[^/]+/comments$	0.101	0.101	0.101	0.101	29

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
56.7%	77	0	0	GET	/api/notifications	566.324	7.355	16.427	16.427	177
10.5%	693	6	0	GET	^/api/images/[^/]+	104.507	0.151	0.992	1.385	159833
8.3%	78	0	0	POST	/api/posts	83.395	1.069	2.523	2.523	638
6.1%	551	0	0	GET	^/api/feeds/[^/]+	60.588	0.110	0.193	0.265	1903
2.9%	203	0	0	GET	/api/memories	29.204	0.144	0.579	0.969	12462
2.8%	165	0	0	POST	/api/auth/login	27.929	0.169	0.818	0.859	253
2.6%	8	0	0	GET	/api/me/export	26.034	3.254	8.541	8.541	3253423
1.4%	93	0	0	GET	/api/posting-window/current	13.858	0.149	0.323	0.323	240
1.1%	101	3	0	GET	^/api/posts/[^/]+$	10.777	0.107	0.152	0.225	595
1.1%	93	10	0	POST	/api/friend-requests	10.767	0.116	0.277	0.277	106
1.0%	89	0	0	GET	^/api/relationships/[^/]+	9.682	0.109	0.330	0.330	723
0.9%	78	0	0	GET	^/api/posts/me/current$	8.952	0.115	0.207	0.207	638
0.7%	44	0	0	POST	^/api/posts/[^/]+/realmojis$	7.205	0.164	0.536	0.536	424
0.6%	59	0	0	GET	^/api/users/[^/]+$	6.419	0.109	0.246	0.246	167
0.5%	38	0	0	POST	/api/auth/logout	5.016	0.132	0.493	0.493	0
0.4%	40	0	0	GET	/api/users	4.368	0.109	0.143	0.143	244
0.3%	27	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	3.273	0.121	0.215	0.215	116
0.3%	26	0	0	GET	^/api/posts/[^/]+/realmojis$	2.834	0.109	0.179	0.179	442
0.3%	22	0	0	GET	^/api/memories/[^/]+	2.636	0.120	0.219	0.219	941
0.3%	14	0	0	POST	/api/auth/register	2.631	0.188	0.294	0.294	229
0.2%	14	0	0	POST	^/api/posts/[^/]+/comments$	2.196	0.157	0.368	0.368	422
0.2%	16	0	0	PUT	^/api/follows/[^/]+$	2.018	0.126	0.205	0.205	0
0.2%	17	0	0	GET	^/api/posts/[^/]+/comments$	1.839	0.108	0.126	0.126	203
0.1%	11	0	0	DELETE	^/api/realmojis/[^/]+$	1.232	0.112	0.145	0.145	0
0.1%	11	1	0	GET	/api/me	1.149	0.104	0.116	0.116	251
0.1%	8	0	0	PATCH	^/api/posts/[^/]+$	0.965	0.121	0.162	0.162	596
0.1%	7	0	0	PUT	/api/me/profile-image	0.882	0.126	0.161	0.161	0
0.1%	5	0	0	DELETE	^/api/posts/[^/]+$	0.713	0.143	0.289	0.289	0
0.1%	5	0	0	PATCH	/api/me	0.650	0.130	0.218	0.218	334
0.0%	3	0	0	DELETE	^/api/comments/[^/]+$	0.411	0.137	0.188	0.188	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 12028 クエリ / 82 種類 / 合計 202.082s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	66.2%	107	133.829	1250.74	9433.37	7175.1	141.8	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	16.5%	18	33.303	1850.18	13183.13	32565.6	1641.1	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
3	5.7%	831	11.568	13.92	236.45	111.8	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	5.0%	2602	10.011	3.85	23.36	10.5	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.6%	241	3.196	13.26	77.07	699.7	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
6	1.2%	1	2.345	2345.01	2345.01	29320.0	1199382.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.7%	1	1.342	1342.23	1342.23	62250.0	1.0	select count(*), count(distinct md5(i.data)) from realmojis r join image_objects i on i.id=r.image_id and i.created_at=r.created_at where r.created_at > (select…
8	0.5%	1	1.064	1063.88	1063.88	59977.0	1.0	select count(*), count(distinct md5(i.data)), count(distinct length(i.data)) from posts p join image_objects i on i.id=p.image_id and i.created_at=p.posted_at w…
9	0.5%	1	1.054	1053.79	1053.79	11528.0	529674.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
10	0.4%	183	0.778	4.25	21.73	40.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
11	0.2%	49	0.409	8.34	38.18	45.9	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
12	0.2%	2	0.387	193.68	379.61	20238.5	2513.5	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
13	0.2%	2	0.334	167.04	178.18	214363.5	2521.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.2%	1	0.320	319.95	319.95	19618.0	1.0	select count(*) from posts where created_at > (select created_at from initial_data_boundary)
15	0.1%	1	0.230	230.15	230.15	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4;
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 7c03052c2f137b2996f44b823585f66ac08a7ad6
Type: cpu
Time: 2026-10-31 11:17:26 JST
Duration: 60.20s, Total samples = 53.12s (88.23%)
Showing nodes accounting for 46.60s, 87.73% of 53.12s total
Dropped 691 nodes (cum <= 0.27s)
Showing top 40 nodes out of 175
      flat  flat%   sum%        cum   cum%
    12.11s 22.80% 22.80%     12.11s 22.80%  crypto/internal/fips140/sha256.blockAVX2
     7.75s 14.59% 37.39%      7.75s 14.59%  runtime.spanClass.sizeclass (inline)
     2.60s  4.89% 42.28%      4.29s  8.08%  image/jpeg.(*decoder).reconstructBlock
     2.50s  4.71% 46.99%     10.73s 20.20%  runtime.tryDeferToSpanScan
     2.01s  3.78% 50.77%      2.01s  3.78%  runtime.memmove
     1.40s  2.64% 53.41%      1.40s  2.64%  runtime.memclrNoHeapPointers
     1.36s  2.56% 55.97%      9.17s 17.26%  image/jpeg.(*decoder).processSOS
     1.32s  2.48% 58.45%      5.55s 10.45%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.23s  2.32% 60.77%      4.97s  9.36%  runtime.scanObjectsSmall
     1.07s  2.01% 62.78%      8.50s 16.00%  runtime.scanObject
     1.06s  2.00% 64.78%      1.06s  2.00%  internal/runtime/syscall/linux.Syscall6
     1.04s  1.96% 66.74%      1.04s  1.96%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.89s  1.68% 68.41%      1.63s  3.07%  image/jpeg.(*decoder).decodeHuffman
     0.87s  1.64% 70.05%      0.89s  1.68%  image/jpeg.(*decoder).receiveExtend
     0.84s  1.58% 71.63%     13.94s 26.24%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.69s  1.30% 72.93%      0.69s  1.30%  aeshashbody
     0.67s  1.26% 74.19%     13.44s 25.30%  crypto/internal/fips140/sha256.(*Digest).Write
     0.65s  1.22% 75.41%      0.86s  1.62%  image/jpeg.idctCols
     0.64s  1.20% 76.62%      0.80s  1.51%  image/jpeg.idctRows
     0.63s  1.19% 77.80%      0.63s  1.19%  memeqbody
     0.59s  1.11% 78.92%     14.84s 27.94%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.48s   0.9% 79.82%     15.93s 29.99%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.48s   0.9% 80.72%      0.75s  1.41%  image/jpeg.(*decoder).ensureNBits
     0.47s  0.88% 81.61%      5.94s 11.18%  runtime.scanSpan
     0.44s  0.83% 82.44%      0.64s  1.20%  example.com/isureal-company/isureal.postLess
     0.37s   0.7% 83.13%      0.37s   0.7%  image/jpeg.dctBox (inline)
     0.35s  0.66% 83.79%      2.16s  4.07%  runtime.mapaccess1_faststr
     0.32s   0.6% 84.39%      0.32s   0.6%  runtime.extractHeapBitsSmall
     0.25s  0.47% 84.86%      0.27s  0.51%  image/jpeg.(*decoder).readByteStuffedByte
     0.22s  0.41% 85.28%      0.70s  1.32%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.22s  0.41% 85.69%      0.85s  1.60%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.16s   0.3% 85.99%      0.66s  1.24%  runtime.mallocgcSmallScanNoHeader
     0.15s  0.28% 86.28%     12.26s 23.08%  crypto/internal/fips140/sha256.block
     0.14s  0.26% 86.54%     17.15s 32.29%  example.com/isureal-company/isureal.imageSignature
     0.13s  0.24% 86.78%      6.10s 11.48%  runtime.mallocgc
     0.12s  0.23% 87.01%      0.35s  0.66%  database/sql.convertAssignRows
     0.12s  0.23% 87.24%      0.46s  0.87%  runtime.scanObjectSmall
     0.11s  0.21% 87.44%      0.27s  0.51%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.08s  0.15% 87.59%      0.27s  0.51%  crypto/internal/fips140/sha256.consumeUint32 (inline)
     0.07s  0.13% 87.73%      1.79s  3.37%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 51.9% / 100% | 22% | 0.2% | 5 | 559 MB | run queue が積もっている |
| isu2 (2 cpu) | 0.8% / 4% | 0% | 0.2% | 1 | 300 MB |  |
| isu3 (2 cpu) | 5.0% / 98% | 2% | 0.4% | 3 | 279 MB |  |
| isu4 (2 cpu) | 61.5% / 100% | 40% | 36.6% | 7 | 95 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 93% / nginx 8% |
| isu2 | 目立つプロセス無し |
| isu3 | kworker/u8:0+events_unbound 2% / kworker/0:0+inode_switch_wbs 2% / isureal 2% |
| isu4 | postgres 43% |

## Diff from previous run

スコア: 15668 → 12786 (-2882)

`8f720c2..d2acb70`

```
 webapp/go/admin_posting_window_api.go | 47 ++++++++++++++++++++++++-----------
 webapp/go/cmd/image-export/main.go    | 11 +-------
 webapp/go/image_local.go              | 37 +++++++++++++++++++++------
 3 files changed, 63 insertions(+), 32 deletions(-)
```
