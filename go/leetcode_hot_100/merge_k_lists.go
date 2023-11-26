package leetcode_hot_100

// TODO again
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeLists(lists[0], lists[1])
	}
	return mergeLists(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
}

func mergeLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var newHead *ListNode
	var p *ListNode
	for l1 != nil && l2 != nil {
		node := l2
		if less(l1, l2) {
			node = l1
			l1 = l1.Next
		} else {
			l2 = l2.Next
		}
		if newHead == nil {
			newHead = node
			p = node
			continue
		}
		p.Next = node
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}

	if l2 != nil {
		p.Next = l2
	}

	return newHead
}
