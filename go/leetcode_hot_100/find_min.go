package leetcode_hot_100

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minNum := 1 << 16
	var dfs func(int, int)
	dfs = func(left, right int) {
		if left > right {
			return
		}

		if nums[left] <= nums[right] {
			if nums[left] < minNum {
				minNum = nums[left]
			}
			return
		}

		mid := (left + right) / 2
		dfs(left, mid)
		dfs(mid+1, right)
	}

	dfs(0, len(nums)-1)

	return minNum
}
