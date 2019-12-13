/*
	给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 */

/*
	思路： 典型的双指针解法，新建一个链表结点node，将其next指向head头结点，定义两个指针，left和right，left指针初始化node，
		  right指针初始化head头指针；让right指针先走n步，然后两指针同时走，直到right指针到达链表尾部，此时left指针刚好在
		  倒数第n个结点的前一个节点，此时便可将倒数第n个结点删除。
 */

package main


func removeNthFromEnd(head *ListNode, n int) *ListNode{
	// 新建node结点，并将其Next指向head头结点
	node := ListNode{Val: -1}
	node.Next = head
	// 初始化left指针为node结点
	left := &node
	// 初始化right指针为head头结点
	right := head
	count := 0
	// 让指针走n步，注意倒数没有0
	for ; count < n - 1 && right != nil; count++{
		right = right.Next
	}
	// 两指针同时走，直到right指针到达链表尾部
	for ; right.Next != nil; {
		left = left.Next
		right = right.Next
	}
	// 如果两指针没动，即需要删除链表的头结点，只需将head结点指针指向head.Next即可
	if left.Next == head{
		head = head.Next
	}else{
		// 执行结点删除
		left.Next = left.Next.Next
	}
	return head
}
