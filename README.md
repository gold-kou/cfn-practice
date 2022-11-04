# cfn-practice
CloudFormationを学習するにあたり実装したテンプレートとAPIサーバです。

backend/ にGoで簡易的なAPIサーバを実装しました。

cfn/ にCloudFormationのテンプレートを実装しました。サーバにECS（Fargate）を利用する場合は全てのテンプレートを実行します。EKSを利用する場合は005_rds.ymlまで実行し、以降はk8s/を利用します。

k8s/ にKubernetesのテンプレートを実装しました。
