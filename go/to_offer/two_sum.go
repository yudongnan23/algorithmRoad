package to_offer

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	l := 0
	r := len(nums) - 1

	for l < r {
		if nums[l]+nums[r] == target {
			return []int{nums[l], nums[r]}
		}

		if nums[l]+nums[r] > target {
			r = r - 1
		} else {
			l = l + 1
		}
	}

	return nil
}
