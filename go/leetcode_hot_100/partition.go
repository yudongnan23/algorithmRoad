package leetcode_hot_100

// TODO three
func partition(s string) [][]string {
	res := make([][]string, 0, 1000)
	var search func(int, []string)
	search = func(start int, cur []string) {
		if start == len(s) {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(s); i++ {
			if moslems(s[start : i+1]) {
				cur = append(cur, s[start:i+1])
				search(i+1, cur)
				cur = cur[:len(cur)-1]
			}
		}
	}

	search(0, make([]string, 0, len(s)))
	return res
}

func moslems(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
