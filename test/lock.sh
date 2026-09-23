#!/bin/sh
# ロック機構(#2) の検証。先に ./test/up.sh を実行しておくこと。
set -eu
cd "$(dirname "$0")"
TEST_DIR=$(pwd)
ROOT=$(cd .. && pwd)

export ISUH_CONFIG="$TEST_DIR/config.test.ts"
export ISUH_SSH_CONFIG="$TEST_DIR/.work/ssh_config"
# 追跡している runs/history.tsv を汚さない
export ISUH_RUNS_DIR="$TEST_DIR/.work/runs-lock"
ISUH="$ROOT/bin/isuh"
# tsx の CLI は子プロセスで node を起動してしまい、kill の検証にならない（bin/isuh と同じ理由）
PROBE="node --import tsx"
WORK="$TEST_DIR/.work"
ERR="$WORK/lock.err"
SSH="ssh -F $ISUH_SSH_CONFIG -o ConnectTimeout=5"

LOCK=/tmp/isucon-bench.lock
TRACE=/tmp/isuh-lock-trace.txt

passed=0
ng() {
  printf '  NG  %s\n' "$*" >&2
  exit 1
}
ok() {
  passed=$((passed + 1))
  printf '  ok  %s\n' "$*"
}

# ロックが空いていれば 0
lock_free() { $SSH isu1 "flock -n $LOCK -c true" >/dev/null 2>&1; }

wait_until_held() {
  i=0
  while lock_free; do
    i=$((i + 1))
    [ "$i" -gt 100 ] && return 1
    sleep 0.1
  done
  return 0
}

wait_until_free() {
  i=0
  until lock_free; do
    i=$((i + 1))
    [ "$i" -gt "$1" ] && return 1
    sleep 0.1
  done
  return 0
}

# bench を通しで走らせるためにポータルの代役を立てる
printf 'sessionid=valid-session; csrftoken=tok123\n' >"$WORK/portal-cookie"
node --import tsx fake-portal.ts >"$WORK/fake-portal.log" 2>&1 &
portal_pid=$!
trap 'kill "$portal_pid" 2>/dev/null || true' EXIT
i=0
until curl -s -o /dev/null "http://127.0.0.1:18080/auth/login/" 2>/dev/null; do
  i=$((i + 1))
  [ "$i" -gt 50 ] && ng "fake-portal が起動しません"
  sleep 0.1
done

$SSH isu1 "rm -f $TRACE $LOCK.owner"
lock_free || ng "開始時点でロックが握られています"

echo "== 1. 2 つ同時に走らせると直列になる"
$PROBE lock-probe.ts A 1500 >/dev/null 2>&1 &
pid_a=$!
sleep 0.5
$PROBE lock-probe.ts B 300 >/dev/null 2>&1 &
pid_b=$!
wait "$pid_a" || ng "probe A が失敗しました"
wait "$pid_b" || ng "probe B が失敗しました"
trace=$($SSH isu1 "cat $TRACE" | tr '\n' ' ')
case "$trace" in
  "ENTER A EXIT A ENTER B EXIT B ") ok "クリティカルセクションが重ならない ($trace)" ;;
  *) ng "重なっています: $trace" ;;
esac

echo "== 2. 取れなければ保持者を出して exit 10"
$SSH isu1 "rm -f $TRACE"
$PROBE lock-probe.ts HOLD -1 >/dev/null 2>&1 &
pid_hold=$!
wait_until_held || ng "HOLD がロックを取れていません"

set +e
"$ISUH" bench main >/dev/null 2>"$ERR"
rc=$?
set -e
[ "$rc" = 10 ] || ng "exit=$rc (10 のはず)"
grep -q "保持者:" "$ERR" || ng "保持者が表示されていません"
grep -q "probe HOLD" "$ERR" || ng "保持者の中身が出ていません: $(grep 保持者 "$ERR")"
ok "待って取れなければ 10、保持者が分かる"

echo "== 3. kill -9 してもロックが残らない"
kill -9 "$pid_hold" 2>/dev/null || true
wait "$pid_hold" 2>/dev/null || true
wait_until_free 100 || ng "kill 後もロックが残っています"
ok "プロセスを kill すればロックは自動で解放される"

echo "== 4. bench が通しで完走する"
set +e
score=$("$ISUH" bench main 2>"$ERR")
rc=$?
set -e
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ "$score" = "12345" ] || ng "stdout='$score' (12345 のはず)"
grep -q "ロックを解放しました" "$ERR" || ng "ロックを解放していません"
lock_free || ng "bench 後もロックが残っています"
# このランの run-id を控える（後続のテストで別のランに上書きされるため）
runid=$(sed -n 's/.*runId=\([^ ]*\).*/\1/p' "$ERR" | head -1)
[ -n "$runid" ] || ng "run-id を取れません"
ok "ロック → deploy → ログリセット → 計測開始 → ベンチ → 回収 → 解放 が通る"

echo "== 4b. 途中で失敗してもロックを解放する"
set +e
"$ISUH" bench broken >/dev/null 2>"$ERR"
rc=$?
set -e
[ "$rc" = 20 ] || ng "exit=$rc (20 のはず)"
lock_free || ng "失敗後もロックが残っています"
grep -q "ロックを解放しました" "$ERR" || ng "ロックを解放していません"
ok "デプロイで落ちても finally で解放される"

echo "== 5. bench が計測の生データを残す"
dir="$ISUH_RUNS_DIR/$runid"
[ -d "$dir" ] || ng "runs/$runid が作られていません"
for f in alp.txt ptquery.txt pprof.txt bench.log; do
  [ -f "$dir/$f" ] || ng "$f がありません"
done
[ -s "$dir/resource/isu1.txt" ] || ng "resource/isu1.txt が空です"
grep -q "^# host=isu1" "$dir/resource/isu1.txt" || ng "resource/isu1.txt の中身が違います"
grep -qE "^ *[0-9]+ +[0-9]+" "$dir/resource/isu1.txt" || ng "vmstat の行がありません"
# alp は setup（test/setup.sh）を流したかどうかで入っていたりいなかったりする。
# 入っていなければ理由、入っていれば集計表が残ること
grep -qE "未取得|COUNT" "$dir/alp.txt" || ng "alp.txt に集計表も未取得の理由もありません"
ok "runs/$runid に alp/ptquery/pprof/resource が揃う（未導入なら理由付きで未取得）"

echo "== 6. ベンチの生出力が残る"
grep -q '"score":12345' "$dir/bench.log" || ng "bench.log にベンチの応答がありません"
ok "bench.log にポータルの応答がそのまま残る"

echo
echo "全 $passed 件パス"
