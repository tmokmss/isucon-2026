# run 20261031-160103-8ff9b1d

## Summary

| | |
|---|---|
| score | 182075 |
| status | pass |
| ref | export-initial-images-split |
| sha | 8ff9b1dbf6c3 |
| 所要時間 | 468s |
| 前回比 | -1256 (-0.7%) vs 20261031-152137-8c13252 |
| ベスト比 | -9741 (-5.1%) vs 20261031-151339-5a230a4 = 191816 |
| note | export: half of initial images from asia1 init/, parallel with PG |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	52	153	447	1023	932
europe	7	21	63	183	499	499
america	6	18	57	165	359	359
```

```
{"id":"26322","status":"done","portalStatus":"finished","passed":true,"score":182075,"raw":182075,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:07:06.994Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:07:09.988Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:07:12.989Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:07:16.024Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T07:07:16.473Z level=INFO msg="user moved" user=makuleo_5175 from=asia to=europe
time=2026-10-31T07:07:17.554Z level=INFO msg="user moved" user=makuleo_5175 from=europe to=asia
time=2026-10-31T07:07:18.990Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T07:07:21.990Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T07:07:25.000Z level=INFO msg="asia 3 日目: 153 人"
time=2026-10-31T07:07:25.497Z level=INFO msg="user moved" user=toke.5461 from=europe to=asia
time=2026-10-31T07:07:25.773Z level=INFO msg="user moved" user=ivymax-5678 from=asia to=europe
time=2026-10-31T07:07:26.506Z level=INFO msg="user moved" user=ivymax-5678 from=europe to=asia
time=2026-10-31T07:07:26.535Z level=INFO msg="user moved" user=toke.5461 from=asia to=america
time=2026-10-31T07:07:26.861Z level=INFO msg="user moved" user=zoehe_5847 from=asia to=america
time=2026-10-31T07:07:27.912Z level=INFO msg="user moved" user=vinnoa6090 from=asia to=europe
time=2026-10-31T07:07:28.004Z level=INFO msg="europe 3 日目: 63 人"
time=2026-10-31T07:07:28.402Z level=INFO msg="user moved" user=futo-6164 from=asia to=europe
time=2026-10-31T07:07:29.379Z level=INFO msg="user moved" user=zoehe_5847 from=america to=asia
time=2026-10-31T07:07:30.989Z level=INFO msg="america 3 日目: 57 人"
time=2026-10-31T07:07:34.117Z level=INFO msg="asia 4 日目: 447 人"
time=2026-10-31T07:07:36.564Z level=INFO msg="user moved" user=ivymax-5678 from=asia to=europe
time=2026-10-31T07:07:37.169Z level=INFO msg="europe 4 日目: 183 人"
time=2026-10-31T07:07:38.994Z level=INFO msg="user moved" user=futo-6164 from=europe to=america
time=2026-10-31T07:07:39.485Z level=INFO msg="user moved" user=satsu_6752 from=america to=asia
time=2026-10-31T07:07:39.567Z level=INFO msg="user moved" user=vinnoa6090 from=europe to=america
time=2026-10-31T07:07:41.474Z level=INFO msg="user moved" user=hevinyo.6765 from=america to=europe
time=2026-10-31T07:07:41.887Z level=INFO msg="america 4 日目: 165 人"
time=2026-10-31T07:07:42.450Z level=INFO msg="user moved" user=leokafu8468 from=asia to=america
time=2026-10-31T07:07:42.539Z level=INFO msg="user moved" user=anevanoa_9642 from=europe to=america
time=2026-10-31T07:07:43.425Z level=INFO msg="asia 5 日目: 1023 人"
time=2026-10-31T07:07:44.459Z level=INFO msg="user moved" user=hora_9595 from=europe to=america
time=2026-10-31T07:07:45.321Z level=INFO msg="user moved" user=kikase_9812 from=america to=europe
time=2026-10-31T07:07:47.161Z level=INFO msg="user moved" user=hevinyo.6765 from=europe to=asia
time=2026-10-31T07:07:49.436Z level=INFO msg="america 5 日目: 359 人"
time=2026-10-31T07:07:50.231Z level=INFO msg="europe 5 日目: 499 人"
time=2026-10-31T07:07:58.325Z level=INFO msg="europe 最終日（グランプリ）: 499 人"
time=2026-10-31T07:07:58.557Z level=INFO msg="asia 最終日（グランプリ）: 932 人"
time=2026-10-31T07:07:59.243Z level=INFO msg="america 最終日（グランプリ）: 359 人"
time=2026-10-31T07:08:07.115Z level=INFO msg="グランプリ: アジア 387 人, ヨーロッパ 317 人, アメリカ 156 人"
{
  "pass": true,
  "score": 182075,
  "messages": [
    "rawa_8397: ログインし直せません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "yumarme10550: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "danyote10591: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noketo-9031: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mimarsol.8896: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "tenoeva_8813: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ululre_10695: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "xiangrikui_shousi_40: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "ting_huang_2867: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "teuripeul_yeongu_2601: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "hina_nakamura_939: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "elnishi_10143: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yoivyhe_10616: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "anne.7609: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "miaista7659: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "benta.7316: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "ren_watanabe_2043: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "marneno.10487: 認定公開アカウントをフォローできません: PUT /api/follows/09ba3816-6c71-4561-b4fa-33a11bedb73d: 通信に失敗しました: Put \"http://52.69.202.109:8080/api/follows/09ba3816-6c71-4561-b4fa-33a11bedb73d\": EOF",
    "lei_zhao_1695: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "torimia-10167: 友達申請を送れません: POST /api/friend-requests: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
48.0%	1567	0	0	GET	/api/notifications	13881.563	8.859	15.757	15.775	171
11.2%	2290	128	0	POST	/api/posts	3230.033	1.410	5.934	66.475	616
7.5%	16613	423	0	GET	^/api/images/[^/]+	2182.290	0.131	1.458	4.848	101719
6.3%	4974	54	0	POST	/api/friend-requests	1818.334	0.366	2.559	4.187	114
3.7%	8324	0	0	GET	^/api/feeds/[^/]+	1080.261	0.130	1.109	1.759	2371
3.4%	4096	0	0	GET	/api/memories	983.626	0.240	2.136	4.023	12926
3.0%	6133	638	0	POST	/api/auth/login	875.955	0.143	2.285	3.980	221
2.8%	667	135	0	PUT	/api/me/profile-image	795.646	1.193	14.019	15.965	3
2.6%	2123	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	737.692	0.347	2.585	3.138	116
2.0%	3324	0	0	GET	/api/posting-window/current	567.110	0.171	1.170	1.759	237
1.2%	2438	0	0	GET	/api/users	336.536	0.138	1.151	1.759	243
1.1%	343	0	0	GET	/api/me/export	317.587	0.926	5.929	6.332	1848591
0.9%	1744	40	0	GET	^/api/posts/me/current$	272.895	0.156	0.948	1.759	637
0.9%	356	14	0	POST	^/api/posts/[^/]+/realmojis$	265.706	0.746	5.437	5.686	455
0.7%	516	1	0	POST	^/api/posts/[^/]+/comments$	200.653	0.389	4.293	4.670	348
0.6%	880	0	0	POST	/api/auth/logout	175.019	0.199	2.544	4.002	0
0.6%	950	1	0	POST	/api/auth/register	165.827	0.175	0.693	2.756	227
0.5%	612	0	0	PUT	^/api/follows/[^/]+$	157.776	0.258	2.511	2.804	0
0.5%	1111	0	0	GET	^/api/relationships/[^/]+	131.893	0.119	1.251	2.364	1254
0.4%	1054	55	0	GET	^/api/posts/[^/]+$	115.908	0.110	1.026	1.458	593
0.4%	432	0	0	PATCH	/api/me	104.215	0.241	2.386	2.804	277
0.3%	889	0	0	GET	^/api/memories/[^/]+	89.255	0.100	1.151	2.136	598
0.3%	184	21	0	PATCH	^/api/posts/[^/]+$	88.552	0.481	2.869	4.018	573
0.3%	83	0	0	DELETE	^/api/posts/[^/]+$	77.873	0.938	3.119	3.119	0
0.2%	518	0	0	GET	^/api/users/[^/]+$	65.157	0.126	1.151	1.759	195
0.2%	542	2	0	GET	/api/me	49.935	0.092	0.602	1.151	302
0.1%	356	0	0	GET	^/api/posts/[^/]+/realmojis$	37.773	0.106	0.883	1.759	1286
0.1%	388	0	0	GET	^/api/posts/[^/]+/comments$	33.513	0.086	0.768	1.759	2147
0.1%	48	0	0	DELETE	^/api/realmojis/[^/]+$	23.122	0.482	2.737	2.737	0
0.1%	1	0	0	POST	/api/initialize	15.593	15.593	15.593	15.593	90
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.9%	755	0	0	GET	/api/notifications	4919.820	6.516	13.029	13.032	161
9.6%	1087	45	0	POST	/api/posts	1054.087	0.970	5.654	7.279	619
7.4%	2549	20	0	POST	/api/friend-requests	814.076	0.319	2.779	3.460	114
6.0%	5543	39	0	GET	^/api/images/[^/]+	657.632	0.119	1.340	3.096	140037
5.2%	3951	0	0	GET	^/api/feeds/[^/]+	575.081	0.146	1.284	2.380	2334
4.9%	1811	0	0	POST	/api/auth/login	532.787	0.294	2.681	3.214	247
4.2%	498	0	0	POST	/api/auth/register	459.931	0.924	2.801	2.861	229
2.3%	1665	0	0	GET	/api/posting-window/current	248.841	0.149	1.102	1.539	239
2.2%	1472	0	0	GET	/api/memories	239.836	0.163	1.950	3.682	14328
1.9%	1264	0	0	GET	/api/users	209.902	0.166	1.264	1.826	244
1.5%	260	0	0	PUT	/api/me/profile-image	166.601	0.641	3.020	3.148	0
1.4%	889	72	0	GET	^/api/posts/me/current$	152.552	0.172	1.296	1.709	595
1.1%	445	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	125.892	0.283	2.679	2.995	116
0.9%	246	0	0	PATCH	/api/me	100.879	0.410	2.663	2.718	272
0.9%	193	2	0	POST	^/api/posts/[^/]+/realmojis$	93.407	0.484	6.137	6.143	464
0.8%	313	0	0	PUT	^/api/follows/[^/]+$	93.100	0.297	2.558	2.664	0
0.7%	552	0	0	GET	^/api/relationships/[^/]+	78.016	0.141	1.223	1.601	618
0.7%	279	0	0	POST	/api/auth/logout	73.472	0.263	2.771	2.818	0
0.6%	569	19	0	GET	^/api/posts/[^/]+$	63.584	0.112	1.135	1.297	601
0.5%	107	13	0	PATCH	^/api/posts/[^/]+$	55.059	0.515	3.343	3.701	574
0.4%	162	1	0	POST	^/api/posts/[^/]+/comments$	44.771	0.276	2.817	3.343	343
0.3%	265	1	0	GET	/api/me	37.829	0.143	1.444	1.549	303
0.3%	43	0	0	DELETE	^/api/posts/[^/]+$	28.564	0.664	3.038	3.038	0
0.2%	245	0	0	GET	^/api/users/[^/]+$	24.901	0.102	0.793	1.017	193
0.2%	262	0	0	GET	^/api/memories/[^/]+	23.561	0.090	0.577	1.280	840
0.2%	211	1	0	GET	^/api/posts/[^/]+/realmojis$	21.265	0.101	0.975	1.196	1792
0.2%	93	0	0	GET	/api/me/export	19.774	0.213	2.735	2.735	2081156
0.1%	162	1	0	GET	^/api/posts/[^/]+/comments$	16.029	0.099	0.882	1.284	2776
0.1%	43	0	0	DELETE	^/api/realmojis/[^/]+$	13.826	0.322	1.063	1.063	0
0.1%	22	0	0	DELETE	^/api/comments/[^/]+$	6.954	0.316	3.119	3.119	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.7%	613	0	0	GET	/api/notifications	3808.761	6.213	10.679	10.680	168
10.0%	912	41	0	POST	/api/posts	892.536	0.979	6.813	7.330	624
7.8%	5244	33	0	GET	^/api/images/[^/]+	694.694	0.132	1.558	2.933	145652
6.7%	1981	21	0	POST	/api/friend-requests	593.152	0.299	2.630	3.357	114
6.6%	4372	0	0	GET	^/api/feeds/[^/]+	589.317	0.135	1.680	2.312	2465
5.2%	1540	0	0	POST	/api/auth/login	461.809	0.300	2.679	3.580	251
2.6%	355	0	0	POST	/api/auth/register	234.040	0.659	2.808	2.917	230
2.4%	1230	0	0	GET	/api/memories	217.293	0.177	2.256	3.746	18121
1.9%	1306	0	0	GET	/api/posting-window/current	167.326	0.128	0.881	1.765	240
1.5%	983	0	0	GET	/api/users	134.496	0.137	1.135	1.713	247
1.5%	702	29	0	GET	^/api/posts/me/current$	131.051	0.187	1.789	2.073	627
1.4%	499	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	126.067	0.253	2.793	3.458	116
1.2%	212	3	0	POST	^/api/posts/[^/]+/realmojis$	107.463	0.507	2.892	5.939	475
1.2%	198	0	0	PUT	/api/me/profile-image	106.726	0.539	3.090	3.482	0
0.9%	524	0	0	GET	^/api/relationships/[^/]+	79.958	0.153	1.700	2.532	678
0.8%	242	0	0	PUT	^/api/follows/[^/]+$	74.093	0.306	2.874	3.391	0
0.8%	176	0	0	PATCH	/api/me	69.462	0.395	2.791	2.797	277
0.7%	593	22	0	GET	^/api/posts/[^/]+$	65.164	0.110	0.898	1.659	605
0.7%	256	0	0	POST	/api/auth/logout	61.488	0.240	2.173	2.488	0
0.6%	87	1	0	PATCH	^/api/posts/[^/]+$	57.865	0.665	3.240	3.240	641
0.5%	173	1	0	POST	^/api/posts/[^/]+/comments$	47.695	0.276	2.938	3.024	338
0.4%	308	0	0	GET	^/api/users/[^/]+$	35.030	0.114	0.717	2.261	202
0.3%	38	0	0	DELETE	^/api/posts/[^/]+$	30.020	0.790	3.440	3.440	0
0.3%	228	0	0	GET	^/api/memories/[^/]+	27.084	0.119	1.700	3.512	940
0.3%	205	0	0	GET	/api/me	24.721	0.121	1.134	1.362	304
0.3%	245	0	0	GET	^/api/posts/[^/]+/realmojis$	23.383	0.095	0.295	0.531	1716
0.2%	182	0	0	GET	^/api/posts/[^/]+/comments$	16.480	0.091	0.472	0.555	1787
0.2%	38	0	0	DELETE	^/api/realmojis/[^/]+$	13.597	0.358	1.212	1.212	0
0.1%	85	0	0	GET	/api/me/export	11.604	0.137	0.667	0.667	3070122
0.1%	31	0	0	DELETE	^/api/comments/[^/]+$	6.354	0.205	0.869	0.869	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 15128 クエリ / 48 種類 / 合計 103.446s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	34.1%	350	35.269	100.77	897.62	590.6	60.1	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	21.0%	3	21.729	7242.98	10019.13	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	15.8%	12317	16.313	1.32	13.84	9.5	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	15.6%	3	16.163	5387.63	7332.56	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	7.8%	3	8.090	2696.73	3798.10	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	3.5%	4	3.662	915.62	3141.11	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	4	0.371	92.69	157.59	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.4%	4	0.367	91.64	155.90	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.255	255.24	255.24	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	0.2%	1	0.249	249.32	249.32	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.228	227.80	227.80	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.2%	1	0.167	167.05	167.05	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.107	106.51	106.51	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.081	81.29	81.29	2463.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.1%	1	0.075	74.79	74.79	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: 6c453a46e688975edbb452a19fef99a2b2472918
Type: cpu
Time: 2026-10-31 16:06:44 JST
Duration: 60s, Total samples = 35.61s (59.35%)
Showing nodes accounting for 25.07s, 70.40% of 35.61s total
Dropped 1064 nodes (cum <= 0.18s)
Showing top 40 nodes out of 302
      flat  flat%   sum%        cum   cum%
     5.52s 15.50% 15.50%      5.52s 15.50%  internal/runtime/syscall/linux.Syscall6
     2.32s  6.52% 22.02%      2.32s  6.52%  crypto/internal/fips140/sha256.blockAVX2
     1.54s  4.32% 26.34%      1.54s  4.32%  runtime.memmove
     1.48s  4.16% 30.50%      5.73s 16.09%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.44s  4.04% 34.54%      1.44s  4.04%  runtime.memclrNoHeapPointers
     1.39s  3.90% 38.44%      1.39s  3.90%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.38s  3.88% 42.32%      5.09s 14.29%  runtime.scanObject
     1.30s  3.65% 45.97%      1.71s  4.80%  runtime.findObject
     1.12s  3.15% 49.12%      1.12s  3.15%  aeshashbody
     1.09s  3.06% 52.18%      1.88s  5.28%  runtime.tryDeferToSpanScan
     0.68s  1.91% 54.09%      2.22s  6.23%  runtime.mapaccess1_faststr
     0.55s  1.54% 55.63%      0.55s  1.54%  runtime.spanClass.sizeclass (inline)
     0.35s  0.98% 56.61%      0.35s  0.98%  memeqbody
     0.34s  0.95% 57.57%      0.51s  1.43%  example.com/isureal-company/isureal.postLess
     0.30s  0.84% 58.41%      0.52s  1.46%  database/sql.convertAssignRows
     0.28s  0.79% 59.20%      0.28s  0.79%  runtime.futex
     0.28s  0.79% 59.98%      0.28s  0.79%  sync/atomic.(*Int32).Add
     0.27s  0.76% 60.74%      0.27s  0.76%  runtime.nextFreeFast (inline)
     0.24s  0.67% 61.42%      1.15s  3.23%  runtime.mallocgcSmallScanNoHeader
     0.23s  0.65% 62.06%      0.29s  0.81%  encoding/json.appendString[go.shape.string]
     0.22s  0.62% 62.68%      0.37s  1.04%  internal/runtime/maps.(*Iter).Next
     0.21s  0.59% 63.27%      0.44s  1.24%  compress/flate.(*compressor).deflate
     0.21s  0.59% 63.86%      1.53s  4.30%  runtime.mapassign_faststr
     0.21s  0.59% 64.45%      0.21s  0.59%  runtime.typePointers.nextFast (inline)
     0.18s  0.51% 64.95%      0.24s  0.67%  runtime.spanOf (inline)
     0.17s  0.48% 65.43%      0.20s  0.56%  runtime.typePointers.next
     0.15s  0.42% 65.85%      2.61s  7.33%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.15s  0.42% 66.27%      2.42s  6.80%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.14s  0.39% 66.67%      2.58s  7.25%  crypto/internal/fips140/sha256.(*Digest).Write
     0.14s  0.39% 67.06%      0.19s  0.53%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.14s  0.39% 67.45%      0.63s  1.77%  runtime.mallocgcSmallNoscan
     0.13s  0.37% 67.82%      0.23s  0.65%  runtime.(*mspan).writeHeapBitsSmall
     0.13s  0.37% 68.18%      0.30s  0.84%  runtime.greyobject
     0.13s  0.37% 68.55%      2.88s  8.09%  runtime.mallocgc
     0.12s  0.34% 68.89%      0.77s  2.16%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
     0.12s  0.34% 69.22%      1.05s  2.95%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.34% 69.56%      0.53s  1.49%  runtime.mapaccess2_faststr
     0.11s  0.31% 69.87%      0.28s  0.79%  internal/stringslite.Index
     0.10s  0.28% 70.15%      1.34s  3.76%  example.com/isureal-company/isureal.listMemories
     0.09s  0.25% 70.40%      0.38s  1.07%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 47.0% / 100% | 30% | 8.6% | 11 | 114 MB | run queue が積もっている |
| isu2 (2 cpu) | 27.9% / 100% | 9% | 0.6% | 5 | 106 MB | run queue が積もっている |
| isu3 (2 cpu) | 26.0% / 100% | 7% | 0.7% | 5 | 138 MB | run queue が積もっている |
| isu4 (2 cpu) | 29.5% / 100% | 8% | 8.7% | 6 | 108 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 58% / nginx 10% / kcompactd0 1% |
| isu2 | isureal 52% / nginx 2% |
| isu3 | isureal 49% / nginx 2% |
| isu4 | isureal 26% / postgres 11% |

## Diff from previous run

スコア: 183331 → 182075 (-1256)

`8c13252..8ff9b1d`

```
 webapp/go/admin_api.go      |  27 +++++--
 webapp/go/export_api.go     | 167 +++++++++++++++++++++++++++++++++-----------
 webapp/go/image_api.go      |  23 +++---
 webapp/go/image_file.go     |  33 +++++----
 webapp/go/image_held.go     |  10 +--
 webapp/go/image_local.go    |  47 +++++--------
 webapp/go/ip_ban.go         |  11 +++
 webapp/go/persist.go        |   2 +-
 webapp/go/realmoji_api.go   |  41 +++++------
 webapp/go/repl.go           |   1 +
 webapp/go/replica_server.go |  13 ++--
 webapp/go/store.go          |   3 +-
 12 files changed, 236 insertions(+), 142 deletions(-)
```
