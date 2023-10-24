package leetcode_hot_100

// TODO again
func longestConsecutive(nums []int) int {
	d := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		d[nums[i]] = true
	}

	maxCount := 0

	for k := range d {
		if !d[k-1] {
			curMaxCount := 1
			for d[k+1] {
				k++
				curMaxCount++
			}

			if curMaxCount > maxCount {
				maxCount = curMaxCount
			}
		}
	}

	return maxCount
}
