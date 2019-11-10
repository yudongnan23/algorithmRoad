/*
	给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
 */

/*
	解题思路：
		- 对分四种情况讨论，分别是在四条边上的转向，并定义一个转向令牌，根据令牌的值进行相应的遍历
 */

package  main

import (
	"fmt"
)

func main() {
	nums :=  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(findDiagonalOrder(nums))

}

func findDiagonalOrder(matrix [][]int) []int{
	result := []int{}
	if len(matrix) == 0{
		return result
	}
	i := 0
	j := 0
	token := "up"
	for ; i < len(matrix) && j < len(matrix[0]);{
		if j == len(matrix[0]) - 1 && token == "up" {			// 当元素遍历来到最右上角时需要进行的转向操作
			result = append(result, matrix[i][j])
			i = i + 1
			token = "down"
		}else if i == 0 && token == "up"{					// 当元素遍历需要在上边进行转向操作时
			result = append(result, matrix[i][j])
			j = j + 1
			token = "down"
		}else if i == len(matrix) - 1 && token == "down"{   // 最左下角
			result = append(result, matrix[i][j])
			j = j + 1
			token = "up"
		}else if j == 0 && token == "down"{    // 左边
			result = append(result, matrix[i][j])
			i = i + 1
			token = "up"
		}else {
			if token == "down"{
				result = append(result, matrix[i][j])
				i = i +1
				j = j -1
			}else if token == "up"{
				result = append(result,matrix[i][j])
				i = i - 1
				j = j + 1
			}
		}
	}
	return result
}