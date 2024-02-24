package leetcode_hot_100

// TODO three - 只出现一次的数字
func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] ^ nums[i-1]
	}
	return nums[len(nums)-1]
}
