# run 20261031-131613-e2387b9

## Summary

| | |
|---|---|
| score | 83453 |
| status | pass |
| ref | repl-replica |
| sha | e2387b97be30 |
| 所要時間 | 962s |
| 前回比 | +18753 (29.0%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +18753 (29.0%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | 複製 検証: shadow（第3段の経路も比べる）＋要約の照合 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	35	81	204	508	507
europe	7	17	39	87	87	86
america	6	14	29	50	51	50
```

```
{"id":"12893","status":"done","portalStatus":"finished","passed":true,"score":83453,"raw":83453,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:30:45.628Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:30:47.927Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:30:51.028Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:30:54.076Z level=INFO msg="asia 2 日目: 35 人"
time=2026-10-31T04:30:57.027Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T04:30:59.058Z level=INFO msg="user moved" user=tormax5260 from=asia to=america
time=2026-10-31T04:31:00.034Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T04:31:03.078Z level=INFO msg="asia 3 日目: 81 人"
time=2026-10-31T04:31:06.775Z level=INFO msg="europe 3 日目: 39 人"
time=2026-10-31T04:31:09.032Z level=INFO msg="america 3 日目: 29 人"
time=2026-10-31T04:31:12.089Z level=INFO msg="asia 4 日目: 204 人"
time=2026-10-31T04:31:16.954Z level=INFO msg="europe 4 日目: 87 人"
time=2026-10-31T04:31:18.148Z level=INFO msg="america 4 日目: 50 人"
time=2026-10-31T04:31:19.528Z level=INFO msg="user moved" user=mimano6609 from=asia to=america
time=2026-10-31T04:31:21.104Z level=INFO msg="asia 5 日目: 508 人"
time=2026-10-31T04:31:24.200Z level=INFO msg="europe 5 日目: 87 人"
time=2026-10-31T04:31:27.176Z level=INFO msg="america 5 日目: 51 人"
time=2026-10-31T04:31:36.887Z level=INFO msg="asia 最終日（グランプリ）: 507 人"
time=2026-10-31T04:31:37.099Z level=INFO msg="europe 最終日（グランプリ）: 86 人"
time=2026-10-31T04:31:37.289Z level=INFO msg="america 最終日（グランプリ）: 50 人"
time=2026-10-31T04:31:44.947Z level=INFO msg="グランプリ: アジア 463 人, ヨーロッパ 56 人, アメリカ 37 人"
{
  "pass": true,
  "score": 83453,
  "messages": [
    "isyuan.5122: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.1%	885	1	0	GET	/api/notifications	7232.456	8.172	15.012	15.016	168
13.1%	1333	4	0	POST	/api/posts	2485.998	1.865	7.001	9.388	653
10.5%	3192	94	0	POST	/api/friend-requests	1981.914	0.621	3.196	6.917	113
7.9%	3826	1479	0	POST	/api/auth/login	1505.192	0.393	3.014	5.295	161
6.1%	8510	334	0	GET	^/api/images/[^/]+	1160.306	0.136	1.638	5.728	167787
5.1%	1379	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	971.795	0.705	3.574	7.303	116
2.3%	409	20	0	POST	^/api/posts/[^/]+/realmojis$	439.621	1.075	5.174	6.819	470
2.0%	432	116	0	PUT	/api/me/profile-image	385.563	0.893	3.743	4.247	6
2.0%	526	1	0	POST	/api/auth/register	381.312	0.725	3.129	3.687	227
1.9%	9478	0	0	GET	^/api/feeds/[^/]+	366.686	0.039	0.838	1.424	2581
1.7%	466	17	0	POST	^/api/posts/[^/]+/comments$	328.827	0.706	4.027	6.022	333
1.5%	2360	0	0	GET	/api/memories	278.980	0.118	0.842	1.526	9590
1.0%	395	0	0	POST	/api/auth/logout	194.950	0.494	3.058	4.315	0
1.0%	334	0	0	PUT	^/api/follows/[^/]+$	186.255	0.558	3.025	3.159	0
0.9%	244	0	0	PATCH	/api/me	177.336	0.727	3.037	3.376	283
0.8%	150	1	0	PATCH	^/api/posts/[^/]+$	155.898	1.039	5.793	6.169	642
0.6%	159	0	0	GET	/api/me/export	120.075	0.755	2.571	3.869	1792811
0.4%	1290	4	0	GET	^/api/posts/me/current$	77.267	0.060	0.992	1.424	653
0.3%	57	0	0	DELETE	^/api/posts/[^/]+$	64.987	1.140	5.995	5.995	0
0.3%	1457	0	0	GET	/api/users	63.588	0.044	0.789	1.422	241
0.3%	1888	0	0	GET	/api/posting-window/current	59.491	0.032	0.477	1.037	237
0.3%	1380	15	0	GET	^/api/posts/[^/]+$	47.933	0.035	0.900	1.422	605
0.2%	86	0	0	DELETE	^/api/comments/[^/]+$	41.090	0.478	2.257	2.257	0
0.2%	1242	0	0	GET	^/api/relationships/[^/]+	40.361	0.032	0.411	1.422	1095
0.2%	57	0	0	DELETE	^/api/realmojis/[^/]+$	37.278	0.654	3.589	3.589	0
0.2%	328	0	0	GET	^/api/memories/[^/]+	31.228	0.095	0.856	0.955	738
0.1%	1	0	0	POST	/api/initialize	21.940	21.940	21.940	21.940	90
0.1%	562	0	0	GET	^/api/posts/[^/]+/comments$	21.693	0.039	0.897	1.422	1603
0.1%	327	3	0	GET	/api/me	20.185	0.062	0.990	1.119	298
0.1%	711	0	0	GET	^/api/users/[^/]+$	18.467	0.026	0.411	1.058	189
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.6%	235	1	0	GET	/api/notifications	1989.084	8.464	12.859	13.017	177
15.9%	292	2	0	POST	/api/posts	778.044	2.665	8.391	9.133	650
6.0%	1446	12	0	GET	^/api/images/[^/]+	293.428	0.203	1.136	2.891	186026
5.9%	543	10	0	POST	/api/friend-requests	291.192	0.536	3.056	3.473	113
5.1%	1252	0	0	GET	^/api/feeds/[^/]+	248.324	0.198	1.736	2.369	3141
4.9%	266	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	242.059	0.910	3.574	5.908	116
4.7%	370	0	0	POST	/api/auth/login	231.081	0.625	3.991	5.285	250
3.5%	427	0	0	GET	/api/memories	171.893	0.403	2.286	2.291	12217
2.2%	386	0	0	GET	/api/posting-window/current	109.221	0.283	1.774	2.073	239
1.2%	55	5	0	POST	^/api/posts/[^/]+/realmojis$	57.667	1.048	4.263	4.263	409
1.0%	223	14	0	GET	^/api/posts/[^/]+$	49.644	0.223	1.856	2.290	580
0.9%	248	0	0	GET	/api/users	45.593	0.184	1.563	1.874	236
0.9%	249	12	0	GET	^/api/posts/me/current$	45.512	0.183	0.406	2.215	625
0.8%	81	0	0	POST	^/api/posts/[^/]+/comments$	41.417	0.511	2.800	2.800	357
0.8%	185	0	0	GET	^/api/relationships/[^/]+	41.283	0.223	1.858	2.286	1217
0.7%	61	0	0	POST	/api/auth/logout	35.682	0.585	2.621	2.621	0
0.7%	21	0	0	GET	/api/me/export	33.510	1.596	3.988	3.988	2321659
0.5%	85	0	0	POST	/api/auth/register	24.905	0.293	0.583	0.583	228
0.4%	55	0	0	PUT	^/api/follows/[^/]+$	20.683	0.376	1.818	1.818	0
0.4%	95	0	0	GET	^/api/users/[^/]+$	19.257	0.203	2.288	2.288	196
0.4%	80	0	0	GET	^/api/posts/[^/]+/realmojis$	18.538	0.232	1.731	1.731	946
0.4%	46	0	0	PUT	/api/me/profile-image	18.537	0.403	0.855	0.855	0
0.4%	61	0	0	GET	^/api/memories/[^/]+	18.193	0.298	2.213	2.213	695
0.3%	80	0	0	GET	^/api/posts/[^/]+/comments$	16.559	0.207	1.734	1.734	1288
0.3%	42	0	0	PATCH	/api/me	15.535	0.370	1.042	1.042	277
0.3%	22	3	0	PATCH	^/api/posts/[^/]+$	15.481	0.704	3.491	3.491	573
0.3%	16	0	0	DELETE	^/api/posts/[^/]+$	13.343	0.834	3.705	3.705	0
0.2%	51	2	0	GET	/api/me	7.531	0.148	0.251	0.251	287
0.1%	8	0	0	DELETE	^/api/comments/[^/]+$	2.972	0.371	0.698	0.698	0
0.1%	1	0	0	POST	/api/me/account-deletion	2.707	2.707	2.707	2.707	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.4%	149	1	0	GET	/api/notifications	1204.458	8.084	10.270	10.270	179
15.3%	177	0	0	POST	/api/posts	457.970	2.587	8.575	8.583	636
6.0%	870	6	0	GET	^/api/images/[^/]+	179.148	0.206	1.221	1.662	187779
5.3%	282	2	0	POST	/api/friend-requests	156.774	0.556	3.402	3.510	114
5.1%	652	0	0	GET	^/api/feeds/[^/]+	153.309	0.235	2.034	2.490	2638
4.7%	221	0	0	POST	/api/auth/login	139.660	0.632	3.806	5.604	251
4.1%	150	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	121.958	0.813	3.143	3.189	116
3.1%	274	0	0	GET	/api/memories	93.139	0.340	1.992	2.465	13861
2.3%	71	3	0	POST	^/api/posts/[^/]+/comments$	68.658	0.967	4.371	4.371	299
1.8%	229	0	0	GET	/api/posting-window/current	53.896	0.235	0.996	1.830	239
1.8%	166	6	0	GET	^/api/posts/me/current$	53.868	0.325	2.499	2.512	613
1.5%	42	4	0	POST	^/api/posts/[^/]+/realmojis$	43.446	1.034	4.437	4.437	417
1.0%	37	0	0	POST	/api/auth/logout	29.822	0.806	4.310	4.310	0
1.0%	141	0	0	GET	/api/users	29.751	0.211	1.973	1.987	234
0.9%	99	0	0	GET	^/api/relationships/[^/]+	26.537	0.268	2.519	2.519	1295
0.9%	13	0	0	GET	/api/me/export	25.869	1.990	5.904	5.904	2234085
0.7%	41	0	0	PUT	^/api/follows/[^/]+$	21.093	0.514	4.002	4.002	0
0.7%	91	5	0	GET	^/api/posts/[^/]+$	19.930	0.219	1.995	1.995	568
0.5%	55	0	0	GET	^/api/users/[^/]+$	13.954	0.254	2.154	2.154	185
0.5%	48	0	0	POST	/api/auth/register	13.813	0.288	0.563	0.563	229
0.4%	40	0	0	GET	^/api/posts/[^/]+/comments$	12.301	0.308	2.071	2.071	1701
0.4%	37	0	0	GET	^/api/memories/[^/]+	11.526	0.312	2.164	2.164	1011
0.4%	22	0	0	PUT	/api/me/profile-image	11.370	0.517	1.970	1.970	0
0.4%	26	0	0	PATCH	/api/me	11.121	0.428	1.981	1.981	269
0.3%	44	0	0	GET	^/api/posts/[^/]+/realmojis$	8.285	0.188	1.821	1.821	867
0.2%	9	0	0	DELETE	^/api/comments/[^/]+$	6.042	0.671	2.577	2.577	0
0.2%	9	0	0	PATCH	^/api/posts/[^/]+$	4.884	0.543	1.185	1.185	621
0.1%	6	0	0	DELETE	^/api/posts/[^/]+$	3.923	0.654	0.881	0.881	0
0.1%	26	1	0	GET	/api/me	3.878	0.149	0.228	0.228	282
0.1%	4	0	0	DELETE	^/api/realmojis/[^/]+$	1.597	0.399	0.475	0.475	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 45930 クエリ / 81 種類 / 合計 313.543s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	47.1%	2661	147.767	55.53	443.86	621.8	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	21.0%	193	65.850	341.19	1349.26	5914.7	120.6	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	6.8%	3	21.444	7148.05	9536.94	30636.0	1200038.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
4	5.6%	1924	17.588	9.14	346.78	119.5	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	5.1%	3	16.064	5354.69	7110.45	23238.0	627924.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
6	4.4%	5139	13.819	2.69	29.01	11.2	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
7	2.5%	6	7.821	1303.57	6960.16	219789.2	2596.3	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
8	2.5%	3	7.718	2572.80	3439.74	12310.0	530201.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	1.4%	1800	4.501	2.50	29.03	25.7	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.6%	6	1.879	313.14	1265.04	16785.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.5%	477	1.659	3.48	16.95	30.6	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
12	0.4%	1	1.097	1097.42	1097.42	2646.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
13	0.3%	598	0.985	1.65	9.25	21.7	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
14	0.3%	6	0.792	132.05	214.05	623.7	25369.8	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
15	0.2%	3882	0.763	0.20	17.26	20.1	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
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
| isu1 (2 cpu) | 42.3% / 100% | 19% | 1.1% | 7 | 117 MB | run queue が積もっている |
| isu2 (2 cpu) | 23.0% / 78% | 0% | 0.5% | 3 | 420 MB |  |
| isu3 (2 cpu) | 22.6% / 97% | 2% | 0.3% | 3 | 860 MB |  |
| isu4 (2 cpu) | 44.5% / 100% | 15% | 17.3% | 24 | 149 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 58% / nginx 11% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 44% |
| isu3 | isureal 43% |
| isu4 | postgres 49% |

## Diff from previous run

スコア: 64700 → 83453 (+18753)

`e0e4cb3..e2387b9`

```
 conf/isu1/home/isucon/env.sh                   |   1 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |   7 +-
 conf/isu2/home/isucon/env.sh                   |   3 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |   7 +-
 conf/isu3/home/isucon/env.sh                   |   3 +
 webapp/go/account_deletion_api.go              |  10 +-
 webapp/go/admin_posting_window_api.go          |  16 +-
 webapp/go/admin_user_api.go                    |   5 +-
 webapp/go/cmd/isureal/main.go                  |   6 +
 webapp/go/comment_api.go                       |   9 +-
 webapp/go/export_api.go                        |   2 +
 webapp/go/feed_api.go                          | 113 +++-
 webapp/go/feed_test.go                         | 331 ++++++++++
 webapp/go/friend_request_api.go                |  15 +-
 webapp/go/image_local.go                       |   1 +
 webapp/go/initialize.go                        |   4 +-
 webapp/go/offload.go                           |  27 +-
 webapp/go/post_api.go                          |  76 ++-
 webapp/go/profile_image_api.go                 |  13 +-
 webapp/go/realmoji_api.go                      |  16 +-
 webapp/go/relationship.go                      |  10 +-
 webapp/go/relationship_api.go                  |  33 +-
 webapp/go/repl.go                              | 870 +++++++++++++++++++++++++
 webapp/go/repl_digest.go                       | 151 +++++
 webapp/go/repl_test.go                         |  62 ++
 webapp/go/replica_server.go                    | 234 +++++++
 webapp/go/server.go                            |   1 +
 webapp/go/store.go                             | 145 ++++-
 webapp/go/user_api.go                          |  34 +-
 29 files changed, 2027 insertions(+), 178 deletions(-)
```
