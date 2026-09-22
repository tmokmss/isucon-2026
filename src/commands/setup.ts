import { loadConfig } from "../config.js";
import type { Command } from "../lib/command.js";
import { notImplemented, usageError } from "../lib/exit.js";
import { log } from "../lib/log.js";

export const command: Command = {
  summary: "インスタンス起動後のセットアップ",
  usage: `usage: isuh setup [--only <step>]

インスタンス起動後のセットアップを 1 コマンドで行う。
IaC は運営から渡されるものを使うため触れない。自動化できるのは「起動した後」だけで、
ここを詰めると当日の最初の 30〜60 分が浮く。

  --only <step>   指定したステップだけ実行する

steps:
  ssh        SSH 鍵の生成と全台への配布、~/.ssh/config（isu1/isu2/isu3）生成
  hosts      /etc/hosts の整備
  dotfiles   .bashrc / .tmux.conf / .vimrc / .inputrc の配置
  tools      alp / percona-toolkit / dstat / htop の導入（オフライン同梱分を優先）
  apprepo    ~/webapp を git 化して自分のリポジトリへ push（元は webapp_bak へ退避）

終了コード: 0 成功 / 64 引数不正 / 65 設定不備`,

  async run(argv) {
    let only = "";
    for (let i = 0; i < argv.length; i++) {
      const arg = argv[i]!;
      if (arg === "--only") {
        const value = argv[++i];
        if (value === undefined) throw usageError("--only に値がありません");
        only = value;
      } else {
        throw usageError(`不明な引数: ${arg}`);
      }
    }
    const cfg = await loadConfig();
    log.info(`setup hosts=[${cfg.hosts.all.join(" ")}] only=${only || "<all>"}`);
    notImplemented(9);
  },
};
