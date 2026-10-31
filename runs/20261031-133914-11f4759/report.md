# run 20261031-133914-11f4759

## Summary

| | |
|---|---|
| score | 109728 |
| status | pass |
| ref | nginx-buffers-tcp |
| sha | 11f4759553c3 |
| 所要時間 | 539s |
| 前回比 | +25984 (31.0%) vs 20261031-132303-5a138c8 |
| ベスト比 | +25984 (31.0%) vs 20261031-132303-5a138c8 = 83744 **ベスト更新** |
| note | nginx-no-tempfile + tcp_slow_start_after_idle=0 on 7e687d9 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	143	406	932	932
europe	7	20	58	159	159	158
america	6	17	51	123	125	125
```

```
{"id":"14163","status":"done","portalStatus":"finished","passed":true,"score":109728,"raw":109728,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:46:47.446Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:46:50.422Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:46:53.418Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:46:56.438Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T04:46:59.417Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T04:47:02.424Z level=INFO msg="america 2 日目: 17 人"
time=2026-10-31T04:47:05.441Z level=INFO msg="asia 3 日目: 143 人"
time=2026-10-31T04:47:06.401Z level=INFO msg="user moved" user=semarha_5671 from=asia to=america
time=2026-10-31T04:47:08.011Z level=INFO msg="user moved" user=semarha_5671 from=america to=europe
time=2026-10-31T04:47:08.432Z level=INFO msg="europe 3 日目: 58 人"
time=2026-10-31T04:47:08.454Z level=INFO msg="user moved" user=yuwa_5972 from=asia to=europe
time=2026-10-31T04:47:11.420Z level=INFO msg="america 3 日目: 51 人"
time=2026-10-31T04:47:14.447Z level=INFO msg="asia 4 日目: 406 人"
time=2026-10-31T04:47:17.506Z level=INFO msg="europe 4 日目: 159 人"
time=2026-10-31T04:47:18.406Z level=INFO msg="user moved" user=sotoror-6861 from=asia to=america
time=2026-10-31T04:47:18.516Z level=INFO msg="user moved" user=zoewa6616 from=america to=europe
time=2026-10-31T04:47:20.403Z level=INFO msg="user moved" user=suivy-7026 from=asia to=america
time=2026-10-31T04:47:20.607Z level=INFO msg="america 4 日目: 123 人"
time=2026-10-31T04:47:22.254Z level=INFO msg="user moved" user=taben-7742 from=asia to=america
time=2026-10-31T04:47:22.620Z level=INFO msg="user moved" user=zoewa6616 from=europe to=america
time=2026-10-31T04:47:23.818Z level=INFO msg="asia 5 日目: 932 人"
time=2026-10-31T04:47:26.509Z level=INFO msg="europe 5 日目: 159 人"
time=2026-10-31T04:47:29.457Z level=INFO msg="america 5 日目: 125 人"
time=2026-10-31T04:47:38.474Z level=INFO msg="europe 最終日（グランプリ）: 158 人"
time=2026-10-31T04:47:38.800Z level=INFO msg="asia 最終日（グランプリ）: 932 人"
time=2026-10-31T04:47:38.852Z level=INFO msg="america 最終日（グランプリ）: 125 人"
time=2026-10-31T04:47:47.450Z level=INFO msg="グランプリ: アジア 454 人, ヨーロッパ 89 人, アメリカ 46 人"
{
  "pass": true,
  "score": 109728,
  "messages": [
    "reel_10053: 認定公開アカウントをフォローできません: PUT /api/follows/28505755-2464-4e8b-b73b-00e18e60d3b8: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
31.6%	1418	3	0	GET	/api/notifications	11712.385	8.260	14.991	15.253	161
14.5%	1984	44	0	POST	/api/posts	5385.628	2.715	9.763	12.510	635
14.5%	4603	66	0	POST	/api/friend-requests	5370.456	1.167	8.333	15.548	114
8.8%	12380	704	0	GET	^/api/images/[^/]+	3274.001	0.264	4.839	13.820	132636
7.6%	5278	1404	0	POST	/api/auth/login	2795.715	0.530	6.895	14.060	186
4.7%	1476	6	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1756.836	1.190	8.397	13.684	116
4.0%	953	1	0	POST	/api/auth/register	1477.798	1.551	8.765	11.950	227
2.5%	633	118	0	PUT	/api/me/profile-image	939.077	1.484	8.061	11.096	4
2.2%	655	0	0	PUT	^/api/follows/[^/]+$	800.405	1.222	7.553	17.708	0
1.8%	456	25	0	POST	^/api/posts/[^/]+/realmojis$	649.512	1.424	8.225	11.549	452
1.2%	409	0	0	PATCH	/api/me	454.774	1.112	8.013	11.023	273
1.2%	406	18	0	POST	^/api/posts/[^/]+/comments$	440.305	1.084	7.789	12.534	330
0.9%	3188	0	0	GET	/api/memories	328.342	0.103	0.695	0.865	10712
0.7%	173	6	0	PATCH	^/api/posts/[^/]+$	260.839	1.508	8.654	12.225	624
0.7%	10206	0	0	GET	^/api/feeds/[^/]+	252.997	0.025	0.176	2.928	1723
0.6%	585	0	0	POST	/api/auth/logout	210.769	0.360	5.861	8.671	0
0.5%	100	0	0	DELETE	^/api/posts/[^/]+$	199.430	1.994	7.772	9.347	0
0.4%	206	0	0	GET	/api/me/export	165.817	0.805	5.224	7.332	1905945
0.3%	2885	0	0	GET	/api/posting-window/current	98.939	0.034	0.204	1.911	237
0.2%	62	0	0	DELETE	^/api/realmojis/[^/]+$	67.760	1.093	9.776	9.776	0
0.2%	82	0	0	DELETE	^/api/comments/[^/]+$	57.936	0.707	7.879	7.879	0
0.1%	1683	79	0	GET	^/api/posts/me/current$	51.491	0.031	0.191	0.386	619
0.1%	2181	0	0	GET	/api/users	51.305	0.024	0.155	0.476	244
0.1%	33	0	0	PUT	^/api/blocks/[^/]+$	36.949	1.120	12.565	12.565	0
0.1%	588	0	0	GET	^/api/memories/[^/]+	33.271	0.057	0.622	0.834	660
0.1%	1399	0	0	GET	^/api/relationships/[^/]+	26.849	0.019	0.151	0.305	895
0.1%	1326	12	0	GET	^/api/posts/[^/]+$	24.808	0.019	0.124	0.291	619
0.0%	7	0	0	DELETE	^/api/follows/[^/]+$	14.912	2.130	6.007	6.007	0
0.0%	14	0	0	PUT	^/api/admin/users/[^/]+/certification$	13.664	0.976	11.412	11.412	0
0.0%	1	0	0	POST	/api/initialize	12.505	12.505	12.505	12.505	90
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.9%	373	1	0	GET	/api/notifications	3037.880	8.144	11.620	12.002	176
15.2%	475	5	0	POST	/api/posts	1248.667	2.629	8.865	9.411	636
12.0%	884	10	0	POST	/api/friend-requests	985.464	1.115	8.647	12.800	114
5.9%	542	0	0	POST	/api/auth/login	484.892	0.895	7.716	10.550	248
5.1%	1928	16	0	GET	^/api/images/[^/]+	417.035	0.216	1.313	1.768	196881
4.3%	1759	0	0	GET	^/api/feeds/[^/]+	354.155	0.201	0.757	1.367	2065
3.7%	233	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	301.600	1.294	8.295	11.508	116
2.1%	648	0	0	GET	/api/posting-window/current	170.522	0.263	0.693	1.035	239
1.9%	589	0	0	GET	/api/memories	153.546	0.261	0.806	1.456	10350
1.4%	90	8	0	POST	^/api/posts/[^/]+/comments$	113.987	1.267	9.656	9.656	319
1.3%	416	33	0	GET	^/api/posts/me/current$	106.539	0.256	0.773	1.112	593
1.1%	75	7	0	POST	^/api/posts/[^/]+/realmojis$	89.974	1.200	7.347	7.347	443
1.1%	112	0	0	PUT	^/api/follows/[^/]+$	89.828	0.802	9.121	11.651	0
1.1%	74	0	0	POST	/api/auth/logout	87.889	1.188	7.059	7.059	0
1.0%	434	0	0	GET	/api/users	84.033	0.194	0.655	0.751	241
1.0%	155	0	0	POST	/api/auth/register	81.771	0.528	1.914	2.459	228
0.9%	45	3	0	PATCH	^/api/posts/[^/]+$	70.259	1.561	14.042	14.042	607
0.8%	35	0	0	GET	/api/me/export	63.608	1.817	5.743	5.743	2146829
0.6%	80	0	0	PUT	/api/me/profile-image	48.469	0.606	2.831	2.831	0
0.5%	73	0	0	PATCH	/api/me	39.555	0.542	6.194	6.194	270
0.5%	226	0	0	GET	^/api/relationships/[^/]+	38.929	0.172	0.587	0.697	796
0.4%	223	10	0	GET	^/api/posts/[^/]+$	36.785	0.165	0.550	0.652	590
0.4%	18	0	0	DELETE	^/api/posts/[^/]+$	28.979	1.610	8.274	8.274	0
0.2%	114	0	0	GET	^/api/users/[^/]+$	18.979	0.166	0.540	0.571	195
0.2%	65	0	0	GET	^/api/memories/[^/]+	14.749	0.227	0.747	0.747	1027
0.2%	93	0	0	GET	^/api/posts/[^/]+/realmojis$	14.468	0.156	0.597	0.597	1197
0.2%	84	2	0	GET	/api/me	14.008	0.167	0.456	0.456	295
0.2%	83	0	0	GET	^/api/posts/[^/]+/comments$	12.852	0.155	0.636	0.636	1297
0.1%	5	0	0	PUT	^/api/blocks/[^/]+$	5.552	1.110	2.807	2.807	0
0.0%	8	0	0	DELETE	^/api/realmojis/[^/]+$	4.024	0.503	2.265	2.265	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.7%	262	0	0	GET	/api/notifications	1958.077	7.474	9.595	9.621	174
14.3%	669	13	0	POST	/api/friend-requests	855.360	1.279	8.299	14.482	113
13.2%	329	5	0	POST	/api/posts	792.381	2.408	7.968	8.194	635
6.7%	382	0	0	POST	/api/auth/login	403.360	1.056	7.106	11.476	249
5.1%	1327	10	0	GET	^/api/images/[^/]+	307.840	0.232	1.175	2.425	184745
4.8%	210	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	285.078	1.358	7.859	8.408	116
3.3%	1015	0	0	GET	^/api/feeds/[^/]+	196.415	0.194	0.675	1.118	2716
2.0%	66	0	0	PUT	/api/me/profile-image	119.153	1.805	9.276	9.276	0
1.9%	452	0	0	GET	/api/posting-window/current	114.666	0.254	0.803	0.882	240
1.7%	61	8	0	POST	^/api/posts/[^/]+/realmojis$	102.895	1.687	10.806	10.806	414
1.7%	398	0	0	GET	/api/memories	99.892	0.251	1.097	1.230	12815
1.5%	64	0	0	POST	/api/auth/logout	88.211	1.378	12.350	12.350	0
1.5%	121	0	0	POST	/api/auth/register	87.412	0.722	6.641	6.728	229
1.2%	345	0	0	GET	/api/users	70.930	0.206	0.688	1.124	239
1.1%	68	0	0	PUT	^/api/follows/[^/]+$	65.761	0.967	5.076	5.076	0
1.1%	52	0	0	PATCH	/api/me	65.391	1.258	5.204	5.204	285
1.0%	278	23	0	GET	^/api/posts/me/current$	62.362	0.224	0.678	0.788	594
1.0%	29	0	0	GET	/api/me/export	57.641	1.988	10.291	10.291	2530752
1.0%	82	5	0	POST	^/api/posts/[^/]+/comments$	57.493	0.701	3.571	3.571	314
0.8%	30	2	0	PATCH	^/api/posts/[^/]+$	49.726	1.658	7.739	7.739	603
0.5%	172	0	0	GET	^/api/relationships/[^/]+	31.457	0.183	0.636	0.704	1169
0.5%	169	6	0	GET	^/api/posts/[^/]+$	27.555	0.163	0.623	0.623	603
0.4%	12	0	0	DELETE	^/api/posts/[^/]+$	22.664	1.889	5.390	5.390	0
0.2%	69	0	0	GET	/api/me	12.872	0.187	0.745	0.745	302
0.2%	50	0	0	GET	^/api/memories/[^/]+	12.535	0.251	0.852	0.852	1055
0.2%	81	0	0	GET	^/api/users/[^/]+$	12.317	0.152	0.345	0.345	197
0.2%	7	0	0	DELETE	^/api/realmojis/[^/]+$	9.435	1.348	4.192	4.192	0
0.2%	59	1	0	GET	^/api/posts/[^/]+/comments$	9.047	0.153	0.514	0.514	791
0.1%	11	0	0	DELETE	^/api/comments/[^/]+$	8.191	0.745	3.624	3.624	0
0.1%	54	0	0	GET	^/api/posts/[^/]+/realmojis$	7.833	0.145	0.433	0.433	862
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 66089 クエリ / 81 種類 / 合計 446.350s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	65.4%	3950	291.829	73.88	644.89	654.2	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	19.5%	270	86.835	321.61	1560.59	1255.5	127.9	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	6.0%	2793	26.854	9.61	585.63	120.6	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	4.5%	8093	20.077	2.48	51.31	10.7	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.6%	2735	7.225	2.64	31.48	24.9	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.5%	1	2.323	2322.83	2322.83	30842.0	1199619.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.4%	552	1.884	3.41	16.81	29.5	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
8	0.4%	1	1.611	1611.15	1611.15	23966.0	626802.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
9	0.3%	6056	1.209	0.20	20.12	19.6	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
10	0.2%	547	0.944	1.73	10.90	20.8	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.2%	1	0.759	759.41	759.41	12545.0	529855.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
12	0.2%	1851	0.698	0.38	25.75	20.7	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
13	0.1%	4798	0.559	0.12	6.40	11.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
14	0.1%	18	0.486	27.02	262.26	284.9	1.0	INSERT INTO posting_windows (id, region, local_date, starts_at, on_time_deadline, created_at, notification_ids, notified_user_ids) VALUES ($1, $2, $3, $4, $5, $…
15	0.1%	2	0.347	173.57	178.24	225715.5	2657.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id

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
Build ID: 321cd7dd08b1e7b091cd431ea2a675ad8aae2999
Type: cpu
Time: 2026-10-31 13:46:30 JST
Duration: 60.04s, Total samples = 48.85s (81.36%)
Showing nodes accounting for 35.60s, 72.88% of 48.85s total
Dropped 1047 nodes (cum <= 0.24s)
Showing top 40 nodes out of 291
      flat  flat%   sum%        cum   cum%
     4.27s  8.74%  8.74%      4.27s  8.74%  internal/runtime/syscall/linux.Syscall6
     3.39s  6.94% 15.68%      3.39s  6.94%  crypto/internal/fips140/sha256.blockAVX2
     2.32s  4.75% 20.43%      2.32s  4.75%  runtime.memclrNoHeapPointers
     2.25s  4.61% 25.04%      2.25s  4.61%  runtime.memmove
     2.18s  4.46% 29.50%      2.91s  5.96%  runtime.findObject
     2.12s  4.34% 33.84%      8.35s 17.09%  runtime.scanObject
     1.76s  3.60% 37.44%      5.98s 12.24%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.73s  3.54% 40.98%      3.01s  6.16%  runtime.tryDeferToSpanScan
     1.56s  3.19% 44.18%      1.56s  3.19%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.31s  2.68% 46.86%      1.31s  2.68%  example.com/isureal-company/isureal.(*UserRow).active
     1.28s  2.62% 49.48%      1.28s  2.62%  aeshashbody
     1.21s  2.48% 51.95%      3.75s  7.68%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.99s  2.03% 53.98%      1.64s  3.36%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.97s  1.99% 55.97%      0.97s  1.99%  runtime.spanClass.sizeclass (inline)
     0.94s  1.92% 57.89%      0.94s  1.92%  memeqbody
     0.86s  1.76% 59.65%      0.90s  1.84%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     0.55s  1.13% 60.78%      0.69s  1.41%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.52s  1.06% 61.84%      2.12s  4.34%  runtime.mapaccess1_faststr
     0.48s  0.98% 62.82%      0.48s  0.98%  indexbody
     0.48s  0.98% 63.81%      0.48s  0.98%  runtime.nextFreeFast (inline)
     0.38s  0.78% 64.59%      0.59s  1.21%  example.com/isureal-company/isureal.postLess
     0.31s  0.63% 65.22%      0.71s  1.45%  database/sql.convertAssignRows
     0.31s  0.63% 65.85%      0.31s  0.63%  runtime.futex
     0.31s  0.63% 66.49%      0.31s  0.63%  sync/atomic.(*Int32).Add
     0.30s  0.61% 67.10%      0.69s  1.41%  runtime.greyobject
     0.29s  0.59% 67.70%      0.34s   0.7%  runtime.spanOf (inline)
     0.25s  0.51% 68.21%      0.43s  0.88%  compress/flate.(*compressor).deflate
     0.25s  0.51% 68.72%      0.25s  0.51%  runtime.(*mspan).divideByElemSize (inline)
     0.21s  0.43% 69.15%      5.71s 11.69%  runtime.mallocgc
     0.20s  0.41% 69.56%      0.28s  0.57%  runtime.(*mspan).markBitsForIndex (inline)
     0.20s  0.41% 69.97%      1.88s  3.85%  runtime.mapassign_faststr
     0.19s  0.39% 70.36%      1.27s  2.60%  runtime.mallocgcSmallScanNoHeader
     0.18s  0.37% 70.73%      2.54s  5.20%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.17s  0.35% 71.07%      0.64s  1.31%  runtime.scanObjectsSmall
     0.16s  0.33% 71.40%      2.93s  6.00%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.16s  0.33% 71.73%      0.30s  0.61%  internal/runtime/maps.(*Iter).Next
     0.15s  0.31% 72.04%      3.65s  7.47%  crypto/internal/fips140/sha256.(*Digest).Write
     0.14s  0.29% 72.32%      0.85s  1.74%  runtime.mallocgcSmallNoscan
     0.14s  0.29% 72.61%      0.63s  1.29%  runtime.wbBufFlush1
     0.13s  0.27% 72.88%      0.72s  1.47%  encoding/json.structEncoder.encode
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 56.1% / 100% | 35% | 0.8% | 11 | 381 MB | CPU が飽和 |
| isu2 (2 cpu) | 21.9% / 98% | 1% | 0.6% | 4 | 1873 MB |  |
| isu3 (2 cpu) | 17.7% / 100% | 2% | 0.6% | 3 | 1658 MB |  |
| isu4 (2 cpu) | 51.2% / 100% | 22% | 15.3% | 25 | 112 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 75% / nginx 12% / journalctl 3% / systemd-journal 2% / kcompactd0 1% |
| isu2 | isureal 41% / nginx 1% |
| isu3 | isureal 33% / nginx 1% |
| isu4 | postgres 62% / isureal 4% |

## Diff from previous run

スコア: 83744 → 109728 (+25984)

`5a138c8..11f4759`

```
 conf/isu1/etc/nginx/sites-enabled/isureal.conf |   2 +
 conf/isu1/home/isucon/env.sh                   |   1 +
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |  13 +++
 conf/isu2/home/isucon/env.sh                   |   2 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |  13 +++
 conf/isu3/home/isucon/env.sh                   |   2 +
 webapp/go/image_local.go                       |   3 +-
 webapp/go/post_api.go                          |   2 +-
 webapp/go/post_forward.go                      | 113 +++++++++++++++++++++++++
 webapp/go/post_forward_test.go                 |  94 ++++++++++++++++++++
 10 files changed, 243 insertions(+), 2 deletions(-)
```
