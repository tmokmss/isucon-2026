# run 20261031-144534-f7ca385

## Summary

| | |
|---|---|
| score | 120563 |
| status | pass |
| ref | f7ca385 |
| sha | f7ca385d5ee0 |
| 所要時間 | 116s |
| 前回比 | -14862 (-11.0%) vs 20261031-143412-10c9074 |
| ベスト比 | -20772 (-14.7%) vs 20261031-141027-bd0fe45 = 141335 |
| note | DB外し4: 停止時スナップショット修正・毎秒syncなし |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	141	405	1007	1007
europe	7	17	42	117	239	238
america	6	15	52	133	164	163
```

```
{"id":"19069","status":"done","portalStatus":"finished","passed":true,"score":120563,"raw":120563,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:45:56.936Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:45:59.927Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:46:02.924Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:46:05.933Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T05:46:07.255Z level=INFO msg="user moved" user=notsu.5219 from=asia to=europe
time=2026-10-31T05:46:08.924Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T05:46:11.924Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T05:46:14.943Z level=INFO msg="asia 3 日目: 141 人"
time=2026-10-31T05:46:17.946Z level=INFO msg="europe 3 日目: 42 人"
time=2026-10-31T05:46:18.230Z level=INFO msg="user moved" user=iswa-6022 from=asia to=america
time=2026-10-31T05:46:18.999Z level=INFO msg="user moved" user=honoanu_6062 from=asia to=america
time=2026-10-31T05:46:19.167Z level=INFO msg="user moved" user=tetomu-6157 from=asia to=america
time=2026-10-31T05:46:20.105Z level=INFO msg="user moved" user=notsu.5219 from=europe to=america
time=2026-10-31T05:46:20.423Z level=INFO msg="user moved" user=honoanu_6062 from=america to=europe
time=2026-10-31T05:46:20.924Z level=INFO msg="america 3 日目: 52 人"
time=2026-10-31T05:46:23.956Z level=INFO msg="asia 4 日目: 405 人"
time=2026-10-31T05:46:26.124Z level=INFO msg="user moved" user=noane6665 from=asia to=america
time=2026-10-31T05:46:27.096Z level=INFO msg="europe 4 日目: 117 人"
time=2026-10-31T05:46:30.029Z level=INFO msg="user moved" user=noane6665 from=america to=asia
time=2026-10-31T05:46:30.274Z level=INFO msg="america 4 日目: 133 人"
time=2026-10-31T05:46:31.083Z level=INFO msg="user moved" user=ullia_7078 from=asia to=america
time=2026-10-31T05:46:31.085Z level=INFO msg="user moved" user=nishiso_7067 from=asia to=america
time=2026-10-31T05:46:31.704Z level=INFO msg="user moved" user=mekiel.8016 from=asia to=europe
time=2026-10-31T05:46:31.989Z level=INFO msg="user moved" user=honoanu_6062 from=europe to=america
time=2026-10-31T05:46:32.991Z level=INFO msg="asia 5 日目: 1007 人"
time=2026-10-31T05:46:35.798Z level=INFO msg="user moved" user=naorto.8812 from=europe to=asia
time=2026-10-31T05:46:36.089Z level=INFO msg="europe 5 日目: 239 人"
time=2026-10-31T05:46:36.402Z level=INFO msg="user moved" user=notsu.5219 from=america to=europe
time=2026-10-31T05:46:36.447Z level=INFO msg="user moved" user=annoa8511 from=europe to=america
time=2026-10-31T05:46:39.944Z level=INFO msg="america 5 日目: 164 人"
time=2026-10-31T05:46:52.585Z level=INFO msg="asia 最終日（グランプリ）: 1007 人"
time=2026-10-31T05:46:52.591Z level=INFO msg="europe 最終日（グランプリ）: 238 人"
time=2026-10-31T05:46:52.591Z level=INFO msg="america 最終日（グランプリ）: 163 人"
time=2026-10-31T05:46:57.008Z level=INFO msg="グランプリ: アジア 30 人, ヨーロッパ 76 人, アメリカ 78 人"
{
  "pass": true,
  "score": 120563,
  "messages": [
    "marni.5056: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
49.0%	2193	1	0	GET	/api/notifications	21442.488	9.778	20.641	20.646	128
14.1%	2045	50	0	POST	/api/posts	6164.991	3.015	10.077	11.430	633
8.1%	16295	469	0	GET	^/api/images/[^/]+	3559.313	0.218	3.195	6.708	105769
3.8%	692	130	0	PUT	/api/me/profile-image	1643.955	2.376	12.036	14.291	3
3.7%	9971	0	0	GET	^/api/feeds/[^/]+	1634.188	0.164	2.657	5.175	2902
3.7%	5152	46	0	POST	/api/friend-requests	1613.419	0.313	2.983	5.175	114
2.6%	6301	897	0	POST	/api/auth/login	1140.668	0.181	2.797	6.479	213
2.5%	3786	0	0	GET	/api/memories	1081.929	0.286	3.243	6.541	13204
1.8%	3303	0	0	GET	/api/posting-window/current	799.228	0.242	2.650	5.175	237
1.6%	2132	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	715.495	0.336	2.983	5.175	116
1.4%	1596	81	0	GET	^/api/posts/me/current$	599.326	0.376	3.990	5.175	616
1.3%	502	10	0	POST	^/api/posts/[^/]+/realmojis$	558.708	1.113	6.479	10.195	468
1.0%	2517	0	0	GET	/api/users	428.366	0.170	2.400	5.175	247
0.8%	368	0	0	GET	/api/me/export	358.827	0.975	5.559	11.530	1886759
0.5%	1526	46	0	GET	^/api/posts/[^/]+$	216.135	0.142	2.601	5.175	608
0.5%	1030	1	0	POST	/api/auth/register	206.618	0.201	0.800	2.742	227
0.5%	541	0	0	POST	^/api/posts/[^/]+/comments$	206.118	0.381	4.456	6.460	339
0.4%	1344	0	0	GET	^/api/relationships/[^/]+	195.848	0.146	2.486	3.142	1063
0.4%	872	0	0	POST	/api/auth/logout	159.847	0.183	3.142	5.208	0
0.3%	640	0	0	PUT	^/api/follows/[^/]+$	145.780	0.228	1.948	3.233	0
0.3%	842	0	0	GET	^/api/memories/[^/]+	142.891	0.170	2.626	3.233	670
0.3%	572	3	0	GET	/api/me	136.699	0.239	1.318	2.580	301
0.3%	461	0	0	PATCH	/api/me	130.046	0.282	1.384	5.174	275
0.3%	765	0	0	GET	^/api/users/[^/]+$	111.054	0.145	2.580	3.142	207
0.2%	157	29	0	PATCH	^/api/posts/[^/]+$	73.207	0.466	3.528	3.758	526
0.1%	100	0	0	DELETE	^/api/realmojis/[^/]+$	64.874	0.649	3.887	4.463	0
0.1%	70	0	0	DELETE	^/api/posts/[^/]+$	59.534	0.850	3.758	3.758	0
0.1%	550	0	0	GET	^/api/posts/[^/]+/comments$	55.245	0.100	1.806	2.909	1689
0.1%	658	0	0	GET	^/api/posts/[^/]+/realmojis$	49.757	0.076	0.847	2.909	1137
0.1%	103	0	0	DELETE	^/api/comments/[^/]+$	36.667	0.356	3.014	3.990	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
37.7%	436	1	0	GET	/api/notifications	3603.168	8.264	14.471	14.583	174
16.3%	2704	17	0	GET	^/api/images/[^/]+	1561.321	0.577	3.529	5.038	148690
12.9%	567	7	0	POST	/api/posts	1237.511	2.183	7.065	9.882	640
5.6%	1225	6	0	POST	/api/friend-requests	539.995	0.441	3.249	3.856	114
3.9%	1762	0	0	GET	^/api/feeds/[^/]+	371.728	0.211	1.147	2.963	2687
3.7%	870	0	0	POST	/api/auth/login	358.192	0.412	2.786	3.765	247
3.7%	127	0	0	PUT	/api/me/profile-image	353.815	2.786	8.941	8.943	0
2.6%	833	0	0	GET	/api/posting-window/current	249.729	0.300	1.509	3.075	239
2.3%	721	0	0	GET	/api/memories	215.971	0.300	2.963	3.230	13920
1.7%	591	0	0	GET	/api/users	158.011	0.267	1.463	3.072	248
1.6%	237	0	0	POST	/api/auth/register	152.549	0.644	1.882	1.929	229
1.4%	445	73	0	GET	^/api/posts/me/current$	133.964	0.301	2.967	3.049	542
1.3%	366	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	127.665	0.349	2.367	3.579	116
0.6%	260	0	0	GET	^/api/relationships/[^/]+	61.826	0.238	1.397	3.080	886
0.6%	148	0	0	PUT	^/api/follows/[^/]+$	60.863	0.411	1.833	1.887	0
0.6%	112	0	0	PATCH	/api/me	54.794	0.489	2.324	4.929	275
0.6%	132	2	0	GET	/api/me	54.047	0.409	1.079	1.479	299
0.5%	133	0	0	POST	/api/auth/logout	47.742	0.359	2.439	4.987	0
0.4%	234	16	0	GET	^/api/posts/[^/]+$	40.478	0.173	1.053	1.329	587
0.3%	47	0	0	GET	/api/me/export	32.414	0.690	5.831	5.831	1820325
0.3%	21	0	0	DELETE	^/api/posts/[^/]+$	24.730	1.178	3.484	3.484	0
0.3%	118	0	0	GET	^/api/memories/[^/]+	23.986	0.203	1.398	1.424	1122
0.2%	32	1	0	POST	^/api/posts/[^/]+/realmojis$	23.170	0.724	4.689	4.689	486
0.2%	36	6	0	PATCH	^/api/posts/[^/]+$	22.843	0.635	2.517	2.517	541
0.2%	131	0	0	GET	^/api/users/[^/]+$	22.303	0.170	0.949	1.148	203
0.2%	52	0	0	POST	^/api/posts/[^/]+/comments$	14.907	0.287	1.521	1.521	354
0.1%	66	0	0	GET	^/api/posts/[^/]+/comments$	8.998	0.136	1.038	1.038	1098
0.1%	77	0	0	GET	^/api/posts/[^/]+/realmojis$	7.295	0.095	0.769	0.769	894
0.0%	4	0	0	DELETE	^/api/realmojis/[^/]+$	3.135	0.784	1.522	1.522	0
0.0%	6	0	0	DELETE	^/api/comments/[^/]+$	1.381	0.230	0.560	0.560	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.7%	358	0	0	GET	/api/notifications	2864.464	8.001	11.053	11.055	177
18.0%	2459	12	0	GET	^/api/images/[^/]+	1331.386	0.541	3.227	4.178	133004
13.6%	469	6	0	POST	/api/posts	1004.690	2.142	6.865	10.077	646
5.8%	879	8	0	POST	/api/friend-requests	430.257	0.489	2.982	3.945	114
3.6%	724	0	0	POST	/api/auth/login	268.271	0.371	2.585	4.902	251
3.4%	1226	0	0	GET	^/api/feeds/[^/]+	255.009	0.208	1.127	2.232	2592
2.3%	658	0	0	GET	/api/posting-window/current	167.463	0.255	1.079	1.839	240
2.2%	680	0	0	GET	/api/memories	162.953	0.240	2.252	4.509	16735
1.9%	336	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	139.993	0.417	1.807	2.589	116
1.7%	409	0	0	GET	/api/users	128.280	0.314	1.100	1.449	248
1.7%	98	0	0	PUT	/api/me/profile-image	127.476	1.301	5.110	5.110	0
0.9%	395	48	0	GET	^/api/posts/me/current$	68.887	0.174	0.963	2.262	577
0.8%	157	0	0	POST	/api/auth/register	57.319	0.365	1.278	1.446	229
0.7%	121	0	0	POST	/api/auth/logout	53.043	0.438	3.201	3.675	0
0.7%	41	0	0	GET	/api/me/export	50.773	1.238	11.961	11.961	2958109
0.5%	177	0	0	GET	^/api/relationships/[^/]+	35.287	0.199	0.977	1.983	1504
0.5%	93	0	0	PUT	^/api/follows/[^/]+$	33.802	0.363	1.801	1.801	0
0.5%	72	0	0	PATCH	/api/me	33.612	0.467	6.071	6.071	286
0.4%	25	3	0	POST	^/api/posts/[^/]+/realmojis$	28.412	1.136	4.925	4.925	399
0.4%	20	0	0	DELETE	^/api/posts/[^/]+$	28.292	1.415	4.094	4.094	0
0.4%	103	1	0	GET	/api/me	27.348	0.266	1.185	1.449	299
0.3%	132	9	0	GET	^/api/posts/[^/]+$	23.223	0.176	0.968	0.992	573
0.3%	106	0	0	GET	^/api/memories/[^/]+	22.580	0.213	1.641	1.694	1181
0.3%	36	1	0	PATCH	^/api/posts/[^/]+$	19.655	0.546	1.730	1.730	628
0.2%	64	0	0	GET	^/api/users/[^/]+$	11.765	0.184	0.969	0.969	207
0.1%	17	0	0	POST	^/api/posts/[^/]+/comments$	10.878	0.640	1.882	1.882	344
0.1%	40	0	0	GET	^/api/posts/[^/]+/realmojis$	7.318	0.183	0.874	0.874	758
0.1%	5	0	0	DELETE	^/api/realmojis/[^/]+$	4.621	0.924	2.055	2.055	0
0.1%	28	0	0	GET	^/api/posts/[^/]+/comments$	4.120	0.147	0.933	0.933	1691
0.0%	1	0	0	DELETE	/api/me/profile-image	2.052	2.052	2.052	2.052	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 14021 クエリ / 48 種類 / 合計 158.268s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	57.2%	384	90.599	235.93	1162.04	1202.3	122.8	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	13.5%	3	21.307	7102.44	9592.94	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.2%	11925	17.745	1.49	6.37	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.0%	3	15.898	5299.41	7192.14	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	4.9%	3	7.787	2595.83	3656.30	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	0.8%	1	1.209	1209.17	1209.17	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
7	0.7%	6	1.068	178.03	215.42	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
8	0.4%	6	0.701	116.88	156.93	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.4%	6	0.688	114.71	157.74	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
10	0.2%	1	0.256	256.37	256.37	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.1%	1	0.220	220.29	220.29	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.137	136.53	136.53	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.107	107.05	107.05	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.097	96.90	96.90	2433.0	0.0	DELETE FROM image_objects WHERE created_at > (SELECT created_at FROM initial_data_boundary)
15	0.1%	1	0.082	81.63	81.63	2470.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT data FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 57c0cf2f6b39f4e9e94144eaf33011b6634cd117
Type: cpu
Time: 2026-10-31 14:45:36 JST
Duration: 60.14s, Total samples = 42.52s (70.70%)
Showing nodes accounting for 29.94s, 70.41% of 42.52s total
Dropped 1153 nodes (cum <= 0.21s)
Showing top 40 nodes out of 287
      flat  flat%   sum%        cum   cum%
     5.19s 12.21% 12.21%      5.19s 12.21%  crypto/internal/fips140/sha256.blockAVX2
     5.04s 11.85% 24.06%      5.04s 11.85%  internal/runtime/syscall/linux.Syscall6
     1.89s  4.44% 28.50%      1.89s  4.44%  runtime.memmove
     1.66s  3.90% 32.41%      2.20s  5.17%  runtime.findObject
     1.61s  3.79% 36.19%      1.61s  3.79%  runtime.memclrNoHeapPointers
     1.60s  3.76% 39.96%      1.60s  3.76%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.60s  3.76% 43.72%      6.34s 14.91%  runtime.scanObject
     1.52s  3.57% 47.30%      5.79s 13.62%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.30s  3.06% 50.35%      2.26s  5.32%  runtime.tryDeferToSpanScan
     1.29s  3.03% 53.39%      1.29s  3.03%  aeshashbody
     0.70s  1.65% 55.03%      0.70s  1.65%  runtime.spanClass.sizeclass (inline)
     0.60s  1.41% 56.44%      2.38s  5.60%  runtime.mapaccess1_faststr
     0.50s  1.18% 57.62%      0.50s  1.18%  memeqbody
     0.41s  0.96% 58.58%      0.57s  1.34%  example.com/isureal-company/isureal.postLess
     0.40s  0.94% 59.52%      0.40s  0.94%  runtime.nextFreeFast (inline)
     0.30s  0.71% 60.23%      5.15s 12.11%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.26s  0.61% 60.84%      0.51s  1.20%  runtime.greyobject
     0.26s  0.61% 61.45%      1.19s  2.80%  runtime.mallocgcSmallScanNoHeader
     0.26s  0.61% 62.06%      0.26s  0.61%  sync/atomic.(*Int32).Add
     0.25s  0.59% 62.65%      5.76s 13.55%  crypto/internal/fips140/sha256.(*Digest).Write
     0.25s  0.59% 63.24%      0.84s  1.98%  database/sql.convertAssignRows
     0.25s  0.59% 63.83%      0.25s  0.59%  runtime.futex
     0.25s  0.59% 64.42%      0.27s  0.63%  runtime.typePointers.next
     0.23s  0.54% 64.96%      1.86s  4.37%  runtime.mapassign_faststr
     0.22s  0.52% 65.48%      0.56s  1.32%  compress/flate.(*compressor).deflate
     0.21s  0.49% 65.97%         5s 11.76%  runtime.mallocgc
     0.20s  0.47% 66.44%      5.79s 13.62%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.20s  0.47% 66.91%      0.24s  0.56%  encoding/json.appendString[go.shape.string]
     0.20s  0.47% 67.38%      0.26s  0.61%  runtime.spanOf (inline)
     0.17s   0.4% 67.78%      0.29s  0.68%  internal/runtime/maps.(*Iter).Next
     0.16s  0.38% 68.16%      0.22s  0.52%  runtime.(*mspan).markBitsForIndex (inline)
     0.16s  0.38% 68.53%      0.59s  1.39%  runtime.mapaccess2_faststr
     0.14s  0.33% 68.86%      0.50s  1.18%  runtime.scanObjectsSmall
     0.10s  0.24% 69.10%      5.36s 12.61%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.10s  0.24% 69.33%      0.32s  0.75%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.10s  0.24% 69.57%      0.72s  1.69%  encoding/json.structEncoder.encode
     0.10s  0.24% 69.80%      0.41s  0.96%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
     0.09s  0.21% 70.01%      5.28s 12.42%  crypto/internal/fips140/sha256.block
     0.09s  0.21% 70.23%      0.78s  1.83%  runtime.mallocgcSmallNoscan
     0.08s  0.19% 70.41%      2.70s  6.35%  database/sql.(*Rows).nextLocked
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 49.6% / 100% | 36% | 5.7% | 9 | 141 MB | CPU が飽和 |
| isu2 (2 cpu) | 26.3% / 100% | 1% | 2.5% | 4 | 603 MB |  |
| isu3 (2 cpu) | 27.9% / 100% | 4% | 3.3% | 5 | 621 MB | run queue が積もっている |
| isu4 (2 cpu) | 36.0% / 99% | 13% | 20.1% | 5 | 105 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（IO 待ちが重い、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 71% / nginx 11% |
| isu2 | isureal 46% / nginx 1% |
| isu3 | isureal 48% / nginx 1% |
| isu4 | isureal 16% / postgres 14% |

## Diff from previous run

スコア: 135425 → 120563 (-14862)

`10c9074..f7ca385`

```
 webapp/go/cmd/isureal/main.go |  2 ++
 webapp/go/persist.go          | 18 +++++++++++++++---
 2 files changed, 17 insertions(+), 3 deletions(-)
```
