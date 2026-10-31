#!/bin/bash
# リーダーボードの上位と自チームの順位を出す（ポータルの観客向けダッシュボード）
# 使い方: setup/leaderboard.sh [上位の件数]
set -euo pipefail
cd "$(dirname "$0")/.."
tmp=$(mktemp)
trap 'rm -f "$tmp"' EXIT
bin/isuh portal get /api/audience/dashboard > "$tmp" 2>/dev/null
node -e '
const [file, top, me] = [process.argv[1], Number(process.argv[2]), process.argv[3]];
const rows = JSON.parse(require("fs").readFileSync(file, "utf8"))["1"]["1"].map((e) => ({
  id: e["16"]?.["1"], name: e["16"]?.["2"], best: Number(e["2"]?.["1"] ?? 0), latest: Number(e["3"]?.["1"] ?? 0),
}));
rows.sort((a, b) => b.best - a.best);
const now = new Date().toLocaleTimeString("ja-JP", { timeZone: "Asia/Tokyo" });
console.log(`${now} 全${rows.length}チーム`);
rows.forEach((r, i) => { if (i < top || r.name === me) console.log(`${i + 1}\t${r.best}\t${r.latest}\t${r.name}`); });
' "$tmp" "${1:-10}" "もなまっち"
