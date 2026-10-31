# run 20261031-121047-249be25

## Summary

| | |
|---|---|
| score | 70128 |
| status | pass |
| ref | origin/main |
| sha | 249be2509f96 |
| 所要時間 | 191s |
| 前回比 | +1782 (2.6%) vs 20261031-120609-9f29066 |
| ベスト比 | +1782 (2.6%) vs 20261031-120609-9f29066 = 68346 **ベスト更新** |
| note | main 9b2e015 compare for pair locks |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	43	118	326	528	527
europe	7	18	47	120	120	120
america	6	14	38	71	72	71
```

```
{"id":"7086","status":"done","portalStatus":"finished","passed":true,"score":70128,"raw":70128,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:12:39.891Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:12:42.857Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:12:45.856Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:12:48.880Z level=INFO msg="asia 2 日目: 43 人"
time=2026-10-31T03:12:51.856Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T03:12:54.856Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T03:12:57.882Z level=INFO msg="asia 3 日目: 118 人"
time=2026-10-31T03:12:59.756Z level=INFO msg="user moved" user=koeva.5500 from=asia to=america
time=2026-10-31T03:13:00.903Z level=INFO msg="europe 3 日目: 47 人"
time=2026-10-31T03:13:01.066Z level=INFO msg="user moved" user=mewase-5561 from=asia to=europe
time=2026-10-31T03:13:02.030Z level=INFO msg="user moved" user=nisu_5974 from=asia to=europe
time=2026-10-31T03:13:03.858Z level=INFO msg="america 3 日目: 38 人"
time=2026-10-31T03:13:06.886Z level=INFO msg="asia 4 日目: 326 人"
time=2026-10-31T03:13:09.894Z level=INFO msg="europe 4 日目: 120 人"
time=2026-10-31T03:13:12.980Z level=INFO msg="america 4 日目: 71 人"
time=2026-10-31T03:13:16.145Z level=INFO msg="user moved" user=miamo.7115 from=asia to=america
time=2026-10-31T03:13:16.169Z level=INFO msg="asia 5 日目: 528 人"
time=2026-10-31T03:13:18.948Z level=INFO msg="europe 5 日目: 120 人"
time=2026-10-31T03:13:21.940Z level=INFO msg="america 5 日目: 72 人"
time=2026-10-31T03:13:31.145Z level=INFO msg="asia 最終日（グランプリ）: 527 人"
time=2026-10-31T03:13:31.306Z level=INFO msg="europe 最終日（グランプリ）: 120 人"
time=2026-10-31T03:13:31.374Z level=INFO msg="america 最終日（グランプリ）: 71 人"
time=2026-10-31T03:13:39.872Z level=INFO msg="グランプリ: アジア 79 人, ヨーロッパ 11 人, アメリカ 12 人"
{
  "pass": true,
  "score": 70128,
  "messages": [
    "nihiru.5474: 友達申請を承認できません: POST /api/friend-requests/19ca07e8-52d9-42f5-9b8b-de84b767e406/accept: 応答がタイムアウトしました",
    "ulsuchi6362: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "marsol.6874: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "liana_7627: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "nuvin_6827: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "noavinke_7402: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "tsuhanoa7698: プロフィール画像を変えられません: PUT /api/me/profile-image: 応答がタイムアウトしました",
    "hano.6556: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "shisolsol.7856: 認定公開アカウントをフォローできません: PUT /api/follows/e393d46e-d6e8-4050-938c-91db8615168b: 応答がタイムアウトしました",
    "nonoaki-6272: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "torsesu5826: 投稿 e9c60ebf-bd47-4808-a9e1-2c02debec966 にコメントできません: POST /api/posts/e9c60ebf-bd47-4808-a9e1-2c02debec966/comments: 応答がタイムアウトしました",
    "kawahi-6988: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "yunoa-7004: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "rodaneva5636: 投稿 e9c60ebf-bd47-4808-a9e1-2c02debec966 に RealMoji を送れません: POST /api/posts/e9c60ebf-bd47-4808-a9e1-2c02debec966/realmojis: 応答がタイムアウトしました",
    "numaro_6597: 認定公開アカウントをフォローできません: PUT /api/follows/e393d46e-d6e8-4050-938c-91db8615168b: 応答がタイムアウトしました",
    "tormax6535: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "zoesa_7727: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "sewa-5033: 投稿 82354180-d390-4f9f-9fa1-0af16c42daa0 にコメントできません: POST /api/posts/82354180-d390-4f9f-9fa1-0af16c42daa0/comments: 応答がタイムアウトしました",
    "yomo-5660: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "elmumo5554: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
30.3%	873	0	0	GET	/api/notifications	7771.234	8.902	15.085	15.315	173
20.5%	2675	66	0	POST	/api/friend-requests	5264.445	1.968	26.894	31.938	113
11.7%	1105	15	0	POST	/api/posts	2987.471	2.704	10.701	18.354	638
7.8%	10289	584	0	GET	^/api/images/[^/]+	1992.703	0.194	2.644	10.664	95960
6.9%	335	0	0	PUT	^/api/follows/[^/]+$	1759.983	5.254	31.006	31.387	0
6.0%	4453	1172	0	POST	/api/auth/login	1531.982	0.344	6.169	27.997	188
3.8%	862	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	966.587	1.121	14.575	32.056	115
3.4%	330	56	0	POST	^/api/posts/[^/]+/realmojis$	864.438	2.620	28.605	31.241	396
2.5%	201	52	0	POST	^/api/posts/[^/]+/comments$	630.914	3.139	28.411	31.382	263
1.2%	2570	0	0	GET	/api/memories	300.355	0.117	0.514	2.750	11652
1.1%	548	1	0	POST	/api/auth/register	289.530	0.528	2.669	6.938	227
0.9%	400	112	0	PUT	/api/me/profile-image	226.690	0.567	3.303	31.913	7
0.8%	1602	0	0	GET	/api/posting-window/current	200.826	0.125	1.171	7.948	237
0.6%	210	0	0	GET	/api/me/export	156.374	0.745	2.939	10.519	1907884
0.5%	6167	0	0	GET	^/api/feeds/[^/]+	137.902	0.022	0.144	1.601	1886
0.5%	559	0	0	POST	/api/auth/logout	128.364	0.230	2.211	10.090	0
0.3%	252	0	0	PATCH	/api/me	79.314	0.315	1.925	2.771	272
0.2%	5	0	0	DELETE	^/api/follows/[^/]+$	55.552	11.110	23.148	23.148	0
0.2%	1301	0	0	GET	/api/users	42.575	0.033	0.422	1.425	233
0.2%	97	1	0	PATCH	^/api/posts/[^/]+$	39.728	0.410	6.814	6.814	645
0.1%	528	0	0	GET	^/api/memories/[^/]+	32.313	0.061	0.327	2.225	569
0.1%	856	7	0	GET	^/api/posts/me/current$	29.671	0.035	0.301	0.706	641
0.1%	37	0	0	DELETE	^/api/posts/[^/]+$	23.976	0.648	2.708	2.708	0
0.1%	17	0	0	PUT	^/api/blocks/[^/]+$	16.395	0.964	6.572	6.572	0
0.1%	1	0	0	POST	/api/initialize	15.314	15.314	15.314	15.314	90
0.1%	886	10	0	GET	^/api/posts/[^/]+$	15.261	0.017	0.106	0.175	613
0.1%	47	0	0	DELETE	^/api/realmojis/[^/]+$	14.618	0.311	1.656	1.656	0
0.1%	910	0	0	GET	^/api/relationships/[^/]+	14.563	0.016	0.127	0.226	964
0.0%	14	0	0	PUT	^/api/admin/users/[^/]+/certification$	12.752	0.911	11.057	11.057	0
0.0%	482	0	0	GET	^/api/users/[^/]+$	7.944	0.016	0.108	0.422	199
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.6%	266	0	0	GET	/api/notifications	2248.220	8.452	12.429	12.429	180
17.5%	680	18	0	POST	/api/friend-requests	1167.931	1.718	27.051	29.584	113
13.2%	317	6	0	POST	/api/posts	881.120	2.780	10.361	11.656	640
5.2%	566	0	0	POST	/api/auth/login	348.137	0.615	7.167	9.516	250
4.2%	232	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	282.095	1.216	13.268	29.388	116
4.2%	1875	7	0	GET	^/api/images/[^/]+	280.578	0.150	0.950	1.209	133471
3.9%	56	0	0	PUT	^/api/follows/[^/]+$	263.597	4.707	29.231	29.231	0
3.5%	119	14	0	POST	^/api/posts/[^/]+/realmojis$	233.950	1.966	15.504	29.389	411
3.2%	1446	0	0	GET	^/api/feeds/[^/]+	215.824	0.149	0.291	1.835	2329
1.7%	540	0	0	GET	/api/memories	112.892	0.209	0.688	1.763	13364
1.5%	425	0	0	GET	/api/posting-window/current	99.989	0.235	1.625	2.106	239
1.5%	42	7	0	POST	^/api/posts/[^/]+/comments$	98.916	2.355	16.777	16.777	259
0.7%	84	0	0	POST	/api/auth/logout	49.905	0.594	8.054	8.054	0
0.7%	33	0	0	GET	/api/me/export	49.713	1.506	6.377	6.377	2327406
0.7%	332	0	0	GET	/api/users	49.442	0.149	0.315	0.689	232
0.6%	71	0	0	PUT	/api/me/profile-image	42.761	0.602	11.494	11.494	0
0.6%	244	13	0	GET	^/api/posts/me/current$	37.815	0.155	0.328	1.034	615
0.5%	116	0	0	POST	/api/auth/register	34.428	0.297	0.837	1.467	228
0.5%	220	7	0	GET	^/api/posts/[^/]+$	32.682	0.149	0.274	0.294	603
0.5%	222	0	0	GET	^/api/relationships/[^/]+	31.065	0.140	0.300	0.618	1042
0.3%	39	4	0	PATCH	^/api/posts/[^/]+$	21.213	0.544	4.528	4.528	584
0.3%	93	0	0	GET	^/api/memories/[^/]+	17.800	0.191	1.275	1.275	794
0.2%	107	0	0	GET	^/api/users/[^/]+$	15.828	0.148	0.248	0.281	196
0.2%	106	1	0	GET	^/api/posts/[^/]+/realmojis$	15.539	0.147	0.287	0.414	889
0.2%	52	0	0	PATCH	/api/me	14.632	0.281	1.175	1.175	277
0.2%	76	1	0	GET	/api/me	10.629	0.140	0.281	0.281	297
0.1%	53	0	0	GET	^/api/posts/[^/]+/comments$	7.651	0.144	0.277	0.277	681
0.1%	5	0	0	PUT	^/api/blocks/[^/]+$	6.416	1.283	4.812	4.812	0
0.1%	16	0	0	DELETE	^/api/realmojis/[^/]+$	5.895	0.368	0.786	0.786	0
0.1%	7	0	0	DELETE	^/api/posts/[^/]+$	3.508	0.501	1.103	1.103	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.5%	171	1	0	GET	/api/notifications	1305.090	7.632	9.478	9.478	178
17.4%	367	4	0	POST	/api/friend-requests	677.859	1.847	26.813	30.919	114
12.5%	198	1	0	POST	/api/posts	487.559	2.462	8.922	10.420	648
5.1%	320	0	0	POST	/api/auth/login	200.586	0.627	8.446	9.529	251
4.8%	1195	15	0	GET	^/api/images/[^/]+	186.856	0.156	0.948	1.617	149165
4.1%	116	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	158.505	1.366	12.208	12.637	115
3.4%	907	0	0	GET	^/api/feeds/[^/]+	132.823	0.146	0.296	1.307	2165
3.3%	44	0	0	PUT	^/api/follows/[^/]+$	126.928	2.885	20.014	20.014	0
2.4%	29	7	0	POST	^/api/posts/[^/]+/realmojis$	94.240	3.250	15.071	15.071	381
2.0%	345	0	0	GET	/api/memories	79.260	0.230	0.836	6.088	14560
1.6%	50	8	0	POST	^/api/posts/[^/]+/comments$	63.704	1.274	15.451	15.451	265
1.6%	267	0	0	GET	/api/posting-window/current	62.772	0.235	1.613	2.120	240
1.1%	64	0	0	POST	/api/auth/logout	41.151	0.643	4.519	4.519	0
0.9%	20	0	0	GET	/api/me/export	35.609	1.780	3.788	3.788	2667392
0.7%	167	0	0	GET	/api/users	26.075	0.156	0.604	0.645	240
0.7%	165	6	0	GET	^/api/posts/me/current$	25.407	0.154	0.374	1.065	628
0.6%	160	6	0	GET	^/api/posts/[^/]+$	23.193	0.145	0.315	0.339	600
0.5%	33	0	0	PATCH	/api/me	19.098	0.579	3.698	3.698	259
0.5%	3	0	0	DELETE	^/api/follows/[^/]+$	18.111	6.037	11.368	11.368	0
0.5%	33	0	0	PUT	/api/me/profile-image	17.752	0.538	2.405	2.405	0
0.4%	128	0	0	GET	^/api/relationships/[^/]+	17.357	0.136	0.276	0.324	880
0.4%	69	0	0	POST	/api/auth/register	17.312	0.251	0.702	0.702	229
0.3%	15	0	0	DELETE	^/api/posts/[^/]+$	13.142	0.876	2.850	2.850	0
0.3%	8	0	0	DELETE	^/api/comments/[^/]+$	12.470	1.559	9.582	9.582	0
0.3%	79	0	0	GET	^/api/users/[^/]+$	11.221	0.142	0.409	0.409	191
0.2%	19	1	0	PATCH	^/api/posts/[^/]+$	9.514	0.501	2.344	2.344	626
0.2%	71	0	0	GET	^/api/posts/[^/]+/comments$	9.396	0.132	0.261	0.261	687
0.2%	59	0	0	GET	^/api/posts/[^/]+/realmojis$	7.803	0.132	0.274	0.274	920
0.2%	34	0	0	GET	^/api/memories/[^/]+	6.510	0.191	0.377	0.377	1115
0.1%	37	1	0	GET	/api/me	5.390	0.146	0.258	0.258	291
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 49550 クエリ / 81 種類 / 合計 275.643s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	50.9%	2391	140.435	58.73	556.97	613.6	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	30.4%	263	83.745	318.42	1145.27	6855.0	129.0	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	6.4%	2545	17.772	6.98	275.51	114.1	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	5.8%	7560	16.109	2.13	41.42	10.3	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.3%	1598	3.591	2.25	33.82	26.2	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.8%	1	2.191	2191.39	2191.39	29737.0	1199731.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.6%	1	1.760	1759.94	1759.94	20756.0	626734.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.5%	2	1.384	692.20	1375.94	16781.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
9	0.5%	401	1.296	3.23	14.27	30.3	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.4%	1	1.093	1092.50	1092.50	2647.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
11	0.3%	1	0.802	802.02	802.02	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.3%	3623	0.755	0.21	14.99	20.3	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
13	0.2%	1	0.677	677.33	677.33	11528.0	529870.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
14	0.2%	2	0.618	309.06	473.26	226008.5	2658.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
15	0.2%	4167	0.475	0.11	10.48	11.2	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: f7c9c67524bdcdb53f296b47113919c97dcd9a6d
Type: cpu
Time: 2026-10-31 12:12:17 JST
Duration: 60.17s, Total samples = 44.89s (74.61%)
Showing nodes accounting for 33.74s, 75.16% of 44.89s total
Dropped 976 nodes (cum <= 0.22s)
Showing top 40 nodes out of 280
      flat  flat%   sum%        cum   cum%
     3.20s  7.13%  7.13%      3.20s  7.13%  crypto/internal/fips140/sha256.blockAVX2
     2.94s  6.55% 13.68%      2.94s  6.55%  internal/runtime/syscall/linux.Syscall6
     2.20s  4.90% 18.58%      7.14s 15.91%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     2.12s  4.72% 23.30%      7.52s 16.75%  runtime.scanObject
     2.07s  4.61% 27.91%      2.07s  4.61%  runtime.memclrNoHeapPointers
     1.96s  4.37% 32.28%      1.96s  4.37%  runtime.memmove
     1.91s  4.25% 36.53%      3.22s  7.17%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.78s  3.97% 40.50%      2.32s  5.17%  runtime.findObject
     1.66s  3.70% 44.20%      5.71s 12.72%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.62s  3.61% 47.81%      1.72s  3.83%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     1.53s  3.41% 51.21%      2.53s  5.64%  runtime.tryDeferToSpanScan
     1.47s  3.27% 54.49%      1.47s  3.27%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.17s  2.61% 57.10%      1.38s  3.07%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     1.07s  2.38% 59.48%      1.07s  2.38%  aeshashbody
     0.72s  1.60% 61.08%      0.72s  1.60%  example.com/isureal-company/isureal.(*UserRow).active
     0.71s  1.58% 62.66%      0.71s  1.58%  runtime.spanClass.sizeclass (inline)
     0.55s  1.23% 63.89%      1.85s  4.12%  runtime.mapaccess1_faststr
     0.52s  1.16% 65.05%      0.52s  1.16%  memeqbody
     0.34s  0.76% 65.81%      0.34s  0.76%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.32s  0.71% 66.52%      0.60s  1.34%  example.com/isureal-company/isureal.postLess
     0.30s  0.67% 67.19%      0.64s  1.43%  database/sql.convertAssignRows
     0.27s   0.6% 67.79%      0.27s   0.6%  runtime.typePointers.nextFast (inline)
     0.25s  0.56% 68.34%      3.66s  8.15%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.25s  0.56% 68.90%      0.25s  0.56%  indexbody
     0.23s  0.51% 69.41%      0.23s  0.51%  runtime.nextFreeFast (inline)
     0.22s  0.49% 69.90%      0.51s  1.14%  runtime.greyobject
     0.22s  0.49% 70.39%      0.23s  0.51%  runtime.typePointers.next
     0.20s  0.45% 70.84%      0.25s  0.56%  runtime.(*mspan).writeHeapBitsSmall
     0.20s  0.45% 71.29%      4.72s 10.51%  runtime.mallocgc
     0.20s  0.45% 71.73%      1.79s  3.99%  runtime.mapassign_faststr
     0.20s  0.45% 72.18%      0.27s   0.6%  runtime.spanOf (inline)
     0.19s  0.42% 72.60%      0.96s  2.14%  runtime.mallocgcSmallScanNoHeader
     0.17s  0.38% 72.98%      0.29s  0.65%  compress/flate.(*compressor).deflate
     0.17s  0.38% 73.36%      3.45s  7.69%  crypto/internal/fips140/sha256.(*Digest).Write
     0.15s  0.33% 73.69%      1.11s  2.47%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.14s  0.31% 74.00%      0.45s  1.00%  runtime.scanObjectsSmall
     0.13s  0.29% 74.29%      4.11s  9.16%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.13s  0.29% 74.58%      3.85s  8.58%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.13s  0.29% 74.87%      1.40s  3.12%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.13s  0.29% 75.16%      3.04s  6.77%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 51.8% / 100% | 36% | 0.0% | 7 | 236 MB | CPU が飽和 |
| isu2 (2 cpu) | 11.9% / 30% | 0% | 0.8% | 3 | 2812 MB |  |
| isu3 (2 cpu) | 9.4% / 39% | 0% | 0.7% | 3 | 2928 MB |  |
| isu4 (2 cpu) | 51.5% / 100% | 20% | 17.9% | 38 | 113 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 85% / nginx 12% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 21% / nginx 1% |
| isu3 | isureal 16% / nginx 1% |
| isu4 | postgres 50% / isureal 14% |

## Diff from previous run

スコア: 68346 → 70128 (+1782)

`9f29066..249be25`

```
 webapp/go/account_deletion_api.go |   2 +-
 webapp/go/admin_user_api.go       |   2 +-
 webapp/go/comment_api.go          |   2 +-
 webapp/go/realmoji_api.go         |   2 +-
 webapp/go/relationship.go         |  10 +---
 webapp/go/store.go                | 100 ++++++--------------------------------
 6 files changed, 22 insertions(+), 96 deletions(-)
```
