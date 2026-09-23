import { command as bench } from "./commands/bench.js";
import { command as check } from "./commands/check.js";
import { command as deploy } from "./commands/deploy.js";
import { command as finalize } from "./commands/finalize.js";
import { command as history } from "./commands/history.js";
import { command as monitor } from "./commands/monitor.js";
import { command as portal } from "./commands/portal.js";
import { command as setup } from "./commands/setup.js";
import { command as verifyRestart } from "./commands/verifyRestart.js";
import type { Command } from "./lib/command.js";
import { EXIT, HarnessError } from "./lib/exit.js";
import { log, out } from "./lib/log.js";

/** 表示順を保つため配列で持つ。 */
const COMMANDS: ReadonlyArray<readonly [string, Command]> = [
  ["setup", setup],
  ["deploy", deploy],
  ["check", check],
  ["bench", bench],
  ["history", history],
  ["portal", portal],
  ["monitor", monitor],
  ["verify-restart", verifyRestart],
  ["finalize", finalize],
];

const topUsage = (): string => {
  const width = Math.max(...COMMANDS.map(([name]) => name.length));
  const lines = COMMANDS.map(([name, cmd]) => `  ${name.padEnd(width)}  ${cmd.summary}`);
  return [
    "usage: isuh <command> [args]",
    "",
    ...lines,
    "",
    "各コマンドの詳細は isuh <command> --help",
    "契約は docs/DESIGN.md",
  ].join("\n");
};

const main = async (): Promise<void> => {
  const [name, ...rest] = process.argv.slice(2);

  if (name === undefined || name === "help" || name === "-h" || name === "--help") {
    out(topUsage());
    return;
  }

  const entry = COMMANDS.find(([n]) => n === name);
  if (!entry) {
    log.error(`不明なコマンド: ${name}`);
    out(topUsage());
    process.exitCode = EXIT.USAGE;
    return;
  }

  const [, cmd] = entry;
  if (rest.includes("-h") || rest.includes("--help")) {
    out(cmd.usage);
    return;
  }

  await cmd.run(rest);
};

main().catch((err: unknown) => {
  if (err instanceof HarnessError) {
    log.error(err.message);
    process.exit(err.code);
  }
  log.error(err instanceof Error ? (err.stack ?? err.message) : String(err));
  process.exit(1);
});
