package to_offer

type CQueue struct {
	Head *Node
	Tail *Node
}

func Constructor() CQueue {
	return CQueue{
		Head: nil,
		Tail: nil,
	}
}

func (cq *CQueue) AppendTail(val int) {
	node := &Node{
		Val: val,
	}

	if cq.Head == nil {
		cq.Head = node
		cq.Tail = node
		return
	}

	cq.Tail.Next = node
	cq.Tail = cq.Tail.Next
	return
}

func (cq *CQueue) DeleteHead() int {
	if cq.Head == nil {
		return -1
	}

	node := cq.Head
	cq.Head = cq.Head.Next

	if node == cq.Tail {
		cq.Tail = nil
	}

	return node.Val
}
