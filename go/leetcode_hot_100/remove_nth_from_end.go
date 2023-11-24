package leetcode_hot_100

// TODO again
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{}
	newHead.Next = head
	r := newHead
	count := 1
	for count <= n && r != nil {
		count++
		r = r.Next
	}

	if r == nil {
		return nil
	}

	l := newHead
	for r.Next != nil {
		r = r.Next
		l = l.Next
	}

	l.Next = l.Next.Next

	return newHead.Next
}
