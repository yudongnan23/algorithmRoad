package sort

func selectSort(nums []int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		minIndex := left
		maxIndex := right
		for i := left; i <= right; i++ {
			if nums[i] < nums[minIndex] && i < right {
				minIndex = i
			}

			if nums[i] > nums[maxIndex] {
				maxIndex = i
			}
		}

		leftVal := nums[left]
		rightVal := nums[right]
		nums[left] = nums[minIndex]
		nums[right] = nums[maxIndex]
		if minIndex != right {
			nums[minIndex] = leftVal
		}
		if maxIndex != left {
			nums[maxIndex] = rightVal
		}

		left++
		right--
	}
}
