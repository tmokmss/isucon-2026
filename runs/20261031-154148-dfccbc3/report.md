# run 20261031-154148-dfccbc3

## Summary

| | |
|---|---|
| score | — (fail) |
| status | fail |
| ref | presign-next-minute |
| sha | dfccbc319c40 |
| 所要時間 | 279s |
| 前回比 | 0 (—) vs 20261031-153557-106c0a8 |
| ベスト比 | -191816 (-100.0%) vs 20261031-151339-5a230a4 = 191816 |
| note | sig cache generations + presign + JPEG prefers asia2 (recursion fixed) |

**失敗の理由**

```
ベンチが失敗しました (status=finished)
  詳細: /Users/tmokmss/projects/isucon-harness/.claude/worktrees/agent-aff106d587106aae5/runs/20261031-154148-dfccbc3/bench.log
```

## Bench

```
リージョン別・日別の利用者数（人数が伸びないリージョンは、その日にエラーかタイムアウトが出ている）

region	1日目	2日目	3日目	4日目	5日目	最終日
asia	18	50	144	423	1063	1062
europe	7	19	48	135	368	368
america	6	20	62	186	303	303
```

```
{"id":"24270","status":"failed","portalStatus":"finished","passed":false,"score":0,"raw":0,"deduction":0,"reason":"走行後の整合性チェックに失敗しました"}

--- stdout ---
time=2026-10-31T06:45:25.166Z level=INFO msg="asia 1 日目: 18 人"
time=2026-10-31T06:45:28.163Z level=INFO msg="europe 1 日目: 7 人"
time=2026-10-31T06:45:31.161Z level=INFO msg="america 1 日目: 6 人"
time=2026-10-31T06:45:34.171Z level=INFO msg="asia 2 日目: 50 人"
time=2026-10-31T06:45:36.044Z level=INFO msg="user moved" user=zoeru_5271 from=asia to=america
time=2026-10-31T06:45:37.061Z level=INFO msg="user moved" user=orme.5331 from=asia to=europe
time=2026-10-31T06:45:37.102Z level=INFO msg="user moved" user=zoeru_5271 from=america to=europe
time=2026-10-31T06:45:37.161Z level=INFO msg="europe 2 日目: 19 人"
time=2026-10-31T06:45:40.161Z level=INFO msg="america 2 日目: 20 人"
time=2026-10-31T06:45:43.171Z level=INFO msg="asia 3 日目: 144 人"
time=2026-10-31T06:45:43.535Z level=INFO msg="user moved" user=tahate.5583 from=asia to=america
time=2026-10-31T06:45:44.243Z level=INFO msg="user moved" user=tahate.5583 from=america to=europe
time=2026-10-31T06:45:46.161Z level=INFO msg="europe 3 日目: 48 人"
time=2026-10-31T06:45:47.867Z level=INFO msg="user moved" user=zoeru_5271 from=europe to=asia
time=2026-10-31T06:45:49.088Z level=INFO msg="user moved" user=orme.5331 from=europe to=asia
time=2026-10-31T06:45:49.161Z level=INFO msg="america 3 日目: 62 人"
time=2026-10-31T06:45:52.173Z level=INFO msg="asia 4 日目: 423 人"
time=2026-10-31T06:45:54.841Z level=INFO msg="user moved" user=zoeru_5271 from=asia to=america
time=2026-10-31T06:45:55.386Z level=INFO msg="europe 4 日目: 135 人"
time=2026-10-31T06:45:55.430Z level=INFO msg="user moved" user=orme.5331 from=asia to=europe
time=2026-10-31T06:45:56.449Z level=INFO msg="user moved" user=ivykechi.7425 from=asia to=europe
time=2026-10-31T06:45:56.478Z level=INFO msg="user moved" user=rehi6671 from=asia to=america
time=2026-10-31T06:45:57.413Z level=INFO msg="user moved" user=mamoya.7376 from=asia to=europe
time=2026-10-31T06:45:57.646Z level=INFO msg="user moved" user=numarne_7851 from=asia to=europe
time=2026-10-31T06:45:57.798Z level=INFO msg="user moved" user=tahate.5583 from=europe to=asia
time=2026-10-31T06:45:58.033Z level=INFO msg="user moved" user=evasu.8341 from=asia to=america
time=2026-10-31T06:45:58.146Z level=INFO msg="user moved" user=fulia8413 from=asia to=america
time=2026-10-31T06:45:58.203Z level=INFO msg="america 4 日目: 186 人"
time=2026-10-31T06:46:01.201Z level=INFO msg="asia 5 日目: 1063 人"
time=2026-10-31T06:46:06.163Z level=INFO msg="europe 5 日目: 368 人"
time=2026-10-31T06:46:07.380Z level=INFO msg="america 5 日目: 303 人"
time=2026-10-31T06:46:16.894Z level=INFO msg="america 最終日（グランプリ）: 303 人"
time=2026-10-31T06:46:17.390Z level=INFO msg="europe 最終日（グランプリ）: 368 人"
time=2026-10-31T06:46:17.794Z level=INFO msg="asia 最終日（グランプリ）: 1062 人"
time=2026-10-31T06:46:25.256Z level=INFO msg="グランプリ: アジア 457 人, ヨーロッパ 173 人, アメリカ 199 人"
{
  "pass": false,
  "score": 0,
  "messages": [
    "走行後の整合性チェックに失敗しました: yozoe12167: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "soya16434: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tsune.14348: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ivyke-8457: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "shiel-7488: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "solnu_13286: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "shime-15081: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "sazoe15910: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "susomia.7212: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kihone-9182: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kiri_15142: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "sotane.6083: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "sodansol-9510: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "furiivy5274: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tokuben.13726: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "toraeva-10585: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "miyuhi-6308: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hizoe_13492: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "maxzoeul.7133: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yufuko.11417: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "rosa13586: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "torki11551: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hezoehi_6771: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "mesol.12146: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "homi-12971: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hehe-16392: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "renamo.9799: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "chikedan12043: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "novin-14550: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "noameleo16083: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "moneor15062: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yunora-16972: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "danse.16591: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tesu_9699: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "seyu-12376: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "samaxan.12463: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "funoaor_9485: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yasol11944: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "sozoeyu10220: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "koriyu-6164: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "solliaul.9239: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kokunoa14707: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ruyo5753: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "mimo_13506: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "roleo-14235: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "torlia-5444: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "shimarnu_9286: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nawaeva_13165: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hiivy_13803: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "leori.5530: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "funoaleo.16095: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "solmafu-6375: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yuzoe.13653: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "maror11160: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yatorme-12998: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "miariis_12734: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yayu7410: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "leorami_14748: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ivyse16156: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kachisa_6592: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "neneri-8934: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "leosol5896: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kasolni_6181: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "mashi-9488: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "evayaso10795: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "marsetsu_14994: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "annevin.12984: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "leonois_13328: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nohael.9600: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tenoaho_15721: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "solel-14044: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ulko-16756: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "solmane.9028: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "honusa.7140: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "noremi.9972: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nenare.7629: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hifu_14332: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "evais10462: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "evasuno9565: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "muanfu6906: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kasune-13359: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "danya.11564: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nero8370: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ansuor_13007: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "none14929: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "chiro.16073: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "torlia_6243: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "fumeben.10112: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "konoamar9086: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "chisol.9416: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "bennuivy-15499: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "mareben15985: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "liamani-5116: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "rodan-5816: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "teko_15921: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tsumoor-7876: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "moleo.10603: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "moki.9005: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "danro_16949: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "toleo12851: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "liaor.6818: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "jihun_park_57: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yoleo5124: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tani_8469: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hievalia-15961: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ivyse_15334: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "elevata_5680: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tortor-10769: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nonayu-7508: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "mahilia_6259: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "teben_5928: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "isyahi_16525: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "uldan-13334: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tewato7222: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "maxshi.5235: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "meso_12441: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yotsuro-6868: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "taeva-13786: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "orivynoa8965: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "fuso.5194: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "elelzoe_11198: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nate_12976: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ormomu9428: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "samar16931: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hefusol-15055: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nadansu9902: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "solraki.14258: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tsunilia_10404: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tasoldan-6438: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "sasu_11718: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yahimax.8911: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "mute.11939: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "menoake_10714: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kivin-5668: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "bennoamia.11726: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kukeis-12922: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "wabenso.9641: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "evasu.8341: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "orma11031: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "toliaivy17009: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "mamar5434: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "neulho_14229: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "rekehe_8791: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "rehoel_10025: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kunanu-8294: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "danta-9056: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "maniha_14145: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kori-16775: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kusazoe8725: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "mumoel10448: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "bensu-6028: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ivysoivy_13819: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "teul-12678: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nona15211: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "roka_7419: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "waki.9741: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kuku6715: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "marna9714: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "liamia16470: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "rusoben-10057: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nusu15380: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ruse-14679: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "orsolia.12083: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "antona7347: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yoreyu.10500: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tenu-14907: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "sochiro16144: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ista7045: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ninuta_16978: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "sesolshi-13518: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kovinme.15514: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tachi_13760: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kemia6064: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "maxku-10272: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "benkunu7296: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ulkisol10767: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "leohe_6350: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "elne.6002: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "liaromu.14996: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hekase.13137: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "shihe-10886: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "bento.9212: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yanamia_11978: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nurake.8055: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yaischi-12904: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "iska.15144: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "orhana_6977: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "danelben-7000: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "liakesol_10843: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "danan-7696: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "runusu-12420: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "benwaso-7569: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "neleo14868: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kumi12345: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "mozoe.7186: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "evanuru-8577: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "reso.11898: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "zoene.14422: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "orisna15107: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "marse_5310: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tasolme_14176: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nuko.9227: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "maxmu7176: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "teor_5304: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "miakisu.8399: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "noanoamar15403: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "maxtorsa_5833: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nakema_11843: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nichi16814: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "mumate_6131: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kedanra_14400: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "fukaru-15440: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "wais_8784: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ninoa.15425: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "noaleo-8760: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "miakusol-5517: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "chihika.8439: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "marrese.10706: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nehanoa.9070: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "evahi.14792: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hehi.7040: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "sori_6766: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "riyo.7862: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "sekiyo.12911: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "liachi-11797: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ulto.16710: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "suvinfu9467: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "reku-7990: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ketoro_14123: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "danyu_16885: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "noasoul.12404: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hisu.11156: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kielmax-13697: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ordan11818: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tsumax-7835: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "momiatsu.7120: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "somareva.11838: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "maxfu-14563: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "naki-11223: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "molialia13450: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "namu.7935: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "leotayo.9451: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "moko.12657: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "miseke_11151: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "zoete.5756: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "issolra.16593: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nuliaha10999: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "dantare8124: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "zoemutor.9864: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "manoase_16313: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tenahe_13027: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "sorusu.5477: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "chiwa6292: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "shimelia8085: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yamaxwa_9161: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hetsu_13106: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "marisu-12794: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "resazoe.16743: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "elro_16368: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "waleo.16058: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "mumiku13603: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yoleoru13382: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kate-16342: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hiisul.14484: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "haul9843: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tosoku.13426: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kuliashi_6150: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "merian.10661: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "heru_7552: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ultsu_16454: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kamu-5766: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "chirusu-6403: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "shini_15590: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "saleoan-8923: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hekoeva.15998: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kihiso15311: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "toleote16140: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "ivyshi11400: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "muivyro9830: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "miki10990: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nekemax13744: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "evamiis-16046: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nayome.8069: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "benmiko_9611: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "koel-13347: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "danniul12842: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "leoisse-13314: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "maxto_9406: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "miakano14960: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "muis13951: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "kuelko-16258: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "fumax_13340: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "hotohe_15892: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "tsutera.5521: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "yowa_13893: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "torvintsu10292: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nuleo-7879: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "danfu.13869: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "nitorben_11712: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "minoa.10379: Memories を取得できません: GET /api/memories?month=2026-10: ステータスコード 502 が返りました（期待値 200、unknown）",
    "bennoavin-5043: 通知を購読できません: GET /api/notifications: ステータスコード 401 が返りました（期待値 200、invalid_session）",
    "saho.16482: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "noato17020: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "maxzoeri-15624: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "benni16320: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "martachi-15259: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ankehe.16035: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "isseri.16015: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ante_16186: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "benrino.16305: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "nichi16814: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "leohieva-15972: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "isfutor.15849: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "ormune-16658: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "meoryo.17002: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "seso.16463: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "honomo_16804: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "rivin16378: 通知を購読できません: GET /api/notifications: ステータスコード 503 が返りました（期待値 200、unknown）",
    "sakura_takahashi_1620: ログインできません: POST /api/auth/login: 通信に失敗しました: Post \"http://52.69.202.109:8080/api/auth/login\": EOF",
    "leozoera.5363: 友達申請を送れません: POST /api/friend-requests: ステータスコード 502 が返りました（期待値 201、unknown）"
  ]
}

--- stderr ---
```

