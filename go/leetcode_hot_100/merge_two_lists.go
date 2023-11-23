package leetcode_hot_100

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head *ListNode
	p := head
	for l1 != nil && l2 != nil {
		var node *ListNode
		if less(l1, l2) {
			node = l1
			l1 = l1.Next
		} else {
			node = l2
			l2 = l2.Next
		}

		if head == nil {
			p = node
			head = node
			continue
		}

		p.Next = node
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}

	if l2 != nil {
		p.Next = l2
	}

	return head
}

func less(node1, node2 *ListNode) bool {
	if node1.Val < node2.Val {
		return true
	}
	return false
}
