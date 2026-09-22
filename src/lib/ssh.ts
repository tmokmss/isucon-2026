import { execa } from "execa";
import { log } from "./log.js";

export interface RunResult {
  stdout: string;
  stderr: string;
  exitCode: number;
}

/**
 * 1 ホストでコマンドを実行する。
 * args は配列で渡す（シェルを経由しないので引用符の事故が起きない）。
 */
export const onHost = async (
  host: string,
  args: readonly string[],
  opts: { allowFailure?: boolean } = {},
): Promise<RunResult> => {
  log.debug(`ssh ${host} -- ${args.join(" ")}`);
  const result = await execa("ssh", ["-o", "ConnectTimeout=5", host, ...args], {
    reject: !opts.allowFailure,
    all: false,
  });
  return {
    stdout: result.stdout ?? "",
    stderr: result.stderr ?? "",
    exitCode: result.exitCode ?? 0,
  };
};

/** 複数ホストで並列に実行する。1 台でも失敗すれば reject する。 */
export const onHosts = async (
  hosts: readonly string[],
  args: readonly string[],
): Promise<Map<string, RunResult>> => {
  const results = await Promise.all(hosts.map(async (h) => [h, await onHost(h, args)] as const));
  return new Map(results);
};

/** ローカルでコマンドを実行する。 */
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
