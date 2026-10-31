# run 20261031-124932-86828dc

## Summary

| | |
|---|---|
| score | 63311 |
| status | pass |
| ref | repl-replica |
| sha | 86828dc8e996 |
| 所要時間 | 210s |
| 前回比 | +63311 (—) vs 20261031-124342-0db3f89 |
| ベスト比 | -1389 (-2.1%) vs 20261031-121235-e0e4cb3 = 64700 |
| note | 複製 第1段 shadow（europe1・america1 が裏でフィード等を計算して asia1 の応答と比べる） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	141	400	566	565
europe	7	15	36	81	83	82
america	6	15	37	64	64	64
```

```
{"id":"9937","status":"done","portalStatus":"finished","passed":true,"score":63311,"raw":63311,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:51:34.741Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:51:37.689Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:51:40.688Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:51:43.839Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T03:51:46.794Z level=INFO msg="europe 2 日目: 15 人"
time=2026-10-31T03:51:49.790Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T03:51:52.845Z level=INFO msg="asia 3 日目: 141 人"
time=2026-10-31T03:51:55.827Z level=INFO msg="europe 3 日目: 36 人"
time=2026-10-31T03:51:58.076Z level=INFO msg="user moved" user=rote.6151 from=asia to=america
time=2026-10-31T03:51:58.794Z level=INFO msg="america 3 日目: 37 人"
time=2026-10-31T03:52:01.869Z level=INFO msg="asia 4 日目: 400 人"
time=2026-10-31T03:52:04.975Z level=INFO msg="europe 4 日目: 81 人"
time=2026-10-31T03:52:08.074Z level=INFO msg="america 4 日目: 64 人"
time=2026-10-31T03:52:09.945Z level=INFO msg="user moved" user=toben_6695 from=asia to=europe
time=2026-10-31T03:52:11.498Z level=INFO msg="asia 5 日目: 566 人"
time=2026-10-31T03:52:12.102Z level=INFO msg="user moved" user=chima6861 from=asia to=europe
time=2026-10-31T03:52:15.634Z level=INFO msg="europe 5 日目: 83 人"
time=2026-10-31T03:52:16.878Z level=INFO msg="america 5 日目: 64 人"
time=2026-10-31T03:52:25.967Z level=INFO msg="america 最終日（グランプリ）: 64 人"
time=2026-10-31T03:52:26.570Z level=INFO msg="asia 最終日（グランプリ）: 565 人"
time=2026-10-31T03:52:26.827Z level=INFO msg="europe 最終日（グランプリ）: 82 人"
time=2026-10-31T03:52:34.703Z level=INFO msg="グランプリ: アジア 1 人, ヨーロッパ 1 人, アメリカ 2 人"
{
  "pass": true,
  "score": 63311,
  "messages": [
    "yomu5095: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
30.9%	1093	1	0	GET	/api/notifications	9671.100	8.848	15.490	15.762	175
22.3%	1466	57	0	POST	/api/posts	6981.365	4.762	15.627	16.557	623
10.9%	3142	53	0	POST	/api/friend-requests	3411.601	1.086	13.076	17.079	114
9.4%	7403	613	0	GET	^/api/images/[^/]+	2945.066	0.398	7.148	17.728	150488
7.6%	3037	863	0	POST	/api/auth/login	2365.279	0.779	11.880	15.890	181
4.9%	1156	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1518.869	1.314	14.168	16.880	116
2.4%	584	2	0	POST	/api/auth/register	752.979	1.289	6.557	10.039	226
1.9%	2070	0	0	GET	/api/posting-window/current	595.036	0.287	4.225	14.371	236
1.5%	423	122	0	PUT	/api/me/profile-image	466.049	1.102	8.797	10.612	7
1.2%	359	7	0	POST	^/api/posts/[^/]+/realmojis$	388.019	1.081	4.294	6.676	470
1.0%	344	5	0	POST	^/api/posts/[^/]+/comments$	320.198	0.931	5.576	18.744	356
0.7%	316	0	0	PUT	^/api/follows/[^/]+$	232.201	0.735	5.191	6.634	0
0.7%	266	0	0	PATCH	/api/me	220.543	0.829	5.592	17.622	271
0.7%	6835	0	0	GET	^/api/feeds/[^/]+	217.606	0.032	0.191	6.066	1699
0.7%	1923	0	0	GET	/api/memories	204.347	0.106	0.626	5.326	9891
0.5%	140	0	0	GET	/api/me/export	169.717	1.212	14.304	15.083	1846190
0.5%	323	0	0	POST	/api/auth/logout	149.221	0.462	4.953	15.040	0
0.3%	104	10	0	PATCH	^/api/posts/[^/]+$	101.663	0.978	5.087	5.970	571
0.3%	53	0	0	DELETE	^/api/posts/[^/]+$	81.152	1.531	14.072	14.072	0
0.3%	54	0	0	DELETE	^/api/comments/[^/]+$	80.464	1.490	16.046	16.046	0
0.2%	57	0	0	DELETE	^/api/realmojis/[^/]+$	67.408	1.183	13.182	13.182	0
0.2%	1574	0	0	GET	/api/users	67.230	0.043	0.333	3.331	241
0.2%	1032	54	0	GET	^/api/posts/me/current$	52.673	0.051	0.441	1.869	615
0.1%	975	17	0	GET	^/api/posts/[^/]+$	23.772	0.024	0.143	1.145	612
0.1%	12	0	0	PUT	^/api/admin/users/[^/]+/certification$	23.734	1.978	10.995	10.995	0
0.1%	23	0	0	PUT	^/api/blocks/[^/]+$	23.238	1.010	3.883	3.883	0
0.1%	985	0	0	GET	^/api/relationships/[^/]+	21.297	0.022	0.171	0.739	1080
0.1%	289	0	0	GET	^/api/memories/[^/]+	20.911	0.072	0.607	1.869	666
0.1%	1	0	0	POST	/api/initialize	18.948	18.948	18.948	18.948	90
0.0%	491	0	0	GET	^/api/users/[^/]+$	13.188	0.027	0.162	0.534	188
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.4%	195	1	0	GET	/api/notifications	1612.445	8.269	11.868	11.874	176
19.6%	233	5	0	POST	/api/posts	976.605	4.191	15.659	16.014	628
12.9%	490	4	0	POST	/api/friend-requests	642.716	1.312	13.161	15.795	114
7.5%	291	0	0	POST	/api/auth/login	373.399	1.283	15.085	17.106	248
6.6%	182	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	329.757	1.812	13.472	16.103	115
3.3%	918	9	0	GET	^/api/images/[^/]+	163.797	0.178	1.012	1.470	185307
2.9%	317	0	0	GET	/api/posting-window/current	143.446	0.453	5.500	7.913	239
2.7%	745	0	0	GET	^/api/feeds/[^/]+	135.578	0.182	0.524	3.792	2174
1.2%	289	0	0	GET	/api/memories	61.406	0.212	0.561	1.376	9825
1.0%	35	0	0	PUT	^/api/follows/[^/]+$	51.322	1.466	7.324	7.324	0
0.9%	21	0	0	GET	/api/me/export	45.501	2.167	8.239	8.239	2678034
0.9%	49	1	0	POST	^/api/posts/[^/]+/comments$	42.694	0.871	5.069	5.069	342
0.8%	40	0	0	PUT	/api/me/profile-image	42.178	1.054	7.349	7.349	0
0.8%	79	0	0	POST	/api/auth/register	41.907	0.530	1.899	1.899	227
0.8%	48	1	0	POST	^/api/posts/[^/]+/realmojis$	41.791	0.871	5.299	5.299	410
0.7%	211	0	0	GET	/api/users	35.192	0.167	0.482	0.547	248
0.6%	175	9	0	GET	^/api/posts/me/current$	31.126	0.178	0.436	3.479	609
0.6%	11	0	0	DELETE	^/api/posts/[^/]+$	30.208	2.746	15.699	15.699	0
0.6%	43	0	0	POST	/api/auth/logout	28.969	0.674	4.223	4.223	0
0.6%	33	0	0	PATCH	/api/me	28.502	0.864	4.490	4.490	285
0.4%	134	0	0	GET	^/api/relationships/[^/]+	20.977	0.157	0.677	1.048	1271
0.4%	121	9	0	GET	^/api/posts/[^/]+$	20.331	0.168	0.558	1.016	575
0.3%	21	0	0	PATCH	^/api/posts/[^/]+$	17.250	0.821	3.149	3.149	647
0.3%	13	0	0	DELETE	^/api/realmojis/[^/]+$	15.780	1.214	10.012	10.012	0
0.2%	60	0	0	GET	^/api/users/[^/]+$	9.740	0.162	0.809	0.809	187
0.2%	38	0	0	GET	^/api/memories/[^/]+	9.019	0.237	2.176	2.176	722
0.2%	45	2	0	GET	/api/me	7.565	0.168	0.340	0.340	285
0.1%	47	0	0	GET	^/api/posts/[^/]+/realmojis$	7.143	0.152	0.550	0.550	700
0.1%	7	0	0	DELETE	^/api/comments/[^/]+$	5.147	0.735	1.709	1.709	0
0.1%	1	0	0	POST	/api/me/account-deletion	3.933	3.933	3.933	3.933	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
30.8%	159	0	0	GET	/api/notifications	1189.377	7.480	9.277	9.278	175
21.6%	177	3	0	POST	/api/posts	834.134	4.713	16.496	16.631	633
11.1%	362	4	0	POST	/api/friend-requests	431.096	1.191	14.094	16.583	114
7.7%	213	0	0	POST	/api/auth/login	299.369	1.405	14.600	15.230	248
4.8%	109	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	185.395	1.701	14.597	14.928	115
4.2%	767	2	0	GET	^/api/images/[^/]+	162.127	0.211	1.017	2.180	186636
3.1%	612	0	0	GET	^/api/feeds/[^/]+	119.523	0.195	0.631	1.760	1943
2.2%	246	0	0	GET	/api/posting-window/current	85.716	0.348	3.462	7.781	240
1.8%	260	0	0	GET	/api/memories	69.499	0.267	1.760	4.360	11502
1.7%	62	0	0	POST	/api/auth/register	66.129	1.067	7.614	7.614	229
1.6%	57	4	0	POST	^/api/posts/[^/]+/realmojis$	63.810	1.119	3.971	3.971	434
1.1%	15	0	0	GET	/api/me/export	42.092	2.806	10.697	10.697	2851778
1.0%	30	0	0	PUT	/api/me/profile-image	40.132	1.338	5.490	5.490	0
1.0%	25	0	0	PATCH	/api/me	36.805	1.472	8.509	8.509	264
0.8%	165	0	0	GET	/api/users	32.684	0.198	0.997	1.764	239
0.8%	134	15	0	GET	^/api/posts/me/current$	31.559	0.236	1.492	4.359	570
0.8%	35	0	0	POST	/api/auth/logout	29.385	0.840	13.024	13.024	0
0.6%	20	2	0	PATCH	^/api/posts/[^/]+$	22.084	1.104	9.025	9.025	588
0.5%	38	0	0	POST	^/api/posts/[^/]+/comments$	21.204	0.558	3.623	3.623	273
0.4%	3	0	0	DELETE	^/api/posts/[^/]+$	16.841	5.614	14.098	14.098	0
0.4%	109	0	0	GET	^/api/relationships/[^/]+	16.322	0.150	0.558	0.707	1242
0.4%	96	7	0	GET	^/api/posts/[^/]+$	16.115	0.168	0.455	0.455	576
0.4%	21	0	0	PUT	^/api/follows/[^/]+$	14.201	0.676	5.614	5.614	0
0.2%	40	0	0	GET	^/api/users/[^/]+$	8.108	0.203	1.888	1.888	178
0.2%	8	0	0	DELETE	^/api/realmojis/[^/]+$	7.291	0.911	2.241	2.241	0
0.2%	33	0	0	GET	/api/me	6.510	0.197	0.838	0.838	297
0.1%	8	0	0	DELETE	^/api/comments/[^/]+$	4.801	0.600	1.354	1.354	0
0.1%	19	0	0	GET	^/api/memories/[^/]+	4.331	0.228	0.449	0.449	824
0.1%	29	0	0	GET	^/api/posts/[^/]+/realmojis$	4.317	0.149	0.320	0.320	995
0.1%	28	0	0	GET	^/api/posts/[^/]+/comments$	4.207	0.150	0.308	0.308	2240
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 41665 クエリ / 81 種類 / 合計 408.518s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	50.0%	2636	204.057	77.41	557.10	650.1	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	15.6%	176	63.581	361.25	1298.03	6391.7	129.8	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	10.4%	5	42.363	8472.65	10701.24	30211.0	1199881.2	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
4	8.4%	5	34.212	6842.33	11066.09	22035.0	627107.2	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	3.6%	5	14.841	2968.23	3706.09	11890.2	529867.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.9%	1488	11.992	8.06	387.18	118.6	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
7	2.6%	8	10.767	1345.88	5944.37	224624.3	2658.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
8	2.6%	4341	10.525	2.42	43.73	11.5	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
9	1.2%	1813	4.941	2.73	56.60	25.7	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.6%	8	2.367	295.93	538.42	25074.6	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.4%	452	1.574	3.48	42.89	30.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
12	0.3%	8	1.286	160.76	332.82	630.0	25594.6	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
13	0.2%	8	1.019	127.35	159.43	641.0	30692.3	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
14	0.2%	3919	0.847	0.22	15.55	20.1	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
15	0.2%	425	0.773	1.82	11.72	22.6	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)

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
File: isureal
Build ID: 111f9f64e733e7a4d506cf8c274a78166c95b1da
Type: cpu
Time: 2026-10-31 12:51:09 JST
Duration: 60.05s, Total samples = 37.94s (63.18%)
Showing nodes accounting for 28.99s, 76.41% of 37.94s total
Dropped 928 nodes (cum <= 0.19s)
Showing top 40 nodes out of 293
      flat  flat%   sum%        cum   cum%
     2.63s  6.93%  6.93%      2.63s  6.93%  internal/runtime/syscall/linux.Syscall6
     1.95s  5.14% 12.07%      1.95s  5.14%  crypto/internal/fips140/sha256.blockAVX2
     1.83s  4.82% 16.90%      5.60s 14.76%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.74s  4.59% 21.48%      2.26s  5.96%  runtime.findObject
     1.73s  4.56% 26.04%      1.73s  4.56%  runtime.memclrNoHeapPointers
     1.61s  4.24% 30.28%      1.61s  4.24%  runtime.memmove
     1.55s  4.09% 34.37%      5.69s 15.00%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.53s  4.03% 38.40%      2.32s  6.11%  runtime.tryDeferToSpanScan
     1.51s  3.98% 42.38%      1.51s  3.98%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.46s  3.85% 46.23%      2.49s  6.56%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.41s  3.72% 49.95%      6.21s 16.37%  runtime.scanObject
     1.25s  3.29% 53.24%      1.25s  3.29%  aeshashbody
     1.20s  3.16% 56.40%      1.28s  3.37%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.89s  2.35% 58.75%      1.10s  2.90%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.67s  1.77% 60.52%      0.67s  1.77%  memeqbody
     0.64s  1.69% 62.20%      0.64s  1.69%  example.com/isureal-company/isureal.(*UserRow).active
     0.52s  1.37% 63.57%      0.52s  1.37%  runtime.spanClass.sizeclass (inline)
     0.47s  1.24% 64.81%      1.95s  5.14%  runtime.mapaccess1_faststr
     0.43s  1.13% 65.95%      0.56s  1.48%  example.com/isureal-company/isureal.postLess
     0.36s  0.95% 66.90%      0.36s  0.95%  runtime.nextFreeFast (inline)
     0.33s  0.87% 67.76%      0.64s  1.69%  database/sql.convertAssignRows
     0.27s  0.71% 68.48%      0.27s  0.71%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.26s  0.69% 69.16%      0.26s  0.69%  indexbody
     0.23s  0.61% 69.77%      0.23s  0.61%  runtime.futex
     0.23s  0.61% 70.37%      0.27s  0.71%  runtime.spanOf (inline)
     0.21s  0.55% 70.93%      0.30s  0.79%  runtime.typePointers.next
     0.21s  0.55% 71.48%      0.21s  0.55%  runtime.typePointers.nextFast (inline)
     0.20s  0.53% 72.01%      0.20s  0.53%  sync/atomic.(*Int32).Add
     0.19s   0.5% 72.51%      0.21s  0.55%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.17s  0.45% 72.96%      3.89s 10.25%  runtime.mallocgc
     0.17s  0.45% 73.41%      0.83s  2.19%  runtime.mallocgcSmallScanNoHeader
     0.16s  0.42% 73.83%      0.19s   0.5%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.16s  0.42% 74.25%      0.34s   0.9%  runtime.greyobject
     0.13s  0.34% 74.59%      1.10s  2.90%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.13s  0.34% 74.93%      1.85s  4.88%  runtime.mapassign_faststr
     0.12s  0.32% 75.25%      0.30s  0.79%  compress/flate.(*compressor).deflate
     0.12s  0.32% 75.57%      2.17s  5.72%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.12s  0.32% 75.88%      0.22s  0.58%  runtime.heapArenaOf (inline)
     0.10s  0.26% 76.15%      2.63s  6.93%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.26% 76.41%      0.55s  1.45%  runtime.mallocgcSmallNoscan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 49.8% / 100% | 34% | 0.2% | 6 | 171 MB | CPU が飽和 |
| isu2 (2 cpu) | 19.7% / 61% | 0% | 0.3% | 6 | 716 MB | run queue が積もっている |
| isu3 (2 cpu) | 19.1% / 75% | 0% | 0.4% | 4 | 700 MB |  |
| isu4 (2 cpu) | 51.4% / 100% | 14% | 17.6% | 37 | 104 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 81% / nginx 11% / journalctl 2% |
| isu2 | isureal 37% |
| isu3 | isureal 36% |
| isu4 | postgres 54% / isureal 8% |

## Diff from previous run

スコア: 0 → 63311 (+63311)

`0db3f89..86828dc`

```
 webapp/go/repl.go  | 65 +++++++++++++++++++++++++++++++-----------------------
 webapp/go/store.go | 32 ++++++++++++++++++++++-----
 2 files changed, 65 insertions(+), 32 deletions(-)
```
