# run 20261031-134616-f256785

## Summary

| | |
|---|---|
| score | 115400 |
| status | pass |
| ref | nginx-buffers-tcp |
| sha | f256785754ab |
| 所要時間 | 249s |
| 前回比 | +31656 (37.8%) vs 20261031-132303-5a138c8 |
| ベスト比 | +31656 (37.8%) vs 20261031-132303-5a138c8 = 83744 **ベスト更新** |
| note | nginx-no-tempfile + tcp_slow_start_after_idle=0 on fbae0df |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	140	400	868	867
europe	7	20	64	167	171	171
america	7	17	47	113	113	113
```

```
{"id":"14337","status":"done","portalStatus":"finished","passed":true,"score":115400,"raw":115400,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:48:52.879Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:48:55.429Z level=INFO msg="user moved" user=shileowa_5023 from=asia to=america
time=2026-10-31T04:48:55.850Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:48:58.849Z level=INFO msg="america 1 日目: 7 人"
time=2026-10-31T04:49:01.871Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T04:49:04.849Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T04:49:07.720Z level=INFO msg="user moved" user=evayamu5369 from=europe to=asia
time=2026-10-31T04:49:07.860Z level=INFO msg="america 2 日目: 17 人"
time=2026-10-31T04:49:10.869Z level=INFO msg="asia 3 日目: 140 人"
time=2026-10-31T04:49:12.752Z level=INFO msg="user moved" user=shileowa_5023 from=america to=asia
time=2026-10-31T04:49:12.826Z level=INFO msg="user moved" user=yabense.5756 from=asia to=europe
time=2026-10-31T04:49:12.994Z level=INFO msg="user moved" user=evadan-5798 from=asia to=europe
time=2026-10-31T04:49:13.094Z level=INFO msg="user moved" user=evayamu5369 from=asia to=europe
time=2026-10-31T04:49:13.850Z level=INFO msg="europe 3 日目: 64 人"
time=2026-10-31T04:49:16.856Z level=INFO msg="america 3 日目: 47 人"
time=2026-10-31T04:49:19.876Z level=INFO msg="asia 4 日目: 400 人"
time=2026-10-31T04:49:22.858Z level=INFO msg="europe 4 日目: 167 人"
time=2026-10-31T04:49:23.395Z level=INFO msg="user moved" user=miarara6596 from=america to=asia
time=2026-10-31T04:49:23.442Z level=INFO msg="user moved" user=shileowa_5023 from=asia to=america
time=2026-10-31T04:49:23.493Z level=INFO msg="user moved" user=vinkovin_6860 from=asia to=europe
time=2026-10-31T04:49:24.844Z level=INFO msg="user moved" user=miarara6596 from=asia to=europe
time=2026-10-31T04:49:26.019Z level=INFO msg="america 4 日目: 113 人"
time=2026-10-31T04:49:27.585Z level=INFO msg="user moved" user=evadan-5798 from=europe to=asia
time=2026-10-31T04:49:28.271Z level=INFO msg="user moved" user=muri-6958 from=asia to=europe
time=2026-10-31T04:49:28.422Z level=INFO msg="user moved" user=sumardan8082 from=asia to=europe
time=2026-10-31T04:49:28.641Z level=INFO msg="user moved" user=danelri7581 from=asia to=europe
time=2026-10-31T04:49:29.114Z level=INFO msg="asia 5 日目: 868 人"
time=2026-10-31T04:49:31.909Z level=INFO msg="europe 5 日目: 171 人"
time=2026-10-31T04:49:35.014Z level=INFO msg="america 5 日目: 113 人"
time=2026-10-31T04:49:44.019Z level=INFO msg="america 最終日（グランプリ）: 113 人"
time=2026-10-31T04:49:44.062Z level=INFO msg="europe 最終日（グランプリ）: 171 人"
time=2026-10-31T04:49:44.234Z level=INFO msg="asia 最終日（グランプリ）: 867 人"
time=2026-10-31T04:49:52.882Z level=INFO msg="グランプリ: アジア 343 人, ヨーロッパ 106 人, アメリカ 65 人"
{
  "pass": true,
  "score": 115400,
  "messages": [
    "hiteke_6456: ログインし直せません: POST /api/auth/login: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.1%	1345	0	0	GET	/api/notifications	11525.189	8.569	15.180	15.197	164
18.4%	1975	60	0	POST	/api/posts	6612.196	3.348	10.004	14.258	632
11.6%	4592	75	0	POST	/api/friend-requests	4149.901	0.904	7.002	12.292	114
8.3%	12646	597	0	GET	^/api/images/[^/]+	2988.219	0.236	4.529	10.657	141870
6.9%	6371	2552	0	POST	/api/auth/login	2474.103	0.388	6.092	12.162	158
4.6%	1644	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1669.423	1.015	6.532	13.215	116
3.9%	610	115	0	PUT	/api/me/profile-image	1414.554	2.319	9.610	12.437	4
2.5%	890	1	0	POST	/api/auth/register	892.027	1.002	7.562	10.622	227
2.4%	556	31	0	POST	^/api/posts/[^/]+/realmojis$	874.550	1.573	9.465	11.214	453
1.5%	504	14	0	POST	^/api/posts/[^/]+/comments$	531.487	1.055	6.556	7.871	332
1.4%	524	0	0	PUT	^/api/follows/[^/]+$	490.601	0.936	7.108	12.327	0
1.1%	433	0	0	PATCH	/api/me	411.733	0.951	6.712	11.288	282
0.9%	3156	0	0	GET	/api/memories	337.840	0.107	0.735	1.224	10985
0.8%	192	8	0	PATCH	^/api/posts/[^/]+$	275.477	1.435	7.485	7.566	624
0.7%	232	0	0	GET	/api/me/export	247.158	1.065	4.798	13.869	1924575
0.6%	11396	0	0	GET	^/api/feeds/[^/]+	204.682	0.018	0.149	1.294	2195
0.5%	595	0	0	POST	/api/auth/logout	196.199	0.330	4.833	7.901	0
0.4%	83	0	0	DELETE	^/api/posts/[^/]+$	126.261	1.521	7.319	7.319	0
0.2%	109	0	0	DELETE	^/api/comments/[^/]+$	76.009	0.697	5.483	6.605	0
0.2%	2858	0	0	GET	/api/posting-window/current	70.722	0.025	0.191	0.818	236
0.2%	79	0	0	DELETE	^/api/realmojis/[^/]+$	58.951	0.746	5.815	5.815	0
0.1%	1559	67	0	GET	^/api/posts/me/current$	46.345	0.030	0.213	3.408	623
0.1%	2130	0	0	GET	/api/users	37.823	0.018	0.142	1.400	242
0.1%	558	0	0	GET	^/api/memories/[^/]+	30.511	0.055	0.597	0.747	674
0.1%	30	0	0	PUT	^/api/blocks/[^/]+$	27.533	0.918	3.530	3.530	0
0.1%	1614	17	0	GET	^/api/posts/[^/]+$	26.628	0.016	0.129	1.258	641
0.1%	1414	0	0	GET	^/api/relationships/[^/]+	24.402	0.017	0.179	0.605	917
0.1%	1	0	0	POST	/api/initialize	19.137	19.137	19.137	19.137	90
0.0%	910	0	0	GET	^/api/users/[^/]+$	12.370	0.014	0.122	0.248	221
0.0%	787	1	0	GET	^/api/posts/[^/]+/realmojis$	11.294	0.014	0.122	1.371	1144
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.5%	423	0	0	GET	/api/notifications	3520.126	8.322	11.687	12.926	178
14.4%	541	5	0	POST	/api/posts	1283.852	2.373	7.828	8.783	649
9.3%	936	19	0	POST	/api/friend-requests	833.355	0.890	6.983	8.293	113
5.4%	575	0	0	POST	/api/auth/login	484.759	0.843	7.297	16.612	247
5.1%	2246	16	0	GET	^/api/images/[^/]+	451.269	0.201	1.183	2.645	210502
4.8%	2178	0	0	GET	^/api/feeds/[^/]+	426.097	0.196	0.683	1.317	2329
3.7%	273	2	0	POST	^/api/friend-requests/[^/]+/[^/]+$	331.392	1.214	8.320	9.456	115
2.2%	737	0	0	GET	/api/posting-window/current	191.980	0.260	0.709	1.330	239
1.8%	161	9	0	POST	^/api/posts/[^/]+/comments$	164.963	1.025	10.427	12.023	337
1.7%	641	0	0	GET	/api/memories	148.106	0.231	0.763	1.298	11013
1.3%	59	8	0	POST	^/api/posts/[^/]+/realmojis$	116.665	1.977	9.877	9.877	379
1.2%	483	25	0	GET	^/api/posts/me/current$	109.014	0.226	0.724	2.117	622
1.2%	87	0	0	POST	/api/auth/logout	102.751	1.181	7.334	7.334	0
1.0%	448	0	0	GET	/api/users	88.260	0.197	1.071	1.328	241
0.9%	163	0	0	POST	/api/auth/register	81.297	0.499	2.147	3.141	228
0.8%	96	0	0	PUT	/api/me/profile-image	74.720	0.778	10.858	10.858	0
0.8%	39	0	0	GET	/api/me/export	69.961	1.794	5.881	5.881	1927403
0.7%	47	5	0	PATCH	^/api/posts/[^/]+$	62.027	1.320	5.796	5.796	587
0.6%	115	0	0	PUT	^/api/follows/[^/]+$	54.346	0.473	2.902	3.181	0
0.6%	288	15	0	GET	^/api/posts/[^/]+$	53.838	0.187	0.688	1.339	608
0.6%	283	0	0	GET	^/api/relationships/[^/]+	51.389	0.182	0.665	1.330	795
0.5%	24	0	0	DELETE	^/api/posts/[^/]+$	40.995	1.708	6.150	6.150	0
0.4%	78	0	0	PATCH	/api/me	39.997	0.513	4.311	4.311	288
0.3%	130	0	0	GET	^/api/users/[^/]+$	24.676	0.190	0.636	0.672	213
0.2%	116	2	0	GET	^/api/posts/[^/]+/comments$	21.290	0.184	0.699	1.313	1411
0.2%	102	1	0	GET	/api/me	18.196	0.178	0.433	0.575	299
0.2%	115	3	0	GET	^/api/posts/[^/]+/realmojis$	18.161	0.158	0.380	0.508	879
0.2%	67	0	0	GET	^/api/memories/[^/]+	14.905	0.222	0.703	0.703	837
0.2%	12	0	0	DELETE	^/api/realmojis/[^/]+$	14.042	1.170	3.035	3.035	0
0.1%	23	0	0	DELETE	^/api/comments/[^/]+$	13.094	0.569	2.365	2.365	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.8%	262	2	0	GET	/api/notifications	1888.864	7.209	9.177	9.192	175
14.1%	335	4	0	POST	/api/posts	764.785	2.283	7.901	8.423	647
12.6%	602	10	0	POST	/api/friend-requests	682.987	1.135	6.542	9.211	114
6.1%	396	0	0	POST	/api/auth/login	333.995	0.843	8.557	11.871	250
5.9%	1476	13	0	GET	^/api/images/[^/]+	321.747	0.218	1.133	2.056	189835
4.2%	1220	0	0	GET	^/api/feeds/[^/]+	227.421	0.186	0.600	1.071	2593
4.1%	205	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	224.540	1.095	6.860	8.025	116
1.9%	458	0	0	GET	/api/posting-window/current	105.470	0.230	0.631	0.778	239
1.8%	435	0	0	GET	/api/memories	99.803	0.229	1.086	1.553	12198
1.4%	65	0	0	PUT	/api/me/profile-image	77.723	1.196	9.635	9.635	0
1.4%	74	0	0	PUT	^/api/follows/[^/]+$	77.099	1.042	7.714	7.714	0
1.3%	66	0	0	POST	/api/auth/logout	69.556	1.054	7.488	7.488	0
1.2%	26	8	0	POST	^/api/posts/[^/]+/realmojis$	63.337	2.436	7.805	7.805	348
1.1%	300	27	0	GET	^/api/posts/me/current$	58.858	0.196	0.593	0.628	599
1.0%	294	0	0	GET	/api/users	55.960	0.190	0.592	0.958	243
1.0%	25	0	0	GET	/api/me/export	54.689	2.188	7.735	7.735	2511442
0.9%	35	2	0	POST	^/api/posts/[^/]+/comments$	49.609	1.417	4.548	4.548	346
0.7%	112	0	0	POST	/api/auth/register	40.121	0.358	2.206	2.810	229
0.7%	62	0	0	PATCH	/api/me	40.108	0.647	5.224	5.224	289
0.6%	180	6	0	GET	^/api/posts/[^/]+$	32.314	0.180	0.600	0.666	628
0.6%	37	10	0	PATCH	^/api/posts/[^/]+$	32.140	0.869	3.912	3.912	491
0.6%	175	0	0	GET	^/api/relationships/[^/]+	30.160	0.172	0.604	0.611	1199
0.5%	13	0	0	DELETE	^/api/posts/[^/]+$	26.295	2.023	4.406	4.406	0
0.3%	81	0	0	GET	^/api/users/[^/]+$	13.703	0.169	1.007	1.007	220
0.2%	68	0	0	GET	^/api/posts/[^/]+/realmojis$	12.023	0.177	0.614	0.614	650
0.2%	61	0	0	GET	^/api/posts/[^/]+/comments$	11.654	0.191	0.786	0.786	889
0.2%	56	0	0	GET	^/api/memories/[^/]+	11.557	0.206	0.557	0.557	1558
0.2%	69	0	0	GET	/api/me	10.340	0.150	0.574	0.574	298
0.2%	3	0	0	PUT	^/api/blocks/[^/]+$	9.341	3.114	6.483	6.483	0
0.1%	2	0	0	DELETE	^/api/realmojis/[^/]+$	4.847	2.423	3.405	3.405	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 67959 クエリ / 82 種類 / 合計 544.476s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	63.4%	4034	344.999	85.52	869.95	655.1	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	20.5%	296	111.420	376.42	1966.02	1244.3	126.4	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	6.5%	3092	35.152	11.37	605.63	121.5	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.2%	8250	22.661	2.75	52.72	10.8	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.6%	2784	8.580	3.08	53.76	24.8	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.0%	2	5.511	2755.69	5358.39	238736.0	2812.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.5%	594	2.640	4.44	40.05	28.6	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
8	0.4%	1	2.341	2341.39	2341.39	30882.0	1199900.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
9	0.4%	1	1.990	1989.61	1989.61	24075.0	627846.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
10	0.3%	6009	1.510	0.25	67.85	19.6	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.3%	675	1.452	2.15	45.01	20.9	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.2%	1	0.879	879.45	879.45	29546.0	2735.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.2%	1	0.823	822.68	822.68	12577.0	530079.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
14	0.1%	2036	0.717	0.35	44.45	20.2	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	4790	0.603	0.13	13.18	10.9	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: d5773d58212d730e446e9b0c85d135f40643caa5
Type: cpu
Time: 2026-10-31 13:48:28 JST
Duration: 60s, Total samples = 35.53s (59.22%)
Showing nodes accounting for 24.96s, 70.25% of 35.53s total
Dropped 990 nodes (cum <= 0.18s)
Showing top 40 nodes out of 289
      flat  flat%   sum%        cum   cum%
     3.57s 10.05% 10.05%      3.57s 10.05%  internal/runtime/syscall/linux.Syscall6
     2.86s  8.05% 18.10%      2.86s  8.05%  crypto/internal/fips140/sha256.blockAVX2
     1.79s  5.04% 23.14%      1.79s  5.04%  runtime.memmove
     1.68s  4.73% 27.86%      6.01s 16.92%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.63s  4.59% 32.45%      1.63s  4.59%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.44s  4.05% 36.50%      1.73s  4.87%  runtime.findObject
     1.42s  4.00% 40.50%      5.51s 15.51%  runtime.scanObject
     1.40s  3.94% 44.44%      1.40s  3.94%  runtime.memclrNoHeapPointers
     1.27s  3.57% 48.02%      1.27s  3.57%  aeshashbody
     1.23s  3.46% 51.48%      2.25s  6.33%  runtime.tryDeferToSpanScan
     0.71s  2.00% 53.48%      0.71s  2.00%  runtime.spanClass.sizeclass (inline)
     0.47s  1.32% 54.80%      0.47s  1.32%  memeqbody
     0.44s  1.24% 56.04%      2.10s  5.91%  runtime.mapaccess1_faststr
     0.38s  1.07% 57.11%      0.57s  1.60%  example.com/isureal-company/isureal.postLess
     0.33s  0.93% 58.04%      0.33s  0.93%  runtime.nextFreeFast (inline)
     0.28s  0.79% 58.82%      3.95s 11.12%  runtime.mallocgc
     0.26s  0.73% 59.56%      0.69s  1.94%  database/sql.convertAssignRows
     0.26s  0.73% 60.29%      0.26s  0.73%  runtime.futex
     0.24s  0.68% 60.96%      0.24s  0.68%  sync/atomic.(*Int32).Add
     0.21s  0.59% 61.55%      1.07s  3.01%  runtime.mallocgcSmallScanNoHeader
     0.20s  0.56% 62.12%      0.23s  0.65%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.20s  0.56% 62.68%      0.44s  1.24%  runtime.greyobject
     0.19s  0.53% 63.21%      0.19s  0.53%  internal/runtime/maps.(*Map).directoryAt (inline)
     0.19s  0.53% 63.75%      0.70s  1.97%  runtime.mapaccess2_faststr
     0.18s  0.51% 64.26%      0.19s  0.53%  encoding/json.appendString[go.shape.string]
     0.18s  0.51% 64.76%      0.24s  0.68%  runtime.heapArenaOf (inline)
     0.18s  0.51% 65.27%      0.19s  0.53%  runtime.typePointers.next
     0.16s  0.45% 65.72%      0.33s  0.93%  compress/flate.(*compressor).deflate
     0.15s  0.42% 66.14%      3.09s  8.70%  crypto/internal/fips140/sha256.(*Digest).Write
     0.15s  0.42% 66.56%      0.77s  2.17%  encoding/json.structEncoder.encode
     0.15s  0.42% 66.99%      0.31s  0.87%  internal/runtime/maps.(*Iter).Next
     0.15s  0.42% 67.41%      1.80s  5.07%  runtime.mapassign_faststr
     0.14s  0.39% 67.80%      2.68s  7.54%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.13s  0.37% 68.17%      0.82s  2.31%  database/sql.(*Rows).scanLocked
     0.13s  0.37% 68.53%      0.26s  0.73%  runtime.scanObjectSmall
     0.13s  0.37% 68.90%      0.42s  1.18%  runtime.scanObjectsSmall
     0.12s  0.34% 69.24%      2.49s  7.01%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.12s  0.34% 69.58%      0.25s   0.7%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
     0.12s  0.34% 69.91%      0.32s   0.9%  internal/stringslite.Index
     0.12s  0.34% 70.25%      0.20s  0.56%  runtime.(*mspan).markBitsForIndex (inline)
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 54.7% / 100% | 27% | 3.5% | 9 | 293 MB | run queue が積もっている |
| isu2 (2 cpu) | 17.9% / 95% | 1% | 0.6% | 5 | 1702 MB | run queue が積もっている |
| isu3 (2 cpu) | 15.2% / 96% | 1% | 0.6% | 5 | 1633 MB | run queue が積もっている |
| isu4 (2 cpu) | 51.6% / 100% | 23% | 16.7% | 37 | 108 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 72% / nginx 11% / find 2% |
| isu2 | isureal 33% / nginx 1% |
| isu3 | isureal 28% / nginx 1% |
| isu4 | postgres 58% / isureal 6% |

## Diff from previous run

スコア: 83744 → 115400 (+31656)

`5a138c8..f256785`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf |   2 +
 conf/isu1/home/isucon/env.sh                   |   1 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |  13 +++
 conf/isu2/home/isucon/env.sh                   |   2 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |  13 +++
 conf/isu3/home/isucon/env.sh                   |   2 +
 webapp/go/account_deletion_api.go              |   1 +
 webapp/go/action_log.go                        |  13 ++-
 webapp/go/action_log_test.go                   |  67 +++++++++++++++
 webapp/go/image_local.go                       |   3 +-
 webapp/go/offload.go                           |  47 +++++++---
 webapp/go/post_api.go                          |   2 +-
 webapp/go/post_forward.go                      | 113 +++++++++++++++++++++++++
 webapp/go/post_forward_test.go                 |  94 ++++++++++++++++++++
 webapp/go/profile_image_api.go                 |   7 +-
 webapp/go/store.go                             |  53 ++++++++++--
 webapp/go/user_api.go                          |   1 +
 webapp/go/user_read_api.go                     |  58 +++++++------
 webapp/go/user_search_test.go                  |  94 ++++++++++++++++++++
 19 files changed, 532 insertions(+), 54 deletions(-)
```
