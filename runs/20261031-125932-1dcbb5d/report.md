# run 20261031-125932-1dcbb5d

## Summary

| | |
|---|---|
| score | 37886 |
| status | pass |
| ref | repl-replica |
| sha | 1dcbb5d020db |
| 所要時間 | 246s |
| 前回比 | -24357 (-39.1%) vs 20261031-125349-02ae5ac |
| ベスト比 | -26814 (-41.4%) vs 20261031-121235-e0e4cb3 = 64700 |
| note | 複製 第1段 shadow（europe1・america1 が裏でフィード等を計算して asia1 の応答と比べる） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	27	66	154	314	313
europe	7	18	25	26	26	25
america	6	15	20	26	26	25
```

```
{"id":"10763","status":"done","portalStatus":"finished","passed":true,"score":37886,"raw":37886,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:02:19.990Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:02:22.957Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:02:25.952Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:02:28.982Z level=INFO msg="asia 2 日目: 27 人"
time=2026-10-31T04:02:35.541Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T04:02:35.653Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T04:02:38.122Z level=INFO msg="asia 3 日目: 66 人"
time=2026-10-31T04:02:45.589Z level=INFO msg="europe 3 日目: 25 人"
time=2026-10-31T04:02:45.796Z level=INFO msg="america 3 日目: 20 人"
time=2026-10-31T04:02:46.498Z level=INFO msg="user moved" user=yohite5518 from=asia to=europe
time=2026-10-31T04:02:47.128Z level=INFO msg="asia 4 日目: 154 人"
time=2026-10-31T04:02:48.067Z level=INFO msg="user moved" user=moka5602 from=asia to=america
time=2026-10-31T04:02:49.474Z level=INFO msg="user moved" user=yomar5780 from=asia to=america
time=2026-10-31T04:02:50.164Z level=INFO msg="europe 4 日目: 26 人"
time=2026-10-31T04:02:57.691Z level=INFO msg="america 4 日目: 26 人"
time=2026-10-31T04:02:57.920Z level=INFO msg="asia 5 日目: 314 人"
time=2026-10-31T04:02:59.181Z level=INFO msg="europe 5 日目: 26 人"
time=2026-10-31T04:03:06.237Z level=INFO msg="america 5 日目: 26 人"
time=2026-10-31T04:03:15.553Z level=INFO msg="america 最終日（グランプリ）: 25 人"
time=2026-10-31T04:03:15.563Z level=INFO msg="europe 最終日（グランプリ）: 25 人"
time=2026-10-31T04:03:16.119Z level=INFO msg="asia 最終日（グランプリ）: 313 人"
time=2026-10-31T04:03:19.957Z level=INFO msg="グランプリ: アジア 115 人, ヨーロッパ 24 人, アメリカ 24 人"
{
  "pass": true,
  "score": 37886,
  "messages": [
    "shileomar-5085: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "ortane.5135: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "fumasol-5045: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.7%	576	2	0	GET	/api/notifications	5957.680	10.343	17.654	17.734	173
9.8%	789	12	0	POST	/api/posts	1310.990	1.662	6.800	7.280	637
6.5%	1679	43	0	POST	/api/friend-requests	862.403	0.514	5.578	6.911	113
5.6%	4740	240	0	GET	^/api/images/[^/]+	742.502	0.157	2.432	4.126	139407
5.5%	1948	551	0	POST	/api/auth/login	730.326	0.375	5.604	6.932	182
4.6%	325	156	0	PUT	/api/me/profile-image	608.344	1.872	12.262	14.770	8
3.2%	738	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	429.469	0.582	6.801	6.921	116
3.2%	333	1	0	POST	/api/auth/register	422.104	1.268	5.660	6.483	227
2.7%	3765	0	0	GET	^/api/feeds/[^/]+	365.082	0.097	2.352	3.402	2791
2.7%	1512	0	0	GET	/api/memories	361.386	0.239	2.998	3.498	8625
1.7%	1114	0	0	GET	/api/posting-window/current	226.272	0.203	3.060	3.329	237
1.4%	158	3	0	POST	^/api/posts/[^/]+/realmojis$	192.814	1.220	6.899	6.918	474
1.0%	211	0	0	PUT	^/api/follows/[^/]+$	139.440	0.661	6.832	6.909	0
1.0%	631	20	0	GET	^/api/posts/me/current$	138.466	0.219	3.289	3.329	626
0.9%	792	0	0	GET	/api/users	116.248	0.147	2.880	3.322	244
0.8%	146	0	0	POST	^/api/posts/[^/]+/comments$	112.313	0.769	5.906	5.954	333
0.8%	217	0	0	POST	/api/auth/logout	103.580	0.477	5.226	6.576	0
0.6%	169	0	0	PATCH	/api/me	82.223	0.487	5.495	6.920	276
0.5%	75	0	0	GET	/api/me/export	69.208	0.923	6.839	6.839	1954578
0.4%	583	8	0	GET	^/api/posts/[^/]+$	46.842	0.080	2.364	3.319	601
0.3%	687	0	0	GET	^/api/relationships/[^/]+	44.638	0.065	2.258	3.321	1393
0.3%	225	0	0	GET	^/api/memories/[^/]+	38.840	0.173	2.356	3.215	737
0.2%	18	0	0	POST	/api/admin/posting-window	32.647	1.814	4.743	4.743	177
0.2%	176	3	0	GET	/api/me	28.295	0.161	3.260	3.350	296
0.2%	256	0	0	GET	^/api/posts/[^/]+/realmojis$	28.188	0.110	2.364	3.308	396
0.2%	67	4	0	PATCH	^/api/posts/[^/]+$	25.010	0.373	4.524	4.524	615
0.2%	287	1	0	GET	^/api/users/[^/]+$	24.170	0.084	2.190	3.292	187
0.2%	202	0	0	GET	^/api/posts/[^/]+/comments$	24.077	0.119	2.371	2.449	555
0.1%	1	0	0	POST	/api/initialize	19.878	19.878	19.878	19.878	90
0.1%	18	0	0	DELETE	^/api/comments/[^/]+$	18.934	1.052	6.887	6.887	0
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
46.7%	109	1	0	GET	/api/notifications	987.428	9.059	16.157	16.157	178
11.3%	127	0	0	POST	/api/posts	238.602	1.879	6.983	7.018	619
7.3%	676	8	0	GET	^/api/images/[^/]+	154.328	0.228	2.529	2.658	178922
6.8%	495	0	0	GET	^/api/feeds/[^/]+	143.575	0.290	4.381	5.304	3792
4.2%	151	0	0	GET	/api/posting-window/current	88.867	0.589	3.373	3.409	239
4.2%	147	0	0	POST	/api/auth/login	88.516	0.602	5.443	6.699	251
3.8%	197	6	0	POST	/api/friend-requests	79.765	0.405	5.462	6.941	112
3.3%	123	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	69.275	0.563	4.524	4.619	115
2.9%	221	0	0	GET	/api/memories	62.336	0.282	3.203	3.538	11269
1.9%	87	0	0	GET	^/api/relationships/[^/]+	39.577	0.455	3.531	3.531	1755
0.9%	22	0	0	POST	/api/auth/logout	19.029	0.865	6.797	6.797	0
0.9%	71	2	0	GET	^/api/posts/[^/]+$	18.753	0.264	4.688	4.688	590
0.8%	125	4	0	GET	^/api/posts/me/current$	16.912	0.135	0.298	0.502	599
0.6%	36	0	0	GET	^/api/posts/[^/]+/comments$	12.837	0.357	4.336	4.336	610
0.6%	11	0	0	GET	/api/me/export	12.180	1.107	2.288	2.288	2549901
0.5%	10	0	0	DELETE	^/api/posts/[^/]+$	11.428	1.143	7.021	7.021	0
0.5%	32	0	0	POST	^/api/posts/[^/]+/comments$	11.337	0.354	4.428	4.428	291
0.5%	87	0	0	GET	/api/users	10.131	0.116	0.230	0.230	233
0.3%	23	0	0	GET	^/api/posts/[^/]+/realmojis$	6.844	0.298	3.940	3.940	278
0.3%	8	0	0	DELETE	^/api/comments/[^/]+$	6.288	0.786	4.853	4.853	0
0.3%	12	2	0	PATCH	^/api/posts/[^/]+$	5.634	0.470	3.349	3.349	543
0.3%	8	0	0	POST	^/api/posts/[^/]+/realmojis$	5.348	0.668	1.031	1.031	474
0.3%	23	0	0	POST	/api/auth/register	5.317	0.231	0.279	0.279	228
0.2%	24	0	0	GET	^/api/memories/[^/]+	4.654	0.194	0.984	0.984	1296
0.2%	12	2	0	GET	/api/me	3.435	0.286	2.002	2.002	231
0.2%	28	0	0	GET	^/api/users/[^/]+$	3.423	0.122	0.295	0.295	183
0.1%	13	0	0	PUT	^/api/follows/[^/]+$	2.613	0.201	0.247	0.247	0
0.1%	11	0	0	PATCH	/api/me	2.498	0.227	0.302	0.302	255
0.1%	7	0	0	PUT	/api/me/profile-image	2.448	0.350	0.405	0.405	0
0.1%	2	0	0	DELETE	^/api/realmojis/[^/]+$	1.281	0.640	0.711	0.711	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.8%	99	1	0	GET	/api/notifications	820.049	8.283	10.845	10.845	177
11.1%	346	0	0	GET	^/api/feeds/[^/]+	199.503	0.577	5.392	5.457	3130
8.3%	108	0	0	POST	/api/posts	148.691	1.377	7.301	8.092	653
6.9%	480	4	0	GET	^/api/images/[^/]+	123.925	0.258	1.320	2.869	224250
4.8%	119	0	0	POST	/api/auth/login	85.803	0.721	5.428	6.752	249
4.1%	132	3	0	POST	/api/friend-requests	74.255	0.563	5.198	6.921	113
3.6%	190	0	0	GET	/api/memories	64.032	0.337	3.294	3.366	8426
2.2%	48	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	38.594	0.804	4.771	4.771	116
1.8%	80	11	0	GET	^/api/posts/[^/]+$	33.047	0.413	5.347	5.347	521
1.6%	107	14	0	GET	^/api/posts/me/current$	29.392	0.275	3.393	3.400	565
1.6%	134	0	0	GET	/api/posting-window/current	28.456	0.212	2.323	3.121	240
0.9%	70	0	0	GET	/api/users	15.354	0.219	2.357	2.357	232
0.8%	23	0	0	POST	/api/auth/logout	14.980	0.651	4.560	4.560	0
0.7%	29	0	0	GET	^/api/users/[^/]+$	13.316	0.459	2.439	2.439	179
0.7%	24	0	0	GET	^/api/memories/[^/]+	12.838	0.535	3.227	3.227	587
0.6%	57	0	0	GET	^/api/relationships/[^/]+	11.485	0.201	2.542	2.542	1313
0.6%	7	0	0	GET	/api/me/export	10.621	1.517	3.617	3.617	2970584
0.5%	12	0	0	PUT	/api/me/profile-image	9.568	0.797	5.271	5.271	0
0.5%	23	0	0	POST	/api/auth/register	9.482	0.412	3.722	3.722	229
0.5%	5	0	0	DELETE	^/api/posts/[^/]+$	9.319	1.864	5.397	5.397	0
0.5%	11	0	0	POST	^/api/posts/[^/]+/comments$	9.169	0.834	4.617	4.617	348
0.3%	17	0	0	GET	^/api/posts/[^/]+/comments$	5.267	0.310	2.353	2.353	840
0.3%	8	0	0	POST	^/api/posts/[^/]+/realmojis$	5.039	0.630	1.574	1.574	540
0.3%	22	0	0	PUT	^/api/follows/[^/]+$	4.829	0.220	0.415	0.415	0
0.2%	18	0	0	GET	^/api/posts/[^/]+/realmojis$	3.602	0.200	0.654	0.654	304
0.2%	8	2	0	PATCH	^/api/posts/[^/]+$	2.830	0.354	1.002	1.002	484
0.2%	11	0	0	PATCH	/api/me	2.797	0.254	0.540	0.540	278
0.1%	16	1	0	GET	/api/me	1.989	0.124	0.322	0.322	276
0.1%	2	0	0	DELETE	^/api/comments/[^/]+$	1.772	0.886	1.493	1.493	0
0.0%	1	0	0	POST	/api/me/account-deletion	0.812	0.812	0.812	0.812	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 25024 クエリ / 80 種類 / 合計 190.194s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	23.6%	1374	44.897	32.68	518.39	640.0	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	19.2%	4	36.444	9111.00	14972.64	30382.0	1199753.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	18.0%	93	34.326	369.10	1342.55	6224.7	134.3	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
4	13.4%	4	25.456	6363.93	9293.34	22480.5	627151.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	6.5%	1006	12.387	12.31	564.44	121.1	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
6	5.9%	4	11.248	2812.11	3577.74	12035.3	529798.8	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
7	4.3%	7	8.101	1157.32	6903.54	219854.1	2603.9	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
8	3.9%	3056	7.416	2.43	32.39	11.6	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
9	1.1%	1014	2.129	2.10	21.50	28.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	1.1%	7	2.022	288.91	457.60	28651.4	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.4%	7	0.842	120.25	160.54	627.9	25388.4	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
12	0.4%	7	0.834	119.17	158.13	622.4	30081.4	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
13	0.4%	171	0.703	4.11	22.57	37.9	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
14	0.2%	1940	0.390	0.20	13.75	21.1	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
15	0.2%	189	0.360	1.91	6.52	25.3	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: a5456bc926e987692497ce8fc87a7552a97a9868
Type: cpu
Time: 2026-10-31 13:01:55 JST
Duration: 60s, Total samples = 32.49s (54.15%)
Showing nodes accounting for 24.43s, 75.19% of 32.49s total
Dropped 849 nodes (cum <= 0.16s)
Showing top 40 nodes out of 255
      flat  flat%   sum%        cum   cum%
     3.99s 12.28% 12.28%      3.99s 12.28%  hash/fnv.(*sum64a).Write (inline)
     2.44s  7.51% 19.79%      8.46s 26.04%  example.com/isureal-company/isureal.(*Store).digestLocked
     1.76s  5.42% 25.21%      5.70s 17.54%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.62s  4.99% 30.19%      1.62s  4.99%  internal/runtime/syscall/linux.Syscall6
     1.40s  4.31% 34.50%      1.40s  4.31%  runtime.memclrNoHeapPointers
     1.30s  4.00% 38.50%      1.30s  4.00%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.16s  3.57% 42.07%      4.13s 12.71%  runtime.scanObject
     1.10s  3.39% 45.46%      1.10s  3.39%  aeshashbody
     1.07s  3.29% 48.75%      1.07s  3.29%  runtime.memmove
     0.97s  2.99% 51.74%      1.38s  4.25%  runtime.findObject
     0.71s  2.19% 53.92%      1.23s  3.79%  runtime.tryDeferToSpanScan
     0.59s  1.82% 55.74%      0.59s  1.82%  crypto/internal/fips140/sha256.blockAVX2
     0.51s  1.57% 57.31%      1.81s  5.57%  runtime.mapaccess1_faststr
     0.44s  1.35% 58.66%      0.44s  1.35%  memeqbody
     0.43s  1.32% 59.99%      0.46s  1.42%  time.appendInt
     0.38s  1.17% 61.16%      0.38s  1.17%  runtime.spanClass.sizeclass (inline)
     0.37s  1.14% 62.30%      0.37s  1.14%  example.com/isureal-company/isureal.(*UserRow).active
     0.36s  1.11% 63.40%      0.53s  1.63%  example.com/isureal-company/isureal.postLess
     0.32s  0.98% 64.39%      0.32s  0.98%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.26s   0.8% 65.19%      0.41s  1.26%  database/sql.convertAssignRows
     0.26s   0.8% 65.99%      0.28s  0.86%  runtime.(*mspan).writeHeapBitsSmall
     0.24s  0.74% 66.73%      1.72s  5.29%  runtime.mapassign_faststr
     0.23s  0.71% 67.44%      0.23s  0.71%  runtime.nextFreeFast (inline)
     0.23s  0.71% 68.14%      0.23s  0.71%  sync/atomic.(*Int32).Add
     0.21s  0.65% 68.79%      0.41s  1.26%  internal/runtime/maps.(*Iter).Next
     0.20s  0.62% 69.41%      2.58s  7.94%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.18s  0.55% 69.96%      0.19s  0.58%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.17s  0.52% 70.48%      1.39s  4.28%  example.com/isureal-company/isureal.tstr
     0.17s  0.52% 71.01%      1.01s  3.11%  runtime.mallocgcSmallScanNoHeader
     0.16s  0.49% 71.50%      0.35s  1.08%  runtime.greyobject
     0.15s  0.46% 71.96%      3.27s 10.06%  runtime.mallocgc
     0.14s  0.43% 72.39%      4.13s 12.71%  example.com/isureal-company/isureal.h64
     0.14s  0.43% 72.82%      0.60s  1.85%  runtime.mallocgcSmallNoscan
     0.13s   0.4% 73.22%      0.17s  0.52%  runtime.spanOf (inline)
     0.12s  0.37% 73.59%      1.12s  3.45%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.37% 73.96%      0.51s  1.57%  runtime.mapaccess2_faststr
     0.12s  0.37% 74.33%      0.33s  1.02%  runtime.scanObjectsSmall
     0.10s  0.31% 74.64%      0.21s  0.65%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.09s  0.28% 74.92%      0.33s  1.02%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.09s  0.28% 75.19%      0.84s  2.59%  time.Time.appendFormatRFC3339
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 46.2% / 100% | 21% | 6.6% | 6 | 109 MB | run queue が積もっている |
| isu2 (2 cpu) | 23.2% / 79% | 0% | 0.2% | 5 | 1047 MB | run queue が積もっている |
| isu3 (2 cpu) | 28.2% / 90% | 1% | 0.2% | 3 | 612 MB |  |
| isu4 (2 cpu) | 37.0% / 100% | 5% | 14.9% | 23 | 105 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 62% / nginx 11% |
| isu2 | isureal 45% |
| isu3 | isureal 55% |
| isu4 | postgres 38% / isureal 1% |

## Diff from previous run

スコア: 62243 → 37886 (-24357)

`02ae5ac..1dcbb5d`

```
 conf/isu1/home/isucon/env.sh |   1 +
 webapp/go/offload.go         |  27 +++++++--
 webapp/go/post_api.go        |  62 ++++++++++++++++----
 webapp/go/repl.go            |   4 +-
 webapp/go/repl_digest.go     | 137 +++++++++++++++++++++++++++++++++++++++++++
 webapp/go/server.go          |   1 +
 6 files changed, 217 insertions(+), 15 deletions(-)
```
