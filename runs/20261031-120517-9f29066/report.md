# run 20261031-120517-9f29066

## Summary

| | |
|---|---|
| score | — (error) |
| status | error |
| ref | relation-pair-locks |
| sha | 9f29066a25e1 |
| 所要時間 | 15s |
| 前回比 | -56314 (-100.0%) vs 20261031-115822-a8d9e7d |
| ベスト比 | -60201 (-100.0%) vs 20261031-115040-1c73ee4 = 60201 |
| note | relation pair locks |

**失敗の理由**

```
デプロイに失敗しました (3/4 台)
  isu1: Command failed with exit code 255: scp -o 'ConnectTimeout=5' -o 'ControlMaster=auto' -o 'ControlPath=/Users/tmokmss/.ssh/isuh-%C' -o 'ControlPersist=300' -o 'ServerAliveInterval=15' -o 'ServerAliveCountMax=3' -q /Users/tmokmss/projects/isucon-harness/.claude/worktrees/agent-a3be7d35e58d79104/.isuh/build/9f29066a25e1-e67f6683/.isuh-conf-isu1.tar 'isu1:/tmp/isuh-conf-28434-isu1.tar'

scp: Connection closed
  isu2: Command failed with exit code 255: scp -o 'ConnectTimeout=5' -o 'ControlMaster=auto' -o 'ControlPath=/Users/tmokmss/.ssh/isuh-%C' -o 'ControlPersist=300' -o 'ServerAliveInterval=15' -o 'ServerAliveCountMax=3' -q /Users/tmokmss/projects/isucon-harness/.claude/worktrees/agent-a3be7d35e58d79104/.isuh/build/9f29066a25e1-e67f6683/.isuh-conf-isu2.tar 'isu2:/tmp/isuh-conf-28434-isu2.tar'

scp: Connection closed
  isu4: isu4: isureal-go.service が起動していません (is-active=unknown)

```

## Bench

```
未取得: ベンチを起動できなかった。
```

## Access log (alp)

```
未取得
```

## Slow query (pt-query-digest)

```
未取得
```

## CPU profile (pprof)

```
未取得
```

## Resources

未取得: リソース計測の結果がありません。

## Diff from previous run

スコア: 56314 → — (error)

`a8d9e7d..9f29066`

```
 conf/isu1/home/isucon/env.sh      |   3 +-
 webapp/go/account_deletion_api.go |   2 +-
 webapp/go/admin_user_api.go       |   2 +-
 webapp/go/comment_api.go          |   2 +-
 webapp/go/offload.go              |  65 +++++++++++++++++++++----
 webapp/go/post_api.go             |   4 +-
 webapp/go/profile_image_api.go    |   3 +-
 webapp/go/realmoji_api.go         |   6 +--
 webapp/go/relationship.go         |  10 +++-
 webapp/go/store.go                | 100 ++++++++++++++++++++++++++++++++------
 10 files changed, 157 insertions(+), 40 deletions(-)
```
