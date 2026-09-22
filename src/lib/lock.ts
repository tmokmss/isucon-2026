import type { Config } from "../config.js";
import { notImplemented } from "./exit.js";

/**
 * ベンチのロック — docs/DESIGN.md §7   実装: #2
 *
 * デプロイだけをロックしても不十分。ベンチ実行中に別のデプロイが走れば計測は壊れる。
 * bench の全ステップがひとつのクリティカルセクション。
 *
 * 保持方法:
 *   hosts.primary 上の flock を、ssh の子プロセスを生かしたまま保持する。
 *
 *     ssh <primary> "exec 9><lockPath>; flock -w <timeout> 9 || exit 10; echo ACQUIRED; cat"
 *
 *   末尾の cat が stdin を開いたまま保つため、ssh プロセスが生きている間だけ
 *   ロックが保持される。子プロセスを kill すれば flock は自動解放される。
 *   Node なら process.on("exit") / try-finally で確実に kill できるので、
 *   Ctrl-C でも異常終了でもロックが残らない。
 *
 *   ローカルのファイルロックではなくサーバ側に置く理由は、worktree・別ターミナル・
 *   別マシンをまたいで効かせるため。
 *
 * 取得できなかった場合は、誰がいつから保持しているかを stderr に出して EXIT.LOCK で落ちる。
 * 保持者は <lockPath>.owner に「ホスト名 / ユーザ / 開始時刻 / git ref」を書いておく。
 */
export interface Lock {
  /** 多重呼び出し安全であること。 */
  release(): Promise<void>;
}

export const acquireLock = async (_owner: string, _cfg: Config): Promise<Lock> => notImplemented(2);
