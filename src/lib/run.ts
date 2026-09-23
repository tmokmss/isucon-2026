import { appendFileSync, existsSync, readFileSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import { historyFile, runDir } from "./paths.js";

/**
 * ラン 1 回分の記録 — docs/DESIGN.md §4
 *
 * meta.json が機械の読むもの、history.tsv が「どの変更でスコアがどう動いたか」の台帳。
 * 前の版へ戻したいときは、ここ（isuh history）で sha を見て isuh deploy <sha> する。
 */

export type RunStatus = "pass" | "fail" | "error";

export interface RunMeta {
  runId: string;
  startedAt: string;
  finishedAt: string;
  ref: string;
  sha: string;
  score: number;
  status: RunStatus;
  note: string;
  benchJobId: string;
  prevRunId: string | null;
  bestRunId: string | null;
}

/** ローカル時刻のオフセット付き ISO8601。当日ログと突き合わせるので UTC にしない。 */
export const localIso = (date = new Date()): string => {
  const p = (n: number): string => String(n).padStart(2, "0");
  const offset = -date.getTimezoneOffset();
  const sign = offset >= 0 ? "+" : "-";
  const hours = p(Math.floor(Math.abs(offset) / 60));
  const minutes = p(Math.abs(offset) % 60);
  return (
    `${date.getFullYear()}-${p(date.getMonth() + 1)}-${p(date.getDate())}` +
    `T${p(date.getHours())}:${p(date.getMinutes())}:${p(date.getSeconds())}${sign}${hours}:${minutes}`
  );
};

// ---------------------------------------------------------------------------
// meta.json
// ---------------------------------------------------------------------------

export const writeMeta = (dir: string, meta: RunMeta): void => {
  writeFileSync(join(dir, "meta.json"), `${JSON.stringify(meta, null, 2)}\n`);
};

export const readMeta = (runId: string): RunMeta | undefined => {
  const path = join(runDir(runId), "meta.json");
  if (!existsSync(path)) return undefined;
  try {
    return JSON.parse(readFileSync(path, "utf8")) as RunMeta;
  } catch {
    return undefined;
  }
};

// ---------------------------------------------------------------------------
// history.tsv
// ---------------------------------------------------------------------------

export const HISTORY_COLUMNS = [
  "run_id",
  "started_at",
  "sha",
  "ref",
  "score",
  "status",
  "note",
] as const;

export interface HistoryRow {
  runId: string;
  startedAt: string;
  sha: string;
  ref: string;
  score: number;
  status: RunStatus;
  note: string;
}

/** note にタブや改行が入ると台帳が壊れる。 */
const sanitize = (text: string): string => text.replace(/[\t\r\n]+/g, " ").trim();

export const readHistory = (): HistoryRow[] => {
  const path = historyFile();
  if (!existsSync(path)) return [];
  const rows: HistoryRow[] = [];
  for (const line of readFileSync(path, "utf8").split("\n")) {
    if (line.trim() === "" || line.startsWith("run_id\t")) continue;
    const fields = line.split("\t");
    if (fields.length < HISTORY_COLUMNS.length) continue;
    const score = Number(fields[4]);
    rows.push({
      runId: fields[0]!,
      startedAt: fields[1]!,
      sha: fields[2]!,
      ref: fields[3]!,
      score: Number.isFinite(score) ? score : 0,
      status: (fields[5] as RunStatus) ?? "error",
      note: fields[6] ?? "",
    });
  }
  return rows;
};

/** 追記のみ。既存行は書き換えない。 */
export const appendHistory = (row: HistoryRow): void => {
  const path = historyFile();
  if (!existsSync(path)) {
    writeFileSync(path, `${HISTORY_COLUMNS.join("\t")}\n`);
  }
  const line = [
    row.runId,
    row.startedAt,
    row.sha,
    row.ref,
    // score は status が pass 以外なら 0（docs/DESIGN.md §4）
    String(row.status === "pass" ? row.score : 0),
    row.status,
    sanitize(row.note),
  ].join("\t");
  appendFileSync(path, `${line}\n`);
};

export const historyRowOf = (meta: RunMeta): HistoryRow => ({
  runId: meta.runId,
  startedAt: meta.startedAt,
  sha: meta.sha,
  ref: meta.ref,
  score: meta.score,
  status: meta.status,
  note: meta.note,
});

/** 直近のラン（成否を問わない）。 */
export const lastRun = (rows: readonly HistoryRow[]): HistoryRow | undefined => rows.at(-1);

/**
 * ベストスコアのラン。
 * pass したものだけが対象。同点なら先に出たほう（= より枯れているほう）を残す。
 */
export const bestRun = (rows: readonly HistoryRow[]): HistoryRow | undefined => {
  let best: HistoryRow | undefined;
  for (const row of rows) {
    if (row.status !== "pass") continue;
    if (best === undefined || row.score > best.score) best = row;
  }
  return best;
};
