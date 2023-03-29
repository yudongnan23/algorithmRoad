package to_offer

func maxValue(gird [][]int) int {
	if len(gird) == 0 || len(gird[0]) == 0 {
		return 0
	}

	for i := 0; i < len(gird); i++ {
		for j := 0; j < len(gird[i]); j++ {
			var leftValue int
			var aboveValue int

			if i > 0 {
				aboveValue = gird[i-1][j]
			}

			if j > 0 {
				leftValue = gird[i][j-1]
			}

			gird[i][j] = max(leftValue+gird[i][j], aboveValue+gird[i][j])
		}
	}

	return gird[len(gird)-1][len(gird[len(gird)-1])-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
