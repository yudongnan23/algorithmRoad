package leetcode_hot_100

// TODO again
func searchII(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	var dfs func(int, int) int
	dfs = func(left, right int) int {
		if left > right {
			return -1
		}

		if right-left == 0 && nums[left] != target {
			return -1
		}

		if right-left == 0 && nums[left] == target {
			return left
		}

		if nums[left] <= nums[right] {
			index := twoWay(nums[left:right+1], target)
			if index != -1 {
				return index + left
			}
		}

		mid := (left + right) / 2
		leftIndex := dfs(left, mid)
		if leftIndex != -1 {
			return leftIndex
		}

		return dfs(mid+1, right)
	}

	return dfs(0, len(nums)-1)
}

func twoWay(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			left = mid + 1
		}

		if nums[mid] < target {
			right = mid - 1
		}
	}

	return -1
}
