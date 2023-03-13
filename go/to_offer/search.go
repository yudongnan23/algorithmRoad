package to_offer

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	var mid int
	left := 0
	right := len(nums) - 1
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
		return 0
	}

	return 1 + findLeft(nums, target, mid-1) + findRight(nums, target, mid+1)
}

func findLeft(nums []int, target int, startIndex int) int {
	var count int
	for i := startIndex; i >= 0; i-- {
		if nums[i] == target {
			count++
		}
	}
	return count
}

func findRight(nums []int, target int, startIndex int) int {
	var count int
	for i := startIndex; i < len(nums); i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}
