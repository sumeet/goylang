#!/bin/bash
set -ex
mkdir -p build
go run ./*.go ./lexer.goy > build/main.go
gofmt -w build/main.go
cat build/main.go
go run build/main.go