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

	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
	COLON  TokenType = "COLON"
	SCOLON TokenType = "SCOLON"

	IDENT TokenType = "IDENT"

	TAG     TokenType = "TAG"
	ATTR    TokenType = "ATTR"
	COMMENT TokenType = "COMMENT"
	TEXT    TokenType = "TEXT"
)
