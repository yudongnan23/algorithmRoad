package leetcode_hot_100

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	isCarry := false
	head := &ListNode{}
	p := head

	for l1 != nil || l2 != nil {
		val := 0
		if l1 != nil {
			val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val = val + l2.Val
			l2 = l2.Next
		}
		if isCarry {
			isCarry = false
			val++
		}

		if val >= 10 {
			val = val - 10
			isCarry = true
		}

		p.Next = &ListNode{Val: val}
		p = p.Next
	}

	if isCarry {
		p.Next = &ListNode{Val: 1}
	}

	return head.Next
}
