# run 20261031-113529-d1c842b

## Summary

| | |
|---|---|
| score | 21198 |
| status | pass |
| ref | limiter-priority |
| sha | d1c842b99f93 |
| 所要時間 | 219s |
| 前回比 | +21198 (—) vs 20261031-113154-1376dd7 |
| ベスト比 | -3603 (-14.5%) vs 20261031-103811-cef5fe3 = 24801 |
| note | limiter post4-8s heavy1 on main-6a7a33a |

## Bench

```
{"id":"4670","status":"done","portalStatus":"finished","passed":true,"score":21198,"raw":21198,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:37:56.042Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:37:58.902Z level=INFO msg="user moved" user=nesane-5052 from=asia to=europe
time=2026-10-31T02:37:59.006Z level=INFO msg="europe 1 日目: 8 人"
time=2026-10-31T02:38:01.998Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:38:05.040Z level=INFO msg="asia 2 日目: 53 人"
time=2026-10-31T02:38:06.844Z level=INFO msg="user moved" user=kona-5177 from=asia to=europe
time=2026-10-31T02:38:08.039Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T02:38:08.460Z level=INFO msg="user moved" user=toneku5308 from=asia to=europe
time=2026-10-31T02:38:11.000Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T02:38:14.024Z level=INFO msg="asia 3 日目: 152 人"
time=2026-10-31T02:38:17.140Z level=INFO msg="europe 3 日目: 49 人"
time=2026-10-31T02:38:20.068Z level=INFO msg="america 3 日目: 18 人"
time=2026-10-31T02:38:23.140Z level=INFO msg="asia 4 日目: 152 人"
time=2026-10-31T02:38:26.142Z level=INFO msg="europe 4 日目: 49 人"
time=2026-10-31T02:38:29.049Z level=INFO msg="america 4 日目: 18 人"
time=2026-10-31T02:38:32.215Z level=INFO msg="asia 5 日目: 152 人"
time=2026-10-31T02:38:35.044Z level=INFO msg="europe 5 日目: 49 人"
time=2026-10-31T02:38:38.035Z level=INFO msg="america 5 日目: 18 人"
time=2026-10-31T02:38:47.064Z level=INFO msg="america 最終日（グランプリ）: 18 人"
time=2026-10-31T02:38:47.119Z level=INFO msg="europe 最終日（グランプリ）: 49 人"
time=2026-10-31T02:38:47.208Z level=INFO msg="asia 最終日（グランプリ）: 152 人"
time=2026-10-31T02:38:55.998Z level=INFO msg="グランプリ: アジア 65 人, ヨーロッパ 17 人, アメリカ 12 人"
{
  "pass": true,
  "score": 21198,
  "messages": [
    "yuxin_wu_4950: 履歴を読めません: GET /api/memories?month=2026-04: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "rita_5296: Memories を取得できません: GET /api/memories?month=2026-11: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "moisre5320: Memories を取得できません: GET /api/memories?month=2026-11: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "evara.5935: プロフィール画像を変えられません: PUT /api/me/profile-image: ステータスコード 503 が返りました（期待値 204、busy）（5 回試行）",
    "techiben_5039: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "zoenoa_6127: プロフィール画像を変えられません: PUT /api/me/profile-image: ステータスコード 503 が返りました（期待値 204、busy）（5 回試行）",
    "jihun_park_57: 履歴を読めません: GET /api/memories?month=2026-05: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "hano6315: プロフィール画像を変えられません: PUT /api/me/profile-image: ステータスコード 503 が返りました（期待値 204、busy）（5 回試行）",
    "musol6294: プロフィール画像を変えられません: PUT /api/me/profile-image: ステータスコード 503 が返りました（期待値 204、busy）（5 回試行）",
    "hina_nakamura_38: 履歴を読めません: GET /api/memories?month=2026-06: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "fukemia-6263: プロフィール画像を変えられません: PUT /api/me/profile-image: ステータスコード 503 が返りました（期待値 204、busy）（5 回試行）",
    "noato.5192: Memories を取得できません: GET /api/memories?month=2026-11: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "yomaxtsu.5207: Memories を取得できません: GET /api/memories?month=2026-11: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "yorimichi_rocket_1474: 履歴を読めません: GET /api/memories/2026-05-27: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "mara.5329: Memories を取得できません: GET /api/memories?month=2026-11: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "salmon_bin_3577: 履歴を読めません: GET /api/memories?month=2026-07: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "gabriel_dubois_2059: 履歴を読めません: GET /api/memories/2025-12-28: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "zoemia-5077: Memories を取得できません: GET /api/memories?month=2026-11: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）",
    "liayunoa.5159: Memories を取得できません: GET /api/memories?month=2026-11: ステータスコード 503 が返りました（期待値 200、busy）（5 回試行）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.0%	522	0	0	GET	/api/notifications	4944.797	9.473	15.009	15.091	177
25.6%	530	6	86	POST	/api/posts	2942.304	5.552	8.388	8.483	536
18.4%	994	0	240	GET	/api/memories	2112.359	2.125	4.106	4.155	5081
8.0%	282	93	110	PUT	/api/me/profile-image	922.402	3.271	5.136	5.511	17
1.7%	78	0	24	GET	^/api/memories/[^/]+	191.387	2.454	4.095	4.095	521
0.9%	42	0	13	GET	/api/me/export	99.280	2.364	5.002	5.002	1437043
0.4%	1129	607	0	POST	/api/auth/login	46.556	0.041	0.220	0.261	127
0.3%	879	17	0	POST	/api/friend-requests	39.942	0.045	0.192	0.349	113
0.3%	1843	188	0	GET	^/api/images/[^/]+	39.057	0.021	0.118	0.217	168697
0.3%	381	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	38.601	0.101	0.361	0.582	116
0.3%	1631	0	0	GET	^/api/feeds/[^/]+	30.479	0.019	0.076	0.128	858
0.1%	171	1	0	POST	/api/auth/register	13.538	0.079	0.284	0.303	228
0.1%	747	0	0	GET	/api/posting-window/current	13.087	0.018	0.079	0.177	237
0.1%	1	0	0	POST	/api/initialize	12.687	12.687	12.687	12.687	90
0.1%	56	0	0	POST	^/api/posts/[^/]+/realmojis$	8.134	0.145	0.674	0.674	503
0.1%	418	0	0	GET	/api/users	7.919	0.019	0.077	0.138	234
0.1%	88	1	0	POST	^/api/posts/[^/]+/comments$	7.243	0.082	0.882	0.882	340
0.1%	350	66	0	GET	^/api/posts/me/current$	6.751	0.019	0.079	0.127	527
0.0%	26	0	0	DELETE	^/api/posts/[^/]+$	4.808	0.185	0.294	0.294	0
0.0%	356	0	0	GET	^/api/relationships/[^/]+	4.445	0.012	0.068	0.079	1642
0.0%	70	0	0	PATCH	/api/me	3.495	0.050	0.202	0.202	271
0.0%	201	2	0	GET	^/api/posts/[^/]+$	3.281	0.016	0.066	0.073	600
0.0%	73	0	0	PUT	^/api/follows/[^/]+$	2.776	0.038	0.145	0.145	0
0.0%	68	0	0	POST	/api/auth/logout	2.242	0.033	0.143	0.143	0
0.0%	111	0	0	GET	^/api/users/[^/]+$	1.722	0.016	0.054	0.061	178
0.0%	45	7	0	PATCH	^/api/posts/[^/]+$	1.581	0.035	0.101	0.101	550
0.0%	18	0	0	POST	/api/admin/posting-window	1.430	0.079	0.214	0.214	177
0.0%	85	2	0	GET	/api/me	1.391	0.016	0.097	0.097	292
0.0%	72	0	0	GET	^/api/posts/[^/]+/comments$	1.287	0.018	0.071	0.071	1041
0.0%	96	0	0	GET	^/api/posts/[^/]+/realmojis$	1.133	0.012	0.061	0.061	383
…（上位 30 / 全 38 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.4%	170	1	0	GET	/api/notifications	1425.285	8.384	12.059	12.060	179
25.7%	149	0	21	POST	/api/posts	929.705	6.240	9.384	9.497	546
18.1%	281	0	74	GET	/api/memories	652.564	2.322	4.245	4.592	5036
5.9%	56	0	35	PUT	/api/me/profile-image	212.771	3.799	4.269	4.269	11
2.1%	370	7	0	GET	^/api/images/[^/]+	76.347	0.206	0.956	1.290	205711
2.0%	24	0	12	GET	^/api/memories/[^/]+	73.540	3.064	4.230	4.230	871
1.0%	245	6	0	POST	/api/friend-requests	36.163	0.148	0.333	0.420	113
1.0%	8	0	2	GET	/api/me/export	35.777	4.472	5.924	5.924	1942374
0.9%	282	0	0	GET	^/api/feeds/[^/]+	33.890	0.120	0.212	0.228	525
0.8%	210	0	0	GET	/api/posting-window/current	27.859	0.133	0.251	0.370	238
0.6%	112	0	0	POST	/api/auth/login	21.319	0.190	0.333	0.356	245
0.4%	71	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	15.602	0.220	0.445	0.445	116
0.4%	114	0	0	GET	/api/users	14.501	0.127	0.231	0.266	240
0.3%	98	36	0	GET	^/api/posts/me/current$	11.791	0.120	0.296	0.296	410
0.3%	93	0	0	GET	^/api/relationships/[^/]+	10.923	0.117	0.203	0.203	1155
0.3%	44	0	0	POST	/api/auth/register	9.376	0.213	0.420	0.420	227
0.1%	33	0	0	PUT	^/api/follows/[^/]+$	4.717	0.143	0.239	0.239	0
0.1%	26	2	0	GET	/api/me	3.003	0.116	0.202	0.202	269
0.1%	15	0	0	PATCH	/api/me	2.500	0.167	0.321	0.321	276
0.1%	7	0	0	DELETE	^/api/posts/[^/]+$	2.181	0.312	0.407	0.407	0
0.1%	19	0	0	GET	^/api/posts/[^/]+$	2.105	0.111	0.177	0.177	618
0.1%	16	0	0	GET	^/api/posts/[^/]+/comments$	2.011	0.126	0.203	0.203	334
0.1%	16	0	0	POST	^/api/posts/[^/]+/comments$	1.907	0.119	0.217	0.217	276
0.0%	13	5	0	PATCH	^/api/posts/[^/]+$	1.794	0.138	0.191	0.191	397
0.0%	9	0	0	POST	/api/auth/logout	1.434	0.159	0.203	0.203	0
0.0%	5	0	0	POST	^/api/posts/[^/]+/realmojis$	1.365	0.273	0.476	0.476	407
0.0%	11	0	0	GET	^/api/users/[^/]+$	1.205	0.110	0.159	0.159	194
0.0%	1	0	0	POST	/api/me/account-deletion	0.525	0.525	0.525	0.525	0
0.0%	3	0	0	DELETE	^/api/comments/[^/]+$	0.486	0.162	0.206	0.206	0
0.0%	3	0	0	GET	^/api/posts/[^/]+/realmojis$	0.360	0.120	0.124	0.124	299
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.9%	80	0	0	GET	/api/notifications	654.217	8.178	9.069	9.069	180
25.1%	62	0	5	POST	/api/posts	382.725	6.173	9.226	9.226	595
18.3%	131	0	29	GET	/api/memories	279.774	2.136	4.264	4.395	6419
3.0%	241	2	0	GET	^/api/images/[^/]+	45.229	0.188	0.951	1.245	240063
1.8%	226	0	0	GET	^/api/feeds/[^/]+	27.200	0.120	0.208	0.214	1360
1.6%	11	0	2	PUT	/api/me/profile-image	24.573	2.234	4.206	4.206	3
1.1%	112	1	0	POST	/api/friend-requests	16.337	0.146	0.319	0.327	114
0.9%	4	0	1	GET	/api/me/export	13.572	3.393	5.770	5.770	2289134
0.9%	7	0	0	GET	^/api/memories/[^/]+	13.112	1.873	3.950	3.950	548
0.8%	83	0	0	GET	/api/posting-window/current	11.503	0.139	0.274	0.274	240
0.6%	55	0	0	POST	/api/auth/login	9.835	0.179	0.388	0.388	244
0.6%	39	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	8.800	0.226	0.409	0.409	116
0.4%	50	0	0	GET	/api/users	6.099	0.122	0.246	0.246	225
0.4%	48	1	0	GET	^/api/posts/[^/]+$	5.819	0.121	0.227	0.227	605
0.4%	47	13	0	GET	^/api/posts/me/current$	5.526	0.118	0.150	0.150	475
0.3%	46	0	0	GET	^/api/relationships/[^/]+	5.233	0.114	0.213	0.213	1516
0.2%	17	0	0	POST	/api/auth/register	3.381	0.199	0.275	0.275	228
0.2%	26	0	0	GET	^/api/posts/[^/]+/realmojis$	3.119	0.120	0.144	0.144	196
0.1%	13	0	0	GET	^/api/users/[^/]+$	1.687	0.130	0.154	0.154	170
0.1%	12	0	0	GET	^/api/posts/[^/]+/comments$	1.541	0.128	0.172	0.172	737
0.1%	12	0	0	GET	/api/me	1.322	0.110	0.154	0.154	296
0.1%	9	0	0	PUT	^/api/follows/[^/]+$	1.319	0.147	0.252	0.252	0
0.1%	7	0	0	PATCH	/api/me	1.271	0.182	0.275	0.275	241
0.1%	3	0	0	DELETE	^/api/posts/[^/]+$	0.782	0.261	0.320	0.320	0
0.0%	5	1	0	PATCH	^/api/posts/[^/]+$	0.755	0.151	0.193	0.193	528
0.0%	6	0	0	POST	/api/auth/logout	0.705	0.118	0.155	0.155	0
0.0%	3	0	0	POST	^/api/posts/[^/]+/realmojis$	0.696	0.232	0.270	0.270	519
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.104	0.104	0.104	0.104	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 12861 クエリ / 79 種類 / 合計 34.358s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	34.7%	38	11.937	314.13	784.66	6929.8	142.6	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	18.3%	796	6.297	7.91	36.25	713.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	14.6%	523	5.016	9.59	210.13	122.8	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	7.1%	1	2.439	2438.78	2438.78	29389.0	1199382.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
5	4.9%	1000	1.684	1.68	5.65	13.3	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
6	4.4%	1	1.507	1506.73	1506.73	19918.0	625617.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	3.1%	1	1.070	1070.30	1070.30	11528.0	529667.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	2.6%	623	0.893	1.43	3.91	26.8	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
9	2.0%	1	0.680	679.97	679.97	20931.0	506.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	1.1%	2	0.380	189.98	214.16	214958.5	2528.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.8%	2	0.263	131.37	254.14	16764.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.7%	1	0.245	245.00	245.00	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
13	0.6%	1	0.206	206.37	206.37	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.6%	1201	0.190	0.16	6.13	20.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
15	0.5%	64	0.188	2.94	5.75	39.7	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: ef4ec007692a195ce1a86c24d88df57725b07656
Type: cpu
Time: 2026-10-31 11:37:37 JST
Duration: 60.09s, Total samples = 70.38s (117.13%)
Showing nodes accounting for 62.03s, 88.14% of 70.38s total
Dropped 750 nodes (cum <= 0.35s)
Showing top 40 nodes out of 173
      flat  flat%   sum%        cum   cum%
    10.29s 14.62% 14.62%     10.29s 14.62%  runtime.spanClass.sizeclass (inline)
     7.76s 11.03% 25.65%     12.86s 18.27%  image/jpeg.(*decoder).reconstructBlock
     6.88s  9.78% 35.42%      6.89s  9.79%  crypto/internal/fips140/sha256.blockAVX2
     4.07s  5.78% 41.20%     28.08s 39.90%  image/jpeg.(*decoder).processSOS
     3.44s  4.89% 46.09%     14.40s 20.46%  runtime.tryDeferToSpanScan
     3.37s  4.79% 50.88%      5.50s  7.81%  image/jpeg.(*decoder).decodeHuffman
     2.58s  3.67% 54.55%      2.70s  3.84%  image/jpeg.(*decoder).receiveExtend
        2s  2.84% 57.39%      2.76s  3.92%  image/jpeg.idctCols
     1.72s  2.44% 59.83%      2.23s  3.17%  image/jpeg.(*decoder).ensureNBits
     1.71s  2.43% 62.26%      2.29s  3.25%  image/jpeg.idctRows
     1.71s  2.43% 64.69%      1.71s  2.43%  runtime.memclrNoHeapPointers
     1.71s  2.43% 67.12%      1.71s  2.43%  runtime.memmove
     1.63s  2.32% 69.44%      6.68s  9.49%  runtime.scanObjectsSmall
     1.45s  2.06% 71.50%      1.45s  2.06%  internal/runtime/syscall/linux.Syscall6
     1.33s  1.89% 73.39%      1.33s  1.89%  image/jpeg.dctBox (inline)
     1.25s  1.78% 75.16%      5.41s  7.69%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.22s  1.73% 76.90%     11.28s 16.03%  runtime.scanObject
     1.13s  1.61% 78.50%      1.13s  1.61%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.84s  1.19% 79.70%      0.84s  1.19%  memeqbody
     0.73s  1.04% 80.73%      0.73s  1.04%  aeshashbody
     0.58s  0.82% 81.56%      7.78s 11.05%  runtime.scanSpan
     0.52s  0.74% 82.30%      7.89s 11.21%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.48s  0.68% 82.98%      0.51s  0.72%  image/jpeg.(*decoder).readByteStuffedByte
     0.48s  0.68% 83.66%      0.48s  0.68%  runtime.extractHeapBitsSmall
     0.41s  0.58% 84.24%      0.41s  0.58%  example.com/isureal-company/isureal.(*UserRow).active
     0.34s  0.48% 84.73%      7.60s 10.80%  crypto/internal/fips140/sha256.(*Digest).Write
     0.32s  0.45% 85.18%      2.07s  2.94%  runtime.mapaccess1_faststr
     0.31s  0.44% 85.62%      0.53s  0.75%  example.com/isureal-company/isureal.postLess
     0.25s  0.36% 85.98%      8.82s 12.53%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.23s  0.33% 86.30%      8.25s 11.72%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.18s  0.26% 86.56%      0.44s  0.63%  runtime.scanObjectSmall
     0.17s  0.24% 86.80%      0.42s   0.6%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.16s  0.23% 87.03%      0.38s  0.54%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.15s  0.21% 87.24%      0.90s  1.28%  runtime.mallocgcSmallScanNoHeader
     0.14s   0.2% 87.44%      9.59s 13.63%  example.com/isureal-company/isureal.imageSignature
     0.13s  0.18% 87.62%     10.82s 15.37%  runtime.mallocgc
     0.11s  0.16% 87.78%      1.13s  1.61%  runtime.mapassign_faststr
     0.09s  0.13% 87.91%      2.29s  3.25%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.08s  0.11% 88.02%      6.97s  9.90%  crypto/internal/fips140/sha256.block
     0.08s  0.11% 88.14%      1.13s  1.61%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 69.8% / 100% | 49% | 0.2% | 6 | 203 MB | CPU が飽和 |
| isu2 (2 cpu) | 1.6% / 43% | 0% | 0.6% | 2 | 494 MB |  |
| isu3 (2 cpu) | 1.1% / 4% | 0% | 0.1% | 2 | 262 MB |  |
| isu4 (2 cpu) | 20.5% / 52% | 0% | 10.8% | 2 | 109 MB |  |

**isu1（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 122% / nginx 12% |
| isu2 | 目立つプロセス無し |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 17% |

## Diff from previous run

スコア: 0 → 21198 (+21198)

`1376dd7..d1c842b`

```
 webapp/go/limiter.go  | 36 ++++++++++++++++++++++++++++++++++++
 webapp/go/server.go   | 10 +++++-----
 webapp/go/user_api.go | 17 +++--------------
 3 files changed, 44 insertions(+), 19 deletions(-)
```
