#!/bin/sh
# deploy(#1) の通し検証。先に ./test/up.sh を実行しておくこと。
set -eu
cd "$(dirname "$0")"
TEST_DIR=$(pwd)
ROOT=$(cd .. && pwd)

export ISUH_CONFIG="$TEST_DIR/config.test.ts"
export ISUH_SSH_CONFIG="$TEST_DIR/.work/ssh_config"
ISUH="$ROOT/bin/isuh"
APP="$TEST_DIR/.work/app-repo"
ERR="$TEST_DIR/.work/last.err"
SSH="ssh -F $ISUH_SSH_CONFIG -o ConnectTimeout=5"

if [ ! -d "$APP/.git" ]; then
  echo "先に ./test/up.sh を実行してください" >&2
  exit 1
fi

passed=0
ng() {
  printf '  NG  %s\n' "$*" >&2
  echo "--- 直前の stderr ---" >&2
  tail -n 30 "$ERR" >&2 || true
  exit 1
}
ok() {
  passed=$((passed + 1))
  printf '  ok  %s\n' "$*"
}

# isuh を走らせて終了コードを rc、stdout を out に入れる
run() {
  set +e
  out=$("$ISUH" "$@" 2>"$ERR")
  rc=$?
  set -e
}

version_on() { $SSH "$1" 'curl -sS --max-time 3 http://127.0.0.1:8080/version'; }
deployed_on() { $SSH "$1" 'cut -f1 "$HOME/.deployed_sha"'; }
active_on() { $SSH "$1" "systemctl is-active $2" 2>/dev/null || true; }

V2=$(git -C "$APP" rev-parse main)
V1=$(git -C "$APP" rev-parse main~1)
BROKEN=$(git -C "$APP" rev-parse broken)

# conf/ の検証はアプリと同じリポジトリの別ブランチで行い、終わったら main に戻す
cleanup() {
  git -C "$APP" checkout -q -f main 2>/dev/null || true
  git -C "$APP" branch -q -D with-conf 2>/dev/null || true
  git -C "$APP" branch -q -D with-sql 2>/dev/null || true
  git -C "$APP" checkout -q -- harness.txt 2>/dev/null || true
}
cleanup
trap cleanup EXIT

echo "== 1. deploy main (v2)"
run deploy main
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ "$out" = "$V2" ] || ng "stdout='$out' ($V2 のはず)"
for h in isu1 isu2; do
  [ "$(version_on $h)" = "v2" ] || ng "$h が v2 になっていない"
  [ "$(deployed_on $h)" = "$V2" ] || ng "$h の .deployed_sha が違う"
  [ "$(active_on $h isuapp.go.service)" = "active" ] || ng "$h でサービスが動いていない"
done
ok "2 台に v2 が載り、sha が記録され、サービスが上がっている"

echo "== 2. deploy <古い sha> (v1)"
run deploy "$V1"
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ "$out" = "$V1" ] || ng "stdout='$out' ($V1 のはず)"
for h in isu1 isu2; do
  [ "$(version_on $h)" = "v1" ] || ng "$h が v1 に戻っていない"
  [ "$(deployed_on $h)" = "$V1" ] || ng "$h の .deployed_sha が v1 でない"
done
ok "古い sha へ戻せる（前の版へ戻すのも deploy でやる）"

echo "== 3. --dry-run は何も変えない"
run deploy main --dry-run
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ "$out" = "$V2" ] || ng "stdout='$out' ($V2 のはず)"
[ "$(version_on isu1)" = "v1" ] || ng "dry-run なのに isu1 が変わった"
ok "dry-run は sha を出すだけで配らない"

echo "== 4. deploy main をもう一度（ビルドキャッシュ）"
run deploy main
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
grep -q "キャッシュを使用" "$ERR" || ng "ビルドキャッシュが効いていない"
[ "$(version_on isu1)" = "v2" ] || ng "isu1 が v2 に戻っていない"
ok "同じ sha の再デプロイはビルドを省く"

echo "== 5. ビルドが壊れた ref は exit 20 で止まる"
run deploy "$BROKEN"
[ "$rc" = 20 ] || ng "exit=$rc (20 のはず)"
for h in isu1 isu2; do
  [ "$(version_on $h)" = "v2" ] || ng "$h が壊れた版に書き換わっている"
done
ok "ビルド失敗時はサーバに触らない（後続のベンチを走らせない）"

echo "== 6. 存在しない ref は exit 64"
run deploy no-such-ref-xyz
[ "$rc" = 64 ] || ng "exit=$rc (64 のはず)"
ok "引数不正は 64"

