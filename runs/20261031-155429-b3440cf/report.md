# run 20261031-155429-b3440cf

## Summary

| | |
|---|---|
| score | 184830 |
| status | pass |
| ref | b3440cf |
| sha | b3440cf1f4d9 |
| 所要時間 | 346s |
| 前回比 | +14266 (8.4%) vs 20261031-155036-b19d232 |
| ベスト比 | -6986 (-3.6%) vs 20261031-151339-5a230a4 = 191816 |
| note | main b3440cf baseline (presign lane) |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	135	390	1036	1037
europe	7	21	63	187	510	509
america	6	14	40	118	249	249
```

```
{"id":"25530","status":"done","portalStatus":"finished","passed":true,"score":184830,"raw":184830,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T06:58:40.128Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:58:43.123Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:58:46.123Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:58:49.136Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T06:58:52.124Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T06:58:55.125Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T06:58:58.133Z level=INFO msg="asia 3 日目: 135 人"
time=2026-10-31T06:58:59.174Z level=INFO msg="user moved" user=riko.5607 from=asia to=america
time=2026-10-31T06:59:01.127Z level=INFO msg="europe 3 日目: 63 人"
time=2026-10-31T06:59:01.180Z level=INFO msg="user moved" user=midan_6016 from=asia to=europe
time=2026-10-31T06:59:03.300Z level=INFO msg="user moved" user=riko.5607 from=america to=europe
time=2026-10-31T06:59:04.124Z level=INFO msg="america 3 日目: 40 人"
time=2026-10-31T06:59:04.677Z level=INFO msg="user moved" user=yore6055 from=europe to=america
time=2026-10-31T06:59:07.137Z level=INFO msg="asia 4 日目: 390 人"
time=2026-10-31T06:59:10.506Z level=INFO msg="europe 4 日目: 187 人"
time=2026-10-31T06:59:10.691Z level=INFO msg="user moved" user=roke7115 from=asia to=europe
time=2026-10-31T06:59:12.299Z level=INFO msg="user moved" user=nusol8128 from=asia to=america
time=2026-10-31T06:59:12.700Z level=INFO msg="user moved" user=evaya_8349 from=asia to=europe
time=2026-10-31T06:59:12.878Z level=INFO msg="user moved" user=dankuwa8342 from=asia to=america
time=2026-10-31T06:59:13.140Z level=INFO msg="america 4 日目: 118 人"
time=2026-10-31T06:59:13.223Z level=INFO msg="user moved" user=nusol8128 from=america to=europe
time=2026-10-31T06:59:15.115Z level=INFO msg="user moved" user=yoleo.7620 from=asia to=america
time=2026-10-31T06:59:15.134Z level=INFO msg="user moved" user=kekaro.7732 from=asia to=america
time=2026-10-31T06:59:15.714Z level=INFO msg="user moved" user=kuwaso-8913 from=europe to=asia
time=2026-10-31T06:59:16.057Z level=INFO msg="user moved" user=benfu.8629 from=europe to=asia
time=2026-10-31T06:59:16.163Z level=INFO msg="asia 5 日目: 1036 人"
time=2026-10-31T06:59:20.511Z level=INFO msg="europe 5 日目: 510 人"
time=2026-10-31T06:59:20.830Z level=INFO msg="user moved" user=zoemaxke-8463 from=europe to=asia
time=2026-10-31T06:59:20.859Z level=INFO msg="user moved" user=suul_9193 from=europe to=asia
time=2026-10-31T06:59:21.853Z level=INFO msg="user moved" user=yore6055 from=america to=europe
time=2026-10-31T06:59:23.485Z level=INFO msg="america 5 日目: 249 人"
time=2026-10-31T06:59:31.293Z level=INFO msg="america 最終日（グランプリ）: 249 人"
time=2026-10-31T06:59:31.384Z level=INFO msg="asia 最終日（グランプリ）: 1037 人"
time=2026-10-31T06:59:31.805Z level=INFO msg="europe 最終日（グランプリ）: 509 人"
time=2026-10-31T06:59:40.202Z level=INFO msg="グランプリ: アジア 492 人, ヨーロッパ 208 人, アメリカ 147 人"
{
  "pass": true,
  "score": 184830,
  "messages": [
    "rokaho.6315: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "razoe.17088: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rilia_12583: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kodanho.10167: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ivymu_14077: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "riketa_15883: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "evamuchi.15135: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "niso-15126: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "marel_16556: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hine-15689: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "shinamar.16535: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "keivy-16356: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "zoeya.15769: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ruan.16373: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "tsuru-15095: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "homia15228: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "solso6901: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "soku10371: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ismi-14006: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "zoero5674: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.4%	1681	0	0	GET	/api/notifications	14745.553	8.772	15.302	15.318	168
10.9%	2643	93	0	POST	/api/posts	3722.062	1.408	5.718	65.651	632
9.9%	5902	73	0	POST	/api/friend-requests	3357.761	0.569	2.198	3.239	114
5.9%	19942	310	0	GET	^/api/images/[^/]+	1995.131	0.100	1.155	3.546	123008
5.0%	6838	529	0	POST	/api/auth/login	1714.831	0.251	1.937	3.904	227
3.7%	3134	11	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1259.560	0.402	2.255	3.382	116
2.7%	4668	0	0	GET	/api/memories	926.592	0.198	1.932	4.609	12981
2.7%	15176	0	0	GET	^/api/feeds/[^/]+	902.003	0.059	0.513	1.689	3724
2.3%	794	16	0	POST	^/api/posts/[^/]+/realmojis$	765.309	0.964	4.246	5.338	468
1.8%	727	125	0	PUT	/api/me/profile-image	600.545	0.826	8.079	10.075	3
1.7%	1058	1	0	POST	/api/auth/register	586.862	0.555	1.908	2.079	227
1.6%	837	3	0	POST	^/api/posts/[^/]+/comments$	546.596	0.653	3.278	4.056	349
1.3%	760	0	0	PUT	^/api/follows/[^/]+$	441.899	0.581	2.255	2.629	0
0.9%	1146	0	0	POST	/api/auth/logout	304.407	0.266	2.125	2.932	0
0.9%	398	0	0	GET	/api/me/export	296.028	0.744	5.653	6.064	1873603
0.9%	524	0	0	PATCH	/api/me	295.966	0.565	2.350	3.348	277
0.8%	3781	0	0	GET	/api/posting-window/current	255.332	0.068	0.444	0.869	237
0.6%	2838	0	0	GET	/api/users	187.897	0.066	0.409	1.416	244
0.5%	2095	5	0	GET	^/api/posts/me/current$	158.186	0.076	0.573	1.397	653
0.4%	226	1	0	PATCH	^/api/posts/[^/]+$	129.382	0.572	2.279	2.940	646
0.3%	2440	25	0	GET	^/api/posts/[^/]+$	114.130	0.047	0.365	1.405	628
0.3%	617	3	0	GET	/api/me	104.670	0.170	1.111	1.176	301
0.3%	1925	0	0	GET	^/api/relationships/[^/]+	104.007	0.054	0.467	1.424	952
0.3%	100	0	0	DELETE	^/api/posts/[^/]+$	93.309	0.933	2.675	2.694	0
0.2%	1015	0	0	GET	^/api/memories/[^/]+	79.924	0.079	0.513	2.616	613
0.2%	115	0	0	DELETE	^/api/realmojis/[^/]+$	74.037	0.644	2.175	2.613	0
0.2%	138	0	0	DELETE	^/api/comments/[^/]+$	55.911	0.405	3.191	3.353	0
0.2%	1203	0	0	GET	^/api/users/[^/]+$	52.753	0.044	0.302	1.386	207
0.1%	1099	0	0	GET	^/api/posts/[^/]+/realmojis$	45.991	0.042	0.288	0.550	1837
0.1%	931	0	0	GET	^/api/posts/[^/]+/comments$	35.764	0.038	0.223	0.491	2147
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.5%	752	0	0	GET	/api/notifications	5123.206	6.813	11.283	11.461	148
12.0%	2451	22	0	POST	/api/friend-requests	1562.936	0.638	2.456	3.232	114
8.9%	1053	48	0	POST	/api/posts	1156.786	1.099	3.503	6.993	621
5.8%	1442	0	0	POST	/api/auth/login	755.737	0.524	2.197	4.177	245
5.2%	4362	31	0	GET	^/api/images/[^/]+	669.150	0.153	1.867	3.526	162460
4.4%	3715	0	0	GET	^/api/feeds/[^/]+	566.282	0.152	1.817	2.831	3184
3.6%	506	0	0	POST	/api/auth/register	465.799	0.921	2.230	2.364	229
2.5%	286	0	0	PUT	/api/me/profile-image	318.734	1.114	3.182	4.091	0
2.3%	358	0	0	PUT	^/api/follows/[^/]+$	298.262	0.833	2.542	2.927	0
2.1%	485	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	267.255	0.551	2.330	2.834	116
1.8%	1638	0	0	GET	/api/posting-window/current	238.808	0.146	1.758	2.364	239
1.8%	245	0	0	PATCH	/api/me	234.654	0.958	2.834	3.917	278
1.4%	1110	0	0	GET	/api/memories	182.942	0.165	1.948	3.390	13210
1.4%	1146	0	0	GET	/api/users	180.219	0.157	1.623	2.038	243
1.2%	767	40	0	GET	^/api/posts/me/current$	160.780	0.210	2.155	2.571	616
1.2%	210	10	0	POST	^/api/posts/[^/]+/realmojis$	157.678	0.751	3.099	5.892	468
0.9%	184	2	0	POST	^/api/posts/[^/]+/comments$	118.996	0.647	3.059	3.604	345
0.7%	202	0	0	POST	/api/auth/logout	91.312	0.452	2.253	2.925	0
0.7%	503	0	0	GET	^/api/relationships/[^/]+	84.810	0.169	1.796	2.406	798
0.5%	89	10	0	PATCH	^/api/posts/[^/]+$	64.412	0.724	2.907	2.907	573
0.4%	533	15	0	GET	^/api/posts/[^/]+$	57.369	0.108	0.368	0.924	610
0.3%	291	1	0	GET	/api/me	35.553	0.122	0.730	0.732	304
0.3%	35	0	0	DELETE	^/api/posts/[^/]+$	35.536	1.015	3.243	3.243	0
0.2%	228	0	0	GET	^/api/users/[^/]+$	27.930	0.122	0.989	1.927	202
0.2%	34	0	0	DELETE	^/api/realmojis/[^/]+$	24.701	0.726	2.245	2.245	0
0.2%	228	0	0	GET	^/api/posts/[^/]+/realmojis$	22.459	0.099	0.266	0.279	2110
0.2%	178	0	0	GET	^/api/memories/[^/]+	22.142	0.124	1.531	3.314	911
0.1%	173	0	0	GET	^/api/posts/[^/]+/comments$	17.201	0.099	0.268	0.313	2858
0.1%	29	0	0	DELETE	^/api/comments/[^/]+$	14.092	0.486	2.263	2.263	0
0.1%	67	0	0	GET	/api/me/export	8.885	0.133	0.834	0.834	2568101
…（上位 30 / 全 35 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.6%	389	1	0	GET	/api/notifications	2741.745	7.048	10.386	10.386	149
10.4%	1293	18	0	POST	/api/friend-requests	684.681	0.530	2.091	2.851	114
7.7%	907	0	0	POST	/api/auth/login	510.022	0.562	2.121	2.894	250
7.5%	538	19	0	POST	/api/posts	495.046	0.920	2.877	7.003	635
6.4%	2897	11	0	GET	^/api/images/[^/]+	422.124	0.146	1.303	1.838	133621
4.4%	1981	0	0	GET	^/api/feeds/[^/]+	288.727	0.146	1.074	1.887	3489
2.8%	245	0	0	POST	/api/auth/register	182.833	0.746	2.079	2.104	229
2.4%	285	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	160.941	0.565	2.066	3.536	116
2.1%	765	0	0	GET	/api/memories	137.289	0.179	1.693	3.839	17144
1.8%	138	0	0	PUT	/api/me/profile-image	120.252	0.871	3.117	3.264	0
1.5%	419	16	0	GET	^/api/posts/me/current$	101.608	0.243	1.491	1.901	634
1.3%	114	5	0	POST	^/api/posts/[^/]+/comments$	88.638	0.778	2.820	2.952	354
1.3%	791	0	0	GET	/api/posting-window/current	87.364	0.110	0.913	1.648	240
1.3%	85	8	0	POST	^/api/posts/[^/]+/realmojis$	86.428	1.017	5.479	5.479	422
1.3%	114	0	0	PATCH	/api/me	84.285	0.739	2.628	2.693	282
1.1%	650	0	0	GET	/api/users	74.588	0.115	1.022	1.335	245
1.0%	117	0	0	PUT	^/api/follows/[^/]+$	68.331	0.584	1.996	2.059	0
0.9%	125	0	0	POST	/api/auth/logout	58.393	0.467	1.998	2.185	0
0.6%	270	0	0	GET	^/api/relationships/[^/]+	36.681	0.136	1.040	1.536	786
0.4%	40	1	0	PATCH	^/api/posts/[^/]+$	26.506	0.663	2.032	2.032	653
0.4%	252	6	0	GET	^/api/posts/[^/]+$	23.131	0.092	0.305	0.938	615
0.3%	61	0	0	GET	/api/me/export	19.758	0.324	4.572	4.572	3259200
0.3%	14	0	0	DELETE	^/api/posts/[^/]+$	18.526	1.323	2.177	2.177	0
0.2%	142	1	0	GET	/api/me	12.123	0.085	0.293	0.774	302
0.2%	109	0	0	GET	^/api/users/[^/]+$	11.693	0.107	0.753	0.905	207
0.2%	106	0	0	GET	^/api/memories/[^/]+	10.693	0.101	0.388	1.409	994
0.2%	103	0	0	GET	^/api/posts/[^/]+/realmojis$	10.255	0.100	0.298	1.571	1968
0.1%	11	0	0	DELETE	^/api/realmojis/[^/]+$	9.391	0.854	2.189	2.189	0
0.1%	78	0	0	GET	^/api/posts/[^/]+/comments$	9.177	0.118	1.202	1.202	2958
0.1%	13	0	0	DELETE	^/api/comments/[^/]+$	8.257	0.635	1.582	1.582	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16986 クエリ / 48 種類 / 合計 155.199s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	55.5%	406	86.101	212.07	1729.14	1192.0	121.9	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	13.7%	3	21.203	7067.80	9514.60	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.9%	14327	18.517	1.29	16.19	9.3	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.3%	3	15.937	5312.43	7172.51	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.1%	3	7.988	2662.54	3774.06	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.1%	4	3.210	802.46	2721.58	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.2%	4	0.375	93.82	157.55	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
8	0.2%	4	0.367	91.78	157.11	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.2%	1	0.256	255.68	255.68	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
10	0.2%	1	0.254	253.76	253.76	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.1%	1	0.232	232.32	232.32	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.163	163.39	163.39	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.109	109.04	109.04	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.082	81.86	81.86	2486.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.074	73.91	73.91	1551.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
Build ID: c7b4d435f98f1ef8930900a3c752289320a0a85d
Type: cpu
Time: 2026-10-31 15:58:18 JST
Duration: 60.28s, Total samples = 42.16s (69.94%)
Showing nodes accounting for 30.47s, 72.27% of 42.16s total
Dropped 1061 nodes (cum <= 0.21s)
Showing top 40 nodes out of 272
      flat  flat%   sum%        cum   cum%
     6.33s 15.01% 15.01%      6.33s 15.01%  crypto/internal/fips140/sha256.blockAVX2
     4.20s  9.96% 24.98%      4.20s  9.96%  internal/runtime/syscall/linux.Syscall6
     1.80s  4.27% 29.25%      1.80s  4.27%  runtime.memmove
     1.60s  3.80% 33.04%      5.42s 12.86%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.58s  3.75% 36.79%      1.98s  4.70%  runtime.findObject
     1.50s  3.56% 40.35%      1.50s  3.56%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.48s  3.51% 43.86%      6.10s 14.47%  runtime.scanObject
     1.47s  3.49% 47.34%      1.47s  3.49%  runtime.memclrNoHeapPointers
     1.44s  3.42% 50.76%      1.44s  3.42%  aeshashbody
     1.19s  2.82% 53.58%      2.34s  5.55%  runtime.tryDeferToSpanScan
     0.77s  1.83% 55.41%      0.77s  1.83%  runtime.spanClass.sizeclass (inline)
     0.55s  1.30% 56.71%      2.27s  5.38%  runtime.mapaccess1_faststr
     0.44s  1.04% 57.76%      0.61s  1.45%  example.com/isureal-company/isureal.postLess
     0.42s     1% 58.75%      0.42s     1%  runtime.nextFreeFast (inline)
     0.40s  0.95% 59.70%      0.80s  1.90%  database/sql.convertAssignRows
     0.35s  0.83% 60.53%      7.22s 17.13%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.35s  0.83% 61.36%      0.35s  0.83%  memeqbody
     0.32s  0.76% 62.12%      0.32s  0.76%  runtime.futex
     0.31s  0.74% 62.86%      6.99s 16.58%  crypto/internal/fips140/sha256.(*Digest).Write
     0.31s  0.74% 63.59%      0.31s  0.74%  sync/atomic.(*Int32).Add
     0.28s  0.66% 64.26%      1.03s  2.44%  runtime.mallocgcSmallScanNoHeader
     0.25s  0.59% 64.85%      0.32s  0.76%  runtime.(*mspan).markBitsForIndex (inline)
     0.24s  0.57% 65.42%      8.08s 19.17%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.23s  0.55% 65.96%      0.42s     1%  compress/flate.(*compressor).deflate
     0.23s  0.55% 66.51%      0.24s  0.57%  encoding/json.appendString[go.shape.string]
     0.23s  0.55% 67.05%      0.24s  0.57%  runtime.(*mspan).writeHeapBitsSmall
     0.23s  0.55% 67.60%      0.64s  1.52%  runtime.greyobject
     0.22s  0.52% 68.12%      7.61s 18.05%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.22s  0.52% 68.64%      0.35s  0.83%  internal/runtime/maps.(*Iter).Next
     0.22s  0.52% 69.17%      4.29s 10.18%  runtime.mallocgc
     0.20s  0.47% 69.64%      0.24s  0.57%  runtime.heapArenaOf (inline)
     0.18s  0.43% 70.07%      0.53s  1.26%  runtime.scanObjectsSmall
     0.16s  0.38% 70.45%      1.87s  4.44%  runtime.mapassign_faststr
     0.15s  0.36% 70.80%      0.27s  0.64%  runtime.scanObjectSmall
     0.12s  0.28% 71.09%      0.45s  1.07%  runtime.mapaccess2_faststr
     0.11s  0.26% 71.35%      1.07s  2.54%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.10s  0.24% 71.58%      6.43s 15.25%  crypto/internal/fips140/sha256.block
     0.10s  0.24% 71.82%      0.77s  1.83%  encoding/json.structEncoder.encode
     0.10s  0.24% 72.06%      1.02s  2.42%  runtime.mallocgcSmallNoscan
     0.09s  0.21% 72.27%      2.75s  6.52%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 46.0% / 100% | 29% | 0.6% | 5 | 113 MB | run queue が積もっている |
| isu2 (2 cpu) | 30.5% / 100% | 6% | 0.6% | 5 | 130 MB | run queue が積もっている |
| isu3 (2 cpu) | 24.5% / 100% | 3% | 0.6% | 5 | 631 MB | run queue が積もっている |
| isu4 (2 cpu) | 40.2% / 100% | 13% | 16.6% | 5 | 100 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 73% / nginx 13% |
| isu2 | isureal 57% / nginx 2% |
| isu3 | isureal 47% / nginx 1% |
| isu4 | isureal 30% / postgres 15% |

## Diff from previous run

スコア: 170564 → 184830 (+14266)

`b19d232..b3440cf`

```
 webapp/go/image_signature.go |  67 ++++++---------------------
 webapp/go/offload.go         | 106 +++++++++----------------------------------
 webapp/go/server.go          |   1 -
 3 files changed, 37 insertions(+), 137 deletions(-)
```
