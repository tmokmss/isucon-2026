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
 ├─ 9. （スコアの良し悪しは判定しない。report.md の前回比・ベスト比を読む）
 └─ 10. ロック解放（try-finally で異常終了時も必ず実行）
```

**スコアの良し悪しは判定しない。** 前回比・ベスト比は `report.md` の Summary に
載っているので、読む側（人かエージェント）が判断する。閾値で機械的に「低下」と
決めると、数 % 〜 10 % あるスコアのブレを誤判定する。

**戻すコマンドは持たない。** 「4. 戻す」は `isuh history` で sha を見て
`isuh deploy <sha>` する（ビルドはキャッシュされるので速い）。自動でも戻さない。
bench は毎回自分の ref をデプロイし直すので、悪い版が載ったままでも次の計測は汚れない。
載っている版が問題になるのは競技終了時だけで、それは終盤の手順（#8）で明示的に扱う。

**`check` は deploy の後に走る。** 新しいコードが載った状態を検証するため。
落ちた場合はベンチを起動せず、貴重なベンチ枠を消費しない。

---

## 2. 設定の契約 — `config.ts`

- 当日埋める値は **すべて** `config.ts` に置く。ソースに直書きしない
- 型は `src/config.ts` の `Config` インターフェースが唯一の定義。
  **項目を増やすときは型と `config.ts` を同時に更新する**
- `config.ts` は**コミットする**。エージェントが作る worktree にも入り、どこからでも `isuh` が動く。
  秘密（ポータルの Cookie）は書かず、`portal.cookieFile` が指す別ファイルに置く
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

ホスト別名そのもの（`~/.ssh/config` の定義）は `isuh setup` が `config.setup` から作る。

| フィールド | 意味 |
|---|---|
| `setup.addresses` | ホスト別名 → `{ public, private, port? }`。当日 CloudFormation の出力から埋める |
| `setup.identityFile` | 手元から入る鍵（ポータルに公開鍵を登録したもの） |

**このリポジトリは ISUCON2026 専用で、アプリ・ハーネス・構成管理を 1 つにまとめる。**

```
.
├── webapp/     # サーバの ~/webapp（app.dir）を写したもの。isuh setup apprepo が取り込む
├── conf/       # サーバの設定ファイル。isuh setup conf が取り込み、isuh deploy が配る
├── src/ bin/   # ハーネス
└── runs/       # 計測結果
```

`isuh deploy <git-ref>` や `isuh bench <git-ref>` の ref は `config.app.repoLocal`
（通常はこのリポジトリ自身）で解決する。**アプリと設定が同じ sha に乗る**ので、
エージェントは worktree 1 つでコードと設定を同時に変えられ、`isuh deploy <古い sha>` で
設定も一緒に戻る。`repoLocal` が別の場所を指すのは test/ の検証環境だけ。

| フィールド | 意味 |
|---|---|
| `app.repoLocal` | アプリと `conf/` を含む git リポジトリ。通常はこのリポジトリ |
| `app.localDir` | その中で、サーバの `app.dir` を写したディレクトリ（`webapp`） |

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
| `setup` | `[--only <step>[,<step>...]]` | 鍵配布・/etc/hosts・dotfiles・計測ツールとログ設定・アプリの git 化（下記） | #9 #11 |
| `deploy` | `[<git-ref>] [--dry-run]` | 指定 ref をビルドして全台へ配布・再起動 | #1 |
| `check` | なし | ベンチ不要の検証ゲート | #12 |
| `bench` | `[<git-ref>] [--skip-check] [--note <text>]` | 上記の通し実行 | #2 #4 |
| `history` | `[-n <count>]` | スコア台帳の表示 | #5 |
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
**HTML 本文も見る**。ポータルのトップは未ログインでも 200 を返すため
（実物で確認済み）、ステータスだけでは判定できない。未ログインなら
`/auth/login/` へのリンク、ログイン済みならログアウトのフォーム
（`<form action="/auth/logout/">`）が出る。

中身の推測は壊れやすいので、`whoamiPath` の既定値は**未ログインだと 302 を返す**
`/auth/settings/team/` にしてある（実物で確認済み。Cookie ありで 200、なしで
`/auth/login/?next=...` へ 302）。

> **ベンチ用のエンドポイントは当日にならないと分からない。**
> URL とパラメータは `config.portal` に集約し、当日は埋めるだけにする。
> 事前に検証できるのは `whoami` と `get` まで。

### `isuh deploy` が配るもの

**アプリのバイナリ** — `app.repoLocal` で ref を sha に解決し、`git archive` で取り出した
ツリーの `app.localDir/app.langDir` でビルドして `hosts.app` へ配る。作業ツリーの内容ではない
（`webapp/` か `conf/` に未コミットの変更があれば警告する。ハーネスや `runs/` の変更は見ない）。配布は `<path>.isuh-new` へ置いてから `mv` で差し替える。
実行中のバイナリを直接上書きすると `ETXTBSY` になるため。

ビルド結果は `.isuh/build/<sha>-<設定のハッシュ>/` にキャッシュする。
同じ sha の再デプロイ（= 前の版へ戻すとき）はビルドを丸ごと省く。

**アプリ以外の設定ファイル** — **バイナリと同じ sha のツリーの** `conf/` 以下を、
サーバ上の絶対パスへそのまま写す。`isuh deploy <古い sha>` で設定も一緒に戻る。
`conf/<host>/` があればアプリの無いホスト（DB 等）にも配る（バイナリと再起動はアプリのホストだけ）。

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
| `/etc/sysctl.conf`、`/etc/sysctl.d/**` | `sysctl --system`（失敗しても警告だけ） |

`/home/` 配下（unit が読む env ファイル等）は作業ユーザの持ち物として置く。それ以外は root。

中身はサーバから `isuh setup --only conf` で取り込む（下記）。チューニングの定番は
`setup/templates/` にあるが、**参考用で自動では入れない**。計測してから写す。

**デプロイの記録** — 各ホストの `$HOME/.deployed_sha` に
`sha <TAB> ref <TAB> ISO8601 <TAB> 実行者` を 1 行で書く。
「今どの版が載っているか」を知りたいときはここを読む。

`isuh deploy` は stdout にデプロイした sha だけを出す（`SHA=$(isuh deploy)` が使える）。

### `isuh setup` がやること

起動直後にまず `config.setup.addresses`（各ホストの public / private アドレス）を埋めて
`isuh setup` を流す。**どのステップも冪等**（何度流しても同じ結果）。途中で落ちたら直して流し直す。

| ステップ | やること | 触る場所 |
|---|---|---|
| `ssh` | `~/.ssh/isuh_config` を生成して `~/.ssh/config` の先頭から `Include`。サーバ間の鍵を配り、1 台目から全台へ入れることを確認 | 手元の `~/.ssh`、全台の `~/.ssh` |
| `hosts` | 全台の `/etc/hosts` に private アドレス | 全台 |
| `dotfiles` | `setup/dotfiles/` を配る。`.bashrc` は置き換えず `~/.isuh.bashrc` を読み込ませる | 全台 |
| `tools` | alp（GitHub のリリース）と pt-query-digest（同梱）を `/usr/local/bin` へ。htop / sysstat は apt で試す（失敗しても止めない） | 全台 |
| `logs` | nginx の既存 `access_log` を無効化して `conf.d/isuh-log.conf`（LTSV）を置く。MySQL に `zz-isuh-slowlog.cnf`（`long_query_time = 0`）を置いて restart | `hosts.nginx`、`hosts.db` |
| `apprepo` | `hosts.app[0]` の `app.dir` を `app.localDir`（`webapp/`）に取り込んでコミット。既にあれば飛ばす | 手元 |
| `conf` | 全台の nginx（`nginx.conf`・`sites-enabled/*`・`conf.d/*`）/ MySQL（`/etc/mysql/**/*.cnf`）/ redis / `sysctl.conf` と、アプリのホストでは systemd unit とそれが読む env ファイルを `conf/<host>/` に取り込んでコミット | 手元 |

- **既存のファイルは管理ブロックだけを書き換える。** `~/.bashrc` / `/etc/hosts` /
  サーバの `~/.ssh/config` は `# >>> isuh >>>` 〜 `# <<< isuh <<<` の間だけを置き換える
