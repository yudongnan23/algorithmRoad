package leetcode_hot_100

// NodeII
// TODO again
type NodeII struct {
	End  bool
	Val  byte
	Sons map[byte]*NodeII
}

func (n *NodeII) Search(word string) (*NodeII, bool) {
	if len(word) == 0 {
		return n, true
	}
	node, ok := n.Sons[word[0]]
	if !ok {
		return nil, false
	}

	return node.Search(word[1:])
}

func (n *NodeII) Insert(word string) {
	if len(word) == 0 {
		return
	}

	end := len(word) == 1

	node, ok := n.Sons[word[0]]
	if ok {
		if end {
			node.End = true
			return
		}
		node.Insert(word[1:])
		return
	}

	newNode := &NodeII{
		Val:  word[0],
		Sons: make(map[byte]*NodeII, 0),
	}

	n.Sons[word[0]] = newNode

	if end {
		newNode.End = true
		return
	}

	newNode.Insert(word[1:])
}

type Trie struct {
	Nodes map[byte]*NodeII
}

func ConstructorII() Trie {
	return Trie{
		Nodes: make(map[byte]*NodeII, 0),
	}
}

func (t *Trie) Insert(word string) {
	node, ok := t.Nodes[word[0]]
	if ok {
		node.Insert(word[1:])
		return
	}
	newNode := &NodeII{
		Val:  word[0],
		Sons: make(map[byte]*NodeII, 0),
	}

	t.Nodes[word[0]] = newNode

	if len(word) == 1 {
		newNode.End = true
		return
	}
	newNode.Insert(word[1:])
}

func (t *Trie) Search(word string) bool {
	node, ok := t.Nodes[word[0]]
	if !ok {
		return false
	}
	node, ok = node.Search(word[1:])
	if !ok {
		return false
	}
	return node.End
}

func (t *Trie) StartsWith(word string) bool {
	node, ok := t.Nodes[word[0]]
	if !ok {
		return false
	}
	_, ok = node.Search(word[1:])
	return ok
}
