# 検証環境

競技サーバの代わりに **sshd コンテナ 3 台**（`isu1` / `isu2` / `isu3`）を立てて、
`isuh` を実際に走らせて確かめるための場所。

当日まで本物のサーバが無いので、ここが無いと `deploy` も `bench` のロックも
「書いたが一度も動かしたことがない」状態になる。それは避ける。

## 使い方

```sh
./test/up.sh        # コンテナ起動 + 鍵/ssh_config + ダミーのリポジトリ（webapp/ と harness.txt）の用意
./test/all.sh       # typecheck + 全スイート
./test/down.sh      # 停止
```

個別に回すこともできる。

| | 確かめること |
|---|---|
| `./test/smoke.sh` | deploy — 配布・古い sha への復帰・dry-run・ビルド失敗・conf を同じ sha から全台へ配布・バイナリ以外の webapp/（sql/ 等）の同期 |
| `./test/lock.sh` | ロックと bench — 直列化・kill 耐性・通しの完走・計測の回収 |
| `./test/portal.sh` | ポータル — Cookie・ログイン切れ判定・CSRF・enqueue/wait/score（コンテナ不要） |
| `./test/report.sh` | report.md の見出しと中身・meta.json・history.tsv・スコア判定 |
| `./test/metrics.sh` | ベンチ中の計測 — pprof をベンチ開始と同時に取る・top・alp の全台集計 |
| `./test/digest.sh` | report.md の alp / pt-query-digest 節 — 列を絞った表・上位 N 件・alpMatches が空のときの警告 |
| `./test/setup.sh` | setup — ssh・/etc/hosts・dotfiles・ツール・nginx と MySQL のログ設定・apprepo・conf の取り込み・冪等性 |
| `./test/bench-local.sh` | `bench.runner=local` — 偽のベンチマーカーを ssh で実行し、出力からスコアと成否を読む・落ちても出力とロックが正しく残る |

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
| `metrics.sh` | ベンチ中の計測の検証 |
| `digest.sh` / `digest-probe.ts` | report.md で alp / pt-query-digest の出力を畳む処理の検証 |
| `setup.sh` / `setup-probe.ts` | `setup` の通し検証と、手元の `~/.ssh` を書き換える部分の検証 |
| `top-probe.ts` | top の要約ロジックの確認 |
| `fake-portal.ts` | ポータルの代役。Django 風の 302 / CSRF / ジョブ進行を返す |
| `lock-probe.ts` | ロックの取り合いを再現するための小さなプログラム |
| `.work/` | 生成物（鍵・ssh_config・ダミーリポジトリ）。git 管理外 |

`up.sh` が作るダミーリポジトリは 3 つの revision を持つ。

| revision | 中身 |
|---|---|
| `main~1` | `/version` が `v1` を返す |
| `main` | `/version` が `v2` を返す |
| `broken` | ビルドが通らない |

## できること・できないこと

**できる** — ssh 越しの配布、実行中バイナリの差し替え、サービス再起動、sha の記録、
設定ファイルの配布、`flock` によるロックの取り合い、ビルド失敗時に止まること、
ポータルの Cookie/CSRF、レポートと台帳の生成。

テストは `ISUH_RUNS_DIR` を `.work/` 以下に向けるので、追跡している
`runs/history.tsv` は汚れない。

**できない** — 本物の systemd、実際の負荷、ベンチマーカー、ポータル。
それらは過去問環境でのリハーサル（#18）で見る。

`~/.ssh/config` は触らない。ホスト別名は `.work/ssh_config` に閉じ、
`ISUH_SSH_CONFIG` で `isuh` に渡す。
