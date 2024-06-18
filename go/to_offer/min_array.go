package to_offer

func minArray(nums []int) int {
	left := 0
	right := len(nums) - 1

	flag := false

	for left < right {
		if nums[left] >= nums[right] {
			flag = true
			left++
			continue
		}

		if nums[left] < nums[right] {
			if flag {
				return nums[left]
			}

			right--
		}
	}
	return nums[left]
}
