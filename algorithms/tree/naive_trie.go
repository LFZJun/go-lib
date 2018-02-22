package tree

type NaiveTrie struct {
	Children map[rune]*NaiveTrie
}

func (n *NaiveTrie) Insert(str string) {
	node := n
	for _, u := range str {
		if _, ok := node.Children[u]; !ok {
			node.Children[u] = NewNaiveTrie()
		}
		node = node.Children[u]
	}
}

func NewNaiveTrie() *NaiveTrie {
	return &NaiveTrie{
		Children: make(map[rune]*NaiveTrie),
	}
}

func (n *NaiveTrie) PrefixSearch(str string) (result []string) {
	node := n
	for _, u := range str {
		if node.Children[u] == nil {
			return
		}
		node = node.Children[u]
	}
	prefixSearch(&result, node, []rune(str))
	return
}

func prefixSearch(result *[]string, node *NaiveTrie, prefix []rune) {
	if len(node.Children) == 0 {
		*result = append(*result, string(prefix))
		return
	}
	for u, child := range node.Children {
		prefixSearch(result, child, append(prefix, u))
	}
}
