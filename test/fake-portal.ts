/**
 * ポータルの代役。isuxportal の挙動だけを真似る。
 *
 *   - Cookie は isuxportal_session。/api/session は未ログインでも 200（contestant が空）
 *   - /api/contestant/* に未ログインで来たら /session/new/github へ 302
 *   - POST は X-Csrf-Token がトップの <meta name="csrf-token"> と一致しなければ 422
 *   - 本体は protobuf。ベンチのジョブは PENDING → RUNNING → FINISHED と進む
 */
import { readFileSync } from "node:fs";
import { createServer } from "node:http";
import { decode, encodeInt, encodeMsg, encodeStr, int } from "../src/lib/proto.js";

const PORT = Number(process.env["ISUH_FAKE_PORTAL_PORT"] ?? "18080");
const VALID_SESSION = "valid-session";
const CSRF_TOKEN = "csrf-for-valid-session";
const TARGET_ID = 101;

/** テストから .work/ のファイルで挙動を変える。fake-score はスコア、fake-status に fail と書くと不合格。 */
const readWork = (name: string): string | undefined => {
  try {
    return readFileSync(new URL(`.work/${name}`, import.meta.url), "utf8").trim();
  } catch {
    return undefined;
  }
};
const currentScore = (): number => Number(readWork("fake-score") ?? "12345");

/** job-id → 状態を返した回数。3 回目で FINISHED にする。 */
const polls = new Map<string, number>();

const cookiesOf = (header: string | undefined): Record<string, string> => {
  const result: Record<string, string> = {};
  for (const part of (header ?? "").split(";")) {
    const index = part.indexOf("=");
    if (index > 0) result[part.slice(0, index).trim()] = part.slice(index + 1).trim();
  }
  return result;
};

// isuxportal/resources/benchmark_job.proto の BenchmarkJob
const job = (id: number, status: number, result?: Uint8Array): Uint8Array =>
  encodeMsg(1, encodeInt(1, id), encodeInt(3, TARGET_ID), encodeInt(4, status), ...(result ? [result] : []));

const server = createServer((req, res) => {
  const url = new URL(req.url ?? "/", `http://127.0.0.1:${PORT}`);
  const authed = cookiesOf(req.headers.cookie)["isuxportal_session"] === VALID_SESSION;
  const chunks: Buffer[] = [];
  req.on("data", (chunk: Buffer) => chunks.push(chunk));
  req.on("end", () => {
    const body = Buffer.concat(chunks);
    const proto = (name: string, bytes: Uint8Array, status = 200): void => {
      res.writeHead(status, { "content-type": `application/protobuf; proto=isuxportal.proto.${name}` });
      res.end(bytes);
    };
    const html = (status: number, text: string): void => {
      res.writeHead(status, { "content-type": "text/html; charset=utf-8" });
      res.end(text);
    };

    if (url.pathname === "/") {
      const token = authed ? CSRF_TOKEN : "csrf-for-guest";
      html(200, `<html><head><title>ISUCON2026 Portal</title><meta name="csrf-token" content="${token}"></head></html>`);
      return;
    }

    if (url.pathname === "/api/session") {
      const session = authed
        ? Buffer.concat([
            encodeMsg(1, encodeInt(1, 7), encodeStr(2, "isuh-test")),
            encodeMsg(2, encodeInt(1, 70), encodeStr(3, "tester")),
            encodeMsg(5, encodeInt(7, TARGET_ID), encodeInt(3, 1), encodeStr(4, "192.0.2.1"), encodeStr(5, "10.0.0.1")),
            encodeMsg(5, encodeInt(7, 102), encodeInt(3, 2), encodeStr(4, "192.0.2.2"), encodeStr(5, "10.0.0.2")),
          ])
        : new Uint8Array();
      proto("services.common.GetCurrentSessionResponse", session);
      return;
    }

    if (!authed) {
      res.writeHead(302, { location: `/session/new/github?back_to=${encodeURIComponent(url.pathname)}` });
      res.end();
      return;
    }

    if (req.method !== "GET" && req.headers["x-csrf-token"] !== CSRF_TOKEN) {
      proto("Error", encodeStr(3, "ActionController::InvalidAuthenticityToken"), 422);
      return;
    }

    if (url.pathname === "/api/contestant/benchmark_jobs" && req.method === "POST") {
      if (int(decode(body), 1) !== TARGET_ID) {
        proto("Error", encodeStr(3, "target_id が違います"), 400);
        return;
      }
      // 登録のたびに進行をリセットする。毎回 PENDING → RUNNING → FINISHED を辿らせて、
      // ベンチがある程度の時間続く状況（= リソース計測が意味を持つ状況）を作る。
      polls.delete("4321");
      proto("services.contestant.EnqueueBenchmarkJobResponse", job(4321, 0));
      return;
    }

    const jobMatch = /^\/api\/contestant\/benchmark_jobs\/(\d+)$/.exec(url.pathname);
    if (jobMatch !== null) {
      const id = Number(jobMatch[1]);
      const seen = (polls.get(String(id)) ?? 0) + 1;
      polls.set(String(id), seen);
      if (seen < 3) {
        proto("services.contestant.GetBenchmarkJobResponse", job(id, seen === 1 ? 0 : 1));
        return;
      }
      const passed = readWork("fake-status") !== "fail";
      // BenchmarkResult { finished=1 passed=2 score=3 execution=5 { reason=1 stdout=2 stderr=3 } }
      const result = encodeMsg(
        18,
        encodeInt(1, 1),
        encodeInt(2, passed ? 1 : 0),
        encodeInt(3, passed ? currentScore() : 0),
        encodeMsg(5, encodeStr(1, passed ? "" : "fail: 整合性チェックに失敗"), encodeStr(2, "bench stdout"), encodeStr(3, "bench stderr")),
      );
      proto("services.contestant.GetBenchmarkJobResponse", job(id, 4, result));
      return;
    }

    if (url.pathname === "/api/contestant/dashboard") {
      // DashboardResponse { leaderboard_item=1 { latest_score=3 { score=1 } } }
      proto("services.contestant.DashboardResponse", encodeMsg(1, encodeMsg(3, encodeInt(1, currentScore()))));
      return;
    }

    html(404, "<html><title>Not Found</title></html>");
  });
});

server.listen(PORT, "127.0.0.1", () => {
  process.stderr.write(`fake-portal listening on http://127.0.0.1:${PORT}\n`);
});
