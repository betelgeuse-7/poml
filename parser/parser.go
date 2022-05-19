package parser

import (
	"fmt"

	"github.com/betelgeuse-7/poml/lexer"
	"github.com/betelgeuse-7/poml/token"
)

type tokens = []token.Token

type Parser struct {
	l           *lexer.Lexer
	tokenStream tokens
	pos         uint
	tok         token.Token
	errors      []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:   l,
		pos: 0,
	}
	for {
		tok := l.Lex()
		p.tokenStream = append(p.tokenStream, tok)
		if tok.Tok == token.EOF {
			break
		}
	}
	p.tok = p.tokenStream[p.pos]
	return p
}

func (p *Parser) advance() {
	if p.pos == uint(len(p.tokenStream)-1) {
		p.tok = token.Token{Tok: token.EOF, Lit: lexer.EOF_LIT}
		return
	}
	p.pos++
	p.tok = p.tokenStream[p.pos]
}

/*
func (p *Parser) peek() token.TokenType {
	if p.pos == uint(len(p.tokenStream)-1) {
		return token.EOF
	}
	return p.tokenStream[p.pos+1].Tok
}
*/
// return all parser errors that happened during subsequent calls to *Parser.Next
func (p *Parser) Errors() []string {
	return p.errors
}

// return next HTML element
func (p *Parser) Next() *Node {
	n := &Node{
		Tag: p.tok.Lit,
	}
	for {
		if p.tok.Tok == token.EOF || p.tok.Tok == token.RPAREN {
			fmt.Println(">>>>>>>>>>>>>>>>> " + p.tok.Tok)
			return n
		}
		// TODO this is buggy
		if p.tok.Tok == token.ILLEGAL {
			panic(fmt.Sprintf("illegal token at line: %d  row: %d\n", p.l.X(), p.l.Y()))
		}
		if p.tok.Tok == token.TEXT {
			n.HasText = true
			n.Text = p.tok.Lit
		}
		if p.tok.Tok == token.ATTR {
			attr := Attr{
				Key: p.tok.Lit,
			}
			p.advance()
			if p.tok.Tok == token.WHITESPACE {
				p.advance()
				if p.tok.Tok == token.TEXT {
					attr.Val = p.tok.Lit
				}
			}
			n.Attrs = append(n.Attrs, attr)
		}
		/*
			if p.tok.Tok == token.TAG {
				n.Children = append(n.Children, p.Next())
			}*/
		p.advance()
	}
}

/*
// next token must match t, otherwise append an error to p.errors
func (p *Parser) matchNext(t token.TokenType) {}
*/
