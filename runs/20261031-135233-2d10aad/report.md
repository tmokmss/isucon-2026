# run 20261031-135233-2d10aad

## Summary

| | |
|---|---|
| score | 119317 |
| status | pass |
| ref | jpegcheck-fast |
| sha | 2d10aadebfbf |
| 所要時間 | 261s |
| 前回比 | +15196 (14.6%) vs 20261031-133618-8b07f57 |
| ベスト比 | +15196 (14.6%) vs 20261031-133618-8b07f57 = 104121 **ベスト更新** |
| note | JPEG の検証を 1.5 倍速く（main c6abd65 の上） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	134	383	923	923
europe	7	16	43	118	171	171
america	6	15	44	104	106	106
```

```
{"id":"14844","status":"done","portalStatus":"finished","passed":true,"score":119317,"raw":119317,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:55:29.660Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:55:32.628Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:55:35.626Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:55:38.649Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T04:55:41.627Z level=INFO msg="europe 2 日目: 16 人"
time=2026-10-31T04:55:44.628Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T04:55:47.649Z level=INFO msg="asia 3 日目: 134 人"
time=2026-10-31T04:55:50.220Z level=INFO msg="user moved" user=shibenma-5804 from=asia to=america
time=2026-10-31T04:55:50.628Z level=INFO msg="europe 3 日目: 43 人"
time=2026-10-31T04:55:53.627Z level=INFO msg="america 3 日目: 44 人"
time=2026-10-31T04:55:56.669Z level=INFO msg="asia 4 日目: 383 人"
time=2026-10-31T04:55:59.581Z level=INFO msg="user moved" user=ruulma_6141 from=europe to=america
time=2026-10-31T04:55:59.635Z level=INFO msg="europe 4 日目: 118 人"
time=2026-10-31T04:56:02.492Z level=INFO msg="user moved" user=suchizoe7432 from=asia to=europe
time=2026-10-31T04:56:02.707Z level=INFO msg="america 4 日目: 104 人"
time=2026-10-31T04:56:04.470Z level=INFO msg="user moved" user=danro_7035 from=asia to=europe
time=2026-10-31T04:56:05.863Z level=INFO msg="asia 5 日目: 923 人"
time=2026-10-31T04:56:08.652Z level=INFO msg="europe 5 日目: 171 人"
time=2026-10-31T04:56:11.677Z level=INFO msg="america 5 日目: 106 人"
time=2026-10-31T04:56:20.819Z level=INFO msg="asia 最終日（グランプリ）: 923 人"
time=2026-10-31T04:56:21.027Z level=INFO msg="america 最終日（グランプリ）: 106 人"
time=2026-10-31T04:56:21.144Z level=INFO msg="europe 最終日（グランプリ）: 171 人"
time=2026-10-31T04:56:29.655Z level=INFO msg="グランプリ: アジア 492 人, ヨーロッパ 43 人, アメリカ 39 人"
{
  "pass": true,
  "score": 119317,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.5%	1511	0	0	GET	/api/notifications	12286.468	8.131	15.133	15.141	166
16.8%	2249	41	0	POST	/api/posts	5996.479	2.666	9.059	11.722	639
12.5%	5177	80	0	POST	/api/friend-requests	4438.414	0.857	5.725	13.690	114
9.1%	13616	390	0	GET	^/api/images/[^/]+	3227.009	0.237	3.655	11.282	157953
5.7%	5456	1373	0	POST	/api/auth/login	2046.674	0.375	4.389	10.654	189
4.9%	1946	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1759.784	0.904	6.017	10.884	116
2.3%	669	22	0	POST	^/api/posts/[^/]+/realmojis$	830.501	1.241	5.862	7.597	461
2.3%	625	121	0	PUT	/api/me/profile-image	826.192	1.322	7.411	10.284	4
2.2%	942	1	0	POST	/api/auth/register	797.667	0.847	3.885	5.895	227
1.7%	667	13	0	POST	^/api/posts/[^/]+/comments$	623.271	0.934	5.706	12.630	329
1.3%	597	0	0	PUT	^/api/follows/[^/]+$	473.912	0.794	4.727	8.517	0
1.2%	3392	0	0	GET	/api/memories	412.453	0.122	0.919	2.446	10987
1.1%	445	0	0	PATCH	/api/me	389.361	0.875	5.111	7.537	274
0.8%	232	9	0	PATCH	^/api/posts/[^/]+$	278.229	1.199	7.593	7.841	624
0.7%	264	0	0	GET	/api/me/export	246.498	0.934	6.449	7.334	1827135
0.6%	680	0	0	POST	/api/auth/logout	211.115	0.310	3.944	6.054	0
0.4%	83	0	0	DELETE	^/api/posts/[^/]+$	157.808	1.901	6.737	6.737	0
0.4%	13475	0	0	GET	^/api/feeds/[^/]+	153.677	0.011	0.131	3.482	1951
0.2%	112	0	0	DELETE	^/api/comments/[^/]+$	86.718	0.774	4.364	6.667	0
0.2%	3241	0	0	GET	/api/posting-window/current	71.840	0.022	0.278	0.640	236
0.2%	101	0	0	DELETE	^/api/realmojis/[^/]+$	70.097	0.694	4.945	7.092	0
0.1%	596	0	0	GET	^/api/memories/[^/]+	38.917	0.065	0.763	0.904	678
0.1%	1887	93	0	GET	^/api/posts/me/current$	35.180	0.019	0.261	1.930	619
0.1%	2501	0	0	GET	/api/users	24.858	0.010	0.106	0.507	245
0.1%	33	0	0	PUT	^/api/blocks/[^/]+$	20.641	0.625	4.129	4.129	0
0.1%	1769	0	0	GET	^/api/relationships/[^/]+	19.308	0.011	0.220	0.630	873
0.1%	1	0	0	POST	/api/initialize	18.308	18.308	18.308	18.308	90
0.0%	1871	10	0	GET	^/api/posts/[^/]+$	14.018	0.007	0.068	0.633	619
0.0%	14	0	0	PUT	^/api/admin/users/[^/]+/certification$	13.651	0.975	9.466	9.466	0
0.0%	9	0	0	DELETE	/api/me/profile-image	8.629	0.959	2.124	2.124	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.2%	352	0	0	GET	/api/notifications	2819.085	8.009	12.018	12.296	168
15.1%	437	9	0	POST	/api/posts	1145.822	2.622	8.629	9.696	635
10.5%	976	16	0	POST	/api/friend-requests	793.445	0.813	5.549	9.177	114
6.1%	1603	13	0	GET	^/api/images/[^/]+	463.609	0.289	1.776	3.807	181163
5.5%	506	0	0	POST	/api/auth/login	414.059	0.818	5.245	6.741	245
3.8%	1284	0	0	GET	^/api/feeds/[^/]+	292.051	0.227	0.774	1.170	2589
3.4%	277	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	260.506	0.940	7.447	8.962	115
2.5%	612	0	0	GET	/api/posting-window/current	186.712	0.305	0.780	1.453	238
1.9%	475	0	0	GET	/api/memories	146.805	0.309	1.400	2.594	10305
1.7%	94	7	0	POST	^/api/posts/[^/]+/realmojis$	126.421	1.345	9.189	9.189	442
1.5%	168	0	0	POST	/api/auth/register	116.556	0.694	3.688	4.319	228
1.3%	457	0	0	GET	/api/users	95.917	0.210	0.754	1.013	243
1.2%	91	0	0	PUT	/api/me/profile-image	91.735	1.008	4.573	4.573	0
1.2%	331	30	0	GET	^/api/posts/me/current$	91.029	0.275	0.815	1.202	589
1.2%	37	0	0	GET	/api/me/export	89.323	2.414	8.042	8.042	2586578
1.1%	99	0	0	PUT	^/api/follows/[^/]+$	81.703	0.825	5.481	5.481	0
0.6%	74	0	0	PATCH	/api/me	42.979	0.581	3.172	3.172	285
0.6%	67	0	0	POST	/api/auth/logout	42.721	0.638	3.150	3.150	0
0.5%	49	1	0	POST	^/api/posts/[^/]+/comments$	41.163	0.840	4.706	4.706	296
0.5%	226	0	0	GET	^/api/relationships/[^/]+	40.229	0.178	0.669	0.690	879
0.5%	36	7	0	PATCH	^/api/posts/[^/]+$	35.728	0.992	5.762	5.762	521
0.4%	186	11	0	GET	^/api/posts/[^/]+$	33.698	0.181	0.662	0.708	583
0.3%	11	0	0	DELETE	^/api/realmojis/[^/]+$	21.025	1.911	14.592	14.592	0
0.3%	14	0	0	DELETE	^/api/posts/[^/]+$	19.398	1.386	4.119	4.119	0
0.2%	95	1	0	GET	/api/me	18.257	0.192	0.640	0.640	299
0.2%	86	0	0	GET	^/api/users/[^/]+$	17.217	0.200	1.200	1.200	188
0.2%	58	0	0	GET	^/api/memories/[^/]+	17.217	0.297	1.439	1.439	1403
0.2%	76	0	0	GET	^/api/posts/[^/]+/realmojis$	14.990	0.197	1.171	1.171	1988
0.1%	13	0	0	DELETE	^/api/comments/[^/]+$	11.214	0.863	3.079	3.079	0
0.1%	57	1	0	GET	^/api/posts/[^/]+/comments$	10.573	0.185	0.692	0.692	1805
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.9%	256	0	0	GET	/api/notifications	1921.680	7.507	9.358	9.360	180
14.4%	316	3	0	POST	/api/posts	728.381	2.305	8.485	13.396	662
9.4%	552	5	0	POST	/api/friend-requests	477.497	0.865	5.586	7.655	114
7.8%	1299	7	0	GET	^/api/images/[^/]+	397.387	0.306	1.906	2.891	168288
5.2%	361	0	0	POST	/api/auth/login	265.501	0.735	6.324	12.229	250
4.4%	197	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	221.085	1.122	6.201	6.471	116
4.0%	861	0	0	GET	^/api/feeds/[^/]+	203.970	0.237	0.709	1.175	2233
2.5%	384	0	0	GET	/api/memories	125.111	0.326	1.453	2.457	13555
2.1%	426	0	0	GET	/api/posting-window/current	107.376	0.252	0.729	1.190	239
1.4%	260	22	0	GET	^/api/posts/me/current$	72.294	0.278	0.749	0.812	611
1.3%	74	0	0	PUT	^/api/follows/[^/]+$	67.266	0.909	5.054	5.054	0
1.1%	70	0	0	PUT	/api/me/profile-image	54.996	0.786	4.469	4.469	0
1.1%	266	0	0	GET	/api/users	53.919	0.203	0.765	1.214	241
0.9%	18	0	0	GET	/api/me/export	46.811	2.601	9.824	9.824	2695214
0.9%	58	0	0	POST	/api/auth/logout	45.006	0.776	5.695	5.695	0
0.7%	44	2	0	POST	^/api/posts/[^/]+/comments$	37.587	0.854	7.478	7.478	370
0.6%	17	1	0	PATCH	^/api/posts/[^/]+$	31.393	1.847	6.312	6.312	632
0.6%	103	0	0	POST	/api/auth/register	31.311	0.304	1.744	3.066	229
0.6%	147	0	0	GET	^/api/relationships/[^/]+	30.226	0.206	0.672	0.703	1096
0.5%	22	3	0	POST	^/api/posts/[^/]+/realmojis$	27.063	1.230	4.114	4.114	384
0.5%	127	7	0	GET	^/api/posts/[^/]+$	25.468	0.201	0.689	0.754	581
0.4%	49	0	0	PATCH	/api/me	22.341	0.456	2.406	2.406	293
0.3%	48	0	0	GET	^/api/memories/[^/]+	14.248	0.297	1.297	1.297	1127
0.2%	62	0	0	GET	^/api/users/[^/]+$	12.586	0.203	0.694	0.694	190
0.2%	73	0	0	GET	/api/me	11.616	0.159	0.618	0.618	303
0.2%	8	0	0	DELETE	^/api/posts/[^/]+$	11.516	1.440	3.650	3.650	0
0.2%	41	0	0	GET	^/api/posts/[^/]+/comments$	8.203	0.200	0.645	0.645	2033
0.1%	14	0	0	DELETE	^/api/comments/[^/]+$	6.188	0.442	2.158	2.158	0
0.1%	28	0	0	GET	^/api/posts/[^/]+/realmojis$	5.046	0.180	0.616	0.616	1140
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	4.316	1.439	4.093	4.093	0
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 71825 クエリ / 80 種類 / 合計 871.878s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	74.4%	4371	648.322	148.32	996.36	722.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	14.0%	319	121.782	381.76	1831.32	1238.0	125.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	4.8%	2567	42.083	16.39	936.53	119.1	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	2.9%	8885	25.225	2.84	76.60	10.8	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.4%	2952	11.911	4.03	136.27	24.8	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.6%	2	5.057	2528.37	4845.95	237729.5	2798.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	753	3.269	4.34	22.29	29.0	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
8	0.3%	1	2.200	2199.82	2199.82	30991.0	1200069.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
9	0.2%	6590	1.823	0.28	133.97	20.1	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	0.2%	1	1.794	1794.26	1794.26	24402.0	627921.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
11	0.2%	744	1.607	2.16	18.68	20.3	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.1%	2352	0.993	0.42	115.22	21.1	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
13	0.1%	1	0.914	913.51	913.51	30023.0	2810.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.730	730.49	730.49	12660.0	530029.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
15	0.1%	4950	0.706	0.14	29.69	11.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

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
Build ID: 13052554f8bcdca52ba475f9529f43fffc026445
Type: cpu
Time: 2026-10-31 13:55:06 JST
Duration: 60.16s, Total samples = 37.24s (61.90%)
Showing nodes accounting for 26.63s, 71.51% of 37.24s total
Dropped 1028 nodes (cum <= 0.19s)
Showing top 40 nodes out of 281
      flat  flat%   sum%        cum   cum%
     4.36s 11.71% 11.71%      4.36s 11.71%  crypto/internal/fips140/sha256.blockAVX2
     3.90s 10.47% 22.18%      3.90s 10.47%  internal/runtime/syscall/linux.Syscall6
     1.83s  4.91% 27.09%      5.98s 16.06%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.48s  3.97% 31.07%      1.48s  3.97%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.45s  3.89% 34.96%      1.96s  5.26%  runtime.findObject
     1.45s  3.89% 38.86%      1.45s  3.89%  runtime.memclrNoHeapPointers
     1.44s  3.87% 42.72%      1.44s  3.87%  runtime.memmove
     1.35s  3.63% 46.35%      5.42s 14.55%  runtime.scanObject
     1.22s  3.28% 49.62%      1.22s  3.28%  aeshashbody
     0.95s  2.55% 52.18%      1.83s  4.91%  runtime.tryDeferToSpanScan
     0.66s  1.77% 53.95%      0.66s  1.77%  runtime.spanClass.sizeclass (inline)
     0.52s  1.40% 55.34%      2.17s  5.83%  runtime.mapaccess1_faststr
     0.43s  1.15% 56.50%      0.43s  1.15%  memeqbody
     0.41s  1.10% 57.60%      0.41s  1.10%  runtime.futex
     0.39s  1.05% 58.65%      0.61s  1.64%  example.com/isureal-company/isureal.postLess
     0.39s  1.05% 59.69%      0.39s  1.05%  runtime.nextFreeFast (inline)
     0.30s  0.81% 60.50%      0.69s  1.85%  database/sql.convertAssignRows
     0.28s  0.75% 61.25%      0.28s  0.75%  sync/atomic.(*Int32).Add
     0.26s   0.7% 61.95%      0.29s  0.78%  encoding/json.appendString[go.shape.string]
     0.25s  0.67% 62.62%      0.25s  0.67%  indexbody
     0.24s  0.64% 63.27%      0.25s  0.67%  runtime.spanOf (inline)
     0.23s  0.62% 63.88%      4.19s 11.25%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.23s  0.62% 64.50%      1.06s  2.85%  runtime.mallocgcSmallScanNoHeader
     0.21s  0.56% 65.06%      0.38s  1.02%  runtime.greyobject
     0.21s  0.56% 65.63%      3.75s 10.07%  runtime.mallocgc
     0.21s  0.56% 66.19%      1.67s  4.48%  runtime.mapassign_faststr
     0.20s  0.54% 66.73%      4.70s 12.62%  crypto/internal/fips140/sha256.(*Digest).Write
     0.18s  0.48% 67.21%      4.71s 12.65%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.17s  0.46% 67.67%      0.29s  0.78%  internal/runtime/maps.(*Iter).Next
     0.16s  0.43% 68.10%      4.43s 11.90%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.16s  0.43% 68.53%      0.19s  0.51%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.16s  0.43% 68.96%      0.19s  0.51%  runtime.typePointers.next
     0.15s   0.4% 69.36%      0.30s  0.81%  compress/flate.(*compressor).deflate
     0.15s   0.4% 69.76%      0.52s  1.40%  runtime.mapaccess2_faststr
     0.12s  0.32% 70.09%      0.51s  1.37%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
     0.11s   0.3% 70.38%      0.27s  0.73%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.11s   0.3% 70.68%      2.79s  7.49%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s   0.3% 70.97%      0.39s  1.05%  internal/stringslite.Index
     0.10s  0.27% 71.24%      0.55s  1.48%  runtime.convTstring
     0.10s  0.27% 71.51%      0.40s  1.07%  runtime.wbBufFlush1
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 47.4% / 100% | 17% | 0.5% | 9 | 106 MB | run queue が積もっている |
| isu2 (2 cpu) | 18.0% / 82% | 0% | 0.4% | 3 | 1789 MB |  |
| isu3 (2 cpu) | 15.8% / 68% | 0% | 0.6% | 4 | 1699 MB |  |
| isu4 (2 cpu) | 59.8% / 100% | 35% | 16.5% | 42 | 115 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 66% / nginx 13% |
| isu2 | isureal 34% / nginx 1% |
| isu3 | isureal 29% / nginx 1% |
| isu4 | postgres 68% / isureal 3% |

## Diff from previous run

スコア: 104121 → 119317 (+15196)

`8b07f57..2d10aad`

```
 conf/_all/etc/sysctl.d/99-isucon.conf          |  2 +
 conf/isu1/etc/nginx/sites-enabled/isureal.conf |  7 +-
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |  7 +-
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |  7 +-
 webapp/go/account_deletion_api.go              |  1 +
 webapp/go/action_log.go                        | 13 +++-
 webapp/go/action_log_test.go                   | 67 ++++++++++++++++++
 webapp/go/jpegcheck/huffman.go                 | 27 +++++++-
 webapp/go/jpegcheck/reader.go                  | 32 +++++----
 webapp/go/jpegcheck/scan.go                    |  7 +-
 webapp/go/offload.go                           | 47 +++++++++----
 webapp/go/post_api.go                          | 16 -----
 webapp/go/profile_image_api.go                 |  7 +-
 webapp/go/store.go                             | 53 +++++++++++++--
 webapp/go/user_api.go                          |  1 +
 webapp/go/user_read_api.go                     | 58 +++++++++-------
 webapp/go/user_search_test.go                  | 94 ++++++++++++++++++++++++++
 17 files changed, 357 insertions(+), 89 deletions(-)
```
