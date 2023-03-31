package to_offer

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newHead *ListNode
	var p *ListNode

	for l1 != nil || l2 != nil {
		if l1 == nil {
			if newHead == nil {
				return l2
			}

			p.Next = l2
			return newHead
		}

		if l2 == nil {
			if newHead == nil {
				return l1
			}

			p.Next = l1
			return newHead
		}

		val := l1.Val
		if l2.Val < val {
			val = l2.Val
			l2 = l2.Next
		} else {
			l1 = l1.Next
		}

		node := ListNode{
			Val: val,
		}

		if newHead == nil {
			newHead = &node
			p = newHead
		} else {
			p.Next = &node
			p = p.Next
		}
	}

	return newHead
}
