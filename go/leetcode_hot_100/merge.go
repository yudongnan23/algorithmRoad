package leetcode_hot_100

import "sort"

// TODO three - 合并区间
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})

	size := len(intervals)
	res := make([][]int, 0)
	res = append(res, intervals[0])

	for i := 1; i < size; i++ {
		curMax := res[len(res)-1][1]
		if intervals[i][0] > curMax {
			res = append(res, intervals[i])
			continue
		}

		if intervals[i][1] > curMax {
			res[len(res)-1][1] = intervals[i][1]
		}
	}

	return res
}
