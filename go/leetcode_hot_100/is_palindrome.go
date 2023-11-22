package leetcode_hot_100

func isPalindrome(head *ListNode) bool {
	vals := make([]int, 0)
	p := head

	for p != nil {
		vals = append(vals, p.Val)
		p = p.Next
	}

	l := 0
	r := len(vals) - 1

	for l < r {
		if vals[l] != vals[r] {
			return false
		}

		l++
		r--
	}

	return true
}
