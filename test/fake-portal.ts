/**
 * ポータルの代役。Django 風の挙動だけを真似る。
 *
 *   - 認証が要るページに未認証で来たら /auth/login/ へ 302
 *   - POST は Cookie の csrftoken と X-CSRFToken が一致しなければ 403
 *   - ベンチのジョブは pending → running → done と進む
 *
 * 本物のポータルは当日までエンドポイントが分からないので、ここで確かめるのは
 * 「Cookie を付けて叩けるか」「ログイン切れを見分けられるか」「CSRF が通るか」
 * の 3 点だけ。
 */
import { readFileSync } from "node:fs";
import { createServer } from "node:http";

const PORT = Number(process.env["ISUH_FAKE_PORTAL_PORT"] ?? "18080");
const VALID_SESSION = "valid-session";

/** 返すスコア。テストから .work/fake-score を書き換えて変えられる。 */
const SCORE_FILE = new URL(".work/fake-score", import.meta.url);
const currentScore = (): number => {
  try {
    const value = Number(readFileSync(SCORE_FILE, "utf8").trim());
    if (Number.isFinite(value)) return value;
  } catch {
    /* 無ければ既定値 */
  }
  return 12345;
};

/** 完了時の状態。.work/fake-status に "fail" と書けばベンチ失敗を再現できる。 */
const FINAL_STATUS_FILE = new URL(".work/fake-status", import.meta.url);
const finalStatus = (): string => {
  try {
    const value = readFileSync(FINAL_STATUS_FILE, "utf8").trim();
    if (value !== "") return value === "fail" ? "failed" : value;
  } catch {
    /* 無ければ既定値 */
  }
  return "done";
};

/** job-id → status を返した回数。3 回目で done にする。 */
const polls = new Map<string, number>();

const cookiesOf = (header: string | undefined): Record<string, string> => {
  const result: Record<string, string> = {};
  for (const part of (header ?? "").split(";")) {
    const index = part.indexOf("=");
    if (index > 0) result[part.slice(0, index).trim()] = part.slice(index + 1).trim();
  }
  return result;
};

const server = createServer((req, res) => {
  const url = new URL(req.url ?? "/", `http://127.0.0.1:${PORT}`);
  const cookies = cookiesOf(req.headers.cookie);
  const authed = cookies["sessionid"] === VALID_SESSION;

  const send = (status: number, contentType: string, body: string): void => {
    res.writeHead(status, { "content-type": contentType, vary: "Cookie" });
    res.end(body);
  };
  const json = (status: number, value: unknown): void =>
    send(status, "application/json", JSON.stringify(value));

  // ログインページは誰でも見られる
  if (url.pathname === "/auth/login/") {
    send(200, "text/html; charset=utf-8", "<html><title>Sign in</title></html>");
    return;
  }

  // トップは未ログインでも 200 を返す（本物のポータルがそうなっている）。
  // 見分けはナビのリンク。
  if (url.pathname === "/") {
    const nav = authed
      ? '<a href="/auth/logout/">ログアウト</a>'
      : '<a href="/auth/login/">ログイン</a>';
    send(200, "text/html; charset=utf-8", `<html><head><title>ISUCON2026 Portal</title></head><body>${nav}</body></html>`);
    return;
  }

  if (!authed) {
    res.writeHead(302, { location: `/auth/login/?next=${encodeURIComponent(url.pathname)}` });
    res.end();
    return;
  }

  if (req.method !== "GET" && req.method !== "HEAD") {
    const sent = req.headers["x-csrftoken"];
    if (typeof sent !== "string" || sent === "" || sent !== cookies["csrftoken"]) {
      send(403, "text/html; charset=utf-8", "<html><title>CSRF verification failed</title></html>");
      return;
    }
  }

  if (url.pathname === "/team/") {
    send(
      200,
      "text/html; charset=utf-8",
      '<html><head><title>チーム情報 | ISUCON</title></head><body><a href="/auth/logout/">ログアウト</a></body></html>',
    );
    return;
  }

  if (url.pathname === "/api/bench/enqueue" && req.method === "POST") {
    // 登録のたびに進行をリセットする。毎回 pending → running → done を辿らせて、
    // ベンチがある程度の時間続く状況（= リソース計測が意味を持つ状況）を作る。
    polls.delete("4321");
    json(200, { job_id: "4321", status: "pending" });
    return;
  }

  const jobMatch = /^\/api\/bench\/job\/([^/]+)$/.exec(url.pathname);
  if (jobMatch !== null) {
    const id = jobMatch[1]!;
    const seen = (polls.get(id) ?? 0) + 1;
    polls.set(id, seen);
    if (seen >= 3) {
      const status = finalStatus();
      json(200, {
        job_id: id,
        status,
        score: status === "done" ? currentScore() : 0,
        reason: status === "done" ? "" : "整合性チェックに失敗しました",
      });
      return;
    }
    json(200, { job_id: id, status: seen === 1 ? "pending" : "running" });
    return;
  }

  if (url.pathname === "/api/score") {
    json(200, { team: "7", score: currentScore() });
    return;
  }

  send(404, "text/html; charset=utf-8", "<html><title>Not Found</title></html>");
});

server.listen(PORT, "127.0.0.1", () => {
  process.stderr.write(`fake-portal listening on http://127.0.0.1:${PORT}\n`);
});
