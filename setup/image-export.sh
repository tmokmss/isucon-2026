#!/bin/bash
# 初期データの画像を、各アプリ台の手元の PostgreSQL（AMI から来た止まったデータ）からファイルに書き出す。
# ルーター越しのコピーはしない。手元のデータが asia2 の初期データと同じことを照合してから書き出す。
# 使い方（手元の PC から）: setup/image-export.sh isu2
#   1. 止まっている PostgreSQL を 127.0.0.1:5433・ソケット /tmp/pgimg で一時的に起動する（5432 のアプリの接続先と取り違えない）
#   2. 境界以前の image_objects の (id, owner, media_type, created_at) と、id 先頭 '000' のバイト列の md5 を asia2 と比べる
#   3. nice・ionice を付けて /home/isucon/images/init/ に書き出す（既にあるファイルは飛ばす）
#   4. PostgreSQL を止め、postgresql.service が disabled・inactive のままであることを確かめる
set -euo pipefail
host="$1"
db_host="${DB_HOST:-isu4}"
here="$(cd "$(dirname "$0")" && pwd)"
gosrc="$here/../webapp/go"

(cd "$gosrc" && GOOS=linux GOARCH=amd64 go build -o /tmp/image-export ./cmd/image-export)
scp -q /tmp/image-export "$host":/tmp/image-export

ssh "$host" 'sudo mkdir -p /tmp/pgimg && sudo chown postgres /tmp/pgimg && (sudo -u postgres /usr/lib/postgresql/16/bin/pg_ctl -D /var/lib/postgresql/16/main status >/dev/null || sudo -u postgres nice -n 10 /usr/lib/postgresql/16/bin/pg_ctl -D /var/lib/postgresql/16/main -l /tmp/pgimg/log -w -o "-c config_file=/etc/postgresql/16/main/postgresql.conf -p 5433 -c listen_addresses=127.0.0.1 -c unix_socket_directories=/tmp/pgimg -c shared_buffers=64MB" start)'

check_sql='SELECT count(*), md5(string_agg(id || chr(58) || owner_user_id || chr(58) || media_type || chr(58) || created_at::text, chr(44) ORDER BY id)) FROM image_objects WHERE created_at <= (SELECT created_at FROM initial_data_boundary); SELECT count(*), md5(string_agg(id || chr(58) || md5(data), chr(44) ORDER BY id)) FROM image_objects WHERE id LIKE chr(48)||chr(48)||chr(48)||chr(37) AND created_at <= (SELECT created_at FROM initial_data_boundary);'
local_sum=$(ssh "$host" "PGPASSWORD=isureal psql -X -h 127.0.0.1 -p 5433 -U isureal -d isureal -At -c \"$check_sql\"")
remote_sum=$(ssh "$db_host" "PGPASSWORD=isureal psql -X -h 127.0.0.1 -p 5432 -U isureal -d isureal -At -c \"$check_sql\"")
echo "local:  $local_sum"
echo "remote: $remote_sum"
if [ "$local_sum" != "$remote_sum" ]; then
	echo "手元のデータが asia2 の初期データと一致しない" >&2
	exit 1
fi

ssh "$host" 'mkdir -p /home/isucon/images/init /home/isucon/images/new && nice -n 19 ionice -c2 -n7 /tmp/image-export'

ssh "$host" 'sudo -u postgres /usr/lib/postgresql/16/bin/pg_ctl -D /var/lib/postgresql/16/main -w stop; systemctl is-enabled postgresql; systemctl is-active postgresql || true; find /home/isucon/images/init -type f | wc -l; du -sh /home/isucon/images/init'
