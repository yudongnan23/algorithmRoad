package leetcode_hot_100

// TODO three
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := findMid(head)

	left := sortList(head)
	right := sortList(mid)

	return mergeTwoLists(left, right)
}

func findMid(head *ListNode) *ListNode {
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return mid
}
