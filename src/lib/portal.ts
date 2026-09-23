import { existsSync, readFileSync } from "node:fs";
import type { Config } from "../config.js";
import { EXIT, HarnessError, configError } from "./exit.js";
import { log } from "./log.js";

/**
 * ポータルへの HTTP — docs/DESIGN.md §3
 *
 * 公開 API は無く、ブラウザ操作が前提。ログインは /auth/login/ からの GitHub 認証で、
 * 中身は Django のセッション Cookie 方式（vary: Cookie）。
 * ブラウザでログインした後の Cookie を持ち出して使う。
 *
 * ベンチ用のエンドポイントは当日にならないと分からないので、パスはすべて
 * config.portal に置く。ここにあるのは「Cookie を付けて叩き、ログイン切れを
 * 見分ける」ところまで。
 */

export interface PortalResponse {
  status: number;
  url: string;
  location: string | null;
  contentType: string;
  body: string;
}

const COOKIE_HELP = `
Cookie の取り出し方:
  1. ブラウザで https://portal.isucon.net/ にログインする
  2. DevTools → Application → Cookies → https://portal.isucon.net
  3. sessionid と csrftoken の値をコピーして、次の 1 行をファイルに保存する

       sessionid=<値>; csrftoken=<値>

  （DevTools の Network タブでリクエストを右クリック →
    Copy → Copy as cURL して、-H 'Cookie: ...' の中身を貼っても同じ）

置き場所は config.ts の portal.cookieFile。`;

/** Cookie ヘッダの中身を読む。`#` 始まりの行は無視し、複数行は `; ` で連結する。 */
export const loadCookie = (cfg: Config): string => {
  const path = cfg.portal.cookieFile;
  if (path === "") throw configError(`portal.cookieFile が未設定です。${COOKIE_HELP}`);
  if (!existsSync(path)) {
    throw configError(`Cookie ファイルがありません: ${path}${COOKIE_HELP}`);
  }
  const cookie = readFileSync(path, "utf8")
    .split("\n")
    .map((line) => line.trim().replace(/^Cookie:\s*/i, ""))
    .filter((line) => line !== "" && !line.startsWith("#"))
    .join("; ")
    .replace(/;\s*;/g, "; ")
    .replace(/;\s*$/, "");
  if (!cookie.includes("=")) {
    throw configError(`Cookie ファイルの形式が読めません: ${path}${COOKIE_HELP}`);
  }
  return cookie;
};

const csrfFrom = (cookie: string): string | undefined =>
  /(?:^|;\s*)csrftoken=([^;]+)/.exec(cookie)?.[1];

export interface RequestOptions {
  /** POST の本体。オブジェクトなら application/x-www-form-urlencoded にする */
  form?: Record<string, string>;
  json?: unknown;
  /** 既定は GET */
  method?: string;
}

/** ポータルを叩く。リダイレクトは追わない（ログイン切れを見分けるため）。 */
export const portalRequest = async (
  cfg: Config,
  path: string,
  opts: RequestOptions = {},
): Promise<PortalResponse> => {
  const cookie = loadCookie(cfg);
  const base = cfg.portal.base.replace(/\/$/, "");
  const url = /^https?:\/\//.test(path) ? path : `${base}${path.startsWith("/") ? "" : "/"}${path}`;
  const method = opts.method ?? (opts.form !== undefined || opts.json !== undefined ? "POST" : "GET");

  const headers: Record<string, string> = {
    Cookie: cookie,
    Accept: "application/json, text/html;q=0.9, */*;q=0.8",
    "User-Agent": "isuh/0.1",
  };

  let body: string | undefined;
  if (opts.json !== undefined) {
    body = JSON.stringify(opts.json);
    headers["Content-Type"] = "application/json";
  } else if (opts.form !== undefined) {
    body = new URLSearchParams(opts.form).toString();
    headers["Content-Type"] = "application/x-www-form-urlencoded";
  }

  if (method !== "GET" && method !== "HEAD") {
    // Django の CSRF は「Cookie の csrftoken と X-CSRFToken が一致するか」を見る。
    // Referer / Origin も https では必須。
    const token = csrfFrom(cookie);
    if (token === undefined) {
      log.warn("Cookie に csrftoken がありません。POST は CSRF で弾かれる可能性があります");
    } else {
      headers["X-CSRFToken"] = token;
    }
    headers["Referer"] = url;
    headers["Origin"] = base;
  }

  log.debug(`portal ${method} ${url}`);
  let response: Response;
  try {
    response = await fetch(url, { method, headers, body, redirect: "manual" });
  } catch (error) {
    throw new HarnessError(
      EXIT.BENCH,
      `ポータルに接続できません: ${url}\n  ${error instanceof Error ? error.message : String(error)}`,
    );
  }

  return {
    status: response.status,
    url,
    location: response.headers.get("location"),
    contentType: response.headers.get("content-type") ?? "",
    body: await response.text(),
  };
};

