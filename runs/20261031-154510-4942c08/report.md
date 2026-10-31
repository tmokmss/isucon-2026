# run 20261031-154510-4942c08

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | eu-us-wait |
| sha | 4942c089b9e0 |
| 所要時間 | 167s |
| 前回比 | -186373 (-100.0%) vs 20261031-153729-0ce5dec |
| ベスト比 | -191816 (-100.0%) vs 20261031-151339-5a230a4 = 191816 |
| note | eu/us: 画像取得の同時要求を1本に・接続の使い回し（main b3440cf を取り込んだ版） |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/agent-a4fb2eece2777c8df/runs/20261031-154510-4942c08/bench.log
```

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	134	391	1070	1071
europe	7	20	59	171	465	465
america	6	21	59	170	330	330
```

```
{"id":"24416","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"負荷走行を打ち切りました"}

--- stdout ---
time=2026-10-31T06:46:55.239Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:46:58.230Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:47:01.228Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:47:04.050Z level=INFO msg="user moved" user=noashimia-5119 from=america to=europe
time=2026-10-31T06:47:04.238Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T06:47:07.020Z level=INFO msg="user moved" user=yuzoe5279 from=asia to=america
time=2026-10-31T06:47:07.227Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T06:47:09.638Z level=INFO msg="user moved" user=noashimia-5119 from=europe to=america
time=2026-10-31T06:47:10.229Z level=INFO msg="america 2 日目: 21 人"
time=2026-10-31T06:47:10.282Z level=INFO msg="user moved" user=yuzoe5279 from=america to=asia
time=2026-10-31T06:47:13.239Z level=INFO msg="asia 3 日目: 134 人"
time=2026-10-31T06:47:13.976Z level=INFO msg="user moved" user=rais-5612 from=asia to=america
time=2026-10-31T06:47:14.709Z level=INFO msg="user moved" user=rais-5612 from=america to=asia
time=2026-10-31T06:47:15.222Z level=INFO msg="user moved" user=waya.5523 from=america to=europe
time=2026-10-31T06:47:16.204Z level=INFO msg="user moved" user=matorno_5812 from=asia to=europe
time=2026-10-31T06:47:16.235Z level=INFO msg="europe 3 日目: 59 人"
time=2026-10-31T06:47:17.036Z level=INFO msg="user moved" user=yuzoe5279 from=asia to=america
time=2026-10-31T06:47:19.228Z level=INFO msg="america 3 日目: 59 人"
time=2026-10-31T06:47:20.740Z level=INFO msg="user moved" user=waya.5523 from=europe to=america
time=2026-10-31T06:47:21.146Z level=INFO msg="user moved" user=maevaan-6381 from=europe to=asia
time=2026-10-31T06:47:21.700Z level=INFO msg="user moved" user=noashimia-5119 from=america to=europe
time=2026-10-31T06:47:22.297Z level=INFO msg="asia 4 日目: 391 人"
time=2026-10-31T06:47:23.285Z level=INFO msg="user moved" user=rais-5612 from=asia to=america
time=2026-10-31T06:47:23.525Z level=INFO msg="user moved" user=benyufu.6718 from=asia to=america
time=2026-10-31T06:47:24.574Z level=INFO msg="user moved" user=ivyho.6681 from=asia to=europe
time=2026-10-31T06:47:24.636Z level=INFO msg="user moved" user=benyufu.6718 from=america to=europe
time=2026-10-31T06:47:24.987Z level=INFO msg="user moved" user=maevaan-6381 from=asia to=europe
time=2026-10-31T06:47:25.340Z level=INFO msg="europe 4 日目: 171 人"
time=2026-10-31T06:47:25.439Z level=INFO msg="user moved" user=tanusu.6957 from=asia to=europe
time=2026-10-31T06:47:27.735Z level=INFO msg="user moved" user=kona.8261 from=asia to=europe
time=2026-10-31T06:47:28.232Z level=INFO msg="user moved" user=kumano-7972 from=asia to=america
time=2026-10-31T06:47:28.334Z level=INFO msg="user moved" user=noashimia-5119 from=europe to=america
time=2026-10-31T06:47:28.335Z level=INFO msg="america 4 日目: 170 人"
time=2026-10-31T06:47:29.158Z level=INFO msg="user moved" user=ruvin8540 from=asia to=europe
time=2026-10-31T06:47:29.570Z level=INFO msg="user moved" user=tsuvinna_8822 from=europe to=asia
time=2026-10-31T06:47:30.495Z level=INFO msg="user moved" user=noatormar-8500 from=europe to=asia
time=2026-10-31T06:47:31.317Z level=INFO msg="asia 5 日目: 1070 人"
time=2026-10-31T06:47:31.593Z level=INFO msg="user moved" user=shiul.9158 from=europe to=asia
time=2026-10-31T06:47:33.645Z level=INFO msg="user moved" user=ulruyo_6635 from=america to=asia
time=2026-10-31T06:47:35.002Z level=INFO msg="europe 5 日目: 465 人"
time=2026-10-31T06:47:35.207Z level=INFO msg="user moved" user=honoatsu-9096 from=europe to=america
time=2026-10-31T06:47:35.746Z level=INFO msg="user moved" user=waya.5523 from=america to=europe
time=2026-10-31T06:47:38.587Z level=INFO msg="america 5 日目: 330 人"
time=2026-10-31T06:47:47.132Z level=INFO msg="europe 最終日（グランプリ）: 465 人"
time=2026-10-31T06:47:47.151Z level=INFO msg="asia 最終日（グランプリ）: 1071 人"
time=2026-10-31T06:47:47.316Z level=INFO msg="america 最終日（グランプリ）: 330 人"
time=2026-10-31T06:47:55.272Z level=INFO msg="グランプリ: アジア 344 人, ヨーロッパ 220 人, アメリカ 103 人"
{
  "pass": false,
  "score": 0,
  "messages": [
    "負荷走行を打ち切りました: 55s〜1m0s のエラー率 76.9%（11603/15081 件）が上限 10.0% を超えました",
    "kuko-15483: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "wanoa17721: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "yasa.17233: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maroni16571: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "neko.8343: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noamiaha6667: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rini13827: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "chito-13072: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "mita-11407: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marru-17985: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "isis.9595: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "teyomar.7801: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ruivy-7559: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "koma17938: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hiro.12788: 投稿できません: POST /api/posts: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/posts\": readfrom tcp 10.27.37.77:59952-\u003e52.69.202.109:8080: write tcp 10.27.37.77:59952-\u003e52.69.202.109:8080: use of closed network connection",
    "evaki-17678: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kibenel-6888: 今日の投稿を取得できません: GET /api/posts/me/current: ステータスコード 502 が返りました（期待値 200、unknown）",
    "minjun_kim_4078: 履歴を読めません: GET /api/images/e2c69dcf-675f-44a2-8c73-067364eac63d?exp=1793515680\u0026sig=lxC2Gck2_La68Akxb-d5b9Z9_N8XvOXbvFTShFcf5qQ: ステータスコード 502 が返りました（期待値 200、unknown）",
    "teheri-6901: 投稿できません: POST /api/posts: ステータスコード 502 が返りました（期待値 201、unknown）",
    "tomu-5631: 投稿 b1cce782-9539-4c08-8c68-507a4ede3ec5 の画像を取得できません: GET /api/images/707af79d-823f-4fa9-85b9-8ced76951b90?exp=1793515680\u0026sig=l4W81CgutXK5WAeURGQuifWIBen4p1Phy6O3u6PHKPg: ステータスコード 502 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
56.2%	3661	1	908	GET	/api/notifications	22608.700	6.176	15.943	15.984	120
8.2%	6073	67	218	POST	/api/friend-requests	3309.286	0.545	2.266	3.540	116
6.9%	2644	111	512	POST	/api/posts	2777.615	1.051	6.877	7.170	532
4.4%	19018	307	251	GET	^/api/images/[^/]+	1753.501	0.092	0.905	2.419	109336
3.7%	10874	610	4043	POST	/api/auth/login	1495.532	0.138	1.645	4.640	203
2.9%	3132	18	269	POST	^/api/friend-requests/[^/]+/[^/]+$	1172.284	0.374	1.991	2.969	120
2.5%	768	125	15	PUT	/api/me/profile-image	991.677	1.291	11.078	14.947	6
2.4%	16213	0	3226	GET	^/api/feeds/[^/]+	955.249	0.059	0.529	1.367	2743
2.2%	4193	0	79	GET	/api/memories	899.049	0.214	1.711	5.256	13925
1.5%	628	9	12	POST	^/api/posts/[^/]+/realmojis$	620.268	0.988	3.442	6.321	466
1.2%	701	7	11	POST	^/api/posts/[^/]+/comments$	498.080	0.711	3.111	3.841	337
1.0%	764	0	15	PUT	^/api/follows/[^/]+$	395.719	0.518	2.063	3.207	3
0.7%	3871	0	63	GET	/api/posting-window/current	282.587	0.073	0.550	1.462	234
0.7%	1111	0	26	POST	/api/auth/logout	272.927	0.246	1.942	2.753	4
0.7%	526	0	7	PATCH	/api/me	269.616	0.513	2.434	4.588	280
0.7%	2867	0	110	GET	/api/users	262.427	0.092	0.567	1.382	243
0.6%	400	0	5	GET	/api/me/export	252.205	0.631	3.490	5.790	1912489
0.6%	1093	1	0	POST	/api/auth/register	249.892	0.229	0.991	1.513	227
0.4%	2022	39	92	GET	^/api/posts/me/current$	175.980	0.087	0.582	1.453	618
0.3%	637	3	0	GET	/api/me	130.044	0.204	0.689	1.386	302
0.3%	1880	24	4	GET	^/api/posts/[^/]+$	117.772	0.063	0.472	0.869	613
0.3%	224	2	51	PATCH	^/api/posts/[^/]+$	117.548	0.525	1.979	2.730	534
0.3%	1804	0	337	GET	^/api/relationships/[^/]+	110.755	0.061	0.496	1.094	775
0.2%	1004	0	1	GET	^/api/memories/[^/]+	94.581	0.094	0.692	1.375	580
0.2%	89	0	14	DELETE	^/api/posts/[^/]+$	79.950	0.898	3.321	3.321	26
0.2%	119	0	0	DELETE	^/api/realmojis/[^/]+$	76.452	0.642	2.703	3.325	0
0.1%	988	0	2	GET	^/api/users/[^/]+$	54.667	0.055	0.443	0.856	199
0.1%	916	0	7	GET	^/api/posts/[^/]+/realmojis$	48.598	0.053	0.456	0.863	1445
0.1%	788	0	7	GET	^/api/posts/[^/]+/comments$	43.959	0.056	0.452	0.700	2301
0.1%	138	0	0	DELETE	^/api/comments/[^/]+$	42.170	0.306	1.738	1.907	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
53.5%	1488	0	442	GET	/api/notifications	7722.590	5.190	11.725	11.728	117
9.0%	2397	27	120	POST	/api/friend-requests	1304.750	0.544	2.271	2.739	108
6.4%	993	32	151	POST	/api/posts	918.018	0.924	6.930	7.191	528
5.3%	1438	0	58	POST	/api/auth/login	758.742	0.528	2.455	3.508	235
3.9%	4435	0	248	GET	^/api/feeds/[^/]+	556.617	0.126	0.823	1.594	3319
3.5%	460	0	0	POST	/api/auth/register	507.627	1.104	2.790	2.851	229
3.3%	4336	15	97	GET	^/api/images/[^/]+	475.546	0.110	0.897	2.445	164292
2.0%	243	0	1	PUT	/api/me/profile-image	282.060	1.161	2.998	3.238	0
1.9%	547	1	11	POST	^/api/friend-requests/[^/]+/[^/]+$	268.398	0.491	1.968	2.478	113
1.6%	335	0	4	PUT	^/api/follows/[^/]+$	231.144	0.690	1.890	2.228	0
1.4%	220	0	1	PATCH	/api/me	198.314	0.901	2.579	3.022	272
1.1%	1530	0	11	GET	/api/posting-window/current	162.247	0.106	0.713	1.615	237
1.0%	1204	0	56	GET	/api/users	151.170	0.126	0.821	1.547	234
0.9%	966	0	28	GET	/api/memories	130.849	0.135	1.358	2.598	14040
0.8%	208	1	1	POST	^/api/posts/[^/]+/comments$	121.022	0.582	2.572	3.782	320
0.7%	810	59	54	GET	^/api/posts/me/current$	99.642	0.123	0.847	1.580	561
0.6%	108	3	2	POST	^/api/posts/[^/]+/realmojis$	87.147	0.807	3.872	3.888	456
0.6%	222	0	3	POST	/api/auth/logout	79.860	0.360	1.939	2.168	0
0.5%	507	0	54	GET	^/api/relationships/[^/]+	70.593	0.139	1.291	1.561	456
0.5%	581	21	3	GET	^/api/posts/[^/]+$	66.572	0.115	0.746	1.011	595
0.3%	93	8	18	PATCH	^/api/posts/[^/]+$	46.757	0.503	2.266	2.266	466
0.3%	295	0	3	GET	^/api/users/[^/]+$	36.172	0.123	0.757	1.485	194
0.2%	249	0	2	GET	^/api/posts/[^/]+/comments$	26.486	0.106	0.780	0.939	2278
0.2%	32	0	10	DELETE	^/api/posts/[^/]+$	24.380	0.762	2.401	2.401	0
0.2%	236	0	3	GET	^/api/posts/[^/]+/realmojis$	23.257	0.099	0.447	0.568	1337
0.1%	247	2	0	GET	/api/me	20.502	0.083	0.215	0.233	303
0.1%	166	0	1	GET	^/api/memories/[^/]+	18.726	0.113	1.165	1.211	1079
0.1%	63	0	0	GET	/api/me/export	14.212	0.226	5.847	5.847	2165209
0.1%	12	0	0	DELETE	^/api/realmojis/[^/]+$	12.388	1.032	2.476	2.476	0
0.1%	38	0	0	DELETE	^/api/comments/[^/]+$	12.004	0.316	1.042	1.042	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
54.8%	1093	0	315	GET	/api/notifications	5623.771	5.145	10.428	10.581	121
7.3%	1513	13	13	POST	/api/friend-requests	744.236	0.492	2.126	2.989	113
7.0%	731	40	146	POST	/api/posts	722.708	0.989	6.848	7.177	487
5.6%	1214	0	20	POST	/api/auth/login	571.139	0.470	2.330	4.715	245
4.6%	3680	16	20	GET	^/api/images/[^/]+	472.385	0.128	1.087	3.565	136564
3.7%	2760	0	43	GET	^/api/feeds/[^/]+	380.985	0.138	1.266	2.362	3031
2.4%	327	0	0	POST	/api/auth/register	248.198	0.759	1.886	1.910	229
1.8%	388	2	9	POST	^/api/friend-requests/[^/]+/[^/]+$	183.297	0.472	2.340	2.493	113
1.3%	1097	0	11	GET	/api/posting-window/current	136.542	0.124	0.949	2.621	238
1.3%	249	0	0	PUT	^/api/follows/[^/]+$	128.219	0.515	2.466	2.492	0
1.2%	191	0	0	PUT	/api/me/profile-image	123.848	0.648	1.975	3.581	0
1.1%	819	0	14	GET	/api/memories	117.543	0.144	1.663	2.935	19022
1.0%	146	0	0	PATCH	/api/me	101.650	0.696	2.016	2.382	282
1.0%	736	0	14	GET	/api/users	99.906	0.136	1.102	2.361	242
0.8%	73	2	2	POST	^/api/posts/[^/]+/realmojis$	82.786	1.134	3.211	3.211	445
0.8%	517	11	11	GET	^/api/posts/me/current$	79.372	0.154	1.830	2.620	626
0.8%	146	1	0	POST	^/api/posts/[^/]+/comments$	77.462	0.531	3.061	3.107	377
0.7%	179	0	2	POST	/api/auth/logout	69.431	0.388	2.463	2.491	0
0.5%	319	0	9	GET	^/api/relationships/[^/]+	48.634	0.152	1.609	1.991	526
0.4%	392	21	0	GET	^/api/posts/[^/]+$	41.835	0.107	0.709	0.914	594
0.3%	50	4	6	PATCH	^/api/posts/[^/]+$	35.741	0.715	2.488	2.488	520
0.3%	76	0	2	GET	/api/me/export	26.703	0.351	4.756	4.756	2920015
0.3%	193	0	0	GET	/api/me	26.259	0.136	0.345	0.373	306
0.2%	215	0	3	GET	^/api/users/[^/]+$	23.076	0.107	0.778	0.969	203
0.2%	21	0	3	DELETE	^/api/posts/[^/]+$	21.844	1.040	2.175	2.175	0
0.2%	172	0	1	GET	^/api/memories/[^/]+	18.041	0.105	0.334	1.936	1076
0.1%	116	0	0	GET	^/api/posts/[^/]+/realmojis$	12.859	0.111	0.312	0.330	1128
0.1%	119	0	0	GET	^/api/posts/[^/]+/comments$	11.618	0.098	0.266	0.892	2470
0.1%	34	0	0	DELETE	^/api/comments/[^/]+$	11.474	0.337	2.589	2.589	0
0.1%	8	0	0	DELETE	^/api/realmojis/[^/]+$	6.334	0.792	3.445	3.445	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 15910 クエリ / 48 種類 / 合計 156.906s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	55.1%	410	86.474	210.91	1038.53	1245.8	127.6	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.4%	3	22.637	7545.54	10066.03	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.0%	13858	17.326	1.25	10.51	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.1%	3	15.882	5294.07	7172.77	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.2%	3	8.160	2719.97	3806.82	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.3%	6	3.574	595.64	2714.44	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.4%	6	0.692	115.37	161.58	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.4%	6	0.665	110.80	161.37	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.249	249.30	249.30	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	0.2%	1	0.247	247.41	247.41	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.1%	1	0.221	220.50	220.50	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.162	162.30	162.30	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.114	113.62	113.62	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.084	84.45	84.45	2486.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.070	69.69	69.69	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT data FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 4652f906c9b030c81aabad690ac9e67388a28e0e
Type: cpu
Time: 2026-10-31 15:46:33 JST
Duration: 60.04s, Total samples = 42.43s (70.67%)
Showing nodes accounting for 30.54s, 71.98% of 42.43s total
Dropped 1104 nodes (cum <= 0.21s)
Showing top 40 nodes out of 278
      flat  flat%   sum%        cum   cum%
     5.62s 13.25% 13.25%      5.62s 13.25%  crypto/internal/fips140/sha256.blockAVX2
     4.85s 11.43% 24.68%      4.85s 11.43%  internal/runtime/syscall/linux.Syscall6
     2.05s  4.83% 29.51%      2.05s  4.83%  runtime.memmove
     1.78s  4.20% 33.70%      5.94s 14.00%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.73s  4.08% 37.78%      2.25s  5.30%  runtime.findObject
     1.69s  3.98% 41.76%      1.69s  3.98%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.47s  3.46% 45.23%      1.47s  3.46%  runtime.memclrNoHeapPointers
     1.33s  3.13% 48.36%      2.40s  5.66%  runtime.tryDeferToSpanScan
     1.30s  3.06% 51.43%      5.97s 14.07%  runtime.scanObject
     1.27s  2.99% 54.42%      1.27s  2.99%  aeshashbody
     0.77s  1.81% 56.23%      0.77s  1.81%  runtime.spanClass.sizeclass (inline)
     0.54s  1.27% 57.51%      2.25s  5.30%  runtime.mapaccess1_faststr
     0.45s  1.06% 58.57%      0.68s  1.60%  example.com/isureal-company/isureal.postLess
     0.41s  0.97% 59.53%      0.41s  0.97%  memeqbody
     0.37s  0.87% 60.41%      0.37s  0.87%  runtime.nextFreeFast (inline)
     0.36s  0.85% 61.25%      0.36s  0.85%  runtime.futex
     0.34s   0.8% 62.06%      6.38s 15.04%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.32s  0.75% 62.81%      6.27s 14.78%  crypto/internal/fips140/sha256.(*Digest).Write
     0.28s  0.66% 63.47%      6.75s 15.91%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.25s  0.59% 64.06%      0.47s  1.11%  compress/flate.(*compressor).deflate
     0.25s  0.59% 64.65%      0.27s  0.64%  encoding/json.appendString[go.shape.string]
     0.24s  0.57% 65.21%      0.24s  0.57%  runtime.(*mspan).divideByElemSize (inline)
     0.24s  0.57% 65.78%      0.46s  1.08%  runtime.greyobject
     0.24s  0.57% 66.34%      0.98s  2.31%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.57% 66.91%      0.25s  0.59%  sync/atomic.(*Int32).Add (inline)
     0.21s  0.49% 67.41%      0.73s  1.72%  database/sql.convertAssignRows
     0.20s  0.47% 67.88%      0.25s  0.59%  runtime.(*mspan).writeHeapBitsSmall
     0.18s  0.42% 68.30%      3.16s  7.45%  runtime.mallocgc
     0.18s  0.42% 68.72%      1.89s  4.45%  runtime.mapassign_faststr
     0.17s   0.4% 69.13%      0.43s  1.01%  runtime.scanObjectsSmall
     0.16s  0.38% 69.50%      0.73s  1.72%  encoding/json.structEncoder.encode
     0.15s  0.35% 69.86%      0.41s  0.97%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.13s  0.31% 70.16%      3.70s  8.72%  example.com/isureal-company/isureal.loadBig
     0.13s  0.31% 70.47%      0.75s  1.77%  runtime.wbBufFlush1
     0.11s  0.26% 70.73%      0.92s  2.17%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.11s  0.26% 70.99%      0.23s  0.54%  internal/runtime/maps.(*Iter).Next
     0.11s  0.26% 71.25%      0.84s  1.98%  runtime.mallocgcSmallNoscan
     0.11s  0.26% 71.51%      0.57s  1.34%  runtime.mapaccess2_faststr
     0.10s  0.24% 71.74%      7.07s 16.66%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.10s  0.24% 71.98%      5.72s 13.48%  crypto/internal/fips140/sha256.block
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 57.4% / 100% | 37% | 2.1% | 6 | 269 MB | CPU が飽和 |
| isu2 (2 cpu) | 40.2% / 100% | 7% | 0.8% | 4 | 330 MB |  |
| isu3 (2 cpu) | 34.9% / 100% | 5% | 0.6% | 4 | 187 MB |  |
| isu4 (2 cpu) | 56.1% / 100% | 22% | 24.2% | 6 | 106 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu4（IO 待ちが重い、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 86% / nginx 19% |
| isu2 | isureal 76% / nginx 3% |
| isu3 | isureal 67% / nginx 2% |
| isu4 | isureal 38% / postgres 22% |

## Diff from previous run

スコア: 186373 → — (fail)

`0ce5dec..4942c08`

```
 conf/_all/etc/sysctl.d/99-isucon.conf |  2 +-
 webapp/go/admin_api.go                | 27 ++++++++++++++++----
 webapp/go/export_api.go               | 12 +++++++--
 webapp/go/image_api.go                | 23 ++++++++++--------
 webapp/go/image_file.go               | 33 ++++++++++++++-----------
 webapp/go/image_held.go               |  9 +++----
 webapp/go/image_local.go              | 46 +++++++++++++++--------------------
 webapp/go/ip_ban.go                   | 11 +++++++++
 webapp/go/repl.go                     |  1 +
 webapp/go/replica_server.go           |  5 +++-
 10 files changed, 104 insertions(+), 65 deletions(-)
```
