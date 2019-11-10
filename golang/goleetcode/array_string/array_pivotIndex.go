/*
	给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。

	我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。

	如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
 */

package main

import "fmt"

func main() {
	nums := []int{-1, -1, -1, -1, -1, -1}
	result := pivotIndex(nums)
	fmt.Printf("middle-index = %d", result)
}

func pivotIndex(nums []int) int {
	// 定义一个动态索引，每次和判断将以该索引指向的元素作为分界点
	pre := 0
	// 定义动态索引指向元素的左边所有元素之和，初始化为0
	fe_sum := 0
	// 定义动态索引指向的元素的右边所有元素之和，初始化为0
	be_sum := 0
	for ; pre < len(nums); pre++{
		if pre == 0{
			for i := 1; i < len(nums); i++{
				be_sum = be_sum + nums[i]
			}
		}
		if fe_sum == be_sum{
			return pre
		}else{
				fe_sum = fe_sum + nums[pre]
				if pre+ 1 < len(nums){
				be_sum = be_sum - nums[pre+1]
				}
		}

	}
	return -1
}