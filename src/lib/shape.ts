/**
 * レスポンスの「構造」— check(#12)
 *
 * レギュレーションはレスポンスの構造（JSON オブジェクトの形）の変更を禁じている。
 * 違反は失格なので、値ではなく **キーと型の形** だけを取り出して比べる。
 *
 * 値を見ないのは、ベンチや initialize のたびに中身が変わるため。
 * 形だけなら「initialize 直後に 1 回記録しておけば、以後ずっと比較できる」。
 */

export type Shape =
  | { t: "null" }
  | { t: "string" }
  | { t: "number" }
  | { t: "boolean" }
  /** of は要素の形の集合（配列の中身が一様でないこともある） */
  | { t: "array"; of: Shape[] }
  | { t: "object"; fields: Record<string, Shape> };

const canonical = (shape: Shape): string => {
  switch (shape.t) {
    case "array":
      return `[${shape.of.map(canonical).sort().join("|")}]`;
    case "object":
      return `{${Object.keys(shape.fields)
        .sort()
        .map((k) => `${k}:${canonical(shape.fields[k]!)}`)
        .join(",")}}`;
    default:
      return shape.t;
  }
};

const uniq = (shapes: Shape[]): Shape[] => {
  const seen = new Map<string, Shape>();
  for (const shape of shapes) seen.set(canonical(shape), shape);
  return [...seen.values()];
};

export const shapeOf = (value: unknown): Shape => {
  if (value === null || value === undefined) return { t: "null" };
  if (Array.isArray(value)) return { t: "array", of: uniq(value.map(shapeOf)) };
  if (typeof value === "object") {
    const fields: Record<string, Shape> = {};
    for (const [key, item] of Object.entries(value)) fields[key] = shapeOf(item);
    return { t: "object", fields };
  }
  if (typeof value === "string") return { t: "string" };
  if (typeof value === "number") return { t: "number" };
  if (typeof value === "boolean") return { t: "boolean" };
  return { t: "null" };
};

/** 片方が null なら互換とみなす（null を取りうるフィールドを毎回の差分にしないため）。 */
const compatible = (expected: Shape, actual: Shape): boolean =>
  diff(expected, actual, "$").length === 0;

const describe = (shape: Shape): string => canonical(shape);

const diff = (expected: Shape, actual: Shape, path: string): string[] => {
  // null を取りうるフィールドは、記録時と実行時で型が揺れる。これは違反ではない。
  if (expected.t === "null" || actual.t === "null") return [];

  if (expected.t !== actual.t) {
    return [`${path}: 型が変わった ${expected.t} → ${actual.t}`];
  }

  if (expected.t === "array" && actual.t === "array") {
    // 記録時か実行時のどちらかが空配列なら、中身は比べようがない
    if (expected.of.length === 0 || actual.of.length === 0) return [];

    // 記録時の要素が 1 種類なら、そこへ潜って「どのキーが変わったか」まで言う。
    // 「要素の形が変わった」だけでは直しようがないため。
    if (expected.of.length === 1) {
      return actual.of.flatMap((item) => diff(expected.of[0]!, item, `${path}[]`));
    }

    const problems: string[] = [];
    for (const item of actual.of) {
      if (!expected.of.some((candidate) => compatible(candidate, item))) {
        problems.push(
          `${path}[]: 要素の形が変わった ${expected.of.map(describe).join("|")} → ${describe(item)}`,
        );
      }
    }
    return problems;
  }

  if (expected.t === "object" && actual.t === "object") {
    const problems: string[] = [];
    const expectedKeys = Object.keys(expected.fields).sort();
    const actualKeys = new Set(Object.keys(actual.fields));

    const removed = expectedKeys.filter((k) => !actualKeys.has(k));
    if (removed.length > 0) problems.push(`${path}: キーが消えた ${removed.join(", ")}`);

    const added = [...actualKeys].filter((k) => !(k in expected.fields)).sort();
    if (added.length > 0) problems.push(`${path}: キーが増えた ${added.join(", ")}`);

    for (const key of expectedKeys) {
      const actualField = actual.fields[key];
      if (actualField === undefined) continue;
      problems.push(...diff(expected.fields[key]!, actualField, `${path}.${key}`));
    }
    return problems;
  }

  return [];
};

/**
 * 記録した形と今の形を比べる。
 * 返るのは人が読める差分の一覧。空なら構造は変わっていない。
 */
export const diffShapes = (expected: Shape, actual: Shape): string[] => diff(expected, actual, "$");
