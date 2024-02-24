package leetcode_hot_100

// TODO three - 寻找重复数
func findDuplicate(nums []int) int {
	i := 0
	for i < len(nums) {
		if i == nums[i] {
			i++
			continue
		}

		if nums[nums[i]] == nums[i] {
			return nums[i]
		}

		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}
	return 0
}
