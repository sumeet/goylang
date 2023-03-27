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
	StructNodeType
	WhileNodeType
	BreakNodeType
	ContinueNodeType
	IfNodeType
	ElseNodeType
	ReturnNodeType
	ArrayAccessNodeType
	BinaryOpNodeType
)

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
	case StructNodeType:
		return "StructNodeType"
	case WhileNodeType:
		return "WhileNodeType"
	case BreakNodeType:
		return "BreakNodeType"
	case ContinueNodeType:
		return "ContinueNodeType"
	case IfNodeType:
		return "IfNodeType"
	case ElseNodeType:
		return "ElseNodeType"
	case ReturnNodeType:
		return "ReturnNodeType"
	case BinaryOpNodeType:
		return "BinaryOpNodeType"
	default:
		panic(fmt.Sprintf("unknown node type %d", n))
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
	TypeExprType
	WhileExprType
	BreakExprType
	ContinueExprType
	IfExprType
	ArrayAccessExprType
	BinaryOpExprType
	FuncDeclExprType
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
	case BreakExprType:
		return "BreakExprType"
	case ContinueExprType:
		return "ContinueExprType"
	case WhileExprType:
		return "WhileExprType"
	case BinaryOpExprType:
		return "BinaryOpExprType"
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

type BinaryOpExpr struct {
	Left  Expr
	Right Expr
	Op    string
}

func (b BinaryOpExpr) Children() []Node {
	return []Node{b.Left, b.Right}
}

func (b BinaryOpExpr) NodeType() NodeType {
	return BinaryOpNodeType
}

func (b BinaryOpExpr) ExprType() ExprType {
	return BinaryOpExprType
}

type Function struct {
	Name       string
	Params     []Param
	Body       Block
	ReturnType *Type
}

func (f Function) ExprType() ExprType {
	return FuncDeclExprType
}

type Param struct {
	Name string
	Type Type
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
	VarNames []string
	Expr     Expr
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
	VarNames []string
	Expr     Expr
}

func (r ReassignmentStmt) Children() []Node {
	return []Node{r.Expr}
}

func (r ReassignmentStmt) NodeType() NodeType {
	return ReassignmentStmtNodeType
}

type FuncCallExpr struct {
	Expr Expr
	Args []Expr
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

type Struct struct {
	Name   string
	Fields []StructField
}

func (s Struct) Children() []Node {
	//TODO implement me
	panic("implement me")
}

func (s Struct) NodeType() NodeType {
	//TODO implement me
	return StructNodeType
}

type StructField struct {
	Name string
	Type Type
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
		case StructDecl:
			var strukt Struct
			strukt, tokens = parseStructDecl(tokens)
			program.Statements = append(program.Statements, strukt)
		case Newline:
			tokens = tokens[1:]
		default:
			panic(fmt.Sprintf("unhandled token type: %s", formatToken(token)))
		}
	}
	return
}

