import { existsSync } from "node:fs";
import { join, resolve } from "node:path";
import { pathToFileURL } from "node:url";
import { configError } from "./lib/exit.js";
import { log } from "./lib/log.js";
import { repoRoot } from "./lib/paths.js";

/**
 * 設定の契約 — docs/DESIGN.md §2
 * 当日埋める値は「すべて」config.ts に置く。ソースに直書きしない。
 */

/** ホストの役割。値は ~/.ssh/config 上のホスト別名。 */
export interface HostConfig {
  /** 全ホスト。再起動試験の対象 */
  all: string[];
  /** アプリを動かすホスト。deploy の配布先 */
  app: string[];
  /** nginx のホスト。access log の収集元 */
  nginx: string[];
  /** DB のホスト。slow query log の収集元 */
  db: string;
  /** ベンチのリクエスト先。ロックの保持先 */
  primary: string;
  /** サーバ上の作業ユーザ */
  user: string;
}

export interface AppConfig {
  /** 手元のアプリリポジトリ。git ref はここで解決する（ハーネス自身ではない） */
  repoLocal: string;
  /** サーバ上のアプリ配置先 */
  dir: string;
  /** dir 配下の言語ディレクトリ */
  langDir: string;
  /** ビルド成果物のファイル名 */
  binary: string;
  /** systemctl のサービス名 */
  service: string;
  /** ビルドコマンド。repoLocal/langDir で実行される */
  buildCommand: string;
  /** check がリクエストを投げる先 */
  baseUrl: string;
  initializePath: string;
}

export interface MiddlewareConfig {
  nginxService: string;
  dbService: string;
  nginxAccessLog: string;
  dbSlowLog: string;
  /** finalize / verify-restart の enable 確認対象に加えるサービス */
  extraServices: string[];
}

export interface MetricsConfig {
  /** nginx のログ形式に合わせた alp のサブコマンド。ltsv にしておくのが楽 */
  alpFormat: "ltsv" | "json";
  /** alp の -m に渡す URI 正規化パターン。当日のエンドポイントに合わせる */
  alpMatches: string[];
  alpSort: "sum" | "avg" | "max" | "count";
  alpLimit: number;
  ptQueryLimit: number;
  /** アプリに pprof を仕込んだら true */
  pprofEnabled: boolean;
  /**
   * CPU プロファイルの URL。**hosts.app[0] から見た URL**（そのホスト上で curl する）。
   * ?seconds=N はベンチの長さに合わせる。ベンチ開始と同時に取り始め、N 秒経つまで待つ。
   */
  pprofUrl: string;
  resourceIntervalSec: number;
}

export interface PortalConfig {
  base: string;
  /** ブラウザから持ち出した sessionid / csrftoken を置くファイル */
  cookieFile: string;
  teamId: string;
  /** whoami が引く「ログインが必要なページ」。Cookie の生死確認だけに使う */
  whoamiPath: string;
  /** 当日判明したら埋める（ベンチ登録） */
  enqueuePath: string;
  /** 当日判明したら埋める（ジョブ状態の取得）。`{job}` を job-id に置換する */
  jobPath: string;
  /** 当日判明したら埋める（確定スコアの取得） */
  scorePath: string;
}

export interface BenchConfig {
  /** hosts.primary 上のロックファイル */
  lockPath: string;
  lockTimeoutSec: number;
  timeoutSec: number;
}

export interface HostAddress {
  /** 手元から入るときのアドレス。~/.ssh/isuh_config の HostName になる */
  public: string;
  /** サーバ間で使うアドレス。/etc/hosts とサーバ上の ~/.ssh/config に書く */
  private: string;
  /** 手元から入るときの ssh のポート。省略時 22（test/ のコンテナ環境で使う） */
  port?: number;
}

export interface SetupConfig {
  /** ホスト別名 → アドレス。hosts.all の全員ぶん。当日 CloudFormation の出力から埋める */
  addresses: Record<string, HostAddress>;
  /** 手元から入るときの秘密鍵。ポータルに公開鍵を登録したもの */
  identityFile: string;
  /** setup apprepo が作ったアプリリポジトリの push 先。空なら push しない */
  appRemote: string;
}

export interface Config {
  hosts: HostConfig;
  app: AppConfig;
  middleware: MiddlewareConfig;
  metrics: MetricsConfig;
  portal: PortalConfig;
  bench: BenchConfig;
  setup: SetupConfig;
}

const validate = (cfg: Config): void => {
  const missing: string[] = [];
  if (cfg.hosts.all.length === 0) missing.push("hosts.all");
  if (cfg.hosts.app.length === 0) missing.push("hosts.app");
  if (!cfg.hosts.primary) missing.push("hosts.primary");
  if (!cfg.app.repoLocal) missing.push("app.repoLocal");
  if (!cfg.app.dir) missing.push("app.dir");
  if (!cfg.app.service) missing.push("app.service");
  if (missing.length > 0) {
    throw configError(`config.ts に未設定の項目があります: ${missing.join(", ")}`);
  }
};

/** ISUH_CONFIG で場所を上書きできる（テスト用）。 */
export const loadConfig = async (): Promise<Config> => {
  const override = process.env["ISUH_CONFIG"];
  const path = override ? resolve(override) : join(repoRoot(), "config.ts");
  if (!existsSync(path)) {
    throw configError(
      `設定ファイルがありません: ${path}\n  cp config.example.ts config.ts して埋めてください。`,
    );
  }
  const mod: { config?: Config; default?: Config } = await import(pathToFileURL(path).href);
  const cfg = mod.config ?? mod.default;
  if (!cfg) throw configError(`config を export していません: ${path}`);
  validate(cfg);
  log.debug(`設定を読み込みました: ${path}`);
  return cfg;
};
