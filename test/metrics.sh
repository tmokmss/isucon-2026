#!/bin/sh
# ベンチ中の計測の検証。先に ./test/up.sh を実行しておくこと。
#   - pprof がベンチの「後」ではなく「開始と同時」に取られていること
#   - top でプロセス単位の CPU が取れ、report.md で要約されること
#   - nginx が複数台なら alp が全台を集計すること
set -eu
cd "$(dirname "$0")"
TEST_DIR=$(pwd)
ROOT=$(cd .. && pwd)
WORK="$TEST_DIR/.work"

export ISUH_CONFIG="$TEST_DIR/config.test.ts"
export ISUH_SSH_CONFIG="$WORK/ssh_config"
export ISUH_RUNS_DIR="$WORK/runs-metrics"
ISUH="$ROOT/bin/isuh"
ERR="$WORK/metrics.err"

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

rm -rf "$ISUH_RUNS_DIR"
rm -f "$WORK/fake-score" "$WORK/fake-status"
printf 'isuxportal_session=valid-session\n' >"$WORK/portal-cookie"

node --import tsx fake-portal.ts >"$WORK/fake-portal-metrics.log" 2>&1 &
portal_pid=$!
trap 'kill "$portal_pid" 2>/dev/null || true' EXIT
i=0
until curl -s -o /dev/null "http://127.0.0.1:18080/api/session" 2>/dev/null; do
  i=$((i + 1))
  [ "$i" -gt 50 ] && ng "fake-portal が起動しません"
  sleep 0.1
done

echo "== 1. bench を 1 回走らせる"
set +e
"$ISUH" bench main >/dev/null 2>"$ERR"
rc=$?
set -e
[ "$rc" = 0 ] || ng "exit=$rc"
runid=$(sed -n 's/.*runId=\([^ ]*\).*/\1/p' "$ERR" | head -1)
DIR="$ISUH_RUNS_DIR/$runid"
[ -d "$DIR" ] || ng "runs/$runid がありません"
ok "完走した ($runid)"

echo "== 2. pprof はベンチの開始と同時に取り始める"
line_pprof=$(grep -n "pprof: .* のプロファイルを開始しました" "$ERR" | head -1 | cut -d: -f1)
line_enqueue=$(grep -n "ベンチを登録しました" "$ERR" | head -1 | cut -d: -f1)
[ -n "$line_pprof" ] || ng "pprof を開始していません"
[ -n "$line_enqueue" ] || ng "ベンチを登録していません"
[ "$line_pprof" -lt "$line_enqueue" ] || ng "pprof がベンチより後に始まっています"
ok "ログ上、pprof の開始がベンチの登録より前"

echo "== 3. プロファイルが回収・整形される"
[ -s "$DIR/cpu.pprof" ] || ng "生のプロファイル cpu.pprof がありません"
grep -q "Duration:" "$DIR/pprof.txt" || ng "pprof.txt が go tool pprof -top の出力になっていません:
$(head -5 "$DIR/pprof.txt")"
grep -q "pprof: 回収しました" "$ERR" || ng "回収のログがありません"
ok "cpu.pprof（生）と pprof.txt（-top）が残る"

echo "== 4. top でプロセス単位の CPU が取れる"
for h in isu1 isu2 isu3; do
  [ -s "$DIR/resource/top/$h.txt" ] || ng "resource/top/$h.txt がありません"
  grep -q "^top - " "$DIR/resource/top/$h.txt" || ng "$h の top の中身が違います"
done
ok "全台の top が resource/top/<host>.txt に残る"

echo "== 5. report.md がプロセス単位の CPU を要約する"
grep -q "CPU を食っているプロセス" "$DIR/report.md" || ng "プロセスの表がありません"
grep -qE "^\| isu1 \| " "$DIR/report.md" || ng "isu1 の行がありません"
ok "Resources 節に「どのプロセスが食っているか」が出る"

echo "== 6. top の要約ロジック"
set +e
cases=$(node --import tsx top-probe.ts 2>"$ERR")
rc=$?
set -e
[ "$rc" = 0 ] || ng "top の要約が期待と違います"
ok "同名プロセスの合算・1 フレーム目の除外・計測プロセスの除外（${cases}）"

echo "== 7. nginx が複数台なら alp は全台を集計する"
grep -q "^== isu1 ==" "$DIR/alp.txt" || ng "isu1 の節がありません"
grep -q "^== isu2 ==" "$DIR/alp.txt" || ng "isu2 の節がありません"
ok "alp.txt にホストごとの節が並ぶ"

echo "== 8. pprofEnabled=false なら取らずに理由を書く"
NOPPROF="$WORK/config.nopprof.ts"
sed -e 's/pprofEnabled: true/pprofEnabled: false/' -e 's|join(here, ".work", |join(here, |g' \
  "$TEST_DIR/config.test.ts" >"$NOPPROF"
set +e
ISUH_CONFIG="$NOPPROF" "$ISUH" bench main >/dev/null 2>"$ERR"
rc=$?
set -e
[ "$rc" = 0 ] || [ "$rc" = 50 ] || ng "exit=$rc"
runid2=$(sed -n 's/.*runId=\([^ ]*\).*/\1/p' "$ERR" | head -1)
grep -q "pprofEnabled = false" "$ISUH_RUNS_DIR/$runid2/pprof.txt" || ng "理由が書かれていません"
grep -q "pprof: .* のプロファイルを開始しました" "$ERR" && ng "無効なのに pprof を取っています"
ok "無効なら待ち時間ゼロで、未取得の理由だけ残す"

echo "== 9. Resources 節: 終盤の張り付きとベンチのホスト"
set +e
cases=$(node --import tsx resource-probe.ts 2>"$ERR")
rc=$?
set -e
[ "$rc" = 0 ] || ng "Resources の要約が期待と違います"
ok "終盤だけの張り付きを名指し・ベンチのホストは別に知らせる（${cases}）"

echo "== 10. bench の前の確認: 前回のランと中身が同じか"
set +e
cases=$(node --import tsx git-probe.ts 2>"$ERR")
rc=$?
set -e
[ "$rc" = 0 ] || ng "sameContent が期待と違います"
ok "同じ・違う・比べられないを区別する（${cases}）"

echo
echo "全 $passed 件パス"
