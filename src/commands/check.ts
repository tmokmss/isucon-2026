import { existsSync, mkdirSync, readFileSync, writeFileSync } from "node:fs";
import { dirname, isAbsolute, join } from "node:path";
import { type CheckEndpoint, type Config, loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { EXIT, HarnessError, usageError } from "../lib/exit.js";
import { log } from "../lib/log.js";
import { repoRoot } from "../lib/paths.js";
import { type Shape, diffShapes, shapeOf } from "../lib/shape.js";

const checkError = (message: string): HarnessError => new HarnessError(EXIT.CHECK, message);

/**
 * 記録した「正常な形」。git で追跡する。
 *
 * レスポンス構造の変更はレギュレーション違反（失格）なので、正しくは
 * **一度記録したら二度と更新しない**。エンドポイントを足したときだけ増える。
 */
interface Snapshot {
  recordedAt: string;
  entries: Record<string, SnapshotEntry>;
}

interface SnapshotEntry {
  status: number;
  json: boolean;
  shape?: Shape;
}

const snapshotPath = (cfg: Config): string => {
  const configured = cfg.check.snapshotFile === "" ? "check/structure.json" : cfg.check.snapshotFile;
  return isAbsolute(configured) ? configured : join(repoRoot(), configured);
};

const loadSnapshot = (path: string): Snapshot => {
  if (!existsSync(path)) return { recordedAt: "", entries: {} };
  try {
    const parsed = JSON.parse(readFileSync(path, "utf8")) as Snapshot;
    return { recordedAt: parsed.recordedAt ?? "", entries: parsed.entries ?? {} };
  } catch {
    throw checkError(`スナップショットを読めません: ${path}\n  壊れているなら削除して取り直してください。`);
  }
};

const saveSnapshot = (path: string, snapshot: Snapshot): void => {
  mkdirSync(dirname(path), { recursive: true });
  writeFileSync(path, `${JSON.stringify(snapshot, null, 2)}\n`);
};

const keyOf = (endpoint: CheckEndpoint): string =>
  `${endpoint.method ?? "GET"} ${endpoint.path}`;

interface Probe {
  key: string;
  status: number;
  json: boolean;
  shape: Shape | undefined;
  error: string | undefined;
}

const request = async (cfg: Config, endpoint: CheckEndpoint): Promise<Probe> => {
  const key = keyOf(endpoint);
  const base = cfg.app.baseUrl.replace(/\/$/, "");
  const url = `${base}${endpoint.path.startsWith("/") ? "" : "/"}${endpoint.path}`;
  const method = endpoint.method ?? "GET";

  const headers: Record<string, string> = { Accept: "application/json", ...endpoint.headers };
  let body: string | undefined;
  if (endpoint.body !== undefined) {
    body = JSON.stringify(endpoint.body);
    headers["Content-Type"] = "application/json";
  }

  try {
    const response = await fetch(url, {
      method,
      headers,
      body,
      signal: AbortSignal.timeout(cfg.check.timeoutSec * 1000),
    });
    const text = await response.text();
    const isJson = /json/i.test(response.headers.get("content-type") ?? "");
    let shape: Shape | undefined;
    if (isJson) {
      try {
        shape = shapeOf(JSON.parse(text));
      } catch {
        return { key, status: response.status, json: true, shape: undefined, error: `JSON として読めません: ${text.slice(0, 200)}` };
      }
    }
    return { key, status: response.status, json: isJson, shape, error: undefined };
  } catch (error) {
    return {
      key,
      status: 0,
      json: false,
      shape: undefined,
      error: `リクエストに失敗: ${error instanceof Error ? error.message : String(error)} (${url})`,
    };
  }
};

const runInitialize = async (cfg: Config): Promise<void> => {
  const base = cfg.app.baseUrl.replace(/\/$/, "");
  const url = `${base}${cfg.app.initializePath}`;
  log.info(`initialize: POST ${url}`);
  const startedAt = Date.now();
  try {
    const response = await fetch(url, {
      method: "POST",
      signal: AbortSignal.timeout(cfg.check.timeoutSec * 1000),
    });
    if (!response.ok) {
      throw checkError(`initialize が失敗しました: HTTP ${response.status} ${url}`);
    }
  } catch (error) {
    if (error instanceof HarnessError) throw error;
    throw checkError(
      `initialize を叩けませんでした: ${url}\n  ${error instanceof Error ? error.message : String(error)}`,
    );
  }
  log.info(`initialize: ok ${((Date.now() - startedAt) / 1000).toFixed(1)}s`);
};

export const command: Command = {
  summary: "ベンチを使わない検証ゲート",
  usage: `usage: isuh check

ベンチを使わずに「壊れていないこと」だけを数秒で確認する。
bench が deploy の直後に自動実行し、落ちた場合はベンチを起動しない。
ベンチは直列でしか回せないため、正しさの確認にベンチ枠を使わないのが狙い。

確認すること:
  - app.initializePath を叩いて成功すること（check.initialize が true のとき）
  - check.endpoints が期待どおりのステータスを返すこと
  - レスポンスの「構造」が記録した形と一致すること（値ではなくキーと型の形）

構造は check.snapshotFile に記録する。知らないエンドポイントは初回に記録され、
以後はそれと比較される。**アプリを触る前に 1 回走らせて、記録をコミットしておくこと。**

レスポンス構造の変更はレギュレーション違反（失格）なので、
記録した形は原則として更新しない。エンドポイントを足したときだけ増える。

値ではなく形だけを見るのは、値が initialize やベンチのたびに変わるため。
null を取りうるフィールドの型の揺れは差分として扱わない。

終了コード: 0 成功 / 30 検証失敗 / 65 設定不備`,

  async run(argv) {
    if (argv.length > 0) throw usageError(`引数は取らない: ${argv.join(" ")}`);
    const cfg = await loadConfig();

    if (cfg.check.endpoints.length === 0 && !cfg.check.initialize) {
      log.warn("check.endpoints が空で initialize も無効です。検証していません（config.ts を埋めてください）");
      return;
    }

    const startedAt = Date.now();
    if (cfg.check.initialize) await runInitialize(cfg);

    if (cfg.check.endpoints.length === 0) {
      log.warn("check.endpoints が空です。構造の検証は行われません（config.ts を埋めてください）");
      return;
    }

    const path = snapshotPath(cfg);
    const snapshot = loadSnapshot(path);
    const probes = await Promise.all(cfg.check.endpoints.map((e) => request(cfg, e)));

    const problems: string[] = [];
    const recorded: string[] = [];

    for (const [index, probe] of probes.entries()) {
      const endpoint = cfg.check.endpoints[index]!;
      const expectedStatus = endpoint.status ?? 200;

      if (probe.error !== undefined) {
        problems.push(`${probe.key}: ${probe.error}`);
        continue;
      }
      if (probe.status !== expectedStatus) {
        problems.push(`${probe.key}: ステータスが違う ${expectedStatus} → ${probe.status}`);
        continue;
      }

      const known = snapshot.entries[probe.key];
      if (known === undefined) {
        snapshot.entries[probe.key] = {
          status: probe.status,
          json: probe.json,
          ...(probe.shape === undefined ? {} : { shape: probe.shape }),
        };
        recorded.push(probe.key);
        continue;
      }

      if (known.json && !probe.json) {
        problems.push(`${probe.key}: JSON を返さなくなった`);
        continue;
      }
      if (known.shape !== undefined && probe.shape !== undefined) {
        const diffs = diffShapes(known.shape, probe.shape);
        for (const d of diffs) problems.push(`${probe.key} ${d}`);
      }
    }

    if (recorded.length > 0) {
      snapshot.recordedAt = snapshot.recordedAt === "" ? new Date().toISOString() : snapshot.recordedAt;
      saveSnapshot(path, snapshot);
      log.warn(
        `構造を新しく記録しました (${recorded.length} 件): ${recorded.join(", ")}\n` +
          `  → ${path}\n` +
          `  これが「正しい形」の基準になります。アプリを壊した後に記録していないか確認し、` +
          `問題なければコミットしてください。`,
      );
    }

    if (problems.length > 0) {
      throw checkError(
        `レスポンスが記録した形と違います (${problems.length} 件)\n` +
          problems.map((p) => `  - ${p}`).join("\n") +
          `\n\n  レスポンス構造の変更はレギュレーション違反です。直すまでベンチは走らせません。` +
          `\n  記録: ${path}`,
      );
    }

    log.info(
      `check ok: ${cfg.check.endpoints.length} エンドポイント ${((Date.now() - startedAt) / 1000).toFixed(1)}s`,
    );
  },
};
