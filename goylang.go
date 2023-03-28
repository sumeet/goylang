package main

import (
	"fmt"
	"github.com/kr/pretty"
	"go/types"
	"os"
	"reflect"
	"strings"
)

func typesEqual(a, b Type) bool {
	if b.Elided {
		return true
	}

	if a.Unknown || b.Unknown {
		return true
	}

	if reflect.DeepEqual(a, b) {
		return true
	}

	if !a.Callable && !b.Callable {
		return false
	}
	//ability to recover if insides are elided
	if len(a.CallableArgs) != len(b.CallableArgs) {
		return false
	}
	for i := 0; i < len(a.CallableArgs); i++ {
		if !typesEqual(*a.CallableArgs[i], *b.CallableArgs[i]) {
			return false
		}
	}
	return true
}

func setTypeForFuncDecl(funcDecl *FunctionDeclaration, paramType Type) {
	if !paramType.Callable {
		panic("paramType must be callable")
	}
	for i := 0; i < len(funcDecl.Params); i++ {
		param := funcDecl.Params[i]
		param.Type = paramType.CallableArgs[i]
	}
	// TODO: funcDecl.ReturnTypeShouldBeAnArray should be an array
	funcDecl.ReturnTypeShouldBeAnArray = paramType.CallableReturns[0]
}

func main() {
	fname := os.Args[1]
	dat, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	tokens := lex(dat)
	module := parse(tokens)
	annotated_module := toAnnotated(&module)
	_ = annotated_module

	// typeAnalysis := typeAnalyze(module)
	// _ = typeAnalysis
	s := Compile(module)
	fmt.Println(s)

	WalkAnnotated(annotated_module, func(node AnnotatedNode) {
		if node.NodeType() == FuncCallExprNodeType {
			funcCall := node.Node.(*FuncCallExpr)

			guessedTypeOfFunction := guessType(funcCall.Expr, node.Scope)
			if guessedTypeOfFunction.Unknown {
				println("unknown: ")
				println(fmt.Sprintf("%#v | %#v", funcCall.Expr, guessedTypeOfFunction))
			} else if guessedTypeOfFunction.Elided {
				println("elided: ")
				println(fmt.Sprintf("%#v | %#v", funcCall.Expr, guessedTypeOfFunction))
			} else if !guessedTypeOfFunction.Callable {
				println(fmt.Sprintf("%#v", guessedTypeOfFunction))
				panic("calling non callable")
			} else {
				if len(guessedTypeOfFunction.CallableArgs) != len(funcCall.Args) {
					panic("wrong number of args")
				}
				for i, callableArg := range guessedTypeOfFunction.CallableArgs {
					gt := guessType(funcCall.Args[i], node.Scope)

					// does guessType have what we need here?
					// gt has elided args, callableArg has concrete args
					if !typesEqual(*callableArg, *gt) {
						panic("wrong type of arg")
					} else {
						fd, ok := funcCall.Args[i].(*FunctionDeclaration)
						// the func decl's type can be changed into the guessed type
						if ok {
							setTypeForFuncDecl(fd, *gt)
						}
						// mutate b to a
					}
				}
				println("callable")
				fmt.Println("-------------------")
				pretty.Println(guessedTypeOfFunction.CallableArgs)
				fmt.Println("-------------------")
			}
		}
	})
}

func WalkAnnotated(node AnnotatedNode, f func(AnnotatedNode)) {
	f(node)
	for _, child := range node.WrappedChildren {
		WalkAnnotated(child, f)
	}
}

type TypeAnalysis struct {
}

func typeAnalyze(module Module) TypeAnalysis {
	a := toAnnotated(&module)
	_ = a
	return TypeAnalysis{}
}

// TODO: this should fall back to go
func getTypeForFuncCall(funcCall FuncCallExpr, scope *Scope) *Type {
	varRef, ok := funcCall.Expr.(*VarRefExpr)
	if !ok {
		panic(fmt.Sprintf("expected var ref expr for func call %#v", funcCall))
	}
	ident := varRef.VarName
	found := scope.Lookup(ident)
	if !found.Callable {
		panic(fmt.Sprintf("expected callable type for func call %#v", funcCall))
	}
	return found.CallableReturns[0]
}

func (s *Scope) Lookup(name string) *Type {
	if val, ok := s.TypeBySymbolName[name]; ok {
		return val
	} else if s.Parent != nil {
		return s.Parent.Lookup(name)
	} else {
		return nil
	}
}

// TODO: delete if this keeps being unused for longer periods of time 3/26/2023
func guessType(expr Expr, scope *Scope) *Type {
	switch e := expr.(type) {
	case *StringLiteralExpr:
		return newTypeStar("string")
	case *FuncCallExpr:
		return getTypeForFuncCall(*e, scope)
	case *IntLiteralExpr:
		return newTypeStar("int")
	case *VarRefExpr:
		v := scope.Lookup(e.VarName)
		if v == nil {
			return newUnknownType()
		} else {
			return v
		}
	case *DotAccessExpr:
		////// here is where "elided" "infects" dot accesses
		t := lookupTypeInNamespace(scope, e.Left, e.Right)
		return &t
	case *InitializerExpr:
		t := guessGolangType(e.LHS)
		return &t
	// anonymous function decl
	case *FunctionDeclaration:
		// use funcDeclToType() to get the type
		decl := expr.(*FunctionDeclaration)
		ft := funcDeclToType(decl)
		return &ft
	}
	//if typ, ok := e.LHS.(*LHS); ok {
	//	return typ.Name
	//} else {
	//}
	//case DotAccessExprType:
	//	dotAccessExpr := expr.(DotAccessExpr)
	//	if dotAccessExpr.Left.NodeType() != VarRefExprNodeType {
	//		break
	//	}
	//	leftNode := dotAccessExpr.Left.(VarRefExpr)
	//	rightNodeName := dotAccessExpr.Right
	//	return fmt.Sprintf("%s%s", leftNode.VarName, rightNodeName)

	panic(fmt.Sprintf("can't guess type for expr: %#v", expr))
}

