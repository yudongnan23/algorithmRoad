package leetcode_hot_100

// TODO again
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		nextNums := make([]int, len(nums)-1)
		copy(nextNums[:i], nums[:i])
		copy(nextNums[i:], nums[i+1:])

		curNums := permute(nextNums)
		for j := 0; j < len(curNums); j++ {
			res = append(res, append([]int{nums[i]}, curNums[j]...))
		}
	}

	return res
}
