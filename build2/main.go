package main

import "os"
import "fmt"
import "strconv"

type TokenType interface {
	_implementsTokenType()
}
type TokenTypeFuncDecl struct {
}

func (_ TokenTypeFuncDecl) _implementsTokenType() {}

type TokenTypeLParen struct {
}

func (_ TokenTypeLParen) _implementsTokenType() {}

type TokenTypeRParen struct {
}

func (_ TokenTypeRParen) _implementsTokenType() {}

type TokenTypeLCurly struct {
}

func (_ TokenTypeLCurly) _implementsTokenType() {}

type TokenTypeRCurly struct {
}

func (_ TokenTypeRCurly) _implementsTokenType() {}

type TokenTypeLBracket struct {
}

func (_ TokenTypeLBracket) _implementsTokenType() {}

type TokenTypeRBracket struct {
}

func (_ TokenTypeRBracket) _implementsTokenType() {}

type TokenTypeIdent struct {
}

func (_ TokenTypeIdent) _implementsTokenType() {}

type TokenTypeIntLiteral struct {
}

func (_ TokenTypeIntLiteral) _implementsTokenType() {}

type TokenTypeBinaryOp struct {
}

func (_ TokenTypeBinaryOp) _implementsTokenType() {}

type TokenTypeNewline struct {
}

func (_ TokenTypeNewline) _implementsTokenType() {}

type TokenTypeComma struct {
}

func (_ TokenTypeComma) _implementsTokenType() {}

type TokenTypeStringLiteral struct {
}

func (_ TokenTypeStringLiteral) _implementsTokenType() {}

type TokenTypeEnumDecl struct {
}

func (_ TokenTypeEnumDecl) _implementsTokenType() {}

type TokenTypeMatch struct {
}

func (_ TokenTypeMatch) _implementsTokenType() {}

type TokenTypeDot struct {
}

func (_ TokenTypeDot) _implementsTokenType() {}

type TokenTypeColon struct {
}

func (_ TokenTypeColon) _implementsTokenType() {}

type TokenTypeEquals struct {
}

func (_ TokenTypeEquals) _implementsTokenType() {}

type TokenTypeImport struct {
}

func (_ TokenTypeImport) _implementsTokenType() {}

type TokenTypeStruct struct {
}

func (_ TokenTypeStruct) _implementsTokenType() {}

type TokenTypeWhile struct {
}

func (_ TokenTypeWhile) _implementsTokenType() {}

type TokenTypeIf struct {
}

func (_ TokenTypeIf) _implementsTokenType() {}

type TokenTypeReturn struct {
}

func (_ TokenTypeReturn) _implementsTokenType() {}

type TokenTypeElse struct {
}

func (_ TokenTypeElse) _implementsTokenType() {}

type TokenTypeBreak struct {
}

func (_ TokenTypeBreak) _implementsTokenType() {}

type TokenTypeContinue struct {
}

func (_ TokenTypeContinue) _implementsTokenType() {}

type Token struct {
	Type  TokenType
	Value string
}

func eqany(xs []byte, x byte) bool {
	imax := len(xs)
	i := 0
	for {
		if i == imax {
			{
				return false
			}

		}
		b := xs[i]
		if x == b {
			{
				return true
			}

		}
		i = i + 1
	}

}

func isAlphanumeric(b byte) bool {
	return isDigit(b) || isAlpha(b)
}

func peekBinaryOp(dat []byte, start int) string {
	binaryOps := []string{"+", "-", "*", "/", "%", "==", "!=", "<=", ">=", "&&", "||", "<", ">"}
	i := 0
	for {
		if i >= len(binaryOps) {
			{
				return ""
			}

		} else {
			if nc(dat, start, binaryOps[i]) {
				{
					return binaryOps[i]
				}

			}
		}
		i = i + 1
	}

}

