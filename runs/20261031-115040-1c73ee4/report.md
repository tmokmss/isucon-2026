# run 20261031-115040-1c73ee4

## Summary

| | |
|---|---|
| score | 60201 |
| status | pass |
| ref | origin/main |
| sha | 1c73ee49c821 |
| 所要時間 | 101s |
| 前回比 | +26634 (79.3%) vs 20261031-114127-f782d77 |
| ベスト比 | +26634 (79.3%) vs 20261031-114127-f782d77 = 33567 **ベスト更新** |
| note | main: +#66 署名を asia2 へ +#69 メモリ削減 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	52	153	394	394	214
europe	7	20	53	104	103	56
america	6	16	36	58	59	51
```

```
{"id":"5622","status":"done","portalStatus":"finished","passed":true,"score":60201,"raw":60201,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:51:02.263Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:51:05.217Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:51:08.217Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:51:11.255Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T02:51:13.861Z level=INFO msg="user moved" user=niku5316 from=asia to=europe
time=2026-10-31T02:51:14.217Z level=INFO msg="europe 2 日目: 20 人"
time=2026-10-31T02:51:17.217Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T02:51:20.244Z level=INFO msg="asia 3 日目: 153 人"
time=2026-10-31T02:51:23.341Z level=INFO msg="europe 3 日目: 53 人"
time=2026-10-31T02:51:23.826Z level=INFO msg="user moved" user=nirois5630 from=asia to=europe
time=2026-10-31T02:51:26.328Z level=INFO msg="america 3 日目: 36 人"
time=2026-10-31T02:51:27.319Z level=INFO msg="user moved" user=yonisa-5900 from=asia to=america
time=2026-10-31T02:51:29.403Z level=INFO msg="asia 4 日目: 394 人"
time=2026-10-31T02:51:33.154Z level=INFO msg="europe 4 日目: 104 人"
time=2026-10-31T02:51:35.262Z level=INFO msg="america 4 日目: 58 人"
time=2026-10-31T02:51:35.920Z level=INFO msg="user moved" user=nirois5630 from=europe to=america
time=2026-10-31T02:51:38.410Z level=INFO msg="asia 5 日目: 394 人"
time=2026-10-31T02:51:41.319Z level=INFO msg="europe 5 日目: 103 人"
time=2026-10-31T02:51:44.459Z level=INFO msg="america 5 日目: 59 人"
time=2026-10-31T02:51:53.317Z level=INFO msg="europe 最終日（グランプリ）: 56 人"
time=2026-10-31T02:51:53.464Z level=INFO msg="asia 最終日（グランプリ）: 214 人"
time=2026-10-31T02:51:53.507Z level=INFO msg="america 最終日（グランプリ）: 51 人"
time=2026-10-31T02:52:02.222Z level=INFO msg="グランプリ: アジア 7 人, ヨーロッパ 11 人, アメリカ 8 人"
{
  "pass": true,
  "score": 60201,
  "messages": [
    "meeva_7069: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "rakiis_7078: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "rose7079: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kika-7091: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "tsudantor.7099: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "miaorru.7113: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kitsu-7120: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "hedanki-7121: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "maxhoul-7128: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kaso_7135: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kone7145: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "zoewa.7155: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ivyni-7161: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "oryu_7163: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "tsuhi.7170: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kiorko_7181: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "soleo.7188: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "evameel.7195: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "furame_7203: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "tayo.7214: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.9%	635	3	0	GET	/api/notifications	5768.062	9.084	15.058	15.157	173
21.8%	413	38	0	POST	/api/auth/register	3810.401	9.226	18.338	18.718	206
12.5%	794	22	0	POST	/api/posts	2191.795	2.760	9.719	10.311	632
9.1%	11150	617	0	GET	^/api/images/[^/]+	1584.716	0.142	2.137	9.788	77967
6.8%	4253	909	0	POST	/api/auth/login	1187.267	0.279	4.343	9.727	201
2.5%	1370	37	0	POST	/api/friend-requests	439.368	0.321	2.861	5.715	113
2.4%	247	116	0	PUT	/api/me/profile-image	428.516	1.735	9.501	10.186	7
2.2%	1025	7	0	POST	^/api/friend-requests/[^/]+/[^/]+$	384.969	0.376	2.877	5.759	115
1.9%	2761	0	0	GET	/api/memories	331.536	0.120	0.689	1.540	13364
1.7%	235	0	0	GET	/api/me/export	304.734	1.297	7.168	11.400	2010039
1.2%	1040	0	0	GET	/api/posting-window/current	208.417	0.200	1.818	3.251	237
1.0%	323	13	0	POST	^/api/posts/[^/]+/realmojis$	166.915	0.517	2.528	2.975	450
0.9%	309	18	0	POST	^/api/posts/[^/]+/comments$	156.564	0.507	2.922	3.911	333
0.8%	5666	0	0	GET	^/api/feeds/[^/]+	137.777	0.024	0.187	1.477	3294
0.6%	677	0	0	POST	/api/auth/logout	105.082	0.155	3.351	9.398	0
0.4%	169	0	0	PUT	^/api/follows/[^/]+$	69.326	0.410	3.849	5.143	0
0.2%	609	0	0	GET	^/api/memories/[^/]+	33.218	0.055	0.345	1.437	720
0.1%	39	0	0	DELETE	^/api/posts/[^/]+$	21.009	0.539	1.787	1.787	0
0.1%	72	0	0	DELETE	^/api/realmojis/[^/]+$	20.075	0.279	1.181	1.181	0
0.1%	915	5	0	GET	^/api/posts/[^/]+$	18.120	0.020	0.141	1.225	630
0.1%	604	6	0	GET	^/api/posts/me/current$	16.904	0.028	0.228	1.164	646
0.1%	56	1	0	PATCH	^/api/posts/[^/]+$	16.655	0.297	1.684	1.684	646
0.1%	626	0	0	GET	/api/users	14.995	0.024	0.165	0.520	245
0.1%	100	0	0	PATCH	/api/me	14.562	0.146	1.218	1.219	278
0.1%	781	0	0	GET	^/api/relationships/[^/]+	11.934	0.015	0.093	0.652	1568
0.1%	1	0	0	POST	/api/initialize	11.701	11.701	11.701	11.701	90
0.1%	484	0	0	GET	^/api/posts/[^/]+/realmojis$	9.687	0.020	0.124	0.166	585
0.1%	472	0	0	GET	^/api/users/[^/]+$	8.792	0.019	0.082	0.179	213
0.0%	60	0	0	DELETE	^/api/comments/[^/]+$	7.697	0.128	0.772	0.772	0
0.0%	408	0	0	GET	^/api/posts/[^/]+/comments$	7.694	0.019	0.090	0.130	862
…（上位 30 / 全 40 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.8%	200	0	0	GET	/api/notifications	1666.741	8.334	11.981	11.987	179
18.5%	100	14	0	POST	/api/auth/register	858.920	8.589	17.845	17.918	196
11.8%	233	4	0	POST	/api/posts	550.365	2.362	9.556	10.358	648
6.5%	513	0	0	POST	/api/auth/login	302.870	0.590	7.197	9.995	253
5.3%	1761	9	0	GET	^/api/images/[^/]+	245.362	0.139	0.976	1.640	134494
3.9%	1263	0	0	GET	^/api/feeds/[^/]+	182.969	0.145	0.448	0.724	4412
2.9%	328	6	0	POST	/api/friend-requests	134.342	0.410	3.633	6.087	113
2.4%	530	0	0	GET	/api/memories	111.566	0.211	0.901	2.937	14180
2.0%	136	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	94.342	0.694	5.339	7.082	115
1.5%	293	0	0	GET	/api/posting-window/current	69.015	0.236	1.878	1.970	238
1.2%	88	8	0	POST	^/api/posts/[^/]+/comments$	56.783	0.645	3.885	3.885	343
1.1%	94	0	0	POST	/api/auth/logout	52.317	0.557	8.028	8.028	0
1.1%	34	0	0	GET	/api/me/export	51.111	1.503	3.485	3.485	2584601
0.9%	52	6	0	POST	^/api/posts/[^/]+/realmojis$	43.717	0.841	4.540	4.540	406
0.9%	278	4	0	GET	^/api/posts/[^/]+$	39.568	0.142	0.445	0.555	631
0.6%	185	3	0	GET	^/api/posts/me/current$	29.852	0.161	0.565	0.696	646
0.6%	207	0	0	GET	^/api/relationships/[^/]+	28.831	0.139	0.449	0.559	980
0.5%	154	0	0	GET	/api/users	21.857	0.142	0.575	0.918	258
0.4%	123	0	0	GET	^/api/users/[^/]+$	17.249	0.140	0.560	0.605	220
0.3%	88	0	0	GET	^/api/memories/[^/]+	15.052	0.171	1.691	1.691	692
0.3%	98	0	0	GET	^/api/posts/[^/]+/realmojis$	13.538	0.138	0.377	0.377	691
0.3%	90	0	0	GET	^/api/posts/[^/]+/comments$	12.910	0.143	0.553	0.553	1052
0.2%	9	0	0	DELETE	^/api/posts/[^/]+$	10.775	1.197	4.181	4.181	0
0.2%	37	0	0	PUT	/api/me/profile-image	8.445	0.228	0.502	0.502	0
0.1%	22	0	0	PATCH	^/api/posts/[^/]+$	6.891	0.313	1.955	1.955	645
0.1%	40	0	0	PUT	^/api/follows/[^/]+$	5.938	0.148	0.396	0.396	0
0.1%	42	2	0	GET	/api/me	5.033	0.120	0.202	0.202	293
0.1%	25	0	0	PATCH	/api/me	4.355	0.174	0.359	0.359	282
0.1%	4	0	0	DELETE	^/api/follows/[^/]+$	3.750	0.938	1.510	1.510	0
0.1%	10	0	0	DELETE	^/api/realmojis/[^/]+$	3.440	0.344	0.614	0.614	0
…（上位 30 / 全 34 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
36.5%	150	0	0	GET	/api/notifications	1126.365	7.509	9.231	9.332	178
14.8%	178	3	0	POST	/api/posts	457.017	2.568	9.555	9.724	630
10.6%	56	3	0	POST	/api/auth/register	327.114	5.841	16.073	16.073	216
6.9%	301	7	0	POST	/api/friend-requests	212.183	0.705	8.012	9.251	113
5.5%	1152	4	0	GET	^/api/images/[^/]+	169.717	0.147	0.951	1.341	159983
4.9%	305	0	0	POST	/api/auth/login	152.587	0.500	7.868	8.177	252
4.1%	829	0	0	GET	^/api/feeds/[^/]+	126.114	0.152	0.584	1.234	4571
2.7%	133	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	84.076	0.632	3.889	4.200	115
2.5%	372	0	0	GET	/api/memories	77.295	0.208	0.988	1.486	14356
2.0%	234	0	0	GET	/api/posting-window/current	61.550	0.263	1.823	2.367	240
1.6%	21	0	0	GET	/api/me/export	48.446	2.307	8.464	8.464	2694725
0.9%	73	0	0	POST	/api/auth/logout	27.614	0.378	3.324	3.324	0
0.8%	38	4	0	POST	^/api/posts/[^/]+/realmojis$	25.950	0.683	2.551	2.551	444
0.8%	138	3	0	GET	^/api/posts/me/current$	25.486	0.185	1.848	3.295	626
0.7%	159	11	0	GET	^/api/posts/[^/]+$	22.839	0.144	0.437	0.442	587
0.7%	40	3	0	POST	^/api/posts/[^/]+/comments$	22.296	0.557	3.894	3.894	328
0.7%	135	0	0	GET	/api/users	21.210	0.157	1.026	1.683	248
0.7%	148	0	0	GET	^/api/relationships/[^/]+	20.329	0.137	0.583	1.247	1268
0.3%	22	0	0	PUT	/api/me/profile-image	10.590	0.481	3.210	3.210	0
0.3%	73	0	0	GET	^/api/users/[^/]+$	10.274	0.141	0.349	0.349	214
0.3%	71	0	0	GET	^/api/posts/[^/]+/realmojis$	9.344	0.132	0.372	0.372	951
0.2%	45	0	0	GET	^/api/memories/[^/]+	7.662	0.170	0.446	0.446	879
0.2%	51	0	0	GET	^/api/posts/[^/]+/comments$	7.070	0.139	0.449	0.449	1395
0.2%	35	0	0	PUT	^/api/follows/[^/]+$	6.740	0.193	1.285	1.285	0
0.2%	20	1	0	PATCH	^/api/posts/[^/]+$	5.978	0.299	1.603	1.603	618
0.1%	4	0	0	DELETE	^/api/posts/[^/]+$	4.208	1.052	2.233	2.233	0
0.1%	2	0	0	DELETE	^/api/follows/[^/]+$	4.134	2.067	2.660	2.660	0
0.1%	23	0	0	PATCH	/api/me	3.742	0.163	0.781	0.781	277
0.1%	26	1	0	GET	/api/me	3.016	0.116	0.220	0.220	282
0.1%	8	0	0	DELETE	^/api/realmojis/[^/]+$	2.217	0.277	0.448	0.448	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 46481 クエリ / 79 種類 / 合計 354.821s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	48.8%	290	173.274	597.50	4164.32	7464.3	136.0	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	32.5%	1764	115.275	65.35	567.15	632.8	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	8.7%	8831	30.786	3.49	43.37	9.7	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	5.6%	2410	19.812	8.22	247.24	114.2	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	1.3%	1184	4.595	3.88	46.21	27.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.6%	1	2.229	2228.62	2228.62	29516.0	1199557.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.5%	1	1.632	1632.27	1632.27	20192.0	625672.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.4%	390	1.438	3.69	14.78	30.4	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
9	0.2%	1	0.786	786.42	786.42	11528.0	529826.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
10	0.2%	408	0.708	1.73	11.39	22.5	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.1%	4162	0.529	0.13	22.34	11.4	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
12	0.1%	1940	0.491	0.25	19.61	20.7	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
13	0.1%	1256	0.346	0.28	18.25	21.8	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
14	0.1%	2	0.321	160.72	162.21	214533.5	2523.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
15	0.1%	2	0.268	133.84	259.19	16758.5	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

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
Build ID: d6610af1bc5062cfd6cbd4fd656e5e73f6eeae2a
Type: cpu
Time: 2026-10-31 11:50:45 JST
Duration: 60.02s, Total samples = 61.91s (103.16%)
Showing nodes accounting for 50.22s, 81.12% of 61.91s total
Dropped 1007 nodes (cum <= 0.31s)
Showing top 40 nodes out of 231
      flat  flat%   sum%        cum   cum%
     8.15s 13.16% 13.16%     25.30s 40.87%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     6.92s 11.18% 24.34%     11.46s 18.51%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     5.45s  8.80% 33.14%      5.68s  9.17%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     3.58s  5.78% 38.93%      4.74s  7.66%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     2.62s  4.23% 43.16%      2.62s  4.23%  internal/runtime/syscall/linux.Syscall6
     2.49s  4.02% 47.18%      2.49s  4.02%  runtime.memclrNoHeapPointers
     2.22s  3.59% 50.77%      8.22s 13.28%  runtime.scanObject
     1.92s  3.10% 53.87%      2.39s  3.86%  runtime.findObject
     1.92s  3.10% 56.97%      1.92s  3.10%  runtime.memmove
     1.69s  2.73% 59.70%      2.99s  4.83%  runtime.tryDeferToSpanScan
     1.58s  2.55% 62.25%      5.64s  9.11%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.43s  2.31% 64.56%      1.43s  2.31%  aeshashbody
     1.41s  2.28% 66.84%      1.41s  2.28%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.31s  2.12% 68.95%      1.31s  2.12%  crypto/internal/fips140/sha256.blockAVX2
     1.13s  1.83% 70.78%      1.16s  1.87%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.85s  1.37% 72.15%      0.85s  1.37%  runtime.spanClass.sizeclass (inline)
     0.68s  1.10% 73.25%      0.68s  1.10%  memeqbody
     0.57s  0.92% 74.17%      0.57s  0.92%  example.com/isureal-company/isureal.(*UserRow).active
     0.39s  0.63% 74.80%      0.57s  0.92%  example.com/isureal-company/isureal.postLess
     0.39s  0.63% 75.43%      2.10s  3.39%  runtime.mapaccess1_faststr
     0.37s   0.6% 76.03%      0.37s   0.6%  runtime.nextFreeFast (inline)
     0.34s  0.55% 76.58%      0.73s  1.18%  runtime.greyobject
     0.31s   0.5% 77.08%      0.31s   0.5%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.27s  0.44% 77.52%      0.57s  0.92%  database/sql.convertAssignRows
     0.26s  0.42% 77.94%      0.33s  0.53%  runtime.typePointers.next
     0.23s  0.37% 78.31%      1.10s  1.78%  runtime.mallocgcSmallScanNoHeader
     0.20s  0.32% 78.63%      0.38s  0.61%  runtime.scanObjectSmall
     0.19s  0.31% 78.94%      2.83s  4.57%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.18s  0.29% 79.23%      0.53s  0.86%  runtime.mapaccess2_faststr
     0.16s  0.26% 79.49%      1.93s  3.12%  runtime.mapassign_faststr
     0.15s  0.24% 79.73%      0.54s  0.87%  runtime.scanObjectsSmall
     0.13s  0.21% 79.94%      1.54s  2.49%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.11s  0.18% 80.12%      0.51s  0.82%  encoding/json.structEncoder.encode
     0.10s  0.16% 80.28%      1.18s  1.91%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.10s  0.16% 80.44%      4.69s  7.58%  runtime.mallocgc
     0.10s  0.16% 80.60%      1.03s  1.66%  runtime.scanSpan
     0.08s  0.13% 80.73%      1.77s  2.86%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.08s  0.13% 80.86%      1.63s  2.63%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.08s  0.13% 80.99%      1.25s  2.02%  example.com/isureal-company/isureal.NewServer.getFeedHandler.func50
     0.08s  0.13% 81.12%      3.70s  5.98%  example.com/isureal-company/isureal.loadBig
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 57.7% / 100% | 41% | 0.4% | 6 | 178 MB | CPU が飽和 |
| isu2 (2 cpu) | 4.7% / 23% | 0% | 0.5% | 2 | 334 MB |  |
| isu3 (2 cpu) | 3.6% / 22% | 0% | 0.5% | 2 | 207 MB |  |
| isu4 (2 cpu) | 60.1% / 100% | 51% | 4.5% | 28 | 106 MB | CPU が飽和 |

**isu1（CPU が飽和、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 97% / nginx 11% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 7% / nginx 1% |
| isu3 | isureal 5% / nginx 1% |
| isu4 | isureal 57% / postgres 49% |

## Diff from previous run

スコア: 33567 → 60201 (+26634)

`f782d77..1c73ee4`

```
 conf/isu1/home/isucon/env.sh                    |   1 +
 conf/isu4/etc/systemd/system/isureal-go.service |  21 +++
 webapp/go/cmd/isureal/main.go                   |   6 +
 webapp/go/export_api.go                         |   9 +
 webapp/go/image_signature.go                    |  33 ++--
 webapp/go/memory_api.go                         |  18 ++
 webapp/go/offload.go                            | 225 ++++++++++++++++++++++++
 webapp/go/store.go                              |  76 +++++++-
 8 files changed, 376 insertions(+), 13 deletions(-)
```
