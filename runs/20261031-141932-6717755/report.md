# run 20261031-141932-6717755

## Summary

| | |
|---|---|
| score | 132853 |
| status | pass |
| ref | jpeg-peer-asia2-2 |
| sha | 6717755df7e0 |
| 所要時間 | 367s |
| 前回比 | +9302 (7.5%) vs 20261031-140544-247167c |
| ベスト比 | +9302 (7.5%) vs 20261031-140544-247167c = 123551 **ベスト更新** |
| note | JPEG の検証を asia2 にも逃がす（main e42c64c の上） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	150	434	1031	1030
europe	7	17	48	138	184	184
america	6	17	48	129	132	131
```

```
{"id":"17214","status":"done","portalStatus":"finished","passed":true,"score":132853,"raw":132853,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:24:03.025Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:24:05.997Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:24:08.995Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:24:12.019Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T05:24:14.993Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T05:24:15.065Z level=INFO msg="user moved" user=kulia.5322 from=asia to=europe
time=2026-10-31T05:24:17.994Z level=INFO msg="america 2 日目: 17 人"
time=2026-10-31T05:24:21.025Z level=INFO msg="asia 3 日目: 150 人"
time=2026-10-31T05:24:23.996Z level=INFO msg="europe 3 日目: 48 人"
time=2026-10-31T05:24:24.303Z level=INFO msg="user moved" user=yamo-6126 from=asia to=europe
time=2026-10-31T05:24:27.028Z level=INFO msg="america 3 日目: 48 人"
time=2026-10-31T05:24:30.030Z level=INFO msg="asia 4 日目: 434 人"
time=2026-10-31T05:24:30.261Z level=INFO msg="user moved" user=hoeva_6487 from=america to=asia
time=2026-10-31T05:24:30.553Z level=INFO msg="user moved" user=hoeva_6487 from=asia to=america
time=2026-10-31T05:24:32.617Z level=INFO msg="user moved" user=vinhohi6961 from=asia to=europe
time=2026-10-31T05:24:32.992Z level=INFO msg="user moved" user=muis7216 from=asia to=america
time=2026-10-31T05:24:33.106Z level=INFO msg="europe 4 日目: 138 人"
time=2026-10-31T05:24:33.767Z level=INFO msg="user moved" user=hanoa-7053 from=asia to=america
time=2026-10-31T05:24:34.042Z level=INFO msg="user moved" user=liayuse.6551 from=america to=europe
time=2026-10-31T05:24:34.756Z level=INFO msg="user moved" user=noevamia_7167 from=asia to=america
time=2026-10-31T05:24:35.030Z level=INFO msg="user moved" user=yuorhe-8618 from=asia to=europe
time=2026-10-31T05:24:35.513Z level=INFO msg="user moved" user=raelku.6623 from=america to=asia
time=2026-10-31T05:24:36.190Z level=INFO msg="america 4 日目: 129 人"
time=2026-10-31T05:24:36.593Z level=INFO msg="user moved" user=vinvin.8712 from=asia to=europe
time=2026-10-31T05:24:37.067Z level=INFO msg="user moved" user=yamo-6126 from=europe to=asia
time=2026-10-31T05:24:37.105Z level=INFO msg="user moved" user=raelku.6623 from=asia to=america
time=2026-10-31T05:24:37.224Z level=INFO msg="user moved" user=roka-7810 from=asia to=europe
time=2026-10-31T05:24:37.366Z level=INFO msg="user moved" user=tahika-7938 from=asia to=america
time=2026-10-31T05:24:38.206Z level=INFO msg="user moved" user=shini.7436 from=asia to=america
time=2026-10-31T05:24:39.074Z level=INFO msg="asia 5 日目: 1031 人"
time=2026-10-31T05:24:42.279Z level=INFO msg="europe 5 日目: 184 人"
time=2026-10-31T05:24:45.071Z level=INFO msg="america 5 日目: 132 人"
time=2026-10-31T05:24:54.183Z level=INFO msg="europe 最終日（グランプリ）: 184 人"
time=2026-10-31T05:24:54.473Z level=INFO msg="america 最終日（グランプリ）: 131 人"
time=2026-10-31T05:24:54.662Z level=INFO msg="asia 最終日（グランプリ）: 1030 人"
time=2026-10-31T05:25:03.040Z level=INFO msg="グランプリ: アジア 408 人, ヨーロッパ 97 人, アメリカ 65 人"
{
  "pass": true,
  "score": 132853,
  "messages": [
    "elul_6390: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "soltor.11803: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noahi5930: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kurama13601: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.3%	1676	0	0	GET	/api/notifications	14129.499	8.430	15.710	15.722	167
16.7%	2568	105	0	POST	/api/posts	6489.751	2.527	9.463	11.871	627
11.3%	5727	86	0	POST	/api/friend-requests	4383.687	0.765	5.271	10.460	114
8.5%	15266	511	0	GET	^/api/images/[^/]+	3321.570	0.218	3.644	7.704	145018
5.7%	5994	1309	0	POST	/api/auth/login	2227.239	0.372	5.070	10.814	196
4.3%	2247	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1689.061	0.752	5.578	7.074	116
2.9%	738	136	0	PUT	/api/me/profile-image	1146.021	1.553	8.016	10.698	4
2.2%	599	52	0	POST	^/api/posts/[^/]+/realmojis$	852.617	1.423	7.771	8.610	437
1.9%	1057	1	0	POST	/api/auth/register	746.353	0.706	5.871	9.567	227
1.7%	661	42	0	POST	^/api/posts/[^/]+/comments$	668.924	1.012	6.714	7.059	330
1.5%	734	0	0	PUT	^/api/follows/[^/]+$	595.622	0.811	5.305	6.752	0
1.3%	3701	0	0	GET	/api/memories	494.981	0.134	0.918	3.800	11086
1.2%	511	0	0	PATCH	/api/me	473.506	0.927	5.930	7.378	280
0.9%	298	0	0	GET	/api/me/export	338.169	1.135	6.790	8.128	2047025
0.7%	13880	0	0	GET	^/api/feeds/[^/]+	267.593	0.019	0.207	6.105	2223
0.6%	761	0	0	POST	/api/auth/logout	221.007	0.290	4.084	6.612	0
0.5%	212	8	0	PATCH	^/api/posts/[^/]+$	200.955	0.948	5.824	5.959	629
0.4%	98	0	0	DELETE	^/api/posts/[^/]+$	148.592	1.516	7.784	7.784	0
0.2%	82	0	0	DELETE	^/api/realmojis/[^/]+$	59.451	0.725	4.030	4.030	0
0.1%	1775	0	0	GET	^/api/relationships/[^/]+	51.219	0.029	0.369	1.599	907
0.1%	108	0	0	DELETE	^/api/comments/[^/]+$	50.257	0.465	2.612	2.959	0
0.1%	3686	0	0	GET	/api/posting-window/current	42.823	0.012	0.123	0.385	237
0.1%	2724	0	0	GET	/api/users	40.125	0.015	0.168	1.964	236
0.1%	1973	103	0	GET	^/api/posts/me/current$	38.077	0.019	0.188	1.254	620
0.1%	676	0	0	GET	^/api/memories/[^/]+	36.173	0.054	0.514	0.678	599
0.1%	2147	18	0	GET	^/api/posts/[^/]+$	35.619	0.017	0.158	1.256	609
0.1%	31	0	0	PUT	^/api/blocks/[^/]+$	23.266	0.751	5.438	5.438	0
0.0%	1	0	0	POST	/api/initialize	18.970	18.970	18.970	18.970	90
0.0%	1044	0	0	GET	^/api/users/[^/]+$	15.908	0.015	0.166	0.668	194
0.0%	802	0	0	GET	^/api/posts/[^/]+/comments$	12.753	0.016	0.105	0.939	1828
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.1%	390	1	0	GET	/api/notifications	3110.187	7.975	11.688	11.688	172
12.1%	504	13	0	POST	/api/posts	961.376	1.907	7.178	9.922	632
9.8%	966	8	0	POST	/api/friend-requests	775.382	0.803	5.399	6.574	114
5.6%	573	0	0	POST	/api/auth/login	445.981	0.778	6.470	8.247	246
5.3%	1998	19	0	GET	^/api/images/[^/]+	417.644	0.209	1.095	2.503	200984
4.7%	1847	0	0	GET	^/api/feeds/[^/]+	373.634	0.202	0.713	2.082	2353
4.3%	345	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	344.150	0.998	5.381	6.295	116
2.2%	593	0	0	GET	/api/memories	175.038	0.295	1.974	2.396	11788
1.8%	176	0	0	POST	/api/auth/register	144.803	0.823	5.704	6.161	229
1.8%	711	0	0	GET	/api/posting-window/current	144.330	0.203	0.678	1.069	239
1.6%	100	5	0	POST	^/api/posts/[^/]+/realmojis$	124.412	1.244	8.210	8.650	448
1.2%	427	27	0	GET	^/api/posts/me/current$	97.382	0.228	1.032	1.674	610
1.1%	458	0	0	GET	/api/users	87.879	0.192	0.662	1.198	238
1.1%	100	0	0	POST	/api/auth/logout	85.301	0.853	5.056	6.169	0
1.0%	100	0	0	PUT	/api/me/profile-image	83.293	0.833	3.546	5.119	0
1.0%	32	0	0	GET	/api/me/export	78.673	2.459	6.985	6.985	2771960
1.0%	90	5	0	POST	^/api/posts/[^/]+/comments$	77.304	0.859	5.851	5.851	297
0.9%	126	0	0	PUT	^/api/follows/[^/]+$	70.326	0.558	3.825	4.111	0
0.8%	285	0	0	GET	^/api/relationships/[^/]+	60.933	0.214	0.847	1.674	887
0.7%	57	9	0	PATCH	^/api/posts/[^/]+$	59.238	1.039	4.496	4.496	554
0.6%	271	9	0	GET	^/api/posts/[^/]+$	48.549	0.179	0.854	0.944	590
0.5%	72	0	0	PATCH	/api/me	42.605	0.592	3.918	3.918	276
0.4%	22	0	0	DELETE	^/api/posts/[^/]+$	29.377	1.335	4.477	4.477	0
0.3%	127	1	0	GET	^/api/users/[^/]+$	21.323	0.168	0.629	1.127	189
0.2%	105	2	0	GET	/api/me	18.442	0.176	0.416	0.679	298
0.2%	108	1	0	GET	^/api/posts/[^/]+/realmojis$	18.032	0.167	0.602	0.946	1055
0.2%	86	0	0	GET	^/api/posts/[^/]+/comments$	16.798	0.195	0.627	0.627	1722
0.2%	71	0	0	GET	^/api/memories/[^/]+	14.557	0.205	0.965	0.965	769
0.1%	13	0	0	DELETE	^/api/realmojis/[^/]+$	7.889	0.607	2.073	2.073	0
0.1%	12	0	0	DELETE	^/api/comments/[^/]+$	7.400	0.617	2.740	2.740	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.8%	306	1	0	GET	/api/notifications	2271.640	7.424	9.467	9.632	173
12.9%	406	10	0	POST	/api/posts	757.565	1.866	6.585	13.103	632
9.7%	643	8	0	POST	/api/friend-requests	568.167	0.884	5.412	6.227	114
5.9%	1674	18	0	GET	^/api/images/[^/]+	347.724	0.208	1.086	1.542	179226
5.5%	455	0	0	POST	/api/auth/login	324.938	0.714	5.242	8.877	249
4.2%	1191	0	0	GET	^/api/feeds/[^/]+	247.954	0.208	0.777	2.001	2242
3.6%	230	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	208.295	0.906	5.314	9.322	115
2.8%	481	0	0	GET	/api/memories	162.906	0.339	2.439	4.034	12270
1.8%	562	0	0	GET	/api/posting-window/current	107.107	0.191	0.671	0.980	240
1.4%	69	0	0	PUT	/api/me/profile-image	84.707	1.228	6.981	6.981	0
1.4%	50	6	0	POST	^/api/posts/[^/]+/realmojis$	79.907	1.598	8.405	8.405	465
1.3%	338	27	0	GET	^/api/posts/me/current$	74.175	0.219	1.464	1.927	595
1.2%	68	2	0	POST	^/api/posts/[^/]+/comments$	70.214	1.033	6.328	6.328	320
1.1%	26	0	0	GET	/api/me/export	63.013	2.424	6.762	6.762	2517952
1.1%	97	0	0	PUT	^/api/follows/[^/]+$	61.551	0.635	5.194	5.194	0
1.0%	73	0	0	POST	/api/auth/logout	58.062	0.795	6.015	6.015	0
0.9%	292	0	0	GET	/api/users	54.269	0.186	0.615	0.663	238
0.9%	65	0	0	PATCH	/api/me	52.107	0.802	6.480	6.480	275
0.8%	127	0	0	POST	/api/auth/register	45.609	0.359	2.735	3.263	229
0.7%	182	0	0	GET	^/api/relationships/[^/]+	40.709	0.224	1.120	1.637	976
0.6%	23	0	0	DELETE	^/api/posts/[^/]+$	37.971	1.651	6.021	6.021	0
0.5%	25	0	0	PATCH	^/api/posts/[^/]+$	27.983	1.119	4.560	4.560	648
0.5%	137	9	0	GET	^/api/posts/[^/]+$	27.597	0.201	0.694	0.764	581
0.3%	64	0	0	GET	^/api/memories/[^/]+	15.502	0.242	0.981	0.981	1358
0.3%	83	0	0	GET	^/api/users/[^/]+$	14.991	0.181	1.173	1.173	178
0.2%	75	1	0	GET	/api/me	13.504	0.180	0.597	0.597	296
0.2%	61	0	0	GET	^/api/posts/[^/]+/realmojis$	10.599	0.174	0.515	0.515	820
0.2%	48	0	0	GET	^/api/posts/[^/]+/comments$	9.069	0.189	0.648	0.648	1886
0.1%	10	0	0	DELETE	^/api/comments/[^/]+$	7.161	0.716	2.001	2.001	0
0.1%	5	0	0	DELETE	^/api/realmojis/[^/]+$	5.905	1.181	1.985	1.985	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 81690 クエリ / 82 種類 / 合計 419.812s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	51.5%	356	216.123	607.09	4452.44	1352.9	137.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	16.4%	4818	68.714	14.26	270.02	584.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	10.2%	3207	43.014	13.41	551.79	119.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	10.0%	10184	41.862	4.11	95.36	11.1	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	4.5%	3359	19.044	5.67	253.85	24.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.6%	2	6.735	3367.25	6578.55	237303.5	2890.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.8%	686	3.446	5.02	43.60	28.6	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
8	0.8%	7219	3.150	0.44	222.70	19.9	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
9	0.7%	2735	3.062	1.12	233.21	20.8	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
10	0.5%	1	2.175	2175.50	2175.50	31512.0	1200065.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
11	0.5%	1	1.889	1889.27	1889.27	24949.0	628466.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
12	0.4%	770	1.764	2.29	20.55	20.9	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
13	0.4%	5713	1.671	0.29	218.36	11.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
14	0.3%	2812	1.315	0.47	218.75	8.5	1.0	UPDATE friend_requests SET status=$2, updated_at=$3 WHERE id=$1
15	0.2%	5716	0.894	0.16	216.21	4.1	0.2	UPDATE login_sessions SET active = $3, invalidated_at = $2, reason = $4 WHERE user_id = $1 AND active = $5

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 082fe256161500d08f65bb98ffc073ef439965c8
Type: cpu
Time: 2026-10-31 14:23:38 JST
Duration: 60.01s, Total samples = 36.99s (61.64%)
Showing nodes accounting for 26.43s, 71.45% of 36.99s total
Dropped 940 nodes (cum <= 0.18s)
Showing top 40 nodes out of 302
      flat  flat%   sum%        cum   cum%
        4s 10.81% 10.81%         4s 10.81%  internal/runtime/syscall/linux.Syscall6
     2.93s  7.92% 18.73%      2.93s  7.92%  crypto/internal/fips140/sha256.blockAVX2
     1.77s  4.79% 23.52%      1.77s  4.79%  runtime.memclrNoHeapPointers
     1.73s  4.68% 28.20%      6.44s 17.41%  runtime.scanObject
     1.64s  4.43% 32.63%      1.64s  4.43%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.62s  4.38% 37.01%      2.14s  5.79%  runtime.findObject
     1.36s  3.68% 40.69%      1.36s  3.68%  aeshashbody
     1.35s  3.65% 44.34%      1.35s  3.65%  runtime.memmove
     1.34s  3.62% 47.96%      5.67s 15.33%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.31s  3.54% 51.50%      2.29s  6.19%  runtime.tryDeferToSpanScan
     0.65s  1.76% 53.26%      0.65s  1.76%  runtime.spanClass.sizeclass (inline)
     0.59s  1.60% 54.85%      2.23s  6.03%  runtime.mapaccess1_faststr
     0.45s  1.22% 56.07%      0.45s  1.22%  memeqbody
     0.38s  1.03% 57.10%      0.38s  1.03%  runtime.nextFreeFast (inline)
     0.33s  0.89% 57.99%      0.35s  0.95%  encoding/json.appendString[go.shape.string]
     0.33s  0.89% 58.88%      0.58s  1.57%  runtime.greyobject
     0.32s  0.87% 59.75%      0.46s  1.24%  example.com/isureal-company/isureal.postLess
     0.32s  0.87% 60.61%      0.32s  0.87%  runtime.futex
     0.31s  0.84% 61.45%      0.73s  1.97%  database/sql.convertAssignRows
     0.29s  0.78% 62.23%      1.35s  3.65%  runtime.mallocgcSmallScanNoHeader
     0.25s  0.68% 62.91%      0.28s  0.76%  runtime.spanOf (inline)
     0.24s  0.65% 63.56%      0.31s  0.84%  runtime.(*mspan).writeHeapBitsSmall
     0.23s  0.62% 64.18%      0.52s  1.41%  compress/flate.(*compressor).deflate
     0.22s  0.59% 64.77%      0.22s  0.59%  runtime.(*mspan).divideByElemSize (inline)
     0.21s  0.57% 65.34%      4.66s 12.60%  runtime.mallocgc
     0.20s  0.54% 65.88%      2.79s  7.54%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.20s  0.54% 66.42%      0.20s  0.54%  runtime.typePointers.nextFast (inline)
     0.19s  0.51% 66.94%      0.20s  0.54%  runtime.typePointers.next
     0.18s  0.49% 67.42%      2.01s  5.43%  runtime.mapassign_faststr
     0.17s  0.46% 67.88%      0.20s  0.54%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.17s  0.46% 68.34%      0.24s  0.65%  runtime.(*mspan).markBitsForIndex (inline)
     0.16s  0.43% 68.78%      0.31s  0.84%  internal/runtime/maps.(*Iter).Next
     0.15s  0.41% 69.18%      0.23s  0.62%  runtime.heapArenaOf (inline)
     0.15s  0.41% 69.59%      0.46s  1.24%  runtime.scanObjectsSmall
     0.13s  0.35% 69.94%      2.79s  7.54%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.12s  0.32% 70.26%      0.60s  1.62%  runtime.mapaccess2_faststr
     0.11s   0.3% 70.56%      3.22s  8.71%  crypto/internal/fips140/sha256.(*Digest).Write
     0.11s   0.3% 70.86%      0.24s  0.65%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.11s   0.3% 71.15%      1.70s  4.60%  example.com/isureal-company/isureal.listMemories
     0.11s   0.3% 71.45%      0.30s  0.81%  internal/stringslite.Index
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 47.0% / 100% | 30% | 0.3% | 7 | 119 MB | CPU が飽和 |
| isu2 (2 cpu) | 15.3% / 70% | 0% | 0.6% | 5 | 1813 MB | run queue が積もっている |
| isu3 (2 cpu) | 12.9% / 78% | 0% | 0.8% | 4 | 2512 MB |  |
| isu4 (2 cpu) | 52.3% / 100% | 28% | 19.8% | 23 | 109 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 75% / nginx 12% |
| isu2 | isureal 28% / nginx 1% |
| isu3 | isureal 23% / nginx 1% |
| isu4 | postgres 40% / isureal 19% / kswapd0 1% |

## Diff from previous run

スコア: 123551 → 132853 (+9302)

`247167c..6717755`

```
 conf/isu1/home/isucon/env.sh   |  2 +-
 webapp/go/jpegcheck/huffman.go | 27 ++++++++++++++++++++++++---
 webapp/go/jpegcheck/reader.go  | 32 ++++++++++++++++++++------------
 webapp/go/jpegcheck/scan.go    |  7 ++++---
 webapp/go/offload.go           |  2 +-
 5 files changed, 50 insertions(+), 20 deletions(-)
```
