package to_offer

func missingNumber(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}

	n := len(nums)
	return ((0+n)*(n+1))/2 - sum
}
