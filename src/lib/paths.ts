import { mkdirSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

/** リポジトリルート（src/lib からふたつ上）。 */
export const repoRoot = (): string => join(dirname(fileURLToPath(import.meta.url)), "..", "..");

/** run-id — docs/DESIGN.md §4。辞書順 = 時系列順になること。 */
export const newRunId = (sha: string, now = new Date()): string => {
  const p = (n: number) => String(n).padStart(2, "0");
  const stamp =
    `${now.getFullYear()}${p(now.getMonth() + 1)}${p(now.getDate())}` +
    `-${p(now.getHours())}${p(now.getMinutes())}${p(now.getSeconds())}`;
  return `${stamp}-${sha.slice(0, 7)}`;
};

/** runs/<run-id>/ の絶対パス。無ければ作る。 */
export const runDir = (runId: string): string => {
  const dir = join(repoRoot(), "runs", runId);
  mkdirSync(dir, { recursive: true });
  return dir;
};

export const historyFile = (): string => join(repoRoot(), "runs", "history.tsv");
