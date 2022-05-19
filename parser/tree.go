package parser

import "fmt"

type Attr struct {
	Key string
	Val string
}

func (a Attr) string() string {
	return fmt.Sprintf("Attr(key: %s, val: %s)", a.Key, a.Val)
}

func attrsString(ax []Attr) string {
	res := "["
	for i, v := range ax {
		if i == len(ax)-1 {
			res += v.string()
		} else {
			res += v.string() + ", "
		}
	}
	res = res + "]"
	return res
}

type Node struct {
	Tag      string
	Attrs    []Attr
	HasText  bool
	Text     string
	Children []*Node
}

func (n *Node) String() string {
	if len(n.Children) == 0 {
		return fmt.Sprintf("Node(tag: %s, attrs: %s, hasText: %v, text: %s, children: 0)", n.Tag, attrsString(n.Attrs), n.HasText, n.Text)
	}
	return fmt.Sprintf("Node(tag: %s, attrs: %s, hasText: %v, text: %s, children: %v)", n.Tag, attrsString(n.Attrs), n.HasText, n.Text, n.String())
}
