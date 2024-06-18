package to_offer

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	duplicateMapping := make(map[byte]int16, 0)
	duplicateMapping[s[0]] = 0
	var maxLength int16 = 1
	var lastIndex int16 = -1
	for i := 1; i < len(s); i++ {
		var curLength int16
		index, ok := duplicateMapping[s[i]]
		if ok && index >= lastIndex {
			curLength = int16(i) - lastIndex - 1
			lastIndex = index
		} else {
			curLength = int16(i) - lastIndex
		}

		if curLength > maxLength {
			maxLength = curLength
		}

		duplicateMapping[s[i]] = int16(i)
	}

	return int(maxLength)
}
