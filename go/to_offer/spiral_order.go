package to_offer

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	startI := 0
	startJ := 0
	maxI := len(matrix) - 1
	maxJ := len(matrix[0]) - 1
	res := make([]int, (maxI+1)*(maxJ+1))
	startIndex := 0

	for {
		if startI > maxI || startJ > maxJ {
			break
		}

		// 右
		for j := startJ; j <= maxJ; j++ {
			res[startIndex] = matrix[startI][j]
			startIndex++
		}

		down := false
		// 下
		for i := startI + 1; i <= maxI; i++ {
			res[startIndex] = matrix[i][maxJ]
			startIndex++
			down = true
		}

		left := false
		// 左
		for j := maxJ - 1; j >= startJ; j-- {
			if down {
				res[startIndex] = matrix[maxI][j]
				startIndex++
				left = true
			}

		}

		// 上
		for i := maxI - 1; i > startI; i-- {
			if left {
				res[startIndex] = matrix[i][startJ]
				startIndex++
			}

		}

		startI++
		startJ++
		maxI--
		maxJ--
	}

	return res
}
