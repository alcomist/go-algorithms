package graph

import (
	"bytes"
	"fmt"
	"strings"
)

type Trie struct {
	root *node
}

type node struct {
	sub      map[rune]*node
	terminal bool
}

func newNode() *node {
	return &node{sub: make(map[rune]*node)}
}

func NewTrie() *Trie {
	return &Trie{root: newNode()}
}

func nodeStrings1(n *node, rs []rune, ss *[]string) {

	if n.terminal {
		*ss = append(*ss, string(rs))
	}

	for k, v := range n.sub {
		rs = append(rs, k)
		nodeStrings1(v, rs, ss)
		rs = rs[:len(rs)-1]
	}
}

func nodeStrings2(n *node, rs []rune) []string {

	var ss []string
	if n.terminal {
		ss = append(ss, string(rs))
	}

	for k, v := range n.sub {
		rs = append(rs, k)
		ss = append(ss, nodeStrings2(v, rs)...)
		rs = rs[:len(rs)-1]
	}

	return ss
}

func (t *Trie) String() string {

	var b bytes.Buffer

	//ss := make([]string, 0)
	//nodeStrings1(t.root, []rune{}, &ss)

	ss := nodeStrings2(t.root, []rune{})

	fmt.Fprintf(&b, "WORDS = %v\n", strings.Join(ss, ", "))
	return b.String()
}

func (t *Trie) Insert(s string) {

	node := t.root

	for _, ch := range s {
		n := node.sub[ch]
		if n == nil {
			n = newNode()
			node.sub[ch] = n
		}
		node = n
	}

	node.terminal = true
}

func (t *Trie) Has(s string) bool {

	node := t.root

	for _, ch := range s {
		n := node.sub[ch]
		if n == nil {
			return false
		} else {
			node = n
		}
	}

	return node.terminal == true
}

func (t *Trie) HasPrefix(s string) bool {

	node := t.root

	for _, ch := range s {
		n, ok := node.sub[ch]
		if !ok {
			return false
		} else {
			node = n
		}
	}

	return true
}