- **計測用の設定は `isuh-` の名前で別ファイルにする。** 本番の最後に `finalize`（#8）が
  このファイルを消せば元に戻る
- **`nginx -t` が落ちたら置いたものを戻して exit 60。** 無効化した `access_log` の元の
  ファイルは `/var/backups/isuh/` に退避する（`sites-enabled/` に置くと読み込まれてしまうため）
- **`conf` は手元に既にあるファイルを上書きしない。** 手元で編集した設定を、サーバの古い状態で
  潰さないため。取り込み直したいファイルは消してから流す
- **`conf` は秘密と計測用の設定を取り込まない。** `debian.cnf`（MySQL の保守用パスワード）と、
  `logs` が置いた `isuh-log.conf` / `zz-isuh-*`（finalize が消す一時的なもの）
- **コミットは取り込んだパスだけ**（他の未コミットの変更は巻き込まない）。push は人がする
- **サーバの webapp は git にしない。** deploy は手元から `git archive` で配るので、
  サーバ側に git も GitHub の鍵も要らない。取り込みは 50MB を超えるファイル・
  ビルド成果物・`node_modules` / `target` を除き、除いたものは `.gitignore` に書く。
  取り込んだコミットがそのまま元の状態の控えになる
- **alp は GitHub のリリースを手元で取り、scp で配る。** サーバの外向きの通信は細いことがあるため、サーバでは取らない。
  チェックサムを固定していて、一致しなければ使わない。`.isuh/tools/` にキャッシュするので、リハーサルで 1 回取っておけば当日は GitHub に行かない
