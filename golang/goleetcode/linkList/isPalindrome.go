/*
	请判断一个链表是否为回文链表。
 */

/*
	思路：定义一个函数使用快慢指针找到链表的中点，再定义一个函数将中点后面的结点进行反转，并返回反转后的头结点
		  对链表进行双向同时遍历按，当遇两个节点的节点值不相等时，代表不是该链表不是回文链表，否则是回文链表。
 */
package main

func isPalindrome(head *ListNode) bool{
	// 调用函数获取链表中点
	middleNode := getMiddleNode(head)
	// 调用反转函数对中点后面的结点进行反转并返回反转后的头结点
	newHead := reverse(middleNode)
	// 左指针
	pointLeft := head
	// 右指针
	pointRight := newHead
	result := true
	// 两指针同时走，一次走一步
	for ; pointLeft != nil && pointRight != nil; {
		// 当遇到两节点值不相等时，代表该链表不是回文链表。
		if pointLeft.Val != pointRight.Val{
			result = false
			break
		}
		pointRight = pointRight.Next
		pointLeft = pointLeft.Next
	}
	return result
}

func reverse(middle *ListNode) *ListNode{
	if middle == nil{
		return middle
	}
	// 递归函数出口，即递归到链表尾结点
	if middle.Next == nil{
		return middle
	}
	// 递归调用本身对链表的每个节点进行反转
	newHead := reverse(middle.Next)
	// 当前结点与next结点进行反转
	cur := middle
	middle.Next.Next = cur
	middle.Next = nil
	return newHead
}

func getMiddleNode(head *ListNode) *ListNode{
	// 定义慢指针，一次走一步
	slow := head
	// 定义快指针，一次走两步
	fast := head
	// 当快指针走到链表尾部时，慢指针刚好在中点位置
	for ; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

