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
| 実行 | `node --import tsx`（`bin/isuh` が exec する。**プロセスは 1 本**→ §7） |
| サブプロセス | `execa`（`ssh` / `scp` / `git` はすべてこれ経由） |
| 型チェック | `npm run typecheck`（`tsc --noEmit`） |
| 実行場所 | **手元の PC**。競技サーバには置かない |

ハーネスは手元で動かすので、単一バイナリ化やクロスコンパイルは不要。
本体のシェルスクリプトは `bin/isuh`（node を exec するだけ）のみに留める
（`test/` の検証スクリプトは別。コンテナや ssh を回すだけなので sh で書く）。

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
| `whoami` | Cookie が有効かの確認 |
| `get <path>` | 任意のページを Cookie 付きで引いて stdout に出す |
| `post <path> [k=v ...]` | 任意のパスへ CSRF 付きで POST する |
| `enqueue` | ベンチを登録し、ジョブ ID を **stdout に出す** |
| `status <job-id>` | ジョブ状態を出す（`pending` / `running` / `done` / `failed`） |
| `wait <job-id>` | 完了まで待ち、結果 JSON を stdout に出す |
| `score` | 最新の確定スコアを stdout に出す |

`get` / `post` は**当日エンドポイントを探るため**にある。ブラウザの DevTools で
フォームの action や fetch 先を見つけたら、まずここで叩いて形を確かめ、
`config.portal` に書く。

**Cookie** — `portal.cookieFile` に Cookie ヘッダの中身を 1 行で置く。
`#` 始まりの行は無視、複数行は `; ` で連結する。

```
sessionid=<値>; csrftoken=<値>
```

POST は Django の CSRF を通すため、Cookie の `csrftoken` を `X-CSRFToken` ヘッダに
載せ、`Referer` と `Origin` も付けて送る。

**ログイン切れの判定** — 3xx で `/auth/login` へ飛ばされる、401/403、のほかに
**HTML 本文のナビを見る**。ポータルのトップは未ログインでも 200 を返すため
（実物で確認済み）、ステータスだけでは判定できない。未ログインなら
`/auth/login/` へのリンク、ログイン済みならログアウトのリンクが出る。

> **ベンチ用のエンドポイントは当日にならないと分からない。**
> URL とパラメータは `config.portal` に集約し、当日は埋めるだけにする。
> 事前に検証できるのは `whoami` と `get` まで。

### `isuh rollback` の戻し先

**good revision = ベスト × `bench.scoreRegressionRatio` 以上のスコアを出した最も新しいラン。**

「直前に pass したラン」ではない。**スコアが低下したランも pass ではある**ので、
それを戻し先にすると悪い版へ戻してしまう。判定の基準を bench のスコア低下判定と
揃えることで、「低下と判定されない範囲で最も新しい版」に戻ることになる。

ベストそのものへ戻さないのは、ベストが古すぎてその後の良い変更まで
巻き戻してしまうのを避けるため。ベストへ戻したいときは `isuh history` で
sha を見て `isuh rollback <sha>` と明示する。

`rollback` は**ベンチと同じロックを取る**。走っている計測を壊さないため。

`bench` は次のときに自動で戻す（`bench.autoRollback` が true のときだけ）。

- ベンチが fail した
- スコアがベスト × `scoreRegressionRatio` を下回った

序盤は false（止めるが戻さない）、終盤は true が使いやすい。
戻したことは stderr のログ、`meta.json` の `rolledBackTo`、
`report.md` の Summary に残る。

> ISUCON のスコアは数 % 〜 10 % 程度は平気でブレる。**ベンチ 1 回で改悪と断定しない。**
> 低下で止まったときは、大きい変更ならもう 1 回測ってから判断する。

### `isuh check` が見るもの

**レスポンス構造の変更はレギュレーション違反（失格）。** ここがその主な防波堤になる。
ベンチは直列でしか回せないので、正しさの確認にベンチ枠を使わないのも狙い。

1. `check.initialize` が true なら `app.initializePath` を POST する
2. `check.endpoints` を叩き、`status`（既定 200）を確かめる
3. JSON の **構造**（キーと型の形）を `check.snapshotFile` の記録と比べる

**値ではなく形だけを見る。** 値は initialize やベンチのたびに変わるため、
形だけなら「一度記録すれば以後ずっと比較できる」。

