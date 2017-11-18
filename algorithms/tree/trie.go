package tree

type SimpleTrieNode struct {
	Children map[rune]*SimpleTrieNode
}

func newSimpleTrieNode() *SimpleTrieNode {
	n := new(SimpleTrieNode)
	n.Children = make(map[rune]*SimpleTrieNode)
	return n
}

type SimpleTrie struct {
	root *SimpleTrieNode
}

func NewSimpleTrie() *SimpleTrie {
	return &SimpleTrie{root: newSimpleTrieNode()}
}

func (t *SimpleTrie) InsertRoot(str string) {
	node := t.root
	for _, u := range str {
		if _, ok := node.Children[u]; !ok {
			node.Children[u] = newSimpleTrieNode()
		}
		node = node.Children[u]
	}
}

func fs(result *[]string, node *SimpleTrieNode, prefix []rune) {
	if len(node.Children) == 0 {
		*result = append(*result, string(prefix))
		return
	}
	for u, child := range node.Children {
		fs(result, child, append(prefix, u))
	}
}

func (t *SimpleTrie) FuzzySearch(str string) (result []string) {
	node := t.root
	for _, u := range str {
		if node.Children[u] == nil {
			return
		}
		node = node.Children[u]
	}
	fs(&result, node, []rune(str))
	return
}
