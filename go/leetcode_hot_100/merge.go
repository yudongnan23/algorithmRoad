package leetcode_hot_100

import "sort"

// TODO again
func merge(intervals [][]int) [][]int {
	res := make([][]int, len(intervals))
	if len(intervals) == 0 {
		return res
	}

	// 先排个序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})

	index := 0
	res[index] = intervals[0]
	index++

	for i := 1; i < len(intervals); i++ {
		lastInterval := res[index-1]
		if lastInterval[1] < intervals[i][0] {
			res[index] = intervals[i]
			index++
			continue
		}
		res[index-1] = []int{min(lastInterval[0], intervals[i][0]), max(lastInterval[1], intervals[i][1])}
	}

	return res[:index]
}
