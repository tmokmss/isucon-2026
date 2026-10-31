# run 20261031-113854-7a0fd2d

## Summary

| | |
|---|---|
| score | 27321 |
| status | pass |
| ref | origin/mem-reduce |
| sha | 7a0fd2db2c37 |
| 所要時間 | 122s |
| 前回比 | +10249 (60.0%) vs 20261031-113441-6a7a33a |
| ベスト比 | +2520 (10.2%) vs 20261031-103811-cef5fe3 = 24801 **ベスト更新** |
| note | メモリ削減: 行をまとめて確保・文字列の共有（GOGC 200 のまま） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	142	142	145	144
europe	7	16	29	29	22	21
america	6	12	19	19	13	12
```

```
{"id":"4793","status":"done","portalStatus":"finished","passed":true,"score":27321,"raw":27321,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:39:41.143Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:39:44.049Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:39:47.080Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:39:50.063Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T02:39:53.134Z level=INFO msg="europe 2 日目: 16 人"
time=2026-10-31T02:39:56.187Z level=INFO msg="america 2 日目: 12 人"
time=2026-10-31T02:39:59.113Z level=INFO msg="asia 3 日目: 142 人"
time=2026-10-31T02:40:02.644Z level=INFO msg="europe 3 日目: 29 人"
time=2026-10-31T02:40:05.463Z level=INFO msg="america 3 日目: 19 人"
time=2026-10-31T02:40:08.337Z level=INFO msg="asia 4 日目: 142 人"
time=2026-10-31T02:40:11.817Z level=INFO msg="europe 4 日目: 29 人"
time=2026-10-31T02:40:14.303Z level=INFO msg="america 4 日目: 19 人"
time=2026-10-31T02:40:17.260Z level=INFO msg="asia 5 日目: 145 人"
time=2026-10-31T02:40:22.272Z level=INFO msg="europe 5 日目: 22 人"
time=2026-10-31T02:40:23.192Z level=INFO msg="america 5 日目: 13 人"
time=2026-10-31T02:40:32.223Z level=INFO msg="asia 最終日（グランプリ）: 144 人"
time=2026-10-31T02:40:32.301Z level=INFO msg="america 最終日（グランプリ）: 12 人"
time=2026-10-31T02:40:32.387Z level=INFO msg="europe 最終日（グランプリ）: 21 人"
time=2026-10-31T02:40:41.038Z level=INFO msg="グランプリ: アジア 3 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 27321,
  "messages": [
    "rutean-6175: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "anyo6178: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "rohanoa-6183: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "nutsunoa_6190: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "vinka6200: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "moritor6206: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "namiku_6214: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "samax.6219: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "evaan.6228: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "solha_6237: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "shishini_6240: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "danneleo6248: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "benivynu.6258: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "liahi_6260: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "benna5122: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.8%	446	0	0	GET	/api/notifications	4043.835	9.067	14.984	15.052	174
16.5%	540	59	0	POST	/api/posts	1919.655	3.555	9.818	10.148	568
9.7%	2495	1218	0	POST	/api/auth/login	1121.722	0.450	4.054	7.744	139
9.1%	4806	679	0	GET	^/api/images/[^/]+	1053.218	0.219	3.086	6.887	111193
6.3%	161	2	0	POST	/api/auth/register	737.197	4.579	14.999	18.774	227
4.9%	201	127	0	PUT	/api/me/profile-image	569.296	2.832	10.465	11.077	8
4.1%	854	32	0	POST	/api/friend-requests	479.227	0.561	4.281	6.976	112
3.7%	511	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	428.868	0.839	5.864	7.354	116
1.7%	3183	0	0	GET	^/api/feeds/[^/]+	203.004	0.064	0.284	1.377	2459
1.7%	255	15	0	POST	^/api/posts/[^/]+/realmojis$	193.095	0.757	2.928	3.118	435
1.4%	712	0	0	GET	/api/posting-window/current	167.562	0.235	1.501	2.414	237
1.2%	1251	0	0	GET	/api/memories	134.346	0.107	1.208	2.451	10461
0.9%	82	0	0	GET	/api/me/export	107.707	1.313	9.654	9.654	1979508
0.6%	98	0	0	PUT	^/api/follows/[^/]+$	67.308	0.687	5.103	5.103	0
0.4%	66	2	0	POST	^/api/posts/[^/]+/comments$	46.652	0.707	2.657	2.657	373
0.4%	228	0	0	POST	/api/auth/logout	42.849	0.188	1.931	6.655	0
0.3%	583	13	0	GET	^/api/posts/[^/]+$	34.594	0.059	0.235	0.305	602
0.3%	405	11	0	GET	^/api/posts/me/current$	31.908	0.079	0.381	0.846	618
0.3%	402	0	0	GET	/api/users	29.296	0.073	0.346	0.853	232
0.2%	51	3	0	PATCH	^/api/posts/[^/]+$	25.376	0.498	2.057	2.057	590
0.2%	495	0	0	GET	^/api/relationships/[^/]+	22.262	0.045	0.258	0.366	1446
0.2%	51	0	0	DELETE	^/api/realmojis/[^/]+$	21.324	0.418	1.242	1.242	0
0.1%	317	0	0	GET	^/api/posts/[^/]+/realmojis$	17.216	0.054	0.179	0.264	487
0.1%	295	0	0	GET	^/api/users/[^/]+$	16.438	0.056	0.233	0.279	191
0.1%	22	0	0	DELETE	^/api/posts/[^/]+$	15.427	0.701	4.351	4.351	0
0.1%	55	0	0	PATCH	/api/me	14.658	0.267	2.387	2.387	266
0.1%	1	0	0	POST	/api/initialize	12.375	12.375	12.375	12.375	90
0.1%	172	0	0	GET	^/api/posts/[^/]+/comments$	11.931	0.069	0.263	0.266	258
0.1%	221	0	0	GET	^/api/memories/[^/]+	11.120	0.050	0.455	1.266	675
0.1%	14	0	0	PUT	^/api/admin/users/[^/]+/certification$	9.582	0.684	7.748	7.748	0
…（上位 30 / 全 44 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.5%	85	1	0	GET	/api/notifications	704.513	8.288	10.745	10.745	177
13.2%	92	2	0	POST	/api/posts	254.653	2.768	9.800	9.800	628
11.2%	27	1	0	POST	/api/auth/register	216.921	8.034	19.703	19.703	219
6.4%	153	0	0	POST	/api/auth/login	123.557	0.808	5.772	7.058	250
5.6%	522	0	0	GET	^/api/feeds/[^/]+	107.950	0.207	0.834	1.076	2184
5.2%	87	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	100.295	1.153	7.263	7.263	116
4.9%	638	4	0	GET	^/api/images/[^/]+	93.843	0.147	0.954	1.272	177488
3.0%	123	4	0	POST	/api/friend-requests	57.804	0.470	3.506	5.164	112
3.0%	184	0	0	GET	/api/memories	57.736	0.314	4.721	9.699	13636
1.4%	18	1	0	POST	^/api/posts/[^/]+/realmojis$	27.824	1.546	5.054	5.054	435
1.4%	110	0	0	GET	/api/posting-window/current	26.911	0.245	1.301	1.404	239
0.9%	9	0	0	GET	/api/me/export	16.904	1.878	4.005	4.005	2492240
0.8%	28	0	0	POST	/api/auth/logout	15.198	0.543	5.022	5.022	0
0.8%	16	2	0	POST	^/api/posts/[^/]+/comments$	15.032	0.940	4.561	4.561	268
0.8%	81	0	0	GET	^/api/relationships/[^/]+	14.814	0.183	0.727	0.727	1296
0.7%	65	4	0	GET	^/api/posts/[^/]+$	13.808	0.212	0.780	0.780	584
0.7%	75	1	0	GET	^/api/posts/me/current$	13.608	0.181	0.778	0.778	632
0.5%	50	0	0	GET	^/api/users/[^/]+$	9.450	0.189	0.622	0.622	200
0.5%	53	0	0	GET	/api/users	9.104	0.172	0.868	0.868	234
0.5%	10	0	0	PATCH	/api/me	8.911	0.891	2.955	2.955	251
0.4%	8	0	0	PUT	/api/me/profile-image	7.698	0.962	4.668	4.668	0
0.3%	31	0	0	GET	^/api/posts/[^/]+/realmojis$	6.566	0.212	0.383	0.383	328
0.3%	30	0	0	GET	^/api/memories/[^/]+	6.400	0.213	1.321	1.321	1038
0.3%	29	0	0	GET	^/api/posts/[^/]+/comments$	6.031	0.208	0.388	0.388	435
0.3%	15	0	0	PUT	^/api/follows/[^/]+$	6.000	0.400	1.203	1.203	0
0.1%	4	0	0	DELETE	^/api/posts/[^/]+$	2.638	0.659	1.394	1.394	0
0.1%	13	2	0	GET	/api/me	1.930	0.148	0.241	0.241	241
0.1%	7	1	0	PATCH	^/api/posts/[^/]+$	1.557	0.222	0.314	0.314	571
0.0%	1	0	0	POST	/api/me/account-deletion	0.912	0.912	0.912	0.912	0
0.0%	2	0	0	PUT	^/api/blocks/[^/]+$	0.744	0.372	0.638	0.638	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.6%	61	1	0	GET	/api/notifications	496.103	8.133	9.320	9.320	177
14.2%	62	0	0	POST	/api/posts	181.763	2.932	10.192	10.192	614
10.4%	18	0	0	POST	/api/auth/register	133.564	7.420	19.657	19.657	229
6.5%	75	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	83.093	1.108	6.033	6.033	115
5.8%	482	3	0	GET	^/api/images/[^/]+	74.940	0.155	1.008	1.011	110071
5.5%	112	0	0	POST	/api/auth/login	70.560	0.630	7.158	7.271	256
3.9%	83	1	0	POST	/api/friend-requests	50.519	0.609	4.457	4.457	114
3.6%	140	0	0	GET	/api/memories	45.844	0.327	3.088	5.875	18595
2.5%	163	0	0	GET	^/api/feeds/[^/]+	31.852	0.195	0.651	0.759	2764
1.5%	71	0	0	GET	/api/posting-window/current	19.532	0.275	1.551	1.551	240
1.3%	5	0	0	GET	/api/me/export	16.051	3.210	6.273	6.273	3568759
0.9%	9	3	0	POST	^/api/posts/[^/]+/realmojis$	11.809	1.312	4.418	4.418	334
0.9%	51	3	0	GET	^/api/posts/me/current$	11.115	0.218	0.930	0.930	576
0.6%	47	0	0	GET	^/api/relationships/[^/]+	7.833	0.167	0.547	0.547	2235
0.6%	12	1	0	POST	^/api/posts/[^/]+/comments$	7.363	0.614	3.154	3.154	287
0.5%	22	0	0	POST	/api/auth/logout	6.383	0.290	1.235	1.235	0
0.5%	36	0	0	GET	/api/users	6.167	0.171	0.474	0.474	235
0.4%	26	3	0	GET	^/api/posts/[^/]+$	4.500	0.173	0.339	0.339	549
0.3%	21	0	0	GET	^/api/memories/[^/]+	3.757	0.179	0.660	0.660	1067
0.3%	5	0	0	PUT	^/api/follows/[^/]+$	3.637	0.727	3.017	3.017	0
0.3%	5	0	0	PUT	/api/me/profile-image	3.250	0.650	2.150	2.150	0
0.2%	16	0	0	GET	^/api/users/[^/]+$	2.629	0.164	0.238	0.238	214
0.2%	3	0	0	DELETE	^/api/posts/[^/]+$	2.550	0.850	1.179	1.179	0
0.2%	3	0	0	PATCH	^/api/posts/[^/]+$	2.165	0.722	1.573	1.573	626
0.1%	2	0	0	PUT	^/api/blocks/[^/]+$	1.468	0.734	1.318	1.318	0
0.1%	9	0	0	GET	^/api/posts/[^/]+/realmojis$	1.377	0.153	0.206	0.206	386
0.1%	9	1	0	GET	/api/me	1.333	0.148	0.312	0.312	252
0.1%	8	0	0	GET	^/api/posts/[^/]+/comments$	1.124	0.141	0.172	0.172	132
0.1%	1	0	0	POST	/api/me/account-deletion	0.758	0.758	0.758	0.758	0
0.0%	3	0	0	DELETE	^/api/comments/[^/]+$	0.554	0.185	0.252	0.252	0
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 20532 クエリ / 81 種類 / 合計 80.860s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	51.3%	96	41.518	432.48	4896.16	6553.6	135.2	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	22.0%	996	17.772	17.84	348.28	620.3	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	7.5%	960	6.086	6.34	305.28	123.2	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	6.8%	3075	5.471	1.78	22.24	11.6	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	2.8%	1	2.304	2304.14	2304.14	29396.0	1199412.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	2.1%	1	1.686	1686.29	1686.29	19950.0	625734.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	1.5%	642	1.173	1.83	14.62	32.0	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
8	1.3%	1	1.088	1087.93	1087.93	11528.0	529719.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	0.8%	263	0.674	2.56	7.67	33.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.5%	2	0.416	208.14	253.45	216361.0	2545.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.3%	2	0.267	133.49	258.70	16766.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.3%	1	0.252	251.74	251.74	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
13	0.3%	1	0.218	217.52	217.52	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
14	0.2%	1014	0.198	0.19	3.53	21.5	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
15	0.2%	1542	0.180	0.12	5.13	12.2	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

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
Build ID: 92445c170d3b17eefc782b3770651d859ce24ae1
Type: cpu
Time: 2026-10-31 11:39:22 JST
Duration: 60.07s, Total samples = 85.73s (142.73%)
Showing nodes accounting for 75.71s, 88.31% of 85.73s total
Dropped 870 nodes (cum <= 0.43s)
Showing top 40 nodes out of 164
      flat  flat%   sum%        cum   cum%
    20.25s 23.62% 23.62%     20.25s 23.62%  crypto/internal/fips140/sha256.blockAVX2
    10.84s 12.64% 36.27%     17.22s 20.09%  image/jpeg.(*decoder).reconstructBlock
     5.14s  6.00% 42.26%     32.77s 38.22%  image/jpeg.(*decoder).processSOS
     3.95s  4.61% 46.87%      6.41s  7.48%  image/jpeg.(*decoder).decodeHuffman
     3.32s  3.87% 50.74%      3.42s  3.99%  image/jpeg.(*decoder).receiveExtend
     2.89s  3.37% 54.11%      2.89s  3.37%  runtime.memmove
     2.44s  2.85% 56.96%      3.18s  3.71%  image/jpeg.idctCols
     2.28s  2.66% 59.62%      3.12s  3.64%  image/jpeg.idctRows
     2.10s  2.45% 62.07%      2.10s  2.45%  runtime.memclrNoHeapPointers
     1.92s  2.24% 64.31%      2.55s  2.97%  image/jpeg.(*decoder).ensureNBits
     1.66s  1.94% 66.24%      6.77s  7.90%  runtime.scanObject
     1.60s  1.87% 68.11%      2.23s  2.60%  runtime.findObject
     1.58s  1.84% 69.95%      1.58s  1.84%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.56s  1.82% 71.77%      1.57s  1.83%  image/jpeg.dctBox (inline)
     1.50s  1.75% 73.52%      5.75s  6.71%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.40s  1.63% 75.15%      1.40s  1.63%  internal/runtime/syscall/linux.Syscall6
     1.35s  1.57% 76.73%      2.23s  2.60%  runtime.tryDeferToSpanScan
     1.24s  1.45% 78.18%     22.94s 26.76%  crypto/internal/fips140/sha256.(*Digest).checkSum
     1.10s  1.28% 79.46%     22.32s 26.04%  crypto/internal/fips140/sha256.(*Digest).Write
     1.03s  1.20% 80.66%      1.03s  1.20%  aeshashbody
     0.81s  0.94% 81.61%     24.23s 28.26%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.65s  0.76% 82.36%     25.90s 30.21%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.61s  0.71% 83.07%      0.61s  0.71%  runtime.spanClass.sizeclass (inline)
     0.54s  0.63% 83.70%      0.60s   0.7%  image/jpeg.(*decoder).readByteStuffedByte
     0.49s  0.57% 84.28%      0.62s  0.72%  example.com/isureal-company/isureal.postLess
     0.48s  0.56% 84.84%      1.77s  2.06%  runtime.mapaccess1_faststr
     0.40s  0.47% 85.30%      0.71s  0.83%  runtime.greyobject
     0.32s  0.37% 85.68%     27.58s 32.17%  example.com/isureal-company/isureal.imageSignature
     0.31s  0.36% 86.04%      0.46s  0.54%  database/sql.convertAssignRows
     0.30s  0.35% 86.39%      1.14s  1.33%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.28s  0.33% 86.71%     20.53s 23.95%  crypto/internal/fips140/sha256.block
     0.27s  0.31% 87.03%      0.86s  1.00%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.22s  0.26% 87.29%      0.81s  0.94%  runtime.mallocgcSmallScanNoHeader
     0.15s  0.17% 87.46%      0.43s   0.5%  runtime.scanObjectsSmall
     0.14s  0.16% 87.62%      0.44s  0.51%  crypto/internal/fips140/hmac.(*HMAC).Write
     0.14s  0.16% 87.79%      1.76s  2.05%  runtime.mapassign_faststr
     0.13s  0.15% 87.94%      0.79s  0.92%  runtime.scanSpan
     0.12s  0.14% 88.08%      2.90s  3.38%  runtime.mallocgc
     0.10s  0.12% 88.20%      2.23s  2.60%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.12% 88.31%      1.11s  1.29%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 72.2% / 100% | 59% | 0.0% | 7 | 153 MB | CPU が飽和 |
| isu2 (2 cpu) | 1.2% / 5% | 0% | 0.2% | 0 | 505 MB |  |
| isu3 (2 cpu) | 1.5% / 4% | 0% | 0.3% | 3 | 260 MB |  |
| isu4 (2 cpu) | 28.6% / 86% | 0% | 15.8% | 17 | 111 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 130% / nginx 10% |
| isu2 | isureal 1% |
| isu3 | isureal 1% |
| isu4 | postgres 23% |

## Diff from previous run

スコア: 17072 → 27321 (+10249)

`6a7a33a..7a0fd2d`

```
 webapp/go/store.go | 76 +++++++++++++++++++++++++++++++++++++++++++++++++++---
 1 file changed, 73 insertions(+), 3 deletions(-)
```
