# run 20261031-130542-c75a0a9

## Summary

| | |
|---|---|
| score | 110454 |
| status | pass |
| ref | HEAD |
| sha | c75a0a91ecce |
| 所要時間 | 353s |
| 前回比 | +25224 (29.6%) vs 20261031-125933-8caa076 |
| ベスト比 | +14761 (15.4%) vs 20261031-125722-c75a0a9 = 95693 **ベスト更新** |
| note | tls-off-export-pk-2 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	141	400	886	886
europe	7	14	33	81	82	82
america	6	14	37	69	70	69
```

```
{"id":"11355","status":"done","portalStatus":"finished","passed":true,"score":110454,"raw":110454,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:10:18.338Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:10:21.304Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:10:24.304Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:10:27.328Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T04:10:30.387Z level=INFO msg="europe 2 日目: 14 人"
time=2026-10-31T04:10:33.304Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T04:10:36.326Z level=INFO msg="asia 3 日目: 141 人"
time=2026-10-31T04:10:39.310Z level=INFO msg="europe 3 日目: 33 人"
time=2026-10-31T04:10:40.073Z level=INFO msg="user moved" user=noaevatsu5960 from=asia to=america
time=2026-10-31T04:10:42.304Z level=INFO msg="america 3 日目: 37 人"
time=2026-10-31T04:10:45.352Z level=INFO msg="asia 4 日目: 400 人"
time=2026-10-31T04:10:48.491Z level=INFO msg="europe 4 日目: 81 人"
time=2026-10-31T04:10:51.536Z level=INFO msg="america 4 日目: 69 人"
time=2026-10-31T04:10:52.008Z level=INFO msg="user moved" user=rieltor.7685 from=asia to=america
time=2026-10-31T04:10:53.964Z level=INFO msg="user moved" user=mianifu.7117 from=asia to=europe
time=2026-10-31T04:10:54.531Z level=INFO msg="asia 5 日目: 886 人"
time=2026-10-31T04:10:57.371Z level=INFO msg="europe 5 日目: 82 人"
time=2026-10-31T04:11:00.589Z level=INFO msg="america 5 日目: 70 人"
time=2026-10-31T04:11:09.423Z level=INFO msg="europe 最終日（グランプリ）: 82 人"
time=2026-10-31T04:11:09.522Z level=INFO msg="asia 最終日（グランプリ）: 886 人"
time=2026-10-31T04:11:09.690Z level=INFO msg="america 最終日（グランプリ）: 69 人"
time=2026-10-31T04:11:18.339Z level=INFO msg="グランプリ: アジア 575 人, ヨーロッパ 21 人, アメリカ 18 人"
{
  "pass": true,
  "score": 110454,
  "messages": [
    "vinmemu.5832: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.0%	1503	1	0	GET	/api/notifications	12317.769	8.195	15.073	15.149	168
14.7%	2284	55	0	POST	/api/posts	4894.754	2.143	7.469	9.391	640
11.2%	5194	71	0	POST	/api/friend-requests	3722.168	0.717	4.444	7.885	114
7.7%	13819	524	0	GET	^/api/images/[^/]+	2575.741	0.186	3.006	9.105	152745
5.7%	5094	944	0	POST	/api/auth/login	1880.186	0.369	4.480	8.882	203
5.5%	2049	10	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1813.067	0.885	5.457	7.646	116
3.0%	658	124	0	PUT	/api/me/profile-image	1013.391	1.540	7.397	9.512	3
2.6%	584	31	0	POST	^/api/posts/[^/]+/realmojis$	848.433	1.453	6.647	8.448	466
2.1%	905	1	0	POST	/api/auth/register	709.506	0.784	5.501	8.423	227
2.0%	779	35	0	POST	^/api/posts/[^/]+/comments$	681.813	0.875	5.194	8.491	330
1.3%	3593	0	0	GET	/api/memories	436.965	0.122	0.663	0.880	10529
1.3%	560	0	0	PUT	^/api/follows/[^/]+$	420.341	0.751	4.095	7.361	0
1.1%	429	0	0	PATCH	/api/me	358.027	0.835	4.503	6.846	281
0.8%	12873	0	0	GET	^/api/feeds/[^/]+	271.687	0.021	0.113	0.789	1914
0.7%	207	5	0	PATCH	^/api/posts/[^/]+$	226.408	1.094	6.044	7.575	645
0.7%	250	0	0	GET	/api/me/export	220.088	0.880	4.793	8.390	1963052
0.7%	708	0	0	POST	/api/auth/logout	217.758	0.308	3.536	6.773	0
0.4%	106	0	0	DELETE	^/api/posts/[^/]+$	135.018	1.274	5.563	7.263	0
0.2%	3267	0	0	GET	/api/posting-window/current	77.309	0.024	0.145	0.347	237
0.2%	149	0	0	DELETE	^/api/comments/[^/]+$	67.822	0.455	2.876	3.705	0
0.2%	1977	43	0	GET	^/api/posts/me/current$	57.142	0.029	0.127	4.074	641
0.2%	2511	0	0	GET	/api/users	50.224	0.020	0.105	0.386	245
0.1%	77	0	0	DELETE	^/api/realmojis/[^/]+$	43.949	0.571	3.447	3.447	0
0.1%	1720	0	0	GET	^/api/relationships/[^/]+	37.451	0.022	0.111	3.134	986
0.1%	680	0	0	GET	^/api/memories/[^/]+	34.523	0.051	0.368	0.679	641
0.1%	1742	17	0	GET	^/api/posts/[^/]+$	31.308	0.018	0.108	0.172	623
0.0%	1	0	0	POST	/api/initialize	16.321	16.321	16.321	16.321	90
0.0%	923	0	0	GET	^/api/users/[^/]+$	15.724	0.017	0.098	0.186	205
0.0%	27	0	0	PUT	^/api/blocks/[^/]+$	13.747	0.509	2.274	2.274	0
0.0%	750	0	0	GET	^/api/posts/[^/]+/comments$	12.905	0.017	0.111	0.215	1945
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.5%	215	0	0	GET	/api/notifications	1826.554	8.496	11.727	12.222	181
18.6%	257	8	0	POST	/api/posts	818.770	3.186	8.742	9.161	637
7.0%	1232	11	0	GET	^/api/images/[^/]+	308.794	0.251	1.658	2.313	149789
6.0%	444	5	0	POST	/api/friend-requests	266.061	0.599	4.684	5.260	114
5.4%	370	0	0	POST	/api/auth/login	239.096	0.646	4.362	8.152	249
3.9%	201	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	173.667	0.864	5.153	6.666	116
3.3%	852	0	0	GET	^/api/feeds/[^/]+	144.841	0.170	0.600	0.744	1783
2.3%	359	0	0	GET	/api/memories	101.945	0.284	1.078	1.702	14372
1.5%	61	8	0	POST	^/api/posts/[^/]+/realmojis$	66.737	1.094	7.410	7.410	455
1.4%	349	0	0	GET	/api/posting-window/current	63.296	0.181	0.483	0.582	239
1.1%	39	5	0	POST	^/api/posts/[^/]+/comments$	46.977	1.205	4.571	4.571	297
0.9%	200	9	0	GET	^/api/posts/me/current$	40.797	0.204	0.645	0.647	626
0.8%	233	0	0	GET	/api/users	37.312	0.160	0.449	0.454	252
0.8%	58	0	0	POST	/api/auth/logout	35.089	0.605	3.550	3.550	0
0.8%	79	0	0	POST	/api/auth/register	34.805	0.441	1.819	1.819	227
0.6%	47	0	0	PUT	/api/me/profile-image	25.044	0.533	2.440	2.440	0
0.6%	16	0	0	GET	/api/me/export	24.299	1.519	3.990	3.990	2085937
0.5%	137	0	0	GET	^/api/relationships/[^/]+	21.941	0.160	0.626	1.017	1068
0.5%	21	1	0	PATCH	^/api/posts/[^/]+$	21.893	1.043	5.848	5.848	631
0.4%	108	5	0	GET	^/api/posts/[^/]+$	17.669	0.164	0.458	0.611	602
0.3%	37	0	0	PATCH	/api/me	15.402	0.416	1.757	1.757	300
0.3%	42	0	0	PUT	^/api/follows/[^/]+$	13.967	0.333	1.764	1.764	0
0.3%	58	0	0	GET	^/api/memories/[^/]+	12.924	0.223	0.885	0.885	1190
0.2%	59	0	0	GET	^/api/users/[^/]+$	9.237	0.157	0.517	0.517	207
0.2%	11	0	0	DELETE	^/api/posts/[^/]+$	9.185	0.835	3.616	3.616	0
0.2%	51	1	0	GET	/api/me	6.907	0.135	0.228	0.228	292
0.2%	45	0	0	GET	^/api/posts/[^/]+/realmojis$	6.865	0.153	0.405	0.405	2113
0.1%	11	0	0	DELETE	^/api/realmojis/[^/]+$	6.308	0.573	2.005	2.005	0
0.1%	30	0	0	GET	^/api/posts/[^/]+/comments$	3.912	0.130	0.192	0.192	1343
0.1%	4	0	0	DELETE	^/api/comments/[^/]+$	2.316	0.579	1.288	1.288	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.1%	194	1	0	GET	/api/notifications	1485.433	7.657	9.366	9.366	178
18.5%	239	7	0	POST	/api/posts	685.446	2.868	7.731	8.047	623
7.8%	385	4	0	POST	/api/friend-requests	289.291	0.751	4.528	8.727	114
7.4%	1079	10	0	GET	^/api/images/[^/]+	273.924	0.254	2.184	2.774	172219
5.0%	289	0	0	POST	/api/auth/login	184.361	0.638	6.325	7.882	249
3.4%	693	0	0	GET	^/api/feeds/[^/]+	125.109	0.181	0.587	0.698	1849
3.3%	161	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	123.485	0.767	4.737	4.900	116
2.2%	298	0	0	GET	/api/memories	83.269	0.279	0.994	1.561	12220
1.5%	54	5	0	POST	^/api/posts/[^/]+/realmojis$	55.503	1.028	5.262	5.262	420
1.5%	317	0	0	GET	/api/posting-window/current	55.279	0.174	0.478	0.561	240
1.1%	17	0	0	GET	/api/me/export	39.042	2.297	7.809	7.809	2523143
1.0%	188	19	0	GET	^/api/posts/me/current$	37.937	0.202	0.614	0.645	576
0.8%	55	0	0	POST	/api/auth/logout	29.287	0.532	3.338	3.338	0
0.8%	175	0	0	GET	/api/users	29.134	0.166	0.483	0.637	247
0.7%	20	2	0	POST	^/api/posts/[^/]+/comments$	26.124	1.306	4.832	4.832	257
0.7%	49	0	0	PUT	^/api/follows/[^/]+$	24.231	0.495	2.282	2.282	0
0.6%	67	0	0	POST	/api/auth/register	23.871	0.356	4.596	4.596	229
0.6%	22	3	0	PATCH	^/api/posts/[^/]+$	23.776	1.081	5.733	5.733	558
0.6%	33	0	0	PUT	/api/me/profile-image	22.939	0.695	5.378	5.378	0
0.5%	110	0	0	GET	^/api/relationships/[^/]+	18.280	0.166	0.500	0.536	1449
0.4%	87	10	0	GET	^/api/posts/[^/]+$	14.806	0.170	0.520	0.520	565
0.3%	26	0	0	PATCH	/api/me	11.361	0.437	2.470	2.470	265
0.2%	10	0	0	DELETE	^/api/posts/[^/]+$	8.418	0.842	1.924	1.924	0
0.2%	35	0	0	GET	^/api/memories/[^/]+	7.956	0.227	0.497	0.497	1181
0.2%	41	0	0	GET	^/api/users/[^/]+$	6.559	0.160	0.406	0.406	192
0.1%	40	0	0	GET	^/api/posts/[^/]+/realmojis$	5.548	0.139	0.354	0.354	929
0.1%	7	0	0	DELETE	^/api/realmojis/[^/]+$	5.428	0.775	2.790	2.790	0
0.1%	38	1	0	GET	/api/me	5.374	0.141	0.422	0.422	287
0.1%	26	1	0	GET	^/api/posts/[^/]+/comments$	3.777	0.145	0.302	0.302	1668
0.0%	3	0	0	PUT	^/api/blocks/[^/]+$	1.346	0.449	1.111	1.111	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 68003 クエリ / 81 種類 / 合計 387.578s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	55.2%	3984	214.112	53.74	458.42	606.8	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	26.5%	283	102.633	362.66	1931.16	1265.7	128.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	5.8%	8833	22.655	2.56	113.49	10.9	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	5.8%	2046	22.508	11.00	655.29	116.5	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	1.8%	2713	6.862	2.53	23.48	25.1	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.3%	2	5.075	2537.73	4910.68	228907.5	2695.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.6%	655	2.239	3.42	19.78	28.3	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
8	0.6%	1	2.151	2150.59	2150.59	30498.0	1199795.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
9	0.5%	1	1.774	1773.93	1773.93	22814.0	627092.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
10	0.3%	796	1.346	1.69	18.05	20.8	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.3%	5920	1.201	0.20	10.41	19.6	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
12	0.2%	1	0.752	751.60	751.60	12160.0	530039.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
13	0.2%	2334	0.688	0.29	13.53	20.7	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
14	0.2%	4809	0.601	0.12	10.13	11.2	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
15	0.1%	2	0.268	133.79	257.43	16776.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 3
SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 5ef7e0ba89049eaa570e417d38efe43a95adeaab
Type: cpu
Time: 2026-10-31 13:09:57 JST
Duration: 60.03s, Total samples = 42.84s (71.37%)
Showing nodes accounting for 30.94s, 72.22% of 42.84s total
Dropped 1015 nodes (cum <= 0.21s)
Showing top 40 nodes out of 290
      flat  flat%   sum%        cum   cum%
     4.04s  9.43%  9.43%      4.04s  9.43%  internal/runtime/syscall/linux.Syscall6
     2.54s  5.93% 15.36%      8.14s 19.00%  runtime.scanObject
     2.10s  4.90% 20.26%      2.10s  4.90%  runtime.memmove
     2.05s  4.79% 25.05%      2.05s  4.79%  crypto/internal/fips140/sha256.blockAVX2
     1.97s  4.60% 29.65%      2.53s  5.91%  runtime.findObject
     1.95s  4.55% 34.20%      1.95s  4.55%  runtime.memclrNoHeapPointers
     1.84s  4.30% 38.49%      5.79s 13.52%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.59s  3.71% 42.20%      2.60s  6.07%  runtime.tryDeferToSpanScan
     1.32s  3.08% 45.28%      1.32s  3.08%  aeshashbody
     1.31s  3.06% 48.34%      1.31s  3.06%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.96s  2.24% 50.58%      2.87s  6.70%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.87s  2.03% 52.61%      0.87s  2.03%  example.com/isureal-company/isureal.(*UserRow).active
     0.72s  1.68% 54.30%      1.29s  3.01%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.72s  1.68% 55.98%      0.72s  1.68%  memeqbody
     0.72s  1.68% 57.66%      0.72s  1.68%  runtime.spanClass.sizeclass (inline)
     0.55s  1.28% 58.94%      0.61s  1.42%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.51s  1.19% 60.13%      0.63s  1.47%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.48s  1.12% 61.25%      1.91s  4.46%  runtime.mapaccess1_faststr
     0.41s  0.96% 62.21%      0.41s  0.96%  runtime.nextFreeFast (inline)
     0.40s  0.93% 63.14%      0.40s  0.93%  indexbody
     0.32s  0.75% 63.89%      0.32s  0.75%  runtime.futex
     0.31s  0.72% 64.61%      0.52s  1.21%  example.com/isureal-company/isureal.postLess
     0.26s  0.61% 65.22%      0.30s   0.7%  runtime.(*mspan).markBitsForIndex (inline)
     0.25s  0.58% 65.80%      0.29s  0.68%  runtime.spanOf (inline)
     0.24s  0.56% 66.36%      0.64s  1.49%  runtime.greyobject
     0.24s  0.56% 66.92%      0.24s  0.56%  sync/atomic.(*Int32).Add
     0.22s  0.51% 67.44%      1.15s  2.68%  runtime.mallocgcSmallScanNoHeader
     0.22s  0.51% 67.95%      0.25s  0.58%  runtime.typePointers.next
     0.19s  0.44% 68.39%      0.39s  0.91%  compress/flate.(*compressor).deflate
     0.19s  0.44% 68.84%      0.68s  1.59%  database/sql.convertAssignRows
     0.19s  0.44% 69.28%      1.92s  4.48%  runtime.mapassign_faststr
     0.16s  0.37% 69.65%      5.47s 12.77%  runtime.mallocgc
     0.15s  0.35% 70.00%      0.40s  0.93%  runtime.scanObjectsSmall
     0.14s  0.33% 70.33%      1.97s  4.60%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.14s  0.33% 70.66%      1.38s  3.22%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.14s  0.33% 70.99%      0.23s  0.54%  runtime.heapArenaOf (inline)
     0.14s  0.33% 71.31%      0.83s  1.94%  runtime.mallocgcSmallNoscan
     0.13s   0.3% 71.62%      2.36s  5.51%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.13s   0.3% 71.92%      0.30s   0.7%  internal/runtime/maps.(*Iter).Next
     0.13s   0.3% 72.22%      0.73s  1.70%  runtime.scanSpan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 53.4% / 100% | 37% | 0.2% | 8 | 181 MB | CPU が飽和 |
| isu2 (2 cpu) | 31.7% / 91% | 2% | 4.9% | 3 | 412 MB |  |
| isu3 (2 cpu) | 18.5% / 64% | 0% | 0.4% | 4 | 2532 MB |  |
| isu4 (2 cpu) | 59.3% / 100% | 29% | 21.9% | 35 | 104 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 79% / nginx 17% / journalctl 3% / systemd-journal 2% |
| isu2 | isureal 41% / du 10% |
| isu3 | isureal 35% / nginx 1% |
| isu4 | postgres 67% / isureal 4% |

## Diff from previous run

スコア: 85230 → 110454 (+25224)

`8caa076..c75a0a9`

```
 webapp/go/config.go     | 2 ++
 webapp/go/export_api.go | 4 +++-
 2 files changed, 5 insertions(+), 1 deletion(-)
```
