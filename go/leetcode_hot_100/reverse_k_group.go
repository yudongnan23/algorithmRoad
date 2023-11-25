package leetcode_hot_100

// TODO again
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{}
	newHead.Next = head

	reverseFunc := func(l *ListNode) (*ListNode, *ListNode, *ListNode) {
		curP := l
		n := 0
		var (
			curHead *ListNode
			tail    *ListNode
		)

		for curP != nil && n < k {
			node := &ListNode{
				Val: curP.Val,
			}

			node.Next = curHead
			curHead = node
			if tail == nil {
				tail = node
			}

			curP = curP.Next
			n++
		}
		if n != k {
			return l, nil, nil
		}
		return curHead, tail, curP
	}

	p := head
	q := newHead

	for p != nil {
		curHead, tail, next := reverseFunc(p)
		q.Next = curHead
		if next == nil {
			break
		}
		q = tail
		p = next
	}

	return newHead.Next
}
