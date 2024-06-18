package leetcode_hot_100

// TODO again
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	temp := head.Next.Next
	right := head.Next
	right.Next = head
	head.Next = swapPairs(temp)
	return right
}
