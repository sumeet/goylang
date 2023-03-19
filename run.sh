#!/bin/bash
set -ex
mkdir -p build
go run *.go > build/prog.asm
nasm -g -felf64 build/prog.asm -l build/prog.lst
gcc -no-pie -m64 -o build/prog build/prog.o
bat --pager=never ./build/prog.asm
./build/prog
