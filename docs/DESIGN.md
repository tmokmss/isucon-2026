# 設計契約

このドキュメントは **全レーン共有の契約** である。
並列で実装するため、ここに書かれたインターフェースを独断で変更すると他レーンが壊れる。
**変更が必要なら実装せずに先に相談すること。**（→ `CLAUDE.md`）

契約の範囲内での実装方法は自由。

---

## 0. 実装スタック

**TypeScript / Node + tsx。** ビルド不要で `.ts` を直接実行する。

| | |
|---|---|
| 実行 | `node_modules/.bin/tsx`（`bin/isuh` が呼ぶ） |
| サブプロセス | `execa`（`ssh` / `scp` / `git` はすべてこれ経由） |
| 型チェック | `npm run typecheck`（`tsc --noEmit`） |
| 実行場所 | **手元の PC**。競技サーバには置かない |

ハーネスは手元で動かすので、単一バイナリ化やクロスコンパイルは不要。
シェルスクリプトは `bin/isuh`（tsx を呼ぶだけの 10 行）のみに留める。

理由: この道具の仕事は「ssh でコマンドを叩く」より「**ポータルへの HTTP
（Cookie + CSRF）／ JSON の構造比較 ／ レポートの整形 ／ スコアの小数比較**」の
比重が大きく、そこは shell では壊れやすい。当日の障害時に読んで直せることを優先する。

---

## 1. 全体像

競技中にやることは 4 つしかない。

1. 変更をデプロイする
2. ベンチを回す
3. 計測結果を読む
4. 良くなっていなければ戻す

`isuh bench` がこの 4 つを 1 コマンドに畳む。**並列に走る他の作業が計測を汚染しない**よう、
デプロイからレポート生成までを 1 つのクリティカルセクションに入れる。

```
isuh bench [<git-ref>]
 ├─ 1. ロック取得 ........................ 失敗 → exit 10
 ├─ 2. git ref を sha に解決
 ├─ 3. deploy <sha> ...................... 失敗 → exit 20
 ├─ 4. check ............................. 失敗 → exit 30（ベンチを起動しない）
 ├─ 5. ログのリセット + リソース計測開始
 ├─ 6. portal でベンチ登録 → 完了待ち .... 失敗 → exit 40
 ├─ 7. 計測停止・ログ収集（alp / pt-query-digest / pprof / リソース）
 ├─ 8. runs/<run-id>/ に report.md と meta.json を生成、history.tsv に追記
 ├─ 9. スコア判定 ........................ 低下 → exit 50（autoRollback なら戻す）
 └─ 10. ロック解放（try-finally で異常終了時も必ず実行）
```

**`check` は deploy の後に走る。** 新しいコードが載った状態を検証するため。
落ちた場合はベンチを起動せず、貴重なベンチ枠を消費しない。

---

## 2. 設定の契約 — `config.ts`

- 当日埋める値は **すべて** `config.ts` に置く。ソースに直書きしない
- 型は `src/config.ts` の `Config` インターフェースが唯一の定義。
  **項目を増やすときは型と `config.example.ts` を同時に更新する**
- 読み込みは `loadConfig()` が行う。各コマンドは自前で読まない
- 必須項目が空なら `EXIT.CONFIG`(65) で落ちる
- `ISUH_CONFIG` 環境変数で場所を上書きできる（テスト用）

**typed config にしている理由**: 当日 `config.ts` を編集するとき、補完と型チェックが効く。
`hosts.all` が配列なので、空白区切り文字列を split する必要もない。

ホストの役割は `config.hosts` で表す。値は `~/.ssh/config` 上のホスト別名。

| フィールド | 意味 |
|---|---|
| `hosts.all` | 全ホスト。再起動試験の対象 |
| `hosts.app` | アプリを動かすホスト。deploy の配布先 |
| `hosts.nginx` | nginx のホスト。access log の収集元 |
| `hosts.db` | DB のホスト。slow query log の収集元 |
| `hosts.primary` | ベンチのリクエスト先。**ロックの保持先** |

アプリのコードは**このハーネスとは別のリポジトリ**にある。
`isuh deploy <git-ref>` や `isuh bench <git-ref>` の ref は `config.app.repoLocal`
（手元のアプリリポジトリのチェックアウト）で解決する。ハーネス自身のリポジトリではない。
これによりアプリ側を worktree で並列に触る運用がそのまま成立する。

---

## 3. CLI の契約 — `bin/isuh`

エントリポイントはひとつ。`bin/isuh <command> [args]`。

全コマンド共通:

- `-h` / `--help` で usage を stdout に出して `exit 0`
- 引数不正は `EXIT.USAGE`(64)、設定不備は `EXIT.CONFIG`(65)
- 引数の検証は設定の読み込みより先に行う
- 各コマンドは `src/commands/<name>.ts` で `Command` を `command` という名前で export する

