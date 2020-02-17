/*
	一个机器人位于一个 m x n 网格的左上角）。机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角。问总共有多少条不同的路径？
 */

package main

import "fmt"

func totalPath(m int, n int)int{
	f := getArray(m, n)

	for i := 0; i < m; i++{
		for j := 0; j < n; j++{
			if i == 0 || j == 0{
				f[i][j] = 1
			}else{
				f[i][j] = f[i - 1][j] + f[i][j - 1]
			}
		}
	}
	return f[m - 1][n - 1]
}

func getArray(m int, n int)[][]int{
	f := make([][]int, m)
	ff := make([]int, n)
	for i := 0; i < m; i++{
		f[i] = ff
	}
	return f
}

func main() {
	m := 3
	n := 3
	fmt.Println(totalPath(m, n))
}