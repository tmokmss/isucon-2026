# run 20261031-141745-a28a701

## Summary

| | |
|---|---|
| score | 140030 |
| status | pass |
| ref | a28a701 |
| sha | a28a701c6be3 |
| 所要時間 | 337s |
| 前回比 | -1305 (-0.9%) vs 20261031-141027-bd0fe45 |
| ベスト比 | -1305 (-0.9%) vs 20261031-141027-bd0fe45 = 141335 |
| note | 複製 local・時刻の印で待つ（main e42c64c 取り込み済み） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	139	404	1006	1005
europe	7	18	50	140	262	262
america	6	20	58	160	195	195
```

```
{"id":"17033","status":"done","portalStatus":"finished","passed":true,"score":140030,"raw":140030,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:21:53.896Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:21:56.861Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:21:59.860Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:22:02.895Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T05:22:05.862Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T05:22:08.278Z level=INFO msg="user moved" user=ulnetor5330 from=europe to=america
time=2026-10-31T05:22:08.862Z level=INFO msg="america 2 日目: 20 人"
time=2026-10-31T05:22:11.889Z level=INFO msg="asia 3 日目: 139 人"
time=2026-10-31T05:22:12.698Z level=INFO msg="user moved" user=kena_5615 from=asia to=america
time=2026-10-31T05:22:13.589Z level=INFO msg="user moved" user=kena_5615 from=america to=asia
time=2026-10-31T05:22:14.862Z level=INFO msg="europe 3 日目: 50 人"
time=2026-10-31T05:22:15.261Z level=INFO msg="user moved" user=ramoyu.6048 from=asia to=europe
time=2026-10-31T05:22:17.900Z level=INFO msg="america 3 日目: 58 人"
time=2026-10-31T05:22:19.679Z level=INFO msg="user moved" user=ulnetor5330 from=america to=europe
time=2026-10-31T05:22:20.901Z level=INFO msg="asia 4 日目: 404 人"
time=2026-10-31T05:22:23.949Z level=INFO msg="europe 4 日目: 140 人"
time=2026-10-31T05:22:24.056Z level=INFO msg="user moved" user=maxnoane.6538 from=america to=europe
time=2026-10-31T05:22:25.357Z level=INFO msg="user moved" user=yuismar-7892 from=asia to=america
time=2026-10-31T05:22:25.403Z level=INFO msg="user moved" user=kena_5615 from=asia to=america
time=2026-10-31T05:22:25.761Z level=INFO msg="user moved" user=kakane_7492 from=asia to=america
time=2026-10-31T05:22:26.855Z level=INFO msg="user moved" user=rito-8358 from=asia to=europe
time=2026-10-31T05:22:26.952Z level=INFO msg="america 4 日目: 160 人"
time=2026-10-31T05:22:26.957Z level=INFO msg="user moved" user=ulnetor5330 from=europe to=america
time=2026-10-31T05:22:26.972Z level=INFO msg="user moved" user=maxnoane.6538 from=europe to=america
time=2026-10-31T05:22:27.425Z level=INFO msg="user moved" user=yuel-8050 from=asia to=america
time=2026-10-31T05:22:27.679Z level=INFO msg="user moved" user=hovin.7728 from=asia to=america
time=2026-10-31T05:22:28.195Z level=INFO msg="user moved" user=marnomia-7165 from=asia to=europe
time=2026-10-31T05:22:28.492Z level=INFO msg="user moved" user=katonoa.7690 from=asia to=europe
time=2026-10-31T05:22:30.066Z level=INFO msg="asia 5 日目: 1006 人"
time=2026-10-31T05:22:33.024Z level=INFO msg="europe 5 日目: 262 人"
time=2026-10-31T05:22:35.922Z level=INFO msg="america 5 日目: 195 人"
time=2026-10-31T05:22:45.404Z level=INFO msg="asia 最終日（グランプリ）: 1005 人"
time=2026-10-31T05:22:45.812Z level=INFO msg="europe 最終日（グランプリ）: 262 人"
time=2026-10-31T05:22:46.021Z level=INFO msg="america 最終日（グランプリ）: 195 人"
time=2026-10-31T05:22:53.901Z level=INFO msg="グランプリ: アジア 516 人, ヨーロッパ 92 人, アメリカ 69 人"
{
  "pass": true,
  "score": 140030,
  "messages": [
    "tsuleo_12032: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "suni.12965: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "leoto12713: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "evaivyru_13158: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kukisa-14862: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nakuku-14521: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nelia14638: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.3%	1604	0	0	GET	/api/notifications	13078.390	8.154	15.540	15.573	167
17.0%	2399	87	0	POST	/api/posts	6284.581	2.620	9.226	11.220	627
10.5%	5594	79	0	POST	/api/friend-requests	3887.197	0.695	4.712	8.496	114
8.9%	15385	444	0	GET	^/api/images/[^/]+	3311.981	0.215	3.558	9.274	133838
5.9%	5880	1099	0	POST	/api/auth/login	2170.213	0.369	4.770	10.819	203
4.6%	688	130	0	PUT	/api/me/profile-image	1701.761	2.473	10.178	10.386	3
4.1%	2265	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1531.160	0.676	4.784	6.496	116
2.6%	669	57	0	POST	^/api/posts/[^/]+/realmojis$	966.623	1.445	8.259	11.322	446
2.1%	1031	1	0	POST	/api/auth/register	792.166	0.768	6.556	8.039	227
1.4%	609	19	0	POST	^/api/posts/[^/]+/comments$	534.246	0.877	5.553	7.587	336
1.4%	634	0	0	PUT	^/api/follows/[^/]+$	521.690	0.823	5.534	9.710	0
1.2%	3710	0	0	GET	/api/memories	434.556	0.117	0.697	1.572	11951
1.1%	492	0	0	PATCH	/api/me	421.136	0.856	5.638	8.494	277
0.7%	300	0	0	GET	/api/me/export	267.472	0.892	5.958	7.158	1941403
0.6%	780	0	0	POST	/api/auth/logout	229.852	0.295	4.111	6.998	0
0.6%	12773	0	0	GET	^/api/feeds/[^/]+	222.110	0.017	0.212	2.729	2663
0.4%	218	12	0	PATCH	^/api/posts/[^/]+$	155.422	0.713	4.685	5.966	615
0.3%	88	0	0	DELETE	^/api/posts/[^/]+$	122.028	1.387	6.702	6.702	0
0.2%	3542	0	0	GET	/api/posting-window/current	65.177	0.018	0.147	1.940	237
0.1%	105	0	0	DELETE	^/api/comments/[^/]+$	50.273	0.479	3.454	4.087	0
0.1%	728	0	0	GET	^/api/memories/[^/]+	46.734	0.064	0.695	1.365	644
0.1%	91	0	0	DELETE	^/api/realmojis/[^/]+$	43.797	0.481	3.474	3.474	0
0.1%	2674	0	0	GET	/api/users	40.961	0.015	0.123	0.278	243
0.1%	1940	70	0	GET	^/api/posts/me/current$	37.640	0.019	0.206	1.722	627
0.1%	1649	0	0	GET	^/api/relationships/[^/]+	28.414	0.017	0.220	0.437	963
0.1%	1925	12	0	GET	^/api/posts/[^/]+$	28.193	0.015	0.123	0.540	609
0.0%	32	0	0	PUT	^/api/blocks/[^/]+$	14.885	0.465	6.183	6.183	0
0.0%	1	0	0	POST	/api/initialize	13.621	13.621	13.621	13.621	90
0.0%	947	0	0	GET	^/api/users/[^/]+$	13.042	0.014	0.117	0.344	189
0.0%	882	0	0	GET	^/api/posts/[^/]+/realmojis$	10.473	0.012	0.093	0.332	1359
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.5%	445	1	0	GET	/api/notifications	3563.516	8.008	12.965	12.966	161
11.4%	1291	19	0	POST	/api/friend-requests	1140.209	0.883	5.236	9.286	114
11.3%	566	12	0	POST	/api/posts	1133.106	2.002	6.903	9.317	640
6.5%	796	0	0	POST	/api/auth/login	651.681	0.819	4.932	7.438	245
6.1%	2549	14	0	GET	^/api/images/[^/]+	616.171	0.242	1.272	3.422	171650
4.6%	2112	0	0	GET	^/api/feeds/[^/]+	461.865	0.219	0.986	1.698	2341
3.4%	828	0	0	GET	/api/posting-window/current	338.957	0.409	1.164	1.619	239
3.3%	304	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	336.260	1.106	5.685	7.991	116
3.0%	257	0	0	POST	/api/auth/register	300.386	1.169	5.994	6.065	228
2.0%	145	0	0	PUT	/api/me/profile-image	196.853	1.358	5.877	6.341	0
1.9%	633	0	0	GET	/api/users	192.468	0.304	1.095	1.287	247
1.6%	716	0	0	GET	/api/memories	162.986	0.228	1.535	2.821	11502
1.2%	71	6	0	POST	^/api/posts/[^/]+/realmojis$	119.787	1.687	8.896	8.896	446
1.1%	157	0	0	PUT	^/api/follows/[^/]+$	107.858	0.687	3.555	3.986	0
0.9%	121	0	0	POST	^/api/posts/[^/]+/comments$	90.898	0.751	4.707	5.131	355
0.9%	450	12	0	GET	^/api/posts/me/current$	90.362	0.201	1.085	1.168	637
0.9%	114	0	0	PATCH	/api/me	85.675	0.752	3.819	4.353	281
0.8%	65	4	0	PATCH	^/api/posts/[^/]+$	77.890	1.198	5.854	5.854	606
0.7%	124	0	0	POST	/api/auth/logout	72.947	0.588	4.142	4.988	0
0.5%	285	0	0	GET	^/api/relationships/[^/]+	48.270	0.169	0.965	1.611	831
0.5%	275	12	0	GET	^/api/posts/[^/]+$	46.592	0.169	0.941	0.985	587
0.4%	40	0	0	GET	/api/me/export	40.123	1.003	7.919	7.919	2180455
0.3%	150	1	0	GET	/api/me	31.622	0.211	0.852	0.856	302
0.2%	134	0	0	GET	^/api/users/[^/]+$	24.167	0.180	0.954	0.977	198
0.2%	13	0	0	DELETE	^/api/realmojis/[^/]+$	23.692	1.822	4.199	4.199	0
0.2%	111	0	0	GET	^/api/posts/[^/]+/realmojis$	22.253	0.200	0.935	0.954	1594
0.2%	15	0	0	DELETE	^/api/posts/[^/]+$	20.619	1.375	5.647	5.647	0
0.2%	89	0	0	GET	^/api/memories/[^/]+	17.780	0.200	1.543	1.543	904
0.1%	86	0	0	GET	^/api/posts/[^/]+/comments$	11.057	0.129	0.817	0.817	2057
0.1%	25	0	0	DELETE	^/api/comments/[^/]+$	9.341	0.374	2.494	2.494	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.7%	377	0	0	GET	/api/notifications	2928.792	7.769	10.324	10.546	173
11.2%	501	18	0	POST	/api/posts	915.969	1.828	5.828	7.810	628
11.1%	984	11	0	POST	/api/friend-requests	914.713	0.930	5.176	7.795	114
8.4%	2298	23	0	GET	^/api/images/[^/]+	687.546	0.299	1.641	3.797	155022
6.0%	691	0	0	POST	/api/auth/login	488.805	0.707	5.364	6.721	249
4.8%	1430	0	0	GET	^/api/feeds/[^/]+	390.797	0.273	1.032	1.290	2433
3.4%	266	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	276.748	1.040	5.034	6.989	116
3.2%	703	0	0	GET	/api/posting-window/current	260.264	0.370	1.076	1.138	239
2.2%	654	0	0	GET	/api/memories	176.588	0.270	1.580	2.260	14785
1.6%	478	0	0	GET	/api/users	134.961	0.282	1.107	1.225	241
1.5%	113	0	0	PUT	/api/me/profile-image	125.872	1.114	5.843	8.184	0
1.4%	409	30	0	GET	^/api/posts/me/current$	111.554	0.273	0.987	1.014	605
1.3%	188	0	0	POST	/api/auth/register	106.903	0.569	4.209	6.597	229
1.1%	119	0	0	PUT	^/api/follows/[^/]+$	92.729	0.779	4.432	8.841	0
0.9%	52	12	0	POST	^/api/posts/[^/]+/realmojis$	76.418	1.470	10.084	10.084	395
0.9%	103	0	0	POST	/api/auth/logout	75.784	0.736	6.517	6.677	0
0.8%	68	3	0	POST	^/api/posts/[^/]+/comments$	62.205	0.915	7.445	7.445	329
0.7%	84	0	0	PATCH	/api/me	60.845	0.724	3.480	3.480	285
0.6%	223	17	0	GET	^/api/posts/[^/]+$	47.196	0.212	0.938	1.106	572
0.6%	37	5	0	PATCH	^/api/posts/[^/]+$	45.724	1.236	5.911	5.911	576
0.5%	203	0	0	GET	^/api/relationships/[^/]+	39.749	0.196	0.927	1.104	1116
0.5%	25	0	0	DELETE	^/api/posts/[^/]+$	37.819	1.513	5.022	5.022	0
0.4%	116	0	0	GET	/api/me	32.735	0.282	0.946	1.015	304
0.4%	32	0	0	GET	/api/me/export	31.410	0.982	9.654	9.654	3231509
0.3%	101	0	0	GET	^/api/memories/[^/]+	26.438	0.262	1.365	1.402	870
0.2%	85	0	0	GET	^/api/users/[^/]+$	19.206	0.226	0.994	0.994	195
0.2%	83	1	0	GET	^/api/posts/[^/]+/realmojis$	13.905	0.168	0.941	0.941	876
0.1%	60	0	0	GET	^/api/posts/[^/]+/comments$	10.632	0.177	0.806	0.806	1828
0.1%	7	0	0	DELETE	^/api/comments/[^/]+$	8.254	1.179	2.450	2.450	0
0.1%	5	0	0	DELETE	^/api/realmojis/[^/]+$	5.169	1.034	2.567	2.567	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 86239 クエリ / 83 種類 / 合計 369.409s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	35.9%	317	132.700	418.61	3980.67	1243.7	127.0	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	17.3%	4888	63.771	13.05	367.28	570.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	13.1%	4059	48.546	11.96	779.04	116.2	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	8.9%	10593	32.893	3.11	70.12	10.8	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	5.8%	3	21.449	7149.53	9641.39	31471.0	1199976.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	5.5%	3355	20.166	6.01	228.46	24.7	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
7	4.4%	3	16.087	5362.19	7197.68	24949.0	628389.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	2.1%	3	7.877	2625.76	3591.96	13129.0	529813.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	1.2%	7744	4.260	0.55	220.47	20.1	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	1.1%	2773	3.905	1.41	221.74	20.8	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
11	0.8%	717	2.791	3.89	49.66	28.3	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
12	0.7%	6	2.482	413.72	1328.76	213595.7	2621.7	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
13	0.5%	6268	1.753	0.28	189.90	10.9	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
14	0.4%	6	1.579	263.14	379.23	27848.2	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	0.4%	776	1.531	1.97	93.59	20.5	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)

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
Build ID: b997c195cef2024b6df7b12ed748694ec109ccf0
Type: cpu
Time: 2026-10-31 14:21:34 JST
Duration: 60.07s, Total samples = 44.74s (74.48%)
Showing nodes accounting for 31.48s, 70.36% of 44.74s total
Dropped 1111 nodes (cum <= 0.22s)
Showing top 40 nodes out of 302
      flat  flat%   sum%        cum   cum%
     5.55s 12.41% 12.41%      5.55s 12.41%  crypto/internal/fips140/sha256.blockAVX2
     4.52s 10.10% 22.51%      4.52s 10.10%  internal/runtime/syscall/linux.Syscall6
     1.92s  4.29% 26.80%      1.92s  4.29%  runtime.memmove
     1.89s  4.22% 31.02%      2.43s  5.43%  runtime.findObject
     1.74s  3.89% 34.91%      2.82s  6.30%  runtime.tryDeferToSpanScan
     1.69s  3.78% 38.69%      5.72s 12.78%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.64s  3.67% 42.36%      7.24s 16.18%  runtime.scanObject
     1.62s  3.62% 45.98%      1.62s  3.62%  runtime.memclrNoHeapPointers
     1.55s  3.46% 49.44%      1.55s  3.46%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.40s  3.13% 52.57%      1.40s  3.13%  aeshashbody
     0.77s  1.72% 54.29%      0.77s  1.72%  runtime.spanClass.sizeclass (inline)
     0.68s  1.52% 55.81%      2.32s  5.19%  runtime.mapaccess1_faststr
     0.48s  1.07% 56.88%      0.48s  1.07%  memeqbody
     0.37s  0.83% 57.71%      0.56s  1.25%  example.com/isureal-company/isureal.postLess
     0.36s   0.8% 58.52%      0.36s   0.8%  runtime.nextFreeFast (inline)
     0.34s  0.76% 59.28%      0.34s  0.76%  runtime.futex
     0.29s  0.65% 59.92%      0.36s   0.8%  runtime.(*mspan).markBitsForIndex (inline)
     0.27s   0.6% 60.53%      0.27s   0.6%  indexbody
     0.27s   0.6% 61.13%      1.87s  4.18%  runtime.mapassign_faststr
     0.27s   0.6% 61.73%      0.27s   0.6%  sync/atomic.(*Int32).Add
     0.26s  0.58% 62.32%      0.61s  1.36%  compress/flate.(*compressor).deflate
     0.26s  0.58% 62.90%      0.79s  1.77%  database/sql.convertAssignRows
     0.24s  0.54% 63.43%      0.35s  0.78%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.24s  0.54% 63.97%      1.10s  2.46%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.54% 64.51%      0.24s  0.54%  runtime.typePointers.nextFast (inline)
     0.23s  0.51% 65.02%      5.96s 13.32%  crypto/internal/fips140/sha256.(*Digest).Write
     0.22s  0.49% 65.51%      0.26s  0.58%  runtime.heapArenaOf (inline)
     0.21s  0.47% 65.98%      5.03s 11.24%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.21s  0.47% 66.45%      0.64s  1.43%  runtime.greyobject
     0.21s  0.47% 66.92%      5.78s 12.92%  runtime.mallocgc
     0.19s  0.42% 67.34%      5.56s 12.43%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.19s  0.42% 67.77%      3.14s  7.02%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.19s  0.42% 68.19%      0.86s  1.92%  runtime.scanSpan
     0.18s   0.4% 68.60%      0.25s  0.56%  runtime.typePointers.next
     0.17s  0.38% 68.98%      0.23s  0.51%  runtime.spanOf (inline)
     0.13s  0.29% 69.27%      5.30s 11.85%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.13s  0.29% 69.56%      0.63s  1.41%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.12s  0.27% 69.83%      1.24s  2.77%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.27% 70.09%      0.25s  0.56%  internal/runtime/maps.(*Iter).Next
     0.12s  0.27% 70.36%      0.88s  1.97%  runtime.mallocgcSmallNoscan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 50.0% / 100% | 34% | 0.3% | 7 | 178 MB | CPU が飽和 |
| isu2 (2 cpu) | 26.8% / 97% | 2% | 1.1% | 5 | 189 MB | run queue が積もっている |
| isu3 (2 cpu) | 28.9% / 100% | 4% | 2.0% | 4 | 108 MB |  |
| isu4 (2 cpu) | 55.2% / 100% | 28% | 20.1% | 29 | 108 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 82% / nginx 13% |
| isu2 | isureal 50% / nginx 1% |
| isu3 | isureal 51% / nginx 2% / kswapd0 1% |
| isu4 | postgres 50% / isureal 13% / kswapd0 1% / kcompactd0 1% |

## Diff from previous run

スコア: 141335 → 140030 (-1305)

`bd0fe45..a28a701`

```
 webapp/go/jpegcheck/huffman.go | 27 ++++++++++++++++++++++++---
 webapp/go/jpegcheck/reader.go  | 32 ++++++++++++++++++++------------
 webapp/go/jpegcheck/scan.go    |  7 ++++---
 webapp/sql/reset.sql           |  3 +++
 4 files changed, 51 insertions(+), 18 deletions(-)
```
