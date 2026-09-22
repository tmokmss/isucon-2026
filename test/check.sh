#!/bin/sh
# 検証ゲート(#12) の検証。先に ./test/up.sh を実行しておくこと。
set -eu
cd "$(dirname "$0")"
TEST_DIR=$(pwd)
ROOT=$(cd .. && pwd)
WORK="$TEST_DIR/.work"

export ISUH_CONFIG="$TEST_DIR/config.test.ts"
export ISUH_SSH_CONFIG="$WORK/ssh_config"
export ISUH_RUNS_DIR="$WORK/runs-check"
ISUH="$ROOT/bin/isuh"
ERR="$WORK/check.err"
SNAPSHOT="$WORK/structure.json"

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
  out=$("$ISUH" "$@" 2>"$ERR")
  rc=$?
  set -e
}

rm -f "$SNAPSHOT"
rm -rf "$ISUH_RUNS_DIR"

echo "== 0. 正常な版を載せる"
run deploy main
[ "$rc" = 0 ] || ng "deploy に失敗: exit=$rc"
ok "main (正しい構造) をデプロイ"

echo "== 1. 初回は今の構造を記録する"
run check
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ -f "$SNAPSHOT" ] || ng "スナップショットが作られていません"
grep -q "構造を新しく記録しました" "$ERR" || ng "記録したことが知らされていません"
grep -q "コミットしてください" "$ERR" || ng "コミットの案内がありません"
ok "ベースラインを記録し、それが基準になることを警告する"

echo "== 2. 記録した形には name / id / tags が入っている"
node -e '
const s = require(process.argv[1]);
const e = s.entries["GET /api/items"];
if (!e) { console.error("エントリが無い: " + Object.keys(s.entries)); process.exit(1); }
const items = e.shape.fields.items;
if (items.t !== "array") { console.error("items が配列でない"); process.exit(1); }
const keys = Object.keys(items.of[0].fields).sort().join(",");
if (keys !== "id,name,tags") { console.error("キーが違う: " + keys); process.exit(1); }
if (e.shape.fields.total.t !== "number") { console.error("total が number でない"); process.exit(1); }
' "$SNAPSHOT" || ng "記録された形が違います"
ok "値ではなくキーと型の形だけが記録される"

echo "== 3. 同じ版なら何度走らせても通る"
run check
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
grep -q "check ok" "$ERR" || ng "成功のログがありません"
grep -q "構造を新しく記録しました" "$ERR" && ng "2 回目なのに記録し直しています"
ok "記録済みなら比較だけして通る"

echo "== 4. 形の比較そのものの挙動"
set +e
cases=$(node --import tsx shape-probe.ts 2>"$ERR")
rc=$?
set -e
[ "$rc" = 0 ] || ng "構造比較の挙動が期待と違います"
ok "null の揺れ・要素数・キー順は許し、キーや型の変更は捕まえる（${cases}）"

echo "== 5. キーをリネームしたら exit 30 で止まる"
run deploy badshape
[ "$rc" = 0 ] || ng "badshape のデプロイに失敗: exit=$rc"
run check
[ "$rc" = 30 ] || ng "exit=$rc (30 のはず)"
grep -q "キーが消えた name" "$ERR" || ng "消えたキーを指摘していません"
grep -q "キーが増えた title" "$ERR" || ng "増えたキーを指摘していません"
grep -q "レギュレーション違反" "$ERR" || ng "違反であることが書かれていません"
ok "レスポンス構造の変更を name → title のレベルで捕まえる"

echo "== 6. bench は check で落ちたらベンチを起動しない"
printf 'sessionid=valid-session; csrftoken=tok123\n' >"$WORK/portal-cookie"
node --import tsx fake-portal.ts >"$WORK/fake-portal-check.log" 2>&1 &
portal_pid=$!
trap 'kill "$portal_pid" 2>/dev/null || true' EXIT
i=0
until curl -s -o /dev/null "http://127.0.0.1:18080/auth/login/" 2>/dev/null; do
  i=$((i + 1))
  [ "$i" -gt 50 ] && ng "fake-portal が起動しません"
  sleep 0.1
done

run bench badshape
[ "$rc" = 30 ] || ng "exit=$rc (30 のはず)"
grep -q "キーが消えた name" "$ERR" || ng "check の理由が出ていません"
grep -q "ベンチを登録しました" "$ERR" && ng "ベンチを起動してしまっています"
grep -q "ロックを解放しました" "$ERR" || ng "ロックを解放していません"
ok "ベンチ枠を消費せずに止まり、ロックは解放される"

echo "== 7. --skip-check なら通る"
run bench badshape --skip-check
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
grep -q "check を省略しました" "$ERR" || ng "省略の警告がありません"
ok "--skip-check で意図的に飛ばせる（警告つき）"

echo "== 8. 引数は取らない"
run check extra
[ "$rc" = 64 ] || ng "exit=$rc (64 のはず)"
ok "引数不正は 64"

echo "== 9. アプリが落ちていれば 30"
ssh -F "$ISUH_SSH_CONFIG" isu1 "sudo systemctl stop isuapp.go.service" >/dev/null 2>&1
run check
[ "$rc" = 30 ] || ng "exit=$rc (30 のはず)"
grep -qE "initialize を叩けませんでした|リクエストに失敗" "$ERR" || ng "接続できない理由が出ていません"
ok "アプリが応答しなければ検証失敗として止まる"

# 後片付け: 正しい版に戻す
run deploy main
echo
echo "全 $passed 件パス"
