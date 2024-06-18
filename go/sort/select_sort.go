package sort

func selectSort(nums []int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		minIndex := left
		maxIndex := right
		for i := left; i <= right; i++ {
			if nums[i] < nums[minIndex] {
				minIndex = i
			}

			if nums[i] > nums[maxIndex] {
				maxIndex = i
			}
		}

		leftVal := nums[left]
		rightVal := nums[right]
		minVal := nums[minIndex]
		maxVal := nums[maxIndex]
		nums[left] = minVal
		nums[right] = maxVal
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
