/*
	给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

 */

package main

import "fmt"

func main() {
	num := 5
	generate(num)
}

func generate(numRows int) [][]int {
	result := [][]int{}
	if numRows == 0{
		return result
	}
	for i := 0; i < numRows; i++{
		if i == 0{
			result = append(result, []int{1})
		}else if i == 1{
			result = append(result, []int{1, 1})
		}else{
			cur_nums := result[len(result)-1]
			new_nums := []int{1}
			for j := 0; j < len(cur_nums) - 1; j++{
				new_nums = append(new_nums, cur_nums[j]+cur_nums[j+1])
			}
			new_nums = append(new_nums, 1)
			result = append(result, new_nums)
		}
	}
	fmt.Println(result)
	return result
}
