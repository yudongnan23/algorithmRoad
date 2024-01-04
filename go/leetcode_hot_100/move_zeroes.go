package leetcode_hot_100

// TODO again
func moveZeroes(nums []int) {
	left := 0
	right := 0

	for right < len(nums) {
		if nums[left] == 0 && nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
		}

		if nums[left] != 0 {
			left++
		}

		right++
	}
}
