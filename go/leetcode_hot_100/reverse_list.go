package leetcode_hot_100

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var ahead *ListNode
	pre := head

	for pre.Next != nil {
		temp := pre.Next
		pre.Next = ahead
		ahead = pre
		pre = temp
	}

	pre.Next = ahead

	return pre
}
