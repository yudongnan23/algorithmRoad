/*
	给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
 */

/*
	思路： 深度优先搜索"1"
 */

package main

func numIslands(grid [][]byte) int {
	numsIslands := 0
	if len(grid) == 0 || len(grid[0]) == 0{
		return numsIslands
	}
	acrossLength := len(grid)
	verticalLength := len(grid[0])

	for across := 0; across < acrossLength; across ++{
		for vertical := 0; vertical < verticalLength; vertical ++{
			if grid[across][vertical] == 49{
				numsIslands ++
				grid[across][vertical] = 48
				dfs(across, vertical, acrossLength, verticalLength, grid)
			}
		}
	}
	return numsIslands
}

func dfs(across int, vertical int, acrossLength int, verticalLength int, grid [][]byte){
	stack := [][]int{[]int{across, vertical}}

	for len(stack) != 0{
		maxIndex := len(stack) - 1
		acrossIndex := stack[maxIndex][0]
		verticalIndex := stack[maxIndex][1]
		stack = stack[0:maxIndex]

		if acrossIndex > 0 && grid[acrossIndex - 1][verticalIndex] == 49{
			grid[acrossIndex - 1][verticalIndex] = 48
			stack = append(stack, []int{acrossIndex - 1, verticalIndex})
		}

		if acrossIndex < acrossLength - 1 && grid[acrossIndex + 1][verticalIndex] == 49{
			grid[acrossIndex + 1][verticalIndex] = 48
			stack = append(stack, []int{acrossIndex + 1, verticalIndex})
		}

		if verticalIndex > 0 && grid[acrossIndex][verticalIndex - 1] == 49{
			grid[acrossIndex][verticalIndex - 1] = 48
			stack = append(stack, []int{acrossIndex, verticalIndex - 1})
		}

		if verticalIndex < verticalLength - 1 && grid[acrossIndex][verticalIndex + 1] == 49{
			grid[acrossIndex][verticalIndex + 1] = 48
			stack = append(stack, []int{acrossIndex, verticalIndex + 1})
		}
	}
}
