# run 20261031-123502-8d603cb

## Summary

| | |
|---|---|
| score | 66783 |
| status | pass |
| ref | feed-author-index |
| sha | 8d603cbc6d49 |
| 所要時間 | 182s |
| 前回比 | +2083 (3.2%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +2083 (3.2%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | feed-author-index |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	45	126	353	735	735
europe	7	14	33	81	81	80
america	6	14	37	68	68	68
```

```
{"id":"8824","status":"done","portalStatus":"finished","passed":true,"score":66783,"raw":66783,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:36:44.316Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:36:47.278Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:36:50.279Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:36:53.301Z level=INFO msg="asia 2 日目: 45 人"
time=2026-10-31T03:36:56.279Z level=INFO msg="europe 2 日目: 14 人"
time=2026-10-31T03:36:59.282Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T03:37:02.305Z level=INFO msg="asia 3 日目: 126 人"
time=2026-10-31T03:37:05.284Z level=INFO msg="europe 3 日目: 33 人"
time=2026-10-31T03:37:05.644Z level=INFO msg="user moved" user=chimotsu.5563 from=asia to=america
time=2026-10-31T03:37:05.831Z level=INFO msg="user moved" user=komo.5893 from=asia to=europe
time=2026-10-31T03:37:08.281Z level=INFO msg="america 3 日目: 37 人"
time=2026-10-31T03:37:11.327Z level=INFO msg="asia 4 日目: 353 人"
time=2026-10-31T03:37:14.332Z level=INFO msg="europe 4 日目: 81 人"
time=2026-10-31T03:37:17.149Z level=INFO msg="user moved" user=solivy-7015 from=asia to=america
time=2026-10-31T03:37:17.326Z level=INFO msg="america 4 日目: 68 人"
time=2026-10-31T03:37:20.561Z level=INFO msg="asia 5 日目: 735 人"
time=2026-10-31T03:37:25.219Z level=INFO msg="europe 5 日目: 81 人"
time=2026-10-31T03:37:26.682Z level=INFO msg="america 5 日目: 68 人"
time=2026-10-31T03:37:35.621Z level=INFO msg="asia 最終日（グランプリ）: 735 人"
time=2026-10-31T03:37:35.785Z level=INFO msg="europe 最終日（グランプリ）: 80 人"
time=2026-10-31T03:37:35.891Z level=INFO msg="america 最終日（グランプリ）: 68 人"
time=2026-10-31T03:37:44.314Z level=INFO msg="グランプリ: アジア 122 人, ヨーロッパ 12 人, アメリカ 10 人"
{
  "pass": true,
  "score": 66783,
  "messages": [
    "ivykivin.9171: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
28.0%	1096	16	0	GET	/api/notifications	9062.940	8.269	15.184	15.321	166
15.6%	1358	53	0	POST	/api/posts	5059.212	3.725	12.498	15.256	629
14.0%	3330	49	0	POST	/api/friend-requests	4534.753	1.362	13.025	18.447	114
9.9%	9024	645	0	GET	^/api/images/[^/]+	3210.175	0.356	7.951	14.883	124624
8.1%	4669	1714	0	POST	/api/auth/login	2626.414	0.563	8.265	14.815	164
6.5%	754	1	0	POST	/api/auth/register	2100.984	2.786	10.420	21.563	227
4.5%	999	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1463.990	1.465	13.712	17.615	116
2.1%	535	112	0	PUT	/api/me/profile-image	689.321	1.288	9.255	16.835	5
1.8%	425	0	0	PUT	^/api/follows/[^/]+$	583.547	1.373	12.628	14.661	0
1.5%	336	0	0	PATCH	/api/me	492.482	1.466	14.811	16.492	278
1.4%	297	4	0	POST	^/api/posts/[^/]+/realmojis$	452.879	1.525	16.061	17.061	479
1.1%	1983	0	0	GET	/api/posting-window/current	368.705	0.186	1.877	5.949	237
1.0%	272	3	0	POST	^/api/posts/[^/]+/comments$	324.514	1.193	12.988	14.195	351
0.8%	2433	0	0	GET	/api/memories	263.350	0.108	0.623	2.083	10059
0.7%	195	0	0	GET	/api/me/export	222.737	1.142	11.021	11.971	1912858
0.6%	6770	0	0	GET	^/api/feeds/[^/]+	197.827	0.029	0.326	4.303	1683
0.6%	434	0	0	POST	/api/auth/logout	179.044	0.413	7.209	13.773	0
0.4%	112	1	0	PATCH	^/api/posts/[^/]+$	130.531	1.165	12.526	14.043	646
0.2%	1609	0	0	GET	/api/users	68.410	0.043	0.464	2.850	237
0.2%	56	0	0	DELETE	^/api/posts/[^/]+$	67.571	1.207	9.885	9.885	0
0.2%	1070	55	0	GET	^/api/posts/me/current$	63.925	0.060	0.623	0.970	622
0.2%	44	0	0	DELETE	^/api/comments/[^/]+$	52.281	1.188	13.347	13.347	0
0.1%	47	0	0	DELETE	^/api/realmojis/[^/]+$	45.036	0.958	8.640	8.640	0
0.1%	433	0	0	GET	^/api/memories/[^/]+	21.561	0.050	0.356	1.387	607
0.1%	1055	0	0	GET	^/api/relationships/[^/]+	20.730	0.020	0.199	1.413	970
0.1%	1019	11	0	GET	^/api/posts/[^/]+$	18.851	0.018	0.120	0.951	623
0.0%	11	0	0	PUT	^/api/admin/users/[^/]+/certification$	15.909	1.446	7.868	7.868	0
0.0%	1	0	0	POST	/api/initialize	12.333	12.333	12.333	12.333	90
0.0%	426	2	0	GET	/api/me	12.147	0.029	0.229	0.464	301
0.0%	447	0	0	GET	^/api/users/[^/]+$	7.933	0.018	0.128	0.217	208
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.3%	201	0	0	GET	/api/notifications	1704.529	8.480	11.331	11.331	180
20.0%	237	3	0	POST	/api/posts	913.843	3.856	11.511	15.069	632
10.7%	428	6	0	POST	/api/friend-requests	487.240	1.138	10.727	13.254	114
6.4%	302	0	0	POST	/api/auth/login	293.107	0.971	8.706	13.023	248
3.8%	1056	8	0	GET	^/api/images/[^/]+	172.393	0.163	0.968	1.426	158316
3.5%	92	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	160.592	1.746	11.287	11.287	116
3.1%	823	0	0	GET	^/api/feeds/[^/]+	142.668	0.173	0.983	3.973	1566
1.8%	318	0	0	GET	/api/posting-window/current	81.602	0.257	1.802	2.801	239
1.6%	49	1	0	POST	^/api/posts/[^/]+/comments$	73.097	1.492	14.579	14.579	293
1.5%	51	0	0	POST	/api/auth/logout	68.545	1.344	12.082	12.082	0
1.4%	305	0	0	GET	/api/memories	61.760	0.202	0.610	2.475	12975
1.1%	27	1	0	POST	^/api/posts/[^/]+/realmojis$	50.460	1.869	14.027	14.027	433
0.9%	78	0	0	POST	/api/auth/register	42.595	0.546	2.552	2.552	228
0.8%	41	0	0	PUT	/api/me/profile-image	36.063	0.880	7.253	7.253	0
0.8%	15	0	0	GET	/api/me/export	35.819	2.388	11.430	11.430	2154828
0.6%	198	0	0	GET	/api/users	28.243	0.143	0.619	0.871	229
0.6%	185	17	0	GET	^/api/posts/me/current$	27.201	0.147	0.419	0.738	582
0.5%	49	0	0	PUT	^/api/follows/[^/]+$	23.419	0.478	6.097	6.097	0
0.5%	10	0	0	DELETE	^/api/posts/[^/]+$	22.616	2.262	12.274	12.274	0
0.4%	20	2	0	PATCH	^/api/posts/[^/]+$	20.495	1.025	4.153	4.153	574
0.4%	102	2	0	GET	^/api/posts/[^/]+$	17.396	0.171	0.983	1.029	623
0.4%	120	0	0	GET	^/api/relationships/[^/]+	17.315	0.144	0.440	0.532	796
0.3%	2	0	0	DELETE	^/api/blocks/[^/]+$	14.923	7.462	14.589	14.589	0
0.3%	29	0	0	PATCH	/api/me	14.721	0.508	3.543	3.543	279
0.3%	17	0	0	DELETE	^/api/comments/[^/]+$	12.187	0.717	7.438	7.438	0
0.2%	54	0	0	GET	^/api/users/[^/]+$	10.197	0.189	1.086	1.086	220
0.1%	47	2	0	GET	/api/me	6.424	0.137	0.282	0.282	283
0.1%	35	0	0	GET	^/api/memories/[^/]+	6.399	0.183	0.449	0.449	1517
0.1%	4	0	0	DELETE	^/api/realmojis/[^/]+$	6.305	1.576	5.061	5.061	0
0.1%	39	0	0	GET	^/api/posts/[^/]+/realmojis$	5.733	0.147	0.309	0.309	706
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.0%	167	0	0	GET	/api/notifications	1308.500	7.835	9.367	9.497	181
15.7%	183	3	0	POST	/api/posts	588.020	3.213	11.560	11.973	632
12.5%	347	4	0	POST	/api/friend-requests	468.518	1.350	12.147	14.691	114
6.6%	235	0	0	POST	/api/auth/login	244.719	1.041	9.783	11.340	251
6.1%	130	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	227.034	1.746	11.837	11.963	116
4.0%	839	9	0	GET	^/api/images/[^/]+	148.086	0.177	1.059	1.287	183019
2.9%	641	0	0	GET	^/api/feeds/[^/]+	107.609	0.168	0.964	2.151	2455
2.2%	23	3	0	PATCH	^/api/posts/[^/]+$	83.219	3.618	13.400	13.400	571
1.5%	18	0	0	GET	/api/me/export	55.997	3.111	14.563	14.563	3113988
1.4%	247	0	0	GET	/api/posting-window/current	53.707	0.217	1.447	3.228	240
1.3%	238	0	0	GET	/api/memories	49.740	0.209	1.001	1.466	12476
1.3%	38	0	0	POST	/api/auth/logout	48.650	1.280	10.488	10.488	0
1.3%	29	0	0	PATCH	/api/me	47.416	1.635	10.189	10.189	265
1.1%	38	1	0	POST	^/api/posts/[^/]+/comments$	41.003	1.079	15.820	15.820	314
0.9%	28	0	0	PUT	^/api/follows/[^/]+$	34.956	1.248	6.860	6.860	0
0.8%	149	15	0	GET	^/api/posts/me/current$	29.091	0.195	2.342	3.286	580
0.7%	163	0	0	GET	/api/users	26.695	0.164	1.057	1.400	231
0.7%	27	0	0	PUT	/api/me/profile-image	26.034	0.964	6.379	6.379	0
0.6%	65	0	0	POST	/api/auth/register	21.232	0.327	1.405	1.405	229
0.6%	15	3	0	POST	^/api/posts/[^/]+/realmojis$	21.135	1.409	8.967	8.967	393
0.6%	120	0	0	GET	^/api/relationships/[^/]+	20.802	0.173	0.945	2.340	1312
0.5%	9	0	0	DELETE	^/api/posts/[^/]+$	19.895	2.211	7.550	7.550	0
0.4%	85	3	0	GET	^/api/posts/[^/]+$	15.047	0.177	1.254	1.254	603
0.2%	5	0	0	PUT	^/api/blocks/[^/]+$	8.462	1.692	7.400	7.400	0
0.2%	2	0	0	DELETE	^/api/realmojis/[^/]+$	7.541	3.770	6.943	6.943	0
0.2%	32	0	0	GET	^/api/memories/[^/]+	7.086	0.221	0.888	0.888	1419
0.2%	38	0	0	GET	^/api/posts/[^/]+/realmojis$	6.191	0.163	1.361	1.361	573
0.2%	31	0	0	GET	^/api/users/[^/]+$	6.078	0.196	1.614	1.614	211
0.1%	33	0	0	GET	^/api/posts/[^/]+/comments$	4.682	0.142	0.361	0.361	1061
0.1%	30	0	0	GET	/api/me	4.160	0.139	0.613	0.613	299
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 45606 クエリ / 81 種類 / 合計 299.779s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	59.5%	2548	178.379	70.01	541.48	637.3	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	26.3%	228	78.846	345.82	2625.04	6888.9	129.8	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	4.3%	1660	12.898	7.77	249.14	117.1	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.2%	5871	12.662	2.16	228.36	10.4	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	2.0%	1724	6.120	3.55	70.46	26.4	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.7%	1	2.228	2227.90	2227.90	29985.0	1199780.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.6%	1	1.674	1673.67	1673.67	21469.0	626756.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.4%	331	1.265	3.82	19.77	32.0	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.3%	4030	0.825	0.20	37.89	20.1	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	0.3%	1	0.755	755.24	755.24	11684.0	530025.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
11	0.2%	354	0.614	1.73	14.97	21.8	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.1%	3492	0.397	0.11	8.26	11.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
13	0.1%	2	0.332	165.87	175.21	223825.5	2634.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.1%	1187	0.309	0.26	22.77	21.2	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	2	0.262	130.91	253.82	16782.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: 0bc0416e5e6396720e14d8b051cebad0b0cb2165
Type: cpu
Time: 2026-10-31 12:36:26 JST
Duration: 60.04s, Total samples = 51.11s (85.12%)
Showing nodes accounting for 39.70s, 77.68% of 51.11s total
Dropped 1004 nodes (cum <= 0.26s)
Showing top 40 nodes out of 281
      flat  flat%   sum%        cum   cum%
     3.44s  6.73%  6.73%     11.41s 22.32%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     3.44s  6.73% 13.46%      3.44s  6.73%  internal/runtime/syscall/linux.Syscall6
     3.20s  6.26% 19.72%      4.87s  9.53%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     3.01s  5.89% 25.61%      3.10s  6.07%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     2.66s  5.20% 30.82%      2.66s  5.20%  runtime.memclrNoHeapPointers
     2.65s  5.18% 36.00%      9.02s 17.65%  runtime.scanObject
     2.22s  4.34% 40.34%      2.71s  5.30%  runtime.findObject
     2.13s  4.17% 44.51%      2.13s  4.17%  runtime.memmove
     1.81s  3.54% 48.05%      3.12s  6.10%  runtime.tryDeferToSpanScan
     1.78s  3.48% 51.54%      1.78s  3.48%  crypto/internal/fips140/sha256.blockAVX2
     1.77s  3.46% 55.00%      5.65s 11.05%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.33s  2.60% 57.60%      1.33s  2.60%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.22s  2.39% 59.99%      1.72s  3.37%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     1.04s  2.03% 62.02%      1.04s  2.03%  aeshashbody
     0.98s  1.92% 63.94%      0.98s  1.92%  runtime.spanClass.sizeclass (inline)
     0.76s  1.49% 65.43%      0.76s  1.49%  example.com/isureal-company/isureal.(*UserRow).active
     0.51s     1% 66.43%      1.82s  3.56%  runtime.mapaccess1_faststr
     0.49s  0.96% 67.38%      0.49s  0.96%  memeqbody
     0.45s  0.88% 68.26%      0.49s  0.96%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.42s  0.82% 69.09%      0.42s  0.82%  runtime.nextFreeFast (inline)
     0.40s  0.78% 69.87%      0.40s  0.78%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.36s   0.7% 70.57%      0.36s   0.7%  indexbody
     0.35s  0.68% 71.26%      0.59s  1.15%  example.com/isureal-company/isureal.postLess
     0.35s  0.68% 71.94%      0.82s  1.60%  runtime.greyobject
     0.29s  0.57% 72.51%      5.93s 11.60%  runtime.mallocgc
     0.27s  0.53% 73.04%      0.27s  0.53%  runtime.typePointers.nextFast (inline)
     0.25s  0.49% 73.53%      1.07s  2.09%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.47% 74.00%      0.59s  1.15%  database/sql.convertAssignRows
     0.24s  0.47% 74.47%      0.33s  0.65%  runtime.(*mspan).markBitsForIndex (inline)
     0.23s  0.45% 74.92%      1.61s  3.15%  runtime.mapassign_faststr
     0.22s  0.43% 75.35%      0.76s  1.49%  runtime.scanObjectsSmall
     0.19s  0.37% 75.72%      0.53s  1.04%  runtime.mapaccess2_faststr
     0.18s  0.35% 76.07%      0.41s   0.8%  compress/flate.(*compressor).deflate
     0.14s  0.27% 76.35%      2.06s  4.03%  crypto/internal/fips140/sha256.(*Digest).Write
     0.13s  0.25% 76.60%      2.12s  4.15%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.13s  0.25% 76.85%      3.34s  6.53%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.13s  0.25% 77.11%      0.32s  0.63%  runtime.scanObjectSmall
     0.11s  0.22% 77.32%      1.20s  2.35%  runtime.scanSpan
     0.09s  0.18% 77.50%      1.21s  2.37%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.09s  0.18% 77.68%      0.55s  1.08%  runtime.mallocgcSmallNoscan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 54.9% / 100% | 37% | 0.4% | 6 | 151 MB | CPU が飽和 |
| isu2 (2 cpu) | 12.0% / 33% | 0% | 0.4% | 3 | 2666 MB |  |
| isu3 (2 cpu) | 8.3% / 32% | 0% | 0.5% | 4 | 2735 MB |  |
| isu4 (2 cpu) | 46.2% / 100% | 18% | 15.4% | 36 | 120 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 89% / nginx 13% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 22% |
| isu3 | isureal 14% / nginx 1% |
| isu4 | postgres 49% / isureal 9% |

## Diff from previous run

スコア: 64700 → 66783 (+2083)

`e0e4cb3..8d603cb`

```
 webapp/go/account_deletion_api.go |   2 +-
 webapp/go/admin_user_api.go       |   2 +-
 webapp/go/comment_api.go          |   2 +-
 webapp/go/feed_api.go             | 113 +++++++++----
 webapp/go/feed_test.go            | 331 ++++++++++++++++++++++++++++++++++++++
 webapp/go/realmoji_api.go         |   2 +-
 webapp/go/relationship.go         |  10 +-
 webapp/go/store.go                | 113 +++++++++++--
 8 files changed, 526 insertions(+), 49 deletions(-)
```
