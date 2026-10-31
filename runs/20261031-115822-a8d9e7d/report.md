# run 20261031-115822-a8d9e7d

## Summary

| | |
|---|---|
| score | 56314 |
| status | pass |
| ref | origin/main |
| sha | a8d9e7d3cca4 |
| 所要時間 | 194s |
| 前回比 | -239 (-0.4%) vs 20261031-115409-773f053 |
| ベスト比 | -3887 (-6.5%) vs 20261031-115040-1c73ee4 = 60201 |
| note | main: 登録の並列化を入れ直した |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	138	392	392	391
europe	7	19	52	116	116	116
america	6	15	36	66	66	65
```

```
{"id":"6245","status":"done","portalStatus":"finished","passed":true,"score":56314,"raw":56314,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:00:20.616Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:00:23.557Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:00:26.560Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:00:29.600Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T03:00:32.558Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T03:00:35.559Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T03:00:38.597Z level=INFO msg="asia 3 日目: 138 人"
time=2026-10-31T03:00:41.677Z level=INFO msg="europe 3 日目: 52 人"
time=2026-10-31T03:00:44.566Z level=INFO msg="america 3 日目: 36 人"
time=2026-10-31T03:00:47.666Z level=INFO msg="asia 4 日目: 392 人"
time=2026-10-31T03:00:50.797Z level=INFO msg="europe 4 日目: 116 人"
time=2026-10-31T03:00:53.610Z level=INFO msg="america 4 日目: 66 人"
time=2026-10-31T03:00:56.799Z level=INFO msg="asia 5 日目: 392 人"
time=2026-10-31T03:00:59.896Z level=INFO msg="europe 5 日目: 116 人"
time=2026-10-31T03:01:02.705Z level=INFO msg="america 5 日目: 66 人"
time=2026-10-31T03:01:11.712Z level=INFO msg="america 最終日（グランプリ）: 65 人"
time=2026-10-31T03:01:11.811Z level=INFO msg="europe 最終日（グランプリ）: 116 人"
time=2026-10-31T03:01:11.919Z level=INFO msg="asia 最終日（グランプリ）: 391 人"
time=2026-10-31T03:01:20.570Z level=INFO msg="グランプリ: アジア 80 人, ヨーロッパ 29 人, アメリカ 11 人"
{
  "pass": true,
  "score": 56314,
  "messages": [
    "keivymi.5750: 友達申請を承認できません: POST /api/friend-requests/700969cc-4026-40d4-a40b-dca85ce20f1e/accept: 応答がタイムアウトしました",
    "teru.7100: プロフィール画像を変えられません: PUT /api/me/profile-image: 応答がタイムアウトしました",
    "nileo_6914: 新しいユーザーを登録できません: POST /api/auth/login: 応答がタイムアウトしました",
    "neivysol-5621: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "rekoya_7121: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "nima-8421: 認定公開アカウントをフォローできません: PUT /api/follows/078a6718-eefd-421d-a8d7-c2e3f30ce87b: 応答がタイムアウトしました",
    "torzoemax_7598: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "danhi_7770: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "komia-5809: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "tsuanmi-8620: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "tordan-7047: 認定公開アカウントをフォローできません: PUT /api/follows/078a6718-eefd-421d-a8d7-c2e3f30ce87b: 応答がタイムアウトしました",
    "tafu.7994: 認定公開アカウントをフォローできません: PUT /api/follows/078a6718-eefd-421d-a8d7-c2e3f30ce87b: 応答がタイムアウトしました",
    "keyoto-5374: 投稿 130f3c6b-0303-49ef-91dc-d2a1b84e7f41 に RealMoji を送れません: POST /api/posts/130f3c6b-0303-49ef-91dc-d2a1b84e7f41/realmojis: 応答がタイムアウトしました",
    "noadan-6791: 認定公開アカウントをフォローできません: PUT /api/follows/078a6718-eefd-421d-a8d7-c2e3f30ce87b: 応答がタイムアウトしました",
    "torne.5266: 投稿 130f3c6b-0303-49ef-91dc-d2a1b84e7f41 にコメントできません: POST /api/posts/130f3c6b-0303-49ef-91dc-d2a1b84e7f41/comments: 応答がタイムアウトしました",
    "romo7180: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "muwaki.5381: 投稿 130f3c6b-0303-49ef-91dc-d2a1b84e7f41 に RealMoji を送れません: POST /api/posts/130f3c6b-0303-49ef-91dc-d2a1b84e7f41/realmojis: 応答がタイムアウトしました",
    "bensueva6592: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "hoanmu6682: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "teheke.6644: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
29.3%	705	0	0	GET	/api/notifications	6375.056	9.043	15.164	15.310	175
19.6%	2038	53	0	POST	/api/friend-requests	4267.017	2.094	23.298	28.627	113
10.9%	904	11	0	POST	/api/posts	2366.267	2.618	9.786	13.887	638
8.0%	7850	561	0	GET	^/api/images/[^/]+	1733.843	0.221	3.049	9.253	106222
7.1%	247	0	0	PUT	^/api/follows/[^/]+$	1533.462	6.208	23.863	26.975	0
7.0%	3573	1223	0	POST	/api/auth/login	1517.756	0.425	5.730	31.434	171
6.0%	866	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1306.368	1.509	16.519	31.459	116
1.8%	408	1	0	POST	/api/auth/register	399.931	0.980	3.847	7.710	227
1.8%	221	37	0	POST	^/api/posts/[^/]+/comments$	382.991	1.733	15.969	26.085	287
1.2%	2072	0	0	GET	/api/memories	261.745	0.126	0.826	2.495	11307
1.1%	322	115	0	PUT	/api/me/profile-image	244.371	0.759	4.344	31.438	9
1.0%	149	23	0	POST	^/api/posts/[^/]+/realmojis$	225.656	1.514	15.302	17.307	410
1.0%	1259	0	0	GET	/api/posting-window/current	207.176	0.165	1.430	2.898	236
0.8%	5298	0	0	GET	^/api/feeds/[^/]+	180.337	0.034	0.209	1.580	2075
0.7%	163	0	0	GET	/api/me/export	162.742	0.998	2.977	9.181	2262154
0.5%	202	0	0	PATCH	/api/me	104.942	0.520	2.591	2.962	270
0.5%	381	0	0	POST	/api/auth/logout	99.873	0.262	3.227	8.660	0
0.2%	87	2	0	PATCH	^/api/posts/[^/]+$	50.151	0.576	2.728	2.728	629
0.2%	37	0	0	DELETE	^/api/posts/[^/]+$	45.544	1.231	7.155	7.155	0
0.2%	952	0	0	GET	/api/users	44.485	0.047	0.448	1.449	238
0.2%	721	15	0	GET	^/api/posts/me/current$	42.101	0.058	0.735	1.370	632
0.1%	860	8	0	GET	^/api/posts/[^/]+$	23.571	0.027	0.167	0.291	602
0.1%	2	0	0	DELETE	^/api/follows/[^/]+$	23.158	11.579	13.946	13.946	0
0.1%	843	0	0	GET	^/api/relationships/[^/]+	22.486	0.027	0.187	0.826	1121
0.1%	341	0	0	GET	^/api/memories/[^/]+	21.936	0.064	0.368	2.861	660
0.1%	20	0	0	PUT	^/api/blocks/[^/]+$	14.588	0.729	3.827	3.827	0
0.1%	455	0	0	GET	^/api/users/[^/]+$	13.010	0.029	0.145	0.241	186
0.1%	1	0	0	POST	/api/initialize	12.733	12.733	12.733	12.733	90
0.1%	18	0	0	DELETE	^/api/realmojis/[^/]+$	11.121	0.618	3.874	3.874	0
0.0%	365	0	0	GET	^/api/posts/[^/]+/realmojis$	10.091	0.028	0.156	0.217	400
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
31.2%	231	0	0	GET	/api/notifications	1896.191	8.209	11.968	11.968	176
15.7%	517	8	0	POST	/api/friend-requests	957.220	1.851	20.881	25.664	113
14.1%	285	6	0	POST	/api/posts	860.437	3.019	10.125	19.940	631
9.7%	73	0	0	PUT	^/api/follows/[^/]+$	591.014	8.096	27.047	27.047	0
6.7%	493	0	0	POST	/api/auth/login	407.838	0.827	7.900	9.545	247
3.7%	1527	9	0	GET	^/api/images/[^/]+	222.689	0.146	1.090	1.498	140335
3.0%	133	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	184.534	1.387	15.370	21.183	116
2.5%	1008	0	0	GET	^/api/feeds/[^/]+	154.642	0.153	0.492	0.779	2110
1.8%	490	0	0	GET	/api/memories	109.377	0.223	1.052	2.048	11078
1.8%	80	14	0	POST	^/api/posts/[^/]+/realmojis$	108.034	1.350	16.720	16.720	380
1.7%	388	0	0	GET	/api/posting-window/current	105.296	0.271	1.633	1.666	239
1.2%	114	0	0	POST	/api/auth/register	73.156	0.642	2.691	2.772	227
0.9%	47	10	0	POST	^/api/posts/[^/]+/comments$	55.406	1.179	13.164	13.164	233
0.8%	73	0	0	POST	/api/auth/logout	50.289	0.689	9.940	9.940	0
0.7%	30	0	0	GET	/api/me/export	42.946	1.432	3.266	3.266	2001908
0.7%	249	0	0	GET	/api/users	42.854	0.172	1.299	1.514	241
0.6%	223	7	0	GET	^/api/posts/me/current$	34.424	0.154	0.435	0.514	626
0.5%	58	0	0	PUT	/api/me/profile-image	32.079	0.553	4.724	4.724	0
0.4%	183	0	0	GET	^/api/relationships/[^/]+	26.515	0.145	0.397	1.688	724
0.4%	19	1	0	PATCH	^/api/posts/[^/]+$	22.598	1.189	6.683	6.683	606
0.3%	50	0	0	PATCH	/api/me	21.122	0.422	2.181	2.181	272
0.3%	120	6	0	GET	^/api/posts/[^/]+$	19.918	0.166	0.764	0.765	578
0.2%	80	0	0	GET	^/api/memories/[^/]+	14.464	0.181	0.404	0.404	804
0.2%	10	0	0	DELETE	^/api/posts/[^/]+$	11.405	1.141	7.456	7.456	0
0.2%	64	0	0	GET	^/api/users/[^/]+$	10.418	0.163	0.786	0.786	181
0.2%	63	1	0	GET	/api/me	9.955	0.158	0.345	0.345	294
0.1%	61	0	0	GET	^/api/posts/[^/]+/realmojis$	9.120	0.150	0.327	0.327	621
0.1%	38	0	0	GET	^/api/posts/[^/]+/comments$	5.278	0.139	0.309	0.309	756
0.0%	10	0	0	DELETE	^/api/realmojis/[^/]+$	2.965	0.296	0.763	0.763	0
0.0%	3	0	0	DELETE	^/api/comments/[^/]+$	1.220	0.407	0.697	0.697	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
30.6%	136	0	0	GET	/api/notifications	1081.337	7.951	9.088	9.089	176
17.8%	303	8	0	POST	/api/friend-requests	628.578	2.075	21.929	27.298	113
11.4%	152	3	0	POST	/api/posts	403.749	2.656	9.741	10.443	626
8.8%	44	0	0	PUT	^/api/follows/[^/]+$	309.271	7.029	25.847	25.847	0
5.8%	280	0	0	POST	/api/auth/login	205.770	0.735	6.534	7.280	250
4.5%	92	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	160.328	1.743	14.708	14.708	116
4.0%	969	5	0	GET	^/api/images/[^/]+	142.766	0.147	0.947	1.364	140377
2.6%	566	0	0	GET	^/api/feeds/[^/]+	92.374	0.163	0.615	0.786	1756
2.0%	43	8	0	POST	^/api/posts/[^/]+/comments$	71.419	1.661	11.209	11.209	277
2.0%	290	0	0	GET	/api/memories	69.940	0.241	1.142	2.361	14652
1.7%	202	0	0	GET	/api/posting-window/current	60.846	0.301	1.650	9.364	239
1.3%	24	0	0	GET	/api/me/export	45.477	1.895	4.756	4.756	2976721
1.1%	20	5	0	POST	^/api/posts/[^/]+/realmojis$	37.912	1.896	8.546	8.546	385
0.9%	39	0	0	POST	/api/auth/logout	31.196	0.800	7.466	7.466	0
0.9%	65	0	0	POST	/api/auth/register	30.121	0.463	2.590	2.590	229
0.7%	32	0	0	PUT	/api/me/profile-image	25.507	0.797	2.627	2.627	0
0.7%	150	0	0	GET	/api/users	23.967	0.160	0.485	0.872	224
0.5%	125	3	0	GET	^/api/posts/me/current$	19.301	0.154	0.333	0.554	621
0.5%	32	0	0	PATCH	/api/me	17.094	0.534	3.004	3.004	263
0.4%	92	0	0	GET	^/api/relationships/[^/]+	13.469	0.146	1.643	1.643	1226
0.3%	61	5	0	GET	^/api/posts/[^/]+$	10.009	0.164	0.518	0.518	560
0.2%	12	1	0	PATCH	^/api/posts/[^/]+$	8.685	0.724	1.344	1.344	613
0.2%	33	0	0	GET	^/api/memories/[^/]+	8.463	0.256	0.882	0.882	969
0.2%	43	0	0	GET	^/api/users/[^/]+$	6.437	0.150	0.284	0.284	187
0.2%	36	1	0	GET	/api/me	5.861	0.163	0.538	0.538	289
0.1%	32	0	0	GET	^/api/posts/[^/]+/comments$	5.154	0.161	0.333	0.333	1133
0.1%	28	0	0	GET	^/api/posts/[^/]+/realmojis$	4.599	0.164	0.251	0.251	478
0.1%	5	0	0	DELETE	^/api/posts/[^/]+$	4.599	0.920	2.076	2.076	0
0.1%	4	0	0	DELETE	^/api/comments/[^/]+$	1.907	0.477	1.186	1.186	0
0.0%	1	0	0	POST	/api/me/account-deletion	1.238	1.238	1.238	1.238	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 38673 クエリ / 81 種類 / 合計 224.458s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	45.7%	1825	102.613	56.23	622.49	654.1	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	36.4%	217	81.750	376.73	2686.54	7649.9	147.3	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	6.4%	2129	14.384	6.76	265.58	114.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	5.4%	5499	12.039	2.19	28.91	10.3	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.3%	1321	2.979	2.25	26.96	27.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.0%	1	2.340	2339.90	2339.90	29643.0	1199736.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.8%	1	1.795	1794.55	1794.55	20465.0	626386.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.5%	1	1.133	1132.94	1132.94	11528.0	529956.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	0.3%	207	0.773	3.73	24.68	34.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.3%	2770	0.627	0.23	14.70	20.5	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.2%	256	0.454	1.77	9.91	26.0	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.2%	3123	0.378	0.12	5.71	11.3	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
13	0.2%	2	0.346	173.06	189.48	220132.0	2590.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.1%	1059	0.274	0.26	11.66	21.7	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	2	0.265	132.32	256.03	16773.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: cdd0337cc59fc3b759890b7cdf7e9d8e795dd3cf
Type: cpu
Time: 2026-10-31 12:00:02 JST
Duration: 60.04s, Total samples = 58.85s (98.02%)
Showing nodes accounting for 47.66s, 80.99% of 58.85s total
Dropped 996 nodes (cum <= 0.29s)
Showing top 40 nodes out of 240
      flat  flat%   sum%        cum   cum%
     7.22s 12.27% 12.27%     22.29s 37.88%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     6.14s 10.43% 22.70%      9.95s 16.91%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     4.91s  8.34% 31.05%      5.11s  8.68%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     2.85s  4.84% 35.89%      3.95s  6.71%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     2.61s  4.44% 40.32%      2.61s  4.44%  runtime.memclrNoHeapPointers
     2.47s  4.20% 44.52%      2.47s  4.20%  internal/runtime/syscall/linux.Syscall6
     2.35s  3.99% 48.51%      8.67s 14.73%  runtime.scanObject
     2.17s  3.69% 52.20%      2.70s  4.59%  runtime.findObject
     1.75s  2.97% 55.17%      1.75s  2.97%  runtime.memmove
     1.73s  2.94% 58.11%      2.89s  4.91%  runtime.tryDeferToSpanScan
     1.63s  2.77% 60.88%      5.63s  9.57%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.41s  2.40% 63.28%      1.41s  2.40%  crypto/internal/fips140/sha256.blockAVX2
     1.29s  2.19% 65.47%      1.29s  2.19%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.18s  2.01% 67.48%      1.18s  2.01%  aeshashbody
     1.04s  1.77% 69.24%      1.10s  1.87%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.80s  1.36% 70.60%      0.80s  1.36%  runtime.spanClass.sizeclass (inline)
     0.72s  1.22% 71.83%      0.72s  1.22%  example.com/isureal-company/isureal.(*UserRow).active
     0.60s  1.02% 72.85%      0.60s  1.02%  memeqbody
     0.57s  0.97% 73.81%      1.99s  3.38%  runtime.mapaccess1_faststr
     0.36s  0.61% 74.43%      0.36s  0.61%  runtime.nextFreeFast (inline)
     0.35s  0.59% 75.02%      0.35s  0.59%  indexbody
     0.33s  0.56% 75.58%      0.71s  1.21%  runtime.greyobject
     0.32s  0.54% 76.13%      0.32s  0.54%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.29s  0.49% 76.62%      0.56s  0.95%  example.com/isureal-company/isureal.postLess
     0.26s  0.44% 77.06%      0.67s  1.14%  database/sql.convertAssignRows
     0.24s  0.41% 77.47%      0.32s  0.54%  runtime.(*mspan).markBitsForIndex (inline)
     0.24s  0.41% 77.88%      0.92s  1.56%  runtime.mallocgcSmallScanNoHeader
     0.21s  0.36% 78.23%      1.71s  2.91%  runtime.mapassign_faststr
     0.21s  0.36% 78.59%      0.59s  1.00%  runtime.scanObjectsSmall
     0.20s  0.34% 78.93%      0.31s  0.53%  runtime.heapArenaOf (inline)
     0.18s  0.31% 79.24%      1.02s  1.73%  runtime.scanSpan
     0.15s  0.25% 79.49%      0.45s  0.76%  runtime.mapaccess2_faststr
     0.14s  0.24% 79.73%      1.57s  2.67%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.13s  0.22% 79.95%      0.42s  0.71%  compress/flate.(*compressor).deflate
     0.13s  0.22% 80.17%      0.65s  1.10%  github.com/jackc/pgx/v5/pgtype.scanPlanString.Scan
     0.13s  0.22% 80.39%      3.14s  5.34%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.17% 80.56%      1.28s  2.18%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.09s  0.15% 80.71%      1.60s  2.72%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.09s  0.15% 80.87%      4.95s  8.41%  runtime.mallocgc
     0.07s  0.12% 80.99%      0.54s  0.92%  runtime.convTstring
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 59.5% / 100% | 43% | 0.1% | 6 | 316 MB | CPU が飽和 |
| isu2 (2 cpu) | 3.7% / 27% | 0% | 0.4% | 5 | 320 MB | run queue が積もっている |
| isu3 (2 cpu) | 2.6% / 14% | 0% | 0.5% | 2 | 247 MB |  |
| isu4 (2 cpu) | 61.7% / 100% | 45% | 7.2% | 44 | 94 MB | CPU が飽和 |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 100% / nginx 12% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 5% / nginx 1% |
| isu3 | isureal 3% / nginx 1% |
| isu4 | isureal 58% / postgres 47% |

## Diff from previous run

スコア: 56553 → 56314 (-239)

`773f053..a8d9e7d`

```
 webapp/go/limiter.go | 37 -------------------------------------
 webapp/go/server.go  | 10 +++++-----
 2 files changed, 5 insertions(+), 42 deletions(-)
```
