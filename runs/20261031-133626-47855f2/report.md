# run 20261031-133626-47855f2

## Summary

| | |
|---|---|
| score | 114291 |
| status | pass |
| ref | asia1-cpu-small |
| sha | 47855f20f0d7 |
| 所要時間 | 458s |
| 前回比 | +8579 (8.1%) vs 20261031-131852-53fe836 |
| ベスト比 | +8579 (8.1%) vs 20261031-131852-53fe836 = 105712 **ベスト更新** |
| note | groupA-on-7e687d9 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	149	428	969	968
europe	7	20	61	169	171	171
america	6	20	56	129	131	131
```

```
{"id":"13818","status":"done","portalStatus":"finished","passed":true,"score":114291,"raw":114291,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:42:34.524Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:42:37.495Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:42:40.493Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:42:43.516Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T04:42:46.131Z level=INFO msg="user moved" user=hasa-5296 from=asia to=america
time=2026-10-31T04:42:46.493Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T04:42:49.493Z level=INFO msg="america 2 日目: 20 人"
time=2026-10-31T04:42:52.516Z level=INFO msg="asia 3 日目: 149 人"
time=2026-10-31T04:42:54.208Z level=INFO msg="user moved" user=tahoshi_5837 from=asia to=america
time=2026-10-31T04:42:54.371Z level=INFO msg="user moved" user=nianka.5858 from=asia to=europe
time=2026-10-31T04:42:55.389Z level=INFO msg="user moved" user=tahoshi_5837 from=america to=europe
time=2026-10-31T04:42:55.434Z level=INFO msg="user moved" user=nianka.5858 from=europe to=asia
time=2026-10-31T04:42:55.497Z level=INFO msg="europe 3 日目: 61 人"
time=2026-10-31T04:42:58.494Z level=INFO msg="america 3 日目: 56 人"
time=2026-10-31T04:42:58.868Z level=INFO msg="user moved" user=rusol-6368 from=europe to=asia
time=2026-10-31T04:43:01.541Z level=INFO msg="asia 4 日目: 428 人"
time=2026-10-31T04:43:04.389Z level=INFO msg="user moved" user=rusol-6368 from=asia to=europe
time=2026-10-31T04:43:04.536Z level=INFO msg="europe 4 日目: 169 人"
time=2026-10-31T04:43:05.682Z level=INFO msg="user moved" user=nianka.5858 from=asia to=europe
time=2026-10-31T04:43:06.277Z level=INFO msg="user moved" user=uleva7421 from=asia to=america
time=2026-10-31T04:43:07.520Z level=INFO msg="america 4 日目: 129 人"
time=2026-10-31T04:43:07.732Z level=INFO msg="user moved" user=leoisma-7300 from=asia to=america
time=2026-10-31T04:43:08.512Z level=INFO msg="user moved" user=seni-8417 from=asia to=europe
time=2026-10-31T04:43:10.653Z level=INFO msg="asia 5 日目: 969 人"
time=2026-10-31T04:43:10.663Z level=INFO msg="user moved" user=anru.8283 from=asia to=america
time=2026-10-31T04:43:13.629Z level=INFO msg="europe 5 日目: 171 人"
time=2026-10-31T04:43:16.730Z level=INFO msg="america 5 日目: 131 人"
time=2026-10-31T04:43:25.646Z level=INFO msg="asia 最終日（グランプリ）: 968 人"
time=2026-10-31T04:43:25.839Z level=INFO msg="europe 最終日（グランプリ）: 171 人"
time=2026-10-31T04:43:26.012Z level=INFO msg="america 最終日（グランプリ）: 131 人"
time=2026-10-31T04:43:34.541Z level=INFO msg="グランプリ: アジア 440 人, ヨーロッパ 64 人, アメリカ 45 人"
{
  "pass": true,
  "score": 114291,
  "messages": [
    "meni.11797: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "rais_10205: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "tami_11772: 友達申請を送れません: POST /api/friend-requests: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests\": EOF",
    "vinme11566: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "noasotsu12846: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "taruzoe.7909: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "elniro.9814: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "leoya.12754: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.7%	1483	0	0	GET	/api/notifications	12499.088	8.428	15.091	15.101	162
17.1%	2190	69	0	POST	/api/posts	6531.205	2.982	10.026	12.431	628
13.3%	4921	48	0	POST	/api/friend-requests	5091.870	1.035	7.013	15.925	114
8.4%	12809	437	0	GET	^/api/images/[^/]+	3202.861	0.250	4.312	11.919	140000
6.2%	6386	2325	0	POST	/api/auth/login	2365.352	0.370	5.572	12.530	165
4.9%	1744	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1866.051	1.070	7.267	11.138	116
4.0%	670	122	0	PUT	/api/me/profile-image	1515.014	2.261	10.192	13.667	4
2.2%	991	1	0	POST	/api/auth/register	854.018	0.862	8.518	15.335	227
1.7%	331	48	0	POST	^/api/posts/[^/]+/realmojis$	633.533	1.914	10.421	13.937	420
1.6%	670	0	0	PUT	^/api/follows/[^/]+$	628.723	0.938	6.133	7.683	0
1.6%	544	41	0	POST	^/api/posts/[^/]+/comments$	598.796	1.101	7.532	14.445	306
1.2%	472	0	0	PATCH	/api/me	470.639	0.997	7.223	10.603	277
1.0%	3315	0	0	GET	/api/memories	383.216	0.116	0.789	2.143	10366
0.9%	696	0	0	POST	/api/auth/logout	330.540	0.475	6.787	9.842	0
0.8%	199	7	0	PATCH	^/api/posts/[^/]+$	299.506	1.505	7.388	7.651	626
0.6%	244	0	0	GET	/api/me/export	214.705	0.880	4.596	7.912	1790221
0.5%	10246	0	0	GET	^/api/feeds/[^/]+	175.882	0.017	0.165	0.502	1935
0.4%	85	0	0	DELETE	^/api/posts/[^/]+$	160.158	1.884	10.391	10.391	0
0.1%	3191	0	0	GET	/api/posting-window/current	53.527	0.017	0.143	3.153	237
0.1%	88	0	0	DELETE	^/api/comments/[^/]+$	50.833	0.578	8.420	8.420	0
0.1%	47	0	0	DELETE	^/api/realmojis/[^/]+$	42.272	0.899	5.214	5.214	0
0.1%	620	0	0	GET	^/api/memories/[^/]+	41.442	0.067	0.733	0.927	659
0.1%	2354	0	0	GET	/api/users	35.667	0.015	0.127	1.054	237
0.1%	1775	71	0	GET	^/api/posts/me/current$	35.487	0.020	0.217	0.536	623
0.1%	24	0	0	PUT	^/api/blocks/[^/]+$	26.710	1.113	5.091	5.091	0
0.1%	1437	0	0	GET	^/api/relationships/[^/]+	22.304	0.016	0.183	0.494	985
0.1%	1490	19	0	GET	^/api/posts/[^/]+$	20.373	0.014	0.113	0.176	613
0.1%	1	0	0	POST	/api/initialize	19.335	19.335	19.335	19.335	90
0.0%	14	0	0	PUT	^/api/admin/users/[^/]+/certification$	17.499	1.250	12.571	12.571	0
0.0%	8	0	0	DELETE	^/api/follows/[^/]+$	17.286	2.161	4.662	4.662	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.6%	408	0	0	GET	/api/notifications	3432.861	8.414	11.882	12.445	178
16.0%	518	6	0	POST	/api/posts	1423.077	2.747	9.767	10.638	640
9.8%	950	8	0	POST	/api/friend-requests	868.176	0.914	7.610	11.271	114
5.9%	1962	20	0	GET	^/api/images/[^/]+	522.642	0.266	1.415	5.613	190053
4.6%	550	0	0	POST	/api/auth/login	406.981	0.740	5.076	10.669	246
4.5%	309	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	396.411	1.283	6.576	13.548	116
4.4%	1731	0	0	GET	^/api/feeds/[^/]+	393.812	0.228	0.770	1.190	2135
2.3%	714	0	0	GET	/api/posting-window/current	203.881	0.286	0.694	1.099	239
2.0%	564	0	0	GET	/api/memories	179.877	0.319	1.376	2.198	11120
1.2%	411	33	0	GET	^/api/posts/me/current$	104.142	0.253	0.844	1.100	594
1.2%	468	0	0	GET	/api/users	103.974	0.222	0.793	1.107	237
1.1%	35	0	0	GET	/api/me/export	101.095	2.888	12.377	12.377	2225383
1.0%	87	0	0	POST	/api/auth/logout	90.290	1.038	6.021	6.021	0
0.9%	60	8	0	POST	^/api/posts/[^/]+/comments$	80.816	1.347	6.902	6.902	319
0.9%	166	0	0	POST	/api/auth/register	78.943	0.476	1.962	1.994	228
0.9%	109	0	0	PUT	^/api/follows/[^/]+$	75.794	0.695	6.137	7.096	0
0.7%	95	0	0	PUT	/api/me/profile-image	62.933	0.662	2.238	2.238	0
0.7%	43	6	0	POST	^/api/posts/[^/]+/realmojis$	58.597	1.363	7.609	7.609	386
0.6%	249	0	0	GET	^/api/relationships/[^/]+	50.487	0.203	0.730	0.825	1049
0.5%	259	17	0	GET	^/api/posts/[^/]+$	48.061	0.186	0.645	0.682	575
0.4%	70	0	0	PATCH	/api/me	33.729	0.482	2.696	2.696	281
0.4%	38	7	0	PATCH	^/api/posts/[^/]+$	31.539	0.830	4.018	4.018	516
0.3%	114	0	0	GET	^/api/users/[^/]+$	24.391	0.214	0.672	0.678	182
0.2%	20	0	0	DELETE	^/api/posts/[^/]+$	21.121	1.056	3.226	3.226	0
0.2%	106	0	0	GET	^/api/posts/[^/]+/realmojis$	20.841	0.197	0.700	0.837	494
0.2%	78	0	0	GET	^/api/memories/[^/]+	20.144	0.258	1.232	1.232	807
0.2%	101	1	0	GET	/api/me	19.286	0.191	0.504	0.683	297
0.2%	82	0	0	GET	^/api/posts/[^/]+/comments$	15.649	0.191	0.761	0.761	1080
0.1%	7	0	0	DELETE	^/api/realmojis/[^/]+$	9.330	1.333	7.129	7.129	0
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	4.774	1.591	4.424	4.424	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.5%	300	1	0	GET	/api/notifications	2193.957	7.313	9.306	9.308	178
14.4%	360	10	0	POST	/api/posts	912.310	2.534	8.220	10.389	626
13.2%	715	3	0	POST	/api/friend-requests	838.853	1.173	7.967	16.858	115
6.3%	1393	10	0	GET	^/api/images/[^/]+	400.821	0.288	1.643	1.954	185015
5.3%	413	0	0	POST	/api/auth/login	335.319	0.812	5.916	9.249	248
4.2%	190	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	266.028	1.400	7.251	9.888	116
3.6%	1063	0	0	GET	^/api/feeds/[^/]+	230.408	0.217	0.755	1.139	2095
2.1%	504	0	0	GET	/api/posting-window/current	133.661	0.265	0.705	0.982	239
2.1%	439	0	0	GET	/api/memories	131.983	0.301	1.280	1.485	14148
1.6%	32	0	0	GET	/api/me/export	98.922	3.091	8.343	8.343	3681852
1.3%	60	0	0	PUT	/api/me/profile-image	82.013	1.367	10.807	10.807	0
1.3%	73	4	0	POST	^/api/posts/[^/]+/comments$	81.622	1.118	8.165	8.165	313
1.3%	338	0	0	GET	/api/users	80.693	0.239	0.731	0.979	238
1.2%	295	33	0	GET	^/api/posts/me/current$	75.805	0.257	0.772	1.142	576
1.0%	39	12	0	POST	^/api/posts/[^/]+/realmojis$	65.360	1.676	7.036	7.036	334
1.0%	54	0	0	PATCH	/api/me	63.779	1.181	10.378	10.378	264
1.0%	57	0	0	POST	/api/auth/logout	62.949	1.104	8.105	8.105	0
0.9%	64	0	0	PUT	^/api/follows/[^/]+$	57.304	0.895	5.968	5.968	0
0.9%	126	0	0	POST	/api/auth/register	54.434	0.432	1.833	2.338	229
0.5%	180	0	0	GET	^/api/relationships/[^/]+	33.570	0.186	0.691	0.717	944
0.5%	23	8	0	PATCH	^/api/posts/[^/]+$	29.154	1.268	6.684	6.684	434
0.4%	134	9	0	GET	^/api/posts/[^/]+$	27.113	0.202	0.704	0.753	576
0.4%	11	0	0	DELETE	^/api/posts/[^/]+$	23.498	2.136	4.048	4.048	0
0.2%	78	0	0	GET	^/api/users/[^/]+$	14.629	0.188	0.759	0.759	189
0.2%	51	0	0	GET	^/api/memories/[^/]+	14.400	0.282	1.145	1.145	697
0.2%	63	0	0	GET	/api/me	13.887	0.220	0.928	0.928	301
0.1%	5	0	0	PUT	^/api/blocks/[^/]+$	7.569	1.514	4.466	4.466	0
0.1%	50	1	0	GET	^/api/posts/[^/]+/comments$	7.279	0.146	0.365	0.365	1163
0.1%	27	0	0	GET	^/api/posts/[^/]+/realmojis$	5.677	0.210	0.766	0.766	462
0.1%	2	0	0	DELETE	^/api/realmojis/[^/]+$	4.892	2.446	4.407	4.407	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 70062 クエリ / 80 種類 / 合計 666.065s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	70.2%	4033	467.813	116.00	746.92	728.3	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	16.7%	311	111.082	357.18	1315.53	1274.3	129.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	5.1%	2958	33.892	11.46	1439.39	119.7	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	3.3%	8518	22.302	2.62	43.11	10.8	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.5%	2983	9.804	3.29	52.53	25.4	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.8%	2	5.497	2748.31	5350.17	239786.5	2824.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.5%	1	3.287	3287.14	3287.14	30806.0	1199865.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.3%	1	1.830	1830.39	1830.39	23769.0	627933.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.2%	347	1.592	4.59	25.00	31.1	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.2%	6513	1.413	0.22	35.24	19.9	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.2%	624	1.289	2.07	28.49	21.1	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.1%	1	0.869	869.25	869.25	29414.0	2822.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.801	800.98	800.98	12497.0	530038.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
14	0.1%	2192	0.768	0.35	57.86	21.0	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	5024	0.712	0.14	59.94	11.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

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
Build ID: 33b0b253ee2fd6812db942a9f9f1ebac590329e4
Type: cpu
Time: 2026-10-31 13:42:08 JST
Duration: 60s, Total samples = 35.22s (58.70%)
Showing nodes accounting for 25.02s, 71.04% of 35.22s total
Dropped 923 nodes (cum <= 0.18s)
Showing top 40 nodes out of 293
      flat  flat%   sum%        cum   cum%
     4.20s 11.93% 11.93%      4.20s 11.93%  crypto/internal/fips140/sha256.blockAVX2
     3.03s  8.60% 20.53%      3.03s  8.60%  internal/runtime/syscall/linux.Syscall6
     1.70s  4.83% 25.35%      1.70s  4.83%  runtime.memmove
     1.53s  4.34% 29.70%      1.53s  4.34%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.51s  4.29% 33.99%      5.68s 16.13%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.49s  4.23% 38.22%      1.95s  5.54%  runtime.findObject
     1.38s  3.92% 42.14%      1.38s  3.92%  aeshashbody
     1.20s  3.41% 45.54%      1.98s  5.62%  runtime.tryDeferToSpanScan
     1.19s  3.38% 48.92%      1.19s  3.38%  runtime.memclrNoHeapPointers
     1.17s  3.32% 52.24%      5.32s 15.11%  runtime.scanObject
     0.49s  1.39% 53.63%      0.49s  1.39%  memeqbody
     0.49s  1.39% 55.03%      0.49s  1.39%  runtime.spanClass.sizeclass (inline)
     0.43s  1.22% 56.25%      0.86s  2.44%  database/sql.convertAssignRows
     0.42s  1.19% 57.44%      2.04s  5.79%  runtime.mapaccess1_faststr
     0.36s  1.02% 58.46%      0.48s  1.36%  example.com/isureal-company/isureal.postLess
     0.33s  0.94% 59.40%      0.33s  0.94%  runtime.futex
     0.31s  0.88% 60.28%      0.31s  0.88%  runtime.nextFreeFast (inline)
     0.29s  0.82% 61.10%      0.29s  0.82%  sync/atomic.(*Int32).Add (inline)
     0.23s  0.65% 61.75%      0.43s  1.22%  runtime.greyobject
     0.22s  0.62% 62.38%      0.22s  0.62%  indexbody
     0.22s  0.62% 63.00%      0.29s  0.82%  runtime.(*mspan).writeHeapBitsSmall
     0.22s  0.62% 63.63%      1.88s  5.34%  runtime.mapassign_faststr
     0.21s   0.6% 64.22%      4.09s 11.61%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.21s   0.6% 64.82%      1.11s  3.15%  runtime.mallocgcSmallScanNoHeader
     0.20s  0.57% 65.39%      3.38s  9.60%  runtime.mallocgc
     0.20s  0.57% 65.96%      0.20s  0.57%  runtime.typePointers.nextFast (inline)
     0.17s  0.48% 66.44%      0.21s   0.6%  runtime.(*mspan).markBitsForIndex (inline)
     0.16s  0.45% 66.89%      4.58s 13.00%  crypto/internal/fips140/sha256.(*Digest).Write
     0.16s  0.45% 67.35%      0.22s  0.62%  runtime.heapArenaOf (inline)
     0.15s  0.43% 67.77%      0.19s  0.54%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.14s   0.4% 68.17%      1.61s  4.57%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.13s  0.37% 68.54%      4.28s 12.15%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.13s  0.37% 68.91%      0.63s  1.79%  runtime.mapaccess2_faststr
     0.12s  0.34% 69.25%      4.60s 13.06%  example.com/isureal-company/isureal.loadBig
     0.11s  0.31% 69.56%      0.38s  1.08%  internal/stringslite.Index
     0.11s  0.31% 69.88%      0.69s  1.96%  runtime.convTstring
     0.11s  0.31% 70.19%      0.61s  1.73%  runtime.scanSpan
     0.10s  0.28% 70.47%      0.30s  0.85%  compress/flate.(*compressor).deflate
     0.10s  0.28% 70.76%      4.52s 12.83%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.10s  0.28% 71.04%      0.57s  1.62%  encoding/json.structEncoder.encode
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 49.1% / 100% | 24% | 2.0% | 8 | 114 MB | run queue が積もっている |
| isu2 (2 cpu) | 20.8% / 100% | 2% | 0.5% | 4 | 1825 MB |  |
| isu3 (2 cpu) | 17.8% / 100% | 2% | 0.6% | 3 | 1689 MB |  |
| isu4 (2 cpu) | 53.2% / 100% | 29% | 17.2% | 36 | 124 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 66% / nginx 12% / journalctl 2% |
| isu2 | isureal 39% / nginx 1% |
| isu3 | isureal 33% / nginx 1% |
| isu4 | postgres 62% / isureal 5% |

## Diff from previous run

スコア: 105712 → 114291 (+8579)

`53fe836..47855f2`

```
 webapp/go/account_deletion_api.go |  1 +
 webapp/go/action_log.go           | 13 ++++--
 webapp/go/action_log_test.go      | 67 ++++++++++++++++++++++++++++
 webapp/go/offload.go              | 47 ++++++++++++++------
 webapp/go/profile_image_api.go    |  7 +--
 webapp/go/store.go                | 53 +++++++++++++++++++---
 webapp/go/user_api.go             |  1 +
 webapp/go/user_read_api.go        | 58 +++++++++++++-----------
 webapp/go/user_search_test.go     | 94 +++++++++++++++++++++++++++++++++++++++
 9 files changed, 289 insertions(+), 52 deletions(-)
```
