/*
	给定一个链表，判断链表中是否有环。
	为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置
	（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	result := false
	slow := head
	quick := head
	for slow.Next != nil && quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
		if slow == quick {
			result = true
			break
		}
	}
	return result
}

func main() {
	head := ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next = head.Next
	fmt.Println(hasCycle(&head))
}
