/*
	给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

 */

/*
    思路：按照杨辉三角的构造原理逐步得到每一个数组。
 */

package main

import "fmt"

func main() {
    // 测试用例
	num := 5
	generate(num)
}

func generate(numRows int) [][]int {
    // 定义一个结果列表，用于接收数组
	result := [][]int{}
	// 当numRows为空时，没有必要进行任何操作，直接返回空列表
	if numRows == 0{
		return result
	}
	// 对结果列表的每一个数组进行有一一构造
	for i := 0; i < numRows; i++{
	    // 当i为0时，直接将[1]添加进结果结果列表
		if i == 0{
			result = append(result, []int{1})
		// 当i为1时，将[1, 1]添加进结果列表
		}else if i == 1{
			result = append(result, []int{1, 1})
		}else{
		    // 获取当前结果列表中的最新数组
			cur_nums := result[len(result)-1]
			// 初始化新列表的第一个元素，为1
			new_nums := []int{1}
			// 新列表的每个元素由当前结果列表中最新数组对应索引与对应索引加1的两个元素之和
			for j := 0; j < len(cur_nums) - 1; j++{
				new_nums = append(new_nums, cur_nums[j]+cur_nums[j+1])
			}
			// 每个数组的最后的一个元素与最前面一个元素都为1
			new_nums = append(new_nums, 1)
			result = append(result, new_nums)
		}
	}
	fmt.Println(result)
	return result
}
