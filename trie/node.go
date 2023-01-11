package trie

type Node struct {
	End      bool
	Children map[rune]*Node
}

func NewNode() *Node {
	return &Node{Children: make(map[rune]*Node), End: false}
}

func (n *Node) Exists(prefix rune) bool {
	_, ok := n.Children[prefix]
	return ok
}

func (n *Node) Child(prefix rune) *Node {
	if node, ok := n.Children[prefix]; ok {
		return node
	}

	return nil
}
