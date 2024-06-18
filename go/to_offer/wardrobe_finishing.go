package to_offer

func wardrobeFinishing(m int, n int, cnt int) int {
	checked := make([][]bool, m)
	for i := 0; i < m; i++ {
		checked[i] = make([]bool, n)
	}
	return backTrackingII(checked, m, n, 0, 0, cnt)
}

func sum(num int) int {
	s := 0

	for num > 0 {
		n := num % 10
		s = s + n
		num = num / 10
	}

	return s
}

func backTrackingII(checked [][]bool, m, n, row, col, cnt int) int {
	if row < 0 || row >= m || col < 0 || col >= n {
		return 0
	}

	if checked[row][col] {
		return 0
	}

	count := 0

	if sum(row)+sum(col) <= cnt {
		count = 1
	} else {
		return 0
	}

	checked[row][col] = true

	return count +
		backTrackingII(checked, m, n, row+1, col, cnt) +
		backTrackingII(checked, m, n, row-1, col, cnt) +
		backTrackingII(checked, m, n, row, col+1, cnt) +
		backTrackingII(checked, m, n, row, col-1, cnt)
}
