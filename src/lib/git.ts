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
