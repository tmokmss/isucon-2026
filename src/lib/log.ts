/**
 * ログ — docs/DESIGN.md §6
 * stdout は機械可読な結果のみ。ログはすべて stderr へ出す。
 */
const timestamp = (): string => new Date().toTimeString().slice(0, 8);

const emit = (level: string, message: string): void => {
  process.stderr.write(`[${timestamp()}] ${level.padEnd(5)} ${message}\n`);
};

export const log = {
  debug: (message: string): void => {
    if (process.env["VERBOSE"] === "1") emit("DEBUG", message);
  },
  info: (message: string): void => emit("INFO", message),
  warn: (message: string): void => emit("WARN", message),
  error: (message: string): void => emit("ERROR", message),
};

/** 機械可読な結果を stdout に出す唯一の口。 */
export const out = (line: string): void => {
  process.stdout.write(`${line}\n`);
};
