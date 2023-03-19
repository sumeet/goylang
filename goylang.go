package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("./hello.goy")
	if err != nil {
		panic(err)
	}
	tokens := lex(dat)
	program := parse(tokens)
	fmt.Printf("%#v\n", program)
}
