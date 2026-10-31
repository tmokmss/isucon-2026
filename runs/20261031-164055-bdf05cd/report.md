# run 20261031-164055-bdf05cd

## Summary

| | |
|---|---|
| score | 232110 |
| status | pass |
| ref | origin/main |
| sha | bdf05cd342f2 |
| 所要時間 | 170s |
| 前回比 | +11570 (5.2%) vs 20261031-163606-15a1e6c |
| ベスト比 | +8368 (3.7%) vs 20261031-163105-6f2cd6f = 223742 **ベスト更新** |
| note | main bdf05cd confirm |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	52	153	422	877	878
europe	7	21	64	189	541	540
america	6	20	61	178	492	490
```

```
{"id":"29679","status":"done","portalStatus":"finished","passed":true,"score":232110,"raw":232110,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:41:26.825Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:41:29.811Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:41:32.810Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:41:35.824Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T07:41:37.900Z level=INFO msg="user moved" user=konuvin-5225 from=asia to=america
time=2026-10-31T07:41:38.811Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T07:41:39.725Z level=INFO msg="user moved" user=konuvin-5225 from=america to=asia
time=2026-10-31T07:41:40.751Z level=INFO msg="user moved" user=shira_5293 from=asia to=europe
time=2026-10-31T07:41:41.811Z level=INFO msg="america 2 日目: 20 人"
time=2026-10-31T07:41:44.823Z level=INFO msg="asia 3 日目: 153 人"
time=2026-10-31T07:41:47.813Z level=INFO msg="europe 3 日目: 64 人"
time=2026-10-31T07:41:50.230Z level=INFO msg="user moved" user=konuvin-5225 from=asia to=america
time=2026-10-31T07:41:50.811Z level=INFO msg="america 3 日目: 61 人"
time=2026-10-31T07:41:53.247Z level=INFO msg="user moved" user=shira_5293 from=europe to=asia
time=2026-10-31T07:41:53.420Z level=INFO msg="user moved" user=evare_6479 from=europe to=asia
time=2026-10-31T07:41:53.823Z level=INFO msg="asia 4 日目: 422 人"
time=2026-10-31T07:41:56.812Z level=INFO msg="europe 4 日目: 189 人"
time=2026-10-31T07:41:58.555Z level=INFO msg="user moved" user=anyu-7852 from=asia to=europe
time=2026-10-31T07:41:58.975Z level=INFO msg="user moved" user=mukehe-7666 from=asia to=america
time=2026-10-31T07:41:59.586Z level=INFO msg="user moved" user=romo.8099 from=asia to=america
time=2026-10-31T07:41:59.776Z level=INFO msg="user moved" user=mukehe-7666 from=america to=asia
time=2026-10-31T07:41:59.817Z level=INFO msg="america 4 日目: 178 人"
time=2026-10-31T07:41:59.831Z level=INFO msg="user moved" user=nachieva.7941 from=asia to=europe
time=2026-10-31T07:42:00.031Z level=INFO msg="user moved" user=sunamar_7088 from=asia to=america
time=2026-10-31T07:42:00.852Z level=INFO msg="user moved" user=hako.7624 from=asia to=europe
time=2026-10-31T07:42:00.891Z level=INFO msg="user moved" user=miaul.8600 from=asia to=europe
time=2026-10-31T07:42:01.145Z level=INFO msg="user moved" user=evare_6479 from=asia to=america
time=2026-10-31T07:42:02.699Z level=INFO msg="user moved" user=iskiku.7113 from=asia to=europe
time=2026-10-31T07:42:02.826Z level=INFO msg="asia 5 日目: 877 人"
time=2026-10-31T07:42:05.651Z level=INFO msg="user moved" user=konoko_10926 from=america to=europe
time=2026-10-31T07:42:05.835Z level=INFO msg="europe 5 日目: 541 人"
time=2026-10-31T07:42:06.540Z level=INFO msg="user moved" user=nare-9898 from=america to=asia
time=2026-10-31T07:42:08.933Z level=INFO msg="america 5 日目: 492 人"
time=2026-10-31T07:42:09.439Z level=INFO msg="user moved" user=konuvin-5225 from=america to=asia
time=2026-10-31T07:42:17.836Z level=INFO msg="america 最終日（グランプリ）: 490 人"
time=2026-10-31T07:42:17.861Z level=INFO msg="europe 最終日（グランプリ）: 540 人"
time=2026-10-31T07:42:18.023Z level=INFO msg="asia 最終日（グランプリ）: 878 人"
time=2026-10-31T07:42:26.925Z level=INFO msg="グランプリ: アジア 506 人, ヨーロッパ 299 人, アメリカ 311 人"
{
  "pass": true,
  "score": 232110,
  "messages": [
    "ulniwa.5028: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "anmiare5098: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "noaul.15439: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "memeya_9470: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "norumax_12898: ログインし直せません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
79.3%	2424	1	0	GET	/api/notifications	66129.085	27.281	60.011	60.011	109
3.9%	12920	0	0	GET	^/api/feeds/[^/]+	3263.905	0.253	1.676	3.276	4494
3.4%	3086	0	0	GET	/api/memories	2813.107	0.912	8.168	12.189	10149
2.3%	2415	53	0	POST	/api/posts	1909.194	0.791	6.420	7.960	640
2.0%	3389	0	0	GET	/api/posting-window/current	1701.062	0.502	2.378	2.788	237
1.2%	2506	0	0	GET	/api/users	1009.685	0.403	1.905	3.110	243
1.0%	212	0	0	GET	/api/me/export	815.337	3.846	23.792	25.557	1963749
0.8%	1963	1	0	GET	^/api/posts/me/current$	684.210	0.349	1.313	1.868	654
0.8%	12788	264	0	GET	^/api/images/[^/]+	637.514	0.050	0.883	1.951	184358
0.7%	5206	69	0	POST	/api/friend-requests	614.259	0.118	0.712	1.344	114
0.7%	633	124	0	PUT	/api/me/profile-image	606.652	0.958	13.678	74.312	3
0.6%	521	3	0	GET	/api/me	486.022	0.933	2.664	2.690	300
0.6%	970	10	0	GET	^/api/posts/[^/]+/realmojis$	484.495	0.499	4.642	5.612	2229
0.4%	3911	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	353.275	0.090	0.726	1.293	116
0.4%	1577	0	0	GET	^/api/relationships/[^/]+	336.706	0.214	1.791	3.392	1440
0.4%	862	12	0	POST	^/api/posts/[^/]+/realmojis$	320.330	0.372	3.176	3.421	472
0.3%	2237	72	0	GET	^/api/posts/[^/]+$	286.176	0.128	0.760	1.158	606
0.2%	526	0	0	GET	^/api/memories/[^/]+	169.801	0.323	2.011	2.480	620
0.2%	1134	0	0	GET	^/api/users/[^/]+$	135.250	0.119	0.780	1.815	199
0.2%	4282	632	0	POST	/api/auth/login	133.053	0.031	0.584	1.391	211
0.1%	863	9	0	GET	^/api/posts/[^/]+/comments$	123.757	0.143	0.854	1.168	2999
0.1%	824	0	0	POST	^/api/posts/[^/]+/comments$	109.636	0.133	0.870	1.531	353
0.1%	98	0	0	DELETE	^/api/posts/[^/]+$	53.636	0.547	1.732	1.732	0
0.0%	125	0	0	DELETE	^/api/realmojis/[^/]+$	33.734	0.270	1.055	1.222	0
0.0%	217	0	0	PATCH	^/api/posts/[^/]+$	30.419	0.140	0.747	0.760	653
0.0%	586	0	0	POST	/api/auth/logout	27.889	0.048	0.666	0.770	0
0.0%	1	0	0	POST	/api/initialize	21.334	21.334	21.334	21.334	90
0.0%	586	0	0	PUT	^/api/follows/[^/]+$	18.781	0.032	0.415	0.542	0
0.0%	440	0	0	PATCH	/api/me	17.285	0.039	0.394	0.845	278
0.0%	156	0	0	DELETE	^/api/comments/[^/]+$	8.004	0.051	0.658	0.733	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
84.7%	1376	1	0	GET	/api/notifications	37260.519	27.079	60.001	60.517	104
2.3%	3066	27	0	POST	/api/friend-requests	1006.671	0.328	1.011	1.564	114
2.2%	7735	0	0	GET	^/api/feeds/[^/]+	950.661	0.123	0.973	1.705	3487
1.9%	6729	36	0	GET	^/api/images/[^/]+	823.871	0.122	1.123	2.845	184463
1.8%	1373	36	0	POST	/api/posts	810.321	0.590	7.066	7.942	637
0.9%	1892	0	0	POST	/api/auth/login	406.547	0.215	0.932	1.166	247
0.7%	1588	0	0	GET	/api/memories	305.819	0.193	1.448	2.202	13651
0.7%	311	0	0	PUT	/api/me/profile-image	287.285	0.924	3.507	4.100	0
0.6%	568	24	0	POST	^/api/posts/[^/]+/realmojis$	255.527	0.450	3.303	3.873	459
0.5%	1059	0	0	GET	^/api/posts/me/current$	205.034	0.194	1.402	1.448	653
0.4%	1965	0	0	GET	/api/posting-window/current	188.165	0.096	0.430	0.951	239
0.4%	711	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	181.471	0.255	1.094	1.648	116
0.4%	534	0	0	POST	/api/auth/register	179.852	0.337	0.924	0.985	229
0.4%	1477	0	0	GET	/api/users	170.600	0.116	0.697	1.448	246
0.3%	449	3	0	POST	^/api/posts/[^/]+/comments$	120.999	0.269	1.124	1.392	348
0.3%	350	0	0	PUT	^/api/follows/[^/]+$	119.010	0.340	1.025	1.402	0
0.3%	1209	30	0	GET	^/api/posts/[^/]+$	115.392	0.095	0.426	1.443	615
0.2%	878	0	0	GET	^/api/relationships/[^/]+	97.050	0.111	0.983	1.449	588
0.2%	272	0	0	PATCH	/api/me	96.747	0.356	0.949	0.970	279
0.1%	317	0	0	POST	/api/auth/logout	61.514	0.194	0.937	1.066	0
0.1%	591	0	0	GET	^/api/users/[^/]+$	55.835	0.094	0.274	1.534	203
0.1%	580	0	0	GET	^/api/posts/[^/]+/realmojis$	52.796	0.091	0.351	1.129	3303
0.1%	435	0	0	GET	^/api/posts/[^/]+/comments$	41.180	0.095	0.480	1.037	3711
0.1%	316	2	0	GET	/api/me	37.037	0.117	0.522	0.726	303
0.1%	117	2	0	PATCH	^/api/posts/[^/]+$	34.010	0.291	0.908	0.930	635
0.1%	47	0	0	DELETE	^/api/posts/[^/]+$	32.673	0.695	2.153	2.153	0
0.1%	72	0	0	DELETE	^/api/realmojis/[^/]+$	27.566	0.383	1.330	1.330	0
0.1%	265	0	0	GET	^/api/memories/[^/]+	26.876	0.101	0.533	1.738	861
0.0%	96	0	0	GET	/api/me/export	17.155	0.179	2.088	2.088	2136818
0.0%	76	0	0	DELETE	^/api/comments/[^/]+$	14.134	0.186	0.852	0.852	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
85.7%	1253	0	0	GET	/api/notifications	33232.998	26.523	60.416	60.870	108
2.1%	1182	26	0	POST	/api/posts	803.039	0.679	7.399	7.815	647
1.9%	5727	26	0	GET	^/api/images/[^/]+	735.362	0.128	1.315	2.351	150110
1.8%	4257	0	0	GET	^/api/feeds/[^/]+	690.057	0.162	1.268	1.593	3059
1.6%	2880	23	0	POST	/api/friend-requests	626.748	0.218	1.086	1.701	114
1.2%	1698	0	0	POST	/api/auth/login	475.152	0.280	1.199	1.349	249
0.7%	489	0	0	POST	/api/auth/register	283.152	0.579	1.390	1.422	230
0.6%	1496	0	0	GET	/api/memories	245.016	0.164	1.388	2.227	15330
0.6%	1758	0	0	GET	/api/posting-window/current	220.788	0.126	1.241	1.358	240
0.5%	293	0	0	PUT	/api/me/profile-image	201.683	0.688	3.243	3.583	0
0.5%	1395	0	0	GET	/api/users	183.157	0.131	1.255	1.390	245
0.4%	948	18	0	GET	^/api/posts/me/current$	167.718	0.177	1.305	1.557	648
0.3%	558	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	134.620	0.241	1.055	1.551	116
0.3%	297	20	0	POST	^/api/posts/[^/]+/realmojis$	118.355	0.399	2.191	2.471	435
0.2%	621	0	0	GET	^/api/relationships/[^/]+	94.218	0.152	1.251	1.385	668
0.2%	232	0	0	PATCH	/api/me	81.480	0.351	1.310	1.489	287
0.2%	282	0	0	PUT	^/api/follows/[^/]+$	70.561	0.250	1.112	1.487	0
0.2%	497	13	0	GET	^/api/posts/[^/]+$	68.278	0.137	1.354	1.499	607
0.2%	227	10	0	POST	^/api/posts/[^/]+/comments$	60.752	0.268	1.313	1.427	337
0.1%	256	0	0	POST	/api/auth/logout	55.058	0.215	1.076	1.160	0
0.1%	266	0	0	GET	^/api/users/[^/]+$	34.451	0.130	1.293	1.359	198
0.1%	236	1	0	GET	^/api/posts/[^/]+/realmojis$	30.858	0.131	1.257	1.336	2432
0.1%	96	12	0	PATCH	^/api/posts/[^/]+$	29.975	0.312	1.004	1.004	580
0.1%	300	1	0	GET	/api/me	29.180	0.097	0.796	0.835	304
0.1%	45	0	0	DELETE	^/api/realmojis/[^/]+$	23.559	0.524	1.669	1.669	0
0.1%	186	1	0	GET	^/api/posts/[^/]+/comments$	21.317	0.115	1.340	1.392	2825
0.1%	101	0	0	GET	/api/me/export	19.854	0.197	1.559	2.924	2800351
0.0%	193	0	0	GET	^/api/memories/[^/]+	19.020	0.099	0.537	1.563	984
0.0%	32	0	0	DELETE	^/api/posts/[^/]+$	17.993	0.562	1.550	1.550	0
0.0%	36	0	0	DELETE	^/api/comments/[^/]+$	8.589	0.239	1.268	1.268	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 12189 クエリ / 48 種類 / 合計 131.429s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	42.3%	217	55.563	256.05	914.81	1261.5	128.2	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	19.5%	4	25.637	6409.15	9555.87	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.3%	4	18.827	4706.82	7205.30	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.7%	7253	11.384	1.57	8.61	10.3	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.9%	4	9.013	2253.20	3452.63	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	5.0%	8	6.517	814.65	3570.03	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.9%	8	1.219	152.40	624.40	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.6%	8	0.748	93.52	161.97	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.5%	1	0.659	659.18	659.18	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
10	0.4%	1	0.492	491.67	491.67	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.233	233.12	233.12	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.179	178.79	178.79	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
13	0.1%	1	0.167	167.43	167.43	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	8	0.126	15.72	22.09	166.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	0.1%	1	0.115	114.92	114.92	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: 8aa4218bb6414425fc1c68325383b1e873e2dcf9
Type: cpu
Time: 2026-10-31 16:41:00 JST
Duration: 60s, Total samples = 27.08s (45.13%)
Showing nodes accounting for 21.06s, 77.77% of 27.08s total
Dropped 828 nodes (cum <= 0.14s)
Showing top 40 nodes out of 229
      flat  flat%   sum%        cum   cum%
     4.43s 16.36% 16.36%      4.43s 16.36%  crypto/internal/fips140/sha256.blockAVX2
     3.04s 11.23% 27.58%      3.04s 11.23%  internal/runtime/syscall/linux.Syscall6
     1.39s  5.13% 32.72%      5.62s 20.75%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.39s  5.13% 37.85%      1.39s  5.13%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.17s  4.32% 42.17%      1.17s  4.32%  runtime.memmove
     1.02s  3.77% 45.94%      1.47s  5.43%  example.com/isureal-company/isureal/jpegcheck.(*decoder).skipAC
        1s  3.69% 49.63%         1s  3.69%  aeshashbody
     0.70s  2.58% 52.22%      0.70s  2.58%  runtime.memclrNoHeapPointers
     0.50s  1.85% 54.06%      0.76s  2.81%  runtime.tryDeferToSpanScan
     0.49s  1.81% 55.87%      0.59s  2.18%  runtime.findObject
     0.45s  1.66% 57.53%      1.90s  7.02%  runtime.mapaccess1_faststr
     0.41s  1.51% 59.05%      0.42s  1.55%  example.com/isureal-company/isureal/jpegcheck.(*decoder).prefetch
     0.40s  1.48% 60.52%      2.01s  7.42%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.40s  1.48% 62.00%      1.92s  7.09%  runtime.scanObject
     0.36s  1.33% 63.33%      0.36s  1.33%  runtime.futex
     0.35s  1.29% 64.62%      0.53s  1.96%  example.com/isureal-company/isureal.postLess
     0.26s  0.96% 65.58%      4.98s 18.39%  crypto/internal/fips140/sha256.(*Digest).Write
     0.25s  0.92% 66.51%      0.25s  0.92%  memeqbody
     0.25s  0.92% 67.43%      0.25s  0.92%  runtime.nextFreeFast (inline)
     0.23s  0.85% 68.28%      5.11s 18.87%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.22s  0.81% 69.09%      0.30s  1.11%  database/sql.convertAssignRows
     0.21s  0.78% 69.87%      5.40s 19.94%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.21s  0.78% 70.64%      0.74s  2.73%  runtime.mallocgcSmallScanNoHeader
     0.18s  0.66% 71.31%      0.20s  0.74%  example.com/isureal-company/isureal.rmLess
     0.18s  0.66% 71.97%      0.18s  0.66%  sync/atomic.(*Int32).Add (inline)
     0.17s  0.63% 72.60%      0.20s  0.74%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.17s  0.63% 73.23%      0.17s  0.63%  runtime.spanClass.sizeclass (inline)
     0.14s  0.52% 73.74%      1.48s  5.47%  runtime.mallocgc
     0.14s  0.52% 74.26%      1.37s  5.06%  runtime.mapassign_faststr
     0.12s  0.44% 74.70%      4.55s 16.80%  crypto/internal/fips140/sha256.block
     0.10s  0.37% 75.07%      5.71s 21.09%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.10s  0.37% 75.44%      0.29s  1.07%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func12
     0.10s  0.37% 75.81%      1.10s  4.06%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.09s  0.33% 76.14%      0.29s  1.07%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.09s  0.33% 76.48%      1.92s  7.09%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.09s  0.33% 76.81%      0.19s   0.7%  runtime.greyobject
     0.07s  0.26% 77.07%      3.51s 12.96%  example.com/isureal-company/isureal.loadBig
     0.07s  0.26% 77.33%      0.14s  0.52%  internal/runtime/maps.(*Iter).Next
     0.06s  0.22% 77.55%      5.98s 22.08%  example.com/isureal-company/isureal.computeImageSignature
     0.06s  0.22% 77.77%      0.27s     1%  runtime.mapaccess2_faststr
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 26.2% / 100% | 9% | 0.7% | 7 | 124 MB | run queue が積もっている |
| isu2 (2 cpu) | 19.3% / 100% | 6% | 0.7% | 4 | 105 MB |  |
| isu3 (2 cpu) | 19.1% / 100% | 6% | 0.7% | 5 | 104 MB | run queue が積もっている |
| isu4 (2 cpu) | 49.4% / 100% | 43% | 1.9% | 7 | 104 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 41% / nginx 8% |
| isu2 | isureal 34% / nginx 2% |
| isu3 | isureal 35% / nginx 2% |
| isu4 | isureal 81% / postgres 11% |

## Diff from previous run

スコア: 220540 → 232110 (+11570)

`15a1e6c..bdf05cd`

アプリと conf/ に差分は無い（ハーネスなど、配らないものだけが変わった）。
