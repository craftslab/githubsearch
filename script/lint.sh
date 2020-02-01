#!/bin/bash

list="cmd,search"

go env -w GOPROXY=https://goproxy.cn,direct

old=$IFS IFS=$','
for item in $list; do
  gofmt -s -w $item/*.go
  golangci-lint run $item/*.go
done
IFS=$old

go mod tidy
