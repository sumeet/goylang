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
	Loc   int
}

func newToken(tt TokenType, value string, loc int) Token {
	return Token{tt, value, loc}
}

func eqany(xs []byte, x byte) bool {
	i := 0
	_ = i

	for {
		if i == len(xs) {
			{
				return false
			}

		}
		b := xs[i]
		_ = b

		if x == b {
			{
				return true
			}

		}
		i = i + 1
		_ = i

	}

}

func isAlphanumeric(b byte) bool {
	return isDigit(b) || isAlpha(b)
}

func peekBinaryOp(dat []byte, start int) string {
	binaryOps := []string{"+", "-", "*", "/", "%", "==", "!=", "<=", ">=", "&&", "||", "<", ">"}
	_ = binaryOps

	i := 0
	_ = i

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
		_ = i

	}

}

func lex(dat []byte) []Token {
	tokens := []Token{}
	_ = tokens

	i := 0
	_ = i

	for {
		if i >= len(dat) {
			{
				break
			}

		} else {
			if nc(dat, i, " ") || nc(dat, i, "\t") {
				{
					i = i + 1
					_ = i

					continue
				}

			} else {
				if nc(dat, i, "//") {
					{
						for {
							if nc(dat, i, "\n") {
								{
									tokens = append(tokens, newToken(TokenTypeNewline{}, "\n", i))
									_ = tokens

									break
								}

							} else {
								{
									i = i + 1
									_ = i

								}

							}
						}

					}

				} else {
					if isAlpha(dat[i]) {
						{
							ident := []byte{}
							_ = ident

							for {
								if isAlphanumeric(dat[i]) {
									{
										ident = append(ident, dat[i])
										_ = ident

										i = i + 1
										_ = i

									}

								} else {
									{
										if string(ident) == "enum" {
											{
												tokens = append(tokens, newToken(TokenTypeEnumDecl{}, string(ident), i))
												_ = tokens

											}

										} else {
											if string(ident) == "import" {
												{
													tokens = append(tokens, newToken(TokenTypeImport{}, string(ident), i))
													_ = tokens

												}

											} else {
												if string(ident) == "struct" {
													{
														tokens = append(tokens, newToken(TokenTypeStruct{}, string(ident), i))
														_ = tokens

													}

												} else {
													if string(ident) == "func" {
														{
															tokens = append(tokens, newToken(TokenTypeFuncDecl{}, string(ident), i))
															_ = tokens

														}

													} else {
														if string(ident) == "while" {
															{
																tokens = append(tokens, newToken(TokenTypeWhile{}, string(ident), i))
																_ = tokens

															}

														} else {
															if string(ident) == "if" {
																{
																	tokens = append(tokens, newToken(TokenTypeIf{}, string(ident), i))
																	_ = tokens

																}

															} else {
																if string(ident) == "return" {
																	{
																		tokens = append(tokens, newToken(TokenTypeReturn{}, string(ident), i))
																		_ = tokens

																	}

																} else {
																	if string(ident) == "else" {
																		{
																			tokens = append(tokens, newToken(TokenTypeElse{}, string(ident), i))
																			_ = tokens

																		}

																	} else {
																		if string(ident) == "break" {
																			{
																				tokens = append(tokens, newToken(TokenTypeBreak{}, string(ident), i))
																				_ = tokens

																			}

																		} else {
																			if string(ident) == "continue" {
																				{
																					tokens = append(tokens, newToken(TokenTypeContinue{}, string(ident), i))
																					_ = tokens

																				}

																			} else {
																				if string(ident) == "match" {
																					{
																						tokens = append(tokens, newToken(TokenTypeMatch{}, string(ident), i))
																						_ = tokens

																					}

																				} else {
																					{
																						tokens = append(tokens, newToken(TokenTypeIdent{}, string(ident), i))
																						_ = tokens

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
								_ = n

								for {
									if isDigit(dat[i]) {
										{
											n = append(n, dat[i])
											_ = n

											i = i + 1
											_ = i

										}

									} else {
										{
											tokens = append(tokens, newToken(TokenTypeIntLiteral{}, string(n), i))
											_ = tokens

											break
										}

									}
								}

							}

						} else {
							if nc(dat, i, "{") {
								{
									tokens = append(tokens, newToken(TokenTypeLCurly{}, "{", i))
									_ = tokens

									i = i + 1
									_ = i

								}

							} else {
								if nc(dat, i, "}") {
									{
										tokens = append(tokens, newToken(TokenTypeRCurly{}, "}", i))
										_ = tokens

										i = i + 1
										_ = i

									}

								} else {
									if nc(dat, i, "[") {
										{
											tokens = append(tokens, newToken(TokenTypeLBracket{}, "[", i))
											_ = tokens

											i = i + 1
											_ = i

										}

									} else {
										if nc(dat, i, "]") {
											{
												tokens = append(tokens, newToken(TokenTypeRBracket{}, "]", i))
												_ = tokens

												i = i + 1
												_ = i

											}

										} else {
											if nc(dat, i, "(") {
												{
													tokens = append(tokens, newToken(TokenTypeLParen{}, "(", i))
													_ = tokens

													i = i + 1
													_ = i

												}

											} else {
												if nc(dat, i, ")") {
													{
														tokens = append(tokens, newToken(TokenTypeRParen{}, ")", i))
														_ = tokens

														i = i + 1
														_ = i

													}

												} else {
													if nc(dat, i, "\n") {
														{
															tokens = append(tokens, newToken(TokenTypeNewline{}, "\n", i))
															_ = tokens

															i = i + 1
															_ = i

														}

													} else {
														if nc(dat, i, "\r\n") {
															{
																tokens = append(tokens, newToken(TokenTypeNewline{}, "\r\n", i))
																_ = tokens

																i = i + 2
																_ = i

															}

														} else {
															if nc(dat, i, ",") {
																{
																	tokens = append(tokens, newToken(TokenTypeComma{}, ",", i))
																	_ = tokens

																	i = i + 1
																	_ = i

																}

															} else {
																if nc(dat, i, ":") {
																	{
																		tokens = append(tokens, newToken(TokenTypeColon{}, ":", i))
																		_ = tokens

																		i = i + 1
																		_ = i

																	}

																} else {
																	if nc(dat, i, ".") {
																		{
																			tokens = append(tokens, newToken(TokenTypeDot{}, ".", i))
																			_ = tokens

																			i = i + 1
																			_ = i

																		}

																	} else {
																		if nc(dat, i, "\"") {
																			{
																				iStart := i
																				_ = iStart

																				str := bs("\"")
																				_ = str

																				i = i + 1
																				_ = i

																				for {
																					if nc(dat, i, "\"") {
																						{
																							str = append(str, dat[i])
																							_ = str

																							i = i + 1
																							_ = i

																							break
																						}

																					} else {
																						if nc(dat, i, "\\\"") {
																							{
																								str = append(str, dat[i], dat[i+1])
																								_ = str

																								i = i + 2
																								_ = i

																							}

																						} else {
																							{
																								str = append(str, dat[i])
																								_ = str

																								i = i + 1
																								_ = i

																							}

																						}
																					}
																				}

																				tokens = append(tokens, newToken(TokenTypeStringLiteral{}, string(str), iStart))
																				_ = tokens

																			}

																		} else {
																			if nc(dat, i, "`") {
																				{
																					iStart := i
																					_ = iStart

																					str := bs("`")
																					_ = str

																					i = i + 1
																					_ = i

																					for {
																						if nc(dat, i, "`") {
																							{
																								str = append(str, dat[i])
																								_ = str

																								i = i + 1
																								_ = i

																								break
																							}

																						} else {
																							if nc(dat, i, "\\\"") {
																								{
																									str = append(str, dat[i], dat[i+1])
																									_ = str

																									i = i + 2
																									_ = i

																								}

																							} else {
																								{
																									str = append(str, dat[i])
																									_ = str

																									i = i + 1
																									_ = i

																								}

																							}
																						}
																					}

																					tokens = append(tokens, newToken(TokenTypeStringLiteral{}, string(str), iStart))
																					_ = tokens

																				}

																			} else {
																				{
																					str := peekBinaryOp(dat, i)
																					_ = str

																					if len(str) > 0 {
																						{
																							tokens = append(tokens, newToken(TokenTypeBinaryOp{}, str, i))
																							_ = tokens

																							i = i + len(str)
																							_ = i

																						}

																					} else {
																						if nc(dat, i, "=") {
																							{
																								tokens = append(tokens, newToken(TokenTypeEquals{}, "=", i))
																								_ = tokens

																								i = i + 1
																								_ = i

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
	_ = t

	t, tokens = consumeToken(tokens, TokenTypeReturn{})
	_ = t
	_ = tokens

	r := Return{}
	_ = r

	if peekToken(tokens, TokenTypeNewline{}) {
		{
			return r, tokens
		}

	}
	e := sentinelExprJank()
	_ = e

	e, tokens = parseExpr(tokens)
	_ = e
	_ = tokens

	r.Exprs = append(r.Exprs, e)
	_ = r.Exprs

	for {
		if not(peekToken(tokens, TokenTypeComma{})) {
			{
				break
			}

		}
		t, tokens = consumeToken(tokens, TokenTypeComma{})
		_ = t
		_ = tokens

		e, tokens = parseExpr(tokens)
		_ = e
		_ = tokens

		r.Exprs = append(r.Exprs, e)
		_ = r.Exprs

	}

	return r, tokens
}

func parseStatement(tokens []Token) (Statement, []Token) {
	if peekToken(tokens, TokenTypeContinue{}) {
		{
			t, tokens := consumeToken(tokens, TokenTypeContinue{})
			_ = t
			_ = tokens

			return StatementContinue{}, tokens
		}

	} else {
		if peekToken(tokens, TokenTypeBreak{}) {
			{
				t, tokens := consumeToken(tokens, TokenTypeBreak{})
				_ = t
				_ = tokens

				return StatementBreak{}, tokens
			}

		} else {
			if peekToken(tokens, TokenTypeWhile{}) {
				{
					w, tokens := parseWhile(tokens)
					_ = w
					_ = tokens

					return StatementWhile{w}, tokens
				}

			} else {
				if peekToken(tokens, TokenTypeIf{}) {
					{
						i, tokens := parseIf(tokens)
						_ = i
						_ = tokens

						return StatementIf{i}, tokens
					}

				} else {
					if peekToken(tokens, TokenTypeReturn{}) {
						{
							r, tokens := parseReturn(tokens)
							_ = r
							_ = tokens

							return StatementReturn{r}, tokens
						}

					} else {
						if peekToken(tokens, TokenTypeMatch{}) {
							{
								m, tokens := parseMatch(tokens)
								_ = m
								_ = tokens

								return StatementMatch{m}, tokens
							}

						}
					}
				}
			}
		}
	}
	isAssignment := false
	_ = isAssignment

	ass := Assignment{}
	_ = ass

	isAssignment, ass, tokens = tryParseAssignment(tokens)
	_ = isAssignment
	_ = ass
	_ = tokens

	if isAssignment {
		{
			return StatementAssignment{ass}, tokens
		}

	}
	expr, tokens := parseExpr(tokens)
	_ = expr
	_ = tokens

	return StatementExpr{expr}, tokens
}

func sentinelLValueJank() LValue {
	return LValueVariable{""}
}

func identsToLValue(idents []string) LValue {
	if len(idents) == 0 {
		{
			panic("requesting idents to lvalue for no idents")
		}

	} else {
		if len(idents) == 1 {
			{
				return LValueVariable{idents[0]}
			}

		}
	}
	l := LValueVariable{idents[0]}
	_ = l

	r := LValueVariable{idents[1]}
	_ = r

	acc := sentinelLValueJank()
	_ = acc

	acc = LValueDot{DotLValue{l, r}}
	_ = acc

	idents = slice(idents, 2)
	_ = idents

	for {
		if len(idents) == 0 {
			{
				return acc
			}

		} else {
			if len(idents) == 1 {
				{
					dlv := DotLValue{acc, LValueVariable{idents[0]}}
					_ = dlv

					return LValueDot{dlv}
				}

			} else {
				{
					l := LValueVariable{idents[0]}
					_ = l

					r := LValueVariable{idents[1]}
					_ = r

					dlv := DotLValue{l, r}
					_ = dlv

					acc = LValueDot{DotLValue{acc, LValueDot{dlv}}}
					_ = acc

					idents = slice(idents, 2)
					_ = idents

				}

			}
		}
	}

}

func tryParseAssignment(tokens []Token) (bool, Assignment, []Token) {
	origTokens := tokens
	_ = origTokens

	t := Token{}
	_ = t

	ass := Assignment{}
	_ = ass

	ass.IsReassignment = false
	_ = ass.IsReassignment

	collectedIdents := []string{}
	_ = collectedIdents

	for {
		if peekTokens(tokens, []TokenType{TokenTypeIdent{}, TokenTypeEquals{}}) {
			{
				ass.IsReassignment = true
				_ = ass.IsReassignment

				t, tokens = consumeToken(tokens, TokenTypeIdent{})
				_ = t
				_ = tokens

				collectedIdents = append(collectedIdents, t.Value)
				_ = collectedIdents

				ass.LValues = append(ass.LValues, identsToLValue(collectedIdents))
				_ = ass.LValues

				collectedIdents = []string{}
				_ = collectedIdents

				t, tokens = consumeToken(tokens, TokenTypeEquals{})
				_ = t
				_ = tokens

				break
			}

		} else {
			if peekTokens(tokens, []TokenType{TokenTypeIdent{}, TokenTypeColon{}, TokenTypeEquals{}}) {
				{
					ass.IsReassignment = false
					_ = ass.IsReassignment

					t, tokens = consumeToken(tokens, TokenTypeIdent{})
					_ = t
					_ = tokens

					collectedIdents = append(collectedIdents, t.Value)
					_ = collectedIdents

					ass.LValues = append(ass.LValues, identsToLValue(collectedIdents))
					_ = ass.LValues

					collectedIdents = []string{}
					_ = collectedIdents

					t, tokens = consumeToken(tokens, TokenTypeColon{})
					_ = t
					_ = tokens

					t, tokens = consumeToken(tokens, TokenTypeEquals{})
					_ = t
					_ = tokens

					break
				}

			} else {
				if peekTokens(tokens, []TokenType{TokenTypeIdent{}, TokenTypeComma{}}) {
					{
						t, tokens = consumeToken(tokens, TokenTypeIdent{})
						_ = t
						_ = tokens

						collectedIdents = append(collectedIdents, t.Value)
						_ = collectedIdents

						ass.LValues = append(ass.LValues, identsToLValue(collectedIdents))
						_ = ass.LValues

						collectedIdents = []string{}
						_ = collectedIdents

						t, tokens = consumeToken(tokens, TokenTypeComma{})
						_ = t
						_ = tokens

					}

				} else {
					if peekTokens(tokens, []TokenType{TokenTypeIdent{}, TokenTypeDot{}}) {
						{
							t, tokens = consumeToken(tokens, TokenTypeIdent{})
							_ = t
							_ = tokens

							collectedIdents = append(collectedIdents, t.Value)
							_ = collectedIdents

							t, tokens = consumeToken(tokens, TokenTypeDot{})
							_ = t
							_ = tokens

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
	_ = ass.RValue
	_ = tokens

	return true, ass, tokens
}

func sentinelExprJank() Expr {
	return ExprVarRef{""}
}

func parseInitializer(tokens []Token, typ string) (Initializer, []Token) {
	t := Token{}
	_ = t

	i := Initializer{}
	_ = i

	i.Type = typ
	_ = i.Type

	t, tokens = consumeToken(tokens, TokenTypeLCurly{})
	_ = t
	_ = tokens

	for {
		if peekToken(tokens, TokenTypeRCurly{}) {
			{
				break
			}

		}
		arg, tokens2 := parseExpr(tokens)
		_ = arg
		_ = tokens2

		tokens = tokens2
		_ = tokens

		i.Params = append(i.Params, arg)
		_ = i.Params

		if peekToken(tokens, TokenTypeComma{}) {
			{
				t, tokens = consumeToken(tokens, TokenTypeComma{})
				_ = t
				_ = tokens

			}

		} else {
			{
				break
			}

		}
	}

	t, tokens = consumeToken(tokens, TokenTypeRCurly{})
	_ = t
	_ = tokens

	return i, tokens
}

func parseExpr(tokens []Token) (Expr, []Token) {
	t := Token{}
	_ = t

	expr := sentinelExprJank()
	_ = expr

	if peekToken(tokens, TokenTypeLBracket{}) {
		{
			i := Initializer{}
			_ = i

			typ := ""
			_ = typ

			typ, tokens = parseType(tokens)
			_ = typ
			_ = tokens

			i, tokens = parseInitializer(tokens, typ)
			_ = i
			_ = tokens

			expr = ExprInitializer{i}
			_ = expr

		}

	} else {
		if peekToken(tokens, TokenTypeStringLiteral{}) {
			{
				t, tokens = consumeToken(tokens, TokenTypeStringLiteral{})
				_ = t
				_ = tokens

				expr = ExprStringLiteral{t.Value}
				_ = expr

			}

		} else {
			if peekToken(tokens, TokenTypeLCurly{}) {
				{
					b := Block{}
					_ = b

					b, tokens = parseBlock(tokens)
					_ = b
					_ = tokens

					expr = ExprBlock{b}
					_ = expr

				}

			} else {
				if peekToken(tokens, TokenTypeIdent{}) {
					{
						t, tokens = consumeToken(tokens, TokenTypeIdent{})
						_ = t
						_ = tokens

						expr = ExprVarRef{t.Value}
						_ = expr

					}

				} else {
					if peekToken(tokens, TokenTypeIntLiteral{}) {
						{
							t, tokens = consumeToken(tokens, TokenTypeIntLiteral{})
							_ = t
							_ = tokens

							expr = ExprIntLiteral{atoi(t.Value)}
							_ = expr

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
		if peekToken(tokens, TokenTypeLParen{}) {
			{
				funcCall := FuncCall{}
				_ = funcCall

				funcCall, tokens = parseFuncCall(tokens, expr)
				_ = funcCall
				_ = tokens

				expr = ExprFuncCall{funcCall}
				_ = expr

				continue
			}

		} else {
			if peekToken(tokens, TokenTypeBinaryOp{}) {
				{
					t, tokens = consumeToken(tokens, TokenTypeBinaryOp{})
					_ = t
					_ = tokens

					binop := BinOp{}
					_ = binop

					binop.LHS = expr
					_ = binop.LHS

					binop.Op = t.Value
					_ = binop.Op

					binop.RHS, tokens = parseExpr(tokens)
					_ = binop.RHS
					_ = tokens

					expr = ExprBinOp{binop}
					_ = expr

					continue
				}

			} else {
				if peekToken(tokens, TokenTypeLBracket{}) {
					{
						t, tokens = consumeToken(tokens, TokenTypeLBracket{})
						_ = t
						_ = tokens

						arrayAccess := ArrayAccess{}
						_ = arrayAccess

						arrayAccess.LHS = expr
						_ = arrayAccess.LHS

						arrayAccess.Index, tokens = parseExpr(tokens)
						_ = arrayAccess.Index
						_ = tokens

						t, tokens = consumeToken(tokens, TokenTypeRBracket{})
						_ = t
						_ = tokens

						expr = ExprArrayAccess{arrayAccess}
						_ = expr

						continue
					}

				} else {
					if peekToken(tokens, TokenTypeLCurly{}) {
						{
							typ := exprToType(expr)
							_ = typ

							i := Initializer{}
							_ = i

							i, tokens = parseInitializer(tokens, typ)
							_ = i
							_ = tokens

							expr = ExprInitializer{i}
							_ = expr

						}

					} else {
						if peekToken(tokens, TokenTypeDot{}) {
							{
								dotAccess := DotAccess{}
								_ = dotAccess

								dotAccess.LHS = expr
								_ = dotAccess.LHS

								t, tokens = consumeToken(tokens, TokenTypeDot{})
								_ = t
								_ = tokens

								t, tokens = consumeToken(tokens, TokenTypeIdent{})
								_ = t
								_ = tokens

								dotAccess.Field = t.Value
								_ = dotAccess.Field

								expr = ExprDotAccess{dotAccess}
								_ = expr

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
		if binding, ok := matchExpr.(ExprVarRef); ok {
			_ = binding
			v := binding.Value
			_ = v
			{
				return v
			}

		} else if binding, ok := matchExpr.(ExprDotAccess); ok {
			_ = binding
			dotAccess := binding.Value
			_ = dotAccess
			{
				return exprToType(dotAccess.LHS) + "." + dotAccess.Field
			}

		}
	}
	panic(fmt.Sprintf("trying to convert expr to type, unhandled expr: %#v", expr))
}

func parseFuncCall(tokens []Token, lhs Expr) (FuncCall, []Token) {
	t := Token{}
	_ = t

	t, tokens = consumeToken(tokens, TokenTypeLParen{})
	_ = t
	_ = tokens

	funcCall := FuncCall{}
	_ = funcCall

	funcCall.LHS = lhs
	_ = funcCall.LHS

	for {
		if peekToken(tokens, TokenTypeRParen{}) {
			{
				break
			}

		}
		nextParam, tokens2 := parseExpr(tokens)
		_ = nextParam
		_ = tokens2

		funcCall.Params = append(funcCall.Params, nextParam)
		_ = funcCall.Params

		tokens = tokens2
		_ = tokens

		if peekToken(tokens, TokenTypeComma{}) {
			{
				t, tokens = consumeToken(tokens, TokenTypeComma{})
				_ = t
				_ = tokens

			}

		} else {
			{
				break
			}

		}
	}

	t, tokens = consumeToken(tokens, TokenTypeRParen{})
	_ = t
	_ = tokens

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
	_ = emp

	emp.Type, tokens = parseType(tokens)
	_ = emp.Type
	_ = tokens

	t := Token{}
	_ = t

	t, tokens = consumeToken(tokens, TokenTypeLCurly{})
	_ = t
	_ = tokens

	if peekToken(tokens, TokenTypeIdent{}) {
		{
			t, tokens = consumeToken(tokens, TokenTypeIdent{})
			_ = t
			_ = tokens

			emp.Binding = t.Value
			_ = emp.Binding

		}

	}
	t, tokens = consumeToken(tokens, TokenTypeRCurly{})
	_ = t
	_ = tokens

	return MatchPatternEnum{emp}, tokens
}

func parseMatch(tokens []Token) (Match, []Token) {
	t := Token{}
	_ = t

	t, tokens = consumeToken(tokens, TokenTypeMatch{})
	_ = t
	_ = tokens

	m := Match{}
	_ = m

	t, tokens = consumeToken(tokens, TokenTypeLParen{})
	_ = t
	_ = tokens

	m.Matched, tokens = parseExpr(tokens)
	_ = m.Matched
	_ = tokens

	t, tokens = consumeToken(tokens, TokenTypeRParen{})
	_ = t
	_ = tokens

	t, tokens = consumeToken(tokens, TokenTypeLCurly{})
	_ = t
	_ = tokens

	for {
		if peekToken(tokens, TokenTypeRCurly{}) {
			{
				break
			}

		}
		arm := MatchArm{}
		_ = arm

		arm.Pattern, tokens = parseMatchPattern(tokens)
		_ = arm.Pattern
		_ = tokens

		t, tokens = consumeToken(tokens, TokenTypeColon{})
		_ = t
		_ = tokens

		arm.Body, tokens = parseExpr(tokens)
		_ = arm.Body
		_ = tokens

		m.Arms = append(m.Arms, arm)
		_ = m.Arms

		if not(peekToken(tokens, TokenTypeComma{})) {
			{
				break
			}

		}
		t, tokens = consumeToken(tokens, TokenTypeComma{})
		_ = t
		_ = tokens

	}

	t, tokens = consumeToken(tokens, TokenTypeRCurly{})
	_ = t
	_ = tokens

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
	_ = t

	t, tokens = consumeToken(tokens, TokenTypeIf{})
	_ = t
	_ = tokens

	i := If{}
	_ = i

	t, tokens = consumeToken(tokens, TokenTypeLParen{})
	_ = t
	_ = tokens

	i.Cond, tokens = parseExpr(tokens)
	_ = i.Cond
	_ = tokens

	t, tokens = consumeToken(tokens, TokenTypeRParen{})
	_ = t
	_ = tokens

	i.IfBody, tokens = parseExpr(tokens)
	_ = i.IfBody
	_ = tokens

	if peekToken(tokens, TokenTypeElse{}) {
		{
			t, tokens = consumeToken(tokens, TokenTypeElse{})
			_ = t
			_ = tokens

			i.ElseBody, tokens = parseStatement(tokens)
			_ = i.ElseBody
			_ = tokens

		}

	}
	return i, tokens
}

func parseWhile(tokens []Token) (While, []Token) {
	t := Token{}
	_ = t

	t, tokens = consumeToken(tokens, TokenTypeWhile{})
	_ = t
	_ = tokens

	w := While{}
	_ = w

	w.Body, tokens = parseBlock(tokens)
	_ = w.Body
	_ = tokens

	return w, tokens
}

func parseBlock(tokens []Token) (Block, []Token) {
	stmts := []Statement{}
	_ = stmts

	t := Token{}
	_ = t

	t, tokens = consumeToken(tokens, TokenTypeLCurly{})
	_ = t
	_ = tokens

	for {
		tokens = skipNewlines(tokens)
		_ = tokens

		stmt, tokens2 := parseStatement(tokens)
		_ = stmt
		_ = tokens2

		tokens = tokens2
		_ = tokens

		stmts = append(stmts, stmt)
		_ = stmts

		if peekToken(tokens, TokenTypeRCurly{}) {
			{
				break
			}

		}
	}

	t, tokens = consumeToken(tokens, TokenTypeRCurly{})
	_ = t
	_ = tokens

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
	t, tokens := consumeToken(tokens, TokenTypeImport{})
	_ = t
	_ = tokens

	t, tokens = consumeToken(tokens, TokenTypeStringLiteral{})
	_ = t
	_ = tokens

	return Import{t.Value}, tokens
}

func parseEnum(tokens []Token) (Enum, []Token) {
	t, tokens := consumeToken(tokens, TokenTypeEnumDecl{})
	_ = t
	_ = tokens

	t, tokens = consumeToken(tokens, TokenTypeIdent{})
	_ = t
	_ = tokens

	e := Enum{}
	_ = e

	e.Name = t.Value
	_ = e.Name

	t, tokens = consumeToken(tokens, TokenTypeLCurly{})
	_ = t
	_ = tokens

	for {
		if peekToken(tokens, TokenTypeRCurly{}) {
			{
				t, tokens = consumeToken(tokens, TokenTypeRCurly{})
				_ = t
				_ = tokens

				break
			}

		}
		t, tokens = consumeToken(tokens, TokenTypeIdent{})
		_ = t
		_ = tokens

		variant := EnumVariant{}
		_ = variant

		variant.Name = t.Value
		_ = variant.Name

		if peekToken(tokens, TokenTypeLParen{}) {
			{
				t, tokens = consumeToken(tokens, TokenTypeLParen{})
				_ = t
				_ = tokens

				variant.Type, tokens = parseType(tokens)
				_ = variant.Type
				_ = tokens

				t, tokens = consumeToken(tokens, TokenTypeRParen{})
				_ = t
				_ = tokens

			}

		}
		e.Variants = append(e.Variants, variant)
		_ = e.Variants

		t, tokens = consumeToken(tokens, TokenTypeComma{})
		_ = t
		_ = tokens

	}

	return e, tokens
}

func parseStruct(tokens []Token) (Struct, []Token) {
	t, tokens := consumeToken(tokens, TokenTypeStruct{})
	_ = t
	_ = tokens

	t, tokens = consumeToken(tokens, TokenTypeIdent{})
	_ = t
	_ = tokens

	s := Struct{}
	_ = s

	s.Name = t.Value
	_ = s.Name

	t, tokens = consumeToken(tokens, TokenTypeLCurly{})
	_ = t
	_ = tokens

	for {
		if peekToken(tokens, TokenTypeRCurly{}) {
			{
				t, tokens = consumeToken(tokens, TokenTypeRCurly{})
				_ = t
				_ = tokens

				break
			}

		}
		field := StructField{}
		_ = field

		t, tokens = consumeToken(tokens, TokenTypeIdent{})
		_ = t
		_ = tokens

		field.Name = t.Value
		_ = field.Name

		field.Type, tokens = parseType(tokens)
		_ = field.Type
		_ = tokens

		s.Fields = append(s.Fields, field)
		_ = s.Fields

		t, tokens = consumeToken(tokens, TokenTypeComma{})
		_ = t
		_ = tokens

	}

	return s, tokens
}

func parseType(tokens []Token) (string, []Token) {
	name := ""
	_ = name

	t := Token{}
	_ = t

	for {
		if peekToken(tokens, TokenTypeIdent{}) {
			{
				t, tokens = consumeToken(tokens, TokenTypeIdent{})
				_ = t
				_ = tokens

				name = name + t.Value
				_ = name

			}

		} else {
			if peekToken(tokens, TokenTypeLBracket{}) {
				{
					t, tokens = consumeToken(tokens, TokenTypeLBracket{})
					_ = t
					_ = tokens

					name = name + t.Value
					_ = name

				}

			} else {
				if peekToken(tokens, TokenTypeRBracket{}) {
					{
						t, tokens = consumeToken(tokens, TokenTypeRBracket{})
						_ = t
						_ = tokens

						name = name + t.Value
						_ = name

					}

				} else {
					if peekToken(tokens, TokenTypeDot{}) {
						{
							t, tokens = consumeToken(tokens, TokenTypeDot{})
							_ = t
							_ = tokens

							name = name + t.Value
							_ = name

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
	t, tokens := consumeToken(tokens, TokenTypeFuncDecl{})
	_ = t
	_ = tokens

	t, tokens = consumeToken(tokens, TokenTypeIdent{})
	_ = t
	_ = tokens

	f := Function{}
	_ = f

	f.Name = t.Value
	_ = f.Name

	t, tokens = consumeToken(tokens, TokenTypeLParen{})
	_ = t
	_ = tokens

	for {
		if peekToken(tokens, TokenTypeRParen{}) {
			{
				t, tokens = consumeToken(tokens, TokenTypeRParen{})
				_ = t
				_ = tokens

				break
			}

		}
		param := FunctionParam{}
		_ = param

		t, tokens = consumeToken(tokens, TokenTypeIdent{})
		_ = t
		_ = tokens

		param.Name = t.Value
		_ = param.Name

		param.Type, tokens = parseType(tokens)
		_ = param.Type
		_ = tokens

		f.Params = append(f.Params, param)
		_ = f.Params

		if peekToken(tokens, TokenTypeComma{}) {
			{
				t, tokens = consumeToken(tokens, TokenTypeComma{})
				_ = t
				_ = tokens

			}

		}
	}

	if not(peekToken(tokens, TokenTypeLCurly{})) {
		{
			f.ReturnTypes, tokens = parseReturnTypes(tokens)
			_ = f.ReturnTypes
			_ = tokens

		}

	}
	f.Body, tokens = parseBlock(tokens)
	_ = f.Body
	_ = tokens

	return f, tokens
}

func parseReturnTypes(tokens []Token) ([]string, []Token) {
	types := []string{}
	_ = types

	typ := ""
	_ = typ

	t := Token{}
	_ = t

	if peekToken(tokens, TokenTypeLParen{}) {
		{
			t, tokens = consumeToken(tokens, TokenTypeLParen{})
			_ = t
			_ = tokens

			for {
				typ, tokens = parseType(tokens)
				_ = typ
				_ = tokens

				types = append(types, typ)
				_ = types

				if peekToken(tokens, TokenTypeComma{}) {
					{
						t, tokens = consumeToken(tokens, TokenTypeComma{})
						_ = t
						_ = tokens

					}

				} else {
					{
						break
					}

				}
			}

			t, tokens = consumeToken(tokens, TokenTypeRParen{})
			_ = t
			_ = tokens

		}

	} else {
		{
			typ, tokens = parseType(tokens)
			_ = typ
			_ = tokens

			types = append(types, typ)
			_ = types

		}

	}
	return types, tokens
}

func parseDeclaration(tokens []Token) (Declaration, []Token) {
	{
		matchExpr := tokens[0].Type
		if binding, ok := matchExpr.(TokenTypeImport); ok {
			_ = binding
			{
				imp, tokens := parseImport(tokens)
				_ = imp
				_ = tokens

				return DeclarationImport{imp}, tokens
			}

		} else if binding, ok := matchExpr.(TokenTypeEnumDecl); ok {
			_ = binding
			{
				e, tokens := parseEnum(tokens)
				_ = e
				_ = tokens

				return DeclarationEnum{e}, tokens
			}

		} else if binding, ok := matchExpr.(TokenTypeStruct); ok {
			_ = binding
			{
				s, tokens := parseStruct(tokens)
				_ = s
				_ = tokens

				return DeclarationStruct{s}, tokens
			}

		} else if binding, ok := matchExpr.(TokenTypeFuncDecl); ok {
			_ = binding
			{
				f, tokens := parseFunction(tokens)
				_ = f
				_ = tokens

				return DeclarationFunction{f}, tokens
			}

		}
	}
	panic(fmt.Sprintf("unexpected token: %#v", tokens[0]))
}

func parseProgram(tokens []Token) Program {
	p := Program{}
	_ = p

	for {
		tokens = skipNewlines(tokens)
		_ = tokens

		if len(tokens) == 0 {
			{
				break
			}

		}
		declaration, tokens2 := parseDeclaration(tokens)
		_ = declaration
		_ = tokens2

		tokens = tokens2
		_ = tokens

		p.Declarations = append(p.Declarations, declaration)
		_ = p.Declarations

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
			if binding, ok := matchExpr.(TokenTypeNewline); ok {
				_ = binding
				{
					tokens = slice(tokens, 1)
					_ = tokens

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
	_ = i

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
		_ = i

	}

}

func peekToken(tokens []Token, expectedType TokenType) bool {
	if len(tokens) == 0 {
		{
			panic("Unexpected end of input")
		}

	}
	nl := TokenTypeNewline{}
	_ = nl

	if expectedType != nl {
		{
			tokens = skipNewlines(tokens)
			_ = tokens

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
	nl := TokenTypeNewline{}
	_ = nl

	if expectedType != nl {
		{
			tokens = skipNewlines(tokens)
			_ = tokens

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
	_ = s

	i := 0
	_ = i

	for {
		if i >= len(program.Declarations) {
			{
				break
			}

		}
		compiled := compileDeclaration(program.Declarations[i])
		_ = compiled

		s = s + compiled + "\n"
		_ = s

		i = i + 1
		_ = i

	}

	s = s + prelude()
	_ = s

	return s
}

func compileDeclaration(d Declaration) string {
	{
		matchExpr := d
		if binding, ok := matchExpr.(DeclarationEnum); ok {
			_ = binding
			e := binding.Value
			_ = e
			{
				return compileEnum(e)
			}

		} else if binding, ok := matchExpr.(DeclarationImport); ok {
			_ = binding
			imp := binding.Value
			_ = imp
			{
				return compileImport(imp)
			}

		} else if binding, ok := matchExpr.(DeclarationStruct); ok {
			_ = binding
			s := binding.Value
			_ = s
			{
				return compileStruct(s)
			}

		} else if binding, ok := matchExpr.(DeclarationFunction); ok {
			_ = binding
			f := binding.Value
			_ = f
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
	_ = s

	i := 0
	_ = i

	for {
		if i >= len(strukt.Fields) {
			{
				break
			}

		}
		field := strukt.Fields[i]
		_ = field

		s = s + field.Name + " " + compileType(field.Type) + "\n"
		_ = s

		i = i + 1
		_ = i

	}

	s = s + "}\n"
	_ = s

	return s
}

func compileFunction(f Function) string {
	s := "func " + f.Name + "("
	_ = s

	i := 0
	_ = i

	for {
		if i >= len(f.Params) {
			{
				break
			}

		}
		s = s + f.Params[i].Name + " " + compileType(f.Params[i].Type)
		_ = s

		if i < len(f.Params)-1 {
			{
				s = s + ", "
				_ = s

			}

		}
		i = i + 1
		_ = i

	}

	s = s + ") " + compileReturnTypes(f.ReturnTypes) + compileBlock(f.Body)
	_ = s

	return s
}

func compileBlock(b Block) string {
	s := "{\n"
	_ = s

	i := 0
	_ = i

	for {
		if i >= len(b.Statements) {
			{
				break
			}

		}
		s = s + compileStatement(b.Statements[i]) + "\n"
		_ = s

		i = i + 1
		_ = i

	}

	s = s + "}\n"
	_ = s

	return s
}

func compileStatement(stmt Statement) string {
	{
		matchExpr := stmt
		if binding, ok := matchExpr.(StatementAssignment); ok {
			_ = binding
			assignment := binding.Value
			_ = assignment
			{
				return compileAssignment(assignment)
			}

		} else if binding, ok := matchExpr.(StatementWhile); ok {
			_ = binding
			w := binding.Value
			_ = w
			{
				return compileWhile(w)
			}

		} else if binding, ok := matchExpr.(StatementIf); ok {
			_ = binding
			ifStmt := binding.Value
			_ = ifStmt
			{
				return compileIf(ifStmt)
			}

		} else if binding, ok := matchExpr.(StatementReturn); ok {
			_ = binding
			r := binding.Value
			_ = r
			{
				return compileReturn(r)
			}

		} else if binding, ok := matchExpr.(StatementBreak); ok {
			_ = binding
			{
				return "break"
			}

		} else if binding, ok := matchExpr.(StatementContinue); ok {
			_ = binding
			{
				return "continue"
			}

		} else if binding, ok := matchExpr.(StatementExpr); ok {
			_ = binding
			expr := binding.Value
			_ = expr
			{
				return compileExpr(expr)
			}

		} else if binding, ok := matchExpr.(StatementMatch); ok {
			_ = binding
			m := binding.Value
			_ = m
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
	_ = s

	s = s + MatchExprVarName() + " := " + compileExpr(m.Matched) + "\n"
	_ = s

	i := 0
	_ = i

	for {
		if i >= len(m.Arms) {
			{
				break
			}

		}
		if i == 0 {
			{
				s = s + "if "
				_ = s

			}

		} else {
			{
				s = s + " else if "
				_ = s

			}

		}
		s = s + compileMatchPatternTestExpr(m.Arms[i].Pattern) + " {\n"
		_ = s

		s = s + stfuUnusedVars(LValueVariable{BindingVarName()}) + "\n"
		_ = s

		s = s + compileBindingForMatchPattern(m.Arms[i].Pattern)
		_ = s

		s = s + compileExpr(m.Arms[i].Body)
		_ = s

		s = s + "\n}"
		_ = s

		i = i + 1
		_ = i

	}

	s = s + "\n}"
	_ = s

	return s
}

func compileBindingForMatchPattern(pattern MatchPattern) string {
	{
		matchExpr := pattern
		if binding, ok := matchExpr.(MatchPatternEnum); ok {
			_ = binding
			emp := binding.Value
			_ = emp
			{
				if emp.Binding == "" {
					{
						return ""
					}

				}
				s := emp.Binding + " := " + BindingVarName() + ".Value\n"
				_ = s

				s = s + stfuUnusedVars(LValueVariable{emp.Binding}) + "\n"
				_ = s

				return s
			}

		}
	}
	panic(fmt.Sprintf("unreachable, must not have handled a type of match pattern: %#v", pattern))
}

func compileMatchPatternTestExpr(pattern MatchPattern) string {
	{
		matchExpr := pattern
		if binding, ok := matchExpr.(MatchPatternEnum); ok {
			_ = binding
			emp := binding.Value
			_ = emp
			{
				s := BindingVarName() + ", ok := " + MatchExprVarName() + ".(" + compileType(emp.Type) + ")"
				_ = s

				s = s + "; ok"
				_ = s

				return s
			}

		}
	}
	panic(fmt.Sprintf("unreachable, must not have handled a type of match pattern: %#v", pattern))
}

func compileReturn(r Return) string {
	s := "return "
	_ = s

	if len(r.Exprs) == 0 {
		{
			return s
		}

	} else {
		{
			return s + compileExprsCSV(r.Exprs)
		}

	}
}

func compileExprsCSV(exprs []Expr) string {
	s := ""
	_ = s

	i := 0
	_ = i

	for {
		if i >= len(exprs) {
			{
				break
			}

		}
		s = s + compileExpr(exprs[i])
		_ = s

		if i < len(exprs)-1 {
			{
				s = s + ", "
				_ = s

			}

		}
		i = i + 1
		_ = i

	}

	return s
}

func compileWhile(w While) string {
	body := w.Body
	_ = body

	return "for " + compileBlock(body)
}

func compileIf(ifStmt If) string {
	s := "if "
	_ = s

	cond := ifStmt.Cond
	_ = cond

	s = s + compileExpr(cond)
	_ = s

	s = s + " {\n"
	_ = s

	ifBody := ifStmt.IfBody
	_ = ifBody

	s = s + compileExpr(ifBody)
	_ = s

	s = s + "\n}"
	_ = s

	elseBody := ifStmt.ElseBody
	_ = elseBody

	if elseBody != nil {
		{
			s = s + "else {\n"
			_ = s

			s = s + compileStatement(elseBody)
			_ = s

			s = s + "\n}"
			_ = s

		}

	}
	return s
}

func compileAssignment(ass Assignment) string {
	lValues := ass.LValues
	_ = lValues

	s := compileLValues(lValues)
	_ = s

	if ass.IsReassignment {
		{
			s = s + " = "
			_ = s

		}

	} else {
		{
			s = s + " := "
			_ = s

		}

	}
	rValue := ass.RValue
	_ = rValue

	s = s + compileExpr(rValue) + "\n"
	_ = s

	i := 0
	_ = i

	for {
		if i >= len(lValues) {
			{
				break
			}

		}
		lValue := lValues[i]
		_ = lValue

		s = s + stfuUnusedVars(lValue) + "\n"
		_ = s

		i = i + 1
		_ = i

	}

	return s
}

func stfuUnusedVars(lValue LValue) string {
	return "_ = " + compileLValue(lValue)
}

func compileExpr(expr Expr) string {
	{
		matchExpr := expr
		if binding, ok := matchExpr.(ExprVarRef); ok {
			_ = binding
			s := binding.Value
			_ = s
			{
				return s
			}

		} else if binding, ok := matchExpr.(ExprFuncCall); ok {
			_ = binding
			funcCall := binding.Value
			_ = funcCall
			{
				lhs := funcCall.LHS
				_ = lhs

				s := compileExpr(lhs) + "("
				_ = s

				params := funcCall.Params
				_ = params

				s = s + compileExprsCSV(params)
				_ = s

				s = s + ")"
				_ = s

				return s
			}

		} else if binding, ok := matchExpr.(ExprIntLiteral); ok {
			_ = binding
			i := binding.Value
			_ = i
			{
				return strconv.Itoa(i)
			}

		} else if binding, ok := matchExpr.(ExprBinOp); ok {
			_ = binding
			binop := binding.Value
			_ = binop
			{
				lhs := binop.LHS
				_ = lhs

				rhs := binop.RHS
				_ = rhs

				op := binop.Op
				_ = op

				return compileExpr(lhs) + " " + op + " " + compileExpr(rhs)
			}

		} else if binding, ok := matchExpr.(ExprBlock); ok {
			_ = binding
			block := binding.Value
			_ = block
			{
				return compileBlock(block)
			}

		} else if binding, ok := matchExpr.(ExprArrayAccess); ok {
			_ = binding
			aa := binding.Value
			_ = aa
			{
				lhs := aa.LHS
				_ = lhs

				index := aa.Index
				_ = index

				return compileExpr(lhs) + "[" + compileExpr(index) + "]"
			}

		} else if binding, ok := matchExpr.(ExprInitializer); ok {
			_ = binding
			init := binding.Value
			_ = init
			{
				typ := init.Type
				_ = typ

				s := compileType(typ) + "{ "
				_ = s

				params := init.Params
				_ = params

				s = s + compileExprsCSV(params)
				_ = s

				s = s + " }"
				_ = s

				return s
			}

		} else if binding, ok := matchExpr.(ExprStringLiteral); ok {
			_ = binding
			sl := binding.Value
			_ = sl
			{
				return sl
			}

		} else if binding, ok := matchExpr.(ExprDotAccess); ok {
			_ = binding
			da := binding.Value
			_ = da
			{
				lhs := da.LHS
				_ = lhs

				field := da.Field
				_ = field

				return compileExpr(lhs) + "." + field
			}

		}
	}
	panic(fmt.Sprintf("unhandled expr: %#v", expr))
}

func compileLValue(lValue LValue) string {
	{
		matchExpr := lValue
		if binding, ok := matchExpr.(LValueVariable); ok {
			_ = binding
			s := binding.Value
			_ = s
			{
				return s
			}

		} else if binding, ok := matchExpr.(LValueDot); ok {
			_ = binding
			dotLValue := binding.Value
			_ = dotLValue
			{
				lhs := dotLValue.LHS
				_ = lhs

				rhs := dotLValue.RHS
				_ = rhs

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
				_ = s

				i := 0
				_ = i

				for {
					if i >= len(lValues) {
						{
							break
						}

					}
					s = s + compileLValue(lValues[i])
					_ = s

					if i < len(lValues)-1 {
						{
							s = s + ", "
							_ = s

						}

					}
					i = i + 1
					_ = i

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
	_ = s

	i := 0
	_ = i

	for {
		if i >= len(returnTypes) {
			{
				break
			}

		}
		s = s + compileType(returnTypes[i])
		_ = s

		if i < len(returnTypes)-1 {
			{
				s = s + ", "
				_ = s

			}

		}
		i = i + 1
		_ = i

	}

	s = s + ")"
	_ = s

	return s
}

func compileType(typ string) string {
	s := ""
	_ = s

	i := 0
	_ = i

	for {
		if i >= len(typ) {
			{
				break
			}

		}
		if typ[i] != c(".") {
			{
				s = s + string(typ[i])
				_ = s

			}

		}
		i = i + 1
		_ = i

	}

	return s
}

func compileEnumInterface(e Enum) string {
	s := "type " + golangInterfaceName(e) + " interface {\n"
	_ = s

	s = s + golangEnumImplementsInterfaceMethodName(e) + "()\n"
	_ = s

	s = s + "}\n"
	_ = s

	return s
}

func structNameForVariant(e Enum, v EnumVariant) string {
	return e.Name + v.Name
}

func compileEnumStructs(e Enum) string {
	s := ""
	_ = s

	i := 0
	_ = i

	for {
		if i >= len(e.Variants) {
			{
				break
			}

		}
		variant := e.Variants[i]
		_ = variant

		structName := structNameForVariant(e, variant)
		_ = structName

		s = s + "type " + structName + " struct {\n"
		_ = s

		if variant.Type != "" {
			{
				typ := variant.Type
				_ = typ

				s = s + "Value " + compileType(typ) + "\n"
				_ = s

			}

		}
		s = s + "}\n"
		_ = s

		s = s + "func (_ " + structName + ") " + golangEnumImplementsInterfaceMethodName(e) + "() {}\n"
		_ = s

		i = i + 1
		_ = i

	}

	return s
}

func main() {
	if len(os.Args) < 2 {
		{
			fmt.Printf("Usage: ./%s <filename.goy>\n", os.Args[0])
			return
		}

	}
	dat, err := os.ReadFile(os.Args[1])
	_ = dat
	_ = err

	if err != nil {
		{
			panic(err)
		}

	}
	tokens := lex(dat)
	_ = tokens

	program := parseProgram(tokens)
	_ = program

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

func slice[T any](s []T, i ...int) []T {
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
