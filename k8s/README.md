# 構築手順
## 前提
以下がインストール済み。

- AWS CLI
- eksctl
- kubectl

cfn/の以下が適用済み。

- 001_iam.yml
- 002_ecr.yml
- 003_network.yml
- 004_bastion.yml
- 005_rds.yml
  - Table Migration included

## EKS Clusterの構築
```
eksctl create cluster \
--vpc-public-subnets <PublicSubnet1,PublicSubnet2> \
--name eks-work-cluster \
--region ap-northeast-1 \
--version 1.23 \
--nodegroup-name eks-work-nodegroup \
--node-type t3.small \
--nodes 2 \
--nodes-min 2 \
--nodes-max 5
```

eksctlの内部でCloudFormationが実行しているため、CloudFormationのコンソールから進捗を確認できる。

## Namespaceとコンテキストの作成
NameSpaceの作成をする。

```
kubectl apply -f 001_namespace.yml

kubectl get namespace
```

コンテキストの作成とNamespaceの紐付けをし、それを使用をします。

```
kubectl config get-contexts

kubectl config set-context eks-work \
--cluster eks-work-cluster.ap-northeast-1.eksctl.io \
--user <AUTHINFO列の値>\
--namespace eks-work

kubectl config use-context eks-work

kubectl config get-contexts
```

## Secretの作成
```
DB_HOST=<パラメータストアのcfn-practice-db-host> \
DB_PASSWORD=<パラメータストアのcfn-practice-db-password> \
envsubst < 002_secret.yml.template | \
kubectl apply -f -

kubectl get secret
```

## Deploymentの作成
```
ECR_HOST=030816431860.dkr.ecr.ap-northeast-1.amazonaws.com \
envsubst < 003_deployment.yml.template | \
kubectl apply -f -

kubectl get deployment
```

## Serviceの作成
```
kubectl apply -f 004_service.yml

kubectl get servic
```

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

# 削除手順
## ServiceとDeploymentの削除
```
kubectl delete service backend-app-service
kubectl delete deployment backend-app
kubectl get all
```

## EKSクラスターの削除
SecretやConfigMapも同時に削除されます。

```
eksctl delete cluster --name eks-work-cluster
```
