package to_offer

func reverseList(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(left *ListNode, right *ListNode) *ListNode {
	if right == nil {
		return left
	}
	nextRight := right.Next
	right.Next = left
	nextLeft := right
	return reverse(nextLeft, nextRight)
}
