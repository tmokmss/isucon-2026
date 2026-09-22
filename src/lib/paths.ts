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

/**
 * runs/ の場所。ISUH_RUNS_DIR で上書きできる（テスト用。ISUH_CONFIG と同じ考え方で、
 * 追跡している runs/history.tsv をテストが汚さないようにするため）。
 */
const runsRoot = (): string => process.env["ISUH_RUNS_DIR"] ?? join(repoRoot(), "runs");

/** runs/<run-id>/ の絶対パス。無ければ作る。 */
export const runDir = (runId: string): string => {
  const dir = join(runsRoot(), runId);
  mkdirSync(dir, { recursive: true });
  return dir;
};

/** runs/history.tsv の絶対パス。親ディレクトリは無ければ作る。 */
export const historyFile = (): string => {
  const root = runsRoot();
  mkdirSync(root, { recursive: true });
  return join(root, "history.tsv");
};
