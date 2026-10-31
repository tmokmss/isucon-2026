# run 20261031-112920-4764fb5

## Summary

| | |
|---|---|
| score | 21430 |
| status | pass |
| ref | origin/main |
| sha | 4764fb57d8a3 |
| 所要時間 | 98s |
| 前回比 | +2155 (11.2%) vs 20261031-111738-5948f35 |
| ベスト比 | -3371 (-13.6%) vs 20261031-103811-cef5fe3 = 24801 |
| note | main: +#65 投稿期間の開始 +#67 無効化の受領待ち |

## Bench

```
{"id":"4114","status":"done","portalStatus":"finished","passed":true,"score":21430,"raw":21430,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:29:41.675Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:29:44.665Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:29:47.649Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:29:50.680Z level=INFO msg="asia 2 日目: 54 人"
time=2026-10-31T02:29:53.841Z level=INFO msg="europe 2 日目: 16 人"
time=2026-10-31T02:29:56.966Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T02:29:59.696Z level=INFO msg="asia 3 日目: 126 人"
time=2026-10-31T02:30:00.842Z level=INFO msg="user moved" user=sumaxyu.5423 from=europe to=america
time=2026-10-31T02:30:04.715Z level=INFO msg="europe 3 日目: 15 人"
time=2026-10-31T02:30:06.702Z level=INFO msg="america 3 日目: 14 人"
time=2026-10-31T02:30:08.860Z level=INFO msg="asia 4 日目: 126 人"
time=2026-10-31T02:30:12.624Z level=INFO msg="europe 4 日目: 15 人"
time=2026-10-31T02:30:14.904Z level=INFO msg="america 4 日目: 14 人"
time=2026-10-31T02:30:18.255Z level=INFO msg="asia 5 日目: 112 人"
time=2026-10-31T02:30:21.634Z level=INFO msg="europe 5 日目: 15 人"
time=2026-10-31T02:30:23.209Z level=INFO msg="user moved" user=sumaxyu.5423 from=america to=asia
time=2026-10-31T02:30:23.964Z level=INFO msg="america 5 日目: 13 人"
time=2026-10-31T02:30:32.846Z level=INFO msg="america 最終日（グランプリ）: 12 人"
time=2026-10-31T02:30:32.942Z level=INFO msg="europe 最終日（グランプリ）: 14 人"
time=2026-10-31T02:30:33.337Z level=INFO msg="asia 最終日（グランプリ）: 112 人"
time=2026-10-31T02:30:41.616Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 21430,
  "messages": [
    "keru-5910: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kuho-5918: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "marwa_5932: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "miadanke5941: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ivyse-5950: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "torvinchi5953: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "yata5961: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "karoho.5968: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "nama.5975: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "keka.5993: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "zoesol.5996: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "hosomu.6007: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "waulleo6010: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "moissu_6021: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
31.8%	359	0	0	GET	/api/notifications	3391.691	9.448	15.138	15.144	175
16.0%	448	42	0	POST	/api/posts	1706.534	3.809	10.881	11.545	585
10.2%	2180	1238	0	POST	/api/auth/login	1091.171	0.501	3.310	7.509	122
9.0%	3745	665	0	GET	^/api/images/[^/]+	963.904	0.257	2.455	7.083	113065
7.5%	142	4	0	POST	/api/auth/register	802.894	5.654	17.331	19.259	224
6.9%	195	130	0	PUT	/api/me/profile-image	739.207	3.791	11.539	11.585	9
3.9%	689	27	0	POST	/api/friend-requests	414.046	0.601	3.818	4.523	112
3.5%	478	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	369.372	0.773	3.896	4.623	116
1.9%	2603	0	0	GET	^/api/feeds/[^/]+	200.272	0.077	0.435	0.867	2912
1.5%	942	0	0	GET	/api/memories	164.486	0.175	1.296	7.101	10114
1.4%	567	0	0	GET	/api/posting-window/current	148.552	0.262	1.509	6.280	237
1.1%	174	9	0	POST	^/api/posts/[^/]+/comments$	116.868	0.672	3.957	4.187	322
1.0%	65	0	0	GET	/api/me/export	110.868	1.706	10.982	10.982	1952258
0.7%	84	2	0	POST	^/api/posts/[^/]+/realmojis$	70.966	0.845	3.974	3.974	477
0.5%	71	0	0	PUT	^/api/follows/[^/]+$	53.807	0.758	4.158	4.158	0
0.4%	173	0	0	POST	/api/auth/logout	39.624	0.229	2.782	3.298	0
0.4%	339	7	0	GET	^/api/posts/me/current$	38.479	0.114	0.484	0.556	630
0.3%	321	0	0	GET	/api/users	33.434	0.104	0.432	0.572	237
0.3%	387	0	0	GET	^/api/relationships/[^/]+	27.436	0.071	0.413	0.437	1774
0.2%	417	9	0	GET	^/api/posts/[^/]+$	25.108	0.060	0.243	0.348	609
0.2%	41	0	0	PATCH	^/api/posts/[^/]+$	20.845	0.508	2.568	2.568	651
0.2%	47	0	0	PATCH	/api/me	20.238	0.431	1.935	1.935	289
0.1%	232	0	0	GET	^/api/users/[^/]+$	14.957	0.064	0.369	0.510	204
0.1%	181	0	0	GET	^/api/posts/[^/]+/realmojis$	11.679	0.065	0.290	0.357	437
0.1%	169	0	0	GET	^/api/memories/[^/]+	11.586	0.069	0.506	0.627	559
0.1%	1	0	0	POST	/api/initialize	11.381	11.381	11.381	11.381	90
0.1%	186	0	0	GET	^/api/posts/[^/]+/comments$	10.672	0.057	0.207	0.219	1085
0.1%	17	0	0	DELETE	^/api/posts/[^/]+$	9.468	0.557	2.741	2.741	0
0.1%	18	0	0	POST	/api/admin/posting-window	8.725	0.485	2.084	2.084	176
0.1%	72	3	0	GET	/api/me	6.522	0.091	0.482	0.482	289
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.5%	75	0	0	GET	/api/notifications	634.081	8.454	11.330	11.330	176
19.5%	77	1	0	POST	/api/posts	278.022	3.611	11.399	11.399	614
5.2%	361	0	0	GET	^/api/feeds/[^/]+	74.300	0.206	0.727	1.057	2840
4.6%	397	6	0	GET	^/api/images/[^/]+	65.659	0.165	1.030	1.035	188083
4.5%	85	0	0	POST	/api/auth/login	64.377	0.757	8.348	8.348	249
4.0%	94	5	0	POST	/api/friend-requests	56.465	0.601	4.742	4.742	111
3.4%	62	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	48.601	0.784	4.021	4.021	116
2.3%	136	0	0	GET	/api/memories	32.629	0.240	1.021	1.422	11210
1.8%	8	0	0	GET	/api/me/export	25.898	3.237	10.029	10.029	3067768
1.8%	90	0	0	GET	/api/posting-window/current	25.650	0.285	1.859	1.859	238
1.0%	77	7	0	GET	^/api/posts/[^/]+$	14.817	0.192	1.159	1.159	555
0.9%	63	0	0	GET	^/api/posts/me/current$	12.986	0.206	0.508	0.508	622
0.8%	65	0	0	GET	^/api/relationships/[^/]+	11.994	0.185	0.595	0.595	1452
0.6%	16	0	0	POST	^/api/posts/[^/]+/realmojis$	8.794	0.550	3.176	3.176	418
0.6%	41	0	0	GET	/api/users	8.314	0.203	0.486	0.486	251
0.5%	3	0	0	PUT	^/api/blocks/[^/]+$	7.358	2.453	5.058	5.058	0
0.5%	12	0	0	PATCH	^/api/posts/[^/]+$	6.817	0.568	1.719	1.719	631
0.4%	15	0	0	POST	^/api/posts/[^/]+/comments$	6.343	0.423	2.128	2.128	279
0.4%	32	0	0	GET	^/api/users/[^/]+$	6.174	0.193	0.508	0.508	182
0.4%	14	0	0	POST	/api/auth/register	5.842	0.417	0.937	0.937	228
0.4%	10	0	0	PUT	^/api/follows/[^/]+$	5.390	0.539	3.388	3.388	0
0.3%	14	0	0	POST	/api/auth/logout	4.698	0.336	0.909	0.909	0
0.2%	23	0	0	GET	^/api/posts/[^/]+/realmojis$	3.562	0.155	0.502	0.502	359
0.2%	22	0	0	GET	^/api/posts/[^/]+/comments$	3.279	0.149	0.455	0.455	631
0.2%	8	0	0	PATCH	/api/me	3.178	0.397	0.750	0.750	244
0.2%	5	0	0	DELETE	^/api/posts/[^/]+$	3.123	0.625	1.454	1.454	0
0.1%	4	0	0	PUT	/api/me/profile-image	1.699	0.425	0.688	0.688	0
0.1%	12	0	0	GET	^/api/memories/[^/]+	1.681	0.140	0.222	0.222	918
0.1%	9	2	0	GET	/api/me	1.254	0.139	0.202	0.202	219
0.1%	3	0	0	DELETE	^/api/comments/[^/]+$	0.815	0.272	0.554	0.554	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.3%	62	0	0	GET	/api/notifications	508.941	8.209	9.733	9.733	176
20.2%	58	0	0	POST	/api/posts	232.573	4.010	11.851	11.851	621
7.6%	83	3	0	POST	/api/friend-requests	87.612	1.056	5.418	5.418	112
4.3%	40	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	49.427	1.236	4.038	4.038	116
4.2%	70	0	0	POST	/api/auth/login	48.163	0.688	7.732	7.732	249
3.7%	246	3	0	GET	^/api/images/[^/]+	42.714	0.174	0.953	0.979	187122
3.2%	7	0	0	GET	/api/me/export	36.775	5.254	13.140	13.140	3760281
2.4%	121	0	0	GET	^/api/feeds/[^/]+	27.057	0.224	0.978	1.928	1980
2.3%	98	0	0	GET	/api/memories	26.593	0.271	2.020	2.020	9745
1.8%	69	0	0	GET	/api/posting-window/current	20.855	0.302	1.266	1.266	239
0.9%	46	3	0	GET	^/api/posts/me/current$	10.623	0.231	0.595	0.595	582
0.8%	38	0	0	GET	/api/users	8.803	0.232	0.645	0.645	248
0.7%	42	0	0	GET	^/api/relationships/[^/]+	8.149	0.194	0.529	0.529	1449
0.5%	15	0	0	POST	/api/auth/logout	6.269	0.418	1.743	1.743	0
0.5%	12	0	0	POST	/api/auth/register	5.664	0.472	0.925	0.925	227
0.5%	14	1	0	POST	^/api/posts/[^/]+/realmojis$	5.387	0.385	1.183	1.183	411
0.4%	6	0	0	PUT	^/api/follows/[^/]+$	4.364	0.727	2.975	2.975	0
0.3%	5	0	0	PATCH	/api/me	3.978	0.796	1.415	1.415	282
0.3%	13	0	0	GET	^/api/memories/[^/]+	2.890	0.222	0.632	0.632	590
0.2%	15	2	0	GET	^/api/posts/[^/]+$	2.460	0.164	0.408	0.408	548
0.2%	3	0	0	DELETE	^/api/posts/[^/]+$	2.448	0.816	1.832	1.832	0
0.1%	8	0	0	GET	^/api/users/[^/]+$	1.694	0.212	0.530	0.530	204
0.1%	1	0	0	POST	/api/me/account-deletion	1.339	1.339	1.339	1.339	0
0.1%	7	0	0	GET	^/api/posts/[^/]+/realmojis$	1.334	0.191	0.428	0.428	828
0.1%	4	0	0	PUT	/api/me/profile-image	1.234	0.308	0.477	0.477	0
0.1%	8	1	0	GET	/api/me	1.160	0.145	0.236	0.236	228
0.1%	5	0	0	GET	^/api/posts/[^/]+/comments$	0.732	0.146	0.185	0.185	3629
0.0%	1	0	0	DELETE	^/api/realmojis/[^/]+$	0.298	0.298	0.298	0.298	0
0.0%	2	0	0	PATCH	^/api/posts/[^/]+$	0.236	0.118	0.132	0.132	639
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.105	0.105	0.105	0.105	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 15745 クエリ / 86 種類 / 合計 68.378s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	62.1%	80	42.488	531.10	7841.76	6478.1	141.8	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	9.9%	732	6.784	9.27	68.42	709.1	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	6.1%	2222	4.183	1.88	6.71	12.5	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	5.4%	565	3.671	6.50	306.10	118.2	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	3.0%	1	2.062	2061.67	2061.67	29373.0	1199415.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	2.4%	1	1.656	1655.54	1655.54	7501.0	1.0	select count(*) from comments
7	2.2%	1	1.523	1522.70	1522.70	19813.0	625489.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	1.4%	1	0.926	926.46	926.46	11528.0	529712.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	1.3%	545	0.896	1.64	6.89	32.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.7%	1	0.462	462.43	462.43	29382.0	1.0	select count(*) total, count(*) filter (where replaced_at is null and deleted_at is null) cur, count(*) filter (where deleted_at is not null) del from realmojis
11	0.5%	1	0.352	352.16	352.16	19849.0	1.0	select count(*), count(*) filter (where deleted_at is not null) from posts
12	0.5%	111	0.328	2.95	6.26	40.0	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
13	0.5%	2	0.310	154.91	155.50	213853.5	2515.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.4%	2	0.285	142.44	276.52	20272.5	2506.5	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	0.4%	180	0.250	1.39	5.27	27.4	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 3
SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: d8aa3436a861a445eb5cd1e65b06a226c5365956
Type: cpu
Time: 2026-10-31 11:29:24 JST
Duration: 60.08s, Total samples = 88.07s (146.60%)
Showing nodes accounting for 78.39s, 89.01% of 88.07s total
Dropped 884 nodes (cum <= 0.44s)
Showing top 40 nodes out of 176
      flat  flat%   sum%        cum   cum%
    13.46s 15.28% 15.28%     13.46s 15.28%  crypto/internal/fips140/sha256.blockAVX2
    12.15s 13.80% 29.08%     12.15s 13.80%  runtime.spanClass.sizeclass (inline)
     9.12s 10.36% 39.43%     14.82s 16.83%  image/jpeg.(*decoder).reconstructBlock
     4.77s  5.42% 44.85%     32.52s 36.93%  image/jpeg.(*decoder).processSOS
     3.98s  4.52% 49.37%     16.94s 19.23%  runtime.tryDeferToSpanScan
     3.86s  4.38% 53.75%      6.29s  7.14%  image/jpeg.(*decoder).decodeHuffman
     3.04s  3.45% 57.20%      3.18s  3.61%  image/jpeg.(*decoder).receiveExtend
     2.53s  2.87% 60.08%      2.53s  2.87%  runtime.memmove
     2.37s  2.69% 62.77%      2.37s  2.69%  runtime.memclrNoHeapPointers
     2.20s  2.50% 65.27%      3.10s  3.52%  image/jpeg.idctCols
     1.93s  2.19% 67.46%      2.50s  2.84%  image/jpeg.idctRows
     1.84s  2.09% 69.55%      2.56s  2.91%  image/jpeg.(*decoder).ensureNBits
     1.70s  1.93% 71.48%      7.53s  8.55%  runtime.scanObjectsSmall
     1.56s  1.77% 73.25%     13.37s 15.18%  runtime.scanObject
     1.46s  1.66% 74.91%      1.47s  1.67%  image/jpeg.dctBox (inline)
     1.28s  1.45% 76.36%      1.28s  1.45%  internal/runtime/syscall/linux.Syscall6
     1.12s  1.27% 77.63%      5.40s  6.13%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.09s  1.24% 78.87%      1.09s  1.24%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.93s  1.06% 79.93%      0.93s  1.06%  aeshashbody
     0.91s  1.03% 80.96%     15.31s 17.38%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.70s  0.79% 81.75%      9.11s 10.34%  runtime.scanSpan
     0.66s  0.75% 82.50%      0.71s  0.81%  image/jpeg.(*decoder).readByteStuffedByte
     0.59s  0.67% 83.17%     14.94s 16.96%  crypto/internal/fips140/sha256.(*Digest).Write
     0.59s  0.67% 83.84%      0.59s  0.67%  memeqbody
     0.55s  0.62% 84.47%     17.27s 19.61%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.52s  0.59% 85.06%     16.10s 18.28%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.50s  0.57% 85.63%      0.50s  0.57%  runtime.extractHeapBitsSmall
     0.40s  0.45% 86.08%      0.45s  0.51%  runtime.(*spanInlineMarkBits).tryAcquire
     0.36s  0.41% 86.49%      1.04s  1.18%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.35s   0.4% 86.89%      0.56s  0.64%  example.com/isureal-company/isureal.postLess
     0.35s   0.4% 87.28%      2.12s  2.41%  runtime.mapaccess1_faststr
     0.27s  0.31% 87.59%     13.73s 15.59%  crypto/internal/fips140/sha256.block
     0.24s  0.27% 87.86%      0.73s  0.83%  runtime.scanObjectSmall
     0.22s  0.25% 88.11%     18.70s 21.23%  example.com/isureal-company/isureal.imageSignature
     0.21s  0.24% 88.35%      0.65s  0.74%  database/sql.convertAssignRows
     0.16s  0.18% 88.53%      1.32s  1.50%  runtime.mapassign_faststr
     0.14s  0.16% 88.69%      0.76s  0.86%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.12s  0.14% 88.83%     12.86s 14.60%  runtime.mallocgc
     0.08s 0.091% 88.92%      2.53s  2.87%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.08s 0.091% 89.01%      0.50s  0.57%  runtime.mallocgcSmallNoscan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 76.2% / 100% | 66% | 0.2% | 7 | 132 MB | CPU が飽和 |
| isu2 (2 cpu) | 1.0% / 4% | 0% | 0.2% | 1 | 270 MB |  |
| isu3 (2 cpu) | 1.2% / 3% | 0% | 0.2% | 1 | 271 MB |  |
| isu4 (2 cpu) | 29.7% / 80% | 0% | 17.6% | 13 | 104 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 136% / nginx 10% |
| isu2 | 目立つプロセス無し |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 21% |

## Diff from previous run

スコア: 19275 → 21430 (+2155)

`5948f35..4764fb5`

```
 webapp/go/admin_posting_window_api.go | 54 ++++++++++++++++++++++-------------
 webapp/go/store.go                    | 21 ++++++++++++++
 2 files changed, 55 insertions(+), 20 deletions(-)
```
