import { mkdirSync } from "node:fs";
import { homedir } from "node:os";
import { join } from "node:path";
import { execa } from "execa";
import { log } from "./log.js";

export interface RunResult {
  stdout: string;
  stderr: string;
  exitCode: number;
}

/** シェルに渡す値をひとつの単語として囲む。 */
export const shQuote = (value: string): string => `'${value.replaceAll("'", `'\\''`)}'`;

/**
 * ssh / scp に共通で渡すオプション。
 *
 * ISUH_SSH_CONFIG を設定すると `-F <path>` を付ける。
 * 手元の ~/.ssh/config を汚さずにホスト別名を定義したいとき（test/ のコンテナ環境や、
 * 当日リポジトリに同梱する ssh_config）に使う。未設定なら通常どおり ~/.ssh/config が効く。
 */
export const sshOpts = (): string[] => {
  const opts = ["-o", "ConnectTimeout=5", ...muxOpts()];
  const configFile = process.env["ISUH_SSH_CONFIG"];
  if (configFile) opts.push("-F", configFile);
  return opts;
};

/**
 * ssh の接続を使い回す（ControlMaster）。
 *
 * 手元から競技サーバへの新規接続は 1 回約 1 秒かかり、使い回すと約 0.04 秒になる（ISUCON14 の
 * リハーサルで実測）。deploy・bench・計測の回収は ssh を数十回張るので、ここが効く。
 * - ControlPersist は deploy → ベンチ → 回収の 1 周（数分）より長くし、次の周でも使い回す
 * - ServerAliveInterval: 相手が再起動などで消えたとき、残った接続を早めに捨てる
 * - ISUH_SSH_MUX=0 で無効にできる（接続まわりを切り分けたいとき）
 */
let muxDirReady = false;
const muxOpts = (): string[] => {
  if (process.env["ISUH_SSH_MUX"] === "0") return [];
  const dir = join(homedir(), ".ssh");
  if (!muxDirReady) {
    mkdirSync(dir, { recursive: true, mode: 0o700 });
    muxDirReady = true;
  }
  return [
    "-o", "ControlMaster=auto",
    // %C は接続先・ユーザ・ポートのハッシュ。UNIX ソケットのパス長の上限（104 字前後）に収まる
    "-o", `ControlPath=${join(dir, "isuh-%C")}`,
    "-o", "ControlPersist=300",
    "-o", "ServerAliveInterval=15",
    "-o", "ServerAliveCountMax=3",
  ];
};

const runSsh = async (
  host: string,
  remoteCommand: string,
  opts: { allowFailure?: boolean },
): Promise<RunResult> => {
  log.debug(`ssh ${host} -- ${remoteCommand}`);
  const result = await execa("ssh", [...sshOpts(), host, remoteCommand], {
    reject: !opts.allowFailure,
    all: false,
  });
  return {
    stdout: result.stdout ?? "",
    stderr: result.stderr ?? "",
    exitCode: result.exitCode ?? 0,
  };
};

/**
 * 1 ホストでコマンドを実行する。args は argv として扱う。
 *
 * ssh は残りの引数を空白で連結した「1 本の文字列」をリモートのログインシェルに渡す。
 * つまり配列で渡してもリモート側では必ず再パースされる。事故を防ぐため、
 * ここで各要素をクォートしてから連結する（`onHost(h, ["echo", "a b"])` は
 * リモートで `echo 'a b'` になる）。
 *
 * リダイレクトや `$HOME` の展開が要るときは onHostShell を使う。
 */
export const onHost = async (
  host: string,
  args: readonly string[],
  opts: { allowFailure?: boolean } = {},
): Promise<RunResult> => runSsh(host, args.map(shQuote).join(" "), opts);

/**
 * 1 ホストでシェルスクリプトを実行する。
 * パイプ・リダイレクト・`$HOME` の展開が必要なときだけ使う。
 * 設定由来の値を埋めるときは shQuote で囲むこと。
 */
export const onHostShell = async (
  host: string,
  script: string,
  opts: { allowFailure?: boolean } = {},
): Promise<RunResult> => runSsh(host, script, opts);

/** 複数ホストで並列に実行する。1 台でも失敗すれば reject する。 */
export const onHosts = async (
  hosts: readonly string[],
  args: readonly string[],
): Promise<Map<string, RunResult>> => {
  const results = await Promise.all(hosts.map(async (h) => [h, await onHost(h, args)] as const));
  return new Map(results);
};

/** ローカルのファイルをホストへ送る。 */
export const copyTo = async (host: string, localPath: string, remotePath: string): Promise<void> => {
  log.debug(`scp ${localPath} ${host}:${remotePath}`);
  await execa("scp", [...sshOpts(), "-q", localPath, `${host}:${remotePath}`]);
};

/** ホスト上のファイルを手元へ持ってくる。取れなければ false（allowFailure 時）。 */
export const copyFrom = async (
  host: string,
  remotePath: string,
  localPath: string,
  opts: { allowFailure?: boolean } = {},
): Promise<boolean> => {
  log.debug(`scp ${host}:${remotePath} ${localPath}`);
  try {
    await execa("scp", [...sshOpts(), "-q", `${host}:${remotePath}`, localPath]);
    return true;
  } catch (error) {
    if (opts.allowFailure) return false;
    throw error;
  }
};

/** ローカルでコマンドを実行する。こちらは本当にシェルを経由しない。 */
export const local = async (
  file: string,
  args: readonly string[],
  opts: { cwd?: string; allowFailure?: boolean } = {},
): Promise<RunResult> => {
  log.debug(`local -- ${file} ${args.join(" ")}`);
  const result = await execa(file, args, {
    cwd: opts.cwd,
    reject: !opts.allowFailure,
  });
  return {
    stdout: result.stdout ?? "",
    stderr: result.stderr ?? "",
    exitCode: result.exitCode ?? 0,
  };
};
