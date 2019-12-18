/*
	给出两个非空的链表用来表示两个非负的整数。其中,它们各自的位数是按照逆序的方式存储的，并且它们的每个节点只能存储一位数字。
	如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
	您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 */


/*
	思路：定义两个指针初始化分别指向两个链表的头结点，同时遍历两个链表，将指针指向的结点的节点值相加，若大于10，则两指针下一步结点值相加
			多加1，否则不加1；将相加得到的数构造成一个链表结点，依次连接。
 */

package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 用于遍历l1的指针
	point1 := l1
	// 用于遍历l2的指针
	point2 := l2
	// 用于告知下一次遍历是否需要加1的token，初始为false
	token := false
	// 指向新链表的头结点的指针
	var newHead *ListNode
	// 指向新链表的尾结点的指针
	var point *ListNode
	// 当point1、point2、token其中之一不为空时，需要对新链表进行构造
	for ; point1 != nil || point2 != nil || token; {
		// 当point1与point2都不为空时，需要获取当前point1与point2指向的结点的结点值，
		// 并根据token的true或者false进行加1操作，然后构造新链表结点并连接至新链表尾结点指针
		if point1 != nil && point2 != nil{
			val1 := point1.Val
			val2 := point2.Val
			var node *ListNode
			// 当token为true时，代表上一次两指针指向的结点的值相加大于10，需要在这一次相加操作进行额外加1操作
			if token == true{
				val := val1 + val2 + 1
				// 当额外加1操作仍然大于10时，token仍为true不变，将和数减10的值构造成链表结点
				if val >= 10{
					node = &ListNode{Val: val - 10}
				// 当额外加1操作小于10时，直接将和数构造成链表节点，再将token置为false，用于告知下一次结点值相加不需要加1
				}else{
					node = &ListNode{Val: val}
					token = false
				}
			// 当token为false，则两指针指向的结点的结点值相加不需要加1，然后将结点值之和构造成链表结点
			}else{
				val := val1 + val2
				// 当和数大于10，需要将和数减10再构造成链表结点，再将token置为true，用于告知下一次结点值相加需要加1
				if val >= 10{
					node = &ListNode{Val: val - 10}
					token = true
				// 当和数小于10，直接将和数构造成链表结点，token不变
				}else{
					node = &ListNode{Val: val}
				}
			}
			// 当新链表头结点指针为空，代表此次相加为第一次，需要将新链表头结点指针以及尾结点指针都初始化为以上构造的结点
			if newHead == nil{
				newHead = node
				point = newHead
			// 如果新链表头结点指针不为空，则将以上构造的结点连接至新链表尾结点，并将新链表尾结点指针往前走一步，指向新的尾结点
			}else{
				point.Next = node
				point = point.Next
			}
			// 两链表指针继续往前走
			point1 = point1.Next
			point2 = point2.Next
		// 当point1不为空而point2不为空时
		}else if point1 != nil{
			// 如果token为true，需要将point1指向的结点的结点值加1在进行结点构造
			if token == true{
				val1 := point1.Val
				val := val1 + 1
				var node *ListNode
				// 如果point1指向的结点的结点值加1后等于10，将0值构造为链表结点，token值不变，用于告知下一个结点值加1操作
				if val == 10{
					node = &ListNode{Val: 0}
				// 如果加1后不等于10，将加1后的和数直接构造成链表结点，并将token值为false，用于告知下一个结点无需加1操作
				}else{
					node  = &ListNode{Val: val}
					token = false
				}
				// 进行新链表尾结点连接以上构造的链表结点
				point.Next = node
				point = point.Next
				// point1指针往前走一步
				point1 = point1.Next
			// 如果token不为true，则将point1指向的结点直接连接至新链表尾结点并break出循环
			}else{
				// 如果新链表头结点为空，则需要将新链表头结点初始化为point1
				if newHead == nil{
					newHead = point1
				// 否则将point1直接连接至新链表尾结点后
				}else{
					point.Next = point1
				}
				break
			}
		// 同上
		}else if point2 != nil{
			if token == true{
				val2 := point2.Val
				val := val2 + 1
				var node *ListNode
				if val == 10{
					node = &ListNode{Val: 0}
				}else{
					node  = &ListNode{Val: val}
					token = false
				}
				point.Next = node
				point = point.Next
				point2 = point2.Next
			}else{
				if newHead == nil{
					newHead = point2
				}else{
					point.Next = point2
				}
				break
			}
		// 当point1与point2为空但是token仍为true时，则还需要构造一个结点值为1的链表结点，连接至链表尾结点后，再break出循环
		}else if token == true{
			node := &ListNode{Val: 1}
			point.Next = node
			break
		}
	}
	return newHead
}

func getValue(head *ListNode) []int{
	point := head
	result := []int{}
	for ; point != nil; {
		result = append(result, point.Val)
		point = point.Next
	}
	return result
}

func main() {
	// 测试用例
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 9}
	l2.Next = &ListNode{Val: 9}
	newHead := addTwoNumbers(l1, l2)
	fmt.Println(getValue(newHead))
}