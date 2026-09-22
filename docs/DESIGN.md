# 設計契約

このドキュメントは **全レーン共有の契約** である。
並列で実装するため、ここに書かれたインターフェースを独断で変更すると他レーンが壊れる。
**変更が必要なら実装せずに先に相談すること。**（→ `CLAUDE.md`）

契約の範囲内での実装方法は自由。

---

## 1. 全体像

競技中にやることは 4 つしかない。

1. 変更をデプロイする
2. ベンチを回す
3. 計測結果を読む
4. 良くなっていなければ戻す

`bin/bench` がこの 4 つを 1 コマンドに畳む。**並列に走る他の作業が計測を汚染しない**よう、
デプロイからレポート生成までを 1 つのクリティカルセクションに入れる。

```
bin/bench [<git-ref>]
 ├─ 1. ロック取得 ........................ 失敗 → exit 10
 ├─ 2. git ref を sha に解決
 ├─ 3. bin/deploy <sha> .................. 失敗 → exit 20
 ├─ 4. bin/check ......................... 失敗 → exit 30（ベンチを起動しない）
 ├─ 5. ログのリセット + リソース計測開始
 ├─ 6. bin/portal enqueue → 完了待ち ..... 失敗 → exit 40
 ├─ 7. 計測停止・ログ収集（alp / pt-query-digest / pprof / dstat）
 ├─ 8. runs/<run-id>/ に report.md と meta.json を生成、history.tsv に追記
 ├─ 9. スコア判定 ........................ 低下 → exit 50（AUTO_ROLLBACK なら戻す）
 └─ 10. ロック解放（trap で異常終了時も必ず実行）
```

**`bin/check` は deploy の後に走る。** 新しいコードが載った状態を検証するため。
落ちた場合はベンチを起動せず、貴重なベンチ枠を消費しない。

---

## 2. 設定の契約 — `config.env`

- 当日埋める値は **すべて** `config.env` に置く。スクリプトに直書きしない
- 項目は `config.env.example` が唯一の定義。**項目を増やすときは example も同時に更新する**
- 読み込みは `bin/_common.sh` の `load_config` が行う。各スクリプトは自前で読まない
- 未設定の必須項目があれば `exit 65` で落ちる

ホストの役割は次の変数で表す。値は `~/.ssh/config` 上のホスト別名。

| 変数 | 意味 |
|---|---|
| `ISU_HOSTS` | 全ホスト。再起動試験の対象 |
| `ISU_APP_HOSTS` | アプリを動かすホスト。deploy の配布先 |
| `ISU_NGINX_HOSTS` | nginx のホスト。access log の収集元 |
| `ISU_DB_HOST` | DB のホスト。slow query log の収集元 |
| `ISU_PRIMARY` | ベンチのリクエスト先。**ロックの保持先** |

アプリのコードは**このハーネスとは別のリポジトリ**にある。
`bin/deploy <git-ref>` や `bin/bench <git-ref>` の ref は `APP_REPO_LOCAL`
（手元のアプリリポジトリのチェックアウト）で解決する。ハーネス自身のリポジトリではない。
これによりアプリ側を worktree で並列に触る運用がそのまま成立する。

---

## 3. CLI の契約 — `bin/*`

全コマンド共通:

- `-h` / `--help` で usage を表示して `exit 0`
- `--dry-run` があるものは「何をするか」を stderr に出して実行しない
- 引数不正は `exit 64`、設定不備は `exit 65`
- 実行前に `load_config` を通す

