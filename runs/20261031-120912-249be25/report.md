# run 20261031-120912-249be25

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | origin/main |
| sha | 249be2509f96 |
| 所要時間 | 171s |
| 前回比 | -56314 (-100.0%) vs 20261031-115822-a8d9e7d |
| ベスト比 | -60201 (-100.0%) vs 20261031-115040-1c73ee4 = 60201 |
| note | main: +#71 JPEG の検証を europe1・america1 へ（再起動試験の後） |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/contest-day-hosts/runs/20261031-120912-249be25/bench.log
```

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	49	141	394	471	470
europe	7	17	44	106	107	106
america	6	14	40	59	60	60
```

```
{"id":"6978","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"負荷走行を打ち切りました"}

--- stdout ---
time=2026-10-31T03:10:59.388Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:11:02.350Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:11:05.351Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:11:08.376Z level=INFO msg="asia 2 日目: 49 人"
time=2026-10-31T03:11:11.353Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T03:11:14.350Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T03:11:17.376Z level=INFO msg="asia 3 日目: 141 人"
time=2026-10-31T03:11:20.398Z level=INFO msg="europe 3 日目: 44 人"
time=2026-10-31T03:11:21.103Z level=INFO msg="user moved" user=hehara.5634 from=asia to=america
time=2026-10-31T03:11:21.161Z level=INFO msg="user moved" user=tato.5981 from=asia to=america
time=2026-10-31T03:11:21.386Z level=INFO msg="user moved" user=funo_6059 from=asia to=america
time=2026-10-31T03:11:22.919Z level=INFO msg="user moved" user=tato.5981 from=america to=asia
time=2026-10-31T03:11:23.351Z level=INFO msg="america 3 日目: 40 人"
time=2026-10-31T03:11:26.456Z level=INFO msg="asia 4 日目: 394 人"
time=2026-10-31T03:11:29.313Z level=INFO msg="user moved" user=tato.5981 from=asia to=america
time=2026-10-31T03:11:29.388Z level=INFO msg="europe 4 日目: 106 人"
time=2026-10-31T03:11:32.395Z level=INFO msg="america 4 日目: 59 人"
time=2026-10-31T03:11:35.407Z level=INFO msg="user moved" user=suta-8134 from=asia to=america
time=2026-10-31T03:11:35.441Z level=INFO msg="asia 5 日目: 471 人"
time=2026-10-31T03:11:35.499Z level=INFO msg="user moved" user=kevinchi-6947 from=asia to=europe
time=2026-10-31T03:11:38.522Z level=INFO msg="europe 5 日目: 107 人"
time=2026-10-31T03:11:41.423Z level=INFO msg="america 5 日目: 60 人"
time=2026-10-31T03:11:50.652Z level=INFO msg="europe 最終日（グランプリ）: 106 人"
time=2026-10-31T03:11:50.795Z level=INFO msg="asia 最終日（グランプリ）: 470 人"
time=2026-10-31T03:11:50.975Z level=INFO msg="america 最終日（グランプリ）: 60 人"
time=2026-10-31T03:11:59.366Z level=INFO msg="グランプリ: アジア 27 人, ヨーロッパ 2 人, アメリカ 2 人"
{
  "pass": false,
  "score": 0,
  "messages": [
    "負荷走行を打ち切りました: 55s〜1m0s のエラー率 13.4%（42/314 件）が上限 10.0% を超えました",
    "maliara-8783: プロフィール画像を取得できません: GET /api/images/dd3ae67f-8d52-43c4-a178-3e4521632092?exp=1793502720\u0026sig=QslaVy2pGZrmZ69-bbCzgi1Xd9bOA7_ZFy5LPbXGEBs: ステータスコード 500 が返りました（期待値 200、failed to connect to `user=isureal database=isureal`:",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)）",
    "kazoemu8799: プロフィール画像を取得できません: GET /api/images/9761de9c-3bc0-4643-916d-e0527bbb38b8?exp=1793502720\u0026sig=UIJVRmMflAYvqDfhH4yx86N_jIiioyTWuEhmWoBr5Qo: ステータスコード 500 が返りました（期待値 200、failed to connect to `user=isureal database=isureal`:",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)）",
    "sushi_on_a_sunflower_4749: 履歴を読めません: GET /api/images/93188968-8cc6-4918-8437-7add6ac84068?exp=1793502720\u0026sig=d2hU17ZXeIxTgAh8HalvYZFeKZ-CQYjy7803P_dTB9c: ステータスコード 500 が返りました（期待値 200、failed to connect to `user=isureal database=isureal`:",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)）",
    "sardinas_en_olla_de_barro_4671: 履歴を読めません: GET /api/images/07ce45ce-1854-40ee-9ef0-2c66b881be67?exp=1793502720\u0026sig=0-HvR4FSw-6XnoM6R7InfQDvUv_OVK54R9G4d_pz-II: ステータスコード 500 が返りました（期待値 200、failed to connect to `user=isureal database=isureal`:",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)）",
    "salmon_in_a_jar_2533: 履歴を読めません: GET /api/images/b9addc54-4149-4b5f-bd67-a4493a187cd3?exp=1793502720\u0026sig=RkssCIWstwjuVl9VjsizRJZfMzu7bloVpCeYKMm3dyc: ステータスコード 500 が返りました（期待値 200、failed to connect to `user=isureal database=isureal`:",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)）",
    "mia_jackson_423: 履歴を読めません: GET /api/images/ac4b8915-7da1-41cc-90b7-abcdb079eeb2?exp=1793502720\u0026sig=3boyEg7XuQBzan_gPMzJOKDVuXlwbNMTIN40EjcFpvI: ステータスコード 500 が返りました（期待値 200、failed to connect to `user=isureal database=isureal`:",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)）",
    "sebastian_ramirez_219: 履歴を読めません: GET /api/images/db408bc4-5ca7-4af9-a145-b4f5efb07c15?exp=1793502720\u0026sig=wNkNjVbLfBSRKCjfd-6Hi8wAWJU5W00G1HUS3KSJldg: ステータスコード 500 が返りました（期待値 200、failed to connect to `user=isureal database=isureal`:",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)）",
    "hehara.5634: 投稿 5255632e-d376-49fc-a56d-9726faa9ab70 の画像を取得できません: GET /api/images/3d0abb3a-693e-4406-86b8-5fab6fe8d818?exp=1793502720\u0026sig=cJDBNbuThAP-p8DOf00Z9x3u2VlqGlFfoWp3L6uusXA: ステータスコード 500 が返りました（期待値 200、failed to connect to `user=isureal database=isureal`:",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)",
    "\t192.168.4.10:5432 (192.168.4.10): server error: FATAL: sorry, too many clients already (SQLSTATE 53300)）",
    "mamarel7735: 認定公開アカウントをフォローできません: PUT /api/follows/7f09e043-412d-4ae0-8a59-0b18a5cfb839: 応答がタイムアウトしました",
    "kano_8287: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "vinhiso_6763: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "tosuben7068: 友達申請を承認できません: POST /api/friend-requests/b2a4b59b-9da3-4f3c-a7c0-2be25227b694/accept: 応答がタイムアウトしました",
    "kesema_9002: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "warina_9137: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "chimi-8979: 認定公開アカウントをフォローできません: PUT /api/follows/7f09e043-412d-4ae0-8a59-0b18a5cfb839: 応答がタイムアウトしました",
    "honofu8243: 認定公開アカウントをフォローできません: PUT /api/follows/7f09e043-412d-4ae0-8a59-0b18a5cfb839: 応答がタイムアウトしました",
    "heku.9326: 認定公開アカウントをフォローできません: PUT /api/follows/7f09e043-412d-4ae0-8a59-0b18a5cfb839: 応答がタイムアウトしました",
    "rumiyo-6595: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "liahari8038: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "noaelre.9385: 認定公開アカウントをフォローできません: PUT /api/follows/7f09e043-412d-4ae0-8a59-0b18a5cfb839: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.6%	926	1	0	GET	/api/notifications	8568.651	9.253	15.510	15.577	175
14.5%	2572	45	0	POST	/api/friend-requests	3586.424	1.394	16.455	29.302	113
13.3%	1204	25	0	POST	/api/posts	3305.956	2.746	10.483	20.324	635
8.6%	10563	516	0	GET	^/api/images/[^/]+	2142.151	0.203	3.286	10.767	98528
5.9%	4226	870	0	POST	/api/auth/login	1459.718	0.345	3.712	20.265	201
5.0%	973	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1250.401	1.285	16.470	23.202	116
3.9%	299	0	0	PUT	^/api/follows/[^/]+$	967.310	3.235	21.037	29.469	0
3.6%	394	77	0	POST	^/api/posts/[^/]+/realmojis$	886.614	2.250	16.833	19.038	389
2.4%	282	49	0	POST	^/api/posts/[^/]+/comments$	593.404	2.104	16.231	17.110	289
1.4%	1674	2	0	GET	/api/posting-window/current	344.309	0.206	2.893	4.086	237
1.3%	2219	0	0	GET	/api/memories	318.387	0.143	0.680	2.248	14369
1.1%	491	1	0	POST	/api/auth/register	282.593	0.576	2.708	8.419	227
0.9%	375	111	0	PUT	/api/me/profile-image	211.487	0.564	4.714	5.508	7
0.7%	241	0	0	GET	/api/me/export	168.499	0.699	3.204	8.695	1738269
0.7%	6335	0	0	GET	^/api/feeds/[^/]+	165.179	0.026	0.198	2.901	1997
0.4%	592	0	0	POST	/api/auth/logout	105.843	0.179	1.986	12.922	0
0.3%	213	0	0	PATCH	/api/me	76.459	0.359	2.074	2.436	272
0.1%	588	0	0	GET	^/api/memories/[^/]+	36.424	0.062	0.383	2.637	620
0.1%	49	0	0	DELETE	^/api/posts/[^/]+$	32.868	0.671	3.891	3.891	0
0.1%	893	21	0	GET	^/api/posts/me/current$	32.284	0.036	0.550	1.096	634
0.1%	1259	0	0	GET	/api/users	31.506	0.025	0.156	0.299	250
0.1%	90	2	0	PATCH	^/api/posts/[^/]+$	30.581	0.340	3.003	3.003	643
0.1%	21	0	0	PUT	^/api/blocks/[^/]+$	27.312	1.301	14.936	14.936	0
0.1%	5	0	0	DELETE	^/api/follows/[^/]+$	21.376	4.275	17.677	17.677	0
0.1%	968	11	0	GET	^/api/posts/[^/]+$	18.167	0.019	0.140	0.436	629
0.1%	58	0	0	DELETE	^/api/realmojis/[^/]+$	17.534	0.302	1.733	1.733	0
0.1%	50	0	0	DELETE	^/api/comments/[^/]+$	16.966	0.339	3.751	3.751	0
0.1%	1	0	0	POST	/api/initialize	16.278	16.278	16.278	16.278	90
0.1%	708	0	0	GET	^/api/relationships/[^/]+	15.717	0.022	0.161	1.134	836
0.1%	1	0	0	DELETE	^/api/blocks/[^/]+$	14.123	14.123	14.123	14.123	0
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.0%	236	1	0	GET	/api/notifications	1976.250	8.374	12.129	12.254	176
15.0%	590	10	0	POST	/api/friend-requests	872.704	1.479	16.475	21.034	114
13.2%	281	4	0	POST	/api/posts	766.786	2.729	10.947	11.176	630
5.9%	500	0	0	POST	/api/auth/login	346.177	0.692	10.164	14.000	248
4.8%	1619	8	0	GET	^/api/images/[^/]+	277.078	0.171	1.265	1.666	138580
4.5%	72	0	0	PUT	^/api/follows/[^/]+$	263.687	3.662	21.328	21.328	0
4.5%	171	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	262.645	1.536	16.707	18.089	116
3.5%	1290	0	0	GET	^/api/feeds/[^/]+	201.481	0.156	0.379	1.562	1826
2.5%	63	17	0	POST	^/api/posts/[^/]+/realmojis$	147.338	2.339	17.205	17.205	328
2.0%	43	8	0	POST	^/api/posts/[^/]+/comments$	113.938	2.650	18.110	18.110	265
1.8%	387	0	0	GET	/api/posting-window/current	105.825	0.273	3.149	3.248	239
1.5%	358	0	0	GET	/api/memories	88.519	0.247	1.102	1.892	17871
1.1%	81	0	0	POST	/api/auth/logout	64.018	0.790	10.000	10.000	0
0.7%	270	0	0	GET	/api/users	40.079	0.148	0.292	1.025	248
0.7%	28	0	0	GET	/api/me/export	38.847	1.387	3.252	3.252	2109913
0.5%	217	8	0	GET	^/api/posts/me/current$	31.070	0.143	0.265	0.404	616
0.5%	104	0	0	POST	/api/auth/register	29.514	0.284	0.821	0.954	228
0.5%	49	0	0	PATCH	/api/me	27.544	0.562	13.986	13.986	276
0.4%	156	11	0	GET	^/api/posts/[^/]+$	24.007	0.154	0.403	0.453	579
0.4%	52	0	0	PUT	/api/me/profile-image	23.791	0.458	2.254	2.254	0
0.4%	151	0	0	GET	^/api/relationships/[^/]+	23.163	0.153	0.327	1.084	651
0.3%	21	0	0	PATCH	^/api/posts/[^/]+$	19.976	0.951	9.886	9.886	654
0.2%	79	0	0	GET	^/api/memories/[^/]+	14.493	0.183	0.473	0.473	982
0.2%	91	0	0	GET	^/api/users/[^/]+$	14.097	0.155	0.465	0.465	203
0.2%	81	0	0	GET	^/api/posts/[^/]+/realmojis$	12.077	0.149	0.341	0.341	796
0.2%	58	0	0	GET	^/api/posts/[^/]+/comments$	9.439	0.163	0.341	0.341	962
0.1%	54	2	0	GET	/api/me	7.947	0.147	0.554	0.554	295
0.1%	7	0	0	DELETE	^/api/realmojis/[^/]+$	6.044	0.863	2.729	2.729	0
0.1%	4	0	0	PUT	^/api/blocks/[^/]+$	4.296	1.074	2.826	2.826	0
0.1%	6	0	0	DELETE	^/api/comments/[^/]+$	3.495	0.583	2.125	2.125	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.4%	155	0	0	GET	/api/notifications	1244.969	8.032	9.524	9.525	181
13.7%	167	4	0	POST	/api/posts	469.040	2.809	10.552	10.649	646
12.4%	311	2	0	POST	/api/friend-requests	425.947	1.370	15.683	17.245	114
6.7%	127	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	228.875	1.802	16.036	17.203	116
5.1%	1070	9	11	GET	^/api/images/[^/]+	175.226	0.164	0.947	1.105	138220
5.0%	296	0	0	POST	/api/auth/login	170.577	0.576	7.492	9.719	252
4.2%	36	0	0	PUT	^/api/follows/[^/]+$	144.104	4.003	17.286	17.286	0
3.0%	45	9	0	POST	^/api/posts/[^/]+/comments$	101.941	2.265	16.877	16.877	267
2.6%	633	0	0	GET	^/api/feeds/[^/]+	90.342	0.143	0.379	0.415	2186
1.8%	226	0	0	GET	/api/posting-window/current	62.639	0.277	3.234	3.260	240
1.7%	244	0	0	GET	/api/memories	58.487	0.240	0.905	1.934	20432
1.0%	50	0	0	POST	/api/auth/logout	34.090	0.682	9.589	9.589	0
0.9%	27	7	0	POST	^/api/posts/[^/]+/realmojis$	30.951	1.146	4.997	4.997	332
0.7%	16	0	0	GET	/api/me/export	23.369	1.461	2.650	2.650	2527060
0.6%	155	0	0	GET	/api/users	21.014	0.136	0.307	0.401	247
0.6%	132	3	0	GET	^/api/posts/me/current$	19.355	0.147	0.358	0.521	647
0.4%	55	0	0	POST	/api/auth/register	14.266	0.259	1.202	1.202	229
0.4%	15	0	0	PATCH	^/api/posts/[^/]+$	13.097	0.873	2.762	2.762	652
0.4%	32	0	0	PUT	/api/me/profile-image	12.855	0.402	1.973	1.973	0
0.4%	101	5	0	GET	^/api/posts/[^/]+$	12.810	0.127	0.222	0.228	585
0.4%	32	0	0	PATCH	/api/me	12.161	0.380	2.269	2.269	282
0.3%	76	0	0	GET	^/api/relationships/[^/]+	11.550	0.152	0.405	0.405	895
0.3%	39	0	0	GET	^/api/memories/[^/]+	8.831	0.226	0.449	0.449	1219
0.2%	9	0	0	DELETE	^/api/posts/[^/]+$	8.279	0.920	2.800	2.800	0
0.2%	52	0	0	GET	^/api/posts/[^/]+/realmojis$	7.554	0.145	0.342	0.342	734
0.2%	42	0	0	GET	^/api/users/[^/]+$	5.966	0.142	0.444	0.444	203
0.1%	37	0	0	GET	^/api/posts/[^/]+/comments$	5.051	0.137	0.419	0.419	794
0.1%	32	0	0	GET	/api/me	4.413	0.138	0.394	0.394	307
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	3.056	1.019	2.369	2.369	0
0.1%	10	0	0	DELETE	^/api/comments/[^/]+$	1.745	0.174	0.390	0.390	0
…（上位 30 / 全 32 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
未取得: pg_stat_statements を読めませんでした (isu4, exit=2)
isuh setup --only logs で有効にしたか確認すること。

psql: error: connection to server on socket "/var/run/postgresql/.s.PGSQL.5432" failed: FATAL:  sorry, too many clients already
```

