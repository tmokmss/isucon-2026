import { loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { notImplemented, usageError } from "../lib/exit.js";
import { log } from "../lib/log.js";

export const command: Command = {
  summary: "直前の good revision へ戻す",
  usage: `usage: isuh rollback [<sha>]

指定した sha、または runs/history.tsv から決めた直前の good revision へ
戻してデプロイする。

ISUCON のスコアは数 % 〜 10 % 程度は平気でブレる。ベンチ 1 回で改悪と
断定しないため、判定には bench.scoreRegressionRatio を使う。
bench は bench.autoRollback が true のときにこれを自動で呼ぶ。

戻したことは stderr のログと report.md に残すこと。`,

  async run(argv) {
    let sha = "";
    for (const arg of argv) {
      if (arg.startsWith("-")) throw usageError(`不明なオプション: ${arg}`);
      sha = arg;
    }
    const cfg = await loadConfig();
    log.info(
      `rollback target=${sha || "<直前の good revision>"} ratio=${cfg.bench.scoreRegressionRatio}`,
    );
    notImplemented(6);
  },
};