const LOGIN_LINK = /href="[^"]*\/auth\/login\/?"/i;
// 実物のポータルのログアウトはリンクではなく POST のフォーム（<form action="/auth/logout/">）
const LOGOUT_LINK = /(?:href|action)="[^"]*\/auth\/logout\/?"/i;

/**
 * ログインが切れていないか。
 *
 * Django で「ログインが必要なページ」に未認証で来ると /auth/login/ へ 302 する。
 * 403 は CSRF か権限。どちらも Cookie を取り直せば直るので同じ扱いにする。
 *
 * ただし **ポータルのトップは未ログインでも 200 を返す**（実物で確認済み）。
 * ステータスだけ見ると「Cookie は有効」と誤判定するので、HTML の中身も見る。
 * 未ログインなら「ログイン」のリンク (/auth/login/)、ログイン済みなら
 * ログアウトのフォーム (action="/auth/logout/") が出る。
 *
 * とはいえ中身の推測は壊れやすい。whoamiPath には、未ログインだと素直に 302 を返す
 * ページ（/auth/settings/team/）を指定するのが確実で、既定値もそうしてある。
 */
export const isLoggedOut = (response: PortalResponse): boolean => {
  if (response.status === 401 || response.status === 403) return true;
  if (response.status >= 300 && response.status < 400) {
    return /\/auth\/login|\/login/.test(response.location ?? "");
  }
  if (response.status === 200 && /html/i.test(response.contentType)) {
    if (LOGOUT_LINK.test(response.body)) return false;
    if (LOGIN_LINK.test(response.body)) return true;
  }
  return false;
};

export const assertLoggedIn = (cfg: Config, response: PortalResponse): void => {
  if (!isLoggedOut(response)) return;
  const where = response.location === null ? "" : `\n  リダイレクト先: ${response.location}`;
  throw configError(
    `ポータルの Cookie が効いていません (HTTP ${response.status}) ${response.url}${where}\n` +
      `  ブラウザで入り直して Cookie を取り直してください。${COOKIE_HELP}`,
  );
};

/** レスポンスが JSON なら解析する。HTML なら undefined。 */
export const asJson = (response: PortalResponse): unknown => {
  if (!/json/i.test(response.contentType)) {
    try {
      return JSON.parse(response.body);
    } catch {
      return undefined;
    }
  }
  try {
    return JSON.parse(response.body);
  } catch {
    return undefined;
  }
};

/** JSON からそれらしいキーを拾う。当日の形が分からないので候補を順に見る。 */
export const pick = (value: unknown, keys: readonly string[]): string | undefined => {
  if (typeof value !== "object" || value === null) return undefined;
  const record = value as Record<string, unknown>;
  for (const key of keys) {
    const found = record[key];
    if (typeof found === "string" && found !== "") return found;
    if (typeof found === "number") return String(found);
  }
  // 1 段だけ潜る（{ "job": { "id": ... } } のような形に対応）
  for (const nested of Object.values(record)) {
    if (typeof nested === "object" && nested !== null && !Array.isArray(nested)) {
      const found = pick(nested, keys);
      if (found !== undefined) return found;
    }
  }
  return undefined;
};

/** HTML から <title> を拾う。whoami で「誰として見えているか」の手がかりにする。 */
export const titleOf = (html: string): string | undefined =>
  /<title[^>]*>([\s\S]*?)<\/title>/i.exec(html)?.[1]?.trim().replace(/\s+/g, " ");
