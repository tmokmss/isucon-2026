# run 20261031-135423-25c2db4

## Summary

| | |
|---|---|
| score | 82880 |
| status | pass |
| ref | image-files |
| sha | 25c2db44de63 |
| 所要時間 | 291s |
| 前回比 | +18180 (28.1%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +18180 (28.1%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | image files, no per-file fsync (sync every 1s) + main c6abd65 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	150	433	450	—
europe	7	18	53	149	283	283
america	6	15	46	107	—	—

最終日（1 人 20 点）が来なかったリージョン: asia, america
```

```
{"id":"15016","status":"done","portalStatus":"finished","passed":true,"score":82880,"raw":82880,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:57:35.206Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:57:38.179Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:57:41.178Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:57:44.199Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T04:57:45.433Z level=INFO msg="user moved" user=tekaben.5218 from=asia to=america
time=2026-10-31T04:57:46.766Z level=INFO msg="user moved" user=tekaben.5218 from=america to=asia
time=2026-10-31T04:57:47.179Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T04:57:50.180Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T04:57:53.203Z level=INFO msg="asia 3 日目: 150 人"
time=2026-10-31T04:57:53.836Z level=INFO msg="user moved" user=miase5580 from=asia to=america
time=2026-10-31T04:57:54.326Z level=INFO msg="user moved" user=tekaben.5218 from=asia to=america
time=2026-10-31T04:57:55.499Z level=INFO msg="user moved" user=wane_5909 from=asia to=europe
time=2026-10-31T04:57:55.927Z level=INFO msg="user moved" user=miis-5671 from=asia to=america
time=2026-10-31T04:57:56.178Z level=INFO msg="europe 3 日目: 53 人"
time=2026-10-31T04:57:56.488Z level=INFO msg="user moved" user=murare_6098 from=asia to=europe
time=2026-10-31T04:57:59.177Z level=INFO msg="america 3 日目: 46 人"
time=2026-10-31T04:58:02.199Z level=INFO msg="asia 4 日目: 433 人"
time=2026-10-31T04:58:02.821Z level=INFO msg="user moved" user=remami6699 from=asia to=europe
time=2026-10-31T04:58:05.510Z level=INFO msg="europe 4 日目: 149 人"
time=2026-10-31T04:58:06.141Z level=INFO msg="user moved" user=remami6699 from=europe to=america
time=2026-10-31T04:58:08.205Z level=INFO msg="america 4 日目: 107 人"
time=2026-10-31T04:58:08.443Z level=INFO msg="user moved" user=wane_5909 from=europe to=america
time=2026-10-31T04:58:08.962Z level=INFO msg="user moved" user=murare_6098 from=europe to=america
time=2026-10-31T04:58:10.677Z level=INFO msg="user moved" user=chiku6832 from=asia to=america
time=2026-10-31T04:58:12.774Z level=INFO msg="asia 5 日目: 450 人"
time=2026-10-31T04:58:14.274Z level=INFO msg="europe 5 日目: 283 人"
time=2026-10-31T04:58:16.434Z level=INFO msg="user moved" user=tekaben.5218 from=america to=asia
time=2026-10-31T04:58:22.213Z level=INFO msg="user moved" user=miis-5671 from=america to=asia
time=2026-10-31T04:58:25.304Z level=INFO msg="user moved" user=kileois-9523 from=america to=asia
time=2026-10-31T04:58:26.268Z level=INFO msg="user moved" user=miase5580 from=america to=europe
time=2026-10-31T04:58:28.967Z level=INFO msg="europe 最終日（グランプリ）: 283 人"
time=2026-10-31T04:58:35.210Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 168 人, アメリカ 0 人"
{
  "pass": true,
  "score": 82880,
  "messages": [
    "america の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました",
    "kileois-9523: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "orneri.7454: 友達申請を送れません: POST /api/friend-requests: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests\": EOF",
    "soel.8636: 友達申請を送れません: POST /api/friend-requests: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests\": EOF",
    "zoemia-7897: 投稿の公開範囲を変えられません: PATCH /api/posts/1e25049e-8db5-4b01-a578-30ad05a9528c: 通信に失敗しました: Patch \"http://52.69.202.109:8080/api/posts/1e25049e-8db5-4b01-a578-30ad05a9528c\": EOF",
    "ivymo7855: 友達申請を送れません: POST /api/friend-requests: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests\": EOF",
    "tazoena7180: 友達申請を送れません: POST /api/friend-requests: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests\": EOF",
    "fuka_8311: 投稿の公開範囲を変えられません: PATCH /api/posts/e06529da-4f03-40e6-ba30-cea4a0ed35ba: 通信に失敗しました: Patch \"http://52.69.202.109:8080/api/posts/e06529da-4f03-40e6-ba30-cea4a0ed35ba\": EOF",
    "noaivy9800: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hajun_jeong_976: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "tankao_mini_huoji_2265: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "zhumian_wushe_3701: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "yowa_8723: 友達申請を送れません: POST /api/friend-requests: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests\": EOF",
    "tsuki10155: 友達申請を送れません: POST /api/friend-requests: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests\": EOF",
    "mero_7343: 投稿の公開範囲を変えられません: PATCH /api/posts/146965ee-c60d-4c4f-8bc3-74e7827d9fcf: 通信に失敗しました: Patch \"http://52.69.202.109:8080/api/posts/146965ee-c60d-4c4f-8bc3-74e7827d9fcf\": EOF",
    "yomiasa6255: 投稿を削除できません: DELETE /api/posts/6fdf784e-2f78-4b04-a18e-6ef1c7a9f858: 通信に失敗しました: Delete \"http://52.69.202.109:8080/api/posts/6fdf784e-2f78-4b04-a18e-6ef1c7a9f858\": EOF",
    "suel_9987: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "subin_yun_3514: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "hitoyasumi_sento_694: ログインできません: POST /api/auth/login: 応答がタイムアウトしました",
    "hupo_shuizuguan_3040: ログインできません: POST /api/auth/login: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
56.1%	981	3	0	GET	/api/notifications	14376.205	14.655	23.735	24.238	158
11.3%	929	1	0	POST	/api/posts	2889.727	3.111	11.985	12.520	647
7.6%	8574	637	0	GET	^/api/images/[^/]+	1954.413	0.228	2.689	11.165	104610
4.3%	4806	0	0	GET	^/api/feeds/[^/]+	1103.452	0.230	2.446	4.810	2355
3.0%	3854	1237	0	POST	/api/auth/login	771.471	0.200	2.446	8.921	174
2.5%	2150	0	0	GET	/api/memories	651.932	0.303	2.272	4.810	11630
2.5%	2366	38	0	POST	/api/friend-requests	636.609	0.269	2.378	2.689	114
2.2%	342	94	0	PUT	/api/me/profile-image	563.542	1.648	16.483	22.052	6
2.1%	914	28	0	GET	^/api/posts/me/current$	530.189	0.580	4.810	4.810	629
2.0%	165	0	0	GET	/api/me/export	520.176	3.153	20.776	20.966	1915057
1.1%	937	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	279.795	0.299	2.230	3.135	116
1.0%	1141	0	0	GET	/api/users	249.661	0.219	2.378	3.688	243
0.7%	473	2	0	POST	/api/auth/register	181.433	0.384	2.134	2.323	226
0.5%	1420	0	0	GET	/api/posting-window/current	140.760	0.099	1.378	4.810	236
0.4%	256	2	0	GET	/api/me	105.924	0.414	2.378	4.810	298
0.4%	693	0	0	GET	^/api/relationships/[^/]+	97.529	0.141	1.378	2.378	1228
0.3%	280	0	0	PUT	^/api/follows/[^/]+$	78.495	0.280	2.378	2.378	0
0.3%	709	16	0	GET	^/api/posts/[^/]+$	68.714	0.097	1.268	2.689	597
0.3%	220	0	0	PATCH	/api/me	67.278	0.306	2.378	2.378	273
0.2%	440	0	0	POST	/api/auth/logout	53.983	0.123	2.131	4.810	0
0.2%	413	0	0	GET	^/api/memories/[^/]+	53.620	0.130	2.272	3.688	628
0.2%	151	1	0	POST	^/api/posts/[^/]+/realmojis$	52.517	0.348	2.461	5.032	483
0.2%	96	5	0	PATCH	^/api/posts/[^/]+$	47.377	0.494	2.230	2.230	618
0.2%	46	0	0	DELETE	^/api/posts/[^/]+$	43.856	0.953	3.909	3.909	0
0.1%	302	0	0	GET	^/api/users/[^/]+$	26.030	0.086	1.268	2.689	185
0.1%	203	1	0	POST	^/api/posts/[^/]+/comments$	24.044	0.118	2.254	2.689	325
0.1%	1	0	0	POST	/api/initialize	20.262	20.262	20.262	20.262	90
0.1%	263	1	0	GET	^/api/posts/[^/]+/realmojis$	17.958	0.068	2.689	2.689	741
0.0%	28	0	0	DELETE	^/api/realmojis/[^/]+$	11.280	0.403	2.689	2.689	0
0.0%	277	0	0	GET	^/api/posts/[^/]+/comments$	7.581	0.027	0.250	2.032	1369
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.5%	519	0	0	GET	/api/notifications	5063.467	9.756	15.720	15.912	178
19.5%	775	2	0	POST	/api/posts	2436.031	3.143	9.878	11.174	649
13.4%	2898	17	0	GET	^/api/images/[^/]+	1680.519	0.580	4.315	6.093	179279
5.0%	2381	0	0	GET	^/api/feeds/[^/]+	626.312	0.263	0.887	2.240	1819
2.7%	1585	17	0	POST	/api/friend-requests	334.530	0.211	0.842	3.697	114
2.5%	983	0	0	GET	/api/memories	308.869	0.314	1.134	2.469	7980
2.2%	1079	0	0	GET	/api/posting-window/current	272.660	0.253	0.912	1.656	239
2.2%	156	0	0	PUT	/api/me/profile-image	270.196	1.732	6.989	7.136	0
2.1%	874	0	0	POST	/api/auth/login	262.239	0.300	1.901	5.909	246
1.7%	47	0	0	GET	/api/me/export	210.737	4.484	18.584	18.584	2068273
1.5%	648	37	0	GET	^/api/posts/me/current$	182.645	0.282	0.911	2.594	613
1.1%	731	0	0	GET	/api/users	143.284	0.196	0.811	1.512	244
0.8%	281	0	0	POST	/api/auth/register	104.946	0.373	1.406	1.551	228
0.8%	94	9	0	POST	^/api/posts/[^/]+/realmojis$	98.756	1.051	6.739	6.739	446
0.7%	361	0	0	GET	^/api/relationships/[^/]+	88.315	0.245	0.818	1.570	695
0.6%	291	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	71.452	0.246	1.647	2.133	116
0.4%	251	11	0	GET	^/api/posts/[^/]+$	53.669	0.214	0.826	1.510	584
0.3%	115	6	0	POST	^/api/posts/[^/]+/comments$	43.165	0.375	3.882	10.017	316
0.3%	159	0	0	PUT	^/api/follows/[^/]+$	35.915	0.226	1.451	1.467	0
0.3%	122	0	0	PATCH	/api/me	31.902	0.261	3.801	4.579	278
0.2%	123	0	0	POST	/api/auth/logout	31.190	0.254	1.728	1.972	0
0.2%	54	11	0	PATCH	^/api/posts/[^/]+$	25.549	0.473	6.974	6.974	529
0.2%	163	2	0	GET	/api/me	24.531	0.150	0.694	1.445	299
0.2%	103	0	0	GET	^/api/memories/[^/]+	23.300	0.226	0.816	0.910	728
0.2%	109	0	0	GET	^/api/users/[^/]+$	22.134	0.203	0.828	1.878	191
0.2%	25	0	0	DELETE	^/api/posts/[^/]+$	19.775	0.791	1.883	1.883	0
0.1%	86	0	0	GET	^/api/posts/[^/]+/comments$	12.634	0.147	0.476	0.476	1112
0.1%	75	0	0	GET	^/api/posts/[^/]+/realmojis$	11.652	0.155	0.525	0.525	910
0.1%	22	0	0	DELETE	^/api/comments/[^/]+$	6.865	0.312	2.170	2.170	0
0.0%	10	0	0	DELETE	^/api/realmojis/[^/]+$	3.166	0.317	0.654	0.654	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
61.5%	196	0	0	GET	/api/notifications	3311.354	16.895	27.833	27.833	148
8.4%	166	0	0	POST	/api/posts	454.542	2.738	11.145	11.236	649
7.1%	2159	7	0	GET	^/api/images/[^/]+	379.951	0.176	1.017	2.997	119604
4.5%	1748	0	0	GET	^/api/feeds/[^/]+	241.685	0.138	0.592	1.094	2622
3.1%	43	0	0	GET	/api/me/export	167.616	3.898	12.907	12.907	2384909
2.3%	645	0	0	POST	/api/auth/login	123.392	0.191	1.661	4.186	252
2.2%	494	0	0	GET	/api/memories	118.718	0.240	1.167	1.976	19998
2.2%	602	8	0	POST	/api/friend-requests	117.032	0.194	1.178	2.153	114
1.1%	60	0	0	PUT	/api/me/profile-image	56.983	0.950	5.914	5.914	0
0.8%	294	0	0	GET	/api/posting-window/current	44.983	0.153	0.616	0.744	239
0.8%	298	0	0	GET	/api/users	44.252	0.148	0.625	0.750	241
0.7%	199	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	40.261	0.202	1.268	1.426	116
0.7%	299	4	0	GET	^/api/posts/[^/]+$	37.917	0.127	0.618	0.751	622
0.6%	255	0	0	GET	^/api/relationships/[^/]+	32.586	0.128	0.571	0.792	700
0.5%	22	4	0	POST	^/api/posts/[^/]+/realmojis$	29.196	1.327	5.586	5.586	374
0.4%	166	0	0	GET	^/api/posts/me/current$	22.986	0.138	0.354	0.381	649
0.4%	101	0	0	GET	^/api/memories/[^/]+	19.097	0.189	0.837	0.900	883
0.3%	39	3	0	POST	^/api/posts/[^/]+/comments$	18.603	0.477	4.393	4.393	329
0.3%	72	0	0	PUT	^/api/follows/[^/]+$	17.035	0.237	1.596	1.596	0
0.3%	110	0	0	POST	/api/auth/logout	17.007	0.155	0.850	1.038	0
0.3%	136	0	0	GET	^/api/users/[^/]+$	16.211	0.119	0.274	0.417	196
0.3%	57	0	0	PATCH	/api/me	16.088	0.282	6.063	6.063	273
0.3%	102	0	0	POST	/api/auth/register	14.490	0.142	0.309	0.454	229
0.2%	65	0	0	GET	/api/me	10.666	0.164	0.582	0.582	299
0.1%	44	0	0	GET	^/api/posts/[^/]+/realmojis$	7.858	0.179	0.926	0.926	496
0.1%	38	0	0	GET	^/api/posts/[^/]+/comments$	5.417	0.143	0.412	0.412	1264
0.1%	20	0	0	PATCH	^/api/posts/[^/]+$	3.383	0.169	0.321	0.321	640
0.0%	7	0	0	DELETE	^/api/posts/[^/]+$	2.632	0.376	0.479	0.479	0
0.0%	2	0	0	DELETE	/api/me/profile-image	2.285	1.143	2.061	2.061	0
0.0%	8	0	0	PUT	^/api/blocks/[^/]+$	2.187	0.273	0.889	0.889	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 46163 クエリ / 80 種類 / 合計 27.287s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	28.0%	4087	7.628	1.87	13.61	97.7	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $2) AS is_profile, EXISTS (SELECT $3 FROM posts p WHERE p.image_id…
2	19.0%	2	5.184	2592.23	5030.30	238697.0	2811.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
3	8.8%	1	2.392	2392.23	2392.23	31037.0	1200101.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
4	8.6%	1867	2.360	1.26	11.87	26.5	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
5	7.1%	1	1.945	1945.02	1945.02	24520.0	628062.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
6	4.0%	1	1.098	1098.03	1098.03	2598.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
7	3.7%	1	1.010	1010.15	1010.15	30423.0	2951.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
8	2.9%	1	0.805	804.53	804.53	12700.0	530237.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	2.2%	4477	0.611	0.14	6.70	20.2	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	2.2%	253	0.608	2.40	13.02	33.4	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
11	1.5%	347	0.398	1.15	8.59	22.9	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	1.4%	4136	0.380	0.09	4.82	11.4	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
13	1.2%	1	0.316	315.77	315.77	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	1.0%	2	0.283	141.35	274.51	16763.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	1.0%	1	0.265	265.02	265.02	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $2) AS is_profile, EXISTS (SELECT $3 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $4 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
-- rank 2
SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id;
-- rank 3
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: d86ffe3d7b351d6a68bdb22ab230ee4f4add448f
Type: cpu
Time: 2026-10-31 13:57:09 JST
Duration: 60.03s, Total samples = 29.84s (49.71%)
Showing nodes accounting for 21.71s, 72.75% of 29.84s total
Dropped 922 nodes (cum <= 0.15s)
Showing top 40 nodes out of 301
      flat  flat%   sum%        cum   cum%
     4.90s 16.42% 16.42%      4.90s 16.42%  internal/runtime/syscall/linux.Syscall6
     2.48s  8.31% 24.73%      2.48s  8.31%  crypto/internal/fips140/sha256.blockAVX2
     1.92s  6.43% 31.17%      5.98s 20.04%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.39s  4.66% 35.82%      1.39s  4.66%  aeshashbody
     1.38s  4.62% 40.45%      1.38s  4.62%  runtime.memclrNoHeapPointers
     1.26s  4.22% 44.67%      1.26s  4.22%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.87s  2.92% 47.59%      0.87s  2.92%  runtime.memmove
     0.79s  2.65% 50.23%      3.08s 10.32%  runtime.scanObject
     0.74s  2.48% 52.71%      1.06s  3.55%  runtime.findObject
     0.66s  2.21% 54.93%      1.12s  3.75%  runtime.tryDeferToSpanScan
     0.49s  1.64% 56.57%      2.03s  6.80%  runtime.mapaccess1_faststr
     0.34s  1.14% 57.71%      0.59s  1.98%  example.com/isureal-company/isureal.postLess
     0.34s  1.14% 58.85%      0.34s  1.14%  memeqbody
     0.31s  1.04% 59.89%      0.31s  1.04%  runtime.futex
     0.30s  1.01% 60.89%      0.44s  1.47%  database/sql.convertAssignRows
     0.29s  0.97% 61.86%      0.29s  0.97%  runtime.nextFreeFast (inline)
     0.29s  0.97% 62.84%      0.29s  0.97%  runtime.spanClass.sizeclass (inline)
     0.25s  0.84% 63.67%      1.76s  5.90%  runtime.mapassign_faststr
     0.22s  0.74% 64.41%      0.24s   0.8%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.21s   0.7% 65.11%      0.21s   0.7%  sync/atomic.(*Int32).Add (inline)
     0.17s  0.57% 65.68%      2.16s  7.24%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.17s  0.57% 66.25%      2.84s  9.52%  runtime.mallocgc
     0.17s  0.57% 66.82%      0.71s  2.38%  runtime.mallocgcSmallNoscan
     0.16s  0.54% 67.36%      1.34s  4.49%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.16s  0.54% 67.90%      0.16s  0.54%  indexbody
     0.16s  0.54% 68.43%      0.17s  0.57%  runtime.(*mspan).writeHeapBitsSmall
     0.16s  0.54% 68.97%      0.90s  3.02%  runtime.mallocgcSmallScanNoHeader
     0.15s   0.5% 69.47%      0.15s   0.5%  runtime.(*spanInlineMarkBits).init
     0.10s  0.34% 69.81%      2.16s  7.24%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.10s  0.34% 70.14%      0.50s  1.68%  encoding/json.structEncoder.encode
     0.09s   0.3% 70.44%      4.08s 13.67%  example.com/isureal-company/isureal.loadBig
     0.09s   0.3% 70.74%      0.17s  0.57%  internal/runtime/maps.(*Iter).Next
     0.08s  0.27% 71.01%      0.21s   0.7%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.08s  0.27% 71.28%      0.25s  0.84%  internal/stringslite.Index
     0.08s  0.27% 71.55%      0.45s  1.51%  runtime.mapaccess2_faststr
     0.08s  0.27% 71.82%      0.36s  1.21%  runtime.newobject
     0.07s  0.23% 72.05%      0.21s   0.7%  compress/flate.(*compressor).deflate
     0.07s  0.23% 72.29%      0.34s  1.14%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
     0.07s  0.23% 72.52%      0.19s  0.64%  runtime.greyobject
     0.07s  0.23% 72.75%      0.77s  2.58%  runtime.slicebytetostring
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 53.6% / 100% | 37% | 20.4% | 7 | 110 MB | CPU が飽和 |
| isu2 (2 cpu) | 17.4% / 100% | 2% | 0.5% | 5 | 1563 MB | run queue が積もっている |
| isu3 (2 cpu) | 9.1% / 65% | 0% | 0.6% | 5 | 1790 MB | run queue が積もっている |
| isu4 (2 cpu) | 18.2% / 86% | 0% | 5.1% | 13 | 110 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 45% / nginx 7% / kcompactd0 1% |
| isu2 | isureal 32% / nginx 1% |
| isu3 | isureal 16% / nginx 1% |
| isu4 | postgres 11% / isureal 3% |

## Diff from previous run

スコア: 64700 → 82880 (+18180)

`e0e4cb3..25c2db4`

```
 conf/_all/etc/sysctl.d/99-isucon.conf              |   2 +
 conf/isu1/etc/nginx/sites-enabled/isureal.conf     |   8 +
 conf/isu1/home/isucon/env.sh                       |   1 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf     |  19 ++
 conf/isu2/home/isucon/env.sh                       |   2 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf     |  19 ++
 conf/isu3/home/isucon/env.sh                       |   2 +
 .../etc/postgresql/16/main/conf.d/zz-tuning.conf   |   2 +
 webapp/go/account_deletion_api.go                  |   3 +-
 webapp/go/action_log.go                            |  13 +-
 webapp/go/action_log_test.go                       |  67 +++++
 webapp/go/admin_user_api.go                        |   2 +-
 webapp/go/comment_api.go                           |   2 +-
 webapp/go/config.go                                |   2 +
 webapp/go/export_api.go                            |  56 +---
 webapp/go/feed_api.go                              | 113 +++++--
 webapp/go/feed_test.go                             | 331 +++++++++++++++++++++
 webapp/go/image_api.go                             |  17 +-
 webapp/go/image_file.go                            | 145 +++++++++
 webapp/go/image_local.go                           |  43 ++-
 webapp/go/initialize.go                            |   1 +
 webapp/go/offload.go                               |  46 ++-
 webapp/go/post_api.go                              |  69 ++++-
 webapp/go/post_forward.go                          | 113 +++++++
 webapp/go/post_forward_test.go                     |  94 ++++++
 webapp/go/profile_image_api.go                     |  12 +-
 webapp/go/realmoji_api.go                          |   7 +-
 webapp/go/relationship.go                          |  10 +-
 webapp/go/server.go                                |   1 +
 webapp/go/store.go                                 | 166 +++++++++--
 webapp/go/user_api.go                              |   1 +
 webapp/go/user_read_api.go                         |  58 ++--
 webapp/go/user_search_test.go                      |  94 ++++++
 33 files changed, 1331 insertions(+), 190 deletions(-)
```
