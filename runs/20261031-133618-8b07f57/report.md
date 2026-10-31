# run 20261031-133618-8b07f57

## Summary

| | |
|---|---|
| score | 104121 |
| status | pass |
| ref | gp-post-path-2 |
| sha | 8b07f57bfd06 |
| 所要時間 | 335s |
| 前回比 | +39421 (60.9%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +39421 (60.9%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | 計測用: 投稿の段階別の時間（main 7e687d9） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	52	152	429	959	957
europe	7	19	52	145	146	146
america	6	17	54	123	123	122
```

```
{"id":"13626","status":"done","portalStatus":"finished","passed":true,"score":104121,"raw":104121,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:40:25.300Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:40:28.267Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:40:31.263Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:40:32.092Z level=INFO msg="user moved" user=nuhete5096 from=europe to=asia
time=2026-10-31T04:40:34.298Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T04:40:36.417Z level=INFO msg="user moved" user=nuhete5096 from=asia to=europe
time=2026-10-31T04:40:37.441Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T04:40:40.264Z level=INFO msg="america 2 日目: 17 人"
time=2026-10-31T04:40:43.299Z level=INFO msg="asia 3 日目: 152 人"
time=2026-10-31T04:40:44.574Z level=INFO msg="user moved" user=madan_5685 from=asia to=america
time=2026-10-31T04:40:45.189Z level=INFO msg="user moved" user=maxyofu-5838 from=asia to=america
time=2026-10-31T04:40:46.313Z level=INFO msg="europe 3 日目: 52 人"
time=2026-10-31T04:40:46.382Z level=INFO msg="user moved" user=tormiaor_5995 from=asia to=europe
time=2026-10-31T04:40:46.535Z level=INFO msg="user moved" user=maxyofu-5838 from=america to=europe
time=2026-10-31T04:40:46.835Z level=INFO msg="user moved" user=danma_6208 from=asia to=america
time=2026-10-31T04:40:48.026Z level=INFO msg="user moved" user=danma_6208 from=america to=europe
time=2026-10-31T04:40:48.855Z level=INFO msg="user moved" user=madan_5685 from=america to=europe
time=2026-10-31T04:40:49.289Z level=INFO msg="america 3 日目: 54 人"
time=2026-10-31T04:40:52.296Z level=INFO msg="asia 4 日目: 429 人"
time=2026-10-31T04:40:55.274Z level=INFO msg="europe 4 日目: 145 人"
time=2026-10-31T04:40:57.701Z level=INFO msg="user moved" user=maxyofu-5838 from=europe to=america
time=2026-10-31T04:40:58.412Z level=INFO msg="america 4 日目: 123 人"
time=2026-10-31T04:41:01.409Z level=INFO msg="user moved" user=leolia-6706 from=asia to=europe
time=2026-10-31T04:41:01.480Z level=INFO msg="asia 5 日目: 959 人"
time=2026-10-31T04:41:01.844Z level=INFO msg="user moved" user=chinuzoe_7195 from=asia to=europe
time=2026-10-31T04:41:04.361Z level=INFO msg="europe 5 日目: 146 人"
time=2026-10-31T04:41:07.509Z level=INFO msg="america 5 日目: 123 人"
time=2026-10-31T04:41:16.343Z level=INFO msg="europe 最終日（グランプリ）: 146 人"
time=2026-10-31T04:41:16.482Z level=INFO msg="america 最終日（グランプリ）: 122 人"
time=2026-10-31T04:41:16.690Z level=INFO msg="asia 最終日（グランプリ）: 957 人"
time=2026-10-31T04:41:25.301Z level=INFO msg="グランプリ: アジア 259 人, ヨーロッパ 42 人, アメリカ 23 人"
{
  "pass": true,
  "score": 104121,
  "messages": [
    "evatolia5034: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "anni10022: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
30.4%	1511	1	0	GET	/api/notifications	12423.240	8.222	15.229	15.243	166
18.1%	2166	60	0	POST	/api/posts	7415.879	3.424	12.229	15.338	632
14.2%	4950	79	0	POST	/api/friend-requests	5797.254	1.171	10.652	16.389	114
9.0%	11891	427	0	GET	^/api/images/[^/]+	3670.917	0.309	7.348	14.678	139540
6.8%	5570	1823	0	POST	/api/auth/login	2769.611	0.497	6.376	12.621	173
5.6%	1672	10	0	POST	^/api/friend-requests/[^/]+/[^/]+$	2285.098	1.367	11.577	13.630	115
3.0%	980	1	0	POST	/api/auth/register	1222.773	1.248	7.456	10.628	227
2.2%	656	125	0	PUT	/api/me/profile-image	892.970	1.361	7.061	11.046	4
1.6%	633	0	0	PUT	^/api/follows/[^/]+$	650.790	1.028	6.055	19.877	0
1.5%	417	22	0	POST	^/api/posts/[^/]+/realmojis$	600.109	1.439	8.505	12.440	458
1.2%	466	0	0	PATCH	/api/me	505.049	1.084	6.729	10.780	275
1.0%	401	14	0	POST	^/api/posts/[^/]+/comments$	419.966	1.047	7.419	15.230	317
0.9%	3044	0	0	GET	/api/memories	372.668	0.122	0.670	1.937	10908
0.7%	620	0	0	POST	/api/auth/logout	297.228	0.479	9.832	13.047	0
0.7%	10646	0	0	GET	^/api/feeds/[^/]+	294.952	0.028	0.292	2.195	1879
0.6%	191	11	0	PATCH	^/api/posts/[^/]+$	254.493	1.332	10.476	11.863	610
0.6%	88	0	0	DELETE	^/api/posts/[^/]+$	239.638	2.723	12.102	12.102	0
0.5%	201	0	0	GET	/api/me/export	195.871	0.974	9.486	11.792	1880109
0.2%	3137	0	0	GET	/api/posting-window/current	89.741	0.029	0.248	0.646	237
0.2%	74	0	0	DELETE	^/api/realmojis/[^/]+$	68.246	0.922	5.908	5.908	0
0.2%	1640	91	0	GET	^/api/posts/me/current$	65.845	0.040	0.549	1.482	615
0.2%	2405	0	0	GET	/api/users	61.884	0.026	0.172	1.078	242
0.1%	73	0	0	DELETE	^/api/comments/[^/]+$	55.229	0.757	8.645	8.645	0
0.1%	1432	0	0	GET	^/api/relationships/[^/]+	37.899	0.026	0.351	1.194	861
0.1%	1501	23	0	GET	^/api/posts/[^/]+$	36.303	0.024	0.169	1.544	618
0.1%	519	0	0	GET	^/api/memories/[^/]+	33.346	0.064	0.500	1.027	693
0.1%	20	0	0	PUT	^/api/blocks/[^/]+$	27.639	1.382	15.382	15.382	0
0.1%	16	0	0	PUT	^/api/admin/users/[^/]+/certification$	21.980	1.374	11.342	11.342	0
0.0%	1	0	0	POST	/api/initialize	19.205	19.205	19.205	19.205	90
0.0%	5	0	0	DELETE	^/api/friendships/[^/]+$	16.437	3.287	8.320	8.320	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.9%	329	0	0	GET	/api/notifications	2712.003	8.243	11.854	12.230	177
18.1%	418	7	0	POST	/api/posts	1408.593	3.370	12.410	14.440	636
12.5%	822	14	0	POST	/api/friend-requests	969.696	1.180	7.411	12.248	114
6.7%	506	0	0	POST	/api/auth/login	519.393	1.026	10.789	14.336	247
4.1%	1755	14	0	GET	^/api/images/[^/]+	320.843	0.183	0.937	1.594	176279
3.9%	1547	0	0	GET	^/api/feeds/[^/]+	306.246	0.198	0.895	1.594	2095
3.7%	220	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	284.414	1.293	9.415	10.043	115
1.8%	465	0	0	GET	/api/memories	137.858	0.296	1.249	1.873	12209
1.6%	575	0	0	GET	/api/posting-window/current	124.483	0.216	0.690	1.032	238
1.5%	101	7	0	POST	^/api/posts/[^/]+/realmojis$	113.646	1.125	10.125	15.650	419
1.0%	429	0	0	GET	/api/users	80.333	0.187	0.647	1.485	242
1.0%	327	25	0	GET	^/api/posts/me/current$	79.031	0.242	1.074	1.296	599
1.0%	139	0	0	POST	/api/auth/register	74.565	0.536	2.529	2.616	227
1.0%	82	0	0	POST	/api/auth/logout	73.894	0.901	7.820	7.820	0
0.9%	77	0	0	PUT	/api/me/profile-image	69.480	0.902	11.574	11.574	0
0.9%	36	0	0	GET	/api/me/export	67.306	1.870	8.206	8.206	2065623
0.9%	83	0	0	PUT	^/api/follows/[^/]+$	66.231	0.798	8.258	8.258	0
0.8%	59	3	0	POST	^/api/posts/[^/]+/comments$	62.611	1.061	7.833	7.833	330
0.7%	33	5	0	PATCH	^/api/posts/[^/]+$	57.497	1.742	9.007	9.007	537
0.5%	204	0	0	GET	^/api/relationships/[^/]+	41.065	0.201	1.173	1.182	868
0.5%	183	12	0	GET	^/api/posts/[^/]+$	39.138	0.214	1.590	1.650	593
0.5%	58	0	0	PATCH	/api/me	36.985	0.638	3.753	3.753	273
0.3%	15	0	0	DELETE	^/api/posts/[^/]+$	26.083	1.739	6.007	6.007	0
0.3%	16	0	0	DELETE	^/api/realmojis/[^/]+$	22.575	1.411	7.938	7.938	0
0.3%	102	0	0	GET	^/api/users/[^/]+$	19.696	0.193	0.670	0.744	202
0.2%	70	0	0	GET	^/api/memories/[^/]+	16.155	0.231	0.794	0.794	819
0.2%	89	0	0	GET	^/api/posts/[^/]+/realmojis$	15.845	0.178	1.113	1.113	1297
0.2%	80	1	0	GET	/api/me	13.038	0.163	0.482	0.482	298
0.1%	62	0	0	GET	^/api/posts/[^/]+/comments$	9.644	0.156	0.459	0.459	1666
0.1%	8	0	0	PUT	^/api/blocks/[^/]+$	5.293	0.662	2.315	2.315	0
…（上位 30 / 全 31 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
30.8%	258	0	0	GET	/api/notifications	1860.133	7.210	9.190	9.220	177
15.4%	315	10	0	POST	/api/posts	929.869	2.952	11.305	12.137	632
14.2%	642	10	0	POST	/api/friend-requests	856.807	1.335	10.826	12.171	114
7.2%	392	0	0	POST	/api/auth/login	435.441	1.111	8.353	9.564	247
5.7%	245	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	343.318	1.401	10.887	12.600	116
5.0%	1344	17	0	GET	^/api/images/[^/]+	299.418	0.223	1.127	2.385	169133
3.1%	936	0	0	GET	^/api/feeds/[^/]+	187.214	0.200	0.686	1.580	2823
2.0%	380	0	0	GET	/api/memories	118.797	0.313	1.373	1.845	14461
1.7%	69	0	0	PUT	/api/me/profile-image	103.802	1.504	6.263	6.263	0
1.6%	447	0	0	GET	/api/posting-window/current	97.489	0.218	0.621	1.078	240
1.6%	121	0	0	POST	/api/auth/register	93.623	0.774	5.429	8.191	229
1.5%	73	0	0	PUT	^/api/follows/[^/]+$	91.339	1.251	6.686	6.686	0
1.2%	63	0	0	PATCH	/api/me	72.922	1.157	6.806	6.806	277
1.2%	22	0	0	GET	/api/me/export	71.696	3.259	13.294	13.294	3146056
1.0%	316	0	0	GET	/api/users	63.135	0.200	0.575	0.941	239
1.0%	248	26	0	GET	^/api/posts/me/current$	60.212	0.243	1.060	1.167	583
1.0%	61	0	0	POST	/api/auth/logout	58.182	0.954	5.377	5.377	0
0.9%	22	0	0	DELETE	^/api/posts/[^/]+$	54.362	2.471	14.958	14.958	0
0.7%	27	1	0	PATCH	^/api/posts/[^/]+$	43.920	1.627	6.967	6.967	648
0.6%	47	1	0	POST	^/api/posts/[^/]+/comments$	37.642	0.801	7.518	7.518	334
0.6%	29	1	0	POST	^/api/posts/[^/]+/realmojis$	35.632	1.229	4.048	4.048	421
0.5%	169	0	0	GET	^/api/relationships/[^/]+	29.033	0.172	0.746	0.817	1069
0.4%	146	5	0	GET	^/api/posts/[^/]+$	26.481	0.181	0.748	0.930	608
0.3%	82	0	0	GET	^/api/users/[^/]+$	15.230	0.186	1.190	1.190	205
0.2%	72	1	0	GET	/api/me	12.538	0.174	0.428	0.428	300
0.2%	43	0	0	GET	^/api/memories/[^/]+	11.368	0.264	1.643	1.643	884
0.1%	52	0	0	GET	^/api/posts/[^/]+/comments$	9.053	0.174	0.484	0.484	1118
0.1%	40	0	0	GET	^/api/posts/[^/]+/realmojis$	7.367	0.184	0.500	0.500	906
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	3.313	1.104	3.100	3.100	0
0.1%	2	0	0	DELETE	^/api/realmojis/[^/]+$	3.025	1.512	2.812	2.812	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 66349 クエリ / 81 種類 / 合計 503.889s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	67.2%	4016	338.403	84.26	788.55	674.0	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	18.2%	259	91.703	354.07	1741.00	1268.6	128.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	5.0%	2660	25.248	9.49	458.80	118.7	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.1%	7776	20.632	2.65	42.99	10.6	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.6%	2822	8.015	2.84	30.45	24.9	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.2%	2	5.820	2910.10	5671.28	240618.5	2833.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.5%	1	2.319	2318.83	2318.83	30770.0	1199950.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.4%	517	1.912	3.70	23.45	29.7	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.4%	1	1.778	1777.87	1777.87	23652.0	627909.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
10	0.2%	6293	1.198	0.19	28.04	19.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.2%	489	0.949	1.94	20.91	21.3	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.2%	1	0.892	891.76	891.76	29249.0	2798.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.2%	1	0.800	799.76	799.76	12468.0	530096.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
14	0.1%	4645	0.607	0.13	15.87	11.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
15	0.1%	2064	0.529	0.26	12.49	20.8	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)

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
Build ID: 30c61df8feb35332ec4509cc4a8a38360838b774
Type: cpu
Time: 2026-10-31 13:40:00 JST
Duration: 60.04s, Total samples = 37.84s (63.02%)
Showing nodes accounting for 27.60s, 72.94% of 37.84s total
Dropped 961 nodes (cum <= 0.19s)
Showing top 40 nodes out of 288
      flat  flat%   sum%        cum   cum%
     3.23s  8.54%  8.54%      3.23s  8.54%  internal/runtime/syscall/linux.Syscall6
     3.15s  8.32% 16.86%      3.15s  8.32%  crypto/internal/fips140/sha256.blockAVX2
     1.87s  4.94% 21.80%      1.87s  4.94%  runtime.memmove
     1.73s  4.57% 26.37%      1.73s  4.57%  runtime.memclrNoHeapPointers
     1.67s  4.41% 30.79%      2.08s  5.50%  runtime.findObject
     1.57s  4.15% 34.94%      6.51s 17.20%  runtime.scanObject
     1.54s  4.07% 39.01%      1.54s  4.07%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.44s  3.81% 42.81%      5.77s 15.25%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.44s  3.81% 46.62%      2.45s  6.47%  runtime.tryDeferToSpanScan
     1.13s  2.99% 49.60%      1.13s  2.99%  aeshashbody
     1.08s  2.85% 52.46%      1.08s  2.85%  example.com/isureal-company/isureal.(*UserRow).active
     0.73s  1.93% 54.39%      0.73s  1.93%  memeqbody
     0.69s  1.82% 56.21%      0.69s  1.82%  runtime.spanClass.sizeclass (inline)
     0.55s  1.45% 57.66%      2.10s  5.55%  runtime.mapaccess1_faststr
     0.39s  1.03% 58.69%      0.99s  2.62%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.39s  1.03% 59.73%      0.39s  1.03%  indexbody
     0.36s  0.95% 60.68%      0.36s  0.95%  runtime.nextFreeFast (inline)
     0.31s  0.82% 61.50%      0.45s  1.19%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.29s  0.77% 62.26%      0.49s  1.29%  example.com/isureal-company/isureal.postLess
     0.29s  0.77% 63.03%      0.33s  0.87%  runtime.typePointers.next
     0.28s  0.74% 63.77%      0.28s  0.74%  runtime.futex
     0.27s  0.71% 64.48%      0.57s  1.51%  database/sql.convertAssignRows
     0.27s  0.71% 65.20%      0.27s  0.71%  sync/atomic.(*Int32).Add
     0.23s  0.61% 65.80%      0.24s  0.63%  example.com/isureal-company/isureal.rmLess
     0.23s  0.61% 66.41%         1s  2.64%  runtime.mallocgcSmallScanNoHeader
     0.21s  0.55% 66.97%      0.24s  0.63%  runtime.spanOf (inline)
     0.20s  0.53% 67.49%      0.45s  1.19%  runtime.greyobject
     0.20s  0.53% 68.02%      0.25s  0.66%  runtime.heapArenaOf (inline)
     0.20s  0.53% 68.55%      0.20s  0.53%  runtime.typePointers.nextFast (inline)
     0.19s   0.5% 69.05%      1.85s  4.89%  runtime.mapassign_faststr
     0.18s  0.48% 69.53%      0.19s   0.5%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.18s  0.48% 70.01%      4.38s 11.58%  runtime.mallocgc
     0.17s  0.45% 70.45%      0.22s  0.58%  runtime.(*mspan).markBitsForIndex (inline)
     0.16s  0.42% 70.88%      0.43s  1.14%  compress/flate.(*compressor).deflate
     0.16s  0.42% 71.30%      1.11s  2.93%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.15s   0.4% 71.70%      0.43s  1.14%  runtime.scanObjectsSmall
     0.12s  0.32% 72.01%      2.94s  7.77%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.12s  0.32% 72.33%      1.93s  5.10%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.12s  0.32% 72.65%      0.55s  1.45%  runtime.mapaccess2_faststr
     0.11s  0.29% 72.94%      2.60s  6.87%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 52.4% / 100% | 32% | 0.5% | 10 | 142 MB | CPU が飽和 |
| isu2 (2 cpu) | 19.0% / 94% | 1% | 0.6% | 3 | 1933 MB |  |
| isu3 (2 cpu) | 17.0% / 95% | 2% | 0.7% | 3 | 1733 MB |  |
| isu4 (2 cpu) | 50.5% / 100% | 20% | 16.7% | 30 | 118 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 72% / nginx 12% / journalctl 3% / systemd-journal 2% |
| isu2 | isureal 36% / nginx 1% |
| isu3 | isureal 31% / nginx 1% |
| isu4 | postgres 58% / isureal 5% |

## Diff from previous run

スコア: 64700 → 104121 (+39421)

`e0e4cb3..8b07f57`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf     |   3 +
 conf/isu1/home/isucon/env.sh                       |   1 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf     |  14 +
 conf/isu2/home/isucon/env.sh                       |   2 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf     |  14 +
 conf/isu3/home/isucon/env.sh                       |   2 +
 .../etc/postgresql/16/main/conf.d/zz-tuning.conf   |   2 +
 webapp/go/account_deletion_api.go                  |   2 +-
 webapp/go/admin_user_api.go                        |   2 +-
 webapp/go/comment_api.go                           |   2 +-
 webapp/go/config.go                                |   2 +
 webapp/go/export_api.go                            |   4 +-
 webapp/go/feed_api.go                              | 113 +++++--
 webapp/go/feed_test.go                             | 331 +++++++++++++++++++++
 webapp/go/image_local.go                           |   3 +-
 webapp/go/offload.go                               |  27 +-
 webapp/go/post_api.go                              |  80 ++++-
 webapp/go/post_forward.go                          | 113 +++++++
 webapp/go/post_forward_test.go                     |  94 ++++++
 webapp/go/realmoji_api.go                          |   2 +-
 webapp/go/relationship.go                          |  10 +-
 webapp/go/store.go                                 | 113 ++++++-
 22 files changed, 870 insertions(+), 66 deletions(-)
```
