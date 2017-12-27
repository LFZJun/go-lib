package tree

import (
	"errors"
	"strings"
)

type (
	Nodes []*Node

	Node struct {
		FullPath string
		Key      string
		Value    interface{}
		Root     bool
		Children Nodes
	}
)

func (nodes Nodes) FuzzyFind(k string, ns *Nodes) {
	for _, n := range nodes {
		lenKey, lenK := len(n.Key), len(k)
		if lenKey > lenK {
			// foo fo
			if strings.HasPrefix(n.Key, k) {
				*ns = append(*ns, n)
			}
			continue
		}
		if lenKey < lenK {
			// fo foo
			if strings.HasPrefix(k, n.Key) {
				n.Children.FuzzyFind(k[lenKey:], ns)
			}
			continue
		}
		// foo foo
		if k == n.Key {
			*ns = append(*ns, n)
		}
	}
}

func (nodes Nodes) Find(k string) *Node {
	for _, n := range nodes {
		if !strings.HasPrefix(k, n.Key) {
			continue
		}
		if len(k) == len(n.Key) {
			return n
		}

		child := n.Children.Find(k[len(n.Key):])
		if child == nil {
			continue
		}
		return child
	}
	return nil
}

func (nodes *Nodes) Add(k string, v interface{}, root bool) (err error) {
	return nodes.addHasPrefix(k, v, root, "")
}

//    key k
// 1. foo foz 有公共项分裂 fo -o -z
// 2. foo bar 没公共项继续遍历 || 在当前节点上添加
// 3. foz fo  有公共项分裂 fo -z
// 4. fo  foz 子节点递归
// 5. foo foo 返回空
func (nodes *Nodes) addHasPrefix(k string, v interface{}, root bool, prefix string) (err error) {
loop:
	for _, n := range *nodes {
		key := n.Key
		minLen := len(key)
		if minLen > len(k) {
			minLen = len(k)
		}

		for i := 0; i < minLen; i++ {
			// 1. foo foz
			if key[i] == k[i] {
				continue
			}
			// 2. foo bar
			if i == 0 {
				continue loop
			}
			// 1. foo foz
			// fo
			// o   z
			fullPath := n.FullPath[:len(n.FullPath)-minLen+i]
			*n = Node{
				FullPath: fullPath,
				Key:      key[:i],
				Children: Nodes{
					{
						FullPath: fullPath + key[i:],
						Key:      key[i:],
						Value:    n.Value,
						Children: n.Children,
					},
					{
						FullPath: fullPath + k[i:],
						Key:      k[i:],
						Value:    v,
					},
				},
				Root: n.Root,
			}
			return
		}

		lenKey, lenK := len(key), len(k)
		// 3. foz fo
		if lenKey > lenK {
			fullPath := n.FullPath[:len(n.FullPath)-lenKey+lenK]
			*n = Node{
				FullPath: fullPath,
				Key:      key[:lenK],
				Value:    v,
				Children: Nodes{
					{
						FullPath: fullPath + key[lenK:],
						Key:      key[lenK:],
						Value:    n.Value,
						Children: n.Children,
					},
				},
				Root: n.Root,
			}
			return
		}
		// 4. fo foz
		if lenKey < lenK {
			keyToAdd := k[lenKey:]
			err = n.Children.addHasPrefix(keyToAdd, v, false, n.FullPath)
			return
		}
		// 5. foo foo
		if v == nil {
			return
		}
		if n.Value != nil {
			return errors.New("duplicate")
		}
		n.Value = v
		return
	}

	n := &Node{
		FullPath: prefix + k,
		Key:      k,
		Value:    v,
		Root:     root,
	}
	*nodes = append(*nodes, n)
	return
}

func (n *Node) Size() (l int) {
	for _, node := range n.Children {
		l += 1 + node.Size()
	}
	return
}
