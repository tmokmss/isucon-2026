import { writeFileSync } from "node:fs";
import { join } from "node:path";
import { type Config, loadConfig } from "../config.js";
import { type ResourceSampler, collectMetrics, resetLogs, startResourceSampling } from "../lib/collect.js";
import type { Command } from "../lib/command.js";
import { EXIT, HarnessError, NotImplementedError, usageError } from "../lib/exit.js";
import { resolveSha } from "../lib/git.js";
import { acquireLock } from "../lib/lock.js";
import { log, out } from "../lib/log.js";
import { newRunId, runDir } from "../lib/paths.js";
import { asJson, pick } from "../lib/portal.js";
import { command as checkCommand } from "./check.js";
import { deploy } from "./deploy.js";
import { SCORE_KEYS, STATUS_KEYS, enqueueBench, isPassStatus, waitBench } from "./portal.js";

interface BenchOutcome {
  jobId: string;
  /** ポータルが返した生の状態 */
  rawStatus: string;
  /** docs/DESIGN.md §4 の status */
  status: "pass" | "fail";
  score: number;
}

/**
 * ベンチを起動して完了まで待つ。
 *
 * ポータルのエンドポイントは当日まで分からないので、パスは config.portal に集約し、
 * ここでは「登録 → 完了待ち → 生出力を bench.log に落とす」だけを行う。
 */
const runBenchmark = async (cfg: Config, dir: string): Promise<BenchOutcome> => {
  const jobId = await enqueueBench(cfg);
  const response = await waitBench(cfg, jobId);
  writeFileSync(join(dir, "bench.log"), response.body.endsWith("\n") ? response.body : `${response.body}\n`);

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
   6. portal でベンチを登録 → 完了待ち
   7. 計測停止・ログ収集（alp / pt-query-digest / pprof / リソース）
   8. runs/<run-id>/ に report.md と meta.json を生成、history.tsv に追記
   9. スコア判定（bench.autoRollback なら低下時に戻す）
  10. ロック解放（try-finally で異常終了時も必ず実行）

ロックは ssh の子プロセスが握っている。Ctrl-C でも異常終了でも、プロセスが
死ねば接続が切れてリモートの flock が解放されるので、ロックは残らない。

終了コード: 0 / 10 ロック / 20 デプロイ / 30 check / 40 ベンチ / 50 スコア低下`,

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
    const runId = newRunId(sha);
    log.info(
      `bench runId=${runId} ref=${ref} sha=${sha.slice(0, 7)} skipCheck=${skipCheck} note='${note}'`,
    );

    // ここから 10 までが 1 つのクリティカルセクション。
    const lock = await acquireLock(`bench ${runId} ref=${ref}`, cfg);
    const dir = runDir(runId);
    let sampler: ResourceSampler | undefined;

    try {
      await deploy(cfg, { ref: sha });
      if (skipCheck) {
        log.warn("check を省略しました（--skip-check）");
      } else {
        await runCheck();
      }

      await resetLogs(cfg);
      sampler = await startResourceSampling(cfg, dir);

      const outcome = await runBenchmark(cfg, dir);
      log.info(
        `ベンチ完了 job=${outcome.jobId} status=${outcome.rawStatus} score=${outcome.score}`,
      );
      if (outcome.status !== "pass") {
        throw new HarnessError(
          EXIT.BENCH,
          `ベンチが失敗しました (status=${outcome.rawStatus})\n  詳細: ${join(dir, "bench.log")}`,
        );
      }
      out(String(outcome.score));

      // 8. レポート生成と台帳追記は #4 / #5。
      // 9. スコア判定とロールバックは #5 / #6。
    } finally {
      // 失敗しても計測は回収する。落ちたランのログこそ次の手がかりになる。
      if (sampler !== undefined) {
        try {
          await sampler.stop();
          await collectMetrics(cfg, dir);
        } catch (error) {
          log.warn(`計測の回収に失敗しました: ${error instanceof Error ? error.message : String(error)}`);
        }
        log.info(`計測の生データ: ${dir}`);
        log.info("report.md と history.tsv の生成は #4 / #5 で実装する");
      }
      try {
        await lock.release();
      } catch (error) {
        log.error(`ロックの解放に失敗しました: ${error instanceof Error ? error.message : String(error)}`);
      }
    }
  },
};
