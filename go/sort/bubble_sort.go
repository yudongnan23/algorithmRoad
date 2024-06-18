package sort

func bubbleSort(nums []int) {
	right := len(nums) - 1

	for right >= 0 {
		left := 0
		exchange := false
		for left < right {
			if nums[left] > nums[left+1] {
				exchange = true
				nums[left], nums[left+1] = nums[left+1], nums[left]
			}
			left++
		}
		if !exchange {
			break
		}
		right--
	}
}
