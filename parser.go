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
	MatchNodeType
	// fields, methods // type level, so variants
	InitializerNodeType
	DotAccessNodeType
	MatchArmNodeType
)

// Blah.Foo.Far
// blah.Foo().Far().Something

type DotAccessExpr struct {
	Left  Expr
	Right string
}

func (d DotAccessExpr) Children() []Node {
	return []Node{d.Left}
}

func (d DotAccessExpr) NodeType() NodeType {
	return DotAccessNodeType
}

func (d DotAccessExpr) ExprType() ExprType {
	return DotAccessExprType
}

type InitializerExpr struct {
	Type Expr
	// TODO: skipping the named params
	Args []Expr
}

func (i InitializerExpr) Children() []Node {
	var children []Node
	for _, arg := range i.Args {
		children = append(children, arg)
	}
	return children
}

func (i InitializerExpr) NodeType() NodeType {
	return InitializerNodeType
}

func (i InitializerExpr) ExprType() ExprType {
	return InitializerExprType
}

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
	case MatchNodeType:
		return "MatchNodeType"
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
	DotAccessExprType
	InitializerExprType
	MatchExprType
	BlockExprType
)

func formatExprType(t ExprType) string {
	switch t {
	case StringLiteralExprType:
		return "StringLiteralExprType"
	case FuncCallExprType:
		return "FuncCallExprType"
	case IntLiteralExprType:
		return "IntLiteralExprType"
	case VarRefExprType:
		return "VarRefExprType"
	case DotAccessExprType:
		return "DotAccessExprType"
	case InitializerExprType:
		return "InitializerExprType"
	case MatchExprType:
		return "MatchExprType"
	default:
		panic(fmt.Sprintf("unknown expr type %d", t))
	}
}

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

func (b Block) ExprType() ExprType {
	return BlockExprType
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

type MatchStmt struct {
	MatchExpr Expr
	Arms      []MatchArm
}

func (ms MatchStmt) Children() []Node {
	var children []Node
	for _, arm := range ms.Arms {
		children = append(children, arm)
	}
	return children
}

func (ms MatchStmt) ExprType() ExprType {
	return MatchExprType
}

func (ms MatchStmt) NodeType() NodeType {
	return MatchNodeType
}

type MatchArm struct {
	// TODO: when we have a less annoying language to program this in,
	// Pattern should be an enum with different variants
	//
	// for now, we'll have this always be an EnumPattern
	Pattern EnumPattern
	Body    Expr
}

func (ma MatchArm) Children() []Node {
	return []Node{ma.Body}
}

func (ma MatchArm) NodeType() NodeType {
	return MatchArmNodeType
}

type EnumPattern struct {
	Expr Expr
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
		// eat the rest of the newlines
		for {
			if len(tokens) == 0 {
				break
			}
			if tokens[0].Type == Newline {
				tokens = tokens[1:]
			} else {
				break
			}
		}
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
	old := func() (Expr, []Token) {
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
		maybeMatchStmt, tokens := tryParseMatchStmt(tokens)
		if maybeMatchStmt != nil {
			return *maybeMatchStmt, tokens
		}

		if tokens[0].Type == LCurly {
			return parseBlock(tokens)
		}

		// therefore, must be a var reference
		thisToken, tokens = consumeToken(tokens, Ident)
		varRef.VarName = thisToken.Value
		return varRef, tokens
	}
	node, tokens := old()

	// handle period
	if len(tokens) > 0 && tokens[0].Type == Dot {
		_, tokens = consumeToken(tokens, Dot)
		var thisToken Token
		thisToken, tokens = consumeToken(tokens, Ident)
		node = DotAccessExpr{Left: node, Right: thisToken.Value}
	}

	// handle initializer
	if len(tokens) > 0 && tokens[0].Type == LCurly {
		node, tokens = consumeInitializer(node, tokens)
	}

	return node, tokens
}

func tryParseMatchStmt(tokens []Token) (*MatchStmt, []Token) {
	if tokens[0].Type != Match {
		return nil, tokens
	}
	_, tokens = consumeToken(tokens, Match)
	_, tokens = consumeToken(tokens, LParen)
	var matchStmt MatchStmt
	matchStmt.MatchExpr, tokens = parseExpr(tokens)
	_, tokens = consumeToken(tokens, RParen)
	_, tokens = consumeToken(tokens, LCurly)

	// this is where the variants go
	for !peekToken(tokens, RCurly) {
		var arm MatchArm
		arm.Pattern, tokens = parsePattern(tokens)
		_, tokens = consumeToken(tokens, Colon)
		arm.Body, tokens = parseExpr(tokens)
		_, tokens = consumeToken(tokens, Comma)
		matchStmt.Arms = append(matchStmt.Arms, arm)
	}

	_, tokens = consumeToken(tokens, RCurly)
	return &matchStmt, tokens
}

func parsePattern(tokens []Token) (EnumPattern, []Token) {
	var pattern EnumPattern
	pattern.Expr, tokens = parseExpr(tokens)
	return pattern, tokens
}

func consumeInitializer(node Expr, tokens []Token) (Expr, []Token) {
	_, tokens = consumeToken(tokens, LCurly)
	var initializer InitializerExpr
	initializer.Type = node

	if peekToken(tokens, RCurly) {
		_, tokens = consumeToken(tokens, RCurly)
		return initializer, tokens
	}

	for {
		var expr Expr
		expr, tokens = parseExpr(tokens)
		initializer.Args = append(initializer.Args, expr)
		if len(tokens) > 0 && tokens[0].Type == Comma {
			_, tokens = consumeToken(tokens, Comma)
		} else {
			break
		}
	}
	_, tokens = consumeToken(tokens, RCurly)
	return initializer, tokens
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

func peekToken(tokens []Token, expectedType TokenType) bool {
	if len(tokens) == 0 {
		panic("Unexpected end of input")
	}
	if expectedType != Newline {
		tokens = skipNewlines(tokens)
	}
	return tokens[0].Type == expectedType
}

func skipNewlines(tokens []Token) []Token {
	for tokens[0].Type == Newline {
		tokens = tokens[1:]
	}
	return tokens
}

func consumeToken(tokens []Token, expectedType TokenType) (Token, []Token) {
	if len(tokens) == 0 {
		panic("Unexpected end of input")
	}
	if expectedType != Newline {
		tokens = skipNewlines(tokens)
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