| コマンド | 引数 | 役割 | 実装 |
|---|---|---|---|
| `bin/setup` | `[--only <step>]` | 鍵配布・/etc/hosts・dotfiles・計測ツール導入 | #9 #11 |
| `bin/deploy` | `[<git-ref>]` (既定 `HEAD`) | 指定 ref をビルドして全台へ配布・再起動 | #1 |
| `bin/check` | なし | ベンチ不要の検証ゲート | #12 |
| `bin/bench` | `[<git-ref>] [--skip-check] [--note <text>]` | 上記の通し実行 | #2 #4 |
| `bin/history` | `[-n <count>]` | スコア台帳の表示 | #5 |
| `bin/rollback` | `[<sha>]` (既定: 直前の good) | 指定 revision へ戻してデプロイ | #6 |
| `bin/portal` | `<subcommand>` | ポータル操作（下記） | #3 |
| `bin/monitor` | なし | 全台のリソースを tmux で 1 画面に | #17 |
| `bin/verify-restart` | `[--times <n>]` | ランダム順で再起動 → initialize 無しでベンチ | #7 |
| `bin/finalize` | なし | 終盤チェックリスト | #8 |

### `bin/portal` のサブコマンド

| サブコマンド | 役割 |
|---|---|
| `whoami` | Cookie が有効かの確認。認証必須ページを引けるか |
| `enqueue` | ベンチを登録し、ジョブ ID を **stdout に出す** |
| `status <job-id>` | ジョブ状態を出す（`pending` / `running` / `done` / `failed`） |
| `wait <job-id>` | 完了まで待ち、結果 JSON を stdout に出す |
| `score` | 最新の確定スコアを stdout に出す |

> **ベンチ用のエンドポイントは当日にならないと分からない。**
> URL とパラメータは `config.env` の `PORTAL_*` に集約し、当日は埋めるだけにする。
> 事前に検証できるのは `whoami` 相当（Cookie 付き curl で認証必須ページを引く）まで。

---

## 4. 成果物の契約 — `runs/`

### run-id

```
<YYYYmmdd>-<HHMMSS>-<sha7>      例: 20261031-103512-a1b2c3d
```

辞書順 = 時系列順になること。

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
  "run_id": "20261031-103512-a1b2c3d",
  "started_at": "2026-10-31T10:35:12+09:00",
  "finished_at": "2026-10-31T10:37:02+09:00",
  "ref": "HEAD",
  "sha": "a1b2c3d4e5f6...",
  "score": 12345,
  "status": "pass",
  "note": "add index on competition_id",
  "bench_job_id": "4321",
  "prev_run_id": "20261031-102011-9f8e7d6",
  "best_run_id": "20261031-103512-a1b2c3d"
}
```

`status` は `pass` / `fail` / `error` のいずれか。

- `pass` — ベンチが完走しスコアが付いた
- `fail` — ベンチが完走したが整合性チェック等で失敗した
- `error` — ベンチを起動できなかった、デプロイに失敗した等

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
<!-- PPROF_ENABLED=false なら「未取得」と書く -->

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
- `score` は `status != pass` のとき `0`
- `note` にタブ・改行を含めない

---

## 5. 終了コードの契約

呼び出し側が失敗理由で分岐するため、**必ずこの値を使う**。

| コード | 意味 |
|---|---|
| `0` | 成功 |
| `10` | ロック取得失敗（タイムアウト） |
| `20` | デプロイ失敗 |
| `30` | 検証ゲート（`bin/check`）失敗 |
| `40` | ベンチ失敗 |
| `50` | スコア低下（`SCORE_REGRESSION_RATIO` を下回った） |
| `64` | 引数不正 |
| `65` | 設定不備 |
| `70` | 未実装 |

`bin/_common.sh` が `EX_LOCK` `EX_DEPLOY` `EX_CHECK` `EX_BENCH` `EX_REGRESSION`
`EX_USAGE` `EX_CONFIG` `EX_NOTIMPL` として定義する。数値を直書きしない。

---

## 6. ログ出力の契約

- **stdout は機械可読な結果のみ。** `bin/portal enqueue` ならジョブ ID だけ、
  `bin/history` なら表だけ
- 進捗・警告・エラーは **すべて stderr**。`log_info` / `log_warn` / `log_error` / `die` を使う
- 形式: `[HH:MM:SS] LEVEL message`
- `VERBOSE=1` で `log_debug` が有効になる

これを守ると `sha=$(bin/deploy HEAD)` のようなパイプ利用が安全にできる。

---

## 7. ベンチのロックプロトコル

**デプロイだけをロックしても不十分。** ベンチ実行中に別のデプロイが走れば計測は壊れる。
`bin/bench` の 1〜10 全体が 1 つのクリティカルセクション。

### 保持方法

`ISU_PRIMARY` 上の `flock` を、**ssh セッションを開いたまま保持する**。

```sh
ssh "$ISU_PRIMARY" "exec 9>$BENCH_LOCK_PATH; flock -w $BENCH_LOCK_TIMEOUT 9 || exit 10; \
                    echo \"\$\$ $owner \$(date +%s)\" >&2; echo ACQUIRED; cat"
