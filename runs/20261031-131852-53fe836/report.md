# run 20261031-131852-53fe836

## Summary

| | |
|---|---|
| score | 105712 |
| status | pass |
| ref | eu-us-post-verify |
| sha | 53fe836f3875 |
| 所要時間 | 929s |
| 前回比 | +41012 (63.4%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +41012 (63.4%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | groupB-eu-us-post-verify |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	133	369	837	838
europe	7	19	55	148	156	156
america	6	16	45	104	106	105
```

```
{"id":"13052","status":"done","portalStatus":"finished","passed":true,"score":105712,"raw":105712,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:32:56.206Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:32:59.173Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:33:02.173Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:33:05.195Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T04:33:08.172Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T04:33:11.173Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T04:33:14.208Z level=INFO msg="asia 3 日目: 133 人"
time=2026-10-31T04:33:17.207Z level=INFO msg="europe 3 日目: 55 人"
time=2026-10-31T04:33:18.408Z level=INFO msg="user moved" user=noleozoe6122 from=asia to=europe
time=2026-10-31T04:33:20.174Z level=INFO msg="america 3 日目: 45 人"
time=2026-10-31T04:33:22.157Z level=INFO msg="user moved" user=sebenyo6350 from=europe to=america
time=2026-10-31T04:33:23.197Z level=INFO msg="asia 4 日目: 369 人"
time=2026-10-31T04:33:26.366Z level=INFO msg="europe 4 日目: 148 人"
time=2026-10-31T04:33:27.708Z level=INFO msg="user moved" user=releodan.6392 from=america to=asia
time=2026-10-31T04:33:27.887Z level=INFO msg="user moved" user=orsol_7344 from=asia to=america
time=2026-10-31T04:33:29.287Z level=INFO msg="america 4 日目: 104 人"
time=2026-10-31T04:33:30.081Z level=INFO msg="user moved" user=releodan.6392 from=asia to=america
time=2026-10-31T04:33:31.807Z level=INFO msg="user moved" user=evator-8045 from=asia to=america
time=2026-10-31T04:33:32.363Z level=INFO msg="asia 5 日目: 837 人"
time=2026-10-31T04:33:35.583Z level=INFO msg="europe 5 日目: 156 人"
time=2026-10-31T04:33:38.310Z level=INFO msg="america 5 日目: 106 人"
time=2026-10-31T04:33:39.611Z level=INFO msg="user moved" user=sebenyo6350 from=america to=asia
time=2026-10-31T04:33:47.435Z level=INFO msg="asia 最終日（グランプリ）: 838 人"
time=2026-10-31T04:33:47.502Z level=INFO msg="america 最終日（グランプリ）: 105 人"
time=2026-10-31T04:33:47.548Z level=INFO msg="europe 最終日（グランプリ）: 156 人"
time=2026-10-31T04:33:56.218Z level=INFO msg="グランプリ: アジア 426 人, ヨーロッパ 73 人, アメリカ 52 人"
{
  "pass": true,
  "score": 105712,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.4%	1295	0	0	GET	/api/notifications	10830.120	8.363	15.110	15.124	164
16.2%	1867	39	0	POST	/api/posts	5425.132	2.906	9.830	14.399	634
12.8%	4415	64	0	POST	/api/friend-requests	4263.702	0.966	6.714	13.887	114
9.2%	12282	746	0	GET	^/api/images/[^/]+	3068.943	0.250	4.608	14.195	138432
7.1%	4844	1140	0	POST	/api/auth/login	2367.659	0.489	6.055	12.074	193
4.7%	1505	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1565.184	1.040	6.548	11.667	116
4.2%	856	1	0	POST	/api/auth/register	1418.047	1.657	8.292	10.963	227
2.5%	599	129	0	PUT	/api/me/profile-image	834.819	1.394	9.815	10.737	4
1.6%	562	0	0	PUT	^/api/follows/[^/]+$	538.429	0.958	6.464	14.440	0
1.6%	448	12	0	POST	^/api/posts/[^/]+/realmojis$	530.649	1.184	7.307	12.026	472
1.3%	382	0	0	PATCH	/api/me	426.545	1.117	6.532	14.812	280
1.1%	417	7	0	POST	^/api/posts/[^/]+/comments$	368.758	0.884	6.586	10.692	330
0.9%	3083	0	0	GET	/api/memories	292.715	0.095	0.552	0.965	11120
0.7%	595	0	0	POST	/api/auth/logout	247.226	0.416	6.219	12.283	0
0.7%	188	7	0	PATCH	^/api/posts/[^/]+$	231.823	1.233	7.211	8.019	628
0.7%	9980	0	0	GET	^/api/feeds/[^/]+	219.062	0.022	0.152	1.327	2033
0.5%	201	0	0	GET	/api/me/export	179.461	0.893	7.209	9.083	1867701
0.3%	60	0	0	DELETE	^/api/posts/[^/]+$	110.412	1.840	11.077	11.077	0
0.3%	2701	0	0	GET	/api/posting-window/current	100.657	0.037	0.257	2.455	236
0.2%	2174	0	0	GET	/api/users	52.965	0.024	0.149	2.359	230
0.2%	75	0	0	DELETE	^/api/realmojis/[^/]+$	50.265	0.670	12.218	12.218	0
0.1%	1572	67	0	GET	^/api/posts/me/current$	48.304	0.031	0.176	0.406	621
0.1%	561	0	0	GET	^/api/memories/[^/]+	29.685	0.053	0.364	0.621	634
0.1%	1504	13	0	GET	^/api/posts/[^/]+$	26.770	0.018	0.139	0.382	617
0.1%	71	0	0	DELETE	^/api/comments/[^/]+$	26.467	0.373	3.216	3.216	0
0.1%	1332	0	0	GET	^/api/relationships/[^/]+	24.410	0.018	0.164	0.281	886
0.1%	1	0	0	POST	/api/initialize	19.525	19.525	19.525	19.525	90
0.1%	13	0	0	PUT	^/api/admin/users/[^/]+/certification$	18.697	1.438	9.369	9.369	0
0.1%	26	0	0	PUT	^/api/blocks/[^/]+$	17.364	0.668	2.941	2.941	0
0.0%	692	0	0	GET	^/api/users/[^/]+$	13.006	0.019	0.138	1.183	199
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.5%	369	0	0	GET	/api/notifications	3083.027	8.355	11.936	12.162	177
16.3%	474	8	0	POST	/api/posts	1269.179	2.678	8.589	9.350	640
9.8%	880	12	0	POST	/api/friend-requests	761.292	0.865	6.496	8.944	114
5.7%	556	0	0	POST	/api/auth/login	444.589	0.800	6.579	11.871	246
5.6%	1926	19	0	GET	^/api/images/[^/]+	434.615	0.226	1.305	1.976	191260
4.1%	257	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	319.871	1.245	7.658	9.696	116
3.6%	1584	0	0	GET	^/api/feeds/[^/]+	283.113	0.179	0.624	6.742	1762
1.6%	592	0	0	GET	/api/memories	126.592	0.214	0.745	1.436	10626
1.5%	643	0	0	GET	/api/posting-window/current	118.141	0.184	0.534	0.911	239
1.2%	73	4	0	POST	^/api/posts/[^/]+/realmojis$	91.469	1.253	6.165	6.165	483
1.2%	91	0	0	POST	/api/auth/logout	89.824	0.987	10.102	10.102	0
1.0%	409	35	0	GET	^/api/posts/me/current$	81.115	0.198	0.678	0.759	596
1.0%	35	0	0	GET	/api/me/export	79.488	2.271	5.604	5.604	2138385
1.0%	154	0	0	POST	/api/auth/register	78.379	0.509	3.634	4.523	228
0.9%	453	0	0	GET	/api/users	73.053	0.161	0.534	0.735	234
0.8%	85	0	0	PUT	/api/me/profile-image	62.497	0.735	7.440	7.440	0
0.8%	68	1	0	POST	^/api/posts/[^/]+/comments$	59.343	0.873	4.747	4.747	357
0.7%	45	2	0	PATCH	^/api/posts/[^/]+$	53.287	1.184	5.175	5.175	616
0.6%	115	0	0	PUT	^/api/follows/[^/]+$	49.651	0.432	5.217	6.070	0
0.5%	251	19	0	GET	^/api/posts/[^/]+$	41.171	0.164	0.578	0.707	580
0.5%	245	0	0	GET	^/api/relationships/[^/]+	38.051	0.155	0.647	0.701	840
0.4%	76	0	0	PATCH	/api/me	34.897	0.459	4.754	4.754	274
0.4%	20	0	0	DELETE	^/api/posts/[^/]+$	32.641	1.632	5.547	5.547	0
0.2%	100	0	0	GET	^/api/posts/[^/]+/realmojis$	15.956	0.160	0.568	0.677	1140
0.2%	102	0	0	GET	^/api/users/[^/]+$	15.612	0.153	0.615	0.635	194
0.2%	89	1	0	GET	/api/me	13.943	0.157	0.498	0.498	298
0.2%	9	0	0	DELETE	^/api/realmojis/[^/]+$	13.649	1.517	8.483	8.483	0
0.2%	68	0	0	GET	^/api/memories/[^/]+	12.218	0.180	0.485	0.485	988
0.2%	75	0	0	GET	^/api/posts/[^/]+/comments$	12.210	0.163	0.613	0.613	1073
0.1%	12	0	0	DELETE	^/api/comments/[^/]+$	8.546	0.712	2.322	2.322	0
…（上位 30 / 全 31 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.7%	240	0	0	GET	/api/notifications	1806.391	7.527	9.238	9.269	178
14.9%	313	4	0	POST	/api/posts	754.353	2.410	7.886	9.607	637
12.4%	584	12	0	POST	/api/friend-requests	626.712	1.073	6.794	8.647	113
6.4%	1309	9	0	GET	^/api/images/[^/]+	324.142	0.248	1.562	2.576	170097
6.3%	363	0	0	POST	/api/auth/login	320.295	0.882	7.687	9.069	250
4.7%	198	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	235.603	1.190	9.756	9.911	116
3.4%	922	0	0	GET	^/api/feeds/[^/]+	171.336	0.186	0.641	0.792	1775
1.9%	401	0	0	GET	/api/memories	97.517	0.243	1.067	3.626	13614
1.7%	422	0	0	GET	/api/posting-window/current	84.474	0.200	0.649	1.178	240
1.3%	51	5	0	POST	^/api/posts/[^/]+/realmojis$	64.960	1.274	7.576	7.576	385
1.2%	24	0	0	GET	/api/me/export	58.585	2.441	7.990	7.990	2799776
1.1%	265	21	0	GET	^/api/posts/me/current$	57.261	0.216	0.649	0.964	595
1.1%	297	0	0	GET	/api/users	53.199	0.179	0.927	0.979	231
1.0%	46	0	0	POST	/api/auth/logout	50.363	1.095	7.499	7.499	0
1.0%	58	0	0	PUT	^/api/follows/[^/]+$	48.317	0.833	5.587	5.587	0
0.9%	29	2	0	PATCH	^/api/posts/[^/]+$	46.111	1.590	6.051	6.051	614
0.8%	51	0	0	PATCH	/api/me	42.133	0.826	6.880	6.880	286
0.8%	102	0	0	POST	/api/auth/register	40.854	0.401	1.352	5.262	229
0.7%	15	0	0	DELETE	^/api/posts/[^/]+$	36.681	2.445	7.892	7.892	0
0.7%	49	0	0	PUT	/api/me/profile-image	36.624	0.747	5.276	5.276	0
0.5%	147	0	0	GET	^/api/relationships/[^/]+	25.186	0.171	0.812	1.016	1166
0.4%	118	3	0	GET	^/api/posts/[^/]+$	18.695	0.158	0.461	0.521	613
0.2%	4	0	0	PUT	^/api/blocks/[^/]+$	9.574	2.393	8.692	8.692	0
0.2%	46	0	0	GET	^/api/memories/[^/]+	9.206	0.200	0.792	0.792	887
0.2%	53	0	0	GET	/api/me	8.457	0.160	0.517	0.517	299
0.1%	4	1	0	POST	^/api/posts/[^/]+/comments$	7.027	1.757	2.793	2.793	272
0.1%	43	0	0	GET	^/api/posts/[^/]+/realmojis$	6.657	0.155	0.361	0.361	1503
0.1%	38	0	0	GET	^/api/posts/[^/]+/comments$	5.804	0.153	0.401	0.401	1062
0.1%	1	0	0	DELETE	/api/me/profile-image	5.691	5.691	5.691	5.691	0
0.1%	30	0	0	GET	^/api/users/[^/]+$	4.380	0.146	0.231	0.231	203
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 64112 クエリ / 79 種類 / 合計 521.758s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	66.2%	3762	345.308	91.79	791.21	659.7	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	17.8%	260	93.081	358.00	1868.98	1246.7	126.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	6.3%	2743	32.823	11.97	641.20	120.1	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	3.8%	7990	19.917	2.49	44.14	10.7	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.6%	2605	8.109	3.11	128.43	24.9	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.5%	2	7.908	3954.02	7757.04	226807.5	2670.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	1	2.262	2261.94	2261.94	30671.0	1199825.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.4%	551	2.069	3.76	32.51	28.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.3%	1	1.619	1619.11	1619.11	23317.0	626911.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
10	0.3%	5782	1.571	0.27	118.78	19.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.2%	480	0.923	1.92	19.54	22.9	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.2%	1907	0.846	0.44	111.96	20.3	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
13	0.1%	1	0.782	782.14	782.14	12349.0	530127.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
14	0.1%	4623	0.622	0.13	65.41	10.9	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
15	0.1%	1956	0.460	0.24	116.00	7.8	1.0	UPDATE friend_requests SET status=$2, updated_at=$3 WHERE id=$1

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
Build ID: 9902038d7a7d1616fb044f696f608dc4cc6f9c2c
Type: cpu
Time: 2026-10-31 13:32:30 JST
Duration: 60.03s, Total samples = 35.18s (58.61%)
Showing nodes accounting for 25.33s, 72.00% of 35.18s total
Dropped 893 nodes (cum <= 0.18s)
Showing top 40 nodes out of 296
      flat  flat%   sum%        cum   cum%
     3.06s  8.70%  8.70%      3.06s  8.70%  internal/runtime/syscall/linux.Syscall6
     2.40s  6.82% 15.52%      2.40s  6.82%  crypto/internal/fips140/sha256.blockAVX2
     1.81s  5.14% 20.67%      1.81s  5.14%  runtime.memclrNoHeapPointers
     1.62s  4.60% 25.27%      1.62s  4.60%  runtime.memmove
     1.61s  4.58% 29.85%      2.08s  5.91%  runtime.findObject
     1.54s  4.38% 34.22%      5.79s 16.46%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.43s  4.06% 38.29%      1.43s  4.06%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.41s  4.01% 42.30%      5.98s 17.00%  runtime.scanObject
     1.27s  3.61% 45.91%      1.27s  3.61%  aeshashbody
     1.27s  3.61% 49.52%      2.13s  6.05%  runtime.tryDeferToSpanScan
     0.88s  2.50% 52.02%      0.88s  2.50%  example.com/isureal-company/isureal.(*UserRow).active
     0.59s  1.68% 53.70%      0.59s  1.68%  runtime.spanClass.sizeclass (inline)
     0.53s  1.51% 55.20%      0.53s  1.51%  memeqbody
     0.39s  1.11% 56.31%      0.87s  2.47%  database/sql.convertAssignRows
     0.36s  1.02% 57.33%      1.97s  5.60%  runtime.mapaccess1_faststr
     0.35s  0.99% 58.33%      0.35s  0.99%  indexbody
     0.34s  0.97% 59.30%      1.11s  3.16%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.33s  0.94% 60.23%      0.35s  0.99%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.29s  0.82% 61.06%      0.59s  1.68%  example.com/isureal-company/isureal.postLess
     0.28s   0.8% 61.85%      1.09s  3.10%  runtime.mallocgcSmallScanNoHeader
     0.27s  0.77% 62.62%      0.27s  0.77%  runtime.futex
     0.26s  0.74% 63.36%      0.42s  1.19%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.23s  0.65% 64.01%      0.50s  1.42%  runtime.greyobject
     0.21s   0.6% 64.61%      0.21s   0.6%  runtime.(*mspan).writeHeapBitsSmall
     0.20s  0.57% 65.18%      0.20s  0.57%  runtime.nextFreeFast (inline)
     0.19s  0.54% 65.72%      4.11s 11.68%  runtime.mallocgc
     0.18s  0.51% 66.23%      0.20s  0.57%  runtime.(*mspan).markBitsForIndex (inline)
     0.18s  0.51% 66.74%      1.79s  5.09%  runtime.mapassign_faststr
     0.18s  0.51% 67.25%      0.19s  0.54%  runtime.typePointers.next
     0.17s  0.48% 67.74%      1.07s  3.04%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.17s  0.48% 68.22%      0.24s  0.68%  runtime.spanOf (inline)
     0.16s  0.45% 68.68%      4.47s 12.71%  database/sql.withLock
     0.16s  0.45% 69.13%      0.20s  0.57%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.15s  0.43% 69.56%      2.64s  7.50%  crypto/internal/fips140/sha256.(*Digest).Write
     0.15s  0.43% 69.98%      0.18s  0.51%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.15s  0.43% 70.41%      0.20s  0.57%  runtime.heapArenaOf (inline)
     0.14s   0.4% 70.81%      0.28s   0.8%  compress/flate.(*compressor).deflate
     0.14s   0.4% 71.21%      1.65s  4.69%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.14s   0.4% 71.60%      0.25s  0.71%  internal/runtime/maps.(*Iter).Next
     0.14s   0.4% 72.00%      0.47s  1.34%  runtime.scanObjectsSmall
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 50.2% / 100% | 30% | 0.8% | 8 | 120 MB | CPU が飽和 |
| isu2 (2 cpu) | 19.6% / 100% | 1% | 0.5% | 4 | 1779 MB |  |
| isu3 (2 cpu) | 16.8% / 97% | 1% | 0.4% | 4 | 2355 MB |  |
| isu4 (2 cpu) | 51.2% / 100% | 23% | 19.1% | 32 | 121 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 67% / nginx 12% / journalctl 3% / systemd-journal 2% |
| isu2 | isureal 36% / nginx 1% |
| isu3 | isureal 31% / nginx 1% |
| isu4 | postgres 57% / isureal 2% / kcompactd0 2% |

## Diff from previous run

スコア: 64700 → 105712 (+41012)

`e0e4cb3..53fe836`

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
 webapp/go/post_api.go                              |  64 +++-
 webapp/go/post_forward.go                          | 113 +++++++
 webapp/go/post_forward_test.go                     |  94 ++++++
 webapp/go/realmoji_api.go                          |   2 +-
 webapp/go/relationship.go                          |  10 +-
 webapp/go/store.go                                 | 113 ++++++-
 22 files changed, 854 insertions(+), 66 deletions(-)
```
