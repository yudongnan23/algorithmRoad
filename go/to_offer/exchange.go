package to_offer

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	evenIndex := 0
	rightIndex := 0

	for {
		if rightIndex == len(nums) {
			return nums
		}

		if !isOdd(nums[rightIndex]) {
			rightIndex++
			continue
		}

		change(nums, rightIndex, evenIndex)
		evenIndex++
		rightIndex++
	}
}

func isOdd(n int) bool {
	return n%2 == 1
}

func change(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
