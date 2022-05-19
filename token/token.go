package token

import "fmt"

type TokenType string

type Token struct {
	Tok TokenType
	Lit string
}

func (t Token) String() string {
	return fmt.Sprintf("(%s, %s)", t.Tok, t.Lit)
}

const (
	EOF        TokenType = "EOF"
	ILLEGAL    TokenType = "ILLEGAL"
	WHITESPACE TokenType = "WHITESPACE"

	RPAREN TokenType = "RPAREN"

	TAG     TokenType = "TAG"
	ATTR    TokenType = "ATTR"
	COMMENT TokenType = "COMMENT"
	TEXT    TokenType = "TEXT"
)
