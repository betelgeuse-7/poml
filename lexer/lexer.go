package lexer

import (
	"fmt"

	"github.com/betelgeuse-7/poml/token"
)

type tokens = []token.Token

const EOF_RUNE = rune(0)

type Lexer struct {
	input      string
	x, y       uint // x -> col, y -> row
	nextX      uint // x + 1
	ch, nextch rune
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
		x:     0,
		y:     1,
		nextX: 1,
	}
	l.ch = rune(l.input[l.x])
	if len(l.input) > 1 {
		l.nextch = rune(l.input[l.nextX])
	} else {
		l.nextch = EOF_RUNE
	}
	return l
}

// for debugging purposes
func (l *Lexer) String() string {
	stats := fmt.Sprintf("Lexer(x: %d, y: %d, nextX: %d, ch: %s, nextCh: %s, len(input): %d)", l.x, l.y, l.nextX, string(l.ch), string(l.nextch), len(l.input))
	if l.nextch == EOF_RUNE {
		stats = fmt.Sprintf("Lexer(x: %d, y: %d, nextX: %d, ch: %s, nextCh: <<<EOF>>>, len(input): %d)", l.x, l.y, l.nextX, string(l.ch), len(l.input))
	}
	return stats
}

func (l *Lexer) advance() {
	if l.nextch == EOF_RUNE {
		l.ch = l.nextch
		l.x++
		return
	}
	if l.nextch == '\n' {
		l.y++
	}
	l.x = l.nextX
	l.ch = rune(l.input[l.x])
	l.nextX++
	if l.nextX == uint(len(l.input)) {
		l.nextch = EOF_RUNE
		return
	}
	l.nextch = rune(l.input[l.nextX])
}

func (l *Lexer) Lex() token.Token {
	if l.ch == EOF_RUNE {
		return token.Token{
			Tok: token.EOF,
			Lit: "<<<EOF>>>",
		}
	}
	if isWhitespace(l.ch) {
		return l.lexWhitespace()
	} else if !isWhitespace(l.ch) && isNotASpecialChar(l.ch) {
		return l.lexIdent()
	}
	switch l.ch {
	case '(':
		l.advance()
		return token.Token{
			Tok: token.LPAREN,
			Lit: "(",
		}
	case ')':
		l.advance()
		return token.Token{
			Tok: token.RPAREN,
			Lit: ")",
		}
	case ':':
		l.advance()
		return token.Token{
			Tok: token.COLON,
			Lit: ":",
		}
	case ';':
		l.advance()
		return token.Token{
			Tok: token.SCOLON,
			Lit: ";",
		}
	}
	l.advance()
	return token.Token{
		Tok: token.ILLEGAL,
		Lit: string(l.ch),
	}
}

func (l *Lexer) lexWhitespace() token.Token {
	start := l.x
	end := l.x
	for isWhitespace(l.ch) {
		end++
		if l.ch == EOF_RUNE {
			break
		}
		l.advance()
	}
	lit := l.input[start:end]
	return token.Token{
		Tok: token.WHITESPACE,
		Lit: lit,
	}
}

func (l *Lexer) lexIdent() token.Token {
	start := l.x
	end := l.x
	for !isWhitespace(l.ch) && isNotASpecialChar(l.ch) {
		end++
		l.advance()
		if l.ch == EOF_RUNE {
			break
		}
	}
	lit := l.input[start:end]
	return token.Token{
		Tok: token.IDENT,
		Lit: lit,
	}
}
