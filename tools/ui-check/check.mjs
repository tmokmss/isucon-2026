// ISUReal のフロントエンド表示確認（閲覧とログインのみ。データは書き換えない）。
// 初回だけ: このディレクトリで npm install && npx playwright install chromium
// 使い方: node check.mjs [--region asia|europe|america | --all] [--out <dir>] [--base <url>] [--user <name> --key <secret>]
// 画面ごとに HTTP/コンソールエラー・/api/ の 4xx/5xx・img の読み込み・フィードの投稿数を確かめる。
// 初期ユーザーが未投稿だとフィードは空（「まずは、あなたの一枚」）が正で、その場合は WARN にして
// 詳細は Memories の過去投稿で確かめる。
import { chromium } from "playwright";
import { mkdirSync } from "node:fs";
import { join } from "node:path";
import { parseArgs } from "node:util";

const { values: args } = parseArgs({
  options: {
    region: { type: "string", default: "asia" },
    all: { type: "boolean", default: false },
    out: { type: "string" },
    base: { type: "string", default: "http://52.69.202.109:8080" },
    user: { type: "string", default: "yuma_kobayashi_186" },
    key: { type: "string", default: "0f6-BAnyBELD2MA-t-fekeDtQN4ynqD0mfgOF8JWwWs" },
  },
});
const COUNTRY = { asia: "JP", europe: "DE", america: "US" };
// 公開の動作確認用ユーザーは asia 所属。他の地域の画面は、その地域のユーザーを --user / --key で渡す
// （asia のユーザーで europe / america に入ると API が 403 を返すのは仕様）。

const regions = args.all ? Object.keys(COUNTRY) : [args.region];
for (const r of regions) {
  if (!COUNTRY[r]) {
    console.error(`不正な --region: ${r}（asia|europe|america）`);
    process.exit(2);
  }
}
const stamp = new Date().toISOString().replace(/[:.]/g, "-").slice(0, 19);

