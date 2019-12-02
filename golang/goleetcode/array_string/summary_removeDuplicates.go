/*
	给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
	不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成
	不要考虑新数组超出新长度后面的元素
*/

package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return len(nums)
	}
	leftPoint := 0
	rightPoint := 0
	for ; rightPoint < len(nums) && leftPoint < len(nums)-1; rightPoint++ {
		if nums[rightPoint] != nums[leftPoint] {
			temp := nums[rightPoint]
			nums[rightPoint] = nums[leftPoint+1]
			nums[leftPoint+1] = temp
			leftPoint++
		}
	}
	return leftPoint + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length := removeDuplicates(nums)
	fmt.Println(length)
}
