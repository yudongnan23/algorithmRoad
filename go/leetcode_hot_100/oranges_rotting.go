package leetcode_hot_100

// TODO again
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	badQueue := make([]int, 0)
	goodQueue := make([]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				goodQueue = append(goodQueue, makeKeyIII(i, j))
			}

			if grid[i][j] == 2 {
				badQueue = append(badQueue, makeKeyIII(i, j))
			}
		}
	}

	if len(goodQueue) == 0 {
		return 0
	}

	minutes := 0
	goBadCount := 0
	for goBadCount != len(goodQueue) && len(badQueue) > 0 {
		newBadQueue := make([]int, 0)
		minutes++
		for i := 0; i < len(badQueue); i++ {
			m, n := parseKeyII(badQueue[i])

			// 上
			if m > 0 && grid[m-1][n] == 1 {
				goBadCount++
				grid[m-1][n] = 2
				newBadQueue = append(newBadQueue, makeKeyIII(m-1, n))
			}

			// 下
			if m < len(grid)-1 && grid[m+1][n] == 1 {
				goBadCount++
				grid[m+1][n] = 2
				newBadQueue = append(newBadQueue, makeKeyIII(m+1, n))
			}

			// 左
			if n > 0 && grid[m][n-1] == 1 {
				goBadCount++
				grid[m][n-1] = 2
				newBadQueue = append(newBadQueue, makeKeyIII(m, n-1))
			}

			// 右
			if n < len(grid[m])-1 && grid[m][n+1] == 1 {
				goBadCount++
				grid[m][n+1] = 2
				newBadQueue = append(newBadQueue, makeKeyIII(m, n+1))
			}
		}

		badQueue = newBadQueue
	}

	if goBadCount != len(goodQueue) {
		return -1
	}

	return minutes
}

func makeKeyIII(i, j int) int {
	return i<<16 | j
}

func parseKeyII(n int) (int, int) {
	i := n >> 16
	j := ((1 << 16) - 1) & n
	return i, j
}
