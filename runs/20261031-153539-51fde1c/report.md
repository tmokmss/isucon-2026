# run 20261031-153539-51fde1c

## Summary

| | |
|---|---|
| score | 184169 |
| status | pass |
| ref | 51fde1c |
| sha | 51fde1c8ae26 |
| 所要時間 | 213s |
| 前回比 | +838 (0.5%) vs 20261031-152137-8c13252 |
| ベスト比 | -7647 (-4.0%) vs 20261031-151339-5a230a4 = 191816 |
| note | nginx listen backlog=4096, tcp_max_syn_backlog=4096 (notifications 503) |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	151	443	1206	1207
europe	7	21	66	194	457	457
america	6	21	62	189	246	246
```

```
{"id":"23520","status":"done","portalStatus":"finished","passed":true,"score":184169,"raw":184169,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:37:25.096Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:37:28.089Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:37:31.089Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:37:34.096Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T06:37:36.776Z level=INFO msg="user moved" user=hiul5284 from=asia to=america
time=2026-10-31T06:37:37.090Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T06:37:40.089Z level=INFO msg="america 2 日目: 21 人"
time=2026-10-31T06:37:43.102Z level=INFO msg="asia 3 日目: 151 人"
time=2026-10-31T06:37:45.737Z level=INFO msg="user moved" user=rumar5161 from=asia to=europe
time=2026-10-31T06:37:46.089Z level=INFO msg="europe 3 日目: 66 人"
time=2026-10-31T06:37:48.195Z level=INFO msg="user moved" user=hael-6291 from=europe to=america
time=2026-10-31T06:37:48.872Z level=INFO msg="user moved" user=hael-6291 from=america to=asia
time=2026-10-31T06:37:49.090Z level=INFO msg="america 3 日目: 62 人"
time=2026-10-31T06:37:52.100Z level=INFO msg="asia 4 日目: 443 人"
time=2026-10-31T06:37:53.440Z level=INFO msg="user moved" user=hael-6291 from=asia to=america
time=2026-10-31T06:37:54.710Z level=INFO msg="user moved" user=orku_7212 from=asia to=europe
time=2026-10-31T06:37:55.528Z level=INFO msg="user moved" user=hayu_7120 from=asia to=america
time=2026-10-31T06:37:55.641Z level=INFO msg="europe 4 日目: 194 人"
time=2026-10-31T06:37:56.543Z level=INFO msg="user moved" user=hesolri-6947 from=asia to=america
time=2026-10-31T06:37:58.062Z level=INFO msg="user moved" user=tsuchi-8773 from=asia to=america
time=2026-10-31T06:37:58.212Z level=INFO msg="america 4 日目: 189 人"
time=2026-10-31T06:38:01.232Z level=INFO msg="asia 5 日目: 1206 人"
time=2026-10-31T06:38:03.512Z level=INFO msg="user moved" user=anho_9236 from=europe to=asia
time=2026-10-31T06:38:06.048Z level=INFO msg="europe 5 日目: 457 人"
time=2026-10-31T06:38:07.487Z level=INFO msg="user moved" user=ulzoewa_10311 from=america to=europe
time=2026-10-31T06:38:07.701Z level=INFO msg="america 5 日目: 246 人"
time=2026-10-31T06:38:17.534Z level=INFO msg="america 最終日（グランプリ）: 246 人"
time=2026-10-31T06:38:18.554Z level=INFO msg="europe 最終日（グランプリ）: 457 人"
time=2026-10-31T06:38:19.127Z level=INFO msg="asia 最終日（グランプリ）: 1207 人"
time=2026-10-31T06:38:25.168Z level=INFO msg="グランプリ: アジア 251 人, ヨーロッパ 172 人, アメリカ 150 人"
{
  "pass": true,
  "score": 184169,
  "messages": [
    "kuis.17268: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "benru_16655: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "elliaku-16188: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "martorro_8666: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "vinel17620: 投稿の公開範囲を変えられません: PATCH /api/posts/e39d1593-b33f-43b4-ae0e-830141272e32: 通信に失敗しました: Patch \"http://52.69.202.109:8080/api/posts/e39d1593-b33f-43b4-ae0e-830141272e32\": EOF",
    "yaivy.16768: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yulia16149: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ansu-16398: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maxshisol_16248: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marmaxse18175: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hobakbit_sujokgwan_3866: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "evaku17203: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "zoeso17893: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rikihi.17749: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kutorri_14009: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.39.248:35958-\u003e52.69.202.109:8080: write tcp 10.27.39.248:35958-\u003e52.69.202.109:8080: use of closed network connection",
    "solnoa-18241: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nifu-18283: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "zoedan-5251: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "memax5800: 友達申請を承認できません: POST /api/friend-requests/5d6d1c81-d9a4-4c71-a5dd-28cfa3562ee8/accept: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests/5d6d1c81-d9a4-4c71-a5dd-28cfa3562ee8/accept\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.8%	1905	0	0	GET	/api/notifications	19731.353	10.358	18.189	18.235	172
10.5%	6871	64	0	POST	/api/friend-requests	4638.219	0.675	2.601	4.322	114
10.3%	2982	255	0	POST	/api/posts	4555.677	1.528	4.917	5.842	598
5.1%	19788	383	0	GET	^/api/images/[^/]+	2266.497	0.115	1.428	3.606	127321
4.0%	6860	555	0	POST	/api/auth/login	1740.719	0.254	2.245	4.322	226
3.7%	3267	11	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1645.886	0.504	2.581	4.640	116
3.6%	4451	0	0	GET	/api/memories	1593.925	0.358	5.036	14.794	13129
3.0%	15627	0	0	GET	^/api/feeds/[^/]+	1333.277	0.085	0.638	1.869	3118
1.8%	818	133	0	PUT	/api/me/profile-image	794.310	0.971	7.893	10.815	3
1.7%	729	12	0	POST	^/api/posts/[^/]+/realmojis$	770.192	1.057	5.528	6.501	466
1.7%	1228	1	0	POST	/api/auth/register	738.480	0.601	2.508	2.785	227
1.1%	532	10	0	POST	^/api/posts/[^/]+/comments$	504.059	0.947	5.377	7.814	328
1.1%	781	0	0	PUT	^/api/follows/[^/]+$	473.026	0.606	2.325	3.538	0
0.8%	399	0	0	GET	/api/me/export	366.857	0.919	13.317	15.367	2060159
0.8%	4326	0	0	GET	/api/posting-window/current	348.219	0.080	0.567	1.845	237
0.7%	615	0	0	PATCH	/api/me	327.743	0.533	2.448	3.231	279
0.7%	3262	0	0	GET	/api/users	320.138	0.098	0.519	1.845	241
0.7%	1084	0	0	POST	/api/auth/logout	312.132	0.288	2.332	3.503	0
0.6%	2107	21	0	GET	^/api/posts/me/current$	242.302	0.115	1.272	1.845	647
0.4%	239	4	0	PATCH	^/api/posts/[^/]+$	170.496	0.713	2.617	4.067	640
0.4%	696	2	0	GET	/api/me	168.581	0.242	0.804	0.890	302
0.4%	2269	50	0	GET	^/api/posts/[^/]+$	166.635	0.073	0.517	1.441	619
0.3%	1835	0	0	GET	^/api/relationships/[^/]+	132.330	0.072	0.536	1.470	1129
0.3%	136	0	0	DELETE	^/api/realmojis/[^/]+$	126.217	0.928	2.849	3.743	0
0.3%	117	0	0	DELETE	^/api/posts/[^/]+$	115.900	0.991	2.887	3.033	0
0.2%	974	0	0	GET	^/api/memories/[^/]+	103.612	0.106	1.190	2.121	621
0.2%	1208	0	0	GET	^/api/users/[^/]+$	77.530	0.064	0.449	1.421	206
0.2%	105	0	0	DELETE	^/api/comments/[^/]+$	66.262	0.631	2.550	2.796	0
0.1%	903	2	0	GET	^/api/posts/[^/]+/realmojis$	63.999	0.071	0.562	0.987	1727
0.1%	715	1	0	GET	^/api/posts/[^/]+/comments$	49.031	0.069	0.518	0.716	1582
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.2%	640	2	0	GET	/api/notifications	5149.382	8.046	12.294	15.097	155
14.2%	2271	27	0	POST	/api/friend-requests	1856.440	0.817	3.036	4.265	114
8.7%	998	67	0	POST	/api/posts	1147.117	1.149	4.854	5.591	605
6.1%	1305	0	0	POST	/api/auth/login	798.445	0.612	2.758	6.248	245
4.5%	3986	0	0	GET	^/api/feeds/[^/]+	588.940	0.148	1.379	2.143	2622
4.4%	4142	34	0	GET	^/api/images/[^/]+	577.439	0.139	1.174	3.586	174979
2.8%	535	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	362.312	0.677	2.953	3.850	115
2.5%	455	0	0	POST	/api/auth/register	331.143	0.728	2.247	2.326	229
2.4%	265	0	0	PUT	/api/me/profile-image	309.405	1.168	3.159	3.881	0
2.1%	300	0	0	PUT	^/api/follows/[^/]+$	280.136	0.934	2.980	3.485	0
1.7%	208	0	0	PATCH	/api/me	226.908	1.091	3.241	4.144	284
1.7%	1058	0	0	GET	/api/memories	218.429	0.206	1.775	8.633	13162
1.2%	1508	0	0	GET	/api/posting-window/current	157.276	0.104	0.475	1.571	239
1.2%	152	7	0	POST	^/api/posts/[^/]+/realmojis$	154.846	1.019	3.713	4.906	449
1.1%	1122	0	0	GET	/api/users	144.714	0.129	0.912	1.865	242
0.9%	170	5	0	POST	^/api/posts/[^/]+/comments$	117.392	0.691	3.796	5.047	338
0.8%	740	39	0	GET	^/api/posts/me/current$	99.818	0.135	1.233	1.869	615
0.8%	199	0	0	POST	/api/auth/logout	98.674	0.496	2.723	2.747	0
0.7%	569	0	0	GET	^/api/relationships/[^/]+	87.399	0.154	1.391	2.401	745
0.6%	517	20	0	GET	^/api/posts/[^/]+$	75.822	0.147	1.622	1.867	597
0.4%	45	0	0	DELETE	^/api/posts/[^/]+$	57.426	1.276	4.282	4.282	0
0.4%	69	7	0	PATCH	^/api/posts/[^/]+$	54.783	0.794	2.322	2.322	584
0.4%	270	2	0	GET	/api/me	50.997	0.189	0.953	1.009	303
0.3%	223	0	0	GET	^/api/users/[^/]+$	33.048	0.148	1.053	1.849	198
0.2%	170	1	0	GET	^/api/posts/[^/]+/realmojis$	30.943	0.182	2.232	2.421	1851
0.2%	178	0	0	GET	^/api/posts/[^/]+/comments$	25.763	0.145	1.784	1.787	1804
0.2%	31	0	0	DELETE	^/api/realmojis/[^/]+$	24.495	0.790	3.114	3.114	0
0.1%	158	0	0	GET	^/api/memories/[^/]+	18.402	0.116	1.345	1.867	939
0.1%	67	0	0	GET	/api/me/export	15.375	0.229	5.064	5.064	2002918
0.1%	36	0	0	DELETE	^/api/comments/[^/]+$	12.411	0.345	1.363	1.363	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.8%	466	0	0	GET	/api/notifications	3611.491	7.750	10.210	10.263	167
11.5%	1317	16	0	POST	/api/friend-requests	996.899	0.757	2.665	3.534	114
9.3%	666	24	0	POST	/api/posts	804.285	1.208	4.717	6.751	629
6.1%	997	0	0	POST	/api/auth/login	530.889	0.532	2.975	5.048	250
5.0%	3390	25	0	GET	^/api/images/[^/]+	430.721	0.127	1.017	3.339	156481
4.1%	2677	0	0	GET	^/api/feeds/[^/]+	358.366	0.134	1.188	1.968	3127
3.2%	387	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	276.449	0.714	3.246	3.928	116
2.5%	910	0	0	GET	/api/memories	219.460	0.241	2.016	12.625	15098
2.0%	938	0	0	GET	/api/posting-window/current	176.802	0.188	1.682	2.268	240
1.5%	176	0	0	PUT	^/api/follows/[^/]+$	129.952	0.738	3.176	3.288	0
1.5%	241	0	0	POST	/api/auth/register	129.740	0.538	2.212	2.294	230
1.4%	112	8	0	POST	^/api/posts/[^/]+/realmojis$	116.725	1.042	4.545	6.107	429
1.3%	657	0	0	GET	/api/users	114.136	0.174	1.720	2.267	242
1.0%	115	0	0	PATCH	/api/me	89.975	0.782	3.480	4.524	271
1.0%	122	0	0	PUT	/api/me/profile-image	89.540	0.734	3.274	4.527	0
1.0%	96	4	0	POST	^/api/posts/[^/]+/comments$	83.046	0.865	4.342	4.342	337
0.9%	146	0	0	POST	/api/auth/logout	81.411	0.558	2.276	3.295	0
0.9%	573	23	0	GET	^/api/posts/me/current$	77.578	0.135	1.273	1.847	628
0.6%	364	0	0	GET	^/api/relationships/[^/]+	53.532	0.147	1.177	1.836	742
0.6%	58	4	0	PATCH	^/api/posts/[^/]+$	49.792	0.858	2.636	2.636	607
0.5%	411	14	0	GET	^/api/posts/[^/]+$	42.641	0.104	0.815	1.837	596
0.4%	34	0	0	DELETE	^/api/posts/[^/]+$	36.243	1.066	2.636	2.636	0
0.3%	170	0	0	GET	^/api/memories/[^/]+	27.800	0.164	2.183	2.483	1196
0.3%	125	0	0	GET	/api/me	21.834	0.175	1.874	1.947	304
0.2%	181	0	0	GET	^/api/users/[^/]+$	21.586	0.119	1.273	2.230	188
0.2%	49	0	0	GET	/api/me/export	16.317	0.333	4.707	4.707	2647895
0.2%	137	0	0	GET	^/api/posts/[^/]+/realmojis$	15.192	0.111	0.488	2.646	1443
0.2%	113	1	0	GET	^/api/posts/[^/]+/comments$	12.994	0.115	0.717	0.783	1086
0.1%	8	0	0	PUT	^/api/blocks/[^/]+$	8.522	1.065	2.510	2.510	0
0.1%	15	0	0	DELETE	^/api/realmojis/[^/]+$	8.410	0.561	1.595	1.595	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 17811 クエリ / 48 種類 / 合計 157.443s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	56.6%	409	89.039	217.70	1123.05	1307.1	133.7	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	13.9%	3	21.874	7291.28	9834.00	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.0%	13850	17.266	1.25	8.27	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.2%	3	16.062	5354.08	7232.69	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.2%	3	8.174	2724.68	3913.62	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	1.8%	4	2.793	698.19	2293.51	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.2%	4	0.364	90.88	164.44	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.2%	4	0.361	90.15	159.05	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.284	283.92	283.92	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.248	247.55	247.55	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.1%	1	0.232	232.37	232.37	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.154	154.47	154.47	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.096	95.68	95.68	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.081	80.55	80.55	2463.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	3335	0.074	0.02	0.44	0.9	30.0	SELECT ip, banned_until FROM ip_bans WHERE banned_until > clock_timestamp()

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
Build ID: 17335183b00efa8c6b15dbe38aa17fa67632c24e
Type: cpu
Time: 2026-10-31 15:37:05 JST
Duration: 60.36s, Total samples = 46.61s (77.22%)
Showing nodes accounting for 33.13s, 71.08% of 46.61s total
Dropped 1147 nodes (cum <= 0.23s)
Showing top 40 nodes out of 300
      flat  flat%   sum%        cum   cum%
     6.08s 13.04% 13.04%      6.08s 13.04%  crypto/internal/fips140/sha256.blockAVX2
     5.45s 11.69% 24.74%      5.45s 11.69%  internal/runtime/syscall/linux.Syscall6
     2.17s  4.66% 29.39%      2.17s  4.66%  runtime.memmove
     1.84s  3.95% 33.34%      1.84s  3.95%  runtime.memclrNoHeapPointers
     1.75s  3.75% 37.10%      2.30s  4.93%  runtime.findObject
     1.71s  3.67% 40.76%      6.76s 14.50%  runtime.scanObject
     1.58s  3.39% 44.15%      1.58s  3.39%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.53s  3.28% 47.44%      5.75s 12.34%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.46s  3.13% 50.57%      1.46s  3.13%  aeshashbody
     1.42s  3.05% 53.62%      2.47s  5.30%  runtime.tryDeferToSpanScan
     0.81s  1.74% 55.35%      0.81s  1.74%  runtime.spanClass.sizeclass (inline)
     0.47s  1.01% 56.36%      2.14s  4.59%  runtime.mapaccess1_faststr
     0.46s  0.99% 57.35%      6.89s 14.78%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.45s  0.97% 58.31%      0.45s  0.97%  runtime.nextFreeFast (inline)
     0.43s  0.92% 59.24%      0.43s  0.92%  memeqbody
     0.39s  0.84% 60.07%      0.39s  0.84%  runtime.futex
     0.37s  0.79% 60.87%      0.57s  1.22%  example.com/isureal-company/isureal.postLess
     0.29s  0.62% 61.49%      7.30s 15.66%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.28s   0.6% 62.09%      0.79s  1.69%  database/sql.convertAssignRows
     0.28s   0.6% 62.69%      0.30s  0.64%  encoding/json.appendString[go.shape.string]
     0.28s   0.6% 63.29%      0.29s  0.62%  runtime.(*mspan).writeHeapBitsSmall
     0.27s  0.58% 63.87%      0.62s  1.33%  runtime.greyobject
     0.26s  0.56% 64.43%      0.26s  0.56%  indexbody
     0.25s  0.54% 64.96%      0.30s  0.64%  runtime.typePointers.next
     0.25s  0.54% 65.50%      0.25s  0.54%  sync/atomic.(*Int32).Add
     0.24s  0.51% 66.02%      0.37s  0.79%  internal/runtime/maps.(*Iter).Next
     0.23s  0.49% 66.51%      6.63s 14.22%  crypto/internal/fips140/sha256.(*Digest).Write
     0.23s  0.49% 67.00%      0.64s  1.37%  runtime.mapaccess2_faststr
     0.23s  0.49% 67.50%      0.28s   0.6%  runtime.spanOf (inline)
     0.22s  0.47% 67.97%      1.05s  2.25%  runtime.mallocgcSmallScanNoHeader
     0.21s  0.45% 68.42%      0.54s  1.16%  compress/flate.(*compressor).deflate
     0.20s  0.43% 68.85%      4.51s  9.68%  runtime.mallocgc
     0.18s  0.39% 69.23%      2.17s  4.66%  runtime.mapassign_faststr
     0.17s  0.36% 69.60%      7.72s 16.56%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.15s  0.32% 69.92%      0.53s  1.14%  runtime.scanObjectsSmall
     0.12s  0.26% 70.18%      0.87s  1.87%  runtime.mallocgcSmallNoscan
     0.11s  0.24% 70.41%      8.06s 17.29%  example.com/isureal-company/isureal.computeImageSignature
     0.11s  0.24% 70.65%      1.19s  2.55%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.11s  0.24% 70.89%      0.81s  1.74%  runtime.scanSpan
     0.09s  0.19% 71.08%      0.25s  0.54%  crypto/internal/fips140/hmac.(*HMAC).Reset
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 41.5% / 100% | 28% | 0.6% | 6 | 161 MB | run queue が積もっている |
| isu2 (2 cpu) | 27.9% / 100% | 4% | 0.5% | 4 | 313 MB |  |
| isu3 (2 cpu) | 25.0% / 98% | 3% | 0.6% | 4 | 258 MB |  |
| isu4 (2 cpu) | 38.5% / 100% | 17% | 15.3% | 17 | 108 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 67% / nginx 13% |
| isu2 | isureal 52% / nginx 2% |
| isu3 | isureal 48% / nginx 1% |
| isu4 | isureal 30% / postgres 14% |

## Diff from previous run

スコア: 183331 → 184169 (+838)

`8c13252..51fde1c`

```
 conf/_all/etc/sysctl.d/99-isucon.conf          |  2 ++
 conf/isu1/etc/nginx/sites-enabled/isureal.conf |  2 +-
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |  2 +-
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |  2 +-
 webapp/go/image_held.go                        |  5 +---
 webapp/go/image_local.go                       |  1 -
 webapp/go/persist.go                           |  2 +-
 webapp/go/realmoji_api.go                      | 41 ++++++++++----------------
 webapp/go/replica_server.go                    |  8 +----
 webapp/go/store.go                             |  3 +-
 10 files changed, 25 insertions(+), 43 deletions(-)
```
