# run 20261031-115145-6850116

## Summary

| | |
|---|---|
| score | 61412 |
| status | pass |
| ref | offload-jpeg |
| sha | 685011625f39 |
| 所要時間 | 256s |
| 前回比 | +27845 (83.0%) vs 20261031-114127-f782d77 |
| ベスト比 | +27845 (83.0%) vs 20261031-114127-f782d77 = 33567 **ベスト更新** |
| note | offload-jpeg-to-asia2 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	48	138	378	376	229
europe	7	15	36	83	83	39
america	6	16	38	78	78	75
```

```
{"id":"5865","status":"done","portalStatus":"finished","passed":true,"score":61412,"raw":61412,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:54:40.041Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:54:42.995Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:54:45.996Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:54:49.032Z level=INFO msg="asia 2 日目: 48 人"
time=2026-10-31T02:54:51.996Z level=INFO msg="europe 2 日目: 15 人"
time=2026-10-31T02:54:54.996Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T02:54:58.033Z level=INFO msg="asia 3 日目: 138 人"
time=2026-10-31T02:55:01.017Z level=INFO msg="europe 3 日目: 36 人"
time=2026-10-31T02:55:03.672Z level=INFO msg="user moved" user=maxnoama_6136 from=asia to=europe
time=2026-10-31T02:55:03.821Z level=INFO msg="user moved" user=nehe5713 from=asia to=america
time=2026-10-31T02:55:04.008Z level=INFO msg="america 3 日目: 38 人"
time=2026-10-31T02:55:06.407Z level=INFO msg="user moved" user=leozoeto_6051 from=europe to=america
time=2026-10-31T02:55:07.035Z level=INFO msg="asia 4 日目: 378 人"
time=2026-10-31T02:55:10.032Z level=INFO msg="europe 4 日目: 83 人"
time=2026-10-31T02:55:12.921Z level=INFO msg="user moved" user=teshifu_6428 from=asia to=america
time=2026-10-31T02:55:13.129Z level=INFO msg="america 4 日目: 78 人"
time=2026-10-31T02:55:13.408Z level=INFO msg="user moved" user=masumia6475 from=asia to=america
time=2026-10-31T02:55:16.139Z level=INFO msg="asia 5 日目: 376 人"
time=2026-10-31T02:55:19.149Z level=INFO msg="europe 5 日目: 83 人"
time=2026-10-31T02:55:21.774Z level=INFO msg="user moved" user=leozoeto_6051 from=america to=asia
time=2026-10-31T02:55:22.025Z level=INFO msg="america 5 日目: 78 人"
time=2026-10-31T02:55:31.127Z level=INFO msg="asia 最終日（グランプリ）: 229 人"
time=2026-10-31T02:55:31.173Z level=INFO msg="america 最終日（グランプリ）: 75 人"
time=2026-10-31T02:55:31.207Z level=INFO msg="europe 最終日（グランプリ）: 39 人"
time=2026-10-31T02:55:40.000Z level=INFO msg="グランプリ: アジア 9 人, ヨーロッパ 0 人, アメリカ 5 人"
{
  "pass": true,
  "score": 61412,
  "messages": [
    "kehiel-7067: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "miakutsu-7082: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "zoekior-7092: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "torleo_7095: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "waevasol.7111: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "shivinya.7117: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "liazoe_7124: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "tesa.7130: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "orse-7138: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "rimaxmu_7145: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "ratote_7150: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "yumiru-7158: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "elrume_7168: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "haorchi.7170: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "kene.7183: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "marsa.7189: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "tohohe_7192: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "miaor_7204: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "numar_7208: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました",
    "hire-7213: 新しいユーザーを登録できません: POST /api/auth/register: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.2%	637	0	0	GET	/api/notifications	5940.551	9.326	14.994	15.127	175
19.3%	396	32	0	POST	/api/auth/register	3259.683	8.232	17.173	17.180	209
15.1%	811	24	0	POST	/api/posts	2548.736	3.143	11.628	12.264	636
7.9%	11784	324	0	GET	^/api/images/[^/]+	1336.942	0.113	1.820	12.055	76516
6.2%	4340	651	0	POST	/api/auth/login	1040.115	0.240	3.253	12.018	214
2.5%	297	0	0	GET	/api/me/export	421.230	1.418	13.733	15.149	1986009
2.3%	250	105	0	PUT	/api/me/profile-image	391.481	1.566	10.263	10.399	6
2.0%	1498	42	0	POST	/api/friend-requests	342.433	0.229	2.017	6.946	113
2.0%	2971	0	0	GET	/api/memories	339.652	0.114	0.578	1.404	13446
1.8%	1006	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	305.835	0.304	1.915	5.077	116
1.2%	1049	0	0	GET	/api/posting-window/current	197.398	0.188	1.855	3.462	237
0.8%	654	0	0	POST	/api/auth/logout	136.172	0.208	6.356	10.896	0
0.8%	6065	0	0	GET	^/api/feeds/[^/]+	128.200	0.021	0.115	2.442	3586
0.7%	256	5	0	POST	^/api/posts/[^/]+/realmojis$	112.906	0.441	1.986	2.154	484
0.4%	220	2	0	POST	^/api/posts/[^/]+/comments$	72.534	0.330	1.877	4.461	367
0.4%	181	0	0	PUT	^/api/follows/[^/]+$	63.572	0.351	2.236	4.152	0
0.2%	671	0	0	GET	^/api/memories/[^/]+	31.360	0.047	0.300	0.847	659
0.1%	832	0	0	GET	^/api/relationships/[^/]+	23.048	0.028	0.234	0.999	1540
0.1%	1195	4	0	GET	^/api/posts/[^/]+$	21.791	0.018	0.095	0.166	628
0.1%	122	0	0	PATCH	/api/me	17.637	0.145	1.521	1.947	288
0.1%	596	6	0	GET	^/api/posts/me/current$	17.011	0.029	0.263	2.280	649
0.1%	66	0	0	PATCH	^/api/posts/[^/]+$	16.282	0.247	2.301	2.301	664
0.1%	51	0	0	DELETE	^/api/realmojis/[^/]+$	16.063	0.315	1.615	1.615	0
0.1%	678	0	0	GET	/api/users	15.130	0.022	0.217	0.497	243
0.1%	1	0	0	POST	/api/initialize	12.982	12.982	12.982	12.982	90
0.1%	579	0	0	GET	^/api/posts/[^/]+/realmojis$	10.197	0.018	0.106	0.209	601
0.1%	585	0	0	GET	^/api/users/[^/]+$	9.846	0.017	0.086	0.126	198
0.1%	465	0	0	GET	^/api/posts/[^/]+/comments$	8.802	0.019	0.128	0.495	624
0.0%	25	0	0	DELETE	^/api/posts/[^/]+$	8.196	0.328	0.735	0.735	0
0.0%	59	0	0	DELETE	^/api/comments/[^/]+$	7.914	0.134	0.869	0.869	0
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.9%	141	0	0	GET	/api/notifications	1227.633	8.707	12.061	12.062	179
21.5%	81	8	0	POST	/api/auth/register	755.928	9.332	17.248	17.248	205
11.9%	156	7	0	POST	/api/posts	418.770	2.684	11.283	11.800	608
6.0%	1560	5	0	GET	^/api/images/[^/]+	211.413	0.136	0.956	1.450	109114
5.8%	447	0	0	POST	/api/auth/login	205.148	0.459	10.795	11.950	252
3.8%	996	0	0	GET	^/api/feeds/[^/]+	132.124	0.133	0.389	0.569	5525
2.4%	427	0	0	GET	/api/memories	83.087	0.195	0.555	1.175	14448
1.7%	163	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	60.211	0.369	2.270	2.305	116
1.7%	95	0	0	POST	/api/auth/logout	58.936	0.620	9.842	9.842	0
1.5%	191	0	0	GET	/api/posting-window/current	53.685	0.281	2.664	2.779	239
1.5%	30	0	0	GET	/api/me/export	53.681	1.789	12.096	12.096	2466979
1.5%	246	10	0	POST	/api/friend-requests	52.538	0.214	1.286	11.757	112
1.2%	93	3	0	POST	^/api/posts/[^/]+/comments$	43.701	0.470	3.209	3.209	323
0.5%	144	0	0	GET	^/api/relationships/[^/]+	18.911	0.131	0.387	0.427	1440
0.5%	135	3	0	GET	^/api/posts/[^/]+$	18.547	0.137	0.454	0.599	620
0.4%	120	0	0	GET	^/api/posts/me/current$	15.193	0.127	0.227	0.359	635
0.4%	6	0	0	DELETE	^/api/posts/[^/]+$	14.301	2.383	12.623	12.623	0
0.4%	89	0	0	GET	^/api/users/[^/]+$	13.526	0.152	0.886	0.886	210
0.4%	90	0	0	GET	^/api/posts/[^/]+/comments$	12.911	0.143	0.500	0.500	681
0.3%	75	0	0	GET	^/api/memories/[^/]+	11.783	0.157	0.381	0.381	1122
0.3%	99	0	0	GET	/api/users	11.565	0.117	0.253	0.253	243
0.3%	65	0	0	GET	^/api/posts/[^/]+/realmojis$	9.298	0.143	0.384	0.384	310
0.2%	34	0	0	PUT	^/api/follows/[^/]+$	7.635	0.225	0.986	0.986	0
0.2%	9	1	0	POST	^/api/posts/[^/]+/realmojis$	6.492	0.721	1.991	1.991	360
0.2%	19	0	0	DELETE	^/api/comments/[^/]+$	6.486	0.341	1.200	1.200	0
0.2%	19	0	0	PUT	/api/me/profile-image	6.372	0.335	1.873	1.873	0
0.1%	16	0	0	PATCH	^/api/posts/[^/]+$	3.137	0.196	0.397	0.397	635
0.1%	23	1	0	GET	/api/me	3.043	0.132	0.531	0.531	281
0.1%	15	0	0	PATCH	/api/me	2.136	0.142	0.316	0.316	269
0.0%	3	0	0	PUT	^/api/blocks/[^/]+$	0.328	0.109	0.114	0.114	0

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
32.4%	184	1	0	GET	/api/notifications	1243.126	6.756	9.138	9.138	179
17.4%	225	7	0	POST	/api/posts	668.664	2.972	11.201	11.448	629
14.1%	74	0	0	POST	/api/auth/register	539.937	7.296	15.121	15.121	229
7.3%	400	6	0	POST	/api/friend-requests	279.294	0.698	11.108	11.513	114
5.6%	1450	6	0	GET	^/api/images/[^/]+	214.088	0.148	0.991	1.396	143083
5.2%	397	0	0	POST	/api/auth/login	198.857	0.501	9.852	11.895	252
3.4%	922	0	0	GET	^/api/feeds/[^/]+	129.498	0.140	0.439	1.411	4464
2.4%	446	0	0	GET	/api/memories	93.043	0.209	0.835	1.654	13913
2.2%	304	0	0	GET	/api/posting-window/current	85.395	0.281	2.760	3.729	240
1.7%	179	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	66.659	0.372	2.259	2.368	115
1.4%	27	0	0	GET	/api/me/export	52.116	1.930	10.637	10.637	2675402
1.0%	63	0	0	POST	/api/auth/logout	38.068	0.604	11.870	11.870	0
0.7%	198	0	0	GET	^/api/relationships/[^/]+	27.882	0.141	0.572	1.028	1314
0.7%	181	0	0	GET	/api/users	27.145	0.150	0.789	3.284	243
0.6%	166	7	0	GET	^/api/posts/me/current$	24.745	0.149	0.682	1.187	617
0.6%	48	2	0	POST	^/api/posts/[^/]+/realmojis$	21.565	0.449	2.510	2.510	439
0.5%	153	2	0	GET	^/api/posts/[^/]+$	20.245	0.132	0.396	0.425	621
0.5%	31	3	0	POST	^/api/posts/[^/]+/comments$	17.509	0.565	2.606	2.606	286
0.4%	56	0	0	PUT	^/api/follows/[^/]+$	17.168	0.307	2.295	2.295	0
0.3%	103	0	0	GET	^/api/posts/[^/]+/realmojis$	13.165	0.128	0.426	0.600	722
0.2%	67	0	0	GET	^/api/users/[^/]+$	9.291	0.139	0.401	0.401	198
0.2%	58	0	0	GET	^/api/memories/[^/]+	8.812	0.152	0.317	0.317	1017
0.2%	40	0	0	PUT	/api/me/profile-image	8.671	0.217	1.468	1.468	0
0.2%	59	0	0	GET	^/api/posts/[^/]+/comments$	7.230	0.123	0.228	0.228	723
0.2%	31	0	0	PATCH	/api/me	6.869	0.222	2.375	2.375	271
0.1%	44	1	0	GET	/api/me	5.356	0.122	0.371	0.371	294
0.1%	12	3	0	PATCH	^/api/posts/[^/]+$	4.161	0.347	1.907	1.907	498
0.1%	6	0	0	DELETE	^/api/posts/[^/]+$	4.055	0.676	1.699	1.699	0
0.1%	10	0	0	DELETE	^/api/realmojis/[^/]+$	3.103	0.310	0.501	0.501	0
0.1%	10	0	0	DELETE	^/api/comments/[^/]+$	2.039	0.204	0.518	0.518	0
…（上位 30 / 全 33 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 48791 クエリ / 82 種類 / 合計 399.570s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	58.6%	354	234.108	661.32	5690.43	7658.0	133.1	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	22.6%	1664	90.137	54.17	512.62	646.3	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	8.8%	9654	35.165	3.64	40.03	9.7	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
4	6.1%	2543	24.480	9.63	318.06	119.3	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
5	1.0%	1155	3.830	3.32	90.09	27.2	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	0.6%	1	2.349	2349.09	2349.09	29566.0	1199628.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	0.4%	1	1.742	1741.88	1741.88	20308.0	626477.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.3%	1	1.115	1114.71	1114.71	11528.0	529766.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	0.3%	305	1.107	3.63	14.81	32.3	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.2%	336	0.641	1.91	17.71	22.7	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
11	0.2%	4533	0.622	0.14	36.20	11.6	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
12	0.1%	2	0.582	290.86	419.19	223824.0	2633.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
13	0.1%	2070	0.527	0.25	57.67	20.3	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
14	0.1%	1309	0.348	0.27	18.03	20.7	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.1%	4535	0.300	0.07	15.02	6.1	0.3	UPDATE login_sessions SET active = $3, invalidated_at = $2, reason = $4 WHERE user_id = $1 AND active = $5

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
Build ID: 6a0dec4b183171bc290422e193a2896a86c5b41b
Type: cpu
Time: 2026-10-31 11:54:21 JST
Duration: 60.04s, Total samples = 53.17s (88.55%)
Showing nodes accounting for 41.27s, 77.62% of 53.17s total
Dropped 1008 nodes (cum <= 0.27s)
Showing top 40 nodes out of 253
      flat  flat%   sum%        cum   cum%
     4.70s  8.84%  8.84%     15.09s 28.38%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     4.26s  8.01% 16.85%      6.79s 12.77%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     3.45s  6.49% 23.34%      3.58s  6.73%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     2.92s  5.49% 28.83%      2.92s  5.49%  internal/runtime/syscall/linux.Syscall6
     2.75s  5.17% 34.00%      2.75s  5.17%  crypto/internal/fips140/sha256.blockAVX2
     2.46s  4.63% 38.63%      2.46s  4.63%  runtime.memclrNoHeapPointers
     2.21s  4.16% 42.79%      7.78s 14.63%  runtime.scanObject
     1.95s  3.67% 46.45%      2.47s  4.65%  runtime.findObject
     1.88s  3.54% 49.99%      2.63s  4.95%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     1.80s  3.39% 53.38%      1.80s  3.39%  runtime.memmove
     1.60s  3.01% 56.39%      5.68s 10.68%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.60s  3.01% 59.39%      2.54s  4.78%  runtime.tryDeferToSpanScan
     1.37s  2.58% 61.97%      1.37s  2.58%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.25s  2.35% 64.32%      1.25s  2.35%  aeshashbody
     0.72s  1.35% 65.68%      0.74s  1.39%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.67s  1.26% 66.94%      0.67s  1.26%  example.com/isureal-company/isureal.(*UserRow).active
     0.64s  1.20% 68.14%      0.64s  1.20%  memeqbody
     0.63s  1.18% 69.32%      0.63s  1.18%  runtime.spanClass.sizeclass (inline)
     0.42s  0.79% 70.11%      1.95s  3.67%  runtime.mapaccess1_faststr
     0.38s  0.71% 70.83%      0.62s  1.17%  example.com/isureal-company/isureal.postLess
     0.37s   0.7% 71.53%      0.73s  1.37%  database/sql.convertAssignRows
     0.33s  0.62% 72.15%      0.33s  0.62%  runtime.nextFreeFast (inline)
     0.31s  0.58% 72.73%      0.64s  1.20%  runtime.greyobject
     0.29s  0.55% 73.27%      0.29s  0.55%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.26s  0.49% 73.76%      1.12s  2.11%  runtime.mallocgcSmallScanNoHeader
     0.23s  0.43% 74.20%      0.27s  0.51%  runtime.(*mspan).markBitsForIndex (inline)
     0.20s  0.38% 74.57%      4.88s  9.18%  runtime.mallocgc
     0.16s   0.3% 74.87%         3s  5.64%  crypto/internal/fips140/sha256.(*Digest).Write
     0.16s   0.3% 75.17%      3.09s  5.81%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.15s  0.28% 75.46%      2.95s  5.55%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.14s  0.26% 75.72%      0.35s  0.66%  compress/flate.(*compressor).deflate
     0.13s  0.24% 75.96%      0.45s  0.85%  runtime.mapaccess2_faststr
     0.13s  0.24% 76.21%      1.78s  3.35%  runtime.mapassign_faststr
     0.13s  0.24% 76.45%      0.39s  0.73%  runtime.scanObjectsSmall
     0.12s  0.23% 76.68%      0.71s  1.34%  runtime.scanSpan
     0.11s  0.21% 76.89%      1.23s  2.31%  github.com/jackc/pgx/v5/pgproto3.(*chunkReader).Next
     0.11s  0.21% 77.09%      1.20s  2.26%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.10s  0.19% 77.28%      1.16s  2.18%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.09s  0.17% 77.45%      1.77s  3.33%  example.com/isureal-company/isureal.listMemories
     0.09s  0.17% 77.62%      0.52s  0.98%  runtime.mallocgcSmallNoscan
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 53.2% / 100% | 36% | 0.2% | 6 | 182 MB | CPU が飽和 |
| isu2 (2 cpu) | 6.6% / 27% | 0% | 0.4% | 3 | 376 MB |  |
| isu3 (2 cpu) | 4.8% / 22% | 0% | 0.5% | 4 | 162 MB |  |
| isu4 (2 cpu) | 60.7% / 100% | 47% | 9.0% | 23 | 109 MB | CPU が飽和 |

**isu1（CPU が飽和、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 88% / nginx 11% / journalctl 1% / systemd-journal 1% |
| isu2 | isureal 11% |
| isu3 | isureal 6% / nginx 1% |
| isu4 | postgres 50% / isureal 49% |

## Diff from previous run

スコア: 33567 → 61412 (+27845)

`f782d77..6850116`

```
 conf/isu1/home/isucon/env.sh                    |   1 +
 conf/isu4/etc/systemd/system/isureal-go.service |  21 ++
 webapp/go/cmd/isureal/main.go                   |   6 +
 webapp/go/export_api.go                         |   9 +
 webapp/go/image_signature.go                    |  33 ++-
 webapp/go/memory_api.go                         |  18 ++
 webapp/go/offload.go                            | 258 ++++++++++++++++++++++++
 webapp/go/post_api.go                           |   4 +-
 webapp/go/profile_image_api.go                  |   3 +-
 webapp/go/realmoji_api.go                       |   4 +-
 webapp/go/store.go                              |  76 ++++++-
 11 files changed, 412 insertions(+), 21 deletions(-)
```
