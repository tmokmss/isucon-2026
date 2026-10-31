# run 20261031-145831-da3beca

## Summary

| | |
|---|---|
| score | 172831 |
| status | pass |
| ref | da3beca |
| sha | da3beca0229f |
| 所要時間 | 128s |
| 前回比 | +36787 (27.0%) vs 20261031-144830-3c8fc33 |
| ベスト比 | -1203 (-0.7%) vs 20261031-144734-fd0f88e = 174034 |
| note | 画像をリージョン内で完結（image-local、main 6ad2e36 の上） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	151	439	1073	1073
europe	7	16	46	130	316	315
america	7	18	58	159	227	226
```

```
{"id":"20155","status":"done","portalStatus":"finished","passed":true,"score":172831,"raw":172831,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:58:58.927Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:59:01.918Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:59:03.936Z level=INFO msg="user moved" user=kehemo_5085 from=europe to=america
time=2026-10-31T05:59:04.915Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T05:59:07.944Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T05:59:10.916Z level=INFO msg="europe 2 日目: 16 人"
time=2026-10-31T05:59:13.915Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T05:59:14.636Z level=INFO msg="user moved" user=ruto5428 from=europe to=asia
time=2026-10-31T05:59:15.464Z level=INFO msg="user moved" user=kehemo_5085 from=america to=europe
time=2026-10-31T05:59:16.927Z level=INFO msg="asia 3 日目: 151 人"
time=2026-10-31T05:59:17.601Z level=INFO msg="user moved" user=tsure.5541 from=asia to=europe
time=2026-10-31T05:59:18.184Z level=INFO msg="user moved" user=ruto5428 from=asia to=america
time=2026-10-31T05:59:18.644Z level=INFO msg="user moved" user=tsure.5541 from=europe to=asia
time=2026-10-31T05:59:19.012Z level=INFO msg="user moved" user=mate-5819 from=asia to=america
time=2026-10-31T05:59:19.920Z level=INFO msg="europe 3 日目: 46 人"
time=2026-10-31T05:59:21.355Z level=INFO msg="user moved" user=kehemo_5085 from=europe to=america
time=2026-10-31T05:59:22.916Z level=INFO msg="america 3 日目: 58 人"
time=2026-10-31T05:59:24.658Z level=INFO msg="user moved" user=chimaxro_6356 from=europe to=asia
time=2026-10-31T05:59:25.289Z level=INFO msg="user moved" user=waul-6455 from=europe to=america
time=2026-10-31T05:59:25.940Z level=INFO msg="asia 4 日目: 439 人"
time=2026-10-31T05:59:27.645Z level=INFO msg="user moved" user=chimaxro_6356 from=asia to=america
time=2026-10-31T05:59:29.012Z level=INFO msg="europe 4 日目: 130 人"
time=2026-10-31T05:59:29.172Z level=INFO msg="user moved" user=tsure.5541 from=asia to=europe
time=2026-10-31T05:59:29.976Z level=INFO msg="user moved" user=miaki-6842 from=asia to=america
time=2026-10-31T05:59:30.648Z level=INFO msg="user moved" user=rachidan_7974 from=asia to=europe
time=2026-10-31T05:59:31.574Z level=INFO msg="user moved" user=maxfu_8093 from=asia to=europe
time=2026-10-31T05:59:32.132Z level=INFO msg="user moved" user=ririho.7648 from=asia to=america
time=2026-10-31T05:59:32.331Z level=INFO msg="america 4 日目: 159 人"
time=2026-10-31T05:59:33.003Z level=INFO msg="user moved" user=suhi_8427 from=asia to=europe
time=2026-10-31T05:59:33.957Z level=INFO msg="user moved" user=ulru.8738 from=asia to=america
time=2026-10-31T05:59:34.360Z level=INFO msg="user moved" user=yuno-8864 from=europe to=america
time=2026-10-31T05:59:34.453Z level=INFO msg="user moved" user=vinevaso8315 from=asia to=america
time=2026-10-31T05:59:34.799Z level=INFO msg="user moved" user=kehemo_5085 from=america to=europe
time=2026-10-31T05:59:35.131Z level=INFO msg="asia 5 日目: 1073 人"
time=2026-10-31T05:59:35.245Z level=INFO msg="user moved" user=elyona-8209 from=asia to=europe
time=2026-10-31T05:59:35.310Z level=INFO msg="user moved" user=waul-6455 from=america to=europe
time=2026-10-31T05:59:37.831Z level=INFO msg="user moved" user=ruto5428 from=america to=asia
time=2026-10-31T05:59:38.630Z level=INFO msg="europe 5 日目: 316 人"
time=2026-10-31T05:59:42.085Z level=INFO msg="america 5 日目: 227 人"
time=2026-10-31T05:59:50.990Z level=INFO msg="asia 最終日（グランプリ）: 1073 人"
time=2026-10-31T05:59:51.289Z level=INFO msg="america 最終日（グランプリ）: 226 人"
time=2026-10-31T05:59:51.367Z level=INFO msg="europe 最終日（グランプリ）: 315 人"
time=2026-10-31T05:59:58.997Z level=INFO msg="グランプリ: アジア 504 人, ヨーロッパ 184 人, アメリカ 79 人"
{
  "pass": true,
  "score": 172831,
  "messages": [
    "nomoso-15784: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tsumoki-15144: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "anfu_15902: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "solliaben-15382: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hose-14489: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nuhozoe.9301: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "washiha_15789: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "soka.16092: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chite15918: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hoor.15572: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "elni13450: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "makuyo_5058: 友達申請を承認できません: POST /api/friend-requests/04dd7d28-f24d-4335-b998-b8a16c2f5f1f/accept: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests/04dd7d28-f24d-4335-b998-b8a16c2f5f1f/accept\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.9%	1768	0	0	GET	/api/notifications	16719.734	9.457	15.978	15.993	171
14.7%	2780	153	0	POST	/api/posts	5715.546	2.056	8.473	10.885	617
8.4%	6340	72	0	POST	/api/friend-requests	3252.768	0.513	2.140	3.335	114
6.1%	20839	326	0	GET	^/api/images/[^/]+	2357.131	0.113	1.391	3.819	125895
3.7%	7054	603	0	POST	/api/auth/login	1429.240	0.203	1.717	4.630	226
3.4%	755	130	0	PUT	/api/me/profile-image	1309.463	1.734	9.511	12.846	3
3.3%	5021	0	0	GET	/api/memories	1266.766	0.252	2.520	6.697	12898
3.1%	3182	10	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1206.268	0.379	2.100	2.919	116
2.7%	15793	0	0	GET	^/api/feeds/[^/]+	1042.129	0.066	0.495	0.990	3478
2.4%	664	25	0	POST	^/api/posts/[^/]+/realmojis$	929.419	1.400	8.566	8.687	463
1.2%	813	14	0	POST	^/api/posts/[^/]+/comments$	459.712	0.565	3.286	6.819	338
0.8%	730	0	0	PUT	^/api/follows/[^/]+$	312.240	0.428	2.073	2.900	0
0.8%	3033	0	0	GET	/api/users	294.435	0.097	0.584	0.916	240
0.8%	3932	0	0	GET	/api/posting-window/current	293.209	0.075	0.557	0.720	237
0.7%	411	0	0	GET	/api/me/export	283.273	0.689	3.635	7.425	1950747
0.7%	1098	1	0	POST	/api/auth/register	281.990	0.257	0.769	1.320	227
0.7%	1082	0	0	POST	/api/auth/logout	262.707	0.243	1.929	2.141	0
0.5%	524	0	0	PATCH	/api/me	214.100	0.409	1.689	1.939	273
0.4%	2209	39	0	GET	^/api/posts/me/current$	169.945	0.077	0.491	0.754	641
0.4%	2533	42	0	GET	^/api/posts/[^/]+$	151.250	0.060	0.443	0.602	610
0.3%	2015	0	0	GET	^/api/relationships/[^/]+	134.791	0.067	0.595	0.762	1008
0.3%	234	3	0	PATCH	^/api/posts/[^/]+$	131.090	0.560	2.258	2.496	648
0.3%	639	2	0	GET	/api/me	121.032	0.189	0.847	1.190	300
0.3%	1005	0	0	GET	^/api/memories/[^/]+	110.449	0.110	0.944	1.246	625
0.2%	92	0	0	DELETE	^/api/posts/[^/]+$	96.887	1.053	2.674	2.674	0
0.2%	126	0	0	DELETE	^/api/realmojis/[^/]+$	80.644	0.640	2.054	3.061	0
0.2%	170	0	0	DELETE	^/api/comments/[^/]+$	79.066	0.465	2.042	2.308	0
0.2%	1174	0	0	GET	^/api/users/[^/]+$	61.018	0.052	0.333	0.538	198
0.1%	1023	0	0	GET	^/api/posts/[^/]+/realmojis$	52.866	0.052	0.488	0.861	1237
0.1%	927	0	0	GET	^/api/posts/[^/]+/comments$	50.977	0.055	0.396	0.774	1826
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.3%	524	0	0	GET	/api/notifications	3827.615	7.305	12.649	12.667	165
10.5%	1519	8	0	POST	/api/friend-requests	1023.785	0.674	2.257	3.080	114
8.1%	750	20	0	POST	/api/posts	792.896	1.057	6.085	6.691	633
6.9%	3156	21	0	GET	^/api/images/[^/]+	672.282	0.213	2.030	2.761	174086
5.8%	2725	0	0	GET	^/api/feeds/[^/]+	567.758	0.208	1.301	1.667	2373
5.1%	956	0	0	POST	/api/auth/login	499.612	0.523	2.114	2.406	245
2.9%	312	0	0	POST	/api/auth/register	284.668	0.912	2.183	2.289	229
2.6%	867	0	0	GET	/api/memories	254.155	0.293	1.643	2.541	11629
2.3%	171	0	0	PUT	/api/me/profile-image	219.407	1.283	3.734	5.845	0
2.2%	363	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	214.654	0.591	2.236	2.736	116
2.1%	1093	0	0	GET	/api/posting-window/current	201.009	0.184	1.248	1.696	238
2.1%	252	0	0	PUT	^/api/follows/[^/]+$	199.950	0.793	2.308	2.898	0
1.3%	143	9	0	POST	^/api/posts/[^/]+/realmojis$	128.828	0.901	4.116	4.167	454
1.3%	581	23	0	GET	^/api/posts/me/current$	126.066	0.217	1.330	2.042	626
1.1%	737	0	0	GET	/api/users	107.872	0.146	1.076	1.449	241
1.1%	143	0	0	PATCH	/api/me	105.530	0.738	2.307	3.476	271
1.0%	130	1	0	POST	^/api/posts/[^/]+/comments$	96.091	0.739	2.383	3.126	355
0.7%	362	0	0	GET	^/api/relationships/[^/]+	64.940	0.179	1.213	1.577	827
0.6%	127	0	0	POST	/api/auth/logout	57.516	0.453	2.148	2.164	0
0.5%	334	25	0	GET	^/api/posts/[^/]+$	50.044	0.150	1.000	1.377	571
0.4%	62	8	0	PATCH	^/api/posts/[^/]+$	43.382	0.700	2.273	2.273	569
0.3%	24	0	0	DELETE	^/api/posts/[^/]+$	25.036	1.043	2.400	2.400	0
0.2%	185	0	0	GET	^/api/users/[^/]+$	23.948	0.129	0.812	0.878	189
0.2%	133	0	0	GET	^/api/memories/[^/]+	23.583	0.177	1.136	1.387	876
0.2%	150	0	0	GET	^/api/posts/[^/]+/realmojis$	22.983	0.153	0.949	0.994	1419
0.2%	42	0	0	GET	/api/me/export	21.227	0.505	5.491	5.491	2138575
0.2%	176	2	0	GET	/api/me	18.473	0.105	0.387	0.564	302
0.2%	109	0	0	GET	^/api/posts/[^/]+/comments$	15.920	0.146	0.994	1.001	2665
0.2%	29	0	0	DELETE	^/api/comments/[^/]+$	15.191	0.524	1.279	1.279	0
0.2%	20	0	0	DELETE	^/api/realmojis/[^/]+$	14.887	0.744	1.741	1.741	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.7%	433	1	0	GET	/api/notifications	3179.873	7.344	10.058	10.332	158
10.9%	1226	12	0	POST	/api/friend-requests	853.308	0.696	2.266	2.973	114
8.0%	571	14	0	POST	/api/posts	626.512	1.097	6.353	6.584	633
7.6%	2835	17	0	GET	^/api/images/[^/]+	592.399	0.209	1.577	5.288	157491
5.8%	1989	0	0	GET	^/api/feeds/[^/]+	453.643	0.228	1.580	2.117	3376
5.7%	877	0	0	POST	/api/auth/login	446.905	0.510	2.312	3.988	250
2.7%	385	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	209.548	0.544	2.243	2.864	116
2.3%	730	0	0	GET	/api/memories	177.040	0.243	1.825	2.237	15780
2.1%	820	0	0	GET	/api/posting-window/current	166.625	0.203	1.105	1.721	240
1.8%	219	0	0	POST	/api/auth/register	139.640	0.638	2.388	2.557	230
1.7%	116	0	0	PUT	/api/me/profile-image	131.704	1.135	3.575	6.948	0
1.5%	619	0	0	GET	/api/users	115.074	0.186	0.904	1.372	237
1.3%	414	10	0	GET	^/api/posts/me/current$	103.075	0.249	1.753	2.057	631
1.0%	62	6	0	POST	^/api/posts/[^/]+/realmojis$	77.251	1.246	8.781	8.781	441
1.0%	149	0	0	POST	/api/auth/logout	76.668	0.515	2.102	2.112	0
0.9%	102	0	0	PATCH	/api/me	68.808	0.675	2.120	2.230	269
0.8%	99	0	0	PUT	^/api/follows/[^/]+$	60.346	0.610	2.185	2.185	0
0.6%	291	17	0	GET	^/api/posts/[^/]+$	46.501	0.160	0.852	0.961	580
0.6%	268	0	0	GET	^/api/relationships/[^/]+	43.791	0.163	1.537	1.830	1038
0.5%	64	1	0	POST	^/api/posts/[^/]+/comments$	41.197	0.644	3.179	3.179	338
0.5%	42	3	0	PATCH	^/api/posts/[^/]+$	35.297	0.840	4.069	4.069	584
0.3%	49	0	0	GET	/api/me/export	25.639	0.523	11.984	11.984	3456935
0.3%	130	0	0	GET	^/api/memories/[^/]+	22.465	0.173	1.060	1.847	1095
0.3%	122	1	0	GET	/api/me	21.917	0.180	0.931	0.962	300
0.3%	18	0	0	DELETE	^/api/posts/[^/]+$	20.318	1.129	2.670	2.670	0
0.2%	139	0	0	GET	^/api/users/[^/]+$	19.215	0.138	0.836	0.843	195
0.2%	87	0	0	GET	^/api/posts/[^/]+/realmojis$	13.487	0.155	0.850	0.850	1554
0.2%	12	0	0	DELETE	^/api/realmojis/[^/]+$	12.918	1.077	2.621	2.621	0
0.1%	85	0	0	GET	^/api/posts/[^/]+/comments$	9.353	0.110	0.623	0.623	2568
0.1%	16	0	0	DELETE	^/api/comments/[^/]+$	7.077	0.442	1.529	1.529	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16958 クエリ / 48 種類 / 合計 150.897s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	54.7%	423	82.499	195.03	1037.83	1243.5	127.2	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.1%	3	21.314	7104.51	9536.95	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.1%	14692	16.726	1.14	7.10	9.3	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.6%	3	15.998	5332.60	7265.39	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.2%	3	7.904	2634.74	3561.44	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	1.7%	6	2.638	439.62	1772.73	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.8%	1	1.239	1238.95	1238.95	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.4%	6	0.674	112.26	164.28	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.4%	6	0.670	111.58	163.94	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.2%	1	0.246	246.18	246.18	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.1%	1	0.221	221.26	221.26	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.150	149.68	149.68	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.114	114.16	114.16	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.084	83.56	83.56	2419.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.1%	1	0.077	77.29	77.29	1507.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
Build ID: 7814d789d714cb40461f04e20d24583875a8f4da
Type: cpu
Time: 2026-10-31 14:58:38 JST
Duration: 60.05s, Total samples = 44.19s (73.59%)
Showing nodes accounting for 30.48s, 68.97% of 44.19s total
Dropped 1189 nodes (cum <= 0.22s)
Showing top 40 nodes out of 282
      flat  flat%   sum%        cum   cum%
     5.30s 11.99% 11.99%      5.30s 11.99%  crypto/internal/fips140/sha256.blockAVX2
     5.09s 11.52% 23.51%      5.09s 11.52%  internal/runtime/syscall/linux.Syscall6
     1.82s  4.12% 27.63%      1.82s  4.12%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.82s  4.12% 31.75%      1.82s  4.12%  runtime.memmove
     1.73s  3.91% 35.66%      5.85s 13.24%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.70s  3.85% 39.51%      2.14s  4.84%  runtime.findObject
     1.70s  3.85% 43.36%      1.70s  3.85%  runtime.memclrNoHeapPointers
     1.64s  3.71% 47.07%      6.14s 13.89%  runtime.scanObject
     1.44s  3.26% 50.33%      2.44s  5.52%  runtime.tryDeferToSpanScan
     1.40s  3.17% 53.50%      1.40s  3.17%  aeshashbody
     0.80s  1.81% 55.31%      0.80s  1.81%  runtime.spanClass.sizeclass (inline)
     0.49s  1.11% 56.42%      0.49s  1.11%  memeqbody
     0.48s  1.09% 57.50%      0.48s  1.09%  runtime.nextFreeFast (inline)
     0.43s  0.97% 58.47%      0.43s  0.97%  runtime.futex
     0.35s  0.79% 59.27%      2.13s  4.82%  runtime.mapaccess1_faststr
     0.34s  0.77% 60.04%      6.04s 13.67%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.31s   0.7% 60.74%      0.55s  1.24%  example.com/isureal-company/isureal.postLess
     0.30s  0.68% 61.42%      4.64s 10.50%  runtime.mallocgc
     0.30s  0.68% 62.10%      2.30s  5.20%  runtime.mapassign_faststr
     0.24s  0.54% 62.64%      5.86s 13.26%  crypto/internal/fips140/sha256.(*Digest).Write
     0.24s  0.54% 63.18%      1.17s  2.65%  runtime.mallocgcSmallScanNoHeader
     0.23s  0.52% 63.70%      0.23s  0.52%  runtime.typePointers.nextFast (inline)
     0.22s   0.5% 64.20%      0.71s  1.61%  database/sql.convertAssignRows
     0.20s  0.45% 64.65%      0.48s  1.09%  compress/flate.(*compressor).deflate
     0.19s  0.43% 65.08%      6.31s 14.28%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.17s  0.38% 65.47%      0.30s  0.68%  internal/runtime/maps.(*Iter).Next
     0.15s  0.34% 65.81%      6.65s 15.05%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.15s  0.34% 66.15%      0.43s  0.97%  runtime.scanObjectsSmall
     0.14s  0.32% 66.46%      0.32s  0.72%  runtime.greyobject
     0.13s  0.29% 66.76%      1.15s  2.60%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.27% 67.03%         1s  2.26%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
     0.12s  0.27% 67.30%      0.98s  2.22%  runtime.mallocgcSmallNoscan
     0.11s  0.25% 67.55%      0.28s  0.63%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.10s  0.23% 67.78%      5.40s 12.22%  crypto/internal/fips140/sha256.block
     0.10s  0.23% 68.00%      3.65s  8.26%  example.com/isureal-company/isureal.listMemories
     0.09s   0.2% 68.21%      0.85s  1.92%  encoding/json.structEncoder.encode
     0.09s   0.2% 68.41%      2.98s  6.74%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.09s   0.2% 68.61%      0.49s  1.11%  runtime.convTstring
     0.08s  0.18% 68.79%      0.23s  0.52%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.08s  0.18% 68.97%      0.52s  1.18%  runtime.mapaccess2_faststr
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 47.1% / 100% | 32% | 0.6% | 6 | 133 MB | CPU が飽和 |
| isu2 (2 cpu) | 27.6% / 100% | 4% | 0.5% | 4 | 302 MB |  |
| isu3 (2 cpu) | 28.5% / 100% | 7% | 0.5% | 4 | 363 MB |  |
| isu4 (2 cpu) | 34.7% / 99% | 10% | 17.6% | 8 | 106 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 75% / nginx 13% |
| isu2 | isureal 53% / nginx 1% |
| isu3 | isureal 54% / nginx 2% |
| isu4 | isureal 18% / postgres 15% |

## Diff from previous run

スコア: 136044 → 172831 (+36787)

`3c8fc33..da3beca`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf |   2 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |   4 +
 conf/isu2/home/isucon/env.sh                   |   1 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |   4 +
 conf/isu3/home/isucon/env.sh                   |   1 +
 webapp/go/export_api.go                        |   4 +-
 webapp/go/image_file.go                        |  11 +-
 webapp/go/image_held.go                        | 172 +++++++++++++++++++++++++
 webapp/go/image_local.go                       |  20 +--
 webapp/go/origin_store.go                      |  59 +++++----
 webapp/go/post_api.go                          |  20 ++-
 webapp/go/post_forward.go                      |  14 +-
 webapp/go/repl.go                              |  17 ++-
 webapp/go/replica_server.go                    |  36 ++++--
 webapp/go/server.go                            |   1 +
 webapp/go/store.go                             |   1 +
 16 files changed, 306 insertions(+), 61 deletions(-)
```