| コマンド | 引数 | 役割 | 実装 |
|---|---|---|---|
| `setup` | `[--only <step>]` | 鍵配布・/etc/hosts・dotfiles・計測ツール導入 | #9 #11 |
| `deploy` | `[<git-ref>] [--dry-run]` | 指定 ref をビルドして全台へ配布・再起動 | #1 |
| `check` | なし | ベンチ不要の検証ゲート | #12 |
| `bench` | `[<git-ref>] [--skip-check] [--note <text>]` | 上記の通し実行 | #2 #4 |
| `history` | `[-n <count>]` | スコア台帳の表示 | #5 |
| `rollback` | `[<sha>]` (既定: 直前の good) | 指定 revision へ戻してデプロイ | #6 |
| `portal` | `<subcommand>` | ポータル操作（下記） | #3 |
| `monitor` | なし | 全台のリソースを tmux で 1 画面に | #17 |
| `verify-restart` | `[--times <n>]` | ランダム順で再起動 → initialize 無しでベンチ | #7 |
| `finalize` | なし | 終盤チェックリスト | #8 |

### `isuh portal` のサブコマンド

| サブコマンド | 役割 |
|---|---|
| `whoami` | Cookie が有効かの確認。認証必須ページを引けるか |
| `enqueue` | ベンチを登録し、ジョブ ID を **stdout に出す** |
| `status <job-id>` | ジョブ状態を出す（`pending` / `running` / `done` / `failed`） |
| `wait <job-id>` | 完了まで待ち、結果 JSON を stdout に出す |
| `score` | 最新の確定スコアを stdout に出す |

> **ベンチ用のエンドポイントは当日にならないと分からない。**
> URL とパラメータは `config.portal` に集約し、当日は埋めるだけにする。
> 事前に検証できるのは `whoami` 相当（Cookie 付きで認証必須ページを引く）まで。

---

## 4. 成果物の契約 — `runs/`

### run-id

```
<YYYYmmdd>-<HHMMSS>-<sha7>      例: 20261031-103512-a1b2c3d
```

辞書順 = 時系列順になること。生成は `newRunId(sha)`。

### ディレクトリ構造

```
runs/
├── history.tsv                       # スコア台帳（git で追跡する）
└── 20261031-103512-a1b2c3d/
    ├── report.md                     # 人間と AI が読む唯一のファイル
    ├── meta.json                     # 機械が読む
    ├── bench.log                     # ベンチの生出力
    ├── alp.txt
    ├── ptquery.txt
    ├── pprof.txt
    └── resource/
        ├── isu1.txt
        ├── isu2.txt
        └── isu3.txt
```

`runs/<run-id>/` は `.gitignore` 済み（`history.tsv` のみ追跡）。

### `meta.json`

```json
{
  "runId": "20261031-103512-a1b2c3d",
  "startedAt": "2026-10-31T10:35:12+09:00",
  "finishedAt": "2026-10-31T10:37:02+09:00",
  "ref": "HEAD",
  "sha": "a1b2c3d4e5f6...",
  "score": 12345,
  "status": "pass",
  "note": "add index on competition_id",
  "benchJobId": "4321",
  "prevRunId": "20261031-102011-9f8e7d6",
  "bestRunId": "20261031-103512-a1b2c3d"
}
```

`status` は `pass` / `fail` / `error` のいずれか。

- `pass` — ベンチが完走しスコアが付いた
- `fail` — ベンチが完走したが整合性チェック等で失敗した
- `error` — ベンチを起動できなかった、デプロイに失敗した等

この形は `src/lib/run.ts` に `RunMeta` として型で置く（#4 で追加）。

### `report.md` の見出し

**見出しは固定。** #4 の生成側と #14 の計測スキルがこの形で揃う。

```markdown
# run <run-id>

## Summary
<!-- score / status / ref / sha / 所要時間 / 前回比 / ベスト比 / note -->

## Bench
<!-- ベンチの出力。status が pass 以外ならエラーを全文載せる -->

## Access log (alp)

## Slow query (pt-query-digest)

## CPU profile (pprof)
<!-- metrics.pprofEnabled が false なら「未取得」と書く -->

## Resources
<!-- ベンチ実行中の各ホストの CPU / mem / IO / load。どのホストが飽和したか -->

## Diff from previous run
<!-- 直前 run からの git diff --stat と、スコアの増減 -->
```

**`report.md` 単体を読めば次に何を直すべきか判断できること**が唯一の品質基準。
これがそのまま Claude Code に渡される。

### `history.tsv`

タブ区切り。1 行目はヘッダ。

```
run_id	started_at	sha	ref	score	status	note
```

- 追記のみ。既存行を書き換えない
- `score` は `status !== "pass"` のとき `0`
- `note` にタブ・改行を含めない

---

## 5. 終了コードの契約

呼び出し側が失敗理由で分岐するため、**数値を直書きせず `EXIT` を使う**。

