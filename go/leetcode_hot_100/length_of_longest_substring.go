package leetcode_hot_100

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var left int16 = 0
	var right int16 = 1

	strIndexMapping := map[byte]int16{
		s[0]: 0,
	}

	length := int16(len(s))
	var max int16
	for right < length {
		index, ok := strIndexMapping[s[right]]
		if ok && index >= left {
			leftLen := index - left + 1
			rightLen := right - index
			curMax := maxThree(leftLen, rightLen, right-left)
			if curMax > max {
				max = curMax
			}

			left = index + 1
		}

		strIndexMapping[s[right]] = right
		right++
	}

	if right-left > max {
		return int(right - left)
	}

	return int(max)
}

func maxThree(i, j, k int16) int16 {
	if i >= j && i >= k {
		return i
	}

	if j >= i && j >= k {
		return j
	}

	return k
}