echo "== 7. conf/ はバイナリと同じ sha から配られる"
git -C "$APP" checkout -q -b with-conf main
CONF="$APP/conf"
mkdir -p "$CONF/_all/etc/isuh-smoke" "$CONF/isu2/etc/isuh-smoke" "$CONF/isu1/etc/nginx" "$CONF/isu3/etc/isuh-smoke"
echo from-all >"$CONF/_all/etc/isuh-smoke/common.conf"
echo from-isu2 >"$CONF/isu2/etc/isuh-smoke/common.conf"
echo "# isuh smoke" >"$CONF/isu1/etc/nginx/isuh-smoke.conf"
echo from-isu3 >"$CONF/isu3/etc/isuh-smoke/db.conf"
git -C "$APP" add conf
git -C "$APP" commit -qm "conf"
WITH_CONF=$(git -C "$APP" rev-parse HEAD)
# コミットしていない変更は配られない（配るのは sha の中身）
echo uncommitted >"$CONF/_all/etc/isuh-smoke/common.conf"
$SSH isu3 'sudo rm -f /etc/isuh-smoke/db.conf'

run deploy with-conf
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ "$out" = "$WITH_CONF" ] || ng "stdout='$out'"
grep -q "未コミットの変更が 1 件" "$ERR" || ng "conf/ の未コミットの変更を警告していません"
[ "$($SSH isu1 'cat /etc/isuh-smoke/common.conf')" = "from-all" ] || ng "isu1 に sha の中身ではないものが届いた"
[ "$($SSH isu2 'cat /etc/isuh-smoke/common.conf')" = "from-isu2" ] || ng "isu2 でホスト指定が優先されていない"
[ "$($SSH isu1 'test -f /etc/nginx/isuh-smoke.conf && echo yes')" = "yes" ] || ng "nginx 設定が届いていない"
[ "$($SSH isu2 'test -f /etc/nginx/isuh-smoke.conf && echo yes || echo no')" = "no" ] || ng "isu2 に余計な配布がされている"
[ "$(active_on isu1 nginx.service)" = "active" ] || ng "nginx が reload されていない"
ok "conf/_all と conf/<host> が sha の中身で配られ、nginx が reload される"

echo "== 8. アプリの無いホスト（DB）にも conf/ が配られる"
[ "$($SSH isu3 'cat /etc/isuh-smoke/db.conf')" = "from-isu3" ] || ng "isu3 に conf/isu3 が届いていない"
[ "$($SSH isu3 'test -f /home/isucon/webapp/go/isuapp.isuh-new && echo yes || echo no')" = "no" ] || ng "isu3 にバイナリを配ろうとした"
grep -q "isu3: restart isuapp" "$ERR" && ng "isu3 でアプリを再起動しようとした"
ok "conf/isu3 は isu3 へ。バイナリと再起動はアプリのホストだけ"

echo "== 9. ハーネス側だけの変更は警告しない"
git -C "$APP" checkout -q -f main
echo changed >>"$APP/harness.txt"
run deploy --dry-run main
[ "$rc" = 0 ] || ng "exit=$rc"
grep -q "未コミット" "$ERR" && ng "配らないファイルの変更で警告しています"
ok "警告の対象はアプリと conf/ だけ"

echo "== 10. バイナリ以外の webapp/ も同じ sha の中身で app.dir へ届く"
git -C "$APP" checkout -q -f main
git -C "$APP" checkout -q -- harness.txt
git -C "$APP" checkout -q -b with-sql main
mkdir -p "$APP/webapp/sql"
echo "-- schema v1" >"$APP/webapp/sql/1-schema.sql"
printf '#!/bin/sh\necho init-ok\n' >"$APP/webapp/sql/init.sh"
chmod +x "$APP/webapp/sql/init.sh"
git -C "$APP" add webapp/sql
git -C "$APP" commit -qm "sql v1"
SQL1=$(git -C "$APP" rev-parse HEAD)
echo "-- schema v2 (index)" >"$APP/webapp/sql/1-schema.sql"
git -C "$APP" commit -qam "sql v2"
SQL2=$(git -C "$APP" rev-parse HEAD)
# サーバにしか無いもの（gitignore されたビルド成果物・node_modules 等）の代わり
for h in isu1 isu2; do
  $SSH "$h" 'rm -rf ~/webapp/sql && mkdir -p ~/webapp/node_modules && echo keep >~/webapp/node_modules/x'
  # 手元より新しい mtime を持つ古いファイル。mtime の比較で取りこぼさないことを見る
  $SSH "$h" 'mkdir -p ~/webapp/sql && echo stale >~/webapp/sql/1-schema.sql && touch -d 2099-01-01 ~/webapp/sql/1-schema.sql'
done

run deploy --dry-run with-sql
[ "$rc" = 0 ] || ng "exit=$rc"
grep -q "isu1: webapp/ を /home/isucon/webapp へ同期" "$ERR" || ng "dry-run に app.dir への同期が出ていない"
grep -q "isu3: webapp/" "$ERR" && ng "dry-run でアプリの無いホストへ同期しようとしている"
[ "$($SSH isu1 'cat ~/webapp/sql/1-schema.sql')" = "stale" ] || ng "dry-run なのに同期された"

