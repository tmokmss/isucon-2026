# run 20261031-140122-bd0fe45

## Summary

| | |
|---|---|
| score | 126129 |
| status | pass |
| ref | bd0fe45 |
| sha | bd0fe45bf2e2 |
| 所要時間 | 542s |
| 前回比 | +15531 (14.0%) vs 20261031-135652-8b16848 |
| ベスト比 | +15531 (14.0%) vs 20261031-135652-8b16848 = 110598 **ベスト更新** |
| note | 複製 local・時刻の印で待つ（main c6abd65 取り込み済み）1回目 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	52	152	439	1012	1012
europe	7	22	67	187	194	194
america	6	16	52	135	139	138
```

```
{"id":"15942","status":"done","portalStatus":"finished","passed":true,"score":126129,"raw":126129,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T05:08:57.980Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T05:09:00.947Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T05:09:03.944Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T05:09:06.978Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T05:09:08.615Z level=INFO msg="user moved" user=maivy_5193 from=asia to=europe
time=2026-10-31T05:09:09.943Z level=INFO msg="europe 2 日目: 22 人"
time=2026-10-31T05:09:10.298Z level=INFO msg="user moved" user=toryulia5348 from=asia to=europe
time=2026-10-31T05:09:11.028Z level=INFO msg="user moved" user=keevato5369 from=asia to=europe
time=2026-10-31T05:09:12.943Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T05:09:14.886Z level=INFO msg="user moved" user=liakaben5489 from=europe to=america
time=2026-10-31T05:09:15.975Z level=INFO msg="asia 3 日目: 152 人"
time=2026-10-31T05:09:16.624Z level=INFO msg="user moved" user=miano_5591 from=asia to=europe
time=2026-10-31T05:09:17.660Z level=INFO msg="user moved" user=miano_5591 from=europe to=asia
time=2026-10-31T05:09:18.279Z level=INFO msg="user moved" user=nehamia.5905 from=asia to=europe
time=2026-10-31T05:09:18.950Z level=INFO msg="europe 3 日目: 67 人"
time=2026-10-31T05:09:19.280Z level=INFO msg="user moved" user=hemuta6145 from=asia to=europe
time=2026-10-31T05:09:19.409Z level=INFO msg="user moved" user=motaso-6195 from=asia to=america
time=2026-10-31T05:09:19.657Z level=INFO msg="user moved" user=marmaxne_6290 from=asia to=america
time=2026-10-31T05:09:20.318Z level=INFO msg="user moved" user=motaso-6195 from=america to=asia
time=2026-10-31T05:09:20.571Z level=INFO msg="user moved" user=marmaxne_6290 from=america to=europe
time=2026-10-31T05:09:21.944Z level=INFO msg="america 3 日目: 52 人"
time=2026-10-31T05:09:22.166Z level=INFO msg="user moved" user=toryulia5348 from=europe to=america
time=2026-10-31T05:09:24.448Z level=INFO msg="user moved" user=liakaben5489 from=america to=europe
time=2026-10-31T05:09:24.989Z level=INFO msg="asia 4 日目: 439 人"
time=2026-10-31T05:09:25.682Z level=INFO msg="user moved" user=tsunara6664 from=america to=asia
time=2026-10-31T05:09:26.769Z level=INFO msg="user moved" user=motaso-6195 from=asia to=america
time=2026-10-31T05:09:27.355Z level=INFO msg="user moved" user=miano_5591 from=asia to=europe
time=2026-10-31T05:09:27.711Z level=INFO msg="user moved" user=maivy_5193 from=europe to=america
time=2026-10-31T05:09:27.958Z level=INFO msg="europe 4 日目: 187 人"
time=2026-10-31T05:09:28.010Z level=INFO msg="user moved" user=rase_6984 from=asia to=america
time=2026-10-31T05:09:28.058Z level=INFO msg="user moved" user=tsunara6664 from=asia to=europe
time=2026-10-31T05:09:28.125Z level=INFO msg="user moved" user=nedanshi5531 from=america to=asia
time=2026-10-31T05:09:30.729Z level=INFO msg="user moved" user=neul_8391 from=asia to=europe
time=2026-10-31T05:09:30.777Z level=INFO msg="user moved" user=nedanshi5531 from=asia to=america
time=2026-10-31T05:09:30.823Z level=INFO msg="user moved" user=fufute.7571 from=asia to=america
time=2026-10-31T05:09:30.849Z level=INFO msg="user moved" user=rase_6984 from=america to=asia
time=2026-10-31T05:09:30.946Z level=INFO msg="america 4 日目: 135 人"
time=2026-10-31T05:09:31.094Z level=INFO msg="user moved" user=maxsu-8588 from=asia to=europe
time=2026-10-31T05:09:31.209Z level=INFO msg="user moved" user=hemuta6145 from=europe to=america
time=2026-10-31T05:09:31.293Z level=INFO msg="user moved" user=torka_7558 from=asia to=europe
time=2026-10-31T05:09:31.889Z level=INFO msg="user moved" user=rufu-7334 from=asia to=america
time=2026-10-31T05:09:32.531Z level=INFO msg="user moved" user=nehamia.5905 from=europe to=asia
time=2026-10-31T05:09:33.207Z level=INFO msg="user moved" user=kata-8776 from=europe to=america
time=2026-10-31T05:09:34.133Z level=INFO msg="user moved" user=liakaben5489 from=europe to=america
time=2026-10-31T05:09:34.229Z level=INFO msg="asia 5 日目: 1012 人"
time=2026-10-31T05:09:37.028Z level=INFO msg="europe 5 日目: 194 人"
time=2026-10-31T05:09:40.302Z level=INFO msg="america 5 日目: 139 人"
time=2026-10-31T05:09:48.973Z level=INFO msg="america 最終日（グランプリ）: 138 人"
time=2026-10-31T05:09:49.089Z level=INFO msg="asia 最終日（グランプリ）: 1012 人"
time=2026-10-31T05:09:49.180Z level=INFO msg="europe 最終日（グランプリ）: 194 人"
time=2026-10-31T05:09:57.984Z level=INFO msg="グランプリ: アジア 344 人, ヨーロッパ 109 人, アメリカ 68 人"
{
  "pass": true,
  "score": 126129,
  "messages": [
    "yumar6256: 友達申請を承認できません: POST /api/friend-requests/e7a4b4e6-f26f-4f9a-a3c3-4661e56a65ae/accept: 応答がタイムアウトしました",
    "tona5085: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "maxtako-8768: 投稿 fc2e8fe0-0f3b-4f0e-90a0-4e986fbb281d に RealMoji を送れません: POST /api/posts/fc2e8fe0-0f3b-4f0e-90a0-4e986fbb281d/realmojis: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
31.0%	1572	0	0	GET	/api/notifications	12550.725	7.984	14.899	14.905	165
18.4%	2302	81	0	POST	/api/posts	7456.768	3.239	10.339	12.369	631
13.1%	5315	63	0	POST	/api/friend-requests	5298.739	0.997	7.399	14.457	114
8.9%	13464	665	0	GET	^/api/images/[^/]+	3595.057	0.267	4.792	14.231	141556
6.8%	6131	1899	0	POST	/api/auth/login	2736.022	0.446	5.863	14.172	177
5.1%	1806	8	0	POST	^/api/friend-requests/[^/]+/[^/]+$	2064.154	1.143	7.086	10.975	116
3.2%	716	122	0	PUT	/api/me/profile-image	1281.382	1.790	8.136	14.457	3
2.8%	1040	1	0	POST	/api/auth/register	1125.351	1.082	7.022	10.759	227
2.0%	496	48	0	POST	^/api/posts/[^/]+/realmojis$	823.326	1.660	9.271	17.294	442
1.7%	716	0	0	PUT	^/api/follows/[^/]+$	678.576	0.948	4.854	6.556	0
1.4%	556	44	0	POST	^/api/posts/[^/]+/comments$	555.669	0.999	7.908	13.986	326
1.3%	515	0	0	PATCH	/api/me	532.458	1.034	6.909	8.993	281
0.9%	3309	0	0	GET	/api/memories	358.715	0.108	0.699	2.830	11057
0.7%	261	0	0	GET	/api/me/export	270.563	1.037	7.507	10.391	1818742
0.6%	650	0	0	POST	/api/auth/logout	254.676	0.392	5.314	10.387	0
0.6%	193	15	0	PATCH	^/api/posts/[^/]+$	224.051	1.161	8.038	9.153	607
0.3%	11498	0	0	GET	^/api/feeds/[^/]+	131.460	0.011	0.146	0.604	1666
0.3%	78	0	0	DELETE	^/api/posts/[^/]+$	119.025	1.526	6.711	6.711	0
0.2%	67	0	0	DELETE	^/api/realmojis/[^/]+$	72.899	1.088	9.291	9.291	0
0.2%	92	0	0	DELETE	^/api/comments/[^/]+$	69.608	0.757	7.002	7.002	0
0.1%	3344	0	0	GET	/api/posting-window/current	59.868	0.018	0.177	0.441	237
0.1%	1778	92	0	GET	^/api/posts/me/current$	34.660	0.019	0.242	0.495	620
0.1%	621	0	0	GET	^/api/memories/[^/]+	33.148	0.053	0.621	1.077	619
0.1%	34	0	0	PUT	^/api/blocks/[^/]+$	27.577	0.811	4.613	4.613	0
0.1%	2583	0	0	GET	/api/users	25.512	0.010	0.086	2.426	244
0.0%	19	0	0	DELETE	^/api/follows/[^/]+$	19.641	1.034	3.989	3.989	0
0.0%	1541	0	0	GET	^/api/relationships/[^/]+	19.577	0.013	0.151	0.459	902
0.0%	1	0	0	POST	/api/initialize	15.231	15.231	15.231	15.231	90
0.0%	13	0	0	PUT	^/api/admin/users/[^/]+/certification$	14.217	1.094	7.084	7.084	0
0.0%	1470	7	0	GET	^/api/posts/[^/]+$	12.639	0.009	0.107	0.401	642
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
38.9%	450	1	0	GET	/api/notifications	3713.484	8.252	11.822	12.498	175
15.0%	570	7	0	POST	/api/posts	1426.954	2.503	8.567	10.387	648
10.3%	1025	8	0	POST	/api/friend-requests	985.283	0.961	5.911	9.822	114
6.0%	733	0	0	POST	/api/auth/login	574.059	0.783	6.543	9.310	247
5.3%	2667	16	0	GET	^/api/images/[^/]+	501.585	0.188	1.042	1.482	185045
5.2%	398	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	497.037	1.249	8.632	15.422	116
3.6%	2685	0	0	GET	^/api/feeds/[^/]+	338.697	0.126	0.879	1.162	2279
2.0%	789	0	0	GET	/api/posting-window/current	187.075	0.237	1.029	1.179	239
1.6%	82	12	0	POST	^/api/posts/[^/]+/realmojis$	151.337	1.846	10.250	10.250	397
1.2%	119	4	0	POST	^/api/posts/[^/]+/comments$	119.118	1.001	7.060	7.880	327
1.1%	767	0	0	GET	/api/memories	105.033	0.137	1.422	2.305	11775
1.0%	187	0	0	POST	/api/auth/register	99.732	0.533	5.130	6.565	228
0.9%	138	0	0	PUT	^/api/follows/[^/]+$	89.205	0.646	4.233	5.753	0
0.9%	113	0	0	PUT	/api/me/profile-image	89.088	0.788	3.536	4.089	0
0.9%	120	0	0	POST	/api/auth/logout	85.962	0.716	4.559	11.796	0
0.8%	492	0	0	GET	/api/users	80.339	0.163	1.035	1.152	239
0.7%	57	4	0	PATCH	^/api/posts/[^/]+$	70.546	1.238	5.240	5.240	614
0.7%	498	24	0	GET	^/api/posts/me/current$	68.748	0.138	0.944	1.164	625
0.7%	106	0	0	PATCH	/api/me	65.846	0.621	4.265	4.593	286
0.5%	381	0	0	GET	^/api/relationships/[^/]+	49.708	0.130	1.051	1.157	880
0.5%	384	11	0	GET	^/api/posts/[^/]+$	47.173	0.123	0.930	1.155	618
0.4%	43	0	0	GET	/api/me/export	39.415	0.917	6.849	6.849	2022820
0.4%	22	0	0	DELETE	^/api/posts/[^/]+$	36.138	1.643	5.376	5.376	0
0.2%	182	0	0	GET	^/api/posts/[^/]+/realmojis$	20.388	0.112	0.948	0.990	1130
0.2%	176	0	0	GET	^/api/users/[^/]+$	18.928	0.108	0.628	0.700	209
0.2%	135	0	0	GET	^/api/posts/[^/]+/comments$	18.741	0.139	1.095	1.164	2097
0.2%	25	0	0	DELETE	^/api/comments/[^/]+$	17.179	0.687	3.312	3.312	0
0.1%	88	0	0	GET	^/api/memories/[^/]+	13.730	0.156	1.582	1.582	1144
0.1%	119	2	0	GET	/api/me	12.640	0.106	0.332	0.342	297
0.1%	6	0	0	PUT	^/api/blocks/[^/]+$	5.070	0.845	2.729	2.729	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.3%	305	1	0	GET	/api/notifications	2171.448	7.120	9.356	9.530	173
14.1%	392	13	0	POST	/api/posts	892.666	2.277	8.894	9.767	634
14.1%	660	9	0	POST	/api/friend-requests	889.154	1.347	7.298	10.518	114
7.0%	521	0	0	POST	/api/auth/login	441.221	0.847	6.391	10.864	249
6.7%	1810	15	0	GET	^/api/images/[^/]+	422.749	0.234	1.163	2.405	167925
3.3%	1291	0	0	GET	^/api/feeds/[^/]+	211.406	0.164	0.954	1.288	1678
3.2%	148	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	203.783	1.377	6.544	9.320	116
2.3%	75	9	0	POST	^/api/posts/[^/]+/realmojis$	144.302	1.924	9.619	9.619	436
2.2%	546	0	0	GET	/api/posting-window/current	141.929	0.260	1.130	1.325	240
1.5%	512	0	0	GET	/api/memories	96.036	0.188	1.564	2.013	16070
1.3%	96	0	0	PUT	^/api/follows/[^/]+$	81.680	0.851	6.059	6.059	0
1.2%	73	0	0	PUT	/api/me/profile-image	73.306	1.004	8.846	8.846	0
1.0%	46	5	0	POST	^/api/posts/[^/]+/comments$	65.623	1.427	7.749	7.749	296
1.0%	311	0	0	GET	/api/users	62.110	0.200	1.047	1.285	240
0.9%	337	29	0	GET	^/api/posts/me/current$	57.776	0.171	1.043	1.417	599
0.9%	68	0	0	POST	/api/auth/logout	54.997	0.809	6.704	6.704	0
0.8%	131	0	0	POST	/api/auth/register	53.629	0.409	2.731	2.792	230
0.7%	61	0	0	PATCH	/api/me	41.782	0.685	3.626	3.626	285
0.6%	34	1	0	PATCH	^/api/posts/[^/]+$	41.035	1.207	4.438	4.438	643
0.5%	20	0	0	DELETE	^/api/posts/[^/]+$	31.517	1.576	6.036	6.036	0
0.4%	23	0	0	GET	/api/me/export	26.599	1.156	5.806	5.806	2932388
0.4%	188	0	0	GET	^/api/relationships/[^/]+	25.486	0.136	0.896	0.936	714
0.3%	145	5	0	GET	^/api/posts/[^/]+$	20.749	0.143	0.839	0.907	613
0.2%	86	0	0	GET	^/api/users/[^/]+$	13.811	0.161	1.280	1.280	217
0.2%	75	0	0	GET	^/api/memories/[^/]+	13.299	0.177	1.545	1.545	1133
0.2%	77	1	0	GET	/api/me	12.430	0.161	1.214	1.214	300
0.2%	73	1	0	GET	^/api/posts/[^/]+/realmojis$	10.831	0.148	1.007	1.007	1749
0.1%	55	1	0	GET	^/api/posts/[^/]+/comments$	7.614	0.138	0.969	0.969	1563
0.1%	7	0	0	DELETE	^/api/realmojis/[^/]+$	5.806	0.829	3.930	3.930	0
0.1%	6	0	0	DELETE	^/api/comments/[^/]+$	5.743	0.957	3.356	3.356	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 76509 クエリ / 87 種類 / 合計 795.425s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	67.4%	4530	535.864	118.29	885.51	694.8	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	12.5%	277	99.800	360.29	1913.31	1182.6	120.6	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
3	6.6%	3796	52.562	13.85	862.44	118.3	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	3.2%	8783	25.291	2.88	45.36	10.9	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	2.7%	3	21.529	7176.40	9554.24	31242.0	1199981.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	2.0%	3	16.298	5432.79	7236.99	24949.0	628394.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	1.6%	3165	12.815	4.05	179.64	25.0	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
8	1.0%	3	7.938	2646.00	3647.52	12889.0	529896.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	0.4%	584	3.547	6.07	136.07	29.1	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.4%	6903	3.379	0.49	161.25	19.9	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.3%	6	2.725	454.09	1356.20	27823.2	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.3%	6	2.291	381.88	1223.89	219930.3	2607.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
13	0.2%	5486	1.903	0.35	150.85	10.9	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
14	0.2%	668	1.567	2.35	113.62	21.4	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
15	0.2%	2273	1.478	0.65	146.62	21.0	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)

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
Build ID: b0705343cf26e9b0d9a736a4928287957110e813
Type: cpu
Time: 2026-10-31 14:08:35 JST
Duration: 60.10s, Total samples = 38.24s (63.63%)
Showing nodes accounting for 26.13s, 68.33% of 38.24s total
Dropped 1032 nodes (cum <= 0.19s)
Showing top 40 nodes out of 288
      flat  flat%   sum%        cum   cum%
     4.24s 11.09% 11.09%      4.24s 11.09%  internal/runtime/syscall/linux.Syscall6
     4.03s 10.54% 21.63%      4.03s 10.54%  crypto/internal/fips140/sha256.blockAVX2
     1.65s  4.31% 25.94%      1.65s  4.31%  runtime.memmove
     1.59s  4.16% 30.10%      1.59s  4.16%  runtime.memclrNoHeapPointers
     1.53s  4.00% 34.10%      5.99s 15.66%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.46s  3.82% 37.92%      5.66s 14.80%  runtime.scanObject
     1.38s  3.61% 41.53%      1.38s  3.61%  aeshashbody
     1.37s  3.58% 45.11%      1.37s  3.58%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.37s  3.58% 48.69%      1.75s  4.58%  runtime.findObject
     1.18s  3.09% 51.78%      2.09s  5.47%  runtime.tryDeferToSpanScan
     0.69s  1.80% 53.58%      0.69s  1.80%  runtime.spanClass.sizeclass (inline)
     0.67s  1.75% 55.33%      2.22s  5.81%  runtime.mapaccess1_faststr
     0.42s  1.10% 56.43%      0.42s  1.10%  runtime.nextFreeFast (inline)
     0.39s  1.02% 57.45%      0.39s  1.02%  runtime.futex
     0.34s  0.89% 58.34%      0.34s  0.89%  memeqbody
     0.32s  0.84% 59.18%      0.52s  1.36%  example.com/isureal-company/isureal.postLess
     0.29s  0.76% 59.94%      1.85s  4.84%  runtime.mapassign_faststr
     0.25s  0.65% 60.59%      0.60s  1.57%  database/sql.convertAssignRows
     0.21s  0.55% 61.14%      3.90s 10.20%  runtime.mallocgc
     0.20s  0.52% 61.66%      0.20s  0.52%  indexbody
     0.20s  0.52% 62.19%      0.24s  0.63%  runtime.typePointers.next
     0.19s   0.5% 62.68%      0.24s  0.63%  runtime.(*mspan).markBitsForIndex (inline)
     0.16s  0.42% 63.10%      3.55s  9.28%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.15s  0.39% 63.49%      1.01s  2.64%  runtime.mallocgcSmallScanNoHeader
     0.14s  0.37% 63.86%      4.36s 11.40%  crypto/internal/fips140/sha256.(*Digest).Write
     0.14s  0.37% 64.23%      0.26s  0.68%  internal/runtime/maps.(*Iter).Next
     0.14s  0.37% 64.59%      0.36s  0.94%  runtime.greyobject
     0.13s  0.34% 64.93%      2.84s  7.43%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.13s  0.34% 65.27%      1.27s  3.32%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.31% 65.59%      4.11s 10.75%  example.com/isureal-company/isureal.computeImageSignature
     0.12s  0.31% 65.90%      0.42s  1.10%  runtime.mapaccess2_faststr
     0.11s  0.29% 66.19%      0.28s  0.73%  compress/flate.(*compressor).deflate
     0.11s  0.29% 66.47%      3.92s 10.25%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.11s  0.29% 66.76%      0.78s  2.04%  runtime.mallocgcSmallNoscan
     0.11s  0.29% 67.05%      0.33s  0.86%  runtime.scanObjectsSmall
     0.10s  0.26% 67.31%      3.72s  9.73%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.10s  0.26% 67.57%      1.89s  4.94%  example.com/isureal-company/isureal.listMemories
     0.10s  0.26% 67.83%      0.72s  1.88%  github.com/jackc/pgx/v5/pgtype.scanPlanString.Scan
     0.10s  0.26% 68.10%      0.51s  1.33%  runtime.wbBufFlush1
     0.09s  0.24% 68.33%      3.96s 10.36%  example.com/isureal-company/isureal.loadBig
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 54.1% / 100% | 26% | 0.6% | 6 | 228 MB | run queue が積もっている |
| isu2 (2 cpu) | 30.2% / 98% | 3% | 1.1% | 5 | 201 MB | run queue が積もっている |
| isu3 (2 cpu) | 28.3% / 100% | 4% | 1.1% | 5 | 113 MB | run queue が積もっている |
| isu4 (2 cpu) | 57.8% / 100% | 34% | 14.4% | 39 | 103 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 79% / nginx 12% |
| isu2 | isureal 57% / nginx 1% |
| isu3 | isureal 53% / nginx 1% |
| isu4 | postgres 77% / kcompactd0 2% / isureal 2% / kswapd0 1% |

## Diff from previous run

スコア: 110598 → 126129 (+15531)

`8b16848..bd0fe45`

```
 conf/isu1/home/isucon/env.sh | 1 -
 conf/isu2/home/isucon/env.sh | 2 +-
 conf/isu3/home/isucon/env.sh | 2 +-
 3 files changed, 2 insertions(+), 3 deletions(-)
```
