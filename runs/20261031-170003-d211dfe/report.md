# run 20261031-170003-d211dfe

## Summary

| | |
|---|---|
| score | 219108 |
| status | pass |
| ref | HEAD |
| sha | d211dfee0e2c |
| 所要時間 | 417s |
| 前回比 | -1432 (-0.6%) vs 20261031-163606-15a1e6c |
| ベスト比 | -4634 (-2.1%) vs 20261031-163105-6f2cd6f = 223742 |
| note | asia2: Memories・エクスポートの画像の次の分の署名を EU/US に先回り |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	52	153	424	757	757
europe	7	19	56	158	433	432
america	6	16	46	133	332	331
```

```
{"id":"32034","status":"done","portalStatus":"finished","passed":true,"score":219108,"raw":219108,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:05:08.204Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:05:11.195Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:05:14.195Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:05:17.208Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T08:05:20.194Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T08:05:23.194Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T08:05:26.206Z level=INFO msg="asia 3 日目: 153 人"
time=2026-10-31T08:05:29.097Z level=INFO msg="user moved" user=muwaso_5947 from=asia to=europe
time=2026-10-31T08:05:29.195Z level=INFO msg="europe 3 日目: 56 人"
time=2026-10-31T08:05:29.742Z level=INFO msg="user moved" user=kuyuta.6086 from=asia to=europe
time=2026-10-31T08:05:29.758Z level=INFO msg="user moved" user=seru.6047 from=asia to=europe
time=2026-10-31T08:05:30.089Z level=INFO msg="user moved" user=mini5746 from=asia to=europe
time=2026-10-31T08:05:32.195Z level=INFO msg="america 3 日目: 46 人"
time=2026-10-31T08:05:35.211Z level=INFO msg="asia 4 日目: 424 人"
time=2026-10-31T08:05:38.196Z level=INFO msg="europe 4 日目: 158 人"
time=2026-10-31T08:05:40.165Z level=INFO msg="user moved" user=muwaso_5947 from=europe to=america
time=2026-10-31T08:05:40.284Z level=INFO msg="user moved" user=seru.6047 from=europe to=asia
time=2026-10-31T08:05:40.353Z level=INFO msg="user moved" user=kuyuta.6086 from=europe to=america
time=2026-10-31T08:05:41.279Z level=INFO msg="user moved" user=maneeva-7126 from=asia to=europe
time=2026-10-31T08:05:41.366Z level=INFO msg="america 4 日目: 133 人"
time=2026-10-31T08:05:42.205Z level=INFO msg="user moved" user=nema-7635 from=asia to=america
time=2026-10-31T08:05:42.205Z level=INFO msg="user moved" user=zoemamia.8880 from=europe to=america
time=2026-10-31T08:05:42.416Z level=INFO msg="user moved" user=raliamax7768 from=asia to=america
time=2026-10-31T08:05:42.603Z level=INFO msg="user moved" user=fuke-7289 from=asia to=europe
time=2026-10-31T08:05:42.939Z level=INFO msg="user moved" user=chielhe.8981 from=europe to=america
time=2026-10-31T08:05:43.490Z level=INFO msg="user moved" user=miakiha_9128 from=europe to=asia
time=2026-10-31T08:05:44.061Z level=INFO msg="user moved" user=hael_7220 from=asia to=europe
time=2026-10-31T08:05:44.221Z level=INFO msg="asia 5 日目: 757 人"
time=2026-10-31T08:05:44.223Z level=INFO msg="user moved" user=numiasol-8939 from=europe to=asia
time=2026-10-31T08:05:47.205Z level=INFO msg="europe 5 日目: 433 人"
time=2026-10-31T08:05:50.819Z level=INFO msg="america 5 日目: 332 人"
time=2026-10-31T08:05:59.259Z level=INFO msg="america 最終日（グランプリ）: 331 人"
time=2026-10-31T08:05:59.263Z level=INFO msg="europe 最終日（グランプリ）: 432 人"
time=2026-10-31T08:05:59.298Z level=INFO msg="asia 最終日（グランプリ）: 757 人"
time=2026-10-31T08:06:08.254Z level=INFO msg="グランプリ: アジア 521 人, ヨーロッパ 425 人, アメリカ 327 人"
{
  "pass": true,
  "score": 219108,
  "messages": [
    "min_zhang_4: 履歴を読めません: GET /api/memories?month=2026-05: 応答がタイムアウトしました",
    "torinu_8811: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "muwaso_5947: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chiben_12547: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "liam_davis_92: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "raleo_10066: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "croquette_sodachi_1965: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "reme-9692: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "makuwa-6053: ログインし直せません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "tordanben_11752: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "hetotor_8533: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
78.5%	2170	0	0	GET	/api/notifications	57857.369	26.662	60.005	60.005	114
4.8%	11928	0	0	GET	^/api/feeds/[^/]+	3560.655	0.299	1.712	3.426	4936
3.8%	3045	0	0	GET	/api/memories	2765.884	0.908	8.720	20.685	9884
2.4%	2165	39	0	POST	/api/posts	1734.098	0.801	3.916	7.602	640
1.8%	2996	0	0	GET	/api/posting-window/current	1321.022	0.441	1.977	2.439	236
1.2%	2094	0	0	GET	/api/users	903.090	0.431	1.830	2.436	236
1.1%	1837	1	0	GET	^/api/posts/me/current$	824.772	0.449	1.492	2.433	650
1.0%	210	0	0	GET	/api/me/export	758.485	3.612	18.883	23.549	2010813
0.7%	12306	250	0	GET	^/api/images/[^/]+	499.861	0.041	0.586	1.626	177076
0.7%	553	129	0	PUT	/api/me/profile-image	499.329	0.903	11.702	16.899	3
0.6%	875	8	0	GET	^/api/posts/[^/]+/realmojis$	453.306	0.518	2.775	3.883	2889
0.5%	4436	46	0	POST	/api/friend-requests	386.126	0.087	0.873	1.254	114
0.5%	1468	0	0	GET	^/api/relationships/[^/]+	379.052	0.258	1.506	2.320	1424
0.4%	2123	88	0	GET	^/api/posts/[^/]+$	304.331	0.143	0.817	1.307	608
0.4%	433	3	0	GET	/api/me	299.827	0.692	2.188	2.424	301
0.3%	510	0	0	GET	^/api/memories/[^/]+	207.849	0.408	2.227	2.812	692
0.2%	796	12	0	POST	^/api/posts/[^/]+/realmojis$	160.733	0.202	2.991	3.356	474
0.2%	900	9	0	GET	^/api/posts/[^/]+/comments$	143.936	0.160	0.983	1.356	3925
0.2%	1074	0	0	GET	^/api/users/[^/]+$	137.562	0.128	0.839	1.268	210
0.2%	3960	527	0	POST	/api/auth/login	117.776	0.030	0.598	0.894	214
0.1%	3449	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	103.972	0.030	0.586	1.101	116
0.1%	86	0	0	DELETE	^/api/posts/[^/]+$	47.202	0.549	1.428	1.428	0
0.0%	187	1	0	PATCH	^/api/posts/[^/]+$	33.781	0.181	0.830	1.054	648
0.0%	529	0	0	PUT	^/api/follows/[^/]+$	31.631	0.060	0.755	0.902	0
0.0%	996	0	0	POST	^/api/posts/[^/]+/comments$	31.128	0.031	0.677	1.195	342
0.0%	527	0	0	POST	/api/auth/logout	25.313	0.048	0.774	0.973	0
0.0%	1	0	0	POST	/api/initialize	22.177	22.177	22.177	22.177	90
0.0%	112	0	0	DELETE	^/api/realmojis/[^/]+$	21.307	0.190	0.719	1.198	0
0.0%	363	0	0	PATCH	/api/me	17.570	0.048	0.683	0.851	271
0.0%	780	1	0	POST	/api/auth/register	4.704	0.006	0.222	0.411	227
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
84.2%	1104	0	0	GET	/api/notifications	29513.271	26.733	60.006	60.006	107
2.5%	7712	0	0	GET	^/api/feeds/[^/]+	877.469	0.114	0.894	1.307	4647
2.3%	6706	42	0	GET	^/api/images/[^/]+	815.546	0.122	1.103	3.195	188437
1.8%	2580	19	0	POST	/api/friend-requests	622.485	0.241	1.106	1.378	114
1.6%	1108	5	0	POST	/api/posts	557.905	0.504	1.328	64.098	656
1.0%	1811	0	0	POST	/api/auth/login	364.195	0.201	1.116	1.215	249
0.8%	1674	0	0	GET	/api/memories	295.710	0.177	1.436	2.033	12594
0.7%	522	10	0	POST	^/api/posts/[^/]+/realmojis$	232.105	0.445	3.308	4.301	471
0.6%	1099	0	0	GET	^/api/posts/me/current$	195.181	0.178	1.011	1.306	659
0.5%	1588	0	0	GET	/api/posting-window/current	167.724	0.106	0.736	1.111	239
0.5%	431	0	0	POST	/api/auth/register	163.383	0.379	1.161	1.207	229
0.4%	1212	0	0	GET	/api/users	149.732	0.124	0.756	1.123	239
0.4%	751	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	146.791	0.195	1.042	1.193	116
0.3%	1217	21	0	GET	^/api/posts/[^/]+$	116.385	0.096	0.752	1.185	629
0.3%	276	0	0	PUT	/api/me/profile-image	109.553	0.397	1.295	3.885	0
0.3%	916	0	0	GET	^/api/relationships/[^/]+	93.492	0.102	0.793	1.019	712
0.2%	311	0	0	PUT	^/api/follows/[^/]+$	81.614	0.262	1.145	1.197	0
0.2%	402	0	0	POST	^/api/posts/[^/]+/comments$	78.334	0.195	0.897	1.149	351
0.2%	663	0	0	GET	^/api/users/[^/]+$	60.560	0.091	0.675	0.902	212
0.2%	212	0	0	PATCH	/api/me	59.939	0.283	1.182	1.213	290
0.1%	588	1	0	GET	^/api/posts/[^/]+/realmojis$	50.862	0.087	0.754	1.133	3290
0.1%	307	0	0	POST	/api/auth/logout	50.354	0.164	0.979	1.118	0
0.1%	130	0	0	PATCH	^/api/posts/[^/]+$	43.528	0.335	1.006	1.029	662
0.1%	491	0	0	GET	^/api/posts/[^/]+/comments$	43.302	0.088	0.574	0.898	4359
0.1%	59	0	0	DELETE	^/api/posts/[^/]+$	39.552	0.670	1.298	1.298	0
0.1%	270	0	0	GET	^/api/memories/[^/]+	27.579	0.102	0.958	1.257	740
0.1%	81	0	0	DELETE	^/api/realmojis/[^/]+$	26.311	0.325	0.910	0.910	0
0.1%	282	2	0	GET	/api/me	24.741	0.088	0.266	0.317	302
0.0%	111	0	0	GET	/api/me/export	17.480	0.157	1.873	3.088	2195867
0.0%	85	0	0	DELETE	^/api/comments/[^/]+$	11.912	0.140	0.596	0.596	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
85.5%	868	1	0	GET	/api/notifications	22409.927	25.818	60.001	60.002	113
2.5%	5595	31	0	GET	^/api/images/[^/]+	658.569	0.118	1.166	3.089	163148
2.4%	4919	0	0	GET	^/api/feeds/[^/]+	626.615	0.127	1.033	1.483	3802
1.5%	852	2	0	POST	/api/posts	382.337	0.449	1.685	7.337	659
1.4%	1981	21	0	POST	/api/friend-requests	371.752	0.188	0.866	1.383	114
1.1%	1562	0	0	POST	/api/auth/login	288.837	0.185	0.894	1.254	252
0.9%	1420	0	0	GET	/api/memories	241.010	0.170	1.477	2.095	15335
0.6%	849	0	0	GET	^/api/posts/me/current$	145.743	0.172	1.204	1.476	660
0.5%	1215	0	0	GET	/api/posting-window/current	134.199	0.110	0.666	1.436	240
0.5%	270	14	0	POST	^/api/posts/[^/]+/realmojis$	127.304	0.471	3.507	4.327	461
0.4%	480	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	106.817	0.223	0.933	1.141	116
0.4%	949	0	0	GET	/api/users	103.667	0.109	0.832	1.440	237
0.3%	638	0	0	GET	^/api/relationships/[^/]+	68.049	0.107	1.010	1.624	601
0.2%	200	0	0	PUT	/api/me/profile-image	63.780	0.319	0.724	1.013	0
0.2%	325	0	0	POST	/api/auth/register	62.896	0.194	0.502	0.519	230
0.2%	652	13	0	GET	^/api/posts/[^/]+$	61.056	0.094	0.673	1.436	627
0.2%	252	0	0	POST	/api/auth/logout	45.887	0.182	1.026	1.440	0
0.2%	227	5	0	POST	^/api/posts/[^/]+/comments$	45.487	0.200	0.843	0.956	339
0.1%	386	0	0	GET	^/api/users/[^/]+$	33.573	0.087	0.624	0.645	209
0.1%	217	0	0	PUT	^/api/follows/[^/]+$	33.077	0.152	0.612	0.840	0
0.1%	294	0	0	GET	^/api/posts/[^/]+/realmojis$	29.726	0.101	0.680	1.200	4865
0.1%	47	0	0	DELETE	^/api/posts/[^/]+$	28.229	0.601	1.440	1.440	0
0.1%	82	1	0	PATCH	^/api/posts/[^/]+$	26.661	0.325	1.305	1.305	654
0.1%	236	0	0	GET	^/api/memories/[^/]+	24.287	0.103	1.011	1.442	1131
0.1%	151	0	0	PATCH	/api/me	22.228	0.147	0.454	0.479	283
0.1%	259	0	0	GET	^/api/posts/[^/]+/comments$	21.878	0.084	0.398	0.723	4390
0.1%	204	1	0	GET	/api/me	21.510	0.105	0.480	0.535	304
0.0%	80	0	0	GET	/api/me/export	11.898	0.149	1.380	1.380	3102907
0.0%	28	0	0	DELETE	^/api/realmojis/[^/]+$	9.480	0.339	0.766	0.766	0
0.0%	35	0	0	DELETE	^/api/comments/[^/]+$	6.472	0.185	0.860	0.860	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 11131 クエリ / 48 種類 / 合計 140.275s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	44.9%	216	63.026	291.79	903.20	1291.4	131.2	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	18.7%	4	26.227	6556.79	10078.75	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	13.6%	4	19.079	4769.82	7208.46	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.9%	6995	12.478	1.78	8.63	10.2	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.4%	4	9.016	2253.96	3436.58	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	4.2%	7	5.843	834.74	3383.09	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.7%	7	1.031	147.22	535.41	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.4%	7	0.608	86.87	161.89	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.4%	1	0.608	607.65	607.65	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	0.3%	1	0.487	487.05	487.05	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.333	332.71	332.71	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.2%	1	0.263	262.68	262.68	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.2%	1	0.223	222.85	222.85	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.187	187.08	187.08	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	1	0.178	177.88	177.88	2433.0	0.0	DELETE FROM image_objects WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: e772fa1a555e831fc3b98f91bfb71824af8f666a
Type: cpu
Time: 2026-10-31 17:04:41 JST
Duration: 60.02s, Total samples = 28.42s (47.35%)
Showing nodes accounting for 21.94s, 77.20% of 28.42s total
Dropped 811 nodes (cum <= 0.14s)
Showing top 40 nodes out of 243
      flat  flat%   sum%        cum   cum%
     4.06s 14.29% 14.29%      4.06s 14.29%  crypto/internal/fips140/sha256.blockAVX2
     2.61s  9.18% 23.47%      2.61s  9.18%  internal/runtime/syscall/linux.Syscall6
     1.74s  6.12% 29.59%      5.65s 19.88%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.40s  4.93% 34.52%      1.40s  4.93%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.28s  4.50% 39.02%      1.28s  4.50%  runtime.memmove
     1.13s  3.98% 43.00%      1.13s  3.98%  aeshashbody
     0.92s  3.24% 46.24%      1.20s  4.22%  runtime.findObject
     0.87s  3.06% 49.30%      3.28s 11.54%  runtime.scanObject
     0.82s  2.89% 52.18%      0.82s  2.89%  runtime.memclrNoHeapPointers
     0.80s  2.81% 55.00%      1.22s  4.29%  example.com/isureal-company/isureal/jpegcheck.(*decoder).skipAC
     0.57s  2.01% 57.00%      1.07s  3.76%  runtime.tryDeferToSpanScan
     0.50s  1.76% 58.76%      1.90s  6.69%  runtime.mapaccess1_faststr
     0.39s  1.37% 60.13%      0.54s  1.90%  example.com/isureal-company/isureal.postLess
     0.39s  1.37% 61.51%      0.39s  1.37%  example.com/isureal-company/isureal/jpegcheck.(*decoder).prefetch
     0.34s  1.20% 62.70%      0.34s  1.20%  runtime.spanClass.sizeclass (inline)
     0.33s  1.16% 63.86%      0.43s  1.51%  database/sql.convertAssignRows
     0.26s  0.91% 64.78%      4.61s 16.22%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.25s  0.88% 65.66%      4.51s 15.87%  crypto/internal/fips140/sha256.(*Digest).Write
     0.25s  0.88% 66.54%      1.55s  5.45%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.24s  0.84% 67.38%      0.24s  0.84%  memeqbody
     0.24s  0.84% 68.23%      0.24s  0.84%  runtime.futex
     0.24s  0.84% 69.07%      0.24s  0.84%  runtime.nextFreeFast (inline)
     0.24s  0.84% 69.92%      0.24s  0.84%  sync/atomic.(*Int32).Add
     0.19s  0.67% 70.58%      2.31s  8.13%  runtime.mallocgc
     0.17s   0.6% 71.18%      5.27s 18.54%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.17s   0.6% 71.78%      0.20s   0.7%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.17s   0.6% 72.38%      0.91s  3.20%  runtime.mallocgcSmallScanNoHeader
     0.16s  0.56% 72.94%      0.28s  0.99%  runtime.greyobject
     0.15s  0.53% 73.47%      4.85s 17.07%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.15s  0.53% 74.00%      0.16s  0.56%  runtime.(*mspan).writeHeapBitsSmall
     0.13s  0.46% 74.45%      1.39s  4.89%  runtime.mapassign_faststr
     0.12s  0.42% 74.88%      0.22s  0.77%  example.com/isureal-company/isureal.(*Store).snapshotLocked
     0.11s  0.39% 75.26%      0.40s  1.41%  runtime.mapaccess2_faststr
     0.09s  0.32% 75.58%      0.28s  0.99%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.09s  0.32% 75.90%      1.88s  6.62%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.08s  0.28% 76.18%      5.51s 19.39%  example.com/isureal-company/isureal.computeImageSignature
     0.08s  0.28% 76.46%      1.08s  3.80%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.07s  0.25% 76.71%      0.19s  0.67%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.07s  0.25% 76.95%      0.50s  1.76%  runtime.scanSpan
     0.07s  0.25% 77.20%      0.44s  1.55%  sort.insertionSort_func
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 30.3% / 100% | 9% | 0.7% | 5 | 168 MB | run queue が積もっている |
| isu2 (2 cpu) | 27.3% / 100% | 11% | 0.7% | 6 | 102 MB | run queue が積もっている |
| isu3 (2 cpu) | 19.8% / 100% | 4% | 0.7% | 5 | 193 MB | run queue が積もっている |
| isu4 (2 cpu) | 58.0% / 100% | 52% | 2.6% | 5 | 108 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 47% / nginx 9% |
| isu2 | isureal 50% / nginx 2% |
| isu3 | isureal 37% / nginx 2% |
| isu4 | isureal 94% / postgres 14% |

## Diff from previous run

スコア: 220540 → 219108 (-1432)

`15a1e6c..d211dfe`

```
 conf/isu4/home/isucon/env.sh |   1 +
 webapp/go/offload.go         | 137 +++++++++++++++++++++++++++++++++++++++++++
 2 files changed, 138 insertions(+)
```
