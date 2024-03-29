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
	ImportStmtNodeType
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
		return "FunctionDeclaration"
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
	case ImportStmtNodeType:
		return "ImportStmtNodeType"
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
	FunctionExprType
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

/* Top-level declarations ( ====================================================
 */

type Module struct{ Declarations []TopLevelDeclaration }

func (_ *Module) NodeType() NodeType { return ModuleNodeType }
func (p *Module) Children() []Node {
	var children []Node
	for _, f := range p.Declarations {
		children = append(children, f)
	}
	return children
}

type TopLevelDeclaration interface {
	Node
	_is_top_level_declaration()
}

type Enum struct {
	Name     string
	Variants []Variant
}
type Struct struct {
	Name   string
	Fields []StructField
}
type StructField struct {
	Name string
	Type Type
}

type ImportStmt struct {
	PackagePath string
	ImportedAs  string
}

type FunctionDeclaration struct {
	Name        string
	Params      []Param
	Body        Block
	ReturnTypes []Type
}

func (x *Struct) Children() []Node                        { return []Node{} }
func (_ *Struct) NodeType() NodeType                      { return StructNodeType }
func (_ *Struct) _is_top_level_declaration()              {}
func (_ *Enum) NodeType() NodeType                        { return EnumNodeType }
func (_ *Enum) Children() []Node                          { return []Node{} /* TODO: expand on this later */ }
func (_ *Enum) _is_top_level_declaration()                {}
func (_ *ImportStmt) Children() []Node                    { return []Node{} }
func (_ *ImportStmt) NodeType() NodeType                  { return ImportStmtNodeType }
func (_ *ImportStmt) _is_top_level_declaration()          {}
func (f *FunctionDeclaration) Children() []Node           { return []Node{&f.Body} }
func (_ *FunctionDeclaration) NodeType() NodeType         { return FunctionNodeType }
func (_ *FunctionDeclaration) _is_top_level_declaration() {}
func (_ *FunctionDeclaration) _is_statement()             {} // HACK
func (_ *FunctionDeclaration) ExprType() ExprType         { return FunctionExprType }

// ) Top-level declarations ====================================================

/* Statements ( ================================================================
 */

type Statement interface {
	Node
	_is_statement()
}

type AssignmentStmt struct {
	VarNames []string
	Expr     Expr
}
type ReassignmentStmt struct {
	VarNames []string
	Expr     Expr
}

func (a *AssignmentStmt) NodeType() NodeType   { return AssignmentStmtNodeType }
func (a *AssignmentStmt) Children() []Node     { return []Node{a.Expr} }
func (a *AssignmentStmt) _is_statement()       {}
func (r *ReassignmentStmt) Children() []Node   { return []Node{r.Expr} }
func (r *ReassignmentStmt) NodeType() NodeType { return ReassignmentStmtNodeType }
func (r *ReassignmentStmt) _is_statement()     {}

// ) Statements ================================================================

/* Expressions ( ===============================================================
 */

type Expr interface {
	Statement
	ExprType() ExprType
}

type Block struct{ Statements []Statement }
type BinaryOpExpr struct {
	Left  Expr
	Right Expr
	Op    string
}
type Param struct {
	Name string
	Type *Type
}
type MatchStmt struct {
	MatchExpr Expr
	Arms      []MatchArm
}
type MatchArm struct {
	// TODO: when we have a less annoying language to program this in,
	// Pattern should be an enum with different variants
	//
	// for now, we'll have this always be an EnumPattern
	Pattern EnumPattern
	Body    Expr
}
type EnumPattern struct{ Expr Expr }
type FuncCallExpr struct {
	Expr Expr
	Args []Expr
}
type IntLiteralExpr struct{ Value int }
type StringLiteralExpr struct{ Value string }
type VarRefExpr struct{ VarName string }
type BreakExpr struct{}
type WhileExpr struct{ Body Expr }
type ContinueExpr struct{}
type IfExpr struct {
	Cond     Expr
	IfBody   Expr
	ElseBody *Expr
}
type ReturnExpr struct{ Exprs []Expr }
type DotAccessExpr struct {
	Left  Expr
	Right string
}
type InitializerExpr struct {
	LHS Expr
	// TODO: skipping the named params
	Args []Expr
}
type ArrayAccess struct {
	Left  Expr
	Right Expr
}

