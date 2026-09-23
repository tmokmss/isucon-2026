#!/bin/sh
# レポート生成(#4) とスコア台帳(#5) の検証。先に ./test/up.sh を実行しておくこと。
set -eu
cd "$(dirname "$0")"
TEST_DIR=$(pwd)
ROOT=$(cd .. && pwd)
WORK="$TEST_DIR/.work"

export ISUH_CONFIG="$TEST_DIR/config.test.ts"
export ISUH_SSH_CONFIG="$WORK/ssh_config"
# 追跡している runs/history.tsv を汚さない
export ISUH_RUNS_DIR="$WORK/runs"
ISUH="$ROOT/bin/isuh"
ERR="$WORK/report.err"
HISTORY="$ISUH_RUNS_DIR/history.tsv"

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
run_id_from_err() { sed -n 's/.*runId=\([^ ]*\).*/\1/p' "$ERR" | head -1; }

rm -rf "$ISUH_RUNS_DIR"
mkdir -p "$ISUH_RUNS_DIR"
rm -f "$WORK/fake-score"
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

echo "== 1. 履歴が空なら history は何も出さずに 0"
run history
[ "$rc" = 0 ] || ng "exit=$rc"
[ -z "$out" ] || ng "stdout='$out' (空のはず)"
grep -q "まだ実行履歴がありません" "$ERR" || ng "案内が出ていません"
ok "履歴ゼロ件でもエラーにしない"

echo "== 2. bench が report.md / meta.json / history.tsv を残す"
run bench main --note "first run"
[ "$rc" = 0 ] || ng "exit=$rc (0 のはず)"
[ "$out" = "12345" ] || ng "stdout='$out'"
RUN1=$(run_id_from_err)
DIR1="$ISUH_RUNS_DIR/$RUN1"
[ -f "$DIR1/report.md" ] || ng "report.md がありません"
[ -f "$DIR1/meta.json" ] || ng "meta.json がありません"
[ -f "$HISTORY" ] || ng "history.tsv がありません"
ok "1 ラン = 1 ディレクトリに揃う ($RUN1)"

echo "== 3. report.md の見出しが契約どおり"
expected="# run $RUN1
## Summary
## Bench
## Access log (alp)
## Slow query (pt-query-digest)
## CPU profile (pprof)
## Resources
## Diff from previous run"
# コードブロックの中（pt-query-digest の出力は # で始まる行だらけ）は見出しではない
actual=$(awk '/^```/ { fenced = !fenced; next } !fenced && /^#{1,2} /' "$DIR1/report.md")
[ "$actual" = "$expected" ] || ng "見出しが違います:
--- 実際 ---
$actual
--- 期待 ---
$expected"
ok "見出しが docs/DESIGN.md §4 と一致し、順序も同じ"

echo "== 4. Summary にスコア・sha・前回比・ベスト比が入る"
grep -q "| score | 12345 |" "$DIR1/report.md" || ng "score がありません"
grep -q "| status | pass |" "$DIR1/report.md" || ng "status がありません"
grep -q "初回" "$DIR1/report.md" || ng "初回であることが書かれていません"
grep -q "| note | first run |" "$DIR1/report.md" || ng "note がありません"
ok "Summary だけで何が起きたか分かる"

echo "== 5. Resources が飽和したホストを名指しする"
grep -q "| host | CPU 使用率" "$DIR1/report.md" || ng "リソースの表がありません"
grep -qE "^\| isu1 \(?[0-9]* ?c?p?u?\)? \|" "$DIR1/report.md" || ng "isu1 の行がありません"
grep -qE "飽和していない|ここから手を付ける" "$DIR1/report.md" || ng "どこを見ればいいかの結論がありません"
ok "vmstat を要約して次に見るべきホストを言う"

echo "== 6. meta.json の形が契約どおり"
node -e '
const m = require(process.argv[1]);
const need = ["runId","startedAt","finishedAt","ref","sha","score","status","note","benchJobId","prevRunId","bestRunId"];
const missing = need.filter((k) => !(k in m));
if (missing.length) { console.error("欠けている: " + missing.join(",")); process.exit(1); }
if (m.status !== "pass" || m.score !== 12345) { console.error("値が違う: " + JSON.stringify(m)); process.exit(1); }
if (m.prevRunId !== null) { console.error("prevRunId は null のはず"); process.exit(1); }
if (m.bestRunId !== m.runId) { console.error("bestRunId が自分ではない"); process.exit(1); }
if (!/^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[+-]\d{2}:\d{2}$/.test(m.startedAt)) { console.error("startedAt の形式: " + m.startedAt); process.exit(1); }
' "$DIR1/meta.json" || ng "meta.json の中身が契約と違います"
ok "meta.json に 11 項目が揃い、時刻はオフセット付き ISO8601"

