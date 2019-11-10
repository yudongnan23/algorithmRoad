/*
	给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 */


package main

import "fmt"

func main() {
	//nums :=  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13 ,14, 15, 16}}
	nums := [][]int{{7}, {9}, {6}}
	fmt.Println(spiralOrder(nums))

}

func spiralOrder(matrix [][]int) []int{
	result := []int{}
	if len(matrix) == 0{
		return result
	}
	result = up(matrix, result)
	return result
}
func up(nums [][]int, result []int) []int{
	if len(nums) == 0 || len(nums[0]) == 0{
		return result
	}
	result_item := []int{}
	result_item = append(result_item, nums[0]...)
	new_nums := nums[1:]
	result = append(result, result_item...)
	return right(new_nums, result)

}
func right(nums [][]int, result []int) []int{
	if len(nums) == 0 || len(nums[0]) == 0{
		return result
	}
	result_item := []int{}
	new_nums := [][]int{}
	for _, num := range nums{
		result_item = append(result_item, num[len(num)-1])
		new_nums = append(new_nums, num[0:len(num)-1])
	}
	result = append(result, result_item...)
	return down(new_nums, result)
}
func down(nums [][]int, result []int) []int{
	if len(nums) == 0 || len(nums[0]) == 0{
		return result
	}
	result_item := []int{}
	result_item = nums[len(nums)-1]
	new_nums := nums[0:len(nums)-1]
	for i := len(result_item) - 1; i >= 0; i -- {
		result = append(result, result_item[i])
	}
	return left(new_nums, result)
}
func left(nums [][]int, result []int)[]int{
	if len(nums) == 0 || len(nums[0]) == 0{
		return result
	}
	result_item := []int{}
	new_nums := [][]int{}
	for _, num := range nums{
		result_item = append(result_item, num[0])
		new_nums = append(new_nums, num[1:])
	}
	for i := len(result_item) - 1; i >= 0; i --{
		result = append(result, result_item[i])
	}
	return up(new_nums, result)
}