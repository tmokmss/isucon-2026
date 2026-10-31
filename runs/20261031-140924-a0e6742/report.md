# run 20261031-140924-a0e6742

## Summary

| | |
|---|---|
| score | 127403 |
| status | pass |
| ref | lane2-tune |
| sha | a0e674282db2 |
| 所要時間 | 325s |
| 前回比 | +14878 (13.2%) vs 20261031-135832-1338f93 |
| ベスト比 | +8086 (6.8%) vs 20261031-135233-2d10aad = 119317 **ベスト更新** |
| note | reset: VACUUM (TRUNCATE false) image_objects on c6abd65 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	142	418	927	926
europe	7	19	54	152	155	155
america	6	17	52	132	135	135
```

```
{"id":"16290","status":"done","portalStatus":"finished","passed":true,"score":127403,"raw":127403,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:13:20.969Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:13:23.941Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:13:26.941Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:13:29.963Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T05:13:32.941Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T05:13:35.940Z level=INFO msg="america 2 日目: 17 人"
time=2026-10-31T05:13:38.960Z level=INFO msg="asia 3 日目: 142 人"
time=2026-10-31T05:13:39.905Z level=INFO msg="user moved" user=hatsuka_5629 from=asia to=america
time=2026-10-31T05:13:40.331Z level=INFO msg="user moved" user=iszoeel-5424 from=europe to=asia
time=2026-10-31T05:13:40.631Z level=INFO msg="user moved" user=iszoeel-5424 from=asia to=america
time=2026-10-31T05:13:41.948Z level=INFO msg="europe 3 日目: 54 人"
time=2026-10-31T05:13:44.644Z level=INFO msg="user moved" user=hatsuka_5629 from=america to=asia
time=2026-10-31T05:13:44.947Z level=INFO msg="america 3 日目: 52 人"
time=2026-10-31T05:13:47.978Z level=INFO msg="asia 4 日目: 418 人"
time=2026-10-31T05:13:48.906Z level=INFO msg="user moved" user=mardanfu6430 from=america to=europe
time=2026-10-31T05:13:51.045Z level=INFO msg="europe 4 日目: 152 人"
time=2026-10-31T05:13:53.184Z level=INFO msg="user moved" user=zoesolia7863 from=asia to=europe
time=2026-10-31T05:13:53.372Z level=INFO msg="user moved" user=yonene.8580 from=asia to=america
time=2026-10-31T05:13:53.596Z level=INFO msg="user moved" user=noan_7090 from=asia to=europe
time=2026-10-31T05:13:53.980Z level=INFO msg="america 4 日目: 132 人"
time=2026-10-31T05:13:54.418Z level=INFO msg="user moved" user=anhi.7020 from=asia to=europe
time=2026-10-31T05:13:55.526Z level=INFO msg="user moved" user=mardanfu6430 from=europe to=america
time=2026-10-31T05:13:56.293Z level=INFO msg="user moved" user=sara_7785 from=asia to=america
time=2026-10-31T05:13:56.832Z level=INFO msg="user moved" user=nihita.7937 from=asia to=europe
time=2026-10-31T05:13:57.002Z level=INFO msg="user moved" user=hatsuka_5629 from=asia to=america
time=2026-10-31T05:13:57.144Z level=INFO msg="asia 5 日目: 927 人"
time=2026-10-31T05:14:00.061Z level=INFO msg="europe 5 日目: 155 人"
time=2026-10-31T05:14:03.081Z level=INFO msg="america 5 日目: 135 人"
time=2026-10-31T05:14:12.173Z level=INFO msg="asia 最終日（グランプリ）: 926 人"
time=2026-10-31T05:14:12.363Z level=INFO msg="europe 最終日（グランプリ）: 155 人"
time=2026-10-31T05:14:12.433Z level=INFO msg="america 最終日（グランプリ）: 135 人"
time=2026-10-31T05:14:20.979Z level=INFO msg="グランプリ: アジア 634 人, ヨーロッパ 98 人, アメリカ 63 人"
{
  "pass": true,
  "score": 127403,
  "messages": [
    "anchite5018: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "solmia5087: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "rihohi-12553: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "leomu9491: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nari-5826: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "museru13372: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tasol-11145: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "totorsol_6037: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ivymevin-10854: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "orhoke_12223: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ketafu.9891: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "redanchi-13169: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "leoevaso_12883: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "namechi_8762: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "meyo_11081: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yasu-10703: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "makosa.12257: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.3%	1479	1	0	GET	/api/notifications	12180.549	8.236	15.135	15.140	159
16.3%	2244	50	0	POST	/api/posts	5768.925	2.571	9.665	13.250	639
10.6%	5061	85	0	POST	/api/friend-requests	3754.937	0.742	4.625	10.586	114
8.7%	14111	599	0	GET	^/api/images/[^/]+	3090.239	0.219	3.188	12.031	144976
6.0%	5388	1206	0	POST	/api/auth/login	2141.957	0.398	4.841	12.930	195
5.4%	702	130	0	PUT	/api/me/profile-image	1927.495	2.746	10.316	13.044	4
4.6%	2118	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1628.492	0.769	4.223	7.707	116
2.6%	950	1	0	POST	/api/auth/register	908.700	0.957	6.376	11.089	227
2.5%	500	36	0	POST	^/api/posts/[^/]+/realmojis$	895.978	1.792	10.083	13.101	448
1.3%	639	0	0	PUT	^/api/follows/[^/]+$	460.932	0.721	4.457	10.468	0
1.3%	3619	0	0	GET	/api/memories	458.107	0.127	0.806	1.202	11035
1.2%	484	13	0	POST	^/api/posts/[^/]+/comments$	418.077	0.864	6.799	11.953	332
0.9%	425	0	0	PATCH	/api/me	336.878	0.793	5.426	7.804	280
0.7%	275	0	0	GET	/api/me/export	260.751	0.948	3.984	5.332	2033878
0.7%	654	0	0	POST	/api/auth/logout	251.461	0.384	4.886	9.041	0
0.6%	12597	0	0	GET	^/api/feeds/[^/]+	213.049	0.017	0.130	1.371	2038
0.6%	234	11	0	PATCH	^/api/posts/[^/]+$	195.252	0.834	4.520	6.089	627
0.4%	101	0	0	DELETE	^/api/posts/[^/]+$	138.386	1.370	4.801	5.371	0
0.2%	3253	0	0	GET	/api/posting-window/current	75.911	0.023	0.188	0.291	237
0.1%	1969	106	0	GET	^/api/posts/me/current$	49.392	0.025	0.189	1.952	618
0.1%	2392	0	0	GET	/api/users	43.828	0.018	0.149	0.346	240
0.1%	51	0	0	DELETE	^/api/realmojis/[^/]+$	37.991	0.745	3.764	3.764	0
0.1%	567	0	0	GET	^/api/memories/[^/]+	31.243	0.055	0.679	0.839	659
0.1%	43	0	0	PUT	^/api/blocks/[^/]+$	30.053	0.699	4.454	4.454	0
0.1%	1688	0	0	GET	^/api/relationships/[^/]+	29.486	0.017	0.141	0.294	991
0.1%	1901	9	0	GET	^/api/posts/[^/]+$	28.862	0.015	0.099	0.270	623
0.1%	74	0	0	DELETE	^/api/comments/[^/]+$	25.554	0.345	3.146	3.146	0
0.1%	1	0	0	POST	/api/initialize	19.185	19.185	19.185	19.185	90
0.0%	956	0	0	GET	^/api/users/[^/]+$	14.808	0.015	0.100	0.214	206
0.0%	846	0	0	GET	^/api/posts/[^/]+/realmojis$	11.872	0.014	0.082	0.124	935
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.7%	375	1	0	GET	/api/notifications	3118.568	8.316	12.347	12.434	175
12.5%	478	1	0	POST	/api/posts	1005.898	2.104	7.471	8.996	651
9.1%	926	16	0	POST	/api/friend-requests	730.228	0.789	5.160	9.113	113
7.4%	1988	18	0	GET	^/api/images/[^/]+	592.644	0.298	2.837	4.230	197259
5.1%	2067	0	0	GET	^/api/feeds/[^/]+	413.989	0.200	0.788	0.981	2185
5.1%	562	0	0	POST	/api/auth/login	413.776	0.736	5.592	8.048	247
4.1%	290	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	331.712	1.144	6.305	6.639	116
3.1%	597	0	0	GET	/api/memories	250.800	0.420	1.530	1.602	10307
2.1%	661	0	0	GET	/api/posting-window/current	172.337	0.261	0.785	1.033	239
1.5%	422	24	0	GET	^/api/posts/me/current$	121.208	0.287	0.876	1.001	615
1.3%	88	7	0	POST	^/api/posts/[^/]+/realmojis$	106.993	1.216	8.555	8.555	419
1.1%	428	0	0	GET	/api/users	86.385	0.202	0.671	4.757	239
1.1%	112	3	0	POST	^/api/posts/[^/]+/comments$	86.140	0.769	7.556	7.810	339
0.8%	150	0	0	POST	/api/auth/register	68.332	0.456	1.675	1.807	228
0.8%	35	0	0	GET	/api/me/export	66.740	1.907	6.750	6.750	1992498
0.7%	97	0	0	POST	/api/auth/logout	60.283	0.621	6.468	6.468	0
0.7%	82	0	0	PUT	/api/me/profile-image	59.408	0.724	10.054	10.054	0
0.7%	247	0	0	GET	^/api/relationships/[^/]+	53.046	0.215	0.777	0.840	913
0.5%	244	12	0	GET	^/api/posts/[^/]+$	42.012	0.172	0.677	0.708	586
0.5%	85	0	0	PATCH	/api/me	40.318	0.474	2.035	2.035	271
0.5%	94	0	0	PUT	^/api/follows/[^/]+$	40.296	0.429	2.334	2.334	0
0.5%	32	2	0	PATCH	^/api/posts/[^/]+$	37.572	1.174	4.245	4.245	601
0.4%	19	0	0	DELETE	^/api/posts/[^/]+$	29.488	1.552	5.574	5.574	0
0.3%	134	0	0	GET	^/api/users/[^/]+$	23.958	0.179	1.016	1.707	216
0.3%	128	1	0	GET	^/api/posts/[^/]+/realmojis$	20.186	0.158	0.609	0.898	814
0.2%	120	0	0	GET	^/api/posts/[^/]+/comments$	20.129	0.168	0.430	0.461	1417
0.2%	83	0	0	GET	^/api/memories/[^/]+	19.569	0.236	0.996	0.996	817
0.2%	19	0	0	DELETE	^/api/comments/[^/]+$	14.669	0.772	5.142	5.142	0
0.2%	87	2	0	GET	/api/me	13.652	0.157	0.614	0.614	296
0.2%	12	0	0	DELETE	^/api/realmojis/[^/]+$	12.260	1.022	3.036	3.036	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.9%	305	0	0	GET	/api/notifications	2212.772	7.255	9.390	9.495	175
12.8%	398	10	0	POST	/api/posts	811.379	2.039	8.927	10.789	634
11.1%	705	18	0	POST	/api/friend-requests	700.760	0.994	5.398	8.077	113
8.1%	1592	17	0	GET	^/api/images/[^/]+	515.915	0.324	2.580	3.814	184805
5.3%	444	0	0	POST	/api/auth/login	337.704	0.761	5.982	8.113	249
4.2%	1179	0	0	GET	^/api/feeds/[^/]+	265.122	0.225	0.847	0.971	1895
3.9%	235	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	245.047	1.043	5.054	7.389	116
3.0%	473	0	0	GET	/api/memories	188.847	0.399	1.352	1.555	13996
2.4%	556	0	0	GET	/api/posting-window/current	152.098	0.274	0.750	0.823	240
1.5%	336	31	0	GET	^/api/posts/me/current$	93.413	0.278	0.814	0.935	589
1.4%	102	0	0	PUT	^/api/follows/[^/]+$	90.076	0.883	6.584	8.536	0
1.3%	28	0	0	GET	/api/me/export	83.778	2.992	8.975	8.975	3312103
1.2%	74	0	0	PUT	/api/me/profile-image	73.720	0.996	10.355	10.355	0
1.1%	335	0	0	GET	/api/users	72.189	0.215	0.785	0.854	243
1.0%	35	6	0	POST	^/api/posts/[^/]+/realmojis$	66.394	1.897	7.016	7.016	367
1.0%	68	3	0	POST	^/api/posts/[^/]+/comments$	65.927	0.970	9.755	9.755	314
0.9%	75	0	0	POST	/api/auth/logout	59.305	0.791	4.806	4.806	0
0.8%	130	0	0	POST	/api/auth/register	50.108	0.385	1.210	2.319	229
0.7%	191	0	0	GET	^/api/relationships/[^/]+	45.349	0.237	0.855	0.925	878
0.6%	20	0	0	DELETE	^/api/posts/[^/]+$	38.865	1.943	5.594	5.594	0
0.6%	56	0	0	PATCH	/api/me	36.283	0.648	5.464	5.464	284
0.5%	169	14	0	GET	^/api/posts/[^/]+$	32.613	0.193	0.768	0.789	569
0.4%	30	6	0	PATCH	^/api/posts/[^/]+$	24.702	0.823	2.986	2.986	540
0.3%	78	0	0	GET	/api/me	16.840	0.216	1.045	1.045	301
0.2%	56	0	0	GET	^/api/memories/[^/]+	12.882	0.230	1.504	1.504	979
0.2%	76	0	0	GET	^/api/users/[^/]+$	12.566	0.165	0.530	0.530	197
0.1%	60	1	0	GET	^/api/posts/[^/]+/comments$	9.423	0.157	0.538	0.538	2182
0.1%	48	0	0	GET	^/api/posts/[^/]+/realmojis$	9.359	0.195	0.625	0.625	1114
0.1%	10	0	0	DELETE	^/api/comments/[^/]+$	7.393	0.739	2.835	2.835	0
0.1%	5	0	0	PUT	^/api/blocks/[^/]+$	4.885	0.977	2.981	2.981	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 74528 クエリ / 83 種類 / 合計 265.981s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	48.1%	338	127.983	378.65	1928.77	1343.9	136.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	16.7%	3046	44.292	14.54	1414.68	101.2	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
3	15.5%	4363	41.239	9.45	171.54	571.9	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
4	9.2%	9223	24.527	2.66	55.32	11.0	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	2.5%	3061	6.710	2.19	16.35	25.0	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	2.4%	2	6.429	3214.45	6270.16	211833.5	2831.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.8%	1	2.162	2162.39	2162.39	31317.0	1199932.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.7%	574	1.889	3.29	16.93	28.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.6%	1	1.715	1715.18	1715.18	24949.0	628106.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
10	0.6%	6550	1.558	0.24	15.42	19.8	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.4%	645	1.129	1.75	17.29	21.7	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.3%	2580	0.871	0.34	49.09	20.7	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
13	0.3%	1	0.787	787.11	787.11	12969.0	530225.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
14	0.3%	1	0.768	767.69	767.69	30940.0	2995.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.2%	5188	0.646	0.12	9.46	11.2	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

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
Build ID: 87adc59d8bba99fee2d50b2481cc0a8797897af7
Type: cpu
Time: 2026-10-31 14:12:57 JST
Duration: 60s, Total samples = 39.39s (65.65%)
Showing nodes accounting for 28.96s, 73.52% of 39.39s total
Dropped 962 nodes (cum <= 0.20s)
Showing top 40 nodes out of 287
      flat  flat%   sum%        cum   cum%
     3.73s  9.47%  9.47%      3.73s  9.47%  internal/runtime/syscall/linux.Syscall6
     3.43s  8.71% 18.18%      3.44s  8.73%  crypto/internal/fips140/sha256.blockAVX2
     1.72s  4.37% 22.54%      2.21s  5.61%  runtime.findObject
     1.58s  4.01% 26.55%      5.51s 13.99%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.56s  3.96% 30.52%      6.30s 15.99%  runtime.scanObject
     1.53s  3.88% 34.40%      1.53s  3.88%  runtime.memmove
     1.52s  3.86% 38.26%      1.52s  3.86%  runtime.memclrNoHeapPointers
     1.50s  3.81% 42.07%      1.50s  3.81%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.29s  3.27% 45.34%      2.15s  5.46%  runtime.tryDeferToSpanScan
     1.19s  3.02% 48.36%      1.19s  3.02%  aeshashbody
     1.06s  2.69% 51.05%      1.64s  4.16%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.91s  2.31% 53.36%      3.37s  8.56%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.79s  2.01% 55.37%      0.82s  2.08%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.57s  1.45% 56.82%      0.57s  1.45%  runtime.spanClass.sizeclass (inline)
     0.49s  1.24% 58.06%      2.08s  5.28%  runtime.mapaccess1_faststr
     0.42s  1.07% 59.13%      0.61s  1.55%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.40s  1.02% 60.14%      0.40s  1.02%  runtime.futex
     0.39s  0.99% 61.13%      0.58s  1.47%  example.com/isureal-company/isureal.postLess
     0.38s  0.96% 62.10%      0.38s  0.96%  memeqbody
     0.31s  0.79% 62.88%      0.56s  1.42%  runtime.greyobject
     0.30s  0.76% 63.65%      0.31s  0.79%  runtime.(*mspan).writeHeapBitsSmall
     0.28s  0.71% 64.36%      4.24s 10.76%  runtime.mallocgc
     0.27s  0.69% 65.04%      0.71s  1.80%  database/sql.convertAssignRows
     0.27s  0.69% 65.73%      1.14s  2.89%  runtime.mallocgcSmallScanNoHeader
     0.27s  0.69% 66.41%      0.27s  0.69%  runtime.typePointers.nextFast (inline)
     0.26s  0.66% 67.07%      0.26s  0.66%  runtime.nextFreeFast (inline)
     0.24s  0.61% 67.68%      0.27s  0.69%  encoding/json.appendString[go.shape.string]
     0.24s  0.61% 68.29%      0.26s  0.66%  runtime.typePointers.next
     0.23s  0.58% 68.88%      3.13s  7.95%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.23s  0.58% 69.46%      0.23s  0.58%  indexbody
     0.21s  0.53% 69.99%      0.43s  1.09%  compress/flate.(*compressor).deflate
     0.19s  0.48% 70.47%      3.75s  9.52%  crypto/internal/fips140/sha256.(*Digest).Write
     0.19s  0.48% 70.96%      4.69s 11.91%  database/sql.withLock
     0.17s  0.43% 71.39%      0.21s  0.53%  runtime.spanOf (inline)
     0.16s  0.41% 71.79%      0.23s  0.58%  runtime.heapArenaOf (inline)
     0.16s  0.41% 72.20%      1.70s  4.32%  runtime.mapassign_faststr
     0.15s  0.38% 72.58%      0.23s  0.58%  runtime.(*mspan).markBitsForIndex (inline)
     0.15s  0.38% 72.96%      0.52s  1.32%  runtime.scanObjectsSmall
     0.12s   0.3% 73.27%      0.71s  1.80%  runtime.mallocgcSmallNoscan
     0.10s  0.25% 73.52%      1.22s  3.10%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 45.2% / 100% | 30% | 0.3% | 6 | 178 MB | run queue が積もっている |
| isu2 (2 cpu) | 20.6% / 100% | 1% | 0.6% | 4 | 1901 MB |  |
| isu3 (2 cpu) | 18.4% / 100% | 2% | 0.7% | 6 | 1644 MB | run queue が積もっている |
| isu4 (2 cpu) | 52.3% / 100% | 21% | 21.4% | 23 | 117 MB | IO 待ちが重い |

**isu1（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 74% / nginx 12% |
| isu2 | isureal 39% / nginx 1% |
| isu3 | isureal 34% / nginx 1% |
| isu4 | postgres 38% / isureal 13% / kswapd0 1% |

## Diff from previous run

スコア: 112525 → 127403 (+14878)

`1338f93..a0e6742`

```
 conf/isu1/etc/systemd/system/isureal-go.service        | 2 +-
 conf/isu4/etc/postgresql/16/main/conf.d/zz-tuning.conf | 2 +-
 webapp/go/server.go                                    | 5 ++---
 webapp/sql/reset.sql                                   | 4 ++++
 4 files changed, 8 insertions(+), 5 deletions(-)
```
