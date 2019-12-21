/*
	给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
 */

/*
    思路： 由题意可知，链表向右移动链表长度个单位，将会得到原链表。则可求出链表长度，使用需要移动的位置个数k对链表长度进行求余step，余数即为需要移动的位置。
            得到正真需要移动的位置数量，则可使用双指针的方法，首先定义一个新的节点，将该节点的next指向head头接单，然后定义left指针与right指针，初始化都指针新的节点
            然后让right指针走step步，此时再让left指针与right指针同时走，直到right.next为空，最后将left与left.next结点断开，将right结点连接至原链表的头结点
            此时left.next即为新链表的头结点；之所以将left与right往前提一步，是为了方便后续的链表的断开与连接；
 */


package main

func rotateRight(head *ListNode, k int) *ListNode {
    // 当链表为空或者k为0时，直接返回head头结点即可
	if head == nil || k == 0{
		return head
	}
	// 调用方法获取链表长度
	length := getLength(head)
	// 使用k对链表进行求余操作，得到链表实际上需要移动的位置的数量
	step := k % length
	// 当余数为0时，代表不需要移动
	if step == 0{
		return head
	}
	// 定义一个新的节点，将新的节点的next指向链表头结点
	preNode := &ListNode{Val: -1}
	preNode.Next = head
	// 定义left指针与right指针
	left := preNode
	right := preNode
	count := 0
	// 两指针同时走，直到right指针走到链表尾结点
	for ; count < step; {
		right = right.Next
	}

    // 两指针同时走，直到right指针走到链表尾结点
	for ; right.Next != nil; {
		left = left.Next
		right = right.Next
	}

	// 将left结点与left.next结点断开，right与head进行连接，新链表结点为left.next，即下面的cur
	cur := left.Next
	newHead := cur
	left.Next = nil
	right.Next = head
	return newHead

}

func getLength(head *ListNode) int {
	length := 0
	point := head
	for ; point != nil; {
		length ++
		point = point.Next
	}
	return length
}