type Type struct {
	Name    string
	Unknown bool
	// Enums
	Elided                 bool
	Callable               bool
	CallableArgs           []*Type
	CallableArgsIsVariadic bool
	CallableReturns        []*Type
	Package                bool
	PackageName            string
	Imported               bool
	ImportedFrom           string
}

func (t *Type) SetImportedFrom(path string) {
	t.Imported = true
	t.ImportedFrom = path
}

func newElidedType() *Type {
	return &Type{
		Name:   "elided",
		Elided: true,
	}
}

func newPkgType(name string) Type {
	return Type{
		Name:        "package",
		Package:     true,
		PackageName: name,
	}
}

func newTypeStar(name string) *Type {
	return &Type{Name: name}
}

func newFunType(argTypes []*Type, returnTypes []*Type) Type {
	return Type{
		Callable:        true,
		CallableArgs:    argTypes,
		CallableReturns: returnTypes,
	}
}

func (b *Block) ExprType() ExprType { return BlockExprType }
func (b *Block) NodeType() NodeType { return BlockNodeType }
func (b *Block) Children() []Node {
	var stmts []Node
	for _, stmt := range b.Statements {
		stmts = append(stmts, stmt)
	}
	return stmts
}
func (_ *Block) _is_statement()            {}
func (b *BinaryOpExpr) Children() []Node   { return []Node{b.Left, b.Right} }
func (b *BinaryOpExpr) NodeType() NodeType { return BinaryOpNodeType }
func (b *BinaryOpExpr) ExprType() ExprType { return BinaryOpExprType }
func (_ *BinaryOpExpr) _is_statement()     {}
func (ms *MatchStmt) Children() []Node {
	var children []Node
	for _, arm := range ms.Arms {
		children = append(children, &arm)
	}
	return children
}
func (ms *MatchStmt) ExprType() ExprType { return MatchExprType }
func (ms *MatchStmt) NodeType() NodeType { return MatchNodeType }
func (ms *MatchStmt) _is_statement()     {}
func (ma *MatchArm) Children() []Node    { return []Node{ma.Body} }
func (ma *MatchArm) NodeType() NodeType  { return MatchArmNodeType }
func (f *FuncCallExpr) Children() []Node {
	var children []Node
	for _, arg := range f.Args {
		children = append(children, arg)
	}
	return children
}
func (f *FuncCallExpr) NodeType() NodeType      { return FuncCallExprNodeType }
func (_ *FuncCallExpr) ExprType() ExprType      { return FuncCallExprType }
func (_ *FuncCallExpr) _is_statement()          {}
func (_ *IntLiteralExpr) Children() []Node      { return []Node{} }
func (_ *IntLiteralExpr) NodeType() NodeType    { return IntLiteralExprNodeType }
func (_ *IntLiteralExpr) ExprType() ExprType    { return IntLiteralExprType }
func (_ *IntLiteralExpr) _is_statement()        {}
func (s *StringLiteralExpr) Children() []Node   { return []Node{} }
func (s *StringLiteralExpr) NodeType() NodeType { return StringLiteralExprNodeType }
func (s *StringLiteralExpr) ExprType() ExprType { return StringLiteralExprType }
func (_ *StringLiteralExpr) _is_statement()     {}
func (v *VarRefExpr) Children() []Node          { return []Node{} }
func (v *VarRefExpr) NodeType() NodeType        { return VarRefExprNodeType }
func (v *VarRefExpr) ExprType() ExprType        { return VarRefExprType }
func (_ *VarRefExpr) _is_statement()            {}
func (_ *BreakExpr) Children() []Node           { return []Node{} }
func (_ *BreakExpr) NodeType() NodeType         { return BreakNodeType }
func (_ *BreakExpr) ExprType() ExprType         { return BreakExprType }
func (_ *BreakExpr) _is_statement()             {}
func (x *WhileExpr) Children() []Node           { return []Node{x.Body} }
func (_ *WhileExpr) NodeType() NodeType         { return WhileNodeType }
func (_ *WhileExpr) ExprType() ExprType         { return WhileExprType }
func (_ *WhileExpr) _is_statement()             {}
func (_ *ContinueExpr) Children() []Node        { return []Node{} }
func (_ *ContinueExpr) NodeType() NodeType      { return ContinueNodeType }
func (_ *ContinueExpr) ExprType() ExprType      { return ContinueExprType }
func (_ *ContinueExpr) _is_statement()          {}
func (x *IfExpr) Children() []Node {
	ret := make([]Node, 0, 3)
	ret = append(ret, x.Cond)
	ret = append(ret, x.IfBody)
	if x.ElseBody != nil {
		ret = append(ret, *x.ElseBody)
	}
	return ret
}
func (_ *IfExpr) NodeType() NodeType { return IfNodeType }
func (_ *IfExpr) ExprType() ExprType { return IfExprType }
func (_ *IfExpr) _is_statement()     {}
func (r *ReturnExpr) Children() []Node {
	var children []Node
	for _, expr := range r.Exprs {
		children = append(children, expr)
	}
	return children
}
func (_ *ReturnExpr) NodeType() NodeType    { return ReturnNodeType }
func (_ *ReturnExpr) ExprType() ExprType    { panic("implement me") }
func (_ *ReturnExpr) _is_statement()        {}
func (d *DotAccessExpr) Children() []Node   { return []Node{d.Left} }
func (_ *DotAccessExpr) NodeType() NodeType { return DotAccessNodeType }
func (_ *DotAccessExpr) ExprType() ExprType { return DotAccessExprType }
func (_ *DotAccessExpr) _is_statement()     {}
func (i *InitializerExpr) Children() []Node {
	var children []Node
	for _, arg := range i.Args {
		children = append(children, arg)
	}
	return children
}
func (_ *InitializerExpr) NodeType() NodeType { return InitializerNodeType }
func (_ *InitializerExpr) ExprType() ExprType { return InitializerExprType }
func (_ *InitializerExpr) _is_statement()     {}
func (x *ArrayAccess) Children() []Node       { return []Node{x.Left, x.Right} }
func (_ *ArrayAccess) NodeType() NodeType     { return ArrayAccessNodeType }
func (_ *ArrayAccess) ExprType() ExprType     { return ArrayAccessExprType }
func (_ *ArrayAccess) _is_statement()         {}
func (_ *Type) Children() []Node              { panic("implement me") }
func (_ *Type) NodeType() NodeType            { panic("implement me") }
func (_ *Type) ExprType() ExprType            { return TypeExprType }
func (_ *Type) _is_statement()                {}

