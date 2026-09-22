import { loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { notImplemented, usageError } from "../lib/exit.js";
import { log } from "../lib/log.js";

export const command: Command = {
  summary: "全ホストのリソースを 1 画面で見る",
  usage: `usage: isuh monitor

hosts.all 全台の dstat / htop を tmux で 1 画面に並べる。
ベンチ実行中にどのホストが飽和しているかを見るのが目的。
凝ったダッシュボードは当日使わない可能性が高いので、まず tmux + dstat で十分。`,

  async run(argv) {
    if (argv.length > 0) throw usageError(`引数は取らない: ${argv.join(" ")}`);
    const cfg = await loadConfig();
    log.info(
      `monitor hosts=[${cfg.hosts.all.join(" ")}] interval=${cfg.metrics.resourceIntervalSec}s`,
    );
    notImplemented(17);
  },
};
