package parser

import (
	"fmt"
	"testing"

	"github.com/betelgeuse-7/poml/lexer"
)

func TestParserNext(t *testing.T) {
	input := `(p :id "cat" :class "some body" "hello")`
	l := lexer.New(input)
	p := New(l)
	n := p.Next()
	fmt.Println(n)
}
