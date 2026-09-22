# 検証環境

競技サーバの代わりに **sshd コンテナ 3 台**（`isu1` / `isu2` / `isu3`）を立てて、
`isuh` を実際に走らせて確かめるための場所。

当日まで本物のサーバが無いので、ここが無いと `deploy` も `bench` のロックも
「書いたが一度も動かしたことがない」状態になる。それは避ける。

## 使い方

```sh
./test/up.sh        # コンテナ起動 + 鍵/ssh_config + ダミーアプリリポジトリの用意
./test/all.sh       # typecheck + 全スイート
./test/down.sh      # 停止
```

個別に回すこともできる。

| | 確かめること |
|---|---|
| `./test/smoke.sh` | deploy — 配布・古い sha への復帰・dry-run・ビルド失敗・conf 配布 |
| `./test/lock.sh` | ロックと bench — 直列化・kill 耐性・通しの完走・計測の回収 |
| `./test/portal.sh` | ポータル — Cookie・ログイン切れ判定・CSRF・enqueue/wait/score（コンテナ不要） |
| `./test/report.sh` | report.md の見出しと中身・meta.json・history.tsv・スコア判定 |
| `./test/check.sh` | 検証ゲート — 構造の記録と比較、違反時にベンチを起動しないこと |

手で触るとき:

```sh
export ISUH_CONFIG=$PWD/test/config.test.ts
export ISUH_SSH_CONFIG=$PWD/test/.work/ssh_config
bin/isuh deploy main
ssh -F test/.work/ssh_config isu1 'curl -s localhost:8080/version'
```

## 中身

| | |
|---|---|
| `Dockerfile` | Ubuntu 24.04 + sshd + sudo + `flock` + 最小の `systemctl` |
| `systemctl` | systemd の代役。ユニット定義（`/etc/isuh-units/<unit>`）のコマンドを pidfile で管理する |
| `docker-compose.yml` | 3 台。ポートは 22101 / 22102 / 22103 |
| `config.test.ts` | この環境向けの `Config`。`ISUH_CONFIG` で指す |
| `fixtures/app/` | ダミーの競技アプリ（`/version` を返すだけの Go） |
| `up.sh` | 鍵生成・ssh_config 生成・コンテナ起動・ダミーリポジトリ作成 |
| `smoke.sh` | `deploy` の通し検証 |
| `lock.sh` | ロックと `bench` の通し検証 |
| `portal.sh` | ポータル操作の検証 |
| `report.sh` | レポートと台帳の検証 |
| `check.sh` | 検証ゲートの検証 |
| `shape-probe.ts` | 構造比較の細かい挙動（null の揺れ・配列・入れ子）の確認 |
| `fake-portal.ts` | ポータルの代役。Django 風の 302 / CSRF / ジョブ進行を返す |
| `lock-probe.ts` | ロックの取り合いを再現するための小さなプログラム |
| `.work/` | 生成物（鍵・ssh_config・ダミーリポジトリ）。git 管理外 |

`up.sh` が作るダミーリポジトリは 3 つの revision を持つ。

| revision | 中身 |
|---|---|
| `main~1` | `/version` が `v1` を返す |
| `main` | `/version` が `v2` を返す |
| `broken` | ビルドが通らない |
| `badshape` | ビルドは通るが `/api/items` の `name` が `title` に変わる（= レギュレーション違反） |

## できること・できないこと

**できる** — ssh 越しの配布、実行中バイナリの差し替え、サービス再起動、sha の記録、
設定ファイルの配布、`flock` によるロックの取り合い、ビルド失敗時に止まること、
ポータルの Cookie/CSRF、レポートと台帳の生成、スコア低下の判定、
レスポンス構造の変更の検出。

テストは `ISUH_RUNS_DIR` を `.work/` 以下に向けるので、追跡している
`runs/history.tsv` は汚れない。

**できない** — 本物の systemd、実際の負荷、ベンチマーカー、ポータル。
それらは過去問環境でのリハーサル（#18）で見る。

`~/.ssh/config` は触らない。ホスト別名は `.work/ssh_config` に閉じ、
`ISUH_SSH_CONFIG` で `isuh` に渡す。
