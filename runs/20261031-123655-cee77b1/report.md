# run 20261031-123655-cee77b1

## Summary

| | |
|---|---|
| score | 100484 |
| status | pass |
| ref | grandprix-jpeg-slots32 |
| sha | cee77b1773ea |
| 所要時間 | 349s |
| 前回比 | +17827 (21.6%) vs 20261031-122906-3a360a3 |
| ベスト比 | +17827 (21.6%) vs 20261031-122906-3a360a3 = 82657 **ベスト更新** |
| note | jpeg fifo 32 + single-copy multipart |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	53	157	445	933	932
europe	7	15	39	92	92	91
america	6	13	34	59	64	64
```

```
{"id":"9188","status":"done","portalStatus":"finished","passed":true,"score":100484,"raw":100484,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:41:21.923Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:41:24.887Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:41:27.889Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:41:30.909Z level=INFO msg="asia 2 日目: 53 人"
time=2026-10-31T03:41:33.892Z level=INFO msg="europe 2 日目: 15 人"
time=2026-10-31T03:41:36.888Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T03:41:39.909Z level=INFO msg="asia 3 日目: 157 人"
time=2026-10-31T03:41:40.328Z level=INFO msg="user moved" user=nita-5497 from=asia to=europe
time=2026-10-31T03:41:42.892Z level=INFO msg="europe 3 日目: 39 人"
time=2026-10-31T03:41:42.908Z level=INFO msg="user moved" user=nita-5497 from=europe to=america
time=2026-10-31T03:41:43.605Z level=INFO msg="user moved" user=evaul.6198 from=asia to=america
time=2026-10-31T03:41:45.823Z level=INFO msg="user moved" user=evaul.6198 from=america to=asia
time=2026-10-31T03:41:45.888Z level=INFO msg="america 3 日目: 34 人"
time=2026-10-31T03:41:48.919Z level=INFO msg="asia 4 日目: 445 人"
time=2026-10-31T03:41:51.281Z level=INFO msg="user moved" user=evaul.6198 from=asia to=america
time=2026-10-31T03:41:52.072Z level=INFO msg="europe 4 日目: 92 人"
time=2026-10-31T03:41:55.123Z level=INFO msg="america 4 日目: 59 人"
time=2026-10-31T03:41:56.061Z level=INFO msg="user moved" user=yano-7294 from=asia to=america
time=2026-10-31T03:41:56.227Z level=INFO msg="user moved" user=leonenoa6939 from=asia to=america
time=2026-10-31T03:41:56.379Z level=INFO msg="user moved" user=tamuivy.6886 from=asia to=america
time=2026-10-31T03:41:56.380Z level=INFO msg="user moved" user=benmuel7486 from=asia to=america
time=2026-10-31T03:41:57.014Z level=INFO msg="user moved" user=sachitor_8080 from=asia to=america
time=2026-10-31T03:41:58.110Z level=INFO msg="asia 5 日目: 933 人"
time=2026-10-31T03:42:01.000Z level=INFO msg="europe 5 日目: 92 人"
time=2026-10-31T03:42:03.932Z level=INFO msg="america 5 日目: 64 人"
time=2026-10-31T03:42:13.004Z level=INFO msg="america 最終日（グランプリ）: 64 人"
time=2026-10-31T03:42:13.045Z level=INFO msg="europe 最終日（グランプリ）: 91 人"
time=2026-10-31T03:42:13.227Z level=INFO msg="asia 最終日（グランプリ）: 932 人"
time=2026-10-31T03:42:21.911Z level=INFO msg="グランプリ: アジア 387 人, ヨーロッパ 14 人, アメリカ 11 人"
{
  "pass": true,
  "score": 100484,
  "messages": [
    "solhisa-5039: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.1%	1545	2	0	GET	/api/notifications	12313.517	7.970	15.169	15.187	167
16.0%	2196	64	0	POST	/api/posts	6119.003	2.786	9.833	11.060	631
13.4%	4951	91	0	POST	/api/friend-requests	5137.511	1.038	8.090	13.376	113
8.5%	13426	885	0	GET	^/api/images/[^/]+	3265.811	0.243	4.374	11.614	138271
7.1%	6040	2001	0	POST	/api/auth/login	2722.545	0.451	5.582	13.168	173
5.2%	1716	9	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1983.456	1.156	8.129	12.097	116
3.0%	956	1	0	POST	/api/auth/register	1136.851	1.189	7.331	8.297	227
2.4%	644	123	0	PUT	/api/me/profile-image	936.828	1.455	7.502	9.947	4
1.9%	546	32	0	POST	^/api/posts/[^/]+/realmojis$	716.480	1.312	10.003	13.180	442
1.6%	653	0	0	PUT	^/api/follows/[^/]+$	596.901	0.914	5.523	11.938	0
1.2%	405	6	0	POST	^/api/posts/[^/]+/comments$	478.745	1.182	7.151	13.966	334
1.2%	455	0	0	PATCH	/api/me	455.235	1.001	6.311	12.348	277
1.1%	3355	0	0	GET	/api/memories	404.817	0.121	0.673	0.898	10762
1.0%	11881	0	0	GET	^/api/feeds/[^/]+	379.800	0.032	0.204	3.200	2427
0.8%	656	0	0	POST	/api/auth/logout	312.384	0.476	6.644	12.354	0
0.8%	215	11	0	PATCH	^/api/posts/[^/]+$	292.050	1.358	7.707	13.063	611
0.5%	221	0	0	GET	/api/me/export	189.582	0.858	6.001	6.743	1829943
0.5%	85	0	0	DELETE	^/api/posts/[^/]+$	174.498	2.053	8.993	8.993	0
0.3%	3202	0	0	GET	/api/posting-window/current	120.659	0.038	0.178	3.862	237
0.2%	73	0	0	DELETE	^/api/comments/[^/]+$	79.404	1.088	11.591	11.591	0
0.2%	2469	0	0	GET	/api/users	78.069	0.032	0.171	1.310	241
0.2%	79	0	0	DELETE	^/api/realmojis/[^/]+$	71.310	0.903	9.768	9.768	0
0.2%	1813	81	0	GET	^/api/posts/me/current$	64.109	0.035	0.204	0.331	621
0.1%	41	0	0	PUT	^/api/blocks/[^/]+$	46.177	1.126	6.726	6.726	0
0.1%	1612	0	0	GET	^/api/relationships/[^/]+	43.686	0.027	0.198	0.479	889
0.1%	606	0	0	GET	^/api/memories/[^/]+	41.775	0.069	0.556	7.648	633
0.1%	1738	18	0	GET	^/api/posts/[^/]+$	35.130	0.020	0.157	0.277	633
0.0%	851	0	0	GET	^/api/users/[^/]+$	18.196	0.021	0.131	0.239	214
0.0%	792	0	0	GET	^/api/posts/[^/]+/realmojis$	16.292	0.021	0.164	0.283	1046
0.0%	588	0	0	GET	^/api/posts/[^/]+/comments$	12.402	0.021	0.139	0.199	1126
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.2%	231	1	0	GET	/api/notifications	1923.455	8.327	11.943	11.943	173
19.3%	273	5	0	POST	/api/posts	971.237	3.558	10.709	11.143	640
9.4%	517	8	0	POST	/api/friend-requests	475.453	0.920	7.366	10.897	114
5.8%	344	0	0	POST	/api/auth/login	290.727	0.845	6.887	9.174	246
5.4%	1223	8	0	GET	^/api/images/[^/]+	271.269	0.222	1.483	2.081	177903
4.2%	208	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	213.689	1.027	9.119	10.102	116
3.7%	1088	0	0	GET	^/api/feeds/[^/]+	186.378	0.171	0.534	1.023	2334
1.7%	338	0	0	GET	/api/memories	88.011	0.260	1.244	1.640	12775
1.6%	381	0	0	GET	/api/posting-window/current	79.308	0.208	0.658	1.172	239
1.0%	24	0	0	GET	/api/me/export	51.557	2.148	9.835	9.835	1801245
0.9%	70	2	0	POST	^/api/posts/[^/]+/comments$	45.983	0.657	3.670	3.670	310
0.9%	90	0	0	POST	/api/auth/register	44.942	0.499	2.271	2.271	228
0.9%	61	0	0	POST	/api/auth/logout	42.859	0.703	4.255	4.255	0
0.8%	242	0	0	GET	/api/users	40.927	0.169	0.515	0.544	242
0.8%	207	11	0	GET	^/api/posts/me/current$	39.763	0.192	0.602	0.721	616
0.6%	188	11	0	GET	^/api/posts/[^/]+$	30.360	0.161	0.515	0.517	593
0.6%	73	0	0	PUT	^/api/follows/[^/]+$	29.004	0.397	2.966	2.966	0
0.6%	27	1	0	PATCH	^/api/posts/[^/]+$	27.791	1.029	3.623	3.623	637
0.5%	50	0	0	PUT	/api/me/profile-image	26.885	0.538	1.868	1.868	0
0.5%	30	3	0	POST	^/api/posts/[^/]+/realmojis$	25.550	0.852	6.227	6.227	472
0.5%	10	0	0	DELETE	^/api/posts/[^/]+$	24.608	2.461	6.924	6.924	0
0.5%	163	0	0	GET	^/api/relationships/[^/]+	24.403	0.150	0.485	0.537	1202
0.4%	41	0	0	PATCH	/api/me	17.677	0.431	1.980	1.980	284
0.3%	73	0	0	GET	^/api/users/[^/]+$	14.074	0.193	0.714	0.714	219
0.2%	49	0	0	GET	^/api/memories/[^/]+	11.420	0.233	1.139	1.139	672
0.2%	4	0	0	DELETE	^/api/realmojis/[^/]+$	10.017	2.504	4.075	4.075	0
0.2%	60	1	0	GET	^/api/posts/[^/]+/comments$	9.554	0.159	0.489	0.489	1143
0.2%	55	2	0	GET	/api/me	8.231	0.150	0.270	0.270	290
0.2%	53	0	0	GET	^/api/posts/[^/]+/realmojis$	8.159	0.154	0.308	0.308	1018
0.1%	1	0	0	POST	/api/me/account-deletion	2.531	2.531	2.531	2.531	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.6%	171	0	0	GET	/api/notifications	1325.150	7.749	9.306	9.307	179
17.9%	194	5	0	POST	/api/posts	613.442	3.162	10.356	10.919	642
8.4%	296	5	0	POST	/api/friend-requests	287.943	0.973	6.798	11.450	114
6.6%	816	5	0	GET	^/api/images/[^/]+	225.338	0.276	2.239	4.245	169673
5.2%	234	0	0	POST	/api/auth/login	180.193	0.770	6.975	8.116	250
5.0%	123	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	170.723	1.388	9.923	10.848	116
3.6%	650	0	0	GET	^/api/feeds/[^/]+	122.785	0.189	0.600	0.848	2034
1.9%	232	0	0	GET	/api/memories	66.586	0.287	1.296	2.698	14100
1.8%	22	0	0	GET	/api/me/export	62.116	2.823	7.597	7.597	2888129
1.7%	51	5	0	POST	^/api/posts/[^/]+/realmojis$	57.934	1.136	8.648	8.648	441
1.5%	258	0	0	GET	/api/posting-window/current	52.912	0.205	0.636	1.290	239
0.9%	150	8	0	GET	^/api/posts/me/current$	30.756	0.205	0.670	0.700	627
0.8%	34	0	0	PUT	^/api/follows/[^/]+$	26.084	0.767	4.629	4.629	0
0.7%	134	0	0	GET	/api/users	23.873	0.178	0.562	0.691	230
0.7%	31	0	0	POST	/api/auth/logout	23.148	0.747	4.887	4.887	0
0.6%	56	0	0	POST	/api/auth/register	22.049	0.394	2.901	2.901	229
0.6%	30	0	0	PATCH	/api/me	19.620	0.654	3.041	3.041	284
0.5%	18	2	0	POST	^/api/posts/[^/]+/comments$	18.170	1.009	2.790	2.790	349
0.5%	29	0	0	PUT	/api/me/profile-image	17.243	0.595	2.882	2.882	0
0.4%	86	0	0	GET	^/api/relationships/[^/]+	14.867	0.173	0.657	0.657	1277
0.4%	12	1	0	PATCH	^/api/posts/[^/]+$	12.414	1.034	4.282	4.282	607
0.4%	65	3	0	GET	^/api/posts/[^/]+$	12.195	0.188	0.458	0.458	618
0.3%	43	0	0	GET	^/api/memories/[^/]+	9.504	0.221	0.827	0.827	1175
0.2%	35	0	0	GET	^/api/users/[^/]+$	5.636	0.161	0.408	0.408	187
0.2%	3	0	0	DELETE	^/api/comments/[^/]+$	5.429	1.810	2.425	2.425	0
0.1%	5	0	0	DELETE	^/api/posts/[^/]+$	5.096	1.019	2.414	2.414	0
0.1%	33	0	0	GET	^/api/posts/[^/]+/realmojis$	4.820	0.146	0.337	0.337	1348
0.1%	32	0	0	GET	/api/me	4.657	0.146	0.469	0.469	294
0.1%	9	0	0	DELETE	^/api/realmojis/[^/]+$	3.932	0.437	1.961	1.961	0
0.1%	2	0	0	DELETE	^/api/friendships/[^/]+$	3.357	1.679	2.641	2.641	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 64057 クエリ / 82 種類 / 合計 367.459s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	61.3%	3780	225.104	59.55	406.37	624.0	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	23.2%	267	85.396	319.83	1855.90	6640.5	122.2	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.3%	8529	19.471	2.28	41.90	10.7	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	4.9%	1749	18.153	10.38	458.29	118.2	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	1.8%	2591	6.458	2.49	38.71	25.0	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.6%	1	2.157	2156.52	2156.52	30014.0	1199367.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.5%	587	1.927	3.28	19.31	29.1	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
8	0.4%	1	1.537	1537.36	1537.36	21555.0	625291.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.3%	5640	1.237	0.22	25.27	19.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	0.2%	483	0.914	1.89	13.48	20.7	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.2%	1	0.733	732.68	732.68	11708.0	529641.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
12	0.2%	1983	0.623	0.31	19.62	20.7	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
13	0.2%	4617	0.605	0.13	17.58	11.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
14	0.1%	2	0.312	155.79	158.85	211856.0	2492.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
15	0.1%	2	0.272	135.96	263.59	16793.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: e2599e1cb56f82e7fa519019c7ef865c4790f21d
Type: cpu
Time: 2026-10-31 12:41:04 JST
Duration: 60s, Total samples = 48.68s (81.13%)
Showing nodes accounting for 34.70s, 71.28% of 48.68s total
Dropped 1106 nodes (cum <= 0.24s)
Showing top 40 nodes out of 297
      flat  flat%   sum%        cum   cum%
     4.17s  8.57%  8.57%      4.17s  8.57%  internal/runtime/syscall/linux.Syscall6
     2.91s  5.98% 14.54%      2.91s  5.98%  crypto/internal/fips140/sha256.blockAVX2
     2.59s  5.32% 19.86%      8.98s 18.45%  runtime.scanObject
     2.44s  5.01% 24.88%      2.44s  5.01%  runtime.memmove
     2.28s  4.68% 29.56%      2.28s  4.68%  runtime.memclrNoHeapPointers
     2.22s  4.56% 34.12%      2.83s  5.81%  runtime.findObject
     1.88s  3.86% 37.98%      3.24s  6.66%  runtime.tryDeferToSpanScan
     1.72s  3.53% 41.52%      5.58s 11.46%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.48s  3.04% 44.56%      1.48s  3.04%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.39s  2.86% 47.41%      1.39s  2.86%  aeshashbody
        1s  2.05% 49.47%      3.16s  6.49%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.96s  1.97% 51.44%      0.96s  1.97%  runtime.spanClass.sizeclass (inline)
     0.95s  1.95% 53.39%      0.95s  1.95%  example.com/isureal-company/isureal.(*UserRow).active
     0.94s  1.93% 55.32%      1.49s  3.06%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.75s  1.54% 56.86%      0.75s  1.54%  memeqbody
     0.62s  1.27% 58.13%      0.66s  1.36%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.51s  1.05% 59.18%      0.51s  1.05%  indexbody
     0.46s  0.94% 60.13%      0.69s  1.42%  example.com/isureal-company/isureal.postLess
     0.46s  0.94% 61.07%      2.15s  4.42%  runtime.mapaccess1_faststr
     0.39s   0.8% 61.87%      0.39s   0.8%  runtime.nextFreeFast (inline)
     0.37s  0.76% 62.63%      0.37s  0.76%  crypto/internal/fips140/aes/gcm.gcmAesEnc
     0.36s  0.74% 63.37%      0.58s  1.19%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.30s  0.62% 63.99%      0.31s  0.64%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.29s   0.6% 64.59%      0.61s  1.25%  database/sql.convertAssignRows
     0.28s  0.58% 65.16%      0.28s  0.58%  runtime.futex
     0.26s  0.53% 65.69%      0.36s  0.74%  runtime.(*mspan).markBitsForIndex (inline)
     0.25s  0.51% 66.21%      0.68s  1.40%  runtime.greyobject
     0.25s  0.51% 66.72%      0.28s  0.58%  runtime.heapArenaOf (inline)
     0.25s  0.51% 67.24%      0.29s   0.6%  runtime.spanOf (inline)
     0.24s  0.49% 67.73%      1.18s  2.42%  runtime.mallocgcSmallScanNoHeader
     0.23s  0.47% 68.20%      2.13s  4.38%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.23s  0.47% 68.67%      2.06s  4.23%  runtime.mapassign_faststr
     0.18s  0.37% 69.04%      0.36s  0.74%  internal/runtime/maps.(*Iter).Next
     0.18s  0.37% 69.41%      0.64s  1.31%  runtime.scanObjectsSmall
     0.17s  0.35% 69.76%      3.64s  7.48%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.17s  0.35% 70.11%      3.22s  6.61%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.17s  0.35% 70.46%      5.95s 12.22%  runtime.mallocgc
     0.14s  0.29% 70.75%      0.38s  0.78%  compress/flate.(*compressor).deflate
     0.13s  0.27% 71.01%      3.18s  6.53%  crypto/internal/fips140/sha256.(*Digest).Write
     0.13s  0.27% 71.28%      2.96s  6.08%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 51.8% / 100% | 33% | 0.1% | 7 | 172 MB | CPU が飽和 |
| isu2 (2 cpu) | 19.8% / 64% | 0% | 0.4% | 3 | 2623 MB |  |
| isu3 (2 cpu) | 17.4% / 60% | 0% | 0.5% | 4 | 2707 MB |  |
| isu4 (2 cpu) | 52.6% / 100% | 21% | 15.5% | 32 | 126 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 78% / nginx 17% / journalctl 3% / systemd-journal 2% |
| isu2 | isureal 38% |
| isu3 | isureal 33% / nginx 1% |
| isu4 | postgres 66% / isureal 4% |

## Diff from previous run

スコア: 82657 → 100484 (+17827)

`3a360a3..cee77b1`

```
 webapp/go/post_api.go | 62 ++++++++++++++++++++++++++++++++++++++++++---------
 1 file changed, 52 insertions(+), 10 deletions(-)
```
