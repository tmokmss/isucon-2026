# run 20261031-163017-8a13cd2

## Summary

| | |
|---|---|
| score | 206779 |
| status | pass |
| ref | skip-abandoned-writes |
| sha | 8a13cd28132d |
| 所要時間 | 162s |
| 前回比 | +1504 (0.7%) vs 20261031-162248-73b6aff |
| ベスト比 | +1504 (0.7%) vs 20261031-162248-73b6aff = 205275 **ベスト更新** |
| note | skip commits when client gave up during JPEG check |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	148	402	647	647
europe	7	19	56	160	452	452
america	6	19	55	158	426	426
```

```
{"id":"28648","status":"done","portalStatus":"finished","passed":true,"score":206779,"raw":206779,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:31:23.183Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:31:26.175Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:31:29.172Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:31:32.186Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T07:31:35.173Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T07:31:35.359Z level=INFO msg="user moved" user=benorzoe-5263 from=asia to=america
time=2026-10-31T07:31:36.390Z level=INFO msg="user moved" user=miroho-5395 from=asia to=europe
time=2026-10-31T07:31:38.173Z level=INFO msg="america 2 日目: 19 人"
time=2026-10-31T07:31:41.188Z level=INFO msg="asia 3 日目: 148 人"
time=2026-10-31T07:31:44.173Z level=INFO msg="europe 3 日目: 56 人"
time=2026-10-31T07:31:45.774Z level=INFO msg="user moved" user=solyuma_6078 from=asia to=america
time=2026-10-31T07:31:46.259Z level=INFO msg="user moved" user=miroho-5395 from=europe to=america
time=2026-10-31T07:31:46.411Z level=INFO msg="user moved" user=solyuma_6078 from=america to=asia
time=2026-10-31T07:31:47.174Z level=INFO msg="america 3 日目: 55 人"
time=2026-10-31T07:31:48.676Z level=INFO msg="user moved" user=haulleo_6343 from=europe to=asia
time=2026-10-31T07:31:49.553Z level=INFO msg="user moved" user=ninosol-6436 from=europe to=america
time=2026-10-31T07:31:49.570Z level=INFO msg="user moved" user=marsol6517 from=europe to=asia
time=2026-10-31T07:31:49.795Z level=INFO msg="user moved" user=sanoaul6566 from=america to=europe
time=2026-10-31T07:31:50.186Z level=INFO msg="asia 4 日目: 402 人"
time=2026-10-31T07:31:53.176Z level=INFO msg="europe 4 日目: 160 人"
time=2026-10-31T07:31:53.661Z level=INFO msg="user moved" user=benorzoe-5263 from=america to=asia
time=2026-10-31T07:31:55.653Z level=INFO msg="user moved" user=solyuma_6078 from=asia to=america
time=2026-10-31T07:31:55.869Z level=INFO msg="user moved" user=marsol6517 from=asia to=europe
time=2026-10-31T07:31:56.184Z level=INFO msg="america 4 日目: 158 人"
time=2026-10-31T07:31:56.317Z level=INFO msg="user moved" user=sanoaul6566 from=europe to=asia
time=2026-10-31T07:31:58.136Z level=INFO msg="user moved" user=miroho-5395 from=america to=europe
time=2026-10-31T07:31:59.194Z level=INFO msg="asia 5 日目: 647 人"
time=2026-10-31T07:32:00.307Z level=INFO msg="user moved" user=ninosol-6436 from=america to=europe
time=2026-10-31T07:32:00.929Z level=INFO msg="user moved" user=reansu-8950 from=europe to=america
time=2026-10-31T07:32:01.186Z level=INFO msg="user moved" user=tesa_8677 from=europe to=asia
time=2026-10-31T07:32:02.188Z level=INFO msg="europe 5 日目: 452 人"
time=2026-10-31T07:32:03.650Z level=INFO msg="user moved" user=wabenmax-9546 from=america to=europe
time=2026-10-31T07:32:05.266Z level=INFO msg="america 5 日目: 426 人"
time=2026-10-31T07:32:14.209Z level=INFO msg="asia 最終日（グランプリ）: 647 人"
time=2026-10-31T07:32:14.254Z level=INFO msg="europe 最終日（グランプリ）: 452 人"
time=2026-10-31T07:32:14.280Z level=INFO msg="america 最終日（グランプリ）: 426 人"
time=2026-10-31T07:32:23.260Z level=INFO msg="グランプリ: アジア 372 人, ヨーロッパ 442 人, アメリカ 177 人"
{
  "pass": true,
  "score": 206779,
  "messages": [
    "benyo5089: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "toan-6263: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hemar-9312: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "haeun_lim_2028: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "byoshin_kobo_2260: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "liamiase7003: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "shine-11528: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ulne-10082: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "anan_11306: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "warire-9352: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "yuto_tanaka_4663: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "anlia-15576: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
47.3%	1300	0	0	GET	/api/notifications	12969.781	9.977	15.029	15.037	173
10.9%	10271	0	0	GET	^/api/feeds/[^/]+	2979.350	0.290	1.786	2.554	5246
9.2%	2936	0	0	GET	/api/memories	2519.744	0.858	5.208	12.039	11147
6.3%	1907	49	0	POST	/api/posts	1731.263	0.908	4.553	6.605	633
5.4%	2616	0	0	GET	/api/posting-window/current	1481.956	0.566	1.863	2.323	237
3.3%	1564	1	0	GET	^/api/posts/me/current$	911.004	0.582	1.814	2.135	648
3.0%	1849	0	0	GET	/api/users	826.590	0.447	1.893	2.359	245
2.8%	224	0	0	GET	/api/me/export	762.046	3.402	22.205	23.193	1841571
1.5%	491	126	0	PUT	/api/me/profile-image	402.018	0.819	11.056	16.195	4
1.4%	11526	244	0	GET	^/api/images/[^/]+	395.321	0.034	0.686	1.742	159656
1.3%	1266	0	0	GET	^/api/relationships/[^/]+	345.669	0.273	1.677	2.445	1576
1.2%	844	4	0	GET	^/api/posts/[^/]+/realmojis$	333.025	0.395	2.917	4.731	2464
1.0%	785	12	0	POST	^/api/posts/[^/]+/realmojis$	282.633	0.360	3.425	4.533	471
0.9%	377	3	0	GET	/api/me	251.819	0.668	2.143	2.376	299
0.9%	1830	53	0	GET	^/api/posts/[^/]+$	233.664	0.128	0.996	1.790	619
0.7%	497	0	0	GET	^/api/memories/[^/]+	180.569	0.363	2.150	2.795	646
0.6%	3464	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	168.925	0.049	0.450	0.664	116
0.5%	3865	43	0	POST	/api/friend-requests	146.928	0.038	0.374	0.544	114
0.4%	968	0	0	GET	^/api/users/[^/]+$	121.015	0.125	0.869	1.752	209
0.4%	681	6	0	GET	^/api/posts/[^/]+/comments$	105.057	0.154	1.023	1.409	2855
0.3%	635	0	0	POST	^/api/posts/[^/]+/comments$	74.859	0.118	0.803	0.981	348
0.2%	3907	552	0	POST	/api/auth/login	47.593	0.012	0.244	0.665	213
0.1%	119	0	0	DELETE	^/api/realmojis/[^/]+$	28.449	0.239	0.690	0.713	0
0.1%	74	0	0	DELETE	^/api/posts/[^/]+$	26.421	0.357	1.024	1.024	0
0.1%	1	0	0	POST	/api/initialize	21.942	21.942	21.942	21.942	90
0.0%	179	0	0	PATCH	^/api/posts/[^/]+$	12.183	0.068	0.406	0.469	647
0.0%	585	0	0	POST	/api/auth/logout	11.332	0.019	0.310	0.363	0
0.0%	135	0	0	DELETE	^/api/comments/[^/]+$	10.034	0.074	0.494	1.030	0
0.0%	403	0	0	PUT	^/api/follows/[^/]+$	5.264	0.013	0.275	0.405	0
0.0%	310	0	0	PATCH	/api/me	3.374	0.011	0.204	0.620	274
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
47.0%	726	1	0	GET	/api/notifications	5531.391	7.619	11.894	11.912	167
8.3%	7877	0	0	GET	^/api/feeds/[^/]+	973.024	0.124	1.360	2.433	4635
7.6%	6724	32	0	GET	^/api/images/[^/]+	893.396	0.133	1.139	3.543	190172
6.9%	1146	6	0	POST	/api/posts	811.460	0.708	2.623	7.648	648
6.1%	2706	25	0	POST	/api/friend-requests	719.070	0.266	0.728	1.720	114
3.0%	1846	0	0	POST	/api/auth/login	355.236	0.192	0.658	1.764	248
2.8%	1741	0	0	GET	/api/memories	327.176	0.188	1.793	3.210	13094
2.0%	1650	0	0	GET	/api/posting-window/current	240.135	0.146	0.742	1.909	238
1.9%	513	20	0	POST	^/api/posts/[^/]+/realmojis$	224.653	0.438	3.590	5.384	466
1.9%	1116	0	0	GET	^/api/posts/me/current$	219.917	0.197	1.637	2.566	652
1.6%	1263	0	0	GET	/api/users	186.118	0.147	0.715	1.569	247
1.2%	699	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	135.310	0.194	0.720	1.558	116
1.1%	264	0	0	PUT	/api/me/profile-image	134.417	0.509	1.843	1.878	0
1.1%	1337	29	0	GET	^/api/posts/[^/]+$	128.481	0.096	0.416	1.520	629
1.1%	452	0	0	POST	/api/auth/register	125.852	0.278	0.616	0.679	229
0.9%	935	0	0	GET	^/api/relationships/[^/]+	103.434	0.111	0.701	2.499	668
0.7%	470	2	0	POST	^/api/posts/[^/]+/comments$	86.992	0.185	1.012	1.632	343
0.7%	299	0	0	PUT	^/api/follows/[^/]+$	85.897	0.287	0.730	1.185	0
0.5%	649	0	0	GET	^/api/users/[^/]+$	59.316	0.091	0.341	1.688	222
0.5%	215	0	0	PATCH	/api/me	56.229	0.262	0.662	0.671	281
0.5%	625	0	0	GET	^/api/posts/[^/]+/realmojis$	53.977	0.086	0.317	1.488	3525
0.4%	299	0	0	POST	/api/auth/logout	51.339	0.172	1.321	1.542	0
0.4%	278	2	0	GET	/api/me	50.067	0.180	0.721	0.775	301
0.4%	510	0	0	GET	^/api/posts/[^/]+/comments$	43.699	0.086	0.298	0.412	4464
0.3%	98	0	0	DELETE	^/api/realmojis/[^/]+$	37.847	0.386	0.964	0.964	0
0.3%	105	1	0	PATCH	^/api/posts/[^/]+$	35.451	0.338	1.724	1.824	638
0.2%	41	0	0	DELETE	^/api/posts/[^/]+$	24.003	0.585	1.056	1.056	0
0.2%	247	0	0	GET	^/api/memories/[^/]+	23.709	0.096	0.554	0.760	708
0.2%	111	0	0	GET	/api/me/export	22.442	0.202	2.193	3.496	1949988
0.1%	88	0	0	DELETE	^/api/comments/[^/]+$	12.887	0.146	0.653	0.653	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.9%	681	0	0	GET	/api/notifications	4118.321	6.047	9.211	9.213	177
10.3%	1087	36	0	POST	/api/posts	1008.315	0.928	3.316	7.629	642
8.4%	3930	0	0	GET	^/api/feeds/[^/]+	825.621	0.210	2.043	2.963	3259
8.3%	5030	31	0	GET	^/api/images/[^/]+	817.383	0.163	1.451	3.473	152802
5.5%	2353	14	0	POST	/api/friend-requests	538.795	0.229	0.778	1.847	114
3.8%	1565	0	0	POST	/api/auth/login	373.549	0.239	1.015	2.524	250
3.2%	1274	0	0	GET	/api/memories	312.421	0.245	2.488	2.994	15894
2.7%	1539	0	0	GET	/api/posting-window/current	262.395	0.170	1.056	2.210	240
2.3%	780	12	0	GET	^/api/posts/me/current$	225.277	0.289	2.475	2.913	654
1.7%	1165	0	0	GET	/api/users	169.205	0.145	0.887	1.659	249
1.6%	425	0	0	POST	/api/auth/register	159.786	0.376	1.021	1.069	229
1.4%	259	0	0	PUT	/api/me/profile-image	139.419	0.538	1.748	1.928	0
1.1%	251	16	0	POST	^/api/posts/[^/]+/realmojis$	111.916	0.446	1.589	4.795	451
1.1%	444	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	105.767	0.238	1.647	2.025	116
0.8%	509	0	0	GET	^/api/relationships/[^/]+	81.127	0.159	1.426	2.449	623
0.8%	520	35	0	GET	^/api/posts/[^/]+$	79.135	0.152	1.029	1.596	605
0.7%	263	0	0	PUT	^/api/follows/[^/]+$	64.238	0.244	0.897	1.015	0
0.6%	250	0	0	POST	/api/auth/logout	58.738	0.235	1.774	2.042	0
0.6%	192	0	0	PATCH	/api/me	54.126	0.282	0.876	1.014	288
0.5%	211	5	0	POST	^/api/posts/[^/]+/comments$	48.961	0.232	0.674	0.716	374
0.4%	255	0	0	GET	^/api/users/[^/]+$	39.304	0.154	1.039	1.099	217
0.4%	108	0	0	GET	/api/me/export	38.176	0.353	4.853	5.218	3027799
0.4%	225	0	0	GET	^/api/posts/[^/]+/realmojis$	36.244	0.161	0.993	1.083	3457
0.3%	222	0	0	GET	^/api/memories/[^/]+	33.810	0.152	1.221	2.233	1027
0.3%	266	1	0	GET	/api/me	33.323	0.125	0.558	0.628	303
0.3%	162	0	0	GET	^/api/posts/[^/]+/comments$	24.891	0.154	0.643	0.705	4300
0.2%	75	3	0	PATCH	^/api/posts/[^/]+$	24.494	0.327	1.847	1.847	632
0.2%	35	0	0	DELETE	^/api/posts/[^/]+$	20.530	0.587	2.769	2.769	0
0.1%	33	0	0	DELETE	^/api/realmojis/[^/]+$	11.409	0.346	0.611	0.611	0
0.1%	27	0	0	DELETE	^/api/comments/[^/]+$	5.660	0.210	0.650	0.650	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 10901 クエリ / 48 種類 / 合計 130.170s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	42.5%	235	55.361	235.58	794.78	1184.0	120.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	20.0%	4	26.012	6503.04	10011.26	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.4%	4	18.715	4678.83	7223.02	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.6%	7243	11.159	1.54	7.79	10.2	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	7.1%	4	9.275	2318.80	3443.81	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	4.5%	7	5.856	836.55	3804.41	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.8%	7	1.072	153.12	610.04	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.5%	7	0.596	85.13	160.99	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.4%	1	0.460	460.45	460.45	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.3%	1	0.383	383.14	383.14	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.2%	1	0.230	230.14	230.14	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.187	186.93	186.93	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.147	147.06	147.06	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.093	92.89	92.89	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
15	0.1%	1	0.093	92.88	92.88	2414.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM posts;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 47f958816702caae388e20a728771507b627cca8
Type: cpu
Time: 2026-10-31 16:30:57 JST
Duration: 60s, Total samples = 26.96s (44.93%)
Showing nodes accounting for 20.59s, 76.37% of 26.96s total
Dropped 807 nodes (cum <= 0.13s)
Showing top 40 nodes out of 235
      flat  flat%   sum%        cum   cum%
     4.53s 16.80% 16.80%      4.53s 16.80%  crypto/internal/fips140/sha256.blockAVX2
     2.64s  9.79% 26.59%      2.64s  9.79%  internal/runtime/syscall/linux.Syscall6
     1.83s  6.79% 33.38%      6.09s 22.59%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.43s  5.30% 38.69%      1.43s  5.30%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.29s  4.78% 43.47%      1.29s  4.78%  runtime.memmove
     1.03s  3.82% 47.29%      1.03s  3.82%  aeshashbody
     0.90s  3.34% 50.63%      0.90s  3.34%  runtime.memclrNoHeapPointers
     0.68s  2.52% 53.15%      0.84s  3.12%  runtime.findObject
     0.56s  2.08% 55.23%      2.36s  8.75%  runtime.scanObject
     0.48s  1.78% 57.01%      1.90s  7.05%  runtime.mapaccess1_faststr
     0.39s  1.45% 58.46%      0.68s  2.52%  runtime.tryDeferToSpanScan
     0.31s  1.15% 59.61%      0.49s  1.82%  example.com/isureal-company/isureal.postLess
     0.28s  1.04% 60.65%      0.28s  1.04%  memeqbody
     0.27s  1.00% 61.65%      0.38s  1.41%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.26s  0.96% 62.61%      0.39s  1.45%  database/sql.convertAssignRows
     0.26s  0.96% 63.58%      0.78s  2.89%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.25s  0.93% 64.50%      4.96s 18.40%  crypto/internal/fips140/sha256.(*Digest).Write
     0.25s  0.93% 65.43%      0.25s  0.93%  runtime.nextFreeFast (inline)
     0.22s  0.82% 66.25%      5.06s 18.77%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.22s  0.82% 67.06%      0.22s  0.82%  runtime.futex
     0.21s  0.78% 67.84%      0.84s  3.12%  runtime.mallocgcSmallScanNoHeader
     0.20s  0.74% 68.58%      0.20s  0.74%  runtime.spanClass.sizeclass (inline)
     0.20s  0.74% 69.32%      0.20s  0.74%  sync/atomic.(*Int32).Add (inline)
     0.19s   0.7% 70.03%      0.19s   0.7%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.19s   0.7% 70.73%      1.52s  5.64%  runtime.mapassign_faststr
     0.13s  0.48% 71.22%      5.26s 19.51%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.13s  0.48% 71.70%      0.14s  0.52%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.13s  0.48% 72.18%      1.20s  4.45%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.45% 72.63%      5.54s 20.55%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.12s  0.45% 73.07%      2.16s  8.01%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.12s  0.45% 73.52%      1.59s  5.90%  runtime.mallocgc
     0.11s  0.41% 73.92%      0.24s  0.89%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.11s  0.41% 74.33%      0.21s  0.78%  runtime.greyobject
     0.10s  0.37% 74.70%      0.36s  1.34%  runtime.mapaccess2_faststr
     0.09s  0.33% 75.04%      0.20s  0.74%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.09s  0.33% 75.37%      3.85s 14.28%  example.com/isureal-company/isureal.loadBig
     0.08s   0.3% 75.67%      0.50s  1.85%  runtime.convTstring
     0.07s  0.26% 75.93%      0.19s   0.7%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.06s  0.22% 76.15%      0.57s  2.11%  github.com/jackc/pgx/v5/pgtype.scanPlanString.Scan
     0.06s  0.22% 76.37%      0.47s  1.74%  sort.partition_func
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 34.4% / 100% | 11% | 0.7% | 8 | 112 MB | run queue が積もっている |
| isu2 (2 cpu) | 26.9% / 100% | 9% | 0.8% | 5 | 170 MB | run queue が積もっている |
| isu3 (2 cpu) | 28.1% / 100% | 10% | 1.3% | 5 | 153 MB | run queue が積もっている |
| isu4 (2 cpu) | 64.9% / 100% | 56% | 2.3% | 7 | 103 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 55% / nginx 10% |
| isu2 | isureal 49% / nginx 3% |
| isu3 | isureal 52% / nginx 2% |
| isu4 | isureal 108% / postgres 16% |

## Diff from previous run

スコア: 205275 → 206779 (+1504)

`73b6aff..8a13cd2`

```
 webapp/go/post_api.go          | 4 ++++
 webapp/go/profile_image_api.go | 4 ++++
 webapp/go/realmoji_api.go      | 4 ++++
 3 files changed, 12 insertions(+)
```
