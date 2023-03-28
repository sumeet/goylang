package main

import (
	"fmt"
	"github.com/kr/pretty"
	"go/types"
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
			gt := guessType(funcCall.Expr, node.Scope)
			if gt == nil {
				fmt.Printf("nil type for: ")
				pretty.Println(funcCall)
			}
			pretty.Println(gt)
			//fmt.Printf("%#v\n", guessType(funcCall.Expr, node.Scope))
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
	return found.Returns[0]
}

func (s *Scope) Lookup(name string) *Type {
	if val, ok := s.Values[name]; ok {
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
		t := lookupTypeInNamespace(scope, e.Left, e.Right)
		return &t
	case *InitializerExpr:
		t := guessGolangType(e.LHS)
		return &t
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
	}
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

	typ := scope.Lookup(nsName)
	if typ.Elided {
		return *typ
	}

	// TODO: also handle the case where it's an enum
	if !typ.Package {
		panic(fmt.Sprintf("expected package type for namespace: %#v", typ))
	}

	golangTyp := getTypeForPackage(typ.PackageName, right)
	ourTyp := golangTypeToType(golangTyp)
	ourTyp.Name = right
	//panic("got to end of lookupTypeInNamespace")
	return *ourTyp
}

func golangTypeToType(golangTyp types.Type) *Type {
	switch t := golangTyp.(type) {
	case *types.Signature:
		s := sigToType(t)
		return &s
	default:
		return newTypeStar(t.String())
	}
	panic("got to end of golangTypeToType")
}

func sigToType(sig *types.Signature) Type {
	params := make([]*Type, sig.Params().Len())
	results := make([]*Type, sig.Results().Len())
	for i := 0; i < sig.Params().Len(); i++ {
		param := sig.Params().At(i)
		params[i] = golangTypeToType(param.Type())
	}
	for i := 0; i < sig.Results().Len(); i++ {
		result := sig.Results().At(i)
		results[i] = golangTypeToType(result.Type())
	}
	return Type{
		Callable: true,
		Args:     params,
		Returns:  results,
	}
}

func compileVarRefExpr(b *strings.Builder, expr VarRefExpr) {
	b.WriteString(expr.VarName)
}

func toAnnotatedAux(node Node, scope *Scope) AnnotatedNode {
	children := node.Children()
	wrappedChildren := make([]AnnotatedNode, 0, len(children))

	switch n := node.(type) {
	case *ImportStmt:
		typ := newPkgType(n.Path)
		scope.Values[n.PkgName()] = &typ
	case *FunctionDeclaration:
		oldScope := scope

		paramTypes := make([]*Type, len(n.Params))

		scope = NewScope(scope)
		for _, param := range n.Params {
			scope.Values[param.Name] = param.Type
			paramTypes = append(paramTypes, param.Type)
		}

		typ := newFunType("function", paramTypes, []*Type{n.ReturnType})
		oldScope.Values[n.Name] = &typ
	}

	switch node.(type) {
	case *Block:
		block_scope := scope
		for _, child := range children {
			switch c := child.(type) {
			case *AssignmentStmt:
				block_scope = NewScope(block_scope)
				for _, varName := range c.VarNames {
					block_scope.Values[varName] = nil
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
	Parent *Scope
	Values map[string]*Type
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		Parent: parent,
		Values: make(map[string]*Type),
	}
}

func (a AnnotatedNode) Children() []Node {
	panic("don't use this")
}

func (a AnnotatedNode) NodeType() NodeType {
	return a.Node.NodeType()
}
