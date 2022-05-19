package lexer

import (
	"fmt"

	"github.com/betelgeuse-7/poml/token"
)

type tokens = []token.Token

const EOF_RUNE = rune(0)

// character position info
type CharPos struct {
	x, y uint
}

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

func (l *Lexer) Lex() (token.Token, CharPos) {
	if l.ch == EOF_RUNE {
		return token.Token{
			Tok: token.EOF,
			Lit: "<<<EOF>>>",
		}, l.charPos()
	}
	if isWhitespace(l.ch) {
		return l.lexWhitespace()
	}
	switch l.ch {
	case '(':
		return l.lexTag()
	case ')':
		l.advance()
		return token.Token{
			Tok: token.RPAREN,
			Lit: ")",
		}, l.charPos()
	case '"':
		return l.lexText()
	case ':':
		return l.lexAttr()
	case ';':
		return l.lexComment()
	}

	l.advance()
	return token.Token{
		Tok: token.ILLEGAL,
		Lit: string(l.ch),
	}, l.charPos()
}

func (l *Lexer) lexWhitespace() (token.Token, CharPos) {
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
	}, l.charPos()
}

func (l *Lexer) lexComment() (token.Token, CharPos) {
	l.advance()
	start := l.x
	for l.ch != '\n' {
		l.advance()
		if l.ch == EOF_RUNE {
			break
		}
	}
	lit := l.input[start:l.x]
	return token.Token{
		Tok: token.COMMENT,
		Lit: lit,
	}, l.charPos()
}

func (l *Lexer) lexTag() (token.Token, CharPos) {
	l.advance()
	start := l.x
	for !isWhitespace(l.ch) {
		l.advance()
		if l.ch == EOF_RUNE {
			break
		}
	}
	lit := l.input[start:l.x]
	return token.Token{
		Tok: token.TAG,
		Lit: lit,
	}, l.charPos()
}

func (l *Lexer) lexText() (token.Token, CharPos) {
	res := ""
	res += string(l.ch)
	for {
		// if we come across an escape slash in a string,
		// we skip it, and add the escaped character to res.
		if l.ch == '\\' {
			if l.nextch == EOF_RUNE {
				res += string(l.ch)
				break
			}
			res += string(l.nextch)
			l.advance()
			continue
		}
		l.advance()
		if l.ch == EOF_RUNE {
			break
		}
		res += string(l.ch)
		// last '"'
		if l.ch == '"' {
			break
		}
	}
	l.advance()
	return token.Token{
		Tok: token.TEXT,
		Lit: res,
	}, l.charPos()
}

func (l *Lexer) lexAttr() (token.Token, CharPos) {
	start := l.x
	for !isWhitespace(l.ch) {
		l.advance()
		if l.ch == EOF_RUNE {
			break
		}
	}
	lit := l.input[start:l.x]
	return token.Token{
		Tok: token.ATTR,
		Lit: lit,
	}, l.charPos()
}

func (l *Lexer) charPos() CharPos {
	return CharPos{
		x: l.x,
		y: l.y,
	}
}
