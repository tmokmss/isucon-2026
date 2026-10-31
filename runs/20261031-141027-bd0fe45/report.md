# run 20261031-141027-bd0fe45

## Summary

| | |
|---|---|
| score | 141335 |
| status | pass |
| ref | bd0fe45 |
| sha | bd0fe45bf2e2 |
| 所要時間 | 389s |
| 前回比 | +15206 (12.1%) vs 20261031-140122-bd0fe45 |
| ベスト比 | +15206 (12.1%) vs 20261031-140122-bd0fe45 = 126129 **ベスト更新** |
| note | 複製 local・時刻の印で待つ（main c6abd65 取り込み済み）2回目 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	135	386	930	930
europe	7	18	55	153	270	269
america	6	15	47	123	191	191
```

```
{"id":"16468","status":"done","portalStatus":"finished","passed":true,"score":141335,"raw":141335,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:15:28.845Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:15:31.808Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:15:34.808Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:15:37.841Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T05:15:40.807Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T05:15:43.808Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T05:15:46.836Z level=INFO msg="asia 3 日目: 135 人"
time=2026-10-31T05:15:48.288Z level=INFO msg="user moved" user=ultato-5618 from=asia to=america
time=2026-10-31T05:15:48.351Z level=INFO msg="user moved" user=ortormo.5633 from=asia to=europe
time=2026-10-31T05:15:49.260Z level=INFO msg="user moved" user=ortormo.5633 from=europe to=asia
time=2026-10-31T05:15:49.661Z level=INFO msg="user moved" user=sumaxsa5840 from=asia to=europe
time=2026-10-31T05:15:49.808Z level=INFO msg="europe 3 日目: 55 人"
time=2026-10-31T05:15:50.087Z level=INFO msg="user moved" user=benrosol.5993 from=asia to=america
time=2026-10-31T05:15:52.215Z level=INFO msg="user moved" user=ultato-5618 from=america to=europe
time=2026-10-31T05:15:52.809Z level=INFO msg="america 3 日目: 47 人"
time=2026-10-31T05:15:55.843Z level=INFO msg="asia 4 日目: 386 人"
time=2026-10-31T05:15:56.945Z level=INFO msg="user moved" user=hezoese-6448 from=america to=asia
time=2026-10-31T05:15:58.482Z level=INFO msg="user moved" user=yoevahi6884 from=asia to=europe
time=2026-10-31T05:15:58.829Z level=INFO msg="europe 4 日目: 153 人"
time=2026-10-31T05:15:58.897Z level=INFO msg="user moved" user=hezoese-6448 from=asia to=america
time=2026-10-31T05:15:59.577Z level=INFO msg="user moved" user=ortormo.5633 from=asia to=europe
time=2026-10-31T05:15:59.700Z level=INFO msg="user moved" user=ulsu-6970 from=asia to=america
time=2026-10-31T05:16:01.867Z level=INFO msg="america 4 日目: 123 人"
time=2026-10-31T05:16:03.214Z level=INFO msg="user moved" user=nihetsu7073 from=asia to=america
time=2026-10-31T05:16:04.988Z level=INFO msg="asia 5 日目: 930 人"
time=2026-10-31T05:16:07.954Z level=INFO msg="europe 5 日目: 270 人"
time=2026-10-31T05:16:10.941Z level=INFO msg="america 5 日目: 191 人"
time=2026-10-31T05:16:19.974Z level=INFO msg="asia 最終日（グランプリ）: 930 人"
time=2026-10-31T05:16:20.090Z level=INFO msg="europe 最終日（グランプリ）: 269 人"
time=2026-10-31T05:16:20.113Z level=INFO msg="america 最終日（グランプリ）: 191 人"
time=2026-10-31T05:16:28.863Z level=INFO msg="グランプリ: アジア 464 人, ヨーロッパ 136 人, アメリカ 104 人"
{
  "pass": true,
  "score": 141335,
  "messages": [
    "nozoe_5100: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "tsusu-6924: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rimima_14254: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sukiya-13638: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noavin.12681: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "nararu13164: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "elchi.12757: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "wawa.13822: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.4%	1537	0	0	GET	/api/notifications	12384.805	8.058	14.982	15.005	167
19.0%	2359	55	0	POST	/api/posts	6639.985	2.815	11.036	14.168	634
9.2%	5401	85	0	POST	/api/friend-requests	3223.507	0.597	3.916	8.237	114
7.8%	15018	390	0	GET	^/api/images/[^/]+	2744.257	0.183	2.703	7.475	146557
6.0%	5394	832	0	POST	/api/auth/login	2099.183	0.389	4.717	8.759	210
4.2%	2355	10	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1484.910	0.631	4.276	7.336	116
3.5%	723	38	0	POST	^/api/posts/[^/]+/realmojis$	1236.049	1.710	10.982	11.257	448
3.0%	638	125	0	PUT	/api/me/profile-image	1034.336	1.621	10.014	11.282	3
2.2%	953	1	0	POST	/api/auth/register	784.692	0.823	5.761	7.427	227
1.6%	681	12	0	POST	^/api/posts/[^/]+/comments$	561.164	0.824	4.856	7.953	334
1.4%	3639	0	0	GET	/api/memories	490.481	0.135	0.815	1.578	11737
1.2%	608	0	0	PUT	^/api/follows/[^/]+$	409.337	0.673	4.154	6.165	0
1.0%	468	0	0	PATCH	/api/me	364.443	0.779	5.768	7.371	273
0.9%	14116	0	0	GET	^/api/feeds/[^/]+	301.589	0.021	0.245	2.625	2192
0.7%	289	0	0	GET	/api/me/export	231.988	0.803	3.676	4.220	1825883
0.6%	747	0	0	POST	/api/auth/logout	201.856	0.270	3.019	6.212	0
0.4%	210	9	0	PATCH	^/api/posts/[^/]+$	147.670	0.703	5.039	5.953	618
0.3%	91	0	0	DELETE	^/api/posts/[^/]+$	112.347	1.235	5.236	5.236	0
0.2%	3382	0	0	GET	/api/posting-window/current	75.900	0.022	0.141	0.457	237
0.2%	106	0	0	DELETE	^/api/realmojis/[^/]+$	65.446	0.617	2.506	3.145	0
0.2%	109	0	0	DELETE	^/api/comments/[^/]+$	60.596	0.556	3.775	5.174	0
0.1%	1898	50	0	GET	^/api/posts/me/current$	41.167	0.022	0.214	0.559	631
0.1%	1800	0	0	GET	^/api/relationships/[^/]+	39.723	0.022	0.294	0.552	939
0.1%	671	0	0	GET	^/api/memories/[^/]+	38.630	0.058	0.610	0.843	614
0.1%	2031	20	0	GET	^/api/posts/[^/]+$	38.298	0.019	0.211	0.594	621
0.1%	2600	0	0	GET	/api/users	38.284	0.015	0.087	0.459	240
0.1%	37	0	0	PUT	^/api/blocks/[^/]+$	28.825	0.779	4.423	4.423	0
0.1%	1048	0	0	GET	^/api/users/[^/]+$	22.050	0.021	0.257	0.572	196
0.1%	1	0	0	POST	/api/initialize	19.727	19.727	19.727	19.727	90
0.0%	926	0	0	GET	^/api/posts/[^/]+/realmojis$	16.628	0.018	0.138	2.188	1752
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.1%	473	1	0	GET	/api/notifications	3757.365	7.944	11.955	13.015	162
12.0%	622	17	0	POST	/api/posts	1250.564	2.011	5.803	8.549	633
10.3%	1319	19	0	POST	/api/friend-requests	1069.953	0.811	4.233	10.569	114
7.2%	2652	16	0	GET	^/api/images/[^/]+	751.924	0.284	1.420	1.957	183227
5.7%	805	0	0	POST	/api/auth/login	594.109	0.738	4.105	6.390	246
5.0%	2118	0	0	GET	^/api/feeds/[^/]+	521.035	0.246	1.180	1.473	2289
3.9%	910	0	0	GET	/api/posting-window/current	402.633	0.442	1.397	2.065	239
2.5%	264	0	0	POST	/api/auth/register	262.031	0.993	4.929	5.827	229
2.2%	282	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	233.019	0.826	3.766	4.660	116
2.0%	625	0	0	GET	/api/users	208.280	0.333	1.175	1.505	238
1.9%	746	0	0	GET	/api/memories	192.811	0.258	1.646	3.090	11782
1.6%	145	0	0	PUT	/api/me/profile-image	170.408	1.175	4.639	6.673	0
1.3%	514	28	0	GET	^/api/posts/me/current$	135.868	0.264	1.060	1.557	615
1.2%	166	0	0	PUT	^/api/follows/[^/]+$	120.865	0.728	3.743	5.205	0
1.1%	139	0	0	PATCH	/api/me	116.134	0.835	3.771	4.083	273
0.8%	48	2	0	POST	^/api/posts/[^/]+/realmojis$	84.587	1.762	9.247	9.247	457
0.7%	114	1	0	POST	^/api/posts/[^/]+/comments$	71.649	0.629	4.611	5.246	329
0.7%	109	0	0	POST	/api/auth/logout	71.181	0.653	4.918	5.362	0
0.6%	282	0	0	GET	^/api/relationships/[^/]+	67.273	0.239	1.014	1.068	757
0.6%	325	12	0	GET	^/api/posts/[^/]+$	58.688	0.181	0.842	0.956	594
0.4%	150	2	0	GET	/api/me	46.677	0.311	0.824	0.962	301
0.4%	46	0	0	GET	/api/me/export	39.478	0.858	9.102	9.102	2407779
0.4%	42	2	0	PATCH	^/api/posts/[^/]+$	39.190	0.933	3.077	3.077	631
0.3%	163	0	0	GET	^/api/users/[^/]+$	30.669	0.188	0.938	0.964	196
0.2%	118	1	0	GET	^/api/posts/[^/]+/realmojis$	22.850	0.194	0.866	0.869	1160
0.2%	127	2	0	GET	^/api/posts/[^/]+/comments$	21.776	0.171	0.988	0.995	1532
0.2%	17	0	0	DELETE	^/api/posts/[^/]+$	21.395	1.259	3.598	3.598	0
0.2%	102	0	0	GET	^/api/memories/[^/]+	21.325	0.209	1.388	1.920	834
0.1%	5	0	0	DELETE	^/api/realmojis/[^/]+$	5.281	1.056	2.531	2.531	0
0.0%	15	0	0	DELETE	^/api/comments/[^/]+$	2.779	0.185	0.460	0.460	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.1%	334	0	0	GET	/api/notifications	2425.021	7.261	9.445	9.447	168
11.9%	954	14	0	POST	/api/friend-requests	845.429	0.886	4.989	7.752	114
10.7%	437	15	0	POST	/api/posts	762.711	1.745	4.801	6.252	625
8.2%	1972	11	0	GET	^/api/images/[^/]+	580.248	0.294	1.405	3.703	160389
6.5%	620	0	0	POST	/api/auth/login	458.516	0.740	4.190	5.705	249
4.1%	1318	0	0	GET	^/api/feeds/[^/]+	290.379	0.220	1.051	1.276	1902
3.6%	635	0	0	GET	/api/posting-window/current	254.863	0.401	1.164	1.509	239
3.0%	239	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	211.539	0.885	3.630	5.243	116
2.1%	574	0	0	GET	/api/memories	151.037	0.263	1.706	2.051	12762
2.1%	471	0	0	GET	/api/users	149.385	0.317	1.125	1.195	235
2.1%	49	6	0	POST	^/api/posts/[^/]+/realmojis$	149.270	3.046	10.424	10.424	412
1.6%	96	0	0	PUT	/api/me/profile-image	111.139	1.158	5.096	5.096	0
1.3%	187	0	0	POST	/api/auth/register	95.426	0.510	3.199	3.687	229
1.3%	376	27	0	GET	^/api/posts/me/current$	95.285	0.253	1.023	1.190	604
1.2%	113	0	0	PUT	^/api/follows/[^/]+$	82.543	0.730	3.759	4.109	0
1.1%	89	0	0	POST	/api/auth/logout	74.969	0.842	5.383	5.383	0
0.9%	83	0	0	PATCH	/api/me	63.211	0.762	5.020	5.020	275
0.7%	51	1	0	POST	^/api/posts/[^/]+/comments$	52.899	1.037	4.473	4.473	370
0.6%	29	0	0	GET	/api/me/export	44.735	1.543	9.685	9.685	3020800
0.5%	201	0	0	GET	^/api/relationships/[^/]+	36.020	0.179	0.949	0.963	960
0.4%	166	11	0	GET	^/api/posts/[^/]+$	28.361	0.171	0.978	1.132	576
0.4%	99	0	0	GET	/api/me	26.328	0.266	1.016	1.016	303
0.4%	83	0	0	GET	^/api/memories/[^/]+	24.962	0.301	1.685	1.685	1161
0.3%	35	5	0	PATCH	^/api/posts/[^/]+$	24.458	0.699	4.606	4.606	566
0.2%	86	0	0	GET	^/api/users/[^/]+$	13.869	0.161	1.132	1.132	194
0.2%	10	0	0	DELETE	^/api/comments/[^/]+$	12.021	1.202	4.304	4.304	0
0.1%	11	0	0	DELETE	^/api/posts/[^/]+$	10.392	0.945	2.430	2.430	0
0.1%	48	0	0	GET	^/api/posts/[^/]+/comments$	8.650	0.180	1.183	1.183	1921
0.1%	3	0	0	DELETE	^/api/follows/[^/]+$	5.952	1.984	2.292	2.292	0
0.1%	4	0	0	DELETE	^/api/realmojis/[^/]+$	4.554	1.139	1.941	1.941	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 84309 クエリ / 81 種類 / 合計 261.569s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	35.9%	305	93.971	308.10	1196.53	1187.8	121.2	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.0%	3958	36.586	9.24	642.60	117.6	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
3	12.8%	4864	33.361	6.86	145.13	545.7	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
4	9.4%	10192	24.513	2.41	72.14	11.1	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	8.1%	3	21.309	7102.83	9550.62	31353.0	1199922.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	6.2%	3	16.255	5418.26	7239.23	24949.0	628171.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	2.9%	3	7.668	2556.04	3428.11	13010.0	529812.7	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	2.8%	6	7.416	1236.02	6610.12	202963.3	2609.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
9	2.7%	3335	6.991	2.10	59.62	24.5	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.9%	774	2.365	3.06	36.92	28.5	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
11	0.7%	7546	1.847	0.24	37.54	19.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
12	0.5%	832	1.313	1.58	51.20	20.3	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
13	0.4%	2781	0.932	0.34	43.01	20.2	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
14	0.3%	6	0.910	151.59	319.42	16773.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	0.3%	5987	0.818	0.14	12.86	11.1	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
-- rank 3
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: b0705343cf26e9b0d9a736a4928287957110e813
Type: cpu
Time: 2026-10-31 14:15:03 JST
Duration: 60s, Total samples = 36.42s (60.70%)
Showing nodes accounting for 25.56s, 70.18% of 36.42s total
Dropped 975 nodes (cum <= 0.18s)
Showing top 40 nodes out of 296
      flat  flat%   sum%        cum   cum%
     4.16s 11.42% 11.42%      4.16s 11.42%  crypto/internal/fips140/sha256.blockAVX2
     3.94s 10.82% 22.24%      3.94s 10.82%  internal/runtime/syscall/linux.Syscall6
     1.60s  4.39% 26.63%      1.60s  4.39%  runtime.memmove
     1.49s  4.09% 30.72%      1.49s  4.09%  runtime.memclrNoHeapPointers
     1.43s  3.93% 34.65%      5.64s 15.49%  runtime.scanObject
     1.38s  3.79% 38.44%      5.48s 15.05%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.31s  3.60% 42.04%      1.31s  3.60%  aeshashbody
     1.27s  3.49% 45.52%      1.27s  3.49%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.26s  3.46% 48.98%      1.71s  4.70%  runtime.findObject
     1.01s  2.77% 51.76%      1.97s  5.41%  runtime.tryDeferToSpanScan
     0.67s  1.84% 53.60%      0.67s  1.84%  runtime.spanClass.sizeclass (inline)
     0.48s  1.32% 54.91%      1.79s  4.91%  runtime.mapaccess1_faststr
     0.35s  0.96% 55.88%      0.58s  1.59%  example.com/isureal-company/isureal.postLess
     0.30s  0.82% 56.70%      1.13s  3.10%  runtime.mallocgcSmallScanNoHeader
     0.30s  0.82% 57.52%      0.30s  0.82%  runtime.nextFreeFast (inline)
     0.28s  0.77% 58.29%      0.72s  1.98%  database/sql.convertAssignRows
     0.27s  0.74% 59.03%      0.27s  0.74%  runtime.futex
     0.26s  0.71% 59.75%      0.41s  1.13%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.25s  0.69% 60.43%      4.44s 12.19%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.25s  0.69% 61.12%      0.52s  1.43%  runtime.greyobject
     0.23s  0.63% 61.75%      4.11s 11.29%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.23s  0.63% 62.38%      0.26s  0.71%  runtime.(*mspan).writeHeapBitsSmall
     0.22s   0.6% 62.99%      0.22s   0.6%  memeqbody
     0.21s  0.58% 63.56%      1.77s  4.86%  runtime.mapassign_faststr
     0.21s  0.58% 64.14%      0.21s  0.58%  runtime.typePointers.nextFast (inline)
     0.20s  0.55% 64.69%      0.20s  0.55%  indexbody
     0.18s  0.49% 65.18%      4.54s 12.47%  crypto/internal/fips140/sha256.(*Digest).Write
     0.18s  0.49% 65.68%      0.22s   0.6%  runtime.heapArenaOf (inline)
     0.17s  0.47% 66.14%      0.38s  1.04%  internal/runtime/maps.(*Iter).Next
     0.16s  0.44% 66.58%      4.29s 11.78%  runtime.mallocgc
     0.15s  0.41% 67.00%      4.73s 12.99%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.14s  0.38% 67.38%      3.08s  8.46%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.14s  0.38% 67.76%      1.28s  3.51%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.14s  0.38% 68.15%      0.19s  0.52%  runtime.(*mspan).markBitsForIndex (inline)
     0.14s  0.38% 68.53%      0.36s  0.99%  runtime.scanObjectsSmall
     0.13s  0.36% 68.89%      0.23s  0.63%  runtime.scanObjectSmall
     0.13s  0.36% 69.25%      0.74s  2.03%  runtime.scanSpan
     0.12s  0.33% 69.58%      0.27s  0.74%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.11s   0.3% 69.88%      0.26s  0.71%  compress/flate.(*compressor).deflate
     0.11s   0.3% 70.18%      0.21s  0.58%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 50.9% / 100% | 36% | 0.2% | 6 | 156 MB | CPU が飽和 |
| isu2 (2 cpu) | 29.4% / 100% | 4% | 0.8% | 4 | 158 MB |  |
| isu3 (2 cpu) | 29.2% / 100% | 3% | 1.4% | 3 | 155 MB |  |
| isu4 (2 cpu) | 53.6% / 100% | 16% | 21.4% | 41 | 115 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 86% / nginx 12% |
| isu2 | isureal 55% / nginx 1% |
| isu3 | isureal 54% / nginx 1% |
| isu4 | postgres 46% / isureal 13% / kswapd0 1% |

## Diff from previous run

スコア: 126129 → 141335 (+15206)

コードの差分は無い（同じ sha bd0fe45bf2e2）。
スコアの違いはブレか、サーバ側の設定変更による。
