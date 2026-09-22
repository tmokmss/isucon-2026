import { existsSync } from "node:fs";
import type { Command } from "../lib/command.js";
import { notImplemented, usageError } from "../lib/exit.js";
import { log } from "../lib/log.js";
import { historyFile } from "../lib/paths.js";

export const command: Command = {
  summary: "スコア台帳の表示",
  usage: `usage: isuh history [-n <count>]

runs/history.tsv を読んでスコア台帳を表示する。ベスト更新を強調する。
「どの変更でスコアがどう動いたか」が一目で分かることが目的で、
rollback が戻し先を決めるための材料にもなる。

  -n <count>   表示する件数（既定: 全件）

カラム: run_id / started_at / sha / ref / score / status / note`,

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
    const file = historyFile();
    if (!existsSync(file)) {
      log.info(`まだ実行履歴がありません: ${file}`);
      return;
    }
    log.info(`history file=${file} count=${count}`);
    notImplemented(5);
  },
};
