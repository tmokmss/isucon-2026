# run 20261031-122906-3a360a3

## Summary

| | |
|---|---|
| score | 82657 |
| status | pass |
| ref | grandprix-jpeg-slots32 |
| sha | 3a360a37a65a |
| 所要時間 | 276s |
| 前回比 | +17957 (27.8%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +17957 (27.8%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | jpeg fifo + 32 slots per peer |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	141	394	760	760
europe	7	17	43	105	106	106
america	6	12	30	51	52	52
```

```
{"id":"8486","status":"done","portalStatus":"finished","passed":true,"score":82657,"raw":82657,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:32:14.050Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:32:17.016Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:32:20.016Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:32:23.046Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T03:32:26.055Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T03:32:29.017Z level=INFO msg="america 2 日目: 12 人"
time=2026-10-31T03:32:30.318Z level=INFO msg="user moved" user=elanse_5413 from=europe to=america
time=2026-10-31T03:32:32.038Z level=INFO msg="asia 3 日目: 141 人"
time=2026-10-31T03:32:35.030Z level=INFO msg="europe 3 日目: 43 人"
time=2026-10-31T03:32:35.596Z level=INFO msg="user moved" user=keru-5996 from=asia to=america
time=2026-10-31T03:32:35.728Z level=INFO msg="user moved" user=kael6013 from=asia to=europe
time=2026-10-31T03:32:38.039Z level=INFO msg="america 3 日目: 30 人"
time=2026-10-31T03:32:41.105Z level=INFO msg="asia 4 日目: 394 人"
time=2026-10-31T03:32:41.359Z level=INFO msg="user moved" user=elanse_5413 from=america to=asia
time=2026-10-31T03:32:43.171Z level=INFO msg="user moved" user=elanse_5413 from=asia to=america
time=2026-10-31T03:32:44.256Z level=INFO msg="europe 4 日目: 105 人"
time=2026-10-31T03:32:47.358Z level=INFO msg="america 4 日目: 51 人"
time=2026-10-31T03:32:47.981Z level=INFO msg="user moved" user=shifufu.6433 from=asia to=europe
time=2026-10-31T03:32:48.251Z level=INFO msg="user moved" user=rore.7477 from=asia to=america
time=2026-10-31T03:32:50.188Z level=INFO msg="asia 5 日目: 760 人"
time=2026-10-31T03:32:53.061Z level=INFO msg="europe 5 日目: 106 人"
time=2026-10-31T03:32:56.102Z level=INFO msg="america 5 日目: 52 人"
time=2026-10-31T03:33:05.500Z level=INFO msg="america 最終日（グランプリ）: 52 人"
time=2026-10-31T03:33:05.623Z level=INFO msg="europe 最終日（グランプリ）: 106 人"
time=2026-10-31T03:33:05.964Z level=INFO msg="asia 最終日（グランプリ）: 760 人"
time=2026-10-31T03:33:14.038Z level=INFO msg="グランプリ: アジア 191 人, ヨーロッパ 15 人, アメリカ 7 人"
{
  "pass": true,
  "score": 82657,
  "messages": [
    "torben-9889: プロフィールを変えられません: PATCH /api/me: 応答がタイムアウトしました",
    "ulsonoa-7441: 投稿 5a23ec9b-ec26-40a7-ac9e-3341d462ffc7 に RealMoji を送れません: POST /api/posts/5a23ec9b-ec26-40a7-ac9e-3341d462ffc7/realmojis: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.4%	1302	0	0	GET	/api/notifications	11607.842	8.915	15.939	15.961	165
17.3%	1780	56	0	POST	/api/posts	6025.281	3.385	11.866	14.777	628
13.1%	4022	93	0	POST	/api/friend-requests	4564.471	1.135	10.889	18.864	113
8.9%	11520	931	0	GET	^/api/images/[^/]+	3079.137	0.267	4.702	12.725	127919
6.7%	5448	2008	0	POST	/api/auth/login	2343.263	0.430	5.716	19.116	164
5.2%	1316	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1792.507	1.362	10.393	17.030	116
1.9%	780	1	0	POST	/api/auth/register	668.593	0.857	4.898	6.396	227
1.8%	539	132	0	PUT	/api/me/profile-image	622.108	1.154	5.354	10.369	5
1.6%	339	19	0	POST	^/api/posts/[^/]+/realmojis$	556.150	1.641	11.740	16.403	440
1.6%	480	29	0	POST	^/api/posts/[^/]+/comments$	545.212	1.136	10.412	13.876	333
1.3%	479	0	0	PUT	^/api/follows/[^/]+$	447.076	0.933	6.165	9.074	0
1.2%	370	0	0	PATCH	/api/me	400.066	1.081	10.657	18.551	280
1.0%	2860	0	0	GET	/api/memories	356.139	0.125	0.993	1.786	10902
0.9%	8761	1	0	GET	^/api/feeds/[^/]+	311.287	0.036	0.317	1.494	2125
0.6%	154	7	0	PATCH	^/api/posts/[^/]+$	224.668	1.459	8.058	9.794	620
0.6%	2543	0	0	GET	/api/posting-window/current	218.628	0.086	0.852	1.496	236
0.6%	569	0	0	POST	/api/auth/logout	206.368	0.363	4.444	10.229	0
0.5%	189	0	0	GET	/api/me/export	159.370	0.843	6.822	8.512	1832258
0.3%	76	0	0	DELETE	^/api/comments/[^/]+$	90.196	1.187	13.085	13.085	0
0.2%	57	0	0	DELETE	^/api/posts/[^/]+$	82.931	1.455	5.635	5.635	0
0.2%	1353	71	0	GET	^/api/posts/me/current$	74.768	0.055	0.542	1.795	615
0.2%	1971	0	0	GET	/api/users	73.151	0.037	0.234	1.462	239
0.2%	43	0	0	DELETE	^/api/realmojis/[^/]+$	57.741	1.343	6.461	6.461	0
0.1%	1333	15	0	GET	^/api/posts/[^/]+$	36.551	0.027	0.221	0.862	622
0.1%	1207	0	0	GET	^/api/relationships/[^/]+	34.345	0.028	0.224	0.771	920
0.1%	504	0	0	GET	^/api/memories/[^/]+	28.980	0.057	0.381	1.400	694
0.1%	31	0	0	PUT	^/api/blocks/[^/]+$	23.579	0.761	6.544	6.544	0
0.1%	674	0	0	GET	^/api/users/[^/]+$	19.388	0.029	0.217	0.608	201
0.1%	1	0	0	POST	/api/me/account-deletion	19.326	19.326	19.326	19.326	0
0.0%	568	0	0	GET	^/api/posts/[^/]+/realmojis$	12.803	0.023	0.151	0.312	653
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.6%	254	1	0	GET	/api/notifications	2163.125	8.516	12.272	12.594	173
18.2%	308	4	0	POST	/api/posts	1102.745	3.580	11.431	18.531	648
11.7%	624	14	0	POST	/api/friend-requests	713.000	1.143	9.280	12.274	113
6.3%	399	0	0	POST	/api/auth/login	381.142	0.955	9.629	13.023	246
5.2%	1493	14	0	GET	^/api/images/[^/]+	317.337	0.213	1.400	1.729	173888
4.5%	214	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	274.234	1.281	10.779	12.543	116
3.7%	1287	0	0	GET	^/api/feeds/[^/]+	226.837	0.176	0.685	1.160	2624
1.7%	428	0	0	GET	/api/posting-window/current	105.764	0.247	0.907	9.161	239
1.7%	404	0	0	GET	/api/memories	102.995	0.255	1.062	2.127	12441
1.1%	316	0	0	GET	/api/users	65.298	0.207	0.807	0.916	249
1.0%	103	0	0	POST	/api/auth/register	59.325	0.576	2.171	2.308	228
1.0%	30	0	0	GET	/api/me/export	58.882	1.963	4.972	4.972	2340138
1.0%	64	0	0	POST	/api/auth/logout	58.512	0.914	6.307	6.307	0
0.9%	32	6	0	POST	^/api/posts/[^/]+/realmojis$	57.072	1.783	9.216	9.216	423
0.8%	64	0	0	PUT	/api/me/profile-image	50.704	0.792	3.216	3.216	0
0.8%	243	12	0	GET	^/api/posts/me/current$	48.682	0.200	0.642	0.737	627
0.6%	60	0	0	PUT	^/api/follows/[^/]+$	35.933	0.599	4.495	4.495	0
0.6%	209	6	0	GET	^/api/posts/[^/]+$	34.830	0.167	0.514	0.924	597
0.5%	46	0	0	PATCH	/api/me	29.794	0.648	4.800	4.800	291
0.5%	183	0	0	GET	^/api/relationships/[^/]+	28.791	0.157	0.909	0.931	1162
0.5%	26	1	0	PATCH	^/api/posts/[^/]+$	28.749	1.106	4.997	4.997	642
0.5%	34	1	0	POST	^/api/posts/[^/]+/comments$	28.501	0.838	8.098	8.098	319
0.3%	95	0	0	GET	^/api/users/[^/]+$	16.434	0.173	0.489	0.489	188
0.2%	75	1	0	GET	^/api/posts/[^/]+/realmojis$	14.852	0.198	2.135	2.135	687
0.2%	15	0	0	DELETE	^/api/posts/[^/]+$	14.710	0.981	2.705	2.705	0
0.2%	48	0	0	GET	^/api/memories/[^/]+	11.892	0.248	1.354	1.354	1054
0.2%	71	2	0	GET	/api/me	11.860	0.167	0.542	0.542	290
0.2%	77	0	0	GET	^/api/posts/[^/]+/comments$	11.263	0.146	0.381	0.381	1274
0.1%	1	0	0	POST	/api/me/account-deletion	7.418	7.418	7.418	7.418	0
0.1%	2	0	0	DELETE	/api/me/profile-image	5.719	2.859	4.542	4.542	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.6%	142	0	0	GET	/api/notifications	1121.957	7.901	9.459	13.276	174
18.2%	152	2	0	POST	/api/posts	556.362	3.660	11.023	11.139	630
8.2%	241	6	0	POST	/api/friend-requests	250.489	1.039	10.128	14.799	113
6.7%	223	0	0	POST	/api/auth/login	205.789	0.923	10.744	11.076	251
5.6%	121	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	170.874	1.412	12.289	16.687	115
5.4%	834	5	0	GET	^/api/images/[^/]+	166.596	0.200	1.067	1.423	143088
3.0%	544	0	0	GET	^/api/feeds/[^/]+	92.414	0.170	0.546	1.470	2491
2.1%	230	0	0	GET	/api/memories	63.865	0.278	1.168	1.517	16225
1.5%	208	0	0	GET	/api/posting-window/current	45.790	0.220	0.826	0.964	240
1.3%	46	3	0	POST	^/api/posts/[^/]+/comments$	40.978	0.891	6.430	6.430	271
1.2%	23	3	0	POST	^/api/posts/[^/]+/realmojis$	37.237	1.619	7.213	7.213	449
1.1%	18	3	0	PATCH	^/api/posts/[^/]+$	33.295	1.850	8.299	8.299	535
1.1%	21	0	0	PUT	/api/me/profile-image	32.724	1.558	5.733	5.733	0
1.0%	34	0	0	POST	/api/auth/logout	29.991	0.882	6.205	6.205	0
0.9%	24	0	0	PATCH	/api/me	27.953	1.165	4.608	4.608	270
0.9%	39	0	0	PUT	^/api/follows/[^/]+$	27.765	0.712	4.105	4.105	0
0.8%	119	13	0	GET	^/api/posts/me/current$	25.577	0.215	0.758	0.784	569
0.8%	48	0	0	POST	/api/auth/register	23.148	0.482	1.914	1.914	229
0.7%	13	0	0	GET	/api/me/export	22.673	1.744	6.810	6.810	2261644
0.6%	110	0	0	GET	/api/users	17.427	0.158	0.494	0.640	229
0.4%	89	0	0	GET	^/api/relationships/[^/]+	13.613	0.153	0.558	0.558	1593
0.4%	79	5	0	GET	^/api/posts/[^/]+$	12.062	0.153	0.476	0.476	574
0.2%	1	0	0	POST	/api/me/account-deletion	7.379	7.379	7.379	7.379	0
0.2%	10	0	0	DELETE	^/api/comments/[^/]+$	6.677	0.668	4.260	4.260	0
0.2%	24	0	0	GET	^/api/memories/[^/]+	5.037	0.210	0.764	0.764	885
0.2%	35	0	0	GET	^/api/posts/[^/]+/realmojis$	4.882	0.139	0.232	0.232	749
0.1%	24	0	0	GET	/api/me	4.395	0.183	0.341	0.341	293
0.1%	31	0	0	GET	^/api/users/[^/]+$	4.054	0.131	0.262	0.262	176
0.1%	33	0	0	GET	^/api/posts/[^/]+/comments$	4.039	0.122	0.176	0.176	2170
0.1%	5	0	0	DELETE	^/api/posts/[^/]+$	4.029	0.806	1.748	1.748	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 54929 クエリ / 81 種類 / 合計 303.053s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	60.3%	3036	182.873	60.23	479.60	644.4	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	23.2%	232	70.309	303.06	1157.28	6443.1	122.8	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.6%	1981	17.110	8.64	461.35	117.2	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	5.5%	7414	16.622	2.24	75.78	10.7	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.6%	2178	4.930	2.26	41.50	25.0	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.7%	1	2.200	2200.37	2200.37	29933.0	1199358.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.5%	1	1.558	1558.02	1558.02	21300.0	625278.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.4%	366	1.334	3.64	14.96	30.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.3%	4756	0.908	0.19	12.02	20.0	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	0.3%	527	0.830	1.58	7.91	21.7	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.2%	1	0.739	738.71	738.71	11620.0	529643.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
12	0.2%	4062	0.493	0.12	10.08	11.4	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
13	0.1%	1604	0.448	0.28	8.05	21.2	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
14	0.1%	2	0.313	156.63	157.73	211898.5	2492.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
15	0.1%	2	0.264	131.95	255.63	16770.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: 9fcd210f6f751fe5c5b6b6c8dd6615a669222d9e
Type: cpu
Time: 2026-10-31 12:31:57 JST
Duration: 60.01s, Total samples = 50.82s (84.69%)
Showing nodes accounting for 36.87s, 72.55% of 50.82s total
Dropped 1090 nodes (cum <= 0.25s)
Showing top 40 nodes out of 303
      flat  flat%   sum%        cum   cum%
     4.34s  8.54%  8.54%      4.34s  8.54%  internal/runtime/syscall/linux.Syscall6
     3.40s  6.69% 15.23%     11.65s 22.92%  runtime.scanObject
     3.26s  6.41% 21.65%      3.26s  6.41%  runtime.memclrNoHeapPointers
     2.67s  5.25% 26.90%      3.43s  6.75%  runtime.findObject
     2.41s  4.74% 31.64%      2.41s  4.74%  runtime.memmove
     2.34s  4.60% 36.25%      3.79s  7.46%  runtime.tryDeferToSpanScan
     1.61s  3.17% 39.41%      1.61s  3.17%  crypto/internal/fips140/sha256.blockAVX2
     1.55s  3.05% 42.46%      5.60s 11.02%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.47s  2.89% 45.36%      1.47s  2.89%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.38s  2.72% 48.07%      1.38s  2.72%  example.com/isureal-company/isureal.(*UserRow).active
     1.26s  2.48% 50.55%      1.26s  2.48%  aeshashbody
     0.99s  1.95% 52.50%      0.99s  1.95%  runtime.spanClass.sizeclass (inline)
     0.93s  1.83% 54.33%      2.81s  5.53%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.84s  1.65% 55.98%      1.23s  2.42%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.67s  1.32% 57.30%      0.67s  1.32%  memeqbody
     0.62s  1.22% 58.52%      0.65s  1.28%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.55s  1.08% 59.60%      2.19s  4.31%  runtime.mapaccess1_faststr
     0.47s  0.92% 60.53%      1.07s  2.11%  runtime.greyobject
     0.44s  0.87% 61.39%      0.64s  1.26%  example.com/isureal-company/isureal.postLess
     0.42s  0.83% 62.22%      0.48s  0.94%  runtime.typePointers.next
     0.39s  0.77% 62.99%      0.39s  0.77%  indexbody
     0.38s  0.75% 63.73%      0.38s  0.75%  crypto/internal/fips140/aes/gcm.gcmAesEnc
     0.37s  0.73% 64.46%      0.49s  0.96%  runtime.(*mspan).markBitsForIndex (inline)
     0.34s  0.67% 65.13%      0.34s  0.67%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.33s  0.65% 65.78%      0.33s  0.65%  runtime.typePointers.nextFast (inline)
     0.31s  0.61% 66.39%      0.46s  0.91%  runtime.scanObjectSmall
     0.29s  0.57% 66.96%      0.29s  0.57%  runtime.(*mspan).divideByElemSize (inline)
     0.29s  0.57% 67.53%      0.29s  0.57%  sync/atomic.(*Int32).Add
     0.27s  0.53% 68.06%      0.42s  0.83%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.27s  0.53% 68.60%      1.95s  3.84%  runtime.mapassign_faststr
     0.26s  0.51% 69.11%      0.26s  0.51%  runtime.nextFreeFast (inline)
     0.25s  0.49% 69.60%      0.70s  1.38%  database/sql.convertAssignRows
     0.25s  0.49% 70.09%      0.84s  1.65%  runtime.scanObjectsSmall
     0.22s  0.43% 70.52%      0.33s  0.65%  runtime.heapArenaOf (inline)
     0.21s  0.41% 70.94%      0.27s  0.53%  runtime.(*mspan).writeHeapBitsSmall
     0.19s  0.37% 71.31%      1.93s  3.80%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.17s  0.33% 71.65%      1.21s  2.38%  runtime.mallocgcSmallScanNoHeader
     0.17s  0.33% 71.98%      1.52s  2.99%  runtime.scanSpan
     0.15s   0.3% 72.27%      2.28s  4.49%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.14s  0.28% 72.55%      1.82s  3.58%  crypto/internal/fips140/sha256.(*Digest).Write
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 54.7% / 100% | 37% | 0.2% | 9 | 181 MB | CPU が飽和 |
| isu2 (2 cpu) | 17.0% / 56% | 0% | 0.5% | 4 | 2657 MB |  |
| isu3 (2 cpu) | 14.7% / 63% | 0% | 0.5% | 4 | 2774 MB |  |
| isu4 (2 cpu) | 45.6% / 99% | 15% | 15.1% | 27 | 117 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 77% / nginx 15% / journalctl 11% / systemd-journal 2% |
| isu2 | isureal 32% / nginx 1% |
| isu3 | isureal 27% |
| isu4 | postgres 54% / isureal 3% |

## Diff from previous run

スコア: 64700 → 82657 (+17957)

`e0e4cb3..3a360a3`

```
 webapp/go/account_deletion_api.go |   2 +-
 webapp/go/admin_user_api.go       |   2 +-
 webapp/go/comment_api.go          |   2 +-
 webapp/go/offload.go              |  27 ++++++++--
 webapp/go/realmoji_api.go         |   2 +-
 webapp/go/relationship.go         |  10 +++-
 webapp/go/store.go                | 100 ++++++++++++++++++++++++++++++++------
 7 files changed, 119 insertions(+), 26 deletions(-)
```
