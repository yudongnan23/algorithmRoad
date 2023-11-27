package leetcode_hot_100

// LRUCache
// TODO again
type LRUCache struct {
	head     *NodeX
	tail     *NodeX
	mmap     map[int]*NodeX
	capacity int
}

type NodeX struct {
	Val  int
	Key  int
	Next *NodeX
	Pre  *NodeX
}

func (n *NodeX) independence() {
	n.Next = nil
	n.Pre = nil
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		head:     nil,
		tail:     nil,
		mmap:     make(map[int]*NodeX, 0),
		capacity: capacity,
	}
}

func (l *LRUCache) addToHead(node *NodeX) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}
	node.Next = l.head
	l.head.Pre = node
	l.head = node
}

func (l *LRUCache) reset() {
	l.head = nil
	l.tail = nil
}

func (l *LRUCache) deleteNode(node *NodeX) {
	if node == l.head && node == l.tail {
		l.reset()
		return
	}
	if node == l.head {
		l.head = l.head.Next
		l.head.Pre = nil
		return
	}
	if node == l.tail {
		l.tail = l.tail.Pre
		l.tail.Next = nil
		return
	}

	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (l *LRUCache) removeTail() {
	if l.tail == nil {
		return
	}
	if l.tail == l.head {
		l.reset()
		return
	}

	l.tail = l.tail.Pre
	l.tail.Next = nil
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.mmap[key]
	if !ok {
		return -1
	}

	l.deleteNode(node)
	node.independence()
	l.addToHead(node)
	return node.Val
}

func (l *LRUCache) Put(key int, value int) {
	node, ok := l.mmap[key]
	if ok {
		node.Val = value
		if node == l.head {
			return
		}
		l.deleteNode(node)
		node.independence()
		l.addToHead(node)
		return
	}
	node = &NodeX{
		Key: key,
		Val: value,
	}
	l.mmap[key] = node
	l.addToHead(node)

	if len(l.mmap) > l.capacity {
		delete(l.mmap, l.tail.Key)
		l.removeTail()
	}
}
