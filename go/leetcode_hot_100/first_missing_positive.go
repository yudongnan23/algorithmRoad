package leetcode_hot_100

// TODO three - 缺失的第一个正数
func firstMissingPositive(nums []int) int {
	size := len(nums)
	for i := 0; i < size; {
		if nums[i] >= 1 && nums[i] <= size && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	for i := 0; i < size; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return size + 1
}
