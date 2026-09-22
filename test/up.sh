#!/bin/sh
# 検証環境を起動する。
#   - isu1 / isu2 / isu3 の sshd コンテナ
#   - 鍵と ssh_config（手元の ~/.ssh/config は触らない）
#   - ダミーのアプリリポジトリ（v1 → v2 の 2 コミット + ビルドが壊れた broken ブランチ）
set -eu
cd "$(dirname "$0")"

WORK=.work
mkdir -p "$WORK"

# --- 鍵と ssh_config ------------------------------------------------------
if [ ! -f "$WORK/id_ed25519" ]; then
  ssh-keygen -q -t ed25519 -N '' -C isuh-test -f "$WORK/id_ed25519"
fi
cp "$WORK/id_ed25519.pub" "$WORK/authorized_keys"

cat >"$WORK/ssh_config" <<EOF
# isuh の検証用。ISUH_SSH_CONFIG でこれを指す。
Host isu1
  Port 22101
Host isu2
  Port 22102
Host isu3
  Port 22103

Host isu1 isu2 isu3
  HostName 127.0.0.1
  User isucon
  IdentityFile $(pwd)/$WORK/id_ed25519
  IdentitiesOnly yes
  StrictHostKeyChecking no
  UserKnownHostsFile /dev/null
  LogLevel ERROR
EOF

# --- コンテナ -------------------------------------------------------------
docker compose up -d --build

echo "sshd の応答を待っています..."
for h in isu1 isu2 isu3; do
  i=0
  until ssh -F "$WORK/ssh_config" -o ConnectTimeout=2 "$h" true 2>/dev/null; do
    i=$((i + 1))
    if [ "$i" -gt 60 ]; then
      echo "$h に ssh できません" >&2
      exit 1
    fi
    sleep 1
  done
  echo "  $h ok"
done

# --- ダミーアプリリポジトリ ------------------------------------------------
APP="$WORK/app-repo"

write_shape() {
  # $1 = items の name に付ける JSON キー
  cat >"$APP/go/shape.go" <<GO
package main

type Item struct {
	ID   int      \`json:"id"\`
	Name string   \`json:"$1"\`
	Tags []string \`json:"tags"\`
}
GO
}

if [ ! -d "$APP/.git" ]; then
  rm -rf "$APP"
  mkdir -p "$APP"
  cp -R fixtures/app/. "$APP/"

  git -C "$APP" init -q -b main
  git -C "$APP" config user.email isuh-test@example.invalid
  git -C "$APP" config user.name "isuh test"

  printf 'package main\n\nconst version = "v1"\n' >"$APP/go/version.go"
  write_shape name
  git -C "$APP" add -A
  git -C "$APP" commit -qm "v1"

  printf 'package main\n\nconst version = "v2"\n' >"$APP/go/version.go"
  git -C "$APP" add -A
  git -C "$APP" commit -qm "v2"

  # ビルドが通らないブランチ。deploy が非ゼロで落ちることの確認に使う。
  git -C "$APP" checkout -q -b broken
  printf 'package main\n\nthis is not go\n' >"$APP/go/version.go"
  git -C "$APP" add -A
  git -C "$APP" commit -qm "broken"
  git -C "$APP" checkout -q main

  # ビルドは通るがレスポンスの構造が変わるブランチ（= レギュレーション違反）。
  # check がこれを捕まえられることの確認に使う。
  git -C "$APP" checkout -q -b badshape
  write_shape title
  git -C "$APP" add -A
  git -C "$APP" commit -qm "rename items[].name to title"
  git -C "$APP" checkout -q main
fi

echo
echo "準備できました。"
echo "  export ISUH_CONFIG=$(pwd)/config.test.ts"
echo "  export ISUH_SSH_CONFIG=$(pwd)/$WORK/ssh_config"
echo "  ./test/all.sh     # typecheck + 全スイート"
