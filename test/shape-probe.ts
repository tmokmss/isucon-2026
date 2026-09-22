/**
 * 構造比較(#12)の細かい挙動の確認。check.sh から呼ぶ。
 * サーバを立てずに確かめられる部分だけをここで見る。
 */
import { diffShapes, shapeOf } from "../src/lib/shape.js";

const cases: { name: string; recorded: unknown; now: unknown; expect: "same" | "diff" }[] = [
  {
    name: "null だったフィールドに値が入っても違反ではない",
    recorded: { id: 1, name: "isucon", email: null },
    now: { id: 1, name: "isucon", email: "a@example.com" },
    expect: "same",
  },
  {
    name: "値が入っていたフィールドが null になっても違反ではない",
    recorded: { id: 1, email: "a@example.com" },
    now: { id: 1, email: null },
    expect: "same",
  },
  {
    name: "型が変わったら違反",
    recorded: { id: 1 },
    now: { id: "1" },
    expect: "diff",
  },
  {
    name: "キーが消えたら違反",
    recorded: { id: 1, name: "a" },
    now: { id: 1 },
    expect: "diff",
  },
  {
    name: "キーが増えても違反",
    recorded: { id: 1 },
    now: { id: 1, extra: true },
    expect: "diff",
  },
  {
    name: "配列の要素の形が変わったら違反",
    recorded: { items: [{ id: 1, name: "a" }] },
    now: { items: [{ id: 1, title: "a" }] },
    expect: "diff",
  },
  {
    name: "要素数が変わっただけなら違反ではない",
    recorded: { items: [{ id: 1 }] },
    now: { items: [{ id: 1 }, { id: 2 }, { id: 3 }] },
    expect: "same",
  },
  {
    name: "記録時が空配列なら中身は比べない",
    recorded: { items: [] },
    now: { items: [{ id: 1 }] },
    expect: "same",
  },
  {
    name: "実行時が空配列でも中身は比べない",
    recorded: { items: [{ id: 1 }] },
    now: { items: [] },
    expect: "same",
  },
  {
    name: "キーの順序が違うだけなら違反ではない",
    recorded: { a: 1, b: "x" },
    now: { b: "x", a: 1 },
    expect: "same",
  },
  {
    name: "入れ子の奥でも見つける",
    recorded: { a: { b: { c: [{ d: 1 }] } } },
    now: { a: { b: { c: [{ d: "1" }] } } },
    expect: "diff",
  },
];

let failed = 0;
for (const testCase of cases) {
  const diffs = diffShapes(shapeOf(testCase.recorded), shapeOf(testCase.now));
  const actual = diffs.length === 0 ? "same" : "diff";
  if (actual !== testCase.expect) {
    failed += 1;
    process.stderr.write(
      `  NG  ${testCase.name}: ${testCase.expect} のはずが ${actual} (${diffs.join(" / ")})\n`,
    );
  }
}

if (failed > 0) process.exit(1);
process.stdout.write(`${cases.length} 件\n`);
