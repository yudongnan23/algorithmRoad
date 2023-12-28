package leetcode_hot_100

// TODO again
func maxProduct(nums []int) int {
	size := len(nums)
	maxRes, curMax, curMin := nums[0], nums[0], nums[0]
	for i := 1; i < size; i++ {
		if nums[i] < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(curMax*nums[i], nums[i])
		curMin = min(curMin*nums[i], nums[i])
		maxRes = max(curMax, maxRes)
	}

	return maxRes
}
