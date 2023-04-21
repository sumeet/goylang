#!/bin/bash
set -ex
./run.sh
go run build/main.go goylang.goy > /tmp/goylangtest.go
go run /tmp/goylangtest.go goylang.goy > /tmp/goylang_output.go
gofmt -w /tmp/goylang_output.go
diff -u ./current_compiler/main.go /tmp/goylang_output.go
