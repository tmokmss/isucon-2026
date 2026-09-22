/**
 * 終了コード — docs/DESIGN.md §5
 * 呼び出し側が失敗理由で分岐するため、数値を直書きせずこの定数を使う。
 */
export const EXIT = {
  OK: 0,
  LOCK: 10, // ロック取得失敗
  DEPLOY: 20, // デプロイ失敗
  CHECK: 30, // 検証ゲート失敗
  BENCH: 40, // ベンチ失敗
  REGRESSION: 50, // スコア低下
  USAGE: 64, // 引数不正
  CONFIG: 65, // 設定不備
  NOT_IMPLEMENTED: 70, // 未実装
} as const;

export type ExitCode = (typeof EXIT)[keyof typeof EXIT];

/** 終了コードを伴うエラー。cli.ts が捕まえて process.exit する。 */
export class HarnessError extends Error {
  constructor(
    readonly code: ExitCode,
    message: string,
  ) {
    super(message);
    this.name = "HarnessError";
  }
}

export const usageError = (message: string) => new HarnessError(EXIT.USAGE, message);
export const configError = (message: string) => new HarnessError(EXIT.CONFIG, message);

/** スタブが呼ばれたときに担当 Issue を示して終了する。 */
export class NotImplementedError extends HarnessError {
  constructor(readonly issue: number) {
    super(
      EXIT.NOT_IMPLEMENTED,
      `未実装です。実装は #${issue}: https://github.com/tmokmss/isucon-harness/issues/${issue}\n` +
        `  契約は docs/DESIGN.md を参照。契約を変える必要があるなら実装せず先に相談すること。`,
    );
  }
}

export const notImplemented = (issue: number): never => {
  throw new NotImplementedError(issue);
};
