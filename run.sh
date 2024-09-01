#! /bin/bash

main="$PWD/cmd/server/"
tmp="$GOPATH/bin/templ"

$tmp generate

go run $main/main.go
