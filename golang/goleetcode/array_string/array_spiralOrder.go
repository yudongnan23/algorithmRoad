/*
	给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 */

/*
    思路：定义四个函数，up函数负责对二维数组矩阵的上边元素进行遍历，right函数负责对二维数组矩阵的右边元素进行遍历，
         down函数负责对二维矩阵的下边元素进行遍历，left函数负责对二维矩阵的左边元素进行比遍历，每个函数遍历完成后将遍历
         过的的元素在二维矩阵中删除，再将二维矩阵传递给下一个函数进行遍历，四个函数遍历顺序为：
            up - right - down - left - up
*/

package main

import "fmt"

func main() {
    // 测试用例
	nums :=  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13 ,14, 15, 16}}
	fmt.Println(spiralOrder(nums))

}

func spiralOrder(matrix [][]int) []int{
	result := []int{}
	if len(matrix) == 0{
		return result
	}
	// 调用up函数对二维矩阵进行遍历
	result = up(matrix, result)
	return result
}

func up(nums [][]int, result []int) []int{
    // 当传入的二维矩阵为空时，返回当前遍历结果切片result
	if len(nums) == 0 || len(nums[0]) == 0{
		return result
	}
	result_item := []int{}
	// 获得二维矩阵的上边所有元素
	result_item = append(result_item, nums[0]...)
	// 获得需要传入下一个函数的二维矩阵，即删除二维矩阵中所有上边元素
	new_nums := nums[1:]
	// 将二维矩阵上边所有的元素添加进遍历结果切片
	result = append(result, result_item...)
	// 调用下一个函数，传入新的二维矩阵和遍历结果切片
	return right(new_nums, result)
}

// 一下函数实现思路同up函数
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