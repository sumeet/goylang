#!/bin/bash
set -ex
mkdir -p build
go run current_compiler/main.go ./lexer.goy > build/main.go
gofmt -w build/main.go

