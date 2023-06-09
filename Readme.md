[![Andromeda](https://github.com/motleylabs/andromeda/actions/workflows/prod.yml/badge.svg)](https://github.com/motleylabs/andromeda/actions/workflows/prod.yml)
[![Andromeda Dev](https://github.com/motleylabs/andromeda/actions/workflows/dev.yml/badge.svg)](https://github.com/motleylabs/andromeda/actions/workflows/dev.yml)

# Andromeda

An abstracted data layer for getting marketplace information on the Solana blockchain.

## Contributing

We welcome contributions to Night Market from the community -- please open a pull request!

Feel free to join the [Motley DAO Discord](https://discord.gg/motleydao) to talk to the team and other community members.

All contributions are automatically licensed under the [AGPL 3.0](https://github.com/motleylabs/andromeda/blob/main/license.md).

## Overview

This service is built using Golang.

The server provides several routes for collections, NFTs, users, general blockchain information, etc.

## How to run

First, you need to install the golang in order to run this project. Then you can execute the following command to run the server:
```
go run app/main.go
```

The swagger API documentation can help you learn how to use the APIs. You can visit the following page to see the full documentation.
```
http://localhost:8080/swagger/index.html
```

Swagger documentation can be generated by running the following command on the root directory.
```
swag init -g ./internal/api/routers/api.go
```

## Environment

The list of required environment variables for the application. To be set through the OS environment or through .env file.

| Name | Description |
|------|-------------|
| PORT | The port to listen on. E.g. `5555` |
| PROVIDER | The blockchain data provider. Currently only supports [Hyperspace](https://docs.hyperspace.xyz/hype/developer-guide/overview). |
| API_KEY | An API key for the data provider. |
| RPC_ENDPOINT | A Solana RPC endpoint. E.g. `https://api.mainnet-beta.solana.com/` |


## How to test

You can test the current API server by using the following command:
```
go clean && go test -v ./...
```

## How to build

You can build this project to run the following command:
```
go build -o andromeda ./app/main.go

```
