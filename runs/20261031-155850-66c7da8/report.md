# run 20261031-155850-66c7da8

## Summary

| | |
|---|---|
| score | 171617 |
| status | pass |
| ref | image-etag-gzip |
| sha | 66c7da811fe9 |
| 所要時間 | 206s |
| 前回比 | -11714 (-6.4%) vs 20261031-152137-8c13252 |
| ベスト比 | -20199 (-10.5%) vs 20261031-151339-5a230a4 = 191816 |
| note | image-etag-304-and-gzip-json |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	47	134	392	1070	1072
europe	7	21	66	190	481	481
america	6	15	39	111	182	181
```

```
{"id":"25713","status":"done","portalStatus":"finished","passed":true,"score":171617,"raw":171617,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T07:00:42.212Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T07:00:45.208Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T07:00:48.206Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T07:00:51.238Z level=INFO msg="asia 2 日目: 47 人"
time=2026-10-31T07:00:52.546Z level=INFO msg="user moved" user=tahe.5214 from=asia to=america
time=2026-10-31T07:00:54.206Z level=INFO msg="europe 2 日目: 21 人"
time=2026-10-31T07:00:57.206Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T07:00:57.311Z level=INFO msg="user moved" user=tahe.5214 from=america to=europe
time=2026-10-31T07:01:00.216Z level=INFO msg="asia 3 日目: 134 人"
time=2026-10-31T07:01:01.736Z level=INFO msg="user moved" user=torchiyo5716 from=asia to=europe
time=2026-10-31T07:01:02.462Z level=INFO msg="user moved" user=solhi5843 from=asia to=america
time=2026-10-31T07:01:02.744Z level=INFO msg="user moved" user=torchiyo5716 from=europe to=america
time=2026-10-31T07:01:03.208Z level=INFO msg="europe 3 日目: 66 人"
time=2026-10-31T07:01:03.802Z level=INFO msg="user moved" user=solhi5843 from=america to=asia
time=2026-10-31T07:01:06.207Z level=INFO msg="america 3 日目: 39 人"
time=2026-10-31T07:01:08.496Z level=INFO msg="user moved" user=tahe.5214 from=europe to=america
time=2026-10-31T07:01:09.220Z level=INFO msg="asia 4 日目: 392 人"
time=2026-10-31T07:01:10.114Z level=INFO msg="user moved" user=hovinleo.6732 from=asia to=america
time=2026-10-31T07:01:11.591Z level=INFO msg="user moved" user=hovinleo.6732 from=america to=asia
time=2026-10-31T07:01:12.582Z level=INFO msg="europe 4 日目: 190 人"
time=2026-10-31T07:01:12.601Z level=INFO msg="user moved" user=solhi5843 from=asia to=america
time=2026-10-31T07:01:12.624Z level=INFO msg="user moved" user=somotsu.7182 from=asia to=europe
time=2026-10-31T07:01:14.675Z level=INFO msg="user moved" user=shiivyshi.6791 from=asia to=america
time=2026-10-31T07:01:15.382Z level=INFO msg="america 4 日目: 111 人"
time=2026-10-31T07:01:16.982Z level=INFO msg="user moved" user=nura_8866 from=europe to=america
time=2026-10-31T07:01:17.760Z level=INFO msg="user moved" user=anan.6915 from=asia to=europe
time=2026-10-31T07:01:17.773Z level=INFO msg="user moved" user=nuho.7864 from=asia to=europe
time=2026-10-31T07:01:17.808Z level=INFO msg="user moved" user=rimeme-8055 from=asia to=america
time=2026-10-31T07:01:17.925Z level=INFO msg="user moved" user=torchiyo5716 from=america to=asia
time=2026-10-31T07:01:18.260Z level=INFO msg="asia 5 日目: 1070 人"
time=2026-10-31T07:01:20.946Z level=INFO msg="user moved" user=kome_8818 from=europe to=asia
time=2026-10-31T07:01:22.550Z level=INFO msg="europe 5 日目: 481 人"
time=2026-10-31T07:01:23.870Z level=INFO msg="user moved" user=meraro.9250 from=europe to=asia
time=2026-10-31T07:01:25.456Z level=INFO msg="america 5 日目: 182 人"
time=2026-10-31T07:01:25.524Z level=INFO msg="user moved" user=ridan_9260 from=america to=europe
time=2026-10-31T07:01:34.406Z level=INFO msg="asia 最終日（グランプリ）: 1072 人"
time=2026-10-31T07:01:34.864Z level=INFO msg="europe 最終日（グランプリ）: 481 人"
time=2026-10-31T07:01:34.950Z level=INFO msg="america 最終日（グランプリ）: 181 人"
time=2026-10-31T07:01:42.302Z level=INFO msg="グランプリ: アジア 528 人, ヨーロッパ 126 人, アメリカ 69 人"
{
  "pass": true,
  "score": 171617,
  "messages": [
    "noamaxto_15652: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "teevare15661: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "orku15572: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noahaor.15758: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "moke.16213: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kima_14965: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "fuha15046: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "zoekise.14696: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noano15202: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nune16038: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "fumu.16701: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "madanel9614: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sahane.16763: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "razoe.17072: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hovinleo.6732: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sovinme.16937: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kimax-16629: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nuni_8983: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "kisesu_15540: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "hochinoa15260: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
43.3%	1697	0	0	GET	/api/notifications	15124.852	8.913	16.188	16.201	170
12.0%	6217	56	0	POST	/api/friend-requests	4204.553	0.676	2.477	4.048	114
9.5%	2704	118	0	POST	/api/posts	3327.408	1.231	3.707	7.195	625
4.9%	6496	550	0	POST	/api/auth/login	1707.751	0.263	2.055	4.711	225
4.7%	18934	307	0	GET	^/api/images/[^/]+	1642.285	0.087	0.869	2.640	128116
3.7%	2837	10	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1287.994	0.454	2.267	3.451	116
3.1%	4528	0	0	GET	/api/memories	1096.352	0.242	3.238	6.812	3990
2.7%	15002	0	0	GET	^/api/feeds/[^/]+	934.199	0.062	0.690	1.876	1203
2.4%	752	130	0	PUT	/api/me/profile-image	854.875	1.137	7.758	9.963	3
2.0%	635	14	0	POST	^/api/posts/[^/]+/realmojis$	694.061	1.093	4.811	8.426	466
1.7%	1093	1	0	POST	/api/auth/register	579.319	0.530	2.224	2.288	227
1.4%	587	8	0	POST	^/api/posts/[^/]+/comments$	476.420	0.812	4.110	5.358	338
1.3%	681	0	0	PUT	^/api/follows/[^/]+$	467.152	0.686	2.798	3.029	0
1.1%	568	0	0	PATCH	/api/me	378.094	0.666	2.951	9.049	274
0.8%	961	0	0	POST	/api/auth/logout	287.885	0.300	2.160	3.034	0
0.8%	350	0	0	GET	/api/me/export	285.026	0.814	7.751	10.213	1785220
0.6%	2971	0	0	GET	/api/users	224.741	0.076	0.407	1.191	245
0.6%	3854	0	0	GET	/api/posting-window/current	203.663	0.053	0.313	0.647	237
0.4%	251	4	0	PATCH	^/api/posts/[^/]+$	137.009	0.546	2.164	2.745	646
0.4%	2169	38	0	GET	^/api/posts/me/current$	134.620	0.062	0.342	0.445	642
0.3%	628	2	0	GET	/api/me	121.987	0.194	0.891	1.147	302
0.3%	2293	39	0	GET	^/api/posts/[^/]+$	117.927	0.051	0.314	0.629	627
0.3%	983	0	0	GET	^/api/memories/[^/]+	107.079	0.109	1.544	2.364	389
0.3%	101	0	0	DELETE	^/api/posts/[^/]+$	106.001	1.050	2.503	2.791	0
0.3%	1894	0	0	GET	^/api/relationships/[^/]+	102.527	0.054	0.321	1.181	494
0.2%	104	0	0	DELETE	^/api/realmojis/[^/]+$	77.335	0.744	2.667	3.033	0
0.2%	110	0	0	DELETE	^/api/comments/[^/]+$	57.100	0.519	3.032	3.602	0
0.2%	943	0	0	GET	^/api/posts/[^/]+/realmojis$	55.585	0.059	0.583	1.237	765
0.1%	1120	0	0	GET	^/api/users/[^/]+$	51.671	0.046	0.277	0.421	219
0.1%	789	0	0	GET	^/api/posts/[^/]+/comments$	38.088	0.048	0.306	0.568	834
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.3%	683	2	0	GET	/api/notifications	5269.671	7.715	12.306	12.307	152
12.4%	2339	25	0	POST	/api/friend-requests	1666.408	0.712	2.335	2.839	114
8.1%	978	44	0	POST	/api/posts	1092.648	1.117	3.401	6.421	621
6.6%	1333	0	0	POST	/api/auth/login	879.057	0.659	2.846	4.566	244
4.8%	4158	0	0	GET	^/api/feeds/[^/]+	640.813	0.154	1.785	3.140	1124
4.5%	4125	30	0	GET	^/api/images/[^/]+	599.537	0.145	2.182	3.559	172205
3.4%	478	0	0	POST	/api/auth/register	452.305	0.946	2.511	2.565	228
2.8%	556	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	379.167	0.682	2.393	3.420	115
2.2%	271	0	0	PUT	/api/me/profile-image	297.604	1.098	3.406	6.230	0
2.0%	1023	0	0	GET	/api/memories	263.549	0.258	2.340	4.800	4346
1.8%	1505	0	0	GET	/api/posting-window/current	244.161	0.162	1.882	2.148	239
1.8%	287	0	0	PUT	^/api/follows/[^/]+$	235.451	0.820	2.361	3.100	0
1.7%	215	6	0	POST	^/api/posts/[^/]+/realmojis$	231.203	1.075	5.146	5.715	457
1.4%	212	0	0	PATCH	/api/me	186.877	0.881	2.314	4.221	279
1.3%	1161	0	0	GET	/api/users	168.450	0.145	1.803	2.584	242
1.0%	687	27	0	GET	^/api/posts/me/current$	128.024	0.186	1.893	2.390	625
0.9%	174	2	0	POST	^/api/posts/[^/]+/comments$	121.436	0.698	3.239	4.867	361
0.7%	198	0	0	POST	/api/auth/logout	95.652	0.483	2.543	2.937	0
0.6%	545	0	0	GET	^/api/relationships/[^/]+	84.654	0.155	1.597	2.290	424
0.5%	70	5	0	PATCH	^/api/posts/[^/]+$	62.278	0.890	4.053	4.053	610
0.5%	619	16	0	GET	^/api/posts/[^/]+$	61.861	0.100	0.273	1.610	619
0.3%	48	0	0	DELETE	^/api/realmojis/[^/]+$	43.848	0.913	2.955	2.955	0
0.2%	32	0	0	DELETE	^/api/posts/[^/]+$	32.085	1.003	2.743	2.743	0
0.2%	271	0	0	GET	^/api/users/[^/]+$	27.790	0.103	0.757	1.364	215
0.2%	277	1	0	GET	/api/me	27.450	0.099	0.313	1.087	303
0.2%	76	0	0	GET	/api/me/export	26.331	0.346	7.407	7.407	2531362
0.2%	199	0	0	GET	^/api/posts/[^/]+/comments$	22.314	0.112	0.826	2.254	903
0.2%	166	0	0	GET	^/api/memories/[^/]+	21.938	0.132	2.385	3.022	492
0.2%	225	0	0	GET	^/api/posts/[^/]+/realmojis$	21.487	0.095	0.296	1.089	843
0.1%	27	0	0	DELETE	^/api/comments/[^/]+$	16.321	0.604	1.988	1.988	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.4%	320	1	0	GET	/api/notifications	2417.202	7.554	11.053	11.073	164
10.2%	854	14	0	POST	/api/friend-requests	580.776	0.680	2.410	2.643	114
9.3%	424	19	0	POST	/api/posts	532.788	1.257	6.423	66.793	632
6.5%	739	0	0	POST	/api/auth/login	370.336	0.501	2.146	4.008	250
5.6%	2341	17	0	GET	^/api/images/[^/]+	317.172	0.135	1.419	3.481	139552
4.5%	1776	0	0	GET	^/api/feeds/[^/]+	254.403	0.143	1.184	2.369	1129
3.0%	268	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	170.964	0.638	2.262	2.679	116
2.3%	176	0	0	POST	/api/auth/register	133.653	0.759	2.359	2.403	229
2.3%	586	0	0	GET	/api/memories	128.373	0.219	3.131	4.685	5460
1.6%	133	0	0	PUT	^/api/follows/[^/]+$	92.637	0.697	2.185	2.778	0
1.5%	105	0	0	PUT	/api/me/profile-image	83.679	0.797	2.757	2.771	0
1.3%	619	0	0	GET	/api/posting-window/current	74.333	0.120	1.072	1.211	239
1.1%	107	0	0	POST	/api/auth/logout	62.161	0.581	2.465	5.262	0
1.0%	307	22	0	GET	^/api/posts/me/current$	57.030	0.186	1.209	1.417	617
1.0%	96	0	0	PATCH	/api/me	55.736	0.581	1.950	1.950	286
0.9%	69	4	0	POST	^/api/posts/[^/]+/comments$	50.372	0.730	5.160	5.160	315
0.9%	49	2	0	POST	^/api/posts/[^/]+/realmojis$	49.412	1.008	4.183	4.183	496
0.8%	406	0	0	GET	/api/users	47.817	0.118	0.810	1.268	249
0.5%	245	12	0	GET	^/api/posts/[^/]+$	28.989	0.118	0.843	1.263	607
0.5%	34	0	0	PATCH	^/api/posts/[^/]+$	28.776	0.846	2.157	2.157	657
0.5%	220	0	0	GET	^/api/relationships/[^/]+	27.035	0.123	0.898	1.101	472
0.4%	51	0	0	GET	/api/me/export	25.051	0.491	5.084	5.084	2817587
0.4%	18	0	0	DELETE	^/api/posts/[^/]+$	22.225	1.235	2.444	2.444	0
0.3%	109	1	0	GET	/api/me	15.211	0.140	1.169	1.170	301
0.3%	138	0	0	GET	^/api/users/[^/]+$	15.193	0.110	0.797	1.054	212
0.3%	112	0	0	GET	^/api/memories/[^/]+	14.912	0.133	1.746	1.788	610
0.2%	101	0	0	GET	^/api/posts/[^/]+/comments$	10.037	0.099	0.660	0.834	870
0.2%	9	0	0	DELETE	^/api/realmojis/[^/]+$	9.228	1.025	1.854	1.854	0
0.2%	11	0	0	DELETE	^/api/comments/[^/]+$	9.051	0.823	2.528	2.528	0
0.1%	73	0	0	GET	^/api/posts/[^/]+/realmojis$	7.258	0.099	0.839	0.839	823
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 15780 クエリ / 48 種類 / 合計 141.355s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	52.2%	362	73.814	203.90	1422.00	1149.7	117.6	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	15.1%	3	21.399	7133.10	9529.58	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.6%	13221	16.450	1.24	9.56	9.5	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	11.4%	3	16.074	5357.87	7232.09	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.5%	3	7.785	2595.03	3541.20	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.5%	4	3.584	896.04	3015.15	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.3%	4	0.365	91.14	156.60	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
8	0.3%	4	0.361	90.30	156.97	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.2%	1	0.273	272.78	272.78	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.258	258.18	258.18	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
11	0.2%	1	0.245	244.67	244.67	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
12	0.1%	1	0.166	165.96	165.96	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.108	108.12	108.12	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.081	80.93	80.93	2463.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.1%	1	0.076	76.34	76.34	13252.0	0.0	DELETE FROM comments WHERE created_at > (SELECT created_at FROM initial_data_boundary)

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
Build ID: 0c4551499267e68505ad8ee7adf1fddebb494575
Type: cpu
Time: 2026-10-31 16:00:21 JST
Duration: 60.03s, Total samples = 41.12s (68.50%)
Showing nodes accounting for 29.80s, 72.47% of 41.12s total
Dropped 1136 nodes (cum <= 0.21s)
Showing top 40 nodes out of 280
      flat  flat%   sum%        cum   cum%
     5.65s 13.74% 13.74%      5.65s 13.74%  crypto/internal/fips140/sha256.blockAVX2
     4.59s 11.16% 24.90%      4.59s 11.16%  internal/runtime/syscall/linux.Syscall6
     2.03s  4.94% 29.84%      2.03s  4.94%  runtime.memmove
     1.59s  3.87% 33.71%      5.88s 14.30%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.56s  3.79% 37.50%      1.56s  3.79%  runtime.memclrNoHeapPointers
     1.53s  3.72% 41.22%      1.53s  3.72%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.41s  3.43% 44.65%      1.41s  3.43%  aeshashbody
     1.35s  3.28% 47.93%      1.83s  4.45%  runtime.findObject
     1.29s  3.14% 51.07%      2.22s  5.40%  runtime.tryDeferToSpanScan
     1.19s  2.89% 53.96%      5.50s 13.38%  runtime.scanObject
     0.66s  1.61% 55.57%      2.19s  5.33%  runtime.mapaccess1_faststr
     0.65s  1.58% 57.15%      0.65s  1.58%  runtime.spanClass.sizeclass (inline)
     0.45s  1.09% 58.24%      0.45s  1.09%  memeqbody
     0.43s  1.05% 59.29%      0.61s  1.48%  example.com/isureal-company/isureal.postLess
     0.40s  0.97% 60.26%      0.40s  0.97%  runtime.nextFreeFast (inline)
     0.35s  0.85% 61.11%      0.35s  0.85%  runtime.futex
     0.30s  0.73% 61.84%      6.24s 15.18%  crypto/internal/fips140/sha256.(*Digest).Write
     0.30s  0.73% 62.57%      6.33s 15.39%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.26s  0.63% 63.21%      0.88s  2.14%  database/sql.convertAssignRows
     0.26s  0.63% 63.84%      2.19s  5.33%  runtime.mapassign_faststr
     0.26s  0.63% 64.47%      0.26s  0.63%  sync/atomic.(*Int32).Add
     0.25s  0.61% 65.08%      0.27s  0.66%  encoding/json.appendString[go.shape.string]
     0.24s  0.58% 65.66%      0.27s  0.66%  runtime.spanOf (inline)
     0.23s  0.56% 66.22%      0.23s  0.56%  runtime.typePointers.nextFast (inline)
     0.22s  0.54% 66.76%      0.26s  0.63%  runtime.typePointers.next
     0.20s  0.49% 67.24%      0.24s  0.58%  runtime.(*mspan).writeHeapBitsSmall
     0.20s  0.49% 67.73%      4.62s 11.24%  runtime.mallocgc
     0.19s  0.46% 68.19%      6.99s 17.00%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.19s  0.46% 68.65%      1.19s  2.89%  runtime.mallocgcSmallScanNoHeader
     0.18s  0.44% 69.09%      0.27s  0.66%  internal/runtime/maps.(*Iter).Next
     0.18s  0.44% 69.53%      0.49s  1.19%  runtime.scanObjectsSmall
     0.17s  0.41% 69.94%      6.64s 16.15%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.16s  0.39% 70.33%      0.22s  0.54%  runtime.(*mspan).markBitsForIndex (inline)
     0.16s  0.39% 70.72%      0.41s     1%  runtime.greyobject
     0.16s  0.39% 71.11%      0.51s  1.24%  runtime.mapaccess2_faststr
     0.13s  0.32% 71.43%      0.35s  0.85%  compress/flate.(*compressor).deflate
     0.13s  0.32% 71.74%      1.30s  3.16%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.11s  0.27% 72.01%      3.07s  7.47%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.24% 72.25%      0.21s  0.51%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.09s  0.22% 72.47%      0.97s  2.36%  database/sql.(*Rows).scanLocked
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 56.7% / 100% | 36% | 8.9% | 9 | 114 MB | CPU が飽和 |
| isu2 (2 cpu) | 31.1% / 100% | 6% | 0.7% | 5 | 316 MB | run queue が積もっている |
| isu3 (2 cpu) | 23.0% / 99% | 2% | 0.5% | 6 | 779 MB | run queue が積もっている |
| isu4 (2 cpu) | 39.5% / 100% | 12% | 15.8% | 8 | 111 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu3（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 68% / nginx 15% / find 2% |
| isu2 | isureal 58% / nginx 2% |
| isu3 | isureal 44% / nginx 1% |
| isu4 | isureal 30% / postgres 15% |

## Diff from previous run

スコア: 183331 → 171617 (-11714)

`8c13252..66c7da8`

```
 conf/isu1/etc/nginx/nginx.conf |  5 +++++
 conf/isu2/etc/nginx/nginx.conf |  5 +++++
 conf/isu3/etc/nginx/nginx.conf |  5 +++++
 webapp/go/admin_api.go         | 27 ++++++++++++++++++-----
 webapp/go/export_api.go        | 12 ++++++++--
 webapp/go/image_api.go         | 36 +++++++++++++++++++++---------
 webapp/go/image_file.go        | 33 ++++++++++++++++------------
 webapp/go/image_held.go        | 10 +++------
 webapp/go/image_local.go       | 50 +++++++++++++++++++-----------------------
 webapp/go/ip_ban.go            | 11 ++++++++++
 webapp/go/persist.go           |  2 +-
 webapp/go/realmoji_api.go      | 41 ++++++++++++++--------------------
 webapp/go/repl.go              |  1 +
 webapp/go/replica_server.go    | 13 +++++------
 webapp/go/store.go             |  3 +--
 15 files changed, 152 insertions(+), 102 deletions(-)
```
