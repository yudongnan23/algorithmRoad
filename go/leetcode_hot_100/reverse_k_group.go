package leetcode_hot_100

// TODO three
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	canReverse, endNode := canReverse(head, k)
	if !canReverse || head == nil {
		return head
	}

	right := head.Next
	left := head
	next := endNode.Next
	for right != endNode {
		temp := right.Next
		right.Next = left
		left = right
		right = temp
	}

	right.Next = left
	head.Next = reverseKGroup(next, k)

	return right
}

func canReverse(head *ListNode, k int) (bool, *ListNode) {
	p := head
	count := 1
	for p != nil && count < k {
		count++
		p = p.Next
	}
	return count == k && p != nil, p
}
