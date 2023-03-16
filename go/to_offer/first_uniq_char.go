package to_offer

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return " "[0]
	}

	ll := NewLinkedList()

	for i := 0; i < len(s); i++ {
		ll.Append(s[i])
	}

	if ll.head != nil {
		return ll.head.v
	}

	return " "[0]
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:        nil,
		tail:        nil,
		nodeMapping: make(map[byte]*doubleLinkedListNode, 0),
	}
}

func (ll *LinkedList) Append(v byte) {
	node, ok := ll.nodeMapping[v]
	if !ok {
		node = &doubleLinkedListNode{
			v: v,
		}

		ll.nodeMapping[v] = node

		if ll.head == nil {
			ll.head = node
			ll.tail = node
			return
		}

		node.pre = ll.tail
		ll.tail.next = node
		ll.tail = ll.tail.next
		return
	}

	if node.isDel {
		return
	}

	node.isDel = true

	if node == ll.head && node == ll.tail {
		ll.head = nil
		ll.tail = nil
		return
	}

	if node == ll.head {
		ll.head = ll.head.next
		return
	}

	if node == ll.tail {
		ll.tail = ll.tail.pre
		ll.tail.next = nil
		return
	}

	preNode := node.pre
	nextNode := node.next
	preNode.next = node.next
	nextNode.pre = preNode
}
