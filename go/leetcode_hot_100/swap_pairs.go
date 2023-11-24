package leetcode_hot_100

// TODO again
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	r := head.Next
	l := head

	newHead := &ListNode{}
	newHead.Next = head
	p := newHead

	for r != nil {
		next := r.Next

		r.Next = l
		l.Next = next
		p.Next = r

		if next == nil {
			break
		}
		l = next
		r = next.Next
		p = p.Next.Next
	}

	return newHead.Next
}
