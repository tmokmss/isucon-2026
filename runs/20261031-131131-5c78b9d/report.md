# run 20261031-131131-5c78b9d

## Summary

| | |
|---|---|
| score | 96430 |
| status | pass |
| ref | repl-replica |
| sha | 5c78b9de0670 |
| 所要時間 | 128s |
| 前回比 | +16312 (20.4%) vs 20261031-130520-28e8717 |
| ベスト比 | +16312 (20.4%) vs 20261031-130520-28e8717 = 80118 **ベスト更新** |
| note | 複製 第3段 local（SSE 以外の読み取りを手元で返す、要約の照合なし） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	143	406	939	939
europe	7	18	48	123	124	123
america	6	12	29	65	66	66
```

```
{"id":"11495","status":"done","portalStatus":"finished","passed":true,"score":96430,"raw":96430,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:12:16.620Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:12:19.580Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:12:22.678Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:12:25.728Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T04:12:28.684Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T04:12:31.677Z level=INFO msg="america 2 日目: 12 人"
time=2026-10-31T04:12:34.738Z level=INFO msg="asia 3 日目: 143 人"
time=2026-10-31T04:12:37.686Z level=INFO msg="europe 3 日目: 48 人"
time=2026-10-31T04:12:40.694Z level=INFO msg="america 3 日目: 29 人"
time=2026-10-31T04:12:42.897Z level=INFO msg="user moved" user=metehi.6226 from=europe to=america
time=2026-10-31T04:12:43.765Z level=INFO msg="asia 4 日目: 406 人"
time=2026-10-31T04:12:45.689Z level=INFO msg="user moved" user=solmeyo.6279 from=europe to=asia
time=2026-10-31T04:12:46.760Z level=INFO msg="europe 4 日目: 123 人"
time=2026-10-31T04:12:49.970Z level=INFO msg="america 4 日目: 65 人"
time=2026-10-31T04:12:50.840Z level=INFO msg="user moved" user=solmeyo.6279 from=asia to=europe
time=2026-10-31T04:12:51.177Z level=INFO msg="user moved" user=yunema.8003 from=asia to=america
time=2026-10-31T04:12:53.060Z level=INFO msg="asia 5 日目: 939 人"
time=2026-10-31T04:12:56.415Z level=INFO msg="europe 5 日目: 124 人"
time=2026-10-31T04:12:59.551Z level=INFO msg="america 5 日目: 66 人"
time=2026-10-31T04:13:08.343Z level=INFO msg="america 最終日（グランプリ）: 66 人"
time=2026-10-31T04:13:09.190Z level=INFO msg="asia 最終日（グランプリ）: 939 人"
time=2026-10-31T04:13:09.430Z level=INFO msg="europe 最終日（グランプリ）: 123 人"
time=2026-10-31T04:13:16.614Z level=INFO msg="グランプリ: アジア 361 人, ヨーロッパ 14 人, アメリカ 30 人"
{
  "pass": true,
  "score": 96430,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
31.2%	1417	0	0	GET	/api/notifications	11924.608	8.415	16.164	16.187	158
16.4%	2027	65	0	POST	/api/posts	6276.312	3.096	9.463	12.252	627
13.8%	4820	78	0	POST	/api/friend-requests	5267.759	1.093	7.047	10.481	114
8.6%	4538	1361	0	POST	/api/auth/login	3275.102	0.722	7.168	13.489	178
7.8%	10246	371	0	GET	^/api/images/[^/]+	2985.105	0.291	5.442	11.399	168054
4.3%	956	1	0	POST	/api/auth/register	1642.298	1.718	8.405	10.919	227
4.3%	1542	12	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1633.145	1.059	6.658	12.791	115
2.8%	625	122	0	PUT	/api/me/profile-image	1077.397	1.724	8.502	9.822	4
1.5%	585	0	0	PUT	^/api/follows/[^/]+$	581.604	0.994	5.930	13.923	0
1.3%	461	0	0	PATCH	/api/me	493.889	1.071	7.287	9.585	277
1.2%	402	17	0	POST	^/api/posts/[^/]+/realmojis$	472.427	1.175	6.707	8.503	457
1.1%	2701	0	0	GET	/api/memories	427.279	0.158	0.912	1.683	9458
1.0%	417	4	0	POST	^/api/posts/[^/]+/comments$	396.942	0.952	7.028	9.375	336
0.9%	517	0	0	POST	/api/auth/logout	327.780	0.634	6.204	8.868	0
0.8%	10128	0	0	GET	^/api/feeds/[^/]+	289.947	0.029	0.335	1.418	2457
0.7%	190	13	0	PATCH	^/api/posts/[^/]+$	270.981	1.426	7.682	8.005	603
0.5%	165	0	0	GET	/api/me/export	183.340	1.111	5.356	7.880	1986915
0.3%	65	0	0	DELETE	^/api/posts/[^/]+$	132.297	2.035	8.147	8.147	0
0.3%	2964	0	0	GET	/api/posting-window/current	124.481	0.042	0.452	1.124	237
0.2%	2347	0	0	GET	/api/users	82.188	0.035	0.348	1.090	237
0.2%	1596	99	0	GET	^/api/posts/me/current$	78.692	0.049	0.574	1.117	609
0.1%	59	0	0	DELETE	^/api/realmojis/[^/]+$	42.883	0.727	3.061	3.061	0
0.1%	1402	0	0	GET	^/api/relationships/[^/]+	40.821	0.029	0.292	0.762	994
0.1%	1647	19	0	GET	^/api/posts/[^/]+$	35.036	0.021	0.236	0.588	608
0.1%	66	0	0	DELETE	^/api/comments/[^/]+$	34.100	0.517	6.428	6.428	0
0.1%	408	0	0	GET	^/api/memories/[^/]+	30.834	0.076	0.573	1.599	576
0.1%	1	0	0	POST	/api/initialize	21.742	21.742	21.742	21.742	90
0.0%	29	0	0	PUT	^/api/blocks/[^/]+$	17.059	0.588	4.412	4.412	0
0.0%	509	2	0	GET	/api/me	17.031	0.033	0.344	0.757	302
0.0%	747	0	0	GET	^/api/users/[^/]+$	14.386	0.019	0.199	0.474	189
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.0%	318	1	0	GET	/api/notifications	2779.778	8.741	13.388	13.802	177
21.9%	391	4	0	POST	/api/posts	1486.849	3.803	9.913	10.815	636
9.0%	698	12	0	POST	/api/friend-requests	608.014	0.871	5.877	9.447	114
7.1%	1921	11	0	GET	^/api/images/[^/]+	483.692	0.252	1.609	2.056	172443
7.1%	532	0	0	POST	/api/auth/login	482.118	0.906	7.640	11.027	248
5.5%	250	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	370.176	1.481	6.884	7.586	116
1.4%	64	4	0	POST	^/api/posts/[^/]+/realmojis$	93.189	1.456	11.925	11.925	462
1.1%	92	0	0	POST	/api/auth/logout	75.761	0.823	6.234	6.234	0
1.1%	102	1	0	POST	^/api/posts/[^/]+/comments$	72.538	0.711	5.657	8.722	323
0.9%	123	0	0	POST	/api/auth/register	60.016	0.488	1.428	1.568	228
0.9%	106	0	0	PUT	^/api/follows/[^/]+$	57.949	0.547	1.801	10.073	0
0.8%	61	0	0	PUT	/api/me/profile-image	51.164	0.839	11.646	11.646	0
0.6%	28	1	0	PATCH	^/api/posts/[^/]+$	42.002	1.500	7.566	7.566	639
0.4%	60	0	0	PATCH	/api/me	26.179	0.436	2.216	2.216	269
0.3%	11	0	0	DELETE	^/api/posts/[^/]+$	18.721	1.702	7.301	7.301	0
0.3%	13	0	0	DELETE	^/api/realmojis/[^/]+$	18.228	1.402	5.639	5.639	0
0.2%	1733	0	0	GET	^/api/feeds/[^/]+	12.818	0.007	0.123	0.495	2922
0.1%	508	0	0	GET	/api/memories	8.865	0.017	0.222	0.940	14899
0.1%	539	0	0	GET	/api/posting-window/current	7.457	0.014	0.254	0.456	239
0.1%	15	0	0	DELETE	^/api/comments/[^/]+$	6.816	0.454	1.758	1.758	0
0.1%	24	0	0	GET	/api/me/export	5.295	0.221	1.241	1.241	2576981
0.1%	344	0	0	GET	/api/users	3.518	0.010	0.121	0.236	230
0.0%	230	0	0	GET	^/api/relationships/[^/]+	2.721	0.012	0.128	1.016	878
0.0%	4	0	0	PUT	^/api/blocks/[^/]+$	2.488	0.622	1.609	1.609	0
0.0%	1	0	0	POST	/api/me/account-deletion	2.088	2.088	2.088	2.088	0
0.0%	74	0	0	GET	^/api/memories/[^/]+	1.521	0.021	1.036	1.036	729
0.0%	259	4	0	GET	^/api/posts/[^/]+$	1.293	0.005	0.103	0.184	603
0.0%	299	12	0	GET	^/api/posts/me/current$	1.274	0.004	0.104	0.125	613
0.0%	146	0	0	GET	^/api/users/[^/]+$	0.787	0.005	0.106	0.148	187
0.0%	1	0	0	DELETE	^/api/follows/[^/]+$	0.777	0.777	0.777	0.777	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.9%	159	0	0	GET	/api/notifications	1183.692	7.445	9.426	9.426	181
16.8%	181	0	0	POST	/api/posts	552.219	3.051	9.578	10.370	659
13.5%	375	8	0	POST	/api/friend-requests	445.816	1.189	7.219	8.238	113
9.4%	1058	4	0	GET	^/api/images/[^/]+	310.378	0.293	2.067	5.439	155606
8.2%	177	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	270.292	1.527	7.930	8.925	115
7.5%	291	0	0	POST	/api/auth/login	245.610	0.844	5.645	9.614	252
1.9%	35	0	0	PUT	/api/me/profile-image	62.444	1.784	9.784	9.784	0
1.0%	33	0	0	PUT	^/api/follows/[^/]+$	32.725	0.992	8.111	8.111	0
0.9%	22	3	0	POST	^/api/posts/[^/]+/comments$	30.776	1.399	11.303	11.303	296
0.8%	28	0	0	PATCH	/api/me	27.841	0.994	6.244	6.244	279
0.7%	63	0	0	POST	/api/auth/register	23.528	0.373	1.407	1.407	229
0.7%	13	4	0	POST	^/api/posts/[^/]+/realmojis$	23.420	1.802	5.879	5.879	314
0.7%	18	2	0	PATCH	^/api/posts/[^/]+$	23.202	1.289	4.283	4.283	567
0.6%	35	0	0	POST	/api/auth/logout	18.755	0.536	2.175	2.175	0
0.3%	718	0	0	GET	^/api/feeds/[^/]+	9.161	0.013	0.231	0.534	3244
0.2%	278	0	0	GET	/api/memories	7.403	0.027	0.338	0.572	13303
0.2%	4	0	0	DELETE	^/api/posts/[^/]+$	5.151	1.288	2.530	2.530	0
0.2%	18	0	0	GET	/api/me/export	5.018	0.279	1.247	1.247	2996147
0.1%	245	0	0	GET	/api/posting-window/current	3.413	0.014	0.167	0.470	240
0.1%	161	0	0	GET	/api/users	3.130	0.019	0.215	0.318	235
0.1%	107	0	0	GET	^/api/relationships/[^/]+	2.294	0.021	0.489	0.514	1564
0.1%	6	0	0	DELETE	^/api/comments/[^/]+$	2.194	0.366	1.120	1.120	0
0.0%	156	14	0	GET	^/api/posts/me/current$	1.555	0.010	0.154	0.432	598
0.0%	2	0	0	DELETE	^/api/realmojis/[^/]+$	1.071	0.535	0.731	0.731	0
0.0%	70	0	0	GET	^/api/users/[^/]+$	0.805	0.012	0.187	0.187	218
0.0%	46	0	0	GET	^/api/memories/[^/]+	0.522	0.011	0.220	0.220	1023
0.0%	111	0	0	GET	^/api/posts/[^/]+$	0.379	0.003	0.030	0.102	648
0.0%	41	0	0	GET	^/api/posts/[^/]+/realmojis$	0.367	0.009	0.200	0.200	549
0.0%	50	0	0	GET	^/api/posts/[^/]+/comments$	0.345	0.007	0.102	0.102	899
0.0%	2	0	0	PUT	^/api/blocks/[^/]+$	0.323	0.162	0.216	0.216	0
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 58922 クエリ / 82 種類 / 合計 454.047s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	60.1%	3588	272.981	76.08	599.99	666.8	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	14.2%	170	64.436	379.03	1483.62	6235.4	127.9	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.9%	2531	26.712	10.55	596.54	116.4	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.7%	3	21.475	7158.30	9573.29	30535.0	1200003.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
5	3.6%	6189	16.235	2.62	47.94	11.2	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
6	3.5%	3	16.039	5346.29	7134.77	22923.0	627823.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	1.9%	6	8.824	1470.63	7962.67	219271.2	2590.2	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
8	1.7%	3	7.635	2545.02	3436.73	12203.0	530275.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	1.5%	2535	7.019	2.77	61.87	25.0	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.4%	454	1.772	3.90	20.81	29.2	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
11	0.3%	5770	1.414	0.25	18.75	19.8	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
12	0.2%	1	1.095	1095.41	1095.41	2646.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
13	0.2%	533	1.002	1.88	18.59	22.2	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
14	0.2%	6	0.942	156.93	322.19	16780.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	0.2%	1	0.899	899.26	899.26	32654.0	1.0	SELECT media_type, count(*) FROM image_objects GROUP BY 1

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
Build ID: eb4c7e01125ae5da8961f80df8f107d1a4811d9f
Type: cpu
Time: 2026-10-31 13:11:51 JST
Duration: 60s, Total samples = 32.12s (53.53%)
Showing nodes accounting for 22.86s, 71.17% of 32.12s total
Dropped 982 nodes (cum <= 0.16s)
Showing top 40 nodes out of 308
      flat  flat%   sum%        cum   cum%
     3.04s  9.46%  9.46%      3.04s  9.46%  internal/runtime/syscall/linux.Syscall6
     1.64s  5.11% 14.57%      5.46s 17.00%  runtime.scanObject
     1.63s  5.07% 19.65%      5.70s 17.75%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.59s  4.95% 24.60%      1.59s  4.95%  runtime.memmove
     1.41s  4.39% 28.99%      1.41s  4.39%  runtime.memclrNoHeapPointers
     1.39s  4.33% 33.31%      1.68s  5.23%  runtime.findObject
     1.30s  4.05% 37.36%      1.30s  4.05%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.27s  3.95% 41.31%      1.27s  3.95%  crypto/internal/fips140/sha256.blockAVX2
     1.18s  3.67% 44.99%      1.18s  3.67%  aeshashbody
     1.17s  3.64% 48.63%      1.90s  5.92%  runtime.tryDeferToSpanScan
     0.62s  1.93% 50.56%      0.62s  1.93%  example.com/isureal-company/isureal.(*UserRow).active
     0.59s  1.84% 52.40%      2.03s  6.32%  runtime.mapaccess1_faststr
     0.54s  1.68% 54.08%      0.54s  1.68%  memeqbody
     0.54s  1.68% 55.76%      0.54s  1.68%  runtime.spanClass.sizeclass (inline)
     0.39s  1.21% 56.97%      0.39s  1.21%  runtime.nextFreeFast (inline)
     0.34s  1.06% 58.03%      0.96s  2.99%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.32s     1% 59.03%      0.54s  1.68%  example.com/isureal-company/isureal.postLess
     0.26s  0.81% 59.84%      0.26s  0.81%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.25s  0.78% 60.62%      0.25s  0.78%  sync/atomic.(*Int32).Add
     0.24s  0.75% 61.36%      0.43s  1.34%  database/sql.convertAssignRows
     0.24s  0.75% 62.11%      0.24s  0.75%  indexbody
     0.24s  0.75% 62.86%      1.12s  3.49%  runtime.mallocgcSmallScanNoHeader
     0.23s  0.72% 63.57%      0.53s  1.65%  runtime.greyobject
     0.22s  0.68% 64.26%      0.39s  1.21%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.22s  0.68% 64.94%      0.23s  0.72%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.22s  0.68% 65.63%      0.22s  0.68%  runtime.futex
     0.21s  0.65% 66.28%      0.24s  0.75%  runtime.(*mspan).markBitsForIndex (inline)
     0.18s  0.56% 66.84%      1.66s  5.17%  runtime.mapassign_faststr
     0.16s   0.5% 67.34%      0.19s  0.59%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.16s   0.5% 67.84%      3.97s 12.36%  runtime.mallocgc
     0.15s  0.47% 68.31%      0.20s  0.62%  runtime.(*mspan).writeHeapBitsSmall
     0.12s  0.37% 68.68%      0.17s  0.53%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.12s  0.37% 69.05%      2.71s  8.44%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.12s  0.37% 69.43%      0.22s  0.68%  internal/runtime/maps.(*Iter).Next
     0.10s  0.31% 69.74%      0.53s  1.65%  database/sql.(*Rows).scanLocked
     0.10s  0.31% 70.05%      1.24s  3.86%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.10s  0.31% 70.36%      0.32s     1%  runtime.scanObjectsSmall
     0.09s  0.28% 70.64%      1.28s  3.99%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.09s  0.28% 70.92%      0.33s  1.03%  runtime.mapaccess2_faststr
     0.08s  0.25% 71.17%      1.48s  4.61%  crypto/internal/fips140/sha256.(*Digest).checkSum
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 45.8% / 100% | 30% | 0.2% | 7 | 113 MB | run queue が積もっている |
| isu2 (2 cpu) | 25.3% / 92% | 1% | 1.4% | 4 | 118 MB |  |
| isu3 (2 cpu) | 24.7% / 99% | 1% | 0.8% | 4 | 856 MB |  |
| isu4 (2 cpu) | 53.8% / 100% | 17% | 17.0% | 25 | 102 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 69% / nginx 14% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 45% / kswapd0 1% / nginx 1% |
| isu3 | isureal 47% |
| isu4 | postgres 66% / isureal 2% |

## Diff from previous run

スコア: 80118 → 96430 (+16312)

`28e8717..5c78b9d`

```
 conf/isu1/home/isucon/env.sh |  1 -
 webapp/go/export_api.go      |  2 ++
 webapp/go/replica_server.go  | 24 +++++++++++++++++++++++-
 3 files changed, 25 insertions(+), 2 deletions(-)
```
