package leetcode_hot_100

// TODO again
func minWindow(s string, t string) string {
	ht := make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		ht[t[i]]++
	}

	l := 0
	r := 0
	valid := 0
	length := len(s)
	minLen := len(s) + 1
	start := 0

	window := make(map[byte]int, 0)

	for r < length {
		char := s[r]
		r++

		if ht[char] > 0 {
			window[char]++
			if window[char] == ht[char] {
				valid++
			}
		}

		for valid == len(ht) {
			if r-l < minLen {
				start = l
				minLen = r - l
			}

			removeChar := s[l]
			l++
			if ht[removeChar] > 0 {
				if window[removeChar] == ht[removeChar] {
					valid--
				}
				window[removeChar]--
			}
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}
