export interface Command {
  /** 一覧表示に出す 1 行説明 */
  summary: string;
  /** --help で出す全文 */
  usage: string;
  run(argv: string[]): Promise<void>;
}
