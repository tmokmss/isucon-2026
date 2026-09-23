#!/bin/sh
set -eu

if [ -f /authorized_keys ]; then
  install -o isucon -g isucon -m 600 /authorized_keys /home/isucon/.ssh/authorized_keys
fi

ssh-keygen -A >/dev/null
exec /usr/sbin/sshd -D -e
