# run 20261031-124715-46aeaca

## Summary

| | |
|---|---|
| score | 100660 |
| status | pass |
| ref | 46aeaca |
| sha | 46aeacab4c1a |
| 所要時間 | 219s |
| 前回比 | +17976 (21.7%) vs 20261031-124309-e42d0b5 |
| ベスト比 | +17976 (21.7%) vs 20261031-124309-e42d0b5 = 82684 **ベスト更新** |
| note | PR73 head merged main 04641dc |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	149	416	924	923
europe	7	16	42	99	99	98
america	6	15	39	72	76	76
```

```
{"id":"9807","status":"done","portalStatus":"finished","passed":true,"score":100660,"raw":100660,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:49:23.556Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:49:26.520Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:49:29.519Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:49:32.542Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T03:49:35.521Z level=INFO msg="europe 2 日目: 16 人"
time=2026-10-31T03:49:38.523Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T03:49:41.544Z level=INFO msg="asia 3 日目: 149 人"
time=2026-10-31T03:49:43.958Z level=INFO msg="user moved" user=koshimu5842 from=asia to=europe
time=2026-10-31T03:49:44.541Z level=INFO msg="europe 3 日目: 42 人"
time=2026-10-31T03:49:47.521Z level=INFO msg="america 3 日目: 39 人"
time=2026-10-31T03:49:50.558Z level=INFO msg="asia 4 日目: 416 人"
time=2026-10-31T03:49:53.661Z level=INFO msg="europe 4 日目: 99 人"
time=2026-10-31T03:49:54.879Z level=INFO msg="user moved" user=hayais6431 from=america to=asia
time=2026-10-31T03:49:56.672Z level=INFO msg="america 4 日目: 72 人"
time=2026-10-31T03:49:57.736Z level=INFO msg="user moved" user=solia7619 from=asia to=america
time=2026-10-31T03:49:58.252Z level=INFO msg="user moved" user=toran8131 from=asia to=america
time=2026-10-31T03:49:58.371Z level=INFO msg="user moved" user=hayais6431 from=asia to=europe
time=2026-10-31T03:49:59.367Z level=INFO msg="user moved" user=tsurari_6894 from=asia to=america
time=2026-10-31T03:49:59.652Z level=INFO msg="user moved" user=koshimu5842 from=europe to=asia
time=2026-10-31T03:49:59.684Z level=INFO msg="asia 5 日目: 924 人"
time=2026-10-31T03:49:59.942Z level=INFO msg="user moved" user=rovin.7275 from=asia to=america
time=2026-10-31T03:50:02.667Z level=INFO msg="europe 5 日目: 99 人"
time=2026-10-31T03:50:05.606Z level=INFO msg="america 5 日目: 76 人"
time=2026-10-31T03:50:14.716Z level=INFO msg="asia 最終日（グランプリ）: 923 人"
time=2026-10-31T03:50:14.791Z level=INFO msg="europe 最終日（グランプリ）: 98 人"
time=2026-10-31T03:50:14.857Z level=INFO msg="america 最終日（グランプリ）: 76 人"
time=2026-10-31T03:50:23.563Z level=INFO msg="グランプリ: アジア 361 人, ヨーロッパ 13 人, アメリカ 7 人"
{
  "pass": true,
  "score": 100660,
  "messages": [
    "muma5025: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.7%	1490	1	0	GET	/api/notifications	12148.647	8.153	15.058	15.092	169
17.1%	2163	59	0	POST	/api/posts	6344.894	2.933	9.962	13.857	636
13.3%	4930	57	0	POST	/api/friend-requests	4942.863	1.003	7.737	12.255	114
8.3%	12625	737	0	GET	^/api/images/[^/]+	3091.456	0.245	4.196	15.909	142241
6.9%	5911	2095	0	POST	/api/auth/login	2554.098	0.432	6.004	12.116	167
5.2%	1634	8	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1921.516	1.176	8.396	12.186	116
2.7%	943	1	0	POST	/api/auth/register	1001.918	1.062	6.199	11.536	227
2.3%	653	122	0	PUT	/api/me/profile-image	836.979	1.282	7.867	10.073	4
2.0%	562	35	0	POST	^/api/posts/[^/]+/realmojis$	743.576	1.323	9.478	10.457	453
1.4%	480	28	0	POST	^/api/posts/[^/]+/comments$	531.649	1.108	9.235	13.479	340
1.4%	581	0	0	PUT	^/api/follows/[^/]+$	520.349	0.896	7.360	13.577	0
1.2%	459	0	0	PATCH	/api/me	435.692	0.949	5.175	7.045	276
1.0%	3151	0	0	GET	/api/memories	380.688	0.121	0.614	1.359	10820
0.8%	632	0	0	POST	/api/auth/logout	288.196	0.456	5.405	11.521	0
0.8%	11174	0	0	GET	^/api/feeds/[^/]+	285.525	0.026	0.202	0.458	1859
0.8%	185	6	0	PATCH	^/api/posts/[^/]+$	284.858	1.540	9.386	9.646	631
0.7%	233	0	0	GET	/api/me/export	252.941	1.086	6.838	8.873	1932902
0.3%	3126	0	0	GET	/api/posting-window/current	107.493	0.034	0.170	3.638	236
0.3%	57	0	0	DELETE	^/api/posts/[^/]+$	98.838	1.734	9.012	9.012	0
0.2%	81	0	0	DELETE	^/api/realmojis/[^/]+$	62.495	0.772	6.871	6.871	0
0.2%	1712	46	0	GET	^/api/posts/me/current$	60.203	0.035	0.205	0.316	638
0.1%	2396	0	0	GET	/api/users	51.701	0.022	0.142	0.283	237
0.1%	1505	0	0	GET	^/api/relationships/[^/]+	34.745	0.023	0.214	0.585	869
0.1%	530	0	0	GET	^/api/memories/[^/]+	28.758	0.054	0.372	0.736	631
0.1%	73	0	0	DELETE	^/api/comments/[^/]+$	28.479	0.390	2.451	2.451	0
0.1%	1547	15	0	GET	^/api/posts/[^/]+$	26.169	0.017	0.154	0.266	622
0.1%	35	0	0	PUT	^/api/blocks/[^/]+$	25.640	0.733	5.215	5.215	0
0.1%	9	0	0	DELETE	^/api/friendships/[^/]+$	20.049	2.228	5.936	5.936	0
0.0%	1	0	0	POST	/api/initialize	15.964	15.964	15.964	15.964	90
0.0%	774	0	0	GET	^/api/users/[^/]+$	15.460	0.020	0.186	0.250	199
…（上位 30 / 全 43 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.4%	255	1	0	GET	/api/notifications	2132.428	8.362	11.811	11.830	178
17.3%	294	2	0	POST	/api/posts	957.893	3.258	10.493	10.867	647
8.8%	592	6	0	POST	/api/friend-requests	490.349	0.828	6.462	11.674	114
6.6%	261	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	366.672	1.405	9.256	11.895	116
6.1%	1391	10	0	GET	^/api/images/[^/]+	337.541	0.243	1.410	2.033	161279
5.3%	395	0	0	POST	/api/auth/login	291.585	0.738	5.524	10.830	247
3.6%	1188	0	0	GET	^/api/feeds/[^/]+	199.224	0.168	0.560	1.084	2334
1.8%	388	0	0	GET	/api/memories	102.140	0.263	1.290	1.654	14287
1.6%	408	0	0	GET	/api/posting-window/current	87.586	0.215	0.551	0.884	239
1.2%	30	0	0	GET	/api/me/export	66.247	2.208	13.345	13.345	2108341
0.9%	60	0	0	POST	/api/auth/logout	51.393	0.857	7.160	7.160	0
0.9%	35	3	0	POST	^/api/posts/[^/]+/realmojis$	50.993	1.457	12.242	12.242	438
0.9%	291	0	0	GET	/api/users	47.714	0.164	0.529	0.630	243
0.8%	96	0	0	POST	/api/auth/register	45.846	0.478	2.054	2.054	228
0.8%	229	9	0	GET	^/api/posts/me/current$	41.632	0.182	0.569	0.668	626
0.7%	37	2	0	POST	^/api/posts/[^/]+/comments$	39.249	1.061	6.251	6.251	321
0.6%	28	2	0	PATCH	^/api/posts/[^/]+$	34.035	1.216	5.564	5.564	599
0.5%	51	0	0	PUT	^/api/follows/[^/]+$	29.469	0.578	3.483	3.483	0
0.5%	153	8	0	GET	^/api/posts/[^/]+$	27.506	0.180	0.504	3.221	598
0.5%	177	0	0	GET	^/api/relationships/[^/]+	25.786	0.146	0.545	0.605	1480
0.5%	52	0	0	PUT	/api/me/profile-image	25.417	0.489	1.895	1.895	0
0.4%	49	0	0	PATCH	/api/me	21.923	0.447	1.897	1.897	278
0.2%	83	0	0	GET	^/api/users/[^/]+$	13.353	0.161	0.545	0.545	192
0.2%	53	0	0	GET	^/api/memories/[^/]+	12.036	0.227	0.907	0.907	824
0.2%	76	0	0	GET	^/api/posts/[^/]+/realmojis$	11.673	0.154	0.442	0.442	886
0.2%	60	0	0	GET	^/api/posts/[^/]+/comments$	8.441	0.141	0.426	0.426	1057
0.1%	57	2	0	GET	/api/me	8.014	0.141	0.244	0.244	290
0.1%	10	0	0	DELETE	^/api/posts/[^/]+$	6.853	0.685	2.238	2.238	0
0.1%	5	0	0	DELETE	^/api/realmojis/[^/]+$	5.249	1.050	2.820	2.820	0
0.1%	1	0	0	DELETE	^/api/friendships/[^/]+$	4.943	4.943	4.943	4.943	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.6%	198	1	0	GET	/api/notifications	1471.599	7.432	9.247	9.315	176
17.7%	230	6	0	POST	/api/posts	710.650	3.090	9.732	9.942	630
8.6%	384	5	0	POST	/api/friend-requests	344.176	0.896	8.225	12.323	114
6.1%	182	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	244.502	1.343	8.636	9.255	115
6.0%	274	0	0	POST	/api/auth/login	240.251	0.877	8.589	12.061	250
5.9%	1027	7	0	GET	^/api/images/[^/]+	237.260	0.231	1.193	2.052	164589
2.9%	675	0	0	GET	^/api/feeds/[^/]+	115.965	0.172	0.534	1.199	2181
1.9%	304	0	0	GET	/api/memories	77.249	0.254	1.123	1.196	14756
1.6%	53	0	0	PUT	^/api/follows/[^/]+$	65.587	1.237	7.415	7.415	0
1.6%	311	0	0	GET	/api/posting-window/current	65.441	0.210	0.571	1.200	240
1.6%	77	7	0	POST	^/api/posts/[^/]+/comments$	62.261	0.809	6.788	6.788	272
1.3%	18	0	0	GET	/api/me/export	53.580	2.977	7.917	7.917	3561524
1.1%	34	5	0	POST	^/api/posts/[^/]+/realmojis$	44.176	1.299	8.197	8.197	456
0.9%	184	19	0	GET	^/api/posts/me/current$	34.515	0.188	0.552	0.580	584
0.8%	193	0	0	GET	/api/users	32.436	0.168	0.661	0.709	236
0.8%	37	0	0	PUT	/api/me/profile-image	31.401	0.849	4.989	4.989	0
0.7%	72	0	0	POST	/api/auth/register	28.830	0.400	2.110	2.110	229
0.6%	37	0	0	POST	/api/auth/logout	24.513	0.663	6.289	6.289	0
0.5%	127	0	0	GET	^/api/relationships/[^/]+	19.369	0.153	0.536	0.537	1405
0.4%	33	0	0	PATCH	/api/me	16.294	0.494	2.249	2.249	289
0.4%	14	3	0	PATCH	^/api/posts/[^/]+$	16.061	1.147	5.534	5.534	489
0.4%	92	5	0	GET	^/api/posts/[^/]+$	14.943	0.162	0.640	0.640	585
0.3%	7	0	0	DELETE	^/api/posts/[^/]+$	12.651	1.807	3.043	3.043	0
0.2%	37	0	0	GET	^/api/memories/[^/]+	7.778	0.210	0.986	0.986	958
0.2%	4	0	0	DELETE	^/api/realmojis/[^/]+$	6.718	1.679	5.215	5.215	0
0.2%	41	1	0	GET	/api/me	6.666	0.163	0.581	0.581	291
0.2%	12	0	0	DELETE	^/api/comments/[^/]+$	6.598	0.550	4.259	4.259	0
0.1%	46	0	0	GET	^/api/posts/[^/]+/comments$	5.991	0.130	0.287	0.287	1713
0.1%	36	0	0	GET	^/api/users/[^/]+$	5.545	0.154	0.507	0.507	188
0.1%	1	0	0	POST	/api/me/account-deletion	5.487	5.487	5.487	5.487	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 63900 クエリ / 82 種類 / 合計 461.721s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	59.4%	3831	274.229	71.58	515.61	646.3	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	25.9%	281	119.612	425.66	2036.26	7204.3	131.5	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	4.7%	2116	21.635	10.22	422.48	117.4	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.7%	8006	21.563	2.69	42.35	10.5	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.6%	2622	7.255	2.77	66.41	25.1	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.8%	2	3.892	1945.95	3739.42	223995.5	2635.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.5%	588	2.450	4.17	39.47	29.6	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
8	0.5%	1	2.174	2173.87	2173.87	30159.0	1199848.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
9	0.4%	1	1.715	1715.38	1715.38	21909.0	626671.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
10	0.3%	5828	1.218	0.21	41.19	19.9	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.2%	557	1.044	1.87	12.47	21.2	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.2%	1	0.736	736.41	736.41	11845.0	529973.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
13	0.1%	4485	0.636	0.14	52.06	11.2	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
14	0.1%	2007	0.580	0.29	29.56	21.0	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	2	0.324	161.98	315.83	16769.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: d86f0901e1d74d9dbc1d8d578a6c085da10f36aa
Type: cpu
Time: 2026-10-31 12:49:02 JST
Duration: 60.02s, Total samples = 42.37s (70.60%)
Showing nodes accounting for 30.14s, 71.14% of 42.37s total
Dropped 1051 nodes (cum <= 0.21s)
Showing top 40 nodes out of 308
      flat  flat%   sum%        cum   cum%
     3.91s  9.23%  9.23%      3.91s  9.23%  internal/runtime/syscall/linux.Syscall6
     2.31s  5.45% 14.68%      2.31s  5.45%  crypto/internal/fips140/sha256.blockAVX2
     2.13s  5.03% 19.71%      7.76s 18.31%  runtime.scanObject
     2.02s  4.77% 24.47%      2.02s  4.77%  runtime.memmove
     1.94s  4.58% 29.05%      1.94s  4.58%  runtime.memclrNoHeapPointers
     1.87s  4.41% 33.47%      2.44s  5.76%  runtime.findObject
     1.59s  3.75% 37.22%      5.63s 13.29%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.52s  3.59% 40.81%      2.69s  6.35%  runtime.tryDeferToSpanScan
     1.42s  3.35% 44.16%      1.42s  3.35%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.20s  2.83% 46.99%      1.20s  2.83%  aeshashbody
     0.94s  2.22% 49.21%      0.94s  2.22%  example.com/isureal-company/isureal.(*UserRow).active
     0.78s  1.84% 51.05%      0.78s  1.84%  runtime.spanClass.sizeclass (inline)
     0.75s  1.77% 52.82%      0.75s  1.77%  memeqbody
     0.71s  1.68% 54.50%      2.22s  5.24%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.53s  1.25% 55.75%      0.54s  1.27%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.48s  1.13% 56.88%      0.97s  2.29%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.47s  1.11% 57.99%      2.09s  4.93%  runtime.mapaccess1_faststr
     0.43s  1.01% 59.00%      0.43s  1.01%  indexbody
     0.38s   0.9% 59.90%      0.50s  1.18%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.37s  0.87% 60.77%      0.54s  1.27%  example.com/isureal-company/isureal.postLess
     0.37s  0.87% 61.65%      0.37s  0.87%  runtime.nextFreeFast (inline)
     0.35s  0.83% 62.47%      0.73s  1.72%  database/sql.convertAssignRows
     0.29s  0.68% 63.16%      0.72s  1.70%  runtime.greyobject
     0.29s  0.68% 63.84%      0.29s  0.68%  runtime.typePointers.nextFast (inline)
     0.27s  0.64% 64.48%      0.27s  0.64%  crypto/internal/fips140/aes/gcm.gcmAesEnc
     0.25s  0.59% 65.07%      0.25s  0.59%  runtime.futex
     0.25s  0.59% 65.66%      1.06s  2.50%  runtime.mallocgcSmallScanNoHeader
     0.23s  0.54% 66.20%      1.96s  4.63%  runtime.mapassign_faststr
     0.20s  0.47% 66.67%      0.25s  0.59%  encoding/json.appendString[go.shape.string]
     0.20s  0.47% 67.15%      0.28s  0.66%  runtime.(*mspan).markBitsForIndex (inline)
     0.20s  0.47% 67.62%      0.24s  0.57%  runtime.typePointers.next
     0.19s  0.45% 68.07%      0.42s  0.99%  compress/flate.(*compressor).deflate
     0.19s  0.45% 68.52%      2.62s  6.18%  crypto/internal/fips140/sha256.(*Digest).Write
     0.19s  0.45% 68.96%      0.32s  0.76%  runtime.heapArenaOf (inline)
     0.19s  0.45% 69.41%      0.22s  0.52%  runtime.spanOf (inline)
     0.16s  0.38% 69.79%      0.47s  1.11%  runtime.scanObjectsSmall
     0.15s  0.35% 70.14%      5.04s 11.90%  runtime.mallocgc
     0.15s  0.35% 70.50%      0.52s  1.23%  runtime.mapaccess2_faststr
     0.14s  0.33% 70.83%      0.84s  1.98%  runtime.slicebytetostring
     0.13s  0.31% 71.14%      1.90s  4.48%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 46.2% / 100% | 32% | 0.2% | 9 | 142 MB | CPU が飽和 |
| isu2 (2 cpu) | 18.2% / 61% | 0% | 0.4% | 2 | 2574 MB |  |
| isu3 (2 cpu) | 16.0% / 66% | 0% | 0.4% | 4 | 2668 MB |  |
| isu4 (2 cpu) | 50.7% / 100% | 28% | 15.8% | 39 | 119 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 69% / nginx 15% / journalctl 2% / systemd-journal 2% |
| isu2 | isureal 34% |
| isu3 | isureal 30% / nginx 1% |
| isu4 | postgres 62% / isureal 3% |

## Diff from previous run

スコア: 82684 → 100660 (+17976)

`e42d0b5..46aeaca`

```
 webapp/go/feed_api.go  | 113 +++++++++++++----
 webapp/go/feed_test.go | 331 +++++++++++++++++++++++++++++++++++++++++++++++++
 webapp/go/offload.go   |   6 +-
 webapp/go/post_api.go  |  62 +++++++--
 webapp/go/store.go     |  13 ++
 5 files changed, 485 insertions(+), 40 deletions(-)
```
