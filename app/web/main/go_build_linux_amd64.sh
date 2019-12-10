#!/bin/sh

target=Webapp
go build -o $target ./Webapp_test_main.go
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $target ./Webapp_test_main.go

