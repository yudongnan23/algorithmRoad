package leetcode_hot_100

// TODO three - 颜色分类
func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1

	findLeft := func(start int) int {
		for i := start; i < len(nums); i++ {
			if nums[i] != 0 {
				return i
			}
		}
		return start
	}

	findRight := func(start int) int {
		for i := start; i >= 0; i-- {
			if nums[i] != 2 {
				return i
			}
		}
		return start
	}

	i := 0
	for i <= right {
		if nums[i] == 1 {
			i++
			continue
		}
		if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}

		if nums[i] == 2 {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		}

		right = findRight(right)
		left = findLeft(left)
		i = left
	}
}
