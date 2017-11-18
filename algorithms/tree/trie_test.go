package tree

import (
	"fmt"
	"testing"
)

func TestNewTrieNode(t *testing.T) {
	tn := newSimpleTrieNode()
	fmt.Println(len(tn.Children) == 128)
}

func TestNewSimpleTrie(t *testing.T) {
	simpleTrie := NewSimpleTrie()
	simpleTrie.InsertRoot("foo")
	simpleTrie.InsertRoot("foz")
	simpleTrie.InsertRoot("fzz")
	result := simpleTrie.FuzzySearch("fo")
	fmt.Println(result)
}
