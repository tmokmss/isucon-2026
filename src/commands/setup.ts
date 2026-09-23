import { chmodSync, copyFileSync, existsSync, mkdirSync, mkdtempSync, readdirSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { homedir, tmpdir } from "node:os";
import { join, posix } from "node:path";
import { type Config, loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { configError, EXIT, HarnessError, usageError } from "../lib/exit.js";
import { log } from "../lib/log.js";
import { repoRoot } from "../lib/paths.js";
import { copyFrom, copyTo, local, onHost, onHostShell, shQuote } from "../lib/ssh.js";

/**
 * isuh setup — docs/DESIGN.md §3「setup」
 *
 * インスタンス起動直後の作業を 1 コマンドにする。どのステップも何度流しても同じ結果になる
 * （冪等）ように書く。当日は途中で失敗して直して流し直す、が普通に起きるため。
 */

const STEPS = ["ssh", "hosts", "dotfiles", "tools", "logs", "apprepo"] as const;
type Step = (typeof STEPS)[number];

const setupError = (message: string) => new HarnessError(EXIT.SETUP, message);

const assetsDir = (): string => join(repoRoot(), "setup");

// ---------------------------------------------------------------------------
// 管理ブロック
//
// 既存のファイル（~/.bashrc・/etc/hosts・~/.ssh/config）は丸ごと置き換えず、
// 印で囲んだブロックだけを書き換える。再実行してもブロックは 1 つのまま。
// ---------------------------------------------------------------------------

const BEGIN = "# >>> isuh >>>";
const END = "# <<< isuh <<<";

/** text 中の管理ブロックを body で置き換える（無ければ末尾に足す）。 */
export const upsertBlock = (text: string, body: string): string => {
  const kept: string[] = [];
  let inside = false;
  for (const line of text.split("\n")) {
    if (line === BEGIN) inside = true;
    else if (line === END) inside = false;
    else if (!inside) kept.push(line);
  }
  while (kept.length > 0 && kept[kept.length - 1] === "") kept.pop();
  const head = kept.length > 0 ? `${kept.join("\n")}\n` : "";
  return `${head}${BEGIN}\n${body.trimEnd()}\n${END}\n`;
};

/** リモートの `~/…` を $HOME 基準のシェル表現にする（クォートは残りの部分だけ）。 */
const remotePathExpr = (path: string): string =>
  path.startsWith("~/") ? `"$HOME"/${shQuote(path.slice(2))}` : shQuote(path);

/**
 * リモートのファイルの管理ブロックを書き換える。
 *
 * mv ではなく上書き（tee）で書く。コンテナの /etc/hosts はバインドマウントで、
 * 置き換えると "Device or resource busy" になるため（本物の EC2 でも害は無い）。
 */
const upsertRemoteBlock = async (
  host: string,
  path: string,
  body: string,
  opts: { sudo?: boolean; mode?: string } = {},
): Promise<void> => {
  const current = await onHostShell(host, `${opts.sudo ? "sudo " : ""}cat ${remotePathExpr(path)} 2>/dev/null || true`);
  const next = upsertBlock(current.stdout, body);
  if (next === current.stdout) return;

  const remoteTmp = `/tmp/isuh-block-${process.pid}`;
  await pushText(host, next, remoteTmp);
  const sudo = opts.sudo ? "sudo " : "";
  const target = remotePathExpr(path);
  const chmod = opts.mode ? ` && ${sudo}chmod ${opts.mode} ${target}` : "";
  await onHostShell(
    host,
    `${sudo}mkdir -p "$(dirname ${target})" && ${sudo}tee ${target} < ${remoteTmp} > /dev/null${chmod}; rc=$?; rm -f ${remoteTmp}; exit $rc`,
  );
};

/** 文字列をリモートのファイルとして送る。手元の一時ファイルは呼び出しごとに別にする（並列で送るため）。 */
const pushText = async (host: string, text: string, remotePath: string): Promise<void> => {
  const dir = mkdtempSync(join(tmpdir(), "isuh-setup-"));
  try {
    const staged = join(dir, "content");
    writeFileSync(staged, text);
    await copyTo(host, staged, remotePath);
  } finally {
    rmSync(dir, { recursive: true, force: true });
  }
};

// ---------------------------------------------------------------------------
// ssh: 手元の ~/.ssh と、サーバ間の鍵
// ---------------------------------------------------------------------------

const addressesOf = (cfg: Config) => {
  const missing = cfg.hosts.all.filter((h) => {
    const a = cfg.setup.addresses[h];
    return a === undefined || a.public === "" || a.private === "";
  });
  if (missing.length > 0) {
    throw configError(
      `setup.addresses に未設定のホストがあります: ${missing.join(", ")}\n` +
        `  当日 CloudFormation の出力（またはポータル）から public / private を埋めてください。`,
    );
  }
  return cfg.hosts.all.map((h) => ({ host: h, ...cfg.setup.addresses[h]! }));
};

const expandHome = (path: string): string => (path.startsWith("~/") ? join(homedir(), path.slice(2)) : path);

/** 手元の ssh 設定。ホスト別名 → public アドレス。 */
export const renderLocalSshConfig = (cfg: Config): string => {
  const lines = ["# isuh setup が生成する。手で編集しない（再実行で上書きされる）。", ""];
  for (const a of addressesOf(cfg)) {
    lines.push(`Host ${a.host}`, `  HostName ${a.public}`);
    if (a.port !== undefined) lines.push(`  Port ${a.port}`);
  }
  lines.push(
    "",
    `Host ${cfg.hosts.all.join(" ")}`,
    `  User ${cfg.hosts.user}`,
    `  IdentityFile ${expandHome(cfg.setup.identityFile)}`,
    "  IdentitiesOnly yes",
    // 競技のサーバは毎回作り直され、ホスト鍵も変わる。known_hosts を汚さない
    "  StrictHostKeyChecking no",
    "  UserKnownHostsFile /dev/null",
    "  LogLevel ERROR",
    "  ServerAliveInterval 30",
    "",
  );
  return lines.join("\n");
};

const INCLUDE_LINE = "Include isuh_config";

/**
 * ~/.ssh/config の先頭に Include を足した内容を返す。既にあれば null。
 * Include は最初の Host より前に置かないと、その Host ブロックの中に取り込まれてしまう。
 */
export const withInclude = (existing: string): string | null => {
  if (existing.split("\n").some((l) => l.trim() === INCLUDE_LINE)) return null;
  const comment = "# isuh setup が追加（isu1 / isu2 / isu3 の定義は ~/.ssh/isuh_config）";
  return `${comment}\n${INCLUDE_LINE}\n\n${existing}`;
};

/**
 * 手元の ssh 設定を書く。
 *
 * ISUH_SSH_CONFIG があればそこに書き、~/.ssh/config には触らない（test/ の環境）。
 * 無ければ ~/.ssh/isuh_config に書き、~/.ssh/config から Include する。
 * こうすると isuh だけでなく、人もエージェントも素の `ssh isu1` で入れる。
 */
export const writeLocalSshConfig = (cfg: Config, sshDir = join(homedir(), ".ssh")): string => {
  const body = renderLocalSshConfig(cfg);
  const override = process.env["ISUH_SSH_CONFIG"];
  if (override) {
    writeFileSync(override, body);
    return override;
  }

  mkdirSync(sshDir, { recursive: true, mode: 0o700 });
  const target = join(sshDir, "isuh_config");
  writeFileSync(target, body, { mode: 0o600 });

  const main = join(sshDir, "config");
  const existing = existsSync(main) ? readFileSync(main, "utf8") : "";
  const next = withInclude(existing);
  if (next !== null) {
    if (existsSync(main) && !existsSync(`${main}.isuh-bak`)) copyFileSync(main, `${main}.isuh-bak`);
    writeFileSync(main, next, { mode: 0o600 });
    chmodSync(main, 0o600);
    log.info(`~/.ssh/config の先頭に「${INCLUDE_LINE}」を足しました（元は config.isuh-bak）`);
  }
  return target;
};

/** サーバ間で使う鍵。手元で 1 本作って全台に配る（.isuh/ は gitignore 済み）。 */
const interServerKey = async (): Promise<string> => {
  const dir = join(repoRoot(), ".isuh", "setup");
  const key = join(dir, "id_ed25519");
  if (!existsSync(key)) {
    mkdirSync(dir, { recursive: true });
    await local("ssh-keygen", ["-q", "-t", "ed25519", "-N", "", "-C", "isuh-inter-server", "-f", key]);
  }
  return key;
};

const REMOTE_KEY = "~/.ssh/isuh_ed25519";

const renderServerSshConfig = (cfg: Config): string => {
  const lines: string[] = [];
  for (const a of addressesOf(cfg)) lines.push(`Host ${a.host}`, `  HostName ${a.private}`);
  lines.push(
    `Host ${cfg.hosts.all.join(" ")}`,
    `  User ${cfg.hosts.user}`,
    `  IdentityFile ${REMOTE_KEY}`,
    "  StrictHostKeyChecking no",
    "  UserKnownHostsFile /dev/null",
    "  LogLevel ERROR",
  );
  return lines.join("\n");
};

const stepSsh = async (cfg: Config): Promise<void> => {
  addressesOf(cfg);
  if (!existsSync(expandHome(cfg.setup.identityFile))) {
    throw configError(`setup.identityFile がありません: ${cfg.setup.identityFile}`);
  }

  const written = writeLocalSshConfig(cfg);
  log.info(`ssh: 手元の設定を書きました ${written}`);

  // まず全台に入れることを確かめる。ここで落ちるなら以降は全部落ちる
  const reach = await Promise.all(
    cfg.hosts.all.map(async (h) => [h, await onHost(h, ["true"], { allowFailure: true })] as const),
  );
  const unreachable = reach.filter(([, r]) => r.exitCode !== 0);
  if (unreachable.length > 0) {
    throw setupError(
      `ssh できないホストがあります: ${unreachable.map(([h]) => h).join(", ")}\n` +
        `  - setup.addresses の public が合っているか\n` +
        `  - ${cfg.setup.identityFile} の公開鍵をポータルに登録したか\n\n` +
        unreachable.map(([h, r]) => `${h}: ${r.stderr.trim()}`).join("\n"),
    );
  }

  const key = await interServerKey();
  const pub = readFileSync(`${key}.pub`, "utf8").trim();
  const serverConfig = renderServerSshConfig(cfg);
  await Promise.all(
    cfg.hosts.all.map(async (h) => {
      await copyTo(h, key, "/tmp/isuh_ed25519");
      await onHostShell(
        h,
        [
          `install -d -m 700 "$HOME/.ssh"`,
          `install -m 600 /tmp/isuh_ed25519 "$HOME/.ssh/isuh_ed25519"`,
          `rm -f /tmp/isuh_ed25519`,
          `touch "$HOME/.ssh/authorized_keys"`,
          `chmod 600 "$HOME/.ssh/authorized_keys"`,
          `grep -qxF ${shQuote(pub)} "$HOME/.ssh/authorized_keys" || printf '%s\\n' ${shQuote(pub)} >> "$HOME/.ssh/authorized_keys"`,
        ].join(" && "),
      );
      await upsertRemoteBlock(h, "~/.ssh/config", serverConfig, { mode: "600" });
    }),
  );

  // サーバ間で入れるかを、1 台目から全台へ確かめる
  const from = cfg.hosts.all[0]!;
  for (const to of cfg.hosts.all) {
    const r = await onHost(from, ["ssh", "-o", "BatchMode=yes", "-o", "ConnectTimeout=5", to, "hostname"], {
      allowFailure: true,
    });
    if (r.exitCode !== 0) throw setupError(`${from} から ${to} へ ssh できません\n${r.stderr.trim()}`);
  }
  log.info(`ssh: ${from} から全台へ入れることを確認しました`);
};

// ---------------------------------------------------------------------------
// hosts: /etc/hosts
// ---------------------------------------------------------------------------

const stepHosts = async (cfg: Config): Promise<void> => {
  const body = addressesOf(cfg)
    .map((a) => `${a.private}\t${a.host}`)
    .join("\n");
  await Promise.all(cfg.hosts.all.map((h) => upsertRemoteBlock(h, "/etc/hosts", body, { sudo: true })));

  for (const h of cfg.hosts.all) {
    for (const a of addressesOf(cfg)) {
      const r = await onHost(h, ["getent", "hosts", a.host], { allowFailure: true });
      const got = r.stdout.trim().split(/\s+/)[0];
      if (got !== a.private) throw setupError(`${h}: ${a.host} が ${a.private} に解決されません (got=${got || "なし"})`);
    }
  }
  log.info("hosts: 全台の /etc/hosts に private アドレスを書きました");
};

// ---------------------------------------------------------------------------
// dotfiles
// ---------------------------------------------------------------------------

/** setup/dotfiles/<name> → ~/.<name>。bashrc だけは ~/.bashrc から読み込ませる。 */
const DOTFILES = ["tmux.conf", "vimrc", "inputrc"];

const stepDotfiles = async (cfg: Config): Promise<void> => {
  const dir = join(assetsDir(), "dotfiles");
  await Promise.all(
    cfg.hosts.all.map(async (h) => {
      for (const name of DOTFILES) {
        const staged = `/tmp/isuh-dot-${name}`;
        await copyTo(h, join(dir, name), staged);
        // 元のファイルは 1 回目だけ退避する
        const dest = `"$HOME"/.${name}`;
        await onHostShell(
          h,
          `{ [ ! -f ${dest} ] || [ -f ${dest}.isuh-bak ] || cp ${dest} ${dest}.isuh-bak; } && install -m 644 ${staged} ${dest} && rm -f ${staged}`,
        );
      }
      await copyTo(h, join(dir, "bashrc"), "/tmp/isuh-dot-bashrc");
      await onHostShell(h, `install -m 644 /tmp/isuh-dot-bashrc "$HOME/.isuh.bashrc" && rm -f /tmp/isuh-dot-bashrc`);
      await upsertRemoteBlock(h, "~/.bashrc", "[ -f ~/.isuh.bashrc ] && . ~/.isuh.bashrc");
    }),
  );
  log.info(`dotfiles: ${["bashrc", ...DOTFILES].join(" / ")} を全台に配りました`);
};

// ---------------------------------------------------------------------------
// tools: 計測ツール（オフライン同梱を優先）
// ---------------------------------------------------------------------------

const ALP_VERSION = "1.0.21";
/** 同梱が無いもの。入らなくても計測は回る（vmstat / top は procps にある）ので警告だけ */
const APT_EXTRAS = ["htop", "sysstat"];

const archOf = async (host: string): Promise<"amd64" | "arm64" | null> => {
  const m = (await onHost(host, ["uname", "-m"])).stdout.trim();
  if (m === "x86_64" || m === "amd64") return "amd64";
  if (m === "aarch64" || m === "arm64") return "arm64";
  return null;
};

const installTools = async (host: string): Promise<void> => {
  const bin = join(assetsDir(), "bin");

  const alp = await onHost(host, ["alp", "--version"], { allowFailure: true });
  if (alp.exitCode === 0 && alp.stdout.includes(ALP_VERSION)) {
    log.info(`${host}: alp ${ALP_VERSION} は導入済み`);
  } else {
    const arch = await archOf(host);
    if (arch === null) throw setupError(`${host}: 同梱の alp が無いアーキテクチャです`);
    await copyTo(host, join(bin, `alp_linux_${arch}.tar.gz`), "/tmp/isuh-alp.tar.gz");
    await onHostShell(
      host,
      "rm -rf /tmp/isuh-alp && mkdir /tmp/isuh-alp && tar -xzf /tmp/isuh-alp.tar.gz -C /tmp/isuh-alp" +
        " && sudo install -m 755 /tmp/isuh-alp/alp /usr/local/bin/alp && rm -rf /tmp/isuh-alp /tmp/isuh-alp.tar.gz",
    );
  }

  // percona-toolkit のうち pt-query-digest だけを同梱している（perl の単体スクリプト）
  await copyTo(host, join(bin, "pt-query-digest"), "/tmp/isuh-pt-query-digest");
  await onHostShell(
    host,
    "sudo install -m 755 /tmp/isuh-pt-query-digest /usr/local/bin/pt-query-digest && rm -f /tmp/isuh-pt-query-digest",
  );

  // sudo で動くこと（collect は sudo で呼ぶ）まで確かめる
  for (const args of [
    ["sudo", "alp", "--version"],
    ["sudo", "pt-query-digest", "--version"],
  ]) {
    const r = await onHost(host, args, { allowFailure: true });
    if (r.exitCode !== 0) throw setupError(`${host}: ${args.slice(1).join(" ")} が動きません\n${r.stderr.trim()}`);
  }

  const missing: string[] = [];
  for (const pkg of APT_EXTRAS) {
    const cmd = pkg === "sysstat" ? "pidstat" : pkg;
    if ((await onHost(host, ["sh", "-c", `command -v ${cmd}`], { allowFailure: true })).exitCode !== 0) {
      missing.push(pkg);
    }
  }
  if (missing.length > 0) {
    // ネットワークが細いと apt は刺さる。時間を区切り、失敗しても止めない
    const r = await onHostShell(
      host,
      `sudo timeout 180 sh -c 'apt-get update -qq && DEBIAN_FRONTEND=noninteractive apt-get install -y -qq ${missing.join(" ")}' > /dev/null 2>&1`,
      { allowFailure: true },
    );
    if (r.exitCode !== 0) log.warn(`${host}: apt で ${missing.join(" ")} を入れられませんでした（計測には不要）`);
  }
  log.info(`${host}: alp / pt-query-digest を導入しました`);
};

const stepTools = async (cfg: Config): Promise<void> => {
  await Promise.all(cfg.hosts.all.map((h) => installTools(h)));
};

// ---------------------------------------------------------------------------
// logs: nginx の LTSV と MySQL の slow query log
// ---------------------------------------------------------------------------

const NGINX_CONF = "/etc/nginx/conf.d/isuh-log.conf";
const NGINX_BACKUP_DIR = "/var/backups/isuh";

const setupNginxLog = async (cfg: Config, host: string): Promise<void> => {
  if ((await onHost(host, ["sh", "-c", "command -v nginx"], { allowFailure: true })).exitCode !== 0) {
    log.warn(`${host}: nginx がありません。access log の設定を飛ばします`);
    return;
  }
  const hasConfD = await onHost(
    host,
    ["grep", "-Eq", "^[[:space:]]*include[[:space:]]+/etc/nginx/conf\\.d/\\*\\.conf", "/etc/nginx/nginx.conf"],
    { allowFailure: true },
  );
  if (hasConfD.exitCode !== 0) {
    log.warn(`${host}: nginx.conf が conf.d/*.conf を読んでいません。access log は conf/ で手当てしてください`);
    return;
  }

  // 既存の access_log を無効にする。残すと同じファイルに 2 形式で書かれ、alp が読めなくなる。
  // `access_log off` は静的ファイル向けなどの意図があるので残す。
  // 退避は nginx が読まない場所へ（sites-enabled/ に置くと include されてしまう）。
  const files = await onHostShell(
    host,
    `sudo grep -lE '^[[:space:]]*access_log[[:space:]]' /etc/nginx/nginx.conf /etc/nginx/sites-enabled/* /etc/nginx/conf.d/*.conf 2>/dev/null | grep -vxF ${shQuote(NGINX_CONF)} || true`,
  );
  for (const file of files.stdout.split("\n").filter((f) => f !== "")) {
    const backup = posix.join(NGINX_BACKUP_DIR, file);
    await onHostShell(
      host,
      `sudo mkdir -p ${shQuote(posix.dirname(backup))} && { [ -f ${shQuote(backup)} ] || sudo cp -L ${shQuote(file)} ${shQuote(backup)}; }` +
        ` && sudo sed -i --follow-symlinks -E '/access_log[[:space:]]+off/! s/^([[:space:]]*)(access_log[[:space:]])/\\1# isuh: \\2/' ${shQuote(file)}`,
    );
    log.info(`${host}: ${file} の access_log を無効にしました（元は ${backup}）`);
  }

  const conf = readFileSync(join(assetsDir(), "conf", "nginx-log.conf"), "utf8").replace(
    "{{ACCESS_LOG}}",
    cfg.middleware.nginxAccessLog,
  );
  await pushText(host, conf, "/tmp/isuh-log.conf");
  await onHostShell(host, `sudo install -m 644 /tmp/isuh-log.conf ${NGINX_CONF} && rm -f /tmp/isuh-log.conf`);

  const test = await onHost(host, ["sudo", "nginx", "-t"], { allowFailure: true });
  if (test.exitCode !== 0) {
    // 壊したまま残さない
    await onHostShell(
      host,
      `sudo rm -f ${NGINX_CONF}; cd ${NGINX_BACKUP_DIR} 2>/dev/null && sudo find . -type f | while read -r f; do sudo cp "$f" "/\${f#./}"; done`,
      { allowFailure: true },
    );
    throw setupError(`${host}: nginx -t が通りません。変更を戻しました\n${test.stderr.trim()}`);
  }
  await onHost(host, ["sudo", "systemctl", "reload", cfg.middleware.nginxService]);

  // 実際に 1 行書かれ、LTSV になっていることを確かめる
  await onHostShell(host, "curl -s -o /dev/null http://localhost/ || true");
  const last = await onHostShell(host, `sudo tail -n 1 ${shQuote(cfg.middleware.nginxAccessLog)} 2>/dev/null || true`);
  if (last.stdout.includes("\treqtime:")) log.info(`${host}: nginx の access log を LTSV にしました`);
  else log.warn(`${host}: access log に LTSV の行が見当たりません（localhost:80 で待ち受けていない？）`);

  if (cfg.metrics.alpFormat !== "ltsv") log.warn(`metrics.alpFormat が ${cfg.metrics.alpFormat} です。ltsv にしてください`);
};

const MYSQL_CONF_DIRS = ["/etc/mysql/mysql.conf.d", "/etc/mysql/mariadb.conf.d", "/etc/mysql/conf.d"];
const MYSQL_CONF_NAME = "zz-isuh-slowlog.cnf";

const setupSlowLog = async (cfg: Config, host: string): Promise<void> => {
  const hasDb = await onHost(host, ["sh", "-c", "command -v mysqld || command -v mariadbd"], { allowFailure: true });
  if (hasDb.exitCode !== 0) {
    log.warn(`${host}: mysqld / mariadbd がありません。slow query log の設定を飛ばします`);
    return;
  }
  const found = await onHostShell(
    host,
    `for d in ${MYSQL_CONF_DIRS.join(" ")}; do [ -d "$d" ] && { echo "$d"; exit 0; }; done; exit 1`,
    { allowFailure: true },
  );
  const confDir = found.stdout.trim();
  if (found.exitCode !== 0 || confDir === "") throw setupError(`${host}: MySQL の conf.d が見つかりません`);
  const dest = `${confDir}/${MYSQL_CONF_NAME}`;

  const conf = readFileSync(join(assetsDir(), "conf", "mysql-slowlog.cnf"), "utf8").replace(
    "{{SLOW_LOG}}",
    cfg.middleware.dbSlowLog,
  );
  await pushText(host, conf, "/tmp/isuh-slowlog.cnf");
  const logDir = posix.dirname(cfg.middleware.dbSlowLog);
  await onHostShell(
    host,
    `sudo install -m 644 /tmp/isuh-slowlog.cnf ${shQuote(dest)} && rm -f /tmp/isuh-slowlog.cnf` +
      ` && { [ -d ${shQuote(logDir)} ] || sudo install -d -o mysql -g mysql ${shQuote(logDir)}; }` +
      // root 所有のログが既にあると mysqld は開けず、黙って記録を止める（error.log に出るだけ）
      ` && sudo touch ${shQuote(cfg.middleware.dbSlowLog)} && sudo chown mysql:mysql ${shQuote(cfg.middleware.dbSlowLog)}`,
  );

  const restarted = await onHost(host, ["sudo", "systemctl", "restart", cfg.middleware.dbService], {
    allowFailure: true,
  });
  if (restarted.exitCode !== 0) {
    await onHost(host, ["sudo", "rm", "-f", dest]);
    await onHost(host, ["sudo", "systemctl", "restart", cfg.middleware.dbService], { allowFailure: true });
    throw setupError(`${host}: slow query log を入れると MySQL が起動しません。戻しました\n${restarted.stderr.trim()}`);
  }

  // 起動直後は接続できないことがあるので少し待つ
  let state = "";
  for (let i = 0; i < 30; i++) {
    const r = await onHost(host, ["sudo", "mysql", "-NBe", "SELECT @@global.slow_query_log, @@global.long_query_time"], {
      allowFailure: true,
    });
    if (r.exitCode === 0) {
      state = r.stdout.trim();
      break;
    }
    await new Promise((resolve) => setTimeout(resolve, 1000));
  }
  if (state === "") {
    log.warn(`${host}: sudo mysql で接続できず、slow query log が効いたか確認できません（${dest} は置きました）`);
  } else if (!/^1\s+0(\.0+)?$/.test(state)) {
    throw setupError(`${host}: slow query log が期待どおりになっていません (slow_query_log, long_query_time = ${state})`);
  } else {
    log.info(`${host}: slow query log を有効にしました（long_query_time = 0, ${cfg.middleware.dbSlowLog}）`);
  }
};

const stepLogs = async (cfg: Config): Promise<void> => {
  // 1 台が落ちても、他の台の再起動を途中で見捨てない（MySQL の restart 中に抜けると壊れうる）
  const results = await Promise.allSettled([
    ...cfg.hosts.nginx.map((h) => setupNginxLog(cfg, h)),
    cfg.hosts.db === "" ? Promise.resolve() : setupSlowLog(cfg, cfg.hosts.db),
  ]);
  const failed = results.find((r): r is PromiseRejectedResult => r.status === "rejected");
  if (failed) throw failed.reason;
};

// ---------------------------------------------------------------------------
// apprepo: サーバの ~/webapp を手元のアプリリポジトリにする
// ---------------------------------------------------------------------------

/** GitHub は 100MB を超えるファイルを拒否する。余裕を見てこれより大きいものは入れない */
const LARGE_FILE_MB = 50;

/**
 * deploy は手元のリポジトリから git archive で配るので、サーバ側を git にする必要は無い。
 * サーバの webapp には触らない（手元の初回コミットがそのまま元の状態の控えになる）。
 */
const stepApprepo = async (cfg: Config): Promise<void> => {
  const dest = cfg.app.repoLocal;
  if (existsSync(join(dest, ".git"))) {
    log.info(`apprepo: ${dest} は既に git リポジトリです。取り込みを飛ばします`);
    return;
  }
  if (existsSync(dest) && readdirSync(dest).length > 0) {
    throw setupError(`apprepo: ${dest} が空ではありません。消すか app.repoLocal を変えてください`);
  }

  const host = cfg.hosts.app[0]!;
  const src = cfg.app.dir.replace(/\/$/, "");
  const binary = `${cfg.app.langDir}/${cfg.app.binary}`;

  const large = await onHostShell(
    host,
    `cd ${shQuote(src)} && sudo find . -type f -size +${LARGE_FILE_MB}M -not -path './.git/*' | sed 's|^\\./||'`,
  );
  const largeFiles = large.stdout.split("\n").filter((f) => f !== "");

  const excludes = [".git", "node_modules", "target", binary, "*.isuh-new", ...largeFiles];
  const archive = "/tmp/isuh-webapp.tgz";
  await onHostShell(
    host,
    `sudo tar -C ${shQuote(src)} -czf ${archive} ${excludes.map((e) => `--exclude=${shQuote(`./${e}`)}`).join(" ")} . && sudo chmod 644 ${archive}`,
  );

  mkdirSync(dest, { recursive: true });
  const localArchive = join(tmpdir(), `isuh-webapp-${process.pid}.tgz`);
  try {
    await copyFrom(host, archive, localArchive);
    await onHost(host, ["sudo", "rm", "-f", archive]);
    await local("tar", ["-xzf", localArchive, "-C", dest]);
  } finally {
    rmSync(localArchive, { force: true });
  }

  const ignore = [
    "# isuh setup apprepo が生成",
    `/${binary}`,
    "*.isuh-new",
    "node_modules/",
    "target/",
  ];
  if (largeFiles.length > 0) {
    ignore.push(`# ${LARGE_FILE_MB}MB を超えるため取り込んでいない（サーバの ${src} にだけある）`);
    ignore.push(...largeFiles.map((f) => `/${f}`));
  }
  const gitignore = join(dest, ".gitignore");
  const prior = existsSync(gitignore) ? `${readFileSync(gitignore, "utf8").trimEnd()}\n\n` : "";
  writeFileSync(gitignore, `${prior}${ignore.join("\n")}\n`);

  await local("git", ["init", "-q", "-b", "main"], { cwd: dest });
  await local("git", ["add", "-A"], { cwd: dest });
  await local("git", ["commit", "-q", "-m", `${host}:${src} の初期状態`], { cwd: dest });
  const sha = (await local("git", ["rev-parse", "--short", "HEAD"], { cwd: dest })).stdout.trim();
  log.info(`apprepo: ${host}:${src} を ${dest} に取り込みました (${sha})`);
  if (largeFiles.length > 0) log.warn(`apprepo: 大きいファイルは取り込んでいません: ${largeFiles.join(", ")}`);

  if (cfg.setup.appRemote === "") {
    log.warn("apprepo: setup.appRemote が空なので push しません");
    return;
  }
  await local("git", ["remote", "add", "origin", cfg.setup.appRemote], { cwd: dest });
  const pushed = await local("git", ["push", "-q", "-u", "origin", "main"], { cwd: dest, allowFailure: true });
  if (pushed.exitCode !== 0) {
    throw setupError(
      `apprepo: ${cfg.setup.appRemote} へ push できません（取り込みとコミットは済んでいます）\n${pushed.stderr.trim()}`,
    );
  }
  log.info(`apprepo: ${cfg.setup.appRemote} へ push しました`);
};

// ---------------------------------------------------------------------------

const RUNNERS: Record<Step, (cfg: Config) => Promise<void>> = {
  ssh: stepSsh,
  hosts: stepHosts,
  dotfiles: stepDotfiles,
  tools: stepTools,
  logs: stepLogs,
  apprepo: stepApprepo,
};

export const command: Command = {
  summary: "インスタンス起動後のセットアップ",
  usage: `usage: isuh setup [--only <step>]

インスタンス起動後のセットアップを 1 コマンドで行う。どのステップも何度流しても同じ結果になる。
IaC は運営から渡されるものを使うため触れない。自動化できるのは「起動した後」だけで、
ここを詰めると当日の最初の 30〜60 分が浮く。

先に config.ts の setup.addresses（各ホストの public / private アドレス）を埋めること。

  --only <step>   指定したステップだけ実行する（カンマ区切りで複数可）

steps（この順で実行する）:
  ssh        手元の ~/.ssh/isuh_config を生成して ~/.ssh/config から Include、
             サーバ間の鍵を配って isu1 → 全台に ssh できるようにする
  hosts      全台の /etc/hosts に private アドレスを書く
  dotfiles   .bashrc / .tmux.conf / .vimrc / .inputrc の配置
  tools      alp / pt-query-digest の導入（リポジトリ同梱のものを配る）。htop / sysstat は apt で試す
  logs       nginx の access log を LTSV に、MySQL の slow query log を long_query_time=0 で有効に
  apprepo    hosts.app[0] の app.dir を手元の app.repoLocal に取り込んで git 化し、setup.appRemote へ push

終了コード: 0 成功 / 64 引数不正 / 65 設定不備 / 60 セットアップ失敗`,

  async run(argv) {
    let steps: Step[] = [...STEPS];
    for (let i = 0; i < argv.length; i++) {
      const arg = argv[i]!;
      if (arg === "--only") {
        const value = argv[++i];
        if (value === undefined) throw usageError("--only に値がありません");
        const picked = value.split(",").filter((s) => s !== "");
        const unknown = picked.filter((s) => !(STEPS as readonly string[]).includes(s));
        if (unknown.length > 0 || picked.length === 0) {
          throw usageError(`不明なステップ: ${unknown.join(", ") || "(空)"}（${STEPS.join(" / ")}）`);
        }
        steps = STEPS.filter((s) => picked.includes(s));
      } else {
        throw usageError(`不明な引数: ${arg}`);
      }
    }
    const cfg = await loadConfig();
    log.info(`setup hosts=[${cfg.hosts.all.join(" ")}] steps=[${steps.join(" ")}]`);

    for (const step of steps) {
      const startedAt = Date.now();
      log.info(`=== ${step}`);
      await RUNNERS[step](cfg);
      log.info(`=== ${step} 完了 ${((Date.now() - startedAt) / 1000).toFixed(1)}s`);
    }
    log.info("setup: すべて完了しました");
  },
};
