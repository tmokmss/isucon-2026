# run 20261031-115658-52e37f7

## Summary

| | |
|---|---|
| score | 69898 |
| status | pass |
| ref | offload-jpeg |
| sha | 52e37f747fca |
| 所要時間 | 168s |
| 前回比 | +36331 (108.2%) vs 20261031-114127-f782d77 |
| ベスト比 | +36331 (108.2%) vs 20261031-114127-f782d77 = 33567 **ベスト更新** |
| note | offload-jpeg-to-eu-us |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	141	397	414	260
europe	7	17	44	106	105	45
america	6	18	40	72	72	55
```

```
{"id":"6102","status":"done","portalStatus":"finished","passed":true,"score":69898,"raw":69898,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:58:25.456Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:58:28.408Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:58:31.408Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:58:34.450Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T02:58:34.924Z level=INFO msg="user moved" user=taisru-5161 from=asia to=america
time=2026-10-31T02:58:37.408Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T02:58:37.613Z level=INFO msg="user moved" user=taisru-5161 from=america to=asia
time=2026-10-31T02:58:38.310Z level=INFO msg="user moved" user=mimeis.5347 from=asia to=america
time=2026-10-31T02:58:40.409Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T02:58:43.438Z level=INFO msg="asia 3 日目: 141 人"
time=2026-10-31T02:58:45.447Z level=INFO msg="user moved" user=taisru-5161 from=asia to=america
time=2026-10-31T02:58:46.429Z level=INFO msg="europe 3 日目: 44 人"
time=2026-10-31T02:58:46.467Z level=INFO msg="user moved" user=elleoul.5923 from=asia to=europe
time=2026-10-31T02:58:47.439Z level=INFO msg="user moved" user=rechimi.6178 from=asia to=europe
time=2026-10-31T02:58:47.675Z level=INFO msg="user moved" user=kueva.5619 from=asia to=europe
time=2026-10-31T02:58:49.408Z level=INFO msg="america 3 日目: 40 人"
time=2026-10-31T02:58:52.585Z level=INFO msg="asia 4 日目: 397 人"
time=2026-10-31T02:58:55.540Z level=INFO msg="europe 4 日目: 106 人"
time=2026-10-31T02:58:58.471Z level=INFO msg="america 4 日目: 72 人"
time=2026-10-31T02:58:58.746Z level=INFO msg="user moved" user=rechimi.6178 from=europe to=america
time=2026-10-31T02:59:01.552Z level=INFO msg="asia 5 日目: 414 人"
time=2026-10-31T02:59:04.556Z level=INFO msg="europe 5 日目: 105 人"
time=2026-10-31T02:59:07.269Z level=INFO msg="user moved" user=taisru-5161 from=america to=asia
time=2026-10-31T02:59:07.461Z level=INFO msg="america 5 日目: 72 人"
time=2026-10-31T02:59:16.451Z level=INFO msg="europe 最終日（グランプリ）: 45 人"
time=2026-10-31T02:59:16.565Z level=INFO msg="asia 最終日（グランプリ）: 260 人"
time=2026-10-31T02:59:16.690Z level=INFO msg="america 最終日（グランプリ）: 55 人"
time=2026-10-31T02:59:25.413Z level=INFO msg="グランプリ: アジア 97 人, ヨーロッパ 11 人, アメリカ 12 人"
{
  "pass": true,
  "score": 69898,
  "messages": [
    "niko_7247: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "hahe-7258: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ulzoe_7261: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "hane_7274: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "torkois7279: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "heru_7287: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "rois_7293: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "sumar_7296: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "martor.7306: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "somi-7315: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kehe-7317: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "nitaeva-7329: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "liawahe_7337: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kataan7344: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "sahi_7350: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "muzoe7358: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "tevinsu7363: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "yunoamar-7368: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "hosol.7373: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "zoevin7384: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.7%	685	1	0	GET	/api/notifications	6157.621	8.989	15.032	15.180	170
20.7%	435	52	0	POST	/api/auth/register	3669.356	8.435	18.412	18.440	200
12.5%	899	18	0	POST	/api/posts	2226.138	2.476	9.354	10.276	635
8.9%	12269	537	0	GET	^/api/images/[^/]+	1578.657	0.129	1.617	9.847	83668
5.8%	4820	1014	0	POST	/api/auth/login	1026.287	0.213	2.293	9.755	201
2.4%	270	113	0	PUT	/api/me/profile-image	428.068	1.585	9.765	9.792	6
2.4%	3102	0	0	GET	/api/memories	427.962	0.138	0.749	1.632	12823
2.3%	1667	59	0	POST	/api/friend-requests	415.887	0.249	2.643	8.339	112
2.0%	1116	13	0	POST	^/api/friend-requests/[^/]+/[^/]+$	359.749	0.322	2.632	5.000	115
1.6%	308	0	0	GET	/api/me/export	284.301	0.923	7.959	10.552	1842465
1.2%	355	9	0	POST	^/api/posts/[^/]+/realmojis$	218.871	0.617	2.751	3.837	469
1.0%	6857	0	0	GET	^/api/feeds/[^/]+	170.676	0.025	0.185	1.351	3138
0.9%	1195	0	0	GET	/api/posting-window/current	167.447	0.140	1.650	1.902	237
0.7%	242	12	0	POST	^/api/posts/[^/]+/comments$	121.674	0.503	2.473	3.522	318
0.7%	202	0	0	PUT	^/api/follows/[^/]+$	116.210	0.575	4.762	4.808	0
0.6%	748	0	0	POST	/api/auth/logout	107.350	0.144	2.358	9.375	0
0.2%	710	0	0	GET	^/api/memories/[^/]+	34.103	0.048	0.287	0.527	704
0.2%	741	12	0	GET	^/api/posts/me/current$	33.994	0.046	0.656	0.980	637
0.1%	36	0	0	DELETE	^/api/posts/[^/]+$	26.121	0.726	5.249	5.249	0
0.1%	89	1	0	PATCH	^/api/posts/[^/]+$	26.114	0.293	6.142	6.142	635
0.1%	1298	8	0	GET	^/api/posts/[^/]+$	26.091	0.020	0.119	0.882	627
0.1%	70	0	0	DELETE	^/api/realmojis/[^/]+$	19.191	0.274	0.833	0.833	0
0.1%	969	0	0	GET	^/api/relationships/[^/]+	17.536	0.018	0.141	1.347	1280
0.1%	130	0	0	PATCH	/api/me	14.284	0.110	0.738	2.452	278
0.1%	732	0	0	GET	/api/users	13.970	0.019	0.101	0.321	232
0.1%	1	0	0	POST	/api/initialize	12.886	12.886	12.886	12.886	90
0.1%	605	0	0	GET	^/api/users/[^/]+$	11.421	0.019	0.097	0.419	208
0.1%	652	0	0	GET	^/api/posts/[^/]+/realmojis$	11.306	0.017	0.084	0.187	682
0.1%	450	0	0	GET	^/api/posts/[^/]+/comments$	8.916	0.020	0.099	0.168	802
0.0%	10	0	0	DELETE	^/api/follows/[^/]+$	8.369	0.837	2.165	2.165	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.7%	173	1	0	GET	/api/notifications	1511.011	8.734	11.851	11.852	177
21.8%	101	20	0	POST	/api/auth/register	974.385	9.647	18.501	18.511	183
10.9%	209	3	0	POST	/api/posts	486.027	2.325	9.553	9.767	643
6.2%	2042	9	0	GET	^/api/images/[^/]+	277.334	0.136	0.975	1.641	125858
5.8%	558	0	0	POST	/api/auth/login	257.802	0.462	7.358	9.357	252
5.0%	1545	0	0	GET	^/api/feeds/[^/]+	224.841	0.146	0.343	0.615	4166
3.0%	530	0	0	GET	/api/memories	132.703	0.250	0.768	2.134	16243
2.0%	339	9	0	POST	/api/friend-requests	90.345	0.267	2.361	4.390	113
1.7%	178	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	76.010	0.427	4.023	9.534	116
1.3%	102	6	0	POST	^/api/posts/[^/]+/comments$	59.860	0.587	3.930	5.103	370
1.3%	250	0	0	GET	/api/posting-window/current	57.282	0.229	1.157	7.100	239
1.2%	41	0	0	GET	/api/me/export	54.437	1.328	7.309	7.309	1800362
1.0%	322	7	0	GET	^/api/posts/[^/]+$	45.760	0.142	0.311	0.398	606
0.8%	89	0	0	POST	/api/auth/logout	37.080	0.417	8.206	8.206	0
0.6%	208	0	0	GET	^/api/relationships/[^/]+	27.792	0.134	0.289	0.358	1186
0.6%	175	0	0	GET	^/api/posts/me/current$	26.787	0.153	0.324	0.349	651
0.5%	152	0	0	GET	^/api/users/[^/]+$	20.915	0.138	0.291	0.338	196
0.4%	109	0	0	GET	^/api/memories/[^/]+	19.069	0.175	0.606	0.856	860
0.4%	144	0	0	GET	/api/users	18.921	0.131	0.236	0.285	246
0.4%	112	0	0	GET	^/api/posts/[^/]+/comments$	16.402	0.146	0.358	0.465	1073
0.4%	116	0	0	GET	^/api/posts/[^/]+/realmojis$	16.114	0.139	0.317	0.328	728
0.1%	25	0	0	DELETE	^/api/comments/[^/]+$	6.486	0.259	1.678	1.678	0
0.1%	25	0	0	PUT	/api/me/profile-image	6.022	0.241	0.495	0.495	0
0.1%	8	1	0	POST	^/api/posts/[^/]+/realmojis$	5.884	0.735	2.316	2.316	378
0.1%	7	0	0	PUT	^/api/blocks/[^/]+$	5.661	0.809	4.032	4.032	0
0.1%	31	0	0	PUT	^/api/follows/[^/]+$	5.144	0.166	1.155	1.155	0
0.1%	7	0	0	DELETE	^/api/posts/[^/]+$	5.061	0.723	1.746	1.746	0
0.1%	14	1	0	PATCH	^/api/posts/[^/]+$	3.984	0.285	0.825	0.825	648
0.1%	24	0	0	PATCH	/api/me	3.675	0.153	0.295	0.295	279
0.1%	30	2	0	GET	/api/me	3.653	0.122	0.248	0.248	277
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.3%	166	1	0	GET	/api/notifications	1276.781	7.691	9.222	9.222	165
14.4%	69	1	0	POST	/api/auth/register	491.888	7.129	16.439	16.439	226
13.5%	189	4	0	POST	/api/posts	460.820	2.438	9.528	9.682	637
5.4%	331	0	0	POST	/api/auth/login	183.180	0.553	7.047	8.761	253
5.3%	1159	11	0	GET	^/api/images/[^/]+	182.462	0.157	0.958	1.303	150664
4.0%	313	3	0	POST	/api/friend-requests	135.303	0.432	6.114	8.318	114
3.4%	799	0	0	GET	^/api/feeds/[^/]+	114.910	0.144	0.434	0.626	4086
2.9%	187	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	98.058	0.524	5.535	7.554	115
2.5%	372	0	0	GET	/api/memories	86.599	0.233	0.900	2.787	13857
2.1%	39	0	0	GET	/api/me/export	71.885	1.843	5.071	5.071	3479813
1.8%	254	0	0	GET	/api/posting-window/current	62.621	0.247	1.876	1.909	240
1.4%	96	6	0	POST	^/api/posts/[^/]+/comments$	47.401	0.494	5.903	5.903	332
0.8%	38	3	0	POST	^/api/posts/[^/]+/realmojis$	25.876	0.681	3.228	3.228	382
0.7%	57	0	0	POST	/api/auth/logout	25.654	0.450	4.193	4.193	0
0.7%	154	0	0	GET	^/api/posts/me/current$	23.322	0.151	0.487	0.795	649
0.6%	142	0	0	GET	/api/users	20.222	0.142	0.356	0.482	234
0.5%	137	9	0	GET	^/api/posts/[^/]+$	18.274	0.133	0.342	0.361	586
0.5%	145	0	0	GET	^/api/relationships/[^/]+	18.155	0.125	0.289	0.354	1814
0.3%	12	0	0	DELETE	^/api/posts/[^/]+$	11.586	0.966	7.109	7.109	0
0.3%	50	0	0	GET	^/api/memories/[^/]+	9.740	0.195	1.296	1.296	711
0.3%	30	0	0	PUT	/api/me/profile-image	9.664	0.322	0.911	0.911	0
0.2%	60	0	0	GET	^/api/users/[^/]+$	7.866	0.131	0.342	0.342	197
0.2%	43	0	0	PUT	^/api/follows/[^/]+$	7.642	0.178	0.885	0.885	0
0.2%	55	1	0	GET	^/api/posts/[^/]+/comments$	6.839	0.124	0.253	0.253	1258
0.1%	35	1	0	GET	/api/me	4.538	0.130	0.284	0.284	290
0.1%	25	0	0	PATCH	/api/me	4.282	0.171	0.365	0.365	268
0.1%	30	0	0	GET	^/api/posts/[^/]+/realmojis$	4.123	0.137	0.487	0.487	840
0.1%	10	0	0	PATCH	^/api/posts/[^/]+$	3.562	0.356	1.680	1.680	631
0.1%	15	0	0	DELETE	^/api/comments/[^/]+$	2.973	0.198	0.470	0.470	0
0.0%	1	0	0	DELETE	/api/me/profile-image	1.465	1.465	1.465	1.465	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 51529 クエリ / 82 種類 / 合計 300.241s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	45.7%	388	137.260	353.76	2186.15	7353.2	127.9	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	34.5%	1878	103.543	55.13	438.83	619.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	8.3%	9710	24.775	2.55	24.38	9.7	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	6.4%	2600	19.303	7.42	319.61	113.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	1.1%	1276	3.311	2.59	89.04	27.2	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.8%	1	2.314	2313.57	2313.57	29618.0	1199930.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.5%	1	1.635	1634.63	1634.63	20406.0	626105.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.5%	388	1.483	3.82	17.19	31.5	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.4%	1	1.093	1093.09	1093.09	11528.0	530029.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
10	0.2%	416	0.744	1.79	8.78	21.1	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.2%	4695	0.636	0.14	54.20	11.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
12	0.2%	2231	0.534	0.24	10.32	20.3	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
13	0.2%	2	0.503	251.35	331.59	217496.0	2559.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.1%	1428	0.382	0.27	13.50	21.1	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	2	0.263	131.51	254.55	16772.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 3
SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: bf0afcfcb69c9e2fbf2526b56f1be4bae074a6fe
Type: cpu
Time: 2026-10-31 11:58:07 JST
Duration: 60.01s, Total samples = 54.87s (91.44%)
Showing nodes accounting for 42.18s, 76.87% of 54.87s total
Dropped 1016 nodes (cum <= 0.27s)
Showing top 40 nodes out of 268
      flat  flat%   sum%        cum   cum%
     8.92s 16.26% 16.26%      8.92s 16.26%  crypto/internal/fips140/sha256.blockAVX2
     3.01s  5.49% 21.74%      3.01s  5.49%  internal/runtime/syscall/linux.Syscall6
     2.54s  4.63% 26.37%      2.54s  4.63%  runtime.memmove
     2.32s  4.23% 30.60%      2.32s  4.23%  runtime.memclrNoHeapPointers
     2.21s  4.03% 34.63%      8.28s 15.09%  runtime.scanObject
     2.05s  3.74% 38.36%      3.21s  5.85%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.99s  3.63% 41.99%      2.65s  4.83%  runtime.findObject
     1.72s  3.13% 45.12%      5.67s 10.33%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.66s  3.03% 48.15%      6.32s 11.52%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.57s  2.86% 51.01%      2.97s  5.41%  runtime.tryDeferToSpanScan
     1.52s  2.77% 53.78%      1.52s  2.77%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.40s  2.55% 56.33%      1.40s  2.55%  aeshashbody
     1.34s  2.44% 58.78%      1.44s  2.62%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     1.03s  1.88% 60.65%      1.03s  1.88%  runtime.spanClass.sizeclass (inline)
     0.99s  1.80% 62.46%      1.25s  2.28%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.64s  1.17% 63.62%      0.64s  1.17%  memeqbody
     0.57s  1.04% 64.66%      2.11s  3.85%  runtime.mapaccess1_faststr
     0.55s  1.00% 65.66%      0.55s  1.00%  example.com/isureal-company/isureal.(*UserRow).active
     0.48s  0.87% 66.54%      9.99s 18.21%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.47s  0.86% 67.40%      9.80s 17.86%  crypto/internal/fips140/sha256.(*Digest).Write
     0.42s  0.77% 68.16%      0.65s  1.18%  example.com/isureal-company/isureal.postLess
     0.37s  0.67% 68.84%      0.37s  0.67%  runtime.nextFreeFast (inline)
     0.34s  0.62% 69.46%      0.34s  0.62%  indexbody
     0.33s   0.6% 70.06%      0.33s   0.6%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.31s  0.56% 70.62%      0.67s  1.22%  runtime.greyobject
     0.30s  0.55% 71.17%      1.14s  2.08%  runtime.mallocgcSmallScanNoHeader
     0.29s  0.53% 71.70%      0.72s  1.31%  database/sql.convertAssignRows
     0.29s  0.53% 72.23%      0.30s  0.55%  runtime.spanOf (inline)
     0.28s  0.51% 72.74%      0.49s  0.89%  compress/flate.(*compressor).deflate
     0.27s  0.49% 73.23%     10.36s 18.88%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.27s  0.49% 73.72%      0.29s  0.53%  runtime.(*mspan).writeHeapBitsSmall
     0.26s  0.47% 74.19%     11.03s 20.10%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.24s  0.44% 74.63%      0.66s  1.20%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.24s  0.44% 75.07%      0.29s  0.53%  runtime.typePointers.next
     0.21s  0.38% 75.45%      0.30s  0.55%  runtime.heapArenaOf (inline)
     0.20s  0.36% 75.82%      0.66s  1.20%  runtime.scanObjectsSmall
     0.15s  0.27% 76.09%      5.72s 10.42%  runtime.mallocgc
     0.15s  0.27% 76.36%      1.95s  3.55%  runtime.mapassign_faststr
     0.14s  0.26% 76.62%      0.56s  1.02%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.14s  0.26% 76.87%      3.24s  5.90%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 53.9% / 100% | 39% | 0.0% | 7 | 167 MB | CPU が飽和 |
| isu2 (2 cpu) | 9.5% / 28% | 0% | 0.5% | 3 | 363 MB |  |
| isu3 (2 cpu) | 8.1% / 35% | 0% | 0.5% | 3 | 217 MB |  |
| isu4 (2 cpu) | 60.4% / 100% | 49% | 6.6% | 19 | 110 MB | CPU が飽和 |

**isu1（CPU が飽和、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 89% / nginx 11% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 17% / nginx 1% |
| isu3 | isureal 13% / nginx 1% |
| isu4 | postgres 53% / isureal 50% |

## Diff from previous run

スコア: 33567 → 69898 (+36331)

`f782d77..52e37f7`

```
 conf/isu1/home/isucon/env.sh                    |   1 +
 conf/isu4/etc/systemd/system/isureal-go.service |  21 ++
 webapp/go/cmd/isureal/main.go                   |   6 +
 webapp/go/export_api.go                         |   9 +
 webapp/go/image_signature.go                    |  33 ++-
 webapp/go/memory_api.go                         |  18 ++
 webapp/go/offload.go                            | 259 ++++++++++++++++++++++++
 webapp/go/post_api.go                           |   4 +-
 webapp/go/profile_image_api.go                  |   3 +-
 webapp/go/realmoji_api.go                       |   4 +-
 webapp/go/store.go                              |  76 ++++++-
 11 files changed, 413 insertions(+), 21 deletions(-)
```