- pt-query-digest は perl の単体スクリプトなので `setup/bin/` に同梱している

**ポートフォワード** — `ssh` ステップの後は別名が使えるので、手元からは次で足りる。

```sh
ssh -N -L 8080:localhost:80 isu1              # アプリを手元のブラウザで見る
ssh -N -L 6060:localhost:6060 isu1            # pprof（go tool pprof -http=:8081 http://localhost:6060/...）
ssh -N -L 13306:localhost:3306 isu2           # MySQL を手元のクライアントで
ssh -N -L 8080:<isu1 の private>:80 isu2      # 踏み台越し（isu2 経由で isu1 へ）
```

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
    ├── alp.txt                       # alp --format tsv の全行（nginx が複数台なら == host == の節ごと）
    ├── ptquery.txt                   # pt-query-digest --output json
    ├── pprof.txt                     # go tool pprof -top の出力
    ├── cpu.pprof                     # 生のプロファイル（go tool pprof -http で後から掘る）
    └── resource/
        ├── isu1.txt                  # vmstat（ホスト単位の CPU / mem / IO）
        ├── isu2.txt
        ├── isu3.txt
        └── top/
            └── isu1.txt ...          # top -b（プロセス単位の CPU）
```

`runs/<run-id>/` は `.gitignore` 済み（`history.tsv` のみ追跡）。

**誰が書くか** — 生ファイル（`alp.txt` / `ptquery.txt` / `pprof.txt` / `resource/*`）は
`lib/collect.ts` がベンチの前後に書く。

| いつ | 何を | どこで |
|---|---|---|
| ベンチ前 | access log / slow query log を `truncate` | `hosts.nginx` / `hosts.db` |
| ベンチ中 | `vmstat`（ホスト単位）と `top -b`（プロセス単位）を常駐させる | `hosts.all` |
| ベンチ中 | `/debug/pprof/profile?seconds=N` を curl する | `hosts.app[0]` |
| ベンチ後 | `alp --format tsv`（nginx が複数台なら全台、ホストごとに並べる） | `hosts.nginx` |
| ベンチ後 | `pt-query-digest --output json` | `hosts.db` |

**生ファイルは機械向けの形式で取り、`report.md` では畳む。** report.md はそのまま
エージェントに渡るので、alp の既定の表（20 列近く・罫線）や pt-query-digest の既定の
レポート（1 クエリ 40 行ほど）を貼ると、判断に要らない装飾で読む側の予算を使い切る。
畳むのは `lib/digest.ts`。形式が想定と違う（未取得の理由が書いてある等）ときは元のまま載せる。

alp の `--limit` は表示件数ではなく **URI の種類数の上限**で、超えると集計ごと落ちる。
`alpMatches` を埋める前は ID 入りのパスがばらけて簡単に超えるので、collect は大きな値を
渡して全行を取り、上位 `metrics.alpLimit` 件への絞り込みは report.md 側で行う。

**pprof はベンチの開始と同時に取り始める。** `?seconds=N` は「叩いてから N 秒間」を
測るので、ベンチの後に叩くとアイドル中を測ってしまう。N はベンチの長さに合わせ、
止めるときは N 秒経つまで待つ（途中で切るとプロファイルが返ってこない）。
`pprofUrl` は `hosts.app[0]` から見た URL で、そのホスト上で curl する。
pprof のポートは外に開いていないことが多いため。

`vmstat` はホスト全体の値しか出さないので、アプリと MySQL が同じホストにいると
どちらが CPU を食っているか分からない。そのため `top -b` も並行で取り、
`report.md` の Resources 節で「どのプロセスが食っているか」まで要約する。それを読んで `report.md` にまとめるのが
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
  "bestRunId": "20261031-103512-a1b2c3d"
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
<!-- 列を絞った TSV: share / count / 4xx / 5xx / method / uri / sum_s / avg_s / p99_s / max_s / avg_bytes。
     上位 metrics.alpLimit 件。alpMatches が空なら「集計は当てにならない」と書く -->

## Slow query (pt-query-digest)
<!-- 全体の件数と合計時間、1 行 1 クエリの TSV:
     rank / share / calls / total_s / avg_ms / p95_ms / rows_examined_avg / rows_sent_avg / query。
     上位 3 件は EXPLAIN にそのまま使える実例を付ける -->

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
| `EXIT.SETUP` | `60` | セットアップ失敗（`isuh setup`） |
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
| `lib/run.ts` | `RunMeta` / `writeMeta` / `readHistory` / `appendHistory` / `bestRun` / `localIso` |
| `lib/report.ts` | `writeReport` |
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
| D | #9 #10 #11 | `src/commands/setup.ts`、`src/commands/deploy.ts`（conf/ の配布） |
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
./test/up.sh      # コンテナ + 鍵 + ssh_config + ダミーのリポジトリ（webapp/ と harness.txt）
./test/all.sh     # typecheck + 全スイート
./test/down.sh

# 個別
./test/smoke.sh   # deploy
./test/lock.sh    # ロックと bench（直列化・kill 耐性・通しの完走）
./test/portal.sh  # ポータル（コンテナ不要。fake-portal を立てる）
./test/report.sh  # report.md / meta.json / history.tsv / スコア判定
./test/metrics.sh # ベンチ中の計測（pprof の開始時機・top・alp の全台集計）
./test/digest.sh  # report.md の alp / pt-query-digest の節が列を絞った表になるか
./test/setup.sh   # setup（本物の nginx / MySQL に設定を入れ、alp / pt-query-digest で読む）
```

- 設定は `test/config.test.ts`（`ISUH_CONFIG` で指す）
- ホスト別名は `test/.work/ssh_config`（`ISUH_SSH_CONFIG` で指す）。`~/.ssh/config` は触らない
- コンテナのアドレスは `10.99.0.11`〜`13` に固定（`setup` が `/etc/hosts` に書く private アドレス）
- nginx と MySQL は本物が入っている（`setup logs` の検証用）
- `systemd` は動かさない。`start` / `stop` / `restart` / `is-active` だけを模した
  スタブを置いている。`deploy` と `bench` が確かめたいのは ssh 越しの配布・差し替え・
  再起動と `flock` なので、systemd の再現度そのものは要らない

**確かめられないもの** — 本物の systemd、実際の負荷、ベンチマーカー、ポータル。
それらは過去問環境でのリハーサル（#18）で見る。

新しいコマンドを実装したら該当のスイートにケースを足す。