async function checkRegion(browser, region) {
  const outDir = args.out
    ? args.all ? join(args.out, region) : args.out
    : join("out", `${region}-${stamp}`);
  mkdirSync(outDir, { recursive: true });

  const ctx = await browser.newContext({
    extraHTTPHeaders: { "X-Country-Code": COUNTRY[region] },
    viewport: { width: 430, height: 900 },
  });
  const page = await ctx.newPage();
  page.setDefaultTimeout(15000);

  let bucket = [];
  const note = (m) => bucket.push(m);
  page.on("console", (m) => m.type() === "error" && note(`console: ${m.text()}`));
  page.on("pageerror", (e) => note(`pageerror: ${e.message}`));
  page.on("requestfailed", (r) => {
    const f = r.failure()?.errorText ?? "";
    if (!f.includes("ABORTED")) note(`requestfailed: ${r.url()} ${f}`);
  });
  page.on("response", (r) => {
    const u = new URL(r.url());
    if (r.status() >= 400 && (u.pathname.startsWith("/api/") || ["document", "script", "stylesheet", "image"].includes(r.request().resourceType()))) {
      note(`HTTP ${r.status()}: ${r.request().method()} ${u.pathname}`);
    }
  });

  let me = null;
  let forbidden = false;
  page.on("response", async (r) => {
    if (new URL(r.url()).pathname !== "/api/me") return;
    if (r.status() === 403) forbidden = true;
    else if (r.ok()) me = await r.json().catch(() => null);
  });

  const results = [];
  let n = 0;
  let skip = false;
  async function step(name, fn, opts = {}) {
    bucket = [];
    if (skip) {
      results.push({ name, problems: [], warns: ["SKIP: この地域に属さないユーザーのため確認できない"], info: "", file: "-" });
      return;
    }
    const problems = [];
    const warns = [];
    let info = "";
    try {
      info = (await fn({ warn: (m) => warns.push(m) })) ?? "";
      await page.waitForLoadState("networkidle", { timeout: 5000 }).catch(() => {});
      await page.waitForTimeout(500);
      const bad = await page.evaluate(() =>
        [...document.images].filter((i) => !(i.complete && i.naturalWidth > 0)).map((i) => i.getAttribute("src")),
      );
      if (bad.length) problems.push(`画像が読めない ${bad.length} 件: ${bad.slice(0, 3).join(", ")}`);
      const imgs = await page.evaluate(() => document.images.length);
      info += ` img=${imgs}`;
    } catch (e) {
      problems.push(`操作失敗: ${e.message.split("\n")[0]}`);
    }
    problems.push(...bucket.filter((m) => !(opts.ignore && opts.ignore.test(m))));
    const file = join(outDir, `${String(++n).padStart(2, "0")}-${name}.png`);
    await page.screenshot({ path: file, fullPage: true }).catch(() => {});
    results.push({ name, problems, warns, info, file });
  }

  const feedState = async () =>
    page.evaluate(() => ({
      posts: document.querySelectorAll("main article").length,
      empty: document.querySelector("main .state-view h2")?.textContent ?? "",
    }));

  await step("top", async () => {
    await page.goto(args.base + "/");
    await page.getByRole("button", { name: "ログイン" }).waitFor();
    return `url=${page.url()}`;
  });

  await step("login", async ({ warn }) => {
    const inputs = page.locator("form input");
    await inputs.nth(0).fill(args.user);
    await inputs.nth(1).fill(args.key);
    await page.getByRole("button", { name: "ログイン" }).click();
    await page.getByRole("navigation", { name: "メインナビゲーション" }).waitFor();
    await page.waitForLoadState("networkidle", { timeout: 5000 }).catch(() => {});
    for (let i = 0; i < 20 && !forbidden && !me; i++) await page.waitForTimeout(250);
    if (forbidden) {
      skip = true;
      warn("ログイン後の /api/me が 403（ユーザーの地域とアクセス先の地域が違う）。以降は確認しない");
    }
    return `url=${page.url()}`;
  }, { ignore: /403/ });

  let anyFeedPost = false;
  for (const [label, name] of [["My friends", "feed-friends"], ["Friends of friends", "feed-fof"], ["公開", "feed-public"], ["フォロー中", "feed-following"]]) {
    await step(name, async ({ warn }) => {
      await page.getByRole("navigation", { name: "フィード" }).getByRole("link", { name: label, exact: true }).click();
      await page.waitForSelector("main article, main .state-view");
      const s = await feedState();
      if (s.posts > 0) anyFeedPost = true;
      else warn(`投稿 0 件（${s.empty}）`);
      return `posts=${s.posts}`;
    });
  }

  await step("post-detail", async ({ warn }) => {
    // フィードに投稿があればそこから、無ければ Memories の過去投稿から開く
    let link = null;
    if (anyFeedPost) {
      for (const label of ["My friends", "Friends of friends", "公開", "フォロー中"]) {
        await page.getByRole("navigation", { name: "フィード" }).getByRole("link", { name: label, exact: true }).click();
        await page.waitForSelector("main article, main .state-view");
        if ((await feedState()).posts > 0) { link = page.locator("main article a").first(); break; }
      }
    }
    if (!link) {
      warn("フィードに投稿が無いため Memories の過去投稿を開いた");
      await page.getByRole("navigation", { name: "メインナビゲーション" }).getByRole("link", { name: "Memories" }).click();
      await page.waitForSelector("main article a");
      link = page.locator("main article a").first();
    }
    await link.click();
    await page.waitForSelector("main img");
    return `url=${page.url()}`;
  });

  await step("profile-menu", async () => {
    await page.getByRole("link", { name: "プロフィールメニュー" }).click();
    await page.getByRole("heading", { name: "プロフィールメニュー" }).waitFor();
  });

  await step("profile-settings", async () => {
    await page.getByRole("link", { name: /プロフィール設定/ }).click();
    await page.waitForURL("**/settings/profile");
    await page.waitForSelector("main input, main textarea");
    const text = await page.locator("main").innerText();
    const vals = await page.locator("main input, main textarea").evaluateAll((e) => e.map((x) => x.value));
    if (me && !(text + vals.join("\n")).includes(me.displayName)) throw new Error(`表示名 ${me.displayName} が画面に無い`);
    return `url=${page.url()}`;
  });

  await step("my-post", async ({ warn }) => {
    await page.goto(args.base + "/settings");
    await page.getByRole("link", { name: /自分の投稿/ }).click();
    await page.waitForSelector("main h1, main .state-view");
    const st = await feedState();
    if (st.empty) warn(`自分の投稿なし（${st.empty}）`);
    return `url=${page.url()}`;
  }, { ignore: /404/ });

  await ctx.close();
  return { region, results };
}

const browser = await chromium.launch();
let failed = false;
for (const region of regions) {
  const { results } = await checkRegion(browser, region);
  for (const r of results) {
    const ok = r.problems.length === 0;
    if (!ok) failed = true;
    console.log(`${ok ? "OK  " : "NG  "} [${region}] ${r.name} ${r.info.trim()}  ${r.file}`);
    for (const w of r.warns) console.log(`     WARN ${w}`);
    for (const p of r.problems) console.log(`     - ${p}`);
  }
}
await browser.close();
console.log(failed ? "結果: 問題あり" : "結果: すべて OK");
process.exit(failed ? 1 : 0);