run deploy with-sql
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ "$out" = "$SQL2" ] || ng "stdout='$out'"
for h in isu1 isu2; do
  [ "$($SSH $h 'cat ~/webapp/sql/1-schema.sql')" = "-- schema v2 (index)" ] || ng "$h に sha の schema が届いていない"
  [ "$($SSH $h '~/webapp/sql/init.sh')" = "init-ok" ] || ng "$h の init.sh が実行できない（実行ビット）"
  [ "$($SSH $h 'cat ~/webapp/node_modules/x')" = "keep" ] || ng "$h のサーバにしか無いファイルが消えた"
  [ "$(version_on $h)" = "v2" ] || ng "$h のバイナリが動いていない"
  [ "$($SSH $h 'test -e ~/webapp/go/isuapp.isuh-new && echo yes || echo no')" = "no" ] || ng "$h に途中のバイナリが残っている"
done
[ "$($SSH isu3 'test -e ~/webapp/sql && echo yes || echo no')" = "no" ] || ng "isu3 に webapp/ を配った"

run deploy "$SQL1"
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ "$($SSH isu1 'cat ~/webapp/sql/1-schema.sql')" = "-- schema v1" ] || ng "古い sha で schema が戻らない"
for h in isu1 isu2; do $SSH "$h" 'rm -rf ~/webapp/sql ~/webapp/node_modules'; done
ok "sql/ が sha の中身で届き（mtime によらず上書き・実行ビット維持）、サーバ側だけのファイルは残る"

echo "== 11. 変わっていなければ配らない・reload / restart しない"
git -C "$APP" checkout -q -f with-conf
run deploy with-conf # 状態をそろえる
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
run deploy with-conf
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
for h in isu1 isu2 isu3; do
  grep -q "$h: 設定ファイル [0-9]* 件はすべて変更なし" "$ERR" || ng "$h で変わっていない conf/ を配った"
done
grep -qE ": (reload|restart) " "$ERR" && ng "変わっていないのに reload / restart した"
grep -q "sysctl --system" "$ERR" && ng "変わっていないのに sysctl を流した"
for h in isu1 isu2; do
  grep -q "$h: webapp/ は前回の同期から変更なし" "$ERR" || ng "$h で変わっていない webapp/ を同期した"
  grep -q "$h: バイナリは変更なし" "$ERR" || ng "$h で変わっていないバイナリを送った"
  [ "$(version_on $h)" = "v2" ] || ng "$h のアプリが再起動後に動いていない"
done
# 別の deploy（記録を書かない古いハーネスや別の worktree）が入ったら、同じツリーでも同期し直す
$SSH isu1 'printf "other\tHEAD\tx\tsomeone\n" >~/.deployed_sha'
run deploy with-conf
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
grep -q "isu1: webapp/ は前回の同期から変更なし" "$ERR" && ng "別の deploy が入った isu1 で同期を省いた"
grep -q "isu2: webapp/ は前回の同期から変更なし" "$ERR" || ng "isu2 で変わっていない webapp/ を同期した"
ok "同じ中身の再デプロイは、アプリの再起動だけ（別の deploy が入ったホストは同期し直す）"

echo "== 12. 変えた設定ファイルだけを配り、対応するサービスだけを反映する"
echo "from-all v2" >"$APP/conf/_all/etc/isuh-smoke/common.conf"
git -C "$APP" commit -qam "conf: common v2"
run deploy with-conf
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
grep -q "isu1: 設定ファイル 1/[0-9]* 件を配布: /etc/isuh-smoke/common.conf" "$ERR" || ng "isu1 に変えたファイルだけが配られていない"
grep -q "isu3: 設定ファイル 1/[0-9]* 件を配布: /etc/isuh-smoke/common.conf" "$ERR" || ng "isu3 に変えたファイルだけが配られていない"
grep -q "isu2: 設定ファイル [0-9]* 件はすべて変更なし" "$ERR" || ng "isu2（ホスト指定で上書き）に _all の変更を配った"
grep -qE ": (reload|restart) " "$ERR" && ng "nginx / MySQL の設定を変えていないのに reload / restart した"
[ "$($SSH isu1 'cat /etc/isuh-smoke/common.conf')" = "from-all v2" ] || ng "isu1 に新しい中身が届いていない"
[ "$($SSH isu2 'cat /etc/isuh-smoke/common.conf')" = "from-isu2" ] || ng "isu2 のホスト指定が崩れた"
grep -q "isu1: webapp/ は前回の同期から変更なし" "$ERR" || ng "conf/ だけの変更で webapp/ を同期した"

echo "# isuh smoke v2" >"$APP/conf/isu1/etc/nginx/isuh-smoke.conf"
git -C "$APP" commit -qam "conf: nginx v2"
run deploy with-conf
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
grep -q "isu1: reload nginx.service" "$ERR" || ng "nginx の設定を変えたのに reload していない"
grep -qE "isu[23]: (reload|restart) " "$ERR" && ng "nginx の設定を変えていないホストで reload / restart した"
[ "$(active_on isu1 nginx.service)" = "active" ] || ng "reload 後に nginx が動いていない"
ok "変えたファイルだけが配られ、nginx は nginx の設定を変えたときだけ reload される"

echo
echo "全 $passed 件パス"