| 定数 | 値 | 意味 |
|---|---|---|
| `EXIT.OK` | `0` | 成功 |
| `EXIT.LOCK` | `10` | ロック取得失敗（タイムアウト） |
| `EXIT.DEPLOY` | `20` | デプロイ失敗 |
| `EXIT.CHECK` | `30` | 検証ゲート失敗 |
| `EXIT.BENCH` | `40` | ベンチ失敗 |
| `EXIT.REGRESSION` | `50` | スコア低下（`bench.scoreRegressionRatio` を下回った） |
| `EXIT.USAGE` | `64` | 引数不正 |
| `EXIT.CONFIG` | `65` | 設定不備 |
| `EXIT.NOT_IMPLEMENTED` | `70` | 未実装 |

失敗は `throw new HarnessError(EXIT.X, message)` で表す。`src/cli.ts` が捕まえて
`log.error` した上で `process.exit(code)` する。**各コマンドで `process.exit` を呼ばない。**

`usageError()` / `configError()` / `notImplemented(issue)` が用意してある。

---

## 6. ログ出力の契約

- **stdout は機械可読な結果のみ。** `isuh portal enqueue` ならジョブ ID だけ、
  `isuh history` なら表だけ。出力には `out()` を使う
- 進捗・警告・エラーは **すべて stderr**。`log.info` / `log.warn` / `log.error` / `log.debug` を使う
- 形式: `[HH:MM:SS] LEVEL message`
- `VERBOSE=1` で `log.debug` が有効になる

これを守ると `JOB=$(isuh portal enqueue)` のようなパイプ利用が安全にできる。

---

## 7. ベンチのロックプロトコル

**デプロイだけをロックしても不十分。** ベンチ実行中に別のデプロイが走れば計測は壊れる。
`isuh bench` の 1〜10 全体が 1 つのクリティカルセクション。

### 保持方法

`hosts.primary` 上の `flock` を、**ssh の子プロセスを生かしたまま保持する**。

```
ssh <primary> "exec 9><lockPath>; flock -w <timeout> 9 || exit 10; echo ACQUIRED; cat"
```

- 末尾の `cat` が stdin を開いたまま保つ → **ssh プロセスが生きている間だけロックが保持される**
- 子プロセスを kill すると接続が切れ、flock が自動解放される
- Node なら `try-finally` と `process.on("exit")` で確実に kill できるため、
  **Ctrl-C でも異常終了でもロックが残らない**（#2 の完了条件）
- ローカルのファイルロックではなく**サーバ側**に置く理由は、worktree・別ターミナル・
  別マシンをまたいで効かせるため

### 保持者の表示

取得できなかった場合、**誰がいつから保持しているか**を stderr に出して `EXIT.LOCK`。
`<lockPath>.owner` に `ホスト名 / ユーザ / 開始時刻 / git ref` を書いておく。

### 迂回の禁止

ロックを取らずに `scp` や `systemctl restart` をすると、走っているベンチの計測が壊れる。
**デプロイは必ず `isuh deploy` / `isuh bench` 経由。**

---

## 8. `src/lib/` が提供するもの

ここにある関数は各コマンドで再実装しない。

| モジュール | 提供するもの |
|---|---|
| `lib/exit.ts` | `EXIT` / `HarnessError` / `usageError` / `configError` / `notImplemented` |
| `lib/log.ts` | `log.debug/info/warn/error`（stderr）、`out`（stdout） |
| `lib/paths.ts` | `repoRoot` / `newRunId` / `runDir` / `historyFile` |
| `lib/ssh.ts` | `onHost` / `onHosts`（並列）/ `local` |
| `lib/git.ts` | `resolveSha`（`config.app.repoLocal` で解決） |
| `lib/lock.ts` | `acquireLock` / `Lock.release`（実装 #2） |
| `lib/command.ts` | `Command` インターフェース |
| `config.ts` | `Config` 型一式 / `loadConfig` |

`ssh` や `git` は必ず `execa` の配列引数で呼ぶ（シェルを経由しないので引用符の事故が起きない）。

---

## 9. 実装レーンの割り当て

触るファイルが重なるものは**同一レーンで直列に**扱う。

| レーン | Issue | 主に触るファイル |
|---|---|---|
| A | #1 #2 | `src/commands/deploy.ts` `src/commands/bench.ts` `src/lib/lock.ts` |
| B | #4 #5 | `src/lib/run.ts`（新規）`src/lib/report.ts`（新規）`src/commands/history.ts` |
| C | #15 #19 | `docs/` のみ |
| D | #9 #10 #11 | `src/commands/setup.ts`、`conf/` |
| E | #3 | `src/commands/portal.ts`（要: 手元のブラウザ Cookie） |

第 2 波（A・B のマージ後）: #6 #12 → #7 #8 → #18。

---

## 10. 未確定事項

当日にならないと確定しないもの。**確定したら `config.ts` を埋めるだけで済む**ように作る。

- ベンチ起動のエンドポイントとパラメータ（`portal.enqueuePath` / `portal.jobPath`）
- サーバの台数と役割（`hosts.*`）
- アプリのサービス名・バイナリ名・ビルドコマンド（`app.*`）
- alp の URI 正規化パターン（`metrics.alpMatches`）
- ログの出力先（`middleware.nginxAccessLog` / `middleware.dbSlowLog`）
