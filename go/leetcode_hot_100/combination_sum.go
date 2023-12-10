package leetcode_hot_100

// TODO again
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0, len(candidates))

	var search func(int, []int, int)
	search = func(start int, path []int, curSum int) {
		if curSum > target {
			return
		}

		if curSum == target {
			res = append(res, path)
			return
		}

		for i := start; i < len(candidates); i++ {
			nextPath := make([]int, len(path)+1)
			copy(nextPath, path)
			nextPath[len(path)] = candidates[i]
			search(i, nextPath, curSum+candidates[i])
		}
	}

	search(0, []int{}, 0)
	return res
}
