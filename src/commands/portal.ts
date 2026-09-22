import { loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { notImplemented, usageError } from "../lib/exit.js";
import { log } from "../lib/log.js";

const SUBCOMMANDS = ["whoami", "enqueue", "status", "wait", "score"] as const;
type Sub = (typeof SUBCOMMANDS)[number];

export const command: Command = {
  summary: "ポータルを Cookie 付きの HTTP で操作する",
  usage: `usage: isuh portal <subcommand> [args]

  whoami            Cookie が有効か確認する（認証必須ページを引けるか）
  enqueue           ベンチを登録し、ジョブ ID を stdout に出す
  status <job-id>   ジョブ状態を出す（pending / running / done / failed）
  wait <job-id>     完了まで待ち、結果 JSON を stdout に出す
  score             最新の確定スコアを stdout に出す

ポータル (https://portal.isucon.net/) に公開 API は無く、ブラウザ操作が前提。
認証は /auth/login/ からの GitHub 認証で、Django のセッション Cookie 方式
（レスポンスに vary: Cookie）。ブラウザでログインした後の sessionid と
csrftoken を portal.cookieFile に置いて使う。
POST には CSRF トークンが必要な想定（X-CSRFToken ヘッダ、または
csrfmiddlewaretoken フォーム値）。

ベンチ用のエンドポイントは当日にならないと分からない。判明したら
config.ts の portal.enqueuePath / portal.jobPath を埋めるだけで済むようにする。
事前に検証できるのは whoami まで。`,

  async run(argv) {
    const sub = argv[0];
    if (sub === undefined) throw usageError("サブコマンドが必要です");
    if (!SUBCOMMANDS.includes(sub as Sub)) throw usageError(`不明なサブコマンド: ${sub}`);
    const cfg = await loadConfig();

    if (sub === "status" || sub === "wait") {
      const jobId = argv[1];
      if (jobId === undefined) throw usageError(`${sub} には job-id が必要です`);
      log.info(`portal ${sub} job=${jobId}`);
    } else {
      log.info(`portal ${sub} base=${cfg.portal.base}`);
    }
    notImplemented(3);
  },
};
