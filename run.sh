#!/bin/bash
set -ex
./build.sh
pushd build
go run main.go $*
popd
