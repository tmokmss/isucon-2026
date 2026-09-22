#!/usr/bin/env bash
# bin/_common.sh — 全スクリプト共通のヘルパ
# 契約: docs/DESIGN.md §8。ここにある関数は各スクリプトで再実装しない。
# shellcheck shell=bash

[[ -n "${_ISUH_COMMON_LOADED:-}" ]] && return 0
_ISUH_COMMON_LOADED=1

ISUH_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
readonly ISUH_ROOT

# ---------------------------------------------------------------------------
# 終了コード — docs/DESIGN.md §5
#   呼び出し側が失敗理由で分岐するため、数値を直書きせずこの定数を使う。
# ---------------------------------------------------------------------------
readonly EX_OK=0
readonly EX_LOCK=10          # ロック取得失敗
readonly EX_DEPLOY=20        # デプロイ失敗
readonly EX_CHECK=30         # 検証ゲート失敗
readonly EX_BENCH=40         # ベンチ失敗
readonly EX_REGRESSION=50    # スコア低下
readonly EX_USAGE=64         # 引数不正
readonly EX_CONFIG=65        # 設定不備
readonly EX_NOTIMPL=70       # 未実装

# ---------------------------------------------------------------------------
# ログ — docs/DESIGN.md §6
#   stdout は機械可読な結果のみ。ログはすべて stderr。
# ---------------------------------------------------------------------------
_log() {
    local level="$1"; shift
    printf '[%s] %-5s %s\n' "$(date +%H:%M:%S)" "$level" "$*" >&2
}
log_debug() { [[ "${VERBOSE:-0}" == "1" ]] && _log DEBUG "$@"; return 0; }
log_info()  { _log INFO  "$@"; }
log_warn()  { _log WARN  "$@"; }
log_error() { _log ERROR "$@"; }

# die <exit-code> <message>...
die() {
    local code="$1"; shift
    log_error "$@"
    exit "$code"
}

# not_implemented <issue-number>
#   スタブが呼ばれたときに担当 Issue を示して終了する。
not_implemented() {
    local issue="${1#\#}"
    log_warn "未実装です。実装は #${issue}: https://github.com/tmokmss/isucon-harness/issues/${issue}"
    log_warn "契約は docs/DESIGN.md を参照。契約を変える必要があるなら実装せず先に相談すること。"
    exit "$EX_NOTIMPL"
}

repo_root() { printf '%s\n' "$ISUH_ROOT"; }

# ---------------------------------------------------------------------------
# 設定 — docs/DESIGN.md §2
# ---------------------------------------------------------------------------

# require_config <name>...
require_config() {
    local missing=() name
    for name in "$@"; do
        [[ -z "${!name:-}" ]] && missing+=("$name")
    done
    (( ${#missing[@]} > 0 )) \
        && die "$EX_CONFIG" "config.env に未設定の項目があります: ${missing[*]}"
    return 0
}

# load_config
#   ISUH_CONFIG で場所を上書きできる（テスト用）。
load_config() {
    local f="${ISUH_CONFIG:-$ISUH_ROOT/config.env}"
    if [[ ! -f "$f" ]]; then
        die "$EX_CONFIG" "設定ファイルがありません: $f
  cp config.env.example config.env して埋めてください。"
    fi
    set -a
    # shellcheck disable=SC1090
    . "$f"
    set +a
    require_config ISU_HOSTS ISU_PRIMARY ISU_APP_HOSTS APP_DIR APP_SERVICE APP_REPO_LOCAL
    log_debug "設定を読み込みました: $f"
}

# ---------------------------------------------------------------------------
# ssh
# ---------------------------------------------------------------------------

# on_host <host> <command>...
on_host() {
    local host="$1"; shift
    log_debug "ssh $host -- $*"
    ssh -o ConnectTimeout=5 "$host" "$@"
}

# on_hosts "<host1 host2 ...>" <command>...
#   1 台でも失敗したら最後の失敗コードを返す。
on_hosts() {
    local hosts="$1"; shift
    local h rc=0
    for h in $hosts; do
        on_host "$h" "$@" || rc=$?
    done
    return "$rc"
}

# ---------------------------------------------------------------------------
# git / run
#   git ref は「アプリリポジトリ」($APP_REPO_LOCAL) で解決する。
#   このハーネス自身のリポジトリではない。
# ---------------------------------------------------------------------------

# resolve_sha [<ref>]
resolve_sha() {
    local ref="${1:-HEAD}"
    git -C "$APP_REPO_LOCAL" rev-parse --verify "${ref}^{commit}" 2>/dev/null \
        || die "$EX_USAGE" "git ref を解決できません: $ref (in $APP_REPO_LOCAL)"
}

# new_run_id <sha>
new_run_id() {
    local sha="${1:?sha required}"
    printf '%s-%s\n' "$(date +%Y%m%d-%H%M%S)" "${sha:0:7}"
}

# run_dir <run-id>
run_dir() {
    local id="${1:?run-id required}"
    local d="$ISUH_ROOT/runs/$id"
    mkdir -p "$d"
    printf '%s\n' "$d"
}

history_file() { printf '%s\n' "$ISUH_ROOT/runs/history.tsv"; }

# ---------------------------------------------------------------------------
# ロック — docs/DESIGN.md §7   実装: #2
#   ISU_PRIMARY 上の flock を ssh セッションで保持する。
#   ssh プロセスが死ねば flock も解放されるため、Ctrl-C や異常終了でロックが残らない。
# ---------------------------------------------------------------------------

# lock_acquire <owner-label>
#   取得できなければ保持者を表示して exit 10。
#   解放用に ISUH_LOCK_PID を設定し、trap で lock_release を呼べるようにする。
lock_acquire() { not_implemented 2; }

# lock_release
#   ロックを保持している ssh プロセスを終了させる。多重呼び出し安全であること。
lock_release() { not_implemented 2; }

# ---------------------------------------------------------------------------
# 引数
# ---------------------------------------------------------------------------

# usage_guard "$@"  — 呼び出し元が usage() を定義していること
usage_guard() {
    local a
    for a in "$@"; do
        case "$a" in
            -h|--help) usage; exit "$EX_OK" ;;
        esac
    done
}
