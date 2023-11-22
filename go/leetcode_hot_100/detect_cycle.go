package leetcode_hot_100

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for slow != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}

	return nil
}
