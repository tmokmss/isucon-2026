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
$EDITOR config.ts           # 当日埋める値はすべてここに集約されている（型が効く）。コミットする

bin/isuh setup              # 起動後のサーバセットアップ（先に config.ts の setup.addresses を埋める）
bin/isuh bench              # deploy → check → ベンチ → レポート生成
bin/isuh history            # スコア台帳
```

ハーネスは**手元の PC で動かす**。競技サーバには置かない。

## ポータルの Cookie を用意する

ポータル（https://portal.isucon.net/ ）に公開 API は無く、ブラウザ操作が前提。
GitHub 認証後の Django セッション Cookie を手元に持ち出して使う。

1. ブラウザで https://portal.isucon.net/ にログインする
2. DevTools → Application → Cookies → `https://portal.isucon.net`
3. `sessionid` と `csrftoken` の値をコピーして、次の **1 行**をファイルに保存する

   ```
   sessionid=<値>; csrftoken=<値>
   ```

4. 置き場所を `config.ts` の `portal.cookieFile` に書く（既定 `~/.isucon/portal-cookie`）

DevTools の Network タブでリクエストを右クリック → Copy → Copy as cURL して、
`-H 'Cookie: ...'` の中身を貼っても同じ。`#` 始まりの行はコメントとして無視される。

```sh
bin/isuh portal whoami        # Cookie が生きているかの確認
bin/isuh portal get /         # 任意のページを Cookie 付きで引く
```

Cookie が切れていれば終了コード 65 で、取り直す手順ごと表示される。

> **ベンチ起動のエンドポイントは当日まで分からない。**
> `isuh portal get` で探って、`config.ts` の `portal.enqueuePath` /
> `jobPath` / `scorePath` を埋める。

## コマンド

```
isuh <command> [args]
```

| コマンド | 役割 | 実装 Issue |
|---|---|---|
| `setup` | 起動後のサーバセットアップ | #9 #11 |
| `deploy` | 指定 ref を全台へ配布して再起動 | #1 ✅ |
| `check` | ベンチ不要の検証ゲート | #12 |
| `bench` | ロック取得〜ベンチ〜レポートまで通し | #2 ✅ #4 ✅ |
| `history` | スコア台帳の表示 | #5 ✅ |
| `portal` | ポータルを Cookie 付きの HTTP で操作 | #3 ✅ |
| `monitor` | 全台のリソースを 1 画面で見る | #17 |
| `verify-restart` | 再起動試験（失格回避） | #7 |
| `finalize` | 終盤チェックリスト | #8 |

詳細は `bin/isuh <command> --help`。

## ディレクトリ

**このリポジトリは ISUCON2026 専用。** アプリ（`webapp/`）・ハーネス・構成管理（`conf/`）を
1 つにまとめ、コードと設定を同じ commit で変えて同じ sha でデプロイする。

| パス | 中身 |
|---|---|
| `docs/DESIGN.md` | **全レーン共有の契約**。実装前に必ず読む |
| `config.ts` | 当日埋める値のすべて。**コミットする**（worktree からも isuh が動くように）。秘密は書かない |
| `src/cli.ts` | エントリポイント。サブコマンドの振り分け |
| `src/commands/` | 各コマンド |
| `src/lib/` | 共通ライブラリ（ログ・終了コード・ssh・git・ロック） |
| `webapp/` | 競技のアプリ（サーバの `~/webapp`）。`isuh setup` が取り込み、`isuh deploy` がビルドしたバイナリと、バイナリ以外のファイル（`sql/` 等）を同じ sha で配る（上書きのみ・削除はしない） |
| `conf/` | サーバの設定ファイル。`isuh setup` が取り込み、`isuh deploy` がアプリと同じ sha のものを絶対パスへ写す |
| `setup/` | setup が配るもの（dotfiles・ログ設定）と、チューニングの参考テンプレ |
| `test/` | sshd コンテナ 3 台の検証環境。実サーバ無しで動作確認する |
| `runs/` | ベンチ 1 回 = 1 ディレクトリ。`history.tsv` がスコア台帳 |
| `CLAUDE.md` | レギュレーション上の禁止事項とリポジトリの規約 |

## 開発

```sh
npm run typecheck      # tsc --noEmit

./test/up.sh           # 検証用の sshd コンテナ 3 台（isu1/isu2/isu3）を起動
./test/all.sh          # typecheck + 全スイート（deploy / ロックと bench / ポータル）
./test/down.sh
```

ビルドは不要。`tsx` が `.ts` を直接実行する。

当日まで本物のサーバが無いので、**実装したら `test/` で必ず一度動かす**。
詳細は `test/README.md`。

## 現在の状態

**中核ループ（deploy → bench → report → history）が通しで動く。**
`isuh bench` 一発で、ロック取得・デプロイ・ログのリセット・リソース計測・
ベンチ実行・計測の回収・`report.md` の生成・台帳への追記・スコア判定まで走る。

残っているのは `check`(#12)、
`verify-restart`(#7)、`finalize`(#8)、`monitor`(#17)。
ベンチ起動のエンドポイントは当日 `config.ts` に埋める。

他のコマンドはスタブで、引数解釈と設定読み込みまでは通り、
未実装部分は担当 Issue 番号を表示して終了コード 70 で終わる。
