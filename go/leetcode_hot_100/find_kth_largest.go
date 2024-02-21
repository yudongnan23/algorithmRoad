package leetcode_hot_100

// TODO three
func findKthLargest(nums []int, k int) int {
	length := len(nums)
	for i := length / 2; i >= 0; i-- {
		heap(nums, length, i)
	}

	sortCount := 0
	for i := length - 1; i > 0 && sortCount < k; i-- {
		swap(nums, 0, i)
		length--
		heap(nums, length, 0)
		sortCount++
	}
	return nums[len(nums)-k]
}

func heap(nums []int, length, index int) {
	left := 2*index + 1
	right := 2*index + 2

	largest := index

	if left < length && nums[left] > nums[largest] {
		largest = left
	}

	if right < length && nums[right] > nums[largest] {
		largest = right
	}

	if largest != index {
		swap(nums, largest, index)
		heap(nums, length, largest)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
