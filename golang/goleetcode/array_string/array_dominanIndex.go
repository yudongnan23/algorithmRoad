/*
	在一个给定的数组nums中，总是存在一个最大元素 。

	查找数组中的最大元素是否至少是数组中每个其他数字的两倍。

	如果是，则返回最大元素的索引，否则返回-1。
 */

package main

import "fmt"

func main()  {
	slice1 := []int{18,66,24,29,57,54,84,16,29,71,22,33,98,34,20,65,21,6,11,
		74,7,93,40,64,97,56,3,72,63,67,72,86,42,29,34,46,97,34,45,82,22,80,94,45,46,96,10,34,98,55}
	fmt.Printf("result=%d\n",dominantIndex(slice1))
}

func dominantIndex(nums []int) int {
	if len(nums) == 0{
		return -1
	}
	var max_num int
	var lt_max_num int
	var max_index int
	max_index, max_num = get_max_num(nums)
	slice1 := nums[0:max_index]
	slice2 := nums[max_index+1:]
	slice3 := append(slice1, slice2...)
	_, lt_max_num = get_max_num(slice3)
	if lt_max_num == 0{
		if max_num == 0{
			return -1
		}
		return max_index
	} else if lt_max_num == max_num{
		return -1
	}else if max_num / lt_max_num < 2{
		return -1
	}else{
		return max_index
	}
}

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