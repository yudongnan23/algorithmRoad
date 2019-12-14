/*
	删除链表中等于给定值 val 的所有节点。
 */

/*
	思路：新建一结点，将其next指向链表头结点，初始化一个指针为该新建的结点，然后让指针往下走，当其Next结点的结点值等于目标值时，
         将其Next结点删除即可，最后返回新建结点的Next结点即可。
 */

package main

func removeElements(head *ListNode, val int) *ListNode {
	// 新建一个结点
	node := &ListNode{Val: -1}
	node.Next = head
	// 初始化指针为新建的结点
	point := node
	// 让指针往下走，直到指针的Next指针为空，代表指针已走到链表尾部
	for ; point.Next != nil; {
		// 当指针Next结点的节点值为目标值，执行删除删除操作
		if point.Next.Val == val{
			point.Next = point.Next.Next
		// 否则指针往前走
		}else {
			point = point.Next
		}
	}
	// 得到链表新的头结点并返回
	head = node.Next
	return head
}