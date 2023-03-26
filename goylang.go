package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fname := os.Args[1]
	dat, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	tokens := lex(dat)
	//println(fmt.Sprintf("%#v\n", tokens))
	module := parse(tokens)
	//println(fmt.Sprintf("%#v\n", module))
	s := Compile(module)
	fmt.Println(s)
}

func prelude() string {
	return `package main

import (
	"os"
	"fmt"
)

func readfile(fname string) []byte {
	dat, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return dat
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func add(n, m int) int {
	return n + m
}

func g[T any](slice []T, i int) T {
	return slice[i]
}

func eq[T comparable](x, y T) bool {
	return x == y
}

func ge(x, y int) bool {
	return x >= y
}

func or(bs ...bool) bool {
	for _, b := range bs {
		if b {	
			return true	
		}	
	}
	return false
}

func c(s string) byte {
	return s[0]
}

func nc(bs []byte, i int, s string) bool {
	for j, c := range s {
		if bs[i+j] != byte(c) {
			return false
		}	
	}
	return true
}

func sprintf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func print(args ...interface{}) {
	for _, arg := range args {
		fmt.Printf("%#v\n", arg)
	}
	return
}

func bs(s string) []byte {
	return []byte(s)
}
`
}

func Compile(module Module) string {
	var b strings.Builder
	b.WriteString(prelude())
	b.WriteString("\n\n")
	for _, statement := range module.Statements {
		compileStatement(&b, statement)
		b.WriteByte('\n')
	}
	return b.String()
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
	case EnumNodeType:
		compileEnum(b, s.(Enum))
	case MatchNodeType:
		compileMatch(b, s.(MatchStmt))
	case StructNodeType:
		compileStruct(b, s.(Struct))
	case WhileNodeType:
		compileWhile(b, s.(WhileExpr))
	case BreakNodeType:
		compileBreak(b, s.(BreakExpr))
	case ContinueNodeType:
		compileContinue(b, s.(ContinueExpr))
	case IfNodeType:
		compileIf(b, s.(IfExpr))
	case ReturnNodeType:
		compileReturn(b, s.(ReturnExpr))
	default:
		panic(fmt.Sprintf("don't know how to compile node type %s", s.NodeType().ToString()))
	}
	return
}

func compileReturn(b *strings.Builder, ret ReturnExpr) {
	b.WriteString("return")
	if ret.Expr != nil {
		b.WriteByte(' ')
		compileExpr(b, *ret.Expr)
	}
}

func compileIf(b *strings.Builder, expr IfExpr) {
	b.WriteString("if ")
	compileExpr(b, expr.Cond)
	compileExpr(b, expr.IfBody)
	if expr.ElseBody != nil {
		b.WriteString(" else ")
		compileExpr(b, *expr.ElseBody)
	}
}

func compileBreak(b *strings.Builder, be BreakExpr) {
	b.WriteString("break")
}

func compileContinue(b *strings.Builder, ce ContinueExpr) {
	b.WriteString("continue")
}

func compileWhile(b *strings.Builder, expr WhileExpr) {
	b.WriteString("for ")
	compileExpr(b, expr.Body)
}

func compileStruct(b *strings.Builder, strukt Struct) {
	b.WriteString("type ")
	b.WriteString(strukt.Name)
	b.WriteString(" struct {\n")
	for _, field := range strukt.Fields {
		b.WriteString(field.Name)
		b.WriteString(" ")
		compileType(b, field.Type)
		b.WriteString("\n")

	}
	b.WriteString("}\n")
}

func golangTypeNameWithBindingsThingTODORename(e Expr) (string, *string) {
	switch e.ExprType() {
	case DotAccessExprType:
		var s strings.Builder
		compileInitializerLHS(&s, e)
		return s.String(), nil
	case InitializerExprType:
		i := e.(InitializerExpr)
		v := i.Args[0].(VarRefExpr).VarName
		var s strings.Builder
		compileInitializerLHS(&s, i.Type)
		return s.String(), &v
	}
	panic(fmt.Sprintf("unable to get golang type name for expr %#v", e))
}

func stfuUnusedVars(b *strings.Builder, varName string) {
	b.WriteString("_ = ") // get the golang compiler to shut up about unused variable
	b.WriteString(varName)
	b.WriteRune('\n')
}

const BindingVarname = "binding"
const MatchExprVarname = "matchExpr"