func parseStructDecl(tokens []Token) (Struct, []Token) {
	var strukt Struct
	var thisToken Token
	_, tokens = consumeToken(tokens, StructDecl)
	thisToken, tokens = consumeToken(tokens, Ident)
	strukt.Name = thisToken.Value

	var field StructField
	_, tokens = consumeToken(tokens, LCurly)
	for len(tokens) > 0 {
		if peekToken(tokens, RCurly) {
			break
		}

		thisToken, tokens = consumeToken(tokens, Ident)
		field.Name = thisToken.Value
		field.Type, tokens = parseType(tokens)

		strukt.Fields = append(strukt.Fields, field)
		_, tokens = consumeToken(tokens, Comma)
	}
	_, tokens = consumeToken(tokens, RCurly)
	return strukt, tokens
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
		//variant.Type = &Type{Name: thisToken.Value}
		variant.Type = &thisToken.Value
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

type Variant struct {
	Name string
	Type *string
}

func parseAnonFuncDecl(tokens []Token) (Function, []Token) {
	var fn Function
	var thisToken Token

	_, tokens = consumeToken(tokens, FuncDecl)
	_, tokens = consumeToken(tokens, LParen)

	for len(tokens) > 0 {
		if peekToken(tokens, RParen) {
			break
		}
		var param Param
		thisToken, tokens = consumeToken(tokens, Ident)
		param.Name = thisToken.Value
		param.Type, tokens = parseType(tokens)
		fn.Params = append(fn.Params, param)

		if peekToken(tokens, RParen) {
			break
		}

		_, tokens = consumeToken(tokens, Comma)
	}

	_, tokens = consumeToken(tokens, RParen)

	if !peekToken(tokens, LCurly) {
		var returnType Type
		returnType, tokens = parseType(tokens)
		fn.ReturnType = &returnType
	}

	fn.Body, tokens = parseBlock(tokens)
	return fn, tokens
}

func parseFuncDecl(tokens []Token) (Function, []Token) {
	var fn Function
	var thisToken Token

	_, tokens = consumeToken(tokens, FuncDecl)
	thisToken, tokens = consumeToken(tokens, Ident)
	fn.Name = thisToken.Value
	_, tokens = consumeToken(tokens, LParen)

	for len(tokens) > 0 {
		if peekToken(tokens, RParen) {
			break
		}
		var param Param
		thisToken, tokens = consumeToken(tokens, Ident)
		param.Name = thisToken.Value
		param.Type, tokens = parseType(tokens)
		fn.Params = append(fn.Params, param)

		if peekToken(tokens, RParen) {
			break
		}

		_, tokens = consumeToken(tokens, Comma)
	}

	_, tokens = consumeToken(tokens, RParen)

	if !peekToken(tokens, LCurly) {
		var returnType Type
		returnType, tokens = parseType(tokens)
		fn.ReturnType = &returnType
	}

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

func getLValuesFor(tokens []Token, tokenType TokenType) ([]string, []Token) {
	origTokens := tokens
	var lValues []string
	var thisToken Token
	for len(tokens) > 0 {
		if peekToken(tokens, Ident, Comma) {
			thisToken, tokens = consumeToken(tokens, Ident)
			lValues = append(lValues, thisToken.Value)
			thisToken, tokens = consumeToken(tokens, Comma)
			continue
		} else if peekToken(tokens, Ident, tokenType) {
			thisToken, tokens = consumeToken(tokens, Ident)
			lValues = append(lValues, thisToken.Value)
			return lValues, tokens
		} else {
			return nil, origTokens
		}
	}
	panic("unreachable")
}

func parseStatement(tokens []Token) (Statement, []Token) {
	var lValues []string
	lValues, tokens = getLValuesFor(tokens, Assignment)
	if len(lValues) > 0 {
		return parseAssignment(tokens, lValues)
	}
	lValues, tokens = getLValuesFor(tokens, Reassignment)
	if len(lValues) > 0 {
		return parseReassignment(tokens, lValues)
	}
	// else must be an expr
	return parseExpr(tokens)
}

func parseAssignment(tokens []Token, lValues []string) (AssignmentStmt, []Token) {
	var stmt AssignmentStmt
	stmt.VarNames = lValues
	_, tokens = consumeToken(tokens, Assignment)
	stmt.Expr, tokens = parseExpr(tokens)
	return stmt, tokens
}

func parseReassignment(tokens []Token, lValues []string) (ReassignmentStmt, []Token) {
	var stmt ReassignmentStmt
	stmt.VarNames = lValues
	_, tokens = consumeToken(tokens, Reassignment)
	stmt.Expr, tokens = parseExpr(tokens)
	return stmt, tokens
}

type ParseExprOpts struct {
	SkipMatch bool
}

func consumeBinaryOperator(leftNode Expr, tokens []Token) (BinaryOpExpr, []Token) {
	var thisToken Token
	var expr BinaryOpExpr
	expr.Left = leftNode
	thisToken, tokens = consumeToken(tokens, BinaryOp)
	expr.Op = thisToken.Value
	expr.Right, tokens = parseExpr(tokens)
	return expr, tokens
}

func parseExpr(tokens []Token) (Expr, []Token) {
	old := func() (Expr, []Token) {
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

		// slices means we're definitely gonna see a type coming, for now
		if peekToken(tokens, LBracket, RBracket) {
			return parseType(tokens)
		}

		if tokens[0].Type == While {
			return parseWhile(tokens)
		}

		if tokens[0].Type == Break {
			return parseBreak(tokens)
		}

		if tokens[0].Type == Continue {
			return parseContinue(tokens)
		}

		if tokens[0].Type == If {
			return parseIf(tokens)
		}

		if tokens[0].Type == Return {
			return parseReturn(tokens)
		}

		if tokens[0].Type == FuncDecl {
			return parseAnonFuncDecl(tokens)
		}

		// therefore, must be a var reference
		thisToken, tokens = consumeToken(tokens, Ident)
		varRef.VarName = thisToken.Value
		return varRef, tokens
	}
	node, tokens := old()

post:
	// handle period
	if peekToken(tokens, Dot) {
		_, tokens = consumeToken(tokens, Dot)
		var thisToken Token
		thisToken, tokens = consumeToken(tokens, Ident)
		node = DotAccessExpr{Left: node, Right: thisToken.Value}
		goto post
	}

	// handle initializer
	if len(tokens) > 0 && tokens[0].Type == LCurly {
		node, tokens = consumeInitializer(node, tokens)
		goto post
	}

	// handle function calls
	if len(tokens) > 0 && tokens[0].Type == LParen {
		node, tokens = consumeFuncCall(node, tokens)
		goto post
	}

	// handle array access
	if len(tokens) > 0 && tokens[0].Type == LBracket {
		node, tokens = consumeArrayAccess(node, tokens)
		goto post
	}

	// handle binary operators
	if len(tokens) > 0 && tokens[0].Type == BinaryOp {
		node, tokens = consumeBinaryOperator(node, tokens)
		goto post
	}

	return node, tokens
}

type ArrayAccess struct {
	Left  Expr
	Right Expr
}

func (a ArrayAccess) Children() []Node {
	//TODO implement me
	panic("implement me")
}

func (a ArrayAccess) NodeType() NodeType {
	return ArrayAccessNodeType
}

func (a ArrayAccess) ExprType() ExprType {
	return ArrayAccessExprType
}

func consumeArrayAccess(node Expr, tokens []Token) (ArrayAccess, []Token) {
	var ac ArrayAccess
	ac.Left = node
	_, tokens = consumeToken(tokens, LBracket)
	ac.Right, tokens = parseExpr(tokens)
	_, tokens = consumeToken(tokens, RBracket)
	return ac, tokens
}

type ReturnExpr struct {
	Exprs []Expr
}

func (r ReturnExpr) Children() []Node {
	var children []Node
	for _, expr := range r.Exprs {
		children = append(children, expr)
	}
	return children
}

func (r ReturnExpr) NodeType() NodeType {
	return ReturnNodeType
}

func (r ReturnExpr) ExprType() ExprType {
	//TODO implement me
	panic("implement me")
}

func parseReturn(tokens []Token) (ReturnExpr, []Token) {
	var ret ReturnExpr
	_, tokens = consumeToken(tokens, Return)

	if !peekToken(tokens, Newline) {
		var expr Expr
		expr, tokens = parseExpr(tokens)
		ret.Exprs = append(ret.Exprs, expr)

		for peekToken(tokens, Comma) {
			_, tokens = consumeToken(tokens, Comma)
			expr, tokens = parseExpr(tokens)
			ret.Exprs = append(ret.Exprs, expr)
		}
	}

	return ret, tokens
}

type IfExpr struct {
	Cond     Expr
	IfBody   Expr
	ElseBody *Expr
}

func (i IfExpr) Children() []Node {
	//TODO implement me
	panic("implement me")
}

func (i IfExpr) NodeType() NodeType {
	return IfNodeType
}

func (i IfExpr) ExprType() ExprType {
	return IfExprType
}

func parseIf(tokens []Token) (IfExpr, []Token) {
	_, tokens = consumeToken(tokens, If)
	var ifExpr IfExpr

	_, tokens = consumeToken(tokens, LParen)
	ifExpr.Cond, tokens = parseExpr(tokens)
	_, tokens = consumeToken(tokens, RParen)

	ifExpr.IfBody, tokens = parseExpr(tokens)
	if peekToken(tokens, Else) {
		_, tokens = consumeToken(tokens, Else)
		var elseBody Expr
		elseBody, tokens = parseExpr(tokens)
		ifExpr.ElseBody = &elseBody
	}
	return ifExpr, tokens
}

type BreakExpr struct {
}

func (b BreakExpr) Children() []Node {
	//TODO implement me
	panic("implement me")
}

func (b BreakExpr) NodeType() NodeType {
	return BreakNodeType
}

func (b BreakExpr) ExprType() ExprType {
	return BreakExprType
}

func parseBreak(tokens []Token) (BreakExpr, []Token) {
	_, tokens = consumeToken(tokens, Break)
	return BreakExpr{}, tokens
}

type ContinueExpr struct {
}

func (b ContinueExpr) Children() []Node {
	//TODO implement me
	panic("implement me")
}

func (c ContinueExpr) NodeType() NodeType {
	return ContinueNodeType
}

func (c ContinueExpr) ExprType() ExprType {
	return ContinueExprType
}

func parseContinue(tokens []Token) (ContinueExpr, []Token) {
	_, tokens = consumeToken(tokens, Continue)
	return ContinueExpr{}, tokens
}

type WhileExpr struct {
	Body Expr
}

func (w WhileExpr) Children() []Node {
	//TODO implement me
	panic("implement me")
}

func (w WhileExpr) NodeType() NodeType {
	return WhileNodeType
}

func (w WhileExpr) ExprType() ExprType {
	return WhileExprType
}

func parseWhile(tokens []Token) (WhileExpr, []Token) {
	_, tokens = consumeToken(tokens, While)
	var w WhileExpr
	w.Body, tokens = parseExpr(tokens)
	return w, tokens
}

type Type struct {
	Name string
}

func (s Type) Children() []Node {
	//TODO implement me
	panic("implement me")
}

func (s Type) NodeType() NodeType {
	//TODO implement me
	panic("implement me")
}

func (s Type) ExprType() ExprType {
	return TypeExprType
}

func parseType(tokens []Token) (Type, []Token) {
	var tn []byte
	var typ Type

	if peekToken(tokens, BinaryOp) && tokens[0].Value == "*" {
		tn = append(tn, '*')
		_, tokens = consumeToken(tokens, BinaryOp)
	}

	// maybe it's a tuple-like list of multiple return values for a function.
	// parse naively as a list of type names (rather, don't call this function
	// again for each type, because you can't have a tuple of tuples)
	if peekToken(tokens, LParen) {
		_, tokens = consumeToken(tokens, LParen)
		var thisToken Token
		thisToken, tokens = consumeToken(tokens, Ident)
		tn = append(tn, '(')
		tn = append(tn, []byte(thisToken.Value)...)
		for peekToken(tokens, Comma) {
			tn = append(tn, []byte(", ")...)
			_, tokens = consumeToken(tokens, Comma)
			thisToken, tokens = consumeToken(tokens, Ident)
			tn = append(tn, []byte(thisToken.Value)...)
		}
		tn = append(tn, ')')
		_, tokens = consumeToken(tokens, RParen)
		typ.Name = string(tn)
		return typ, tokens
	}

	// maybe it's a slice type
	if peekToken(tokens, LBracket) {
		_, tokens = consumeToken(tokens, LBracket)
		_, tokens = consumeToken(tokens, RBracket)
		tn = append(tn, '[', ']')
	}

	// otherwise it's an ident
	var thisToken Token
	thisToken, tokens = consumeToken(tokens, Ident)
	tn = append(tn, []byte(thisToken.Value)...)
	typ.Name = string(tn)
	return typ, tokens
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

func consumeFuncCall(expr Expr, tokens []Token) (FuncCallExpr, []Token) {
	var funcCall FuncCallExpr
	var thisExpr Expr

	funcCall.Expr = expr
	_, tokens = consumeToken(tokens, LParen)
	for true {
		if tokens[0].Type == RParen {
			_, tokens = consumeToken(tokens, RParen)
			break
		}

		thisExpr, tokens = parseExpr(tokens)
		funcCall.Args = append(funcCall.Args, thisExpr)

		if tokens[0].Type == Comma {
			_, tokens = consumeToken(tokens, Comma)
		}
	}
	return funcCall, tokens
}

func peekToken(tokens []Token, expectedTypes ...TokenType) bool {
	if len(tokens) == 0 {
		panic("Unexpected end of input")
	}
	for _, expectedType := range expectedTypes {
		if expectedType != Newline {
			tokens = skipNewlines(tokens)
		}
		if tokens[0].Type != expectedType {
			return false
		}
		tokens = tokens[1:]
	}
	return true
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
