package leetcode_hot_100

// TODO three - 旋转图像
func rotateMatrix(matrix [][]int) {
	length := len(matrix)
	// 竖中轴线交换
	for i := 0; i < length; i++ {
		l := 0
		r := length - 1
		for l < r {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}

	// 右对角线交换
	for i := 0; i < length; i++ {
		for j := 0; j+i < length-1; j++ {
			matrix[i][j], matrix[length-j-1][length-i-1] = matrix[length-j-1][length-i-1], matrix[i][j]
		}
	}
}
