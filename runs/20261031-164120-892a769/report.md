# run 20261031-164120-892a769

## Summary

| | |
|---|---|
| score | 229591 |
| status | pass |
| ref | presign-on-fast-peers |
| sha | 892a76907598 |
| 所要時間 | 306s |
| 前回比 | +9051 (4.1%) vs 20261031-163606-15a1e6c |
| ベスト比 | +5849 (2.6%) vs 20261031-163105-6f2cd6f = 223742 **ベスト更新** |
| note | 次の分の署名を毎分30秒過ぎに EU/US で先に計算（main bdf05cd の上） |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	138	368	688	688
europe	7	21	62	187	541	541
america	6	16	49	142	404	403
```

```
{"id":"29950","status":"done","portalStatus":"finished","passed":true,"score":229591,"raw":229591,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:44:13.430Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:44:16.421Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:44:19.420Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:44:22.434Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T07:44:25.421Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T07:44:28.420Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T07:44:30.576Z level=INFO msg="user moved" user=evanute5396 from=europe to=asia
time=2026-10-31T07:44:31.434Z level=INFO msg="asia 3 日目: 138 人"
time=2026-10-31T07:44:34.420Z level=INFO msg="europe 3 日目: 62 人"
time=2026-10-31T07:44:37.056Z level=INFO msg="user moved" user=ronoa-5771 from=asia to=america
time=2026-10-31T07:44:37.421Z level=INFO msg="america 3 日目: 49 人"
time=2026-10-31T07:44:37.970Z level=INFO msg="user moved" user=evanute5396 from=asia to=europe
time=2026-10-31T07:44:38.530Z level=INFO msg="user moved" user=solchi-6199 from=asia to=europe
time=2026-10-31T07:44:39.595Z level=INFO msg="user moved" user=mari6243 from=europe to=asia
time=2026-10-31T07:44:40.443Z level=INFO msg="user moved" user=kedanmo-6451 from=america to=europe
time=2026-10-31T07:44:40.458Z level=INFO msg="asia 4 日目: 368 人"
time=2026-10-31T07:44:43.421Z level=INFO msg="europe 4 日目: 187 人"
time=2026-10-31T07:44:45.722Z level=INFO msg="user moved" user=maror_6847 from=asia to=america
time=2026-10-31T07:44:46.429Z level=INFO msg="america 4 日目: 142 人"
time=2026-10-31T07:44:46.737Z level=INFO msg="user moved" user=evatsuta8367 from=asia to=america
time=2026-10-31T07:44:47.244Z level=INFO msg="user moved" user=hike-7339 from=asia to=america
time=2026-10-31T07:44:47.537Z level=INFO msg="user moved" user=hosora7424 from=asia to=europe
time=2026-10-31T07:44:47.999Z level=INFO msg="user moved" user=mefu_8924 from=europe to=america
time=2026-10-31T07:44:48.771Z level=INFO msg="user moved" user=kedanmo-6451 from=europe to=america
time=2026-10-31T07:44:49.435Z level=INFO msg="asia 5 日目: 688 人"
time=2026-10-31T07:44:50.173Z level=INFO msg="user moved" user=solchi-6199 from=europe to=america
time=2026-10-31T07:44:52.629Z level=INFO msg="europe 5 日目: 541 人"
time=2026-10-31T07:44:52.636Z level=INFO msg="user moved" user=maulyo_9020 from=america to=asia
time=2026-10-31T07:44:53.113Z level=INFO msg="user moved" user=ronoa-5771 from=america to=europe
time=2026-10-31T07:44:55.431Z level=INFO msg="america 5 日目: 404 人"
time=2026-10-31T07:45:04.456Z level=INFO msg="asia 最終日（グランプリ）: 688 人"
time=2026-10-31T07:45:04.521Z level=INFO msg="america 最終日（グランプリ）: 403 人"
time=2026-10-31T07:45:04.576Z level=INFO msg="europe 最終日（グランプリ）: 541 人"
time=2026-10-31T07:45:13.475Z level=INFO msg="グランプリ: アジア 450 人, ヨーロッパ 410 人, アメリカ 308 人"
{
  "pass": true,
  "score": 229591,
  "messages": [
    "noamunoa-5015: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "miivymax16034: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marra-5507: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hemaho-15693: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sane_11191: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "vinzoe.8684: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": EOF",
    "ethan_miller_4503: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "anna_becker_2684: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
78.4%	1970	1	0	GET	/api/notifications	53417.410	27.115	60.005	60.008	113
5.3%	10734	0	0	GET	^/api/feeds/[^/]+	3584.263	0.334	1.952	5.325	5225
3.5%	2874	0	0	GET	/api/memories	2413.243	0.840	6.915	9.303	10965
2.1%	1939	58	0	POST	/api/posts	1437.018	0.741	7.345	7.865	640
1.9%	2673	0	0	GET	/api/posting-window/current	1297.536	0.485	1.925	2.321	237
1.6%	1582	1	0	GET	^/api/posts/me/current$	1057.748	0.669	1.943	2.366	659
1.5%	1889	0	0	GET	/api/users	1047.971	0.555	1.936	2.169	256
1.0%	213	0	0	GET	/api/me/export	672.680	3.158	13.172	19.318	1935627
0.7%	551	123	0	PUT	/api/me/profile-image	486.351	0.883	13.557	16.566	3
0.7%	829	14	0	GET	^/api/posts/[^/]+/realmojis$	457.775	0.552	3.605	4.609	2687
0.6%	1305	0	0	GET	^/api/relationships/[^/]+	425.149	0.326	1.885	3.821	1671
0.6%	11716	237	0	GET	^/api/images/[^/]+	394.543	0.034	0.704	1.918	174091
0.4%	1846	95	0	GET	^/api/posts/[^/]+$	290.567	0.157	1.206	1.867	600
0.3%	446	3	0	GET	/api/me	235.797	0.529	1.595	1.977	299
0.3%	749	6	0	GET	^/api/posts/[^/]+/comments$	182.522	0.244	1.710	2.138	3103
0.2%	484	0	0	GET	^/api/memories/[^/]+	152.746	0.316	2.049	2.272	629
0.2%	836	16	0	POST	^/api/posts/[^/]+/realmojis$	141.103	0.169	0.942	3.068	477
0.2%	872	0	0	GET	^/api/users/[^/]+$	133.541	0.153	1.201	1.901	203
0.1%	4050	36	0	POST	/api/friend-requests	55.104	0.014	0.146	0.973	114
0.1%	3465	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	51.158	0.015	0.185	1.134	116
0.0%	3850	568	0	POST	/api/auth/login	31.562	0.008	0.141	0.894	212
0.0%	804	0	0	POST	^/api/posts/[^/]+/comments$	23.596	0.029	0.280	0.802	351
0.0%	112	0	0	DELETE	^/api/realmojis/[^/]+$	21.849	0.195	0.493	0.552	0
0.0%	72	0	0	DELETE	^/api/posts/[^/]+$	21.160	0.294	1.114	1.114	0
0.0%	1	0	0	POST	/api/initialize	16.792	16.792	16.792	16.792	90
0.0%	709	1	0	POST	/api/auth/register	12.294	0.017	0.184	0.234	227
0.0%	529	0	0	POST	/api/auth/logout	6.846	0.013	0.142	0.998	0
0.0%	171	0	0	PATCH	^/api/posts/[^/]+$	5.911	0.035	0.721	0.891	660
0.0%	456	0	0	PUT	^/api/follows/[^/]+$	5.696	0.012	0.136	0.253	0
0.0%	335	0	0	PATCH	/api/me	5.537	0.017	0.163	0.263	282
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
85.1%	1372	1	0	GET	/api/notifications	37788.399	27.543	60.101	60.170	104
2.6%	9279	0	0	GET	^/api/feeds/[^/]+	1163.237	0.125	1.045	2.108	5115
2.1%	1378	23	0	POST	/api/posts	948.665	0.688	2.998	63.744	642
2.0%	7294	39	0	GET	^/api/images/[^/]+	891.586	0.122	1.032	2.774	196706
1.5%	3178	34	0	POST	/api/friend-requests	652.430	0.205	0.829	1.646	114
0.8%	2021	0	0	POST	/api/auth/login	333.638	0.165	0.679	1.718	247
0.7%	1746	0	0	GET	/api/memories	309.047	0.177	1.816	2.584	13067
0.6%	1984	0	0	GET	/api/posting-window/current	271.723	0.137	0.552	1.808	238
0.6%	1192	0	0	GET	^/api/posts/me/current$	249.019	0.209	1.562	2.083	652
0.6%	1562	0	0	GET	/api/users	244.841	0.157	0.589	1.644	255
0.4%	520	8	0	POST	^/api/posts/[^/]+/realmojis$	174.651	0.336	1.073	3.775	468
0.4%	1616	40	0	GET	^/api/posts/[^/]+$	157.329	0.097	0.360	1.314	616
0.3%	846	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	147.385	0.174	1.126	1.368	116
0.3%	1114	0	0	GET	^/api/relationships/[^/]+	141.691	0.127	1.264	1.715	578
0.3%	314	0	0	PUT	/api/me/profile-image	111.571	0.355	1.013	1.365	0
0.2%	540	0	0	POST	/api/auth/register	106.787	0.198	0.495	0.584	228
0.2%	544	2	0	POST	^/api/posts/[^/]+/comments$	95.242	0.175	0.515	0.859	354
0.2%	771	0	0	GET	^/api/posts/[^/]+/realmojis$	77.497	0.101	0.441	0.838	3456
0.2%	810	0	0	GET	^/api/users/[^/]+$	77.140	0.095	0.368	1.318	206
0.2%	385	0	0	PUT	^/api/follows/[^/]+$	73.025	0.190	0.496	1.288	0
0.1%	621	2	0	GET	^/api/posts/[^/]+/comments$	60.444	0.097	0.382	1.546	3822
0.1%	337	0	0	POST	/api/auth/logout	55.699	0.165	1.107	1.405	0
0.1%	277	0	0	PATCH	/api/me	49.716	0.179	0.405	0.469	279
0.1%	323	2	0	GET	/api/me	44.555	0.138	0.476	0.552	302
0.1%	130	1	0	PATCH	^/api/posts/[^/]+$	41.617	0.320	1.413	1.569	646
0.1%	101	0	0	DELETE	^/api/realmojis/[^/]+$	31.230	0.309	0.679	0.780	0
0.1%	291	0	0	GET	^/api/memories/[^/]+	28.971	0.100	0.689	2.382	826
0.1%	46	0	0	DELETE	^/api/posts/[^/]+$	25.349	0.551	2.251	2.251	0
0.1%	114	0	0	GET	/api/me/export	23.242	0.204	1.355	4.087	2221179
0.0%	123	0	0	DELETE	^/api/comments/[^/]+$	19.187	0.156	0.416	0.520	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
85.9%	1033	0	0	GET	/api/notifications	27723.657	26.838	60.103	60.520	107
2.4%	6459	30	0	GET	^/api/images/[^/]+	761.096	0.118	1.385	2.846	159976
2.3%	5896	0	0	GET	^/api/feeds/[^/]+	741.825	0.126	1.138	1.897	4277
1.9%	1023	20	0	POST	/api/posts	621.082	0.607	3.067	7.686	644
1.4%	2287	29	0	POST	/api/friend-requests	464.910	0.203	0.971	1.672	114
0.9%	1783	0	0	POST	/api/auth/login	281.868	0.158	0.949	1.857	251
0.6%	1522	0	0	GET	/api/memories	200.372	0.132	1.253	2.640	16140
0.5%	1453	0	0	GET	/api/posting-window/current	159.472	0.110	0.777	1.107	240
0.5%	1128	0	0	GET	/api/users	154.904	0.137	0.835	1.510	252
0.5%	923	0	0	GET	^/api/posts/me/current$	145.698	0.158	1.243	1.718	657
0.4%	372	13	0	POST	^/api/posts/[^/]+/realmojis$	137.555	0.370	3.033	3.281	455
0.3%	719	0	0	GET	^/api/relationships/[^/]+	97.035	0.135	1.373	1.537	541
0.3%	460	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	88.025	0.191	1.110	1.276	116
0.3%	864	29	0	GET	^/api/posts/[^/]+$	87.119	0.101	0.656	1.057	609
0.2%	399	0	0	POST	/api/auth/register	74.075	0.186	0.636	0.687	230
0.2%	293	2	0	POST	^/api/posts/[^/]+/comments$	58.874	0.201	1.004	1.113	343
0.2%	220	0	0	PUT	/api/me/profile-image	56.980	0.259	0.715	3.009	0
0.2%	280	0	0	PUT	^/api/follows/[^/]+$	51.066	0.182	0.668	0.879	0
0.1%	298	0	0	POST	/api/auth/logout	46.992	0.158	0.895	1.821	0
0.1%	454	0	0	GET	^/api/users/[^/]+$	46.293	0.102	0.686	1.219	201
0.1%	402	1	0	GET	^/api/posts/[^/]+/realmojis$	40.631	0.101	0.554	0.903	3533
0.1%	50	0	0	DELETE	^/api/posts/[^/]+$	38.609	0.772	1.873	1.873	0
0.1%	94	0	0	PATCH	^/api/posts/[^/]+$	37.576	0.400	1.373	1.373	652
0.1%	206	0	0	PATCH	/api/me	34.962	0.170	0.555	0.694	276
0.1%	286	0	0	GET	^/api/posts/[^/]+/comments$	27.313	0.095	0.448	0.847	4401
0.1%	94	0	0	GET	/api/me/export	24.454	0.260	5.130	5.130	2941004
0.1%	228	1	0	GET	/api/me	22.919	0.101	0.518	0.540	302
0.1%	277	0	0	GET	^/api/memories/[^/]+	22.511	0.081	0.364	0.927	1076
0.0%	47	0	0	DELETE	^/api/realmojis/[^/]+$	14.643	0.312	0.966	0.966	0
0.0%	49	0	0	DELETE	^/api/comments/[^/]+$	7.605	0.155	0.584	0.584	0
…（上位 30 / 全 36 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 11651 クエリ / 48 種類 / 合計 131.782s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	44.0%	221	57.968	262.30	2071.22	1257.2	127.5	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	19.8%	4	26.080	6520.00	9988.74	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.5%	4	19.079	4769.70	7296.20	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.4%	7025	11.074	1.58	10.72	9.9	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.9%	4	9.078	2269.49	3466.49	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	4.1%	5	5.396	1079.10	3034.06	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	1	0.558	557.76	557.76	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
8	0.4%	1	0.521	521.02	521.02	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
9	0.3%	5	0.439	87.73	163.10	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
10	0.3%	5	0.426	85.25	156.88	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
11	0.2%	1	0.224	224.22	224.22	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.165	165.43	165.43	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.105	104.50	104.50	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	4187	0.101	0.02	2.35	0.9	29.3	SELECT ip, banned_until FROM ip_bans WHERE banned_until > clock_timestamp()
15	0.1%	1	0.098	98.26	98.26	70247.0	10020.0	INSERT INTO follows (id, follower_id, followee_id, created_at, updated_at) SELECT id, follower_id, followee_id, created_at, updated_at FROM initial_follows

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
Build ID: 3c7cf242f088bea4bac57f665442655ec7f6e366
Type: cpu
Time: 2026-10-31 16:43:51 JST
Duration: 60.02s, Total samples = 30.03s (50.04%)
Showing nodes accounting for 22.88s, 76.19% of 30.03s total
Dropped 891 nodes (cum <= 0.15s)
Showing top 40 nodes out of 240
      flat  flat%   sum%        cum   cum%
     4.97s 16.55% 16.55%      4.97s 16.55%  crypto/internal/fips140/sha256.blockAVX2
     3.05s 10.16% 26.71%      3.05s 10.16%  internal/runtime/syscall/linux.Syscall6
     1.75s  5.83% 32.53%      5.61s 18.68%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.42s  4.73% 37.26%      1.42s  4.73%  runtime.memmove
     1.31s  4.36% 41.63%      1.31s  4.36%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.04s  3.46% 45.09%      1.04s  3.46%  aeshashbody
     0.89s  2.96% 48.05%      1.13s  3.76%  runtime.findObject
     0.74s  2.46% 50.52%      2.95s  9.82%  runtime.scanObject
     0.71s  2.36% 52.88%      0.71s  2.36%  runtime.memclrNoHeapPointers
     0.60s  2.00% 54.88%      1.94s  6.46%  runtime.mapaccess1_faststr
     0.55s  1.83% 56.71%      0.94s  3.13%  runtime.tryDeferToSpanScan
     0.54s  1.80% 58.51%      0.81s  2.70%  example.com/isureal-company/isureal/jpegcheck.(*decoder).skipAC
     0.40s  1.33% 59.84%      5.63s 18.75%  crypto/internal/fips140/sha256.(*Digest).Write
     0.39s  1.30% 61.14%      5.79s 19.28%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.35s  1.17% 62.30%      0.56s  1.86%  example.com/isureal-company/isureal.postLess
     0.33s  1.10% 63.40%      0.45s  1.50%  database/sql.convertAssignRows
     0.32s  1.07% 64.47%      0.32s  1.07%  runtime.futex
     0.29s  0.97% 65.43%      0.29s  0.97%  runtime.spanClass.sizeclass (inline)
     0.28s  0.93% 66.37%      0.28s  0.93%  example.com/isureal-company/isureal/jpegcheck.(*decoder).prefetch
     0.26s  0.87% 67.23%      0.26s  0.87%  runtime.nextFreeFast (inline)
     0.25s  0.83% 68.07%      0.25s  0.83%  memeqbody
     0.24s   0.8% 68.86%      1.17s  3.90%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.19s  0.63% 69.50%      1.32s  4.40%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.19s  0.63% 70.13%      1.81s  6.03%  runtime.mallocgc
     0.17s  0.57% 70.70%      0.86s  2.86%  runtime.mallocgcSmallScanNoHeader
     0.16s  0.53% 71.23%      6.29s 20.95%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.16s  0.53% 71.76%      0.19s  0.63%  runtime.(*mspan).writeHeapBitsSmall
     0.16s  0.53% 72.29%      1.34s  4.46%  runtime.mapassign_faststr
     0.14s  0.47% 72.76%      0.19s  0.63%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.12s   0.4% 73.16%      5.97s 19.88%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.12s   0.4% 73.56%      0.17s  0.57%  runtime.greyobject
     0.10s  0.33% 73.89%      0.28s  0.93%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.10s  0.33% 74.23%      2.07s  6.89%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.33% 74.56%      0.34s  1.13%  runtime.mapaccess2_faststr
     0.09s   0.3% 74.86%      5.06s 16.85%  crypto/internal/fips140/sha256.block
     0.09s   0.3% 75.16%      0.19s  0.63%  internal/runtime/maps.(*Iter).Next
     0.09s   0.3% 75.46%      0.41s  1.37%  runtime.mallocgcSmallNoscan
     0.08s  0.27% 75.72%      6.70s 22.31%  example.com/isureal-company/isureal.computeImageSignature
     0.07s  0.23% 75.96%      1.26s  4.20%  encoding/gob.(*Encoder).encodeStruct
     0.07s  0.23% 76.19%      0.21s   0.7%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 26.3% / 100% | 6% | 0.6% | 6 | 112 MB | run queue が積もっている |
| isu2 (2 cpu) | 25.8% / 100% | 10% | 1.1% | 4 | 144 MB |  |
| isu3 (2 cpu) | 19.3% / 100% | 5% | 0.9% | 5 | 175 MB | run queue が積もっている |
| isu4 (2 cpu) | 49.8% / 100% | 44% | 1.8% | 5 | 100 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 42% / nginx 7% |
| isu2 | isureal 46% / nginx 3% |
| isu3 | isureal 35% / nginx 2% |
| isu4 | isureal 83% / postgres 12% |

## Diff from previous run

スコア: 220540 → 229591 (+9051)

`15a1e6c..892a769`

```
 conf/isu1/home/isucon/env.sh |  1 +
 conf/isu4/home/isucon/env.sh |  1 +
 webapp/go/offload.go         | 62 ++++++++++++++++++++++++++++++++++++++++++++
 3 files changed, 64 insertions(+)
```
