#! /bin/bash

GOOS=windows GOARCH=amd64 go build -o build/htmltox.exe main.go

GOOS=linux GOARCH=amd64 go build -o build/htmltox-linux-amd64 main.go
GOOS=linux GOARCH=arm64 go build -o build/htmltox-linux-arm64 main.go

GOOS=darwin GOARCH=amd64 go build -o build/htmltox-darwin-amd64 main.go
GOOS=darwin GOARCH=arm64 go build -o build/htmltox-darwin-arm64 main.go
