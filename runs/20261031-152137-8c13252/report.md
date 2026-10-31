# run 20261031-152137-8c13252

## Summary

| | |
|---|---|
| score | 183331 |
| status | pass |
| ref | 8c13252 |
| sha | 8c13252c3064 |
| 所要時間 | 125s |
| 前回比 | +147 (0.1%) vs 20261031-151835-2fe5b6e |
| ベスト比 | -8485 (-4.4%) vs 20261031-151339-5a230a4 = 191816 |
| note | RealMoji の画像も受けた台に置く（main 5fcd3e6 の上）2回目 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	53	153	454	1241	1241
europe	7	21	57	161	431	430
america	6	26	73	203	318	317
```

```
{"id":"22166","status":"done","portalStatus":"finished","passed":true,"score":183331,"raw":183331,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:22:03.917Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:22:06.913Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:22:09.913Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:22:12.961Z level=INFO msg="asia 2 日目: 53 人"
time=2026-10-31T06:22:13.688Z level=INFO msg="user moved" user=yaevano.5169 from=asia to=america
time=2026-10-31T06:22:14.372Z level=INFO msg="user moved" user=nita-5163 from=asia to=europe
time=2026-10-31T06:22:14.636Z level=INFO msg="user moved" user=matone5254 from=asia to=america
time=2026-10-31T06:22:14.957Z level=INFO msg="user moved" user=yaevano.5169 from=america to=europe
time=2026-10-31T06:22:15.722Z level=INFO msg="user moved" user=yaor.5338 from=asia to=america
time=2026-10-31T06:22:15.913Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T06:22:16.614Z level=INFO msg="user moved" user=namarra_5393 from=asia to=america
time=2026-10-31T06:22:16.808Z level=INFO msg="user moved" user=yaor.5338 from=america to=europe
time=2026-10-31T06:22:18.915Z level=INFO msg="america 2 日目: 26 人"
time=2026-10-31T06:22:21.930Z level=INFO msg="asia 3 日目: 153 人"
time=2026-10-31T06:22:22.790Z level=INFO msg="user moved" user=nivinyo_5729 from=asia to=europe
time=2026-10-31T06:22:22.832Z level=INFO msg="user moved" user=senoa.5534 from=america to=europe
time=2026-10-31T06:22:23.075Z level=INFO msg="user moved" user=noasoltsu5611 from=asia to=america
time=2026-10-31T06:22:23.483Z level=INFO msg="user moved" user=nivinyo_5729 from=europe to=america
time=2026-10-31T06:22:24.914Z level=INFO msg="europe 3 日目: 57 人"
time=2026-10-31T06:22:26.028Z level=INFO msg="user moved" user=yaor.5338 from=europe to=america
time=2026-10-31T06:22:26.794Z level=INFO msg="user moved" user=noasoltsu5611 from=america to=europe
time=2026-10-31T06:22:27.914Z level=INFO msg="america 3 日目: 73 人"
time=2026-10-31T06:22:28.263Z level=INFO msg="user moved" user=yaevano.5169 from=europe to=asia
time=2026-10-31T06:22:31.023Z level=INFO msg="asia 4 日目: 454 人"
time=2026-10-31T06:22:32.538Z level=INFO msg="user moved" user=evami_6612 from=america to=asia
time=2026-10-31T06:22:34.107Z level=INFO msg="europe 4 日目: 161 人"
time=2026-10-31T06:22:34.983Z level=INFO msg="user moved" user=chilia-6927 from=asia to=america
time=2026-10-31T06:22:34.993Z level=INFO msg="user moved" user=evami_6612 from=asia to=europe
time=2026-10-31T06:22:35.051Z level=INFO msg="user moved" user=matone5254 from=america to=asia
time=2026-10-31T06:22:35.104Z level=INFO msg="user moved" user=namarra_5393 from=america to=europe
time=2026-10-31T06:22:35.154Z level=INFO msg="user moved" user=mutone-7973 from=asia to=europe
time=2026-10-31T06:22:35.244Z level=INFO msg="user moved" user=haben_7618 from=asia to=america
time=2026-10-31T06:22:35.420Z level=INFO msg="user moved" user=hoshimu.6708 from=america to=asia
time=2026-10-31T06:22:36.898Z level=INFO msg="user moved" user=yunoalia_8705 from=asia to=america
time=2026-10-31T06:22:37.011Z level=INFO msg="user moved" user=haben_7618 from=america to=asia
time=2026-10-31T06:22:37.179Z level=INFO msg="america 4 日目: 203 人"
time=2026-10-31T06:22:37.626Z level=INFO msg="user moved" user=nemuro_8943 from=asia to=europe
time=2026-10-31T06:22:37.715Z level=INFO msg="user moved" user=yaevano.5169 from=asia to=america
time=2026-10-31T06:22:37.913Z level=INFO msg="user moved" user=hoshimu.6708 from=asia to=america
time=2026-10-31T06:22:38.040Z level=INFO msg="user moved" user=matone5254 from=asia to=america
time=2026-10-31T06:22:38.434Z level=INFO msg="user moved" user=evakomi-7197 from=asia to=america
time=2026-10-31T06:22:38.497Z level=INFO msg="user moved" user=yaor.5338 from=america to=europe
time=2026-10-31T06:22:40.070Z level=INFO msg="asia 5 日目: 1241 人"
time=2026-10-31T06:22:41.645Z level=INFO msg="user moved" user=sudanso9422 from=europe to=america
time=2026-10-31T06:22:43.477Z level=INFO msg="europe 5 日目: 431 人"
time=2026-10-31T06:22:43.570Z level=INFO msg="user moved" user=nivinyo_5729 from=america to=asia
time=2026-10-31T06:22:46.779Z level=INFO msg="america 5 日目: 318 人"
time=2026-10-31T06:22:55.961Z level=INFO msg="asia 最終日（グランプリ）: 1241 人"
time=2026-10-31T06:22:56.512Z level=INFO msg="america 最終日（グランプリ）: 317 人"
time=2026-10-31T06:22:56.782Z level=INFO msg="europe 最終日（グランプリ）: 430 人"
time=2026-10-31T06:23:04.067Z level=INFO msg="グランプリ: アジア 492 人, ヨーロッパ 135 人, アメリカ 89 人"
{
  "pass": true,
  "score": 183331,
  "messages": [
    "muse_5048: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "vinno.5090: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "orre_9040: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kura_16491: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hine-17551: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "heshitsu_17366: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kunayo_8858: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ultene16553: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "elri.17799: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kimiha11529: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "solzoe-10283: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "more_12434: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yayaya16153: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "teroke.8842: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "soorfu17530: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tormo_17577: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sete17915: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "watsune17102: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "fusol9542: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "vinruzoe-17637: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.4%	1935	1	0	GET	/api/notifications	17318.456	8.950	15.907	15.918	161
12.3%	3028	254	0	POST	/api/posts	5138.928	1.697	6.875	8.157	598
12.0%	6888	97	0	POST	/api/friend-requests	5010.125	0.727	3.000	4.965	114
5.5%	20199	305	0	GET	^/api/images/[^/]+	2313.974	0.115	1.387	3.453	118346
4.9%	7354	655	0	POST	/api/auth/login	2063.690	0.281	2.273	4.213	224
3.5%	2911	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1453.261	0.499	2.570	3.870	116
3.0%	4859	0	0	GET	/api/memories	1259.228	0.259	2.594	6.231	12662
2.4%	852	129	0	PUT	/api/me/profile-image	1007.712	1.183	9.703	12.592	2
2.4%	14129	0	0	GET	^/api/feeds/[^/]+	998.827	0.071	0.740	2.089	3236
1.6%	614	7	0	POST	^/api/posts/[^/]+/realmojis$	656.990	1.070	4.322	6.426	480
1.4%	786	0	0	PUT	^/api/follows/[^/]+$	575.557	0.732	3.206	4.329	0
1.2%	620	5	0	POST	^/api/posts/[^/]+/comments$	482.331	0.778	3.878	6.060	337
1.1%	1269	1	0	POST	/api/auth/register	451.358	0.356	1.675	1.772	227
1.0%	605	0	0	PATCH	/api/me	405.423	0.670	2.923	3.772	279
0.9%	1150	0	0	POST	/api/auth/logout	370.003	0.322	2.619	3.681	0
0.8%	446	0	0	GET	/api/me/export	326.009	0.731	6.990	10.531	1758440
0.8%	3338	0	0	GET	/api/users	324.152	0.097	0.734	1.021	251
0.8%	4419	0	0	GET	/api/posting-window/current	318.891	0.072	0.590	1.116	237
0.5%	2274	60	0	GET	^/api/posts/me/current$	205.435	0.090	0.733	1.359	636
0.5%	258	9	0	PATCH	^/api/posts/[^/]+$	192.770	0.747	3.080	3.952	631
0.4%	738	3	0	GET	/api/me	165.514	0.224	0.804	1.162	301
0.3%	1842	0	0	GET	^/api/relationships/[^/]+	132.847	0.072	0.813	2.601	1056
0.3%	107	0	0	DELETE	^/api/posts/[^/]+$	121.734	1.138	3.098	3.372	0
0.3%	1053	0	0	GET	^/api/memories/[^/]+	112.304	0.107	1.052	2.596	640
0.2%	1876	33	0	GET	^/api/posts/[^/]+$	101.742	0.054	0.524	0.907	612
0.1%	89	0	0	DELETE	^/api/realmojis/[^/]+$	61.488	0.691	3.024	3.024	0
0.1%	115	0	0	DELETE	^/api/comments/[^/]+$	57.743	0.502	2.356	2.980	0
0.1%	980	1	0	GET	^/api/users/[^/]+$	48.826	0.050	0.400	0.863	199
0.1%	777	0	0	GET	^/api/posts/[^/]+/realmojis$	40.903	0.053	0.552	0.847	1586
0.1%	701	0	0	GET	^/api/posts/[^/]+/comments$	34.751	0.050	0.451	1.052	1918
…（上位 30 / 全 43 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.1%	634	1	0	GET	/api/notifications	4795.892	7.564	13.163	13.218	150
11.9%	1939	19	0	POST	/api/friend-requests	1423.339	0.734	2.931	4.034	114
9.0%	863	69	0	POST	/api/posts	1070.313	1.240	6.638	6.851	591
6.8%	1257	0	0	POST	/api/auth/login	817.967	0.651	2.999	4.846	245
4.2%	422	0	0	POST	/api/auth/register	499.695	1.184	2.813	2.841	229
4.1%	3809	20	0	GET	^/api/images/[^/]+	495.052	0.130	1.378	1.875	157314
4.0%	3099	0	0	GET	^/api/feeds/[^/]+	478.847	0.155	1.314	3.125	2679
2.7%	218	0	0	PUT	/api/me/profile-image	322.764	1.481	3.895	4.708	0
2.6%	469	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	316.575	0.675	3.108	3.565	116
2.1%	259	0	0	PUT	^/api/follows/[^/]+$	245.866	0.949	3.057	4.059	0
1.9%	197	0	0	PATCH	/api/me	223.885	1.136	3.790	3.805	274
1.5%	943	0	0	GET	/api/memories	183.166	0.194	2.115	4.307	13440
1.5%	1322	0	0	GET	/api/posting-window/current	176.612	0.134	1.259	1.605	239
1.1%	164	0	0	POST	^/api/posts/[^/]+/comments$	127.511	0.778	5.060	5.367	332
1.1%	937	0	0	GET	/api/users	125.890	0.134	1.039	2.208	246
0.9%	592	35	0	GET	^/api/posts/me/current$	109.953	0.186	1.418	1.775	604
0.8%	118	1	0	POST	^/api/posts/[^/]+/realmojis$	100.783	0.854	4.095	4.526	446
0.6%	170	0	0	POST	/api/auth/logout	73.639	0.433	2.491	3.795	0
0.5%	64	1	0	PATCH	^/api/posts/[^/]+$	62.963	0.984	3.288	3.288	632
0.5%	406	0	0	GET	^/api/relationships/[^/]+	55.603	0.137	1.498	1.914	823
0.4%	412	17	0	GET	^/api/posts/[^/]+$	49.911	0.121	0.710	0.809	605
0.3%	24	0	0	DELETE	^/api/posts/[^/]+$	35.617	1.484	3.747	3.747	0
0.2%	208	0	0	GET	^/api/users/[^/]+$	24.068	0.116	0.626	0.705	200
0.2%	223	2	0	GET	/api/me	23.372	0.105	0.695	0.883	302
0.2%	37	0	0	DELETE	^/api/comments/[^/]+$	23.055	0.623	1.889	1.889	0
0.2%	163	0	0	GET	^/api/memories/[^/]+	19.986	0.123	1.022	1.919	977
0.2%	158	0	0	GET	^/api/posts/[^/]+/realmojis$	18.229	0.115	0.711	0.720	1561
0.1%	151	0	0	GET	^/api/posts/[^/]+/comments$	17.102	0.113	0.716	0.745	1932
0.1%	20	0	0	DELETE	^/api/realmojis/[^/]+$	12.866	0.643	2.432	2.432	0
0.1%	64	0	0	GET	/api/me/export	12.832	0.201	3.659	3.659	2374837
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.9%	574	1	0	GET	/api/notifications	4291.520	7.477	10.030	10.404	154
10.5%	762	51	0	POST	/api/posts	1074.070	1.410	6.549	6.867	615
10.4%	1617	23	0	POST	/api/friend-requests	1063.340	0.658	2.837	3.836	114
5.5%	1150	0	0	POST	/api/auth/login	562.851	0.489	2.507	3.478	249
5.2%	3764	23	0	GET	^/api/images/[^/]+	535.666	0.142	1.342	2.529	158072
4.9%	3180	0	0	GET	^/api/feeds/[^/]+	500.103	0.157	1.509	2.241	3615
2.7%	313	0	0	POST	/api/auth/register	275.042	0.879	2.785	2.853	230
2.5%	462	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	255.691	0.553	2.549	4.090	116
1.8%	1126	0	0	GET	/api/posting-window/current	185.312	0.165	1.239	2.712	239
1.7%	943	0	0	GET	/api/memories	173.552	0.184	2.292	3.839	15960
1.6%	182	0	0	PUT	/api/me/profile-image	162.457	0.893	3.225	3.278	0
1.5%	226	0	0	PUT	^/api/follows/[^/]+$	149.216	0.660	2.474	2.505	0
1.3%	164	0	0	PATCH	/api/me	132.742	0.809	3.386	3.804	280
1.2%	172	3	0	POST	^/api/posts/[^/]+/comments$	121.780	0.708	3.537	3.606	347
1.1%	223	0	0	POST	/api/auth/logout	112.478	0.504	2.197	2.484	0
1.1%	809	0	0	GET	/api/users	111.263	0.138	1.209	2.802	255
1.0%	543	23	0	GET	^/api/posts/me/current$	105.671	0.195	2.096	2.303	632
0.7%	81	2	0	POST	^/api/posts/[^/]+/realmojis$	76.160	0.940	3.338	3.338	484
0.5%	388	0	0	GET	^/api/relationships/[^/]+	53.770	0.139	1.445	1.475	857
0.5%	482	14	0	GET	^/api/posts/[^/]+$	51.651	0.107	0.723	1.262	608
0.4%	59	4	0	PATCH	^/api/posts/[^/]+$	43.675	0.740	2.977	2.977	616
0.4%	27	0	0	DELETE	^/api/posts/[^/]+$	36.876	1.366	3.859	3.859	0
0.3%	241	0	0	GET	^/api/users/[^/]+$	27.616	0.115	0.702	0.839	198
0.2%	187	1	0	GET	/api/me	25.557	0.137	0.905	1.813	304
0.2%	162	0	0	GET	^/api/memories/[^/]+	25.015	0.154	1.443	1.691	932
0.2%	66	0	0	GET	/api/me/export	21.716	0.329	2.854	2.854	2979482
0.2%	172	0	0	GET	^/api/posts/[^/]+/comments$	17.270	0.100	0.673	0.785	1750
0.2%	39	0	0	DELETE	^/api/comments/[^/]+$	15.590	0.400	2.022	2.022	0
0.2%	152	0	0	GET	^/api/posts/[^/]+/realmojis$	15.521	0.102	0.679	0.993	1513
0.1%	13	0	0	DELETE	^/api/realmojis/[^/]+$	12.721	0.979	2.530	2.530	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 18335 クエリ / 48 種類 / 合計 157.538s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	55.8%	454	87.928	193.67	1921.93	1117.9	114.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.0%	3	21.999	7332.94	10134.85	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.6%	14552	18.233	1.25	7.76	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.1%	3	15.957	5318.99	7193.44	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.0%	3	7.857	2619.10	3532.48	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.1%	4	3.326	831.60	2819.25	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.2%	4	0.352	87.97	156.40	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
8	0.2%	4	0.350	87.45	156.08	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.2%	1	0.290	290.14	290.14	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.251	250.98	250.98	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.1%	1	0.227	227.37	227.37	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.162	162.48	162.48	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.106	105.93	105.93	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.080	79.57	79.57	2412.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.074	74.31	74.31	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: 39be65b6c87c38f3c098ee9cd8339ed0239294a3
Type: cpu
Time: 2026-10-31 15:21:43 JST
Duration: 60.09s, Total samples = 44.80s (74.56%)
Showing nodes accounting for 30.81s, 68.77% of 44.80s total
Dropped 1168 nodes (cum <= 0.22s)
Showing top 40 nodes out of 294
      flat  flat%   sum%        cum   cum%
     5.38s 12.01% 12.01%      5.38s 12.01%  internal/runtime/syscall/linux.Syscall6
     5.01s 11.18% 23.19%      5.01s 11.18%  crypto/internal/fips140/sha256.blockAVX2
     1.85s  4.13% 27.32%      1.85s  4.13%  runtime.memmove
     1.83s  4.08% 31.41%      2.32s  5.18%  runtime.findObject
     1.69s  3.77% 35.18%      1.69s  3.77%  runtime.memclrNoHeapPointers
     1.67s  3.73% 38.91%      1.67s  3.73%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.61s  3.59% 42.50%      6.63s 14.80%  runtime.scanObject
     1.49s  3.33% 45.83%      5.91s 13.19%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.32s  2.95% 48.77%      2.26s  5.04%  runtime.tryDeferToSpanScan
     1.24s  2.77% 51.54%      1.24s  2.77%  aeshashbody
     0.73s  1.63% 53.17%      0.73s  1.63%  runtime.spanClass.sizeclass (inline)
     0.70s  1.56% 54.73%      2.49s  5.56%  runtime.mapaccess1_faststr
     0.49s  1.09% 55.83%      0.49s  1.09%  runtime.nextFreeFast (inline)
     0.48s  1.07% 56.90%      0.48s  1.07%  memeqbody
     0.43s  0.96% 57.86%      0.60s  1.34%  example.com/isureal-company/isureal.postLess
     0.33s  0.74% 58.59%      5.69s 12.70%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.29s  0.65% 59.24%      0.83s  1.85%  database/sql.convertAssignRows
     0.29s  0.65% 59.89%      0.29s  0.65%  runtime.futex
     0.27s   0.6% 60.49%      2.19s  4.89%  runtime.mapassign_faststr
     0.27s   0.6% 61.09%      0.27s   0.6%  sync/atomic.(*Int32).Add (inline)
     0.25s  0.56% 61.65%      5.52s 12.32%  crypto/internal/fips140/sha256.(*Digest).Write
     0.25s  0.56% 62.21%      5.17s 11.54%  runtime.mallocgc
     0.25s  0.56% 62.77%      1.62s  3.62%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.54% 63.30%      0.27s   0.6%  runtime.typePointers.next
     0.23s  0.51% 63.82%      0.23s  0.51%  runtime.typePointers.nextFast (inline)
     0.22s  0.49% 64.31%      0.34s  0.76%  internal/runtime/maps.(*Iter).Next
     0.22s  0.49% 64.80%      0.49s  1.09%  runtime.greyobject
     0.20s  0.45% 65.25%      0.54s  1.21%  compress/flate.(*compressor).deflate
     0.20s  0.45% 65.69%      5.98s 13.35%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.18s   0.4% 66.09%      6.38s 14.24%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.17s  0.38% 66.47%      0.42s  0.94%  runtime.scanObjectsSmall
     0.15s  0.33% 66.81%      0.52s  1.16%  runtime.mapaccess2_faststr
     0.14s  0.31% 67.12%      0.39s  0.87%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.12s  0.27% 67.39%      3.72s  8.30%  example.com/isureal-company/isureal.loadBig
     0.12s  0.27% 67.66%      0.74s  1.65%  runtime.wbBufFlush1
     0.11s  0.25% 67.90%      0.34s  0.76%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.11s  0.25% 68.15%      0.32s  0.71%  internal/stringslite.Index
     0.10s  0.22% 68.37%      0.26s  0.58%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.09s   0.2% 68.57%      0.78s  1.74%  encoding/json.structEncoder.encode
     0.09s   0.2% 68.77%      2.91s  6.50%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 46.9% / 100% | 32% | 0.4% | 4 | 107 MB | CPU が飽和 |
| isu2 (2 cpu) | 27.5% / 100% | 5% | 0.6% | 4 | 314 MB |  |
| isu3 (2 cpu) | 27.6% / 100% | 5% | 0.8% | 5 | 262 MB | run queue が積もっている |
| isu4 (2 cpu) | 39.4% / 100% | 17% | 13.9% | 22 | 105 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 77% / nginx 13% |
| isu2 | isureal 52% / nginx 2% |
| isu3 | isureal 52% / nginx 2% |
| isu4 | isureal 33% / postgres 15% |

## Diff from previous run

スコア: 183184 → 183331 (+147)

`2fe5b6e..8c13252`

```
 webapp/go/image_local.go    |  7 +------
 webapp/go/ip_ban.go         | 35 +++++++++++++++--------------------
 webapp/go/replica_server.go |  1 +
 webapp/go/server.go         |  1 +
 4 files changed, 18 insertions(+), 26 deletions(-)
```
