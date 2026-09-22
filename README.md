# isucon-harness

ISUCON2026 用の作業ハーネス。**計測 → 改善 → 検証**のループを 1 コマンドで回し、
複数の Claude Code セッションを並列で走らせても計測が壊れないようにするための道具箱。

- 競技日: **2026-10-31(土) 10:00 - 18:00 JST**
- レギュレーション: https://isucon.net/archives/59966826.html
- ポータル: https://portal.isucon.net/

## これは何をするものか

競技中にやることは、突き詰めると次の 4 つしかない。

1. 変更をデプロイする
2. ベンチを回す
3. 計測結果を読む
4. 良くなっていなければ戻す

このリポジトリは、この 4 つを `bin/bench` 一発に畳み、**同時に走っている他の作業が
計測を汚染しないようにロックする**ことだけを目的にしている。

## クイックスタート

```sh
cp config.env.example config.env
$EDITOR config.env          # 当日埋める値はすべてここに集約されている
bin/setup                   # 起動後のサーバセットアップ
bin/bench                   # deploy → check → ベンチ → レポート生成
bin/history                 # スコア台帳
```

## ディレクトリ

| パス | 中身 |
|---|---|
| `docs/DESIGN.md` | **全レーン共有の契約**。実装前に必ず読む |
| `config.env.example` | 当日埋める値の一覧 |
| `bin/` | 作業コマンド一式 |
| `runs/` | ベンチ 1 回 = 1 ディレクトリ。`history.tsv` がスコア台帳 |
| `CLAUDE.md` | レギュレーション上の禁止事項とリポジトリの規約 |

## コマンド

| コマンド | 役割 | 実装 Issue |
|---|---|---|
| `bin/setup` | 起動後のサーバセットアップ | #9 |
| `bin/deploy` | 指定 ref を全台へ配布して再起動 | #1 |
| `bin/check` | ベンチ不要の検証ゲート | #12 |
| `bin/bench` | ロック取得〜ベンチ〜レポートまで通し | #2 #4 |
| `bin/history` | スコア台帳の表示 | #5 |
| `bin/rollback` | 直前の good revision へ戻す | #6 |
| `bin/portal` | ポータルを Cookie 付き curl で操作 | #3 |
| `bin/monitor` | 全台のリソースを 1 画面で見る | #17 |
| `bin/verify-restart` | 再起動試験（失格回避） | #7 |
| `bin/finalize` | 終盤チェックリスト | #8 |

## 現在の状態

契約とスタブのみ（#21）。各コマンドの中身は担当 Issue で実装する。
スタブは引数解釈と設定読み込みまでは通り、未実装部分は担当 Issue 番号を表示して終了する。
