package to_offer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	for i := 0; i < len(matrix); i++ {
		if target < matrix[i][0] {
			return false
		}

		if target > matrix[i][len(matrix[i])-1] {
			continue
		}

		if findNum(matrix[i], target) {
			return true
		}
	}

	return false
}

func findNum(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}

		if nums[mid] > target {
			right = mid - 1
			continue
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}

	return false
}
