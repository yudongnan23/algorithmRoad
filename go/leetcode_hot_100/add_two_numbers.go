package leetcode_hot_100

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newHead *ListNode
	var p *ListNode
	var more bool
	for l1 != nil || l2 != nil {
		var res int
		if l1 != nil {
			res = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			res = res + l2.Val
			l2 = l2.Next
		}
		if more {
			res = res + 1
		}

		more = false
		if res > 9 {
			more = true
		}

		node := ListNode{
			Val: res % 10,
		}

		if newHead == nil {
			newHead = &node
			p = &node
		} else {
			p.Next = &node
			p = p.Next
		}
	}

	if more {
		p.Next = &ListNode{
			Val: 1,
		}
	}

	return newHead
}
