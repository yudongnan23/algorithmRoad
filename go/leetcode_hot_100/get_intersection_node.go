package leetcode_hot_100

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	changeA := false
	changeB := false

	a := headA
	b := headB

	for {
		if a == b {
			return a
		}

		a = a.Next
		b = b.Next

		if (a == nil && changeA) || (b == nil && changeB) {
			return nil
		}

		if a == nil {
			changeA = true
			a = headB
		}

		if b == nil {
			changeB = true
			b = headA
		}
	}
}