| 変化 | 判定 |
|---|---|
| キーが増えた / 消えた | **違反** |
| 型が変わった（`number` → `string`） | **違反** |
| 配列の要素の形が変わった | **違反**（どのキーが変わったかまで出す） |
| `null` ↔ 値 | 許す（null を取りうるフィールドの揺れ） |
| 配列の要素数 | 許す |
| キーの順序 | 許す |
| どちらかが空配列 | 比べない |

知らないエンドポイントは**初回に記録される**。
`check.snapshotFile`（既定 `check/structure.json`）は git で追跡し、
**アプリを触る前に 1 回走らせて記録をコミットしておく。**
記録は原則として更新しない。エンドポイントを足したときだけ増える。

### `isuh deploy` が配るもの

**アプリのバイナリ** — `app.repoLocal` で ref を sha に解決し、`git archive` で取り出した
ツリーをビルドして `hosts.app` へ配る。作業ツリーの内容ではない（未コミットの変更が
あれば警告する）。配布は `<path>.isuh-new` へ置いてから `mv` で差し替える。
実行中のバイナリを直接上書きすると `ETXTBSY` になるため。

ビルド結果は `.isuh/build/<sha>-<設定のハッシュ>/` にキャッシュする。
同じ sha の再デプロイ（= ロールバック）はビルドを丸ごと省く。

**アプリ以外の設定ファイル** — `conf/` 以下を、サーバ上の絶対パスへそのまま写す。

```
conf/_all/etc/sysctl.d/99-isucon.conf   → 全ホストの /etc/sysctl.d/99-isucon.conf
conf/isu1/etc/nginx/nginx.conf          → isu1 の /etc/nginx/nginx.conf
```

同じパスが両方にあればホスト指定側が勝つ。配った後に反映が要るものは自動で叩く。

| 配ったパス | 反映 |
|---|---|
| `/etc/nginx/**` | `middleware.nginxService` を reload |
| `/etc/mysql/**`、`/etc/my.cnf` | `middleware.dbService` を restart |
| `/etc/systemd/**` | `daemon-reload`（その後アプリを restart） |

`conf/` の中身をサーバから取り込んで管理下に置く仕組みは #10。

