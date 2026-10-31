# run 20261031-130155-1deeabf

## Summary

| | |
|---|---|
| score | 84880 |
| status | pass |
| ref | tcp-slow-start-idle |
| sha | 1deeabf08fa9 |
| 所要時間 | 344s |
| 前回比 | -16277 (-16.1%) vs 20261031-125928-1deeabf |
| ベスト比 | -16277 (-16.1%) vs 20261031-125928-1deeabf = 101157 |
| note | tcp_slow_start_after_idle=0 2回目 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	51	150	426	888	885
europe	7	18	50	124	125	125
america	6	15	36	72	74	74
```

```
{"id":"11038","status":"done","portalStatus":"finished","passed":true,"score":84880,"raw":84880,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:06:13.961Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:06:16.923Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:06:19.924Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:06:22.951Z level=INFO msg="asia 2 日目: 51 人"
time=2026-10-31T04:06:25.924Z level=INFO msg="europe 2 日目: 18 人"
time=2026-10-31T04:06:28.922Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T04:06:31.947Z level=INFO msg="asia 3 日目: 150 人"
time=2026-10-31T04:06:32.740Z level=INFO msg="user moved" user=tase-5586 from=asia to=europe
time=2026-10-31T04:06:35.010Z level=INFO msg="europe 3 日目: 50 人"
time=2026-10-31T04:06:37.928Z level=INFO msg="america 3 日目: 36 人"
time=2026-10-31T04:06:40.958Z level=INFO msg="asia 4 日目: 426 人"
time=2026-10-31T04:06:42.614Z level=INFO msg="user moved" user=miamia_6634 from=asia to=europe
time=2026-10-31T04:06:44.066Z level=INFO msg="europe 4 日目: 124 人"
time=2026-10-31T04:06:45.972Z level=INFO msg="user moved" user=yaru6804 from=asia to=america
time=2026-10-31T04:06:47.064Z level=INFO msg="america 4 日目: 72 人"
time=2026-10-31T04:06:49.775Z level=INFO msg="user moved" user=hanu.7186 from=asia to=europe
time=2026-10-31T04:06:50.189Z level=INFO msg="asia 5 日目: 888 人"
time=2026-10-31T04:06:52.419Z level=INFO msg="user moved" user=evaya-7225 from=asia to=america
time=2026-10-31T04:06:53.054Z level=INFO msg="europe 5 日目: 125 人"
time=2026-10-31T04:06:55.382Z level=INFO msg="user moved" user=anshi-8009 from=asia to=america
time=2026-10-31T04:06:56.240Z level=INFO msg="america 5 日目: 74 人"
time=2026-10-31T04:07:05.341Z level=INFO msg="asia 最終日（グランプリ）: 885 人"
time=2026-10-31T04:07:05.586Z level=INFO msg="europe 最終日（グランプリ）: 125 人"
time=2026-10-31T04:07:05.692Z level=INFO msg="america 最終日（グランプリ）: 74 人"
time=2026-10-31T04:07:13.960Z level=INFO msg="グランプリ: アジア 171 人, ヨーロッパ 4 人, アメリカ 4 人"
{
  "pass": true,
  "score": 84880,
  "messages": [
    "shihi-5026: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "daniel_rodriguez_4850: ログインできません: POST /api/auth/login: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
30.6%	1363	1	0	GET	/api/notifications	11757.490	8.626	15.330	15.367	165
17.3%	1882	73	0	POST	/api/posts	6624.311	3.520	11.208	15.407	625
13.5%	4393	76	0	POST	/api/friend-requests	5181.343	1.179	10.689	15.411	114
8.9%	12628	1538	0	GET	^/api/images/[^/]+	3412.244	0.270	6.154	16.640	119834
8.3%	6018	2315	0	POST	/api/auth/login	3198.201	0.531	7.838	12.336	161
4.3%	1273	3	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1665.492	1.308	10.415	15.018	116
4.1%	908	1	0	POST	/api/auth/register	1589.336	1.750	8.835	13.514	227
2.3%	608	123	0	PUT	/api/me/profile-image	880.063	1.447	8.003	12.240	4
1.5%	581	0	0	PUT	^/api/follows/[^/]+$	583.658	1.005	9.667	12.297	0
1.2%	435	0	0	PATCH	/api/me	473.147	1.088	7.981	13.496	275
1.2%	397	15	0	POST	^/api/posts/[^/]+/realmojis$	447.641	1.128	8.882	13.798	459
1.1%	440	25	0	POST	^/api/posts/[^/]+/comments$	433.144	0.984	9.735	14.035	320
1.0%	2946	0	0	GET	/api/memories	388.904	0.132	1.202	1.819	11187
0.9%	170	9	0	PATCH	^/api/posts/[^/]+$	326.381	1.920	10.771	13.179	615
0.7%	8539	0	0	GET	^/api/feeds/[^/]+	269.863	0.032	0.392	12.444	2049
0.6%	574	0	0	POST	/api/auth/logout	217.001	0.378	6.456	9.121	0
0.5%	207	0	0	GET	/api/me/export	186.663	0.902	6.042	8.761	1860634
0.4%	75	0	0	DELETE	^/api/posts/[^/]+$	147.040	1.961	9.575	9.575	0
0.3%	2704	0	0	GET	/api/posting-window/current	125.267	0.046	0.433	6.318	237
0.2%	1407	88	0	GET	^/api/posts/me/current$	89.358	0.064	0.575	1.952	610
0.2%	2162	0	0	GET	/api/users	73.177	0.034	0.422	1.102	237
0.2%	62	0	0	DELETE	^/api/realmojis/[^/]+$	70.732	1.141	12.098	12.098	0
0.1%	68	0	0	DELETE	^/api/comments/[^/]+$	47.846	0.704	12.051	12.051	0
0.1%	1198	0	0	GET	^/api/relationships/[^/]+	36.928	0.031	0.479	0.838	917
0.1%	529	0	0	GET	^/api/memories/[^/]+	31.506	0.060	0.392	1.690	656
0.1%	1273	12	0	GET	^/api/posts/[^/]+$	28.203	0.022	0.327	0.601	613
0.1%	24	0	0	PUT	^/api/blocks/[^/]+$	21.095	0.879	5.877	5.877	0
0.0%	1	0	0	POST	/api/initialize	17.619	17.619	17.619	17.619	90
0.0%	493	3	0	GET	/api/me	12.344	0.025	0.186	0.441	301
0.0%	619	0	0	GET	^/api/users/[^/]+$	12.108	0.020	0.147	0.448	195
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.9%	297	0	0	GET	/api/notifications	2545.635	8.571	12.214	12.763	179
20.2%	358	12	0	POST	/api/posts	1396.985	3.902	11.200	15.454	624
10.3%	578	12	0	POST	/api/friend-requests	709.702	1.228	9.985	15.638	113
6.5%	469	0	0	POST	/api/auth/login	449.157	0.958	9.704	13.820	248
4.8%	1631	13	0	GET	^/api/images/[^/]+	333.658	0.205	1.440	3.643	142701
4.4%	186	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	303.462	1.632	10.807	11.226	116
2.8%	1107	0	0	GET	^/api/feeds/[^/]+	191.137	0.173	0.634	1.101	1720
1.7%	492	0	0	GET	/api/posting-window/current	120.696	0.245	0.906	1.125	239
1.7%	457	0	0	GET	/api/memories	115.591	0.253	1.041	1.947	15410
1.5%	106	0	0	PUT	^/api/follows/[^/]+$	105.831	0.998	9.681	9.895	0
1.2%	55	7	0	POST	^/api/posts/[^/]+/realmojis$	85.333	1.552	15.483	15.483	433
0.9%	120	0	0	POST	/api/auth/register	64.683	0.539	2.426	2.776	228
0.8%	279	0	0	GET	/api/users	54.981	0.197	0.915	1.175	237
0.8%	66	0	0	POST	/api/auth/logout	54.397	0.824	13.968	13.968	0
0.7%	260	20	0	GET	^/api/posts/me/current$	51.001	0.196	0.744	0.816	595
0.7%	64	0	0	PUT	/api/me/profile-image	47.468	0.742	4.345	4.345	0
0.5%	14	0	0	DELETE	^/api/posts/[^/]+$	33.986	2.428	7.871	7.871	0
0.5%	171	0	0	GET	^/api/relationships/[^/]+	31.454	0.184	0.746	1.729	919
0.4%	17	0	0	GET	/api/me/export	30.866	1.816	7.588	7.588	1540450
0.4%	22	3	0	PATCH	^/api/posts/[^/]+$	26.287	1.195	11.550	11.550	575
0.4%	55	0	0	PATCH	/api/me	26.272	0.478	1.911	1.911	290
0.3%	141	9	0	GET	^/api/posts/[^/]+$	23.130	0.164	0.635	0.641	568
0.3%	43	3	0	POST	^/api/posts/[^/]+/comments$	22.261	0.518	6.331	6.331	295
0.2%	1	0	0	POST	/api/me/account-deletion	15.921	15.921	15.921	15.921	0
0.2%	68	1	0	GET	/api/me	13.146	0.193	1.020	1.020	297
0.2%	80	0	0	GET	^/api/users/[^/]+$	12.287	0.154	0.590	0.590	182
0.1%	53	0	0	GET	^/api/memories/[^/]+	10.223	0.193	0.679	0.679	772
0.1%	9	0	0	DELETE	^/api/realmojis/[^/]+$	9.188	1.021	3.212	3.212	0
0.1%	49	0	0	GET	^/api/posts/[^/]+/realmojis$	7.680	0.157	0.539	0.539	848
0.1%	50	0	0	GET	^/api/posts/[^/]+/comments$	6.529	0.131	0.334	0.334	748
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.6%	171	0	0	GET	/api/notifications	1329.208	7.773	9.574	9.686	179
18.0%	205	5	0	POST	/api/posts	692.899	3.380	10.829	11.984	622
10.5%	347	3	0	POST	/api/friend-requests	405.180	1.168	9.319	10.722	114
7.1%	244	0	0	POST	/api/auth/login	273.430	1.121	10.294	16.966	248
5.5%	144	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	213.376	1.482	8.684	8.775	116
4.6%	836	15	0	GET	^/api/images/[^/]+	177.250	0.212	1.108	3.401	174403
2.6%	553	0	0	GET	^/api/feeds/[^/]+	101.202	0.183	0.631	1.009	1973
2.1%	45	0	0	PUT	^/api/follows/[^/]+$	81.914	1.820	10.857	10.857	0
1.8%	286	0	0	GET	/api/posting-window/current	67.835	0.237	0.683	1.494	240
1.7%	240	0	0	GET	/api/memories	65.045	0.271	1.490	2.226	13389
1.4%	17	0	0	DELETE	^/api/posts/[^/]+$	52.472	3.087	6.720	6.720	0
1.1%	70	0	0	POST	/api/auth/register	43.215	0.617	8.797	8.797	229
1.1%	16	0	0	GET	/api/me/export	42.726	2.670	6.011	6.011	3295883
1.0%	175	0	0	GET	/api/users	39.050	0.223	0.931	1.571	240
1.0%	39	0	0	POST	/api/auth/logout	38.536	0.988	7.521	7.521	0
0.9%	35	0	0	PUT	/api/me/profile-image	35.166	1.005	5.825	5.825	0
0.9%	155	16	0	GET	^/api/posts/me/current$	33.451	0.216	1.035	1.085	568
0.7%	18	1	0	POST	^/api/posts/[^/]+/realmojis$	27.155	1.509	4.178	4.178	467
0.7%	27	2	0	POST	^/api/posts/[^/]+/comments$	26.361	0.976	7.725	7.725	301
0.5%	93	0	0	GET	^/api/relationships/[^/]+	20.789	0.224	1.662	1.662	1310
0.5%	23	0	0	PATCH	/api/me	20.251	0.880	3.589	3.589	291
0.3%	10	0	0	PATCH	^/api/posts/[^/]+$	10.910	1.091	3.159	3.159	648
0.3%	57	5	0	GET	^/api/posts/[^/]+$	10.188	0.179	0.577	0.577	560
0.2%	30	0	0	GET	^/api/memories/[^/]+	9.260	0.309	1.716	1.716	1130
0.2%	3	0	0	DELETE	^/api/comments/[^/]+$	6.891	2.297	6.218	6.218	0
0.1%	38	0	0	GET	/api/me	4.934	0.130	0.429	0.429	298
0.1%	29	0	0	GET	^/api/posts/[^/]+/realmojis$	4.636	0.160	0.356	0.356	642
0.1%	25	0	0	GET	^/api/users/[^/]+$	4.138	0.166	0.395	0.395	183
0.1%	27	0	0	GET	^/api/posts/[^/]+/comments$	3.853	0.143	0.265	0.265	1312
0.0%	1	0	0	DELETE	^/api/friendships/[^/]+$	1.772	1.772	1.772	1.772	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 59277 クエリ / 81 種類 / 合計 416.446s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	63.0%	3390	262.181	77.34	500.73	657.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	21.3%	240	88.611	369.21	1551.80	6531.8	123.9	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.1%	2213	21.412	9.68	559.45	115.9	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.7%	7842	19.386	2.47	41.25	10.3	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.5%	2358	6.447	2.73	50.43	26.2	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.3%	2	5.249	2624.68	4758.89	236215.5	2782.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.5%	1	2.271	2271.42	2271.42	30431.0	1199843.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
8	0.5%	1	1.880	1879.57	1879.57	22627.0	627393.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.4%	447	1.829	4.09	16.96	30.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.3%	5204	1.207	0.23	17.46	20.0	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.2%	480	0.970	2.02	15.88	22.5	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.2%	1	0.788	788.30	788.30	12086.0	530031.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
13	0.1%	4416	0.549	0.12	17.08	11.1	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
14	0.1%	1564	0.542	0.35	39.24	21.6	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	18	0.313	17.36	220.36	281.9	1.0	INSERT INTO posting_windows (id, region, local_date, starts_at, on_time_deadline, created_at, notification_ids, notified_user_ids) VALUES ($1, $2, $3, $4, $5, $…

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
Build ID: 357f5242f3cecf617bc0cfc64f996e4d3bf04765
Type: cpu
Time: 2026-10-31 13:05:51 JST
Duration: 60.04s, Total samples = 41.17s (68.57%)
Showing nodes accounting for 29.85s, 72.50% of 41.17s total
Dropped 961 nodes (cum <= 0.21s)
Showing top 40 nodes out of 324
      flat  flat%   sum%        cum   cum%
     3.67s  8.91%  8.91%      3.67s  8.91%  internal/runtime/syscall/linux.Syscall6
     2.34s  5.68% 14.60%      2.34s  5.68%  runtime.memmove
     2.28s  5.54% 20.14%      2.28s  5.54%  runtime.memclrNoHeapPointers
     2.11s  5.13% 25.26%      7.50s 18.22%  runtime.scanObject
     1.84s  4.47% 29.73%      1.84s  4.47%  crypto/internal/fips140/sha256.blockAVX2
     1.79s  4.35% 34.08%      5.77s 14.02%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.79s  4.35% 38.43%      2.44s  5.93%  runtime.findObject
     1.48s  3.59% 42.02%      2.62s  6.36%  runtime.tryDeferToSpanScan
     1.41s  3.42% 45.45%      1.41s  3.42%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.18s  2.87% 48.31%      1.18s  2.87%  example.com/isureal-company/isureal.(*UserRow).active
     1.17s  2.84% 51.15%      1.17s  2.84%  aeshashbody
     0.78s  1.89% 53.05%      0.78s  1.89%  runtime.spanClass.sizeclass (inline)
     0.77s  1.87% 54.92%      0.77s  1.87%  memeqbody
     0.58s  1.41% 56.33%      1.89s  4.59%  runtime.mapaccess1_faststr
     0.51s  1.24% 57.57%      1.48s  3.59%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.44s  1.07% 58.63%      0.60s  1.46%  example.com/isureal-company/isureal.postLess
     0.38s  0.92% 59.56%      0.38s  0.92%  runtime.nextFreeFast (inline)
     0.36s  0.87% 60.43%      0.63s  1.53%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.33s   0.8% 61.23%      0.34s  0.83%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.31s  0.75% 61.99%      0.31s  0.75%  crypto/internal/fips140/aes/gcm.gcmAesEnc
     0.30s  0.73% 62.72%      0.30s  0.73%  indexbody
     0.30s  0.73% 63.44%      0.30s  0.73%  runtime.futex
     0.29s   0.7% 64.15%      0.75s  1.82%  database/sql.convertAssignRows
     0.28s  0.68% 64.83%      1.14s  2.77%  runtime.mallocgcSmallScanNoHeader
     0.28s  0.68% 65.51%      1.97s  4.79%  runtime.mapassign_faststr
     0.27s  0.66% 66.16%      0.27s  0.66%  runtime.(*mspan).divideByElemSize (inline)
     0.26s  0.63% 66.80%      0.26s  0.63%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.26s  0.63% 67.43%      0.26s  0.63%  sync/atomic.(*Int32).Add
     0.24s  0.58% 68.01%      0.24s  0.58%  runtime.typePointers.nextFast (inline)
     0.23s  0.56% 68.57%      0.28s  0.68%  runtime.typePointers.next
     0.21s  0.51% 69.08%      0.28s  0.68%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.21s  0.51% 69.59%      0.25s  0.61%  runtime.heapArenaOf (inline)
     0.20s  0.49% 70.08%      0.23s  0.56%  runtime.(*mspan).writeHeapBitsSmall
     0.17s  0.41% 70.49%      0.25s  0.61%  runtime.spanOf (inline)
     0.16s  0.39% 70.88%      0.44s  1.07%  runtime.greyobject
     0.16s  0.39% 71.27%      4.86s 11.80%  runtime.mallocgc
     0.14s  0.34% 71.61%      3.06s  7.43%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.13s  0.32% 71.92%      0.42s  1.02%  runtime.scanObjectsSmall
     0.12s  0.29% 72.21%      2.28s  5.54%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.12s  0.29% 72.50%      2.09s  5.08%  crypto/internal/fips140/sha256.(*Digest).Write
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 49.3% / 100% | 33% | 0.2% | 6 | 136 MB | CPU が飽和 |
| isu2 (2 cpu) | 14.6% / 52% | 0% | 0.3% | 3 | 2515 MB |  |
| isu3 (2 cpu) | 14.8% / 60% | 0% | 0.3% | 3 | 2590 MB |  |
| isu4 (2 cpu) | 50.9% / 100% | 23% | 17.0% | 30 | 99 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 75% / nginx 14% / journalctl 2% / systemd-journal 2% |
| isu2 | isureal 27% / nginx 1% |
| isu3 | isureal 28% / nginx 1% |
| isu4 | postgres 58% / isureal 5% |

## Diff from previous run

スコア: 101157 → 84880 (-16277)

コードの差分は無い（同じ sha 1deeabf08fa9）。
スコアの違いはブレか、サーバ側の設定変更による。
