import { loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { notImplemented, usageError } from "../lib/exit.js";
import { resolveSha } from "../lib/git.js";
import { log } from "../lib/log.js";
import { newRunId } from "../lib/paths.js";

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
    // 実装時は try-finally で必ず lock.release() すること（docs/DESIGN.md §7）
    notImplemented(2);
  },
};
