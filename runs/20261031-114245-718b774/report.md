# run 20261031-114245-718b774

## Summary

| | |
|---|---|
| score | 48796 |
| status | pass |
| ref | offload-cpu |
| sha | 718b774f966f |
| 所要時間 | 329s |
| 前回比 | +31724 (185.8%) vs 20261031-113441-6a7a33a |
| ベスト比 | +23995 (96.8%) vs 20261031-103811-cef5fe3 = 24801 **ベスト更新** |
| note | offload-sign-only-on-main-f782d77 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	146	402	402	173
europe	7	17	45	50	50	46
america	6	16	38	39	38	37
```

```
{"id":"5303","status":"done","portalStatus":"finished","passed":true,"score":48796,"raw":48796,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:46:52.373Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:46:55.340Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:46:58.337Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:47:01.368Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T02:47:04.363Z level=INFO msg="europe 2 日目: 17 人"
time=2026-10-31T02:47:07.336Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T02:47:10.570Z level=INFO msg="asia 3 日目: 146 人"
time=2026-10-31T02:47:13.539Z level=INFO msg="europe 3 日目: 45 人"
time=2026-10-31T02:47:14.078Z level=INFO msg="user moved" user=moul.5672 from=asia to=europe
time=2026-10-31T02:47:15.535Z level=INFO msg="user moved" user=ranoaki_5863 from=asia to=america
time=2026-10-31T02:47:16.373Z level=INFO msg="america 3 日目: 38 人"
time=2026-10-31T02:47:16.431Z level=INFO msg="user moved" user=nusomi.6074 from=asia to=america
time=2026-10-31T02:47:16.576Z level=INFO msg="user moved" user=miatate-5532 from=asia to=europe
time=2026-10-31T02:47:19.604Z level=INFO msg="asia 4 日目: 402 人"
time=2026-10-31T02:47:24.168Z level=INFO msg="europe 4 日目: 50 人"
time=2026-10-31T02:47:25.641Z level=INFO msg="america 4 日目: 39 人"
time=2026-10-31T02:47:29.467Z level=INFO msg="asia 5 日目: 402 人"
time=2026-10-31T02:47:30.263Z level=INFO msg="user moved" user=nusomi.6074 from=america to=europe
time=2026-10-31T02:47:30.356Z level=INFO msg="user moved" user=moul.5672 from=europe to=asia
time=2026-10-31T02:47:32.306Z level=INFO msg="europe 5 日目: 50 人"
time=2026-10-31T02:47:34.411Z level=INFO msg="america 5 日目: 38 人"
time=2026-10-31T02:47:43.404Z level=INFO msg="america 最終日（グランプリ）: 37 人"
time=2026-10-31T02:47:43.414Z level=INFO msg="europe 最終日（グランプリ）: 46 人"
time=2026-10-31T02:47:43.461Z level=INFO msg="asia 最終日（グランプリ）: 173 人"
time=2026-10-31T02:47:52.338Z level=INFO msg="グランプリ: アジア 24 人, ヨーロッパ 0 人, アメリカ 4 人"
{
  "pass": true,
  "score": 48796,
  "messages": [
    "eltor6749: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "elvinso_6763: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "hesol_6765: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kuyo-6773: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "benkefu-6781: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "liaben_6790: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ramar6796: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "rimi6799: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "saroku.6810: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "rure-6813: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "hihi-6822: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "elha_6833: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "metsu-6836: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "yoseku-6847: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "teta.6850: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kesolta6855: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "tachi6867: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "nianki6869: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ronichi.6877: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "naliare-6887: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
27.4%	556	0	0	GET	/api/notifications	5127.737	9.223	13.989	13.989	175
24.9%	422	53	0	POST	/api/auth/register	4652.367	11.025	20.659	21.630	199
12.7%	8700	607	0	GET	^/api/images/[^/]+	2373.346	0.273	4.876	12.034	87745
12.0%	682	15	0	POST	/api/posts	2249.043	3.298	11.867	12.000	636
7.0%	3837	1315	0	POST	/api/auth/login	1312.815	0.342	4.769	11.645	173
3.2%	1193	40	0	POST	/api/friend-requests	589.561	0.494	6.119	8.107	112
2.4%	822	6	0	POST	^/api/friend-requests/[^/]+/[^/]+$	453.524	0.552	5.435	12.164	115
1.7%	2150	0	0	GET	/api/memories	322.845	0.150	1.106	8.060	12617
1.7%	352	35	0	POST	^/api/posts/[^/]+/realmojis$	322.031	0.915	6.776	7.870	419
1.0%	5050	0	0	GET	^/api/feeds/[^/]+	185.858	0.037	0.401	1.838	2842
1.0%	202	0	0	GET	/api/me/export	185.335	0.917	6.835	8.337	1810241
0.9%	859	0	0	GET	/api/posting-window/current	167.378	0.195	1.267	2.073	237
0.9%	193	91	0	PUT	/api/me/profile-image	161.170	0.835	7.767	8.126	10
0.8%	205	16	0	POST	^/api/posts/[^/]+/comments$	142.865	0.697	7.692	8.027	326
0.4%	483	0	0	POST	/api/auth/logout	80.996	0.168	4.090	9.458	0
0.4%	122	0	0	PUT	^/api/follows/[^/]+$	66.870	0.548	4.416	6.286	0
0.2%	467	0	0	GET	^/api/memories/[^/]+	36.505	0.078	0.566	1.984	663
0.2%	540	2	0	GET	^/api/posts/me/current$	33.611	0.062	0.786	1.077	647
0.1%	880	14	0	GET	^/api/posts/[^/]+$	25.603	0.029	0.184	0.996	614
0.1%	694	0	0	GET	^/api/relationships/[^/]+	22.602	0.033	0.225	1.012	1621
0.1%	63	0	0	DELETE	^/api/realmojis/[^/]+$	19.702	0.313	2.701	2.701	0
0.1%	28	0	0	DELETE	^/api/posts/[^/]+$	19.527	0.697	4.102	4.102	0
0.1%	63	0	0	PATCH	^/api/posts/[^/]+$	18.083	0.287	3.263	3.263	641
0.1%	95	0	0	PATCH	/api/me	17.377	0.183	3.262	3.262	274
0.1%	521	0	0	GET	/api/users	15.907	0.031	0.311	1.011	244
0.1%	451	0	0	GET	^/api/posts/[^/]+/realmojis$	12.545	0.028	0.189	0.254	858
0.1%	445	0	0	GET	^/api/users/[^/]+$	12.243	0.028	0.195	0.384	200
0.1%	1	0	0	POST	/api/initialize	12.125	12.125	12.125	12.125	90
0.1%	22	0	0	PUT	^/api/blocks/[^/]+$	9.452	0.430	2.529	2.529	0
0.0%	326	0	0	GET	^/api/posts/[^/]+/comments$	8.025	0.025	0.155	0.287	659
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
41.2%	172	0	0	GET	/api/notifications	1435.108	8.344	11.071	11.072	177
19.1%	191	5	0	POST	/api/posts	667.421	3.494	11.395	12.076	631
5.9%	301	0	0	POST	/api/auth/login	204.358	0.679	7.986	11.710	252
5.5%	314	9	0	POST	/api/friend-requests	190.650	0.607	5.894	6.472	113
4.8%	1164	5	0	GET	^/api/images/[^/]+	166.461	0.143	0.938	0.964	155686
4.6%	973	0	0	GET	^/api/feeds/[^/]+	160.389	0.165	0.701	1.898	3536
2.5%	46	0	0	POST	/api/auth/register	88.589	1.926	19.397	19.397	229
2.3%	362	0	0	GET	/api/memories	80.004	0.221	1.042	1.531	13583
2.0%	111	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	68.028	0.613	3.295	4.230	116
1.8%	242	0	0	GET	/api/posting-window/current	64.285	0.266	1.346	1.543	238
1.4%	53	6	0	POST	^/api/posts/[^/]+/comments$	48.415	0.913	6.400	6.400	268
1.2%	65	1	0	POST	^/api/posts/[^/]+/realmojis$	43.277	0.666	2.975	2.975	486
1.0%	20	0	0	GET	/api/me/export	34.113	1.706	2.996	2.996	2654221
0.8%	21	0	0	PUT	^/api/follows/[^/]+$	26.166	1.246	6.609	6.609	0
0.7%	160	7	0	GET	^/api/posts/[^/]+$	25.913	0.162	0.708	0.728	597
0.7%	163	0	0	GET	^/api/relationships/[^/]+	25.147	0.154	1.047	1.153	1102
0.7%	147	18	0	GET	^/api/posts/me/current$	24.230	0.165	0.472	0.488	568
0.6%	138	0	0	GET	/api/users	20.466	0.148	0.384	0.387	252
0.4%	45	0	0	POST	/api/auth/logout	15.278	0.340	3.085	3.085	0
0.4%	22	3	0	PATCH	^/api/posts/[^/]+$	15.249	0.693	3.040	3.040	558
0.4%	87	0	0	GET	^/api/users/[^/]+$	15.042	0.173	0.880	0.880	204
0.3%	52	0	0	GET	^/api/memories/[^/]+	12.070	0.232	1.958	1.958	715
0.3%	77	0	0	GET	^/api/posts/[^/]+/realmojis$	11.071	0.144	0.413	0.413	845
0.3%	66	0	0	GET	^/api/posts/[^/]+/comments$	10.382	0.157	0.681	0.681	696
0.3%	5	0	0	DELETE	^/api/posts/[^/]+$	9.194	1.839	4.390	4.390	0
0.2%	22	0	0	PATCH	/api/me	7.319	0.333	3.000	3.000	282
0.1%	22	0	0	PUT	/api/me/profile-image	4.564	0.207	0.488	0.488	0
0.1%	28	2	0	GET	/api/me	3.571	0.128	0.237	0.237	271
0.1%	4	0	0	PUT	^/api/blocks/[^/]+$	3.435	0.859	2.506	2.506	0
0.1%	10	0	0	DELETE	^/api/comments/[^/]+$	2.188	0.219	0.889	0.889	0
…（上位 30 / 全 33 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
40.2%	129	0	0	GET	/api/notifications	1038.267	8.049	9.263	9.263	178
19.7%	150	4	0	POST	/api/posts	508.252	3.388	12.023	12.405	622
7.3%	203	2	0	POST	/api/friend-requests	189.121	0.932	6.383	7.172	114
5.0%	198	0	0	POST	/api/auth/login	130.027	0.657	9.178	9.995	253
4.9%	776	4	0	GET	^/api/images/[^/]+	127.762	0.165	0.959	1.535	162690
3.3%	529	0	0	GET	^/api/feeds/[^/]+	84.290	0.159	0.505	1.350	3149
2.8%	272	0	0	GET	/api/memories	73.004	0.268	1.956	2.083	14016
2.7%	103	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	69.344	0.673	5.955	8.007	116
2.3%	189	0	0	GET	/api/posting-window/current	58.703	0.311	1.936	4.499	240
2.0%	63	6	0	POST	^/api/posts/[^/]+/realmojis$	51.756	0.822	5.090	5.090	404
1.3%	13	0	0	GET	/api/me/export	33.651	2.589	12.269	12.269	3306354
1.2%	32	0	0	PUT	^/api/follows/[^/]+$	31.193	0.975	7.180	7.180	0
1.0%	30	6	0	POST	^/api/posts/[^/]+/comments$	26.455	0.882	5.606	5.606	247
0.9%	32	0	0	POST	/api/auth/logout	23.289	0.728	6.992	6.992	0
0.8%	114	5	0	GET	^/api/posts/me/current$	20.254	0.178	0.728	1.199	611
0.7%	97	0	0	GET	/api/users	16.876	0.174	0.600	0.600	240
0.6%	36	0	0	POST	/api/auth/register	16.220	0.451	7.463	7.463	229
0.6%	106	0	0	GET	^/api/relationships/[^/]+	15.750	0.149	0.432	1.303	1277
0.6%	11	1	0	PATCH	^/api/posts/[^/]+$	15.080	1.371	4.761	4.761	598
0.5%	21	0	0	PATCH	/api/me	14.034	0.668	4.035	4.035	277
0.5%	80	3	0	GET	^/api/posts/[^/]+$	11.654	0.146	0.688	0.688	605
0.2%	39	0	0	GET	^/api/posts/[^/]+/realmojis$	5.940	0.152	0.939	0.939	974
0.2%	33	0	0	GET	^/api/users/[^/]+$	4.802	0.146	0.301	0.301	207
0.2%	24	0	0	GET	^/api/memories/[^/]+	4.663	0.194	0.662	0.662	800
0.2%	15	0	0	PUT	/api/me/profile-image	3.908	0.261	0.714	0.714	0
0.1%	27	0	0	GET	^/api/posts/[^/]+/comments$	3.275	0.121	0.179	0.179	1026
0.1%	20	1	0	GET	/api/me	2.489	0.124	0.199	0.199	270
0.1%	1	0	0	DELETE	^/api/follows/[^/]+$	1.660	1.660	1.660	1.660	0
0.1%	5	0	0	DELETE	^/api/realmojis/[^/]+$	1.440	0.288	0.345	0.345	0
0.1%	4	0	0	DELETE	^/api/posts/[^/]+$	1.299	0.325	0.375	0.375	0
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 36268 クエリ / 82 種類 / 合計 150.181s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	51.0%	235	76.529	325.66	2610.50	6562.0	125.6	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	24.8%	1576	37.247	23.63	311.00	571.4	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	8.8%	6533	13.221	2.02	12.34	10.0	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	7.1%	1662	10.688	6.43	228.16	115.6	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	1.6%	1	2.457	2457.22	2457.22	29473.0	1199543.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
6	1.3%	999	1.938	1.94	12.79	29.3	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
7	1.1%	1	1.644	1644.27	1644.27	20113.0	625795.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.8%	438	1.256	2.87	12.30	29.9	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.7%	1	1.109	1108.88	1108.88	11528.0	529757.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
10	0.3%	260	0.416	1.60	12.77	24.6	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.3%	3021	0.404	0.13	10.42	12.0	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
12	0.2%	1646	0.370	0.23	8.23	20.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
13	0.2%	2	0.338	168.99	186.59	215468.5	2534.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.2%	2	0.283	141.65	274.80	16757.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL
15	0.2%	1012	0.267	0.26	6.32	22.6	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 3
SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 7c3332686307b5e4997be1acb474422a0cd7fde7
Type: cpu
Time: 2026-10-31 11:46:34 JST
Duration: 60.04s, Total samples = 66.36s (110.52%)
Showing nodes accounting for 55.57s, 83.74% of 66.36s total
Dropped 988 nodes (cum <= 0.33s)
Showing top 40 nodes out of 227
      flat  flat%   sum%        cum   cum%
    11.73s 17.68% 17.68%     11.73s 17.68%  runtime.spanClass.sizeclass (inline)
     6.05s  9.12% 26.79%     19.70s 29.69%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     5.37s  8.09% 34.89%      8.86s 13.35%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     4.56s  6.87% 41.76%      4.76s  7.17%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     4.39s  6.62% 48.37%     17.06s 25.71%  runtime.tryDeferToSpanScan
     2.53s  3.81% 52.19%      3.65s  5.50%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     2.41s  3.63% 55.82%      2.41s  3.63%  internal/runtime/syscall/linux.Syscall6
     2.27s  3.42% 59.24%      2.27s  3.42%  runtime.memclrNoHeapPointers
     1.87s  2.82% 62.06%      1.87s  2.82%  runtime.memmove
     1.62s  2.44% 64.50%     13.06s 19.68%  runtime.scanObject
     1.45s  2.19% 66.68%      7.58s 11.42%  runtime.scanObjectsSmall
     1.41s  2.12% 68.81%      1.41s  2.12%  crypto/internal/fips140/sha256.blockAVX2
     1.31s  1.97% 70.78%      5.37s  8.09%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.29s  1.94% 72.72%      1.29s  1.94%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.05s  1.58% 74.31%      1.11s  1.67%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.93s  1.40% 75.71%      0.93s  1.40%  memeqbody
     0.64s  0.96% 76.67%      0.64s  0.96%  aeshashbody
     0.59s  0.89% 77.56%      9.10s 13.71%  runtime.scanSpan
     0.50s  0.75% 78.32%      0.50s  0.75%  example.com/isureal-company/isureal.(*UserRow).active
     0.42s  0.63% 78.95%      0.42s  0.63%  runtime.extractHeapBitsSmall
     0.39s  0.59% 79.54%      0.57s  0.86%  example.com/isureal-company/isureal.postLess
     0.37s  0.56% 80.09%      0.49s  0.74%  runtime.(*spanInlineMarkBits).tryAcquire
     0.33s   0.5% 80.59%      2.10s  3.16%  runtime.mapaccess1_faststr
     0.26s  0.39% 80.98%      0.34s  0.51%  runtime.heapArenaOf (inline)
     0.21s  0.32% 81.30%      0.78s  1.18%  database/sql.convertAssignRows
     0.16s  0.24% 81.54%      0.79s  1.19%  runtime.scanObjectSmall
     0.15s  0.23% 81.77%     12.34s 18.60%  runtime.mallocgc
     0.15s  0.23% 81.99%      1.06s  1.60%  runtime.mallocgcSmallScanNoHeader
     0.13s   0.2% 82.19%      0.57s  0.86%  runtime.sweepone
     0.12s  0.18% 82.37%      3.23s  4.87%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.12s  0.18% 82.55%      1.45s  2.19%  runtime.mapassign_faststr
     0.11s  0.17% 82.72%      1.61s  2.43%  crypto/internal/fips140/sha256.(*Digest).Write
     0.11s  0.17% 82.88%      0.46s  0.69%  runtime.wbBufFlush1
     0.10s  0.15% 83.03%      1.64s  2.47%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.09s  0.14% 83.17%      1.20s  1.81%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
     0.09s  0.14% 83.30%      0.60s   0.9%  runtime.mallocgcSmallNoscan
     0.08s  0.12% 83.42%      0.51s  0.77%  runtime.(*sweepLocked).sweep
     0.07s  0.11% 83.53%      1.89s  2.85%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.07s  0.11% 83.63%      1.76s  2.65%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.07s  0.11% 83.74%      0.97s  1.46%  example.com/isureal-company/isureal.listMemories
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 61.7% / 100% | 49% | 0.1% | 6 | 125 MB | CPU が飽和 |
| isu2 (2 cpu) | 7.4% / 28% | 0% | 0.4% | 3 | 391 MB |  |
| isu3 (2 cpu) | 3.2% / 26% | 0% | 0.4% | 3 | 220 MB |  |
| isu4 (2 cpu) | 38.6% / 99% | 8% | 19.0% | 19 | 107 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu4（run queue が積もっている、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 108% / nginx 10% / journalctl 1% |
| isu2 | isureal 13% |
| isu3 | isureal 4% |
| isu4 | postgres 35% |

## Diff from previous run

スコア: 17072 → 48796 (+31724)

`6a7a33a..718b774`

```
 conf/isu1/home/isucon/env.sh                    |   1 +
 conf/isu4/etc/systemd/system/isureal-go.service |  21 +
 webapp/go/cmd/isureal/main.go                   |   6 +
 webapp/go/export_api.go                         |   9 +
 webapp/go/image_signature.go                    |  33 +-
 webapp/go/jpegcheck/check_test.go               | 130 +++++
 webapp/go/jpegcheck/huffman.go                  | 249 +++++++++
 webapp/go/jpegcheck/reader.go                   | 669 ++++++++++++++++++++++++
 webapp/go/jpegcheck/scan.go                     | 481 +++++++++++++++++
 webapp/go/memory_api.go                         |  18 +
 webapp/go/offload.go                            | 225 ++++++++
 webapp/go/post_api.go                           |   6 +-
 webapp/go/profile_image_api.go                  |   3 +-
 webapp/go/realmoji_api.go                       |   6 +-
 14 files changed, 1840 insertions(+), 17 deletions(-)
```
