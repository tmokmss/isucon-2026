import { writeFileSync } from "node:fs";
import { join } from "node:path";
import { execa } from "execa";
import { type Config, loadConfig } from "../config.js";
import {
  type ResourceSampler,
  collectMetrics,
  resetLogs,
  startResourceSampling,
} from "../lib/collect.js";
import type { Command } from "../lib/command.js";
import { EXIT, HarnessError, NotImplementedError, usageError } from "../lib/exit.js";
import { resolveSha } from "../lib/git.js";
import { acquireLock } from "../lib/lock.js";
import { log, out } from "../lib/log.js";
import { newRunId, runDir } from "../lib/paths.js";
import { asJson, pick } from "../lib/portal.js";
import { writeReport } from "../lib/report.js";
import { sshOpts } from "../lib/ssh.js";
import {
  type RunMeta,
  appendHistory,
  bestRun,
  historyRowOf,
  lastRun,
  localIso,
  readHistory,
  writeMeta,
} from "../lib/run.js";
import { command as checkCommand } from "./check.js";
import { deploy } from "./deploy.js";
import { SCORE_KEYS, STATUS_KEYS, enqueueBench, isPassStatus, waitBench } from "./portal.js";

interface BenchOutcome {
  /** ポータルのジョブ ID。local では "local" */
  jobId: string;
  /** ポータルが返した生の状態（local では pass / fail / exit=<n>） */
  rawStatus: string;
  /** docs/DESIGN.md §4 の status */
  status: "pass" | "fail";
  score: number;
}

const withNewline = (text: string): string => (text.endsWith("\n") ? text : `${text}\n`);

/**
 * ポータルでベンチを起動して完了まで待つ。
 *
 * ポータルのエンドポイントは当日まで分からないので、パスは config.portal に集約し、
 * ここでは「登録 → 完了待ち → 生出力を bench.log に落とす」だけを行う。
 */
const runPortalBenchmark = async (cfg: Config, dir: string): Promise<BenchOutcome> => {
  const jobId = await enqueueBench(cfg);
  const response = await waitBench(cfg, jobId);
  writeFileSync(join(dir, "bench.log"), withNewline(response.body));

  const json = asJson(response);
  const rawStatus = (pick(json, STATUS_KEYS) ?? "unknown").toLowerCase();
  const parsed = Number(pick(json, SCORE_KEYS) ?? "0");
  return {
    jobId,
    rawStatus,
    status: isPassStatus(rawStatus) ? "pass" : "fail",
    score: Number.isFinite(parsed) ? parsed : 0,
  };
};

/**
 * ベンチマーカーを ssh で直接実行して、出力からスコアと成否を読む（過去問のリハーサル用）。
 *
 * 出力は stdout と stderr を混ぜた順のまま bench.log に落とす。
 * スコアを読めない・passPattern が当たらない・終了コードが 0 でない、のどれかなら fail。
 */
const runLocalBenchmark = async (cfg: Config, dir: string): Promise<BenchOutcome> => {
  const { host, command, scorePattern, passPattern } = cfg.bench.local;
  log.info(`ベンチマーカーを実行します host=${host}`);
  log.debug(`ssh ${host} -- ${command}`);
  const result = await execa("ssh", [...sshOpts(), host, command], {
    reject: false,
    all: true,
    timeout: cfg.bench.timeoutSec * 1000,
  });
  const output = result.all ?? "";
  writeFileSync(join(dir, "bench.log"), withNewline(output));

  if (result.timedOut) {
    throw new HarnessError(EXIT.BENCH, `ベンチが ${cfg.bench.timeoutSec}s で終わりませんでした`);
  }
  const exitCode = result.exitCode ?? -1;
  const scoreText = new RegExp(scorePattern, "m").exec(output)?.[1];
  const parsed = Number(scoreText ?? Number.NaN);
  const passed = new RegExp(passPattern, "m").test(output);
  const readable = Number.isFinite(parsed);
  if (!readable) log.warn(`ベンチの出力からスコアを読めません（bench.local.scorePattern を確認）`);

  const pass = exitCode === 0 && passed && readable;
  return {
    jobId: "local",
    rawStatus: exitCode !== 0 ? `exit=${exitCode}` : pass ? "pass" : "fail",
    status: pass ? "pass" : "fail",
    score: readable ? parsed : 0,
  };
};

/** ベンチを起動して完了まで待つ。起動方法は bench.runner で選ぶ。 */
const runBenchmark = (cfg: Config, dir: string): Promise<BenchOutcome> =>
  cfg.bench.runner === "local" ? runLocalBenchmark(cfg, dir) : runPortalBenchmark(cfg, dir);

/**
 * check を挟む。deploy の「後」に走らせるのは、新しいコードが載った状態を検証するため。
 * 落ちたらベンチを起動せず、貴重なベンチ枠を消費しない。
 */
const runCheck = async (): Promise<void> => {
  try {
    await checkCommand.run([]);
  } catch (error) {
    if (error instanceof NotImplementedError) {
      log.warn("check は未実装のため飛ばしました（実装は #12）。構造を壊していないか自分で確認すること");
      return;
    }
    throw error;
  }
};

