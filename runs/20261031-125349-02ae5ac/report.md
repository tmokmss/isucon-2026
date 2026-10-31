# run 20261031-125349-02ae5ac

## Summary

| | |
|---|---|
| score | 62243 |
| status | pass |
| ref | repl-replica |
| sha | 02ae5ac4acd0 |
| 所要時間 | 208s |
| 前回比 | -1068 (-1.7%) vs 20261031-124932-86828dc |
| ベスト比 | -2457 (-3.8%) vs 20261031-121235-e0e4cb3 = 64700 |
| note | 複製 第1段 shadow（europe1・america1 が裏でフィード等を計算して asia1 の応答と比べる） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	132	370	656	656
europe	7	16	40	98	98	98
america	6	13	30	52	52	52
```

```
{"id":"10279","status":"done","portalStatus":"finished","passed":true,"score":62243,"raw":62243,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:55:45.536Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:55:48.497Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:55:51.596Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:55:54.644Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T03:55:57.617Z level=INFO msg="europe 2 日目: 16 人"
time=2026-10-31T03:56:00.598Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T03:56:03.675Z level=INFO msg="asia 3 日目: 132 人"
time=2026-10-31T03:56:06.629Z level=INFO msg="europe 3 日目: 40 人"
time=2026-10-31T03:56:09.491Z level=INFO msg="user moved" user=nete.5120 from=america to=europe
time=2026-10-31T03:56:09.596Z level=INFO msg="america 3 日目: 30 人"
time=2026-10-31T03:56:12.690Z level=INFO msg="asia 4 日目: 370 人"
time=2026-10-31T03:56:15.773Z level=INFO msg="europe 4 日目: 98 人"
time=2026-10-31T03:56:18.764Z level=INFO msg="america 4 日目: 52 人"
time=2026-10-31T03:56:21.801Z level=INFO msg="asia 5 日目: 656 人"
time=2026-10-31T03:56:25.981Z level=INFO msg="europe 5 日目: 98 人"
time=2026-10-31T03:56:27.661Z level=INFO msg="america 5 日目: 52 人"
time=2026-10-31T03:56:36.951Z level=INFO msg="asia 最終日（グランプリ）: 656 人"
time=2026-10-31T03:56:37.101Z level=INFO msg="europe 最終日（グランプリ）: 98 人"
time=2026-10-31T03:56:37.506Z level=INFO msg="america 最終日（グランプリ）: 52 人"
time=2026-10-31T03:56:45.519Z level=INFO msg="グランプリ: アジア 12 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 62243,
  "messages": [
    "sonoa_5397: フィード friends を取得できません: GET /api/feeds/friends: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
28.9%	1107	2	0	GET	/api/notifications	9368.013	8.463	15.189	15.311	171
18.9%	1426	73	0	POST	/api/posts	6122.943	4.294	14.869	15.569	618
12.5%	3248	55	0	POST	/api/friend-requests	4059.128	1.250	14.056	18.263	114
9.3%	7643	815	0	GET	^/api/images/[^/]+	2994.563	0.392	9.375	17.859	145341
7.9%	3334	1073	0	POST	/api/auth/login	2568.568	0.770	12.387	16.249	173
4.9%	1079	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1581.626	1.466	15.406	19.634	116
3.9%	672	1	0	POST	/api/auth/register	1249.253	1.859	9.588	11.257	227
2.1%	2042	0	0	GET	/api/posting-window/current	670.983	0.329	4.908	7.303	237
1.9%	487	120	0	PUT	/api/me/profile-image	605.935	1.244	7.761	9.183	5
1.6%	353	15	0	POST	^/api/posts/[^/]+/realmojis$	529.661	1.500	15.696	18.744	473
1.4%	416	0	0	PUT	^/api/follows/[^/]+$	450.908	1.084	13.646	16.631	0
1.1%	348	14	0	POST	^/api/posts/[^/]+/comments$	368.533	1.059	6.290	13.369	337
0.9%	301	0	0	PATCH	/api/me	292.771	0.973	7.741	14.909	278
0.6%	6474	0	0	GET	^/api/feeds/[^/]+	187.995	0.029	0.194	4.923	2163
0.6%	1966	0	0	GET	/api/memories	185.786	0.094	0.881	2.986	9924
0.5%	323	0	0	POST	/api/auth/logout	175.718	0.544	6.492	15.250	0
0.4%	118	5	0	PATCH	^/api/posts/[^/]+$	140.050	1.187	10.952	15.946	620
0.4%	54	0	0	DELETE	^/api/realmojis/[^/]+$	131.963	2.444	15.705	15.705	0
0.4%	123	0	0	GET	/api/me/export	125.990	1.024	6.293	12.907	2102980
0.4%	52	0	0	DELETE	^/api/posts/[^/]+$	116.781	2.246	13.747	13.747	0
0.3%	69	0	0	DELETE	^/api/comments/[^/]+$	102.202	1.481	18.053	18.053	0
0.2%	1508	0	0	GET	/api/users	79.222	0.053	0.712	4.927	246
0.2%	1025	78	0	GET	^/api/posts/me/current$	69.659	0.068	0.968	14.242	601
0.1%	967	0	0	GET	^/api/relationships/[^/]+	25.016	0.026	0.195	3.071	994
0.1%	1070	17	0	GET	^/api/posts/[^/]+$	24.102	0.023	0.143	0.191	611
0.1%	301	0	0	GET	^/api/memories/[^/]+	21.098	0.070	0.456	1.771	678
0.1%	1	0	0	POST	/api/initialize	19.388	19.388	19.388	19.388	90
0.1%	22	0	0	PUT	^/api/blocks/[^/]+$	17.336	0.788	4.598	4.598	0
0.1%	503	0	0	GET	^/api/users/[^/]+$	16.860	0.034	0.183	2.567	192
0.0%	374	2	0	GET	/api/me	12.381	0.033	0.167	0.742	300
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.2%	220	0	0	GET	/api/notifications	1858.553	8.448	11.405	15.665	177
20.3%	266	4	0	POST	/api/posts	1168.896	4.394	16.132	18.153	640
12.8%	510	5	0	POST	/api/friend-requests	736.676	1.444	12.798	16.618	114
7.0%	328	0	0	POST	/api/auth/login	402.270	1.226	14.599	15.095	246
4.9%	1207	11	0	GET	^/api/images/[^/]+	280.194	0.232	2.094	2.553	165951
3.6%	847	0	0	GET	^/api/feeds/[^/]+	208.196	0.246	0.622	39.396	2284
3.6%	167	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	205.533	1.231	14.195	15.686	116
2.4%	362	0	0	GET	/api/posting-window/current	135.919	0.375	3.156	7.411	239
1.4%	54	0	0	POST	/api/auth/logout	81.112	1.502	13.103	13.103	0
1.3%	329	0	0	GET	/api/memories	77.187	0.235	1.038	1.888	11221
1.3%	41	3	0	POST	^/api/posts/[^/]+/realmojis$	75.777	1.848	13.206	13.206	455
1.2%	53	0	0	POST	^/api/posts/[^/]+/comments$	69.613	1.313	15.012	15.012	316
1.0%	58	0	0	PUT	/api/me/profile-image	59.985	1.034	5.680	5.680	0
0.9%	95	0	0	POST	/api/auth/register	53.626	0.564	2.558	2.558	227
0.8%	58	0	0	PUT	^/api/follows/[^/]+$	48.346	0.834	5.743	5.743	0
0.8%	232	0	0	GET	/api/users	47.085	0.203	1.036	2.367	243
0.8%	197	16	0	GET	^/api/posts/me/current$	44.465	0.226	0.711	5.567	597
0.6%	20	1	0	PATCH	^/api/posts/[^/]+$	33.581	1.679	14.595	14.595	601
0.5%	38	0	0	PATCH	/api/me	31.009	0.816	5.209	5.209	283
0.4%	135	0	0	GET	^/api/relationships/[^/]+	22.999	0.170	0.503	0.763	1320
0.4%	127	4	0	GET	^/api/posts/[^/]+$	22.692	0.179	0.478	0.667	597
0.2%	1	0	0	POST	/api/me/account-deletion	14.268	14.268	14.268	14.268	0
0.2%	3	0	0	DELETE	^/api/realmojis/[^/]+$	13.132	4.377	7.675	7.675	0
0.2%	9	0	0	GET	/api/me/export	11.985	1.332	3.807	3.807	1887961
0.2%	71	0	0	GET	^/api/users/[^/]+$	11.328	0.160	0.414	0.414	192
0.2%	63	1	0	GET	/api/me	10.602	0.168	0.392	0.392	293
0.2%	11	0	0	DELETE	^/api/posts/[^/]+$	10.524	0.957	1.892	1.892	0
0.2%	42	0	0	GET	^/api/memories/[^/]+	8.685	0.207	0.460	0.460	1096
0.1%	41	0	0	GET	^/api/posts/[^/]+/comments$	8.330	0.203	1.314	1.314	909
0.1%	33	0	0	GET	^/api/posts/[^/]+/realmojis$	5.599	0.170	0.420	0.420	650
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
31.9%	131	0	0	GET	/api/notifications	996.067	7.604	9.737	9.739	181
18.2%	148	1	0	POST	/api/posts	567.773	3.836	16.025	16.049	648
13.9%	320	7	0	POST	/api/friend-requests	434.325	1.357	14.248	15.972	113
8.2%	197	0	0	POST	/api/auth/login	254.780	1.293	14.332	14.686	248
6.2%	109	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	194.855	1.788	14.836	16.611	116
4.3%	771	4	0	GET	^/api/images/[^/]+	134.167	0.174	0.989	1.428	174777
3.4%	624	0	0	GET	^/api/feeds/[^/]+	107.459	0.172	0.658	1.846	2808
1.7%	192	0	0	GET	/api/posting-window/current	53.505	0.279	2.799	3.235	240
1.6%	30	0	0	PUT	/api/me/profile-image	49.208	1.640	6.699	6.699	0
1.3%	213	0	0	GET	/api/memories	41.257	0.194	0.470	0.591	14537
1.3%	52	0	0	POST	/api/auth/register	39.848	0.766	7.054	7.054	229
0.8%	21	0	0	PATCH	/api/me	26.448	1.259	8.025	8.025	270
0.8%	165	0	0	GET	/api/users	25.736	0.156	0.740	0.756	249
0.7%	18	2	0	POST	^/api/posts/[^/]+/realmojis$	22.509	1.250	6.742	6.742	387
0.7%	23	0	0	POST	/api/auth/logout	20.426	0.888	7.998	7.998	0
0.6%	122	12	0	GET	^/api/posts/[^/]+$	19.025	0.156	0.786	0.865	551
0.6%	118	8	0	GET	^/api/posts/me/current$	18.923	0.160	0.467	0.738	610
0.6%	5	0	0	DELETE	^/api/posts/[^/]+$	18.453	3.691	9.700	9.700	0
0.5%	9	0	0	GET	/api/me/export	14.389	1.599	3.768	3.768	2218937
0.4%	11	2	0	PATCH	^/api/posts/[^/]+$	13.992	1.272	3.840	3.840	537
0.4%	96	0	0	GET	^/api/relationships/[^/]+	13.527	0.141	0.787	0.787	1373
0.3%	23	0	0	POST	^/api/posts/[^/]+/comments$	8.908	0.387	1.728	1.728	340
0.3%	19	0	0	PUT	^/api/follows/[^/]+$	8.223	0.433	2.253	2.253	0
0.3%	55	0	0	GET	^/api/users/[^/]+$	8.111	0.147	0.623	0.623	195
0.2%	53	0	0	GET	^/api/posts/[^/]+/realmojis$	7.481	0.141	0.350	0.350	861
0.2%	44	1	0	GET	^/api/posts/[^/]+/comments$	6.092	0.138	0.467	0.467	921
0.2%	26	0	0	GET	^/api/memories/[^/]+	5.793	0.223	0.699	0.699	1178
0.1%	33	0	0	GET	/api/me	4.244	0.129	0.242	0.242	300
0.1%	6	0	0	DELETE	^/api/comments/[^/]+$	2.439	0.407	0.637	0.637	0
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	2.214	0.738	1.260	1.260	0
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 42335 クエリ / 80 種類 / 合計 417.231s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	51.7%	2613	215.779	82.58	739.20	653.5	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	12.9%	141	54.020	383.12	1383.69	6715.9	134.4	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	10.3%	5	42.801	8560.15	10836.45	30282.2	1199793.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
4	8.2%	5	34.080	6816.03	10747.70	22206.4	626866.6	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	4.0%	1733	16.814	9.70	440.84	123.6	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
6	3.8%	5	15.682	3136.39	4541.73	11940.8	529835.4	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
7	2.8%	8	11.681	1460.13	6529.54	225843.9	2672.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
8	2.6%	4396	11.033	2.51	45.98	11.1	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
9	1.2%	1765	5.189	2.94	43.85	26.8	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.4%	392	1.569	4.00	23.01	31.0	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
11	0.3%	8	1.428	178.51	307.00	20937.4	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.3%	8	1.113	139.10	216.39	629.1	25524.3	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
13	0.2%	8	0.998	124.71	163.16	640.0	30596.1	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
14	0.2%	3996	0.925	0.23	28.40	20.3	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
15	0.2%	410	0.782	1.91	19.26	23.2	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: adfcccac25c3860caaccb2c09865215defd3aac8
Type: cpu
Time: 2026-10-31 12:55:19 JST
Duration: 60.05s, Total samples = 33.97s (56.57%)
Showing nodes accounting for 25.36s, 74.65% of 33.97s total
Dropped 881 nodes (cum <= 0.17s)
Showing top 40 nodes out of 306
      flat  flat%   sum%        cum   cum%
     2.33s  6.86%  6.86%      2.33s  6.86%  internal/runtime/syscall/linux.Syscall6
     1.74s  5.12% 11.98%      1.74s  5.12%  runtime.memclrNoHeapPointers
     1.68s  4.95% 16.93%      5.50s 16.19%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.51s  4.45% 21.37%      1.51s  4.45%  crypto/internal/fips140/sha256.blockAVX2
     1.46s  4.30% 25.67%      4.44s 13.07%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.36s  4.00% 29.67%      5.08s 14.95%  runtime.scanObject
     1.34s  3.94% 33.62%      1.67s  4.92%  runtime.findObject
     1.34s  3.94% 37.56%      1.34s  3.94%  runtime.memmove
     1.33s  3.92% 41.48%      1.33s  3.92%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.24s  3.65% 45.13%      1.24s  3.65%  aeshashbody
     1.23s  3.62% 48.75%      2.05s  6.03%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
        1s  2.94% 51.69%      1.71s  5.03%  runtime.tryDeferToSpanScan
     0.87s  2.56% 54.25%      0.92s  2.71%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.68s  2.00% 56.26%      0.86s  2.53%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.54s  1.59% 57.85%      0.54s  1.59%  example.com/isureal-company/isureal.(*UserRow).active
     0.52s  1.53% 59.38%      0.52s  1.53%  runtime.spanClass.sizeclass (inline)
     0.46s  1.35% 60.73%      0.46s  1.35%  memeqbody
     0.44s  1.30% 62.03%      1.75s  5.15%  runtime.mapaccess1_faststr
     0.34s  1.00% 63.03%      0.51s  1.50%  example.com/isureal-company/isureal.postLess
     0.30s  0.88% 63.91%      0.30s  0.88%  indexbody
     0.26s  0.77% 64.67%      0.26s  0.77%  runtime.nextFreeFast (inline)
     0.24s  0.71% 65.38%      0.43s  1.27%  database/sql.convertAssignRows
     0.23s  0.68% 66.06%      0.23s  0.68%  sync/atomic.(*Int32).Add
     0.22s  0.65% 66.71%      0.41s  1.21%  runtime.greyobject
     0.21s  0.62% 67.32%      0.21s  0.62%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.20s  0.59% 67.91%      1.19s  3.50%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.20s  0.59% 68.50%      0.22s  0.65%  runtime.typePointers.next
     0.19s  0.56% 69.06%      0.19s  0.56%  runtime.futex
     0.19s  0.56% 69.62%      0.96s  2.83%  runtime.mallocgcSmallScanNoHeader
     0.18s  0.53% 70.15%      0.18s  0.53%  crypto/internal/fips140/aes/gcm.gcmAesEnc
     0.18s  0.53% 70.68%      3.14s  9.24%  runtime.mallocgc
     0.18s  0.53% 71.21%      1.75s  5.15%  runtime.mapassign_faststr
     0.17s   0.5% 71.71%      0.33s  0.97%  runtime.scanObjectsSmall
     0.17s   0.5% 72.21%      0.17s   0.5%  runtime.typePointers.nextFast (inline)
     0.16s  0.47% 72.68%      0.19s  0.56%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.14s  0.41% 73.09%      0.17s   0.5%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.14s  0.41% 73.51%      0.19s  0.56%  runtime.heapArenaOf (inline)
     0.13s  0.38% 73.89%      1.71s  5.03%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.13s  0.38% 74.27%      0.49s  1.44%  runtime.convTstring
     0.13s  0.38% 74.65%      0.48s  1.41%  runtime.mapaccess2_faststr
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 45.2% / 100% | 30% | 0.1% | 6 | 166 MB | run queue が積もっている |
| isu2 (2 cpu) | 17.3% / 65% | 0% | 0.3% | 5 | 682 MB | run queue が積もっている |
| isu3 (2 cpu) | 18.8% / 63% | 0% | 0.3% | 3 | 646 MB |  |
| isu4 (2 cpu) | 48.4% / 100% | 15% | 17.2% | 33 | 106 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 74% / nginx 10% / journalctl 2% |
| isu2 | isureal 33% |
| isu3 | isureal 36% |
| isu4 | postgres 51% / isureal 6% |

## Diff from previous run

スコア: 63311 → 62243 (-1068)

`86828dc..02ae5ac`

```
 webapp/go/repl.go           | 25 +++++++++++++------------
 webapp/go/replica_server.go | 24 ++++++++++++++++++------
 2 files changed, 31 insertions(+), 18 deletions(-)
```
