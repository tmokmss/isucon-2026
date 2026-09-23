import { setTimeout as delay } from "node:timers/promises";
import { type Config, loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { EXIT, HarnessError, configError, usageError } from "../lib/exit.js";
import { log, out } from "../lib/log.js";
import {
  type PortalResponse,
  asJson,
  assertLoggedIn,
  pick,
  portalRequest,
  titleOf,
} from "../lib/portal.js";

const SUBCOMMANDS = ["whoami", "get", "post", "enqueue", "status", "wait", "score"] as const;
type Sub = (typeof SUBCOMMANDS)[number];

const benchError = (message: string): HarnessError => new HarnessError(EXIT.BENCH, message);

/** 当日埋める設定が空のときに、何をすればいいかまで書いて落ちる。 */
const requirePath = (value: string, name: string): string => {
  if (value !== "") return value;
  throw configError(
    `config.ts の portal.${name} が空です。当日、ポータルの該当エンドポイントを埋めてください。\n` +
      `  探すには: isuh portal get / でページを引いて、フォームの action や fetch 先を見る。`,
  );
};

const dump = (response: PortalResponse): void => {
  out(response.body);
};

/** ジョブ状態の候補キー。当日の JSON の形が分からないので順に見る。 */
export const STATUS_KEYS = ["status", "state", "job_status", "jobStatus"] as const;
export const JOB_ID_KEYS = ["job_id", "jobId", "id", "uuid"] as const;
export const SCORE_KEYS = ["score", "total_score", "totalScore"] as const;

const TERMINAL = new Set(["done", "finished", "completed", "success", "failed", "error", "aborted"]);

/** 完走してスコアが付いた状態か。 */
export const isPassStatus = (status: string): boolean =>
  ["done", "finished", "completed", "success", "passed"].includes(status.toLowerCase());

const jobPathFor = (cfg: Config, jobId: string): string =>
  requirePath(cfg.portal.jobPath, "jobPath").replace("{job}", encodeURIComponent(jobId));

const fetchStatus = async (cfg: Config, jobId: string): Promise<{ status: string; response: PortalResponse }> => {
  const response = await portalRequest(cfg, jobPathFor(cfg, jobId));
  assertLoggedIn(cfg, response);
  if (response.status >= 400) {
    throw benchError(`ジョブ状態を取得できません (HTTP ${response.status}) ${response.url}`);
  }
  const json = asJson(response);
  const status = pick(json, STATUS_KEYS);
  if (status === undefined) {
    throw benchError(
      `ジョブ状態を読み取れません。返ってきたものを見て portal.jobPath を直してください。\n` +
        `  ${response.url} (HTTP ${response.status}, ${response.contentType})\n` +
        `${response.body.slice(0, 800)}`,
    );
  }
  return { status: status.toLowerCase(), response };
};

/** ベンチを登録して job-id を返す。bench(#2) から呼ぶ。 */
export const enqueueBench = async (cfg: Config): Promise<string> => {
  const path = requirePath(cfg.portal.enqueuePath, "enqueuePath");
  const form: Record<string, string> = {};
  if (cfg.portal.teamId !== "") form["team"] = cfg.portal.teamId;
  const response = await portalRequest(cfg, path, { form });
  assertLoggedIn(cfg, response);
  if (response.status >= 400) {
    throw benchError(
      `ベンチを登録できません (HTTP ${response.status}) ${response.url}\n${response.body.slice(0, 800)}`,
    );
  }
  // JSON に job id があればそれ。無ければリダイレクト先の末尾を見る。
  const jobId =
    pick(asJson(response), JOB_ID_KEYS) ?? /\/(\w[\w-]*)\/?$/.exec(response.location ?? "")?.[1];
  if (jobId === undefined) {
    throw benchError(
      `ジョブ ID を読み取れません。返ってきたものを見て実装を合わせてください。\n` +
        `  HTTP ${response.status} ${response.contentType} Location=${response.location}\n` +
        `${response.body.slice(0, 800)}`,
    );
  }
  log.info(`ベンチを登録しました job=${jobId}`);
  return jobId;
};

/** 完了まで待って最後のレスポンスを返す。bench(#2) から呼ぶ。 */
export const waitBench = async (cfg: Config, jobId: string): Promise<PortalResponse> => {
  const deadline = Date.now() + cfg.bench.timeoutSec * 1000;
  let last = "";
  for (;;) {
    const { status, response } = await fetchStatus(cfg, jobId);
    if (status !== last) {
      log.info(`job=${jobId} ${status}`);
      last = status;
    }
    if (TERMINAL.has(status)) return response;
    if (Date.now() > deadline) {
      throw benchError(
        `ベンチの完了を待てませんでした (${cfg.bench.timeoutSec}s, 最後の状態: ${status})`,
      );
    }
    await delay(2000);
  }
};

export const command: Command = {
  summary: "ポータルを Cookie 付きの HTTP で操作する",
  usage: `usage: isuh portal <subcommand> [args]

  whoami            Cookie が有効か確認する（認証必須ページを引けるか）
  get <path>        任意のページを Cookie 付きで引いて stdout に出す
  post <path> [k=v ...]
                    任意のパスへ CSRF 付きで POST する
  enqueue           ベンチを登録し、ジョブ ID を stdout に出す
  status <job-id>   ジョブ状態を出す（pending / running / done / failed）
  wait <job-id>     完了まで待ち、結果 JSON を stdout に出す
  score             最新の確定スコアを stdout に出す

ポータル (https://portal.isucon.net/) に公開 API は無く、ブラウザ操作が前提。
認証は /auth/login/ からの GitHub 認証で、Django のセッション Cookie 方式
（レスポンスに vary: Cookie）。ブラウザでログインした後の sessionid と
csrftoken を portal.cookieFile に 1 行で置いて使う。

    sessionid=<値>; csrftoken=<値>

POST には CSRF が要る。Cookie の csrftoken を X-CSRFToken ヘッダに載せ、
Referer / Origin も付けて送る。

ベンチ用のエンドポイントは当日にならないと分からない。
まず get で探り、判明したら config.ts の portal.enqueuePath / jobPath /
scorePath を埋める。事前に検証できるのは whoami と get まで。

終了コード: 0 / 40 ポータルの応答が想定外 / 64 引数不正 / 65 Cookie 切れ・設定不備`,

  async run(argv) {
    const sub = argv[0];
    if (sub === undefined) throw usageError("サブコマンドが必要です");
    if (!SUBCOMMANDS.includes(sub as Sub)) throw usageError(`不明なサブコマンド: ${sub}`);

    // 引数の検証は設定の読み込みより先。
    if (sub === "status" || sub === "wait") {
      if (argv[1] === undefined) throw usageError(`${sub} には job-id が必要です`);
    }
    if (sub === "get" || sub === "post") {
      if (argv[1] === undefined) throw usageError(`${sub} にはパスが必要です`);
    }

    const cfg = await loadConfig();

    switch (sub) {
      case "whoami": {
        const path = cfg.portal.whoamiPath === "" ? "/" : cfg.portal.whoamiPath;
        const response = await portalRequest(cfg, path);
        assertLoggedIn(cfg, response);
        if (response.status >= 400) {
          throw benchError(`ポータルが ${response.status} を返しました: ${response.url}`);
        }
        const title = titleOf(response.body);
        log.info(`Cookie は有効です (HTTP ${response.status}) ${response.url}`);
        out(title ?? `ok ${response.status}`);
        return;
      }

      case "get": {
        const response = await portalRequest(cfg, argv[1]!);
        assertLoggedIn(cfg, response);
        log.info(`HTTP ${response.status} ${response.contentType} ${response.url}`);
        if (response.location !== null) log.info(`Location: ${response.location}`);
        dump(response);
        return;
      }

      case "post": {
        const form: Record<string, string> = {};
        for (const pair of argv.slice(2)) {
          const index = pair.indexOf("=");
          if (index <= 0) throw usageError(`key=value の形で渡してください: ${pair}`);
          form[pair.slice(0, index)] = pair.slice(index + 1);
        }
        const response = await portalRequest(cfg, argv[1]!, { form });
        assertLoggedIn(cfg, response);
        log.info(`HTTP ${response.status} ${response.contentType} ${response.url}`);
        if (response.location !== null) log.info(`Location: ${response.location}`);
        dump(response);
        return;
      }

      case "enqueue": {
        out(await enqueueBench(cfg));
        return;
      }

      case "status": {
        const { status } = await fetchStatus(cfg, argv[1]!);
        out(status);
        return;
      }

      case "wait": {
        out((await waitBench(cfg, argv[1]!)).body);
        return;
      }

      case "score": {
        const path = requirePath(cfg.portal.scorePath, "scorePath");
        const response = await portalRequest(cfg, path);
        assertLoggedIn(cfg, response);
        if (response.status >= 400) {
          throw benchError(`スコアを取得できません (HTTP ${response.status}) ${response.url}`);
        }
        const score = pick(asJson(response), SCORE_KEYS);
        if (score === undefined) {
          throw benchError(
            `スコアを読み取れません。返ってきたものを見て portal.scorePath を直してください。\n` +
              `  ${response.url} (HTTP ${response.status}, ${response.contentType})\n` +
              `${response.body.slice(0, 800)}`,
          );
        }
        out(score);
        return;
      }
    }
  },
};
