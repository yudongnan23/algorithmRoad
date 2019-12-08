/*
	设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，
	next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的
 */


package main

import (
	"fmt"
)

type MyLinkedList struct {
	val *int
	next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	node := MyLinkedList{}
	return node
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this.val == nil{
		return -1
	}else{
		var result = -1
		point := this
		for count := 0; point != nil; count++{
			if count == index{
				result = *(point.val)
				break
			}
			point = point.next
		}
		return result
	}
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.val == nil{
		node := Constructor()
		node.val = &val
		*this = node
	}else{
		node := Constructor()
		node.val = &val
		point := *this
		*this = node
		this.next = &point
	}
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	if this.val == nil{
		tail := Constructor()
		tail.val = &val
		*this = tail
	}else{
		point := this
		tail := Constructor()
		tail.val = &val
		for ; point.next != nil;{
			point = point.next
		}
		point.next = &tail

	}

}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0{
		this.AddAtHead(val)
	}else {
		node := Constructor()
		node.val = &val
		length := this.getLength()
		if index == length{
			this.AddAtTail(val)
		}else if index < length{
				point := this
				for count := 0; count < index-1; count++{
					point = point.next
				}
				node.next = point.next
				point.next = &node
		}
	}
}

func (this *MyLinkedList)getLength()int{
	length := 0
	point := this
	for ; point != nil; {
		length++
		point = point.next
	}
	return length
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	length := this.getLength()
	if index >= 0 && index <= length - 1{
		val := -1
		pre := &MyLinkedList{}
		pre.val = &val
		pre.next = this
		point := pre
		count := -1
		for ; count < index - 1;count++{
			point = point.next
		}
		if count == -1{
			if this.next != nil{
				*this = *this.next
			}else{
				this = nil
			}
		}else if count == length - 1{
			point.next = nil
		}else{
			point.next = point.next.next
		}
	}
}


func (this *MyLinkedList)getValue() []int{
	result := []int{}
	point := this
	for ; point != nil;{
		result = append(result, *point.val)
		point = point.next
	}
	return result
}

func main(){
	head := Constructor()
	head.AddAtHead(3)
	head.AddAtHead(2)
	head.AddAtHead(1)
	head.AddAtTail(4)
	head.AddAtTail(5)
	head.AddAtTail(6)
	head.AddAtIndex(6, 8)
	fmt.Println(head.getValue())
	head.DeleteAtIndex(6)
	fmt.Println(head.getValue())
	fmt.Println(head.getLength())
}


