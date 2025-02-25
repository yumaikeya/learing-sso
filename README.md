# Planner Backend V2
API GW + Lambda(image: distroless) + DynamoDB

## Presentation Layer
API GW requestをparseし、commandに変換する

## Application Layer
`internal/{service}/application/usecase`が担当しています。
IFで抽象化しています。

## Domain Layer
`internal/{service}/domain/model`に定義しています。
modelが持つfieldを書いています。

## Infrastructure Layer
`internal/{service}/infrastructure/repo`でDynamoDBとの接続を行っています。
Dynamo特有のロジックはここに書くようにしています。
DBとのやり取りはdbModelを介して行います。

## How to Test?

1. `make up`
2. `docker compose exec api_private sh`
3. `go run pkg/scripts/migrate.go` => Migrate your local dynamo to create tables
4. Open Postman to call `http://localhost:9090/2015-03-31/functions/function/invocations` (Lambda-RIE endpoint)