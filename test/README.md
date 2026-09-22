# 検証環境

競技サーバの代わりに **sshd コンテナ 3 台**（`isu1` / `isu2` / `isu3`）を立てて、
`isuh` を実際に走らせて確かめるための場所。

当日まで本物のサーバが無いので、ここが無いと `deploy` も `bench` のロックも
「書いたが一度も動かしたことがない」状態になる。それは避ける。

## 使い方

```sh
./test/up.sh        # コンテナ起動 + 鍵/ssh_config + ダミーアプリリポジトリの用意
./test/smoke.sh     # deploy の通し検証
./test/down.sh      # 停止
```

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
| `.work/` | 生成物（鍵・ssh_config・ダミーリポジトリ）。git 管理外 |

`up.sh` が作るダミーリポジトリは 3 つの revision を持つ。

| revision | 中身 |
|---|---|
| `main~1` | `/version` が `v1` を返す |
| `main` | `/version` が `v2` を返す |
| `broken` | ビルドが通らない |

## できること・できないこと

**できる** — ssh 越しの配布、実行中バイナリの差し替え、サービス再起動、sha の記録、
設定ファイルの配布、`flock` によるロックの取り合い、ビルド失敗時に止まること。

**できない** — 本物の systemd、実際の負荷、ベンチマーカー、ポータル。
それらは過去問環境でのリハーサル（#18）で見る。

`~/.ssh/config` は触らない。ホスト別名は `.work/ssh_config` に閉じ、
`ISUH_SSH_CONFIG` で `isuh` に渡す。
