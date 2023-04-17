# SecretsManagerExtension
A local HTTP server that runs at loopback address and proxies retrieval of secrets from AWS Secrets Manager, for improved performance and added security.

To retrieve a secret, use ```http://localhost:8080/secrets/<your-secret-id>```. ```<your-secret-id>``` can contain special character ```/``` to form a logical hierachy of secrets with folders, for example, ```AppSecretsPlayground/DB/Admin```, however, it shall not start with special character ```/```. In other words, secret name does not start with ```/```. 


## Commands

### Commands for Local Agent

```
git clone https://github.com/JinlianWang/SecretsManagerExtension.git
go mod init SecretsManagerExtension
go mod tidy
go run mainAgent.go
http://localhost:8080/secrets/AppSecretsPlayground/DB/Admin
```

### Commands to Create IAM Role

```
aws iam create-role --role-name lambda-cloudwatch-logs-role --assume-role-policy-document file://trust-policy.json
aws iam create-policy --policy-name LambdaCloudWatchLogsPolicy --policy-document file://logging-policy.json
aws iam attach-role-policy --role-name lambda-cloudwatch-logs-role --policy-arn arn:aws:iam::975156237701:policy/LambdaCloudWatchLogsPolicy
```

### Commands to Create Lambda Extension

```
GOOS=linux GOARCH=amd64 go build -o extension main.go
zip extension.zip extension
aws lambda publish-layer-version --layer-name secretsmanager-extension --zip-file fileb://extension.zip --compatible-runtimes provided
```

### Commands to Create Lambda Function

```
zip function.zip index.js
aws lambda create-function --function-name secretsmanager-demo-lambda --runtime nodejs14.x --role arn:aws:iam::975156237701:role/lambda-cloudwatch-logs-role --handler index.handler --zip-file fileb://function.zip --layers arn:aws:lambda:us-east-1:975156237701:layer:secretsmanager-extension:1
aws lambda invoke --function-name secretsmanager-demo-lambda --cli-binary-format raw-in-base64-out --payload '{}' /dev/stdout
```