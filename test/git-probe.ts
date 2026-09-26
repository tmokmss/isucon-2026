/**
 * src/lib/git.ts の sameContent の確認。このリポジトリの履歴そのものを使う。
 * 通ったケース数を stdout に出す。
 */
import { execa } from "execa";
import { sameContent } from "../src/lib/git.js";

let passed = 0;
const check = (name: string, cond: boolean, detail = ""): void => {
  if (!cond) {
    process.stderr.write(`NG: ${name}\n${detail}\n`);
    process.exit(1);
  }
  passed++;
};

const root = (await execa("git", ["rev-parse", "--show-toplevel"])).stdout.trim();
const rev = async (ref: string): Promise<string> =>
  (await execa("git", ["-C", root, "rev-parse", ref])).stdout.trim();

// 直前の commit で src/ が変わっていないパスと、変わったパスを 1 つずつ探す代わりに、
// 同じ commit 同士（必ず同じ）と、初回の commit と HEAD の src/（必ず違う）で比べる
const head = await rev("HEAD");
const first = (await execa("git", ["-C", root, "rev-list", "--max-parents=0", "HEAD"])).stdout.trim().split("\n")[0]!;

check("同じ commit 同士は同じ", (await sameContent(root, head, head, ["src"])) === true);
check("初回の commit と HEAD の src/ は違う", (await sameContent(root, first, head, ["src"])) === false);
check(
  "存在しないパスは差分なし = 同じ",
  (await sameContent(root, first, head, ["no-such-dir-for-probe"])) === true,
);
check(
  "存在しない commit は比べられない",
  (await sameContent(root, "0000000000000000000000000000000000000000", head, ["src"])) === undefined,
);

process.stdout.write(`${passed}\n`);
