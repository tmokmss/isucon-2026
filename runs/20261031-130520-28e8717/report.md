# run 20261031-130520-28e8717

## Summary

| | |
|---|---|
| score | 80118 |
| status | pass |
| ref | repl-replica |
| sha | 28e87177fb42 |
| 所要時間 | 263s |
| 前回比 | +42232 (111.5%) vs 20261031-125932-1dcbb5d |
| ベスト比 | +15418 (23.8%) vs 20261031-121235-e0e4cb3 = 64700 **ベスト更新** |
| note | 複製 第2段 local（フィード・投稿詳細・コメント/RealMoji 一覧を europe1・america1 が手元で返す）＋要約の照合 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	39	93	241	592	590
europe	7	15	37	85	85	84
america	6	15	33	67	69	68
```

```
{"id":"11207","status":"done","portalStatus":"finished","passed":true,"score":80118,"raw":80118,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T04:08:20.692Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T04:08:22.883Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T04:08:25.881Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T04:08:29.074Z level=INFO msg="asia 2 日目: 39 人"
time=2026-10-31T04:08:31.983Z level=INFO msg="europe 2 日目: 15 人"
time=2026-10-31T04:08:34.606Z level=INFO msg="user moved" user=leochi.5260 from=asia to=america
time=2026-10-31T04:08:34.984Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T04:08:38.042Z level=INFO msg="asia 3 日目: 93 人"
time=2026-10-31T04:08:41.669Z level=INFO msg="europe 3 日目: 37 人"
time=2026-10-31T04:08:42.112Z level=INFO msg="user moved" user=kieva-5550 from=asia to=america
time=2026-10-31T04:08:44.007Z level=INFO msg="america 3 日目: 33 人"
time=2026-10-31T04:08:47.053Z level=INFO msg="asia 4 日目: 241 人"
time=2026-10-31T04:08:52.003Z level=INFO msg="user moved" user=rinuto-6203 from=asia to=europe
time=2026-10-31T04:08:52.025Z level=INFO msg="europe 4 日目: 85 人"
time=2026-10-31T04:08:53.056Z level=INFO msg="america 4 日目: 67 人"
time=2026-10-31T04:08:53.731Z level=INFO msg="user moved" user=evaanka-6336 from=asia to=america
time=2026-10-31T04:08:56.116Z level=INFO msg="asia 5 日目: 592 人"
time=2026-10-31T04:08:56.151Z level=INFO msg="user moved" user=tesu_6546 from=asia to=america
time=2026-10-31T04:08:59.140Z level=INFO msg="europe 5 日目: 85 人"
time=2026-10-31T04:09:02.277Z level=INFO msg="america 5 日目: 69 人"
time=2026-10-31T04:09:12.422Z level=INFO msg="asia 最終日（グランプリ）: 590 人"
time=2026-10-31T04:09:12.923Z level=INFO msg="america 最終日（グランプリ）: 68 人"
time=2026-10-31T04:09:13.057Z level=INFO msg="europe 最終日（グランプリ）: 84 人"
time=2026-10-31T04:09:19.898Z level=INFO msg="グランプリ: アジア 438 人, ヨーロッパ 27 人, アメリカ 14 人"
{
  "pass": true,
  "score": 80118,
  "messages": [
    "vinyoko_5032: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "retsu.5125: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.9%	1007	1	0	GET	/api/notifications	8100.855	8.045	15.467	15.493	167
15.1%	1502	15	0	POST	/api/posts	3498.571	2.329	8.517	9.653	645
11.2%	3344	88	0	POST	/api/friend-requests	2605.818	0.779	4.165	7.922	113
8.1%	3860	1382	0	POST	/api/auth/login	1883.503	0.488	3.771	7.696	166
6.5%	8744	389	0	GET	^/api/images/[^/]+	1514.066	0.173	2.165	7.015	164792
5.3%	1250	6	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1232.016	0.986	5.237	7.420	116
2.5%	612	1	0	POST	/api/auth/register	578.325	0.945	3.932	5.451	227
2.3%	367	22	0	POST	^/api/posts/[^/]+/realmojis$	543.205	1.480	6.271	7.068	443
2.0%	468	118	0	PUT	/api/me/profile-image	471.571	1.008	4.042	4.655	5
1.6%	425	13	0	POST	^/api/posts/[^/]+/comments$	373.681	0.879	4.786	5.879	337
1.5%	2412	0	0	GET	/api/memories	343.774	0.143	0.907	1.902	9844
1.4%	429	0	0	PUT	^/api/follows/[^/]+$	330.252	0.770	3.037	6.343	0
1.3%	8275	0	0	GET	^/api/feeds/[^/]+	294.358	0.036	0.816	1.554	2329
1.0%	399	0	0	POST	/api/auth/logout	234.365	0.587	5.273	7.185	0
0.9%	276	0	0	PATCH	/api/me	218.592	0.792	4.162	4.484	275
0.8%	154	4	0	PATCH	^/api/posts/[^/]+$	188.265	1.222	6.561	7.325	633
0.5%	133	0	0	GET	/api/me/export	109.391	0.822	4.007	5.737	1829466
0.4%	59	0	0	DELETE	^/api/posts/[^/]+$	99.447	1.686	6.842	6.842	0
0.4%	2150	0	0	GET	/api/posting-window/current	91.583	0.043	0.310	1.324	237
0.3%	1612	0	0	GET	/api/users	77.639	0.048	0.999	1.325	240
0.2%	75	0	0	DELETE	^/api/comments/[^/]+$	55.263	0.737	5.190	5.190	0
0.2%	1352	44	0	GET	^/api/posts/me/current$	49.417	0.037	0.372	1.554	631
0.2%	48	0	0	DELETE	^/api/realmojis/[^/]+$	35.872	0.747	4.568	4.568	0
0.1%	1239	0	0	GET	^/api/relationships/[^/]+	34.076	0.028	0.372	1.554	973
0.1%	1224	10	0	GET	^/api/posts/[^/]+$	33.398	0.027	0.372	1.530	619
0.1%	351	0	0	GET	^/api/memories/[^/]+	33.236	0.095	0.993	1.242	715
0.1%	359	3	0	GET	/api/me	23.439	0.065	1.121	1.174	300
0.1%	20	0	0	PUT	^/api/blocks/[^/]+$	21.059	1.053	5.114	5.114	0
0.1%	1	0	0	POST	/api/initialize	19.851	19.851	19.851	19.851	90
0.1%	610	0	0	GET	^/api/users/[^/]+$	19.180	0.031	0.819	1.534	203
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.4%	221	1	0	GET	/api/notifications	1985.628	8.985	13.972	14.203	178
18.3%	274	6	0	POST	/api/posts	817.932	2.985	8.349	8.435	634
6.1%	1181	10	0	GET	^/api/images/[^/]+	272.533	0.231	1.415	3.075	191303
5.9%	408	6	0	POST	/api/friend-requests	265.480	0.651	3.656	4.870	114
5.0%	327	0	0	POST	/api/auth/login	224.390	0.686	3.580	5.251	246
3.6%	179	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	158.868	0.888	3.826	4.100	115
2.3%	370	0	0	GET	/api/memories	103.650	0.280	1.347	1.704	10412
2.3%	96	4	0	POST	^/api/posts/[^/]+/realmojis$	103.560	1.079	6.982	6.982	465
2.3%	369	0	0	GET	/api/posting-window/current	101.090	0.274	1.271	1.362	238
1.2%	26	0	0	GET	/api/me/export	51.933	1.997	7.109	7.109	2707086
1.1%	53	0	0	POST	/api/auth/logout	47.055	0.888	4.586	4.586	0
0.9%	216	8	0	GET	^/api/posts/me/current$	39.818	0.184	0.393	1.028	624
0.8%	197	0	0	GET	/api/users	33.836	0.172	0.471	1.254	237
0.7%	45	1	0	POST	^/api/posts/[^/]+/comments$	32.831	0.730	3.694	3.694	287
0.7%	69	0	0	PUT	^/api/follows/[^/]+$	30.892	0.448	3.458	3.458	0
0.6%	167	0	0	GET	^/api/relationships/[^/]+	27.300	0.163	1.206	1.309	964
0.5%	82	0	0	POST	/api/auth/register	24.035	0.293	0.474	0.474	228
0.5%	18	0	0	DELETE	^/api/realmojis/[^/]+$	21.698	1.205	3.914	3.914	0
0.5%	45	0	0	PUT	/api/me/profile-image	20.574	0.457	1.669	1.669	0
0.4%	1189	0	0	GET	^/api/feeds/[^/]+	18.135	0.015	0.234	1.049	2289
0.4%	36	0	0	PATCH	/api/me	17.292	0.480	3.409	3.409	282
0.4%	24	1	0	PATCH	^/api/posts/[^/]+$	16.855	0.702	3.034	3.034	627
0.3%	86	0	0	GET	^/api/users/[^/]+$	14.964	0.174	1.214	1.214	184
0.2%	10	0	0	DELETE	^/api/posts/[^/]+$	9.764	0.976	4.004	4.004	0
0.2%	34	0	0	GET	^/api/memories/[^/]+	7.396	0.218	1.238	1.238	389
0.2%	50	2	0	GET	/api/me	7.311	0.146	0.379	0.379	288
0.1%	6	0	0	PUT	^/api/blocks/[^/]+$	4.189	0.698	1.409	1.409	0
0.1%	1	0	0	DELETE	^/api/friendships/[^/]+$	3.949	3.949	3.949	3.949	0
0.1%	7	0	0	DELETE	^/api/comments/[^/]+$	3.151	0.450	1.059	1.059	0
0.0%	1	0	0	POST	/api/me/account-deletion	1.765	1.765	1.765	1.765	0
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.9%	190	1	0	GET	/api/notifications	1564.538	8.234	10.893	10.894	179
18.6%	228	6	0	POST	/api/posts	695.961	3.052	8.196	8.383	627
8.2%	393	12	0	POST	/api/friend-requests	306.666	0.780	3.923	5.175	112
6.4%	995	11	0	GET	^/api/images/[^/]+	237.251	0.238	1.154	2.040	205305
5.0%	259	0	0	POST	/api/auth/login	187.047	0.722	3.849	5.615	249
4.2%	146	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	156.440	1.072	4.855	5.532	116
2.1%	296	0	0	GET	/api/memories	78.229	0.264	1.359	1.445	12401
1.9%	301	0	0	GET	/api/posting-window/current	71.029	0.236	0.615	1.124	240
1.3%	49	0	0	POST	/api/auth/logout	47.155	0.962	4.784	4.784	0
1.2%	39	0	0	POST	^/api/posts/[^/]+/comments$	43.512	1.116	4.456	4.456	345
1.0%	190	0	0	GET	/api/users	38.309	0.202	0.858	1.403	238
1.0%	180	18	0	GET	^/api/posts/me/current$	36.791	0.204	1.200	1.224	582
0.8%	49	0	0	PUT	^/api/follows/[^/]+$	30.546	0.623	3.488	3.488	0
0.7%	11	0	0	GET	/api/me/export	26.310	2.392	3.890	3.890	3036234
0.7%	125	0	0	GET	^/api/relationships/[^/]+	25.267	0.202	1.357	1.358	1200
0.6%	13	3	0	POST	^/api/posts/[^/]+/realmojis$	23.976	1.844	4.950	4.950	398
0.5%	886	0	0	GET	^/api/feeds/[^/]+	19.919	0.022	0.375	1.024	2799
0.5%	64	0	0	POST	/api/auth/register	19.516	0.305	1.524	1.524	230
0.5%	31	0	0	PUT	/api/me/profile-image	19.029	0.614	3.366	3.366	0
0.5%	69	0	0	GET	^/api/users/[^/]+$	18.594	0.269	1.436	1.436	188
0.5%	23	5	0	PATCH	^/api/posts/[^/]+$	16.830	0.732	2.916	2.916	504
0.4%	10	0	0	DELETE	^/api/posts/[^/]+$	16.720	1.672	3.950	3.950	0
0.4%	31	0	0	PATCH	/api/me	14.158	0.457	2.228	2.228	277
0.3%	44	0	0	GET	^/api/memories/[^/]+	9.697	0.220	1.014	1.014	723
0.2%	6	0	0	DELETE	^/api/comments/[^/]+$	8.914	1.486	3.924	3.924	0
0.1%	36	1	0	GET	/api/me	5.104	0.142	0.399	0.399	290
0.1%	4	0	0	PUT	^/api/blocks/[^/]+$	4.716	1.179	2.867	2.867	0
0.1%	3	0	0	DELETE	^/api/follows/[^/]+$	4.100	1.367	2.703	2.703	0
0.1%	1	0	0	POST	/api/me/account-deletion	2.977	2.977	2.977	2.977	0
0.0%	50	0	0	GET	^/api/posts/[^/]+/comments$	1.613	0.032	0.329	0.329	1243
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 46429 クエリ / 82 種類 / 合計 387.684s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	57.6%	2854	223.416	78.28	433.40	665.1	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	16.3%	170	63.347	372.63	1567.89	6458.8	130.4	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
3	5.5%	3	21.432	7143.85	9535.79	30468.0	1199795.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
4	4.9%	1805	19.116	10.59	340.43	120.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	4.1%	3	15.848	5282.65	7075.90	22729.0	627468.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
6	3.9%	5387	14.983	2.78	55.79	11.3	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
7	2.0%	6	7.672	1278.59	6779.34	218216.5	2587.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
8	2.0%	3	7.657	2552.40	3451.31	12122.0	529765.3	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	1.3%	1981	5.162	2.61	37.61	25.7	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.5%	447	1.777	3.98	18.78	29.2	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
11	0.2%	495	0.950	1.92	18.76	22.3	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.2%	6	0.941	156.91	311.21	16775.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
13	0.2%	4026	0.829	0.21	19.10	20.3	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
14	0.2%	6	0.789	131.42	212.61	621.3	25256.2	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
15	0.2%	6	0.667	111.22	156.73	616.0	29856.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 3
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 3dfb77218fd567f1823fd487544a8b67a0c10f6a
Type: cpu
Time: 2026-10-31 13:07:55 JST
Duration: 60.04s, Total samples = 31.76s (52.90%)
Showing nodes accounting for 22.92s, 72.17% of 31.76s total
Dropped 950 nodes (cum <= 0.16s)
Showing top 40 nodes out of 300
      flat  flat%   sum%        cum   cum%
     2.27s  7.15%  7.15%      2.27s  7.15%  internal/runtime/syscall/linux.Syscall6
     1.79s  5.64% 12.78%      5.82s 18.32%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.62s  5.10% 17.88%      3.40s 10.71%  example.com/isureal-company/isureal.(*Store).digestLocked
     1.40s  4.41% 22.29%      1.40s  4.41%  runtime.memclrNoHeapPointers
     1.30s  4.09% 26.39%      1.62s  5.10%  runtime.findObject
     1.29s  4.06% 30.45%      1.29s  4.06%  runtime.memmove
     1.28s  4.03% 34.48%      1.28s  4.03%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.25s  3.94% 38.41%      4.83s 15.21%  runtime.scanObject
     1.21s  3.81% 42.22%      1.21s  3.81%  aeshashbody
     1.20s  3.78% 46.00%      1.21s  3.81%  example.com/isureal-company/isureal.h64 (inline)
     0.96s  3.02% 49.02%      1.49s  4.69%  runtime.tryDeferToSpanScan
     0.94s  2.96% 51.98%      0.94s  2.96%  crypto/internal/fips140/sha256.blockAVX2
     0.57s  1.79% 53.78%      0.57s  1.79%  memeqbody
     0.52s  1.64% 55.42%      0.52s  1.64%  example.com/isureal-company/isureal.(*UserRow).active
     0.47s  1.48% 56.90%      1.85s  5.82%  runtime.mapaccess1_faststr
     0.37s  1.16% 58.06%      0.62s  1.95%  example.com/isureal-company/isureal.postLess
     0.35s  1.10% 59.16%      0.35s  1.10%  runtime.spanClass.sizeclass (inline)
     0.32s  1.01% 60.17%      0.32s  1.01%  runtime.nextFreeFast (inline)
     0.28s  0.88% 61.05%      1.89s  5.95%  runtime.mapassign_faststr
     0.26s  0.82% 61.87%      0.49s  1.54%  runtime.greyobject
     0.23s  0.72% 62.59%      0.47s  1.48%  database/sql.convertAssignRows
     0.23s  0.72% 63.32%      0.65s  2.05%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     0.22s  0.69% 64.01%      0.22s  0.69%  sync/atomic.(*Int32).Add
     0.21s  0.66% 64.67%      0.21s  0.66%  indexbody
     0.18s  0.57% 65.24%      0.18s  0.57%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.18s  0.57% 65.81%      0.18s  0.57%  runtime.futex
     0.17s  0.54% 66.34%      0.20s  0.63%  example.com/isureal-company/isureal.(*Store).pruneAfterBoundary
     0.17s  0.54% 66.88%      0.28s  0.88%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     0.17s  0.54% 67.41%      0.86s  2.71%  runtime.mallocgcSmallScanNoHeader
     0.16s   0.5% 67.92%      0.16s   0.5%  example.com/isureal-company/isureal.rmLess
     0.16s   0.5% 68.42%      2.55s  8.03%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.16s   0.5% 68.92%      0.36s  1.13%  internal/runtime/maps.(*Iter).Next
     0.15s  0.47% 69.40%      0.45s  1.42%  runtime.mapaccess2_faststr
     0.15s  0.47% 69.87%      0.16s   0.5%  runtime.typePointers.next
     0.14s  0.44% 70.31%      0.17s  0.54%  runtime.spanOf (inline)
     0.13s  0.41% 70.72%      0.20s  0.63%  compress/flate.(*compressor).deflate
     0.13s  0.41% 71.13%      0.16s   0.5%  runtime.heapArenaOf (inline)
     0.13s  0.41% 71.54%      2.86s  9.01%  runtime.mallocgc
     0.12s  0.38% 71.91%      0.30s  0.94%  runtime.scanObjectsSmall
     0.08s  0.25% 72.17%      0.24s  0.76%  example.com/isureal-company/isureal.(*Store).rebuildIndexes.func4
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 42.7% / 100% | 21% | 0.3% | 6 | 177 MB | run queue が積もっている |
| isu2 (2 cpu) | 23.3% / 71% | 0% | 0.3% | 4 | 935 MB |  |
| isu3 (2 cpu) | 24.8% / 84% | 0% | 0.4% | 3 | 948 MB |  |
| isu4 (2 cpu) | 46.1% / 100% | 17% | 15.6% | 32 | 101 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 66% / nginx 12% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 45% |
| isu3 | isureal 48% |
| isu4 | postgres 54% / isureal 3% |

## Diff from previous run

スコア: 37886 → 80118 (+42232)

`1dcbb5d..28e8717`

```
 conf/isu2/home/isucon/env.sh  |  2 +-
 conf/isu3/home/isucon/env.sh  |  2 +-
 webapp/go/relationship_api.go |  2 +-
 webapp/go/repl.go             | 42 +++++++++++++++++++++++++++++++++---------
 webapp/go/repl_digest.go      | 30 ++++++++++++++++++++++--------
 5 files changed, 58 insertions(+), 20 deletions(-)
```
