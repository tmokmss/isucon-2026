# run 20261031-124342-0db3f89

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | repl-replica |
| sha | 0db3f895ec8c |
| 所要時間 | 304s |
| 前回比 | -64700 (-100.0%) vs 20261031-121235-e0e4cb3 |
| ベスト比 | -64700 (-100.0%) vs 20261031-121235-e0e4cb3 = 64700 |
| note | 複製 第1段 shadow（europe1・america1 が裏でフィード等を計算して asia1 の応答と比べる） |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/agent-a7f5c44d209df1790/runs/20261031-124342-0db3f89/bench.log
```

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	143	400	479	480
europe	7	17	34	34	37	37
america	6	10	23	22	22	21
```

```
{"id":"9653","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"負荷走行を打ち切りました"}

--- stdout ---
time=2026-10-31T03:47:41.649Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T03:47:44.619Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T03:47:47.615Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T03:47:50.666Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T03:47:53.611Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T03:47:56.613Z level=INFO msg="america 2 日目: 10 人"
time=2026-10-31T03:47:56.674Z level=INFO msg="user moved" user=heel-5135 from=america to=asia
time=2026-10-31T03:47:57.131Z level=INFO msg="user moved" user=heel-5135 from=asia to=europe
time=2026-10-31T03:47:59.639Z level=INFO msg="asia 3 日目: 143 人"
time=2026-10-31T03:48:02.129Z level=INFO msg="user moved" user=hesol.5726 from=asia to=europe
time=2026-10-31T03:48:02.498Z level=INFO msg="user moved" user=ruyu.5729 from=asia to=america
time=2026-10-31T03:48:02.652Z level=INFO msg="europe 3 日目: 34 人"
time=2026-10-31T03:48:03.239Z level=INFO msg="user moved" user=shiis5942 from=asia to=america
time=2026-10-31T03:48:03.907Z level=INFO msg="user moved" user=mukolia-6019 from=asia to=america
time=2026-10-31T03:48:05.635Z level=INFO msg="america 3 日目: 23 人"
time=2026-10-31T03:48:08.664Z level=INFO msg="asia 4 日目: 400 人"
time=2026-10-31T03:48:11.913Z level=INFO msg="europe 4 日目: 34 人"
time=2026-10-31T03:48:14.829Z level=INFO msg="america 4 日目: 22 人"
time=2026-10-31T03:48:16.689Z level=INFO msg="user moved" user=miarimar_7891 from=asia to=europe
time=2026-10-31T03:48:17.165Z level=INFO msg="user moved" user=hitano-6469 from=asia to=europe
time=2026-10-31T03:48:17.901Z level=INFO msg="asia 5 日目: 479 人"
time=2026-10-31T03:48:21.262Z level=INFO msg="europe 5 日目: 37 人"
time=2026-10-31T03:48:24.007Z level=INFO msg="america 5 日目: 22 人"
time=2026-10-31T03:48:24.477Z level=INFO msg="user moved" user=mukolia-6019 from=america to=asia
time=2026-10-31T03:48:32.756Z level=INFO msg="asia 最終日（グランプリ）: 480 人"
time=2026-10-31T03:48:32.841Z level=INFO msg="europe 最終日（グランプリ）: 37 人"
time=2026-10-31T03:48:32.873Z level=INFO msg="america 最終日（グランプリ）: 21 人"
time=2026-10-31T03:48:41.616Z level=INFO msg="グランプリ: アジア 6 人, ヨーロッパ 0 人, アメリカ 1 人"
{
  "pass": false,
  "score": 0,
  "messages": [
    "負荷走行を打ち切りました: 55s〜1m0s のエラー率 35.3%（59/167 件）が上限 10.0% を超えました",
    "salmon_in_a_jar_2533: ログアウトできません: POST /api/auth/logout: ステータスコード 502 が返りました（期待値 204、unknown）",
    "nusa6249: 友達申請を送れません: POST /api/friend-requests: ステータスコード 502 が返りました（期待値 201、unknown）",
    "ivymutsu_5473: 投稿できません: POST /api/posts: ステータスコード 502 が返りました（期待値 201、unknown）",
    "elnamax-5459: 投稿できません: POST /api/posts: ステータスコード 502 が返りました（期待値 201、unknown）",
    "muma.6237: 友達申請を送れません: POST /api/friend-requests: ステータスコード 502 が返りました（期待値 201、unknown）",
    "shidantsu6259: 友達申請を送れません: POST /api/friend-requests: ステータスコード 502 が返りました（期待値 201、unknown）",
    "liase-6269: 新しいユーザーを登録できません: POST /api/auth/register: ステータスコード 502 が返りました（期待値 201、unknown）",
    "helena_santos_4199: 投稿できません: POST /api/posts: ステータスコード 502 が返りました（期待値 201、unknown）",
    "leoliaivy6223: 投稿できません: POST /api/posts: ステータスコード 502 が返りました（期待値 201、unknown）",
    "fuhomax5106: 友達申請を承認できません: POST /api/friend-requests/da5c3edc-c88d-4964-a530-1be0bf1b3ebe/accept: ステータスコード 502 が返りました（期待値 200、unknown）",
    "reho.5439: 友達申請を承認できません: POST /api/friend-requests/4692d97c-ef2a-4a69-93c8-3ac72c2a7e93/accept: ステータスコード 502 が返りました（期待値 200、unknown）",
    "reyadan-6149: 友達申請を送れません: POST /api/friend-requests: ステータスコード 502 が返りました（期待値 201、unknown）",
    "benul_6241: 投稿できません: POST /api/posts: ステータスコード 502 が返りました（期待値 201、unknown）",
    "ramar_6159: 友達申請を送れません: POST /api/friend-requests: ステータスコード 502 が返りました（期待値 201、unknown）",
    "evachi_5394: 投稿 8ee04eed-5ab8-4124-adf0-0e1c1d294b7f に RealMoji を送れません: POST /api/posts/8ee04eed-5ab8-4124-adf0-0e1c1d294b7f/realmojis: ステータスコード 502 が返りました（期待値 201、unknown）",
    "danlia6202: 投稿できません: POST /api/posts: ステータスコード 502 が返りました（期待値 201、unknown）",
    "elias_hoffmann_8: 履歴を読めません: GET /api/memories?month=2026-08: 通信に失敗しました: 応答の本体を読めません: unexpected EOF",
    "rian-5374: 投稿 75629362-2f08-447d-bb08-f997592fb37e に RealMoji を送れません: POST /api/posts/75629362-2f08-447d-bb08-f997592fb37e/realmojis: ステータスコード 502 が返りました（期待値 201、unknown）",
    "zoehena_5452: 友達申請を承認できません: POST /api/friend-requests/5dc48186-771c-4fa3-b054-951c1ae39945/accept: ステータスコード 502 が返りました（期待値 200、unknown）",
    "salmon_in_a_jar_2533: ログインできません: POST /api/auth/login: ステータスコード 502 が返りました（期待値 200、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.1%	1041	0	0	GET	/api/notifications	9217.408	8.854	14.890	15.206	173
22.1%	1413	34	0	POST	/api/posts	5788.568	4.097	12.575	13.088	636
11.2%	9369	888	0	GET	^/api/images/[^/]+	2942.413	0.314	6.706	13.102	138631
6.7%	2781	62	0	POST	/api/friend-requests	1750.352	0.629	5.249	8.823	113
6.6%	3523	935	0	POST	/api/auth/login	1734.084	0.492	7.989	12.774	186
3.5%	1261	5	0	POST	^/api/friend-requests/[^/]+/[^/]+$	926.587	0.735	8.192	13.174	116
2.6%	501	1	0	POST	/api/auth/register	676.955	1.351	5.911	8.149	227
2.3%	1931	0	0	GET	/api/posting-window/current	608.752	0.315	3.867	7.486	237
1.4%	470	7	0	POST	^/api/posts/[^/]+/realmojis$	373.262	0.794	3.922	8.039	464
1.3%	408	123	0	PUT	/api/me/profile-image	332.664	0.815	5.474	6.559	7
1.1%	1796	0	0	GET	/api/memories	281.304	0.157	0.850	4.872	13751
0.9%	8243	0	0	GET	^/api/feeds/[^/]+	230.726	0.028	0.140	6.123	2156
0.8%	391	3	0	POST	^/api/posts/[^/]+/comments$	202.578	0.518	3.552	3.918	362
0.7%	319	0	0	PUT	^/api/follows/[^/]+$	178.987	0.561	3.214	6.994	0
0.6%	249	0	0	PATCH	/api/me	154.747	0.621	3.779	7.346	279
0.6%	153	0	0	GET	/api/me/export	150.237	0.982	7.973	12.234	2061549
0.4%	452	0	0	POST	/api/auth/logout	116.898	0.259	4.696	11.054	0
0.4%	139	6	0	PATCH	^/api/posts/[^/]+$	110.683	0.796	5.126	5.516	618
0.3%	79	0	0	DELETE	^/api/realmojis/[^/]+$	86.001	1.089	11.703	11.703	0
0.2%	42	0	0	DELETE	^/api/posts/[^/]+$	50.983	1.214	6.108	6.108	0
0.2%	1318	0	0	GET	/api/users	45.407	0.034	0.311	1.089	232
0.2%	982	39	0	GET	^/api/posts/me/current$	41.853	0.043	0.369	1.839	623
0.1%	398	0	0	GET	^/api/memories/[^/]+	38.164	0.096	1.254	4.431	689
0.1%	84	0	0	DELETE	^/api/comments/[^/]+$	33.926	0.404	7.532	7.532	0
0.1%	1408	11	0	GET	^/api/posts/[^/]+$	33.652	0.024	0.113	0.169	629
0.1%	892	0	0	GET	^/api/relationships/[^/]+	26.571	0.030	0.172	1.319	711
0.1%	1	0	0	POST	/api/initialize	18.325	18.325	18.325	18.325	90
0.1%	684	0	0	GET	^/api/users/[^/]+$	17.078	0.025	0.129	0.825	211
0.1%	6	0	0	DELETE	/api/me/profile-image	16.112	2.685	12.113	12.113	0
0.1%	612	0	0	GET	^/api/posts/[^/]+/realmojis$	16.012	0.026	0.121	0.230	1077
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.6%	140	0	0	GET	/api/notifications	1217.190	8.694	11.752	11.752	180
13.3%	660	3	0	GET	^/api/images/[^/]+	362.653	0.549	4.226	6.408	201004
11.2%	140	1	26	POST	/api/posts	306.722	2.191	5.457	5.743	549
6.3%	154	0	8	POST	^/api/friend-requests/[^/]+/[^/]+$	172.023	1.117	5.997	9.067	119
4.9%	617	0	0	GET	^/api/feeds/[^/]+	133.024	0.216	0.674	0.995	2233
4.1%	174	0	6	POST	/api/auth/login	112.222	0.645	5.898	6.852	249
3.1%	207	3	3	POST	/api/friend-requests	85.675	0.414	4.646	6.483	115
2.7%	174	0	0	GET	/api/posting-window/current	73.049	0.420	5.773	5.774	239
1.9%	172	0	0	GET	/api/memories	51.621	0.300	1.988	3.317	14314
1.1%	37	1	2	POST	^/api/posts/[^/]+/realmojis$	30.504	0.824	3.526	3.526	448
0.7%	110	4	0	GET	^/api/posts/me/current$	20.434	0.186	0.338	0.634	620
0.6%	90	0	0	GET	/api/users	17.526	0.195	1.646	1.646	235
0.6%	82	8	0	GET	^/api/posts/[^/]+$	16.233	0.198	0.845	0.845	555
0.5%	15	1	0	POST	^/api/posts/[^/]+/comments$	13.371	0.891	2.659	2.659	295
0.5%	27	0	1	POST	/api/auth/logout	13.011	0.482	4.888	4.888	6
0.5%	9	0	0	GET	/api/me/export	12.678	1.409	2.738	2.738	1300994
0.5%	66	0	0	GET	^/api/relationships/[^/]+	12.391	0.188	0.367	0.367	1161
0.4%	17	0	0	PATCH	^/api/posts/[^/]+$	11.321	0.666	3.579	3.579	642
0.4%	46	0	0	GET	^/api/users/[^/]+$	9.874	0.215	0.691	0.691	198
0.3%	39	0	0	GET	^/api/posts/[^/]+/realmojis$	9.466	0.243	0.744	0.744	1026
0.3%	2	0	0	DELETE	^/api/comments/[^/]+$	8.173	4.087	7.516	7.516	0
0.3%	34	0	0	GET	^/api/posts/[^/]+/comments$	7.572	0.223	0.608	0.608	872
0.2%	31	0	0	POST	/api/auth/register	6.397	0.206	0.933	0.933	228
0.2%	24	0	0	GET	^/api/memories/[^/]+	5.041	0.210	0.395	0.395	674
0.2%	18	0	0	PUT	^/api/follows/[^/]+$	4.872	0.271	1.855	1.855	0
0.2%	15	0	0	PUT	/api/me/profile-image	4.760	0.317	0.406	0.406	0
0.1%	5	0	0	DELETE	^/api/realmojis/[^/]+$	4.089	0.818	1.052	1.052	0
0.1%	16	1	0	GET	/api/me	2.949	0.184	0.281	0.281	287
0.1%	3	0	0	DELETE	^/api/posts/[^/]+$	2.436	0.812	1.730	1.730	0
0.1%	10	0	0	PATCH	/api/me	2.402	0.240	0.770	0.770	282
…（上位 30 / 全 32 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.2%	86	0	0	GET	/api/notifications	705.731	8.206	9.420	9.420	180
10.9%	86	1	19	POST	/api/posts	170.967	1.988	7.488	7.488	518
10.1%	344	3	0	GET	^/api/images/[^/]+	158.165	0.460	3.079	4.789	191518
8.2%	107	1	4	POST	^/api/friend-requests/[^/]+/[^/]+$	127.720	1.194	7.434	8.317	117
5.0%	132	2	3	POST	/api/friend-requests	78.295	0.593	5.159	5.704	115
4.5%	102	0	2	POST	/api/auth/login	70.602	0.692	7.947	8.376	253
3.0%	237	0	0	GET	^/api/feeds/[^/]+	46.506	0.196	0.612	0.728	2542
2.5%	109	0	0	GET	/api/posting-window/current	39.422	0.362	3.144	4.403	240
1.6%	90	0	0	GET	/api/memories	24.773	0.275	1.019	1.019	14908
1.5%	8	0	0	GET	/api/me/export	24.027	3.003	5.036	5.036	4216255
0.9%	66	3	0	GET	^/api/posts/me/current$	13.520	0.205	0.792	0.792	598
0.9%	31	2	0	POST	^/api/posts/[^/]+/comments$	13.435	0.433	2.608	2.608	283
0.7%	61	0	0	GET	/api/users	11.397	0.187	0.680	0.680	243
0.7%	16	0	2	POST	/api/auth/logout	10.469	0.654	3.876	3.876	21
0.5%	7	1	0	PATCH	^/api/posts/[^/]+$	8.338	1.191	5.285	5.285	540
0.5%	3	0	0	DELETE	^/api/posts/[^/]+$	8.038	2.679	4.111	4.111	0
0.5%	43	2	0	GET	^/api/posts/[^/]+$	7.882	0.183	0.657	0.657	584
0.4%	32	0	0	GET	^/api/relationships/[^/]+	6.374	0.199	0.553	0.553	1614
0.4%	7	0	0	PUT	/api/me/profile-image	5.850	0.836	1.891	1.891	0
0.4%	8	0	0	PATCH	/api/me	5.516	0.690	2.959	2.959	273
0.3%	25	0	0	GET	^/api/users/[^/]+$	5.060	0.202	0.650	0.650	196
0.3%	16	0	0	PUT	^/api/follows/[^/]+$	4.047	0.253	1.815	1.815	0
0.2%	20	0	1	POST	/api/auth/register	3.380	0.169	0.434	0.434	227
0.2%	18	0	0	GET	^/api/memories/[^/]+	3.219	0.179	0.443	0.443	827
0.2%	2	0	0	POST	^/api/posts/[^/]+/realmojis$	3.020	1.510	2.086	2.086	541
0.2%	17	0	0	GET	^/api/posts/[^/]+/comments$	2.758	0.162	0.268	0.268	727
0.1%	9	0	0	GET	^/api/posts/[^/]+/realmojis$	1.640	0.182	0.233	0.233	1602
0.1%	7	0	0	GET	/api/me	1.606	0.229	0.703	0.703	306
0.0%	4	0	0	DELETE	^/api/comments/[^/]+$	0.688	0.172	0.220	0.220	0
0.0%	1	0	0	PUT	^/api/blocks/[^/]+$	0.116	0.116	0.116	0.116	0
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 40721 クエリ / 81 種類 / 合計 430.503s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	42.2%	2344	181.887	77.60	708.37	623.8	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
2	16.1%	7	69.221	9888.73	12471.99	30127.9	1199642.6	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	15.6%	170	67.310	395.94	3018.27	6954.7	137.7	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
4	12.8%	7	55.311	7901.61	12143.17	21843.7	626161.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	3.9%	581	16.841	28.99	645.71	122.6	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
6	3.3%	6070	14.248	2.35	79.84	10.5	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
7	1.8%	3	7.623	2541.07	3465.95	11805.0	529802.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
8	1.4%	4	5.870	1467.55	5281.60	219429.8	2597.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
9	1.1%	1537	4.559	2.97	38.67	26.0	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
10	0.4%	500	1.849	3.70	13.15	29.8	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
11	0.2%	431	0.819	1.90	12.48	21.9	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
12	0.2%	3028	0.662	0.22	20.87	20.1	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
13	0.1%	4	0.406	101.44	185.42	622.8	25405.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
14	0.1%	2859	0.371	0.13	16.71	11.3	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
15	0.1%	4	0.358	89.51	156.87	622.0	30064.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 1f79951b369c68c7bdfd36bbf2952d778294a3bd
Type: cpu
Time: 2026-10-31 12:47:16 JST
Duration: 60.11s, Total samples = 42.38s (70.50%)
Showing nodes accounting for 32.94s, 77.73% of 42.38s total
Dropped 979 nodes (cum <= 0.21s)
Showing top 40 nodes out of 252
      flat  flat%   sum%        cum   cum%
     3.40s  8.02%  8.02%      3.40s  8.02%  crypto/internal/fips140/sha256.blockAVX2
     2.90s  6.84% 14.87%      9.61s 22.68%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     2.59s  6.11% 20.98%      4.28s 10.10%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     2.40s  5.66% 26.64%      2.40s  5.66%  internal/runtime/syscall/linux.Syscall6
     2.36s  5.57% 32.21%      2.42s  5.71%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     1.70s  4.01% 36.22%      1.70s  4.01%  runtime.memmove
     1.69s  3.99% 40.21%      1.69s  3.99%  runtime.memclrNoHeapPointers
     1.64s  3.87% 44.08%      5.47s 12.91%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.50s  3.54% 47.62%      1.91s  4.51%  runtime.findObject
     1.46s  3.45% 51.06%      5.91s 13.95%  runtime.scanObject
     1.33s  3.14% 54.20%      1.33s  3.14%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.21s  2.86% 57.06%      1.73s  4.08%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     1.14s  2.69% 59.75%      1.97s  4.65%  runtime.tryDeferToSpanScan
     1.04s  2.45% 62.20%      1.04s  2.45%  aeshashbody
     0.52s  1.23% 63.43%      0.52s  1.23%  runtime.spanClass.sizeclass (inline)
     0.50s  1.18% 64.61%      1.88s  4.44%  runtime.mapaccess1_faststr
     0.49s  1.16% 65.76%      0.52s  1.23%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.44s  1.04% 66.80%      0.44s  1.04%  example.com/isureal-company/isureal.(*UserRow).active
     0.39s  0.92% 67.72%      0.39s  0.92%  memeqbody
     0.32s  0.76% 68.48%      0.32s  0.76%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.31s  0.73% 69.21%      0.52s  1.23%  example.com/isureal-company/isureal.postLess
     0.31s  0.73% 69.94%      0.31s  0.73%  runtime.nextFreeFast (inline)
     0.30s  0.71% 70.65%      0.60s  1.42%  database/sql.convertAssignRows
     0.26s  0.61% 71.26%      0.26s  0.61%  indexbody
     0.26s  0.61% 71.87%      0.50s  1.18%  runtime.greyobject
     0.25s  0.59% 72.46%      0.83s  1.96%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.57% 73.03%      1.62s  3.82%  runtime.mapassign_faststr
     0.23s  0.54% 73.57%      3.86s  9.11%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.23s  0.54% 74.12%      0.23s  0.54%  runtime.typePointers.nextFast (inline)
     0.20s  0.47% 74.59%      3.12s  7.36%  runtime.mallocgc
     0.19s  0.45% 75.04%      2.90s  6.84%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.17s   0.4% 75.44%      0.22s  0.52%  runtime.(*mspan).writeHeapBitsSmall
     0.16s  0.38% 75.81%      0.22s  0.52%  runtime.heapArenaOf (inline)
     0.14s  0.33% 76.14%      0.25s  0.59%  internal/runtime/maps.(*Iter).Next
     0.12s  0.28% 76.43%      1.24s  2.93%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.12s  0.28% 76.71%      0.48s  1.13%  runtime.mapaccess2_faststr
     0.11s  0.26% 76.97%      3.76s  8.87%  crypto/internal/fips140/sha256.(*Digest).Write
     0.11s  0.26% 77.23%      1.03s  2.43%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.11s  0.26% 77.49%      0.53s  1.25%  runtime.mallocgcSmallNoscan
     0.10s  0.24% 77.73%      4.07s  9.60%  crypto/internal/fips140/sha256.(*Digest).Sum
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 60.8% / 100% | 42% | 0.1% | 6 | 131 MB | CPU が飽和 |
| isu2 (2 cpu) | 18.0% / 41% | 0% | 0.2% | 3 | 1935 MB |  |
| isu3 (2 cpu) | 18.0% / 45% | 0% | 0.2% | 4 | 1977 MB |  |
| isu4 (2 cpu) | 61.0% / 100% | 30% | 16.2% | 29 | 112 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 98% / nginx 14% / journalctl 2% / systemd-journal 2% |
| isu2 | isureal 32% |
| isu3 | isureal 34% |
| isu4 | postgres 63% / isureal 21% / kcompactd0 1% |

## Diff from previous run

スコア: 64700 → — (fail)

`e0e4cb3..0db3f89`

```
 conf/isu2/etc/nginx/sites-enabled/isureal.conf |   7 +-
 conf/isu2/home/isucon/env.sh                   |   3 +
 conf/isu3/etc/nginx/sites-enabled/isureal.conf |   7 +-
 conf/isu3/home/isucon/env.sh                   |   3 +
 webapp/go/account_deletion_api.go              |  10 +-
 webapp/go/admin_posting_window_api.go          |  16 +-
 webapp/go/admin_user_api.go                    |   5 +-
 webapp/go/cmd/isureal/main.go                  |   6 +
 webapp/go/comment_api.go                       |   9 +-
 webapp/go/feed_api.go                          | 113 +++-
 webapp/go/feed_test.go                         | 331 ++++++++++
 webapp/go/friend_request_api.go                |  15 +-
 webapp/go/image_local.go                       |   1 +
 webapp/go/initialize.go                        |   4 +-
 webapp/go/post_api.go                          |  14 +-
 webapp/go/profile_image_api.go                 |  13 +-
 webapp/go/realmoji_api.go                      |  16 +-
 webapp/go/relationship.go                      |  10 +-
 webapp/go/relationship_api.go                  |  33 +-
 webapp/go/repl.go                              | 832 +++++++++++++++++++++++++
 webapp/go/repl_test.go                         |  62 ++
 webapp/go/replica_server.go                    | 158 +++++
 webapp/go/store.go                             | 113 +++-
 webapp/go/user_api.go                          |  34 +-
 24 files changed, 1656 insertions(+), 159 deletions(-)
```
