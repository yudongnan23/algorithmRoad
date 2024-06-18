package leetcode_hot_100

// Trie
// TODO three - 前缀树
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func ConstructorII() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return true
}
