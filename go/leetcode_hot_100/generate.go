package leetcode_hot_100

func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)

	res = append(res, []int{1})
	for i := 2; i <= numRows; i++ {
		curNums := make([]int, i)
		curNums[0] = 1
		curNums[i-1] = 1

		for j := 1; j < i-1; j++ {
			curNums[j] = res[i-2][j-1] + res[i-2][j]
		}

		res = append(res, curNums)
	}
	return res
}
