# About
This is a sample backend application. No tests code.

# Push docker image
## IAM User
cfn/security.ymlを適用しIAMユーザとグループを作成する。

~/.aws/credential を編集し、作成したユーザのprofileを設定する。

## Build and push image
```
aws ecr get-login-password --profile admin-user-XXX --region ap-northeast-1 | docker login --username AWS --password-stdin XXX.dkr.ecr.ap-northeast-1.amazonaws.com

docker build -t cfn-practice-repository .

docker tag cfn-practice-repository:latest XXX.dkr.ecr.ap-northeast-1.amazonaws.com/cfn-practice-repository:latest

docker push XXX.dkr.ecr.ap-northeast-1.amazonaws.com/cfn-practice-repository:latest
```

# Local check
## Launch
```
docker compose up
```

## Request
### Post
```
curl --location --request POST 'http://localhost:80/messages' \
--header 'Content-Type: application/json' \
--data-raw '{
  "message": "Hello world."
}'
```

### Get
```
curl --location --request GET 'http://localhost:80/messages/1'
```

### Delete
```
curl --location --request DELETE 'http://localhost:80/messages/1'
```
