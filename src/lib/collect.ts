import { mkdirSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import type { Config } from "../config.js";
import { log } from "./log.js";
import { copyFrom, local, onHost, onHostShell, shQuote } from "./ssh.js";

/**
 * 計測の生データを集める — docs/DESIGN.md §4
 *
 * ここが作るのは runs/<run-id>/ 以下の生ファイルだけ。
 * それを読んで report.md に整形するのは #4（src/lib/report.ts）。
 *
 * 計測ツールが入っていないことは普通に起きる（導入は #11、pprof は仕込むかどうか自体が
 * 判断）。その場合は落とさず「未取得」と理由をファイルに書く。ベンチ 1 回は高価なので、
 * 補助的な計測が無いことでランを失敗にはしない。
 */

const RESOURCE_OUT = "/tmp/isuh-resource.txt";
const RESOURCE_PID = "/tmp/isuh-resource.pid";

const write = (dir: string, name: string, body: string): void => {
  writeFileSync(join(dir, name), body.endsWith("\n") ? body : `${body}\n`);
};

// ---------------------------------------------------------------------------
// ベンチ前
// ---------------------------------------------------------------------------

/**
 * ログを空にする。
 * alp と pt-query-digest がそのラン分だけを見るようにするため、ベンチ直前に必ず通る。
 */
export const resetLogs = async (cfg: Config): Promise<void> => {
  const targets: { host: string; path: string; what: string }[] = [
    ...cfg.hosts.nginx.map((host) => ({
      host,
      path: cfg.middleware.nginxAccessLog,
      what: "access log",
    })),
    { host: cfg.hosts.db, path: cfg.middleware.dbSlowLog, what: "slow query log" },
  ];

  await Promise.all(
    targets.map(async ({ host, path, what }) => {
      if (path === "") return;
      const result = await onHost(host, ["sudo", "truncate", "-s", "0", path], {
        allowFailure: true,
      });
      if (result.exitCode !== 0) {
        log.warn(`${host}: ${what} を空にできません (${path})`);
      }
    }),
  );
};

// ---------------------------------------------------------------------------
// ベンチ中
// ---------------------------------------------------------------------------

export interface ResourceSampler {
  /** 計測を止めて runs/<run-id>/resource/<host>.txt に回収する。多重呼び出し安全。 */
  stop(): Promise<void>;
}

/**
 * 全ホストで vmstat を回し始める。
 * 「どのホストが飽和したか」は report.md の Resources 節で一番効く情報なので、
 * ベンチの間ずっと取り続ける。
 */
export const startResourceSampling = async (
  cfg: Config,
  runDir: string,
): Promise<ResourceSampler> => {
  const hosts = cfg.hosts.all;
  const interval = Math.max(1, Math.floor(cfg.metrics.resourceIntervalSec));
  const out = shQuote(RESOURCE_OUT);
  const pid = shQuote(RESOURCE_PID);

  await Promise.all(
    hosts.map(async (host) => {
      // vmstat を直接 background にする（サブシェルを挟むと $! が vmstat の PID でなくなる）
      const script = [
        `rm -f ${out} ${pid}`,
        `{ echo "# host=$(hostname) nproc=$(nproc)"; free -m 2>/dev/null | sed 's/^/# /'; } > ${out} 2>&1`,
        `nohup vmstat -n ${interval} >> ${out} 2>&1 &`,
        `echo $! > ${pid}`,
      ].join("\n");
      const result = await onHostShell(host, script, { allowFailure: true });
      if (result.exitCode !== 0) log.warn(`${host}: リソース計測を開始できません`);
    }),
  );
  log.info(`リソース計測を開始しました (${interval}s 間隔, ${hosts.length} 台)`);

  const dir = join(runDir, "resource");
  mkdirSync(dir, { recursive: true });

  let stopped = false;
  return {
    async stop() {
      if (stopped) return;
      stopped = true;
      await Promise.all(
        hosts.map(async (host) => {
          const script = [
            `if [ -f ${pid} ]; then kill "$(cat ${pid})" 2>/dev/null || true; fi`,
            `rm -f ${pid}`,
            `{ echo; echo "# uptime: $(uptime)"; } >> ${out} 2>/dev/null || true`,
          ].join("\n");
          await onHostShell(host, script, { allowFailure: true });
          const got = await copyFrom(host, RESOURCE_OUT, join(dir, `${host}.txt`), {
            allowFailure: true,
          });
          if (!got) {
            write(dir, `${host}.txt`, `未取得: ${host} からリソース計測結果を回収できませんでした`);
          }
        }),
      );
      log.info("リソース計測を回収しました");
    },
  };
};

// ---------------------------------------------------------------------------
// ベンチ後
// ---------------------------------------------------------------------------

const collectAlp = async (cfg: Config, runDir: string): Promise<void> => {
  const host = cfg.hosts.nginx[0];
  if (host === undefined || cfg.middleware.nginxAccessLog === "") {
    write(runDir, "alp.txt", "未取得: hosts.nginx / middleware.nginxAccessLog が未設定");
    return;
  }
  if (cfg.metrics.alpMatches.length === 0) {
    log.warn("metrics.alpMatches が空です。URI が正規化されず、集計が意味を持ちません");
  }

  const args = [
    "sudo",
    "alp",
    cfg.metrics.alpFormat,
    "--file",
    cfg.middleware.nginxAccessLog,
    "--sort",
    cfg.metrics.alpSort,
    "-r",
    "--limit",
    String(cfg.metrics.alpLimit),
  ];
  if (cfg.metrics.alpMatches.length > 0) args.push("-m", cfg.metrics.alpMatches.join(","));

  const result = await onHost(host, args, { allowFailure: true });
  if (result.exitCode === 0 && result.stdout.trim() !== "") {
    write(runDir, "alp.txt", result.stdout);
  } else {
    write(
      runDir,
      "alp.txt",
      `未取得: alp を実行できませんでした (${host}, exit=${result.exitCode})\n` +
        `導入は #11。手元で確認するには:\n  ssh ${host} ${args.join(" ")}\n\n${result.stderr}`,
    );
    log.warn(`alp を実行できませんでした (${host})`);
  }
};

const collectSlowQuery = async (cfg: Config, runDir: string): Promise<void> => {
  const host = cfg.hosts.db;
  if (host === "" || cfg.middleware.dbSlowLog === "") {
    write(runDir, "ptquery.txt", "未取得: hosts.db / middleware.dbSlowLog が未設定");
    return;
  }

  const args = [
    "sudo",
    "pt-query-digest",
    "--limit",
    String(cfg.metrics.ptQueryLimit),
    cfg.middleware.dbSlowLog,
  ];
  const result = await onHost(host, args, { allowFailure: true });
  if (result.exitCode === 0 && result.stdout.trim() !== "") {
    write(runDir, "ptquery.txt", result.stdout);
  } else {
    write(
      runDir,
      "ptquery.txt",
      `未取得: pt-query-digest を実行できませんでした (${host}, exit=${result.exitCode})\n` +
        `導入は #11。slow query log が有効になっているかも確認すること。\n\n${result.stderr}`,
    );
    log.warn(`pt-query-digest を実行できませんでした (${host})`);
  }
};

const collectPprof = async (cfg: Config, runDir: string): Promise<void> => {
  if (!cfg.metrics.pprofEnabled) {
    write(runDir, "pprof.txt", "未取得: metrics.pprofEnabled = false（アプリに pprof を仕込んでいない）");
    return;
  }
  if (cfg.metrics.pprofUrl === "") {
    write(runDir, "pprof.txt", "未取得: metrics.pprofUrl が未設定");
    return;
  }

  // pprof の取得は手元で行う。競技サーバに go tool を要求しないため。
  const result = await local("go", ["tool", "pprof", "-top", "-nodecount=40", cfg.metrics.pprofUrl], {
    allowFailure: true,
  });
  if (result.exitCode === 0 && result.stdout.trim() !== "") {
    write(runDir, "pprof.txt", result.stdout);
  } else {
    write(
      runDir,
      "pprof.txt",
      `未取得: go tool pprof を実行できませんでした (exit=${result.exitCode})\n` +
        `  ${cfg.metrics.pprofUrl}\n\n${result.stderr}`,
    );
    log.warn("pprof を取得できませんでした");
  }
};

/**
 * ベンチ後の計測を集める。
 * pprof はベンチ中に取らないと意味が無いように見えるが、pprofUrl は
 * `?seconds=N` 付きで「その場で N 秒プロファイルする」形を想定している。
 * ベンチ中に取りたい場合は bench 側から呼ぶ（#4 で詰める）。
 */
export const collectMetrics = async (cfg: Config, runDir: string): Promise<void> => {
  await Promise.all([
    collectAlp(cfg, runDir),
    collectSlowQuery(cfg, runDir),
    collectPprof(cfg, runDir),
  ]);
  log.info("計測結果を回収しました");
};
