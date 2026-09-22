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
    nginx: ["isu1"],
    db: "isu3",
    primary: "isu1",
    user: "isucon",
  },

  app: {
    repoLocal: join(here, ".work", "app-repo"),
    dir: "/home/isucon/webapp",
    langDir: "go",
    binary: "isuapp",
    service: "isuapp.go.service",
    buildCommand: `GOOS=linux GOARCH=${goarch} CGO_ENABLED=0 go build -o isuapp .`,
    baseUrl: "http://127.0.0.1:18081",
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
    pprofEnabled: false,
    pprofUrl: "",
    resourceIntervalSec: 1,
  },

  // test/fake-portal.ts が 127.0.0.1:18080 で Django 風の応答を返す
  portal: {
    base: process.env["ISUH_TEST_PORTAL"] ?? "http://127.0.0.1:18080",
    cookieFile: join(here, ".work", "portal-cookie"),
    teamId: "7",
    whoamiPath: "/team/",
    enqueuePath: "/api/bench/enqueue",
    jobPath: "/api/bench/job/{job}",
    scorePath: "/api/score",
  },

  check: {
    initialize: true,
    endpoints: [{ path: "/api/items" }, { path: "/api/user" }],
    snapshotFile: join(here, ".work", "structure.json"),
    timeoutSec: 10,
  },

  bench: {
    lockPath: "/tmp/isucon-bench.lock",
    lockTimeoutSec: 3,
    timeoutSec: 60,
    scoreRegressionRatio: 0.9,
    autoRollback: false,
  },
};
