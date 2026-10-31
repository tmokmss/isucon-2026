# run 20261031-135832-1338f93

## Summary

| | |
|---|---|
| score | 112525 |
| status | pass |
| ref | lane2-tune |
| sha | 1338f938d563 |
| 所要時間 | 442s |
| 前回比 | -2875 (-2.5%) vs 20261031-134616-f256785 |
| ベスト比 | -2875 (-2.5%) vs 20261031-134616-f256785 = 115400 |
| note | pool 100 + GOGC 400 on c6abd65 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	144	415	606	603
europe	8	19	54	153	154	154
america	7	16	52	104	108	107
```

```
{"id":"15550","status":"done","portalStatus":"finished","passed":true,"score":112525,"raw":112525,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:04:13.471Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:04:15.432Z level=INFO msg="user moved" user=solnomu-5028 from=asia to=europe
time=2026-10-31T05:04:16.438Z level=INFO msg="europe 1 日目: 8 人"
time=2026-10-31T05:04:17.153Z level=INFO msg="user moved" user=yohe5062 from=asia to=america
time=2026-10-31T05:04:19.439Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T05:04:22.460Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T05:04:25.439Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T05:04:28.437Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T05:04:29.563Z level=INFO msg="user moved" user=solnomu-5028 from=europe to=america
time=2026-10-31T05:04:30.516Z level=INFO msg="user moved" user=temu_5450 from=europe to=america
time=2026-10-31T05:04:31.463Z level=INFO msg="asia 3 日目: 144 人"
time=2026-10-31T05:04:33.670Z level=INFO msg="user moved" user=yohe5062 from=america to=europe
time=2026-10-31T05:04:33.919Z level=INFO msg="user moved" user=nasetsu-5795 from=asia to=america
time=2026-10-31T05:04:34.443Z level=INFO msg="europe 3 日目: 54 人"
time=2026-10-31T05:04:34.594Z level=INFO msg="user moved" user=zoeri_5999 from=asia to=america
time=2026-10-31T05:04:37.438Z level=INFO msg="america 3 日目: 52 人"
time=2026-10-31T05:04:39.446Z level=INFO msg="user moved" user=temu_5450 from=america to=europe
time=2026-10-31T05:04:40.211Z level=INFO msg="user moved" user=yohe5062 from=europe to=asia
time=2026-10-31T05:04:40.461Z level=INFO msg="asia 4 日目: 415 人"
time=2026-10-31T05:04:43.179Z level=INFO msg="user moved" user=isre.6572 from=america to=asia
time=2026-10-31T05:04:43.710Z level=INFO msg="europe 4 日目: 153 人"
time=2026-10-31T05:04:44.343Z level=INFO msg="user moved" user=isre.6572 from=asia to=america
time=2026-10-31T05:04:46.639Z level=INFO msg="america 4 日目: 104 人"
time=2026-10-31T05:04:46.908Z level=INFO msg="user moved" user=resu7131 from=asia to=america
time=2026-10-31T05:04:47.431Z level=INFO msg="user moved" user=kosolre_6965 from=asia to=europe
time=2026-10-31T05:04:48.668Z level=INFO msg="user moved" user=yohe5062 from=asia to=america
time=2026-10-31T05:04:49.745Z level=INFO msg="asia 5 日目: 606 人"
time=2026-10-31T05:04:49.783Z level=INFO msg="user moved" user=temu_5450 from=europe to=america
time=2026-10-31T05:04:49.926Z level=INFO msg="user moved" user=kusu8019 from=asia to=america
time=2026-10-31T05:04:50.064Z level=INFO msg="user moved" user=liasu7352 from=asia to=europe
time=2026-10-31T05:04:52.835Z level=INFO msg="europe 5 日目: 154 人"
time=2026-10-31T05:04:55.477Z level=INFO msg="america 5 日目: 108 人"
time=2026-10-31T05:05:04.909Z level=INFO msg="asia 最終日（グランプリ）: 603 人"
time=2026-10-31T05:05:04.993Z level=INFO msg="europe 最終日（グランプリ）: 154 人"
time=2026-10-31T05:05:05.200Z level=INFO msg="america 最終日（グランプリ）: 107 人"
time=2026-10-31T05:05:13.464Z level=INFO msg="グランプリ: アジア 312 人, ヨーロッパ 108 人, アメリカ 76 人"
{
  "pass": true,
  "score": 112525,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.9%	1259	0	0	GET	/api/notifications	11588.796	9.205	15.230	15.368	171
19.4%	1796	30	0	POST	/api/posts	5351.816	2.980	9.401	11.884	645
6.9%	12618	452	0	GET	^/api/images/[^/]+	1913.635	0.152	2.172	9.071	139260
6.5%	3799	72	0	POST	/api/friend-requests	1800.824	0.474	3.409	6.233	113
4.9%	1971	14	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1345.630	0.683	4.128	7.415	115
4.2%	4613	906	0	POST	/api/auth/login	1167.424	0.253	3.402	6.812	202
2.9%	574	20	0	POST	^/api/posts/[^/]+/realmojis$	798.782	1.392	9.048	9.366	472
1.8%	495	129	0	PUT	/api/me/profile-image	487.245	0.984	5.757	8.104	6
1.7%	3134	0	0	GET	/api/memories	471.279	0.150	0.924	3.948	11009
1.6%	538	17	0	POST	^/api/posts/[^/]+/comments$	439.243	0.816	4.520	5.491	328
1.4%	11149	0	0	GET	^/api/feeds/[^/]+	381.280	0.034	0.300	1.491	2601
1.0%	251	0	0	GET	/api/me/export	285.131	1.136	8.497	9.248	1939466
0.8%	628	1	0	POST	/api/auth/register	230.065	0.366	2.536	3.718	227
0.7%	407	0	0	PUT	^/api/follows/[^/]+$	200.349	0.492	2.968	6.494	0
0.6%	696	0	0	POST	/api/auth/logout	170.331	0.245	3.133	4.040	0
0.5%	308	0	0	PATCH	/api/me	151.561	0.492	3.138	3.908	279
0.5%	182	5	0	PATCH	^/api/posts/[^/]+$	147.924	0.813	5.200	5.531	640
0.4%	85	0	0	DELETE	^/api/posts/[^/]+$	107.469	1.264	5.195	5.195	0
0.3%	2479	0	0	GET	/api/posting-window/current	89.300	0.036	0.269	1.566	237
0.3%	1508	44	0	GET	^/api/posts/me/current$	78.685	0.052	0.387	0.505	636
0.2%	93	0	0	DELETE	^/api/comments/[^/]+$	53.225	0.572	5.252	5.252	0
0.2%	1749	17	0	GET	^/api/posts/[^/]+$	48.385	0.028	0.245	0.429	606
0.2%	1821	0	0	GET	/api/users	46.503	0.026	0.301	0.937	242
0.2%	1421	0	0	GET	^/api/relationships/[^/]+	43.472	0.031	0.273	0.529	1050
0.2%	622	0	0	GET	^/api/memories/[^/]+	42.867	0.069	0.664	1.148	676
0.1%	74	0	0	DELETE	^/api/realmojis/[^/]+$	40.627	0.549	1.950	1.950	0
0.1%	37	0	0	PUT	^/api/blocks/[^/]+$	32.250	0.872	6.931	6.931	0
0.1%	751	0	0	GET	^/api/posts/[^/]+/realmojis$	21.644	0.029	0.284	0.552	977
0.1%	887	0	0	GET	^/api/users/[^/]+$	21.456	0.024	0.259	0.437	191
0.1%	1	0	0	POST	/api/initialize	18.254	18.254	18.254	18.254	90
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.8%	378	0	0	GET	/api/notifications	3160.479	8.361	12.278	12.551	177
15.7%	496	1	0	POST	/api/posts	1246.909	2.514	6.373	7.466	643
7.2%	906	21	0	POST	/api/friend-requests	571.187	0.630	3.741	5.674	113
6.4%	2088	19	0	GET	^/api/images/[^/]+	510.800	0.245	1.430	3.195	183544
4.9%	579	0	0	POST	/api/auth/login	387.653	0.670	3.919	5.014	248
4.9%	1834	0	0	GET	^/api/feeds/[^/]+	387.584	0.211	0.689	1.197	2520
2.7%	267	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	217.479	0.815	3.915	4.287	116
2.7%	674	0	0	GET	/api/posting-window/current	211.883	0.314	0.765	1.206	238
2.4%	650	0	0	GET	/api/memories	188.605	0.290	1.294	3.757	9460
1.7%	130	3	0	POST	^/api/posts/[^/]+/realmojis$	137.462	1.057	7.027	7.515	455
1.4%	446	35	0	GET	^/api/posts/me/current$	109.441	0.245	0.713	1.100	594
1.2%	41	0	0	GET	/api/me/export	97.562	2.380	7.987	7.987	2204714
1.2%	119	2	0	POST	^/api/posts/[^/]+/comments$	91.526	0.769	3.342	5.673	334
1.1%	436	0	0	GET	/api/users	85.752	0.197	0.684	1.008	244
0.9%	151	0	0	POST	/api/auth/register	72.823	0.482	2.525	2.746	228
0.7%	76	0	0	PUT	/api/me/profile-image	57.005	0.750	2.024	2.024	0
0.7%	84	0	0	POST	/api/auth/logout	52.060	0.620	6.121	6.121	0
0.6%	60	7	0	PATCH	^/api/posts/[^/]+$	48.185	0.803	3.087	3.087	576
0.6%	246	0	0	GET	^/api/relationships/[^/]+	45.010	0.183	0.616	1.206	1085
0.5%	227	11	0	GET	^/api/posts/[^/]+$	43.585	0.192	0.660	0.947	579
0.5%	25	0	0	DELETE	^/api/posts/[^/]+$	39.927	1.597	3.519	3.519	0
0.5%	74	0	0	PATCH	/api/me	37.433	0.506	2.889	2.889	271
0.4%	97	0	0	PUT	^/api/follows/[^/]+$	31.891	0.329	1.613	1.613	0
0.2%	100	0	0	GET	^/api/users/[^/]+$	17.469	0.175	0.376	0.423	188
0.2%	78	0	0	GET	^/api/posts/[^/]+/comments$	16.976	0.218	0.688	0.688	1673
0.2%	22	0	0	DELETE	^/api/realmojis/[^/]+$	16.839	0.765	5.860	5.860	0
0.2%	78	0	0	GET	^/api/memories/[^/]+	16.214	0.208	0.695	0.695	825
0.2%	85	0	0	GET	^/api/posts/[^/]+/realmojis$	15.842	0.186	1.221	1.221	1189
0.2%	81	1	0	GET	/api/me	12.661	0.156	1.177	1.177	297
0.1%	17	0	0	DELETE	^/api/comments/[^/]+$	8.329	0.490	1.480	1.480	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.9%	277	1	0	GET	/api/notifications	2094.230	7.560	9.726	9.727	178
14.3%	342	5	0	POST	/api/posts	769.311	2.249	6.467	6.961	642
7.7%	603	12	0	POST	/api/friend-requests	414.992	0.688	4.445	6.236	113
7.3%	1460	9	0	GET	^/api/images/[^/]+	395.236	0.271	2.912	3.525	182570
5.0%	297	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	271.315	0.914	4.652	8.221	116
4.5%	1130	0	0	GET	^/api/feeds/[^/]+	241.422	0.214	0.720	1.337	2642
4.4%	371	0	0	POST	/api/auth/login	235.937	0.636	3.594	3.939	250
2.9%	493	0	0	GET	/api/memories	157.582	0.320	1.504	2.933	9839
2.3%	462	0	0	GET	/api/posting-window/current	123.551	0.267	0.712	1.300	239
1.4%	311	14	0	GET	^/api/posts/me/current$	73.560	0.237	0.798	0.810	622
1.3%	304	0	0	GET	/api/users	68.752	0.226	0.859	1.220	252
1.3%	47	7	0	POST	^/api/posts/[^/]+/realmojis$	68.180	1.451	8.173	8.173	397
1.3%	25	0	0	GET	/api/me/export	67.427	2.697	7.582	7.582	3293814
1.0%	69	0	0	PUT	^/api/follows/[^/]+$	51.671	0.749	7.303	7.303	0
0.9%	230	0	0	GET	^/api/relationships/[^/]+	48.111	0.209	0.720	1.092	1264
0.9%	51	0	0	PUT	/api/me/profile-image	46.795	0.918	3.333	3.333	0
0.8%	46	0	0	POST	/api/auth/logout	42.644	0.927	4.740	4.740	0
0.7%	100	0	0	POST	/api/auth/register	37.624	0.376	1.601	1.610	229
0.7%	60	3	0	POST	^/api/posts/[^/]+/comments$	35.454	0.591	3.858	3.858	336
0.5%	167	11	0	GET	^/api/posts/[^/]+$	28.588	0.171	0.506	0.547	576
0.4%	58	0	0	PATCH	/api/me	24.001	0.414	3.121	3.121	277
0.3%	54	1	0	GET	/api/me	14.711	0.272	1.344	1.344	296
0.3%	55	0	0	GET	^/api/memories/[^/]+	14.091	0.256	1.106	1.106	895
0.3%	12	0	0	DELETE	^/api/posts/[^/]+$	13.735	1.145	2.921	2.921	0
0.2%	55	0	0	GET	^/api/users/[^/]+$	10.668	0.194	0.491	0.491	186
0.2%	60	1	0	GET	^/api/posts/[^/]+/realmojis$	10.074	0.168	0.566	0.566	747
0.2%	16	3	0	PATCH	^/api/posts/[^/]+$	9.004	0.563	1.730	1.730	540
0.1%	47	0	0	GET	^/api/posts/[^/]+/comments$	8.079	0.172	0.799	0.799	1321
0.1%	18	0	0	DELETE	^/api/comments/[^/]+$	5.410	0.301	1.565	1.565	0
0.0%	6	0	0	DELETE	^/api/realmojis/[^/]+$	2.382	0.397	0.785	0.785	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 67133 クエリ / 85 種類 / 合計 1689.216s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	73.3%	3817	1237.378	324.18	2484.44	706.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	9.8%	317	164.779	519.81	2765.56	1326.1	133.0	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	3.9%	3113	66.152	21.25	552.36	127.1	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	3.7%	2602	61.968	23.82	414.04	27.2	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
5	2.3%	8545	38.802	4.54	227.83	12.1	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
6	1.6%	5178	27.292	5.27	366.58	20.4	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
7	0.9%	2447	14.466	5.91	360.61	21.4	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
8	0.8%	4657	13.674	2.94	413.45	11.6	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
9	0.6%	721	9.854	13.67	375.00	33.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.4%	2521	6.879	2.73	367.65	7.8	1.0	UPDATE friend_requests SET status=$2, updated_at=$3 WHERE id=$1
11	0.4%	695	6.346	9.13	393.23	24.2	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.3%	2	5.247	2623.73	5097.97	237234.5	2793.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
13	0.3%	4659	4.964	1.07	353.97	4.2	0.3	UPDATE login_sessions SET active = $3, invalidated_at = $2, reason = $4 WHERE user_id = $1 AND active = $5
14	0.3%	878	4.957	5.65	284.90	8.7	1.0	INSERT INTO users (id, username, secret_key_hash, display_name, region, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $6)
15	0.3%	493	4.798	9.73	372.05	8.3	1.0	UPDATE users SET profile_image_id=$2,updated_at=$3 WHERE id=$1

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
Build ID: a0d2dd7a9b4d95e76c5f8f31701bf28d5e92508d
Type: cpu
Time: 2026-10-31 14:03:50 JST
Duration: 60s, Total samples = 37.12s (61.86%)
Showing nodes accounting for 26.37s, 71.04% of 37.12s total
Dropped 985 nodes (cum <= 0.19s)
Showing top 40 nodes out of 312
      flat  flat%   sum%        cum   cum%
     3.88s 10.45% 10.45%      3.88s 10.45%  internal/runtime/syscall/linux.Syscall6
     3.08s  8.30% 18.75%      3.08s  8.30%  crypto/internal/fips140/sha256.blockAVX2
     1.78s  4.80% 23.55%      2.23s  6.01%  runtime.findObject
     1.67s  4.50% 28.04%      5.97s 16.08%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.65s  4.45% 32.49%      1.65s  4.45%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.52s  4.09% 36.58%      1.52s  4.09%  runtime.memclrNoHeapPointers
     1.47s  3.96% 40.54%      1.47s  3.96%  aeshashbody
     1.45s  3.91% 44.45%      1.45s  3.91%  runtime.memmove
     1.41s  3.80% 48.25%      6.09s 16.41%  runtime.scanObject
     1.24s  3.34% 51.59%      2.31s  6.22%  runtime.tryDeferToSpanScan
     0.84s  2.26% 53.85%      0.84s  2.26%  runtime.spanClass.sizeclass (inline)
     0.51s  1.37% 55.23%      2.10s  5.66%  runtime.mapaccess1_faststr
     0.44s  1.19% 56.41%      0.44s  1.19%  runtime.nextFreeFast (inline)
     0.41s  1.10% 57.52%      0.59s  1.59%  example.com/isureal-company/isureal.postLess
     0.39s  1.05% 58.57%      0.39s  1.05%  memeqbody
     0.33s  0.89% 59.46%      0.33s  0.89%  runtime.futex
     0.30s  0.81% 60.26%      1.34s  3.61%  runtime.mallocgcSmallScanNoHeader
     0.29s  0.78% 61.05%      0.29s  0.78%  indexbody
     0.28s  0.75% 61.80%      0.57s  1.54%  database/sql.convertAssignRows
     0.28s  0.75% 62.55%      3.64s  9.81%  runtime.mallocgc
     0.26s   0.7% 63.25%      0.29s  0.78%  runtime.(*mspan).writeHeapBitsSmall
     0.25s  0.67% 63.93%      0.56s  1.51%  runtime.greyobject
     0.25s  0.67% 64.60%      2.20s  5.93%  runtime.mapassign_faststr
     0.23s  0.62% 65.22%      0.23s  0.62%  sync/atomic.(*Int32).Add (inline)
     0.21s  0.57% 65.79%      0.25s  0.67%  runtime.(*mspan).markBitsForIndex (inline)
     0.19s  0.51% 66.30%      0.37s     1%  internal/runtime/maps.(*Iter).Next
     0.19s  0.51% 66.81%      0.19s  0.51%  runtime.typePointers.nextFast (inline)
     0.18s  0.48% 67.30%      0.35s  0.94%  compress/flate.(*compressor).deflate
     0.16s  0.43% 67.73%      0.54s  1.45%  runtime.mapaccess2_faststr
     0.15s   0.4% 68.13%      0.48s  1.29%  runtime.scanObjectsSmall
     0.14s  0.38% 68.51%      3.32s  8.94%  crypto/internal/fips140/sha256.(*Digest).Write
     0.13s  0.35% 68.86%      2.71s  7.30%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.13s  0.35% 69.21%      1.09s  2.94%  example.com/isureal-company/isureal.listMemories
     0.13s  0.35% 69.56%      2.83s  7.62%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.12s  0.32% 69.88%      0.86s  2.32%  runtime.mallocgcSmallNoscan
     0.11s   0.3% 70.18%      0.50s  1.35%  encoding/json.structEncoder.encode
     0.09s  0.24% 70.42%      3.85s 10.37%  example.com/isureal-company/isureal.loadBig
     0.08s  0.22% 70.64%      2.84s  7.65%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.08s  0.22% 70.85%      0.55s  1.48%  runtime.convTstring
     0.07s  0.19% 71.04%      3.03s  8.16%  crypto/internal/fips140/hmac.(*HMAC).Sum
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 45.9% / 100% | 24% | 0.4% | 7 | 120 MB | run queue が積もっている |
| isu2 (2 cpu) | 17.2% / 100% | 2% | 0.6% | 3 | 1847 MB |  |
| isu3 (2 cpu) | 14.7% / 99% | 1% | 0.5% | 4 | 1655 MB |  |
| isu4 (2 cpu) | 54.5% / 100% | 30% | 14.8% | 77 | 100 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 68% / nginx 10% |
| isu2 | isureal 32% / nginx 1% |
| isu3 | isureal 27% / nginx 1% |
| isu4 | postgres 57% / kcompactd0 5% / isureal 3% |

## Diff from previous run

スコア: 115400 → 112525 (-2875)

`f256785..1338f93`

```
 conf/isu1/etc/systemd/system/isureal-go.service        | 2 +-
 conf/isu4/etc/postgresql/16/main/conf.d/zz-tuning.conf | 2 +-
 webapp/go/server.go                                    | 5 +++--
 3 files changed, 5 insertions(+), 4 deletions(-)
```
