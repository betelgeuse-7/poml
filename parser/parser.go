package parser

import (
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
	n := &Node{}
	if p.tok.Tok == token.EOF {
		return n
	}
	if p.tok.Tok == token.TAG {
		n.Tag = p.tok.Lit
	}
	for p.tok.Tok == token.ATTR {
		attr := Attr{Key: p.tok.Lit}
		attr.Val = p.parseAttrVal()
		n.Attrs = append(n.Attrs, attr)
	}
	if p.tok.Tok == token.TEXT {
		n.HasText = true
		n.Text = p.parseTextContent()
	}
	/*
		p.parseNodeChildren(n)
	*/
	p.advance()
	return n
}

func (p *Parser) parseAttrVal() string {
	p.advance()
	if cur := p.tok.Tok; cur == token.TEXT {
		return p.tok.Lit
	}
	return ""
}

func (p *Parser) parseTextContent() string {
	p.advance()
	if cur := p.tok.Tok; cur == token.TEXT {
		return p.tok.Lit
	}
	return ""
}

/*
func (p *Parser) parseNodeChildren(n *Node) {
	p.advance()
	if cur := p.tok; cur.Tok == token.TAG {
		nn := &Node{
			Tag: cur.Lit,
		}
		p.advance()
		switch cur.Tok {
		case token.TEXT:
			nn.HasText = true
			nn.Text = p.parseTextContent()
		case token.ATTR:
			for p.tok.Tok == token.ATTR {
				attr := Attr{Key: p.tok.Lit}
				attr.Val = p.parseAttrVal()
				nn.Attrs = append(nn.Attrs, attr)
			}
		case token.TAG:
			p.parseNodeChildren(nn)
		}
		n.Children = append(n.Children, nn)
	}
}
*/
