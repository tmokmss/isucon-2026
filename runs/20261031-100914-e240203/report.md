# run 20261031-100914-e240203

## Summary

| | |
|---|---|
| score | 2674 |
| status | pass |
| ref | HEAD |
| sha | e240203cde40 |
| 所要時間 | 137s |
| 前回比 | +2674 (—) vs 20261031-100757-ff46021 |
| ベスト比 | —（まだ pass したランが無い） |
| note | 初期状態 |

## Bench

```
{"id":"151","status":"done","portalStatus":"finished","passed":true,"score":2674,"raw":2674,"deduction":0,"reason":"実行終了"}

--- stdout ---
time=2026-10-31T01:09:58.980Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T01:09:59.137Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T01:09:59.268Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T01:10:08.342Z level=INFO msg="asia 2 日目: 18 人"
time=2026-10-31T01:10:10.229Z level=INFO msg="europe 2 日目: 7 人"
time=2026-10-31T01:10:10.422Z level=INFO msg="america 2 日目: 6 人"
time=2026-10-31T01:10:18.425Z level=INFO msg="europe 3 日目: 7 人"
time=2026-10-31T01:10:18.458Z level=INFO msg="asia 3 日目: 18 人"
time=2026-10-31T01:10:19.992Z level=INFO msg="america 3 日目: 6 人"
time=2026-10-31T01:10:27.946Z level=INFO msg="europe 4 日目: 7 人"
time=2026-10-31T01:10:29.184Z level=INFO msg="america 4 日目: 6 人"
time=2026-10-31T01:10:34.574Z level=INFO msg="asia 5 日目: 18 人"
time=2026-10-31T01:10:36.047Z level=INFO msg="europe 5 日目: 7 人"
time=2026-10-31T01:10:40.591Z level=INFO msg="america 5 日目: 6 人"
time=2026-10-31T01:10:49.146Z level=INFO msg="asia 最終日（グランプリ）: 17 人"
time=2026-10-31T01:10:50.043Z level=INFO msg="europe 最終日（グランプリ）: 7 人"
time=2026-10-31T01:10:51.467Z level=INFO msg="グランプリ: アジア 7 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": true,
  "score": 2674,
  "messages": [
    "asia の投稿期間を開けません: POST /api/admin/posting-window: 応答がタイムアウトしました",
    "leoruho_5081: 投稿が開いていない投稿期間 4edf340f-d1b4-4185-8326-4491ed619e69 に入りました",
    "seoyeon_lee_20: 投稿が開いていない投稿期間 4edf340f-d1b4-4185-8326-4491ed619e69 に入りました",
    "hina_nakamura_38: 投稿が開いていない投稿期間 4edf340f-d1b4-4185-8326-4491ed619e69 に入りました",
    "suta-5044: 投稿が開いていない投稿期間 4edf340f-d1b4-4185-8326-4491ed619e69 に入りました",
    "yomi.5031: 投稿が開いていない投稿期間 4edf340f-d1b4-4185-8326-4491ed619e69 に入りました",
    "noaro.5058: 投稿が開いていない投稿期間 4edf340f-d1b4-4185-8326-4491ed619e69 に入りました",
    "yachisa_5021: 投稿が開いていない投稿期間 4edf340f-d1b4-4185-8326-4491ed619e69 に入りました",
    "yuse_5050: 投稿が開いていない投稿期間 4edf340f-d1b4-4185-8326-4491ed619e69 に入りました",
    "susol_5063: 投稿が開いていない投稿期間 4edf340f-d1b4-4185-8326-4491ed619e69 に入りました"
  ]
}

--- stderr ---
```

## Access log (alp)

