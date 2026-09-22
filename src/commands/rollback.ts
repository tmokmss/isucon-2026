import { type Config, loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { EXIT, HarnessError, usageError } from "../lib/exit.js";
import { resolveSha } from "../lib/git.js";
import { acquireLock } from "../lib/lock.js";
import { log, out } from "../lib/log.js";
import { type HistoryRow, lastGoodRun, readHistory } from "../lib/run.js";
import { deploy } from "./deploy.js";

export interface RollbackTarget {
  sha: string;
  /** どうやって決めたか。ログとレポートに出す */
  reason: string;
  /** 台帳から決めた場合の元ラン */
  run: HistoryRow | undefined;
}

/**
 * 戻し先を決める。
 *
 * 引数があればそれ、無ければ台帳の good revision。
 * good は「ベスト × bench.scoreRegressionRatio 以上のスコアを出した最も新しいラン」。
 * 単に「直前に pass したラン」にすると、**低下したラン自身も pass なので
 * そこへ戻してしまう**。ベストそのものへ戻さないのは、ベストが古すぎて
 * その後の良い変更まで巻き戻してしまうのを避けるため。
 */
export const resolveRollbackTarget = async (
  cfg: Config,
  ref: string | undefined,
): Promise<RollbackTarget> => {
  if (ref !== undefined && ref !== "") {
    return { sha: await resolveSha(cfg.app.repoLocal, ref), reason: `指定 (${ref})`, run: undefined };
  }

  const good = lastGoodRun(readHistory(), cfg.bench.scoreRegressionRatio);
  if (good === undefined) {
    throw new HarnessError(
      EXIT.USAGE,
      "戻し先が決められません。まだ pass したランが台帳にありません。\n" +
        "  戻したい revision を明示してください: isuh rollback <sha>",
    );
  }
  return {
    sha: good.sha,
    reason: `直前の good revision (${good.runId}, score=${good.score})`,
    run: good,
  };
};

/**
 * 戻す。**ロックは呼び出し側で取っていること。**
 * bench(#2) はクリティカルセクションの中からこれを呼ぶ。
 */
export const rollbackTo = async (cfg: Config, target: RollbackTarget): Promise<void> => {
  log.warn(`ロールバックします → ${target.sha.slice(0, 12)} : ${target.reason}`);
  await deploy(cfg, { ref: target.sha });
  log.warn(`ロールバック完了: ${target.sha.slice(0, 12)} が全台に載っています`);
};

export const command: Command = {
  summary: "直前の good revision へ戻す",
  usage: `usage: isuh rollback [<sha>]

指定した sha、または runs/history.tsv から決めた直前の good revision へ
戻してデプロイする。

  <sha>   戻し先。省略すると台帳の「直前に pass したラン」

ベストではなく直前の good を既定にするのは、ベストが古すぎて、その後の
良い変更まで巻き戻してしまうのを避けるため。ベストへ戻したいときは
isuh history で sha を見て明示する。

ベンチが走っている最中に戻すと計測が壊れるので、ベンチと同じロックを取る。
取れなければ exit 10 で、誰が保持しているかを出す。

ISUCON のスコアは数 % 〜 10 % 程度は平気でブレる。ベンチ 1 回で改悪と
断定しないため、判定には bench.scoreRegressionRatio を使う。
bench は bench.autoRollback が true のときにこれを自動で呼ぶ。

戻したことは stderr のログと、次に生成される report.md に残る。
stdout には戻した sha だけを出す。

終了コード: 0 / 10 ロック / 20 デプロイ失敗 / 64 引数不正 / 65 設定不備`,

  async run(argv) {
    let ref = "";
    for (const arg of argv) {
      if (arg.startsWith("-")) throw usageError(`不明なオプション: ${arg}`);
      ref = arg;
    }

    const cfg = await loadConfig();
    const target = await resolveRollbackTarget(cfg, ref);
    log.info(`rollback target=${target.sha.slice(0, 12)} (${target.reason})`);

    // ベンチ中に配布すると計測が壊れる。ベンチと同じロックを取る。
    const lock = await acquireLock(`rollback → ${target.sha.slice(0, 12)}`, cfg);
    try {
      await rollbackTo(cfg, target);
    } finally {
      await lock.release();
    }
    out(target.sha);
  },
};
