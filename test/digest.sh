#!/bin/sh
# report.md の Access log / Slow query 節が、列を絞った表に畳まれることの検証。
# 先に ./test/up.sh を実行しておくこと（alp と slow query log は冒頭で setup が入れる）。
#   - alp / pt-query-digest の生の出力（TSV / JSON）が alp.txt / ptquery.txt に残る
#   - report.md には判断に要る列だけの表が載る
#   - alpMatches が空なら、集計が当てにならないことを書く
set -eu
cd "$(dirname "$0")"
TEST_DIR=$(pwd)
ROOT=$(cd .. && pwd)
WORK="$TEST_DIR/.work"
ISUH="$ROOT/bin/isuh"
ERR="$WORK/digest.err"
SSH="ssh -F $WORK/ssh_config"

export ISUH_SSH_CONFIG="$WORK/ssh_config"
export ISUH_RUNS_DIR="$WORK/runs-digest"

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

# alpMatches を入れた設定（当日の使い方）
CFG="$WORK/config.digest.ts"
sed -e 's|join(here, ".work", |join(here, |g' \
  -e 's|alpMatches: \[\]|alpMatches: ["/api/user/.+"]|' \
  "$TEST_DIR/config.test.ts" >"$CFG"
grep -q '"/api/user/.+"' "$CFG" || ng "alpMatches を差し込めません"

rm -rf "$ISUH_RUNS_DIR"
rm -f "$WORK/fake-score" "$WORK/fake-status"
printf 'sessionid=valid-session; csrftoken=tok123\n' >"$WORK/portal-cookie"

ISUH_CONFIG="$CFG" "$ISUH" setup --only tools,logs >/dev/null 2>"$ERR" || ng "setup が失敗しました"

node --import tsx fake-portal.ts >"$WORK/fake-portal-digest.log" 2>&1 &
portal_pid=$!
load_pids=""
trap 'kill $portal_pid $load_pids 2>/dev/null || true' EXIT
i=0
until curl -s -o /dev/null "http://127.0.0.1:18080/auth/login/" 2>/dev/null; do
  i=$((i + 1))
  [ "$i" -gt 50 ] && ng "fake-portal が起動しません"
  sleep 0.1
done

# ベンチ中に nginx と MySQL へ負荷の代わりを流す（ログのリセットの後にも届くよう長めに）
start_load() {
  $SSH isu1 'for i in $(seq 1 125); do curl -s -o /dev/null http://localhost/api/user/$i; sleep 0.2; done' &
  load_pids="$load_pids $!"
  $SSH isu3 'for i in $(seq 1 125); do sudo mysql -NBe "SELECT $i AS isuh_digest_probe" >/dev/null; sleep 0.2; done' &
  load_pids="$load_pids $!"
}
stop_load() {
  kill $load_pids 2>/dev/null || true
  wait $load_pids 2>/dev/null || true
  load_pids=""
}
bench() {
  start_load
  set +e
  "$ISUH" bench main >/dev/null 2>"$ERR"
  rc=$?
  set -e
  stop_load
  [ "$rc" = 0 ] || ng "bench exit=$rc"
  runid=$(sed -n 's/.*runId=\([^ ]*\).*/\1/p' "$ERR" | head -1)
  DIR="$ISUH_RUNS_DIR/$runid"
  REPORT="$DIR/report.md"
  [ -f "$REPORT" ] || ng "report.md がありません"
}
# report.md の見出し ## <name> から次の見出しまで
section() { awk -v h="## $1" '$0 == h { on = 1; next } on && /^## / { exit } on' "$REPORT"; }

echo "== 1. bench を 1 回走らせる（alpMatches あり）"
ISUH_CONFIG="$CFG" bench
ok "完走した ($runid)"