## CPU profile (pprof)

```
File: isureal
Build ID: bccf26f46231835381414f798d155fa0085ac095
Type: cpu
Time: 2026-10-31 12:10:37 JST
Duration: 60.20s, Total samples = 45.84s (76.15%)
Showing nodes accounting for 34.45s, 75.15% of 45.84s total
Dropped 967 nodes (cum <= 0.23s)
Showing top 40 nodes out of 278
      flat  flat%   sum%        cum   cum%
     3.05s  6.65%  6.65%      3.05s  6.65%  internal/runtime/syscall/linux.Syscall6
     2.47s  5.39% 12.04%      8.67s 18.91%  runtime.scanObject
     2.41s  5.26% 17.30%      2.41s  5.26%  runtime.memclrNoHeapPointers
     2.21s  4.82% 22.12%      2.21s  4.82%  crypto/internal/fips140/sha256.blockAVX2
     1.99s  4.34% 26.46%      6.82s 14.88%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     1.99s  4.34% 30.80%      2.52s  5.50%  runtime.findObject
     1.86s  4.06% 34.86%      3.24s  7.07%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     1.82s  3.97% 38.83%      1.82s  3.97%  runtime.memmove
     1.65s  3.60% 42.43%      5.72s 12.48%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.63s  3.56% 45.99%      3.14s  6.85%  runtime.tryDeferToSpanScan
     1.53s  3.34% 49.32%      1.53s  3.34%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.50s  3.27% 52.60%      1.57s  3.42%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     1.09s  2.38% 54.97%      1.09s  2.38%  runtime.spanClass.sizeclass (inline)
     1.04s  2.27% 57.24%      1.04s  2.27%  aeshashbody
     1.03s  2.25% 59.49%      1.44s  3.14%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     0.82s  1.79% 61.28%      0.82s  1.79%  example.com/isureal-company/isureal.(*UserRow).active
     0.69s  1.51% 62.78%      0.69s  1.51%  memeqbody
     0.57s  1.24% 64.03%      1.93s  4.21%  runtime.mapaccess1_faststr
     0.37s  0.81% 64.83%      0.40s  0.87%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.33s  0.72% 65.55%      0.71s  1.55%  runtime.greyobject
     0.32s   0.7% 66.25%      0.32s   0.7%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.32s   0.7% 66.95%      0.32s   0.7%  runtime.nextFreeFast (inline)
     0.31s  0.68% 67.63%      0.54s  1.18%  example.com/isureal-company/isureal.postLess
     0.29s  0.63% 68.26%      0.29s  0.63%  indexbody
     0.29s  0.63% 68.89%      0.32s   0.7%  runtime.(*mspan).markBitsForIndex (inline)
     0.27s  0.59% 69.48%      0.31s  0.68%  runtime.spanOf (inline)
     0.24s  0.52% 70.00%      0.24s  0.52%  crypto/internal/fips140/aes/gcm.gcmAesEnc
     0.24s  0.52% 70.53%      1.01s  2.20%  runtime.mallocgcSmallScanNoHeader
     0.23s   0.5% 71.03%      0.23s   0.5%  runtime.futex
     0.22s  0.48% 71.51%      0.27s  0.59%  runtime.typePointers.next
     0.21s  0.46% 71.97%      0.39s  0.85%  compress/flate.(*compressor).deflate
     0.21s  0.46% 72.43%      0.59s  1.29%  database/sql.convertAssignRows
     0.20s  0.44% 72.86%      0.56s  1.22%  runtime.scanObjectsSmall
     0.18s  0.39% 73.25%      0.34s  0.74%  runtime.heapArenaOf (inline)
     0.18s  0.39% 73.65%      6.19s 13.50%  runtime.mallocgc
     0.17s  0.37% 74.02%      1.81s  3.95%  runtime.mapassign_faststr
     0.13s  0.28% 74.30%      2.47s  5.39%  crypto/internal/fips140/sha256.(*Digest).Write
     0.13s  0.28% 74.59%      3.11s  6.78%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.13s  0.28% 74.87%      1.29s  2.81%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.13s  0.28% 75.15%      0.25s  0.55%  internal/runtime/maps.(*Iter).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 62.0% / 100% | 44% | 0.2% | 7 | 199 MB | CPU が飽和 |
| isu2 (2 cpu) | 14.0% / 31% | 0% | 0.9% | 3 | 2886 MB |  |
| isu3 (2 cpu) | 10.3% / 33% | 0% | 0.8% | 6 | 3003 MB | run queue が積もっている |
| isu4 (2 cpu) | 58.0% / 100% | 25% | 20.1% | 34 | 101 MB | IO 待ちが重い |

**isu1（CPU が飽和、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（IO 待ちが重い、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 101% / nginx 15% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 25% / nginx 1% |
| isu3 | isureal 17% / nginx 1% |
| isu4 | postgres 58% / isureal 14% |

## Diff from previous run

スコア: 56314 → — (fail)

`a8d9e7d..249be25`

```
 conf/isu1/home/isucon/env.sh   |  3 +-
 webapp/go/offload.go           | 65 ++++++++++++++++++++++++++++++++++++------
 webapp/go/post_api.go          |  4 +--
 webapp/go/profile_image_api.go |  3 +-
 webapp/go/realmoji_api.go      |  4 +--
 5 files changed, 61 insertions(+), 18 deletions(-)
```
