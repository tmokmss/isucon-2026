# run 20261031-111806-214e79f

## Summary

| | |
|---|---|
| score | 21074 |
| status | pass |
| ref | HEAD |
| sha | 214e79f301c0 |
| 所要時間 | 371s |
| 前回比 | +5406 (34.5%) vs 20261031-110450-8f720c2 |
| ベスト比 | -3727 (-15.0%) vs 20261031-103811-cef5fe3 = 24801 |
| note | 画像の無効化を同期にする（受領待ち）、init/ 全件書き出し済み |

## Bench

```
{"id":"3653","status":"done","portalStatus":"finished","passed":true,"score":21074,"raw":21074,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:23:00.464Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:23:02.822Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:23:05.048Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:23:08.730Z level=INFO msg="asia 2 日目: 20 人"
time=2026-10-31T02:23:10.940Z level=INFO msg="europe 2 日目: 8 人"
time=2026-10-31T02:23:13.953Z level=INFO msg="america 2 日目: 9 人"
time=2026-10-31T02:23:17.654Z level=INFO msg="asia 3 日目: 44 人"
time=2026-10-31T02:23:20.168Z level=INFO msg="europe 3 日目: 21 人"
time=2026-10-31T02:23:23.550Z level=INFO msg="america 3 日目: 20 人"
time=2026-10-31T02:23:24.735Z level=INFO msg="user moved" user=nemaris5377 from=asia to=america
time=2026-10-31T02:23:28.352Z level=INFO msg="asia 4 日目: 105 人"
time=2026-10-31T02:23:29.837Z level=INFO msg="europe 4 日目: 27 人"
time=2026-10-31T02:23:33.004Z level=INFO msg="america 4 日目: 21 人"
time=2026-10-31T02:23:36.219Z level=INFO msg="user moved" user=rihamar-5557 from=asia to=america
time=2026-10-31T02:23:36.887Z level=INFO msg="asia 5 日目: 104 人"
time=2026-10-31T02:23:42.009Z level=INFO msg="europe 5 日目: 27 人"
time=2026-10-31T02:23:43.627Z level=INFO msg="america 5 日目: 22 人"
time=2026-10-31T02:23:51.104Z level=INFO msg="europe 最終日（グランプリ）: 27 人"
time=2026-10-31T02:23:52.676Z level=INFO msg="asia 最終日（グランプリ）: 103 人"
time=2026-10-31T02:23:53.551Z level=INFO msg="america 最終日（グランプリ）: 22 人"
time=2026-10-31T02:23:58.905Z level=INFO msg="グランプリ: アジア 1 人, ヨーロッパ 15 人, アメリカ 0 人"
{
  "pass": true,
  "score": 21074,
  "messages": [
    "orleo.5029: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.0%	269	1	0	GET	/api/notifications	2585.949	9.613	17.640	17.798	171
15.9%	325	28	0	POST	/api/posts	1142.736	3.516	10.823	11.368	575
10.7%	1906	845	0	POST	/api/auth/login	770.738	0.404	4.062	6.792	150
6.3%	3545	286	0	GET	^/api/images/[^/]+	455.773	0.129	1.591	5.215	93239
5.7%	537	18	0	POST	/api/friend-requests	410.239	0.764	5.669	7.756	112
5.3%	186	141	0	PUT	/api/me/profile-image	379.196	2.039	9.999	10.056	15
3.9%	296	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	283.986	0.959	8.739	10.776	116
3.5%	122	1	0	POST	/api/auth/register	249.903	2.048	10.531	11.141	229
2.0%	1889	0	0	GET	^/api/feeds/[^/]+	144.084	0.076	0.384	0.616	1650
1.9%	929	0	0	GET	/api/memories	135.482	0.146	0.958	2.999	11537
1.4%	84	2	0	POST	^/api/posts/[^/]+/realmojis$	98.875	1.177	5.043	5.043	450
1.2%	424	0	0	GET	/api/posting-window/current	84.062	0.198	1.023	1.767	236
1.0%	76	0	0	GET	/api/me/export	73.018	0.961	11.320	11.320	2015848
0.7%	86	0	0	PUT	^/api/follows/[^/]+$	53.281	0.620	7.903	7.903	0
0.6%	53	1	0	POST	^/api/posts/[^/]+/comments$	43.960	0.829	3.881	3.881	320
0.6%	309	0	0	GET	^/api/relationships/[^/]+	43.384	0.140	0.516	2.691	1545
0.4%	200	0	0	POST	/api/auth/logout	28.468	0.142	1.512	4.910	0
0.4%	262	0	0	GET	/api/users	27.537	0.105	0.763	0.968	234
0.4%	18	0	0	POST	/api/admin/posting-window	25.928	1.440	4.062	4.062	176
0.3%	235	7	0	GET	^/api/posts/me/current$	23.762	0.101	0.427	0.521	609
0.3%	48	0	0	PATCH	/api/me	19.837	0.413	1.891	1.891	270
0.2%	243	4	0	GET	^/api/posts/[^/]+$	17.610	0.072	0.299	0.311	600
0.2%	1	0	0	POST	/api/initialize	12.263	12.263	12.263	12.263	90
0.2%	16	0	0	DELETE	^/api/posts/[^/]+$	11.197	0.700	2.227	2.227	0
0.1%	4	0	0	DELETE	^/api/follows/[^/]+$	9.492	2.373	7.752	7.752	0
0.1%	192	0	0	GET	^/api/memories/[^/]+	9.288	0.048	0.642	0.971	584
0.1%	145	0	0	GET	^/api/users/[^/]+$	8.587	0.059	0.279	0.308	193
0.1%	27	3	0	PATCH	^/api/posts/[^/]+$	7.971	0.295	1.282	1.282	546
0.1%	101	0	0	GET	^/api/posts/[^/]+/realmojis$	5.600	0.055	0.320	0.338	437
0.1%	52	3	0	GET	/api/me	5.387	0.104	0.500	0.500	280
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.2%	91	0	0	GET	/api/notifications	791.620	8.699	12.520	12.520	177
13.7%	107	3	0	POST	/api/posts	264.235	2.469	9.481	11.720	620
5.6%	168	8	0	POST	/api/friend-requests	107.501	0.640	7.630	7.906	111
5.3%	146	0	0	POST	/api/auth/login	101.033	0.692	5.966	6.390	252
5.0%	575	5	0	GET	^/api/images/[^/]+	95.633	0.166	0.955	1.267	169436
5.0%	404	0	0	GET	^/api/feeds/[^/]+	95.515	0.236	1.313	2.813	1890
4.3%	80	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	83.632	1.045	6.864	6.864	116
3.8%	215	0	0	GET	/api/memories	72.936	0.339	1.590	3.818	13327
2.5%	25	0	0	POST	/api/auth/register	48.381	1.935	11.071	11.071	229
1.5%	130	0	0	GET	/api/posting-window/current	29.703	0.228	0.600	0.832	239
1.3%	10	1	0	POST	^/api/posts/[^/]+/realmojis$	24.619	2.462	6.290	6.290	481
1.3%	33	3	0	POST	^/api/posts/[^/]+/comments$	24.232	0.734	5.999	5.999	300
1.2%	95	3	0	GET	^/api/posts/me/current$	23.825	0.251	1.286	1.286	618
1.2%	83	0	0	GET	^/api/relationships/[^/]+	23.103	0.278	1.772	1.772	1421
1.2%	8	0	0	DELETE	^/api/posts/[^/]+$	22.324	2.791	6.912	6.912	0
1.0%	18	0	0	PUT	^/api/follows/[^/]+$	19.680	1.093	6.355	6.355	0
0.8%	8	0	0	GET	/api/me/export	16.186	2.023	4.114	4.114	2184962
0.8%	26	0	0	POST	/api/auth/logout	14.603	0.562	4.860	4.860	0
0.7%	77	0	0	GET	/api/users	12.926	0.168	0.583	0.583	225
0.5%	46	4	0	GET	^/api/posts/[^/]+$	9.524	0.207	0.896	0.896	560
0.4%	35	0	0	GET	^/api/users/[^/]+$	7.676	0.219	0.882	0.882	188
0.4%	6	1	0	PATCH	^/api/posts/[^/]+$	7.595	1.266	5.343	5.343	576
0.4%	11	0	0	PUT	/api/me/profile-image	6.805	0.619	2.318	2.318	0
0.2%	19	0	0	GET	^/api/memories/[^/]+	3.533	0.186	0.380	0.380	609
0.2%	2	0	0	DELETE	^/api/realmojis/[^/]+$	3.322	1.661	2.739	2.739	0
0.2%	15	1	0	GET	/api/me	3.281	0.219	0.958	0.958	264
0.1%	4	0	0	DELETE	^/api/comments/[^/]+$	2.725	0.681	2.233	2.233	0
0.1%	14	0	0	GET	^/api/posts/[^/]+/comments$	2.261	0.161	0.360	0.360	634
0.1%	7	0	0	PATCH	/api/me	1.734	0.248	0.488	0.488	316
0.1%	2	0	0	PUT	^/api/blocks/[^/]+$	1.638	0.819	1.530	1.530	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
46.3%	88	1	0	GET	/api/notifications	773.033	8.784	10.622	10.622	178
19.1%	83	3	0	POST	/api/posts	318.763	3.841	9.609	9.609	641
4.7%	471	2	0	GET	^/api/images/[^/]+	78.110	0.166	0.957	1.250	180418
4.5%	123	0	0	POST	/api/auth/login	74.475	0.605	5.531	6.987	255
3.8%	321	0	0	GET	^/api/feeds/[^/]+	64.188	0.200	0.664	1.031	1241
3.7%	116	11	0	POST	/api/friend-requests	61.134	0.527	4.597	7.709	107
3.1%	162	0	0	GET	/api/memories	51.257	0.316	0.905	1.474	13735
2.6%	13	0	0	GET	/api/me/export	43.322	3.332	11.791	11.791	2811756
2.3%	26	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	38.471	1.480	8.642	8.642	116
2.0%	99	0	0	GET	/api/posting-window/current	33.404	0.337	1.054	1.054	239
1.1%	22	1	0	POST	^/api/posts/[^/]+/realmojis$	18.196	0.827	2.979	2.979	391
1.0%	64	8	0	GET	^/api/posts/me/current$	17.312	0.271	1.023	1.023	575
1.0%	68	0	0	GET	^/api/relationships/[^/]+	16.943	0.249	0.749	0.749	773
0.6%	58	0	0	GET	/api/users	9.949	0.172	0.470	0.470	254
0.6%	14	1	0	POST	^/api/posts/[^/]+/comments$	9.360	0.669	2.118	2.118	272
0.5%	7	0	0	PATCH	^/api/posts/[^/]+$	9.169	1.310	2.323	2.323	626
0.5%	48	5	0	GET	^/api/posts/[^/]+$	8.730	0.182	0.703	0.703	560
0.4%	1	0	0	POST	/api/me/account-deletion	7.275	7.275	7.275	7.275	0
0.4%	19	0	0	POST	/api/auth/register	6.224	0.328	0.711	0.711	230
0.3%	35	0	0	GET	^/api/users/[^/]+$	5.675	0.162	0.595	0.595	187
0.3%	11	0	0	PUT	/api/me/profile-image	5.383	0.489	1.041	1.041	0
0.3%	21	0	0	POST	/api/auth/logout	4.592	0.219	0.718	0.718	0
0.2%	21	0	0	GET	^/api/memories/[^/]+	3.493	0.166	0.311	0.311	750
0.2%	18	0	0	GET	^/api/posts/[^/]+/realmojis$	3.342	0.186	0.793	0.793	346
0.1%	14	0	0	GET	/api/me	2.188	0.156	0.297	0.297	293
0.1%	8	0	0	PATCH	/api/me	2.029	0.254	0.484	0.484	316
0.1%	10	0	0	PUT	^/api/follows/[^/]+$	1.580	0.158	0.274	0.274	0
0.1%	1	0	0	DELETE	^/api/posts/[^/]+$	1.258	1.258	1.258	1.258	0
0.1%	7	0	0	GET	^/api/posts/[^/]+/comments$	1.102	0.157	0.276	0.276	248
0.0%	1	0	0	DELETE	^/api/realmojis/[^/]+$	0.573	0.573	0.573	0.573	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16305 クエリ / 79 種類 / 合計 76.715s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	53.4%	97	40.972	422.39	7266.70	6770.3	136.4	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	10.9%	665	8.332	12.53	165.58	701.1	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	9.1%	16	6.993	437.08	2132.27	36292.3	1855.9	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4
4	8.4%	931	6.461	6.94	222.32	116.2	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	6.1%	2599	4.701	1.81	10.15	11.8	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
6	3.1%	1	2.388	2388.13	2388.13	29348.0	1199407.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	2.1%	1	1.618	1618.35	1618.35	19695.0	625549.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	1.4%	1	1.097	1097.01	1097.01	11528.0	529669.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	1.1%	485	0.841	1.73	12.40	31.4	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.5%	2	0.356	178.05	187.19	214661.0	2525.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.4%	112	0.335	2.99	6.99	44.1	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
12	0.4%	2	0.293	146.56	281.44	21004.5	2519.5	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
13	0.4%	1	0.280	279.69	279.69	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
14	0.3%	1	0.225	224.74	224.74	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
15	0.3%	1	0.208	208.32	208.32	69894.0	29782.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 3
INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT unnest($1::text[]), unnest($2::text[]), $3, $4;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 7cd16b6bea0378f791aec0ac2e89144ff6b9f9a3
Type: cpu
Time: 2026-10-31 11:22:41 JST
Duration: 60.27s, Total samples = 85.75s (142.28%)
Showing nodes accounting for 75.89s, 88.50% of 85.75s total
Dropped 876 nodes (cum <= 0.43s)
Showing top 40 nodes out of 175
      flat  flat%   sum%        cum   cum%
       20s 23.32% 23.32%        20s 23.32%  crypto/internal/fips140/sha256.blockAVX2
    11.01s 12.84% 36.16%     11.01s 12.84%  runtime.spanClass.sizeclass (inline)
     6.92s  8.07% 44.23%     10.59s 12.35%  image/jpeg.(*decoder).reconstructBlock
     3.62s  4.22% 48.45%     15.33s 17.88%  runtime.tryDeferToSpanScan
     2.94s  3.43% 51.88%     22.26s 25.96%  image/jpeg.(*decoder).processSOS
     2.75s  3.21% 55.09%      2.75s  3.21%  runtime.memmove
     2.49s  2.90% 57.99%      4.07s  4.75%  image/jpeg.(*decoder).decodeHuffman
     2.35s  2.74% 60.73%      2.35s  2.74%  runtime.memclrNoHeapPointers
     2.01s  2.34% 63.08%      2.12s  2.47%  image/jpeg.(*decoder).receiveExtend
     1.63s  1.90% 64.98%      6.95s  8.10%  runtime.scanObjectsSmall
     1.46s  1.70% 66.68%      1.46s  1.70%  internal/runtime/syscall/linux.Syscall6
     1.45s  1.69% 68.37%     12.15s 14.17%  runtime.scanObject
     1.36s  1.59% 69.96%      1.36s  1.59%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.35s  1.57% 71.53%      1.86s  2.17%  image/jpeg.idctCols
     1.26s  1.47% 73.00%      5.38s  6.27%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.25s  1.46% 74.46%     22.67s 26.44%  crypto/internal/fips140/sha256.(*Digest).checkSum
     1.19s  1.39% 75.85%      1.73s  2.02%  image/jpeg.idctRows
     1.16s  1.35% 77.20%      1.65s  1.92%  image/jpeg.(*decoder).ensureNBits
     1.07s  1.25% 78.45%     22.21s 25.90%  crypto/internal/fips140/sha256.(*Digest).Write
     1.04s  1.21% 79.66%      1.04s  1.21%  image/jpeg.dctBox (inline)
     0.78s  0.91% 80.57%      0.78s  0.91%  aeshashbody
     0.71s  0.83% 81.40%     25.54s 29.78%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.68s  0.79% 82.19%     23.75s 27.70%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.65s  0.76% 82.95%      0.65s  0.76%  memeqbody
     0.51s  0.59% 83.55%      8.11s  9.46%  runtime.scanSpan
     0.43s   0.5% 84.05%      0.43s   0.5%  runtime.extractHeapBitsSmall
     0.42s  0.49% 84.54%      0.48s  0.56%  image/jpeg.(*decoder).readByteStuffedByte
     0.41s  0.48% 85.01%      1.16s  1.35%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.39s  0.45% 85.47%      0.45s  0.52%  runtime.(*spanInlineMarkBits).tryAcquire
     0.38s  0.44% 85.91%     27.59s 32.17%  example.com/isureal-company/isureal.imageSignature
     0.36s  0.42% 86.33%      0.55s  0.64%  example.com/isureal-company/isureal.postLess
     0.35s  0.41% 86.74%      0.96s  1.12%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.34s   0.4% 87.14%         2s  2.33%  runtime.mapaccess1_faststr
     0.29s  0.34% 87.48%     20.29s 23.66%  crypto/internal/fips140/sha256.block
     0.20s  0.23% 87.71%      1.08s  1.26%  runtime.mallocgcSmallScanNoHeader
     0.18s  0.21% 87.92%      0.51s  0.59%  database/sql.convertAssignRows
     0.15s  0.17% 88.09%      1.40s  1.63%  runtime.mapassign_faststr
     0.13s  0.15% 88.24%     10.91s 12.72%  runtime.mallocgc
     0.11s  0.13% 88.37%      0.53s  0.62%  crypto/internal/fips140/hmac.(*HMAC).Write
     0.11s  0.13% 88.50%      0.58s  0.68%  runtime.scanObjectSmall
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 75.1% / 100% | 60% | 0.0% | 9 | 176 MB | CPU が飽和 |
| isu2 (2 cpu) | 1.3% / 17% | 0% | 0.3% | 2 | 292 MB |  |
| isu3 (2 cpu) | 1.6% / 5% | 0% | 0.3% | 2 | 258 MB |  |
| isu4 (2 cpu) | 30.7% / 84% | 0% | 18.7% | 4 | 105 MB |  |

**isu1（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 136% / nginx 10% |
| isu2 | 目立つプロセス無し |
| isu3 | isureal 1% |
| isu4 | postgres 20% |

## Diff from previous run

スコア: 15668 → 21074 (+5406)

`8f720c2..214e79f`

```
 conf/isu1/home/isucon/env.sh       |   1 +
 webapp/go/cmd/image-export/main.go |  11 +--
 webapp/go/image_local.go           | 183 +++++++++++++++++++++++++++++--------
 3 files changed, 149 insertions(+), 46 deletions(-)
```
