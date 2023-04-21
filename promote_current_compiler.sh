#!/bin/bash
set -ex
./run.sh
cp build/main.go current_compiler/main.go
