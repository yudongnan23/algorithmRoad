package leetcode_hot_100

// TODO again 二分查找优化
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	queue := make([]int, 0)
	index := 0
	for i := 0; i < len(nums); i++ {
		queue = insert(nums, i, queue, k)
		if i >= k-1 {
			res[index] = nums[queue[0]]
			index++
		}
	}

	return res
}

func insert(nums []int, index int, queue []int, k int) []int {
	deleteFunc := func() {
		for i := 0; i < len(queue); i++ {
			if index-queue[i] <= k-1 {
				queue = queue[i:]
				break
			}
		}
	}

	if len(queue) > 0 && nums[index] > nums[queue[0]] {
		queue = []int{index}
		return queue
	}

	if len(queue) > 0 && nums[index] < nums[queue[len(queue)-1]] {
		queue = append(queue, index)
		deleteFunc()
		return queue
	}

	for i := len(queue) - 1; i >= 0; i-- {
		if nums[queue[i]] >= nums[index] {
			queue = queue[:i+1]
			break
		}
	}

	queue = append(queue, index)
	deleteFunc()

	return queue
}
