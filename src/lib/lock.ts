import { hostname, userInfo } from "node:os";
import { setTimeout as delay } from "node:timers/promises";
import { execa } from "execa";
import type { Config } from "../config.js";
import { EXIT, HarnessError } from "./exit.js";
import { log } from "./log.js";
import { shQuote, sshOpts } from "./ssh.js";

/**
 * ベンチのロック — docs/DESIGN.md §7
 *
 * デプロイだけをロックしても不十分。ベンチ実行中に別のデプロイが走れば計測は壊れる。
 * bench の全ステップがひとつのクリティカルセクション。
 *
 * 保持方法:
 *   hosts.primary 上の flock を、ssh の子プロセスを生かしたまま保持する。
 *   末尾の `cat` が stdin を開いたまま保つため、ssh プロセスが生きている間だけ
 *   ロックが保持される。子プロセスが死ねば接続が切れ、リモートのシェルも死んで
 *   fd 9 が閉じ、flock が自動解放される。
 *
 *   ローカルのファイルロックではなくサーバ側に置く理由は、worktree・別ターミナル・
 *   別マシンをまたいで効かせるため。
 */
export interface Lock {
  /** 多重呼び出し安全。 */
  release(): Promise<void>;
}

/**
 * ロックを取りに行くリモートスクリプト。
 *
 * 取れなかったときに BUSY と保持者を出してから exit 10 する。保持者の記録は
 * ロックを取った「後」に書くので、待っている側が読むのは必ず現在の保持者。
 */
const remoteScript = (lockPath: string, timeoutSec: number, ownerLine: string): string => {
  const lock = shQuote(lockPath);
  const owner = shQuote(`${lockPath}.owner`);
  return [
    `exec 9>${lock}`,
    `if ! flock -w ${Math.max(0, Math.floor(timeoutSec))} 9; then`,
    `  echo BUSY`,
    `  cat ${owner} 2>/dev/null`,
    `  exit 10`,
    `fi`,
    `printf '%s\\n' ${shQuote(ownerLine)} > ${owner}`,
    `echo ACQUIRED`,
    `cat`,
  ].join("\n");
};

/**
 * ロックを取る。取れなければ EXIT.LOCK で落ちる。
 *
 * owner には「何のために取ったか」を入れる（例: run-id や git ref）。
 * 待たされた側の stderr にそのまま出るので、誰の何を待っているかが分かる。
 */
export const acquireLock = async (owner: string, cfg: Config): Promise<Lock> => {
  const host = cfg.hosts.primary;
  const { lockPath, lockTimeoutSec } = cfg.bench;
  const ownerLine = [
    `${userInfo().username}@${hostname()}`,
    `started=${new Date().toISOString()}`,
    `pid=${process.pid}`,
    owner,
  ].join("  ");

  log.info(`ロックを取得します: ${host}:${lockPath}`);
  const child = execa("ssh", [...sshOpts(), host, remoteScript(lockPath, lockTimeoutSec, ownerLine)], {
    stdin: "pipe",
    // execa に溜め込ませず自分でストリームを読む。ACQUIRED が来た時点で先へ進みたいため。
    buffer: false,
    reject: false,
  });

  let stdout = "";
  let stderr = "";
  let onAcquired!: () => void;
  const acquiredSignal = new Promise<boolean>((resolve) => {
    onAcquired = () => resolve(true);
  });
  child.stdout?.on("data", (chunk: Buffer) => {
    stdout += chunk.toString();
    if (stdout.includes("ACQUIRED")) onAcquired();
  });
  child.stderr?.on("data", (chunk: Buffer) => {
    stderr += chunk.toString();
  });

  // ACQUIRED が来るか、ssh が先に終わるか。取れていれば ssh は cat で生き続ける。
  const acquired = await Promise.race([acquiredSignal, child.then(() => false)]);

  if (!acquired) {
    const result = await child;
    if (result.exitCode === 10) {
      const holder = stdout.replace(/^\s*BUSY\s*/, "").trim();
      throw new HarnessError(
        EXIT.LOCK,
        `ロックを取得できませんでした（${lockTimeoutSec}s 待機）: ${host}:${lockPath}\n` +
          `  保持者: ${holder === "" ? "不明（.owner が未作成）" : holder}\n` +
          `  ベンチが走っています。終わるのを待つか、bench.lockTimeoutSec を伸ばしてください。`,
      );
    }
    const detail = [stdout, stderr].filter((s) => s.trim() !== "").join("\n").trim();
    throw new HarnessError(
      EXIT.LOCK,
      `ロックの取得に失敗しました: ${host}:${lockPath} (exit=${result.exitCode})\n${detail}`,
    );
  }

  // 異常終了しても ssh を道連れにする。ssh が死ねばリモートの fd が閉じてロックが解ける。
  const killChild = (): void => {
    try {
      child.kill("SIGKILL");
    } catch {
      /* 既に終了している */
    }
  };
  process.once("exit", killChild);

  log.info(`ロックを取得しました: ${host}:${lockPath}`);

  let released = false;
  return {
    async release() {
      if (released) return;
      released = true;
      process.off("exit", killChild);

      // stdin を閉じるとリモートの cat が EOF で終わる = 正常解放。
      child.stdin?.end();
      const exited = await Promise.race([
        child.then(
          () => true,
          () => true,
        ),
        delay(3000).then(() => false),
      ]);
      if (!exited) {
        log.warn("ロック保持プロセスが終了しないため kill します");
        child.kill("SIGKILL");
        await child.catch(() => undefined);
      }
      log.info("ロックを解放しました");
    },
  };
};
