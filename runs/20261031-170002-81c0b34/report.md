# run 20261031-170002-81c0b34

## Summary

| | |
|---|---|
| score | 197896 |
| status | pass |
| ref | asia-read-split |
| sha | 81c0b34af983 |
| 所要時間 | 274s |
| 前回比 | -22644 (-10.3%) vs 20261031-163606-15a1e6c |
| ベスト比 | -25846 (-11.6%) vs 20261031-163105-6f2cd6f = 223742 |
| note | asia1 の nginx が GET を asia1 と asia2 に 1:1 で送る、asia2 の GOGC=200（main bdf05cd の上） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	43	119	330	798	798
europe	7	21	63	188	540	539
america	6	14	38	110	310	310
```

```
{"id":"31814","status":"done","portalStatus":"finished","passed":true,"score":197896,"raw":197896,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:02:56.880Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:02:59.868Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:03:02.868Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:03:05.881Z level=INFO msg="asia 2 日目: 43 人"
time=2026-10-31T08:03:08.868Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T08:03:11.868Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T08:03:14.881Z level=INFO msg="asia 3 日目: 119 人"
time=2026-10-31T08:03:17.868Z level=INFO msg="europe 3 日目: 63 人"
time=2026-10-31T08:03:20.895Z level=INFO msg="america 3 日目: 38 人"
time=2026-10-31T08:03:21.302Z level=INFO msg="user moved" user=nimi_5977 from=asia to=america
time=2026-10-31T08:03:22.390Z level=INFO msg="user moved" user=nichi6112 from=europe to=america
time=2026-10-31T08:03:23.889Z level=INFO msg="asia 4 日目: 330 人"
time=2026-10-31T08:03:26.930Z level=INFO msg="europe 4 日目: 188 人"
time=2026-10-31T08:03:29.871Z level=INFO msg="america 4 日目: 110 人"
time=2026-10-31T08:03:30.561Z level=INFO msg="user moved" user=ivyha.6838 from=asia to=america
time=2026-10-31T08:03:30.766Z level=INFO msg="user moved" user=funoyu.7410 from=asia to=america
time=2026-10-31T08:03:31.532Z level=INFO msg="user moved" user=leote.8478 from=europe to=america
time=2026-10-31T08:03:32.165Z level=INFO msg="user moved" user=robenmar6890 from=asia to=america
time=2026-10-31T08:03:32.882Z level=INFO msg="asia 5 日目: 798 人"
time=2026-10-31T08:03:34.116Z level=INFO msg="user moved" user=nichi6112 from=america to=europe
time=2026-10-31T08:03:34.370Z level=INFO msg="user moved" user=nimi_5977 from=america to=asia
time=2026-10-31T08:03:36.377Z level=INFO msg="europe 5 日目: 540 人"
time=2026-10-31T08:03:39.510Z level=INFO msg="america 5 日目: 310 人"
time=2026-10-31T08:03:48.319Z level=INFO msg="america 最終日（グランプリ）: 310 人"
time=2026-10-31T08:03:48.483Z level=INFO msg="europe 最終日（グランプリ）: 539 人"
time=2026-10-31T08:03:48.609Z level=INFO msg="asia 最終日（グランプリ）: 798 人"
time=2026-10-31T08:03:56.986Z level=INFO msg="グランプリ: アジア 426 人, ヨーロッパ 243 人, アメリカ 192 人"
{
  "pass": true,
  "score": 197896,
  "messages": [
    "rehi5100: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "tormi_10688: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tonoa8741: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hitor9010: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yoso9905: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "vinnimia_10593: ログインし直せません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "misaki_ito_4605: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
49.9%	1355	0	0	GET	/api/notifications	14179.258	10.464	15.734	15.737	170
8.2%	2103	60	0	POST	/api/posts	2333.088	1.109	6.840	7.404	632
7.8%	12090	0	0	GET	^/api/feeds/[^/]+	2207.422	0.183	1.131	2.373	4647
4.9%	3743	0	0	GET	/api/memories	1399.310	0.374	2.130	2.927	12533
3.6%	15814	248	0	GET	^/api/images/[^/]+	1026.723	0.065	0.776	3.048	129863
3.3%	4697	51	0	POST	/api/friend-requests	946.750	0.202	1.154	1.633	114
2.4%	2976	0	0	GET	/api/posting-window/current	676.253	0.227	1.251	1.680	237
2.2%	580	123	0	PUT	/api/me/profile-image	627.305	1.082	11.210	14.424	3
1.9%	702	10	0	POST	^/api/posts/[^/]+/realmojis$	538.924	0.768	4.364	4.563	483
1.9%	3163	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	528.950	0.167	1.203	1.780	116
1.8%	2283	0	0	GET	/api/users	510.413	0.224	1.218	1.658	239
1.8%	5429	592	0	POST	/api/auth/login	499.028	0.092	0.928	1.945	220
1.6%	329	0	0	GET	/api/me/export	463.071	1.408	6.959	8.816	1969397
1.5%	1703	1	0	GET	^/api/posts/me/current$	423.508	0.249	1.295	1.668	652
1.2%	2196	58	0	GET	^/api/posts/[^/]+$	332.984	0.152	1.030	1.402	619
0.9%	1484	0	0	GET	^/api/relationships/[^/]+	263.422	0.178	1.185	1.779	1304
0.7%	788	0	0	POST	^/api/posts/[^/]+/comments$	210.848	0.268	1.632	1.889	347
0.6%	750	0	0	GET	^/api/memories/[^/]+	173.454	0.231	1.632	2.737	634
0.5%	895	1	0	GET	^/api/posts/[^/]+/realmojis$	153.907	0.172	1.180	1.546	1828
0.5%	1135	0	0	GET	^/api/users/[^/]+$	153.582	0.135	1.015	1.257	212
0.5%	468	3	0	GET	/api/me	131.480	0.281	1.315	1.410	300
0.5%	782	1	0	GET	^/api/posts/[^/]+/comments$	130.193	0.166	1.063	1.430	2406
0.4%	876	0	0	POST	/api/auth/logout	114.910	0.131	1.191	1.842	0
0.2%	519	0	0	PUT	^/api/follows/[^/]+$	62.183	0.120	0.865	1.226	0
0.2%	397	0	0	PATCH	/api/me	58.972	0.149	0.931	1.672	276
0.2%	134	0	0	DELETE	^/api/realmojis/[^/]+$	48.322	0.361	1.461	1.547	0
0.2%	188	0	0	PATCH	^/api/posts/[^/]+$	47.696	0.254	1.174	1.252	655
0.2%	84	0	0	DELETE	^/api/posts/[^/]+$	47.106	0.561	1.587	1.587	0
0.2%	818	1	0	POST	/api/auth/register	46.737	0.057	0.335	0.905	227
0.1%	172	0	0	DELETE	^/api/comments/[^/]+$	34.098	0.198	1.400	1.780	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.3%	839	1	0	GET	/api/notifications	6073.082	7.238	12.075	12.287	174
8.5%	1352	53	0	POST	/api/posts	1137.261	0.841	6.696	7.516	628
8.1%	2985	31	0	POST	/api/friend-requests	1087.775	0.364	1.420	1.729	114
6.6%	5988	0	0	GET	^/api/feeds/[^/]+	883.463	0.148	1.700	2.255	3576
5.3%	6014	33	0	GET	^/api/images/[^/]+	711.459	0.118	1.148	3.606	168738
4.5%	1825	0	0	POST	/api/auth/login	605.908	0.332	1.293	1.961	247
2.5%	305	0	0	PUT	/api/me/profile-image	337.967	1.108	4.478	4.760	0
2.3%	1510	0	0	GET	/api/memories	302.595	0.200	1.968	2.869	13398
2.2%	539	0	0	POST	/api/auth/register	292.154	0.542	1.161	1.220	229
1.9%	373	6	0	POST	^/api/posts/[^/]+/realmojis$	254.062	0.681	3.988	4.552	475
1.7%	1927	0	0	GET	/api/posting-window/current	223.940	0.116	0.474	1.913	239
1.6%	639	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	218.601	0.342	1.430	1.961	116
1.3%	970	3	0	GET	^/api/posts/me/current$	180.158	0.186	1.806	2.205	651
1.3%	1490	0	0	GET	/api/users	168.083	0.113	0.512	1.942	238
1.2%	370	0	0	PUT	^/api/follows/[^/]+$	155.315	0.420	1.111	1.541	0
1.0%	263	0	0	PATCH	/api/me	130.361	0.496	1.384	2.028	275
0.7%	744	0	0	GET	^/api/relationships/[^/]+	98.393	0.132	1.462	2.145	617
0.6%	224	1	0	POST	^/api/posts/[^/]+/comments$	85.804	0.383	1.898	2.122	359
0.6%	801	36	0	GET	^/api/posts/[^/]+$	81.826	0.102	0.545	1.950	611
0.5%	277	0	0	POST	/api/auth/logout	71.132	0.257	1.361	1.602	0
0.4%	97	0	0	PATCH	^/api/posts/[^/]+$	52.471	0.541	1.942	1.942	656
0.3%	422	0	0	GET	^/api/users/[^/]+$	43.607	0.103	0.702	1.433	209
0.3%	388	0	0	GET	^/api/posts/[^/]+/realmojis$	40.262	0.104	0.448	1.949	2365
0.2%	34	0	0	DELETE	^/api/posts/[^/]+$	32.060	0.943	2.074	2.074	0
0.2%	313	2	0	GET	/api/me	30.872	0.099	0.482	1.259	302
0.2%	304	0	0	GET	^/api/posts/[^/]+/comments$	28.734	0.095	0.425	0.704	3124
0.2%	56	0	0	DELETE	^/api/realmojis/[^/]+$	26.493	0.473	1.870	1.870	0
0.2%	236	0	0	GET	^/api/memories/[^/]+	22.301	0.094	0.571	1.129	997
0.1%	51	0	0	DELETE	^/api/comments/[^/]+$	14.846	0.291	1.354	1.354	0
0.1%	93	0	0	GET	/api/me/export	9.411	0.101	0.580	0.580	2130166
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.4%	492	0	0	GET	/api/notifications	3120.208	6.342	9.766	9.801	171
8.2%	760	18	0	POST	/api/posts	561.391	0.739	6.831	7.294	643
7.1%	4267	17	0	GET	^/api/images/[^/]+	485.527	0.114	1.104	2.962	149605
6.9%	1834	16	0	POST	/api/friend-requests	472.786	0.258	1.359	2.039	114
6.6%	3455	0	0	GET	^/api/feeds/[^/]+	451.254	0.131	1.055	1.511	3665
5.5%	1280	0	0	POST	/api/auth/login	376.835	0.294	1.307	1.759	251
2.6%	1106	0	0	GET	/api/memories	179.843	0.163	1.228	2.725	16841
2.1%	305	0	0	POST	/api/auth/register	141.342	0.463	1.033	1.053	230
1.7%	394	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	117.081	0.297	1.479	1.681	116
1.6%	628	15	0	GET	^/api/posts/me/current$	109.001	0.174	1.074	1.218	643
1.5%	1097	0	0	GET	/api/posting-window/current	104.658	0.095	0.744	1.053	239
1.3%	181	0	0	PUT	/api/me/profile-image	91.489	0.505	1.571	1.702	0
1.2%	901	0	0	GET	/api/users	84.695	0.094	0.694	1.059	244
1.1%	115	6	0	POST	^/api/posts/[^/]+/realmojis$	74.876	0.651	4.048	4.059	441
1.0%	178	3	0	POST	^/api/posts/[^/]+/comments$	67.057	0.377	1.528	1.535	348
0.8%	516	15	0	GET	^/api/posts/[^/]+$	52.762	0.102	0.897	1.255	614
0.8%	454	0	0	GET	^/api/relationships/[^/]+	51.674	0.114	1.019	1.113	815
0.7%	213	0	0	POST	/api/auth/logout	51.076	0.240	0.966	1.576	0
0.7%	173	0	0	PUT	^/api/follows/[^/]+$	45.076	0.261	1.475	1.561	0
0.6%	149	0	0	PATCH	/api/me	38.132	0.256	0.799	0.895	281
0.5%	74	6	0	PATCH	^/api/posts/[^/]+$	34.661	0.468	1.489	1.489	603
0.3%	193	0	0	GET	^/api/posts/[^/]+/comments$	22.597	0.117	1.073	1.230	3373
0.3%	249	0	0	GET	^/api/users/[^/]+$	22.152	0.089	0.677	0.951	205
0.3%	26	0	0	DELETE	^/api/posts/[^/]+$	22.050	0.848	1.697	1.697	0
0.3%	194	0	0	GET	^/api/posts/[^/]+/realmojis$	21.081	0.109	1.057	1.058	1959
0.3%	86	0	0	GET	/api/me/export	19.086	0.222	4.798	4.798	3153674
0.2%	181	0	0	GET	^/api/memories/[^/]+	16.783	0.093	0.761	0.987	1008
0.2%	185	0	0	GET	/api/me	14.691	0.079	0.154	0.167	306
0.2%	30	0	0	DELETE	^/api/comments/[^/]+$	13.388	0.446	1.370	1.370	0
0.2%	23	0	0	DELETE	^/api/realmojis/[^/]+$	11.141	0.484	1.553	1.553	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 14895 クエリ / 48 種類 / 合計 170.286s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	52.0%	335	88.529	264.26	1618.46	1252.9	127.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.7%	4	25.112	6278.03	9523.11	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.0%	4	18.673	4668.35	7166.62	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	10.3%	11003	17.616	1.60	9.97	9.7	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	5.5%	4	9.289	2322.17	3423.50	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	4.6%	7	7.875	1125.01	4541.24	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	7	0.618	88.24	161.34	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.4%	7	0.605	86.39	162.06	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.424	423.73	423.73	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.334	333.77	333.77	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.1%	1	0.231	230.80	230.80	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.155	155.39	155.39	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.115	114.70	114.70	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	7	0.109	15.54	25.90	166.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	0.1%	1	0.096	96.41	96.41	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows

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
Build ID: b26eb1556a3c14267af0a6063c83c187a13763f7
Type: cpu
Time: 2026-10-31 17:02:29 JST
Duration: 60s, Total samples = 27.25s (45.42%)
Showing nodes accounting for 20.21s, 74.17% of 27.25s total
Dropped 806 nodes (cum <= 0.14s)
Showing top 40 nodes out of 277
      flat  flat%   sum%        cum   cum%
     3.06s 11.23% 11.23%      3.06s 11.23%  internal/runtime/syscall/linux.Syscall6
     2.84s 10.42% 21.65%      2.84s 10.42%  crypto/internal/fips140/sha256.blockAVX2
     1.59s  5.83% 27.49%      5.94s 21.80%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.49s  5.47% 32.95%      1.49s  5.47%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.43s  5.25% 38.20%      1.43s  5.25%  runtime.memmove
     1.11s  4.07% 42.28%      1.11s  4.07%  aeshashbody
     0.97s  3.56% 45.83%      0.97s  3.56%  runtime.memclrNoHeapPointers
     0.80s  2.94% 48.77%      0.99s  3.63%  runtime.findObject
     0.72s  2.64% 51.41%      2.97s 10.90%  runtime.scanObject
     0.58s  2.13% 53.54%      0.93s  3.41%  runtime.tryDeferToSpanScan
     0.54s  1.98% 55.52%      0.76s  2.79%  example.com/isureal-company/isureal/jpegcheck.(*decoder).skipAC
     0.44s  1.61% 57.14%      2.02s  7.41%  runtime.mapaccess1_faststr
     0.36s  1.32% 58.46%      0.36s  1.32%  runtime.futex
     0.35s  1.28% 59.74%      0.53s  1.94%  example.com/isureal-company/isureal.postLess
     0.29s  1.06% 60.81%      0.29s  1.06%  memeqbody
     0.28s  1.03% 61.83%      0.28s  1.03%  runtime.nextFreeFast (inline)
     0.27s  0.99% 62.83%      0.27s  0.99%  runtime.spanClass.sizeclass (inline)
     0.25s  0.92% 63.74%      1.56s  5.72%  runtime.mapassign_faststr
     0.24s  0.88% 64.62%      0.24s  0.88%  sync/atomic.(*Int32).Add (inline)
     0.23s  0.84% 65.47%      0.38s  1.39%  database/sql.convertAssignRows
     0.22s  0.81% 66.28%      0.92s  3.38%  runtime.mallocgcSmallScanNoHeader
     0.21s  0.77% 67.05%      1.08s  3.96%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.19s   0.7% 67.74%      0.19s   0.7%  example.com/isureal-company/isureal/jpegcheck.(*decoder).prefetch
     0.16s  0.59% 68.33%      0.19s   0.7%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.15s  0.55% 68.88%      0.16s  0.59%  example.com/isureal-company/isureal.rmLess
     0.15s  0.55% 69.43%      2.06s  7.56%  runtime.mallocgc
     0.14s  0.51% 69.94%      0.17s  0.62%  runtime.(*mspan).writeHeapBitsSmall
     0.13s  0.48% 70.42%      3.17s 11.63%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.12s  0.44% 70.86%      3.49s 12.81%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.10s  0.37% 71.23%      2.29s  8.40%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.37% 71.60%      0.23s  0.84%  internal/runtime/maps.(*Iter).Next
     0.10s  0.37% 71.96%      0.52s  1.91%  runtime.mallocgcSmallNoscan
     0.09s  0.33% 72.29%      0.16s  0.59%  compress/flate.(*compressor).deflate
     0.09s  0.33% 72.62%      3.09s 11.34%  crypto/internal/fips140/sha256.(*Digest).Write
     0.09s  0.33% 72.95%      0.28s  1.03%  runtime.greyobject
     0.08s  0.29% 73.25%      0.24s  0.88%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.07s  0.26% 73.50%      0.19s   0.7%  runtime.scanObjectsSmall
     0.06s  0.22% 73.72%      0.78s  2.86%  encoding/gob.(*Encoder).encodeStruct
     0.06s  0.22% 73.94%      0.40s  1.47%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func3
     0.06s  0.22% 74.17%      3.51s 12.88%  example.com/isureal-company/isureal.loadBig
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 35.6% / 100% | 14% | 0.6% | 4 | 133 MB |  |
| isu2 (2 cpu) | 24.1% / 100% | 7% | 0.7% | 6 | 126 MB | run queue が積もっている |
| isu3 (2 cpu) | 19.9% / 100% | 2% | 0.8% | 4 | 170 MB |  |
| isu4 (2 cpu) | 62.2% / 100% | 54% | 3.7% | 7 | 103 MB | CPU が飽和 |

**isu2（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 56% / nginx 10% |
| isu2 | isureal 44% / nginx 2% |
| isu3 | isureal 37% / nginx 2% |
| isu4 | isureal 96% / postgres 17% |

## Diff from previous run

スコア: 220540 → 197896 (-22644)

`15a1e6c..81c0b34`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf  | 6 +++---
 conf/isu4/etc/systemd/system/isureal-go.service | 2 +-
 2 files changed, 4 insertions(+), 4 deletions(-)
```
