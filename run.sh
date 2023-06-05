#!/bin/bash
set -ex
./build.sh
pushd build >/dev/null
go run main.go $*
popd >/dev/null
