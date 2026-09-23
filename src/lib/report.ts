import { existsSync, readFileSync, readdirSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import type { Config } from "../config.js";
import type { HistoryRow, RunMeta } from "./run.js";
import { summarizeAlp, summarizePtQuery } from "./digest.js";
import { local } from "./ssh.js";

/**
 * report.md の生成 — docs/DESIGN.md §4
 *
 * **見出しは固定。** 生成側（ここ）と計測スキル(#14)がこの形で揃う。
 *
 * 品質基準はひとつだけ: **report.md 単体を読んで「次に何を直すべきか」が判断できること。**
 * これがそのまま Claude Code に渡される。だから
 *   - 生ログをそのまま貼るのではなく、飽和しているホストを名指しする
 *   - 長すぎるものは切り詰め、元ファイルの場所を書く（読む側の予算は有限）
 *   - 前回との差分（スコアと git diff）を必ず載せる
 */

const MAX_LINES = 120;

const readOr = (path: string, fallback: string): string =>
  existsSync(path) ? readFileSync(path, "utf8").trimEnd() : fallback;

const clip = (text: string, source: string, max = MAX_LINES): string => {
  const lines = text.split("\n");
  if (lines.length <= max) return text;
  return `${lines.slice(0, max).join("\n")}\n…（${lines.length - max} 行省略。全文は ${source}）`;
};

const fence = (text: string, lang = ""): string => `\`\`\`${lang}\n${text}\n\`\`\``;

const signed = (n: number): string => (n > 0 ? `+${n}` : String(n));

const ratio = (value: number, base: number): string =>
  base === 0 ? "—" : `${((value / base - 1) * 100).toFixed(1)}%`;

// ---------------------------------------------------------------------------
// vmstat の要約
// ---------------------------------------------------------------------------

interface ResourceSummary {
  host: string;
  nproc: number | undefined;
  samples: number;
  cpuBusyAvg: number;
  cpuBusyMax: number;
  ioWaitAvg: number;
  runQueueMax: number;
  freeMinMb: number;
  problem: string | undefined;
}

type VmstatResult =
  | { kind: "ok"; summary: ResourceSummary }
  /** ベンチが短すぎて、起動からの平均しか取れなかった */
  | { kind: "insufficient" }
  /** vmstat の出力として読めなかった（未取得も含む） */
  | { kind: "unparsable" };

/**
 * vmstat の出力を読んで「どのホストが飽和したか」に答えられる形にする。
 *
 * 最初のデータ行は起動からの平均なので捨てる（ベンチ中の値ではない）。
 * 列位置は環境で変わるのでヘッダ行から引く。
 */
const summarizeVmstat = (host: string, text: string): VmstatResult => {
  const lines = text.split("\n");
  const nproc = Number(/nproc=(\d+)/.exec(text)?.[1] ?? "");

  let columns: string[] | undefined;
  const rows: number[][] = [];
  for (const line of lines) {
    const tokens = line.trim().split(/\s+/);
    if (columns === undefined) {
      if (tokens.includes("r") && tokens.includes("id") && tokens.includes("wa")) columns = tokens;
      continue;
    }
    if (tokens.length < columns.length) continue;
    if (!tokens.every((t) => /^-?\d+$/.test(t))) continue;
    rows.push(tokens.map(Number));
  }
  if (columns === undefined || rows.length === 0) return { kind: "unparsable" };
  if (rows.length === 1) return { kind: "insufficient" };

  const at = (row: number[], name: string): number => row[columns.indexOf(name)] ?? 0;
  // 1 行目は起動からの累積平均
  const samples = rows.slice(1);

  let busySum = 0;
  let busyMax = 0;
  let waitSum = 0;
  let queueMax = 0;
  let freeMin = Number.POSITIVE_INFINITY;
  for (const row of samples) {
    const busy = 100 - at(row, "id");
    busySum += busy;
    if (busy > busyMax) busyMax = busy;
    waitSum += at(row, "wa");
    queueMax = Math.max(queueMax, at(row, "r"));
    freeMin = Math.min(freeMin, at(row, "free"));
  }

  const cpuBusyAvg = busySum / samples.length;
  const ioWaitAvg = waitSum / samples.length;
  const runQueueMax = queueMax;

  let problem: string | undefined;
  if (cpuBusyAvg >= 80) problem = "CPU が飽和";
  else if (ioWaitAvg >= 20) problem = "IO 待ちが重い";
  else if (Number.isFinite(nproc) && runQueueMax > nproc * 2) problem = "run queue が積もっている";

  return {
    kind: "ok",
    summary: {
      host,
      nproc: Number.isFinite(nproc) ? nproc : undefined,
      samples: samples.length,
      cpuBusyAvg,
      cpuBusyMax: busyMax,
      ioWaitAvg,
      runQueueMax,
      freeMinMb: Number.isFinite(freeMin) ? Math.round(freeMin / 1024) : 0,
      problem,
    },
  };
};

// ---------------------------------------------------------------------------
// top の要約（プロセス単位の CPU）
// ---------------------------------------------------------------------------

interface ProcessCpu {
  command: string;
  /** フレーム平均の %CPU。100% = 1 コア */
  cpu: number;
}

/** 自分たちが走らせている計測プロセスは除く。 */
const OWN_PROCESSES = new Set(["top", "vmstat", "curl"]);

/**
 * top -b の出力から「どのプロセスが CPU を食っていたか」を出す。
 *
 * フレームは `top - ` で始まる。最初のフレームは起動からの累積に近い値なので捨てる。
 * 同じコマンド名のプロセス（nginx の worker 等）はフレームごとに合算してから平均する。
 * 列位置は環境で変わるのでヘッダ行（PID ... %CPU ... COMMAND）から引く。
 */
export const summarizeTop = (text: string, limit = 5): ProcessCpu[] | undefined => {
  const frames: Map<string, number>[] = [];
  let current: Map<string, number> | undefined;
  let cpuIndex = -1;
  let commandIndex = -1;

  for (const line of text.split("\n")) {
    if (line.startsWith("top - ")) {
      current = new Map();
      frames.push(current);
      cpuIndex = -1;
      continue;
    }
    const tokens = line.trim().split(/\s+/);
    if (tokens[0] === "PID" && tokens.includes("%CPU") && tokens.includes("COMMAND")) {
      cpuIndex = tokens.indexOf("%CPU");
      commandIndex = tokens.indexOf("COMMAND");
      continue;
    }
    if (current === undefined || cpuIndex < 0 || !/^\d+$/.test(tokens[0] ?? "")) continue;

    const cpu = Number((tokens[cpuIndex] ?? "").replace(",", "."));
    const command = tokens.slice(commandIndex).join(" ");
    if (!Number.isFinite(cpu) || command === "" || OWN_PROCESSES.has(command)) continue;
    current.set(command, (current.get(command) ?? 0) + cpu);
  }

  const samples = frames.slice(1);
  if (samples.length === 0) return undefined;

  const totals = new Map<string, number>();
  for (const frame of samples) {
    for (const [command, cpu] of frame) totals.set(command, (totals.get(command) ?? 0) + cpu);
  }
  return [...totals]
    .map(([command, total]) => ({ command, cpu: total / samples.length }))
    .filter((p) => p.cpu >= 1)
    .sort((a, b) => b.cpu - a.cpu)
    .slice(0, limit);
};

const topByHost = (resourceDir: string): Map<string, ProcessCpu[] | undefined> => {
  const result = new Map<string, ProcessCpu[] | undefined>();
  const topDir = join(resourceDir, "top");
  if (!existsSync(topDir)) return result;
  for (const file of readdirSync(topDir).sort()) {
    if (!file.endsWith(".txt")) continue;
    result.set(file.replace(/\.txt$/, ""), summarizeTop(readFileSync(join(topDir, file), "utf8")));
  }
  return result;
};

const formatProcesses = (processes: ProcessCpu[]): string =>
  processes.map((p) => `${p.command} ${p.cpu.toFixed(0)}%`).join(" / ");

const resourceSection = (dir: string): string => {
  const resourceDir = join(dir, "resource");
  if (!existsSync(resourceDir)) return "未取得: リソース計測の結果がありません。";

  const processes = topByHost(resourceDir);
  const summaries: ResourceSummary[] = [];
  const insufficient: string[] = [];
  const unparsed: string[] = [];
  for (const file of readdirSync(resourceDir).sort()) {
    if (!file.endsWith(".txt")) continue;
    const host = file.replace(/\.txt$/, "");
    const result = summarizeVmstat(host, readFileSync(join(resourceDir, file), "utf8"));
    if (result.kind === "ok") summaries.push(result.summary);
    else if (result.kind === "insufficient") insufficient.push(host);
    else unparsed.push(host);
  }

  const parts: string[] = [];
  if (summaries.length > 0) {
    parts.push("| host | CPU 使用率 平均 / 最大 | IO 待ち 平均 | run queue 最大 | free 最小 | |");
    parts.push("|---|---|---|---|---|---|");
    for (const s of summaries) {
      const cpus = s.nproc === undefined ? "" : ` (${s.nproc} cpu)`;
      parts.push(
        `| ${s.host}${cpus} | ${s.cpuBusyAvg.toFixed(1)}% / ${s.cpuBusyMax.toFixed(0)}% ` +
          `| ${s.ioWaitAvg.toFixed(1)}% | ${s.runQueueMax} | ${s.freeMinMb} MB ` +
          `| ${s.problem ?? ""} |`,
      );
    }

    const hot = summaries.filter((s) => s.problem !== undefined);
    parts.push("");
    if (hot.length === 0) {
      parts.push(
        "どのホストも飽和していない。**ボトルネックはアプリのロジックか DB のクエリ側**にある可能性が高い。",
      );
    } else {
      // 飽和したホストで一番 CPU を食っていたプロセスまで名指しする
      const culprit = (host: string): string => {
        const first = processes.get(host)?.[0];
        return first === undefined ? "" : `、主に ${first.command}`;
      };
      parts.push(
        `**${hot.map((s) => `${s.host}（${s.problem}${culprit(s.host)}）`).join("、")}**。ここから手を付ける。`,
      );
    }
  }

  if (processes.size > 0) {
    parts.push("", "**CPU を食っているプロセス**（ベンチ中の平均 %CPU。100% = 1 コア）", "");
    parts.push("| host | 上位プロセス |", "|---|---|");
    for (const [host, list] of processes) {
      const cell =
        list === undefined
          ? "未取得（top のサンプルが足りない）"
          : list.length === 0
            ? "目立つプロセス無し"
            : formatProcesses(list);
      parts.push(`| ${host} | ${cell} |`);
    }
  }
  if (insufficient.length > 0) {
    parts.push(
      "",
      `サンプルが足りず要約できず: ${insufficient.join(", ")}。` +
        "ベンチが短すぎて vmstat が 1 行しか取れなかった（起動からの平均のみ）。" +
        "生の値は resource/<host>.txt にある。",
    );
  }
  if (unparsed.length > 0) {
    parts.push("", `未取得または vmstat として読めず: ${unparsed.join(", ")}`);
  }
  return parts.join("\n");
};

// ---------------------------------------------------------------------------
// 本体
// ---------------------------------------------------------------------------

export interface ReportContext {
  cfg: Config;
  meta: RunMeta;
  prev: HistoryRow | undefined;
  best: HistoryRow | undefined;
  /** status が pass 以外のときの理由 */
  failure: string | undefined;
}

const summarySection = (ctx: ReportContext): string => {
  const { meta, prev, best } = ctx;
  const started = Date.parse(meta.startedAt);
  const finished = Date.parse(meta.finishedAt);
  const elapsed =
    Number.isFinite(started) && Number.isFinite(finished)
      ? `${((finished - started) / 1000).toFixed(0)}s`
      : "—";

  const rows = [
    ["score", meta.status === "pass" ? String(meta.score) : `— (${meta.status})`],
    ["status", meta.status],
    ["ref", meta.ref],
    ["sha", meta.sha.slice(0, 12)],
    ["所要時間", elapsed],
    [
      "前回比",
      prev === undefined
        ? "—（初回）"
        : `${signed(meta.score - prev.score)} (${ratio(meta.score, prev.score)}) vs ${prev.runId}`,
    ],
    [
      "ベスト比",
      best === undefined
        ? "—（まだ pass したランが無い）"
        : `${signed(meta.score - best.score)} (${ratio(meta.score, best.score)}) vs ${best.runId} = ${best.score}` +
          (meta.status === "pass" && meta.score > best.score ? " **ベスト更新**" : ""),
    ],
    ["note", meta.note === "" ? "—" : meta.note],
  ];

  const table = ["| | |", "|---|---|", ...rows.map(([k, v]) => `| ${k} | ${v} |`)].join("\n");

  if (ctx.failure === undefined) return table;
  return `${table}\n\n**失敗の理由**\n\n${fence(ctx.failure)}`;
};

const diffSection = async (ctx: ReportContext): Promise<string> => {
  const { cfg, meta, prev } = ctx;
  if (prev === undefined) return "前回のランが無い（これが最初のラン）。";

  const scoreLine =
    `スコア: ${prev.score} → ${meta.status === "pass" ? meta.score : `— (${meta.status})`}` +
    (meta.status === "pass" ? ` (${signed(meta.score - prev.score)})` : "");

  if (prev.sha === meta.sha) {
    return `${scoreLine}\n\nコードの差分は無い（同じ sha ${meta.sha.slice(0, 12)}）。\nスコアの違いはブレか、サーバ側の設定変更による。`;
  }

  const result = await local(
    "git",
    ["-C", cfg.app.repoLocal, "diff", "--stat", prev.sha, meta.sha],
    { allowFailure: true },
  );
  const stat =
    result.exitCode === 0 && result.stdout.trim() !== ""
      ? fence(clip(result.stdout.trimEnd(), "git diff", 60))
      : `差分を取れなかった（${prev.sha.slice(0, 7)}..${meta.sha.slice(0, 7)}）。`;

  return `${scoreLine}\n\n\`${prev.sha.slice(0, 7)}..${meta.sha.slice(0, 7)}\`\n\n${stat}`;
};

/**
 * 計測ツールの出力は、要約できれば列を絞った表に、できなければ（未取得の理由など）元のまま載せる。
 * 要約は src/lib/digest.ts。生の出力は元ファイルにある。
 */
const digestSection = (dir: string, file: string, summarize: (text: string) => string | null): string => {
  const raw = readOr(join(dir, file), "未取得");
  const summary = summarize(raw);
  if (summary === null) return fence(clip(raw, file));
  return `${fence(clip(summary, file))}\n\n（生の出力は ${file}）`;
};

const alpSection = (dir: string, cfg: Config): string => {
  const body = digestSection(dir, "alp.txt", (text) => summarizeAlp(text, cfg.metrics.alpLimit));
  // 正規化しないと /api/user/1, /api/user/2 ... が別の行になり、上位が意味を持たない
  if (cfg.metrics.alpMatches.length > 0) return body;
  return `**注意: metrics.alpMatches が空で URI が正規化されていない。** ID を含むパスが別々の行に割れているので、集計は当てにならない。\n\n${body}`;
};

const ptQuerySection = (dir: string): string => digestSection(dir, "ptquery.txt", summarizePtQuery);

/** report.md を書く。見出しは docs/DESIGN.md §4 の通りで固定。 */
export const writeReport = async (dir: string, ctx: ReportContext): Promise<void> => {
  const benchLog = readOr(join(dir, "bench.log"), "未取得: ベンチを起動できなかった。");
  // pass 以外なら全文。原因はたいてい末尾ではなく途中にある。
  const bench = ctx.meta.status === "pass" ? clip(benchLog, "bench.log") : benchLog;

  const body = [
    `# run ${ctx.meta.runId}`,
    "",
    "## Summary",
    "",
    summarySection(ctx),
    "",
    "## Bench",
    "",
    fence(bench),
    "",
    "## Access log (alp)",
    "",
    alpSection(dir, ctx.cfg),
    "",
    "## Slow query (pt-query-digest)",
    "",
    ptQuerySection(dir),
    "",
    "## CPU profile (pprof)",
    "",
    fence(clip(readOr(join(dir, "pprof.txt"), "未取得"), "pprof.txt")),
    "",
    "## Resources",
    "",
    resourceSection(dir),
    "",
    "## Diff from previous run",
    "",
    await diffSection(ctx),
    "",
  ].join("\n");

  writeFileSync(join(dir, "report.md"), body);
};
