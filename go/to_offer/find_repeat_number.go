package to_offer

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] != i {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		} else {
			i++
		}
	}
	return -1
}
