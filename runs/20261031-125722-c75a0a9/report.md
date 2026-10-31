# run 20261031-125722-c75a0a9

## Summary

| | |
|---|---|
| score | 95693 |
| status | pass |
| ref | HEAD |
| sha | c75a0a91ecce |
| 所要時間 | 119s |
| 前回比 | +30993 (47.9%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +30993 (47.9%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | tls-off-export-pk |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	141	406	966	965
europe	7	16	42	104	105	105
america	6	12	27	51	51	51
```

```
{"id":"10436","status":"done","portalStatus":"finished","passed":true,"score":95693,"raw":95693,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:57:53.884Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:57:56.853Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:57:59.850Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:58:02.877Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T03:58:05.852Z level=INFO msg="europe 2 日目: 16 人"
time=2026-10-31T03:58:08.851Z level=INFO msg="america 2 日目: 12 人"
time=2026-10-31T03:58:11.874Z level=INFO msg="asia 3 日目: 141 人"
time=2026-10-31T03:58:12.027Z level=INFO msg="user moved" user=futa-5444 from=america to=asia
time=2026-10-31T03:58:13.066Z level=INFO msg="user moved" user=futa-5444 from=asia to=america
time=2026-10-31T03:58:14.854Z level=INFO msg="europe 3 日目: 42 人"
time=2026-10-31T03:58:17.882Z level=INFO msg="america 3 日目: 27 人"
time=2026-10-31T03:58:20.891Z level=INFO msg="asia 4 日目: 406 人"
time=2026-10-31T03:58:23.884Z level=INFO msg="europe 4 日目: 104 人"
time=2026-10-31T03:58:26.954Z level=INFO msg="america 4 日目: 51 人"
time=2026-10-31T03:58:28.645Z level=INFO msg="user moved" user=soldan.7033 from=asia to=europe
time=2026-10-31T03:58:30.065Z level=INFO msg="asia 5 日目: 966 人"
time=2026-10-31T03:58:32.949Z level=INFO msg="europe 5 日目: 105 人"
time=2026-10-31T03:58:35.909Z level=INFO msg="america 5 日目: 51 人"
time=2026-10-31T03:58:44.923Z level=INFO msg="america 最終日（グランプリ）: 51 人"
time=2026-10-31T03:58:45.038Z level=INFO msg="asia 最終日（グランプリ）: 965 人"
time=2026-10-31T03:58:45.177Z level=INFO msg="europe 最終日（グランプリ）: 105 人"
time=2026-10-31T03:58:53.871Z level=INFO msg="グランプリ: アジア 199 人, ヨーロッパ 3 人, アメリカ 5 人"
{
  "pass": true,
  "score": 95693,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
30.1%	1497	1	0	GET	/api/notifications	12246.120	8.180	14.999	15.030	161
20.2%	2137	69	0	POST	/api/posts	8210.617	3.842	12.688	16.130	630
12.6%	4915	49	0	POST	/api/friend-requests	5108.187	1.039	10.529	18.848	114
9.7%	13161	895	0	GET	^/api/images/[^/]+	3928.048	0.298	6.579	15.493	136323
7.3%	6025	1956	0	POST	/api/auth/login	2958.377	0.491	6.796	14.255	174
4.1%	1644	8	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1672.318	1.017	10.330	13.912	116
2.9%	983	1	0	POST	/api/auth/register	1159.661	1.180	7.388	10.424	227
2.5%	681	115	0	PUT	/api/me/profile-image	1031.123	1.514	7.491	9.571	4
1.7%	674	0	0	PUT	^/api/follows/[^/]+$	681.753	1.012	6.387	9.297	0
1.5%	506	30	0	POST	^/api/posts/[^/]+/realmojis$	597.119	1.180	7.966	10.686	458
1.3%	455	0	0	PATCH	/api/me	532.958	1.171	7.453	11.905	278
1.1%	607	22	0	POST	^/api/posts/[^/]+/comments$	445.023	0.733	6.335	12.111	317
0.8%	3192	0	0	GET	/api/memories	330.817	0.104	0.564	1.402	11291
0.8%	656	0	0	POST	/api/auth/logout	308.957	0.471	6.438	14.179	0
0.7%	10561	0	0	GET	^/api/feeds/[^/]+	270.899	0.026	0.325	1.429	2796
0.6%	194	10	0	PATCH	^/api/posts/[^/]+$	261.398	1.347	9.528	10.256	612
0.5%	251	0	0	GET	/api/me/export	202.850	0.808	8.986	12.211	1883260
0.4%	73	0	0	DELETE	^/api/posts/[^/]+$	167.483	2.294	12.834	12.834	0
0.2%	3101	0	0	GET	/api/posting-window/current	100.340	0.032	0.247	1.172	237
0.2%	1526	78	0	GET	^/api/posts/me/current$	72.331	0.047	0.497	1.417	617
0.1%	2361	0	0	GET	/api/users	57.463	0.024	0.197	1.429	240
0.1%	77	0	0	DELETE	^/api/realmojis/[^/]+$	52.187	0.678	5.606	5.606	0
0.1%	1484	0	0	GET	^/api/relationships/[^/]+	38.014	0.026	0.393	0.848	890
0.1%	93	0	0	DELETE	^/api/comments/[^/]+$	33.665	0.362	2.784	2.784	0
0.1%	580	0	0	GET	^/api/memories/[^/]+	31.320	0.054	0.437	0.992	667
0.1%	38	0	0	PUT	^/api/blocks/[^/]+$	28.602	0.753	7.453	7.453	0
0.1%	1561	16	0	GET	^/api/posts/[^/]+$	27.992	0.018	0.141	1.539	622
0.0%	9	0	0	DELETE	^/api/friendships/[^/]+$	14.994	1.666	8.712	8.712	0
0.0%	713	0	0	GET	^/api/posts/[^/]+/comments$	14.061	0.020	0.281	1.171	1446
0.0%	726	0	0	GET	^/api/posts/[^/]+/realmojis$	14.056	0.019	0.266	0.832	1064
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.9%	272	0	0	GET	/api/notifications	2315.481	8.513	12.063	12.350	179
22.7%	334	6	0	POST	/api/posts	1387.536	4.154	12.765	13.111	632
7.0%	601	10	0	POST	/api/friend-requests	430.851	0.717	4.873	17.644	114
5.6%	408	0	0	POST	/api/auth/login	344.943	0.845	9.336	12.724	249
5.1%	1539	6	0	GET	^/api/images/[^/]+	309.367	0.201	1.186	1.994	183029
4.1%	221	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	252.687	1.143	14.759	21.111	116
3.5%	1288	0	0	GET	^/api/feeds/[^/]+	213.027	0.165	0.498	0.938	2934
2.4%	141	13	0	POST	^/api/posts/[^/]+/realmojis$	146.857	1.042	7.572	12.039	428
1.6%	455	0	0	GET	/api/posting-window/current	99.726	0.219	0.520	0.828	239
1.3%	388	0	0	GET	/api/memories	81.513	0.210	0.803	0.934	14479
1.0%	58	8	0	POST	^/api/posts/[^/]+/comments$	59.432	1.025	12.146	12.146	262
0.9%	27	0	0	GET	/api/me/export	57.646	2.135	9.884	9.884	2086469
0.9%	61	0	0	POST	/api/auth/logout	55.406	0.908	9.724	9.724	0
0.7%	280	0	0	GET	/api/users	45.344	0.162	0.466	0.496	233
0.7%	244	13	0	GET	^/api/posts/me/current$	42.701	0.175	0.484	0.565	607
0.6%	102	0	0	POST	/api/auth/register	37.055	0.363	1.268	1.881	228
0.6%	56	0	0	PUT	/api/me/profile-image	34.116	0.609	2.060	2.060	0
0.5%	207	13	0	GET	^/api/posts/[^/]+$	31.668	0.153	0.423	0.492	575
0.4%	172	0	0	GET	^/api/relationships/[^/]+	26.748	0.156	0.524	0.546	1248
0.3%	54	0	0	PATCH	/api/me	21.242	0.393	2.324	2.324	266
0.3%	69	0	0	PUT	^/api/follows/[^/]+$	19.258	0.279	2.213	2.213	0
0.3%	18	2	0	PATCH	^/api/posts/[^/]+$	16.580	0.921	3.677	3.677	563
0.3%	105	0	0	GET	^/api/users/[^/]+$	15.475	0.147	0.446	0.506	182
0.2%	93	0	0	GET	^/api/posts/[^/]+/realmojis$	14.085	0.151	0.517	0.517	1322
0.2%	11	0	0	DELETE	^/api/realmojis/[^/]+$	11.919	1.084	6.185	6.185	0
0.2%	55	0	0	GET	^/api/memories/[^/]+	10.871	0.198	0.947	0.947	786
0.2%	62	1	0	GET	/api/me	9.412	0.152	0.432	0.432	293
0.1%	58	0	0	GET	^/api/posts/[^/]+/comments$	8.703	0.150	0.419	0.419	1258
0.1%	7	0	0	DELETE	^/api/posts/[^/]+$	6.918	0.988	2.064	2.064	0
0.1%	7	0	0	DELETE	^/api/comments/[^/]+$	3.313	0.473	2.311	2.311	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.7%	140	0	0	GET	/api/notifications	1050.514	7.504	9.091	9.091	177
18.1%	154	4	0	POST	/api/posts	518.840	3.369	12.356	12.370	654
10.6%	305	5	0	POST	/api/friend-requests	303.796	0.996	10.548	12.801	114
7.9%	142	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	226.800	1.597	14.111	14.468	116
6.1%	225	0	0	POST	/api/auth/login	174.740	0.777	9.713	10.688	253
5.5%	790	4	0	GET	^/api/images/[^/]+	157.297	0.199	0.840	2.230	145299
2.4%	404	0	0	GET	^/api/feeds/[^/]+	67.918	0.168	0.465	0.894	2932
1.5%	210	0	0	GET	/api/memories	43.449	0.207	0.636	0.796	16578
1.4%	204	0	0	GET	/api/posting-window/current	39.073	0.192	0.460	0.819	240
1.1%	13	0	0	GET	/api/me/export	30.196	2.323	5.575	5.575	2600643
1.0%	32	0	0	POST	/api/auth/logout	29.039	0.907	10.120	10.120	0
0.9%	144	0	0	GET	/api/users	25.510	0.177	0.474	1.107	231
0.7%	38	0	0	PUT	^/api/follows/[^/]+$	21.029	0.553	2.151	2.151	0
0.7%	120	16	0	GET	^/api/posts/me/current$	20.170	0.168	0.428	0.428	585
0.7%	30	0	0	PUT	/api/me/profile-image	19.274	0.642	5.312	5.312	0
0.7%	22	0	0	PATCH	/api/me	19.245	0.875	6.891	6.891	289
0.6%	13	1	0	POST	^/api/posts/[^/]+/realmojis$	17.592	1.353	7.032	7.032	484
0.6%	50	0	0	POST	/api/auth/register	16.346	0.327	1.758	1.758	229
0.6%	12	2	0	PATCH	^/api/posts/[^/]+$	16.120	1.343	6.136	6.136	568
0.6%	8	0	0	POST	^/api/posts/[^/]+/comments$	15.772	1.972	5.667	5.667	390
0.5%	81	0	0	GET	^/api/relationships/[^/]+	13.071	0.161	0.703	0.703	1352
0.3%	54	4	0	GET	^/api/posts/[^/]+$	8.992	0.167	0.782	0.782	575
0.2%	29	0	0	GET	^/api/memories/[^/]+	5.427	0.187	0.478	0.478	1129
0.2%	30	0	0	GET	^/api/users/[^/]+$	4.897	0.163	0.715	0.715	205
0.2%	33	0	0	GET	/api/me	4.374	0.133	0.353	0.353	295
0.1%	22	0	0	GET	^/api/posts/[^/]+/realmojis$	3.623	0.165	0.370	0.370	916
0.1%	4	0	0	DELETE	^/api/posts/[^/]+$	3.274	0.819	1.496	1.496	0
0.1%	13	0	0	GET	^/api/posts/[^/]+/comments$	1.560	0.120	0.151	0.151	1635
0.0%	2	0	0	PUT	^/api/blocks/[^/]+$	0.947	0.473	0.840	0.840	0
0.0%	1	0	0	DELETE	^/api/comments/[^/]+$	0.398	0.398	0.398	0.398	0
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 65022 クエリ / 82 種類 / 合計 418.586s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	63.3%	3818	264.843	69.37	619.49	627.7	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	21.4%	291	89.510	307.59	1997.69	1221.6	124.2	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	5.5%	1974	23.186	11.75	623.42	116.9	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.7%	8544	19.848	2.32	52.84	10.5	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	2.0%	2550	8.168	3.20	62.57	25.3	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.5%	616	2.165	3.51	33.26	28.0	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
7	0.5%	1	2.033	2033.28	2033.28	30299.0	1199740.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.5%	1	1.941	1941.12	1941.12	22272.0	626875.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.3%	5743	1.139	0.20	38.47	19.9	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	0.2%	643	1.017	1.58	11.52	22.1	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.2%	1	0.689	689.36	689.36	11957.0	529962.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
12	0.1%	4702	0.542	0.12	7.24	11.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
13	0.1%	1943	0.465	0.24	14.31	21.0	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
14	0.1%	2	0.345	172.30	185.63	230647.5	2714.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
15	0.1%	2	0.264	131.92	255.56	16793.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: 5ef7e0ba89049eaa570e417d38efe43a95adeaab
Type: cpu
Time: 2026-10-31 12:57:37 JST
Duration: 60.15s, Total samples = 49.35s (82.04%)
Showing nodes accounting for 35.82s, 72.58% of 49.35s total
Dropped 1058 nodes (cum <= 0.25s)
Showing top 40 nodes out of 277
      flat  flat%   sum%        cum   cum%
     4.60s  9.32%  9.32%      4.60s  9.32%  internal/runtime/syscall/linux.Syscall6
     3.12s  6.32% 15.64%     10.34s 20.95%  runtime.scanObject
     2.78s  5.63% 21.28%      3.49s  7.07%  runtime.findObject
     2.42s  4.90% 26.18%      2.42s  4.90%  runtime.memclrNoHeapPointers
     2.16s  4.38% 30.56%      2.16s  4.38%  crypto/internal/fips140/sha256.blockAVX2
     2.12s  4.30% 34.85%      2.12s  4.30%  runtime.memmove
     2.03s  4.11% 38.97%      3.30s  6.69%  runtime.tryDeferToSpanScan
     1.91s  3.87% 42.84%      5.80s 11.75%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.58s  3.20% 46.04%      1.58s  3.20%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.49s  3.02% 49.06%      1.49s  3.02%  aeshashbody
     1.18s  2.39% 51.45%      1.18s  2.39%  example.com/isureal-company/isureal.(*UserRow).active
     1.03s  2.09% 53.54%      3.35s  6.79%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.97s  1.97% 55.50%      1.56s  3.16%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.89s  1.80% 57.30%      0.89s  1.80%  runtime.spanClass.sizeclass (inline)
     0.73s  1.48% 58.78%      0.75s  1.52%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.67s  1.36% 60.14%      0.67s  1.36%  memeqbody
     0.51s  1.03% 61.18%      2.26s  4.58%  runtime.mapaccess1_faststr
     0.46s  0.93% 62.11%      0.46s  0.93%  indexbody
     0.44s  0.89% 63.00%      1.02s  2.07%  runtime.greyobject
     0.42s  0.85% 63.85%      0.59s  1.20%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.38s  0.77% 64.62%      0.52s  1.05%  example.com/isureal-company/isureal.postLess
     0.37s  0.75% 65.37%      0.37s  0.75%  runtime.nextFreeFast (inline)
     0.36s  0.73% 66.10%      0.42s  0.85%  runtime.(*mspan).markBitsForIndex (inline)
     0.30s  0.61% 66.71%      0.35s  0.71%  runtime.spanOf (inline)
     0.25s  0.51% 67.21%      0.33s  0.67%  runtime.heapArenaOf (inline)
     0.25s  0.51% 67.72%      0.28s  0.57%  runtime.typePointers.next
     0.24s  0.49% 68.21%      0.28s  0.57%  encoding/json.appendString[go.shape.string]
     0.22s  0.45% 68.65%      0.70s  1.42%  runtime.scanObjectsSmall
     0.20s  0.41% 69.06%      0.58s  1.18%  database/sql.convertAssignRows
     0.20s  0.41% 69.46%      0.82s  1.66%  encoding/json.structEncoder.encode
     0.19s  0.39% 69.85%      0.39s  0.79%  compress/flate.(*compressor).deflate
     0.19s  0.39% 70.23%      1.15s  2.33%  runtime.mallocgcSmallScanNoHeader
     0.18s  0.36% 70.60%      7.05s 14.29%  runtime.mallocgc
     0.15s   0.3% 70.90%      2.18s  4.42%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.15s   0.3% 71.21%      1.99s  4.03%  runtime.mapassign_faststr
     0.15s   0.3% 71.51%      0.31s  0.63%  runtime.scanObjectSmall
     0.14s  0.28% 71.79%      2.44s  4.94%  crypto/internal/fips140/sha256.(*Digest).Write
     0.14s  0.28% 72.08%      2.51s  5.09%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.13s  0.26% 72.34%      2.87s  5.82%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.12s  0.24% 72.58%      2.78s  5.63%  crypto/internal/fips140/hmac.(*HMAC).Sum
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 51.5% / 100% | 34% | 0.1% | 8 | 184 MB | CPU が飽和 |
| isu2 (2 cpu) | 19.5% / 58% | 0% | 0.4% | 3 | 2516 MB |  |
| isu3 (2 cpu) | 14.8% / 61% | 0% | 0.4% | 5 | 2610 MB | run queue が積もっている |
| isu4 (2 cpu) | 52.0% / 100% | 28% | 18.9% | 31 | 103 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 78% / nginx 16% / journalctl 3% / systemd-journal 2% |
| isu2 | isureal 37% / nginx 1% |
| isu3 | isureal 28% |
| isu4 | postgres 59% / isureal 3% |

## Diff from previous run

スコア: 64700 → 95693 (+30993)

`e0e4cb3..c75a0a9`

```
 webapp/go/account_deletion_api.go |   2 +-
 webapp/go/admin_user_api.go       |   2 +-
 webapp/go/comment_api.go          |   2 +-
 webapp/go/config.go               |   2 +
 webapp/go/export_api.go           |   4 +-
 webapp/go/feed_api.go             | 113 +++++++++----
 webapp/go/feed_test.go            | 331 ++++++++++++++++++++++++++++++++++++++
 webapp/go/offload.go              |  27 +++-
 webapp/go/post_api.go             |  62 +++++--
 webapp/go/realmoji_api.go         |   2 +-
 webapp/go/relationship.go         |  10 +-
 webapp/go/store.go                | 113 +++++++++++--
 12 files changed, 606 insertions(+), 64 deletions(-)
```
