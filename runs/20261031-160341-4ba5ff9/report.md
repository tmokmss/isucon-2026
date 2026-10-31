# run 20261031-160341-4ba5ff9

## Summary

| | |
|---|---|
| score | 175365 |
| status | pass |
| ref | image-etag-gzip |
| sha | 4ba5ff900bf0 |
| 所要時間 | 439s |
| 前回比 | +3748 (2.2%) vs 20261031-155850-66c7da8 |
| ベスト比 | -16451 (-8.6%) vs 20261031-151339-5a230a4 = 191816 |
| note | gzip-json-only |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	142	409	1127	1128
europe	7	21	72	203	548	547
america	6	17	44	127	249	248
```

```
{"id":"26514","status":"done","portalStatus":"finished","passed":true,"score":175365,"raw":175365,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:09:17.479Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:09:20.480Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:09:23.474Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:09:26.595Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T07:09:27.125Z level=INFO msg="user moved" user=riis.5175 from=asia to=america
time=2026-10-31T07:09:29.486Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T07:09:32.475Z level=INFO msg="america 2 日目: 17 人"
time=2026-10-31T07:09:35.492Z level=INFO msg="asia 3 日目: 142 人"
time=2026-10-31T07:09:36.383Z level=INFO msg="user moved" user=kuyu_5535 from=asia to=europe
time=2026-10-31T07:09:37.285Z level=INFO msg="user moved" user=seliaor_5744 from=asia to=europe
time=2026-10-31T07:09:37.800Z level=INFO msg="user moved" user=danuldan5885 from=asia to=europe
time=2026-10-31T07:09:38.413Z level=INFO msg="user moved" user=ulnoawa-6010 from=asia to=europe
time=2026-10-31T07:09:38.474Z level=INFO msg="europe 3 日目: 72 人"
time=2026-10-31T07:09:41.474Z level=INFO msg="america 3 日目: 44 人"
time=2026-10-31T07:09:42.561Z level=INFO msg="user moved" user=nazoe.6419 from=europe to=america
time=2026-10-31T07:09:44.486Z level=INFO msg="asia 4 日目: 409 人"
time=2026-10-31T07:09:46.186Z level=INFO msg="user moved" user=riis.5175 from=america to=europe
time=2026-10-31T07:09:47.709Z level=INFO msg="europe 4 日目: 203 人"
time=2026-10-31T07:09:48.431Z level=INFO msg="user moved" user=ruevahe_7042 from=asia to=america
time=2026-10-31T07:09:50.260Z level=INFO msg="user moved" user=leolia8249 from=asia to=america
time=2026-10-31T07:09:50.663Z level=INFO msg="america 4 日目: 127 人"
time=2026-10-31T07:09:50.669Z level=INFO msg="user moved" user=kiel-8041 from=asia to=europe
time=2026-10-31T07:09:51.937Z level=INFO msg="user moved" user=nazoe.6419 from=america to=europe
time=2026-10-31T07:09:53.329Z level=INFO msg="user moved" user=tokaso7950 from=asia to=america
time=2026-10-31T07:09:53.621Z level=INFO msg="asia 5 日目: 1127 人"
time=2026-10-31T07:09:57.115Z level=INFO msg="user moved" user=danuldan5885 from=europe to=asia
time=2026-10-31T07:09:57.453Z level=INFO msg="europe 5 日目: 548 人"
time=2026-10-31T07:09:58.000Z level=INFO msg="user moved" user=kuyu_5535 from=europe to=america
time=2026-10-31T07:10:00.894Z level=INFO msg="america 5 日目: 249 人"
time=2026-10-31T07:10:09.038Z level=INFO msg="user moved" user=yukomu-10238 from=america to=asia
time=2026-10-31T07:10:10.568Z level=INFO msg="america 最終日（グランプリ）: 248 人"
time=2026-10-31T07:10:11.695Z level=INFO msg="europe 最終日（グランプリ）: 547 人"
time=2026-10-31T07:10:12.683Z level=INFO msg="asia 最終日（グランプリ）: 1128 人"
time=2026-10-31T07:10:17.618Z level=INFO msg="グランプリ: アジア 273 人, ヨーロッパ 194 人, アメリカ 122 人"
{
  "pass": true,
  "score": 175365,
  "messages": [
    "elias_hoffmann_8: 履歴を読めません: GET /api/memories?month=2026-05: 応答がタイムアウトしました",
    "raranu.15861: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yoshimax.16315: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "munutsu17259: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "vinchiben-15284: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "namia.17058: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "funuul.14822: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sanufu_17523: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rani.17043: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "soshisol.17272: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "muma.16715: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "liais_17617: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hemia.15970: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "koke.16516: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "toleo18322: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "michiki.10695: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "kunato13839: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.38.96:52432-\u003e52.69.202.109:8080: write tcp 10.27.38.96:52432-\u003e52.69.202.109:8080: use of closed network connection",
    "ulzoeyo.17113: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "roshi_16640: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nora_18033: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.1%	1753	1	0	GET	/api/notifications	18879.120	10.770	19.071	19.087	173
11.8%	6486	87	0	POST	/api/friend-requests	4952.944	0.764	3.232	4.559	114
8.7%	2743	274	0	POST	/api/posts	3650.368	1.331	4.410	5.406	589
5.4%	19168	323	0	GET	^/api/images/[^/]+	2263.034	0.118	1.492	3.273	118506
4.5%	6789	687	0	POST	/api/auth/login	1894.410	0.279	2.909	5.331	222
3.6%	2806	14	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1504.072	0.536	3.018	4.675	116
3.5%	14545	0	0	GET	^/api/feeds/[^/]+	1450.787	0.100	0.808	2.745	1386
2.8%	4398	0	0	GET	/api/memories	1179.780	0.268	3.655	10.609	4218
1.8%	804	128	0	PUT	/api/me/profile-image	746.126	0.928	4.621	6.695	3
1.5%	570	6	0	POST	^/api/posts/[^/]+/realmojis$	638.814	1.121	5.540	6.843	486
1.2%	728	0	0	PUT	^/api/follows/[^/]+$	505.584	0.694	3.008	4.602	0
1.1%	595	7	0	POST	^/api/posts/[^/]+/comments$	464.834	0.781	4.882	6.061	325
1.0%	530	0	0	PATCH	/api/me	402.285	0.759	2.977	4.222	279
0.9%	3143	0	0	GET	/api/users	375.800	0.120	0.770	1.702	236
0.8%	358	0	0	GET	/api/me/export	345.051	0.964	15.569	16.685	1931347
0.8%	1152	1	0	POST	/api/auth/register	344.055	0.299	1.086	2.033	227
0.8%	3993	0	0	GET	/api/posting-window/current	318.985	0.080	0.665	1.740	237
0.7%	1080	0	0	POST	/api/auth/logout	310.474	0.287	2.989	5.019	0
0.6%	1958	43	0	GET	^/api/posts/me/current$	252.159	0.129	0.956	1.690	640
0.5%	235	11	0	PATCH	^/api/posts/[^/]+$	215.091	0.915	3.727	4.074	625
0.4%	2135	49	0	GET	^/api/posts/[^/]+$	187.739	0.088	0.672	1.690	606
0.4%	689	3	0	GET	/api/me	173.869	0.252	0.931	1.888	302
0.4%	1769	0	0	GET	^/api/relationships/[^/]+	157.089	0.089	0.766	2.357	540
0.3%	83	0	0	DELETE	^/api/posts/[^/]+$	130.577	1.573	5.414	5.414	0
0.2%	141	0	0	DELETE	^/api/comments/[^/]+$	93.880	0.666	4.032	4.724	0
0.2%	975	0	0	GET	^/api/memories/[^/]+	93.729	0.096	0.867	2.250	375
0.2%	105	0	0	DELETE	^/api/realmojis/[^/]+$	92.276	0.879	3.691	4.099	0
0.2%	1094	0	0	GET	^/api/users/[^/]+$	84.774	0.077	0.544	1.027	199
0.1%	787	0	0	GET	^/api/posts/[^/]+/realmojis$	60.248	0.077	0.761	1.871	911
0.1%	666	0	0	GET	^/api/posts/[^/]+/comments$	38.152	0.057	0.381	0.771	766
…（上位 30 / 全 44 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.6%	706	4	0	GET	/api/notifications	5777.812	8.184	14.209	14.213	149
14.8%	2624	21	0	POST	/api/friend-requests	2342.481	0.893	3.361	4.334	114
8.7%	1091	117	0	POST	/api/posts	1372.067	1.258	4.258	5.079	579
6.1%	1421	0	0	POST	/api/auth/login	968.949	0.682	3.646	5.723	243
4.2%	4416	23	0	GET	^/api/images/[^/]+	656.747	0.149	1.315	3.392	177227
4.0%	541	0	0	POST	/api/auth/register	636.458	1.176	2.675	3.294	228
3.9%	4320	0	0	GET	^/api/feeds/[^/]+	612.538	0.142	0.999	1.793	1361
3.7%	310	0	0	PUT	/api/me/profile-image	590.964	1.906	4.572	7.565	0
2.6%	337	0	0	PUT	^/api/follows/[^/]+$	408.751	1.213	3.618	4.059	0
2.4%	551	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	381.212	0.692	3.046	3.554	115
2.2%	258	0	0	PATCH	/api/me	343.079	1.330	3.956	5.691	281
1.7%	1694	0	0	GET	/api/posting-window/current	269.059	0.159	0.898	1.737	239
1.6%	1147	0	0	GET	/api/memories	254.561	0.222	1.958	15.465	3682
1.5%	237	8	0	POST	^/api/posts/[^/]+/realmojis$	241.373	1.018	4.815	5.494	458
1.0%	1288	0	0	GET	/api/users	159.734	0.124	0.642	1.682	239
0.6%	779	46	0	GET	^/api/posts/me/current$	99.035	0.127	1.208	1.389	609
0.5%	104	2	0	POST	^/api/posts/[^/]+/comments$	86.502	0.832	4.828	5.301	342
0.5%	636	34	0	GET	^/api/posts/[^/]+$	85.509	0.134	0.931	1.414	586
0.5%	73	1	0	PATCH	^/api/posts/[^/]+$	78.196	1.071	3.896	3.896	646
0.5%	182	0	0	POST	/api/auth/logout	77.580	0.426	4.327	4.336	0
0.4%	521	0	0	GET	^/api/relationships/[^/]+	59.934	0.115	0.901	1.068	389
0.3%	30	0	0	DELETE	^/api/posts/[^/]+$	53.402	1.780	4.413	4.413	0
0.3%	42	0	0	DELETE	^/api/realmojis/[^/]+$	52.218	1.243	3.368	3.368	0
0.3%	315	1	0	GET	/api/me	45.798	0.145	0.759	0.913	304
0.2%	321	0	0	GET	^/api/users/[^/]+$	34.953	0.109	0.475	0.901	198
0.1%	55	0	0	GET	/api/me/export	22.120	0.402	4.530	4.530	2319704
0.1%	221	0	0	GET	^/api/posts/[^/]+/realmojis$	21.717	0.098	0.296	0.732	1009
0.1%	170	0	0	GET	^/api/posts/[^/]+/comments$	18.209	0.107	0.709	0.746	777
0.1%	21	0	0	DELETE	^/api/comments/[^/]+$	16.017	0.763	4.230	4.230	0
0.1%	167	0	0	GET	^/api/memories/[^/]+	14.607	0.087	0.282	0.331	564
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.1%	374	1	0	GET	/api/notifications	2880.107	7.701	10.234	10.917	155
13.0%	1251	10	0	POST	/api/friend-requests	958.416	0.766	3.256	4.554	114
9.4%	528	51	0	POST	/api/posts	691.344	1.309	4.228	5.625	591
7.8%	922	0	0	POST	/api/auth/login	577.217	0.626	3.196	4.120	250
5.8%	3019	15	0	GET	^/api/images/[^/]+	423.562	0.140	1.101	3.423	136104
3.9%	244	0	0	POST	/api/auth/register	284.344	1.165	3.159	3.196	230
3.2%	1887	0	0	GET	^/api/feeds/[^/]+	238.070	0.126	0.934	1.470	1379
2.4%	276	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	173.813	0.630	2.891	3.776	116
2.2%	144	0	0	PUT	/api/me/profile-image	162.313	1.127	3.853	4.515	0
1.6%	149	0	0	PUT	^/api/follows/[^/]+$	117.495	0.789	3.049	3.433	0
1.5%	107	7	0	POST	^/api/posts/[^/]+/realmojis$	114.066	1.066	4.319	4.862	450
1.5%	764	0	0	GET	/api/memories	112.654	0.147	1.085	2.334	5386
1.5%	100	0	0	PATCH	/api/me	109.305	1.093	3.366	3.536	286
1.2%	780	0	0	GET	/api/posting-window/current	87.095	0.112	0.747	1.121	240
1.0%	663	0	0	GET	/api/users	71.739	0.108	0.535	0.935	236
0.8%	116	0	0	POST	/api/auth/logout	61.340	0.529	2.183	3.733	0
0.6%	414	13	0	GET	^/api/posts/me/current$	46.685	0.113	0.996	1.071	632
0.5%	49	0	0	POST	^/api/posts/[^/]+/comments$	34.940	0.713	4.725	4.725	321
0.5%	259	0	0	GET	^/api/relationships/[^/]+	33.710	0.130	0.869	1.049	458
0.4%	247	10	0	GET	^/api/posts/[^/]+$	31.410	0.127	1.022	1.246	586
0.3%	18	0	0	DELETE	^/api/posts/[^/]+$	22.619	1.257	2.434	2.434	0
0.3%	41	6	0	PATCH	^/api/posts/[^/]+$	22.618	0.552	2.156	2.156	558
0.3%	45	0	0	GET	/api/me/export	20.925	0.465	9.963	9.963	2792504
0.2%	140	0	0	GET	^/api/memories/[^/]+	14.791	0.106	0.443	1.022	548
0.2%	125	0	0	GET	^/api/users/[^/]+$	14.755	0.118	0.748	0.881	189
0.2%	148	0	0	GET	/api/me	13.646	0.092	0.208	0.897	304
0.1%	100	0	0	GET	^/api/posts/[^/]+/realmojis$	10.612	0.106	0.912	0.912	1105
0.1%	73	0	0	GET	^/api/posts/[^/]+/comments$	8.073	0.111	0.914	0.914	668
0.1%	14	0	0	DELETE	^/api/realmojis/[^/]+$	6.380	0.456	1.329	1.329	0
0.1%	1	0	0	DELETE	/api/me/profile-image	5.509	5.509	5.509	5.509	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16343 クエリ / 48 種類 / 合計 148.648s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	55.1%	367	81.865	223.06	1632.92	1227.5	125.7	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.4%	3	21.424	7141.45	9535.77	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.4%	13601	16.981	1.25	7.30	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.8%	3	16.106	5368.61	7319.26	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.2%	3	7.779	2592.93	3541.15	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	1.5%	4	2.281	570.14	1556.66	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.2%	4	0.356	89.06	161.06	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.2%	4	0.356	89.05	164.22	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.280	279.87	279.87	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.258	258.14	258.14	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.2%	1	0.239	238.78	238.78	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.162	162.36	162.36	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.109	109.40	109.40	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.080	79.50	79.50	2412.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.074	73.52	73.52	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT data FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 7d34a202997143424da2a3e80ec0d525a72fa058
Type: cpu
Time: 2026-10-31 16:08:57 JST
Duration: 60.60s, Total samples = 44.93s (74.14%)
Showing nodes accounting for 31.38s, 69.84% of 44.93s total
Dropped 1144 nodes (cum <= 0.22s)
Showing top 40 nodes out of 293
      flat  flat%   sum%        cum   cum%
     5.12s 11.40% 11.40%      5.12s 11.40%  crypto/internal/fips140/sha256.blockAVX2
     5.05s 11.24% 22.64%      5.05s 11.24%  internal/runtime/syscall/linux.Syscall6
     2.21s  4.92% 27.55%      2.21s  4.92%  runtime.memmove
     1.80s  4.01% 31.56%      2.10s  4.67%  runtime.findObject
     1.77s  3.94% 35.50%      1.77s  3.94%  runtime.memclrNoHeapPointers
     1.72s  3.83% 39.33%      6.72s 14.96%  runtime.scanObject
     1.65s  3.67% 43.00%      5.58s 12.42%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.59s  3.54% 46.54%      2.76s  6.14%  runtime.tryDeferToSpanScan
     1.55s  3.45% 49.99%      1.55s  3.45%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.46s  3.25% 53.24%      1.46s  3.25%  aeshashbody
     0.91s  2.03% 55.26%      0.91s  2.03%  runtime.spanClass.sizeclass (inline)
     0.59s  1.31% 56.58%      2.35s  5.23%  runtime.mapaccess1_faststr
     0.48s  1.07% 57.65%      0.48s  1.07%  runtime.nextFreeFast (inline)
     0.45s  1.00% 58.65%      0.45s  1.00%  memeqbody
     0.33s  0.73% 59.38%      0.75s  1.67%  database/sql.convertAssignRows
     0.31s  0.69% 60.07%      0.39s  0.87%  encoding/json.appendString[go.shape.string]
     0.31s  0.69% 60.76%      0.31s  0.69%  runtime.typePointers.next
     0.30s  0.67% 61.43%      0.58s  1.29%  example.com/isureal-company/isureal.postLess
     0.29s  0.65% 62.07%      0.64s  1.42%  runtime.greyobject
     0.27s   0.6% 62.68%      0.59s  1.31%  runtime.scanObjectsSmall
     0.25s  0.56% 63.23%      1.14s  2.54%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.53% 63.77%      5.60s 12.46%  crypto/internal/fips140/sha256.(*Digest).Write
     0.20s  0.45% 64.21%      3.89s  8.66%  runtime.mallocgc
     0.19s  0.42% 64.63%      2.01s  4.47%  runtime.mapassign_faststr
     0.18s   0.4% 65.03%      0.46s  1.02%  compress/flate.(*compressor).deflate
     0.18s   0.4% 65.44%      6.24s 13.89%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.18s   0.4% 65.84%      5.57s 12.40%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.18s   0.4% 66.24%      0.27s   0.6%  runtime.(*mspan).markBitsForIndex (inline)
     0.17s  0.38% 66.61%      5.82s 12.95%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.17s  0.38% 66.99%      1.02s  2.27%  runtime.scanSpan
     0.15s  0.33% 67.33%      0.32s  0.71%  internal/runtime/maps.(*Iter).Next
     0.14s  0.31% 67.64%      1.18s  2.63%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.13s  0.29% 67.93%      0.95s  2.11%  encoding/json.structEncoder.encode
     0.13s  0.29% 68.22%      0.42s  0.93%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
     0.13s  0.29% 68.51%      2.78s  6.19%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.13s  0.29% 68.80%      0.93s  2.07%  runtime.mallocgcSmallNoscan
     0.13s  0.29% 69.09%      0.98s  2.18%  runtime.wbBufFlush1
     0.12s  0.27% 69.35%      0.68s  1.51%  runtime.bulkBarrierPreWrite
     0.11s  0.24% 69.60%      0.31s  0.69%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.11s  0.24% 69.84%      3.86s  8.59%  example.com/isureal-company/isureal.listMemories
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 44.5% / 100% | 27% | 0.7% | 9 | 101 MB | run queue が積もっている |
| isu2 (2 cpu) | 29.5% / 100% | 4% | 0.7% | 4 | 413 MB |  |
| isu3 (2 cpu) | 24.1% / 100% | 2% | 1.8% | 3 | 709 MB |  |
| isu4 (2 cpu) | 39.0% / 100% | 12% | 15.2% | 7 | 107 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 64% / nginx 14% |
| isu2 | isureal 55% / nginx 2% |
| isu3 | isureal 44% / nginx 2% |
| isu4 | isureal 28% / postgres 14% |

## Diff from previous run

スコア: 171617 → 175365 (+3748)

`66c7da8..4ba5ff9`

```
 webapp/go/image_api.go   | 13 -------------
 webapp/go/image_local.go |  3 ---
 2 files changed, 16 deletions(-)
```
