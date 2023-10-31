package leetcode_hot_100

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 1
	m := make(map[byte]int, 0)

	l := 0
	r := 1

	m[s[l]] = 0

	for r < len(s) {
		index, ok := m[s[r]]
		if !ok || index < l {
			if r-l+1 > max {
				max = r - l + 1
			}
		} else {
			l = index + 1
		}

		m[s[r]] = r
		r++
	}

	return max
}
