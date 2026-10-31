# run 20261031-135652-8b16848

## Summary

| | |
|---|---|
| score | 110598 |
| status | pass |
| ref | 8b16848 |
| sha | 8b168481ecd6 |
| 所要時間 | 267s |
| 前回比 | +45898 (70.9%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +45898 (70.9%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | 複製 検証: shadow＋照合（利用者を最後に読む修正、main c6abd65 取り込み済み） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	129	327	715	715
europe	7	19	54	147	164	163
america	6	18	54	99	114	114
```

```
{"id":"15209","status":"done","portalStatus":"finished","passed":true,"score":110598,"raw":110598,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:59:48.057Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:59:51.021Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:59:54.021Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:59:57.871Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T05:00:00.023Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T05:00:01.539Z level=INFO msg="user moved" user=koyori5393 from=asia to=america
time=2026-10-31T05:00:02.645Z level=INFO msg="user moved" user=koyori5393 from=america to=asia
time=2026-10-31T05:00:03.021Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T05:00:06.045Z level=INFO msg="asia 3 日目: 129 人"
time=2026-10-31T05:00:06.719Z level=INFO msg="user moved" user=noazoe5561 from=asia to=america
time=2026-10-31T05:00:08.068Z level=INFO msg="user moved" user=koyori5393 from=asia to=america
time=2026-10-31T05:00:09.048Z level=INFO msg="europe 3 日目: 54 人"
time=2026-10-31T05:00:09.943Z level=INFO msg="user moved" user=kehozoe-5850 from=asia to=america
time=2026-10-31T05:00:12.069Z level=INFO msg="america 3 日目: 54 人"
time=2026-10-31T05:00:12.118Z level=INFO msg="user moved" user=noazoe5561 from=america to=asia
time=2026-10-31T05:00:15.050Z level=INFO msg="asia 4 日目: 327 人"
time=2026-10-31T05:00:18.325Z level=INFO msg="europe 4 日目: 147 人"
time=2026-10-31T05:00:19.098Z level=INFO msg="user moved" user=noazoe5561 from=asia to=america
time=2026-10-31T05:00:19.555Z level=INFO msg="user moved" user=suto_7109 from=asia to=america
time=2026-10-31T05:00:19.747Z level=INFO msg="user moved" user=beneva.6547 from=asia to=europe
time=2026-10-31T05:00:21.161Z level=INFO msg="america 4 日目: 99 人"
time=2026-10-31T05:00:21.163Z level=INFO msg="user moved" user=ulsolchi-6903 from=asia to=europe
time=2026-10-31T05:00:21.195Z level=INFO msg="user moved" user=ischileo_7507 from=asia to=america
time=2026-10-31T05:00:23.172Z level=INFO msg="user moved" user=teshi.7262 from=asia to=america
time=2026-10-31T05:00:23.685Z level=INFO msg="user moved" user=koyori5393 from=america to=asia
time=2026-10-31T05:00:23.855Z level=INFO msg="user moved" user=sean-7883 from=europe to=asia
time=2026-10-31T05:00:24.496Z level=INFO msg="asia 5 日目: 715 人"
time=2026-10-31T05:00:28.096Z level=INFO msg="europe 5 日目: 164 人"
time=2026-10-31T05:00:30.229Z level=INFO msg="america 5 日目: 114 人"
time=2026-10-31T05:00:39.497Z level=INFO msg="asia 最終日（グランプリ）: 715 人"
time=2026-10-31T05:00:39.705Z level=INFO msg="america 最終日（グランプリ）: 114 人"
time=2026-10-31T05:00:39.889Z level=INFO msg="europe 最終日（グランプリ）: 163 人"
time=2026-10-31T05:00:48.048Z level=INFO msg="グランプリ: アジア 349 人, ヨーロッパ 63 人, アメリカ 73 人"
{
  "pass": true,
  "score": 110598,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.5%	1258	0	0	GET	/api/notifications	10954.967	8.708	15.281	15.288	169
18.1%	1888	35	0	POST	/api/posts	5291.375	2.803	9.203	10.164	639
9.5%	4236	78	0	POST	/api/friend-requests	2776.466	0.655	5.006	7.917	113
8.0%	12314	349	0	GET	^/api/images/[^/]+	2345.056	0.190	2.761	6.576	148539
5.0%	4783	1123	0	POST	/api/auth/login	1472.404	0.308	3.578	8.226	193
4.7%	2192	11	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1381.344	0.630	5.387	7.018	116
2.8%	583	9	0	POST	^/api/posts/[^/]+/realmojis$	806.360	1.383	9.090	9.355	469
2.3%	536	126	0	PUT	/api/me/profile-image	666.193	1.243	6.249	8.441	5
1.9%	537	6	0	POST	^/api/posts/[^/]+/comments$	569.106	1.060	6.233	6.925	343
1.7%	737	1	0	POST	/api/auth/register	502.789	0.682	3.129	4.651	227
1.4%	3033	0	0	GET	/api/memories	406.426	0.134	0.960	2.849	11181
0.9%	451	0	0	PUT	^/api/follows/[^/]+$	274.296	0.608	4.444	6.452	0
0.9%	11863	0	0	GET	^/api/feeds/[^/]+	268.871	0.023	0.197	5.308	3279
0.8%	363	0	0	PATCH	/api/me	242.458	0.668	4.435	7.005	274
0.7%	632	0	0	POST	/api/auth/logout	209.338	0.331	3.836	7.015	0
0.6%	219	0	0	GET	/api/me/export	182.391	0.833	3.792	4.898	2051963
0.6%	93	0	0	DELETE	^/api/posts/[^/]+$	181.944	1.956	7.935	7.935	0
0.5%	173	6	0	PATCH	^/api/posts/[^/]+$	155.305	0.898	5.915	6.866	624
0.3%	2659	0	0	GET	/api/posting-window/current	87.909	0.033	0.740	2.802	237
0.2%	1529	39	0	GET	^/api/posts/me/current$	55.656	0.036	0.239	2.853	633
0.2%	102	0	0	DELETE	^/api/comments/[^/]+$	54.799	0.537	3.352	5.072	0
0.2%	89	0	0	DELETE	^/api/realmojis/[^/]+$	51.504	0.579	3.763	3.763	0
0.1%	611	0	0	GET	^/api/memories/[^/]+	41.278	0.068	0.569	1.236	702
0.1%	1956	10	0	GET	^/api/posts/[^/]+$	37.000	0.019	0.152	2.850	622
0.1%	1563	0	0	GET	^/api/relationships/[^/]+	33.675	0.022	0.182	2.848	1098
0.1%	2052	0	0	GET	/api/users	32.932	0.016	0.135	1.078	243
0.1%	941	0	0	GET	^/api/users/[^/]+$	20.219	0.021	0.211	1.241	195
0.0%	29	0	0	PUT	^/api/blocks/[^/]+$	13.753	0.474	2.411	2.411	0
0.0%	913	0	0	GET	^/api/posts/[^/]+/realmojis$	12.980	0.014	0.145	1.234	909
0.0%	1	0	0	POST	/api/initialize	12.664	12.664	12.664	12.664	90
…（上位 30 / 全 44 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.9%	387	1	0	GET	/api/notifications	3320.683	8.581	11.695	13.141	176
12.9%	510	7	0	POST	/api/posts	1133.856	2.223	7.206	8.153	636
7.2%	1420	0	0	GET	^/api/feeds/[^/]+	634.336	0.447	3.521	3.770	2697
6.3%	898	11	0	POST	/api/friend-requests	555.737	0.619	4.668	7.583	114
6.2%	1898	20	0	GET	^/api/images/[^/]+	546.136	0.288	1.392	3.075	185647
4.5%	555	0	0	POST	/api/auth/login	397.494	0.716	5.406	7.121	247
3.9%	683	0	0	GET	/api/posting-window/current	342.445	0.501	1.520	3.286	239
3.5%	621	0	0	GET	/api/memories	305.436	0.492	3.736	3.756	10910
3.3%	271	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	293.353	1.082	5.683	6.143	116
2.0%	417	26	0	GET	^/api/posts/me/current$	172.325	0.413	1.618	3.756	603
1.8%	435	0	0	GET	/api/users	159.995	0.368	1.680	3.699	242
1.2%	227	0	0	GET	^/api/relationships/[^/]+	103.080	0.454	3.521	3.619	960
1.0%	97	3	0	POST	^/api/posts/[^/]+/realmojis$	91.315	0.941	7.172	7.172	429
1.0%	182	19	0	GET	^/api/posts/[^/]+$	90.321	0.496	3.557	3.756	553
0.9%	161	0	0	POST	/api/auth/register	81.866	0.508	5.393	5.616	228
0.7%	78	0	0	POST	/api/auth/logout	63.429	0.813	5.146	5.146	0
0.7%	28	0	0	GET	/api/me/export	60.520	2.161	5.263	5.263	2058296
0.7%	86	0	0	PUT	/api/me/profile-image	59.791	0.695	3.529	3.529	0
0.6%	60	2	0	POST	^/api/posts/[^/]+/comments$	49.263	0.821	3.863	3.863	317
0.5%	81	1	0	GET	^/api/posts/[^/]+/realmojis$	40.159	0.496	3.756	3.756	830
0.4%	46	8	0	PATCH	^/api/posts/[^/]+$	36.317	0.789	4.361	4.361	543
0.4%	78	0	0	GET	^/api/memories/[^/]+	33.555	0.430	2.108	2.108	811
0.3%	22	0	0	DELETE	^/api/posts/[^/]+$	30.467	1.385	4.187	4.187	0
0.3%	79	0	0	GET	^/api/users/[^/]+$	29.525	0.374	2.183	2.183	196
0.3%	72	0	0	PUT	^/api/follows/[^/]+$	28.270	0.393	2.975	2.975	0
0.3%	75	0	0	PATCH	/api/me	26.907	0.359	1.904	1.904	272
0.3%	92	2	0	GET	/api/me	26.346	0.286	1.544	1.544	295
0.2%	41	0	0	GET	^/api/posts/[^/]+/comments$	20.156	0.492	3.756	3.756	850
0.2%	16	0	0	DELETE	^/api/realmojis/[^/]+$	14.889	0.931	2.474	2.474	0
0.1%	6	0	0	PUT	^/api/blocks/[^/]+$	6.274	1.046	4.381	4.381	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.7%	282	0	0	GET	/api/notifications	2167.885	7.688	9.677	9.679	173
10.9%	340	4	0	POST	/api/posts	612.163	1.800	7.232	7.498	641
7.5%	581	5	0	POST	/api/friend-requests	420.541	0.724	5.047	7.002	114
6.6%	828	0	0	GET	^/api/feeds/[^/]+	369.079	0.446	3.829	3.989	2880
5.6%	1318	19	0	GET	^/api/images/[^/]+	313.274	0.238	1.090	2.593	177499
4.8%	383	0	0	POST	/api/auth/login	265.849	0.694	4.921	7.159	250
3.8%	466	0	0	GET	/api/posting-window/current	215.334	0.462	2.576	3.971	240
3.8%	193	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	213.633	1.107	7.402	7.850	116
3.2%	432	0	0	GET	/api/memories	176.262	0.408	1.920	3.865	10974
2.7%	310	23	0	GET	^/api/posts/me/current$	149.055	0.481	3.862	3.865	603
2.2%	283	0	0	GET	/api/users	124.970	0.442	3.691	3.970	241
1.3%	27	0	0	GET	/api/me/export	69.948	2.591	5.906	5.906	2801861
1.0%	140	0	0	GET	^/api/relationships/[^/]+	56.207	0.401	3.888	4.013	1317
1.0%	57	0	0	POST	/api/auth/logout	53.383	0.937	4.025	4.025	0
0.9%	71	0	0	PUT	^/api/follows/[^/]+$	48.391	0.682	3.912	3.912	0
0.9%	31	3	0	POST	^/api/posts/[^/]+/realmojis$	47.654	1.537	4.902	4.902	418
0.7%	75	13	0	GET	^/api/posts/[^/]+$	37.305	0.497	3.858	3.858	514
0.7%	108	0	0	POST	/api/auth/register	37.207	0.345	2.154	2.518	229
0.6%	53	0	0	PUT	/api/me/profile-image	33.692	0.636	3.564	3.564	0
0.6%	30	3	0	PATCH	^/api/posts/[^/]+$	33.481	1.116	5.917	5.917	593
0.6%	21	0	0	DELETE	^/api/posts/[^/]+$	31.675	1.508	5.204	5.204	0
0.4%	50	0	0	PATCH	/api/me	22.208	0.444	2.659	2.659	270
0.4%	51	0	0	GET	^/api/memories/[^/]+	21.021	0.412	1.590	1.590	1080
0.3%	56	0	0	GET	/api/me	18.493	0.330	0.929	0.929	302
0.3%	45	0	0	GET	^/api/users/[^/]+$	17.212	0.382	2.592	2.592	197
0.2%	29	0	0	GET	^/api/posts/[^/]+/comments$	10.705	0.369	2.365	2.365	704
0.2%	43	0	0	POST	^/api/posts/[^/]+/comments$	9.395	0.218	1.839	1.839	283
0.1%	4	0	0	DELETE	^/api/realmojis/[^/]+$	7.972	1.993	2.569	2.569	0
0.1%	16	0	0	GET	^/api/posts/[^/]+/realmojis$	7.649	0.478	1.933	1.933	1352
0.0%	1	0	0	DELETE	^/api/follows/[^/]+$	1.748	1.748	1.748	1.748	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 67877 クエリ / 81 種類 / 合計 655.476s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	62.3%	3948	408.680	103.52	716.49	665.4	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	17.5%	274	114.907	419.37	1752.29	1338.5	136.3	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	5.8%	2857	38.026	13.31	1559.24	120.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	3.8%	8213	25.116	3.06	45.25	10.7	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	3.3%	3	21.639	7213.09	9630.29	31064.0	1199601.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	2.4%	3	15.865	5288.48	7178.17	24601.0	626978.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	1.3%	2703	8.326	3.08	24.24	24.8	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
8	1.2%	3	7.979	2659.65	3663.26	12724.0	529726.3	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	0.5%	696	3.001	4.31	58.54	28.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.3%	6	1.980	330.07	457.19	27823.2	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.2%	5602	1.362	0.24	32.10	19.6	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
12	0.2%	632	1.316	2.08	20.56	19.9	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
13	0.2%	6	1.190	198.27	305.43	216791.2	2570.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.1%	2581	0.833	0.32	15.25	20.0	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	6	0.696	116.05	172.87	615.7	29768.7	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: c0ced6a2e0f524e3d888a6993a9e0bdafb027692
Type: cpu
Time: 2026-10-31 13:59:29 JST
Duration: 60s, Total samples = 46.24s (77.06%)
Showing nodes accounting for 33.19s, 71.78% of 46.24s total
Dropped 1085 nodes (cum <= 0.23s)
Showing top 40 nodes out of 281
      flat  flat%   sum%        cum   cum%
     5.33s 11.53% 11.53%      5.34s 11.55%  crypto/internal/fips140/sha256.blockAVX2
     4.66s 10.08% 21.60%      4.66s 10.08%  internal/runtime/syscall/linux.Syscall6
     2.08s  4.50% 26.10%      4.51s  9.75%  example.com/isureal-company/isureal.(*Store).digestLocked
     1.83s  3.96% 30.06%      2.42s  5.23%  runtime.findObject
     1.75s  3.78% 33.85%      1.75s  3.78%  runtime.memmove
     1.64s  3.55% 37.39%      1.64s  3.55%  example.com/isureal-company/isureal.h64 (inline)
     1.62s  3.50% 40.90%      5.79s 12.52%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.48s  3.20% 44.10%      1.48s  3.20%  aeshashbody
     1.48s  3.20% 47.30%      1.48s  3.20%  runtime.memclrNoHeapPointers
     1.45s  3.14% 50.43%      1.45s  3.14%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.33s  2.88% 53.31%      6.48s 14.01%  runtime.scanObject
     1.09s  2.36% 55.67%      2.24s  4.84%  runtime.tryDeferToSpanScan
     0.79s  1.71% 57.37%      0.79s  1.71%  runtime.spanClass.sizeclass (inline)
     0.64s  1.38% 58.76%      2.15s  4.65%  runtime.mapaccess1_faststr
     0.44s  0.95% 59.71%      0.44s  0.95%  runtime.nextFreeFast (inline)
     0.36s  0.78% 60.49%      0.51s  1.10%  internal/runtime/maps.(*Iter).Next
     0.36s  0.78% 61.27%      0.36s  0.78%  memeqbody
     0.36s  0.78% 62.05%      2.18s  4.71%  runtime.mapassign_faststr
     0.32s  0.69% 62.74%      0.56s  1.21%  example.com/isureal-company/isureal.postLess
     0.30s  0.65% 63.39%      5.79s 12.52%  crypto/internal/fips140/sha256.(*Digest).Write
     0.29s  0.63% 64.01%      0.32s  0.69%  encoding/json.appendString[go.shape.string]
     0.28s  0.61% 64.62%      0.28s  0.61%  runtime.futex
     0.28s  0.61% 65.22%      4.87s 10.53%  runtime.mallocgc
     0.28s  0.61% 65.83%      1.19s  2.57%  runtime.mallocgcSmallScanNoHeader
     0.26s  0.56% 66.39%      0.65s  1.41%  database/sql.convertAssignRows
     0.25s  0.54% 66.93%      5.08s 10.99%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.24s  0.52% 67.45%      0.32s  0.69%  runtime.(*mspan).markBitsForIndex (inline)
     0.22s  0.48% 67.93%      0.27s  0.58%  runtime.spanOf (inline)
     0.20s  0.43% 68.36%      0.59s  1.28%  runtime.greyobject
     0.19s  0.41% 68.77%      0.53s  1.15%  runtime.scanObjectsSmall
     0.16s  0.35% 69.12%      5.72s 12.37%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.16s  0.35% 69.46%      0.28s  0.61%  runtime.heapArenaOf (inline)
     0.15s  0.32% 69.79%      3.53s  7.63%  example.com/isureal-company/isureal.listMemories
     0.15s  0.32% 70.11%      0.25s  0.54%  runtime.typePointers.next
     0.14s   0.3% 70.42%      5.35s 11.57%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.13s  0.28% 70.70%      0.29s  0.63%  compress/flate.(*compressor).deflate
     0.13s  0.28% 70.98%      0.27s  0.58%  internal/stringslite.Index
     0.13s  0.28% 71.26%      0.54s  1.17%  runtime.mapaccess2_faststr
     0.12s  0.26% 71.52%      0.35s  0.76%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.12s  0.26% 71.78%      2.72s  5.88%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 54.0% / 100% | 28% | 0.4% | 10 | 293 MB | run queue が積もっている |
| isu2 (2 cpu) | 30.4% / 100% | 4% | 0.7% | 4 | 124 MB |  |
| isu3 (2 cpu) | 30.8% / 100% | 3% | 1.4% | 4 | 120 MB |  |
| isu4 (2 cpu) | 55.9% / 100% | 30% | 13.7% | 35 | 107 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 81% / nginx 11% |
| isu2 | isureal 58% / nginx 1% |
| isu3 | isureal 57% / nginx 1% |
| isu4 | postgres 66% / isureal 3% |

## Diff from previous run

スコア: 64700 → 110598 (+45898)

`e0e4cb3..8b16848`

```
 conf/_all/etc/sysctl.d/99-isucon.conf              |   2 +
 conf/isu1/etc/nginx/sites-enabled/isureal.conf     |   8 +
 conf/isu1/home/isucon/env.sh                       |   2 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf     |  26 +-
 conf/isu2/home/isucon/env.sh                       |   5 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf     |  26 +-
 conf/isu3/home/isucon/env.sh                       |   5 +
 .../etc/postgresql/16/main/conf.d/zz-tuning.conf   |   2 +
 webapp/go/account_deletion_api.go                  |  10 +-
 webapp/go/action_log.go                            |  13 +-
 webapp/go/action_log_test.go                       |  67 ++
 webapp/go/admin_posting_window_api.go              |  16 +-
 webapp/go/admin_user_api.go                        |   5 +-
 webapp/go/cmd/isureal/main.go                      |   6 +
 webapp/go/comment_api.go                           |   9 +-
 webapp/go/config.go                                |   2 +
 webapp/go/export_api.go                            |   6 +-
 webapp/go/feed_api.go                              | 113 ++-
 webapp/go/feed_test.go                             | 331 ++++++++
 webapp/go/friend_request_api.go                    |  15 +-
 webapp/go/image_local.go                           |   4 +-
 webapp/go/initialize.go                            |   4 +-
 webapp/go/offload.go                               |  46 +-
 webapp/go/post_api.go                              |  78 +-
 webapp/go/post_forward.go                          | 113 +++
 webapp/go/post_forward_test.go                     |  94 +++
 webapp/go/profile_image_api.go                     |  20 +-
 webapp/go/realmoji_api.go                          |  16 +-
 webapp/go/relationship.go                          |  10 +-
 webapp/go/relationship_api.go                      |  33 +-
 webapp/go/repl.go                                  | 922 +++++++++++++++++++++
 webapp/go/repl_digest.go                           | 151 ++++
 webapp/go/repl_test.go                             |  62 ++
 webapp/go/replica_server.go                        | 255 ++++++
 webapp/go/server.go                                |   1 +
 webapp/go/store.go                                 | 272 ++++--
 webapp/go/user_api.go                              |  34 +-
 webapp/go/user_read_api.go                         |  58 +-
 webapp/go/user_search_test.go                      |  94 +++
 39 files changed, 2680 insertions(+), 256 deletions(-)
```
