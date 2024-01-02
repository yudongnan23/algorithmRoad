package leetcode_hot_100

// TODO again
func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] ^= nums[i-1]
	}
	return nums[len(nums)-1]
}
