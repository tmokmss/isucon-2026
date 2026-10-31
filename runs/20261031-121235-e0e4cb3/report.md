# run 20261031-121235-e0e4cb3

## Summary

| | |
|---|---|
| score | 64700 |
| status | pass |
| ref | origin/main |
| sha | e0e4cb333b1a |
| 所要時間 | 207s |
| 前回比 | +64700 (—) vs 20261031-120912-249be25 |
| ベスト比 | +4499 (7.5%) vs 20261031-115040-1c73ee4 = 60201 **ベスト更新** |
| note | main: max_connections を 300 に |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	45	125	354	428	427
europe	7	17	45	113	113	112
america	6	14	34	64	64	64
```

```
{"id":"7221","status":"done","portalStatus":"finished","passed":true,"score":64700,"raw":64700,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T03:14:37.354Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:14:40.321Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:14:43.318Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:14:46.342Z level=INFO msg="asia 2 日目: 45 人"
time=2026-10-31T03:14:49.319Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T03:14:49.372Z level=INFO msg="user moved" user=muse.5294 from=asia to=europe
time=2026-10-31T03:14:52.317Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T03:14:55.341Z level=INFO msg="asia 3 日目: 125 人"
time=2026-10-31T03:14:57.068Z level=INFO msg="user moved" user=waku_5521 from=asia to=america
time=2026-10-31T03:14:58.324Z level=INFO msg="europe 3 日目: 45 人"
time=2026-10-31T03:14:59.680Z level=INFO msg="user moved" user=waku_5521 from=america to=europe
time=2026-10-31T03:15:01.320Z level=INFO msg="america 3 日目: 34 人"
time=2026-10-31T03:15:01.537Z level=INFO msg="user moved" user=muse.5294 from=europe to=america
time=2026-10-31T03:15:04.361Z level=INFO msg="asia 4 日目: 354 人"
time=2026-10-31T03:15:07.373Z level=INFO msg="europe 4 日目: 113 人"
time=2026-10-31T03:15:10.340Z level=INFO msg="america 4 日目: 64 人"
time=2026-10-31T03:15:13.414Z level=INFO msg="asia 5 日目: 428 人"
time=2026-10-31T03:15:16.786Z level=INFO msg="europe 5 日目: 113 人"
time=2026-10-31T03:15:19.369Z level=INFO msg="america 5 日目: 64 人"
time=2026-10-31T03:15:28.437Z level=INFO msg="europe 最終日（グランプリ）: 112 人"
time=2026-10-31T03:15:28.566Z level=INFO msg="asia 最終日（グランプリ）: 427 人"
time=2026-10-31T03:15:28.617Z level=INFO msg="america 最終日（グランプリ）: 64 人"
time=2026-10-31T03:15:37.330Z level=INFO msg="グランプリ: アジア 27 人, ヨーロッパ 1 人, アメリカ 0 人"
{
  "pass": true,
  "score": 64700,
  "messages": [
    "noako.5068: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "yaleo.8384: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "keku.5496: 友達申請を承認できません: POST /api/friend-requests/10281b7a-ec61-4669-bfe4-b39b485cc457/accept: 応答がタイムアウトしました",
    "rumuis.5872: 友達申請を承認できません: POST /api/friend-requests/eacf3d5d-0491-40de-b5d4-553a2989544d/accept: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.7%	848	1	0	GET	/api/notifications	7657.841	9.030	15.178	15.244	175
17.1%	1121	27	0	POST	/api/posts	3884.173	3.465	11.788	12.531	639
13.3%	2343	46	0	POST	/api/friend-requests	3020.908	1.289	12.239	17.012	113
9.0%	9864	567	0	GET	^/api/images/[^/]+	2043.160	0.207	3.720	12.062	93972
6.0%	4160	995	0	POST	/api/auth/login	1351.781	0.325	4.322	14.560	193
4.3%	982	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	981.361	0.999	12.401	21.034	116
3.6%	305	0	0	PUT	^/api/follows/[^/]+$	823.364	2.700	14.335	16.949	0
2.5%	317	74	0	POST	^/api/posts/[^/]+/comments$	562.183	1.773	15.031	16.710	272
2.1%	287	51	0	POST	^/api/posts/[^/]+/realmojis$	486.003	1.693	15.178	15.999	398
1.3%	2470	0	0	GET	/api/memories	305.771	0.124	0.820	4.437	12691
1.3%	1559	0	0	GET	/api/posting-window/current	302.587	0.194	2.955	10.375	237
1.3%	446	1	0	POST	/api/auth/register	288.227	0.646	3.229	8.539	227
0.9%	366	119	0	PUT	/api/me/profile-image	196.603	0.537	5.853	7.639	8
0.8%	235	0	0	GET	/api/me/export	175.780	0.748	4.600	10.755	1855337
0.6%	5458	0	0	GET	^/api/feeds/[^/]+	133.726	0.025	0.194	3.004	1991
0.5%	568	0	0	POST	/api/auth/logout	110.568	0.195	2.574	12.293	0
0.3%	204	0	0	PATCH	/api/me	59.861	0.293	2.380	2.666	279
0.2%	94	1	0	PATCH	^/api/posts/[^/]+$	48.940	0.521	5.331	5.331	635
0.2%	807	13	0	GET	^/api/posts/me/current$	40.038	0.050	0.853	1.803	644
0.2%	588	0	0	GET	^/api/memories/[^/]+	37.148	0.063	0.478	1.822	664
0.1%	1121	0	0	GET	/api/users	32.077	0.029	0.186	1.453	244
0.1%	45	0	0	DELETE	^/api/posts/[^/]+$	31.649	0.703	5.196	5.196	0
0.1%	1	0	0	POST	/api/initialize	20.692	20.692	20.692	20.692	90
0.1%	842	0	0	GET	^/api/relationships/[^/]+	16.504	0.020	0.122	1.966	1102
0.1%	789	11	0	GET	^/api/posts/[^/]+$	16.477	0.021	0.116	0.655	622
0.0%	33	0	0	DELETE	^/api/realmojis/[^/]+$	10.294	0.312	1.064	1.064	0
0.0%	413	0	0	GET	^/api/users/[^/]+$	9.039	0.022	0.145	0.908	208
0.0%	7	0	0	DELETE	^/api/follows/[^/]+$	8.870	1.267	5.371	5.371	0
0.0%	354	0	0	GET	^/api/posts/[^/]+/realmojis$	6.869	0.019	0.151	0.302	603
0.0%	48	0	0	DELETE	^/api/comments/[^/]+$	6.507	0.136	0.887	0.887	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.3%	246	0	0	GET	/api/notifications	1950.600	7.929	11.753	11.753	177
18.4%	312	5	0	POST	/api/posts	1113.567	3.569	11.887	14.018	644
12.9%	597	10	0	POST	/api/friend-requests	776.813	1.301	12.013	21.378	114
6.1%	579	0	0	POST	/api/auth/login	369.523	0.638	10.045	11.350	250
4.6%	1886	4	0	GET	^/api/images/[^/]+	280.118	0.149	1.004	1.613	134078
4.4%	71	0	0	PUT	^/api/follows/[^/]+$	265.670	3.742	15.396	15.396	0
3.4%	194	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	203.882	1.051	11.758	15.028	116
2.9%	1212	0	0	GET	^/api/feeds/[^/]+	177.355	0.146	0.411	1.224	2175
2.2%	432	0	0	GET	/api/posting-window/current	131.053	0.303	2.694	3.193	239
1.9%	97	16	0	POST	^/api/posts/[^/]+/comments$	115.064	1.186	14.701	14.701	255
1.8%	511	0	0	GET	/api/memories	110.213	0.216	0.814	2.253	13019
1.4%	63	9	0	POST	^/api/posts/[^/]+/realmojis$	81.782	1.298	6.901	6.901	394
1.1%	96	0	0	POST	/api/auth/logout	68.396	0.712	11.467	11.467	0
1.1%	36	0	0	GET	/api/me/export	64.250	1.785	8.231	8.231	1987242
0.8%	65	0	0	PUT	/api/me/profile-image	46.365	0.713	5.221	5.221	0
0.8%	297	0	0	GET	/api/users	46.252	0.156	0.436	1.224	242
0.6%	222	13	0	GET	^/api/posts/me/current$	35.971	0.162	0.266	3.603	615
0.6%	110	0	0	POST	/api/auth/register	33.390	0.304	0.912	0.934	228
0.5%	207	0	0	GET	^/api/relationships/[^/]+	29.326	0.142	0.371	0.500	1037
0.5%	201	16	0	GET	^/api/posts/[^/]+$	27.773	0.138	0.310	0.521	574
0.3%	3	0	0	DELETE	^/api/follows/[^/]+$	17.743	5.914	9.602	9.602	0
0.3%	95	0	0	GET	^/api/memories/[^/]+	17.180	0.181	0.937	0.937	801
0.3%	53	0	0	PATCH	/api/me	16.007	0.302	1.426	1.426	279
0.2%	96	0	0	GET	^/api/users/[^/]+$	13.564	0.141	0.528	0.528	196
0.2%	18	3	0	PATCH	^/api/posts/[^/]+$	11.934	0.663	6.807	6.807	575
0.2%	70	2	0	GET	/api/me	10.340	0.148	0.353	0.353	293
0.2%	74	1	0	GET	^/api/posts/[^/]+/realmojis$	10.210	0.138	0.677	0.677	504
0.2%	78	0	0	GET	^/api/posts/[^/]+/comments$	10.072	0.129	0.281	0.281	522
0.1%	18	0	0	DELETE	^/api/comments/[^/]+$	3.663	0.204	0.885	0.885	0
0.0%	5	0	0	DELETE	^/api/realmojis/[^/]+$	2.089	0.418	0.945	0.945	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.8%	151	0	0	GET	/api/notifications	1170.301	7.750	9.245	9.245	180
16.5%	181	4	0	POST	/api/posts	572.163	3.161	11.156	11.795	638
11.1%	314	6	0	POST	/api/friend-requests	385.104	1.226	13.180	15.034	113
5.5%	276	0	0	POST	/api/auth/login	189.743	0.687	9.665	11.281	250
5.5%	50	0	0	PUT	^/api/follows/[^/]+$	189.687	3.794	15.461	15.461	0
4.9%	121	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	169.141	1.398	13.737	17.196	116
4.7%	991	12	0	GET	^/api/images/[^/]+	162.768	0.164	1.094	1.484	156435
2.9%	685	0	0	GET	^/api/feeds/[^/]+	98.909	0.144	0.432	1.224	2187
1.9%	40	13	0	POST	^/api/posts/[^/]+/comments$	65.524	1.638	9.985	9.985	215
1.9%	238	0	0	GET	/api/posting-window/current	64.522	0.271	1.920	3.189	240
1.8%	294	0	0	GET	/api/memories	62.616	0.213	1.234	2.200	13644
1.2%	23	5	0	POST	^/api/posts/[^/]+/realmojis$	43.003	1.870	13.352	13.352	381
1.2%	21	0	0	GET	/api/me/export	42.151	2.007	6.407	6.407	3197278
0.7%	33	0	0	PUT	/api/me/profile-image	24.206	0.734	4.365	4.365	0
0.7%	35	0	0	PATCH	/api/me	24.184	0.691	4.129	4.129	275
0.7%	49	0	0	POST	/api/auth/logout	24.120	0.492	10.076	10.076	0
0.7%	13	0	0	DELETE	^/api/posts/[^/]+$	24.010	1.847	5.452	5.452	0
0.6%	62	0	0	POST	/api/auth/register	22.049	0.356	4.105	4.105	229
0.6%	138	3	0	GET	^/api/posts/me/current$	19.723	0.143	0.340	0.470	635
0.6%	132	0	0	GET	/api/users	19.600	0.148	0.334	0.450	235
0.5%	125	7	0	GET	^/api/posts/[^/]+$	17.668	0.141	0.481	0.513	578
0.5%	117	0	0	GET	^/api/relationships/[^/]+	16.259	0.139	0.428	0.465	1127
0.4%	3	0	0	DELETE	^/api/realmojis/[^/]+$	12.619	4.206	12.178	12.178	0
0.2%	52	0	0	GET	^/api/users/[^/]+$	7.962	0.153	0.481	0.481	194
0.2%	53	0	0	GET	^/api/posts/[^/]+/realmojis$	7.172	0.135	0.530	0.530	625
0.2%	35	0	0	GET	^/api/memories/[^/]+	6.317	0.180	0.319	0.319	1382
0.2%	36	0	0	GET	/api/me	5.650	0.157	0.476	0.476	298
0.1%	39	0	0	GET	^/api/posts/[^/]+/comments$	5.011	0.128	0.278	0.278	939
0.1%	9	0	0	PATCH	^/api/posts/[^/]+$	4.188	0.465	1.176	1.176	662
0.1%	9	0	0	DELETE	^/api/comments/[^/]+$	2.622	0.291	0.821	0.821	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 48019 クエリ / 82 種類 / 合計 306.957s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	49.5%	2234	152.016	68.05	511.05	658.1	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	31.3%	292	96.131	329.22	1176.75	7011.0	126.0	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	6.2%	2400	19.168	7.99	328.24	120.9	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	5.4%	7384	16.656	2.26	27.38	10.0	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	2.7%	2	8.220	4109.77	8072.07	227710.0	2679.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
6	1.2%	1579	3.775	2.39	31.96	26.3	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
7	0.7%	1	2.167	2167.20	2167.20	29768.0	1199749.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.5%	1	1.645	1645.00	1645.00	20832.0	626709.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.4%	308	1.190	3.86	20.48	30.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.3%	2	0.980	489.86	971.39	16787.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
11	0.2%	1	0.737	737.18	737.18	11528.0	529820.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
12	0.2%	351	0.710	2.02	11.93	23.5	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
13	0.2%	3172	0.588	0.19	11.55	20.1	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
14	0.2%	4020	0.514	0.13	19.57	11.3	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
15	0.1%	1283	0.299	0.23	11.50	21.0	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)

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
Build ID: ef3ae2231ac32d0839fdf7028b105e6fa0a46c70
Type: cpu
Time: 2026-10-31 12:14:11 JST
Duration: 60.11s, Total samples = 38.29s (63.70%)
Showing nodes accounting for 28.51s, 74.46% of 38.29s total
Dropped 896 nodes (cum <= 0.19s)
Showing top 40 nodes out of 286
      flat  flat%   sum%        cum   cum%
     2.28s  5.95%  5.95%      2.28s  5.95%  internal/runtime/syscall/linux.Syscall6
     2.22s  5.80% 11.75%      2.22s  5.80%  crypto/internal/fips140/sha256.blockAVX2
     1.84s  4.81% 16.56%      1.84s  4.81%  runtime.memclrNoHeapPointers
     1.81s  4.73% 21.28%      6.32s 16.51%  runtime.scanObject
     1.79s  4.67% 25.96%      1.79s  4.67%  runtime.memmove
     1.71s  4.47% 30.43%      5.78s 15.10%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.58s  4.13% 34.55%      2.08s  5.43%  runtime.findObject
     1.45s  3.79% 38.34%      2.22s  5.80%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.44s  3.76% 42.10%      1.44s  3.76%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.40s  3.66% 45.76%      4.78s 12.48%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.16s  3.03% 48.79%      1.16s  3.03%  aeshashbody
     1.12s  2.93% 51.71%      1.16s  3.03%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     1.12s  2.93% 54.64%      1.94s  5.07%  runtime.tryDeferToSpanScan
     0.71s  1.85% 56.49%      0.71s  1.85%  example.com/isureal-company/isureal.(*UserRow).active
     0.70s  1.83% 58.32%      0.70s  1.83%  memeqbody
     0.58s  1.51% 59.83%      0.81s  2.12%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.57s  1.49% 61.32%      0.57s  1.49%  runtime.spanClass.sizeclass (inline)
     0.39s  1.02% 62.34%      0.61s  1.59%  example.com/isureal-company/isureal.postLess
     0.37s  0.97% 63.31%      1.77s  4.62%  runtime.mapaccess1_faststr
     0.31s  0.81% 64.12%      0.31s  0.81%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.31s  0.81% 64.93%      0.55s  1.44%  database/sql.convertAssignRows
     0.25s  0.65% 65.58%      0.25s  0.65%  indexbody
     0.25s  0.65% 66.23%      0.58s  1.51%  runtime.greyobject
     0.25s  0.65% 66.88%      0.25s  0.65%  sync/atomic.(*Int32).Add
     0.24s  0.63% 67.51%      0.24s  0.63%  runtime.futex
     0.24s  0.63% 68.14%      0.24s  0.63%  runtime.nextFreeFast (inline)
     0.22s  0.57% 68.71%      0.23s   0.6%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.21s  0.55% 69.26%      0.21s  0.55%  runtime.(*mspan).divideByElemSize (inline)
     0.20s  0.52% 69.78%      0.24s  0.63%  runtime.(*mspan).markBitsForIndex (inline)
     0.20s  0.52% 70.31%      1.82s  4.75%  runtime.mapassign_faststr
     0.20s  0.52% 70.83%      0.23s   0.6%  runtime.spanOf (inline)
     0.18s  0.47% 71.30%      0.22s  0.57%  runtime.typePointers.next
     0.17s  0.44% 71.74%      0.33s  0.86%  compress/flate.(*compressor).deflate
     0.17s  0.44% 72.19%      0.23s   0.6%  runtime.(*mspan).writeHeapBitsSmall
     0.16s  0.42% 72.60%      2.96s  7.73%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.16s  0.42% 73.02%      0.78s  2.04%  runtime.mallocgcSmallScanNoHeader
     0.15s  0.39% 73.41%      2.50s  6.53%  crypto/internal/fips140/sha256.(*Digest).Write
     0.14s  0.37% 73.78%      2.55s  6.66%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.14s  0.37% 74.14%      0.35s  0.91%  runtime.scanObjectsSmall
     0.12s  0.31% 74.46%      2.68s  7.00%  crypto/internal/fips140/sha256.(*Digest).Sum
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 48.1% / 100% | 33% | 0.1% | 6 | 198 MB | CPU が飽和 |
| isu2 (2 cpu) | 10.8% / 30% | 0% | 0.6% | 4 | 2783 MB |  |
| isu3 (2 cpu) | 7.6% / 28% | 0% | 0.6% | 4 | 2902 MB |  |
| isu4 (2 cpu) | 47.6% / 100% | 21% | 17.0% | 25 | 107 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 79% / nginx 11% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 19% / nginx 1% |
| isu3 | isureal 13% / nginx 1% |
| isu4 | postgres 46% / isureal 13% |

## Diff from previous run

スコア: 0 → 64700 (+64700)

`249be25..e0e4cb3`

```
 conf/isu4/etc/postgresql/16/main/conf.d/zz-tuning.conf | 2 ++
 1 file changed, 2 insertions(+)
```
