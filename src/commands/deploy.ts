import { loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { notImplemented, usageError } from "../lib/exit.js";
import { resolveSha } from "../lib/git.js";
import { log } from "../lib/log.js";

export const command: Command = {
  summary: "指定した git ref をビルドして全台へ配布・再起動する",
  usage: `usage: isuh deploy [<git-ref>] [--dry-run]

指定した git ref のアプリをビルドし、hosts.app へ配布して再起動する。
ref は app.repoLocal（手元のアプリリポジトリ）で解決する。既定は HEAD。

  <git-ref>   デプロイする ref
  --dry-run   何をするかを表示して実行しない

デプロイ後、各ホストに現在の sha を記録すること。
nginx.conf や my.cnf などアプリ以外の配布もこのコマンドで行う。

終了コード: 0 成功 / 20 デプロイ失敗 / 64 引数不正 / 65 設定不備`,

  async run(argv) {
    let ref = "HEAD";
    let dryRun = false;
    for (const arg of argv) {
      if (arg === "--dry-run") dryRun = true;
      else if (arg.startsWith("-")) throw usageError(`不明なオプション: ${arg}`);
      else ref = arg;
    }
    const cfg = await loadConfig();
    const sha = await resolveSha(cfg.app.repoLocal, ref);
    log.info(
      `deploy ref=${ref} sha=${sha.slice(0, 7)} hosts=[${cfg.hosts.app.join(" ")}] dryRun=${dryRun}`,
    );
    notImplemented(1);
  },
};
