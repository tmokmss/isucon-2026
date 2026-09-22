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

このリポジトリは、この 4 つを `isuh bench` 一発に畳み、**同時に走っている他の作業が
計測を汚染しないようにロックする**ことだけを目的にしている。

## クイックスタート

```sh
npm install
cp config.example.ts config.ts
$EDITOR config.ts           # 当日埋める値はすべてここに集約されている（型が効く）

bin/isuh setup              # 起動後のサーバセットアップ
bin/isuh bench              # deploy → check → ベンチ → レポート生成
bin/isuh history            # スコア台帳
```

ハーネスは**手元の PC で動かす**。競技サーバには置かない。

## コマンド

```
isuh <command> [args]
```

| コマンド | 役割 | 実装 Issue |
|---|---|---|
| `setup` | 起動後のサーバセットアップ | #9 #11 |
| `deploy` | 指定 ref を全台へ配布して再起動 | #1 |
| `check` | ベンチ不要の検証ゲート | #12 |
| `bench` | ロック取得〜ベンチ〜レポートまで通し | #2 #4 |
| `history` | スコア台帳の表示 | #5 |
| `rollback` | 直前の good revision へ戻す | #6 |
| `portal` | ポータルを Cookie 付きの HTTP で操作 | #3 |
| `monitor` | 全台のリソースを 1 画面で見る | #17 |
| `verify-restart` | 再起動試験（失格回避） | #7 |
| `finalize` | 終盤チェックリスト | #8 |

詳細は `bin/isuh <command> --help`。

## ディレクトリ

| パス | 中身 |
|---|---|
| `docs/DESIGN.md` | **全レーン共有の契約**。実装前に必ず読む |
| `config.example.ts` | 当日埋める値の一覧。`config.ts` にコピーして使う |
| `src/cli.ts` | エントリポイント。サブコマンドの振り分け |
| `src/commands/` | 各コマンド |
| `src/lib/` | 共通ライブラリ（ログ・終了コード・ssh・git・ロック） |
| `runs/` | ベンチ 1 回 = 1 ディレクトリ。`history.tsv` がスコア台帳 |
| `CLAUDE.md` | レギュレーション上の禁止事項とリポジトリの規約 |

## 開発

```sh
npm run typecheck      # tsc --noEmit
```

ビルドは不要。`tsx` が `.ts` を直接実行する。

## 現在の状態

契約とスタブのみ（#21）。各コマンドの中身は担当 Issue で実装する。
スタブは引数解釈と設定読み込みまでは通り、未実装部分は担当 Issue 番号を表示して
終了コード 70 で終わる。
