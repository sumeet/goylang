package main

import (
	"fmt"
	"strconv"
)

type NodeType uint8

const (
	ModuleNodeType NodeType = iota
	BlockNodeType
	AssignmentStmtNodeType
	ReassignmentStmtNodeType
	StringLiteralExprNodeType
	FuncCallExprNodeType
	IntLiteralExprNodeType
	VarRefExprNodeType
	FunctionNodeType
	EnumNodeType
)

func (n NodeType) ToString() string {
	switch n {
	case ModuleNodeType:
		return "Module"
	case BlockNodeType:
		return "Block"
	case AssignmentStmtNodeType:
		return "AssignmentStmt"
	case ReassignmentStmtNodeType:
		return "ReassignmentStmt"
	case StringLiteralExprNodeType:
		return "StringLiteralExpr"
	case FuncCallExprNodeType:
		return "FuncCallExpr"
	case IntLiteralExprNodeType:
		return "IntLiteralExpr"
	case VarRefExprNodeType:
		return "VarRefExpr"
	case FunctionNodeType:
		return "Function"
	case EnumNodeType:
		return "EnumNodeType"
	default:
		panic(fmt.Sprintf("unknown node type %d", n))
	}
}

func Walk(n Node, fn func(Node)) {
	fn(n)
	for _, child := range n.Children() {
		Walk(child, fn)
	}
}

type Node interface {
	Children() []Node
	NodeType() NodeType
}

type ExprType uint8

const (
	StringLiteralExprType ExprType = iota
	FuncCallExprType
	IntLiteralExprType
	VarRefExprType
)

type Statement interface {
	Node
}

type Expr interface {
	Statement
	ExprType() ExprType
}

type Block struct {
	Statements []Statement
}

func (b Block) NodeType() NodeType {
	return BlockNodeType
}

func (b Block) Children() []Node {
	var stmts []Node
	for _, stmt := range b.Statements {
		stmts = append(stmts, stmt)
	}
	return stmts
}

type Function struct {
	Name string
	Body Block
}

func (f Function) Children() []Node {
	return f.Body.Children()
}

func (f Function) NodeType() NodeType {
	return FunctionNodeType
}

type Module struct {
	Statements []Statement
}

func (p Module) Children() []Node {
	var children []Node
	for _, f := range p.Statements {
		children = append(children, f)
	}
	return children
}

func (p Module) NodeType() NodeType {
	return ModuleNodeType
}

type AssignmentStmt struct {
	VarName string
	Expr    Expr
}

func (a AssignmentStmt) NodeType() NodeType {
	return AssignmentStmtNodeType
}

func (a AssignmentStmt) Children() []Node {
	return []Node{a.Expr}
}

type ReassignmentStmt struct {
	VarName string
	Expr    Expr
}

func (r ReassignmentStmt) Children() []Node {
	return []Node{r.Expr}
}

func (r ReassignmentStmt) NodeType() NodeType {
	return ReassignmentStmtNodeType
}

type FuncCallExpr struct {
	FuncName string
	Args     []Expr
}

func (f FuncCallExpr) Children() []Node {
	var children []Node
	for _, arg := range f.Args {
		children = append(children, arg)
	}
	return children
}

func (f FuncCallExpr) NodeType() NodeType {
	return FuncCallExprNodeType
}

func (_ FuncCallExpr) ExprType() ExprType {
	return FuncCallExprType
}

type IntLiteralExpr struct {
	Value int
}

func (_ IntLiteralExpr) Children() []Node {
	return []Node{}
}

func (_ IntLiteralExpr) NodeType() NodeType {
	return IntLiteralExprNodeType
}

func (_ IntLiteralExpr) ExprType() ExprType {
	return IntLiteralExprType
}

type StringLiteralExpr struct {
	Value string
}

func (s StringLiteralExpr) Children() []Node {
	return []Node{}
}

func (s StringLiteralExpr) NodeType() NodeType {
	return StringLiteralExprNodeType
}

func (s StringLiteralExpr) ExprType() ExprType {
	return StringLiteralExprType
}

type VarRefExpr struct {
	VarName string
}

func (v VarRefExpr) Children() []Node {
	return []Node{}
}

func (v VarRefExpr) NodeType() NodeType {
	return VarRefExprNodeType
}

func (v VarRefExpr) ExprType() ExprType {
	return VarRefExprType
}

func parse(tokens []Token) (program Module) {
	for len(tokens) > 0 {
		token := tokens[0]
		switch token.Type {
		case FuncDecl:
			var fn Function
			fn, tokens = parseFuncDecl(tokens)
			program.Statements = append(program.Statements, fn)
		case EnumDecl:
			var eneom Enum
			eneom, tokens = parseEnumDecl(tokens)
			program.Statements = append(program.Statements, eneom)

		case Newline:
			tokens = tokens[1:]
		default:
			panic(fmt.Sprintf("unhandled token type: %s", formatToken(token)))
		}
	}
	return
}

func parseEnumDecl(tokens []Token) (Enum, []Token) {
	var enum Enum
	var thisToken Token

	_, tokens = consumeToken(tokens, EnumDecl)
	thisToken, tokens = consumeToken(tokens, Ident)
	enum.Name = thisToken.Value
	_, tokens = consumeToken(tokens, LCurly)
	for len(tokens) > 0 {
		_, tokens = consumeToken(tokens, Newline)
		if len(tokens) > 0 && tokens[0].Type == RCurly {
			break
		}
		var variant Variant
		variant, tokens = consumeEnumVariant(tokens)
		enum.Variants = append(enum.Variants, variant)
		_, tokens = consumeToken(tokens, Comma)
	}
	_, tokens = consumeToken(tokens, RCurly)
	return enum, tokens
}

