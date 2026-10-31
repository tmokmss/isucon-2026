# run 20261031-143155-1ce8dee

## Summary

| | |
|---|---|
| score | 124239 |
| status | pass |
| ref | 1ce8dee |
| sha | 1ce8deed43a8 |
| 所要時間 | 124s |
| 前回比 | +13142 (11.8%) vs 20261031-142817-cb92b74 |
| ベスト比 | -17096 (-12.1%) vs 20261031-141027-bd0fe45 = 141335 |
| note | DB外し2: 複製の再接続はスナップショットで |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	45	128	374	907	906
europe	7	18	53	147	269	269
america	6	17	49	131	133	133
```

```
{"id":"17894","status":"done","portalStatus":"finished","passed":true,"score":124239,"raw":124239,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:32:17.741Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:32:20.740Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:32:21.304Z level=INFO msg="user moved" user=nosota_5029 from=asia to=europe
time=2026-10-31T05:32:23.734Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:32:26.739Z level=INFO msg="asia 2 日目: 45 人"
time=2026-10-31T05:32:29.737Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T05:32:32.736Z level=INFO msg="america 2 日目: 17 人"
time=2026-10-31T05:32:35.742Z level=INFO msg="asia 3 日目: 128 人"
time=2026-10-31T05:32:37.125Z level=INFO msg="user moved" user=vindanan.5625 from=asia to=europe
time=2026-10-31T05:32:37.410Z level=INFO msg="user moved" user=nosota_5029 from=europe to=asia
time=2026-10-31T05:32:38.700Z level=INFO msg="user moved" user=nosota_5029 from=asia to=europe
time=2026-10-31T05:32:38.734Z level=INFO msg="europe 3 日目: 53 人"
time=2026-10-31T05:32:39.451Z level=INFO msg="user moved" user=hirimax_5437 from=america to=europe
time=2026-10-31T05:32:41.734Z level=INFO msg="america 3 日目: 49 人"
time=2026-10-31T05:32:44.828Z level=INFO msg="asia 4 日目: 374 人"
time=2026-10-31T05:32:45.246Z level=INFO msg="user moved" user=hirimax_5437 from=europe to=america
time=2026-10-31T05:32:46.403Z level=INFO msg="user moved" user=nufuka6348 from=europe to=asia
time=2026-10-31T05:32:47.827Z level=INFO msg="europe 4 日目: 147 人"
time=2026-10-31T05:32:49.942Z level=INFO msg="user moved" user=ivytor7330 from=asia to=america
time=2026-10-31T05:32:50.942Z level=INFO msg="america 4 日目: 131 人"
time=2026-10-31T05:32:53.452Z level=INFO msg="user moved" user=nufuka6348 from=asia to=america
time=2026-10-31T05:32:53.513Z level=INFO msg="user moved" user=reha-7888 from=asia to=america
time=2026-10-31T05:32:53.913Z level=INFO msg="asia 5 日目: 907 人"
time=2026-10-31T05:32:57.476Z level=INFO msg="europe 5 日目: 269 人"
time=2026-10-31T05:33:01.238Z level=INFO msg="america 5 日目: 133 人"
time=2026-10-31T05:33:10.459Z level=INFO msg="america 最終日（グランプリ）: 133 人"
time=2026-10-31T05:33:10.483Z level=INFO msg="europe 最終日（グランプリ）: 269 人"
time=2026-10-31T05:33:10.606Z level=INFO msg="asia 最終日（グランプリ）: 906 人"
time=2026-10-31T05:33:17.787Z level=INFO msg="グランプリ: アジア 319 人, ヨーロッパ 103 人, アメリカ 37 人"
{
  "pass": true,
  "score": 124239,
  "messages": [
    "sovin.5038: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "isya.13425: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marri-13813: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "liamaki_12812: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "muanmia14069: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.5%	1514	1	0	GET	/api/notifications	14635.786	9.667	17.105	17.113	169
14.8%	2358	76	0	POST	/api/posts	5222.380	2.215	6.995	7.907	635
7.0%	15332	348	0	GET	^/api/images/[^/]+	2475.771	0.161	1.378	2.395	127105
5.7%	5162	72	0	POST	/api/friend-requests	2004.898	0.388	2.312	3.231	114
4.8%	11160	0	0	GET	^/api/feeds/[^/]+	1682.367	0.151	1.217	3.919	3176
4.0%	672	138	0	PUT	/api/me/profile-image	1421.670	2.116	5.575	8.457	4
3.6%	3728	0	0	GET	/api/memories	1271.783	0.341	3.376	8.760	12249
3.3%	5530	747	0	POST	/api/auth/login	1161.287	0.210	1.794	2.538	214
2.4%	2298	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	858.405	0.374	2.262	3.243	116
1.8%	446	23	0	POST	^/api/posts/[^/]+/realmojis$	622.431	1.396	5.660	5.964	449
1.3%	927	1	0	POST	/api/auth/register	466.076	0.503	1.737	1.979	227
1.2%	3350	0	0	GET	/api/posting-window/current	439.679	0.131	0.886	1.897	237
1.0%	2464	0	0	GET	/api/users	353.030	0.143	1.113	1.865	243
0.8%	1762	71	0	GET	^/api/posts/me/current$	298.331	0.169	1.132	1.865	627
0.8%	283	0	0	GET	/api/me/export	270.759	0.957	9.295	15.035	1920205
0.7%	499	4	0	POST	^/api/posts/[^/]+/comments$	259.500	0.520	3.530	4.873	346
0.6%	1688	52	0	GET	^/api/posts/[^/]+$	229.286	0.136	0.961	1.865	624
0.5%	1474	0	0	GET	^/api/relationships/[^/]+	185.433	0.126	1.077	2.292	1008
0.5%	589	0	0	PUT	^/api/follows/[^/]+$	179.335	0.304	2.261	2.775	0
0.5%	820	0	0	POST	/api/auth/logout	177.517	0.216	2.034	3.143	0
0.5%	541	3	0	GET	/api/me	160.191	0.296	0.773	0.915	302
0.4%	452	0	0	PATCH	/api/me	158.458	0.351	2.002	2.272	277
0.4%	241	15	0	PATCH	^/api/posts/[^/]+$	128.361	0.533	2.284	2.304	614
0.3%	742	0	0	GET	^/api/memories/[^/]+	118.917	0.160	1.642	1.866	641
0.3%	797	0	0	GET	^/api/users/[^/]+$	112.964	0.142	1.491	1.866	215
0.3%	87	0	0	DELETE	^/api/posts/[^/]+$	109.113	1.254	3.517	3.517	0
0.2%	650	0	0	GET	^/api/posts/[^/]+/realmojis$	70.239	0.108	0.908	1.845	886
0.2%	635	0	0	GET	^/api/posts/[^/]+/comments$	70.167	0.110	0.872	2.046	1462
0.1%	99	0	0	DELETE	^/api/comments/[^/]+$	47.339	0.478	3.447	3.447	0
0.1%	65	0	0	DELETE	^/api/realmojis/[^/]+$	46.533	0.716	4.680	4.680	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.0%	474	1	0	GET	/api/notifications	4025.030	8.492	12.740	13.564	150
14.0%	2379	15	0	GET	^/api/images/[^/]+	1518.160	0.638	3.819	5.177	162801
10.0%	597	12	0	POST	/api/posts	1086.856	1.821	6.461	7.860	646
7.5%	1308	15	0	POST	/api/friend-requests	819.062	0.626	2.785	3.052	114
5.0%	1834	0	0	GET	^/api/feeds/[^/]+	542.913	0.296	1.813	2.578	2328
3.9%	767	0	0	POST	/api/auth/login	419.069	0.546	2.752	3.024	245
3.7%	716	0	0	GET	/api/memories	400.216	0.559	3.629	5.970	12115
2.5%	881	0	0	GET	/api/posting-window/current	273.565	0.311	1.530	2.361	239
2.3%	158	0	0	PUT	/api/me/profile-image	253.239	1.603	4.033	4.035	0
2.0%	266	0	0	POST	/api/auth/register	222.761	0.837	2.738	2.752	228
1.9%	629	0	0	GET	/api/users	207.038	0.329	1.751	2.761	249
1.8%	454	24	0	GET	^/api/posts/me/current$	199.781	0.440	1.868	2.749	623
1.7%	326	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	188.694	0.579	2.567	2.781	116
1.0%	173	0	0	PUT	^/api/follows/[^/]+$	109.839	0.635	2.784	2.792	0
0.9%	163	2	0	GET	/api/me	99.793	0.612	1.858	1.858	300
0.8%	120	0	0	PATCH	/api/me	87.500	0.729	2.476	2.788	283
0.7%	267	0	0	GET	^/api/relationships/[^/]+	75.160	0.281	1.336	1.703	957
0.5%	98	4	0	POST	^/api/posts/[^/]+/realmojis$	58.340	0.595	2.892	2.892	474
0.5%	46	0	0	GET	/api/me/export	50.134	1.090	14.891	14.891	2214923
0.4%	96	0	0	POST	/api/auth/logout	47.414	0.494	2.841	2.841	0
0.3%	47	2	0	PATCH	^/api/posts/[^/]+$	33.436	0.711	2.221	2.221	618
0.3%	238	16	0	GET	^/api/posts/[^/]+$	32.056	0.135	0.965	1.221	582
0.3%	19	0	0	DELETE	^/api/posts/[^/]+$	30.632	1.612	3.403	3.403	0
0.3%	82	0	0	GET	^/api/memories/[^/]+	27.904	0.340	2.295	2.295	702
0.2%	106	0	0	GET	^/api/users/[^/]+$	16.374	0.154	0.913	0.997	207
0.1%	41	1	0	POST	^/api/posts/[^/]+/comments$	16.261	0.397	2.672	2.672	324
0.1%	100	1	0	GET	^/api/posts/[^/]+/realmojis$	10.742	0.107	0.866	0.971	742
0.1%	15	0	0	DELETE	^/api/realmojis/[^/]+$	9.276	0.618	2.183	2.183	0
0.1%	60	0	0	GET	^/api/posts/[^/]+/comments$	7.240	0.121	1.143	1.143	856
0.0%	7	0	0	PUT	^/api/blocks/[^/]+$	2.241	0.320	0.770	0.770	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.3%	311	0	0	GET	/api/notifications	2474.976	7.958	10.695	10.696	177
15.3%	1758	12	0	GET	^/api/images/[^/]+	1016.236	0.578	3.981	4.816	149014
11.7%	396	5	0	POST	/api/posts	778.691	1.966	6.218	7.197	643
7.8%	720	9	0	POST	/api/friend-requests	517.934	0.719	2.821	3.223	114
3.7%	544	0	0	GET	/api/posting-window/current	248.316	0.456	1.785	2.381	240
3.6%	520	0	0	POST	/api/auth/login	237.922	0.458	2.497	3.175	250
3.5%	854	0	0	GET	^/api/feeds/[^/]+	233.726	0.274	1.402	1.849	2541
3.4%	491	0	0	GET	/api/memories	226.647	0.462	2.408	5.721	15042
2.3%	259	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	154.185	0.595	2.698	2.883	116
1.9%	351	0	0	GET	/api/users	123.271	0.351	1.384	2.094	234
1.7%	321	39	0	GET	^/api/posts/me/current$	110.674	0.345	1.113	1.626	572
1.1%	31	0	0	GET	/api/me/export	70.435	2.272	21.604	21.604	2862267
0.8%	66	0	0	PUT	/api/me/profile-image	52.606	0.797	4.487	4.487	0
0.7%	78	0	0	POST	/api/auth/logout	47.238	0.606	2.540	2.540	0
0.7%	129	0	0	POST	/api/auth/register	46.456	0.360	2.070	2.095	229
0.7%	35	9	0	POST	^/api/posts/[^/]+/realmojis$	45.699	1.306	5.301	5.301	330
0.6%	65	0	0	PATCH	/api/me	39.278	0.604	3.008	3.008	280
0.6%	145	0	0	GET	^/api/relationships/[^/]+	37.852	0.261	1.560	1.574	1291
0.4%	79	0	0	PUT	^/api/follows/[^/]+$	28.344	0.359	2.911	2.911	0
0.4%	59	1	0	POST	^/api/posts/[^/]+/comments$	26.247	0.445	2.920	2.920	282
0.3%	112	8	0	GET	^/api/posts/[^/]+$	19.757	0.176	1.024	1.613	590
0.3%	26	12	0	PATCH	^/api/posts/[^/]+$	18.801	0.723	1.783	1.783	358
0.3%	16	0	0	DELETE	^/api/posts/[^/]+$	17.307	1.082	2.503	2.503	0
0.3%	72	0	0	GET	^/api/memories/[^/]+	17.301	0.240	1.301	1.301	837
0.2%	69	0	0	GET	/api/me	15.637	0.227	1.159	1.159	301
0.1%	53	0	0	GET	^/api/users/[^/]+$	9.600	0.181	1.514	1.514	210
0.1%	27	0	0	GET	^/api/posts/[^/]+/realmojis$	6.731	0.249	2.950	2.950	896
0.1%	45	0	0	GET	^/api/posts/[^/]+/comments$	5.752	0.128	1.037	1.037	1135
0.1%	14	0	0	DELETE	^/api/comments/[^/]+$	4.898	0.350	2.826	2.826	0
0.0%	7	0	0	DELETE	^/api/realmojis/[^/]+$	2.891	0.413	0.699	0.699	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 12658 クエリ / 48 種類 / 合計 135.851s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	53.1%	308	72.149	234.25	1758.72	1223.9	125.1	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	15.7%	3	21.297	7099.09	9575.71	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.8%	3	16.057	5352.18	7184.71	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	10.8%	10691	14.688	1.37	8.62	9.6	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.0%	3	8.089	2696.35	3765.38	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	0.8%	5	1.041	208.30	367.81	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	5	0.553	110.58	156.92	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.4%	5	0.517	103.41	157.37	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.246	246.26	246.26	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.221	220.56	220.56	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.2%	1	0.216	215.98	215.98	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.159	158.76	158.76	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.128	127.95	127.95	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.085	85.38	85.38	2419.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.1%	1	0.078	77.81	77.81	1507.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM posts;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 36588577aa8caaa559dd65bedeeb3134fae0c26d
Type: cpu
Time: 2026-10-31 14:32:01 JST
Duration: 60.02s, Total samples = 55.12s (91.84%)
Showing nodes accounting for 40.09s, 72.73% of 55.12s total
Dropped 1128 nodes (cum <= 0.28s)
Showing top 40 nodes out of 288
      flat  flat%   sum%        cum   cum%
     6.62s 12.01% 12.01%      6.63s 12.03%  crypto/internal/fips140/sha256.blockAVX2
     6.10s 11.07% 23.08%      6.10s 11.07%  internal/runtime/syscall/linux.Syscall6
     4.07s  7.38% 30.46%      4.07s  7.38%  runtime.memmove
     2.64s  4.79% 35.25%     10.19s 18.49%  runtime.scanObject
     2.62s  4.75% 40.00%      3.37s  6.11%  runtime.findObject
     2.18s  3.96% 43.96%      2.18s  3.96%  runtime.memclrNoHeapPointers
     2.16s  3.92% 47.88%      3.57s  6.48%  runtime.tryDeferToSpanScan
     1.64s  2.98% 50.85%      5.61s 10.18%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.47s  2.67% 53.52%      1.47s  2.67%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.15s  2.09% 55.61%      1.15s  2.09%  aeshashbody
     1.08s  1.96% 57.57%      1.08s  1.96%  runtime.spanClass.sizeclass (inline)
     0.52s  0.94% 58.51%      2.11s  3.83%  runtime.mapaccess1_faststr
     0.48s  0.87% 59.38%      0.95s  1.72%  runtime.greyobject
     0.46s  0.83% 60.21%      0.46s  0.83%  memeqbody
     0.44s   0.8% 61.01%      0.44s   0.8%  runtime.nextFreeFast (inline)
     0.38s  0.69% 61.70%      0.58s  1.05%  example.com/isureal-company/isureal.postLess
     0.38s  0.69% 62.39%      0.38s  0.69%  runtime.(*mspan).divideByElemSize (inline)
     0.34s  0.62% 63.01%      0.69s  1.25%  example.com/isureal-company/isureal.(*Store).snapshotLocked
     0.33s   0.6% 63.61%      0.33s   0.6%  runtime.typePointers.nextFast (inline)
     0.32s  0.58% 64.19%      0.34s  0.62%  runtime.(*mspan).writeHeapBitsSmall
     0.31s  0.56% 64.75%      0.35s  0.63%  runtime.(*mspan).markBitsForIndex (inline)
     0.30s  0.54% 65.29%      7.32s 13.28%  crypto/internal/fips140/sha256.(*Digest).Write
     0.29s  0.53% 65.82%      6.13s 11.12%  runtime.mallocgc
     0.29s  0.53% 66.35%      1.47s  2.67%  runtime.mallocgcSmallScanNoHeader
     0.29s  0.53% 66.87%      0.33s   0.6%  runtime.typePointers.next
     0.28s  0.51% 67.38%      0.28s  0.51%  runtime.futex
     0.27s  0.49% 67.87%      6.59s 11.96%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.27s  0.49% 68.36%      0.57s  1.03%  database/sql.convertAssignRows
     0.26s  0.47% 68.83%      0.46s  0.83%  compress/flate.(*compressor).deflate
     0.26s  0.47% 69.30%      5.75s 10.43%  encoding/gob.(*Encoder).encodeStruct
     0.25s  0.45% 69.76%      0.47s  0.85%  internal/runtime/maps.(*Iter).Next
     0.23s  0.42% 70.17%      7.07s 12.83%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.21s  0.38% 70.56%      6.22s 11.28%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.19s  0.34% 70.90%      0.63s  1.14%  runtime.scanObjectsSmall
     0.18s  0.33% 71.23%      1.17s  2.12%  runtime.scanSpan
     0.17s  0.31% 71.53%         4s  7.26%  example.com/isureal-company/isureal.listMemories
     0.17s  0.31% 71.84%      0.89s  1.61%  runtime.mallocgcSmallNoscan
     0.17s  0.31% 72.15%      0.32s  0.58%  runtime.scanObjectSmall
     0.16s  0.29% 72.44%      2.77s  5.03%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.16s  0.29% 72.73%      1.90s  3.45%  runtime.mapassign_faststr
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 49.6% / 100% | 33% | 1.5% | 6 | 104 MB | CPU が飽和 |
| isu2 (2 cpu) | 25.1% / 100% | 1% | 0.9% | 4 | 422 MB |  |
| isu3 (2 cpu) | 24.7% / 100% | 2% | 1.4% | 7 | 871 MB | run queue が積もっている |
| isu4 (2 cpu) | 31.6% / 95% | 6% | 17.5% | 4 | 101 MB |  |

**isu1（CPU が飽和、主に isureal）、isu3（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 81% / nginx 12% / kcompactd0 1% |
| isu2 | isureal 47% / nginx 1% |
| isu3 | isureal 45% / nginx 1% |
| isu4 | isureal 13% / postgres 13% |

## Diff from previous run

スコア: 111097 → 124239 (+13142)

`cb92b74..1ce8dee`

```
 webapp/go/origin_store.go | 13 +++++++++++++
 webapp/go/repl.go         | 16 ++++++++++++++--
 2 files changed, 27 insertions(+), 2 deletions(-)
```
