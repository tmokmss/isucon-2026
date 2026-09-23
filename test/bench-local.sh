#!/bin/sh
# bench.runner=local（ベンチマーカーを ssh で直接実行する）の検証。先に ./test/up.sh を実行しておくこと。
set -eu
cd "$(dirname "$0")"
TEST_DIR=$(pwd)
ROOT=$(cd .. && pwd)

export ISUH_CONFIG="$TEST_DIR/config.test.ts"
export ISUH_SSH_CONFIG="$TEST_DIR/.work/ssh_config"
export ISUH_TEST_RUNNER=local
# 追跡している runs/history.tsv を汚さない
export ISUH_RUNS_DIR="$TEST_DIR/.work/runs-bench-local"
ISUH="$ROOT/bin/isuh"
WORK="$TEST_DIR/.work"
ERR="$WORK/bench-local.err"
SSH="ssh -F $ISUH_SSH_CONFIG -o ConnectTimeout=5"

passed=0
ng() {
  printf '  NG  %s\n' "$*" >&2
  exit 1
}
ok() {
  passed=$((passed + 1))
  printf '  ok  %s\n' "$*"
}

rm -rf "$ISUH_RUNS_DIR"

# 偽のベンチマーカー。ISUCON14 と同じ形の結果の行を出す。
# 途中経過は stderr、結果は stdout（本物も分かれている）。挙動は ~/fake-bench.mode で切り替える
$SSH isu1 'cat > fake-bench.sh' <<'EOF'
mode=$(cat fake-bench.mode 2>/dev/null || echo pass)
echo "level=DEBUG msg=eval reqs=4" >&2
echo "level=INFO msg=負荷走行を開始します"
case "$mode" in
  pass) echo "level=INFO msg=結果 pass=true スコア=759 種別エラー数=map[26:1]" ;;
  fail) echo "level=INFO msg=結果 pass=false スコア=0 種別エラー数=map[1:3]" ;;
  crash) echo "panic: connection refused" >&2; exit 3 ;;
esac
echo "level=INFO msg=負荷走行が終了しました"
EOF

# bench を 1 回走らせ、rc と run-id を控える
run_bench() {
  $SSH isu1 "echo $1 > fake-bench.mode"
  set +e
  score=$("$ISUH" bench main --skip-check --note "$1" 2>"$ERR")
  rc=$?
  set -e
  runid=$(sed -n 's/.*runId=\([^ ]*\).*/\1/p' "$ERR" | head -1)
  [ -n "$runid" ] || ng "run-id を取れません: $(tail -3 "$ERR")"
  dir="$ISUH_RUNS_DIR/$runid"
}

echo "== 1. ベンチマーカーを ssh で実行し、出力からスコアを読む"
run_bench pass
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず): $(tail -3 "$ERR")"
[ "$score" = "759" ] || ng "stdout='$score' (759 のはず)"
grep -q '"status": *"pass"' "$dir/meta.json" || ng "meta.json の status が pass ではありません"
grep -q '"benchJobId": *"local"' "$dir/meta.json" || ng "meta.json の benchJobId が local ではありません"
ok "pass=true スコア=759 → exit 0、stdout は 759"

echo "== 2. stdout と stderr の両方が bench.log に残る"
grep -q "スコア=759" "$dir/bench.log" || ng "bench.log に結果の行がありません"
grep -q "level=DEBUG" "$dir/bench.log" || ng "bench.log に stderr の行がありません"
grep -q "スコア=759" "$dir/report.md" || ng "report.md の Bench 節に出力がありません"
ok "bench.log と report.md の Bench 節にベンチの出力が残る"

echo "== 3. pass=false なら exit 40、スコアは 0 で記録"
run_bench fail
[ "$rc" = 40 ] || ng "exit=$rc (40 のはず)"
grep -q '"status": *"fail"' "$dir/meta.json" || ng "meta.json の status が fail ではありません"
grep -q '"score": *0' "$dir/meta.json" || ng "meta.json の score が 0 ではありません"
grep -q "pass=false" "$dir/report.md" || ng "report.md に失敗したベンチの出力がありません"
ok "失敗したベンチは fail として台帳に残り、出力は report.md で読める"

echo "== 4. ベンチマーカーが落ちたら exit 40、出力は残る"
run_bench crash
[ "$rc" = 40 ] || ng "exit=$rc (40 のはず)"
grep -q "exit=3" "$ERR" || ng "終了コードが表示されていません"
grep -q "panic: connection refused" "$dir/bench.log" || ng "bench.log に落ちたときの出力がありません"
ok "終了コードが 0 でなければ fail（exit=3 と表示）"

echo "== 5. ロックが残らない"
$SSH isu1 "flock -n /tmp/isucon-bench.lock -c true" || ng "bench 後もロックが残っています"
ok "失敗したランの後もロックは解放されている"

$SSH isu1 'rm -f fake-bench.sh fake-bench.mode'

echo
echo "全 $passed 件パス"
