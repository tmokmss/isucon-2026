# run 20261031-125933-8caa076

## Summary

| | |
|---|---|
| score | 85230 |
| status | pass |
| ref | 8caa076 |
| sha | 8caa076c49d3 |
| 所要時間 | 364s |
| 前回比 | -10463 (-10.9%) vs 20261031-125722-c75a0a9 |
| ベスト比 | -10463 (-10.9%) vs 20261031-125722-c75a0a9 = 95693 |
| note | main-8caa076-control |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	150	428	866	866
europe	7	18	49	126	128	128
america	6	13	31	60	63	63
```

```
{"id":"10907","status":"done","portalStatus":"finished","passed":true,"score":85230,"raw":85230,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:04:11.179Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:04:14.142Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:04:17.141Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:04:20.166Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T04:04:23.141Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T04:04:26.141Z level=INFO msg="america 2 日目: 13 人"
time=2026-10-31T04:04:29.166Z level=INFO msg="asia 3 日目: 150 人"
time=2026-10-31T04:04:32.158Z level=INFO msg="europe 3 日目: 49 人"
time=2026-10-31T04:04:32.757Z level=INFO msg="user moved" user=shivin-6009 from=asia to=europe
time=2026-10-31T04:04:35.144Z level=INFO msg="america 3 日目: 31 人"
time=2026-10-31T04:04:38.211Z level=INFO msg="asia 4 日目: 428 人"
time=2026-10-31T04:04:41.348Z level=INFO msg="europe 4 日目: 126 人"
time=2026-10-31T04:04:41.652Z level=INFO msg="user moved" user=razoeis-6564 from=asia to=europe
time=2026-10-31T04:04:44.348Z level=INFO msg="america 4 日目: 60 人"
time=2026-10-31T04:04:44.807Z level=INFO msg="user moved" user=kuke-7291 from=asia to=europe
time=2026-10-31T04:04:45.265Z level=INFO msg="user moved" user=ruso_6833 from=asia to=america
time=2026-10-31T04:04:46.312Z level=INFO msg="user moved" user=hitatsu-7336 from=asia to=america
time=2026-10-31T04:04:46.502Z level=INFO msg="user moved" user=maxivyki.8029 from=asia to=america
time=2026-10-31T04:04:47.408Z level=INFO msg="asia 5 日目: 866 人"
time=2026-10-31T04:04:50.226Z level=INFO msg="europe 5 日目: 128 人"
time=2026-10-31T04:04:53.414Z level=INFO msg="america 5 日目: 63 人"
time=2026-10-31T04:05:02.340Z level=INFO msg="asia 最終日（グランプリ）: 866 人"
time=2026-10-31T04:05:02.510Z level=INFO msg="america 最終日（グランプリ）: 63 人"
time=2026-10-31T04:05:02.580Z level=INFO msg="europe 最終日（グランプリ）: 128 人"
time=2026-10-31T04:05:11.173Z level=INFO msg="グランプリ: アジア 158 人, ヨーロッパ 4 人, アメリカ 1 人"
{
  "pass": true,
  "score": 85230,
  "messages": [
    "homeya6648: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "hire6259: 投稿 ce5c939b-8220-4a67-83e0-8e8b5484111f に RealMoji を送れません: POST /api/posts/ce5c939b-8220-4a67-83e0-8e8b5484111f/realmojis: 応答がタイムアウトしました",
    "fuzoe_5465: 投稿できません: POST /api/posts: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
30.5%	1368	0	0	GET	/api/notifications	11610.385	8.487	14.988	15.219	167
18.0%	1866	81	0	POST	/api/posts	6834.196	3.662	11.539	22.030	624
13.3%	4452	49	0	POST	/api/friend-requests	5048.679	1.134	10.996	16.154	114
9.4%	11681	927	0	GET	^/api/images/[^/]+	3561.589	0.305	7.059	14.493	123694
7.8%	5537	1892	0	POST	/api/auth/login	2981.255	0.538	7.787	14.265	170
4.7%	1385	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1792.213	1.294	10.859	16.182	116
4.0%	888	1	0	POST	/api/auth/register	1509.987	1.700	8.057	10.132	227
2.0%	619	118	0	PUT	/api/me/profile-image	767.883	1.241	8.011	10.100	4
1.5%	378	20	0	POST	^/api/posts/[^/]+/realmojis$	568.876	1.505	9.324	15.270	469
1.2%	438	18	0	POST	^/api/posts/[^/]+/comments$	473.986	1.082	14.599	15.565	324
1.2%	497	0	0	PUT	^/api/follows/[^/]+$	442.800	0.891	9.593	17.912	0
1.0%	2894	0	0	GET	/api/memories	396.907	0.137	0.997	2.138	11155
1.0%	414	0	0	PATCH	/api/me	383.644	0.927	4.755	12.570	285
0.7%	8447	0	0	GET	^/api/feeds/[^/]+	251.916	0.030	0.344	1.670	1874
0.7%	168	10	0	PATCH	^/api/posts/[^/]+$	248.303	1.478	10.134	11.141	606
0.6%	562	0	0	POST	/api/auth/logout	223.772	0.398	8.083	11.849	0
0.5%	214	0	0	GET	/api/me/export	202.915	0.948	6.021	9.513	1850178
0.4%	76	0	0	DELETE	^/api/posts/[^/]+$	164.564	2.165	15.899	15.899	0
0.3%	2696	0	0	GET	/api/posting-window/current	121.748	0.045	0.676	5.837	237
0.2%	1383	92	0	GET	^/api/posts/me/current$	76.058	0.055	0.663	1.165	609
0.2%	64	0	0	DELETE	^/api/comments/[^/]+$	64.909	1.014	12.289	12.289	0
0.2%	2211	0	0	GET	/api/users	58.412	0.026	0.244	0.620	246
0.1%	42	0	0	DELETE	^/api/realmojis/[^/]+$	42.337	1.008	10.536	10.536	0
0.1%	24	0	0	PUT	^/api/blocks/[^/]+$	40.686	1.695	14.024	14.024	0
0.1%	538	0	0	GET	^/api/memories/[^/]+	27.954	0.052	0.454	0.861	639
0.1%	1185	0	0	GET	^/api/relationships/[^/]+	26.996	0.023	0.293	0.795	917
0.1%	1300	11	0	GET	^/api/posts/[^/]+$	25.832	0.020	0.163	0.796	620
0.0%	13	0	0	PUT	^/api/admin/users/[^/]+/certification$	13.178	1.014	12.460	12.460	0
0.0%	1	0	0	POST	/api/initialize	12.638	12.638	12.638	12.638	90
0.0%	515	2	0	GET	/api/me	11.384	0.022	0.118	0.182	300
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.6%	283	0	0	GET	/api/notifications	2361.755	8.345	11.997	12.492	178
19.1%	351	8	0	POST	/api/posts	1341.033	3.821	11.312	16.859	634
12.8%	681	4	0	POST	/api/friend-requests	903.612	1.327	10.222	14.635	114
6.0%	442	0	0	POST	/api/auth/login	423.966	0.959	10.511	14.579	247
4.7%	200	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	330.798	1.654	11.828	16.090	116
4.5%	1502	16	0	GET	^/api/images/[^/]+	314.539	0.209	1.348	3.489	175848
3.2%	1244	0	0	GET	^/api/feeds/[^/]+	228.520	0.184	0.610	3.061	2072
1.8%	99	8	0	POST	^/api/posts/[^/]+/realmojis$	127.987	1.293	18.679	18.679	437
1.8%	479	0	0	GET	/api/posting-window/current	126.406	0.264	0.730	1.064	239
1.6%	434	0	0	GET	/api/memories	111.503	0.257	1.146	2.375	13630
1.5%	79	0	0	PUT	^/api/follows/[^/]+$	103.401	1.309	7.891	7.891	0
1.1%	74	0	0	POST	/api/auth/logout	77.529	1.048	9.290	9.290	0
1.0%	123	0	0	POST	/api/auth/register	71.853	0.584	2.275	2.362	228
1.0%	322	0	0	GET	/api/users	68.004	0.211	0.674	0.931	250
0.8%	263	26	0	GET	^/api/posts/me/current$	59.599	0.227	0.880	1.123	586
0.8%	27	0	0	GET	/api/me/export	56.655	2.098	12.880	12.880	2012916
0.8%	70	0	0	PUT	/api/me/profile-image	54.936	0.785	4.794	4.794	0
0.7%	67	0	0	PATCH	/api/me	46.780	0.698	5.440	5.440	280
0.6%	35	6	0	POST	^/api/posts/[^/]+/comments$	43.932	1.255	4.928	4.928	305
0.5%	184	0	0	GET	^/api/relationships/[^/]+	34.919	0.190	0.817	1.223	972
0.4%	181	10	0	GET	^/api/posts/[^/]+$	30.966	0.171	0.593	0.598	585
0.4%	16	0	0	DELETE	^/api/posts/[^/]+$	29.194	1.825	8.926	8.926	0
0.2%	21	4	0	PATCH	^/api/posts/[^/]+$	14.144	0.674	2.262	2.262	518
0.2%	78	0	0	GET	^/api/users/[^/]+$	13.982	0.179	0.670	0.670	192
0.2%	83	0	0	GET	^/api/posts/[^/]+/realmojis$	12.785	0.154	0.506	0.506	999
0.2%	74	1	0	GET	/api/me	12.625	0.171	0.557	0.557	299
0.1%	53	0	0	GET	^/api/memories/[^/]+	10.106	0.191	0.600	0.600	506
0.1%	38	0	0	GET	^/api/posts/[^/]+/comments$	8.859	0.233	0.631	0.631	970
0.1%	3	0	0	DELETE	^/api/comments/[^/]+$	7.942	2.647	5.532	5.532	0
0.1%	14	0	0	DELETE	^/api/realmojis/[^/]+$	6.321	0.451	1.317	1.317	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.6%	145	0	0	GET	/api/notifications	1137.664	7.846	9.226	12.102	177
17.1%	157	5	0	POST	/api/posts	576.862	3.674	10.722	11.094	623
12.9%	289	1	0	POST	/api/friend-requests	436.743	1.511	12.547	13.644	115
7.3%	211	0	0	POST	/api/auth/login	247.596	1.173	10.433	11.052	248
4.8%	144	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	161.418	1.121	9.258	9.494	115
4.7%	769	2	0	GET	^/api/images/[^/]+	159.873	0.208	1.258	1.668	178220
3.4%	669	0	0	GET	^/api/feeds/[^/]+	115.403	0.173	0.553	1.069	2302
1.6%	215	0	0	GET	/api/memories	53.228	0.248	1.057	1.421	15275
1.5%	209	0	0	GET	/api/posting-window/current	49.621	0.237	0.666	1.025	240
1.2%	48	2	0	POST	^/api/posts/[^/]+/realmojis$	39.306	0.819	5.106	5.106	456
1.1%	28	0	0	PUT	/api/me/profile-image	38.860	1.388	9.234	9.234	0
1.1%	59	0	0	POST	/api/auth/register	36.723	0.622	4.761	4.761	228
1.0%	17	0	0	GET	/api/me/export	32.849	1.932	6.629	6.629	3024916
0.9%	46	0	0	PUT	^/api/follows/[^/]+$	31.502	0.685	8.506	8.506	0
0.9%	36	0	0	POST	/api/auth/logout	30.123	0.837	5.527	5.527	0
0.9%	143	0	0	GET	/api/users	29.532	0.207	0.648	0.831	237
0.9%	23	0	0	PATCH	/api/me	28.941	1.258	5.282	5.282	265
0.7%	20	1	0	POST	^/api/posts/[^/]+/comments$	23.675	1.184	5.293	5.293	359
0.6%	112	10	0	GET	^/api/posts/me/current$	21.670	0.193	0.573	1.096	587
0.6%	4	0	0	DELETE	^/api/posts/[^/]+$	20.064	5.016	11.134	11.134	0
0.5%	94	0	0	GET	^/api/relationships/[^/]+	17.848	0.190	0.822	0.822	1148
0.5%	20	0	0	PATCH	^/api/posts/[^/]+$	17.750	0.888	3.860	3.860	625
0.5%	110	8	0	GET	^/api/posts/[^/]+$	17.012	0.155	0.391	0.425	573
0.4%	1	0	0	POST	/api/me/account-deletion	13.263	13.263	13.263	13.263	0
0.3%	2	0	0	DELETE	^/api/comments/[^/]+$	9.803	4.902	8.605	8.605	0
0.2%	49	0	0	GET	^/api/users/[^/]+$	7.108	0.145	0.364	0.364	182
0.2%	48	0	0	GET	^/api/posts/[^/]+/realmojis$	6.447	0.134	0.287	0.287	1133
0.2%	8	0	0	DELETE	^/api/realmojis/[^/]+$	5.716	0.714	3.203	3.203	0
0.2%	31	0	0	GET	/api/me	5.693	0.184	0.528	0.528	296
0.2%	35	1	0	GET	^/api/posts/[^/]+/comments$	5.223	0.149	0.409	0.409	1726
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 58540 クエリ / 82 種類 / 合計 382.111s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	61.3%	3381	234.327	69.31	687.68	632.1	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	23.5%	258	89.808	348.09	1747.18	6742.7	124.5	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.0%	1951	19.257	9.87	760.92	118.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.9%	7642	18.780	2.46	76.40	10.4	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.8%	2286	6.830	2.99	59.95	25.5	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.6%	1	2.256	2255.68	2255.68	30399.0	1199519.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.5%	1	1.869	1869.48	1869.48	22529.0	626124.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.5%	495	1.779	3.59	21.34	29.7	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.3%	5356	1.095	0.20	24.80	20.0	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	0.2%	468	0.877	1.87	9.06	21.8	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.2%	1	0.745	745.31	745.31	12051.0	529793.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
12	0.2%	4298	0.580	0.13	17.08	11.1	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
13	0.1%	1683	0.523	0.31	43.14	21.2	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
14	0.1%	2	0.494	246.89	342.17	219541.5	2583.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
15	0.1%	622	0.277	0.45	25.20	22.4	1.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$4)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 583b9bf630c3053678fd2bca5c7ed7aaa7db2cdd
Type: cpu
Time: 2026-10-31 13:03:53 JST
Duration: 60.10s, Total samples = 49.54s (82.43%)
Showing nodes accounting for 36.20s, 73.07% of 49.54s total
Dropped 1065 nodes (cum <= 0.25s)
Showing top 40 nodes out of 311
      flat  flat%   sum%        cum   cum%
     3.92s  7.91%  7.91%      3.92s  7.91%  internal/runtime/syscall/linux.Syscall6
     2.46s  4.97% 12.88%      9.35s 18.87%  runtime.scanObject
     2.39s  4.82% 17.70%      3.09s  6.24%  runtime.findObject
     2.31s  4.66% 22.37%      2.31s  4.66%  runtime.memmove
     2.25s  4.54% 26.91%      2.25s  4.54%  runtime.memclrNoHeapPointers
     2.16s  4.36% 31.27%      2.16s  4.36%  crypto/internal/fips140/sha256.blockAVX2
     1.83s  3.69% 34.96%      3.28s  6.62%  runtime.tryDeferToSpanScan
     1.77s  3.57% 38.53%      5.64s 11.38%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.49s  3.01% 41.54%      5.19s 10.48%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.46s  2.95% 44.49%      2.50s  5.05%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.46s  2.95% 47.44%      1.46s  2.95%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.29s  2.60% 50.04%      1.29s  2.60%  aeshashbody
     1.14s  2.30% 52.34%      1.20s  2.42%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     1.02s  2.06% 54.40%      1.02s  2.06%  runtime.spanClass.sizeclass (inline)
     0.97s  1.96% 56.36%      0.97s  1.96%  example.com/isureal-company/isureal.(*UserRow).active
     0.87s  1.76% 58.11%      0.87s  1.76%  memeqbody
     0.77s  1.55% 59.67%      1.10s  2.22%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.58s  1.17% 60.84%      2.21s  4.46%  runtime.mapaccess1_faststr
     0.47s  0.95% 61.79%      0.47s  0.95%  indexbody
     0.38s  0.77% 62.56%      0.58s  1.17%  example.com/isureal-company/isureal.postLess
     0.36s  0.73% 63.28%      0.36s  0.73%  runtime.futex
     0.35s  0.71% 63.99%      0.35s  0.71%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.33s  0.67% 64.65%      0.81s  1.64%  runtime.greyobject
     0.31s  0.63% 65.28%      0.31s  0.63%  runtime.nextFreeFast (inline)
     0.30s  0.61% 65.89%      0.30s  0.61%  crypto/internal/fips140/aes/gcm.gcmAesEnc
     0.30s  0.61% 66.49%      0.32s  0.65%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.30s  0.61% 67.10%      0.37s  0.75%  runtime.spanOf (inline)
     0.29s  0.59% 67.68%      0.40s  0.81%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.28s  0.57% 68.25%      1.89s  3.82%  runtime.mapassign_faststr
     0.27s  0.55% 68.79%      0.34s  0.69%  runtime.typePointers.next
     0.27s  0.55% 69.34%      0.27s  0.55%  sync/atomic.(*Int32).Add
     0.25s   0.5% 69.84%      0.25s   0.5%  runtime.typePointers.nextFast (inline)
     0.23s  0.46% 70.31%      0.32s  0.65%  runtime.heapArenaOf (inline)
     0.22s  0.44% 70.75%      6.13s 12.37%  runtime.mallocgc
     0.22s  0.44% 71.19%      1.14s  2.30%  runtime.scanSpan
     0.20s   0.4% 71.60%      2.47s  4.99%  crypto/internal/fips140/sha256.(*Digest).Write
     0.19s  0.38% 71.98%      0.78s  1.57%  database/sql.convertAssignRows
     0.19s  0.38% 72.37%      1.06s  2.14%  runtime.mallocgcSmallScanNoHeader
     0.19s  0.38% 72.75%      0.59s  1.19%  runtime.scanObjectsSmall
     0.16s  0.32% 73.07%      0.50s  1.01%  runtime.mapaccess2_faststr
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 51.4% / 100% | 34% | 0.1% | 7 | 116 MB | CPU が飽和 |
| isu2 (2 cpu) | 16.7% / 52% | 0% | 0.4% | 4 | 2497 MB |  |
| isu3 (2 cpu) | 12.4% / 60% | 0% | 0.3% | 3 | 2613 MB |  |
| isu4 (2 cpu) | 49.3% / 100% | 21% | 14.8% | 32 | 100 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 79% / nginx 15% / journalctl 2% / systemd-journal 2% |
| isu2 | isureal 31% / nginx 1% |
| isu3 | isureal 23% |
| isu4 | postgres 59% / isureal 6% |

## Diff from previous run

スコア: 95693 → 85230 (-10463)

`c75a0a9..8caa076`

```
 webapp/go/config.go     | 2 --
 webapp/go/export_api.go | 4 +---
 2 files changed, 1 insertion(+), 5 deletions(-)
```
