# run 20261031-123155-fdbe826

## Summary

| | |
|---|---|
| score | 65037 |
| status | pass |
| ref | origin/main |
| sha | fdbe82670a24 |
| 所要時間 | 233s |
| 前回比 | +59815 (1145.4%) vs 20261031-122906-3d5209e |
| ベスト比 | -1680 (-2.5%) vs 20261031-122401-fdbe826 = 66717 |
| note | 低い版の試走の直後の main |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	145	407	405	404
europe	8	17	43	93	93	92
america	6	13	29	40	43	42
```

```
{"id":"8649","status":"done","portalStatus":"finished","passed":true,"score":65037,"raw":65037,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:34:22.375Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:34:23.781Z level=INFO msg="user moved" user=muna5039 from=asia to=europe
time=2026-10-31T03:34:25.334Z level=INFO msg="europe 1 日目: 8 人"
time=2026-10-31T03:34:28.334Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:34:31.361Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T03:34:34.342Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T03:34:36.761Z level=INFO msg="user moved" user=muna5039 from=europe to=america
time=2026-10-31T03:34:37.333Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T03:34:38.711Z level=INFO msg="user moved" user=zoeel.5408 from=europe to=asia
time=2026-10-31T03:34:40.366Z level=INFO msg="asia 3 日目: 145 人"
time=2026-10-31T03:34:41.377Z level=INFO msg="user moved" user=shisovin-5498 from=asia to=america
time=2026-10-31T03:34:41.723Z level=INFO msg="user moved" user=zoeel.5408 from=asia to=europe
time=2026-10-31T03:34:43.140Z level=INFO msg="user moved" user=marka.5750 from=asia to=america
time=2026-10-31T03:34:43.430Z level=INFO msg="europe 3 日目: 43 人"
time=2026-10-31T03:34:44.351Z level=INFO msg="user moved" user=liaan_6026 from=asia to=europe
time=2026-10-31T03:34:44.880Z level=INFO msg="user moved" user=marka.5750 from=america to=europe
time=2026-10-31T03:34:45.596Z level=INFO msg="user moved" user=lianotsu5478 from=america to=europe
time=2026-10-31T03:34:46.456Z level=INFO msg="america 3 日目: 29 人"
time=2026-10-31T03:34:48.757Z level=INFO msg="user moved" user=lianotsu5478 from=europe to=asia
time=2026-10-31T03:34:49.383Z level=INFO msg="asia 4 日目: 407 人"
time=2026-10-31T03:34:51.487Z level=INFO msg="user moved" user=haseor6342 from=europe to=asia
time=2026-10-31T03:34:52.377Z level=INFO msg="europe 4 日目: 93 人"
time=2026-10-31T03:34:54.437Z level=INFO msg="user moved" user=haseor6342 from=asia to=america
time=2026-10-31T03:34:55.401Z level=INFO msg="america 4 日目: 40 人"
time=2026-10-31T03:34:55.402Z level=INFO msg="user moved" user=lianotsu5478 from=asia to=america
time=2026-10-31T03:34:58.360Z level=INFO msg="user moved" user=take-8350 from=asia to=europe
time=2026-10-31T03:34:58.521Z level=INFO msg="asia 5 日目: 405 人"
time=2026-10-31T03:34:58.836Z level=INFO msg="user moved" user=zoeel.5408 from=europe to=asia
time=2026-10-31T03:34:58.952Z level=INFO msg="user moved" user=naku.6753 from=asia to=america
time=2026-10-31T03:35:01.821Z level=INFO msg="europe 5 日目: 93 人"
time=2026-10-31T03:35:03.533Z level=INFO msg="user moved" user=liaan_6026 from=europe to=america
time=2026-10-31T03:35:04.528Z level=INFO msg="america 5 日目: 43 人"
time=2026-10-31T03:35:13.468Z level=INFO msg="asia 最終日（グランプリ）: 404 人"
time=2026-10-31T03:35:13.516Z level=INFO msg="america 最終日（グランプリ）: 42 人"
time=2026-10-31T03:35:13.571Z level=INFO msg="europe 最終日（グランプリ）: 92 人"
time=2026-10-31T03:35:22.343Z level=INFO msg="グランプリ: アジア 30 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 65037,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.7%	940	0	0	GET	/api/notifications	8806.554	9.369	14.955	15.033	173
22.0%	1267	42	0	POST	/api/posts	5269.946	4.159	14.025	14.786	629
9.1%	9097	679	0	GET	^/api/images/[^/]+	2171.336	0.239	3.474	14.903	125204
7.1%	2475	68	0	POST	/api/friend-requests	1711.595	0.692	5.147	10.375	113
6.9%	3481	899	0	POST	/api/auth/login	1646.762	0.473	7.208	13.845	189
3.9%	1434	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	929.450	0.648	5.855	13.935	116
1.6%	426	1	0	POST	/api/auth/register	392.838	0.922	4.298	12.871	227
1.6%	359	130	0	PUT	/api/me/profile-image	382.090	1.064	8.413	9.529	8
1.5%	1719	0	0	GET	/api/posting-window/current	368.898	0.215	3.531	7.070	237
1.2%	2330	0	0	GET	/api/memories	299.157	0.128	0.979	5.167	10877
1.2%	369	6	0	POST	^/api/posts/[^/]+/realmojis$	297.141	0.805	3.813	5.385	479
1.1%	7349	0	0	GET	^/api/feeds/[^/]+	267.037	0.036	0.242	6.002	2629
1.1%	409	5	0	POST	^/api/posts/[^/]+/comments$	253.813	0.621	4.266	8.317	349
0.9%	260	0	0	PUT	^/api/follows/[^/]+$	213.244	0.820	6.116	7.079	0
0.7%	474	0	0	POST	/api/auth/logout	167.424	0.353	7.870	12.680	0
0.6%	147	0	0	GET	/api/me/export	141.405	0.962	11.209	12.091	2086550
0.6%	226	0	0	PATCH	/api/me	136.912	0.606	5.149	7.217	272
0.3%	102	5	0	PATCH	^/api/posts/[^/]+$	81.422	0.798	5.874	5.900	622
0.3%	917	48	0	GET	^/api/posts/me/current$	73.110	0.080	1.067	5.201	617
0.3%	47	0	0	DELETE	^/api/posts/[^/]+$	68.289	1.453	10.846	10.846	0
0.2%	1333	12	0	GET	^/api/posts/[^/]+$	41.148	0.031	0.185	0.449	626
0.2%	1200	0	0	GET	/api/users	39.959	0.033	0.227	0.927	242
0.1%	1091	0	0	GET	^/api/relationships/[^/]+	31.319	0.029	0.240	1.193	1211
0.1%	58	0	0	DELETE	^/api/realmojis/[^/]+$	29.893	0.515	7.984	7.984	0
0.1%	428	0	0	GET	^/api/memories/[^/]+	29.478	0.069	0.309	2.579	747
0.1%	93	0	0	DELETE	^/api/comments/[^/]+$	26.595	0.286	2.173	2.173	0
0.1%	1	0	0	POST	/api/initialize	19.682	19.682	19.682	19.682	90
0.1%	596	0	0	GET	^/api/users/[^/]+$	17.999	0.030	0.212	0.447	205
0.1%	559	0	0	GET	^/api/posts/[^/]+/realmojis$	17.337	0.031	0.150	0.339	1800
0.1%	463	0	0	GET	^/api/posts/[^/]+/comments$	13.644	0.029	0.163	0.278	1886
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.8%	227	0	0	GET	/api/notifications	1907.194	8.402	11.758	12.284	177
23.5%	293	5	0	POST	/api/posts	1289.147	4.400	13.995	14.275	643
7.8%	555	7	0	POST	/api/friend-requests	429.076	0.773	5.465	11.978	114
7.0%	389	0	0	POST	/api/auth/login	381.058	0.980	12.827	14.067	249
4.1%	1481	15	0	GET	^/api/images/[^/]+	223.886	0.151	0.990	1.654	173505
3.8%	229	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	206.617	0.902	11.829	14.263	116
3.6%	1223	0	0	GET	^/api/feeds/[^/]+	197.715	0.162	0.517	1.159	2606
2.5%	393	0	0	GET	/api/posting-window/current	134.411	0.342	3.835	7.759	239
1.6%	400	0	0	GET	/api/memories	88.160	0.220	0.618	1.114	12452
1.1%	60	0	0	POST	/api/auth/logout	60.748	1.012	12.901	12.901	0
1.1%	54	0	0	PUT	/api/me/profile-image	59.279	1.098	7.931	7.931	0
0.8%	90	0	0	POST	/api/auth/register	45.749	0.508	2.843	2.843	228
0.8%	51	0	0	PATCH	/api/me	43.935	0.861	4.544	4.544	279
0.8%	47	0	0	PUT	^/api/follows/[^/]+$	42.107	0.896	4.328	4.328	0
0.8%	64	2	0	POST	^/api/posts/[^/]+/realmojis$	42.100	0.658	3.590	3.590	462
0.8%	28	0	0	GET	/api/me/export	41.670	1.488	7.362	7.362	1866331
0.8%	257	0	0	GET	/api/users	41.296	0.161	0.426	0.535	248
0.6%	208	13	0	GET	^/api/posts/[^/]+$	34.111	0.164	0.519	0.638	592
0.6%	210	6	0	GET	^/api/posts/me/current$	32.457	0.155	0.323	0.356	634
0.6%	64	1	0	POST	^/api/posts/[^/]+/comments$	30.388	0.475	2.703	2.703	304
0.5%	194	0	0	GET	^/api/relationships/[^/]+	29.904	0.154	0.517	0.639	1299
0.4%	23	0	0	PATCH	^/api/posts/[^/]+$	23.560	1.024	3.323	3.323	660
0.3%	95	0	0	GET	^/api/users/[^/]+$	15.511	0.163	0.639	0.639	205
0.3%	15	0	0	DELETE	^/api/posts/[^/]+$	14.307	0.954	3.006	3.006	0
0.3%	87	0	0	GET	^/api/posts/[^/]+/realmojis$	14.250	0.164	0.639	0.639	1623
0.2%	76	0	0	GET	^/api/posts/[^/]+/comments$	12.488	0.164	0.541	0.541	2610
0.2%	58	1	0	GET	/api/me	9.482	0.163	0.436	0.436	295
0.2%	51	0	0	GET	^/api/memories/[^/]+	8.837	0.173	0.434	0.434	799
0.1%	3	0	0	DELETE	^/api/follows/[^/]+$	5.880	1.960	5.025	5.025	0
0.1%	7	0	0	DELETE	^/api/realmojis/[^/]+$	3.498	0.500	0.947	0.947	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.0%	129	0	0	GET	/api/notifications	985.687	7.641	9.205	9.242	176
21.9%	133	3	0	POST	/api/posts	552.315	4.153	13.265	13.287	645
7.2%	192	0	0	POST	/api/auth/login	180.869	0.942	12.971	14.380	254
6.9%	235	6	0	POST	/api/friend-requests	174.503	0.743	5.461	6.168	113
4.7%	647	4	0	GET	^/api/images/[^/]+	118.886	0.184	0.957	1.111	148876
4.3%	128	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	107.643	0.841	14.154	14.446	115
3.1%	188	0	0	GET	/api/posting-window/current	78.899	0.420	5.660	7.727	240
2.2%	363	0	0	GET	^/api/feeds/[^/]+	55.352	0.152	0.483	1.008	2785
1.7%	198	0	0	GET	/api/memories	43.618	0.220	0.831	2.098	11634
1.5%	19	0	0	GET	/api/me/export	38.189	2.010	8.497	8.497	2903745
1.5%	50	0	0	POST	/api/auth/logout	36.948	0.739	7.171	7.171	0
0.8%	20	0	0	PATCH	/api/me	19.602	0.980	4.003	4.003	291
0.7%	21	0	0	PUT	/api/me/profile-image	17.604	0.838	3.242	3.242	0
0.7%	106	0	0	GET	/api/users	17.603	0.166	0.503	0.662	242
0.6%	97	7	0	GET	^/api/posts/me/current$	15.516	0.160	0.715	0.715	611
0.5%	37	0	0	POST	/api/auth/register	13.780	0.372	1.854	1.854	230
0.4%	75	0	0	GET	^/api/relationships/[^/]+	10.398	0.139	0.453	0.453	1878
0.4%	29	2	0	POST	^/api/posts/[^/]+/comments$	9.287	0.320	0.735	0.735	317
0.4%	66	2	0	GET	^/api/posts/[^/]+$	9.043	0.137	0.242	0.242	622
0.3%	18	0	0	PUT	^/api/follows/[^/]+$	7.786	0.433	3.106	3.106	0
0.2%	8	1	0	PATCH	^/api/posts/[^/]+$	5.380	0.672	1.704	1.704	589
0.2%	26	0	0	GET	^/api/memories/[^/]+	5.198	0.200	1.008	1.008	804
0.2%	4	0	0	DELETE	^/api/posts/[^/]+$	4.171	1.043	2.513	2.513	0
0.1%	27	0	0	GET	^/api/posts/[^/]+/comments$	3.784	0.140	0.437	0.437	1989
0.1%	25	1	0	GET	/api/me	3.269	0.131	0.279	0.279	281
0.1%	21	0	0	GET	^/api/users/[^/]+$	2.849	0.136	0.207	0.207	214
0.1%	7	0	0	POST	^/api/posts/[^/]+/realmojis$	2.535	0.362	0.625	0.625	522
0.1%	2	0	0	DELETE	^/api/comments/[^/]+$	2.293	1.147	1.757	1.757	0
0.1%	11	0	0	GET	^/api/posts/[^/]+/realmojis$	1.454	0.132	0.180	0.180	748
0.0%	2	0	0	PUT	^/api/blocks/[^/]+$	0.596	0.298	0.488	0.488	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 45058 クエリ / 81 種類 / 合計 278.662s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	58.1%	2382	161.930	67.98	523.05	644.5	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	24.4%	194	67.958	350.30	1575.54	7212.2	136.5	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	4.7%	1789	13.214	7.39	265.03	116.4	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.7%	6110	12.981	2.12	25.33	10.5	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	2.2%	2	6.233	3116.67	6014.09	231973.5	2731.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
6	1.6%	1646	4.358	2.65	38.53	26.1	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
7	0.8%	1	2.331	2331.18	2331.18	29963.0	1199714.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.6%	1	1.784	1784.47	1784.47	21390.0	627289.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.5%	432	1.362	3.15	24.44	30.4	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.4%	2	0.991	495.39	982.25	16776.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.3%	494	0.813	1.65	12.69	21.2	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.3%	1	0.787	786.56	786.56	11661.0	530067.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
13	0.2%	3171	0.639	0.20	7.76	20.6	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
14	0.1%	1737	0.399	0.23	4.95	21.5	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	3163	0.382	0.12	15.02	11.8	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

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
Build ID: 93cebbb64943b8a20fb2be2ee828625a8aafe067
Type: cpu
Time: 2026-10-31 12:33:56 JST
Duration: 60.07s, Total samples = 41.54s (69.15%)
Showing nodes accounting for 31.68s, 76.26% of 41.54s total
Dropped 897 nodes (cum <= 0.21s)
Showing top 40 nodes out of 285
      flat  flat%   sum%        cum   cum%
     4.22s 10.16% 10.16%      4.22s 10.16%  crypto/internal/fips140/sha256.blockAVX2
     2.51s  6.04% 16.20%      2.51s  6.04%  internal/runtime/syscall/linux.Syscall6
     2.18s  5.25% 21.45%      2.18s  5.25%  runtime.memclrNoHeapPointers
     2.18s  5.25% 26.70%      2.18s  5.25%  runtime.memmove
     1.80s  4.33% 31.03%      2.22s  5.34%  runtime.findObject
     1.78s  4.29% 35.32%      5.93s 14.28%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.77s  4.26% 39.58%      6.62s 15.94%  runtime.scanObject
     1.39s  3.35% 42.92%      4.49s 10.81%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.39s  3.35% 46.27%      1.39s  3.35%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.23s  2.96% 49.23%      2.17s  5.22%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.21s  2.91% 52.14%      2.19s  5.27%  runtime.tryDeferToSpanScan
     1.12s  2.70% 54.84%      1.12s  2.70%  aeshashbody
     0.92s  2.21% 57.05%      0.93s  2.24%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.66s  1.59% 58.64%      0.66s  1.59%  example.com/isureal-company/isureal.(*UserRow).active
     0.66s  1.59% 60.23%      0.66s  1.59%  runtime.spanClass.sizeclass (inline)
     0.64s  1.54% 61.77%      0.94s  2.26%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.52s  1.25% 63.02%      2.01s  4.84%  runtime.mapaccess1_faststr
     0.49s  1.18% 64.20%      0.49s  1.18%  memeqbody
     0.36s  0.87% 65.07%      0.56s  1.35%  example.com/isureal-company/isureal.postLess
     0.35s  0.84% 65.91%      0.72s  1.73%  database/sql.convertAssignRows
     0.31s  0.75% 66.66%      0.60s  1.44%  runtime.greyobject
     0.29s   0.7% 67.36%      0.29s   0.7%  runtime.nextFreeFast (inline)
     0.28s  0.67% 68.03%      0.29s   0.7%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.27s  0.65% 68.68%      4.72s 11.36%  crypto/internal/fips140/sha256.(*Digest).Write
     0.25s   0.6% 69.28%      0.25s   0.6%  sync/atomic.(*Int32).Add
     0.24s  0.58% 69.86%      0.24s  0.58%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.24s  0.58% 70.44%      0.24s  0.58%  indexbody
     0.24s  0.58% 71.02%      0.26s  0.63%  runtime.typePointers.next
     0.23s  0.55% 71.57%      0.25s   0.6%  runtime.(*mspan).markBitsForIndex (inline)
     0.21s  0.51% 72.08%      0.21s  0.51%  crypto/internal/fips140/aes/gcm.gcmAesEnc
     0.21s  0.51% 72.58%      0.21s  0.51%  runtime.futex
     0.21s  0.51% 73.09%      0.86s  2.07%  runtime.mallocgcSmallScanNoHeader
     0.19s  0.46% 73.54%      4.73s 11.39%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.17s  0.41% 73.95%      1.58s  3.80%  runtime.mapassign_faststr
     0.17s  0.41% 74.36%      0.21s  0.51%  runtime.spanOf (inline)
     0.16s  0.39% 74.75%      4.96s 11.94%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.16s  0.39% 75.13%      0.28s  0.67%  internal/runtime/maps.(*Iter).Next
     0.16s  0.39% 75.52%      5.11s 12.30%  runtime.mallocgc
     0.16s  0.39% 75.90%      0.48s  1.16%  runtime.scanObjectsSmall
     0.15s  0.36% 76.26%      0.23s  0.55%  runtime.heapArenaOf (inline)
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 52.5% / 100% | 38% | 0.2% | 8 | 149 MB | CPU が飽和 |
| isu2 (2 cpu) | 9.9% / 38% | 0% | 0.6% | 6 | 2642 MB | run queue が積もっている |
| isu3 (2 cpu) | 6.3% / 25% | 0% | 0.3% | 3 | 2785 MB |  |
| isu4 (2 cpu) | 44.6% / 100% | 14% | 18.0% | 28 | 108 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 84% / nginx 12% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 17% |
| isu3 | isureal 11% |
| isu4 | postgres 42% / isureal 8% |

## Diff from previous run

スコア: 5222 → 65037 (+59815)

`3d5209e..fdbe826`

```
 webapp/go/account_deletion_api.go |   2 +-
 webapp/go/admin_user_api.go       |   2 +-
 webapp/go/comment_api.go          |   2 +-
 webapp/go/realmoji_api.go         |   2 +-
 webapp/go/relationship.go         |  10 +++-
 webapp/go/server.go               |   1 -
 webapp/go/slow.go                 |  21 --------
 webapp/go/store.go                | 100 ++++++++++++++++++++++++++++++++------
 8 files changed, 96 insertions(+), 44 deletions(-)
```
