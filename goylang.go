package main

import (
	"fmt"
	"hash/fnv"
	"os"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	if _, err := h.Write([]byte(s)); err != nil {
		panic(err)
	}
	return h.Sum32()
}

type EndProgram struct {
	stringLiterals map[uint32]byte
}

func main() {
	dat, err := os.ReadFile("./hello.goy")
	if err != nil {
		panic(err)
	}
	tokens := lex(dat)
	program := parse(tokens)
	//fmt.Printf("%#v\n", program)
	fmt.Printf("%#v\n", program.Children())
}
