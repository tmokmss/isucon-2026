# run 20261031-144734-fd0f88e

## Summary

| | |
|---|---|
| score | 174034 |
| status | pass |
| ref | fd0f88e |
| sha | fd0f88e6dbc8 |
| 所要時間 | 117s |
| 前回比 | +38609 (28.5%) vs 20261031-143412-10c9074 |
| ベスト比 | +32699 (23.1%) vs 20261031-141027-bd0fe45 = 141335 **ベスト更新** |
| note | 画像をリージョン内で完結（image-local、inmemory-persist f7ca385 の上、repl-lag の計測入り） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	140	417	1074	1074
europe	7	19	51	147	391	390
america	7	16	48	136	245	245
```

```
{"id":"19234","status":"done","portalStatus":"finished","passed":true,"score":174034,"raw":174034,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:47:58.680Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:48:01.487Z level=INFO msg="user moved" user=rileoho-5014 from=asia to=america
time=2026-10-31T05:48:01.673Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:48:04.669Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T05:48:07.680Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T05:48:10.670Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T05:48:13.671Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T05:48:16.682Z level=INFO msg="asia 3 日目: 140 人"
time=2026-10-31T05:48:19.672Z level=INFO msg="europe 3 日目: 51 人"
time=2026-10-31T05:48:20.023Z level=INFO msg="user moved" user=solmaxta.6025 from=asia to=america
time=2026-10-31T05:48:21.838Z level=INFO msg="user moved" user=rileoho-5014 from=america to=asia
time=2026-10-31T05:48:22.671Z level=INFO msg="america 3 日目: 48 人"
time=2026-10-31T05:48:23.075Z level=INFO msg="user moved" user=rileoho-5014 from=asia to=america
time=2026-10-31T05:48:24.247Z level=INFO msg="user moved" user=nina-6196 from=europe to=asia
time=2026-10-31T05:48:25.683Z level=INFO msg="asia 4 日目: 417 人"
time=2026-10-31T05:48:28.729Z level=INFO msg="europe 4 日目: 147 人"
time=2026-10-31T05:48:31.342Z level=INFO msg="user moved" user=rima7652 from=asia to=america
time=2026-10-31T05:48:31.694Z level=INFO msg="user moved" user=nahomia7397 from=asia to=europe
time=2026-10-31T05:48:31.751Z level=INFO msg="america 4 日目: 136 人"
time=2026-10-31T05:48:33.736Z level=INFO msg="user moved" user=kavin_8574 from=asia to=europe
time=2026-10-31T05:48:34.693Z level=INFO msg="asia 5 日目: 1074 人"
time=2026-10-31T05:48:34.903Z level=INFO msg="user moved" user=mutor-8730 from=europe to=asia
time=2026-10-31T05:48:38.286Z level=INFO msg="europe 5 日目: 391 人"
time=2026-10-31T05:48:41.464Z level=INFO msg="america 5 日目: 245 人"
time=2026-10-31T05:48:50.831Z level=INFO msg="america 最終日（グランプリ）: 245 人"
time=2026-10-31T05:48:51.676Z level=INFO msg="europe 最終日（グランプリ）: 390 人"
time=2026-10-31T05:48:52.714Z level=INFO msg="asia 最終日（グランプリ）: 1074 人"
time=2026-10-31T05:48:58.740Z level=INFO msg="グランプリ: アジア 355 人, ヨーロッパ 188 人, アメリカ 159 人"
{
  "pass": true,
  "score": 174034,
  "messages": [
    "leobenna-5027: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "tsumaxhi-15579: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "eltorsu_15724: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ronesa14606: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "solsasu11783: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "elraya_16279: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "leoto-14782: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maseel-15925: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nusol.16451: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nehika_16594: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sotor_16676: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rezoe.16192: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nudan.16889: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maxhiyo12817: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "rikina15571: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
47.1%	1769	1	0	GET	/api/notifications	18404.852	10.404	18.086	18.128	168
12.9%	2728	152	0	POST	/api/posts	5040.234	1.848	5.571	5.905	620
8.2%	6053	74	0	POST	/api/friend-requests	3184.039	0.526	2.127	3.481	114
5.0%	19757	337	0	GET	^/api/images/[^/]+	1948.724	0.099	0.939	3.772	123891
4.2%	7041	676	0	POST	/api/auth/login	1637.817	0.233	1.836	3.111	223
3.4%	2808	9	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1327.363	0.473	2.042	2.804	116
3.2%	771	134	0	PUT	/api/me/profile-image	1244.999	1.615	7.240	10.306	3
2.7%	4576	0	0	GET	/api/memories	1050.484	0.230	2.077	8.713	12707
2.5%	15446	0	0	GET	^/api/feeds/[^/]+	979.024	0.063	0.372	0.859	2813
1.7%	598	27	0	POST	^/api/posts/[^/]+/realmojis$	666.886	1.115	4.188	5.932	456
1.0%	593	1	0	POST	^/api/posts/[^/]+/comments$	372.083	0.627	2.799	3.116	348
0.9%	1157	0	0	POST	/api/auth/logout	356.979	0.309	2.031	2.765	0
0.8%	700	0	0	PUT	^/api/follows/[^/]+$	308.867	0.441	1.752	1.997	0
0.7%	429	0	0	GET	/api/me/export	286.974	0.669	4.369	7.944	2046210
0.7%	1094	1	0	POST	/api/auth/register	277.804	0.254	1.524	1.561	227
0.7%	554	0	0	PATCH	/api/me	268.623	0.485	1.998	2.410	281
0.6%	2907	0	0	GET	/api/users	245.606	0.084	0.400	0.738	244
0.6%	3931	0	0	GET	/api/posting-window/current	229.956	0.058	0.338	0.843	237
0.4%	2047	18	0	GET	^/api/posts/me/current$	153.688	0.075	0.402	0.670	651
0.4%	2485	33	0	GET	^/api/posts/[^/]+$	143.506	0.058	0.394	0.845	627
0.3%	234	3	0	PATCH	^/api/posts/[^/]+$	129.945	0.555	2.540	3.178	644
0.3%	1049	0	0	GET	^/api/memories/[^/]+	108.157	0.103	0.845	1.210	645
0.3%	1932	0	0	GET	^/api/relationships/[^/]+	101.064	0.052	0.329	0.689	898
0.3%	95	0	0	DELETE	^/api/posts/[^/]+$	99.483	1.047	2.559	2.559	0
0.2%	652	3	0	GET	/api/me	90.502	0.139	0.560	0.974	301
0.2%	127	0	0	DELETE	^/api/realmojis/[^/]+$	85.620	0.674	2.206	2.252	0
0.2%	142	0	0	DELETE	^/api/comments/[^/]+$	84.958	0.598	2.037	2.809	0
0.2%	1165	0	0	GET	^/api/users/[^/]+$	62.737	0.054	0.460	0.851	206
0.1%	888	0	0	GET	^/api/posts/[^/]+/realmojis$	52.228	0.059	0.332	0.497	1599
0.1%	749	0	0	GET	^/api/posts/[^/]+/comments$	46.446	0.062	0.386	0.591	2196
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.7%	587	1	0	GET	/api/notifications	4437.321	7.559	13.202	13.305	152
12.5%	1999	21	0	POST	/api/friend-requests	1475.281	0.738	2.413	4.569	114
6.8%	871	19	0	POST	/api/posts	801.544	0.920	4.304	5.207	638
6.4%	3692	25	0	GET	^/api/images/[^/]+	751.149	0.203	1.512	6.208	185704
6.0%	3704	0	0	GET	^/api/feeds/[^/]+	706.279	0.191	1.330	1.893	2983
5.8%	1145	0	0	POST	/api/auth/login	680.094	0.594	2.200	2.965	245
3.0%	389	0	0	POST	/api/auth/register	347.927	0.894	2.057	2.211	229
2.8%	222	0	0	PUT	/api/me/profile-image	327.916	1.477	3.898	3.910	0
2.5%	941	0	0	GET	/api/memories	293.261	0.312	1.611	2.056	12176
2.3%	422	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	268.820	0.637	2.079	2.202	116
2.2%	1286	0	0	GET	/api/posting-window/current	256.291	0.199	0.984	1.436	239
1.6%	236	0	0	PUT	^/api/follows/[^/]+$	183.639	0.778	2.538	2.744	0
1.4%	987	0	0	GET	/api/users	160.425	0.163	1.095	1.481	246
1.3%	185	0	0	PATCH	/api/me	157.355	0.851	2.185	2.237	280
1.3%	145	17	0	POST	^/api/posts/[^/]+/realmojis$	156.439	1.079	3.514	3.792	416
1.1%	454	0	0	GET	^/api/relationships/[^/]+	131.382	0.289	1.500	1.934	630
1.1%	663	31	0	GET	^/api/posts/me/current$	126.455	0.191	1.327	1.568	623
0.7%	568	23	0	GET	^/api/posts/[^/]+$	84.132	0.148	0.879	1.522	602
0.6%	158	0	0	POST	/api/auth/logout	72.707	0.460	1.958	2.200	0
0.4%	83	0	0	POST	^/api/posts/[^/]+/comments$	52.265	0.630	2.905	2.905	360
0.4%	35	0	0	DELETE	^/api/posts/[^/]+$	48.097	1.374	3.763	3.763	0
0.3%	264	0	0	GET	^/api/users/[^/]+$	39.534	0.150	0.764	0.957	197
0.3%	230	2	0	GET	/api/me	35.590	0.155	0.695	0.956	301
0.3%	52	3	0	PATCH	^/api/posts/[^/]+$	34.577	0.665	2.002	2.002	630
0.3%	32	0	0	DELETE	^/api/realmojis/[^/]+$	32.239	1.007	2.238	2.238	0
0.2%	193	0	0	GET	^/api/posts/[^/]+/realmojis$	27.403	0.142	0.769	1.126	1799
0.2%	143	0	0	GET	^/api/memories/[^/]+	22.404	0.157	1.209	1.800	932
0.2%	144	1	0	GET	^/api/posts/[^/]+/comments$	20.532	0.143	0.786	0.825	2131
0.2%	22	0	0	DELETE	^/api/comments/[^/]+$	18.626	0.847	2.221	2.221	0
0.1%	46	0	0	GET	/api/me/export	13.526	0.294	4.193	4.193	2337295
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.6%	411	2	0	GET	/api/notifications	3012.562	7.330	9.705	10.100	154
11.0%	1217	14	0	POST	/api/friend-requests	836.122	0.687	2.211	3.057	114
7.2%	2821	17	0	GET	^/api/images/[^/]+	546.079	0.194	1.467	4.753	166046
6.6%	558	14	0	POST	/api/posts	498.929	0.894	4.529	5.103	636
5.9%	852	0	0	POST	/api/auth/login	446.761	0.524	2.172	3.531	249
5.7%	2259	0	0	GET	^/api/feeds/[^/]+	430.100	0.190	1.086	1.272	2965
3.0%	375	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	225.844	0.602	2.231	2.884	116
2.7%	750	0	0	GET	/api/memories	203.781	0.272	1.601	2.123	14160
2.4%	241	0	0	POST	/api/auth/register	182.923	0.759	2.195	2.229	229
1.9%	810	0	0	GET	/api/posting-window/current	143.556	0.177	0.923	1.131	240
1.8%	137	0	0	PUT	/api/me/profile-image	133.824	0.977	3.742	3.896	0
1.4%	116	8	0	POST	^/api/posts/[^/]+/realmojis$	107.156	0.924	4.572	5.250	411
1.4%	153	0	0	PUT	^/api/follows/[^/]+$	103.585	0.677	1.822	1.847	0
1.3%	580	0	0	GET	/api/users	98.172	0.169	1.048	1.249	244
1.3%	291	0	0	GET	^/api/relationships/[^/]+	95.754	0.329	1.188	1.249	1054
1.1%	117	0	0	PATCH	/api/me	85.892	0.734	2.610	3.288	285
1.1%	111	2	0	POST	^/api/posts/[^/]+/comments$	81.293	0.732	2.561	2.728	304
1.0%	482	18	0	GET	^/api/posts/me/current$	76.502	0.159	1.011	1.149	628
0.8%	120	0	0	POST	/api/auth/logout	61.659	0.514	2.088	2.175	0
0.7%	299	12	0	GET	^/api/posts/[^/]+$	50.351	0.168	1.005	1.125	603
0.4%	21	0	0	DELETE	^/api/posts/[^/]+$	27.002	1.286	2.137	2.137	0
0.3%	34	3	0	PATCH	^/api/posts/[^/]+$	23.295	0.685	2.152	2.152	594
0.3%	125	0	0	GET	^/api/posts/[^/]+/realmojis$	20.673	0.165	1.012	1.148	1644
0.3%	142	0	0	GET	/api/me	20.661	0.145	1.048	1.394	304
0.2%	129	0	0	GET	^/api/memories/[^/]+	18.557	0.144	1.794	1.903	1003
0.2%	117	0	0	GET	^/api/users/[^/]+$	14.712	0.126	0.744	0.973	203
0.2%	79	0	0	GET	^/api/posts/[^/]+/comments$	14.278	0.181	0.976	0.976	1581
0.2%	53	0	0	GET	/api/me/export	13.523	0.255	1.839	1.839	3291255
0.2%	17	0	0	DELETE	^/api/comments/[^/]+$	11.588	0.682	1.776	1.776	0
0.1%	16	0	0	DELETE	^/api/realmojis/[^/]+$	10.233	0.640	2.237	2.237	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16266 クエリ / 48 種類 / 合計 161.645s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	58.3%	437	94.313	215.82	986.28	1301.9	133.3	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	13.3%	3	21.478	7159.45	9737.41	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.1%	14136	17.981	1.27	7.21	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	9.9%	3	15.940	5313.21	7179.19	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.0%	3	8.078	2692.75	3723.11	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	0.7%	6	1.139	189.89	217.23	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	6	0.691	115.19	165.34	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.4%	6	0.678	113.02	164.40	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.1%	1	0.240	239.98	239.98	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.1%	1	0.200	200.04	200.04	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.1%	1	0.196	196.32	196.32	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.149	149.15	149.15	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.120	119.70	119.70	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.0%	1	0.072	71.51	71.51	2419.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.064	64.24	64.24	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: bbdd1fdf1fdf36ac5323dd26ea635ea0e9ffe2d6
Type: cpu
Time: 2026-10-31 14:47:40 JST
Duration: 60.30s, Total samples = 48.40s (80.27%)
Showing nodes accounting for 34.71s, 71.71% of 48.40s total
Dropped 1163 nodes (cum <= 0.24s)
Showing top 40 nodes out of 281
      flat  flat%   sum%        cum   cum%
     6.88s 14.21% 14.21%      6.88s 14.21%  crypto/internal/fips140/sha256.blockAVX2
     5.42s 11.20% 25.41%      5.42s 11.20%  internal/runtime/syscall/linux.Syscall6
     2.21s  4.57% 29.98%      2.21s  4.57%  runtime.memmove
     2.07s  4.28% 34.26%      2.75s  5.68%  runtime.findObject
     1.95s  4.03% 38.29%      3.17s  6.55%  runtime.tryDeferToSpanScan
     1.73s  3.57% 41.86%      7.65s 15.81%  runtime.scanObject
     1.55s  3.20% 45.06%      1.55s  3.20%  runtime.memclrNoHeapPointers
     1.52s  3.14% 48.20%      1.52s  3.14%  aeshashbody
     1.48s  3.06% 51.26%      5.64s 11.65%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.45s  3.00% 54.26%      1.45s  3.00%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.96s  1.98% 56.24%      0.96s  1.98%  runtime.spanClass.sizeclass (inline)
     0.67s  1.38% 57.62%      2.47s  5.10%  runtime.mapaccess1_faststr
     0.57s  1.18% 58.80%      7.85s 16.22%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.47s  0.97% 59.77%      0.47s  0.97%  memeqbody
     0.40s  0.83% 60.60%      0.40s  0.83%  runtime.nextFreeFast (inline)
     0.35s  0.72% 61.32%      0.56s  1.16%  example.com/isureal-company/isureal.postLess
     0.33s  0.68% 62.00%      0.34s   0.7%  runtime.spanOf (inline)
     0.31s  0.64% 62.64%      7.48s 15.45%  crypto/internal/fips140/sha256.(*Digest).Write
     0.29s   0.6% 63.24%      0.71s  1.47%  database/sql.convertAssignRows
     0.28s  0.58% 63.82%      0.28s  0.58%  sync/atomic.(*Int32).Add
     0.26s  0.54% 64.36%      0.26s  0.54%  runtime.typePointers.nextFast (inline)
     0.24s   0.5% 64.86%      0.51s  1.05%  compress/flate.(*compressor).deflate
     0.24s   0.5% 65.35%      8.24s 17.02%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.24s   0.5% 65.85%      2.12s  4.38%  runtime.mapassign_faststr
     0.23s  0.48% 66.32%      0.33s  0.68%  runtime.(*mspan).markBitsForIndex (inline)
     0.22s  0.45% 66.78%      0.61s  1.26%  runtime.greyobject
     0.22s  0.45% 67.23%      1.19s  2.46%  runtime.mallocgcSmallScanNoHeader
     0.22s  0.45% 67.69%      0.60s  1.24%  runtime.mapaccess2_faststr
     0.22s  0.45% 68.14%      0.28s  0.58%  runtime.typePointers.next
     0.21s  0.43% 68.57%      8.67s 17.91%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.20s  0.41% 68.99%      0.52s  1.07%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.20s  0.41% 69.40%      4.10s  8.47%  runtime.mallocgc
     0.17s  0.35% 69.75%      0.53s  1.10%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
     0.16s  0.33% 70.08%      0.56s  1.16%  runtime.scanObjectsSmall
     0.16s  0.33% 70.41%      1.09s  2.25%  runtime.wbBufFlush1
     0.15s  0.31% 70.72%      0.53s  1.10%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.13s  0.27% 70.99%      0.74s  1.53%  encoding/json.structEncoder.encode
     0.12s  0.25% 71.24%      2.71s  5.60%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.12s  0.25% 71.49%      0.35s  0.72%  internal/runtime/maps.(*Iter).Next
     0.11s  0.23% 71.71%      0.26s  0.54%  runtime.scanObjectSmall
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 46.8% / 100% | 28% | 0.7% | 5 | 123 MB | run queue が積もっている |
| isu2 (2 cpu) | 33.5% / 100% | 5% | 0.9% | 4 | 132 MB |  |
| isu3 (2 cpu) | 33.1% / 100% | 6% | 1.2% | 6 | 212 MB | run queue が積もっている |
| isu4 (2 cpu) | 37.6% / 99% | 10% | 21.2% | 9 | 103 MB | IO 待ちが重い |

**isu1（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 74% / nginx 14% |
| isu2 | isureal 63% / nginx 2% |
| isu3 | isureal 62% / nginx 2% |
| isu4 | postgres 16% / isureal 14% |

## Diff from previous run

スコア: 135425 → 174034 (+38609)

`10c9074..fd0f88e`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf |   2 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |   4 +
 conf/isu2/home/isucon/env.sh                   |   1 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |   4 +
 conf/isu3/home/isucon/env.sh                   |   1 +
 webapp/go/cmd/isureal/main.go                  |   2 +
 webapp/go/export_api.go                        |   4 +-
 webapp/go/image_file.go                        |  11 +-
 webapp/go/image_held.go                        | 172 +++++++++++++++++++++++++
 webapp/go/image_local.go                       |  20 +--
 webapp/go/origin_store.go                      |  59 +++++----
 webapp/go/persist.go                           |  18 ++-
 webapp/go/post_api.go                          |  20 ++-
 webapp/go/post_forward.go                      |  14 +-
 webapp/go/repl.go                              |  17 ++-
 webapp/go/replica_server.go                    |  36 ++++--
 webapp/go/server.go                            |   1 +
 webapp/go/store.go                             |   1 +
 18 files changed, 323 insertions(+), 64 deletions(-)
```
