package leetcode_hot_100

func setZeroes(matrix [][]int) {
	locMap := make(map[int]bool, 0)
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				locMap[makeKey(i, j)] = true
			}
		}
	}

	markFunc := func(i, j int) {
		for k := 0; k < n; k++ {
			matrix[i][k] = 0
		}

		for l := 0; l < m; l++ {
			matrix[l][j] = 0
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if locMap[makeKey(i, j)] {
				markFunc(i, j)
			}
		}
	}
}

func makeKey(i, j int) int {
	return (i << 32) | j
}
