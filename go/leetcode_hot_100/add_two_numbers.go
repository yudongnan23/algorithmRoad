package leetcode_hot_100

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newHead *ListNode
	p := newHead

	flag := false

	createNewNode := func(a, b *ListNode) *ListNode {
		var val int
		if a != nil {
			val = a.Val
		}
		if b != nil {
			val = val + b.Val
		}
		if flag {
			val++
			flag = false
		}
		if val >= 10 {
			flag = true
			val = val - 10
		}

		return &ListNode{
			Val: val,
		}
	}

	for l1 != nil || l2 != nil {
		newNode := createNewNode(l1, l2)
		if newHead == nil {
			newHead = newNode
			p = newNode
		} else {
			p.Next = newNode
			p = p.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if flag {
		p.Next = &ListNode{
			Val: 1,
		}
	}

	return newHead
}
