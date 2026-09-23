#!/bin/sh
# isuh setup の検証。先に ./test/up.sh を実行しておくこと。
#   - 生成した ssh 設定だけで全台に入れ、サーバ間でも ssh できること
#   - /etc/hosts・dotfiles・計測ツール・ログ設定が入り、alp / pt-query-digest が本物のログを読めること
#   - nginx -t が落ちたら変更を戻すこと
#   - ~/webapp が webapp/ に、サーバの設定が conf/<host>/ に取り込まれること（大きいファイル・秘密は入れない）
#   - 2 回流しても同じ結果になること（冪等）
set -eu
cd "$(dirname "$0")"
TEST_DIR=$(pwd)
ROOT=$(cd .. && pwd)
WORK="$TEST_DIR/.work"
ISUH="$ROOT/bin/isuh"
ERR="$WORK/setup.err"

# app.repoLocal だけ差し替えた設定（既存の app-repo は他のテストが使う）
CFG="$WORK/config.setup.ts"
sed -e 's|join(here, ".work", |join(here, |g' \
  -e 's|join(here, "app-repo")|join(here, "apprepo-setup")|' \
  "$TEST_DIR/config.test.ts" >"$CFG"
export ISUH_CONFIG="$CFG"
# setup が生成する ssh 設定。up.sh が作った ssh_config とは別にして、生成物だけで入れることを見る
export ISUH_SSH_CONFIG="$WORK/ssh_config.setup"
SSH="ssh -F $ISUH_SSH_CONFIG"

passed=0
ng() {
  printf '  NG  %s\n' "$*" >&2
  echo "--- 直前の stderr ---" >&2
  tail -n 25 "$ERR" >&2 || true
  exit 1
}
ok() {
  passed=$((passed + 1))
  printf '  ok  %s\n' "$*"
}
run() {
  set +e
  "$ISUH" setup "$@" >/dev/null 2>"$ERR"
  rc=$?
  set -e
}
count_on() { $SSH "$1" "grep -cxF '$2' $3 || true"; }

rm -f "$ISUH_SSH_CONFIG"
# 取り込み先 = アプリ・ハーネス・conf/ を 1 つにまとめたリポジトリ（ハーネス側のファイルだけある状態）
REPO="$WORK/apprepo-setup"
rm -rf "$REPO" "$WORK/app-remote.git"
mkdir -p "$REPO"
git -C "$REPO" init -q -b main
git -C "$REPO" config user.email isuh-test@example.invalid
git -C "$REPO" config user.name "isuh test"
echo harness >"$REPO/harness.txt"
git -C "$REPO" add -A
git -C "$REPO" commit -qm "harness"

# サーバの webapp を当日っぽくする: ソース・ビルド成果物・巨大ファイル
UP="ssh -F $WORK/ssh_config"
$UP isu1 'mkdir -p ~/webapp/go ~/webapp/sql
  printf "package main\n" > ~/webapp/go/main.go
  printf "CREATE TABLE t (id INT);\n" > ~/webapp/sql/init.sql
  [ -f ~/webapp/go/isuapp ] || printf "binary" > ~/webapp/go/isuapp
  truncate -s 60M ~/webapp/big.bin'
# 前回の実行の名残（nginx のロールバック検証用ファイル）を消す
$UP isu1 'sudo rm -f /etc/nginx/conf.d/zz-broken.conf'
# alp をリリースから取ってくる経路を毎回通す。手元のキャッシュは壊れた中身にしておき、
# チェックサムが合わないキャッシュは使わずに取り直すことを見る
for h in isu1 isu2 isu3; do $UP "$h" 'sudo rm -f /usr/local/bin/alp'; done
TOOLS="$ROOT/.isuh/tools"
mkdir -p "$TOOLS"
for a in amd64 arm64; do printf 'broken' >"$TOOLS/alp_1.0.21_linux_$a.tar.gz"; done

echo "== 1. 引数と設定の検証"
run --only bogus
[ "$rc" = 64 ] || ng "不明なステップで exit=$rc（64 のはず）"
BADCFG="$WORK/config.setup-bad.ts"
sed -e 's|private: "10.99.0.13"|private: ""|' "$CFG" >"$BADCFG"
set +e
ISUH_CONFIG="$BADCFG" "$ISUH" setup --only ssh >/dev/null 2>"$ERR"
rc=$?
set -e
[ "$rc" = 65 ] || ng "アドレス未設定で exit=$rc（65 のはず）"
grep -q "isu3" "$ERR" || ng "未設定のホスト名が出ていません"
ok "不明なステップは 64、アドレス未設定は 65（どのホストか出る）"

