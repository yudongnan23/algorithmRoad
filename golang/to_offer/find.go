package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param target int整型
 * @param array int整型二维数组
 * @return bool布尔型
 */
func Find( target int ,  array [][]int ) bool {
	var result bool

	if len(array) == 0 || len(array[0]) == 0{
		return result
	}

	queen := [][]int{{0, 0}}
	dict := make(map[string]bool)

	for len(queen) != 0 {
		cur := queen[len(queen) - 1]
		queen = queen[0:len(queen) - 1]

		i := cur[0]
		j := cur[1]
		if array[i][j] == target {
			result = true
			break
		}

		if exist, _ := dict[fmt.Sprintf("%d:%d", i + 1, j)]; !exist &&  i + 1 < len(array)  && array[i + 1][j] <= target {
			if array[i + 1][j] == target {
				result = true
				break
			}
			dict[fmt.Sprintf("%d:%d", i + 1, j)] = true
			queen  = append(queen, []int{i + 1, j})
		}
		if exist, _ := dict[fmt.Sprintf("%d:%d", i, j + 1)]; !exist && j + 1 < len(array[i])  && array[i][j + 1] <= target {
			if array[i][j + 1] == target {
				result = true
				break
			}
			dict[fmt.Sprintf("%d:%d", i, j + 1)] = true
			queen = append(queen, []int{i, j + 1})
		}
	}

	return result
}

func main(){
	Find(11, [][]int{{1, 2, 3, 4},
							{5, 6, 7, 8},
							{9, 10, 11, 12},
							{13, 14, 15, 16}})
}