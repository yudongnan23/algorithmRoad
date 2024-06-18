package leetcode_hot_100

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		head = &ListNode{}
		p    = head
	)

	for list1 != nil && list2 != nil {
		node := list1
		if node.Val > list2.Val {
			node = list2
			list2 = list2.Next
		} else {
			list1 = list1.Next
		}

		p.Next = node
		p = p.Next
	}

	if list1 == nil {
		p.Next = list2
	}
	if list2 == nil {
		p.Next = list1
	}

	return head.Next
}
