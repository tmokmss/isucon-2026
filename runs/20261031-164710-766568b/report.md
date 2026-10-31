# run 20261031-164710-766568b

## Summary

| | |
|---|---|
| score | 230652 |
| status | pass |
| ref | HEAD |
| sha | 766568bb6e59 |
| 所要時間 | 158s |
| 前回比 | +10112 (4.6%) vs 20261031-163606-15a1e6c |
| ベスト比 | +6910 (3.1%) vs 20261031-163105-6f2cd6f = 223742 **ベスト更新** |
| note | 閲覧者の友達・ブロックを写さず、友達の友達は使うときだけ作る（main bdf05cd の上） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	133	350	613	613
europe	7	20	54	153	437	437
america	7	18	52	150	420	418
```

```
{"id":"30278","status":"done","portalStatus":"finished","passed":true,"score":230652,"raw":230652,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:47:43.651Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:47:45.049Z level=INFO msg="user moved" user=benel5057 from=asia to=america
time=2026-10-31T07:47:46.644Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:47:49.640Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T07:47:52.653Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T07:47:54.509Z level=INFO msg="user moved" user=liamar.5193 from=asia to=america
time=2026-10-31T07:47:54.731Z level=INFO msg="user moved" user=royoka-5222 from=asia to=europe
time=2026-10-31T07:47:55.588Z level=INFO msg="user moved" user=royoka-5222 from=europe to=asia
time=2026-10-31T07:47:55.636Z level=INFO msg="user moved" user=liamar.5193 from=america to=europe
time=2026-10-31T07:47:55.652Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T07:47:58.639Z level=INFO msg="america 2 日目: 18 人"
time=2026-10-31T07:48:01.659Z level=INFO msg="asia 3 日目: 133 人"
time=2026-10-31T07:48:04.240Z level=INFO msg="user moved" user=royoka-5222 from=asia to=europe
time=2026-10-31T07:48:04.639Z level=INFO msg="europe 3 日目: 54 人"
time=2026-10-31T07:48:05.761Z level=INFO msg="user moved" user=noawaya5944 from=asia to=europe
time=2026-10-31T07:48:06.577Z level=INFO msg="user moved" user=benel5057 from=america to=asia
time=2026-10-31T07:48:06.959Z level=INFO msg="user moved" user=liamar.5193 from=europe to=america
time=2026-10-31T07:48:07.639Z level=INFO msg="america 3 日目: 52 人"
time=2026-10-31T07:48:10.367Z level=INFO msg="user moved" user=maeva6214 from=europe to=america
time=2026-10-31T07:48:10.653Z level=INFO msg="asia 4 日目: 350 人"
time=2026-10-31T07:48:10.833Z level=INFO msg="user moved" user=ronu_6466 from=america to=europe
time=2026-10-31T07:48:13.640Z level=INFO msg="europe 4 日目: 153 人"
time=2026-10-31T07:48:14.484Z level=INFO msg="user moved" user=noawaya5944 from=europe to=asia
time=2026-10-31T07:48:15.321Z level=INFO msg="user moved" user=ronu_6466 from=europe to=america
time=2026-10-31T07:48:16.355Z level=INFO msg="user moved" user=royoka-5222 from=europe to=asia
time=2026-10-31T07:48:16.409Z level=INFO msg="user moved" user=ivymi-7229 from=asia to=america
time=2026-10-31T07:48:16.651Z level=INFO msg="america 4 日目: 150 人"
time=2026-10-31T07:48:18.116Z level=INFO msg="user moved" user=liamar.5193 from=america to=europe
time=2026-10-31T07:48:18.916Z level=INFO msg="user moved" user=mimu-8353 from=europe to=america
time=2026-10-31T07:48:19.093Z level=INFO msg="user moved" user=machi_8873 from=america to=europe
time=2026-10-31T07:48:19.474Z level=INFO msg="user moved" user=hemarse.8050 from=asia to=europe
time=2026-10-31T07:48:19.660Z level=INFO msg="asia 5 日目: 613 人"
time=2026-10-31T07:48:21.487Z level=INFO msg="user moved" user=semieva-9564 from=america to=asia
time=2026-10-31T07:48:22.652Z level=INFO msg="europe 5 日目: 437 人"
time=2026-10-31T07:48:25.747Z level=INFO msg="america 5 日目: 420 人"
time=2026-10-31T07:48:25.863Z level=INFO msg="user moved" user=maeva6214 from=america to=europe
time=2026-10-31T07:48:34.724Z level=INFO msg="america 最終日（グランプリ）: 418 人"
time=2026-10-31T07:48:34.750Z level=INFO msg="asia 最終日（グランプリ）: 613 人"
time=2026-10-31T07:48:34.799Z level=INFO msg="europe 最終日（グランプリ）: 437 人"
time=2026-10-31T07:48:43.676Z level=INFO msg="グランプリ: アジア 471 人, ヨーロッパ 427 人, アメリカ 408 人"
{
  "pass": true,
  "score": 230652,
  "messages": [
    "muko_12573: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "roro_14133: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ruyolia.11814: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "jiwoo_choi_3305: ログアウトできません: POST /api/auth/logout: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/logout\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
79.3%	1798	1	0	GET	/api/notifications	47516.225	26.427	60.005	60.005	113
4.4%	12560	0	0	GET	^/api/feeds/[^/]+	2656.349	0.211	1.517	2.153	7082
3.3%	2982	0	0	GET	/api/memories	1999.210	0.670	5.721	12.926	11009
2.2%	2452	0	0	GET	/api/posting-window/current	1338.482	0.546	1.662	1.881	237
2.0%	1780	24	0	POST	/api/posts	1222.993	0.687	3.407	7.553	642
1.4%	1633	1	0	GET	^/api/posts/me/current$	829.525	0.508	1.734	1.852	651
1.3%	1712	0	0	GET	/api/users	790.245	0.462	1.647	1.911	233
1.0%	12369	245	0	GET	^/api/images/[^/]+	578.065	0.047	0.874	1.803	173401
0.9%	210	0	0	GET	/api/me/export	510.767	2.432	12.003	19.640	1768715
0.7%	468	123	0	PUT	/api/me/profile-image	421.359	0.900	10.619	16.095	4
0.5%	1019	5	0	GET	^/api/posts/[^/]+/realmojis$	275.512	0.270	1.402	2.118	2371
0.4%	1425	0	0	GET	^/api/relationships/[^/]+	253.587	0.178	1.276	1.750	1478
0.4%	2404	43	0	GET	^/api/posts/[^/]+$	249.223	0.104	0.767	1.760	619
0.4%	357	3	0	GET	/api/me	217.231	0.608	1.602	1.648	299
0.3%	730	7	0	POST	^/api/posts/[^/]+/realmojis$	180.536	0.247	2.373	3.356	468
0.3%	506	0	0	GET	^/api/memories/[^/]+	163.372	0.323	1.611	2.062	677
0.2%	1177	1	0	GET	^/api/users/[^/]+$	123.047	0.105	0.777	1.322	204
0.2%	3546	8	0	POST	^/api/friend-requests/[^/]+/[^/]+$	120.236	0.034	0.574	1.092	116
0.2%	3735	62	0	POST	/api/friend-requests	113.063	0.030	0.356	1.004	114
0.2%	1026	4	0	GET	^/api/posts/[^/]+/comments$	110.626	0.108	0.684	1.593	3159
0.1%	3794	405	0	POST	/api/auth/login	54.645	0.014	0.392	1.104	221
0.1%	830	0	0	POST	^/api/posts/[^/]+/comments$	38.829	0.047	0.504	1.076	349
0.1%	87	0	0	DELETE	^/api/posts/[^/]+$	35.026	0.403	1.436	1.436	0
0.0%	159	0	0	DELETE	^/api/realmojis/[^/]+$	28.133	0.177	0.602	0.642	0
0.0%	1	0	0	POST	/api/initialize	21.147	21.147	21.147	21.147	90
0.0%	166	0	0	PATCH	^/api/posts/[^/]+$	20.786	0.125	0.979	0.987	652
0.0%	599	0	0	POST	/api/auth/logout	13.399	0.022	0.449	1.005	0
0.0%	182	0	0	DELETE	^/api/comments/[^/]+$	4.113	0.023	0.236	0.297	0
0.0%	318	0	0	PATCH	/api/me	3.201	0.010	0.146	1.127	270
0.0%	453	0	0	PUT	^/api/follows/[^/]+$	2.712	0.006	0.137	0.729	0
…（上位 30 / 全 44 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
84.7%	1123	0	0	GET	/api/notifications	30128.632	26.829	60.012	60.122	102
2.6%	8060	0	0	GET	^/api/feeds/[^/]+	922.470	0.114	1.098	1.825	5868
2.2%	7520	29	0	GET	^/api/images/[^/]+	778.207	0.103	0.955	2.887	176909
1.8%	2748	44	0	POST	/api/friend-requests	639.996	0.233	1.045	1.344	114
1.4%	1102	2	0	POST	/api/posts	509.641	0.462	1.904	7.063	651
0.9%	2034	0	0	POST	/api/auth/login	327.940	0.161	0.768	1.305	249
0.8%	647	10	0	POST	^/api/posts/[^/]+/realmojis$	292.032	0.451	2.748	3.895	477
0.8%	1836	0	0	GET	/api/memories	268.084	0.146	1.287	3.442	14057
0.6%	1601	0	0	GET	/api/posting-window/current	200.609	0.125	0.662	1.493	238
0.5%	1091	0	0	GET	^/api/posts/me/current$	181.537	0.166	0.935	1.673	652
0.5%	1318	0	0	GET	/api/users	177.676	0.135	0.505	1.257	239
0.4%	616	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	131.559	0.214	1.137	1.356	115
0.3%	1364	12	0	GET	^/api/posts/[^/]+$	115.595	0.085	0.324	1.501	619
0.3%	1005	0	0	GET	^/api/relationships/[^/]+	107.144	0.107	0.830	1.138	536
0.3%	246	0	0	PUT	/api/me/profile-image	95.018	0.386	2.199	2.945	0
0.3%	434	0	0	POST	/api/auth/register	91.206	0.210	0.435	0.457	229
0.2%	394	2	0	POST	^/api/posts/[^/]+/comments$	75.665	0.192	1.039	1.217	350
0.2%	727	1	0	GET	^/api/posts/[^/]+/realmojis$	61.109	0.084	0.456	1.138	3117
0.2%	694	0	0	GET	^/api/users/[^/]+$	60.996	0.088	0.460	0.842	204
0.2%	367	0	0	POST	/api/auth/logout	55.053	0.150	0.819	1.185	0
0.1%	246	0	0	PUT	^/api/follows/[^/]+$	53.094	0.216	0.552	0.983	0
0.1%	218	0	0	PATCH	/api/me	42.710	0.196	0.517	0.546	276
0.1%	119	0	0	PATCH	^/api/posts/[^/]+$	42.318	0.356	1.343	1.345	648
0.1%	450	0	0	GET	^/api/posts/[^/]+/comments$	40.651	0.090	0.353	0.510	3649
0.1%	127	0	0	DELETE	^/api/realmojis/[^/]+$	40.157	0.316	0.721	1.549	0
0.1%	52	0	0	DELETE	^/api/posts/[^/]+$	39.707	0.764	1.666	1.666	0
0.1%	253	2	0	GET	/api/me	35.545	0.140	0.388	0.965	301
0.1%	295	0	0	GET	^/api/memories/[^/]+	28.323	0.096	0.787	1.575	973
0.0%	110	0	0	GET	/api/me/export	17.610	0.160	1.849	4.078	2233553
0.0%	74	0	0	DELETE	^/api/comments/[^/]+$	9.633	0.130	0.390	0.390	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
85.2%	1077	0	0	GET	/api/notifications	28528.185	26.489	60.002	60.124	107
2.5%	6398	0	0	GET	^/api/feeds/[^/]+	831.099	0.130	1.096	1.766	4365
2.4%	7084	31	0	GET	^/api/images/[^/]+	805.523	0.114	0.966	2.687	168935
1.6%	1071	2	0	POST	/api/posts	537.490	0.502	2.064	7.039	660
1.6%	2505	35	0	POST	/api/friend-requests	523.671	0.209	1.064	1.305	114
1.0%	1896	0	0	POST	/api/auth/login	331.740	0.175	0.933	1.416	252
0.8%	1757	0	0	GET	/api/memories	254.726	0.145	1.316	2.173	15535
0.5%	347	7	0	POST	^/api/posts/[^/]+/realmojis$	173.107	0.499	2.747	2.983	462
0.5%	1058	0	0	GET	^/api/posts/me/current$	158.835	0.150	1.215	1.420	662
0.4%	1522	0	0	GET	/api/posting-window/current	147.842	0.097	0.407	1.402	240
0.4%	1205	0	0	GET	/api/users	146.727	0.122	0.652	1.497	238
0.4%	258	0	0	PUT	/api/me/profile-image	132.214	0.512	1.155	2.882	0
0.4%	592	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	119.570	0.202	1.168	1.359	116
0.3%	919	26	0	GET	^/api/posts/[^/]+$	97.643	0.106	0.598	1.215	609
0.3%	793	0	0	GET	^/api/relationships/[^/]+	95.577	0.121	0.873	1.370	691
0.3%	418	0	0	POST	/api/auth/register	84.089	0.201	0.478	0.534	230
0.2%	314	0	0	POST	^/api/posts/[^/]+/comments$	64.671	0.206	0.946	1.271	358
0.2%	284	0	0	PUT	^/api/follows/[^/]+$	57.926	0.204	0.415	1.316	0
0.2%	313	0	0	POST	/api/auth/logout	52.078	0.166	0.985	1.209	0
0.1%	447	0	0	GET	^/api/users/[^/]+$	45.537	0.102	0.483	1.129	203
0.1%	134	0	0	PATCH	^/api/posts/[^/]+$	45.025	0.336	1.256	1.263	662
0.1%	416	0	0	GET	^/api/posts/[^/]+/realmojis$	43.959	0.106	0.552	0.812	3988
0.1%	202	0	0	PATCH	/api/me	41.827	0.207	0.413	0.426	285
0.1%	332	0	0	GET	^/api/posts/[^/]+/comments$	34.307	0.103	0.580	0.825	3868
0.1%	50	0	0	DELETE	^/api/posts/[^/]+$	33.113	0.662	1.555	1.555	0
0.1%	259	0	0	GET	^/api/memories/[^/]+	28.215	0.109	1.009	1.853	1013
0.1%	263	1	0	GET	/api/me	27.435	0.104	0.275	0.302	303
0.1%	52	0	0	DELETE	^/api/realmojis/[^/]+$	18.734	0.360	1.344	1.344	0
0.1%	118	0	0	GET	/api/me/export	17.432	0.148	0.855	2.159	3174132
0.0%	52	0	0	DELETE	^/api/comments/[^/]+$	7.434	0.143	0.324	0.324	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 12141 クエリ / 48 種類 / 合計 126.158s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	39.7%	216	50.028	231.61	1092.13	1136.3	116.0	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	21.0%	4	26.522	6630.55	10116.42	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.8%	4	18.651	4662.64	7205.13	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	9.5%	7541	11.990	1.59	8.51	9.9	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	7.4%	4	9.392	2348.07	3429.44	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	5.0%	7	6.260	894.35	3989.61	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.5%	7	0.631	90.13	162.00	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.5%	7	0.625	89.28	168.17	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.3%	1	0.389	389.16	389.16	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.3%	1	0.386	386.13	386.13	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.2%	1	0.230	229.60	229.60	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.169	169.17	169.17	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.133	132.56	132.56	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
14	0.1%	1	0.109	109.16	109.16	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	7	0.099	14.19	28.19	167.4	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: 48adea8f435c46fd6b106e6a8ce4d3f346c64d11
Type: cpu
Time: 2026-10-31 16:47:15 JST
Duration: 60s, Total samples = 23.76s (39.60%)
Showing nodes accounting for 18.77s, 79.00% of 23.76s total
Dropped 725 nodes (cum <= 0.12s)
Showing top 40 nodes out of 221
      flat  flat%   sum%        cum   cum%
     4.34s 18.27% 18.27%      4.34s 18.27%  crypto/internal/fips140/sha256.blockAVX2
     2.17s  9.13% 27.40%      2.17s  9.13%  internal/runtime/syscall/linux.Syscall6
     1.62s  6.82% 34.22%      5.54s 23.32%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.35s  5.68% 39.90%      1.35s  5.68%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.17s  4.92% 44.82%      1.17s  4.92%  runtime.memmove
     1.05s  4.42% 49.24%      1.05s  4.42%  aeshashbody
     0.68s  2.86% 52.10%      0.68s  2.86%  runtime.memclrNoHeapPointers
     0.66s  2.78% 54.88%      2.07s  8.71%  runtime.scanObject
     0.54s  2.27% 57.15%      0.69s  2.90%  runtime.findObject
     0.41s  1.73% 58.88%      0.58s  2.44%  runtime.tryDeferToSpanScan
     0.36s  1.52% 60.40%      1.62s  6.82%  runtime.mapaccess1_faststr
     0.31s  1.30% 61.70%      0.31s  1.30%  runtime.nextFreeFast (inline)
     0.30s  1.26% 62.96%      0.47s  1.98%  example.com/isureal-company/isureal.postLess
     0.30s  1.26% 64.23%      1.06s  4.46%  runtime.mallocgcSmallScanNoHeader
     0.29s  1.22% 65.45%      0.29s  1.22%  sync/atomic.(*Int32).Add (inline)
     0.25s  1.05% 66.50%      4.88s 20.54%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.25s  1.05% 67.55%      0.25s  1.05%  memeqbody
     0.25s  1.05% 68.60%      0.25s  1.05%  runtime.futex
     0.22s  0.93% 69.53%      4.74s 19.95%  crypto/internal/fips140/sha256.(*Digest).Write
     0.21s  0.88% 70.41%      0.28s  1.18%  database/sql.convertAssignRows
     0.17s  0.72% 71.13%      1.52s  6.40%  runtime.mapassign_faststr
     0.16s  0.67% 71.80%      5.13s 21.59%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.16s  0.67% 72.47%      0.19s   0.8%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.16s  0.67% 73.15%      0.21s  0.88%  runtime.(*mspan).writeHeapBitsSmall
     0.15s  0.63% 73.78%      1.67s  7.03%  runtime.mallocgc
     0.13s  0.55% 74.33%      3.65s 15.36%  example.com/isureal-company/isureal.loadBig
     0.12s  0.51% 74.83%      1.15s  4.84%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.11s  0.46% 75.29%      0.23s  0.97%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.10s  0.42% 75.72%      5.37s 22.60%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.10s  0.42% 76.14%      0.12s  0.51%  time.Time.Before
     0.09s  0.38% 76.52%      0.14s  0.59%  internal/runtime/maps.(*Iter).Next
     0.08s  0.34% 76.85%      0.22s  0.93%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.07s  0.29% 77.15%      0.37s  1.56%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func3
     0.07s  0.29% 77.44%      0.54s  2.27%  runtime.convTstring
     0.07s  0.29% 77.74%      0.16s  0.67%  runtime.greyobject
     0.07s  0.29% 78.03%      0.33s  1.39%  runtime.mallocgcSmallNoscan
     0.07s  0.29% 78.32%      0.28s  1.18%  runtime.mapaccess2_faststr
     0.06s  0.25% 78.58%      1.89s  7.95%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.05s  0.21% 78.79%      4.39s 18.48%  crypto/internal/fips140/sha256.block
     0.05s  0.21% 79.00%      0.28s  1.18%  sort.insertionSort_func
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 24.9% / 100% | 7% | 0.8% | 5 | 157 MB | run queue が積もっている |
| isu2 (2 cpu) | 19.2% / 100% | 5% | 0.7% | 6 | 178 MB | run queue が積もっている |
| isu3 (2 cpu) | 19.8% / 100% | 7% | 0.9% | 6 | 101 MB | run queue が積もっている |
| isu4 (2 cpu) | 51.4% / 100% | 46% | 2.5% | 5 | 103 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 39% / nginx 8% |
| isu2 | isureal 34% / nginx 2% |
| isu3 | isureal 36% / nginx 2% |
| isu4 | isureal 83% / postgres 12% |

## Diff from previous run

スコア: 220540 → 230652 (+10112)

`15a1e6c..766568b`

```
 webapp/go/comment.go  |  2 +-
 webapp/go/feed_api.go | 36 +++++++++++++---------
 webapp/go/post.go     | 85 ++++++++++++++++++++++++++++++---------------------
 webapp/go/post_api.go |  3 +-
 4 files changed, 74 insertions(+), 52 deletions(-)
```
