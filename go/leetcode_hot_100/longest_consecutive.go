package leetcode_hot_100

// TODO again
func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}

	maxLength := 0
	for k := range m {
		if m[k-1] {
			continue
		}
		start := k
		for m[start] {
			start++
		}

		curLength := start - k

		if curLength > maxLength {
			maxLength = curLength
		}
	}
	return maxLength
}
