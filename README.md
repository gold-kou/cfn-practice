# cfn-practice
CloudFormationを学習するにあたり実装したテンプレートとAPIサーバです。

backend/ にGoで簡易的なAPIサーバを実装しました。MySQLのmessagesテーブルに文字列をPOST/GET/DELETEする機能があります。

cfn/ にCloudFormationのテンプレートを実装しました。サーバはECS（Fargate）です。
