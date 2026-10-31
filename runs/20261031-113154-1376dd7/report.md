# run 20261031-113154-1376dd7

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | origin/main |
| sha | 1376dd789552 |
| 所要時間 | 130s |
| 前回比 | -21430 (-100.0%) vs 20261031-112920-4764fb5 |
| ベスト比 | -24801 (-100.0%) vs 20261031-103811-cef5fe3 = 24801 |
| note | main: +#67 無効化の受領待ち +登録のロックを予約だけに |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/contest-day-hosts/runs/20261031-113154-1376dd7/bench.log
```

## Bench

```
{"id":"4363","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"負荷走行を打ち切りました"}

--- stdout ---
time=2026-10-31T02:33:11.802Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:33:14.819Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:33:17.754Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:33:21.077Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T02:33:23.937Z level=INFO msg="europe 2 日目: 13 人"
time=2026-10-31T02:33:27.401Z level=INFO msg="america 2 日目: 15 人"
time=2026-10-31T02:33:30.816Z level=INFO msg="asia 3 日目: 48 人"
time=2026-10-31T02:33:33.614Z level=INFO msg="europe 3 日目: 13 人"
time=2026-10-31T02:33:38.191Z level=INFO msg="america 3 日目: 15 人"
time=2026-10-31T02:33:43.054Z level=INFO msg="asia 4 日目: 48 人"
time=2026-10-31T02:33:43.573Z level=INFO msg="europe 4 日目: 13 人"
time=2026-10-31T02:33:48.320Z level=INFO msg="america 4 日目: 15 人"
time=2026-10-31T02:34:01.753Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": false,
  "score": 0,
  "messages": [
    "負荷走行を打ち切りました: 45s〜50s のエラー率 131.4%（205/156 件）が上限 10.0% を超えました",
    "teru_5134: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "hina_nakamura_38: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "satoran5062: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "sutor.5178: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "seoyeon_lee_63: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "fan_yang_65: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "kame.5217: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "mini_truthahn_vom_holzkohlegrill_4646: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "leoivy5107: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "kina-5411: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "namu-5212: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "bensushi_5446: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "mechi5325: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "mahihe.5359: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "jihun_park_57: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "hais-5025: 友達申請を承認できません: POST /api/friend-requests/1a8a2d8e-2660-4663-96c4-fbc9d220d77b/accept: 応答がタイムアウトしました",
    "roto_5252: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "shiel5183: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "nean.5173: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "benkeleo.5464: 投稿できません: POST /api/posts: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
24.1%	224	0	48	GET	/api/notifications	1947.755	8.695	18.609	18.623	137
23.0%	2385	492	1071	GET	^/api/images/[^/]+	1865.481	0.782	12.536	14.080	24127
22.4%	1750	932	509	POST	/api/auth/login	1812.222	1.036	12.556	15.181	107
10.0%	108	1	25	POST	/api/posts	811.858	7.517	25.111	26.272	518
7.7%	107	73	14	PUT	/api/me/profile-image	621.369	5.807	22.226	23.232	39
3.1%	65	0	16	POST	^/api/friend-requests/[^/]+/[^/]+$	254.778	3.920	23.136	23.136	128
3.1%	231	8	19	POST	/api/friend-requests	254.672	1.102	13.745	18.710	116
1.2%	168	0	24	GET	/api/posting-window/current	96.152	0.572	12.138	12.313	227
1.1%	30	0	3	GET	/api/me/export	86.533	2.884	30.239	30.239	1940054
0.8%	181	0	1	GET	/api/memories	60.981	0.337	4.839	12.250	18053
0.7%	79	21	2	GET	^/api/posts/me/current$	54.668	0.692	12.102	12.102	456
0.7%	61	1	5	POST	/api/auth/register	53.161	0.871	13.696	13.696	224
0.6%	313	0	3	GET	^/api/feeds/[^/]+	45.007	0.144	1.672	2.698	435
0.4%	15	0	3	POST	/api/admin/posting-window	35.484	2.366	13.930	13.930	175
0.4%	6	0	2	DELETE	^/api/posts/[^/]+$	31.145	5.191	13.899	13.899	55
0.2%	106	0	5	GET	/api/users	16.225	0.153	1.660	1.665	247
0.1%	1	0	0	POST	/api/initialize	11.694	11.694	11.694	11.694	90
0.1%	49	0	0	POST	/api/auth/logout	10.912	0.223	2.700	2.700	0
0.1%	55	0	2	GET	^/api/relationships/[^/]+	7.810	0.142	0.884	0.884	3434
0.1%	44	0	0	GET	^/api/posts/[^/]+$	4.840	0.110	0.368	0.368	603
0.0%	13	0	0	PATCH	/api/me	3.417	0.263	0.799	0.799	276
0.0%	15	0	0	PUT	^/api/follows/[^/]+$	2.001	0.133	0.540	0.540	0
0.0%	15	0	0	GET	^/api/posts/[^/]+/comments$	1.858	0.124	0.368	0.368	29
0.0%	16	0	0	GET	^/api/posts/[^/]+/realmojis$	1.239	0.077	0.310	0.310	80
0.0%	23	2	0	GET	/api/me	1.214	0.053	0.274	0.274	274
0.0%	46	0	0	GET	^/api/memories/[^/]+	1.196	0.026	0.574	0.574	644
0.0%	7	0	0	PUT	^/api/admin/users/[^/]+/certification$	1.087	0.155	0.837	0.837	0
0.0%	8	2	2	PATCH	^/api/posts/[^/]+$	1.021	0.128	0.467	0.467	353
0.0%	5	0	0	POST	^/api/posts/[^/]+/realmojis$	0.944	0.189	0.300	0.300	407
0.0%	15	0	0	GET	^/api/users/[^/]+$	0.562	0.037	0.221	0.221	161
…（上位 30 / 全 35 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.4%	56	8	4	GET	/api/notifications	469.178	8.378	17.915	17.915	119
18.7%	24	0	5	POST	/api/posts	211.524	8.813	25.924	25.924	513
10.5%	44	0	6	POST	/api/friend-requests	118.468	2.692	18.322	18.322	122
5.3%	36	0	3	GET	/api/memories	59.797	1.661	13.403	13.403	22561
5.0%	47	0	3	POST	/api/auth/login	56.602	1.204	17.619	17.619	245
4.6%	36	0	8	GET	/api/posting-window/current	51.792	1.439	10.919	10.919	223
2.7%	12	0	1	POST	^/api/friend-requests/[^/]+/[^/]+$	31.040	2.587	18.319	18.319	120
2.6%	151	1	0	GET	^/api/images/[^/]+	29.638	0.196	1.241	1.245	135696
2.6%	84	0	1	GET	^/api/feeds/[^/]+	29.576	0.352	12.338	12.338	666
1.3%	29	0	2	GET	/api/users	14.344	0.495	3.320	3.320	244
0.8%	4	0	0	PUT	/api/me/profile-image	9.001	2.250	3.957	3.957	0
0.7%	10	0	0	PUT	^/api/follows/[^/]+$	8.455	0.846	2.035	2.035	0
0.6%	16	2	0	GET	^/api/posts/me/current$	7.223	0.451	3.648	3.648	522
0.6%	7	1	0	POST	^/api/posts/[^/]+/comments$	6.264	0.895	3.498	3.498	243
0.5%	9	0	0	POST	/api/auth/logout	5.932	0.659	1.576	1.576	0
0.5%	11	0	0	POST	/api/auth/register	5.135	0.467	1.340	1.340	228
0.2%	5	1	0	GET	/api/me	2.485	0.497	0.984	0.984	250
0.2%	5	0	0	POST	^/api/posts/[^/]+/realmojis$	2.385	0.477	1.359	1.359	405
0.2%	2	0	0	PATCH	/api/me	2.126	1.063	1.138	1.138	275
0.2%	12	0	0	GET	^/api/relationships/[^/]+	2.102	0.175	0.557	0.557	400
0.2%	13	0	0	GET	^/api/posts/[^/]+$	2.003	0.154	0.321	0.321	585
0.2%	11	0	0	GET	^/api/posts/[^/]+/comments$	1.919	0.174	0.395	0.395	255
0.1%	1	0	0	DELETE	^/api/posts/[^/]+$	1.246	1.246	1.246	1.246	0
0.1%	9	0	0	GET	^/api/users/[^/]+$	1.230	0.137	0.209	0.209	161
0.1%	3	0	0	DELETE	^/api/comments/[^/]+$	0.999	0.333	0.470	0.470	0
0.1%	6	0	0	GET	^/api/posts/[^/]+/realmojis$	0.801	0.134	0.181	0.181	164
0.0%	1	0	0	GET	/api/me/export	0.496	0.496	0.496	0.496	450367
0.0%	3	0	0	GET	^/api/memories/[^/]+	0.314	0.105	0.110	0.110	448
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.106	0.106	0.106	0.106	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.5%	55	6	6	GET	/api/notifications	439.722	7.995	13.722	13.722	121
15.8%	44	1	10	POST	/api/friend-requests	166.975	3.795	20.971	20.971	125
15.1%	18	1	5	POST	/api/posts	160.101	8.895	25.218	25.218	431
6.2%	44	0	4	POST	/api/auth/login	65.803	1.496	13.747	13.747	235
3.9%	34	0	7	GET	/api/posting-window/current	41.822	1.230	12.376	12.376	225
2.5%	16	0	3	GET	/api/users	26.866	1.679	12.734	12.734	231
2.0%	112	1	0	GET	^/api/images/[^/]+	21.613	0.193	0.950	1.235	98249
1.6%	37	0	1	GET	^/api/feeds/[^/]+	17.433	0.471	7.376	7.376	790
1.6%	11	0	0	PUT	^/api/follows/[^/]+$	17.378	1.580	4.157	4.157	0
1.4%	20	0	0	GET	/api/memories	14.987	0.749	4.191	4.191	28014
1.3%	6	0	0	POST	/api/auth/logout	13.961	2.327	6.154	6.154	0
1.1%	6	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	12.048	2.008	5.323	5.323	116
0.9%	6	0	0	PATCH	/api/me	9.167	1.528	2.876	2.876	256
0.9%	4	0	0	PUT	/api/me/profile-image	9.057	2.264	3.065	3.065	0
0.8%	14	0	0	POST	/api/auth/register	8.919	0.637	1.121	1.121	229
0.8%	7	0	0	POST	^/api/posts/[^/]+/realmojis$	8.564	1.223	2.439	2.439	408
0.6%	3	0	0	GET	/api/me/export	6.251	2.084	3.193	3.193	2994805
0.4%	11	2	0	GET	^/api/posts/me/current$	4.461	0.406	0.793	0.793	479
0.2%	2	0	0	POST	^/api/posts/[^/]+/comments$	2.455	1.228	1.993	1.993	264
0.2%	2	0	0	PATCH	^/api/posts/[^/]+$	2.409	1.205	2.164	2.164	569
0.2%	4	0	0	GET	/api/me	2.226	0.556	0.786	0.786	307
0.2%	7	0	0	GET	^/api/memories/[^/]+	2.187	0.312	0.654	0.654	870
0.1%	1	0	0	DELETE	^/api/realmojis/[^/]+$	1.448	1.448	1.448	1.448	0
0.1%	4	0	0	GET	^/api/relationships/[^/]+	1.385	0.346	0.827	0.827	656
0.1%	6	0	0	GET	^/api/posts/[^/]+$	1.246	0.208	0.332	0.332	589
0.0%	1	0	0	DELETE	^/api/posts/[^/]+$	0.461	0.461	0.461	0.461	0
0.0%	1	0	0	GET	^/api/users/[^/]+$	0.352	0.352	0.352	0.352	169
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.108	0.108	0.108	0.108	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 5241 クエリ / 166 種類 / 合計 45.716s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	36.0%	32	16.449	514.04	6840.09	6424.8	141.2	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	26.5%	1	12.094	12093.55	12093.55	70708.0	19054.0	COPY public.image_objects_202510 (id, owner_user_id, data, media_type, created_at) TO stdout
3	8.6%	1	3.936	3935.98	3935.98	11528.0	546016.0	COPY public.comments (id, post_id, author_id, body, created_at, deleted_at) TO stdout
4	5.9%	241	2.686	11.14	205.52	109.9	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	5.1%	1	2.349	2348.51	2348.51	29389.0	1199467.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	3.9%	1	1.776	1776.05	1776.05	19872.0	625555.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
7	3.2%	171	1.474	8.62	56.45	712.7	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
8	2.6%	701	1.190	1.70	8.30	13.2	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
9	1.5%	1	0.702	701.86	701.86	11528.0	529703.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
10	0.7%	2	0.336	168.20	179.56	215553.5	2535.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
11	0.6%	2	0.262	131.11	253.71	16763.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
12	0.5%	1	0.246	246.12	246.12	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
13	0.5%	124	0.224	1.81	3.82	68.2	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
14	0.4%	1	0.193	193.45	193.45	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
15	0.4%	1	0.177	176.95	176.95	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
COPY public.image_objects_202510 (id, owner_user_id, data, media_type, created_at) TO stdout;
-- rank 3
COPY public.comments (id, post_id, author_id, body, created_at, deleted_at) TO stdout;
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
| isu1 (2 cpu) | 80.7% / 100% | 65% | 4.1% | 8 | 81 MB | CPU が飽和 |
| isu2 (2 cpu) | 38.4% / 99% | 1% | 31.2% | 5 | 418 MB | IO 待ちが重い |
| isu3 (2 cpu) | 0.8% / 3% | 0% | 0.1% | 2 | 299 MB |  |
| isu4 (2 cpu) | 18.0% / 72% | 0% | 8.7% | 8 | 111 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（IO 待ちが重い、主に du）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 90% / pg_dump 41% / nginx 9% / sshd 4% / udisksd 3% |
| isu2 | du 5% / tar 3% / kswapd0 3% / sshd 1% |
| isu3 | 目立つプロセス無し |
| isu4 | postgres 15% |

## Diff from previous run

スコア: 21430 → — (fail)

`4764fb5..1376dd7`

```
 conf/isu1/home/isucon/env.sh |   1 +
 webapp/go/image_local.go     | 146 ++++++++++++++++++++++++++++++++++---------
 webapp/go/user_api.go        |  17 ++++-
 3 files changed, 133 insertions(+), 31 deletions(-)
```
