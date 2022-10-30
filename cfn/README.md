# 構築手順
## SSMパラメータストアの設定
- cfn-practice-db-host
  - RDSエンドポイントの値
- cfn-practice-db-password
  - 任意のDBパスワードの値

## Apply templates
Apply all cfn/ templates.

## DB Initial Settings
adminユーザのパスワード変更、テーブルマイグレーション、バックエンド用ユーザの作成を実施します。

1. EC2コンソールの EC2 Instance Connect を利用し、踏み台サーバのインスタンスに接続する。
2. `mysql -u admin -p -h <RDSエンドポイント>` を実行する。RDSエンドポイントはコンソールの `接続とセキュリティ` から確認可能。初期パスワードはrds.ymlのMasterUserPasswordを参照する。
3. `SET PASSWORD = 'XXXXX';` を実行してadminユーザのパスワードを変更する。パスワード値は任意の値。パスワードはどこかに保存しておくこと。
4. `USE practicedb;` を実行する。
5. `backend/sql/mysql/entrypoint/001_create_tables.sql` の内容を実行し、テーブルマイグレーションする。
6. `backend/sql/mysql/create_user.sql` の内容を実行し、バックエンド用のDBユーザを作成する。IDENTIFIED BYに設定するパスワード値はパラメータストアに設定したcfn-practice-db-passwordの値。
7. exitする。

# API Request
## DNS名
AWSコンソールより、ALBのDNS名をコピーする。

## Post
```
curl --location --request POST 'http://<ALBのDNS名>/messages' \
--header 'Content-Type: application/json' \
--data-raw '{
  "message": "Hello world."
}'
```

## Get
```
curl --location --request GET 'http://<ALBのDNS名>/messages/1'
```

## Delete
```
curl --location --request DELETE 'http://<ALBのDNS名>/messages/1'
```
