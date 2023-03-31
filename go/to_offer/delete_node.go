package to_offer

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	left := head
	right := head.Next
	for right != nil {
		if right.Val == val {
			left.Next = right.Next
			break
		}

		left = left.Next
		right = right.Next
	}

	return head
}