**デプロイの記録** — 各ホストの `$HOME/.deployed_sha` に
`sha <TAB> ref <TAB> ISO8601 <TAB> 実行者` を 1 行で書く。
`rollback`(#6) と `report`(#4) はここを読んで「今どの版が載っているか」を知る。

`isuh deploy` は stdout にデプロイした sha だけを出す（`SHA=$(isuh deploy)` が使える）。

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

**誰が書くか** — 生ファイル（`alp.txt` / `ptquery.txt` / `pprof.txt` / `resource/*`）は
`lib/collect.ts` がベンチの前後に書く。それを読んで `report.md` にまとめるのが
`lib/report.ts`、`meta.json` と `history.tsv` が `lib/run.ts`。

`runs/` の場所は `ISUH_RUNS_DIR` で上書きできる（テスト用。`ISUH_CONFIG` と同じ考え方で、
追跡している `runs/history.tsv` をテストが汚さないようにするため）。

計測ツールが入っていない場合（導入は #11）、`collect.ts` はランを失敗させず、
そのファイルに **「未取得」と理由と再現コマンド** を書く。ベンチ 1 回は高価なので、
補助的な計測が欠けただけで結果を捨てない。

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
  "bestRunId": "20261031-103512-a1b2c3d",
  "rolledBackTo": null
}
```

`status` は `pass` / `fail` / `error` のいずれか。

- `pass` — ベンチが完走しスコアが付いた
- `fail` — ベンチが完走したが整合性チェック等で失敗した
- `error` — ベンチを起動できなかった、デプロイに失敗した等

この形は `src/lib/run.ts` の `RunMeta`。

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
これがそのまま Claude Code に渡される。だから:

- 生ログを貼るだけにしない。`Resources` は `vmstat` を要約して
  **飽和しているホストを名指しする**（CPU 飽和 / IO 待ち / run queue）
- 長いものは切り詰め、元ファイルの場所を書く。読む側の予算は有限
- `Summary` に前回比とベスト比、`Diff from previous run` に
  `git diff --stat` とスコアの増減を必ず載せる
- 計測ツールが無くても「未取得」と理由を書いて欠測を明示する

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
- 正常時は `release()` が stdin を閉じる。リモートの `cat` が EOF で終わり、
  シェルが終了して fd 9 が閉じ、flock が解放される
- ローカルのファイルロックではなく**サーバ側**に置く理由は、worktree・別ターミナル・
  別マシンをまたいで効かせるため

### 異常終了でロックが残らない仕組み

3 段構えになっている。どれか 1 つでも効けばロックは解ける。

| 死に方 | 効く仕組み |
|---|---|
| 例外・`exit` | `try-finally` の `release()` |
| Ctrl-C | SIGINT はプロセスグループ全体に届くので ssh も一緒に死ぬ |
| `kill -9` | **node が死ぬと stdin のパイプが閉じる** → ssh が EOF を転送 → リモートの `cat` が終わる |

3 つめが要で、これがあるので `process.on("exit")` が走らない殺され方でも残らない。

**そのため `isuh` はプロセスが 1 本でなければならない。**
`node_modules/.bin/tsx` は子プロセスとして node を起動するので、`isuh` の PID を
kill しても実作業のプロセスが生き残り、ロックが残る。`bin/isuh` は
`node --import tsx src/cli.ts` を `exec` すること。

検証は `test/lock.sh`（`kill -9` 後にロックが解放されることを実測している）。

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
| `lib/ssh.ts` | `onHost` / `onHostShell` / `onHosts`（並列）/ `copyTo`（scp）/ `local` / `shQuote` |
| `lib/git.ts` | `resolveSha`（`config.app.repoLocal` で解決） |
| `lib/lock.ts` | `acquireLock` / `Lock.release` |
| `lib/collect.ts` | `resetLogs` / `startResourceSampling` / `collectMetrics` |
| `lib/portal.ts` | `loadCookie` / `portalRequest` / `assertLoggedIn` / `asJson` / `pick` |
| `lib/run.ts` | `RunMeta` / `writeMeta` / `readHistory` / `appendHistory` / `bestRun` / `lastGoodRun` / `localIso` |
| `lib/report.ts` | `writeReport` |
| `lib/shape.ts` | `shapeOf` / `diffShapes`（レスポンス構造の比較） |
| `lib/command.ts` | `Command` インターフェース |
| `config.ts` | `Config` 型一式 / `loadConfig` |

`ssh` や `git` は必ず `execa` の配列引数で呼ぶ。手元のシェルを経由しないので、
ローカルでの引用符の事故が起きない。

**ただし ssh のリモート側は別。** ssh は残りの引数を空白で連結した 1 本の文字列を
リモートのログインシェルに渡すため、配列で渡しても**リモートでは必ず再パースされる**。
そのため `onHost` は各要素を `shQuote` してから連結する。

- `onHost(host, ["echo", "a b"])` → リモートで `echo 'a b'`（argv として正しく届く）
- リダイレクト・パイプ・`$HOME` の展開が要るときだけ `onHostShell(host, script)` を使い、
  設定由来の値は `shQuote` で囲む

`ISUH_SSH_CONFIG` を設定すると `ssh` / `scp` に `-F <path>` が付く。
手元の `~/.ssh/config` を汚さずにホスト別名を定義したいときに使う（`test/` がこれを使う）。

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

---

## 11. 検証環境 — `test/`

当日まで本物のサーバが無い。**一度も動かしたことがないコードを当日持ち込まない**ために、
`sshd` コンテナ 3 台（`isu1` / `isu2` / `isu3`）を立てて実際に `isuh` を走らせる。

```sh
./test/up.sh      # コンテナ + 鍵 + ssh_config + ダミーのアプリリポジトリ
./test/all.sh     # typecheck + 全スイート
./test/down.sh

# 個別
./test/smoke.sh   # deploy
./test/lock.sh    # ロックと bench（直列化・kill 耐性・通しの完走）
./test/portal.sh  # ポータル（コンテナ不要。fake-portal を立てる）
./test/report.sh  # report.md / meta.json / history.tsv / スコア判定
./test/check.sh   # 検証ゲート（構造の記録と比較・ベンチを起動しないこと）
./test/rollback.sh # 戻し先の決定・ロック・自動ロールバック
```

- 設定は `test/config.test.ts`（`ISUH_CONFIG` で指す）
- ホスト別名は `test/.work/ssh_config`（`ISUH_SSH_CONFIG` で指す）。`~/.ssh/config` は触らない
- `systemd` は動かさない。`start` / `stop` / `restart` / `is-active` だけを模した
  スタブを置いている。`deploy` と `bench` が確かめたいのは ssh 越しの配布・差し替え・
  再起動と `flock` なので、systemd の再現度そのものは要らない

**確かめられないもの** — 本物の systemd、実際の負荷、ベンチマーカー、ポータル。
それらは過去問環境でのリハーサル（#18）で見る。

新しいコマンドを実装したら該当のスイートにケースを足す。
