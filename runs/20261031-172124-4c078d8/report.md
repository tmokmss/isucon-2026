# run 20261031-172124-4c078d8

## Summary

| | |
|---|---|
| score | 223434 |
| status | pass |
| ref | 4c078d8 |
| sha | 4c078d8c5a8c |
| 所要時間 | 249s |
| 前回比 | +3705 (1.7%) vs 20261031-171010-53bad13 |
| ベスト比 | -8676 (-3.7%) vs 20261031-164055-bdf05cd = 232110 |
| note | presign-wide EU/US precompute next-minute signatures, 2-gen exp cache |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	143	385	773	772
europe	7	19	56	165	473	473
america	6	16	46	122	336	335
```

```
{"id":"33949","status":"done","portalStatus":"finished","passed":true,"score":223434,"raw":223434,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T08:23:35.698Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T08:23:38.690Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T08:23:41.687Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T08:23:44.698Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T08:23:47.687Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T08:23:48.337Z level=INFO msg="user moved" user=mosaso5351 from=asia to=america
time=2026-10-31T08:23:48.956Z level=INFO msg="user moved" user=mosaso5351 from=america to=europe
time=2026-10-31T08:23:50.687Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T08:23:53.710Z level=INFO msg="asia 3 日目: 143 人"
time=2026-10-31T08:23:56.687Z level=INFO msg="europe 3 日目: 56 人"
time=2026-10-31T08:23:57.907Z level=INFO msg="user moved" user=mosaso5351 from=europe to=america
time=2026-10-31T08:23:59.568Z level=INFO msg="user moved" user=miaeleva-5926 from=asia to=america
time=2026-10-31T08:23:59.687Z level=INFO msg="america 3 日目: 46 人"
time=2026-10-31T08:24:02.722Z level=INFO msg="asia 4 日目: 385 人"
time=2026-10-31T08:24:05.688Z level=INFO msg="europe 4 日目: 165 人"
time=2026-10-31T08:24:07.604Z level=INFO msg="user moved" user=liamaxko_7046 from=asia to=europe
time=2026-10-31T08:24:08.723Z level=INFO msg="user moved" user=wafutor8028 from=asia to=america
time=2026-10-31T08:24:08.743Z level=INFO msg="america 4 日目: 122 人"
time=2026-10-31T08:24:10.311Z level=INFO msg="user moved" user=ivyso8091 from=europe to=america
time=2026-10-31T08:24:10.446Z level=INFO msg="user moved" user=mosaso5351 from=america to=europe
time=2026-10-31T08:24:10.857Z level=INFO msg="user moved" user=maxkiul-6701 from=asia to=europe
time=2026-10-31T08:24:10.920Z level=INFO msg="user moved" user=liakaben7223 from=asia to=america
time=2026-10-31T08:24:11.705Z level=INFO msg="asia 5 日目: 773 人"
time=2026-10-31T08:24:14.821Z level=INFO msg="europe 5 日目: 473 人"
time=2026-10-31T08:24:15.296Z level=INFO msg="user moved" user=miaeleva-5926 from=america to=europe
time=2026-10-31T08:24:17.812Z level=INFO msg="america 5 日目: 336 人"
time=2026-10-31T08:24:26.788Z level=INFO msg="europe 最終日（グランプリ）: 473 人"
time=2026-10-31T08:24:26.792Z level=INFO msg="asia 最終日（グランプリ）: 772 人"
time=2026-10-31T08:24:26.892Z level=INFO msg="america 最終日（グランプリ）: 335 人"
time=2026-10-31T08:24:35.757Z level=INFO msg="グランプリ: アジア 515 人, ヨーロッパ 399 人, アメリカ 321 人"
{
  "pass": true,
  "score": 223434,
  "messages": [
    "hemar-5034: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "hina_nakamura_2667: 履歴を読めません: GET /api/memories?month=2026-10: 応答がタイムアウトしました",
    "yure_5248: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "manoa.9026: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maxzoemia.8711: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "elsolben14647: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "naru9041: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.35.75:48036-\u003e52.69.202.109:8080: write tcp 10.27.35.75:48036-\u003e52.69.202.109:8080: use of closed network connection",
    "marwa-14415: ログインし直せません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "vintsu_14223: ログインし直せません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "manuela_almeida_2494: ログアウトできません: POST /api/auth/logout: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/logout\": EOF",
    "hatorchi.8664: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "seoyeon_lee_4917: ログアウトできません: POST /api/auth/logout: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/logout\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
77.1%	2160	1	0	GET	/api/notifications	59485.753	27.540	60.004	60.008	112
4.6%	12673	0	0	GET	^/api/feeds/[^/]+	3536.476	0.279	1.960	3.312	5312
4.4%	2170	77	0	POST	/api/posts	3392.965	1.564	64.426	66.455	631
3.6%	2865	0	0	GET	/api/memories	2744.152	0.958	9.713	15.296	9791
1.8%	2960	0	0	GET	/api/posting-window/current	1417.177	0.479	2.397	3.273	237
1.4%	2192	0	0	GET	/api/users	1049.776	0.479	2.364	3.039	254
1.2%	182	0	0	GET	/api/me/export	890.616	4.893	31.011	31.265	2029176
1.1%	1851	1	0	GET	^/api/posts/me/current$	869.671	0.470	2.277	3.141	654
0.7%	952	1	0	GET	^/api/posts/[^/]+/realmojis$	549.995	0.578	3.694	4.550	2491
0.6%	11522	238	0	GET	^/api/images/[^/]+	446.902	0.039	0.782	1.307	204608
0.6%	596	126	0	PUT	/api/me/profile-image	434.997	0.730	11.345	14.459	4
0.5%	486	3	0	GET	/api/me	387.248	0.797	2.394	3.272	300
0.5%	1543	0	0	GET	^/api/relationships/[^/]+	373.039	0.242	1.580	4.587	1308
0.4%	2249	54	0	GET	^/api/posts/[^/]+$	297.217	0.132	0.760	1.088	611
0.3%	913	11	0	POST	^/api/posts/[^/]+/realmojis$	233.437	0.256	2.925	3.298	471
0.2%	4800	55	0	POST	/api/friend-requests	180.465	0.038	0.552	1.065	114
0.2%	966	4	0	GET	^/api/posts/[^/]+/comments$	173.608	0.180	1.437	1.876	3630
0.2%	437	0	0	GET	^/api/memories/[^/]+	165.331	0.378	2.245	3.050	697
0.2%	1186	0	0	GET	^/api/users/[^/]+$	141.950	0.120	0.681	1.666	203
0.2%	3541	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	130.084	0.037	0.656	1.102	116
0.1%	3634	505	0	POST	/api/auth/login	72.179	0.020	0.585	1.073	212
0.1%	1034	2	0	POST	^/api/posts/[^/]+/comments$	47.783	0.046	0.633	1.004	352
0.0%	82	0	0	DELETE	^/api/posts/[^/]+$	34.124	0.416	1.395	1.395	0
0.0%	200	1	0	PATCH	^/api/posts/[^/]+$	29.886	0.149	0.767	0.781	651
0.0%	141	0	0	DELETE	^/api/realmojis/[^/]+$	26.163	0.186	0.440	0.464	0
0.0%	1	0	0	POST	/api/initialize	18.979	18.979	18.979	18.979	90
0.0%	512	0	0	POST	/api/auth/logout	11.684	0.023	0.502	0.653	0
0.0%	475	0	0	PUT	^/api/follows/[^/]+$	3.747	0.008	0.145	0.541	0
0.0%	202	0	0	DELETE	^/api/comments/[^/]+$	3.681	0.018	0.234	0.660	0
0.0%	11	0	0	DELETE	/api/me/profile-image	3.255	0.296	0.928	0.928	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
83.5%	1203	0	0	GET	/api/notifications	33240.873	27.632	60.008	60.282	106
3.9%	1209	26	0	POST	/api/posts	1556.635	1.288	64.563	66.054	641
2.3%	8239	0	0	GET	^/api/feeds/[^/]+	903.635	0.110	0.752	1.319	4677
2.0%	7499	32	0	GET	^/api/images/[^/]+	787.095	0.105	0.941	2.019	174911
1.5%	2833	47	0	POST	/api/friend-requests	614.329	0.217	0.882	1.463	114
1.0%	1793	0	0	GET	/api/memories	379.361	0.212	1.248	1.660	13768
0.9%	2000	0	0	POST	/api/auth/login	348.153	0.174	0.863	1.298	248
0.6%	1083	0	0	GET	^/api/posts/me/current$	219.768	0.203	1.180	1.250	654
0.5%	1709	0	0	GET	/api/posting-window/current	196.496	0.115	0.584	0.974	239
0.4%	489	2	0	POST	^/api/posts/[^/]+/realmojis$	178.245	0.365	2.930	3.713	472
0.4%	1315	0	0	GET	/api/users	162.383	0.123	0.600	1.242	258
0.4%	800	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	157.446	0.197	0.983	1.291	116
0.3%	1013	0	0	GET	^/api/relationships/[^/]+	131.268	0.130	0.694	1.330	639
0.3%	1415	17	0	GET	^/api/posts/[^/]+$	118.838	0.084	0.567	0.756	619
0.3%	469	0	0	POST	/api/auth/register	100.406	0.214	0.460	0.480	228
0.2%	270	0	0	PUT	/api/me/profile-image	93.372	0.346	0.735	3.461	0
0.2%	334	0	0	PUT	^/api/follows/[^/]+$	71.487	0.214	0.534	0.737	0
0.2%	352	0	0	POST	^/api/posts/[^/]+/comments$	61.893	0.176	0.741	1.086	332
0.2%	751	0	0	GET	^/api/posts/[^/]+/realmojis$	61.454	0.082	0.524	0.790	3462
0.1%	705	0	0	GET	^/api/users/[^/]+$	58.275	0.083	0.385	0.655	200
0.1%	367	0	0	POST	/api/auth/logout	57.898	0.158	0.790	1.247	0
0.1%	253	0	0	PATCH	/api/me	54.255	0.214	0.513	0.550	279
0.1%	132	0	0	PATCH	^/api/posts/[^/]+$	47.111	0.357	1.354	1.589	657
0.1%	557	0	0	GET	^/api/posts/[^/]+/comments$	44.007	0.079	0.308	0.603	4425
0.1%	276	2	0	GET	/api/me	33.646	0.122	0.445	0.468	303
0.1%	105	0	0	DELETE	^/api/realmojis/[^/]+$	33.391	0.318	0.870	1.998	0
0.1%	48	0	0	DELETE	^/api/posts/[^/]+$	30.280	0.631	1.999	1.999	0
0.1%	291	0	0	GET	^/api/memories/[^/]+	28.396	0.098	0.974	1.487	911
0.0%	106	0	0	GET	/api/me/export	15.372	0.145	0.759	1.194	2225434
0.0%	72	0	0	DELETE	^/api/comments/[^/]+$	10.333	0.144	0.567	0.567	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
86.4%	871	1	0	GET	/api/notifications	23234.791	26.676	60.001	60.115	109
2.2%	5691	29	0	GET	^/api/images/[^/]+	582.147	0.102	0.937	2.204	156843
2.0%	852	6	0	POST	/api/posts	550.182	0.646	1.915	65.413	653
2.0%	4876	0	0	GET	^/api/feeds/[^/]+	545.444	0.112	0.819	1.626	3294
1.3%	2041	29	0	POST	/api/friend-requests	350.872	0.172	0.704	1.452	114
1.0%	1626	0	0	POST	/api/auth/login	278.735	0.171	0.770	1.401	252
0.7%	1440	0	0	GET	/api/memories	197.115	0.137	1.333	2.209	17483
0.5%	824	0	0	GET	^/api/posts/me/current$	142.715	0.173	1.292	1.372	657
0.5%	1209	0	0	GET	/api/posting-window/current	121.859	0.101	0.516	1.075	240
0.4%	1010	0	0	GET	/api/users	110.106	0.109	0.589	1.380	251
0.4%	283	8	0	POST	^/api/posts/[^/]+/realmojis$	109.366	0.386	2.602	2.619	463
0.3%	332	0	0	POST	/api/auth/register	72.303	0.218	0.507	0.682	230
0.2%	387	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	66.827	0.173	0.757	1.320	116
0.2%	188	0	0	PUT	/api/me/profile-image	62.764	0.334	2.824	2.934	0
0.2%	624	0	0	GET	^/api/relationships/[^/]+	62.211	0.100	0.662	1.161	522
0.2%	665	31	0	GET	^/api/posts/[^/]+$	57.823	0.087	0.425	1.175	596
0.2%	268	0	0	POST	/api/auth/logout	44.779	0.167	1.025	1.262	0
0.2%	257	2	0	POST	^/api/posts/[^/]+/comments$	41.814	0.163	0.730	1.269	336
0.1%	211	0	0	PUT	^/api/follows/[^/]+$	35.639	0.169	0.625	0.733	0
0.1%	253	0	0	GET	^/api/memories/[^/]+	27.913	0.110	0.688	1.118	1072
0.1%	315	0	0	GET	^/api/users/[^/]+$	27.804	0.088	0.427	0.691	204
0.1%	170	0	0	PATCH	/api/me	27.800	0.164	0.408	0.436	280
0.1%	301	0	0	GET	^/api/posts/[^/]+/realmojis$	27.402	0.091	0.486	1.177	4557
0.1%	97	0	0	PATCH	^/api/posts/[^/]+$	25.956	0.268	1.426	1.426	666
0.1%	285	0	0	GET	^/api/posts/[^/]+/comments$	24.222	0.085	0.322	0.666	6241
0.1%	42	0	0	DELETE	^/api/posts/[^/]+$	22.484	0.535	1.573	1.573	0
0.1%	194	1	0	GET	/api/me	21.219	0.109	0.522	0.522	302
0.1%	89	0	0	GET	/api/me/export	13.944	0.157	1.203	1.203	3092888
0.0%	38	0	0	DELETE	^/api/realmojis/[^/]+$	10.461	0.275	0.749	0.749	0
0.0%	39	0	0	DELETE	^/api/comments/[^/]+$	5.799	0.149	0.954	0.954	0
…（上位 30 / 全 36 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 10385 クエリ / 48 種類 / 合計 127.862s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	40.2%	189	51.368	271.79	785.13	1300.3	132.1	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	20.7%	4	26.444	6610.96	10015.79	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	15.1%	4	19.353	4838.19	7217.89	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.0%	6210	10.261	1.65	7.30	10.1	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	7.1%	4	9.049	2262.33	3413.03	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	6.3%	6	8.009	1334.91	4258.29	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.6%	1	0.797	797.00	797.00	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
8	0.3%	6	0.444	73.92	159.43	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.3%	6	0.417	69.54	157.62	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.3%	1	0.363	362.62	362.62	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.3%	1	0.355	355.14	355.14	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.165	164.78	164.78	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.134	133.93	133.93	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.100	99.71	99.71	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows
15	0.1%	3759	0.092	0.02	0.78	0.9	31.3	SELECT ip, banned_until FROM ip_bans WHERE banned_until > clock_timestamp()

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
Build ID: 83cdb01db1b56b7d5c0b801389319fd68aad44c5
Type: cpu
Time: 2026-10-31 17:23:12 JST
Duration: 60s, Total samples = 29.54s (49.23%)
Showing nodes accounting for 22.41s, 75.86% of 29.54s total
Dropped 920 nodes (cum <= 0.15s)
Showing top 40 nodes out of 244
      flat  flat%   sum%        cum   cum%
     4.47s 15.13% 15.13%      4.47s 15.13%  crypto/internal/fips140/sha256.blockAVX2
     3.11s 10.53% 25.66%      3.11s 10.53%  internal/runtime/syscall/linux.Syscall6
     1.64s  5.55% 31.21%      5.76s 19.50%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.43s  4.84% 36.05%      1.43s  4.84%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.35s  4.57% 40.62%      1.35s  4.57%  runtime.memmove
     1.15s  3.89% 44.52%      1.15s  3.89%  aeshashbody
     0.88s  2.98% 47.49%      0.88s  2.98%  runtime.memclrNoHeapPointers
     0.76s  2.57% 50.07%      1.17s  3.96%  runtime.tryDeferToSpanScan
     0.66s  2.23% 52.30%      2.95s  9.99%  runtime.scanObject
     0.64s  2.17% 54.47%      0.88s  2.98%  runtime.findObject
     0.63s  2.13% 56.60%      0.89s  3.01%  example.com/isureal-company/isureal/jpegcheck.(*decoder).skipAC
     0.41s  1.39% 57.99%      1.78s  6.03%  runtime.mapaccess1_faststr
     0.36s  1.22% 59.21%      0.36s  1.22%  memeqbody
     0.35s  1.18% 60.39%      0.59s  2.00%  example.com/isureal-company/isureal.postLess
     0.33s  1.12% 61.51%      5.15s 17.43%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.33s  1.12% 62.63%      0.51s  1.73%  database/sql.convertAssignRows
     0.31s  1.05% 63.68%      5.03s 17.03%  crypto/internal/fips140/sha256.(*Digest).Write
     0.30s  1.02% 64.69%      1.29s  4.37%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.27s  0.91% 65.61%      0.27s  0.91%  runtime.nextFreeFast (inline)
     0.26s  0.88% 66.49%      0.26s  0.88%  example.com/isureal-company/isureal/jpegcheck.(*decoder).prefetch
     0.23s  0.78% 67.26%      0.23s  0.78%  runtime.spanClass.sizeclass (inline)
     0.22s  0.74% 68.01%      0.22s  0.74%  runtime.futex
     0.22s  0.74% 68.75%      1.96s  6.64%  runtime.mallocgc
     0.21s  0.71% 69.47%      0.76s  2.57%  runtime.mallocgcSmallScanNoHeader
     0.20s  0.68% 70.14%      1.74s  5.89%  runtime.mapassign_faststr
     0.17s  0.58% 70.72%      5.70s 19.30%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.17s  0.58% 71.29%      0.52s  1.76%  runtime.mallocgcSmallNoscan
     0.16s  0.54% 71.83%      5.38s 18.21%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.16s  0.54% 72.38%      0.16s  0.54%  sync/atomic.(*Int32).Add (inline)
     0.15s  0.51% 72.88%      0.30s  1.02%  runtime.greyobject
     0.14s  0.47% 73.36%      0.19s  0.64%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.13s  0.44% 73.80%      1.26s  4.27%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.13s  0.44% 74.24%      0.21s  0.71%  internal/runtime/maps.(*Iter).Next
     0.09s   0.3% 74.54%      0.24s  0.81%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.08s  0.27% 74.81%      3.93s 13.30%  example.com/isureal-company/isureal.loadBig
     0.07s  0.24% 75.05%      0.21s  0.71%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.07s  0.24% 75.29%      2.15s  7.28%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.06s   0.2% 75.49%      4.53s 15.34%  crypto/internal/fips140/sha256.block
     0.06s   0.2% 75.69%      0.52s  1.76%  runtime.convTstring
     0.05s  0.17% 75.86%      0.31s  1.05%  encoding/gob.(*Encoder).encodeGobEncoder
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 29.4% / 100% | 8% | 0.7% | 7 | 109 MB | run queue が積もっている |
| isu2 (2 cpu) | 24.5% / 100% | 7% | 0.8% | 5 | 114 MB | run queue が積もっている |
| isu3 (2 cpu) | 19.6% / 100% | 4% | 0.7% | 6 | 197 MB | run queue が積もっている |
| isu4 (2 cpu) | 56.8% / 100% | 50% | 2.7% | 7 | 102 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 45% / nginx 9% |
| isu2 | isureal 44% / nginx 3% |
| isu3 | isureal 36% / nginx 2% |
| isu4 | isureal 93% / postgres 13% |

## Diff from previous run

スコア: 219729 → 223434 (+3705)

`53bad13..4c078d8`

```
 conf/isu1/home/isucon/env.sh                    |   3 +-
 conf/isu4/etc/systemd/system/isureal-go.service |   5 +-
 conf/isu4/home/isucon/env.sh                    |   1 +
 webapp/go/image_signature.go                    |  61 ++++++++++----
 webapp/go/offload.go                            |  22 ++---
 webapp/go/presign.go                            | 103 ++++++++++++++++++++++++
 6 files changed, 163 insertions(+), 32 deletions(-)
```
