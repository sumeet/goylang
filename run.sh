#!/bin/bash
set -ex
mkdir -p build
mkdir -p build2
mkdir -p build3
mkdir -p build4
mkdir -p build5

# original compiler
go run ./*.go ./lexer.goy > build/main.go
gofmt -w build/main.go

# compiled one time compiler
go run build/main.go > build2/main.go
gofmt -w build2/main.go

# compiled two times compiler
go run build2/main.go > build3/main.go
gofmt -w build3/main.go
go run build3/main.go

# compiled three times compiler
go run build3/main.go > build4/main.go
gofmt -w build4/main.go
go run build4/main.go

# compiled four times compiler
go run build4/main.go > build5/main.go
gofmt -w build5/main.go
go run build5/main.go
