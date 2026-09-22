import { writeFileSync } from "node:fs";
import { join } from "node:path";
import { type Config, loadConfig } from "../config.js";
import {
  type ResourceSampler,
  collectMetrics,
  resetLogs,
  startResourceSampling,
} from "../lib/collect.js";
import type { Command } from "../lib/command.js";
import { EXIT, HarnessError, usageError } from "../lib/exit.js";
import { resolveSha } from "../lib/git.js";
import { acquireLock } from "../lib/lock.js";
import { log, out } from "../lib/log.js";
import { newRunId, runDir } from "../lib/paths.js";
import { asJson, pick } from "../lib/portal.js";
import { writeReport } from "../lib/report.js";
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
import { resolveRollbackTarget, rollbackTo } from "./rollback.js";
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
  writeFileSync(
    join(dir, "bench.log"),
    response.body.endsWith("\n") ? response.body : `${response.body}\n`,
  );

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
   9. スコア判定（bench.scoreRegressionRatio を下回れば exit 50）
  10. ロック解放（try-finally で異常終了時も必ず実行）

途中で失敗しても 7〜10 は必ず走る。落ちたランのログこそ次の手がかりになる。

ロックは ssh の子プロセスが握っている。Ctrl-C でも異常終了でも、プロセスが
死ねば接続が切れてリモートの flock が解放されるので、ロックは残らない。

stdout にはスコアだけを出す。読むべきものは runs/<run-id>/report.md。

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
    let rolledBackTo: string | null = null;

    try {
      await deploy(cfg, { ref: sha });
      if (skipCheck) {
        log.warn("check を省略しました（--skip-check）");
      } else {
        // check は deploy の「後」。新しいコードが載った状態を検証するため。
        // 落ちたらベンチを起動せず、貴重なベンチ枠を消費しない。
        await checkCommand.run([]);
      }

      await resetLogs(cfg);
      sampler = await startResourceSampling(cfg, dir);

      outcome = await runBenchmark(cfg, dir);
      log.info(`ベンチ完了 job=${outcome.jobId} status=${outcome.rawStatus} score=${outcome.score}`);

      // 9. スコア判定。ISUCON のスコアは数 % 〜 10 % 平気でブレるので、
      // ベスト × bench.scoreRegressionRatio を下回ったときだけ「低下」とする。
      const threshold = best === undefined ? 0 : best.score * cfg.bench.scoreRegressionRatio;
      const regressed =
        outcome.status === "pass" && best !== undefined && outcome.score < threshold;

      if (outcome.status !== "pass" || regressed) {
        rolledBackTo = await maybeRollback(cfg, sha);
      }

      if (outcome.status !== "pass") {
        throw new HarnessError(
          EXIT.BENCH,
          `ベンチが失敗しました (status=${outcome.rawStatus})\n  詳細: ${join(dir, "bench.log")}`,
        );
      }

      // 低下していてもスコアは出す。判定は終了コードで伝える。
      out(String(outcome.score));

      if (regressed) {
        throw new HarnessError(
          EXIT.REGRESSION,
          `スコアが低下しました: ${outcome.score} < ベスト ${best?.score} × ${cfg.bench.scoreRegressionRatio} = ${threshold.toFixed(0)}\n` +
            (rolledBackTo === null
              ? `  戻すなら: isuh rollback\n`
              : `  ${rolledBackTo.slice(0, 12)} へ戻しました\n`) +
            `  ブレの範囲かもしれません。大きい変更なら 2 回測って判断してください。`,
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
          rolledBackTo,
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
  },
};

/**
 * 失敗・低下を受けて戻す。ロックはすでに取れている（bench のクリティカルセクション内）。
 *
 * 戻したかどうかは bench.autoRollback で切り替える。序盤は手動（false）、
 * 終盤は自動（true）が使いやすい。戻せなかった場合もランは記録して先へ進む。
 */
const maybeRollback = async (cfg: Config, currentSha: string): Promise<string | null> => {
  if (!cfg.bench.autoRollback) {
    log.warn("bench.autoRollback が false のため戻していません。戻すなら: isuh rollback");
    return null;
  }
  try {
    const target = await resolveRollbackTarget(cfg, undefined);
    if (target.sha === currentSha) {
      log.warn("戻し先が今デプロイした revision と同じなので、戻しませんでした");
      return null;
    }
    await rollbackTo(cfg, target);
    return target.sha;
  } catch (error) {
    log.error(
      `ロールバックに失敗しました: ${error instanceof Error ? error.message : String(error)}`,
    );
    return null;
  }
};