## Access log (alp)

```
== isu1 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
57.6%	3296	2	508	GET	/api/notifications	24861.067	7.543	16.693	16.703	116
7.1%	2726	35	542	POST	/api/posts	3046.642	1.118	5.830	64.649	547
7.0%	6187	85	25	POST	/api/friend-requests	3018.281	0.488	2.157	3.396	114
4.7%	20025	411	148	GET	^/api/images/[^/]+	2013.429	0.101	1.038	2.755	124652
3.3%	7123	795	101	POST	/api/auth/login	1435.009	0.201	1.793	3.526	219
2.9%	4472	0	236	GET	/api/memories	1267.798	0.283	2.823	8.206	12788
2.9%	3139	9	71	POST	^/api/friend-requests/[^/]+/[^/]+$	1243.004	0.396	2.081	3.724	117
2.4%	16032	0	181	GET	^/api/feeds/[^/]+	1032.468	0.064	0.565	2.102	2942
1.5%	721	130	3	PUT	/api/me/profile-image	654.833	0.908	8.877	10.737	4
1.5%	666	2	31	POST	^/api/posts/[^/]+/realmojis$	639.118	0.960	3.543	4.308	459
1.4%	914	4	20	POST	^/api/posts/[^/]+/comments$	617.959	0.676	3.431	4.667	341
1.0%	1087	1	0	POST	/api/auth/register	445.551	0.410	2.066	2.157	227
0.7%	395	0	11	GET	/api/me/export	296.322	0.750	7.679	9.238	1865452
0.7%	695	0	2	PUT	^/api/follows/[^/]+$	285.783	0.411	2.124	2.804	0
0.6%	1055	0	13	POST	/api/auth/logout	272.846	0.259	2.006	3.226	2
0.6%	3023	0	12	GET	/api/users	245.314	0.081	0.499	1.041	234
0.5%	3894	2	4	GET	/api/posting-window/current	223.125	0.057	0.399	1.016	237
0.4%	531	0	1	PATCH	/api/me	188.231	0.354	1.991	3.222	276
0.4%	2136	3	27	GET	^/api/posts/me/current$	186.159	0.087	0.646	1.018	645
0.4%	2522	44	7	GET	^/api/posts/[^/]+$	152.134	0.060	0.496	1.005	612
0.3%	262	0	10	PATCH	^/api/posts/[^/]+$	150.759	0.575	3.001	3.079	632
0.3%	600	3	0	GET	/api/me	143.824	0.240	1.037	1.327	302
0.3%	1799	0	25	GET	^/api/relationships/[^/]+	113.124	0.063	0.571	2.265	785
0.2%	1007	0	10	GET	^/api/memories/[^/]+	101.716	0.101	0.829	1.463	690
0.2%	112	0	7	DELETE	^/api/posts/[^/]+$	94.237	0.841	2.225	2.256	10
0.2%	119	0	0	DELETE	^/api/realmojis/[^/]+$	82.786	0.696	3.023	3.499	0
0.2%	208	0	4	DELETE	^/api/comments/[^/]+$	78.609	0.378	2.029	2.069	3
0.2%	1217	0	10	GET	^/api/users/[^/]+$	69.188	0.057	0.457	0.675	199
0.1%	851	2	9	GET	^/api/posts/[^/]+/realmojis$	56.619	0.067	0.548	1.005	2048
0.1%	788	0	7	GET	^/api/posts/[^/]+/comments$	48.271	0.061	0.464	1.086	2994
…（上位 30 / 全 42 種類）

== isu2 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
52.5%	914	66	29	GET	/api/notifications	6123.522	6.700	11.202	11.203	93
10.9%	1857	15	25	POST	/api/friend-requests	1270.537	0.684	2.382	3.039	113
5.3%	1206	0	6	POST	/api/auth/login	623.022	0.517	2.225	2.805	246
4.6%	788	22	102	POST	/api/posts	532.532	0.676	2.444	6.501	547
4.2%	3666	23	22	GET	^/api/images/[^/]+	489.056	0.133	1.609	2.980	159604
4.2%	3418	0	72	GET	^/api/feeds/[^/]+	485.410	0.142	1.265	1.656	2943
2.3%	362	0	0	POST	/api/auth/register	272.655	0.753	2.370	2.408	228
2.1%	409	1	10	POST	^/api/friend-requests/[^/]+/[^/]+$	245.446	0.600	2.217	2.584	113
1.7%	204	0	0	PUT	/api/me/profile-image	199.222	0.977	2.726	3.302	0
1.7%	190	0	0	PATCH	/api/me	197.895	1.042	2.733	2.749	274
1.6%	239	0	0	PUT	^/api/follows/[^/]+$	182.872	0.765	2.247	2.822	0
1.3%	890	0	103	GET	/api/memories	147.053	0.165	1.632	2.023	15041
1.0%	201	0	0	POST	^/api/posts/[^/]+/comments$	112.766	0.561	2.613	2.909	333
0.9%	1185	2	6	GET	/api/posting-window/current	109.548	0.092	0.817	1.309	237
0.8%	105	4	0	POST	^/api/posts/[^/]+/realmojis$	97.987	0.933	3.175	3.493	428
0.8%	906	0	5	GET	/api/users	97.634	0.108	1.156	1.620	239
0.7%	583	38	23	GET	^/api/posts/me/current$	80.308	0.138	1.302	1.572	583
0.5%	153	0	2	POST	/api/auth/logout	59.689	0.390	1.989	2.258	0
0.5%	362	0	8	GET	^/api/relationships/[^/]+	55.057	0.152	1.267	1.590	527
0.5%	455	12	4	GET	^/api/posts/[^/]+$	54.264	0.119	1.005	1.591	609
0.4%	61	6	3	PATCH	^/api/posts/[^/]+$	44.660	0.732	2.872	2.872	559
0.2%	241	0	0	GET	^/api/users/[^/]+$	27.230	0.113	0.688	0.804	202
0.2%	31	0	2	DELETE	^/api/posts/[^/]+$	26.209	0.845	2.613	2.613	0
0.2%	208	2	0	GET	/api/me	21.122	0.102	0.265	1.280	303
0.2%	182	0	0	GET	^/api/posts/[^/]+/comments$	19.928	0.109	1.063	1.516	2972
0.2%	41	0	1	DELETE	^/api/comments/[^/]+$	19.291	0.471	2.619	2.619	0
0.2%	62	0	1	GET	/api/me/export	17.701	0.285	4.973	4.973	2238950
0.1%	171	0	0	GET	^/api/memories/[^/]+	16.521	0.097	0.520	1.178	654
0.1%	138	1	1	GET	^/api/posts/[^/]+/realmojis$	15.573	0.113	0.803	1.198	1541
0.1%	14	0	0	DELETE	^/api/realmojis/[^/]+$	11.088	0.792	1.784	1.784	0
…（上位 30 / 全 36 種類）

== isu3 ==
share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes
52.9%	828	28	31	GET	/api/notifications	5787.580	6.990	9.502	9.742	108
8.7%	1590	15	13	POST	/api/friend-requests	952.382	0.599	2.211	2.638	113
5.8%	745	20	51	POST	/api/posts	638.302	0.857	3.195	6.382	592
5.3%	1140	0	8	POST	/api/auth/login	576.653	0.506	2.221	3.826	248
4.6%	3707	28	25	GET	^/api/images/[^/]+	507.489	0.137	1.219	2.174	166605
4.2%	3358	0	51	GET	^/api/feeds/[^/]+	463.337	0.138	1.066	1.437	3144
2.9%	570	3	17	POST	^/api/friend-requests/[^/]+/[^/]+$	321.096	0.563	2.129	2.452	112
1.9%	298	0	0	POST	/api/auth/register	209.968	0.705	2.365	2.387	230
1.5%	1088	0	5	GET	/api/posting-window/current	159.107	0.146	0.895	1.184	239
1.3%	938	0	95	GET	/api/memories	143.314	0.153	1.312	3.932	15910
1.2%	162	0	0	PUT	/api/me/profile-image	127.933	0.790	2.971	3.331	0
1.1%	131	4	1	POST	^/api/posts/[^/]+/realmojis$	118.075	0.901	3.299	4.340	476
1.0%	173	0	0	PUT	^/api/follows/[^/]+$	109.045	0.630	2.497	2.634	0
1.0%	619	13	28	GET	^/api/posts/me/current$	108.111	0.175	1.221	1.406	612
0.9%	148	0	0	PATCH	/api/me	99.839	0.675	2.343	2.771	279
0.9%	722	0	8	GET	/api/users	96.761	0.134	0.840	1.361	236
0.8%	103	2	2	POST	^/api/posts/[^/]+/comments$	91.050	0.884	3.293	3.341	316
0.8%	190	0	5	POST	/api/auth/logout	88.911	0.468	2.167	2.933	0
0.6%	353	0	6	GET	^/api/relationships/[^/]+	60.598	0.172	1.194	1.385	687
0.4%	437	19	2	GET	^/api/posts/[^/]+$	48.368	0.111	0.856	1.244	598
0.4%	69	3	7	PATCH	^/api/posts/[^/]+$	45.369	0.658	2.401	2.401	573
0.3%	31	0	0	DELETE	^/api/posts/[^/]+$	30.095	0.971	2.339	2.339	0
0.2%	203	0	1	GET	^/api/posts/[^/]+/realmojis$	21.876	0.108	0.756	0.877	1511
0.2%	201	0	1	GET	^/api/users/[^/]+$	20.080	0.100	0.650	1.357	200
0.2%	154	0	1	GET	^/api/memories/[^/]+	19.776	0.128	1.116	2.000	1004
0.2%	161	0	0	GET	^/api/posts/[^/]+/comments$	17.269	0.107	0.438	1.552	2122
0.2%	77	0	1	GET	/api/me/export	16.587	0.215	1.790	1.790	3140364
0.2%	24	0	0	DELETE	^/api/comments/[^/]+$	16.504	0.688	1.491	1.491	0
0.2%	163	0	0	GET	/api/me	16.421	0.101	0.366	0.366	307
0.1%	24	0	0	DELETE	^/api/realmojis/[^/]+$	15.058	0.627	2.149	2.149	0
…（上位 30 / 全 35 種類）
```

