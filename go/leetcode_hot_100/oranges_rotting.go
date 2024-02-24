package leetcode_hot_100

// TODO three - 腐烂的橘子
func orangesRotting(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	badQueue := make([]int, 0)
	goodCount := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 2 {
				badQueue = append(badQueue, makeKeyIII(i, j))
			}
			if grid[i][j] == 1 {
				goodCount++
			}
		}
	}

	corrupt := func(i, j int) {
		// 上
		if i > 0 && grid[i-1][j] == 1 {
			goodCount--
			grid[i-1][j] = 2
			badQueue = append(badQueue, makeKeyIII(i-1, j))
		}

		// 下
		if i < n-1 && grid[i+1][j] == 1 {
			goodCount--
			grid[i+1][j] = 2
			badQueue = append(badQueue, makeKeyIII(i+1, j))
		}

		// 左
		if j > 0 && grid[i][j-1] == 1 {
			goodCount--
			grid[i][j-1] = 2
			badQueue = append(badQueue, makeKeyIII(i, j-1))
		}

		// 右
		if j < m-1 && grid[i][j+1] == 1 {
			goodCount--
			grid[i][j+1] = 2
			badQueue = append(badQueue, makeKeyIII(i, j+1))
		}
	}

	if goodCount == 0 {
		return 0
	}

	minute := -1
	for len(badQueue) != 0 {
		minute++
		curBadQueue := make([]int, len(badQueue))
		copy(curBadQueue, badQueue)
		badQueue = badQueue[0:0]
		for _, bad := range curBadQueue {
			i, j := parseKeyII(bad)
			corrupt(i, j)
		}
	}
	if goodCount > 0 {
		minute = -1
	}
	return minute
}

func makeKeyIII(i, j int) int {
	return i<<16 | j
}

func parseKeyII(n int) (int, int) {
	i := n >> 16
	j := ((1 << 16) - 1) & n
	return i, j
}
