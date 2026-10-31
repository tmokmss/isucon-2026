# run 20261031-114137-25f2cad

## Summary

| | |
|---|---|
| score | 27654 |
| status | pass |
| ref | origin/mem-reduce-gogc400 |
| sha | 25f2cada3c82 |
| 所要時間 | 282s |
| 前回比 | +333 (1.2%) vs 20261031-113854-7a0fd2d |
| ベスト比 | +333 (1.2%) vs 20261031-113854-7a0fd2d = 27321 **ベスト更新** |
| note | メモリ削減 + GOGC 400 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	133	131	131	130
europe	7	17	37	37	31	30
america	6	15	24	24	18	16
```

```
{"id":"5173","status":"done","portalStatus":"finished","passed":true,"score":27654,"raw":27654,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:45:04.365Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:45:07.313Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:45:10.323Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:45:13.398Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T02:45:16.393Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T02:45:17.353Z level=INFO msg="user moved" user=toisni5274 from=asia to=america
time=2026-10-31T02:45:19.350Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T02:45:19.385Z level=INFO msg="user moved" user=toisni5274 from=america to=europe
time=2026-10-31T02:45:22.251Z level=INFO msg="user moved" user=monoa.5362 from=europe to=asia
time=2026-10-31T02:45:22.351Z level=INFO msg="asia 3 日目: 133 人"
time=2026-10-31T02:45:25.799Z level=INFO msg="user moved" user=monoa.5362 from=asia to=america
time=2026-10-31T02:45:26.084Z level=INFO msg="europe 3 日目: 37 人"
time=2026-10-31T02:45:26.614Z level=INFO msg="user moved" user=haro_5486 from=asia to=america
time=2026-10-31T02:45:28.637Z level=INFO msg="america 3 日目: 24 人"
time=2026-10-31T02:45:31.750Z level=INFO msg="asia 4 日目: 131 人"
time=2026-10-31T02:45:35.467Z level=INFO msg="europe 4 日目: 37 人"
time=2026-10-31T02:45:37.733Z level=INFO msg="america 4 日目: 24 人"
time=2026-10-31T02:45:40.676Z level=INFO msg="asia 5 日目: 131 人"
time=2026-10-31T02:45:45.300Z level=INFO msg="europe 5 日目: 31 人"
time=2026-10-31T02:45:46.723Z level=INFO msg="america 5 日目: 18 人"
time=2026-10-31T02:45:55.369Z level=INFO msg="europe 最終日（グランプリ）: 30 人"
time=2026-10-31T02:45:55.427Z level=INFO msg="asia 最終日（グランプリ）: 130 人"
time=2026-10-31T02:45:55.439Z level=INFO msg="america 最終日（グランプリ）: 16 人"
time=2026-10-31T02:46:04.311Z level=INFO msg="グランプリ: アジア 5 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 27654,
  "messages": [
    "hina.6189: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "soleoel-6203: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "temax6209: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ulkuse_6217: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "nokaleo-6218: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "yohate_6227: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "orsuya.6233: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "yare-6239: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kafu.6248: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "harese.6258: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "yaro.6263: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ultaeva.6267: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "keormia_6276: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.6%	431	1	0	GET	/api/notifications	3876.304	8.994	14.751	14.751	173
17.3%	516	27	0	POST	/api/posts	1945.322	3.770	10.325	11.217	616
10.5%	4654	676	0	GET	^/api/images/[^/]+	1175.705	0.253	3.005	7.660	107430
9.8%	2152	920	0	POST	/api/auth/login	1096.257	0.509	3.895	8.574	152
5.1%	149	1	0	POST	/api/auth/register	571.651	3.837	17.553	18.347	229
4.9%	201	127	0	PUT	/api/me/profile-image	552.847	2.750	10.559	12.645	8
4.1%	807	30	0	POST	/api/friend-requests	463.348	0.574	5.113	9.383	112
3.2%	485	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	360.103	0.742	4.166	6.374	116
1.6%	2806	0	0	GET	^/api/feeds/[^/]+	181.405	0.065	0.312	2.232	2930
1.6%	666	0	0	GET	/api/posting-window/current	177.542	0.267	1.460	2.628	237
1.2%	180	14	0	POST	^/api/posts/[^/]+/realmojis$	134.282	0.746	2.790	3.420	429
1.2%	1190	0	0	GET	/api/memories	132.546	0.111	0.905	4.155	10690
1.0%	88	0	0	GET	/api/me/export	114.595	1.302	9.567	9.567	2096565
0.7%	104	5	0	POST	^/api/posts/[^/]+/comments$	73.680	0.708	3.683	4.865	341
0.6%	72	0	0	PUT	^/api/follows/[^/]+$	63.855	0.887	8.312	8.312	0
0.3%	396	15	0	GET	^/api/posts/me/current$	34.731	0.088	0.927	1.050	625
0.3%	228	0	0	POST	/api/auth/logout	33.412	0.147	3.295	3.696	0
0.3%	504	10	0	GET	^/api/posts/[^/]+$	29.039	0.058	0.209	0.244	617
0.2%	364	0	0	GET	/api/users	27.853	0.077	0.448	0.674	245
0.2%	453	0	0	GET	^/api/relationships/[^/]+	23.595	0.052	0.197	0.358	1488
0.2%	61	0	0	PATCH	/api/me	18.595	0.305	2.559	2.559	276
0.1%	41	1	0	PATCH	^/api/posts/[^/]+$	14.591	0.356	1.748	1.748	630
0.1%	29	0	0	DELETE	^/api/realmojis/[^/]+$	14.391	0.496	3.467	3.467	0
0.1%	238	1	0	GET	^/api/users/[^/]+$	13.633	0.057	0.215	0.242	200
0.1%	1	0	0	POST	/api/initialize	12.546	12.546	12.546	12.546	90
0.1%	219	0	0	GET	^/api/memories/[^/]+	12.295	0.056	0.688	1.034	641
0.1%	23	0	0	DELETE	^/api/posts/[^/]+$	12.254	0.533	1.159	1.159	0
0.1%	231	0	0	GET	^/api/posts/[^/]+/realmojis$	12.059	0.052	0.192	0.257	621
0.1%	183	0	0	GET	^/api/posts/[^/]+/comments$	9.576	0.052	0.149	0.158	590
0.1%	18	0	0	POST	/api/admin/posting-window	6.421	0.357	1.991	1.991	177
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.0%	99	0	0	GET	/api/notifications	817.902	8.262	10.549	10.549	177
17.1%	121	1	0	POST	/api/posts	412.183	3.406	10.802	10.859	627
11.3%	35	2	0	POST	/api/auth/register	272.735	7.792	18.945	18.945	216
5.9%	185	0	0	POST	/api/auth/login	141.119	0.763	6.195	6.940	252
4.6%	544	0	0	GET	^/api/feeds/[^/]+	111.628	0.205	0.857	1.712	2320
4.5%	189	5	0	POST	/api/friend-requests	107.485	0.569	5.121	5.650	113
4.2%	704	6	0	GET	^/api/images/[^/]+	101.714	0.144	0.947	1.235	176970
3.0%	85	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	71.956	0.847	8.193	8.193	116
2.4%	218	0	0	GET	/api/memories	58.269	0.267	1.417	9.935	9498
1.6%	146	0	0	GET	/api/posting-window/current	37.466	0.257	1.204	1.413	239
1.5%	17	0	0	GET	/api/me/export	36.905	2.171	10.948	10.948	2767258
1.2%	21	0	0	PUT	^/api/follows/[^/]+$	27.995	1.333	4.848	4.848	0
1.1%	12	0	0	PUT	/api/me/profile-image	25.587	2.132	6.193	6.193	0
1.0%	37	2	0	POST	^/api/posts/[^/]+/comments$	25.100	0.678	2.610	2.610	301
0.8%	86	7	0	GET	^/api/posts/[^/]+$	18.359	0.213	1.424	1.424	565
0.7%	28	2	0	POST	^/api/posts/[^/]+/realmojis$	17.592	0.628	2.676	2.676	426
0.7%	92	2	0	GET	^/api/posts/me/current$	16.589	0.180	0.485	0.485	614
0.7%	32	0	0	POST	/api/auth/logout	16.382	0.512	1.967	1.967	0
0.6%	94	0	0	GET	^/api/relationships/[^/]+	14.468	0.154	0.487	0.487	1369
0.6%	80	0	0	GET	/api/users	13.818	0.173	0.629	0.629	241
0.5%	15	0	0	PATCH	/api/me	11.469	0.765	2.383	2.383	279
0.4%	42	1	0	GET	^/api/posts/[^/]+/realmojis$	9.921	0.236	1.505	1.505	511
0.4%	50	0	0	GET	^/api/users/[^/]+$	9.675	0.194	0.791	0.791	203
0.4%	36	0	0	GET	^/api/memories/[^/]+	8.724	0.242	1.178	1.178	1047
0.2%	24	0	0	GET	^/api/posts/[^/]+/comments$	5.591	0.233	0.789	0.789	767
0.2%	6	0	0	DELETE	^/api/posts/[^/]+$	4.553	0.759	2.267	2.267	0
0.2%	17	2	0	GET	/api/me	4.163	0.245	0.593	0.593	254
0.1%	9	0	0	PATCH	^/api/posts/[^/]+$	3.342	0.371	0.904	0.904	630
0.1%	4	0	0	DELETE	^/api/comments/[^/]+$	1.293	0.323	0.825	0.825	0
0.1%	3	0	0	DELETE	^/api/realmojis/[^/]+$	1.210	0.403	0.525	0.525	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.4%	79	0	0	GET	/api/notifications	627.996	7.949	9.315	9.315	175
15.8%	75	1	0	POST	/api/posts	245.209	3.269	10.551	10.551	637
9.1%	21	1	0	POST	/api/auth/register	142.218	6.772	18.630	18.630	219
5.6%	115	0	0	POST	/api/auth/login	87.586	0.762	5.992	8.225	250
5.4%	75	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	83.823	1.118	5.547	5.547	116
4.5%	424	2	0	GET	^/api/images/[^/]+	70.537	0.166	0.937	0.945	155969
3.8%	102	5	0	POST	/api/friend-requests	59.576	0.584	2.632	3.192	111
3.4%	228	0	0	GET	^/api/feeds/[^/]+	52.273	0.229	1.282	2.397	2453
2.2%	139	0	0	GET	/api/memories	33.663	0.242	1.164	1.205	9318
1.9%	97	0	0	GET	/api/posting-window/current	30.039	0.310	1.730	1.730	240
1.6%	24	0	0	POST	/api/auth/logout	25.460	1.061	7.089	7.089	0
1.3%	8	0	0	GET	/api/me/export	19.861	2.483	6.354	6.354	2732528
1.0%	60	3	0	GET	^/api/posts/me/current$	14.829	0.247	0.795	0.795	614
0.8%	61	0	0	GET	^/api/relationships/[^/]+	11.953	0.196	1.054	1.054	1489
0.5%	44	0	0	GET	/api/users	7.389	0.168	0.522	0.522	233
0.3%	4	1	0	POST	^/api/posts/[^/]+/realmojis$	5.260	1.315	3.551	3.551	417
0.3%	23	0	0	GET	^/api/users/[^/]+$	4.464	0.194	0.555	0.555	171
0.3%	7	0	0	PATCH	^/api/posts/[^/]+$	4.297	0.614	2.056	2.056	644
0.3%	25	0	0	GET	^/api/memories/[^/]+	4.063	0.163	0.567	0.567	1209
0.2%	7	0	0	PATCH	/api/me	3.772	0.539	2.215	2.215	266
0.2%	11	1	0	POST	^/api/posts/[^/]+/comments$	3.733	0.339	0.681	0.681	251
0.2%	21	1	0	GET	^/api/posts/[^/]+$	3.650	0.174	0.291	0.291	576
0.2%	9	0	0	PUT	^/api/follows/[^/]+$	2.722	0.302	1.115	1.115	0
0.2%	5	0	0	PUT	/api/me/profile-image	2.429	0.486	1.486	1.486	0
0.1%	2	0	0	DELETE	^/api/posts/[^/]+$	2.051	1.026	1.453	1.453	0
0.1%	14	0	0	GET	^/api/posts/[^/]+/realmojis$	2.020	0.144	0.211	0.211	668
0.1%	12	0	0	GET	^/api/posts/[^/]+/comments$	1.978	0.165	0.302	0.302	411
0.1%	9	1	0	GET	/api/me	1.170	0.130	0.213	0.213	243
0.0%	2	0	0	PUT	^/api/blocks/[^/]+$	0.605	0.302	0.501	0.501	0
0.0%	1	0	0	POST	/api/me/account-deletion	0.443	0.443	0.443	0.443	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 20219 クエリ / 80 種類 / 合計 101.836s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	51.6%	113	52.521	464.79	7687.36	7010.0	143.4	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	28.0%	977	28.563	29.24	474.19	679.3	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	5.2%	973	5.258	5.40	212.90	115.3	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	5.1%	2975	5.219	1.75	12.09	11.4	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	2.2%	1	2.273	2273.34	2273.34	29457.0	1199607.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	1.6%	686	1.640	2.39	84.17	33.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
7	1.5%	1	1.546	1546.01	1546.01	20070.0	625853.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	1.1%	1	1.070	1070.13	1070.13	11528.0	529842.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	0.5%	195	0.535	2.74	17.93	36.5	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.3%	2	0.340	169.82	181.36	215803.5	2539.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.3%	2	0.265	132.50	256.52	16770.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.2%	1049	0.231	0.22	4.10	22.5	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
13	0.2%	1	0.226	226.13	226.13	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
14	0.2%	1	0.203	202.62	202.62	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
15	0.2%	144	0.201	1.40	3.31	29.3	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)

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
Build ID: c2e728aaf4b1a72d9946415195c2f6c5010290a7
Type: cpu
Time: 2026-10-31 11:44:47 JST
Duration: 60.11s, Total samples = 88.43s (147.10%)
Showing nodes accounting for 78.11s, 88.33% of 88.43s total
Dropped 813 nodes (cum <= 0.44s)
Showing top 40 nodes out of 167
      flat  flat%   sum%        cum   cum%
    19.83s 22.42% 22.42%     19.83s 22.42%  crypto/internal/fips140/sha256.blockAVX2
    10.88s 12.30% 34.73%     17.56s 19.86%  image/jpeg.(*decoder).reconstructBlock
     5.04s  5.70% 40.43%     34.40s 38.90%  image/jpeg.(*decoder).processSOS
     4.32s  4.89% 45.31%      7.02s  7.94%  image/jpeg.(*decoder).decodeHuffman
     3.65s  4.13% 49.44%      3.82s  4.32%  image/jpeg.(*decoder).receiveExtend
     2.80s  3.17% 52.61%      2.80s  3.17%  runtime.memmove
     2.51s  2.84% 55.44%      3.43s  3.88%  image/jpeg.idctCols
     2.34s  2.65% 58.09%      3.10s  3.51%  image/jpeg.idctRows
     2.21s  2.50% 60.59%      8.08s  9.14%  runtime.scanObject
     2.04s  2.31% 62.90%      2.85s  3.22%  image/jpeg.(*decoder).ensureNBits
     1.94s  2.19% 65.09%      1.94s  2.19%  runtime.memclrNoHeapPointers
     1.92s  2.17% 67.26%      1.92s  2.17%  internal/runtime/syscall/linux.Syscall6
     1.90s  2.15% 69.41%      2.34s  2.65%  runtime.findObject
     1.68s  1.90% 71.31%      1.68s  1.90%  image/jpeg.dctBox (inline)
     1.52s  1.72% 73.03%      2.74s  3.10%  runtime.tryDeferToSpanScan
     1.40s  1.58% 74.61%      5.70s  6.45%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.24s  1.40% 76.01%      1.24s  1.40%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.15s  1.30% 77.32%      1.15s  1.30%  aeshashbody
     1.13s  1.28% 78.59%     22.31s 25.23%  crypto/internal/fips140/sha256.(*Digest).checkSum
     1.08s  1.22% 79.81%     21.85s 24.71%  crypto/internal/fips140/sha256.(*Digest).Write
     0.90s  1.02% 80.83%     25.46s 28.79%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.86s  0.97% 81.80%      0.86s  0.97%  runtime.spanClass.sizeclass (inline)
     0.78s  0.88% 82.69%     23.54s 26.62%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.76s  0.86% 83.55%      0.79s  0.89%  image/jpeg.(*decoder).readByteStuffedByte
     0.48s  0.54% 84.09%      0.48s  0.54%  memeqbody
     0.45s  0.51% 84.60%      0.45s  0.51%  example.com/isureal-company/isureal.(*UserRow).active
     0.45s  0.51% 85.11%      1.83s  2.07%  runtime.mapaccess1_faststr
     0.40s  0.45% 85.56%      1.29s  1.46%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.40s  0.45% 86.01%      0.53s   0.6%  example.com/isureal-company/isureal.postLess
     0.37s  0.42% 86.43%     27.35s 30.93%  example.com/isureal-company/isureal.imageSignature
     0.31s  0.35% 86.78%      0.69s  0.78%  runtime.greyobject
     0.27s  0.31% 87.09%      0.89s  1.01%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.27% 87.36%      0.91s  1.03%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.23s  0.26% 87.62%     20.06s 22.68%  crypto/internal/fips140/sha256.block
     0.12s  0.14% 87.75%      0.87s  0.98%  runtime.scanSpan
     0.11s  0.12% 87.88%      2.36s  2.67%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s  0.12% 88.00%      1.15s  1.30%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.11s  0.12% 88.13%      1.44s  1.63%  runtime.mapassign_faststr
     0.09s   0.1% 88.23%      0.47s  0.53%  crypto/internal/fips140/hmac.(*HMAC).Write
     0.09s   0.1% 88.33%      6.62s  7.49%  image/jpeg.idct
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 74.3% / 100% | 59% | 0.1% | 8 | 161 MB | CPU が飽和 |
| isu2 (2 cpu) | 1.8% / 45% | 0% | 0.6% | 4 | 428 MB |  |
| isu3 (2 cpu) | 1.4% / 5% | 0% | 0.3% | 2 | 264 MB |  |
| isu4 (2 cpu) | 30.8% / 81% | 0% | 16.9% | 29 | 103 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 134% / nginx 10% |
| isu2 | isureal 1% |
| isu3 | isureal 1% |
| isu4 | postgres 25% |

## Diff from previous run

スコア: 27321 → 27654 (+333)

`7a0fd2d..25f2cad`

```
 conf/isu1/etc/systemd/system/isureal-go.service | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)
```