（生の出力は alp.txt）

## Slow query (pt-query-digest)

```
全 16906 クエリ / 48 種類 / 合計 152.402s（上位 15 種類。share は全体の合計時間に対する割合。blks_avg は 1 回あたりに読んだ共有バッファのページ数で、rows_avg よりずっと大きければインデックスが効いていない）

rank	share	calls	total_s	avg_ms	max_ms	blks_avg	rows_avg	query
1	55.1%	399	83.945	210.39	2528.31	1230.3	126.2	SELECT i.id, i.data FROM unnest($1::text[], $2::timestamptz[]) AS k(id, created_at) CROSS JOIN LATERAL (SELECT id, data FROM image_objects WHERE id = k.id AND c…
2	14.0%	3	21.277	7092.34	9539.76	31603.0	1199348.0	SELECT id, post_id, reactor_id, image_id, reaction_type, created_at, replaced_at, deleted_at FROM realmojis
3	11.7%	14121	17.845	1.26	11.29	9.4	1.0	SELECT data FROM image_objects WHERE id = $1 AND created_at = $2
4	10.4%	3	15.925	5308.31	7156.48	24949.0	625111.0	SELECT id, author_id, posting_window_id, image_id, caption, visibility, location_kind, latitude, longitude, locality, posted_at, created_at, deleted_at FROM pos…
5	5.3%	3	8.060	2686.63	3798.27	13251.0	529632.0	SELECT id, post_id, author_id, body, created_at FROM comments WHERE deleted_at IS NULL
6	2.1%	4	3.148	786.96	2622.03	204500.0	2491.0	SELECT u.id, i.created_at FROM users u JOIN image_objects i ON i.id = u.profile_image_id
7	0.2%	4	0.355	88.77	156.63	592.0	28970.0	SELECT id, requester_id, recipient_id, status, created_at, updated_at FROM friend_requests
8	0.2%	4	0.348	87.06	156.63	610.0	24985.0	SELECT id, user_low_id, user_high_id, source_request_id, created_at, updated_at FROM friendships
9	0.2%	1	0.271	270.85	270.85	237331.0	28970.0	INSERT INTO friend_requests (id, requester_id, recipient_id, status, created_at, updated_at) SELECT id, requester_id, recipient_id, status, created_at, updated_…
10	0.2%	1	0.248	247.62	247.62	33183.0	0.0	DELETE FROM notifications WHERE created_at > (SELECT created_at FROM initial_data_boundary)
11	0.2%	1	0.240	240.26	240.26	201964.0	24985.0	INSERT INTO friendships (id, user_low_id, user_high_id, source_request_id, created_at, updated_at) SELECT id, user_low_id, user_high_id, source_request_id, crea…
12	0.1%	1	0.171	170.51	170.51	31604.0	0.0	DELETE FROM realmojis WHERE created_at > (SELECT created_at FROM initial_data_boundary)
13	0.1%	1	0.103	103.03	103.03	24950.0	0.0	DELETE FROM posts WHERE created_at > (SELECT created_at FROM initial_data_boundary)
14	0.1%	1	0.083	82.79	82.79	2412.0	0.0	DO $$ DECLARE partition_name text; BEGIN FOR partition_name IN SELECT c.relname FROM pg_inherits i JOIN pg_class c ON c.oid = i.inhrelid WHERE i.inhparent = 'im…
15	0.0%	1	0.069	68.68	68.68	1503.0	0.0	TRUNCATE ip_bans, login_sessions, administrator_sessions, friendships, friend_requests, follows, block_relations

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
Build ID: c81bcb19f8760b473535e00f5ef01d1630c0c3ad
Type: cpu
Time: 2026-10-31 15:45:03 JST
Duration: 60.05s, Total samples = 43.40s (72.28%)
Showing nodes accounting for 31.07s, 71.59% of 43.40s total
Dropped 1126 nodes (cum <= 0.22s)
Showing top 40 nodes out of 296
      flat  flat%   sum%        cum   cum%
     5.90s 13.59% 13.59%      5.90s 13.59%  crypto/internal/fips140/sha256.blockAVX2
     4.83s 11.13% 24.72%      4.83s 11.13%  internal/runtime/syscall/linux.Syscall6
     2.11s  4.86% 29.59%      2.11s  4.86%  runtime.memmove
     1.76s  4.06% 33.64%      5.67s 13.06%  example.com/isureal-company/isureal.(*Store).rebuildIndexes
     1.60s  3.69% 37.33%      6.28s 14.47%  runtime.scanObject
     1.58s  3.64% 40.97%      2.09s  4.82%  runtime.findObject
     1.54s  3.55% 44.52%      1.54s  3.55%  runtime.memclrNoHeapPointers
     1.47s  3.39% 47.90%      2.42s  5.58%  runtime.tryDeferToSpanScan
     1.46s  3.36% 51.27%      1.46s  3.36%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     1.28s  2.95% 54.22%      1.28s  2.95%  aeshashbody
     0.68s  1.57% 55.78%      0.68s  1.57%  runtime.spanClass.sizeclass (inline)
     0.57s  1.31% 57.10%      2.21s  5.09%  runtime.mapaccess1_faststr
     0.49s  1.13% 58.23%      0.49s  1.13%  memeqbody
     0.43s  0.99% 59.22%      0.43s  0.99%  runtime.nextFreeFast (inline)
     0.42s  0.97% 60.18%      0.63s  1.45%  example.com/isureal-company/isureal.postLess
     0.34s  0.78% 60.97%      0.34s  0.78%  runtime.futex
     0.29s  0.67% 61.64%      6.96s 16.04%  crypto/internal/fips140/sha256.(*Digest).Sum
     0.29s  0.67% 62.30%      6.55s 15.09%  crypto/internal/fips140/sha256.(*Digest).Write
     0.28s  0.65% 62.95%      0.58s  1.34%  runtime.greyobject
     0.28s  0.65% 63.59%      1.26s  2.90%  runtime.mallocgcSmallScanNoHeader
     0.24s  0.55% 64.15%      0.55s  1.27%  database/sql.convertAssignRows
     0.24s  0.55% 64.70%      0.28s  0.65%  encoding/json.appendString[go.shape.string]
     0.24s  0.55% 65.25%      3.91s  9.01%  runtime.mallocgc
     0.24s  0.55% 65.81%      1.87s  4.31%  runtime.mapassign_faststr
     0.22s  0.51% 66.31%      0.42s  0.97%  compress/flate.(*compressor).deflate
     0.22s  0.51% 66.82%      0.22s  0.51%  runtime.typePointers.nextFast (inline)
     0.21s  0.48% 67.30%      0.56s  1.29%  runtime.scanObjectsSmall
     0.21s  0.48% 67.79%      0.24s  0.55%  runtime.typePointers.next
     0.20s  0.46% 68.25%      7.38s 17.00%  crypto/internal/fips140/hmac.(*HMAC).Sum
     0.20s  0.46% 68.71%      6.57s 15.14%  crypto/internal/fips140/sha256.(*Digest).checkSum
     0.18s  0.41% 69.12%      2.94s  6.77%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.17s  0.39% 69.52%      0.22s  0.51%  runtime.(*mspan).writeHeapBitsSmall
     0.17s  0.39% 69.91%      0.22s  0.51%  runtime.spanOf (inline)
     0.12s  0.28% 70.18%      0.30s  0.69%  crypto/internal/fips140/hmac.(*HMAC).Reset
     0.12s  0.28% 70.46%      0.22s  0.51%  internal/runtime/maps.(*Iter).Next
     0.11s  0.25% 70.71%      0.94s  2.17%  runtime.mallocgcSmallNoscan
     0.10s  0.23% 70.94%      0.74s  1.71%  encoding/json.structEncoder.encode
     0.10s  0.23% 71.18%      7.84s 18.06%  example.com/isureal-company/isureal.computeImageSignature
     0.09s  0.21% 71.38%      0.27s  0.62%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.09s  0.21% 71.59%      5.99s 13.80%  crypto/internal/fips140/sha256.block
```