echo "== 7. history.tsv がタブ区切りで追記される"
header=$(head -1 "$HISTORY")
[ "$header" = "$(printf 'run_id\tstarted_at\tsha\tref\tscore\tstatus\tnote')" ] || ng "ヘッダが違います: $header"
[ "$(wc -l <"$HISTORY" | tr -d ' ')" = "2" ] || ng "行数が合いません"
row=$(sed -n 2p "$HISTORY")
[ "$(printf '%s' "$row" | cut -f5)" = "12345" ] || ng "score 列が違います"
[ "$(printf '%s' "$row" | cut -f6)" = "pass" ] || ng "status 列が違います"
[ "$(printf '%s' "$row" | cut -f7)" = "first run" ] || ng "note 列が違います"
ok "7 カラムが揃い、追記のみで壊れない"

echo "== 8. 2 回目は前回比が出て、ベストが更新される"
echo 20000 >"$WORK/fake-score"
run bench main --note "second"
[ "$rc" = 0 ] || ng "exit=$rc"
[ "$out" = "20000" ] || ng "stdout='$out'"
RUN2=$(run_id_from_err)
DIR2="$ISUH_RUNS_DIR/$RUN2"
grep -q "+7655" "$DIR2/report.md" || ng "前回比が出ていません"
grep -q "$RUN1" "$DIR2/report.md" || ng "前回の run-id が参照されていません"
grep -q "ベスト更新" "$ERR" || ng "ベスト更新が知らされていません"
ok "スコアの増減が前回・ベストとの比較で出る"

echo "== 8b. 別の sha なら Diff にコードの差分が出る"
echo 21000 >"$WORK/fake-score"
run bench main~1 --note "old revision"
[ "$rc" = 0 ] || ng "exit=$rc"
RUN3=$(run_id_from_err)
DIR3="$ISUH_RUNS_DIR/$RUN3"
grep -q "version.go" "$DIR3/report.md" || ng "git diff --stat が出ていません"
grep -q "$RUN2" "$DIR3/report.md" || ng "比較対象の run-id がありません"
ok "前回と sha が違えば何を変えたかが出る"

echo "== 9. スコアが落ちたら exit 50 で止まる"
echo 1000 >"$WORK/fake-score"
run bench main --note "regression"
[ "$rc" = 50 ] || ng "exit=$rc (50 のはず)"
grep -q "スコアが低下しました" "$ERR" || ng "低下の説明がありません"
grep -q "isuh deploy" "$ERR" || ng "戻し方が案内されていません"
[ "$(wc -l <"$HISTORY" | tr -d ' ')" = "5" ] || ng "低下したランが台帳に残っていません"
ok "ベスト × scoreRegressionRatio を下回れば 50（戻し方つき）"

echo "== 10. 失敗したランも台帳とレポートに残る"
run bench broken
[ "$rc" = 20 ] || ng "exit=$rc (20 のはず)"
RUN4=$(run_id_from_err)
DIR4="$ISUH_RUNS_DIR/$RUN4"
[ -f "$DIR4/report.md" ] || ng "失敗時に report.md がありません"
grep -q "失敗の理由" "$DIR4/report.md" || ng "失敗の理由が書かれていません"
grep -q "ビルドに失敗しました" "$DIR4/report.md" || ng "エラーの中身がありません"
last=$(tail -1 "$HISTORY")
[ "$(printf '%s' "$last" | cut -f6)" = "error" ] || ng "status が error ではありません"
[ "$(printf '%s' "$last" | cut -f5)" = "0" ] || ng "失敗時の score は 0 のはず"
ok "デプロイで落ちても何を試したかが残る"

echo "== 11. history がベストを強調して並べる"
run history
[ "$rc" = 0 ] || ng "exit=$rc"
[ "$(printf '%s\n' "$out" | wc -l | tr -d ' ')" = "6" ] || ng "行数が合いません:
$out"
printf '%s\n' "$out" | grep -q "^\*  *$RUN3" || ng "ベストに * が付いていません:
$out"
printf '%s\n' "$out" | grep -q "+7655" || ng "delta 列がありません"
printf '%s\n' "$out" | grep -q "error" || ng "失敗したランが出ていません"
ok "どの変更でスコアがどう動いたかが一目で分かる"

echo "== 12. -n で件数を絞れる"
run history -n 2
[ "$(printf '%s\n' "$out" | wc -l | tr -d ' ')" = "3" ] || ng "-n 2 で 3 行（ヘッダ込み）のはず"
run history -n
[ "$rc" = 64 ] || ng "exit=$rc (64 のはず)"
run history --bogus
[ "$rc" = 64 ] || ng "exit=$rc (64 のはず)"
ok "-n が効き、引数不正は 64"

rm -f "$WORK/fake-score"
echo
echo "全 $passed 件パス"
