package leetcode_hot_100

import "math"

// TODO three - 子集
func subsets(nums []int) [][]int {
	ans := make([][]int, 1, int(math.Pow(2, float64(len(nums))))+1)
	ans[0] = []int{}
	for _, x := range nums {
		for _, arr := range ans {
			a := make([]int, len(arr), len(arr)+1)
			copy(a, arr)
			ans = append(ans, append(a, x))
		}
	}
	return ans
}
