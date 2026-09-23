# 競技サーバの代わり。sshd + sudo + flock + 最小の systemctl を持つ Ubuntu。
#
# 本物の systemd は動かさない（コンテナで動かすと重く、壊れやすい）。
# 代わりに /usr/local/bin/systemctl が start/stop/restart/is-active を模す。
# deploy と lock が検証したいのは「ssh 越しに配って入れ替えて再起動できるか」なので、
# systemd の再現度そのものは要らない。
FROM ubuntu:24.04

# nginx / mysql-server / perl は setup の logs・tools ステップの検証用（本物の設定を書いて -t と
# 再起動が通るか、alp / pt-query-digest が本物のログを読めるかを見る）。
RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
      openssh-server openssh-client sudo util-linux procps curl ca-certificates \
      nginx mysql-server perl \
 && rm -rf /var/lib/apt/lists/* \
 && install -d -o mysql -g mysql /var/run/mysqld

RUN useradd -m -s /bin/bash isucon \
 && printf 'isucon ALL=(ALL) NOPASSWD:ALL\n' > /etc/sudoers.d/90-isucon \
 && chmod 0440 /etc/sudoers.d/90-isucon \
 && install -d -m 700 -o isucon -g isucon /home/isucon/.ssh \
 && install -d -m 755 -o isucon -g isucon /home/isucon/webapp \
 && install -d -m 755 -o isucon -g isucon /home/isucon/webapp/go \
 && install -d /run/sshd /etc/isuh-units /var/run/isuh /var/log/isuh

COPY systemctl /usr/local/bin/systemctl
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod 0755 /usr/local/bin/systemctl /usr/local/bin/entrypoint.sh

# 本物の systemd unit と env ファイル（競技環境の形）。setup conf が取り込む対象。
# 起動には使わない（起動は下の /etc/isuh-units）。
RUN install -d /etc/systemd/system \
 && printf '[Unit]\nDescription=isuapp\n\n[Service]\nWorkingDirectory=/home/isucon/webapp/go\nEnvironmentFile=/home/isucon/env.sh\nExecStart=/home/isucon/webapp/go/isuapp\n\n[Install]\nWantedBy=multi-user.target\n' \
      > /etc/systemd/system/isuapp.go.service \
 && printf 'ISUCON_DB_HOST=isu3\nISUCON_DB_USER=isucon\n' > /home/isucon/env.sh \
 && chown isucon:isucon /home/isucon/env.sh

# ユニット定義 = 起動コマンドを 1 行書いたファイル。先頭の as-root は root のまま動かす印
# （nginx は 80 番を、mysqld は自分で mysql ユーザへ切り替えるため）。
RUN printf '/home/isucon/webapp/go/isuapp\n' > /etc/isuh-units/isuapp.go.service \
 && printf "as-root nginx -g 'daemon off;'\n" > /etc/isuh-units/nginx.service \
 && printf 'as-root mysqld\n' > /etc/isuh-units/mysql.service

EXPOSE 22
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
