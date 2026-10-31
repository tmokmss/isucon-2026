# run 20261031-133224-e43e4ce

## Summary

| | |
|---|---|
| score | 104467 |
| status | pass |
| ref | repl-replica |
| sha | e43e4ce69db8 |
| 所要時間 | 441s |
| 前回比 | +21014 (25.2%) vs 20261031-131613-e2387b9 |
| ベスト比 | +21014 (25.2%) vs 20261031-131613-e2387b9 = 83453 **ベスト更新** |
| note | 複製 local・書き込みは返事を待たず読み取り側が時刻の印で待つ（main 860a676 を取り込み済み） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	136	392	990	990
europe	7	16	42	115	115	115
america	6	14	40	91	91	91
```

```
{"id":"13455","status":"done","portalStatus":"finished","passed":true,"score":104467,"raw":104467,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:38:07.558Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:38:10.523Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:38:13.522Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:38:16.554Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T04:38:19.521Z level=INFO msg="europe 2 日目: 16 人"
time=2026-10-31T04:38:22.521Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T04:38:25.543Z level=INFO msg="asia 3 日目: 136 人"
time=2026-10-31T04:38:28.529Z level=INFO msg="europe 3 日目: 42 人"
time=2026-10-31T04:38:28.694Z level=INFO msg="user moved" user=roneta5915 from=asia to=america
time=2026-10-31T04:38:31.522Z level=INFO msg="america 3 日目: 40 人"
time=2026-10-31T04:38:34.551Z level=INFO msg="asia 4 日目: 392 人"
time=2026-10-31T04:38:37.654Z level=INFO msg="europe 4 日目: 115 人"
time=2026-10-31T04:38:39.456Z level=INFO msg="user moved" user=tatsueva_7021 from=asia to=europe
time=2026-10-31T04:38:40.568Z level=INFO msg="america 4 日目: 91 人"
time=2026-10-31T04:38:43.515Z level=INFO msg="user moved" user=ivysu8416 from=europe to=asia
time=2026-10-31T04:38:43.747Z level=INFO msg="asia 5 日目: 990 人"
time=2026-10-31T04:38:46.776Z level=INFO msg="europe 5 日目: 115 人"
time=2026-10-31T04:38:49.831Z level=INFO msg="america 5 日目: 91 人"
time=2026-10-31T04:38:59.000Z level=INFO msg="europe 最終日（グランプリ）: 115 人"
time=2026-10-31T04:38:59.495Z level=INFO msg="america 最終日（グランプリ）: 91 人"
time=2026-10-31T04:38:59.734Z level=INFO msg="asia 最終日（グランプリ）: 990 人"
time=2026-10-31T04:39:07.550Z level=INFO msg="グランプリ: アジア 247 人, ヨーロッパ 38 人, アメリカ 22 人"
{
  "pass": true,
  "score": 104467,
  "messages": [
    "kasoel.5122: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
29.9%	1551	0	0	GET	/api/notifications	12797.137	8.251	16.024	16.250	162
21.0%	2274	83	0	POST	/api/posts	8986.936	3.952	13.888	17.628	631
12.7%	5194	75	0	POST	/api/friend-requests	5442.942	1.048	10.545	18.315	114
9.5%	13435	735	0	GET	^/api/images/[^/]+	4069.904	0.303	6.353	15.896	131898
6.8%	5315	1133	0	POST	/api/auth/login	2896.606	0.545	6.800	15.960	197
4.1%	1622	9	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1766.891	1.089	11.158	15.469	116
2.9%	1007	1	0	POST	/api/auth/register	1233.237	1.225	6.828	12.076	227
2.3%	701	126	0	PUT	/api/me/profile-image	990.210	1.413	7.796	9.865	4
1.7%	538	18	0	POST	^/api/posts/[^/]+/realmojis$	732.038	1.361	14.265	16.200	464
1.4%	647	0	0	PUT	^/api/follows/[^/]+$	616.924	0.954	6.094	16.235	0
1.2%	491	0	0	PATCH	/api/me	514.666	1.048	7.156	15.764	278
1.1%	3229	0	0	GET	/api/memories	469.035	0.145	1.021	1.879	11325
1.1%	482	10	0	POST	^/api/posts/[^/]+/comments$	456.012	0.946	11.149	17.219	332
0.7%	165	11	0	PATCH	^/api/posts/[^/]+$	285.405	1.730	10.883	11.242	612
0.7%	650	0	0	POST	/api/auth/logout	279.153	0.429	6.682	16.015	0
0.6%	223	0	0	GET	/api/me/export	255.345	1.145	9.966	13.320	1963036
0.6%	10648	0	0	GET	^/api/feeds/[^/]+	250.450	0.024	0.206	1.299	2742
0.5%	79	0	0	DELETE	^/api/posts/[^/]+$	197.632	2.502	11.820	11.820	0
0.2%	3321	0	0	GET	/api/posting-window/current	105.548	0.032	0.392	0.944	236
0.2%	1661	95	0	GET	^/api/posts/me/current$	77.517	0.047	0.496	9.738	618
0.1%	2479	0	0	GET	/api/users	62.705	0.025	0.198	0.694	239
0.1%	84	0	0	DELETE	^/api/realmojis/[^/]+$	57.230	0.681	5.465	5.465	0
0.1%	79	0	0	DELETE	^/api/comments/[^/]+$	40.174	0.509	3.688	3.688	0
0.1%	630	0	0	GET	^/api/memories/[^/]+	34.967	0.056	0.384	0.906	607
0.1%	1454	0	0	GET	^/api/relationships/[^/]+	33.129	0.023	0.221	0.879	951
0.1%	36	0	0	PUT	^/api/blocks/[^/]+$	32.441	0.901	9.139	9.139	0
0.1%	1797	13	0	GET	^/api/posts/[^/]+$	30.848	0.017	0.124	1.169	622
0.0%	821	0	0	GET	^/api/users/[^/]+$	15.692	0.019	0.135	1.425	200
0.0%	583	2	0	GET	/api/me	14.011	0.024	0.144	0.264	302
0.0%	1	0	0	POST	/api/initialize	11.649	11.649	11.649	11.649	90
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.0%	300	0	0	GET	/api/notifications	2547.882	8.493	11.810	13.176	178
20.1%	361	6	0	POST	/api/posts	1346.546	3.730	15.226	16.685	647
6.6%	618	8	0	POST	/api/friend-requests	442.301	0.716	8.324	11.494	114
6.2%	505	0	0	POST	/api/auth/login	414.984	0.822	9.030	15.219	248
6.0%	1798	15	0	GET	^/api/images/[^/]+	404.500	0.225	1.137	2.929	169509
4.8%	229	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	318.436	1.391	9.680	15.371	116
4.5%	1483	0	0	GET	^/api/feeds/[^/]+	299.501	0.202	1.131	2.789	3186
1.7%	476	0	0	GET	/api/memories	112.181	0.236	1.746	1.850	12758
1.5%	502	0	0	GET	/api/posting-window/current	100.953	0.201	1.094	1.266	239
1.4%	59	6	0	POST	^/api/posts/[^/]+/realmojis$	90.500	1.534	15.276	15.276	407
1.3%	82	0	0	POST	/api/auth/logout	90.283	1.101	13.920	13.920	0
1.0%	295	16	0	GET	^/api/posts/me/current$	67.081	0.227	1.138	1.211	621
0.8%	61	2	0	POST	^/api/posts/[^/]+/comments$	52.271	0.857	6.663	6.663	320
0.7%	113	0	0	POST	/api/auth/register	44.524	0.394	1.551	1.594	228
0.7%	279	0	0	GET	/api/users	43.597	0.156	0.966	1.280	245
0.6%	28	6	0	PATCH	^/api/posts/[^/]+$	41.753	1.491	9.332	9.332	511
0.6%	43	0	0	GET	/api/me/export	39.530	0.919	13.160	13.160	2354760
0.6%	261	7	0	GET	^/api/posts/[^/]+$	37.007	0.142	0.982	1.143	620
0.5%	67	0	0	PUT	/api/me/profile-image	35.898	0.536	1.191	1.191	0
0.5%	185	0	0	GET	^/api/relationships/[^/]+	35.024	0.189	0.984	1.004	1109
0.4%	81	0	0	PUT	^/api/follows/[^/]+$	30.150	0.372	2.670	2.670	0
0.3%	54	0	0	PATCH	/api/me	20.604	0.382	2.292	2.292	285
0.3%	124	0	0	GET	^/api/users/[^/]+$	17.901	0.144	0.824	0.898	209
0.2%	16	0	0	DELETE	^/api/posts/[^/]+$	16.751	1.047	2.628	2.628	0
0.2%	92	0	0	GET	^/api/posts/[^/]+/comments$	13.418	0.146	1.009	1.009	855
0.2%	69	0	0	GET	^/api/memories/[^/]+	12.296	0.178	1.782	1.782	740
0.2%	89	0	0	GET	^/api/posts/[^/]+/realmojis$	10.879	0.122	0.853	0.853	800
0.1%	71	1	0	GET	/api/me	9.118	0.128	0.579	0.579	299
0.1%	8	0	0	DELETE	^/api/realmojis/[^/]+$	4.011	0.501	1.533	1.533	0
0.0%	13	0	0	DELETE	^/api/comments/[^/]+$	2.186	0.168	0.452	0.452	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.0%	221	1	0	GET	/api/notifications	1688.032	7.638	9.753	9.754	171
18.7%	257	6	0	POST	/api/posts	928.149	3.611	13.986	16.224	641
9.5%	450	15	0	POST	/api/friend-requests	471.026	1.047	9.639	12.683	112
6.7%	1386	11	0	GET	^/api/images/[^/]+	332.682	0.240	1.364	2.843	142992
6.3%	178	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	313.580	1.762	14.173	14.362	116
6.3%	376	0	0	POST	/api/auth/login	312.563	0.831	9.279	13.714	251
3.2%	913	0	0	GET	^/api/feeds/[^/]+	159.297	0.174	0.983	1.310	2537
2.1%	366	0	0	GET	/api/memories	102.417	0.280	2.311	2.514	15261
1.6%	360	0	0	GET	/api/posting-window/current	80.742	0.224	1.133	1.585	240
1.4%	48	0	0	POST	/api/auth/logout	69.935	1.457	13.456	13.456	0
1.1%	31	2	0	POST	^/api/posts/[^/]+/realmojis$	54.136	1.746	15.305	15.305	431
1.1%	198	12	0	GET	^/api/posts/me/current$	52.415	0.265	1.063	4.912	615
1.0%	58	0	0	PUT	/api/me/profile-image	50.064	0.863	6.324	6.324	0
1.0%	63	0	0	PUT	^/api/follows/[^/]+$	47.418	0.753	6.179	6.179	0
0.9%	189	0	0	GET	/api/users	44.681	0.236	0.996	1.100	241
0.7%	21	0	0	GET	/api/me/export	32.741	1.559	8.511	8.511	3013588
0.6%	89	0	0	POST	/api/auth/register	28.650	0.322	2.297	2.297	229
0.5%	40	0	0	PATCH	/api/me	27.286	0.682	3.851	3.851	291
0.5%	12	0	0	DELETE	^/api/posts/[^/]+$	25.456	2.121	6.556	6.556	0
0.4%	133	0	0	GET	^/api/relationships/[^/]+	22.290	0.168	1.010	1.107	1268
0.4%	158	5	0	GET	^/api/posts/[^/]+$	19.788	0.125	0.813	1.003	592
0.3%	36	1	0	POST	^/api/posts/[^/]+/comments$	16.484	0.458	3.218	3.218	303
0.3%	19	4	0	PATCH	^/api/posts/[^/]+$	14.645	0.771	5.334	5.334	519
0.2%	63	1	0	GET	/api/me	11.819	0.188	1.604	1.604	294
0.2%	88	0	0	GET	^/api/users/[^/]+$	10.713	0.122	0.756	0.756	179
0.2%	6	0	0	DELETE	^/api/realmojis/[^/]+$	8.250	1.375	3.254	3.254	0
0.2%	65	0	0	GET	^/api/posts/[^/]+/realmojis$	7.448	0.115	0.729	0.729	526
0.1%	5	0	0	PUT	^/api/blocks/[^/]+$	7.154	1.431	5.665	5.665	0
0.1%	58	0	0	GET	^/api/posts/[^/]+/comments$	6.846	0.118	0.964	0.964	865
0.1%	44	0	0	GET	^/api/memories/[^/]+	6.270	0.143	0.965	0.965	920
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 69367 クエリ / 83 種類 / 合計 584.050s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	64.3%	4100	375.516	91.59	704.32	660.4	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	14.9%	240	86.866	361.94	1901.24	1276.6	129.7	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	5.0%	2742	29.318	10.69	1079.10	122.8	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.0%	8906	23.159	2.60	48.76	10.7	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	3.7%	3	21.441	7146.87	9739.65	30732.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	2.7%	3	16.018	5339.40	7184.61	23539.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	1.7%	2798	9.930	3.55	117.21	25.2	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
8	1.4%	3	7.915	2638.22	3706.60	12428.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	0.4%	602	2.162	3.59	22.48	28.6	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.3%	6144	1.474	0.24	31.45	19.9	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.2%	6	1.441	240.08	476.07	22305.3	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.2%	566	1.107	1.96	31.22	22.0	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
13	0.2%	6	1.075	179.23	219.06	209821.2	2492.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.1%	1	0.820	820.47	820.47	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	1981	0.784	0.40	25.66	20.4	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
未取得: プロファイルを取れませんでした (isu1)
  http://localhost:6060/debug/pprof/profile?seconds=60

curl: (22) The requested URL returned error: 500
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 51.1% / 100% | 32% | 1.1% | 10 | 127 MB | CPU が飽和 |
| isu2 (2 cpu) | 26.2% / 86% | 0% | 1.9% | 4 | 109 MB |  |
| isu3 (2 cpu) | 30.1% / 100% | 2% | 3.4% | 4 | 100 MB |  |
| isu4 (2 cpu) | 49.4% / 100% | 24% | 14.5% | 25 | 118 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 72% / nginx 12% / journalctl 3% / kswapd0 2% / systemd-journal 2% |
| isu2 | isureal 45% / kswapd0 1% |
| isu3 | isureal 46% / kswapd0 3% / du 2% / nginx 1% |
| isu4 | postgres 61% / isureal 3% |

## Diff from previous run

スコア: 83453 → 104467 (+21014)

`e2387b9..e43e4ce`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf     |  1 +
 conf/isu1/home/isucon/env.sh                       |  1 -
 conf/isu2/etc/nginx/sites-enabled/isureal.conf     |  1 +
 conf/isu2/home/isucon/env.sh                       |  2 +-
 conf/isu3/etc/nginx/sites-enabled/isureal.conf     |  1 +
 conf/isu3/home/isucon/env.sh                       |  2 +-
 .../etc/postgresql/16/main/conf.d/zz-tuning.conf   |  2 +
 webapp/go/config.go                                |  2 +
 webapp/go/export_api.go                            |  4 +-
 webapp/go/repl.go                                  | 63 ++++++++++++++++++----
 webapp/go/replica_server.go                        | 36 +++++++++----
 11 files changed, 92 insertions(+), 23 deletions(-)
```
