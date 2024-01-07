package leetcode_hot_100

// TODO three
func maxSlidingWindow(nums []int, k int) []int {
	size := len(nums)
	res := make([]int, 0, size-k+1)
	queue := make([]int, 0, k)

	deleteFunc := func(index int) {
		for i := 0; i < len(queue); i++ {
			if index-queue[i] <= k-1 {
				queue = queue[i:]
				break
			}
		}
	}

	enqueue := func(index int) {
		curNum := nums[index]
		queueSize := len(queue)
		if queueSize > 0 && nums[queue[0]] < curNum {
			queue = []int{index}
			return
		}

		left := 0
		right := queueSize - 1
		for left < right {
			mid := (left + right) / 2
			if nums[queue[mid]] == curNum {
				right = mid
				break
			}

			if nums[queue[mid]] > curNum {
				left = mid + 1

			}
			if nums[queue[mid]] < curNum {
				right = mid - 1
			}
		}

		if queueSize > 0 {
			if nums[queue[right]] > curNum {
				right++
			}
			queue = queue[:right]
		}

		queue = append(queue, index)
		deleteFunc(index)
	}

	for i := 0; i < size; i++ {
		enqueue(i)
		if i-k+1 >= 0 {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}
