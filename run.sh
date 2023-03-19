#!/bin/bash
set -ex
mkdir -p build
go run *.go > build/main.go
gofmt -w build/main.go
cat build/main.go
go run build/main.go