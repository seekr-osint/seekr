#!/bin/sh

go generate ./...
tsc --project web
go run main.go
