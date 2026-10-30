import { dirname } from "node:path";
import { fileURLToPath } from "node:url";
import type { Config } from "./src/config.js";

/**
 * isucon-harness 設定
 *   このリポジトリは ISUCON2026 専用・private なので、実設定をそのままコミットする。
 *   エージェントが作る worktree にもこのファイルが入り、どこからでも isuh が動く。
 *   秘密（ポータルの Cookie）はここに書かない。portal.cookieFile が指す別ファイルに置く。
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

  // 公開 API は無くブラウザ操作が前提。GitHub ログイン後のセッション Cookie
  // （isuxportal_session）を手元に持ち出して使う。詳細は src/lib/portal.ts。
  portal: {
    base: "https://portal.isucon.net",
    cookieFile: `${process.env["HOME"]}/.isucon/portal-cookie`,
    teamId: "",
    // isuxportal の API（ポータルの JS から読み取ったもの）
    whoamiPath: "/api/session",
    enqueuePath: "/api/contestant/benchmark_jobs",
    jobPath: "/api/contestant/benchmark_jobs/{job}",
    scorePath: "/api/contestant/dashboard",
  },

  bench: {
    lockPath: "/tmp/isucon-bench.lock", // hosts.primary 上
    lockTimeoutSec: 1800,
    timeoutSec: 600,
    // 本番は portal。過去問のリハーサルでは local（ベンチマーカーを ssh で直接実行）
    runner: "portal",
    // 例（ISUCON14 の AMI、アプリと同じホストで動かす）:
    //   host: "isu1",
    //   command: "./bench run --addr 127.0.0.1:443 --target https://isuride.xiv.isucon.net --payment-url http://127.0.0.1:12346 --payment-bind-port 12346",
    //   scorePattern: "スコア=(\\d+)",
    //   passPattern: "pass=true",
    local: { host: "", command: "", scorePattern: "", passPattern: "" },
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
