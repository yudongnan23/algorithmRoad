package to_offer

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	minNode, maxNode := sortNode(l1, l2)
	newHead := minNode

	p1 := minNode.Next
	p2 := maxNode

	p := newHead

	for {
		if p1 == nil {
			p.Next = p2
			break
		}

		if p2 == nil {
			p.Next = p1
			break
		}

		minNode, maxNode = sortNode(p1, p2)

		p.Next = minNode
		p = p.Next

		p1 = maxNode
		p2 = minNode.Next
	}

	return newHead

}

func sortNode(l1, l2 *ListNode) (*ListNode, *ListNode) {
	if l1.Val > l2.Val {
		return l2, l1
	}
	return l1, l2
}
