// The lexer outputs token.IDENTs, but we don't know if that token.IDENT is a tag, attr, or a raw text.
// This package is a middleware between the lexer and the parser, that eases the job of the parser
// by determining if the current token.IDENT is a token.TAG, token.ATTR, or a token.TEXT.
//
// Returns a token stream.
package ident

import "github.com/betelgeuse-7/poml/token"

type tokens = []token.Token

var EOF_TOK = token.Token{Tok: token.EOF, Lit: "<<<EOF>>>"}

type Identifier struct {
	tokenStream     tokens
	pos             uint
	curTok, nextTok token.Token
}

func New(tokenStream tokens) *Identifier {
	i := &Identifier{
		tokenStream: tokenStream,
		pos:         0,
	}
	i.curTok = i.tokenStream[i.pos]
	if len(i.tokenStream) > 1 {
		i.nextTok = i.tokenStream[i.pos+1]
	} else {
		i.nextTok = EOF_TOK
	}
	return i
}

/*
func (i *Identifier) advance() {
	if i.pos == uint(len(i.tokenStream)-1) {
		i.nextTok = EOF_TOK
		return
	}
	i.pos++
	i.curTok = i.tokenStream[i.pos]
	if i.pos+1 > uint(len(i.tokenStream)-1) {
		i.nextTok = EOF_TOK
		return
	}
	i.nextTok = i.tokenStream[i.pos+1]
}
*/

func (i *Identifier) MarkWithProperTokenTypes(ts tokens) tokens {
	for _, v := range ts {
		if v.Tok == token.IDENT {
			if isATag(i, v) {
				v.Tok = token.TAG
			} else if isAnAttr(i, v) {
				v.Tok = token.ATTR
			} else if isAText(i, v) {
				v.Tok = token.TEXT
			} else if isAFunctionCall(i, v) {
				v.Tok = token.FUNCTIONCALL
			} else {
				panic("not a tag, attr, text, nor a function call\n")
			}
		}
	}
	i.concatenateTextGroups(ts)
	return ts
}

// concatenate token.IDENTs that are marked token.TEXTs, that are related
// (single string. they should be together)
func (i *Identifier) concatenateTextGroups(ts tokens) {

}

func isATag(i *Identifier, tok token.Token) bool {
	prev := i.tokenStream[i.pos-1]
	return prev.Tok == token.LPAREN
}

func isAnAttr(i *Identifier, tok token.Token) bool {
	prevPrev := i.tokenStream[i.pos-2]
	if isATag(i, prevPrev) {
		prev := i.tokenStream[i.pos-1]
		if prev.Tok == token.WHITESPACE {
			if i.pos == uint(len(i.tokenStream)-1) {
				return false
			}
			next := i.tokenStream[i.pos+1]
			if next.Tok == token.COLON {
				return true
			}
		}
	}
	return false
}

func isAText(i *Identifier, tok token.Token) bool {
	prev := token.Token{}
	// how far we must go to find an attr. (hopefully it is text attr.)
	n := 0
	for !(isAnAttr(i, prev)) {
		// this is our base case, essentially
		if uint(n) == i.pos {
			return false
		}
		n++
	}
	// came across an attr.
	attrTok := i.tokenStream[i.pos-uint(n)]
	if attrTok.Lit == "text" {
		return true
	}
	return isAText(i, tok)
}

func isAFunctionCall(i *Identifier, tok token.Token) bool {
	return true
}
