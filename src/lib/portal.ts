import { existsSync, readFileSync } from "node:fs";
import type { Config } from "../config.js";
import { EXIT, HarnessError, configError } from "./exit.js";
import { log } from "./log.js";

/**
 * ポータルへの HTTP — docs/DESIGN.md §3
 *
 * ポータルは isuxportal（Rails）。公開 API は無く、ブラウザの JS が叩く API を使う。
 *   - 認証は GitHub ログイン後の Cookie `isuxportal_session`。ブラウザから持ち出して使う
 *   - API の本体は protobuf（application/vnd.google.protobuf）。中身は lib/proto.ts で読む
 *   - POST の CSRF は X-Csrf-Token ヘッダ。値はページの <meta name="csrf-token"> から取る
 *   - 未ログインで /api/contestant/* を叩くと /session/new/github へ 302
 */

export interface PortalResponse {
  status: number;
  url: string;
  location: string | null;
  contentType: string;
  bytes: Uint8Array;
  body: string;
}

const COOKIE_HELP = `
Cookie の取り出し方:
  1. ブラウザで https://portal.isucon.net/ にログインする
  2. DevTools → Application → Cookies → https://portal.isucon.net
  3. isuxportal_session の値をコピーして、次の 1 行をファイルに保存する

       isuxportal_session=<値>

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

export interface RequestOptions {
  /** POST の本体。オブジェクトなら application/x-www-form-urlencoded にする */
  form?: Record<string, string>;
  /** POST の本体（protobuf） */
  proto?: Uint8Array;
  /** 既定は GET */
  method?: string;
}

const send = async (
  url: string,
  method: string,
  headers: Record<string, string>,
  body?: string | Uint8Array,
): Promise<PortalResponse> => {
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
  const bytes = new Uint8Array(await response.arrayBuffer());
  return {
    status: response.status,
    url,
    location: response.headers.get("location"),
    contentType: response.headers.get("content-type") ?? "",
    bytes,
    body: Buffer.from(bytes).toString("utf8"),
  };
};

/** トップページの <meta name="csrf-token"> を取る。トークンはセッションに紐づく。 */
const fetchCsrfToken = async (base: string, headers: Record<string, string>): Promise<string> => {
  const page = await send(`${base}/`, "GET", { ...headers, Accept: "text/html" });
  const token = /<meta name="csrf-token" content="([^"]+)"/.exec(page.body)?.[1];
  if (token === undefined) {
    log.warn(`ポータルのトップに csrf-token がありません (HTTP ${page.status})。POST は弾かれる可能性があります`);
    return "";
  }
  return token;
};

/** ポータルを叩く。リダイレクトは追わない（ログイン切れを見分けるため）。 */
export const portalRequest = async (
  cfg: Config,
  path: string,
  opts: RequestOptions = {},
): Promise<PortalResponse> => {
  const cookie = loadCookie(cfg);
  const base = cfg.portal.base.replace(/\/$/, "");
  const url = /^https?:\/\//.test(path) ? path : `${base}${path.startsWith("/") ? "" : "/"}${path}`;
  const method = opts.method ?? (opts.form !== undefined || opts.proto !== undefined ? "POST" : "GET");

  const headers: Record<string, string> = {
    Cookie: cookie,
    Accept: "application/protobuf, application/vnd.google.protobuf, text/html;q=0.9, */*;q=0.8",
    "User-Agent": "isuh/0.1",
  };

  let body: string | Uint8Array | undefined;
  if (opts.proto !== undefined) {
    body = opts.proto;
    headers["Content-Type"] = "application/vnd.google.protobuf";
  } else if (opts.form !== undefined) {
    body = new URLSearchParams(opts.form).toString();
    headers["Content-Type"] = "application/x-www-form-urlencoded";
  }

  if (method !== "GET" && method !== "HEAD") {
    headers["X-Csrf-Token"] = await fetchCsrfToken(base, headers);
    headers["Referer"] = `${base}/`;
    headers["Origin"] = base;
  }

  return send(url, method, headers, body);
};

/** protobuf で返ってきたか。エラーは `; proto=isuxportal.proto.Error` 付きで返る。 */
export const isProto = (response: PortalResponse): boolean =>
  /^application\/(vnd\.google\.)?protobuf/i.test(response.contentType) &&
  !/proto=isuxportal\.proto\.Error/.test(response.contentType);

/**
 * ログインが切れていないか。
 *
 * 未ログインで認証が要る API に来ると /session/new/github へ 302 する。
 * 403 は CSRF か権限。どちらも Cookie を取り直せば直るので同じ扱いにする。
 * /api/session は未ログインでも 200 を返すので、そちらは中身（contestant の有無）で見る。
 */
export const isLoggedOut = (response: PortalResponse): boolean => {
  if (response.status === 401 || response.status === 403) return true;
  if (response.status >= 300 && response.status < 400) {
    return /\/session\/new|\/login/.test(response.location ?? "");
  }
  return false;
};

export const loggedOutError = (response: PortalResponse): HarnessError => {
  const where = response.location === null ? "" : `\n  リダイレクト先: ${response.location}`;
  return configError(
    `ポータルの Cookie が効いていません (HTTP ${response.status}) ${response.url}${where}\n` +
      `  ブラウザで入り直して Cookie を取り直してください。${COOKIE_HELP}`,
  );
};

export const assertLoggedIn = (response: PortalResponse): void => {
  if (isLoggedOut(response)) throw loggedOutError(response);
};
