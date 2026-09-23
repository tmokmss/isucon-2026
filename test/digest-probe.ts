/**
 * src/lib/digest.ts（alp / pt-query-digest の出力を畳む）の検証。
 * 入力は検証用コンテナで実際に出したものと同じ形。通ったケース数を stdout に出す。
 */
import { summarizeAlp, summarizePtQuery } from "../src/lib/digest.js";

let passed = 0;
const check = (name: string, cond: boolean, detail = ""): void => {
  if (!cond) {
    process.stderr.write(`NG: ${name}\n${detail}\n`);
    process.exit(1);
  }
  passed++;
};

// --- alp -----------------------------------------------------------------
const header =
  "Count\t1xx\t2xx\t3xx\t4xx\t5xx\tMethod\tUri\tMin\tMax\tSum\tAvg\tP90\tP95\tP99\tStddev\tMin(Body)\tMax(Body)\tSum(Body)\tAvg(Body)";
const alpTsv = [
  header,
  "300\t0\t290\t0\t0\t10\tGET\t/api/user/.+\t0.010\t0.900\t3.000\t0.010\t0.020\t0.030\t0.500\t0.050\t100.000\t900.000\t60000.000\t200.400",
  "10\t0\t10\t0\t0\t0\tPOST\t/api/login\t0.050\t0.200\t1.000\t0.100\t0.150\t0.180\t0.200\t0.030\t20.000\t20.000\t200.000\t20.000",
].join("\n");

const alp = summarizeAlp(alpTsv) ?? "";
const alpLines = alp.split("\n");
check(
  "alp: 列を絞る",
  alpLines[0] === "share\tcount\t4xx\t5xx\tmethod\turi\tsum_s\tavg_s\tp99_s\tmax_s\tavg_bytes",
  alp,
);
check("alp: 合計時間に対する割合", alpLines[1]?.startsWith("75.0%\t300\t0\t10\tGET\t/api/user/.+\t3.000") === true, alp);
check("alp: 本文サイズは整数", alpLines[1]?.endsWith("\t200") === true, alp);
check("alp: 罫線や不要な列（P90 / Stddev / 1xx）が無い", !/[|+]-|P90|Stddev|1xx/.test(alp), alp);

const multi = summarizeAlp(`== isu1 ==\n${alpTsv}\n\n== isu2 ==\n未取得: alp を実行できませんでした (isu2, exit=127)`) ?? "";
check("alp: ホストごとの節を保つ", multi.includes("== isu1 ==\nshare\tcount") && multi.includes("== isu2 ==\n未取得"), multi);
check("alp: 全部未取得なら null（元のテキストを載せる）", summarizeAlp("未取得: alp を実行できませんでした") === null);
check("alp: リクエストが無い", summarizeAlp(header) === "（リクエストなし）");
const top1 = (summarizeAlp(alpTsv, 1) ?? "").split("\n");
check(
  "alp: 上位だけ表示し、割合は全行に対して出す",
  top1.length === 3 && top1[1]?.startsWith("75.0%") === true && top1[2] === "…（上位 1 / 全 2 種類）",
  top1.join("\n"),
);

// --- pt-query-digest -----------------------------------------------------
const metric = (sum: string, avg: string, p95: string) => ({ sum, avg, pct_95: p95, max: p95 });
const pt = JSON.stringify({
  global: { query_count: 1300, unique_query_count: 7, metrics: { Query_time: metric("10.000000", "0.007", "0.02") } },
  classes: [
    {
      fingerprint: "select * from users where id = ?",
      query_count: 1000,
      example: { query: "SELECT * FROM users\n  WHERE id = 42" },
      metrics: {
        Query_time: metric("8.000000", "0.008000", "0.012000"),
        Rows_examined: metric("100000", "100", "100"),
        Rows_sent: metric("1000", "1", "1"),
      },
    },
    {
      fingerprint: `select ${"x, ".repeat(100)}y from t`,
      query_count: 300,
      example: { query: "SELECT 1" },
      metrics: {
        Query_time: metric("1.500000", "0.005000", "0.006000"),
        Rows_examined: metric("0", "0", "0"),
        Rows_sent: metric("300", "1", "1"),
      },
    },
  ],
});
const ptOut = summarizePtQuery(pt) ?? "";
const ptLines = ptOut.split("\n");
check("pt: 全体の件数と合計時間", ptLines[0]?.startsWith("全 1300 クエリ / 7 種類 / 合計 10.000s") === true, ptOut);
check(
  "pt: 1 行 1 クエリの表",
  ptOut.includes("rank\tshare\tcalls\ttotal_s\tavg_ms\tp95_ms\trows_examined_avg\trows_sent_avg\tquery\n") &&
    ptOut.includes("1\t80.0%\t1000\t8.000\t8.00\t12.00\t100\t1\tselect * from users where id = ?"),
  ptOut,
);
check("pt: 長いクエリは切り詰める", ptOut.includes("x, x, …") || /x, …/.test(ptOut), ptOut);
check("pt: 実例を 1 行にして載せる", ptOut.includes("-- rank 1\nSELECT * FROM users WHERE id = 42;"), ptOut);
check("pt: JSON でなければ null", summarizePtQuery("未取得: pt-query-digest を実行できませんでした") === null);
check("pt: クエリが無い", (summarizePtQuery(JSON.stringify({ global: {}, classes: [] })) ?? "").startsWith("（クエリなし"));

process.stdout.write(`${passed} 件`);
