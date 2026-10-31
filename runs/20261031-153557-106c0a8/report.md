# run 20261031-153557-106c0a8

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | presign-next-minute |
| sha | 106c0a8b7664 |
| 所要時間 | 270s |
| 前回比 | -186885 (-100.0%) vs 20261031-153030-dca9471 |
| ベスト比 | -191816 (-100.0%) vs 20261031-151339-5a230a4 = 191816 |
| note | sig cache generations fixed-size keys, presign next minute, JPEG prefers asia2 |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/agent-aff106d587106aae5/runs/20261031-153557-106c0a8/bench.log
```

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目
asia	18	50	145	423
europe	7	19	55	—
america	6	21	62	—

最終日（1 人 20 点）が来なかったリージョン: asia, europe, america
```

```
{"id":"23713","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"負荷走行を打ち切りました"}

--- stdout ---
time=2026-10-31T06:39:39.129Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:39:42.125Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:39:45.124Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:39:48.134Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T06:39:49.242Z level=INFO msg="user moved" user=niru.5156 from=asia to=america
time=2026-10-31T06:39:51.123Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T06:39:54.123Z level=INFO msg="america 2 日目: 21 人"
time=2026-10-31T06:39:57.134Z level=INFO msg="asia 3 日目: 145 人"
time=2026-10-31T06:39:59.693Z level=INFO msg="user moved" user=chiwa6020 from=asia to=america
time=2026-10-31T06:40:00.212Z level=INFO msg="europe 3 日目: 55 人"
time=2026-10-31T06:40:00.376Z level=INFO msg="user moved" user=chiwa6020 from=america to=asia
time=2026-10-31T06:40:03.125Z level=INFO msg="america 3 日目: 62 人"
time=2026-10-31T06:40:03.855Z level=INFO msg="user moved" user=dankator-6369 from=europe to=asia
time=2026-10-31T06:40:06.166Z level=INFO msg="asia 4 日目: 423 人"
time=2026-10-31T06:40:07.937Z level=INFO msg="user moved" user=chiwa6020 from=asia to=america
time=2026-10-31T06:40:24.138Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": false,
  "score": 0,
  "messages": [
    "負荷走行を打ち切りました: 40s〜45s のエラー率 100.0%（323/323 件）が上限 10.0% を超えました",
    "renoa-5727: 投稿 d0fd33e5-a765-4cae-8698-47d70c8d160f の画像を取得できません: GET /api/images/ff7bd8b6-7e8f-4e6c-8aa0-5ce844a9d065?exp=1793515260\u0026sig=9MKTmTm138LkVjG1Y-BJT0Li9ECYsk6w__ZkUIXqSLU: ステータスコード 500 が返りました（期待値 200、context deadline exceeded）",
    "danyosol-6784: 投稿 8f0b28ac-a0ee-47e2-8569-38dab0899000 の画像を取得できません: GET /api/images/9ef56661-31d3-4a74-9817-91cfa3f43149?exp=1793515260\u0026sig=TbvAAZH21J_YfQMg6XklAuLiRfuq1SjoWw_0OFXFmog: ステータスコード 500 が返りました（期待値 200、Get \"http://192.168.3.10:8081/_internal/image?id=9ef56661-31d3-4a74-9817-91cfa3f43149\": context deadline exceeded）",
    "chibeneva.6847: 投稿 759a55d9-34b2-4423-b814-8bb445b6bdf6 の画像を取得できません: GET /api/images/2f6fad34-cf8e-43ac-b13d-745043f29f54?exp=1793515260\u0026sig=KC2YgkGWh7_yuAjMXH2kPqM56bKkxLnK5mlvWoO5cMo: ステータスコード 500 が返りました（期待値 200、context deadline exceeded）",
    "tsumax5641: 投稿 d0fd33e5-a765-4cae-8698-47d70c8d160f の画像を取得できません: GET /api/images/ff7bd8b6-7e8f-4e6c-8aa0-5ce844a9d065?exp=1793515260\u0026sig=9MKTmTm138LkVjG1Y-BJT0Li9ECYsk6w__ZkUIXqSLU: ステータスコード 500 が返りました（期待値 200、context deadline exceeded）",
    "europe の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました",
    "manuela_almeida_2494: 履歴を読めません: GET /api/images/4aa3bb4c-2b40-4644-ad36-3071bed34093?exp=1793515260\u0026sig=2EfkuTlIVYRG7PX94dV0qflLOeRY2-I1G7KFuR5lWJs: ステータスコード 500 が返りました（期待値 200、Get \"http://192.168.1.10:8081/_internal/image-check?id=4aa3bb4c-2b40-4644-ad36-3071bed34093\": context deadline exceeded (Client.Timeout exceeded while awaiting headers)）",
    "diego_gonzalez_4290: 履歴を読めません: GET /api/images/9f1bf4cb-ec6c-472f-926d-be2c0d6a90cb?exp=1793515260\u0026sig=bmzsWnIlwLFVuRvm1LMVP7Ip8ilr6Yz5jaHtZNRQoiI: ステータスコード 500 が返りました（期待値 200、Get \"http://192.168.1.10:8081/_internal/image-check?id=9f1bf4cb-ec6c-472f-926d-be2c0d6a90cb\": context deadline exceeded (Client.Timeout exceeded while awaiting headers)）",
    "capybaras_en_promenade_4662: 履歴を読めません: GET /api/images/161090aa-3402-48e8-9526-92796080b370?exp=1793515260\u0026sig=UL2Za_hREKHFkW5frwC24nRqFOQUCKGoYi5nsCxTvbs: ステータスコード 500 が返りました（期待値 200、Get \"http://192.168.1.10:8081/_internal/image-check?id=161090aa-3402-48e8-9526-92796080b370\": context deadline exceeded (Client.Timeout exceeded while awaiting headers)）",
    "charcoal_mini_turkey_2371: 履歴を読めません: GET /api/images/361e649c-2c88-4b9d-93ec-bd017ae405be?exp=1793515260\u0026sig=dn6bLOCaNw11o1MNemzr-6LNZvkmvHVOjsfk8Tq74tQ: ステータスコード 500 が返りました（期待値 200、Get \"http://192.168.1.10:8081/_internal/image-check?id=361e649c-2c88-4b9d-93ec-bd017ae405be\": context deadline exceeded (Client.Timeout exceeded while awaiting headers)）",
    "america の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました",
    "sushi_on_a_sunflower_4749: 履歴を読めません: GET /api/images/af425b18-8c6c-4cb2-b3c4-3464a10f8e0f?exp=1793515260\u0026sig=0mcbwvKjsmhgKhOJtbtUOjKRmBoIJnJwGAmrzcPq4ds: ステータスコード 500 が返りました（期待値 200、Get \"http://192.168.1.10:8081/_internal/image-check?id=af425b18-8c6c-4cb2-b3c4-3464a10f8e0f\": context deadline exceeded (Client.Timeout exceeded while awaiting headers)）",
    "miyochi-6978: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "seoyeon_lee_20: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "kifu-5667: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "solhozoe.6828: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "mone.6742: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "hahator-5827: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "murito_6578: 投稿 759a55d9-34b2-4423-b814-8bb445b6bdf6 に RealMoji を送れません: POST /api/posts/759a55d9-34b2-4423-b814-8bb445b6bdf6/realmojis: 応答がタイムアウトしました",
    "yunoana.6106: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "leokune-6898: 投稿できません: POST /api/posts: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
22.3%	8847	186	852	GET	^/api/images/[^/]+	5657.861	0.640	15.853	18.954	36429
21.9%	431	1	2	GET	/api/notifications	5561.232	12.903	20.144	20.144	95
18.6%	3742	294	626	POST	/api/auth/login	4708.330	1.258	17.078	17.884	215
15.5%	402	47	82	POST	/api/posts	3940.468	9.802	64.014	64.683	478
8.7%	434	28	99	POST	/api/auth/register	2203.671	5.078	18.119	18.316	198
3.1%	206	50	43	PUT	/api/me/profile-image	790.689	3.838	19.449	19.453	40
2.5%	1732	0	23	GET	/api/memories	641.451	0.370	14.799	18.381	16508
1.5%	578	0	20	GET	/api/posting-window/current	388.551	0.672	17.545	18.038	235
1.1%	2254	0	13	GET	^/api/feeds/[^/]+	285.416	0.127	3.583	18.257	1669
0.8%	1005	30	11	POST	/api/friend-requests	200.663	0.200	12.178	18.203	113
0.7%	253	1	11	GET	^/api/posts/me/current$	186.706	0.738	17.257	17.773	629
0.6%	199	0	5	GET	/api/me/export	161.382	0.811	19.538	19.635	1752539
0.5%	497	0	7	GET	/api/users	131.409	0.264	17.178	17.513	249
0.5%	96	0	7	PATCH	/api/me	120.049	1.251	17.469	17.469	261
0.3%	505	0	7	POST	/api/auth/logout	73.795	0.146	1.167	13.975	2
0.3%	126	0	4	PUT	^/api/follows/[^/]+$	72.399	0.575	17.469	17.472	5
0.2%	148	0	2	POST	^/api/posts/[^/]+/realmojis$	47.419	0.320	18.419	19.379	473
0.2%	532	0	1	GET	^/api/memories/[^/]+	46.216	0.087	0.248	17.343	766
0.2%	14	0	4	POST	/api/admin/posting-window	42.937	3.067	17.159	17.159	174
0.2%	117	2	2	GET	/api/me	41.761	0.357	17.194	17.510	295
0.1%	251	0	1	GET	^/api/relationships/[^/]+	18.376	0.073	0.132	17.492	1029
0.1%	151	0	1	GET	^/api/posts/[^/]+/realmojis$	17.817	0.118	0.092	17.480	1284
0.1%	1	0	0	POST	/api/initialize	15.126	15.126	15.126	15.126	90
0.0%	11	0	0	DELETE	^/api/posts/[^/]+$	7.012	0.637	5.831	5.831	0
0.0%	18	0	0	DELETE	^/api/realmojis/[^/]+$	2.231	0.124	0.194	0.194	0
0.0%	319	0	0	GET	^/api/posts/[^/]+$	2.085	0.007	0.039	0.927	605
0.0%	364	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	0.945	0.003	0.024	0.027	116
0.0%	124	0	0	GET	^/api/posts/[^/]+/comments$	0.209	0.002	0.025	0.033	1450
0.0%	132	0	0	POST	^/api/posts/[^/]+/comments$	0.195	0.001	0.022	0.032	313
0.0%	148	0	0	GET	^/api/users/[^/]+$	0.185	0.001	0.017	0.025	184
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
56.1%	88	0	0	GET	/api/notifications	1506.565	17.120	26.085	26.085	68
21.6%	86	9	0	POST	/api/posts	579.444	6.738	64.236	64.236	582
5.7%	399	0	10	POST	/api/auth/login	151.800	0.380	15.508	18.360	246
5.1%	1243	2	9	GET	^/api/images/[^/]+	138.033	0.111	1.694	5.302	95923
2.9%	918	0	2	GET	^/api/feeds/[^/]+	77.481	0.084	0.131	17.587	2080
2.2%	251	0	1	GET	/api/memories	57.866	0.231	7.658	18.463	18378
1.2%	319	4	0	POST	/api/friend-requests	33.437	0.105	0.129	0.148	114
0.9%	27	0	1	GET	/api/me/export	23.682	0.877	18.291	18.291	3064103
0.6%	126	0	1	GET	/api/posting-window/current	16.790	0.133	0.119	9.330	237
0.5%	65	0	1	POST	/api/auth/logout	14.538	0.224	7.099	7.099	0
0.3%	150	0	0	GET	/api/users	9.143	0.061	0.107	0.124	278
0.3%	139	0	0	GET	^/api/posts/[^/]+$	8.618	0.062	0.123	0.134	629
0.3%	75	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	7.835	0.104	0.123	0.123	116
0.2%	53	0	0	POST	/api/auth/register	6.235	0.118	0.273	0.273	228
0.2%	100	0	0	GET	^/api/relationships/[^/]+	6.076	0.061	0.117	0.120	354
0.2%	41	0	0	POST	^/api/posts/[^/]+/realmojis$	5.904	0.144	0.405	0.405	466
0.2%	77	0	0	GET	^/api/posts/me/current$	4.687	0.061	0.113	0.113	650
0.2%	42	0	0	POST	^/api/posts/[^/]+/comments$	4.347	0.104	0.122	0.122	324
0.2%	63	0	0	GET	^/api/memories/[^/]+	4.301	0.068	0.245	0.245	947
0.2%	67	0	0	GET	^/api/posts/[^/]+/realmojis$	4.256	0.064	0.131	0.131	1141
0.2%	64	0	0	GET	^/api/users/[^/]+$	4.186	0.065	0.165	0.165	194
0.1%	33	0	0	PUT	/api/me/profile-image	3.917	0.119	0.306	0.306	0
0.1%	17	0	0	DELETE	^/api/realmojis/[^/]+$	3.753	0.221	0.344	0.344	0
0.1%	50	0	0	GET	^/api/posts/[^/]+/comments$	3.139	0.063	0.122	0.122	1641
0.1%	27	0	0	PUT	^/api/follows/[^/]+$	2.926	0.108	0.135	0.135	0
0.1%	22	0	0	PATCH	/api/me	2.262	0.103	0.108	0.108	265
0.1%	34	1	0	GET	/api/me	2.162	0.064	0.232	0.232	299
0.0%	11	0	0	PATCH	^/api/posts/[^/]+$	1.127	0.102	0.107	0.107	646
0.0%	3	0	0	DELETE	^/api/posts/[^/]+$	0.610	0.203	0.205	0.205	0
0.0%	5	0	0	PUT	^/api/blocks/[^/]+$	0.506	0.101	0.102	0.102	0
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
49.2%	91	0	0	GET	/api/notifications	1378.218	15.145	23.174	23.174	68
8.6%	628	2	17	GET	^/api/feeds/[^/]+	239.738	0.382	17.115	18.498	2139
5.9%	337	0	13	POST	/api/auth/login	165.541	0.491	13.934	18.366	243
5.8%	1016	2	11	GET	^/api/images/[^/]+	161.512	0.159	5.001	6.630	104985
5.6%	85	2	1	POST	/api/posts	157.525	1.853	65.306	65.306	618
5.3%	295	6	9	POST	/api/friend-requests	149.041	0.505	17.425	18.355	110
2.6%	138	0	4	GET	/api/posting-window/current	71.634	0.519	18.466	18.504	233
2.3%	34	0	3	POST	^/api/posts/[^/]+/realmojis$	64.293	1.891	19.473	19.473	452
2.1%	164	0	3	GET	/api/users	57.799	0.352	18.456	18.888	247
1.9%	192	0	1	GET	/api/memories	54.043	0.281	6.748	18.455	25659
1.9%	29	0	3	PATCH	/api/me	53.946	1.860	18.415	18.415	244
1.7%	66	0	3	POST	^/api/friend-requests/[^/]+/[^/]+$	47.758	0.724	18.218	18.218	111
1.1%	58	0	1	GET	^/api/memories/[^/]+	29.924	0.516	18.466	18.466	930
0.9%	49	0	3	PUT	^/api/follows/[^/]+$	24.161	0.493	12.164	12.164	0
0.8%	49	0	1	GET	^/api/posts/[^/]+/comments$	21.921	0.447	18.461	18.461	1304
0.8%	74	0	1	GET	^/api/relationships/[^/]+	21.236	0.287	14.871	14.871	512
0.7%	82	0	1	GET	^/api/posts/me/current$	19.777	0.241	12.203	12.203	632
0.7%	1	0	1	DELETE	/api/me/profile-image	18.272	18.272	18.272	18.272	0
0.3%	6	0	1	DELETE	^/api/realmojis/[^/]+$	9.406	1.568	4.491	4.491	0
0.3%	60	0	0	POST	/api/auth/register	7.649	0.127	0.254	0.254	229
0.3%	52	0	0	POST	/api/auth/logout	7.551	0.145	0.787	0.787	0
0.2%	91	0	0	GET	^/api/posts/[^/]+$	5.997	0.066	0.132	0.132	615
0.2%	35	0	0	POST	^/api/posts/[^/]+/comments$	5.749	0.164	0.756	0.756	300
0.2%	29	0	0	PUT	/api/me/profile-image	5.240	0.181	0.933	0.933	0
0.2%	20	0	0	GET	/api/me/export	5.148	0.257	2.068	2.068	3323034
0.2%	64	0	0	GET	^/api/posts/[^/]+/realmojis$	4.308	0.067	0.150	0.150	1222
0.2%	9	0	1	DELETE	^/api/comments/[^/]+$	4.256	0.473	3.376	3.376	0
0.1%	49	0	0	GET	^/api/users/[^/]+$	3.303	0.067	0.151	0.151	193
0.1%	29	0	0	GET	/api/me	2.778	0.096	1.054	1.054	307
0.0%	9	0	0	PATCH	^/api/posts/[^/]+$	0.997	0.111	0.171	0.171	655
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 9205 クエリ / 48 種類 / 合計 109.855s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	44.6%	202	48.992	242.54	9542.43	1184.1	120.6	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	19.9%	3	21.880	7293.32	9816.80	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	14.7%	3	16.137	5379.09	7247.04	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	8.2%	7090	8.967	1.26	11.43	9.8	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	7.7%	3	8.445	2815.00	3908.45	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	3.0%	4	3.257	814.26	2745.94	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.3%	4	0.361	90.18	158.90	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
8	0.3%	4	0.361	90.15	159.54	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.2%	1	0.271	271.16	271.16	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.251	250.55	250.55	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.2%	1	0.220	219.56	219.56	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.164	163.87	163.87	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.105	104.69	104.69	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.080	80.36	80.36	2412.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.1%	1	0.067	67.11	67.11	1503.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
未取得: プロファイルを取れませんでした (isu1)
  http://localhost:6060/debug/pprof/profile?seconds=60

curl: (52) Empty reply from server
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 48.6% / 100% | 26% | 3.3% | 5 | 105 MB | run queue が積もっている |
| isu2 (2 cpu) | 18.9% / 79% | 0% | 0.2% | 3 | 1308 MB |  |
| isu3 (2 cpu) | 16.3% / 78% | 0% | 0.3% | 3 | 989 MB |  |
| isu4 (2 cpu) | 35.4% / 100% | 13% | 17.9% | 3 | 104 MB |  |

**isu1（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 78% / nginx 8% / amazon-ssm-agen 1% |
| isu2 | isureal 36% |
| isu3 | isureal 31% |
| isu4 | postgres 17% / isureal 15% |

## Diff from previous run

スコア: 186885 → — (fail)

`dca9471..106c0a8`

```
 webapp/go/image_signature.go |  67 ++++++++++++++++++++-------
 webapp/go/offload.go         | 105 ++++++++++++++++++++++---------------------
 2 files changed, 106 insertions(+), 66 deletions(-)
```
