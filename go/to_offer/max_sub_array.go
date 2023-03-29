package to_offer

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		sum = nums[i] + sum
		if nums[i] > sum {
			if nums[i] > max {
				max = nums[i]
			}
			sum = nums[i]
		} else {
			if sum > max {
				max = sum
			}
		}
	}
	return max
}
