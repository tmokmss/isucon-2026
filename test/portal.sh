#!/bin/sh
# ポータル操作(#3) の検証。コンテナは不要（fake-portal をローカルで立てる）。
set -eu
cd "$(dirname "$0")"
TEST_DIR=$(pwd)
ROOT=$(cd .. && pwd)

export ISUH_CONFIG="$TEST_DIR/config.test.ts"
ISUH="$ROOT/bin/isuh"
WORK="$TEST_DIR/.work"
COOKIE="$WORK/portal-cookie"
ERR="$WORK/portal.err"
mkdir -p "$WORK"

passed=0
ng() {
  printf '  NG  %s\n' "$*" >&2
  echo "--- 直前の stderr ---" >&2
  tail -n 20 "$ERR" >&2 || true
  exit 1
}
ok() {
  passed=$((passed + 1))
  printf '  ok  %s\n' "$*"
}
run() {
  set +e
  out=$("$ISUH" "$@" 2>"$ERR")
  rc=$?
  set -e
}

# fake-portal を起動
node --import tsx fake-portal.ts >"$WORK/fake-portal.log" 2>&1 &
portal_pid=$!
trap 'kill "$portal_pid" 2>/dev/null || true' EXIT

i=0
until curl -s -o /dev/null "http://127.0.0.1:18080/auth/login/" 2>/dev/null; do
  i=$((i + 1))
  [ "$i" -gt 50 ] && ng "fake-portal が起動しません"
  sleep 0.1
done

echo "== 1. Cookie ファイルが無ければ exit 65 と手順を出す"
rm -f "$COOKIE"
run portal whoami
[ "$rc" = 65 ] || ng "exit=$rc (65 のはず)"
grep -q "DevTools" "$ERR" || ng "Cookie の取り出し手順が出ていません"
ok "Cookie が無いときに取り方を教えて 65 で落ちる"

echo "== 2. Cookie が切れていれば exit 65"
printf 'sessionid=expired; csrftoken=tok123\n' >"$COOKIE"
run portal whoami
[ "$rc" = 65 ] || ng "exit=$rc (65 のはず)"
grep -q "Cookie が効いていません" "$ERR" || ng "ログイン切れと言っていません"
grep -q "/auth/login" "$ERR" || ng "リダイレクト先が出ていません"
ok "302 → /auth/login をログイン切れとして見分ける"

echo "== 3. 有効な Cookie なら whoami が通る"
printf '# ブラウザから持ち出した Cookie\nsessionid=valid-session; csrftoken=tok123\n' >"$COOKIE"
run portal whoami
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ "$out" = "チーム情報 | ISUCON" ] || ng "stdout='$out'"
ok "認証必須ページを Cookie 付きで引ける（#3 の完了条件）"

echo "== 3b. 200 を返すトップページでもログイン切れを見分ける"
# 本物のポータルの / は未ログインでも 200 を返す。ステータスだけでは判定できない。
TOPCFG="$WORK/config.top.ts"
sed -e 's|whoamiPath: "/team/"|whoamiPath: "/"|' -e 's|join(here, ".work", "portal-cookie")|join(here, "portal-cookie")|' \
  "$TEST_DIR/config.test.ts" >"$TOPCFG"
printf 'sessionid=expired; csrftoken=tok123\n' >"$WORK/portal-cookie"
set +e
out=$(ISUH_CONFIG="$TOPCFG" "$ISUH" portal whoami 2>"$ERR")
rc=$?
set -e
[ "$rc" = 65 ] || ng "exit=$rc (65 のはず) — 200 のトップを有効と誤判定している"
printf '# ブラウザから持ち出した Cookie\nsessionid=valid-session; csrftoken=tok123\n' >"$WORK/portal-cookie"
set +e
out=$(ISUH_CONFIG="$TOPCFG" "$ISUH" portal whoami 2>"$ERR")
rc=$?
set -e
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ "$out" = "ISUCON2026 Portal" ] || ng "stdout='$out'"
ok "ナビのログイン/ログアウトリンクで見分ける（実物のポータルで確認済みの挙動）"

echo "== 4. get は任意のページを引ける"
run portal get /team/
[ "$rc" = 0 ] || ng "exit=$rc"
case "$out" in
  *"チーム情報"*) ok "get で生のページが取れる" ;;
  *) ng "stdout='$out'" ;;
esac

echo "== 5. POST は CSRF が通る"
run portal post /api/bench/enqueue team=7
[ "$rc" = 0 ] || ng "exit=$rc (CSRF で弾かれた可能性): $(tail -2 "$ERR")"
case "$out" in
  *'"job_id":"4321"'*) ok "X-CSRFToken が受理される" ;;
  *) ng "stdout='$out'" ;;
esac

echo "== 6. csrftoken が無ければ POST は 403 になる（警告が出る）"
printf 'sessionid=valid-session\n' >"$COOKIE"
run portal post /api/bench/enqueue
[ "$rc" = 65 ] || ng "exit=$rc (65 のはず)"
grep -q "csrftoken がありません" "$ERR" || ng "csrftoken が無いことを警告していません"
ok "csrftoken 欠けを警告し、403 をログイン切れとして扱う"

printf 'sessionid=valid-session; csrftoken=tok123\n' >"$COOKIE"

echo "== 7. enqueue が job-id を stdout に出す"
run portal enqueue
[ "$rc" = 0 ] || ng "exit=$rc"
[ "$out" = "4321" ] || ng "stdout='$out' (4321 のはず)"
ok "JOB=\$(isuh portal enqueue) が使える"

echo "== 8. status と wait が状態を追える"
run portal status 4321
[ "$rc" = 0 ] || ng "exit=$rc"
[ "$out" = "pending" ] || ng "stdout='$out' (pending のはず)"
run portal wait 4321
[ "$rc" = 0 ] || ng "exit=$rc"
case "$out" in
  *'"status":"done"'*) ok "wait が完了まで追って結果 JSON を出す" ;;
  *) ng "stdout='$out'" ;;
esac

echo "== 9. score が取れる"
run portal score
[ "$rc" = 0 ] || ng "exit=$rc"
[ "$out" = "12345" ] || ng "stdout='$out' (12345 のはず)"
ok "score が確定スコアだけを出す"

echo "== 10. 未設定のエンドポイントは 65 で「当日埋める」と言う"
ISUH_CONFIG="$TEST_DIR/.work/config.noendpoint.ts"
sed 's|scorePath: "/api/score"|scorePath: ""|' "$TEST_DIR/config.test.ts" >"$ISUH_CONFIG"
export ISUH_CONFIG
run portal score
[ "$rc" = 65 ] || ng "exit=$rc (65 のはず)"
grep -q "当日" "$ERR" || ng "当日埋めるという案内が出ていません"
export ISUH_CONFIG="$TEST_DIR/config.test.ts"
ok "未設定のパスは設定不備として落ちる"

echo "== 11. 引数不正は 64"
run portal
[ "$rc" = 64 ] || ng "exit=$rc"
run portal nosuchsub
[ "$rc" = 64 ] || ng "exit=$rc"
run portal status
[ "$rc" = 64 ] || ng "exit=$rc"
run portal get
[ "$rc" = 64 ] || ng "exit=$rc"
ok "サブコマンド・引数の不正は 64"

echo
echo "全 $passed 件パス"