echo "== 2. nginx -t が落ちたら戻す"
# 既存の設定と名前がぶつかって nginx -t が落ちる状況を作る
$UP isu1 "printf 'log_format isuh_ltsv \"x\";\n' | sudo tee /etc/nginx/conf.d/zz-broken.conf >/dev/null"
run --only ssh,logs
[ "$rc" = 60 ] || ng "nginx -t が落ちたのに exit=$rc（60 のはず）"
grep -q "変更を戻しました" "$ERR" || ng "戻したことが出ていません"
[ "$($SSH isu1 'test -f /etc/nginx/conf.d/isuh-log.conf && echo yes || echo no')" = no ] || ng "isuh-log.conf が残っています"
$SSH isu1 'sudo rm -f /etc/nginx/conf.d/zz-broken.conf; sudo nginx -t' 2>/dev/null || ng "戻した後の nginx -t が通りません"
ok "nginx -t が落ちたら exit 60 で、置いた設定を戻す"

echo "== 3. 全ステップを通す"
# ハーネス側の未コミットの変更。apprepo / conf のコミットに巻き込まないことを後で見る
echo dirty >>"$REPO/harness.txt"
run
[ "$rc" = 0 ] || ng "exit=$rc"
ok "isuh setup が完走した"

echo "== 4. ssh"
grep -q "Port 22101" "$ISUH_SSH_CONFIG" || ng "生成した ssh 設定にポートがありません"
[ "$($SSH isu1 hostname)" = isu1 ] || ng "生成した設定で isu1 に入れません"
[ "$($SSH isu1 ssh isu3 hostname)" = isu3 ] || ng "isu1 から isu3 へ ssh できません"
[ "$($SSH isu3 ssh isu2 hostname)" = isu2 ] || ng "isu3 から isu2 へ ssh できません"
ok "生成した設定で全台に入れ、サーバ間でも ssh できる"

echo "== 5. /etc/hosts"
for h in isu1 isu2 isu3; do
  [ "$($SSH "$h" 'getent hosts isu3' | awk 'NR==1{print $1}')" = 10.99.0.13 ] || ng "$h で isu3 が private に解決されません"
done
ok "全台で isuN が private アドレスに解決される"

echo "== 6. dotfiles"
$SSH isu2 'test -f ~/.isuh.bashrc && test -f ~/.vimrc && test -f ~/.tmux.conf && test -f ~/.inputrc' || ng "dotfiles が揃っていません"
[ "$($SSH isu2 'bash -ic "type -t ll" 2>/dev/null')" = alias ] || ng "対話シェルで ~/.isuh.bashrc が読まれていません"
ok "dotfiles が配られ、対話シェルで読まれる"

echo "== 7. 計測ツール"
$SSH isu2 'sudo alp --version' | grep -q "1.0.21" || ng "alp 1.0.21 が入っていません"
$SSH isu3 'sudo pt-query-digest --version' | grep -q "pt-query-digest" || ng "pt-query-digest が動きません"
grep -q "alp を取得しています\|releases/download/v1.0.21" "$ERR" || ng "alp をリリースから取得していません"
[ "$(wc -c <"$TOOLS/alp_1.0.21_linux_$($UP isu1 'uname -m | sed "s/x86_64/amd64/; s/aarch64/arm64/"').tar.gz" | tr -d ' ')" -gt 1000000 ] ||
  ng "壊れたキャッシュが取り直されていません"
ok "alp をリリースから取得して配り（壊れたキャッシュは取り直す）、pt-query-digest と共に sudo で動く"

echo "== 8. nginx の access log が LTSV になり、alp が読める"
[ "$($SSH isu1 "grep -c '^[[:space:]]*access_log /var/log/nginx/access.log;' /etc/nginx/nginx.conf || true")" = 0 ] ||
  ng "既存の access_log が残っています"
$SSH isu1 'curl -s -o /dev/null http://localhost/isuh-setup-probe'
$SSH isu1 'sudo alp ltsv --file /var/log/nginx/access.log' | grep -q "/isuh-setup-probe" ||
  ng "alp が access log を読めません"
ok "既存の access_log を無効化し、alp ltsv で集計できる"

echo "== 9. MySQL の slow query log"
[ "$($SSH isu3 "sudo mysql -NBe 'SELECT @@global.long_query_time'")" = "0.000000" ] || ng "long_query_time が 0 ではありません"
# 他のクエリが溜まっていると、この 1 件は pt-query-digest の上位から外れて MISC にまとめられる
$SSH isu3 'sudo truncate -s 0 /var/log/mysql/mysql-slow.log'
$SSH isu3 "sudo mysql -NBe 'SELECT 42 AS isuh_setup_probe'" >/dev/null
$SSH isu3 'sudo pt-query-digest /var/log/mysql/mysql-slow.log' | grep -q "isuh_setup_probe" ||
  ng "pt-query-digest がクエリを拾えません"
ok "long_query_time=0 で記録され、pt-query-digest が読める"

