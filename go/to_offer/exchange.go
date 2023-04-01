package to_offer

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	l := 0
	for r := 1; r < len(nums); r++ {
		if nums[r]%2 == 1 {
			nums[l], nums[r] = nums[r], nums[l]
		}

		if nums[l]%2 == 1 {
			l = l + 1
		}
	}

	return nums
}
