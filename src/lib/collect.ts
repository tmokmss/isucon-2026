import { existsSync, mkdirSync, statSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import { setTimeout as delay } from "node:timers/promises";
import type { Config } from "../config.js";
import { log } from "./log.js";
import { copyFrom, local, onHost, onHostShell, shQuote } from "./ssh.js";

/**
 * 計測の生データを集める — docs/DESIGN.md §4
 *
 * ここが作るのは runs/<run-id>/ 以下の生ファイルだけ。
 * それを読んで report.md に整形するのは src/lib/report.ts。
 *
 * 計測ツールが入っていないことは普通に起きる（導入は #11、pprof は仕込むかどうか自体が
 * 判断）。その場合は落とさず「未取得」と理由をファイルに書く。ベンチ 1 回は高価なので、
 * 補助的な計測が無いことでランを失敗にはしない。
 */

const write = (dir: string, name: string, body: string): void => {
  writeFileSync(join(dir, name), body.endsWith("\n") ? body : `${body}\n`);
};

const errorText = (error: unknown): string => (error instanceof Error ? error.message : String(error));

/**
 * PostgreSQL には slow query log の代わりに pg_stat_statements を使う。
 * 集計はサーバ全体で 1 つなので、拡張は postgres DB に作り、アプリの DB の分だけを読む。
 */
export const usesPostgres = (cfg: Config): boolean => cfg.middleware.dbService.startsWith("postgresql");

export const psqlOn = (host: string, sql: string, opts: { allowFailure?: boolean } = {}) =>
  onHost(host, ["sudo", "-u", "postgres", "psql", "-d", "postgres", "-XAtc", sql], opts);

// ---------------------------------------------------------------------------
// リモートで常駐させる計測プロセス
// ---------------------------------------------------------------------------

/**
 * リモートでコマンドをバックグラウンド起動し、PID を控える。
 * コマンドを直接 background にする（サブシェルを挟むと $! が本体の PID でなくなる）。
 * ssh はすぐ戻るので、手元に長生きするプロセスは残らない。
 */
const startRemote = async (
  host: string,
  command: string,
  out: string,
  pid: string,
  preamble = "",
): Promise<boolean> => {
  const script = [
    `rm -f ${shQuote(out)} ${shQuote(pid)}`,
    preamble === "" ? `: > ${shQuote(out)}` : `{ ${preamble}; } > ${shQuote(out)} 2>&1`,
    `nohup ${command} >> ${shQuote(out)} 2>&1 &`,
    `echo $! > ${shQuote(pid)}`,
  ].join("\n");
  const result = await onHostShell(host, script, { allowFailure: true });
  return result.exitCode === 0;
};

const stopRemote = async (host: string, pid: string): Promise<void> => {
  await onHostShell(
    host,
    `if [ -f ${shQuote(pid)} ]; then kill "$(cat ${shQuote(pid)})" 2>/dev/null || true; fi; rm -f ${shQuote(pid)}`,
    { allowFailure: true },
  );
};

/** リモートのプロセスがまだ動いているか。 */
const remoteAlive = async (host: string, pid: string): Promise<boolean> => {
  const result = await onHostShell(
    host,
    `[ -f ${shQuote(pid)} ] && [ -d "/proc/$(cat ${shQuote(pid)})" ]`,
    { allowFailure: true },
  );
  return result.exitCode === 0;
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

  if (usesPostgres(cfg) && cfg.hosts.db !== "") {
    const result = await psqlOn(cfg.hosts.db, "SELECT pg_stat_statements_reset()", { allowFailure: true });
    if (result.exitCode !== 0) log.warn(`${cfg.hosts.db}: pg_stat_statements を空にできません（isuh setup --only logs が未実行？）`);
  }

  await Promise.all(
    targets.map(async ({ host, path, what }) => {
      if (host === "" || path === "") return;
      // -c: 無いファイルは作らない。root 所有で作ると mysqld が slow log を開けず、記録が止まる
      const result = await onHost(host, ["sudo", "truncate", "-c", "-s", "0", path], {
        allowFailure: true,
      });
      if (result.exitCode !== 0) {
        log.warn(`${host}: ${what} を空にできません (${path})`);
      }
    }),
  );
};

// ---------------------------------------------------------------------------
// ベンチ中: vmstat / top（全台）
// ---------------------------------------------------------------------------

const VMSTAT_OUT = "/tmp/isuh-resource.txt";
const VMSTAT_PID = "/tmp/isuh-resource.pid";
const TOP_OUT = "/tmp/isuh-top.txt";
const TOP_PID = "/tmp/isuh-top.pid";

/**
 * vmstat はホスト全体の値しか出さない。アプリと MySQL が同じホストにいると
 * どちらが CPU を食っているか分からないので、top でプロセス単位も取る。
 * top は procps に入っていて、sysstat（pidstat）と違い必ずある。
 */
const startHostSampling = async (cfg: Config, dir: string): Promise<() => Promise<void>> => {
  const hosts = cfg.hosts.all;
  const interval = Math.max(1, Math.floor(cfg.metrics.resourceIntervalSec));
  mkdirSync(join(dir, "top"), { recursive: true });

  await Promise.all(
    hosts.map(async (host) => {
      const vmstat = await startRemote(
        host,
        `vmstat -n ${interval}`,
        VMSTAT_OUT,
        VMSTAT_PID,
        `echo "# host=$(hostname) nproc=$(nproc)"; free -m 2>/dev/null | sed 's/^/# /'`,
      );
      if (!vmstat) log.warn(`${host}: vmstat を開始できません`);
      // -w で COMMAND が切れないようにし、-o で CPU 順に並べる
      const top = await startRemote(host, `top -b -d ${interval} -w 512 -o %CPU`, TOP_OUT, TOP_PID);
      if (!top) log.warn(`${host}: top を開始できません`);
    }),
  );
  log.info(`リソース計測を開始しました (vmstat + top, ${interval}s 間隔, ${hosts.length} 台)`);

  return async () => {
    await Promise.all(
      hosts.map(async (host) => {
        await Promise.all([stopRemote(host, VMSTAT_PID), stopRemote(host, TOP_PID)]);
        await onHostShell(host, `{ echo; echo "# uptime: $(uptime)"; } >> ${shQuote(VMSTAT_OUT)}`, {
          allowFailure: true,
        });
        const [gotVmstat, gotTop] = await Promise.all([
          copyFrom(host, VMSTAT_OUT, join(dir, `${host}.txt`), { allowFailure: true }),
          copyFrom(host, TOP_OUT, join(dir, "top", `${host}.txt`), { allowFailure: true }),
        ]);
        if (!gotVmstat) write(dir, `${host}.txt`, `未取得: ${host} から vmstat の結果を回収できませんでした`);
        if (!gotTop) write(join(dir, "top"), `${host}.txt`, `未取得: ${host} から top の結果を回収できませんでした`);
      }),
    );
    log.info("リソース計測を回収しました");
  };
};

// ---------------------------------------------------------------------------
// ベンチ中: pprof
// ---------------------------------------------------------------------------

const PPROF_OUT = "/tmp/isuh-cpu.pprof";
const PPROF_PID = "/tmp/isuh-pprof.pid";
const PPROF_ERR = "/tmp/isuh-pprof.err";

/** pprofUrl の ?seconds=N。無ければ pprof の既定の 30 秒。 */
const profileSeconds = (url: string): number => {
  const seconds = Number(/[?&]seconds=(\d+)/.exec(url)?.[1] ?? "30");
  return Number.isFinite(seconds) && seconds > 0 ? seconds : 30;
};

/**
 * CPU プロファイルを**ベンチの開始と同時に**取り始める。
 *
 * /debug/pprof/profile?seconds=N は「叩いてから N 秒間」を測る。ベンチの後に叩くと
 * アイドル中を測ることになり役に立たない。N はベンチの長さに合わせておく。
 *
 * 取得は hosts.app[0] 上で curl する（pprofUrl はそのホストから見た URL）。
 * pprof のポートはセキュリティグループで外に開いていないことが多いため。
 * 途中で接続を切るとプロファイルは返ってこないので、止めるときは N 秒経つまで待つ。
 */
const startProfiling = async (cfg: Config, dir: string): Promise<() => Promise<void>> => {
  if (!cfg.metrics.pprofEnabled) {
    write(dir, "pprof.txt", "未取得: metrics.pprofEnabled = false（アプリに pprof を仕込んでいない）");
    return async () => {};
  }
  const host = cfg.hosts.app[0];
  if (host === undefined || cfg.metrics.pprofUrl === "") {
    write(dir, "pprof.txt", "未取得: hosts.app / metrics.pprofUrl が未設定");
    return async () => {};
  }

  const url = cfg.metrics.pprofUrl;
  const seconds = profileSeconds(url);
  const startedAt = Date.now();
  const started = await onHostShell(
    host,
    [
      `rm -f ${shQuote(PPROF_OUT)} ${shQuote(PPROF_ERR)} ${shQuote(PPROF_PID)}`,
      `nohup curl -sS --fail --max-time ${seconds + 15} -o ${shQuote(PPROF_OUT)} ${shQuote(url)} > ${shQuote(PPROF_ERR)} 2>&1 &`,
      `echo $! > ${shQuote(PPROF_PID)}`,
    ].join("\n"),
    { allowFailure: true },
  );
  if (started.exitCode !== 0) {
    write(dir, "pprof.txt", `未取得: ${host} で pprof の取得を開始できませんでした\n${started.stderr}`);
    return async () => {};
  }
  log.info(`pprof: ${host} で ${seconds}s のプロファイルを開始しました`);

  return async () => {
    const remaining = Math.ceil((startedAt + seconds * 1000 - Date.now()) / 1000);
    if (remaining > 0) {
      log.info(`pprof: プロファイルの完了を待っています（残り約 ${remaining}s。seconds はベンチの長さに合わせる）`);
    }
    const deadline = startedAt + (seconds + 20) * 1000;
    while (Date.now() < deadline && (await remoteAlive(host, PPROF_PID))) {
      await delay(1000);
    }
    await stopRemote(host, PPROF_PID);

    const raw = join(dir, "cpu.pprof");
    const got = await copyFrom(host, PPROF_OUT, raw, { allowFailure: true });
    if (!got || !existsSync(raw) || statSync(raw).size === 0) {
      const err = await onHostShell(host, `cat ${shQuote(PPROF_ERR)} 2>/dev/null`, { allowFailure: true });
      write(
        dir,
        "pprof.txt",
        `未取得: プロファイルを取れませんでした (${host})\n  ${url}\n\n${err.stdout}`,
      );
      log.warn(`pprof を取得できませんでした (${host})`);
      return;
    }

    // 生のプロファイルも残す（go tool pprof -http で後から掘れる）。整形は手元で行う。
    const top = await local("go", ["tool", "pprof", "-top", "-nodecount=40", raw], { allowFailure: true });
    if (top.exitCode === 0 && top.stdout.trim() !== "") {
      write(dir, "pprof.txt", top.stdout);
      log.info("pprof: 回収しました");
    } else {
      write(
        dir,
        "pprof.txt",
        `未取得: go tool pprof で読めませんでした (exit=${top.exitCode})\n` +
          `  生のプロファイルは cpu.pprof にある\n\n${top.stderr}`,
      );
      log.warn("pprof を整形できませんでした");
    }
  };
};

// ---------------------------------------------------------------------------
// ベンチ中の計測の入口
// ---------------------------------------------------------------------------

export interface ResourceSampler {
  /** 計測を止めて runs/<run-id>/ に回収する。多重呼び出し安全。 */
  stop(): Promise<void>;
}

/**
 * ベンチの直前に呼び、ベンチの間ずっと測る。
 *   - vmstat: ホストごとの CPU / mem / IO（どのホストが飽和したか）
 *   - top:    プロセスごとの CPU（どのプロセスが食っているか）
 *   - pprof:  アプリのどの関数が食っているか（pprofEnabled のとき）
 */
export const startResourceSampling = async (
  cfg: Config,
  runDir: string,
): Promise<ResourceSampler> => {
  const dir = join(runDir, "resource");
  mkdirSync(dir, { recursive: true });

  const [stopHosts, stopProfile] = await Promise.all([
    startHostSampling(cfg, dir),
    startProfiling(cfg, runDir),
  ]);

  let stopped = false;
  return {
    async stop() {
      if (stopped) return;
      stopped = true;
      // ホストの計測はベンチが終わった瞬間に止め、pprof は N 秒経つのを待つ
      await Promise.all([
        stopHosts().catch((error) => log.warn(`リソース計測の回収に失敗: ${errorText(error)}`)),
        stopProfile().catch((error) => log.warn(`pprof の回収に失敗: ${errorText(error)}`)),
      ]);
    },
  };
};

// ---------------------------------------------------------------------------
// ベンチ後
// ---------------------------------------------------------------------------

const ALP_URI_LIMIT = 100000;

const alpOn = async (cfg: Config, host: string): Promise<string> => {
  const args = [
    "sudo",
    "alp",
    cfg.metrics.alpFormat,
    // 機械向けの TSV。report.md では src/lib/digest.ts が列を絞った表に畳む
    "--format",
    "tsv",
    "--file",
    cfg.middleware.nginxAccessLog,
    "--sort",
    cfg.metrics.alpSort,
    "-r",
    // alp の --limit は「表示する行数」ではなく「許容する URI の種類数」で、超えると
    // `Too many URI's` で集計ごと落ちる。alpMatches を埋める前（ID 入りのパスがばらける）でも
    // 取れるよう大きくしておき、上位 alpLimit 件への絞り込みは report.md 側で行う。
    "--limit",
    String(ALP_URI_LIMIT),
  ];
  if (cfg.metrics.alpMatches.length > 0) args.push("-m", cfg.metrics.alpMatches.join(","));

  const result = await onHost(host, args, { allowFailure: true });
  if (result.exitCode === 0 && result.stdout.trim() !== "") return result.stdout.trimEnd();

  log.warn(`alp を実行できませんでした (${host})`);
  return (
    `未取得: alp を実行できませんでした (${host}, exit=${result.exitCode})\n` +
    `導入は #11。手元で確認するには:\n  ssh ${host} ${args.join(" ")}\n\n${result.stderr.trimEnd()}`
  );
};

/**
 * nginx が複数台なら全台で集計する。
 * alp はサーバ側にしか無いので 1 本に合算はせず、ホストごとに並べる。
 */
const collectAlp = async (cfg: Config, runDir: string): Promise<void> => {
  const hosts = cfg.hosts.nginx;
  if (hosts.length === 0 || cfg.middleware.nginxAccessLog === "") {
    write(runDir, "alp.txt", "未取得: hosts.nginx / middleware.nginxAccessLog が未設定");
    return;
  }
  if (cfg.metrics.alpMatches.length === 0) {
    log.warn("metrics.alpMatches が空です。URI が正規化されず、集計が意味を持ちません");
  }

  const outputs = await Promise.all(hosts.map((host) => alpOn(cfg, host)));
  const body =
    hosts.length === 1
      ? outputs[0]!
      : hosts.map((host, i) => `== ${host} ==\n${outputs[i]}`).join("\n\n");
  write(runDir, "alp.txt", body);
};

const pgStatSql = (limit: number): string => `
WITH s AS (
  SELECT p.* FROM pg_stat_statements p JOIN pg_database d ON d.oid = p.dbid WHERE d.datname <> 'postgres'
)
SELECT json_build_object(
  'source', 'pg_stat_statements',
  'total_ms', coalesce(sum(total_exec_time), 0),
  'calls', coalesce(sum(calls), 0),
  'unique', count(*),
  'top', (SELECT coalesce(json_agg(t), '[]') FROM (
    SELECT calls, total_exec_time AS total_ms, mean_exec_time AS mean_ms, max_exec_time AS max_ms,
           rows, shared_blks_hit + shared_blks_read AS blks, query
    FROM s ORDER BY total_exec_time DESC LIMIT ${limit}
  ) t)
) FROM s`;

const collectPgStat = async (cfg: Config, runDir: string): Promise<void> => {
  const host = cfg.hosts.db;
  const result = await psqlOn(host, pgStatSql(cfg.metrics.ptQueryLimit), { allowFailure: true });
  if (result.exitCode === 0 && result.stdout.trim() !== "") {
    write(runDir, "ptquery.txt", result.stdout);
  } else {
    write(
      runDir,
      "ptquery.txt",
      `未取得: pg_stat_statements を読めませんでした (${host}, exit=${result.exitCode})\n` +
        `isuh setup --only logs で有効にしたか確認すること。\n\n${result.stderr}`,
    );
    log.warn(`pg_stat_statements を読めませんでした (${host})`);
  }
};

const collectSlowQuery = async (cfg: Config, runDir: string): Promise<void> => {
  const host = cfg.hosts.db;
  if (host !== "" && usesPostgres(cfg)) {
    await collectPgStat(cfg, runDir);
    return;
  }
  if (host === "" || cfg.middleware.dbSlowLog === "") {
    write(runDir, "ptquery.txt", "未取得: hosts.db / middleware.dbSlowLog が未設定");
    return;
  }

  const args = [
    "sudo",
    "pt-query-digest",
    // 機械向けの JSON。report.md では src/lib/digest.ts が 1 行 1 クエリの表に畳む
    "--output",
    "json",
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

/** ベンチ後に、ログを集計する（alp / pt-query-digest）。pprof はベンチ中に取っている。 */
export const collectMetrics = async (cfg: Config, runDir: string): Promise<void> => {
  await Promise.all([collectAlp(cfg, runDir), collectSlowQuery(cfg, runDir)]);
  log.info("計測結果を回収しました");
};
