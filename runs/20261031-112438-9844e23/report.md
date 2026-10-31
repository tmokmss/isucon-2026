# run 20261031-112438-9844e23

## Summary

| | |
|---|---|
| score | 18671 |
| status | pass |
| ref | offload-cpu |
| sha | 9844e230cdda |
| 所要時間 | 192s |
| 前回比 | -604 (-3.1%) vs 20261031-111738-5948f35 |
| ベスト比 | -6130 (-24.7%) vs 20261031-103811-cef5fe3 = 24801 |
| note | offload-with-slots |

## Bench

```
{"id":"3909","status":"done","portalStatus":"finished","passed":true,"score":18671,"raw":18671,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:26:38.902Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:26:38.935Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:26:41.890Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:26:47.578Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T02:26:47.881Z level=INFO msg="europe 2 日目: 9 人"
time=2026-10-31T02:26:51.109Z level=INFO msg="america 2 日目: 9 人"
time=2026-10-31T02:26:54.932Z level=INFO msg="asia 3 日目: 18 人"
time=2026-10-31T02:26:56.886Z level=INFO msg="europe 3 日目: 19 人"
time=2026-10-31T02:26:59.892Z level=INFO msg="america 3 日目: 18 人"
time=2026-10-31T02:27:03.873Z level=INFO msg="asia 4 日目: 23 人"
time=2026-10-31T02:27:06.109Z level=INFO msg="europe 4 日目: 37 人"
time=2026-10-31T02:27:08.890Z level=INFO msg="america 4 日目: 29 人"
time=2026-10-31T02:27:13.801Z level=INFO msg="asia 5 日目: 35 人"
time=2026-10-31T02:27:15.790Z level=INFO msg="user moved" user=nokasol-5460 from=america to=europe
time=2026-10-31T02:27:17.131Z level=INFO msg="europe 5 日目: 54 人"
time=2026-10-31T02:27:18.523Z level=INFO msg="america 5 日目: 28 人"
time=2026-10-31T02:27:27.385Z level=INFO msg="europe 最終日（グランプリ）: 53 人"
time=2026-10-31T02:27:29.242Z level=INFO msg="asia 最終日（グランプリ）: 34 人"
time=2026-10-31T02:27:30.380Z level=INFO msg="america 最終日（グランプリ）: 27 人"
time=2026-10-31T02:27:35.864Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 22 人, アメリカ 0 人"
{
  "pass": true,
  "score": 18671,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.9%	131	0	0	GET	/api/notifications	1135.565	8.668	15.463	15.465	175
17.1%	1879	1061	0	POST	/api/auth/login	511.665	0.272	1.893	3.771	125
12.7%	173	161	0	PUT	/api/me/profile-image	381.826	2.207	9.349	9.540	26
10.5%	124	6	0	POST	/api/posts	315.302	2.543	7.816	8.316	586
7.5%	2971	407	0	GET	^/api/images/[^/]+	224.130	0.075	0.965	3.773	54514
2.2%	644	0	0	GET	/api/memories	64.492	0.100	0.594	1.243	14697
1.9%	172	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	57.315	0.333	4.765	5.028	116
1.6%	207	20	0	POST	/api/friend-requests	48.518	0.234	2.759	4.474	107
1.6%	59	0	0	GET	/api/me/export	47.509	0.805	6.797	6.797	2042850
1.3%	51	1	0	POST	/api/auth/register	38.146	0.748	4.725	4.725	234
1.2%	880	0	0	GET	^/api/feeds/[^/]+	35.031	0.040	0.579	1.126	2374
0.8%	144	0	0	GET	/api/posting-window/current	23.200	0.161	0.989	1.103	237
0.7%	18	0	0	POST	/api/admin/posting-window	19.762	1.098	3.516	3.516	176
0.4%	1	0	0	POST	/api/initialize	12.267	12.267	12.267	12.267	90
0.4%	156	0	0	POST	/api/auth/logout	12.142	0.078	2.156	2.687	0
0.4%	23	0	0	PUT	^/api/follows/[^/]+$	11.807	0.513	3.066	3.066	0
0.3%	146	0	0	GET	^/api/relationships/[^/]+	7.829	0.054	0.470	0.708	2360
0.2%	22	0	0	POST	^/api/posts/[^/]+/realmojis$	6.366	0.289	3.583	3.583	415
0.2%	86	0	0	GET	/api/users	5.825	0.068	0.663	0.663	225
0.2%	144	1	0	GET	^/api/posts/[^/]+$	5.511	0.038	0.529	0.626	596
0.2%	135	0	0	GET	^/api/memories/[^/]+	5.071	0.038	0.261	0.418	775
0.2%	93	1	0	GET	^/api/posts/me/current$	4.829	0.052	0.551	0.551	606
0.2%	5	0	0	PUT	^/api/blocks/[^/]+$	4.503	0.901	3.948	3.948	0
0.1%	7	0	0	PATCH	/api/me	3.669	0.524	1.282	1.282	263
0.1%	49	0	0	POST	^/api/posts/[^/]+/comments$	2.902	0.059	1.118	1.118	273
0.1%	19	3	0	GET	/api/me	2.745	0.144	0.636	0.636	251
0.1%	1	0	0	POST	/api/admin/auth/logout	2.396	2.396	2.396	2.396	0
0.1%	83	0	0	GET	^/api/users/[^/]+$	1.863	0.022	0.417	0.417	188
0.0%	7	0	0	DELETE	^/api/realmojis/[^/]+$	1.276	0.182	0.824	0.824	0
0.0%	9	0	0	PUT	^/api/admin/users/[^/]+/certification$	1.081	0.120	0.621	0.621	0
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.7%	126	0	0	GET	/api/notifications	1064.165	8.446	11.096	11.690	165
20.2%	151	1	0	POST	/api/posts	540.800	3.581	9.610	9.929	639
5.9%	747	6	0	GET	^/api/images/[^/]+	158.788	0.213	1.569	2.154	149542
5.5%	232	0	0	POST	/api/auth/login	146.727	0.632	4.116	4.323	249
4.8%	286	16	0	POST	/api/friend-requests	129.777	0.454	3.294	3.965	110
4.2%	480	0	0	GET	^/api/feeds/[^/]+	113.425	0.236	0.898	1.486	1618
3.8%	316	0	0	GET	/api/memories	103.177	0.327	1.317	1.426	8546
3.2%	51	0	0	POST	/api/auth/register	85.713	1.681	7.275	7.275	228
2.1%	209	0	0	GET	/api/posting-window/current	56.383	0.270	0.766	1.053	238
1.8%	15	0	0	GET	/api/me/export	47.069	3.138	10.245	10.245	2152812
1.5%	46	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	41.571	0.904	5.357	5.357	116
1.2%	151	0	0	GET	/api/users	32.380	0.214	0.847	0.916	232
1.2%	118	1	0	GET	^/api/posts/me/current$	31.534	0.267	0.912	0.940	635
1.0%	115	0	0	GET	^/api/relationships/[^/]+	26.005	0.226	0.623	0.932	672
0.6%	44	0	0	POST	/api/auth/logout	14.878	0.338	3.692	3.692	0
0.5%	22	0	0	PUT	/api/me/profile-image	12.244	0.557	2.502	2.502	0
0.4%	24	0	0	PATCH	/api/me	11.059	0.461	1.891	1.891	266
0.4%	53	6	0	GET	^/api/posts/[^/]+$	10.227	0.193	0.469	0.469	539
0.4%	25	0	0	POST	^/api/posts/[^/]+/comments$	9.981	0.399	2.238	2.238	319
0.3%	24	0	0	PUT	^/api/follows/[^/]+$	8.504	0.354	1.608	1.608	0
0.3%	28	0	0	GET	^/api/memories/[^/]+	7.465	0.267	0.896	0.896	911
0.2%	27	2	0	GET	/api/me	5.240	0.194	0.561	0.561	274
0.1%	5	0	0	PATCH	^/api/posts/[^/]+$	3.920	0.784	2.196	2.196	598
0.1%	5	0	0	POST	^/api/posts/[^/]+/realmojis$	3.788	0.758	1.335	1.335	427
0.1%	6	0	0	DELETE	^/api/posts/[^/]+$	3.759	0.626	2.125	2.125	0
0.1%	20	0	0	GET	^/api/users/[^/]+$	3.453	0.173	0.538	0.538	177
0.1%	19	0	0	GET	^/api/posts/[^/]+/realmojis$	2.882	0.152	0.366	0.366	153
0.1%	19	0	0	GET	^/api/posts/[^/]+/comments$	2.796	0.147	0.288	0.288	499
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	2.322	0.774	1.465	1.465	0
0.1%	7	0	0	DELETE	^/api/comments/[^/]+$	1.898	0.271	0.489	0.489	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.7%	95	0	0	GET	/api/notifications	845.542	8.900	12.075	12.075	178
18.1%	112	4	0	POST	/api/posts	350.410	3.129	7.341	7.381	623
6.5%	737	3	0	GET	^/api/images/[^/]+	125.434	0.170	0.960	1.590	165255
5.6%	521	0	0	GET	^/api/feeds/[^/]+	108.341	0.208	0.716	1.255	2554
5.2%	175	0	0	POST	/api/auth/login	100.870	0.576	4.581	4.824	252
4.0%	194	13	0	POST	/api/friend-requests	76.491	0.394	3.231	4.050	109
3.6%	226	0	0	GET	/api/memories	69.172	0.306	1.283	1.540	12490
2.3%	135	0	0	GET	/api/posting-window/current	44.464	0.329	1.065	1.187	240
2.0%	40	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	39.163	0.979	3.789	3.789	116
1.1%	41	1	0	POST	^/api/posts/[^/]+/realmojis$	21.031	0.513	2.265	2.265	431
0.9%	92	0	0	GET	/api/users	18.292	0.199	2.188	2.188	214
0.9%	87	1	0	GET	^/api/posts/me/current$	16.843	0.194	0.483	0.483	634
0.8%	84	0	0	GET	^/api/relationships/[^/]+	16.392	0.195	0.507	0.507	857
0.8%	80	9	0	GET	^/api/posts/[^/]+$	15.008	0.188	0.683	0.683	543
0.5%	7	0	0	GET	/api/me/export	10.362	1.480	1.967	1.967	1956519
0.5%	31	0	0	POST	^/api/posts/[^/]+/comments$	9.803	0.316	1.574	1.574	319
0.5%	28	0	0	POST	/api/auth/register	9.524	0.340	1.451	1.451	228
0.5%	44	0	0	GET	^/api/users/[^/]+$	9.275	0.211	0.844	0.844	191
0.4%	29	0	0	POST	/api/auth/logout	7.543	0.260	1.379	1.379	0
0.3%	17	0	0	PUT	/api/me/profile-image	6.072	0.357	1.232	1.232	0
0.2%	15	0	0	PATCH	/api/me	4.697	0.313	0.775	0.775	263
0.2%	13	0	0	PUT	^/api/follows/[^/]+$	4.480	0.345	1.285	1.285	0
0.2%	21	0	0	GET	^/api/memories/[^/]+	4.251	0.202	0.882	0.882	1340
0.2%	7	0	0	DELETE	^/api/realmojis/[^/]+$	4.221	0.603	1.524	1.524	0
0.2%	22	1	0	GET	/api/me	4.170	0.190	0.487	0.487	275
0.2%	23	3	0	GET	^/api/posts/[^/]+/realmojis$	3.292	0.143	0.446	0.446	345
0.1%	8	1	0	PATCH	^/api/posts/[^/]+$	2.355	0.294	0.481	0.481	593
0.1%	15	0	0	GET	^/api/posts/[^/]+/comments$	2.331	0.155	0.427	0.427	564
0.1%	3	0	0	DELETE	^/api/posts/[^/]+$	1.471	0.490	0.590	0.590	0
0.1%	1	0	0	DELETE	/api/me/profile-image	1.230	1.230	1.230	1.230	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 14351 クエリ / 79 種類 / 合計 64.577s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	45.7%	81	29.532	364.60	3972.02	6578.3	131.4	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	14.1%	18	9.110	506.11	3016.85	32779.3	1648.7	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
3	14.1%	1275	9.074	7.12	358.23	119.6	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	6.8%	498	4.364	8.76	73.66	715.0	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
5	6.1%	2232	3.941	1.77	15.25	10.8	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
6	3.7%	1	2.384	2384.36	2384.36	29367.0	1199433.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	2.4%	1	1.558	1558.42	1558.42	19778.0	625629.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	1.7%	1	1.101	1100.80	1100.80	11528.0	529676.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	0.9%	379	0.588	1.55	5.23	35.5	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.5%	2	0.343	171.68	186.19	214746.0	2526.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.4%	2	0.279	139.45	270.31	20231.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.4%	1	0.277	277.23	277.23	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
13	0.3%	67	0.213	3.18	7.59	41.9	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
14	0.3%	1	0.190	190.43	190.43	40112.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.3%	1	0.189	189.03	189.03	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…

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
Build ID: 7f0ae8e060ff4cf87ab6afc34a237b596364f2a0
Type: cpu
Time: 2026-10-31 11:26:18 JST
Duration: 60.18s, Total samples = 67.19s (111.65%)
Showing nodes accounting for 57.68s, 85.85% of 67.19s total
Dropped 893 nodes (cum <= 0.34s)
Showing top 40 nodes out of 205
      flat  flat%   sum%        cum   cum%
    13.37s 19.90% 19.90%     13.37s 19.90%  runtime.spanClass.sizeclass (inline)
     9.62s 14.32% 34.22%      9.62s 14.32%  crypto/internal/fips140/sha256.blockAVX2
     4.37s  6.50% 40.72%     18.49s 27.52%  runtime.tryDeferToSpanScan
     3.08s  4.58% 45.30%      3.08s  4.58%  runtime.memclrNoHeapPointers
     2.95s  4.39% 49.69%      4.75s  7.07%  image/jpeg.(*decoder).reconstructBlock
     2.60s  3.87% 53.56%      2.60s  3.87%  runtime.memmove
     2.25s  3.35% 56.91%      8.74s 13.01%  runtime.scanObjectsSmall
     1.91s  2.84% 59.76%     14.78s 22.00%  runtime.scanObject
     1.88s  2.80% 62.55%      1.88s  2.80%  internal/runtime/syscall/linux.Syscall6
     1.38s  2.05% 64.61%      5.66s  8.42%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.32s  1.96% 66.57%     10.38s 15.45%  image/jpeg.(*decoder).processSOS
     1.15s  1.71% 68.28%      1.15s  1.71%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.91s  1.35% 69.64%      1.80s  2.68%  image/jpeg.(*decoder).decodeHuffman
     0.87s  1.29% 70.93%      0.89s  1.32%  image/jpeg.(*decoder).receiveExtend
     0.85s  1.27% 72.20%      0.85s  1.27%  aeshashbody
     0.84s  1.25% 73.45%     10.26s 15.27%  runtime.scanSpan
     0.74s  1.10% 74.55%      0.97s  1.44%  image/jpeg.idctCols
     0.67s     1% 75.55%      0.90s  1.34%  image/jpeg.(*decoder).ensureNBits
     0.60s  0.89% 76.44%        11s 16.37%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.59s  0.88% 77.32%      0.79s  1.18%  image/jpeg.idctRows
     0.52s  0.77% 78.09%      0.52s  0.77%  memeqbody
     0.52s  0.77% 78.87%      0.52s  0.77%  runtime.extractHeapBitsSmall
     0.49s  0.73% 79.60%     10.70s 15.92%  crypto/internal/fips140/sha256.(*Digest).Write
     0.48s  0.71% 80.31%      2.08s  3.10%  runtime.mapaccess1_faststr
     0.43s  0.64% 80.95%      0.43s  0.64%  image/jpeg.dctBox (inline)
     0.40s   0.6% 81.54%     11.62s 17.29%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.36s  0.54% 82.08%      0.43s  0.64%  runtime.(*spanInlineMarkBits).tryAcquire
     0.33s  0.49% 82.57%      0.55s  0.82%  example.com/isureal-company/isureal.postLess
     0.29s  0.43% 83.00%      0.35s  0.52%  runtime.findObject
     0.25s  0.37% 83.38%     12.40s 18.46%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.22s  0.33% 83.70%      0.63s  0.94%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.22s  0.33% 84.03%      1.10s  1.64%  runtime.mallocgcSmallScanNoHeader
     0.19s  0.28% 84.31%     15.84s 23.57%  runtime.mallocgc
     0.18s  0.27% 84.58%      9.80s 14.59%  crypto/internal/fips140/sha256.block
     0.17s  0.25% 84.83%      0.56s  0.83%  database/sql.convertAssignRows
     0.16s  0.24% 85.07%      0.66s  0.98%  runtime.scanObjectSmall
     0.15s  0.22% 85.30%      1.29s  1.92%  runtime.mapassign_faststr
     0.14s  0.21% 85.50%      2.92s  4.35%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.13s  0.19% 85.70%      0.44s  0.65%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.10s  0.15% 85.85%      0.45s  0.67%  runtime.(*sweepLocked).sweep
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 68.9% / 100% | 50% | 0.6% | 6 | 138 MB | CPU が飽和 |
| isu2 (2 cpu) | 4.6% / 26% | 0% | 0.4% | 4 | 176 MB |  |
| isu3 (2 cpu) | 4.6% / 21% | 0% | 0.5% | 2 | 191 MB |  |
| isu4 (2 cpu) | 29.7% / 93% | 1% | 18.4% | 7 | 108 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 118% / nginx 11% |
| isu2 | isureal 8% |
| isu3 | isureal 7% |
| isu4 | postgres 19% |

## Diff from previous run

スコア: 19275 → 18671 (-604)

`5948f35..9844e23`

```
 conf/isu1/home/isucon/env.sh                    |   1 +
 conf/isu4/etc/systemd/system/isureal-go.service |  17 ++
 webapp/go/cmd/isureal/main.go                   |   6 +
 webapp/go/export_api.go                         |   9 +
 webapp/go/image_signature.go                    |  33 ++-
 webapp/go/memory_api.go                         |  18 ++
 webapp/go/offload.go                            | 258 ++++++++++++++++++++++++
 webapp/go/post_api.go                           |   8 +-
 webapp/go/profile_image_api.go                  |  11 +-
 webapp/go/realmoji_api.go                       |   8 +-
 10 files changed, 341 insertions(+), 28 deletions(-)
```
