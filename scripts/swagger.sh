#!/bin/bash
go install github.com/swaggo/swag/cmd/swag@latest

# execute swag command by export PATH
export PATH=$(go env GOPATH)/bin:$PATH
export GO111MODULE=on

rm -f -r ./docs/swagger/docs/*
go mod download
swag init -g cmd/main.go -o ./docs/swagger/docs