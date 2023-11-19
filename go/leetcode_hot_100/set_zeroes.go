package leetcode_hot_100

func setZeroes(matrix [][]int) {
	zeroMap := make(map[int]bool, len(matrix)*len(matrix[0]))
	row := make(map[int]bool, len(matrix))
	column := make(map[int]bool, len(matrix[0]))

	changeFunc := func(i, j int) {
		if !row[i] {
			for m := 0; m < len(matrix[i]); m++ {
				if matrix[i][m] == 0 {
					continue
				}
				zeroMap[makeKey(i, m)] = true
				matrix[i][m] = 0
			}
			row[i] = true
		}

		if !column[j] {
			for n := 0; n < len(matrix); n++ {
				if matrix[n][j] == 0 {
					continue
				}
				zeroMap[makeKey(n, j)] = true
				matrix[n][j] = 0
			}
			column[j] = true
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if !zeroMap[makeKey(i, j)] && matrix[i][j] == 0 {
				changeFunc(i, j)
			}
		}
	}
}

func makeKey(i, j int) int {
	return i<<32 | j
}
