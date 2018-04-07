package trie

type Node struct {
	value    rune
	terminal bool
	children map[rune]*Node
}

type Trie struct {
	root *Node
}

func (t *Trie) insert(value string) {
	node := t.root
	for pos, char := range value {
		if _, ok := node.children[char]; !ok {
			node.children[char] = &Node{
				value:    char,
				children: make(map[rune]*Node),
			}
		}

		node = node.children[char]
		if pos == len(value)-1 {
			node.terminal = true
		}
	}
}

func (t *Trie) find(value string) (*Node, bool) {
	node := t.root
	for _, char := range value {
		if _, ok := node.children[char]; !ok {
			return nil, false
		}

		node = node.children[char]
	}

	return node, true
}

func New() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[rune]*Node),
		},
	}
}
