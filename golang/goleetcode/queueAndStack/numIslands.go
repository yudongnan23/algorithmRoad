/*
	给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
 */

/*
     思路： 本题采用广度优先的思路解题，遍历二维数组，当遇到49时，将该元素改48，再将其两个索引放入队列，然后以队列不为空为循环条件，弹出队列中的队首元素索引，对其周围
            元素进行判断，当遇到49，改48，再次索引入队，直到队列为空，代表一个岛屿结束。直到下一个岛屿的出现执行同样的操作
            注：go语言中byte类型49对应字符串"1", 48对应字符串"0"
 */

package main

import "fmt"

func numIslands(grid [][]byte) int {
     // 当二维数组为空时，及时返回0
    if len(grid) == 0 || len(grid) == 1 && len(grid[0]) == 0{
        return 0
    }
	acrossLength := len(grid[0])
	verticalLength := len(grid)
	island := 0
	// 对二维数组进行遍历
	for i := 0; i < verticalLength; i ++{
		for j := 0; j < acrossLength; j ++{
		    // 当遇到时，将该元素改48后再对其周围元素尽心处理
			if grid[i][j] == 49{
				island ++
				grid[i][j] = 0
				otherIsland(grid, i , j, acrossLength, verticalLength)
			}else{
				continue
			}
		}
	}
	return island
}

func otherIsland(grid [][]byte, i int, j int, acrossLength int, verticalLength int){
    // 队列初始化入参的两个索引
	var queue [][]int
	queue = append(queue, []int{i, j})
	// 当队列不为空时，执行该循环
	for ; len(queue) != 0; {
	    // 获取队列中的两个索引
		m := queue[0][0]
		n := queue[0][1]
		queue = queue[1:]
		// 判断当前索引指向元素在矩阵中的上面的元素
		if m > 0 && grid[m-1][n] == 49 {
			grid[m-1][n] = 0
			queue = append(queue, []int{m - 1, n})
		}
		// 判断当前索引指向元素在矩阵中的下面元素
		if m < verticalLength - 1 && grid[m+1][n] == 49 {
			grid[m+1][n] = 0
			queue = append(queue, []int{m + 1, n})
		}
		// 判断当前元素指向元素在矩阵中的左边元素
		if n > 0 && grid[m][n-1] == 49 {
			grid[m][n-1] = 0
			queue = append(queue, []int{m, n - 1})
		}
		// 判断当前索引指向元素在矩阵中的右边元素
		if n < acrossLength - 1 && grid[m][n+1] == 49 {
			grid[m][n+1] = 0
			queue = append(queue, []int{m, n + 1})
		}
	}
}

func main() {
    // 测试用例
	list := [][]byte{{49, 49, 48, 48, 48},{49, 49, 48, 48, 48},{48, 48, 49, 48, 48},{48, 48, 48, 49, 49}}
	fmt.Println(numIslands(list))

}
