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

func collectStringLiterals(program Program) map[uint32]string {
	stringLiterals := make(map[uint32]string)
	Walk(program, func(node Node) {
		if node.NodeType() == StringLiteralExprNodeType {
			child := node.(StringLiteralExpr)
			stringLiterals[hash(child.Value)] = child.Value
		}
	})
	return stringLiterals
}

func Compile(program Program) string {
	var buf bytes.Buffer

	// data section
	buf.WriteString("section .data\n")
	for hash, value := range collectStringLiterals(program) {
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

	for _, function := range program.functions {
		buf.WriteString(compileFunction(function))
	}

	return buf.String()
}

type Type uint8

const (
	IntType Type = iota
)

func guessType(expr Expr) Type {
	switch expr.NodeType() {
	case IntLiteralExprNodeType:
		return IntType
	default:
		panic("unsupported expression type")
	}
}

func compileFunction(function Function) string {
	var buf bytes.Buffer
	stackVars := make(map[string]Type)
	buf.WriteString(fmt.Sprintf("%s:\n", function.Name))
	buf.WriteString("ret\n")
	return buf.String()
}

func findVariables(program Program) map[string]Type {
	variables := make(map[string]Type)
	Walk(program, func(node Node) {
		if node.NodeType() == VariableDeclarationNodeType {
			child := node.(VariableDeclaration)
			variables[child.Name] = guessType(child.Expr)
		}
	})
	return variables
}

func main() {
	dat, err := os.ReadFile("./hello.goy")
	if err != nil {
		panic(err)
	}
	tokens := lex(dat)
	program := parse(tokens)
	fmt.Println(Compile(program))
}
