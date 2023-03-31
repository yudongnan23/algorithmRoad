package to_offer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	right := head
	for i := 0; i < k; i++ {
		if right == nil {
			return nil
		}

		right = right.Next
	}

	left := head
	for right != nil {
		left = left.Next
		right = right.Next
	}

	return left
}
