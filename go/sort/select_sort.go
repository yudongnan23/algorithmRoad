package sort

func selectSort(nums []int) {
	left := 0
	for left < len(nums) {
		minIndex := left
		for i := left + 1; i < len(nums); i++ {
			if nums[i] < nums[minIndex] {
				minIndex = i
			}
		}
		if minIndex != left {
			nums[left], nums[minIndex] = nums[minIndex], nums[left]
		}

		left++
	}
}
