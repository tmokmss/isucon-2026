# run 20261031-115123-0904d44

## Summary

| | |
|---|---|
| score | 58879 |
| status | pass |
| ref | try-register-only |
| sha | 0904d44b884b |
| 所要時間 | 163s |
| 前回比 | +25312 (75.4%) vs 20261031-114127-f782d77 |
| ベスト比 | +25312 (75.4%) vs 20261031-114127-f782d77 = 33567 **ベスト更新** |
| note | main + 登録の並列化 |

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	52	154	397	398	398
europe	7	19	51	99	98	97
america	6	16	40	56	56	56
```

```
{"id":"5734","status":"done","portalStatus":"finished","passed":true,"score":58879,"raw":58879,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T02:52:55.004Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T02:52:57.956Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T02:53:00.955Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T02:53:03.999Z level=INFO msg="asia 2 日目: 52 人"
time=2026-10-31T02:53:06.957Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T02:53:09.958Z level=INFO msg="america 2 日目: 16 人"
time=2026-10-31T02:53:12.990Z level=INFO msg="asia 3 日目: 154 人"
time=2026-10-31T02:53:16.110Z level=INFO msg="europe 3 日目: 51 人"
time=2026-10-31T02:53:19.088Z level=INFO msg="america 3 日目: 40 人"
time=2026-10-31T02:53:19.443Z level=INFO msg="user moved" user=hemichi_5789 from=asia to=america
time=2026-10-31T02:53:19.447Z level=INFO msg="user moved" user=anulis.6247 from=asia to=europe
time=2026-10-31T02:53:22.030Z level=INFO msg="asia 4 日目: 397 人"
time=2026-10-31T02:53:25.877Z level=INFO msg="europe 4 日目: 99 人"
time=2026-10-31T02:53:28.227Z level=INFO msg="america 4 日目: 56 人"
time=2026-10-31T02:53:28.629Z level=INFO msg="user moved" user=anulis.6247 from=europe to=asia
time=2026-10-31T02:53:31.289Z level=INFO msg="asia 5 日目: 398 人"
time=2026-10-31T02:53:34.286Z level=INFO msg="europe 5 日目: 98 人"
time=2026-10-31T02:53:37.118Z level=INFO msg="america 5 日目: 56 人"
time=2026-10-31T02:53:46.289Z level=INFO msg="america 最終日（グランプリ）: 56 人"
time=2026-10-31T02:53:46.357Z level=INFO msg="europe 最終日（グランプリ）: 97 人"
time=2026-10-31T02:53:46.561Z level=INFO msg="asia 最終日（グランプリ）: 398 人"
time=2026-10-31T02:53:54.969Z level=INFO msg="グランプリ: アジア 221 人, ヨーロッパ 41 人, アメリカ 23 人"
{
  "pass": true,
  "score": 58879,
  "messages": [
    "shitorta_5041: 投稿 b8cd4cd2-a03d-4e12-9a97-14967583420d に RealMoji を送れません: POST /api/posts/b8cd4cd2-a03d-4e12-9a97-14967583420d/realmojis: 応答がタイムアウトしました",
    "ivyivy5304: 投稿 b8cd4cd2-a03d-4e12-9a97-14967583420d にコメントできません: POST /api/posts/b8cd4cd2-a03d-4e12-9a97-14967583420d/comments: 応答がタイムアウトしました",
    "kuhine.7517: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "tahaso.7256: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "benrini_8016: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "iwashi_pilaf_stew_2048: ログインできません: POST /api/auth/login: 応答がタイムアウトしました",
    "kuni-5535: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "hihiko_8604: 認定公開アカウントをフォローできません: PUT /api/follows/09b79cf2-0a39-4a70-aacf-db51fa171b8d: 応答がタイムアウトしました",
    "hamana.8486: 認定公開アカウントをフォローできません: PUT /api/follows/a54496d4-83f8-44fc-a311-9425ecd27c92: 応答がタイムアウトしました",
    "leoka.8450: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "susone8547: 認定公開アカウントをフォローできません: PUT /api/follows/a54496d4-83f8-44fc-a311-9425ecd27c92: 応答がタイムアウトしました",
    "kuelul.8096: 認定公開アカウントをフォローできません: PUT /api/follows/a54496d4-83f8-44fc-a311-9425ecd27c92: 応答がタイムアウトしました",
    "maxsolsu_8115: 認定公開アカウントをフォローできません: PUT /api/follows/a54496d4-83f8-44fc-a311-9425ecd27c92: 応答がタイムアウトしました",
    "nosuha-6669: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "yukois8245: 認定公開アカウントをフォローできません: PUT /api/follows/a54496d4-83f8-44fc-a311-9425ecd27c92: 応答がタイムアウトしました",
    "vinisdan_8262: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました",
    "miya8640: 認定公開アカウントをフォローできません: PUT /api/follows/a54496d4-83f8-44fc-a311-9425ecd27c92: 応答がタイムアウトしました",
    "masosu.5052: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "meyo5169: 投稿できません: POST /api/posts: 応答がタイムアウトしました",
    "marmiafu.6925: 友達申請を送れません: POST /api/friend-requests: 応答がタイムアウトしました"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
31.1%	757	0	0	GET	/api/notifications	6677.156	8.821	15.316	15.432	174
19.0%	2029	33	0	POST	/api/friend-requests	4087.251	2.014	15.888	23.760	114
10.1%	964	3	0	POST	/api/posts	2163.569	2.244	8.071	23.675	645
8.5%	7720	766	0	GET	^/api/images/[^/]+	1825.241	0.236	3.931	8.706	128276
5.8%	3857	1667	0	POST	/api/auth/login	1239.942	0.321	3.741	17.708	152
5.3%	821	4	0	POST	^/api/friend-requests/[^/]+/[^/]+$	1144.040	1.393	12.371	16.958	116
5.1%	214	0	0	PUT	^/api/follows/[^/]+$	1088.827	5.088	21.961	25.006	0
3.0%	415	1	0	POST	/api/auth/register	651.527	1.570	5.414	6.447	227
2.7%	263	44	0	POST	^/api/posts/[^/]+/realmojis$	569.592	2.166	15.203	19.866	405
2.0%	203	41	0	POST	^/api/posts/[^/]+/comments$	437.220	2.154	12.324	15.305	282
1.0%	5707	0	0	GET	^/api/feeds/[^/]+	220.518	0.039	0.184	0.924	1412
1.0%	346	126	0	PUT	/api/me/profile-image	218.441	0.631	3.072	8.570	8
1.0%	2090	0	0	GET	/api/memories	213.615	0.102	0.604	1.591	9835
1.0%	1328	0	0	GET	/api/posting-window/current	207.577	0.156	1.115	3.754	236
0.7%	150	0	0	GET	/api/me/export	141.960	0.946	4.585	4.891	1893418
0.6%	353	0	0	POST	/api/auth/logout	119.685	0.339	5.015	16.084	0
0.4%	106	2	0	PATCH	^/api/posts/[^/]+$	78.564	0.741	3.531	3.918	639
0.3%	181	0	0	PATCH	/api/me	67.217	0.371	3.490	3.689	276
0.2%	914	22	0	GET	^/api/posts/me/current$	49.848	0.055	0.413	1.294	632
0.2%	44	0	0	DELETE	^/api/posts/[^/]+$	43.119	0.980	7.571	7.571	0
0.2%	968	0	0	GET	/api/users	38.320	0.040	0.206	0.518	232
0.1%	891	0	0	GET	^/api/relationships/[^/]+	29.756	0.033	0.197	0.322	953
0.1%	44	0	0	DELETE	^/api/realmojis/[^/]+$	27.652	0.628	6.074	6.074	0
0.1%	849	13	0	GET	^/api/posts/[^/]+$	25.726	0.030	0.175	0.292	604
0.1%	10	0	0	PUT	^/api/blocks/[^/]+$	17.901	1.790	15.092	15.092	0
0.1%	352	0	0	GET	^/api/memories/[^/]+	17.836	0.051	0.357	0.523	646
0.1%	1	0	0	POST	/api/initialize	12.765	12.765	12.765	12.765	90
0.1%	402	0	0	GET	^/api/users/[^/]+$	12.462	0.031	0.154	0.274	198
0.1%	1	0	0	POST	/api/me/account-deletion	10.827	10.827	10.827	10.827	0
0.0%	370	0	0	GET	^/api/posts/[^/]+/realmojis$	10.043	0.027	0.139	0.199	745
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
33.3%	216	0	0	GET	/api/notifications	1757.694	8.137	12.227	12.227	177
19.1%	508	10	0	POST	/api/friend-requests	1007.163	1.983	14.997	18.411	113
9.9%	247	1	0	POST	/api/posts	524.544	2.124	8.815	15.158	639
6.8%	425	0	0	POST	/api/auth/login	360.526	0.848	6.493	13.127	249
5.5%	135	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	289.263	2.143	15.333	15.438	116
4.7%	53	0	0	PUT	^/api/follows/[^/]+$	248.154	4.682	21.145	21.145	0
4.1%	1491	8	0	GET	^/api/images/[^/]+	216.961	0.146	1.056	1.462	150411
3.3%	1013	0	0	GET	^/api/feeds/[^/]+	172.484	0.170	0.590	1.037	1231
1.9%	469	0	0	GET	/api/memories	101.196	0.216	0.836	1.015	12536
1.7%	43	11	0	POST	^/api/posts/[^/]+/realmojis$	88.054	2.048	12.897	12.897	379
1.3%	327	0	0	GET	/api/posting-window/current	68.700	0.210	1.239	1.305	239
1.0%	96	0	0	POST	/api/auth/register	53.460	0.557	2.471	2.471	227
1.0%	46	0	0	PUT	/api/me/profile-image	52.574	1.143	14.798	14.798	0
0.9%	24	0	0	GET	/api/me/export	45.955	1.915	4.930	4.930	2413709
0.8%	229	0	0	GET	/api/users	40.129	0.175	0.766	1.377	230
0.7%	69	0	0	POST	/api/auth/logout	38.759	0.562	3.484	3.484	0
0.7%	229	13	0	GET	^/api/posts/me/current$	38.644	0.169	0.357	2.145	608
0.6%	45	0	0	PATCH	/api/me	31.899	0.709	3.607	3.607	261
0.6%	176	0	0	GET	^/api/relationships/[^/]+	31.023	0.176	1.005	1.122	901
0.4%	119	11	0	GET	^/api/posts/[^/]+$	21.604	0.182	0.555	1.005	565
0.4%	28	0	0	PATCH	^/api/posts/[^/]+$	19.304	0.689	3.850	3.850	643
0.4%	2	2	0	POST	^/api/posts/[^/]+/comments$	19.020	9.510	10.505	10.505	34
0.2%	68	0	0	GET	^/api/users/[^/]+$	11.534	0.170	0.528	0.528	204
0.2%	53	0	0	GET	^/api/memories/[^/]+	10.481	0.198	0.719	0.719	943
0.2%	46	0	0	GET	^/api/posts/[^/]+/realmojis$	8.504	0.185	1.037	1.037	820
0.1%	50	2	0	GET	/api/me	7.869	0.157	0.366	0.366	287
0.1%	43	0	0	GET	^/api/posts/[^/]+/comments$	7.551	0.176	1.225	1.225	123
0.1%	10	0	0	DELETE	^/api/posts/[^/]+$	6.844	0.684	1.630	1.630	0
0.1%	5	0	0	DELETE	^/api/realmojis/[^/]+$	3.491	0.698	2.214	2.214	0
0.0%	1	0	0	POST	/api/me/account-deletion	0.986	0.986	0.986	0.986	0
…（上位 30 / 全 31 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
35.4%	146	0	0	GET	/api/notifications	1131.479	7.750	9.281	9.283	177
17.3%	299	9	0	POST	/api/friend-requests	553.467	1.851	17.237	19.735	112
12.5%	161	2	0	POST	/api/posts	399.949	2.484	13.824	16.009	628
6.0%	224	0	0	POST	/api/auth/login	190.660	0.851	7.780	10.913	253
4.3%	75	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	137.381	1.832	10.079	10.079	116
3.9%	780	4	0	GET	^/api/images/[^/]+	125.283	0.161	0.949	1.272	170569
3.7%	32	0	0	PUT	^/api/follows/[^/]+$	118.604	3.706	25.219	25.219	0
2.5%	482	0	0	GET	^/api/feeds/[^/]+	80.077	0.166	0.668	1.246	986
1.9%	36	7	0	POST	^/api/posts/[^/]+/realmojis$	61.670	1.713	10.909	10.909	368
1.8%	276	0	0	GET	/api/memories	58.516	0.212	1.018	1.415	12883
1.6%	213	0	0	GET	/api/posting-window/current	50.147	0.235	1.499	5.729	240
1.3%	15	0	0	GET	/api/me/export	41.452	2.763	4.920	4.920	3016423
1.0%	11	3	0	POST	^/api/posts/[^/]+/comments$	31.914	2.901	12.341	12.341	234
0.9%	153	0	0	GET	/api/users	28.281	0.185	1.036	1.829	230
0.9%	54	0	0	POST	/api/auth/register	27.194	0.504	3.992	3.992	230
0.8%	143	6	0	GET	^/api/posts/me/current$	25.545	0.179	0.616	1.113	607
0.8%	28	0	0	PUT	/api/me/profile-image	25.105	0.897	4.076	4.076	0
0.7%	35	0	0	POST	/api/auth/logout	22.582	0.645	4.574	4.574	0
0.5%	27	0	0	PATCH	/api/me	17.160	0.636	2.542	2.542	281
0.5%	17	0	0	PATCH	^/api/posts/[^/]+$	15.345	0.903	3.939	3.939	618
0.4%	85	0	0	GET	^/api/relationships/[^/]+	13.364	0.157	0.712	0.712	1074
0.3%	52	1	0	GET	^/api/posts/[^/]+$	8.256	0.159	0.546	0.546	596
0.2%	27	0	0	GET	^/api/memories/[^/]+	5.418	0.201	0.541	0.541	775
0.2%	5	0	0	DELETE	^/api/posts/[^/]+$	5.004	1.001	1.736	1.736	0
0.1%	31	0	0	GET	^/api/users/[^/]+$	4.588	0.148	0.350	0.350	198
0.1%	31	0	0	GET	/api/me	4.334	0.140	0.284	0.284	296
0.1%	23	0	0	GET	^/api/posts/[^/]+/realmojis$	3.138	0.136	0.239	0.239	426
0.1%	15	0	0	GET	^/api/posts/[^/]+/comments$	2.274	0.152	0.257	0.257	333
0.1%	3	0	0	PUT	^/api/blocks/[^/]+$	1.968	0.656	1.395	1.395	0
0.0%	2	0	0	DELETE	^/api/realmojis/[^/]+$	1.381	0.691	1.176	1.176	0
…（上位 30 / 全 31 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 36697 クエリ / 81 種類 / 合計 167.228s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	41.1%	189	68.660	363.28	2754.93	6857.5	130.9	SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2)
2	36.0%	1941	60.254	31.04	348.35	601.6	1.0	INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4)
3	8.5%	1922	14.148	7.36	306.63	116.0	1.0	SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id…
4	6.5%	4871	10.837	2.22	16.26	10.6	1.0	SELECT data, media_type FROM image_objects WHERE id = $1 AND created_at = $2
5	1.6%	1367	2.758	2.02	19.08	26.6	1.0	INSERT INTO posts (id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at) VALUES …
6	1.4%	1	2.341	2341.26	2341.26	29545.0	1199738.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
7	1.0%	1	1.741	1740.87	1740.87	20250.0	626293.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
8	0.7%	1	1.107	1107.05	1107.05	11528.0	529960.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
9	0.6%	280	0.943	3.37	12.16	32.6	1.0	INSERT INTO realmojis(id,post_id,reactor_id,image_id,reaction_type,created_at) VALUES($1,$2,$3,$4,$5,$6)
10	0.4%	2772	0.698	0.25	14.39	20.4	1.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) VALUES ($1,$2,$3,$5,$4,$4)
11	0.2%	2839	0.370	0.13	11.89	11.5	1.0	INSERT INTO login_sessions (id, user_id, token_hash, created_at) VALUES ($1, $2, $3, $4)
12	0.2%	170	0.366	2.15	13.65	25.9	1.0	INSERT INTO comments(id,post_id,author_id,body,created_at) VALUES($1,$2,$3,$4,$5)
13	0.2%	2	0.341	170.30	185.52	219373.5	2581.5	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
14	0.2%	991	0.269	0.27	5.44	22.4	1.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$5)
15	0.2%	2	0.264	131.83	255.16	16763.0	0.0	SELECT n.id, n.user_id, n.posting_window_id, n.created_at FROM notifications n JOIN posting_windows w ON w.id = n.posting_window_id WHERE w.ended_at IS NULL

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT id, data FROM image_objects WHERE id = ANY($1) AND created_at = ANY($2);
-- rank 2
INSERT INTO image_objects(id,owner_user_id,data,media_type,created_at) VALUES ($1,$2,$3,$5,$4);
-- rank 3
SELECT u.secret_key_hash, u.id AS owner_id, i.media_type, COALESCE(u.profile_image_id = i.id, $3) AS is_profile, EXISTS (SELECT $4 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL) AS has_post, (SELECT p.author_id || $5 || p.image_id FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.author_id AND author…;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 2acde90f3514c98366ea696a3d85156144680946
Type: cpu
Time: 2026-10-31 11:52:36 JST
Duration: 60.18s, Total samples = 59.27s (98.48%)
Showing nodes accounting for 48.45s, 81.74% of 59.27s total
Dropped 903 nodes (cum <= 0.30s)
Showing top 40 nodes out of 230
      flat  flat%   sum%        cum   cum%
     6.86s 11.57% 11.57%     22.38s 37.76%  example.com/isureal-company/isureal/jpegcheck.(*decoder).processSOS
     6.07s 10.24% 21.82%     10.19s 17.19%  example.com/isureal-company/isureal/jpegcheck.(*decoder).decodeHuffman
     5.06s  8.54% 30.35%      5.32s  8.98%  example.com/isureal-company/isureal/jpegcheck.(*decoder).receiveExtend
     3.12s  5.26% 35.62%      4.31s  7.27%  example.com/isureal-company/isureal/jpegcheck.(*decoder).ensureNBits
     2.96s  4.99% 40.61%      2.96s  4.99%  runtime.memclrNoHeapPointers
     2.68s  4.52% 45.13%      8.45s 14.26%  runtime.scanObject
     2.37s  4.00% 49.13%      2.37s  4.00%  internal/runtime/syscall/linux.Syscall6
     2.04s  3.44% 52.57%      2.04s  3.44%  runtime.memmove
     1.92s  3.24% 55.81%      2.66s  4.49%  runtime.findObject
     1.77s  2.99% 58.80%      5.67s  9.57%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.75s  2.95% 61.75%      1.75s  2.95%  crypto/internal/fips140/sha256.blockAVX2
     1.74s  2.94% 64.69%      2.70s  4.56%  runtime.tryDeferToSpanScan
     1.48s  2.50% 67.18%      1.48s  2.50%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.28s  2.16% 69.34%      1.28s  2.16%  aeshashbody
     1.10s  1.86% 71.20%      1.17s  1.97%  example.com/isureal-company/isureal/jpegcheck.(*decoder).readByteStuffedByte
     0.87s  1.47% 72.67%      0.87s  1.47%  example.com/isureal-company/isureal.(*UserRow).active
     0.67s  1.13% 73.80%      0.67s  1.13%  runtime.spanClass.sizeclass (inline)
     0.64s  1.08% 74.88%      0.64s  1.08%  memeqbody
     0.56s  0.94% 75.82%      2.14s  3.61%  runtime.mapaccess1_faststr
     0.40s  0.67% 76.50%      0.61s  1.03%  example.com/isureal-company/isureal.postLess
     0.31s  0.52% 77.02%      0.32s  0.54%  crypto/internal/fips140/aes/gcm.gcmAesDec
     0.30s  0.51% 77.53%      0.55s  0.93%  runtime.greyobject
     0.29s  0.49% 78.02%      0.66s  1.11%  database/sql.convertAssignRows
     0.28s  0.47% 78.49%      0.35s  0.59%  runtime.spanOf (inline)
     0.26s  0.44% 78.93%      0.31s  0.52%  runtime.typePointers.next
     0.19s  0.32% 79.25%         1s  1.69%  runtime.mallocgcSmallScanNoHeader
     0.16s  0.27% 79.52%      0.54s  0.91%  runtime.scanObjectsSmall
     0.15s  0.25% 79.77%      4.51s  7.61%  runtime.mallocgc
     0.13s  0.22% 79.99%      1.99s  3.36%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.13s  0.22% 80.21%      1.15s  1.94%  github.com/jackc/pgx/v5/stdlib.(*Rows).initValueFuncs.func14
     0.13s  0.22% 80.43%      1.69s  2.85%  runtime.mapassign_faststr
     0.12s   0.2% 80.63%      1.98s  3.34%  crypto/internal/fips140/sha256.(*Digest).Write
     0.12s   0.2% 80.83%      2.86s  4.83%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.11s  0.19% 81.02%      0.44s  0.74%  runtime.mapaccess2_faststr
     0.09s  0.15% 81.17%      1.56s  2.63%  example.com/isureal-company/isureal.NewServer.searchUsersHandler.func11
     0.09s  0.15% 81.32%      0.91s  1.54%  runtime.scanSpan
     0.07s  0.12% 81.44%      0.43s  0.73%  runtime.wbBufFlush1
     0.06s   0.1% 81.54%      0.91s  1.54%  example.com/isureal-company/isureal.listMemories
     0.06s   0.1% 81.64%      4.08s  6.88%  example.com/isureal-company/isureal.loadBig
     0.06s   0.1% 81.74%      0.36s  0.61%  internal/stringslite.Index
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 62.9% / 100% | 48% | 0.1% | 6 | 360 MB | CPU が飽和 |
| isu2 (2 cpu) | 3.0% / 21% | 0% | 0.4% | 2 | 336 MB |  |
| isu3 (2 cpu) | 2.4% / 15% | 0% | 0.4% | 3 | 217 MB |  |
| isu4 (2 cpu) | 61.3% / 100% | 35% | 8.0% | 29 | 107 MB | CPU が飽和 |

**isu1（CPU が飽和、主に isureal）、isu4（CPU が飽和、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 104% / nginx 13% / journalctl 2% / systemd-journal 1% |
| isu2 | isureal 4% / nginx 1% |
| isu3 | isureal 2% / nginx 1% |
| isu4 | isureal 55% / postgres 47% |

## Diff from previous run

スコア: 33567 → 58879 (+25312)

`f782d77..0904d44`

```
 conf/isu1/home/isucon/env.sh                    |   1 +
 conf/isu4/etc/systemd/system/isureal-go.service |  21 +++
 webapp/go/cmd/isureal/main.go                   |   6 +
 webapp/go/export_api.go                         |   9 +
 webapp/go/image_signature.go                    |  33 ++--
 webapp/go/memory_api.go                         |  18 ++
 webapp/go/offload.go                            | 225 ++++++++++++++++++++++++
 webapp/go/store.go                              |  76 +++++++-
 webapp/go/user_api.go                           |  17 +-
 9 files changed, 390 insertions(+), 16 deletions(-)
```
