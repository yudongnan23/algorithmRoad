package to_offer

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	old2NewMapping := make(map[*Node]*Node, 0)
	new2OldMapping := make(map[*Node]*Node, 0)
	randomMapping := make(map[*Node]*Node, 0)

	newHead := Node{
		Val: head.Val,
	}

	new2OldMapping[&newHead] = head
	old2NewMapping[head] = &newHead
	if head.Random != nil {
		randomMapping[head] = head.Random
	}

	newP := &newHead
	p := head.Next

	for p != nil {
		newNode := Node{
			Val: p.Val,
		}
		newP.Next = &newNode
		old2NewMapping[p] = &newNode
		new2OldMapping[&newNode] = p

		if p.Random != nil {
			randomMapping[p] = p.Random
		}

		newP = newP.Next
		p = p.Next
	}

	newP = &newHead
	for newP != nil {
		random, ok := randomMapping[new2OldMapping[newP]]
		if ok {
			newP.Random = old2NewMapping[random]
		}

		newP = newP.Next
	}

	return &newHead
}
