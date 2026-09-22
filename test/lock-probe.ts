/**
 * ロックの検証用。acquireLock で取って、リモートのトレースファイルに
 * ENTER / EXIT を追記してから解放する。
 *
 *   tsx test/lock-probe.ts <id> <保持ミリ秒>
 *
 * 保持ミリ秒が負なら ENTER を書いたまま永久に持ち続ける（kill の検証用）。
 */
import { loadConfig } from "../src/config.js";
import { acquireLock } from "../src/lib/lock.js";
import { onHostShell, shQuote } from "../src/lib/ssh.js";

const TRACE = "/tmp/isuh-lock-trace.txt";

const id = process.argv[2] ?? "x";
const holdMs = Number(process.argv[3] ?? "1000");

const cfg = await loadConfig();
const append = (text: string): Promise<unknown> =>
  onHostShell(cfg.hosts.primary, `printf '%s\\n' ${shQuote(text)} >> ${shQuote(TRACE)}`);

const lock = await acquireLock(`probe ${id}`, cfg);
try {
  await append(`ENTER ${id}`);
  if (holdMs < 0) {
    // kill されるまで持ち続ける
    await new Promise(() => {});
  }
  await new Promise((resolve) => setTimeout(resolve, holdMs));
  await append(`EXIT ${id}`);
} finally {
  await lock.release();
}
