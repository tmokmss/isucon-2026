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

# conf/ は #10 が育てる場所。既にあるなら壊さないようにテストを飛ばす。
CONF="$ROOT/conf"
CONF_TEST=1
[ -e "$CONF" ] && CONF_TEST=0
cleanup() { [ "$CONF_TEST" = 1 ] && rm -rf "$CONF"; }
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

if [ "$CONF_TEST" = 1 ]; then
  echo "== 7. conf/ の配布とサービス再読込"
  mkdir -p "$CONF/_all/etc/isuh-smoke" "$CONF/isu2/etc/isuh-smoke" "$CONF/isu1/etc/nginx"
  echo from-all >"$CONF/_all/etc/isuh-smoke/common.conf"
  echo from-isu2 >"$CONF/isu2/etc/isuh-smoke/common.conf"
  echo "# isuh smoke" >"$CONF/isu1/etc/nginx/isuh-smoke.conf"

  run deploy main
  [ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
  [ "$($SSH isu1 'cat /etc/isuh-smoke/common.conf')" = "from-all" ] || ng "isu1 に _all が届いていない"
  [ "$($SSH isu2 'cat /etc/isuh-smoke/common.conf')" = "from-isu2" ] || ng "isu2 でホスト指定が優先されていない"
  [ "$($SSH isu1 'test -f /etc/nginx/isuh-smoke.conf && echo yes')" = "yes" ] || ng "nginx 設定が届いていない"
  [ "$($SSH isu2 'test -f /etc/nginx/isuh-smoke.conf && echo yes || echo no')" = "no" ] || ng "isu2 に余計な配布がされている"
  [ "$(active_on isu1 nginx.service)" = "active" ] || ng "nginx が reload されていない"
  ok "conf/_all と conf/<host> が配られ、nginx が reload される"
else
  echo "== 7. conf/ が既にあるため設定配布のテストは飛ばします"
fi

echo
echo "全 $passed 件パス"
