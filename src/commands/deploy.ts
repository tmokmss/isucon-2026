import { createHash } from "node:crypto";
import { existsSync, mkdirSync, readdirSync, rmSync } from "node:fs";
import { hostname, userInfo } from "node:os";
import { join, sep } from "node:path";
import { setTimeout as delay } from "node:timers/promises";
import { execa } from "execa";
import { type Config, loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { EXIT, HarnessError, usageError } from "../lib/exit.js";
import { resolveSha } from "../lib/git.js";
import { log, out } from "../lib/log.js";
import { repoRoot } from "../lib/paths.js";
import { copyTo, local, onHost, onHostShell, shQuote } from "../lib/ssh.js";

const deployError = (message: string): HarnessError => new HarnessError(EXIT.DEPLOY, message);

/** サーバ上で「今どの sha が動いているか」を記録するファイル（ログインユーザの $HOME 直下）。 */
const DEPLOYED_SHA_FILE = "$HOME/.deployed_sha";

// ---------------------------------------------------------------------------
// ビルド
// ---------------------------------------------------------------------------

/**
 * ビルド成果物の置き場。sha + ビルド設定ごとにキャッシュする。
 *
 * 同じ sha を再度デプロイするとき（= ロールバック）にビルドを丸ごと省ける。
 * 終盤に「とにかく速く直前の状態へ戻す」のが効くので、ここは効かせておく価値がある。
 */
const buildDirFor = (cfg: Config, sha: string): string => {
  const key = createHash("sha256")
    .update([cfg.app.langDir, cfg.app.binary, cfg.app.buildCommand].join("\u0000"))
    .digest("hex")
    .slice(0, 8);
  return join(repoRoot(), ".isuh", "build", `${sha.slice(0, 12)}-${key}`);
};

/**
 * 指定 sha のツリーを取り出す。
 *
 * 作業ツリーではなく `git archive` の出力を使う。並列にアプリを編集していても
 * デプロイされる内容が sha と一致することを保証するため。
 */
const exportTree = async (repoLocal: string, sha: string, dest: string): Promise<void> => {
  rmSync(dest, { recursive: true, force: true });
  mkdirSync(dest, { recursive: true });
  const tar = `${dest}.tar`;
  await local("git", ["-C", repoLocal, "archive", "--format=tar", "-o", tar, sha]);
  await local("tar", ["-xf", tar, "-C", dest]);
  rmSync(tar, { force: true });
};

/** ビルドして、成果物のローカルパスを返す。 */
const buildBinary = async (cfg: Config, sha: string): Promise<string> => {
  const dir = buildDirFor(cfg, sha);
  const binary = join(dir, cfg.app.langDir, cfg.app.binary);
  if (existsSync(binary)) {
    log.info(`build: キャッシュを使用 ${binary}`);
    return binary;
  }

  log.info(`build: ${sha.slice(0, 7)} を取り出しています`);
  await exportTree(cfg.app.repoLocal, sha, dir);

  const cwd = join(dir, cfg.app.langDir);
  if (!existsSync(cwd)) {
    rmSync(dir, { recursive: true, force: true });
    throw deployError(`app.langDir がリポジトリにありません: ${cfg.app.langDir} (at ${sha.slice(0, 7)})`);
  }

  log.info(`build: ${cfg.app.buildCommand}`);
  const startedAt = Date.now();
  try {
    // buildCommand は設定由来のシェル文字列（`GOOS=linux go build ...`）なので shell で実行する。
    // ssh / git のような引数の組み立てではないため、ここだけは配列引数にできない。
    await execa(cfg.app.buildCommand, { cwd, shell: true, stdout: "pipe", stderr: "pipe" });
  } catch (error) {
    const detail = error instanceof Error && "stderr" in error ? String(error.stderr) : String(error);
    rmSync(dir, { recursive: true, force: true });
    throw deployError(`ビルドに失敗しました\n${detail.trimEnd()}`);
  }

  if (!existsSync(binary)) {
    rmSync(dir, { recursive: true, force: true });
    throw deployError(
      `ビルドは成功しましたが成果物がありません: ${cfg.app.binary}\n` +
        `  app.binary と app.buildCommand の出力先が一致しているか確認してください。`,
    );
  }
  log.info(`build: 完了 ${((Date.now() - startedAt) / 1000).toFixed(1)}s`);
  return binary;
};

// ---------------------------------------------------------------------------
// アプリ以外の設定ファイル（conf/）
// ---------------------------------------------------------------------------

/**
 * conf/ のレイアウト。サーバ上の絶対パスをそのまま写す。
 *
 *   conf/_all/etc/sysctl.d/99-isucon.conf  → 全ホストの /etc/sysctl.d/99-isucon.conf
 *   conf/isu1/etc/nginx/nginx.conf         → isu1 の /etc/nginx/nginx.conf
 *
 * 同じパスが両方にあればホスト指定側が勝つ。中身の git 管理（サーバからの取り込み）は #10。
 */
interface ConfFile {
  localPath: string;
  remotePath: string;
}

const walkFiles = (dir: string, prefix = ""): string[] => {
  const found: string[] = [];
  for (const entry of readdirSync(dir, { withFileTypes: true })) {
    if (entry.name.startsWith(".")) continue;
    const rel = prefix === "" ? entry.name : `${prefix}${sep}${entry.name}`;
    if (entry.isDirectory()) found.push(...walkFiles(join(dir, entry.name), rel));
    else found.push(rel);
  }
  return found;
};

const confFilesFor = (host: string): ConfFile[] => {
  const byRemote = new Map<string, ConfFile>();
  for (const scope of ["_all", host]) {
    const base = join(repoRoot(), "conf", scope);
    if (!existsSync(base)) continue;
    for (const rel of walkFiles(base)) {
      const remotePath = `/${rel.split(sep).join("/")}`;
      byRemote.set(remotePath, { localPath: join(base, rel), remotePath });
    }
  }
  return [...byRemote.values()].sort((a, b) => a.remotePath.localeCompare(b.remotePath));
};

/**
 * 配った設定ファイルから、反映に必要なサービス操作を決める。
 *
 * 明示的に対応づけたものだけを触る。MySQL の restart は重いので、my.cnf を
 * 配ったときにだけ走る。
 */
const reloadsFor = (
  cfg: Config,
  remotePaths: readonly string[],
): { service: string; action: "reload" | "restart" }[] => {
  const hits = new Map<string, "reload" | "restart">();
  for (const path of remotePaths) {
    if (path.startsWith("/etc/nginx/")) hits.set(cfg.middleware.nginxService, "reload");
    else if (path.startsWith("/etc/mysql/") || path === "/etc/my.cnf")
      hits.set(cfg.middleware.dbService, "restart");
  }
  return [...hits].map(([service, action]) => ({ service, action }));
};

const pushConf = async (cfg: Config, host: string, files: readonly ConfFile[]): Promise<void> => {
  for (const [index, file] of files.entries()) {
    const staged = `/tmp/isuh-conf-${index}`;
    await copyTo(host, file.localPath, staged);
    // install -D なら親ディレクトリごと作る。/etc 配下なので sudo が要る。
    await onHost(host, ["sudo", "install", "-D", "-m", "0644", staged, file.remotePath]);
    await onHost(host, ["rm", "-f", staged]);
  }

  const paths = files.map((f) => f.remotePath);
  if (paths.some((p) => p.startsWith("/etc/systemd/"))) {
    await onHost(host, ["sudo", "systemctl", "daemon-reload"]);
  }
  for (const { service, action } of reloadsFor(cfg, paths)) {
    log.info(`${host}: ${action} ${service}`);
    await onHost(host, ["sudo", "systemctl", action, service]);
  }
};

// ---------------------------------------------------------------------------
// 配布・再起動
// ---------------------------------------------------------------------------

const remoteBinaryPath = (cfg: Config): string =>
  `${cfg.app.dir.replace(/\/$/, "")}/${cfg.app.langDir}/${cfg.app.binary}`;

const pushBinary = async (cfg: Config, host: string, localBinary: string): Promise<void> => {
  const dest = remoteBinaryPath(cfg);
  const staged = `${dest}.isuh-new`;
  await copyTo(host, localBinary, staged);
  // 実行中のバイナリを直接上書きすると ETXTBSY になる。rename なら原子的に差し替わる。
  await onHostShell(host, `chmod +x ${shQuote(staged)} && mv -f ${shQuote(staged)} ${shQuote(dest)}`);
};

const restartApp = async (cfg: Config, host: string): Promise<void> => {
  await onHost(host, ["sudo", "systemctl", "restart", cfg.app.service]);
  // 起動直後に落ちる事故（設定ミス・ポート衝突）は is-active を一拍おいて見ると捕まる。
  await delay(800);
  const result = await onHost(host, ["systemctl", "is-active", cfg.app.service], {
    allowFailure: true,
  });
  const state = result.stdout.trim();
  if (state !== "active") {
    const status = await onHostShell(
      host,
      `sudo systemctl status --no-pager -l ${shQuote(cfg.app.service)} 2>&1 | tail -30`,
      { allowFailure: true },
    );
    throw deployError(
      `${host}: ${cfg.app.service} が起動していません (is-active=${state || "unknown"})\n${status.stdout}`,
    );
  }
};

const recordSha = async (cfg: Config, host: string, sha: string, ref: string): Promise<void> => {
  const line = [sha, ref, new Date().toISOString(), `${userInfo().username}@${hostname()}`].join("\t");
  await onHostShell(host, `printf '%s\\n' ${shQuote(line)} > "${DEPLOYED_SHA_FILE}"`);
};

/** そのホストで今動いている sha。未記録なら null。rollback(#6) / report(#4) から使う。 */
export const deployedSha = async (host: string): Promise<string | null> => {
  const result = await onHostShell(host, `cat "${DEPLOYED_SHA_FILE}" 2>/dev/null || true`, {
    allowFailure: true,
  });
  const sha = result.stdout.split("\n")[0]?.split("\t")[0]?.trim();
  return sha ? sha : null;
};

// ---------------------------------------------------------------------------
// 本体
// ---------------------------------------------------------------------------

export interface DeployOptions {
  ref?: string;
  dryRun?: boolean;
}

/** deploy の本体。bench(#2) から直接呼ぶ。戻り値はデプロイした sha。 */
export const deploy = async (cfg: Config, opts: DeployOptions = {}): Promise<string> => {
  const ref = opts.ref ?? "HEAD";
  const sha = await resolveSha(cfg.app.repoLocal, ref);
  const hosts = cfg.hosts.app;
  const startedAt = Date.now();

  const dirty = await local("git", ["-C", cfg.app.repoLocal, "status", "--porcelain"], {
    allowFailure: true,
  });
  if (dirty.stdout.trim() !== "") {
    const files = dirty.stdout.trim().split("\n").length;
    log.warn(
      `アプリリポジトリに未コミットの変更が ${files} 件あります。` +
        `デプロイされるのは ${sha.slice(0, 7)} の内容だけです。`,
    );
  }

  if (opts.dryRun) {
    log.info(`[dry-run] ref=${ref} sha=${sha.slice(0, 7)}`);
    log.info(`[dry-run] build: ${cfg.app.buildCommand} (in ${cfg.app.langDir})`);
    for (const host of hosts) {
      log.info(`[dry-run] ${host}: ${remoteBinaryPath(cfg)} を差し替え`);
      const confFiles = confFilesFor(host);
      for (const file of confFiles) log.info(`[dry-run] ${host}: ${file.remotePath} を配布`);
      for (const { service, action } of reloadsFor(cfg, confFiles.map((f) => f.remotePath))) {
        log.info(`[dry-run] ${host}: ${action} ${service}`);
      }
      log.info(`[dry-run] ${host}: restart ${cfg.app.service}`);
    }
    return sha;
  }

  const binary = await buildBinary(cfg, sha);

  // ホストごとに直列（配布 → 再起動）、ホスト間は並列。
  // 1 台でも失敗したら全体を失敗にする（後続のベンチを走らせないため）。
  const results = await Promise.allSettled(
    hosts.map(async (host) => {
      const confFiles = confFilesFor(host);
      if (confFiles.length > 0) {
        log.info(`${host}: 設定ファイル ${confFiles.length} 件を配布`);
        await pushConf(cfg, host, confFiles);
      }
      await pushBinary(cfg, host, binary);
      await restartApp(cfg, host);
      await recordSha(cfg, host, sha, ref);
      log.info(`${host}: ok`);
    }),
  );

  const failures = results.flatMap((result, index) =>
    result.status === "rejected" ? [{ host: hosts[index]!, reason: result.reason }] : [],
  );
  if (failures.length > 0) {
    const detail = failures
      .map(({ host, reason }) => `  ${host}: ${reason instanceof Error ? reason.message : String(reason)}`)
      .join("\n");
    throw deployError(`デプロイに失敗しました (${failures.length}/${hosts.length} 台)\n${detail}`);
  }

  log.info(
    `deploy 完了 sha=${sha.slice(0, 7)} hosts=${hosts.length} ${((Date.now() - startedAt) / 1000).toFixed(1)}s`,
  );
  return sha;
};

export const command: Command = {
  summary: "指定した git ref をビルドして全台へ配布・再起動する",
  usage: `usage: isuh deploy [<git-ref>] [--dry-run]

指定した git ref のアプリをビルドし、hosts.app へ配布して再起動する。
ref は app.repoLocal（手元のアプリリポジトリ）で解決する。既定は HEAD。

  <git-ref>   デプロイする ref
  --dry-run   何をするかを表示して実行しない

デプロイされるのは作業ツリーではなく ref が指す commit の内容（git archive）。
未コミットの変更があれば警告するが、配られはしない。

各ホストの $HOME/.deployed_sha に「sha / ref / 時刻 / 実行者」を記録する。
conf/_all/** と conf/<host>/** はサーバ上の絶対パスへそのまま配布し、
nginx / MySQL の設定を含む場合は該当サービスを reload / restart する。

stdout にはデプロイした sha だけを出す（SHA=$(isuh deploy) が使える）。

終了コード: 0 成功 / 20 デプロイ失敗 / 64 引数不正 / 65 設定不備`,

  async run(argv) {
    let ref = "HEAD";
    let dryRun = false;
    for (const arg of argv) {
      if (arg === "--dry-run") dryRun = true;
      else if (arg.startsWith("-")) throw usageError(`不明なオプション: ${arg}`);
      else ref = arg;
    }
    const cfg = await loadConfig();
    log.info(`deploy ref=${ref} hosts=[${cfg.hosts.app.join(" ")}] dryRun=${dryRun}`);
    const sha = await deploy(cfg, { ref, dryRun });
    out(sha);
  },
};
