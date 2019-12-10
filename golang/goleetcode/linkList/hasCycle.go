/*
	给定一个链表，判断链表中是否有环。
	为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置
	（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
*/

/*
	思路：判断链表中是否有环，定义快慢两个指针，快指针一次走两步，慢指针一次走一步，当两指针相遇时，代表链表中有环
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	result := false
	// 定义快慢指针，初始化为链表头结点
	slow := head
	quick := head
	// 指针往前走的必要条件，否则会抛出异常
	for slow.Next != nil && quick.Next != nil && quick.Next.Next != nil {
		// 慢指针走一步
		slow = slow.Next
		// 快指针走两步
		quick = quick.Next.Next
		// 两指针相等，即相遇
		if slow == quick {
			result = true
			break
		}
	}
	return result
}

// main函数测试用例
func main() {
	head := ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next = head.Next
	fmt.Println(hasCycle(&head))
}