// ) Expressions ===============================================================

func parse(tokens []Token) (program Module) {
	for len(tokens) > 0 {
		token := tokens[0]
		switch token.Type {
		case FuncDecl:
			var fn FunctionDeclaration
			fn, tokens = parse_top_level_function_declaration(tokens)
			program.Declarations = append(program.Declarations, &fn)
		case EnumDecl:
			var enum Enum
			enum, tokens = parseEnumDecl(tokens)
			program.Declarations = append(program.Declarations, &enum)
		case StructDecl:
			var struct_ Struct
			struct_, tokens = parseStructDecl(tokens)
			program.Declarations = append(program.Declarations, &struct_)
		case Import:
			var import_ ImportStmt
			import_, tokens = parseImportStmt(tokens)
			program.Declarations = append(program.Declarations, &import_)
		case Newline:
			tokens = tokens[1:]
		default:
			panic(fmt.Sprintf("unhandled token type: %s", formatToken(token)))
		}
	}
	return
}

func parseImportStmt(tokens []Token) (ImportStmt, []Token) {
	var importStmt ImportStmt
	var thisToken Token
	_, tokens = consumeToken(tokens, Import)
	thisToken, tokens = consumeToken(tokens, StringLiteral)
	var err error
	importStmt.PackagePath, err = strconv.Unquote(thisToken.Value)
	importStmt.ImportedAs = packageScopeNameFromPackagePath(importStmt.PackagePath)
	if err != nil {
		panic(err)
	}
	return importStmt, tokens
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
		//variant.LHS = &LHS{Name: thisToken.Value}
		variant.Type = &thisToken.Value
		_, tokens = consumeToken(tokens, RParen)
	}
	return variant, tokens
}

type Variant struct {
	Name string
	Type *string
}

