#!/bin/sh
# 自動ロールバック(#6) の検証。先に ./test/up.sh を実行しておくこと。
set -eu
cd "$(dirname "$0")"
TEST_DIR=$(pwd)
ROOT=$(cd .. && pwd)
WORK="$TEST_DIR/.work"

export ISUH_CONFIG="$TEST_DIR/config.test.ts"
export ISUH_SSH_CONFIG="$WORK/ssh_config"
export ISUH_RUNS_DIR="$WORK/runs-rollback"
ISUH="$ROOT/bin/isuh"
ERR="$WORK/rollback.err"
APP="$WORK/app-repo"
SSH="ssh -F $ISUH_SSH_CONFIG -o ConnectTimeout=5"

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
version_on() { $SSH "$1" 'curl -sS --max-time 3 http://127.0.0.1:8080/version'; }

V2=$(git -C "$APP" rev-parse main)
V1=$(git -C "$APP" rev-parse main~1)

rm -rf "$ISUH_RUNS_DIR"
rm -f "$WORK/structure.json" "$WORK/fake-score"
printf 'sessionid=valid-session; csrftoken=tok123\n' >"$WORK/portal-cookie"

# autoRollback を有効にした設定
AUTOCFG="$WORK/config.autorollback.ts"
sed -e 's/autoRollback: false/autoRollback: true/' \
  -e 's|join(here, ".work", |join(here, |g' \
  "$TEST_DIR/config.test.ts" >"$AUTOCFG"

node --import tsx fake-portal.ts >"$WORK/fake-portal-rb.log" 2>&1 &
portal_pid=$!
trap 'kill "$portal_pid" 2>/dev/null || true' EXIT
i=0
until curl -s -o /dev/null "http://127.0.0.1:18080/auth/login/" 2>/dev/null; do
  i=$((i + 1))
  [ "$i" -gt 50 ] && ng "fake-portal が起動しません"
  sleep 0.1
done

echo "== 1. 履歴が無ければ戻し先を決められない"
run rollback
[ "$rc" = 64 ] || ng "exit=$rc (64 のはず)"
grep -q "isuh rollback <sha>" "$ERR" || ng "明示する方法が案内されていません"
ok "pass したランが無ければ 64 で、sha を明示しろと言う"

echo "== 2. good revision を 1 つ作る"
run bench main --note "good"
[ "$rc" = 0 ] || ng "exit=$rc"
[ "$out" = "12345" ] || ng "stdout='$out'"
[ "$(version_on isu1)" = "v2" ] || ng "isu1 が v2 でない"
ok "main (v2, score=12345) が台帳に pass として入った"

echo "== 3. 別の revision を載せてから rollback で戻る"
run deploy "$V1"
[ "$rc" = 0 ] || ng "deploy に失敗"
[ "$(version_on isu1)" = "v1" ] || ng "isu1 が v1 でない"
run rollback
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ "$out" = "$V2" ] || ng "stdout='$out' ($V2 のはず)"
[ "$(version_on isu1)" = "v2" ] || ng "isu1 が戻っていない"
[ "$(version_on isu2)" = "v2" ] || ng "isu2 が戻っていない"
grep -q "直前の good revision" "$ERR" || ng "どう決めたかが出ていません"
ok "引数なしで台帳の直前の good へ全台戻る"

echo "== 4. sha を明示しても戻せる"
run rollback "$V1"
[ "$rc" = 0 ] || ng "exit=$rc"
[ "$out" = "$V1" ] || ng "stdout='$out'"
[ "$(version_on isu1)" = "v1" ] || ng "isu1 が v1 でない"
grep -q "指定" "$ERR" || ng "指定によるものだと出ていません"
ok "isuh rollback <sha> で任意の revision へ戻せる"

echo "== 5. ベンチ中は戻せない（ロックを取る）"
node --import tsx lock-probe.ts HOLD -1 >/dev/null 2>&1 &
pid_hold=$!
i=0
while $SSH isu1 "flock -n /tmp/isucon-bench.lock -c true" >/dev/null 2>&1; do
  i=$((i + 1))
  [ "$i" -gt 100 ] && ng "HOLD がロックを取れていません"
  sleep 0.1
