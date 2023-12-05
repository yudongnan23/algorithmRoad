package leetcode_hot_100

const (
	zero byte = 48
	one  byte = 49
)

// TODO again
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == zero {
				continue
			}

			count++

			queue := make([]int, 1)
			queue[0] = makeKey(i, j)
			grid[i][j] = zero

			for len(queue) != 0 {
				cur := queue[0]
				queue = queue[1:]

				m, n := parseKey(cur)

				// 上
				if m > 0 && grid[m-1][n] == one {
					grid[m-1][n] = zero
					queue = append(queue, makeKey(m-1, n))
				}

				// 下
				if m < len(grid)-1 && grid[m+1][n] == one {
					grid[m+1][n] = zero
					queue = append(queue, makeKey(m+1, n))
				}

				// 左
				if n > 0 && grid[m][n-1] == one {
					grid[m][n-1] = zero
					queue = append(queue, makeKey(m, n-1))
				}

				// 右
				if n < len(grid[i])-1 && grid[m][n+1] == one {
					grid[m][n+1] = zero
					queue = append(queue, makeKeyII(m, n+1))
				}
			}
		}
	}

	return count
}

func makeKeyII(i, j int) int {
	return i<<16 | j
}

func parseKey(n int) (int, int) {
	i := n >> 16
	j := ((1 << 16) - 1) & n
	return i, j
}
