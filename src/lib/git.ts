import { usageError } from "./exit.js";
import { local } from "./ssh.js";

/**
 * git ref を完全な sha に解決する。
 * ref は「アプリリポジトリ」(config.app.repoLocal) で解決する。
 * このハーネス自身のリポジトリではない。
 */
export const resolveSha = async (repoLocal: string, ref = "HEAD"): Promise<string> => {
  try {
    const { stdout } = await local("git", ["-C", repoLocal, "rev-parse", "--verify", `${ref}^{commit}`]);
    return stdout.trim();
  } catch {
    throw usageError(`git ref を解決できません: ${ref} (in ${repoLocal})`);
  }
};

/**
 * 2 つの commit で、paths 配下の中身が同じか。commit が無い等で比べられなければ undefined。
 * ベンチの前に「狙った変更がコミットに入っていない」（revert やコミットの失敗で、ラン記録だけの
 * コミットになった等）を見つけるのに使う。
 */
export const sameContent = async (
  repoLocal: string,
  a: string,
  b: string,
  paths: readonly string[],
): Promise<boolean | undefined> => {
  const { exitCode } = await local("git", ["-C", repoLocal, "diff", "--quiet", a, b, "--", ...paths], {
    allowFailure: true,
  });
  // --quiet は差分が無ければ 0、あれば 1。それ以外は比べられなかった
  if (exitCode === 0) return true;
  if (exitCode === 1) return false;
  return undefined;
};
