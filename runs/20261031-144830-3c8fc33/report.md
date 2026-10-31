# run 20261031-144830-3c8fc33

## Summary

| | |
|---|---|
| score | 136044 |
| status | pass |
| ref | 3c8fc33 |
| sha | 3c8fc3341aa7 |
| 所要時間 | 203s |
| 前回比 | +15481 (12.8%) vs 20261031-144534-f7ca385 |
| ベスト比 | -5291 (-3.7%) vs 20261031-141027-bd0fe45 = 141335 |
| note | DB外し5: 追記でfsyncしない |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	149	433	1032	1032
europe	7	18	53	151	245	244
america	6	19	62	169	169	169
```

```
{"id":"19400","status":"done","portalStatus":"finished","passed":true,"score":136044,"raw":136044,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:49:58.047Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:50:01.051Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:50:04.035Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:50:07.045Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T05:50:09.189Z level=INFO msg="user moved" user=miasana.5278 from=asia to=america
time=2026-10-31T05:50:10.036Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T05:50:10.319Z level=INFO msg="user moved" user=miasana.5278 from=america to=europe
time=2026-10-31T05:50:13.036Z level=INFO msg="america 2 日目: 19 人"
time=2026-10-31T05:50:16.046Z level=INFO msg="asia 3 日目: 149 人"
time=2026-10-31T05:50:18.443Z level=INFO msg="user moved" user=benheha_5890 from=asia to=america
time=2026-10-31T05:50:19.035Z level=INFO msg="europe 3 日目: 53 人"
time=2026-10-31T05:50:20.214Z level=INFO msg="user moved" user=yumi6030 from=asia to=america
time=2026-10-31T05:50:21.507Z level=INFO msg="user moved" user=miasana.5278 from=europe to=asia
time=2026-10-31T05:50:22.036Z level=INFO msg="america 3 日目: 62 人"
time=2026-10-31T05:50:25.051Z level=INFO msg="asia 4 日目: 433 人"
time=2026-10-31T05:50:27.062Z level=INFO msg="user moved" user=watomu_5452 from=europe to=america
time=2026-10-31T05:50:27.570Z level=INFO msg="user moved" user=miasana.5278 from=asia to=america
time=2026-10-31T05:50:28.766Z level=INFO msg="europe 4 日目: 151 人"
time=2026-10-31T05:50:29.092Z level=INFO msg="user moved" user=ulhahi-7261 from=asia to=europe
time=2026-10-31T05:50:29.401Z level=INFO msg="user moved" user=ivyse-7394 from=asia to=america
time=2026-10-31T05:50:31.297Z level=INFO msg="america 4 日目: 169 人"
time=2026-10-31T05:50:32.220Z level=INFO msg="user moved" user=torta_8619 from=asia to=europe
time=2026-10-31T05:50:34.120Z level=INFO msg="asia 5 日目: 1032 人"
time=2026-10-31T05:50:38.243Z level=INFO msg="europe 5 日目: 245 人"
time=2026-10-31T05:50:41.757Z level=INFO msg="america 5 日目: 169 人"
time=2026-10-31T05:50:50.194Z level=INFO msg="asia 最終日（グランプリ）: 1032 人"
time=2026-10-31T05:50:50.609Z level=INFO msg="europe 最終日（グランプリ）: 244 人"
time=2026-10-31T05:50:50.632Z level=INFO msg="america 最終日（グランプリ）: 169 人"
time=2026-10-31T05:50:58.091Z level=INFO msg="グランプリ: アジア 414 人, ヨーロッパ 120 人, アメリカ 87 人"
{
  "pass": true,
  "score": 136044,
  "messages": [
    "kochi.7418: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "kohe-5181: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "andanmo-7563: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "chisoeva_5594: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "vinno_5906: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "koleose10384: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "zoeissol_7663: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "seivyshi7547: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "kael6041: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "namais8592: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "sesemar-8615: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "mean.7306: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "maxzoe.5819: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "solnetsu.5326: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "fumi-5726: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "sume-7025: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "yufu-8500: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "maryano-8017: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "solan_5648: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "toyu8042: 投稿できません: POST /api/posts: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
67.1%	2342	0	0	GET	/api/notifications	63367.895	27.057	60.003	60.367	109
12.9%	2330	370	0	POST	/api/posts	12216.362	5.243	24.427	64.045	548
6.0%	736	129	0	PUT	/api/me/profile-image	5632.565	7.653	23.520	24.603	3
3.0%	17576	362	0	GET	^/api/images/[^/]+	2852.365	0.162	1.791	3.497	97780
2.0%	4978	68	0	POST	/api/friend-requests	1853.446	0.372	2.083	2.972	114
1.4%	6672	713	0	POST	/api/auth/login	1300.034	0.195	1.803	3.688	221
1.3%	4387	0	0	GET	/api/memories	1263.570	0.288	2.848	6.390	12830
1.1%	11125	0	0	GET	^/api/feeds/[^/]+	1027.646	0.092	0.732	2.202	2575
0.9%	513	21	0	POST	^/api/posts/[^/]+/realmojis$	893.227	1.741	22.692	23.002	448
0.7%	1054	1	0	POST	/api/auth/register	658.456	0.625	2.557	2.610	227
0.7%	2042	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	649.728	0.318	1.956	2.548	116
0.4%	3338	0	0	GET	/api/posting-window/current	377.481	0.113	0.859	1.514	237
0.3%	391	0	0	GET	/api/me/export	327.859	0.839	6.910	8.587	1914905
0.3%	2403	0	0	GET	/api/users	245.441	0.102	0.657	1.533	253
0.3%	708	0	0	PUT	^/api/follows/[^/]+$	239.298	0.338	1.992	2.667	0
0.2%	986	0	0	POST	/api/auth/logout	208.658	0.212	1.746	2.552	0
0.2%	1560	19	0	GET	^/api/posts/me/current$	204.352	0.131	0.810	1.182	640
0.2%	502	0	0	PATCH	/api/me	167.301	0.333	2.034	2.923	259
0.2%	417	0	0	POST	^/api/posts/[^/]+/comments$	158.175	0.379	2.057	2.909	348
0.1%	1445	0	0	GET	^/api/relationships/[^/]+	127.484	0.088	0.735	1.515	994
0.1%	1600	17	0	GET	^/api/posts/[^/]+$	123.738	0.077	0.641	1.517	614
0.1%	942	0	0	GET	^/api/memories/[^/]+	113.654	0.121	0.989	1.785	659
0.1%	212	3	0	PATCH	^/api/posts/[^/]+$	79.324	0.374	1.919	2.164	630
0.1%	78	0	0	DELETE	^/api/posts/[^/]+$	74.010	0.949	3.108	3.108	0
0.1%	788	0	0	GET	^/api/users/[^/]+$	52.443	0.067	0.494	0.892	200
0.1%	410	2	0	GET	/api/me	51.842	0.126	0.656	1.336	301
0.1%	728	0	0	GET	^/api/posts/[^/]+/realmojis$	49.492	0.068	0.676	0.993	1159
0.0%	612	0	0	GET	^/api/posts/[^/]+/comments$	41.437	0.068	0.696	1.150	1551
0.0%	82	0	0	DELETE	^/api/realmojis/[^/]+$	37.350	0.455	1.527	1.527	0
0.0%	78	0	0	DELETE	^/api/comments/[^/]+$	18.648	0.239	1.203	1.203	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
72.4%	655	1	0	GET	/api/notifications	15988.995	24.411	60.610	61.097	117
5.8%	2708	17	0	GET	^/api/images/[^/]+	1290.065	0.476	2.606	3.558	159558
4.7%	620	11	0	POST	/api/posts	1029.417	1.660	4.540	5.540	634
3.2%	1157	6	0	POST	/api/friend-requests	699.408	0.605	2.276	2.634	114
2.2%	2152	0	0	GET	^/api/feeds/[^/]+	483.167	0.225	1.705	2.020	2373
2.0%	840	0	0	POST	/api/auth/login	438.251	0.522	2.116	2.527	247
1.8%	135	0	0	PUT	/api/me/profile-image	399.409	2.959	22.716	22.732	0
1.5%	874	0	0	GET	/api/posting-window/current	321.689	0.368	1.377	2.054	239
0.8%	772	0	0	GET	/api/memories	186.729	0.242	2.427	4.605	12662
0.8%	279	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	168.036	0.602	2.576	2.619	116
0.7%	242	0	0	POST	/api/auth/register	150.464	0.622	2.103	2.188	228
0.5%	564	0	0	GET	/api/users	120.506	0.214	1.425	2.278	251
0.5%	494	12	0	GET	^/api/posts/me/current$	120.132	0.243	1.593	1.801	629
0.5%	109	8	0	POST	^/api/posts/[^/]+/realmojis$	115.424	1.059	19.678	21.118	404
0.4%	171	0	0	PUT	^/api/follows/[^/]+$	82.790	0.484	1.982	2.267	0
0.3%	116	1	0	PATCH	/api/me	68.085	0.587	1.986	1.997	274
0.3%	141	1	0	POST	^/api/posts/[^/]+/comments$	65.849	0.467	2.119	2.451	310
0.3%	136	0	0	POST	/api/auth/logout	62.980	0.463	1.954	2.105	0
0.2%	275	0	0	GET	^/api/relationships/[^/]+	54.608	0.199	1.253	2.018	764
0.2%	44	0	0	GET	/api/me/export	39.417	0.896	8.101	8.101	2392825
0.2%	274	12	0	GET	^/api/posts/[^/]+$	38.047	0.139	1.248	1.274	584
0.1%	126	2	0	GET	/api/me	31.742	0.252	1.281	1.284	299
0.1%	44	4	0	PATCH	^/api/posts/[^/]+$	29.967	0.681	2.190	2.190	597
0.1%	146	0	0	GET	^/api/users/[^/]+$	22.671	0.155	1.245	1.612	190
0.1%	20	0	0	DELETE	^/api/posts/[^/]+$	20.550	1.028	2.436	2.436	0
0.1%	100	0	0	GET	^/api/memories/[^/]+	19.378	0.194	1.280	1.566	617
0.1%	99	0	0	GET	^/api/posts/[^/]+/realmojis$	15.632	0.158	0.969	0.969	1064
0.1%	86	0	0	GET	^/api/posts/[^/]+/comments$	11.086	0.129	0.941	0.941	1323
0.0%	19	0	0	DELETE	^/api/comments/[^/]+$	10.040	0.528	2.103	2.103	0
0.0%	24	0	0	DELETE	^/api/realmojis/[^/]+$	9.241	0.385	1.172	1.172	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
71.2%	561	0	0	GET	/api/notifications	12727.894	22.688	60.491	60.770	126
8.0%	2564	18	0	GET	^/api/images/[^/]+	1423.342	0.555	2.907	3.306	160382
4.8%	520	16	0	POST	/api/posts	859.351	1.653	4.357	4.816	622
3.2%	906	6	0	POST	/api/friend-requests	573.036	0.632	2.278	2.657	114
1.9%	716	0	0	POST	/api/auth/login	339.022	0.473	2.045	2.698	250
1.7%	1709	0	0	GET	^/api/feeds/[^/]+	304.056	0.178	1.097	1.356	2549
1.4%	706	0	0	GET	/api/posting-window/current	256.588	0.363	1.287	1.880	240
1.2%	330	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	211.472	0.641	2.270	2.295	116
1.1%	75	10	0	POST	^/api/posts/[^/]+/realmojis$	189.323	2.524	23.598	23.598	399
0.9%	78	0	0	PUT	/api/me/profile-image	167.667	2.150	22.996	22.996	0
0.9%	672	0	0	GET	/api/memories	159.955	0.238	1.930	6.855	14672
0.6%	449	0	0	GET	/api/users	105.807	0.236	1.195	1.320	257
0.4%	444	22	0	GET	^/api/posts/me/current$	78.590	0.177	1.151	1.287	610
0.4%	130	0	0	POST	/api/auth/logout	74.160	0.570	2.083	2.284	0
0.3%	109	0	0	PUT	^/api/follows/[^/]+$	49.851	0.457	2.270	2.664	0
0.3%	163	0	0	POST	/api/auth/register	47.549	0.292	0.814	0.867	229
0.3%	217	0	0	GET	^/api/relationships/[^/]+	46.533	0.214	1.785	1.908	1104
0.2%	82	0	0	POST	^/api/posts/[^/]+/comments$	38.583	0.471	2.109	2.109	306
0.2%	78	0	0	PATCH	/api/me	36.845	0.472	2.270	2.270	269
0.2%	122	0	0	GET	^/api/memories/[^/]+	28.563	0.234	1.515	2.138	1039
0.2%	191	12	0	GET	^/api/posts/[^/]+$	28.058	0.147	1.279	1.322	576
0.1%	37	3	0	PATCH	^/api/posts/[^/]+$	25.766	0.696	2.264	2.264	586
0.1%	19	0	0	DELETE	^/api/posts/[^/]+$	24.041	1.265	2.402	2.402	0
0.1%	78	0	0	GET	/api/me	22.031	0.282	1.286	1.286	302
0.1%	39	0	0	GET	/api/me/export	20.590	0.528	7.321	7.321	2760987
0.1%	120	0	0	GET	^/api/users/[^/]+$	19.806	0.165	0.899	1.265	196
0.1%	59	0	0	GET	^/api/posts/[^/]+/realmojis$	9.991	0.169	1.264	1.264	1145
0.0%	60	0	0	GET	^/api/posts/[^/]+/comments$	8.919	0.149	0.671	0.671	1427
0.0%	17	0	0	DELETE	^/api/comments/[^/]+$	5.002	0.294	0.778	0.778	0
0.0%	4	0	0	PUT	^/api/blocks/[^/]+$	2.742	0.685	1.708	1.708	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 15758 クエリ / 48 種類 / 合計 152.077s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	54.1%	405	82.272	203.14	1069.97	1224.6	125.4	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.2%	3	21.547	7182.41	9540.11	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.1%	13404	16.915	1.26	6.63	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.5%	3	15.967	5322.46	7186.18	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.1%	3	7.787	2595.78	3546.08	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.5%	6	3.768	627.93	2866.93	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.8%	1	1.191	1191.45	1191.45	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	0.5%	6	0.685	114.10	157.54	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.4%	6	0.659	109.83	159.42	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.2%	1	0.248	247.96	247.96	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.1%	1	0.225	225.17	225.17	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.156	156.36	156.36	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.113	112.84	112.84	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.106	106.11	106.11	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
15	0.1%	1	0.082	82.38	82.38	2493.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…

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
Build ID: 0ed586a5c14f3975f2896bb26b35cd11731c1c95
Type: cpu
Time: 2026-10-31 14:49:36 JST
Duration: 60.54s, Total samples = 45.28s (74.79%)
Showing nodes accounting for 31.84s, 70.32% of 45.28s total
Dropped 1140 nodes (cum <= 0.23s)
Showing top 40 nodes out of 281
      flat  flat%   sum%        cum   cum%
     5.88s 12.99% 12.99%      5.88s 12.99%  crypto/internal/fips140/sha256.blockAVX2
        5s 11.04% 24.03%         5s 11.04%  internal/runtime/syscall/linux.Syscall6
     2.06s  4.55% 28.58%      2.06s  4.55%  runtime.memmove
     1.83s  4.04% 32.62%      2.29s  5.06%  runtime.findObject
     1.79s  3.95% 36.57%      5.70s 12.59%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.63s  3.60% 40.17%      1.63s  3.60%  runtime.memclrNoHeapPointers
     1.55s  3.42% 43.60%      7.07s 15.61%  runtime.scanObject
     1.40s  3.09% 46.69%      1.40s  3.09%  aeshashbody
     1.40s  3.09% 49.78%      2.75s  6.07%  runtime.tryDeferToSpanScan
     1.33s  2.94% 52.72%      1.33s  2.94%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.94s  2.08% 54.79%      0.94s  2.08%  runtime.spanClass.sizeclass (inline)
     0.59s  1.30% 56.10%      2.26s  4.99%  runtime.mapaccess1_faststr
     0.58s  1.28% 57.38%      0.58s  1.28%  memeqbody
     0.41s  0.91% 58.28%      0.41s  0.91%  runtime.nextFreeFast (inline)
     0.34s  0.75% 59.03%      0.57s  1.26%  example.com/isureal-company/isureal.postLess
     0.30s  0.66% 59.70%      4.29s  9.47%  runtime.mallocgc
     0.29s  0.64% 60.34%      5.71s 12.61%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.28s  0.62% 60.95%      0.50s  1.10%  compress/flate.(*compressor).deflate
     0.28s  0.62% 61.57%      0.32s  0.71%  runtime.typePointers.next
     0.27s   0.6% 62.17%      0.70s  1.55%  runtime.greyobject
     0.26s  0.57% 62.74%      0.26s  0.57%  runtime.futex
     0.25s  0.55% 63.30%      0.28s  0.62%  encoding/json.appendString[go.shape.string]
     0.23s  0.51% 63.80%      1.12s  2.47%  runtime.mallocgcSmallScanNoHeader
     0.23s  0.51% 64.31%      0.23s  0.51%  sync/atomic.(*Int32).Add
     0.22s  0.49% 64.80%      6.45s 14.24%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.22s  0.49% 65.28%         6s 13.25%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.22s  0.49% 65.77%      0.27s   0.6%  runtime.heapArenaOf (inline)
     0.22s  0.49% 66.25%      0.52s  1.15%  runtime.scanObjectsSmall
     0.21s  0.46% 66.72%      6.37s 14.07%  crypto/internal/fips140/sha256.(*Digest).Write
     0.20s  0.44% 67.16%      0.23s  0.51%  runtime.(*mspan).markBitsForIndex (inline)
     0.20s  0.44% 67.60%      2.02s  4.46%  runtime.mapassign_faststr
     0.19s  0.42% 68.02%      0.64s  1.41%  database/sql.convertAssignRows
     0.18s   0.4% 68.42%      0.32s  0.71%  internal/runtime/maps.(*Iter).Next
     0.18s   0.4% 68.82%      0.90s  1.99%  runtime.scanSpan
     0.17s  0.38% 69.19%      0.91s  2.01%  runtime.mallocgcSmallNoscan
     0.12s  0.27% 69.46%      1.16s  2.56%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.11s  0.24% 69.70%      0.29s  0.64%  internal/stringslite.Index
     0.10s  0.22% 69.92%      0.74s  1.63%  database/sql.(*Rows).scanLocked
     0.09s   0.2% 70.12%      5.97s 13.18%  crypto/internal/fips140/sha256.block
     0.09s   0.2% 70.32%      0.73s  1.61%  encoding/json.structEncoder.encode
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 46.8% / 100% | 33% | 0.7% | 7 | 152 MB | CPU が飽和 |
| isu2 (2 cpu) | 20.9% / 100% | 1% | 0.5% | 4 | 668 MB |  |
| isu3 (2 cpu) | 22.0% / 100% | 1% | 0.6% | 5 | 610 MB | run queue が積もっている |
| isu4 (2 cpu) | 30.8% / 100% | 7% | 15.9% | 5 | 106 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 75% / nginx 10% |
| isu2 | isureal 39% / nginx 1% |
| isu3 | isureal 42% / nginx 1% |
| isu4 | isureal 15% / postgres 13% |

## Diff from previous run

スコア: 120563 → 136044 (+15481)

`f7ca385..3c8fc33`

```
 webapp/go/persist.go | 17 +++++------------
 1 file changed, 5 insertions(+), 12 deletions(-)
```
