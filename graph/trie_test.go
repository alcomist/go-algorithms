package graph

import (
	"testing"
)

func TestTrie(t *testing.T) {

	trie := NewTrie()
	trie.Insert("")
	trie.Insert("sea")
	trie.Insert("saw")
	trie.Insert("seen")
	trie.Insert("see")
	trie.Insert("season")
	trie.Insert("")
	trie.Insert("")

	var tests1 = []struct {
		input string
		want  bool
	}{
		{"sea", true},
		{"sean", false},
	}

	for _, test := range tests1 {
		if got := trie.Has(test.input); got != test.want {
			t.Errorf("trie.Has(%q) = %v (WANT:%v)", test.input, got, test.want)
		}
	}

	var tests2 = []struct {
		input string
		want  bool
	}{
		{"seaso", true},
		{"sean", false},
	}

	for _, test := range tests2 {
		if got := trie.HasPrefix(test.input); got != test.want {
			t.Errorf("trie.HasPrefix(%q) = %v (WANT:%v)", test.input, got, test.want)
		}
	}
}
