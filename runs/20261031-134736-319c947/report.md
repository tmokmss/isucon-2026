# run 20261031-134736-319c947

## Summary

| | |
|---|---|
| score | 104789 |
| status | pass |
| ref | repl-replica |
| sha | 319c9472f411 |
| 所要時間 | 434s |
| 前回比 | +40089 (62.0%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | +40089 (62.0%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | 複製 検証: 時刻の印で待つ形の shadow＋照合（main fbae0df 取り込み済み） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	143	417	910	908
europe	7	17	42	97	99	100
america	6	16	47	84	83	83
```

```
{"id":"14685","status":"done","portalStatus":"finished","passed":true,"score":104789,"raw":104789,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:53:15.487Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:53:18.451Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:53:21.449Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:53:24.260Z level=INFO msg="user moved" user=danhi.5127 from=america to=asia
time=2026-10-31T04:53:24.477Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T04:53:25.079Z level=INFO msg="user moved" user=danhi.5127 from=asia to=europe
time=2026-10-31T04:53:27.453Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T04:53:30.449Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T04:53:33.490Z level=INFO msg="asia 3 日目: 143 人"
time=2026-10-31T04:53:37.050Z level=INFO msg="europe 3 日目: 42 人"
time=2026-10-31T04:53:37.911Z level=INFO msg="user moved" user=fuvin_6108 from=asia to=america
time=2026-10-31T04:53:39.453Z level=INFO msg="america 3 日目: 47 人"
time=2026-10-31T04:53:39.491Z level=INFO msg="user moved" user=danhi.5127 from=europe to=asia
time=2026-10-31T04:53:42.357Z level=INFO msg="user moved" user=fulia.6286 from=europe to=asia
time=2026-10-31T04:53:42.493Z level=INFO msg="asia 4 日目: 417 人"
time=2026-10-31T04:53:45.567Z level=INFO msg="europe 4 日目: 97 人"
time=2026-10-31T04:53:48.180Z level=INFO msg="user moved" user=marnete-6845 from=asia to=america
time=2026-10-31T04:53:48.467Z level=INFO msg="user moved" user=anyuno-7471 from=asia to=europe
time=2026-10-31T04:53:48.477Z level=INFO msg="america 4 日目: 84 人"
time=2026-10-31T04:53:51.696Z level=INFO msg="asia 5 日目: 910 人"
time=2026-10-31T04:53:51.869Z level=INFO msg="user moved" user=mavinma-6782 from=asia to=europe
time=2026-10-31T04:53:54.564Z level=INFO msg="europe 5 日目: 99 人"
time=2026-10-31T04:53:58.027Z level=INFO msg="user moved" user=fuvin_6108 from=america to=europe
time=2026-10-31T04:53:58.047Z level=INFO msg="america 5 日目: 83 人"
time=2026-10-31T04:54:08.238Z level=INFO msg="asia 最終日（グランプリ）: 908 人"
time=2026-10-31T04:54:08.476Z level=INFO msg="america 最終日（グランプリ）: 83 人"
time=2026-10-31T04:54:08.544Z level=INFO msg="europe 最終日（グランプリ）: 100 人"
time=2026-10-31T04:54:15.487Z level=INFO msg="グランプリ: アジア 384 人, ヨーロッパ 14 人, アメリカ 8 人"
{
  "pass": true,
  "score": 104789,
  "messages": [
    "ruvin_5135: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "zoenu-5085: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.0%	1545	0	0	GET	/api/notifications	14096.389	9.124	16.562	16.569	166
17.0%	2352	76	0	POST	/api/posts	6325.042	2.689	8.070	9.957	633
10.2%	5123	77	0	POST	/api/friend-requests	3779.738	0.738	4.868	7.717	114
7.7%	13403	447	0	GET	^/api/images/[^/]+	2863.439	0.214	3.117	8.868	146744
6.1%	5889	1749	0	POST	/api/auth/login	2274.583	0.386	4.198	8.035	180
4.2%	1864	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1548.822	0.831	5.425	8.856	116
2.9%	658	122	0	PUT	/api/me/profile-image	1080.535	1.642	6.352	9.035	3
2.6%	638	36	0	POST	^/api/posts/[^/]+/realmojis$	970.694	1.521	6.176	7.852	445
1.6%	927	1	0	POST	/api/auth/register	595.172	0.642	3.964	5.714	227
1.6%	474	23	0	POST	^/api/posts/[^/]+/comments$	581.064	1.226	6.696	8.163	328
1.4%	3424	0	0	GET	/api/memories	504.677	0.147	0.993	1.751	10732
1.2%	602	0	0	PUT	^/api/follows/[^/]+$	437.550	0.727	3.836	5.004	0
0.9%	454	0	0	PATCH	/api/me	341.263	0.752	4.053	5.100	275
0.8%	12020	0	0	GET	^/api/feeds/[^/]+	279.371	0.023	0.265	2.377	2061
0.7%	711	0	0	POST	/api/auth/logout	251.667	0.354	3.432	7.464	0
0.7%	228	3	0	PATCH	^/api/posts/[^/]+$	246.035	1.079	6.347	7.340	645
0.6%	261	0	0	GET	/api/me/export	233.804	0.896	4.895	5.040	1813298
0.4%	103	0	0	DELETE	^/api/posts/[^/]+$	141.657	1.375	5.105	6.960	0
0.2%	3340	0	0	GET	/api/posting-window/current	92.346	0.028	0.361	1.618	235
0.2%	2485	0	0	GET	/api/users	77.896	0.031	0.260	7.472	246
0.2%	1859	37	0	GET	^/api/posts/me/current$	64.870	0.035	0.358	1.607	640
0.2%	84	0	0	DELETE	^/api/comments/[^/]+$	59.937	0.714	4.761	4.761	0
0.1%	89	0	0	DELETE	^/api/realmojis/[^/]+$	55.591	0.625	3.878	3.878	0
0.1%	619	0	0	GET	^/api/memories/[^/]+	42.821	0.069	0.614	1.139	587
0.1%	1637	0	0	GET	^/api/relationships/[^/]+	34.892	0.021	0.321	1.505	958
0.1%	40	0	0	PUT	^/api/blocks/[^/]+$	33.321	0.833	3.347	3.347	0
0.1%	1609	14	0	GET	^/api/posts/[^/]+$	28.030	0.017	0.130	1.592	604
0.0%	918	0	0	GET	^/api/users/[^/]+$	17.313	0.019	0.137	1.426	185
0.0%	545	3	0	GET	/api/me	14.721	0.027	0.150	0.315	301
0.0%	1	0	0	POST	/api/initialize	12.800	12.800	12.800	12.800	90
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.8%	248	1	0	GET	/api/notifications	2237.804	9.023	13.684	14.063	178
14.4%	297	2	0	POST	/api/posts	811.918	2.734	7.476	8.146	641
7.2%	535	7	0	POST	/api/friend-requests	405.742	0.758	4.239	6.152	114
5.6%	804	0	0	GET	^/api/feeds/[^/]+	314.503	0.391	2.375	2.661	2354
5.5%	1223	9	0	GET	^/api/images/[^/]+	310.142	0.254	1.398	2.825	150992
4.9%	412	0	0	GET	/api/posting-window/current	277.576	0.674	2.004	2.492	238
4.6%	355	0	0	POST	/api/auth/login	261.250	0.736	4.173	7.306	247
3.2%	346	0	0	GET	/api/memories	182.178	0.527	2.220	6.429	14291
2.5%	152	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	141.961	0.934	5.481	6.978	115
1.6%	105	7	0	POST	^/api/posts/[^/]+/realmojis$	88.152	0.840	5.291	8.372	436
1.5%	243	0	0	GET	/api/users	85.979	0.354	1.339	1.798	241
1.2%	202	10	0	GET	^/api/posts/me/current$	69.856	0.346	1.658	1.799	614
0.9%	96	0	0	POST	/api/auth/register	47.990	0.500	2.974	2.974	228
0.7%	122	0	0	GET	^/api/relationships/[^/]+	37.947	0.311	0.920	2.508	1039
0.6%	53	0	0	PUT	/api/me/profile-image	35.331	0.667	2.326	2.326	0
0.6%	96	6	0	GET	^/api/posts/[^/]+$	35.323	0.368	2.508	2.508	571
0.6%	17	0	0	GET	/api/me/export	35.160	2.068	4.083	4.083	1856980
0.6%	51	0	0	POST	/api/auth/logout	35.105	0.688	5.179	5.179	0
0.6%	68	0	0	GET	^/api/users/[^/]+$	33.487	0.492	2.508	2.508	188
0.6%	60	0	0	PUT	^/api/follows/[^/]+$	31.235	0.521	2.863	2.863	0
0.5%	54	0	0	GET	^/api/memories/[^/]+	29.783	0.552	2.117	2.117	964
0.4%	39	0	0	PATCH	/api/me	23.344	0.599	2.618	2.618	281
0.3%	67	0	0	GET	^/api/posts/[^/]+/realmojis$	19.172	0.286	0.920	0.920	1536
0.3%	58	2	0	GET	/api/me	17.521	0.302	0.951	0.951	290
0.2%	19	0	0	DELETE	^/api/realmojis/[^/]+$	13.957	0.735	4.859	4.859	0
0.2%	10	0	0	DELETE	^/api/posts/[^/]+$	11.924	1.192	2.326	2.326	0
0.2%	8	0	0	POST	^/api/posts/[^/]+/comments$	11.416	1.427	2.226	2.226	347
0.1%	18	1	0	GET	^/api/posts/[^/]+/comments$	7.990	0.444	2.375	2.375	1548
0.1%	12	3	0	PATCH	^/api/posts/[^/]+$	6.953	0.579	1.562	1.562	498
0.1%	1	0	0	POST	/api/me/account-deletion	3.452	3.452	3.452	3.452	0
…（上位 30 / 全 31 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.3%	213	1	0	GET	/api/notifications	1733.623	8.139	10.528	11.081	174
10.8%	245	8	0	POST	/api/posts	504.534	2.059	6.873	7.293	631
10.6%	532	11	0	POST	/api/friend-requests	492.324	0.925	5.422	7.652	113
5.7%	967	5	0	GET	^/api/images/[^/]+	267.151	0.276	1.328	3.890	152220
5.2%	501	0	0	GET	^/api/feeds/[^/]+	242.735	0.485	5.044	5.234	2089
4.3%	277	0	0	POST	/api/auth/login	198.401	0.716	5.401	6.769	248
4.0%	175	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	184.117	1.052	4.981	5.585	115
3.5%	345	0	0	GET	/api/posting-window/current	164.931	0.478	2.283	5.082	239
3.4%	294	0	0	GET	/api/memories	160.521	0.546	1.916	2.147	14020
3.3%	271	0	0	GET	/api/users	151.747	0.560	4.907	5.150	234
2.2%	189	18	0	GET	^/api/posts/me/current$	101.456	0.537	5.083	5.217	591
1.2%	16	0	0	GET	/api/me/export	56.479	3.530	13.257	13.257	2885954
1.1%	49	0	0	POST	/api/auth/logout	52.358	1.069	5.993	5.993	0
1.1%	97	0	0	GET	^/api/relationships/[^/]+	52.335	0.540	5.118	5.118	1452
0.7%	46	0	0	PUT	/api/me/profile-image	32.087	0.698	5.605	5.605	0
0.6%	44	0	0	PUT	^/api/follows/[^/]+$	29.984	0.681	4.647	4.647	0
0.6%	82	0	0	POST	/api/auth/register	29.976	0.366	1.730	1.730	229
0.6%	21	4	0	PATCH	^/api/posts/[^/]+$	28.242	1.345	2.682	2.682	559
0.5%	27	2	0	POST	^/api/posts/[^/]+/comments$	25.108	0.930	5.556	5.556	326
0.5%	42	7	0	GET	^/api/posts/[^/]+$	23.996	0.571	5.239	5.239	511
0.5%	22	1	0	POST	^/api/posts/[^/]+/realmojis$	21.358	0.971	3.760	3.760	473
0.4%	42	0	0	PATCH	/api/me	20.198	0.481	2.785	2.785	276
0.3%	51	1	0	GET	/api/me	16.070	0.315	1.012	1.012	292
0.3%	15	1	0	GET	^/api/posts/[^/]+/realmojis$	14.368	0.958	5.103	5.103	965
0.3%	32	0	0	GET	^/api/memories/[^/]+	13.200	0.413	2.367	2.367	877
0.2%	26	0	0	GET	^/api/users/[^/]+$	10.180	0.392	2.417	2.417	170
0.2%	18	0	0	GET	^/api/posts/[^/]+/comments$	9.338	0.519	2.284	2.284	710
0.2%	1	0	0	POST	/api/me/account-deletion	7.688	7.688	7.688	7.688	0
0.1%	5	0	0	DELETE	^/api/posts/[^/]+$	4.391	0.878	1.586	1.586	0
0.1%	3	0	0	DELETE	^/api/realmojis/[^/]+$	2.576	0.859	1.445	1.445	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 68001 クエリ / 81 種類 / 合計 662.895s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	66.4%	4169	440.170	105.58	815.35	667.7	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	16.1%	294	106.840	363.40	1795.76	1185.7	120.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	3.8%	8818	25.092	2.85	62.67	10.9	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	3.2%	3	21.525	7174.97	9571.55	30951.0	1199757.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
5	3.2%	1931	21.411	11.09	789.09	117.1	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
6	2.4%	3	15.999	5333.15	7201.61	24286.0	627545.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	1.5%	2810	10.065	3.58	105.32	24.8	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
8	1.2%	3	7.935	2644.86	3593.35	12635.0	529719.7	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	0.4%	721	2.366	3.28	24.57	28.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.2%	6	1.543	257.22	344.82	27825.2	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.2%	6076	1.474	0.24	79.23	19.9	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
12	0.2%	6	1.139	189.87	229.88	218621.0	2591.7	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
13	0.2%	2122	1.025	0.48	81.17	20.6	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
14	0.1%	484	0.986	2.04	34.92	21.4	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
15	0.1%	4772	0.782	0.16	87.13	11.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 3
SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 3ef271a5e1bf5b02f27cb5ef0067ed72018072af
Type: cpu
Time: 2026-10-31 13:52:58 JST
Duration: 60.04s, Total samples = 46.90s (78.12%)
Showing nodes accounting for 33.56s, 71.56% of 46.90s total
Dropped 1082 nodes (cum <= 0.23s)
Showing top 40 nodes out of 287
      flat  flat%   sum%        cum   cum%
     5.30s 11.30% 11.30%      5.30s 11.30%  internal/runtime/syscall/linux.Syscall6
     4.11s  8.76% 20.06%      4.11s  8.76%  crypto/internal/fips140/sha256.blockAVX2
     2.11s  4.50% 24.56%      2.11s  4.50%  runtime.memmove
     1.95s  4.16% 28.72%      4.51s  9.62%  example.com/isureal-company/isureal.(*Store).digestLocked
     1.86s  3.97% 32.69%      1.86s  3.97%  runtime.memclrNoHeapPointers
     1.72s  3.67% 36.35%      6.87s 14.65%  runtime.scanObject
     1.67s  3.56% 39.91%      5.82s 12.41%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.67s  3.56% 43.48%      2.22s  4.73%  runtime.findObject
     1.59s  3.39% 46.87%      1.59s  3.39%  example.com/isureal-company/isureal.h64 (inline)
     1.52s  3.24% 50.11%      1.52s  3.24%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.31s  2.79% 52.90%      2.41s  5.14%  runtime.tryDeferToSpanScan
     1.28s  2.73% 55.63%      1.28s  2.73%  aeshashbody
     0.85s  1.81% 57.44%      0.85s  1.81%  runtime.spanClass.sizeclass (inline)
     0.73s  1.56% 59.00%      2.29s  4.88%  runtime.mapaccess1_faststr
     0.42s   0.9% 59.89%      0.42s   0.9%  memeqbody
     0.35s  0.75% 60.64%      0.56s  1.19%  example.com/isureal-company/isureal.postLess
     0.34s  0.72% 61.36%      0.34s  0.72%  runtime.futex
     0.28s   0.6% 61.96%      0.66s  1.41%  database/sql.convertAssignRows
     0.28s   0.6% 62.56%      0.28s   0.6%  indexbody
     0.28s   0.6% 63.16%      0.74s  1.58%  runtime.greyobject
     0.28s   0.6% 63.75%      1.96s  4.18%  runtime.mapassign_faststr
     0.27s  0.58% 64.33%      1.11s  2.37%  runtime.mallocgcSmallScanNoHeader
     0.27s  0.58% 64.90%      0.27s  0.58%  runtime.typePointers.nextFast (inline)
     0.26s  0.55% 65.46%      4.63s  9.87%  crypto/internal/fips140/sha256.(*Digest).Write
     0.26s  0.55% 66.01%      0.48s  1.02%  internal/runtime/maps.(*Iter).Next
     0.26s  0.55% 66.57%      0.28s   0.6%  runtime.(*mspan).markBitsForIndex (inline)
     0.26s  0.55% 67.12%      0.29s  0.62%  runtime.spanOf (inline)
     0.25s  0.53% 67.65%      5.03s 10.72%  runtime.mallocgc
     0.21s  0.45% 68.10%      4.68s  9.98%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.20s  0.43% 68.53%      0.33s   0.7%  compress/flate.(*compressor).deflate
     0.20s  0.43% 68.96%      0.24s  0.51%  runtime.typePointers.next
     0.18s  0.38% 69.34%      4.09s  8.72%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.18s  0.38% 69.72%      0.57s  1.22%  runtime.mapaccess2_faststr
     0.15s  0.32% 70.04%      2.88s  6.14%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.15s  0.32% 70.36%      0.81s  1.73%  runtime.mallocgcSmallNoscan
     0.12s  0.26% 70.62%      4.93s 10.51%  example.com/isureal-company/isureal.computeImageSignature
     0.12s  0.26% 70.87%      0.45s  0.96%  runtime.scanObjectsSmall
     0.11s  0.23% 71.11%      4.31s  9.19%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.11s  0.23% 71.34%      0.87s  1.86%  runtime.slicebytetostring
     0.10s  0.21% 71.56%      2.27s  4.84%  example.com/isureal-company/isureal.listMemories
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 51.2% / 100% | 23% | 0.5% | 9 | 162 MB | run queue が積もっている |
| isu2 (2 cpu) | 30.0% / 100% | 3% | 0.4% | 4 | 306 MB |  |
| isu3 (2 cpu) | 32.1% / 100% | 4% | 0.8% | 4 | 128 MB |  |
| isu4 (2 cpu) | 50.7% / 100% | 28% | 11.7% | 27 | 100 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 72% / nginx 13% |
| isu2 | isureal 58% |
| isu3 | isureal 61% |
| isu4 | postgres 66% / isureal 6% |

## Diff from previous run

スコア: 64700 → 104789 (+40089)

`e0e4cb3..319c947`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf     |   3 +
 conf/isu1/home/isucon/env.sh                       |   2 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf     |  21 +-
 conf/isu2/home/isucon/env.sh                       |   5 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf     |  21 +-
 conf/isu3/home/isucon/env.sh                       |   5 +
 .../etc/postgresql/16/main/conf.d/zz-tuning.conf   |   2 +
 webapp/go/account_deletion_api.go                  |  10 +-
 webapp/go/action_log.go                            |  13 +-
 webapp/go/action_log_test.go                       |  67 ++
 webapp/go/admin_posting_window_api.go              |  16 +-
 webapp/go/admin_user_api.go                        |   5 +-
 webapp/go/cmd/isureal/main.go                      |   6 +
 webapp/go/comment_api.go                           |   9 +-
 webapp/go/config.go                                |   2 +
 webapp/go/export_api.go                            |   6 +-
 webapp/go/feed_api.go                              | 113 ++-
 webapp/go/feed_test.go                             | 331 ++++++++
 webapp/go/friend_request_api.go                    |  15 +-
 webapp/go/image_local.go                           |   4 +-
 webapp/go/initialize.go                            |   4 +-
 webapp/go/offload.go                               |  46 +-
 webapp/go/post_api.go                              |  78 +-
 webapp/go/post_forward.go                          | 113 +++
 webapp/go/post_forward_test.go                     |  94 +++
 webapp/go/profile_image_api.go                     |  20 +-
 webapp/go/realmoji_api.go                          |  16 +-
 webapp/go/relationship.go                          |  10 +-
 webapp/go/relationship_api.go                      |  33 +-
 webapp/go/repl.go                                  | 922 +++++++++++++++++++++
 webapp/go/repl_digest.go                           | 151 ++++
 webapp/go/repl_test.go                             |  62 ++
 webapp/go/replica_server.go                        | 255 ++++++
 webapp/go/server.go                                |   1 +
 webapp/go/store.go                                 | 198 ++++-
 webapp/go/user_api.go                              |  34 +-
 webapp/go/user_read_api.go                         |  58 +-
 webapp/go/user_search_test.go                      |  94 +++
 38 files changed, 2626 insertions(+), 219 deletions(-)
```
