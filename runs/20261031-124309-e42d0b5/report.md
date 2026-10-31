# run 20261031-124309-e42d0b5

## Summary

| | |
|---|---|
| score | 82684 |
| status | pass |
| ref | grandprix-post-fifo |
| sha | e42d0b56cef8 |
| 所要時間 | 232s |
| 前回比 | +17984 (27.8%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +17984 (27.8%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | PR73 merged with main 04641dc |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	145	407	620	620
europe	7	15	39	92	93	93
america	7	13	32	59	59	57
```

```
{"id":"9486","status":"done","portalStatus":"finished","passed":true,"score":82684,"raw":82684,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:45:29.384Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:45:30.876Z level=INFO msg="user moved" user=numax_5069 from=asia to=america
time=2026-10-31T03:45:32.348Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:45:35.347Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T03:45:38.375Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T03:45:41.347Z level=INFO msg="europe 2 日目: 15 人"
time=2026-10-31T03:45:44.348Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T03:45:47.378Z level=INFO msg="asia 3 日目: 145 人"
time=2026-10-31T03:45:49.823Z level=INFO msg="user moved" user=hiannoa-5520 from=asia to=europe
time=2026-10-31T03:45:50.376Z level=INFO msg="europe 3 日目: 39 人"
time=2026-10-31T03:45:50.585Z level=INFO msg="user moved" user=numax_5069 from=america to=asia
time=2026-10-31T03:45:51.046Z level=INFO msg="user moved" user=numax_5069 from=asia to=america
time=2026-10-31T03:45:53.347Z level=INFO msg="america 3 日目: 32 人"
time=2026-10-31T03:45:56.442Z level=INFO msg="asia 4 日目: 407 人"
time=2026-10-31T03:45:59.540Z level=INFO msg="europe 4 日目: 92 人"
time=2026-10-31T03:46:02.521Z level=INFO msg="user moved" user=namahi-6599 from=asia to=europe
time=2026-10-31T03:46:02.523Z level=INFO msg="america 4 日目: 59 人"
time=2026-10-31T03:46:05.875Z level=INFO msg="asia 5 日目: 620 人"
time=2026-10-31T03:46:08.669Z level=INFO msg="europe 5 日目: 93 人"
time=2026-10-31T03:46:11.853Z level=INFO msg="america 5 日目: 59 人"
time=2026-10-31T03:46:12.974Z level=INFO msg="user moved" user=numax_5069 from=america to=asia
time=2026-10-31T03:46:20.586Z level=INFO msg="america 最終日（グランプリ）: 57 人"
time=2026-10-31T03:46:20.901Z level=INFO msg="europe 最終日（グランプリ）: 93 人"
time=2026-10-31T03:46:21.303Z level=INFO msg="asia 最終日（グランプリ）: 620 人"
time=2026-10-31T03:46:29.364Z level=INFO msg="グランプリ: アジア 248 人, ヨーロッパ 29 人, アメリカ 26 人"
{
  "pass": true,
  "score": 82684,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.4%	1201	0	0	GET	/api/notifications	10698.262	8.908	15.492	15.764	171
17.3%	1723	36	0	POST	/api/posts	4959.241	2.878	9.344	10.534	638
10.5%	3589	59	0	POST	/api/friend-requests	2989.919	0.833	5.017	8.378	114
8.4%	11008	602	0	GET	^/api/images/[^/]+	2405.325	0.219	3.334	9.859	139842
5.7%	4453	1241	0	POST	/api/auth/login	1644.624	0.369	4.128	8.538	184
4.4%	1351	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1260.808	0.933	5.535	7.605	116
2.3%	494	135	0	PUT	/api/me/profile-image	655.481	1.327	7.090	8.347	5
2.0%	559	8	0	POST	^/api/posts/[^/]+/comments$	577.199	1.033	6.124	9.856	331
2.0%	402	9	0	POST	^/api/posts/[^/]+/realmojis$	576.098	1.433	6.889	7.662	466
1.6%	639	1	0	POST	/api/auth/register	470.185	0.736	4.126	8.120	227
1.1%	405	0	0	PUT	^/api/follows/[^/]+$	326.365	0.806	4.575	6.002	0
1.1%	8992	0	0	GET	^/api/feeds/[^/]+	325.050	0.036	0.214	0.990	1943
1.0%	2837	0	0	GET	/api/memories	296.816	0.105	0.452	1.485	10654
0.8%	280	0	0	PATCH	/api/me	234.221	0.837	5.931	7.564	276
0.6%	169	5	0	PATCH	^/api/posts/[^/]+$	181.608	1.075	5.295	5.399	628
0.6%	188	0	0	GET	/api/me/export	162.224	0.863	4.746	6.214	1920921
0.5%	2414	0	0	GET	/api/posting-window/current	153.476	0.064	0.561	1.137	237
0.5%	571	0	0	POST	/api/auth/logout	145.881	0.255	2.331	3.209	0
0.3%	68	0	0	DELETE	^/api/posts/[^/]+$	93.213	1.371	3.863	3.863	0
0.2%	1367	60	0	GET	^/api/posts/me/current$	64.992	0.048	0.382	0.657	623
0.2%	1719	0	0	GET	/api/users	62.914	0.037	0.184	0.780	243
0.2%	75	0	0	DELETE	^/api/comments/[^/]+$	48.224	0.643	4.070	4.070	0
0.2%	62	0	0	DELETE	^/api/realmojis/[^/]+$	45.478	0.734	4.483	4.483	0
0.1%	1233	0	0	GET	^/api/relationships/[^/]+	41.553	0.034	0.239	1.574	960
0.1%	1248	20	0	GET	^/api/posts/[^/]+$	39.770	0.032	0.180	0.337	624
0.1%	493	0	0	GET	^/api/memories/[^/]+	29.004	0.059	0.357	0.982	655
0.1%	693	0	0	GET	^/api/users/[^/]+$	22.510	0.032	0.265	0.624	200
0.1%	33	0	0	PUT	^/api/blocks/[^/]+$	17.655	0.535	2.572	2.572	0
0.1%	1	0	0	POST	/api/initialize	17.477	17.477	17.477	17.477	90
0.1%	536	0	0	GET	^/api/posts/[^/]+/realmojis$	15.768	0.029	0.158	0.269	858
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.0%	226	0	0	GET	/api/notifications	1937.561	8.573	12.286	12.287	179
19.3%	287	3	0	POST	/api/posts	958.795	3.341	9.799	10.607	636
9.5%	515	9	0	POST	/api/friend-requests	470.742	0.914	4.861	10.170	114
6.1%	377	0	0	POST	/api/auth/login	304.153	0.807	4.816	6.601	249
4.6%	1425	11	0	GET	^/api/images/[^/]+	226.415	0.159	0.948	1.656	161239
4.5%	224	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	221.638	0.989	4.141	5.184	116
3.4%	1061	0	0	GET	^/api/feeds/[^/]+	170.705	0.161	0.424	1.553	2133
2.0%	412	0	0	GET	/api/memories	97.346	0.236	0.671	1.082	11750
1.3%	390	0	0	GET	/api/posting-window/current	63.903	0.164	0.654	1.025	239
1.2%	58	0	0	POST	/api/auth/logout	57.576	0.993	5.469	5.469	0
1.0%	89	0	0	POST	/api/auth/register	51.929	0.583	2.246	2.246	228
1.0%	50	1	0	POST	^/api/posts/[^/]+/realmojis$	50.046	1.001	4.947	4.947	491
0.8%	63	0	0	PUT	^/api/follows/[^/]+$	38.736	0.615	3.688	3.688	0
0.8%	241	0	0	GET	/api/users	37.374	0.155	0.351	0.413	238
0.7%	227	11	0	GET	^/api/posts/me/current$	36.601	0.161	0.465	0.724	613
0.7%	17	0	0	GET	/api/me/export	35.824	2.107	5.775	5.775	2542565
0.7%	47	0	0	PUT	/api/me/profile-image	33.205	0.706	3.729	3.729	0
0.6%	43	0	0	PATCH	/api/me	28.380	0.660	2.501	2.501	272
0.5%	160	0	0	GET	^/api/relationships/[^/]+	22.917	0.143	0.338	0.470	1462
0.5%	142	6	0	GET	^/api/posts/[^/]+$	22.602	0.159	0.345	0.447	619
0.4%	19	1	0	PATCH	^/api/posts/[^/]+$	19.302	1.016	4.076	4.076	602
0.3%	13	0	0	DELETE	^/api/posts/[^/]+$	16.817	1.294	4.685	4.685	0
0.3%	13	0	0	POST	^/api/posts/[^/]+/comments$	15.482	1.191	3.194	3.194	348
0.2%	76	0	0	GET	^/api/users/[^/]+$	11.685	0.154	0.321	0.321	203
0.2%	72	0	0	GET	^/api/posts/[^/]+/realmojis$	10.594	0.147	0.346	0.346	1059
0.2%	42	0	0	GET	^/api/memories/[^/]+	8.492	0.202	0.510	0.510	888
0.1%	50	1	0	GET	/api/me	7.060	0.141	0.308	0.308	293
0.1%	33	0	0	GET	^/api/posts/[^/]+/comments$	5.544	0.168	0.293	0.293	1760
0.1%	7	0	0	DELETE	^/api/realmojis/[^/]+$	3.857	0.551	1.498	1.498	0
0.0%	3	0	0	PUT	^/api/blocks/[^/]+$	1.121	0.374	0.892	0.892	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.9%	152	0	0	GET	/api/notifications	1124.582	7.399	9.451	9.451	178
15.9%	174	4	0	POST	/api/posts	482.951	2.776	9.718	10.543	636
10.1%	333	6	0	POST	/api/friend-requests	306.470	0.920	5.511	7.065	113
6.7%	245	0	0	POST	/api/auth/login	204.914	0.836	6.079	9.747	252
5.4%	1023	4	0	GET	^/api/images/[^/]+	165.345	0.162	1.034	1.464	167914
4.8%	131	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	145.318	1.109	4.203	7.095	115
4.1%	753	0	0	GET	^/api/feeds/[^/]+	124.425	0.165	0.679	2.947	2161
2.0%	256	0	0	GET	/api/memories	62.110	0.243	0.710	1.281	14961
1.4%	32	0	0	PUT	/api/me/profile-image	44.149	1.380	8.283	8.283	0
1.3%	57	0	0	POST	/api/auth/register	38.140	0.669	4.332	4.332	229
1.2%	233	0	0	GET	/api/posting-window/current	37.859	0.162	0.646	1.172	239
1.0%	13	0	0	GET	/api/me/export	31.042	2.388	4.030	4.030	3675094
1.0%	30	0	0	POST	/api/auth/logout	30.416	1.014	6.680	6.680	0
1.0%	28	0	0	PATCH	/api/me	29.633	1.058	3.817	3.817	273
0.8%	27	2	0	POST	^/api/posts/[^/]+/realmojis$	25.461	0.943	7.319	7.319	432
0.8%	156	0	0	GET	/api/users	25.056	0.161	0.413	1.078	232
0.8%	146	9	0	GET	^/api/posts/me/current$	24.344	0.167	0.893	1.009	610
0.7%	20	1	0	PATCH	^/api/posts/[^/]+$	20.836	1.042	4.678	4.678	631
0.6%	33	0	0	PUT	^/api/follows/[^/]+$	19.066	0.578	3.729	3.729	0
0.6%	117	0	0	GET	^/api/relationships/[^/]+	18.999	0.162	0.709	1.026	1022
0.6%	116	10	0	GET	^/api/posts/[^/]+$	17.468	0.151	0.300	0.312	579
0.5%	22	0	0	POST	^/api/posts/[^/]+/comments$	16.219	0.737	4.533	4.533	323
0.3%	69	0	0	GET	^/api/users/[^/]+$	10.657	0.154	0.450	0.450	202
0.3%	6	0	0	DELETE	^/api/posts/[^/]+$	7.842	1.307	2.748	2.748	0
0.3%	42	0	0	GET	^/api/memories/[^/]+	7.714	0.184	0.380	0.380	833
0.2%	46	1	0	GET	^/api/posts/[^/]+/realmojis$	7.113	0.155	0.418	0.418	710
0.2%	36	1	0	GET	/api/me	5.639	0.157	0.331	0.331	293
0.2%	36	1	0	GET	^/api/posts/[^/]+/comments$	5.111	0.142	0.211	0.211	1523
0.1%	5	0	0	DELETE	^/api/realmojis/[^/]+$	2.718	0.544	1.708	1.708	0
0.1%	1	0	0	DELETE	^/api/comments/[^/]+$	1.589	1.589	1.589	1.589	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 53519 クエリ / 82 種類 / 合計 257.185s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	52.2%	3049	134.202	44.02	453.61	622.6	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	27.7%	218	71.358	327.33	1161.01	6898.0	132.7	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	6.0%	7286	15.415	2.12	58.41	10.8	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	5.5%	2059	14.112	6.85	257.26	116.5	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	2.0%	2143	5.217	2.43	71.09	25.3	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.9%	2	4.762	2380.84	4605.70	229957.5	2707.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.9%	1	2.228	2227.78	2227.78	30076.0	1199701.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.7%	1	1.723	1723.45	1723.45	21742.0	626943.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.6%	467	1.446	3.10	18.54	29.7	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.4%	586	0.904	1.54	9.74	21.6	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.3%	4350	0.887	0.20	25.92	19.8	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
12	0.3%	2	0.815	407.64	806.88	16757.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
13	0.3%	1	0.751	751.28	751.28	11774.0	529937.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
14	0.2%	1655	0.419	0.25	23.43	20.7	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.2%	3834	0.410	0.11	7.24	11.1	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 914bfaf5d2c97645b53090fbd8ce1b981377fb49
Type: cpu
Time: 2026-10-31 12:45:06 JST
Duration: 60.06s, Total samples = 43.20s (71.93%)
Showing nodes accounting for 32.68s, 75.65% of 43.20s total
Dropped 952 nodes (cum <= 0.22s)
Showing top 40 nodes out of 280
      flat  flat%   sum%        cum   cum%
     3.11s  7.20%  7.20%      3.11s  7.20%  internal/runtime/syscall/linux.Syscall6
     2.60s  6.02% 13.22%      2.60s  6.02%  runtime.memclrNoHeapPointers
     2.46s  5.69% 18.91%      8.65s 20.02%  runtime.scanObject
     2.12s  4.91% 23.82%      2.74s  6.34%  runtime.findObject
     1.96s  4.54% 28.36%      1.96s  4.54%  runtime.memmove
     1.86s  4.31% 32.66%      2.97s  6.88%  runtime.tryDeferToSpanScan
     1.70s  3.94% 36.60%      5.13s 11.88%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.69s  3.91% 40.51%      1.69s  3.91%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.63s  3.77% 44.28%      1.63s  3.77%  crypto/internal/fips140/sha256.blockAVX2
     1.52s  3.52% 47.80%      5.52s 12.78%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.32s  3.06% 50.86%      2.16s  5.00%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.23s  2.85% 53.70%      1.27s  2.94%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     1.22s  2.82% 56.53%      1.22s  2.82%  aeshashbody
     0.77s  1.78% 58.31%      0.77s  1.78%  example.com/isureal-company/isureal.(*UserRow).active
     0.76s  1.76% 60.07%      0.76s  1.76%  runtime.spanClass.sizeclass (inline)
     0.63s  1.46% 61.53%      0.87s  2.01%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.48s  1.11% 62.64%      0.48s  1.11%  memeqbody
     0.47s  1.09% 63.73%      1.94s  4.49%  runtime.mapaccess1_faststr
     0.36s  0.83% 64.56%      0.36s  0.83%  runtime.nextFreeFast (inline)
     0.34s  0.79% 65.35%      0.34s  0.79%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.33s  0.76% 66.11%      0.56s  1.30%  example.com/isureal-company/isureal.postLess
     0.33s  0.76% 66.88%      0.68s  1.57%  runtime.greyobject
     0.33s  0.76% 67.64%      0.33s  0.76%  runtime.typePointers.nextFast (inline)
     0.29s  0.67% 68.31%      0.39s   0.9%  runtime.typePointers.next
     0.28s  0.65% 68.96%      0.66s  1.53%  database/sql.convertAssignRows
     0.27s  0.62% 69.58%      0.27s  0.62%  runtime.futex
     0.23s  0.53% 70.12%      0.23s  0.53%  indexbody
     0.23s  0.53% 70.65%      0.23s  0.53%  runtime.(*mspan).divideByElemSize (inline)
     0.22s  0.51% 71.16%      0.24s  0.56%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.22s  0.51% 71.67%      0.22s  0.51%  sync/atomic.(*Int32).Add (inline)
     0.21s  0.49% 72.15%      0.91s  2.11%  runtime.mallocgcSmallScanNoHeader
     0.21s  0.49% 72.64%      0.25s  0.58%  runtime.spanOf (inline)
     0.18s  0.42% 73.06%      0.26s   0.6%  internal/runtime/maps.(*Iter).Next
     0.18s  0.42% 73.47%      1.90s  4.40%  runtime.mapassign_faststr
     0.17s  0.39% 73.87%      0.93s  2.15%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.17s  0.39% 74.26%      0.24s  0.56%  runtime.heapArenaOf (inline)
     0.16s  0.37% 74.63%      2.71s  6.27%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.16s  0.37% 75.00%      0.63s  1.46%  runtime.scanObjectsSmall
     0.15s  0.35% 75.35%      1.86s  4.31%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.13s   0.3% 75.65%      3.87s  8.96%  example.com/isureal-company/isureal.loadBig
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 47.4% / 100% | 35% | 0.1% | 6 | 123 MB | CPU が飽和 |
| isu2 (2 cpu) | 14.2% / 49% | 0% | 0.4% | 5 | 2593 MB | run queue が積もっている |
| isu3 (2 cpu) | 10.7% / 48% | 0% | 0.4% | 4 | 2685 MB |  |
| isu4 (2 cpu) | 47.4% / 100% | 11% | 20.1% | 25 | 115 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 75% / nginx 13% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 27% |
| isu3 | isureal 19% / nginx 1% |
| isu4 | postgres 50% / isureal 1% |

## Diff from previous run

スコア: 64700 → 82684 (+17984)

`e0e4cb3..e42d0b5`

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
