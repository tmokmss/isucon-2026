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
    repoLocal: `${process.env["HOME"]}/projects/isucon2026-app`,
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
    pprofUrl: "http://isu1:6060/debug/pprof/profile?seconds=30",
    resourceIntervalSec: 1,
  },

  // 公開 API は無くブラウザ操作が前提。GitHub 認証の Django セッション Cookie を
  // 手元に持ち出して使う。詳細は #3 と docs/DESIGN.md §7。
  portal: {
    base: "https://portal.isucon.net",
    cookieFile: `${process.env["HOME"]}/.isucon/portal-cookie`,
    teamId: "",
    // Cookie の生死確認に使うページ。ログインが要るところなら何でもよい。
    whoamiPath: "/",
    // 当日埋める。まずは `isuh portal get <path>` で探る。
    enqueuePath: "",
    jobPath: "", // 例: "/api/job/{job}" — {job} が job-id に置換される
    scorePath: "",
  },

  // ベンチを使わない検証ゲート。レスポンス構造の変更（= 失格）を捕まえる防波堤。
  // 初回実行時に「今の形」を snapshotFile に記録し、以後はそれと比較する。
  // **アプリを触る前に 1 回走らせて、記録した形をコミットしておくこと。**
  check: {
    initialize: true,
    // 当日の主要エンドポイントを並べる。ログインが要るものは headers に Cookie を。
    endpoints: [
      // { path: "/api/player/competitions" },
      // { path: "/api/organizer/players", headers: { Cookie: "isuports_session=..." } },
    ],
    snapshotFile: "check/structure.json",
    timeoutSec: 30, // initialize が重いことがあるので長めに
  },

  bench: {
    lockPath: "/tmp/isucon-bench.lock", // hosts.primary 上
    lockTimeoutSec: 1800,
    timeoutSec: 600,
    scoreRegressionRatio: 0.9,
    autoRollback: false, // 序盤 false、終盤 true を推奨
  },
};
