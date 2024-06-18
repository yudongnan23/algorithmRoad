package leetcode_hot_100

// TODO three - 搜索旋转排序数组
func searchII(nums []int, target int) int {
	return dfsSearchII(nums, 0, len(nums)-1, target)
}

func dfsSearchII(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	if nums[left] <= nums[right] {
		return searchInNums(nums[left:right+1], left, target)
	}
	mid := (left + right) / 2
	leftIndex := dfsSearchII(nums, left, mid, target)
	if leftIndex > -1 {
		return leftIndex
	}
	rightIndex := dfsSearchII(nums, mid+1, right, target)
	if rightIndex > -1 {
		return rightIndex
	}
	return -1
}

func searchInNums(nums []int, addIndex, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid + addIndex
		}

		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}
