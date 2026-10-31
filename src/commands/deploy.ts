import { createHash } from "node:crypto";
import { existsSync, mkdirSync, readFileSync, readdirSync, renameSync, rmSync, writeFileSync } from "node:fs";
import { hostname, userInfo } from "node:os";
import { join, sep } from "node:path";
import { setTimeout as delay } from "node:timers/promises";
import { execa } from "execa";
import { type Config, loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { EXIT, HarnessError, usageError } from "../lib/exit.js";
import { resolveSha } from "../lib/git.js";
import { acquireLock } from "../lib/lock.js";
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
 * 同じ sha を再度デプロイするとき（= 前の版へ戻すとき）にビルドを丸ごと省ける。
 * 終盤に「とにかく速く直前の状態へ戻す」のが効くので、ここは効かせておく価値がある。
 */
const buildDirFor = (cfg: Config, sha: string): string => {
  const key = createHash("sha256")
    .update([cfg.app.localDir, cfg.app.langDir, cfg.app.binary, cfg.app.buildCommand].join("\u0000"))
    .digest("hex")
    .slice(0, 8);
  return join(repoRoot(), ".isuh", "build", `${sha.slice(0, 12)}-${key}`);
};

/** 取り出しが最後まで終わった印。途中で落ちたツリーを使わないため。 */
const TREE_DONE = ".isuh-tree-done";

/**
 * 指定 sha のツリーを取り出す（取り出し済みならそれを使う）。
 *
 * 作業ツリーではなく `git archive` の出力を使う。並列に編集していても
 * デプロイされる内容（アプリも conf/ も）が sha と一致することを保証するため。
 */
const treeFor = async (cfg: Config, sha: string): Promise<string> => {
  const dest = buildDirFor(cfg, sha);
  if (existsSync(join(dest, TREE_DONE))) return dest;
  log.info(`build: ${sha.slice(0, 7)} を取り出しています`);
  rmSync(dest, { recursive: true, force: true });
  mkdirSync(dest, { recursive: true });
  const tar = `${dest}.tar`;
  await local("git", ["-C", cfg.app.repoLocal, "archive", "--format=tar", "-o", tar, sha]);
  await local("tar", ["-xf", tar, "-C", dest]);
  rmSync(tar, { force: true });
  writeFileSync(join(dest, TREE_DONE), sha);
  return dest;
};

/** ビルドして、成果物のローカルパスを返す。 */
const buildBinary = async (cfg: Config, sha: string, tree: string): Promise<string> => {
  const cwd = join(tree, cfg.app.localDir, cfg.app.langDir);
  const binary = join(cwd, cfg.app.binary);
  if (existsSync(binary)) {
    log.info(`build: キャッシュを使用 ${binary}`);
    return binary;
  }

  if (!existsSync(cwd)) {
    throw deployError(
      `アプリのディレクトリがリポジトリにありません: ${join(cfg.app.localDir, cfg.app.langDir)} (at ${sha.slice(0, 7)})\n` +
        `  app.localDir / app.langDir を確認するか、isuh setup --only apprepo で取り込んでください。`,
    );
  }

  log.info(`build: ${cfg.app.buildCommand}`);
  const startedAt = Date.now();
  try {
    // buildCommand は設定由来のシェル文字列（`GOOS=linux go build ...`）なので shell で実行する。
    // ssh / git のような引数の組み立てではないため、ここだけは配列引数にできない。
    await execa(cfg.app.buildCommand, { cwd, shell: true, stdout: "pipe", stderr: "pipe" });
  } catch (error) {
    const detail = error instanceof Error && "stderr" in error ? String(error.stderr) : String(error);
    rmSync(binary, { force: true });
    throw deployError(`ビルドに失敗しました\n${detail.trimEnd()}`);
  }

  if (!existsSync(binary)) {
    throw deployError(
      `ビルドは成功しましたが成果物がありません: ${cfg.app.binary}\n` +
        `  app.binary と app.buildCommand の出力先が一致しているか確認してください。`,
    );
  }
  log.info(`build: 完了 ${((Date.now() - startedAt) / 1000).toFixed(1)}s`);
  return binary;
};

// ---------------------------------------------------------------------------
// アプリのバイナリ以外のファイル（スキーマ・スクリプト・テンプレートなど）
// ---------------------------------------------------------------------------

/** バイナリの app.localDir からの相対パス。同期から外す（pushBinary が原子的に差し替えるため）。 */
const binaryRelPath = (cfg: Config): string => `${cfg.app.langDir.replace(/\/+$/, "")}/${cfg.app.binary}`;

/** app.localDir 以下を固めた tar。ツリーと同じ場所に sha ごとにキャッシュする。 */
const APP_TAR = ".isuh-app.tar";

/**
 * 最後に同期した app.localDir のツリー ID と、そのときデプロイした sha を記録するファイル
 * （サーバのログインユーザの $HOME 直下。`<ツリー ID> <TAB> <sha>`）。
 */
const APP_TREE_FILE = "$HOME/.isuh-app-tree";

/**
 * app.localDir 以下を sha の中身で tar に固める。
 *
 * `git archive <sha>:<localDir>` を使うので、中身は sha と一致し、実行ビット（sql/init.sh 等）も
 * git の記録どおりになる。取り出したツリーを手元の tar（macOS は bsdtar）で固め直さないのは、
 * 拡張属性などのヘッダを混ぜてサーバの GNU tar に警告を出させないため。
 */
const appTarFor = async (cfg: Config, sha: string, tree: string): Promise<string> => {
  const tar = join(tree, APP_TAR);
  if (existsSync(tar)) return tar;
  const localDir = cfg.app.localDir.replace(/\/+$/, "");
  const treeish = localDir === "" || localDir === "." ? sha : `${sha}:${localDir}`;
  const partial = `${tar}.${process.pid}.partial`;
  try {
    await local("git", ["-C", cfg.app.repoLocal, "archive", "--format=tar", "-o", partial, treeish]);
  } catch (error) {
    rmSync(partial, { force: true });
    const detail = error instanceof Error && "stderr" in error ? String(error.stderr) : String(error);
    throw deployError(`${cfg.app.localDir} を取り出せません (at ${sha.slice(0, 7)})\n${detail.trimEnd()}`);
  }
  renameSync(partial, tar);
  return tar;
};

/**
 * app.localDir 以下を app.dir へ同期する（上書きのみ。削除はしない）。
 *
 * - 削除しないのは、サーバ側にしか無いもの（gitignore されたビルド成果物・node_modules・
 *   rust/target・アップロードされた画像など）を消さないため。リポジトリで消したファイルは残る
 * - 同期するときはすべて書き直すので、mtime の比較による取りこぼしは起きない
 * - 前回同期したツリー ID と同じで、その後に別の deploy が入っていなければ、同期そのものを省く
 *   （conf/ だけを変えたとき等）
 * - バイナリは外す。実行中のファイルを上書きすると ETXTBSY になるため
 * - rsync ではなく tar にしているのは、サーバに追加の道具が要らないため（検証用コンテナにも
 *   rsync は無い）。手元の rsync も macOS では openrsync で、GNU rsync と挙動が揃わない
 */
const syncAppFiles = async (
  cfg: Config,
  host: string,
  appTar: string,
  treeId: string,
  sha: string,
): Promise<void> => {
  // 省いてよいのは「記録したツリーが同じ」かつ「その後に別の deploy が入っていない」とき。
  // 別の deploy（記録を書かない古いハーネスや、別の worktree）が入ると .deployed_sha が変わる。
  // そのときはサーバのファイルが記録と違いうるので、同期し直す
  const state = await onHostShell(
    host,
    `cat "${APP_TREE_FILE}" 2>/dev/null || true; echo; cut -f1 "${DEPLOYED_SHA_FILE}" 2>/dev/null || true`,
    { allowFailure: true },
  );
  const [recordedLine = "", deployedNow = ""] = state.stdout.split("\n").filter((line) => line.trim() !== "");
  const [recordedTree, recordedSha] = recordedLine.split("\t");
  if (recordedTree === treeId && recordedSha !== undefined && recordedSha === deployedNow.trim()) {
    log.info(`${host}: ${cfg.app.localDir}/ は前回の同期から変更なし`);
    return;
  }
  const dir = cfg.app.dir.replace(/\/$/, "");
  const staged = `/tmp/isuh-app-${process.pid}.tar`;
  const exclude = binaryRelPath(cfg);
  await copyTo(host, appTar, staged);
  // 実行ビットは git の記録から残る（root 以外の展開は umask を掛けるだけ。0664 → 0644、0775 → 0755）。
  // -m: mtime は今にする（commit 時刻で古く見えないように）
  // --no-overwrite-dir: 既存ディレクトリの権限を変えない
  const script =
    `mkdir -p ${shQuote(dir)} && ` +
    `tar -x -m --no-overwrite-dir -f ${shQuote(staged)} -C ${shQuote(dir)} ` +
    `--anchored --exclude=${shQuote(exclude)} --exclude=${shQuote(`./${exclude}`)}`;
  try {
    await onHostShell(host, script);
  } finally {
    await onHost(host, ["rm", "-f", staged], { allowFailure: true });
  }
  // 同期が最後まで済んでから記録する（途中で落ちたら次回は同期し直す）
  // sha は、この deploy が最後に .deployed_sha に書く値。途中で落ちて書かれなければ、次回は同期し直す
  await onHostShell(host, `printf '%s\\t%s\\n' ${shQuote(treeId)} ${shQuote(sha)} > "${APP_TREE_FILE}"`);
};

/**
 * app.localDir のツリー ID（git の tree オブジェクト）。中身が同じなら sha が違っても同じになる。
 * 同期を省くかどうかの判定に使う。
 */
const appTreeIdFor = async (cfg: Config, sha: string): Promise<string> => {
  const localDir = cfg.app.localDir.replace(/\/+$/, "");
  const treeish = localDir === "" || localDir === "." ? `${sha}^{tree}` : `${sha}:${localDir}`;
  const result = await local("git", ["-C", cfg.app.repoLocal, "rev-parse", treeish]);
  return result.stdout.trim();
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
 * 同じパスが両方にあればホスト指定側が勝つ。読むのは app.repoLocal の作業ツリーではなく、
 * バイナリと同じ sha のツリー（deploy <古い sha> で設定も一緒に戻る）。
 * 中身はサーバから `isuh setup --only conf` で取り込む。
 */
interface ConfFile {
  localPath: string;
  /** ツリーの中の相対パス（git archive の中のパス。`conf/isu1/etc/nginx/nginx.conf`） */
  treePath: string;
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

const confFilesFor = (tree: string, host: string): ConfFile[] => {
  const byRemote = new Map<string, ConfFile>();
  for (const scope of ["_all", host]) {
    const base = join(tree, "conf", scope);
    if (!existsSync(base)) continue;
    for (const rel of walkFiles(base)) {
      const posix = rel.split(sep).join("/");
      const remotePath = `/${posix}`;
      byRemote.set(remotePath, { localPath: join(base, rel), treePath: `conf/${scope}/${posix}`, remotePath });
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
    else if (path.startsWith("/etc/mysql/") || path === "/etc/my.cnf" || path.startsWith("/etc/postgresql/"))
      hits.set(cfg.middleware.dbService, "restart");
  }
  return [...hits].map(([service, action]) => ({ service, action }));
};

const sha256 = (data: Buffer): string => createHash("sha256").update(data).digest("hex");

/** サーバ上のファイルの sha256。無いファイルは結果に含まれない（1 回の ssh で全部取る）。 */
const remoteHashes = async (
  host: string,
  paths: readonly string[],
  sudo: boolean,
): Promise<Map<string, string>> => {
  const hashes = new Map<string, string>();
  if (paths.length === 0) return hashes;
  const cmd = `${sudo ? "sudo " : ""}sha256sum ${paths.map(shQuote).join(" ")} 2>/dev/null || true`;
  const result = await onHostShell(host, cmd, { allowFailure: true });
  for (const line of result.stdout.split("\n")) {
    const m = /^([0-9a-f]{64})\s+\*?(.+)$/.exec(line.trim());
    if (m) hashes.set(m[2]!, m[1]!);
  }
  return hashes;
};

/** conf/_all と conf/<host> を sha の中身で固めた tar（ツリーと同じ場所にキャッシュする）。 */
const confTarFor = async (cfg: Config, sha: string, tree: string, host: string): Promise<string> => {
  const tar = join(tree, `.isuh-conf-${host}.tar`);
  if (existsSync(tar)) return tar;
  const scopes = ["_all", host]
    .filter((scope) => existsSync(join(tree, "conf", scope)))
    .map((scope) => `conf/${scope}`);
  const partial = `${tar}.${process.pid}.partial`;
  await local("git", ["-C", cfg.app.repoLocal, "archive", "--format=tar", "-o", partial, sha, "--", ...scopes]);
  renameSync(partial, tar);
  return tar;
};

/**
 * conf/ のうち、サーバ上の中身と違うファイルだけを配り、反映に必要な操作をする。
 *
 * - 比べるのは中身の sha256。同じなら配らず、reload / restart もしない（MySQL の restart は重い）
 * - 配るときは conf/ の tar を 1 回送り、1 回の ssh で全部 install する（ファイルごとに ssh を張らない）
 * - 戻り値は配ったパス
 */
const pushConf = async (
  cfg: Config,
  sha: string,
  tree: string,
  host: string,
  files: readonly ConfFile[],
): Promise<string[]> => {
  const remote = await remoteHashes(
    host,
    files.map((f) => f.remotePath),
    true,
  );
  const changed = files.filter((f) => remote.get(f.remotePath) !== sha256(readFileSync(f.localPath)));
  if (changed.length === 0) {
    log.info(`${host}: 設定ファイル ${files.length} 件はすべて変更なし`);
    return [];
  }
  log.info(
    `${host}: 設定ファイル ${changed.length}/${files.length} 件を配布: ${changed.map((f) => f.remotePath).join(" ")}`,
  );

  const tar = await confTarFor(cfg, sha, tree, host);
  const staged = `/tmp/isuh-conf-${process.pid}-${host}.tar`;
  await copyTo(host, tar, staged);
  // install -D なら親ディレクトリごと作る。/etc 配下なので sudo が要る。
  // /home 配下（systemd unit が読む env ファイルなど）は作業ユーザの持ち物のままにする
  const installs = changed.map((file) => {
    const owner = file.remotePath.startsWith("/home/")
      ? `-o ${shQuote(cfg.hosts.user)} -g ${shQuote(cfg.hosts.user)} `
      : "";
    return `sudo install -D -m 0644 ${owner}"$d"/${shQuote(file.treePath)} ${shQuote(file.remotePath)}`;
  });
  const script = [
    "set -e",
    "d=$(mktemp -d)",
    `trap 'rm -rf "$d" ${shQuote(staged).replaceAll("'", `'\\''`)}' EXIT`,
    `tar -xf ${shQuote(staged)} -C "$d"`,
    ...installs,
  ].join("\n");
  await onHostShell(host, script);

  const paths = changed.map((f) => f.remotePath);
  if (paths.some((p) => p.startsWith("/etc/systemd/"))) {
    await onHost(host, ["sudo", "systemctl", "daemon-reload"]);
  }
  if (paths.some((p) => p === "/etc/sysctl.conf" || p.startsWith("/etc/sysctl.d/"))) {
    log.info(`${host}: sysctl --system`);
    // 一部のキーが書けない環境（コンテナ等）でも、配布そのものは済んでいるので止めない
    const applied = await onHost(host, ["sudo", "sysctl", "--system"], { allowFailure: true });
    if (applied.exitCode !== 0) log.warn(`${host}: sysctl --system が一部失敗しました\n${applied.stderr.trim()}`);
  }
  for (const { service, action } of reloadsFor(cfg, paths)) {
    log.info(`${host}: ${action} ${service}`);
    await onHost(host, ["sudo", "systemctl", action, service]);
  }
  return paths;
};

// ---------------------------------------------------------------------------
// 配布・再起動
// ---------------------------------------------------------------------------

const remoteBinaryPath = (cfg: Config): string =>
  `${cfg.app.dir.replace(/\/$/, "")}/${cfg.app.langDir}/${cfg.app.binary}`;

const pushBinary = async (cfg: Config, host: string, localBinary: string): Promise<void> => {
  const dest = remoteBinaryPath(cfg);
  const remote = await remoteHashes(host, [dest], false);
  if (remote.get(dest) === sha256(readFileSync(localBinary))) {
    log.info(`${host}: バイナリは変更なし（送らずに再起動だけする）`);
    return;
  }
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

/** そのホストで今動いている sha。未記録なら null。 */
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
  const appHosts = new Set(cfg.hosts.app);
  const startedAt = Date.now();

  // 見るのは配るものだけ。ハーネスと同じリポジトリなので、runs/history.tsv などは常に動いている
  const dirty = await local(
    "git",
    ["-C", cfg.app.repoLocal, "status", "--porcelain", "--", cfg.app.localDir, "conf"],
    { allowFailure: true },
  );
  if (dirty.stdout.trim() !== "") {
    const files = dirty.stdout.trim().split("\n").length;
    log.warn(
      `アプリか conf/ に未コミットの変更が ${files} 件あります。` +
        `デプロイされるのは ${sha.slice(0, 7)} の内容だけです。`,
    );
  }

  const tree = await treeFor(cfg, sha);
  // バイナリは hosts.app へ、conf/ はそのホスト宛てがあれば全台へ（DB ホストの my.cnf など）
  const hosts = [...new Set([...cfg.hosts.all, ...cfg.hosts.app])].filter(
    (h) => appHosts.has(h) || confFilesFor(tree, h).length > 0,
  );

  if (opts.dryRun) {
    log.info(`[dry-run] ref=${ref} sha=${sha.slice(0, 7)}`);
    log.info(`[dry-run] build: ${cfg.app.buildCommand} (in ${join(cfg.app.localDir, cfg.app.langDir)})`);
    for (const host of hosts) {
      if (appHosts.has(host)) {
        log.info(
          `[dry-run] ${host}: ${cfg.app.localDir}/ を ${cfg.app.dir} へ同期（上書きのみ・${binaryRelPath(cfg)} は除く。前回から変わっていれば）`,
        );
        log.info(`[dry-run] ${host}: ${remoteBinaryPath(cfg)} を差し替え（中身が変わっていれば）`);
      }
      const confFiles = confFilesFor(tree, host);
      // dry-run はサーバを見ないので、配る候補と、変わっていた場合の反映を出す
      for (const file of confFiles) log.info(`[dry-run] ${host}: ${file.remotePath} を配布（中身が変わっていれば）`);
      for (const { service, action } of reloadsFor(cfg, confFiles.map((f) => f.remotePath))) {
        log.info(`[dry-run] ${host}: ${action} ${service}（対応する設定ファイルが変わっていれば）`);
      }
      if (appHosts.has(host)) log.info(`[dry-run] ${host}: restart ${cfg.app.service}`);
    }
    return sha;
  }

  const binary = await buildBinary(cfg, sha, tree);
  const appTar = await appTarFor(cfg, sha, tree);
  const appTreeId = await appTreeIdFor(cfg, sha);

  // ホストごとに直列（配布 → 再起動）、ホスト間は並列。
  // 1 台でも失敗したら全体を失敗にする（後続のベンチを走らせないため）。
  const results = await Promise.allSettled(
    hosts.map(async (host) => {
      const confFiles = confFilesFor(tree, host);
      if (confFiles.length > 0) await pushConf(cfg, sha, tree, host, confFiles);
      if (appHosts.has(host)) {
        await syncAppFiles(cfg, host, appTar, appTreeId, sha);
        await pushBinary(cfg, host, binary);
        await restartApp(cfg, host);
      }
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
バイナリ以外の app.localDir 以下（sql/・スクリプト・テンプレート等）も app.dir へ上書きで
同期する。削除はしないので、リポジトリで消したファイルはサーバに残る。
ref は app.repoLocal（通常はこのリポジトリ）で解決する。既定は HEAD。

  <git-ref>   デプロイする ref
  --dry-run   何をするかを表示して実行しない

デプロイされるのは作業ツリーではなく ref が指す commit の内容（git archive）。
未コミットの変更があれば警告するが、配られはしない。

各ホストの $HOME/.deployed_sha に「sha / ref / 時刻 / 実行者」を記録する。
同じ sha の conf/_all/** と conf/<host>/** のうち、サーバ上の中身と違うものだけを
サーバ上の絶対パスへ配布し、nginx / MySQL の設定が変わった場合は該当サービスを
reload / restart する。

ベンチと同じロック（hosts.primary の bench.lockPath）を取ってから配る。
走っているベンチの最中にアプリを差し替えて、計測を壊さないため。
ベンチが走っていれば終わるまで待つ（bench.lockTimeoutSec まで）。--dry-run は取らない。

stdout にはデプロイした sha だけを出す（SHA=$(isuh deploy) が使える）。

終了コード: 0 成功 / 10 ロックを取れない / 20 デプロイ失敗 / 64 引数不正 / 65 設定不備`,

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
    if (dryRun) {
      out(await deploy(cfg, { ref, dryRun }));
      return;
    }
    // bench は自分でロックを握ってから deploy() を呼ぶ。単体の deploy はここで同じロックを取る
    const lock = await acquireLock(`deploy ${ref}`, cfg);
    try {
      out(await deploy(cfg, { ref, dryRun }));
    } finally {
      await lock.release();
    }
  },
};
