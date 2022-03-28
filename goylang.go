package main

import (
	"bytes"
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

type IR struct {
	stringLiterals map[uint32]string
}

func GenerateIR(program Program) IR {
	stringLiterals := make(map[uint32]string)
	Walk(program, func(node Node) {
		if node.NodeType() == StringLiteralExprNodeType {
			child := node.(StringLiteralExpr)
			stringLiterals[hash(child.Value)] = child.Value
		}
	})
	return IR{stringLiterals: stringLiterals}
}

func Compile(ir IR) string {
	var buf bytes.Buffer

	// data section
	buf.WriteString("section .data\n")
	for hash, value := range ir.stringLiterals {
		var csBytes bytes.Buffer
		for i, b := range []byte(value) {
			csBytes.WriteString(fmt.Sprintf("0x%02x", b))
			if i < len(value)-1 {
				csBytes.WriteString(",")
			} else {
				// TODO: should strings be null terminated?
				//       they are for now for compatibility with C
				csBytes.WriteString(",0")
			}
		}
		buf.WriteString(fmt.Sprintf("string_%d: db %s\n", hash, csBytes.String()))
	}

	// text section
	buf.WriteString("section .text\n")
	buf.WriteString("global main\n")
	buf.WriteString("extern printf\n")

	return buf.String()
}

func main() {
	dat, err := os.ReadFile("./hello.goy")
	if err != nil {
		panic(err)
	}
	tokens := lex(dat)
	program := parse(tokens)
	endProgram := GenerateIR(program)
	fmt.Println(Compile(endProgram))
}
