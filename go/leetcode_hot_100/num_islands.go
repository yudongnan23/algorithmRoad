package leetcode_hot_100

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	n := len(grid)
	m := len(grid[0])
	count := 0
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	var visit func(i, j int)

	visit = func(i, j int) {
		visited[i][j] = true
		// 上
		if i > 0 && grid[i-1][j] == '1' && !visited[i-1][j] {
			visit(i-1, j)
		}

		// 下
		if i < n-1 && grid[i+1][j] == '1' && !visited[i+1][j] {
			visit(i+1, j)
		}

		// 左
		if j > 0 && grid[i][j-1] == '1' && !visited[i][j-1] {
			visit(i, j-1)
		}

		// 右
		if j < m-1 && grid[i][j+1] == '1' && !visited[i][j+1] {
			visit(i, j+1)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				continue
			}

			if !visited[i][j] {
				visit(i, j)
				count++
			}
		}
	}

	return count
}
