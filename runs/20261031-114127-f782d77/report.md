# run 20261031-114127-f782d77

## Summary

| | |
|---|---|
| score | 33567 |
| status | pass |
| ref | origin/main |
| sha | f782d7707145 |
| 所要時間 | 185s |
| 前回比 | +16495 (96.6%) vs 20261031-113441-6a7a33a |
| ベスト比 | +8766 (35.3%) vs 20261031-103811-cef5fe3 = 24801 **ベスト更新** |
| note | main: +#68 検証だけの JPEG デコーダ |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	134	262	261	144
europe	7	13	28	28	29	28
america	6	13	23	23	23	21
```

```
{"id":"5047","status":"done","portalStatus":"finished","passed":true,"score":33567,"raw":33567,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:43:17.277Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:43:20.329Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:43:23.226Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:43:26.242Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T02:43:29.414Z level=INFO msg="europe 2 日目: 13 人"
time=2026-10-31T02:43:32.269Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T02:43:35.316Z level=INFO msg="asia 3 日目: 134 人"
time=2026-10-31T02:43:38.135Z level=INFO msg="user moved" user=orfuivy-5373 from=europe to=asia
time=2026-10-31T02:43:38.282Z level=INFO msg="europe 3 日目: 28 人"
time=2026-10-31T02:43:41.390Z level=INFO msg="america 3 日目: 23 人"
time=2026-10-31T02:43:44.363Z level=INFO msg="asia 4 日目: 262 人"
time=2026-10-31T02:43:48.635Z level=INFO msg="europe 4 日目: 28 人"
time=2026-10-31T02:43:50.603Z level=INFO msg="america 4 日目: 23 人"
time=2026-10-31T02:43:51.609Z level=INFO msg="user moved" user=tsuisnu-5843 from=asia to=europe
time=2026-10-31T02:43:53.984Z level=INFO msg="asia 5 日目: 261 人"
time=2026-10-31T02:43:57.356Z level=INFO msg="europe 5 日目: 29 人"
time=2026-10-31T02:43:59.405Z level=INFO msg="america 5 日目: 23 人"
time=2026-10-31T02:44:08.291Z level=INFO msg="europe 最終日（グランプリ）: 28 人"
time=2026-10-31T02:44:08.298Z level=INFO msg="america 最終日（グランプリ）: 21 人"
time=2026-10-31T02:44:08.334Z level=INFO msg="asia 最終日（グランプリ）: 144 人"
time=2026-10-31T02:44:17.223Z level=INFO msg="グランプリ: アジア 65 人, ヨーロッパ 12 人, アメリカ 9 人"
{
  "pass": true,
  "score": 33567,
  "messages": [
    "sumeshi_6308: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ketor-6310: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "sechifu.6324: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "benhe-6333: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "wate-6341: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "zoeso.6346: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ormiamu-6352: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ivybenru.6361: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "niso-6371: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kinu-6373: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "noyu.6380: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "liateul_6390: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kiulne6397: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "noameta_6405: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "mezoeben_6410: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "orneis6420: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "wahe.6426: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "rimise.6431: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ulmoshi-6441: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ivyel_6446: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
30.6%	478	0	0	GET	/api/notifications	4343.164	9.086	14.334	14.338	178
18.2%	277	28	0	POST	/api/auth/register	2576.684	9.302	19.842	20.207	205
11.9%	579	8	0	POST	/api/posts	1690.760	2.920	9.516	10.204	636
9.3%	5794	892	0	GET	^/api/images/[^/]+	1313.365	0.227	2.811	8.520	109951
7.2%	3270	1817	0	POST	/api/auth/login	1026.006	0.314	3.159	9.008	127
3.7%	906	26	0	POST	/api/friend-requests	527.947	0.583	4.440	9.927	113
3.6%	550	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	507.608	0.923	6.318	8.266	116
3.0%	212	134	0	PUT	/api/me/profile-image	429.785	2.027	8.390	8.562	10
1.8%	3805	0	0	GET	^/api/feeds/[^/]+	255.766	0.067	0.356	0.720	2237
1.7%	1454	0	0	GET	/api/memories	239.833	0.165	1.416	2.256	10949
1.7%	246	11	0	POST	^/api/posts/[^/]+/realmojis$	239.049	0.972	4.841	6.357	451
1.3%	238	15	0	POST	^/api/posts/[^/]+/comments$	191.145	0.803	4.993	6.075	306
1.0%	722	0	0	GET	/api/posting-window/current	139.325	0.193	0.895	1.221	237
1.0%	111	0	0	GET	/api/me/export	137.767	1.241	6.617	6.740	1673607
0.8%	126	0	0	PUT	^/api/follows/[^/]+$	106.842	0.848	7.193	7.409	0
0.6%	255	0	0	POST	/api/auth/logout	91.272	0.358	6.052	7.068	0
0.4%	505	5	0	GET	^/api/posts/me/current$	52.006	0.103	0.492	0.571	637
0.3%	683	11	0	GET	^/api/posts/[^/]+$	37.823	0.055	0.299	0.506	626
0.3%	63	2	0	PATCH	^/api/posts/[^/]+$	37.704	0.598	3.305	3.305	628
0.2%	540	0	0	GET	^/api/relationships/[^/]+	33.386	0.062	0.270	0.676	1450
0.2%	453	0	0	GET	/api/users	29.139	0.064	0.437	0.780	245
0.2%	49	0	0	DELETE	^/api/realmojis/[^/]+$	23.279	0.475	3.504	3.504	0
0.1%	339	0	0	GET	^/api/users/[^/]+$	19.824	0.058	0.299	0.368	210
0.1%	24	0	0	DELETE	^/api/posts/[^/]+$	18.498	0.771	3.161	3.161	0
0.1%	293	0	0	GET	^/api/posts/[^/]+/realmojis$	17.143	0.059	0.238	0.289	638
0.1%	67	0	0	PATCH	/api/me	15.341	0.229	1.672	1.672	269
0.1%	248	0	0	GET	^/api/memories/[^/]+	12.298	0.050	0.458	0.601	724
0.1%	1	0	0	POST	/api/initialize	12.218	12.218	12.218	12.218	90
0.1%	217	0	0	GET	^/api/posts/[^/]+/comments$	11.110	0.051	0.231	0.269	835
0.1%	45	0	0	DELETE	^/api/comments/[^/]+$	10.787	0.240	3.117	3.117	0
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.2%	90	1	0	GET	/api/notifications	749.614	8.329	12.324	12.324	175
14.6%	97	1	0	POST	/api/posts	293.336	3.024	9.598	9.598	632
8.4%	164	5	0	POST	/api/friend-requests	168.990	1.030	5.831	5.936	112
6.2%	144	0	0	POST	/api/auth/login	124.389	0.864	5.866	7.121	251
5.0%	106	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	101.705	0.959	6.120	6.244	116
4.8%	505	0	0	GET	^/api/feeds/[^/]+	95.758	0.190	0.568	0.856	1879
4.4%	624	3	0	GET	^/api/images/[^/]+	89.428	0.143	0.940	1.321	169678
4.0%	27	0	0	POST	/api/auth/register	81.396	3.015	7.273	7.273	228
2.4%	195	0	0	GET	/api/memories	49.100	0.252	1.386	3.939	9750
1.6%	124	0	0	GET	/api/posting-window/current	32.015	0.258	0.780	0.988	239
1.3%	13	0	0	GET	/api/me/export	27.048	2.081	7.064	7.064	2032108
1.2%	16	0	0	PUT	^/api/follows/[^/]+$	23.356	1.460	5.825	5.825	0
1.1%	14	0	0	PUT	/api/me/profile-image	23.105	1.650	4.178	4.178	0
1.1%	25	1	0	POST	^/api/posts/[^/]+/comments$	21.356	0.854	5.953	5.953	359
1.0%	38	0	0	POST	/api/auth/logout	21.026	0.553	3.550	3.550	0
0.8%	83	6	0	GET	^/api/posts/me/current$	16.166	0.195	0.482	0.482	594
0.7%	94	4	0	GET	^/api/posts/[^/]+$	14.923	0.159	0.497	0.497	593
0.6%	79	0	0	GET	^/api/relationships/[^/]+	12.327	0.156	0.506	0.506	2065
0.6%	65	0	0	GET	/api/users	11.966	0.184	0.560	0.560	245
0.5%	52	0	0	GET	^/api/posts/[^/]+/realmojis$	9.498	0.183	0.850	0.850	595
0.5%	18	0	0	POST	^/api/posts/[^/]+/realmojis$	9.238	0.513	1.568	1.568	422
0.4%	10	0	0	PATCH	/api/me	8.892	0.889	3.321	3.321	295
0.4%	8	1	0	PATCH	^/api/posts/[^/]+$	8.466	1.058	5.171	5.171	566
0.3%	38	0	0	GET	^/api/posts/[^/]+/comments$	6.359	0.167	0.480	0.480	905
0.2%	30	0	0	GET	^/api/users/[^/]+$	4.785	0.160	0.361	0.361	203
0.2%	19	2	0	GET	/api/me	3.725	0.196	0.440	0.440	259
0.1%	9	0	0	GET	^/api/memories/[^/]+	2.166	0.241	0.762	0.762	2274
0.1%	1	0	0	POST	/api/me/account-deletion	1.043	1.043	1.043	1.043	0
0.1%	3	0	0	DELETE	^/api/posts/[^/]+$	1.034	0.345	0.478	0.478	0
0.0%	2	0	0	DELETE	^/api/realmojis/[^/]+$	0.891	0.446	0.512	0.512	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.4%	74	0	0	GET	/api/notifications	598.533	8.088	9.234	9.234	180
15.0%	76	1	0	POST	/api/posts	234.488	3.085	9.868	9.868	620
7.2%	140	4	0	POST	/api/friend-requests	112.485	0.803	5.024	6.385	113
6.6%	73	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	102.368	1.402	9.168	9.168	115
5.8%	22	0	0	POST	/api/auth/register	90.420	4.110	15.768	15.768	230
5.1%	113	0	0	POST	/api/auth/login	79.695	0.705	4.408	4.520	252
4.4%	398	3	0	GET	^/api/images/[^/]+	68.817	0.173	0.951	1.250	158911
3.1%	249	0	0	GET	^/api/feeds/[^/]+	48.760	0.196	0.531	0.930	1572
3.1%	146	0	0	GET	/api/memories	48.131	0.330	1.955	3.888	11636
1.7%	95	0	0	GET	/api/posting-window/current	26.096	0.275	1.708	1.708	240
1.5%	10	0	0	GET	/api/me/export	23.665	2.366	6.630	6.630	3048388
1.2%	18	5	0	POST	^/api/posts/[^/]+/comments$	18.822	1.046	4.120	4.120	234
0.9%	66	4	0	GET	^/api/posts/me/current$	13.932	0.211	0.566	0.566	593
0.9%	68	0	0	GET	/api/users	13.296	0.196	0.509	0.509	261
0.7%	9	0	0	PUT	^/api/follows/[^/]+$	10.686	1.187	3.676	3.676	0
0.7%	55	0	0	GET	^/api/relationships/[^/]+	10.434	0.190	0.486	0.486	1812
0.6%	17	0	0	POST	/api/auth/logout	8.613	0.507	2.238	2.238	0
0.4%	4	0	0	DELETE	^/api/posts/[^/]+$	6.968	1.742	3.581	3.581	0
0.4%	7	1	0	POST	^/api/posts/[^/]+/realmojis$	6.050	0.864	3.220	3.220	452
0.3%	31	0	0	GET	^/api/posts/[^/]+$	5.094	0.164	0.419	0.419	614
0.3%	7	0	0	PUT	/api/me/profile-image	4.871	0.696	1.818	1.818	0
0.3%	11	1	0	GET	/api/me	3.951	0.359	2.409	2.409	251
0.2%	8	0	0	PATCH	^/api/posts/[^/]+$	3.867	0.483	1.776	1.776	632
0.2%	2	0	0	PUT	^/api/blocks/[^/]+$	3.446	1.723	3.341	3.341	0
0.2%	7	0	0	PATCH	/api/me	3.042	0.435	0.927	0.927	235
0.2%	20	0	0	GET	^/api/memories/[^/]+	2.925	0.146	0.308	0.308	790
0.2%	17	0	0	GET	^/api/posts/[^/]+/comments$	2.699	0.159	0.309	0.309	1056
0.1%	17	0	0	GET	^/api/users/[^/]+$	2.331	0.137	0.226	0.226	203
0.1%	12	0	0	GET	^/api/posts/[^/]+/realmojis$	1.878	0.157	0.304	0.304	824
0.1%	2	0	0	DELETE	^/api/comments/[^/]+$	1.171	0.586	0.705	0.705	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 23014 クエリ / 79 種類 / 合計 70.405s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	53.5%	134	37.638	280.88	1148.11	5726.4	116.0	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	16.0%	1100	11.282	10.26	127.55	593.1	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	9.1%	3624	6.372	1.76	9.01	11.0	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	7.0%	889	4.901	5.51	314.86	115.2	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	3.4%	1	2.424	2423.92	2423.92	29434.0	1199605.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	2.2%	1	1.558	1558.36	1558.36	20022.0	625651.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	1.7%	742	1.173	1.58	6.22	29.7	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
8	1.5%	1	1.078	1077.75	1077.75	11528.0	529810.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	0.9%	259	0.640	2.47	6.82	34.0	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.5%	260	0.336	1.29	4.45	24.0	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.5%	2	0.328	163.89	175.48	214278.5	2520.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
12	0.4%	1163	0.263	0.23	6.43	22.4	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
13	0.4%	2	0.262	130.95	253.57	16767.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
14	0.3%	1	0.223	222.85	222.85	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
15	0.3%	1	0.209	209.29	209.29	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…

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
Build ID: 67ea584c04b810b5e8d65f3fb5524b8641b26c42
Type: cpu
Time: 2026-10-31 11:42:59 JST
Duration: 60.08s, Total samples = 83.57s (139.10%)
Showing nodes accounting for 73.15s, 87.53% of 83.57s total
Dropped 912 nodes (cum <= 0.42s)
Showing top 40 nodes out of 189
      flat  flat%   sum%        cum   cum%
    21.90s 26.21% 26.21%     21.90s 26.21%  crypto/internal/fips140/sha256.blockAVX2
    11.02s 13.19% 39.39%     11.02s 13.19%  runtime.spanClass.sizeclass (inline)
     4.40s  5.27% 44.66%     14.59s 17.46%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     4.14s  4.95% 49.61%      6.52s  7.80%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     3.49s  4.18% 53.79%      3.65s  4.37%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     3.47s  4.15% 57.94%     15.13s 18.10%  runtime.tryDeferToSpanScan
     3.17s  3.79% 61.73%      3.17s  3.79%  runtime.memmove
     2.18s  2.61% 64.34%      2.18s  2.61%  runtime.memclrNoHeapPointers
     1.78s  2.13% 66.47%      1.78s  2.13%  internal/runtime/syscall/linux.Syscall6
     1.76s  2.11% 68.58%      2.52s  3.02%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     1.52s  1.82% 70.40%     12.01s 14.37%  runtime.scanObject
     1.39s  1.66% 72.06%      5.52s  6.61%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.37s  1.64% 73.70%      6.63s  7.93%  runtime.scanObjectsSmall
     1.19s  1.42% 75.12%     24.56s 29.39%  crypto/internal/fips140/sha256.(*Digest).checkSum
     1.18s  1.41% 76.53%     24.16s 28.91%  crypto/internal/fips140/sha256.(*Digest).Write
     1.10s  1.32% 77.85%      1.10s  1.32%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.98s  1.17% 79.02%      0.98s  1.17%  aeshashbody
     0.73s  0.87% 79.90%     25.75s 30.81%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.71s  0.85% 80.75%      0.76s  0.91%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.64s  0.77% 81.51%      0.64s  0.77%  memeqbody
     0.62s  0.74% 82.25%      8.01s  9.58%  runtime.scanSpan
     0.60s  0.72% 82.97%     27.36s 32.74%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.47s  0.56% 83.53%      0.47s  0.56%  runtime.extractHeapBitsSmall
     0.41s  0.49% 84.03%      0.54s  0.65%  example.com/isureal-company/isureal.postLess
     0.38s  0.45% 84.48%      2.15s  2.57%  runtime.mapaccess1_faststr
     0.36s  0.43% 84.91%      1.29s  1.54%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.31s  0.37% 85.28%     29.52s 35.32%  example.com/isureal-company/isureal.imageSignature
     0.29s  0.35% 85.63%      1.04s  1.24%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.25s   0.3% 85.93%     22.15s 26.50%  crypto/internal/fips140/sha256.block
     0.19s  0.23% 86.16%     11.64s 13.93%  runtime.mallocgc
     0.18s  0.22% 86.37%      0.90s  1.08%  runtime.mallocgcSmallScanNoHeader
     0.17s   0.2% 86.57%      0.72s  0.86%  database/sql.convertAssignRows
     0.17s   0.2% 86.78%      1.29s  1.54%  runtime.mapassign_faststr
     0.16s  0.19% 86.97%      0.63s  0.75%  runtime.scanObjectSmall
     0.12s  0.14% 87.11%      2.94s  3.52%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.08s 0.096% 87.21%      0.49s  0.59%  runtime.wbBufFlush1
     0.07s 0.084% 87.29%      0.51s  0.61%  crypto/internal/fips140/hmac.(*HMAC).Write
     0.07s 0.084% 87.38%      1.23s  1.47%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.07s 0.084% 87.46%      0.47s  0.56%  runtime.convTstring
     0.06s 0.072% 87.53%     18.24s 21.83%  example.com/isureal-company/isureal.listMemories
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 72.9% / 100% | 59% | 0.1% | 5 | 103 MB | CPU が飽和 |
| isu2 (2 cpu) | 1.2% / 4% | 0% | 0.2% | 2 | 497 MB |  |
| isu3 (2 cpu) | 1.4% / 4% | 0% | 0.3% | 2 | 240 MB |  |
| isu4 (2 cpu) | 30.8% / 79% | 0% | 16.9% | 5 | 106 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 128% / nginx 10% / journalctl 1% |
| isu2 | isureal 1% |
| isu3 | isureal 1% |
| isu4 | postgres 24% |

## Diff from previous run

スコア: 17072 → 33567 (+16495)

`6a7a33a..f782d77`

```
 webapp/go/jpegcheck/check_test.go | 130 ++++++++
 webapp/go/jpegcheck/huffman.go    | 249 ++++++++++++++
 webapp/go/jpegcheck/reader.go     | 669 ++++++++++++++++++++++++++++++++++++++
 webapp/go/jpegcheck/scan.go       | 481 +++++++++++++++++++++++++++
 webapp/go/post_api.go             |   6 +-
 webapp/go/profile_image_api.go    |   3 +-
 webapp/go/realmoji_api.go         |   6 +-
 7 files changed, 1537 insertions(+), 7 deletions(-)
```
