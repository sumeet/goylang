package main

import (
	"fmt"
	"strconv"
)

type NodeType uint8

const (
	ProgramNodeType NodeType = iota
	BlockNodeType
	AssignmentStmtNodeType
	ReassignmentStmtNodeType
	StringLiteralExprNodeType
	FuncCallExprNodeType
	IntLiteralExprNodeType
	VarRefExprNodeType
)

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
	statements []Statement
}

func (b Block) NodeType() NodeType {
	return BlockNodeType
}

func (b Block) Children() []Node {
	var stmts []Node
	for _, stmt := range b.statements {
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
	return ProgramNodeType
}

type Program struct {
	functions []Function
}

func (p Program) Children() []Function {
	return p.functions
}

func (p Program) NodeType() NodeType {
	return ProgramNodeType
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

func parse(tokens []Token) (program Program) {
	for len(tokens) > 0 {
		token := tokens[0]
		switch token.Type {
		case FuncDecl:
			var fn Function
			fn, tokens = parseFuncDecl(tokens)
			program.functions = append(program.functions, fn)
		case Newline:
			tokens = tokens[1:]
		default:
			panic("unhandled token type")
		}
	}
	return
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
		block.statements = append(block.statements, thisStatement)
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
