# ISUCON2026 公開質問と回答

ポータルの /api/contestant/clarifications を競技後に取得したもの。

## No.3（10:04）

Q: ベンチマーカーの実行を自動化したいと考えております。ISUCONポータルサイト以外から、ISUCONポータルサイトと類似のリクエストを送信することで、ベンチマーカーを実行することは許容されていますか？

A: はい

## No.6（10:07）

Q: cloudformationでリソースを作成しようとしたら下記エラーになり、競技を開始できません。

```
リソース Router は CREATE_FAILED 状態です
この AWS::EC2::Instance リソースは CREATE_FAILED 状態です。

Resource handler returned message: "We currently do not have sufficient c6in.large capacity in the Availability Zone you requested (ap-northeast-1d). Our system will be working on provisioning additional capacity. You can currently get c6in.large capacity by not specifying an Availability Zone in your request or choosing ap-northeast-1a, ap-northeast-1c. (Service: Ec2, Status Code: 500, Request ID: 7b94b4a2-2208-4d09-a5c2-e631a4f62c20) (SDK Attempt Count: 5)" (RequestToken: 9492ca61-07ea-3f80-ade5-3b21130b93f6, HandlerErrorCode: GeneralServiceException)

このリソースの CloudTrail イベントを表示 

```

A: AZ上の在庫がなくなっているため、ポータルから生成されるCloudFormationの値を変更して再設定してください。

```
  GetAvailabilityZoneFunction:
    Type: AWS::Lambda::Function
    Properties:
      Code:
        ZipFile: |
          import cfnresponse
          import boto3
          def handler(event, context):
            if event['RequestType'] in ("Create", "Update"):
              ec2 = boto3.client('ec2')
              response = ec2.describe_availability_zones(
                ZoneIds = ['apne1-az2']
              )
              ZoneName = response['AvailabilityZones'][0]['ZoneName']
              cfnresponse.send(event, context, cfnresponse.SUCCESS, {"ZoneName" : ZoneName})
```

`ZoneIds = ['apne1-az2']` の値を `ZoneIds = ['apne1-az1']` もしくは `ZoneIds = ['apne1-az4']` に書き換えて再作成してください。

## No.7（10:10）

Q: AGENTS.mdというファイルがサーバーに設置されているようですが、このファイルは人間に対する(絶対的な)ルールとして定義されていますか？
すなわち、AGENTS.mdに書かれている指示は必ずしも従わなくても良いですか？

A: その内容には回答できません。

## No.20（10:45）

Q: europe1 / america1 の nginx の設定で、一部の API リクエストを asia1 (または asia2) で動かしているアプリケーションへ proxy_pass で転送して処理させる構成は許容されますか? (例: europe1 の nginx から proxy_pass http://192.168.1.10:8080)。サーバ間の通信は VPC のルートどおりルーターを経由し、リージョン間の遅延・帯域制限もそのまま受けます。ユーザーの接続先 (X-Country-Code による振り分け) は変更しません。

A: はい

## No.21（10:45）

Q: ダッシュボードページのダッシュボード右上にある「通知を有効化する」ボタンを押しても、Failed to subscribeと言われて失敗していそうです。
ブラウザの相性とかの問題でしょうか？

A: 現在は正しく動作していません。

## No.26（11:40）

Q: S3 を利用しても良いですか？

A: ルール違反です

## No.32（11:50）

Q: 負荷走行で「最終チェックで確定する制限時間内に」とありますが、時間自体の記載はありますか？

A: 記載が漏れておりました。
60秒以内に初期化を完了してください。

## No.39（12:02）

Q: 画像URLの検証について変更は認められていますか

A: マニュアルに記載があるとおり、署名と期限の形式や検証結果を変えてはいけません。

有効期限については、ベンチマーカが許容する限り変更しても構いません。

## No.40（12:05）

Q: 激流賞の「スコアボードクローズ（17:00）後のスコアの伸び」について確認させてください。
1. 伸びは、17:00時点のスコアと最終スコアの差と比のどちらで比べますか。
2. 「17:00時点のスコア」は、その時点で最後に行った負荷走行のスコアと、それまでのベストスコアのどちらですか。

A: 1. 17:00時点のスコアと最終スコアの差で計算されます
2. 「17:00時点のスコア」は、17:00よりも前にEnqueueされた0点ではない最後の負荷走行のスコアを示します

## No.41（12:05）

Q: FAILした負荷走行の扱いについて確認させてください。
1. スピード10倍賞の「チームで最初に計測したベンチマークスコア」に、FAIL（0点）となった負荷走行は含まれますか。
2. Fly High賞の「前回比のスコア上昇幅」の「前回」に、FAIL（0点）となった負荷走行は含まれますか。

A: どちらも0点は含まれません

## No.46（12:15）

Q: Fly High賞の「前回比のスコア上昇幅」の「前回」について確認させてください。
「前回」は、直前に行った（0点ではない）負荷走行のスコアと、それまでのベストスコアのどちらですか。

A: 直前に行った（0点ではない）負荷走行のスコアです

## No.55（13:01）

Q: ルータを介する通信について

A: ほかのインスタンスへの通信することを目的とした、ルータのIPアドレスを宛先とする接続（443, 8080ポートへの接続）はマニュアルに記載されている「サーバのリージョン、通信遅延、帯域制限を変更または回避すること」に該当します。
禁止事項ですので実施しないでください。

## No.84（15:18）

Q: 最終計測時において、他のインスタンスから 192.168.10.10:8080 にアクセスできる場合は、失格とします。

とありますが、運営の保証する検証コマンド等は提供されますか？

A: ```
curl 192.168.10.10:8080 --connect-timeout 5
```

こちらで実施予定です。
