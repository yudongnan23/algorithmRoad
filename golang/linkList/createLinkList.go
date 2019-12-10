/*
	设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，
	next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的
 */

/*
	    思路：自定义一个函数用于获取当前链表的长度
        Get：获得链表的长度，判断索引是否正确，正确则定义一个指针，让指针走index步，即可得到index位置上结点的值
        AddAtHead：根据传入的结点值初始化一个结点，判断头结点是否为空，若为空，将头结点指针指向初始化的结点即可；否则将该结点的next
                   指定为头结点，再将头结点指针指向初始化的结点。
        AddAtTail：根据初入的结点初始化一个结点，判断头结点是否是空，若为空，将头结点指针指向初始化的结点即可；否则，定义一个动态指针
                   让指针走到链表的表尾，然后将指针指向的结点的next指向初始化的结点。
        AddAtIndex：首先根据传入的结点值初始化一个结点，然后获取当链表的长度，判断传入的index是否小于或等于0，若是，执行AddAtHead操作；
                    若不是，判断index是否小于链表的长度，若是，则定义一个动态指针，初始化为链表的头部，让该指针走index-1步，将初始的结点
                    的next指向当前指针指向的结点的next结点，再将当前指针指向的结点的next指向初始化的结点；如不是，判断index是否等于
                    链表的长度，若是，则进行AddAtTail操作；若不是，则代表index大于链表的长度，则不进行任何操作
 */

package main

import (
	"fmt"
)

type MyLinkedList struct {
	// 结点的值为一个指向int型值的指针
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
	// 判断结点值是否为nil指针，若是nil指针代表链表为空，返回-1
	if this.val == nil{
		return -1
	}else{
		// 定义动态指针，初始化为头结点
		var result = -1
		point := this
		// 让指针走index步
		for count := 0; point != nil; count++{
			if count == index {
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
	// 判断链表是否为空
	if this.val == nil{
		// 空链表，初始化结点，将头结点指针指向该初始化的结点
		node := Constructor()
		node.val = &val
		*this = node
	}else{
		// 链表非空，初始化新的结点，将新的节点的next指向头结点，再将头结点指针指向初始化的结点
		node := Constructor()
		node.val = &val
		point := *this
		*this = node
		this.next = &point
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	// 判断头结点是否为空
	if this.val == nil{
		// 空链表，初始化结点，将头结点指针指向该初始化的结点
		tail := Constructor()
		tail.val = &val
		*this = tail
	}else{
		// 不为空链表，定义一个指向表头的指针，让该指针走到表尾
		point := this
		tail := Constructor()
		tail.val = &val
		for ; point.next != nil;{
			point = point.next
		}
		// 将表尾结点的next指向初始化的新结点
		point.next = &tail

	}

}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 判断index是否小于等于0，若是，执行AddAtHead操作
	if index <= 0{
		this.AddAtHead(val)
	}else {
		// 初始化一个新的链表
		node := Constructor()
		node.val = &val
		// 获取当前链表的长度
		length := this.getLength()
		// 若index等于链表的长度，执行AddAtTail操作
		if index == length{
			this.AddAtTail(val)
		}else if index < length{
			// index小于链表长度，执行插入操作
				point := this
				// 让指针走index-1步
				for count := 0; count < index-1; count++{
					point = point.next
				}
				// 执行插入第index结点前的操作
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
	// index必须在0-length之间才有效，否则不进行任何操作
	if index >= 0 && index <= length - 1{
		// 初始化一个新的结点
		val := -1
		pre := &MyLinkedList{}
		pre.val = &val
		// 将新结点next指向链表头结点
		pre.next = this
		// 动态指针指向新结点
		point := pre
		count := -1
		// 动态指针走index步
		for ; count < index - 1;count++{
			point = point.next
		}
		// 当count为-1，代表index为0时，删除的头结点，需要重新定义头结点
		if count == -1{
			if this.next != nil{
				*this = *this.next
			}else{
				this = nil
			}
			// 当count为length-1，代表index等于length，则此时动态指针指向的结点为表尾结点的前一结点，将动态指针指向的next指向nil指针即可
		}else if count == length - 1{
			point.next = nil
		}else{
			// 执行正常的删除操作
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
