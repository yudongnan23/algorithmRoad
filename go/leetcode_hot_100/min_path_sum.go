package leetcode_hot_100

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := grid[i][j]
			leftAdded := false
			if i > 0 {
				sum = grid[i][j] + grid[i-1][j]
				leftAdded = true
			}
			if j > 0 && (grid[i][j]+grid[i][j-1] < sum || !leftAdded) {
				sum = grid[i][j] + grid[i][j-1]
			}
			grid[i][j] = sum

		}
	}
	return grid[m-1][n-1]
}