**注意: metrics.alpMatches が空で URI が正規化されていない。** ID を含むパスが別々の行に割れているので、集計は当てにならない。

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
49.9%	108	0	0	GET	/api/notifications	1130.378	10.466	15.010	16.017	177
7.4%	64	0	0	POST	/api/posts	168.586	2.634	8.790	8.790	601
5.1%	18	0	0	POST	/api/admin/posting-window	115.806	6.434	9.616	9.616	177
4.7%	34	1	0	POST	/api/auth/register	106.111	3.121	9.420	9.420	236
4.1%	136	125	10	PUT	/api/me/profile-image	92.656	0.681	3.035	3.088	27
1.8%	184	0	0	GET	/api/memories	41.057	0.223	0.888	0.971	13827
1.6%	69	4	0	POST	/api/friend-requests	36.341	0.527	7.582	7.582	110
1.4%	792	595	0	POST	/api/auth/login	31.996	0.040	0.222	0.496	85
1.2%	127	0	0	GET	/api/feeds/friends	28.008	0.221	0.520	0.686	582
1.2%	15	0	0	GET	/api/me/export	26.111	1.741	5.432	5.432	1944312
0.6%	52	0	0	GET	/api/feeds/friends-of-friends	13.975	0.269	0.530	0.530	1196
0.4%	49	0	0	GET	/api/feeds/public	9.669	0.197	0.516	0.516	1169
0.3%	1	0	0	POST	/api/friend-requests/32d8644c-c4ed-45e9-bff8-d4e5e737dfb8/accept	7.272	7.272	7.272	7.272	115
0.3%	40	0	0	GET	/api/feeds/following	7.208	0.180	0.430	0.430	195
0.3%	83	0	0	GET	/api/posting-window/current	6.262	0.075	0.340	0.340	237
0.2%	24	0	0	GET	/api/users	4.749	0.198	0.570	0.570	264
0.2%	1	0	0	POST	/api/initialize	4.747	4.747	4.747	4.747	69
0.2%	31	31	0	GET	/api/images/777d5e88-9eb9-4084-ac7d-b4c3743ff421	4.571	0.147	0.379	0.379	27
0.2%	2	1	0	GET	/api/images/8ba9f25e-29b7-400e-8db4-c5fff7d29dcb	3.781	1.891	3.475	3.475	223024
0.2%	49	1	0	GET	/api/posts/me/current	3.669	0.075	0.469	0.469	590
0.2%	32	0	0	GET	/api/relationships/friends	3.647	0.114	0.473	0.473	5873
0.1%	1	0	0	GET	/api/images/f45f3780-1220-4e69-97a0-d7302d9c31f8	3.046	3.046	3.046	3.046	3013
0.1%	1	0	0	GET	/api/images/0f154bf3-8b2e-4999-87f6-20358b073c23	3.043	3.043	3.043	3.043	3609
0.1%	1	0	0	GET	/api/images/21199c3c-2194-4de1-a69d-6ce5e55d14b6	2.959	2.959	2.959	2.959	3602
0.1%	1	0	0	POST	/api/friend-requests/6cc25d4d-3641-4e56-86ca-b59970748e6f/accept	2.923	2.923	2.923	2.923	116
0.1%	1	0	0	GET	/api/images/eb5bda9a-906a-4641-91be-3420de51125a	2.916	2.916	2.916	2.916	3546
0.1%	1	0	0	POST	/api/friend-requests/c0baf152-0f84-4471-a2d4-5802e584a5e7/accept	2.898	2.898	2.898	2.898	116
0.1%	1	0	0	GET	/api/images/ca87bd73-d589-4350-b29d-905d98d88d63	2.894	2.894	2.894	2.894	3655
0.1%	1	0	0	GET	/api/images/98424d21-00e1-4b0f-bf09-5b20c9238ada	2.858	2.858	2.858	2.858	4791
0.1%	1	0	0	GET	/api/images/e1144f6f-5f0a-4371-a963-8f41b7db3857	2.851	2.851	2.851	2.851	4029
…（上位 30 / 全 703 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
46.0%	43	0	0	GET	/api/notifications	452.721	10.528	15.306	15.306	177
9.3%	25	0	0	POST	/api/posts	91.183	3.647	8.336	8.336	600
7.2%	29	3	0	POST	/api/friend-requests	70.892	2.445	9.017	9.017	107
5.0%	52	0	0	GET	/api/memories	49.656	0.955	2.118	2.118	5063
4.1%	28	0	0	GET	/api/feeds/friends	40.570	1.449	2.273	2.273	373
2.9%	21	0	0	POST	/api/auth/login	28.864	1.374	2.300	2.300	248
2.0%	4	0	0	PUT	/api/follows/d9f109b1-95c1-40d7-a469-b1178d887bec	19.637	4.909	6.511	6.511	0
1.8%	5	0	0	POST	/api/auth/register	18.200	3.640	5.930	5.930	231
1.7%	28	0	0	GET	/api/posting-window/current	16.463	0.588	1.383	1.383	239
1.5%	23	10	0	GET	/api/posts/me/current	14.429	0.627	1.366	1.366	349
1.2%	13	0	0	GET	/api/users	11.820	0.909	1.316	1.316	211
0.6%	1	0	0	POST	/api/me/account-deletion	5.911	5.911	5.911	5.911	0
0.6%	13	0	0	GET	/api/relationships/friends	5.753	0.443	0.625	0.625	1252
0.6%	2	0	0	GET	/api/me/export	5.474	2.737	3.237	3.237	2979408
0.5%	1	0	0	GET	/api/images/3a1440b8-d520-44e5-86a0-4e461e7a8609	4.613	4.613	4.613	4.613	2954
0.5%	1	0	0	GET	/api/images/e988b4fb-3420-4aaf-a68f-1262e0dc74b6	4.596	4.596	4.596	4.596	3897
0.5%	1	0	0	GET	/api/images/c5a1fead-c4c1-43c0-890f-951f491ed58f	4.545	4.545	4.545	4.545	3336
0.4%	1	0	0	POST	/api/friend-requests/d28f8f8e-20f7-4ffa-bcd2-1de03a2fb678/accept	3.745	3.745	3.745	3.745	116
0.3%	7	0	0	GET	/api/relationships/incoming-requests	3.423	0.489	0.710	0.710	68
0.3%	1	0	0	GET	/api/images/f9ff798f-b3f8-4137-971f-6c62b3a8a5da	3.400	3.400	3.400	3.400	3140
0.3%	2	1	0	GET	/api/images/d2f251ef-fd5d-47e7-9134-e0d524a2cc71	3.323	1.661	2.501	2.501	258921
0.3%	2	0	0	GET	/api/feeds/public	3.154	1.577	1.711	1.711	1216
0.3%	1	0	0	GET	/api/images/f13e7391-0d14-416c-80ea-eff7a6c0f4f3	3.080	3.080	3.080	3.080	3591
0.3%	1	0	0	GET	/api/images/23df1d17-d9d4-4574-8d34-1967bc386bfb	2.938	2.938	2.938	2.938	3689
0.3%	3	0	0	POST	/api/auth/logout	2.900	0.967	1.147	1.147	0
0.3%	1	0	0	GET	/api/images/78d2671e-11ea-4163-b11e-eb3397d2dd59	2.883	2.883	2.883	2.883	2277
0.3%	1	0	0	GET	/api/images/56d02a70-e1e2-42c5-b8c1-7774bb621174	2.701	2.701	2.701	2.701	3255
0.3%	2	0	0	GET	/api/feeds/friends-of-friends	2.620	1.310	1.396	1.396	656
0.3%	1	0	0	GET	/api/images/f9081e42-eb06-4943-ac81-042ec400830d	2.586	2.586	2.586	2.586	4640
0.3%	1	0	0	GET	/api/images/fb9e3d64-ef24-4d84-bf8d-f3bd23faf67b	2.571	2.571	2.571	2.571	3166
…（上位 30 / 全 111 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
45.8%	36	0	0	GET	/api/notifications	394.575	10.960	14.184	14.184	182
8.5%	22	0	0	POST	/api/posts	72.761	3.307	5.935	5.935	591
7.3%	32	3	0	POST	/api/friend-requests	63.014	1.969	6.834	6.834	107
4.5%	47	0	0	GET	/api/memories	39.112	0.832	1.863	1.863	4767
3.6%	21	0	0	GET	/api/feeds/friends	31.114	1.482	1.916	1.916	512
3.5%	5	0	0	POST	/api/auth/register	30.405	6.081	8.097	8.097	232
3.1%	19	0	0	POST	/api/auth/login	26.286	1.383	2.007	2.007	245
2.0%	25	0	0	GET	/api/posting-window/current	17.578	0.703	1.485	1.485	240
1.7%	4	0	0	PUT	/api/follows/eb8e1fbf-b54d-4618-9146-58f9f4727edb	15.060	3.765	5.137	5.137	0
1.4%	5	0	0	GET	/api/me/export	12.139	2.428	3.860	3.860	2124740
1.4%	14	0	0	GET	/api/users	12.010	0.858	1.260	1.260	251
1.4%	20	6	0	GET	/api/posts/me/current	11.757	0.588	0.922	0.922	422
0.5%	3	0	0	GET	/api/feeds/public	4.398	1.466	1.585	1.585	627
0.5%	1	0	0	GET	/api/images/67ef4dbf-57af-4e20-8baa-e15ff1eafd38	4.057	4.057	4.057	4.057	2613
0.5%	12	0	0	GET	/api/relationships/friends	4.049	0.337	0.548	0.548	1217
0.5%	1	0	0	GET	/api/images/181b9c41-0745-4ac6-a42c-a053d6cf0b3f	4.023	4.023	4.023	4.023	4296
0.5%	1	0	0	GET	/api/images/b9c4f8b2-9d99-4d0e-9e15-673691c99722	4.011	4.011	4.011	4.011	3400
0.4%	1	0	0	GET	/api/images/68c93fef-e91c-4850-b8a7-8fa07d02b28b	3.685	3.685	3.685	3.685	3155
0.4%	1	0	0	GET	/api/images/63bb7b29-b103-45f4-b265-448b9b877f9f	3.655	3.655	3.655	3.655	2765
0.4%	1	0	0	GET	/api/images/8af3f04c-0425-415c-a3f2-d37495cf4097	3.649	3.649	3.649	3.649	3512
0.4%	1	0	0	GET	/api/images/684fda99-7c79-4518-976e-384e281da4ce	3.551	3.551	3.551	3.551	3353
0.4%	1	0	0	GET	/api/images/8f2f093b-53f3-440c-8225-a4ad414f0268	3.471	3.471	3.471	3.471	2708
0.4%	1	0	0	POST	/api/friend-requests/bf4392b5-14bd-4db0-80db-4b8ccd5430d6/accept	3.252	3.252	3.252	3.252	116
0.4%	2	0	0	GET	/api/feeds/following	3.155	1.578	1.721	1.721	29
0.3%	1	0	0	GET	/api/images/0cf38190-22fc-44c6-8c8c-e28f1ca3e68a	2.744	2.744	2.744	2.744	3483
0.3%	3	0	0	POST	/api/auth/logout	2.678	0.893	0.935	0.935	0
0.3%	1	0	0	GET	/api/images/9a13cd44-5912-4c73-b2bd-0abd1813bf34	2.593	2.593	2.593	2.593	3677
0.3%	1	0	0	POST	/api/friend-requests/5b33b32c-1dab-499b-b5ba-6640e60a30ff/accept	2.532	2.532	2.532	2.532	116
0.3%	1	0	0	GET	/api/images/0851362b-57b1-4df3-b38a-5a7e8863478f	2.454	2.454	2.454	2.454	3599
0.3%	1	0	0	GET	/api/images/d1bc3e18-c943-4800-855f-11045bceb117	2.422	2.422	2.422	2.422	3806
…（上位 30 / 全 97 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 26091 クエリ / 168 種類 / 合計 1071.907s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	36.0%	749	386.391	515.88	4013.76	7607.6	1.0	SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXI…
2	18.1%	62	194.389	3135.31	9340.13	5.3	0.0	LOCK TABLE users IN SHARE ROW EXCLUSIVE MODE
3	17.2%	129	184.658	1431.46	8110.54	3.4	0.9	SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE
4	8.9%	663	95.806	144.50	7242.45	4.3	1.0	SELECT id FROM users WHERE id = $1 AND status = $2 FOR NO KEY UPDATE
5	3.9%	18	42.267	2348.14	7566.21	33027.2	1639.1	INSERT INTO notifications (id, user_id, posting_window_id, created_at) SELECT gen_random_uuid()::text, id, $2, $3 FROM users WHERE region = $1 AND status = $4
6	3.0%	22	32.159	1461.76	4406.41	13509.5	132.4	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
7	2.2%	526	24.025	45.67	162.17	610.5	11.6	SELECT user_low_id, user_high_id FROM friendships WHERE user_low_id = $1 OR user_high_id = $1
8	1.2%	18	13.055	725.29	4185.09	13.3	0.8	UPDATE posting_windows SET ended_at = $2 WHERE region = $1 AND ended_at IS NULL
9	1.0%	963	10.801	11.22	48.79	87.0	1.0	SELECT u.secret_key_hash FROM image_objects i JOIN users u ON u.id = i.owner_user_id WHERE i.id = $1 AND u.status = $2
10	0.9%	111	9.889	89.09	335.79	980.9	1.0	INSERT INTO image_objects (id, owner_user_id, data, media_type, created_at) VALUES ($1, $2, $3, $5, $4)
11	0.9%	51	9.473	185.74	435.01	225.5	4917.7	SELECT id, username, secret_key_hash, display_name, bio, profile_image_id, region, account_type, status FROM users WHERE status=$4 AND id<>$1 AND ($2=$5 OR (use…
12	0.6%	129	6.235	48.34	129.87	594.7	0.1	SELECT id FROM friend_requests WHERE status=$3 AND ((requester_id=$1 AND recipient_id=$2) OR (requester_id=$2 AND recipient_id=$1))
13	0.5%	161	5.599	34.77	1247.35	4.3	1.0	SELECT id FROM users WHERE id = $1 FOR NO KEY UPDATE
14	0.5%	310	5.397	17.41	169.56	116.8	5.6	SELECT p.id, p.author_id, p.posting_window_id, p.image_id, p.caption, p.visibility, p.location_kind, p.latitude, p.longitude, p.locality, p.posted_at, w.on_time…
15	0.5%	61	5.396	88.46	255.88	610.1	165.3	SELECT user_low_id, user_high_id FROM friendships WHERE user_low_id IN ($1, $2, $3, $4, $5, $6, $7) OR user_high_id IN ($8, $9, $10, $11, $12, $13, $14)

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.data, i.media_type FROM image_objects i JOIN users u ON u.id = i.owner_user_id AND u.status = $2 WHERE i.id = $1 AND ( u.profile_image_id = i.id OR EXISTS ( SELECT $3 FROM posts p WHERE p.image_id = i.id AND p.author_id = u.id AND p.deleted_at IS NULL ) OR EXISTS ( SELECT $4 FROM realmojis r JOIN posts p ON p.id = r.post_id AND p.deleted_at IS NULL JOIN users author ON author.id = p.autho…;
-- rank 2
LOCK TABLE users IN SHARE ROW EXCLUSIVE MODE;
-- rank 3
SELECT id, local_date, starts_at FROM posting_windows WHERE region = $1 AND ended_at IS NULL FOR UPDATE;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
未取得: metrics.pprofEnabled = false（アプリに pprof を仕込んでいない）
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 19.3% / 92% | 2% | 0.1% | 5 | 1856 MB | run queue が積もっている |
| isu2 (2 cpu) | 1.7% / 12% | 0% | 0.0% | 3 | 3228 MB |  |
| isu3 (2 cpu) | 2.1% / 14% | 0% | 0.0% | 3 | 3226 MB |  |
| isu4 (2 cpu) | 59.5% / 100% | 50% | 6.3% | 29 | 124 MB | CPU が飽和 |

**isu1（run queue が積もっている、主に isureal）、isu4（CPU が飽和、主に postgres）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 32% / nginx 5% |
| isu2 | isureal 3% |
| isu3 | isureal 4% |
| isu4 | postgres 60% |

## Diff from previous run

スコア: 0 → 2674 (+2674)

`ff46021..e240203`

```
 webapp/._perl                                             | Bin 163 -> 0 bytes
 webapp/._rust                                             | Bin 163 -> 0 bytes
 webapp/.gitignore                                         |   1 +
 webapp/go/._Makefile                                      | Bin 163 -> 0 bytes
 webapp/go/._admin_session.go                              | Bin 163 -> 0 bytes
 webapp/go/._post.go                                       | Bin 163 -> 0 bytes
 webapp/go/._post_api.go                                   | Bin 163 -> 0 bytes
 webapp/go/._realmoji.go                                   | Bin 163 -> 0 bytes
 webapp/go/._relationship_list_api.go                      | Bin 163 -> 0 bytes
 webapp/go/._user_read_api.go                              | Bin 163 -> 0 bytes
 webapp/nodejs/src/._account_deletion_api.ts               | Bin 163 -> 0 bytes
 webapp/nodejs/src/._action_log.ts                         | Bin 163 -> 0 bytes
 webapp/nodejs/src/._export_api.ts                         | Bin 163 -> 0 bytes
 webapp/nodejs/src/._main.ts                               | Bin 163 -> 0 bytes
 webapp/nodejs/src/._post_api.ts                           | Bin 163 -> 0 bytes
 webapp/nodejs/src/._server.ts                             | Bin 163 -> 0 bytes
 webapp/nodejs/src/._session.ts                            | Bin 163 -> 0 bytes
 webapp/perl/lib/Isureal/._APIError.pm                     | Bin 163 -> 0 bytes
 webapp/perl/lib/Isureal/._Comment.pm                      | Bin 163 -> 0 bytes
 webapp/perl/lib/Isureal/._Config.pm                       | Bin 163 -> 0 bytes
 webapp/perl/lib/Isureal/._Util.pm                         | Bin 163 -> 0 bytes
 webapp/perl/lib/Isureal/Controller/._AdminSession.pm      | Bin 163 -> 0 bytes
 webapp/perl/lib/Isureal/Controller/._Image.pm             | Bin 163 -> 0 bytes
 webapp/perl/lib/Isureal/Controller/._ProfileImage.pm      | Bin 163 -> 0 bytes
 webapp/php/._php-fpm.conf                                 | Bin 163 -> 0 bytes
 webapp/php/._phpcs.xml                                    | Bin 163 -> 0 bytes
 webapp/php/._src                                          | Bin 163 -> 0 bytes
 webapp/php/src/._Jpeg.php                                 | Bin 163 -> 0 bytes
 webapp/php/src/._Time.php                                 | Bin 163 -> 0 bytes
 webapp/php/src/Handler/._InitializeHandler.php            | Bin 163 -> 0 bytes
 webapp/php/src/Handler/._PostingWindowHandler.php         | Bin 163 -> 0 bytes
 webapp/php/src/Handler/._RelationshipHandler.php          | Bin 163 -> 0 bytes
 webapp/php/src/Http/._Middleware                          | Bin 163 -> 0 bytes
 .../php/src/Http/Middleware/._AdminSessionMiddleware.php  | Bin 163 -> 0 bytes
 webapp/php/src/Http/Middleware/._RegionMiddleware.php     | Bin 163 -> 0 bytes
 webapp/php/src/Model/._PostingWindow.php                  | Bin 163 -> 0 bytes
 webapp/php/src/Repository/._AdminSessionRepository.php    | Bin 163 -> 0 bytes
 webapp/python/._.python-version                           | Bin 163 -> 0 bytes
 webapp/python/._gunicorn.conf.py                          | Bin 163 -> 0 bytes
 webapp/python/._pyproject.toml                            | Bin 163 -> 0 bytes
 webapp/python/isureal/.___init__.py                       | Bin 163 -> 0 bytes
 webapp/python/isureal/._config.py                         | Bin 163 -> 0 bytes
 webapp/python/isureal/._db.py                             | Bin 163 -> 0 bytes
 webapp/python/isureal/models/._post.py                    | Bin 163 -> 0 bytes
 webapp/python/isureal/routers/._initialize.py             | Bin 163 -> 0 bytes
 webapp/python/isureal/routers/._memories.py               | Bin 163 -> 0 bytes
 webapp/python/isureal/routers/._notifications.py          | Bin 163 -> 0 bytes
 webapp/ruby/._Gemfile.lock                                | Bin 163 -> 0 bytes
 webapp/ruby/._Makefile                                    | Bin 163 -> 0 bytes
 webapp/ruby/lib/._isureal                                 | Bin 163 -> 0 bytes
 webapp/ruby/lib/isureal/._action_log.rb                   | Bin 163 -> 0 bytes
 webapp/ruby/lib/isureal/._admin_user_api.rb               | Bin 163 -> 0 bytes
 webapp/ruby/lib/isureal/._comment.rb                      | Bin 163 -> 0 bytes
 webapp/ruby/lib/isureal/._initialize.rb                   | Bin 163 -> 0 bytes
 webapp/ruby/lib/isureal/._realmoji_api.rb                 | Bin 163 -> 0 bytes
 webapp/ruby/lib/isureal/._region.rb                       | Bin 163 -> 0 bytes
 webapp/ruby/lib/isureal/._relationship.rb                 | Bin 163 -> 0 bytes
 webapp/rust/._.gitignore                                  | Bin 163 -> 0 bytes
 webapp/rust/._Cargo.toml                                  | Bin 163 -> 0 bytes
 webapp/rust/src/._admin_posting_window_api.rs             | Bin 163 -> 0 bytes
…（11 行省略。全文は git diff）
```
