package lexer

import (
	"fmt"
	"testing"

	"github.com/betelgeuse-7/poml/token"
)

var (
	TAG = func(lit string) token.Token {
		return token.Token{
			Tok: token.TAG,
			Lit: lit,
		}
	}
	WS = func(lit string) token.Token {
		return token.Token{
			Tok: token.WHITESPACE,
			Lit: lit,
		}
	}
	COMMENT = func(lit string) token.Token {
		return token.Token{
			Tok: token.COMMENT,
			Lit: lit,
		}
	}
	TEXT = func(lit string) token.Token {
		return token.Token{
			Tok: token.TEXT,
			Lit: lit,
		}
	}
	ATTR = func(lit string) token.Token {
		return token.Token{
			Tok: token.ATTR,
			Lit: lit,
		}
	}
	RPAREN = func() token.Token {
		return token.Token{
			Tok: token.RPAREN,
			Lit: ")",
		}
	}
)

func TestLexerLex(t *testing.T) {
	input := `(div 
		; this is a comment
		(h3 "A \"Cat\" Picture")
		(a :href "https://google.com" "Google")
		(div :class "cat-div container"
			(img :id "catphoto" :src "https://example.com/img/cat.jpg")
		)
		(button :onclick "doSomething()" "Click Me" :style "background-color: blue; border-radius: 3px;")
	)`
	want := tokens{
		TAG("div"),
		WS(" \n		"),
		COMMENT(" this is a comment"),
		WS("\n		"),
		TAG("h3"),
		WS(" "),
		TEXT("\"A \\\"Cat\\\" Picture\""),
		RPAREN(),
		WS("\n		"),
		TAG("a"),
		WS(" "),
		ATTR(":href"),
		WS(" "),
		TEXT("\"https://google.com\""),
		WS(" "),
		TEXT("\"Google\""),
		RPAREN(),
		WS("\n		"),
		TAG("div"),
		WS(" "),
		ATTR(":class"),
		WS(" "),
		TEXT("\"cat-div container\""),
		WS("\n			"),
		TAG("img"),
		WS(" "),
		ATTR(":id"),
		WS(" "),
		TEXT("\"catphoto\""),
		WS(" "),
		ATTR(":src"),
		WS(" "),
		TEXT("\"https://example.com/img/cat.jpg\""),
		RPAREN(),
		WS("\n		"),
		RPAREN(),
		WS("\n		"),
		TAG("button"),
		WS(" "),
		ATTR(":onclick"),
		WS(" "),
		TEXT("\"doSomething()\""),
		WS(" "),
		TEXT("\"Click Me\""),
		WS(" "),
		ATTR(":style"),
		WS(" "),
		TEXT("\"background-color: blue; border-radius: 3px;\""),
		RPAREN(),
		WS("\n	"),
		RPAREN(),
	}
	l := New(input)
	got := tokens{}
	for {
		tok, _ := l.Lex()
		got = append(got, tok)
		if tok.Tok == token.EOF {
			break
		}
	}
	for i, v := range want {
		if i == len(got) {
			t.Errorf("index out of range of 'got': %d\n", i)
		}
		if got[i].Tok != v.Tok || got[i].Lit != v.Lit {
			t.Errorf("wanted %s, but got %s\n", v.String(), got[i].String())
		}
	}
}

func TestLexerLexWhitespace(t *testing.T) {
	input := "    "
	want := token.Token{Tok: token.WHITESPACE, Lit: "    "}
	l := New(input)
	got, _ := l.lexWhitespace()
	fmt.Println(l.String())
	if got.Tok != want.Tok || got.Lit != want.Lit {
		t.Errorf("expected %s (length: %d), but got %s (length: %d)\n", want.String(), len(want.Lit), got.String(), len(got.Lit))
	}
}

func TestLexerAdvance(t *testing.T) {
	input := "(p text:\nyes; )   q"
	l := New(input)
	_assertEqRune(t, l.ch, '(', "l.ch not correct")
	_assertEqRune(t, l.nextch, 'p', "l.nextCh not correct")
	_assertEqUint(t, l.x, 0, "l.x must be 0")
	l.advance()
	_assertEqRune(t, l.ch, 'p', "l.ch must be p")
	_assertEqRune(t, l.nextch, ' ', "l.nextCh must be ' '")
	_assertEqUint(t, l.nextX, 2, "l.nextX must be 2")
	for {
		if l.x == uint(len(input)-1) {
			break
		}
		l.advance()
	}
	_assertEqRune(t, l.ch, 'q', "l.ch must be q")
	_assertEqRune(t, l.nextch, EOF_RUNE, "l.nextch must be EOF_RUNE")
	_assertEqUint(t, l.y, 2, "l.y must be 2")
}

func _assertEqUint(t *testing.T, val, val2 uint, errmsg string) {
	if val != val2 {
		t.Errorf(errmsg + "\n")
	}
}

func _assertEqRune(t *testing.T, val, val2 rune, errmsg string) {
	if val != val2 {
		t.Errorf(errmsg, "\n")
	}
}
