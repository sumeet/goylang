package main

import (
	"bytes"
	"fmt"
)

type TokenType int

const (
	FuncDecl TokenType = iota
	LParen
	RParen
	LCurly
	RCurly
	Ident
	IntLiteral
	Assignment
	Reassignment
	BinaryOp
	Newline
	Comma
	StringLiteral
)

func formatToken(t Token) string {
	switch t.Type {
	case FuncDecl:
		return "FuncDecl"
	case LParen:
		return "LParen"
	case RParen:
		return "RParen"
	case LCurly:
		return "LCurly"
	case RCurly:
		return "RCurly"
	case Ident:
		return fmt.Sprintf("Ident(%s)", t.Value)
	case IntLiteral:
		return fmt.Sprintf("IntLiteral(%s)", t.Value)
	case Assignment:
		return "Assignment"
	case Reassignment:
		return "Reassignment"
	case BinaryOp:
		return fmt.Sprintf("BinaryOp(%s)", t.Value)
	case Newline:
		return "Newline"
	case Comma:
		return "Comma"
	case StringLiteral:
		return fmt.Sprintf("StringLiteral(%s)", t.Value)
	default:
		panic(fmt.Sprintf("unknown token type %d", t.Type))
	}
}

func formatTokenType(t TokenType) string {
	switch t {
	case FuncDecl:
		return "FuncDecl"
	case LParen:
		return "LParen"
	case RParen:
		return "RParen"
	case LCurly:
		return "LCurly"
	case RCurly:
		return "RCurly"
	case Ident:
		return "Ident"
	case IntLiteral:
		return "IntLiteral"
	case Assignment:
		return "Assignment"
	case Reassignment:
		return "Reassignment"
	case BinaryOp:
		return "BinaryOp"
	case Newline:
		return "Newline"
	case Comma:
		return "Comma"
	case StringLiteral:
		return "StringLiteral"
	default:
		panic(fmt.Sprintf("unknown token type %d", t))
	}
}

func printTokens(tokens []Token) {
	for _, t := range tokens {
		fmt.Printf("%s ", formatToken(t))
	}
	fmt.Println()
}

type Token struct {
	Type  TokenType
	Value string
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isNumeric(b byte) bool {
	return b >= '0' && b <= '9'
}

func lex(dat []byte) []Token {
	var tokens []Token

	for i := 0; i < len(dat); i += 1 {
		if isAlpha(dat[i]) {
			thisIdent := []byte{dat[i]}

			for i += 1; isAlpha(dat[i]); i += 1 {
				thisIdent = append(thisIdent, dat[i])
			}
			i -= 1
			if bytes.Compare(thisIdent, []byte("func")) == 0 {
				tokens = append(tokens, Token{FuncDecl, string(thisIdent)})
			} else {
				tokens = append(tokens, Token{Ident, string(thisIdent)})
			}
		} else if isNumeric(dat[i]) {
			thisInt := []byte{dat[i]}
			for i += 1; isNumeric(dat[i]); i += 1 {
				thisInt = append(thisInt, dat[i])
			}
			i -= 1
			tokens = append(tokens, Token{IntLiteral, string(thisInt)})
		} else if dat[i] == ' ' || dat[i] == '\t' {
			// ignore
		} else if dat[i] == '\n' {
			tokens = append(tokens, Token{Newline, "\n"})
		} else if bytes.Compare(dat[i:i+2], []byte{'\r', '\n'}) == 0 {
			tokens = append(tokens, Token{Newline, "\r\n"})
		} else if dat[i] == '(' {
			tokens = append(tokens, Token{LParen, string([]byte{dat[i]})})
		} else if dat[i] == ')' {
			tokens = append(tokens, Token{RParen, string([]byte{dat[i]})})
		} else if dat[i] == '{' {
			tokens = append(tokens, Token{LCurly, string([]byte{dat[i]})})
		} else if dat[i] == '}' {
			tokens = append(tokens, Token{RCurly, string([]byte{dat[i]})})
		} else if bytes.Compare(dat[i:i+2], []byte{':', '='}) == 0 {
			tokens = append(tokens, Token{Assignment, string([]byte{dat[i]})})
			i += 1
		} else if dat[i] == '=' {
			tokens = append(tokens, Token{Reassignment, string([]byte{dat[i]})})
		} else if dat[i] == '+' {
			tokens = append(tokens, Token{BinaryOp, string([]byte{dat[i]})})
		} else if dat[i] == ',' {
			tokens = append(tokens, Token{Comma, string([]byte{dat[i]})})
		} else if dat[i] == '"' {
			thisStringLit := []byte{dat[i]}
			i += 1
			// TODO: this doesn't handle an escaped "
			for dat[i] != '"' {
				thisStringLit = append(thisStringLit, dat[i])
				i += 1
			}
			thisStringLit = append(thisStringLit, '"')
			tokens = append(tokens, Token{StringLiteral, string(thisStringLit)})
		} else {
			panic(fmt.Sprintf("unrecognized character %c", dat[i]))
		}
	}
	return tokens
}
