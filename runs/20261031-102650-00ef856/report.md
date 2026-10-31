# run 20261031-102650-00ef856

## Summary

| | |
|---|---|
| score | 5421 |
| status | pass |
| ref | db-indexes-pgconf |
| sha | 00ef856a3021 |
| 所要時間 | 156s |
| 前回比 | +2747 (102.7%) vs 20261031-100914-e240203 |
| ベスト比 | +2747 (102.7%) vs 20261031-100914-e240203 = 2674 **ベスト更新** |
| note | インデックス追加とusers.username一意制約 |

## Bench

```
{"id":"752","status":"done","portalStatus":"finished","passed":true,"score":5421,"raw":5421,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:27:42.652Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:27:42.716Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:27:42.791Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:27:51.627Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T01:27:52.987Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T01:27:55.823Z level=INFO msg="america 2 日目: 6 人"
time=2026-10-31T01:27:59.158Z level=INFO msg="asia 3 日目: 18 人"
time=2026-10-31T01:27:59.220Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T01:28:00.437Z level=INFO msg="america 3 日目: 6 人"
time=2026-10-31T01:28:06.778Z level=INFO msg="asia 4 日目: 18 人"
time=2026-10-31T01:28:09.392Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T01:28:11.223Z level=INFO msg="america 4 日目: 6 人"
time=2026-10-31T01:28:14.526Z level=INFO msg="asia 5 日目: 18 人"
time=2026-10-31T01:28:16.243Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T01:28:18.302Z level=INFO msg="america 5 日目: 6 人"
time=2026-10-31T01:28:30.482Z level=INFO msg="asia 最終日（グランプリ）: 17 人"
time=2026-10-31T01:28:31.606Z level=INFO msg="america 最終日（グランプリ）: 6 人"
time=2026-10-31T01:28:31.978Z level=INFO msg="europe 最終日（グランプリ）: 6 人"
time=2026-10-31T01:28:36.209Z level=INFO msg="グランプリ: アジア 16 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 5421,
  "messages": []
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
56.2%	110	1	0	GET	/api/notifications	1108.145	10.074	16.558	38.030	176
10.9%	1566	219	0	GET	^/api/images/[^/]+	215.178	0.137	0.435	0.547	79602
3.9%	693	0	0	GET	^/api/feeds/[^/]+	76.995	0.111	0.410	0.628	1203
3.9%	1083	674	0	POST	/api/auth/login	76.536	0.071	0.284	0.424	114
3.9%	363	0	0	GET	/api/memories	76.332	0.210	0.891	1.023	12712
3.8%	90	0	0	POST	/api/posts	75.271	0.836	6.166	6.166	604
2.9%	18	0	0	POST	/api/admin/posting-window	58.103	3.228	6.443	6.443	176
2.4%	29	0	0	GET	/api/me/export	47.715	1.645	5.747	5.747	2065359
2.3%	34	1	0	POST	/api/auth/register	45.304	1.332	5.197	5.197	237
2.3%	117	108	7	PUT	/api/me/profile-image	44.673	0.382	2.987	4.567	29
2.1%	101	10	0	POST	/api/friend-requests	41.315	0.409	3.638	5.218	107
1.1%	74	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	22.080	0.298	2.767	2.767	116
0.6%	132	3	0	GET	^/api/posts/[^/]+$	12.383	0.094	0.355	0.393	587
0.5%	98	0	0	GET	/api/posting-window/current	9.509	0.097	0.310	0.310	237
0.4%	42	0	0	GET	/api/users	7.551	0.180	0.365	0.365	265
0.4%	91	1	0	GET	^/api/posts/me/current$	7.450	0.082	0.350	0.350	597
0.4%	13	0	0	PUT	^/api/follows/[^/]+$	7.196	0.554	2.646	2.646	0
0.3%	1	0	0	POST	/api/initialize	5.876	5.876	5.876	5.876	69
0.3%	109	0	0	GET	^/api/relationships/[^/]+	5.587	0.051	0.190	0.214	2255
0.3%	86	0	0	GET	^/api/users/[^/]+$	5.024	0.058	0.301	0.301	171
0.2%	68	0	0	GET	^/api/memories/[^/]+	3.675	0.054	0.280	0.280	729
0.2%	2	0	0	DELETE	^/api/follows/[^/]+$	3.391	1.696	1.792	1.792	0
0.2%	33	0	0	GET	^/api/posts/[^/]+/comments$	3.385	0.103	0.277	0.277	37
0.2%	85	0	0	POST	/api/auth/logout	3.220	0.038	0.249	0.249	0
0.2%	5	0	0	POST	^/api/posts/[^/]+/comments$	2.963	0.593	1.425	1.425	272
0.1%	20	0	0	GET	^/api/posts/[^/]+/realmojis$	2.105	0.105	0.308	0.308	29
0.1%	13	0	0	PATCH	^/api/posts/[^/]+$	1.731	0.133	0.273	0.273	615
0.0%	1	0	0	DELETE	^/api/friendships/[^/]+$	0.887	0.887	0.887	0.887	0
0.0%	6	0	0	DELETE	^/api/posts/[^/]+$	0.687	0.115	0.222	0.222	0
0.0%	6	0	0	PUT	^/api/blocks/[^/]+$	0.409	0.068	0.188	0.188	0
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.5%	42	0	0	GET	/api/notifications	422.531	10.060	17.162	17.162	177
10.5%	27	0	0	POST	/api/posts	97.947	3.628	6.013	6.013	627
9.3%	99	3	0	GET	^/api/images/[^/]+	86.646	0.875	1.447	1.447	180472
6.4%	45	0	0	GET	^/api/feeds/[^/]+	59.673	1.326	2.295	2.295	411
6.2%	33	0	0	POST	/api/friend-requests	57.860	1.753	2.723	2.723	115
5.5%	54	0	0	GET	/api/memories	50.829	0.941	1.946	1.946	7251
4.2%	27	0	0	POST	/api/auth/login	38.623	1.430	1.971	1.971	256
1.9%	20	0	0	GET	/api/users	18.041	0.902	1.387	1.387	200
1.8%	33	0	0	GET	/api/posting-window/current	16.667	0.505	1.512	1.512	239
1.8%	4	0	0	PUT	^/api/follows/[^/]+$	16.639	4.160	5.527	5.527	0
1.4%	23	7	0	GET	^/api/posts/me/current$	13.056	0.568	0.826	0.826	449
1.3%	5	0	0	POST	/api/auth/register	12.476	2.495	3.915	3.915	230
1.1%	4	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	10.638	2.660	4.453	4.453	116
0.6%	14	0	0	GET	^/api/relationships/[^/]+	5.789	0.413	0.609	0.609	1219
0.5%	3	0	0	DELETE	^/api/posts/[^/]+$	4.724	1.575	1.744	1.744	0
0.4%	4	0	0	POST	/api/auth/logout	3.408	0.852	1.032	1.032	0
0.4%	2	0	0	GET	/api/me/export	3.387	1.693	2.167	2.167	1783805
0.3%	2	0	0	PATCH	^/api/posts/[^/]+$	2.457	1.228	1.235	1.235	692
0.2%	2	0	0	GET	^/api/memories/[^/]+	1.888	0.944	0.948	0.948	3178
0.2%	1	0	0	PUT	^/api/blocks/[^/]+$	1.718	1.718	1.718	1.718	0
0.2%	5	2	0	GET	/api/me	1.519	0.304	0.508	0.508	127
0.1%	1	1	0	GET	^/api/posts/[^/]+$	1.247	1.247	1.247	1.247	27
0.1%	1	0	0	POST	/api/me/account-deletion	1.018	1.018	1.018	1.018	0
0.0%	1	0	0	GET	^/api/users/[^/]+$	0.404	0.404	0.404	0.404	166

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
42.9%	37	0	0	GET	/api/notifications	365.562	9.880	14.072	14.072	177
10.7%	25	0	0	POST	/api/posts	91.414	3.657	8.558	8.558	595
9.4%	102	1	0	GET	^/api/images/[^/]+	80.197	0.786	1.479	1.505	174152
6.7%	31	0	0	POST	/api/friend-requests	57.272	1.847	3.929	3.929	115
5.5%	49	0	0	GET	/api/memories	46.450	0.948	1.870	1.870	6424
4.4%	28	0	0	GET	^/api/feeds/[^/]+	37.792	1.350	1.671	1.671	693
3.9%	26	0	0	POST	/api/auth/login	33.157	1.275	1.950	1.950	251
3.6%	5	0	0	POST	/api/auth/register	30.233	6.047	7.467	7.467	230
2.2%	31	0	0	GET	/api/posting-window/current	19.115	0.617	1.356	1.356	240
1.9%	4	0	0	PUT	^/api/follows/[^/]+$	16.551	4.138	5.530	5.530	0
1.9%	17	0	0	GET	/api/users	15.871	0.934	1.863	1.863	223
1.7%	6	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	14.867	2.478	4.660	4.660	116
1.4%	22	8	0	GET	^/api/posts/me/current$	12.310	0.560	0.933	0.933	390
0.9%	18	0	0	GET	^/api/relationships/[^/]+	7.713	0.428	0.625	0.625	1121
0.6%	1	0	0	POST	/api/me/account-deletion	5.127	5.127	5.127	5.127	0
0.6%	6	0	0	POST	/api/auth/logout	5.046	0.841	1.032	1.032	0
0.5%	6	0	0	GET	^/api/memories/[^/]+	4.508	0.751	1.026	1.026	578
0.2%	2	2	0	PATCH	^/api/posts/[^/]+$	1.968	0.984	1.247	1.247	30
0.2%	1	0	0	PUT	^/api/blocks/[^/]+$	1.719	1.719	1.719	1.719	0
0.2%	1	0	0	DELETE	^/api/posts/[^/]+$	1.324	1.324	1.324	1.324	0
0.2%	4	1	0	GET	/api/me	1.320	0.330	0.406	0.406	153
0.1%	1	0	0	GET	/api/me/export	1.233	1.233	1.233	1.233	749608
0.1%	1	0	0	GET	^/api/users/[^/]+$	0.610	0.610	0.610	0.610	158
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 39036 クエリ / 197 種類 / 合計 451.547s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	21.2%	145	95.556	659.00	5691.25	3.1	1.0	SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE
2	20.1%	1031	90.639	87.91	5144.05	4.4	1.0	SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE
3	20.0%	62	90.517	1459.95	5817.17	5.8	0.0	LOCK TABLE users IN SHARE ROW EXCLUSIVE MODE
4	8.4%	1554	37.982	24.44	261.10	353.8	1.0	SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXI…
5	7.1%	32	32.032	1001.00	3660.22	12979.6	129.3	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
6	5.5%	18	24.944	1385.76	6419.81	32875.7	1639.1	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT gen_random_uuid()::text, id, $2, $3 FROM users WHERE region = $1 AND status = $4
7	3.5%	1753	15.612	8.91	47.34	87.1	1.0	SELECT u.secret_key_hash FROM image_objects i JOIN users u ON u.id = i.owner_user_id WHERE i.id = $1 AND u.status = $2
8	2.5%	79	11.446	144.89	380.74	225.8	4917.3	SELECT id, username, secret_key_hash, display_name, bio, profile_image_id, region, account_type, status FROM users WHERE status=$4 AND id<>$1 AND ($2=$5 OR (use…
9	2.2%	18	9.858	547.65	2609.53	13.3	0.8	UPDATE posting_windows SET ended_at = $2 WHERE region = $1 AND ended_at IS NULL
10	1.6%	142	7.243	51.01	212.07	931.3	1.0	INSERT INTO image_objects (id, owner_user_id, data, media_type, created_at) VALUES ($1, $2, $3, $5, $4)
11	1.6%	275	7.023	25.54	1176.59	4.4	1.0	SELECT id FROM users WHERE id = $1 FOR NO KEY UPDATE
12	0.8%	542	3.537	6.53	51.22	110.4	5.7	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
13	0.5%	55	2.277	41.41	110.73	610.8	1926.3	SELECT user_low_id, user_high_id FROM friendships WHERE user_low_id IN ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $2…
14	0.4%	1	1.826	1826.16	1826.16	29357.0	0.0	CREATE INDEX IF NOT EXISTS realmojis_image_id ON realmojis(image_id)
15	0.3%	1	1.221	1221.26	1221.26	63691.0	29504.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE;
-- rank 2
SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE;
-- rank 3
LOCK TABLE users IN SHARE ROW EXCLUSIVE MODE;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
未取得: metrics.pprofEnabled = false（アプリに pprof を仕込んでいない）
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 33.7% / 98% | 9% | 0.0% | 7 | 1913 MB | run queue が積もっている |
| isu2 (2 cpu) | 2.0% / 17% | 0% | 0.0% | 3 | 3206 MB |  |
| isu3 (2 cpu) | 2.3% / 12% | 0% | 0.0% | 3 | 3232 MB |  |
| isu4 (2 cpu) | 44.2% / 99% | 26% | 7.7% | 29 | 110 MB | run queue が積もっている |

**isu1（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 60% / nginx 5% |
| isu2 | isureal 3% |
| isu3 | isureal 4% |
| isu4 | postgres 16% |

## Diff from previous run

スコア: 2674 → 5421 (+2747)

`e240203..00ef856`

```
 webapp/sql/002-indexes.sql | 10 ++++++++++
 webapp/sql/reset.sql       |  2 ++
 2 files changed, 12 insertions(+)
```