echo "== 2. 生の出力は機械向けの形式で残る"
grep -q "^Count	1xx	2xx" "$DIR/alp.txt" || ng "alp.txt が TSV ではありません: $(head -3 "$DIR/alp.txt")"
node -e 'JSON.parse(require("fs").readFileSync(process.argv[1], "utf8")).classes.length' "$DIR/ptquery.txt" ||
  ng "ptquery.txt が JSON ではありません: $(head -c 300 "$DIR/ptquery.txt")"
ok "alp.txt は TSV、ptquery.txt は JSON"

echo "== 3. Access log 節は列を絞った表"
ALP=$(section "Access log (alp)")
printf '%s\n' "$ALP" | grep -q "^share	count	4xx	5xx	method	uri	sum_s	avg_s	p99_s	max_s	avg_bytes$" ||
  ng "alp の表の列が違います:
$ALP"
printf '%s\n' "$ALP" | grep -qE "^([0-9.]+%|—)	[0-9]+	[0-9]+	0	GET	/api/user/\.\+	" || ng "正規化した行がありません:
$ALP"
printf '%s\n' "$ALP" | grep -q "metrics.alpMatches が空" && ng "alpMatches があるのに警告が出ています"
printf '%s\n' "$ALP" | grep -q "生の出力は alp.txt" || ng "生の出力の場所がありません"
ok "share / count / 4xx / 5xx / uri / 時間 / 本文サイズだけの表になる"

echo "== 4. Slow query 節は 1 行 1 クエリの表と実例"
PT=$(section "Slow query (pt-query-digest)")
printf '%s\n' "$PT" | grep -qE "^全 [0-9]+ クエリ / [0-9]+ 種類 / 合計 " || ng "全体の件数がありません:
$PT"
printf '%s\n' "$PT" | grep -q "^rank	share	calls	total_s	avg_ms	p95_ms	rows_examined_avg	rows_sent_avg	query$" ||
  ng "pt の表の列が違います:
$PT"
printf '%s\n' "$PT" | grep -q "isuh_digest_probe" || ng "流したクエリがありません:
$PT"
printf '%s\n' "$PT" | grep -q "^-- 上位の実例（EXPLAIN 用）" || ng "実例がありません"
printf '%s\n' "$PT" | grep -q "^# Query 1" && ng "pt-query-digest の生の出力が混ざっています"
ok "1 行 1 クエリの表に畳まれ、上位の実例が付く"

echo "== 5. 小さくなっている"
raw=$(($(wc -c <"$DIR/alp.txt") + $(wc -c <"$DIR/ptquery.txt")))
digest=$(($(printf '%s' "$ALP" | wc -c) + $(printf '%s' "$PT" | wc -c)))
[ "$digest" -lt "$raw" ] || ng "畳んだ方が大きい: report=$digest raw=$raw"
ok "2 節の合計 ${digest} バイト（生の出力は ${raw} バイト）"

echo "== 6. alpMatches が空でも alp は落ちず、集計が当てにならないと書く"
ISUH_CONFIG="$TEST_DIR/config.test.ts" bench
ALP=$(section "Access log (alp)")
printf '%s\n' "$ALP" | grep -q "metrics.alpMatches が空" || ng "警告がありません"
# /api/user/1, /api/user/2 ... が別々になり alpLimit(30) を超える。alp の --limit は URI の
# 種類数の上限なので、そのまま渡すと "Too many URI's" で集計ごと落ちていた
printf '%s\n' "$ALP" | grep -q "Too many URI" && ng "alp が URI の種類数の上限で落ちています:
$ALP"
printf '%s\n' "$ALP" | grep -qE "^…（上位 30 / 全 [0-9]+ 種類）$" || ng "上位 30 件に絞られていません:
$ALP"
ok "正規化していないことを明示し、上位 alpLimit 件に絞る"

echo "== 7. 畳む処理そのもの"
set +e
cases=$(node --import tsx digest-probe.ts 2>"$ERR")
rc=$?
set -e
[ "$rc" = 0 ] || ng "digest-probe が失敗しました"
ok "列の選び方・割合・切り詰め・未取得のときの扱い（${cases}）"

echo
echo "全 $passed 件パス"
