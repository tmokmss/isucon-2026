import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import type { Config } from "../src/config.js";

/**
 * 検証環境（test/up.sh が起動するコンテナ）向けの設定。
 *   ISUH_CONFIG=test/config.test.ts ISUH_SSH_CONFIG=test/.work/ssh_config bin/isuh ...
 */
const here = dirname(fileURLToPath(import.meta.url));

// コンテナは手元の Docker と同じアーキテクチャで動く
const goarch = process.arch === "arm64" ? "arm64" : "amd64";

export const config: Config = {
  hosts: {
    all: ["isu1", "isu2", "isu3"],
    // 2 台に配ることで「ホスト間は並列・1 台でも失敗したら全体失敗」を検証する
    app: ["isu1", "isu2"],
    // nginx を 2 台にして、alp が全台を集計することを確かめる
    nginx: ["isu1", "isu2"],
    db: "isu3",
    primary: "isu1",
    user: "isucon",
  },

  app: {
    // 本番はこのリポジトリ自身。テストはハーネスを汚さないよう別のリポジトリを使う（up.sh が作る）
    repoLocal: join(here, ".work", "app-repo"),
    localDir: "webapp",
    dir: "/home/isucon/webapp",
    langDir: "go",
    binary: "isuapp",
    service: "isuapp.go.service",
    buildCommand: `GOOS=linux GOARCH=${goarch} CGO_ENABLED=0 go build -o isuapp .`,
    baseUrl: "http://127.0.0.1:8080",
    initializePath: "/initialize",
  },

  middleware: {
    nginxService: "nginx.service",
    dbService: "mysql.service",
    nginxAccessLog: "/var/log/nginx/access.log",
    dbSlowLog: "/var/log/mysql/mysql-slow.log",
    extraServices: [],
  },

  metrics: {
    alpFormat: "ltsv",
    alpMatches: [],
    alpSort: "sum",
    alpLimit: 30,
    ptQueryLimit: 15,
    pprofEnabled: true,
    // テストのベンチは数秒で終わるので短く
    pprofUrl: "http://localhost:8080/debug/pprof/profile?seconds=3",
    resourceIntervalSec: 1,
  },

  // test/fake-portal.ts が 127.0.0.1:18080 で isuxportal 風の応答を返す
  portal: {
    base: process.env["ISUH_TEST_PORTAL"] ?? "http://127.0.0.1:18080",
    cookieFile: join(here, ".work", "portal-cookie"),
    teamId: "7",
    whoamiPath: "/api/session",
    enqueuePath: "/api/contestant/benchmark_jobs",
    jobPath: "/api/contestant/benchmark_jobs/{job}",
    scorePath: "/api/contestant/dashboard",
  },

  bench: {
    lockPath: "/tmp/isucon-bench.lock",
    lockTimeoutSec: 3,
    timeoutSec: 60,
    // bench-local.sh が ISUH_TEST_RUNNER=local で切り替える。偽のベンチマーカーは isu1 の ~/fake-bench.sh
    runner: process.env["ISUH_TEST_RUNNER"] === "local" ? "local" : "portal",
    local: {
      host: "isu1",
      command: "sh fake-bench.sh",
      scorePattern: "スコア=(\\d+)",
      passPattern: "pass=true",
    },
  },

  // 手元からは 127.0.0.1 の転送ポート、サーバ間は docker-compose.yml で固定したアドレス
  setup: {
    addresses: {
      isu1: { public: "127.0.0.1", port: 22101, private: "10.99.0.11" },
      isu2: { public: "127.0.0.1", port: 22102, private: "10.99.0.12" },
      isu3: { public: "127.0.0.1", port: 22103, private: "10.99.0.13" },
    },
    identityFile: join(here, ".work", "id_ed25519"),
  },
};
