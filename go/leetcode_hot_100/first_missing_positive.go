package leetcode_hot_100

import "sort"

// TODO again
func firstMissingPositive(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < nums[j] {
			return true
		}
		return false
	})

	length := len(nums)
	min := 1
	for i := 0; i < length; i++ {
		if nums[i] <= 0 {
			continue
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		if min != nums[i] {
			break
		}
		min++
	}
	return min
}
