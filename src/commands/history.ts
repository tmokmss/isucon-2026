import type { Command } from "../lib/command.js";
import { usageError } from "../lib/exit.js";
import { log, out } from "../lib/log.js";
import { historyFile } from "../lib/paths.js";
import { type HistoryRow, bestRun, readHistory } from "../lib/run.js";

const pad = (text: string, width: number): string => text.padEnd(width, " ");
const padStart = (text: string, width: number): string => text.padStart(width, " ");

const signed = (n: number): string => (n > 0 ? `+${n}` : String(n));

export const command: Command = {
  summary: "スコア台帳の表示",
  usage: `usage: isuh history [-n <count>]

runs/history.tsv を読んでスコア台帳を表示する。ベスト更新を強調する。
「どの変更でスコアがどう動いたか」が一目で分かることが目的で、
前の版へ戻すときの材料にもなる（sha を見て isuh deploy <sha>）。

  -n <count>   表示する件数（既定: 全件）

  列: * = ベスト / run_id / score / Δ（直前に pass したランとの差）/
      status / sha / note

stdout には表だけを出す。件数やファイル位置は stderr。`,

  async run(argv) {
    let count = 0;
    for (let i = 0; i < argv.length; i++) {
      const arg = argv[i]!;
      if (arg === "-n") {
        const value = argv[++i];
        if (value === undefined) throw usageError("-n に値がありません");
        count = Number.parseInt(value, 10);
        if (Number.isNaN(count)) throw usageError(`-n は数値で指定する: ${value}`);
      } else {
        throw usageError(`不明な引数: ${arg}`);
      }
    }

    const rows = readHistory();
    if (rows.length === 0) {
      log.info(`まだ実行履歴がありません: ${historyFile()}`);
      return;
    }

    // Δ は「直前に pass したラン」との差。fail を挟んでも比較が続く。
    const deltas = new Map<string, string>();
    let previousPass: HistoryRow | undefined;
    for (const row of rows) {
      if (row.status !== "pass") {
        deltas.set(row.runId, "—");
        continue;
      }
      deltas.set(row.runId, previousPass === undefined ? "—" : signed(row.score - previousPass.score));
      previousPass = row;
    }

    const best = bestRun(rows);
    const shown = count > 0 ? rows.slice(-count) : rows;

    const cells = shown.map((row) => ({
      mark: best !== undefined && row.runId === best.runId ? "*" : " ",
      runId: row.runId,
      score: row.status === "pass" ? String(row.score) : "—",
      delta: deltas.get(row.runId) ?? "—",
      status: row.status,
      sha: row.sha.slice(0, 7),
      note: row.note,
    }));

    const width = (key: keyof (typeof cells)[number], header: string): number =>
      Math.max(header.length, ...cells.map((c) => c[key].length));

    const w = {
      runId: width("runId", "run_id"),
      score: width("score", "score"),
      delta: width("delta", "delta"),
      status: width("status", "status"),
      sha: width("sha", "sha"),
    };

    out(
      [
        " ",
        pad("run_id", w.runId),
        padStart("score", w.score),
        padStart("delta", w.delta),
        pad("status", w.status),
        pad("sha", w.sha),
        "note",
      ].join("  "),
    );
    for (const c of cells) {
      out(
        [
          c.mark,
          pad(c.runId, w.runId),
          padStart(c.score, w.score),
          padStart(c.delta, w.delta),
          pad(c.status, w.status),
          pad(c.sha, w.sha),
          c.note,
        ]
          .join("  ")
          .trimEnd(),
      );
    }

    const passed = rows.filter((r) => r.status === "pass").length;
    log.info(
      `${shown.length}/${rows.length} 件を表示（pass ${passed}）` +
        (best === undefined ? "" : ` / ベスト ${best.score} = ${best.runId} (${best.sha.slice(0, 7)})`),
    );
  },
};
