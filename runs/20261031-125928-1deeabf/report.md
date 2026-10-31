# run 20261031-125928-1deeabf

## Summary

| | |
|---|---|
| score | 101157 |
| status | pass |
| ref | tcp-slow-start-idle |
| sha | 1deeabf08fa9 |
| 所要時間 | 132s |
| 前回比 | +36457 (56.3%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +36457 (56.3%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | tcp_slow_start_after_idle=0 1回目 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	147	414	931	930
europe	7	18	49	119	118	117
america	6	18	42	87	88	87
```

```
{"id":"10609","status":"done","portalStatus":"finished","passed":true,"score":101157,"raw":101157,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:00:09.555Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:00:12.522Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:00:15.518Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:00:18.544Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T04:00:20.341Z level=INFO msg="user moved" user=nehemu5184 from=asia to=america
time=2026-10-31T04:00:21.525Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T04:00:24.517Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T04:00:26.557Z level=INFO msg="user moved" user=reri5433 from=europe to=asia
time=2026-10-31T04:00:27.541Z level=INFO msg="asia 3 日目: 147 人"
time=2026-10-31T04:00:29.798Z level=INFO msg="user moved" user=reri5433 from=asia to=europe
time=2026-10-31T04:00:30.480Z level=INFO msg="user moved" user=torlia.5976 from=asia to=europe
time=2026-10-31T04:00:30.522Z level=INFO msg="europe 3 日目: 49 人"
time=2026-10-31T04:00:30.797Z level=INFO msg="user moved" user=chiwaeva6033 from=asia to=europe
time=2026-10-31T04:00:33.548Z level=INFO msg="america 3 日目: 42 人"
time=2026-10-31T04:00:36.553Z level=INFO msg="asia 4 日目: 414 人"
time=2026-10-31T04:00:39.671Z level=INFO msg="europe 4 日目: 119 人"
time=2026-10-31T04:00:41.674Z level=INFO msg="user moved" user=resawa_7204 from=asia to=europe
time=2026-10-31T04:00:42.798Z level=INFO msg="america 4 日目: 87 人"
time=2026-10-31T04:00:44.370Z level=INFO msg="user moved" user=reri5433 from=europe to=asia
time=2026-10-31T04:00:44.825Z level=INFO msg="user moved" user=chiwaeva6033 from=europe to=america
time=2026-10-31T04:00:45.991Z level=INFO msg="asia 5 日目: 931 人"
time=2026-10-31T04:00:48.610Z level=INFO msg="europe 5 日目: 118 人"
time=2026-10-31T04:00:51.561Z level=INFO msg="america 5 日目: 88 人"
time=2026-10-31T04:01:01.034Z level=INFO msg="asia 最終日（グランプリ）: 930 人"
time=2026-10-31T04:01:01.285Z level=INFO msg="europe 最終日（グランプリ）: 117 人"
time=2026-10-31T04:01:01.332Z level=INFO msg="america 最終日（グランプリ）: 87 人"
time=2026-10-31T04:01:09.550Z level=INFO msg="グランプリ: アジア 327 人, ヨーロッパ 19 人, アメリカ 9 人"
{
  "pass": true,
  "score": 101157,
  "messages": [
    "reor.5110: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "annu.10550: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "tsuku5057: 投稿 37fd3dc8-8a5c-455a-a00f-e16a792d0fc1 に RealMoji を送れません: POST /api/posts/37fd3dc8-8a5c-455a-a00f-e16a792d0fc1/realmojis: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
31.6%	1445	0	0	GET	/api/notifications	11982.244	8.292	15.157	15.180	167
16.0%	2040	41	0	POST	/api/posts	6081.170	2.981	10.451	14.110	637
14.2%	4718	66	0	POST	/api/friend-requests	5388.510	1.142	10.082	17.202	114
8.7%	12756	698	0	GET	^/api/images/[^/]+	3284.103	0.257	4.622	14.386	138632
7.3%	5493	1589	0	POST	/api/auth/login	2758.268	0.502	6.258	13.215	181
5.0%	1540	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1902.993	1.236	9.586	15.065	116
3.2%	950	1	0	POST	/api/auth/register	1217.372	1.281	6.681	9.469	227
2.1%	629	114	0	PUT	/api/me/profile-image	780.033	1.240	6.877	11.045	4
1.7%	620	0	0	PUT	^/api/follows/[^/]+$	646.934	1.043	7.690	15.069	0
1.7%	548	24	0	POST	^/api/posts/[^/]+/comments$	632.835	1.155	8.965	16.377	327
1.5%	383	18	0	POST	^/api/posts/[^/]+/realmojis$	575.122	1.502	11.775	18.090	473
1.3%	478	0	0	PATCH	/api/me	478.179	1.000	5.809	10.844	281
1.1%	3217	0	0	GET	/api/memories	432.311	0.134	0.809	1.058	10777
0.7%	185	7	0	PATCH	^/api/posts/[^/]+$	283.952	1.535	8.419	8.626	624
0.7%	575	0	0	POST	/api/auth/logout	278.574	0.484	5.383	10.932	0
0.7%	11401	0	0	GET	^/api/feeds/[^/]+	260.560	0.023	0.182	4.799	2060
0.6%	220	0	0	GET	/api/me/export	208.795	0.949	5.242	9.931	1905545
0.5%	81	0	0	DELETE	^/api/posts/[^/]+$	177.016	2.185	8.738	8.738	0
0.4%	2947	0	0	GET	/api/posting-window/current	134.831	0.046	0.279	4.425	237
0.2%	70	0	0	DELETE	^/api/realmojis/[^/]+$	74.401	1.063	9.646	9.646	0
0.1%	90	0	0	DELETE	^/api/comments/[^/]+$	56.811	0.631	4.558	4.558	0
0.1%	1669	67	0	GET	^/api/posts/me/current$	52.097	0.031	0.215	0.353	623
0.1%	2298	0	0	GET	/api/users	50.539	0.022	0.167	1.075	237
0.1%	555	0	0	GET	^/api/memories/[^/]+	31.832	0.057	0.574	0.790	614
0.1%	1564	0	0	GET	^/api/relationships/[^/]+	31.055	0.020	0.169	0.276	799
0.1%	16	0	0	PUT	^/api/admin/users/[^/]+/certification$	30.236	1.890	11.708	11.708	0
0.1%	27	0	0	PUT	^/api/blocks/[^/]+$	21.717	0.804	4.257	4.257	0
0.1%	1560	15	0	GET	^/api/posts/[^/]+$	20.147	0.013	0.134	0.231	629
0.1%	1	0	0	POST	/api/initialize	19.678	19.678	19.678	19.678	90
0.0%	7	0	0	DELETE	^/api/follows/[^/]+$	12.871	1.839	10.821	10.821	0
…（上位 30 / 全 43 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.4%	305	1	0	GET	/api/notifications	2595.879	8.511	12.260	12.260	177
17.2%	363	5	0	POST	/api/posts	1164.343	3.208	10.171	10.517	630
8.5%	654	9	0	POST	/api/friend-requests	578.481	0.885	6.556	12.297	114
5.9%	232	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	402.049	1.733	11.616	13.316	116
5.9%	1639	13	0	GET	^/api/images/[^/]+	396.424	0.242	1.852	2.677	173520
5.0%	440	0	0	POST	/api/auth/login	338.121	0.768	5.702	12.931	246
3.7%	1301	0	0	GET	^/api/feeds/[^/]+	250.640	0.193	0.661	1.334	2316
2.0%	497	0	0	GET	/api/posting-window/current	137.475	0.277	0.758	1.218	239
1.8%	457	0	0	GET	/api/memories	118.819	0.260	1.283	1.322	13363
1.3%	72	0	0	POST	/api/auth/logout	86.045	1.195	10.334	10.334	0
1.0%	62	3	0	POST	^/api/posts/[^/]+/realmojis$	67.714	1.092	12.228	12.228	398
1.0%	289	29	0	GET	^/api/posts/me/current$	64.443	0.223	0.693	1.038	574
0.9%	115	0	0	POST	/api/auth/register	60.862	0.529	1.798	2.269	228
0.9%	61	4	0	POST	^/api/posts/[^/]+/comments$	60.059	0.985	9.175	9.175	320
0.8%	291	0	0	GET	/api/users	56.950	0.196	0.678	1.076	238
0.7%	21	0	0	GET	/api/me/export	48.955	2.331	8.276	8.276	2050735
0.7%	36	3	0	PATCH	^/api/posts/[^/]+$	46.849	1.301	6.861	6.861	574
0.6%	57	0	0	PUT	/api/me/profile-image	43.166	0.757	4.576	4.576	0
0.5%	204	0	0	GET	^/api/relationships/[^/]+	36.968	0.181	0.643	1.165	1091
0.5%	222	10	0	GET	^/api/posts/[^/]+$	36.899	0.166	0.595	1.120	596
0.5%	96	0	0	PUT	^/api/follows/[^/]+$	32.042	0.334	1.944	1.944	0
0.3%	14	0	0	DELETE	^/api/posts/[^/]+$	21.845	1.560	8.172	8.172	0
0.3%	14	0	0	DELETE	^/api/comments/[^/]+$	19.888	1.421	8.521	8.521	0
0.3%	51	0	0	PATCH	/api/me	17.633	0.346	1.434	1.434	279
0.2%	88	0	0	GET	^/api/users/[^/]+$	15.620	0.178	0.625	0.625	192
0.2%	54	0	0	GET	^/api/memories/[^/]+	15.055	0.279	1.077	1.077	913
0.2%	88	1	0	GET	^/api/posts/[^/]+/realmojis$	12.093	0.137	0.594	0.594	597
0.2%	4	0	0	PUT	^/api/blocks/[^/]+$	12.062	3.015	11.503	11.503	0
0.2%	75	0	0	GET	^/api/posts/[^/]+/comments$	12.002	0.160	0.619	0.619	1440
0.2%	63	2	0	GET	/api/me	11.982	0.190	0.593	0.593	291
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.8%	219	1	0	GET	/api/notifications	1684.062	7.690	9.816	9.818	180
16.4%	254	4	0	POST	/api/posts	770.367	3.033	9.888	10.107	634
9.8%	442	6	0	POST	/api/friend-requests	463.105	1.048	7.895	12.177	114
5.6%	1120	8	0	GET	^/api/images/[^/]+	261.932	0.234	1.131	1.935	176202
5.3%	168	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	250.108	1.489	10.834	13.262	116
4.9%	313	0	0	POST	/api/auth/login	230.783	0.737	6.687	7.650	251
3.5%	790	0	0	GET	^/api/feeds/[^/]+	163.536	0.207	0.745	1.284	2018
2.2%	86	10	0	POST	^/api/posts/[^/]+/realmojis$	102.240	1.189	9.702	9.702	406
2.0%	328	0	0	GET	/api/memories	92.270	0.281	1.341	4.534	12376
1.9%	344	0	0	GET	/api/posting-window/current	89.617	0.261	1.043	1.221	240
1.5%	57	0	0	POST	/api/auth/logout	68.559	1.203	8.858	8.858	0
1.0%	198	25	0	GET	^/api/posts/me/current$	46.796	0.236	0.783	1.025	567
1.0%	42	0	0	PUT	/api/me/profile-image	46.337	1.103	5.941	5.941	0
0.9%	57	0	0	PUT	^/api/follows/[^/]+$	43.400	0.761	5.039	5.039	0
0.9%	85	0	0	POST	/api/auth/register	43.159	0.508	8.466	8.466	228
0.9%	41	5	0	POST	^/api/posts/[^/]+/comments$	42.993	1.049	5.265	5.265	323
0.9%	213	0	0	GET	/api/users	41.587	0.195	0.699	0.709	239
0.8%	13	0	0	DELETE	^/api/posts/[^/]+$	37.252	2.866	5.942	5.942	0
0.7%	19	0	0	GET	/api/me/export	34.761	1.830	5.827	5.827	2959023
0.7%	20	1	0	PATCH	^/api/posts/[^/]+$	32.207	1.610	11.232	11.232	602
0.7%	47	0	0	PATCH	/api/me	31.613	0.673	5.210	5.210	276
0.6%	138	10	0	GET	^/api/posts/[^/]+$	25.907	0.188	0.554	0.742	590
0.5%	139	0	0	GET	^/api/relationships/[^/]+	25.628	0.184	0.700	0.737	1022
0.5%	10	0	0	DELETE	^/api/realmojis/[^/]+$	22.369	2.237	5.899	5.899	0
0.2%	45	0	0	GET	^/api/memories/[^/]+	11.231	0.250	0.564	0.564	1144
0.2%	56	0	0	GET	^/api/users/[^/]+$	11.038	0.197	0.708	0.708	206
0.2%	56	0	0	GET	^/api/posts/[^/]+/realmojis$	9.768	0.174	0.703	0.703	845
0.2%	46	1	0	GET	/api/me	8.494	0.185	0.754	0.754	293
0.1%	27	0	0	GET	^/api/posts/[^/]+/comments$	6.428	0.238	0.731	0.731	1374
0.1%	7	0	0	DELETE	^/api/comments/[^/]+$	5.847	0.835	4.175	4.175	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 64180 クエリ / 81 種類 / 合計 494.236s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	65.6%	3722	324.135	87.09	665.90	681.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	19.0%	260	93.938	361.30	1795.40	6932.5	127.6	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.3%	2389	26.372	11.04	996.66	118.5	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.3%	8294	21.142	2.55	45.15	10.8	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.5%	2	7.316	3658.04	7161.53	238466.5	2807.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
6	1.5%	2607	7.265	2.79	35.32	25.3	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
7	0.5%	1	2.230	2229.86	2229.86	30338.0	1199964.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.4%	500	2.147	4.29	33.46	29.4	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.4%	1	1.805	1805.16	1805.16	22371.0	627661.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
10	0.3%	617	1.280	2.07	18.52	20.8	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.3%	5716	1.274	0.22	50.44	20.1	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
12	0.2%	1	0.761	760.59	760.59	11998.0	530174.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
13	0.1%	4657	0.561	0.12	7.31	11.1	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
14	0.1%	1865	0.550	0.29	19.19	21.3	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	18	0.462	25.69	359.63	282.8	1.0	INSERT INTO posting_windows (id, region, local_date, starts_at, on_time_deadline, created_at, notification_ids, notified_user_ids) VALUES ($1, $2, $3, $4, $5, $…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 357f5242f3cecf617bc0cfc64f996e4d3bf04765
Type: cpu
Time: 2026-10-31 12:59:43 JST
Duration: 60s, Total samples = 36.32s (60.53%)
Showing nodes accounting for 25.47s, 70.13% of 36.32s total
Dropped 970 nodes (cum <= 0.18s)
Showing top 40 nodes out of 306
      flat  flat%   sum%        cum   cum%
     3.27s  9.00%  9.00%      3.27s  9.00%  internal/runtime/syscall/linux.Syscall6
     1.84s  5.07% 14.07%      1.84s  5.07%  runtime.memmove
     1.82s  5.01% 19.08%      6.34s 17.46%  runtime.scanObject
     1.74s  4.79% 23.87%      1.74s  4.79%  runtime.memclrNoHeapPointers
     1.70s  4.68% 28.55%      5.65s 15.56%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.63s  4.49% 33.04%      1.63s  4.49%  crypto/internal/fips140/sha256.blockAVX2
     1.55s  4.27% 37.31%      1.90s  5.23%  runtime.findObject
     1.34s  3.69% 41.00%      1.34s  3.69%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.32s  3.63% 44.63%      2.15s  5.92%  runtime.tryDeferToSpanScan
     1.24s  3.41% 48.05%      1.24s  3.41%  aeshashbody
     0.84s  2.31% 50.36%      0.84s  2.31%  example.com/isureal-company/isureal.(*UserRow).active
     0.66s  1.82% 52.18%      0.66s  1.82%  memeqbody
     0.48s  1.32% 53.50%      1.99s  5.48%  runtime.mapaccess1_faststr
     0.48s  1.32% 54.82%      0.48s  1.32%  runtime.spanClass.sizeclass (inline)
     0.42s  1.16% 55.97%      0.42s  1.16%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.39s  1.07% 57.05%      0.55s  1.51%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.37s  1.02% 58.07%      1.12s  3.08%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.33s  0.91% 58.98%      0.67s  1.84%  database/sql.convertAssignRows
     0.33s  0.91% 59.88%      0.56s  1.54%  example.com/isureal-company/isureal.postLess
     0.32s  0.88% 60.77%      0.32s  0.88%  indexbody
     0.25s  0.69% 61.45%      0.25s  0.69%  runtime.nextFreeFast (inline)
     0.23s  0.63% 62.09%      0.23s  0.63%  runtime.typePointers.nextFast (inline)
     0.22s  0.61% 62.69%      0.28s  0.77%  runtime.heapArenaOf (inline)
     0.21s  0.58% 63.27%      0.26s  0.72%  runtime.(*mspan).writeHeapBitsSmall
     0.21s  0.58% 63.85%      1.08s  2.97%  runtime.mallocgcSmallScanNoHeader
     0.20s  0.55% 64.40%      0.20s  0.55%  runtime.futex
     0.19s  0.52% 64.92%      0.19s  0.52%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.19s  0.52% 65.45%      0.43s  1.18%  runtime.greyobject
     0.18s   0.5% 65.94%      3.20s  8.81%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.17s  0.47% 66.41%      0.19s  0.52%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.17s  0.47% 66.88%      1.61s  4.43%  runtime.mapassign_faststr
     0.17s  0.47% 67.35%      0.20s  0.55%  runtime.typePointers.next
     0.16s  0.44% 67.79%      0.45s  1.24%  runtime.scanObjectsSmall
     0.15s  0.41% 68.20%      1.78s  4.90%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.13s  0.36% 68.56%      0.28s  0.77%  compress/flate.(*compressor).deflate
     0.13s  0.36% 68.92%      1.10s  3.03%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.33% 69.25%      4.32s 11.89%  runtime.mallocgc
     0.11s   0.3% 69.55%      2.17s  5.97%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.11s   0.3% 69.85%      1.84s  5.07%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.10s  0.28% 70.13%      0.19s  0.52%  crypto/internal/fips140/hmac.(*HMAC).Reset
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 44.1% / 100% | 26% | 0.1% | 6 | 145 MB | run queue が積もっている |
| isu2 (2 cpu) | 17.0% / 55% | 0% | 0.4% | 4 | 2506 MB |  |
| isu3 (2 cpu) | 15.1% / 59% | 0% | 0.5% | 5 | 2556 MB | run queue が積もっている |
| isu4 (2 cpu) | 50.7% / 100% | 22% | 16.3% | 39 | 109 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 66% / nginx 14% / journalctl 2% / systemd-journal 2% |
| isu2 | isureal 32% / nginx 1% |
| isu3 | isureal 28% / nginx 1% |
| isu4 | postgres 60% / isureal 5% |

## Diff from previous run

スコア: 64700 → 101157 (+36457)

`e0e4cb3..1deeabf`

```
 conf/isu1/etc/sysctl.d/99-tcp-slow-start.conf |   1 +
 conf/isu2/etc/sysctl.d/99-tcp-slow-start.conf |   1 +
 conf/isu3/etc/sysctl.d/99-tcp-slow-start.conf |   1 +
 conf/isu4/etc/sysctl.d/99-tcp-slow-start.conf |   1 +
 webapp/go/account_deletion_api.go             |   2 +-
 webapp/go/admin_user_api.go                   |   2 +-
 webapp/go/comment_api.go                      |   2 +-
 webapp/go/feed_api.go                         | 113 ++++++---
 webapp/go/feed_test.go                        | 331 ++++++++++++++++++++++++++
 webapp/go/offload.go                          |  27 ++-
 webapp/go/post_api.go                         |  62 ++++-
 webapp/go/realmoji_api.go                     |   2 +-
 webapp/go/relationship.go                     |  10 +-
 webapp/go/store.go                            | 113 +++++++--
 14 files changed, 605 insertions(+), 63 deletions(-)
```
