# run 20261031-120609-9f29066

## Summary

| | |
|---|---|
| score | 68346 |
| status | pass |
| ref | relation-pair-locks |
| sha | 9f29066a25e1 |
| 所要時間 | 253s |
| 前回比 | +68346 (—) vs 20261031-120517-9f29066 |
| ベスト比 | +8145 (13.5%) vs 20261031-115040-1c73ee4 = 60201 **ベスト更新** |
| note | relation pair locks |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	53	156	440	535	535
europe	7	18	47	109	108	107
america	6	13	35	62	63	63
```

```
{"id":"6824","status":"done","portalStatus":"finished","passed":true,"score":68346,"raw":68346,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:08:49.074Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:08:52.038Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:08:55.039Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:08:58.068Z level=INFO msg="asia 2 日目: 53 人"
time=2026-10-31T03:09:01.097Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T03:09:04.038Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T03:09:07.070Z level=INFO msg="asia 3 日目: 156 人"
time=2026-10-31T03:09:08.885Z level=INFO msg="user moved" user=tekedan.5538 from=asia to=america
time=2026-10-31T03:09:09.713Z level=INFO msg="user moved" user=rimumia-5765 from=asia to=europe
time=2026-10-31T03:09:10.068Z level=INFO msg="europe 3 日目: 47 人"
time=2026-10-31T03:09:10.079Z level=INFO msg="user moved" user=hikemu5757 from=asia to=europe
time=2026-10-31T03:09:10.367Z level=INFO msg="user moved" user=teke.6048 from=asia to=america
time=2026-10-31T03:09:10.397Z level=INFO msg="user moved" user=tesolnu6028 from=asia to=europe
time=2026-10-31T03:09:12.634Z level=INFO msg="user moved" user=teke.6048 from=america to=europe
time=2026-10-31T03:09:13.039Z level=INFO msg="america 3 日目: 35 人"
time=2026-10-31T03:09:16.125Z level=INFO msg="asia 4 日目: 440 人"
time=2026-10-31T03:09:19.142Z level=INFO msg="europe 4 日目: 109 人"
time=2026-10-31T03:09:22.327Z level=INFO msg="america 4 日目: 62 人"
time=2026-10-31T03:09:24.201Z level=INFO msg="user moved" user=hemax.6773 from=asia to=europe
time=2026-10-31T03:09:25.186Z level=INFO msg="asia 5 日目: 535 人"
time=2026-10-31T03:09:27.015Z level=INFO msg="user moved" user=teke.6048 from=europe to=america
time=2026-10-31T03:09:28.194Z level=INFO msg="user moved" user=tesolnu6028 from=europe to=asia
time=2026-10-31T03:09:28.696Z level=INFO msg="europe 5 日目: 108 人"
time=2026-10-31T03:09:31.329Z level=INFO msg="america 5 日目: 63 人"
time=2026-10-31T03:09:40.331Z level=INFO msg="europe 最終日（グランプリ）: 107 人"
time=2026-10-31T03:09:40.712Z level=INFO msg="asia 最終日（グランプリ）: 535 人"
time=2026-10-31T03:09:40.824Z level=INFO msg="america 最終日（グランプリ）: 63 人"
time=2026-10-31T03:09:49.056Z level=INFO msg="グランプリ: アジア 8 人, ヨーロッパ 1 人, アメリカ 0 人"
{
  "pass": true,
  "score": 68346,
  "messages": [
    "anlia_5041: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.3%	1094	1	0	GET	/api/notifications	10018.918	9.158	15.512	15.626	173
20.3%	1446	57	0	POST	/api/posts	6114.981	4.229	14.202	17.379	626
10.8%	3062	66	0	POST	/api/friend-requests	3256.783	1.064	12.112	15.849	113
8.9%	8725	525	0	GET	^/api/images/[^/]+	2678.029	0.307	6.272	15.579	135165
7.0%	3430	833	0	POST	/api/auth/login	2117.836	0.617	9.679	17.000	191
4.9%	1118	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1472.046	1.317	13.996	16.313	115
1.9%	403	23	0	POST	^/api/posts/[^/]+/comments$	569.390	1.413	15.587	15.968	309
1.9%	389	22	0	POST	^/api/posts/[^/]+/realmojis$	559.946	1.439	14.834	16.401	465
1.7%	557	1	0	POST	/api/auth/register	521.316	0.936	6.222	12.619	227
1.6%	1992	0	0	GET	/api/posting-window/current	469.469	0.236	3.599	5.231	234
1.5%	441	131	0	PUT	/api/me/profile-image	447.638	1.015	7.630	11.039	7
0.8%	302	0	0	PUT	^/api/follows/[^/]+$	232.265	0.769	5.385	8.542	0
0.7%	256	0	0	PATCH	/api/me	214.459	0.838	10.112	16.394	279
0.7%	7186	0	0	GET	^/api/feeds/[^/]+	210.285	0.029	0.178	1.141	2301
0.7%	2227	0	0	GET	/api/memories	207.684	0.093	0.438	4.121	10328
0.7%	427	0	0	POST	/api/auth/logout	203.273	0.476	9.381	15.884	0
0.6%	173	0	0	GET	/api/me/export	176.773	1.022	12.725	16.210	2077212
0.3%	55	0	0	DELETE	^/api/realmojis/[^/]+$	90.496	1.645	13.597	13.597	0
0.3%	101	11	0	PATCH	^/api/posts/[^/]+$	85.281	0.844	4.740	7.247	584
0.2%	68	0	0	DELETE	^/api/comments/[^/]+$	65.965	0.970	15.276	15.276	0
0.2%	55	0	0	DELETE	^/api/posts/[^/]+$	58.676	1.067	13.347	13.347	0
0.2%	1459	0	0	GET	/api/users	56.223	0.039	0.246	2.629	243
0.1%	1040	64	0	GET	^/api/posts/me/current$	44.304	0.043	0.270	3.107	612
0.1%	1110	14	0	GET	^/api/posts/[^/]+$	32.208	0.029	0.181	1.241	609
0.1%	1031	0	0	GET	^/api/relationships/[^/]+	25.448	0.025	0.180	1.065	965
0.1%	14	0	0	PUT	^/api/admin/users/[^/]+/certification$	25.247	1.803	14.162	14.162	0
0.1%	384	0	0	GET	^/api/memories/[^/]+	24.869	0.065	0.335	2.792	648
0.1%	18	0	0	PUT	^/api/blocks/[^/]+$	18.728	1.040	4.712	4.712	0
0.0%	543	0	0	GET	^/api/users/[^/]+$	14.812	0.027	0.147	0.262	191
0.0%	434	0	0	GET	^/api/posts/[^/]+/comments$	11.447	0.026	0.171	0.276	1335
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.8%	255	1	0	GET	/api/notifications	2086.435	8.182	11.717	11.718	178
23.3%	328	8	0	POST	/api/posts	1479.838	4.512	14.610	17.045	635
10.5%	531	8	0	POST	/api/friend-requests	666.859	1.256	12.140	16.208	114
6.9%	370	0	0	POST	/api/auth/login	435.778	1.178	13.993	16.243	246
3.7%	1424	15	0	GET	^/api/images/[^/]+	236.468	0.166	1.057	1.591	200474
3.5%	164	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	225.100	1.373	13.576	13.604	116
3.1%	1240	0	0	GET	^/api/feeds/[^/]+	197.023	0.159	0.346	2.039	2962
2.2%	443	0	0	GET	/api/posting-window/current	142.752	0.322	3.755	5.595	239
1.7%	116	13	0	POST	^/api/posts/[^/]+/comments$	110.012	0.948	7.080	9.589	308
1.2%	399	0	0	GET	/api/memories	79.052	0.198	1.058	2.309	10068
1.2%	58	0	0	POST	/api/auth/logout	73.590	1.269	12.855	12.855	0
1.1%	103	0	0	POST	/api/auth/register	69.668	0.676	2.484	5.297	228
1.1%	59	0	0	PUT	/api/me/profile-image	68.092	1.154	6.751	6.751	0
1.0%	74	0	0	PUT	^/api/follows/[^/]+$	60.660	0.820	6.400	6.400	0
0.9%	61	3	0	POST	^/api/posts/[^/]+/realmojis$	57.534	0.943	6.011	6.011	433
0.8%	24	0	0	GET	/api/me/export	51.929	2.164	10.406	10.406	2203077
0.7%	240	23	0	GET	^/api/posts/me/current$	43.826	0.183	1.108	2.133	592
0.7%	255	0	0	GET	/api/users	43.545	0.171	0.405	2.130	244
0.6%	41	0	0	PATCH	/api/me	35.717	0.871	6.574	6.574	277
0.5%	195	7	0	GET	^/api/posts/[^/]+$	31.665	0.162	0.346	0.386	592
0.5%	189	0	0	GET	^/api/relationships/[^/]+	30.403	0.161	1.126	1.287	848
0.4%	16	0	0	DELETE	^/api/comments/[^/]+$	22.828	1.427	5.680	5.680	0
0.4%	19	3	0	PATCH	^/api/posts/[^/]+$	22.771	1.198	6.719	6.719	556
0.2%	15	0	0	DELETE	^/api/posts/[^/]+$	15.833	1.056	3.373	3.373	0
0.2%	90	1	0	GET	^/api/users/[^/]+$	13.963	0.155	0.418	0.418	183
0.2%	83	0	0	GET	^/api/posts/[^/]+/comments$	11.051	0.133	0.246	0.246	1073
0.2%	51	0	0	GET	^/api/memories/[^/]+	10.231	0.201	0.498	0.498	1372
0.2%	62	0	0	GET	^/api/posts/[^/]+/realmojis$	9.999	0.161	0.302	0.302	1133
0.2%	64	2	0	GET	/api/me	9.934	0.155	1.008	1.008	296
0.1%	6	0	0	DELETE	^/api/realmojis/[^/]+$	7.805	1.301	4.395	4.395	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.5%	160	0	0	GET	/api/notifications	1218.578	7.616	9.480	9.480	181
20.1%	174	4	0	POST	/api/posts	731.276	4.203	14.298	14.322	642
9.7%	310	6	0	POST	/api/friend-requests	352.277	1.136	12.693	13.289	113
8.3%	213	0	0	POST	/api/auth/login	301.399	1.415	13.416	15.260	248
5.3%	104	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	191.631	1.843	13.222	13.959	115
3.8%	824	7	0	GET	^/api/images/[^/]+	138.763	0.168	0.980	1.286	172014
2.4%	543	0	0	GET	^/api/feeds/[^/]+	88.602	0.163	0.436	3.991	2701
2.0%	236	0	0	GET	/api/posting-window/current	73.358	0.311	2.429	4.163	240
1.9%	28	5	0	POST	^/api/posts/[^/]+/realmojis$	70.835	2.530	16.938	16.938	439
1.7%	60	0	0	POST	/api/auth/register	62.371	1.040	5.371	5.371	229
1.7%	62	7	0	POST	^/api/posts/[^/]+/comments$	61.774	0.996	16.034	16.034	301
1.2%	237	0	0	GET	/api/memories	45.339	0.191	0.512	1.142	12881
1.1%	26	0	0	PATCH	/api/me	39.692	1.527	14.441	14.441	275
1.0%	43	0	0	PUT	^/api/follows/[^/]+$	34.521	0.803	8.591	8.591	0
0.8%	13	0	0	GET	/api/me/export	27.979	2.152	4.469	4.469	2837717
0.8%	31	0	0	PUT	/api/me/profile-image	27.284	0.880	4.422	4.422	0
0.7%	134	22	0	GET	^/api/posts/me/current$	25.321	0.189	0.811	1.100	548
0.7%	34	0	0	POST	/api/auth/logout	25.000	0.735	4.237	4.237	0
0.7%	148	0	0	GET	/api/users	23.710	0.160	0.451	0.980	235
0.5%	15	1	0	PATCH	^/api/posts/[^/]+$	19.142	1.276	3.228	3.228	603
0.4%	7	0	0	DELETE	^/api/posts/[^/]+$	16.008	2.287	5.137	5.137	0
0.4%	96	0	0	GET	^/api/relationships/[^/]+	13.539	0.141	1.170	1.170	1488
0.3%	83	5	0	GET	^/api/posts/[^/]+$	11.409	0.137	0.264	0.264	580
0.2%	52	0	0	GET	^/api/users/[^/]+$	6.919	0.133	0.442	0.442	177
0.2%	29	0	0	GET	^/api/memories/[^/]+	5.999	0.207	0.814	0.814	953
0.1%	34	0	0	GET	/api/me	5.076	0.149	0.651	0.651	296
0.1%	3	0	0	DELETE	^/api/realmojis/[^/]+$	4.256	1.419	2.155	2.155	0
0.1%	28	0	0	GET	^/api/posts/[^/]+/realmojis$	3.987	0.142	0.285	0.285	2003
0.1%	22	0	0	GET	^/api/posts/[^/]+/comments$	3.380	0.154	0.290	0.290	1350
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	2.546	0.849	2.265	2.265	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 45447 クエリ / 80 種類 / 合計 291.330s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	59.7%	2729	173.858	63.71	533.90	646.5	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	25.7%	210	74.992	357.10	1353.45	7193.0	137.0	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.3%	1902	15.436	8.12	276.29	118.2	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.4%	5727	12.746	2.23	47.11	10.8	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.5%	1881	4.335	2.30	23.51	25.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.8%	1	2.193	2193.50	2193.50	29677.0	1199604.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.6%	448	1.631	3.64	36.99	30.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
8	0.3%	538	0.869	1.62	8.93	21.0	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
9	0.3%	3801	0.866	0.23	28.22	20.1	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	0.3%	1	0.752	752.29	752.29	11528.0	529860.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
11	0.2%	3180	0.453	0.14	22.95	11.2	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
12	0.1%	1354	0.425	0.31	27.15	21.3	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
13	0.1%	2	0.322	160.89	172.90	225679.0	2656.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.1%	2	0.260	129.86	252.13	16763.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	0.1%	1	0.217	216.96	216.96	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…

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
Build ID: 09de871e65949e92948bda21273b5824f63ed68d
Type: cpu
Time: 2026-10-31 12:08:32 JST
Duration: 60s, Total samples = 55.27s (92.11%)
Showing nodes accounting for 42.93s, 77.67% of 55.27s total
Dropped 1022 nodes (cum <= 0.28s)
Showing top 40 nodes out of 292
      flat  flat%   sum%        cum   cum%
     3.66s  6.62%  6.62%      3.66s  6.62%  internal/runtime/syscall/linux.Syscall6
     3.35s  6.06% 12.68%     11.30s 20.45%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     3.24s  5.86% 18.55%      5.21s  9.43%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     3.24s  5.86% 24.41%     11.22s 20.30%  runtime.scanObject
     2.88s  5.21% 29.62%      2.88s  5.21%  runtime.memclrNoHeapPointers
     2.78s  5.03% 34.65%      3.55s  6.42%  runtime.findObject
     2.57s  4.65% 39.30%      2.71s  4.90%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     2.21s  4.00% 43.30%      3.78s  6.84%  runtime.tryDeferToSpanScan
     2.04s  3.69% 46.99%      2.04s  3.69%  runtime.memmove
     1.85s  3.35% 50.33%      1.85s  3.35%  crypto/internal/fips140/sha256.blockAVX2
     1.76s  3.18% 53.52%      5.74s 10.39%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.47s  2.66% 56.18%      2.10s  3.80%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     1.42s  2.57% 58.75%      1.42s  2.57%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.25s  2.26% 61.01%      1.25s  2.26%  aeshashbody
     1.17s  2.12% 63.13%      1.17s  2.12%  runtime.spanClass.sizeclass (inline)
     1.09s  1.97% 65.10%      1.09s  1.97%  example.com/isureal-company/isureal.(*UserRow).active
     0.68s  1.23% 66.33%      0.68s  1.23%  memeqbody
     0.59s  1.07% 67.40%      0.61s  1.10%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.48s  0.87% 68.26%      1.96s  3.55%  runtime.mapaccess1_faststr
     0.46s  0.83% 69.10%      0.46s  0.83%  indexbody
     0.41s  0.74% 69.84%      0.61s  1.10%  example.com/isureal-company/isureal.postLess
     0.39s  0.71% 70.54%      0.84s  1.52%  runtime.greyobject
     0.35s  0.63% 71.18%      0.37s  0.67%  runtime.spanOf (inline)
     0.32s  0.58% 71.76%      0.32s  0.58%  runtime.(*mspan).divideByElemSize (inline)
     0.32s  0.58% 72.34%      0.32s  0.58%  runtime.nextFreeFast (inline)
     0.31s  0.56% 72.90%      0.31s  0.56%  runtime.typePointers.nextFast (inline)
     0.29s  0.52% 73.42%      0.29s  0.52%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.29s  0.52% 73.95%      2.10s  3.80%  runtime.mapassign_faststr
     0.26s  0.47% 74.42%      0.33s   0.6%  runtime.(*mspan).markBitsForIndex (inline)
     0.24s  0.43% 74.85%      0.28s  0.51%  runtime.typePointers.next
     0.22s   0.4% 75.25%      0.74s  1.34%  runtime.scanObjectsSmall
     0.20s  0.36% 75.61%      7.66s 13.86%  runtime.mallocgc
     0.17s  0.31% 75.92%      0.48s  0.87%  database/sql.convertAssignRows
     0.17s  0.31% 76.23%      2.03s  3.67%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.16s  0.29% 76.52%      0.35s  0.63%  runtime.scanObjectSmall
     0.15s  0.27% 76.79%      2.08s  3.76%  crypto/internal/fips140/sha256.(*Digest).Write
     0.13s  0.24% 77.02%      0.98s  1.77%  runtime.mallocgcSmallScanNoHeader
     0.13s  0.24% 77.26%      1.28s  2.32%  runtime.scanSpan
     0.12s  0.22% 77.47%      2.60s  4.70%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s   0.2% 77.67%      0.45s  0.81%  encoding/json.structEncoder.encode
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 52.0% / 100% | 39% | 0.1% | 8 | 110 MB | CPU が飽和 |
| isu2 (2 cpu) | 11.3% / 32% | 0% | 0.9% | 4 | 2965 MB |  |
| isu3 (2 cpu) | 7.8% / 32% | 0% | 0.8% | 5 | 3083 MB | run queue が積もっている |
| isu4 (2 cpu) | 41.6% / 100% | 17% | 14.2% | 32 | 159 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 85% / nginx 12% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 20% / nginx 1% |
| isu3 | isureal 13% |
| isu4 | postgres 44% / isureal 7% |

## Diff from previous run

スコア: 0 → 68346 (+68346)

コードの差分は無い（同じ sha 9f29066a25e1）。
スコアの違いはブレか、サーバ側の設定変更による。
