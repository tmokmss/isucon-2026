# run 20261031-140017-04990e1

## Summary

| | |
|---|---|
| score | 134480 |
| status | pass |
| ref | image-files |
| sha | 04990e175c3e |
| 所要時間 | 492s |
| 前回比 | +51600 (62.3%) vs 20261031-135423-25c2db4 |
| ベスト比 | +15163 (12.7%) vs 20261031-135233-2d10aad = 119317 **ベスト更新** |
| note | hybrid: new images in asia1 files, initial from DB + main c6abd65 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	145	424	978	977
europe	7	18	53	146	170	170
america	6	16	49	111	112	112
```

```
{"id":"15747","status":"done","portalStatus":"finished","passed":true,"score":134480,"raw":134480,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:06:36.285Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:06:39.256Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:06:42.250Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:06:45.269Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T05:06:45.632Z level=INFO msg="user moved" user=noeva-5109 from=europe to=asia
time=2026-10-31T05:06:48.255Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T05:06:51.252Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T05:06:54.280Z level=INFO msg="asia 3 日目: 145 人"
time=2026-10-31T05:06:57.193Z level=INFO msg="user moved" user=noeva-5109 from=asia to=america
time=2026-10-31T05:06:57.250Z level=INFO msg="europe 3 日目: 53 人"
time=2026-10-31T05:07:00.436Z level=INFO msg="america 3 日目: 49 人"
time=2026-10-31T05:07:00.486Z level=INFO msg="user moved" user=kane_6158 from=europe to=asia
time=2026-10-31T05:07:03.284Z level=INFO msg="asia 4 日目: 424 人"
time=2026-10-31T05:07:06.502Z level=INFO msg="europe 4 日目: 146 人"
time=2026-10-31T05:07:07.035Z level=INFO msg="user moved" user=kane_6158 from=asia to=america
time=2026-10-31T05:07:09.203Z level=INFO msg="user moved" user=ruismi-8364 from=asia to=america
time=2026-10-31T05:07:09.286Z level=INFO msg="america 4 日目: 111 人"
time=2026-10-31T05:07:11.555Z level=INFO msg="user moved" user=maleoko-7854 from=asia to=america
time=2026-10-31T05:07:12.293Z level=INFO msg="user moved" user=nufuya7977 from=asia to=europe
time=2026-10-31T05:07:12.381Z level=INFO msg="asia 5 日目: 978 人"
time=2026-10-31T05:07:15.429Z level=INFO msg="europe 5 日目: 170 人"
time=2026-10-31T05:07:18.294Z level=INFO msg="america 5 日目: 112 人"
time=2026-10-31T05:07:28.769Z level=INFO msg="asia 最終日（グランプリ）: 977 人"
time=2026-10-31T05:07:29.197Z level=INFO msg="europe 最終日（グランプリ）: 170 人"
time=2026-10-31T05:07:29.200Z level=INFO msg="america 最終日（グランプリ）: 112 人"
time=2026-10-31T05:07:36.281Z level=INFO msg="グランプリ: アジア 442 人, ヨーロッパ 67 人, アメリカ 41 人"
{
  "pass": true,
  "score": 134480,
  "messages": [
    "rimuno.5020: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "muko_13198: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maxmuki_10314: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "zoemi.8107: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chivintor_8616: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marro-5773: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "anteki-7656: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ruvin13325: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "niki_6909: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kosote_10832: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "himax.5291: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "anor_10163: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yomufu.5647: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rean_10579: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "senaki.12127: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marnoa_6224: 友達申請を承認できません: POST /api/friend-requests/09cb568a-9c6a-4c8f-8b1b-f4cdbb722cbf/accept: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests/09cb568a-9c6a-4c8f-8b1b-f4cdbb722cbf/accept\": EOF",
    "subin_yun_2244: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "hanako_sato_2460: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "ulleo7936: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
66.3%	2085	1	0	GET	/api/notifications	41884.926	20.089	60.106	60.139	132
9.9%	2547	86	0	POST	/api/posts	6230.104	2.446	9.386	11.315	632
4.5%	16717	444	0	GET	^/api/images/[^/]+	2869.318	0.172	2.645	5.213	137119
4.3%	5910	87	0	POST	/api/friend-requests	2738.686	0.463	3.114	5.275	114
2.4%	5719	574	0	POST	/api/auth/login	1495.895	0.262	2.827	6.193	221
2.1%	2553	10	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1301.466	0.510	3.223	4.544	116
2.0%	699	127	0	PUT	/api/me/profile-image	1251.810	1.791	6.401	8.548	3
1.5%	596	22	0	POST	^/api/posts/[^/]+/realmojis$	921.101	1.545	8.079	9.519	460
1.0%	14568	0	0	GET	^/api/feeds/[^/]+	642.086	0.044	0.493	2.235	2820
1.0%	4054	0	0	GET	/api/memories	625.392	0.154	1.074	2.191	11461
1.0%	997	1	0	POST	/api/auth/register	614.579	0.616	3.103	5.142	227
0.7%	880	6	0	POST	^/api/posts/[^/]+/comments$	425.730	0.484	3.500	5.691	347
0.4%	660	0	0	PUT	^/api/follows/[^/]+$	281.404	0.426	2.848	4.696	0
0.4%	302	0	0	GET	/api/me/export	243.407	0.806	4.039	4.219	1907721
0.4%	473	0	0	PATCH	/api/me	224.677	0.475	3.569	5.062	280
0.3%	848	0	0	POST	/api/auth/logout	193.356	0.228	2.725	3.758	0
0.3%	3619	0	0	GET	/api/posting-window/current	174.554	0.048	0.449	1.132	237
0.2%	227	3	0	PATCH	^/api/posts/[^/]+$	124.878	0.550	3.093	4.169	640
0.2%	102	0	0	DELETE	^/api/posts/[^/]+$	123.284	1.209	3.190	3.633	0
0.2%	2856	0	0	GET	/api/users	97.751	0.034	0.259	1.257	238
0.1%	2094	34	0	GET	^/api/posts/[^/]+$	89.669	0.043	0.493	1.339	622
0.1%	2011	15	0	GET	^/api/posts/me/current$	88.008	0.044	0.555	2.763	648
0.1%	1854	0	0	GET	^/api/relationships/[^/]+	79.679	0.043	0.555	1.292	971
0.1%	178	0	0	DELETE	^/api/comments/[^/]+$	77.643	0.436	2.888	3.529	0
0.1%	872	0	0	GET	^/api/memories/[^/]+	65.474	0.075	0.646	0.776	640
0.1%	87	0	0	DELETE	^/api/realmojis/[^/]+$	58.220	0.669	2.723	2.723	0
0.1%	1054	0	0	GET	^/api/users/[^/]+$	49.045	0.047	0.505	1.593	215
0.0%	846	0	0	GET	^/api/posts/[^/]+/realmojis$	27.909	0.033	0.312	0.555	892
0.0%	890	0	0	GET	^/api/posts/[^/]+/comments$	27.794	0.031	0.261	1.095	2241
0.0%	44	0	0	PUT	^/api/blocks/[^/]+$	24.764	0.563	2.553	2.553	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.5%	386	0	0	GET	/api/notifications	3239.579	8.393	12.463	12.543	172
13.4%	2102	9	0	GET	^/api/images/[^/]+	1101.293	0.524	3.052	3.768	166277
10.8%	500	15	0	POST	/api/posts	883.443	1.767	5.343	7.431	624
6.6%	965	9	0	POST	/api/friend-requests	541.644	0.561	3.269	4.087	114
5.0%	1511	0	0	GET	^/api/feeds/[^/]+	409.944	0.271	1.074	1.195	2579
4.7%	619	0	0	POST	/api/auth/login	386.472	0.624	3.359	4.096	247
3.3%	688	0	0	GET	/api/posting-window/current	268.422	0.390	1.079	1.124	239
2.8%	396	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	230.418	0.582	3.071	3.736	116
2.7%	595	0	0	GET	/api/memories	219.132	0.368	1.459	1.970	12758
1.6%	409	23	0	GET	^/api/posts/me/current$	127.768	0.312	0.864	1.261	605
1.3%	462	0	0	GET	/api/users	110.684	0.240	0.813	1.020	232
1.2%	169	0	0	POST	/api/auth/register	99.199	0.587	3.559	3.774	228
1.1%	85	0	0	PUT	/api/me/profile-image	86.936	1.023	4.454	4.454	0
0.7%	89	0	0	POST	/api/auth/logout	56.727	0.637	3.554	3.554	0
0.7%	36	0	0	GET	/api/me/export	56.203	1.561	5.754	5.754	1569506
0.7%	94	0	0	PUT	^/api/follows/[^/]+$	53.905	0.573	3.415	3.415	0
0.6%	49	5	0	POST	^/api/posts/[^/]+/realmojis$	52.571	1.073	4.096	4.096	419
0.6%	220	0	0	GET	^/api/relationships/[^/]+	49.935	0.227	0.725	1.031	1518
0.5%	73	0	0	PATCH	/api/me	37.681	0.516	3.500	3.500	276
0.4%	162	9	0	GET	^/api/posts/[^/]+$	34.358	0.212	0.665	0.687	591
0.4%	64	0	0	POST	^/api/posts/[^/]+/comments$	29.436	0.460	3.373	3.373	326
0.3%	78	0	0	GET	^/api/memories/[^/]+	24.605	0.315	1.308	1.308	971
0.2%	89	1	0	GET	/api/me	19.546	0.220	0.873	0.873	299
0.2%	32	6	0	PATCH	^/api/posts/[^/]+$	16.591	0.518	1.590	1.590	515
0.2%	81	0	0	GET	^/api/users/[^/]+$	15.324	0.189	0.604	0.604	190
0.2%	84	0	0	GET	^/api/posts/[^/]+/realmojis$	14.726	0.175	0.574	0.574	712
0.2%	12	0	0	DELETE	^/api/posts/[^/]+$	13.762	1.147	2.795	2.795	0
0.1%	64	1	0	GET	^/api/posts/[^/]+/comments$	10.929	0.171	0.553	0.553	1383
0.1%	11	0	0	DELETE	^/api/realmojis/[^/]+$	9.655	0.878	1.936	1.936	0
0.1%	13	0	0	DELETE	^/api/comments/[^/]+$	4.246	0.327	0.992	0.992	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.3%	275	0	0	GET	/api/notifications	2116.756	7.697	9.944	10.395	179
12.6%	1387	11	0	GET	^/api/images/[^/]+	662.901	0.478	2.947	4.213	161989
11.7%	343	10	0	POST	/api/posts	616.996	1.799	5.965	6.345	622
7.2%	607	4	0	POST	/api/friend-requests	378.772	0.624	3.176	3.876	114
4.4%	904	0	0	GET	^/api/feeds/[^/]+	229.866	0.254	0.909	1.144	2093
4.2%	413	0	0	POST	/api/auth/login	222.014	0.538	2.776	3.257	250
3.4%	476	0	0	GET	/api/posting-window/current	179.817	0.378	1.013	1.699	239
2.6%	401	0	0	GET	/api/memories	138.945	0.346	1.166	1.824	13334
2.3%	200	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	120.761	0.604	3.060	3.334	115
1.6%	299	0	0	GET	/api/users	82.443	0.276	0.654	0.659	229
1.5%	275	16	0	GET	^/api/posts/me/current$	79.774	0.290	0.824	1.150	604
1.3%	26	0	0	GET	/api/me/export	66.905	2.573	6.868	6.868	3185991
0.9%	55	0	0	PUT	/api/me/profile-image	49.431	0.899	4.609	4.609	0
0.6%	63	0	0	POST	/api/auth/logout	34.116	0.542	3.106	3.106	0
0.6%	107	0	0	POST	/api/auth/register	34.080	0.319	1.132	1.150	229
0.5%	41	0	0	PATCH	/api/me	27.666	0.675	3.198	3.198	285
0.5%	66	0	0	PUT	^/api/follows/[^/]+$	25.782	0.391	2.759	2.759	0
0.5%	58	1	0	POST	^/api/posts/[^/]+/comments$	24.169	0.417	2.682	2.682	299
0.4%	110	14	0	GET	^/api/posts/[^/]+$	23.501	0.214	0.602	0.610	541
0.4%	115	0	0	GET	^/api/relationships/[^/]+	22.741	0.198	0.608	0.632	1158
0.4%	18	3	0	POST	^/api/posts/[^/]+/realmojis$	18.783	1.044	3.506	3.506	413
0.4%	65	0	0	GET	^/api/memories/[^/]+	18.720	0.288	0.775	0.775	955
0.3%	28	1	0	PATCH	^/api/posts/[^/]+$	16.685	0.596	1.516	1.516	618
0.3%	11	0	0	DELETE	^/api/posts/[^/]+$	13.884	1.262	2.953	2.953	0
0.2%	58	0	0	GET	/api/me	12.754	0.220	0.658	0.658	300
0.2%	47	0	0	GET	^/api/users/[^/]+$	11.017	0.234	0.626	0.626	214
0.1%	31	0	0	GET	^/api/posts/[^/]+/comments$	7.116	0.230	0.527	0.527	1582
0.1%	23	0	0	GET	^/api/posts/[^/]+/realmojis$	4.833	0.210	0.757	0.757	615
0.1%	12	0	0	DELETE	^/api/comments/[^/]+$	4.601	0.383	1.182	1.182	0
0.1%	6	0	0	PUT	^/api/blocks/[^/]+$	3.689	0.615	1.227	1.227	0
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 85250 クエリ / 87 種類 / 合計 141.614s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	62.5%	364	88.440	242.97	918.57	1234.7	125.9	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	13.0%	11153	18.440	1.65	13.27	9.5	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
3	5.7%	3162	8.091	2.56	17.13	97.4	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $2) AS is_profile, EXISTS (SELECT $3 FROM posts p WHERE p.image_id…
4	5.4%	2	7.709	3854.29	7561.72	231805.5	2729.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
5	3.3%	3283	4.619	1.41	17.52	24.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.7%	1	2.377	2376.56	2376.56	31200.0	1200069.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	1.2%	1	1.637	1637.11	1637.11	24949.0	627713.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	1.1%	633	1.588	2.51	9.81	28.7	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.9%	2	1.290	644.90	1281.98	16763.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
10	0.9%	7367	1.249	0.17	5.61	19.5	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.8%	995	1.139	1.14	8.92	19.6	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.6%	1	0.857	857.15	857.15	12846.0	530199.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
13	0.5%	6177	0.672	0.11	8.76	10.9	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
14	0.4%	3053	0.518	0.17	8.08	19.9	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.3%	4629	0.467	0.10	8.86	8.5	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$4::bytea,$5,$3)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT data FROM image_objects WHERE id = $1 AND created_at = $2;
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $2) AS is_profile, EXISTS (SELECT $3 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $4 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 311fa1597f6042f4c9f0ff7e04b5b1a249277462
Type: cpu
Time: 2026-10-31 14:06:10 JST
Duration: 60.05s, Total samples = 35.28s (58.75%)
Showing nodes accounting for 24.92s, 70.63% of 35.28s total
Dropped 1032 nodes (cum <= 0.18s)
Showing top 40 nodes out of 291
      flat  flat%   sum%        cum   cum%
     4.22s 11.96% 11.96%      4.22s 11.96%  crypto/internal/fips140/sha256.blockAVX2
     3.89s 11.03% 22.99%      3.89s 11.03%  internal/runtime/syscall/linux.Syscall6
     1.55s  4.39% 27.38%      5.62s 15.93%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.43s  4.05% 31.43%      1.43s  4.05%  runtime.memmove
     1.34s  3.80% 35.23%      1.34s  3.80%  aeshashbody
     1.28s  3.63% 38.86%      1.28s  3.63%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.28s  3.63% 42.49%      1.28s  3.63%  runtime.memclrNoHeapPointers
     1.28s  3.63% 46.12%      4.66s 13.21%  runtime.scanObject
     1.18s  3.34% 49.46%      1.59s  4.51%  runtime.findObject
     0.85s  2.41% 51.87%      1.49s  4.22%  runtime.tryDeferToSpanScan
     0.57s  1.62% 53.49%      2.19s  6.21%  runtime.mapaccess1_faststr
     0.42s  1.19% 54.68%      0.42s  1.19%  memeqbody
     0.42s  1.19% 55.87%      0.42s  1.19%  runtime.spanClass.sizeclass (inline)
     0.37s  1.05% 56.92%      0.60s  1.70%  example.com/isureal-company/isureal.postLess
     0.32s  0.91% 57.82%      0.32s  0.91%  runtime.futex
     0.30s  0.85% 58.67%      4.28s 12.13%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.30s  0.85% 59.52%      0.47s  1.33%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.29s  0.82% 60.35%      1.07s  3.03%  runtime.mallocgcSmallScanNoHeader
     0.28s  0.79% 61.14%      0.28s  0.79%  runtime.nextFreeFast (inline)
     0.27s  0.77% 61.90%      0.59s  1.67%  database/sql.convertAssignRows
     0.21s   0.6% 62.50%      3.20s  9.07%  runtime.mallocgc
     0.20s  0.57% 63.07%      0.40s  1.13%  compress/flate.(*compressor).deflate
     0.20s  0.57% 63.63%      0.20s  0.57%  indexbody
     0.20s  0.57% 64.20%      0.20s  0.57%  sync/atomic.(*Int32).Add (inline)
     0.19s  0.54% 64.74%      4.62s 13.10%  crypto/internal/fips140/sha256.(*Digest).Write
     0.19s  0.54% 65.28%      0.20s  0.57%  encoding/json.appendString[go.shape.string]
     0.18s  0.51% 65.79%      0.18s  0.51%  runtime.(*mspan).divideByElemSize (inline)
     0.18s  0.51% 66.30%      1.80s  5.10%  runtime.mapassign_faststr
     0.17s  0.48% 66.78%      0.18s  0.51%  example.com/isureal-company/isureal.rmLess
     0.16s  0.45% 67.23%      4.46s 12.64%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.16s  0.45% 67.69%      2.71s  7.68%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.15s  0.43% 68.11%      4.69s 13.29%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.15s  0.43% 68.54%      0.18s  0.51%  runtime.(*mspan).writeHeapBitsSmall
     0.12s  0.34% 68.88%      0.69s  1.96%  encoding/json.structEncoder.encode
     0.11s  0.31% 69.19%      5.01s 14.20%  example.com/isureal-company/isureal.computeImageSignature
     0.11s  0.31% 69.50%      0.34s  0.96%  internal/runtime/maps.(*Iter).Next
     0.10s  0.28% 69.78%      0.48s  1.36%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
     0.10s  0.28% 70.07%      3.69s 10.46%  example.com/isureal-company/isureal.loadBig
     0.10s  0.28% 70.35%      0.38s  1.08%  runtime.greyobject
     0.10s  0.28% 70.63%      0.54s  1.53%  runtime.scanSpan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 45.3% / 100% | 29% | 1.6% | 8 | 123 MB | run queue が積もっている |
| isu2 (2 cpu) | 17.5% / 99% | 2% | 0.5% | 3 | 1885 MB |  |
| isu3 (2 cpu) | 16.3% / 99% | 1% | 0.5% | 4 | 1689 MB |  |
| isu4 (2 cpu) | 36.7% / 100% | 11% | 17.8% | 14 | 107 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 64% / nginx 11% |
| isu2 | isureal 33% |
| isu3 | isureal 30% |
| isu4 | postgres 19% / isureal 15% |

## Diff from previous run

スコア: 82880 → 134480 (+51600)

`25c2db4..04990e1`

```
 webapp/go/export_api.go | 60 +++++++++++++++++++++++++++++++++++++++++--------
 webapp/go/image_file.go | 20 +++++++++--------
 2 files changed, 62 insertions(+), 18 deletions(-)
```
