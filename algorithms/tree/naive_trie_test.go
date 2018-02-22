package tree

import (
	"fmt"
	"testing"
)

func TestNewTrieNode(t *testing.T) {
	tn := NewNaiveTrie()
	fmt.Println(len(tn.Children) == 128)
}

func TestNewSimpleTrie(t *testing.T) {
	naiveTrie := NewNaiveTrie()
	naiveTrie.Insert("foo")
	naiveTrie.Insert("foz")
	naiveTrie.Insert("fzz")
	result := naiveTrie.PrefixSearch("fo")
	fmt.Println(result)
}
