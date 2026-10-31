# run 20261031-112112-976eb11

## Summary

| | |
|---|---|
| score | 18863 |
| status | pass |
| ref | posting-window-fast |
| sha | 976eb112bab2 |
| 所要時間 | 291s |
| 前回比 | +3195 (20.4%) vs 20261031-110450-8f720c2 |
| ベスト比 | -5938 (-23.9%) vs 20261031-103811-cef5fe3 = 24801 |
| note | 投稿期間: 専用プール・1往復＋通知を期間の行の配列に |

## Bench

```
{"id":"3782","status":"done","portalStatus":"finished","passed":true,"score":18863,"raw":18863,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:24:49.076Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:24:52.105Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:24:55.024Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:24:58.143Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T02:25:01.221Z level=INFO msg="europe 2 日目: 14 人"
time=2026-10-31T02:25:04.421Z level=INFO msg="america 2 日目: 12 人"
time=2026-10-31T02:25:07.357Z level=INFO msg="asia 3 日目: 127 人"
time=2026-10-31T02:25:11.600Z level=INFO msg="europe 3 日目: 16 人"
time=2026-10-31T02:25:13.174Z level=INFO msg="america 3 日目: 12 人"
time=2026-10-31T02:25:16.631Z level=INFO msg="asia 4 日目: 127 人"
time=2026-10-31T02:25:20.313Z level=INFO msg="europe 4 日目: 16 人"
time=2026-10-31T02:25:22.247Z level=INFO msg="america 4 日目: 12 人"
time=2026-10-31T02:25:25.523Z level=INFO msg="asia 5 日目: 126 人"
time=2026-10-31T02:25:30.763Z level=INFO msg="europe 5 日目: 15 人"
time=2026-10-31T02:25:32.166Z level=INFO msg="america 5 日目: 12 人"
time=2026-10-31T02:25:40.157Z level=INFO msg="europe 最終日（グランプリ）: 14 人"
time=2026-10-31T02:25:40.221Z level=INFO msg="america 最終日（グランプリ）: 12 人"
time=2026-10-31T02:25:40.377Z level=INFO msg="asia 最終日（グランプリ）: 118 人"
time=2026-10-31T02:25:49.022Z level=INFO msg="グランプリ: アジア 5 人, ヨーロッパ 1 人, アメリカ 1 人"
{
  "pass": true,
  "score": 18863,
  "messages": [
    "nean5929: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "wamiake5961: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "nevinchi-5972: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "isorvin.5975: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "torohi.5980: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "wael.5989: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ulyu5998: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "furu6004: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kekino-6014: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.6%	374	0	0	GET	/api/notifications	3511.651	9.389	14.871	15.288	176
17.3%	439	44	0	POST	/api/posts	1859.991	4.237	10.266	10.849	579
10.4%	2065	1272	0	POST	/api/auth/login	1118.995	0.542	3.513	7.639	112
10.3%	3492	799	0	GET	^/api/images/[^/]+	1113.439	0.319	2.551	5.074	116076
6.9%	143	2	0	POST	/api/auth/register	740.187	5.176	15.593	15.686	227
4.9%	209	146	0	PUT	/api/me/profile-image	526.100	2.517	10.068	10.355	11
4.3%	721	19	0	POST	/api/friend-requests	464.467	0.644	4.551	7.129	113
3.6%	378	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	385.364	1.019	7.636	9.068	116
1.8%	1940	0	0	GET	^/api/feeds/[^/]+	192.822	0.099	0.485	3.322	1904
1.4%	563	0	0	GET	/api/posting-window/current	147.983	0.263	1.575	2.683	237
1.2%	877	0	0	GET	/api/memories	128.214	0.146	1.368	3.670	9716
1.0%	52	0	0	GET	/api/me/export	107.144	2.060	13.155	13.155	2400167
0.5%	71	5	0	POST	^/api/posts/[^/]+/realmojis$	52.986	0.746	3.559	3.559	427
0.5%	337	4	0	GET	^/api/posts/me/current$	52.491	0.156	1.040	1.362	636
0.4%	136	0	0	POST	/api/auth/logout	45.001	0.331	4.854	5.239	0
0.4%	76	0	0	PUT	^/api/follows/[^/]+$	43.681	0.575	3.175	3.175	0
0.4%	66	6	0	POST	^/api/posts/[^/]+/comments$	41.316	0.626	3.710	3.710	313
0.4%	325	0	0	GET	/api/users	40.052	0.123	0.690	1.661	236
0.3%	41	0	0	PATCH	^/api/posts/[^/]+$	27.073	0.660	2.980	2.980	640
0.3%	367	0	0	GET	^/api/relationships/[^/]+	26.980	0.074	0.676	2.679	1632
0.2%	309	9	0	GET	^/api/posts/[^/]+$	24.063	0.078	0.281	0.503	591
0.2%	50	0	0	PATCH	/api/me	17.339	0.347	2.688	2.688	267
0.2%	143	0	0	GET	^/api/users/[^/]+$	16.663	0.117	1.045	3.080	197
0.1%	123	0	0	GET	^/api/memories/[^/]+	13.175	0.107	1.605	1.608	631
0.1%	1	0	0	POST	/api/initialize	12.034	12.034	12.034	12.034	90
0.1%	20	0	0	DELETE	^/api/posts/[^/]+$	11.421	0.571	3.057	3.057	0
0.1%	117	0	0	GET	^/api/posts/[^/]+/realmojis$	10.179	0.087	0.344	0.416	326
0.1%	18	0	0	POST	/api/admin/posting-window	10.030	0.557	2.741	2.741	176
0.1%	100	0	0	GET	^/api/posts/[^/]+/comments$	8.543	0.085	0.340	0.360	210
0.1%	70	3	0	GET	/api/me	6.384	0.091	1.120	1.120	286
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.9%	71	0	0	GET	/api/notifications	606.703	8.545	10.337	10.337	176
17.3%	70	1	0	POST	/api/posts	245.001	3.500	10.897	10.897	642
5.9%	300	0	0	GET	^/api/feeds/[^/]+	83.982	0.280	1.032	1.744	1591
5.4%	77	0	0	POST	/api/auth/login	76.201	0.990	7.928	7.928	252
4.2%	349	3	0	GET	^/api/images/[^/]+	58.993	0.169	0.957	0.959	213160
3.9%	42	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	55.545	1.322	4.732	4.732	114
3.5%	87	4	0	POST	/api/friend-requests	49.248	0.566	4.170	4.170	111
3.2%	23	5	0	POST	^/api/posts/[^/]+/realmojis$	45.090	1.960	5.348	5.348	392
2.5%	124	0	0	GET	/api/memories	35.131	0.283	1.851	3.612	11570
2.4%	14	0	0	POST	/api/auth/register	33.654	2.404	15.491	15.491	228
1.5%	79	0	0	GET	/api/posting-window/current	20.733	0.262	1.116	1.116	238
1.0%	57	2	0	GET	^/api/posts/me/current$	14.563	0.255	1.117	1.117	628
0.9%	44	7	0	GET	^/api/posts/[^/]+$	13.182	0.300	0.794	0.794	500
0.7%	49	0	0	GET	^/api/relationships/[^/]+	9.802	0.200	1.152	1.152	1451
0.7%	9	0	0	PATCH	^/api/posts/[^/]+$	9.651	1.072	4.672	4.672	666
0.7%	6	0	0	GET	/api/me/export	9.443	1.574	3.551	3.551	2614334
0.5%	37	0	0	GET	/api/users	7.772	0.210	0.565	0.565	233
0.5%	25	0	0	GET	^/api/posts/[^/]+/realmojis$	7.515	0.301	1.031	1.031	335
0.4%	10	0	0	POST	/api/auth/logout	6.102	0.610	1.990	1.990	0
0.3%	19	0	0	GET	^/api/users/[^/]+$	3.905	0.206	0.537	0.537	176
0.3%	2	0	0	DELETE	^/api/realmojis/[^/]+$	3.889	1.945	2.294	2.294	0
0.3%	7	0	0	PUT	/api/me/profile-image	3.883	0.555	0.915	0.915	0
0.2%	13	0	0	GET	^/api/memories/[^/]+	3.311	0.255	0.952	0.952	826
0.2%	9	0	0	PUT	^/api/follows/[^/]+$	3.273	0.364	1.189	1.189	0
0.1%	12	2	0	GET	/api/me	1.856	0.155	0.233	0.233	230
0.1%	10	0	0	GET	^/api/posts/[^/]+/comments$	1.818	0.182	0.388	0.388	29
0.1%	3	0	0	DELETE	^/api/posts/[^/]+$	1.299	0.433	0.552	0.552	0
0.1%	1	0	0	POST	/api/me/account-deletion	1.265	1.265	1.265	1.265	0
0.1%	3	0	0	PATCH	/api/me	0.868	0.289	0.467	0.467	333
0.0%	2	0	0	PUT	^/api/blocks/[^/]+$	0.262	0.131	0.154	0.154	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.1%	57	0	0	GET	/api/notifications	487.269	8.549	10.005	10.005	182
18.7%	55	1	0	POST	/api/posts	211.888	3.853	11.138	11.138	608
6.1%	46	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	69.216	1.505	6.047	6.047	116
5.6%	82	2	0	POST	/api/friend-requests	62.923	0.767	6.087	6.087	113
4.6%	69	0	0	POST	/api/auth/login	51.890	0.752	8.531	8.531	257
4.4%	281	2	0	GET	^/api/images/[^/]+	49.459	0.176	0.950	0.951	169368
3.7%	142	0	0	GET	^/api/feeds/[^/]+	42.133	0.297	0.890	1.112	1193
3.2%	105	0	0	GET	/api/memories	35.781	0.341	1.691	7.114	11340
2.2%	5	0	0	GET	/api/me/export	25.425	5.085	15.023	15.023	3626832
1.7%	63	0	0	GET	/api/posting-window/current	19.471	0.309	2.421	2.421	240
1.2%	44	6	0	GET	^/api/posts/me/current$	13.173	0.299	0.799	0.799	527
0.9%	39	0	0	GET	^/api/relationships/[^/]+	9.760	0.250	1.224	1.224	2049
0.9%	35	0	0	GET	/api/users	9.645	0.276	1.144	1.144	236
0.7%	11	0	0	POST	/api/auth/register	8.139	0.740	5.439	5.439	228
0.6%	5	2	0	PATCH	^/api/posts/[^/]+$	7.063	1.413	5.510	5.510	396
0.5%	17	2	0	GET	^/api/posts/[^/]+$	5.655	0.333	0.959	0.959	534
0.3%	1	0	0	POST	^/api/posts/[^/]+/comments$	3.788	3.788	3.788	3.788	387
0.3%	11	0	0	GET	^/api/users/[^/]+$	3.520	0.320	0.557	0.557	181
0.3%	14	0	0	GET	^/api/memories/[^/]+	2.981	0.213	0.569	0.569	951
0.2%	9	0	0	POST	/api/auth/logout	2.784	0.309	0.925	0.925	0
0.2%	4	0	0	PUT	/api/me/profile-image	2.006	0.501	1.306	1.306	0
0.2%	7	0	0	GET	^/api/posts/[^/]+/realmojis$	1.789	0.256	0.469	0.469	336
0.1%	7	0	0	PUT	^/api/follows/[^/]+$	1.395	0.199	0.640	0.640	0
0.1%	7	0	0	GET	/api/me	0.961	0.137	0.317	0.317	256
0.1%	4	0	0	GET	^/api/posts/[^/]+/comments$	0.879	0.220	0.334	0.334	29
0.1%	2	0	0	DELETE	^/api/posts/[^/]+$	0.816	0.408	0.633	0.633	0
0.1%	2	0	0	PATCH	/api/me	0.606	0.303	0.461	0.461	327
0.0%	1	0	0	POST	^/api/posts/[^/]+/realmojis$	0.329	0.329	0.329	0.329	407
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.106	0.106	0.106	0.106	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 13872 クエリ / 80 種類 / 合計 86.149s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	60.2%	63	51.853	823.07	8989.56	7445.1	161.7	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	20.9%	677	18.000	26.59	446.96	737.5	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	4.2%	540	3.660	6.78	204.19	118.8	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.0%	1882	3.470	1.84	12.59	12.8	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	2.8%	1	2.433	2433.44	2433.44	29357.0	1199460.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	1.8%	1	1.550	1550.06	1550.06	19735.0	625595.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	1.3%	1	1.095	1094.54	1094.54	11528.0	529706.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	1.2%	518	0.992	1.92	11.39	36.1	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
9	0.4%	2	0.334	166.94	176.63	214618.5	2524.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
10	0.3%	2	0.286	142.92	277.20	20229.0	2525.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.3%	85	0.273	3.21	5.32	51.3	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
12	0.3%	1	0.232	232.33	232.33	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
13	0.2%	1	0.211	211.44	211.44	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.2%	1	0.208	207.88	207.88	69904.0	29792.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.2%	847	0.167	0.20	1.89	22.2	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)

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
Build ID: 59405859360b31b4e1e6b39e2c6c3633a658653d
Type: cpu
Time: 2026-10-31 11:24:32 JST
Duration: 60.32s, Total samples = 88.73s (147.11%)
Showing nodes accounting for 78.99s, 89.02% of 88.73s total
Dropped 755 nodes (cum <= 0.44s)
Showing top 40 nodes out of 177
      flat  flat%   sum%        cum   cum%
    14.94s 16.84% 16.84%     14.94s 16.84%  crypto/internal/fips140/sha256.blockAVX2
    12.08s 13.61% 30.45%     12.08s 13.61%  runtime.spanClass.sizeclass (inline)
     8.85s  9.97% 40.43%     14.32s 16.14%  image/jpeg.(*decoder).reconstructBlock
     4.38s  4.94% 45.36%     30.86s 34.78%  image/jpeg.(*decoder).processSOS
     4.11s  4.63% 49.99%     16.80s 18.93%  runtime.tryDeferToSpanScan
     3.40s  3.83% 53.83%      5.54s  6.24%  image/jpeg.(*decoder).decodeHuffman
     2.87s  3.23% 57.06%         3s  3.38%  image/jpeg.(*decoder).receiveExtend
     2.52s  2.84% 59.90%      2.52s  2.84%  runtime.memclrNoHeapPointers
     2.41s  2.72% 62.62%      2.41s  2.72%  runtime.memmove
     2.29s  2.58% 65.20%      2.92s  3.29%  image/jpeg.idctCols
     1.86s  2.10% 67.29%      2.46s  2.77%  image/jpeg.idctRows
     1.82s  2.05% 69.35%      7.80s  8.79%  runtime.scanObjectsSmall
     1.74s  1.96% 71.31%     13.47s 15.18%  runtime.scanObject
     1.68s  1.89% 73.20%      2.23s  2.51%  image/jpeg.(*decoder).ensureNBits
     1.63s  1.84% 75.04%      1.63s  1.84%  internal/runtime/syscall/linux.Syscall6
     1.34s  1.51% 76.55%      5.44s  6.13%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.21s  1.36% 77.91%      1.21s  1.36%  image/jpeg.dctBox (inline)
     1.05s  1.18% 79.09%      1.05s  1.18%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.92s  1.04% 80.13%     17.03s 19.19%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.90s  1.01% 81.15%      0.90s  1.01%  aeshashbody
     0.79s  0.89% 82.04%     16.67s 18.79%  crypto/internal/fips140/sha256.(*Digest).Write
     0.69s  0.78% 82.81%      0.69s  0.78%  memeqbody
     0.60s  0.68% 83.49%      9.25s 10.42%  runtime.scanSpan
     0.57s  0.64% 84.13%     17.93s 20.21%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.52s  0.59% 84.72%      0.54s  0.61%  image/jpeg.(*decoder).readByteStuffedByte
     0.52s  0.59% 85.30%      0.52s  0.59%  runtime.extractHeapBitsSmall
     0.47s  0.53% 85.83%     19.25s 21.70%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.38s  0.43% 86.26%      0.95s  1.07%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.36s  0.41% 86.67%      0.60s  0.68%  example.com/isureal-company/isureal.postLess
     0.30s  0.34% 87.01%     20.74s 23.37%  example.com/isureal-company/isureal.imageSignature
     0.25s  0.28% 87.29%      2.03s  2.29%  runtime.mapaccess1_faststr
     0.24s  0.27% 87.56%     15.18s 17.11%  crypto/internal/fips140/sha256.block
     0.24s  0.27% 87.83%      0.69s  0.78%  database/sql.convertAssignRows
     0.23s  0.26% 88.09%      0.71s   0.8%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.19s  0.21% 88.30%      0.77s  0.87%  runtime.scanObjectSmall
     0.17s  0.19% 88.49%     13.80s 15.55%  runtime.mallocgc
     0.16s  0.18% 88.67%      1.13s  1.27%  runtime.mapassign_faststr
     0.12s  0.14% 88.81%      1.13s  1.27%  runtime.mallocgcSmallScanNoHeader
     0.11s  0.12% 88.93%      0.48s  0.54%  runtime.(*sweepLocked).sweep
     0.08s  0.09% 89.02%      0.77s  0.87%  database/sql.(*Rows).scanLocked
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 75.1% / 100% | 65% | 0.3% | 6 | 106 MB | CPU が飽和 |
| isu2 (2 cpu) | 0.9% / 3% | 0% | 0.2% | 2 | 325 MB |  |
| isu3 (2 cpu) | 1.1% / 3% | 0% | 0.2% | 0 | 269 MB |  |
| isu4 (2 cpu) | 25.0% / 78% | 0% | 14.3% | 5 | 100 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 134% / nginx 11% |
| isu2 | 目立つプロセス無し |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 18% |

## Diff from previous run

スコア: 15668 → 18863 (+3195)

`8f720c2..976eb11`

```
 webapp/go/admin_posting_window_api.go | 54 ++++++++++++++++++++++-------------
 webapp/go/cmd/image-export/main.go    | 11 +------
 webapp/go/image_local.go              | 37 ++++++++++++++++++------
 webapp/go/store.go                    | 21 ++++++++++++++
 4 files changed, 85 insertions(+), 38 deletions(-)
```
