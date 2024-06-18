package leetcode_hot_100

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	leftFind := func(start int) int {
		for start >= 0 {
			if nums[start] != target {
				break
			}
			start--
		}
		if start >= 0 {
			return start + 1
		}
		return 0
	}

	rightFind := func(start int) int {
		for start < len(nums) {
			if nums[start] != target {
				break
			}
			start++
		}
		if start <= len(nums)-1 {
			return start - 1
		}
		return len(nums) - 1
	}

	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			break
		}

		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}

	if nums[mid] != target {
		return []int{-1, -1}
	}

	return []int{leftFind(mid), rightFind(mid)}
}
