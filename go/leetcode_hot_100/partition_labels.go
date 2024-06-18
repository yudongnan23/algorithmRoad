package leetcode_hot_100

// TODO three - 划分字母区间
func partitionLabels(s string) []int {
	res := make([]int, 0)
	d := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		d[s[i]] = i
	}

	left := 0
	right := 0
	curMaxIndex := d[s[0]]
	for right < len(s) {
		if right > curMaxIndex {
			res = append(res, right-left)
			left = right
			curMaxIndex = d[s[right]]
		} else {
			if d[s[right]] > curMaxIndex {
				curMaxIndex = d[s[right]]
			}
		}
		right++
	}

	res = append(res, right-left)

	return res
}
