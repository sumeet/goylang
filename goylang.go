package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type X interface{}
type A struct{}
type B struct{}

func main() {
	a := A{}
	b := B{}
	fn := func(x X) {
		switch x.(type) {
		case A:
			fmt.Println("A")
		case B:
			fmt.Println("B")
		}
	}
	fn(b)
	fn(a)
}

func main2() {
	dat, err := os.ReadFile("./hello.goy")
	if err != nil {
		panic(err)
	}
	tokens := lex(dat)
	module := parse(tokens)
	s := Compile(module)
	fmt.Println(s)
}

func Compile(module Module) string {
	var b strings.Builder
	b.WriteString("package main\n\n")
	for _, statement := range module.Statements {
		compileStatement(&b, statement)
		b.WriteByte('\n')
	}
	return b.String()
}

func blah(goSource string) {
	outFile := "./output/program.go"
	err := ioutil.WriteFile(outFile, []byte(goSource), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %s", err.Error())
		os.Exit(1)
	}

	cmd := exec.Command("go", "run", outFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// is there output?
		if len(output) > 0 {
			fmt.Println(string(output))
		}

		fmt.Printf("Error compiling Go source: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func compileStatement(b *strings.Builder, s Statement) {
	switch s.NodeType() {
	case ModuleNodeType:
		panic("module node type should not be compiled through this function")
	case BlockNodeType:
		compileBlock(b, s.(Block))
	case AssignmentStmtNodeType:
		compileAssignmentStmt(b, s.(AssignmentStmt))
	case ReassignmentStmtNodeType:
		compileReassignmentStmt(b, s.(ReassignmentStmt))
	case StringLiteralExprNodeType:
		compileStringLiteralExpr(b, s.(StringLiteralExpr))
	case FuncCallExprNodeType:
		compileFuncCallExpr(b, s.(FuncCallExpr))
	case IntLiteralExprNodeType:
		compileIntLiteralExpr(b, s.(IntLiteralExpr))
	case VarRefExprNodeType:
		compileVarRefExpr(b, s.(VarRefExpr))
	case FunctionNodeType:
		compileFunction(b, s.(Function))
	default:
		panic(fmt.Sprintf("unknown node type %s", s.NodeType().ToString()))
	}
	return
}

func compileExpr(b *strings.Builder, e Expr) {
	switch e.ExprType() {
	case StringLiteralExprType:
		compileStringLiteralExpr(b, e.(StringLiteralExpr))
	case FuncCallExprType:
		compileFuncCallExpr(b, e.(FuncCallExpr))
	case IntLiteralExprType:
		compileIntLiteralExpr(b, e.(IntLiteralExpr))
	case VarRefExprType:
		compileVarRefExpr(b, e.(VarRefExpr))
	default:
		panic(fmt.Sprintf("unknown expr type %d", e.ExprType()))
	}
}

func compileBlock(b *strings.Builder, block Block) {
	b.WriteString("{\n")
	for _, statement := range block.Statements {
		compileStatement(b, statement)
		b.WriteString("\n")
	}
	b.WriteString("}")
}

func compileAssignmentStmt(b *strings.Builder, stmt AssignmentStmt) {
	b.WriteString(stmt.VarName)
	b.WriteString(" := ")
	compileExpr(b, stmt.Expr)
}

func compileReassignmentStmt(b *strings.Builder, stmt ReassignmentStmt) {
	b.WriteString(stmt.VarName)
	b.WriteString(" = ")
	compileExpr(b, stmt.Expr)
}

func compileStringLiteralExpr(b *strings.Builder, expr StringLiteralExpr) {
	b.WriteString(fmt.Sprintf("%q", expr.Value))
}

func compileFuncCallExpr(b *strings.Builder, expr FuncCallExpr) {
	b.WriteString(expr.FuncName)
	b.WriteString("(")
	for i, arg := range expr.Args {
		if i != 0 {
			b.WriteString(", ")
		}
		compileExpr(b, arg)
	}
	b.WriteString(")")
}

func compileIntLiteralExpr(b *strings.Builder, expr IntLiteralExpr) {
	b.WriteString(fmt.Sprintf("%d", expr.Value))
}

func compileVarRefExpr(b *strings.Builder, expr VarRefExpr) {
	b.WriteString(expr.VarName)
}

func compileFunction(b *strings.Builder, f Function) {
	b.WriteString("func ")
	b.WriteString(f.Name)
	b.WriteString("(")
	//for i, arg := range f.Args {
	//	if i != 0 {
	//		b.WriteString(", ")
	//	}
	//	b.WriteString(arg)
	//}
	b.WriteString(") ")
	compileBlock(b, f.Body)
}
