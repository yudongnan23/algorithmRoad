/*
	编写一个程序，找到两个单链表相交的起始节点
 */

package main


func getIntersectionNode(headA, headB *ListNode) *ListNode{
	pointA := headA
	pointB := headB
	countA := 0
	countB := 0
	for ; pointA != nil || pointB != nil; {
		if pointA != nil{
			pointA = pointA.Next
		}else{
			countA ++
		}
		if pointB != nil{
			pointB = pointB.Next
		}else{
			countB ++
		}
	}
	pointA = headA
	pointB = headB
	if countA != 0{
		count := 0
		for ; count < countA; {
			pointB = pointB.Next
			count++
		}
	}else if countB != 0{
		count := 0
		for ; count < countB; {
			pointA = pointA.Next
			count++
		}
	}
	for ; pointA != nil && pointB != nil; {
		if pointB == pointA{
			return pointB
		}
		pointA = pointA.Next
		pointB = pointB.Next
	}
	return nil
}
