# SecretsManagerExtension
A local HTTP server that runs at loopback address and proxies retrieval of secrets from AWS Secrets Manager, for improved performance and added security.

To retrieve a secret, use ```http://localhost:8080/secrets/<your-secret-id>```. ```<your-secret-id>``` can contain special character ```/``` to form a logical hierachy of secrets with folders, for example, ```AppSecretsPlayground/DB/Admin```, however, it shall not start with special character ```/```. In other words, secret name does not start with ```/```. 


## Commands

```
git clone https://github.com/JinlianWang/SecretsManagerExtension.git
go mod init SecretsManagerExtension
go mod tidy
go run mainAgent.go
http://localhost:8080/secrets/AppSecretsPlayground/DB/Admin
```
