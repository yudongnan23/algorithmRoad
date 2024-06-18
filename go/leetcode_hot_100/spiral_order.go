package leetcode_hot_100

const checkedNum = 100

// TODO three - 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	columns := len(matrix[0])
	res := make([]int, rows*columns)
	index := 0
	i := 0
	j := 0

	checked := func(i, j int) bool {
		return matrix[i][j] > checkedNum
	}

	appendFunc := func() {
		if index < rows*columns && !checked(i, j) {
			res[index] = matrix[i][j]
			matrix[i][j] = 101
			index++
		}
	}

	right := func() {
		for {
			appendFunc()
			if j+1 < len(matrix[0]) && !checked(i, j+1) {
				j++
				continue
			}
			break
		}
	}

	down := func() {
		for {
			appendFunc()
			if i+1 < len(matrix) && !checked(i+1, j) {
				i++
				continue
			}
			break
		}
	}

	left := func() {
		for {
			appendFunc()
			if j-1 >= 0 && !checked(i, j-1) {
				j--
				continue
			}
			break
		}
	}

	up := func() {
		for {
			appendFunc()
			if i-1 >= 0 && !checked(i-1, j) {
				i--
				continue
			}
			break
		}
	}

	for index < rows*columns {
		// 右
		right()

		// 下
		down()

		// 左
		left()

		// 上
		up()
	}

	return res
}
