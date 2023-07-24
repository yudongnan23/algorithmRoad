package to_offer

func reversePrint(head *ListNode) []int {
	values := make([]int, 10000)
	p := head
	i := 9999
	for p != nil {
		values[i] = p.Val
		p = p.Next
		i--
	}
	return values[i+1:]
}
