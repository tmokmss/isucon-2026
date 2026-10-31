# run 20261031-120044-95362e7

## Summary

| | |
|---|---|
| score | 63045 |
| status | pass |
| ref | offload-jpeg |
| sha | 95362e7e2214 |
| 所要時間 | 164s |
| 前回比 | +6492 (11.5%) vs 20261031-115409-773f053 |
| ベスト比 | +2844 (4.7%) vs 20261031-115040-1c73ee4 = 60201 **ベスト更新** |
| note | jpeg-eu-us-sign-far-first |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	146	412	465	464
europe	7	14	33	79	81	82
america	7	13	33	53	52	51
```

```
{"id":"6372","status":"done","portalStatus":"finished","passed":true,"score":63045,"raw":63045,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:02:10.822Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:02:13.782Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:02:13.889Z level=INFO msg="user moved" user=yomu-5045 from=asia to=america
time=2026-10-31T03:02:16.782Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T03:02:19.812Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T03:02:22.782Z level=INFO msg="europe 2 日目: 14 人"
time=2026-10-31T03:02:25.782Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T03:02:26.292Z level=INFO msg="user moved" user=kemoma.5388 from=europe to=america
time=2026-10-31T03:02:28.804Z level=INFO msg="asia 3 日目: 146 人"
time=2026-10-31T03:02:31.241Z level=INFO msg="user moved" user=yomu-5045 from=america to=europe
time=2026-10-31T03:02:31.634Z level=INFO msg="user moved" user=evate_5701 from=asia to=america
time=2026-10-31T03:02:31.792Z level=INFO msg="europe 3 日目: 33 人"
time=2026-10-31T03:02:33.706Z level=INFO msg="user moved" user=evate_5701 from=america to=asia
time=2026-10-31T03:02:34.782Z level=INFO msg="america 3 日目: 33 人"
time=2026-10-31T03:02:37.300Z level=INFO msg="user moved" user=kemoma.5388 from=america to=asia
time=2026-10-31T03:02:37.447Z level=INFO msg="user moved" user=yomu-5045 from=europe to=america
time=2026-10-31T03:02:37.826Z level=INFO msg="asia 4 日目: 412 人"
time=2026-10-31T03:02:40.813Z level=INFO msg="europe 4 日目: 79 人"
time=2026-10-31T03:02:40.888Z level=INFO msg="user moved" user=kemoma.5388 from=asia to=europe
time=2026-10-31T03:02:43.943Z level=INFO msg="america 4 日目: 53 人"
time=2026-10-31T03:02:46.699Z level=INFO msg="user moved" user=benyore6821 from=asia to=europe
time=2026-10-31T03:02:46.983Z level=INFO msg="asia 5 日目: 465 人"
time=2026-10-31T03:02:49.851Z level=INFO msg="europe 5 日目: 81 人"
time=2026-10-31T03:02:51.160Z level=INFO msg="user moved" user=yomu-5045 from=america to=europe
time=2026-10-31T03:02:52.839Z level=INFO msg="america 5 日目: 52 人"
time=2026-10-31T03:03:01.953Z level=INFO msg="europe 最終日（グランプリ）: 82 人"
time=2026-10-31T03:03:02.250Z level=INFO msg="asia 最終日（グランプリ）: 464 人"
time=2026-10-31T03:03:02.252Z level=INFO msg="america 最終日（グランプリ）: 51 人"
time=2026-10-31T03:03:10.794Z level=INFO msg="グランプリ: アジア 77 人, ヨーロッパ 6 人, アメリカ 2 人"
{
  "pass": true,
  "score": 63045,
  "messages": [
    "miatsuhe-5030: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "senaso-7845: 認定公開アカウントをフォローできません: PUT /api/follows/e89a6267-97ae-4277-b44e-3ac17cf958b2: 応答がタイムアウトしました",
    "nayo8513: 認定公開アカウントをフォローできません: PUT /api/follows/e89a6267-97ae-4277-b44e-3ac17cf958b2: 応答がタイムアウトしました",
    "tani.8835: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "fuchira_7129: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "toeva-7223: 投稿 7bccfe64-8603-4640-9c07-028eecb0b981 に RealMoji を送れません: POST /api/posts/7bccfe64-8603-4640-9c07-028eecb0b981/realmojis: 応答がタイムアウトしました",
    "suyo-8943: 認定公開アカウントをフォローできません: PUT /api/follows/e89a6267-97ae-4277-b44e-3ac17cf958b2: 応答がタイムアウトしました",
    "shiannoa.9022: 認定公開アカウントをフォローできません: PUT /api/follows/05ff6308-230f-4812-a5e7-f0fe1ef745e8: 応答がタイムアウトしました",
    "moben-5979: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "rimiyo_9005: 認定公開アカウントをフォローできません: PUT /api/follows/05ff6308-230f-4812-a5e7-f0fe1ef745e8: 応答がタイムアウトしました",
    "zoena_6725: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "riku-7775: 投稿 7bccfe64-8603-4640-9c07-028eecb0b981 にコメントできません: POST /api/posts/7bccfe64-8603-4640-9c07-028eecb0b981/comments: 応答がタイムアウトしました",
    "vinke.7582: 投稿 7bccfe64-8603-4640-9c07-028eecb0b981 にコメントできません: POST /api/posts/7bccfe64-8603-4640-9c07-028eecb0b981/comments: 応答がタイムアウトしました",
    "kuruchi7485: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "shiri-5712: 投稿 7bccfe64-8603-4640-9c07-028eecb0b981 にコメントできません: POST /api/posts/7bccfe64-8603-4640-9c07-028eecb0b981/comments: 応答がタイムアウトしました",
    "heivymu_7593: 投稿 7bccfe64-8603-4640-9c07-028eecb0b981 に RealMoji を送れません: POST /api/posts/7bccfe64-8603-4640-9c07-028eecb0b981/realmojis: 応答がタイムアウトしました",
    "kaseis6011: 投稿 7bccfe64-8603-4640-9c07-028eecb0b981 にコメントできません: POST /api/posts/7bccfe64-8603-4640-9c07-028eecb0b981/comments: 応答がタイムアウトしました",
    "saheyo-7531: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "teyo6346: 友達申請を承認できません: POST /api/friend-requests/1fc5c730-bd74-4ff1-9b3a-ccd304422f09/accept: 応答がタイムアウトしました",
    "noaoran.9098: 認定公開アカウントをフォローできません: PUT /api/follows/c988aaf3-5f9e-4710-8e98-b056dec918e2: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.1%	959	2	0	GET	/api/notifications	8613.719	8.982	15.180	15.270	173
13.5%	1230	16	0	POST	/api/posts	3319.793	2.699	10.434	13.570	642
13.3%	2586	34	0	POST	/api/friend-requests	3271.016	1.265	12.407	18.671	114
8.6%	10825	718	0	GET	^/api/images/[^/]+	2119.130	0.196	2.728	10.846	96929
5.9%	4376	1070	0	POST	/api/auth/login	1434.950	0.328	4.359	15.229	192
4.4%	1029	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1088.186	1.058	10.053	15.264	116
4.1%	302	0	0	PUT	^/api/follows/[^/]+$	1001.628	3.317	17.597	18.128	0
3.6%	303	73	0	POST	^/api/posts/[^/]+/comments$	884.815	2.920	18.029	18.136	265
3.0%	262	56	0	POST	^/api/posts/[^/]+/realmojis$	746.764	2.850	17.947	18.260	393
1.5%	2772	0	0	GET	/api/memories	362.515	0.131	0.946	7.670	12144
1.3%	483	1	0	POST	/api/auth/register	326.571	0.676	3.022	8.651	227
1.1%	1724	1	0	GET	/api/posting-window/current	271.670	0.158	1.995	3.331	237
0.7%	5847	0	0	GET	^/api/feeds/[^/]+	177.101	0.030	0.285	1.841	1616
0.7%	380	117	0	PUT	/api/me/profile-image	174.227	0.458	2.644	3.752	7
0.7%	225	0	0	GET	/api/me/export	170.182	0.756	3.325	11.038	1847885
0.5%	574	0	0	POST	/api/auth/logout	117.455	0.205	3.326	9.743	0
0.2%	195	0	0	PATCH	/api/me	54.630	0.280	1.623	3.595	277
0.2%	106	4	0	PATCH	^/api/posts/[^/]+$	43.764	0.413	2.401	7.877	628
0.2%	968	17	0	GET	^/api/posts/me/current$	41.323	0.043	0.406	1.407	637
0.2%	62	0	0	DELETE	^/api/posts/[^/]+$	39.438	0.636	4.401	4.401	0
0.1%	1234	0	0	GET	/api/users	36.322	0.029	0.230	1.009	238
0.1%	43	0	0	DELETE	^/api/realmojis/[^/]+$	35.831	0.833	10.378	10.378	0
0.1%	509	0	0	GET	^/api/memories/[^/]+	31.246	0.061	0.411	1.995	579
0.1%	4	0	0	DELETE	^/api/follows/[^/]+$	22.906	5.726	13.933	13.933	0
0.1%	15	0	0	PUT	^/api/blocks/[^/]+$	22.290	1.486	14.500	14.500	0
0.1%	932	0	0	GET	^/api/relationships/[^/]+	21.323	0.023	0.231	0.499	1004
0.1%	878	11	0	GET	^/api/posts/[^/]+$	18.617	0.021	0.169	0.339	610
0.1%	15	0	0	PUT	^/api/admin/users/[^/]+/certification$	16.469	1.098	11.244	11.244	0
0.1%	1	0	0	POST	/api/initialize	12.988	12.988	12.988	12.988	90
0.0%	469	0	0	GET	^/api/users/[^/]+$	9.960	0.021	0.237	0.269	198
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.9%	180	0	0	GET	/api/notifications	1462.519	8.125	12.087	12.087	178
16.3%	218	4	0	POST	/api/posts	703.708	3.228	10.623	11.282	640
12.2%	402	11	0	POST	/api/friend-requests	528.539	1.315	11.525	15.312	113
5.6%	62	0	0	PUT	^/api/follows/[^/]+$	241.948	3.902	18.640	18.640	0
5.4%	394	0	0	POST	/api/auth/login	233.737	0.593	8.445	9.736	250
4.5%	1234	4	0	GET	^/api/images/[^/]+	196.237	0.159	1.021	1.420	146639
4.1%	195	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	178.368	0.915	8.561	10.044	116
3.3%	876	0	0	GET	^/api/feeds/[^/]+	143.083	0.163	0.543	1.871	1816
2.1%	304	0	0	GET	/api/posting-window/current	91.160	0.300	2.101	3.497	239
1.9%	31	7	0	POST	^/api/posts/[^/]+/comments$	80.965	2.612	17.756	17.756	238
1.8%	368	0	0	GET	/api/memories	79.605	0.216	1.091	1.961	12290
1.4%	29	0	0	GET	/api/me/export	58.423	2.015	9.490	9.490	2154518
1.2%	18	6	0	POST	^/api/posts/[^/]+/realmojis$	51.724	2.874	14.361	14.361	353
1.0%	55	0	0	POST	/api/auth/logout	42.951	0.781	10.420	10.420	0
0.7%	206	0	0	GET	/api/users	30.772	0.149	0.369	2.101	237
0.6%	78	0	0	POST	/api/auth/register	24.989	0.320	1.295	1.295	228
0.6%	158	6	0	GET	^/api/posts/me/current$	24.848	0.157	0.477	1.134	629
0.5%	154	0	0	GET	^/api/relationships/[^/]+	21.810	0.142	0.352	0.627	1277
0.4%	121	5	0	GET	^/api/posts/[^/]+$	19.199	0.159	0.503	0.514	595
0.4%	43	0	0	PUT	/api/me/profile-image	17.524	0.408	0.905	0.905	0
0.3%	66	0	0	GET	^/api/memories/[^/]+	12.074	0.183	1.042	1.042	820
0.3%	36	0	0	PATCH	/api/me	10.844	0.301	1.085	1.085	284
0.2%	5	0	0	DELETE	^/api/posts/[^/]+$	9.743	1.949	6.711	6.711	0
0.2%	61	0	0	GET	^/api/users/[^/]+$	8.962	0.147	0.409	0.409	185
0.2%	60	0	0	GET	^/api/posts/[^/]+/comments$	8.698	0.145	0.419	0.419	644
0.2%	51	0	0	GET	^/api/posts/[^/]+/realmojis$	7.184	0.141	0.378	0.378	323
0.2%	2	0	0	DELETE	^/api/follows/[^/]+$	7.170	3.585	4.194	4.194	0
0.2%	48	1	0	GET	/api/me	6.691	0.139	0.258	0.258	291
0.2%	1	0	0	DELETE	/api/me/profile-image	6.605	6.605	6.605	6.605	0
0.1%	15	0	0	PATCH	^/api/posts/[^/]+$	6.115	0.408	2.180	2.180	666
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.1%	141	0	0	GET	/api/notifications	1082.740	7.679	9.442	9.448	174
14.7%	306	2	0	POST	/api/friend-requests	440.674	1.440	15.013	17.211	114
11.3%	134	1	0	POST	/api/posts	338.161	2.524	9.936	10.300	662
6.3%	110	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	188.249	1.711	14.359	18.063	116
5.6%	282	0	0	POST	/api/auth/login	168.220	0.597	7.652	11.190	251
4.9%	969	7	0	GET	^/api/images/[^/]+	146.794	0.151	0.947	1.545	128992
3.5%	44	8	0	POST	^/api/posts/[^/]+/realmojis$	105.947	2.408	17.811	17.811	390
2.6%	498	0	0	GET	^/api/feeds/[^/]+	78.525	0.158	0.512	1.124	2179
2.3%	27	0	0	PUT	^/api/follows/[^/]+$	69.843	2.587	14.166	14.166	0
2.3%	260	0	0	GET	/api/memories	68.643	0.264	1.345	9.513	15988
1.7%	189	0	0	GET	/api/posting-window/current	49.759	0.263	2.351	3.154	239
1.1%	18	5	0	POST	^/api/posts/[^/]+/comments$	31.732	1.763	15.890	15.890	205
1.0%	33	0	0	POST	/api/auth/logout	29.743	0.901	7.787	7.787	0
0.8%	15	0	0	GET	/api/me/export	25.366	1.691	4.163	4.163	4307218
0.7%	32	0	0	PATCH	/api/me	21.107	0.660	2.805	2.805	276
0.7%	143	0	0	GET	/api/users	20.457	0.143	0.390	0.504	236
0.6%	51	0	0	POST	/api/auth/register	17.234	0.338	2.459	2.459	229
0.5%	109	5	0	GET	^/api/posts/me/current$	16.307	0.150	0.350	0.355	636
0.5%	30	0	0	PUT	/api/me/profile-image	15.041	0.501	2.000	2.000	0
0.5%	2	0	0	DELETE	^/api/follows/[^/]+$	13.978	6.989	13.869	13.869	0
0.5%	86	0	0	GET	^/api/relationships/[^/]+	13.718	0.160	0.571	0.571	1518
0.4%	7	0	0	DELETE	^/api/realmojis/[^/]+$	11.790	1.684	7.725	7.725	0
0.3%	51	0	0	GET	^/api/memories/[^/]+	9.702	0.190	0.647	0.647	941
0.3%	52	2	0	GET	^/api/posts/[^/]+$	7.901	0.152	0.422	0.422	592
0.2%	44	0	0	GET	^/api/users/[^/]+$	6.556	0.149	0.384	0.384	210
0.2%	34	1	0	GET	/api/me	5.255	0.155	0.386	0.386	289
0.1%	6	0	0	PATCH	^/api/posts/[^/]+$	4.408	0.735	1.448	1.448	669
0.1%	33	0	0	GET	^/api/posts/[^/]+/realmojis$	4.278	0.130	0.299	0.299	682
0.1%	17	1	0	GET	^/api/posts/[^/]+/comments$	2.449	0.144	0.243	0.243	899
0.1%	7	0	0	DELETE	^/api/posts/[^/]+$	2.363	0.338	0.476	0.476	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 47426 クエリ / 82 種類 / 合計 274.889s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	49.8%	2153	137.018	63.64	483.13	661.3	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	32.8%	269	90.250	335.50	1233.23	6893.1	129.0	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	6.7%	7935	18.505	2.33	49.28	10.1	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	5.1%	1883	14.000	7.43	262.93	114.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	1.3%	1563	3.633	2.32	15.01	26.4	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.9%	1	2.369	2369.25	2369.25	29657.0	1199555.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.6%	1	1.751	1751.32	1751.32	20524.0	626432.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.4%	1	1.105	1104.90	1104.90	11528.0	529837.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	0.4%	254	0.978	3.85	16.11	31.5	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.3%	3222	0.771	0.24	20.67	20.3	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.2%	267	0.544	2.04	14.79	24.4	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.2%	3982	0.499	0.13	12.17	11.3	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
13	0.2%	2	0.477	238.59	249.27	223783.5	2634.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.1%	1318	0.366	0.28	6.24	21.7	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	2	0.264	131.77	254.93	16757.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 4a2bbac5e3e93de5a5e666695cd437e080564ff9
Type: cpu
Time: 2026-10-31 12:01:51 JST
Duration: 60.15s, Total samples = 49.37s (82.08%)
Showing nodes accounting for 37.52s, 76.00% of 49.37s total
Dropped 1022 nodes (cum <= 0.25s)
Showing top 40 nodes out of 285
      flat  flat%   sum%        cum   cum%
     3.28s  6.64%  6.64%      3.28s  6.64%  internal/runtime/syscall/linux.Syscall6
     2.98s  6.04% 12.68%        10s 20.26%  runtime.scanObject
     2.91s  5.89% 18.57%      2.91s  5.89%  runtime.memclrNoHeapPointers
     2.38s  4.82% 23.39%      3.11s  6.30%  runtime.findObject
     2.29s  4.64% 28.03%      7.52s 15.23%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     2.24s  4.54% 32.57%      3.55s  7.19%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     2.22s  4.50% 37.07%      2.22s  4.50%  runtime.memmove
     2.04s  4.13% 41.20%      2.04s  4.13%  crypto/internal/fips140/sha256.blockAVX2
     1.91s  3.87% 45.07%      3.30s  6.68%  runtime.tryDeferToSpanScan
     1.83s  3.71% 48.77%      5.62s 11.38%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.62s  3.28% 52.06%      1.67s  3.38%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     1.30s  2.63% 54.69%      1.30s  2.63%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.23s  2.49% 57.18%      1.23s  2.49%  aeshashbody
     1.05s  2.13% 59.31%      1.35s  2.73%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.93s  1.88% 61.19%      0.93s  1.88%  runtime.spanClass.sizeclass (inline)
     0.88s  1.78% 62.97%      0.88s  1.78%  example.com/isureal-company/isureal.(*UserRow).active
     0.63s  1.28% 64.25%      0.63s  1.28%  memeqbody
     0.48s  0.97% 65.22%      1.99s  4.03%  runtime.mapaccess1_faststr
     0.39s  0.79% 66.01%      0.77s  1.56%  runtime.greyobject
     0.37s  0.75% 66.76%      0.65s  1.32%  example.com/isureal-company/isureal.postLess
     0.34s  0.69% 67.45%      0.34s  0.69%  runtime.typePointers.nextFast (inline)
     0.33s  0.67% 68.12%      0.33s  0.67%  runtime.nextFreeFast (inline)
     0.31s  0.63% 68.75%      0.31s  0.63%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.31s  0.63% 69.37%      0.31s  0.63%  runtime.madvise
     0.29s  0.59% 69.96%      0.29s  0.59%  indexbody
     0.29s  0.59% 70.55%      0.29s  0.59%  runtime.(*mspan).divideByElemSize (inline)
     0.26s  0.53% 71.08%      0.30s  0.61%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.25s  0.51% 71.58%      0.33s  0.67%  runtime.typePointers.next
     0.24s  0.49% 72.07%      1.56s  3.16%  runtime.mapassign_faststr
     0.22s  0.45% 72.51%      0.72s  1.46%  database/sql.convertAssignRows
     0.21s  0.43% 72.94%      1.06s  2.15%  runtime.mallocgcSmallScanNoHeader
     0.20s  0.41% 73.34%      0.25s  0.51%  runtime.spanOf (inline)
     0.19s  0.38% 73.73%      0.29s  0.59%  runtime.heapArenaOf (inline)
     0.18s  0.36% 74.09%      0.25s  0.51%  runtime.(*mspan).markBitsForIndex (inline)
     0.18s  0.36% 74.46%      0.54s  1.09%  runtime.mapaccess2_faststr
     0.16s  0.32% 74.78%      0.30s  0.61%  runtime.scanObjectSmall
     0.16s  0.32% 75.11%         1s  2.03%  runtime.scanSpan
     0.15s   0.3% 75.41%      6.30s 12.76%  runtime.mallocgc
     0.15s   0.3% 75.71%      0.54s  1.09%  runtime.scanObjectsSmall
     0.14s  0.28% 76.00%      0.34s  0.69%  compress/flate.(*compressor).deflate
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 55.4% / 100% | 38% | 0.1% | 8 | 222 MB | CPU が飽和 |
| isu2 (2 cpu) | 11.9% / 33% | 0% | 0.5% | 3 | 375 MB |  |
| isu3 (2 cpu) | 9.0% / 34% | 0% | 0.5% | 3 | 259 MB |  |
| isu4 (2 cpu) | 49.8% / 100% | 22% | 15.7% | 30 | 103 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 90% / nginx 13% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 22% |
| isu3 | isureal 15% / nginx 1% |
| isu4 | postgres 50% / isureal 13% |

## Diff from previous run

スコア: 56553 → 63045 (+6492)

`773f053..95362e7`

```
 conf/isu1/home/isucon/env.sh   |  3 +-
 webapp/go/limiter.go           | 37 ------------------------
 webapp/go/offload.go           | 65 ++++++++++++++++++++++++++++++++++++------
 webapp/go/post_api.go          |  4 +--
 webapp/go/profile_image_api.go |  3 +-
 webapp/go/realmoji_api.go      |  4 +--
 webapp/go/server.go            | 10 +++----
 7 files changed, 66 insertions(+), 60 deletions(-)
```
