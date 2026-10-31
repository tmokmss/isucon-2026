# run 20261031-150350-0343458

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | lane2-tune |
| sha | 034345856596 |
| 所要時間 | 72s |
| 前回比 | -133663 (-100.0%) vs 20261031-145346-4535807 |
| ベスト比 | -141335 (-100.0%) vs 20261031-141027-bd0fe45 = 141335 |
| note | exp 5min + ipban bg reload + image cache 100MB + GOMEMLIMIT 2200 on 23db9c5 |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/agent-a50022a8e53f0ade9/runs/20261031-150350-0343458/bench.log
```

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目
asia	18	46	132	379	1051
europe	8	19	57	159	366
america	6	14	40	119	184

最終日（1 人 20 点）が来なかったリージョン: asia, europe, america
```

```
{"id":"20606","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"整合性の違反"}

--- stdout ---
time=2026-10-31T06:04:18.070Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:04:19.504Z level=INFO msg="user moved" user=reyomia5065 from=asia to=europe
time=2026-10-31T06:04:21.069Z level=INFO msg="europe 1 日目: 8 人"
time=2026-10-31T06:04:24.063Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:04:27.070Z level=INFO msg="asia 2 日目: 46 人"
time=2026-10-31T06:04:30.063Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T06:04:32.903Z level=INFO msg="user moved" user=reyomia5065 from=europe to=asia
time=2026-10-31T06:04:33.064Z level=INFO msg="america 2 日目: 14 人"
time=2026-10-31T06:04:36.076Z level=INFO msg="asia 3 日目: 132 人"
time=2026-10-31T06:04:37.674Z level=INFO msg="user moved" user=reyomia5065 from=asia to=europe
time=2026-10-31T06:04:38.840Z level=INFO msg="user moved" user=kake5863 from=asia to=europe
time=2026-10-31T06:04:39.063Z level=INFO msg="europe 3 日目: 57 人"
time=2026-10-31T06:04:42.065Z level=INFO msg="america 3 日目: 40 人"
time=2026-10-31T06:04:45.076Z level=INFO msg="asia 4 日目: 379 人"
time=2026-10-31T06:04:46.058Z level=INFO msg="user moved" user=tsunoben_6256 from=europe to=asia
time=2026-10-31T06:04:48.145Z level=INFO msg="europe 4 日目: 159 人"
time=2026-10-31T06:04:48.179Z level=INFO msg="user moved" user=ivymax_6601 from=asia to=america
time=2026-10-31T06:04:48.181Z level=INFO msg="user moved" user=evasolhi6293 from=europe to=asia
time=2026-10-31T06:04:48.722Z level=INFO msg="user moved" user=tsunoben_6256 from=asia to=europe
time=2026-10-31T06:04:48.919Z level=INFO msg="user moved" user=tasol.6995 from=asia to=america
time=2026-10-31T06:04:50.629Z level=INFO msg="user moved" user=evasolhi6293 from=asia to=europe
time=2026-10-31T06:04:50.940Z level=INFO msg="user moved" user=tasol.6995 from=america to=europe
time=2026-10-31T06:04:51.159Z level=INFO msg="america 4 日目: 119 人"
time=2026-10-31T06:04:52.416Z level=INFO msg="user moved" user=solku_7257 from=asia to=america
time=2026-10-31T06:04:54.238Z level=INFO msg="asia 5 日目: 1051 人"
time=2026-10-31T06:04:56.756Z level=INFO msg="user moved" user=kake5863 from=europe to=america
time=2026-10-31T06:04:57.567Z level=INFO msg="europe 5 日目: 366 人"
time=2026-10-31T06:05:00.462Z level=INFO msg="america 5 日目: 184 人"
time=2026-10-31T06:05:01.158Z level=INFO msg="グランプリ: アジア 0 人, ヨーロッパ 0 人, アメリカ 0 人"
{
  "pass": false,
  "score": 0,
  "messages": [
    "整合性の違反: ren_watanabe_3699: 履歴が初期データと違います: 画像 URL \"/api/images/7fb44c75-84c4-4c6b-ad3d-bc9875c4753d?exp=1793513400\u0026sig=Bn6eyPtR6xAcHmYILhSB1hSW34FMK-12_wGc3J3k_qk\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "ren_watanabe_3699: 履歴が初期データと違います: 画像 URL \"/api/images/7fb44c75-84c4-4c6b-ad3d-bc9875c4753d?exp=1793513400\u0026sig=Bn6eyPtR6xAcHmYILhSB1hSW34FMK-12_wGc3J3k_qk\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "jihun_park_3721: 履歴が初期データと違います: 画像 URL \"/api/images/5bfb745b-5102-4308-83de-39dd3d911d2d?exp=1793513400\u0026sig=OyNMsEgia2tqm30DQrs9CigqNha8BCrtANa0K6kBW0k\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "diego_gonzalez_4290: 履歴が初期データと違います: 画像 URL \"/api/images/a80e6004-ef64-4dc1-b47d-bdfcf96fd0fe?exp=1793513400\u0026sig=boAyb9EWe4o6Bzl-THstasy2wtNdCZPhKuLl3SHu948\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "hugo_leroy_4540: 履歴が初期データと違います: 画像 URL \"/api/images/1ffb8b6a-0532-454d-baa8-fb6598c14c73?exp=1793513400\u0026sig=MU3tYJ6MJt3UcYaJNlOaO1EmlHyI1FE2heQqvgZS3a4\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "soba_und_geschmeidiges_rust_2334: 履歴が初期データと違います: エクスポート: 画像 URL \"/api/images/07b22b13-a18b-4ce1-bb30-c6c64b2e2361?exp=1793513400\u0026sig=MabUbYp-ercpSa_PpFWsQCcGfS2LmDkCns37Fb0p9NM\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "shadingyu_menfan_1205: 履歴が初期データと違います: 画像 URL \"/api/images/e0d355fb-39e5-49a0-8659-da92d12534e8?exp=1793513400\u0026sig=-z1a_FnrWNlqhRX1uGghF-jtGY2vUnjRpbRkElqq-jI\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "gansik_yesan_4079: 履歴が初期データと違います: エクスポート: 画像 URL \"/api/images/acb76cfe-de97-4d84-874f-c0c6d734bba0?exp=1793513400\u0026sig=Do8DwujwgvjCzFJQfpPtIidTowWAk0ZfsCyw1WXUTWE\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "donabe_iwashi_53: 履歴が初期データと違います: 画像 URL \"/api/images/26ed3642-2216-4667-8d58-71af669ae66d?exp=1793513400\u0026sig=sFi3qbpCV39-ft-ZH2vePTUgQmqGpt_KwKr3XKIt9B0\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "roor.7957: フィード friends: 投稿 5db5cb50-29ca-4215-8be1-7c67cd0ec50d の画像 URL が不正です: 画像 URL \"/api/images/b2dd8b1c-c9b7-4b9d-afff-0f4557d0f71f?exp=1793513400\u0026sig=A9oguOmPWipoV_gknrOEc6uhvE9EhGdTEmTrPGK9E1E\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "netorna9183: フィード following: 投稿 43875a3d-0f38-42a6-8816-49894a9ff230 の画像 URL が不正です: 画像 URL \"/api/images/1f0c2189-5fb5-4980-afcd-b12fe8c41ceb?exp=1793513400\u0026sig=KllWdeTsBPm-MYQxRZ7MwrdHcj8sOq3M_TlCWf19f8I\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "anwa_7679: フィード friends: 投稿 5aa917a1-ebc2-42f0-92bc-baa9664825f4 の画像 URL が不正です: 画像 URL \"/api/images/09bb3215-c935-4878-a7dc-7c403fc43cc6?exp=1793513400\u0026sig=FeweDxhFXgW2Pba3hbJ_XgaxNfW9GnNE10oGvO6PSB8\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "waelmax5606: フィード friends: 投稿 174935d5-5f98-4566-9329-00658d907a4a の画像 URL が不正です: 画像 URL \"/api/images/d0cb4bdc-cb3a-4482-b315-e696952f04d9?exp=1793513400\u0026sig=xX_jjigzSUNpjs6Licqtm311VYCJX2ff9NAsQIuzP4Y\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "tsuruel.5546: フィード friends: 投稿 83470f1e-6b4f-4d09-b016-87bca5890ef7 の画像 URL が不正です: 画像 URL \"/api/images/100f92c8-cc59-4dbc-b349-189d331c01e9?exp=1793513400\u0026sig=ZDZ3dL4cnCkZAbahxflcMNE9zyBEkoXJU1XNrh8OY9s\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "mitoto-9851: 投稿 e7bd4473-d58d-4294-af3b-78bb37e35f39 の画像 URL が不正です: 画像 URL \"/api/images/89544c01-6ccb-48c1-89e1-d99352d5e84c?exp=1793513400\u0026sig=yLMHJ-6vTTRjCSTOmPfo4Td33toD3jxKmSMTEDdSMcw\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "mifulia8003: フィード friends-of-friends: 投稿 5aa917a1-ebc2-42f0-92bc-baa9664825f4 の画像 URL が不正です: 画像 URL \"/api/images/09bb3215-c935-4878-a7dc-7c403fc43cc6?exp=1793513400\u0026sig=FeweDxhFXgW2Pba3hbJ_XgaxNfW9GnNE10oGvO6PSB8\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "benyumi7636: フィード friends: 投稿 009ec879-9e5e-4235-bac6-b4c498b9bf61 の画像 URL が不正です: 画像 URL \"/api/images/0140bc6a-7e73-425f-8b88-80c5eae305a5?exp=1793513400\u0026sig=WYBKEsX2-JPAC7122YSZJiUxpB6tXbWYsL8Tu_JHxTo\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "chinoaho_7940: フィード friends: 投稿 72a2f15c-80de-4021-98a9-5b5799f85b6c の画像 URL が不正です: 画像 URL \"/api/images/7b19c886-1b9a-4a7f-a4cc-7ad359e42431?exp=1793513400\u0026sig=2WhLbr7hpa6lLS6plasyfuQedUYq2aqV9JWAredsw_E\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "hofuse.7365: 投稿 718c5e04-9b9d-48af-9bd6-dcf6097a0e55 の画像 URL が不正です: 画像 URL \"/api/images/3f3e6684-6440-41ca-ad1d-44d529c74c67?exp=1793513400\u0026sig=jWQTYROdzdkOhrV9cvqGb86mGepyYJ5T2GW30S-QiY0\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "ulel.5821: フィード friends: 投稿 ff1856bd-28c3-4bf2-8116-c76d91915586 の画像 URL が不正です: 画像 URL \"/api/images/36bff905-52f5-4918-87b8-96384969e4eb?exp=1793513400\u0026sig=X9HPbNXIidNQsVRvjTkQaJiDdoTeNVs4TUVTGYSEHRk\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "ulel.5821: フィード friends: 投稿 5afd1eae-31bc-4e19-bc1a-a720cec9c681 の画像 URL が不正です: 画像 URL \"/api/images/68b83d67-daf9-4f82-876b-d0d2ea714037?exp=1793513400\u0026sig=bLSU8TOTN2qklKKqUA2rp8TaOq0aEp4X_G8kNp69Mbw\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "totanoa.5345: フィード friends: 投稿 25720e6b-f343-4097-a788-de8d1548f5ed の画像 URL が不正です: 画像 URL \"/api/images/20075b37-1281-4357-a854-01abbd18a449?exp=1793513400\u0026sig=KANeaYAADExdLL0-c90Hpk9xNLnOpP_7ZhYRH_MklMk\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "totanoa.5345: フィード friends: 投稿 2868a2a8-b6ce-4717-a449-fa6fe05b62f4 の画像 URL が不正です: 画像 URL \"/api/images/3aebc723-81df-4d6b-b4e8-08680165c346?exp=1793513400\u0026sig=LxDRAHUkzWmxRYAl8TlDbSkria71z_j9BhC2VpXsvlU\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "mahe.5743: フィード friends: 投稿 ed8d69b5-0762-4bcb-9917-91303da8e9ae の画像 URL が不正です: 画像 URL \"/api/images/a3233f8b-58a8-408b-8af2-14e337437bc4?exp=1793513400\u0026sig=OFUu-XY2yla26wTfcQDHjooG6q6X1zqKjrCZbHdfjrI\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "nimuro_6882: フィード friends: 投稿 2718d18a-b5e6-4f28-b3bb-d3888adab24f の画像 URL が不正です: 画像 URL \"/api/images/eac0e224-a7e9-477f-b366-a381c70ab22d?exp=1793513400\u0026sig=HKUZiehhmNBWvtgoJZQ7LqLqEWgf4ZO47SaNjyt6egw\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "nimuro_6882: フィード friends: 投稿 78ce89c0-611a-445e-9dea-0e4bfd13d668 の画像 URL が不正です: 画像 URL \"/api/images/bba2f71b-d40d-4258-87ae-913ae9a0d8be?exp=1793513400\u0026sig=JFtbtC9C9AfsCX6-3VUQ8DjTFdreEU_akCVcmt_w3Nc\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "kiho6090: フィード friends: 投稿 c115cc8e-293b-4b5d-86b9-8351fbecb1a5 の画像 URL が不正です: 画像 URL \"/api/images/43efef64-01fb-4887-acb4-0c30a4b5e2a0?exp=1793513400\u0026sig=F3AdKjkYWlLasO-MfNEurB_hFi9_H_7AJx2iMu91oLk\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "kiho6090: フィード friends: 投稿 e63ce4a5-d062-42f8-a8ba-7c15511ac9ae の画像 URL が不正です: 画像 URL \"/api/images/f21fae16-3784-406a-bffc-54f3d557c299?exp=1793513400\u0026sig=wbijRzj8Pt5n413R5dYdNbKS6v4AP0SLnqlMf0jsauc\" の有効期限 2026-11-01T06:10:00Z が範囲外です",
    "kemax-5759: フィード friends: 投稿 c86e54d0-6c4d-4079-9293-cd9f0c5d6efb の画像 URL が不正です: 画像 URL \"/api/images/caf62beb-43ce-40a9-9865-7cf61a82083c?exp=1793513400\u0026sig=QtQesio6q89JPEYF3dL1Q4G2ErYutmyxhKB1F3_mblY\" の有効期限 2026-11-01T06:10:00Z が範囲外です"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
34.0%	721	9	0	GET	/api/notifications	4432.292	6.147	9.246	9.249	138
9.3%	3202	44	0	POST	/api/friend-requests	1212.601	0.379	2.541	3.225	114
8.3%	12423	153	0	GET	^/api/images/[^/]+	1087.810	0.088	1.112	2.040	97219
8.0%	5081	429	0	POST	/api/auth/login	1045.468	0.206	2.350	3.049	225
7.4%	1026	5	0	POST	/api/posts	960.646	0.936	3.952	4.449	646
5.6%	674	82	0	PUT	/api/me/profile-image	732.917	1.087	3.382	4.209	3
4.1%	2766	0	0	GET	/api/memories	538.187	0.195	1.131	1.718	14224
3.4%	1068	1	0	POST	/api/auth/register	444.098	0.416	2.288	2.354	227
3.0%	7824	0	0	GET	^/api/feeds/[^/]+	395.428	0.051	0.489	0.858	2930
2.5%	1276	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	327.580	0.257	2.492	2.543	116
2.1%	456	0	0	PATCH	/api/me	268.894	0.590	2.558	3.006	267
1.8%	518	0	0	PUT	^/api/follows/[^/]+$	237.275	0.458	2.542	2.551	0
1.5%	441	7	0	POST	^/api/posts/[^/]+/realmojis$	195.399	0.443	3.265	3.959	485
1.3%	307	0	0	GET	/api/me/export	167.026	0.544	3.017	5.636	1755756
1.1%	1609	0	0	GET	/api/posting-window/current	149.207	0.093	0.533	0.805	237
1.1%	1551	0	0	GET	/api/users	147.660	0.095	0.549	1.025	236
0.7%	371	2	0	POST	^/api/posts/[^/]+/comments$	96.341	0.260	2.315	2.483	340
0.6%	531	2	0	GET	/api/me	81.804	0.154	0.605	0.825	302
0.6%	757	0	0	POST	/api/auth/logout	73.550	0.097	1.005	2.215	0
0.5%	955	1	0	GET	^/api/posts/me/current$	70.353	0.074	0.542	0.814	650
0.5%	714	0	0	GET	^/api/memories/[^/]+	65.072	0.091	0.694	0.988	632
0.5%	1242	7	0	GET	^/api/posts/[^/]+$	58.702	0.047	0.360	0.809	619
0.4%	125	0	0	PATCH	^/api/posts/[^/]+$	48.162	0.385	2.424	2.912	651
0.3%	841	0	0	GET	^/api/relationships/[^/]+	40.583	0.048	0.468	0.790	902
0.2%	39	0	0	DELETE	^/api/posts/[^/]+$	31.517	0.808	2.485	2.485	0
0.2%	604	0	0	GET	^/api/posts/[^/]+/realmojis$	27.911	0.046	0.422	0.545	1512
0.2%	622	0	0	GET	^/api/users/[^/]+$	26.373	0.042	0.327	0.533	201
0.2%	79	0	0	DELETE	^/api/realmojis/[^/]+$	24.212	0.306	0.917	0.917	0
0.2%	500	0	0	GET	^/api/posts/[^/]+/comments$	23.511	0.047	0.368	0.758	2024
0.1%	1	0	0	POST	/api/initialize	12.548	12.548	12.548	12.548	90
…（上位 30 / 全 41 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
39.8%	255	0	0	GET	/api/notifications	1758.699	6.897	9.667	9.684	177
13.3%	2138	5	0	GET	^/api/images/[^/]+	590.023	0.276	1.393	2.087	133198
7.9%	364	14	0	POST	/api/posts	351.394	0.965	2.766	3.356	616
7.3%	356	25	0	POST	/api/auth/register	323.489	0.909	2.496	2.640	212
7.3%	836	0	0	POST	/api/auth/login	321.906	0.385	2.370	2.954	244
5.7%	838	12	0	POST	/api/friend-requests	253.510	0.303	2.137	2.467	114
4.8%	1572	0	0	GET	^/api/feeds/[^/]+	213.762	0.136	0.856	0.983	2100
1.8%	185	1	0	POST	^/api/friend-requests/[^/]+/[^/]+$	78.490	0.424	2.824	3.026	116
1.4%	458	0	0	GET	/api/memories	63.800	0.139	1.116	1.894	19206
1.4%	533	0	0	GET	/api/posting-window/current	61.138	0.115	0.657	0.824	239
1.1%	431	0	0	GET	/api/users	49.033	0.114	0.701	0.879	236
1.0%	84	5	0	POST	^/api/posts/[^/]+/realmojis$	45.302	0.539	2.101	2.101	455
0.8%	85	0	0	PUT	/api/me/profile-image	34.170	0.402	1.074	1.074	0
0.7%	118	0	0	POST	/api/auth/logout	29.692	0.252	1.251	2.350	0
0.7%	254	1	0	GET	^/api/posts/me/current$	29.264	0.115	0.732	0.735	636
0.6%	112	0	0	POST	^/api/posts/[^/]+/comments$	27.126	0.242	0.924	1.119	310
0.6%	189	10	0	GET	^/api/posts/[^/]+$	26.843	0.142	0.760	1.096	581
0.5%	163	0	0	GET	^/api/relationships/[^/]+	22.249	0.136	0.951	0.970	579
0.5%	102	0	0	PUT	^/api/follows/[^/]+$	21.118	0.207	0.766	0.813	0
0.4%	76	0	0	PATCH	/api/me	19.170	0.252	0.869	0.869	273
0.3%	99	0	0	GET	^/api/users/[^/]+$	14.133	0.143	0.859	0.859	188
0.3%	22	0	0	DELETE	^/api/comments/[^/]+$	13.564	0.617	3.125	3.125	0
0.3%	44	0	0	GET	/api/me/export	11.070	0.252	2.673	2.673	2084755
0.2%	85	0	0	GET	^/api/memories/[^/]+	10.917	0.128	1.321	1.321	973
0.2%	32	2	0	PATCH	^/api/posts/[^/]+$	10.840	0.339	1.312	1.312	624
0.2%	12	0	0	DELETE	^/api/realmojis/[^/]+$	10.234	0.853	3.532	3.532	0
0.2%	81	0	0	GET	^/api/posts/[^/]+/realmojis$	9.673	0.119	0.736	0.736	1505
0.2%	73	0	0	GET	^/api/posts/[^/]+/comments$	8.884	0.122	0.962	0.962	1230
0.2%	80	1	0	GET	/api/me	6.999	0.087	0.276	0.276	301
0.1%	5	0	0	DELETE	^/api/posts/[^/]+$	2.264	0.453	1.217	1.217	0
…（上位 30 / 全 31 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
44.6%	195	2	0	GET	/api/notifications	1197.600	6.142	10.140	10.140	148
15.0%	1245	5	0	GET	^/api/images/[^/]+	402.912	0.324	1.717	2.361	137989
10.4%	546	4	0	POST	/api/friend-requests	279.029	0.511	2.633	3.129	114
4.6%	413	0	0	POST	/api/auth/login	122.393	0.296	2.144	2.383	249
4.2%	912	0	0	GET	^/api/feeds/[^/]+	112.888	0.124	0.829	1.057	3013
3.5%	150	0	0	POST	/api/posts	93.732	0.625	2.992	3.260	652
2.3%	153	0	0	POST	^/api/friend-requests/[^/]+/[^/]+$	62.252	0.407	2.377	2.622	116
1.5%	272	0	0	GET	/api/users	39.818	0.146	0.789	0.818	231
1.2%	256	0	0	GET	/api/posting-window/current	33.368	0.130	0.751	1.052	239
1.2%	128	3	0	POST	/api/auth/register	31.799	0.248	1.407	1.621	223
1.2%	252	0	0	GET	/api/memories	31.617	0.125	1.048	1.336	23448
1.2%	69	0	0	PUT	/api/me/profile-image	31.003	0.449	2.842	2.842	0
1.1%	77	1	0	POST	^/api/posts/[^/]+/comments$	30.322	0.394	2.933	2.933	311
0.9%	145	3	0	GET	^/api/posts/me/current$	24.757	0.171	1.053	1.056	640
0.9%	65	0	0	POST	/api/auth/logout	23.487	0.361	2.384	2.384	0
0.9%	51	0	0	PATCH	/api/me	23.178	0.454	2.645	2.645	296
0.9%	58	0	0	PUT	^/api/follows/[^/]+$	22.974	0.396	2.382	2.382	0
0.7%	153	3	0	GET	^/api/posts/[^/]+$	17.974	0.117	0.848	0.915	607
0.6%	16	0	0	PATCH	^/api/posts/[^/]+$	15.118	0.945	2.523	2.523	674
0.5%	17	0	0	POST	^/api/posts/[^/]+/realmojis$	14.770	0.869	2.105	2.105	533
0.5%	37	0	0	GET	/api/me/export	12.946	0.350	2.579	2.579	3128157
0.5%	102	0	0	GET	^/api/relationships/[^/]+	12.550	0.123	0.688	0.774	975
0.3%	70	0	0	GET	/api/me	7.541	0.108	0.402	0.402	304
0.3%	60	0	0	GET	^/api/posts/[^/]+/comments$	6.811	0.114	1.041	1.041	1811
0.2%	17	0	0	DELETE	^/api/comments/[^/]+$	6.225	0.366	1.976	1.976	0
0.2%	53	0	0	GET	^/api/users/[^/]+$	5.771	0.109	0.617	0.617	188
0.2%	45	0	0	GET	^/api/memories/[^/]+	5.424	0.121	0.291	0.291	1011
0.2%	6	0	0	DELETE	^/api/realmojis/[^/]+$	5.333	0.889	1.263	1.263	0
0.2%	5	0	0	DELETE	^/api/posts/[^/]+$	4.615	0.923	2.499	2.499	0
0.1%	28	0	0	GET	^/api/posts/[^/]+/realmojis$	3.782	0.135	0.786	0.786	1256
…（上位 30 / 全 34 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 11700 クエリ / 48 種類 / 合計 120.750s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	49.3%	316	59.491	188.26	1459.18	1130.0	115.4	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	18.4%	3	22.175	7391.55	10452.45	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	13.1%	3	15.832	5277.45	7145.39	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
4	9.9%	9533	11.972	1.26	6.58	9.7	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
5	6.4%	3	7.674	2558.12	3539.62	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	0.8%	1	0.996	995.92	995.92	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
7	0.6%	4	0.686	171.53	174.56	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
8	0.3%	4	0.388	96.91	158.39	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
9	0.3%	4	0.380	94.91	166.92	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
10	0.2%	1	0.252	252.50	252.50	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
11	0.2%	1	0.207	206.83	206.83	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.154	154.28	154.28	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.114	113.92	113.92	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.074	74.18	74.18	1525.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations
15	0.1%	1	0.071	71.41	71.41	2481.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…

-- 上位のクエリ（$1 のまま EXPLAIN (GENERIC_PLAN) に渡せる）
-- rank 1
SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND created_at = k.created_at) i;
-- rank 2
SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis;
-- rank 3
SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM posts;
```