done
run rollback
[ "$rc" = 10 ] || ng "exit=$rc (10 のはず)"
grep -q "保持者:" "$ERR" || ng "保持者が出ていません"
kill -9 "$pid_hold" 2>/dev/null || true
wait "$pid_hold" 2>/dev/null || true
i=0
until $SSH isu1 "flock -n /tmp/isucon-bench.lock -c true" >/dev/null 2>&1; do
  i=$((i + 1))
  [ "$i" -gt 100 ] && ng "ロックが解放されません"
  sleep 0.1
done
ok "ベンチと同じロックを取るので計測を壊さない"

echo "== 6. autoRollback=false ならスコアが落ちても戻さない"
echo 1000 >"$WORK/fake-score"
run bench "$V1" --note "regression, no auto"
[ "$rc" = 50 ] || ng "exit=$rc (50 のはず)"
[ "$out" = "1000" ] || ng "stdout='$out' (低下してもスコアは出す)"
grep -q "autoRollback が false" "$ERR" || ng "戻さなかった理由が出ていません"
grep -q "戻すなら: isuh rollback" "$ERR" || ng "戻し方が案内されていません"
[ "$(version_on isu1)" = "v1" ] || ng "戻していないはずなのに変わっている"
ok "序盤の運用（手動）: 止めるが戻さない"

echo "== 7. autoRollback=true なら自動で戻る"
run deploy main
echo 1000 >"$WORK/fake-score"
set +e
out=$(ISUH_CONFIG="$AUTOCFG" "$ISUH" bench "$V1" --note "regression, auto" 2>"$ERR")
rc=$?
set -e
[ "$rc" = 50 ] || ng "exit=$rc (50 のはず)"
grep -q "ロールバックします" "$ERR" || ng "ロールバックのログがありません"
grep -q "ロールバック完了" "$ERR" || ng "完了のログがありません"
[ "$(version_on isu1)" = "v2" ] || ng "isu1 が good に戻っていない"
[ "$(version_on isu2)" = "v2" ] || ng "isu2 が good に戻っていない"
ok "終盤の運用（自動）: 低下したら直前の good へ戻す"

echo "== 8. 戻したことが report.md と meta.json に残る"
runid=$(sed -n 's/.*runId=\([^ ]*\).*/\1/p' "$ERR" | head -1)
[ -n "$runid" ] || ng "run-id を取れません"
dir="$ISUH_RUNS_DIR/$runid"
grep -q "ロールバック" "$dir/report.md" || ng "report.md に残っていません"
grep -q "この変更はサーバに載っていない" "$dir/report.md" || ng "今の状態が分かりません"
node -e '
const m = require(process.argv[1]);
if (typeof m.rolledBackTo !== "string" || m.rolledBackTo.length < 40) {
  console.error("rolledBackTo が sha でない: " + JSON.stringify(m.rolledBackTo)); process.exit(1);
}
' "$dir/meta.json" || ng "meta.json に残っていません"
ok "「戻した」ことがログとレポートと meta.json に残る"

echo "== 9. ベンチが fail したときも戻る"
rm -f "$WORK/fake-score"
run deploy "$V1"
# ジョブを failed で返させる
printf 'fail\n' >"$WORK/fake-status"
set +e
ISUH_CONFIG="$AUTOCFG" "$ISUH" bench "$V1" --note "bench failed" >/dev/null 2>"$ERR"
rc=$?
set -e
rm -f "$WORK/fake-status"
[ "$rc" = 40 ] || ng "exit=$rc (40 のはず)"
grep -q "ロールバック完了" "$ERR" || ng "fail のときに戻していません"
[ "$(version_on isu1)" = "v2" ] || ng "isu1 が good に戻っていない"
ok "スコア低下だけでなくベンチ失敗でも戻す"

echo "== 10. 引数不正は 64"
run rollback --bogus
[ "$rc" = 64 ] || ng "exit=$rc"
ok "不明なオプションは 64"

run deploy main
echo
echo "全 $passed 件パス"
