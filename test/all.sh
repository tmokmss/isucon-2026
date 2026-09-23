#!/bin/sh
# 全部まとめて回す。先に ./test/up.sh を実行しておくこと。
set -eu
cd "$(dirname "$0")"

echo "### typecheck"
(cd .. && npm run --silent typecheck)
echo "  ok"
echo

for suite in smoke lock portal report; do
  echo "### $suite"
  "./$suite.sh"
  echo
done

echo "すべて通りました"
