package to_offer

func missingNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	return ((len(nums) * (len(nums) + 1)) / 2) - nums[len(nums)-1]
}
