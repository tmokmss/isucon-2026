import { loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { notImplemented, usageError } from "../lib/exit.js";
import { log } from "../lib/log.js";

export const command: Command = {
  summary: "ベンチを使わない検証ゲート",
  usage: `usage: isuh check

ベンチを使わずに「壊れていないこと」だけを数秒で確認する。
bench が deploy の直後に自動実行し、落ちた場合はベンチを起動しない。
ベンチは直列でしか回せないため、正しさの確認にベンチ枠を使わないのが狙い。

確認すること:
  - app.initializePath を叩いて成功すること
  - 主要エンドポイントが 200 を返すこと
  - レスポンスの「構造」が既知の形と一致すること（値ではなくキーの形）

レスポンス構造の変更はレギュレーション違反なので、ここが主な防波堤になる。

終了コード: 0 成功 / 30 検証失敗 / 65 設定不備`,

  async run(argv) {
    if (argv.length > 0) throw usageError(`引数は取らない: ${argv.join(" ")}`);
    const cfg = await loadConfig();
    log.info(`check target=${cfg.app.baseUrl}`);
    notImplemented(12);
  },
};
