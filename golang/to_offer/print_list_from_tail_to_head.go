package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param head ListNode类
 * @return int整型一维数组
 */
func printListFromTailToHead(head *ListNode) []int {
	result := make([]int, 0)

	if head == nil {
		return result
	}

	cur := head
	for cur != nil {
		result = append(result, cur.Val)

		cur = cur.Next
	}


	return result[:]
}

func main(){
	fmt.Println(printListFromTailToHead(&ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}))
}
