import { loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { notImplemented, usageError } from "../lib/exit.js";
import { log } from "../lib/log.js";

export const command: Command = {
  summary: "再起動試験（失格回避）",
  usage: `usage: isuh verify-restart [--times <n>]

hosts.all を「ランダム順」で再起動し、全台の起動を待ってから
initialize を叩かずにベンチを実行して、直前スコアと比較する。

  --times <n>   繰り返し回数（既定: 1）。順序依存の事故を炙り出すため複数回まわす

レギュレーションより:
  - 主催者がインスタンスを再起動する順序は保証しない
  - 正常にベンチが完了しない、または競技中のスコアに近い結果が再現されない場合は失格
  - ベンチで書き込まれたデータは再起動後にも取得できること

あわせて確認すること:
  - systemd の enable 漏れ（systemctl is-enabled）
  - 再起動後にデータが残っていること

オンメモリ化・キャッシュ戦略はここで死ぬ。マージ前に必ず通す。

終了コード: 0 成功 / 40 ベンチ失敗 / 50 スコア低下`,

  async run(argv) {
    let times = 1;
    for (let i = 0; i < argv.length; i++) {
      const arg = argv[i]!;
      if (arg === "--times") {
        const value = argv[++i];
        if (value === undefined) throw usageError("--times に値がありません");
        times = Number.parseInt(value, 10);
        if (Number.isNaN(times) || times < 1) throw usageError(`--times は 1 以上の数値: ${value}`);
      } else {
        throw usageError(`不明な引数: ${arg}`);
      }
    }
    const cfg = await loadConfig();
    log.info(`verify-restart hosts=[${cfg.hosts.all.join(" ")}] times=${times}`);
    notImplemented(7);
  },
};
