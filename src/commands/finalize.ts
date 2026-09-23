import { loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { notImplemented, usageError } from "../lib/exit.js";
import { log } from "../lib/log.js";

export const command: Command = {
  summary: "終盤チェックリスト",
  usage: `usage: isuh finalize

競技終盤に 1 コマンドで確認する。「終盤 1 時間は改善禁止、finalize に充てる」
という運用とセットで使う。

確認すること:
  - 全ログの無効化: nginx access log / DB slow query log / アプリの debug logger / pprof
  - systemd の enable 確認（再起動後に全部上がるか）
  - デバッグ用に開けたポート（6060 等）が閉じているか
  - 不要になったデバッグコードが残っていないか

その後、最終ベンチ → verify-restart まで流す。

終了コード: 0 成功 / 40 ベンチ失敗`,

  async run(argv) {
    if (argv.length > 0) throw usageError(`引数は取らない: ${argv.join(" ")}`);
    const cfg = await loadConfig();
    log.info(`finalize hosts=[${cfg.hosts.all.join(" ")}]`);
    notImplemented(8);
  },
};