func compileMatch(b *strings.Builder, match MatchStmt) {
	b.WriteString("{\n")
	b.WriteString(MatchExprVarname)
	b.WriteString(" := ")
	compileExpr(b, match.MatchExpr)
	b.WriteString("\n")
	stfuUnusedVars(b, MatchExprVarname)

	for i, matchArm := range match.Arms {
		// right now it's just an enum variant, but could be other stuff that you might want to match in the future
		golangTypeName, binding := golangTypeNameWithBindingsThingTODORename(matchArm.Pattern.Expr)
		if i == 0 {
			b.WriteString(fmt.Sprintf("if %s, ok := matchExpr.(%s); ok {\n", BindingVarname, golangTypeName))
		} else {
			b.WriteString(fmt.Sprintf("} else if %s, ok := matchExpr.(%s); ok {\n", BindingVarname, golangTypeName))
		}
		stfuUnusedVars(b, BindingVarname)

		if binding != nil {
			b.WriteString(*binding)
			b.WriteString(" := ")
			b.WriteString(BindingVarname)
			b.WriteString(".Value")
			b.WriteString("\n")
			stfuUnusedVars(b, *binding)
		}
		compileExpr(b, matchArm.Body)
	}
	// TODO: but like, is an empty match even valid?
	if len(match.Arms) > 0 {
		b.WriteString("}\n") // end of if-chain
	}
	b.WriteString("}\n") // end of anonymous block
}

var CompiledEnums []Enum
var CompiledFuncs map[string]Function = make(map[string]Function)

func findEnumInTable(name string) *Enum {
	for _, e := range CompiledEnums {
		if e.Name == name {
			return &e
		}
	}
	return nil
}

func findFuncInTable(name string) *Function {
	f, ok := CompiledFuncs[name]
	if !ok {
		return nil
	} else {
		return &f
	}
}

func compileEnum(b *strings.Builder, enum Enum) {
	CompiledEnums = append(CompiledEnums, enum)

	// iota constants for Type enum
	compileIotaConstants(b, enum)
	// interface
	compileEnumInterfaces(b, enum)
	// structs
	compileEnumStructs(b, enum)
}

func typeName(e Enum) string {
	return fmt.Sprintf("%sType", e.Name)
}

func enumVariantTag(e Enum, variant Variant) string {
	return fmt.Sprintf("%sType%s", e.Name, variant.Name)
}

func compileIotaConstants(b *strings.Builder, enum Enum) {
	b.WriteString(fmt.Sprintf("type %s uint8\n", typeName(enum)))

	b.WriteString("const (\n")
	for i, variant := range enum.Variants {
		b.WriteString(enumVariantTag(enum, variant))
		if i == 0 {
			b.WriteString(" = iota")
		}
		b.WriteString("\n")
	}
	b.WriteString(")\n")
}

func golangEnumTagMethodName(e Enum) string {
	return fmt.Sprintf("%sTag", e.Name)
}

func golangInterfaceName(e Enum) string {
	return e.Name
}

func compileEnumInterfaces(b *strings.Builder, enum Enum) {
	b.WriteString(fmt.Sprintf("type %s interface {\n", golangInterfaceName(enum)))
	b.WriteString(fmt.Sprintf("%s() %s\n", golangEnumTagMethodName(enum), typeName(enum)))
	b.WriteString("}\n")
}

func compileEnumStructs(b *strings.Builder, enum Enum) {
	for _, variant := range enum.Variants {
		someName := fmt.Sprintf("%s%s", enum.Name, variant.Name)
		b.WriteString(fmt.Sprintf("type %s struct {\n", someName))
		if variant.Type != nil {
			b.WriteString(fmt.Sprintf("Value %s", *variant.Type))
		}
		b.WriteString("}\n")

		b.WriteString(fmt.Sprintf("func (i %s) %s() %s {\n", someName, golangEnumTagMethodName(enum), typeName(enum)))
		b.WriteString(fmt.Sprintf("return %s", enumVariantTag(enum, variant)))
		b.WriteString("}\n")
	}
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
	case DotAccessExprType:
		compileDotAccessExpr(b, e.(DotAccessExpr))
	case InitializerExprType:
		compileInitializerExpr(b, e.(InitializerExpr))
	case BlockExprType:
		compileBlock(b, e.(Block))
	case WhileExprType:
		compileWhile(b, e.(WhileExpr))
	case IfExprType:
		compileIf(b, e.(IfExpr))
	default:
		panic(fmt.Sprintf("unable to compile expr: %#v", e))
	}
}

func getVarName(e Expr) *string {
	switch e.ExprType() {
	case VarRefExprType:
		vn := e.(VarRefExpr).VarName
		return &vn
	default:
		return nil
	}
}

func compileDotAccessExpr(b *strings.Builder, expr DotAccessExpr) {
	vn := getVarName(expr.Left)
	if vn != nil && findEnumInTable(*vn) != nil {
		compileInitializerLHS(b, expr)
	} else {
		compileExpr(b, expr.Left)
		b.WriteString(".")
		b.WriteString(expr.Right)
	}
}

func compileInitializerExpr(b *strings.Builder, expr InitializerExpr) {
	compileInitializerLHS(b, expr.Type)
	b.WriteString("{ ")
	for i, arg := range expr.Args {
		if i > 0 {
			b.WriteString(", ")
		}
		compileExpr(b, arg)
	}
	b.WriteString(" }")
}

