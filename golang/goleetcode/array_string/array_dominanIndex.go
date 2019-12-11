/*
	在一个给定的数组nums中，总是存在一个最大元素 。
	查找数组中的最大元素是否至少是数组中每个其他数字的两倍。
	如果是，则返回最大元素的索引，否则返回-1。
 */

/*
    思路： 找出数组中的最大数，再找出剩余元素的最大数，判断数组最大数是否是剩余元素的最大数的两倍即可。
*/

package main

import "fmt"

// 测试用例
func main()  {
	slice1 := []int{18,66,24,29,57,54,84,16,29,71,22,33,98,34,20,65,21,6,11,
		74,7,93,40,64,97,56,3,72,63,67,72,86,42,29,34,46,97,34,45,82,22,80,94,45,46,96,10,34,98,55}
	fmt.Printf("result=%d\n",dominantIndex(slice1))
}

func dominantIndex(nums []int) int {
    // 切片为空时，直接返回-1
	if len(nums) == 0{
		return -1
	}
	// 声明切片最大数
	var max_num int
	// 声明切片第二大叔
	var lt_max_num int
	// 声明切片最大数索引
	var max_index int
	// 调用函数数获取数组的最大数和最大数索引
	max_index, max_num = get_max_num(nums)
	slice1 := nums[0:max_index]
	slice2 := nums[max_index+1:]
	// 拼接获取除最大数之外所有元素的切片
	slice3 := append(slice1, slice2...)
	// 调用函数获取切片第二大元素
	_, lt_max_num = get_max_num(slice3)
    // 判断切片第二大元素是否等于0
	if lt_max_num == 0{
	    // 判断最大元素是否等于0，若是，则返回-1，若不是返回最大数索引
		if max_num == 0{
			return -1
		}
		return max_index
	// 判断第二大元素是否等于最大数，若是，返回-1
	} else if lt_max_num == max_num{
		return -1
	// 判断第二大元素是否大于最大数的一半，若是，返回-1
	}else if max_num / lt_max_num < 2{
		return -1
	}else{
		return max_index
	}
}

// 获取切片最大数函数
func get_max_num(nums []int)(int, int){
	var max_index int
	var max_num int
	for index, num := range nums{
		if num > max_num{
			max_num = num
			max_index = index
		}
	}
	return max_index, max_num
}