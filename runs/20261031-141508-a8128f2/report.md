# run 20261031-141508-a8128f2

## Summary

| | |
|---|---|
| score | 134342 |
| status | pass |
| ref | image-files |
| sha | a8128f2ba308 |
| 所要時間 | 370s |
| 前回比 | +10791 (8.7%) vs 20261031-140544-247167c |
| ベスト比 | -138 (-0.1%) vs 20261031-140017-04990e1 = 134480 |
| note | image-files (no fsync/sync) on main e42c64c |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	52	153	448	1086	1085
europe	7	19	61	169	220	220
america	6	13	35	87	115	115
```

```
{"id":"16844","status":"done","portalStatus":"finished","passed":true,"score":134342,"raw":134342,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:19:43.813Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:19:46.784Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:19:49.784Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:19:52.805Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T05:19:55.784Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T05:19:58.784Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T05:20:01.806Z level=INFO msg="asia 3 日目: 153 人"
time=2026-10-31T05:20:02.553Z level=INFO msg="user moved" user=vinyoto.5539 from=asia to=america
time=2026-10-31T05:20:03.812Z level=INFO msg="user moved" user=vinyoto.5539 from=america to=asia
time=2026-10-31T05:20:04.096Z level=INFO msg="user moved" user=nura5931 from=asia to=europe
time=2026-10-31T05:20:04.128Z level=INFO msg="user moved" user=kaevato5651 from=asia to=europe
time=2026-10-31T05:20:04.784Z level=INFO msg="europe 3 日目: 61 人"
time=2026-10-31T05:20:07.785Z level=INFO msg="america 3 日目: 35 人"
time=2026-10-31T05:20:09.989Z level=INFO msg="user moved" user=shiyuha_6502 from=europe to=asia
time=2026-10-31T05:20:10.467Z level=INFO msg="user moved" user=danhoor_6412 from=europe to=asia
time=2026-10-31T05:20:10.815Z level=INFO msg="asia 4 日目: 448 人"
time=2026-10-31T05:20:12.670Z level=INFO msg="user moved" user=shiyuha_6502 from=asia to=europe
time=2026-10-31T05:20:12.709Z level=INFO msg="user moved" user=vinyoto.5539 from=asia to=america
time=2026-10-31T05:20:13.032Z level=INFO msg="user moved" user=danhoor_6412 from=asia to=europe
time=2026-10-31T05:20:13.884Z level=INFO msg="user moved" user=zoeivy.7349 from=asia to=europe
time=2026-10-31T05:20:13.939Z level=INFO msg="europe 4 日目: 169 人"
time=2026-10-31T05:20:14.703Z level=INFO msg="user moved" user=takotsu_6997 from=asia to=europe
time=2026-10-31T05:20:15.882Z level=INFO msg="user moved" user=yaul.7415 from=asia to=europe
time=2026-10-31T05:20:16.994Z level=INFO msg="america 4 日目: 87 人"
time=2026-10-31T05:20:19.112Z level=INFO msg="user moved" user=nura5931 from=europe to=america
time=2026-10-31T05:20:19.584Z level=INFO msg="user moved" user=somumi-7859 from=asia to=america
time=2026-10-31T05:20:19.704Z level=INFO msg="user moved" user=neya.8273 from=asia to=america
time=2026-10-31T05:20:19.891Z level=INFO msg="asia 5 日目: 1086 人"
time=2026-10-31T05:20:22.872Z level=INFO msg="europe 5 日目: 220 人"
time=2026-10-31T05:20:26.035Z level=INFO msg="america 5 日目: 115 人"
time=2026-10-31T05:20:35.159Z level=INFO msg="america 最終日（グランプリ）: 115 人"
time=2026-10-31T05:20:35.348Z level=INFO msg="europe 最終日（グランプリ）: 220 人"
time=2026-10-31T05:20:35.480Z level=INFO msg="asia 最終日（グランプリ）: 1085 人"
time=2026-10-31T05:20:43.822Z level=INFO msg="グランプリ: アジア 468 人, ヨーロッパ 102 人, アメリカ 61 人"
{
  "pass": true,
  "score": 134342,
  "messages": [
    "miamar_13166: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "horere_14217: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ordan-12252: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "suno_13767: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kalia12219: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yaha.13627: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "neliano-11369: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mahako_11810: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chisol13841: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "liatortor_13761: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chisa_13420: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kini9944: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nerileo10906: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "reru10105: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "furoor-12894: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yonuan.11474: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rosolivy_13286: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hesol12852: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "miul_12553: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chisazoe11527: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.7%	1704	0	0	GET	/api/notifications	14607.729	8.573	15.647	15.655	151
18.3%	2527	125	0	POST	/api/posts	7481.353	2.961	9.915	13.412	617
9.6%	5878	74	0	POST	/api/friend-requests	3923.791	0.668	5.187	9.620	114
7.9%	16976	545	0	GET	^/api/images/[^/]+	3234.070	0.191	3.237	11.405	127572
6.9%	745	127	0	PUT	/api/me/profile-image	2812.774	3.776	11.686	13.083	3
5.1%	6021	696	0	POST	/api/auth/login	2102.553	0.349	4.169	10.272	218
3.8%	2392	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1537.747	0.643	4.703	11.449	116
2.2%	577	32	0	POST	^/api/posts/[^/]+/realmojis$	888.690	1.540	9.543	12.417	462
1.7%	1110	1	0	POST	/api/auth/register	681.510	0.614	4.221	10.144	227
1.2%	750	0	0	PUT	^/api/follows/[^/]+$	508.902	0.679	4.508	9.630	0
1.2%	4082	0	0	GET	/api/memories	495.811	0.121	0.733	1.046	11905
1.0%	533	0	0	PATCH	/api/me	416.439	0.781	4.967	8.380	275
1.0%	514	13	0	POST	^/api/posts/[^/]+/comments$	398.321	0.775	5.697	7.088	341
0.9%	13326	0	0	GET	^/api/feeds/[^/]+	374.836	0.028	0.260	1.741	2470
0.6%	863	0	0	POST	/api/auth/logout	231.649	0.268	3.857	5.789	0
0.5%	311	0	0	GET	/api/me/export	220.578	0.709	3.651	5.016	1882996
0.4%	228	9	0	PATCH	^/api/posts/[^/]+$	171.965	0.754	4.316	7.007	626
0.4%	3729	0	0	GET	/api/posting-window/current	163.295	0.044	0.321	0.659	237
0.3%	106	0	0	DELETE	^/api/posts/[^/]+$	140.854	1.329	4.357	7.059	0
0.2%	2829	0	0	GET	/api/users	82.866	0.029	0.181	0.640	236
0.1%	1935	71	0	GET	^/api/posts/me/current$	58.990	0.030	0.194	0.612	626
0.1%	814	0	0	GET	^/api/memories/[^/]+	56.985	0.070	0.586	0.949	723
0.1%	96	0	0	DELETE	^/api/realmojis/[^/]+$	53.880	0.561	4.675	4.675	0
0.1%	2037	25	0	GET	^/api/posts/[^/]+$	51.763	0.025	0.150	1.461	609
0.1%	1724	0	0	GET	^/api/relationships/[^/]+	49.049	0.028	0.275	0.805	958
0.1%	98	0	0	DELETE	^/api/comments/[^/]+$	37.523	0.383	6.299	6.299	0
0.1%	1046	0	0	GET	^/api/users/[^/]+$	23.649	0.023	0.160	0.674	186
0.0%	1	0	0	POST	/api/initialize	19.110	19.110	19.110	19.110	90
0.0%	878	0	0	GET	^/api/posts/[^/]+/realmojis$	19.029	0.022	0.140	0.623	1016
0.0%	703	0	0	GET	^/api/posts/[^/]+/comments$	17.070	0.024	0.198	0.659	1754
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.1%	466	0	0	GET	/api/notifications	3840.171	8.241	12.372	13.200	165
11.4%	617	14	0	POST	/api/posts	1117.964	1.812	6.599	10.397	629
9.0%	1162	9	0	POST	/api/friend-requests	888.433	0.765	4.652	9.722	114
8.9%	2316	21	0	GET	^/api/images/[^/]+	879.250	0.380	2.682	4.043	199731
5.1%	672	0	0	POST	/api/auth/login	498.919	0.742	4.993	9.201	246
4.6%	2107	0	0	GET	^/api/feeds/[^/]+	448.470	0.213	0.773	1.487	2376
2.8%	330	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	272.436	0.826	4.804	5.754	116
2.6%	862	0	0	GET	/api/posting-window/current	258.475	0.300	0.816	1.122	239
1.8%	703	0	0	GET	/api/memories	181.543	0.258	1.096	1.545	11360
1.8%	113	0	0	PUT	/api/me/profile-image	176.382	1.561	8.649	8.763	0
1.6%	214	0	0	POST	/api/auth/register	161.380	0.754	4.509	6.960	229
1.5%	70	16	0	POST	^/api/posts/[^/]+/realmojis$	144.296	2.061	8.502	8.502	366
1.3%	595	0	0	GET	/api/users	128.603	0.216	0.766	1.485	237
1.3%	141	9	0	POST	^/api/posts/[^/]+/comments$	126.787	0.899	6.042	6.268	328
1.2%	504	25	0	GET	^/api/posts/me/current$	118.953	0.236	0.790	0.999	612
0.9%	111	0	0	POST	/api/auth/logout	90.683	0.817	4.209	4.479	0
0.8%	138	0	0	PUT	^/api/follows/[^/]+$	77.130	0.559	4.435	4.538	0
0.6%	37	0	0	GET	/api/me/export	60.331	1.631	3.846	3.846	2295500
0.6%	294	0	0	GET	^/api/relationships/[^/]+	58.958	0.201	0.783	1.284	895
0.5%	287	22	0	GET	^/api/posts/[^/]+$	50.736	0.177	0.704	0.775	565
0.5%	96	0	0	PATCH	/api/me	48.988	0.510	3.600	3.600	268
0.5%	52	10	0	PATCH	^/api/posts/[^/]+$	46.543	0.895	3.499	3.499	520
0.3%	139	0	0	GET	^/api/users/[^/]+$	26.547	0.191	0.733	0.903	184
0.2%	24	0	0	DELETE	^/api/posts/[^/]+$	24.009	1.000	3.172	3.172	0
0.2%	119	1	0	GET	/api/me	21.406	0.180	0.421	0.606	299
0.2%	99	0	0	GET	^/api/posts/[^/]+/comments$	20.360	0.206	0.762	0.762	1622
0.2%	78	0	0	GET	^/api/memories/[^/]+	19.393	0.249	1.322	1.322	899
0.2%	94	0	0	GET	^/api/posts/[^/]+/realmojis$	16.453	0.175	1.390	1.390	759
0.1%	20	0	0	DELETE	^/api/comments/[^/]+$	10.340	0.517	3.781	3.781	0
0.1%	6	0	0	DELETE	^/api/realmojis/[^/]+$	5.290	0.882	2.848	2.848	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.9%	235	0	0	GET	/api/notifications	1725.421	7.342	9.330	9.330	164
10.5%	284	14	0	POST	/api/posts	488.758	1.721	7.106	7.870	613
9.9%	597	11	0	POST	/api/friend-requests	464.840	0.779	4.882	6.543	113
8.8%	1262	5	0	GET	^/api/images/[^/]+	411.657	0.326	2.047	2.881	164484
6.1%	390	0	0	POST	/api/auth/login	286.528	0.735	4.975	8.610	249
4.2%	189	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	194.344	1.028	6.092	7.272	115
3.9%	918	0	0	GET	^/api/feeds/[^/]+	183.438	0.200	0.609	1.055	2869
2.3%	403	0	0	GET	/api/posting-window/current	105.470	0.262	0.641	0.765	240
2.1%	362	0	0	GET	/api/memories	96.038	0.265	1.101	1.441	14195
1.9%	55	0	0	PUT	/api/me/profile-image	87.297	1.587	12.205	12.205	0
1.8%	34	5	0	POST	^/api/posts/[^/]+/realmojis$	85.929	2.527	8.584	8.584	412
1.4%	57	0	0	PATCH	/api/me	65.876	1.156	6.245	6.245	284
1.3%	110	0	0	POST	/api/auth/register	59.144	0.538	3.651	3.788	229
1.2%	282	0	0	GET	/api/users	57.313	0.203	0.635	1.100	233
1.2%	237	17	0	GET	^/api/posts/me/current$	55.673	0.235	0.822	2.246	600
1.0%	64	0	0	POST	/api/auth/logout	45.823	0.716	4.685	4.685	0
1.0%	25	0	0	GET	/api/me/export	45.325	1.813	7.089	7.089	2613742
0.8%	59	0	0	PUT	^/api/follows/[^/]+$	38.771	0.657	7.029	7.029	0
0.6%	141	0	0	GET	^/api/relationships/[^/]+	28.244	0.200	1.088	1.107	1131
0.6%	32	0	0	POST	^/api/posts/[^/]+/comments$	27.311	0.853	5.210	5.210	331
0.5%	123	7	0	GET	^/api/posts/[^/]+$	21.690	0.176	0.569	0.827	580
0.4%	19	1	0	PATCH	^/api/posts/[^/]+$	18.960	0.998	2.913	2.913	600
0.4%	11	0	0	DELETE	^/api/posts/[^/]+$	16.627	1.512	3.537	3.537	0
0.3%	61	0	0	GET	^/api/memories/[^/]+	15.645	0.256	0.899	0.899	1267
0.3%	70	0	0	GET	^/api/users/[^/]+$	12.144	0.173	0.554	0.554	187
0.2%	58	0	0	GET	/api/me	11.010	0.190	0.924	0.924	301
0.2%	50	0	0	GET	^/api/posts/[^/]+/realmojis$	7.958	0.159	0.546	0.546	833
0.1%	47	0	0	GET	^/api/posts/[^/]+/comments$	7.001	0.149	0.531	0.531	1661
0.1%	4	0	0	DELETE	^/api/comments/[^/]+$	2.990	0.748	2.259	2.259	0
0.0%	2	0	0	DELETE	^/api/follows/[^/]+$	2.043	1.022	1.941	1.941	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 85895 クエリ / 82 種類 / 合計 139.576s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	62.9%	373	87.746	235.25	850.82	1236.2	126.3	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.0%	11681	19.605	1.68	19.62	9.5	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
3	5.4%	3067	7.571	2.47	19.67	103.7	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $2) AS is_profile, EXISTS (SELECT $3 FROM posts p WHERE p.image_id…
4	4.8%	2	6.656	3328.23	6498.04	233612.5	2846.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
5	3.3%	3279	4.626	1.41	13.47	24.9	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.6%	1	2.168	2168.26	2168.26	31433.0	1199905.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	1.3%	1	1.828	1827.53	1827.53	24949.0	628397.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	1.1%	628	1.558	2.48	11.63	28.6	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.9%	7533	1.314	0.17	5.49	19.8	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	0.6%	665	0.878	1.32	8.71	22.2	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.6%	1	0.792	792.42	792.42	13091.0	530281.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
12	0.5%	6387	0.664	0.10	5.95	11.2	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
13	0.4%	2811	0.548	0.19	4.49	20.5	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
14	0.4%	4693	0.513	0.11	6.26	8.7	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$4::bytea,$5,$3)
15	0.2%	2	0.266	132.81	256.94	16781.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT data FROM image_objects WHERE id = $1 AND created_at = $2;
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $2) AS is_profile, EXISTS (SELECT $3 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $4 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 94d7b5d4548b5b078dfb5dcefb31f72c2147d132
Type: cpu
Time: 2026-10-31 14:19:20 JST
Duration: 60.02s, Total samples = 39.92s (66.52%)
Showing nodes accounting for 28.31s, 70.92% of 39.92s total
Dropped 1032 nodes (cum <= 0.20s)
Showing top 40 nodes out of 306
      flat  flat%   sum%        cum   cum%
     4.81s 12.05% 12.05%      4.81s 12.05%  crypto/internal/fips140/sha256.blockAVX2
     4.75s 11.90% 23.95%      4.75s 11.90%  internal/runtime/syscall/linux.Syscall6
     1.75s  4.38% 28.33%      5.62s 14.08%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.59s  3.98% 32.31%      1.59s  3.98%  runtime.memclrNoHeapPointers
     1.52s  3.81% 36.12%      1.52s  3.81%  runtime.memmove
     1.39s  3.48% 39.60%      5.60s 14.03%  runtime.scanObject
     1.33s  3.33% 42.94%      1.33s  3.33%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.30s  3.26% 46.19%      1.30s  3.26%  aeshashbody
     1.27s  3.18% 49.37%      1.66s  4.16%  runtime.findObject
     1.24s  3.11% 52.48%      2.24s  5.61%  runtime.tryDeferToSpanScan
     0.73s  1.83% 54.31%      0.73s  1.83%  runtime.spanClass.sizeclass (inline)
     0.53s  1.33% 55.64%      0.53s  1.33%  memeqbody
     0.52s  1.30% 56.94%      2.27s  5.69%  runtime.mapaccess1_faststr
     0.47s  1.18% 58.12%      0.47s  1.18%  runtime.nextFreeFast (inline)
     0.34s  0.85% 58.97%      0.34s  0.85%  sync/atomic.(*Int32).Add
     0.32s   0.8% 59.77%      0.50s  1.25%  example.com/isureal-company/isureal.postLess
     0.28s   0.7% 60.47%      0.80s  2.00%  database/sql.convertAssignRows
     0.26s  0.65% 61.12%      1.76s  4.41%  runtime.mapassign_faststr
     0.25s  0.63% 61.75%      5.26s 13.18%  crypto/internal/fips140/sha256.(*Digest).Write
     0.24s   0.6% 62.35%      0.30s  0.75%  encoding/json.appendString[go.shape.string]
     0.24s   0.6% 62.95%      0.24s   0.6%  indexbody
     0.24s   0.6% 63.55%      0.25s  0.63%  runtime.(*mspan).writeHeapBitsSmall
     0.23s  0.58% 64.13%      0.23s  0.58%  runtime.futex
     0.22s  0.55% 64.68%      1.19s  2.98%  runtime.mallocgcSmallScanNoHeader
     0.21s  0.53% 65.21%      4.55s 11.40%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.19s  0.48% 65.68%      0.36s   0.9%  compress/flate.(*compressor).deflate
     0.19s  0.48% 66.16%      0.20s   0.5%  runtime.spanOf (inline)
     0.19s  0.48% 66.63%      0.22s  0.55%  runtime.typePointers.next
     0.17s  0.43% 67.06%      0.42s  1.05%  runtime.greyobject
     0.16s   0.4% 67.46%      3.06s  7.67%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.16s   0.4% 67.86%      1.23s  3.08%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.16s   0.4% 68.26%      0.20s   0.5%  runtime.heapArenaOf (inline)
     0.15s  0.38% 68.64%      4.81s 12.05%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.15s  0.38% 69.01%      4.61s 11.55%  runtime.mallocgc
     0.15s  0.38% 69.39%      0.40s  1.00%  runtime.scanObjectsSmall
     0.14s  0.35% 69.74%      0.24s   0.6%  internal/runtime/maps.(*Iter).Next
     0.13s  0.33% 70.07%      0.77s  1.93%  encoding/json.structEncoder.encode
     0.12s   0.3% 70.37%      0.54s  1.35%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
     0.11s  0.28% 70.64%      3.21s  8.04%  example.com/isureal-company/isureal.listMemories
     0.11s  0.28% 70.92%      0.70s  1.75%  runtime.scanSpan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 48.8% / 100% | 35% | 0.6% | 7 | 106 MB | CPU が飽和 |
| isu2 (2 cpu) | 15.6% / 70% | 0% | 0.7% | 5 | 1811 MB | run queue が積もっている |
| isu3 (2 cpu) | 13.0% / 73% | 0% | 0.6% | 5 | 1750 MB | run queue が積もっている |
| isu4 (2 cpu) | 41.3% / 100% | 14% | 18.0% | 25 | 105 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 80% / nginx 12% |
| isu2 | isureal 29% / nginx 1% |
| isu3 | isureal 24% / nginx 1% |
| isu4 | postgres 24% / isureal 19% |

## Diff from previous run

スコア: 123551 → 134342 (+10791)

`247167c..a8128f2`

```
 webapp/go/export_api.go        |   2 +
 webapp/go/image_api.go         |  17 +-----
 webapp/go/image_file.go        | 132 +++++++++++++++++++++++++++++++++++++++++
 webapp/go/image_local.go       |  40 ++++++-------
 webapp/go/initialize.go        |   1 +
 webapp/go/jpegcheck/huffman.go |  27 ++++++++-
 webapp/go/jpegcheck/reader.go  |  32 ++++++----
 webapp/go/jpegcheck/scan.go    |   7 ++-
 webapp/go/post_api.go          |   5 +-
 webapp/go/profile_image_api.go |   5 +-
 webapp/go/realmoji_api.go      |   5 +-
 webapp/go/server.go            |   1 +
 12 files changed, 216 insertions(+), 58 deletions(-)
```