func lex(dat []byte) []Token {
	tokens := []Token{}
	i := 0
	for {
		if i >= len(dat) {
			{
				break
			}

		} else {
			if nc(dat, i, " ") || nc(dat, i, "\t") {
				{
					i = i + 1
					continue
				}

			} else {
				if nc(dat, i, "//") {
					{
						for {
							if nc(dat, i, "\n") {
								{
									tokens = append(tokens, Token{TokenType.Newline{}, "\n"})
									break
								}

							} else {
								{
									i = i + 1
								}

							}
						}

					}

				} else {
					if isAlpha(dat[i]) {
						{
							ident := []byte{}
							for {
								if isAlphanumeric(dat[i]) {
									{
										ident = append(ident, dat[i])
										i = i + 1
									}

								} else {
									{
										if string(ident) == "enum" {
											{
												tokens = append(tokens, Token{TokenType.EnumDecl{}, string(ident)})
											}

										} else {
											if string(ident) == "import" {
												{
													tokens = append(tokens, Token{TokenType.Import{}, string(ident)})
												}

											} else {
												if string(ident) == "struct" {
													{
														tokens = append(tokens, Token{TokenType.Struct{}, string(ident)})
													}

												} else {
													if string(ident) == "func" {
														{
															tokens = append(tokens, Token{TokenType.FuncDecl{}, string(ident)})
														}

													} else {
														if string(ident) == "while" {
															{
																tokens = append(tokens, Token{TokenType.While{}, string(ident)})
															}

														} else {
															if string(ident) == "if" {
																{
																	tokens = append(tokens, Token{TokenType.If{}, string(ident)})
																}

															} else {
																if string(ident) == "return" {
																	{
																		tokens = append(tokens, Token{TokenType.Return{}, string(ident)})
																	}

																} else {
																	if string(ident) == "else" {
																		{
																			tokens = append(tokens, Token{TokenType.Else{}, string(ident)})
																		}

																	} else {
																		if string(ident) == "break" {
																			{
																				tokens = append(tokens, Token{TokenType.Break{}, string(ident)})
																			}

																		} else {
																			if string(ident) == "continue" {
																				{
																					tokens = append(tokens, Token{TokenType.Continue{}, string(ident)})
																				}

																			} else {
																				if string(ident) == "match" {
																					{
																						tokens = append(tokens, Token{TokenType.Match{}, string(ident)})
																					}

																				} else {
																					{
																						tokens = append(tokens, Token{TokenType.Ident{}, string(ident)})
																					}

																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
										break
									}

								}
							}

							continue
						}

					} else {
						if isDigit(dat[i]) {
							{
								n := []byte{}
								for {
									if isDigit(dat[i]) {
										{
											n = append(n, dat[i])
											i = i + 1
										}

									} else {
										{
											tokens = append(tokens, Token{TokenType.IntLiteral{}, string(n)})
											break
										}

									}
								}

							}

						} else {
							if nc(dat, i, "{") {
								{
									tokens = append(tokens, Token{TokenType.LCurly{}, "{"})
									i = i + 1
								}

							} else {
								if nc(dat, i, "}") {
									{
										tokens = append(tokens, Token{TokenType.RCurly{}, "}"})
										i = i + 1
									}

								} else {
									if nc(dat, i, "[") {
										{
											tokens = append(tokens, Token{TokenType.LBracket{}, "["})
											i = i + 1
										}

									} else {
										if nc(dat, i, "]") {
											{
												tokens = append(tokens, Token{TokenType.RBracket{}, "]"})
												i = i + 1
											}

										} else {
											if nc(dat, i, "(") {
												{
													tokens = append(tokens, Token{TokenType.LParen{}, "("})
													i = i + 1
												}

											} else {
												if nc(dat, i, ")") {
													{
														tokens = append(tokens, Token{TokenType.RParen{}, ")"})
														i = i + 1
													}

												} else {
													if nc(dat, i, "\n") {
														{
															tokens = append(tokens, Token{TokenType.Newline{}, "\n"})
															i = i + 1
														}

													} else {
														if nc(dat, i, "\r\n") {
															{
																tokens = append(tokens, Token{TokenType.Newline{}, "\r\n"})
																i = i + 2
															}

														} else {
															if nc(dat, i, ",") {
																{
																	tokens = append(tokens, Token{TokenType.Comma{}, ","})
																	i = i + 1
																}

															} else {
																if nc(dat, i, ":") {
																	{
																		tokens = append(tokens, Token{TokenType.Colon{}, ":"})
																		i = i + 1
																	}

																} else {
																	if nc(dat, i, ".") {
																		{
																			tokens = append(tokens, Token{TokenType.Dot{}, "."})
																			i = i + 1
																		}

																	} else {
																		if nc(dat, i, "\"") {
																			{
																				str := bs("\"")
																				i = i + 1
																				for {
																					if nc(dat, i, "\"") {
																						{
																							str = append(str, dat[i])
																							i = i + 1
																							break
																						}

																					} else {
																						if nc(dat, i, "\\\"") {
																							{
																								str = append(str, dat[i], dat[i+1])
																								i = i + 2
																							}

																						} else {
																							{
																								str = append(str, dat[i])
																								i = i + 1
																							}

																						}
																					}
																				}

																				tokens = append(tokens, Token{TokenType.StringLiteral{}, string(str)})
																			}

																		} else {
																			if nc(dat, i, "`") {
																				{
																					str := bs("`")
																					i = i + 1
																					for {
																						if nc(dat, i, "`") {
																							{
																								str = append(str, dat[i])
																								i = i + 1
																								break
																							}

																						} else {
																							if nc(dat, i, "\\\"") {
																								{
																									str = append(str, dat[i], dat[i+1])
																									i = i + 2
																								}

																							} else {
																								{
																									str = append(str, dat[i])
																									i = i + 1
																								}

																							}
																						}
																					}

																					tokens = append(tokens, Token{TokenType.StringLiteral{}, string(str)})
																				}

																			} else {
																				{
																					str := peekBinaryOp(dat, i)
																					if len(str) > 0 {
																						{
																							tokens = append(tokens, Token{TokenType.BinaryOp{}, str})
																							i = i + len(str)
																						}

																					} else {
																						if nc(dat, i, "=") {
																							{
																								tokens = append(tokens, Token{TokenType.Equals{}, "="})
																								i = i + 1
																							}

																						} else {
																							{
																								panic(sprintf("unexpected token: |%c|\n\ntokens so far: %#v", dat[i], tokens))
																							}

																						}
																					}
																				}

																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return tokens
}

type Declaration interface {
	_implementsDeclaration()
}
type DeclarationEnum struct {
	Value Enum
}

func (_ DeclarationEnum) _implementsDeclaration() {}

type DeclarationImport struct {
	Value Import
}

func (_ DeclarationImport) _implementsDeclaration() {}

type DeclarationStruct struct {
	Value Struct
}

func (_ DeclarationStruct) _implementsDeclaration() {}

type DeclarationFunction struct {
	Value Function
}

func (_ DeclarationFunction) _implementsDeclaration() {}

type FunctionParam struct {
	Name string
	Type string
}

type Function struct {
	Name        string
	Params      []FunctionParam
	ReturnTypes []string
	Body        Block
}

type LValue interface {
	_implementsLValue()
}
type LValueVariable struct {
	Value string
}

func (_ LValueVariable) _implementsLValue() {}

type LValueDot struct {
	Value DotLValue
}

func (_ LValueDot) _implementsLValue() {}

type DotLValue struct {
	LHS LValue
	RHS LValue
}

type FuncCall struct {
	LHS    Expr
	Params []Expr
}

type Expr interface {
	_implementsExpr()
}
type ExprVarRef struct {
	Value string
}

func (_ ExprVarRef) _implementsExpr() {}

type ExprFuncCall struct {
	Value FuncCall
}

func (_ ExprFuncCall) _implementsExpr() {}

type ExprIntLiteral struct {
	Value int
}

func (_ ExprIntLiteral) _implementsExpr() {}

type ExprBinOp struct {
	Value BinOp
}

func (_ ExprBinOp) _implementsExpr() {}

type ExprBlock struct {
	Value Block
}

func (_ ExprBlock) _implementsExpr() {}

type ExprArrayAccess struct {
	Value ArrayAccess
}

func (_ ExprArrayAccess) _implementsExpr() {}

type ExprInitializer struct {
	Value Initializer
}

func (_ ExprInitializer) _implementsExpr() {}

type ExprStringLiteral struct {
	Value string
}

func (_ ExprStringLiteral) _implementsExpr() {}

type ExprDotAccess struct {
	Value DotAccess
}

func (_ ExprDotAccess) _implementsExpr() {}

type DotAccess struct {
	LHS   Expr
	Field string
}

type Initializer struct {
	Type   string
	Params []Expr
}

type ArrayAccess struct {
	LHS   Expr
	Index Expr
}

type BinOp struct {
	LHS Expr
	RHS Expr
	Op  string
}

type Assignment struct {
	LValues        []LValue
	RValue         Expr
	IsReassignment bool
}

type If struct {
	Cond     Expr
	IfBody   Expr
	ElseBody Statement
}

func parseReturn(tokens []Token) (Return, []Token) {
	t := Token{}
	t, tokens = consumeToken(tokens, TokenType.Return{})
	r := Return{}
	e := sentinelExprJank()
	e, tokens = parseExpr(tokens)
	r.Exprs = append(r.Exprs, e)
	for {
		if not(peekToken(tokens, TokenType.Comma{})) {
			{
				break
			}

		}
		t, tokens = consumeToken(tokens, TokenType.Comma{})
		e, tokens = parseExpr(tokens)
		r.Exprs = append(r.Exprs, e)
	}

	return r, tokens
}

func parseStatement(tokens []Token) (Statement, []Token) {
	if peekToken(tokens, TokenType.Continue{}) {
		{
			t, tokens := consumeToken(tokens, TokenType.Continue{})
			return Statement.Continue{}, tokens
		}

	} else {
		if peekToken(tokens, TokenType.Break{}) {
			{
				t, tokens := consumeToken(tokens, TokenType.Break{})
				return Statement.Break{}, tokens
			}

		} else {
			if peekToken(tokens, TokenType.While{}) {
				{
					w, tokens := parseWhile(tokens)
					return Statement.While{w}, tokens
				}

			} else {
				if peekToken(tokens, TokenType.If{}) {
					{
						i, tokens := parseIf(tokens)
						return Statement.If{i}, tokens
					}

				} else {
					if peekToken(tokens, TokenType.Return{}) {
						{
							r, tokens := parseReturn(tokens)
							return Statement.Return{r}, tokens
						}

					} else {
						if peekToken(tokens, TokenType.Match{}) {
							{
								m, tokens := parseMatch(tokens)
								return Statement.Match{m}, tokens
							}

						}
					}
				}
			}
		}
	}
	isAssignment := false
	ass := Assignment{}
	isAssignment, ass, tokens = tryParseAssignment(tokens)
	if isAssignment {
		{
			return Statement.Assignment{ass}, tokens
		}

	}
	expr, tokens := parseExpr(tokens)
	return Statement.Expr{expr}, tokens
}

func sentinelLValueJank() LValue {
	return LValue.Variable{""}
}

func identsToLValue(idents []string) LValue {
	if len(idents) == 0 {
		{
			panic("requesting idents to lvalue for no idents")
		}

	} else {
		if len(idents) == 1 {
			{
				return LValue.Variable{idents[0]}
			}

		}
	}
	l := LValue.Variable{idents[0]}
	r := LValue.Variable{idents[1]}
	acc := sentinelLValueJank()
	acc = LValue.Dot{DotLValue{l, r}}
	idents = slice(idents, 2)
	for {
		if len(idents) == 0 {
			{
				return acc
			}

		} else {
			if len(idents) == 1 {
				{
					dlv := DotLValue{acc, LValue.Variable{idents[0]}}
					return LValue.Dot{dlv}
				}

			} else {
				{
					l := LValue.Variable{idents[0]}
					r := LValue.Variable{idents[1]}
					dlv := DotLValue{l, r}
					acc = LValue.Dot{DotLValue{acc, LValue.Dot{dlv}}}
					idents = slice(idents, 2)
				}

			}
		}
	}

}

func tryParseAssignment(tokens []Token) (bool, Assignment, []Token) {
	origTokens := tokens
	t := Token{}
	ass := Assignment{}
	ass.IsReassignment = false
	collectedIdents := []string{}
	for {
		if peekTokens(tokens, []TokenType{TokenType.Ident{}, TokenType.Equals{}}) {
			{
				ass.IsReassignment = true
				t, tokens = consumeToken(tokens, TokenType.Ident{})
				collectedIdents = append(collectedIdents, t.Value)
				ass.LValues = append(ass.LValues, identsToLValue(collectedIdents))
				collectedIdents = []string{}
				t, tokens = consumeToken(tokens, TokenType.Equals{})
				break
			}

		} else {
			if peekTokens(tokens, []TokenType{TokenType.Ident{}, TokenType.Colon{}, TokenType.Equals{}}) {
				{
					ass.IsReassignment = false
					t, tokens = consumeToken(tokens, TokenType.Ident{})
					collectedIdents = append(collectedIdents, t.Value)
					ass.LValues = append(ass.LValues, identsToLValue(collectedIdents))
					collectedIdents = []string{}
					t, tokens = consumeToken(tokens, TokenType.Colon{})
					t, tokens = consumeToken(tokens, TokenType.Equals{})
					break
				}

			} else {
				if peekTokens(tokens, []TokenType{TokenType.Ident{}, TokenType.Comma{}}) {
					{
						t, tokens = consumeToken(tokens, TokenType.Ident{})
						collectedIdents = append(collectedIdents, t.Value)
						ass.LValues = append(ass.LValues, identsToLValue(collectedIdents))
						collectedIdents = []string{}
						t, tokens = consumeToken(tokens, TokenType.Comma{})
					}

				} else {
					if peekTokens(tokens, []TokenType{TokenType.Ident{}, TokenType.Dot{}}) {
						{
							t, tokens = consumeToken(tokens, TokenType.Ident{})
							collectedIdents = append(collectedIdents, t.Value)
							t, tokens = consumeToken(tokens, TokenType.Dot{})
							continue
						}

					} else {
						{
							return false, Assignment{}, origTokens
						}

					}
				}
			}
		}
	}

	ass.RValue, tokens = parseExpr(tokens)
	return true, ass, tokens
}

func sentinelExprJank() Expr {
	return Expr.VarRef{""}
}

func parseInitializer(tokens []Token, typ string) (Initializer, []Token) {
	t := Token{}
	i := Initializer{}
	i.Type = typ
	t, tokens = consumeToken(tokens, TokenType.LCurly{})
	for {
		if peekToken(tokens, TokenType.RCurly{}) {
			{
				break
			}

		}
		arg, tokens2 := parseExpr(tokens)
		tokens = tokens2
		i.Params = append(i.Params, arg)
		if peekToken(tokens, TokenType.Comma{}) {
			{
				t, tokens = consumeToken(tokens, TokenType.Comma{})
			}

		} else {
			{
				break
			}

		}
	}

	t, tokens = consumeToken(tokens, TokenType.RCurly{})
	return i, tokens
}

func parseExpr(tokens []Token) (Expr, []Token) {
	t := Token{}
	expr := sentinelExprJank()
	if peekToken(tokens, TokenType.LBracket{}) {
		{
			i := Initializer{}
			typ := ""
			typ, tokens = parseType(tokens)
			i, tokens = parseInitializer(tokens, typ)
			expr = Expr.Initializer{i}
		}

	} else {
		if peekToken(tokens, TokenType.StringLiteral{}) {
			{
				t, tokens = consumeToken(tokens, TokenType.StringLiteral{})
				expr = Expr.StringLiteral{t.Value}
			}

		} else {
			if peekToken(tokens, TokenType.LCurly{}) {
				{
					b := Block{}
					b, tokens = parseBlock(tokens)
					expr = Expr.Block{b}
				}

			} else {
				if peekToken(tokens, TokenType.Ident{}) {
					{
						t, tokens = consumeToken(tokens, TokenType.Ident{})
						expr = Expr.VarRef{t.Value}
					}

				} else {
					if peekToken(tokens, TokenType.IntLiteral{}) {
						{
							t, tokens = consumeToken(tokens, TokenType.IntLiteral{})
							expr = Expr.IntLiteral{atoi(t.Value)}
						}

					} else {
						{
							panic(fmt.Sprintf("unrecognized token when parsing expr: %#v", tokens[0]))
						}

					}
				}
			}
		}
	}
	for {
		if peekToken(tokens, TokenType.LParen{}) {
			{
				funcCall := FuncCall{}
				funcCall, tokens = parseFuncCall(tokens, expr)
				expr = Expr.FuncCall{funcCall}
				continue
			}

		} else {
			if peekToken(tokens, TokenType.BinaryOp{}) {
				{
					t, tokens = consumeToken(tokens, TokenType.BinaryOp{})
					binop := BinOp{}
					binop.LHS = expr
					binop.Op = t.Value
					binop.RHS, tokens = parseExpr(tokens)
					expr = Expr.BinOp{binop}
					continue
				}

			} else {
				if peekToken(tokens, TokenType.LBracket{}) {
					{
						t, tokens = consumeToken(tokens, TokenType.LBracket{})
						arrayAccess := ArrayAccess{}
						arrayAccess.LHS = expr
						arrayAccess.Index, tokens = parseExpr(tokens)
						t, tokens = consumeToken(tokens, TokenType.RBracket{})
						expr = Expr.ArrayAccess{arrayAccess}
						continue
					}

				} else {
					if peekToken(tokens, TokenType.LCurly{}) {
						{
							typ := exprToType(expr)
							i := Initializer{}
							i, tokens = parseInitializer(tokens, typ)
							expr = Expr.Initializer{i}
						}

					} else {
						if peekToken(tokens, TokenType.Dot{}) {
							{
								dotAccess := DotAccess{}
								dotAccess.LHS = expr
								t, tokens = consumeToken(tokens, TokenType.Dot{})
								t, tokens = consumeToken(tokens, TokenType.Ident{})
								dotAccess.Field = t.Value
								expr = Expr.DotAccess{dotAccess}
							}

						} else {
							{
								break
							}

						}
					}
				}
			}
		}
	}

	return expr, tokens
}

func exprToType(expr Expr) string {
	{
		matchExpr := expr
		if binding, ok := matchExpr.(Expr.VarRef); ok {
			v := binding.Value
			{
				return v
			}

		} else if binding, ok := matchExpr.(Expr.DotAccess); ok {
			dotAccess := binding.Value
			{
				l := dotAccess.LHS
				f := dotAccess.Field
				t := exprToType(l)
				return t + "." + f
			}

		}
	}
	panic(fmt.Sprintf("trying to convert expr to type, unhandled expr: %#v", expr))
}

func parseFuncCall(tokens []Token, lhs Expr) (FuncCall, []Token) {
	t := Token{}
	t, tokens = consumeToken(tokens, TokenType.LParen{})
	funcCall := FuncCall{}
	funcCall.LHS = lhs
	for {
		if peekToken(tokens, TokenType.RParen{}) {
			{
				break
			}

		}
		nextParam, tokens2 := parseExpr(tokens)
		funcCall.Params = append(funcCall.Params, nextParam)
		tokens = tokens2
		if peekToken(tokens, TokenType.Comma{}) {
			{
				t, tokens = consumeToken(tokens, TokenType.Comma{})
			}

		} else {
			{
				break
			}

		}
	}

	t, tokens = consumeToken(tokens, TokenType.RParen{})
	return funcCall, tokens
}

type Statement interface {
	_implementsStatement()
}
type StatementAssignment struct {
	Value Assignment
}

func (_ StatementAssignment) _implementsStatement() {}

type StatementWhile struct {
	Value While
}

func (_ StatementWhile) _implementsStatement() {}

type StatementIf struct {
	Value If
}

func (_ StatementIf) _implementsStatement() {}

type StatementReturn struct {
	Value Return
}

func (_ StatementReturn) _implementsStatement() {}

type StatementBreak struct {
}

func (_ StatementBreak) _implementsStatement() {}

type StatementContinue struct {
}

func (_ StatementContinue) _implementsStatement() {}

type StatementExpr struct {
	Value Expr
}

func (_ StatementExpr) _implementsStatement() {}

type StatementMatch struct {
	Value Match
}

func (_ StatementMatch) _implementsStatement() {}

type Match struct {
	Matched Expr
	Arms    []MatchArm
}

type MatchArm struct {
	Pattern MatchPattern
	Body    Expr
}

type MatchPattern interface {
	_implementsMatchPattern()
}
type MatchPatternEnum struct {
	Value EnumMatchPattern
}

func (_ MatchPatternEnum) _implementsMatchPattern() {}

type EnumMatchPattern struct {
	Type    string
	Binding string
}

func parseMatchPattern(tokens []Token) (MatchPattern, []Token) {
	emp := EnumMatchPattern{}
	emp.Type, tokens = parseType(tokens)
	t := Token{}
	t, tokens = consumeToken(tokens, TokenType.LCurly{})
	if peekToken(tokens, TokenType.Ident{}) {
		{
			t, tokens = consumeToken(tokens, TokenType.Ident{})
			emp.Binding = t.Value
		}

	}
	t, tokens = consumeToken(tokens, TokenType.RCurly{})
	return MatchPattern.Enum{emp}, tokens
}

func parseMatch(tokens []Token) (Match, []Token) {
	t := Token{}
	t, tokens = consumeToken(tokens, TokenType.Match{})
	m := Match{}
	t, tokens = consumeToken(tokens, TokenType.LParen{})
	m.Matched, tokens = parseExpr(tokens)
	t, tokens = consumeToken(tokens, TokenType.RParen{})
	t, tokens = consumeToken(tokens, TokenType.LCurly{})
	for {
		if peekToken(tokens, TokenType.RCurly{}) {
			{
				break
			}

		}
		arm := MatchArm{}
		arm.Pattern, tokens = parseMatchPattern(tokens)
		t, tokens = consumeToken(tokens, TokenType.Colon{})
		arm.Body, tokens = parseExpr(tokens)
		m.Arms = append(m.Arms, arm)
		if not(peekToken(tokens, TokenType.Comma{})) {
			{
				break
			}

		}
		t, tokens = consumeToken(tokens, TokenType.Comma{})
	}

	t, tokens = consumeToken(tokens, TokenType.RCurly{})
	return m, tokens
}

type Return struct {
	Exprs []Expr
}

type Block struct {
	Statements []Statement
}

type While struct {
	Body Block
}

func parseIf(tokens []Token) (If, []Token) {
	t := Token{}
	t, tokens = consumeToken(tokens, TokenType.If{})
	i := If{}
	t, tokens = consumeToken(tokens, TokenType.LParen{})
	i.Cond, tokens = parseExpr(tokens)
	t, tokens = consumeToken(tokens, TokenType.RParen{})
	i.IfBody, tokens = parseExpr(tokens)
	if peekToken(tokens, TokenType.Else{}) {
		{
			t, tokens = consumeToken(tokens, TokenType.Else{})
			i.ElseBody, tokens = parseStatement(tokens)
		}

	}
	return i, tokens
}

func parseWhile(tokens []Token) (While, []Token) {
	t := Token{}
	t, tokens = consumeToken(tokens, TokenType.While{})
	w := While{}
	w.Body, tokens = parseBlock(tokens)
	return w, tokens
}

func parseBlock(tokens []Token) (Block, []Token) {
	stmts := []Statement{}
	t := Token{}
	t, tokens = consumeToken(tokens, TokenType.LCurly{})
	for {
		tokens = skipNewlines(tokens)
		stmt, tokens2 := parseStatement(tokens)
		tokens = tokens2
		stmts = append(stmts, stmt)
		if peekToken(tokens, TokenType.RCurly{}) {
			{
				break
			}

		}
	}

	t, tokens = consumeToken(tokens, TokenType.RCurly{})
	return Block{stmts}, tokens
}

type Program struct {
	Declarations []Declaration
}

type Import struct {
	Path string
}

type Enum struct {
	Name     string
	Variants []EnumVariant
}

type EnumVariant struct {
	Name string
	Type string
}

type Struct struct {
	Name   string
	Fields []StructField
}

type StructField struct {
	Name string
	Type string
}

func parseImport(tokens []Token) (Import, []Token) {
	t, tokens := consumeToken(tokens, TokenType.Import{})
	t, tokens = consumeToken(tokens, TokenType.StringLiteral{})
	name := t.Value
	return Import{name}, tokens
}

func parseEnum(tokens []Token) (Enum, []Token) {
	t, tokens := consumeToken(tokens, TokenType.EnumDecl{})
	t, tokens = consumeToken(tokens, TokenType.Ident{})
	e := Enum{}
	e.Name = t.Value
	t, tokens = consumeToken(tokens, TokenType.LCurly{})
	for {
		if peekToken(tokens, TokenType.RCurly{}) {
			{
				t, tokens = consumeToken(tokens, TokenType.RCurly{})
				break
			}

		}
		t, tokens = consumeToken(tokens, TokenType.Ident{})
		variant := EnumVariant{}
		variant.Name = t.Value
		if peekToken(tokens, TokenType.LParen{}) {
			{
				t, tokens = consumeToken(tokens, TokenType.LParen{})
				typ := ""
				typ, tokens = parseType(tokens)
				variant.Type = typ
				t, tokens = consumeToken(tokens, TokenType.RParen{})
			}

		}
		e.Variants = append(e.Variants, variant)
		t, tokens = consumeToken(tokens, TokenType.Comma{})
	}

	return e, tokens
}

func parseStruct(tokens []Token) (Struct, []Token) {
	t, tokens := consumeToken(tokens, TokenType.Struct{})
	t, tokens = consumeToken(tokens, TokenType.Ident{})
	s := Struct{}
	s.Name = t.Value
	t, tokens = consumeToken(tokens, TokenType.LCurly{})
	for {
		if peekToken(tokens, TokenType.RCurly{}) {
			{
				t, tokens = consumeToken(tokens, TokenType.RCurly{})
				break
			}

		}
		field := StructField{}
		t, tokens = consumeToken(tokens, TokenType.Ident{})
		field.Name = t.Value
		field.Type, tokens = parseType(tokens)
		s.Fields = append(s.Fields, field)
		t, tokens = consumeToken(tokens, TokenType.Comma{})
	}

	return s, tokens
}

func parseType(tokens []Token) (string, []Token) {
	name := ""
	t := Token{}
	for {
		if peekToken(tokens, TokenType.Ident{}) {
			{
				t, tokens = consumeToken(tokens, TokenType.Ident{})
				name = name + t.Value
			}

		} else {
			if peekToken(tokens, TokenType.LBracket{}) {
				{
					t, tokens = consumeToken(tokens, TokenType.LBracket{})
					name = name + t.Value
				}

			} else {
				if peekToken(tokens, TokenType.RBracket{}) {
					{
						t, tokens = consumeToken(tokens, TokenType.RBracket{})
						name = name + t.Value
					}

				} else {
					if peekToken(tokens, TokenType.Dot{}) {
						{
							t, tokens = consumeToken(tokens, TokenType.Dot{})
							name = name + t.Value
						}

					} else {
						{
							break
						}

					}
				}
			}
		}
	}

	return name, tokens
}

func parseFunction(tokens []Token) (Function, []Token) {
	t, tokens := consumeToken(tokens, TokenType.FuncDecl{})
	t, tokens = consumeToken(tokens, TokenType.Ident{})
	f := Function{}
	f.Name = t.Value
	t, tokens = consumeToken(tokens, TokenType.LParen{})
	for {
		if peekToken(tokens, TokenType.RParen{}) {
			{
				t, tokens = consumeToken(tokens, TokenType.RParen{})
				break
			}

		}
		param := FunctionParam{}
		t, tokens = consumeToken(tokens, TokenType.Ident{})
		param.Name = t.Value
		param.Type, tokens = parseType(tokens)
		f.Params = append(f.Params, param)
		if peekToken(tokens, TokenType.Comma{}) {
			{
				t, tokens = consumeToken(tokens, TokenType.Comma{})
			}

		}
	}

	if not(peekToken(tokens, TokenType.LCurly{})) {
		{
			typ := ""
			f.ReturnTypes, tokens = parseReturnTypes(tokens)
		}

	}
	f.Body, tokens = parseBlock(tokens)
	return f, tokens
}

func parseReturnTypes(tokens []Token) ([]string, []Token) {
	types := []string{}
	typ := ""
	t := Token{}
	if peekToken(tokens, TokenType.LParen{}) {
		{
			t, tokens = consumeToken(tokens, TokenType.LParen{})
			for {
				typ, tokens = parseType(tokens)
				types = append(types, typ)
				if peekToken(tokens, TokenType.Comma{}) {
					{
						t, tokens = consumeToken(tokens, TokenType.Comma{})
					}

				} else {
					{
						break
					}

				}
			}

			t, tokens = consumeToken(tokens, TokenType.RParen{})
		}

	} else {
		{
			typ, tokens = parseType(tokens)
			types = append(types, typ)
		}

	}
	return types, tokens
}

func parseDeclaration(tokens []Token) (Declaration, []Token) {
	{
		matchExpr := tokens[0].Type
		if binding, ok := matchExpr.(TokenType.Import); ok {
			{
				imp, tokens := parseImport(tokens)
				return Declaration.Import{imp}, tokens
			}

		} else if binding, ok := matchExpr.(TokenType.EnumDecl); ok {
			{
				e, tokens := parseEnum(tokens)
				return Declaration.Enum{e}, tokens
			}

		} else if binding, ok := matchExpr.(TokenType.Struct); ok {
			{
				s, tokens := parseStruct(tokens)
				return Declaration.Struct{s}, tokens
			}

		} else if binding, ok := matchExpr.(TokenType.FuncDecl); ok {
			{
				f, tokens := parseFunction(tokens)
				return Declaration.Function{f}, tokens
			}

		}
	}
	panic(fmt.Sprintf("unexpected token: %#v", tokens[0]))
}

func parseProgram(tokens []Token) Program {
	p := Program{}
	for {
		tokens = skipNewlines(tokens)
		if len(tokens) == 0 {
			{
				break
			}

		}
		declaration, tokens2 := parseDeclaration(tokens)
		tokens = tokens2
		p.Declarations = append(p.Declarations, declaration)
	}

	return p
}

func skipNewlines(tokens []Token) []Token {
	for {
		if len(tokens) == 0 {
			{
				return tokens
			}

		}
		{
			matchExpr := tokens[0].Type
			if binding, ok := matchExpr.(TokenType.Newline); ok {
				{
					tokens = slice(tokens, 1)
					continue
				}

			}
		}
		break
	}

	return tokens
}

func peekTokens(tokens []Token, expectedTypes []TokenType) bool {
	i := 0
	for {
		if i >= len(expectedTypes) {
			{
				return true
			}

		}
		if len(tokens) == 0 {
			{
				return false
			}

		}
		if tokens[i].Type != expectedTypes[i] {
			{
				return false
			}

		}
		i = i + 1
	}

}

func peekToken(tokens []Token, expectedType TokenType) bool {
	if len(tokens) == 0 {
		{
			panic("Unexpected end of input")
		}

	}
	nl := TokenType.Newline{}
	if expectedType != nl {
		{
			tokens = skipNewlines(tokens)
		}

	}
	return tokens[0].Type == expectedType
}

func consumeToken(tokens []Token, expectedType TokenType) (Token, []Token) {
	if len(tokens) == 0 {
		{
			panic("Unexpected end of input")
		}

	}
	nl := TokenType.Newline{}
	if expectedType != nl {
		{
			tokens = skipNewlines(tokens)
		}

	}
	if tokens[0].Type != expectedType {
		{
			panic(fmt.Sprintf("Was expecting token %#v, got %#v", expectedType, tokens[0]))
		}

	}
	return tokens[0], slice(tokens, 1)
}

func compile(program Program) string {
	s := "package main\n\n"
	i := 0
	for {
		if i >= len(program.Declarations) {
			{
				break
			}

		}
		d := program.Declarations[i]
		compiled := compileDeclaration(d)
		s = s + compiled
		s = s + "\n"
		i = i + 1
	}

	return s
}

func compileDeclaration(d Declaration) string {
	{
		matchExpr := d
		if binding, ok := matchExpr.(Declaration.Enum); ok {
			e := binding.Value
			{
				return compileEnum(e)
			}

		} else if binding, ok := matchExpr.(Declaration.Import); ok {
			imp := binding.Value
			{
				return compileImport(imp)
			}

		} else if binding, ok := matchExpr.(Declaration.Struct); ok {
			s := binding.Value
			{
				return compileStruct(s)
			}

		} else if binding, ok := matchExpr.(Declaration.Function); ok {
			f := binding.Value
			{
				return compileFunction(f)
			}

		}
	}
	panic("unreachable")
}

func golangInterfaceName(e Enum) string {
	return e.Name
}

func golangEnumImplementsInterfaceMethodName(e Enum) string {
	return "_implements" + golangInterfaceName(e)
}

func compileEnum(e Enum) string {
	return compileEnumInterface(e) + compileEnumStructs(e)
}

func compileImport(imp Import) string {
	return "import " + imp.Path
}

func compileStruct(strukt Struct) string {
	s := "type " + strukt.Name + " struct {\n"
	i := 0
	for {
		if i >= len(strukt.Fields) {
			{
				break
			}

		}
		field := strukt.Fields[i]
		s = s + field.Name + " " + field.Type + "\n"
		i = i + 1
	}

	s = s + "}\n"
	return s
}

func compileFunction(f Function) string {
	s := "func " + f.Name + "("
	i := 0
	for {
		if i >= len(f.Params) {
			{
				break
			}

		}
		param := f.Params[i]
		s = s + param.Name + " " + param.Type
		if i < len(f.Params)-1 {
			{
				s = s + ", "
			}

		}
		i = i + 1
	}

	returnTypes := f.ReturnTypes
	body := f.Body
	s = s + ") " + compileReturnTypes(returnTypes) + compileBlock(body)
	return s
}

func compileBlock(b Block) string {
	s := "{\n"
	i := 0
	for {
		if i >= len(b.Statements) {
			{
				break
			}

		}
		statement := b.Statements[i]
		s = s + compileStatement(statement)
		s = s + "\n"
		i = i + 1
	}

	s = s + "}\n"
	return s
}

func compileStatement(stmt Statement) string {
	{
		matchExpr := stmt
		if binding, ok := matchExpr.(Statement.Assignment); ok {
			assignment := binding.Value
			{
				return compileAssignment(assignment)
			}

		} else if binding, ok := matchExpr.(Statement.While); ok {
			w := binding.Value
			{
				return compileWhile(w)
			}

		} else if binding, ok := matchExpr.(Statement.If); ok {
			ifStmt := binding.Value
			{
				return compileIf(ifStmt)
			}

		} else if binding, ok := matchExpr.(Statement.Return); ok {
			r := binding.Value
			{
				return compileReturn(r)
			}

		} else if binding, ok := matchExpr.(Statement.Break); ok {
			{
				return "break"
			}

		} else if binding, ok := matchExpr.(Statement.Continue); ok {
			{
				return "continue"
			}

		} else if binding, ok := matchExpr.(Statement.Expr); ok {
			expr := binding.Value
			{
				return compileExpr(expr)
			}

		} else if binding, ok := matchExpr.(Statement.Match); ok {
			m := binding.Value
			{
				return compileMatch(m)
			}

		}
	}
	panic(fmt.Sprintf("unreachable, must not have handled a type of statement: %#v", stmt))
}

func MatchExprVarName() string {
	return "matchExpr"
}

func BindingVarName() string {
	return "binding"
}

func compileMatch(m Match) string {
	s := "{\n"
	matched := m.Matched
	s = s + MatchExprVarName() + " := " + compileExpr(matched) + "\n"
	i := 0
	arms := m.Arms
	for {
		if i >= len(arms) {
			{
				break
			}

		}
		arm := arms[i]
		if i == 0 {
			{
				s = s + "if "
			}

		} else {
			{
				s = s + " else if "
			}

		}
		pat := arm.Pattern
		s = s + compileMatchPatternTestExpr(pat) + " {\n"
		s = s + compileBindingForMatchPattern(pat)
		body := arm.Body
		s = s + compileExpr(body)
		s = s + "\n}"
		i = i + 1
	}

	s = s + "\n}"
	return s
}

func compileBindingForMatchPattern(pattern MatchPattern) string {
	{
		matchExpr := pattern
		if binding, ok := matchExpr.(MatchPattern.Enum); ok {
			emp := binding.Value
			{
				if emp.Binding == "" {
					{
						return ""
					}

				}
				return emp.Binding + " := " + BindingVarName() + ".Value\n"
			}

		}
	}
	panic(fmt.Sprintf("unreachable, must not have handled a type of match pattern: %#v", pattern))
}

func compileMatchPatternTestExpr(pattern MatchPattern) string {
	{
		matchExpr := pattern
		if binding, ok := matchExpr.(MatchPattern.Enum); ok {
			emp := binding.Value
			{
				s := BindingVarName() + ", ok := " + MatchExprVarName() + ".(" + emp.Type + ")"
				s = s + "; ok"
				return s
			}

		}
	}
	panic(fmt.Sprintf("unreachable, must not have handled a type of match pattern: %#v", pattern))
}

func compileReturn(r Return) string {
	s := "return "
	exprs := r.Exprs
	if len(exprs) == 0 {
		{
			return s
		}

	} else {
		{
			return s + compileExprsCSV(exprs)
		}

	}
}

func compileExprsCSV(exprs []Expr) string {
	s := ""
	i := 0
	for {
		if i >= len(exprs) {
			{
				break
			}

		}
		expr := exprs[i]
		s = s + compileExpr(expr)
		if i < len(exprs)-1 {
			{
				s = s + ", "
			}

		}
		i = i + 1
	}

	return s
}

func compileWhile(w While) string {
	body := w.Body
	return "for " + compileBlock(body)
}

func compileIf(ifStmt If) string {
	s := "if "
	cond := ifStmt.Cond
	s = s + compileExpr(cond)
	s = s + " {\n"
	ifBody := ifStmt.IfBody
	s = s + compileExpr(ifBody)
	s = s + "\n}"
	elseBody := ifStmt.ElseBody
	if elseBody != nil {
		{
			s = s + "else {\n"
			s = s + compileStatement(elseBody)
			s = s + "\n}"
		}

	}
	return s
}

func compileAssignment(ass Assignment) string {
	lValues := ass.LValues
	s := compileLValues(lValues)
	if ass.IsReassignment {
		{
			s = s + " = "
		}

	} else {
		{
			s = s + " := "
		}

	}
	rValue := ass.RValue
	s = s + compileExpr(rValue)
	return s
}

func compileExpr(expr Expr) string {
	{
		matchExpr := expr
		if binding, ok := matchExpr.(Expr.VarRef); ok {
			s := binding.Value
			{
				return s
			}

		} else if binding, ok := matchExpr.(Expr.FuncCall); ok {
			funcCall := binding.Value
			{
				lhs := funcCall.LHS
				s := compileExpr(lhs) + "("
				params := funcCall.Params
				s = s + compileExprsCSV(params)
				s = s + ")"
				return s
			}

		} else if binding, ok := matchExpr.(Expr.IntLiteral); ok {
			i := binding.Value
			{
				return strconv.Itoa(i)
			}

		} else if binding, ok := matchExpr.(Expr.BinOp); ok {
			binop := binding.Value
			{
				lhs := binop.LHS
				rhs := binop.RHS
				op := binop.Op
				return compileExpr(lhs) + " " + op + " " + compileExpr(rhs)
			}

		} else if binding, ok := matchExpr.(Expr.Block); ok {
			block := binding.Value
			{
				return compileBlock(block)
			}

		} else if binding, ok := matchExpr.(Expr.ArrayAccess); ok {
			aa := binding.Value
			{
				lhs := aa.LHS
				index := aa.Index
				return compileExpr(lhs) + "[" + compileExpr(index) + "]"
			}

		} else if binding, ok := matchExpr.(ExprInitializer); ok {
			init := binding.Value
			{
				s := init.Type + "{ "
				params := init.Params
				s = s + compileExprsCSV(params)
				s = s + " }"
				return s
			}

		} else if binding, ok := matchExpr.(Expr.StringLiteral); ok {
			sl := binding.Value
			{
				return sl
			}

		} else if binding, ok := matchExpr.(Expr.DotAccess); ok {
			da := binding.Value
			{
				lhs := da.LHS
				field := da.Field
				return compileExpr(lhs) + "." + field
			}

		}
	}
	panic(fmt.Sprintf("unhandled expr: %#v", expr))
}

func compileLValue(lValue LValue) string {
	{
		matchExpr := lValue
		if binding, ok := matchExpr.(LValue.Variable); ok {
			s := binding.Value
			{
				return s
			}

		} else if binding, ok := matchExpr.(LValue.Dot); ok {
			dotLValue := binding.Value
			{
				lhs := dotLValue.LHS
				rhs := dotLValue.RHS
				return compileLValue(lhs) + "." + compileLValue(rhs)
			}

		}
	}
	panic("unreachable")
}

func compileLValues(lValues []LValue) string {
	if len(lValues) == 0 {
		{
			panic("must have at least one lvalue")
		}

	} else {
		if len(lValues) == 1 {
			{
				return compileLValue(lValues[0])
			}

		} else {
			{
				s := ""
				i := 0
				for {
					if i >= len(lValues) {
						{
							break
						}

					}
					s = s + compileLValue(lValues[i])
					if i < len(lValues)-1 {
						{
							s = s + ", "
						}

					}
					i = i + 1
				}

				return s
			}

		}
	}
}

func compileReturnTypes(returnTypes []string) string {
	if len(returnTypes) == 0 {
		{
			return ""
		}

	}
	if len(returnTypes) == 1 {
		{
			return returnTypes[0]
		}

	}
	s := "("
	i := 0
	for {
		if i >= len(returnTypes) {
			{
				break
			}

		}
		s = s + returnTypes[i]
		if i < len(returnTypes)-1 {
			{
				s = s + ", "
			}

		}
		i = i + 1
	}

	s = s + ")"
	return s
}

func compileEnumInterface(e Enum) string {
	s := "type " + golangInterfaceName(e) + " interface {\n"
	s = s + golangEnumImplementsInterfaceMethodName(e) + "()\n"
	s = s + "}\n"
	return s
}

func structNameForVariant(e Enum, v EnumVariant) string {
	return e.Name + v.Name
}

func compileEnumStructs(e Enum) string {
	s := ""
	i := 0
	for {
		if i >= len(e.Variants) {
			{
				break
			}

		}
		variant := e.Variants[i]
		structName := structNameForVariant(e, variant)
		s = s + "type " + structName + " struct {\n"
		if variant.Type != "" {
			{
				s = s + "Value " + variant.Type + "\n"
			}

		}
		s = s + "}\n"
		s = s + "func (_ " + structName + ") " + golangEnumImplementsInterfaceMethodName(e) + "() {}\n"
		i = i + 1
	}

	return s
}

func main() {
	dat, err := os.ReadFile("lexer.goy")
	if err != nil {
		{
			panic(err)
		}

	}
	tokens := lex(dat)
	program := parseProgram(tokens)
	fmt.Println(compile(program))
}

func prelude() string {
	return `func slice[T any](s []T, i... int) []T {
	if len(i) == 0 {
		return s
	}
	if len(i) == 1 {
		return s[i[0]:]
	}
	if len(i) > 2 {
		panic("slice takes at most 2 arguments")
	}
    return s[i[0]:i[1]]
}

func atoi(s string) int {
	var i int
	for _, c := range s {
		i *= 10
		i += int(c - '0')
	}
	return i
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func not(b bool) bool {
	return !b
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
