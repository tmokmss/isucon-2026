# run 20261031-151629-200016f

## Summary

| | |
|---|---|
| score | 180317 |
| status | pass |
| ref | lane2-ipban |
| sha | 200016f64d91 |
| 所要時間 | 122s |
| 前回比 | +10019 (5.9%) vs 20261031-151136-dc74738 |
| ベスト比 | +6283 (3.6%) vs 20261031-144734-fd0f88e = 174034 **ベスト更新** |
| note | ip_bans bg reload only on 5a230a4 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	53	159	464	1258	1259
europe	7	20	63	187	432	432
america	7	18	54	159	221	221
```

```
{"id":"21720","status":"done","portalStatus":"finished","passed":true,"score":180317,"raw":180317,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:16:54.661Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:16:57.655Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:17:00.133Z level=INFO msg="user moved" user=nano_5098 from=europe to=america
time=2026-10-31T06:17:00.655Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T06:17:03.662Z level=INFO msg="asia 2 日目: 53 人"
time=2026-10-31T06:17:06.655Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T06:17:09.655Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T06:17:12.492Z level=INFO msg="user moved" user=nano_5098 from=america to=asia
time=2026-10-31T06:17:12.666Z level=INFO msg="asia 3 日目: 159 人"
time=2026-10-31T06:17:13.400Z level=INFO msg="user moved" user=isvin5621 from=asia to=europe
time=2026-10-31T06:17:14.767Z level=INFO msg="user moved" user=nano_5098 from=asia to=america
time=2026-10-31T06:17:15.655Z level=INFO msg="europe 3 日目: 63 人"
time=2026-10-31T06:17:16.261Z level=INFO msg="user moved" user=musemar-6226 from=asia to=europe
time=2026-10-31T06:17:18.656Z level=INFO msg="america 3 日目: 54 人"
time=2026-10-31T06:17:21.677Z level=INFO msg="asia 4 日目: 464 人"
time=2026-10-31T06:17:23.170Z level=INFO msg="user moved" user=rateru_6254 from=asia to=america
time=2026-10-31T06:17:25.239Z level=INFO msg="europe 4 日目: 187 人"
time=2026-10-31T06:17:25.338Z level=INFO msg="user moved" user=rateru_6254 from=america to=asia
time=2026-10-31T06:17:25.399Z level=INFO msg="user moved" user=marleono7598 from=asia to=america
time=2026-10-31T06:17:26.525Z level=INFO msg="user moved" user=marleono7598 from=america to=asia
time=2026-10-31T06:17:26.676Z level=INFO msg="user moved" user=ivyelna-8131 from=asia to=europe
time=2026-10-31T06:17:26.733Z level=INFO msg="user moved" user=nois7521 from=asia to=europe
time=2026-10-31T06:17:27.413Z level=INFO msg="user moved" user=hisuis.8025 from=asia to=america
time=2026-10-31T06:17:27.458Z level=INFO msg="user moved" user=toshiku.8074 from=asia to=europe
time=2026-10-31T06:17:27.884Z level=INFO msg="america 4 日目: 159 人"
time=2026-10-31T06:17:28.536Z level=INFO msg="user moved" user=anma.8490 from=asia to=europe
time=2026-10-31T06:17:28.652Z level=INFO msg="user moved" user=nuevachi.8851 from=asia to=europe
time=2026-10-31T06:17:29.544Z level=INFO msg="user moved" user=torkamo7888 from=asia to=europe
time=2026-10-31T06:17:30.696Z level=INFO msg="asia 5 日目: 1258 人"
time=2026-10-31T06:17:31.257Z level=INFO msg="user moved" user=nano_5098 from=america to=asia
time=2026-10-31T06:17:34.488Z level=INFO msg="europe 5 日目: 432 人"
time=2026-10-31T06:17:38.532Z level=INFO msg="america 5 日目: 221 人"
time=2026-10-31T06:17:46.592Z level=INFO msg="america 最終日（グランプリ）: 221 人"
time=2026-10-31T06:17:48.908Z level=INFO msg="asia 最終日（グランプリ）: 1259 人"
time=2026-10-31T06:17:49.411Z level=INFO msg="europe 最終日（グランプリ）: 432 人"
time=2026-10-31T06:17:54.736Z level=INFO msg="グランプリ: アジア 356 人, ヨーロッパ 69 人, アメリカ 141 人"
{
  "pass": true,
  "score": 180317,
  "messages": [
    "hetor-17729: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tsuvin-17224: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yuorto.17076: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kirihi.6585: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "furi.16857: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "torkimia17994: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nota_17630: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chisochi-8182: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "suheso_10850: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yominu_17836: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hezoe17890: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mialiashi.14672: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.41.143:55782-\u003e52.69.202.109:8080: write tcp 10.27.41.143:55782-\u003e52.69.202.109:8080: use of closed network connection",
    "sebensa_18206: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "wanoa_18132: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "meyoel18067: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yuyoka-18239: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "komu_17946: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.0%	1957	0	0	GET	/api/notifications	18977.934	9.697	18.196	18.211	166
13.0%	7055	93	0	POST	/api/friend-requests	5728.580	0.812	3.163	4.792	114
9.5%	3087	244	0	POST	/api/posts	4176.952	1.353	4.801	5.721	600
6.1%	7118	574	0	POST	/api/auth/login	2676.108	0.376	2.553	5.156	226
4.2%	20276	347	0	GET	^/api/images/[^/]+	1854.376	0.091	0.902	3.192	133241
3.7%	2965	6	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1644.562	0.555	2.735	4.513	116
2.4%	4663	0	0	GET	/api/memories	1060.790	0.227	3.061	6.564	12449
2.2%	16772	0	0	GET	^/api/feeds/[^/]+	971.915	0.058	0.384	0.982	3043
2.1%	857	132	0	PUT	/api/me/profile-image	942.818	1.100	7.443	9.535	3
1.9%	1283	1	0	POST	/api/auth/register	827.430	0.645	1.824	2.471	227
1.8%	773	12	0	POST	^/api/posts/[^/]+/realmojis$	807.270	1.044	4.408	5.623	476
1.5%	779	6	0	POST	^/api/posts/[^/]+/comments$	647.402	0.831	4.571	7.193	340
1.4%	813	0	0	PUT	^/api/follows/[^/]+$	622.068	0.765	2.872	3.513	0
1.1%	622	0	0	PATCH	/api/me	477.243	0.767	2.872	4.186	277
1.0%	1092	0	0	POST	/api/auth/logout	439.795	0.403	2.602	3.414	0
0.7%	4519	0	0	GET	/api/posting-window/current	289.649	0.064	0.405	0.777	235
0.6%	3409	0	0	GET	/api/users	265.138	0.078	0.366	0.638	246
0.6%	439	0	0	GET	/api/me/export	260.959	0.594	2.434	8.169	1869402
0.4%	271	11	0	PATCH	^/api/posts/[^/]+$	192.365	0.710	3.010	4.148	622
0.4%	2308	74	0	GET	^/api/posts/me/current$	161.862	0.070	0.406	0.796	630
0.3%	740	3	0	GET	/api/me	152.168	0.206	0.713	0.795	302
0.3%	124	0	0	DELETE	^/api/posts/[^/]+$	149.647	1.207	3.670	3.880	0
0.3%	142	0	0	DELETE	^/api/realmojis/[^/]+$	139.724	0.984	3.855	4.062	0
0.3%	2577	39	0	GET	^/api/posts/[^/]+$	133.982	0.052	0.366	0.701	607
0.2%	1984	0	0	GET	^/api/relationships/[^/]+	102.171	0.051	0.374	0.701	885
0.2%	1042	0	0	GET	^/api/memories/[^/]+	95.123	0.091	0.773	1.623	616
0.2%	182	0	0	DELETE	^/api/comments/[^/]+$	89.978	0.494	2.388	3.116	0
0.1%	1197	0	0	GET	^/api/users/[^/]+$	63.209	0.053	0.369	0.684	188
0.1%	929	0	0	GET	^/api/posts/[^/]+/realmojis$	53.464	0.058	0.383	0.669	2076
0.1%	55	0	0	PUT	^/api/blocks/[^/]+$	44.533	0.810	3.556	3.556	0
…（上位 30 / 全 43 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.3%	660	0	0	GET	/api/notifications	5390.124	8.167	14.950	15.944	160
14.3%	2108	22	0	POST	/api/friend-requests	1862.833	0.884	3.356	4.795	114
8.2%	905	69	0	POST	/api/posts	1066.682	1.179	4.525	5.129	600
6.0%	1208	0	0	POST	/api/auth/login	777.947	0.644	2.745	3.557	245
3.8%	3804	29	0	GET	^/api/images/[^/]+	501.387	0.132	1.020	2.435	178958
3.8%	3939	0	0	GET	^/api/feeds/[^/]+	501.116	0.127	0.935	1.869	2713
3.1%	423	0	0	POST	/api/auth/register	404.016	0.955	2.776	2.851	229
3.0%	508	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	387.999	0.764	3.063	3.803	116
2.5%	244	0	0	PUT	/api/me/profile-image	328.301	1.345	3.994	5.771	0
1.9%	245	0	0	PUT	^/api/follows/[^/]+$	249.335	1.018	3.168	3.542	0
1.8%	193	0	0	PATCH	/api/me	234.761	1.216	3.494	4.969	279
1.4%	1386	0	0	GET	/api/posting-window/current	183.988	0.133	1.305	1.897	239
1.4%	926	0	0	GET	/api/memories	179.468	0.194	1.370	2.203	13038
1.0%	123	5	0	POST	^/api/posts/[^/]+/realmojis$	129.637	1.054	3.967	4.404	444
0.9%	1009	0	0	GET	/api/users	119.115	0.118	0.739	1.390	245
0.8%	633	35	0	GET	^/api/posts/me/current$	101.294	0.160	1.371	1.689	614
0.7%	180	0	0	POST	/api/auth/logout	97.634	0.542	3.020	3.556	0
0.7%	96	0	0	POST	^/api/posts/[^/]+/comments$	92.642	0.965	4.805	4.805	367
0.6%	481	0	0	GET	^/api/relationships/[^/]+	83.710	0.174	0.954	1.471	836
0.5%	566	18	0	GET	^/api/posts/[^/]+$	65.587	0.116	0.614	0.930	595
0.4%	64	7	0	PATCH	^/api/posts/[^/]+$	51.790	0.809	2.516	2.516	587
0.4%	31	0	0	DELETE	^/api/posts/[^/]+$	48.915	1.578	2.927	2.927	0
0.2%	261	0	0	GET	^/api/users/[^/]+$	29.613	0.113	0.668	1.192	189
0.2%	252	1	0	GET	/api/me	29.243	0.116	0.740	0.741	302
0.2%	17	0	0	DELETE	^/api/realmojis/[^/]+$	22.725	1.337	3.298	3.298	0
0.2%	191	0	0	GET	^/api/posts/[^/]+/realmojis$	20.579	0.108	0.580	0.740	1448
0.1%	164	0	0	GET	^/api/memories/[^/]+	18.251	0.111	0.910	1.372	813
0.1%	162	0	0	GET	^/api/posts/[^/]+/comments$	18.104	0.112	0.577	0.891	1975
0.1%	17	0	0	DELETE	^/api/comments/[^/]+$	12.347	0.726	2.208	2.208	0
0.1%	56	0	0	GET	/api/me/export	10.872	0.194	1.680	1.680	2529916
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.7%	399	0	0	GET	/api/notifications	2925.962	7.333	10.942	10.942	166
12.9%	1168	14	0	POST	/api/friend-requests	925.739	0.793	2.814	3.483	114
9.3%	570	13	0	POST	/api/posts	667.493	1.171	4.190	5.479	638
7.8%	884	0	0	POST	/api/auth/login	560.544	0.634	3.358	5.136	250
5.4%	2983	18	0	GET	^/api/images/[^/]+	391.600	0.131	1.146	3.445	146754
3.6%	1966	0	0	GET	^/api/feeds/[^/]+	256.397	0.130	0.852	1.165	2751
3.2%	313	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	230.143	0.735	2.710	3.459	116
2.1%	219	0	0	POST	/api/auth/register	151.272	0.691	2.823	2.939	229
2.0%	811	0	0	GET	/api/memories	140.566	0.173	0.962	1.601	15069
1.5%	114	0	0	PUT	/api/me/profile-image	104.273	0.915	3.038	3.241	0
1.4%	109	0	0	PATCH	/api/me	103.685	0.951	3.522	4.073	278
1.4%	150	0	0	POST	/api/auth/logout	98.663	0.658	2.714	4.824	0
1.2%	806	0	0	GET	/api/posting-window/current	88.741	0.110	0.585	0.762	240
1.0%	112	0	0	PUT	^/api/follows/[^/]+$	75.157	0.671	2.654	2.688	0
1.0%	582	0	0	GET	/api/users	72.876	0.125	0.929	1.204	255
1.0%	71	3	0	POST	^/api/posts/[^/]+/realmojis$	71.417	1.006	3.672	3.672	438
0.9%	497	23	0	GET	^/api/posts/me/current$	61.451	0.124	0.775	0.954	624
0.6%	53	8	0	PATCH	^/api/posts/[^/]+$	42.151	0.795	2.683	2.683	549
0.6%	63	3	0	POST	^/api/posts/[^/]+/comments$	40.690	0.646	2.290	2.290	335
0.6%	265	0	0	GET	^/api/relationships/[^/]+	40.512	0.153	1.054	1.114	895
0.4%	19	0	0	DELETE	^/api/posts/[^/]+$	31.567	1.661	3.949	3.949	0
0.2%	175	8	0	GET	^/api/posts/[^/]+$	17.551	0.100	0.289	0.973	594
0.2%	133	0	0	GET	^/api/users/[^/]+$	16.741	0.126	1.067	1.101	192
0.2%	132	0	0	GET	^/api/memories/[^/]+	13.172	0.100	0.340	0.664	1002
0.2%	117	0	0	GET	/api/me	11.960	0.102	0.533	0.555	305
0.2%	52	0	0	GET	/api/me/export	11.391	0.219	1.848	1.848	2480895
0.1%	11	0	0	DELETE	^/api/comments/[^/]+$	10.650	0.968	2.267	2.267	0
0.1%	72	0	0	GET	^/api/posts/[^/]+/realmojis$	7.141	0.099	0.646	0.646	1555
0.1%	4	0	0	DELETE	^/api/realmojis/[^/]+$	6.346	1.587	3.091	3.091	0
0.1%	55	0	0	GET	^/api/posts/[^/]+/comments$	4.749	0.086	0.249	0.249	2239
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 17759 クエリ / 48 種類 / 合計 163.100s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	57.7%	447	94.184	210.70	4042.74	1190.6	121.6	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	13.1%	3	21.384	7127.92	9535.10	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	10.9%	14010	17.734	1.27	10.76	9.3	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	9.8%	3	15.917	5305.70	7190.84	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.2%	3	8.446	2815.44	4178.34	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.0%	4	3.224	806.10	2710.80	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.2%	4	0.374	93.45	157.70	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.2%	4	0.343	85.86	157.50	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.251	251.44	251.44	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	0.2%	1	0.248	248.40	248.40	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.1%	1	0.225	225.37	225.37	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.166	166.40	166.40	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.104	103.99	103.99	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.0%	1	0.081	80.54	80.54	2488.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.071	71.07	71.07	1551.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
Build ID: 632cd8db6d3a0ab64d707c84cd74e78578f77f3c
Type: cpu
Time: 2026-10-31 15:16:34 JST
Duration: 60.06s, Total samples = 46.69s (77.74%)
Showing nodes accounting for 32.82s, 70.29% of 46.69s total
Dropped 1164 nodes (cum <= 0.23s)
Showing top 40 nodes out of 292
      flat  flat%   sum%        cum   cum%
     5.69s 12.19% 12.19%      5.69s 12.19%  internal/runtime/syscall/linux.Syscall6
     4.88s 10.45% 22.64%      4.88s 10.45%  crypto/internal/fips140/sha256.blockAVX2
     2.27s  4.86% 27.50%      2.27s  4.86%  runtime.memmove
     2.19s  4.69% 32.19%      2.69s  5.76%  runtime.findObject
     1.89s  4.05% 36.24%      1.89s  4.05%  runtime.memclrNoHeapPointers
     1.83s  3.92% 40.16%      7.34s 15.72%  runtime.scanObject
     1.65s  3.53% 43.69%      5.70s 12.21%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.54s  3.30% 46.99%      1.54s  3.30%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.50s  3.21% 50.20%      2.67s  5.72%  runtime.tryDeferToSpanScan
     1.28s  2.74% 52.94%      1.28s  2.74%  aeshashbody
     0.84s  1.80% 54.74%      0.84s  1.80%  runtime.spanClass.sizeclass (inline)
     0.61s  1.31% 56.05%      2.37s  5.08%  runtime.mapaccess1_faststr
     0.54s  1.16% 57.21%      0.54s  1.16%  memeqbody
     0.45s  0.96% 58.17%      0.45s  0.96%  runtime.nextFreeFast (inline)
     0.38s  0.81% 58.98%      0.58s  1.24%  example.com/isureal-company/isureal.postLess
     0.32s  0.69% 59.67%      0.32s  0.69%  runtime.futex
     0.32s  0.69% 60.36%      0.32s  0.69%  sync/atomic.(*Int32).Add
     0.31s  0.66% 61.02%      0.78s  1.67%  database/sql.convertAssignRows
     0.29s  0.62% 61.64%      2.12s  4.54%  runtime.mapassign_faststr
     0.29s  0.62% 62.26%      0.30s  0.64%  runtime.typePointers.next
     0.28s   0.6% 62.86%      5.47s 11.72%  crypto/internal/fips140/sha256.(*Digest).Write
     0.26s  0.56% 63.42%      1.51s  3.23%  runtime.mallocgcSmallScanNoHeader
     0.25s  0.54% 63.95%      5.57s 11.93%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.25s  0.54% 64.49%      0.30s  0.64%  encoding/json.appendString[go.shape.string]
     0.23s  0.49% 64.98%      0.50s  1.07%  compress/flate.(*compressor).deflate
     0.23s  0.49% 65.47%      5.96s 12.77%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.23s  0.49% 65.97%      0.26s  0.56%  runtime.spanOf (inline)
     0.22s  0.47% 66.44%      0.49s  1.05%  runtime.greyobject
     0.21s  0.45% 66.89%      0.25s  0.54%  runtime.(*mspan).writeHeapBitsSmall
     0.20s  0.43% 67.32%      5.52s 11.82%  runtime.mallocgc
     0.19s  0.41% 67.72%      6.37s 13.64%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.18s  0.39% 68.11%      0.37s  0.79%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.16s  0.34% 68.45%      0.27s  0.58%  internal/runtime/maps.(*Iter).Next
     0.14s   0.3% 68.75%      3.12s  6.68%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.14s   0.3% 69.05%      0.55s  1.18%  runtime.mapaccess2_faststr
     0.13s  0.28% 69.33%      0.45s  0.96%  runtime.scanObjectsSmall
     0.12s  0.26% 69.59%      0.86s  1.84%  encoding/json.structEncoder.encode
     0.12s  0.26% 69.84%      0.71s  1.52%  runtime.wbBufFlush1
     0.11s  0.24% 70.08%      0.81s  1.73%  runtime.scanSpan
     0.10s  0.21% 70.29%      0.32s  0.69%  crypto/internal/fips140/hmac.(*HMAC).Reset
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 45.1% / 100% | 27% | 0.9% | 5 | 108 MB | run queue が積もっている |
| isu2 (2 cpu) | 29.0% / 100% | 4% | 0.7% | 4 | 370 MB |  |
| isu3 (2 cpu) | 25.2% / 99% | 5% | 0.6% | 4 | 299 MB |  |
| isu4 (2 cpu) | 42.6% / 100% | 18% | 16.1% | 8 | 112 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 70% / nginx 14% |
| isu2 | isureal 54% / nginx 2% |
| isu3 | isureal 48% / nginx 1% |
| isu4 | isureal 35% / postgres 16% |

## Diff from previous run

スコア: 170298 → 180317 (+10019)

`dc74738..200016f`

```
 conf/isu1/etc/systemd/system/isureal-go.service | 2 +-
 webapp/go/image_cache.go                        | 2 +-
 2 files changed, 2 insertions(+), 2 deletions(-)
```
