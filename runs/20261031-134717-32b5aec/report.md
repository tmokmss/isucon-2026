# run 20261031-134717-32b5aec

## Summary

| | |
|---|---|
| score | 91214 |
| status | pass |
| ref | image-files |
| sha | 32b5aecfa0ec |
| 所要時間 | 325s |
| 前回比 | +16002 (21.3%) vs 20261031-133811-f498277 |
| ベスト比 | +16002 (21.3%) vs 20261031-133811-f498277 = 75212 **ベスト更新** |
| note | image bodies to files + main fbae0df (#78) |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	150	430	789	789
europe	7	18	51	137	140	140
america	6	14	49	125	124	124
```

```
{"id":"14514","status":"done","portalStatus":"finished","passed":true,"score":91214,"raw":91214,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:51:03.546Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:51:06.517Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:51:09.518Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:51:12.539Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T04:51:15.520Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T04:51:18.519Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T04:51:21.538Z level=INFO msg="asia 3 日目: 150 人"
time=2026-10-31T04:51:22.152Z level=INFO msg="user moved" user=mosueva-5546 from=asia to=america
time=2026-10-31T04:51:22.516Z level=INFO msg="user moved" user=toliama-5634 from=asia to=america
time=2026-10-31T04:51:24.011Z level=INFO msg="user moved" user=toliama-5634 from=america to=europe
time=2026-10-31T04:51:24.565Z level=INFO msg="europe 3 日目: 51 人"
time=2026-10-31T04:51:24.588Z level=INFO msg="user moved" user=wari5814 from=asia to=america
time=2026-10-31T04:51:24.845Z level=INFO msg="user moved" user=miko_5885 from=asia to=america
time=2026-10-31T04:51:27.094Z level=INFO msg="user moved" user=shihofu-6172 from=asia to=america
time=2026-10-31T04:51:27.118Z level=INFO msg="user moved" user=mosueva-5546 from=america to=europe
time=2026-10-31T04:51:27.539Z level=INFO msg="america 3 日目: 49 人"
time=2026-10-31T04:51:30.315Z level=INFO msg="user moved" user=ivyme_6321 from=europe to=asia
time=2026-10-31T04:51:30.573Z level=INFO msg="asia 4 日目: 430 人"
time=2026-10-31T04:51:34.937Z level=INFO msg="europe 4 日目: 137 人"
time=2026-10-31T04:51:35.115Z level=INFO msg="user moved" user=sarore_6696 from=asia to=europe
time=2026-10-31T04:51:36.050Z level=INFO msg="user moved" user=sari-7463 from=asia to=america
time=2026-10-31T04:51:37.777Z level=INFO msg="user moved" user=sari-7463 from=america to=europe
time=2026-10-31T04:51:38.412Z level=INFO msg="user moved" user=nohichi.7871 from=asia to=europe
time=2026-10-31T04:51:38.474Z level=INFO msg="america 4 日目: 125 人"
time=2026-10-31T04:51:39.333Z level=INFO msg="user moved" user=toliama-5634 from=europe to=america
time=2026-10-31T04:51:39.776Z level=INFO msg="asia 5 日目: 789 人"
time=2026-10-31T04:51:40.153Z level=INFO msg="user moved" user=ivyme_6321 from=asia to=europe
time=2026-10-31T04:51:49.172Z level=INFO msg="user moved" user=rata.9399 from=america to=europe
time=2026-10-31T04:51:49.320Z level=INFO msg="user moved" user=mosueva-5546 from=europe to=asia
time=2026-10-31T04:51:51.368Z level=INFO msg="user moved" user=miko_5885 from=america to=asia
time=2026-10-31T04:51:52.920Z level=INFO msg="america 5 日目: 124 人"
time=2026-10-31T04:51:53.063Z level=INFO msg="europe 5 日目: 140 人"
time=2026-10-31T04:51:56.240Z level=INFO msg="asia 最終日（グランプリ）: 789 人"
time=2026-10-31T04:51:57.354Z level=INFO msg="america 最終日（グランプリ）: 124 人"
time=2026-10-31T04:51:58.075Z level=INFO msg="europe 最終日（グランプリ）: 140 人"
time=2026-10-31T04:52:03.551Z level=INFO msg="グランプリ: アジア 56 人, ヨーロッパ 75 人, アメリカ 122 人"
{
  "pass": true,
  "score": 91214,
  "messages": [
    "tao_zhou_85: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "yuultsu-8414: 友達申請を送れません: POST /api/friend-requests: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests\": EOF",
    "moisre.9781: 友達申請を送れません: POST /api/friend-requests: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests\": EOF",
    "duncai_ayi_2630: ログアウトできません: POST /api/auth/logout: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/logout\": EOF",
    "seoyeon_lee_2250: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "dalbam_mechurial_4472: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "liara9447: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ren_watanabe_4484: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "min_zhang_4687: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "mosueva-5546: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sasolsu10443: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noazoete_10159: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tenere-9923: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mihi.10209: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "shisuno-9898: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kemo10532: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chiya9531: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kuyose9727: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "suki-9577: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ornuno_10694: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.0%	1341	51	0	GET	/api/notifications	11155.602	8.319	16.423	16.452	165
10.4%	1606	26	0	POST	/api/posts	2690.925	1.676	5.654	7.253	641
8.7%	9730	560	0	GET	^/api/images/[^/]+	2242.979	0.231	3.834	8.389	117678
5.8%	3566	50	0	POST	/api/friend-requests	1499.491	0.420	3.123	5.792	114
4.0%	5697	0	0	GET	^/api/feeds/[^/]+	1033.249	0.181	1.764	5.185	2215
3.7%	584	128	0	PUT	/api/me/profile-image	952.749	1.631	10.854	13.607	4
3.5%	4151	1004	0	POST	/api/auth/login	916.467	0.221	2.647	4.163	189
3.5%	2419	1	0	GET	/api/posting-window/current	901.166	0.373	3.834	5.185	237
3.1%	1424	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	803.864	0.565	3.986	7.253	116
2.6%	2558	0	0	GET	/api/memories	678.365	0.265	3.834	5.185	10889
2.3%	187	0	0	GET	/api/me/export	597.095	3.193	18.665	18.703	1985892
1.6%	1742	0	0	GET	/api/users	415.918	0.239	2.860	5.185	240
1.5%	1242	108	0	GET	^/api/posts/me/current$	376.622	0.303	2.860	5.185	597
1.0%	812	1	0	POST	/api/auth/register	255.320	0.314	1.451	3.981	227
0.8%	297	9	0	POST	^/api/posts/[^/]+/realmojis$	200.042	0.674	3.346	6.780	469
0.8%	486	0	0	PUT	^/api/follows/[^/]+$	195.180	0.402	3.834	4.019	0
0.6%	353	0	0	PATCH	/api/me	154.001	0.436	3.834	5.185	279
0.6%	465	3	0	GET	/api/me	143.617	0.309	1.632	3.986	300
0.5%	909	0	0	GET	^/api/relationships/[^/]+	121.971	0.134	1.580	3.834	1244
0.4%	461	0	0	POST	/api/auth/logout	101.483	0.220	3.621	3.986	0
0.3%	793	25	0	GET	^/api/posts/[^/]+$	82.093	0.104	1.057	1.764	596
0.3%	254	5	0	POST	^/api/posts/[^/]+/comments$	79.157	0.312	2.491	2.681	341
0.3%	415	0	0	GET	^/api/memories/[^/]+	76.110	0.183	3.123	4.237	663
0.3%	66	0	0	DELETE	^/api/posts/[^/]+$	65.585	0.994	6.487	6.487	0
0.2%	120	6	0	PATCH	^/api/posts/[^/]+$	48.337	0.403	2.554	2.761	624
0.2%	399	0	0	GET	^/api/users/[^/]+$	44.550	0.112	1.338	2.860	192
0.1%	1	0	0	POST	/api/initialize	18.510	18.510	18.510	18.510	90
0.1%	35	0	0	DELETE	^/api/realmojis/[^/]+$	16.395	0.468	4.413	4.413	0
0.1%	281	0	0	GET	^/api/posts/[^/]+/realmojis$	14.494	0.052	0.329	0.579	1032
0.0%	246	0	0	GET	^/api/posts/[^/]+/comments$	12.326	0.050	0.297	0.334	1010
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
47.3%	384	1	0	GET	/api/notifications	3417.739	8.900	17.815	17.843	174
11.3%	2639	13	0	GET	^/api/images/[^/]+	813.613	0.308	2.669	4.553	161476
10.4%	465	2	0	POST	/api/posts	749.659	1.612	4.032	4.182	654
5.9%	2625	0	0	GET	^/api/feeds/[^/]+	425.334	0.162	0.636	1.415	3235
3.9%	842	16	0	POST	/api/friend-requests	277.987	0.330	2.036	3.336	113
3.3%	726	0	0	POST	/api/auth/login	236.841	0.326	2.160	2.808	250
2.7%	51	0	0	GET	/api/me/export	196.182	3.847	11.332	11.332	1889073
2.5%	736	0	0	GET	/api/memories	178.682	0.243	1.015	3.831	12819
2.3%	631	0	0	GET	/api/posting-window/current	165.034	0.262	0.746	1.875	239
1.3%	355	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	90.969	0.256	1.801	2.928	116
1.2%	399	25	0	GET	^/api/posts/me/current$	87.933	0.220	1.261	1.477	615
1.0%	387	0	0	GET	/api/users	73.951	0.191	0.657	0.795	239
0.9%	352	0	0	GET	^/api/relationships/[^/]+	68.442	0.194	0.836	1.454	841
0.8%	418	4	0	GET	^/api/posts/[^/]+$	60.912	0.146	0.552	0.824	627
0.8%	84	0	0	PUT	/api/me/profile-image	55.694	0.663	3.164	3.164	0
0.6%	74	1	0	POST	^/api/posts/[^/]+/realmojis$	39.959	0.540	3.238	3.238	440
0.5%	97	0	0	PUT	^/api/follows/[^/]+$	35.990	0.371	1.928	1.928	0
0.5%	134	0	0	POST	/api/auth/register	34.154	0.255	1.045	1.070	228
0.5%	124	0	0	POST	/api/auth/logout	32.563	0.263	1.407	1.611	0
0.4%	221	0	0	GET	^/api/users/[^/]+$	30.884	0.140	0.451	0.587	207
0.4%	78	0	0	PATCH	/api/me	29.198	0.374	1.876	1.876	274
0.3%	19	0	0	DELETE	^/api/posts/[^/]+$	22.058	1.161	2.984	2.984	0
0.3%	44	10	0	PATCH	^/api/posts/[^/]+$	20.419	0.464	2.102	2.102	506
0.2%	94	0	0	GET	^/api/memories/[^/]+	17.545	0.187	1.002	1.002	838
0.2%	90	2	0	GET	/api/me	16.710	0.186	0.530	0.530	296
0.2%	80	0	0	GET	^/api/posts/[^/]+/realmojis$	11.267	0.141	0.604	0.604	799
0.1%	27	0	0	DELETE	^/api/realmojis/[^/]+$	9.444	0.350	1.393	1.393	0
0.1%	33	0	0	POST	^/api/posts/[^/]+/comments$	7.535	0.228	1.234	1.234	358
0.1%	44	0	0	GET	^/api/posts/[^/]+/comments$	5.565	0.126	0.404	0.404	592
0.0%	7	0	0	PUT	^/api/blocks/[^/]+$	2.623	0.375	1.085	1.085	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.6%	337	1	0	GET	/api/notifications	2394.678	7.106	14.626	15.290	177
13.6%	1953	13	0	GET	^/api/images/[^/]+	744.680	0.381	2.361	3.160	197020
10.3%	382	0	0	POST	/api/posts	565.732	1.481	5.949	6.612	646
5.8%	1843	0	0	GET	^/api/feeds/[^/]+	318.068	0.173	0.667	0.981	2514
3.4%	33	0	0	GET	/api/me/export	188.897	5.724	13.059	13.059	3025299
3.4%	520	0	0	POST	/api/auth/login	187.690	0.361	2.295	2.616	250
3.1%	744	9	0	POST	/api/friend-requests	171.965	0.231	1.649	2.360	114
2.9%	604	0	0	GET	/api/memories	160.872	0.266	1.148	1.742	13529
2.3%	550	0	0	GET	/api/posting-window/current	124.197	0.226	0.755	1.104	240
1.5%	379	64	0	GET	^/api/posts/me/current$	79.812	0.211	0.751	0.792	542
1.2%	121	0	0	POST	/api/auth/register	65.778	0.544	1.886	2.255	229
1.2%	374	0	0	GET	/api/users	63.299	0.169	0.658	0.693	240
1.1%	222	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	62.228	0.280	1.704	1.829	115
1.1%	59	0	0	PUT	/api/me/profile-image	58.722	0.995	4.253	4.253	0
1.0%	283	0	0	GET	^/api/relationships/[^/]+	54.819	0.194	0.622	0.770	732
0.7%	282	8	0	GET	^/api/posts/[^/]+$	40.190	0.143	0.437	0.719	615
0.6%	51	3	0	POST	^/api/posts/[^/]+/realmojis$	34.240	0.671	3.814	3.814	465
0.5%	84	0	0	POST	/api/auth/logout	25.888	0.308	2.529	2.529	0
0.4%	144	0	0	GET	^/api/users/[^/]+$	22.368	0.155	0.473	0.536	207
0.4%	40	7	0	PATCH	^/api/posts/[^/]+$	21.945	0.549	1.926	1.926	540
0.3%	21	0	0	DELETE	^/api/posts/[^/]+$	17.763	0.846	1.602	1.602	0
0.3%	81	0	0	PUT	^/api/follows/[^/]+$	17.009	0.210	0.792	0.792	0
0.3%	80	0	0	GET	^/api/memories/[^/]+	16.626	0.208	0.712	0.712	812
0.2%	53	0	0	PATCH	/api/me	13.674	0.258	1.483	1.483	265
0.2%	52	0	0	POST	^/api/posts/[^/]+/comments$	10.886	0.209	0.851	0.851	287
0.2%	63	1	0	GET	/api/me	10.165	0.161	0.476	0.476	296
0.1%	43	0	0	GET	^/api/posts/[^/]+/comments$	6.737	0.157	0.432	0.432	908
0.1%	38	0	0	GET	^/api/posts/[^/]+/realmojis$	5.252	0.138	0.540	0.540	1412
0.1%	9	0	0	DELETE	^/api/realmojis/[^/]+$	3.237	0.360	0.704	0.704	0
0.1%	15	0	0	DELETE	^/api/comments/[^/]+$	3.027	0.202	0.603	0.603	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 53503 クエリ / 79 種類 / 合計 25.970s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	26.5%	3618	6.891	1.90	17.99	104.9	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $2) AS is_profile, EXISTS (SELECT $3 FROM posts p WHERE p.image_id…
2	20.4%	2	5.300	2649.94	5140.65	238485.5	2809.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
3	11.7%	2435	3.047	1.25	9.50	25.2	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
4	9.3%	1	2.424	2423.59	2423.59	30923.0	1199942.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
5	7.4%	1	1.920	1920.49	1920.49	24184.0	627893.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
6	3.6%	409	0.936	2.29	9.86	31.9	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
7	3.0%	1	0.791	791.07	791.07	12606.0	530173.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	2.6%	5063	0.675	0.13	7.64	19.9	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
9	1.6%	334	0.415	1.24	5.21	22.6	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
10	1.5%	4393	0.397	0.09	6.41	11.3	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
11	1.3%	3443	0.333	0.10	10.18	9.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$4::bytea,$5,$3)
12	1.1%	1949	0.296	0.15	4.63	20.3	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
13	1.0%	2	0.265	132.55	257.13	16790.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
14	0.9%	1	0.244	244.45	244.45	29749.0	2782.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.9%	1	0.244	244.11	244.11	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $2) AS is_profile, EXISTS (SELECT $3 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $4 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
-- rank 2
SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id;
-- rank 3
INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $11);
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: c795435da5b949af002d8aac1d28b1779c7e22f4
Type: cpu
Time: 2026-10-31 13:50:40 JST
Duration: 60.06s, Total samples = 33.46s (55.71%)
Showing nodes accounting for 24.49s, 73.19% of 33.46s total
Dropped 962 nodes (cum <= 0.17s)
Showing top 40 nodes out of 286
      flat  flat%   sum%        cum   cum%
     5.64s 16.86% 16.86%      5.64s 16.86%  internal/runtime/syscall/linux.Syscall6
     2.60s  7.77% 24.63%      2.60s  7.77%  crypto/internal/fips140/sha256.blockAVX2
     1.91s  5.71% 30.33%      5.91s 17.66%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.57s  4.69% 35.03%      1.57s  4.69%  runtime.memclrNoHeapPointers
     1.44s  4.30% 39.33%      1.44s  4.30%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.14s  3.41% 42.74%      1.14s  3.41%  aeshashbody
     1.08s  3.23% 45.97%      1.39s  4.15%  runtime.findObject
     1.05s  3.14% 49.10%      4.38s 13.09%  runtime.scanObject
     1.03s  3.08% 52.18%      1.82s  5.44%  runtime.tryDeferToSpanScan
     0.91s  2.72% 54.90%      0.91s  2.72%  runtime.memmove
     0.62s  1.85% 56.75%      2.03s  6.07%  runtime.mapaccess1_faststr
     0.56s  1.67% 58.43%      0.56s  1.67%  runtime.spanClass.sizeclass (inline)
     0.42s  1.26% 59.68%      0.42s  1.26%  memeqbody
     0.42s  1.26% 60.94%      0.42s  1.26%  runtime.nextFreeFast (inline)
     0.40s  1.20% 62.13%      0.62s  1.85%  example.com/isureal-company/isureal.postLess
     0.31s  0.93% 63.06%      0.43s  1.29%  database/sql.convertAssignRows
     0.23s  0.69% 63.75%      0.25s  0.75%  runtime.(*mspan).writeHeapBitsSmall
     0.22s  0.66% 64.41%      0.22s  0.66%  indexbody
     0.21s  0.63% 65.03%      1.14s  3.41%  runtime.mallocgcSmallScanNoHeader
     0.21s  0.63% 65.66%      0.21s  0.63%  sync/atomic.(*Int32).Add (inline)
     0.19s  0.57% 66.23%      0.20s   0.6%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.18s  0.54% 66.77%      1.66s  4.96%  runtime.mapassign_faststr
     0.16s  0.48% 67.24%      0.19s  0.57%  runtime.heapArenaOf (inline)
     0.16s  0.48% 67.72%      3.05s  9.12%  runtime.mallocgc
     0.14s  0.42% 68.14%      0.31s  0.93%  compress/flate.(*compressor).deflate
     0.14s  0.42% 68.56%      0.51s  1.52%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
     0.14s  0.42% 68.98%      0.29s  0.87%  internal/runtime/maps.(*Iter).Next
     0.14s  0.42% 69.40%      0.39s  1.17%  runtime.greyobject
     0.13s  0.39% 69.78%      2.25s  6.72%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.13s  0.39% 70.17%      0.79s  2.36%  example.com/isureal-company/isureal.listMemories
     0.13s  0.39% 70.56%      0.19s  0.57%  runtime.(*mspan).markBitsForIndex (inline)
     0.12s  0.36% 70.92%      0.17s  0.51%  encoding/json.appendString[go.shape.string]
     0.11s  0.33% 71.25%      0.25s  0.75%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.10s   0.3% 71.55%      4.01s 11.98%  example.com/isureal-company/isureal.loadBig
     0.10s   0.3% 71.85%      0.65s  1.94%  runtime.mallocgcSmallNoscan
     0.10s   0.3% 72.15%      0.53s  1.58%  runtime.mapaccess2_faststr
     0.09s  0.27% 72.41%      2.36s  7.05%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.09s  0.27% 72.68%      2.12s  6.34%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.09s  0.27% 72.95%      0.34s  1.02%  internal/stringslite.Index
     0.08s  0.24% 73.19%      2.73s  8.16%  crypto/internal/fips140/sha256.(*Digest).Write
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 55.2% / 100% | 39% | 16.5% | 8 | 125 MB | CPU が飽和 |
| isu2 (2 cpu) | 16.8% / 98% | 2% | 0.6% | 3 | 1815 MB |  |
| isu3 (2 cpu) | 14.8% / 100% | 2% | 0.6% | 4 | 1623 MB |  |
| isu4 (2 cpu) | 13.5% / 74% | 0% | 5.1% | 21 | 202 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 54% / nginx 8% / kcompactd0 1% |
| isu2 | isureal 31% / nginx 1% |
| isu3 | isureal 26% / nginx 1% |
| isu4 | postgres 11% / isureal 3% |

## Diff from previous run

スコア: 75212 → 91214 (+16002)

`f498277..32b5aec`

```
 webapp/go/account_deletion_api.go |  1 +
 webapp/go/action_log.go           | 13 ++++--
 webapp/go/action_log_test.go      | 67 ++++++++++++++++++++++++++++
 webapp/go/offload.go              | 47 ++++++++++++++------
 webapp/go/profile_image_api.go    |  7 +--
 webapp/go/store.go                | 53 +++++++++++++++++++---
 webapp/go/user_api.go             |  1 +
 webapp/go/user_read_api.go        | 58 +++++++++++++-----------
 webapp/go/user_search_test.go     | 94 +++++++++++++++++++++++++++++++++++++++
 9 files changed, 289 insertions(+), 52 deletions(-)
```