// TODO: this feels funny. feels like instead we should just be able to
// guessType(expr) and then Type -> Golang type
func guessGolangTypeName(expr Expr) string {
	switch expr.ExprType() {
	case DotAccessExprType:
		dotAccessExpr := expr.(DotAccessExpr)
		if dotAccessExpr.Left.NodeType() != VarRefExprNodeType {
			break
		}
		return dotAccessExpr.Left.(VarRefExpr).VarName
	case VarRefExprType:
		return expr.(VarRefExpr).VarName
	}
	panic(fmt.Sprintf("couldn't print golang type name for %#v", expr))
}

func compileType(b *strings.Builder, t Type) {
	b.WriteString(t.Name)
}

func compileInitializerLHS(b *strings.Builder, expr Expr) {
	switch expr.ExprType() {
	case DotAccessExprType:
		dotAccessExpr := expr.(DotAccessExpr)
		if dotAccessExpr.Left.NodeType() != VarRefExprNodeType {
			break
		}
		leftNode := dotAccessExpr.Left.(VarRefExpr)
		rightNodeName := dotAccessExpr.Right
		b.WriteString(leftNode.VarName)
		b.WriteString(rightNodeName)
	case TypeExprType:
		compileType(b, expr.(Type))
	case VarRefExprType:
		varRef := expr.(VarRefExpr)
		b.WriteString(varRef.VarName)
	default:
		panic(fmt.Sprintf("unknown initializer LHS type %#v", expr))
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
	// TODO: this doesn't handle multiple assignments
	b.WriteString("var ")
	b.WriteString(stmt.VarNames[0])
	b.WriteString(" ")
	b.WriteString(guessType(stmt.Expr))
	b.WriteString("\n")

	b.WriteString(stmt.VarNames[0])
	b.WriteString(" = ")
	compileExpr(b, stmt.Expr)
	b.WriteString("\n")

	b.WriteString("_ = ")
	b.WriteString(stmt.VarNames[0])
	b.WriteString("\n")
}

func getTypeForFuncCall(funcCall FuncCallExpr) string {
	varRef, ok := funcCall.Expr.(VarRefExpr)
	if !ok {
		panic(fmt.Sprintf("expected var ref expr for func call %#v", funcCall))
	}
	ident := varRef.VarName
	if ident == "readfile" {
		return "[]byte"
	} else if ident == "bs" {
		return "[]byte"
	} else {
		if f := findFuncInTable(ident); f != nil {
			return f.ReturnType.Name
		}
		panic(fmt.Sprintf("don't know type of func call %s", ident))
	}
}

func guessType(expr Expr) string {
	switch expr.ExprType() {
	case StringLiteralExprType:
		return "string"
	case FuncCallExprType:
		funcCall := expr.(FuncCallExpr)
		return getTypeForFuncCall(funcCall)
	case IntLiteralExprType:
		return "int"
	case VarRefExprType:
		panic(fmt.Sprintf("can't guess type for var ref %#v", expr))
	case InitializerExprType:
		init := expr.(InitializerExpr)
		if typ, ok := init.Type.(Type); ok {
			return typ.Name
		} else {
			return guessGolangTypeName(init.Type)
		}
		//case DotAccessExprType:
		//	dotAccessExpr := expr.(DotAccessExpr)
		//	if dotAccessExpr.Left.NodeType() != VarRefExprNodeType {
		//		break
		//	}
		//	leftNode := dotAccessExpr.Left.(VarRefExpr)
		//	rightNodeName := dotAccessExpr.Right
		//	return fmt.Sprintf("%s%s", leftNode.VarName, rightNodeName)

	}
	panic(fmt.Sprintf("can't guess type for expr: %#v", expr))
}

func compileReassignmentStmt(b *strings.Builder, stmt ReassignmentStmt) {
	for i, varName := range stmt.VarNames {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(varName)
	}
	b.WriteString(" = ")
	compileExpr(b, stmt.Expr)
}

func compileStringLiteralExpr(b *strings.Builder, expr StringLiteralExpr) {
	b.WriteString(fmt.Sprintf("%q", expr.Value))
}

func compileFuncCallExpr(b *strings.Builder, funcCall FuncCallExpr) {
	compileExpr(b, funcCall.Expr)
	b.WriteString("(")
	for i, arg := range funcCall.Args {
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
	CompiledFuncs[f.Name] = f

	b.WriteString("func ")
	b.WriteString(f.Name)
	b.WriteString("(")
	for i, param := range f.Params {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(param.Name)
		b.WriteByte(' ')
		compileType(b, param.Type)
	}
	b.WriteString(") ")
	if f.ReturnType != nil {
		compileType(b, *f.ReturnType)
		b.WriteString(" ")
	}
	compileBlock(b, f.Body)
}