echo "== 10. apprepo — webapp/ に取り込む"
git -C "$REPO" ls-files --error-unmatch webapp/go/main.go webapp/sql/init.sql >/dev/null 2>&1 || ng "ソースが取り込まれていません"
git -C "$REPO" ls-files --error-unmatch webapp/big.bin >/dev/null 2>&1 && ng "60MB のファイルが入っています"
git -C "$REPO" ls-files --error-unmatch webapp/go/isuapp >/dev/null 2>&1 && ng "ビルド成果物が入っています"
grep -qx "/big.bin" "$REPO/webapp/.gitignore" || ng "大きいファイルが webapp/.gitignore にありません"
[ "$(git -C "$REPO" log --format=%s | grep -c 'webapp/ に取り込む')" = 1 ] || ng "取り込みのコミットがありません"
ok "ソースだけを webapp/ に取り込んでコミットした"

echo "== 11. conf — サーバの設定を conf/<host>/ に取り込む"
C="$REPO/conf"
[ -f "$C/isu1/etc/nginx/nginx.conf" ] || ng "nginx.conf がありません"
grep -q "# isuh: access_log" "$C/isu1/etc/nginx/nginx.conf" || ng "取り込んだ nginx.conf がサーバの今の状態ではありません"
[ -f "$C/isu1/etc/nginx/sites-enabled/default" ] || ng "sites-enabled（シンボリックリンク）の中身がありません"
[ -f "$C/isu3/etc/mysql/mysql.conf.d/mysqld.cnf" ] || ng "mysqld.cnf がありません"
[ -f "$C/isu1/etc/systemd/system/isuapp.go.service" ] || ng "アプリの systemd unit がありません"
grep -qx "ISUCON_DB_HOST=isu3" "$C/isu1/home/isucon/env.sh" || ng "unit が読む env ファイルがありません"
[ -e "$C/isu3/etc/systemd/system/isuapp.go.service" ] && ng "アプリの無いホストの unit まで取り込んでいます"
[ -n "$(find "$C" -name 'debian.cnf')" ] && ng "パスワードを含む debian.cnf を取り込んでいます"
[ -n "$(find "$C" -name 'zz-isuh-*' -o -name 'isuh-log.conf')" ] && ng "setup logs が置いた計測用の設定を取り込んでいます"
cmp -s "$C/isu1/home/isucon/env.sh" /dev/stdin <<EOF || ng "中身がバイト単位で一致しません（末尾の改行など）"
ISUCON_DB_HOST=isu3
ISUCON_DB_USER=isucon
EOF
git -C "$REPO" ls-files --error-unmatch conf/isu3/etc/mysql/mysql.conf.d/mysqld.cnf >/dev/null 2>&1 || ng "conf/ がコミットされていません"
[ "$(git -C "$REPO" status --porcelain -- harness.txt)" = " M harness.txt" ] || ng "他のファイルの変更を巻き込んでコミットしています"
ok "nginx / MySQL / unit / env を取り込み、秘密と計測用の設定は除いてコミットした"

echo "== 12. 2 回目も通り、何も増えない"
run
[ "$rc" = 0 ] || ng "2 回目 exit=$rc"
grep -q "取り込みを飛ばします" "$ERR" || ng "apprepo が 2 回目も取り込もうとしています"
[ "$(git -C "$REPO" status --porcelain -- conf)" = "" ] || ng "2 回目の conf で差分が出ています"
for h in isu1 isu3; do
  [ "$(count_on "$h" '# >>> isuh >>>' /etc/hosts)" = 1 ] || ng "$h の /etc/hosts のブロックが 1 つではありません"
  [ "$(count_on "$h" '# >>> isuh >>>' '~/.bashrc')" = 1 ] || ng "$h の ~/.bashrc のブロックが 1 つではありません"
done
[ "$(count_on isu1 "$(cat "$ROOT/.isuh/setup/id_ed25519.pub")" '~/.ssh/authorized_keys')" = 1 ] ||
  ng "authorized_keys に鍵が重複しています"
ok "2 回目も exit 0 で、管理ブロックと鍵は 1 つのまま"

echo "== 13. 手元で編集した conf は、取り込み直しで潰さない"
echo "# edited locally" >>"$C/isu1/etc/nginx/nginx.conf"
run --only conf
[ "$rc" = 0 ] || ng "exit=$rc"
tail -1 "$C/isu1/etc/nginx/nginx.conf" | grep -qx "# edited locally" || ng "手元の編集が消えました"
grep -q "上書きしない" "$ERR" || ng "上書きしなかったことが出ていません"
ok "既にあるファイルは上書きしない"

echo "== 14. 手元の ~/.ssh の書き換え（Include の挿入・退避・冪等）"
set +e
cases=$(node --import tsx setup-probe.ts 2>"$ERR")
rc=$?
set -e
[ "$rc" = 0 ] || ng "setup-probe が失敗しました"
ok "管理ブロックと Include の書き換え（${cases}）"

echo
echo "全 $passed 件パス"