export const command: Command = {
  summary: "ロック取得からレポート生成までの通し実行",
  usage: `usage: isuh bench [<git-ref>] [--skip-check] [--note <text>]

デプロイからレポート生成までを 1 つのクリティカルセクションで実行する。
並列に走る他の作業が計測を汚染しないよう、全体をロックで保護する。

  <git-ref>       デプロイする ref（既定: HEAD、app.repoLocal で解決）
  --skip-check    check を省略する（非推奨）
  --note <text>   history.tsv に残すメモ

流れ:
   1. ロック取得（hosts.primary 上の flock を ssh の子プロセスで保持）
   2. git ref を sha に解決
   3. deploy <sha>
   4. check                        ← deploy の後。落ちたらベンチを起動しない
   5. ログのリセット + リソース計測開始
   6. ベンチを起動して完了待ち（bench.runner が portal ならポータルに登録、
      local なら bench.local.host でベンチマーカーを ssh で直接実行）
   7. 計測停止・ログ収集（alp / pt-query-digest / pprof / リソース）
   8. runs/<run-id>/ に report.md と meta.json を生成、history.tsv に追記
   9. （スコアの良し悪しは判定しない。report.md の前回比・ベスト比を読む）
  10. ロック解放（try-finally で異常終了時も必ず実行）

途中で失敗しても 7〜10 は必ず走る。落ちたランのログこそ次の手がかりになる。

ロックは ssh の子プロセスが握っている。Ctrl-C でも異常終了でも、プロセスが
死ねば接続が切れてリモートの flock が解放されるので、ロックは残らない。

stdout にはスコアだけを出す。読むべきものは runs/<run-id>/report.md。

終了コード: 0 / 10 ロック / 20 デプロイ / 30 check / 40 ベンチ`,

  async run(argv) {
    let ref = "HEAD";
    let skipCheck = false;
    let note = "";
    for (let i = 0; i < argv.length; i++) {
      const arg = argv[i]!;
      if (arg === "--skip-check") {
        skipCheck = true;
      } else if (arg === "--note") {
        const value = argv[++i];
        if (value === undefined) throw usageError("--note に値がありません");
        note = value;
      } else if (arg.startsWith("-")) {
        throw usageError(`不明なオプション: ${arg}`);
      } else {
        ref = arg;
      }
    }

    const cfg = await loadConfig();
    const sha = await resolveSha(cfg.app.repoLocal, ref);
    const startedAt = new Date();
    const runId = newRunId(sha, startedAt);
    log.info(
      `bench runId=${runId} ref=${ref} sha=${sha.slice(0, 7)} skipCheck=${skipCheck} note='${note}'`,
    );

    // このランより前の履歴。前回比・ベスト比・スコア判定に使う。
    const history = readHistory();
    const prev = lastRun(history);
    const best = bestRun(history);

    // ここから 10 までが 1 つのクリティカルセクション。
    const lock = await acquireLock(`bench ${runId} ref=${ref}`, cfg);
    const dir = runDir(runId);
    let sampler: ResourceSampler | undefined;
    let outcome: BenchOutcome | undefined;
    let failure: Error | undefined;

    try {
      await deploy(cfg, { ref: sha });
      if (skipCheck) {
        log.warn("check を省略しました（--skip-check）");
      } else {
        await runCheck();
      }

      await resetLogs(cfg);
      sampler = await startResourceSampling(cfg, dir);

      outcome = await runBenchmark(cfg, dir);
      log.info(`ベンチ完了 job=${outcome.jobId} status=${outcome.rawStatus} score=${outcome.score}`);
      if (outcome.status !== "pass") {
        throw new HarnessError(
          EXIT.BENCH,
          `ベンチが失敗しました (status=${outcome.rawStatus})\n  詳細: ${join(dir, "bench.log")}`,
        );
      }
    } catch (error) {
      failure = error instanceof Error ? error : new Error(String(error));
    } finally {
      // 失敗しても計測は回収し、レポートと台帳は必ず残す。
      if (sampler !== undefined) {
        try {
          await sampler.stop();
          await collectMetrics(cfg, dir);
        } catch (error) {
          log.warn(
            `計測の回収に失敗しました: ${error instanceof Error ? error.message : String(error)}`,
          );
        }
      }

      try {
        const status = outcome?.status ?? "error";
        const score = status === "pass" ? (outcome?.score ?? 0) : 0;
        const updatesBest = status === "pass" && (best === undefined || score > best.score);
        const meta: RunMeta = {
          runId,
          startedAt: localIso(startedAt),
          finishedAt: localIso(),
          ref,
          sha,
          score,
          status,
          note,
          benchJobId: outcome?.jobId ?? "",
          prevRunId: prev?.runId ?? null,
          bestRunId: updatesBest ? runId : (best?.runId ?? null),
        };
        writeMeta(dir, meta);
        await writeReport(dir, { cfg, meta, prev, best, failure: failure?.message });
        appendHistory(historyRowOf(meta));
        log.info(`レポート: ${join(dir, "report.md")}`);
        if (updatesBest) log.info(`ベスト更新 ${best === undefined ? "" : `${best.score} → `}${score}`);
      } catch (error) {
        log.warn(
          `レポートの生成に失敗しました: ${error instanceof Error ? error.message : String(error)}`,
        );
      }

      try {
        await lock.release();
      } catch (error) {
        log.error(
          `ロックの解放に失敗しました: ${error instanceof Error ? error.message : String(error)}`,
        );
      }
    }

    if (failure !== undefined) throw failure;

    // 良くなったか悪くなったかの判定はしない。前回比・ベスト比は report.md に載っているので、
    // 読む側（人かエージェント）が判断する。閾値で機械的に決めるとスコアのブレを誤判定する。
    out(String((outcome as BenchOutcome).score));
  },
};
