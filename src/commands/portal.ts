import { setTimeout as delay } from "node:timers/promises";
import { type Config, loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { EXIT, HarnessError, configError, usageError } from "../lib/exit.js";
import { log, out } from "../lib/log.js";
import {
  type PortalResponse,
  assertLoggedIn,
  isProto,
  loggedOutError,
  portalRequest,
} from "../lib/portal.js";
import { type Fields, bool, decode, dumpFields, encodeInt, int, msg, msgs, str } from "../lib/proto.js";

const SUBCOMMANDS = ["whoami", "get", "post", "enqueue", "status", "wait", "score"] as const;
type Sub = (typeof SUBCOMMANDS)[number];

const benchError = (message: string): HarnessError => new HarnessError(EXIT.BENCH, message);

/** 当日埋める設定が空のときに、何をすればいいかまで書いて落ちる。 */
const requirePath = (value: string, name: string): string => {
  if (value !== "") return value;
  throw configError(
    `config.ts の portal.${name} が空です。当日、ポータルの該当エンドポイントを埋めてください。\n` +
      `  探すには: isuh portal get <path> で API を引いて中身を見る。`,
  );
};

/** API を引いて protobuf を読む。ログイン切れ・エラー応答はここで落とす。 */
const fetchProto = async (
  cfg: Config,
  path: string,
  what: string,
  proto?: Uint8Array,
): Promise<Fields> => {
  const response = await portalRequest(cfg, path, proto === undefined ? {} : { proto });
  assertLoggedIn(response);
  if (response.status >= 400 || !isProto(response)) {
    throw benchError(
      `${what}に失敗しました (HTTP ${response.status}, ${response.contentType}) ${response.url}\n` +
        `${describe(response).slice(0, 800)}`,
    );
  }
  return decode(response.bytes);
};

/** エラー応答の中身。protobuf の Error（human_message 等）なら読めるように展開する。 */
const describe = (response: PortalResponse): string => {
  if (!/protobuf/i.test(response.contentType)) return response.body;
  try {
    return JSON.stringify(dumpFields(decode(response.bytes)));
  } catch {
    return response.body;
  }
};

/** isuxportal.proto.resources.BenchmarkJob.Status */
const JOB_STATUS = ["pending", "running", "errored", "cancelled", "finished"] as const;

export interface BenchJob {
  id: string;
  /** docs/DESIGN.md の isuh portal status の値（pending / running / done / failed） */
  status: "pending" | "running" | "done" | "failed";
  /** ポータルの BenchmarkJob.Status をそのまま小文字にしたもの */
  portalStatus: string;
  passed: boolean;
  score: number;
  reason: string;
  stdout: string;
  stderr: string;
}

/** {Get,Enqueue}BenchmarkJobResponse の job（field 1）を読む */
const jobOf = (response: Fields): BenchJob => {
  const job = msg(response, 1);
  const result = msg(job, 18);
  const execution = msg(result, 5);
  const portalStatus = JOB_STATUS[int(job, 4)] ?? `unknown(${int(job, 4)})`;
  const status =
    portalStatus === "pending" || portalStatus === "running"
      ? portalStatus
      : portalStatus === "finished" && bool(result, 2)
        ? "done"
        : "failed";
  return {
    id: String(int(job, 1)),
    status,
    portalStatus,
    passed: bool(result, 2),
    score: int(result, 3),
    reason: str(execution, 1),
    stdout: str(execution, 2),
    stderr: str(execution, 3),
  };
};

interface Session {
  team: string;
  contestant: string;
  instances: { id: number; number: number; publicIp: string; privateIp: string }[];
}

/** GetCurrentSessionResponse。/api/session は未ログインでも 200 を返すので contestant の有無で見る。 */
const fetchSession = async (cfg: Config): Promise<Session> => {
  const path = cfg.portal.whoamiPath === "" ? "/api/session" : cfg.portal.whoamiPath;
  const response = await portalRequest(cfg, path);
  assertLoggedIn(response);
  if (response.status >= 400 || !isProto(response)) {
    throw benchError(`ポータルが想定外の応答を返しました (HTTP ${response.status}, ${response.contentType}) ${response.url}`);
  }
  const session = decode(response.bytes);
  const contestant = msg(session, 2);
  if (contestant.size === 0) throw loggedOutError(response);
  return {
    team: str(msg(session, 1), 2),
    contestant: str(contestant, 3),
    instances: msgs(session, 5).map((instance) => ({
      id: int(instance, 7),
      number: int(instance, 3),
      publicIp: str(instance, 4),
      privateIp: str(instance, 5),
    })),
  };
};

const fetchJob = async (cfg: Config, jobId: string): Promise<BenchJob> => {
  const path = requirePath(cfg.portal.jobPath, "jobPath").replace("{job}", encodeURIComponent(jobId));
  return jobOf(await fetchProto(cfg, path, "ジョブ状態の取得"));
};

/**
 * ベンチを登録して job-id を返す。bench(#2) から呼ぶ。
 * 対象はポータルの画面と同じく、セッションの contestant_instances の先頭。
 */
export const enqueueBench = async (cfg: Config): Promise<string> => {
  const path = requirePath(cfg.portal.enqueuePath, "enqueuePath");
  const target = (await fetchSession(cfg)).instances[0];
  if (target === undefined) {
    throw benchError("ポータルにチームのインスタンスがありません（ベンチの対象が決まりません）");
  }
  // EnqueueBenchmarkJobRequest { int64 target_id = 1 }
  const job = jobOf(await fetchProto(cfg, path, "ベンチの登録", encodeInt(1, target.id)));
  log.info(`ベンチを登録しました job=${job.id} target=${target.publicIp}`);
  return job.id;
};

/** 完了まで待って最後のジョブを返す。bench(#2) から呼ぶ。 */
export const waitBench = async (cfg: Config, jobId: string): Promise<BenchJob> => {
  const deadline = Date.now() + cfg.bench.timeoutSec * 1000;
  let last = "";
  for (;;) {
    const job = await fetchJob(cfg, jobId);
    if (job.portalStatus !== last) {
      log.info(`job=${jobId} ${job.portalStatus}`);
      last = job.portalStatus;
    }
    if (job.status === "done" || job.status === "failed") return job;
    if (Date.now() > deadline) {
      throw benchError(
        `ベンチの完了を待てませんでした (${cfg.bench.timeoutSec}s, 最後の状態: ${job.portalStatus})`,
      );
    }
    await delay(2000);
  }
};

const dump = (response: PortalResponse): void => {
  out(/protobuf/i.test(response.contentType) ? JSON.stringify(dumpFields(decode(response.bytes)), null, 2) : response.body);
};

export const command: Command = {
  summary: "ポータルを Cookie 付きの HTTP で操作する",
  usage: `usage: isuh portal <subcommand> [args]

  whoami            Cookie が有効か確認し、チーム名と参加者名を出す
  get <path>        任意の API を Cookie 付きで引いて stdout に出す（protobuf はフィールド番号で展開）
  post <path> [k=v ...]
                    任意のパスへ CSRF 付きで POST する
  enqueue           ベンチを登録し、ジョブ ID を stdout に出す
  status <job-id>   ジョブ状態を出す（pending / running / done / failed）
  wait <job-id>     完了まで待ち、結果 JSON を stdout に出す
  score             最新の確定スコアを stdout に出す

ポータル (https://portal.isucon.net/) は isuxportal。公開 API は無く、
ブラウザの JS が叩く /api/... を protobuf で使う。認証は GitHub ログイン後の
セッション Cookie。ブラウザから isuxportal_session を portal.cookieFile に
1 行で置いて使う。

    isuxportal_session=<値>

POST の CSRF トークンはトップページの <meta name="csrf-token"> から取って
X-Csrf-Token ヘッダに載せる。ベンチの対象はポータルの画面と同じく、
チームのインスタンス一覧の先頭。

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
        const session = await fetchSession(cfg);
        log.info("Cookie は有効です");
        for (const instance of session.instances) {
          log.info(
            `  instance #${instance.number} id=${instance.id} public=${instance.publicIp} private=${instance.privateIp}`,
          );
        }
        out(`${session.team} / ${session.contestant}`);
        return;
      }
      case "get": {
        const response = await portalRequest(cfg, argv[1]!);
        assertLoggedIn(response);
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
        assertLoggedIn(response);
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
        out((await fetchJob(cfg, argv[1]!)).status);
        return;
      }
      case "wait": {
        out(JSON.stringify(await waitBench(cfg, argv[1]!)));
        return;
      }
      case "score": {
        // DashboardResponse { LeaderboardItem leaderboard_item = 1 { LeaderboardScore latest_score = 3 { int64 score = 1 } } }
        const dashboard = await fetchProto(cfg, requirePath(cfg.portal.scorePath, "scorePath"), "スコアの取得");
        const item = msg(dashboard, 1);
        if (!item.has(3)) throw benchError("確定したスコアがまだありません");
        out(String(int(msg(item, 3), 1)));
        return;
      }
    }
  },
};
