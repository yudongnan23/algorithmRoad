package leetcode_hot_100

func searchMatrix(matrix [][]int, target int) bool {
	length := len(matrix)
	for i := 0; i < length; i++ {
		if matrix[i][len(matrix[i])-1] < target {
			continue
		}

		if matrix[i][0] > target {
			return false
		}

		if search(matrix[i], target) {
			return true
		}
	}

	return false
}

func search(nums []int, target int) bool {
	l := 0
	r := len(nums)

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}

		if nums[mid] > target {
			r = mid - 1
			continue
		}

		if nums[mid] < target {
			l = mid + 1
			continue
		}
	}

	return false
}
