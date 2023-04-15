package main

import (
	"fmt"
	"strings"
)

func prelude() string {
	return `func slice[T any](s []T, i, j int) []T {
    return s[i:j]
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
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

func gather_imports(module Module) ([]TopLevelDeclaration, []*ImportStmt) {
	var imports []*ImportStmt
	var rest []TopLevelDeclaration
	for _, declaration := range module.Declarations {
		switch st := declaration.(type) {
		case *ImportStmt:
			imports = append(imports, st)
		default:
			rest = append(rest, declaration)
		}
	}
	return rest, imports
}

func Compile(module Module) string {
	var b strings.Builder
	b.WriteString("package main\n\n")

	declarations, imports := gather_imports(module)
	for _, imp := range imports {
		compile_import(&b, *imp)
		b.WriteByte('\n')
	}

	b.WriteString(prelude())
	b.WriteString("\n\n")
	for _, declaration := range declarations {
		compile_declaration(&b, declaration)
		b.WriteByte('\n')
	}
	return b.String()
}

func compile_declaration(b *strings.Builder, tld TopLevelDeclaration) {
	switch d := tld.(type) {
	case *ImportStmt:
		panic("compile_declaration: ImportStmt should be handled via gather_imports")
	case *Enum:
		compileEnum(b, *d)
	case *Struct:
		compileStruct(b, *d)
	case *FunctionDeclaration:
		compile_function_declaration(b, d)
	default:
		panic(fmt.Sprintf("compile_declaration: unrecognized node type %s", tld.NodeType().ToString()))
	}
}

func compileStatement(b *strings.Builder, s Statement) {
	switch st := s.(type) {
	case *Block:
		compileBlock(b, *st)
	case *AssignmentStmt:
		compileAssignmentStmt(b, *st)
	case *ReassignmentStmt:
		compileReassignmentStmt(b, *st)
	case *IntLiteralExpr:
		compileIntLiteralExpr(b, *st)
	case *StringLiteralExpr:
		compileStringLiteralExpr(b, *st)
	case *VarRefExpr:
		compileVarRefExpr(b, *st)
	case *FuncCallExpr:
		compileFuncCallExpr(b, *st)
	case *MatchStmt:
		compileMatch(b, *st)
	case *WhileExpr:
		compileWhile(b, *st)
	case *BreakExpr:
		compileBreak(b, *st)
	case *ContinueExpr:
		compileContinue(b, *st)
	case *IfExpr:
		compileIf(b, *st)
	case *ReturnExpr:
		compileReturn(b, *st)
	case *BinaryOpExpr:
		compileBinaryOp(b, *st)
	default:
		panic(fmt.Sprintf("don't know how to compile node type %s", s.NodeType().ToString()))
	}
}

func compile_import(b *strings.Builder, stmt ImportStmt) {
	b.WriteString("import ")
	b.WriteString(stmt.ImportedAs)
	b.WriteString(" ")
	b.WriteString(fmt.Sprintf("%q", stmt.PackagePath))
}

func compileBinaryOp(b *strings.Builder, expr BinaryOpExpr) {
	compileExpr(b, expr.Left)
	b.WriteString(" ")
	b.WriteString(expr.Op)
	b.WriteString(" ")
	compileExpr(b, expr.Right)
}

func compileReturn(b *strings.Builder, ret ReturnExpr) {
	b.WriteString("return")

	// multiple return values are possible
	if len(ret.Exprs) > 0 {
		b.WriteString(" ")
	}

	for i, expr := range ret.Exprs {
		if i != 0 {
			b.WriteString(", ")
		}
		compileExpr(b, expr)
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

func golangTypeNameWithBindingsThingTODORename(expr Expr) (string, *string) {
	switch e := expr.(type) {
	case *DotAccessExpr:
		var s strings.Builder
		compileInitializerLHS(&s, e)
		return s.String(), nil
	case *InitializerExpr:
		var s strings.Builder
		compileInitializerLHS(&s, e.LHS)
		if len(e.Args) == 1 {
			v := e.Args[0].(*VarRefExpr).VarName
			return s.String(), &v
		} else if len(e.Args) > 1 {
			panic("initializer with more than one arg in binding")
		} else {
			return s.String(), nil
		}
	}
	panic(fmt.Sprintf("unable to get golang type name for expr %#v", expr))
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
var CompiledFuncs map[string]*FunctionDeclaration = make(map[string]*FunctionDeclaration)

func findEnumInTable(name string) *Enum {
	for _, e := range CompiledEnums {
		if e.Name == name {
			return &e
		}
	}
	return nil
}

func findFuncInTable(name string) *FunctionDeclaration {
	f, ok := CompiledFuncs[name]
	if !ok {
		return nil
	} else {
		return f
	}
}

func compileEnum(b *strings.Builder, enum Enum) {
	CompiledEnums = append(CompiledEnums, enum)

	// iota constants for LHS enum
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

func compileExpr(b *strings.Builder, expr Expr) {
	switch e := expr.(type) {
	case *StringLiteralExpr:
		compileStringLiteralExpr(b, *e)
	case *FuncCallExpr:
		compileFuncCallExpr(b, *e)
	case *IntLiteralExpr:
		compileIntLiteralExpr(b, *e)
	case *VarRefExpr:
		compileVarRefExpr(b, *e)
	case *DotAccessExpr:
		compileDotAccessExpr(b, *e)
	case *InitializerExpr:
		compileInitializerExpr(b, *e)
	case *Block:
		compileBlock(b, *e)
	case *WhileExpr:
		compileWhile(b, *e)
	case *IfExpr:
		compileIf(b, *e)
	case *ArrayAccess:
		compileArrayAccess(b, *e)
	case *BinaryOpExpr:
		compileBinaryOp(b, *e)
	case *FunctionDeclaration:
		compile_anonymous_function_expr(b, e)
	default:
		panic(fmt.Sprintf("unable to compile expr: %#v", e))
	}
}

func compile_anonymous_function_expr(b *strings.Builder, declaration *FunctionDeclaration) {
	if len(declaration.Name) > 0 {
		panic(fmt.Sprintf("compile_anonymous_function_type: expected anonymous function; found named function with name %s", declaration.Name))
	}
	compile_named_or_anonymous_function_aux(b, declaration)
}

func compileArrayAccess(b *strings.Builder, ac ArrayAccess) {
	compileExpr(b, ac.Left)
	b.WriteString("[")
	compileExpr(b, ac.Right)
	b.WriteString("]")
}

func getVarName(expr Expr) *string {
	switch e := expr.(type) {
	case *VarRefExpr:
		vn := e.VarName
		return &vn
	default:
		return nil
	}
}

func compileDotAccessExpr(b *strings.Builder, expr DotAccessExpr) {
	vn := getVarName(expr.Left)
	if vn != nil && findEnumInTable(*vn) != nil {
		compileInitializerLHS(b, &expr)
	} else {
		compileExpr(b, expr.Left)
		b.WriteString(".")
		b.WriteString(expr.Right)
	}
}

func compileInitializerExpr(b *strings.Builder, expr InitializerExpr) {
	compileInitializerLHS(b, expr.LHS)
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
func guessGolangType(expr Expr) Type {
	switch e := expr.(type) {
	case *DotAccessExpr:
		if e.Left.NodeType() != VarRefExprNodeType {
			break
		}
		return *newTypeStar(e.Left.(*VarRefExpr).VarName)
	case *VarRefExpr:
		return *newTypeStar(e.VarName)
	}
	panic(fmt.Sprintf("couldn't print golang type name for %#v", expr))
}

func compileType(b *strings.Builder, t Type) {
	if len(t.Name) == 0 {
		panic(fmt.Sprintf("compileType: expected type name to be non-empty"))
	}

	isPointer := false
	name := t.Name

	// TODO: Pointer should be a field on Type?
	if t.Name[0] == '*' {
		isPointer = true
		name = t.Name[1:]
	}

	if isPointer {
		b.WriteString("*")
	}

	if t.Imported {
		packageImportedAs := importNameFromPackagePath(t.ImportedFrom)
		b.WriteString(packageImportedAs)
		b.WriteString(".")
	}

	b.WriteString(name)
}

func compileInitializerLHS(b *strings.Builder, expr Expr) {
	switch e := expr.(type) {
	case *DotAccessExpr:
		if e.Left.NodeType() != VarRefExprNodeType {
			break
		}
		leftNode := e.Left.(*VarRefExpr)
		rightNodeName := e.Right
		b.WriteString(leftNode.VarName)
		b.WriteString(rightNodeName)
	case *Type:
		compileType(b, *e)
	case *VarRefExpr:
		b.WriteString(e.VarName)
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

func declareVar(b *strings.Builder, varName string, varType string) {
	b.WriteString("var ")
	b.WriteString(varName)
	b.WriteString(" ")
	b.WriteString(varType)
	b.WriteString("\n")

	stfuUnusedVars(b, varName)
}

func typeListFromTypeString(typ string) []string {
	var typeList []string
	// if exprType is a (a, b, c) tuple, then we need to extract the type names
	if strings.HasPrefix(typ, "(") && strings.HasSuffix(typ, ")") {
		// split the tuple into its component types
		typ = strings.TrimPrefix(typ, "(")
		typ = strings.TrimSuffix(typ, ")")

		typeList = strings.Split(typ, ", ")
	} else {
		typeList = []string{typ}
	}
	return typeList
}

func compileAssignmentStmt(b *strings.Builder, stmt AssignmentStmt) {
	for i, varName := range stmt.VarNames {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(varName)
	}
	b.WriteString(" := ")
	compileExpr(b, stmt.Expr)
	b.WriteString("\n")
	for _, varName := range stmt.VarNames {
		stfuUnusedVars(b, varName)
	}

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

func compile_function_declaration(b *strings.Builder, f *FunctionDeclaration) {
	CompiledFuncs[f.Name] = f
	compile_named_or_anonymous_function_aux(b, f)
}

func compile_named_or_anonymous_function_aux(b *strings.Builder, f *FunctionDeclaration) {
	b.WriteString("func ")
	b.WriteString(f.Name)
	b.WriteString("(")
	for i, param := range f.Params {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(param.Name)
		b.WriteByte(' ')
		if param.Type == nil {
			b.WriteString("_elided_")
		} else {
			compileType(b, *param.Type)
		}
	}
	b.WriteString(") ")
	if f.ReturnTypeShouldBeAnArray != nil {
		compileType(b, *f.ReturnTypeShouldBeAnArray)
		b.WriteString(" ")
	}
	compileBlock(b, f.Body)
}