func consumeEnumVariant(tokens []Token) (Variant, []Token) {
	var thisToken Token
	thisToken, tokens = consumeToken(tokens, Ident)
	variant := Variant{Name: thisToken.Value}
	if len(tokens) > 0 && tokens[0].Type == LParen {
		_, tokens = consumeToken(tokens, LParen)
		thisToken, tokens = consumeToken(tokens, Ident)
		variant.Type = &Type{Name: thisToken.Value}
		_, tokens = consumeToken(tokens, RParen)
	}
	return variant, tokens
}

type Enum struct {
	Name     string
	Variants []Variant
}

func (e Enum) NodeType() NodeType {
	return EnumNodeType
}

func (e Enum) Children() []Node {
	return []Node{} // TODO: expand on this later
}

type Type struct {
	Name string
}

type Variant struct {
	Name string
	Type *Type
}

func parseFuncDecl(tokens []Token) (Function, []Token) {
	var fn Function
	var thisToken Token

	_, tokens = consumeToken(tokens, FuncDecl)
	thisToken, tokens = consumeToken(tokens, Ident)
	fn.Name = thisToken.Value
	_, tokens = consumeToken(tokens, LParen)
	// TODO: args
	_, tokens = consumeToken(tokens, RParen)
	fn.Body, tokens = parseBlock(tokens)
	return fn, tokens
}

func parseBlock(tokens []Token) (Block, []Token) {
	var block Block
	var thisStatement Statement

	_, tokens = consumeToken(tokens, LCurly)
	for len(tokens) > 0 {
		_, tokens = consumeToken(tokens, Newline)
		if tokens[0].Type == RCurly {
			_, tokens = consumeToken(tokens, RCurly)
			break
		}
		thisStatement, tokens = parseStatement(tokens)
		block.Statements = append(block.Statements, thisStatement)
	}
	return block, tokens
}

func parseStatement(tokens []Token) (Statement, []Token) {
	if tokens[0].Type == Ident && tokens[1].Type == Assignment {
		return parseAssignment(tokens)
	} else if tokens[0].Type == Ident && tokens[1].Type == Reassignment {
		return parseReassignment(tokens)
	} else {
		// must be an expr
		return parseExpr(tokens)
	}
}

func parseAssignment(tokens []Token) (AssignmentStmt, []Token) {
	var assignmentStmt AssignmentStmt
	var thisToken Token
	thisToken, tokens = consumeToken(tokens, Ident)
	assignmentStmt.VarName = thisToken.Value
	_, tokens = consumeToken(tokens, Assignment)
	assignmentStmt.Expr, tokens = parseExpr(tokens)
	return assignmentStmt, tokens
}

func parseExpr(tokens []Token) (Expr, []Token) {
	var maybeFuncCall *FuncCallExpr
	var maybeIntLiteral *IntLiteralExpr
	var varRef VarRefExpr
	var thisToken Token

	if tokens[0].Type == StringLiteral {
		if value, err := strconv.Unquote(tokens[0].Value); err != nil {
			panic(fmt.Sprintf("unable to unquote string: %s: %s", tokens[0].Value, err))
		} else {
			return StringLiteralExpr{Value: value}, tokens[1:]
		}
	}

	maybeFuncCall, tokens = tryParseFuncCall(tokens)
	if maybeFuncCall != nil {
		return *maybeFuncCall, tokens
	}
	maybeIntLiteral, tokens = tryParseIntLiteral(tokens)
	if maybeIntLiteral != nil {
		return *maybeIntLiteral, tokens
	}

	// therefore, must be a var reference
	thisToken, tokens = consumeToken(tokens, Ident)
	varRef.VarName = thisToken.Value
	return varRef, tokens
}

func tryParseFuncCall(tokens []Token) (*FuncCallExpr, []Token) {
	if tokens[0].Type == Ident && tokens[1].Type == LParen {
		var funcCall FuncCallExpr
		funcCall, tokens = parseFunctionCall(tokens)
		return &funcCall, tokens
	} else {
		return nil, tokens
	}
}

func tryParseIntLiteral(tokens []Token) (*IntLiteralExpr, []Token) {
	var thisToken Token
	var intLiteral IntLiteralExpr
	var err error

	if tokens[0].Type == IntLiteral {
		thisToken, tokens = consumeToken(tokens, IntLiteral)
		if intLiteral.Value, err = strconv.Atoi(thisToken.Value); err != nil {
			panic(err)
		}
		return &intLiteral, tokens
	} else {
		return nil, tokens
	}
}

func parseReassignment(tokens []Token) (ReassignmentStmt, []Token) {
	return ReassignmentStmt{}, tokens
}

func parseFunctionCall(tokens []Token) (FuncCallExpr, []Token) {
	var funcCall FuncCallExpr
	var thisToken Token
	var thisExpr Expr

	thisToken, tokens = consumeToken(tokens, Ident)
	funcCall.FuncName = thisToken.Value
	_, tokens = consumeToken(tokens, LParen)
	for true {
		thisExpr, tokens = parseExpr(tokens)
		funcCall.Args = append(funcCall.Args, thisExpr)
		if tokens[0].Type == RParen {
			_, tokens = consumeToken(tokens, RParen)
			break
		} else {
			_, tokens = consumeToken(tokens, Comma)
		}
	}
	return funcCall, tokens
}

func consumeToken(tokens []Token, expectedType TokenType) (Token, []Token) {
	if len(tokens) == 0 {
		panic("Unexpected end of input")
	}
	token := tokens[0]
	if token.Type != expectedType {
		fmt.Println("remaining tokens:")
		printTokens(tokens)
		panic(fmt.Sprintf("Was expecting token %s, got %s", formatTokenType(expectedType),
			formatToken(token)))
	}
	return token, tokens[1:]
}
