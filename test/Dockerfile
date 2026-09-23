# 競技サーバの代わり。sshd + sudo + flock + 最小の systemctl を持つ Ubuntu。
#
# 本物の systemd は動かさない（コンテナで動かすと重く、壊れやすい）。
# 代わりに /usr/local/bin/systemctl が start/stop/restart/is-active を模す。
# deploy と lock が検証したいのは「ssh 越しに配って入れ替えて再起動できるか」なので、
# systemd の再現度そのものは要らない。
FROM ubuntu:24.04

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
      openssh-server sudo util-linux procps curl ca-certificates \
 && rm -rf /var/lib/apt/lists/*

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

# ユニット定義 = 起動コマンドを 1 行書いたファイル。
RUN printf '/home/isucon/webapp/go/isuapp\n' > /etc/isuh-units/isuapp.go.service \
 && printf 'sleep infinity\n' > /etc/isuh-units/nginx.service \
 && printf 'sleep infinity\n' > /etc/isuh-units/mysql.service

EXPOSE 22
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
