/**
 * top の要約（report.ts の summarizeTop）の細かい挙動の確認。metrics.sh から呼ぶ。
 */
import { summarizeTop } from "../src/lib/report.js";

const frame = (rows: string[]): string =>
  [
    "top - 10:00:00 up 1 day,  1 user,  load average: 1.00, 1.00, 1.00",
    "Tasks: 100 total,   1 running",
    "%Cpu(s): 50.0 us,  0.0 sy",
    "",
    "    PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND",
    ...rows,
    "",
  ].join("\n");

const text = [
  // 1 フレーム目は起動からの累積に近いので捨てられるはず
  frame(["      1 root      20   0  100000  10000   1000 S 999.0   0.1   0:01.00 bogus"]),
  frame([
    "    100 isucon    20   0  100000  10000   1000 S 180.0   1.0   0:10.00 isuapp",
    "    200 mysql     20   0  100000  10000   1000 S  80.0   5.0   0:10.00 mysqld",
    "    300 www-data  20   0  100000  10000   1000 S  10.0   0.1   0:01.00 nginx",
    "    301 www-data  20   0  100000  10000   1000 S  10.0   0.1   0:01.00 nginx",
    "    400 isucon    20   0  100000  10000   1000 R   5.0   0.1   0:01.00 top",
  ]),
  frame([
    "    100 isucon    20   0  100000  10000   1000 S 220.0   1.0   0:10.00 isuapp",
    "    200 mysql     20   0  100000  10000   1000 S 120.0   5.0   0:10.00 mysqld",
    "    300 www-data  20   0  100000  10000   1000 S   0.0   0.1   0:01.00 nginx",
  ]),
].join("\n");

const fail = (message: string): never => {
  process.stderr.write(`  NG  ${message}\n`);
  process.exit(1);
};

const result = summarizeTop(text);
if (result === undefined) fail("要約できない");
const got = result!.map((p) => `${p.command}=${p.cpu}`).join(",");

// isuapp: (180+220)/2=200, mysqld: (80+120)/2=100, nginx: (10+10+0)/2=10
if (got !== "isuapp=200,mysqld=100,nginx=10") fail(`集計が違う: ${got}`);
if (got.includes("bogus")) fail("1 フレーム目を捨てていない");
if (got.includes("top")) fail("計測プロセス自身を除いていない");

if (summarizeTop(frame(["    100 isucon 20 0 1 1 1 S 50.0 1.0 0:01.00 isuapp"])) !== undefined) {
  fail("1 フレームしか無いときは要約できないはず");
}

process.stdout.write("5 件\n");
