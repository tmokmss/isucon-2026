/**
 * report.ts の Resources 節（resourceSection）の確認。metrics.sh から呼ぶ。
 * - 平均は低くても、終盤に張り付いたホストを飽和として名指しする
 * - ベンチマーカーのホストは「ここから手を付ける」から外し、飽和したら別に知らせる
 * 通ったケース数を stdout に出す。
 */
import { mkdtempSync, mkdirSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { resourceSection } from "../src/lib/report.js";

let passed = 0;
const check = (name: string, cond: boolean, detail = ""): void => {
  if (!cond) {
    process.stderr.write(`NG: ${name}\n${detail}\n`);
    process.exit(1);
  }
  passed++;
};

/** vmstat 1 の出力（1 行目は起動からの平均なので捨てられる）。idles は 2 行目以降の id 列 */
const vmstat = (nproc: number, idles: number[]): string =>
  [
    `# host=h nproc=${nproc}`,
    "procs -----------memory---------- ---swap-- -----io---- -system-- -------cpu-------",
    " r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st gu",
    " 1  0      0 2307824  39184 877720    0    0    72   172 1504    6  6  1 93  0  0  0",
    ...idles.map((id) => ` 1  0      0 2307824  39184 877720    0    0     0  1000 5000 5000 ${100 - id - 1}  1 ${id}  0  0  0`),
    "",
  ].join("\n");

// 前半は空いていて、後半だけ張り付く（平均は 80% 未満、90% 以上の時間は 4 割）
const lateSaturated = [...Array(12).fill(60), ...Array(8).fill(3)];
// ずっと空いている
const idle = Array(20).fill(70);

const dir = mkdtempSync(join(tmpdir(), "resource-probe-"));
try {
  mkdirSync(join(dir, "resource"));
  writeFileSync(join(dir, "resource", "isu1.txt"), vmstat(2, lateSaturated));
  writeFileSync(join(dir, "resource", "isu2.txt"), vmstat(2, idle));
  writeFileSync(join(dir, "resource", "isu3.txt"), vmstat(8, lateSaturated));

  const withBench = resourceSection(dir, "isu3");
  const culpritLine = withBench.split("\n").find((l) => l.includes("ここから手を付ける")) ?? "";
  check("終盤だけ張り付いたホストを飽和として名指しする", culpritLine.includes("isu1（CPU が飽和"), withBench);
  check("空いているホストは名指ししない", !culpritLine.includes("isu2"), withBench);
  check("ベンチのホストは「ここから手を付ける」に入らない", !culpritLine.includes("isu3"), withBench);
  check("ベンチのホストの飽和を知らせる", withBench.includes("ベンチマーカーのホスト isu3 が飽和した"), withBench);
  check("表でベンチのホストが分かる", withBench.includes("| isu3（ベンチ） (8 cpu) |"), withBench);
  check("張り付いた時間の列がある", withBench.includes("90% 以上の時間") && /\| isu1 \(2 cpu\) \|[^|]+\| 40% \|/.test(withBench), withBench);

  const withoutBench = resourceSection(dir);
  check("ベンチのホストを渡さなければ警告しない", !withoutBench.includes("ベンチマーカーのホスト"), withoutBench);
  check(
    "ベンチのホストを渡さなければ従来どおり名指しする",
    (withoutBench.split("\n").find((l) => l.includes("ここから手を付ける")) ?? "").includes("isu3"),
    withoutBench,
  );

  // ベンチのホストが空いていれば警告しない
  writeFileSync(join(dir, "resource", "isu3.txt"), vmstat(8, idle));
  check("ベンチのホストが空いていれば警告しない", !resourceSection(dir, "isu3").includes("ベンチマーカーのホスト"));
} finally {
  rmSync(dir, { recursive: true, force: true });
}

process.stdout.write(`${passed}\n`);