## Resources

| host | CPU 使用率 平均 / 最大 | 90% 以上の時間 | IO 待ち 平均 | run queue 最大 | free 最小 | |
|---|---|---|---|---|---|---|
| isu1 (2 cpu) | 58.1% / 100% | 39% | 0.7% | 6 | 313 MB | CPU が飽和 |
| isu2 (2 cpu) | 36.2% / 100% | 7% | 0.7% | 6 | 229 MB | run queue が積もっている |
| isu3 (2 cpu) | 35.0% / 100% | 8% | 0.6% | 4 | 291 MB |  |
| isu4 (2 cpu) | 59.2% / 100% | 28% | 19.5% | 41 | 101 MB | run queue が積もっている |

**isu1（CPU が飽和、主に isureal）、isu2（run queue が積もっている、主に isureal）、isu4（run queue が積もっている、主に isureal）**。ここから手を付ける。

**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）

| host | 上位プロセス |
|---|---|
| isu1 | isureal 92% / nginx 19% |
| isu2 | isureal 68% / nginx 2% |
| isu3 | isureal 67% / nginx 3% |
| isu4 | isureal 54% / postgres 21% |

## Diff from previous run

スコア: 0 → — (fail)

`106c0a8..dfccbc3`

```
 webapp/go/offload.go | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)
```
