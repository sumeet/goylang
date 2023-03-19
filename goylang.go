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
	IntType Type = iota // i64
)

func guessType(expr Expr) Type {
	switch expr.NodeType() {
	case IntLiteralExprNodeType:
		return IntType
	default:
		panic("unsupported expression type")
	}
}

func sizeOf(t Type) int {
	switch t {
	case IntType:
		return 8
	default:
		panic("unsupported type")
	}
}

func compileFunction(function Function) string {
	var buf bytes.Buffer
	// TODO: strike here
	buf.WriteString(fmt.Sprintf("%s:\n", function.Name))

	stackVars := findStackVars(function)
	buf.WriteString(fmt.Sprintf("enter %d, 0\n", stackVars.StackSize()))
	buf.WriteString(compileBlock(function.Body, stackVars))
	buf.WriteString("leave\n")

	// TODO: implement return, every function just always returns 0 now
	buf.WriteString("mov rax, 0\n")
	buf.WriteString("ret\n")
	return buf.String()
}

func compileBlock(body Block, stackVars stackVars) string {
	var buf bytes.Buffer
	for _, stmt := range body.statements {
		buf.WriteString(compileStatement(stmt, stackVars))
	}
	return buf.String()
}

func compileStatement(stmt Statement, stackVars stackVars) string {
	switch stmt.NodeType() {
	case AssignmentStmtNodeType:
		return compileAssignmentStmt(stmt.(AssignmentStmt), stackVars)
	default:
		panic("unsupported node type")
	}
}

func compileAssignmentStmt(stmt AssignmentStmt, vars stackVars) string {
	return "TODO: unimplemented\n"
}

type stackVars struct {
	typeByName   map[string]Type
	offsetByName map[string]int
}

func newStackVars(typeByName map[string]Type) stackVars {

	return stackVars{
		typeByName: typeByName,
	}
}

func (sv stackVars) StackSize() int {
	size := 0
	for _, t := range sv.typeByName {
		size += sizeOf(t)
	}
	return size
}

func findStackVars(function Function) stackVars {
	variables := make(map[string]Type)
	Walk(function, func(node Node) {
		if node.NodeType() == AssignmentStmtNodeType {
			child := node.(AssignmentStmt)
			variables[child.VarName] = guessType(child.Expr)
		}
	})
	return newStackVars(variables)
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
