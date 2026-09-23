/**
 * 計測ツールの出力を report.md 向けに畳む — docs/DESIGN.md §4
 *
 * report.md はそのままエージェントに渡る。alp の既定の表は 20 列近くあり罫線も多く、
 * pt-query-digest は 1 クエリにつき 40 行ほど出る。どちらも人間向けの装飾が大半で、
 * 読む側の予算（トークン）を食う。
 *
 * そこで collect はツールに機械向けの形式（alp は TSV、pt-query-digest は JSON）で出させて
 * そのまま alp.txt / ptquery.txt に残し、ここで「次に何を直すか」の判断に要る列だけの
 * 表にする。生のファイルは深掘りしたいときに読む。
 *
 * 形式が想定と違う（未取得の理由が書いてある・ツールの版が違う）ときは null を返し、
 * 呼び出し側が元のテキストをそのまま載せる。
 */

// ---------------------------------------------------------------------------
// alp
// ---------------------------------------------------------------------------

/** 表に出す列。alp の TSV のヘッダ名 → report.md での名前 */
const ALP_COLUMNS: ReadonlyArray<readonly [string, string]> = [
  ["Count", "count"],
  ["4xx", "4xx"],
  ["5xx", "5xx"],
  ["Method", "method"],
  ["Uri", "uri"],
  ["Sum", "sum_s"],
  ["Avg", "avg_s"],
  ["P99", "p99_s"],
  ["Max", "max_s"],
  ["Avg(Body)", "avg_bytes"],
];

/** alp の TSV（1 ホスト分）を畳む。 */
const summarizeAlpTable = (tsv: string, limit: number): string | null => {
  const lines = tsv.split("\n").filter((l) => l.trim() !== "");
  const header = lines[0]?.split("\t");
  if (header === undefined || !header.includes("Uri") || !header.includes("Sum")) return null;

  const index = (name: string): number => header.indexOf(name);
  const picked = ALP_COLUMNS.filter(([name]) => index(name) >= 0);
  const rows = lines.slice(1).map((l) => l.split("\t"));
  if (rows.length === 0) return "（リクエストなし）";

  const sumAt = index("Sum");
  const total = rows.reduce((acc, r) => acc + Number(r[sumAt] ?? 0), 0);

  const out = [["share", ...picked.map(([, label]) => label)].join("\t")];
  // alp は並べ替え済み（collect が --sort と -r を渡す）。割合は全行の合計に対して出し、表示は上位だけ
  for (const r of rows.slice(0, limit)) {
    const share = total > 0 ? `${((Number(r[sumAt] ?? 0) / total) * 100).toFixed(1)}%` : "—";
    const cells = picked.map(([name]) => {
      const v = r[index(name)] ?? "";
      // 本文サイズは小数が付いて出る（615.000）。バイト数なので整数にする
      return name === "Avg(Body)" ? String(Math.round(Number(v))) : v;
    });
    out.push([share, ...cells].join("\t"));
  }
  if (rows.length > limit) out.push(`…（上位 ${limit} / 全 ${rows.length} 種類）`);
  return out.join("\n");
};

/**
 * alp.txt を畳む。nginx が複数台なら `== host ==` の節ごとに畳む。
 * 未取得の節（理由のテキスト）はそのまま残す。
 */
export const summarizeAlp = (text: string, limit = Number.POSITIVE_INFINITY): string | null => {
  const sections = splitHostSections(text);
  const parts = sections.map(({ host, body }) => {
    const table = summarizeAlpTable(body, limit);
    const content = table ?? body.trim();
    return host === null ? content : `== ${host} ==\n${content}`;
  });
  if (sections.every(({ body }) => summarizeAlpTable(body, limit) === null)) return null;
  return parts.join("\n\n");
};

const splitHostSections = (text: string): { host: string | null; body: string }[] => {
  const re = /^== (.+) ==$/gm;
  const heads = [...text.matchAll(re)];
  if (heads.length === 0) return [{ host: null, body: text }];
  return heads.map((m, i) => {
    const start = (m.index ?? 0) + m[0].length;
    const end = i + 1 < heads.length ? (heads[i + 1]!.index ?? text.length) : text.length;
    return { host: m[1]!, body: text.slice(start, end) };
  });
};

// ---------------------------------------------------------------------------
// pt-query-digest
// ---------------------------------------------------------------------------

interface PtMetric {
  sum?: string;
  avg?: string;
  pct_95?: string;
  max?: string;
}

interface PtClass {
  fingerprint?: string;
  query_count?: number;
  example?: { query?: string };
  metrics?: Record<string, PtMetric | undefined>;
}

interface PtDigest {
  global?: { query_count?: number; unique_query_count?: number; metrics?: Record<string, PtMetric | undefined> };
  classes?: PtClass[];
}

const FINGERPRINT_MAX = 160;
const EXAMPLE_MAX = 400;
/** 実例（EXPLAIN にそのまま使える形）を載せる上位の件数 */
const EXAMPLES = 3;

const shorten = (s: string, max: number): string => {
  const one = s.replace(/\s+/g, " ").trim();
  return one.length <= max ? one : `${one.slice(0, max)}…`;
};

const ms = (sec: string | undefined): string => (sec === undefined ? "—" : (Number(sec) * 1000).toFixed(2));

/** pt-query-digest --output json を畳む。 */
export const summarizePtQuery = (text: string): string | null => {
  let digest: PtDigest;
  try {
    digest = JSON.parse(text) as PtDigest;
  } catch {
    return null;
  }
  const classes = digest.classes;
  if (!Array.isArray(classes)) return null;
  if (classes.length === 0) return "（クエリなし。slow query log が空か、long_query_time が 0 になっていない）";

  const totalTime = Number(digest.global?.metrics?.["Query_time"]?.sum ?? 0);
  const head =
    `全 ${digest.global?.query_count ?? "?"} クエリ / ${digest.global?.unique_query_count ?? "?"} 種類 / ` +
    `合計 ${totalTime.toFixed(3)}s（上位 ${classes.length} 種類を表示。share は全体の合計時間に対する割合）`;

  const rows = [["rank", "share", "calls", "total_s", "avg_ms", "p95_ms", "rows_examined_avg", "rows_sent_avg", "query"].join("\t")];
  classes.forEach((c, i) => {
    const qt = c.metrics?.["Query_time"];
    const sum = Number(qt?.sum ?? 0);
    rows.push(
      [
        String(i + 1),
        totalTime > 0 ? `${((sum / totalTime) * 100).toFixed(1)}%` : "—",
        String(c.query_count ?? "?"),
        sum.toFixed(3),
        ms(qt?.avg),
        ms(qt?.pct_95),
        c.metrics?.["Rows_examined"]?.avg ?? "—",
        c.metrics?.["Rows_sent"]?.avg ?? "—",
        shorten(c.fingerprint ?? "", FINGERPRINT_MAX),
      ].join("\t"),
    );
  });

  const examples = classes
    .slice(0, EXAMPLES)
    .map((c, i) => (c.example?.query ? `-- rank ${i + 1}\n${shorten(c.example.query, EXAMPLE_MAX)};` : null))
    .filter((e): e is string => e !== null);

  return [head, "", rows.join("\n"), ...(examples.length > 0 ? ["", "-- 上位の実例（EXPLAIN 用）", examples.join("\n")] : [])].join(
    "\n",
  );
};
