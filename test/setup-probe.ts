/**
 * setup の「手元のファイルを書き換える」部分の検証。
 * ~/.ssh は本物を触れないので、一時ディレクトリに対して writeLocalSshConfig を呼ぶ。
 * 通ったケース数を stdout に出し、失敗したら非ゼロで終わる。
 */
import { mkdtempSync, readFileSync, rmSync, statSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { upsertBlock, withInclude, writeLocalSshConfig } from "../src/commands/setup.js";
import { config } from "./config.test.js";

let passed = 0;
const check = (name: string, cond: boolean, detail = ""): void => {
  if (!cond) {
    process.stderr.write(`NG: ${name}\n${detail}\n`);
    process.exit(1);
  }
  passed++;
};

// --- 管理ブロック ---------------------------------------------------------
const base = "127.0.0.1 localhost\n";
const once = upsertBlock(base, "10.0.0.1\tisu1");
const twice = upsertBlock(once, "10.0.0.1\tisu1");
check("ブロックが末尾に足される", once === `${base}# >>> isuh >>>\n10.0.0.1\tisu1\n# <<< isuh <<<\n`, once);
check("同じ内容で 2 回書いても変わらない", twice === once, twice);
const changed = upsertBlock(`${once}# after\n`, "10.0.0.2\tisu1");
check(
  "中身だけ差し替わり、ブロック外の行は残る",
  changed.includes("# after") && changed.includes("10.0.0.2") && !changed.includes("10.0.0.1"),
  changed,
);
check("空のファイルにも書ける", upsertBlock("", "x") === "# >>> isuh >>>\nx\n# <<< isuh <<<\n");

// --- Include -------------------------------------------------------------
const existing = "Host github.com\n  User git\n";
const inc = withInclude(existing);
check(
  "Include は最初の Host より前に入る",
  inc !== null && inc.indexOf("Include isuh_config") < inc.indexOf("Host github.com"),
  String(inc),
);
check("Include 済みなら何もしない", withInclude(inc ?? "") === null);

// --- ~/.ssh の書き換え ----------------------------------------------------
const saved = process.env["ISUH_SSH_CONFIG"];
delete process.env["ISUH_SSH_CONFIG"];
const dir = mkdtempSync(join(tmpdir(), "isuh-setup-probe-"));
try {
  writeFileSync(join(dir, "config"), existing);
  writeLocalSshConfig(config, dir);
  writeLocalSshConfig(config, dir);
  const main = readFileSync(join(dir, "config"), "utf8");
  const generated = readFileSync(join(dir, "isuh_config"), "utf8");
  check("2 回流しても Include は 1 行", main.split("\n").filter((l) => l === "Include isuh_config").length === 1, main);
  check("元の ~/.ssh/config が退避される", readFileSync(join(dir, "config.isuh-bak"), "utf8") === existing);
  check("元の設定は残る", main.includes("Host github.com"));
  check("isuh_config に全台の定義", ["isu1", "isu2", "isu3"].every((h) => generated.includes(`Host ${h}\n`)), generated);
  check("isuh_config は 600", (statSync(join(dir, "isuh_config")).mode & 0o777) === 0o600);
} finally {
  rmSync(dir, { recursive: true, force: true });
  if (saved !== undefined) process.env["ISUH_SSH_CONFIG"] = saved;
}

process.stdout.write(`${passed} 件`);
