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
rm -f "$WORK/fake-score" "$WORK/fake-status"

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
until curl -s -o /dev/null "http://127.0.0.1:18080/api/session" 2>/dev/null; do
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

echo "== 2. 200 を返す /api/session でもログイン切れを見分ける"
printf 'isuxportal_session=expired\n' >"$COOKIE"
run portal whoami
[ "$rc" = 65 ] || ng "exit=$rc (65 のはず) — contestant が空のセッションを有効と誤判定している"
grep -q "Cookie が効いていません" "$ERR" || ng "ログイン切れと言っていません"
ok "contestant の無いセッションをログイン切れとして扱う（実物のポータルで確認済みの挙動）"

echo "== 3. 認証が要る API の 302 → /session/new をログイン切れとして見分ける"
run portal score
[ "$rc" = 65 ] || ng "exit=$rc (65 のはず)"
grep -q "/session/new/github" "$ERR" || ng "リダイレクト先が出ていません"
ok "302 → /session/new/github を 65 にする"

echo "== 4. 有効な Cookie なら whoami がチーム名と参加者名を出す"
printf '# ブラウザから持ち出した Cookie\nisuxportal_session=valid-session\n' >"$COOKIE"
run portal whoami
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ "$out" = "isuh-test / tester" ] || ng "stdout='$out'"
grep -q "instance #1 id=101 public=192.0.2.1" "$ERR" || ng "インスタンス一覧が出ていません"
ok "セッションを protobuf で読める"

echo "== 5. get は protobuf をフィールド番号で展開して出す"
run portal get /api/session
[ "$rc" = 0 ] || ng "exit=$rc"
case "$out" in
  *'"2": "isuh-test"'*) ok "get で API の中身が読める" ;;
  *) ng "stdout='$out'" ;;
esac

echo "== 6. POST はトップの csrf-token を X-Csrf-Token に載せる"
run portal post /api/contestant/benchmark_jobs
# 本体が空なので target_id 違いの 400 になる。CSRF で弾かれたなら 422
[ "$rc" = 0 ] || ng "exit=$rc"
grep -q "HTTP 400" "$ERR" || ng "CSRF を通っていません: $(grep HTTP "$ERR")"
ok "CSRF を通る"

echo "== 7. enqueue が job-id を stdout に出す（対象はインスタンス一覧の先頭）"
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
  *'"status":"done"'*'"score":12345,"raw":12445,"deduction":100'*) ok "wait が完了まで追って減点の内訳つきの結果 JSON を出す" ;;
  *) ng "stdout='$out'" ;;
esac

echo "== 9. 不合格のジョブは failed になる"
echo fail >"$WORK/fake-status"
run portal enqueue
run portal wait 4321
rm -f "$WORK/fake-status"
[ "$rc" = 0 ] || ng "exit=$rc"
case "$out" in
  *'"status":"failed"'*'"portalStatus":"finished"'*'整合性チェック'*) ok "passed=false を failed として理由つきで出す" ;;
  *) ng "stdout='$out'" ;;
esac

echo "== 10. score が取れる"
run portal score
[ "$rc" = 0 ] || ng "exit=$rc"
[ "$out" = "12345" ] || ng "stdout='$out' (12345 のはず)"
ok "score が確定スコアだけを出す"

echo "== 11. 未設定のエンドポイントは 65"
ISUH_CONFIG="$TEST_DIR/.work/config.noendpoint.ts"
sed 's|scorePath: "/api/contestant/dashboard"|scorePath: ""|' "$TEST_DIR/config.test.ts" >"$ISUH_CONFIG"
export ISUH_CONFIG
run portal score
[ "$rc" = 65 ] || ng "exit=$rc (65 のはず)"
grep -q "scorePath が空" "$ERR" || ng "どの設定が空かを言っていません"
export ISUH_CONFIG="$TEST_DIR/config.test.ts"
ok "未設定のパスは設定不備として落ちる"

echo "== 12. 引数不正は 64"
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
