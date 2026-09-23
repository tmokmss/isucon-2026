import { dirname } from "node:path";
import { fileURLToPath } from "node:url";
import type { Config } from "./src/config.js";

/**
 * isucon-harness 設定
 *   cp config.example.ts config.ts して埋める。config.ts は .gitignore 済み。
 *   当日埋める値は「すべて」このファイルに集約する。ソースに直書きしない。
 */
export const config: Config = {
  // ホストの値は ~/.ssh/config 上のホスト別名。setup が生成する。
  hosts: {
    all: ["isu1", "isu2", "isu3"],
    app: ["isu1"],
    nginx: ["isu1"],
    db: "isu2",
    primary: "isu1",
    user: "isucon",
  },

  app: {
    // アプリ・ハーネス・構成管理（conf/）は同じリポジトリ。config.ts はリポジトリ直下に置くので、その場所
    repoLocal: dirname(fileURLToPath(import.meta.url)),
    // サーバの app.dir を写す場所（isuh setup apprepo が取り込む）
    localDir: "webapp",
    dir: "/home/isucon/webapp",
    langDir: "go",
    binary: "isuports",
    service: "isuports.go.service",
    buildCommand: "GOOS=linux GOARCH=amd64 go build -o isuports ./cmd/isuports",
    baseUrl: "http://isu1",
    initializePath: "/initialize",
  },

  middleware: {
    nginxService: "nginx.service",
    dbService: "mysql.service",
    nginxAccessLog: "/var/log/nginx/access.log",
    dbSlowLog: "/var/log/mysql/mysql-slow.log",
    extraServices: [], // 例: ["redis-server.service"]
  },

  metrics: {
    // nginx の log_format に合わせる。conf/ で ltsv にしておくと楽。
    alpFormat: "ltsv",
    // 当日のエンドポイントに合わせる。
    // 例: ["/api/player/competition/[^/]+/ranking", "/api/player/player/[^/]+"]
    alpMatches: [],
    alpSort: "sum",
    alpLimit: 30,
    ptQueryLimit: 15,
    pprofEnabled: false, // アプリに pprof を仕込んだら true
    // hosts.app[0] 上で curl する URL。seconds はベンチの長さ（通常 60 秒）に合わせる。
    pprofUrl: "http://localhost:6060/debug/pprof/profile?seconds=60",
    resourceIntervalSec: 1,
  },

  // 公開 API は無くブラウザ操作が前提。GitHub 認証の Django セッション Cookie を
  // 手元に持ち出して使う。詳細は #3 と docs/DESIGN.md §7。
  portal: {
    base: "https://portal.isucon.net",
    cookieFile: `${process.env["HOME"]}/.isucon/portal-cookie`,
    teamId: "",
    // Cookie の生死確認に使うページ。未ログインだと /auth/login/ へ 302 するところを選ぶ。
    // トップ (/) は未ログインでも 200 を返すので向かない（実物で確認済み）。
    whoamiPath: "/auth/settings/team/",
    // 当日埋める。まずは `isuh portal get <path>` で探る。
    enqueuePath: "",
    jobPath: "", // 例: "/api/job/{job}" — {job} が job-id に置換される
    scorePath: "",
  },

  bench: {
    lockPath: "/tmp/isucon-bench.lock", // hosts.primary 上
    lockTimeoutSec: 1800,
    timeoutSec: 600,
  },

  // isuh setup が使う。起動直後にまずここを埋めて `isuh setup` する。
  setup: {
    // 当日 CloudFormation の出力（またはポータル）から埋める。hosts.all の全員ぶん。
    addresses: {
      isu1: { public: "", private: "" },
      isu2: { public: "", private: "" },
      isu3: { public: "", private: "" },
    },
    // ポータルに公開鍵を登録した鍵。サーバの isucon ユーザにはこの鍵で入れる。
    identityFile: `${process.env["HOME"]}/.ssh/id_ed25519`,
  },
};
