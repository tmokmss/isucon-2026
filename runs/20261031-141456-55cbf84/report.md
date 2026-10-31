# run 20261031-141456-55cbf84

## Summary

| | |
|---|---|
| score | 134100 |
| status | pass |
| ref | gp-measure-2 |
| sha | 55cbf84f128a |
| 所要時間 | 248s |
| 前回比 | +10549 (8.5%) vs 20261031-140544-247167c |
| ベスト比 | +10549 (8.5%) vs 20261031-140544-247167c = 123551 **ベスト更新** |
| note | 計測用: 投稿・転送の段階別の時間（main e42c64c） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	147	429	999	1000
europe	8	26	66	179	208	207
america	6	15	45	101	116	116
```

```
{"id":"16657","status":"done","portalStatus":"finished","passed":true,"score":134100,"raw":134100,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:17:36.414Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:17:38.943Z level=INFO msg="user moved" user=reantor-5042 from=asia to=europe
time=2026-10-31T05:17:39.385Z level=INFO msg="europe 1 日目: 8 人"
time=2026-10-31T05:17:42.386Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:17:45.006Z level=INFO msg="user moved" user=tetorul-5133 from=america to=asia
time=2026-10-31T05:17:45.411Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T05:17:45.659Z level=INFO msg="user moved" user=saho.5155 from=asia to=europe
time=2026-10-31T05:17:46.312Z level=INFO msg="user moved" user=tetorul-5133 from=asia to=europe
time=2026-10-31T05:17:46.817Z level=INFO msg="user moved" user=keshi.5229 from=asia to=europe
time=2026-10-31T05:17:48.385Z level=INFO msg="europe 2 日目: 26 人"
time=2026-10-31T05:17:51.385Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T05:17:52.247Z level=INFO msg="user moved" user=reantor-5042 from=europe to=america
time=2026-10-31T05:17:54.414Z level=INFO msg="asia 3 日目: 147 人"
time=2026-10-31T05:17:57.396Z level=INFO msg="europe 3 日目: 66 人"
time=2026-10-31T05:17:59.326Z level=INFO msg="user moved" user=tetorul-5133 from=europe to=america
time=2026-10-31T05:18:00.584Z level=INFO msg="america 3 日目: 45 人"
time=2026-10-31T05:18:03.430Z level=INFO msg="asia 4 日目: 429 人"
time=2026-10-31T05:18:03.852Z level=INFO msg="user moved" user=saho.5155 from=europe to=america
time=2026-10-31T05:18:06.408Z level=INFO msg="europe 4 日目: 179 人"
time=2026-10-31T05:18:08.182Z level=INFO msg="user moved" user=nator8170 from=asia to=america
time=2026-10-31T05:18:08.263Z level=INFO msg="user moved" user=seivyyu8074 from=asia to=europe
time=2026-10-31T05:18:09.455Z level=INFO msg="america 4 日目: 101 人"
time=2026-10-31T05:18:10.673Z level=INFO msg="user moved" user=yameel_8171 from=asia to=america
time=2026-10-31T05:18:12.547Z level=INFO msg="asia 5 日目: 999 人"
time=2026-10-31T05:18:15.465Z level=INFO msg="europe 5 日目: 208 人"
time=2026-10-31T05:18:15.691Z level=INFO msg="user moved" user=mimimax-8425 from=europe to=asia
time=2026-10-31T05:18:18.527Z level=INFO msg="america 5 日目: 116 人"
time=2026-10-31T05:18:27.751Z level=INFO msg="america 最終日（グランプリ）: 116 人"
time=2026-10-31T05:18:27.855Z level=INFO msg="asia 最終日（グランプリ）: 1000 人"
time=2026-10-31T05:18:28.060Z level=INFO msg="europe 最終日（グランプリ）: 207 人"
time=2026-10-31T05:18:36.418Z level=INFO msg="グランプリ: アジア 590 人, ヨーロッパ 108 人, アメリカ 53 人"
{
  "pass": true,
  "score": 134100,
  "messages": [
    "miaeva13811: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "keneyo8720: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hani.13721: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sora_11591: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tokuno.13312: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "makuno_13706: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mera-11370: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "menoa12789: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ulmaxke.13302: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nisoul7705: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "zoena.12178: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yosa_13074: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ivysu-12744: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sutbul_mini_chilmyeonjo_1752: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "maxul-13132: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sora_11591: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "liarato_12119: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tsuraku.13086: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.6%	1612	0	0	GET	/api/notifications	14374.659	8.917	15.436	15.464	161
16.1%	2439	78	0	POST	/api/posts	5983.325	2.453	9.081	12.436	629
10.1%	5561	82	0	POST	/api/friend-requests	3754.121	0.675	5.198	11.151	114
8.4%	14631	484	0	GET	^/api/images/[^/]+	3129.578	0.214	3.200	8.867	152555
5.1%	2554	6	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1894.380	0.742	5.238	10.428	116
5.0%	676	125	0	PUT	/api/me/profile-image	1849.399	2.736	10.710	11.811	3
4.4%	5427	966	0	POST	/api/auth/login	1641.965	0.303	3.309	11.659	205
2.2%	483	33	0	POST	^/api/posts/[^/]+/realmojis$	833.266	1.725	8.898	11.617	442
1.4%	638	18	0	POST	^/api/posts/[^/]+/comments$	524.454	0.822	5.690	9.126	330
1.4%	3708	0	0	GET	/api/memories	513.925	0.139	0.879	2.118	11040
1.1%	631	0	0	PUT	^/api/follows/[^/]+$	424.021	0.672	6.011	7.955	0
1.1%	515	0	0	PATCH	/api/me	421.779	0.819	5.906	9.419	277
1.0%	1021	1	0	POST	/api/auth/register	378.032	0.370	2.800	5.154	227
0.7%	13789	0	0	GET	^/api/feeds/[^/]+	251.423	0.018	0.136	0.508	2224
0.7%	713	0	0	POST	/api/auth/logout	245.138	0.344	3.638	7.263	0
0.6%	263	0	0	GET	/api/me/export	241.791	0.919	4.439	4.612	1860433
0.5%	200	3	0	PATCH	^/api/posts/[^/]+$	173.789	0.869	4.082	5.108	637
0.3%	93	0	0	DELETE	^/api/posts/[^/]+$	116.965	1.258	4.303	4.303	0
0.2%	3534	0	0	GET	/api/posting-window/current	66.309	0.019	0.138	1.875	237
0.1%	105	0	0	DELETE	^/api/comments/[^/]+$	49.653	0.473	3.670	3.920	0
0.1%	60	0	0	DELETE	^/api/realmojis/[^/]+$	44.846	0.747	3.390	3.390	0
0.1%	2056	66	0	GET	^/api/posts/me/current$	44.589	0.022	0.138	0.392	628
0.1%	2686	0	0	GET	/api/users	43.614	0.016	0.099	0.347	238
0.1%	682	0	0	GET	^/api/memories/[^/]+	42.574	0.062	0.660	1.051	638
0.1%	1862	0	0	GET	^/api/relationships/[^/]+	36.496	0.020	0.132	0.472	967
0.1%	1982	26	0	GET	^/api/posts/[^/]+$	33.787	0.017	0.116	0.335	605
0.1%	1	0	0	POST	/api/initialize	19.668	19.668	19.668	19.668	90
0.0%	1061	0	0	GET	^/api/users/[^/]+$	18.283	0.017	0.105	0.324	195
0.0%	828	0	0	GET	^/api/posts/[^/]+/comments$	16.345	0.020	0.121	3.675	1990
0.0%	850	0	0	GET	^/api/posts/[^/]+/realmojis$	14.200	0.017	0.092	1.033	837
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.8%	479	0	0	GET	/api/notifications	3989.665	8.329	12.235	12.855	170
13.1%	618	5	0	POST	/api/posts	1311.666	2.122	7.177	8.423	641
8.3%	1099	16	0	POST	/api/friend-requests	835.429	0.760	4.525	7.298	114
7.3%	2271	24	0	GET	^/api/images/[^/]+	736.134	0.324	2.150	4.328	217204
5.0%	2171	0	0	GET	^/api/feeds/[^/]+	500.956	0.231	0.777	0.971	2125
4.7%	644	0	0	POST	/api/auth/login	469.029	0.728	5.474	9.369	245
3.0%	343	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	301.920	0.880	4.690	6.897	116
2.8%	705	0	0	GET	/api/memories	276.914	0.393	1.432	1.857	11116
2.7%	862	0	0	GET	/api/posting-window/current	273.214	0.317	0.788	1.535	239
1.5%	518	18	0	GET	^/api/posts/me/current$	153.318	0.296	0.784	0.980	622
1.4%	525	0	0	GET	/api/users	135.743	0.259	0.719	1.359	238
1.3%	204	0	0	POST	/api/auth/register	131.532	0.645	3.219	7.034	228
1.2%	88	13	0	POST	^/api/posts/[^/]+/realmojis$	124.395	1.414	6.722	6.722	403
1.1%	116	0	0	PUT	/api/me/profile-image	115.196	0.993	6.233	7.091	0
0.9%	108	8	0	POST	^/api/posts/[^/]+/comments$	95.022	0.880	5.636	7.126	311
0.8%	35	0	0	GET	/api/me/export	84.115	2.403	9.042	9.042	2152344
0.7%	145	0	0	PUT	^/api/follows/[^/]+$	68.440	0.472	2.546	2.755	0
0.6%	97	0	0	POST	/api/auth/logout	63.039	0.650	3.741	3.741	0
0.6%	284	0	0	GET	^/api/relationships/[^/]+	62.732	0.221	0.773	1.359	973
0.5%	275	25	0	GET	^/api/posts/[^/]+$	54.642	0.199	0.752	0.897	556
0.4%	87	0	0	PATCH	/api/me	44.369	0.510	4.219	4.219	273
0.4%	46	5	0	PATCH	^/api/posts/[^/]+$	43.951	0.955	5.456	5.456	574
0.3%	26	0	0	DELETE	^/api/posts/[^/]+$	32.231	1.240	3.525	3.525	0
0.3%	142	0	0	GET	^/api/users/[^/]+$	26.131	0.184	0.733	0.940	188
0.2%	83	0	0	GET	^/api/memories/[^/]+	21.934	0.264	1.209	1.209	839
0.2%	121	1	0	GET	/api/me	21.727	0.180	0.402	0.696	300
0.2%	107	0	0	GET	^/api/posts/[^/]+/comments$	18.438	0.172	0.419	0.453	1941
0.2%	100	0	0	GET	^/api/posts/[^/]+/realmojis$	17.585	0.176	0.751	0.793	786
0.1%	10	0	0	DELETE	^/api/realmojis/[^/]+$	10.653	1.065	3.037	3.037	0
0.1%	12	0	0	DELETE	^/api/comments/[^/]+$	6.357	0.530	2.303	2.303	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.9%	258	0	0	GET	/api/notifications	1908.488	7.397	9.327	9.327	175
11.0%	314	4	0	POST	/api/posts	568.431	1.810	6.495	7.302	644
10.8%	608	3	0	POST	/api/friend-requests	557.461	0.917	5.094	8.571	114
7.0%	1275	8	0	GET	^/api/images/[^/]+	363.209	0.285	1.513	2.761	175248
5.5%	363	0	0	POST	/api/auth/login	282.963	0.780	6.623	9.613	249
5.0%	203	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	255.859	1.260	6.274	7.149	116
4.1%	849	0	0	GET	^/api/feeds/[^/]+	210.752	0.248	0.785	1.213	2218
2.4%	361	0	0	GET	/api/memories	126.428	0.350	1.374	1.540	12919
2.3%	441	0	0	GET	/api/posting-window/current	120.531	0.273	0.731	3.808	240
2.0%	47	15	0	POST	^/api/posts/[^/]+/realmojis$	104.463	2.223	9.543	9.543	311
1.7%	65	0	0	PUT	/api/me/profile-image	87.919	1.353	8.667	8.667	0
1.5%	265	27	0	GET	^/api/posts/me/current$	75.572	0.285	0.786	1.010	585
1.3%	278	0	0	GET	/api/users	69.664	0.251	0.819	2.767	238
1.3%	20	0	0	GET	/api/me/export	64.597	3.230	7.832	7.832	3118344
0.9%	111	0	0	POST	/api/auth/register	48.731	0.439	2.938	3.966	229
0.9%	56	0	0	POST	^/api/posts/[^/]+/comments$	47.661	0.851	4.846	4.846	352
0.8%	64	0	0	PUT	^/api/follows/[^/]+$	40.166	0.628	4.766	4.766	0
0.8%	51	0	0	POST	/api/auth/logout	39.106	0.767	5.088	5.088	0
0.6%	55	0	0	PATCH	/api/me	30.218	0.549	2.851	2.851	286
0.5%	124	0	0	GET	^/api/relationships/[^/]+	28.130	0.227	0.688	0.746	1175
0.5%	11	0	0	DELETE	^/api/posts/[^/]+$	27.838	2.531	7.678	7.678	0
0.5%	107	7	0	GET	^/api/posts/[^/]+$	23.310	0.218	0.585	0.609	563
0.4%	21	4	0	PATCH	^/api/posts/[^/]+$	21.261	1.012	3.283	3.283	526
0.3%	69	0	0	GET	/api/me	17.494	0.254	1.220	1.220	301
0.3%	53	0	0	GET	^/api/memories/[^/]+	14.218	0.268	1.070	1.070	1032
0.2%	48	0	0	GET	^/api/users/[^/]+$	8.852	0.184	0.528	0.528	188
0.2%	35	0	0	GET	^/api/posts/[^/]+/comments$	8.193	0.234	0.658	0.658	2286
0.1%	40	0	0	GET	^/api/posts/[^/]+/realmojis$	7.628	0.191	0.539	0.539	844
0.1%	10	0	0	DELETE	^/api/comments/[^/]+$	3.343	0.334	1.247	1.247	0
0.1%	2	0	0	DELETE	^/api/realmojis/[^/]+$	2.870	1.435	2.665	2.665	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 79503 クエリ / 83 種類 / 合計 239.808s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	42.7%	318	102.306	321.72	1341.05	1239.8	126.2	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	15.4%	4577	36.813	8.04	187.02	573.7	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	14.0%	3065	33.578	10.96	1091.43	129.1	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	9.5%	9553	22.876	2.39	62.95	11.0	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	6.4%	1	15.249	15249.01	15249.01	83456.0	8.0	select $1,count(*) from users union all select $2,count(*) from friendships union all select $3,count(*) from follows union all select $4,count(*) from block_re…
6	3.0%	3286	7.262	2.21	38.66	24.5	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
7	2.8%	2	6.601	3300.62	6443.38	234317.0	2854.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
8	0.9%	1	2.239	2239.19	2239.19	31394.0	1200122.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
9	0.8%	1	1.817	1817.37	1817.37	24949.0	628446.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
10	0.7%	557	1.753	3.15	15.24	28.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
11	0.7%	7149	1.629	0.23	19.78	19.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
12	0.5%	776	1.225	1.58	11.15	20.5	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
13	0.4%	3008	0.938	0.31	19.58	20.2	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
14	0.4%	1	0.876	875.57	875.57	31620.0	3335.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.3%	1	0.810	810.14	810.14	13050.0	530330.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 10c2b7710b2125468051e39e45f099ee49bda00a
Type: cpu
Time: 2026-10-31 14:17:11 JST
Duration: 60.15s, Total samples = 38.39s (63.82%)
Showing nodes accounting for 27s, 70.33% of 38.39s total
Dropped 976 nodes (cum <= 0.19s)
Showing top 40 nodes out of 284
      flat  flat%   sum%        cum   cum%
     5.19s 13.52% 13.52%      5.19s 13.52%  crypto/internal/fips140/sha256.blockAVX2
     3.43s  8.93% 22.45%      3.43s  8.93%  internal/runtime/syscall/linux.Syscall6
     1.86s  4.85% 27.30%      5.79s 15.08%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.78s  4.64% 31.94%      1.78s  4.64%  runtime.memmove
     1.65s  4.30% 36.23%      2.05s  5.34%  runtime.findObject
     1.58s  4.12% 40.35%      1.58s  4.12%  runtime.memclrNoHeapPointers
     1.38s  3.59% 43.94%      5.59s 14.56%  runtime.scanObject
     1.36s  3.54% 47.49%      1.36s  3.54%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.25s  3.26% 50.74%      2.08s  5.42%  runtime.tryDeferToSpanScan
     1.23s  3.20% 53.95%      1.23s  3.20%  aeshashbody
     0.65s  1.69% 55.64%      0.65s  1.69%  runtime.spanClass.sizeclass (inline)
     0.43s  1.12% 56.76%      1.93s  5.03%  runtime.mapaccess1_faststr
     0.40s  1.04% 57.80%      0.59s  1.54%  example.com/isureal-company/isureal.postLess
     0.39s  1.02% 58.82%      0.39s  1.02%  memeqbody
     0.33s  0.86% 59.68%         5s 13.02%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.31s  0.81% 60.48%      0.31s  0.81%  runtime.nextFreeFast (inline)
     0.28s  0.73% 61.21%      0.28s  0.73%  sync/atomic.(*Int32).Add
     0.27s   0.7% 61.92%      0.63s  1.64%  database/sql.convertAssignRows
     0.27s   0.7% 62.62%      1.83s  4.77%  runtime.mapassign_faststr
     0.22s  0.57% 63.19%      5.30s 13.81%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.21s  0.55% 63.74%      0.22s  0.57%  runtime.(*mspan).writeHeapBitsSmall
     0.20s  0.52% 64.26%      0.42s  1.09%  compress/flate.(*compressor).deflate
     0.19s  0.49% 64.76%      1.01s  2.63%  runtime.mallocgcSmallScanNoHeader
     0.18s  0.47% 65.23%      0.35s  0.91%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.16s  0.42% 65.64%      0.25s  0.65%  encoding/json.appendString[go.shape.string]
     0.15s  0.39% 66.03%      5.69s 14.82%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.15s  0.39% 66.42%      0.41s  1.07%  runtime.greyobject
     0.14s  0.36% 66.79%      5.62s 14.64%  crypto/internal/fips140/sha256.(*Digest).Write
     0.14s  0.36% 67.15%      2.20s  5.73%  example.com/isureal-company/isureal.listMemories
     0.14s  0.36% 67.52%      0.21s  0.55%  runtime.(*mspan).markBitsForIndex (inline)
     0.13s  0.34% 67.86%      4.19s 10.91%  runtime.mallocgc
     0.13s  0.34% 68.19%      0.35s  0.91%  runtime.scanObjectsSmall
     0.13s  0.34% 68.53%      0.20s  0.52%  runtime.typePointers.next
     0.12s  0.31% 68.85%      0.78s  2.03%  runtime.mallocgcSmallNoscan
     0.11s  0.29% 69.13%      0.35s  0.91%  internal/stringslite.Index
     0.10s  0.26% 69.39%      0.21s  0.55%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.10s  0.26% 69.65%      1.23s  3.20%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.09s  0.23% 69.89%      6.07s 15.81%  example.com/isureal-company/isureal.computeImageSignature
     0.09s  0.23% 70.12%      0.21s  0.55%  internal/runtime/maps.(*Iter).Next
     0.08s  0.21% 70.33%      0.41s  1.07%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 58.0% / 100% | 40% | 0.5% | 7 | 136 MB | CPU が飽和 |
| isu2 (2 cpu) | 17.8% / 95% | 1% | 0.6% | 3 | 1835 MB |  |
| isu3 (2 cpu) | 14.7% / 73% | 0% | 0.6% | 4 | 1754 MB |  |
| isu4 (2 cpu) | 52.5% / 100% | 16% | 24.5% | 23 | 113 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 77% / journalctl 20% / nginx 12% |
| isu2 | isureal 33% / nginx 1% |
| isu3 | isureal 27% / nginx 1% |
| isu4 | postgres 40% / isureal 10% / kswapd0 1% |

## Diff from previous run

スコア: 123551 → 134100 (+10549)

`247167c..55cbf84`

```
 webapp/go/jpegcheck/huffman.go | 27 ++++++++++++++++++++++++---
 webapp/go/jpegcheck/reader.go  | 32 ++++++++++++++++++++------------
 webapp/go/jpegcheck/scan.go    |  7 ++++---
 webapp/go/post_api.go          | 16 ++++++++++++++++
 webapp/go/post_forward.go      |  5 +++++
 5 files changed, 69 insertions(+), 18 deletions(-)
```
