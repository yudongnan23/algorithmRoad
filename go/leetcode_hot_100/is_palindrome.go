package leetcode_hot_100

// TODO three - 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}

	left := head
	right := head.Next
	left.Next = nil

	for i := 1; i < count/2; i++ {
		temp := right.Next
		right.Next = left
		left = right
		right = temp
	}

	if count%2 == 1 && right != nil {
		right = right.Next
	}

	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}
