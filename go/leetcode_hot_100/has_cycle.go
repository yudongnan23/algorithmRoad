package leetcode_hot_100

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
