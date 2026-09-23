#!/bin/sh
# ロック機構(#2) の検証。先に ./test/up.sh を実行しておくこと。
set -eu
cd "$(dirname "$0")"
TEST_DIR=$(pwd)
ROOT=$(cd .. && pwd)

export ISUH_CONFIG="$TEST_DIR/config.test.ts"
export ISUH_SSH_CONFIG="$TEST_DIR/.work/ssh_config"
ISUH="$ROOT/bin/isuh"
# tsx の CLI は子プロセスで node を起動してしまい、kill の検証にならない（bin/isuh と同じ理由）
PROBE="node --import tsx"
ERR="$TEST_DIR/.work/lock.err"
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

echo "== 4. bench は失敗してもロックを解放する"
$SSH isu1 "rm -f $TRACE"
set +e
"$ISUH" bench main >/dev/null 2>"$ERR"
rc=$?
set -e
# ベンチ起動(#3)が未実装なので 70 で終わる。そこまでの手順は本物。
[ "$rc" = 70 ] || ng "exit=$rc (70 のはず)"
grep -q "ロックを解放しました" "$ERR" || ng "ロックを解放していません"
lock_free || ng "bench 後もロックが残っています"
ok "途中で失敗しても finally で解放される"

echo "== 5. bench が計測の生データを残す"
runid=$(ls -1t "$ROOT/runs" | grep -E '^[0-9]{8}-[0-9]{6}-' | head -1)
[ -n "$runid" ] || ng "runs/<run-id>/ が作られていません"
dir="$ROOT/runs/$runid"
for f in alp.txt ptquery.txt pprof.txt; do
  [ -f "$dir/$f" ] || ng "$f がありません"
done
[ -s "$dir/resource/isu1.txt" ] || ng "resource/isu1.txt が空です"
grep -q "^# host=isu1" "$dir/resource/isu1.txt" || ng "resource/isu1.txt の中身が違います"
grep -qE "^ *[0-9]+ +[0-9]+" "$dir/resource/isu1.txt" || ng "vmstat の行がありません"
grep -q "未取得" "$dir/alp.txt" || ng "alp 未導入時に理由が書かれていません"
ok "runs/$runid に alp/ptquery/pprof/resource が揃う（未導入なら理由付きで未取得）"

echo "== 6. ログのリセットが走っている"
grep -q "リソース計測を開始しました" "$ERR" || ng "リソース計測が開始されていません"
ok "ログのリセットとリソース計測がベンチ前に走る"

echo
echo "全 $passed 件パス"