func parseAnonFuncDecl(tokens []Token) (FunctionDeclaration, []Token) {
	var fn FunctionDeclaration
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
		if !peekToken(tokens, RParen) && !peekToken(tokens, Comma) {
			var typ Type
			typ, tokens = parseType(tokens)
			param.Type = &typ
		} else {
			param.Type = newElidedType()
		}
		fn.Params = append(fn.Params, param)

		if peekToken(tokens, RParen) {
			break
		}

		_, tokens = consumeToken(tokens, Comma)
	}

	_, tokens = consumeToken(tokens, RParen)

	if !peekToken(tokens, LCurly) {
		var returnTypes []Type
		returnTypes, tokens = parseReturnTypes(tokens)
		fn.ReturnTypes = returnTypes
	}

	fn.Body, tokens = parseBlock(tokens)
	return fn, tokens
}

func parse_top_level_function_declaration(tokens []Token) (FunctionDeclaration, []Token) {
	var fn FunctionDeclaration
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
		var typ Type
		typ, tokens = parseType(tokens)
		param.Type = &typ
		fn.Params = append(fn.Params, param)

		if peekToken(tokens, RParen) {
			break
		}

		_, tokens = consumeToken(tokens, Comma)
	}

	_, tokens = consumeToken(tokens, RParen)

	if !peekToken(tokens, LCurly) {
		var returnTypes []Type
		returnTypes, tokens = parseReturnTypes(tokens)
		fn.ReturnTypes = returnTypes
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

// TODO: lvalues can be like dot accesses or array accesses
func getLValues(tokens []Token) ([]string, []Token, TokenType) {
	origTokens := tokens
	var lValues []string
	var thisToken Token
	for len(tokens) > 0 {
		thisLValueName := ""
		for peekToken(tokens, Ident, Dot) {
			thisToken, tokens = consumeToken(tokens, Ident)
			thisLValueName += thisToken.Value
			thisToken, tokens = consumeToken(tokens, Dot)
			thisLValueName += "."
			continue
		}
		if peekToken(tokens, Ident, Comma) {
			thisToken, tokens = consumeToken(tokens, Ident)
			thisLValueName += thisToken.Value
			lValues = append(lValues, thisLValueName)
			thisToken, tokens = consumeToken(tokens, Comma)
			continue
		} else if peekToken(tokens, Ident, Assignment) {
			thisToken, tokens = consumeToken(tokens, Ident)
			thisLValueName += thisToken.Value
			lValues = append(lValues, thisLValueName)
			return lValues, tokens, Assignment
		} else if peekToken(tokens, Ident, Reassignment) {
			thisToken, tokens = consumeToken(tokens, Ident)
			thisLValueName += thisToken.Value
			lValues = append(lValues, thisLValueName)
			return lValues, tokens, Reassignment
		} else {
			return nil, origTokens, Assignment // last value doesn't matter, it's an error
		}
	}
	panic("unreachable")
}

func parseStatement(tokens []Token) (Statement, []Token) {
	var lValues []string
	var assignType TokenType
	lValues, tokens, assignType = getLValues(tokens)
	println("assignType", formatTokenType(assignType))
	println(fmt.Sprintf("lValues: %#v", lValues))
	if len(lValues) > 0 && assignType == Assignment {
		return parseAssignment(tokens, lValues)
	} else if len(lValues) > 0 && assignType == Reassignment {
		return parseReassignment(tokens, lValues)
	} else if len(lValues) > 0 {
		panic("wrong type of assignment??????")
	}
	// else must be an expr
	println("parsing as expr instead")
	return parseExpr(tokens)
}

func parseAssignment(tokens []Token, lValues []string) (*AssignmentStmt, []Token) {
	var stmt AssignmentStmt
	stmt.VarNames = lValues
	_, tokens = consumeToken(tokens, Assignment)
	stmt.Expr, tokens = parseExpr(tokens)
	return &stmt, tokens
}

func parseReassignment(tokens []Token, lValues []string) (*ReassignmentStmt, []Token) {
	var stmt ReassignmentStmt
	stmt.VarNames = lValues
	_, tokens = consumeToken(tokens, Reassignment)
	stmt.Expr, tokens = parseExpr(tokens)
	return &stmt, tokens
}

type ParseExprOpts struct {
	SkipMatch bool
}

func consumeBinaryOperator(leftNode Expr, tokens []Token) (*BinaryOpExpr, []Token) {
	var thisToken Token
	var expr BinaryOpExpr
	expr.Left = leftNode
	thisToken, tokens = consumeToken(tokens, BinaryOp)
	expr.Op = thisToken.Value
	expr.Right, tokens = parseExpr(tokens)
	return &expr, tokens
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
				return &StringLiteralExpr{Value: value}, tokens[1:]
			}
		}

		maybeIntLiteral, tokens = tryParseIntLiteral(tokens)
		if maybeIntLiteral != nil {
			return maybeIntLiteral, tokens
		}
		maybeMatchStmt, tokens := tryParseMatchStmt(tokens)
		if maybeMatchStmt != nil {
			return maybeMatchStmt, tokens
		}

		if tokens[0].Type == LCurly {
			block, tokens := parseBlock(tokens)
			return &block, tokens
		}

		// slices means we're definitely gonna see a type coming, for now
		if peekToken(tokens, LBracket, RBracket) {
			type_, tokens := parseType(tokens)
			return &type_, tokens
		}

		if tokens[0].Type == While {
			while_, tokens := parseWhile(tokens)
			return &while_, tokens
		}

		if tokens[0].Type == Break {
			break_, tokens := parseBreak(tokens)
			return &break_, tokens
		}

		if tokens[0].Type == Continue {
			continue_, tokens := parseContinue(tokens)
			return &continue_, tokens
		}

		if tokens[0].Type == If {
			if_, tokens := parseIf(tokens)
			return &if_, tokens
		}

		if tokens[0].Type == Return {
			return_, tokens := parseReturn(tokens)
			return &return_, tokens
		}

		if tokens[0].Type == FuncDecl {
			f, tokens := parseAnonFuncDecl(tokens)
			return &f, tokens
		}

		// therefore, must be a var reference
		thisToken, tokens = consumeToken(tokens, Ident)
		varRef.VarName = thisToken.Value
		return &varRef, tokens
	}
	node, tokens := old()