```

- 末尾の `cat` が stdin を開いたまま保つ → **ssh プロセスが生きている間だけロックが保持される**
- ローカルの ssh プロセスを kill すると接続が切れ、flock が自動解放される
  → **Ctrl-C でも異常終了でもロックが残らない**（#2 の完了条件）
- ローカルのファイルロックではなく**サーバ側**に置く理由は、worktree・別ターミナル・
  別マシンをまたいで効かせるため

### 保持者の表示

取得できなかった場合、**誰がいつから保持しているか**を stderr に出して `exit 10`。
`$BENCH_LOCK_PATH.owner` に `ホスト名 / ユーザ / 開始時刻 / git ref` を書いておく。

### 迂回の禁止

ロックを取らずに `scp` や `systemctl restart` をすると、走っているベンチの計測が壊れる。
**デプロイは必ず `bin/deploy` / `bin/bench` 経由。**

---

## 8. `bin/_common.sh` が提供するもの

全スクリプトが先頭で読み込む。ここにある関数は再実装しない。

| 関数 | 役割 |
|---|---|
| `repo_root` | リポジトリルートの絶対パス |
| `load_config` | `config.env` を読み、必須項目を検証（不備は `exit 65`） |
| `require_config <名前>...` | 指定した変数が空でないことを検証 |
| `log_debug/info/warn/error` | stderr へ時刻付きで出力 |
| `die <code> <message>` | エラーを出して指定コードで終了 |
| `not_implemented <issue>` | 未実装である旨と担当 Issue を出して `exit 70` |
| `on_host <host> <cmd>` | 1 ホストでコマンド実行 |
| `on_hosts "<hosts>" <cmd>` | 複数ホストで順に実行 |
| `resolve_sha <ref>` | git ref を完全な sha に解決 |
| `new_run_id <sha>` | run-id を生成 |
| `run_dir <run-id>` | `runs/<run-id>` の絶対パスを返し、無ければ作る |
| `usage_guard "$@"` | `-h` / `--help` を処理 |
| `lock_acquire <owner>` | ベンチのロックを取得。失敗は `exit 10`（実装 #2） |
| `lock_release` | ロックを解放。多重呼び出し安全であること（実装 #2） |

---

## 9. 実装レーンの割り当て

触るファイルが重なるものは**同一レーンで直列に**扱う。

| レーン | Issue | 主に触るファイル |
|---|---|---|
| A | #1 #2 | `bin/deploy` `bin/bench` `bin/_common.sh`（ロック） |
| B | #4 #5 | `runs/` 生成まわり、`bin/history` |
| C | #15 #19 | `docs/` のみ |
| D | #9 #10 #11 | `bin/setup`、`conf/` |
| E | #3 | `bin/portal`（要: 手元のブラウザ Cookie） |

第 2 波（A・B のマージ後）: #6 #12 → #7 #8 → #18。

---

## 10. 未確定事項

当日にならないと確定しないもの。**確定したら `config.env` を埋めるだけで済む**ように作る。

- ベンチ起動のエンドポイントとパラメータ（`PORTAL_ENQUEUE_PATH` / `PORTAL_JOB_PATH`）
- サーバの台数と役割（`ISU_*`）
- アプリのサービス名・バイナリ名・ビルドコマンド（`APP_*`）
- alp の URI 正規化パターン（`ALP_MATCHES`）
- ログの出力先（`NGINX_ACCESS_LOG` / `DB_SLOW_LOG`）
