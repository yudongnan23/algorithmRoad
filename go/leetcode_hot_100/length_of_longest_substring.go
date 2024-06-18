package leetcode_hot_100

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int, 0)
	maxLength := 0
	l := 0
	r := 0
	size := len(s)
	for r < size {
		index, ok := m[s[r]]
		if ok && index >= l {
			if r-l > maxLength {
				maxLength = r - l
			}
			l = index + 1
		}
		m[s[r]] = r
		r++
	}

	if r-l > maxLength {
		maxLength = r - l
	}

	return maxLength
}