post:
	// handle period
	if peekToken(tokens, Dot) {
		_, tokens = consumeToken(tokens, Dot)
		var thisToken Token
		thisToken, tokens = consumeToken(tokens, Ident)
		node = &DotAccessExpr{Left: node, Right: thisToken.Value}
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

func consumeArrayAccess(node Expr, tokens []Token) (*ArrayAccess, []Token) {
	var ac ArrayAccess
	ac.Left = node
	_, tokens = consumeToken(tokens, LBracket)
	ac.Right, tokens = parseExpr(tokens)
	_, tokens = consumeToken(tokens, RBracket)
	return &ac, tokens
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

func parseBreak(tokens []Token) (BreakExpr, []Token) {
	_, tokens = consumeToken(tokens, Break)
	return BreakExpr{}, tokens
}

func parseContinue(tokens []Token) (ContinueExpr, []Token) {
	_, tokens = consumeToken(tokens, Continue)
	return ContinueExpr{}, tokens
}

func parseWhile(tokens []Token) (WhileExpr, []Token) {
	_, tokens = consumeToken(tokens, While)
	var w WhileExpr
	w.Body, tokens = parseExpr(tokens)
	return w, tokens
}

func parseReturnTypes(tokens []Token) ([]Type, []Token) {
	var types []Type
	if peekToken(tokens, LParen) {
		_, tokens = consumeToken(tokens, LParen)
		for !peekToken(tokens, RParen) {
			var typ Type
			typ, tokens = parseType(tokens)
			types = append(types, typ)
			if peekToken(tokens, Comma) {
				_, tokens = consumeToken(tokens, Comma)
			}
		}
		_, tokens = consumeToken(tokens, RParen)
	} else {
		var typ Type
		typ, tokens = parseType(tokens)
		types = append(types, typ)
	}
	return types, tokens
}

func parseType(tokens []Token) (Type, []Token) {
	var tn []byte
	var typ Type

	if peekToken(tokens, BinaryOp) && tokens[0].Value == "*" {
		tn = append(tn, '*')
		_, tokens = consumeToken(tokens, BinaryOp)
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

	// is it followed by a dot?
	if peekToken(tokens, Dot) {
		tn = append(tn, '.')
		_, tokens = consumeToken(tokens, Dot)
		thisToken, tokens = consumeToken(tokens, Ident)
		tn = append(tn, []byte(thisToken.Value)...)
	}

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
	initializer.LHS = node

	if peekToken(tokens, RCurly) {
		_, tokens = consumeToken(tokens, RCurly)
		return &initializer, tokens
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
	return &initializer, tokens
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

func consumeFuncCall(expr Expr, tokens []Token) (*FuncCallExpr, []Token) {
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
	return &funcCall, tokens
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
