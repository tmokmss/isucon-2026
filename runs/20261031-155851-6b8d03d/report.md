# run 20261031-155851-6b8d03d

## Summary

| | |
|---|---|
| score | 171210 |
| status | pass |
| ref | mem-headroom |
| sha | 6b8d03da9689 |
| 所要時間 | 344s |
| 前回比 | -16700 (-8.9%) vs 20261031-155347-d7fb2c5 |
| ベスト比 | -20606 (-10.7%) vs 20261031-151339-5a230a4 = 191816 |
| note | mem headroom + nginx buffers shrink on isu1 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	149	442	1147	1146
europe	7	21	68	194	448	446
america	6	18	53	153	254	254
```

```
{"id":"25906","status":"done","portalStatus":"finished","passed":true,"score":171210,"raw":171210,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:02:46.767Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:02:49.753Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:02:52.754Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:02:55.762Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T07:02:58.753Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T07:02:59.043Z level=INFO msg="user moved" user=nite.5370 from=asia to=europe
time=2026-10-31T07:03:01.755Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T07:03:04.556Z level=INFO msg="user moved" user=tokosol.5506 from=america to=europe
time=2026-10-31T07:03:04.765Z level=INFO msg="asia 3 日目: 149 人"
time=2026-10-31T07:03:06.581Z level=INFO msg="user moved" user=kakuya_5853 from=asia to=europe
time=2026-10-31T07:03:07.752Z level=INFO msg="europe 3 日目: 68 人"
time=2026-10-31T07:03:10.267Z level=INFO msg="user moved" user=tokosol.5506 from=europe to=asia
time=2026-10-31T07:03:10.754Z level=INFO msg="america 3 日目: 53 人"
time=2026-10-31T07:03:12.893Z level=INFO msg="user moved" user=nite.5370 from=europe to=asia
time=2026-10-31T07:03:13.777Z level=INFO msg="asia 4 日目: 442 人"
time=2026-10-31T07:03:17.199Z level=INFO msg="europe 4 日目: 194 人"
time=2026-10-31T07:03:17.317Z level=INFO msg="user moved" user=tokosol.5506 from=asia to=europe
time=2026-10-31T07:03:18.712Z level=INFO msg="user moved" user=sonuke7614 from=asia to=america
time=2026-10-31T07:03:18.943Z level=INFO msg="user moved" user=tabenel_8038 from=asia to=europe
time=2026-10-31T07:03:18.952Z level=INFO msg="user moved" user=anke_7496 from=asia to=europe
time=2026-10-31T07:03:19.208Z level=INFO msg="user moved" user=benzoe_8147 from=asia to=europe
time=2026-10-31T07:03:19.389Z level=INFO msg="user moved" user=nite.5370 from=asia to=europe
time=2026-10-31T07:03:19.658Z level=INFO msg="user moved" user=maro_6724 from=america to=europe
time=2026-10-31T07:03:19.823Z level=INFO msg="user moved" user=sonuke7614 from=america to=asia
time=2026-10-31T07:03:19.875Z level=INFO msg="america 4 日目: 153 人"
time=2026-10-31T07:03:19.899Z level=INFO msg="user moved" user=nuben8876 from=asia to=europe
time=2026-10-31T07:03:20.520Z level=INFO msg="user moved" user=maliaben8687 from=asia to=america
time=2026-10-31T07:03:21.725Z level=INFO msg="user moved" user=evaben_9108 from=europe to=america
time=2026-10-31T07:03:21.916Z level=INFO msg="user moved" user=kenoachi_7549 from=asia to=europe
time=2026-10-31T07:03:22.795Z level=INFO msg="asia 5 日目: 1147 人"
time=2026-10-31T07:03:22.932Z level=INFO msg="user moved" user=lialiasu.8604 from=asia to=europe
time=2026-10-31T07:03:27.151Z level=INFO msg="europe 5 日目: 448 人"
time=2026-10-31T07:03:27.157Z level=INFO msg="user moved" user=kakuya_5853 from=europe to=america
time=2026-10-31T07:03:27.181Z level=INFO msg="user moved" user=anmiana9725 from=america to=asia
time=2026-10-31T07:03:30.390Z level=INFO msg="america 5 日目: 254 人"
time=2026-10-31T07:03:38.554Z level=INFO msg="asia 最終日（グランプリ）: 1146 人"
time=2026-10-31T07:03:39.118Z level=INFO msg="europe 最終日（グランプリ）: 446 人"
time=2026-10-31T07:03:39.135Z level=INFO msg="america 最終日（グランプリ）: 254 人"
time=2026-10-31T07:03:46.870Z level=INFO msg="グランプリ: アジア 444 人, ヨーロッパ 105 人, アメリカ 98 人"
{
  "pass": true,
  "score": 171210,
  "messages": [
    "roreni_14276: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "suke.14264: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tehasol_14927: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "anrahe17104: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "none17643: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "miamaxro-16189: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "numar-16656: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nomoeva.13566: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nosol-17310: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "wachimia15971: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "naansa_16253: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "isruan_16045: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "molia_6142: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mileo-17565: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maben-11060: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maxivywa.15682: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yuleochi16754: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "robense_10615: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hemia.11213: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "royo_14825: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.5%	1804	0	0	GET	/api/notifications	16592.456	9.198	15.813	15.916	169
11.4%	2878	224	0	POST	/api/posts	4568.486	1.587	6.599	7.322	602
10.2%	6379	72	0	POST	/api/friend-requests	4068.015	0.638	2.560	3.894	114
7.6%	19260	388	0	GET	^/api/images/[^/]+	3042.113	0.158	1.705	3.508	119825
3.9%	6895	669	0	POST	/api/auth/login	1563.565	0.227	2.239	3.700	222
3.2%	2801	9	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1281.100	0.457	2.443	3.351	116
3.1%	12897	0	0	GET	^/api/feeds/[^/]+	1247.038	0.097	0.757	1.745	2945
2.9%	4536	0	0	GET	/api/memories	1167.294	0.257	2.645	10.552	12689
2.0%	791	133	0	PUT	/api/me/profile-image	803.037	1.015	9.654	11.167	3
1.7%	1171	1	0	POST	/api/auth/register	671.250	0.573	2.002	2.281	227
1.6%	620	27	0	POST	^/api/posts/[^/]+/realmojis$	634.082	1.023	4.588	5.089	455
1.1%	632	7	0	POST	^/api/posts/[^/]+/comments$	456.950	0.723	3.848	4.780	352
1.1%	4098	0	0	GET	/api/posting-window/current	427.286	0.104	0.774	1.449	237
1.1%	751	0	0	PUT	^/api/follows/[^/]+$	421.900	0.562	2.310	3.216	0
1.0%	3105	0	0	GET	/api/users	388.578	0.125	0.748	1.313	250
0.9%	559	0	0	PATCH	/api/me	346.007	0.619	2.326	3.203	279
0.8%	400	0	0	GET	/api/me/export	338.170	0.845	7.358	12.044	1781087
0.7%	1032	0	0	POST	/api/auth/logout	269.930	0.262	2.204	3.164	0
0.7%	2192	62	0	GET	^/api/posts/me/current$	265.949	0.121	0.855	1.597	634
0.6%	668	3	0	GET	/api/me	230.802	0.346	1.655	1.872	301
0.4%	1643	0	0	GET	^/api/relationships/[^/]+	174.981	0.107	0.898	1.927	1080
0.4%	1788	34	0	GET	^/api/posts/[^/]+$	169.415	0.095	0.759	1.431	609
0.4%	131	0	0	DELETE	^/api/posts/[^/]+$	155.809	1.189	2.887	3.755	0
0.4%	236	10	0	PATCH	^/api/posts/[^/]+$	148.510	0.629	2.363	2.524	627
0.3%	1039	0	0	GET	^/api/memories/[^/]+	112.329	0.108	1.035	2.492	601
0.2%	791	2	0	GET	^/api/posts/[^/]+/realmojis$	77.057	0.097	0.730	1.175	1469
0.2%	962	0	0	GET	^/api/users/[^/]+$	76.232	0.079	0.509	1.150	199
0.2%	106	0	0	DELETE	^/api/realmojis/[^/]+$	70.177	0.662	3.184	4.150	0
0.1%	724	0	0	GET	^/api/posts/[^/]+/comments$	57.230	0.079	0.497	1.093	1845
0.1%	118	0	0	DELETE	^/api/comments/[^/]+$	56.598	0.480	2.389	2.917	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.7%	682	1	0	GET	/api/notifications	5278.847	7.740	11.941	12.344	151
11.5%	2028	19	0	POST	/api/friend-requests	1491.449	0.735	2.901	3.713	114
10.6%	964	80	0	POST	/api/posts	1379.396	1.431	6.623	65.441	597
5.5%	1277	0	0	POST	/api/auth/login	716.716	0.561	2.394	6.024	245
5.2%	4004	21	0	GET	^/api/images/[^/]+	675.867	0.169	1.553	3.226	161146
4.5%	3569	0	0	GET	^/api/feeds/[^/]+	582.113	0.163	1.672	3.054	2741
2.8%	437	0	0	POST	/api/auth/register	367.893	0.842	2.237	2.288	228
2.5%	251	0	0	PUT	/api/me/profile-image	318.085	1.267	3.550	3.822	0
2.2%	438	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	280.916	0.641	2.709	3.216	116
1.9%	264	0	0	PUT	^/api/follows/[^/]+$	245.143	0.929	2.853	3.483	0
1.6%	204	0	0	PATCH	/api/me	212.536	1.042	2.896	3.716	280
1.6%	1433	0	0	GET	/api/posting-window/current	202.982	0.142	1.489	2.383	239
1.5%	179	10	0	POST	^/api/posts/[^/]+/realmojis$	193.143	1.079	4.842	4.878	457
1.3%	1031	0	0	GET	/api/memories	162.563	0.158	1.866	3.282	12433
1.2%	634	29	0	GET	^/api/posts/me/current$	153.598	0.242	1.812	2.274	621
0.9%	1030	0	0	GET	/api/users	121.907	0.118	0.750	2.152	253
0.9%	145	1	0	POST	^/api/posts/[^/]+/comments$	120.605	0.832	3.741	3.919	347
0.7%	179	0	0	POST	/api/auth/logout	88.764	0.496	3.219	3.364	0
0.5%	506	23	0	GET	^/api/posts/[^/]+$	68.165	0.135	1.376	2.234	590
0.4%	451	0	0	GET	^/api/relationships/[^/]+	57.471	0.127	1.570	2.060	752
0.4%	65	4	0	PATCH	^/api/posts/[^/]+$	56.613	0.871	2.746	2.746	605
0.3%	236	0	0	GET	^/api/users/[^/]+$	39.632	0.168	1.348	1.380	198
0.2%	23	0	0	DELETE	^/api/posts/[^/]+$	31.449	1.367	2.714	2.714	0
0.2%	257	2	0	GET	/api/me	24.922	0.097	0.214	0.725	302
0.2%	203	0	0	GET	^/api/posts/[^/]+/realmojis$	24.403	0.120	0.918	1.371	1506
0.1%	23	0	0	DELETE	^/api/realmojis/[^/]+$	19.105	0.831	2.875	2.875	0
0.1%	134	0	0	GET	^/api/posts/[^/]+/comments$	16.523	0.123	1.494	1.521	1689
0.1%	155	0	0	GET	^/api/memories/[^/]+	16.076	0.104	0.593	1.055	772
0.1%	70	0	0	GET	/api/me/export	11.202	0.160	3.255	3.255	2452441
0.1%	28	0	0	DELETE	^/api/comments/[^/]+$	8.118	0.290	0.904	0.904	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.0%	439	1	0	GET	/api/notifications	3261.212	7.429	10.710	10.743	151
9.5%	1206	11	0	POST	/api/friend-requests	735.996	0.610	2.540	3.369	114
8.8%	590	24	0	POST	/api/posts	685.902	1.163	3.197	6.940	631
7.6%	960	0	0	POST	/api/auth/login	591.445	0.616	2.556	3.607	250
6.3%	3053	27	0	GET	^/api/images/[^/]+	489.804	0.160	1.521	3.486	133920
4.3%	2100	0	0	GET	^/api/feeds/[^/]+	335.842	0.160	1.152	1.672	2609
2.7%	253	0	0	POST	/api/auth/register	208.413	0.824	2.423	2.501	229
2.5%	290	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	196.543	0.678	3.062	3.282	115
1.8%	143	0	0	PUT	/api/me/profile-image	140.172	0.980	3.408	3.568	0
1.7%	859	0	0	GET	/api/posting-window/current	133.412	0.155	1.093	1.557	240
1.5%	761	0	0	GET	/api/memories	114.385	0.150	1.286	2.886	17863
1.3%	89	4	0	POST	^/api/posts/[^/]+/realmojis$	100.570	1.130	3.655	3.655	474
1.3%	447	14	0	GET	^/api/posts/me/current$	99.197	0.222	1.380	1.480	638
1.1%	131	0	0	POST	/api/auth/logout	83.374	0.636	2.670	3.189	0
1.0%	103	1	0	POST	^/api/posts/[^/]+/comments$	81.027	0.787	3.341	3.821	350
1.0%	145	0	0	PUT	^/api/follows/[^/]+$	78.368	0.540	2.670	3.449	0
1.0%	613	0	0	GET	/api/users	77.058	0.126	0.965	1.379	251
0.9%	116	0	0	PATCH	/api/me	71.699	0.618	2.543	2.815	270
0.6%	256	0	0	GET	^/api/relationships/[^/]+	43.599	0.170	1.177	1.398	809
0.5%	41	1	0	PATCH	^/api/posts/[^/]+$	40.971	0.999	4.337	4.337	646
0.5%	28	0	0	DELETE	^/api/posts/[^/]+$	38.448	1.373	2.864	2.864	0
0.3%	221	18	0	GET	^/api/posts/[^/]+$	24.662	0.112	0.717	0.950	568
0.2%	59	0	0	GET	/api/me/export	18.902	0.320	3.319	3.319	2972850
0.2%	143	0	0	GET	^/api/memories/[^/]+	17.734	0.124	1.335	1.335	749
0.2%	128	0	0	GET	^/api/users/[^/]+$	16.921	0.132	0.966	1.347	190
0.2%	149	1	0	GET	/api/me	15.843	0.106	0.272	1.552	302
0.2%	12	0	0	DELETE	^/api/realmojis/[^/]+$	13.551	1.129	1.999	1.999	0
0.1%	73	0	0	GET	^/api/posts/[^/]+/realmojis$	9.008	0.123	1.353	1.353	1574
0.1%	78	0	0	GET	^/api/posts/[^/]+/comments$	7.317	0.094	0.235	0.235	2139
0.1%	12	0	0	DELETE	^/api/comments/[^/]+$	6.070	0.506	0.968	0.968	0
…（上位 30 / 全 36 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16633 クエリ / 48 種類 / 合計 147.394s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	53.3%	408	78.547	192.52	1450.54	1131.8	115.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.9%	3	21.899	7299.73	10011.79	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.7%	13684	17.274	1.26	8.83	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	11.0%	3	16.244	5414.74	7182.76	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.5%	3	8.098	2699.33	3449.98	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.1%	4	3.132	782.95	2645.42	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.3%	4	0.375	93.68	156.45	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.2%	4	0.351	87.74	157.10	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.255	254.61	254.61	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	0.2%	1	0.249	249.21	249.21	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.227	227.32	227.32	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.159	158.85	158.85	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.114	113.59	113.59	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.080	79.74	79.74	2412.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.071	71.18	71.18	1503.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
Build ID: d718c8869ccdc11d01b7c17d16facbc2ada10760
Type: cpu
Time: 2026-10-31 16:02:23 JST
Duration: 60.09s, Total samples = 41.63s (69.28%)
Showing nodes accounting for 29.82s, 71.63% of 41.63s total
Dropped 1101 nodes (cum <= 0.21s)
Showing top 40 nodes out of 288
      flat  flat%   sum%        cum   cum%
     4.23s 10.16% 10.16%      4.23s 10.16%  internal/runtime/syscall/linux.Syscall6
     3.75s  9.01% 19.17%      3.75s  9.01%  crypto/internal/fips140/sha256.blockAVX2
     2.03s  4.88% 24.05%      2.65s  6.37%  runtime.findObject
     1.97s  4.73% 28.78%      1.97s  4.73%  runtime.memmove
     1.89s  4.54% 33.32%      7.57s 18.18%  runtime.scanObject
     1.74s  4.18% 37.50%      1.74s  4.18%  runtime.memclrNoHeapPointers
     1.57s  3.77% 41.27%      6.18s 14.85%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.51s  3.63% 44.90%      2.61s  6.27%  runtime.tryDeferToSpanScan
     1.48s  3.56% 48.45%      1.48s  3.56%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.45s  3.48% 51.93%      1.45s  3.48%  aeshashbody
     0.70s  1.68% 53.62%      0.70s  1.68%  runtime.spanClass.sizeclass (inline)
     0.63s  1.51% 55.13%      2.23s  5.36%  runtime.mapaccess1_faststr
     0.47s  1.13% 56.26%      0.47s  1.13%  memeqbody
     0.42s  1.01% 57.27%      0.65s  1.56%  example.com/isureal-company/isureal.postLess
     0.40s  0.96% 58.23%      0.40s  0.96%  runtime.nextFreeFast (inline)
     0.36s  0.86% 59.09%      0.82s  1.97%  runtime.greyobject
     0.34s  0.82% 59.91%      0.82s  1.97%  database/sql.convertAssignRows
     0.29s   0.7% 60.61%      0.29s   0.7%  runtime.futex
     0.28s  0.67% 61.28%      0.36s  0.86%  runtime.heapArenaOf (inline)
     0.28s  0.67% 61.95%      2.16s  5.19%  runtime.mapassign_faststr
     0.27s  0.65% 62.60%      0.58s  1.39%  compress/flate.(*compressor).deflate
     0.26s  0.62% 63.22%      0.29s   0.7%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.26s  0.62% 63.85%      0.35s  0.84%  runtime.(*mspan).markBitsForIndex (inline)
     0.25s   0.6% 64.45%      4.33s 10.40%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.25s   0.6% 65.05%      0.25s   0.6%  sync/atomic.(*Int32).Add
     0.24s  0.58% 65.63%      0.26s  0.62%  encoding/json.appendString[go.shape.string]
     0.22s  0.53% 66.15%      4.19s 10.06%  crypto/internal/fips140/sha256.(*Digest).Write
     0.22s  0.53% 66.68%      0.23s  0.55%  runtime.(*mspan).writeHeapBitsSmall
     0.21s   0.5% 67.19%      0.21s   0.5%  runtime.(*mspan).divideByElemSize (inline)
     0.21s   0.5% 67.69%      0.29s   0.7%  runtime.scanObjectSmall
     0.20s  0.48% 68.17%      3.04s  7.30%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.19s  0.46% 68.63%      0.23s  0.55%  runtime.typePointers.next
     0.18s  0.43% 69.06%      0.32s  0.77%  internal/runtime/maps.(*Iter).Next
     0.18s  0.43% 69.49%      1.15s  2.76%  runtime.mallocgcSmallScanNoHeader
     0.17s  0.41% 69.90%      0.24s  0.58%  runtime.spanOf (inline)
     0.16s  0.38% 70.29%      4.56s 10.95%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.15s  0.36% 70.65%      0.57s  1.37%  runtime.mapaccess2_faststr
     0.15s  0.36% 71.01%      0.46s  1.10%  runtime.scanObjectsSmall
     0.13s  0.31% 71.32%      0.25s   0.6%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.13s  0.31% 71.63%      0.55s  1.32%  runtime.convTstring
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 59.6% / 100% | 37% | 9.3% | 8 | 105 MB | CPU が飽和 |
| isu2 (2 cpu) | 25.9% / 100% | 3% | 0.7% | 4 | 324 MB |  |
| isu3 (2 cpu) | 21.9% / 100% | 2% | 0.5% | 5 | 597 MB | run queue が積もっている |
| isu4 (2 cpu) | 34.8% / 100% | 13% | 13.6% | 5 | 104 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 67% / nginx 12% / du 9% / kswapd0 3% / kcompactd0 1% |
| isu2 | isureal 48% / nginx 2% |
| isu3 | isureal 42% / nginx 1% |
| isu4 | isureal 27% / postgres 13% |

## Diff from previous run

スコア: 187910 → 171210 (-16700)

`d7fb2c5..6b8d03d`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf | 4 ++--
 1 file changed, 2 insertions(+), 2 deletions(-)
```
