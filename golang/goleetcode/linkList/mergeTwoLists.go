/*
	将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/

/*
	思路： 定义四个指针，指针point1，用于遍历l1链表，指针point2，用于遍历l2，指针point，用于指向新链表的尾结点，指针newHead用于指向新链表头结点
			对两个链表进行遍历，将两指针指向的结点结点值较小的一个结点构造成链表结点连接至新链表尾结点，然后重新进行下一轮遍历与比较。
 */

package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode)*ListNode{
	point1 := l1
	point2 := l2
	var newHead *ListNode
	point := newHead
	// 当两指针其一不为nil时
	for ; point1 != nil || point2 != nil; {
		// 当两指针都不为nil时，代表point1与point2指向的结点都有值
		if point1 != nil && point2 != nil{
			val1 := point1.Val
			val2 := point2.Val
			// 当val1小于val2时，将val1构造成结点放置新链表尾部
			if val1 <= val2{
				node := &ListNode{Val: val1}
				// 判断新链表头结点是否为nil指针，若是，则将新链表头结点与尾结点指针初始化
				if newHead == nil{
					newHead = node
					point = newHead
				// 若新链表头结点不为nil，则进行尾部连接操作
				}else{
					point.Next = node
					point = point.Next
				}
				point1 = point1.Next
			// 当val2小于val1时，将val2构造成结点放置新链表尾部
			}else{
				node := &ListNode{Val: val2}
				if newHead == nil{
					newHead = node
					point = newHead
				}else{
					point.Next = node
					point = point.Next
				}
				point2 = point2.Next
			}
		// 当l2已遍历完而l1未遍历完时，直接将point1指针连接至新链表尾部
		}else if point1 != nil{
			// 当point指针为nil时，代表l2链表为空，此时直接将newHead初始化为point1即可
			if point != nil{
				point.Next = point1
			// 若point不为空，则直接进行尾部连接操作
			}else{
				newHead = point1
			}
			// 直接退出当前循环
			break
		// 当l1已遍历完而l2未遍历完时，直接将point2指针连接至新链表尾部
		}else if point2 != nil{
			// 当point指针为nil时，代表l1链表为空，此时直接将newHead初始化为point2即可
			if point != nil{
				point.Next = point2
			// 若point不为空，则直接进行尾部连接操作
			}else{
				newHead = point2
			}
			// 直接退出当前循环
			break
		}
	}
	return newHead
}