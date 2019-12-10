/*
	编写一个程序，找到两个单链表相交的起始节点
*/

/*
	思路：求出两个链表的长度差，然后让指向长的的链表的指针先走长度差个单位，然后指向两个链表的指针同时，当两个指针相等时，代表两个链表的相交点。
*/

package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 定义两个指针，初始化指向两个链表的表头
	pointA := headA
	pointB := headB
	// 定义两个常量，初始化为空，用于记录长度差
	countA := 0
	countB := 0
	// 对两个两个链表进行遍历，遍历条件为两个链表不同时为空，当其中一个为空时，另外一个链表继续遍历，由count_a或者count_b记录长度差，分别对应的链表
	for pointA != nil || pointB != nil {
		// 指针若不为空，继续往下走，否则，将长度差常量进行加1操作，下同。
		if pointA != nil {
			pointA = pointA.Next
		} else {
			countA++
		}
		if pointB != nil {
			pointB = pointB.Next
		} else {
			countB++
		}
	}
	// 重新定义指针，初始化为两个链表的表头
	pointA = headA
	pointB = headB
	// 判断记录长度差的常量是否为0，若不为0则代表相应的链表长度较短，对指向另外一个链表表头的指针进行走长度差个单位操作
	if countA != 0 {
		count := 0
		for count < countA {
			pointB = pointB.Next
			count++
		}
	} else if countB != 0 {
		count := 0
		for count < countB {
			pointA = pointA.Next
			count++
		}
	}
	// 此时两指针已经相对同步，可让其同时往下走，直到两个指针相等。
	for pointA != nil && pointB != nil {
		if pointB == pointA {
			return pointB
		}
		pointA = pointA.Next
		pointB = pointB.Next
	}
	return nil
}
