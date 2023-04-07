# Andromeda

A data layer for getting marketplace information on Solana blockchain.

## Overview

The HTTP server was built using Golang 1.20.
The server provides several routes for collections, NFTs, and users.

## How to run

First, you need to install the golang in order to run this project. Then you can execute the following command to run the server:
```
go run app/main.go
```

The swagger API documentation can help you learn how to use the APIs. You can visit the following page to see the full documentation.
```
http://localhost:8080/swagger/index.html
```

## How to test

You can test the current API server by using the following command:
```
go clean -testcase && go test -v ./...
```

## How to build

You can build this project to run the following command:
```
go build -o andromeda ./app/main.go
```