（生の出力は ptquery.txt）

## CPU profile (pprof)

```
File: isureal
Build ID: 591b59650465535111f9a420364fbf88bceacb3c
Type: cpu
Time: 2026-10-31 15:03:58 JST
Duration: 60.19s, Total samples = 46.45s (77.17%)
Showing nodes accounting for 32.67s, 70.33% of 46.45s total
Dropped 1187 nodes (cum <= 0.23s)
Showing top 40 nodes out of 301
      flat  flat%   sum%        cum   cum%
     6.44s 13.86% 13.86%      6.44s 13.86%  crypto/internal/fips140/sha256.blockAVX2
     5.51s 11.86% 25.73%      5.51s 11.86%  internal/runtime/syscall/linux.Syscall6
     2.01s  4.33% 30.05%      2.01s  4.33%  runtime.memmove
     1.95s  4.20% 34.25%      2.52s  5.43%  runtime.findObject
     1.84s  3.96% 38.21%      1.84s  3.96%  runtime.memclrNoHeapPointers
     1.82s  3.92% 42.13%      2.71s  5.83%  runtime.tryDeferToSpanScan
     1.73s  3.72% 45.86%      7.33s 15.78%  runtime.scanObject
     1.42s  3.06% 48.91%      5.61s 12.08%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.28s  2.76% 51.67%      1.28s  2.76%  aeshashbody
     1.26s  2.71% 54.38%      1.26s  2.71%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.67s  1.44% 55.82%      0.67s  1.44%  runtime.spanClass.sizeclass (inline)
     0.57s  1.23% 57.05%      2.08s  4.48%  runtime.mapaccess1_faststr
     0.46s  0.99% 58.04%      0.46s  0.99%  memeqbody
     0.44s  0.95% 58.99%      0.66s  1.42%  example.com/isureal-company/isureal.postLess
     0.37s   0.8% 59.78%      6.48s 13.95%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.36s  0.78% 60.56%      0.82s  1.77%  database/sql.convertAssignRows
     0.34s  0.73% 61.29%      7.08s 15.24%  crypto/internal/fips140/sha256.(*Digest).Write
     0.31s  0.67% 61.96%      0.31s  0.67%  runtime.futex
     0.31s  0.67% 62.63%      0.59s  1.27%  runtime.greyobject
     0.28s   0.6% 63.23%      0.28s   0.6%  runtime.typePointers.nextFast (inline)
     0.27s  0.58% 63.81%      1.19s  2.56%  runtime.mallocgcSmallScanNoHeader
     0.27s  0.58% 64.39%      0.27s  0.58%  runtime.nextFreeFast (inline)
     0.25s  0.54% 64.93%      0.29s  0.62%  runtime.typePointers.next
     0.25s  0.54% 65.47%      0.25s  0.54%  sync/atomic.(*Int32).Add
     0.23s   0.5% 65.96%      0.25s  0.54%  runtime.spanOf (inline)
     0.19s  0.41% 66.37%      1.89s  4.07%  runtime.mapassign_faststr
     0.18s  0.39% 66.76%      0.37s   0.8%  compress/flate.(*compressor).deflate
     0.18s  0.39% 67.15%      0.38s  0.82%  internal/runtime/maps.(*Iter).Next
     0.18s  0.39% 67.53%      5.36s 11.54%  runtime.mallocgc
     0.15s  0.32% 67.86%      0.94s  2.02%  runtime.scanSpan
     0.14s   0.3% 68.16%      6.70s 14.42%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.13s  0.28% 68.44%      7.37s 15.87%  example.com/isureal-company/isureal.computeImageSignature
     0.13s  0.28% 68.72%      0.34s  0.73%  internal/stringslite.Index
     0.12s  0.26% 68.98%      0.32s  0.69%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.12s  0.26% 69.24%      0.56s  1.21%  runtime.scanObjectsSmall
     0.11s  0.24% 69.47%      7.03s 15.13%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.11s  0.24% 69.71%      6.55s 14.10%  crypto/internal/fips140/sha256.block
     0.10s  0.22% 69.92%      2.85s  6.14%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.10s  0.22% 70.14%      0.45s  0.97%  runtime.mapaccess2_faststr
     0.09s  0.19% 70.33%      0.45s  0.97%  example.com/isureal-company/isureal.(*Store).searchUsersLocked
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 51.9% / 100% | 30% | 0.3% | 5 | 230 MB | run queue が積もっている |
| isu2 (2 cpu) | 31.2% / 87% | 0% | 0.5% | 4 | 902 MB |  |
| isu3 (2 cpu) | 26.4% / 74% | 0% | 0.5% | 4 | 925 MB |  |
| isu4 (2 cpu) | 48.1% / 100% | 14% | 22.5% | 15 | 107 MB | IO 待ちが重い |

**isu1（run queue が積もっている、主に isureal）、isu4（IO 待ちが重い、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 84% / nginx 13% |
| isu2 | isureal 58% / nginx 1% |
| isu3 | isureal 51% / nginx 1% |
| isu4 | isureal 26% / postgres 22% |

## Diff from previous run

スコア: 133663 → — (fail)

`4535807..0343458`

```
 conf/isu1/etc/systemd/system/isureal-go.service |  2 +-
 webapp/go/image_cache.go                        |  2 +-
 webapp/go/image_local.go                        |  7 +----
 webapp/go/image_signature.go                    |  6 +++--
 webapp/go/ip_ban.go                             | 35 +++++++++++--------------
 webapp/go/replica_server.go                     |  1 +
 webapp/go/server.go                             |  1 +
 7 files changed, 24 insertions(+), 30 deletions(-)
```
