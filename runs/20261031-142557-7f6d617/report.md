# run 20261031-142557-7f6d617

## Summary

| | |
|---|---|
| score | 137881 |
| status | pass |
| ref | 7f6d617 |
| sha | 7f6d61729ef7 |
| 所要時間 | 126s |
| 前回比 | -2149 (-1.5%) vs 20261031-141745-a28a701 |
| ベスト比 | -3454 (-2.4%) vs 20261031-141027-bd0fe45 = 141335 |
| note | main: +#75 複製（確認） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	142	412	1081	1080
europe	7	20	63	175	282	282
america	6	15	49	137	137	137
```

```
{"id":"17391","status":"done","portalStatus":"finished","passed":true,"score":137881,"raw":137881,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:26:25.329Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:26:28.294Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:26:31.292Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:26:34.320Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T05:26:37.292Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T05:26:40.292Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T05:26:43.319Z level=INFO msg="asia 3 日目: 142 人"
time=2026-10-31T05:26:44.288Z level=INFO msg="user moved" user=torsa-5629 from=asia to=america
time=2026-10-31T05:26:44.614Z level=INFO msg="user moved" user=mizoe5650 from=asia to=america
time=2026-10-31T05:26:45.190Z level=INFO msg="user moved" user=wadan-5807 from=asia to=europe
time=2026-10-31T05:26:45.684Z level=INFO msg="user moved" user=torsa-5629 from=america to=europe
time=2026-10-31T05:26:46.088Z level=INFO msg="user moved" user=wadan-5807 from=europe to=asia
time=2026-10-31T05:26:46.183Z level=INFO msg="user moved" user=solbenni.5726 from=asia to=europe
time=2026-10-31T05:26:46.291Z level=INFO msg="europe 3 日目: 63 人"
time=2026-10-31T05:26:46.543Z level=INFO msg="user moved" user=nelia-6082 from=asia to=america
time=2026-10-31T05:26:47.669Z level=INFO msg="user moved" user=nelia-6082 from=america to=asia
time=2026-10-31T05:26:48.809Z level=INFO msg="user moved" user=mizoe5650 from=america to=europe
time=2026-10-31T05:26:49.292Z level=INFO msg="america 3 日目: 49 人"
time=2026-10-31T05:26:50.839Z level=INFO msg="user moved" user=tofumo6373 from=europe to=america
time=2026-10-31T05:26:52.324Z level=INFO msg="asia 4 日目: 412 人"
time=2026-10-31T05:26:54.212Z level=INFO msg="user moved" user=nelia-6082 from=asia to=america
time=2026-10-31T05:26:55.437Z level=INFO msg="europe 4 日目: 175 人"
time=2026-10-31T05:26:56.157Z level=INFO msg="user moved" user=wadan-5807 from=asia to=europe
time=2026-10-31T05:26:56.201Z level=INFO msg="user moved" user=tasare7070 from=asia to=america
time=2026-10-31T05:26:56.434Z level=INFO msg="user moved" user=hefuchi.7231 from=asia to=america
time=2026-10-31T05:26:56.914Z level=INFO msg="user moved" user=semahe.7561 from=asia to=america
time=2026-10-31T05:26:58.248Z level=INFO msg="user moved" user=hetorke8140 from=asia to=america
time=2026-10-31T05:26:58.315Z level=INFO msg="user moved" user=ivyeva_8493 from=asia to=america
time=2026-10-31T05:26:58.396Z level=INFO msg="america 4 日目: 137 人"
time=2026-10-31T05:26:59.274Z level=INFO msg="user moved" user=yaru-8266 from=asia to=europe
time=2026-10-31T05:27:00.629Z level=INFO msg="user moved" user=torsa-5629 from=europe to=america
time=2026-10-31T05:27:01.575Z level=INFO msg="asia 5 日目: 1081 人"
time=2026-10-31T05:27:04.339Z level=INFO msg="user moved" user=tofumo6373 from=america to=europe
time=2026-10-31T05:27:04.347Z level=INFO msg="europe 5 日目: 282 人"
time=2026-10-31T05:27:07.498Z level=INFO msg="america 5 日目: 137 人"
time=2026-10-31T05:27:16.720Z level=INFO msg="europe 最終日（グランプリ）: 282 人"
time=2026-10-31T05:27:17.302Z level=INFO msg="asia 最終日（グランプリ）: 1080 人"
time=2026-10-31T05:27:17.540Z level=INFO msg="america 最終日（グランプリ）: 137 人"
time=2026-10-31T05:27:25.337Z level=INFO msg="グランプリ: アジア 324 人, ヨーロッパ 149 人, アメリカ 65 人"
{
  "pass": true,
  "score": 137881,
  "messages": [
    "mitaho.5060: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "sari.13360: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "vinmi6942: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "croquette_sodachi_2784: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "hemawa_7051: 友達申請を承認できません: POST /api/friend-requests/93b8a7a9-e608-41bd-bd42-1c36376a6afd/accept: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests/93b8a7a9-e608-41bd-bd42-1c36376a6afd/accept\": EOF",
    "wase7342: 友達申請を承認できません: POST /api/friend-requests/c48d005b-31f8-4ebe-a5bb-9fb3a68cf21b/accept: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/friend-requests/c48d005b-31f8-4ebe-a5bb-9fb3a68cf21b/accept\": EOF",
    "hamoni15213: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "miivyzoe_14765: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.3%	1702	1	0	GET	/api/notifications	13995.540	8.223	15.899	15.920	166
18.5%	2520	100	0	POST	/api/posts	7538.329	2.991	8.842	11.362	625
11.9%	5960	78	0	POST	/api/friend-requests	4857.938	0.815	5.132	10.370	114
8.4%	15955	471	0	GET	^/api/images/[^/]+	3411.140	0.214	3.454	8.809	131838
6.0%	6191	1080	0	POST	/api/auth/login	2460.095	0.397	5.012	10.184	205
4.2%	2255	8	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1704.629	0.756	4.636	7.391	116
3.5%	738	129	0	PUT	/api/me/profile-image	1416.879	1.920	9.874	11.197	3
2.0%	715	45	0	POST	^/api/posts/[^/]+/realmojis$	805.064	1.126	7.775	8.966	443
1.5%	1108	1	0	POST	/api/auth/register	610.577	0.551	5.171	7.141	227
1.5%	707	0	0	PUT	^/api/follows/[^/]+$	605.401	0.856	5.188	8.769	0
1.4%	3658	0	0	GET	/api/memories	578.297	0.158	1.060	2.322	12455
1.2%	557	0	0	PATCH	/api/me	476.913	0.856	5.294	7.885	278
1.1%	13450	0	0	GET	^/api/feeds/[^/]+	441.702	0.033	0.252	1.292	2813
1.0%	548	15	0	POST	^/api/posts/[^/]+/comments$	398.141	0.727	4.911	6.328	322
0.6%	806	0	0	POST	/api/auth/logout	253.529	0.315	4.271	6.345	0
0.6%	307	0	0	GET	/api/me/export	248.553	0.810	4.080	5.947	1846911
0.4%	230	8	0	PATCH	^/api/posts/[^/]+$	181.812	0.790	5.338	6.222	623
0.3%	3731	0	0	GET	/api/posting-window/current	134.369	0.036	0.233	0.843	237
0.3%	96	0	0	DELETE	^/api/posts/[^/]+$	130.933	1.364	6.391	6.391	0
0.2%	2882	0	0	GET	/api/users	97.526	0.034	0.224	0.565	235
0.2%	1836	91	0	GET	^/api/posts/me/current$	74.292	0.040	0.300	0.851	618
0.2%	124	0	0	DELETE	^/api/realmojis/[^/]+$	72.670	0.586	3.530	3.891	0
0.2%	2000	23	0	GET	^/api/posts/[^/]+$	61.266	0.031	0.261	1.119	626
0.1%	1696	0	0	GET	^/api/relationships/[^/]+	58.293	0.034	0.304	3.319	947
0.1%	806	0	0	GET	^/api/memories/[^/]+	51.127	0.063	0.493	1.017	633
0.1%	986	0	0	GET	^/api/users/[^/]+$	28.961	0.029	0.236	0.438	210
0.1%	621	3	0	GET	/api/me	27.091	0.044	0.253	1.381	301
0.1%	85	0	0	DELETE	^/api/comments/[^/]+$	22.633	0.266	2.445	2.445	0
0.1%	952	0	0	GET	^/api/posts/[^/]+/realmojis$	21.517	0.023	0.218	0.506	1413
0.0%	714	0	0	GET	^/api/posts/[^/]+/comments$	17.012	0.024	0.184	0.249	1783
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.1%	495	1	0	GET	/api/notifications	4142.164	8.368	12.323	12.949	165
11.9%	664	14	0	POST	/api/posts	1293.318	1.948	6.951	9.793	634
10.5%	1314	9	0	POST	/api/friend-requests	1138.196	0.866	4.543	7.520	114
8.9%	2829	23	0	GET	^/api/images/[^/]+	962.232	0.340	2.621	3.381	174504
6.1%	864	0	0	POST	/api/auth/login	666.755	0.772	5.490	7.535	246
3.8%	2192	0	0	GET	^/api/feeds/[^/]+	418.213	0.191	1.018	1.614	1762
3.0%	302	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	330.561	1.095	5.055	6.459	116
2.8%	276	0	0	POST	/api/auth/register	302.563	1.096	7.243	8.231	228
1.7%	958	0	0	GET	/api/posting-window/current	187.979	0.196	0.951	1.275	239
1.6%	151	0	0	PUT	/api/me/profile-image	169.930	1.125	4.498	4.710	0
1.4%	178	0	0	PUT	^/api/follows/[^/]+$	147.188	0.827	4.287	4.418	0
1.2%	835	0	0	GET	/api/memories	135.483	0.162	1.501	2.236	12184
1.2%	644	0	0	GET	/api/users	127.520	0.198	0.969	1.205	232
1.1%	142	0	0	PATCH	/api/me	119.125	0.839	4.346	4.515	276
1.1%	84	13	0	POST	^/api/posts/[^/]+/realmojis$	116.224	1.384	7.810	7.810	377
0.9%	564	33	0	GET	^/api/posts/me/current$	101.232	0.179	1.189	1.275	609
0.8%	127	0	0	POST	/api/auth/logout	89.271	0.703	3.908	4.759	0
0.8%	102	1	0	POST	^/api/posts/[^/]+/comments$	85.433	0.838	4.380	4.494	348
0.5%	30	0	0	DELETE	^/api/posts/[^/]+$	52.820	1.761	7.570	7.570	0
0.5%	268	13	0	GET	^/api/posts/[^/]+$	50.410	0.188	0.937	1.214	592
0.4%	314	0	0	GET	^/api/relationships/[^/]+	45.425	0.145	0.920	1.249	882
0.3%	49	3	0	PATCH	^/api/posts/[^/]+$	37.397	0.763	3.034	3.034	595
0.2%	160	1	0	GET	/api/me	26.485	0.166	0.849	0.879	298
0.2%	45	0	0	GET	/api/me/export	20.408	0.454	6.805	6.805	2674858
0.2%	121	0	0	GET	^/api/users/[^/]+$	19.127	0.158	0.997	1.162	196
0.2%	103	1	0	GET	^/api/posts/[^/]+/comments$	18.547	0.180	0.913	0.914	1430
0.2%	83	1	0	GET	^/api/posts/[^/]+/realmojis$	17.058	0.206	1.088	1.088	1004
0.2%	111	0	0	GET	^/api/memories/[^/]+	16.559	0.149	1.382	1.588	908
0.1%	6	0	0	DELETE	/api/me/profile-image	11.885	1.981	3.245	3.245	0
0.1%	9	0	0	PUT	^/api/blocks/[^/]+$	10.417	1.157	3.931	3.931	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.4%	298	0	0	GET	/api/notifications	2217.826	7.442	10.201	10.565	170
12.3%	729	10	0	POST	/api/friend-requests	747.808	1.026	5.762	7.938	114
10.7%	356	12	0	POST	/api/posts	652.710	1.833	5.749	6.558	627
10.5%	1823	6	0	GET	^/api/images/[^/]+	638.888	0.350	2.546	3.308	148101
6.3%	533	0	0	POST	/api/auth/login	385.139	0.723	4.603	6.738	250
3.5%	1116	0	0	GET	^/api/feeds/[^/]+	214.001	0.192	1.140	1.636	2803
3.1%	163	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	189.422	1.162	4.726	6.514	116
2.2%	516	0	0	GET	/api/posting-window/current	131.836	0.255	1.158	1.287	239
1.7%	472	0	0	GET	/api/memories	104.091	0.221	1.781	1.933	15927
1.6%	68	0	0	PUT	/api/me/profile-image	99.198	1.459	6.755	6.755	0
1.4%	357	0	0	GET	/api/users	83.267	0.233	1.035	1.328	242
1.3%	80	0	0	PUT	^/api/follows/[^/]+$	81.422	1.018	5.509	5.509	0
1.3%	129	0	0	POST	/api/auth/register	77.491	0.601	2.574	2.779	229
1.2%	81	0	0	POST	/api/auth/logout	71.368	0.881	5.780	5.780	0
1.1%	64	0	0	PATCH	/api/me	69.602	1.088	9.033	9.033	267
1.0%	291	17	0	GET	^/api/posts/me/current$	63.415	0.218	1.151	1.267	612
0.5%	47	1	0	POST	^/api/posts/[^/]+/comments$	28.436	0.605	4.521	4.521	343
0.4%	26	3	0	PATCH	^/api/posts/[^/]+$	26.800	1.031	3.818	3.818	573
0.4%	39	0	0	POST	^/api/posts/[^/]+/realmojis$	25.441	0.652	3.505	3.505	431
0.4%	161	8	0	GET	^/api/posts/[^/]+$	25.281	0.157	0.886	0.886	607
0.4%	27	0	0	GET	/api/me/export	25.012	0.926	4.242	4.242	2788444
0.4%	70	0	0	GET	/api/me	24.202	0.346	0.965	0.965	304
0.4%	149	0	0	GET	^/api/relationships/[^/]+	23.238	0.156	1.029	1.356	877
0.3%	79	0	0	GET	^/api/memories/[^/]+	18.122	0.229	1.612	1.612	1096
0.3%	9	0	0	DELETE	^/api/posts/[^/]+$	16.966	1.885	4.955	4.955	0
0.3%	98	0	0	GET	^/api/users/[^/]+$	16.415	0.167	0.944	0.944	205
0.2%	76	0	0	GET	^/api/posts/[^/]+/comments$	10.030	0.132	0.895	0.895	1599
0.1%	65	0	0	GET	^/api/posts/[^/]+/realmojis$	8.947	0.138	0.834	0.834	1332
0.1%	11	0	0	DELETE	^/api/realmojis/[^/]+$	7.154	0.650	2.869	2.869	0
0.0%	5	0	0	PUT	^/api/blocks/[^/]+$	2.350	0.470	1.778	1.778	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 87811 クエリ / 85 種類 / 合計 175.236s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	47.0%	319	82.290	257.96	895.89	1202.2	122.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	12.1%	3	21.163	7054.31	9534.19	31556.0	1200034.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	10.6%	10815	18.631	1.72	15.19	9.6	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	9.1%	3	15.993	5330.92	7213.86	24949.0	628470.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.2%	3871	9.087	2.35	17.05	96.5	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $2) AS is_profile, EXISTS (SELECT $3 FROM posts p WHERE p.image_id…
6	4.5%	3	7.907	2635.83	3677.80	13206.0	529845.3	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
7	2.6%	3426	4.617	1.35	91.98	24.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
8	2.1%	6	3.723	620.54	2565.19	213146.2	2616.2	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
9	1.1%	6	1.847	307.91	457.35	27833.2	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
10	1.0%	780	1.833	2.35	11.15	28.5	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
11	0.7%	7885	1.293	0.16	4.17	19.6	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
12	0.5%	680	0.827	1.22	5.08	21.8	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
13	0.4%	6	0.694	115.59	156.99	624.0	30194.5	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
14	0.4%	6	0.692	115.40	158.18	625.2	25448.5	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
15	0.4%	6508	0.645	0.10	6.21	11.1	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

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
Build ID: 902005c20b08fa1e3c1c5623abc19a0b90cf0801
Type: cpu
Time: 2026-10-31 14:26:03 JST
Duration: 60s, Total samples = 41.27s (68.78%)
Showing nodes accounting for 28.73s, 69.61% of 41.27s total
Dropped 1137 nodes (cum <= 0.21s)
Showing top 40 nodes out of 310
      flat  flat%   sum%        cum   cum%
     5.74s 13.91% 13.91%      5.74s 13.91%  crypto/internal/fips140/sha256.blockAVX2
     5.47s 13.25% 27.16%      5.47s 13.25%  internal/runtime/syscall/linux.Syscall6
     1.75s  4.24% 31.40%      1.75s  4.24%  runtime.memmove
     1.61s  3.90% 35.30%      1.61s  3.90%  runtime.memclrNoHeapPointers
     1.57s  3.80% 39.11%      1.57s  3.80%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.48s  3.59% 42.69%      5.76s 13.96%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.27s  3.08% 45.77%      1.71s  4.14%  runtime.findObject
     1.16s  2.81% 48.58%      5.09s 12.33%  runtime.scanObject
     1.16s  2.81% 51.39%      1.91s  4.63%  runtime.tryDeferToSpanScan
     1.10s  2.67% 54.06%      1.10s  2.67%  aeshashbody
     0.62s  1.50% 55.56%      2.16s  5.23%  runtime.mapaccess1_faststr
     0.48s  1.16% 56.72%      0.48s  1.16%  runtime.spanClass.sizeclass (inline)
     0.46s  1.11% 57.84%      0.46s  1.11%  runtime.nextFreeFast (inline)
     0.40s  0.97% 58.81%      0.58s  1.41%  example.com/isureal-company/isureal.postLess
     0.38s  0.92% 59.73%      0.38s  0.92%  runtime.futex
     0.34s  0.82% 60.55%      0.34s  0.82%  memeqbody
     0.25s  0.61% 61.16%      5.79s 14.03%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.25s  0.61% 61.76%      0.25s  0.61%  sync/atomic.(*Int32).Add
     0.24s  0.58% 62.35%      0.59s  1.43%  database/sql.convertAssignRows
     0.24s  0.58% 62.93%      1.11s  2.69%  runtime.mallocgcSmallScanNoHeader
     0.23s  0.56% 63.48%      0.52s  1.26%  runtime.greyobject
     0.23s  0.56% 64.04%      1.80s  4.36%  runtime.mapassign_faststr
     0.22s  0.53% 64.57%      3.45s  8.36%  runtime.mallocgc
     0.21s  0.51% 65.08%      6.19s 15.00%  crypto/internal/fips140/sha256.(*Digest).Write
     0.21s  0.51% 65.59%      0.22s  0.53%  runtime.typePointers.next
     0.17s  0.41% 66.00%      5.47s 13.25%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.16s  0.39% 66.39%      6.08s 14.73%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.15s  0.36% 66.76%      0.41s  0.99%  runtime.scanObjectsSmall
     0.14s  0.34% 67.09%      0.31s  0.75%  compress/flate.(*compressor).deflate
     0.13s  0.31% 67.41%      0.27s  0.65%  internal/runtime/maps.(*Iter).Next
     0.12s  0.29% 67.70%      0.46s  1.11%  runtime.mapaccess2_faststr
     0.11s  0.27% 67.97%      6.51s 15.77%  example.com/isureal-company/isureal.computeImageSignature
     0.10s  0.24% 68.21%      3.45s  8.36%  example.com/isureal-company/isureal.listMemories
     0.09s  0.22% 68.43%      0.24s  0.58%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.09s  0.22% 68.65%      2.64s  6.40%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.08s  0.19% 68.84%      0.22s  0.53%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.08s  0.19% 69.03%      0.23s  0.56%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.08s  0.19% 69.23%      0.62s  1.50%  encoding/json.structEncoder.encode
     0.08s  0.19% 69.42%      0.28s  0.68%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
     0.08s  0.19% 69.61%      0.96s  2.33%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 47.9% / 100% | 33% | 0.7% | 5 | 112 MB | CPU が飽和 |
| isu2 (2 cpu) | 23.7% / 86% | 0% | 0.9% | 5 | 182 MB | run queue が積もっている |
| isu3 (2 cpu) | 24.4% / 92% | 1% | 1.2% | 4 | 774 MB |  |
| isu4 (2 cpu) | 39.9% / 99% | 12% | 18.1% | 19 | 100 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 77% / nginx 12% |
| isu2 | isureal 44% / nginx 1% |
| isu3 | isureal 45% / nginx 1% |
| isu4 | postgres 27% / isureal 12% |

## Diff from previous run

スコア: 140030 → 137881 (-2149)

`a28a701..7f6d617`

```
 webapp/go/export_api.go        |   2 +
 webapp/go/image_api.go         |  17 +-----
 webapp/go/image_file.go        | 132 +++++++++++++++++++++++++++++++++++++++++
 webapp/go/image_local.go       |  40 ++++++-------
 webapp/go/initialize.go        |   1 +
 webapp/go/post_api.go          |   5 +-
 webapp/go/profile_image_api.go |   5 +-
 webapp/go/realmoji_api.go      |   5 +-
 webapp/go/replica_server.go    |  13 +++-
 webapp/go/server.go            |   1 +
 10 files changed, 180 insertions(+), 41 deletions(-)
```