func newUnknownType() *Type {
	return &Type{
		Unknown: true,
	}
}

func lookupTypeInNamespace(scope *Scope, left Expr, right string) Type {
	var nsName string
	switch left.(type) {
	case *VarRefExpr:
		varRef := left.(*VarRefExpr)
		nsName = varRef.VarName
	default:
		panic(fmt.Sprintf("expected var ref expr for left side of dot access expr: %#v", left))
	}

	////// here is where "elided" "infects" dot accesses
	typ := scope.Lookup(nsName)
	if typ.Elided {
		return *typ
	}

	// TODO: also handle the case where it's an enum
	if !typ.Package {
		panic(fmt.Sprintf("expected package type for namespace: %#v", typ))
	}

	golangTyp := getTypeForPackage(typ.PackageName, right)
	ourTyp := convertGolangTypToOurTyp(golangTyp)
	ourTyp.Name = right
	//panic("got to end of lookupTypeInNamespace")
	return *ourTyp
}

func convertGolangTypToOurTyp(golangTyp types.Type) *Type {
	var res *Type

	switch golangTyp := golangTyp.(type) {
	case *types.Signature:
		s := golangFuncSignatureToOurType(golangTyp)
		res = &s
		return res
	default:
		// TODO: there doesn't seem to be a way to get go/types to tell us the package name
		// separately from the type name, so we're parsing it ourselves for now
		queriedTypeName := golangTyp.String()
		isPointer := false
		if queriedTypeName[0] == '*' {
			queriedTypeName = queriedTypeName[1:]
			isPointer = true
		}

		sep := strings.Split(queriedTypeName, ".")
		if len(sep) == 1 {
			res = newTypeStar(sep[0])
			return res
		}
		pkgName := sep[0]
		symbol := sep[1]
		if isPointer {
			symbol = "*" + symbol
		}
		res = newTypeStar(symbol)
		res.SetImportedFrom(pkgName)
		return res
	}

	panic("got to end of convertGolangTypToOurTyp")
}

func golangFuncSignatureToOurType(sig *types.Signature) Type {
	params := make([]*Type, sig.Params().Len())
	results := make([]*Type, sig.Results().Len())
	for i := 0; i < sig.Params().Len(); i++ {
		param := sig.Params().At(i)
		params[i] = convertGolangTypToOurTyp(param.Type())
	}
	for i := 0; i < sig.Results().Len(); i++ {
		result := sig.Results().At(i)
		results[i] = convertGolangTypToOurTyp(result.Type())
	}
	return Type{
		Callable:        true,
		CallableArgs:    params,
		CallableReturns: results,
	}
}

func compileVarRefExpr(b *strings.Builder, expr VarRefExpr) {
	b.WriteString(expr.VarName)
}

func funcDeclToType(node *FunctionDeclaration) Type {
	params := make([]*Type, len(node.Params))
	for i, param := range node.Params {
		params[i] = param.Type
	}
	return newFunType(params, []*Type{node.ReturnTypeShouldBeAnArray})
}

func toAnnotatedAux(node Node, scope *Scope) AnnotatedNode {
	children := node.Children()
	wrappedChildren := make([]AnnotatedNode, 0, len(children))

	switch n := node.(type) {
	case *ImportStmt:
		typ := newPkgType(n.Path)
		scope.TypeBySymbolName[n.PkgName()] = &typ
	case *FunctionDeclaration:
		oldScope := scope
		typ := funcDeclToType(n)
		scope = NewScope(scope)
		for _, param := range n.Params {
			scope.TypeBySymbolName[param.Name] = param.Type
		}
		oldScope.TypeBySymbolName[n.Name] = &typ
	}

	switch node.(type) {
	case *Block:
		block_scope := scope
		for _, child := range children {
			switch c := child.(type) {
			case *AssignmentStmt:
				block_scope = NewScope(block_scope)
				for _, varName := range c.VarNames {
					block_scope.TypeBySymbolName[varName] = newUnknownType()
					// TODO: we should be calling guessType here
					//block_scope.TypeBySymbolName[varName] = guessType(c.Expr, block_scope)
				}
			}
			wrappedChild := toAnnotatedAux(child, block_scope)
			wrappedChildren = append(wrappedChildren, wrappedChild)
		}
	default:
		for _, child := range children {
			wrappedChild := toAnnotatedAux(child, scope)
			wrappedChildren = append(wrappedChildren, wrappedChild)
		}
	}

	return AnnotatedNode{
		Node:            node,
		Scope:           scope,
		WrappedChildren: wrappedChildren,
	}
}

func toAnnotated(root Node) AnnotatedNode {
	rootScope := NewScope(nil)
	newRoot := toAnnotatedAux(root, rootScope)
	return newRoot
}

type AnnotatedNode struct {
	Node            Node
	Scope           *Scope
	WrappedChildren []AnnotatedNode
}

type Scope struct {
	Parent           *Scope
	TypeBySymbolName map[string]*Type
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		Parent:           parent,
		TypeBySymbolName: make(map[string]*Type),
	}
}

func (a AnnotatedNode) Children() []Node {
	panic("don't use this")
}

func (a AnnotatedNode) NodeType() NodeType {
	return a.Node.NodeType()
}
