#!/bin/sh
# 検証環境を落とす。鍵・ssh_config・ダミーリポジトリ（.work/）は残す。
# 完全に消すなら: rm -rf test/.work
set -eu
cd "$(dirname "$0")"
docker compose down -v --remove-orphans
